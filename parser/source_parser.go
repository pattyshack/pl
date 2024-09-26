package parser

import (
	"fmt"
	"io"

	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser/lr"
	reducerImpl "github.com/pattyshack/pl/parser/reducer"
)

type bufferLexer struct {
	*lexutil.BufferedReader[lr.Token]
	end lexutil.Location
}

func (l bufferLexer) CurrentLocation() lexutil.Location {
	peeked, err := l.Peek(1)
	if err != nil || len(peeked) == 0 {
		return l.end
	}

	return peeked[0].Loc()
}

// This splits the source token stream into definition segments, then parse
// each definition segments individually.  Each segment is terminated by either
// newlines or eof.  Note that newlines inside paren/brace scopes are
// ignored.  Parse errors are returned as part of the definition list.
type sourceParser struct {
	lexer lr.Lexer
	*reducerImpl.Reducer
	ParserOptions
}

func newSourceParser(
	fileName string,
	options ParserOptions,
) (
	*sourceParser,
	error,
) {
	reducer := reducerImpl.NewReducer()
	lexer, err := options.NewLexer(fileName, reducer)
	if err != nil {
		return nil, err
	}

	return &sourceParser{
		lexer:         lexer,
		Reducer:       reducer,
		ParserOptions: options,
	}, nil
}

func (parser *sourceParser) readDefinitionSegment() ([]lr.Token, error) {
	parenScope := 0
	braceScope := 0
	result := []lr.Token{}
	for {
		token, err := parser.lexer.Next()
		if err != nil {
			return result, err
		}

		switch token.Id() {
		case lr.NewlinesToken:
			if parenScope == 0 && braceScope == 0 {
				return result, nil
			}
		case lr.LparenToken:
			parenScope++
		case lr.RparenToken:
			parenScope--
		case lr.LbraceToken:
			braceScope++
		case lr.RbraceToken:
			braceScope--
		}

		result = append(result, token)

		if parenScope < 0 || braceScope < 0 {
			// mismatch paren/brace.  This will always result in an parse error.
			return result, nil
		}
	}
}

func (parser *sourceParser) _parseSource() (*ast.DefinitionList, error) {
	if parser.UseLRParseSource {
		return lr.ParseSource(parser.lexer, parser)
	}

	list := ast.NewDefinitionList()
	for {
		start := parser.lexer.CurrentLocation()
		segment, err := parser.readDefinitionSegment()
		end := parser.lexer.CurrentLocation()

		if len(segment) > 0 {
			def, err := lr.ParseDefinition(
				&bufferLexer{
					BufferedReader: lexutil.NewBufferedReaderFromSlice(segment),
					end:            end,
				},
				parser)
			if err != nil {
				parser.ParseErrors = append(parser.ParseErrors, err)
				def = ast.NewParseErrorNode(ast.NewStartEndPos(start, end), err)
			}

			if len(list.Elements) == 0 {
				list.Add(def)
			} else {
				list.ReduceAdd(&lr.TokenValue{}, def)
			}
		}

		if err != nil {
			if err == io.EOF {
				err = nil
			}

			return list, err
		}
	}
}

func (parser *sourceParser) groupCaseStatements(stmts *ast.StatementsExpr) {
	newStatements := []ast.Statement{}
	var current *ast.BranchStatement
	for _, stmt := range stmts.Statements {
		branch, ok := stmt.(*ast.BranchStatement)
		if !ok {
			if current != nil {
				current.Body.Statements = append(current.Body.Statements, stmt)
				current.EndPos = stmt.End()
			} else {
				err := fmt.Errorf(
					"statement does not belong to any branch: %s",
					stmt.Loc())
				parser.ParseErrors = append(parser.ParseErrors, err)
				newStatements = append(
					newStatements,
					ast.NewParseErrorNode(stmt.StartEnd(), err))
			}
			continue
		}

		current = branch
		newStatements = append(newStatements, branch)

		// Flatten nested case statements
		for {
			if len(current.Body.Statements) == 0 {
				break
			}

			if len(current.Body.Statements) != 1 {
				panic("should never happen")
			}

			nested, ok := current.Body.Statements[0].(*ast.BranchStatement)
			if !ok {
				break
			}

			current.EndPos = nested.Loc()
			current.Body.Statements = nil

			current = nested
			newStatements = append(newStatements, nested)
		}
	}

	stmts.Statements = newStatements
}

