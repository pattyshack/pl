package util

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/pattyshack/pl/ast"
)

const (
	indent = "  "
)

func TreeString(node ast.Node, indent string) string {
	buffer := &bytes.Buffer{}
	_ = PrintTree(buffer, node, indent)
	return buffer.String()
}

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
		for i := size - 1; i >= 0; i-- {
			printer.labelStack = append(
				printer.labelStack,
				fmt.Sprintf("%s%d=", elementType, i))
		}

		// push in reverse order
		printer.push(argLabels...)
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
		printer.write("[ParseErrorExpr: %s]", node.Error)
	case *ast.LiteralExpr:
		printer.write("[LiteralExpr: Kind=%s Value=%s]", node.Kind, node.Value)
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
			fmt.Sprintf("[ImplicitStructExpr: Kind=%v", node.Kind),
			nil,
			"Argument",
			len(node.Arguments))
	case *ast.CallExpr:
		argLabels := []string{"FuncExpr=", "GenericArguments="}
		printer.list("[CallExpr", argLabels, "Argument", len(node.Arguments))
	case *ast.IndexExpr:
		printer.list(
			"[IndexExpr:",
			[]string{"Accessible="},
			"IndexArg",
			len(node.IndexArgs))
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
			fmt.Sprintf("[IfExpr: Label=%s", node.Label),
			nil,
			"Branch",
			len(node.ConditionBranches))
	case *ast.SwitchExpr:
		printer.list(
			fmt.Sprintf("[SwitchExpr: Label=%s", node.Label),
			[]string{"Operand="},
			"Branch",
			len(node.ConditionBranches))
	case *ast.SelectExpr:
		printer.list(
			fmt.Sprintf("[SelectExpr: Label=%s", node.Label),
			nil,
			"Branch",
			len(node.ConditionBranches))
	case *ast.LoopExpr:
		printer.write("[LoopExpr: Kind=%s Label=%s", node.Kind, node.Label)
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
			fmt.Sprintf("[StatementsExpr: Label=%s", node.Label),
			nil,
			"Statement",
			len(node.Statements))
	case *ast.AssignPattern:
		printer.write("[AssignPattern: Kind=(%s)", node.Kind)
		printer.push("Pattern=", "Value=")
	case *ast.CasePatterns:
		printer.list("[CasePatterns:", nil, "Pattern", len(node.Patterns))
	case *ast.AddrDeclPattern:
		printer.write("[AddrDeclPattern: IsVar=%v", node.IsVar)
		printer.push("Pattern=", "TypeExpr=")
	case *ast.AssignToAddrPattern:
		printer.write("[AssignToAddrPattern:")
		printer.push("Pattern=")
	case *ast.EnumPattern:
		printer.write("[EnumPattern: EnumValue=%s", node.EnumValue)
		printer.push("Pattern=")
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
		printer.write("[InferredTypeExpr: IsImplicit=%v]", node.IsImplicit)
	case *ast.NamedTypeExpr:
		printer.write("[NamedTypeExpr: Pkg=%s Name=%s", node.Pkg, node.Name)
		printer.push("GenericArguments=")
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
		if node.GenericParameters != nil {
			labels = append(labels, "GenericParameters=")
		}
		labels = append(labels, "Parameters=", "ReturnType=")
		printer.push(labels...)

	case *ast.ImportStmt:
		printer.list("[ImportStmt:", nil, "ImportClause", len(node.ImportClauses))
	case *ast.JumpStmt:
		printer.write("[JumpStmt: Op=%s Label=%s", node.Op, node.Label)
		printer.push("Value=")
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
	case *ast.BlockAddrDeclStmt:
		printer.list(
			fmt.Sprintf("[BlockAddrDeclStmt: IsVar=%v", node.IsVar),
			nil,
			"Decl",
			len(node.Patterns))

	case *ast.FloatingComment:
		printer.write("[FloatingComment]")
	case *ast.FuncDefinition:
		printer.write("[FuncDefinition")
		printer.push("Signature=", "Body=")
	case *ast.TypeDef:
		printer.write("[TypeDef: Name=%s", node.Name)
		printer.push("GenericParameters=", "BaseType=", "Constraint=")
	case *ast.AliasDef:
		printer.write("[Alias: Alias=%s", node.Alias)
		printer.push("GenericParameters=", "Value=")
	case *ast.Parameter:
		printer.write("[Parameter: Kind=%v Name=%s", node.Kind, node.Name)
		printer.push("Type=")
	case *ast.Argument:
		printer.write("[Argument: Kind=%v Name=%s", node.Kind, node.Name)
		if node.Expr != nil {
			printer.push("Expr=")
		} else {
			printer.write("]")
		}
	case *ast.FieldDef:
		printer.write("FieldDef: IsDefault=%v Name=%s", node.IsDefault, node.Name)
		printer.push("Type=")
	case *ast.GenericParameter:
		printer.write("GenericParameter: Name=%s", node.Name)
		printer.push("Constraint=")
	case *ast.ImportClause:
		printer.write(
			"[ImportClause: Alias=%s PackageID=%s]",
			node.Alias,
			node.PackageID)

	case *ast.StatementList:
		printer.list("[StatementList:", nil, "Statement", len(node.Elements))
	case *ast.TypeExpressionList:
		printer.list(
			fmt.Sprintf("[TypeExpressionList: IsImplicit=%v", node.IsImplicit),
			nil,
			"TypeExpression",
			len(node.Elements))
	case *ast.ParameterList:
		printer.list("[ParameterList:", nil, "Parameter", len(node.Elements))
	case *ast.GenericParameterList:
		printer.list(
			fmt.Sprintf("[GenericParameterList: IsImplicit=%v", node.IsImplicit),
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
	case *ast.AddrDeclPattern:
		printer.endNode()
	case *ast.AssignToAddrPattern:
		printer.endNode()
	case *ast.AssignPattern:
		printer.endNode()
	case *ast.CasePatterns:
		printer.endList(len(node.Patterns))
	case *ast.EnumPattern:
		printer.endNode()
	case *ast.SliceTypeExpr:
		printer.endNode()
	case *ast.ArrayTypeExpr:
		printer.endNode()
	case *ast.MapTypeExpr:
		printer.endNode()
	case *ast.NamedTypeExpr:
		printer.endNode()
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
		printer.endNode()
	case *ast.ConditionBranchStmt:
		printer.endNode()
	case *ast.BlockAddrDeclStmt:
		printer.endList(len(node.Patterns))

	case *ast.FuncDefinition:
		printer.endNode()
	case *ast.TypeDef:
		printer.endNode()
	case *ast.AliasDef:
		printer.endNode()

	case *ast.Parameter:
		printer.endNode()
	case *ast.Argument:
		if node.Expr != nil {
			printer.endNode()
		}
	case *ast.FieldDef:
		printer.endNode()
	case *ast.GenericParameter:
		printer.endNode()

	case *ast.StatementList:
		printer.endList(len(node.Elements))
	case *ast.TypeExpressionList:
		printer.endList(len(node.Elements))
	case *ast.ParameterList:
		printer.endList(len(node.Elements))
	case *ast.GenericParameterList:
		printer.endList(len(node.Elements))
	}
}
