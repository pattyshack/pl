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
	// 106:2: literal -> TRUE: ...
	TrueToLiteral(True_ *GenericSymbol) (*GenericSymbol, error)

	// 107:2: literal -> FALSE: ...
	FalseToLiteral(False_ *GenericSymbol) (*GenericSymbol, error)

	// 108:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 109:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 110:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 111:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 116:2: anonymous_struct_expr -> explicit: ...
	ExplicitToAnonymousStructExpr(ExplicitStructType_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 117:2: anonymous_struct_expr -> implicit: ...
	ImplicitToAnonymousStructExpr(Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 123:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 124:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 125:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 126:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 127:2: atom_expr -> anonymous_struct_expr: ...
	AnonymousStructExprToAtomExpr(AnonymousStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 128:2: atom_expr -> LEX_ERROR: ...
	LexErrorToAtomExpr(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 132:2: generic_arguments -> value_type: ...
	ValueTypeToGenericArguments(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 133:2: generic_arguments -> add: ...
	AddToGenericArguments(GenericArguments_ *GenericSymbol, Comma_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 136:2: optional_generic_arguments -> generic_arguments: ...
	GenericArgumentsToOptionalGenericArguments(GenericArguments_ *GenericSymbol) (*GenericSymbol, error)

	// 137:2: optional_generic_arguments -> nil: ...
	NilToOptionalGenericArguments() (*GenericSymbol, error)

	// 140:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ *GenericSymbol, OptionalGenericArguments_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 141:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (*GenericSymbol, error)

	// 145:2: argument -> positional: ...
	PositionalToArgument(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 148:2: arguments -> argument: ...
	ArgumentToArguments(Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 149:2: arguments -> add: ...
	AddToArguments(Arguments_ *GenericSymbol, Comma_ *GenericSymbol, Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 152:2: optional_arguments -> arguments: ...
	ArgumentsToOptionalArguments(Arguments_ *GenericSymbol) (*GenericSymbol, error)

	// 153:2: optional_arguments -> nil: ...
	NilToOptionalArguments() (*GenericSymbol, error)

	// 155:13: call_expr -> ...
	ToCallExpr(AccessExpr_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalArguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 158:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 159:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 160:2: access_expr -> call_expr: ...
	CallExprToAccessExpr(CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 162:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Expression_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 165:2: postfix_unary_expr -> access_expr: ...
	AccessExprToPostfixUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 166:2: postfix_unary_expr -> question: ...
	QuestionToPostfixUnaryExpr(AccessExpr_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 169:2: prefix_unary_op -> NOT: ...
	NotToPrefixUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 170:2: prefix_unary_op -> BIT_NEG: ...
	BitNegToPrefixUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 171:2: prefix_unary_op -> SUB: ...
	SubToPrefixUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 174:2: prefix_unary_op -> MUL: ...
	MulToPrefixUnaryOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 177:2: prefix_unary_op -> BIT_AND: ...
	BitAndToPrefixUnaryOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 180:2: prefix_unary_expr -> postfix_unary_expr: ...
	PostfixUnaryExprToPrefixUnaryExpr(PostfixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 181:2: prefix_unary_expr -> prefix_op: ...
	PrefixOpToPrefixUnaryExpr(PrefixUnaryOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 184:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 185:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 186:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 187:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 188:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 189:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 192:2: mul_expr -> prefix_unary_expr: ...
	PrefixUnaryExprToMulExpr(PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 193:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 196:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 197:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 198:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 199:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 202:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 203:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 206:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 207:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 208:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 209:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 210:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 211:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 214:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 215:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 218:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 219:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 222:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 223:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 225:17: sequence_expr -> ...
	ToSequenceExpr(OrExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 228:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 229:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 232:2: optional_expression_or_implicit_struct -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct(ExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 233:2: optional_expression_or_implicit_struct -> nil: ...
	NilToOptionalExpressionOrImplicitStruct() (*GenericSymbol, error)

	// 236:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 237:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 238:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 241:2: op_one_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToOpOneAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 242:2: op_one_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToOpOneAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 245:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 246:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 247:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 248:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 249:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 250:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 251:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 252:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 253:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 254:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 255:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 258:2: expression_or_implicit_struct -> expression: ...
	ExpressionToExpressionOrImplicitStruct(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 259:2: expression_or_implicit_struct -> implicit_struct: ...
	ImplicitStructToExpressionOrImplicitStruct(ExpressionOrImplicitStruct_ *GenericSymbol, Comma_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 263:20: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ *GenericSymbol, Less_ *GenericSymbol, Identifier_ *GenericSymbol, Greater_ *GenericSymbol, StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 266:2: statement_body -> unsafe_statement: ...
	UnsafeStatementToStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 268:2: statement_body -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToStatementBody(ExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 270:2: statement_body -> async: ...
	AsyncToStatementBody(Async_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 273:2: statement_body -> jump: ...
	JumpToStatementBody(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 286:2: statement_body -> op_one_assign: ...
	OpOneAssignToStatementBody(AccessExpr_ *GenericSymbol, OpOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 287:2: statement_body -> binary_op_assign: ...
	BinaryOpAssignToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 290:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 291:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 294:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 295:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 298:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 299:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (*GenericSymbol, error)

	// 301:14: block_body -> ...
	ToBlockBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 302:14: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 307:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 308:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, BlockBody_2 *GenericSymbol) (*GenericSymbol, error)

	// 309:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 312:15: switch_expr -> ...
	ToSwitchExpr(Switch_ *GenericSymbol, Lbrace_ *GenericSymbol, Case_ *GenericSymbol, Default_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 315:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(For_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 316:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 321:2: expression -> sequence_expr: ...
	SequenceExprToExpression(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 322:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 323:2: expression -> switch_expr: ...
	SwitchExprToExpression(OptionalLabelDecl_ *GenericSymbol, SwitchExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 324:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ *GenericSymbol, LoopExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 327:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 328:2: atom_type -> explicit_struct_type: ...
	ExplicitStructTypeToAtomType(ExplicitStructType_ *GenericSymbol) (*GenericSymbol, error)

	// 329:2: atom_type -> implicit_struct_type: ...
	ImplicitStructTypeToAtomType(ImplicitStructType_ *GenericSymbol) (*GenericSymbol, error)

	// 330:2: atom_type -> implicit_enum_type: ...
	ImplicitEnumTypeToAtomType(ImplicitEnumType_ *GenericSymbol) (*GenericSymbol, error)

	// 331:2: atom_type -> explicit_enum_type: ...
	ExplicitEnumTypeToAtomType(ExplicitEnumType_ *GenericSymbol) (*GenericSymbol, error)

	// 332:2: atom_type -> trait_type: ...
	TraitTypeToAtomType(TraitType_ *GenericSymbol) (*GenericSymbol, error)

	// 335:2: traitable_type -> atom_type: ...
	AtomTypeToTraitableType(AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 336:2: traitable_type -> method_interface: ...
	MethodInterfaceToTraitableType(Exclaim_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 337:2: traitable_type -> trait: ...
	TraitToTraitableType(ExclaimExclaim_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 340:2: trait_mul_type -> traitable_type: ...
	TraitableTypeToTraitMulType(TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 341:2: trait_mul_type -> intersect: ...
	IntersectToTraitMulType(TraitMulType_ *GenericSymbol, Mul_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 344:2: trait_add_type -> trait_mul_type: ...
	TraitMulTypeToTraitAddType(TraitMulType_ *GenericSymbol) (*GenericSymbol, error)

	// 345:2: trait_add_type -> union: ...
	UnionToTraitAddType(TraitAddType_ *GenericSymbol, Add_ *GenericSymbol, TraitMulType_ *GenericSymbol) (*GenericSymbol, error)

	// 350:2: trait_add_type -> difference: ...
	DifferenceToTraitAddType(TraitAddType_ *GenericSymbol, Sub_ *GenericSymbol, TraitMulType_ *GenericSymbol) (*GenericSymbol, error)

	// 355:2: field_decl -> explicit: ...
	ExplicitToFieldDecl(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 356:2: field_decl -> implicit: ...
	ImplicitToFieldDecl(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 359:2: field_decls -> field_decl: ...
	FieldDeclToFieldDecls(FieldDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 360:2: field_decls -> implicit: ...
	ImplicitToFieldDecls(FieldDecls_ *GenericSymbol, Newlines_ *GenericSymbol, FieldDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 361:2: field_decls -> explicit: ...
	ExplicitToFieldDecls(FieldDecls_ *GenericSymbol, Comma_ *GenericSymbol, FieldDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 364:2: optional_field_decls -> field_decls: ...
	FieldDeclsToOptionalFieldDecls(FieldDecls_ *GenericSymbol) (*GenericSymbol, error)

	// 365:2: optional_field_decls -> nil: ...
	NilToOptionalFieldDecls() (*GenericSymbol, error)

	// 367:24: implicit_struct_type -> ...
	ToImplicitStructType(Lparen_ *GenericSymbol, OptionalFieldDecls_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 370:2: field_decl_or_list -> pair: ...
	PairToFieldDeclOrList(FieldDecl_ *GenericSymbol, Or_ *GenericSymbol, FieldDecl_2 *GenericSymbol) (*GenericSymbol, error)

	// 371:2: field_decl_or_list -> add: ...
	AddToFieldDeclOrList(FieldDeclOrList_ *GenericSymbol, Or_ *GenericSymbol, FieldDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 373:22: implicit_enum_type -> ...
	ToImplicitEnumType(Lparen_ *GenericSymbol, FieldDeclOrList_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 376:24: explicit_struct_type -> ...
	ToExplicitStructType(Struct_ *GenericSymbol, ImplicitStructType_ *GenericSymbol) (*GenericSymbol, error)

	// 379:14: trait_type -> ...
	ToTraitType(Trait_ *GenericSymbol, Lparen_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 382:22: explicit_enum_type -> ...
	ToExplicitEnumType(Enum_ *GenericSymbol, Lparen_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 385:2: value_type -> trait_add_type: ...
	TraitAddTypeToValueType(TraitAddType_ *GenericSymbol) (*GenericSymbol, error)

	// 386:2: value_type -> reference: ...
	ReferenceToValueType(BitAnd_ *GenericSymbol, TraitAddType_ *GenericSymbol) (*GenericSymbol, error)

	// 387:2: value_type -> func_type: ...
	FuncTypeToValueType(FuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 390:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 391:2: type_def -> alias: ...
	AliasToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, Equal_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 399:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 400:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 403:2: generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToGenericParameterDefs(GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 404:2: generic_parameter_defs -> add: ...
	AddToGenericParameterDefs(GenericParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 407:2: optional_generic_parameter_defs -> generic_parameter_defs: ...
	GenericParameterDefsToOptionalGenericParameterDefs(GenericParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 408:2: optional_generic_parameter_defs -> nil: ...
	NilToOptionalGenericParameterDefs() (*GenericSymbol, error)

	// 411:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ *GenericSymbol, OptionalGenericParameterDefs_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 412:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (*GenericSymbol, error)

	// 419:2: parameter_type -> value_type: ...
	ValueTypeToParameterType(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 420:2: parameter_type -> QUESTION: ...
	QuestionToParameterType(Question_ *GenericSymbol) (*GenericSymbol, error)

	// 423:2: return_type -> value_type: ...
	ValueTypeToReturnType(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 424:2: return_type -> QUESTION: ...
	QuestionToReturnType(Question_ *GenericSymbol) (*GenericSymbol, error)

	// 425:2: return_type -> nil: ...
	NilToReturnType() (*GenericSymbol, error)

	// 428:2: parameter_decl -> parameter_def: ...
	ParameterDefToParameterDecl(ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 429:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ParameterType_ *GenericSymbol) (*GenericSymbol, error)

	// 430:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(Dotdotdot_ *GenericSymbol, ParameterType_ *GenericSymbol) (*GenericSymbol, error)

	// 433:2: parameter_decls -> parameter_decl: ...
	ParameterDeclToParameterDecls(ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 434:2: parameter_decls -> add: ...
	AddToParameterDecls(ParameterDecls_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 437:2: optional_parameter_decls -> parameter_decls: ...
	ParameterDeclsToOptionalParameterDecls(ParameterDecls_ *GenericSymbol) (*GenericSymbol, error)

	// 438:2: optional_parameter_decls -> nil: ...
	NilToOptionalParameterDecls() (*GenericSymbol, error)

	// 440:13: func_type -> ...
	ToFuncType(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 443:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ *GenericSymbol, ParameterType_ *GenericSymbol) (*GenericSymbol, error)

	// 444:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ParameterType_ *GenericSymbol) (*GenericSymbol, error)

	// 447:2: parameter_defs -> parameter_def: ...
	ParameterDefToParameterDefs(ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 448:2: parameter_defs -> add: ...
	AddToParameterDefs(ParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 451:2: optional_parameter_defs -> parameter_defs: ...
	ParameterDefsToOptionalParameterDefs(ParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 452:2: optional_parameter_defs -> nil: ...
	NilToOptionalParameterDefs() (*GenericSymbol, error)

	// 455:2: optional_receiver -> receiver: ...
	ReceiverToOptionalReceiver(Lparen_ *GenericSymbol, ParameterDef_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 456:2: optional_receiver -> nil: ...
	NilToOptionalReceiver() (*GenericSymbol, error)

	// 459:2: named_func_def -> ...
	ToNamedFuncDef(Func_ *GenericSymbol, OptionalReceiver_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 462:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 469:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 470:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, PackageStatements_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 472:26: package_statement_body -> ...
	ToPackageStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 475:2: package_statement -> implicit: ...
	ImplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 476:2: package_statement -> explicit: ...
	ExplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 479:2: package_statements -> empty_list: ...
	EmptyListToPackageStatements() (*GenericSymbol, error)

	// 480:2: package_statements -> add: ...
	AddToPackageStatements(PackageStatements_ *GenericSymbol, PackageStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 484:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 485:2: lex_internal_tokens -> COMMENT: ...
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

func ParsePackageDef(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParsePackageDefWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParsePackageDefWithCustomErrorHandler(
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

func ParseTypeDef(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParseTypeDefWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseTypeDefWithCustomErrorHandler(
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

func ParseNamedFuncDef(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParseNamedFuncDefWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseNamedFuncDefWithCustomErrorHandler(
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
	case FieldDeclOrListType:
		return "field_decl_or_list"
	case ImplicitEnumTypeType:
		return "implicit_enum_type"
	case ExplicitStructTypeType:
		return "explicit_struct_type"
	case TraitTypeType:
		return "trait_type"
	case ExplicitEnumTypeType:
		return "explicit_enum_type"
	case ValueTypeType:
		return "value_type"
	case TypeDefType:
		return "type_def"
	case GenericParameterDefType:
		return "generic_parameter_def"
	case GenericParameterDefsType:
		return "generic_parameter_defs"
	case OptionalGenericParameterDefsType:
		return "optional_generic_parameter_defs"
	case OptionalGenericParametersType:
		return "optional_generic_parameters"
	case ParameterTypeType:
		return "parameter_type"
	case ReturnTypeType:
		return "return_type"
	case ParameterDeclType:
		return "parameter_decl"
	case ParameterDeclsType:
		return "parameter_decls"
	case OptionalParameterDeclsType:
		return "optional_parameter_decls"
	case FuncTypeType:
		return "func_type"
	case ParameterDefType:
		return "parameter_def"
	case ParameterDefsType:
		return "parameter_defs"
	case OptionalParameterDefsType:
		return "optional_parameter_defs"
	case OptionalReceiverType:
		return "optional_receiver"
	case NamedFuncDefType:
		return "named_func_def"
	case AnonymousFuncExprType:
		return "anonymous_func_expr"
	case PackageDefType:
		return "package_def"
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
	AnonymousStructExprType                = SymbolId(334)
	AtomExprType                           = SymbolId(335)
	GenericArgumentsType                   = SymbolId(336)
	OptionalGenericArgumentsType           = SymbolId(337)
	OptionalGenericBindingType             = SymbolId(338)
	ArgumentType                           = SymbolId(339)
	ArgumentsType                          = SymbolId(340)
	OptionalArgumentsType                  = SymbolId(341)
	CallExprType                           = SymbolId(342)
	AccessExprType                         = SymbolId(343)
	PostfixUnaryExprType                   = SymbolId(344)
	PrefixUnaryOpType                      = SymbolId(345)
	PrefixUnaryExprType                    = SymbolId(346)
	MulOpType                              = SymbolId(347)
	MulExprType                            = SymbolId(348)
	AddOpType                              = SymbolId(349)
	AddExprType                            = SymbolId(350)
	CmpOpType                              = SymbolId(351)
	CmpExprType                            = SymbolId(352)
	AndExprType                            = SymbolId(353)
	OrExprType                             = SymbolId(354)
	SequenceExprType                       = SymbolId(355)
	OptionalJumpLabelType                  = SymbolId(356)
	OptionalExpressionOrImplicitStructType = SymbolId(357)
	JumpTypeType                           = SymbolId(358)
	OpOneAssignType                        = SymbolId(359)
	BinaryOpAssignType                     = SymbolId(360)
	ExpressionOrImplicitStructType         = SymbolId(361)
	UnsafeStatementType                    = SymbolId(362)
	StatementBodyType                      = SymbolId(363)
	StatementType                          = SymbolId(364)
	StatementsType                         = SymbolId(365)
	OptionalLabelDeclType                  = SymbolId(366)
	BlockBodyType                          = SymbolId(367)
	BlockExprType                          = SymbolId(368)
	IfExprType                             = SymbolId(369)
	SwitchExprType                         = SymbolId(370)
	LoopExprType                           = SymbolId(371)
	ExpressionType                         = SymbolId(372)
	AtomTypeType                           = SymbolId(373)
	TraitableTypeType                      = SymbolId(374)
	TraitMulTypeType                       = SymbolId(375)
	TraitAddTypeType                       = SymbolId(376)
	FieldDeclType                          = SymbolId(377)
	FieldDeclsType                         = SymbolId(378)
	OptionalFieldDeclsType                 = SymbolId(379)
	ImplicitStructTypeType                 = SymbolId(380)
	FieldDeclOrListType                    = SymbolId(381)
	ImplicitEnumTypeType                   = SymbolId(382)
	ExplicitStructTypeType                 = SymbolId(383)
	TraitTypeType                          = SymbolId(384)
	ExplicitEnumTypeType                   = SymbolId(385)
	ValueTypeType                          = SymbolId(386)
	TypeDefType                            = SymbolId(387)
	GenericParameterDefType                = SymbolId(388)
	GenericParameterDefsType               = SymbolId(389)
	OptionalGenericParameterDefsType       = SymbolId(390)
	OptionalGenericParametersType          = SymbolId(391)
	ParameterTypeType                      = SymbolId(392)
	ReturnTypeType                         = SymbolId(393)
	ParameterDeclType                      = SymbolId(394)
	ParameterDeclsType                     = SymbolId(395)
	OptionalParameterDeclsType             = SymbolId(396)
	FuncTypeType                           = SymbolId(397)
	ParameterDefType                       = SymbolId(398)
	ParameterDefsType                      = SymbolId(399)
	OptionalParameterDefsType              = SymbolId(400)
	OptionalReceiverType                   = SymbolId(401)
	NamedFuncDefType                       = SymbolId(402)
	AnonymousFuncExprType                  = SymbolId(403)
	PackageDefType                         = SymbolId(404)
	PackageStatementBodyType               = SymbolId(405)
	PackageStatementType                   = SymbolId(406)
	PackageStatementsType                  = SymbolId(407)
	LexInternalTokensType                  = SymbolId(408)
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
	_ReduceExplicitToAnonymousStructExpr                                  = _ReduceType(7)
	_ReduceImplicitToAnonymousStructExpr                                  = _ReduceType(8)
	_ReduceLiteralToAtomExpr                                              = _ReduceType(9)
	_ReduceIdentifierToAtomExpr                                           = _ReduceType(10)
	_ReduceBlockExprToAtomExpr                                            = _ReduceType(11)
	_ReduceAnonymousFuncExprToAtomExpr                                    = _ReduceType(12)
	_ReduceAnonymousStructExprToAtomExpr                                  = _ReduceType(13)
	_ReduceLexErrorToAtomExpr                                             = _ReduceType(14)
	_ReduceValueTypeToGenericArguments                                    = _ReduceType(15)
	_ReduceAddToGenericArguments                                          = _ReduceType(16)
	_ReduceGenericArgumentsToOptionalGenericArguments                     = _ReduceType(17)
	_ReduceNilToOptionalGenericArguments                                  = _ReduceType(18)
	_ReduceBindingToOptionalGenericBinding                                = _ReduceType(19)
	_ReduceNilToOptionalGenericBinding                                    = _ReduceType(20)
	_ReducePositionalToArgument                                           = _ReduceType(21)
	_ReduceArgumentToArguments                                            = _ReduceType(22)
	_ReduceAddToArguments                                                 = _ReduceType(23)
	_ReduceArgumentsToOptionalArguments                                   = _ReduceType(24)
	_ReduceNilToOptionalArguments                                         = _ReduceType(25)
	_ReduceToCallExpr                                                     = _ReduceType(26)
	_ReduceAtomExprToAccessExpr                                           = _ReduceType(27)
	_ReduceAccessToAccessExpr                                             = _ReduceType(28)
	_ReduceCallExprToAccessExpr                                           = _ReduceType(29)
	_ReduceIndexToAccessExpr                                              = _ReduceType(30)
	_ReduceAccessExprToPostfixUnaryExpr                                   = _ReduceType(31)
	_ReduceQuestionToPostfixUnaryExpr                                     = _ReduceType(32)
	_ReduceNotToPrefixUnaryOp                                             = _ReduceType(33)
	_ReduceBitNegToPrefixUnaryOp                                          = _ReduceType(34)
	_ReduceSubToPrefixUnaryOp                                             = _ReduceType(35)
	_ReduceMulToPrefixUnaryOp                                             = _ReduceType(36)
	_ReduceBitAndToPrefixUnaryOp                                          = _ReduceType(37)
	_ReducePostfixUnaryExprToPrefixUnaryExpr                              = _ReduceType(38)
	_ReducePrefixOpToPrefixUnaryExpr                                      = _ReduceType(39)
	_ReduceMulToMulOp                                                     = _ReduceType(40)
	_ReduceDivToMulOp                                                     = _ReduceType(41)
	_ReduceModToMulOp                                                     = _ReduceType(42)
	_ReduceBitAndToMulOp                                                  = _ReduceType(43)
	_ReduceBitLshiftToMulOp                                               = _ReduceType(44)
	_ReduceBitRshiftToMulOp                                               = _ReduceType(45)
	_ReducePrefixUnaryExprToMulExpr                                       = _ReduceType(46)
	_ReduceOpToMulExpr                                                    = _ReduceType(47)
	_ReduceAddToAddOp                                                     = _ReduceType(48)
	_ReduceSubToAddOp                                                     = _ReduceType(49)
	_ReduceBitOrToAddOp                                                   = _ReduceType(50)
	_ReduceBitXorToAddOp                                                  = _ReduceType(51)
	_ReduceMulExprToAddExpr                                               = _ReduceType(52)
	_ReduceOpToAddExpr                                                    = _ReduceType(53)
	_ReduceEqualToCmpOp                                                   = _ReduceType(54)
	_ReduceNotEqualToCmpOp                                                = _ReduceType(55)
	_ReduceLessToCmpOp                                                    = _ReduceType(56)
	_ReduceLessOrEqualToCmpOp                                             = _ReduceType(57)
	_ReduceGreaterToCmpOp                                                 = _ReduceType(58)
	_ReduceGreaterOrEqualToCmpOp                                          = _ReduceType(59)
	_ReduceAddExprToCmpExpr                                               = _ReduceType(60)
	_ReduceOpToCmpExpr                                                    = _ReduceType(61)
	_ReduceCmpExprToAndExpr                                               = _ReduceType(62)
	_ReduceOpToAndExpr                                                    = _ReduceType(63)
	_ReduceAndExprToOrExpr                                                = _ReduceType(64)
	_ReduceOpToOrExpr                                                     = _ReduceType(65)
	_ReduceToSequenceExpr                                                 = _ReduceType(66)
	_ReduceJumpLabelToOptionalJumpLabel                                   = _ReduceType(67)
	_ReduceUnlabelledToOptionalJumpLabel                                  = _ReduceType(68)
	_ReduceExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct = _ReduceType(69)
	_ReduceNilToOptionalExpressionOrImplicitStruct                        = _ReduceType(70)
	_ReduceReturnToJumpType                                               = _ReduceType(71)
	_ReduceBreakToJumpType                                                = _ReduceType(72)
	_ReduceContinueToJumpType                                             = _ReduceType(73)
	_ReduceAddOneAssignToOpOneAssign                                      = _ReduceType(74)
	_ReduceSubOneAssignToOpOneAssign                                      = _ReduceType(75)
	_ReduceAddAssignToBinaryOpAssign                                      = _ReduceType(76)
	_ReduceSubAssignToBinaryOpAssign                                      = _ReduceType(77)
	_ReduceMulAssignToBinaryOpAssign                                      = _ReduceType(78)
	_ReduceDivAssignToBinaryOpAssign                                      = _ReduceType(79)
	_ReduceModAssignToBinaryOpAssign                                      = _ReduceType(80)
	_ReduceBitNegAssignToBinaryOpAssign                                   = _ReduceType(81)
	_ReduceBitAndAssignToBinaryOpAssign                                   = _ReduceType(82)
	_ReduceBitOrAssignToBinaryOpAssign                                    = _ReduceType(83)
	_ReduceBitXorAssignToBinaryOpAssign                                   = _ReduceType(84)
	_ReduceBitLshiftAssignToBinaryOpAssign                                = _ReduceType(85)
	_ReduceBitRshiftAssignToBinaryOpAssign                                = _ReduceType(86)
	_ReduceExpressionToExpressionOrImplicitStruct                         = _ReduceType(87)
	_ReduceImplicitStructToExpressionOrImplicitStruct                     = _ReduceType(88)
	_ReduceToUnsafeStatement                                              = _ReduceType(89)
	_ReduceUnsafeStatementToStatementBody                                 = _ReduceType(90)
	_ReduceExpressionOrImplicitStructToStatementBody                      = _ReduceType(91)
	_ReduceAsyncToStatementBody                                           = _ReduceType(92)
	_ReduceJumpToStatementBody                                            = _ReduceType(93)
	_ReduceOpOneAssignToStatementBody                                     = _ReduceType(94)
	_ReduceBinaryOpAssignToStatementBody                                  = _ReduceType(95)
	_ReduceImplicitToStatement                                            = _ReduceType(96)
	_ReduceExplicitToStatement                                            = _ReduceType(97)
	_ReduceEmptyListToStatements                                          = _ReduceType(98)
	_ReduceAddToStatements                                                = _ReduceType(99)
	_ReduceLabelDeclToOptionalLabelDecl                                   = _ReduceType(100)
	_ReduceUnlabelledToOptionalLabelDecl                                  = _ReduceType(101)
	_ReduceToBlockBody                                                    = _ReduceType(102)
	_ReduceToBlockExpr                                                    = _ReduceType(103)
	_ReduceNoElseToIfExpr                                                 = _ReduceType(104)
	_ReduceIfElseToIfExpr                                                 = _ReduceType(105)
	_ReduceMultiIfElseToIfExpr                                            = _ReduceType(106)
	_ReduceToSwitchExpr                                                   = _ReduceType(107)
	_ReduceInfiniteToLoopExpr                                             = _ReduceType(108)
	_ReduceWhileToLoopExpr                                                = _ReduceType(109)
	_ReduceSequenceExprToExpression                                       = _ReduceType(110)
	_ReduceIfExprToExpression                                             = _ReduceType(111)
	_ReduceSwitchExprToExpression                                         = _ReduceType(112)
	_ReduceLoopExprToExpression                                           = _ReduceType(113)
	_ReduceNamedToAtomType                                                = _ReduceType(114)
	_ReduceExplicitStructTypeToAtomType                                   = _ReduceType(115)
	_ReduceImplicitStructTypeToAtomType                                   = _ReduceType(116)
	_ReduceImplicitEnumTypeToAtomType                                     = _ReduceType(117)
	_ReduceExplicitEnumTypeToAtomType                                     = _ReduceType(118)
	_ReduceTraitTypeToAtomType                                            = _ReduceType(119)
	_ReduceAtomTypeToTraitableType                                        = _ReduceType(120)
	_ReduceMethodInterfaceToTraitableType                                 = _ReduceType(121)
	_ReduceTraitToTraitableType                                           = _ReduceType(122)
	_ReduceTraitableTypeToTraitMulType                                    = _ReduceType(123)
	_ReduceIntersectToTraitMulType                                        = _ReduceType(124)
	_ReduceTraitMulTypeToTraitAddType                                     = _ReduceType(125)
	_ReduceUnionToTraitAddType                                            = _ReduceType(126)
	_ReduceDifferenceToTraitAddType                                       = _ReduceType(127)
	_ReduceExplicitToFieldDecl                                            = _ReduceType(128)
	_ReduceImplicitToFieldDecl                                            = _ReduceType(129)
	_ReduceFieldDeclToFieldDecls                                          = _ReduceType(130)
	_ReduceImplicitToFieldDecls                                           = _ReduceType(131)
	_ReduceExplicitToFieldDecls                                           = _ReduceType(132)
	_ReduceFieldDeclsToOptionalFieldDecls                                 = _ReduceType(133)
	_ReduceNilToOptionalFieldDecls                                        = _ReduceType(134)
	_ReduceToImplicitStructType                                           = _ReduceType(135)
	_ReducePairToFieldDeclOrList                                          = _ReduceType(136)
	_ReduceAddToFieldDeclOrList                                           = _ReduceType(137)
	_ReduceToImplicitEnumType                                             = _ReduceType(138)
	_ReduceToExplicitStructType                                           = _ReduceType(139)
	_ReduceToTraitType                                                    = _ReduceType(140)
	_ReduceToExplicitEnumType                                             = _ReduceType(141)
	_ReduceTraitAddTypeToValueType                                        = _ReduceType(142)
	_ReduceReferenceToValueType                                           = _ReduceType(143)
	_ReduceFuncTypeToValueType                                            = _ReduceType(144)
	_ReduceDefinitionToTypeDef                                            = _ReduceType(145)
	_ReduceAliasToTypeDef                                                 = _ReduceType(146)
	_ReduceUnconstrainedToGenericParameterDef                             = _ReduceType(147)
	_ReduceConstrainedToGenericParameterDef                               = _ReduceType(148)
	_ReduceGenericParameterDefToGenericParameterDefs                      = _ReduceType(149)
	_ReduceAddToGenericParameterDefs                                      = _ReduceType(150)
	_ReduceGenericParameterDefsToOptionalGenericParameterDefs             = _ReduceType(151)
	_ReduceNilToOptionalGenericParameterDefs                              = _ReduceType(152)
	_ReduceGenericToOptionalGenericParameters                             = _ReduceType(153)
	_ReduceNilToOptionalGenericParameters                                 = _ReduceType(154)
	_ReduceValueTypeToParameterType                                       = _ReduceType(155)
	_ReduceQuestionToParameterType                                        = _ReduceType(156)
	_ReduceValueTypeToReturnType                                          = _ReduceType(157)
	_ReduceQuestionToReturnType                                           = _ReduceType(158)
	_ReduceNilToReturnType                                                = _ReduceType(159)
	_ReduceParameterDefToParameterDecl                                    = _ReduceType(160)
	_ReduceUnamedToParameterDecl                                          = _ReduceType(161)
	_ReduceUnnamedVarargToParameterDecl                                   = _ReduceType(162)
	_ReduceParameterDeclToParameterDecls                                  = _ReduceType(163)
	_ReduceAddToParameterDecls                                            = _ReduceType(164)
	_ReduceParameterDeclsToOptionalParameterDecls                         = _ReduceType(165)
	_ReduceNilToOptionalParameterDecls                                    = _ReduceType(166)
	_ReduceToFuncType                                                     = _ReduceType(167)
	_ReduceArgToParameterDef                                              = _ReduceType(168)
	_ReduceVarargToParameterDef                                           = _ReduceType(169)
	_ReduceParameterDefToParameterDefs                                    = _ReduceType(170)
	_ReduceAddToParameterDefs                                             = _ReduceType(171)
	_ReduceParameterDefsToOptionalParameterDefs                           = _ReduceType(172)
	_ReduceNilToOptionalParameterDefs                                     = _ReduceType(173)
	_ReduceReceiverToOptionalReceiver                                     = _ReduceType(174)
	_ReduceNilToOptionalReceiver                                          = _ReduceType(175)
	_ReduceToNamedFuncDef                                                 = _ReduceType(176)
	_ReduceToAnonymousFuncExpr                                            = _ReduceType(177)
	_ReduceNoSpecToPackageDef                                             = _ReduceType(178)
	_ReduceWithSpecToPackageDef                                           = _ReduceType(179)
	_ReduceToPackageStatementBody                                         = _ReduceType(180)
	_ReduceImplicitToPackageStatement                                     = _ReduceType(181)
	_ReduceExplicitToPackageStatement                                     = _ReduceType(182)
	_ReduceEmptyListToPackageStatements                                   = _ReduceType(183)
	_ReduceAddToPackageStatements                                         = _ReduceType(184)
	_ReduceSpacesToLexInternalTokens                                      = _ReduceType(185)
	_ReduceCommentToLexInternalTokens                                     = _ReduceType(186)
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
	case _ReduceNamedToAtomType:
		return "NamedToAtomType"
	case _ReduceExplicitStructTypeToAtomType:
		return "ExplicitStructTypeToAtomType"
	case _ReduceImplicitStructTypeToAtomType:
		return "ImplicitStructTypeToAtomType"
	case _ReduceImplicitEnumTypeToAtomType:
		return "ImplicitEnumTypeToAtomType"
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
	case _ReducePairToFieldDeclOrList:
		return "PairToFieldDeclOrList"
	case _ReduceAddToFieldDeclOrList:
		return "AddToFieldDeclOrList"
	case _ReduceToImplicitEnumType:
		return "ToImplicitEnumType"
	case _ReduceToExplicitStructType:
		return "ToExplicitStructType"
	case _ReduceToTraitType:
		return "ToTraitType"
	case _ReduceToExplicitEnumType:
		return "ToExplicitEnumType"
	case _ReduceTraitAddTypeToValueType:
		return "TraitAddTypeToValueType"
	case _ReduceReferenceToValueType:
		return "ReferenceToValueType"
	case _ReduceFuncTypeToValueType:
		return "FuncTypeToValueType"
	case _ReduceDefinitionToTypeDef:
		return "DefinitionToTypeDef"
	case _ReduceAliasToTypeDef:
		return "AliasToTypeDef"
	case _ReduceUnconstrainedToGenericParameterDef:
		return "UnconstrainedToGenericParameterDef"
	case _ReduceConstrainedToGenericParameterDef:
		return "ConstrainedToGenericParameterDef"
	case _ReduceGenericParameterDefToGenericParameterDefs:
		return "GenericParameterDefToGenericParameterDefs"
	case _ReduceAddToGenericParameterDefs:
		return "AddToGenericParameterDefs"
	case _ReduceGenericParameterDefsToOptionalGenericParameterDefs:
		return "GenericParameterDefsToOptionalGenericParameterDefs"
	case _ReduceNilToOptionalGenericParameterDefs:
		return "NilToOptionalGenericParameterDefs"
	case _ReduceGenericToOptionalGenericParameters:
		return "GenericToOptionalGenericParameters"
	case _ReduceNilToOptionalGenericParameters:
		return "NilToOptionalGenericParameters"
	case _ReduceValueTypeToParameterType:
		return "ValueTypeToParameterType"
	case _ReduceQuestionToParameterType:
		return "QuestionToParameterType"
	case _ReduceValueTypeToReturnType:
		return "ValueTypeToReturnType"
	case _ReduceQuestionToReturnType:
		return "QuestionToReturnType"
	case _ReduceNilToReturnType:
		return "NilToReturnType"
	case _ReduceParameterDefToParameterDecl:
		return "ParameterDefToParameterDecl"
	case _ReduceUnamedToParameterDecl:
		return "UnamedToParameterDecl"
	case _ReduceUnnamedVarargToParameterDecl:
		return "UnnamedVarargToParameterDecl"
	case _ReduceParameterDeclToParameterDecls:
		return "ParameterDeclToParameterDecls"
	case _ReduceAddToParameterDecls:
		return "AddToParameterDecls"
	case _ReduceParameterDeclsToOptionalParameterDecls:
		return "ParameterDeclsToOptionalParameterDecls"
	case _ReduceNilToOptionalParameterDecls:
		return "NilToOptionalParameterDecls"
	case _ReduceToFuncType:
		return "ToFuncType"
	case _ReduceArgToParameterDef:
		return "ArgToParameterDef"
	case _ReduceVarargToParameterDef:
		return "VarargToParameterDef"
	case _ReduceParameterDefToParameterDefs:
		return "ParameterDefToParameterDefs"
	case _ReduceAddToParameterDefs:
		return "AddToParameterDefs"
	case _ReduceParameterDefsToOptionalParameterDefs:
		return "ParameterDefsToOptionalParameterDefs"
	case _ReduceNilToOptionalParameterDefs:
		return "NilToOptionalParameterDefs"
	case _ReduceReceiverToOptionalReceiver:
		return "ReceiverToOptionalReceiver"
	case _ReduceNilToOptionalReceiver:
		return "NilToOptionalReceiver"
	case _ReduceToNamedFuncDef:
		return "ToNamedFuncDef"
	case _ReduceToAnonymousFuncExpr:
		return "ToAnonymousFuncExpr"
	case _ReduceNoSpecToPackageDef:
		return "NoSpecToPackageDef"
	case _ReduceWithSpecToPackageDef:
		return "WithSpecToPackageDef"
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
	_State273 = _StateId(273)
	_State274 = _StateId(274)
	_State275 = _StateId(275)
	_State276 = _StateId(276)
	_State277 = _StateId(277)
	_State278 = _StateId(278)
	_State279 = _StateId(279)
	_State280 = _StateId(280)
	_State281 = _StateId(281)
	_State282 = _StateId(282)
	_State283 = _StateId(283)
	_State284 = _StateId(284)
	_State285 = _StateId(285)
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
	case _ReduceImplicitEnumTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ImplicitEnumTypeToAtomType(args[0].Generic_)
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
	case _ReducePairToFieldDeclOrList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = FieldDeclOrListType
		symbol.Generic_, err = reducer.PairToFieldDeclOrList(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceAddToFieldDeclOrList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = FieldDeclOrListType
		symbol.Generic_, err = reducer.AddToFieldDeclOrList(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToImplicitEnumType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitEnumTypeType
		symbol.Generic_, err = reducer.ToImplicitEnumType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	case _ReduceTraitAddTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.TraitAddTypeToValueType(args[0].Generic_)
	case _ReduceReferenceToValueType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.ReferenceToValueType(args[0].Generic_, args[1].Generic_)
	case _ReduceFuncTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.FuncTypeToValueType(args[0].Generic_)
	case _ReduceDefinitionToTypeDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = TypeDefType
		symbol.Generic_, err = reducer.DefinitionToTypeDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceAliasToTypeDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = TypeDefType
		symbol.Generic_, err = reducer.AliasToTypeDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceUnconstrainedToGenericParameterDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericParameterDefType
		symbol.Generic_, err = reducer.UnconstrainedToGenericParameterDef(args[0].Generic_)
	case _ReduceConstrainedToGenericParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = GenericParameterDefType
		symbol.Generic_, err = reducer.ConstrainedToGenericParameterDef(args[0].Generic_, args[1].Generic_)
	case _ReduceGenericParameterDefToGenericParameterDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericParameterDefsType
		symbol.Generic_, err = reducer.GenericParameterDefToGenericParameterDefs(args[0].Generic_)
	case _ReduceAddToGenericParameterDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = GenericParameterDefsType
		symbol.Generic_, err = reducer.AddToGenericParameterDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceGenericParameterDefsToOptionalGenericParameterDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalGenericParameterDefsType
		symbol.Generic_, err = reducer.GenericParameterDefsToOptionalGenericParameterDefs(args[0].Generic_)
	case _ReduceNilToOptionalGenericParameterDefs:
		symbol.SymbolId_ = OptionalGenericParameterDefsType
		symbol.Generic_, err = reducer.NilToOptionalGenericParameterDefs()
	case _ReduceGenericToOptionalGenericParameters:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OptionalGenericParametersType
		symbol.Generic_, err = reducer.GenericToOptionalGenericParameters(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceNilToOptionalGenericParameters:
		symbol.SymbolId_ = OptionalGenericParametersType
		symbol.Generic_, err = reducer.NilToOptionalGenericParameters()
	case _ReduceValueTypeToParameterType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterTypeType
		symbol.Generic_, err = reducer.ValueTypeToParameterType(args[0].Generic_)
	case _ReduceQuestionToParameterType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterTypeType
		symbol.Generic_, err = reducer.QuestionToParameterType(args[0].Generic_)
	case _ReduceValueTypeToReturnType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnTypeType
		symbol.Generic_, err = reducer.ValueTypeToReturnType(args[0].Generic_)
	case _ReduceQuestionToReturnType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnTypeType
		symbol.Generic_, err = reducer.QuestionToReturnType(args[0].Generic_)
	case _ReduceNilToReturnType:
		symbol.SymbolId_ = ReturnTypeType
		symbol.Generic_, err = reducer.NilToReturnType()
	case _ReduceParameterDefToParameterDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.ParameterDefToParameterDecl(args[0].Generic_)
	case _ReduceUnamedToParameterDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.UnamedToParameterDecl(args[0].Generic_)
	case _ReduceUnnamedVarargToParameterDecl:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.UnnamedVarargToParameterDecl(args[0].Generic_, args[1].Generic_)
	case _ReduceParameterDeclToParameterDecls:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDeclsType
		symbol.Generic_, err = reducer.ParameterDeclToParameterDecls(args[0].Generic_)
	case _ReduceAddToParameterDecls:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDeclsType
		symbol.Generic_, err = reducer.AddToParameterDecls(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceParameterDeclsToOptionalParameterDecls:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalParameterDeclsType
		symbol.Generic_, err = reducer.ParameterDeclsToOptionalParameterDecls(args[0].Generic_)
	case _ReduceNilToOptionalParameterDecls:
		symbol.SymbolId_ = OptionalParameterDeclsType
		symbol.Generic_, err = reducer.NilToOptionalParameterDecls()
	case _ReduceToFuncType:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = FuncTypeType
		symbol.Generic_, err = reducer.ToFuncType(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceArgToParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.ArgToParameterDef(args[0].Generic_, args[1].Generic_)
	case _ReduceVarargToParameterDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.VarargToParameterDef(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceParameterDefToParameterDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDefsType
		symbol.Generic_, err = reducer.ParameterDefToParameterDefs(args[0].Generic_)
	case _ReduceAddToParameterDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDefsType
		symbol.Generic_, err = reducer.AddToParameterDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceParameterDefsToOptionalParameterDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalParameterDefsType
		symbol.Generic_, err = reducer.ParameterDefsToOptionalParameterDefs(args[0].Generic_)
	case _ReduceNilToOptionalParameterDefs:
		symbol.SymbolId_ = OptionalParameterDefsType
		symbol.Generic_, err = reducer.NilToOptionalParameterDefs()
	case _ReduceReceiverToOptionalReceiver:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OptionalReceiverType
		symbol.Generic_, err = reducer.ReceiverToOptionalReceiver(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceNilToOptionalReceiver:
		symbol.SymbolId_ = OptionalReceiverType
		symbol.Generic_, err = reducer.NilToOptionalReceiver()
	case _ReduceToNamedFuncDef:
		args := stack[len(stack)-9:]
		stack = stack[:len(stack)-9]
		symbol.SymbolId_ = NamedFuncDefType
		symbol.Generic_, err = reducer.ToNamedFuncDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_, args[6].Generic_, args[7].Generic_, args[8].Generic_)
	case _ReduceToAnonymousFuncExpr:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = AnonymousFuncExprType
		symbol.Generic_, err = reducer.ToAnonymousFuncExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_)
	case _ReduceNoSpecToPackageDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PackageDefType
		symbol.Generic_, err = reducer.NoSpecToPackageDef(args[0].Generic_, args[1].Generic_)
	case _ReduceWithSpecToPackageDef:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = PackageDefType
		symbol.Generic_, err = reducer.WithSpecToPackageDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
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
	_GotoState273Action                                                         = &_Action{_ShiftAction, _State273, 0}
	_GotoState274Action                                                         = &_Action{_ShiftAction, _State274, 0}
	_GotoState275Action                                                         = &_Action{_ShiftAction, _State275, 0}
	_GotoState276Action                                                         = &_Action{_ShiftAction, _State276, 0}
	_GotoState277Action                                                         = &_Action{_ShiftAction, _State277, 0}
	_GotoState278Action                                                         = &_Action{_ShiftAction, _State278, 0}
	_GotoState279Action                                                         = &_Action{_ShiftAction, _State279, 0}
	_GotoState280Action                                                         = &_Action{_ShiftAction, _State280, 0}
	_GotoState281Action                                                         = &_Action{_ShiftAction, _State281, 0}
	_GotoState282Action                                                         = &_Action{_ShiftAction, _State282, 0}
	_GotoState283Action                                                         = &_Action{_ShiftAction, _State283, 0}
	_GotoState284Action                                                         = &_Action{_ShiftAction, _State284, 0}
	_GotoState285Action                                                         = &_Action{_ShiftAction, _State285, 0}
	_ReduceTrueToLiteralAction                                                  = &_Action{_ReduceAction, 0, _ReduceTrueToLiteral}
	_ReduceFalseToLiteralAction                                                 = &_Action{_ReduceAction, 0, _ReduceFalseToLiteral}
	_ReduceIntegerLiteralToLiteralAction                                        = &_Action{_ReduceAction, 0, _ReduceIntegerLiteralToLiteral}
	_ReduceFloatLiteralToLiteralAction                                          = &_Action{_ReduceAction, 0, _ReduceFloatLiteralToLiteral}
	_ReduceRuneLiteralToLiteralAction                                           = &_Action{_ReduceAction, 0, _ReduceRuneLiteralToLiteral}
	_ReduceStringLiteralToLiteralAction                                         = &_Action{_ReduceAction, 0, _ReduceStringLiteralToLiteral}
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
	_ReduceNamedToAtomTypeAction                                                = &_Action{_ReduceAction, 0, _ReduceNamedToAtomType}
	_ReduceExplicitStructTypeToAtomTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceExplicitStructTypeToAtomType}
	_ReduceImplicitStructTypeToAtomTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceImplicitStructTypeToAtomType}
	_ReduceImplicitEnumTypeToAtomTypeAction                                     = &_Action{_ReduceAction, 0, _ReduceImplicitEnumTypeToAtomType}
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
	_ReducePairToFieldDeclOrListAction                                          = &_Action{_ReduceAction, 0, _ReducePairToFieldDeclOrList}
	_ReduceAddToFieldDeclOrListAction                                           = &_Action{_ReduceAction, 0, _ReduceAddToFieldDeclOrList}
	_ReduceToImplicitEnumTypeAction                                             = &_Action{_ReduceAction, 0, _ReduceToImplicitEnumType}
	_ReduceToExplicitStructTypeAction                                           = &_Action{_ReduceAction, 0, _ReduceToExplicitStructType}
	_ReduceToTraitTypeAction                                                    = &_Action{_ReduceAction, 0, _ReduceToTraitType}
	_ReduceToExplicitEnumTypeAction                                             = &_Action{_ReduceAction, 0, _ReduceToExplicitEnumType}
	_ReduceTraitAddTypeToValueTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceTraitAddTypeToValueType}
	_ReduceReferenceToValueTypeAction                                           = &_Action{_ReduceAction, 0, _ReduceReferenceToValueType}
	_ReduceFuncTypeToValueTypeAction                                            = &_Action{_ReduceAction, 0, _ReduceFuncTypeToValueType}
	_ReduceDefinitionToTypeDefAction                                            = &_Action{_ReduceAction, 0, _ReduceDefinitionToTypeDef}
	_ReduceAliasToTypeDefAction                                                 = &_Action{_ReduceAction, 0, _ReduceAliasToTypeDef}
	_ReduceUnconstrainedToGenericParameterDefAction                             = &_Action{_ReduceAction, 0, _ReduceUnconstrainedToGenericParameterDef}
	_ReduceConstrainedToGenericParameterDefAction                               = &_Action{_ReduceAction, 0, _ReduceConstrainedToGenericParameterDef}
	_ReduceGenericParameterDefToGenericParameterDefsAction                      = &_Action{_ReduceAction, 0, _ReduceGenericParameterDefToGenericParameterDefs}
	_ReduceAddToGenericParameterDefsAction                                      = &_Action{_ReduceAction, 0, _ReduceAddToGenericParameterDefs}
	_ReduceGenericParameterDefsToOptionalGenericParameterDefsAction             = &_Action{_ReduceAction, 0, _ReduceGenericParameterDefsToOptionalGenericParameterDefs}
	_ReduceNilToOptionalGenericParameterDefsAction                              = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericParameterDefs}
	_ReduceGenericToOptionalGenericParametersAction                             = &_Action{_ReduceAction, 0, _ReduceGenericToOptionalGenericParameters}
	_ReduceNilToOptionalGenericParametersAction                                 = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericParameters}
	_ReduceValueTypeToParameterTypeAction                                       = &_Action{_ReduceAction, 0, _ReduceValueTypeToParameterType}
	_ReduceQuestionToParameterTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceQuestionToParameterType}
	_ReduceValueTypeToReturnTypeAction                                          = &_Action{_ReduceAction, 0, _ReduceValueTypeToReturnType}
	_ReduceQuestionToReturnTypeAction                                           = &_Action{_ReduceAction, 0, _ReduceQuestionToReturnType}
	_ReduceNilToReturnTypeAction                                                = &_Action{_ReduceAction, 0, _ReduceNilToReturnType}
	_ReduceParameterDefToParameterDeclAction                                    = &_Action{_ReduceAction, 0, _ReduceParameterDefToParameterDecl}
	_ReduceUnamedToParameterDeclAction                                          = &_Action{_ReduceAction, 0, _ReduceUnamedToParameterDecl}
	_ReduceUnnamedVarargToParameterDeclAction                                   = &_Action{_ReduceAction, 0, _ReduceUnnamedVarargToParameterDecl}
	_ReduceParameterDeclToParameterDeclsAction                                  = &_Action{_ReduceAction, 0, _ReduceParameterDeclToParameterDecls}
	_ReduceAddToParameterDeclsAction                                            = &_Action{_ReduceAction, 0, _ReduceAddToParameterDecls}
	_ReduceParameterDeclsToOptionalParameterDeclsAction                         = &_Action{_ReduceAction, 0, _ReduceParameterDeclsToOptionalParameterDecls}
	_ReduceNilToOptionalParameterDeclsAction                                    = &_Action{_ReduceAction, 0, _ReduceNilToOptionalParameterDecls}
	_ReduceToFuncTypeAction                                                     = &_Action{_ReduceAction, 0, _ReduceToFuncType}
	_ReduceArgToParameterDefAction                                              = &_Action{_ReduceAction, 0, _ReduceArgToParameterDef}
	_ReduceVarargToParameterDefAction                                           = &_Action{_ReduceAction, 0, _ReduceVarargToParameterDef}
	_ReduceParameterDefToParameterDefsAction                                    = &_Action{_ReduceAction, 0, _ReduceParameterDefToParameterDefs}
	_ReduceAddToParameterDefsAction                                             = &_Action{_ReduceAction, 0, _ReduceAddToParameterDefs}
	_ReduceParameterDefsToOptionalParameterDefsAction                           = &_Action{_ReduceAction, 0, _ReduceParameterDefsToOptionalParameterDefs}
	_ReduceNilToOptionalParameterDefsAction                                     = &_Action{_ReduceAction, 0, _ReduceNilToOptionalParameterDefs}
	_ReduceReceiverToOptionalReceiverAction                                     = &_Action{_ReduceAction, 0, _ReduceReceiverToOptionalReceiver}
	_ReduceNilToOptionalReceiverAction                                          = &_Action{_ReduceAction, 0, _ReduceNilToOptionalReceiver}
	_ReduceToNamedFuncDefAction                                                 = &_Action{_ReduceAction, 0, _ReduceToNamedFuncDef}
	_ReduceToAnonymousFuncExprAction                                            = &_Action{_ReduceAction, 0, _ReduceToAnonymousFuncExpr}
	_ReduceNoSpecToPackageDefAction                                             = &_Action{_ReduceAction, 0, _ReduceNoSpecToPackageDef}
	_ReduceWithSpecToPackageDefAction                                           = &_Action{_ReduceAction, 0, _ReduceWithSpecToPackageDef}
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
	{_State1, PackageDefType}:                           _GotoState6Action,
	{_State2, TypeToken}:                                _GotoState12Action,
	{_State2, TypeDefType}:                              _GotoState7Action,
	{_State3, FuncToken}:                                _GotoState13Action,
	{_State3, NamedFuncDefType}:                         _GotoState8Action,
	{_State4, IntegerLiteralToken}:                      _GotoState20Action,
	{_State4, FloatLiteralToken}:                        _GotoState17Action,
	{_State4, RuneLiteralToken}:                         _GotoState26Action,
	{_State4, StringLiteralToken}:                       _GotoState27Action,
	{_State4, IdentifierToken}:                          _GotoState19Action,
	{_State4, TrueToken}:                                _GotoState30Action,
	{_State4, FalseToken}:                               _GotoState16Action,
	{_State4, StructToken}:                              _GotoState28Action,
	{_State4, FuncToken}:                                _GotoState18Action,
	{_State4, LparenToken}:                              _GotoState23Action,
	{_State4, LabelDeclToken}:                           _GotoState21Action,
	{_State4, NotToken}:                                 _GotoState25Action,
	{_State4, SubToken}:                                 _GotoState29Action,
	{_State4, MulToken}:                                 _GotoState24Action,
	{_State4, BitNegToken}:                              _GotoState15Action,
	{_State4, BitAndToken}:                              _GotoState14Action,
	{_State4, LexErrorToken}:                            _GotoState22Action,
	{_State4, LiteralType}:                              _GotoState41Action,
	{_State4, AnonymousStructExprType}:                  _GotoState35Action,
	{_State4, AtomExprType}:                             _GotoState36Action,
	{_State4, CallExprType}:                             _GotoState38Action,
	{_State4, AccessExprType}:                           _GotoState31Action,
	{_State4, PostfixUnaryExprType}:                     _GotoState45Action,
	{_State4, PrefixUnaryOpType}:                        _GotoState47Action,
	{_State4, PrefixUnaryExprType}:                      _GotoState46Action,
	{_State4, MulExprType}:                              _GotoState42Action,
	{_State4, AddExprType}:                              _GotoState32Action,
	{_State4, CmpExprType}:                              _GotoState39Action,
	{_State4, AndExprType}:                              _GotoState33Action,
	{_State4, OrExprType}:                               _GotoState44Action,
	{_State4, SequenceExprType}:                         _GotoState48Action,
	{_State4, OptionalLabelDeclType}:                    _GotoState43Action,
	{_State4, BlockExprType}:                            _GotoState37Action,
	{_State4, ExpressionType}:                           _GotoState9Action,
	{_State4, ExplicitStructTypeType}:                   _GotoState40Action,
	{_State4, AnonymousFuncExprType}:                    _GotoState34Action,
	{_State5, SpacesToken}:                              _GotoState50Action,
	{_State5, CommentToken}:                             _GotoState49Action,
	{_State5, LexInternalTokensType}:                    _GotoState10Action,
	{_State11, IdentifierToken}:                         _GotoState51Action,
	{_State12, IdentifierToken}:                         _GotoState52Action,
	{_State13, LparenToken}:                             _GotoState53Action,
	{_State13, OptionalReceiverType}:                    _GotoState54Action,
	{_State18, LparenToken}:                             _GotoState55Action,
	{_State23, IntegerLiteralToken}:                     _GotoState20Action,
	{_State23, FloatLiteralToken}:                       _GotoState17Action,
	{_State23, RuneLiteralToken}:                        _GotoState26Action,
	{_State23, StringLiteralToken}:                      _GotoState27Action,
	{_State23, IdentifierToken}:                         _GotoState19Action,
	{_State23, TrueToken}:                               _GotoState30Action,
	{_State23, FalseToken}:                              _GotoState16Action,
	{_State23, StructToken}:                             _GotoState28Action,
	{_State23, FuncToken}:                               _GotoState18Action,
	{_State23, LparenToken}:                             _GotoState23Action,
	{_State23, LabelDeclToken}:                          _GotoState21Action,
	{_State23, NotToken}:                                _GotoState25Action,
	{_State23, SubToken}:                                _GotoState29Action,
	{_State23, MulToken}:                                _GotoState24Action,
	{_State23, BitNegToken}:                             _GotoState15Action,
	{_State23, BitAndToken}:                             _GotoState14Action,
	{_State23, LexErrorToken}:                           _GotoState22Action,
	{_State23, LiteralType}:                             _GotoState41Action,
	{_State23, AnonymousStructExprType}:                 _GotoState35Action,
	{_State23, AtomExprType}:                            _GotoState36Action,
	{_State23, ArgumentType}:                            _GotoState56Action,
	{_State23, ArgumentsType}:                           _GotoState57Action,
	{_State23, CallExprType}:                            _GotoState38Action,
	{_State23, AccessExprType}:                          _GotoState31Action,
	{_State23, PostfixUnaryExprType}:                    _GotoState45Action,
	{_State23, PrefixUnaryOpType}:                       _GotoState47Action,
	{_State23, PrefixUnaryExprType}:                     _GotoState46Action,
	{_State23, MulExprType}:                             _GotoState42Action,
	{_State23, AddExprType}:                             _GotoState32Action,
	{_State23, CmpExprType}:                             _GotoState39Action,
	{_State23, AndExprType}:                             _GotoState33Action,
	{_State23, OrExprType}:                              _GotoState44Action,
	{_State23, SequenceExprType}:                        _GotoState48Action,
	{_State23, OptionalLabelDeclType}:                   _GotoState43Action,
	{_State23, BlockExprType}:                           _GotoState37Action,
	{_State23, ExpressionType}:                          _GotoState58Action,
	{_State23, ExplicitStructTypeType}:                  _GotoState40Action,
	{_State23, AnonymousFuncExprType}:                   _GotoState34Action,
	{_State28, LparenToken}:                             _GotoState59Action,
	{_State28, ImplicitStructTypeType}:                  _GotoState60Action,
	{_State31, LbracketToken}:                           _GotoState63Action,
	{_State31, DotToken}:                                _GotoState62Action,
	{_State31, QuestionToken}:                           _GotoState64Action,
	{_State31, DollarLbracketToken}:                     _GotoState61Action,
	{_State31, OptionalGenericBindingType}:              _GotoState65Action,
	{_State32, AddToken}:                                _GotoState66Action,
	{_State32, SubToken}:                                _GotoState69Action,
	{_State32, BitXorToken}:                             _GotoState68Action,
	{_State32, BitOrToken}:                              _GotoState67Action,
	{_State32, AddOpType}:                               _GotoState70Action,
	{_State33, AndToken}:                                _GotoState71Action,
	{_State39, EqualToken}:                              _GotoState72Action,
	{_State39, NotEqualToken}:                           _GotoState77Action,
	{_State39, LessToken}:                               _GotoState75Action,
	{_State39, LessOrEqualToken}:                        _GotoState76Action,
	{_State39, GreaterToken}:                            _GotoState73Action,
	{_State39, GreaterOrEqualToken}:                     _GotoState74Action,
	{_State39, CmpOpType}:                               _GotoState78Action,
	{_State40, LparenToken}:                             _GotoState79Action,
	{_State42, MulToken}:                                _GotoState85Action,
	{_State42, DivToken}:                                _GotoState83Action,
	{_State42, ModToken}:                                _GotoState84Action,
	{_State42, BitAndToken}:                             _GotoState80Action,
	{_State42, BitLshiftToken}:                          _GotoState81Action,
	{_State42, BitRshiftToken}:                          _GotoState82Action,
	{_State42, MulOpType}:                               _GotoState86Action,
	{_State43, IfToken}:                                 _GotoState88Action,
	{_State43, SwitchToken}:                             _GotoState90Action,
	{_State43, ForToken}:                                _GotoState87Action,
	{_State43, LbraceToken}:                             _GotoState89Action,
	{_State43, BlockBodyType}:                           _GotoState91Action,
	{_State43, IfExprType}:                              _GotoState92Action,
	{_State43, SwitchExprType}:                          _GotoState94Action,
	{_State43, LoopExprType}:                            _GotoState93Action,
	{_State44, OrToken}:                                 _GotoState95Action,
	{_State47, IntegerLiteralToken}:                     _GotoState20Action,
	{_State47, FloatLiteralToken}:                       _GotoState17Action,
	{_State47, RuneLiteralToken}:                        _GotoState26Action,
	{_State47, StringLiteralToken}:                      _GotoState27Action,
	{_State47, IdentifierToken}:                         _GotoState19Action,
	{_State47, TrueToken}:                               _GotoState30Action,
	{_State47, FalseToken}:                              _GotoState16Action,
	{_State47, StructToken}:                             _GotoState28Action,
	{_State47, FuncToken}:                               _GotoState18Action,
	{_State47, LparenToken}:                             _GotoState23Action,
	{_State47, LabelDeclToken}:                          _GotoState21Action,
	{_State47, NotToken}:                                _GotoState25Action,
	{_State47, SubToken}:                                _GotoState29Action,
	{_State47, MulToken}:                                _GotoState24Action,
	{_State47, BitNegToken}:                             _GotoState15Action,
	{_State47, BitAndToken}:                             _GotoState14Action,
	{_State47, LexErrorToken}:                           _GotoState22Action,
	{_State47, LiteralType}:                             _GotoState41Action,
	{_State47, AnonymousStructExprType}:                 _GotoState35Action,
	{_State47, AtomExprType}:                            _GotoState36Action,
	{_State47, CallExprType}:                            _GotoState38Action,
	{_State47, AccessExprType}:                          _GotoState31Action,
	{_State47, PostfixUnaryExprType}:                    _GotoState45Action,
	{_State47, PrefixUnaryOpType}:                       _GotoState47Action,
	{_State47, PrefixUnaryExprType}:                     _GotoState97Action,
	{_State47, OptionalLabelDeclType}:                   _GotoState96Action,
	{_State47, BlockExprType}:                           _GotoState37Action,
	{_State47, ExplicitStructTypeType}:                  _GotoState40Action,
	{_State47, AnonymousFuncExprType}:                   _GotoState34Action,
	{_State51, LparenToken}:                             _GotoState98Action,
	{_State52, DollarLbracketToken}:                     _GotoState99Action,
	{_State52, EqualToken}:                              _GotoState100Action,
	{_State52, OptionalGenericParametersType}:           _GotoState101Action,
	{_State53, IdentifierToken}:                         _GotoState102Action,
	{_State53, ParameterDefType}:                        _GotoState103Action,
	{_State54, IdentifierToken}:                         _GotoState104Action,
	{_State55, IdentifierToken}:                         _GotoState102Action,
	{_State55, ParameterDefType}:                        _GotoState106Action,
	{_State55, ParameterDefsType}:                       _GotoState107Action,
	{_State55, OptionalParameterDefsType}:               _GotoState105Action,
	{_State57, RparenToken}:                             _GotoState109Action,
	{_State57, CommaToken}:                              _GotoState108Action,
	{_State59, IdentifierToken}:                         _GotoState115Action,
	{_State59, StructToken}:                             _GotoState28Action,
	{_State59, EnumToken}:                               _GotoState111Action,
	{_State59, TraitToken}:                              _GotoState117Action,
	{_State59, FuncToken}:                               _GotoState114Action,
	{_State59, LparenToken}:                             _GotoState116Action,
	{_State59, ExclaimToken}:                            _GotoState112Action,
	{_State59, ExclaimExclaimToken}:                     _GotoState113Action,
	{_State59, BitAndToken}:                             _GotoState110Action,
	{_State59, AtomTypeType}:                            _GotoState118Action,
	{_State59, TraitableTypeType}:                       _GotoState130Action,
	{_State59, TraitMulTypeType}:                        _GotoState128Action,
	{_State59, TraitAddTypeType}:                        _GotoState127Action,
	{_State59, FieldDeclType}:                           _GotoState121Action,
	{_State59, FieldDeclsType}:                          _GotoState122Action,
	{_State59, OptionalFieldDeclsType}:                  _GotoState126Action,
	{_State59, ImplicitStructTypeType}:                  _GotoState125Action,
	{_State59, ImplicitEnumTypeType}:                    _GotoState124Action,
	{_State59, ExplicitStructTypeType}:                  _GotoState120Action,
	{_State59, TraitTypeType}:                           _GotoState129Action,
	{_State59, ExplicitEnumTypeType}:                    _GotoState119Action,
	{_State59, ValueTypeType}:                           _GotoState131Action,
	{_State59, FuncTypeType}:                            _GotoState123Action,
	{_State61, IdentifierToken}:                         _GotoState132Action,
	{_State61, StructToken}:                             _GotoState28Action,
	{_State61, EnumToken}:                               _GotoState111Action,
	{_State61, TraitToken}:                              _GotoState117Action,
	{_State61, FuncToken}:                               _GotoState114Action,
	{_State61, LparenToken}:                             _GotoState116Action,
	{_State61, ExclaimToken}:                            _GotoState112Action,
	{_State61, ExclaimExclaimToken}:                     _GotoState113Action,
	{_State61, BitAndToken}:                             _GotoState110Action,
	{_State61, GenericArgumentsType}:                    _GotoState133Action,
	{_State61, OptionalGenericArgumentsType}:            _GotoState134Action,
	{_State61, AtomTypeType}:                            _GotoState118Action,
	{_State61, TraitableTypeType}:                       _GotoState130Action,
	{_State61, TraitMulTypeType}:                        _GotoState128Action,
	{_State61, TraitAddTypeType}:                        _GotoState127Action,
	{_State61, ImplicitStructTypeType}:                  _GotoState125Action,
	{_State61, ImplicitEnumTypeType}:                    _GotoState124Action,
	{_State61, ExplicitStructTypeType}:                  _GotoState120Action,
	{_State61, TraitTypeType}:                           _GotoState129Action,
	{_State61, ExplicitEnumTypeType}:                    _GotoState119Action,
	{_State61, ValueTypeType}:                           _GotoState135Action,
	{_State61, FuncTypeType}:                            _GotoState123Action,
	{_State62, IdentifierToken}:                         _GotoState136Action,
	{_State63, IntegerLiteralToken}:                     _GotoState20Action,
	{_State63, FloatLiteralToken}:                       _GotoState17Action,
	{_State63, RuneLiteralToken}:                        _GotoState26Action,
	{_State63, StringLiteralToken}:                      _GotoState27Action,
	{_State63, IdentifierToken}:                         _GotoState19Action,
	{_State63, TrueToken}:                               _GotoState30Action,
	{_State63, FalseToken}:                              _GotoState16Action,
	{_State63, StructToken}:                             _GotoState28Action,
	{_State63, FuncToken}:                               _GotoState18Action,
	{_State63, LparenToken}:                             _GotoState23Action,
	{_State63, LabelDeclToken}:                          _GotoState21Action,
	{_State63, NotToken}:                                _GotoState25Action,
	{_State63, SubToken}:                                _GotoState29Action,
	{_State63, MulToken}:                                _GotoState24Action,
	{_State63, BitNegToken}:                             _GotoState15Action,
	{_State63, BitAndToken}:                             _GotoState14Action,
	{_State63, LexErrorToken}:                           _GotoState22Action,
	{_State63, LiteralType}:                             _GotoState41Action,
	{_State63, AnonymousStructExprType}:                 _GotoState35Action,
	{_State63, AtomExprType}:                            _GotoState36Action,
	{_State63, CallExprType}:                            _GotoState38Action,
	{_State63, AccessExprType}:                          _GotoState31Action,
	{_State63, PostfixUnaryExprType}:                    _GotoState45Action,
	{_State63, PrefixUnaryOpType}:                       _GotoState47Action,
	{_State63, PrefixUnaryExprType}:                     _GotoState46Action,
	{_State63, MulExprType}:                             _GotoState42Action,
	{_State63, AddExprType}:                             _GotoState32Action,
	{_State63, CmpExprType}:                             _GotoState39Action,
	{_State63, AndExprType}:                             _GotoState33Action,
	{_State63, OrExprType}:                              _GotoState44Action,
	{_State63, SequenceExprType}:                        _GotoState48Action,
	{_State63, OptionalLabelDeclType}:                   _GotoState43Action,
	{_State63, BlockExprType}:                           _GotoState37Action,
	{_State63, ExpressionType}:                          _GotoState137Action,
	{_State63, ExplicitStructTypeType}:                  _GotoState40Action,
	{_State63, AnonymousFuncExprType}:                   _GotoState34Action,
	{_State65, LparenToken}:                             _GotoState138Action,
	{_State70, IntegerLiteralToken}:                     _GotoState20Action,
	{_State70, FloatLiteralToken}:                       _GotoState17Action,
	{_State70, RuneLiteralToken}:                        _GotoState26Action,
	{_State70, StringLiteralToken}:                      _GotoState27Action,
	{_State70, IdentifierToken}:                         _GotoState19Action,
	{_State70, TrueToken}:                               _GotoState30Action,
	{_State70, FalseToken}:                              _GotoState16Action,
	{_State70, StructToken}:                             _GotoState28Action,
	{_State70, FuncToken}:                               _GotoState18Action,
	{_State70, LparenToken}:                             _GotoState23Action,
	{_State70, LabelDeclToken}:                          _GotoState21Action,
	{_State70, NotToken}:                                _GotoState25Action,
	{_State70, SubToken}:                                _GotoState29Action,
	{_State70, MulToken}:                                _GotoState24Action,
	{_State70, BitNegToken}:                             _GotoState15Action,
	{_State70, BitAndToken}:                             _GotoState14Action,
	{_State70, LexErrorToken}:                           _GotoState22Action,
	{_State70, LiteralType}:                             _GotoState41Action,
	{_State70, AnonymousStructExprType}:                 _GotoState35Action,
	{_State70, AtomExprType}:                            _GotoState36Action,
	{_State70, CallExprType}:                            _GotoState38Action,
	{_State70, AccessExprType}:                          _GotoState31Action,
	{_State70, PostfixUnaryExprType}:                    _GotoState45Action,
	{_State70, PrefixUnaryOpType}:                       _GotoState47Action,
	{_State70, PrefixUnaryExprType}:                     _GotoState46Action,
	{_State70, MulExprType}:                             _GotoState139Action,
	{_State70, OptionalLabelDeclType}:                   _GotoState96Action,
	{_State70, BlockExprType}:                           _GotoState37Action,
	{_State70, ExplicitStructTypeType}:                  _GotoState40Action,
	{_State70, AnonymousFuncExprType}:                   _GotoState34Action,
	{_State71, IntegerLiteralToken}:                     _GotoState20Action,
	{_State71, FloatLiteralToken}:                       _GotoState17Action,
	{_State71, RuneLiteralToken}:                        _GotoState26Action,
	{_State71, StringLiteralToken}:                      _GotoState27Action,
	{_State71, IdentifierToken}:                         _GotoState19Action,
	{_State71, TrueToken}:                               _GotoState30Action,
	{_State71, FalseToken}:                              _GotoState16Action,
	{_State71, StructToken}:                             _GotoState28Action,
	{_State71, FuncToken}:                               _GotoState18Action,
	{_State71, LparenToken}:                             _GotoState23Action,
	{_State71, LabelDeclToken}:                          _GotoState21Action,
	{_State71, NotToken}:                                _GotoState25Action,
	{_State71, SubToken}:                                _GotoState29Action,
	{_State71, MulToken}:                                _GotoState24Action,
	{_State71, BitNegToken}:                             _GotoState15Action,
	{_State71, BitAndToken}:                             _GotoState14Action,
	{_State71, LexErrorToken}:                           _GotoState22Action,
	{_State71, LiteralType}:                             _GotoState41Action,
	{_State71, AnonymousStructExprType}:                 _GotoState35Action,
	{_State71, AtomExprType}:                            _GotoState36Action,
	{_State71, CallExprType}:                            _GotoState38Action,
	{_State71, AccessExprType}:                          _GotoState31Action,
	{_State71, PostfixUnaryExprType}:                    _GotoState45Action,
	{_State71, PrefixUnaryOpType}:                       _GotoState47Action,
	{_State71, PrefixUnaryExprType}:                     _GotoState46Action,
	{_State71, MulExprType}:                             _GotoState42Action,
	{_State71, AddExprType}:                             _GotoState32Action,
	{_State71, CmpExprType}:                             _GotoState140Action,
	{_State71, OptionalLabelDeclType}:                   _GotoState96Action,
	{_State71, BlockExprType}:                           _GotoState37Action,
	{_State71, ExplicitStructTypeType}:                  _GotoState40Action,
	{_State71, AnonymousFuncExprType}:                   _GotoState34Action,
	{_State78, IntegerLiteralToken}:                     _GotoState20Action,
	{_State78, FloatLiteralToken}:                       _GotoState17Action,
	{_State78, RuneLiteralToken}:                        _GotoState26Action,
	{_State78, StringLiteralToken}:                      _GotoState27Action,
	{_State78, IdentifierToken}:                         _GotoState19Action,
	{_State78, TrueToken}:                               _GotoState30Action,
	{_State78, FalseToken}:                              _GotoState16Action,
	{_State78, StructToken}:                             _GotoState28Action,
	{_State78, FuncToken}:                               _GotoState18Action,
	{_State78, LparenToken}:                             _GotoState23Action,
	{_State78, LabelDeclToken}:                          _GotoState21Action,
	{_State78, NotToken}:                                _GotoState25Action,
	{_State78, SubToken}:                                _GotoState29Action,
	{_State78, MulToken}:                                _GotoState24Action,
	{_State78, BitNegToken}:                             _GotoState15Action,
	{_State78, BitAndToken}:                             _GotoState14Action,
	{_State78, LexErrorToken}:                           _GotoState22Action,
	{_State78, LiteralType}:                             _GotoState41Action,
	{_State78, AnonymousStructExprType}:                 _GotoState35Action,
	{_State78, AtomExprType}:                            _GotoState36Action,
	{_State78, CallExprType}:                            _GotoState38Action,
	{_State78, AccessExprType}:                          _GotoState31Action,
	{_State78, PostfixUnaryExprType}:                    _GotoState45Action,
	{_State78, PrefixUnaryOpType}:                       _GotoState47Action,
	{_State78, PrefixUnaryExprType}:                     _GotoState46Action,
	{_State78, MulExprType}:                             _GotoState42Action,
	{_State78, AddExprType}:                             _GotoState141Action,
	{_State78, OptionalLabelDeclType}:                   _GotoState96Action,
	{_State78, BlockExprType}:                           _GotoState37Action,
	{_State78, ExplicitStructTypeType}:                  _GotoState40Action,
	{_State78, AnonymousFuncExprType}:                   _GotoState34Action,
	{_State79, IntegerLiteralToken}:                     _GotoState20Action,
	{_State79, FloatLiteralToken}:                       _GotoState17Action,
	{_State79, RuneLiteralToken}:                        _GotoState26Action,
	{_State79, StringLiteralToken}:                      _GotoState27Action,
	{_State79, IdentifierToken}:                         _GotoState19Action,
	{_State79, TrueToken}:                               _GotoState30Action,
	{_State79, FalseToken}:                              _GotoState16Action,
	{_State79, StructToken}:                             _GotoState28Action,
	{_State79, FuncToken}:                               _GotoState18Action,
	{_State79, LparenToken}:                             _GotoState23Action,
	{_State79, LabelDeclToken}:                          _GotoState21Action,
	{_State79, NotToken}:                                _GotoState25Action,
	{_State79, SubToken}:                                _GotoState29Action,
	{_State79, MulToken}:                                _GotoState24Action,
	{_State79, BitNegToken}:                             _GotoState15Action,
	{_State79, BitAndToken}:                             _GotoState14Action,
	{_State79, LexErrorToken}:                           _GotoState22Action,
	{_State79, LiteralType}:                             _GotoState41Action,
	{_State79, AnonymousStructExprType}:                 _GotoState35Action,
	{_State79, AtomExprType}:                            _GotoState36Action,
	{_State79, ArgumentType}:                            _GotoState56Action,
	{_State79, ArgumentsType}:                           _GotoState142Action,
	{_State79, CallExprType}:                            _GotoState38Action,
	{_State79, AccessExprType}:                          _GotoState31Action,
	{_State79, PostfixUnaryExprType}:                    _GotoState45Action,
	{_State79, PrefixUnaryOpType}:                       _GotoState47Action,
	{_State79, PrefixUnaryExprType}:                     _GotoState46Action,
	{_State79, MulExprType}:                             _GotoState42Action,
	{_State79, AddExprType}:                             _GotoState32Action,
	{_State79, CmpExprType}:                             _GotoState39Action,
	{_State79, AndExprType}:                             _GotoState33Action,
	{_State79, OrExprType}:                              _GotoState44Action,
	{_State79, SequenceExprType}:                        _GotoState48Action,
	{_State79, OptionalLabelDeclType}:                   _GotoState43Action,
	{_State79, BlockExprType}:                           _GotoState37Action,
	{_State79, ExpressionType}:                          _GotoState58Action,
	{_State79, ExplicitStructTypeType}:                  _GotoState40Action,
	{_State79, AnonymousFuncExprType}:                   _GotoState34Action,
	{_State86, IntegerLiteralToken}:                     _GotoState20Action,
	{_State86, FloatLiteralToken}:                       _GotoState17Action,
	{_State86, RuneLiteralToken}:                        _GotoState26Action,
	{_State86, StringLiteralToken}:                      _GotoState27Action,
	{_State86, IdentifierToken}:                         _GotoState19Action,
	{_State86, TrueToken}:                               _GotoState30Action,
	{_State86, FalseToken}:                              _GotoState16Action,
	{_State86, StructToken}:                             _GotoState28Action,
	{_State86, FuncToken}:                               _GotoState18Action,
	{_State86, LparenToken}:                             _GotoState23Action,
	{_State86, LabelDeclToken}:                          _GotoState21Action,
	{_State86, NotToken}:                                _GotoState25Action,
	{_State86, SubToken}:                                _GotoState29Action,
	{_State86, MulToken}:                                _GotoState24Action,
	{_State86, BitNegToken}:                             _GotoState15Action,
	{_State86, BitAndToken}:                             _GotoState14Action,
	{_State86, LexErrorToken}:                           _GotoState22Action,
	{_State86, LiteralType}:                             _GotoState41Action,
	{_State86, AnonymousStructExprType}:                 _GotoState35Action,
	{_State86, AtomExprType}:                            _GotoState36Action,
	{_State86, CallExprType}:                            _GotoState38Action,
	{_State86, AccessExprType}:                          _GotoState31Action,
	{_State86, PostfixUnaryExprType}:                    _GotoState45Action,
	{_State86, PrefixUnaryOpType}:                       _GotoState47Action,
	{_State86, PrefixUnaryExprType}:                     _GotoState143Action,
	{_State86, OptionalLabelDeclType}:                   _GotoState96Action,
	{_State86, BlockExprType}:                           _GotoState37Action,
	{_State86, ExplicitStructTypeType}:                  _GotoState40Action,
	{_State86, AnonymousFuncExprType}:                   _GotoState34Action,
	{_State87, IntegerLiteralToken}:                     _GotoState20Action,
	{_State87, FloatLiteralToken}:                       _GotoState17Action,
	{_State87, RuneLiteralToken}:                        _GotoState26Action,
	{_State87, StringLiteralToken}:                      _GotoState27Action,
	{_State87, IdentifierToken}:                         _GotoState19Action,
	{_State87, TrueToken}:                               _GotoState30Action,
	{_State87, FalseToken}:                              _GotoState16Action,
	{_State87, StructToken}:                             _GotoState28Action,
	{_State87, FuncToken}:                               _GotoState18Action,
	{_State87, LparenToken}:                             _GotoState23Action,
	{_State87, LabelDeclToken}:                          _GotoState21Action,
	{_State87, NotToken}:                                _GotoState25Action,
	{_State87, SubToken}:                                _GotoState29Action,
	{_State87, MulToken}:                                _GotoState24Action,
	{_State87, BitNegToken}:                             _GotoState15Action,
	{_State87, BitAndToken}:                             _GotoState14Action,
	{_State87, LexErrorToken}:                           _GotoState22Action,
	{_State87, LiteralType}:                             _GotoState41Action,
	{_State87, AnonymousStructExprType}:                 _GotoState35Action,
	{_State87, AtomExprType}:                            _GotoState36Action,
	{_State87, CallExprType}:                            _GotoState38Action,
	{_State87, AccessExprType}:                          _GotoState31Action,
	{_State87, PostfixUnaryExprType}:                    _GotoState45Action,
	{_State87, PrefixUnaryOpType}:                       _GotoState47Action,
	{_State87, PrefixUnaryExprType}:                     _GotoState46Action,
	{_State87, MulExprType}:                             _GotoState42Action,
	{_State87, AddExprType}:                             _GotoState32Action,
	{_State87, CmpExprType}:                             _GotoState39Action,
	{_State87, AndExprType}:                             _GotoState33Action,
	{_State87, OrExprType}:                              _GotoState44Action,
	{_State87, SequenceExprType}:                        _GotoState145Action,
	{_State87, OptionalLabelDeclType}:                   _GotoState96Action,
	{_State87, BlockExprType}:                           _GotoState144Action,
	{_State87, ExplicitStructTypeType}:                  _GotoState40Action,
	{_State87, AnonymousFuncExprType}:                   _GotoState34Action,
	{_State88, IntegerLiteralToken}:                     _GotoState20Action,
	{_State88, FloatLiteralToken}:                       _GotoState17Action,
	{_State88, RuneLiteralToken}:                        _GotoState26Action,
	{_State88, StringLiteralToken}:                      _GotoState27Action,
	{_State88, IdentifierToken}:                         _GotoState19Action,
	{_State88, TrueToken}:                               _GotoState30Action,
	{_State88, FalseToken}:                              _GotoState16Action,
	{_State88, StructToken}:                             _GotoState28Action,
	{_State88, FuncToken}:                               _GotoState18Action,
	{_State88, LparenToken}:                             _GotoState23Action,
	{_State88, LabelDeclToken}:                          _GotoState21Action,
	{_State88, NotToken}:                                _GotoState25Action,
	{_State88, SubToken}:                                _GotoState29Action,
	{_State88, MulToken}:                                _GotoState24Action,
	{_State88, BitNegToken}:                             _GotoState15Action,
	{_State88, BitAndToken}:                             _GotoState14Action,
	{_State88, LexErrorToken}:                           _GotoState22Action,
	{_State88, LiteralType}:                             _GotoState41Action,
	{_State88, AnonymousStructExprType}:                 _GotoState35Action,
	{_State88, AtomExprType}:                            _GotoState36Action,
	{_State88, CallExprType}:                            _GotoState38Action,
	{_State88, AccessExprType}:                          _GotoState31Action,
	{_State88, PostfixUnaryExprType}:                    _GotoState45Action,
	{_State88, PrefixUnaryOpType}:                       _GotoState47Action,
	{_State88, PrefixUnaryExprType}:                     _GotoState46Action,
	{_State88, MulExprType}:                             _GotoState42Action,
	{_State88, AddExprType}:                             _GotoState32Action,
	{_State88, CmpExprType}:                             _GotoState39Action,
	{_State88, AndExprType}:                             _GotoState33Action,
	{_State88, OrExprType}:                              _GotoState44Action,
	{_State88, SequenceExprType}:                        _GotoState146Action,
	{_State88, OptionalLabelDeclType}:                   _GotoState96Action,
	{_State88, BlockExprType}:                           _GotoState37Action,
	{_State88, ExplicitStructTypeType}:                  _GotoState40Action,
	{_State88, AnonymousFuncExprType}:                   _GotoState34Action,
	{_State89, StatementsType}:                          _GotoState147Action,
	{_State90, LbraceToken}:                             _GotoState148Action,
	{_State95, IntegerLiteralToken}:                     _GotoState20Action,
	{_State95, FloatLiteralToken}:                       _GotoState17Action,
	{_State95, RuneLiteralToken}:                        _GotoState26Action,
	{_State95, StringLiteralToken}:                      _GotoState27Action,
	{_State95, IdentifierToken}:                         _GotoState19Action,
	{_State95, TrueToken}:                               _GotoState30Action,
	{_State95, FalseToken}:                              _GotoState16Action,
	{_State95, StructToken}:                             _GotoState28Action,
	{_State95, FuncToken}:                               _GotoState18Action,
	{_State95, LparenToken}:                             _GotoState23Action,
	{_State95, LabelDeclToken}:                          _GotoState21Action,
	{_State95, NotToken}:                                _GotoState25Action,
	{_State95, SubToken}:                                _GotoState29Action,
	{_State95, MulToken}:                                _GotoState24Action,
	{_State95, BitNegToken}:                             _GotoState15Action,
	{_State95, BitAndToken}:                             _GotoState14Action,
	{_State95, LexErrorToken}:                           _GotoState22Action,
	{_State95, LiteralType}:                             _GotoState41Action,
	{_State95, AnonymousStructExprType}:                 _GotoState35Action,
	{_State95, AtomExprType}:                            _GotoState36Action,
	{_State95, CallExprType}:                            _GotoState38Action,
	{_State95, AccessExprType}:                          _GotoState31Action,
	{_State95, PostfixUnaryExprType}:                    _GotoState45Action,
	{_State95, PrefixUnaryOpType}:                       _GotoState47Action,
	{_State95, PrefixUnaryExprType}:                     _GotoState46Action,
	{_State95, MulExprType}:                             _GotoState42Action,
	{_State95, AddExprType}:                             _GotoState32Action,
	{_State95, CmpExprType}:                             _GotoState39Action,
	{_State95, AndExprType}:                             _GotoState149Action,
	{_State95, OptionalLabelDeclType}:                   _GotoState96Action,
	{_State95, BlockExprType}:                           _GotoState37Action,
	{_State95, ExplicitStructTypeType}:                  _GotoState40Action,
	{_State95, AnonymousFuncExprType}:                   _GotoState34Action,
	{_State96, LbraceToken}:                             _GotoState89Action,
	{_State96, BlockBodyType}:                           _GotoState91Action,
	{_State98, PackageStatementsType}:                   _GotoState150Action,
	{_State99, IdentifierToken}:                         _GotoState151Action,
	{_State99, GenericParameterDefType}:                 _GotoState152Action,
	{_State99, GenericParameterDefsType}:                _GotoState153Action,
	{_State99, OptionalGenericParameterDefsType}:        _GotoState154Action,
	{_State100, IdentifierToken}:                        _GotoState132Action,
	{_State100, StructToken}:                            _GotoState28Action,
	{_State100, EnumToken}:                              _GotoState111Action,
	{_State100, TraitToken}:                             _GotoState117Action,
	{_State100, FuncToken}:                              _GotoState114Action,
	{_State100, LparenToken}:                            _GotoState116Action,
	{_State100, ExclaimToken}:                           _GotoState112Action,
	{_State100, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State100, BitAndToken}:                            _GotoState110Action,
	{_State100, AtomTypeType}:                           _GotoState118Action,
	{_State100, TraitableTypeType}:                      _GotoState130Action,
	{_State100, TraitMulTypeType}:                       _GotoState128Action,
	{_State100, TraitAddTypeType}:                       _GotoState127Action,
	{_State100, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State100, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State100, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State100, TraitTypeType}:                          _GotoState129Action,
	{_State100, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State100, ValueTypeType}:                          _GotoState155Action,
	{_State100, FuncTypeType}:                           _GotoState123Action,
	{_State101, IdentifierToken}:                        _GotoState132Action,
	{_State101, StructToken}:                            _GotoState28Action,
	{_State101, EnumToken}:                              _GotoState111Action,
	{_State101, TraitToken}:                             _GotoState117Action,
	{_State101, FuncToken}:                              _GotoState114Action,
	{_State101, LparenToken}:                            _GotoState116Action,
	{_State101, ExclaimToken}:                           _GotoState112Action,
	{_State101, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State101, BitAndToken}:                            _GotoState110Action,
	{_State101, AtomTypeType}:                           _GotoState118Action,
	{_State101, TraitableTypeType}:                      _GotoState130Action,
	{_State101, TraitMulTypeType}:                       _GotoState128Action,
	{_State101, TraitAddTypeType}:                       _GotoState127Action,
	{_State101, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State101, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State101, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State101, TraitTypeType}:                          _GotoState129Action,
	{_State101, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State101, ValueTypeType}:                          _GotoState156Action,
	{_State101, FuncTypeType}:                           _GotoState123Action,
	{_State102, IdentifierToken}:                        _GotoState132Action,
	{_State102, StructToken}:                            _GotoState28Action,
	{_State102, EnumToken}:                              _GotoState111Action,
	{_State102, TraitToken}:                             _GotoState117Action,
	{_State102, FuncToken}:                              _GotoState114Action,
	{_State102, LparenToken}:                            _GotoState116Action,
	{_State102, QuestionToken}:                          _GotoState158Action,
	{_State102, DotdotdotToken}:                         _GotoState157Action,
	{_State102, ExclaimToken}:                           _GotoState112Action,
	{_State102, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State102, BitAndToken}:                            _GotoState110Action,
	{_State102, AtomTypeType}:                           _GotoState118Action,
	{_State102, TraitableTypeType}:                      _GotoState130Action,
	{_State102, TraitMulTypeType}:                       _GotoState128Action,
	{_State102, TraitAddTypeType}:                       _GotoState127Action,
	{_State102, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State102, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State102, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State102, TraitTypeType}:                          _GotoState129Action,
	{_State102, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State102, ValueTypeType}:                          _GotoState160Action,
	{_State102, ParameterTypeType}:                      _GotoState159Action,
	{_State102, FuncTypeType}:                           _GotoState123Action,
	{_State103, RparenToken}:                            _GotoState161Action,
	{_State104, DollarLbracketToken}:                    _GotoState99Action,
	{_State104, OptionalGenericParametersType}:          _GotoState162Action,
	{_State105, RparenToken}:                            _GotoState163Action,
	{_State107, CommaToken}:                             _GotoState164Action,
	{_State108, IntegerLiteralToken}:                    _GotoState20Action,
	{_State108, FloatLiteralToken}:                      _GotoState17Action,
	{_State108, RuneLiteralToken}:                       _GotoState26Action,
	{_State108, StringLiteralToken}:                     _GotoState27Action,
	{_State108, IdentifierToken}:                        _GotoState19Action,
	{_State108, TrueToken}:                              _GotoState30Action,
	{_State108, FalseToken}:                             _GotoState16Action,
	{_State108, StructToken}:                            _GotoState28Action,
	{_State108, FuncToken}:                              _GotoState18Action,
	{_State108, LparenToken}:                            _GotoState23Action,
	{_State108, LabelDeclToken}:                         _GotoState21Action,
	{_State108, NotToken}:                               _GotoState25Action,
	{_State108, SubToken}:                               _GotoState29Action,
	{_State108, MulToken}:                               _GotoState24Action,
	{_State108, BitNegToken}:                            _GotoState15Action,
	{_State108, BitAndToken}:                            _GotoState14Action,
	{_State108, LexErrorToken}:                          _GotoState22Action,
	{_State108, LiteralType}:                            _GotoState41Action,
	{_State108, AnonymousStructExprType}:                _GotoState35Action,
	{_State108, AtomExprType}:                           _GotoState36Action,
	{_State108, ArgumentType}:                           _GotoState165Action,
	{_State108, CallExprType}:                           _GotoState38Action,
	{_State108, AccessExprType}:                         _GotoState31Action,
	{_State108, PostfixUnaryExprType}:                   _GotoState45Action,
	{_State108, PrefixUnaryOpType}:                      _GotoState47Action,
	{_State108, PrefixUnaryExprType}:                    _GotoState46Action,
	{_State108, MulExprType}:                            _GotoState42Action,
	{_State108, AddExprType}:                            _GotoState32Action,
	{_State108, CmpExprType}:                            _GotoState39Action,
	{_State108, AndExprType}:                            _GotoState33Action,
	{_State108, OrExprType}:                             _GotoState44Action,
	{_State108, SequenceExprType}:                       _GotoState48Action,
	{_State108, OptionalLabelDeclType}:                  _GotoState43Action,
	{_State108, BlockExprType}:                          _GotoState37Action,
	{_State108, ExpressionType}:                         _GotoState58Action,
	{_State108, ExplicitStructTypeType}:                 _GotoState40Action,
	{_State108, AnonymousFuncExprType}:                  _GotoState34Action,
	{_State110, IdentifierToken}:                        _GotoState132Action,
	{_State110, StructToken}:                            _GotoState28Action,
	{_State110, EnumToken}:                              _GotoState111Action,
	{_State110, TraitToken}:                             _GotoState117Action,
	{_State110, LparenToken}:                            _GotoState116Action,
	{_State110, ExclaimToken}:                           _GotoState112Action,
	{_State110, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State110, AtomTypeType}:                           _GotoState118Action,
	{_State110, TraitableTypeType}:                      _GotoState130Action,
	{_State110, TraitMulTypeType}:                       _GotoState128Action,
	{_State110, TraitAddTypeType}:                       _GotoState166Action,
	{_State110, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State110, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State110, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State110, TraitTypeType}:                          _GotoState129Action,
	{_State110, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State111, LparenToken}:                            _GotoState167Action,
	{_State112, IdentifierToken}:                        _GotoState132Action,
	{_State112, StructToken}:                            _GotoState28Action,
	{_State112, EnumToken}:                              _GotoState111Action,
	{_State112, TraitToken}:                             _GotoState117Action,
	{_State112, LparenToken}:                            _GotoState116Action,
	{_State112, AtomTypeType}:                           _GotoState168Action,
	{_State112, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State112, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State112, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State112, TraitTypeType}:                          _GotoState129Action,
	{_State112, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State113, IdentifierToken}:                        _GotoState132Action,
	{_State113, StructToken}:                            _GotoState28Action,
	{_State113, EnumToken}:                              _GotoState111Action,
	{_State113, TraitToken}:                             _GotoState117Action,
	{_State113, LparenToken}:                            _GotoState116Action,
	{_State113, AtomTypeType}:                           _GotoState169Action,
	{_State113, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State113, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State113, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State113, TraitTypeType}:                          _GotoState129Action,
	{_State113, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State114, LparenToken}:                            _GotoState170Action,
	{_State115, IdentifierToken}:                        _GotoState132Action,
	{_State115, StructToken}:                            _GotoState28Action,
	{_State115, EnumToken}:                              _GotoState111Action,
	{_State115, TraitToken}:                             _GotoState117Action,
	{_State115, FuncToken}:                              _GotoState114Action,
	{_State115, LparenToken}:                            _GotoState116Action,
	{_State115, DollarLbracketToken}:                    _GotoState61Action,
	{_State115, ExclaimToken}:                           _GotoState112Action,
	{_State115, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State115, BitAndToken}:                            _GotoState110Action,
	{_State115, OptionalGenericBindingType}:             _GotoState171Action,
	{_State115, AtomTypeType}:                           _GotoState118Action,
	{_State115, TraitableTypeType}:                      _GotoState130Action,
	{_State115, TraitMulTypeType}:                       _GotoState128Action,
	{_State115, TraitAddTypeType}:                       _GotoState127Action,
	{_State115, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State115, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State115, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State115, TraitTypeType}:                          _GotoState129Action,
	{_State115, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State115, ValueTypeType}:                          _GotoState172Action,
	{_State115, FuncTypeType}:                           _GotoState123Action,
	{_State116, IdentifierToken}:                        _GotoState115Action,
	{_State116, StructToken}:                            _GotoState28Action,
	{_State116, EnumToken}:                              _GotoState111Action,
	{_State116, TraitToken}:                             _GotoState117Action,
	{_State116, FuncToken}:                              _GotoState114Action,
	{_State116, LparenToken}:                            _GotoState116Action,
	{_State116, ExclaimToken}:                           _GotoState112Action,
	{_State116, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State116, BitAndToken}:                            _GotoState110Action,
	{_State116, AtomTypeType}:                           _GotoState118Action,
	{_State116, TraitableTypeType}:                      _GotoState130Action,
	{_State116, TraitMulTypeType}:                       _GotoState128Action,
	{_State116, TraitAddTypeType}:                       _GotoState127Action,
	{_State116, FieldDeclType}:                          _GotoState173Action,
	{_State116, FieldDeclsType}:                         _GotoState122Action,
	{_State116, OptionalFieldDeclsType}:                 _GotoState126Action,
	{_State116, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State116, FieldDeclOrListType}:                    _GotoState174Action,
	{_State116, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State116, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State116, TraitTypeType}:                          _GotoState129Action,
	{_State116, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State116, ValueTypeType}:                          _GotoState131Action,
	{_State116, FuncTypeType}:                           _GotoState123Action,
	{_State117, LparenToken}:                            _GotoState175Action,
	{_State122, NewlinesToken}:                          _GotoState177Action,
	{_State122, CommaToken}:                             _GotoState176Action,
	{_State126, RparenToken}:                            _GotoState178Action,
	{_State127, AddToken}:                               _GotoState179Action,
	{_State127, SubToken}:                               _GotoState180Action,
	{_State128, MulToken}:                               _GotoState181Action,
	{_State132, DollarLbracketToken}:                    _GotoState61Action,
	{_State132, OptionalGenericBindingType}:             _GotoState171Action,
	{_State133, CommaToken}:                             _GotoState182Action,
	{_State134, RbracketToken}:                          _GotoState183Action,
	{_State137, RbracketToken}:                          _GotoState184Action,
	{_State138, IntegerLiteralToken}:                    _GotoState20Action,
	{_State138, FloatLiteralToken}:                      _GotoState17Action,
	{_State138, RuneLiteralToken}:                       _GotoState26Action,
	{_State138, StringLiteralToken}:                     _GotoState27Action,
	{_State138, IdentifierToken}:                        _GotoState19Action,
	{_State138, TrueToken}:                              _GotoState30Action,
	{_State138, FalseToken}:                             _GotoState16Action,
	{_State138, StructToken}:                            _GotoState28Action,
	{_State138, FuncToken}:                              _GotoState18Action,
	{_State138, LparenToken}:                            _GotoState23Action,
	{_State138, LabelDeclToken}:                         _GotoState21Action,
	{_State138, NotToken}:                               _GotoState25Action,
	{_State138, SubToken}:                               _GotoState29Action,
	{_State138, MulToken}:                               _GotoState24Action,
	{_State138, BitNegToken}:                            _GotoState15Action,
	{_State138, BitAndToken}:                            _GotoState14Action,
	{_State138, LexErrorToken}:                          _GotoState22Action,
	{_State138, LiteralType}:                            _GotoState41Action,
	{_State138, AnonymousStructExprType}:                _GotoState35Action,
	{_State138, AtomExprType}:                           _GotoState36Action,
	{_State138, ArgumentType}:                           _GotoState56Action,
	{_State138, ArgumentsType}:                          _GotoState185Action,
	{_State138, OptionalArgumentsType}:                  _GotoState186Action,
	{_State138, CallExprType}:                           _GotoState38Action,
	{_State138, AccessExprType}:                         _GotoState31Action,
	{_State138, PostfixUnaryExprType}:                   _GotoState45Action,
	{_State138, PrefixUnaryOpType}:                      _GotoState47Action,
	{_State138, PrefixUnaryExprType}:                    _GotoState46Action,
	{_State138, MulExprType}:                            _GotoState42Action,
	{_State138, AddExprType}:                            _GotoState32Action,
	{_State138, CmpExprType}:                            _GotoState39Action,
	{_State138, AndExprType}:                            _GotoState33Action,
	{_State138, OrExprType}:                             _GotoState44Action,
	{_State138, SequenceExprType}:                       _GotoState48Action,
	{_State138, OptionalLabelDeclType}:                  _GotoState43Action,
	{_State138, BlockExprType}:                          _GotoState37Action,
	{_State138, ExpressionType}:                         _GotoState58Action,
	{_State138, ExplicitStructTypeType}:                 _GotoState40Action,
	{_State138, AnonymousFuncExprType}:                  _GotoState34Action,
	{_State139, MulToken}:                               _GotoState85Action,
	{_State139, DivToken}:                               _GotoState83Action,
	{_State139, ModToken}:                               _GotoState84Action,
	{_State139, BitAndToken}:                            _GotoState80Action,
	{_State139, BitLshiftToken}:                         _GotoState81Action,
	{_State139, BitRshiftToken}:                         _GotoState82Action,
	{_State139, MulOpType}:                              _GotoState86Action,
	{_State140, EqualToken}:                             _GotoState72Action,
	{_State140, NotEqualToken}:                          _GotoState77Action,
	{_State140, LessToken}:                              _GotoState75Action,
	{_State140, LessOrEqualToken}:                       _GotoState76Action,
	{_State140, GreaterToken}:                           _GotoState73Action,
	{_State140, GreaterOrEqualToken}:                    _GotoState74Action,
	{_State140, CmpOpType}:                              _GotoState78Action,
	{_State141, AddToken}:                               _GotoState66Action,
	{_State141, SubToken}:                               _GotoState69Action,
	{_State141, BitXorToken}:                            _GotoState68Action,
	{_State141, BitOrToken}:                             _GotoState67Action,
	{_State141, AddOpType}:                              _GotoState70Action,
	{_State142, RparenToken}:                            _GotoState187Action,
	{_State142, CommaToken}:                             _GotoState108Action,
	{_State145, LabelDeclToken}:                         _GotoState21Action,
	{_State145, OptionalLabelDeclType}:                  _GotoState96Action,
	{_State145, BlockExprType}:                          _GotoState188Action,
	{_State146, LbraceToken}:                            _GotoState89Action,
	{_State146, BlockBodyType}:                          _GotoState189Action,
	{_State147, IntegerLiteralToken}:                    _GotoState20Action,
	{_State147, FloatLiteralToken}:                      _GotoState17Action,
	{_State147, RuneLiteralToken}:                       _GotoState26Action,
	{_State147, StringLiteralToken}:                     _GotoState27Action,
	{_State147, IdentifierToken}:                        _GotoState19Action,
	{_State147, TrueToken}:                              _GotoState30Action,
	{_State147, FalseToken}:                             _GotoState16Action,
	{_State147, ReturnToken}:                            _GotoState194Action,
	{_State147, BreakToken}:                             _GotoState191Action,
	{_State147, ContinueToken}:                          _GotoState192Action,
	{_State147, UnsafeToken}:                            _GotoState195Action,
	{_State147, StructToken}:                            _GotoState28Action,
	{_State147, FuncToken}:                              _GotoState18Action,
	{_State147, AsyncToken}:                             _GotoState190Action,
	{_State147, RbraceToken}:                            _GotoState193Action,
	{_State147, LparenToken}:                            _GotoState23Action,
	{_State147, LabelDeclToken}:                         _GotoState21Action,
	{_State147, NotToken}:                               _GotoState25Action,
	{_State147, SubToken}:                               _GotoState29Action,
	{_State147, MulToken}:                               _GotoState24Action,
	{_State147, BitNegToken}:                            _GotoState15Action,
	{_State147, BitAndToken}:                            _GotoState14Action,
	{_State147, LexErrorToken}:                          _GotoState22Action,
	{_State147, LiteralType}:                            _GotoState41Action,
	{_State147, AnonymousStructExprType}:                _GotoState35Action,
	{_State147, AtomExprType}:                           _GotoState36Action,
	{_State147, CallExprType}:                           _GotoState38Action,
	{_State147, AccessExprType}:                         _GotoState196Action,
	{_State147, PostfixUnaryExprType}:                   _GotoState45Action,
	{_State147, PrefixUnaryOpType}:                      _GotoState47Action,
	{_State147, PrefixUnaryExprType}:                    _GotoState46Action,
	{_State147, MulExprType}:                            _GotoState42Action,
	{_State147, AddExprType}:                            _GotoState32Action,
	{_State147, CmpExprType}:                            _GotoState39Action,
	{_State147, AndExprType}:                            _GotoState33Action,
	{_State147, OrExprType}:                             _GotoState44Action,
	{_State147, SequenceExprType}:                       _GotoState48Action,
	{_State147, JumpTypeType}:                           _GotoState199Action,
	{_State147, ExpressionOrImplicitStructType}:         _GotoState198Action,
	{_State147, UnsafeStatementType}:                    _GotoState202Action,
	{_State147, StatementBodyType}:                      _GotoState201Action,
	{_State147, StatementType}:                          _GotoState200Action,
	{_State147, OptionalLabelDeclType}:                  _GotoState43Action,
	{_State147, BlockExprType}:                          _GotoState37Action,
	{_State147, ExpressionType}:                         _GotoState197Action,
	{_State147, ExplicitStructTypeType}:                 _GotoState40Action,
	{_State147, AnonymousFuncExprType}:                  _GotoState34Action,
	{_State148, CaseToken}:                              _GotoState203Action,
	{_State149, AndToken}:                               _GotoState71Action,
	{_State150, UnsafeToken}:                            _GotoState195Action,
	{_State150, RparenToken}:                            _GotoState204Action,
	{_State150, UnsafeStatementType}:                    _GotoState207Action,
	{_State150, PackageStatementBodyType}:               _GotoState206Action,
	{_State150, PackageStatementType}:                   _GotoState205Action,
	{_State151, IdentifierToken}:                        _GotoState132Action,
	{_State151, StructToken}:                            _GotoState28Action,
	{_State151, EnumToken}:                              _GotoState111Action,
	{_State151, TraitToken}:                             _GotoState117Action,
	{_State151, FuncToken}:                              _GotoState114Action,
	{_State151, LparenToken}:                            _GotoState116Action,
	{_State151, ExclaimToken}:                           _GotoState112Action,
	{_State151, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State151, BitAndToken}:                            _GotoState110Action,
	{_State151, AtomTypeType}:                           _GotoState118Action,
	{_State151, TraitableTypeType}:                      _GotoState130Action,
	{_State151, TraitMulTypeType}:                       _GotoState128Action,
	{_State151, TraitAddTypeType}:                       _GotoState127Action,
	{_State151, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State151, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State151, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State151, TraitTypeType}:                          _GotoState129Action,
	{_State151, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State151, ValueTypeType}:                          _GotoState208Action,
	{_State151, FuncTypeType}:                           _GotoState123Action,
	{_State153, CommaToken}:                             _GotoState209Action,
	{_State154, RbracketToken}:                          _GotoState210Action,
	{_State157, IdentifierToken}:                        _GotoState132Action,
	{_State157, StructToken}:                            _GotoState28Action,
	{_State157, EnumToken}:                              _GotoState111Action,
	{_State157, TraitToken}:                             _GotoState117Action,
	{_State157, FuncToken}:                              _GotoState114Action,
	{_State157, LparenToken}:                            _GotoState116Action,
	{_State157, QuestionToken}:                          _GotoState158Action,
	{_State157, ExclaimToken}:                           _GotoState112Action,
	{_State157, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State157, BitAndToken}:                            _GotoState110Action,
	{_State157, AtomTypeType}:                           _GotoState118Action,
	{_State157, TraitableTypeType}:                      _GotoState130Action,
	{_State157, TraitMulTypeType}:                       _GotoState128Action,
	{_State157, TraitAddTypeType}:                       _GotoState127Action,
	{_State157, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State157, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State157, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State157, TraitTypeType}:                          _GotoState129Action,
	{_State157, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State157, ValueTypeType}:                          _GotoState160Action,
	{_State157, ParameterTypeType}:                      _GotoState211Action,
	{_State157, FuncTypeType}:                           _GotoState123Action,
	{_State162, LparenToken}:                            _GotoState212Action,
	{_State163, IdentifierToken}:                        _GotoState132Action,
	{_State163, StructToken}:                            _GotoState28Action,
	{_State163, EnumToken}:                              _GotoState111Action,
	{_State163, TraitToken}:                             _GotoState117Action,
	{_State163, FuncToken}:                              _GotoState114Action,
	{_State163, LparenToken}:                            _GotoState116Action,
	{_State163, QuestionToken}:                          _GotoState213Action,
	{_State163, ExclaimToken}:                           _GotoState112Action,
	{_State163, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State163, BitAndToken}:                            _GotoState110Action,
	{_State163, AtomTypeType}:                           _GotoState118Action,
	{_State163, TraitableTypeType}:                      _GotoState130Action,
	{_State163, TraitMulTypeType}:                       _GotoState128Action,
	{_State163, TraitAddTypeType}:                       _GotoState127Action,
	{_State163, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State163, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State163, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State163, TraitTypeType}:                          _GotoState129Action,
	{_State163, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State163, ValueTypeType}:                          _GotoState215Action,
	{_State163, ReturnTypeType}:                         _GotoState214Action,
	{_State163, FuncTypeType}:                           _GotoState123Action,
	{_State164, IdentifierToken}:                        _GotoState102Action,
	{_State164, ParameterDefType}:                       _GotoState216Action,
	{_State166, AddToken}:                               _GotoState179Action,
	{_State166, SubToken}:                               _GotoState180Action,
	{_State167, RparenToken}:                            _GotoState217Action,
	{_State170, IdentifierToken}:                        _GotoState219Action,
	{_State170, StructToken}:                            _GotoState28Action,
	{_State170, EnumToken}:                              _GotoState111Action,
	{_State170, TraitToken}:                             _GotoState117Action,
	{_State170, FuncToken}:                              _GotoState114Action,
	{_State170, LparenToken}:                            _GotoState116Action,
	{_State170, QuestionToken}:                          _GotoState158Action,
	{_State170, DotdotdotToken}:                         _GotoState218Action,
	{_State170, ExclaimToken}:                           _GotoState112Action,
	{_State170, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State170, BitAndToken}:                            _GotoState110Action,
	{_State170, AtomTypeType}:                           _GotoState118Action,
	{_State170, TraitableTypeType}:                      _GotoState130Action,
	{_State170, TraitMulTypeType}:                       _GotoState128Action,
	{_State170, TraitAddTypeType}:                       _GotoState127Action,
	{_State170, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State170, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State170, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State170, TraitTypeType}:                          _GotoState129Action,
	{_State170, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State170, ValueTypeType}:                          _GotoState160Action,
	{_State170, ParameterTypeType}:                      _GotoState224Action,
	{_State170, ParameterDeclType}:                      _GotoState221Action,
	{_State170, ParameterDeclsType}:                     _GotoState222Action,
	{_State170, OptionalParameterDeclsType}:             _GotoState220Action,
	{_State170, FuncTypeType}:                           _GotoState123Action,
	{_State170, ParameterDefType}:                       _GotoState223Action,
	{_State173, OrToken}:                                _GotoState225Action,
	{_State174, RparenToken}:                            _GotoState227Action,
	{_State174, OrToken}:                                _GotoState226Action,
	{_State175, RparenToken}:                            _GotoState228Action,
	{_State176, IdentifierToken}:                        _GotoState115Action,
	{_State176, StructToken}:                            _GotoState28Action,
	{_State176, EnumToken}:                              _GotoState111Action,
	{_State176, TraitToken}:                             _GotoState117Action,
	{_State176, FuncToken}:                              _GotoState114Action,
	{_State176, LparenToken}:                            _GotoState116Action,
	{_State176, ExclaimToken}:                           _GotoState112Action,
	{_State176, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State176, BitAndToken}:                            _GotoState110Action,
	{_State176, AtomTypeType}:                           _GotoState118Action,
	{_State176, TraitableTypeType}:                      _GotoState130Action,
	{_State176, TraitMulTypeType}:                       _GotoState128Action,
	{_State176, TraitAddTypeType}:                       _GotoState127Action,
	{_State176, FieldDeclType}:                          _GotoState229Action,
	{_State176, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State176, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State176, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State176, TraitTypeType}:                          _GotoState129Action,
	{_State176, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State176, ValueTypeType}:                          _GotoState131Action,
	{_State176, FuncTypeType}:                           _GotoState123Action,
	{_State177, IdentifierToken}:                        _GotoState115Action,
	{_State177, StructToken}:                            _GotoState28Action,
	{_State177, EnumToken}:                              _GotoState111Action,
	{_State177, TraitToken}:                             _GotoState117Action,
	{_State177, FuncToken}:                              _GotoState114Action,
	{_State177, LparenToken}:                            _GotoState116Action,
	{_State177, ExclaimToken}:                           _GotoState112Action,
	{_State177, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State177, BitAndToken}:                            _GotoState110Action,
	{_State177, AtomTypeType}:                           _GotoState118Action,
	{_State177, TraitableTypeType}:                      _GotoState130Action,
	{_State177, TraitMulTypeType}:                       _GotoState128Action,
	{_State177, TraitAddTypeType}:                       _GotoState127Action,
	{_State177, FieldDeclType}:                          _GotoState230Action,
	{_State177, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State177, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State177, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State177, TraitTypeType}:                          _GotoState129Action,
	{_State177, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State177, ValueTypeType}:                          _GotoState131Action,
	{_State177, FuncTypeType}:                           _GotoState123Action,
	{_State179, IdentifierToken}:                        _GotoState132Action,
	{_State179, StructToken}:                            _GotoState28Action,
	{_State179, EnumToken}:                              _GotoState111Action,
	{_State179, TraitToken}:                             _GotoState117Action,
	{_State179, LparenToken}:                            _GotoState116Action,
	{_State179, ExclaimToken}:                           _GotoState112Action,
	{_State179, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State179, AtomTypeType}:                           _GotoState118Action,
	{_State179, TraitableTypeType}:                      _GotoState130Action,
	{_State179, TraitMulTypeType}:                       _GotoState231Action,
	{_State179, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State179, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State179, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State179, TraitTypeType}:                          _GotoState129Action,
	{_State179, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State180, IdentifierToken}:                        _GotoState132Action,
	{_State180, StructToken}:                            _GotoState28Action,
	{_State180, EnumToken}:                              _GotoState111Action,
	{_State180, TraitToken}:                             _GotoState117Action,
	{_State180, LparenToken}:                            _GotoState116Action,
	{_State180, ExclaimToken}:                           _GotoState112Action,
	{_State180, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State180, AtomTypeType}:                           _GotoState118Action,
	{_State180, TraitableTypeType}:                      _GotoState130Action,
	{_State180, TraitMulTypeType}:                       _GotoState232Action,
	{_State180, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State180, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State180, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State180, TraitTypeType}:                          _GotoState129Action,
	{_State180, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State181, IdentifierToken}:                        _GotoState132Action,
	{_State181, StructToken}:                            _GotoState28Action,
	{_State181, EnumToken}:                              _GotoState111Action,
	{_State181, TraitToken}:                             _GotoState117Action,
	{_State181, LparenToken}:                            _GotoState116Action,
	{_State181, ExclaimToken}:                           _GotoState112Action,
	{_State181, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State181, AtomTypeType}:                           _GotoState118Action,
	{_State181, TraitableTypeType}:                      _GotoState233Action,
	{_State181, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State181, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State181, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State181, TraitTypeType}:                          _GotoState129Action,
	{_State181, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State182, IdentifierToken}:                        _GotoState132Action,
	{_State182, StructToken}:                            _GotoState28Action,
	{_State182, EnumToken}:                              _GotoState111Action,
	{_State182, TraitToken}:                             _GotoState117Action,
	{_State182, FuncToken}:                              _GotoState114Action,
	{_State182, LparenToken}:                            _GotoState116Action,
	{_State182, ExclaimToken}:                           _GotoState112Action,
	{_State182, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State182, BitAndToken}:                            _GotoState110Action,
	{_State182, AtomTypeType}:                           _GotoState118Action,
	{_State182, TraitableTypeType}:                      _GotoState130Action,
	{_State182, TraitMulTypeType}:                       _GotoState128Action,
	{_State182, TraitAddTypeType}:                       _GotoState127Action,
	{_State182, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State182, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State182, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State182, TraitTypeType}:                          _GotoState129Action,
	{_State182, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State182, ValueTypeType}:                          _GotoState234Action,
	{_State182, FuncTypeType}:                           _GotoState123Action,
	{_State185, CommaToken}:                             _GotoState108Action,
	{_State186, RparenToken}:                            _GotoState235Action,
	{_State189, ElseToken}:                              _GotoState236Action,
	{_State190, IntegerLiteralToken}:                    _GotoState20Action,
	{_State190, FloatLiteralToken}:                      _GotoState17Action,
	{_State190, RuneLiteralToken}:                       _GotoState26Action,
	{_State190, StringLiteralToken}:                     _GotoState27Action,
	{_State190, IdentifierToken}:                        _GotoState19Action,
	{_State190, TrueToken}:                              _GotoState30Action,
	{_State190, FalseToken}:                             _GotoState16Action,
	{_State190, StructToken}:                            _GotoState28Action,
	{_State190, FuncToken}:                              _GotoState18Action,
	{_State190, LparenToken}:                            _GotoState23Action,
	{_State190, LabelDeclToken}:                         _GotoState21Action,
	{_State190, LexErrorToken}:                          _GotoState22Action,
	{_State190, LiteralType}:                            _GotoState41Action,
	{_State190, AnonymousStructExprType}:                _GotoState35Action,
	{_State190, AtomExprType}:                           _GotoState36Action,
	{_State190, CallExprType}:                           _GotoState238Action,
	{_State190, AccessExprType}:                         _GotoState237Action,
	{_State190, OptionalLabelDeclType}:                  _GotoState96Action,
	{_State190, BlockExprType}:                          _GotoState37Action,
	{_State190, ExplicitStructTypeType}:                 _GotoState40Action,
	{_State190, AnonymousFuncExprType}:                  _GotoState34Action,
	{_State195, LessToken}:                              _GotoState239Action,
	{_State196, LbracketToken}:                          _GotoState63Action,
	{_State196, DotToken}:                               _GotoState62Action,
	{_State196, QuestionToken}:                          _GotoState64Action,
	{_State196, DollarLbracketToken}:                    _GotoState61Action,
	{_State196, AddAssignToken}:                         _GotoState240Action,
	{_State196, SubAssignToken}:                         _GotoState251Action,
	{_State196, MulAssignToken}:                         _GotoState250Action,
	{_State196, DivAssignToken}:                         _GotoState248Action,
	{_State196, ModAssignToken}:                         _GotoState249Action,
	{_State196, AddOneAssignToken}:                      _GotoState241Action,
	{_State196, SubOneAssignToken}:                      _GotoState252Action,
	{_State196, BitNegAssignToken}:                      _GotoState244Action,
	{_State196, BitAndAssignToken}:                      _GotoState242Action,
	{_State196, BitOrAssignToken}:                       _GotoState245Action,
	{_State196, BitXorAssignToken}:                      _GotoState247Action,
	{_State196, BitLshiftAssignToken}:                   _GotoState243Action,
	{_State196, BitRshiftAssignToken}:                   _GotoState246Action,
	{_State196, OptionalGenericBindingType}:             _GotoState65Action,
	{_State196, OpOneAssignType}:                        _GotoState254Action,
	{_State196, BinaryOpAssignType}:                     _GotoState253Action,
	{_State198, CommaToken}:                             _GotoState255Action,
	{_State199, JumpLabelToken}:                         _GotoState256Action,
	{_State199, OptionalJumpLabelType}:                  _GotoState257Action,
	{_State201, NewlinesToken}:                          _GotoState258Action,
	{_State201, SemicolonToken}:                         _GotoState259Action,
	{_State203, DefaultToken}:                           _GotoState260Action,
	{_State206, NewlinesToken}:                          _GotoState261Action,
	{_State206, SemicolonToken}:                         _GotoState262Action,
	{_State209, IdentifierToken}:                        _GotoState151Action,
	{_State209, GenericParameterDefType}:                _GotoState263Action,
	{_State212, IdentifierToken}:                        _GotoState102Action,
	{_State212, ParameterDefType}:                       _GotoState106Action,
	{_State212, ParameterDefsType}:                      _GotoState107Action,
	{_State212, OptionalParameterDefsType}:              _GotoState264Action,
	{_State214, LbraceToken}:                            _GotoState89Action,
	{_State214, BlockBodyType}:                          _GotoState265Action,
	{_State218, IdentifierToken}:                        _GotoState132Action,
	{_State218, StructToken}:                            _GotoState28Action,
	{_State218, EnumToken}:                              _GotoState111Action,
	{_State218, TraitToken}:                             _GotoState117Action,
	{_State218, FuncToken}:                              _GotoState114Action,
	{_State218, LparenToken}:                            _GotoState116Action,
	{_State218, QuestionToken}:                          _GotoState158Action,
	{_State218, ExclaimToken}:                           _GotoState112Action,
	{_State218, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State218, BitAndToken}:                            _GotoState110Action,
	{_State218, AtomTypeType}:                           _GotoState118Action,
	{_State218, TraitableTypeType}:                      _GotoState130Action,
	{_State218, TraitMulTypeType}:                       _GotoState128Action,
	{_State218, TraitAddTypeType}:                       _GotoState127Action,
	{_State218, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State218, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State218, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State218, TraitTypeType}:                          _GotoState129Action,
	{_State218, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State218, ValueTypeType}:                          _GotoState160Action,
	{_State218, ParameterTypeType}:                      _GotoState266Action,
	{_State218, FuncTypeType}:                           _GotoState123Action,
	{_State219, IdentifierToken}:                        _GotoState132Action,
	{_State219, StructToken}:                            _GotoState28Action,
	{_State219, EnumToken}:                              _GotoState111Action,
	{_State219, TraitToken}:                             _GotoState117Action,
	{_State219, FuncToken}:                              _GotoState114Action,
	{_State219, LparenToken}:                            _GotoState116Action,
	{_State219, QuestionToken}:                          _GotoState158Action,
	{_State219, DollarLbracketToken}:                    _GotoState61Action,
	{_State219, DotdotdotToken}:                         _GotoState157Action,
	{_State219, ExclaimToken}:                           _GotoState112Action,
	{_State219, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State219, BitAndToken}:                            _GotoState110Action,
	{_State219, OptionalGenericBindingType}:             _GotoState171Action,
	{_State219, AtomTypeType}:                           _GotoState118Action,
	{_State219, TraitableTypeType}:                      _GotoState130Action,
	{_State219, TraitMulTypeType}:                       _GotoState128Action,
	{_State219, TraitAddTypeType}:                       _GotoState127Action,
	{_State219, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State219, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State219, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State219, TraitTypeType}:                          _GotoState129Action,
	{_State219, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State219, ValueTypeType}:                          _GotoState160Action,
	{_State219, ParameterTypeType}:                      _GotoState159Action,
	{_State219, FuncTypeType}:                           _GotoState123Action,
	{_State220, RparenToken}:                            _GotoState267Action,
	{_State222, CommaToken}:                             _GotoState268Action,
	{_State225, IdentifierToken}:                        _GotoState115Action,
	{_State225, StructToken}:                            _GotoState28Action,
	{_State225, EnumToken}:                              _GotoState111Action,
	{_State225, TraitToken}:                             _GotoState117Action,
	{_State225, FuncToken}:                              _GotoState114Action,
	{_State225, LparenToken}:                            _GotoState116Action,
	{_State225, ExclaimToken}:                           _GotoState112Action,
	{_State225, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State225, BitAndToken}:                            _GotoState110Action,
	{_State225, AtomTypeType}:                           _GotoState118Action,
	{_State225, TraitableTypeType}:                      _GotoState130Action,
	{_State225, TraitMulTypeType}:                       _GotoState128Action,
	{_State225, TraitAddTypeType}:                       _GotoState127Action,
	{_State225, FieldDeclType}:                          _GotoState269Action,
	{_State225, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State225, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State225, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State225, TraitTypeType}:                          _GotoState129Action,
	{_State225, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State225, ValueTypeType}:                          _GotoState131Action,
	{_State225, FuncTypeType}:                           _GotoState123Action,
	{_State226, IdentifierToken}:                        _GotoState115Action,
	{_State226, StructToken}:                            _GotoState28Action,
	{_State226, EnumToken}:                              _GotoState111Action,
	{_State226, TraitToken}:                             _GotoState117Action,
	{_State226, FuncToken}:                              _GotoState114Action,
	{_State226, LparenToken}:                            _GotoState116Action,
	{_State226, ExclaimToken}:                           _GotoState112Action,
	{_State226, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State226, BitAndToken}:                            _GotoState110Action,
	{_State226, AtomTypeType}:                           _GotoState118Action,
	{_State226, TraitableTypeType}:                      _GotoState130Action,
	{_State226, TraitMulTypeType}:                       _GotoState128Action,
	{_State226, TraitAddTypeType}:                       _GotoState127Action,
	{_State226, FieldDeclType}:                          _GotoState270Action,
	{_State226, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State226, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State226, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State226, TraitTypeType}:                          _GotoState129Action,
	{_State226, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State226, ValueTypeType}:                          _GotoState131Action,
	{_State226, FuncTypeType}:                           _GotoState123Action,
	{_State231, MulToken}:                               _GotoState181Action,
	{_State232, MulToken}:                               _GotoState181Action,
	{_State236, IfToken}:                                _GotoState88Action,
	{_State236, LbraceToken}:                            _GotoState89Action,
	{_State236, BlockBodyType}:                          _GotoState271Action,
	{_State236, IfExprType}:                             _GotoState272Action,
	{_State237, LbracketToken}:                          _GotoState63Action,
	{_State237, DotToken}:                               _GotoState62Action,
	{_State237, DollarLbracketToken}:                    _GotoState61Action,
	{_State237, OptionalGenericBindingType}:             _GotoState65Action,
	{_State239, IdentifierToken}:                        _GotoState273Action,
	{_State253, IntegerLiteralToken}:                    _GotoState20Action,
	{_State253, FloatLiteralToken}:                      _GotoState17Action,
	{_State253, RuneLiteralToken}:                       _GotoState26Action,
	{_State253, StringLiteralToken}:                     _GotoState27Action,
	{_State253, IdentifierToken}:                        _GotoState19Action,
	{_State253, TrueToken}:                              _GotoState30Action,
	{_State253, FalseToken}:                             _GotoState16Action,
	{_State253, StructToken}:                            _GotoState28Action,
	{_State253, FuncToken}:                              _GotoState18Action,
	{_State253, LparenToken}:                            _GotoState23Action,
	{_State253, LabelDeclToken}:                         _GotoState21Action,
	{_State253, NotToken}:                               _GotoState25Action,
	{_State253, SubToken}:                               _GotoState29Action,
	{_State253, MulToken}:                               _GotoState24Action,
	{_State253, BitNegToken}:                            _GotoState15Action,
	{_State253, BitAndToken}:                            _GotoState14Action,
	{_State253, LexErrorToken}:                          _GotoState22Action,
	{_State253, LiteralType}:                            _GotoState41Action,
	{_State253, AnonymousStructExprType}:                _GotoState35Action,
	{_State253, AtomExprType}:                           _GotoState36Action,
	{_State253, CallExprType}:                           _GotoState38Action,
	{_State253, AccessExprType}:                         _GotoState31Action,
	{_State253, PostfixUnaryExprType}:                   _GotoState45Action,
	{_State253, PrefixUnaryOpType}:                      _GotoState47Action,
	{_State253, PrefixUnaryExprType}:                    _GotoState46Action,
	{_State253, MulExprType}:                            _GotoState42Action,
	{_State253, AddExprType}:                            _GotoState32Action,
	{_State253, CmpExprType}:                            _GotoState39Action,
	{_State253, AndExprType}:                            _GotoState33Action,
	{_State253, OrExprType}:                             _GotoState44Action,
	{_State253, SequenceExprType}:                       _GotoState48Action,
	{_State253, OptionalLabelDeclType}:                  _GotoState43Action,
	{_State253, BlockExprType}:                          _GotoState37Action,
	{_State253, ExpressionType}:                         _GotoState274Action,
	{_State253, ExplicitStructTypeType}:                 _GotoState40Action,
	{_State253, AnonymousFuncExprType}:                  _GotoState34Action,
	{_State255, IntegerLiteralToken}:                    _GotoState20Action,
	{_State255, FloatLiteralToken}:                      _GotoState17Action,
	{_State255, RuneLiteralToken}:                       _GotoState26Action,
	{_State255, StringLiteralToken}:                     _GotoState27Action,
	{_State255, IdentifierToken}:                        _GotoState19Action,
	{_State255, TrueToken}:                              _GotoState30Action,
	{_State255, FalseToken}:                             _GotoState16Action,
	{_State255, StructToken}:                            _GotoState28Action,
	{_State255, FuncToken}:                              _GotoState18Action,
	{_State255, LparenToken}:                            _GotoState23Action,
	{_State255, LabelDeclToken}:                         _GotoState21Action,
	{_State255, NotToken}:                               _GotoState25Action,
	{_State255, SubToken}:                               _GotoState29Action,
	{_State255, MulToken}:                               _GotoState24Action,
	{_State255, BitNegToken}:                            _GotoState15Action,
	{_State255, BitAndToken}:                            _GotoState14Action,
	{_State255, LexErrorToken}:                          _GotoState22Action,
	{_State255, LiteralType}:                            _GotoState41Action,
	{_State255, AnonymousStructExprType}:                _GotoState35Action,
	{_State255, AtomExprType}:                           _GotoState36Action,
	{_State255, CallExprType}:                           _GotoState38Action,
	{_State255, AccessExprType}:                         _GotoState31Action,
	{_State255, PostfixUnaryExprType}:                   _GotoState45Action,
	{_State255, PrefixUnaryOpType}:                      _GotoState47Action,
	{_State255, PrefixUnaryExprType}:                    _GotoState46Action,
	{_State255, MulExprType}:                            _GotoState42Action,
	{_State255, AddExprType}:                            _GotoState32Action,
	{_State255, CmpExprType}:                            _GotoState39Action,
	{_State255, AndExprType}:                            _GotoState33Action,
	{_State255, OrExprType}:                             _GotoState44Action,
	{_State255, SequenceExprType}:                       _GotoState48Action,
	{_State255, OptionalLabelDeclType}:                  _GotoState43Action,
	{_State255, BlockExprType}:                          _GotoState37Action,
	{_State255, ExpressionType}:                         _GotoState275Action,
	{_State255, ExplicitStructTypeType}:                 _GotoState40Action,
	{_State255, AnonymousFuncExprType}:                  _GotoState34Action,
	{_State257, IntegerLiteralToken}:                    _GotoState20Action,
	{_State257, FloatLiteralToken}:                      _GotoState17Action,
	{_State257, RuneLiteralToken}:                       _GotoState26Action,
	{_State257, StringLiteralToken}:                     _GotoState27Action,
	{_State257, IdentifierToken}:                        _GotoState19Action,
	{_State257, TrueToken}:                              _GotoState30Action,
	{_State257, FalseToken}:                             _GotoState16Action,
	{_State257, StructToken}:                            _GotoState28Action,
	{_State257, FuncToken}:                              _GotoState18Action,
	{_State257, LparenToken}:                            _GotoState23Action,
	{_State257, LabelDeclToken}:                         _GotoState21Action,
	{_State257, NotToken}:                               _GotoState25Action,
	{_State257, SubToken}:                               _GotoState29Action,
	{_State257, MulToken}:                               _GotoState24Action,
	{_State257, BitNegToken}:                            _GotoState15Action,
	{_State257, BitAndToken}:                            _GotoState14Action,
	{_State257, LexErrorToken}:                          _GotoState22Action,
	{_State257, LiteralType}:                            _GotoState41Action,
	{_State257, AnonymousStructExprType}:                _GotoState35Action,
	{_State257, AtomExprType}:                           _GotoState36Action,
	{_State257, CallExprType}:                           _GotoState38Action,
	{_State257, AccessExprType}:                         _GotoState31Action,
	{_State257, PostfixUnaryExprType}:                   _GotoState45Action,
	{_State257, PrefixUnaryOpType}:                      _GotoState47Action,
	{_State257, PrefixUnaryExprType}:                    _GotoState46Action,
	{_State257, MulExprType}:                            _GotoState42Action,
	{_State257, AddExprType}:                            _GotoState32Action,
	{_State257, CmpExprType}:                            _GotoState39Action,
	{_State257, AndExprType}:                            _GotoState33Action,
	{_State257, OrExprType}:                             _GotoState44Action,
	{_State257, SequenceExprType}:                       _GotoState48Action,
	{_State257, OptionalExpressionOrImplicitStructType}: _GotoState277Action,
	{_State257, ExpressionOrImplicitStructType}:         _GotoState276Action,
	{_State257, OptionalLabelDeclType}:                  _GotoState43Action,
	{_State257, BlockExprType}:                          _GotoState37Action,
	{_State257, ExpressionType}:                         _GotoState197Action,
	{_State257, ExplicitStructTypeType}:                 _GotoState40Action,
	{_State257, AnonymousFuncExprType}:                  _GotoState34Action,
	{_State260, RbraceToken}:                            _GotoState278Action,
	{_State264, RparenToken}:                            _GotoState279Action,
	{_State267, IdentifierToken}:                        _GotoState132Action,
	{_State267, StructToken}:                            _GotoState28Action,
	{_State267, EnumToken}:                              _GotoState111Action,
	{_State267, TraitToken}:                             _GotoState117Action,
	{_State267, FuncToken}:                              _GotoState114Action,
	{_State267, LparenToken}:                            _GotoState116Action,
	{_State267, QuestionToken}:                          _GotoState213Action,
	{_State267, ExclaimToken}:                           _GotoState112Action,
	{_State267, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State267, BitAndToken}:                            _GotoState110Action,
	{_State267, AtomTypeType}:                           _GotoState118Action,
	{_State267, TraitableTypeType}:                      _GotoState130Action,
	{_State267, TraitMulTypeType}:                       _GotoState128Action,
	{_State267, TraitAddTypeType}:                       _GotoState127Action,
	{_State267, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State267, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State267, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State267, TraitTypeType}:                          _GotoState129Action,
	{_State267, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State267, ValueTypeType}:                          _GotoState215Action,
	{_State267, ReturnTypeType}:                         _GotoState280Action,
	{_State267, FuncTypeType}:                           _GotoState123Action,
	{_State268, IdentifierToken}:                        _GotoState219Action,
	{_State268, StructToken}:                            _GotoState28Action,
	{_State268, EnumToken}:                              _GotoState111Action,
	{_State268, TraitToken}:                             _GotoState117Action,
	{_State268, FuncToken}:                              _GotoState114Action,
	{_State268, LparenToken}:                            _GotoState116Action,
	{_State268, QuestionToken}:                          _GotoState158Action,
	{_State268, DotdotdotToken}:                         _GotoState218Action,
	{_State268, ExclaimToken}:                           _GotoState112Action,
	{_State268, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State268, BitAndToken}:                            _GotoState110Action,
	{_State268, AtomTypeType}:                           _GotoState118Action,
	{_State268, TraitableTypeType}:                      _GotoState130Action,
	{_State268, TraitMulTypeType}:                       _GotoState128Action,
	{_State268, TraitAddTypeType}:                       _GotoState127Action,
	{_State268, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State268, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State268, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State268, TraitTypeType}:                          _GotoState129Action,
	{_State268, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State268, ValueTypeType}:                          _GotoState160Action,
	{_State268, ParameterTypeType}:                      _GotoState224Action,
	{_State268, ParameterDeclType}:                      _GotoState281Action,
	{_State268, FuncTypeType}:                           _GotoState123Action,
	{_State268, ParameterDefType}:                       _GotoState223Action,
	{_State273, GreaterToken}:                           _GotoState282Action,
	{_State276, CommaToken}:                             _GotoState255Action,
	{_State279, IdentifierToken}:                        _GotoState132Action,
	{_State279, StructToken}:                            _GotoState28Action,
	{_State279, EnumToken}:                              _GotoState111Action,
	{_State279, TraitToken}:                             _GotoState117Action,
	{_State279, FuncToken}:                              _GotoState114Action,
	{_State279, LparenToken}:                            _GotoState116Action,
	{_State279, QuestionToken}:                          _GotoState213Action,
	{_State279, ExclaimToken}:                           _GotoState112Action,
	{_State279, ExclaimExclaimToken}:                    _GotoState113Action,
	{_State279, BitAndToken}:                            _GotoState110Action,
	{_State279, AtomTypeType}:                           _GotoState118Action,
	{_State279, TraitableTypeType}:                      _GotoState130Action,
	{_State279, TraitMulTypeType}:                       _GotoState128Action,
	{_State279, TraitAddTypeType}:                       _GotoState127Action,
	{_State279, ImplicitStructTypeType}:                 _GotoState125Action,
	{_State279, ImplicitEnumTypeType}:                   _GotoState124Action,
	{_State279, ExplicitStructTypeType}:                 _GotoState120Action,
	{_State279, TraitTypeType}:                          _GotoState129Action,
	{_State279, ExplicitEnumTypeType}:                   _GotoState119Action,
	{_State279, ValueTypeType}:                          _GotoState215Action,
	{_State279, ReturnTypeType}:                         _GotoState283Action,
	{_State279, FuncTypeType}:                           _GotoState123Action,
	{_State282, StringLiteralToken}:                     _GotoState284Action,
	{_State283, LbraceToken}:                            _GotoState89Action,
	{_State283, BlockBodyType}:                          _GotoState285Action,
	{_State4, _WildcardMarker}:                          _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State13, IdentifierToken}:                         _ReduceNilToOptionalReceiverAction,
	{_State14, _WildcardMarker}:                         _ReduceBitAndToPrefixUnaryOpAction,
	{_State15, _WildcardMarker}:                         _ReduceBitNegToPrefixUnaryOpAction,
	{_State16, _WildcardMarker}:                         _ReduceFalseToLiteralAction,
	{_State17, _WildcardMarker}:                         _ReduceFloatLiteralToLiteralAction,
	{_State19, _WildcardMarker}:                         _ReduceIdentifierToAtomExprAction,
	{_State20, _WildcardMarker}:                         _ReduceIntegerLiteralToLiteralAction,
	{_State21, _WildcardMarker}:                         _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State22, _WildcardMarker}:                         _ReduceLexErrorToAtomExprAction,
	{_State23, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State24, _WildcardMarker}:                         _ReduceMulToPrefixUnaryOpAction,
	{_State25, _WildcardMarker}:                         _ReduceNotToPrefixUnaryOpAction,
	{_State26, _WildcardMarker}:                         _ReduceRuneLiteralToLiteralAction,
	{_State27, _WildcardMarker}:                         _ReduceStringLiteralToLiteralAction,
	{_State29, _WildcardMarker}:                         _ReduceSubToPrefixUnaryOpAction,
	{_State30, _WildcardMarker}:                         _ReduceTrueToLiteralAction,
	{_State31, _WildcardMarker}:                         _ReduceAccessExprToPostfixUnaryExprAction,
	{_State31, LparenToken}:                             _ReduceNilToOptionalGenericBindingAction,
	{_State32, _WildcardMarker}:                         _ReduceAddExprToCmpExprAction,
	{_State33, _WildcardMarker}:                         _ReduceAndExprToOrExprAction,
	{_State34, _WildcardMarker}:                         _ReduceAnonymousFuncExprToAtomExprAction,
	{_State35, _WildcardMarker}:                         _ReduceAnonymousStructExprToAtomExprAction,
	{_State36, _WildcardMarker}:                         _ReduceAtomExprToAccessExprAction,
	{_State37, _WildcardMarker}:                         _ReduceBlockExprToAtomExprAction,
	{_State38, _WildcardMarker}:                         _ReduceCallExprToAccessExprAction,
	{_State39, _WildcardMarker}:                         _ReduceCmpExprToAndExprAction,
	{_State41, _WildcardMarker}:                         _ReduceLiteralToAtomExprAction,
	{_State42, _WildcardMarker}:                         _ReduceMulExprToAddExprAction,
	{_State44, _EndMarker}:                              _ReduceToSequenceExprAction,
	{_State45, _WildcardMarker}:                         _ReducePostfixUnaryExprToPrefixUnaryExprAction,
	{_State46, _WildcardMarker}:                         _ReducePrefixUnaryExprToMulExprAction,
	{_State47, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State48, _EndMarker}:                              _ReduceSequenceExprToExpressionAction,
	{_State49, _EndMarker}:                              _ReduceCommentToLexInternalTokensAction,
	{_State50, _EndMarker}:                              _ReduceSpacesToLexInternalTokensAction,
	{_State51, _EndMarker}:                              _ReduceNoSpecToPackageDefAction,
	{_State52, _WildcardMarker}:                         _ReduceNilToOptionalGenericParametersAction,
	{_State55, RparenToken}:                             _ReduceNilToOptionalParameterDefsAction,
	{_State56, _WildcardMarker}:                         _ReduceArgumentToArgumentsAction,
	{_State58, _WildcardMarker}:                         _ReducePositionalToArgumentAction,
	{_State59, RparenToken}:                             _ReduceNilToOptionalFieldDeclsAction,
	{_State60, _WildcardMarker}:                         _ReduceToExplicitStructTypeAction,
	{_State61, RbracketToken}:                           _ReduceNilToOptionalGenericArgumentsAction,
	{_State63, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State64, _WildcardMarker}:                         _ReduceQuestionToPostfixUnaryExprAction,
	{_State66, _WildcardMarker}:                         _ReduceAddToAddOpAction,
	{_State67, _WildcardMarker}:                         _ReduceBitOrToAddOpAction,
	{_State68, _WildcardMarker}:                         _ReduceBitXorToAddOpAction,
	{_State69, _WildcardMarker}:                         _ReduceSubToAddOpAction,
	{_State70, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State71, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State72, _WildcardMarker}:                         _ReduceEqualToCmpOpAction,
	{_State73, _WildcardMarker}:                         _ReduceGreaterToCmpOpAction,
	{_State74, _WildcardMarker}:                         _ReduceGreaterOrEqualToCmpOpAction,
	{_State75, _WildcardMarker}:                         _ReduceLessToCmpOpAction,
	{_State76, _WildcardMarker}:                         _ReduceLessOrEqualToCmpOpAction,
	{_State77, _WildcardMarker}:                         _ReduceNotEqualToCmpOpAction,
	{_State78, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State79, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State80, _WildcardMarker}:                         _ReduceBitAndToMulOpAction,
	{_State81, _WildcardMarker}:                         _ReduceBitLshiftToMulOpAction,
	{_State82, _WildcardMarker}:                         _ReduceBitRshiftToMulOpAction,
	{_State83, _WildcardMarker}:                         _ReduceDivToMulOpAction,
	{_State84, _WildcardMarker}:                         _ReduceModToMulOpAction,
	{_State85, _WildcardMarker}:                         _ReduceMulToMulOpAction,
	{_State86, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State87, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State88, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State89, _WildcardMarker}:                         _ReduceEmptyListToStatementsAction,
	{_State91, _EndMarker}:                              _ReduceToBlockExprAction,
	{_State92, _EndMarker}:                              _ReduceIfExprToExpressionAction,
	{_State93, _EndMarker}:                              _ReduceLoopExprToExpressionAction,
	{_State94, _EndMarker}:                              _ReduceSwitchExprToExpressionAction,
	{_State95, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State97, _WildcardMarker}:                         _ReducePrefixOpToPrefixUnaryExprAction,
	{_State98, _WildcardMarker}:                         _ReduceEmptyListToPackageStatementsAction,
	{_State99, RbracketToken}:                           _ReduceNilToOptionalGenericParameterDefsAction,
	{_State104, LparenToken}:                            _ReduceNilToOptionalGenericParametersAction,
	{_State106, _WildcardMarker}:                        _ReduceParameterDefToParameterDefsAction,
	{_State107, RparenToken}:                            _ReduceParameterDefsToOptionalParameterDefsAction,
	{_State108, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State109, _WildcardMarker}:                        _ReduceImplicitToAnonymousStructExprAction,
	{_State115, _WildcardMarker}:                        _ReduceNilToOptionalGenericBindingAction,
	{_State116, RparenToken}:                            _ReduceNilToOptionalFieldDeclsAction,
	{_State118, _WildcardMarker}:                        _ReduceAtomTypeToTraitableTypeAction,
	{_State119, _WildcardMarker}:                        _ReduceExplicitEnumTypeToAtomTypeAction,
	{_State120, _WildcardMarker}:                        _ReduceExplicitStructTypeToAtomTypeAction,
	{_State121, _WildcardMarker}:                        _ReduceFieldDeclToFieldDeclsAction,
	{_State122, RparenToken}:                            _ReduceFieldDeclsToOptionalFieldDeclsAction,
	{_State123, _EndMarker}:                             _ReduceFuncTypeToValueTypeAction,
	{_State124, _WildcardMarker}:                        _ReduceImplicitEnumTypeToAtomTypeAction,
	{_State125, _WildcardMarker}:                        _ReduceImplicitStructTypeToAtomTypeAction,
	{_State127, _EndMarker}:                             _ReduceTraitAddTypeToValueTypeAction,
	{_State128, _WildcardMarker}:                        _ReduceTraitMulTypeToTraitAddTypeAction,
	{_State129, _WildcardMarker}:                        _ReduceTraitTypeToAtomTypeAction,
	{_State130, _WildcardMarker}:                        _ReduceTraitableTypeToTraitMulTypeAction,
	{_State131, _WildcardMarker}:                        _ReduceImplicitToFieldDeclAction,
	{_State132, _WildcardMarker}:                        _ReduceNilToOptionalGenericBindingAction,
	{_State133, RbracketToken}:                          _ReduceGenericArgumentsToOptionalGenericArgumentsAction,
	{_State135, _WildcardMarker}:                        _ReduceValueTypeToGenericArgumentsAction,
	{_State136, _WildcardMarker}:                        _ReduceAccessToAccessExprAction,
	{_State138, RparenToken}:                            _ReduceNilToOptionalArgumentsAction,
	{_State138, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State139, _WildcardMarker}:                        _ReduceOpToAddExprAction,
	{_State140, _WildcardMarker}:                        _ReduceOpToAndExprAction,
	{_State141, _WildcardMarker}:                        _ReduceOpToCmpExprAction,
	{_State143, _WildcardMarker}:                        _ReduceOpToMulExprAction,
	{_State144, _WildcardMarker}:                        _ReduceBlockExprToAtomExprAction,
	{_State144, _EndMarker}:                             _ReduceInfiniteToLoopExprAction,
	{_State145, LbraceToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State147, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State149, _WildcardMarker}:                        _ReduceOpToOrExprAction,
	{_State151, _WildcardMarker}:                        _ReduceUnconstrainedToGenericParameterDefAction,
	{_State152, _WildcardMarker}:                        _ReduceGenericParameterDefToGenericParameterDefsAction,
	{_State153, RbracketToken}:                          _ReduceGenericParameterDefsToOptionalGenericParameterDefsAction,
	{_State155, _EndMarker}:                             _ReduceAliasToTypeDefAction,
	{_State156, _EndMarker}:                             _ReduceDefinitionToTypeDefAction,
	{_State158, _WildcardMarker}:                        _ReduceQuestionToParameterTypeAction,
	{_State159, _WildcardMarker}:                        _ReduceArgToParameterDefAction,
	{_State160, _WildcardMarker}:                        _ReduceValueTypeToParameterTypeAction,
	{_State161, IdentifierToken}:                        _ReduceReceiverToOptionalReceiverAction,
	{_State163, LbraceToken}:                            _ReduceNilToReturnTypeAction,
	{_State165, _WildcardMarker}:                        _ReduceAddToArgumentsAction,
	{_State166, _EndMarker}:                             _ReduceReferenceToValueTypeAction,
	{_State168, _WildcardMarker}:                        _ReduceMethodInterfaceToTraitableTypeAction,
	{_State169, _WildcardMarker}:                        _ReduceTraitToTraitableTypeAction,
	{_State170, RparenToken}:                            _ReduceNilToOptionalParameterDeclsAction,
	{_State171, _WildcardMarker}:                        _ReduceNamedToAtomTypeAction,
	{_State172, _WildcardMarker}:                        _ReduceExplicitToFieldDeclAction,
	{_State173, _WildcardMarker}:                        _ReduceFieldDeclToFieldDeclsAction,
	{_State178, _WildcardMarker}:                        _ReduceToImplicitStructTypeAction,
	{_State183, _WildcardMarker}:                        _ReduceBindingToOptionalGenericBindingAction,
	{_State184, _WildcardMarker}:                        _ReduceIndexToAccessExprAction,
	{_State185, RparenToken}:                            _ReduceArgumentsToOptionalArgumentsAction,
	{_State187, _WildcardMarker}:                        _ReduceExplicitToAnonymousStructExprAction,
	{_State188, _EndMarker}:                             _ReduceWhileToLoopExprAction,
	{_State189, _WildcardMarker}:                        _ReduceNoElseToIfExprAction,
	{_State190, LbraceToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State191, _WildcardMarker}:                        _ReduceBreakToJumpTypeAction,
	{_State192, _WildcardMarker}:                        _ReduceContinueToJumpTypeAction,
	{_State193, _EndMarker}:                             _ReduceToBlockBodyAction,
	{_State194, _WildcardMarker}:                        _ReduceReturnToJumpTypeAction,
	{_State196, _WildcardMarker}:                        _ReduceAccessExprToPostfixUnaryExprAction,
	{_State196, LparenToken}:                            _ReduceNilToOptionalGenericBindingAction,
	{_State197, _WildcardMarker}:                        _ReduceExpressionToExpressionOrImplicitStructAction,
	{_State198, _WildcardMarker}:                        _ReduceExpressionOrImplicitStructToStatementBodyAction,
	{_State199, _WildcardMarker}:                        _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State200, _WildcardMarker}:                        _ReduceAddToStatementsAction,
	{_State202, _WildcardMarker}:                        _ReduceUnsafeStatementToStatementBodyAction,
	{_State204, _EndMarker}:                             _ReduceWithSpecToPackageDefAction,
	{_State205, _WildcardMarker}:                        _ReduceAddToPackageStatementsAction,
	{_State207, _WildcardMarker}:                        _ReduceToPackageStatementBodyAction,
	{_State208, _WildcardMarker}:                        _ReduceConstrainedToGenericParameterDefAction,
	{_State210, _WildcardMarker}:                        _ReduceGenericToOptionalGenericParametersAction,
	{_State211, _WildcardMarker}:                        _ReduceVarargToParameterDefAction,
	{_State212, RparenToken}:                            _ReduceNilToOptionalParameterDefsAction,
	{_State213, _EndMarker}:                             _ReduceQuestionToReturnTypeAction,
	{_State215, _EndMarker}:                             _ReduceValueTypeToReturnTypeAction,
	{_State216, _WildcardMarker}:                        _ReduceAddToParameterDefsAction,
	{_State217, _WildcardMarker}:                        _ReduceToExplicitEnumTypeAction,
	{_State219, _WildcardMarker}:                        _ReduceNilToOptionalGenericBindingAction,
	{_State221, _WildcardMarker}:                        _ReduceParameterDeclToParameterDeclsAction,
	{_State222, RparenToken}:                            _ReduceParameterDeclsToOptionalParameterDeclsAction,
	{_State223, _WildcardMarker}:                        _ReduceParameterDefToParameterDeclAction,
	{_State224, _WildcardMarker}:                        _ReduceUnamedToParameterDeclAction,
	{_State227, _WildcardMarker}:                        _ReduceToImplicitEnumTypeAction,
	{_State228, _WildcardMarker}:                        _ReduceToTraitTypeAction,
	{_State229, _WildcardMarker}:                        _ReduceExplicitToFieldDeclsAction,
	{_State230, _WildcardMarker}:                        _ReduceImplicitToFieldDeclsAction,
	{_State231, _WildcardMarker}:                        _ReduceUnionToTraitAddTypeAction,
	{_State232, _WildcardMarker}:                        _ReduceDifferenceToTraitAddTypeAction,
	{_State233, _WildcardMarker}:                        _ReduceIntersectToTraitMulTypeAction,
	{_State234, _WildcardMarker}:                        _ReduceAddToGenericArgumentsAction,
	{_State235, _WildcardMarker}:                        _ReduceToCallExprAction,
	{_State237, LparenToken}:                            _ReduceNilToOptionalGenericBindingAction,
	{_State238, _WildcardMarker}:                        _ReduceCallExprToAccessExprAction,
	{_State238, NewlinesToken}:                          _ReduceAsyncToStatementBodyAction,
	{_State238, SemicolonToken}:                         _ReduceAsyncToStatementBodyAction,
	{_State240, _WildcardMarker}:                        _ReduceAddAssignToBinaryOpAssignAction,
	{_State241, _WildcardMarker}:                        _ReduceAddOneAssignToOpOneAssignAction,
	{_State242, _WildcardMarker}:                        _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State243, _WildcardMarker}:                        _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State244, _WildcardMarker}:                        _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State245, _WildcardMarker}:                        _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State246, _WildcardMarker}:                        _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State247, _WildcardMarker}:                        _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State248, _WildcardMarker}:                        _ReduceDivAssignToBinaryOpAssignAction,
	{_State249, _WildcardMarker}:                        _ReduceModAssignToBinaryOpAssignAction,
	{_State250, _WildcardMarker}:                        _ReduceMulAssignToBinaryOpAssignAction,
	{_State251, _WildcardMarker}:                        _ReduceSubAssignToBinaryOpAssignAction,
	{_State252, _WildcardMarker}:                        _ReduceSubOneAssignToOpOneAssignAction,
	{_State253, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State254, _WildcardMarker}:                        _ReduceOpOneAssignToStatementBodyAction,
	{_State255, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State256, _WildcardMarker}:                        _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State257, NewlinesToken}:                          _ReduceNilToOptionalExpressionOrImplicitStructAction,
	{_State257, SemicolonToken}:                         _ReduceNilToOptionalExpressionOrImplicitStructAction,
	{_State257, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State258, _WildcardMarker}:                        _ReduceImplicitToStatementAction,
	{_State259, _WildcardMarker}:                        _ReduceExplicitToStatementAction,
	{_State261, _WildcardMarker}:                        _ReduceImplicitToPackageStatementAction,
	{_State262, _WildcardMarker}:                        _ReduceExplicitToPackageStatementAction,
	{_State263, _WildcardMarker}:                        _ReduceAddToGenericParameterDefsAction,
	{_State265, _WildcardMarker}:                        _ReduceToAnonymousFuncExprAction,
	{_State266, _WildcardMarker}:                        _ReduceUnnamedVarargToParameterDeclAction,
	{_State267, _WildcardMarker}:                        _ReduceNilToReturnTypeAction,
	{_State269, _WildcardMarker}:                        _ReducePairToFieldDeclOrListAction,
	{_State270, _WildcardMarker}:                        _ReduceAddToFieldDeclOrListAction,
	{_State271, _EndMarker}:                             _ReduceIfElseToIfExprAction,
	{_State272, _EndMarker}:                             _ReduceMultiIfElseToIfExprAction,
	{_State274, _WildcardMarker}:                        _ReduceBinaryOpAssignToStatementBodyAction,
	{_State275, _WildcardMarker}:                        _ReduceImplicitStructToExpressionOrImplicitStructAction,
	{_State276, _WildcardMarker}:                        _ReduceExpressionOrImplicitStructToOptionalExpressionOrImplicitStructAction,
	{_State277, _WildcardMarker}:                        _ReduceJumpToStatementBodyAction,
	{_State278, _EndMarker}:                             _ReduceToSwitchExprAction,
	{_State279, LbraceToken}:                            _ReduceNilToReturnTypeAction,
	{_State280, _EndMarker}:                             _ReduceToFuncTypeAction,
	{_State281, _WildcardMarker}:                        _ReduceAddToParameterDeclsAction,
	{_State284, _WildcardMarker}:                        _ReduceToUnsafeStatementAction,
	{_State285, _EndMarker}:                             _ReduceToNamedFuncDefAction,
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.package_def
    Reduce:
      (nil)
    Goto:
      PACKAGE -> State 11
      package_def -> State 6

  State 2:
    Kernel Items:
      #accept: ^.type_def
    Reduce:
      (nil)
    Goto:
      TYPE -> State 12
      type_def -> State 7

  State 3:
    Kernel Items:
      #accept: ^.named_func_def
    Reduce:
      (nil)
    Goto:
      FUNC -> State 13
      named_func_def -> State 8

  State 4:
    Kernel Items:
      #accept: ^.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 42
      add_expr -> State 32
      cmp_expr -> State 39
      and_expr -> State 33
      or_expr -> State 44
      sequence_expr -> State 48
      optional_label_decl -> State 43
      block_expr -> State 37
      expression -> State 9
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 5:
    Kernel Items:
      #accept: ^.lex_internal_tokens
    Reduce:
      (nil)
    Goto:
      SPACES -> State 50
      COMMENT -> State 49
      lex_internal_tokens -> State 10

  State 6:
    Kernel Items:
      #accept: ^ package_def., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 7:
    Kernel Items:
      #accept: ^ type_def., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      #accept: ^ named_func_def., $
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
      package_def: PACKAGE.IDENTIFIER
      package_def: PACKAGE.IDENTIFIER LPAREN package_statements RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 51

  State 12:
    Kernel Items:
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_def: TYPE.IDENTIFIER EQUAL value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 52

  State 13:
    Kernel Items:
      named_func_def: FUNC.optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      LPAREN -> State 53
      optional_receiver -> State 54

  State 14:
    Kernel Items:
      prefix_unary_op: BIT_AND., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 15:
    Kernel Items:
      prefix_unary_op: BIT_NEG., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 16:
    Kernel Items:
      literal: FALSE., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 17:
    Kernel Items:
      literal: FLOAT_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 18:
    Kernel Items:
      anonymous_func_expr: FUNC.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 55

  State 19:
    Kernel Items:
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 20:
    Kernel Items:
      literal: INTEGER_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      optional_label_decl: LABEL_DECL., *
    Reduce:
      * -> [optional_label_decl]
    Goto:
      (nil)

  State 22:
    Kernel Items:
      atom_expr: LEX_ERROR., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      anonymous_struct_expr: LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      argument -> State 56
      arguments -> State 57
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 42
      add_expr -> State 32
      cmp_expr -> State 39
      and_expr -> State 33
      or_expr -> State 44
      sequence_expr -> State 48
      optional_label_decl -> State 43
      block_expr -> State 37
      expression -> State 58
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 24:
    Kernel Items:
      prefix_unary_op: MUL., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 25:
    Kernel Items:
      prefix_unary_op: NOT., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      literal: RUNE_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 27:
    Kernel Items:
      literal: STRING_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 28:
    Kernel Items:
      explicit_struct_type: STRUCT.implicit_struct_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 59
      implicit_struct_type -> State 60

  State 29:
    Kernel Items:
      prefix_unary_op: SUB., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 30:
    Kernel Items:
      literal: TRUE., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 31:
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
      LBRACKET -> State 63
      DOT -> State 62
      QUESTION -> State 64
      DOLLAR_LBRACKET -> State 61
      optional_generic_binding -> State 65

  State 32:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 66
      SUB -> State 69
      BIT_XOR -> State 68
      BIT_OR -> State 67
      add_op -> State 70

  State 33:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 71

  State 34:
    Kernel Items:
      atom_expr: anonymous_func_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 35:
    Kernel Items:
      atom_expr: anonymous_struct_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 36:
    Kernel Items:
      access_expr: atom_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 37:
    Kernel Items:
      atom_expr: block_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 38:
    Kernel Items:
      access_expr: call_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 39:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 72
      NOT_EQUAL -> State 77
      LESS -> State 75
      LESS_OR_EQUAL -> State 76
      GREATER -> State 73
      GREATER_OR_EQUAL -> State 74
      cmp_op -> State 78

  State 40:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 79

  State 41:
    Kernel Items:
      atom_expr: literal., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 42:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 85
      DIV -> State 83
      MOD -> State 84
      BIT_AND -> State 80
      BIT_LSHIFT -> State 81
      BIT_RSHIFT -> State 82
      mul_op -> State 86

  State 43:
    Kernel Items:
      block_expr: optional_label_decl.block_body
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 88
      SWITCH -> State 90
      FOR -> State 87
      LBRACE -> State 89
      block_body -> State 91
      if_expr -> State 92
      switch_expr -> State 94
      loop_expr -> State 93

  State 44:
    Kernel Items:
      or_expr: or_expr.OR and_expr
      sequence_expr: or_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      OR -> State 95

  State 45:
    Kernel Items:
      prefix_unary_expr: postfix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 46:
    Kernel Items:
      mul_expr: prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 47:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op.prefix_unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 97
      optional_label_decl -> State 96
      block_expr -> State 37
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 48:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 49:
    Kernel Items:
      lex_internal_tokens: COMMENT., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 50:
    Kernel Items:
      lex_internal_tokens: SPACES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 51:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER., $
      package_def: PACKAGE IDENTIFIER.LPAREN package_statements RPAREN
    Reduce:
      $ -> [package_def]
    Goto:
      LPAREN -> State 98

  State 52:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_def: TYPE IDENTIFIER.EQUAL value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 99
      EQUAL -> State 100
      optional_generic_parameters -> State 101

  State 53:
    Kernel Items:
      optional_receiver: LPAREN.parameter_def RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 102
      parameter_def -> State 103

  State 54:
    Kernel Items:
      named_func_def: FUNC optional_receiver.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 104

  State 55:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 102
      parameter_def -> State 106
      parameter_defs -> State 107
      optional_parameter_defs -> State 105

  State 56:
    Kernel Items:
      arguments: argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 57:
    Kernel Items:
      anonymous_struct_expr: LPAREN arguments.RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 109
      COMMA -> State 108

  State 58:
    Kernel Items:
      argument: expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 59:
    Kernel Items:
      implicit_struct_type: LPAREN.optional_field_decls RPAREN
    Reduce:
      RPAREN -> [optional_field_decls]
    Goto:
      IDENTIFIER -> State 115
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      field_decl -> State 121
      field_decls -> State 122
      optional_field_decls -> State 126
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 131
      func_type -> State 123

  State 60:
    Kernel Items:
      explicit_struct_type: STRUCT implicit_struct_type., *
    Reduce:
      * -> [explicit_struct_type]
    Goto:
      (nil)

  State 61:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      generic_arguments -> State 133
      optional_generic_arguments -> State 134
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 135
      func_type -> State 123

  State 62:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 136

  State 63:
    Kernel Items:
      access_expr: access_expr LBRACKET.expression RBRACKET
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 42
      add_expr -> State 32
      cmp_expr -> State 39
      and_expr -> State 33
      or_expr -> State 44
      sequence_expr -> State 48
      optional_label_decl -> State 43
      block_expr -> State 37
      expression -> State 137
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 64:
    Kernel Items:
      postfix_unary_expr: access_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      call_expr: access_expr optional_generic_binding.LPAREN optional_arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 138

  State 66:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 67:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 68:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 69:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 70:
    Kernel Items:
      add_expr: add_expr add_op.mul_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 139
      optional_label_decl -> State 96
      block_expr -> State 37
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 71:
    Kernel Items:
      and_expr: and_expr AND.cmp_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 42
      add_expr -> State 32
      cmp_expr -> State 140
      optional_label_decl -> State 96
      block_expr -> State 37
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 72:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 73:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 74:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 75:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 76:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 77:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 78:
    Kernel Items:
      cmp_expr: cmp_expr cmp_op.add_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 42
      add_expr -> State 141
      optional_label_decl -> State 96
      block_expr -> State 37
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 79:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      argument -> State 56
      arguments -> State 142
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 42
      add_expr -> State 32
      cmp_expr -> State 39
      and_expr -> State 33
      or_expr -> State 44
      sequence_expr -> State 48
      optional_label_decl -> State 43
      block_expr -> State 37
      expression -> State 58
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 80:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 81:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 82:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 84:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 85:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 86:
    Kernel Items:
      mul_expr: mul_expr mul_op.prefix_unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 143
      optional_label_decl -> State 96
      block_expr -> State 37
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 87:
    Kernel Items:
      loop_expr: FOR.block_expr
      loop_expr: FOR.sequence_expr block_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 42
      add_expr -> State 32
      cmp_expr -> State 39
      and_expr -> State 33
      or_expr -> State 44
      sequence_expr -> State 145
      optional_label_decl -> State 96
      block_expr -> State 144
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 88:
    Kernel Items:
      if_expr: IF.sequence_expr block_body
      if_expr: IF.sequence_expr block_body ELSE block_body
      if_expr: IF.sequence_expr block_body ELSE if_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 42
      add_expr -> State 32
      cmp_expr -> State 39
      and_expr -> State 33
      or_expr -> State 44
      sequence_expr -> State 146
      optional_label_decl -> State 96
      block_expr -> State 37
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 89:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 147

  State 90:
    Kernel Items:
      switch_expr: SWITCH.LBRACE CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 148

  State 91:
    Kernel Items:
      block_expr: optional_label_decl block_body., $
    Reduce:
      $ -> [block_expr]
    Goto:
      (nil)

  State 92:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 93:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 94:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 95:
    Kernel Items:
      or_expr: or_expr OR.and_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 42
      add_expr -> State 32
      cmp_expr -> State 39
      and_expr -> State 149
      optional_label_decl -> State 96
      block_expr -> State 37
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 96:
    Kernel Items:
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 89
      block_body -> State 91

  State 97:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 98:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN.package_statements RPAREN
    Reduce:
      * -> [package_statements]
    Goto:
      package_statements -> State 150

  State 99:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.optional_generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 151
      generic_parameter_def -> State 152
      generic_parameter_defs -> State 153
      optional_generic_parameter_defs -> State 154

  State 100:
    Kernel Items:
      type_def: TYPE IDENTIFIER EQUAL.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 155
      func_type -> State 123

  State 101:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 156
      func_type -> State 123

  State 102:
    Kernel Items:
      parameter_def: IDENTIFIER.parameter_type
      parameter_def: IDENTIFIER.DOTDOTDOT parameter_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      QUESTION -> State 158
      DOTDOTDOT -> State 157
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 160
      parameter_type -> State 159
      func_type -> State 123

  State 103:
    Kernel Items:
      optional_receiver: LPAREN parameter_def.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 161

  State 104:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 99
      optional_generic_parameters -> State 162

  State 105:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 163

  State 106:
    Kernel Items:
      parameter_defs: parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 107:
    Kernel Items:
      parameter_defs: parameter_defs.COMMA parameter_def
      optional_parameter_defs: parameter_defs., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      COMMA -> State 164

  State 108:
    Kernel Items:
      arguments: arguments COMMA.argument
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      argument -> State 165
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 42
      add_expr -> State 32
      cmp_expr -> State 39
      and_expr -> State 33
      or_expr -> State 44
      sequence_expr -> State 48
      optional_label_decl -> State 43
      block_expr -> State 37
      expression -> State 58
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 109:
    Kernel Items:
      anonymous_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 110:
    Kernel Items:
      value_type: BIT_AND.trait_add_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      LPAREN -> State 116
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 166
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119

  State 111:
    Kernel Items:
      explicit_enum_type: ENUM.LPAREN RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 167

  State 112:
    Kernel Items:
      traitable_type: EXCLAIM.atom_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      LPAREN -> State 116
      atom_type -> State 168
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119

  State 113:
    Kernel Items:
      traitable_type: EXCLAIM_EXCLAIM.atom_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      LPAREN -> State 116
      atom_type -> State 169
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119

  State 114:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 170

  State 115:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      field_decl: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      DOLLAR_LBRACKET -> State 61
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      optional_generic_binding -> State 171
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 172
      func_type -> State 123

  State 116:
    Kernel Items:
      implicit_struct_type: LPAREN.optional_field_decls RPAREN
      implicit_enum_type: LPAREN.field_decl_or_list RPAREN
    Reduce:
      RPAREN -> [optional_field_decls]
    Goto:
      IDENTIFIER -> State 115
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      field_decl -> State 173
      field_decls -> State 122
      optional_field_decls -> State 126
      implicit_struct_type -> State 125
      field_decl_or_list -> State 174
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 131
      func_type -> State 123

  State 117:
    Kernel Items:
      trait_type: TRAIT.LPAREN RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 175

  State 118:
    Kernel Items:
      traitable_type: atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 119:
    Kernel Items:
      atom_type: explicit_enum_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      atom_type: explicit_struct_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 121:
    Kernel Items:
      field_decls: field_decl., *
    Reduce:
      * -> [field_decls]
    Goto:
      (nil)

  State 122:
    Kernel Items:
      field_decls: field_decls.NEWLINES field_decl
      field_decls: field_decls.COMMA field_decl
      optional_field_decls: field_decls., RPAREN
    Reduce:
      RPAREN -> [optional_field_decls]
    Goto:
      NEWLINES -> State 177
      COMMA -> State 176

  State 123:
    Kernel Items:
      value_type: func_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      atom_type: implicit_enum_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      atom_type: implicit_struct_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 126:
    Kernel Items:
      implicit_struct_type: LPAREN optional_field_decls.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 178

  State 127:
    Kernel Items:
      trait_add_type: trait_add_type.ADD trait_mul_type
      trait_add_type: trait_add_type.SUB trait_mul_type
      value_type: trait_add_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      ADD -> State 179
      SUB -> State 180

  State 128:
    Kernel Items:
      trait_mul_type: trait_mul_type.MUL traitable_type
      trait_add_type: trait_mul_type., *
    Reduce:
      * -> [trait_add_type]
    Goto:
      MUL -> State 181

  State 129:
    Kernel Items:
      atom_type: trait_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 130:
    Kernel Items:
      trait_mul_type: traitable_type., *
    Reduce:
      * -> [trait_mul_type]
    Goto:
      (nil)

  State 131:
    Kernel Items:
      field_decl: value_type., *
    Reduce:
      * -> [field_decl]
    Goto:
      (nil)

  State 132:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 61
      optional_generic_binding -> State 171

  State 133:
    Kernel Items:
      generic_arguments: generic_arguments.COMMA value_type
      optional_generic_arguments: generic_arguments., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 182

  State 134:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 183

  State 135:
    Kernel Items:
      generic_arguments: value_type., *
    Reduce:
      * -> [generic_arguments]
    Goto:
      (nil)

  State 136:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 137:
    Kernel Items:
      access_expr: access_expr LBRACKET expression.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 184

  State 138:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN.optional_arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_arguments]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      argument -> State 56
      arguments -> State 185
      optional_arguments -> State 186
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 42
      add_expr -> State 32
      cmp_expr -> State 39
      and_expr -> State 33
      or_expr -> State 44
      sequence_expr -> State 48
      optional_label_decl -> State 43
      block_expr -> State 37
      expression -> State 58
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 139:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 85
      DIV -> State 83
      MOD -> State 84
      BIT_AND -> State 80
      BIT_LSHIFT -> State 81
      BIT_RSHIFT -> State 82
      mul_op -> State 86

  State 140:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 72
      NOT_EQUAL -> State 77
      LESS -> State 75
      LESS_OR_EQUAL -> State 76
      GREATER -> State 73
      GREATER_OR_EQUAL -> State 74
      cmp_op -> State 78

  State 141:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 66
      SUB -> State 69
      BIT_XOR -> State 68
      BIT_OR -> State 67
      add_op -> State 70

  State 142:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type LPAREN arguments.RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 187
      COMMA -> State 108

  State 143:
    Kernel Items:
      mul_expr: mul_expr mul_op prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 144:
    Kernel Items:
      atom_expr: block_expr., *
      loop_expr: FOR block_expr., $
    Reduce:
      * -> [atom_expr]
      $ -> [loop_expr]
    Goto:
      (nil)

  State 145:
    Kernel Items:
      loop_expr: FOR sequence_expr.block_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      LABEL_DECL -> State 21
      optional_label_decl -> State 96
      block_expr -> State 188

  State 146:
    Kernel Items:
      if_expr: IF sequence_expr.block_body
      if_expr: IF sequence_expr.block_body ELSE block_body
      if_expr: IF sequence_expr.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 89
      block_body -> State 189

  State 147:
    Kernel Items:
      statements: statements.statement
      block_body: LBRACE statements.RBRACE
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      RETURN -> State 194
      BREAK -> State 191
      CONTINUE -> State 192
      UNSAFE -> State 195
      STRUCT -> State 28
      FUNC -> State 18
      ASYNC -> State 190
      RBRACE -> State 193
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      call_expr -> State 38
      access_expr -> State 196
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 42
      add_expr -> State 32
      cmp_expr -> State 39
      and_expr -> State 33
      or_expr -> State 44
      sequence_expr -> State 48
      jump_type -> State 199
      expression_or_implicit_struct -> State 198
      unsafe_statement -> State 202
      statement_body -> State 201
      statement -> State 200
      optional_label_decl -> State 43
      block_expr -> State 37
      expression -> State 197
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 148:
    Kernel Items:
      switch_expr: SWITCH LBRACE.CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      CASE -> State 203

  State 149:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 71

  State 150:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements.RPAREN
      package_statements: package_statements.package_statement
    Reduce:
      (nil)
    Goto:
      UNSAFE -> State 195
      RPAREN -> State 204
      unsafe_statement -> State 207
      package_statement_body -> State 206
      package_statement -> State 205

  State 151:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 208
      func_type -> State 123

  State 152:
    Kernel Items:
      generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 153:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs.COMMA generic_parameter_def
      optional_generic_parameter_defs: generic_parameter_defs., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      COMMA -> State 209

  State 154:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 210

  State 155:
    Kernel Items:
      type_def: TYPE IDENTIFIER EQUAL value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      (nil)

  State 156:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      (nil)

  State 157:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT.parameter_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      QUESTION -> State 158
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 160
      parameter_type -> State 211
      func_type -> State 123

  State 158:
    Kernel Items:
      parameter_type: QUESTION., *
    Reduce:
      * -> [parameter_type]
    Goto:
      (nil)

  State 159:
    Kernel Items:
      parameter_def: IDENTIFIER parameter_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      (nil)

  State 160:
    Kernel Items:
      parameter_type: value_type., *
    Reduce:
      * -> [parameter_type]
    Goto:
      (nil)

  State 161:
    Kernel Items:
      optional_receiver: LPAREN parameter_def RPAREN., IDENTIFIER
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      (nil)

  State 162:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 212

  State 163:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      QUESTION -> State 213
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 215
      return_type -> State 214
      func_type -> State 123

  State 164:
    Kernel Items:
      parameter_defs: parameter_defs COMMA.parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 102
      parameter_def -> State 216

  State 165:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 166:
    Kernel Items:
      trait_add_type: trait_add_type.ADD trait_mul_type
      trait_add_type: trait_add_type.SUB trait_mul_type
      value_type: BIT_AND trait_add_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      ADD -> State 179
      SUB -> State 180

  State 167:
    Kernel Items:
      explicit_enum_type: ENUM LPAREN.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 217

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
      func_type: FUNC LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 219
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      QUESTION -> State 158
      DOTDOTDOT -> State 218
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 160
      parameter_type -> State 224
      parameter_decl -> State 221
      parameter_decls -> State 222
      optional_parameter_decls -> State 220
      func_type -> State 123
      parameter_def -> State 223

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
    Reduce:
      * -> [field_decl]
    Goto:
      (nil)

  State 173:
    Kernel Items:
      field_decls: field_decl., *
      field_decl_or_list: field_decl.OR field_decl
    Reduce:
      * -> [field_decls]
    Goto:
      OR -> State 225

  State 174:
    Kernel Items:
      field_decl_or_list: field_decl_or_list.OR field_decl
      implicit_enum_type: LPAREN field_decl_or_list.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 227
      OR -> State 226

  State 175:
    Kernel Items:
      trait_type: TRAIT LPAREN.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 228

  State 176:
    Kernel Items:
      field_decls: field_decls COMMA.field_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 115
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      field_decl -> State 229
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 131
      func_type -> State 123

  State 177:
    Kernel Items:
      field_decls: field_decls NEWLINES.field_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 115
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      field_decl -> State 230
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 131
      func_type -> State 123

  State 178:
    Kernel Items:
      implicit_struct_type: LPAREN optional_field_decls RPAREN., *
    Reduce:
      * -> [implicit_struct_type]
    Goto:
      (nil)

  State 179:
    Kernel Items:
      trait_add_type: trait_add_type ADD.trait_mul_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      LPAREN -> State 116
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 231
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119

  State 180:
    Kernel Items:
      trait_add_type: trait_add_type SUB.trait_mul_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      LPAREN -> State 116
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 232
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119

  State 181:
    Kernel Items:
      trait_mul_type: trait_mul_type MUL.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      LPAREN -> State 116
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      atom_type -> State 118
      traitable_type -> State 233
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119

  State 182:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 234
      func_type -> State 123

  State 183:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 184:
    Kernel Items:
      access_expr: access_expr LBRACKET expression RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 185:
    Kernel Items:
      arguments: arguments.COMMA argument
      optional_arguments: arguments., RPAREN
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 108

  State 186:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 235

  State 187:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 188:
    Kernel Items:
      loop_expr: FOR sequence_expr block_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 189:
    Kernel Items:
      if_expr: IF sequence_expr block_body., *
      if_expr: IF sequence_expr block_body.ELSE block_body
      if_expr: IF sequence_expr block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 236

  State 190:
    Kernel Items:
      statement_body: ASYNC.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      call_expr -> State 238
      access_expr -> State 237
      optional_label_decl -> State 96
      block_expr -> State 37
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 191:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 192:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 193:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 194:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 195:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 239

  State 196:
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
      LBRACKET -> State 63
      DOT -> State 62
      QUESTION -> State 64
      DOLLAR_LBRACKET -> State 61
      ADD_ASSIGN -> State 240
      SUB_ASSIGN -> State 251
      MUL_ASSIGN -> State 250
      DIV_ASSIGN -> State 248
      MOD_ASSIGN -> State 249
      ADD_ONE_ASSIGN -> State 241
      SUB_ONE_ASSIGN -> State 252
      BIT_NEG_ASSIGN -> State 244
      BIT_AND_ASSIGN -> State 242
      BIT_OR_ASSIGN -> State 245
      BIT_XOR_ASSIGN -> State 247
      BIT_LSHIFT_ASSIGN -> State 243
      BIT_RSHIFT_ASSIGN -> State 246
      optional_generic_binding -> State 65
      op_one_assign -> State 254
      binary_op_assign -> State 253

  State 197:
    Kernel Items:
      expression_or_implicit_struct: expression., *
    Reduce:
      * -> [expression_or_implicit_struct]
    Goto:
      (nil)

  State 198:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct.COMMA expression
      statement_body: expression_or_implicit_struct., *
    Reduce:
      * -> [statement_body]
    Goto:
      COMMA -> State 255

  State 199:
    Kernel Items:
      statement_body: jump_type.optional_jump_label optional_expression_or_implicit_struct
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 256
      optional_jump_label -> State 257

  State 200:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 201:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 258
      SEMICOLON -> State 259

  State 202:
    Kernel Items:
      statement_body: unsafe_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 203:
    Kernel Items:
      switch_expr: SWITCH LBRACE CASE.DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 260

  State 204:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements RPAREN., $
    Reduce:
      $ -> [package_def]
    Goto:
      (nil)

  State 205:
    Kernel Items:
      package_statements: package_statements package_statement., *
    Reduce:
      * -> [package_statements]
    Goto:
      (nil)

  State 206:
    Kernel Items:
      package_statement: package_statement_body.NEWLINES
      package_statement: package_statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 261
      SEMICOLON -> State 262

  State 207:
    Kernel Items:
      package_statement_body: unsafe_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 208:
    Kernel Items:
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      (nil)

  State 209:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA.generic_parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 151
      generic_parameter_def -> State 263

  State 210:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 211:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT parameter_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      (nil)

  State 212:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 102
      parameter_def -> State 106
      parameter_defs -> State 107
      optional_parameter_defs -> State 264

  State 213:
    Kernel Items:
      return_type: QUESTION., $
    Reduce:
      $ -> [return_type]
    Goto:
      (nil)

  State 214:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 89
      block_body -> State 265

  State 215:
    Kernel Items:
      return_type: value_type., $
    Reduce:
      $ -> [return_type]
    Goto:
      (nil)

  State 216:
    Kernel Items:
      parameter_defs: parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      explicit_enum_type: ENUM LPAREN RPAREN., *
    Reduce:
      * -> [explicit_enum_type]
    Goto:
      (nil)

  State 218:
    Kernel Items:
      parameter_decl: DOTDOTDOT.parameter_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      QUESTION -> State 158
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 160
      parameter_type -> State 266
      func_type -> State 123

  State 219:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      parameter_def: IDENTIFIER.parameter_type
      parameter_def: IDENTIFIER.DOTDOTDOT parameter_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      QUESTION -> State 158
      DOLLAR_LBRACKET -> State 61
      DOTDOTDOT -> State 157
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      optional_generic_binding -> State 171
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 160
      parameter_type -> State 159
      func_type -> State 123

  State 220:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 267

  State 221:
    Kernel Items:
      parameter_decls: parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 222:
    Kernel Items:
      parameter_decls: parameter_decls.COMMA parameter_decl
      optional_parameter_decls: parameter_decls., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      COMMA -> State 268

  State 223:
    Kernel Items:
      parameter_decl: parameter_def., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 224:
    Kernel Items:
      parameter_decl: parameter_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 225:
    Kernel Items:
      field_decl_or_list: field_decl OR.field_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 115
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      field_decl -> State 269
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 131
      func_type -> State 123

  State 226:
    Kernel Items:
      field_decl_or_list: field_decl_or_list OR.field_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 115
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      field_decl -> State 270
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 131
      func_type -> State 123

  State 227:
    Kernel Items:
      implicit_enum_type: LPAREN field_decl_or_list RPAREN., *
    Reduce:
      * -> [implicit_enum_type]
    Goto:
      (nil)

  State 228:
    Kernel Items:
      trait_type: TRAIT LPAREN RPAREN., *
    Reduce:
      * -> [trait_type]
    Goto:
      (nil)

  State 229:
    Kernel Items:
      field_decls: field_decls COMMA field_decl., *
    Reduce:
      * -> [field_decls]
    Goto:
      (nil)

  State 230:
    Kernel Items:
      field_decls: field_decls NEWLINES field_decl., *
    Reduce:
      * -> [field_decls]
    Goto:
      (nil)

  State 231:
    Kernel Items:
      trait_mul_type: trait_mul_type.MUL traitable_type
      trait_add_type: trait_add_type ADD trait_mul_type., *
    Reduce:
      * -> [trait_add_type]
    Goto:
      MUL -> State 181

  State 232:
    Kernel Items:
      trait_mul_type: trait_mul_type.MUL traitable_type
      trait_add_type: trait_add_type SUB trait_mul_type., *
    Reduce:
      * -> [trait_add_type]
    Goto:
      MUL -> State 181

  State 233:
    Kernel Items:
      trait_mul_type: trait_mul_type MUL traitable_type., *
    Reduce:
      * -> [trait_mul_type]
    Goto:
      (nil)

  State 234:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
    Reduce:
      * -> [generic_arguments]
    Goto:
      (nil)

  State 235:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 236:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE.block_body
      if_expr: IF sequence_expr block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 88
      LBRACE -> State 89
      block_body -> State 271
      if_expr -> State 272

  State 237:
    Kernel Items:
      call_expr: access_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET expression RBRACKET
    Reduce:
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 63
      DOT -> State 62
      DOLLAR_LBRACKET -> State 61
      optional_generic_binding -> State 65

  State 238:
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

  State 239:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 273

  State 240:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      op_one_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 242:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 243:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 244:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 245:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 246:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 247:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 248:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 249:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 250:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 251:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 252:
    Kernel Items:
      op_one_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 253:
    Kernel Items:
      statement_body: access_expr binary_op_assign.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 42
      add_expr -> State 32
      cmp_expr -> State 39
      and_expr -> State 33
      or_expr -> State 44
      sequence_expr -> State 48
      optional_label_decl -> State 43
      block_expr -> State 37
      expression -> State 274
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 254:
    Kernel Items:
      statement_body: access_expr op_one_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 255:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct COMMA.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 42
      add_expr -> State 32
      cmp_expr -> State 39
      and_expr -> State 33
      or_expr -> State 44
      sequence_expr -> State 48
      optional_label_decl -> State 43
      block_expr -> State 37
      expression -> State 275
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 256:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 257:
    Kernel Items:
      statement_body: jump_type optional_jump_label.optional_expression_or_implicit_struct
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_expression_or_implicit_struct]
      SEMICOLON -> [optional_expression_or_implicit_struct]
    Goto:
      INTEGER_LITERAL -> State 20
      FLOAT_LITERAL -> State 17
      RUNE_LITERAL -> State 26
      STRING_LITERAL -> State 27
      IDENTIFIER -> State 19
      TRUE -> State 30
      FALSE -> State 16
      STRUCT -> State 28
      FUNC -> State 18
      LPAREN -> State 23
      LABEL_DECL -> State 21
      NOT -> State 25
      SUB -> State 29
      MUL -> State 24
      BIT_NEG -> State 15
      BIT_AND -> State 14
      LEX_ERROR -> State 22
      literal -> State 41
      anonymous_struct_expr -> State 35
      atom_expr -> State 36
      call_expr -> State 38
      access_expr -> State 31
      postfix_unary_expr -> State 45
      prefix_unary_op -> State 47
      prefix_unary_expr -> State 46
      mul_expr -> State 42
      add_expr -> State 32
      cmp_expr -> State 39
      and_expr -> State 33
      or_expr -> State 44
      sequence_expr -> State 48
      optional_expression_or_implicit_struct -> State 277
      expression_or_implicit_struct -> State 276
      optional_label_decl -> State 43
      block_expr -> State 37
      expression -> State 197
      explicit_struct_type -> State 40
      anonymous_func_expr -> State 34

  State 258:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 259:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 260:
    Kernel Items:
      switch_expr: SWITCH LBRACE CASE DEFAULT.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 278

  State 261:
    Kernel Items:
      package_statement: package_statement_body NEWLINES., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 262:
    Kernel Items:
      package_statement: package_statement_body SEMICOLON., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 263:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 264:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 279

  State 265:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 266:
    Kernel Items:
      parameter_decl: DOTDOTDOT parameter_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 267:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      QUESTION -> State 213
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 215
      return_type -> State 280
      func_type -> State 123

  State 268:
    Kernel Items:
      parameter_decls: parameter_decls COMMA.parameter_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 219
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      QUESTION -> State 158
      DOTDOTDOT -> State 218
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 160
      parameter_type -> State 224
      parameter_decl -> State 281
      func_type -> State 123
      parameter_def -> State 223

  State 269:
    Kernel Items:
      field_decl_or_list: field_decl OR field_decl., *
    Reduce:
      * -> [field_decl_or_list]
    Goto:
      (nil)

  State 270:
    Kernel Items:
      field_decl_or_list: field_decl_or_list OR field_decl., *
    Reduce:
      * -> [field_decl_or_list]
    Goto:
      (nil)

  State 271:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 272:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 273:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 282

  State 274:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 275:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct COMMA expression., *
    Reduce:
      * -> [expression_or_implicit_struct]
    Goto:
      (nil)

  State 276:
    Kernel Items:
      optional_expression_or_implicit_struct: expression_or_implicit_struct., *
      expression_or_implicit_struct: expression_or_implicit_struct.COMMA expression
    Reduce:
      * -> [optional_expression_or_implicit_struct]
    Goto:
      COMMA -> State 255

  State 277:
    Kernel Items:
      statement_body: jump_type optional_jump_label optional_expression_or_implicit_struct., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      switch_expr: SWITCH LBRACE CASE DEFAULT RBRACE., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 132
      STRUCT -> State 28
      ENUM -> State 111
      TRAIT -> State 117
      FUNC -> State 114
      LPAREN -> State 116
      QUESTION -> State 213
      EXCLAIM -> State 112
      EXCLAIM_EXCLAIM -> State 113
      BIT_AND -> State 110
      atom_type -> State 118
      traitable_type -> State 130
      trait_mul_type -> State 128
      trait_add_type -> State 127
      implicit_struct_type -> State 125
      implicit_enum_type -> State 124
      explicit_struct_type -> State 120
      trait_type -> State 129
      explicit_enum_type -> State 119
      value_type -> State 215
      return_type -> State 283
      func_type -> State 123

  State 280:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN return_type., $
    Reduce:
      $ -> [func_type]
    Goto:
      (nil)

  State 281:
    Kernel Items:
      parameter_decls: parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 282:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 284

  State 283:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 89
      block_body -> State 285

  State 284:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 285:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

Number of states: 285
Number of shift actions: 1402
Number of reduce actions: 226
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
