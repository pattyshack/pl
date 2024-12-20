package source_passes

import (
	"fmt"

	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
)

type patternState int

const (
	// Real expressions
	nonPattern = patternState(0)

	// Statement level AddrDeclPattern not associated with AssignPattern or
	// CasePatterns.  Pattern body must either be a name or (nested) tuple of
	// name, without skip patterns, and be explicitly typed.
	addrDeclPattern = patternState(1)

	// AddrDeclPattern associated with AssignPattern or CasePatterns.  Pattern
	// body must either be a name or (nested) tuple of names with skip patterns.
	assignToNewAddrPattern = patternState(2)

	// AssignToAddrPattern associated with AssignPattern or CasePatterns.
	// Pattern body must either be a real expression or (nested) tuples of
	// real expressions with skip patterns.
	assignToExistingAddrPattern = patternState(3)

	// Pattern associated with AssignPattern or CasePatterns.  Pattern body
	// may either be a real expression, or AddrDeclPattern, or
	// AssignToAddrPattern, or a implicit struct pattern (i.e., (nested) tuples
	// of	real-expr/AddrDeclPattern/AssignToAddrPattern with skip patterns).
	// Note that AssignToAddrPattern is redundant; all real expr and
	// AssignToAddrPattern are existing variable addresses.
	addressPattern = patternState(4)

	// Pattern associated with AssignPattern or CasePatterns.  Pattern body
	// may either be a real expression, or AddrDeclPattern, or
	// AssignToAddrPattern, or a implicit struct pattern (i.e., (nested) tuples
	// of	real-expr/AddrDeclPattern/AssignToAddrPattern with skip patterns).
	matchPattern = patternState(5)
)

type PatternsAnalyzer struct {
	// pattern -> parent (either implicit struct expr or assign expr)
	redundantAssignToAddr map[*ast.AssignToAddrPattern]ast.Expression

	*parseutil.Emitter
}

func NewPatternsAnalyzer(emitter *parseutil.Emitter) *PatternsAnalyzer {
	return &PatternsAnalyzer{
		redundantAssignToAddr: map[*ast.AssignToAddrPattern]ast.Expression{},
		Emitter:               emitter,
	}
}

type patternsAnalyzePass struct {
	*PatternsAnalyzer
}

func (analyzer *patternsAnalyzePass) Process(list *ast.StatementList) {
	visitors := make(map[ast.Statement]*patternsAnalyzer, len(list.Elements))
	for _, stmt := range list.Elements {
		visitors[stmt] = &patternsAnalyzer{
			root:                      stmt,
			stateStack:                []patternState{},
			validAssignPatterns:       map[*ast.AssignPattern]bool{},
			validCasePatterns:         map[*ast.CasePatterns]struct{}{},
			validEnumPatterns:         map[*ast.EnumPattern]struct{}{},
			validAddrDeclPatterns:     map[*ast.AddrDeclPattern]patternState{},
			validAssignToAddrPatterns: map[*ast.AssignToAddrPattern]struct{}{},
			redundantAssignToAddr:     map[*ast.AssignToAddrPattern]ast.Expression{},
			Emitter:                   analyzer.Emitter,
		}
	}

	process.ParallelWalk(
		list,
		func(stmt ast.Statement) ast.Visitor { return visitors[stmt] })

	for _, visitor := range visitors {
		for pattern, parent := range visitor.redundantAssignToAddr {
			analyzer.redundantAssignToAddr[pattern] = parent
		}
	}
}

func (analyzer *PatternsAnalyzer) Validate() process.Pass {
	return &patternsAnalyzePass{analyzer}
}

type patternsTransformPass struct {
	*PatternsAnalyzer
}

func (analyzer *patternsTransformPass) Process(list *ast.StatementList) {
	for redundant, parent := range analyzer.redundantAssignToAddr {
		switch expr := parent.(type) {
		case *ast.ImplicitStructExpr:
			for _, arg := range expr.Arguments {
				if arg.Expr == redundant {
					arg.Expr = redundant.Pattern
				}
			}
		case *ast.AssignPattern:
			expr.Pattern = redundant.Pattern
		default:
			panic(fmt.Sprintf("unexpected parent expression: %v", parent))
		}
	}
}

func (analyzer *PatternsAnalyzer) Transform() process.Pass {
	return &patternsTransformPass{analyzer}
}

