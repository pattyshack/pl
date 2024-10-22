package ast

//
// DirectivesDecl
//

type DirectivesDecl struct {
	IsStmt
	StartEndPos
	LeadingTrailingComments

	Directives []*Directive
}

var _ Statement = &DirectivesDecl{}

func (decl *DirectivesDecl) Walk(visitor Visitor) {
	visitor.Enter(decl)
	for _, dir := range decl.Directives {
		dir.Walk(visitor)
	}
	visitor.Exit(decl)
}

//
// Directive
//

type Directive struct {
	StartEndPos
	LeadingTrailingComments

	Name    string
	SubName string

	Arguments []DirectiveExpression
}

var _ Node = &Directive{}

func (dir *Directive) Walk(visitor Visitor) {
	visitor.Enter(dir)
	for _, arg := range dir.Arguments {
		arg.Walk(visitor)
	}
	visitor.Exit(dir)
}

//
// Directive expressions
//

type DirectiveValueExpr struct {
	IsDirectiveExpr
	StartEndPos
	LeadingTrailingComments

	Value string
}

var _ DirectiveExpression = &DirectiveValueExpr{}

func (expr *DirectiveValueExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	visitor.Exit(expr)
}

type DirectiveGroupExpr struct {
	IsDirectiveExpr
	StartEndPos
	LeadingTrailingComments

	Expr DirectiveExpression
}

var _ DirectiveExpression = &DirectiveGroupExpr{}

func (expr *DirectiveGroupExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Expr.Walk(visitor)
	visitor.Exit(expr)
}

type DirectiveNotExpr struct {
	IsDirectiveExpr
	StartEndPos
	LeadingTrailingComments

	Operand DirectiveExpression
}

var _ DirectiveExpression = &DirectiveNotExpr{}

func (expr *DirectiveNotExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Operand.Walk(visitor)
	visitor.Exit(expr)
}

type DirectiveAndExpr struct {
	IsDirectiveExpr
	StartEndPos
	LeadingTrailingComments

	Left  DirectiveExpression
	Right DirectiveExpression
}

var _ DirectiveExpression = &DirectiveAndExpr{}

func (expr *DirectiveAndExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Left.Walk(visitor)
	expr.Right.Walk(visitor)
	visitor.Exit(expr)
}

type DirectiveOrExpr struct {
	IsDirectiveExpr
	StartEndPos
	LeadingTrailingComments

	Left  DirectiveExpression
	Right DirectiveExpression
}

var _ DirectiveExpression = &DirectiveOrExpr{}

func (expr *DirectiveOrExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Left.Walk(visitor)
	expr.Right.Walk(visitor)
	visitor.Exit(expr)
}
