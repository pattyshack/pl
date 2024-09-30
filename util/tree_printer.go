package util

import (
	"fmt"
	"io"
	"strings"

	"github.com/pattyshack/pl/ast"
)

const (
	indent = "  "
)

func PrintTree(output io.Writer, node ast.Node, indent string) error {
	printer := &treePrinter{
		indent:     indent,
		labelStack: []string{},
		writer:     output,
	}
	node.Walk(printer)
	return printer.err
}

type treePrinter struct {
	indent     string
	labelStack []string
	writer     io.Writer
	err        error
}

func (printer *treePrinter) write(format string, args ...interface{}) {
	if printer.err != nil {
		return
	}

	if len(args) == 0 {
		_, printer.err = printer.writer.Write([]byte(format))
	} else {
		_, printer.err = fmt.Fprintf(printer.writer, format, args...)
	}
}

func (printer *treePrinter) writeLabel() {
	label := ""
	if len(printer.labelStack) > 0 {
		label = printer.labelStack[len(printer.labelStack)-1]
		printer.labelStack = printer.labelStack[:len(printer.labelStack)-1]
	}

	if len(label) > 0 {
		printer.write("\n")
		printer.write(printer.indent)
		printer.write(label)
	} else {
		printer.write(printer.indent)
	}
}

func (printer *treePrinter) endNode() {
	printer.indent = printer.indent[:len(printer.indent)-len(indent)]
	printer.write("\n")
	printer.write(printer.indent)
	printer.write("]")
}

func (printer *treePrinter) push(labels ...string) {
	printer.indent += indent

	for len(labels) > 0 {
		last := labels[len(labels)-1]
		labels = labels[:len(labels)-1]

		printer.labelStack = append(printer.labelStack, last)
	}
}

func (printer *treePrinter) list(
	header string,
	argLabels []string,
	elementType string,
	size int,
) {
	printer.write(header)
	if size == 0 && len(argLabels) == 0 {
		printer.write("]")
	} else {
		printer.push(argLabels...)

		for i := size - 1; i >= 0; i-- {
			printer.labelStack = append(
				printer.labelStack,
				fmt.Sprintf("%s%d=", elementType, i))
		}
	}
}

func (printer *treePrinter) endList(size int) {
	if size > 0 {
		printer.endNode()
	}
}