type patternsAnalyzer struct {
	root ast.Node

	stateStack []patternState

	// true for conditional assignment (must be paired with CasePatterns).
	// false for unconditional assignment.
	validAssignPatterns map[*ast.AssignPattern]bool

	validCasePatterns         map[*ast.CasePatterns]struct{}
	validEnumPatterns         map[*ast.EnumPattern]struct{}
	validAssignToAddrPatterns map[*ast.AssignToAddrPattern]struct{}
	// state is either addrDeclPattern or assignToNewAddrPattern
	validAddrDeclPatterns map[*ast.AddrDeclPattern]patternState

	// pattern -> parent (either implicit struct expr or assign expr)
	redundantAssignToAddr map[*ast.AssignToAddrPattern]ast.Expression

	*parseutil.Emitter
}

func (analyzer *patternsAnalyzer) push(state patternState) {
	analyzer.stateStack = append(analyzer.stateStack, state)
}

func (analyzer *patternsAnalyzer) pop() {
	analyzer.stateStack = analyzer.stateStack[:len(analyzer.stateStack)-1]
}

func (analyzer *patternsAnalyzer) current() patternState {
	state := nonPattern
	if len(analyzer.stateStack) > 0 {
		state = analyzer.stateStack[len(analyzer.stateStack)-1]
	}

	return state
}

func (analyzer *patternsAnalyzer) Enter(n ast.Node) {
	analyzer.collectPatternEntryPoints(n)

	e, ok := n.(ast.Expression)
	if !ok {
		return
	}

	switch expr := e.(type) {
	case *ast.NamedExpr:
		analyzer.push(nonPattern)

	case *ast.ImplicitStructExpr:
		current := analyzer.current()

		allowSkipPattern := current != nonPattern && current != addrDeclPattern
		skipPatternCount := 0
		for _, arg := range expr.Arguments {
			if arg.Kind == ast.SkipPatternArgument {
				skipPatternCount++
				if !allowSkipPattern {
					analyzer.Emit(arg.Loc(), "unexpected %s argument", arg.Kind)
				} else if skipPatternCount > 1 {
					analyzer.Emit(
						arg.Loc(),
						"each tuple can have at most one skip (...) pattern")
				}
				continue
			}

			if arg.Expr == nil {
				continue // error already emitted by nodeValidators
			}

			if current == addressPattern || current == matchPattern {
				switch argExpr := arg.Expr.(type) {
				case *ast.AddrDeclPattern:
					analyzer.validAddrDeclPatterns[argExpr] = assignToNewAddrPattern
				case *ast.AssignToAddrPattern:
					analyzer.validAssignToAddrPatterns[argExpr] = struct{}{}
					if current == addressPattern {
						analyzer.redundantAssignToAddr[argExpr] = expr
					}
				}
			}
		}

		analyzer.push(current)
	case *ast.AssignPattern:
		conditional, ok := analyzer.validAssignPatterns[expr]
		if !ok {
			analyzer.Emit(expr.Loc(), "unexpected assign pattern")
		}

		if conditional {
			casePatterns, ok := expr.Pattern.(*ast.CasePatterns)
			if ok {
				analyzer.validCasePatterns[casePatterns] = struct{}{}
			} else {
				analyzer.Emit(
					expr.Loc(),
					"invalid ast construction. "+
						"assign pattern must be paired case patterns")
			}
		} else {
			switch pattern := expr.Pattern.(type) {
			case *ast.AddrDeclPattern:
				analyzer.validAddrDeclPatterns[pattern] = assignToNewAddrPattern
			case *ast.AssignToAddrPattern:
				analyzer.validAssignToAddrPatterns[pattern] = struct{}{}
				analyzer.redundantAssignToAddr[pattern] = expr
			}
		}

		analyzer.push(addressPattern)
	case *ast.CasePatterns:
		_, ok := analyzer.validCasePatterns[expr]
		if !ok {
			analyzer.Emit(expr.Loc(), "unexpected case patterns")
		}

		for _, pattern := range expr.Patterns {
			enum, ok := pattern.(*ast.EnumPattern)
			if ok {
				analyzer.validEnumPatterns[enum] = struct{}{}
			}
		}

		analyzer.push(matchPattern)
	case *ast.AddrDeclPattern:
		state, ok := analyzer.validAddrDeclPatterns[expr]
		if !ok {
			analyzer.Emit(expr.Loc(), "unexpected address declaration pattern")
			state = assignToNewAddrPattern
		}

		analyzer.push(state)
	case *ast.AssignToAddrPattern:
		_, ok := analyzer.validAssignToAddrPatterns[expr]
		if !ok {
			analyzer.Emit(expr.Loc(), "unexpected assign to address (>) pattern")
		}

		analyzer.push(assignToExistingAddrPattern)
	case *ast.EnumPattern:
		_, ok := analyzer.validEnumPatterns[expr]
		if !ok {
			analyzer.Emit(expr.Loc(), "unexpected enum pattern")
		}

		analyzer.push(matchPattern)
	default: // Non pattern expressions
		state := analyzer.current()
		if state == addrDeclPattern || state == assignToNewAddrPattern {
			analyzer.Emit(
				e.Loc(),
				"unexpected expression, expecting name or tuple name pattern")
		}

		analyzer.push(nonPattern)
	}
}