func (parser *sourceParser) pruneUnreachableStatements(
	stmts *ast.StatementsExpr,
) {
	for idx, stmt := range stmts.Statements {
		_, ok := stmt.(*ast.JumpStatement)
		if !ok {
			continue
		}

		if idx < len(stmts.Statements)-1 {
			err := fmt.Errorf(
				"unreachable statement: %s",
				stmts.Statements[idx+1].Loc())
			parser.ParseErrors = append(parser.ParseErrors, err)
			stmts.Statements[idx+1] = ast.NewParseErrorNode(
				stmts.Statements[idx+1].StartEnd(),
				err)
			stmts.Statements = stmts.Statements[:idx+2]
			break
		}
	}
}

func (parser *sourceParser) rejectUnexpectedStatements() {
	pkgBodies := map[*ast.StatementsExpr]struct{}{}
	for _, pkg := range parser.PackageDefs {
		pkgBodies[pkg.Body] = struct{}{}
	}

	branchBodies := map[*ast.StatementsExpr]struct{}{}
	for _, switchExpr := range parser.SwitchExprs {
		branchBodies[switchExpr.Branches] = struct{}{}
	}
	for _, selectExpr := range parser.SelectExprs {
		branchBodies[selectExpr.Branches] = struct{}{}
	}

	caseBodies := map[*ast.StatementsExpr]struct{}{}
	for stmts, _ := range branchBodies {
		for _, stmt := range stmts.Statements {
			switch branch := stmt.(type) {
			case *ast.BranchStatement:
				caseBodies[branch.Body] = struct{}{}
			}
		}
	}

	for _, stmts := range parser.StatementsExprs {
		for idx, stmt := range stmts.Statements {
			switch casted := stmt.(type) {
			case *ast.UnsafeStatement:
				// usable by all
			case *ast.ParseErrorNode:
				// usable by all
			case *ast.ImportStatement:
				_, ok := pkgBodies[stmts]
				if ok {
					continue
				}
				err := fmt.Errorf("unexpected import statement: %s", stmt.Loc())
				parser.ParseErrors = append(parser.ParseErrors, err)
				stmts.Statements[idx] = ast.NewParseErrorNode(stmt.StartEnd(), err)

			case *ast.BranchStatement:
				_, ok := branchBodies[stmts]
				if ok {
					continue
				}
				err := fmt.Errorf("unexpected branch statement: %s", stmt.Loc())
				parser.ParseErrors = append(parser.ParseErrors, err)
				stmts.Statements[idx] = ast.NewParseErrorNode(stmt.StartEnd(), err)

			case *ast.JumpStatement:
				_, isPkg := pkgBodies[stmts]
				_, isBranch := branchBodies[stmts]
				if isPkg || isBranch {
					err := fmt.Errorf("unexpected jump statement: %s", stmt.Loc())
					parser.ParseErrors = append(parser.ParseErrors, err)
					stmts.Statements[idx] = ast.NewParseErrorNode(stmt.StartEnd(), err)
					continue
				}

				if casted.Op == ast.FallthroughOp {
					_, ok := caseBodies[stmts]
					if !ok {
						err := fmt.Errorf(
							"unexpected fallthrough statement: %s",
							stmt.Loc())
						parser.ParseErrors = append(parser.ParseErrors, err)
						stmts.Statements[idx] = ast.NewParseErrorNode(stmt.StartEnd(), err)
					}
				}
			case ast.Expression:
				_, isPkg := pkgBodies[stmts]
				_, isBranch := branchBodies[stmts]
				if !isPkg && !isBranch {
					continue
				}
				err := fmt.Errorf("unexpected expression statement: %s", stmt.Loc())
				parser.ParseErrors = append(parser.ParseErrors, err)
				stmts.Statements[idx] = ast.NewParseErrorNode(stmt.StartEnd(), err)

			default:
				panic(fmt.Sprintf("unexpected statement type: %v", stmt))
			}
		}
	}
}

