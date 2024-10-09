package ast

import (
	"github.com/pattyshack/gt/lexutil"
)

//
// ImportClause
//

type ImportClause struct {
	StartEndPos
	LeadingTrailingComments

	Alias   string // Identifier or underscore or dot or ""
	Package string
}

var _ Node = &ImportClause{}

func (clause *ImportClause) Walk(visitor Visitor) {
	visitor.Enter(clause)
	visitor.Exit(clause)
}

//
// ImportStmt
//

type ImportStmt struct {
	IsStmt
	StartEndPos
	LeadingTrailingComments

	ImportClauses []*ImportClause
}

var _ Statement = &ImportStmt{}

func (stmt *ImportStmt) Walk(visitor Visitor) {
	visitor.Enter(stmt)
	for _, clause := range stmt.ImportClauses {
		clause.Walk(visitor)
	}
	visitor.Exit(stmt)
}

//
// JumpStmt
//

type JumpOp string

const (
	FallthroughOp = JumpOp("fallthrough")
	ReturnOp      = JumpOp("return")
	ContinueOp    = JumpOp("continue")
	BreakOp       = JumpOp("break")
)

type JumpStmt struct {
	IsStmt
	StartEndPos
	LeadingTrailingComments

	Op    JumpOp
	Label string     // optional
	Value Expression // use improper unit struct for no value
}

var _ Statement = &JumpStmt{}
var _ Validator = &JumpStmt{}

func (stmt *JumpStmt) Walk(visitor Visitor) {
	visitor.Enter(stmt)
	stmt.Value.Walk(visitor)
	visitor.Exit(stmt)
}

func (stmt *JumpStmt) Validate(emitter *lexutil.ErrorEmitter) {
	switch stmt.Op {
	case FallthroughOp, ReturnOp, ContinueOp, BreakOp: // ok
	default:
		emitter.Emit(
			stmt.Loc(),
			"invalid ast construction. unexpected jump statement kind (%s)",
			stmt.Op)
	}
}

func NewJumpStmt(
	op TokenValue,
	atToken TokenValue,
	labelToken TokenValue,
	value Expression,
) *JumpStmt {
	start := op.Loc()
	end := op.End()
	leading := op.TakeLeading()
	trailing := op.TakeTrailing()

	label := ""
	if labelToken != nil {
		label = labelToken.Val()
		end = labelToken.End()
		trailing.Append(atToken.TakeLeading())
		trailing.Append(atToken.TakeTrailing())
		trailing.Append(labelToken.TakeLeading())
		trailing.Append(labelToken.TakeTrailing())
	}

	if value != nil {
		value.PrependToLeading(trailing)
		end = value.End()
		trailing = value.TakeTrailing()
	}

	stmt := &JumpStmt{
		StartEndPos: NewStartEndPos(start, end),
		Op:          JumpOp(op.Val()),
		Label:       label,
		Value:       value,
	}
	stmt.LeadingComment = leading
	stmt.TrailingComment = trailing

	return stmt
}

//
// UnsafeStmt
//

type UnsafeStmt struct {
	IsStmt
	StartEndPos
	LeadingTrailingComments

	Language       string
	VerbatimSource string
}

var _ Statement = &UnsafeStmt{}

func (stmt *UnsafeStmt) Walk(visitor Visitor) {
	visitor.Enter(stmt)
	visitor.Exit(stmt)
}

//
// ConditionBranchStmt
//

type ConditionBranchStmt struct {
	IsStmt
	StartEndPos
	LeadingTrailingComments

	// either default branch in SwitchExpr/SelectExpr, or else branch in IfExpr
	IsDefaultBranch bool

	Condition Expression // nil when IsDefaultBranch is true
	Branch    *StatementsExpr
}

var _ Statement = &ConditionBranchStmt{}

func (cb *ConditionBranchStmt) Walk(visitor Visitor) {
	visitor.Enter(cb)
	if cb.Condition != nil {
		cb.Condition.Walk(visitor)
	}
	cb.Branch.Walk(visitor)
	visitor.Exit(cb)
}

//
// FloatingComment
//

type FloatingComment struct {
	IsStmt
	StartEndPos
	LeadingTrailingComments
}

var _ Statement = &FloatingComment{}

func (def *FloatingComment) Walk(visitor Visitor) {
	visitor.Enter(def)
	visitor.Exit(def)
}

//
// TypeDef
//

type TypeDef struct {
	IsStmt
	StartEndPos
	LeadingTrailingComments

	Name              string
	GenericParameters *GenericParameterList // optional
	BaseType          TypeExpression
	Constraint        TypeExpression // use any trait if unconstrained
}

var _ Statement = &TypeDef{}

func (def *TypeDef) Walk(visitor Visitor) {
	visitor.Enter(def)
	if def.GenericParameters != nil {
		def.GenericParameters.Walk(visitor)
	}
	def.BaseType.Walk(visitor)
	def.Constraint.Walk(visitor)
	visitor.Exit(def)
}

//
// BlockAddrDeclStmt
//

type BlockAddrDeclStmt struct {
	IsStmt
	StartEndPos
	LeadingTrailingComments

	IsVar bool // true = var, false = let

	Patterns []Expression // either AddrDeclPattern or AssignPattern
}

var _ Statement = &BlockAddrDeclStmt{}
var _ Validator = &BlockAddrDeclStmt{}

func (block *BlockAddrDeclStmt) Walk(visitor Visitor) {
	visitor.Enter(block)
	for _, pattern := range block.Patterns {
		pattern.Walk(visitor)
	}
	visitor.Exit(block)
}

func (block *BlockAddrDeclStmt) Validate(emitter *lexutil.ErrorEmitter) {
	for _, expr := range block.Patterns {
		switch pattern := expr.(type) {
		case *AssignPattern:
			if pattern.Kind != EqualAssign {
				emitter.Emit(
					pattern.Loc(),
					"invalid ast construction, unexpected assign pattern kind (%s)",
					pattern.Kind)
			}

			decl, ok := pattern.Pattern.(*AddrDeclPattern)
			if !ok {
				emitter.Emit(
					pattern.Pattern.Loc(),
					"invalid ast construction, expected addr decl in assign pattern")
			} else if block.IsVar != decl.IsVar {
				emitter.Emit(
					decl.Loc(),
					"invalid ast construction, addr decl type does not much block")
			}
		case *AddrDeclPattern:
			if block.IsVar != pattern.IsVar {
				emitter.Emit(
					pattern.Loc(),
					"invalid ast construction, addr decl type does not much block")
			}
		default:
			emitter.Emit(
				expr.Loc(),
				"invalid ast construction, expected assign or addr decl pattern")
		}
	}
}
