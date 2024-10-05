package ast

import (
	"github.com/pattyshack/gt/lexutil"
)

//
// ParseErrorExpr
//

type ParseErrorExpr struct {
	IsExpr

	StartEndPos
	LeadingTrailingComments

	Error error
}

func (s *ParseErrorExpr) Walk(visitor Visitor) {
	visitor.Enter(s)
	visitor.Exit(s)
}

//
// (Bool/Int/Float/Rune/String)LiteralExpr
//

type LiteralKind string

const (
	BoolLiteral   = LiteralKind("bool")
	IntLiteral    = LiteralKind("int")
	FloatLiteral  = LiteralKind("float")
	RuneLiteral   = LiteralKind("rune")
	StringLiteral = LiteralKind("string")
)

type LiteralExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Kind LiteralKind

	Value string
}

var _ Expression = &LiteralExpr{}
var _ Validator = &LiteralExpr{}

func (expr *LiteralExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	visitor.Exit(expr)
}

func (expr *LiteralExpr) Validate(emitter *lexutil.ErrorEmitter) {
	switch expr.Kind {
	case BoolLiteral, IntLiteral, FloatLiteral, RuneLiteral, StringLiteral: // ok
	default:
		emitter.Emit(
			expr.Loc(),
			"invalid ast construction. unexpected literal kind (%s)",
			expr.Kind)
	}
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

var _ Expression = &NamedExpr{}

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

var _ Expression = &AccessExpr{}

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

var _ Expression = &UnaryExpr{}
var _ Validator = &UnaryExpr{}

func (expr *UnaryExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Operand.Walk(visitor)
	visitor.Exit(expr)
}

func (expr *UnaryExpr) Validate(emitter *lexutil.ErrorEmitter) {
	switch expr.Op {
	case UnaryReturnDefaultEnumOp,
		UnaryPanicOnDefaultEnumOp,
		UnaryAssignAddOneOp,
		UnaryAssignSubOneOp,
		UnaryNotOp,
		UnaryBitwiseComplementOp,
		UnaryPlusOp,
		UnaryMinusOp,
		UnaryDerefOp,
		UnaryRefOp,
		UnaryAsyncOp,
		UnaryDeferOp,
		UnaryAssignToOp,
		UnaryRecvOp:
		// ok
	default:
		emitter.Emit(
			expr.Loc(),
			"invalid ast construction. unexpected unary op (%s)",
			expr.Op)
	}
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

var _ Expression = &BinaryExpr{}
var _ Validator = &BinaryExpr{}

func (expr *BinaryExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Left.Walk(visitor)
	expr.Right.Walk(visitor)
	visitor.Exit(expr)
}

func (expr *BinaryExpr) Validate(emitter *lexutil.ErrorEmitter) {
	switch expr.Op {
	case BinaryAddOp,
		BinarySubOp,
		BinaryMulOp,
		BinaryDivOp,
		BinaryModOp,
		BinaryBitAndOp,
		BinaryBitOrOp,
		BinaryBitXorOp,
		BinaryBitLshiftOp,
		BinaryBitRshiftOp,
		BinaryEqualOp,
		BinaryNotEqualOp,
		BinaryLessOp,
		BinaryLessOrEqualOp,
		BinaryGreaterOp,
		BinaryGreaterOrEqualOp,
		BinaryAndOp,
		BinaryOrOp,
		BinarySendOp,
		BinaryAddAssignOp,
		BinarySubAssignOp,
		BinaryMulAssignOp,
		BinaryDivAssignOp,
		BinaryModAssignOp,
		BinaryBitAndAssignOp,
		BinaryBitOrAssignOp,
		BinaryBitXorAssignOp,
		BinaryBitLshiftAssignOp,
		BinaryBitRshiftAssignOp:
		//ok
	default:
		emitter.Emit(
			expr.Loc(),
			"invalid ast construction. unexpected binary op (%s)",
			expr.Op)
	}
}

//
// ImplicitStructExpr
//

type ImplicitStructKind string

const (
	// Comma separated list of expressions left/right paren.
	ProperImplicitStruct = ImplicitStructKind("proper")

	// Comma separated list of expressions without left/right paren.
	ImproperImplicitStruct = ImplicitStructKind("improper")

	// Colon seperated list of (optional) expressions
	ColonImplicitStruct = ImplicitStructKind("colon")
)

type ImplicitStructExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Kind ImplicitStructKind

	Arguments []*Argument
}

var _ Expression = &ImplicitStructExpr{}
var _ Validator = &ImplicitStructExpr{}

func (expr *ImplicitStructExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	for _, arg := range expr.Arguments {
		arg.Walk(visitor)
	}
	visitor.Exit(expr)
}

func (expr *ImplicitStructExpr) Validate(emitter *lexutil.ErrorEmitter) {
	switch expr.Kind {
	case ProperImplicitStruct:
		// NOTE: proper implicit struct argument validation are handled by
		// unexpectedPatternsDetector
	case ImproperImplicitStruct, ColonImplicitStruct:
		// NOTE: additional argument validation are handled by
		// unexpectedPatternsDetector
		for _, arg := range expr.Arguments {
			if arg.Kind == SkipPatternArgument || arg.Kind == VariadicArgument {
				emitter.Emit(arg.Loc(), "unexpected %s argument", arg.Kind)
			}
		}
	default:
		emitter.Emit(
			expr.Loc(),
			"invalid ast construction. unexpected implicit struct expr kind (%s)",
			expr.Kind)
	}
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
	Arguments        []*Argument
}

var _ Expression = &CallExpr{}
var _ Validator = &CallExpr{}

func (expr *CallExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.FuncExpr.Walk(visitor)
	if expr.GenericArguments != nil {
		expr.GenericArguments.Walk(visitor)
	}
	for _, arg := range expr.Arguments {
		arg.Walk(visitor)
	}
	visitor.Exit(expr)
}

func (expr *CallExpr) Validate(emitter *lexutil.ErrorEmitter) {
	for idx, arg := range expr.Arguments {
		if arg.Kind == SkipPatternArgument {
			emitter.Emit(arg.Loc(), "unexpected %s argument", arg.Kind)
		}

		if arg.Kind == VariadicArgument && idx != len(expr.Arguments)-1 {
			emitter.Emit(
				arg.Loc(),
				"%s argument must be the last argument in the list",
				arg.Kind)
		}
	}
}

//
// IndexExpr
//

type IndexExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Accessible Expression

	// NOTE: colon expression is decomposed back into expression list.
	IndexArgs []Expression
}

