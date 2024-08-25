// Auto-generated from source: grammar.lr

package parser

import (
	fmt "fmt"
	io "io"
	sort "sort"
)

type SymbolId int

const (
	SpacesToken          = SymbolId(256)
	NewlinesToken        = SymbolId(257)
	CommentToken         = SymbolId(258)
	IntegerLiteralToken  = SymbolId(259)
	FloatLiteralToken    = SymbolId(260)
	RuneLiteralToken     = SymbolId(261)
	StringLiteralToken   = SymbolId(262)
	IdentifierToken      = SymbolId(263)
	TrueToken            = SymbolId(264)
	FalseToken           = SymbolId(265)
	IfToken              = SymbolId(266)
	ElseToken            = SymbolId(267)
	SwitchToken          = SymbolId(268)
	CaseToken            = SymbolId(269)
	DefaultToken         = SymbolId(270)
	ForToken             = SymbolId(271)
	ReturnToken          = SymbolId(272)
	BreakToken           = SymbolId(273)
	ContinueToken        = SymbolId(274)
	PackageToken         = SymbolId(275)
	UnsafeToken          = SymbolId(276)
	TypeToken            = SymbolId(277)
	StructToken          = SymbolId(278)
	EnumToken            = SymbolId(279)
	TraitToken           = SymbolId(280)
	FuncToken            = SymbolId(281)
	AsyncToken           = SymbolId(282)
	LbraceToken          = SymbolId(283)
	RbraceToken          = SymbolId(284)
	LparenToken          = SymbolId(285)
	RparenToken          = SymbolId(286)
	LbracketToken        = SymbolId(287)
	RbracketToken        = SymbolId(288)
	DotToken             = SymbolId(289)
	CommaToken           = SymbolId(290)
	QuestionToken        = SymbolId(291)
	SemicolonToken       = SymbolId(292)
	DollarLbracketToken  = SymbolId(293)
	DotdotdotToken       = SymbolId(294)
	ExclaimToken         = SymbolId(295)
	ExclaimExclaimToken  = SymbolId(296)
	LabelDeclToken       = SymbolId(297)
	JumpLabelToken       = SymbolId(298)
	AddAssignToken       = SymbolId(299)
	SubAssignToken       = SymbolId(300)
	MulAssignToken       = SymbolId(301)
	DivAssignToken       = SymbolId(302)
	ModAssignToken       = SymbolId(303)
	AddOneAssignToken    = SymbolId(304)
	SubOneAssignToken    = SymbolId(305)
	BitNegAssignToken    = SymbolId(306)
	BitAndAssignToken    = SymbolId(307)
	BitOrAssignToken     = SymbolId(308)
	BitXorAssignToken    = SymbolId(309)
	BitLshiftAssignToken = SymbolId(310)
	BitRshiftAssignToken = SymbolId(311)
	NotToken             = SymbolId(312)
	AndToken             = SymbolId(313)
	OrToken              = SymbolId(314)
	AddToken             = SymbolId(315)
	SubToken             = SymbolId(316)
	MulToken             = SymbolId(317)
	DivToken             = SymbolId(318)
	ModToken             = SymbolId(319)
	BitNegToken          = SymbolId(320)
	BitAndToken          = SymbolId(321)
	BitXorToken          = SymbolId(322)
	BitOrToken           = SymbolId(323)
	BitLshiftToken       = SymbolId(324)
	BitRshiftToken       = SymbolId(325)
	EqualToken           = SymbolId(326)
	NotEqualToken        = SymbolId(327)
	LessToken            = SymbolId(328)
	LessOrEqualToken     = SymbolId(329)
	GreaterToken         = SymbolId(330)
	GreaterOrEqualToken  = SymbolId(331)
	LexErrorToken        = SymbolId(332)
)

type Location struct {
	FileName string
	Line     int
	Column   int
}

func (l Location) String() string {
	return fmt.Sprintf("%v:%v:%v", l.FileName, l.Line, l.Column)
}

func (l Location) ShortString() string {
	return fmt.Sprintf("%v:%v", l.Line, l.Column)
}

type Token interface {
	Id() SymbolId
	Loc() Location
}

type GenericSymbol struct {
	SymbolId
	Location
}

func (t *GenericSymbol) Id() SymbolId { return t.SymbolId }

func (t *GenericSymbol) Loc() Location { return t.Location }

type Lexer interface {
	// Note: Return io.EOF to indicate end of stream
	// Token with unspecified value type should return *GenericSymbol
	Next() (Token, error)

	CurrentLocation() Location
}

