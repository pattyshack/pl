package ast

import (
	"fmt"
)

//
// (Bool/Int/Float/Rune/String)LiteralExpr
//

type BoolLiteralExpr struct {
	IsExpr
	ValuedNode
}

func (expr BoolLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[BoolLiteralExpr: %s]", indent, label, expr.Val())
}

type IntLiteralExpr struct {
	IsExpr
	ValuedNode
}

func (expr IntLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[IntLiteralExpr: %s]", indent, label, expr.Val())
}

type FloatLiteralExpr struct {
	IsExpr
	ValuedNode
}

func (expr FloatLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[FloatLiteralExpr: %s]", indent, label, expr.Val())
}

type RuneLiteralExpr struct {
	IsExpr
	ValuedNode
}

func (expr RuneLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[RuneLiteralExpr: %s]", indent, label, expr.Val())
}

type StringLiteralExpr struct {
	IsExpr
	ValuedNode
}

func (expr StringLiteralExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[StringLiteralExpr: %s]", indent, label, expr.Val())
}

//
// NamedExpr
//

type NamedExpr struct {
	IsExpr
	ValuedNode
}

func (expr NamedExpr) TreeString(indent string, label string) string {
	return fmt.Sprintf("%s%s[NamedExpr: %s]", indent, label, expr.Val())
}

//
// AccessExpr
//

type AccessExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Operand Expression
	Field   ValuedNode
}

func (expr AccessExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[AccessExpr: Field=%s\n", indent, label, expr.Field.Val())
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

type BinaryExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Left  Expression
	Op    BinaryOp
	Right Expression
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

	Arguments []*Argument

	// An improper struct is the a comma separated list of expressions without
	// left/right paren.  e.g., return 1, 2, 3
	IsImproper bool
}

func (expr *ImplicitStructExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[ImplicitStructExpr: IsImproper=%v\n",
		indent,
		label,
		expr.IsImproper)
	if len(expr.Arguments) == 0 {
		return result + "]"
	}

	result += elementsTreeString(expr.Arguments, indent+"  ", "Argument")
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

	Arguments []*Argument
}

var _ Expression = &ColonExpr{}

func (expr ColonExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[ColonExpr\n", indent, label)
	result += ListTreeString(
		expr.Arguments,
		indent+"  ",
		"Arguments=",
		"Argument")
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
	GenericArguments []TypeExpression
	Arguments        []*Argument
}

var _ Expression = &CallExpr{}

func (expr CallExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[CallExpr\n", indent, label)
	result += expr.FuncExpr.TreeString(indent+"  ", "FuncExpr=") + "\n"
	if len(expr.GenericArguments) > 0 {
		result += ListTreeString(
			expr.GenericArguments,
			indent+"  ",
			"GenericArguments=",
			"GenericArgument")
		result += "\n"
	}
	result += ListTreeString(
		expr.Arguments,
		indent+"  ",
		"Arguments=",
		"Argument")
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
	Index      Argument
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
	Arguments     []*Argument
}

func (expr InitializeExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[InitializeExpr:\n", indent, label)
	result += expr.Initializable.TreeString(indent+"  ", "Initialiable=") + "\n"
	result += ListTreeString(
		expr.Arguments,
		indent+"  ",
		"Arguments=",
		"Argument")
	result += "\n" + indent + "]"
	return result
}

//
// IfExpr
//

type ConditionBranch struct {
	Condition Expression
	Branch    Expression
}

func (cb ConditionBranch) TreeString(indent string, label string) string {
	result := fmt.Sprintf("%s%s[ConditionBranch:\n", indent, label)
	result += cb.Condition.TreeString(indent+"  ", "Condition=") + "\n"
	result += cb.Branch.TreeString(indent+"  ", "Branch=")
	result += "\n" + indent + "]"
	return result
}

type IfExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	LabelDecl         string // optional
	ConditionBranches []ConditionBranch
	ElseBranch        Expression // optional
}

var _ Expression = &IfExpr{}

func (expr IfExpr) TreeString(indent string, label string) string {
	result := fmt.Sprintf(
		"%s%s[IfExpr: LabelDecl=%s\n",
		indent,
		label,
		expr.LabelDecl)
	for idx, condBranch := range expr.ConditionBranches {
		branchLabel := fmt.Sprintf("Branch%d=", idx)
		result += condBranch.TreeString(indent+"  ", branchLabel) + "\n"
	}

	if expr.ElseBranch != nil {
		result += expr.ElseBranch.TreeString(indent+"  ", "ElseBranch=") + "\n"
	}

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