var _ Expression = &IndexExpr{}

func (expr *IndexExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Accessible.Walk(visitor)
	for _, arg := range expr.IndexArgs {
		arg.Walk(visitor)
	}
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

var _ Expression = &AsExpr{}

func (expr *AsExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Accessible.Walk(visitor)
	expr.CastType.Walk(visitor)
	visitor.Exit(expr)
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

var _ Expression = &InitializeExpr{}
var _ Validator = &InitializeExpr{}

func (expr *InitializeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Initializable.Walk(visitor)
	for _, arg := range expr.Arguments {
		arg.Walk(visitor)
	}
	visitor.Exit(expr)
}

func (expr *InitializeExpr) Validate(emitter *lexutil.ErrorEmitter) {
	for _, arg := range expr.Arguments {
		if arg.Kind == SkipPatternArgument || arg.Kind == VariadicArgument {
			emitter.Emit(arg.Loc(), "unexpected %s argument", arg.Kind)
		}
	}
}

//
// StatementsExpr
//

type StatementsExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Label
	Statements []Statement
}

var _ Expression = &StatementsExpr{}
var _ Validator = &StatementsExpr{}

func (expr *StatementsExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	for _, stmt := range expr.Statements {
		stmt.Walk(visitor)
	}
	visitor.Exit(expr)
}

func (expr *StatementsExpr) Validate(emitter *lexutil.ErrorEmitter) {
	for idx, stmt := range expr.Statements {
		_, ok := stmt.(*JumpStmt)
		if !ok {
			continue
		}

		if idx < len(expr.Statements)-1 {
			emitter.Emit(expr.Statements[idx+1].Loc(), "unreachable statement")
			return
		}
	}
}

//
// IfExpr
//

func validateDefaultBranch(
	emitter *lexutil.ErrorEmitter,
	parent Expression,
	branches []*ConditionBranchStmt,
) {
	if len(branches) == 0 {
		emitter.Emit(parent.Loc(), "expected at least one branch")
	}

	for idx, branch := range branches {
		if branch.IsDefaultBranch {
			if idx != len(branches)-1 {
				emitter.Emit(branch.Loc(), "default branch is not the last branch")
			}
		}
	}
}

type IfExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Label
	ConditionBranches []*ConditionBranchStmt
}

