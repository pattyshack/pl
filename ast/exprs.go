package ast

//
// (Bool/Int/Float/Rune/String)LiteralExpr
//

type BoolLiteralExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Value string
}

func (expr *BoolLiteralExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	visitor.Exit(expr)
}

type IntLiteralExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Value string
}

func (expr *IntLiteralExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	visitor.Exit(expr)
}

type FloatLiteralExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Value string
}

func (expr *FloatLiteralExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	visitor.Exit(expr)
}

type RuneLiteralExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Value string
}

func (expr *RuneLiteralExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	visitor.Exit(expr)
}

type StringLiteralExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Value string
}

func (expr *StringLiteralExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	visitor.Exit(expr)
}

//
// NamedExpr
//

type NamedExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Name string
}

func (expr *NamedExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	visitor.Exit(expr)
}

//
// AccessExpr
//

type AccessExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Operand Expression
	Field   string
}

func (expr *AccessExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Operand.Walk(visitor)
	visitor.Exit(expr)
}

//
// UnaryExpr
//

type UnaryOp string

const (
	UnaryReturnDefaultEnumOp  = UnaryOp("?")
	UnaryPanicOnDefaultEnumOp = UnaryOp("!")
	UnaryAssignAddOneOp       = UnaryOp("++")
	UnaryAssignSubOneOp       = UnaryOp("--")
	UnaryNotOp                = UnaryOp("not")
	UnaryBitwiseComplementOp  = UnaryOp("^")
	UnaryPlusOp               = UnaryOp("+")
	UnaryMinusOp              = UnaryOp("-")
	UnaryDerefOp              = UnaryOp("*")
	UnaryRefOp                = UnaryOp("&")
	UnaryAsyncOp              = UnaryOp("async")
	UnaryDeferOp              = UnaryOp("defer")
	UnaryAssignToOp           = UnaryOp(">")
	UnaryRecvOp               = UnaryOp("<-")
)

type UnaryExpr struct {
	IsExpr

	StartEndPos
	LeadingTrailingComments

	IsPrefix bool

	Op      UnaryOp
	Operand Expression
}

func (expr *UnaryExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Operand.Walk(visitor)
	visitor.Exit(expr)
}

//
// BinaryExpr
//

type BinaryOp string

const (
	BinaryAddOp             = BinaryOp("+")
	BinarySubOp             = BinaryOp("-")
	BinaryMulOp             = BinaryOp("*")
	BinaryDivOp             = BinaryOp("/")
	BinaryModOp             = BinaryOp("%")
	BinaryBitAndOp          = BinaryOp("&")
	BinaryBitOrOp           = BinaryOp("|")
	BinaryBitXorOp          = BinaryOp("^")
	BinaryBitLshiftOp       = BinaryOp("<<")
	BinaryBitRshiftOp       = BinaryOp(">>")
	BinaryEqualOp           = BinaryOp("==")
	BinaryNotEqualOp        = BinaryOp("!=")
	BinaryLessOp            = BinaryOp("<")
	BinaryLessOrEqualOp     = BinaryOp("<=")
	BinaryGreaterOp         = BinaryOp(">")
	BinaryGreaterOrEqualOp  = BinaryOp(">=")
	BinaryAndOp             = BinaryOp("and")
	BinaryOrOp              = BinaryOp("or")
	BinarySendOp            = BinaryOp("<-")
	BinaryAssignOp          = BinaryOp("=")
	BinaryInOp              = BinaryOp("in")
	BinaryAddAssignOp       = BinaryOp("+=")
	BinarySubAssignOp       = BinaryOp("-=")
	BinaryMulAssignOp       = BinaryOp("*=")
	BinaryDivAssignOp       = BinaryOp("/=")
	BinaryModAssignOp       = BinaryOp("%=")
	BinaryBitAndAssignOp    = BinaryOp("&=")
	BinaryBitOrAssignOp     = BinaryOp("|=")
	BinaryBitXorAssignOp    = BinaryOp("^=")
	BinaryBitLshiftAssignOp = BinaryOp("<<=")
	BinaryBitRshiftAssignOp = BinaryOp(">>=")
)

type BinaryExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Left  Expression
	Op    BinaryOp
	Right Expression
}

func (expr *BinaryExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Left.Walk(visitor)
	expr.Right.Walk(visitor)
	visitor.Exit(expr)
}

//
// ImplicitStructExpr
//

type ImplicitStructExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Arguments *ArgumentList

	// An improper struct is the a comma separated list of expressions without
	// left/right paren.  e.g., return 1, 2, 3
	IsImproper bool
}

func (expr *ImplicitStructExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Arguments.Walk(visitor)
	visitor.Exit(expr)
}

//
// ColonExpr
//

type ColonExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Arguments *ArgumentList
}

var _ Expression = &ColonExpr{}

func (expr *ColonExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Arguments.Walk(visitor)
	visitor.Exit(expr)
}

//
// CallExpr
//

type CallExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	FuncExpr         Expression
	GenericArguments *TypeExpressionList // optional
	Arguments        *ArgumentList
}

var _ Expression = &CallExpr{}

func (expr *CallExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.FuncExpr.Walk(visitor)
	if expr.GenericArguments != nil {
		expr.GenericArguments.Walk(visitor)
	}
	expr.Arguments.Walk(visitor)
	visitor.Exit(expr)
}

//
// IndexExpr
//

type IndexExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Accessible Expression
	Index      *Argument
}

func (expr *IndexExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Accessible.Walk(visitor)
	expr.Index.Walk(visitor)
	visitor.Exit(expr)
}

//
// AsExpr
//

type AsExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Accessible Expression
	CastType   TypeExpression
}

func (expr *AsExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Accessible.Walk(visitor)
	expr.CastType.Walk(visitor)
	visitor.Exit(expr)
}

var _ Expression = &AsExpr{}

//
// InitializeExpr
//

type InitializeExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Initializable TypeExpression
	Arguments     *ArgumentList
}

func (expr *InitializeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Initializable.Walk(visitor)
	expr.Arguments.Walk(visitor)
	visitor.Exit(expr)
}

//
// IfExpr
//

type CaseConditionExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	SwitchablePatterns *ExpressionList
	Value              Expression
}

var _ Expression = &CaseConditionExpr{}

func (cond *CaseConditionExpr) Walk(visitor Visitor) {
	visitor.Enter(cond)
	cond.SwitchablePatterns.Walk(visitor)
	cond.Value.Walk(visitor)
	visitor.Exit(cond)
}

type ConditionBranch struct {
	StartEndPos
	LeadingTrailingComments

	IsElse    bool
	Condition Expression // nil when IsElse is true
	Branch    Expression
}

var _ Node = &ConditionBranch{}

func (cb *ConditionBranch) Walk(visitor Visitor) {
	visitor.Enter(cb)
	if cb.Condition != nil {
		cb.Condition.Walk(visitor)
	}
	cb.Branch.Walk(visitor)
	visitor.Exit(cb)
}

type IfExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	LabelDecl         string // optional
	ConditionBranches []*ConditionBranch
}

var _ Expression = &IfExpr{}

func (expr *IfExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	for _, condBranch := range expr.ConditionBranches {
		condBranch.Walk(visitor)
	}
	visitor.Exit(expr)
}

//
// SwitchExpr
//

type SwitchExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	LabelDecl string // optional
	Operand   Expression
	Branches  *StatementsExpr
}

var _ Expression = &SwitchExpr{}

func (expr *SwitchExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Operand.Walk(visitor)
	expr.Branches.Walk(visitor)
	visitor.Exit(expr)
}

//
// SelectExpr
//

type SelectExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	LabelDecl string // optional
	Branches  *StatementsExpr
}

var _ Expression = &SelectExpr{}

func (expr *SelectExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Branches.Walk(visitor)
	visitor.Exit(expr)
}

//
// LoopExpr
//

type LoopKind string

const (
	InfiniteLoop = LoopKind("infinite")
	DoWhileLoop  = LoopKind("do-while")
	WhileLoop    = LoopKind("while")
	IteratorLoop = LoopKind("iterator")
	ForLoop      = LoopKind("for")
)

type LoopExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	LoopKind

	LabelDecl string

	Init      Statement  // optional. only applicable to traditional-for loop
	Condition Expression // optional. not applicable to infinite loop
	Post      Statement  // optional. only applicable to traditional-for loop

	Body *StatementsExpr
}

var _ Expression = &LoopExpr{}

func (expr *LoopExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	if expr.Init != nil {
		expr.Init.Walk(visitor)
	}
	if expr.Condition != nil {
		expr.Condition.Walk(visitor)
	}
	if expr.Post != nil {
		expr.Post.Walk(visitor)
	}
	expr.Body.Walk(visitor)
	visitor.Exit(expr)
}