func (printer *treePrinter) Enter(n ast.Node) {
	printer.writeLabel()

	switch node := n.(type) {
	case *ast.ParseErrorExpr:
		printer.write("[ParseErrorExpr: %s %s]", node.Error, node.StartPos)
	case *ast.BoolLiteralExpr:
		printer.write("[BoolLiteralExpr: %s]", node.Value)
	case *ast.IntLiteralExpr:
		printer.write("[IntLiteralExpr: %s]", node.Value)
	case *ast.FloatLiteralExpr:
		printer.write("[FloatLiteralExpr: %s]", node.Value)
	case *ast.RuneLiteralExpr:
		printer.write("[RuneLiteralExpr: %s]", node.Value)
	case *ast.StringLiteralExpr:
		printer.write("[StringLiteralExpr: %s]", node.Value)
	case *ast.NamedExpr:
		printer.write("[NamedExpr: %s]", node.Name)
	case *ast.AccessExpr:
		printer.write("[AccessExpr: Field=%s", node.Field)
		printer.push("Operand=")
	case *ast.UnaryExpr:
		printer.write("[UnaryExpr: IsPrefix=%v Op=(%s)", node.IsPrefix, node.Op)
		printer.push("Operand=")
	case *ast.BinaryExpr:
		printer.write("[BinaryExpr: Op=(%s)", node.Op)
		printer.push("Left=", "Right=")
	case *ast.ImplicitStructExpr:
		printer.list(
			fmt.Sprintf("[ImplicitStructExpr: IsImproper=%v", node.IsImproper),
			nil,
			"Argument",
			len(node.Arguments))
	case *ast.ColonExpr:
		printer.list("[ColonExpr:", nil, "Argument", len(node.Arguments))
	case *ast.CallExpr:
		argLabels := []string{"FuncExpr="}
		if node.GenericArguments != nil {
			argLabels = append(argLabels, "GenericArguments=")
		}
		printer.list("[CallExpr", argLabels, "Argument", len(node.Arguments))
	case *ast.IndexExpr:
		printer.write("[IndexExpr:")
		printer.push("Accessible=", "Argument=")
	case *ast.AsExpr:
		printer.write("[AsExpr:")
		printer.push("Accessible=", "CastType=")
	case *ast.InitializeExpr:
		printer.list(
			"[InitializeExpr:",
			[]string{"Initializable="},
			"Argument",
			len(node.Arguments))
	case *ast.IfExpr:
		printer.list(
			fmt.Sprintf("[IfExpr: LabelDecl=%s", node.LabelDecl),
			nil,
			"Branch",
			len(node.ConditionBranches))
	case *ast.SwitchExpr:
		printer.list(
			fmt.Sprintf("[SwitchExpr: LabelDecl=%s", node.LabelDecl),
			[]string{"Operand="},
			"Branch",
			len(node.ConditionBranches))
	case *ast.SelectExpr:
		printer.list(
			fmt.Sprintf("[SwitchExpr: LabelDecl=%s", node.LabelDecl),
			nil,
			"Branch",
			len(node.ConditionBranches))
	case *ast.LoopExpr:
		printer.write(
			"[LoopExpr: LoopKind=%s LabelDecl=%s",
			node.LoopKind,
			node.LabelDecl)
		labels := []string{}
		if node.Init != nil {
			labels = append(labels, "Init=")
		}
		if node.Condition != nil {
			labels = append(labels, "Condition=")
		}
		if node.Post != nil {
			labels = append(labels, "Post=")
		}
		labels = append(labels, "Body=")
		printer.push(labels...)
	case *ast.StatementsExpr:
		printer.list(
			fmt.Sprintf("[StatementsExpr: LabelDecl=%s", node.LabelDecl),
			nil,
			"Statement",
			len(node.Statements))
	case *ast.VarPattern:
		printer.write("[VarPattern: Kind=%s", node.Kind)
		if node.Type != nil {
			printer.push("VarPattern=", "TypeExpr=")
		} else {
			printer.push("VarPattern=")
		}
	case *ast.CaseEnumPattern:
		printer.write("[CaseEnumPattern: EnumValue=%s", node.EnumValue)
		if node.VarPattern != nil {
			printer.push("VarPattern=")
		} else {
			printer.write("]")
		}
	case *ast.CasePatternExpr:
		printer.write("[CasePatternExpr:")
		if node.Value != nil {
			printer.push("Patterns=", "Value=")
		} else {
			printer.push("Patterns=")
		}

	case *ast.SliceTypeExpr:
		printer.write("[SliceTypeExpr:")
		printer.push("Value=")
	case *ast.ArrayTypeExpr:
		printer.write("[ArrayTypeExpr: Size=%s", node.Size)
		printer.push("Value=")
	case *ast.MapTypeExpr:
		printer.write("[MapTypeExpr:")
		printer.push("Key=", "Value=")
	case *ast.InferredTypeExpr:
		printer.write("[InferredTypeExpr: InferMutable=%v]", node.InferMutable)
	case *ast.NamedTypeExpr:
		printer.write("[NamedTypeExpr: Pkg=%s Name=%s", node.Pkg, node.Name)
		if node.GenericArguments != nil {
			printer.push("GenericArguments=")
		} else {
			printer.write("]")
		}
	case *ast.UnaryTypeExpr:
		printer.write("[UnaryTypeExpr: Op=(%s)", node.Op)
		printer.push("Operand=")
	case *ast.BinaryTypeExpr:
		printer.write("[BinaryTypeExpr: Op=(%s)", node.Op)
		printer.push("Left=", "Right=")
	case *ast.PropertiesTypeExpr:
		printer.list(
			fmt.Sprintf(
				"[PropertiesTypeExpr: Kind=%s IsImplicit=%v",
				node.Kind,
				node.IsImplicit),
			nil,
			"Property",
			len(node.Properties))
	case *ast.FuncSignature:
		printer.write("[FuncSignature: Name=%s", node.Name)
		labels := []string{}
		if node.Receiver != nil {
			labels = append(labels, "Receiver=")
		}
		if node.GenericParameters != nil {
			labels = append(labels, "GenericParameters=")
		}
		labels = append(labels, "Parameters=")
		if node.ReturnType != nil {
			labels = append(labels, "ReturnType=")
		}
		printer.push(labels...)

	case *ast.ImportStmt:
		printer.list("[ImportStmt:", nil, "ImportClause", len(node.ImportClauses))
	case *ast.JumpStmt:
		printer.write("[JumpStmt: Op=%s Label=%s", node.Op, node.Label)
		if node.Value != nil {
			printer.push("Value=")
		} else {
			printer.write("]")
		}
	case *ast.UnsafeStmt:
		printer.write(
			"[UnsafeStmt: Language=%s VerbatimSource\n",
			node.Language)
		for _, line := range strings.Split(node.VerbatimSource, "\n") {
			printer.write(printer.indent)
			printer.write("  |")
			printer.write(line)
			printer.write("\n")
		}
		printer.write(printer.indent)
		printer.write("]")
	case *ast.ConditionBranchStmt:
		printer.write(
			"[ConditionBranchStmt: IsDefaultBranch=%v",
			node.IsDefaultBranch)
		if node.Condition != nil {
			printer.push("Condition=", "Branch=")
		} else {
			printer.push("Branch=")
		}

	case *ast.FloatingComment:
		printer.write("[FloatingComment]")
	case *ast.PackageDef:
		printer.write("[PackageDef")
		printer.push("Body=")
	case *ast.FuncDefinition:
		printer.write("[FuncDefinition")
		printer.push("Signature=", "Body=")
	case *ast.TypeDef:
		printer.write("[TypeDef: Name=%s IsAlias=%v", node.Name, node.IsAlias)
		labels := []string{}
		if node.GenericParameters != nil {
			labels = append(labels, "GenericParameters=")
		}
		labels = append(labels, "BaseType=")
		if node.Constraint != nil {
			labels = append(labels, "Constraint=")
		}
		printer.push(labels...)

	case *ast.Parameter:
		printer.write(
			"[Parameter: Kind=%v Name=%s HasEllipsis=%v",
			node.Kind,
			node.Name,
			node.HasEllipsis)
		if node.Type != nil {
			printer.push("Type=")
		} else {
			printer.write("]")
		}
	case *ast.Argument:
		printer.write(
			"[Argument: Kind=%v OptionalName=%s HasEllipsis=%v",
			node.Kind,
			node.OptionalName,
			node.HasEllipsis)
		if node.Expr != nil {
			printer.push("Expr=")
		} else {
			printer.write("]")
		}
	case *ast.FieldDef:
		printer.write("FieldDef: Kind=%s Name=%s", node.Kind, node.Name)
		printer.push("Type=")
	case *ast.GenericParameter:
		printer.write("GenericParameter: Name=%s", node.Name)
		if node.Constraint != nil {
			printer.push("Constraint=")
		} else {
			printer.write("]")
		}
	case *ast.ImportClause:
		printer.write(
			"[ImportClause: Alias=%s Package=%s]",
			node.Alias,
			node.Package)

	case *ast.DefinitionList:
		printer.list("[DefinitionList:", nil, "Definition", len(node.Elements))
	case *ast.ExpressionList:
		printer.list("[ExpressionList:", nil, "Expression", len(node.Elements))
	case *ast.TypeExpressionList:
		printer.list(
			"[TypeExpressionList:",
			nil,
			"TypeExpression",
			len(node.Elements))
	case *ast.ParameterList:
		printer.list("[ParameterList:", nil, "Parameter", len(node.Elements))
	case *ast.ArgumentList:
		printer.list("[ArgumentList:", nil, "Argument", len(node.Elements))
	case *ast.GenericParameterList:
		printer.list(
			"[GenericParameterList:",
			nil,
			"GenericParameter",
			len(node.Elements))
	default:
		printer.write("unhandled node: %v", n)
	}
}

