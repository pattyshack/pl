package ast

import (
	"fmt"
	"io"
	"strings"
)

const (
	indent = "  "
)

func PrintTree(output io.Writer, node Node, indent string) error {
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

func (printer *treePrinter) list(elementType string, size int) {
	printer.write("[%sList:", elementType)
	if size == 0 {
		printer.write("]")
	} else {
		printer.indent += indent

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

func (printer *treePrinter) Enter(n Node) {
	printer.writeLabel()

	switch node := n.(type) {
	case *ParseErrorNode:
		printer.write("[ParseErrorNode: %s %s]", node.Errors[0], node.StartPos)
	case *BoolLiteralExpr:
		printer.write("[BoolLiteralExpr: %s]", node.Value)
	case *IntLiteralExpr:
		printer.write("[IntLiteralExpr: %s]", node.Value)
	case *FloatLiteralExpr:
		printer.write("[FloatLiteralExpr: %s]", node.Value)
	case *RuneLiteralExpr:
		printer.write("[RuneLiteralExpr: %s]", node.Value)
	case *StringLiteralExpr:
		printer.write("[StringLiteralExpr: %s]", node.Value)
	case *NamedExpr:
		printer.write("[NamedExpr: %s]", node.Name)
	case *AccessExpr:
		printer.write("[AccessExpr: Field=%s", node.Field)
		printer.push("Operand=")
	case *UnaryExpr:
		printer.write("[UnaryExpr: IsPrefix=%v Op=(%s)", node.IsPrefix, node.Op)
		printer.push("Operand=")
	case *BinaryExpr:
		printer.write("[BinaryExpr: Op=(%s)", node.Op)
		printer.push("Left=", "Right=")
	case *ImplicitStructExpr:
		printer.write("[ImplicitStructExpr: IsImproper=%v", node.IsImproper)
		printer.push("Arguments=")
	case *ColonExpr:
		printer.write("[ColonExpr:")
		printer.push("Arguments=")
	case *CallExpr:
		printer.write("[CallExpr")
		if node.GenericArguments != nil {
			printer.push("FuncExpr=", "GenericArguments=", "Arguments=")
		} else {
			printer.push("FuncExpr=", "Arguments=")
		}
	case *IndexExpr:
		printer.write("[IndexExpr:")
		printer.push("Accessible=", "Argument=")
	case *AsExpr:
		printer.write("[AsExpr:")
		printer.push("Accessible=", "CastType=")
	case *InitializeExpr:
		printer.write("[InitializeExpr:")
		printer.push("Initializable=", "Arguments=")
	case *IfExpr:
		printer.write("[IfExpr: LabelDecl=%s", node.LabelDecl)
		printer.push("Branches=")
	case *SwitchExpr:
		printer.write("[SwitchExpr: LabelDecl=%s", node.LabelDecl)
		printer.push("Operand=", "Branches=")
	case *SelectExpr:
		printer.write("[SwitchExpr: LabelDecl=%s", node.LabelDecl)
		printer.push("Branches=")
	case *LoopExpr:
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
	case *StatementsExpr:
		printer.write("[StatementsExpr: LabelDecl=%s", node.LabelDecl)
		printer.push("Statements=")
	case *VarPattern:
		printer.write("[VarPattern: Kind=%s", node.Kind)
		if node.Type != nil {
			printer.push("VarPattern=", "TypeExpr=")
		} else {
			printer.push("VarPattern=")
		}
	case *CaseAssignPattern:
		printer.write("[CaseAssignPattern:")
		printer.push("AssignPattern=", "Value=")
	case *CaseEnumPattern:
		printer.write("[CaseEnumPattern: EnumValue=%s", node.EnumValue)
		if node.VarPattern != nil {
			printer.push("VarPattern=")
		} else {
			printer.write("]")
		}

	case *SliceTypeExpr:
		printer.write("[SliceTypeExpr:")
		printer.push("Value=")
	case *ArrayTypeExpr:
		printer.write("[ArrayTypeExpr: Size=%s", node.Size)
		printer.push("Value=")
	case *MapTypeExpr:
		printer.write("[MapTypeExpr:")
		printer.push("Key=", "Value=")
	case *InferredTypeExpr:
		printer.write("[InferredTypeExpr: InferMutable=%v]", node.InferMutable)
	case *NamedTypeExpr:
		printer.write("[NamedTypeExpr: Pkg=%s Name=%s", node.Pkg, node.Name)
		if node.GenericArguments != nil {
			printer.push("GenericArguments=")
		} else {
			printer.write("]")
		}
	case *UnaryTypeExpr:
		printer.write("[UnaryTypeExpr: Op=(%s)", node.Op)
		printer.push("Operand=")
	case *BinaryTypeExpr:
		printer.write("[BinaryTypeExpr: Op=(%s)", node.Op)
		printer.push("Left=", "Right=")
	case *PropertiesTypeExpr:
		printer.write(
			"[PropertiesTypeExpr: Kind=%s IsImplicit=%v",
			node.Kind,
			node.IsImplicit)
		printer.push("Properties=")
	case *FuncSignature:
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

	case *ImportStatement:
		printer.write("[ImportStatement:")
		printer.push("ImportClauses=")
	case *JumpStatement:
		printer.write("[JumpStatement: Op=%s Label=%s", node.Op, node.Label)
		if node.Value != nil {
			printer.push("Value=")
		} else {
			printer.write("]")
		}
	case *UnsafeStatement:
		printer.write(
			"[UnsafeStatement: Language=%s VerbatimSource\n",
			node.Language)
		for _, line := range strings.Split(node.VerbatimSource, "\n") {
			printer.write(printer.indent)
			printer.write("  |")
			printer.write(line)
			printer.write("\n")
		}
		printer.write(printer.indent)
		printer.write("]")
	case *BranchStatement:
		printer.write("[BranchStatement: IsDefault=%v", node.IsDefault)
		if node.CasePatterns != nil {
			printer.push("CasePatterns=", "Body=")
		} else {
			printer.push("Body=")
		}

	case *FloatingComment:
		printer.write("[FloatingComment]")
	case *PackageDef:
		printer.write("[PackageDef")
		printer.push("Body=")
	case *FuncDefinition:
		printer.write("[FuncDefinition")
		printer.push("Signature=", "Body=")
	case *TypeDef:
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

	case *Parameter:
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
	case *Argument:
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
	case *FieldDef:
		printer.write("FieldDef: Kind=%s Name=%s", node.Kind, node.Name)
		printer.push("Type=")
	case *GenericParameter:
		printer.write("GenericParameter: Name=%s", node.Name)
		if node.Constraint != nil {
			printer.push("Constraint=")
		} else {
			printer.write("]")
		}
	case *ImportClause:
		printer.write(
			"[ImportClause: Alias=%s Package=%s]",
			node.Alias,
			node.Package)
	case *ConditionBranch:
		printer.write("[ConditionBranch: IsElse=%v", node.IsElse)
		if node.Condition != nil {
			printer.push("Condition=", "Branch=")
		} else {
			printer.push("Branch=")
		}

	case *DefinitionList:
		printer.list("Definition", len(node.Elements))
	case *ExpressionList:
		printer.list("Expression", len(node.Elements))
	case *TypePropertyList:
		printer.list("TypeProperty", len(node.Elements))
	case *TypeExpressionList:
		printer.list("TypeExpression", len(node.Elements))
	case *StatementList:
		printer.list("Statement", len(node.Elements))
	case *ParameterList:
		printer.list("Parameter", len(node.Elements))
	case *ArgumentList:
		printer.list("Argument", len(node.Elements))
	case *GenericParameterList:
		printer.list("GenericParameter", len(node.Elements))
	case *ImportClauseList:
		printer.list("ImportClause", len(node.Elements))
	case *ConditionBranchList:
		printer.list("ConditionBranch", len(node.Elements))
	default:
		printer.write("unhandled node: %v", n)
	}
}

func (printer *treePrinter) Exit(n Node) {
	switch node := n.(type) {
	case *AccessExpr:
		printer.endNode()
	case *UnaryExpr:
		printer.endNode()
	case *BinaryExpr:
		printer.endNode()
	case *ImplicitStructExpr:
		printer.endNode()
	case *ColonExpr:
		printer.endNode()
	case *CallExpr:
		printer.endNode()
	case *IndexExpr:
		printer.endNode()
	case *AsExpr:
		printer.endNode()
	case *InitializeExpr:
		printer.endNode()
	case *IfExpr:
		printer.endNode()
	case *SwitchExpr:
		printer.endNode()
	case *SelectExpr:
		printer.endNode()
	case *LoopExpr:
		printer.endNode()
	case *StatementsExpr:
		printer.endNode()
	case *VarPattern:
		printer.endNode()
	case *CaseAssignPattern:
		printer.endNode()
	case *CaseEnumPattern:
		if node.VarPattern != nil {
			printer.endNode()
		}

	case *SliceTypeExpr:
		printer.endNode()
	case *ArrayTypeExpr:
		printer.endNode()
	case *MapTypeExpr:
		printer.endNode()
	case *NamedTypeExpr:
		if node.GenericArguments != nil {
			printer.endNode()
		}
	case *UnaryTypeExpr:
		printer.endNode()
	case *BinaryTypeExpr:
		printer.endNode()
	case *PropertiesTypeExpr:
		printer.endNode()
	case *FuncSignature:
		printer.endNode()

	case *ImportStatement:
		printer.endNode()
	case *JumpStatement:
		if node.Value != nil {
			printer.endNode()
		}
	case *BranchStatement:
		printer.endNode()

	case *PackageDef:
		printer.endNode()
	case *FuncDefinition:
		printer.endNode()
	case *TypeDef:
		printer.endNode()

	case *Parameter:
		if node.Type != nil {
			printer.endNode()
		}
	case *Argument:
		if node.Expr != nil {
			printer.endNode()
		}
	case *FieldDef:
		printer.endNode()
	case *GenericParameter:
		if node.Constraint != nil {
			printer.endNode()
		}
	case *ConditionBranch:
		printer.endNode()

	case *DefinitionList:
		printer.endList(len(node.Elements))
	case *ExpressionList:
		printer.endList(len(node.Elements))
	case *TypePropertyList:
		printer.endList(len(node.Elements))
	case *TypeExpressionList:
		printer.endList(len(node.Elements))
	case *StatementList:
		printer.endList(len(node.Elements))
	case *ParameterList:
		printer.endList(len(node.Elements))
	case *ArgumentList:
		printer.endList(len(node.Elements))
	case *GenericParameterList:
		printer.endList(len(node.Elements))
	case *ImportClauseList:
		printer.endList(len(node.Elements))
	case *ConditionBranchList:
		printer.endList(len(node.Elements))
	}
}