func (analyzer *patternsAnalyzer) Exit(n ast.Node) {
	_, ok := n.(ast.Expression)
	if ok {
		analyzer.pop()
	}
}

func (analyzer *patternsAnalyzer) collectPatternEntryPoints(
	n ast.Node,
) {
	if n == analyzer.root {
		analyzer.processStatement(n.(ast.Statement))
	}

	switch node := n.(type) {
	case *ast.StatementsExpr:
		for _, stmt := range node.Statements {
			analyzer.processStatement(stmt)
		}
	case *ast.LoopExpr:
		if node.Condition != nil {
			assign, ok := node.Condition.(*ast.AssignPattern)
			if ok {
				analyzer.validAssignPatterns[assign] = false
			}
		}
	case *ast.IfExpr:
		for _, branch := range node.ConditionBranches {
			if branch.Condition == nil {
				continue
			}

			assign, ok := branch.Condition.(*ast.AssignPattern)
			if !ok {
				continue
			}
			analyzer.validAssignPatterns[assign] = true
		}
	case *ast.SwitchExpr:
		for _, branch := range node.ConditionBranches {
			if branch.Condition == nil {
				continue
			}

			pattern, ok := branch.Condition.(*ast.CasePatterns)
			if ok {
				analyzer.validCasePatterns[pattern] = struct{}{}
			}
		}
	case *ast.SelectExpr:
		for _, branch := range node.ConditionBranches {
			if branch.Condition == nil {
				continue
			}

			cond := branch.Condition
			assign, ok := branch.Condition.(*ast.AssignPattern)
			if ok {
				analyzer.validAssignPatterns[assign] = false
				cond = assign.Value
			}

			isValid := false
			switch expr := cond.(type) {
			case *ast.UnaryExpr:
				isValid = expr.Op == ast.UnaryRecvOp
			case *ast.BinaryExpr:
				isValid = expr.Op == ast.BinarySendOp
			}

			if !isValid {
				analyzer.Emit(
					cond.Loc(),
					"unexpected expression, expecting send/recv expression")
			}
		}
	}
}

func (analyzer *patternsAnalyzer) processStatement(
	s ast.Statement,
) {
	block, ok := s.(*ast.BlockAddrDeclStmt)
	if ok {
		for _, expr := range block.Patterns {
			switch pattern := expr.(type) {
			case *ast.AssignPattern:
				analyzer.validAssignPatterns[pattern] = false
			case *ast.AddrDeclPattern:
				analyzer.validAddrDeclPatterns[pattern] = addrDeclPattern
			}
		}
		return
	}

	assign, ok := s.(*ast.AssignPattern)
	if ok {
		if assign.Kind != ast.EqualAssign {
			analyzer.Emit(
				assign.Loc(),
				"invalid ast construction. unexpected assign pattern kind (%s)",
				assign.Kind)
		}
		analyzer.validAssignPatterns[assign] = false
		return
	}

	jmp, ok := s.(*ast.JumpStmt)
	if ok {
		s = jmp.Value
	}

	switch expr := s.(type) {
	case *ast.AddrDeclPattern:
		analyzer.validAddrDeclPatterns[expr] = addrDeclPattern
	case *ast.ImplicitStructExpr:
		// As a special case, we'll also allow variable declaration in improper
		// implicit struct's top level elements.
		if expr.Kind != ast.ImproperImplicitStruct {
			return
		}

		for _, arg := range expr.Arguments {
			if arg.Expr == nil {
				continue
			}

			decl, ok := arg.Expr.(*ast.AddrDeclPattern)
			if ok {
				analyzer.validAddrDeclPatterns[decl] = addrDeclPattern
			}
		}
	}
}