var _ Expression = &IfExpr{}
var _ Validator = &IfExpr{}

func (expr *IfExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	for _, condBranch := range expr.ConditionBranches {
		condBranch.Walk(visitor)
	}
	visitor.Exit(expr)
}

func (expr *IfExpr) Validate(emitter *lexutil.ErrorEmitter) {
	validateDefaultBranch(emitter, expr, expr.ConditionBranches)
}

//
// SwitchExpr
//

type SwitchExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Label
	Operand           Expression
	ConditionBranches []*ConditionBranchStmt
}

var _ Expression = &SwitchExpr{}
var _ Validator = &SwitchExpr{}

func (expr *SwitchExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.Operand.Walk(visitor)
	for _, condBranch := range expr.ConditionBranches {
		condBranch.Walk(visitor)
	}
	visitor.Exit(expr)
}

func (expr *SwitchExpr) Validate(emitter *lexutil.ErrorEmitter) {
	validateDefaultBranch(emitter, expr, expr.ConditionBranches)
}

//
// SelectExpr
//

type SelectExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Label
	ConditionBranches []*ConditionBranchStmt
}

var _ Expression = &SelectExpr{}
var _ Validator = &SelectExpr{}

func (expr *SelectExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	for _, condBranch := range expr.ConditionBranches {
		condBranch.Walk(visitor)
	}
	visitor.Exit(expr)
}

func (expr *SelectExpr) Validate(emitter *lexutil.ErrorEmitter) {
	validateDefaultBranch(emitter, expr, expr.ConditionBranches)
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

	Kind LoopKind

	Label

	Init      Statement  // optional. only applicable to traditional-for loop
	Condition Expression // optional. not applicable to infinite loop
	Post      Statement  // optional. only applicable to traditional-for loop

	Body *StatementsExpr
}

var _ Expression = &LoopExpr{}
var _ Validator = &LoopExpr{}

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

func (expr *LoopExpr) Validate(emitter *lexutil.ErrorEmitter) {
	switch expr.Kind {
	case InfiniteLoop,
		DoWhileLoop,
		WhileLoop,
		IteratorLoop,
		ForLoop:
		// ok
	default:
		emitter.Emit(
			expr.Loc(),
			"invalid ast construction.  unexpected loop kind (%s)",
			expr.Kind)
	}

	if expr.Kind != ForLoop {
		if expr.Init != nil {
			emitter.Emit(
				expr.Init.Loc(),
				"invalid ast construction. init statement set in %s loop",
				expr.Kind)
		}

		if expr.Post != nil {
			emitter.Emit(
				expr.Post.Loc(),
				"invalid ast construction. post statement set in %s loop",
				expr.Kind)
		}
	}

	if expr.Kind == InfiniteLoop {
		if expr.Condition != nil {
			emitter.Emit(
				expr.Condition.Loc(),
				"invalid ast construction. condition expression set in %s loop",
				expr.Kind)
		}
		return
	}

	if expr.Condition == nil {
		emitter.Emit(
			expr.Loc(),
			"invalid ast construction. nil condition expression in %s loop",
			expr.Kind)
		return
	}

	if expr.Kind != IteratorLoop {
		return
	}

	assign, ok := expr.Condition.(*AssignPattern)
	if !ok || assign.Kind != InAssign {
		emitter.Emit(
			expr.Condition.Loc(),
			"invalid ast construction. condition expresion set in %s loop",
			expr.Kind)
	}
}

//
// FuncDefinition
//

type FuncDefinition struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	Signature *FuncSignature
	Body      Expression
}

var _ Expression = &FuncDefinition{}

func (def *FuncDefinition) Walk(visitor Visitor) {
	visitor.Enter(def)
	def.Signature.Walk(visitor)
	def.Body.Walk(visitor)
	visitor.Exit(def)
}
