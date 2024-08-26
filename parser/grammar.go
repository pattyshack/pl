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
	ImplementsToken      = SymbolId(281)
	FuncToken            = SymbolId(282)
	AsyncToken           = SymbolId(283)
	LbraceToken          = SymbolId(284)
	RbraceToken          = SymbolId(285)
	LparenToken          = SymbolId(286)
	RparenToken          = SymbolId(287)
	LbracketToken        = SymbolId(288)
	RbracketToken        = SymbolId(289)
	DotToken             = SymbolId(290)
	CommaToken           = SymbolId(291)
	QuestionToken        = SymbolId(292)
	SemicolonToken       = SymbolId(293)
	DollarLbracketToken  = SymbolId(294)
	DotdotdotToken       = SymbolId(295)
	TildeTildeToken      = SymbolId(296)
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
	// 110:2: literal -> TRUE: ...
	TrueToLiteral(True_ *GenericSymbol) (*GenericSymbol, error)

	// 111:2: literal -> FALSE: ...
	FalseToLiteral(False_ *GenericSymbol) (*GenericSymbol, error)

	// 112:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 113:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 114:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 115:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 118:2: anonymous_struct_expr -> explicit: ...
	ExplicitToAnonymousStructExpr(ExplicitStructDef_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 119:2: anonymous_struct_expr -> implicit: ...
	ImplicitToAnonymousStructExpr(Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 125:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 126:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 127:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 128:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 129:2: atom_expr -> anonymous_struct_expr: ...
	AnonymousStructExprToAtomExpr(AnonymousStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 130:2: atom_expr -> LEX_ERROR: ...
	LexErrorToAtomExpr(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 134:2: generic_arguments -> value_type: ...
	ValueTypeToGenericArguments(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 135:2: generic_arguments -> add: ...
	AddToGenericArguments(GenericArguments_ *GenericSymbol, Comma_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 138:2: optional_generic_arguments -> generic_arguments: ...
	GenericArgumentsToOptionalGenericArguments(GenericArguments_ *GenericSymbol) (*GenericSymbol, error)

	// 139:2: optional_generic_arguments -> nil: ...
	NilToOptionalGenericArguments() (*GenericSymbol, error)

	// 142:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ *GenericSymbol, OptionalGenericArguments_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 143:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (*GenericSymbol, error)

	// 147:2: argument -> positional: ...
	PositionalToArgument(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 150:2: arguments -> argument: ...
	ArgumentToArguments(Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 151:2: arguments -> add: ...
	AddToArguments(Arguments_ *GenericSymbol, Comma_ *GenericSymbol, Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 154:2: optional_arguments -> arguments: ...
	ArgumentsToOptionalArguments(Arguments_ *GenericSymbol) (*GenericSymbol, error)

	// 155:2: optional_arguments -> nil: ...
	NilToOptionalArguments() (*GenericSymbol, error)

	// 157:13: call_expr -> ...
	ToCallExpr(AccessExpr_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalArguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 160:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 161:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 162:2: access_expr -> call_expr: ...
	CallExprToAccessExpr(CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 164:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Expression_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 167:2: postfix_unary_expr -> access_expr: ...
	AccessExprToPostfixUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 168:2: postfix_unary_expr -> question: ...
	QuestionToPostfixUnaryExpr(AccessExpr_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 171:2: prefix_unary_op -> NOT: ...
	NotToPrefixUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 172:2: prefix_unary_op -> BIT_NEG: ...
	BitNegToPrefixUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 173:2: prefix_unary_op -> SUB: ...
	SubToPrefixUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 176:2: prefix_unary_op -> MUL: ...
	MulToPrefixUnaryOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 179:2: prefix_unary_op -> BIT_AND: ...
	BitAndToPrefixUnaryOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 182:2: prefix_unary_expr -> postfix_unary_expr: ...
	PostfixUnaryExprToPrefixUnaryExpr(PostfixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 183:2: prefix_unary_expr -> prefix_op: ...
	PrefixOpToPrefixUnaryExpr(PrefixUnaryOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 186:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 187:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 188:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 189:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 190:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 191:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 194:2: mul_expr -> prefix_unary_expr: ...
	PrefixUnaryExprToMulExpr(PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 195:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 198:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 199:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 200:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 201:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 204:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 205:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 208:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 209:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 210:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 211:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 212:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 213:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 216:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 217:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 220:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 221:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 224:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 225:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 227:17: sequence_expr -> ...
	ToSequenceExpr(OrExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 230:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 231:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 234:2: optional_expression_or_implicit_struct -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct(ExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 235:2: optional_expression_or_implicit_struct -> nil: ...
	NilToOptionalExpressionOrImplicitStruct() (*GenericSymbol, error)

	// 238:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 239:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 240:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 243:2: op_one_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToOpOneAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 244:2: op_one_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToOpOneAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 247:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 248:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 249:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 250:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 251:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 252:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 253:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 254:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 255:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 256:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 257:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 260:2: expression_or_implicit_struct -> expression: ...
	ExpressionToExpressionOrImplicitStruct(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 261:2: expression_or_implicit_struct -> implicit_struct: ...
	ImplicitStructToExpressionOrImplicitStruct(ExpressionOrImplicitStruct_ *GenericSymbol, Comma_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 265:20: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ *GenericSymbol, Less_ *GenericSymbol, Identifier_ *GenericSymbol, Greater_ *GenericSymbol, StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 268:2: statement_body -> unsafe_statement: ...
	UnsafeStatementToStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 270:2: statement_body -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToStatementBody(ExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 272:2: statement_body -> async: ...
	AsyncToStatementBody(Async_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 275:2: statement_body -> jump: ...
	JumpToStatementBody(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 288:2: statement_body -> op_one_assign: ...
	OpOneAssignToStatementBody(AccessExpr_ *GenericSymbol, OpOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 289:2: statement_body -> binary_op_assign: ...
	BinaryOpAssignToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 292:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 293:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 296:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 297:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 300:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 301:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (*GenericSymbol, error)

	// 303:14: block_body -> ...
	ToBlockBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 304:14: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 309:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 310:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, BlockBody_2 *GenericSymbol) (*GenericSymbol, error)

	// 311:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 314:15: switch_expr -> ...
	ToSwitchExpr(Switch_ *GenericSymbol, Lbrace_ *GenericSymbol, Case_ *GenericSymbol, Default_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 317:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(For_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 318:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 323:2: expression -> sequence_expr: ...
	SequenceExprToExpression(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 324:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 325:2: expression -> switch_expr: ...
	SwitchExprToExpression(OptionalLabelDecl_ *GenericSymbol, SwitchExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 326:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ *GenericSymbol, LoopExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 329:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 330:2: atom_type -> explicit_struct_def: ...
	ExplicitStructDefToAtomType(ExplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 331:2: atom_type -> implicit_struct_def: ...
	ImplicitStructDefToAtomType(ImplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 332:2: atom_type -> explicit_enum_def: ...
	ExplicitEnumDefToAtomType(ExplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 333:2: atom_type -> implicit_enum_def: ...
	ImplicitEnumDefToAtomType(ImplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 334:2: atom_type -> trait_def: ...
	TraitDefToAtomType(TraitDef_ *GenericSymbol) (*GenericSymbol, error)

	// 335:2: atom_type -> QUESTION: ...
	QuestionToAtomType(Question_ *GenericSymbol) (*GenericSymbol, error)

	// 338:2: traitable_type -> atom_type: ...
	AtomTypeToTraitableType(AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 339:2: traitable_type -> method_interface: ...
	MethodInterfaceToTraitableType(BitNeg_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 340:2: traitable_type -> trait: ...
	TraitToTraitableType(TildeTilde_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 343:2: trait_mul_type -> traitable_type: ...
	TraitableTypeToTraitMulType(TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 344:2: trait_mul_type -> intersect: ...
	IntersectToTraitMulType(TraitMulType_ *GenericSymbol, Mul_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 347:2: trait_add_type -> trait_mul_type: ...
	TraitMulTypeToTraitAddType(TraitMulType_ *GenericSymbol) (*GenericSymbol, error)

	// 348:2: trait_add_type -> union: ...
	UnionToTraitAddType(TraitAddType_ *GenericSymbol, Add_ *GenericSymbol, TraitMulType_ *GenericSymbol) (*GenericSymbol, error)

	// 353:2: trait_add_type -> difference: ...
	DifferenceToTraitAddType(TraitAddType_ *GenericSymbol, Sub_ *GenericSymbol, TraitMulType_ *GenericSymbol) (*GenericSymbol, error)

	// 356:2: value_type -> trait_add_type: ...
	TraitAddTypeToValueType(TraitAddType_ *GenericSymbol) (*GenericSymbol, error)

	// 357:2: value_type -> reference: ...
	ReferenceToValueType(BitAnd_ *GenericSymbol, TraitAddType_ *GenericSymbol) (*GenericSymbol, error)

	// 358:2: value_type -> func_type: ...
	FuncTypeToValueType(FuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 361:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 362:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol, Implements_ *GenericSymbol, ValueType_2 *GenericSymbol) (*GenericSymbol, error)

	// 363:2: type_def -> alias: ...
	AliasToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, Equal_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 370:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 371:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 374:2: generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToGenericParameterDefs(GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 375:2: generic_parameter_defs -> add: ...
	AddToGenericParameterDefs(GenericParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 378:2: optional_generic_parameter_defs -> generic_parameter_defs: ...
	GenericParameterDefsToOptionalGenericParameterDefs(GenericParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 379:2: optional_generic_parameter_defs -> nil: ...
	NilToOptionalGenericParameterDefs() (*GenericSymbol, error)

	// 382:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ *GenericSymbol, OptionalGenericParameterDefs_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 383:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (*GenericSymbol, error)

	// 390:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 391:2: field_def -> implicit: ...
	ImplicitToFieldDef(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 394:2: implicit_field_defs -> field_def: ...
	FieldDefToImplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 395:2: implicit_field_defs -> add: ...
	AddToImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 398:2: optional_implicit_field_defs -> implicit_field_defs: ...
	ImplicitFieldDefsToOptionalImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 399:2: optional_implicit_field_defs -> nil: ...
	NilToOptionalImplicitFieldDefs() (*GenericSymbol, error)

	// 401:23: implicit_struct_def -> ...
	ToImplicitStructDef(Lparen_ *GenericSymbol, OptionalImplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 404:2: explicit_field_defs -> field_def: ...
	FieldDefToExplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 405:2: explicit_field_defs -> implicit: ...
	ImplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Newlines_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 406:2: explicit_field_defs -> explicit: ...
	ExplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 409:2: optional_explicit_field_defs -> explicit_field_defs: ...
	ExplicitFieldDefsToOptionalExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 410:2: optional_explicit_field_defs -> nil: ...
	NilToOptionalExplicitFieldDefs() (*GenericSymbol, error)

	// 412:23: explicit_struct_def -> ...
	ToExplicitStructDef(Struct_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalExplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 422:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(FieldDef_ *GenericSymbol, Or_ *GenericSymbol, FieldDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 423:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 425:21: implicit_enum_def -> ...
	ToImplicitEnumDef(Lparen_ *GenericSymbol, ImplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 428:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(FieldDef_ *GenericSymbol, Or_ *GenericSymbol, FieldDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 429:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(FieldDef_ *GenericSymbol, Newlines_ *GenericSymbol, FieldDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 430:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 431:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Newlines_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 433:21: explicit_enum_def -> ...
	ToExplicitEnumDef(Enum_ *GenericSymbol, Lparen_ *GenericSymbol, ExplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 440:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 441:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ *GenericSymbol) (*GenericSymbol, error)

	// 444:2: trait_properties -> trait_property: ...
	TraitPropertyToTraitProperties(TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 445:2: trait_properties -> implicit: ...
	ImplicitToTraitProperties(TraitProperties_ *GenericSymbol, Newlines_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 446:2: trait_properties -> explicit: ...
	ExplicitToTraitProperties(TraitProperties_ *GenericSymbol, Comma_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 449:2: optional_trait_properties -> trait_properties: ...
	TraitPropertiesToOptionalTraitProperties(TraitProperties_ *GenericSymbol) (*GenericSymbol, error)

	// 450:2: optional_trait_properties -> nil: ...
	NilToOptionalTraitProperties() (*GenericSymbol, error)

	// 452:13: trait_def -> ...
	ToTraitDef(Trait_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalTraitProperties_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 460:2: return_type -> value_type: ...
	ValueTypeToReturnType(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 461:2: return_type -> nil: ...
	NilToReturnType() (*GenericSymbol, error)

	// 464:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 465:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 466:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 467:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 470:2: parameter_decls -> parameter_decl: ...
	ParameterDeclToParameterDecls(ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 471:2: parameter_decls -> add: ...
	AddToParameterDecls(ParameterDecls_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 474:2: optional_parameter_decls -> parameter_decls: ...
	ParameterDeclsToOptionalParameterDecls(ParameterDecls_ *GenericSymbol) (*GenericSymbol, error)

	// 475:2: optional_parameter_decls -> nil: ...
	NilToOptionalParameterDecls() (*GenericSymbol, error)

	// 477:13: func_type -> ...
	ToFuncType(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 485:20: method_signature -> ...
	ToMethodSignature(Func_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 488:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 489:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 492:2: parameter_defs -> parameter_def: ...
	ParameterDefToParameterDefs(ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 493:2: parameter_defs -> add: ...
	AddToParameterDefs(ParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 496:2: optional_parameter_defs -> parameter_defs: ...
	ParameterDefsToOptionalParameterDefs(ParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 497:2: optional_parameter_defs -> nil: ...
	NilToOptionalParameterDefs() (*GenericSymbol, error)

	// 500:2: optional_receiver -> receiver: ...
	ReceiverToOptionalReceiver(Lparen_ *GenericSymbol, ParameterDef_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 501:2: optional_receiver -> nil: ...
	NilToOptionalReceiver() (*GenericSymbol, error)

	// 504:2: named_func_def -> ...
	ToNamedFuncDef(Func_ *GenericSymbol, OptionalReceiver_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 507:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 514:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 515:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, PackageStatements_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 517:26: package_statement_body -> ...
	ToPackageStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 520:2: package_statement -> implicit: ...
	ImplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 521:2: package_statement -> explicit: ...
	ExplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 524:2: package_statements -> empty_list: ...
	EmptyListToPackageStatements() (*GenericSymbol, error)

	// 525:2: package_statements -> add: ...
	AddToPackageStatements(PackageStatements_ *GenericSymbol, PackageStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 529:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 530:2: lex_internal_tokens -> COMMENT: ...
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

func ParseTraitDef(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParseTraitDefWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseTraitDefWithCustomErrorHandler(
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

	item, err := _Parse(lexer, reducer, errHandler, _State4)
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

	item, err := _Parse(lexer, reducer, errHandler, _State5)
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

	item, err := _Parse(lexer, reducer, errHandler, _State6)
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
	case ImplementsToken:
		return "IMPLEMENTS"
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
	case TildeTildeToken:
		return "TILDE_TILDE"
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
	case FieldDefType:
		return "field_def"
	case ImplicitFieldDefsType:
		return "implicit_field_defs"
	case OptionalImplicitFieldDefsType:
		return "optional_implicit_field_defs"
	case ImplicitStructDefType:
		return "implicit_struct_def"
	case ExplicitFieldDefsType:
		return "explicit_field_defs"
	case OptionalExplicitFieldDefsType:
		return "optional_explicit_field_defs"
	case ExplicitStructDefType:
		return "explicit_struct_def"
	case ImplicitEnumValueDefsType:
		return "implicit_enum_value_defs"
	case ImplicitEnumDefType:
		return "implicit_enum_def"
	case ExplicitEnumValueDefsType:
		return "explicit_enum_value_defs"
	case ExplicitEnumDefType:
		return "explicit_enum_def"
	case TraitPropertyType:
		return "trait_property"
	case TraitPropertiesType:
		return "trait_properties"
	case OptionalTraitPropertiesType:
		return "optional_trait_properties"
	case TraitDefType:
		return "trait_def"
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
	case MethodSignatureType:
		return "method_signature"
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
	ValueTypeType                          = SymbolId(377)
	TypeDefType                            = SymbolId(378)
	GenericParameterDefType                = SymbolId(379)
	GenericParameterDefsType               = SymbolId(380)
	OptionalGenericParameterDefsType       = SymbolId(381)
	OptionalGenericParametersType          = SymbolId(382)
	FieldDefType                           = SymbolId(383)
	ImplicitFieldDefsType                  = SymbolId(384)
	OptionalImplicitFieldDefsType          = SymbolId(385)
	ImplicitStructDefType                  = SymbolId(386)
	ExplicitFieldDefsType                  = SymbolId(387)
	OptionalExplicitFieldDefsType          = SymbolId(388)
	ExplicitStructDefType                  = SymbolId(389)
	ImplicitEnumValueDefsType              = SymbolId(390)
	ImplicitEnumDefType                    = SymbolId(391)
	ExplicitEnumValueDefsType              = SymbolId(392)
	ExplicitEnumDefType                    = SymbolId(393)
	TraitPropertyType                      = SymbolId(394)
	TraitPropertiesType                    = SymbolId(395)
	OptionalTraitPropertiesType            = SymbolId(396)
	TraitDefType                           = SymbolId(397)
	ReturnTypeType                         = SymbolId(398)
	ParameterDeclType                      = SymbolId(399)
	ParameterDeclsType                     = SymbolId(400)
	OptionalParameterDeclsType             = SymbolId(401)
	FuncTypeType                           = SymbolId(402)
	MethodSignatureType                    = SymbolId(403)
	ParameterDefType                       = SymbolId(404)
	ParameterDefsType                      = SymbolId(405)
	OptionalParameterDefsType              = SymbolId(406)
	OptionalReceiverType                   = SymbolId(407)
	NamedFuncDefType                       = SymbolId(408)
	AnonymousFuncExprType                  = SymbolId(409)
	PackageDefType                         = SymbolId(410)
	PackageStatementBodyType               = SymbolId(411)
	PackageStatementType                   = SymbolId(412)
	PackageStatementsType                  = SymbolId(413)
	LexInternalTokensType                  = SymbolId(414)
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
	_ReduceExplicitStructDefToAtomType                                    = _ReduceType(115)
	_ReduceImplicitStructDefToAtomType                                    = _ReduceType(116)
	_ReduceExplicitEnumDefToAtomType                                      = _ReduceType(117)
	_ReduceImplicitEnumDefToAtomType                                      = _ReduceType(118)
	_ReduceTraitDefToAtomType                                             = _ReduceType(119)
	_ReduceQuestionToAtomType                                             = _ReduceType(120)
	_ReduceAtomTypeToTraitableType                                        = _ReduceType(121)
	_ReduceMethodInterfaceToTraitableType                                 = _ReduceType(122)
	_ReduceTraitToTraitableType                                           = _ReduceType(123)
	_ReduceTraitableTypeToTraitMulType                                    = _ReduceType(124)
	_ReduceIntersectToTraitMulType                                        = _ReduceType(125)
	_ReduceTraitMulTypeToTraitAddType                                     = _ReduceType(126)
	_ReduceUnionToTraitAddType                                            = _ReduceType(127)
	_ReduceDifferenceToTraitAddType                                       = _ReduceType(128)
	_ReduceTraitAddTypeToValueType                                        = _ReduceType(129)
	_ReduceReferenceToValueType                                           = _ReduceType(130)
	_ReduceFuncTypeToValueType                                            = _ReduceType(131)
	_ReduceDefinitionToTypeDef                                            = _ReduceType(132)
	_ReduceConstrainedDefToTypeDef                                        = _ReduceType(133)
	_ReduceAliasToTypeDef                                                 = _ReduceType(134)
	_ReduceUnconstrainedToGenericParameterDef                             = _ReduceType(135)
	_ReduceConstrainedToGenericParameterDef                               = _ReduceType(136)
	_ReduceGenericParameterDefToGenericParameterDefs                      = _ReduceType(137)
	_ReduceAddToGenericParameterDefs                                      = _ReduceType(138)
	_ReduceGenericParameterDefsToOptionalGenericParameterDefs             = _ReduceType(139)
	_ReduceNilToOptionalGenericParameterDefs                              = _ReduceType(140)
	_ReduceGenericToOptionalGenericParameters                             = _ReduceType(141)
	_ReduceNilToOptionalGenericParameters                                 = _ReduceType(142)
	_ReduceExplicitToFieldDef                                             = _ReduceType(143)
	_ReduceImplicitToFieldDef                                             = _ReduceType(144)
	_ReduceFieldDefToImplicitFieldDefs                                    = _ReduceType(145)
	_ReduceAddToImplicitFieldDefs                                         = _ReduceType(146)
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefs                   = _ReduceType(147)
	_ReduceNilToOptionalImplicitFieldDefs                                 = _ReduceType(148)
	_ReduceToImplicitStructDef                                            = _ReduceType(149)
	_ReduceFieldDefToExplicitFieldDefs                                    = _ReduceType(150)
	_ReduceImplicitToExplicitFieldDefs                                    = _ReduceType(151)
	_ReduceExplicitToExplicitFieldDefs                                    = _ReduceType(152)
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefs                   = _ReduceType(153)
	_ReduceNilToOptionalExplicitFieldDefs                                 = _ReduceType(154)
	_ReduceToExplicitStructDef                                            = _ReduceType(155)
	_ReducePairToImplicitEnumValueDefs                                    = _ReduceType(156)
	_ReduceAddToImplicitEnumValueDefs                                     = _ReduceType(157)
	_ReduceToImplicitEnumDef                                              = _ReduceType(158)
	_ReduceExplicitPairToExplicitEnumValueDefs                            = _ReduceType(159)
	_ReduceImplicitPairToExplicitEnumValueDefs                            = _ReduceType(160)
	_ReduceExplicitAddToExplicitEnumValueDefs                             = _ReduceType(161)
	_ReduceImplicitAddToExplicitEnumValueDefs                             = _ReduceType(162)
	_ReduceToExplicitEnumDef                                              = _ReduceType(163)
	_ReduceFieldDefToTraitProperty                                        = _ReduceType(164)
	_ReduceMethodSignatureToTraitProperty                                 = _ReduceType(165)
	_ReduceTraitPropertyToTraitProperties                                 = _ReduceType(166)
	_ReduceImplicitToTraitProperties                                      = _ReduceType(167)
	_ReduceExplicitToTraitProperties                                      = _ReduceType(168)
	_ReduceTraitPropertiesToOptionalTraitProperties                       = _ReduceType(169)
	_ReduceNilToOptionalTraitProperties                                   = _ReduceType(170)
	_ReduceToTraitDef                                                     = _ReduceType(171)
	_ReduceValueTypeToReturnType                                          = _ReduceType(172)
	_ReduceNilToReturnType                                                = _ReduceType(173)
	_ReduceArgToParameterDecl                                             = _ReduceType(174)
	_ReduceVarargToParameterDecl                                          = _ReduceType(175)
	_ReduceUnamedToParameterDecl                                          = _ReduceType(176)
	_ReduceUnnamedVarargToParameterDecl                                   = _ReduceType(177)
	_ReduceParameterDeclToParameterDecls                                  = _ReduceType(178)
	_ReduceAddToParameterDecls                                            = _ReduceType(179)
	_ReduceParameterDeclsToOptionalParameterDecls                         = _ReduceType(180)
	_ReduceNilToOptionalParameterDecls                                    = _ReduceType(181)
	_ReduceToFuncType                                                     = _ReduceType(182)
	_ReduceToMethodSignature                                              = _ReduceType(183)
	_ReduceArgToParameterDef                                              = _ReduceType(184)
	_ReduceVarargToParameterDef                                           = _ReduceType(185)
	_ReduceParameterDefToParameterDefs                                    = _ReduceType(186)
	_ReduceAddToParameterDefs                                             = _ReduceType(187)
	_ReduceParameterDefsToOptionalParameterDefs                           = _ReduceType(188)
	_ReduceNilToOptionalParameterDefs                                     = _ReduceType(189)
	_ReduceReceiverToOptionalReceiver                                     = _ReduceType(190)
	_ReduceNilToOptionalReceiver                                          = _ReduceType(191)
	_ReduceToNamedFuncDef                                                 = _ReduceType(192)
	_ReduceToAnonymousFuncExpr                                            = _ReduceType(193)
	_ReduceNoSpecToPackageDef                                             = _ReduceType(194)
	_ReduceWithSpecToPackageDef                                           = _ReduceType(195)
	_ReduceToPackageStatementBody                                         = _ReduceType(196)
	_ReduceImplicitToPackageStatement                                     = _ReduceType(197)
	_ReduceExplicitToPackageStatement                                     = _ReduceType(198)
	_ReduceEmptyListToPackageStatements                                   = _ReduceType(199)
	_ReduceAddToPackageStatements                                         = _ReduceType(200)
	_ReduceSpacesToLexInternalTokens                                      = _ReduceType(201)
	_ReduceCommentToLexInternalTokens                                     = _ReduceType(202)
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
	case _ReduceExplicitStructDefToAtomType:
		return "ExplicitStructDefToAtomType"
	case _ReduceImplicitStructDefToAtomType:
		return "ImplicitStructDefToAtomType"
	case _ReduceExplicitEnumDefToAtomType:
		return "ExplicitEnumDefToAtomType"
	case _ReduceImplicitEnumDefToAtomType:
		return "ImplicitEnumDefToAtomType"
	case _ReduceTraitDefToAtomType:
		return "TraitDefToAtomType"
	case _ReduceQuestionToAtomType:
		return "QuestionToAtomType"
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
	case _ReduceTraitAddTypeToValueType:
		return "TraitAddTypeToValueType"
	case _ReduceReferenceToValueType:
		return "ReferenceToValueType"
	case _ReduceFuncTypeToValueType:
		return "FuncTypeToValueType"
	case _ReduceDefinitionToTypeDef:
		return "DefinitionToTypeDef"
	case _ReduceConstrainedDefToTypeDef:
		return "ConstrainedDefToTypeDef"
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
	case _ReduceExplicitToFieldDef:
		return "ExplicitToFieldDef"
	case _ReduceImplicitToFieldDef:
		return "ImplicitToFieldDef"
	case _ReduceFieldDefToImplicitFieldDefs:
		return "FieldDefToImplicitFieldDefs"
	case _ReduceAddToImplicitFieldDefs:
		return "AddToImplicitFieldDefs"
	case _ReduceImplicitFieldDefsToOptionalImplicitFieldDefs:
		return "ImplicitFieldDefsToOptionalImplicitFieldDefs"
	case _ReduceNilToOptionalImplicitFieldDefs:
		return "NilToOptionalImplicitFieldDefs"
	case _ReduceToImplicitStructDef:
		return "ToImplicitStructDef"
	case _ReduceFieldDefToExplicitFieldDefs:
		return "FieldDefToExplicitFieldDefs"
	case _ReduceImplicitToExplicitFieldDefs:
		return "ImplicitToExplicitFieldDefs"
	case _ReduceExplicitToExplicitFieldDefs:
		return "ExplicitToExplicitFieldDefs"
	case _ReduceExplicitFieldDefsToOptionalExplicitFieldDefs:
		return "ExplicitFieldDefsToOptionalExplicitFieldDefs"
	case _ReduceNilToOptionalExplicitFieldDefs:
		return "NilToOptionalExplicitFieldDefs"
	case _ReduceToExplicitStructDef:
		return "ToExplicitStructDef"
	case _ReducePairToImplicitEnumValueDefs:
		return "PairToImplicitEnumValueDefs"
	case _ReduceAddToImplicitEnumValueDefs:
		return "AddToImplicitEnumValueDefs"
	case _ReduceToImplicitEnumDef:
		return "ToImplicitEnumDef"
	case _ReduceExplicitPairToExplicitEnumValueDefs:
		return "ExplicitPairToExplicitEnumValueDefs"
	case _ReduceImplicitPairToExplicitEnumValueDefs:
		return "ImplicitPairToExplicitEnumValueDefs"
	case _ReduceExplicitAddToExplicitEnumValueDefs:
		return "ExplicitAddToExplicitEnumValueDefs"
	case _ReduceImplicitAddToExplicitEnumValueDefs:
		return "ImplicitAddToExplicitEnumValueDefs"
	case _ReduceToExplicitEnumDef:
		return "ToExplicitEnumDef"
	case _ReduceFieldDefToTraitProperty:
		return "FieldDefToTraitProperty"
	case _ReduceMethodSignatureToTraitProperty:
		return "MethodSignatureToTraitProperty"
	case _ReduceTraitPropertyToTraitProperties:
		return "TraitPropertyToTraitProperties"
	case _ReduceImplicitToTraitProperties:
		return "ImplicitToTraitProperties"
	case _ReduceExplicitToTraitProperties:
		return "ExplicitToTraitProperties"
	case _ReduceTraitPropertiesToOptionalTraitProperties:
		return "TraitPropertiesToOptionalTraitProperties"
	case _ReduceNilToOptionalTraitProperties:
		return "NilToOptionalTraitProperties"
	case _ReduceToTraitDef:
		return "ToTraitDef"
	case _ReduceValueTypeToReturnType:
		return "ValueTypeToReturnType"
	case _ReduceNilToReturnType:
		return "NilToReturnType"
	case _ReduceArgToParameterDecl:
		return "ArgToParameterDecl"
	case _ReduceVarargToParameterDecl:
		return "VarargToParameterDecl"
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
	case _ReduceToMethodSignature:
		return "ToMethodSignature"
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
	_State286 = _StateId(286)
	_State287 = _StateId(287)
	_State288 = _StateId(288)
	_State289 = _StateId(289)
	_State290 = _StateId(290)
	_State291 = _StateId(291)
	_State292 = _StateId(292)
	_State293 = _StateId(293)
	_State294 = _StateId(294)
	_State295 = _StateId(295)
	_State296 = _StateId(296)
	_State297 = _StateId(297)
	_State298 = _StateId(298)
	_State299 = _StateId(299)
	_State300 = _StateId(300)
	_State301 = _StateId(301)
	_State302 = _StateId(302)
	_State303 = _StateId(303)
	_State304 = _StateId(304)
	_State305 = _StateId(305)
	_State306 = _StateId(306)
	_State307 = _StateId(307)
	_State308 = _StateId(308)
	_State309 = _StateId(309)
	_State310 = _StateId(310)
	_State311 = _StateId(311)
	_State312 = _StateId(312)
	_State313 = _StateId(313)
	_State314 = _StateId(314)
	_State315 = _StateId(315)
	_State316 = _StateId(316)
	_State317 = _StateId(317)
	_State318 = _StateId(318)
	_State319 = _StateId(319)
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
	case _EndMarker, SpacesToken, NewlinesToken, CommentToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, ReturnToken, BreakToken, ContinueToken, PackageToken, UnsafeToken, TypeToken, StructToken, EnumToken, TraitToken, ImplementsToken, FuncToken, AsyncToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, DollarLbracketToken, DotdotdotToken, TildeTildeToken, LabelDeclToken, JumpLabelToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, NotToken, AndToken, OrToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, LexErrorToken:
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
	case _ReduceExplicitStructDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ExplicitStructDefToAtomType(args[0].Generic_)
	case _ReduceImplicitStructDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ImplicitStructDefToAtomType(args[0].Generic_)
	case _ReduceExplicitEnumDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ExplicitEnumDefToAtomType(args[0].Generic_)
	case _ReduceImplicitEnumDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ImplicitEnumDefToAtomType(args[0].Generic_)
	case _ReduceTraitDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.TraitDefToAtomType(args[0].Generic_)
	case _ReduceQuestionToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.QuestionToAtomType(args[0].Generic_)
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
	case _ReduceConstrainedDefToTypeDef:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = TypeDefType
		symbol.Generic_, err = reducer.ConstrainedDefToTypeDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_)
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
	case _ReduceExplicitToFieldDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = FieldDefType
		symbol.Generic_, err = reducer.ExplicitToFieldDef(args[0].Generic_, args[1].Generic_)
	case _ReduceImplicitToFieldDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldDefType
		symbol.Generic_, err = reducer.ImplicitToFieldDef(args[0].Generic_)
	case _ReduceFieldDefToImplicitFieldDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ImplicitFieldDefsType
		symbol.Generic_, err = reducer.FieldDefToImplicitFieldDefs(args[0].Generic_)
	case _ReduceAddToImplicitFieldDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitFieldDefsType
		symbol.Generic_, err = reducer.AddToImplicitFieldDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceImplicitFieldDefsToOptionalImplicitFieldDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalImplicitFieldDefsType
		symbol.Generic_, err = reducer.ImplicitFieldDefsToOptionalImplicitFieldDefs(args[0].Generic_)
	case _ReduceNilToOptionalImplicitFieldDefs:
		symbol.SymbolId_ = OptionalImplicitFieldDefsType
		symbol.Generic_, err = reducer.NilToOptionalImplicitFieldDefs()
	case _ReduceToImplicitStructDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitStructDefType
		symbol.Generic_, err = reducer.ToImplicitStructDef(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceFieldDefToExplicitFieldDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExplicitFieldDefsType
		symbol.Generic_, err = reducer.FieldDefToExplicitFieldDefs(args[0].Generic_)
	case _ReduceImplicitToExplicitFieldDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitFieldDefsType
		symbol.Generic_, err = reducer.ImplicitToExplicitFieldDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceExplicitToExplicitFieldDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitFieldDefsType
		symbol.Generic_, err = reducer.ExplicitToExplicitFieldDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceExplicitFieldDefsToOptionalExplicitFieldDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalExplicitFieldDefsType
		symbol.Generic_, err = reducer.ExplicitFieldDefsToOptionalExplicitFieldDefs(args[0].Generic_)
	case _ReduceNilToOptionalExplicitFieldDefs:
		symbol.SymbolId_ = OptionalExplicitFieldDefsType
		symbol.Generic_, err = reducer.NilToOptionalExplicitFieldDefs()
	case _ReduceToExplicitStructDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ExplicitStructDefType
		symbol.Generic_, err = reducer.ToExplicitStructDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReducePairToImplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitEnumValueDefsType
		symbol.Generic_, err = reducer.PairToImplicitEnumValueDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceAddToImplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitEnumValueDefsType
		symbol.Generic_, err = reducer.AddToImplicitEnumValueDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToImplicitEnumDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitEnumDefType
		symbol.Generic_, err = reducer.ToImplicitEnumDef(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceExplicitPairToExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ExplicitPairToExplicitEnumValueDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceImplicitPairToExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ImplicitPairToExplicitEnumValueDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceExplicitAddToExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ExplicitAddToExplicitEnumValueDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceImplicitAddToExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ImplicitAddToExplicitEnumValueDefs(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToExplicitEnumDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ExplicitEnumDefType
		symbol.Generic_, err = reducer.ToExplicitEnumDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceFieldDefToTraitProperty:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitPropertyType
		symbol.Generic_, err = reducer.FieldDefToTraitProperty(args[0].Generic_)
	case _ReduceMethodSignatureToTraitProperty:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitPropertyType
		symbol.Generic_, err = reducer.MethodSignatureToTraitProperty(args[0].Generic_)
	case _ReduceTraitPropertyToTraitProperties:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitPropertiesType
		symbol.Generic_, err = reducer.TraitPropertyToTraitProperties(args[0].Generic_)
	case _ReduceImplicitToTraitProperties:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TraitPropertiesType
		symbol.Generic_, err = reducer.ImplicitToTraitProperties(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceExplicitToTraitProperties:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TraitPropertiesType
		symbol.Generic_, err = reducer.ExplicitToTraitProperties(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceTraitPropertiesToOptionalTraitProperties:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalTraitPropertiesType
		symbol.Generic_, err = reducer.TraitPropertiesToOptionalTraitProperties(args[0].Generic_)
	case _ReduceNilToOptionalTraitProperties:
		symbol.SymbolId_ = OptionalTraitPropertiesType
		symbol.Generic_, err = reducer.NilToOptionalTraitProperties()
	case _ReduceToTraitDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = TraitDefType
		symbol.Generic_, err = reducer.ToTraitDef(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceValueTypeToReturnType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnTypeType
		symbol.Generic_, err = reducer.ValueTypeToReturnType(args[0].Generic_)
	case _ReduceNilToReturnType:
		symbol.SymbolId_ = ReturnTypeType
		symbol.Generic_, err = reducer.NilToReturnType()
	case _ReduceArgToParameterDecl:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.ArgToParameterDecl(args[0].Generic_, args[1].Generic_)
	case _ReduceVarargToParameterDecl:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.VarargToParameterDecl(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	case _ReduceToMethodSignature:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = MethodSignatureType
		symbol.Generic_, err = reducer.ToMethodSignature(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_)
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
	_GotoState286Action                                                         = &_Action{_ShiftAction, _State286, 0}
	_GotoState287Action                                                         = &_Action{_ShiftAction, _State287, 0}
	_GotoState288Action                                                         = &_Action{_ShiftAction, _State288, 0}
	_GotoState289Action                                                         = &_Action{_ShiftAction, _State289, 0}
	_GotoState290Action                                                         = &_Action{_ShiftAction, _State290, 0}
	_GotoState291Action                                                         = &_Action{_ShiftAction, _State291, 0}
	_GotoState292Action                                                         = &_Action{_ShiftAction, _State292, 0}
	_GotoState293Action                                                         = &_Action{_ShiftAction, _State293, 0}
	_GotoState294Action                                                         = &_Action{_ShiftAction, _State294, 0}
	_GotoState295Action                                                         = &_Action{_ShiftAction, _State295, 0}
	_GotoState296Action                                                         = &_Action{_ShiftAction, _State296, 0}
	_GotoState297Action                                                         = &_Action{_ShiftAction, _State297, 0}
	_GotoState298Action                                                         = &_Action{_ShiftAction, _State298, 0}
	_GotoState299Action                                                         = &_Action{_ShiftAction, _State299, 0}
	_GotoState300Action                                                         = &_Action{_ShiftAction, _State300, 0}
	_GotoState301Action                                                         = &_Action{_ShiftAction, _State301, 0}
	_GotoState302Action                                                         = &_Action{_ShiftAction, _State302, 0}
	_GotoState303Action                                                         = &_Action{_ShiftAction, _State303, 0}
	_GotoState304Action                                                         = &_Action{_ShiftAction, _State304, 0}
	_GotoState305Action                                                         = &_Action{_ShiftAction, _State305, 0}
	_GotoState306Action                                                         = &_Action{_ShiftAction, _State306, 0}
	_GotoState307Action                                                         = &_Action{_ShiftAction, _State307, 0}
	_GotoState308Action                                                         = &_Action{_ShiftAction, _State308, 0}
	_GotoState309Action                                                         = &_Action{_ShiftAction, _State309, 0}
	_GotoState310Action                                                         = &_Action{_ShiftAction, _State310, 0}
	_GotoState311Action                                                         = &_Action{_ShiftAction, _State311, 0}
	_GotoState312Action                                                         = &_Action{_ShiftAction, _State312, 0}
	_GotoState313Action                                                         = &_Action{_ShiftAction, _State313, 0}
	_GotoState314Action                                                         = &_Action{_ShiftAction, _State314, 0}
	_GotoState315Action                                                         = &_Action{_ShiftAction, _State315, 0}
	_GotoState316Action                                                         = &_Action{_ShiftAction, _State316, 0}
	_GotoState317Action                                                         = &_Action{_ShiftAction, _State317, 0}
	_GotoState318Action                                                         = &_Action{_ShiftAction, _State318, 0}
	_GotoState319Action                                                         = &_Action{_ShiftAction, _State319, 0}
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
	_ReduceExplicitStructDefToAtomTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceExplicitStructDefToAtomType}
	_ReduceImplicitStructDefToAtomTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceImplicitStructDefToAtomType}
	_ReduceExplicitEnumDefToAtomTypeAction                                      = &_Action{_ReduceAction, 0, _ReduceExplicitEnumDefToAtomType}
	_ReduceImplicitEnumDefToAtomTypeAction                                      = &_Action{_ReduceAction, 0, _ReduceImplicitEnumDefToAtomType}
	_ReduceTraitDefToAtomTypeAction                                             = &_Action{_ReduceAction, 0, _ReduceTraitDefToAtomType}
	_ReduceQuestionToAtomTypeAction                                             = &_Action{_ReduceAction, 0, _ReduceQuestionToAtomType}
	_ReduceAtomTypeToTraitableTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceAtomTypeToTraitableType}
	_ReduceMethodInterfaceToTraitableTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceMethodInterfaceToTraitableType}
	_ReduceTraitToTraitableTypeAction                                           = &_Action{_ReduceAction, 0, _ReduceTraitToTraitableType}
	_ReduceTraitableTypeToTraitMulTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceTraitableTypeToTraitMulType}
	_ReduceIntersectToTraitMulTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceIntersectToTraitMulType}
	_ReduceTraitMulTypeToTraitAddTypeAction                                     = &_Action{_ReduceAction, 0, _ReduceTraitMulTypeToTraitAddType}
	_ReduceUnionToTraitAddTypeAction                                            = &_Action{_ReduceAction, 0, _ReduceUnionToTraitAddType}
	_ReduceDifferenceToTraitAddTypeAction                                       = &_Action{_ReduceAction, 0, _ReduceDifferenceToTraitAddType}
	_ReduceTraitAddTypeToValueTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceTraitAddTypeToValueType}
	_ReduceReferenceToValueTypeAction                                           = &_Action{_ReduceAction, 0, _ReduceReferenceToValueType}
	_ReduceFuncTypeToValueTypeAction                                            = &_Action{_ReduceAction, 0, _ReduceFuncTypeToValueType}
	_ReduceDefinitionToTypeDefAction                                            = &_Action{_ReduceAction, 0, _ReduceDefinitionToTypeDef}
	_ReduceConstrainedDefToTypeDefAction                                        = &_Action{_ReduceAction, 0, _ReduceConstrainedDefToTypeDef}
	_ReduceAliasToTypeDefAction                                                 = &_Action{_ReduceAction, 0, _ReduceAliasToTypeDef}
	_ReduceUnconstrainedToGenericParameterDefAction                             = &_Action{_ReduceAction, 0, _ReduceUnconstrainedToGenericParameterDef}
	_ReduceConstrainedToGenericParameterDefAction                               = &_Action{_ReduceAction, 0, _ReduceConstrainedToGenericParameterDef}
	_ReduceGenericParameterDefToGenericParameterDefsAction                      = &_Action{_ReduceAction, 0, _ReduceGenericParameterDefToGenericParameterDefs}
	_ReduceAddToGenericParameterDefsAction                                      = &_Action{_ReduceAction, 0, _ReduceAddToGenericParameterDefs}
	_ReduceGenericParameterDefsToOptionalGenericParameterDefsAction             = &_Action{_ReduceAction, 0, _ReduceGenericParameterDefsToOptionalGenericParameterDefs}
	_ReduceNilToOptionalGenericParameterDefsAction                              = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericParameterDefs}
	_ReduceGenericToOptionalGenericParametersAction                             = &_Action{_ReduceAction, 0, _ReduceGenericToOptionalGenericParameters}
	_ReduceNilToOptionalGenericParametersAction                                 = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericParameters}
	_ReduceExplicitToFieldDefAction                                             = &_Action{_ReduceAction, 0, _ReduceExplicitToFieldDef}
	_ReduceImplicitToFieldDefAction                                             = &_Action{_ReduceAction, 0, _ReduceImplicitToFieldDef}
	_ReduceFieldDefToImplicitFieldDefsAction                                    = &_Action{_ReduceAction, 0, _ReduceFieldDefToImplicitFieldDefs}
	_ReduceAddToImplicitFieldDefsAction                                         = &_Action{_ReduceAction, 0, _ReduceAddToImplicitFieldDefs}
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction                   = &_Action{_ReduceAction, 0, _ReduceImplicitFieldDefsToOptionalImplicitFieldDefs}
	_ReduceNilToOptionalImplicitFieldDefsAction                                 = &_Action{_ReduceAction, 0, _ReduceNilToOptionalImplicitFieldDefs}
	_ReduceToImplicitStructDefAction                                            = &_Action{_ReduceAction, 0, _ReduceToImplicitStructDef}
	_ReduceFieldDefToExplicitFieldDefsAction                                    = &_Action{_ReduceAction, 0, _ReduceFieldDefToExplicitFieldDefs}
	_ReduceImplicitToExplicitFieldDefsAction                                    = &_Action{_ReduceAction, 0, _ReduceImplicitToExplicitFieldDefs}
	_ReduceExplicitToExplicitFieldDefsAction                                    = &_Action{_ReduceAction, 0, _ReduceExplicitToExplicitFieldDefs}
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction                   = &_Action{_ReduceAction, 0, _ReduceExplicitFieldDefsToOptionalExplicitFieldDefs}
	_ReduceNilToOptionalExplicitFieldDefsAction                                 = &_Action{_ReduceAction, 0, _ReduceNilToOptionalExplicitFieldDefs}
	_ReduceToExplicitStructDefAction                                            = &_Action{_ReduceAction, 0, _ReduceToExplicitStructDef}
	_ReducePairToImplicitEnumValueDefsAction                                    = &_Action{_ReduceAction, 0, _ReducePairToImplicitEnumValueDefs}
	_ReduceAddToImplicitEnumValueDefsAction                                     = &_Action{_ReduceAction, 0, _ReduceAddToImplicitEnumValueDefs}
	_ReduceToImplicitEnumDefAction                                              = &_Action{_ReduceAction, 0, _ReduceToImplicitEnumDef}
	_ReduceExplicitPairToExplicitEnumValueDefsAction                            = &_Action{_ReduceAction, 0, _ReduceExplicitPairToExplicitEnumValueDefs}
	_ReduceImplicitPairToExplicitEnumValueDefsAction                            = &_Action{_ReduceAction, 0, _ReduceImplicitPairToExplicitEnumValueDefs}
	_ReduceExplicitAddToExplicitEnumValueDefsAction                             = &_Action{_ReduceAction, 0, _ReduceExplicitAddToExplicitEnumValueDefs}
	_ReduceImplicitAddToExplicitEnumValueDefsAction                             = &_Action{_ReduceAction, 0, _ReduceImplicitAddToExplicitEnumValueDefs}
	_ReduceToExplicitEnumDefAction                                              = &_Action{_ReduceAction, 0, _ReduceToExplicitEnumDef}
	_ReduceFieldDefToTraitPropertyAction                                        = &_Action{_ReduceAction, 0, _ReduceFieldDefToTraitProperty}
	_ReduceMethodSignatureToTraitPropertyAction                                 = &_Action{_ReduceAction, 0, _ReduceMethodSignatureToTraitProperty}
	_ReduceTraitPropertyToTraitPropertiesAction                                 = &_Action{_ReduceAction, 0, _ReduceTraitPropertyToTraitProperties}
	_ReduceImplicitToTraitPropertiesAction                                      = &_Action{_ReduceAction, 0, _ReduceImplicitToTraitProperties}
	_ReduceExplicitToTraitPropertiesAction                                      = &_Action{_ReduceAction, 0, _ReduceExplicitToTraitProperties}
	_ReduceTraitPropertiesToOptionalTraitPropertiesAction                       = &_Action{_ReduceAction, 0, _ReduceTraitPropertiesToOptionalTraitProperties}
	_ReduceNilToOptionalTraitPropertiesAction                                   = &_Action{_ReduceAction, 0, _ReduceNilToOptionalTraitProperties}
	_ReduceToTraitDefAction                                                     = &_Action{_ReduceAction, 0, _ReduceToTraitDef}
	_ReduceValueTypeToReturnTypeAction                                          = &_Action{_ReduceAction, 0, _ReduceValueTypeToReturnType}
	_ReduceNilToReturnTypeAction                                                = &_Action{_ReduceAction, 0, _ReduceNilToReturnType}
	_ReduceArgToParameterDeclAction                                             = &_Action{_ReduceAction, 0, _ReduceArgToParameterDecl}
	_ReduceVarargToParameterDeclAction                                          = &_Action{_ReduceAction, 0, _ReduceVarargToParameterDecl}
	_ReduceUnamedToParameterDeclAction                                          = &_Action{_ReduceAction, 0, _ReduceUnamedToParameterDecl}
	_ReduceUnnamedVarargToParameterDeclAction                                   = &_Action{_ReduceAction, 0, _ReduceUnnamedVarargToParameterDecl}
	_ReduceParameterDeclToParameterDeclsAction                                  = &_Action{_ReduceAction, 0, _ReduceParameterDeclToParameterDecls}
	_ReduceAddToParameterDeclsAction                                            = &_Action{_ReduceAction, 0, _ReduceAddToParameterDecls}
	_ReduceParameterDeclsToOptionalParameterDeclsAction                         = &_Action{_ReduceAction, 0, _ReduceParameterDeclsToOptionalParameterDecls}
	_ReduceNilToOptionalParameterDeclsAction                                    = &_Action{_ReduceAction, 0, _ReduceNilToOptionalParameterDecls}
	_ReduceToFuncTypeAction                                                     = &_Action{_ReduceAction, 0, _ReduceToFuncType}
	_ReduceToMethodSignatureAction                                              = &_Action{_ReduceAction, 0, _ReduceToMethodSignature}
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
	{_State7, _EndMarker}:                               &_Action{_AcceptAction, 0, 0},
	{_State8, _EndMarker}:                               &_Action{_AcceptAction, 0, 0},
	{_State9, _EndMarker}:                               &_Action{_AcceptAction, 0, 0},
	{_State10, _EndMarker}:                              &_Action{_AcceptAction, 0, 0},
	{_State11, _EndMarker}:                              &_Action{_AcceptAction, 0, 0},
	{_State12, _EndMarker}:                              &_Action{_AcceptAction, 0, 0},
	{_State1, PackageToken}:                             _GotoState13Action,
	{_State1, PackageDefType}:                           _GotoState7Action,
	{_State2, TypeToken}:                                _GotoState14Action,
	{_State2, TypeDefType}:                              _GotoState8Action,
	{_State3, TraitToken}:                               _GotoState15Action,
	{_State3, TraitDefType}:                             _GotoState9Action,
	{_State4, FuncToken}:                                _GotoState16Action,
	{_State4, NamedFuncDefType}:                         _GotoState10Action,
	{_State5, IntegerLiteralToken}:                      _GotoState23Action,
	{_State5, FloatLiteralToken}:                        _GotoState20Action,
	{_State5, RuneLiteralToken}:                         _GotoState29Action,
	{_State5, StringLiteralToken}:                       _GotoState30Action,
	{_State5, IdentifierToken}:                          _GotoState22Action,
	{_State5, TrueToken}:                                _GotoState33Action,
	{_State5, FalseToken}:                               _GotoState19Action,
	{_State5, StructToken}:                              _GotoState31Action,
	{_State5, FuncToken}:                                _GotoState21Action,
	{_State5, LparenToken}:                              _GotoState26Action,
	{_State5, LabelDeclToken}:                           _GotoState24Action,
	{_State5, NotToken}:                                 _GotoState28Action,
	{_State5, SubToken}:                                 _GotoState32Action,
	{_State5, MulToken}:                                 _GotoState27Action,
	{_State5, BitNegToken}:                              _GotoState18Action,
	{_State5, BitAndToken}:                              _GotoState17Action,
	{_State5, LexErrorToken}:                            _GotoState25Action,
	{_State5, LiteralType}:                              _GotoState44Action,
	{_State5, AnonymousStructExprType}:                  _GotoState38Action,
	{_State5, AtomExprType}:                             _GotoState39Action,
	{_State5, CallExprType}:                             _GotoState41Action,
	{_State5, AccessExprType}:                           _GotoState34Action,
	{_State5, PostfixUnaryExprType}:                     _GotoState48Action,
	{_State5, PrefixUnaryOpType}:                        _GotoState50Action,
	{_State5, PrefixUnaryExprType}:                      _GotoState49Action,
	{_State5, MulExprType}:                              _GotoState45Action,
	{_State5, AddExprType}:                              _GotoState35Action,
	{_State5, CmpExprType}:                              _GotoState42Action,
	{_State5, AndExprType}:                              _GotoState36Action,
	{_State5, OrExprType}:                               _GotoState47Action,
	{_State5, SequenceExprType}:                         _GotoState51Action,
	{_State5, OptionalLabelDeclType}:                    _GotoState46Action,
	{_State5, BlockExprType}:                            _GotoState40Action,
	{_State5, ExpressionType}:                           _GotoState11Action,
	{_State5, ExplicitStructDefType}:                    _GotoState43Action,
	{_State5, AnonymousFuncExprType}:                    _GotoState37Action,
	{_State6, SpacesToken}:                              _GotoState53Action,
	{_State6, CommentToken}:                             _GotoState52Action,
	{_State6, LexInternalTokensType}:                    _GotoState12Action,
	{_State13, IdentifierToken}:                         _GotoState54Action,
	{_State14, IdentifierToken}:                         _GotoState55Action,
	{_State15, LparenToken}:                             _GotoState56Action,
	{_State16, LparenToken}:                             _GotoState57Action,
	{_State16, OptionalReceiverType}:                    _GotoState58Action,
	{_State21, LparenToken}:                             _GotoState59Action,
	{_State26, IntegerLiteralToken}:                     _GotoState23Action,
	{_State26, FloatLiteralToken}:                       _GotoState20Action,
	{_State26, RuneLiteralToken}:                        _GotoState29Action,
	{_State26, StringLiteralToken}:                      _GotoState30Action,
	{_State26, IdentifierToken}:                         _GotoState22Action,
	{_State26, TrueToken}:                               _GotoState33Action,
	{_State26, FalseToken}:                              _GotoState19Action,
	{_State26, StructToken}:                             _GotoState31Action,
	{_State26, FuncToken}:                               _GotoState21Action,
	{_State26, LparenToken}:                             _GotoState26Action,
	{_State26, LabelDeclToken}:                          _GotoState24Action,
	{_State26, NotToken}:                                _GotoState28Action,
	{_State26, SubToken}:                                _GotoState32Action,
	{_State26, MulToken}:                                _GotoState27Action,
	{_State26, BitNegToken}:                             _GotoState18Action,
	{_State26, BitAndToken}:                             _GotoState17Action,
	{_State26, LexErrorToken}:                           _GotoState25Action,
	{_State26, LiteralType}:                             _GotoState44Action,
	{_State26, AnonymousStructExprType}:                 _GotoState38Action,
	{_State26, AtomExprType}:                            _GotoState39Action,
	{_State26, ArgumentType}:                            _GotoState60Action,
	{_State26, ArgumentsType}:                           _GotoState61Action,
	{_State26, CallExprType}:                            _GotoState41Action,
	{_State26, AccessExprType}:                          _GotoState34Action,
	{_State26, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State26, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State26, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State26, MulExprType}:                             _GotoState45Action,
	{_State26, AddExprType}:                             _GotoState35Action,
	{_State26, CmpExprType}:                             _GotoState42Action,
	{_State26, AndExprType}:                             _GotoState36Action,
	{_State26, OrExprType}:                              _GotoState47Action,
	{_State26, SequenceExprType}:                        _GotoState51Action,
	{_State26, OptionalLabelDeclType}:                   _GotoState46Action,
	{_State26, BlockExprType}:                           _GotoState40Action,
	{_State26, ExpressionType}:                          _GotoState62Action,
	{_State26, ExplicitStructDefType}:                   _GotoState43Action,
	{_State26, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State31, LparenToken}:                             _GotoState63Action,
	{_State34, LbracketToken}:                           _GotoState66Action,
	{_State34, DotToken}:                                _GotoState65Action,
	{_State34, QuestionToken}:                           _GotoState67Action,
	{_State34, DollarLbracketToken}:                     _GotoState64Action,
	{_State34, OptionalGenericBindingType}:              _GotoState68Action,
	{_State35, AddToken}:                                _GotoState69Action,
	{_State35, SubToken}:                                _GotoState72Action,
	{_State35, BitXorToken}:                             _GotoState71Action,
	{_State35, BitOrToken}:                              _GotoState70Action,
	{_State35, AddOpType}:                               _GotoState73Action,
	{_State36, AndToken}:                                _GotoState74Action,
	{_State42, EqualToken}:                              _GotoState75Action,
	{_State42, NotEqualToken}:                           _GotoState80Action,
	{_State42, LessToken}:                               _GotoState78Action,
	{_State42, LessOrEqualToken}:                        _GotoState79Action,
	{_State42, GreaterToken}:                            _GotoState76Action,
	{_State42, GreaterOrEqualToken}:                     _GotoState77Action,
	{_State42, CmpOpType}:                               _GotoState81Action,
	{_State43, LparenToken}:                             _GotoState82Action,
	{_State45, MulToken}:                                _GotoState88Action,
	{_State45, DivToken}:                                _GotoState86Action,
	{_State45, ModToken}:                                _GotoState87Action,
	{_State45, BitAndToken}:                             _GotoState83Action,
	{_State45, BitLshiftToken}:                          _GotoState84Action,
	{_State45, BitRshiftToken}:                          _GotoState85Action,
	{_State45, MulOpType}:                               _GotoState89Action,
	{_State46, IfToken}:                                 _GotoState91Action,
	{_State46, SwitchToken}:                             _GotoState93Action,
	{_State46, ForToken}:                                _GotoState90Action,
	{_State46, LbraceToken}:                             _GotoState92Action,
	{_State46, BlockBodyType}:                           _GotoState94Action,
	{_State46, IfExprType}:                              _GotoState95Action,
	{_State46, SwitchExprType}:                          _GotoState97Action,
	{_State46, LoopExprType}:                            _GotoState96Action,
	{_State47, OrToken}:                                 _GotoState98Action,
	{_State50, IntegerLiteralToken}:                     _GotoState23Action,
	{_State50, FloatLiteralToken}:                       _GotoState20Action,
	{_State50, RuneLiteralToken}:                        _GotoState29Action,
	{_State50, StringLiteralToken}:                      _GotoState30Action,
	{_State50, IdentifierToken}:                         _GotoState22Action,
	{_State50, TrueToken}:                               _GotoState33Action,
	{_State50, FalseToken}:                              _GotoState19Action,
	{_State50, StructToken}:                             _GotoState31Action,
	{_State50, FuncToken}:                               _GotoState21Action,
	{_State50, LparenToken}:                             _GotoState26Action,
	{_State50, LabelDeclToken}:                          _GotoState24Action,
	{_State50, NotToken}:                                _GotoState28Action,
	{_State50, SubToken}:                                _GotoState32Action,
	{_State50, MulToken}:                                _GotoState27Action,
	{_State50, BitNegToken}:                             _GotoState18Action,
	{_State50, BitAndToken}:                             _GotoState17Action,
	{_State50, LexErrorToken}:                           _GotoState25Action,
	{_State50, LiteralType}:                             _GotoState44Action,
	{_State50, AnonymousStructExprType}:                 _GotoState38Action,
	{_State50, AtomExprType}:                            _GotoState39Action,
	{_State50, CallExprType}:                            _GotoState41Action,
	{_State50, AccessExprType}:                          _GotoState34Action,
	{_State50, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State50, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State50, PrefixUnaryExprType}:                     _GotoState100Action,
	{_State50, OptionalLabelDeclType}:                   _GotoState99Action,
	{_State50, BlockExprType}:                           _GotoState40Action,
	{_State50, ExplicitStructDefType}:                   _GotoState43Action,
	{_State50, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State54, LparenToken}:                             _GotoState101Action,
	{_State55, DollarLbracketToken}:                     _GotoState102Action,
	{_State55, EqualToken}:                              _GotoState103Action,
	{_State55, OptionalGenericParametersType}:           _GotoState104Action,
	{_State56, IdentifierToken}:                         _GotoState109Action,
	{_State56, StructToken}:                             _GotoState31Action,
	{_State56, EnumToken}:                               _GotoState107Action,
	{_State56, TraitToken}:                              _GotoState15Action,
	{_State56, FuncToken}:                               _GotoState108Action,
	{_State56, LparenToken}:                             _GotoState110Action,
	{_State56, QuestionToken}:                           _GotoState111Action,
	{_State56, TildeTildeToken}:                         _GotoState112Action,
	{_State56, BitNegToken}:                             _GotoState106Action,
	{_State56, BitAndToken}:                             _GotoState105Action,
	{_State56, AtomTypeType}:                            _GotoState113Action,
	{_State56, TraitableTypeType}:                       _GotoState127Action,
	{_State56, TraitMulTypeType}:                        _GotoState124Action,
	{_State56, TraitAddTypeType}:                        _GotoState122Action,
	{_State56, ValueTypeType}:                           _GotoState128Action,
	{_State56, FieldDefType}:                            _GotoState116Action,
	{_State56, ImplicitStructDefType}:                   _GotoState119Action,
	{_State56, ExplicitStructDefType}:                   _GotoState115Action,
	{_State56, ImplicitEnumDefType}:                     _GotoState118Action,
	{_State56, ExplicitEnumDefType}:                     _GotoState114Action,
	{_State56, TraitPropertyType}:                       _GotoState126Action,
	{_State56, TraitPropertiesType}:                     _GotoState125Action,
	{_State56, OptionalTraitPropertiesType}:             _GotoState121Action,
	{_State56, TraitDefType}:                            _GotoState123Action,
	{_State56, FuncTypeType}:                            _GotoState117Action,
	{_State56, MethodSignatureType}:                     _GotoState120Action,
	{_State57, IdentifierToken}:                         _GotoState129Action,
	{_State57, ParameterDefType}:                        _GotoState130Action,
	{_State58, IdentifierToken}:                         _GotoState131Action,
	{_State59, IdentifierToken}:                         _GotoState129Action,
	{_State59, ParameterDefType}:                        _GotoState133Action,
	{_State59, ParameterDefsType}:                       _GotoState134Action,
	{_State59, OptionalParameterDefsType}:               _GotoState132Action,
	{_State61, RparenToken}:                             _GotoState136Action,
	{_State61, CommaToken}:                              _GotoState135Action,
	{_State63, IdentifierToken}:                         _GotoState109Action,
	{_State63, StructToken}:                             _GotoState31Action,
	{_State63, EnumToken}:                               _GotoState107Action,
	{_State63, TraitToken}:                              _GotoState15Action,
	{_State63, FuncToken}:                               _GotoState137Action,
	{_State63, LparenToken}:                             _GotoState110Action,
	{_State63, QuestionToken}:                           _GotoState111Action,
	{_State63, TildeTildeToken}:                         _GotoState112Action,
	{_State63, BitNegToken}:                             _GotoState106Action,
	{_State63, BitAndToken}:                             _GotoState105Action,
	{_State63, AtomTypeType}:                            _GotoState113Action,
	{_State63, TraitableTypeType}:                       _GotoState127Action,
	{_State63, TraitMulTypeType}:                        _GotoState124Action,
	{_State63, TraitAddTypeType}:                        _GotoState122Action,
	{_State63, ValueTypeType}:                           _GotoState128Action,
	{_State63, FieldDefType}:                            _GotoState139Action,
	{_State63, ImplicitStructDefType}:                   _GotoState119Action,
	{_State63, ExplicitFieldDefsType}:                   _GotoState138Action,
	{_State63, OptionalExplicitFieldDefsType}:           _GotoState140Action,
	{_State63, ExplicitStructDefType}:                   _GotoState115Action,
	{_State63, ImplicitEnumDefType}:                     _GotoState118Action,
	{_State63, ExplicitEnumDefType}:                     _GotoState114Action,
	{_State63, TraitDefType}:                            _GotoState123Action,
	{_State63, FuncTypeType}:                            _GotoState117Action,
	{_State64, IdentifierToken}:                         _GotoState141Action,
	{_State64, StructToken}:                             _GotoState31Action,
	{_State64, EnumToken}:                               _GotoState107Action,
	{_State64, TraitToken}:                              _GotoState15Action,
	{_State64, FuncToken}:                               _GotoState137Action,
	{_State64, LparenToken}:                             _GotoState110Action,
	{_State64, QuestionToken}:                           _GotoState111Action,
	{_State64, TildeTildeToken}:                         _GotoState112Action,
	{_State64, BitNegToken}:                             _GotoState106Action,
	{_State64, BitAndToken}:                             _GotoState105Action,
	{_State64, GenericArgumentsType}:                    _GotoState142Action,
	{_State64, OptionalGenericArgumentsType}:            _GotoState143Action,
	{_State64, AtomTypeType}:                            _GotoState113Action,
	{_State64, TraitableTypeType}:                       _GotoState127Action,
	{_State64, TraitMulTypeType}:                        _GotoState124Action,
	{_State64, TraitAddTypeType}:                        _GotoState122Action,
	{_State64, ValueTypeType}:                           _GotoState144Action,
	{_State64, ImplicitStructDefType}:                   _GotoState119Action,
	{_State64, ExplicitStructDefType}:                   _GotoState115Action,
	{_State64, ImplicitEnumDefType}:                     _GotoState118Action,
	{_State64, ExplicitEnumDefType}:                     _GotoState114Action,
	{_State64, TraitDefType}:                            _GotoState123Action,
	{_State64, FuncTypeType}:                            _GotoState117Action,
	{_State65, IdentifierToken}:                         _GotoState145Action,
	{_State66, IntegerLiteralToken}:                     _GotoState23Action,
	{_State66, FloatLiteralToken}:                       _GotoState20Action,
	{_State66, RuneLiteralToken}:                        _GotoState29Action,
	{_State66, StringLiteralToken}:                      _GotoState30Action,
	{_State66, IdentifierToken}:                         _GotoState22Action,
	{_State66, TrueToken}:                               _GotoState33Action,
	{_State66, FalseToken}:                              _GotoState19Action,
	{_State66, StructToken}:                             _GotoState31Action,
	{_State66, FuncToken}:                               _GotoState21Action,
	{_State66, LparenToken}:                             _GotoState26Action,
	{_State66, LabelDeclToken}:                          _GotoState24Action,
	{_State66, NotToken}:                                _GotoState28Action,
	{_State66, SubToken}:                                _GotoState32Action,
	{_State66, MulToken}:                                _GotoState27Action,
	{_State66, BitNegToken}:                             _GotoState18Action,
	{_State66, BitAndToken}:                             _GotoState17Action,
	{_State66, LexErrorToken}:                           _GotoState25Action,
	{_State66, LiteralType}:                             _GotoState44Action,
	{_State66, AnonymousStructExprType}:                 _GotoState38Action,
	{_State66, AtomExprType}:                            _GotoState39Action,
	{_State66, CallExprType}:                            _GotoState41Action,
	{_State66, AccessExprType}:                          _GotoState34Action,
	{_State66, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State66, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State66, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State66, MulExprType}:                             _GotoState45Action,
	{_State66, AddExprType}:                             _GotoState35Action,
	{_State66, CmpExprType}:                             _GotoState42Action,
	{_State66, AndExprType}:                             _GotoState36Action,
	{_State66, OrExprType}:                              _GotoState47Action,
	{_State66, SequenceExprType}:                        _GotoState51Action,
	{_State66, OptionalLabelDeclType}:                   _GotoState46Action,
	{_State66, BlockExprType}:                           _GotoState40Action,
	{_State66, ExpressionType}:                          _GotoState146Action,
	{_State66, ExplicitStructDefType}:                   _GotoState43Action,
	{_State66, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State68, LparenToken}:                             _GotoState147Action,
	{_State73, IntegerLiteralToken}:                     _GotoState23Action,
	{_State73, FloatLiteralToken}:                       _GotoState20Action,
	{_State73, RuneLiteralToken}:                        _GotoState29Action,
	{_State73, StringLiteralToken}:                      _GotoState30Action,
	{_State73, IdentifierToken}:                         _GotoState22Action,
	{_State73, TrueToken}:                               _GotoState33Action,
	{_State73, FalseToken}:                              _GotoState19Action,
	{_State73, StructToken}:                             _GotoState31Action,
	{_State73, FuncToken}:                               _GotoState21Action,
	{_State73, LparenToken}:                             _GotoState26Action,
	{_State73, LabelDeclToken}:                          _GotoState24Action,
	{_State73, NotToken}:                                _GotoState28Action,
	{_State73, SubToken}:                                _GotoState32Action,
	{_State73, MulToken}:                                _GotoState27Action,
	{_State73, BitNegToken}:                             _GotoState18Action,
	{_State73, BitAndToken}:                             _GotoState17Action,
	{_State73, LexErrorToken}:                           _GotoState25Action,
	{_State73, LiteralType}:                             _GotoState44Action,
	{_State73, AnonymousStructExprType}:                 _GotoState38Action,
	{_State73, AtomExprType}:                            _GotoState39Action,
	{_State73, CallExprType}:                            _GotoState41Action,
	{_State73, AccessExprType}:                          _GotoState34Action,
	{_State73, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State73, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State73, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State73, MulExprType}:                             _GotoState148Action,
	{_State73, OptionalLabelDeclType}:                   _GotoState99Action,
	{_State73, BlockExprType}:                           _GotoState40Action,
	{_State73, ExplicitStructDefType}:                   _GotoState43Action,
	{_State73, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State74, IntegerLiteralToken}:                     _GotoState23Action,
	{_State74, FloatLiteralToken}:                       _GotoState20Action,
	{_State74, RuneLiteralToken}:                        _GotoState29Action,
	{_State74, StringLiteralToken}:                      _GotoState30Action,
	{_State74, IdentifierToken}:                         _GotoState22Action,
	{_State74, TrueToken}:                               _GotoState33Action,
	{_State74, FalseToken}:                              _GotoState19Action,
	{_State74, StructToken}:                             _GotoState31Action,
	{_State74, FuncToken}:                               _GotoState21Action,
	{_State74, LparenToken}:                             _GotoState26Action,
	{_State74, LabelDeclToken}:                          _GotoState24Action,
	{_State74, NotToken}:                                _GotoState28Action,
	{_State74, SubToken}:                                _GotoState32Action,
	{_State74, MulToken}:                                _GotoState27Action,
	{_State74, BitNegToken}:                             _GotoState18Action,
	{_State74, BitAndToken}:                             _GotoState17Action,
	{_State74, LexErrorToken}:                           _GotoState25Action,
	{_State74, LiteralType}:                             _GotoState44Action,
	{_State74, AnonymousStructExprType}:                 _GotoState38Action,
	{_State74, AtomExprType}:                            _GotoState39Action,
	{_State74, CallExprType}:                            _GotoState41Action,
	{_State74, AccessExprType}:                          _GotoState34Action,
	{_State74, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State74, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State74, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State74, MulExprType}:                             _GotoState45Action,
	{_State74, AddExprType}:                             _GotoState35Action,
	{_State74, CmpExprType}:                             _GotoState149Action,
	{_State74, OptionalLabelDeclType}:                   _GotoState99Action,
	{_State74, BlockExprType}:                           _GotoState40Action,
	{_State74, ExplicitStructDefType}:                   _GotoState43Action,
	{_State74, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State81, IntegerLiteralToken}:                     _GotoState23Action,
	{_State81, FloatLiteralToken}:                       _GotoState20Action,
	{_State81, RuneLiteralToken}:                        _GotoState29Action,
	{_State81, StringLiteralToken}:                      _GotoState30Action,
	{_State81, IdentifierToken}:                         _GotoState22Action,
	{_State81, TrueToken}:                               _GotoState33Action,
	{_State81, FalseToken}:                              _GotoState19Action,
	{_State81, StructToken}:                             _GotoState31Action,
	{_State81, FuncToken}:                               _GotoState21Action,
	{_State81, LparenToken}:                             _GotoState26Action,
	{_State81, LabelDeclToken}:                          _GotoState24Action,
	{_State81, NotToken}:                                _GotoState28Action,
	{_State81, SubToken}:                                _GotoState32Action,
	{_State81, MulToken}:                                _GotoState27Action,
	{_State81, BitNegToken}:                             _GotoState18Action,
	{_State81, BitAndToken}:                             _GotoState17Action,
	{_State81, LexErrorToken}:                           _GotoState25Action,
	{_State81, LiteralType}:                             _GotoState44Action,
	{_State81, AnonymousStructExprType}:                 _GotoState38Action,
	{_State81, AtomExprType}:                            _GotoState39Action,
	{_State81, CallExprType}:                            _GotoState41Action,
	{_State81, AccessExprType}:                          _GotoState34Action,
	{_State81, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State81, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State81, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State81, MulExprType}:                             _GotoState45Action,
	{_State81, AddExprType}:                             _GotoState150Action,
	{_State81, OptionalLabelDeclType}:                   _GotoState99Action,
	{_State81, BlockExprType}:                           _GotoState40Action,
	{_State81, ExplicitStructDefType}:                   _GotoState43Action,
	{_State81, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State82, IntegerLiteralToken}:                     _GotoState23Action,
	{_State82, FloatLiteralToken}:                       _GotoState20Action,
	{_State82, RuneLiteralToken}:                        _GotoState29Action,
	{_State82, StringLiteralToken}:                      _GotoState30Action,
	{_State82, IdentifierToken}:                         _GotoState22Action,
	{_State82, TrueToken}:                               _GotoState33Action,
	{_State82, FalseToken}:                              _GotoState19Action,
	{_State82, StructToken}:                             _GotoState31Action,
	{_State82, FuncToken}:                               _GotoState21Action,
	{_State82, LparenToken}:                             _GotoState26Action,
	{_State82, LabelDeclToken}:                          _GotoState24Action,
	{_State82, NotToken}:                                _GotoState28Action,
	{_State82, SubToken}:                                _GotoState32Action,
	{_State82, MulToken}:                                _GotoState27Action,
	{_State82, BitNegToken}:                             _GotoState18Action,
	{_State82, BitAndToken}:                             _GotoState17Action,
	{_State82, LexErrorToken}:                           _GotoState25Action,
	{_State82, LiteralType}:                             _GotoState44Action,
	{_State82, AnonymousStructExprType}:                 _GotoState38Action,
	{_State82, AtomExprType}:                            _GotoState39Action,
	{_State82, ArgumentType}:                            _GotoState60Action,
	{_State82, ArgumentsType}:                           _GotoState151Action,
	{_State82, CallExprType}:                            _GotoState41Action,
	{_State82, AccessExprType}:                          _GotoState34Action,
	{_State82, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State82, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State82, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State82, MulExprType}:                             _GotoState45Action,
	{_State82, AddExprType}:                             _GotoState35Action,
	{_State82, CmpExprType}:                             _GotoState42Action,
	{_State82, AndExprType}:                             _GotoState36Action,
	{_State82, OrExprType}:                              _GotoState47Action,
	{_State82, SequenceExprType}:                        _GotoState51Action,
	{_State82, OptionalLabelDeclType}:                   _GotoState46Action,
	{_State82, BlockExprType}:                           _GotoState40Action,
	{_State82, ExpressionType}:                          _GotoState62Action,
	{_State82, ExplicitStructDefType}:                   _GotoState43Action,
	{_State82, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State89, IntegerLiteralToken}:                     _GotoState23Action,
	{_State89, FloatLiteralToken}:                       _GotoState20Action,
	{_State89, RuneLiteralToken}:                        _GotoState29Action,
	{_State89, StringLiteralToken}:                      _GotoState30Action,
	{_State89, IdentifierToken}:                         _GotoState22Action,
	{_State89, TrueToken}:                               _GotoState33Action,
	{_State89, FalseToken}:                              _GotoState19Action,
	{_State89, StructToken}:                             _GotoState31Action,
	{_State89, FuncToken}:                               _GotoState21Action,
	{_State89, LparenToken}:                             _GotoState26Action,
	{_State89, LabelDeclToken}:                          _GotoState24Action,
	{_State89, NotToken}:                                _GotoState28Action,
	{_State89, SubToken}:                                _GotoState32Action,
	{_State89, MulToken}:                                _GotoState27Action,
	{_State89, BitNegToken}:                             _GotoState18Action,
	{_State89, BitAndToken}:                             _GotoState17Action,
	{_State89, LexErrorToken}:                           _GotoState25Action,
	{_State89, LiteralType}:                             _GotoState44Action,
	{_State89, AnonymousStructExprType}:                 _GotoState38Action,
	{_State89, AtomExprType}:                            _GotoState39Action,
	{_State89, CallExprType}:                            _GotoState41Action,
	{_State89, AccessExprType}:                          _GotoState34Action,
	{_State89, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State89, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State89, PrefixUnaryExprType}:                     _GotoState152Action,
	{_State89, OptionalLabelDeclType}:                   _GotoState99Action,
	{_State89, BlockExprType}:                           _GotoState40Action,
	{_State89, ExplicitStructDefType}:                   _GotoState43Action,
	{_State89, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State90, IntegerLiteralToken}:                     _GotoState23Action,
	{_State90, FloatLiteralToken}:                       _GotoState20Action,
	{_State90, RuneLiteralToken}:                        _GotoState29Action,
	{_State90, StringLiteralToken}:                      _GotoState30Action,
	{_State90, IdentifierToken}:                         _GotoState22Action,
	{_State90, TrueToken}:                               _GotoState33Action,
	{_State90, FalseToken}:                              _GotoState19Action,
	{_State90, StructToken}:                             _GotoState31Action,
	{_State90, FuncToken}:                               _GotoState21Action,
	{_State90, LparenToken}:                             _GotoState26Action,
	{_State90, LabelDeclToken}:                          _GotoState24Action,
	{_State90, NotToken}:                                _GotoState28Action,
	{_State90, SubToken}:                                _GotoState32Action,
	{_State90, MulToken}:                                _GotoState27Action,
	{_State90, BitNegToken}:                             _GotoState18Action,
	{_State90, BitAndToken}:                             _GotoState17Action,
	{_State90, LexErrorToken}:                           _GotoState25Action,
	{_State90, LiteralType}:                             _GotoState44Action,
	{_State90, AnonymousStructExprType}:                 _GotoState38Action,
	{_State90, AtomExprType}:                            _GotoState39Action,
	{_State90, CallExprType}:                            _GotoState41Action,
	{_State90, AccessExprType}:                          _GotoState34Action,
	{_State90, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State90, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State90, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State90, MulExprType}:                             _GotoState45Action,
	{_State90, AddExprType}:                             _GotoState35Action,
	{_State90, CmpExprType}:                             _GotoState42Action,
	{_State90, AndExprType}:                             _GotoState36Action,
	{_State90, OrExprType}:                              _GotoState47Action,
	{_State90, SequenceExprType}:                        _GotoState154Action,
	{_State90, OptionalLabelDeclType}:                   _GotoState99Action,
	{_State90, BlockExprType}:                           _GotoState153Action,
	{_State90, ExplicitStructDefType}:                   _GotoState43Action,
	{_State90, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State91, IntegerLiteralToken}:                     _GotoState23Action,
	{_State91, FloatLiteralToken}:                       _GotoState20Action,
	{_State91, RuneLiteralToken}:                        _GotoState29Action,
	{_State91, StringLiteralToken}:                      _GotoState30Action,
	{_State91, IdentifierToken}:                         _GotoState22Action,
	{_State91, TrueToken}:                               _GotoState33Action,
	{_State91, FalseToken}:                              _GotoState19Action,
	{_State91, StructToken}:                             _GotoState31Action,
	{_State91, FuncToken}:                               _GotoState21Action,
	{_State91, LparenToken}:                             _GotoState26Action,
	{_State91, LabelDeclToken}:                          _GotoState24Action,
	{_State91, NotToken}:                                _GotoState28Action,
	{_State91, SubToken}:                                _GotoState32Action,
	{_State91, MulToken}:                                _GotoState27Action,
	{_State91, BitNegToken}:                             _GotoState18Action,
	{_State91, BitAndToken}:                             _GotoState17Action,
	{_State91, LexErrorToken}:                           _GotoState25Action,
	{_State91, LiteralType}:                             _GotoState44Action,
	{_State91, AnonymousStructExprType}:                 _GotoState38Action,
	{_State91, AtomExprType}:                            _GotoState39Action,
	{_State91, CallExprType}:                            _GotoState41Action,
	{_State91, AccessExprType}:                          _GotoState34Action,
	{_State91, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State91, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State91, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State91, MulExprType}:                             _GotoState45Action,
	{_State91, AddExprType}:                             _GotoState35Action,
	{_State91, CmpExprType}:                             _GotoState42Action,
	{_State91, AndExprType}:                             _GotoState36Action,
	{_State91, OrExprType}:                              _GotoState47Action,
	{_State91, SequenceExprType}:                        _GotoState155Action,
	{_State91, OptionalLabelDeclType}:                   _GotoState99Action,
	{_State91, BlockExprType}:                           _GotoState40Action,
	{_State91, ExplicitStructDefType}:                   _GotoState43Action,
	{_State91, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State92, StatementsType}:                          _GotoState156Action,
	{_State93, LbraceToken}:                             _GotoState157Action,
	{_State98, IntegerLiteralToken}:                     _GotoState23Action,
	{_State98, FloatLiteralToken}:                       _GotoState20Action,
	{_State98, RuneLiteralToken}:                        _GotoState29Action,
	{_State98, StringLiteralToken}:                      _GotoState30Action,
	{_State98, IdentifierToken}:                         _GotoState22Action,
	{_State98, TrueToken}:                               _GotoState33Action,
	{_State98, FalseToken}:                              _GotoState19Action,
	{_State98, StructToken}:                             _GotoState31Action,
	{_State98, FuncToken}:                               _GotoState21Action,
	{_State98, LparenToken}:                             _GotoState26Action,
	{_State98, LabelDeclToken}:                          _GotoState24Action,
	{_State98, NotToken}:                                _GotoState28Action,
	{_State98, SubToken}:                                _GotoState32Action,
	{_State98, MulToken}:                                _GotoState27Action,
	{_State98, BitNegToken}:                             _GotoState18Action,
	{_State98, BitAndToken}:                             _GotoState17Action,
	{_State98, LexErrorToken}:                           _GotoState25Action,
	{_State98, LiteralType}:                             _GotoState44Action,
	{_State98, AnonymousStructExprType}:                 _GotoState38Action,
	{_State98, AtomExprType}:                            _GotoState39Action,
	{_State98, CallExprType}:                            _GotoState41Action,
	{_State98, AccessExprType}:                          _GotoState34Action,
	{_State98, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State98, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State98, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State98, MulExprType}:                             _GotoState45Action,
	{_State98, AddExprType}:                             _GotoState35Action,
	{_State98, CmpExprType}:                             _GotoState42Action,
	{_State98, AndExprType}:                             _GotoState158Action,
	{_State98, OptionalLabelDeclType}:                   _GotoState99Action,
	{_State98, BlockExprType}:                           _GotoState40Action,
	{_State98, ExplicitStructDefType}:                   _GotoState43Action,
	{_State98, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State99, LbraceToken}:                             _GotoState92Action,
	{_State99, BlockBodyType}:                           _GotoState94Action,
	{_State101, PackageStatementsType}:                  _GotoState159Action,
	{_State102, IdentifierToken}:                        _GotoState160Action,
	{_State102, GenericParameterDefType}:                _GotoState161Action,
	{_State102, GenericParameterDefsType}:               _GotoState162Action,
	{_State102, OptionalGenericParameterDefsType}:       _GotoState163Action,
	{_State103, IdentifierToken}:                        _GotoState141Action,
	{_State103, StructToken}:                            _GotoState31Action,
	{_State103, EnumToken}:                              _GotoState107Action,
	{_State103, TraitToken}:                             _GotoState15Action,
	{_State103, FuncToken}:                              _GotoState137Action,
	{_State103, LparenToken}:                            _GotoState110Action,
	{_State103, QuestionToken}:                          _GotoState111Action,
	{_State103, TildeTildeToken}:                        _GotoState112Action,
	{_State103, BitNegToken}:                            _GotoState106Action,
	{_State103, BitAndToken}:                            _GotoState105Action,
	{_State103, AtomTypeType}:                           _GotoState113Action,
	{_State103, TraitableTypeType}:                      _GotoState127Action,
	{_State103, TraitMulTypeType}:                       _GotoState124Action,
	{_State103, TraitAddTypeType}:                       _GotoState122Action,
	{_State103, ValueTypeType}:                          _GotoState164Action,
	{_State103, ImplicitStructDefType}:                  _GotoState119Action,
	{_State103, ExplicitStructDefType}:                  _GotoState115Action,
	{_State103, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State103, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State103, TraitDefType}:                           _GotoState123Action,
	{_State103, FuncTypeType}:                           _GotoState117Action,
	{_State104, IdentifierToken}:                        _GotoState141Action,
	{_State104, StructToken}:                            _GotoState31Action,
	{_State104, EnumToken}:                              _GotoState107Action,
	{_State104, TraitToken}:                             _GotoState15Action,
	{_State104, FuncToken}:                              _GotoState137Action,
	{_State104, LparenToken}:                            _GotoState110Action,
	{_State104, QuestionToken}:                          _GotoState111Action,
	{_State104, TildeTildeToken}:                        _GotoState112Action,
	{_State104, BitNegToken}:                            _GotoState106Action,
	{_State104, BitAndToken}:                            _GotoState105Action,
	{_State104, AtomTypeType}:                           _GotoState113Action,
	{_State104, TraitableTypeType}:                      _GotoState127Action,
	{_State104, TraitMulTypeType}:                       _GotoState124Action,
	{_State104, TraitAddTypeType}:                       _GotoState122Action,
	{_State104, ValueTypeType}:                          _GotoState165Action,
	{_State104, ImplicitStructDefType}:                  _GotoState119Action,
	{_State104, ExplicitStructDefType}:                  _GotoState115Action,
	{_State104, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State104, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State104, TraitDefType}:                           _GotoState123Action,
	{_State104, FuncTypeType}:                           _GotoState117Action,
	{_State105, IdentifierToken}:                        _GotoState141Action,
	{_State105, StructToken}:                            _GotoState31Action,
	{_State105, EnumToken}:                              _GotoState107Action,
	{_State105, TraitToken}:                             _GotoState15Action,
	{_State105, LparenToken}:                            _GotoState110Action,
	{_State105, QuestionToken}:                          _GotoState111Action,
	{_State105, TildeTildeToken}:                        _GotoState112Action,
	{_State105, BitNegToken}:                            _GotoState106Action,
	{_State105, AtomTypeType}:                           _GotoState113Action,
	{_State105, TraitableTypeType}:                      _GotoState127Action,
	{_State105, TraitMulTypeType}:                       _GotoState124Action,
	{_State105, TraitAddTypeType}:                       _GotoState166Action,
	{_State105, ImplicitStructDefType}:                  _GotoState119Action,
	{_State105, ExplicitStructDefType}:                  _GotoState115Action,
	{_State105, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State105, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State105, TraitDefType}:                           _GotoState123Action,
	{_State106, IdentifierToken}:                        _GotoState141Action,
	{_State106, StructToken}:                            _GotoState31Action,
	{_State106, EnumToken}:                              _GotoState107Action,
	{_State106, TraitToken}:                             _GotoState15Action,
	{_State106, LparenToken}:                            _GotoState110Action,
	{_State106, QuestionToken}:                          _GotoState111Action,
	{_State106, AtomTypeType}:                           _GotoState167Action,
	{_State106, ImplicitStructDefType}:                  _GotoState119Action,
	{_State106, ExplicitStructDefType}:                  _GotoState115Action,
	{_State106, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State106, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State106, TraitDefType}:                           _GotoState123Action,
	{_State107, LparenToken}:                            _GotoState168Action,
	{_State108, IdentifierToken}:                        _GotoState169Action,
	{_State108, LparenToken}:                            _GotoState170Action,
	{_State109, IdentifierToken}:                        _GotoState141Action,
	{_State109, StructToken}:                            _GotoState31Action,
	{_State109, EnumToken}:                              _GotoState107Action,
	{_State109, TraitToken}:                             _GotoState15Action,
	{_State109, FuncToken}:                              _GotoState137Action,
	{_State109, LparenToken}:                            _GotoState110Action,
	{_State109, QuestionToken}:                          _GotoState111Action,
	{_State109, DollarLbracketToken}:                    _GotoState64Action,
	{_State109, TildeTildeToken}:                        _GotoState112Action,
	{_State109, BitNegToken}:                            _GotoState106Action,
	{_State109, BitAndToken}:                            _GotoState105Action,
	{_State109, OptionalGenericBindingType}:             _GotoState171Action,
	{_State109, AtomTypeType}:                           _GotoState113Action,
	{_State109, TraitableTypeType}:                      _GotoState127Action,
	{_State109, TraitMulTypeType}:                       _GotoState124Action,
	{_State109, TraitAddTypeType}:                       _GotoState122Action,
	{_State109, ValueTypeType}:                          _GotoState172Action,
	{_State109, ImplicitStructDefType}:                  _GotoState119Action,
	{_State109, ExplicitStructDefType}:                  _GotoState115Action,
	{_State109, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State109, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State109, TraitDefType}:                           _GotoState123Action,
	{_State109, FuncTypeType}:                           _GotoState117Action,
	{_State110, IdentifierToken}:                        _GotoState109Action,
	{_State110, StructToken}:                            _GotoState31Action,
	{_State110, EnumToken}:                              _GotoState107Action,
	{_State110, TraitToken}:                             _GotoState15Action,
	{_State110, FuncToken}:                              _GotoState137Action,
	{_State110, LparenToken}:                            _GotoState110Action,
	{_State110, QuestionToken}:                          _GotoState111Action,
	{_State110, TildeTildeToken}:                        _GotoState112Action,
	{_State110, BitNegToken}:                            _GotoState106Action,
	{_State110, BitAndToken}:                            _GotoState105Action,
	{_State110, AtomTypeType}:                           _GotoState113Action,
	{_State110, TraitableTypeType}:                      _GotoState127Action,
	{_State110, TraitMulTypeType}:                       _GotoState124Action,
	{_State110, TraitAddTypeType}:                       _GotoState122Action,
	{_State110, ValueTypeType}:                          _GotoState128Action,
	{_State110, FieldDefType}:                           _GotoState173Action,
	{_State110, ImplicitFieldDefsType}:                  _GotoState175Action,
	{_State110, OptionalImplicitFieldDefsType}:          _GotoState176Action,
	{_State110, ImplicitStructDefType}:                  _GotoState119Action,
	{_State110, ExplicitStructDefType}:                  _GotoState115Action,
	{_State110, ImplicitEnumValueDefsType}:              _GotoState174Action,
	{_State110, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State110, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State110, TraitDefType}:                           _GotoState123Action,
	{_State110, FuncTypeType}:                           _GotoState117Action,
	{_State112, IdentifierToken}:                        _GotoState141Action,
	{_State112, StructToken}:                            _GotoState31Action,
	{_State112, EnumToken}:                              _GotoState107Action,
	{_State112, TraitToken}:                             _GotoState15Action,
	{_State112, LparenToken}:                            _GotoState110Action,
	{_State112, QuestionToken}:                          _GotoState111Action,
	{_State112, AtomTypeType}:                           _GotoState177Action,
	{_State112, ImplicitStructDefType}:                  _GotoState119Action,
	{_State112, ExplicitStructDefType}:                  _GotoState115Action,
	{_State112, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State112, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State112, TraitDefType}:                           _GotoState123Action,
	{_State121, RparenToken}:                            _GotoState178Action,
	{_State122, AddToken}:                               _GotoState179Action,
	{_State122, SubToken}:                               _GotoState180Action,
	{_State124, MulToken}:                               _GotoState181Action,
	{_State125, NewlinesToken}:                          _GotoState183Action,
	{_State125, CommaToken}:                             _GotoState182Action,
	{_State129, IdentifierToken}:                        _GotoState141Action,
	{_State129, StructToken}:                            _GotoState31Action,
	{_State129, EnumToken}:                              _GotoState107Action,
	{_State129, TraitToken}:                             _GotoState15Action,
	{_State129, FuncToken}:                              _GotoState137Action,
	{_State129, LparenToken}:                            _GotoState110Action,
	{_State129, QuestionToken}:                          _GotoState111Action,
	{_State129, DotdotdotToken}:                         _GotoState184Action,
	{_State129, TildeTildeToken}:                        _GotoState112Action,
	{_State129, BitNegToken}:                            _GotoState106Action,
	{_State129, BitAndToken}:                            _GotoState105Action,
	{_State129, AtomTypeType}:                           _GotoState113Action,
	{_State129, TraitableTypeType}:                      _GotoState127Action,
	{_State129, TraitMulTypeType}:                       _GotoState124Action,
	{_State129, TraitAddTypeType}:                       _GotoState122Action,
	{_State129, ValueTypeType}:                          _GotoState185Action,
	{_State129, ImplicitStructDefType}:                  _GotoState119Action,
	{_State129, ExplicitStructDefType}:                  _GotoState115Action,
	{_State129, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State129, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State129, TraitDefType}:                           _GotoState123Action,
	{_State129, FuncTypeType}:                           _GotoState117Action,
	{_State130, RparenToken}:                            _GotoState186Action,
	{_State131, DollarLbracketToken}:                    _GotoState102Action,
	{_State131, OptionalGenericParametersType}:          _GotoState187Action,
	{_State132, RparenToken}:                            _GotoState188Action,
	{_State134, CommaToken}:                             _GotoState189Action,
	{_State135, IntegerLiteralToken}:                    _GotoState23Action,
	{_State135, FloatLiteralToken}:                      _GotoState20Action,
	{_State135, RuneLiteralToken}:                       _GotoState29Action,
	{_State135, StringLiteralToken}:                     _GotoState30Action,
	{_State135, IdentifierToken}:                        _GotoState22Action,
	{_State135, TrueToken}:                              _GotoState33Action,
	{_State135, FalseToken}:                             _GotoState19Action,
	{_State135, StructToken}:                            _GotoState31Action,
	{_State135, FuncToken}:                              _GotoState21Action,
	{_State135, LparenToken}:                            _GotoState26Action,
	{_State135, LabelDeclToken}:                         _GotoState24Action,
	{_State135, NotToken}:                               _GotoState28Action,
	{_State135, SubToken}:                               _GotoState32Action,
	{_State135, MulToken}:                               _GotoState27Action,
	{_State135, BitNegToken}:                            _GotoState18Action,
	{_State135, BitAndToken}:                            _GotoState17Action,
	{_State135, LexErrorToken}:                          _GotoState25Action,
	{_State135, LiteralType}:                            _GotoState44Action,
	{_State135, AnonymousStructExprType}:                _GotoState38Action,
	{_State135, AtomExprType}:                           _GotoState39Action,
	{_State135, ArgumentType}:                           _GotoState190Action,
	{_State135, CallExprType}:                           _GotoState41Action,
	{_State135, AccessExprType}:                         _GotoState34Action,
	{_State135, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State135, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State135, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State135, MulExprType}:                            _GotoState45Action,
	{_State135, AddExprType}:                            _GotoState35Action,
	{_State135, CmpExprType}:                            _GotoState42Action,
	{_State135, AndExprType}:                            _GotoState36Action,
	{_State135, OrExprType}:                             _GotoState47Action,
	{_State135, SequenceExprType}:                       _GotoState51Action,
	{_State135, OptionalLabelDeclType}:                  _GotoState46Action,
	{_State135, BlockExprType}:                          _GotoState40Action,
	{_State135, ExpressionType}:                         _GotoState62Action,
	{_State135, ExplicitStructDefType}:                  _GotoState43Action,
	{_State135, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State137, LparenToken}:                            _GotoState170Action,
	{_State138, NewlinesToken}:                          _GotoState192Action,
	{_State138, CommaToken}:                             _GotoState191Action,
	{_State140, RparenToken}:                            _GotoState193Action,
	{_State141, DollarLbracketToken}:                    _GotoState64Action,
	{_State141, OptionalGenericBindingType}:             _GotoState171Action,
	{_State142, CommaToken}:                             _GotoState194Action,
	{_State143, RbracketToken}:                          _GotoState195Action,
	{_State146, RbracketToken}:                          _GotoState196Action,
	{_State147, IntegerLiteralToken}:                    _GotoState23Action,
	{_State147, FloatLiteralToken}:                      _GotoState20Action,
	{_State147, RuneLiteralToken}:                       _GotoState29Action,
	{_State147, StringLiteralToken}:                     _GotoState30Action,
	{_State147, IdentifierToken}:                        _GotoState22Action,
	{_State147, TrueToken}:                              _GotoState33Action,
	{_State147, FalseToken}:                             _GotoState19Action,
	{_State147, StructToken}:                            _GotoState31Action,
	{_State147, FuncToken}:                              _GotoState21Action,
	{_State147, LparenToken}:                            _GotoState26Action,
	{_State147, LabelDeclToken}:                         _GotoState24Action,
	{_State147, NotToken}:                               _GotoState28Action,
	{_State147, SubToken}:                               _GotoState32Action,
	{_State147, MulToken}:                               _GotoState27Action,
	{_State147, BitNegToken}:                            _GotoState18Action,
	{_State147, BitAndToken}:                            _GotoState17Action,
	{_State147, LexErrorToken}:                          _GotoState25Action,
	{_State147, LiteralType}:                            _GotoState44Action,
	{_State147, AnonymousStructExprType}:                _GotoState38Action,
	{_State147, AtomExprType}:                           _GotoState39Action,
	{_State147, ArgumentType}:                           _GotoState60Action,
	{_State147, ArgumentsType}:                          _GotoState197Action,
	{_State147, OptionalArgumentsType}:                  _GotoState198Action,
	{_State147, CallExprType}:                           _GotoState41Action,
	{_State147, AccessExprType}:                         _GotoState34Action,
	{_State147, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State147, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State147, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State147, MulExprType}:                            _GotoState45Action,
	{_State147, AddExprType}:                            _GotoState35Action,
	{_State147, CmpExprType}:                            _GotoState42Action,
	{_State147, AndExprType}:                            _GotoState36Action,
	{_State147, OrExprType}:                             _GotoState47Action,
	{_State147, SequenceExprType}:                       _GotoState51Action,
	{_State147, OptionalLabelDeclType}:                  _GotoState46Action,
	{_State147, BlockExprType}:                          _GotoState40Action,
	{_State147, ExpressionType}:                         _GotoState62Action,
	{_State147, ExplicitStructDefType}:                  _GotoState43Action,
	{_State147, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State148, MulToken}:                               _GotoState88Action,
	{_State148, DivToken}:                               _GotoState86Action,
	{_State148, ModToken}:                               _GotoState87Action,
	{_State148, BitAndToken}:                            _GotoState83Action,
	{_State148, BitLshiftToken}:                         _GotoState84Action,
	{_State148, BitRshiftToken}:                         _GotoState85Action,
	{_State148, MulOpType}:                              _GotoState89Action,
	{_State149, EqualToken}:                             _GotoState75Action,
	{_State149, NotEqualToken}:                          _GotoState80Action,
	{_State149, LessToken}:                              _GotoState78Action,
	{_State149, LessOrEqualToken}:                       _GotoState79Action,
	{_State149, GreaterToken}:                           _GotoState76Action,
	{_State149, GreaterOrEqualToken}:                    _GotoState77Action,
	{_State149, CmpOpType}:                              _GotoState81Action,
	{_State150, AddToken}:                               _GotoState69Action,
	{_State150, SubToken}:                               _GotoState72Action,
	{_State150, BitXorToken}:                            _GotoState71Action,
	{_State150, BitOrToken}:                             _GotoState70Action,
	{_State150, AddOpType}:                              _GotoState73Action,
	{_State151, RparenToken}:                            _GotoState199Action,
	{_State151, CommaToken}:                             _GotoState135Action,
	{_State154, LabelDeclToken}:                         _GotoState24Action,
	{_State154, OptionalLabelDeclType}:                  _GotoState99Action,
	{_State154, BlockExprType}:                          _GotoState200Action,
	{_State155, LbraceToken}:                            _GotoState92Action,
	{_State155, BlockBodyType}:                          _GotoState201Action,
	{_State156, IntegerLiteralToken}:                    _GotoState23Action,
	{_State156, FloatLiteralToken}:                      _GotoState20Action,
	{_State156, RuneLiteralToken}:                       _GotoState29Action,
	{_State156, StringLiteralToken}:                     _GotoState30Action,
	{_State156, IdentifierToken}:                        _GotoState22Action,
	{_State156, TrueToken}:                              _GotoState33Action,
	{_State156, FalseToken}:                             _GotoState19Action,
	{_State156, ReturnToken}:                            _GotoState206Action,
	{_State156, BreakToken}:                             _GotoState203Action,
	{_State156, ContinueToken}:                          _GotoState204Action,
	{_State156, UnsafeToken}:                            _GotoState207Action,
	{_State156, StructToken}:                            _GotoState31Action,
	{_State156, FuncToken}:                              _GotoState21Action,
	{_State156, AsyncToken}:                             _GotoState202Action,
	{_State156, RbraceToken}:                            _GotoState205Action,
	{_State156, LparenToken}:                            _GotoState26Action,
	{_State156, LabelDeclToken}:                         _GotoState24Action,
	{_State156, NotToken}:                               _GotoState28Action,
	{_State156, SubToken}:                               _GotoState32Action,
	{_State156, MulToken}:                               _GotoState27Action,
	{_State156, BitNegToken}:                            _GotoState18Action,
	{_State156, BitAndToken}:                            _GotoState17Action,
	{_State156, LexErrorToken}:                          _GotoState25Action,
	{_State156, LiteralType}:                            _GotoState44Action,
	{_State156, AnonymousStructExprType}:                _GotoState38Action,
	{_State156, AtomExprType}:                           _GotoState39Action,
	{_State156, CallExprType}:                           _GotoState41Action,
	{_State156, AccessExprType}:                         _GotoState208Action,
	{_State156, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State156, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State156, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State156, MulExprType}:                            _GotoState45Action,
	{_State156, AddExprType}:                            _GotoState35Action,
	{_State156, CmpExprType}:                            _GotoState42Action,
	{_State156, AndExprType}:                            _GotoState36Action,
	{_State156, OrExprType}:                             _GotoState47Action,
	{_State156, SequenceExprType}:                       _GotoState51Action,
	{_State156, JumpTypeType}:                           _GotoState211Action,
	{_State156, ExpressionOrImplicitStructType}:         _GotoState210Action,
	{_State156, UnsafeStatementType}:                    _GotoState214Action,
	{_State156, StatementBodyType}:                      _GotoState213Action,
	{_State156, StatementType}:                          _GotoState212Action,
	{_State156, OptionalLabelDeclType}:                  _GotoState46Action,
	{_State156, BlockExprType}:                          _GotoState40Action,
	{_State156, ExpressionType}:                         _GotoState209Action,
	{_State156, ExplicitStructDefType}:                  _GotoState43Action,
	{_State156, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State157, CaseToken}:                              _GotoState215Action,
	{_State158, AndToken}:                               _GotoState74Action,
	{_State159, UnsafeToken}:                            _GotoState207Action,
	{_State159, RparenToken}:                            _GotoState216Action,
	{_State159, UnsafeStatementType}:                    _GotoState219Action,
	{_State159, PackageStatementBodyType}:               _GotoState218Action,
	{_State159, PackageStatementType}:                   _GotoState217Action,
	{_State160, IdentifierToken}:                        _GotoState141Action,
	{_State160, StructToken}:                            _GotoState31Action,
	{_State160, EnumToken}:                              _GotoState107Action,
	{_State160, TraitToken}:                             _GotoState15Action,
	{_State160, FuncToken}:                              _GotoState137Action,
	{_State160, LparenToken}:                            _GotoState110Action,
	{_State160, QuestionToken}:                          _GotoState111Action,
	{_State160, TildeTildeToken}:                        _GotoState112Action,
	{_State160, BitNegToken}:                            _GotoState106Action,
	{_State160, BitAndToken}:                            _GotoState105Action,
	{_State160, AtomTypeType}:                           _GotoState113Action,
	{_State160, TraitableTypeType}:                      _GotoState127Action,
	{_State160, TraitMulTypeType}:                       _GotoState124Action,
	{_State160, TraitAddTypeType}:                       _GotoState122Action,
	{_State160, ValueTypeType}:                          _GotoState220Action,
	{_State160, ImplicitStructDefType}:                  _GotoState119Action,
	{_State160, ExplicitStructDefType}:                  _GotoState115Action,
	{_State160, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State160, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State160, TraitDefType}:                           _GotoState123Action,
	{_State160, FuncTypeType}:                           _GotoState117Action,
	{_State162, CommaToken}:                             _GotoState221Action,
	{_State163, RbracketToken}:                          _GotoState222Action,
	{_State165, ImplementsToken}:                        _GotoState223Action,
	{_State166, AddToken}:                               _GotoState179Action,
	{_State166, SubToken}:                               _GotoState180Action,
	{_State168, IdentifierToken}:                        _GotoState109Action,
	{_State168, StructToken}:                            _GotoState31Action,
	{_State168, EnumToken}:                              _GotoState107Action,
	{_State168, TraitToken}:                             _GotoState15Action,
	{_State168, FuncToken}:                              _GotoState137Action,
	{_State168, LparenToken}:                            _GotoState110Action,
	{_State168, QuestionToken}:                          _GotoState111Action,
	{_State168, TildeTildeToken}:                        _GotoState112Action,
	{_State168, BitNegToken}:                            _GotoState106Action,
	{_State168, BitAndToken}:                            _GotoState105Action,
	{_State168, AtomTypeType}:                           _GotoState113Action,
	{_State168, TraitableTypeType}:                      _GotoState127Action,
	{_State168, TraitMulTypeType}:                       _GotoState124Action,
	{_State168, TraitAddTypeType}:                       _GotoState122Action,
	{_State168, ValueTypeType}:                          _GotoState128Action,
	{_State168, FieldDefType}:                           _GotoState225Action,
	{_State168, ImplicitStructDefType}:                  _GotoState119Action,
	{_State168, ExplicitStructDefType}:                  _GotoState115Action,
	{_State168, ImplicitEnumValueDefsType}:              _GotoState226Action,
	{_State168, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State168, ExplicitEnumValueDefsType}:              _GotoState224Action,
	{_State168, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State168, TraitDefType}:                           _GotoState123Action,
	{_State168, FuncTypeType}:                           _GotoState117Action,
	{_State169, LparenToken}:                            _GotoState227Action,
	{_State170, IdentifierToken}:                        _GotoState229Action,
	{_State170, StructToken}:                            _GotoState31Action,
	{_State170, EnumToken}:                              _GotoState107Action,
	{_State170, TraitToken}:                             _GotoState15Action,
	{_State170, FuncToken}:                              _GotoState137Action,
	{_State170, LparenToken}:                            _GotoState110Action,
	{_State170, QuestionToken}:                          _GotoState111Action,
	{_State170, DotdotdotToken}:                         _GotoState228Action,
	{_State170, TildeTildeToken}:                        _GotoState112Action,
	{_State170, BitNegToken}:                            _GotoState106Action,
	{_State170, BitAndToken}:                            _GotoState105Action,
	{_State170, AtomTypeType}:                           _GotoState113Action,
	{_State170, TraitableTypeType}:                      _GotoState127Action,
	{_State170, TraitMulTypeType}:                       _GotoState124Action,
	{_State170, TraitAddTypeType}:                       _GotoState122Action,
	{_State170, ValueTypeType}:                          _GotoState233Action,
	{_State170, ImplicitStructDefType}:                  _GotoState119Action,
	{_State170, ExplicitStructDefType}:                  _GotoState115Action,
	{_State170, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State170, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State170, TraitDefType}:                           _GotoState123Action,
	{_State170, ParameterDeclType}:                      _GotoState231Action,
	{_State170, ParameterDeclsType}:                     _GotoState232Action,
	{_State170, OptionalParameterDeclsType}:             _GotoState230Action,
	{_State170, FuncTypeType}:                           _GotoState117Action,
	{_State173, OrToken}:                                _GotoState234Action,
	{_State174, RparenToken}:                            _GotoState236Action,
	{_State174, OrToken}:                                _GotoState235Action,
	{_State175, CommaToken}:                             _GotoState237Action,
	{_State176, RparenToken}:                            _GotoState238Action,
	{_State179, IdentifierToken}:                        _GotoState141Action,
	{_State179, StructToken}:                            _GotoState31Action,
	{_State179, EnumToken}:                              _GotoState107Action,
	{_State179, TraitToken}:                             _GotoState15Action,
	{_State179, LparenToken}:                            _GotoState110Action,
	{_State179, QuestionToken}:                          _GotoState111Action,
	{_State179, TildeTildeToken}:                        _GotoState112Action,
	{_State179, BitNegToken}:                            _GotoState106Action,
	{_State179, AtomTypeType}:                           _GotoState113Action,
	{_State179, TraitableTypeType}:                      _GotoState127Action,
	{_State179, TraitMulTypeType}:                       _GotoState239Action,
	{_State179, ImplicitStructDefType}:                  _GotoState119Action,
	{_State179, ExplicitStructDefType}:                  _GotoState115Action,
	{_State179, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State179, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State179, TraitDefType}:                           _GotoState123Action,
	{_State180, IdentifierToken}:                        _GotoState141Action,
	{_State180, StructToken}:                            _GotoState31Action,
	{_State180, EnumToken}:                              _GotoState107Action,
	{_State180, TraitToken}:                             _GotoState15Action,
	{_State180, LparenToken}:                            _GotoState110Action,
	{_State180, QuestionToken}:                          _GotoState111Action,
	{_State180, TildeTildeToken}:                        _GotoState112Action,
	{_State180, BitNegToken}:                            _GotoState106Action,
	{_State180, AtomTypeType}:                           _GotoState113Action,
	{_State180, TraitableTypeType}:                      _GotoState127Action,
	{_State180, TraitMulTypeType}:                       _GotoState240Action,
	{_State180, ImplicitStructDefType}:                  _GotoState119Action,
	{_State180, ExplicitStructDefType}:                  _GotoState115Action,
	{_State180, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State180, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State180, TraitDefType}:                           _GotoState123Action,
	{_State181, IdentifierToken}:                        _GotoState141Action,
	{_State181, StructToken}:                            _GotoState31Action,
	{_State181, EnumToken}:                              _GotoState107Action,
	{_State181, TraitToken}:                             _GotoState15Action,
	{_State181, LparenToken}:                            _GotoState110Action,
	{_State181, QuestionToken}:                          _GotoState111Action,
	{_State181, TildeTildeToken}:                        _GotoState112Action,
	{_State181, BitNegToken}:                            _GotoState106Action,
	{_State181, AtomTypeType}:                           _GotoState113Action,
	{_State181, TraitableTypeType}:                      _GotoState241Action,
	{_State181, ImplicitStructDefType}:                  _GotoState119Action,
	{_State181, ExplicitStructDefType}:                  _GotoState115Action,
	{_State181, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State181, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State181, TraitDefType}:                           _GotoState123Action,
	{_State182, IdentifierToken}:                        _GotoState109Action,
	{_State182, StructToken}:                            _GotoState31Action,
	{_State182, EnumToken}:                              _GotoState107Action,
	{_State182, TraitToken}:                             _GotoState15Action,
	{_State182, FuncToken}:                              _GotoState108Action,
	{_State182, LparenToken}:                            _GotoState110Action,
	{_State182, QuestionToken}:                          _GotoState111Action,
	{_State182, TildeTildeToken}:                        _GotoState112Action,
	{_State182, BitNegToken}:                            _GotoState106Action,
	{_State182, BitAndToken}:                            _GotoState105Action,
	{_State182, AtomTypeType}:                           _GotoState113Action,
	{_State182, TraitableTypeType}:                      _GotoState127Action,
	{_State182, TraitMulTypeType}:                       _GotoState124Action,
	{_State182, TraitAddTypeType}:                       _GotoState122Action,
	{_State182, ValueTypeType}:                          _GotoState128Action,
	{_State182, FieldDefType}:                           _GotoState116Action,
	{_State182, ImplicitStructDefType}:                  _GotoState119Action,
	{_State182, ExplicitStructDefType}:                  _GotoState115Action,
	{_State182, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State182, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State182, TraitPropertyType}:                      _GotoState242Action,
	{_State182, TraitDefType}:                           _GotoState123Action,
	{_State182, FuncTypeType}:                           _GotoState117Action,
	{_State182, MethodSignatureType}:                    _GotoState120Action,
	{_State183, IdentifierToken}:                        _GotoState109Action,
	{_State183, StructToken}:                            _GotoState31Action,
	{_State183, EnumToken}:                              _GotoState107Action,
	{_State183, TraitToken}:                             _GotoState15Action,
	{_State183, FuncToken}:                              _GotoState108Action,
	{_State183, LparenToken}:                            _GotoState110Action,
	{_State183, QuestionToken}:                          _GotoState111Action,
	{_State183, TildeTildeToken}:                        _GotoState112Action,
	{_State183, BitNegToken}:                            _GotoState106Action,
	{_State183, BitAndToken}:                            _GotoState105Action,
	{_State183, AtomTypeType}:                           _GotoState113Action,
	{_State183, TraitableTypeType}:                      _GotoState127Action,
	{_State183, TraitMulTypeType}:                       _GotoState124Action,
	{_State183, TraitAddTypeType}:                       _GotoState122Action,
	{_State183, ValueTypeType}:                          _GotoState128Action,
	{_State183, FieldDefType}:                           _GotoState116Action,
	{_State183, ImplicitStructDefType}:                  _GotoState119Action,
	{_State183, ExplicitStructDefType}:                  _GotoState115Action,
	{_State183, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State183, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State183, TraitPropertyType}:                      _GotoState243Action,
	{_State183, TraitDefType}:                           _GotoState123Action,
	{_State183, FuncTypeType}:                           _GotoState117Action,
	{_State183, MethodSignatureType}:                    _GotoState120Action,
	{_State184, IdentifierToken}:                        _GotoState141Action,
	{_State184, StructToken}:                            _GotoState31Action,
	{_State184, EnumToken}:                              _GotoState107Action,
	{_State184, TraitToken}:                             _GotoState15Action,
	{_State184, FuncToken}:                              _GotoState137Action,
	{_State184, LparenToken}:                            _GotoState110Action,
	{_State184, QuestionToken}:                          _GotoState111Action,
	{_State184, TildeTildeToken}:                        _GotoState112Action,
	{_State184, BitNegToken}:                            _GotoState106Action,
	{_State184, BitAndToken}:                            _GotoState105Action,
	{_State184, AtomTypeType}:                           _GotoState113Action,
	{_State184, TraitableTypeType}:                      _GotoState127Action,
	{_State184, TraitMulTypeType}:                       _GotoState124Action,
	{_State184, TraitAddTypeType}:                       _GotoState122Action,
	{_State184, ValueTypeType}:                          _GotoState244Action,
	{_State184, ImplicitStructDefType}:                  _GotoState119Action,
	{_State184, ExplicitStructDefType}:                  _GotoState115Action,
	{_State184, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State184, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State184, TraitDefType}:                           _GotoState123Action,
	{_State184, FuncTypeType}:                           _GotoState117Action,
	{_State187, LparenToken}:                            _GotoState245Action,
	{_State188, IdentifierToken}:                        _GotoState141Action,
	{_State188, StructToken}:                            _GotoState31Action,
	{_State188, EnumToken}:                              _GotoState107Action,
	{_State188, TraitToken}:                             _GotoState15Action,
	{_State188, FuncToken}:                              _GotoState137Action,
	{_State188, LparenToken}:                            _GotoState110Action,
	{_State188, QuestionToken}:                          _GotoState111Action,
	{_State188, TildeTildeToken}:                        _GotoState112Action,
	{_State188, BitNegToken}:                            _GotoState106Action,
	{_State188, BitAndToken}:                            _GotoState105Action,
	{_State188, AtomTypeType}:                           _GotoState113Action,
	{_State188, TraitableTypeType}:                      _GotoState127Action,
	{_State188, TraitMulTypeType}:                       _GotoState124Action,
	{_State188, TraitAddTypeType}:                       _GotoState122Action,
	{_State188, ValueTypeType}:                          _GotoState247Action,
	{_State188, ImplicitStructDefType}:                  _GotoState119Action,
	{_State188, ExplicitStructDefType}:                  _GotoState115Action,
	{_State188, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State188, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State188, TraitDefType}:                           _GotoState123Action,
	{_State188, ReturnTypeType}:                         _GotoState246Action,
	{_State188, FuncTypeType}:                           _GotoState117Action,
	{_State189, IdentifierToken}:                        _GotoState129Action,
	{_State189, ParameterDefType}:                       _GotoState248Action,
	{_State191, IdentifierToken}:                        _GotoState109Action,
	{_State191, StructToken}:                            _GotoState31Action,
	{_State191, EnumToken}:                              _GotoState107Action,
	{_State191, TraitToken}:                             _GotoState15Action,
	{_State191, FuncToken}:                              _GotoState137Action,
	{_State191, LparenToken}:                            _GotoState110Action,
	{_State191, QuestionToken}:                          _GotoState111Action,
	{_State191, TildeTildeToken}:                        _GotoState112Action,
	{_State191, BitNegToken}:                            _GotoState106Action,
	{_State191, BitAndToken}:                            _GotoState105Action,
	{_State191, AtomTypeType}:                           _GotoState113Action,
	{_State191, TraitableTypeType}:                      _GotoState127Action,
	{_State191, TraitMulTypeType}:                       _GotoState124Action,
	{_State191, TraitAddTypeType}:                       _GotoState122Action,
	{_State191, ValueTypeType}:                          _GotoState128Action,
	{_State191, FieldDefType}:                           _GotoState249Action,
	{_State191, ImplicitStructDefType}:                  _GotoState119Action,
	{_State191, ExplicitStructDefType}:                  _GotoState115Action,
	{_State191, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State191, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State191, TraitDefType}:                           _GotoState123Action,
	{_State191, FuncTypeType}:                           _GotoState117Action,
	{_State192, IdentifierToken}:                        _GotoState109Action,
	{_State192, StructToken}:                            _GotoState31Action,
	{_State192, EnumToken}:                              _GotoState107Action,
	{_State192, TraitToken}:                             _GotoState15Action,
	{_State192, FuncToken}:                              _GotoState137Action,
	{_State192, LparenToken}:                            _GotoState110Action,
	{_State192, QuestionToken}:                          _GotoState111Action,
	{_State192, TildeTildeToken}:                        _GotoState112Action,
	{_State192, BitNegToken}:                            _GotoState106Action,
	{_State192, BitAndToken}:                            _GotoState105Action,
	{_State192, AtomTypeType}:                           _GotoState113Action,
	{_State192, TraitableTypeType}:                      _GotoState127Action,
	{_State192, TraitMulTypeType}:                       _GotoState124Action,
	{_State192, TraitAddTypeType}:                       _GotoState122Action,
	{_State192, ValueTypeType}:                          _GotoState128Action,
	{_State192, FieldDefType}:                           _GotoState250Action,
	{_State192, ImplicitStructDefType}:                  _GotoState119Action,
	{_State192, ExplicitStructDefType}:                  _GotoState115Action,
	{_State192, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State192, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State192, TraitDefType}:                           _GotoState123Action,
	{_State192, FuncTypeType}:                           _GotoState117Action,
	{_State194, IdentifierToken}:                        _GotoState141Action,
	{_State194, StructToken}:                            _GotoState31Action,
	{_State194, EnumToken}:                              _GotoState107Action,
	{_State194, TraitToken}:                             _GotoState15Action,
	{_State194, FuncToken}:                              _GotoState137Action,
	{_State194, LparenToken}:                            _GotoState110Action,
	{_State194, QuestionToken}:                          _GotoState111Action,
	{_State194, TildeTildeToken}:                        _GotoState112Action,
	{_State194, BitNegToken}:                            _GotoState106Action,
	{_State194, BitAndToken}:                            _GotoState105Action,
	{_State194, AtomTypeType}:                           _GotoState113Action,
	{_State194, TraitableTypeType}:                      _GotoState127Action,
	{_State194, TraitMulTypeType}:                       _GotoState124Action,
	{_State194, TraitAddTypeType}:                       _GotoState122Action,
	{_State194, ValueTypeType}:                          _GotoState251Action,
	{_State194, ImplicitStructDefType}:                  _GotoState119Action,
	{_State194, ExplicitStructDefType}:                  _GotoState115Action,
	{_State194, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State194, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State194, TraitDefType}:                           _GotoState123Action,
	{_State194, FuncTypeType}:                           _GotoState117Action,
	{_State197, CommaToken}:                             _GotoState135Action,
	{_State198, RparenToken}:                            _GotoState252Action,
	{_State201, ElseToken}:                              _GotoState253Action,
	{_State202, IntegerLiteralToken}:                    _GotoState23Action,
	{_State202, FloatLiteralToken}:                      _GotoState20Action,
	{_State202, RuneLiteralToken}:                       _GotoState29Action,
	{_State202, StringLiteralToken}:                     _GotoState30Action,
	{_State202, IdentifierToken}:                        _GotoState22Action,
	{_State202, TrueToken}:                              _GotoState33Action,
	{_State202, FalseToken}:                             _GotoState19Action,
	{_State202, StructToken}:                            _GotoState31Action,
	{_State202, FuncToken}:                              _GotoState21Action,
	{_State202, LparenToken}:                            _GotoState26Action,
	{_State202, LabelDeclToken}:                         _GotoState24Action,
	{_State202, LexErrorToken}:                          _GotoState25Action,
	{_State202, LiteralType}:                            _GotoState44Action,
	{_State202, AnonymousStructExprType}:                _GotoState38Action,
	{_State202, AtomExprType}:                           _GotoState39Action,
	{_State202, CallExprType}:                           _GotoState255Action,
	{_State202, AccessExprType}:                         _GotoState254Action,
	{_State202, OptionalLabelDeclType}:                  _GotoState99Action,
	{_State202, BlockExprType}:                          _GotoState40Action,
	{_State202, ExplicitStructDefType}:                  _GotoState43Action,
	{_State202, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State207, LessToken}:                              _GotoState256Action,
	{_State208, LbracketToken}:                          _GotoState66Action,
	{_State208, DotToken}:                               _GotoState65Action,
	{_State208, QuestionToken}:                          _GotoState67Action,
	{_State208, DollarLbracketToken}:                    _GotoState64Action,
	{_State208, AddAssignToken}:                         _GotoState257Action,
	{_State208, SubAssignToken}:                         _GotoState268Action,
	{_State208, MulAssignToken}:                         _GotoState267Action,
	{_State208, DivAssignToken}:                         _GotoState265Action,
	{_State208, ModAssignToken}:                         _GotoState266Action,
	{_State208, AddOneAssignToken}:                      _GotoState258Action,
	{_State208, SubOneAssignToken}:                      _GotoState269Action,
	{_State208, BitNegAssignToken}:                      _GotoState261Action,
	{_State208, BitAndAssignToken}:                      _GotoState259Action,
	{_State208, BitOrAssignToken}:                       _GotoState262Action,
	{_State208, BitXorAssignToken}:                      _GotoState264Action,
	{_State208, BitLshiftAssignToken}:                   _GotoState260Action,
	{_State208, BitRshiftAssignToken}:                   _GotoState263Action,
	{_State208, OptionalGenericBindingType}:             _GotoState68Action,
	{_State208, OpOneAssignType}:                        _GotoState271Action,
	{_State208, BinaryOpAssignType}:                     _GotoState270Action,
	{_State210, CommaToken}:                             _GotoState272Action,
	{_State211, JumpLabelToken}:                         _GotoState273Action,
	{_State211, OptionalJumpLabelType}:                  _GotoState274Action,
	{_State213, NewlinesToken}:                          _GotoState275Action,
	{_State213, SemicolonToken}:                         _GotoState276Action,
	{_State215, DefaultToken}:                           _GotoState277Action,
	{_State218, NewlinesToken}:                          _GotoState278Action,
	{_State218, SemicolonToken}:                         _GotoState279Action,
	{_State221, IdentifierToken}:                        _GotoState160Action,
	{_State221, GenericParameterDefType}:                _GotoState280Action,
	{_State223, IdentifierToken}:                        _GotoState141Action,
	{_State223, StructToken}:                            _GotoState31Action,
	{_State223, EnumToken}:                              _GotoState107Action,
	{_State223, TraitToken}:                             _GotoState15Action,
	{_State223, FuncToken}:                              _GotoState137Action,
	{_State223, LparenToken}:                            _GotoState110Action,
	{_State223, QuestionToken}:                          _GotoState111Action,
	{_State223, TildeTildeToken}:                        _GotoState112Action,
	{_State223, BitNegToken}:                            _GotoState106Action,
	{_State223, BitAndToken}:                            _GotoState105Action,
	{_State223, AtomTypeType}:                           _GotoState113Action,
	{_State223, TraitableTypeType}:                      _GotoState127Action,
	{_State223, TraitMulTypeType}:                       _GotoState124Action,
	{_State223, TraitAddTypeType}:                       _GotoState122Action,
	{_State223, ValueTypeType}:                          _GotoState281Action,
	{_State223, ImplicitStructDefType}:                  _GotoState119Action,
	{_State223, ExplicitStructDefType}:                  _GotoState115Action,
	{_State223, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State223, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State223, TraitDefType}:                           _GotoState123Action,
	{_State223, FuncTypeType}:                           _GotoState117Action,
	{_State224, RparenToken}:                            _GotoState282Action,
	{_State225, NewlinesToken}:                          _GotoState283Action,
	{_State225, OrToken}:                                _GotoState284Action,
	{_State226, NewlinesToken}:                          _GotoState285Action,
	{_State226, OrToken}:                                _GotoState286Action,
	{_State227, IdentifierToken}:                        _GotoState229Action,
	{_State227, StructToken}:                            _GotoState31Action,
	{_State227, EnumToken}:                              _GotoState107Action,
	{_State227, TraitToken}:                             _GotoState15Action,
	{_State227, FuncToken}:                              _GotoState137Action,
	{_State227, LparenToken}:                            _GotoState110Action,
	{_State227, QuestionToken}:                          _GotoState111Action,
	{_State227, DotdotdotToken}:                         _GotoState228Action,
	{_State227, TildeTildeToken}:                        _GotoState112Action,
	{_State227, BitNegToken}:                            _GotoState106Action,
	{_State227, BitAndToken}:                            _GotoState105Action,
	{_State227, AtomTypeType}:                           _GotoState113Action,
	{_State227, TraitableTypeType}:                      _GotoState127Action,
	{_State227, TraitMulTypeType}:                       _GotoState124Action,
	{_State227, TraitAddTypeType}:                       _GotoState122Action,
	{_State227, ValueTypeType}:                          _GotoState233Action,
	{_State227, ImplicitStructDefType}:                  _GotoState119Action,
	{_State227, ExplicitStructDefType}:                  _GotoState115Action,
	{_State227, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State227, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State227, TraitDefType}:                           _GotoState123Action,
	{_State227, ParameterDeclType}:                      _GotoState231Action,
	{_State227, ParameterDeclsType}:                     _GotoState232Action,
	{_State227, OptionalParameterDeclsType}:             _GotoState287Action,
	{_State227, FuncTypeType}:                           _GotoState117Action,
	{_State228, IdentifierToken}:                        _GotoState141Action,
	{_State228, StructToken}:                            _GotoState31Action,
	{_State228, EnumToken}:                              _GotoState107Action,
	{_State228, TraitToken}:                             _GotoState15Action,
	{_State228, FuncToken}:                              _GotoState137Action,
	{_State228, LparenToken}:                            _GotoState110Action,
	{_State228, QuestionToken}:                          _GotoState111Action,
	{_State228, TildeTildeToken}:                        _GotoState112Action,
	{_State228, BitNegToken}:                            _GotoState106Action,
	{_State228, BitAndToken}:                            _GotoState105Action,
	{_State228, AtomTypeType}:                           _GotoState113Action,
	{_State228, TraitableTypeType}:                      _GotoState127Action,
	{_State228, TraitMulTypeType}:                       _GotoState124Action,
	{_State228, TraitAddTypeType}:                       _GotoState122Action,
	{_State228, ValueTypeType}:                          _GotoState288Action,
	{_State228, ImplicitStructDefType}:                  _GotoState119Action,
	{_State228, ExplicitStructDefType}:                  _GotoState115Action,
	{_State228, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State228, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State228, TraitDefType}:                           _GotoState123Action,
	{_State228, FuncTypeType}:                           _GotoState117Action,
	{_State229, IdentifierToken}:                        _GotoState141Action,
	{_State229, StructToken}:                            _GotoState31Action,
	{_State229, EnumToken}:                              _GotoState107Action,
	{_State229, TraitToken}:                             _GotoState15Action,
	{_State229, FuncToken}:                              _GotoState137Action,
	{_State229, LparenToken}:                            _GotoState110Action,
	{_State229, QuestionToken}:                          _GotoState111Action,
	{_State229, DollarLbracketToken}:                    _GotoState64Action,
	{_State229, DotdotdotToken}:                         _GotoState289Action,
	{_State229, TildeTildeToken}:                        _GotoState112Action,
	{_State229, BitNegToken}:                            _GotoState106Action,
	{_State229, BitAndToken}:                            _GotoState105Action,
	{_State229, OptionalGenericBindingType}:             _GotoState171Action,
	{_State229, AtomTypeType}:                           _GotoState113Action,
	{_State229, TraitableTypeType}:                      _GotoState127Action,
	{_State229, TraitMulTypeType}:                       _GotoState124Action,
	{_State229, TraitAddTypeType}:                       _GotoState122Action,
	{_State229, ValueTypeType}:                          _GotoState290Action,
	{_State229, ImplicitStructDefType}:                  _GotoState119Action,
	{_State229, ExplicitStructDefType}:                  _GotoState115Action,
	{_State229, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State229, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State229, TraitDefType}:                           _GotoState123Action,
	{_State229, FuncTypeType}:                           _GotoState117Action,
	{_State230, RparenToken}:                            _GotoState291Action,
	{_State232, CommaToken}:                             _GotoState292Action,
	{_State234, IdentifierToken}:                        _GotoState109Action,
	{_State234, StructToken}:                            _GotoState31Action,
	{_State234, EnumToken}:                              _GotoState107Action,
	{_State234, TraitToken}:                             _GotoState15Action,
	{_State234, FuncToken}:                              _GotoState137Action,
	{_State234, LparenToken}:                            _GotoState110Action,
	{_State234, QuestionToken}:                          _GotoState111Action,
	{_State234, TildeTildeToken}:                        _GotoState112Action,
	{_State234, BitNegToken}:                            _GotoState106Action,
	{_State234, BitAndToken}:                            _GotoState105Action,
	{_State234, AtomTypeType}:                           _GotoState113Action,
	{_State234, TraitableTypeType}:                      _GotoState127Action,
	{_State234, TraitMulTypeType}:                       _GotoState124Action,
	{_State234, TraitAddTypeType}:                       _GotoState122Action,
	{_State234, ValueTypeType}:                          _GotoState128Action,
	{_State234, FieldDefType}:                           _GotoState293Action,
	{_State234, ImplicitStructDefType}:                  _GotoState119Action,
	{_State234, ExplicitStructDefType}:                  _GotoState115Action,
	{_State234, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State234, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State234, TraitDefType}:                           _GotoState123Action,
	{_State234, FuncTypeType}:                           _GotoState117Action,
	{_State235, IdentifierToken}:                        _GotoState109Action,
	{_State235, StructToken}:                            _GotoState31Action,
	{_State235, EnumToken}:                              _GotoState107Action,
	{_State235, TraitToken}:                             _GotoState15Action,
	{_State235, FuncToken}:                              _GotoState137Action,
	{_State235, LparenToken}:                            _GotoState110Action,
	{_State235, QuestionToken}:                          _GotoState111Action,
	{_State235, TildeTildeToken}:                        _GotoState112Action,
	{_State235, BitNegToken}:                            _GotoState106Action,
	{_State235, BitAndToken}:                            _GotoState105Action,
	{_State235, AtomTypeType}:                           _GotoState113Action,
	{_State235, TraitableTypeType}:                      _GotoState127Action,
	{_State235, TraitMulTypeType}:                       _GotoState124Action,
	{_State235, TraitAddTypeType}:                       _GotoState122Action,
	{_State235, ValueTypeType}:                          _GotoState128Action,
	{_State235, FieldDefType}:                           _GotoState294Action,
	{_State235, ImplicitStructDefType}:                  _GotoState119Action,
	{_State235, ExplicitStructDefType}:                  _GotoState115Action,
	{_State235, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State235, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State235, TraitDefType}:                           _GotoState123Action,
	{_State235, FuncTypeType}:                           _GotoState117Action,
	{_State237, IdentifierToken}:                        _GotoState109Action,
	{_State237, StructToken}:                            _GotoState31Action,
	{_State237, EnumToken}:                              _GotoState107Action,
	{_State237, TraitToken}:                             _GotoState15Action,
	{_State237, FuncToken}:                              _GotoState137Action,
	{_State237, LparenToken}:                            _GotoState110Action,
	{_State237, QuestionToken}:                          _GotoState111Action,
	{_State237, TildeTildeToken}:                        _GotoState112Action,
	{_State237, BitNegToken}:                            _GotoState106Action,
	{_State237, BitAndToken}:                            _GotoState105Action,
	{_State237, AtomTypeType}:                           _GotoState113Action,
	{_State237, TraitableTypeType}:                      _GotoState127Action,
	{_State237, TraitMulTypeType}:                       _GotoState124Action,
	{_State237, TraitAddTypeType}:                       _GotoState122Action,
	{_State237, ValueTypeType}:                          _GotoState128Action,
	{_State237, FieldDefType}:                           _GotoState295Action,
	{_State237, ImplicitStructDefType}:                  _GotoState119Action,
	{_State237, ExplicitStructDefType}:                  _GotoState115Action,
	{_State237, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State237, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State237, TraitDefType}:                           _GotoState123Action,
	{_State237, FuncTypeType}:                           _GotoState117Action,
	{_State239, MulToken}:                               _GotoState181Action,
	{_State240, MulToken}:                               _GotoState181Action,
	{_State245, IdentifierToken}:                        _GotoState129Action,
	{_State245, ParameterDefType}:                       _GotoState133Action,
	{_State245, ParameterDefsType}:                      _GotoState134Action,
	{_State245, OptionalParameterDefsType}:              _GotoState296Action,
	{_State246, LbraceToken}:                            _GotoState92Action,
	{_State246, BlockBodyType}:                          _GotoState297Action,
	{_State253, IfToken}:                                _GotoState91Action,
	{_State253, LbraceToken}:                            _GotoState92Action,
	{_State253, BlockBodyType}:                          _GotoState298Action,
	{_State253, IfExprType}:                             _GotoState299Action,
	{_State254, LbracketToken}:                          _GotoState66Action,
	{_State254, DotToken}:                               _GotoState65Action,
	{_State254, DollarLbracketToken}:                    _GotoState64Action,
	{_State254, OptionalGenericBindingType}:             _GotoState68Action,
	{_State256, IdentifierToken}:                        _GotoState300Action,
	{_State270, IntegerLiteralToken}:                    _GotoState23Action,
	{_State270, FloatLiteralToken}:                      _GotoState20Action,
	{_State270, RuneLiteralToken}:                       _GotoState29Action,
	{_State270, StringLiteralToken}:                     _GotoState30Action,
	{_State270, IdentifierToken}:                        _GotoState22Action,
	{_State270, TrueToken}:                              _GotoState33Action,
	{_State270, FalseToken}:                             _GotoState19Action,
	{_State270, StructToken}:                            _GotoState31Action,
	{_State270, FuncToken}:                              _GotoState21Action,
	{_State270, LparenToken}:                            _GotoState26Action,
	{_State270, LabelDeclToken}:                         _GotoState24Action,
	{_State270, NotToken}:                               _GotoState28Action,
	{_State270, SubToken}:                               _GotoState32Action,
	{_State270, MulToken}:                               _GotoState27Action,
	{_State270, BitNegToken}:                            _GotoState18Action,
	{_State270, BitAndToken}:                            _GotoState17Action,
	{_State270, LexErrorToken}:                          _GotoState25Action,
	{_State270, LiteralType}:                            _GotoState44Action,
	{_State270, AnonymousStructExprType}:                _GotoState38Action,
	{_State270, AtomExprType}:                           _GotoState39Action,
	{_State270, CallExprType}:                           _GotoState41Action,
	{_State270, AccessExprType}:                         _GotoState34Action,
	{_State270, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State270, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State270, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State270, MulExprType}:                            _GotoState45Action,
	{_State270, AddExprType}:                            _GotoState35Action,
	{_State270, CmpExprType}:                            _GotoState42Action,
	{_State270, AndExprType}:                            _GotoState36Action,
	{_State270, OrExprType}:                             _GotoState47Action,
	{_State270, SequenceExprType}:                       _GotoState51Action,
	{_State270, OptionalLabelDeclType}:                  _GotoState46Action,
	{_State270, BlockExprType}:                          _GotoState40Action,
	{_State270, ExpressionType}:                         _GotoState301Action,
	{_State270, ExplicitStructDefType}:                  _GotoState43Action,
	{_State270, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State272, IntegerLiteralToken}:                    _GotoState23Action,
	{_State272, FloatLiteralToken}:                      _GotoState20Action,
	{_State272, RuneLiteralToken}:                       _GotoState29Action,
	{_State272, StringLiteralToken}:                     _GotoState30Action,
	{_State272, IdentifierToken}:                        _GotoState22Action,
	{_State272, TrueToken}:                              _GotoState33Action,
	{_State272, FalseToken}:                             _GotoState19Action,
	{_State272, StructToken}:                            _GotoState31Action,
	{_State272, FuncToken}:                              _GotoState21Action,
	{_State272, LparenToken}:                            _GotoState26Action,
	{_State272, LabelDeclToken}:                         _GotoState24Action,
	{_State272, NotToken}:                               _GotoState28Action,
	{_State272, SubToken}:                               _GotoState32Action,
	{_State272, MulToken}:                               _GotoState27Action,
	{_State272, BitNegToken}:                            _GotoState18Action,
	{_State272, BitAndToken}:                            _GotoState17Action,
	{_State272, LexErrorToken}:                          _GotoState25Action,
	{_State272, LiteralType}:                            _GotoState44Action,
	{_State272, AnonymousStructExprType}:                _GotoState38Action,
	{_State272, AtomExprType}:                           _GotoState39Action,
	{_State272, CallExprType}:                           _GotoState41Action,
	{_State272, AccessExprType}:                         _GotoState34Action,
	{_State272, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State272, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State272, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State272, MulExprType}:                            _GotoState45Action,
	{_State272, AddExprType}:                            _GotoState35Action,
	{_State272, CmpExprType}:                            _GotoState42Action,
	{_State272, AndExprType}:                            _GotoState36Action,
	{_State272, OrExprType}:                             _GotoState47Action,
	{_State272, SequenceExprType}:                       _GotoState51Action,
	{_State272, OptionalLabelDeclType}:                  _GotoState46Action,
	{_State272, BlockExprType}:                          _GotoState40Action,
	{_State272, ExpressionType}:                         _GotoState302Action,
	{_State272, ExplicitStructDefType}:                  _GotoState43Action,
	{_State272, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State274, IntegerLiteralToken}:                    _GotoState23Action,
	{_State274, FloatLiteralToken}:                      _GotoState20Action,
	{_State274, RuneLiteralToken}:                       _GotoState29Action,
	{_State274, StringLiteralToken}:                     _GotoState30Action,
	{_State274, IdentifierToken}:                        _GotoState22Action,
	{_State274, TrueToken}:                              _GotoState33Action,
	{_State274, FalseToken}:                             _GotoState19Action,
	{_State274, StructToken}:                            _GotoState31Action,
	{_State274, FuncToken}:                              _GotoState21Action,
	{_State274, LparenToken}:                            _GotoState26Action,
	{_State274, LabelDeclToken}:                         _GotoState24Action,
	{_State274, NotToken}:                               _GotoState28Action,
	{_State274, SubToken}:                               _GotoState32Action,
	{_State274, MulToken}:                               _GotoState27Action,
	{_State274, BitNegToken}:                            _GotoState18Action,
	{_State274, BitAndToken}:                            _GotoState17Action,
	{_State274, LexErrorToken}:                          _GotoState25Action,
	{_State274, LiteralType}:                            _GotoState44Action,
	{_State274, AnonymousStructExprType}:                _GotoState38Action,
	{_State274, AtomExprType}:                           _GotoState39Action,
	{_State274, CallExprType}:                           _GotoState41Action,
	{_State274, AccessExprType}:                         _GotoState34Action,
	{_State274, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State274, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State274, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State274, MulExprType}:                            _GotoState45Action,
	{_State274, AddExprType}:                            _GotoState35Action,
	{_State274, CmpExprType}:                            _GotoState42Action,
	{_State274, AndExprType}:                            _GotoState36Action,
	{_State274, OrExprType}:                             _GotoState47Action,
	{_State274, SequenceExprType}:                       _GotoState51Action,
	{_State274, OptionalExpressionOrImplicitStructType}: _GotoState304Action,
	{_State274, ExpressionOrImplicitStructType}:         _GotoState303Action,
	{_State274, OptionalLabelDeclType}:                  _GotoState46Action,
	{_State274, BlockExprType}:                          _GotoState40Action,
	{_State274, ExpressionType}:                         _GotoState209Action,
	{_State274, ExplicitStructDefType}:                  _GotoState43Action,
	{_State274, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State277, RbraceToken}:                            _GotoState305Action,
	{_State283, IdentifierToken}:                        _GotoState109Action,
	{_State283, StructToken}:                            _GotoState31Action,
	{_State283, EnumToken}:                              _GotoState107Action,
	{_State283, TraitToken}:                             _GotoState15Action,
	{_State283, FuncToken}:                              _GotoState137Action,
	{_State283, LparenToken}:                            _GotoState110Action,
	{_State283, QuestionToken}:                          _GotoState111Action,
	{_State283, TildeTildeToken}:                        _GotoState112Action,
	{_State283, BitNegToken}:                            _GotoState106Action,
	{_State283, BitAndToken}:                            _GotoState105Action,
	{_State283, AtomTypeType}:                           _GotoState113Action,
	{_State283, TraitableTypeType}:                      _GotoState127Action,
	{_State283, TraitMulTypeType}:                       _GotoState124Action,
	{_State283, TraitAddTypeType}:                       _GotoState122Action,
	{_State283, ValueTypeType}:                          _GotoState128Action,
	{_State283, FieldDefType}:                           _GotoState306Action,
	{_State283, ImplicitStructDefType}:                  _GotoState119Action,
	{_State283, ExplicitStructDefType}:                  _GotoState115Action,
	{_State283, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State283, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State283, TraitDefType}:                           _GotoState123Action,
	{_State283, FuncTypeType}:                           _GotoState117Action,
	{_State284, IdentifierToken}:                        _GotoState109Action,
	{_State284, StructToken}:                            _GotoState31Action,
	{_State284, EnumToken}:                              _GotoState107Action,
	{_State284, TraitToken}:                             _GotoState15Action,
	{_State284, FuncToken}:                              _GotoState137Action,
	{_State284, LparenToken}:                            _GotoState110Action,
	{_State284, QuestionToken}:                          _GotoState111Action,
	{_State284, TildeTildeToken}:                        _GotoState112Action,
	{_State284, BitNegToken}:                            _GotoState106Action,
	{_State284, BitAndToken}:                            _GotoState105Action,
	{_State284, AtomTypeType}:                           _GotoState113Action,
	{_State284, TraitableTypeType}:                      _GotoState127Action,
	{_State284, TraitMulTypeType}:                       _GotoState124Action,
	{_State284, TraitAddTypeType}:                       _GotoState122Action,
	{_State284, ValueTypeType}:                          _GotoState128Action,
	{_State284, FieldDefType}:                           _GotoState307Action,
	{_State284, ImplicitStructDefType}:                  _GotoState119Action,
	{_State284, ExplicitStructDefType}:                  _GotoState115Action,
	{_State284, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State284, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State284, TraitDefType}:                           _GotoState123Action,
	{_State284, FuncTypeType}:                           _GotoState117Action,
	{_State285, IdentifierToken}:                        _GotoState109Action,
	{_State285, StructToken}:                            _GotoState31Action,
	{_State285, EnumToken}:                              _GotoState107Action,
	{_State285, TraitToken}:                             _GotoState15Action,
	{_State285, FuncToken}:                              _GotoState137Action,
	{_State285, LparenToken}:                            _GotoState110Action,
	{_State285, QuestionToken}:                          _GotoState111Action,
	{_State285, TildeTildeToken}:                        _GotoState112Action,
	{_State285, BitNegToken}:                            _GotoState106Action,
	{_State285, BitAndToken}:                            _GotoState105Action,
	{_State285, AtomTypeType}:                           _GotoState113Action,
	{_State285, TraitableTypeType}:                      _GotoState127Action,
	{_State285, TraitMulTypeType}:                       _GotoState124Action,
	{_State285, TraitAddTypeType}:                       _GotoState122Action,
	{_State285, ValueTypeType}:                          _GotoState128Action,
	{_State285, FieldDefType}:                           _GotoState308Action,
	{_State285, ImplicitStructDefType}:                  _GotoState119Action,
	{_State285, ExplicitStructDefType}:                  _GotoState115Action,
	{_State285, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State285, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State285, TraitDefType}:                           _GotoState123Action,
	{_State285, FuncTypeType}:                           _GotoState117Action,
	{_State286, IdentifierToken}:                        _GotoState109Action,
	{_State286, StructToken}:                            _GotoState31Action,
	{_State286, EnumToken}:                              _GotoState107Action,
	{_State286, TraitToken}:                             _GotoState15Action,
	{_State286, FuncToken}:                              _GotoState137Action,
	{_State286, LparenToken}:                            _GotoState110Action,
	{_State286, QuestionToken}:                          _GotoState111Action,
	{_State286, TildeTildeToken}:                        _GotoState112Action,
	{_State286, BitNegToken}:                            _GotoState106Action,
	{_State286, BitAndToken}:                            _GotoState105Action,
	{_State286, AtomTypeType}:                           _GotoState113Action,
	{_State286, TraitableTypeType}:                      _GotoState127Action,
	{_State286, TraitMulTypeType}:                       _GotoState124Action,
	{_State286, TraitAddTypeType}:                       _GotoState122Action,
	{_State286, ValueTypeType}:                          _GotoState128Action,
	{_State286, FieldDefType}:                           _GotoState309Action,
	{_State286, ImplicitStructDefType}:                  _GotoState119Action,
	{_State286, ExplicitStructDefType}:                  _GotoState115Action,
	{_State286, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State286, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State286, TraitDefType}:                           _GotoState123Action,
	{_State286, FuncTypeType}:                           _GotoState117Action,
	{_State287, RparenToken}:                            _GotoState310Action,
	{_State289, IdentifierToken}:                        _GotoState141Action,
	{_State289, StructToken}:                            _GotoState31Action,
	{_State289, EnumToken}:                              _GotoState107Action,
	{_State289, TraitToken}:                             _GotoState15Action,
	{_State289, FuncToken}:                              _GotoState137Action,
	{_State289, LparenToken}:                            _GotoState110Action,
	{_State289, QuestionToken}:                          _GotoState111Action,
	{_State289, TildeTildeToken}:                        _GotoState112Action,
	{_State289, BitNegToken}:                            _GotoState106Action,
	{_State289, BitAndToken}:                            _GotoState105Action,
	{_State289, AtomTypeType}:                           _GotoState113Action,
	{_State289, TraitableTypeType}:                      _GotoState127Action,
	{_State289, TraitMulTypeType}:                       _GotoState124Action,
	{_State289, TraitAddTypeType}:                       _GotoState122Action,
	{_State289, ValueTypeType}:                          _GotoState311Action,
	{_State289, ImplicitStructDefType}:                  _GotoState119Action,
	{_State289, ExplicitStructDefType}:                  _GotoState115Action,
	{_State289, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State289, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State289, TraitDefType}:                           _GotoState123Action,
	{_State289, FuncTypeType}:                           _GotoState117Action,
	{_State291, IdentifierToken}:                        _GotoState141Action,
	{_State291, StructToken}:                            _GotoState31Action,
	{_State291, EnumToken}:                              _GotoState107Action,
	{_State291, TraitToken}:                             _GotoState15Action,
	{_State291, FuncToken}:                              _GotoState137Action,
	{_State291, LparenToken}:                            _GotoState110Action,
	{_State291, QuestionToken}:                          _GotoState111Action,
	{_State291, TildeTildeToken}:                        _GotoState112Action,
	{_State291, BitNegToken}:                            _GotoState106Action,
	{_State291, BitAndToken}:                            _GotoState105Action,
	{_State291, AtomTypeType}:                           _GotoState113Action,
	{_State291, TraitableTypeType}:                      _GotoState127Action,
	{_State291, TraitMulTypeType}:                       _GotoState124Action,
	{_State291, TraitAddTypeType}:                       _GotoState122Action,
	{_State291, ValueTypeType}:                          _GotoState247Action,
	{_State291, ImplicitStructDefType}:                  _GotoState119Action,
	{_State291, ExplicitStructDefType}:                  _GotoState115Action,
	{_State291, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State291, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State291, TraitDefType}:                           _GotoState123Action,
	{_State291, ReturnTypeType}:                         _GotoState312Action,
	{_State291, FuncTypeType}:                           _GotoState117Action,
	{_State292, IdentifierToken}:                        _GotoState229Action,
	{_State292, StructToken}:                            _GotoState31Action,
	{_State292, EnumToken}:                              _GotoState107Action,
	{_State292, TraitToken}:                             _GotoState15Action,
	{_State292, FuncToken}:                              _GotoState137Action,
	{_State292, LparenToken}:                            _GotoState110Action,
	{_State292, QuestionToken}:                          _GotoState111Action,
	{_State292, DotdotdotToken}:                         _GotoState228Action,
	{_State292, TildeTildeToken}:                        _GotoState112Action,
	{_State292, BitNegToken}:                            _GotoState106Action,
	{_State292, BitAndToken}:                            _GotoState105Action,
	{_State292, AtomTypeType}:                           _GotoState113Action,
	{_State292, TraitableTypeType}:                      _GotoState127Action,
	{_State292, TraitMulTypeType}:                       _GotoState124Action,
	{_State292, TraitAddTypeType}:                       _GotoState122Action,
	{_State292, ValueTypeType}:                          _GotoState233Action,
	{_State292, ImplicitStructDefType}:                  _GotoState119Action,
	{_State292, ExplicitStructDefType}:                  _GotoState115Action,
	{_State292, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State292, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State292, TraitDefType}:                           _GotoState123Action,
	{_State292, ParameterDeclType}:                      _GotoState313Action,
	{_State292, FuncTypeType}:                           _GotoState117Action,
	{_State296, RparenToken}:                            _GotoState314Action,
	{_State300, GreaterToken}:                           _GotoState315Action,
	{_State303, CommaToken}:                             _GotoState272Action,
	{_State310, IdentifierToken}:                        _GotoState141Action,
	{_State310, StructToken}:                            _GotoState31Action,
	{_State310, EnumToken}:                              _GotoState107Action,
	{_State310, TraitToken}:                             _GotoState15Action,
	{_State310, FuncToken}:                              _GotoState137Action,
	{_State310, LparenToken}:                            _GotoState110Action,
	{_State310, QuestionToken}:                          _GotoState111Action,
	{_State310, TildeTildeToken}:                        _GotoState112Action,
	{_State310, BitNegToken}:                            _GotoState106Action,
	{_State310, BitAndToken}:                            _GotoState105Action,
	{_State310, AtomTypeType}:                           _GotoState113Action,
	{_State310, TraitableTypeType}:                      _GotoState127Action,
	{_State310, TraitMulTypeType}:                       _GotoState124Action,
	{_State310, TraitAddTypeType}:                       _GotoState122Action,
	{_State310, ValueTypeType}:                          _GotoState247Action,
	{_State310, ImplicitStructDefType}:                  _GotoState119Action,
	{_State310, ExplicitStructDefType}:                  _GotoState115Action,
	{_State310, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State310, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State310, TraitDefType}:                           _GotoState123Action,
	{_State310, ReturnTypeType}:                         _GotoState316Action,
	{_State310, FuncTypeType}:                           _GotoState117Action,
	{_State314, IdentifierToken}:                        _GotoState141Action,
	{_State314, StructToken}:                            _GotoState31Action,
	{_State314, EnumToken}:                              _GotoState107Action,
	{_State314, TraitToken}:                             _GotoState15Action,
	{_State314, FuncToken}:                              _GotoState137Action,
	{_State314, LparenToken}:                            _GotoState110Action,
	{_State314, QuestionToken}:                          _GotoState111Action,
	{_State314, TildeTildeToken}:                        _GotoState112Action,
	{_State314, BitNegToken}:                            _GotoState106Action,
	{_State314, BitAndToken}:                            _GotoState105Action,
	{_State314, AtomTypeType}:                           _GotoState113Action,
	{_State314, TraitableTypeType}:                      _GotoState127Action,
	{_State314, TraitMulTypeType}:                       _GotoState124Action,
	{_State314, TraitAddTypeType}:                       _GotoState122Action,
	{_State314, ValueTypeType}:                          _GotoState247Action,
	{_State314, ImplicitStructDefType}:                  _GotoState119Action,
	{_State314, ExplicitStructDefType}:                  _GotoState115Action,
	{_State314, ImplicitEnumDefType}:                    _GotoState118Action,
	{_State314, ExplicitEnumDefType}:                    _GotoState114Action,
	{_State314, TraitDefType}:                           _GotoState123Action,
	{_State314, ReturnTypeType}:                         _GotoState317Action,
	{_State314, FuncTypeType}:                           _GotoState117Action,
	{_State315, StringLiteralToken}:                     _GotoState318Action,
	{_State317, LbraceToken}:                            _GotoState92Action,
	{_State317, BlockBodyType}:                          _GotoState319Action,
	{_State5, _WildcardMarker}:                          _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State16, IdentifierToken}:                         _ReduceNilToOptionalReceiverAction,
	{_State17, _WildcardMarker}:                         _ReduceBitAndToPrefixUnaryOpAction,
	{_State18, _WildcardMarker}:                         _ReduceBitNegToPrefixUnaryOpAction,
	{_State19, _WildcardMarker}:                         _ReduceFalseToLiteralAction,
	{_State20, _WildcardMarker}:                         _ReduceFloatLiteralToLiteralAction,
	{_State22, _WildcardMarker}:                         _ReduceIdentifierToAtomExprAction,
	{_State23, _WildcardMarker}:                         _ReduceIntegerLiteralToLiteralAction,
	{_State24, _WildcardMarker}:                         _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State25, _WildcardMarker}:                         _ReduceLexErrorToAtomExprAction,
	{_State26, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State27, _WildcardMarker}:                         _ReduceMulToPrefixUnaryOpAction,
	{_State28, _WildcardMarker}:                         _ReduceNotToPrefixUnaryOpAction,
	{_State29, _WildcardMarker}:                         _ReduceRuneLiteralToLiteralAction,
	{_State30, _WildcardMarker}:                         _ReduceStringLiteralToLiteralAction,
	{_State32, _WildcardMarker}:                         _ReduceSubToPrefixUnaryOpAction,
	{_State33, _WildcardMarker}:                         _ReduceTrueToLiteralAction,
	{_State34, _WildcardMarker}:                         _ReduceAccessExprToPostfixUnaryExprAction,
	{_State34, LparenToken}:                             _ReduceNilToOptionalGenericBindingAction,
	{_State35, _WildcardMarker}:                         _ReduceAddExprToCmpExprAction,
	{_State36, _WildcardMarker}:                         _ReduceAndExprToOrExprAction,
	{_State37, _WildcardMarker}:                         _ReduceAnonymousFuncExprToAtomExprAction,
	{_State38, _WildcardMarker}:                         _ReduceAnonymousStructExprToAtomExprAction,
	{_State39, _WildcardMarker}:                         _ReduceAtomExprToAccessExprAction,
	{_State40, _WildcardMarker}:                         _ReduceBlockExprToAtomExprAction,
	{_State41, _WildcardMarker}:                         _ReduceCallExprToAccessExprAction,
	{_State42, _WildcardMarker}:                         _ReduceCmpExprToAndExprAction,
	{_State44, _WildcardMarker}:                         _ReduceLiteralToAtomExprAction,
	{_State45, _WildcardMarker}:                         _ReduceMulExprToAddExprAction,
	{_State47, _EndMarker}:                              _ReduceToSequenceExprAction,
	{_State48, _WildcardMarker}:                         _ReducePostfixUnaryExprToPrefixUnaryExprAction,
	{_State49, _WildcardMarker}:                         _ReducePrefixUnaryExprToMulExprAction,
	{_State50, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State51, _EndMarker}:                              _ReduceSequenceExprToExpressionAction,
	{_State52, _EndMarker}:                              _ReduceCommentToLexInternalTokensAction,
	{_State53, _EndMarker}:                              _ReduceSpacesToLexInternalTokensAction,
	{_State54, _EndMarker}:                              _ReduceNoSpecToPackageDefAction,
	{_State55, _WildcardMarker}:                         _ReduceNilToOptionalGenericParametersAction,
	{_State56, RparenToken}:                             _ReduceNilToOptionalTraitPropertiesAction,
	{_State59, RparenToken}:                             _ReduceNilToOptionalParameterDefsAction,
	{_State60, _WildcardMarker}:                         _ReduceArgumentToArgumentsAction,
	{_State62, _WildcardMarker}:                         _ReducePositionalToArgumentAction,
	{_State63, RparenToken}:                             _ReduceNilToOptionalExplicitFieldDefsAction,
	{_State64, RbracketToken}:                           _ReduceNilToOptionalGenericArgumentsAction,
	{_State66, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State67, _WildcardMarker}:                         _ReduceQuestionToPostfixUnaryExprAction,
	{_State69, _WildcardMarker}:                         _ReduceAddToAddOpAction,
	{_State70, _WildcardMarker}:                         _ReduceBitOrToAddOpAction,
	{_State71, _WildcardMarker}:                         _ReduceBitXorToAddOpAction,
	{_State72, _WildcardMarker}:                         _ReduceSubToAddOpAction,
	{_State73, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State74, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State75, _WildcardMarker}:                         _ReduceEqualToCmpOpAction,
	{_State76, _WildcardMarker}:                         _ReduceGreaterToCmpOpAction,
	{_State77, _WildcardMarker}:                         _ReduceGreaterOrEqualToCmpOpAction,
	{_State78, _WildcardMarker}:                         _ReduceLessToCmpOpAction,
	{_State79, _WildcardMarker}:                         _ReduceLessOrEqualToCmpOpAction,
	{_State80, _WildcardMarker}:                         _ReduceNotEqualToCmpOpAction,
	{_State81, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State82, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State83, _WildcardMarker}:                         _ReduceBitAndToMulOpAction,
	{_State84, _WildcardMarker}:                         _ReduceBitLshiftToMulOpAction,
	{_State85, _WildcardMarker}:                         _ReduceBitRshiftToMulOpAction,
	{_State86, _WildcardMarker}:                         _ReduceDivToMulOpAction,
	{_State87, _WildcardMarker}:                         _ReduceModToMulOpAction,
	{_State88, _WildcardMarker}:                         _ReduceMulToMulOpAction,
	{_State89, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State90, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State91, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State92, _WildcardMarker}:                         _ReduceEmptyListToStatementsAction,
	{_State94, _EndMarker}:                              _ReduceToBlockExprAction,
	{_State95, _EndMarker}:                              _ReduceIfExprToExpressionAction,
	{_State96, _EndMarker}:                              _ReduceLoopExprToExpressionAction,
	{_State97, _EndMarker}:                              _ReduceSwitchExprToExpressionAction,
	{_State98, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State100, _WildcardMarker}:                        _ReducePrefixOpToPrefixUnaryExprAction,
	{_State101, _WildcardMarker}:                        _ReduceEmptyListToPackageStatementsAction,
	{_State102, RbracketToken}:                          _ReduceNilToOptionalGenericParameterDefsAction,
	{_State109, _WildcardMarker}:                        _ReduceNilToOptionalGenericBindingAction,
	{_State110, RparenToken}:                            _ReduceNilToOptionalImplicitFieldDefsAction,
	{_State111, _WildcardMarker}:                        _ReduceQuestionToAtomTypeAction,
	{_State113, _WildcardMarker}:                        _ReduceAtomTypeToTraitableTypeAction,
	{_State114, _WildcardMarker}:                        _ReduceExplicitEnumDefToAtomTypeAction,
	{_State115, _WildcardMarker}:                        _ReduceExplicitStructDefToAtomTypeAction,
	{_State116, _WildcardMarker}:                        _ReduceFieldDefToTraitPropertyAction,
	{_State117, _EndMarker}:                             _ReduceFuncTypeToValueTypeAction,
	{_State118, _WildcardMarker}:                        _ReduceImplicitEnumDefToAtomTypeAction,
	{_State119, _WildcardMarker}:                        _ReduceImplicitStructDefToAtomTypeAction,
	{_State120, _WildcardMarker}:                        _ReduceMethodSignatureToTraitPropertyAction,
	{_State122, _EndMarker}:                             _ReduceTraitAddTypeToValueTypeAction,
	{_State123, _WildcardMarker}:                        _ReduceTraitDefToAtomTypeAction,
	{_State124, _WildcardMarker}:                        _ReduceTraitMulTypeToTraitAddTypeAction,
	{_State125, RparenToken}:                            _ReduceTraitPropertiesToOptionalTraitPropertiesAction,
	{_State126, _WildcardMarker}:                        _ReduceTraitPropertyToTraitPropertiesAction,
	{_State127, _WildcardMarker}:                        _ReduceTraitableTypeToTraitMulTypeAction,
	{_State128, _WildcardMarker}:                        _ReduceImplicitToFieldDefAction,
	{_State131, LparenToken}:                            _ReduceNilToOptionalGenericParametersAction,
	{_State133, _WildcardMarker}:                        _ReduceParameterDefToParameterDefsAction,
	{_State134, RparenToken}:                            _ReduceParameterDefsToOptionalParameterDefsAction,
	{_State135, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State136, _WildcardMarker}:                        _ReduceImplicitToAnonymousStructExprAction,
	{_State138, RparenToken}:                            _ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction,
	{_State139, _WildcardMarker}:                        _ReduceFieldDefToExplicitFieldDefsAction,
	{_State141, _WildcardMarker}:                        _ReduceNilToOptionalGenericBindingAction,
	{_State142, RbracketToken}:                          _ReduceGenericArgumentsToOptionalGenericArgumentsAction,
	{_State144, _WildcardMarker}:                        _ReduceValueTypeToGenericArgumentsAction,
	{_State145, _WildcardMarker}:                        _ReduceAccessToAccessExprAction,
	{_State147, RparenToken}:                            _ReduceNilToOptionalArgumentsAction,
	{_State147, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State148, _WildcardMarker}:                        _ReduceOpToAddExprAction,
	{_State149, _WildcardMarker}:                        _ReduceOpToAndExprAction,
	{_State150, _WildcardMarker}:                        _ReduceOpToCmpExprAction,
	{_State152, _WildcardMarker}:                        _ReduceOpToMulExprAction,
	{_State153, _WildcardMarker}:                        _ReduceBlockExprToAtomExprAction,
	{_State153, _EndMarker}:                             _ReduceInfiniteToLoopExprAction,
	{_State154, LbraceToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State156, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State158, _WildcardMarker}:                        _ReduceOpToOrExprAction,
	{_State160, _WildcardMarker}:                        _ReduceUnconstrainedToGenericParameterDefAction,
	{_State161, _WildcardMarker}:                        _ReduceGenericParameterDefToGenericParameterDefsAction,
	{_State162, RbracketToken}:                          _ReduceGenericParameterDefsToOptionalGenericParameterDefsAction,
	{_State164, _EndMarker}:                             _ReduceAliasToTypeDefAction,
	{_State165, _EndMarker}:                             _ReduceDefinitionToTypeDefAction,
	{_State166, _EndMarker}:                             _ReduceReferenceToValueTypeAction,
	{_State167, _WildcardMarker}:                        _ReduceMethodInterfaceToTraitableTypeAction,
	{_State170, RparenToken}:                            _ReduceNilToOptionalParameterDeclsAction,
	{_State171, _WildcardMarker}:                        _ReduceNamedToAtomTypeAction,
	{_State172, _WildcardMarker}:                        _ReduceExplicitToFieldDefAction,
	{_State173, _WildcardMarker}:                        _ReduceFieldDefToImplicitFieldDefsAction,
	{_State175, RparenToken}:                            _ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction,
	{_State177, _WildcardMarker}:                        _ReduceTraitToTraitableTypeAction,
	{_State178, _EndMarker}:                             _ReduceToTraitDefAction,
	{_State185, _WildcardMarker}:                        _ReduceArgToParameterDefAction,
	{_State186, IdentifierToken}:                        _ReduceReceiverToOptionalReceiverAction,
	{_State188, LbraceToken}:                            _ReduceNilToReturnTypeAction,
	{_State190, _WildcardMarker}:                        _ReduceAddToArgumentsAction,
	{_State193, _WildcardMarker}:                        _ReduceToExplicitStructDefAction,
	{_State195, _WildcardMarker}:                        _ReduceBindingToOptionalGenericBindingAction,
	{_State196, _WildcardMarker}:                        _ReduceIndexToAccessExprAction,
	{_State197, RparenToken}:                            _ReduceArgumentsToOptionalArgumentsAction,
	{_State199, _WildcardMarker}:                        _ReduceExplicitToAnonymousStructExprAction,
	{_State200, _EndMarker}:                             _ReduceWhileToLoopExprAction,
	{_State201, _WildcardMarker}:                        _ReduceNoElseToIfExprAction,
	{_State202, LbraceToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State203, _WildcardMarker}:                        _ReduceBreakToJumpTypeAction,
	{_State204, _WildcardMarker}:                        _ReduceContinueToJumpTypeAction,
	{_State205, _EndMarker}:                             _ReduceToBlockBodyAction,
	{_State206, _WildcardMarker}:                        _ReduceReturnToJumpTypeAction,
	{_State208, _WildcardMarker}:                        _ReduceAccessExprToPostfixUnaryExprAction,
	{_State208, LparenToken}:                            _ReduceNilToOptionalGenericBindingAction,
	{_State209, _WildcardMarker}:                        _ReduceExpressionToExpressionOrImplicitStructAction,
	{_State210, _WildcardMarker}:                        _ReduceExpressionOrImplicitStructToStatementBodyAction,
	{_State211, _WildcardMarker}:                        _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State212, _WildcardMarker}:                        _ReduceAddToStatementsAction,
	{_State214, _WildcardMarker}:                        _ReduceUnsafeStatementToStatementBodyAction,
	{_State216, _EndMarker}:                             _ReduceWithSpecToPackageDefAction,
	{_State217, _WildcardMarker}:                        _ReduceAddToPackageStatementsAction,
	{_State219, _WildcardMarker}:                        _ReduceToPackageStatementBodyAction,
	{_State220, _WildcardMarker}:                        _ReduceConstrainedToGenericParameterDefAction,
	{_State222, _WildcardMarker}:                        _ReduceGenericToOptionalGenericParametersAction,
	{_State227, RparenToken}:                            _ReduceNilToOptionalParameterDeclsAction,
	{_State229, _WildcardMarker}:                        _ReduceNilToOptionalGenericBindingAction,
	{_State231, _WildcardMarker}:                        _ReduceParameterDeclToParameterDeclsAction,
	{_State232, RparenToken}:                            _ReduceParameterDeclsToOptionalParameterDeclsAction,
	{_State233, _WildcardMarker}:                        _ReduceUnamedToParameterDeclAction,
	{_State236, _WildcardMarker}:                        _ReduceToImplicitEnumDefAction,
	{_State238, _WildcardMarker}:                        _ReduceToImplicitStructDefAction,
	{_State239, _WildcardMarker}:                        _ReduceUnionToTraitAddTypeAction,
	{_State240, _WildcardMarker}:                        _ReduceDifferenceToTraitAddTypeAction,
	{_State241, _WildcardMarker}:                        _ReduceIntersectToTraitMulTypeAction,
	{_State242, _WildcardMarker}:                        _ReduceExplicitToTraitPropertiesAction,
	{_State243, _WildcardMarker}:                        _ReduceImplicitToTraitPropertiesAction,
	{_State244, _WildcardMarker}:                        _ReduceVarargToParameterDefAction,
	{_State245, RparenToken}:                            _ReduceNilToOptionalParameterDefsAction,
	{_State247, _EndMarker}:                             _ReduceValueTypeToReturnTypeAction,
	{_State248, _WildcardMarker}:                        _ReduceAddToParameterDefsAction,
	{_State249, _WildcardMarker}:                        _ReduceExplicitToExplicitFieldDefsAction,
	{_State250, _WildcardMarker}:                        _ReduceImplicitToExplicitFieldDefsAction,
	{_State251, _WildcardMarker}:                        _ReduceAddToGenericArgumentsAction,
	{_State252, _WildcardMarker}:                        _ReduceToCallExprAction,
	{_State254, LparenToken}:                            _ReduceNilToOptionalGenericBindingAction,
	{_State255, _WildcardMarker}:                        _ReduceCallExprToAccessExprAction,
	{_State255, NewlinesToken}:                          _ReduceAsyncToStatementBodyAction,
	{_State255, SemicolonToken}:                         _ReduceAsyncToStatementBodyAction,
	{_State257, _WildcardMarker}:                        _ReduceAddAssignToBinaryOpAssignAction,
	{_State258, _WildcardMarker}:                        _ReduceAddOneAssignToOpOneAssignAction,
	{_State259, _WildcardMarker}:                        _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State260, _WildcardMarker}:                        _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State261, _WildcardMarker}:                        _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State262, _WildcardMarker}:                        _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State263, _WildcardMarker}:                        _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State264, _WildcardMarker}:                        _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State265, _WildcardMarker}:                        _ReduceDivAssignToBinaryOpAssignAction,
	{_State266, _WildcardMarker}:                        _ReduceModAssignToBinaryOpAssignAction,
	{_State267, _WildcardMarker}:                        _ReduceMulAssignToBinaryOpAssignAction,
	{_State268, _WildcardMarker}:                        _ReduceSubAssignToBinaryOpAssignAction,
	{_State269, _WildcardMarker}:                        _ReduceSubOneAssignToOpOneAssignAction,
	{_State270, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State271, _WildcardMarker}:                        _ReduceOpOneAssignToStatementBodyAction,
	{_State272, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State273, _WildcardMarker}:                        _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State274, NewlinesToken}:                          _ReduceNilToOptionalExpressionOrImplicitStructAction,
	{_State274, SemicolonToken}:                         _ReduceNilToOptionalExpressionOrImplicitStructAction,
	{_State274, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State275, _WildcardMarker}:                        _ReduceImplicitToStatementAction,
	{_State276, _WildcardMarker}:                        _ReduceExplicitToStatementAction,
	{_State278, _WildcardMarker}:                        _ReduceImplicitToPackageStatementAction,
	{_State279, _WildcardMarker}:                        _ReduceExplicitToPackageStatementAction,
	{_State280, _WildcardMarker}:                        _ReduceAddToGenericParameterDefsAction,
	{_State281, _EndMarker}:                             _ReduceConstrainedDefToTypeDefAction,
	{_State282, _WildcardMarker}:                        _ReduceToExplicitEnumDefAction,
	{_State288, _WildcardMarker}:                        _ReduceUnnamedVarargToParameterDeclAction,
	{_State290, _WildcardMarker}:                        _ReduceArgToParameterDeclAction,
	{_State291, _WildcardMarker}:                        _ReduceNilToReturnTypeAction,
	{_State293, _WildcardMarker}:                        _ReducePairToImplicitEnumValueDefsAction,
	{_State294, _WildcardMarker}:                        _ReduceAddToImplicitEnumValueDefsAction,
	{_State295, _WildcardMarker}:                        _ReduceAddToImplicitFieldDefsAction,
	{_State297, _WildcardMarker}:                        _ReduceToAnonymousFuncExprAction,
	{_State298, _EndMarker}:                             _ReduceIfElseToIfExprAction,
	{_State299, _EndMarker}:                             _ReduceMultiIfElseToIfExprAction,
	{_State301, _WildcardMarker}:                        _ReduceBinaryOpAssignToStatementBodyAction,
	{_State302, _WildcardMarker}:                        _ReduceImplicitStructToExpressionOrImplicitStructAction,
	{_State303, _WildcardMarker}:                        _ReduceExpressionOrImplicitStructToOptionalExpressionOrImplicitStructAction,
	{_State304, _WildcardMarker}:                        _ReduceJumpToStatementBodyAction,
	{_State305, _EndMarker}:                             _ReduceToSwitchExprAction,
	{_State306, RparenToken}:                            _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State307, _WildcardMarker}:                        _ReducePairToImplicitEnumValueDefsAction,
	{_State307, RparenToken}:                            _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State308, RparenToken}:                            _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State309, _WildcardMarker}:                        _ReduceAddToImplicitEnumValueDefsAction,
	{_State309, RparenToken}:                            _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State310, _WildcardMarker}:                        _ReduceNilToReturnTypeAction,
	{_State311, _WildcardMarker}:                        _ReduceVarargToParameterDeclAction,
	{_State312, _EndMarker}:                             _ReduceToFuncTypeAction,
	{_State313, _WildcardMarker}:                        _ReduceAddToParameterDeclsAction,
	{_State314, LbraceToken}:                            _ReduceNilToReturnTypeAction,
	{_State316, _WildcardMarker}:                        _ReduceToMethodSignatureAction,
	{_State318, _WildcardMarker}:                        _ReduceToUnsafeStatementAction,
	{_State319, _EndMarker}:                             _ReduceToNamedFuncDefAction,
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.package_def
    Reduce:
      (nil)
    Goto:
      PACKAGE -> State 13
      package_def -> State 7

  State 2:
    Kernel Items:
      #accept: ^.type_def
    Reduce:
      (nil)
    Goto:
      TYPE -> State 14
      type_def -> State 8

  State 3:
    Kernel Items:
      #accept: ^.trait_def
    Reduce:
      (nil)
    Goto:
      TRAIT -> State 15
      trait_def -> State 9

  State 4:
    Kernel Items:
      #accept: ^.named_func_def
    Reduce:
      (nil)
    Goto:
      FUNC -> State 16
      named_func_def -> State 10

  State 5:
    Kernel Items:
      #accept: ^.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      sequence_expr -> State 51
      optional_label_decl -> State 46
      block_expr -> State 40
      expression -> State 11
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 6:
    Kernel Items:
      #accept: ^.lex_internal_tokens
    Reduce:
      (nil)
    Goto:
      SPACES -> State 53
      COMMENT -> State 52
      lex_internal_tokens -> State 12

  State 7:
    Kernel Items:
      #accept: ^ package_def., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      #accept: ^ type_def., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      #accept: ^ trait_def., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      #accept: ^ named_func_def., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 11:
    Kernel Items:
      #accept: ^ expression., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 12:
    Kernel Items:
      #accept: ^ lex_internal_tokens., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      package_def: PACKAGE.IDENTIFIER
      package_def: PACKAGE.IDENTIFIER LPAREN package_statements RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 54

  State 14:
    Kernel Items:
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE.IDENTIFIER EQUAL value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 55

  State 15:
    Kernel Items:
      trait_def: TRAIT.LPAREN optional_trait_properties RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 56

  State 16:
    Kernel Items:
      named_func_def: FUNC.optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      LPAREN -> State 57
      optional_receiver -> State 58

  State 17:
    Kernel Items:
      prefix_unary_op: BIT_AND., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 18:
    Kernel Items:
      prefix_unary_op: BIT_NEG., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 19:
    Kernel Items:
      literal: FALSE., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 20:
    Kernel Items:
      literal: FLOAT_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      anonymous_func_expr: FUNC.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 59

  State 22:
    Kernel Items:
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      literal: INTEGER_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 24:
    Kernel Items:
      optional_label_decl: LABEL_DECL., *
    Reduce:
      * -> [optional_label_decl]
    Goto:
      (nil)

  State 25:
    Kernel Items:
      atom_expr: LEX_ERROR., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      anonymous_struct_expr: LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      argument -> State 60
      arguments -> State 61
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      sequence_expr -> State 51
      optional_label_decl -> State 46
      block_expr -> State 40
      expression -> State 62
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 27:
    Kernel Items:
      prefix_unary_op: MUL., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 28:
    Kernel Items:
      prefix_unary_op: NOT., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 29:
    Kernel Items:
      literal: RUNE_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 30:
    Kernel Items:
      literal: STRING_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 31:
    Kernel Items:
      explicit_struct_def: STRUCT.LPAREN optional_explicit_field_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 63

  State 32:
    Kernel Items:
      prefix_unary_op: SUB., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 33:
    Kernel Items:
      literal: TRUE., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 34:
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
      LBRACKET -> State 66
      DOT -> State 65
      QUESTION -> State 67
      DOLLAR_LBRACKET -> State 64
      optional_generic_binding -> State 68

  State 35:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 69
      SUB -> State 72
      BIT_XOR -> State 71
      BIT_OR -> State 70
      add_op -> State 73

  State 36:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 74

  State 37:
    Kernel Items:
      atom_expr: anonymous_func_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 38:
    Kernel Items:
      atom_expr: anonymous_struct_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 39:
    Kernel Items:
      access_expr: atom_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 40:
    Kernel Items:
      atom_expr: block_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 41:
    Kernel Items:
      access_expr: call_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 42:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 75
      NOT_EQUAL -> State 80
      LESS -> State 78
      LESS_OR_EQUAL -> State 79
      GREATER -> State 76
      GREATER_OR_EQUAL -> State 77
      cmp_op -> State 81

  State 43:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_def.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 82

  State 44:
    Kernel Items:
      atom_expr: literal., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 45:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 88
      DIV -> State 86
      MOD -> State 87
      BIT_AND -> State 83
      BIT_LSHIFT -> State 84
      BIT_RSHIFT -> State 85
      mul_op -> State 89

  State 46:
    Kernel Items:
      block_expr: optional_label_decl.block_body
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 91
      SWITCH -> State 93
      FOR -> State 90
      LBRACE -> State 92
      block_body -> State 94
      if_expr -> State 95
      switch_expr -> State 97
      loop_expr -> State 96

  State 47:
    Kernel Items:
      or_expr: or_expr.OR and_expr
      sequence_expr: or_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      OR -> State 98

  State 48:
    Kernel Items:
      prefix_unary_expr: postfix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 49:
    Kernel Items:
      mul_expr: prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 50:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op.prefix_unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 100
      optional_label_decl -> State 99
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 51:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 52:
    Kernel Items:
      lex_internal_tokens: COMMENT., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 53:
    Kernel Items:
      lex_internal_tokens: SPACES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 54:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER., $
      package_def: PACKAGE IDENTIFIER.LPAREN package_statements RPAREN
    Reduce:
      $ -> [package_def]
    Goto:
      LPAREN -> State 101

  State 55:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE IDENTIFIER.EQUAL value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 102
      EQUAL -> State 103
      optional_generic_parameters -> State 104

  State 56:
    Kernel Items:
      trait_def: TRAIT LPAREN.optional_trait_properties RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      IDENTIFIER -> State 109
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 108
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 128
      field_def -> State 116
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_property -> State 126
      trait_properties -> State 125
      optional_trait_properties -> State 121
      trait_def -> State 123
      func_type -> State 117
      method_signature -> State 120

  State 57:
    Kernel Items:
      optional_receiver: LPAREN.parameter_def RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 129
      parameter_def -> State 130

  State 58:
    Kernel Items:
      named_func_def: FUNC optional_receiver.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131

  State 59:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 129
      parameter_def -> State 133
      parameter_defs -> State 134
      optional_parameter_defs -> State 132

  State 60:
    Kernel Items:
      arguments: argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 61:
    Kernel Items:
      anonymous_struct_expr: LPAREN arguments.RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 136
      COMMA -> State 135

  State 62:
    Kernel Items:
      argument: expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 63:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN.optional_explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      IDENTIFIER -> State 109
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 128
      field_def -> State 139
      implicit_struct_def -> State 119
      explicit_field_defs -> State 138
      optional_explicit_field_defs -> State 140
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 64:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      generic_arguments -> State 142
      optional_generic_arguments -> State 143
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 144
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 65:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 145

  State 66:
    Kernel Items:
      access_expr: access_expr LBRACKET.expression RBRACKET
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      sequence_expr -> State 51
      optional_label_decl -> State 46
      block_expr -> State 40
      expression -> State 146
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 67:
    Kernel Items:
      postfix_unary_expr: access_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 68:
    Kernel Items:
      call_expr: access_expr optional_generic_binding.LPAREN optional_arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 147

  State 69:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 70:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 71:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 72:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 73:
    Kernel Items:
      add_expr: add_expr add_op.mul_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 148
      optional_label_decl -> State 99
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 74:
    Kernel Items:
      and_expr: and_expr AND.cmp_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 149
      optional_label_decl -> State 99
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 75:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 76:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 77:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 78:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 79:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 80:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 81:
    Kernel Items:
      cmp_expr: cmp_expr cmp_op.add_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 150
      optional_label_decl -> State 99
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 82:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_def LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      argument -> State 60
      arguments -> State 151
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      sequence_expr -> State 51
      optional_label_decl -> State 46
      block_expr -> State 40
      expression -> State 62
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 83:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 84:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 85:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 86:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 87:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      mul_expr: mul_expr mul_op.prefix_unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 152
      optional_label_decl -> State 99
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 90:
    Kernel Items:
      loop_expr: FOR.block_expr
      loop_expr: FOR.sequence_expr block_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      sequence_expr -> State 154
      optional_label_decl -> State 99
      block_expr -> State 153
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 91:
    Kernel Items:
      if_expr: IF.sequence_expr block_body
      if_expr: IF.sequence_expr block_body ELSE block_body
      if_expr: IF.sequence_expr block_body ELSE if_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      sequence_expr -> State 155
      optional_label_decl -> State 99
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 92:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 156

  State 93:
    Kernel Items:
      switch_expr: SWITCH.LBRACE CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 157

  State 94:
    Kernel Items:
      block_expr: optional_label_decl block_body., $
    Reduce:
      $ -> [block_expr]
    Goto:
      (nil)

  State 95:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 96:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 97:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 98:
    Kernel Items:
      or_expr: or_expr OR.and_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 158
      optional_label_decl -> State 99
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 99:
    Kernel Items:
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 92
      block_body -> State 94

  State 100:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 101:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN.package_statements RPAREN
    Reduce:
      * -> [package_statements]
    Goto:
      package_statements -> State 159

  State 102:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.optional_generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 160
      generic_parameter_def -> State 161
      generic_parameter_defs -> State 162
      optional_generic_parameter_defs -> State 163

  State 103:
    Kernel Items:
      type_def: TYPE IDENTIFIER EQUAL.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 164
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 104:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 165
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 105:
    Kernel Items:
      value_type: BIT_AND.trait_add_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 166
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123

  State 106:
    Kernel Items:
      traitable_type: BIT_NEG.atom_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      LPAREN -> State 110
      QUESTION -> State 111
      atom_type -> State 167
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123

  State 107:
    Kernel Items:
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 168

  State 108:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 169
      LPAREN -> State 170

  State 109:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      DOLLAR_LBRACKET -> State 64
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      optional_generic_binding -> State 171
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 172
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 110:
    Kernel Items:
      implicit_struct_def: LPAREN.optional_implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      IDENTIFIER -> State 109
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 128
      field_def -> State 173
      implicit_field_defs -> State 175
      optional_implicit_field_defs -> State 176
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_value_defs -> State 174
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 111:
    Kernel Items:
      atom_type: QUESTION., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 112:
    Kernel Items:
      traitable_type: TILDE_TILDE.atom_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      LPAREN -> State 110
      QUESTION -> State 111
      atom_type -> State 177
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123

  State 113:
    Kernel Items:
      traitable_type: atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 114:
    Kernel Items:
      atom_type: explicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      atom_type: explicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 117:
    Kernel Items:
      value_type: func_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      (nil)

  State 118:
    Kernel Items:
      atom_type: implicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 119:
    Kernel Items:
      atom_type: implicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 121:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 178

  State 122:
    Kernel Items:
      trait_add_type: trait_add_type.ADD trait_mul_type
      trait_add_type: trait_add_type.SUB trait_mul_type
      value_type: trait_add_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      ADD -> State 179
      SUB -> State 180

  State 123:
    Kernel Items:
      atom_type: trait_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      trait_mul_type: trait_mul_type.MUL traitable_type
      trait_add_type: trait_mul_type., *
    Reduce:
      * -> [trait_add_type]
    Goto:
      MUL -> State 181

  State 125:
    Kernel Items:
      trait_properties: trait_properties.NEWLINES trait_property
      trait_properties: trait_properties.COMMA trait_property
      optional_trait_properties: trait_properties., RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      NEWLINES -> State 183
      COMMA -> State 182

  State 126:
    Kernel Items:
      trait_properties: trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 127:
    Kernel Items:
      trait_mul_type: traitable_type., *
    Reduce:
      * -> [trait_mul_type]
    Goto:
      (nil)

  State 128:
    Kernel Items:
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 129:
    Kernel Items:
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      DOTDOTDOT -> State 184
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 185
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 130:
    Kernel Items:
      optional_receiver: LPAREN parameter_def.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 186

  State 131:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 102
      optional_generic_parameters -> State 187

  State 132:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 188

  State 133:
    Kernel Items:
      parameter_defs: parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 134:
    Kernel Items:
      parameter_defs: parameter_defs.COMMA parameter_def
      optional_parameter_defs: parameter_defs., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      COMMA -> State 189

  State 135:
    Kernel Items:
      arguments: arguments COMMA.argument
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      argument -> State 190
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      sequence_expr -> State 51
      optional_label_decl -> State 46
      block_expr -> State 40
      expression -> State 62
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 136:
    Kernel Items:
      anonymous_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 137:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 170

  State 138:
    Kernel Items:
      explicit_field_defs: explicit_field_defs.NEWLINES field_def
      explicit_field_defs: explicit_field_defs.COMMA field_def
      optional_explicit_field_defs: explicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      NEWLINES -> State 192
      COMMA -> State 191

  State 139:
    Kernel Items:
      explicit_field_defs: field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 140:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 193

  State 141:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 64
      optional_generic_binding -> State 171

  State 142:
    Kernel Items:
      generic_arguments: generic_arguments.COMMA value_type
      optional_generic_arguments: generic_arguments., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 194

  State 143:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 195

  State 144:
    Kernel Items:
      generic_arguments: value_type., *
    Reduce:
      * -> [generic_arguments]
    Goto:
      (nil)

  State 145:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 146:
    Kernel Items:
      access_expr: access_expr LBRACKET expression.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 196

  State 147:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN.optional_arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_arguments]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      argument -> State 60
      arguments -> State 197
      optional_arguments -> State 198
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      sequence_expr -> State 51
      optional_label_decl -> State 46
      block_expr -> State 40
      expression -> State 62
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 148:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 88
      DIV -> State 86
      MOD -> State 87
      BIT_AND -> State 83
      BIT_LSHIFT -> State 84
      BIT_RSHIFT -> State 85
      mul_op -> State 89

  State 149:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 75
      NOT_EQUAL -> State 80
      LESS -> State 78
      LESS_OR_EQUAL -> State 79
      GREATER -> State 76
      GREATER_OR_EQUAL -> State 77
      cmp_op -> State 81

  State 150:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 69
      SUB -> State 72
      BIT_XOR -> State 71
      BIT_OR -> State 70
      add_op -> State 73

  State 151:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_def LPAREN arguments.RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 199
      COMMA -> State 135

  State 152:
    Kernel Items:
      mul_expr: mul_expr mul_op prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 153:
    Kernel Items:
      atom_expr: block_expr., *
      loop_expr: FOR block_expr., $
    Reduce:
      * -> [atom_expr]
      $ -> [loop_expr]
    Goto:
      (nil)

  State 154:
    Kernel Items:
      loop_expr: FOR sequence_expr.block_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      LABEL_DECL -> State 24
      optional_label_decl -> State 99
      block_expr -> State 200

  State 155:
    Kernel Items:
      if_expr: IF sequence_expr.block_body
      if_expr: IF sequence_expr.block_body ELSE block_body
      if_expr: IF sequence_expr.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 92
      block_body -> State 201

  State 156:
    Kernel Items:
      statements: statements.statement
      block_body: LBRACE statements.RBRACE
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      RETURN -> State 206
      BREAK -> State 203
      CONTINUE -> State 204
      UNSAFE -> State 207
      STRUCT -> State 31
      FUNC -> State 21
      ASYNC -> State 202
      RBRACE -> State 205
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      call_expr -> State 41
      access_expr -> State 208
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      sequence_expr -> State 51
      jump_type -> State 211
      expression_or_implicit_struct -> State 210
      unsafe_statement -> State 214
      statement_body -> State 213
      statement -> State 212
      optional_label_decl -> State 46
      block_expr -> State 40
      expression -> State 209
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 157:
    Kernel Items:
      switch_expr: SWITCH LBRACE.CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      CASE -> State 215

  State 158:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 74

  State 159:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements.RPAREN
      package_statements: package_statements.package_statement
    Reduce:
      (nil)
    Goto:
      UNSAFE -> State 207
      RPAREN -> State 216
      unsafe_statement -> State 219
      package_statement_body -> State 218
      package_statement -> State 217

  State 160:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 220
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 161:
    Kernel Items:
      generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 162:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs.COMMA generic_parameter_def
      optional_generic_parameter_defs: generic_parameter_defs., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      COMMA -> State 221

  State 163:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 222

  State 164:
    Kernel Items:
      type_def: TYPE IDENTIFIER EQUAL value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      (nil)

  State 165:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., $
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      $ -> [type_def]
    Goto:
      IMPLEMENTS -> State 223

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
      traitable_type: BIT_NEG atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 168:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 109
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 128
      field_def -> State 225
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_value_defs -> State 226
      implicit_enum_def -> State 118
      explicit_enum_value_defs -> State 224
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 169:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 227

  State 170:
    Kernel Items:
      func_type: FUNC LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 229
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      DOTDOTDOT -> State 228
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 233
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      parameter_decl -> State 231
      parameter_decls -> State 232
      optional_parameter_decls -> State 230
      func_type -> State 117

  State 171:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 172:
    Kernel Items:
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 173:
    Kernel Items:
      implicit_field_defs: field_def., *
      implicit_enum_value_defs: field_def.OR field_def
    Reduce:
      * -> [implicit_field_defs]
    Goto:
      OR -> State 234

  State 174:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR field_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 236
      OR -> State 235

  State 175:
    Kernel Items:
      implicit_field_defs: implicit_field_defs.COMMA field_def
      optional_implicit_field_defs: implicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      COMMA -> State 237

  State 176:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 238

  State 177:
    Kernel Items:
      traitable_type: TILDE_TILDE atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 178:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties RPAREN., $
    Reduce:
      $ -> [trait_def]
    Goto:
      (nil)

  State 179:
    Kernel Items:
      trait_add_type: trait_add_type ADD.trait_mul_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 239
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123

  State 180:
    Kernel Items:
      trait_add_type: trait_add_type SUB.trait_mul_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 240
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123

  State 181:
    Kernel Items:
      trait_mul_type: trait_mul_type MUL.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      atom_type -> State 113
      traitable_type -> State 241
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123

  State 182:
    Kernel Items:
      trait_properties: trait_properties COMMA.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 109
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 108
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 128
      field_def -> State 116
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_property -> State 242
      trait_def -> State 123
      func_type -> State 117
      method_signature -> State 120

  State 183:
    Kernel Items:
      trait_properties: trait_properties NEWLINES.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 109
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 108
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 128
      field_def -> State 116
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_property -> State 243
      trait_def -> State 123
      func_type -> State 117
      method_signature -> State 120

  State 184:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 244
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 185:
    Kernel Items:
      parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      (nil)

  State 186:
    Kernel Items:
      optional_receiver: LPAREN parameter_def RPAREN., IDENTIFIER
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      (nil)

  State 187:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 245

  State 188:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 247
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      return_type -> State 246
      func_type -> State 117

  State 189:
    Kernel Items:
      parameter_defs: parameter_defs COMMA.parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 129
      parameter_def -> State 248

  State 190:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 191:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 109
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 128
      field_def -> State 249
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 192:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 109
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 128
      field_def -> State 250
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 193:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_def]
    Goto:
      (nil)

  State 194:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 251
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 195:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 196:
    Kernel Items:
      access_expr: access_expr LBRACKET expression RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 197:
    Kernel Items:
      arguments: arguments.COMMA argument
      optional_arguments: arguments., RPAREN
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 135

  State 198:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 252

  State 199:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_def LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 200:
    Kernel Items:
      loop_expr: FOR sequence_expr block_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 201:
    Kernel Items:
      if_expr: IF sequence_expr block_body., *
      if_expr: IF sequence_expr block_body.ELSE block_body
      if_expr: IF sequence_expr block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 253

  State 202:
    Kernel Items:
      statement_body: ASYNC.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      call_expr -> State 255
      access_expr -> State 254
      optional_label_decl -> State 99
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 203:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 204:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 205:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 206:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 207:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 256

  State 208:
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
      LBRACKET -> State 66
      DOT -> State 65
      QUESTION -> State 67
      DOLLAR_LBRACKET -> State 64
      ADD_ASSIGN -> State 257
      SUB_ASSIGN -> State 268
      MUL_ASSIGN -> State 267
      DIV_ASSIGN -> State 265
      MOD_ASSIGN -> State 266
      ADD_ONE_ASSIGN -> State 258
      SUB_ONE_ASSIGN -> State 269
      BIT_NEG_ASSIGN -> State 261
      BIT_AND_ASSIGN -> State 259
      BIT_OR_ASSIGN -> State 262
      BIT_XOR_ASSIGN -> State 264
      BIT_LSHIFT_ASSIGN -> State 260
      BIT_RSHIFT_ASSIGN -> State 263
      optional_generic_binding -> State 68
      op_one_assign -> State 271
      binary_op_assign -> State 270

  State 209:
    Kernel Items:
      expression_or_implicit_struct: expression., *
    Reduce:
      * -> [expression_or_implicit_struct]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct.COMMA expression
      statement_body: expression_or_implicit_struct., *
    Reduce:
      * -> [statement_body]
    Goto:
      COMMA -> State 272

  State 211:
    Kernel Items:
      statement_body: jump_type.optional_jump_label optional_expression_or_implicit_struct
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 273
      optional_jump_label -> State 274

  State 212:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 275
      SEMICOLON -> State 276

  State 214:
    Kernel Items:
      statement_body: unsafe_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 215:
    Kernel Items:
      switch_expr: SWITCH LBRACE CASE.DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 277

  State 216:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements RPAREN., $
    Reduce:
      $ -> [package_def]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      package_statements: package_statements package_statement., *
    Reduce:
      * -> [package_statements]
    Goto:
      (nil)

  State 218:
    Kernel Items:
      package_statement: package_statement_body.NEWLINES
      package_statement: package_statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 278
      SEMICOLON -> State 279

  State 219:
    Kernel Items:
      package_statement_body: unsafe_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 220:
    Kernel Items:
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      (nil)

  State 221:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA.generic_parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 160
      generic_parameter_def -> State 280

  State 222:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 223:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 281
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 224:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 282

  State 225:
    Kernel Items:
      implicit_enum_value_defs: field_def.OR field_def
      explicit_enum_value_defs: field_def.OR field_def
      explicit_enum_value_defs: field_def.NEWLINES field_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 283
      OR -> State 284

  State 226:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR field_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR field_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES field_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 285
      OR -> State 286

  State 227:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 229
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      DOTDOTDOT -> State 228
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 233
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      parameter_decl -> State 231
      parameter_decls -> State 232
      optional_parameter_decls -> State 287
      func_type -> State 117

  State 228:
    Kernel Items:
      parameter_decl: DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 288
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 229:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      DOLLAR_LBRACKET -> State 64
      DOTDOTDOT -> State 289
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      optional_generic_binding -> State 171
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 290
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 230:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 291

  State 231:
    Kernel Items:
      parameter_decls: parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 232:
    Kernel Items:
      parameter_decls: parameter_decls.COMMA parameter_decl
      optional_parameter_decls: parameter_decls., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      COMMA -> State 292

  State 233:
    Kernel Items:
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 234:
    Kernel Items:
      implicit_enum_value_defs: field_def OR.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 109
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 128
      field_def -> State 293
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 235:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 109
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 128
      field_def -> State 294
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 236:
    Kernel Items:
      implicit_enum_def: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_def]
    Goto:
      (nil)

  State 237:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 109
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 128
      field_def -> State 295
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 238:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_def]
    Goto:
      (nil)

  State 239:
    Kernel Items:
      trait_mul_type: trait_mul_type.MUL traitable_type
      trait_add_type: trait_add_type ADD trait_mul_type., *
    Reduce:
      * -> [trait_add_type]
    Goto:
      MUL -> State 181

  State 240:
    Kernel Items:
      trait_mul_type: trait_mul_type.MUL traitable_type
      trait_add_type: trait_add_type SUB trait_mul_type., *
    Reduce:
      * -> [trait_add_type]
    Goto:
      MUL -> State 181

  State 241:
    Kernel Items:
      trait_mul_type: trait_mul_type MUL traitable_type., *
    Reduce:
      * -> [trait_mul_type]
    Goto:
      (nil)

  State 242:
    Kernel Items:
      trait_properties: trait_properties COMMA trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 243:
    Kernel Items:
      trait_properties: trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 244:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      (nil)

  State 245:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 129
      parameter_def -> State 133
      parameter_defs -> State 134
      optional_parameter_defs -> State 296

  State 246:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 92
      block_body -> State 297

  State 247:
    Kernel Items:
      return_type: value_type., $
    Reduce:
      $ -> [return_type]
    Goto:
      (nil)

  State 248:
    Kernel Items:
      parameter_defs: parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 249:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 250:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 251:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
    Reduce:
      * -> [generic_arguments]
    Goto:
      (nil)

  State 252:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 253:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE.block_body
      if_expr: IF sequence_expr block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 91
      LBRACE -> State 92
      block_body -> State 298
      if_expr -> State 299

  State 254:
    Kernel Items:
      call_expr: access_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET expression RBRACKET
    Reduce:
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 66
      DOT -> State 65
      DOLLAR_LBRACKET -> State 64
      optional_generic_binding -> State 68

  State 255:
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

  State 256:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 300

  State 257:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 258:
    Kernel Items:
      op_one_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 259:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 260:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 261:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 262:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 263:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 264:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 265:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 266:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 267:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 268:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 269:
    Kernel Items:
      op_one_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 270:
    Kernel Items:
      statement_body: access_expr binary_op_assign.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      sequence_expr -> State 51
      optional_label_decl -> State 46
      block_expr -> State 40
      expression -> State 301
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 271:
    Kernel Items:
      statement_body: access_expr op_one_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 272:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct COMMA.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      sequence_expr -> State 51
      optional_label_decl -> State 46
      block_expr -> State 40
      expression -> State 302
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 273:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 274:
    Kernel Items:
      statement_body: jump_type optional_jump_label.optional_expression_or_implicit_struct
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_expression_or_implicit_struct]
      SEMICOLON -> [optional_expression_or_implicit_struct]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 22
      TRUE -> State 33
      FALSE -> State 19
      STRUCT -> State 31
      FUNC -> State 21
      LPAREN -> State 26
      LABEL_DECL -> State 24
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      literal -> State 44
      anonymous_struct_expr -> State 38
      atom_expr -> State 39
      call_expr -> State 41
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      sequence_expr -> State 51
      optional_expression_or_implicit_struct -> State 304
      expression_or_implicit_struct -> State 303
      optional_label_decl -> State 46
      block_expr -> State 40
      expression -> State 209
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 275:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 276:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 277:
    Kernel Items:
      switch_expr: SWITCH LBRACE CASE DEFAULT.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 305

  State 278:
    Kernel Items:
      package_statement: package_statement_body NEWLINES., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      package_statement: package_statement_body SEMICOLON., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 280:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 281:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      (nil)

  State 282:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_def]
    Goto:
      (nil)

  State 283:
    Kernel Items:
      explicit_enum_value_defs: field_def NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 109
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 128
      field_def -> State 306
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 284:
    Kernel Items:
      implicit_enum_value_defs: field_def OR.field_def
      explicit_enum_value_defs: field_def OR.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 109
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 128
      field_def -> State 307
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 285:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 109
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 128
      field_def -> State 308
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 286:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.field_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 109
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 128
      field_def -> State 309
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 287:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 310

  State 288:
    Kernel Items:
      parameter_decl: DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 289:
    Kernel Items:
      parameter_decl: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 311
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      func_type -> State 117

  State 290:
    Kernel Items:
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 291:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 247
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      return_type -> State 312
      func_type -> State 117

  State 292:
    Kernel Items:
      parameter_decls: parameter_decls COMMA.parameter_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 229
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      DOTDOTDOT -> State 228
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 233
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      parameter_decl -> State 313
      func_type -> State 117

  State 293:
    Kernel Items:
      implicit_enum_value_defs: field_def OR field_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 294:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR field_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 295:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [implicit_field_defs]
    Goto:
      (nil)

  State 296:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 314

  State 297:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 298:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 299:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 300:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 315

  State 301:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 302:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct COMMA expression., *
    Reduce:
      * -> [expression_or_implicit_struct]
    Goto:
      (nil)

  State 303:
    Kernel Items:
      optional_expression_or_implicit_struct: expression_or_implicit_struct., *
      expression_or_implicit_struct: expression_or_implicit_struct.COMMA expression
    Reduce:
      * -> [optional_expression_or_implicit_struct]
    Goto:
      COMMA -> State 272

  State 304:
    Kernel Items:
      statement_body: jump_type optional_jump_label optional_expression_or_implicit_struct., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 305:
    Kernel Items:
      switch_expr: SWITCH LBRACE CASE DEFAULT RBRACE., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 306:
    Kernel Items:
      explicit_enum_value_defs: field_def NEWLINES field_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 307:
    Kernel Items:
      implicit_enum_value_defs: field_def OR field_def., *
      explicit_enum_value_defs: field_def OR field_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 308:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES field_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 309:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR field_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR field_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 310:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 247
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      return_type -> State 316
      func_type -> State 117

  State 311:
    Kernel Items:
      parameter_decl: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 312:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN return_type., $
    Reduce:
      $ -> [func_type]
    Goto:
      (nil)

  State 313:
    Kernel Items:
      parameter_decls: parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 314:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 141
      STRUCT -> State 31
      ENUM -> State 107
      TRAIT -> State 15
      FUNC -> State 137
      LPAREN -> State 110
      QUESTION -> State 111
      TILDE_TILDE -> State 112
      BIT_NEG -> State 106
      BIT_AND -> State 105
      atom_type -> State 113
      traitable_type -> State 127
      trait_mul_type -> State 124
      trait_add_type -> State 122
      value_type -> State 247
      implicit_struct_def -> State 119
      explicit_struct_def -> State 115
      implicit_enum_def -> State 118
      explicit_enum_def -> State 114
      trait_def -> State 123
      return_type -> State 317
      func_type -> State 117

  State 315:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 318

  State 316:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 317:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 92
      block_body -> State 319

  State 318:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 319:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

Number of states: 319
Number of shift actions: 1723
Number of reduce actions: 245
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
