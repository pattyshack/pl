package source_passes

import (
	"github.com/pattyshack/gt/parseutil"

	"github.com/pattyshack/pl/analyze/process"
	"github.com/pattyshack/pl/ast"
)

type DirectivesValidator struct {
	buildDirective *ast.Directive
	*parseutil.Emitter
}

func ValidateDirectives(emitter *parseutil.Emitter) process.Pass {
	return &DirectivesValidator{
		Emitter: emitter,
	}
}

func (validator *DirectivesValidator) Process(list *ast.StatementList) {
	if len(list.Elements) > 0 {
		decl, ok := list.Elements[0].(*ast.DirectivesDecl)
		if ok && len(decl.Directives) > 0 && decl.Directives[0].Name == "build" {
			validator.buildDirective = decl.Directives[0]
		}
	}

	process.ParallelWalk(
		list,
		func(ast.Statement) ast.Visitor { return validator })
}

func (validator *DirectivesValidator) Enter(n ast.Node) {
	switch node := n.(type) {
	case *ast.Directive:
		validator.validate(node)
	}
}

func (validator *DirectivesValidator) Exit(n ast.Node) {
}

func (validator *DirectivesValidator) validate(directive *ast.Directive) {
	switch directive.Name {
	case "build":
		if directive != validator.buildDirective {
			validator.Emit(
				directive.Loc(),
				"build directive must be the first directive declaration at the top "+
					"of the source file")
		}

		if directive.SubName != "" {
			validator.Emit(
				directive.Loc(),
				"unexpected build sub-directive (%s)",
				directive.SubName)
		}

		if len(directive.Arguments) != 1 {
			validator.Emit(
				directive.Loc(),
				"build directive expects one constraint argument")
		}
	default:
		validator.Emit(directive.Loc(), "unexpected directive (%s)", directive.Name)
	}
}
