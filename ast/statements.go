package ast

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
	Value Expression // optional
}

func (stmt *JumpStmt) Walk(visitor Visitor) {
	visitor.Enter(stmt)
	if stmt.Value != nil {
		stmt.Value.Walk(visitor)
	}
	visitor.Exit(stmt)
}

func NewJumpStmt(
	op TokenValue,
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
	IsTypeProp
	StartEndPos
	LeadingTrailingComments

	Language       string
	VerbatimSource string
}

var _ TypeProperty = &UnsafeStmt{}

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

	Condition Expression // nil when IsElse is true
	Branch    *StatementsExpr
}

func (cb *ConditionBranchStmt) Walk(visitor Visitor) {
	visitor.Enter(cb)
	if cb.Condition != nil {
		cb.Condition.Walk(visitor)
	}
	cb.Branch.Walk(visitor)
	visitor.Exit(cb)
}
