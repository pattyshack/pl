package ast

import (
	"fmt"
)

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

func (expr BoolLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[BoolLiteralExpr: %s]", indent, label, expr.Value)
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

func (expr IntLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[IntLiteralExpr: %s]", indent, label, expr.Value)
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

func (expr FloatLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[FloatLiteralExpr: %s]", indent, label, expr.Value)
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

func (expr RuneLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[RuneLiteralExpr: %s]", indent, label, expr.Value)
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

func (expr StringLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[StringLiteralExpr: %s]", indent, label, expr.Value)
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

func (expr NamedExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[NamedExpr: %s]", indent, label, expr.Name)
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

func (expr AccessExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[AccessExpr: Field=%s\n", indent, label, expr.Field)
	result += expr.Operand.TreeString(indent+"  ", "Operand=")
	result += "\n" + indent + "]"
	return result
}

//
// UnaryExpr
//

type UnaryOp string

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

func (expr UnaryExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[UnaryExpr: IsPrefix=%v Op=(%s)\n",
		indent,
		label,
		expr.IsPrefix,
		expr.Op)
	result += expr.Operand.TreeString(indent+"  ", "Operand=")
	result += "\n" + indent + "]"
	return result
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

func (expr BinaryExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[BinaryExpr: Op=(%s)\n", indent, label, expr.Op)
	result += expr.Left.TreeString(indent+"  ", "Left=") + "\n"
	result += expr.Right.TreeString(indent+"  ", "Right=")
	result += "\n" + indent + "]"
	return result
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

func (expr *ImplicitStructExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[ImplicitStructExpr: IsImproper=%v\n",
		indent,
		label,
		expr.IsImproper)
	result += expr.Arguments.TreeString(indent+"  ", "Arguments=")
	result += "\n" + indent + "]\n"
	return result
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

func (expr ColonExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[ColonExpr\n", indent, label)
	result += expr.Arguments.TreeString(indent+"  ", "Arguments=")
	result += "\n" + indent + "]"
	return result
}

//
// CallExpr
//

type CallExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	FuncExpr         Expression
	GenericArguments *TypeExpressionList
	Arguments        *ArgumentList
}

var _ Expression = &CallExpr{}

func (expr *CallExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.FuncExpr.Walk(visitor)
	expr.GenericArguments.Walk(visitor)
	expr.Arguments.Walk(visitor)
	visitor.Exit(expr)
}

func (expr CallExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[CallExpr\n", indent, label)
	result += expr.FuncExpr.TreeString(indent+"  ", "FuncExpr=") + "\n"
	result += expr.GenericArguments.TreeString(
		indent+"  ",
		"GenericArguments=")
	result += expr.Arguments.TreeString(indent+"  ", "Arguments=")
	result += "\n" + indent + "]"
	return result
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

func (expr IndexExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[IndexExpr\n", indent, label)
	result += expr.Accessible.TreeString(indent+"  ", "Accessible=")
	result += "\n" + expr.Index.TreeString(indent+"  ", "Argument=")
	result += "\n" + indent + "]"
	return result
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

func (expr AsExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[AsExpr:\n", indent, label)
	result += expr.Accessible.TreeString(indent+"  ", "Accessible=") + "\n"
	result += expr.CastType.TreeString(indent+"  ", "CastType=") + "\n"
	result += "\n" + indent + "]"
	return result
}

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

func (expr InitializeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[InitializeExpr:\n", indent, label)
	result += expr.Initializable.TreeString(indent+"  ", "Initialiable=") + "\n"
	result += expr.Arguments.TreeString(indent+"  ", "Arguments=")
	result += "\n" + indent + "]"
	return result
}

//
// IfExpr
//

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
	if cb.Condition == nil {
		cb.Condition.Walk(visitor)
	}
	cb.Branch.Walk(visitor)
	visitor.Exit(cb)
}

func (cb ConditionBranch) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[ConditionBranch: IsElse=%v\n",
		indent,
		label,
		cb.IsElse)
	if cb.Condition != nil {
		result += cb.Condition.TreeString(indent+"  ", "Condition=") + "\n"
	}
	result += cb.Branch.TreeString(indent+"  ", "Branch=")
	result += "\n" + indent + "]"
	return result
}

type IfExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	LabelDecl         string // optional
	ConditionBranches *ConditionBranchList
}

var _ Expression = &IfExpr{}

func (expr *IfExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.ConditionBranches.Walk(visitor)
	visitor.Exit(expr)
}

func (expr IfExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[IfExpr: LabelDecl=%s\n",
		indent,
		label,
		expr.LabelDecl)
	expr.ConditionBranches.TreeString(indent+"  ", "Branches=")
	result += indent + "]"
	return result
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

func (expr SwitchExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[SwitchExpr:\n", indent, label)
	result += expr.Operand.TreeString(indent+"  ", "Operand=") + "\n"
	result += expr.Branches.TreeString(indent+"  ", "Branches=")
	result += "\n" + indent + "]"
	return result
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

func (expr SelectExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[SelectExpr:\n", indent, label)
	result += expr.Branches.TreeString(indent+"  ", "Branches=")
	result += "\n" + indent + "]"
	return result
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

func (expr LoopExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[LoopExpr: LoopKind=%s LabelDecl=%s",
		indent,
		label,
		expr.LoopKind,
		expr.LabelDecl)

	if expr.Init != nil {
		result += "\n" + expr.Init.TreeString(indent+"  ", "Init=")
	}

	result += "\n" + expr.Condition.TreeString(indent+"  ", "Condition=")

	if expr.Post != nil {
		result += "\n" + expr.Post.TreeString(indent+"  ", "Post=")
	}

	result += "\n" + expr.Body.TreeString(indent+"  ", "Body=")
	result += "\n" + indent + "]"
	return result
}