func (printer *treePrinter) Exit(n ast.Node) {
	switch node := n.(type) {
	case *ast.AccessExpr:
		printer.endNode()
	case *ast.UnaryExpr:
		printer.endNode()
	case *ast.BinaryExpr:
		printer.endNode()
	case *ast.ImplicitStructExpr:
		printer.endList(len(node.Arguments))
	case *ast.ColonExpr:
		printer.endList(len(node.Arguments))
	case *ast.CallExpr:
		printer.endNode()
	case *ast.IndexExpr:
		printer.endNode()
	case *ast.AsExpr:
		printer.endNode()
	case *ast.InitializeExpr:
		printer.endNode()
	case *ast.IfExpr:
		printer.endList(len(node.ConditionBranches))
	case *ast.SwitchExpr:
		printer.endNode()
	case *ast.SelectExpr:
		printer.endList(len(node.ConditionBranches))
	case *ast.LoopExpr:
		printer.endNode()
	case *ast.StatementsExpr:
		printer.endList(len(node.Statements))
	case *ast.VarPattern:
		printer.endNode()
	case *ast.CaseEnumPattern:
		if node.VarPattern != nil {
			printer.endNode()
		}
	case *ast.CasePatternExpr:
		printer.endNode()

	case *ast.SliceTypeExpr:
		printer.endNode()
	case *ast.ArrayTypeExpr:
		printer.endNode()
	case *ast.MapTypeExpr:
		printer.endNode()
	case *ast.NamedTypeExpr:
		if node.GenericArguments != nil {
			printer.endNode()
		}
	case *ast.UnaryTypeExpr:
		printer.endNode()
	case *ast.BinaryTypeExpr:
		printer.endNode()
	case *ast.PropertiesTypeExpr:
		printer.endList(len(node.Properties))
	case *ast.FuncSignature:
		printer.endNode()

	case *ast.ImportStmt:
		printer.endList(len(node.ImportClauses))
	case *ast.JumpStmt:
		if node.Value != nil {
			printer.endNode()
		}
	case *ast.ConditionBranchStmt:
		printer.endNode()

	case *ast.PackageDef:
		printer.endNode()
	case *ast.FuncDefinition:
		printer.endNode()
	case *ast.TypeDef:
		printer.endNode()

	case *ast.Parameter:
		if node.Type != nil {
			printer.endNode()
		}
	case *ast.Argument:
		if node.Expr != nil {
			printer.endNode()
		}
	case *ast.FieldDef:
		printer.endNode()
	case *ast.GenericParameter:
		if node.Constraint != nil {
			printer.endNode()
		}

	case *ast.DefinitionList:
		printer.endList(len(node.Elements))
	case *ast.ExpressionList:
		printer.endList(len(node.Elements))
	case *ast.TypeExpressionList:
		printer.endList(len(node.Elements))
	case *ast.ParameterList:
		printer.endList(len(node.Elements))
	case *ast.ArgumentList:
		printer.endList(len(node.Elements))
	case *ast.GenericParameterList:
		printer.endList(len(node.Elements))
	}
}