func (parser *sourceParser) rejectUnexpectedArguments() {
	for _, indexExpr := range parser.IndexExprs {
		switch indexExpr.Index.Kind {
		case ast.PositionalArgument, ast.ColonExprArgument:
			// ok
		default:
			err := fmt.Errorf(
				"unexpected %s argument: %s",
				indexExpr.Index.Kind,
				indexExpr.Index.Loc())
			parser.ParseErrors = append(parser.ParseErrors, err)
			indexExpr.Index = ast.NewPositionalArgument(
				ast.NewParseErrorNode(indexExpr.Index.StartEnd(), err))
		}
	}

	for _, callExpr := range parser.CallExprs {
		for idx, arg := range callExpr.Arguments {
			if arg.Kind == ast.SkipPatternArgument {
				err := fmt.Errorf("unexpected %s argument: %s", arg.Kind, arg.Loc())
				parser.ParseErrors = append(parser.ParseErrors, err)
				callExpr.Arguments[idx] = ast.NewPositionalArgument(
					ast.NewParseErrorNode(arg.StartEnd(), err))
			}
		}
	}

	for _, initExpr := range parser.InitializeExprs {
		for idx, arg := range initExpr.Arguments {
			if arg.Kind == ast.SkipPatternArgument ||
				arg.Kind == ast.VarargAssignmentArgument {

				err := fmt.Errorf("unexpected %s argument: %s", arg.Kind, arg.Loc())
				parser.ParseErrors = append(parser.ParseErrors, err)
				initExpr.Arguments[idx] = ast.NewPositionalArgument(
					ast.NewParseErrorNode(arg.StartEnd(), err))
			}
		}
	}

	// TODO: reject skip pattern in implicit struct if the struct is not a pattern
}

func (parser *sourceParser) analyze() {
	for _, expr := range parser.SwitchExprs {
		parser.groupCaseStatements(expr.Branches)
	}

	for _, expr := range parser.SelectExprs {
		parser.groupCaseStatements(expr.Branches)
	}

	for _, stmts := range parser.StatementsExprs {
		parser.pruneUnreachableStatements(stmts)
	}

	parser.rejectUnexpectedStatements()
	parser.rejectUnexpectedArguments()
}

func (parser *sourceParser) parseSource() (
	*reducerImpl.Reducer,
	*ast.DefinitionList,
	error,
) {
	source, err := parser._parseSource()
	if err != nil {
		return nil, nil, err
	}

	parser.analyze()
	return parser.Reducer, source, nil
}

func (parser *sourceParser) parseExpr() (
	*reducerImpl.Reducer,
	ast.Expression,
	error,
) {
	expr, err := lr.ParseExpr(parser.lexer, parser)
	if err != nil {
		return nil, nil, err
	}

	parser.analyze()
	return parser.Reducer, expr, nil
}

func (parser *sourceParser) parseTypeExpr() (
	*reducerImpl.Reducer,
	ast.TypeExpression,
	error,
) {
	typeExpr, err := lr.ParseTypeExpr(parser.lexer, parser)
	if err != nil {
		return nil, nil, err
	}

	parser.analyze()
	return parser.Reducer, typeExpr, nil
}

func (parser *sourceParser) parseStatement() (
	*reducerImpl.Reducer,
	ast.Statement,
	error,
) {
	stmt, err := lr.ParseStatement(parser.lexer, parser)
	if err != nil {
		return nil, nil, err
	}

	parser.analyze()
	return parser.Reducer, stmt, nil
}

func (parser *sourceParser) parseDefinition() (
	*reducerImpl.Reducer,
	ast.Definition,
	error,
) {
	def, err := lr.ParseDefinition(parser.lexer, parser)
	if err != nil {
		return nil, nil, err
	}

	parser.analyze()
	return parser.Reducer, def, nil
}

func ParseExpr(
	fileName string,
	options ParserOptions,
) (
	*reducerImpl.Reducer,
	ast.Expression,
	error,
) {
	parser, err := newSourceParser(fileName, options)
	if err != nil {
		return nil, nil, err
	}

	return parser.parseExpr()
}

func ParseTypeExpr(
	fileName string,
	options ParserOptions,
) (
	*reducerImpl.Reducer,
	ast.TypeExpression,
	error,
) {
	parser, err := newSourceParser(fileName, options)
	if err != nil {
		return nil, nil, err
	}

	return parser.parseTypeExpr()
}

func ParseStatement(
	fileName string,
	options ParserOptions,
) (
	*reducerImpl.Reducer,
	ast.Statement,
	error,
) {
	parser, err := newSourceParser(fileName, options)
	if err != nil {
		return nil, nil, err
	}

	return parser.parseStatement()
}

func ParseDefinition(
	fileName string,
	options ParserOptions,
) (
	*reducerImpl.Reducer,
	ast.Definition,
	error,
) {
	parser, err := newSourceParser(fileName, options)
	if err != nil {
		return nil, nil, err
	}

	return parser.parseDefinition()
}

func ParseSource(
	fileName string,
	options ParserOptions,
) (
	*reducerImpl.Reducer,
	*ast.DefinitionList,
	error,
) {
	parser, err := newSourceParser(fileName, options)
	if err != nil {
		return nil, nil, err
	}

	return parser.parseSource()
}