type Reducer interface {
	// 99:2: literal -> TRUE: ...
	TrueToLiteral(True_ *GenericSymbol) (*GenericSymbol, error)

	// 100:2: literal -> FALSE: ...
	FalseToLiteral(False_ *GenericSymbol) (*GenericSymbol, error)

	// 101:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 102:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 103:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 104:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 106:23: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(ExplicitFuncType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 111:2: anonymous_struct_expr -> explicit: ...
	ExplicitToAnonymousStructExpr(ExplicitStructType_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 112:2: anonymous_struct_expr -> implicit: ...
	ImplicitToAnonymousStructExpr(Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 118:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 119:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 120:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 122:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 123:2: atom_expr -> anonymous_struct_expr: ...
	AnonymousStructExprToAtomExpr(AnonymousStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 124:2: atom_expr -> LEX_ERROR: ...
	LexErrorToAtomExpr(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 128:2: generic_arguments -> value_type: ...
	ValueTypeToGenericArguments(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 129:2: generic_arguments -> add: ...
	AddToGenericArguments(GenericArguments_ *GenericSymbol, Comma_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 132:2: optional_generic_arguments -> generic_arguments: ...
	GenericArgumentsToOptionalGenericArguments(GenericArguments_ *GenericSymbol) (*GenericSymbol, error)

	// 133:2: optional_generic_arguments -> nil: ...
	NilToOptionalGenericArguments() (*GenericSymbol, error)

	// 136:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ *GenericSymbol, OptionalGenericArguments_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 137:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (*GenericSymbol, error)

	// 141:2: argument -> positional: ...
	PositionalToArgument(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 144:2: arguments -> argument: ...
	ArgumentToArguments(Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 145:2: arguments -> add: ...
	AddToArguments(Arguments_ *GenericSymbol, Comma_ *GenericSymbol, Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 148:2: optional_arguments -> arguments: ...
	ArgumentsToOptionalArguments(Arguments_ *GenericSymbol) (*GenericSymbol, error)

	// 149:2: optional_arguments -> nil: ...
	NilToOptionalArguments() (*GenericSymbol, error)

	// 151:13: call_expr -> ...
	ToCallExpr(AccessExpr_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalArguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 154:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 155:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 156:2: access_expr -> call_expr: ...
	CallExprToAccessExpr(CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 158:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Expression_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 161:2: postfix_unary_expr -> access_expr: ...
	AccessExprToPostfixUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 162:2: postfix_unary_expr -> question: ...
	QuestionToPostfixUnaryExpr(AccessExpr_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 165:2: prefix_unary_op -> NOT: ...
	NotToPrefixUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 166:2: prefix_unary_op -> BIT_NEG: ...
	BitNegToPrefixUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 167:2: prefix_unary_op -> SUB: ...
	SubToPrefixUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 170:2: prefix_unary_op -> MUL: ...
	MulToPrefixUnaryOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 173:2: prefix_unary_op -> BIT_AND: ...
	BitAndToPrefixUnaryOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 176:2: prefix_unary_expr -> postfix_unary_expr: ...
	PostfixUnaryExprToPrefixUnaryExpr(PostfixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 177:2: prefix_unary_expr -> prefix_op: ...
	PrefixOpToPrefixUnaryExpr(PrefixUnaryOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 180:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 181:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 182:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 183:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 184:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 185:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 188:2: mul_expr -> prefix_unary_expr: ...
	PrefixUnaryExprToMulExpr(PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 189:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 192:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 193:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 194:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 195:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 198:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 199:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 202:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 203:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 204:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 205:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 206:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 207:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 210:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 211:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 214:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 215:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 218:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 219:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 221:17: sequence_expr -> ...
	ToSequenceExpr(OrExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 224:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 225:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 228:2: optional_expression_or_implicit_struct -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct(ExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 229:2: optional_expression_or_implicit_struct -> nil: ...
	NilToOptionalExpressionOrImplicitStruct() (*GenericSymbol, error)

	// 232:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 233:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 234:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 237:2: op_one_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToOpOneAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 238:2: op_one_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToOpOneAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 241:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 242:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 243:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 244:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 245:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 246:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 247:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 248:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 249:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 250:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 251:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 254:2: expression_or_implicit_struct -> expression: ...
	ExpressionToExpressionOrImplicitStruct(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 255:2: expression_or_implicit_struct -> implicit_struct: ...
	ImplicitStructToExpressionOrImplicitStruct(ExpressionOrImplicitStruct_ *GenericSymbol, Comma_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 259:20: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ *GenericSymbol, Less_ *GenericSymbol, Identifier_ *GenericSymbol, Greater_ *GenericSymbol, StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 262:2: statement_body -> unsafe_statement: ...
	UnsafeStatementToStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 264:2: statement_body -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToStatementBody(ExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 266:2: statement_body -> async: ...
	AsyncToStatementBody(Async_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 269:2: statement_body -> jump: ...
	JumpToStatementBody(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 282:2: statement_body -> op_one_assign: ...
	OpOneAssignToStatementBody(AccessExpr_ *GenericSymbol, OpOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 283:2: statement_body -> binary_op_assign: ...
	BinaryOpAssignToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 286:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 287:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 290:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 291:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 294:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 295:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (*GenericSymbol, error)

	// 297:14: block_body -> ...
	ToBlockBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 298:14: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 303:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 304:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, BlockBody_2 *GenericSymbol) (*GenericSymbol, error)

	// 305:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 308:15: switch_expr -> ...
	ToSwitchExpr(Switch_ *GenericSymbol, Lbrace_ *GenericSymbol, Case_ *GenericSymbol, Default_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 311:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(For_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 312:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 317:2: expression -> sequence_expr: ...
	SequenceExprToExpression(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 318:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 319:2: expression -> switch_expr: ...
	SwitchExprToExpression(OptionalLabelDecl_ *GenericSymbol, SwitchExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 320:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ *GenericSymbol, LoopExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 323:2: generic_parameter -> unconstrained: ...
	UnconstrainedToGenericParameter(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 324:2: generic_parameter -> constrained: ...
	ConstrainedToGenericParameter(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 327:2: generic_parameters -> generic_parameter: ...
	GenericParameterToGenericParameters(GenericParameter_ *GenericSymbol) (*GenericSymbol, error)

	// 328:2: generic_parameters -> add: ...
	AddToGenericParameters(GenericParameters_ *GenericSymbol, Comma_ *GenericSymbol, GenericParameter_ *GenericSymbol) (*GenericSymbol, error)

	// 331:2: optional_generic_parameters -> generic_parameters: ...
	GenericParametersToOptionalGenericParameters(GenericParameters_ *GenericSymbol) (*GenericSymbol, error)

	// 332:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (*GenericSymbol, error)

	// 335:2: optional_generic_parameter_decl -> generic: ...
	GenericToOptionalGenericParameterDecl(DollarLbracket_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 336:2: optional_generic_parameter_decl -> nil: ...
	NilToOptionalGenericParameterDecl() (*GenericSymbol, error)

	// 339:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 340:2: atom_type -> explicit_struct_type: ...
	ExplicitStructTypeToAtomType(ExplicitStructType_ *GenericSymbol) (*GenericSymbol, error)

	// 341:2: atom_type -> implicit_struct_type: ...
	ImplicitStructTypeToAtomType(ImplicitStructType_ *GenericSymbol) (*GenericSymbol, error)

	// 342:2: atom_type -> explicit_enum_type: ...
	ExplicitEnumTypeToAtomType(ExplicitEnumType_ *GenericSymbol) (*GenericSymbol, error)

	// 343:2: atom_type -> trait_type: ...
	TraitTypeToAtomType(TraitType_ *GenericSymbol) (*GenericSymbol, error)

	// 346:2: traitable_type -> atom_type: ...
	AtomTypeToTraitableType(AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 347:2: traitable_type -> method_interface: ...
	MethodInterfaceToTraitableType(Exclaim_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 348:2: traitable_type -> trait: ...
	TraitToTraitableType(ExclaimExclaim_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 351:2: trait_mul_type -> traitable_type: ...
	TraitableTypeToTraitMulType(TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 352:2: trait_mul_type -> intersect: ...
	IntersectToTraitMulType(TraitMulType_ *GenericSymbol, Mul_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 355:2: trait_add_type -> trait_mul_type: ...
	TraitMulTypeToTraitAddType(TraitMulType_ *GenericSymbol) (*GenericSymbol, error)

	// 356:2: trait_add_type -> union: ...
	UnionToTraitAddType(TraitAddType_ *GenericSymbol, Add_ *GenericSymbol, TraitMulType_ *GenericSymbol) (*GenericSymbol, error)

	// 361:2: trait_add_type -> difference: ...
	DifferenceToTraitAddType(TraitAddType_ *GenericSymbol, Sub_ *GenericSymbol, TraitMulType_ *GenericSymbol) (*GenericSymbol, error)

	// 366:2: field_decl -> explicit: ...
	ExplicitToFieldDecl(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 367:2: field_decl -> implicit: ...
	ImplicitToFieldDecl(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 370:2: field_decls -> field_decl: ...
	FieldDeclToFieldDecls(FieldDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 371:2: field_decls -> implicit: ...
	ImplicitToFieldDecls(FieldDecls_ *GenericSymbol, Newlines_ *GenericSymbol, FieldDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 372:2: field_decls -> explicit: ...
	ExplicitToFieldDecls(FieldDecls_ *GenericSymbol, Comma_ *GenericSymbol, FieldDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 375:2: optional_field_decls -> field_decls: ...
	FieldDeclsToOptionalFieldDecls(FieldDecls_ *GenericSymbol) (*GenericSymbol, error)

	// 376:2: optional_field_decls -> nil: ...
	NilToOptionalFieldDecls() (*GenericSymbol, error)

	// 378:24: implicit_struct_type -> ...
	ToImplicitStructType(Lparen_ *GenericSymbol, OptionalFieldDecls_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 381:24: explicit_struct_type -> ...
	ToExplicitStructType(Struct_ *GenericSymbol, ImplicitStructType_ *GenericSymbol) (*GenericSymbol, error)

	// 383:14: trait_type -> ...
	ToTraitType(Trait_ *GenericSymbol, Lparen_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 386:22: explicit_enum_type -> ...
	ToExplicitEnumType(Enum_ *GenericSymbol, Lparen_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 389:2: func_type -> trait_add_type: ...
	TraitAddTypeToFuncType(TraitAddType_ *GenericSymbol) (*GenericSymbol, error)

	// 390:2: func_type -> reference: ...
	ReferenceToFuncType(BitAnd_ *GenericSymbol, TraitAddType_ *GenericSymbol) (*GenericSymbol, error)

	// 391:2: func_type -> explicit_func_type: ...
	ExplicitFuncTypeToFuncType(ExplicitFuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 394:2: value_type -> func_type: ...
	FuncTypeToValueType(FuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 395:2: value_type -> implicit_enum: ...
	ImplicitEnumToValueType(ValueType_ *GenericSymbol, Or_ *GenericSymbol, FuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 398:2: type_decl -> definition: ...
	DefinitionToTypeDecl(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameterDecl_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 399:2: type_decl -> alias: ...
	AliasToTypeDecl(Type_ *GenericSymbol, Identifier_ *GenericSymbol, Equal_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 402:2: parameter_decl -> field_decl: ...
	FieldDeclToParameterDecl(FieldDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 403:2: parameter_decl -> inferred: ...
	InferredToParameterDecl(Identifier_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 406:2: non_empty_parameters -> parameter_decl: ...
	ParameterDeclToNonEmptyParameters(ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 407:2: non_empty_parameters -> add: ...
	AddToNonEmptyParameters(NonEmptyParameters_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 410:2: vararg -> explicit: ...
	ExplicitToVararg(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 411:2: vararg -> implicit: ...
	ImplicitToVararg(Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 412:2: vararg -> inferred: ...
	InferredToVararg(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 415:2: optional_vararg -> vararg: ...
	VarargToOptionalVararg(Vararg_ *GenericSymbol) (*GenericSymbol, error)

	// 416:2: optional_vararg -> vararg2: ...
	Vararg2ToOptionalVararg(Vararg_ *GenericSymbol, Comma_ *GenericSymbol) (*GenericSymbol, error)

	// 417:2: optional_vararg -> nil: ...
	NilToOptionalVararg() (*GenericSymbol, error)

	// 420:2: parameters -> non_empty_parameters: ...
	NonEmptyParametersToParameters(NonEmptyParameters_ *GenericSymbol) (*GenericSymbol, error)

	// 421:2: parameters -> mixed: ...
	MixedToParameters(NonEmptyParameters_ *GenericSymbol, Comma_ *GenericSymbol, OptionalVararg_ *GenericSymbol) (*GenericSymbol, error)

	// 422:2: parameters -> optional_vararg: ...
	OptionalVarargToParameters(OptionalVararg_ *GenericSymbol) (*GenericSymbol, error)

	// 425:2: implicit_func_type -> typed: ...
	TypedToImplicitFuncType(Lparen_ *GenericSymbol, Parameters_ *GenericSymbol, Rparen_ *GenericSymbol, FuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 426:2: implicit_func_type -> inferred: ...
	InferredToImplicitFuncType(Lparen_ *GenericSymbol, Parameters_ *GenericSymbol, Rparen_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 427:2: implicit_func_type -> implicit_unit: ...
	ImplicitUnitToImplicitFuncType(Lparen_ *GenericSymbol, Parameters_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 429:22: explicit_func_type -> ...
	ToExplicitFuncType(Func_ *GenericSymbol, ImplicitFuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 432:2: optional_receiver -> receiver: ...
	ReceiverToOptionalReceiver(Lparen_ *GenericSymbol, ParameterDecl_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 433:2: optional_receiver -> nil: ...
	NilToOptionalReceiver() (*GenericSymbol, error)

	// 436:2: implicit_func_decl -> ...
	ToImplicitFuncDecl(Identifier_ *GenericSymbol, OptionalGenericParameterDecl_ *GenericSymbol, ImplicitFuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 439:2: explicit_func_decl -> ...
	ToExplicitFuncDecl(Func_ *GenericSymbol, OptionalReceiver_ *GenericSymbol, ImplicitFuncDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 441:12: func_def -> ...
	ToFuncDef(ExplicitFuncDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 444:2: package_decl -> no_spec: ...
	NoSpecToPackageDecl(Package_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 445:2: package_decl -> with_spec: ...
	WithSpecToPackageDecl(Package_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, PackageStatements_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 447:26: package_statement_body -> ...
	ToPackageStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 450:2: package_statement -> implicit: ...
	ImplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 451:2: package_statement -> explicit: ...
	ExplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 454:2: package_statements -> empty_list: ...
	EmptyListToPackageStatements() (*GenericSymbol, error)

	// 455:2: package_statements -> add: ...
	AddToPackageStatements(PackageStatements_ *GenericSymbol, PackageStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 459:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 460:2: lex_internal_tokens -> COMMENT: ...
	CommentToLexInternalTokens(Comment_ *GenericSymbol) (*GenericSymbol, error)
}

type ParseErrorHandler interface {
	Error(nextToken Token, parseStack _Stack) error
}

type DefaultParseErrorHandler struct{}

func (DefaultParseErrorHandler) Error(nextToken Token, stack _Stack) error {
	return fmt.Errorf(
		"Syntax error: unexpected symbol %v. Expecting %v (%v)",
		nextToken.Id(),
		ExpectedTerminals(stack[len(stack)-1].StateId),
		nextToken.Loc())
}

func ExpectedTerminals(id _StateId) []SymbolId {
	result := []SymbolId{}
	for key, _ := range _ActionTable {
		if key._StateId != id {
			continue
		}
		result = append(result, key.SymbolId)
	}

	sort.Slice(result, func(i int, j int) bool { return result[i] < result[j] })
	return result
}

func ParsePackageDecl(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParsePackageDeclWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParsePackageDeclWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	*GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State1)
	if err != nil {
		var errRetVal *GenericSymbol
		return errRetVal, err
	}
	return item.Generic_, nil
}

func ParseTypeDecl(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParseTypeDeclWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseTypeDeclWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	*GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State2)
	if err != nil {
		var errRetVal *GenericSymbol
		return errRetVal, err
	}
	return item.Generic_, nil
}

func ParseFuncDef(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParseFuncDefWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseFuncDefWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	*GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State3)
	if err != nil {
		var errRetVal *GenericSymbol
		return errRetVal, err
	}
	return item.Generic_, nil
}

func ParseExpression(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParseExpressionWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseExpressionWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	*GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State4)
	if err != nil {
		var errRetVal *GenericSymbol
		return errRetVal, err
	}
	return item.Generic_, nil
}

func ParseLexInternalTokens(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParseLexInternalTokensWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseLexInternalTokensWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	*GenericSymbol,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State5)
	if err != nil {
		var errRetVal *GenericSymbol
		return errRetVal, err
	}
	return item.Generic_, nil
}

// ================================================================
// Parser internal implementation
// User should normally avoid directly accessing the following code
// ================================================================

func _Parse(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler,
	startState _StateId) (
	*_StackItem,
	error) {

	stateStack := _Stack{
		// Note: we don't have to populate the start symbol since its value
		// is never accessed.
		&_StackItem{startState, nil},
	}

	symbolStack := &_PseudoSymbolStack{lexer: lexer}

	for {
		nextSymbol, err := symbolStack.Top()
		if err != nil {
			return nil, err
		}

		action, ok := _ActionTable.Get(
			stateStack[len(stateStack)-1].StateId,
			nextSymbol.Id())
		if !ok {
			return nil, errHandler.Error(nextSymbol, stateStack)
		}

		if action.ActionType == _ShiftAction {
			stateStack = append(stateStack, action.ShiftItem(nextSymbol))

			_, err = symbolStack.Pop()
			if err != nil {
				return nil, err
			}
		} else if action.ActionType == _ReduceAction {
			var reduceSymbol *Symbol
			stateStack, reduceSymbol, err = action.ReduceSymbol(
				reducer,
				stateStack)
			if err != nil {
				return nil, err
			}

			symbolStack.Push(reduceSymbol)
		} else if action.ActionType == _AcceptAction {
			if len(stateStack) != 2 {
				panic("This should never happen")
			}
			return stateStack[1], nil
		} else {
			panic("Unknown action type: " + action.ActionType.String())
		}
	}
}

func (i SymbolId) String() string {
	switch i {
	case _EndMarker:
		return "$"
	case _WildcardMarker:
		return "*"
	case SpacesToken:
		return "SPACES"
	case NewlinesToken:
		return "NEWLINES"
	case CommentToken:
		return "COMMENT"
	case IntegerLiteralToken:
		return "INTEGER_LITERAL"
	case FloatLiteralToken:
		return "FLOAT_LITERAL"
	case RuneLiteralToken:
		return "RUNE_LITERAL"
	case StringLiteralToken:
		return "STRING_LITERAL"
	case IdentifierToken:
		return "IDENTIFIER"
	case TrueToken:
		return "TRUE"
	case FalseToken:
		return "FALSE"
	case IfToken:
		return "IF"
	case ElseToken:
		return "ELSE"
	case SwitchToken:
		return "SWITCH"
	case CaseToken:
		return "CASE"
	case DefaultToken:
		return "DEFAULT"
	case ForToken:
		return "FOR"
	case ReturnToken:
		return "RETURN"
	case BreakToken:
		return "BREAK"
	case ContinueToken:
		return "CONTINUE"
	case PackageToken:
		return "PACKAGE"
	case UnsafeToken:
		return "UNSAFE"
	case TypeToken:
		return "TYPE"
	case StructToken:
		return "STRUCT"
	case EnumToken:
		return "ENUM"
	case TraitToken:
		return "TRAIT"
	case FuncToken:
		return "FUNC"
	case AsyncToken:
		return "ASYNC"
	case LbraceToken:
		return "LBRACE"
	case RbraceToken:
		return "RBRACE"
	case LparenToken:
		return "LPAREN"
	case RparenToken:
		return "RPAREN"
	case LbracketToken:
		return "LBRACKET"
	case RbracketToken:
		return "RBRACKET"
	case DotToken:
		return "DOT"
	case CommaToken:
		return "COMMA"
	case QuestionToken:
		return "QUESTION"
	case SemicolonToken:
		return "SEMICOLON"
	case DollarLbracketToken:
		return "DOLLAR_LBRACKET"
	case DotdotdotToken:
		return "DOTDOTDOT"
	case ExclaimToken:
		return "EXCLAIM"
	case ExclaimExclaimToken:
		return "EXCLAIM_EXCLAIM"
	case LabelDeclToken:
		return "LABEL_DECL"
	case JumpLabelToken:
		return "JUMP_LABEL"
	case AddAssignToken:
		return "ADD_ASSIGN"
	case SubAssignToken:
		return "SUB_ASSIGN"
	case MulAssignToken:
		return "MUL_ASSIGN"
	case DivAssignToken:
		return "DIV_ASSIGN"
	case ModAssignToken:
		return "MOD_ASSIGN"
	case AddOneAssignToken:
		return "ADD_ONE_ASSIGN"
	case SubOneAssignToken:
		return "SUB_ONE_ASSIGN"
	case BitNegAssignToken:
		return "BIT_NEG_ASSIGN"
	case BitAndAssignToken:
		return "BIT_AND_ASSIGN"
	case BitOrAssignToken:
		return "BIT_OR_ASSIGN"
	case BitXorAssignToken:
		return "BIT_XOR_ASSIGN"
	case BitLshiftAssignToken:
		return "BIT_LSHIFT_ASSIGN"
	case BitRshiftAssignToken:
		return "BIT_RSHIFT_ASSIGN"
	case NotToken:
		return "NOT"
	case AndToken:
		return "AND"
	case OrToken:
		return "OR"
	case AddToken:
		return "ADD"
	case SubToken:
		return "SUB"
	case MulToken:
		return "MUL"
	case DivToken:
		return "DIV"
	case ModToken:
		return "MOD"
	case BitNegToken:
		return "BIT_NEG"
	case BitAndToken:
		return "BIT_AND"
	case BitXorToken:
		return "BIT_XOR"
	case BitOrToken:
		return "BIT_OR"
	case BitLshiftToken:
		return "BIT_LSHIFT"
	case BitRshiftToken:
		return "BIT_RSHIFT"
	case EqualToken:
		return "EQUAL"
	case NotEqualToken:
		return "NOT_EQUAL"
	case LessToken:
		return "LESS"
	case LessOrEqualToken:
		return "LESS_OR_EQUAL"
	case GreaterToken:
		return "GREATER"
	case GreaterOrEqualToken:
		return "GREATER_OR_EQUAL"
	case LexErrorToken:
		return "LEX_ERROR"
	case LiteralType:
		return "literal"
	case AnonymousFuncExprType:
		return "anonymous_func_expr"
	case AnonymousStructExprType:
		return "anonymous_struct_expr"
	case AtomExprType:
		return "atom_expr"
	case GenericArgumentsType:
		return "generic_arguments"
	case OptionalGenericArgumentsType:
		return "optional_generic_arguments"
	case OptionalGenericBindingType:
		return "optional_generic_binding"
	case ArgumentType:
		return "argument"
	case ArgumentsType:
		return "arguments"
	case OptionalArgumentsType:
		return "optional_arguments"
	case CallExprType:
		return "call_expr"
	case AccessExprType:
		return "access_expr"
	case PostfixUnaryExprType:
		return "postfix_unary_expr"
	case PrefixUnaryOpType:
		return "prefix_unary_op"
	case PrefixUnaryExprType:
		return "prefix_unary_expr"
	case MulOpType:
		return "mul_op"
	case MulExprType:
		return "mul_expr"
	case AddOpType:
		return "add_op"
	case AddExprType:
		return "add_expr"
	case CmpOpType:
		return "cmp_op"
	case CmpExprType:
		return "cmp_expr"
	case AndExprType:
		return "and_expr"
	case OrExprType:
		return "or_expr"
	case SequenceExprType:
		return "sequence_expr"
	case OptionalJumpLabelType:
		return "optional_jump_label"
	case OptionalExpressionOrImplicitStructType:
		return "optional_expression_or_implicit_struct"
	case JumpTypeType:
		return "jump_type"
	case OpOneAssignType:
		return "op_one_assign"
	case BinaryOpAssignType:
		return "binary_op_assign"
	case ExpressionOrImplicitStructType:
		return "expression_or_implicit_struct"
	case UnsafeStatementType:
		return "unsafe_statement"
	case StatementBodyType:
		return "statement_body"
	case StatementType:
		return "statement"
	case StatementsType:
		return "statements"
	case OptionalLabelDeclType:
		return "optional_label_decl"
	case BlockBodyType:
		return "block_body"
	case BlockExprType:
		return "block_expr"
	case IfExprType:
		return "if_expr"
	case SwitchExprType:
		return "switch_expr"
	case LoopExprType:
		return "loop_expr"
	case ExpressionType:
		return "expression"
	case GenericParameterType:
		return "generic_parameter"
	case GenericParametersType:
		return "generic_parameters"
	case OptionalGenericParametersType:
		return "optional_generic_parameters"
	case OptionalGenericParameterDeclType:
		return "optional_generic_parameter_decl"
	case AtomTypeType:
		return "atom_type"
	case TraitableTypeType:
		return "traitable_type"
	case TraitMulTypeType:
		return "trait_mul_type"
	case TraitAddTypeType:
		return "trait_add_type"
	case FieldDeclType:
		return "field_decl"
	case FieldDeclsType:
		return "field_decls"
	case OptionalFieldDeclsType:
		return "optional_field_decls"
	case ImplicitStructTypeType:
		return "implicit_struct_type"
	case ExplicitStructTypeType:
		return "explicit_struct_type"
	case TraitTypeType:
		return "trait_type"
	case ExplicitEnumTypeType:
		return "explicit_enum_type"
	case FuncTypeType:
		return "func_type"
	case ValueTypeType:
		return "value_type"
	case TypeDeclType:
		return "type_decl"
	case ParameterDeclType:
		return "parameter_decl"
	case NonEmptyParametersType:
		return "non_empty_parameters"
	case VarargType:
		return "vararg"
	case OptionalVarargType:
		return "optional_vararg"
	case ParametersType:
		return "parameters"
	case ImplicitFuncTypeType:
		return "implicit_func_type"
	case ExplicitFuncTypeType:
		return "explicit_func_type"
	case OptionalReceiverType:
		return "optional_receiver"
	case ImplicitFuncDeclType:
		return "implicit_func_decl"
	case ExplicitFuncDeclType:
		return "explicit_func_decl"
	case FuncDefType:
		return "func_def"
	case PackageDeclType:
		return "package_decl"
	case PackageStatementBodyType:
		return "package_statement_body"
	case PackageStatementType:
		return "package_statement"
	case PackageStatementsType:
		return "package_statements"
	case LexInternalTokensType:
		return "lex_internal_tokens"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_EndMarker      = SymbolId(0)
	_WildcardMarker = SymbolId(-1)

	LiteralType                            = SymbolId(333)
	AnonymousFuncExprType                  = SymbolId(334)
	AnonymousStructExprType                = SymbolId(335)
	AtomExprType                           = SymbolId(336)
	GenericArgumentsType                   = SymbolId(337)
	OptionalGenericArgumentsType           = SymbolId(338)
	OptionalGenericBindingType             = SymbolId(339)
	ArgumentType                           = SymbolId(340)
	ArgumentsType                          = SymbolId(341)
	OptionalArgumentsType                  = SymbolId(342)
	CallExprType                           = SymbolId(343)
	AccessExprType                         = SymbolId(344)
	PostfixUnaryExprType                   = SymbolId(345)
	PrefixUnaryOpType                      = SymbolId(346)
	PrefixUnaryExprType                    = SymbolId(347)
	MulOpType                              = SymbolId(348)
	MulExprType                            = SymbolId(349)
	AddOpType                              = SymbolId(350)
	AddExprType                            = SymbolId(351)
	CmpOpType                              = SymbolId(352)
	CmpExprType                            = SymbolId(353)
	AndExprType                            = SymbolId(354)
	OrExprType                             = SymbolId(355)
	SequenceExprType                       = SymbolId(356)
	OptionalJumpLabelType                  = SymbolId(357)
	OptionalExpressionOrImplicitStructType = SymbolId(358)
	JumpTypeType                           = SymbolId(359)
	OpOneAssignType                        = SymbolId(360)
	BinaryOpAssignType                     = SymbolId(361)
	ExpressionOrImplicitStructType         = SymbolId(362)
	UnsafeStatementType                    = SymbolId(363)
	StatementBodyType                      = SymbolId(364)
	StatementType                          = SymbolId(365)
	StatementsType                         = SymbolId(366)
	OptionalLabelDeclType                  = SymbolId(367)
	BlockBodyType                          = SymbolId(368)
	BlockExprType                          = SymbolId(369)
	IfExprType                             = SymbolId(370)
	SwitchExprType                         = SymbolId(371)
	LoopExprType                           = SymbolId(372)
	ExpressionType                         = SymbolId(373)
	GenericParameterType                   = SymbolId(374)
	GenericParametersType                  = SymbolId(375)
	OptionalGenericParametersType          = SymbolId(376)
	OptionalGenericParameterDeclType       = SymbolId(377)
	AtomTypeType                           = SymbolId(378)
	TraitableTypeType                      = SymbolId(379)
	TraitMulTypeType                       = SymbolId(380)
	TraitAddTypeType                       = SymbolId(381)
	FieldDeclType                          = SymbolId(382)
	FieldDeclsType                         = SymbolId(383)
	OptionalFieldDeclsType                 = SymbolId(384)
	ImplicitStructTypeType                 = SymbolId(385)
	ExplicitStructTypeType                 = SymbolId(386)
	TraitTypeType                          = SymbolId(387)
	ExplicitEnumTypeType                   = SymbolId(388)
	FuncTypeType                           = SymbolId(389)
	ValueTypeType                          = SymbolId(390)
	TypeDeclType                           = SymbolId(391)
	ParameterDeclType                      = SymbolId(392)
	NonEmptyParametersType                 = SymbolId(393)
	VarargType                             = SymbolId(394)
	OptionalVarargType                     = SymbolId(395)
	ParametersType                         = SymbolId(396)
	ImplicitFuncTypeType                   = SymbolId(397)
	ExplicitFuncTypeType                   = SymbolId(398)
	OptionalReceiverType                   = SymbolId(399)
	ImplicitFuncDeclType                   = SymbolId(400)
	ExplicitFuncDeclType                   = SymbolId(401)
	FuncDefType                            = SymbolId(402)
	PackageDeclType                        = SymbolId(403)
	PackageStatementBodyType               = SymbolId(404)
	PackageStatementType                   = SymbolId(405)
	PackageStatementsType                  = SymbolId(406)
	LexInternalTokensType                  = SymbolId(407)
)

type _ActionType int

const (
	// NOTE: error action is implicit
	_ShiftAction  = _ActionType(0)
	_ReduceAction = _ActionType(1)
	_AcceptAction = _ActionType(2)
)

func (i _ActionType) String() string {
	switch i {
	case _ShiftAction:
		return "shift"
	case _ReduceAction:
		return "reduce"
	case _AcceptAction:
		return "accept"
	default:
		return fmt.Sprintf("?Unknown action %d?", int(i))
	}
}

type _ReduceType int

const (
	_ReduceTrueToLiteral                                                  = _ReduceType(1)
	_ReduceFalseToLiteral                                                 = _ReduceType(2)
	_ReduceIntegerLiteralToLiteral                                        = _ReduceType(3)
	_ReduceFloatLiteralToLiteral                                          = _ReduceType(4)
	_ReduceRuneLiteralToLiteral                                           = _ReduceType(5)
	_ReduceStringLiteralToLiteral                                         = _ReduceType(6)
	_ReduceToAnonymousFuncExpr                                            = _ReduceType(7)
	_ReduceExplicitToAnonymousStructExpr                                  = _ReduceType(8)
	_ReduceImplicitToAnonymousStructExpr                                  = _ReduceType(9)
	_ReduceLiteralToAtomExpr                                              = _ReduceType(10)
	_ReduceIdentifierToAtomExpr                                           = _ReduceType(11)
	_ReduceBlockExprToAtomExpr                                            = _ReduceType(12)
	_ReduceAnonymousFuncExprToAtomExpr                                    = _ReduceType(13)
	_ReduceAnonymousStructExprToAtomExpr                                  = _ReduceType(14)
	_ReduceLexErrorToAtomExpr                                             = _ReduceType(15)
	_ReduceValueTypeToGenericArguments                                    = _ReduceType(16)
	_ReduceAddToGenericArguments                                          = _ReduceType(17)
	_ReduceGenericArgumentsToOptionalGenericArguments                     = _ReduceType(18)
	_ReduceNilToOptionalGenericArguments                                  = _ReduceType(19)
	_ReduceBindingToOptionalGenericBinding                                = _ReduceType(20)
	_ReduceNilToOptionalGenericBinding                                    = _ReduceType(21)
	_ReducePositionalToArgument                                           = _ReduceType(22)
	_ReduceArgumentToArguments                                            = _ReduceType(23)
	_ReduceAddToArguments                                                 = _ReduceType(24)
	_ReduceArgumentsToOptionalArguments                                   = _ReduceType(25)
	_ReduceNilToOptionalArguments                                         = _ReduceType(26)
	_ReduceToCallExpr                                                     = _ReduceType(27)
	_ReduceAtomExprToAccessExpr                                           = _ReduceType(28)
	_ReduceAccessToAccessExpr                                             = _ReduceType(29)
	_ReduceCallExprToAccessExpr                                           = _ReduceType(30)
	_ReduceIndexToAccessExpr                                              = _ReduceType(31)
	_ReduceAccessExprToPostfixUnaryExpr                                   = _ReduceType(32)
	_ReduceQuestionToPostfixUnaryExpr                                     = _ReduceType(33)
	_ReduceNotToPrefixUnaryOp                                             = _ReduceType(34)
	_ReduceBitNegToPrefixUnaryOp                                          = _ReduceType(35)
	_ReduceSubToPrefixUnaryOp                                             = _ReduceType(36)
	_ReduceMulToPrefixUnaryOp                                             = _ReduceType(37)
	_ReduceBitAndToPrefixUnaryOp                                          = _ReduceType(38)
	_ReducePostfixUnaryExprToPrefixUnaryExpr                              = _ReduceType(39)
	_ReducePrefixOpToPrefixUnaryExpr                                      = _ReduceType(40)
	_ReduceMulToMulOp                                                     = _ReduceType(41)
	_ReduceDivToMulOp                                                     = _ReduceType(42)
	_ReduceModToMulOp                                                     = _ReduceType(43)
	_ReduceBitAndToMulOp                                                  = _ReduceType(44)
	_ReduceBitLshiftToMulOp                                               = _ReduceType(45)
	_ReduceBitRshiftToMulOp                                               = _ReduceType(46)
	_ReducePrefixUnaryExprToMulExpr                                       = _ReduceType(47)
	_ReduceOpToMulExpr                                                    = _ReduceType(48)
	_ReduceAddToAddOp                                                     = _ReduceType(49)
	_ReduceSubToAddOp                                                     = _ReduceType(50)
	_ReduceBitOrToAddOp                                                   = _ReduceType(51)
	_ReduceBitXorToAddOp                                                  = _ReduceType(52)
	_ReduceMulExprToAddExpr                                               = _ReduceType(53)
	_ReduceOpToAddExpr                                                    = _ReduceType(54)
	_ReduceEqualToCmpOp                                                   = _ReduceType(55)
	_ReduceNotEqualToCmpOp                                                = _ReduceType(56)
	_ReduceLessToCmpOp                                                    = _ReduceType(57)
	_ReduceLessOrEqualToCmpOp                                             = _ReduceType(58)
	_ReduceGreaterToCmpOp                                                 = _ReduceType(59)
	_ReduceGreaterOrEqualToCmpOp                                          = _ReduceType(60)
	_ReduceAddExprToCmpExpr                                               = _ReduceType(61)
	_ReduceOpToCmpExpr                                                    = _ReduceType(62)
	_ReduceCmpExprToAndExpr                                               = _ReduceType(63)
	_ReduceOpToAndExpr                                                    = _ReduceType(64)
	_ReduceAndExprToOrExpr                                                = _ReduceType(65)
	_ReduceOpToOrExpr                                                     = _ReduceType(66)
	_ReduceToSequenceExpr                                                 = _ReduceType(67)
	_ReduceJumpLabelToOptionalJumpLabel                                   = _ReduceType(68)
	_ReduceUnlabelledToOptionalJumpLabel                                  = _ReduceType(69)
	_ReduceExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct = _ReduceType(70)
	_ReduceNilToOptionalExpressionOrImplicitStruct                        = _ReduceType(71)
	_ReduceReturnToJumpType                                               = _ReduceType(72)
	_ReduceBreakToJumpType                                                = _ReduceType(73)
	_ReduceContinueToJumpType                                             = _ReduceType(74)
	_ReduceAddOneAssignToOpOneAssign                                      = _ReduceType(75)
	_ReduceSubOneAssignToOpOneAssign                                      = _ReduceType(76)
	_ReduceAddAssignToBinaryOpAssign                                      = _ReduceType(77)
	_ReduceSubAssignToBinaryOpAssign                                      = _ReduceType(78)
	_ReduceMulAssignToBinaryOpAssign                                      = _ReduceType(79)
	_ReduceDivAssignToBinaryOpAssign                                      = _ReduceType(80)
	_ReduceModAssignToBinaryOpAssign                                      = _ReduceType(81)
	_ReduceBitNegAssignToBinaryOpAssign                                   = _ReduceType(82)
	_ReduceBitAndAssignToBinaryOpAssign                                   = _ReduceType(83)
	_ReduceBitOrAssignToBinaryOpAssign                                    = _ReduceType(84)
	_ReduceBitXorAssignToBinaryOpAssign                                   = _ReduceType(85)
	_ReduceBitLshiftAssignToBinaryOpAssign                                = _ReduceType(86)
	_ReduceBitRshiftAssignToBinaryOpAssign                                = _ReduceType(87)
	_ReduceExpressionToExpressionOrImplicitStruct                         = _ReduceType(88)
	_ReduceImplicitStructToExpressionOrImplicitStruct                     = _ReduceType(89)
	_ReduceToUnsafeStatement                                              = _ReduceType(90)
	_ReduceUnsafeStatementToStatementBody                                 = _ReduceType(91)
	_ReduceExpressionOrImplicitStructToStatementBody                      = _ReduceType(92)
	_ReduceAsyncToStatementBody                                           = _ReduceType(93)
	_ReduceJumpToStatementBody                                            = _ReduceType(94)
	_ReduceOpOneAssignToStatementBody                                     = _ReduceType(95)
	_ReduceBinaryOpAssignToStatementBody                                  = _ReduceType(96)
	_ReduceImplicitToStatement                                            = _ReduceType(97)
	_ReduceExplicitToStatement                                            = _ReduceType(98)
	_ReduceEmptyListToStatements                                          = _ReduceType(99)
	_ReduceAddToStatements                                                = _ReduceType(100)
	_ReduceLabelDeclToOptionalLabelDecl                                   = _ReduceType(101)
	_ReduceUnlabelledToOptionalLabelDecl                                  = _ReduceType(102)
	_ReduceToBlockBody                                                    = _ReduceType(103)
	_ReduceToBlockExpr                                                    = _ReduceType(104)
	_ReduceNoElseToIfExpr                                                 = _ReduceType(105)
	_ReduceIfElseToIfExpr                                                 = _ReduceType(106)
	_ReduceMultiIfElseToIfExpr                                            = _ReduceType(107)
	_ReduceToSwitchExpr                                                   = _ReduceType(108)
	_ReduceInfiniteToLoopExpr                                             = _ReduceType(109)
	_ReduceWhileToLoopExpr                                                = _ReduceType(110)
	_ReduceSequenceExprToExpression                                       = _ReduceType(111)
	_ReduceIfExprToExpression                                             = _ReduceType(112)
	_ReduceSwitchExprToExpression                                         = _ReduceType(113)
	_ReduceLoopExprToExpression                                           = _ReduceType(114)
	_ReduceUnconstrainedToGenericParameter                                = _ReduceType(115)
	_ReduceConstrainedToGenericParameter                                  = _ReduceType(116)
	_ReduceGenericParameterToGenericParameters                            = _ReduceType(117)
	_ReduceAddToGenericParameters                                         = _ReduceType(118)
	_ReduceGenericParametersToOptionalGenericParameters                   = _ReduceType(119)
	_ReduceNilToOptionalGenericParameters                                 = _ReduceType(120)
	_ReduceGenericToOptionalGenericParameterDecl                          = _ReduceType(121)
	_ReduceNilToOptionalGenericParameterDecl                              = _ReduceType(122)
	_ReduceNamedToAtomType                                                = _ReduceType(123)
	_ReduceExplicitStructTypeToAtomType                                   = _ReduceType(124)
	_ReduceImplicitStructTypeToAtomType                                   = _ReduceType(125)
	_ReduceExplicitEnumTypeToAtomType                                     = _ReduceType(126)
	_ReduceTraitTypeToAtomType                                            = _ReduceType(127)
	_ReduceAtomTypeToTraitableType                                        = _ReduceType(128)
	_ReduceMethodInterfaceToTraitableType                                 = _ReduceType(129)
	_ReduceTraitToTraitableType                                           = _ReduceType(130)
	_ReduceTraitableTypeToTraitMulType                                    = _ReduceType(131)
	_ReduceIntersectToTraitMulType                                        = _ReduceType(132)
	_ReduceTraitMulTypeToTraitAddType                                     = _ReduceType(133)
	_ReduceUnionToTraitAddType                                            = _ReduceType(134)
	_ReduceDifferenceToTraitAddType                                       = _ReduceType(135)
	_ReduceExplicitToFieldDecl                                            = _ReduceType(136)
	_ReduceImplicitToFieldDecl                                            = _ReduceType(137)
	_ReduceFieldDeclToFieldDecls                                          = _ReduceType(138)
	_ReduceImplicitToFieldDecls                                           = _ReduceType(139)
	_ReduceExplicitToFieldDecls                                           = _ReduceType(140)
	_ReduceFieldDeclsToOptionalFieldDecls                                 = _ReduceType(141)
	_ReduceNilToOptionalFieldDecls                                        = _ReduceType(142)
	_ReduceToImplicitStructType                                           = _ReduceType(143)
	_ReduceToExplicitStructType                                           = _ReduceType(144)
	_ReduceToTraitType                                                    = _ReduceType(145)
	_ReduceToExplicitEnumType                                             = _ReduceType(146)
	_ReduceTraitAddTypeToFuncType                                         = _ReduceType(147)
	_ReduceReferenceToFuncType                                            = _ReduceType(148)
	_ReduceExplicitFuncTypeToFuncType                                     = _ReduceType(149)
	_ReduceFuncTypeToValueType                                            = _ReduceType(150)
	_ReduceImplicitEnumToValueType                                        = _ReduceType(151)
	_ReduceDefinitionToTypeDecl                                           = _ReduceType(152)
	_ReduceAliasToTypeDecl                                                = _ReduceType(153)
	_ReduceFieldDeclToParameterDecl                                       = _ReduceType(154)
	_ReduceInferredToParameterDecl                                        = _ReduceType(155)
	_ReduceParameterDeclToNonEmptyParameters                              = _ReduceType(156)
	_ReduceAddToNonEmptyParameters                                        = _ReduceType(157)
	_ReduceExplicitToVararg                                               = _ReduceType(158)
	_ReduceImplicitToVararg                                               = _ReduceType(159)
	_ReduceInferredToVararg                                               = _ReduceType(160)
	_ReduceVarargToOptionalVararg                                         = _ReduceType(161)
	_ReduceVararg2ToOptionalVararg                                        = _ReduceType(162)
	_ReduceNilToOptionalVararg                                            = _ReduceType(163)
	_ReduceNonEmptyParametersToParameters                                 = _ReduceType(164)
	_ReduceMixedToParameters                                              = _ReduceType(165)
	_ReduceOptionalVarargToParameters                                     = _ReduceType(166)
	_ReduceTypedToImplicitFuncType                                        = _ReduceType(167)
	_ReduceInferredToImplicitFuncType                                     = _ReduceType(168)
	_ReduceImplicitUnitToImplicitFuncType                                 = _ReduceType(169)
	_ReduceToExplicitFuncType                                             = _ReduceType(170)
	_ReduceReceiverToOptionalReceiver                                     = _ReduceType(171)
	_ReduceNilToOptionalReceiver                                          = _ReduceType(172)
	_ReduceToImplicitFuncDecl                                             = _ReduceType(173)
	_ReduceToExplicitFuncDecl                                             = _ReduceType(174)
	_ReduceToFuncDef                                                      = _ReduceType(175)
	_ReduceNoSpecToPackageDecl                                            = _ReduceType(176)
	_ReduceWithSpecToPackageDecl                                          = _ReduceType(177)
	_ReduceToPackageStatementBody                                         = _ReduceType(178)
	_ReduceImplicitToPackageStatement                                     = _ReduceType(179)
	_ReduceExplicitToPackageStatement                                     = _ReduceType(180)
	_ReduceEmptyListToPackageStatements                                   = _ReduceType(181)
	_ReduceAddToPackageStatements                                         = _ReduceType(182)
	_ReduceSpacesToLexInternalTokens                                      = _ReduceType(183)
	_ReduceCommentToLexInternalTokens                                     = _ReduceType(184)
)

func (i _ReduceType) String() string {
	switch i {
	case _ReduceTrueToLiteral:
		return "TrueToLiteral"
	case _ReduceFalseToLiteral:
		return "FalseToLiteral"
	case _ReduceIntegerLiteralToLiteral:
		return "IntegerLiteralToLiteral"
	case _ReduceFloatLiteralToLiteral:
		return "FloatLiteralToLiteral"
	case _ReduceRuneLiteralToLiteral:
		return "RuneLiteralToLiteral"
	case _ReduceStringLiteralToLiteral:
		return "StringLiteralToLiteral"
	case _ReduceToAnonymousFuncExpr:
		return "ToAnonymousFuncExpr"
	case _ReduceExplicitToAnonymousStructExpr:
		return "ExplicitToAnonymousStructExpr"
	case _ReduceImplicitToAnonymousStructExpr:
		return "ImplicitToAnonymousStructExpr"
	case _ReduceLiteralToAtomExpr:
		return "LiteralToAtomExpr"
	case _ReduceIdentifierToAtomExpr:
		return "IdentifierToAtomExpr"
	case _ReduceBlockExprToAtomExpr:
		return "BlockExprToAtomExpr"
	case _ReduceAnonymousFuncExprToAtomExpr:
		return "AnonymousFuncExprToAtomExpr"
	case _ReduceAnonymousStructExprToAtomExpr:
		return "AnonymousStructExprToAtomExpr"
	case _ReduceLexErrorToAtomExpr:
		return "LexErrorToAtomExpr"
	case _ReduceValueTypeToGenericArguments:
		return "ValueTypeToGenericArguments"
	case _ReduceAddToGenericArguments:
		return "AddToGenericArguments"
	case _ReduceGenericArgumentsToOptionalGenericArguments:
		return "GenericArgumentsToOptionalGenericArguments"
	case _ReduceNilToOptionalGenericArguments:
		return "NilToOptionalGenericArguments"
	case _ReduceBindingToOptionalGenericBinding:
		return "BindingToOptionalGenericBinding"
	case _ReduceNilToOptionalGenericBinding:
		return "NilToOptionalGenericBinding"
	case _ReducePositionalToArgument:
		return "PositionalToArgument"
	case _ReduceArgumentToArguments:
		return "ArgumentToArguments"
	case _ReduceAddToArguments:
		return "AddToArguments"
	case _ReduceArgumentsToOptionalArguments:
		return "ArgumentsToOptionalArguments"
	case _ReduceNilToOptionalArguments:
		return "NilToOptionalArguments"
	case _ReduceToCallExpr:
		return "ToCallExpr"
	case _ReduceAtomExprToAccessExpr:
		return "AtomExprToAccessExpr"
	case _ReduceAccessToAccessExpr:
		return "AccessToAccessExpr"
	case _ReduceCallExprToAccessExpr:
		return "CallExprToAccessExpr"
	case _ReduceIndexToAccessExpr:
		return "IndexToAccessExpr"
	case _ReduceAccessExprToPostfixUnaryExpr:
		return "AccessExprToPostfixUnaryExpr"
	case _ReduceQuestionToPostfixUnaryExpr:
		return "QuestionToPostfixUnaryExpr"
	case _ReduceNotToPrefixUnaryOp:
		return "NotToPrefixUnaryOp"
	case _ReduceBitNegToPrefixUnaryOp:
		return "BitNegToPrefixUnaryOp"
	case _ReduceSubToPrefixUnaryOp:
		return "SubToPrefixUnaryOp"
	case _ReduceMulToPrefixUnaryOp:
		return "MulToPrefixUnaryOp"
	case _ReduceBitAndToPrefixUnaryOp:
		return "BitAndToPrefixUnaryOp"
	case _ReducePostfixUnaryExprToPrefixUnaryExpr:
		return "PostfixUnaryExprToPrefixUnaryExpr"
	case _ReducePrefixOpToPrefixUnaryExpr:
		return "PrefixOpToPrefixUnaryExpr"
	case _ReduceMulToMulOp:
		return "MulToMulOp"
	case _ReduceDivToMulOp:
		return "DivToMulOp"
	case _ReduceModToMulOp:
		return "ModToMulOp"
	case _ReduceBitAndToMulOp:
		return "BitAndToMulOp"
	case _ReduceBitLshiftToMulOp:
		return "BitLshiftToMulOp"
	case _ReduceBitRshiftToMulOp:
		return "BitRshiftToMulOp"
	case _ReducePrefixUnaryExprToMulExpr:
		return "PrefixUnaryExprToMulExpr"
	case _ReduceOpToMulExpr:
		return "OpToMulExpr"
	case _ReduceAddToAddOp:
		return "AddToAddOp"
	case _ReduceSubToAddOp:
		return "SubToAddOp"
	case _ReduceBitOrToAddOp:
		return "BitOrToAddOp"
	case _ReduceBitXorToAddOp:
		return "BitXorToAddOp"
	case _ReduceMulExprToAddExpr:
		return "MulExprToAddExpr"
	case _ReduceOpToAddExpr:
		return "OpToAddExpr"
	case _ReduceEqualToCmpOp:
		return "EqualToCmpOp"
	case _ReduceNotEqualToCmpOp:
		return "NotEqualToCmpOp"
	case _ReduceLessToCmpOp:
		return "LessToCmpOp"
	case _ReduceLessOrEqualToCmpOp:
		return "LessOrEqualToCmpOp"
	case _ReduceGreaterToCmpOp:
		return "GreaterToCmpOp"
	case _ReduceGreaterOrEqualToCmpOp:
		return "GreaterOrEqualToCmpOp"
	case _ReduceAddExprToCmpExpr:
		return "AddExprToCmpExpr"
	case _ReduceOpToCmpExpr:
		return "OpToCmpExpr"
	case _ReduceCmpExprToAndExpr:
		return "CmpExprToAndExpr"
	case _ReduceOpToAndExpr:
		return "OpToAndExpr"
	case _ReduceAndExprToOrExpr:
		return "AndExprToOrExpr"
	case _ReduceOpToOrExpr:
		return "OpToOrExpr"
	case _ReduceToSequenceExpr:
		return "ToSequenceExpr"
	case _ReduceJumpLabelToOptionalJumpLabel:
		return "JumpLabelToOptionalJumpLabel"
	case _ReduceUnlabelledToOptionalJumpLabel:
		return "UnlabelledToOptionalJumpLabel"
	case _ReduceExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct:
		return "ExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct"
	case _ReduceNilToOptionalExpressionOrImplicitStruct:
		return "NilToOptionalExpressionOrImplicitStruct"
	case _ReduceReturnToJumpType:
		return "ReturnToJumpType"
	case _ReduceBreakToJumpType:
		return "BreakToJumpType"
	case _ReduceContinueToJumpType:
		return "ContinueToJumpType"
	case _ReduceAddOneAssignToOpOneAssign:
		return "AddOneAssignToOpOneAssign"
	case _ReduceSubOneAssignToOpOneAssign:
		return "SubOneAssignToOpOneAssign"
	case _ReduceAddAssignToBinaryOpAssign:
		return "AddAssignToBinaryOpAssign"
	case _ReduceSubAssignToBinaryOpAssign:
		return "SubAssignToBinaryOpAssign"
	case _ReduceMulAssignToBinaryOpAssign:
		return "MulAssignToBinaryOpAssign"
	case _ReduceDivAssignToBinaryOpAssign:
		return "DivAssignToBinaryOpAssign"
	case _ReduceModAssignToBinaryOpAssign:
		return "ModAssignToBinaryOpAssign"
	case _ReduceBitNegAssignToBinaryOpAssign:
		return "BitNegAssignToBinaryOpAssign"
	case _ReduceBitAndAssignToBinaryOpAssign:
		return "BitAndAssignToBinaryOpAssign"
	case _ReduceBitOrAssignToBinaryOpAssign:
		return "BitOrAssignToBinaryOpAssign"
	case _ReduceBitXorAssignToBinaryOpAssign:
		return "BitXorAssignToBinaryOpAssign"
	case _ReduceBitLshiftAssignToBinaryOpAssign:
		return "BitLshiftAssignToBinaryOpAssign"
	case _ReduceBitRshiftAssignToBinaryOpAssign:
		return "BitRshiftAssignToBinaryOpAssign"
	case _ReduceExpressionToExpressionOrImplicitStruct:
		return "ExpressionToExpressionOrImplicitStruct"
	case _ReduceImplicitStructToExpressionOrImplicitStruct:
		return "ImplicitStructToExpressionOrImplicitStruct"
	case _ReduceToUnsafeStatement:
		return "ToUnsafeStatement"
	case _ReduceUnsafeStatementToStatementBody:
		return "UnsafeStatementToStatementBody"
	case _ReduceExpressionOrImplicitStructToStatementBody:
		return "ExpressionOrImplicitStructToStatementBody"
	case _ReduceAsyncToStatementBody:
		return "AsyncToStatementBody"
	case _ReduceJumpToStatementBody:
		return "JumpToStatementBody"
	case _ReduceOpOneAssignToStatementBody:
		return "OpOneAssignToStatementBody"
	case _ReduceBinaryOpAssignToStatementBody:
		return "BinaryOpAssignToStatementBody"
	case _ReduceImplicitToStatement:
		return "ImplicitToStatement"
	case _ReduceExplicitToStatement:
		return "ExplicitToStatement"
	case _ReduceEmptyListToStatements:
		return "EmptyListToStatements"
	case _ReduceAddToStatements:
		return "AddToStatements"
	case _ReduceLabelDeclToOptionalLabelDecl:
		return "LabelDeclToOptionalLabelDecl"
	case _ReduceUnlabelledToOptionalLabelDecl:
		return "UnlabelledToOptionalLabelDecl"
	case _ReduceToBlockBody:
		return "ToBlockBody"
	case _ReduceToBlockExpr:
		return "ToBlockExpr"
	case _ReduceNoElseToIfExpr:
		return "NoElseToIfExpr"
	case _ReduceIfElseToIfExpr:
		return "IfElseToIfExpr"
	case _ReduceMultiIfElseToIfExpr:
		return "MultiIfElseToIfExpr"
	case _ReduceToSwitchExpr:
		return "ToSwitchExpr"
	case _ReduceInfiniteToLoopExpr:
		return "InfiniteToLoopExpr"
	case _ReduceWhileToLoopExpr:
		return "WhileToLoopExpr"
	case _ReduceSequenceExprToExpression:
		return "SequenceExprToExpression"
	case _ReduceIfExprToExpression:
		return "IfExprToExpression"
	case _ReduceSwitchExprToExpression:
		return "SwitchExprToExpression"
	case _ReduceLoopExprToExpression:
		return "LoopExprToExpression"
	case _ReduceUnconstrainedToGenericParameter:
		return "UnconstrainedToGenericParameter"
	case _ReduceConstrainedToGenericParameter:
		return "ConstrainedToGenericParameter"
	case _ReduceGenericParameterToGenericParameters:
		return "GenericParameterToGenericParameters"
	case _ReduceAddToGenericParameters:
		return "AddToGenericParameters"
	case _ReduceGenericParametersToOptionalGenericParameters:
		return "GenericParametersToOptionalGenericParameters"
	case _ReduceNilToOptionalGenericParameters:
		return "NilToOptionalGenericParameters"
	case _ReduceGenericToOptionalGenericParameterDecl:
		return "GenericToOptionalGenericParameterDecl"
	case _ReduceNilToOptionalGenericParameterDecl:
		return "NilToOptionalGenericParameterDecl"
	case _ReduceNamedToAtomType:
		return "NamedToAtomType"
	case _ReduceExplicitStructTypeToAtomType:
		return "ExplicitStructTypeToAtomType"
	case _ReduceImplicitStructTypeToAtomType:
		return "ImplicitStructTypeToAtomType"
	case _ReduceExplicitEnumTypeToAtomType:
		return "ExplicitEnumTypeToAtomType"
	case _ReduceTraitTypeToAtomType:
		return "TraitTypeToAtomType"
	case _ReduceAtomTypeToTraitableType:
		return "AtomTypeToTraitableType"
	case _ReduceMethodInterfaceToTraitableType:
		return "MethodInterfaceToTraitableType"
	case _ReduceTraitToTraitableType:
		return "TraitToTraitableType"
	case _ReduceTraitableTypeToTraitMulType:
		return "TraitableTypeToTraitMulType"
	case _ReduceIntersectToTraitMulType:
		return "IntersectToTraitMulType"
	case _ReduceTraitMulTypeToTraitAddType:
		return "TraitMulTypeToTraitAddType"
	case _ReduceUnionToTraitAddType:
		return "UnionToTraitAddType"
	case _ReduceDifferenceToTraitAddType:
		return "DifferenceToTraitAddType"
	case _ReduceExplicitToFieldDecl:
		return "ExplicitToFieldDecl"
	case _ReduceImplicitToFieldDecl:
		return "ImplicitToFieldDecl"
	case _ReduceFieldDeclToFieldDecls:
		return "FieldDeclToFieldDecls"
	case _ReduceImplicitToFieldDecls:
		return "ImplicitToFieldDecls"
	case _ReduceExplicitToFieldDecls:
		return "ExplicitToFieldDecls"
	case _ReduceFieldDeclsToOptionalFieldDecls:
		return "FieldDeclsToOptionalFieldDecls"
	case _ReduceNilToOptionalFieldDecls:
		return "NilToOptionalFieldDecls"
	case _ReduceToImplicitStructType:
		return "ToImplicitStructType"
	case _ReduceToExplicitStructType:
		return "ToExplicitStructType"
	case _ReduceToTraitType:
		return "ToTraitType"
	case _ReduceToExplicitEnumType:
		return "ToExplicitEnumType"
	case _ReduceTraitAddTypeToFuncType:
		return "TraitAddTypeToFuncType"
	case _ReduceReferenceToFuncType:
		return "ReferenceToFuncType"
	case _ReduceExplicitFuncTypeToFuncType:
		return "ExplicitFuncTypeToFuncType"
	case _ReduceFuncTypeToValueType:
		return "FuncTypeToValueType"
	case _ReduceImplicitEnumToValueType:
		return "ImplicitEnumToValueType"
	case _ReduceDefinitionToTypeDecl:
		return "DefinitionToTypeDecl"
	case _ReduceAliasToTypeDecl:
		return "AliasToTypeDecl"
	case _ReduceFieldDeclToParameterDecl:
		return "FieldDeclToParameterDecl"
	case _ReduceInferredToParameterDecl:
		return "InferredToParameterDecl"
	case _ReduceParameterDeclToNonEmptyParameters:
		return "ParameterDeclToNonEmptyParameters"
	case _ReduceAddToNonEmptyParameters:
		return "AddToNonEmptyParameters"
	case _ReduceExplicitToVararg:
		return "ExplicitToVararg"
	case _ReduceImplicitToVararg:
		return "ImplicitToVararg"
	case _ReduceInferredToVararg:
		return "InferredToVararg"
	case _ReduceVarargToOptionalVararg:
		return "VarargToOptionalVararg"
	case _ReduceVararg2ToOptionalVararg:
		return "Vararg2ToOptionalVararg"
	case _ReduceNilToOptionalVararg:
		return "NilToOptionalVararg"
	case _ReduceNonEmptyParametersToParameters:
		return "NonEmptyParametersToParameters"
	case _ReduceMixedToParameters:
		return "MixedToParameters"
	case _ReduceOptionalVarargToParameters:
		return "OptionalVarargToParameters"
	case _ReduceTypedToImplicitFuncType:
		return "TypedToImplicitFuncType"
	case _ReduceInferredToImplicitFuncType:
		return "InferredToImplicitFuncType"
	case _ReduceImplicitUnitToImplicitFuncType:
		return "ImplicitUnitToImplicitFuncType"
	case _ReduceToExplicitFuncType:
		return "ToExplicitFuncType"
	case _ReduceReceiverToOptionalReceiver:
		return "ReceiverToOptionalReceiver"
	case _ReduceNilToOptionalReceiver:
		return "NilToOptionalReceiver"
	case _ReduceToImplicitFuncDecl:
		return "ToImplicitFuncDecl"
	case _ReduceToExplicitFuncDecl:
		return "ToExplicitFuncDecl"
	case _ReduceToFuncDef:
		return "ToFuncDef"
	case _ReduceNoSpecToPackageDecl:
		return "NoSpecToPackageDecl"
	case _ReduceWithSpecToPackageDecl:
		return "WithSpecToPackageDecl"
	case _ReduceToPackageStatementBody:
		return "ToPackageStatementBody"
	case _ReduceImplicitToPackageStatement:
		return "ImplicitToPackageStatement"
	case _ReduceExplicitToPackageStatement:
		return "ExplicitToPackageStatement"
	case _ReduceEmptyListToPackageStatements:
		return "EmptyListToPackageStatements"
	case _ReduceAddToPackageStatements:
		return "AddToPackageStatements"
	case _ReduceSpacesToLexInternalTokens:
		return "SpacesToLexInternalTokens"
	case _ReduceCommentToLexInternalTokens:
		return "CommentToLexInternalTokens"
	default:
		return fmt.Sprintf("?unknown reduce type %d?", int(i))
	}
}

type _StateId int

func (id _StateId) String() string {
	return fmt.Sprintf("State %d", int(id))
}

const (
	_State1   = _StateId(1)
	_State2   = _StateId(2)
	_State3   = _StateId(3)
	_State4   = _StateId(4)
	_State5   = _StateId(5)
	_State6   = _StateId(6)
	_State7   = _StateId(7)
	_State8   = _StateId(8)
	_State9   = _StateId(9)
	_State10  = _StateId(10)
	_State11  = _StateId(11)
	_State12  = _StateId(12)
	_State13  = _StateId(13)
	_State14  = _StateId(14)
	_State15  = _StateId(15)
	_State16  = _StateId(16)
	_State17  = _StateId(17)
	_State18  = _StateId(18)
	_State19  = _StateId(19)
	_State20  = _StateId(20)
	_State21  = _StateId(21)
	_State22  = _StateId(22)
	_State23  = _StateId(23)
	_State24  = _StateId(24)
	_State25  = _StateId(25)
	_State26  = _StateId(26)
	_State27  = _StateId(27)
	_State28  = _StateId(28)
	_State29  = _StateId(29)
	_State30  = _StateId(30)
	_State31  = _StateId(31)
	_State32  = _StateId(32)
	_State33  = _StateId(33)
	_State34  = _StateId(34)
	_State35  = _StateId(35)
	_State36  = _StateId(36)
	_State37  = _StateId(37)
	_State38  = _StateId(38)
	_State39  = _StateId(39)
	_State40  = _StateId(40)
	_State41  = _StateId(41)
	_State42  = _StateId(42)
	_State43  = _StateId(43)
	_State44  = _StateId(44)
	_State45  = _StateId(45)
	_State46  = _StateId(46)
	_State47  = _StateId(47)
	_State48  = _StateId(48)
	_State49  = _StateId(49)
	_State50  = _StateId(50)
	_State51  = _StateId(51)
	_State52  = _StateId(52)
	_State53  = _StateId(53)
	_State54  = _StateId(54)
	_State55  = _StateId(55)
	_State56  = _StateId(56)
	_State57  = _StateId(57)
	_State58  = _StateId(58)
	_State59  = _StateId(59)
	_State60  = _StateId(60)
	_State61  = _StateId(61)
	_State62  = _StateId(62)
	_State63  = _StateId(63)
	_State64  = _StateId(64)
	_State65  = _StateId(65)
	_State66  = _StateId(66)
	_State67  = _StateId(67)
	_State68  = _StateId(68)
	_State69  = _StateId(69)
	_State70  = _StateId(70)
	_State71  = _StateId(71)
	_State72  = _StateId(72)
	_State73  = _StateId(73)
	_State74  = _StateId(74)
	_State75  = _StateId(75)
	_State76  = _StateId(76)
	_State77  = _StateId(77)
	_State78  = _StateId(78)
	_State79  = _StateId(79)
	_State80  = _StateId(80)
	_State81  = _StateId(81)
	_State82  = _StateId(82)
	_State83  = _StateId(83)
	_State84  = _StateId(84)
	_State85  = _StateId(85)
	_State86  = _StateId(86)
	_State87  = _StateId(87)
	_State88  = _StateId(88)
	_State89  = _StateId(89)
	_State90  = _StateId(90)
	_State91  = _StateId(91)
	_State92  = _StateId(92)
	_State93  = _StateId(93)
	_State94  = _StateId(94)
	_State95  = _StateId(95)
	_State96  = _StateId(96)
	_State97  = _StateId(97)
	_State98  = _StateId(98)
	_State99  = _StateId(99)
	_State100 = _StateId(100)
	_State101 = _StateId(101)
	_State102 = _StateId(102)
	_State103 = _StateId(103)
	_State104 = _StateId(104)
	_State105 = _StateId(105)
	_State106 = _StateId(106)
	_State107 = _StateId(107)
	_State108 = _StateId(108)
	_State109 = _StateId(109)
	_State110 = _StateId(110)
	_State111 = _StateId(111)
	_State112 = _StateId(112)
	_State113 = _StateId(113)
	_State114 = _StateId(114)
	_State115 = _StateId(115)
	_State116 = _StateId(116)
	_State117 = _StateId(117)
	_State118 = _StateId(118)
	_State119 = _StateId(119)
	_State120 = _StateId(120)
	_State121 = _StateId(121)
	_State122 = _StateId(122)
	_State123 = _StateId(123)
	_State124 = _StateId(124)
	_State125 = _StateId(125)
	_State126 = _StateId(126)
	_State127 = _StateId(127)
	_State128 = _StateId(128)
	_State129 = _StateId(129)
	_State130 = _StateId(130)
	_State131 = _StateId(131)
	_State132 = _StateId(132)
	_State133 = _StateId(133)
	_State134 = _StateId(134)
	_State135 = _StateId(135)
	_State136 = _StateId(136)
	_State137 = _StateId(137)
	_State138 = _StateId(138)
	_State139 = _StateId(139)
	_State140 = _StateId(140)
	_State141 = _StateId(141)
	_State142 = _StateId(142)
	_State143 = _StateId(143)
	_State144 = _StateId(144)
	_State145 = _StateId(145)
	_State146 = _StateId(146)
	_State147 = _StateId(147)
	_State148 = _StateId(148)
	_State149 = _StateId(149)
	_State150 = _StateId(150)
	_State151 = _StateId(151)
	_State152 = _StateId(152)
	_State153 = _StateId(153)
	_State154 = _StateId(154)
	_State155 = _StateId(155)
	_State156 = _StateId(156)
	_State157 = _StateId(157)
	_State158 = _StateId(158)
	_State159 = _StateId(159)
	_State160 = _StateId(160)
	_State161 = _StateId(161)
	_State162 = _StateId(162)
	_State163 = _StateId(163)
	_State164 = _StateId(164)
	_State165 = _StateId(165)
	_State166 = _StateId(166)
	_State167 = _StateId(167)
	_State168 = _StateId(168)
	_State169 = _StateId(169)
	_State170 = _StateId(170)
	_State171 = _StateId(171)
	_State172 = _StateId(172)
	_State173 = _StateId(173)
	_State174 = _StateId(174)
	_State175 = _StateId(175)
	_State176 = _StateId(176)
	_State177 = _StateId(177)
	_State178 = _StateId(178)
	_State179 = _StateId(179)
	_State180 = _StateId(180)
	_State181 = _StateId(181)
	_State182 = _StateId(182)
	_State183 = _StateId(183)
	_State184 = _StateId(184)
	_State185 = _StateId(185)
	_State186 = _StateId(186)
	_State187 = _StateId(187)
	_State188 = _StateId(188)
	_State189 = _StateId(189)
	_State190 = _StateId(190)
	_State191 = _StateId(191)
	_State192 = _StateId(192)
	_State193 = _StateId(193)
	_State194 = _StateId(194)
	_State195 = _StateId(195)
	_State196 = _StateId(196)
	_State197 = _StateId(197)
	_State198 = _StateId(198)
	_State199 = _StateId(199)
	_State200 = _StateId(200)
	_State201 = _StateId(201)
	_State202 = _StateId(202)
	_State203 = _StateId(203)
	_State204 = _StateId(204)
	_State205 = _StateId(205)
	_State206 = _StateId(206)
	_State207 = _StateId(207)
	_State208 = _StateId(208)
	_State209 = _StateId(209)
	_State210 = _StateId(210)
	_State211 = _StateId(211)
	_State212 = _StateId(212)
	_State213 = _StateId(213)
	_State214 = _StateId(214)
	_State215 = _StateId(215)
	_State216 = _StateId(216)
	_State217 = _StateId(217)
	_State218 = _StateId(218)
	_State219 = _StateId(219)
	_State220 = _StateId(220)
	_State221 = _StateId(221)
	_State222 = _StateId(222)
	_State223 = _StateId(223)
	_State224 = _StateId(224)
	_State225 = _StateId(225)
	_State226 = _StateId(226)
	_State227 = _StateId(227)
	_State228 = _StateId(228)
	_State229 = _StateId(229)
	_State230 = _StateId(230)
	_State231 = _StateId(231)
	_State232 = _StateId(232)
	_State233 = _StateId(233)
	_State234 = _StateId(234)
	_State235 = _StateId(235)
	_State236 = _StateId(236)
	_State237 = _StateId(237)
	_State238 = _StateId(238)
	_State239 = _StateId(239)
	_State240 = _StateId(240)
	_State241 = _StateId(241)
	_State242 = _StateId(242)
	_State243 = _StateId(243)
	_State244 = _StateId(244)
	_State245 = _StateId(245)
	_State246 = _StateId(246)
	_State247 = _StateId(247)
	_State248 = _StateId(248)
	_State249 = _StateId(249)
	_State250 = _StateId(250)
	_State251 = _StateId(251)
	_State252 = _StateId(252)
	_State253 = _StateId(253)
	_State254 = _StateId(254)
	_State255 = _StateId(255)
	_State256 = _StateId(256)
	_State257 = _StateId(257)
	_State258 = _StateId(258)
	_State259 = _StateId(259)
	_State260 = _StateId(260)
	_State261 = _StateId(261)
	_State262 = _StateId(262)
	_State263 = _StateId(263)
	_State264 = _StateId(264)
	_State265 = _StateId(265)
	_State266 = _StateId(266)
	_State267 = _StateId(267)
	_State268 = _StateId(268)
	_State269 = _StateId(269)
	_State270 = _StateId(270)
	_State271 = _StateId(271)
	_State272 = _StateId(272)
)

type Symbol struct {
	SymbolId_ SymbolId

	Generic_ *GenericSymbol
}

func NewSymbol(token Token) (*Symbol, error) {
	symbol, ok := token.(*Symbol)
	if ok {
		return symbol, nil
	}

	symbol = &Symbol{SymbolId_: token.Id()}
	switch token.Id() {
	case _EndMarker, SpacesToken, NewlinesToken, CommentToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, ReturnToken, BreakToken, ContinueToken, PackageToken, UnsafeToken, TypeToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, DollarLbracketToken, DotdotdotToken, ExclaimToken, ExclaimExclaimToken, LabelDeclToken, JumpLabelToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, NotToken, AndToken, OrToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, LexErrorToken:
		val, ok := token.(*GenericSymbol)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting *GenericSymbol (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.Generic_ = val
	default:
		return nil, fmt.Errorf("Unexpected token type: %s", symbol.Id())
	}
	return symbol, nil
}

func (s *Symbol) Id() SymbolId {
	return s.SymbolId_
}

func (s *Symbol) Loc() Location {
	type locator interface{ Loc() Location }
	switch s.SymbolId_ {
	}
	if s.Generic_ != nil {
		return s.Generic_.Loc()
	}
	return Location{}
}

type _PseudoSymbolStack struct {
	lexer Lexer
	top   []*Symbol
}

func (stack *_PseudoSymbolStack) Top() (*Symbol, error) {
	if len(stack.top) == 0 {
		token, err := stack.lexer.Next()
		if err != nil {
			if err != io.EOF {
				return nil, fmt.Errorf("Unexpected lex error: %s", err)
			}
			token = &GenericSymbol{_EndMarker, stack.lexer.CurrentLocation()}
		}
		item, err := NewSymbol(token)
		if err != nil {
			return nil, err
		}
		stack.top = append(stack.top, item)
	}
	return stack.top[len(stack.top)-1], nil
}

func (stack *_PseudoSymbolStack) Push(symbol *Symbol) {
	stack.top = append(stack.top, symbol)
}

func (stack *_PseudoSymbolStack) Pop() (*Symbol, error) {
	if len(stack.top) == 0 {
		return nil, fmt.Errorf("internal error: cannot pop an empty top")
	}
	ret := stack.top[len(stack.top)-1]
	stack.top = stack.top[:len(stack.top)-1]
	return ret, nil
}

type _StackItem struct {
	StateId _StateId

	*Symbol
}

type _Stack []*_StackItem

type _Action struct {
	ActionType _ActionType

	ShiftStateId _StateId
	ReduceType   _ReduceType
}

func (act *_Action) ShiftItem(symbol *Symbol) *_StackItem {
	return &_StackItem{StateId: act.ShiftStateId, Symbol: symbol}
}

func (act *_Action) ReduceSymbol(
	reducer Reducer,
	stack _Stack) (
	_Stack,
	*Symbol,
	error) {

	var err error
	symbol := &Symbol{}
	switch act.ReduceType {
	case _ReduceTrueToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.TrueToLiteral(args[0].Generic_)
	case _ReduceFalseToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.FalseToLiteral(args[0].Generic_)
	case _ReduceIntegerLiteralToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.IntegerLiteralToLiteral(args[0].Generic_)
	case _ReduceFloatLiteralToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.FloatLiteralToLiteral(args[0].Generic_)
	case _ReduceRuneLiteralToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.RuneLiteralToLiteral(args[0].Generic_)
	case _ReduceStringLiteralToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.StringLiteralToLiteral(args[0].Generic_)
	case _ReduceToAnonymousFuncExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = AnonymousFuncExprType
		symbol.Generic_, err = reducer.ToAnonymousFuncExpr(args[0].Generic_, args[1].Generic_)
	case _ReduceExplicitToAnonymousStructExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = AnonymousStructExprType
		symbol.Generic_, err = reducer.ExplicitToAnonymousStructExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceImplicitToAnonymousStructExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = AnonymousStructExprType
		symbol.Generic_, err = reducer.ImplicitToAnonymousStructExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceLiteralToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.LiteralToAtomExpr(args[0].Generic_)
	case _ReduceIdentifierToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.IdentifierToAtomExpr(args[0].Generic_)
	case _ReduceBlockExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.BlockExprToAtomExpr(args[0].Generic_)
	case _ReduceAnonymousFuncExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.AnonymousFuncExprToAtomExpr(args[0].Generic_)
	case _ReduceAnonymousStructExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.AnonymousStructExprToAtomExpr(args[0].Generic_)
	case _ReduceLexErrorToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.LexErrorToAtomExpr(args[0].Generic_)
	case _ReduceValueTypeToGenericArguments:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericArgumentsType
		symbol.Generic_, err = reducer.ValueTypeToGenericArguments(args[0].Generic_)
	case _ReduceAddToGenericArguments:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = GenericArgumentsType
		symbol.Generic_, err = reducer.AddToGenericArguments(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceGenericArgumentsToOptionalGenericArguments:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalGenericArgumentsType
		symbol.Generic_, err = reducer.GenericArgumentsToOptionalGenericArguments(args[0].Generic_)
	case _ReduceNilToOptionalGenericArguments:
		symbol.SymbolId_ = OptionalGenericArgumentsType
		symbol.Generic_, err = reducer.NilToOptionalGenericArguments()
	case _ReduceBindingToOptionalGenericBinding:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OptionalGenericBindingType
		symbol.Generic_, err = reducer.BindingToOptionalGenericBinding(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceNilToOptionalGenericBinding:
		symbol.SymbolId_ = OptionalGenericBindingType
		symbol.Generic_, err = reducer.NilToOptionalGenericBinding()
	case _ReducePositionalToArgument:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentType
		symbol.Generic_, err = reducer.PositionalToArgument(args[0].Generic_)
	case _ReduceArgumentToArguments:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentsType
		symbol.Generic_, err = reducer.ArgumentToArguments(args[0].Generic_)
	case _ReduceAddToArguments:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ArgumentsType
		symbol.Generic_, err = reducer.AddToArguments(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceArgumentsToOptionalArguments:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalArgumentsType
		symbol.Generic_, err = reducer.ArgumentsToOptionalArguments(args[0].Generic_)
	case _ReduceNilToOptionalArguments:
		symbol.SymbolId_ = OptionalArgumentsType
		symbol.Generic_, err = reducer.NilToOptionalArguments()
	case _ReduceToCallExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CallExprType
		symbol.Generic_, err = reducer.ToCallExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceAtomExprToAccessExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessExprType
		symbol.Generic_, err = reducer.AtomExprToAccessExpr(args[0].Generic_)
	case _ReduceAccessToAccessExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = AccessExprType
		symbol.Generic_, err = reducer.AccessToAccessExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceCallExprToAccessExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessExprType
		symbol.Generic_, err = reducer.CallExprToAccessExpr(args[0].Generic_)
	case _ReduceIndexToAccessExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = AccessExprType
		symbol.Generic_, err = reducer.IndexToAccessExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceAccessExprToPostfixUnaryExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixUnaryExprType
		symbol.Generic_, err = reducer.AccessExprToPostfixUnaryExpr(args[0].Generic_)
	case _ReduceQuestionToPostfixUnaryExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PostfixUnaryExprType
		symbol.Generic_, err = reducer.QuestionToPostfixUnaryExpr(args[0].Generic_, args[1].Generic_)
	case _ReduceNotToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		symbol.Generic_, err = reducer.NotToPrefixUnaryOp(args[0].Generic_)
	case _ReduceBitNegToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		symbol.Generic_, err = reducer.BitNegToPrefixUnaryOp(args[0].Generic_)
	case _ReduceSubToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		symbol.Generic_, err = reducer.SubToPrefixUnaryOp(args[0].Generic_)
	case _ReduceMulToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		symbol.Generic_, err = reducer.MulToPrefixUnaryOp(args[0].Generic_)
	case _ReduceBitAndToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		symbol.Generic_, err = reducer.BitAndToPrefixUnaryOp(args[0].Generic_)
	case _ReducePostfixUnaryExprToPrefixUnaryExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryExprType
		symbol.Generic_, err = reducer.PostfixUnaryExprToPrefixUnaryExpr(args[0].Generic_)
	case _ReducePrefixOpToPrefixUnaryExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PrefixUnaryExprType
		symbol.Generic_, err = reducer.PrefixOpToPrefixUnaryExpr(args[0].Generic_, args[1].Generic_)
	case _ReduceMulToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		symbol.Generic_, err = reducer.MulToMulOp(args[0].Generic_)
	case _ReduceDivToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		symbol.Generic_, err = reducer.DivToMulOp(args[0].Generic_)
	case _ReduceModToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		symbol.Generic_, err = reducer.ModToMulOp(args[0].Generic_)
	case _ReduceBitAndToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		symbol.Generic_, err = reducer.BitAndToMulOp(args[0].Generic_)
	case _ReduceBitLshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		symbol.Generic_, err = reducer.BitLshiftToMulOp(args[0].Generic_)
	case _ReduceBitRshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		symbol.Generic_, err = reducer.BitRshiftToMulOp(args[0].Generic_)
	case _ReducePrefixUnaryExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		symbol.Generic_, err = reducer.PrefixUnaryExprToMulExpr(args[0].Generic_)
	case _ReduceOpToMulExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = MulExprType
		symbol.Generic_, err = reducer.OpToMulExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceAddToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		symbol.Generic_, err = reducer.AddToAddOp(args[0].Generic_)
	case _ReduceSubToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		symbol.Generic_, err = reducer.SubToAddOp(args[0].Generic_)
	case _ReduceBitOrToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		symbol.Generic_, err = reducer.BitOrToAddOp(args[0].Generic_)
	case _ReduceBitXorToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		symbol.Generic_, err = reducer.BitXorToAddOp(args[0].Generic_)
	case _ReduceMulExprToAddExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddExprType
		symbol.Generic_, err = reducer.MulExprToAddExpr(args[0].Generic_)
	case _ReduceOpToAddExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = AddExprType
		symbol.Generic_, err = reducer.OpToAddExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		symbol.Generic_, err = reducer.EqualToCmpOp(args[0].Generic_)
	case _ReduceNotEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		symbol.Generic_, err = reducer.NotEqualToCmpOp(args[0].Generic_)
	case _ReduceLessToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		symbol.Generic_, err = reducer.LessToCmpOp(args[0].Generic_)
	case _ReduceLessOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		symbol.Generic_, err = reducer.LessOrEqualToCmpOp(args[0].Generic_)
	case _ReduceGreaterToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		symbol.Generic_, err = reducer.GreaterToCmpOp(args[0].Generic_)
	case _ReduceGreaterOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		symbol.Generic_, err = reducer.GreaterOrEqualToCmpOp(args[0].Generic_)
	case _ReduceAddExprToCmpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpExprType
		symbol.Generic_, err = reducer.AddExprToCmpExpr(args[0].Generic_)
	case _ReduceOpToCmpExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CmpExprType
		symbol.Generic_, err = reducer.OpToCmpExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceCmpExprToAndExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AndExprType
		symbol.Generic_, err = reducer.CmpExprToAndExpr(args[0].Generic_)
	case _ReduceOpToAndExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = AndExprType
		symbol.Generic_, err = reducer.OpToAndExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceAndExprToOrExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OrExprType
		symbol.Generic_, err = reducer.AndExprToOrExpr(args[0].Generic_)
	case _ReduceOpToOrExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OrExprType
		symbol.Generic_, err = reducer.OpToOrExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToSequenceExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SequenceExprType
		symbol.Generic_, err = reducer.ToSequenceExpr(args[0].Generic_)
	case _ReduceJumpLabelToOptionalJumpLabel:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalJumpLabelType
		symbol.Generic_, err = reducer.JumpLabelToOptionalJumpLabel(args[0].Generic_)
	case _ReduceUnlabelledToOptionalJumpLabel:
		symbol.SymbolId_ = OptionalJumpLabelType
		symbol.Generic_, err = reducer.UnlabelledToOptionalJumpLabel()
	case _ReduceExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalExpressionOrImplicitStructType
		symbol.Generic_, err = reducer.ExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct(args[0].Generic_)
	case _ReduceNilToOptionalExpressionOrImplicitStruct:
		symbol.SymbolId_ = OptionalExpressionOrImplicitStructType
		symbol.Generic_, err = reducer.NilToOptionalExpressionOrImplicitStruct()
	case _ReduceReturnToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		symbol.Generic_, err = reducer.ReturnToJumpType(args[0].Generic_)
	case _ReduceBreakToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		symbol.Generic_, err = reducer.BreakToJumpType(args[0].Generic_)
	case _ReduceContinueToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		symbol.Generic_, err = reducer.ContinueToJumpType(args[0].Generic_)
	case _ReduceAddOneAssignToOpOneAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OpOneAssignType
		symbol.Generic_, err = reducer.AddOneAssignToOpOneAssign(args[0].Generic_)
	case _ReduceSubOneAssignToOpOneAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OpOneAssignType
		symbol.Generic_, err = reducer.SubOneAssignToOpOneAssign(args[0].Generic_)
	case _ReduceAddAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.AddAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceSubAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.SubAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceMulAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.MulAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceDivAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.DivAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceModAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.ModAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitNegAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitNegAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitAndAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitAndAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitOrAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitOrAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitXorAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitXorAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitLshiftAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitLshiftAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceBitRshiftAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		symbol.Generic_, err = reducer.BitRshiftAssignToBinaryOpAssign(args[0].Generic_)
	case _ReduceExpressionToExpressionOrImplicitStruct:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionOrImplicitStructType
		symbol.Generic_, err = reducer.ExpressionToExpressionOrImplicitStruct(args[0].Generic_)
	case _ReduceImplicitStructToExpressionOrImplicitStruct:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExpressionOrImplicitStructType
		symbol.Generic_, err = reducer.ImplicitStructToExpressionOrImplicitStruct(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToUnsafeStatement:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = UnsafeStatementType
		symbol.Generic_, err = reducer.ToUnsafeStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceUnsafeStatementToStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.UnsafeStatementToStatementBody(args[0].Generic_)
	case _ReduceExpressionOrImplicitStructToStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.ExpressionOrImplicitStructToStatementBody(args[0].Generic_)
	case _ReduceAsyncToStatementBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.AsyncToStatementBody(args[0].Generic_, args[1].Generic_)
	case _ReduceJumpToStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.JumpToStatementBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceOpOneAssignToStatementBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.OpOneAssignToStatementBody(args[0].Generic_, args[1].Generic_)
	case _ReduceBinaryOpAssignToStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.BinaryOpAssignToStatementBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceImplicitToStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementType
		symbol.Generic_, err = reducer.ImplicitToStatement(args[0].Generic_, args[1].Generic_)
	case _ReduceExplicitToStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementType
		symbol.Generic_, err = reducer.ExplicitToStatement(args[0].Generic_, args[1].Generic_)
	case _ReduceEmptyListToStatements:
		symbol.SymbolId_ = StatementsType
		symbol.Generic_, err = reducer.EmptyListToStatements()
	case _ReduceAddToStatements:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementsType
		symbol.Generic_, err = reducer.AddToStatements(args[0].Generic_, args[1].Generic_)
	case _ReduceLabelDeclToOptionalLabelDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalLabelDeclType
		symbol.Generic_, err = reducer.LabelDeclToOptionalLabelDecl(args[0].Generic_)
	case _ReduceUnlabelledToOptionalLabelDecl:
		symbol.SymbolId_ = OptionalLabelDeclType
		symbol.Generic_, err = reducer.UnlabelledToOptionalLabelDecl()
	case _ReduceToBlockBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BlockBodyType
		symbol.Generic_, err = reducer.ToBlockBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToBlockExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = BlockExprType
		symbol.Generic_, err = reducer.ToBlockExpr(args[0].Generic_, args[1].Generic_)
	case _ReduceNoElseToIfExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = IfExprType
		symbol.Generic_, err = reducer.NoElseToIfExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceIfElseToIfExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = IfExprType
		symbol.Generic_, err = reducer.IfElseToIfExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceMultiIfElseToIfExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = IfExprType
		symbol.Generic_, err = reducer.MultiIfElseToIfExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceToSwitchExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = SwitchExprType
		symbol.Generic_, err = reducer.ToSwitchExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceInfiniteToLoopExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LoopExprType
		symbol.Generic_, err = reducer.InfiniteToLoopExpr(args[0].Generic_, args[1].Generic_)
	case _ReduceWhileToLoopExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = LoopExprType
		symbol.Generic_, err = reducer.WhileToLoopExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceSequenceExprToExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.SequenceExprToExpression(args[0].Generic_)
	case _ReduceIfExprToExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.IfExprToExpression(args[0].Generic_, args[1].Generic_)
	case _ReduceSwitchExprToExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.SwitchExprToExpression(args[0].Generic_, args[1].Generic_)
	case _ReduceLoopExprToExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.LoopExprToExpression(args[0].Generic_, args[1].Generic_)
	case _ReduceUnconstrainedToGenericParameter:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericParameterType
		symbol.Generic_, err = reducer.UnconstrainedToGenericParameter(args[0].Generic_)
	case _ReduceConstrainedToGenericParameter:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = GenericParameterType
		symbol.Generic_, err = reducer.ConstrainedToGenericParameter(args[0].Generic_, args[1].Generic_)
	case _ReduceGenericParameterToGenericParameters:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericParametersType
		symbol.Generic_, err = reducer.GenericParameterToGenericParameters(args[0].Generic_)
	case _ReduceAddToGenericParameters:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = GenericParametersType
		symbol.Generic_, err = reducer.AddToGenericParameters(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceGenericParametersToOptionalGenericParameters:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalGenericParametersType
		symbol.Generic_, err = reducer.GenericParametersToOptionalGenericParameters(args[0].Generic_)
	case _ReduceNilToOptionalGenericParameters:
		symbol.SymbolId_ = OptionalGenericParametersType
		symbol.Generic_, err = reducer.NilToOptionalGenericParameters()
	case _ReduceGenericToOptionalGenericParameterDecl:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OptionalGenericParameterDeclType
		symbol.Generic_, err = reducer.GenericToOptionalGenericParameterDecl(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceNilToOptionalGenericParameterDecl:
		symbol.SymbolId_ = OptionalGenericParameterDeclType
		symbol.Generic_, err = reducer.NilToOptionalGenericParameterDecl()
	case _ReduceNamedToAtomType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.NamedToAtomType(args[0].Generic_, args[1].Generic_)
	case _ReduceExplicitStructTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ExplicitStructTypeToAtomType(args[0].Generic_)
	case _ReduceImplicitStructTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ImplicitStructTypeToAtomType(args[0].Generic_)
	case _ReduceExplicitEnumTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ExplicitEnumTypeToAtomType(args[0].Generic_)
	case _ReduceTraitTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.TraitTypeToAtomType(args[0].Generic_)
	case _ReduceAtomTypeToTraitableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitableTypeType
		symbol.Generic_, err = reducer.AtomTypeToTraitableType(args[0].Generic_)
	case _ReduceMethodInterfaceToTraitableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = TraitableTypeType
		symbol.Generic_, err = reducer.MethodInterfaceToTraitableType(args[0].Generic_, args[1].Generic_)
	case _ReduceTraitToTraitableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = TraitableTypeType
		symbol.Generic_, err = reducer.TraitToTraitableType(args[0].Generic_, args[1].Generic_)
	case _ReduceTraitableTypeToTraitMulType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitMulTypeType
		symbol.Generic_, err = reducer.TraitableTypeToTraitMulType(args[0].Generic_)
	case _ReduceIntersectToTraitMulType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TraitMulTypeType
		symbol.Generic_, err = reducer.IntersectToTraitMulType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceTraitMulTypeToTraitAddType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitAddTypeType
		symbol.Generic_, err = reducer.TraitMulTypeToTraitAddType(args[0].Generic_)
	case _ReduceUnionToTraitAddType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TraitAddTypeType
		symbol.Generic_, err = reducer.UnionToTraitAddType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceDifferenceToTraitAddType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TraitAddTypeType
		symbol.Generic_, err = reducer.DifferenceToTraitAddType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceExplicitToFieldDecl:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = FieldDeclType
		symbol.Generic_, err = reducer.ExplicitToFieldDecl(args[0].Generic_, args[1].Generic_)
	case _ReduceImplicitToFieldDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldDeclType
		symbol.Generic_, err = reducer.ImplicitToFieldDecl(args[0].Generic_)
	case _ReduceFieldDeclToFieldDecls:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldDeclsType
		symbol.Generic_, err = reducer.FieldDeclToFieldDecls(args[0].Generic_)
	case _ReduceImplicitToFieldDecls:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = FieldDeclsType
		symbol.Generic_, err = reducer.ImplicitToFieldDecls(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceExplicitToFieldDecls:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = FieldDeclsType
		symbol.Generic_, err = reducer.ExplicitToFieldDecls(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceFieldDeclsToOptionalFieldDecls:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalFieldDeclsType
		symbol.Generic_, err = reducer.FieldDeclsToOptionalFieldDecls(args[0].Generic_)
	case _ReduceNilToOptionalFieldDecls:
		symbol.SymbolId_ = OptionalFieldDeclsType
		symbol.Generic_, err = reducer.NilToOptionalFieldDecls()
	case _ReduceToImplicitStructType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitStructTypeType
		symbol.Generic_, err = reducer.ToImplicitStructType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToExplicitStructType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExplicitStructTypeType
		symbol.Generic_, err = reducer.ToExplicitStructType(args[0].Generic_, args[1].Generic_)
	case _ReduceToTraitType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TraitTypeType
		symbol.Generic_, err = reducer.ToTraitType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToExplicitEnumType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumTypeType
		symbol.Generic_, err = reducer.ToExplicitEnumType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceTraitAddTypeToFuncType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FuncTypeType
		symbol.Generic_, err = reducer.TraitAddTypeToFuncType(args[0].Generic_)
	case _ReduceReferenceToFuncType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = FuncTypeType
		symbol.Generic_, err = reducer.ReferenceToFuncType(args[0].Generic_, args[1].Generic_)
	case _ReduceExplicitFuncTypeToFuncType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FuncTypeType
		symbol.Generic_, err = reducer.ExplicitFuncTypeToFuncType(args[0].Generic_)
	case _ReduceFuncTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.FuncTypeToValueType(args[0].Generic_)
	case _ReduceImplicitEnumToValueType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.ImplicitEnumToValueType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceDefinitionToTypeDecl:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = TypeDeclType
		symbol.Generic_, err = reducer.DefinitionToTypeDecl(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceAliasToTypeDecl:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = TypeDeclType
		symbol.Generic_, err = reducer.AliasToTypeDecl(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceFieldDeclToParameterDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.FieldDeclToParameterDecl(args[0].Generic_)
	case _ReduceInferredToParameterDecl:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.InferredToParameterDecl(args[0].Generic_, args[1].Generic_)
	case _ReduceParameterDeclToNonEmptyParameters:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = NonEmptyParametersType
		symbol.Generic_, err = reducer.ParameterDeclToNonEmptyParameters(args[0].Generic_)
	case _ReduceAddToNonEmptyParameters:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = NonEmptyParametersType
		symbol.Generic_, err = reducer.AddToNonEmptyParameters(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceExplicitToVararg:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = VarargType
		symbol.Generic_, err = reducer.ExplicitToVararg(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceImplicitToVararg:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = VarargType
		symbol.Generic_, err = reducer.ImplicitToVararg(args[0].Generic_, args[1].Generic_)
	case _ReduceInferredToVararg:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = VarargType
		symbol.Generic_, err = reducer.InferredToVararg(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceVarargToOptionalVararg:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalVarargType
		symbol.Generic_, err = reducer.VarargToOptionalVararg(args[0].Generic_)
	case _ReduceVararg2ToOptionalVararg:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = OptionalVarargType
		symbol.Generic_, err = reducer.Vararg2ToOptionalVararg(args[0].Generic_, args[1].Generic_)
	case _ReduceNilToOptionalVararg:
		symbol.SymbolId_ = OptionalVarargType
		symbol.Generic_, err = reducer.NilToOptionalVararg()
	case _ReduceNonEmptyParametersToParameters:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParametersType
		symbol.Generic_, err = reducer.NonEmptyParametersToParameters(args[0].Generic_)
	case _ReduceMixedToParameters:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParametersType
		symbol.Generic_, err = reducer.MixedToParameters(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceOptionalVarargToParameters:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParametersType
		symbol.Generic_, err = reducer.OptionalVarargToParameters(args[0].Generic_)
	case _ReduceTypedToImplicitFuncType:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ImplicitFuncTypeType
		symbol.Generic_, err = reducer.TypedToImplicitFuncType(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceInferredToImplicitFuncType:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ImplicitFuncTypeType
		symbol.Generic_, err = reducer.InferredToImplicitFuncType(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceImplicitUnitToImplicitFuncType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitFuncTypeType
		symbol.Generic_, err = reducer.ImplicitUnitToImplicitFuncType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToExplicitFuncType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExplicitFuncTypeType
		symbol.Generic_, err = reducer.ToExplicitFuncType(args[0].Generic_, args[1].Generic_)
	case _ReduceReceiverToOptionalReceiver:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OptionalReceiverType
		symbol.Generic_, err = reducer.ReceiverToOptionalReceiver(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceNilToOptionalReceiver:
		symbol.SymbolId_ = OptionalReceiverType
		symbol.Generic_, err = reducer.NilToOptionalReceiver()
	case _ReduceToImplicitFuncDecl:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitFuncDeclType
		symbol.Generic_, err = reducer.ToImplicitFuncDecl(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToExplicitFuncDecl:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitFuncDeclType
		symbol.Generic_, err = reducer.ToExplicitFuncDecl(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToFuncDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = FuncDefType
		symbol.Generic_, err = reducer.ToFuncDef(args[0].Generic_, args[1].Generic_)
	case _ReduceNoSpecToPackageDecl:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PackageDeclType
		symbol.Generic_, err = reducer.NoSpecToPackageDecl(args[0].Generic_, args[1].Generic_)
	case _ReduceWithSpecToPackageDecl:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = PackageDeclType
		symbol.Generic_, err = reducer.WithSpecToPackageDecl(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceToPackageStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PackageStatementBodyType
		symbol.Generic_, err = reducer.ToPackageStatementBody(args[0].Generic_)
	case _ReduceImplicitToPackageStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PackageStatementType
		symbol.Generic_, err = reducer.ImplicitToPackageStatement(args[0].Generic_, args[1].Generic_)
	case _ReduceExplicitToPackageStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PackageStatementType
		symbol.Generic_, err = reducer.ExplicitToPackageStatement(args[0].Generic_, args[1].Generic_)
	case _ReduceEmptyListToPackageStatements:
		symbol.SymbolId_ = PackageStatementsType
		symbol.Generic_, err = reducer.EmptyListToPackageStatements()
	case _ReduceAddToPackageStatements:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PackageStatementsType
		symbol.Generic_, err = reducer.AddToPackageStatements(args[0].Generic_, args[1].Generic_)
	case _ReduceSpacesToLexInternalTokens:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LexInternalTokensType
		symbol.Generic_, err = reducer.SpacesToLexInternalTokens(args[0].Generic_)
	case _ReduceCommentToLexInternalTokens:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LexInternalTokensType
		symbol.Generic_, err = reducer.CommentToLexInternalTokens(args[0].Generic_)
	default:
		panic("Unknown reduce type: " + act.ReduceType.String())
	}

	if err != nil {
		err = fmt.Errorf("Unexpected %s reduce error: %s", act.ReduceType, err)
	}

	return stack, symbol, err
}

type _ActionTableKey struct {
	_StateId
	SymbolId
}

type _ActionTableType map[_ActionTableKey]*_Action

func (table _ActionTableType) Get(
	stateId _StateId,
	symbolId SymbolId) (
	*_Action,
	bool) {

	action, ok := table[_ActionTableKey{stateId, symbolId}]
	if ok {
		return action, ok
	}

	action, ok = table[_ActionTableKey{stateId, _WildcardMarker}]
	return action, ok
}

var (
	_GotoState1Action                                                           = &_Action{_ShiftAction, _State1, 0}
	_GotoState2Action                                                           = &_Action{_ShiftAction, _State2, 0}
	_GotoState3Action                                                           = &_Action{_ShiftAction, _State3, 0}
	_GotoState4Action                                                           = &_Action{_ShiftAction, _State4, 0}
	_GotoState5Action                                                           = &_Action{_ShiftAction, _State5, 0}
	_GotoState6Action                                                           = &_Action{_ShiftAction, _State6, 0}
	_GotoState7Action                                                           = &_Action{_ShiftAction, _State7, 0}
	_GotoState8Action                                                           = &_Action{_ShiftAction, _State8, 0}
	_GotoState9Action                                                           = &_Action{_ShiftAction, _State9, 0}
	_GotoState10Action                                                          = &_Action{_ShiftAction, _State10, 0}
	_GotoState11Action                                                          = &_Action{_ShiftAction, _State11, 0}
	_GotoState12Action                                                          = &_Action{_ShiftAction, _State12, 0}
	_GotoState13Action                                                          = &_Action{_ShiftAction, _State13, 0}
	_GotoState14Action                                                          = &_Action{_ShiftAction, _State14, 0}
	_GotoState15Action                                                          = &_Action{_ShiftAction, _State15, 0}
	_GotoState16Action                                                          = &_Action{_ShiftAction, _State16, 0}
	_GotoState17Action                                                          = &_Action{_ShiftAction, _State17, 0}
	_GotoState18Action                                                          = &_Action{_ShiftAction, _State18, 0}
	_GotoState19Action                                                          = &_Action{_ShiftAction, _State19, 0}
	_GotoState20Action                                                          = &_Action{_ShiftAction, _State20, 0}
	_GotoState21Action                                                          = &_Action{_ShiftAction, _State21, 0}
	_GotoState22Action                                                          = &_Action{_ShiftAction, _State22, 0}
	_GotoState23Action                                                          = &_Action{_ShiftAction, _State23, 0}
	_GotoState24Action                                                          = &_Action{_ShiftAction, _State24, 0}
	_GotoState25Action                                                          = &_Action{_ShiftAction, _State25, 0}
	_GotoState26Action                                                          = &_Action{_ShiftAction, _State26, 0}
	_GotoState27Action                                                          = &_Action{_ShiftAction, _State27, 0}
	_GotoState28Action                                                          = &_Action{_ShiftAction, _State28, 0}
	_GotoState29Action                                                          = &_Action{_ShiftAction, _State29, 0}
	_GotoState30Action                                                          = &_Action{_ShiftAction, _State30, 0}
	_GotoState31Action                                                          = &_Action{_ShiftAction, _State31, 0}
	_GotoState32Action                                                          = &_Action{_ShiftAction, _State32, 0}
	_GotoState33Action                                                          = &_Action{_ShiftAction, _State33, 0}
	_GotoState34Action                                                          = &_Action{_ShiftAction, _State34, 0}
	_GotoState35Action                                                          = &_Action{_ShiftAction, _State35, 0}
	_GotoState36Action                                                          = &_Action{_ShiftAction, _State36, 0}
	_GotoState37Action                                                          = &_Action{_ShiftAction, _State37, 0}
	_GotoState38Action                                                          = &_Action{_ShiftAction, _State38, 0}
	_GotoState39Action                                                          = &_Action{_ShiftAction, _State39, 0}
	_GotoState40Action                                                          = &_Action{_ShiftAction, _State40, 0}
	_GotoState41Action                                                          = &_Action{_ShiftAction, _State41, 0}
	_GotoState42Action                                                          = &_Action{_ShiftAction, _State42, 0}
	_GotoState43Action                                                          = &_Action{_ShiftAction, _State43, 0}
	_GotoState44Action                                                          = &_Action{_ShiftAction, _State44, 0}
	_GotoState45Action                                                          = &_Action{_ShiftAction, _State45, 0}
	_GotoState46Action                                                          = &_Action{_ShiftAction, _State46, 0}
	_GotoState47Action                                                          = &_Action{_ShiftAction, _State47, 0}
	_GotoState48Action                                                          = &_Action{_ShiftAction, _State48, 0}
	_GotoState49Action                                                          = &_Action{_ShiftAction, _State49, 0}
	_GotoState50Action                                                          = &_Action{_ShiftAction, _State50, 0}
	_GotoState51Action                                                          = &_Action{_ShiftAction, _State51, 0}
	_GotoState52Action                                                          = &_Action{_ShiftAction, _State52, 0}
	_GotoState53Action                                                          = &_Action{_ShiftAction, _State53, 0}
	_GotoState54Action                                                          = &_Action{_ShiftAction, _State54, 0}
	_GotoState55Action                                                          = &_Action{_ShiftAction, _State55, 0}
	_GotoState56Action                                                          = &_Action{_ShiftAction, _State56, 0}
	_GotoState57Action                                                          = &_Action{_ShiftAction, _State57, 0}
	_GotoState58Action                                                          = &_Action{_ShiftAction, _State58, 0}
	_GotoState59Action                                                          = &_Action{_ShiftAction, _State59, 0}
	_GotoState60Action                                                          = &_Action{_ShiftAction, _State60, 0}
	_GotoState61Action                                                          = &_Action{_ShiftAction, _State61, 0}
	_GotoState62Action                                                          = &_Action{_ShiftAction, _State62, 0}
	_GotoState63Action                                                          = &_Action{_ShiftAction, _State63, 0}
	_GotoState64Action                                                          = &_Action{_ShiftAction, _State64, 0}
	_GotoState65Action                                                          = &_Action{_ShiftAction, _State65, 0}
	_GotoState66Action                                                          = &_Action{_ShiftAction, _State66, 0}
	_GotoState67Action                                                          = &_Action{_ShiftAction, _State67, 0}
	_GotoState68Action                                                          = &_Action{_ShiftAction, _State68, 0}
	_GotoState69Action                                                          = &_Action{_ShiftAction, _State69, 0}
	_GotoState70Action                                                          = &_Action{_ShiftAction, _State70, 0}
	_GotoState71Action                                                          = &_Action{_ShiftAction, _State71, 0}
	_GotoState72Action                                                          = &_Action{_ShiftAction, _State72, 0}
	_GotoState73Action                                                          = &_Action{_ShiftAction, _State73, 0}
	_GotoState74Action                                                          = &_Action{_ShiftAction, _State74, 0}
	_GotoState75Action                                                          = &_Action{_ShiftAction, _State75, 0}
	_GotoState76Action                                                          = &_Action{_ShiftAction, _State76, 0}
	_GotoState77Action                                                          = &_Action{_ShiftAction, _State77, 0}
	_GotoState78Action                                                          = &_Action{_ShiftAction, _State78, 0}
	_GotoState79Action                                                          = &_Action{_ShiftAction, _State79, 0}
	_GotoState80Action                                                          = &_Action{_ShiftAction, _State80, 0}
	_GotoState81Action                                                          = &_Action{_ShiftAction, _State81, 0}
	_GotoState82Action                                                          = &_Action{_ShiftAction, _State82, 0}
	_GotoState83Action                                                          = &_Action{_ShiftAction, _State83, 0}
	_GotoState84Action                                                          = &_Action{_ShiftAction, _State84, 0}
	_GotoState85Action                                                          = &_Action{_ShiftAction, _State85, 0}
	_GotoState86Action                                                          = &_Action{_ShiftAction, _State86, 0}
	_GotoState87Action                                                          = &_Action{_ShiftAction, _State87, 0}
	_GotoState88Action                                                          = &_Action{_ShiftAction, _State88, 0}
	_GotoState89Action                                                          = &_Action{_ShiftAction, _State89, 0}
	_GotoState90Action                                                          = &_Action{_ShiftAction, _State90, 0}
	_GotoState91Action                                                          = &_Action{_ShiftAction, _State91, 0}
	_GotoState92Action                                                          = &_Action{_ShiftAction, _State92, 0}
	_GotoState93Action                                                          = &_Action{_ShiftAction, _State93, 0}
	_GotoState94Action                                                          = &_Action{_ShiftAction, _State94, 0}
	_GotoState95Action                                                          = &_Action{_ShiftAction, _State95, 0}
	_GotoState96Action                                                          = &_Action{_ShiftAction, _State96, 0}
	_GotoState97Action                                                          = &_Action{_ShiftAction, _State97, 0}
	_GotoState98Action                                                          = &_Action{_ShiftAction, _State98, 0}
	_GotoState99Action                                                          = &_Action{_ShiftAction, _State99, 0}
	_GotoState100Action                                                         = &_Action{_ShiftAction, _State100, 0}
	_GotoState101Action                                                         = &_Action{_ShiftAction, _State101, 0}
	_GotoState102Action                                                         = &_Action{_ShiftAction, _State102, 0}
	_GotoState103Action                                                         = &_Action{_ShiftAction, _State103, 0}
	_GotoState104Action                                                         = &_Action{_ShiftAction, _State104, 0}
	_GotoState105Action                                                         = &_Action{_ShiftAction, _State105, 0}
	_GotoState106Action                                                         = &_Action{_ShiftAction, _State106, 0}
	_GotoState107Action                                                         = &_Action{_ShiftAction, _State107, 0}
	_GotoState108Action                                                         = &_Action{_ShiftAction, _State108, 0}
	_GotoState109Action                                                         = &_Action{_ShiftAction, _State109, 0}
	_GotoState110Action                                                         = &_Action{_ShiftAction, _State110, 0}
	_GotoState111Action                                                         = &_Action{_ShiftAction, _State111, 0}
	_GotoState112Action                                                         = &_Action{_ShiftAction, _State112, 0}
	_GotoState113Action                                                         = &_Action{_ShiftAction, _State113, 0}
	_GotoState114Action                                                         = &_Action{_ShiftAction, _State114, 0}
	_GotoState115Action                                                         = &_Action{_ShiftAction, _State115, 0}
	_GotoState116Action                                                         = &_Action{_ShiftAction, _State116, 0}
	_GotoState117Action                                                         = &_Action{_ShiftAction, _State117, 0}
	_GotoState118Action                                                         = &_Action{_ShiftAction, _State118, 0}
	_GotoState119Action                                                         = &_Action{_ShiftAction, _State119, 0}
	_GotoState120Action                                                         = &_Action{_ShiftAction, _State120, 0}
	_GotoState121Action                                                         = &_Action{_ShiftAction, _State121, 0}
	_GotoState122Action                                                         = &_Action{_ShiftAction, _State122, 0}
	_GotoState123Action                                                         = &_Action{_ShiftAction, _State123, 0}
	_GotoState124Action                                                         = &_Action{_ShiftAction, _State124, 0}
	_GotoState125Action                                                         = &_Action{_ShiftAction, _State125, 0}
	_GotoState126Action                                                         = &_Action{_ShiftAction, _State126, 0}
	_GotoState127Action                                                         = &_Action{_ShiftAction, _State127, 0}
	_GotoState128Action                                                         = &_Action{_ShiftAction, _State128, 0}
	_GotoState129Action                                                         = &_Action{_ShiftAction, _State129, 0}
	_GotoState130Action                                                         = &_Action{_ShiftAction, _State130, 0}
	_GotoState131Action                                                         = &_Action{_ShiftAction, _State131, 0}
	_GotoState132Action                                                         = &_Action{_ShiftAction, _State132, 0}
	_GotoState133Action                                                         = &_Action{_ShiftAction, _State133, 0}
	_GotoState134Action                                                         = &_Action{_ShiftAction, _State134, 0}
	_GotoState135Action                                                         = &_Action{_ShiftAction, _State135, 0}
	_GotoState136Action                                                         = &_Action{_ShiftAction, _State136, 0}
	_GotoState137Action                                                         = &_Action{_ShiftAction, _State137, 0}
	_GotoState138Action                                                         = &_Action{_ShiftAction, _State138, 0}
	_GotoState139Action                                                         = &_Action{_ShiftAction, _State139, 0}
	_GotoState140Action                                                         = &_Action{_ShiftAction, _State140, 0}
	_GotoState141Action                                                         = &_Action{_ShiftAction, _State141, 0}
	_GotoState142Action                                                         = &_Action{_ShiftAction, _State142, 0}
	_GotoState143Action                                                         = &_Action{_ShiftAction, _State143, 0}
	_GotoState144Action                                                         = &_Action{_ShiftAction, _State144, 0}
	_GotoState145Action                                                         = &_Action{_ShiftAction, _State145, 0}
	_GotoState146Action                                                         = &_Action{_ShiftAction, _State146, 0}
	_GotoState147Action                                                         = &_Action{_ShiftAction, _State147, 0}
	_GotoState148Action                                                         = &_Action{_ShiftAction, _State148, 0}
	_GotoState149Action                                                         = &_Action{_ShiftAction, _State149, 0}
	_GotoState150Action                                                         = &_Action{_ShiftAction, _State150, 0}
	_GotoState151Action                                                         = &_Action{_ShiftAction, _State151, 0}
	_GotoState152Action                                                         = &_Action{_ShiftAction, _State152, 0}
	_GotoState153Action                                                         = &_Action{_ShiftAction, _State153, 0}
	_GotoState154Action                                                         = &_Action{_ShiftAction, _State154, 0}
	_GotoState155Action                                                         = &_Action{_ShiftAction, _State155, 0}
	_GotoState156Action                                                         = &_Action{_ShiftAction, _State156, 0}
	_GotoState157Action                                                         = &_Action{_ShiftAction, _State157, 0}
	_GotoState158Action                                                         = &_Action{_ShiftAction, _State158, 0}
	_GotoState159Action                                                         = &_Action{_ShiftAction, _State159, 0}
	_GotoState160Action                                                         = &_Action{_ShiftAction, _State160, 0}
	_GotoState161Action                                                         = &_Action{_ShiftAction, _State161, 0}
	_GotoState162Action                                                         = &_Action{_ShiftAction, _State162, 0}
	_GotoState163Action                                                         = &_Action{_ShiftAction, _State163, 0}
	_GotoState164Action                                                         = &_Action{_ShiftAction, _State164, 0}
	_GotoState165Action                                                         = &_Action{_ShiftAction, _State165, 0}
	_GotoState166Action                                                         = &_Action{_ShiftAction, _State166, 0}
	_GotoState167Action                                                         = &_Action{_ShiftAction, _State167, 0}
	_GotoState168Action                                                         = &_Action{_ShiftAction, _State168, 0}
	_GotoState169Action                                                         = &_Action{_ShiftAction, _State169, 0}
	_GotoState170Action                                                         = &_Action{_ShiftAction, _State170, 0}
	_GotoState171Action                                                         = &_Action{_ShiftAction, _State171, 0}
	_GotoState172Action                                                         = &_Action{_ShiftAction, _State172, 0}
	_GotoState173Action                                                         = &_Action{_ShiftAction, _State173, 0}
	_GotoState174Action                                                         = &_Action{_ShiftAction, _State174, 0}
	_GotoState175Action                                                         = &_Action{_ShiftAction, _State175, 0}
	_GotoState176Action                                                         = &_Action{_ShiftAction, _State176, 0}
	_GotoState177Action                                                         = &_Action{_ShiftAction, _State177, 0}
	_GotoState178Action                                                         = &_Action{_ShiftAction, _State178, 0}
	_GotoState179Action                                                         = &_Action{_ShiftAction, _State179, 0}
	_GotoState180Action                                                         = &_Action{_ShiftAction, _State180, 0}
	_GotoState181Action                                                         = &_Action{_ShiftAction, _State181, 0}
	_GotoState182Action                                                         = &_Action{_ShiftAction, _State182, 0}
	_GotoState183Action                                                         = &_Action{_ShiftAction, _State183, 0}
	_GotoState184Action                                                         = &_Action{_ShiftAction, _State184, 0}
	_GotoState185Action                                                         = &_Action{_ShiftAction, _State185, 0}
	_GotoState186Action                                                         = &_Action{_ShiftAction, _State186, 0}
	_GotoState187Action                                                         = &_Action{_ShiftAction, _State187, 0}
	_GotoState188Action                                                         = &_Action{_ShiftAction, _State188, 0}
	_GotoState189Action                                                         = &_Action{_ShiftAction, _State189, 0}
	_GotoState190Action                                                         = &_Action{_ShiftAction, _State190, 0}
	_GotoState191Action                                                         = &_Action{_ShiftAction, _State191, 0}
	_GotoState192Action                                                         = &_Action{_ShiftAction, _State192, 0}
	_GotoState193Action                                                         = &_Action{_ShiftAction, _State193, 0}
	_GotoState194Action                                                         = &_Action{_ShiftAction, _State194, 0}
	_GotoState195Action                                                         = &_Action{_ShiftAction, _State195, 0}
	_GotoState196Action                                                         = &_Action{_ShiftAction, _State196, 0}
	_GotoState197Action                                                         = &_Action{_ShiftAction, _State197, 0}
	_GotoState198Action                                                         = &_Action{_ShiftAction, _State198, 0}
	_GotoState199Action                                                         = &_Action{_ShiftAction, _State199, 0}
	_GotoState200Action                                                         = &_Action{_ShiftAction, _State200, 0}
	_GotoState201Action                                                         = &_Action{_ShiftAction, _State201, 0}
	_GotoState202Action                                                         = &_Action{_ShiftAction, _State202, 0}
	_GotoState203Action                                                         = &_Action{_ShiftAction, _State203, 0}
	_GotoState204Action                                                         = &_Action{_ShiftAction, _State204, 0}
	_GotoState205Action                                                         = &_Action{_ShiftAction, _State205, 0}
	_GotoState206Action                                                         = &_Action{_ShiftAction, _State206, 0}
	_GotoState207Action                                                         = &_Action{_ShiftAction, _State207, 0}
	_GotoState208Action                                                         = &_Action{_ShiftAction, _State208, 0}
	_GotoState209Action                                                         = &_Action{_ShiftAction, _State209, 0}
	_GotoState210Action                                                         = &_Action{_ShiftAction, _State210, 0}
	_GotoState211Action                                                         = &_Action{_ShiftAction, _State211, 0}
	_GotoState212Action                                                         = &_Action{_ShiftAction, _State212, 0}
	_GotoState213Action                                                         = &_Action{_ShiftAction, _State213, 0}
	_GotoState214Action                                                         = &_Action{_ShiftAction, _State214, 0}
	_GotoState215Action                                                         = &_Action{_ShiftAction, _State215, 0}
	_GotoState216Action                                                         = &_Action{_ShiftAction, _State216, 0}
	_GotoState217Action                                                         = &_Action{_ShiftAction, _State217, 0}
	_GotoState218Action                                                         = &_Action{_ShiftAction, _State218, 0}
	_GotoState219Action                                                         = &_Action{_ShiftAction, _State219, 0}
	_GotoState220Action                                                         = &_Action{_ShiftAction, _State220, 0}
	_GotoState221Action                                                         = &_Action{_ShiftAction, _State221, 0}
	_GotoState222Action                                                         = &_Action{_ShiftAction, _State222, 0}
	_GotoState223Action                                                         = &_Action{_ShiftAction, _State223, 0}
	_GotoState224Action                                                         = &_Action{_ShiftAction, _State224, 0}
	_GotoState225Action                                                         = &_Action{_ShiftAction, _State225, 0}
	_GotoState226Action                                                         = &_Action{_ShiftAction, _State226, 0}
	_GotoState227Action                                                         = &_Action{_ShiftAction, _State227, 0}
	_GotoState228Action                                                         = &_Action{_ShiftAction, _State228, 0}
	_GotoState229Action                                                         = &_Action{_ShiftAction, _State229, 0}
	_GotoState230Action                                                         = &_Action{_ShiftAction, _State230, 0}
	_GotoState231Action                                                         = &_Action{_ShiftAction, _State231, 0}
	_GotoState232Action                                                         = &_Action{_ShiftAction, _State232, 0}
	_GotoState233Action                                                         = &_Action{_ShiftAction, _State233, 0}
	_GotoState234Action                                                         = &_Action{_ShiftAction, _State234, 0}
	_GotoState235Action                                                         = &_Action{_ShiftAction, _State235, 0}
	_GotoState236Action                                                         = &_Action{_ShiftAction, _State236, 0}
	_GotoState237Action                                                         = &_Action{_ShiftAction, _State237, 0}
	_GotoState238Action                                                         = &_Action{_ShiftAction, _State238, 0}
	_GotoState239Action                                                         = &_Action{_ShiftAction, _State239, 0}
	_GotoState240Action                                                         = &_Action{_ShiftAction, _State240, 0}
	_GotoState241Action                                                         = &_Action{_ShiftAction, _State241, 0}
	_GotoState242Action                                                         = &_Action{_ShiftAction, _State242, 0}
	_GotoState243Action                                                         = &_Action{_ShiftAction, _State243, 0}
	_GotoState244Action                                                         = &_Action{_ShiftAction, _State244, 0}
	_GotoState245Action                                                         = &_Action{_ShiftAction, _State245, 0}
	_GotoState246Action                                                         = &_Action{_ShiftAction, _State246, 0}
	_GotoState247Action                                                         = &_Action{_ShiftAction, _State247, 0}
	_GotoState248Action                                                         = &_Action{_ShiftAction, _State248, 0}
	_GotoState249Action                                                         = &_Action{_ShiftAction, _State249, 0}
	_GotoState250Action                                                         = &_Action{_ShiftAction, _State250, 0}
	_GotoState251Action                                                         = &_Action{_ShiftAction, _State251, 0}
	_GotoState252Action                                                         = &_Action{_ShiftAction, _State252, 0}
	_GotoState253Action                                                         = &_Action{_ShiftAction, _State253, 0}
	_GotoState254Action                                                         = &_Action{_ShiftAction, _State254, 0}
	_GotoState255Action                                                         = &_Action{_ShiftAction, _State255, 0}
	_GotoState256Action                                                         = &_Action{_ShiftAction, _State256, 0}
	_GotoState257Action                                                         = &_Action{_ShiftAction, _State257, 0}
	_GotoState258Action                                                         = &_Action{_ShiftAction, _State258, 0}
	_GotoState259Action                                                         = &_Action{_ShiftAction, _State259, 0}
	_GotoState260Action                                                         = &_Action{_ShiftAction, _State260, 0}
	_GotoState261Action                                                         = &_Action{_ShiftAction, _State261, 0}
	_GotoState262Action                                                         = &_Action{_ShiftAction, _State262, 0}
	_GotoState263Action                                                         = &_Action{_ShiftAction, _State263, 0}
	_GotoState264Action                                                         = &_Action{_ShiftAction, _State264, 0}
	_GotoState265Action                                                         = &_Action{_ShiftAction, _State265, 0}
	_GotoState266Action                                                         = &_Action{_ShiftAction, _State266, 0}
	_GotoState267Action                                                         = &_Action{_ShiftAction, _State267, 0}
	_GotoState268Action                                                         = &_Action{_ShiftAction, _State268, 0}
	_GotoState269Action                                                         = &_Action{_ShiftAction, _State269, 0}
	_GotoState270Action                                                         = &_Action{_ShiftAction, _State270, 0}
	_GotoState271Action                                                         = &_Action{_ShiftAction, _State271, 0}
	_GotoState272Action                                                         = &_Action{_ShiftAction, _State272, 0}
	_ReduceTrueToLiteralAction                                                  = &_Action{_ReduceAction, 0, _ReduceTrueToLiteral}
	_ReduceFalseToLiteralAction                                                 = &_Action{_ReduceAction, 0, _ReduceFalseToLiteral}
	_ReduceIntegerLiteralToLiteralAction                                        = &_Action{_ReduceAction, 0, _ReduceIntegerLiteralToLiteral}
	_ReduceFloatLiteralToLiteralAction                                          = &_Action{_ReduceAction, 0, _ReduceFloatLiteralToLiteral}
	_ReduceRuneLiteralToLiteralAction                                           = &_Action{_ReduceAction, 0, _ReduceRuneLiteralToLiteral}
	_ReduceStringLiteralToLiteralAction                                         = &_Action{_ReduceAction, 0, _ReduceStringLiteralToLiteral}
	_ReduceToAnonymousFuncExprAction                                            = &_Action{_ReduceAction, 0, _ReduceToAnonymousFuncExpr}
	_ReduceExplicitToAnonymousStructExprAction                                  = &_Action{_ReduceAction, 0, _ReduceExplicitToAnonymousStructExpr}
	_ReduceImplicitToAnonymousStructExprAction                                  = &_Action{_ReduceAction, 0, _ReduceImplicitToAnonymousStructExpr}
	_ReduceLiteralToAtomExprAction                                              = &_Action{_ReduceAction, 0, _ReduceLiteralToAtomExpr}
	_ReduceIdentifierToAtomExprAction                                           = &_Action{_ReduceAction, 0, _ReduceIdentifierToAtomExpr}
	_ReduceBlockExprToAtomExprAction                                            = &_Action{_ReduceAction, 0, _ReduceBlockExprToAtomExpr}
	_ReduceAnonymousFuncExprToAtomExprAction                                    = &_Action{_ReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}
	_ReduceAnonymousStructExprToAtomExprAction                                  = &_Action{_ReduceAction, 0, _ReduceAnonymousStructExprToAtomExpr}
	_ReduceLexErrorToAtomExprAction                                             = &_Action{_ReduceAction, 0, _ReduceLexErrorToAtomExpr}
	_ReduceValueTypeToGenericArgumentsAction                                    = &_Action{_ReduceAction, 0, _ReduceValueTypeToGenericArguments}
	_ReduceAddToGenericArgumentsAction                                          = &_Action{_ReduceAction, 0, _ReduceAddToGenericArguments}
	_ReduceGenericArgumentsToOptionalGenericArgumentsAction                     = &_Action{_ReduceAction, 0, _ReduceGenericArgumentsToOptionalGenericArguments}
	_ReduceNilToOptionalGenericArgumentsAction                                  = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericArguments}
	_ReduceBindingToOptionalGenericBindingAction                                = &_Action{_ReduceAction, 0, _ReduceBindingToOptionalGenericBinding}
	_ReduceNilToOptionalGenericBindingAction                                    = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericBinding}
	_ReducePositionalToArgumentAction                                           = &_Action{_ReduceAction, 0, _ReducePositionalToArgument}
	_ReduceArgumentToArgumentsAction                                            = &_Action{_ReduceAction, 0, _ReduceArgumentToArguments}
	_ReduceAddToArgumentsAction                                                 = &_Action{_ReduceAction, 0, _ReduceAddToArguments}
	_ReduceArgumentsToOptionalArgumentsAction                                   = &_Action{_ReduceAction, 0, _ReduceArgumentsToOptionalArguments}
	_ReduceNilToOptionalArgumentsAction                                         = &_Action{_ReduceAction, 0, _ReduceNilToOptionalArguments}
	_ReduceToCallExprAction                                                     = &_Action{_ReduceAction, 0, _ReduceToCallExpr}
	_ReduceAtomExprToAccessExprAction                                           = &_Action{_ReduceAction, 0, _ReduceAtomExprToAccessExpr}
	_ReduceAccessToAccessExprAction                                             = &_Action{_ReduceAction, 0, _ReduceAccessToAccessExpr}
	_ReduceCallExprToAccessExprAction                                           = &_Action{_ReduceAction, 0, _ReduceCallExprToAccessExpr}
	_ReduceIndexToAccessExprAction                                              = &_Action{_ReduceAction, 0, _ReduceIndexToAccessExpr}
	_ReduceAccessExprToPostfixUnaryExprAction                                   = &_Action{_ReduceAction, 0, _ReduceAccessExprToPostfixUnaryExpr}
	_ReduceQuestionToPostfixUnaryExprAction                                     = &_Action{_ReduceAction, 0, _ReduceQuestionToPostfixUnaryExpr}
	_ReduceNotToPrefixUnaryOpAction                                             = &_Action{_ReduceAction, 0, _ReduceNotToPrefixUnaryOp}
	_ReduceBitNegToPrefixUnaryOpAction                                          = &_Action{_ReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}
	_ReduceSubToPrefixUnaryOpAction                                             = &_Action{_ReduceAction, 0, _ReduceSubToPrefixUnaryOp}
	_ReduceMulToPrefixUnaryOpAction                                             = &_Action{_ReduceAction, 0, _ReduceMulToPrefixUnaryOp}
	_ReduceBitAndToPrefixUnaryOpAction                                          = &_Action{_ReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}
	_ReducePostfixUnaryExprToPrefixUnaryExprAction                              = &_Action{_ReduceAction, 0, _ReducePostfixUnaryExprToPrefixUnaryExpr}
	_ReducePrefixOpToPrefixUnaryExprAction                                      = &_Action{_ReduceAction, 0, _ReducePrefixOpToPrefixUnaryExpr}
	_ReduceMulToMulOpAction                                                     = &_Action{_ReduceAction, 0, _ReduceMulToMulOp}
	_ReduceDivToMulOpAction                                                     = &_Action{_ReduceAction, 0, _ReduceDivToMulOp}
	_ReduceModToMulOpAction                                                     = &_Action{_ReduceAction, 0, _ReduceModToMulOp}
	_ReduceBitAndToMulOpAction                                                  = &_Action{_ReduceAction, 0, _ReduceBitAndToMulOp}
	_ReduceBitLshiftToMulOpAction                                               = &_Action{_ReduceAction, 0, _ReduceBitLshiftToMulOp}
	_ReduceBitRshiftToMulOpAction                                               = &_Action{_ReduceAction, 0, _ReduceBitRshiftToMulOp}
	_ReducePrefixUnaryExprToMulExprAction                                       = &_Action{_ReduceAction, 0, _ReducePrefixUnaryExprToMulExpr}
	_ReduceOpToMulExprAction                                                    = &_Action{_ReduceAction, 0, _ReduceOpToMulExpr}
	_ReduceAddToAddOpAction                                                     = &_Action{_ReduceAction, 0, _ReduceAddToAddOp}
	_ReduceSubToAddOpAction                                                     = &_Action{_ReduceAction, 0, _ReduceSubToAddOp}
	_ReduceBitOrToAddOpAction                                                   = &_Action{_ReduceAction, 0, _ReduceBitOrToAddOp}
	_ReduceBitXorToAddOpAction                                                  = &_Action{_ReduceAction, 0, _ReduceBitXorToAddOp}
	_ReduceMulExprToAddExprAction                                               = &_Action{_ReduceAction, 0, _ReduceMulExprToAddExpr}
	_ReduceOpToAddExprAction                                                    = &_Action{_ReduceAction, 0, _ReduceOpToAddExpr}
	_ReduceEqualToCmpOpAction                                                   = &_Action{_ReduceAction, 0, _ReduceEqualToCmpOp}
	_ReduceNotEqualToCmpOpAction                                                = &_Action{_ReduceAction, 0, _ReduceNotEqualToCmpOp}
	_ReduceLessToCmpOpAction                                                    = &_Action{_ReduceAction, 0, _ReduceLessToCmpOp}
	_ReduceLessOrEqualToCmpOpAction                                             = &_Action{_ReduceAction, 0, _ReduceLessOrEqualToCmpOp}
	_ReduceGreaterToCmpOpAction                                                 = &_Action{_ReduceAction, 0, _ReduceGreaterToCmpOp}
	_ReduceGreaterOrEqualToCmpOpAction                                          = &_Action{_ReduceAction, 0, _ReduceGreaterOrEqualToCmpOp}
	_ReduceAddExprToCmpExprAction                                               = &_Action{_ReduceAction, 0, _ReduceAddExprToCmpExpr}
	_ReduceOpToCmpExprAction                                                    = &_Action{_ReduceAction, 0, _ReduceOpToCmpExpr}
	_ReduceCmpExprToAndExprAction                                               = &_Action{_ReduceAction, 0, _ReduceCmpExprToAndExpr}
	_ReduceOpToAndExprAction                                                    = &_Action{_ReduceAction, 0, _ReduceOpToAndExpr}
	_ReduceAndExprToOrExprAction                                                = &_Action{_ReduceAction, 0, _ReduceAndExprToOrExpr}
	_ReduceOpToOrExprAction                                                     = &_Action{_ReduceAction, 0, _ReduceOpToOrExpr}
	_ReduceToSequenceExprAction                                                 = &_Action{_ReduceAction, 0, _ReduceToSequenceExpr}
	_ReduceJumpLabelToOptionalJumpLabelAction                                   = &_Action{_ReduceAction, 0, _ReduceJumpLabelToOptionalJumpLabel}
	_ReduceUnlabelledToOptionalJumpLabelAction                                  = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalJumpLabel}
	_ReduceExpressionOrImplicitStructToOptionalExpressionOrImplicitStructAction = &_Action{_ReduceAction, 0, _ReduceExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct}
	_ReduceNilToOptionalExpressionOrImplicitStructAction                        = &_Action{_ReduceAction, 0, _ReduceNilToOptionalExpressionOrImplicitStruct}
	_ReduceReturnToJumpTypeAction                                               = &_Action{_ReduceAction, 0, _ReduceReturnToJumpType}
	_ReduceBreakToJumpTypeAction                                                = &_Action{_ReduceAction, 0, _ReduceBreakToJumpType}
	_ReduceContinueToJumpTypeAction                                             = &_Action{_ReduceAction, 0, _ReduceContinueToJumpType}
	_ReduceAddOneAssignToOpOneAssignAction                                      = &_Action{_ReduceAction, 0, _ReduceAddOneAssignToOpOneAssign}
	_ReduceSubOneAssignToOpOneAssignAction                                      = &_Action{_ReduceAction, 0, _ReduceSubOneAssignToOpOneAssign}
	_ReduceAddAssignToBinaryOpAssignAction                                      = &_Action{_ReduceAction, 0, _ReduceAddAssignToBinaryOpAssign}
	_ReduceSubAssignToBinaryOpAssignAction                                      = &_Action{_ReduceAction, 0, _ReduceSubAssignToBinaryOpAssign}
	_ReduceMulAssignToBinaryOpAssignAction                                      = &_Action{_ReduceAction, 0, _ReduceMulAssignToBinaryOpAssign}
	_ReduceDivAssignToBinaryOpAssignAction                                      = &_Action{_ReduceAction, 0, _ReduceDivAssignToBinaryOpAssign}
	_ReduceModAssignToBinaryOpAssignAction                                      = &_Action{_ReduceAction, 0, _ReduceModAssignToBinaryOpAssign}
	_ReduceBitNegAssignToBinaryOpAssignAction                                   = &_Action{_ReduceAction, 0, _ReduceBitNegAssignToBinaryOpAssign}
	_ReduceBitAndAssignToBinaryOpAssignAction                                   = &_Action{_ReduceAction, 0, _ReduceBitAndAssignToBinaryOpAssign}
	_ReduceBitOrAssignToBinaryOpAssignAction                                    = &_Action{_ReduceAction, 0, _ReduceBitOrAssignToBinaryOpAssign}
	_ReduceBitXorAssignToBinaryOpAssignAction                                   = &_Action{_ReduceAction, 0, _ReduceBitXorAssignToBinaryOpAssign}
	_ReduceBitLshiftAssignToBinaryOpAssignAction                                = &_Action{_ReduceAction, 0, _ReduceBitLshiftAssignToBinaryOpAssign}
	_ReduceBitRshiftAssignToBinaryOpAssignAction                                = &_Action{_ReduceAction, 0, _ReduceBitRshiftAssignToBinaryOpAssign}
	_ReduceExpressionToExpressionOrImplicitStructAction                         = &_Action{_ReduceAction, 0, _ReduceExpressionToExpressionOrImplicitStruct}
	_ReduceImplicitStructToExpressionOrImplicitStructAction                     = &_Action{_ReduceAction, 0, _ReduceImplicitStructToExpressionOrImplicitStruct}
	_ReduceToUnsafeStatementAction                                              = &_Action{_ReduceAction, 0, _ReduceToUnsafeStatement}
	_ReduceUnsafeStatementToStatementBodyAction                                 = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToStatementBody}
	_ReduceExpressionOrImplicitStructToStatementBodyAction                      = &_Action{_ReduceAction, 0, _ReduceExpressionOrImplicitStructToStatementBody}
	_ReduceAsyncToStatementBodyAction                                           = &_Action{_ReduceAction, 0, _ReduceAsyncToStatementBody}
	_ReduceJumpToStatementBodyAction                                            = &_Action{_ReduceAction, 0, _ReduceJumpToStatementBody}
	_ReduceOpOneAssignToStatementBodyAction                                     = &_Action{_ReduceAction, 0, _ReduceOpOneAssignToStatementBody}
	_ReduceBinaryOpAssignToStatementBodyAction                                  = &_Action{_ReduceAction, 0, _ReduceBinaryOpAssignToStatementBody}
	_ReduceImplicitToStatementAction                                            = &_Action{_ReduceAction, 0, _ReduceImplicitToStatement}
	_ReduceExplicitToStatementAction                                            = &_Action{_ReduceAction, 0, _ReduceExplicitToStatement}
	_ReduceEmptyListToStatementsAction                                          = &_Action{_ReduceAction, 0, _ReduceEmptyListToStatements}
	_ReduceAddToStatementsAction                                                = &_Action{_ReduceAction, 0, _ReduceAddToStatements}
	_ReduceLabelDeclToOptionalLabelDeclAction                                   = &_Action{_ReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}
	_ReduceUnlabelledToOptionalLabelDeclAction                                  = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}
	_ReduceToBlockBodyAction                                                    = &_Action{_ReduceAction, 0, _ReduceToBlockBody}
	_ReduceToBlockExprAction                                                    = &_Action{_ReduceAction, 0, _ReduceToBlockExpr}
	_ReduceNoElseToIfExprAction                                                 = &_Action{_ReduceAction, 0, _ReduceNoElseToIfExpr}
	_ReduceIfElseToIfExprAction                                                 = &_Action{_ReduceAction, 0, _ReduceIfElseToIfExpr}
	_ReduceMultiIfElseToIfExprAction                                            = &_Action{_ReduceAction, 0, _ReduceMultiIfElseToIfExpr}
	_ReduceToSwitchExprAction                                                   = &_Action{_ReduceAction, 0, _ReduceToSwitchExpr}
	_ReduceInfiniteToLoopExprAction                                             = &_Action{_ReduceAction, 0, _ReduceInfiniteToLoopExpr}
	_ReduceWhileToLoopExprAction                                                = &_Action{_ReduceAction, 0, _ReduceWhileToLoopExpr}
	_ReduceSequenceExprToExpressionAction                                       = &_Action{_ReduceAction, 0, _ReduceSequenceExprToExpression}
	_ReduceIfExprToExpressionAction                                             = &_Action{_ReduceAction, 0, _ReduceIfExprToExpression}
	_ReduceSwitchExprToExpressionAction                                         = &_Action{_ReduceAction, 0, _ReduceSwitchExprToExpression}
	_ReduceLoopExprToExpressionAction                                           = &_Action{_ReduceAction, 0, _ReduceLoopExprToExpression}
	_ReduceUnconstrainedToGenericParameterAction                                = &_Action{_ReduceAction, 0, _ReduceUnconstrainedToGenericParameter}
	_ReduceConstrainedToGenericParameterAction                                  = &_Action{_ReduceAction, 0, _ReduceConstrainedToGenericParameter}
	_ReduceGenericParameterToGenericParametersAction                            = &_Action{_ReduceAction, 0, _ReduceGenericParameterToGenericParameters}
	_ReduceAddToGenericParametersAction                                         = &_Action{_ReduceAction, 0, _ReduceAddToGenericParameters}
	_ReduceGenericParametersToOptionalGenericParametersAction                   = &_Action{_ReduceAction, 0, _ReduceGenericParametersToOptionalGenericParameters}
	_ReduceNilToOptionalGenericParametersAction                                 = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericParameters}
	_ReduceGenericToOptionalGenericParameterDeclAction                          = &_Action{_ReduceAction, 0, _ReduceGenericToOptionalGenericParameterDecl}
	_ReduceNilToOptionalGenericParameterDeclAction                              = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericParameterDecl}
	_ReduceNamedToAtomTypeAction                                                = &_Action{_ReduceAction, 0, _ReduceNamedToAtomType}
	_ReduceExplicitStructTypeToAtomTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceExplicitStructTypeToAtomType}
	_ReduceImplicitStructTypeToAtomTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceImplicitStructTypeToAtomType}
	_ReduceExplicitEnumTypeToAtomTypeAction                                     = &_Action{_ReduceAction, 0, _ReduceExplicitEnumTypeToAtomType}
	_ReduceTraitTypeToAtomTypeAction                                            = &_Action{_ReduceAction, 0, _ReduceTraitTypeToAtomType}
	_ReduceAtomTypeToTraitableTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceAtomTypeToTraitableType}
	_ReduceMethodInterfaceToTraitableTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceMethodInterfaceToTraitableType}
	_ReduceTraitToTraitableTypeAction                                           = &_Action{_ReduceAction, 0, _ReduceTraitToTraitableType}
	_ReduceTraitableTypeToTraitMulTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceTraitableTypeToTraitMulType}
	_ReduceIntersectToTraitMulTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceIntersectToTraitMulType}
	_ReduceTraitMulTypeToTraitAddTypeAction                                     = &_Action{_ReduceAction, 0, _ReduceTraitMulTypeToTraitAddType}
	_ReduceUnionToTraitAddTypeAction                                            = &_Action{_ReduceAction, 0, _ReduceUnionToTraitAddType}
	_ReduceDifferenceToTraitAddTypeAction                                       = &_Action{_ReduceAction, 0, _ReduceDifferenceToTraitAddType}
	_ReduceExplicitToFieldDeclAction                                            = &_Action{_ReduceAction, 0, _ReduceExplicitToFieldDecl}
	_ReduceImplicitToFieldDeclAction                                            = &_Action{_ReduceAction, 0, _ReduceImplicitToFieldDecl}
	_ReduceFieldDeclToFieldDeclsAction                                          = &_Action{_ReduceAction, 0, _ReduceFieldDeclToFieldDecls}
	_ReduceImplicitToFieldDeclsAction                                           = &_Action{_ReduceAction, 0, _ReduceImplicitToFieldDecls}
	_ReduceExplicitToFieldDeclsAction                                           = &_Action{_ReduceAction, 0, _ReduceExplicitToFieldDecls}
	_ReduceFieldDeclsToOptionalFieldDeclsAction                                 = &_Action{_ReduceAction, 0, _ReduceFieldDeclsToOptionalFieldDecls}
	_ReduceNilToOptionalFieldDeclsAction                                        = &_Action{_ReduceAction, 0, _ReduceNilToOptionalFieldDecls}
	_ReduceToImplicitStructTypeAction                                           = &_Action{_ReduceAction, 0, _ReduceToImplicitStructType}
	_ReduceToExplicitStructTypeAction                                           = &_Action{_ReduceAction, 0, _ReduceToExplicitStructType}
	_ReduceToTraitTypeAction                                                    = &_Action{_ReduceAction, 0, _ReduceToTraitType}
	_ReduceToExplicitEnumTypeAction                                             = &_Action{_ReduceAction, 0, _ReduceToExplicitEnumType}
	_ReduceTraitAddTypeToFuncTypeAction                                         = &_Action{_ReduceAction, 0, _ReduceTraitAddTypeToFuncType}
	_ReduceReferenceToFuncTypeAction                                            = &_Action{_ReduceAction, 0, _ReduceReferenceToFuncType}
	_ReduceExplicitFuncTypeToFuncTypeAction                                     = &_Action{_ReduceAction, 0, _ReduceExplicitFuncTypeToFuncType}
	_ReduceFuncTypeToValueTypeAction                                            = &_Action{_ReduceAction, 0, _ReduceFuncTypeToValueType}
	_ReduceImplicitEnumToValueTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceImplicitEnumToValueType}
	_ReduceDefinitionToTypeDeclAction                                           = &_Action{_ReduceAction, 0, _ReduceDefinitionToTypeDecl}
	_ReduceAliasToTypeDeclAction                                                = &_Action{_ReduceAction, 0, _ReduceAliasToTypeDecl}
	_ReduceFieldDeclToParameterDeclAction                                       = &_Action{_ReduceAction, 0, _ReduceFieldDeclToParameterDecl}
	_ReduceInferredToParameterDeclAction                                        = &_Action{_ReduceAction, 0, _ReduceInferredToParameterDecl}
	_ReduceParameterDeclToNonEmptyParametersAction                              = &_Action{_ReduceAction, 0, _ReduceParameterDeclToNonEmptyParameters}
	_ReduceAddToNonEmptyParametersAction                                        = &_Action{_ReduceAction, 0, _ReduceAddToNonEmptyParameters}
	_ReduceExplicitToVarargAction                                               = &_Action{_ReduceAction, 0, _ReduceExplicitToVararg}
	_ReduceImplicitToVarargAction                                               = &_Action{_ReduceAction, 0, _ReduceImplicitToVararg}
	_ReduceInferredToVarargAction                                               = &_Action{_ReduceAction, 0, _ReduceInferredToVararg}
	_ReduceVarargToOptionalVarargAction                                         = &_Action{_ReduceAction, 0, _ReduceVarargToOptionalVararg}
	_ReduceVararg2ToOptionalVarargAction                                        = &_Action{_ReduceAction, 0, _ReduceVararg2ToOptionalVararg}
	_ReduceNilToOptionalVarargAction                                            = &_Action{_ReduceAction, 0, _ReduceNilToOptionalVararg}
	_ReduceNonEmptyParametersToParametersAction                                 = &_Action{_ReduceAction, 0, _ReduceNonEmptyParametersToParameters}
	_ReduceMixedToParametersAction                                              = &_Action{_ReduceAction, 0, _ReduceMixedToParameters}
	_ReduceOptionalVarargToParametersAction                                     = &_Action{_ReduceAction, 0, _ReduceOptionalVarargToParameters}
	_ReduceTypedToImplicitFuncTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceTypedToImplicitFuncType}
	_ReduceInferredToImplicitFuncTypeAction                                     = &_Action{_ReduceAction, 0, _ReduceInferredToImplicitFuncType}
	_ReduceImplicitUnitToImplicitFuncTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceImplicitUnitToImplicitFuncType}
	_ReduceToExplicitFuncTypeAction                                             = &_Action{_ReduceAction, 0, _ReduceToExplicitFuncType}
	_ReduceReceiverToOptionalReceiverAction                                     = &_Action{_ReduceAction, 0, _ReduceReceiverToOptionalReceiver}
	_ReduceNilToOptionalReceiverAction                                          = &_Action{_ReduceAction, 0, _ReduceNilToOptionalReceiver}
	_ReduceToImplicitFuncDeclAction                                             = &_Action{_ReduceAction, 0, _ReduceToImplicitFuncDecl}
	_ReduceToExplicitFuncDeclAction                                             = &_Action{_ReduceAction, 0, _ReduceToExplicitFuncDecl}
	_ReduceToFuncDefAction                                                      = &_Action{_ReduceAction, 0, _ReduceToFuncDef}
	_ReduceNoSpecToPackageDeclAction                                            = &_Action{_ReduceAction, 0, _ReduceNoSpecToPackageDecl}
	_ReduceWithSpecToPackageDeclAction                                          = &_Action{_ReduceAction, 0, _ReduceWithSpecToPackageDecl}
	_ReduceToPackageStatementBodyAction                                         = &_Action{_ReduceAction, 0, _ReduceToPackageStatementBody}
	_ReduceImplicitToPackageStatementAction                                     = &_Action{_ReduceAction, 0, _ReduceImplicitToPackageStatement}
	_ReduceExplicitToPackageStatementAction                                     = &_Action{_ReduceAction, 0, _ReduceExplicitToPackageStatement}
	_ReduceEmptyListToPackageStatementsAction                                   = &_Action{_ReduceAction, 0, _ReduceEmptyListToPackageStatements}
	_ReduceAddToPackageStatementsAction                                         = &_Action{_ReduceAction, 0, _ReduceAddToPackageStatements}
	_ReduceSpacesToLexInternalTokensAction                                      = &_Action{_ReduceAction, 0, _ReduceSpacesToLexInternalTokens}
	_ReduceCommentToLexInternalTokensAction                                     = &_Action{_ReduceAction, 0, _ReduceCommentToLexInternalTokens}
)

var _ActionTable = _ActionTableType{
	{_State6, _EndMarker}:                               &_Action{_AcceptAction, 0, 0},
	{_State7, _EndMarker}:                               &_Action{_AcceptAction, 0, 0},
	{_State8, _EndMarker}:                               &_Action{_AcceptAction, 0, 0},
	{_State9, _EndMarker}:                               &_Action{_AcceptAction, 0, 0},
	{_State10, _EndMarker}:                              &_Action{_AcceptAction, 0, 0},
	{_State1, PackageToken}:                             _GotoState11Action,
	{_State1, PackageDeclType}:                          _GotoState6Action,
	{_State2, TypeToken}:                                _GotoState12Action,
	{_State2, TypeDeclType}:                             _GotoState7Action,
	{_State3, FuncToken}:                                _GotoState13Action,
	{_State3, ExplicitFuncDeclType}:                     _GotoState14Action,
	{_State3, FuncDefType}:                              _GotoState8Action,
	{_State4, IntegerLiteralToken}:                      _GotoState21Action,
	{_State4, FloatLiteralToken}:                        _GotoState18Action,
	{_State4, RuneLiteralToken}:                         _GotoState27Action,
	{_State4, StringLiteralToken}:                       _GotoState28Action,
	{_State4, IdentifierToken}:                          _GotoState20Action,
	{_State4, TrueToken}:                                _GotoState31Action,
	{_State4, FalseToken}:                               _GotoState17Action,
	{_State4, StructToken}:                              _GotoState29Action,
	{_State4, FuncToken}:                                _GotoState19Action,
	{_State4, LparenToken}:                              _GotoState24Action,
	{_State4, LabelDeclToken}:                           _GotoState22Action,
	{_State4, NotToken}:                                 _GotoState26Action,
	{_State4, SubToken}:                                 _GotoState30Action,
	{_State4, MulToken}:                                 _GotoState25Action,
	{_State4, BitNegToken}:                              _GotoState16Action,
	{_State4, BitAndToken}:                              _GotoState15Action,
	{_State4, LexErrorToken}:                            _GotoState23Action,
	{_State4, LiteralType}:                              _GotoState43Action,
	{_State4, AnonymousFuncExprType}:                    _GotoState35Action,
	{_State4, AnonymousStructExprType}:                  _GotoState36Action,
	{_State4, AtomExprType}:                             _GotoState37Action,
	{_State4, CallExprType}:                             _GotoState39Action,
	{_State4, AccessExprType}:                           _GotoState32Action,
	{_State4, PostfixUnaryExprType}:                     _GotoState47Action,
	{_State4, PrefixUnaryOpType}:                        _GotoState49Action,
	{_State4, PrefixUnaryExprType}:                      _GotoState48Action,
	{_State4, MulExprType}:                              _GotoState44Action,
	{_State4, AddExprType}:                              _GotoState33Action,
	{_State4, CmpExprType}:                              _GotoState40Action,
	{_State4, AndExprType}:                              _GotoState34Action,
	{_State4, OrExprType}:                               _GotoState46Action,
	{_State4, SequenceExprType}:                         _GotoState50Action,
	{_State4, OptionalLabelDeclType}:                    _GotoState45Action,
	{_State4, BlockExprType}:                            _GotoState38Action,
	{_State4, ExpressionType}:                           _GotoState9Action,
	{_State4, ExplicitStructTypeType}:                   _GotoState42Action,
	{_State4, ExplicitFuncTypeType}:                     _GotoState41Action,
	{_State5, SpacesToken}:                              _GotoState52Action,
	{_State5, CommentToken}:                             _GotoState51Action,
	{_State5, LexInternalTokensType}:                    _GotoState10Action,
	{_State11, IdentifierToken}:                         _GotoState53Action,
	{_State12, IdentifierToken}:                         _GotoState54Action,
	{_State13, LparenToken}:                             _GotoState55Action,
	{_State13, OptionalReceiverType}:                    _GotoState56Action,
	{_State14, LbraceToken}:                             _GotoState57Action,
	{_State14, BlockBodyType}:                           _GotoState58Action,
	{_State19, LparenToken}:                             _GotoState59Action,
	{_State19, ImplicitFuncTypeType}:                    _GotoState60Action,
	{_State24, IntegerLiteralToken}:                     _GotoState21Action,
	{_State24, FloatLiteralToken}:                       _GotoState18Action,
	{_State24, RuneLiteralToken}:                        _GotoState27Action,
	{_State24, StringLiteralToken}:                      _GotoState28Action,
	{_State24, IdentifierToken}:                         _GotoState20Action,
	{_State24, TrueToken}:                               _GotoState31Action,
	{_State24, FalseToken}:                              _GotoState17Action,
	{_State24, StructToken}:                             _GotoState29Action,
	{_State24, FuncToken}:                               _GotoState19Action,
	{_State24, LparenToken}:                             _GotoState24Action,
	{_State24, LabelDeclToken}:                          _GotoState22Action,
	{_State24, NotToken}:                                _GotoState26Action,
	{_State24, SubToken}:                                _GotoState30Action,
	{_State24, MulToken}:                                _GotoState25Action,
	{_State24, BitNegToken}:                             _GotoState16Action,
	{_State24, BitAndToken}:                             _GotoState15Action,
	{_State24, LexErrorToken}:                           _GotoState23Action,
	{_State24, LiteralType}:                             _GotoState43Action,
	{_State24, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State24, AnonymousStructExprType}:                 _GotoState36Action,
	{_State24, AtomExprType}:                            _GotoState37Action,
	{_State24, ArgumentType}:                            _GotoState61Action,
	{_State24, ArgumentsType}:                           _GotoState62Action,
	{_State24, CallExprType}:                            _GotoState39Action,
	{_State24, AccessExprType}:                          _GotoState32Action,
	{_State24, PostfixUnaryExprType}:                    _GotoState47Action,
	{_State24, PrefixUnaryOpType}:                       _GotoState49Action,
	{_State24, PrefixUnaryExprType}:                     _GotoState48Action,
	{_State24, MulExprType}:                             _GotoState44Action,
	{_State24, AddExprType}:                             _GotoState33Action,
	{_State24, CmpExprType}:                             _GotoState40Action,
	{_State24, AndExprType}:                             _GotoState34Action,
	{_State24, OrExprType}:                              _GotoState46Action,
	{_State24, SequenceExprType}:                        _GotoState50Action,
	{_State24, OptionalLabelDeclType}:                   _GotoState45Action,
	{_State24, BlockExprType}:                           _GotoState38Action,
	{_State24, ExpressionType}:                          _GotoState63Action,
	{_State24, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State24, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State29, LparenToken}:                             _GotoState64Action,
	{_State29, ImplicitStructTypeType}:                  _GotoState65Action,
	{_State32, LbracketToken}:                           _GotoState68Action,
	{_State32, DotToken}:                                _GotoState67Action,
	{_State32, QuestionToken}:                           _GotoState69Action,
	{_State32, DollarLbracketToken}:                     _GotoState66Action,
	{_State32, OptionalGenericBindingType}:              _GotoState70Action,
	{_State33, AddToken}:                                _GotoState71Action,
	{_State33, SubToken}:                                _GotoState74Action,
	{_State33, BitXorToken}:                             _GotoState73Action,
	{_State33, BitOrToken}:                              _GotoState72Action,
	{_State33, AddOpType}:                               _GotoState75Action,
	{_State34, AndToken}:                                _GotoState76Action,
	{_State40, EqualToken}:                              _GotoState77Action,
	{_State40, NotEqualToken}:                           _GotoState82Action,
	{_State40, LessToken}:                               _GotoState80Action,
	{_State40, LessOrEqualToken}:                        _GotoState81Action,
	{_State40, GreaterToken}:                            _GotoState78Action,
	{_State40, GreaterOrEqualToken}:                     _GotoState79Action,
	{_State40, CmpOpType}:                               _GotoState83Action,
	{_State41, LbraceToken}:                             _GotoState57Action,
	{_State41, BlockBodyType}:                           _GotoState84Action,
	{_State42, LparenToken}:                             _GotoState85Action,
	{_State44, MulToken}:                                _GotoState91Action,
	{_State44, DivToken}:                                _GotoState89Action,
	{_State44, ModToken}:                                _GotoState90Action,
	{_State44, BitAndToken}:                             _GotoState86Action,
	{_State44, BitLshiftToken}:                          _GotoState87Action,
	{_State44, BitRshiftToken}:                          _GotoState88Action,
	{_State44, MulOpType}:                               _GotoState92Action,
	{_State45, IfToken}:                                 _GotoState94Action,
	{_State45, SwitchToken}:                             _GotoState95Action,
	{_State45, ForToken}:                                _GotoState93Action,
	{_State45, LbraceToken}:                             _GotoState57Action,
	{_State45, BlockBodyType}:                           _GotoState96Action,
	{_State45, IfExprType}:                              _GotoState97Action,
	{_State45, SwitchExprType}:                          _GotoState99Action,
	{_State45, LoopExprType}:                            _GotoState98Action,
	{_State46, OrToken}:                                 _GotoState100Action,
	{_State49, IntegerLiteralToken}:                     _GotoState21Action,
	{_State49, FloatLiteralToken}:                       _GotoState18Action,
	{_State49, RuneLiteralToken}:                        _GotoState27Action,
	{_State49, StringLiteralToken}:                      _GotoState28Action,
	{_State49, IdentifierToken}:                         _GotoState20Action,
	{_State49, TrueToken}:                               _GotoState31Action,
	{_State49, FalseToken}:                              _GotoState17Action,
	{_State49, StructToken}:                             _GotoState29Action,
	{_State49, FuncToken}:                               _GotoState19Action,
	{_State49, LparenToken}:                             _GotoState24Action,
	{_State49, LabelDeclToken}:                          _GotoState22Action,
	{_State49, NotToken}:                                _GotoState26Action,
	{_State49, SubToken}:                                _GotoState30Action,
	{_State49, MulToken}:                                _GotoState25Action,
	{_State49, BitNegToken}:                             _GotoState16Action,
	{_State49, BitAndToken}:                             _GotoState15Action,
	{_State49, LexErrorToken}:                           _GotoState23Action,
	{_State49, LiteralType}:                             _GotoState43Action,
	{_State49, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State49, AnonymousStructExprType}:                 _GotoState36Action,
	{_State49, AtomExprType}:                            _GotoState37Action,
	{_State49, CallExprType}:                            _GotoState39Action,
	{_State49, AccessExprType}:                          _GotoState32Action,
	{_State49, PostfixUnaryExprType}:                    _GotoState47Action,
	{_State49, PrefixUnaryOpType}:                       _GotoState49Action,
	{_State49, PrefixUnaryExprType}:                     _GotoState102Action,
	{_State49, OptionalLabelDeclType}:                   _GotoState101Action,
	{_State49, BlockExprType}:                           _GotoState38Action,
	{_State49, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State49, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State53, LparenToken}:                             _GotoState103Action,
	{_State54, DollarLbracketToken}:                     _GotoState104Action,
	{_State54, EqualToken}:                              _GotoState105Action,
	{_State54, OptionalGenericParameterDeclType}:        _GotoState106Action,
	{_State55, IdentifierToken}:                         _GotoState111Action,
	{_State55, StructToken}:                             _GotoState29Action,
	{_State55, EnumToken}:                               _GotoState108Action,
	{_State55, TraitToken}:                              _GotoState112Action,
	{_State55, FuncToken}:                               _GotoState19Action,
	{_State55, LparenToken}:                             _GotoState64Action,
	{_State55, ExclaimToken}:                            _GotoState109Action,
	{_State55, ExclaimExclaimToken}:                     _GotoState110Action,
	{_State55, BitAndToken}:                             _GotoState107Action,
	{_State55, AtomTypeType}:                            _GotoState113Action,
	{_State55, TraitableTypeType}:                       _GotoState124Action,
	{_State55, TraitMulTypeType}:                        _GotoState122Action,
	{_State55, TraitAddTypeType}:                        _GotoState121Action,
	{_State55, FieldDeclType}:                           _GotoState117Action,
	{_State55, ImplicitStructTypeType}:                  _GotoState119Action,
	{_State55, ExplicitStructTypeType}:                  _GotoState116Action,
	{_State55, TraitTypeType}:                           _GotoState123Action,
	{_State55, ExplicitEnumTypeType}:                    _GotoState114Action,
	{_State55, FuncTypeType}:                            _GotoState118Action,
	{_State55, ValueTypeType}:                           _GotoState125Action,
	{_State55, ParameterDeclType}:                       _GotoState120Action,
	{_State55, ExplicitFuncTypeType}:                    _GotoState115Action,
	{_State56, IdentifierToken}:                         _GotoState126Action,
	{_State56, ImplicitFuncDeclType}:                    _GotoState127Action,
	{_State57, StatementsType}:                          _GotoState128Action,
	{_State59, IdentifierToken}:                         _GotoState130Action,
	{_State59, StructToken}:                             _GotoState29Action,
	{_State59, EnumToken}:                               _GotoState108Action,
	{_State59, TraitToken}:                              _GotoState112Action,
	{_State59, FuncToken}:                               _GotoState19Action,
	{_State59, LparenToken}:                             _GotoState64Action,
	{_State59, DotdotdotToken}:                          _GotoState129Action,
	{_State59, ExclaimToken}:                            _GotoState109Action,
	{_State59, ExclaimExclaimToken}:                     _GotoState110Action,
	{_State59, BitAndToken}:                             _GotoState107Action,
	{_State59, AtomTypeType}:                            _GotoState113Action,
	{_State59, TraitableTypeType}:                       _GotoState124Action,
	{_State59, TraitMulTypeType}:                        _GotoState122Action,
	{_State59, TraitAddTypeType}:                        _GotoState121Action,
	{_State59, FieldDeclType}:                           _GotoState117Action,
	{_State59, ImplicitStructTypeType}:                  _GotoState119Action,
	{_State59, ExplicitStructTypeType}:                  _GotoState116Action,
	{_State59, TraitTypeType}:                           _GotoState123Action,
	{_State59, ExplicitEnumTypeType}:                    _GotoState114Action,
	{_State59, FuncTypeType}:                            _GotoState118Action,
	{_State59, ValueTypeType}:                           _GotoState125Action,
	{_State59, ParameterDeclType}:                       _GotoState133Action,
	{_State59, NonEmptyParametersType}:                  _GotoState131Action,
	{_State59, VarargType}:                              _GotoState135Action,
	{_State59, OptionalVarargType}:                      _GotoState132Action,
	{_State59, ParametersType}:                          _GotoState134Action,
	{_State59, ExplicitFuncTypeType}:                    _GotoState115Action,
	{_State62, RparenToken}:                             _GotoState137Action,
	{_State62, CommaToken}:                              _GotoState136Action,
	{_State64, IdentifierToken}:                         _GotoState138Action,
	{_State64, StructToken}:                             _GotoState29Action,
	{_State64, EnumToken}:                               _GotoState108Action,
	{_State64, TraitToken}:                              _GotoState112Action,
	{_State64, FuncToken}:                               _GotoState19Action,
	{_State64, LparenToken}:                             _GotoState64Action,
	{_State64, ExclaimToken}:                            _GotoState109Action,
	{_State64, ExclaimExclaimToken}:                     _GotoState110Action,
	{_State64, BitAndToken}:                             _GotoState107Action,
	{_State64, AtomTypeType}:                            _GotoState113Action,
	{_State64, TraitableTypeType}:                       _GotoState124Action,
	{_State64, TraitMulTypeType}:                        _GotoState122Action,
	{_State64, TraitAddTypeType}:                        _GotoState121Action,
	{_State64, FieldDeclType}:                           _GotoState139Action,
	{_State64, FieldDeclsType}:                          _GotoState140Action,
	{_State64, OptionalFieldDeclsType}:                  _GotoState141Action,
	{_State64, ImplicitStructTypeType}:                  _GotoState119Action,
	{_State64, ExplicitStructTypeType}:                  _GotoState116Action,
	{_State64, TraitTypeType}:                           _GotoState123Action,
	{_State64, ExplicitEnumTypeType}:                    _GotoState114Action,
	{_State64, FuncTypeType}:                            _GotoState118Action,
	{_State64, ValueTypeType}:                           _GotoState125Action,
	{_State64, ExplicitFuncTypeType}:                    _GotoState115Action,
	{_State66, IdentifierToken}:                         _GotoState142Action,
	{_State66, StructToken}:                             _GotoState29Action,
	{_State66, EnumToken}:                               _GotoState108Action,
	{_State66, TraitToken}:                              _GotoState112Action,
	{_State66, FuncToken}:                               _GotoState19Action,
	{_State66, LparenToken}:                             _GotoState64Action,
	{_State66, ExclaimToken}:                            _GotoState109Action,
	{_State66, ExclaimExclaimToken}:                     _GotoState110Action,
	{_State66, BitAndToken}:                             _GotoState107Action,
	{_State66, GenericArgumentsType}:                    _GotoState143Action,
	{_State66, OptionalGenericArgumentsType}:            _GotoState144Action,
	{_State66, AtomTypeType}:                            _GotoState113Action,
	{_State66, TraitableTypeType}:                       _GotoState124Action,
	{_State66, TraitMulTypeType}:                        _GotoState122Action,
	{_State66, TraitAddTypeType}:                        _GotoState121Action,
	{_State66, ImplicitStructTypeType}:                  _GotoState119Action,
	{_State66, ExplicitStructTypeType}:                  _GotoState116Action,
	{_State66, TraitTypeType}:                           _GotoState123Action,
	{_State66, ExplicitEnumTypeType}:                    _GotoState114Action,
	{_State66, FuncTypeType}:                            _GotoState118Action,
	{_State66, ValueTypeType}:                           _GotoState145Action,
	{_State66, ExplicitFuncTypeType}:                    _GotoState115Action,
	{_State67, IdentifierToken}:                         _GotoState146Action,
	{_State68, IntegerLiteralToken}:                     _GotoState21Action,
	{_State68, FloatLiteralToken}:                       _GotoState18Action,
	{_State68, RuneLiteralToken}:                        _GotoState27Action,
	{_State68, StringLiteralToken}:                      _GotoState28Action,
	{_State68, IdentifierToken}:                         _GotoState20Action,
	{_State68, TrueToken}:                               _GotoState31Action,
	{_State68, FalseToken}:                              _GotoState17Action,
	{_State68, StructToken}:                             _GotoState29Action,
	{_State68, FuncToken}:                               _GotoState19Action,
	{_State68, LparenToken}:                             _GotoState24Action,
	{_State68, LabelDeclToken}:                          _GotoState22Action,
	{_State68, NotToken}:                                _GotoState26Action,
	{_State68, SubToken}:                                _GotoState30Action,
	{_State68, MulToken}:                                _GotoState25Action,
	{_State68, BitNegToken}:                             _GotoState16Action,
	{_State68, BitAndToken}:                             _GotoState15Action,
	{_State68, LexErrorToken}:                           _GotoState23Action,
	{_State68, LiteralType}:                             _GotoState43Action,
	{_State68, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State68, AnonymousStructExprType}:                 _GotoState36Action,
	{_State68, AtomExprType}:                            _GotoState37Action,
	{_State68, CallExprType}:                            _GotoState39Action,
	{_State68, AccessExprType}:                          _GotoState32Action,
	{_State68, PostfixUnaryExprType}:                    _GotoState47Action,
	{_State68, PrefixUnaryOpType}:                       _GotoState49Action,
	{_State68, PrefixUnaryExprType}:                     _GotoState48Action,
	{_State68, MulExprType}:                             _GotoState44Action,
	{_State68, AddExprType}:                             _GotoState33Action,
	{_State68, CmpExprType}:                             _GotoState40Action,
	{_State68, AndExprType}:                             _GotoState34Action,
	{_State68, OrExprType}:                              _GotoState46Action,
	{_State68, SequenceExprType}:                        _GotoState50Action,
	{_State68, OptionalLabelDeclType}:                   _GotoState45Action,
	{_State68, BlockExprType}:                           _GotoState38Action,
	{_State68, ExpressionType}:                          _GotoState147Action,
	{_State68, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State68, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State70, LparenToken}:                             _GotoState148Action,
	{_State75, IntegerLiteralToken}:                     _GotoState21Action,
	{_State75, FloatLiteralToken}:                       _GotoState18Action,
	{_State75, RuneLiteralToken}:                        _GotoState27Action,
	{_State75, StringLiteralToken}:                      _GotoState28Action,
	{_State75, IdentifierToken}:                         _GotoState20Action,
	{_State75, TrueToken}:                               _GotoState31Action,
	{_State75, FalseToken}:                              _GotoState17Action,
	{_State75, StructToken}:                             _GotoState29Action,
	{_State75, FuncToken}:                               _GotoState19Action,
	{_State75, LparenToken}:                             _GotoState24Action,
	{_State75, LabelDeclToken}:                          _GotoState22Action,
	{_State75, NotToken}:                                _GotoState26Action,
	{_State75, SubToken}:                                _GotoState30Action,
	{_State75, MulToken}:                                _GotoState25Action,
	{_State75, BitNegToken}:                             _GotoState16Action,
	{_State75, BitAndToken}:                             _GotoState15Action,
	{_State75, LexErrorToken}:                           _GotoState23Action,
	{_State75, LiteralType}:                             _GotoState43Action,
	{_State75, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State75, AnonymousStructExprType}:                 _GotoState36Action,
	{_State75, AtomExprType}:                            _GotoState37Action,
	{_State75, CallExprType}:                            _GotoState39Action,
	{_State75, AccessExprType}:                          _GotoState32Action,
	{_State75, PostfixUnaryExprType}:                    _GotoState47Action,
	{_State75, PrefixUnaryOpType}:                       _GotoState49Action,
	{_State75, PrefixUnaryExprType}:                     _GotoState48Action,
	{_State75, MulExprType}:                             _GotoState149Action,
	{_State75, OptionalLabelDeclType}:                   _GotoState101Action,
	{_State75, BlockExprType}:                           _GotoState38Action,
	{_State75, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State75, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State76, IntegerLiteralToken}:                     _GotoState21Action,
	{_State76, FloatLiteralToken}:                       _GotoState18Action,
	{_State76, RuneLiteralToken}:                        _GotoState27Action,
	{_State76, StringLiteralToken}:                      _GotoState28Action,
	{_State76, IdentifierToken}:                         _GotoState20Action,
	{_State76, TrueToken}:                               _GotoState31Action,
	{_State76, FalseToken}:                              _GotoState17Action,
	{_State76, StructToken}:                             _GotoState29Action,
	{_State76, FuncToken}:                               _GotoState19Action,
	{_State76, LparenToken}:                             _GotoState24Action,
	{_State76, LabelDeclToken}:                          _GotoState22Action,
	{_State76, NotToken}:                                _GotoState26Action,
	{_State76, SubToken}:                                _GotoState30Action,
	{_State76, MulToken}:                                _GotoState25Action,
	{_State76, BitNegToken}:                             _GotoState16Action,
	{_State76, BitAndToken}:                             _GotoState15Action,
	{_State76, LexErrorToken}:                           _GotoState23Action,
	{_State76, LiteralType}:                             _GotoState43Action,
	{_State76, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State76, AnonymousStructExprType}:                 _GotoState36Action,
	{_State76, AtomExprType}:                            _GotoState37Action,
	{_State76, CallExprType}:                            _GotoState39Action,
	{_State76, AccessExprType}:                          _GotoState32Action,
	{_State76, PostfixUnaryExprType}:                    _GotoState47Action,
	{_State76, PrefixUnaryOpType}:                       _GotoState49Action,
	{_State76, PrefixUnaryExprType}:                     _GotoState48Action,
	{_State76, MulExprType}:                             _GotoState44Action,
	{_State76, AddExprType}:                             _GotoState33Action,
	{_State76, CmpExprType}:                             _GotoState150Action,
	{_State76, OptionalLabelDeclType}:                   _GotoState101Action,
	{_State76, BlockExprType}:                           _GotoState38Action,
	{_State76, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State76, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State83, IntegerLiteralToken}:                     _GotoState21Action,
	{_State83, FloatLiteralToken}:                       _GotoState18Action,
	{_State83, RuneLiteralToken}:                        _GotoState27Action,
	{_State83, StringLiteralToken}:                      _GotoState28Action,
	{_State83, IdentifierToken}:                         _GotoState20Action,
	{_State83, TrueToken}:                               _GotoState31Action,
	{_State83, FalseToken}:                              _GotoState17Action,
	{_State83, StructToken}:                             _GotoState29Action,
	{_State83, FuncToken}:                               _GotoState19Action,
	{_State83, LparenToken}:                             _GotoState24Action,
	{_State83, LabelDeclToken}:                          _GotoState22Action,
	{_State83, NotToken}:                                _GotoState26Action,
	{_State83, SubToken}:                                _GotoState30Action,
	{_State83, MulToken}:                                _GotoState25Action,
	{_State83, BitNegToken}:                             _GotoState16Action,
	{_State83, BitAndToken}:                             _GotoState15Action,
	{_State83, LexErrorToken}:                           _GotoState23Action,
	{_State83, LiteralType}:                             _GotoState43Action,
	{_State83, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State83, AnonymousStructExprType}:                 _GotoState36Action,
	{_State83, AtomExprType}:                            _GotoState37Action,
	{_State83, CallExprType}:                            _GotoState39Action,
	{_State83, AccessExprType}:                          _GotoState32Action,
	{_State83, PostfixUnaryExprType}:                    _GotoState47Action,
	{_State83, PrefixUnaryOpType}:                       _GotoState49Action,
	{_State83, PrefixUnaryExprType}:                     _GotoState48Action,
	{_State83, MulExprType}:                             _GotoState44Action,
	{_State83, AddExprType}:                             _GotoState151Action,
	{_State83, OptionalLabelDeclType}:                   _GotoState101Action,
	{_State83, BlockExprType}:                           _GotoState38Action,
	{_State83, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State83, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State85, IntegerLiteralToken}:                     _GotoState21Action,
	{_State85, FloatLiteralToken}:                       _GotoState18Action,
	{_State85, RuneLiteralToken}:                        _GotoState27Action,
	{_State85, StringLiteralToken}:                      _GotoState28Action,
	{_State85, IdentifierToken}:                         _GotoState20Action,
	{_State85, TrueToken}:                               _GotoState31Action,
	{_State85, FalseToken}:                              _GotoState17Action,
	{_State85, StructToken}:                             _GotoState29Action,
	{_State85, FuncToken}:                               _GotoState19Action,
	{_State85, LparenToken}:                             _GotoState24Action,
	{_State85, LabelDeclToken}:                          _GotoState22Action,
	{_State85, NotToken}:                                _GotoState26Action,
	{_State85, SubToken}:                                _GotoState30Action,
	{_State85, MulToken}:                                _GotoState25Action,
	{_State85, BitNegToken}:                             _GotoState16Action,
	{_State85, BitAndToken}:                             _GotoState15Action,
	{_State85, LexErrorToken}:                           _GotoState23Action,
	{_State85, LiteralType}:                             _GotoState43Action,
	{_State85, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State85, AnonymousStructExprType}:                 _GotoState36Action,
	{_State85, AtomExprType}:                            _GotoState37Action,
	{_State85, ArgumentType}:                            _GotoState61Action,
	{_State85, ArgumentsType}:                           _GotoState152Action,
	{_State85, CallExprType}:                            _GotoState39Action,
	{_State85, AccessExprType}:                          _GotoState32Action,
	{_State85, PostfixUnaryExprType}:                    _GotoState47Action,
	{_State85, PrefixUnaryOpType}:                       _GotoState49Action,
	{_State85, PrefixUnaryExprType}:                     _GotoState48Action,
	{_State85, MulExprType}:                             _GotoState44Action,
	{_State85, AddExprType}:                             _GotoState33Action,
	{_State85, CmpExprType}:                             _GotoState40Action,
	{_State85, AndExprType}:                             _GotoState34Action,
	{_State85, OrExprType}:                              _GotoState46Action,
	{_State85, SequenceExprType}:                        _GotoState50Action,
	{_State85, OptionalLabelDeclType}:                   _GotoState45Action,
	{_State85, BlockExprType}:                           _GotoState38Action,
	{_State85, ExpressionType}:                          _GotoState63Action,
	{_State85, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State85, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State92, IntegerLiteralToken}:                     _GotoState21Action,
	{_State92, FloatLiteralToken}:                       _GotoState18Action,
	{_State92, RuneLiteralToken}:                        _GotoState27Action,
	{_State92, StringLiteralToken}:                      _GotoState28Action,
	{_State92, IdentifierToken}:                         _GotoState20Action,
	{_State92, TrueToken}:                               _GotoState31Action,
	{_State92, FalseToken}:                              _GotoState17Action,
	{_State92, StructToken}:                             _GotoState29Action,
	{_State92, FuncToken}:                               _GotoState19Action,
	{_State92, LparenToken}:                             _GotoState24Action,
	{_State92, LabelDeclToken}:                          _GotoState22Action,
	{_State92, NotToken}:                                _GotoState26Action,
	{_State92, SubToken}:                                _GotoState30Action,
	{_State92, MulToken}:                                _GotoState25Action,
	{_State92, BitNegToken}:                             _GotoState16Action,
	{_State92, BitAndToken}:                             _GotoState15Action,
	{_State92, LexErrorToken}:                           _GotoState23Action,
	{_State92, LiteralType}:                             _GotoState43Action,
	{_State92, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State92, AnonymousStructExprType}:                 _GotoState36Action,
	{_State92, AtomExprType}:                            _GotoState37Action,
	{_State92, CallExprType}:                            _GotoState39Action,
	{_State92, AccessExprType}:                          _GotoState32Action,
	{_State92, PostfixUnaryExprType}:                    _GotoState47Action,
	{_State92, PrefixUnaryOpType}:                       _GotoState49Action,
	{_State92, PrefixUnaryExprType}:                     _GotoState153Action,
	{_State92, OptionalLabelDeclType}:                   _GotoState101Action,
	{_State92, BlockExprType}:                           _GotoState38Action,
	{_State92, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State92, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State93, IntegerLiteralToken}:                     _GotoState21Action,
	{_State93, FloatLiteralToken}:                       _GotoState18Action,
	{_State93, RuneLiteralToken}:                        _GotoState27Action,
	{_State93, StringLiteralToken}:                      _GotoState28Action,
	{_State93, IdentifierToken}:                         _GotoState20Action,
	{_State93, TrueToken}:                               _GotoState31Action,
	{_State93, FalseToken}:                              _GotoState17Action,
	{_State93, StructToken}:                             _GotoState29Action,
	{_State93, FuncToken}:                               _GotoState19Action,
	{_State93, LparenToken}:                             _GotoState24Action,
	{_State93, LabelDeclToken}:                          _GotoState22Action,
	{_State93, NotToken}:                                _GotoState26Action,
	{_State93, SubToken}:                                _GotoState30Action,
	{_State93, MulToken}:                                _GotoState25Action,
	{_State93, BitNegToken}:                             _GotoState16Action,
	{_State93, BitAndToken}:                             _GotoState15Action,
	{_State93, LexErrorToken}:                           _GotoState23Action,
	{_State93, LiteralType}:                             _GotoState43Action,
	{_State93, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State93, AnonymousStructExprType}:                 _GotoState36Action,
	{_State93, AtomExprType}:                            _GotoState37Action,
	{_State93, CallExprType}:                            _GotoState39Action,
	{_State93, AccessExprType}:                          _GotoState32Action,
	{_State93, PostfixUnaryExprType}:                    _GotoState47Action,
	{_State93, PrefixUnaryOpType}:                       _GotoState49Action,
	{_State93, PrefixUnaryExprType}:                     _GotoState48Action,
	{_State93, MulExprType}:                             _GotoState44Action,
	{_State93, AddExprType}:                             _GotoState33Action,
	{_State93, CmpExprType}:                             _GotoState40Action,
	{_State93, AndExprType}:                             _GotoState34Action,
	{_State93, OrExprType}:                              _GotoState46Action,
	{_State93, SequenceExprType}:                        _GotoState155Action,
	{_State93, OptionalLabelDeclType}:                   _GotoState101Action,
	{_State93, BlockExprType}:                           _GotoState154Action,
	{_State93, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State93, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State94, IntegerLiteralToken}:                     _GotoState21Action,
	{_State94, FloatLiteralToken}:                       _GotoState18Action,
	{_State94, RuneLiteralToken}:                        _GotoState27Action,
	{_State94, StringLiteralToken}:                      _GotoState28Action,
	{_State94, IdentifierToken}:                         _GotoState20Action,
	{_State94, TrueToken}:                               _GotoState31Action,
	{_State94, FalseToken}:                              _GotoState17Action,
	{_State94, StructToken}:                             _GotoState29Action,
	{_State94, FuncToken}:                               _GotoState19Action,
	{_State94, LparenToken}:                             _GotoState24Action,
	{_State94, LabelDeclToken}:                          _GotoState22Action,
	{_State94, NotToken}:                                _GotoState26Action,
	{_State94, SubToken}:                                _GotoState30Action,
	{_State94, MulToken}:                                _GotoState25Action,
	{_State94, BitNegToken}:                             _GotoState16Action,
	{_State94, BitAndToken}:                             _GotoState15Action,
	{_State94, LexErrorToken}:                           _GotoState23Action,
	{_State94, LiteralType}:                             _GotoState43Action,
	{_State94, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State94, AnonymousStructExprType}:                 _GotoState36Action,
	{_State94, AtomExprType}:                            _GotoState37Action,
	{_State94, CallExprType}:                            _GotoState39Action,
	{_State94, AccessExprType}:                          _GotoState32Action,
	{_State94, PostfixUnaryExprType}:                    _GotoState47Action,
	{_State94, PrefixUnaryOpType}:                       _GotoState49Action,
	{_State94, PrefixUnaryExprType}:                     _GotoState48Action,
	{_State94, MulExprType}:                             _GotoState44Action,
	{_State94, AddExprType}:                             _GotoState33Action,
	{_State94, CmpExprType}:                             _GotoState40Action,
	{_State94, AndExprType}:                             _GotoState34Action,
	{_State94, OrExprType}:                              _GotoState46Action,
	{_State94, SequenceExprType}:                        _GotoState156Action,
	{_State94, OptionalLabelDeclType}:                   _GotoState101Action,
	{_State94, BlockExprType}:                           _GotoState38Action,
	{_State94, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State94, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State95, LbraceToken}:                             _GotoState157Action,
	{_State100, IntegerLiteralToken}:                    _GotoState21Action,
	{_State100, FloatLiteralToken}:                      _GotoState18Action,
	{_State100, RuneLiteralToken}:                       _GotoState27Action,
	{_State100, StringLiteralToken}:                     _GotoState28Action,
	{_State100, IdentifierToken}:                        _GotoState20Action,
	{_State100, TrueToken}:                              _GotoState31Action,
	{_State100, FalseToken}:                             _GotoState17Action,
	{_State100, StructToken}:                            _GotoState29Action,
	{_State100, FuncToken}:                              _GotoState19Action,
	{_State100, LparenToken}:                            _GotoState24Action,
	{_State100, LabelDeclToken}:                         _GotoState22Action,
	{_State100, NotToken}:                               _GotoState26Action,
	{_State100, SubToken}:                               _GotoState30Action,
	{_State100, MulToken}:                               _GotoState25Action,
	{_State100, BitNegToken}:                            _GotoState16Action,
	{_State100, BitAndToken}:                            _GotoState15Action,
	{_State100, LexErrorToken}:                          _GotoState23Action,
	{_State100, LiteralType}:                            _GotoState43Action,
	{_State100, AnonymousFuncExprType}:                  _GotoState35Action,
	{_State100, AnonymousStructExprType}:                _GotoState36Action,
	{_State100, AtomExprType}:                           _GotoState37Action,
	{_State100, CallExprType}:                           _GotoState39Action,
	{_State100, AccessExprType}:                         _GotoState32Action,
	{_State100, PostfixUnaryExprType}:                   _GotoState47Action,
	{_State100, PrefixUnaryOpType}:                      _GotoState49Action,
	{_State100, PrefixUnaryExprType}:                    _GotoState48Action,
	{_State100, MulExprType}:                            _GotoState44Action,
	{_State100, AddExprType}:                            _GotoState33Action,
	{_State100, CmpExprType}:                            _GotoState40Action,
	{_State100, AndExprType}:                            _GotoState158Action,
	{_State100, OptionalLabelDeclType}:                  _GotoState101Action,
	{_State100, BlockExprType}:                          _GotoState38Action,
	{_State100, ExplicitStructTypeType}:                 _GotoState42Action,
	{_State100, ExplicitFuncTypeType}:                   _GotoState41Action,
	{_State101, LbraceToken}:                            _GotoState57Action,
	{_State101, BlockBodyType}:                          _GotoState96Action,
	{_State103, PackageStatementsType}:                  _GotoState159Action,
	{_State104, IdentifierToken}:                        _GotoState160Action,
	{_State104, GenericParameterType}:                   _GotoState161Action,
	{_State104, GenericParametersType}:                  _GotoState162Action,
	{_State104, OptionalGenericParametersType}:          _GotoState163Action,
	{_State105, IdentifierToken}:                        _GotoState142Action,
	{_State105, StructToken}:                            _GotoState29Action,
	{_State105, EnumToken}:                              _GotoState108Action,
	{_State105, TraitToken}:                             _GotoState112Action,
	{_State105, FuncToken}:                              _GotoState19Action,
	{_State105, LparenToken}:                            _GotoState64Action,
	{_State105, ExclaimToken}:                           _GotoState109Action,
	{_State105, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State105, BitAndToken}:                            _GotoState107Action,
	{_State105, AtomTypeType}:                           _GotoState113Action,
	{_State105, TraitableTypeType}:                      _GotoState124Action,
	{_State105, TraitMulTypeType}:                       _GotoState122Action,
	{_State105, TraitAddTypeType}:                       _GotoState121Action,
	{_State105, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State105, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State105, TraitTypeType}:                          _GotoState123Action,
	{_State105, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State105, FuncTypeType}:                           _GotoState118Action,
	{_State105, ValueTypeType}:                          _GotoState164Action,
	{_State105, ExplicitFuncTypeType}:                   _GotoState115Action,
	{_State106, IdentifierToken}:                        _GotoState142Action,
	{_State106, StructToken}:                            _GotoState29Action,
	{_State106, EnumToken}:                              _GotoState108Action,
	{_State106, TraitToken}:                             _GotoState112Action,
	{_State106, FuncToken}:                              _GotoState19Action,
	{_State106, LparenToken}:                            _GotoState64Action,
	{_State106, ExclaimToken}:                           _GotoState109Action,
	{_State106, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State106, BitAndToken}:                            _GotoState107Action,
	{_State106, AtomTypeType}:                           _GotoState113Action,
	{_State106, TraitableTypeType}:                      _GotoState124Action,
	{_State106, TraitMulTypeType}:                       _GotoState122Action,
	{_State106, TraitAddTypeType}:                       _GotoState121Action,
	{_State106, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State106, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State106, TraitTypeType}:                          _GotoState123Action,
	{_State106, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State106, FuncTypeType}:                           _GotoState118Action,
	{_State106, ValueTypeType}:                          _GotoState165Action,
	{_State106, ExplicitFuncTypeType}:                   _GotoState115Action,
	{_State107, IdentifierToken}:                        _GotoState142Action,
	{_State107, StructToken}:                            _GotoState29Action,
	{_State107, EnumToken}:                              _GotoState108Action,
	{_State107, TraitToken}:                             _GotoState112Action,
	{_State107, LparenToken}:                            _GotoState64Action,
	{_State107, ExclaimToken}:                           _GotoState109Action,
	{_State107, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State107, AtomTypeType}:                           _GotoState113Action,
	{_State107, TraitableTypeType}:                      _GotoState124Action,
	{_State107, TraitMulTypeType}:                       _GotoState122Action,
	{_State107, TraitAddTypeType}:                       _GotoState166Action,
	{_State107, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State107, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State107, TraitTypeType}:                          _GotoState123Action,
	{_State107, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State108, LparenToken}:                            _GotoState167Action,
	{_State109, IdentifierToken}:                        _GotoState142Action,
	{_State109, StructToken}:                            _GotoState29Action,
	{_State109, EnumToken}:                              _GotoState108Action,
	{_State109, TraitToken}:                             _GotoState112Action,
	{_State109, LparenToken}:                            _GotoState64Action,
	{_State109, AtomTypeType}:                           _GotoState168Action,
	{_State109, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State109, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State109, TraitTypeType}:                          _GotoState123Action,
	{_State109, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State110, IdentifierToken}:                        _GotoState142Action,
	{_State110, StructToken}:                            _GotoState29Action,
	{_State110, EnumToken}:                              _GotoState108Action,
	{_State110, TraitToken}:                             _GotoState112Action,
	{_State110, LparenToken}:                            _GotoState64Action,
	{_State110, AtomTypeType}:                           _GotoState169Action,
	{_State110, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State110, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State110, TraitTypeType}:                          _GotoState123Action,
	{_State110, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State111, IdentifierToken}:                        _GotoState142Action,
	{_State111, StructToken}:                            _GotoState29Action,
	{_State111, EnumToken}:                              _GotoState108Action,
	{_State111, TraitToken}:                             _GotoState112Action,
	{_State111, FuncToken}:                              _GotoState19Action,
	{_State111, LparenToken}:                            _GotoState64Action,
	{_State111, QuestionToken}:                          _GotoState170Action,
	{_State111, DollarLbracketToken}:                    _GotoState66Action,
	{_State111, ExclaimToken}:                           _GotoState109Action,
	{_State111, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State111, BitAndToken}:                            _GotoState107Action,
	{_State111, OptionalGenericBindingType}:             _GotoState171Action,
	{_State111, AtomTypeType}:                           _GotoState113Action,
	{_State111, TraitableTypeType}:                      _GotoState124Action,
	{_State111, TraitMulTypeType}:                       _GotoState122Action,
	{_State111, TraitAddTypeType}:                       _GotoState121Action,
	{_State111, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State111, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State111, TraitTypeType}:                          _GotoState123Action,
	{_State111, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State111, FuncTypeType}:                           _GotoState118Action,
	{_State111, ValueTypeType}:                          _GotoState172Action,
	{_State111, ExplicitFuncTypeType}:                   _GotoState115Action,
	{_State112, LparenToken}:                            _GotoState173Action,
	{_State120, RparenToken}:                            _GotoState174Action,
	{_State121, AddToken}:                               _GotoState175Action,
	{_State121, SubToken}:                               _GotoState176Action,
	{_State122, MulToken}:                               _GotoState177Action,
	{_State125, OrToken}:                                _GotoState178Action,
	{_State126, DollarLbracketToken}:                    _GotoState104Action,
	{_State126, OptionalGenericParameterDeclType}:       _GotoState179Action,
	{_State128, IntegerLiteralToken}:                    _GotoState21Action,
	{_State128, FloatLiteralToken}:                      _GotoState18Action,
	{_State128, RuneLiteralToken}:                       _GotoState27Action,
	{_State128, StringLiteralToken}:                     _GotoState28Action,
	{_State128, IdentifierToken}:                        _GotoState20Action,
	{_State128, TrueToken}:                              _GotoState31Action,
	{_State128, FalseToken}:                             _GotoState17Action,
	{_State128, ReturnToken}:                            _GotoState184Action,
	{_State128, BreakToken}:                             _GotoState181Action,
	{_State128, ContinueToken}:                          _GotoState182Action,
	{_State128, UnsafeToken}:                            _GotoState185Action,
	{_State128, StructToken}:                            _GotoState29Action,
	{_State128, FuncToken}:                              _GotoState19Action,
	{_State128, AsyncToken}:                             _GotoState180Action,
	{_State128, RbraceToken}:                            _GotoState183Action,
	{_State128, LparenToken}:                            _GotoState24Action,
	{_State128, LabelDeclToken}:                         _GotoState22Action,
	{_State128, NotToken}:                               _GotoState26Action,
	{_State128, SubToken}:                               _GotoState30Action,
	{_State128, MulToken}:                               _GotoState25Action,
	{_State128, BitNegToken}:                            _GotoState16Action,
	{_State128, BitAndToken}:                            _GotoState15Action,
	{_State128, LexErrorToken}:                          _GotoState23Action,
	{_State128, LiteralType}:                            _GotoState43Action,
	{_State128, AnonymousFuncExprType}:                  _GotoState35Action,
	{_State128, AnonymousStructExprType}:                _GotoState36Action,
	{_State128, AtomExprType}:                           _GotoState37Action,
	{_State128, CallExprType}:                           _GotoState39Action,
	{_State128, AccessExprType}:                         _GotoState186Action,
	{_State128, PostfixUnaryExprType}:                   _GotoState47Action,
	{_State128, PrefixUnaryOpType}:                      _GotoState49Action,
	{_State128, PrefixUnaryExprType}:                    _GotoState48Action,
	{_State128, MulExprType}:                            _GotoState44Action,
	{_State128, AddExprType}:                            _GotoState33Action,
	{_State128, CmpExprType}:                            _GotoState40Action,
	{_State128, AndExprType}:                            _GotoState34Action,
	{_State128, OrExprType}:                             _GotoState46Action,
	{_State128, SequenceExprType}:                       _GotoState50Action,
	{_State128, JumpTypeType}:                           _GotoState189Action,
	{_State128, ExpressionOrImplicitStructType}:         _GotoState188Action,
	{_State128, UnsafeStatementType}:                    _GotoState192Action,
	{_State128, StatementBodyType}:                      _GotoState191Action,
	{_State128, StatementType}:                          _GotoState190Action,
	{_State128, OptionalLabelDeclType}:                  _GotoState45Action,
	{_State128, BlockExprType}:                          _GotoState38Action,
	{_State128, ExpressionType}:                         _GotoState187Action,
	{_State128, ExplicitStructTypeType}:                 _GotoState42Action,
	{_State128, ExplicitFuncTypeType}:                   _GotoState41Action,
	{_State129, IdentifierToken}:                        _GotoState142Action,
	{_State129, StructToken}:                            _GotoState29Action,
	{_State129, EnumToken}:                              _GotoState108Action,
	{_State129, TraitToken}:                             _GotoState112Action,
	{_State129, FuncToken}:                              _GotoState19Action,
	{_State129, LparenToken}:                            _GotoState64Action,
	{_State129, ExclaimToken}:                           _GotoState109Action,
	{_State129, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State129, BitAndToken}:                            _GotoState107Action,
	{_State129, AtomTypeType}:                           _GotoState113Action,
	{_State129, TraitableTypeType}:                      _GotoState124Action,
	{_State129, TraitMulTypeType}:                       _GotoState122Action,
	{_State129, TraitAddTypeType}:                       _GotoState121Action,
	{_State129, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State129, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State129, TraitTypeType}:                          _GotoState123Action,
	{_State129, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State129, FuncTypeType}:                           _GotoState118Action,
	{_State129, ValueTypeType}:                          _GotoState193Action,
	{_State129, ExplicitFuncTypeType}:                   _GotoState115Action,
	{_State130, IdentifierToken}:                        _GotoState142Action,
	{_State130, StructToken}:                            _GotoState29Action,
	{_State130, EnumToken}:                              _GotoState108Action,
	{_State130, TraitToken}:                             _GotoState112Action,
	{_State130, FuncToken}:                              _GotoState19Action,
	{_State130, LparenToken}:                            _GotoState64Action,
	{_State130, QuestionToken}:                          _GotoState170Action,
	{_State130, DollarLbracketToken}:                    _GotoState66Action,
	{_State130, DotdotdotToken}:                         _GotoState194Action,
	{_State130, ExclaimToken}:                           _GotoState109Action,
	{_State130, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State130, BitAndToken}:                            _GotoState107Action,
	{_State130, OptionalGenericBindingType}:             _GotoState171Action,
	{_State130, AtomTypeType}:                           _GotoState113Action,
	{_State130, TraitableTypeType}:                      _GotoState124Action,
	{_State130, TraitMulTypeType}:                       _GotoState122Action,
	{_State130, TraitAddTypeType}:                       _GotoState121Action,
	{_State130, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State130, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State130, TraitTypeType}:                          _GotoState123Action,
	{_State130, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State130, FuncTypeType}:                           _GotoState118Action,
	{_State130, ValueTypeType}:                          _GotoState172Action,
	{_State130, ExplicitFuncTypeType}:                   _GotoState115Action,
	{_State131, CommaToken}:                             _GotoState195Action,
	{_State134, RparenToken}:                            _GotoState196Action,
	{_State135, CommaToken}:                             _GotoState197Action,
	{_State136, IntegerLiteralToken}:                    _GotoState21Action,
	{_State136, FloatLiteralToken}:                      _GotoState18Action,
	{_State136, RuneLiteralToken}:                       _GotoState27Action,
	{_State136, StringLiteralToken}:                     _GotoState28Action,
	{_State136, IdentifierToken}:                        _GotoState20Action,
	{_State136, TrueToken}:                              _GotoState31Action,
	{_State136, FalseToken}:                             _GotoState17Action,
	{_State136, StructToken}:                            _GotoState29Action,
	{_State136, FuncToken}:                              _GotoState19Action,
	{_State136, LparenToken}:                            _GotoState24Action,
	{_State136, LabelDeclToken}:                         _GotoState22Action,
	{_State136, NotToken}:                               _GotoState26Action,
	{_State136, SubToken}:                               _GotoState30Action,
	{_State136, MulToken}:                               _GotoState25Action,
	{_State136, BitNegToken}:                            _GotoState16Action,
	{_State136, BitAndToken}:                            _GotoState15Action,
	{_State136, LexErrorToken}:                          _GotoState23Action,
	{_State136, LiteralType}:                            _GotoState43Action,
	{_State136, AnonymousFuncExprType}:                  _GotoState35Action,
	{_State136, AnonymousStructExprType}:                _GotoState36Action,
	{_State136, AtomExprType}:                           _GotoState37Action,
	{_State136, ArgumentType}:                           _GotoState198Action,
	{_State136, CallExprType}:                           _GotoState39Action,
	{_State136, AccessExprType}:                         _GotoState32Action,
	{_State136, PostfixUnaryExprType}:                   _GotoState47Action,
	{_State136, PrefixUnaryOpType}:                      _GotoState49Action,
	{_State136, PrefixUnaryExprType}:                    _GotoState48Action,
	{_State136, MulExprType}:                            _GotoState44Action,
	{_State136, AddExprType}:                            _GotoState33Action,
	{_State136, CmpExprType}:                            _GotoState40Action,
	{_State136, AndExprType}:                            _GotoState34Action,
	{_State136, OrExprType}:                             _GotoState46Action,
	{_State136, SequenceExprType}:                       _GotoState50Action,
	{_State136, OptionalLabelDeclType}:                  _GotoState45Action,
	{_State136, BlockExprType}:                          _GotoState38Action,
	{_State136, ExpressionType}:                         _GotoState63Action,
	{_State136, ExplicitStructTypeType}:                 _GotoState42Action,
	{_State136, ExplicitFuncTypeType}:                   _GotoState41Action,
	{_State138, IdentifierToken}:                        _GotoState142Action,
	{_State138, StructToken}:                            _GotoState29Action,
	{_State138, EnumToken}:                              _GotoState108Action,
	{_State138, TraitToken}:                             _GotoState112Action,
	{_State138, FuncToken}:                              _GotoState19Action,
	{_State138, LparenToken}:                            _GotoState64Action,
	{_State138, DollarLbracketToken}:                    _GotoState66Action,
	{_State138, ExclaimToken}:                           _GotoState109Action,
	{_State138, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State138, BitAndToken}:                            _GotoState107Action,
	{_State138, OptionalGenericBindingType}:             _GotoState171Action,
	{_State138, AtomTypeType}:                           _GotoState113Action,
	{_State138, TraitableTypeType}:                      _GotoState124Action,
	{_State138, TraitMulTypeType}:                       _GotoState122Action,
	{_State138, TraitAddTypeType}:                       _GotoState121Action,
	{_State138, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State138, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State138, TraitTypeType}:                          _GotoState123Action,
	{_State138, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State138, FuncTypeType}:                           _GotoState118Action,
	{_State138, ValueTypeType}:                          _GotoState172Action,
	{_State138, ExplicitFuncTypeType}:                   _GotoState115Action,
	{_State140, NewlinesToken}:                          _GotoState200Action,
	{_State140, CommaToken}:                             _GotoState199Action,
	{_State141, RparenToken}:                            _GotoState201Action,
	{_State142, DollarLbracketToken}:                    _GotoState66Action,
	{_State142, OptionalGenericBindingType}:             _GotoState171Action,
	{_State143, CommaToken}:                             _GotoState202Action,
	{_State144, RbracketToken}:                          _GotoState203Action,
	{_State145, OrToken}:                                _GotoState178Action,
	{_State147, RbracketToken}:                          _GotoState204Action,
	{_State148, IntegerLiteralToken}:                    _GotoState21Action,
	{_State148, FloatLiteralToken}:                      _GotoState18Action,
	{_State148, RuneLiteralToken}:                       _GotoState27Action,
	{_State148, StringLiteralToken}:                     _GotoState28Action,
	{_State148, IdentifierToken}:                        _GotoState20Action,
	{_State148, TrueToken}:                              _GotoState31Action,
	{_State148, FalseToken}:                             _GotoState17Action,
	{_State148, StructToken}:                            _GotoState29Action,
	{_State148, FuncToken}:                              _GotoState19Action,
	{_State148, LparenToken}:                            _GotoState24Action,
	{_State148, LabelDeclToken}:                         _GotoState22Action,
	{_State148, NotToken}:                               _GotoState26Action,
	{_State148, SubToken}:                               _GotoState30Action,
	{_State148, MulToken}:                               _GotoState25Action,
	{_State148, BitNegToken}:                            _GotoState16Action,
	{_State148, BitAndToken}:                            _GotoState15Action,
	{_State148, LexErrorToken}:                          _GotoState23Action,
	{_State148, LiteralType}:                            _GotoState43Action,
	{_State148, AnonymousFuncExprType}:                  _GotoState35Action,
	{_State148, AnonymousStructExprType}:                _GotoState36Action,
	{_State148, AtomExprType}:                           _GotoState37Action,
	{_State148, ArgumentType}:                           _GotoState61Action,
	{_State148, ArgumentsType}:                          _GotoState205Action,
	{_State148, OptionalArgumentsType}:                  _GotoState206Action,
	{_State148, CallExprType}:                           _GotoState39Action,
	{_State148, AccessExprType}:                         _GotoState32Action,
	{_State148, PostfixUnaryExprType}:                   _GotoState47Action,
	{_State148, PrefixUnaryOpType}:                      _GotoState49Action,
	{_State148, PrefixUnaryExprType}:                    _GotoState48Action,
	{_State148, MulExprType}:                            _GotoState44Action,
	{_State148, AddExprType}:                            _GotoState33Action,
	{_State148, CmpExprType}:                            _GotoState40Action,
	{_State148, AndExprType}:                            _GotoState34Action,
	{_State148, OrExprType}:                             _GotoState46Action,
	{_State148, SequenceExprType}:                       _GotoState50Action,
	{_State148, OptionalLabelDeclType}:                  _GotoState45Action,
	{_State148, BlockExprType}:                          _GotoState38Action,
	{_State148, ExpressionType}:                         _GotoState63Action,
	{_State148, ExplicitStructTypeType}:                 _GotoState42Action,
	{_State148, ExplicitFuncTypeType}:                   _GotoState41Action,
	{_State149, MulToken}:                               _GotoState91Action,
	{_State149, DivToken}:                               _GotoState89Action,
	{_State149, ModToken}:                               _GotoState90Action,
	{_State149, BitAndToken}:                            _GotoState86Action,
	{_State149, BitLshiftToken}:                         _GotoState87Action,
	{_State149, BitRshiftToken}:                         _GotoState88Action,
	{_State149, MulOpType}:                              _GotoState92Action,
	{_State150, EqualToken}:                             _GotoState77Action,
	{_State150, NotEqualToken}:                          _GotoState82Action,
	{_State150, LessToken}:                              _GotoState80Action,
	{_State150, LessOrEqualToken}:                       _GotoState81Action,
	{_State150, GreaterToken}:                           _GotoState78Action,
	{_State150, GreaterOrEqualToken}:                    _GotoState79Action,
	{_State150, CmpOpType}:                              _GotoState83Action,
	{_State151, AddToken}:                               _GotoState71Action,
	{_State151, SubToken}:                               _GotoState74Action,
	{_State151, BitXorToken}:                            _GotoState73Action,
	{_State151, BitOrToken}:                             _GotoState72Action,
	{_State151, AddOpType}:                              _GotoState75Action,
	{_State152, RparenToken}:                            _GotoState207Action,
	{_State152, CommaToken}:                             _GotoState136Action,
	{_State155, LabelDeclToken}:                         _GotoState22Action,
	{_State155, OptionalLabelDeclType}:                  _GotoState101Action,
	{_State155, BlockExprType}:                          _GotoState208Action,
	{_State156, LbraceToken}:                            _GotoState57Action,
	{_State156, BlockBodyType}:                          _GotoState209Action,
	{_State157, CaseToken}:                              _GotoState210Action,
	{_State158, AndToken}:                               _GotoState76Action,
	{_State159, UnsafeToken}:                            _GotoState185Action,
	{_State159, RparenToken}:                            _GotoState211Action,
	{_State159, UnsafeStatementType}:                    _GotoState214Action,
	{_State159, PackageStatementBodyType}:               _GotoState213Action,
	{_State159, PackageStatementType}:                   _GotoState212Action,
	{_State160, IdentifierToken}:                        _GotoState142Action,
	{_State160, StructToken}:                            _GotoState29Action,
	{_State160, EnumToken}:                              _GotoState108Action,
	{_State160, TraitToken}:                             _GotoState112Action,
	{_State160, FuncToken}:                              _GotoState19Action,
	{_State160, LparenToken}:                            _GotoState64Action,
	{_State160, ExclaimToken}:                           _GotoState109Action,
	{_State160, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State160, BitAndToken}:                            _GotoState107Action,
	{_State160, AtomTypeType}:                           _GotoState113Action,
	{_State160, TraitableTypeType}:                      _GotoState124Action,
	{_State160, TraitMulTypeType}:                       _GotoState122Action,
	{_State160, TraitAddTypeType}:                       _GotoState121Action,
	{_State160, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State160, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State160, TraitTypeType}:                          _GotoState123Action,
	{_State160, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State160, FuncTypeType}:                           _GotoState118Action,
	{_State160, ValueTypeType}:                          _GotoState215Action,
	{_State160, ExplicitFuncTypeType}:                   _GotoState115Action,
	{_State162, CommaToken}:                             _GotoState216Action,
	{_State163, RbracketToken}:                          _GotoState217Action,
	{_State164, OrToken}:                                _GotoState178Action,
	{_State165, OrToken}:                                _GotoState178Action,
	{_State166, AddToken}:                               _GotoState175Action,
	{_State166, SubToken}:                               _GotoState176Action,
	{_State167, RparenToken}:                            _GotoState218Action,
	{_State172, OrToken}:                                _GotoState178Action,
	{_State173, RparenToken}:                            _GotoState219Action,
	{_State175, IdentifierToken}:                        _GotoState142Action,
	{_State175, StructToken}:                            _GotoState29Action,
	{_State175, EnumToken}:                              _GotoState108Action,
	{_State175, TraitToken}:                             _GotoState112Action,
	{_State175, LparenToken}:                            _GotoState64Action,
	{_State175, ExclaimToken}:                           _GotoState109Action,
	{_State175, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State175, AtomTypeType}:                           _GotoState113Action,
	{_State175, TraitableTypeType}:                      _GotoState124Action,
	{_State175, TraitMulTypeType}:                       _GotoState220Action,
	{_State175, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State175, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State175, TraitTypeType}:                          _GotoState123Action,
	{_State175, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State176, IdentifierToken}:                        _GotoState142Action,
	{_State176, StructToken}:                            _GotoState29Action,
	{_State176, EnumToken}:                              _GotoState108Action,
	{_State176, TraitToken}:                             _GotoState112Action,
	{_State176, LparenToken}:                            _GotoState64Action,
	{_State176, ExclaimToken}:                           _GotoState109Action,
	{_State176, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State176, AtomTypeType}:                           _GotoState113Action,
	{_State176, TraitableTypeType}:                      _GotoState124Action,
	{_State176, TraitMulTypeType}:                       _GotoState221Action,
	{_State176, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State176, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State176, TraitTypeType}:                          _GotoState123Action,
	{_State176, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State177, IdentifierToken}:                        _GotoState142Action,
	{_State177, StructToken}:                            _GotoState29Action,
	{_State177, EnumToken}:                              _GotoState108Action,
	{_State177, TraitToken}:                             _GotoState112Action,
	{_State177, LparenToken}:                            _GotoState64Action,
	{_State177, ExclaimToken}:                           _GotoState109Action,
	{_State177, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State177, AtomTypeType}:                           _GotoState113Action,
	{_State177, TraitableTypeType}:                      _GotoState222Action,
	{_State177, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State177, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State177, TraitTypeType}:                          _GotoState123Action,
	{_State177, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State178, IdentifierToken}:                        _GotoState142Action,
	{_State178, StructToken}:                            _GotoState29Action,
	{_State178, EnumToken}:                              _GotoState108Action,
	{_State178, TraitToken}:                             _GotoState112Action,
	{_State178, FuncToken}:                              _GotoState19Action,
	{_State178, LparenToken}:                            _GotoState64Action,
	{_State178, ExclaimToken}:                           _GotoState109Action,
	{_State178, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State178, BitAndToken}:                            _GotoState107Action,
	{_State178, AtomTypeType}:                           _GotoState113Action,
	{_State178, TraitableTypeType}:                      _GotoState124Action,
	{_State178, TraitMulTypeType}:                       _GotoState122Action,
	{_State178, TraitAddTypeType}:                       _GotoState121Action,
	{_State178, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State178, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State178, TraitTypeType}:                          _GotoState123Action,
	{_State178, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State178, FuncTypeType}:                           _GotoState223Action,
	{_State178, ExplicitFuncTypeType}:                   _GotoState115Action,
	{_State179, LparenToken}:                            _GotoState59Action,
	{_State179, ImplicitFuncTypeType}:                   _GotoState224Action,
	{_State180, IntegerLiteralToken}:                    _GotoState21Action,
	{_State180, FloatLiteralToken}:                      _GotoState18Action,
	{_State180, RuneLiteralToken}:                       _GotoState27Action,
	{_State180, StringLiteralToken}:                     _GotoState28Action,
	{_State180, IdentifierToken}:                        _GotoState20Action,
	{_State180, TrueToken}:                              _GotoState31Action,
	{_State180, FalseToken}:                             _GotoState17Action,
	{_State180, StructToken}:                            _GotoState29Action,
	{_State180, FuncToken}:                              _GotoState19Action,
	{_State180, LparenToken}:                            _GotoState24Action,
	{_State180, LabelDeclToken}:                         _GotoState22Action,
	{_State180, LexErrorToken}:                          _GotoState23Action,
	{_State180, LiteralType}:                            _GotoState43Action,
	{_State180, AnonymousFuncExprType}:                  _GotoState35Action,
	{_State180, AnonymousStructExprType}:                _GotoState36Action,
	{_State180, AtomExprType}:                           _GotoState37Action,
	{_State180, CallExprType}:                           _GotoState226Action,
	{_State180, AccessExprType}:                         _GotoState225Action,
	{_State180, OptionalLabelDeclType}:                  _GotoState101Action,
	{_State180, BlockExprType}:                          _GotoState38Action,
	{_State180, ExplicitStructTypeType}:                 _GotoState42Action,
	{_State180, ExplicitFuncTypeType}:                   _GotoState41Action,
	{_State185, LessToken}:                              _GotoState227Action,
	{_State186, LbracketToken}:                          _GotoState68Action,
	{_State186, DotToken}:                               _GotoState67Action,
	{_State186, QuestionToken}:                          _GotoState69Action,
	{_State186, DollarLbracketToken}:                    _GotoState66Action,
	{_State186, AddAssignToken}:                         _GotoState228Action,
	{_State186, SubAssignToken}:                         _GotoState239Action,
	{_State186, MulAssignToken}:                         _GotoState238Action,
	{_State186, DivAssignToken}:                         _GotoState236Action,
	{_State186, ModAssignToken}:                         _GotoState237Action,
	{_State186, AddOneAssignToken}:                      _GotoState229Action,
	{_State186, SubOneAssignToken}:                      _GotoState240Action,
	{_State186, BitNegAssignToken}:                      _GotoState232Action,
	{_State186, BitAndAssignToken}:                      _GotoState230Action,
	{_State186, BitOrAssignToken}:                       _GotoState233Action,
	{_State186, BitXorAssignToken}:                      _GotoState235Action,
	{_State186, BitLshiftAssignToken}:                   _GotoState231Action,
	{_State186, BitRshiftAssignToken}:                   _GotoState234Action,
	{_State186, OptionalGenericBindingType}:             _GotoState70Action,
	{_State186, OpOneAssignType}:                        _GotoState242Action,
	{_State186, BinaryOpAssignType}:                     _GotoState241Action,
	{_State188, CommaToken}:                             _GotoState243Action,
	{_State189, JumpLabelToken}:                         _GotoState244Action,
	{_State189, OptionalJumpLabelType}:                  _GotoState245Action,
	{_State191, NewlinesToken}:                          _GotoState246Action,
	{_State191, SemicolonToken}:                         _GotoState247Action,
	{_State193, OrToken}:                                _GotoState178Action,
	{_State194, IdentifierToken}:                        _GotoState142Action,
	{_State194, StructToken}:                            _GotoState29Action,
	{_State194, EnumToken}:                              _GotoState108Action,
	{_State194, TraitToken}:                             _GotoState112Action,
	{_State194, FuncToken}:                              _GotoState19Action,
	{_State194, LparenToken}:                            _GotoState64Action,
	{_State194, QuestionToken}:                          _GotoState248Action,
	{_State194, ExclaimToken}:                           _GotoState109Action,
	{_State194, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State194, BitAndToken}:                            _GotoState107Action,
	{_State194, AtomTypeType}:                           _GotoState113Action,
	{_State194, TraitableTypeType}:                      _GotoState124Action,
	{_State194, TraitMulTypeType}:                       _GotoState122Action,
	{_State194, TraitAddTypeType}:                       _GotoState121Action,
	{_State194, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State194, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State194, TraitTypeType}:                          _GotoState123Action,
	{_State194, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State194, FuncTypeType}:                           _GotoState118Action,
	{_State194, ValueTypeType}:                          _GotoState249Action,
	{_State194, ExplicitFuncTypeType}:                   _GotoState115Action,
	{_State195, IdentifierToken}:                        _GotoState130Action,
	{_State195, StructToken}:                            _GotoState29Action,
	{_State195, EnumToken}:                              _GotoState108Action,
	{_State195, TraitToken}:                             _GotoState112Action,
	{_State195, FuncToken}:                              _GotoState19Action,
	{_State195, LparenToken}:                            _GotoState64Action,
	{_State195, DotdotdotToken}:                         _GotoState129Action,
	{_State195, ExclaimToken}:                           _GotoState109Action,
	{_State195, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State195, BitAndToken}:                            _GotoState107Action,
	{_State195, AtomTypeType}:                           _GotoState113Action,
	{_State195, TraitableTypeType}:                      _GotoState124Action,
	{_State195, TraitMulTypeType}:                       _GotoState122Action,
	{_State195, TraitAddTypeType}:                       _GotoState121Action,
	{_State195, FieldDeclType}:                          _GotoState117Action,
	{_State195, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State195, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State195, TraitTypeType}:                          _GotoState123Action,
	{_State195, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State195, FuncTypeType}:                           _GotoState118Action,
	{_State195, ValueTypeType}:                          _GotoState125Action,
	{_State195, ParameterDeclType}:                      _GotoState251Action,
	{_State195, VarargType}:                             _GotoState135Action,
	{_State195, OptionalVarargType}:                     _GotoState250Action,
	{_State195, ExplicitFuncTypeType}:                   _GotoState115Action,
	{_State196, IdentifierToken}:                        _GotoState142Action,
	{_State196, StructToken}:                            _GotoState29Action,
	{_State196, EnumToken}:                              _GotoState108Action,
	{_State196, TraitToken}:                             _GotoState112Action,
	{_State196, FuncToken}:                              _GotoState19Action,
	{_State196, LparenToken}:                            _GotoState64Action,
	{_State196, QuestionToken}:                          _GotoState252Action,
	{_State196, ExclaimToken}:                           _GotoState109Action,
	{_State196, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State196, BitAndToken}:                            _GotoState107Action,
	{_State196, AtomTypeType}:                           _GotoState113Action,
	{_State196, TraitableTypeType}:                      _GotoState124Action,
	{_State196, TraitMulTypeType}:                       _GotoState122Action,
	{_State196, TraitAddTypeType}:                       _GotoState121Action,
	{_State196, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State196, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State196, TraitTypeType}:                          _GotoState123Action,
	{_State196, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State196, FuncTypeType}:                           _GotoState253Action,
	{_State196, ExplicitFuncTypeType}:                   _GotoState115Action,
	{_State199, IdentifierToken}:                        _GotoState138Action,
	{_State199, StructToken}:                            _GotoState29Action,
	{_State199, EnumToken}:                              _GotoState108Action,
	{_State199, TraitToken}:                             _GotoState112Action,
	{_State199, FuncToken}:                              _GotoState19Action,
	{_State199, LparenToken}:                            _GotoState64Action,
	{_State199, ExclaimToken}:                           _GotoState109Action,
	{_State199, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State199, BitAndToken}:                            _GotoState107Action,
	{_State199, AtomTypeType}:                           _GotoState113Action,
	{_State199, TraitableTypeType}:                      _GotoState124Action,
	{_State199, TraitMulTypeType}:                       _GotoState122Action,
	{_State199, TraitAddTypeType}:                       _GotoState121Action,
	{_State199, FieldDeclType}:                          _GotoState254Action,
	{_State199, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State199, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State199, TraitTypeType}:                          _GotoState123Action,
	{_State199, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State199, FuncTypeType}:                           _GotoState118Action,
	{_State199, ValueTypeType}:                          _GotoState125Action,
	{_State199, ExplicitFuncTypeType}:                   _GotoState115Action,
	{_State200, IdentifierToken}:                        _GotoState138Action,
	{_State200, StructToken}:                            _GotoState29Action,
	{_State200, EnumToken}:                              _GotoState108Action,
	{_State200, TraitToken}:                             _GotoState112Action,
	{_State200, FuncToken}:                              _GotoState19Action,
	{_State200, LparenToken}:                            _GotoState64Action,
	{_State200, ExclaimToken}:                           _GotoState109Action,
	{_State200, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State200, BitAndToken}:                            _GotoState107Action,
	{_State200, AtomTypeType}:                           _GotoState113Action,
	{_State200, TraitableTypeType}:                      _GotoState124Action,
	{_State200, TraitMulTypeType}:                       _GotoState122Action,
	{_State200, TraitAddTypeType}:                       _GotoState121Action,
	{_State200, FieldDeclType}:                          _GotoState255Action,
	{_State200, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State200, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State200, TraitTypeType}:                          _GotoState123Action,
	{_State200, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State200, FuncTypeType}:                           _GotoState118Action,
	{_State200, ValueTypeType}:                          _GotoState125Action,
	{_State200, ExplicitFuncTypeType}:                   _GotoState115Action,
	{_State202, IdentifierToken}:                        _GotoState142Action,
	{_State202, StructToken}:                            _GotoState29Action,
	{_State202, EnumToken}:                              _GotoState108Action,
	{_State202, TraitToken}:                             _GotoState112Action,
	{_State202, FuncToken}:                              _GotoState19Action,
	{_State202, LparenToken}:                            _GotoState64Action,
	{_State202, ExclaimToken}:                           _GotoState109Action,
	{_State202, ExclaimExclaimToken}:                    _GotoState110Action,
	{_State202, BitAndToken}:                            _GotoState107Action,
	{_State202, AtomTypeType}:                           _GotoState113Action,
	{_State202, TraitableTypeType}:                      _GotoState124Action,
	{_State202, TraitMulTypeType}:                       _GotoState122Action,
	{_State202, TraitAddTypeType}:                       _GotoState121Action,
	{_State202, ImplicitStructTypeType}:                 _GotoState119Action,
	{_State202, ExplicitStructTypeType}:                 _GotoState116Action,
	{_State202, TraitTypeType}:                          _GotoState123Action,
	{_State202, ExplicitEnumTypeType}:                   _GotoState114Action,
	{_State202, FuncTypeType}:                           _GotoState118Action,
	{_State202, ValueTypeType}:                          _GotoState256Action,
	{_State202, ExplicitFuncTypeType}:                   _GotoState115Action,
	{_State205, CommaToken}:                             _GotoState136Action,
	{_State206, RparenToken}:                            _GotoState257Action,
	{_State209, ElseToken}:                              _GotoState258Action,
	{_State210, DefaultToken}:                           _GotoState259Action,
	{_State213, NewlinesToken}:                          _GotoState260Action,
	{_State213, SemicolonToken}:                         _GotoState261Action,
	{_State215, OrToken}:                                _GotoState178Action,
	{_State216, IdentifierToken}:                        _GotoState160Action,
	{_State216, GenericParameterType}:                   _GotoState262Action,
	{_State220, MulToken}:                               _GotoState177Action,
	{_State221, MulToken}:                               _GotoState177Action,
	{_State225, LbracketToken}:                          _GotoState68Action,
	{_State225, DotToken}:                               _GotoState67Action,
	{_State225, DollarLbracketToken}:                    _GotoState66Action,
	{_State225, OptionalGenericBindingType}:             _GotoState70Action,
	{_State227, IdentifierToken}:                        _GotoState263Action,
	{_State241, IntegerLiteralToken}:                    _GotoState21Action,
	{_State241, FloatLiteralToken}:                      _GotoState18Action,
	{_State241, RuneLiteralToken}:                       _GotoState27Action,
	{_State241, StringLiteralToken}:                     _GotoState28Action,
	{_State241, IdentifierToken}:                        _GotoState20Action,
	{_State241, TrueToken}:                              _GotoState31Action,
	{_State241, FalseToken}:                             _GotoState17Action,
	{_State241, StructToken}:                            _GotoState29Action,
	{_State241, FuncToken}:                              _GotoState19Action,
	{_State241, LparenToken}:                            _GotoState24Action,
	{_State241, LabelDeclToken}:                         _GotoState22Action,
	{_State241, NotToken}:                               _GotoState26Action,
	{_State241, SubToken}:                               _GotoState30Action,
	{_State241, MulToken}:                               _GotoState25Action,
	{_State241, BitNegToken}:                            _GotoState16Action,
	{_State241, BitAndToken}:                            _GotoState15Action,
	{_State241, LexErrorToken}:                          _GotoState23Action,
	{_State241, LiteralType}:                            _GotoState43Action,
	{_State241, AnonymousFuncExprType}:                  _GotoState35Action,
	{_State241, AnonymousStructExprType}:                _GotoState36Action,
	{_State241, AtomExprType}:                           _GotoState37Action,
	{_State241, CallExprType}:                           _GotoState39Action,
	{_State241, AccessExprType}:                         _GotoState32Action,
	{_State241, PostfixUnaryExprType}:                   _GotoState47Action,
	{_State241, PrefixUnaryOpType}:                      _GotoState49Action,
	{_State241, PrefixUnaryExprType}:                    _GotoState48Action,
	{_State241, MulExprType}:                            _GotoState44Action,
	{_State241, AddExprType}:                            _GotoState33Action,
	{_State241, CmpExprType}:                            _GotoState40Action,
	{_State241, AndExprType}:                            _GotoState34Action,
	{_State241, OrExprType}:                             _GotoState46Action,
	{_State241, SequenceExprType}:                       _GotoState50Action,
	{_State241, OptionalLabelDeclType}:                  _GotoState45Action,
	{_State241, BlockExprType}:                          _GotoState38Action,
	{_State241, ExpressionType}:                         _GotoState264Action,
	{_State241, ExplicitStructTypeType}:                 _GotoState42Action,
	{_State241, ExplicitFuncTypeType}:                   _GotoState41Action,
	{_State243, IntegerLiteralToken}:                    _GotoState21Action,
	{_State243, FloatLiteralToken}:                      _GotoState18Action,
	{_State243, RuneLiteralToken}:                       _GotoState27Action,
	{_State243, StringLiteralToken}:                     _GotoState28Action,
	{_State243, IdentifierToken}:                        _GotoState20Action,
	{_State243, TrueToken}:                              _GotoState31Action,
	{_State243, FalseToken}:                             _GotoState17Action,
	{_State243, StructToken}:                            _GotoState29Action,
	{_State243, FuncToken}:                              _GotoState19Action,
	{_State243, LparenToken}:                            _GotoState24Action,
	{_State243, LabelDeclToken}:                         _GotoState22Action,
	{_State243, NotToken}:                               _GotoState26Action,
	{_State243, SubToken}:                               _GotoState30Action,
	{_State243, MulToken}:                               _GotoState25Action,
	{_State243, BitNegToken}:                            _GotoState16Action,
	{_State243, BitAndToken}:                            _GotoState15Action,
	{_State243, LexErrorToken}:                          _GotoState23Action,
	{_State243, LiteralType}:                            _GotoState43Action,
	{_State243, AnonymousFuncExprType}:                  _GotoState35Action,
	{_State243, AnonymousStructExprType}:                _GotoState36Action,
	{_State243, AtomExprType}:                           _GotoState37Action,
	{_State243, CallExprType}:                           _GotoState39Action,
	{_State243, AccessExprType}:                         _GotoState32Action,
	{_State243, PostfixUnaryExprType}:                   _GotoState47Action,
	{_State243, PrefixUnaryOpType}:                      _GotoState49Action,
	{_State243, PrefixUnaryExprType}:                    _GotoState48Action,
	{_State243, MulExprType}:                            _GotoState44Action,
	{_State243, AddExprType}:                            _GotoState33Action,
	{_State243, CmpExprType}:                            _GotoState40Action,
	{_State243, AndExprType}:                            _GotoState34Action,
	{_State243, OrExprType}:                             _GotoState46Action,
	{_State243, SequenceExprType}:                       _GotoState50Action,
	{_State243, OptionalLabelDeclType}:                  _GotoState45Action,
	{_State243, BlockExprType}:                          _GotoState38Action,
	{_State243, ExpressionType}:                         _GotoState265Action,
	{_State243, ExplicitStructTypeType}:                 _GotoState42Action,
	{_State243, ExplicitFuncTypeType}:                   _GotoState41Action,
	{_State245, IntegerLiteralToken}:                    _GotoState21Action,
	{_State245, FloatLiteralToken}:                      _GotoState18Action,
	{_State245, RuneLiteralToken}:                       _GotoState27Action,
	{_State245, StringLiteralToken}:                     _GotoState28Action,
	{_State245, IdentifierToken}:                        _GotoState20Action,
	{_State245, TrueToken}:                              _GotoState31Action,
	{_State245, FalseToken}:                             _GotoState17Action,
	{_State245, StructToken}:                            _GotoState29Action,
	{_State245, FuncToken}:                              _GotoState19Action,
	{_State245, LparenToken}:                            _GotoState24Action,
	{_State245, LabelDeclToken}:                         _GotoState22Action,
	{_State245, NotToken}:                               _GotoState26Action,
	{_State245, SubToken}:                               _GotoState30Action,
	{_State245, MulToken}:                               _GotoState25Action,
	{_State245, BitNegToken}:                            _GotoState16Action,
	{_State245, BitAndToken}:                            _GotoState15Action,
	{_State245, LexErrorToken}:                          _GotoState23Action,
	{_State245, LiteralType}:                            _GotoState43Action,
	{_State245, AnonymousFuncExprType}:                  _GotoState35Action,
	{_State245, AnonymousStructExprType}:                _GotoState36Action,
	{_State245, AtomExprType}:                           _GotoState37Action,
	{_State245, CallExprType}:                           _GotoState39Action,
	{_State245, AccessExprType}:                         _GotoState32Action,
	{_State245, PostfixUnaryExprType}:                   _GotoState47Action,
	{_State245, PrefixUnaryOpType}:                      _GotoState49Action,
	{_State245, PrefixUnaryExprType}:                    _GotoState48Action,
	{_State245, MulExprType}:                            _GotoState44Action,
	{_State245, AddExprType}:                            _GotoState33Action,
	{_State245, CmpExprType}:                            _GotoState40Action,
	{_State245, AndExprType}:                            _GotoState34Action,
	{_State245, OrExprType}:                             _GotoState46Action,
	{_State245, SequenceExprType}:                       _GotoState50Action,
	{_State245, OptionalExpressionOrImplicitStructType}: _GotoState267Action,
	{_State245, ExpressionOrImplicitStructType}:         _GotoState266Action,
	{_State245, OptionalLabelDeclType}:                  _GotoState45Action,
	{_State245, BlockExprType}:                          _GotoState38Action,
	{_State245, ExpressionType}:                         _GotoState187Action,
	{_State245, ExplicitStructTypeType}:                 _GotoState42Action,
	{_State245, ExplicitFuncTypeType}:                   _GotoState41Action,
	{_State249, OrToken}:                                _GotoState178Action,
	{_State256, OrToken}:                                _GotoState178Action,
	{_State258, IfToken}:                                _GotoState94Action,
	{_State258, LbraceToken}:                            _GotoState57Action,
	{_State258, BlockBodyType}:                          _GotoState268Action,
	{_State258, IfExprType}:                             _GotoState269Action,
	{_State259, RbraceToken}:                            _GotoState270Action,
	{_State263, GreaterToken}:                           _GotoState271Action,
	{_State266, CommaToken}:                             _GotoState243Action,
	{_State271, StringLiteralToken}:                     _GotoState272Action,
	{_State4, _WildcardMarker}:                          _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State13, IdentifierToken}:                         _ReduceNilToOptionalReceiverAction,
	{_State15, _WildcardMarker}:                         _ReduceBitAndToPrefixUnaryOpAction,
	{_State16, _WildcardMarker}:                         _ReduceBitNegToPrefixUnaryOpAction,
	{_State17, _WildcardMarker}:                         _ReduceFalseToLiteralAction,
	{_State18, _WildcardMarker}:                         _ReduceFloatLiteralToLiteralAction,
	{_State20, _WildcardMarker}:                         _ReduceIdentifierToAtomExprAction,
	{_State21, _WildcardMarker}:                         _ReduceIntegerLiteralToLiteralAction,
	{_State22, _WildcardMarker}:                         _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State23, _WildcardMarker}:                         _ReduceLexErrorToAtomExprAction,
	{_State24, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State25, _WildcardMarker}:                         _ReduceMulToPrefixUnaryOpAction,
	{_State26, _WildcardMarker}:                         _ReduceNotToPrefixUnaryOpAction,
	{_State27, _WildcardMarker}:                         _ReduceRuneLiteralToLiteralAction,
	{_State28, _WildcardMarker}:                         _ReduceStringLiteralToLiteralAction,
	{_State30, _WildcardMarker}:                         _ReduceSubToPrefixUnaryOpAction,
	{_State31, _WildcardMarker}:                         _ReduceTrueToLiteralAction,
	{_State32, _WildcardMarker}:                         _ReduceAccessExprToPostfixUnaryExprAction,
	{_State32, LparenToken}:                             _ReduceNilToOptionalGenericBindingAction,
	{_State33, _WildcardMarker}:                         _ReduceAddExprToCmpExprAction,
	{_State34, _WildcardMarker}:                         _ReduceAndExprToOrExprAction,
	{_State35, _WildcardMarker}:                         _ReduceAnonymousFuncExprToAtomExprAction,
	{_State36, _WildcardMarker}:                         _ReduceAnonymousStructExprToAtomExprAction,
	{_State37, _WildcardMarker}:                         _ReduceAtomExprToAccessExprAction,
	{_State38, _WildcardMarker}:                         _ReduceBlockExprToAtomExprAction,
	{_State39, _WildcardMarker}:                         _ReduceCallExprToAccessExprAction,
	{_State40, _WildcardMarker}:                         _ReduceCmpExprToAndExprAction,
	{_State43, _WildcardMarker}:                         _ReduceLiteralToAtomExprAction,
	{_State44, _WildcardMarker}:                         _ReduceMulExprToAddExprAction,
	{_State46, _EndMarker}:                              _ReduceToSequenceExprAction,
	{_State47, _WildcardMarker}:                         _ReducePostfixUnaryExprToPrefixUnaryExprAction,
	{_State48, _WildcardMarker}:                         _ReducePrefixUnaryExprToMulExprAction,
	{_State49, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State50, _EndMarker}:                              _ReduceSequenceExprToExpressionAction,
	{_State51, _EndMarker}:                              _ReduceCommentToLexInternalTokensAction,
	{_State52, _EndMarker}:                              _ReduceSpacesToLexInternalTokensAction,
	{_State53, _EndMarker}:                              _ReduceNoSpecToPackageDeclAction,
	{_State54, _WildcardMarker}:                         _ReduceNilToOptionalGenericParameterDeclAction,
	{_State57, _WildcardMarker}:                         _ReduceEmptyListToStatementsAction,
	{_State58, _EndMarker}:                              _ReduceToFuncDefAction,
	{_State59, RparenToken}:                             _ReduceNilToOptionalVarargAction,
	{_State60, _WildcardMarker}:                         _ReduceToExplicitFuncTypeAction,
	{_State61, _WildcardMarker}:                         _ReduceArgumentToArgumentsAction,
	{_State63, _WildcardMarker}:                         _ReducePositionalToArgumentAction,
	{_State64, RparenToken}:                             _ReduceNilToOptionalFieldDeclsAction,
	{_State65, _WildcardMarker}:                         _ReduceToExplicitStructTypeAction,
	{_State66, RbracketToken}:                           _ReduceNilToOptionalGenericArgumentsAction,
	{_State68, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State69, _WildcardMarker}:                         _ReduceQuestionToPostfixUnaryExprAction,
	{_State71, _WildcardMarker}:                         _ReduceAddToAddOpAction,
	{_State72, _WildcardMarker}:                         _ReduceBitOrToAddOpAction,
	{_State73, _WildcardMarker}:                         _ReduceBitXorToAddOpAction,
	{_State74, _WildcardMarker}:                         _ReduceSubToAddOpAction,
	{_State75, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State76, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State77, _WildcardMarker}:                         _ReduceEqualToCmpOpAction,
	{_State78, _WildcardMarker}:                         _ReduceGreaterToCmpOpAction,
	{_State79, _WildcardMarker}:                         _ReduceGreaterOrEqualToCmpOpAction,
	{_State80, _WildcardMarker}:                         _ReduceLessToCmpOpAction,
	{_State81, _WildcardMarker}:                         _ReduceLessOrEqualToCmpOpAction,
	{_State82, _WildcardMarker}:                         _ReduceNotEqualToCmpOpAction,
	{_State83, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State84, _WildcardMarker}:                         _ReduceToAnonymousFuncExprAction,
	{_State85, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State86, _WildcardMarker}:                         _ReduceBitAndToMulOpAction,
	{_State87, _WildcardMarker}:                         _ReduceBitLshiftToMulOpAction,
	{_State88, _WildcardMarker}:                         _ReduceBitRshiftToMulOpAction,
	{_State89, _WildcardMarker}:                         _ReduceDivToMulOpAction,
	{_State90, _WildcardMarker}:                         _ReduceModToMulOpAction,
	{_State91, _WildcardMarker}:                         _ReduceMulToMulOpAction,
	{_State92, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State93, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State94, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State96, _EndMarker}:                              _ReduceToBlockExprAction,
	{_State97, _EndMarker}:                              _ReduceIfExprToExpressionAction,
	{_State98, _EndMarker}:                              _ReduceLoopExprToExpressionAction,
	{_State99, _EndMarker}:                              _ReduceSwitchExprToExpressionAction,
	{_State100, LbraceToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State102, _WildcardMarker}:                        _ReducePrefixOpToPrefixUnaryExprAction,
	{_State103, _WildcardMarker}:                        _ReduceEmptyListToPackageStatementsAction,
	{_State104, RbracketToken}:                          _ReduceNilToOptionalGenericParametersAction,
	{_State111, _WildcardMarker}:                        _ReduceNilToOptionalGenericBindingAction,
	{_State113, _WildcardMarker}:                        _ReduceAtomTypeToTraitableTypeAction,
	{_State114, _WildcardMarker}:                        _ReduceExplicitEnumTypeToAtomTypeAction,
	{_State115, _WildcardMarker}:                        _ReduceExplicitFuncTypeToFuncTypeAction,
	{_State116, _WildcardMarker}:                        _ReduceExplicitStructTypeToAtomTypeAction,
	{_State117, _WildcardMarker}:                        _ReduceFieldDeclToParameterDeclAction,
	{_State118, _WildcardMarker}:                        _ReduceFuncTypeToValueTypeAction,
	{_State119, _WildcardMarker}:                        _ReduceImplicitStructTypeToAtomTypeAction,
	{_State121, _WildcardMarker}:                        _ReduceTraitAddTypeToFuncTypeAction,
	{_State122, _WildcardMarker}:                        _ReduceTraitMulTypeToTraitAddTypeAction,
	{_State123, _WildcardMarker}:                        _ReduceTraitTypeToAtomTypeAction,
	{_State124, _WildcardMarker}:                        _ReduceTraitableTypeToTraitMulTypeAction,
	{_State125, _WildcardMarker}:                        _ReduceImplicitToFieldDeclAction,
	{_State126, LparenToken}:                            _ReduceNilToOptionalGenericParameterDeclAction,
	{_State127, LbraceToken}:                            _ReduceToExplicitFuncDeclAction,
	{_State128, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State130, _WildcardMarker}:                        _ReduceNilToOptionalGenericBindingAction,
	{_State131, RparenToken}:                            _ReduceNonEmptyParametersToParametersAction,
	{_State132, RparenToken}:                            _ReduceOptionalVarargToParametersAction,
	{_State133, _WildcardMarker}:                        _ReduceParameterDeclToNonEmptyParametersAction,
	{_State135, RparenToken}:                            _ReduceVarargToOptionalVarargAction,
	{_State136, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State137, _WildcardMarker}:                        _ReduceImplicitToAnonymousStructExprAction,
	{_State138, _WildcardMarker}:                        _ReduceNilToOptionalGenericBindingAction,
	{_State139, _WildcardMarker}:                        _ReduceFieldDeclToFieldDeclsAction,
	{_State140, RparenToken}:                            _ReduceFieldDeclsToOptionalFieldDeclsAction,
	{_State142, _WildcardMarker}:                        _ReduceNilToOptionalGenericBindingAction,
	{_State143, RbracketToken}:                          _ReduceGenericArgumentsToOptionalGenericArgumentsAction,
	{_State145, _WildcardMarker}:                        _ReduceValueTypeToGenericArgumentsAction,
	{_State146, _WildcardMarker}:                        _ReduceAccessToAccessExprAction,
	{_State148, RparenToken}:                            _ReduceNilToOptionalArgumentsAction,
	{_State148, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State149, _WildcardMarker}:                        _ReduceOpToAddExprAction,
	{_State150, _WildcardMarker}:                        _ReduceOpToAndExprAction,
	{_State151, _WildcardMarker}:                        _ReduceOpToCmpExprAction,
	{_State153, _WildcardMarker}:                        _ReduceOpToMulExprAction,
	{_State154, _WildcardMarker}:                        _ReduceBlockExprToAtomExprAction,
	{_State154, _EndMarker}:                             _ReduceInfiniteToLoopExprAction,
	{_State155, LbraceToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State158, _WildcardMarker}:                        _ReduceOpToOrExprAction,
	{_State160, _WildcardMarker}:                        _ReduceUnconstrainedToGenericParameterAction,
	{_State161, _WildcardMarker}:                        _ReduceGenericParameterToGenericParametersAction,
	{_State162, RbracketToken}:                          _ReduceGenericParametersToOptionalGenericParametersAction,
	{_State164, _EndMarker}:                             _ReduceAliasToTypeDeclAction,
	{_State165, _EndMarker}:                             _ReduceDefinitionToTypeDeclAction,
	{_State166, _WildcardMarker}:                        _ReduceReferenceToFuncTypeAction,
	{_State168, _WildcardMarker}:                        _ReduceMethodInterfaceToTraitableTypeAction,
	{_State169, _WildcardMarker}:                        _ReduceTraitToTraitableTypeAction,
	{_State170, _WildcardMarker}:                        _ReduceInferredToParameterDeclAction,
	{_State171, _WildcardMarker}:                        _ReduceNamedToAtomTypeAction,
	{_State172, _WildcardMarker}:                        _ReduceExplicitToFieldDeclAction,
	{_State174, IdentifierToken}:                        _ReduceReceiverToOptionalReceiverAction,
	{_State180, LbraceToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State181, _WildcardMarker}:                        _ReduceBreakToJumpTypeAction,
	{_State182, _WildcardMarker}:                        _ReduceContinueToJumpTypeAction,
	{_State183, _EndMarker}:                             _ReduceToBlockBodyAction,
	{_State184, _WildcardMarker}:                        _ReduceReturnToJumpTypeAction,
	{_State186, _WildcardMarker}:                        _ReduceAccessExprToPostfixUnaryExprAction,
	{_State186, LparenToken}:                            _ReduceNilToOptionalGenericBindingAction,
	{_State187, _WildcardMarker}:                        _ReduceExpressionToExpressionOrImplicitStructAction,
	{_State188, _WildcardMarker}:                        _ReduceExpressionOrImplicitStructToStatementBodyAction,
	{_State189, _WildcardMarker}:                        _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State190, _WildcardMarker}:                        _ReduceAddToStatementsAction,
	{_State192, _WildcardMarker}:                        _ReduceUnsafeStatementToStatementBodyAction,
	{_State193, _WildcardMarker}:                        _ReduceImplicitToVarargAction,
	{_State195, RparenToken}:                            _ReduceNilToOptionalVarargAction,
	{_State196, _WildcardMarker}:                        _ReduceImplicitUnitToImplicitFuncTypeAction,
	{_State197, RparenToken}:                            _ReduceVararg2ToOptionalVarargAction,
	{_State198, _WildcardMarker}:                        _ReduceAddToArgumentsAction,
	{_State201, _WildcardMarker}:                        _ReduceToImplicitStructTypeAction,
	{_State203, _WildcardMarker}:                        _ReduceBindingToOptionalGenericBindingAction,
	{_State204, _WildcardMarker}:                        _ReduceIndexToAccessExprAction,
	{_State205, RparenToken}:                            _ReduceArgumentsToOptionalArgumentsAction,
	{_State207, _WildcardMarker}:                        _ReduceExplicitToAnonymousStructExprAction,
	{_State208, _EndMarker}:                             _ReduceWhileToLoopExprAction,
	{_State209, _WildcardMarker}:                        _ReduceNoElseToIfExprAction,
	{_State211, _EndMarker}:                             _ReduceWithSpecToPackageDeclAction,
	{_State212, _WildcardMarker}:                        _ReduceAddToPackageStatementsAction,
	{_State214, _WildcardMarker}:                        _ReduceToPackageStatementBodyAction,
	{_State215, _WildcardMarker}:                        _ReduceConstrainedToGenericParameterAction,
	{_State217, _WildcardMarker}:                        _ReduceGenericToOptionalGenericParameterDeclAction,
	{_State218, _WildcardMarker}:                        _ReduceToExplicitEnumTypeAction,
	{_State219, _WildcardMarker}:                        _ReduceToTraitTypeAction,
	{_State220, _WildcardMarker}:                        _ReduceUnionToTraitAddTypeAction,
	{_State221, _WildcardMarker}:                        _ReduceDifferenceToTraitAddTypeAction,
	{_State222, _WildcardMarker}:                        _ReduceIntersectToTraitMulTypeAction,
	{_State223, _WildcardMarker}:                        _ReduceImplicitEnumToValueTypeAction,
	{_State224, LbraceToken}:                            _ReduceToImplicitFuncDeclAction,
	{_State225, LparenToken}:                            _ReduceNilToOptionalGenericBindingAction,
	{_State226, _WildcardMarker}:                        _ReduceCallExprToAccessExprAction,
	{_State226, NewlinesToken}:                          _ReduceAsyncToStatementBodyAction,
	{_State226, SemicolonToken}:                         _ReduceAsyncToStatementBodyAction,
	{_State228, _WildcardMarker}:                        _ReduceAddAssignToBinaryOpAssignAction,
	{_State229, _WildcardMarker}:                        _ReduceAddOneAssignToOpOneAssignAction,
	{_State230, _WildcardMarker}:                        _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State231, _WildcardMarker}:                        _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State232, _WildcardMarker}:                        _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State233, _WildcardMarker}:                        _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State234, _WildcardMarker}:                        _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State235, _WildcardMarker}:                        _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State236, _WildcardMarker}:                        _ReduceDivAssignToBinaryOpAssignAction,
	{_State237, _WildcardMarker}:                        _ReduceModAssignToBinaryOpAssignAction,
	{_State238, _WildcardMarker}:                        _ReduceMulAssignToBinaryOpAssignAction,
	{_State239, _WildcardMarker}:                        _ReduceSubAssignToBinaryOpAssignAction,
	{_State240, _WildcardMarker}:                        _ReduceSubOneAssignToOpOneAssignAction,
	{_State241, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State242, _WildcardMarker}:                        _ReduceOpOneAssignToStatementBodyAction,
	{_State243, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State244, _WildcardMarker}:                        _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State245, NewlinesToken}:                          _ReduceNilToOptionalExpressionOrImplicitStructAction,
	{_State245, SemicolonToken}:                         _ReduceNilToOptionalExpressionOrImplicitStructAction,
	{_State245, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State246, _WildcardMarker}:                        _ReduceImplicitToStatementAction,
	{_State247, _WildcardMarker}:                        _ReduceExplicitToStatementAction,
	{_State248, _WildcardMarker}:                        _ReduceInferredToVarargAction,
	{_State249, _WildcardMarker}:                        _ReduceExplicitToVarargAction,
	{_State250, RparenToken}:                            _ReduceMixedToParametersAction,
	{_State251, _WildcardMarker}:                        _ReduceAddToNonEmptyParametersAction,
	{_State252, _WildcardMarker}:                        _ReduceInferredToImplicitFuncTypeAction,
	{_State253, _WildcardMarker}:                        _ReduceTypedToImplicitFuncTypeAction,
	{_State254, _WildcardMarker}:                        _ReduceExplicitToFieldDeclsAction,
	{_State255, _WildcardMarker}:                        _ReduceImplicitToFieldDeclsAction,
	{_State256, _WildcardMarker}:                        _ReduceAddToGenericArgumentsAction,
	{_State257, _WildcardMarker}:                        _ReduceToCallExprAction,
	{_State260, _WildcardMarker}:                        _ReduceImplicitToPackageStatementAction,
	{_State261, _WildcardMarker}:                        _ReduceExplicitToPackageStatementAction,
	{_State262, _WildcardMarker}:                        _ReduceAddToGenericParametersAction,
	{_State264, _WildcardMarker}:                        _ReduceBinaryOpAssignToStatementBodyAction,
	{_State265, _WildcardMarker}:                        _ReduceImplicitStructToExpressionOrImplicitStructAction,
	{_State266, _WildcardMarker}:                        _ReduceExpressionOrImplicitStructToOptionalExpressionOrImplicitStructAction,
	{_State267, _WildcardMarker}:                        _ReduceJumpToStatementBodyAction,
	{_State268, _EndMarker}:                             _ReduceIfElseToIfExprAction,
	{_State269, _EndMarker}:                             _ReduceMultiIfElseToIfExprAction,
	{_State270, _EndMarker}:                             _ReduceToSwitchExprAction,
	{_State272, _WildcardMarker}:                        _ReduceToUnsafeStatementAction,
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.package_decl
    Reduce:
      (nil)
    Goto:
      PACKAGE -> State 11
      package_decl -> State 6

  State 2:
    Kernel Items:
      #accept: ^.type_decl
    Reduce:
      (nil)
    Goto:
      TYPE -> State 12
      type_decl -> State 7

  State 3:
    Kernel Items:
      #accept: ^.func_def
    Reduce:
      (nil)
    Goto:
      FUNC -> State 13
      explicit_func_decl -> State 14
      func_def -> State 8

  State 4:
    Kernel Items:
      #accept: ^.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 50
      optional_label_decl -> State 45
      block_expr -> State 38
      expression -> State 9
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 5:
    Kernel Items:
      #accept: ^.lex_internal_tokens
    Reduce:
      (nil)
    Goto:
      SPACES -> State 52
      COMMENT -> State 51
      lex_internal_tokens -> State 10

  State 6:
    Kernel Items:
      #accept: ^ package_decl., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 7:
    Kernel Items:
      #accept: ^ type_decl., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      #accept: ^ func_def., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      #accept: ^ expression., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      #accept: ^ lex_internal_tokens., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 11:
    Kernel Items:
      package_decl: PACKAGE.IDENTIFIER
      package_decl: PACKAGE.IDENTIFIER LPAREN package_statements RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 53

  State 12:
    Kernel Items:
      type_decl: TYPE.IDENTIFIER optional_generic_parameter_decl value_type
      type_decl: TYPE.IDENTIFIER EQUAL value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 54

  State 13:
    Kernel Items:
      explicit_func_decl: FUNC.optional_receiver implicit_func_decl
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      LPAREN -> State 55
      optional_receiver -> State 56

  State 14:
    Kernel Items:
      func_def: explicit_func_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 57
      block_body -> State 58

  State 15:
    Kernel Items:
      prefix_unary_op: BIT_AND., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 16:
    Kernel Items:
      prefix_unary_op: BIT_NEG., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 17:
    Kernel Items:
      literal: FALSE., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 18:
    Kernel Items:
      literal: FLOAT_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 19:
    Kernel Items:
      explicit_func_type: FUNC.implicit_func_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 59
      implicit_func_type -> State 60

  State 20:
    Kernel Items:
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      literal: INTEGER_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 22:
    Kernel Items:
      optional_label_decl: LABEL_DECL., *
    Reduce:
      * -> [optional_label_decl]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      atom_expr: LEX_ERROR., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 24:
    Kernel Items:
      anonymous_struct_expr: LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      argument -> State 61
      arguments -> State 62
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 50
      optional_label_decl -> State 45
      block_expr -> State 38
      expression -> State 63
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 25:
    Kernel Items:
      prefix_unary_op: MUL., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      prefix_unary_op: NOT., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 27:
    Kernel Items:
      literal: RUNE_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 28:
    Kernel Items:
      literal: STRING_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 29:
    Kernel Items:
      explicit_struct_type: STRUCT.implicit_struct_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 64
      implicit_struct_type -> State 65

  State 30:
    Kernel Items:
      prefix_unary_op: SUB., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 31:
    Kernel Items:
      literal: TRUE., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      call_expr: access_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET expression RBRACKET
      postfix_unary_expr: access_expr., *
      postfix_unary_expr: access_expr.QUESTION
    Reduce:
      * -> [postfix_unary_expr]
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 68
      DOT -> State 67
      QUESTION -> State 69
      DOLLAR_LBRACKET -> State 66
      optional_generic_binding -> State 70

  State 33:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 71
      SUB -> State 74
      BIT_XOR -> State 73
      BIT_OR -> State 72
      add_op -> State 75

  State 34:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 76

  State 35:
    Kernel Items:
      atom_expr: anonymous_func_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 36:
    Kernel Items:
      atom_expr: anonymous_struct_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 37:
    Kernel Items:
      access_expr: atom_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 38:
    Kernel Items:
      atom_expr: block_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 39:
    Kernel Items:
      access_expr: call_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 40:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 77
      NOT_EQUAL -> State 82
      LESS -> State 80
      LESS_OR_EQUAL -> State 81
      GREATER -> State 78
      GREATER_OR_EQUAL -> State 79
      cmp_op -> State 83

  State 41:
    Kernel Items:
      anonymous_func_expr: explicit_func_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 57
      block_body -> State 84

  State 42:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 85

  State 43:
    Kernel Items:
      atom_expr: literal., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 44:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 91
      DIV -> State 89
      MOD -> State 90
      BIT_AND -> State 86
      BIT_LSHIFT -> State 87
      BIT_RSHIFT -> State 88
      mul_op -> State 92

  State 45:
    Kernel Items:
      block_expr: optional_label_decl.block_body
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 94
      SWITCH -> State 95
      FOR -> State 93
      LBRACE -> State 57
      block_body -> State 96
      if_expr -> State 97
      switch_expr -> State 99
      loop_expr -> State 98

  State 46:
    Kernel Items:
      or_expr: or_expr.OR and_expr
      sequence_expr: or_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      OR -> State 100

  State 47:
    Kernel Items:
      prefix_unary_expr: postfix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      mul_expr: prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 49:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op.prefix_unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 102
      optional_label_decl -> State 101
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 50:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 51:
    Kernel Items:
      lex_internal_tokens: COMMENT., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 52:
    Kernel Items:
      lex_internal_tokens: SPACES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 53:
    Kernel Items:
      package_decl: PACKAGE IDENTIFIER., $
      package_decl: PACKAGE IDENTIFIER.LPAREN package_statements RPAREN
    Reduce:
      $ -> [package_decl]
    Goto:
      LPAREN -> State 103

  State 54:
    Kernel Items:
      type_decl: TYPE IDENTIFIER.optional_generic_parameter_decl value_type
      type_decl: TYPE IDENTIFIER.EQUAL value_type
    Reduce:
      * -> [optional_generic_parameter_decl]
    Goto:
      DOLLAR_LBRACKET -> State 104
      EQUAL -> State 105
      optional_generic_parameter_decl -> State 106

  State 55:
    Kernel Items:
      optional_receiver: LPAREN.parameter_decl RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 111
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      field_decl -> State 117
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 125
      parameter_decl -> State 120
      explicit_func_type -> State 115

  State 56:
    Kernel Items:
      explicit_func_decl: FUNC optional_receiver.implicit_func_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 126
      implicit_func_decl -> State 127

  State 57:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 128

  State 58:
    Kernel Items:
      func_def: explicit_func_decl block_body., $
    Reduce:
      $ -> [func_def]
    Goto:
      (nil)

  State 59:
    Kernel Items:
      implicit_func_type: LPAREN.parameters RPAREN func_type
      implicit_func_type: LPAREN.parameters RPAREN QUESTION
      implicit_func_type: LPAREN.parameters RPAREN
    Reduce:
      RPAREN -> [optional_vararg]
    Goto:
      IDENTIFIER -> State 130
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      DOTDOTDOT -> State 129
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      field_decl -> State 117
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 125
      parameter_decl -> State 133
      non_empty_parameters -> State 131
      vararg -> State 135
      optional_vararg -> State 132
      parameters -> State 134
      explicit_func_type -> State 115

  State 60:
    Kernel Items:
      explicit_func_type: FUNC implicit_func_type., *
    Reduce:
      * -> [explicit_func_type]
    Goto:
      (nil)

  State 61:
    Kernel Items:
      arguments: argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      anonymous_struct_expr: LPAREN arguments.RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 137
      COMMA -> State 136

  State 63:
    Kernel Items:
      argument: expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 64:
    Kernel Items:
      implicit_struct_type: LPAREN.optional_field_decls RPAREN
    Reduce:
      RPAREN -> [optional_field_decls]
    Goto:
      IDENTIFIER -> State 138
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      field_decl -> State 139
      field_decls -> State 140
      optional_field_decls -> State 141
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 125
      explicit_func_type -> State 115

  State 65:
    Kernel Items:
      explicit_struct_type: STRUCT implicit_struct_type., *
    Reduce:
      * -> [explicit_struct_type]
    Goto:
      (nil)

  State 66:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      generic_arguments -> State 143
      optional_generic_arguments -> State 144
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 145
      explicit_func_type -> State 115

  State 67:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146

  State 68:
    Kernel Items:
      access_expr: access_expr LBRACKET.expression RBRACKET
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 50
      optional_label_decl -> State 45
      block_expr -> State 38
      expression -> State 147
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 69:
    Kernel Items:
      postfix_unary_expr: access_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 70:
    Kernel Items:
      call_expr: access_expr optional_generic_binding.LPAREN optional_arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 148

  State 71:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 72:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 73:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 74:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 75:
    Kernel Items:
      add_expr: add_expr add_op.mul_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 149
      optional_label_decl -> State 101
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 76:
    Kernel Items:
      and_expr: and_expr AND.cmp_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 150
      optional_label_decl -> State 101
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 77:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 78:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 79:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 80:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 81:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 82:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      cmp_expr: cmp_expr cmp_op.add_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 151
      optional_label_decl -> State 101
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 84:
    Kernel Items:
      anonymous_func_expr: explicit_func_type block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 85:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      argument -> State 61
      arguments -> State 152
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 50
      optional_label_decl -> State 45
      block_expr -> State 38
      expression -> State 63
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 86:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 87:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 90:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 91:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 92:
    Kernel Items:
      mul_expr: mul_expr mul_op.prefix_unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 153
      optional_label_decl -> State 101
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 93:
    Kernel Items:
      loop_expr: FOR.block_expr
      loop_expr: FOR.sequence_expr block_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 155
      optional_label_decl -> State 101
      block_expr -> State 154
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 94:
    Kernel Items:
      if_expr: IF.sequence_expr block_body
      if_expr: IF.sequence_expr block_body ELSE block_body
      if_expr: IF.sequence_expr block_body ELSE if_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 156
      optional_label_decl -> State 101
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 95:
    Kernel Items:
      switch_expr: SWITCH.LBRACE CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 157

  State 96:
    Kernel Items:
      block_expr: optional_label_decl block_body., $
    Reduce:
      $ -> [block_expr]
    Goto:
      (nil)

  State 97:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 98:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 99:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 100:
    Kernel Items:
      or_expr: or_expr OR.and_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 158
      optional_label_decl -> State 101
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 101:
    Kernel Items:
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 57
      block_body -> State 96

  State 102:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 103:
    Kernel Items:
      package_decl: PACKAGE IDENTIFIER LPAREN.package_statements RPAREN
    Reduce:
      * -> [package_statements]
    Goto:
      package_statements -> State 159

  State 104:
    Kernel Items:
      optional_generic_parameter_decl: DOLLAR_LBRACKET.optional_generic_parameters RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameters]
    Goto:
      IDENTIFIER -> State 160
      generic_parameter -> State 161
      generic_parameters -> State 162
      optional_generic_parameters -> State 163

  State 105:
    Kernel Items:
      type_decl: TYPE IDENTIFIER EQUAL.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 164
      explicit_func_type -> State 115

  State 106:
    Kernel Items:
      type_decl: TYPE IDENTIFIER optional_generic_parameter_decl.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 165
      explicit_func_type -> State 115

  State 107:
    Kernel Items:
      func_type: BIT_AND.trait_add_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      LPAREN -> State 64
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 166
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114

  State 108:
    Kernel Items:
      explicit_enum_type: ENUM.LPAREN RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 167

  State 109:
    Kernel Items:
      traitable_type: EXCLAIM.atom_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      LPAREN -> State 64
      atom_type -> State 168
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114

  State 110:
    Kernel Items:
      traitable_type: EXCLAIM_EXCLAIM.atom_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      LPAREN -> State 64
      atom_type -> State 169
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114

  State 111:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      field_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.QUESTION
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      QUESTION -> State 170
      DOLLAR_LBRACKET -> State 66
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      optional_generic_binding -> State 171
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 172
      explicit_func_type -> State 115

  State 112:
    Kernel Items:
      trait_type: TRAIT.LPAREN RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 173

  State 113:
    Kernel Items:
      traitable_type: atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 114:
    Kernel Items:
      atom_type: explicit_enum_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      func_type: explicit_func_type., *
    Reduce:
      * -> [func_type]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      atom_type: explicit_struct_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 117:
    Kernel Items:
      parameter_decl: field_decl., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 118:
    Kernel Items:
      value_type: func_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 119:
    Kernel Items:
      atom_type: implicit_struct_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      optional_receiver: LPAREN parameter_decl.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 174

  State 121:
    Kernel Items:
      trait_add_type: trait_add_type.ADD trait_mul_type
      trait_add_type: trait_add_type.SUB trait_mul_type
      func_type: trait_add_type., *
    Reduce:
      * -> [func_type]
    Goto:
      ADD -> State 175
      SUB -> State 176

  State 122:
    Kernel Items:
      trait_mul_type: trait_mul_type.MUL traitable_type
      trait_add_type: trait_mul_type., *
    Reduce:
      * -> [trait_add_type]
    Goto:
      MUL -> State 177

  State 123:
    Kernel Items:
      atom_type: trait_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      trait_mul_type: traitable_type., *
    Reduce:
      * -> [trait_mul_type]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      field_decl: value_type., *
      value_type: value_type.OR func_type
    Reduce:
      * -> [field_decl]
    Goto:
      OR -> State 178

  State 126:
    Kernel Items:
      implicit_func_decl: IDENTIFIER.optional_generic_parameter_decl implicit_func_type
    Reduce:
      LPAREN -> [optional_generic_parameter_decl]
    Goto:
      DOLLAR_LBRACKET -> State 104
      optional_generic_parameter_decl -> State 179

  State 127:
    Kernel Items:
      explicit_func_decl: FUNC optional_receiver implicit_func_decl., LBRACE
    Reduce:
      LBRACE -> [explicit_func_decl]
    Goto:
      (nil)

  State 128:
    Kernel Items:
      statements: statements.statement
      block_body: LBRACE statements.RBRACE
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      RETURN -> State 184
      BREAK -> State 181
      CONTINUE -> State 182
      UNSAFE -> State 185
      STRUCT -> State 29
      FUNC -> State 19
      ASYNC -> State 180
      RBRACE -> State 183
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      call_expr -> State 39
      access_expr -> State 186
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 50
      jump_type -> State 189
      expression_or_implicit_struct -> State 188
      unsafe_statement -> State 192
      statement_body -> State 191
      statement -> State 190
      optional_label_decl -> State 45
      block_expr -> State 38
      expression -> State 187
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 129:
    Kernel Items:
      vararg: DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 193
      explicit_func_type -> State 115

  State 130:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      field_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.QUESTION
      vararg: IDENTIFIER.DOTDOTDOT value_type
      vararg: IDENTIFIER.DOTDOTDOT QUESTION
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      QUESTION -> State 170
      DOLLAR_LBRACKET -> State 66
      DOTDOTDOT -> State 194
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      optional_generic_binding -> State 171
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 172
      explicit_func_type -> State 115

  State 131:
    Kernel Items:
      non_empty_parameters: non_empty_parameters.COMMA parameter_decl
      parameters: non_empty_parameters., RPAREN
      parameters: non_empty_parameters.COMMA optional_vararg
    Reduce:
      RPAREN -> [parameters]
    Goto:
      COMMA -> State 195

  State 132:
    Kernel Items:
      parameters: optional_vararg., RPAREN
    Reduce:
      RPAREN -> [parameters]
    Goto:
      (nil)

  State 133:
    Kernel Items:
      non_empty_parameters: parameter_decl., *
    Reduce:
      * -> [non_empty_parameters]
    Goto:
      (nil)

  State 134:
    Kernel Items:
      implicit_func_type: LPAREN parameters.RPAREN func_type
      implicit_func_type: LPAREN parameters.RPAREN QUESTION
      implicit_func_type: LPAREN parameters.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 196

  State 135:
    Kernel Items:
      optional_vararg: vararg., RPAREN
      optional_vararg: vararg.COMMA
    Reduce:
      RPAREN -> [optional_vararg]
    Goto:
      COMMA -> State 197

  State 136:
    Kernel Items:
      arguments: arguments COMMA.argument
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      argument -> State 198
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 50
      optional_label_decl -> State 45
      block_expr -> State 38
      expression -> State 63
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 137:
    Kernel Items:
      anonymous_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 138:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      field_decl: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      DOLLAR_LBRACKET -> State 66
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      optional_generic_binding -> State 171
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 172
      explicit_func_type -> State 115

  State 139:
    Kernel Items:
      field_decls: field_decl., *
    Reduce:
      * -> [field_decls]
    Goto:
      (nil)

  State 140:
    Kernel Items:
      field_decls: field_decls.NEWLINES field_decl
      field_decls: field_decls.COMMA field_decl
      optional_field_decls: field_decls., RPAREN
    Reduce:
      RPAREN -> [optional_field_decls]
    Goto:
      NEWLINES -> State 200
      COMMA -> State 199

  State 141:
    Kernel Items:
      implicit_struct_type: LPAREN optional_field_decls.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 201

  State 142:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 66
      optional_generic_binding -> State 171

  State 143:
    Kernel Items:
      generic_arguments: generic_arguments.COMMA value_type
      optional_generic_arguments: generic_arguments., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 202

  State 144:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 203

  State 145:
    Kernel Items:
      generic_arguments: value_type., *
      value_type: value_type.OR func_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      OR -> State 178

  State 146:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 147:
    Kernel Items:
      access_expr: access_expr LBRACKET expression.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 204

  State 148:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN.optional_arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_arguments]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      argument -> State 61
      arguments -> State 205
      optional_arguments -> State 206
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 50
      optional_label_decl -> State 45
      block_expr -> State 38
      expression -> State 63
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 149:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 91
      DIV -> State 89
      MOD -> State 90
      BIT_AND -> State 86
      BIT_LSHIFT -> State 87
      BIT_RSHIFT -> State 88
      mul_op -> State 92

  State 150:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 77
      NOT_EQUAL -> State 82
      LESS -> State 80
      LESS_OR_EQUAL -> State 81
      GREATER -> State 78
      GREATER_OR_EQUAL -> State 79
      cmp_op -> State 83

  State 151:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 71
      SUB -> State 74
      BIT_XOR -> State 73
      BIT_OR -> State 72
      add_op -> State 75

  State 152:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type LPAREN arguments.RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 207
      COMMA -> State 136

  State 153:
    Kernel Items:
      mul_expr: mul_expr mul_op prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 154:
    Kernel Items:
      atom_expr: block_expr., *
      loop_expr: FOR block_expr., $
    Reduce:
      * -> [atom_expr]
      $ -> [loop_expr]
    Goto:
      (nil)

  State 155:
    Kernel Items:
      loop_expr: FOR sequence_expr.block_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      LABEL_DECL -> State 22
      optional_label_decl -> State 101
      block_expr -> State 208

  State 156:
    Kernel Items:
      if_expr: IF sequence_expr.block_body
      if_expr: IF sequence_expr.block_body ELSE block_body
      if_expr: IF sequence_expr.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 57
      block_body -> State 209

  State 157:
    Kernel Items:
      switch_expr: SWITCH LBRACE.CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      CASE -> State 210

  State 158:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 76

  State 159:
    Kernel Items:
      package_decl: PACKAGE IDENTIFIER LPAREN package_statements.RPAREN
      package_statements: package_statements.package_statement
    Reduce:
      (nil)
    Goto:
      UNSAFE -> State 185
      RPAREN -> State 211
      unsafe_statement -> State 214
      package_statement_body -> State 213
      package_statement -> State 212

  State 160:
    Kernel Items:
      generic_parameter: IDENTIFIER., *
      generic_parameter: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter]
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 215
      explicit_func_type -> State 115

  State 161:
    Kernel Items:
      generic_parameters: generic_parameter., *
    Reduce:
      * -> [generic_parameters]
    Goto:
      (nil)

  State 162:
    Kernel Items:
      generic_parameters: generic_parameters.COMMA generic_parameter
      optional_generic_parameters: generic_parameters., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameters]
    Goto:
      COMMA -> State 216

  State 163:
    Kernel Items:
      optional_generic_parameter_decl: DOLLAR_LBRACKET optional_generic_parameters.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 217

  State 164:
    Kernel Items:
      value_type: value_type.OR func_type
      type_decl: TYPE IDENTIFIER EQUAL value_type., $
    Reduce:
      $ -> [type_decl]
    Goto:
      OR -> State 178

  State 165:
    Kernel Items:
      value_type: value_type.OR func_type
      type_decl: TYPE IDENTIFIER optional_generic_parameter_decl value_type., $
    Reduce:
      $ -> [type_decl]
    Goto:
      OR -> State 178

  State 166:
    Kernel Items:
      trait_add_type: trait_add_type.ADD trait_mul_type
      trait_add_type: trait_add_type.SUB trait_mul_type
      func_type: BIT_AND trait_add_type., *
    Reduce:
      * -> [func_type]
    Goto:
      ADD -> State 175
      SUB -> State 176

  State 167:
    Kernel Items:
      explicit_enum_type: ENUM LPAREN.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 218

  State 168:
    Kernel Items:
      traitable_type: EXCLAIM atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 169:
    Kernel Items:
      traitable_type: EXCLAIM_EXCLAIM atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 170:
    Kernel Items:
      parameter_decl: IDENTIFIER QUESTION., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 171:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 172:
    Kernel Items:
      field_decl: IDENTIFIER value_type., *
      value_type: value_type.OR func_type
    Reduce:
      * -> [field_decl]
    Goto:
      OR -> State 178

  State 173:
    Kernel Items:
      trait_type: TRAIT LPAREN.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 219

  State 174:
    Kernel Items:
      optional_receiver: LPAREN parameter_decl RPAREN., IDENTIFIER
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      (nil)

  State 175:
    Kernel Items:
      trait_add_type: trait_add_type ADD.trait_mul_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      LPAREN -> State 64
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 220
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114

  State 176:
    Kernel Items:
      trait_add_type: trait_add_type SUB.trait_mul_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      LPAREN -> State 64
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 221
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114

  State 177:
    Kernel Items:
      trait_mul_type: trait_mul_type MUL.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      LPAREN -> State 64
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      atom_type -> State 113
      traitable_type -> State 222
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114

  State 178:
    Kernel Items:
      value_type: value_type OR.func_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 223
      explicit_func_type -> State 115

  State 179:
    Kernel Items:
      implicit_func_decl: IDENTIFIER optional_generic_parameter_decl.implicit_func_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 59
      implicit_func_type -> State 224

  State 180:
    Kernel Items:
      statement_body: ASYNC.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      call_expr -> State 226
      access_expr -> State 225
      optional_label_decl -> State 101
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 181:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 182:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 183:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 184:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 185:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 227

  State 186:
    Kernel Items:
      call_expr: access_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET expression RBRACKET
      postfix_unary_expr: access_expr., *
      postfix_unary_expr: access_expr.QUESTION
      statement_body: access_expr.op_one_assign
      statement_body: access_expr.binary_op_assign expression
    Reduce:
      * -> [postfix_unary_expr]
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 68
      DOT -> State 67
      QUESTION -> State 69
      DOLLAR_LBRACKET -> State 66
      ADD_ASSIGN -> State 228
      SUB_ASSIGN -> State 239
      MUL_ASSIGN -> State 238
      DIV_ASSIGN -> State 236
      MOD_ASSIGN -> State 237
      ADD_ONE_ASSIGN -> State 229
      SUB_ONE_ASSIGN -> State 240
      BIT_NEG_ASSIGN -> State 232
      BIT_AND_ASSIGN -> State 230
      BIT_OR_ASSIGN -> State 233
      BIT_XOR_ASSIGN -> State 235
      BIT_LSHIFT_ASSIGN -> State 231
      BIT_RSHIFT_ASSIGN -> State 234
      optional_generic_binding -> State 70
      op_one_assign -> State 242
      binary_op_assign -> State 241

  State 187:
    Kernel Items:
      expression_or_implicit_struct: expression., *
    Reduce:
      * -> [expression_or_implicit_struct]
    Goto:
      (nil)

  State 188:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct.COMMA expression
      statement_body: expression_or_implicit_struct., *
    Reduce:
      * -> [statement_body]
    Goto:
      COMMA -> State 243

  State 189:
    Kernel Items:
      statement_body: jump_type.optional_jump_label optional_expression_or_implicit_struct
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 244
      optional_jump_label -> State 245

  State 190:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 191:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 246
      SEMICOLON -> State 247

  State 192:
    Kernel Items:
      statement_body: unsafe_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 193:
    Kernel Items:
      value_type: value_type.OR func_type
      vararg: DOTDOTDOT value_type., *
    Reduce:
      * -> [vararg]
    Goto:
      OR -> State 178

  State 194:
    Kernel Items:
      vararg: IDENTIFIER DOTDOTDOT.value_type
      vararg: IDENTIFIER DOTDOTDOT.QUESTION
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      QUESTION -> State 248
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 249
      explicit_func_type -> State 115

  State 195:
    Kernel Items:
      non_empty_parameters: non_empty_parameters COMMA.parameter_decl
      parameters: non_empty_parameters COMMA.optional_vararg
    Reduce:
      RPAREN -> [optional_vararg]
    Goto:
      IDENTIFIER -> State 130
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      DOTDOTDOT -> State 129
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      field_decl -> State 117
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 125
      parameter_decl -> State 251
      vararg -> State 135
      optional_vararg -> State 250
      explicit_func_type -> State 115

  State 196:
    Kernel Items:
      implicit_func_type: LPAREN parameters RPAREN.func_type
      implicit_func_type: LPAREN parameters RPAREN.QUESTION
      implicit_func_type: LPAREN parameters RPAREN., *
    Reduce:
      * -> [implicit_func_type]
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      QUESTION -> State 252
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 253
      explicit_func_type -> State 115

  State 197:
    Kernel Items:
      optional_vararg: vararg COMMA., RPAREN
    Reduce:
      RPAREN -> [optional_vararg]
    Goto:
      (nil)

  State 198:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 199:
    Kernel Items:
      field_decls: field_decls COMMA.field_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 138
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      field_decl -> State 254
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 125
      explicit_func_type -> State 115

  State 200:
    Kernel Items:
      field_decls: field_decls NEWLINES.field_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 138
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      field_decl -> State 255
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 125
      explicit_func_type -> State 115

  State 201:
    Kernel Items:
      implicit_struct_type: LPAREN optional_field_decls RPAREN., *
    Reduce:
      * -> [implicit_struct_type]
    Goto:
      (nil)

  State 202:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 142
      STRUCT -> State 29
      ENUM -> State 108
      TRAIT -> State 112
      FUNC -> State 19
      LPAREN -> State 64
      EXCLAIM -> State 109
      EXCLAIM_EXCLAIM -> State 110
      BIT_AND -> State 107
      atom_type -> State 113
      traitable_type -> State 124
      trait_mul_type -> State 122
      trait_add_type -> State 121
      implicit_struct_type -> State 119
      explicit_struct_type -> State 116
      trait_type -> State 123
      explicit_enum_type -> State 114
      func_type -> State 118
      value_type -> State 256
      explicit_func_type -> State 115

  State 203:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 204:
    Kernel Items:
      access_expr: access_expr LBRACKET expression RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 205:
    Kernel Items:
      arguments: arguments.COMMA argument
      optional_arguments: arguments., RPAREN
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 136

  State 206:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 257

  State 207:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 208:
    Kernel Items:
      loop_expr: FOR sequence_expr block_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 209:
    Kernel Items:
      if_expr: IF sequence_expr block_body., *
      if_expr: IF sequence_expr block_body.ELSE block_body
      if_expr: IF sequence_expr block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 258

  State 210:
    Kernel Items:
      switch_expr: SWITCH LBRACE CASE.DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 259

  State 211:
    Kernel Items:
      package_decl: PACKAGE IDENTIFIER LPAREN package_statements RPAREN., $
    Reduce:
      $ -> [package_decl]
    Goto:
      (nil)

  State 212:
    Kernel Items:
      package_statements: package_statements package_statement., *
    Reduce:
      * -> [package_statements]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      package_statement: package_statement_body.NEWLINES
      package_statement: package_statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 260
      SEMICOLON -> State 261

  State 214:
    Kernel Items:
      package_statement_body: unsafe_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 215:
    Kernel Items:
      generic_parameter: IDENTIFIER value_type., *
      value_type: value_type.OR func_type
    Reduce:
      * -> [generic_parameter]
    Goto:
      OR -> State 178

  State 216:
    Kernel Items:
      generic_parameters: generic_parameters COMMA.generic_parameter
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 160
      generic_parameter -> State 262

  State 217:
    Kernel Items:
      optional_generic_parameter_decl: DOLLAR_LBRACKET optional_generic_parameters RBRACKET., *
    Reduce:
      * -> [optional_generic_parameter_decl]
    Goto:
      (nil)

  State 218:
    Kernel Items:
      explicit_enum_type: ENUM LPAREN RPAREN., *
    Reduce:
      * -> [explicit_enum_type]
    Goto:
      (nil)

  State 219:
    Kernel Items:
      trait_type: TRAIT LPAREN RPAREN., *
    Reduce:
      * -> [trait_type]
    Goto:
      (nil)

  State 220:
    Kernel Items:
      trait_mul_type: trait_mul_type.MUL traitable_type
      trait_add_type: trait_add_type ADD trait_mul_type., *
    Reduce:
      * -> [trait_add_type]
    Goto:
      MUL -> State 177

  State 221:
    Kernel Items:
      trait_mul_type: trait_mul_type.MUL traitable_type
      trait_add_type: trait_add_type SUB trait_mul_type., *
    Reduce:
      * -> [trait_add_type]
    Goto:
      MUL -> State 177

  State 222:
    Kernel Items:
      trait_mul_type: trait_mul_type MUL traitable_type., *
    Reduce:
      * -> [trait_mul_type]
    Goto:
      (nil)

  State 223:
    Kernel Items:
      value_type: value_type OR func_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 224:
    Kernel Items:
      implicit_func_decl: IDENTIFIER optional_generic_parameter_decl implicit_func_type., LBRACE
    Reduce:
      LBRACE -> [implicit_func_decl]
    Goto:
      (nil)

  State 225:
    Kernel Items:
      call_expr: access_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET expression RBRACKET
    Reduce:
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 68
      DOT -> State 67
      DOLLAR_LBRACKET -> State 66
      optional_generic_binding -> State 70

  State 226:
    Kernel Items:
      access_expr: call_expr., *
      statement_body: ASYNC call_expr., NEWLINES
      statement_body: ASYNC call_expr., SEMICOLON
    Reduce:
      * -> [access_expr]
      NEWLINES -> [statement_body]
      SEMICOLON -> [statement_body]
    Goto:
      (nil)

  State 227:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 263

  State 228:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 229:
    Kernel Items:
      op_one_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 230:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 231:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 232:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 233:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 234:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 235:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 236:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 237:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 238:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 239:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      op_one_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      statement_body: access_expr binary_op_assign.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 50
      optional_label_decl -> State 45
      block_expr -> State 38
      expression -> State 264
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 242:
    Kernel Items:
      statement_body: access_expr op_one_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 243:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct COMMA.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 50
      optional_label_decl -> State 45
      block_expr -> State 38
      expression -> State 265
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 244:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 245:
    Kernel Items:
      statement_body: jump_type optional_jump_label.optional_expression_or_implicit_struct
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_expression_or_implicit_struct]
      SEMICOLON -> [optional_expression_or_implicit_struct]
    Goto:
      INTEGER_LITERAL -> State 21
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 27
      STRING_LITERAL -> State 28
      IDENTIFIER -> State 20
      TRUE -> State 31
      FALSE -> State 17
      STRUCT -> State 29
      FUNC -> State 19
      LPAREN -> State 24
      LABEL_DECL -> State 22
      NOT -> State 26
      SUB -> State 30
      MUL -> State 25
      BIT_NEG -> State 16
      BIT_AND -> State 15
      LEX_ERROR -> State 23
      literal -> State 43
      anonymous_func_expr -> State 35
      anonymous_struct_expr -> State 36
      atom_expr -> State 37
      call_expr -> State 39
      access_expr -> State 32
      postfix_unary_expr -> State 47
      prefix_unary_op -> State 49
      prefix_unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 50
      optional_expression_or_implicit_struct -> State 267
      expression_or_implicit_struct -> State 266
      optional_label_decl -> State 45
      block_expr -> State 38
      expression -> State 187
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 246:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 247:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 248:
    Kernel Items:
      vararg: IDENTIFIER DOTDOTDOT QUESTION., *
    Reduce:
      * -> [vararg]
    Goto:
      (nil)

  State 249:
    Kernel Items:
      value_type: value_type.OR func_type
      vararg: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [vararg]
    Goto:
      OR -> State 178

  State 250:
    Kernel Items:
      parameters: non_empty_parameters COMMA optional_vararg., RPAREN
    Reduce:
      RPAREN -> [parameters]
    Goto:
      (nil)

  State 251:
    Kernel Items:
      non_empty_parameters: non_empty_parameters COMMA parameter_decl., *
    Reduce:
      * -> [non_empty_parameters]
    Goto:
      (nil)

  State 252:
    Kernel Items:
      implicit_func_type: LPAREN parameters RPAREN QUESTION., *
    Reduce:
      * -> [implicit_func_type]
    Goto:
      (nil)

  State 253:
    Kernel Items:
      implicit_func_type: LPAREN parameters RPAREN func_type., *
    Reduce:
      * -> [implicit_func_type]
    Goto:
      (nil)

  State 254:
    Kernel Items:
      field_decls: field_decls COMMA field_decl., *
    Reduce:
      * -> [field_decls]
    Goto:
      (nil)

  State 255:
    Kernel Items:
      field_decls: field_decls NEWLINES field_decl., *
    Reduce:
      * -> [field_decls]
    Goto:
      (nil)

  State 256:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
      value_type: value_type.OR func_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      OR -> State 178

  State 257:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 258:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE.block_body
      if_expr: IF sequence_expr block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 94
      LBRACE -> State 57
      block_body -> State 268
      if_expr -> State 269

  State 259:
    Kernel Items:
      switch_expr: SWITCH LBRACE CASE DEFAULT.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 270

  State 260:
    Kernel Items:
      package_statement: package_statement_body NEWLINES., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 261:
    Kernel Items:
      package_statement: package_statement_body SEMICOLON., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 262:
    Kernel Items:
      generic_parameters: generic_parameters COMMA generic_parameter., *
    Reduce:
      * -> [generic_parameters]
    Goto:
      (nil)

  State 263:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 271

  State 264:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 265:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct COMMA expression., *
    Reduce:
      * -> [expression_or_implicit_struct]
    Goto:
      (nil)

  State 266:
    Kernel Items:
      optional_expression_or_implicit_struct: expression_or_implicit_struct., *
      expression_or_implicit_struct: expression_or_implicit_struct.COMMA expression
    Reduce:
      * -> [optional_expression_or_implicit_struct]
    Goto:
      COMMA -> State 243

  State 267:
    Kernel Items:
      statement_body: jump_type optional_jump_label optional_expression_or_implicit_struct., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 268:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 269:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 270:
    Kernel Items:
      switch_expr: SWITCH LBRACE CASE DEFAULT RBRACE., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 271:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 272

  State 272:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

Number of states: 272
Number of shift actions: 1335
Number of reduce actions: 221
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
