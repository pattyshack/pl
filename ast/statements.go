package ast

//
// StatementsExpr
//

type StatementsExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	LabelDecl  string // optional
	Statements []Statement
}

func (expr *StatementsExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	for _, stmt := range expr.Statements {
		stmt.Walk(visitor)
	}
	visitor.Exit(expr)
}

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
// ImportStatement
//

type ImportStatement struct {
	IsStmt
	StartEndPos
	LeadingTrailingComments

	ImportClauses []*ImportClause
}

var _ Statement = &ImportStatement{}

func (stmt *ImportStatement) Walk(visitor Visitor) {
	visitor.Enter(stmt)
	for _, clause := range stmt.ImportClauses {
		clause.Walk(visitor)
	}
	visitor.Exit(stmt)
}

//
// JumpStatement
//

type JumpOp string

const (
	FallthroughOp = JumpOp("fallthrough")
	ReturnOp      = JumpOp("return")
	ContinueOp    = JumpOp("continue")
	BreakOp       = JumpOp("break")
)

type JumpStatement struct {
	IsStmt
	StartEndPos
	LeadingTrailingComments

	Op    JumpOp
	Label string     // optional
	Value Expression // optional
}

var _ Statement = &JumpStatement{}

func (stmt *JumpStatement) Walk(visitor Visitor) {
	visitor.Enter(stmt)
	if stmt.Value != nil {
		stmt.Value.Walk(visitor)
	}
	visitor.Exit(stmt)
}

func NewJumpStatement(
	op TokenValue,
	labelToken TokenValue,
	value Expression,
) *JumpStatement {
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

	stmt := &JumpStatement{
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
// UnsafeStatement
//

type UnsafeStatement struct {
	IsStmt
	IsTypeProp
	StartEndPos
	LeadingTrailingComments

	Language       string
	VerbatimSource string
}

var _ Statement = &UnsafeStatement{}
var _ TypeProperty = &UnsafeStatement{}

func (stmt *UnsafeStatement) Walk(visitor Visitor) {
	visitor.Enter(stmt)
	visitor.Exit(stmt)
}

//
// BranchStatement
//

type BranchStatement struct {
	IsStmt
	StartEndPos
	LeadingTrailingComments

	IsDefault    bool
	CasePatterns *ExpressionList // optional
	Body         *StatementsExpr
}

var _ Statement = &BranchStatement{}

func (stmt *BranchStatement) Walk(visitor Visitor) {
	visitor.Enter(stmt)
	if stmt.CasePatterns != nil {
		stmt.CasePatterns.Walk(visitor)
	}
	stmt.Body.Walk(visitor)
	visitor.Exit(stmt)
}
