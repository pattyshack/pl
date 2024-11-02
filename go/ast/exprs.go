package ast

import (
	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/errors"
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

func (expr *LiteralExpr) Validate(emitter *errors.Emitter) {
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

func (expr *NamedExpr) IsPadding() bool {
	return expr.Name == "_"
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

func (expr *UnaryExpr) Validate(emitter *errors.Emitter) {
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

func (expr *BinaryExpr) Validate(emitter *errors.Emitter) {
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

func NewImproperUnit(pos lexutil.Location) *ImplicitStructExpr {
	return &ImplicitStructExpr{
		StartEndPos: NewStartEndPos(pos, pos),
		Kind:        ImproperImplicitStruct,
	}
}

func (expr *ImplicitStructExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	for _, arg := range expr.Arguments {
		arg.Walk(visitor)
	}
	visitor.Exit(expr)
}

func (expr *ImplicitStructExpr) Validate(emitter *errors.Emitter) {
	// NOTE: additional argument validation are handled by
	// unexpectedPatternsDetector
	switch expr.Kind {
	case ProperImplicitStruct:
		for _, arg := range expr.Arguments {
			if arg.Kind == VariadicArgument {
				emitter.Emit(arg.Loc(), "unexpected %s argument", arg.Kind)
			}
		}
	case ImproperImplicitStruct, ColonImplicitStruct:
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
// ParameterizedExpr
//

// NOTE: ParameterizedExpr fields should be identical to NamedTypeExpr
type ParameterizedExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	// NOTE: Pkg is an identifier or empty string for normal ast construction.
	// But for the purpose of type cataloging / checking, Pkg is the fully
	// qualified package path (even for local names)
	Pkg  string // optional.  "" = local
	Name string

	Parameters []TypeExpression
}

var _ Expression = &ParameterizedExpr{}

func (expr *ParameterizedExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	for _, arg := range expr.Parameters {
		arg.Walk(visitor)
	}
	visitor.Exit(expr)
}

//
// CallExpr
//

type CallExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	FuncExpr  Expression
	Arguments []*Argument
}

var _ Expression = &CallExpr{}
var _ Validator = &CallExpr{}

func (expr *CallExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.FuncExpr.Walk(visitor)
	for _, arg := range expr.Arguments {
		arg.Walk(visitor)
	}
	visitor.Exit(expr)
}

func (expr *CallExpr) Validate(emitter *errors.Emitter) {
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
// MakeExpr
//

type MakeExpr struct {
	IsExpr
	StartEndPos
	LeadingTrailingComments

	VariableSizedType TypeExpression
	Size              Expression
	Capacity          Expression // optional. only applicable to slice
	Value             Expression // optional. only applicable to slice
}

var _ Expression = &MakeExpr{}
var _ Validator = &MakeExpr{}

func (expr *MakeExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	expr.VariableSizedType.Walk(visitor)
	expr.Size.Walk(visitor)
	if expr.Capacity != nil {
		expr.Capacity.Walk(visitor)
	}
	if expr.Value != nil {
		expr.Value.Walk(visitor)
	}
	visitor.Exit(expr)
}

func (expr *MakeExpr) Validate(emitter *errors.Emitter) {
	switch expr.VariableSizedType.(type) {
	case *SliceTypeExpr: // ok
	case *MapTypeExpr:
		if expr.Capacity != nil {
			emitter.Emit(
				expr.Capacity.Loc(),
				"unexpected capacity specified for map type")
		}

		if expr.Value != nil {
			emitter.Emit(
				expr.Value.Loc(),
				"unexpected initial value specified for map type")
		}
	default:
		emitter.Emit(
			expr.VariableSizedType.Loc(),
			"unexpected fixed size type, make only operate on slice or map types")
	}
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

func (expr *InitializeExpr) Validate(emitter *errors.Emitter) {
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

func (expr *StatementsExpr) Validate(emitter *errors.Emitter) {
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
	emitter *errors.Emitter,
	parent Expression,
	branches []*ConditionBranchStmt,
	checkCondition func(Expression),
) {
	if len(branches) == 0 {
		emitter.Emit(parent.Loc(), "expected at least one branch")
	}

	for idx, branch := range branches {
		if branch.IsDefaultBranch {
			if idx != len(branches)-1 {
				emitter.Emit(branch.Loc(), "default branch is not the last branch")
			}

			if branch.Condition != nil {
				emitter.Emit(
					branch.Condition.Loc(),
					"invalid ast construction. condition set on default branch")
			}
		} else if branch.Condition == nil {
			emitter.Emit(
				branch.Loc(),
				"invalid ast construction. condition not set on non-default branch")
		} else {
			checkCondition(branch.Condition)
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

func (expr *IfExpr) Validate(emitter *errors.Emitter) {
	validateDefaultBranch(
		emitter,
		expr,
		expr.ConditionBranches,
		func(cond Expression) {
			assign, ok := cond.(*AssignPattern)
			if !ok {
				return
			}

			if assign.Kind != EqualAssign {
				emitter.Emit(
					cond.Loc(),
					"invalid ast construction.  unexpected assign pattern kind (%s)",
					assign.Kind)
			}
		})
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

func (expr *SwitchExpr) Validate(emitter *errors.Emitter) {
	validateDefaultBranch(
		emitter,
		expr,
		expr.ConditionBranches,
		func(cond Expression) {
			_, ok := cond.(*CasePatterns)
			if !ok {
				emitter.Emit(
					cond.Loc(),
					"invalid ast construction. expecting case patterns")
			}
		})
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

func (expr *SelectExpr) Validate(emitter *errors.Emitter) {
	validateDefaultBranch(
		emitter,
		expr,
		expr.ConditionBranches,
		func(condition Expression) {
			assign, ok := condition.(*AssignPattern)
			if ok {
				if assign.Kind != EqualAssign {
					emitter.Emit(
						condition.Loc(),
						"invalid ast construction. unexpected assign pattern kind (%s)",
						assign.Kind)
				}
				condition = assign.Value
			}

			switch nary := condition.(type) {
			case *CasePatterns:
				emitter.Emit(
					condition.Loc(),
					"unexpected pattern list, expecting only one pattern per case")
			case *UnaryExpr:
				if nary.Op != UnaryRecvOp {
					emitter.Emit(
						condition.Loc(),
						"unexpected expression, expecting send/recv expression")
				}
			case *BinaryExpr:
				if nary.Op != BinarySendOp {
					emitter.Emit(
						condition.Loc(),
						"unexpected expression, expecting send/recv expression")
				}
			}
		})
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

	Condition Expression // nil for infinite loop. Required otherwise
	Post      Expression // required for semi-traditional for loop. nil otherwise

	Body *StatementsExpr
}

var _ Expression = &LoopExpr{}
var _ Validator = &LoopExpr{}

func (expr *LoopExpr) Walk(visitor Visitor) {
	visitor.Enter(expr)
	if expr.Condition != nil {
		expr.Condition.Walk(visitor)
	}
	if expr.Post != nil {
		expr.Post.Walk(visitor)
	}
	expr.Body.Walk(visitor)
	visitor.Exit(expr)
}

func (expr *LoopExpr) Validate(emitter *errors.Emitter) {
	switch expr.Kind {
	case InfiniteLoop:
		if expr.Condition != nil {
			emitter.Emit(
				expr.Condition.Loc(),
				"invalid ast construction. condition expression set in %s loop",
				expr.Kind)
		}
	case DoWhileLoop, WhileLoop, IteratorLoop, ForLoop:
		if expr.Condition == nil {
			emitter.Emit(
				expr.Condition.Loc(),
				"invalid ast construction. nil condition expression in %s loop",
				expr.Kind)
		}
	default:
		emitter.Emit(
			expr.Loc(),
			"invalid ast construction.  unexpected loop kind (%s)",
			expr.Kind)
	}

	switch expr.Kind {
	case InfiniteLoop, DoWhileLoop, WhileLoop, IteratorLoop:
		if expr.Post != nil {
			emitter.Emit(
				expr.Post.Loc(),
				"invalid ast construction. post expression set in %s loop",
				expr.Kind)
		}
	case ForLoop:
		if expr.Post == nil {
			emitter.Emit(
				expr.Post.Loc(),
				"invalid ast construction. nil post expression in %s loop",
				expr.Kind)
		}
	}

	if expr.Kind == InfiniteLoop {
		return
	}

	assign, ok := expr.Condition.(*AssignPattern)
	if ok {
		if expr.Kind != IteratorLoop || assign.Kind != InAssign {
			emitter.Emit(
				expr.Condition.Loc(),
				"invalid ast construction. unexpected condition set in %s loop",
				expr.Kind)
		}
	} else {
		if expr.Kind == IteratorLoop {
			emitter.Emit(
				expr.Condition.Loc(),
				"invalid ast construction. unexpected condition set in %s loop",
				expr.Kind)
		}
	}
}

//
// FuncDefinition
//

type FuncDefinition struct {
	IsExpr
	IsTypeProp
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
