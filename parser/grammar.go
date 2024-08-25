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
	DollarLbraceToken    = SymbolId(293)
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
	// 95:2: literal -> TRUE: ...
	TrueToLiteral(True_ *GenericSymbol) (*GenericSymbol, error)

	// 96:2: literal -> FALSE: ...
	FalseToLiteral(False_ *GenericSymbol) (*GenericSymbol, error)

	// 97:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 98:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 99:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 100:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 102:23: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(ExplicitFuncType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 107:2: anonymous_struct_expr -> explicit: ...
	ExplicitToAnonymousStructExpr(ExplicitStructType_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 108:2: anonymous_struct_expr -> implicit: ...
	ImplicitToAnonymousStructExpr(Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 114:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 115:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 116:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 118:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 119:2: atom_expr -> anonymous_struct_expr: ...
	AnonymousStructExprToAtomExpr(AnonymousStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 120:2: atom_expr -> LEX_ERROR: ...
	LexErrorToAtomExpr(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 123:0: generic_arguments -> ...
	ToGenericArguments() (*GenericSymbol, error)

	// 126:0: arguments -> ...
	ToArguments() (*GenericSymbol, error)

	// 129:2: call_expr -> concrete: ...
	ConcreteToCallExpr(AccessExpr_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 130:2: call_expr -> generic: ...
	GenericToCallExpr(AccessExpr_ *GenericSymbol, DollarLbrace_ *GenericSymbol, GenericArguments_ *GenericSymbol, Rbrace_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 133:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 134:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 135:2: access_expr -> call_expr: ...
	CallExprToAccessExpr(CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 137:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Expression_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 140:2: unary_op -> NOT: ...
	NotToUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 141:2: unary_op -> BIT_NEG: ...
	BitNegToUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 142:2: unary_op -> SUB: ...
	SubToUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 145:2: unary_op -> MUL: ...
	MulToUnaryOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 148:2: unary_op -> BIT_AND: ...
	BitAndToUnaryOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 151:2: unary_expr -> access_expr: ...
	AccessExprToUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 152:2: unary_expr -> op: ...
	OpToUnaryExpr(UnaryOp_ *GenericSymbol, UnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 155:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 156:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 157:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 158:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 159:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 160:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 163:2: mul_expr -> unary_expr: ...
	UnaryExprToMulExpr(UnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 164:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, UnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 167:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 168:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 169:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 170:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 173:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 174:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 177:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 178:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 179:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 180:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 181:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 182:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 185:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 186:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 189:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 190:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 193:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 194:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 196:17: sequence_expr -> ...
	ToSequenceExpr(OrExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 199:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 200:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 203:2: optional_expression_or_implicit_struct -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct(ExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 204:2: optional_expression_or_implicit_struct -> nil: ...
	NilToOptionalExpressionOrImplicitStruct() (*GenericSymbol, error)

	// 207:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 208:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 209:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 212:2: op_one_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToOpOneAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 213:2: op_one_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToOpOneAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 216:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 217:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 218:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 219:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 220:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 221:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 222:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 223:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 224:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 225:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 226:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 229:2: expression_or_implicit_struct -> expression: ...
	ExpressionToExpressionOrImplicitStruct(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 230:2: expression_or_implicit_struct -> implicit_struct: ...
	ImplicitStructToExpressionOrImplicitStruct(ExpressionOrImplicitStruct_ *GenericSymbol, Comma_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 234:20: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ *GenericSymbol, Less_ *GenericSymbol, Identifier_ *GenericSymbol, Greater_ *GenericSymbol, StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 237:2: statement_body -> unsafe_statement: ...
	UnsafeStatementToStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 239:2: statement_body -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToStatementBody(ExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 241:2: statement_body -> async: ...
	AsyncToStatementBody(Async_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 244:2: statement_body -> jump: ...
	JumpToStatementBody(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 257:2: statement_body -> op_one_assign: ...
	OpOneAssignToStatementBody(AccessExpr_ *GenericSymbol, OpOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 258:2: statement_body -> binary_op_assign: ...
	BinaryOpAssignToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 261:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 262:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 265:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 266:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 269:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 270:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (*GenericSymbol, error)

	// 272:14: block_body -> ...
	ToBlockBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 273:14: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 278:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 279:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, BlockBody_2 *GenericSymbol) (*GenericSymbol, error)

	// 280:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 283:15: switch_expr -> ...
	ToSwitchExpr(Switch_ *GenericSymbol, Case_ *GenericSymbol, Default_ *GenericSymbol) (*GenericSymbol, error)

	// 286:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(For_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 287:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 293:2: expression -> sequence_expr: ...
	SequenceExprToExpression(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 294:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 295:2: expression -> switch_expr: ...
	SwitchExprToExpression(OptionalLabelDecl_ *GenericSymbol, SwitchExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 296:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ *GenericSymbol, LoopExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 299:0: generic_parameters -> ...
	ToGenericParameters() (*GenericSymbol, error)

	// 302:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbrace_ *GenericSymbol, GenericParameters_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 303:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (*GenericSymbol, error)

	// 306:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol) (*GenericSymbol, error)

	// 307:2: atom_type -> explicit_struct_type: ...
	ExplicitStructTypeToAtomType(ExplicitStructType_ *GenericSymbol) (*GenericSymbol, error)

	// 308:2: atom_type -> implicit_struct_type: ...
	ImplicitStructTypeToAtomType(ImplicitStructType_ *GenericSymbol) (*GenericSymbol, error)

	// 309:2: atom_type -> explicit_enum_type: ...
	ExplicitEnumTypeToAtomType(ExplicitEnumType_ *GenericSymbol) (*GenericSymbol, error)

	// 310:2: atom_type -> trait_type: ...
	TraitTypeToAtomType(TraitType_ *GenericSymbol) (*GenericSymbol, error)

	// 313:2: traitable_type -> atom_type: ...
	AtomTypeToTraitableType(AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 314:2: traitable_type -> method_interface: ...
	MethodInterfaceToTraitableType(Exclaim_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 315:2: traitable_type -> trait: ...
	TraitToTraitableType(ExclaimExclaim_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 318:2: trait_intersect_type -> traitable_type: ...
	TraitableTypeToTraitIntersectType(TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 319:2: trait_intersect_type -> intersect: ...
	IntersectToTraitIntersectType(TraitIntersectType_ *GenericSymbol, BitAnd_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 320:2: trait_intersect_type -> difference: ...
	DifferenceToTraitIntersectType(TraitIntersectType_ *GenericSymbol, Sub_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 323:2: trait_union_type -> trait_intersect_type: ...
	TraitIntersectTypeToTraitUnionType(TraitIntersectType_ *GenericSymbol) (*GenericSymbol, error)

	// 324:2: trait_union_type -> union: ...
	UnionToTraitUnionType(TraitUnionType_ *GenericSymbol, BitOr_ *GenericSymbol, TraitIntersectType_ *GenericSymbol) (*GenericSymbol, error)

	// 327:24: implicit_struct_type -> ...
	ToImplicitStructType(Lparen_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 329:24: explicit_struct_type -> ...
	ToExplicitStructType(Struct_ *GenericSymbol, ImplicitStructType_ *GenericSymbol) (*GenericSymbol, error)

	// 331:14: trait_type -> ...
	ToTraitType(Trait_ *GenericSymbol, ImplicitStructType_ *GenericSymbol) (*GenericSymbol, error)

	// 334:22: explicit_enum_type -> ...
	ToExplicitEnumType(Enum_ *GenericSymbol, Lparen_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 337:2: value_type -> trait_union_type: ...
	TraitUnionTypeToValueType(TraitUnionType_ *GenericSymbol) (*GenericSymbol, error)

	// 338:2: value_type -> implicit_enum: ...
	ImplicitEnumToValueType(ValueType_ *GenericSymbol, Or_ *GenericSymbol, TraitUnionType_ *GenericSymbol) (*GenericSymbol, error)

	// 341:2: type_decl -> definition: ...
	DefinitionToTypeDecl(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 342:2: type_decl -> alias: ...
	AliasToTypeDecl(Type_ *GenericSymbol, Identifier_ *GenericSymbol, Equal_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 345:2: field_decl -> explicit: ...
	ExplicitToFieldDecl(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 346:2: field_decl -> implicit: ...
	ImplicitToFieldDecl(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 349:2: parameter_decl -> field_decl: ...
	FieldDeclToParameterDecl(FieldDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 350:2: parameter_decl -> inferred: ...
	InferredToParameterDecl(Identifier_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 353:2: non_empty_parameters -> parameter_decl: ...
	ParameterDeclToNonEmptyParameters(ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 354:2: non_empty_parameters -> add: ...
	AddToNonEmptyParameters(NonEmptyParameters_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 357:2: vararg -> explicit: ...
	ExplicitToVararg(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 358:2: vararg -> implicit: ...
	ImplicitToVararg(Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 359:2: vararg -> inferred: ...
	InferredToVararg(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 362:2: optional_vararg -> vararg: ...
	VarargToOptionalVararg(Vararg_ *GenericSymbol) (*GenericSymbol, error)

	// 363:2: optional_vararg -> vararg2: ...
	Vararg2ToOptionalVararg(Vararg_ *GenericSymbol, Comma_ *GenericSymbol) (*GenericSymbol, error)

	// 364:2: optional_vararg -> nil: ...
	NilToOptionalVararg() (*GenericSymbol, error)

	// 367:2: parameters -> non_empty_parameters: ...
	NonEmptyParametersToParameters(NonEmptyParameters_ *GenericSymbol) (*GenericSymbol, error)

	// 368:2: parameters -> mixed: ...
	MixedToParameters(NonEmptyParameters_ *GenericSymbol, Comma_ *GenericSymbol, OptionalVararg_ *GenericSymbol) (*GenericSymbol, error)

	// 369:2: parameters -> optional_vararg: ...
	OptionalVarargToParameters(OptionalVararg_ *GenericSymbol) (*GenericSymbol, error)

	// 372:2: implicit_func_type -> typed: ...
	TypedToImplicitFuncType(Lparen_ *GenericSymbol, Parameters_ *GenericSymbol, Rparen_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 373:2: implicit_func_type -> inferred: ...
	InferredToImplicitFuncType(Lparen_ *GenericSymbol, Parameters_ *GenericSymbol, Rparen_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 374:2: implicit_func_type -> implicit_unit: ...
	ImplicitUnitToImplicitFuncType(Lparen_ *GenericSymbol, Parameters_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 376:22: explicit_func_type -> ...
	ToExplicitFuncType(Func_ *GenericSymbol, ImplicitFuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 379:2: optional_receiver -> receiver: ...
	ReceiverToOptionalReceiver(Lparen_ *GenericSymbol, ParameterDecl_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 380:2: optional_receiver -> nil: ...
	NilToOptionalReceiver() (*GenericSymbol, error)

	// 383:2: func_decl -> ...
	ToFuncDecl(Func_ *GenericSymbol, OptionalReceiver_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ImplicitFuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 385:12: func_def -> ...
	ToFuncDef(FuncDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 388:2: package_decl -> no_spec: ...
	NoSpecToPackageDecl(Package_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 389:2: package_decl -> with_spec: ...
	WithSpecToPackageDecl(Package_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, PackageStatements_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 391:26: package_statement_body -> ...
	ToPackageStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 394:2: package_statement -> implicit: ...
	ImplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 395:2: package_statement -> explicit: ...
	ExplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 398:2: package_statements -> empty_list: ...
	EmptyListToPackageStatements() (*GenericSymbol, error)

	// 399:2: package_statements -> add: ...
	AddToPackageStatements(PackageStatements_ *GenericSymbol, PackageStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 403:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 404:2: lex_internal_tokens -> COMMENT: ...
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
	case DollarLbraceToken:
		return "DOLLAR_LBRACE"
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
	case ArgumentsType:
		return "arguments"
	case CallExprType:
		return "call_expr"
	case AccessExprType:
		return "access_expr"
	case UnaryOpType:
		return "unary_op"
	case UnaryExprType:
		return "unary_expr"
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
	case GenericParametersType:
		return "generic_parameters"
	case OptionalGenericParametersType:
		return "optional_generic_parameters"
	case AtomTypeType:
		return "atom_type"
	case TraitableTypeType:
		return "traitable_type"
	case TraitIntersectTypeType:
		return "trait_intersect_type"
	case TraitUnionTypeType:
		return "trait_union_type"
	case ImplicitStructTypeType:
		return "implicit_struct_type"
	case ExplicitStructTypeType:
		return "explicit_struct_type"
	case TraitTypeType:
		return "trait_type"
	case ExplicitEnumTypeType:
		return "explicit_enum_type"
	case ValueTypeType:
		return "value_type"
	case TypeDeclType:
		return "type_decl"
	case FieldDeclType:
		return "field_decl"
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
	case FuncDeclType:
		return "func_decl"
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
	ArgumentsType                          = SymbolId(338)
	CallExprType                           = SymbolId(339)
	AccessExprType                         = SymbolId(340)
	UnaryOpType                            = SymbolId(341)
	UnaryExprType                          = SymbolId(342)
	MulOpType                              = SymbolId(343)
	MulExprType                            = SymbolId(344)
	AddOpType                              = SymbolId(345)
	AddExprType                            = SymbolId(346)
	CmpOpType                              = SymbolId(347)
	CmpExprType                            = SymbolId(348)
	AndExprType                            = SymbolId(349)
	OrExprType                             = SymbolId(350)
	SequenceExprType                       = SymbolId(351)
	OptionalJumpLabelType                  = SymbolId(352)
	OptionalExpressionOrImplicitStructType = SymbolId(353)
	JumpTypeType                           = SymbolId(354)
	OpOneAssignType                        = SymbolId(355)
	BinaryOpAssignType                     = SymbolId(356)
	ExpressionOrImplicitStructType         = SymbolId(357)
	UnsafeStatementType                    = SymbolId(358)
	StatementBodyType                      = SymbolId(359)
	StatementType                          = SymbolId(360)
	StatementsType                         = SymbolId(361)
	OptionalLabelDeclType                  = SymbolId(362)
	BlockBodyType                          = SymbolId(363)
	BlockExprType                          = SymbolId(364)
	IfExprType                             = SymbolId(365)
	SwitchExprType                         = SymbolId(366)
	LoopExprType                           = SymbolId(367)
	ExpressionType                         = SymbolId(368)
	GenericParametersType                  = SymbolId(369)
	OptionalGenericParametersType          = SymbolId(370)
	AtomTypeType                           = SymbolId(371)
	TraitableTypeType                      = SymbolId(372)
	TraitIntersectTypeType                 = SymbolId(373)
	TraitUnionTypeType                     = SymbolId(374)
	ImplicitStructTypeType                 = SymbolId(375)
	ExplicitStructTypeType                 = SymbolId(376)
	TraitTypeType                          = SymbolId(377)
	ExplicitEnumTypeType                   = SymbolId(378)
	ValueTypeType                          = SymbolId(379)
	TypeDeclType                           = SymbolId(380)
	FieldDeclType                          = SymbolId(381)
	ParameterDeclType                      = SymbolId(382)
	NonEmptyParametersType                 = SymbolId(383)
	VarargType                             = SymbolId(384)
	OptionalVarargType                     = SymbolId(385)
	ParametersType                         = SymbolId(386)
	ImplicitFuncTypeType                   = SymbolId(387)
	ExplicitFuncTypeType                   = SymbolId(388)
	OptionalReceiverType                   = SymbolId(389)
	FuncDeclType                           = SymbolId(390)
	FuncDefType                            = SymbolId(391)
	PackageDeclType                        = SymbolId(392)
	PackageStatementBodyType               = SymbolId(393)
	PackageStatementType                   = SymbolId(394)
	PackageStatementsType                  = SymbolId(395)
	LexInternalTokensType                  = SymbolId(396)
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
	_ReduceToGenericArguments                                             = _ReduceType(16)
	_ReduceToArguments                                                    = _ReduceType(17)
	_ReduceConcreteToCallExpr                                             = _ReduceType(18)
	_ReduceGenericToCallExpr                                              = _ReduceType(19)
	_ReduceAtomExprToAccessExpr                                           = _ReduceType(20)
	_ReduceAccessToAccessExpr                                             = _ReduceType(21)
	_ReduceCallExprToAccessExpr                                           = _ReduceType(22)
	_ReduceIndexToAccessExpr                                              = _ReduceType(23)
	_ReduceNotToUnaryOp                                                   = _ReduceType(24)
	_ReduceBitNegToUnaryOp                                                = _ReduceType(25)
	_ReduceSubToUnaryOp                                                   = _ReduceType(26)
	_ReduceMulToUnaryOp                                                   = _ReduceType(27)
	_ReduceBitAndToUnaryOp                                                = _ReduceType(28)
	_ReduceAccessExprToUnaryExpr                                          = _ReduceType(29)
	_ReduceOpToUnaryExpr                                                  = _ReduceType(30)
	_ReduceMulToMulOp                                                     = _ReduceType(31)
	_ReduceDivToMulOp                                                     = _ReduceType(32)
	_ReduceModToMulOp                                                     = _ReduceType(33)
	_ReduceBitAndToMulOp                                                  = _ReduceType(34)
	_ReduceBitLshiftToMulOp                                               = _ReduceType(35)
	_ReduceBitRshiftToMulOp                                               = _ReduceType(36)
	_ReduceUnaryExprToMulExpr                                             = _ReduceType(37)
	_ReduceOpToMulExpr                                                    = _ReduceType(38)
	_ReduceAddToAddOp                                                     = _ReduceType(39)
	_ReduceSubToAddOp                                                     = _ReduceType(40)
	_ReduceBitOrToAddOp                                                   = _ReduceType(41)
	_ReduceBitXorToAddOp                                                  = _ReduceType(42)
	_ReduceMulExprToAddExpr                                               = _ReduceType(43)
	_ReduceOpToAddExpr                                                    = _ReduceType(44)
	_ReduceEqualToCmpOp                                                   = _ReduceType(45)
	_ReduceNotEqualToCmpOp                                                = _ReduceType(46)
	_ReduceLessToCmpOp                                                    = _ReduceType(47)
	_ReduceLessOrEqualToCmpOp                                             = _ReduceType(48)
	_ReduceGreaterToCmpOp                                                 = _ReduceType(49)
	_ReduceGreaterOrEqualToCmpOp                                          = _ReduceType(50)
	_ReduceAddExprToCmpExpr                                               = _ReduceType(51)
	_ReduceOpToCmpExpr                                                    = _ReduceType(52)
	_ReduceCmpExprToAndExpr                                               = _ReduceType(53)
	_ReduceOpToAndExpr                                                    = _ReduceType(54)
	_ReduceAndExprToOrExpr                                                = _ReduceType(55)
	_ReduceOpToOrExpr                                                     = _ReduceType(56)
	_ReduceToSequenceExpr                                                 = _ReduceType(57)
	_ReduceJumpLabelToOptionalJumpLabel                                   = _ReduceType(58)
	_ReduceUnlabelledToOptionalJumpLabel                                  = _ReduceType(59)
	_ReduceExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct = _ReduceType(60)
	_ReduceNilToOptionalExpressionOrImplicitStruct                        = _ReduceType(61)
	_ReduceReturnToJumpType                                               = _ReduceType(62)
	_ReduceBreakToJumpType                                                = _ReduceType(63)
	_ReduceContinueToJumpType                                             = _ReduceType(64)
	_ReduceAddOneAssignToOpOneAssign                                      = _ReduceType(65)
	_ReduceSubOneAssignToOpOneAssign                                      = _ReduceType(66)
	_ReduceAddAssignToBinaryOpAssign                                      = _ReduceType(67)
	_ReduceSubAssignToBinaryOpAssign                                      = _ReduceType(68)
	_ReduceMulAssignToBinaryOpAssign                                      = _ReduceType(69)
	_ReduceDivAssignToBinaryOpAssign                                      = _ReduceType(70)
	_ReduceModAssignToBinaryOpAssign                                      = _ReduceType(71)
	_ReduceBitNegAssignToBinaryOpAssign                                   = _ReduceType(72)
	_ReduceBitAndAssignToBinaryOpAssign                                   = _ReduceType(73)
	_ReduceBitOrAssignToBinaryOpAssign                                    = _ReduceType(74)
	_ReduceBitXorAssignToBinaryOpAssign                                   = _ReduceType(75)
	_ReduceBitLshiftAssignToBinaryOpAssign                                = _ReduceType(76)
	_ReduceBitRshiftAssignToBinaryOpAssign                                = _ReduceType(77)
	_ReduceExpressionToExpressionOrImplicitStruct                         = _ReduceType(78)
	_ReduceImplicitStructToExpressionOrImplicitStruct                     = _ReduceType(79)
	_ReduceToUnsafeStatement                                              = _ReduceType(80)
	_ReduceUnsafeStatementToStatementBody                                 = _ReduceType(81)
	_ReduceExpressionOrImplicitStructToStatementBody                      = _ReduceType(82)
	_ReduceAsyncToStatementBody                                           = _ReduceType(83)
	_ReduceJumpToStatementBody                                            = _ReduceType(84)
	_ReduceOpOneAssignToStatementBody                                     = _ReduceType(85)
	_ReduceBinaryOpAssignToStatementBody                                  = _ReduceType(86)
	_ReduceImplicitToStatement                                            = _ReduceType(87)
	_ReduceExplicitToStatement                                            = _ReduceType(88)
	_ReduceEmptyListToStatements                                          = _ReduceType(89)
	_ReduceAddToStatements                                                = _ReduceType(90)
	_ReduceLabelDeclToOptionalLabelDecl                                   = _ReduceType(91)
	_ReduceUnlabelledToOptionalLabelDecl                                  = _ReduceType(92)
	_ReduceToBlockBody                                                    = _ReduceType(93)
	_ReduceToBlockExpr                                                    = _ReduceType(94)
	_ReduceNoElseToIfExpr                                                 = _ReduceType(95)
	_ReduceIfElseToIfExpr                                                 = _ReduceType(96)
	_ReduceMultiIfElseToIfExpr                                            = _ReduceType(97)
	_ReduceToSwitchExpr                                                   = _ReduceType(98)
	_ReduceInfiniteToLoopExpr                                             = _ReduceType(99)
	_ReduceWhileToLoopExpr                                                = _ReduceType(100)
	_ReduceSequenceExprToExpression                                       = _ReduceType(101)
	_ReduceIfExprToExpression                                             = _ReduceType(102)
	_ReduceSwitchExprToExpression                                         = _ReduceType(103)
	_ReduceLoopExprToExpression                                           = _ReduceType(104)
	_ReduceToGenericParameters                                            = _ReduceType(105)
	_ReduceGenericToOptionalGenericParameters                             = _ReduceType(106)
	_ReduceNilToOptionalGenericParameters                                 = _ReduceType(107)
	_ReduceNamedToAtomType                                                = _ReduceType(108)
	_ReduceExplicitStructTypeToAtomType                                   = _ReduceType(109)
	_ReduceImplicitStructTypeToAtomType                                   = _ReduceType(110)
	_ReduceExplicitEnumTypeToAtomType                                     = _ReduceType(111)
	_ReduceTraitTypeToAtomType                                            = _ReduceType(112)
	_ReduceAtomTypeToTraitableType                                        = _ReduceType(113)
	_ReduceMethodInterfaceToTraitableType                                 = _ReduceType(114)
	_ReduceTraitToTraitableType                                           = _ReduceType(115)
	_ReduceTraitableTypeToTraitIntersectType                              = _ReduceType(116)
	_ReduceIntersectToTraitIntersectType                                  = _ReduceType(117)
	_ReduceDifferenceToTraitIntersectType                                 = _ReduceType(118)
	_ReduceTraitIntersectTypeToTraitUnionType                             = _ReduceType(119)
	_ReduceUnionToTraitUnionType                                          = _ReduceType(120)
	_ReduceToImplicitStructType                                           = _ReduceType(121)
	_ReduceToExplicitStructType                                           = _ReduceType(122)
	_ReduceToTraitType                                                    = _ReduceType(123)
	_ReduceToExplicitEnumType                                             = _ReduceType(124)
	_ReduceTraitUnionTypeToValueType                                      = _ReduceType(125)
	_ReduceImplicitEnumToValueType                                        = _ReduceType(126)
	_ReduceDefinitionToTypeDecl                                           = _ReduceType(127)
	_ReduceAliasToTypeDecl                                                = _ReduceType(128)
	_ReduceExplicitToFieldDecl                                            = _ReduceType(129)
	_ReduceImplicitToFieldDecl                                            = _ReduceType(130)
	_ReduceFieldDeclToParameterDecl                                       = _ReduceType(131)
	_ReduceInferredToParameterDecl                                        = _ReduceType(132)
	_ReduceParameterDeclToNonEmptyParameters                              = _ReduceType(133)
	_ReduceAddToNonEmptyParameters                                        = _ReduceType(134)
	_ReduceExplicitToVararg                                               = _ReduceType(135)
	_ReduceImplicitToVararg                                               = _ReduceType(136)
	_ReduceInferredToVararg                                               = _ReduceType(137)
	_ReduceVarargToOptionalVararg                                         = _ReduceType(138)
	_ReduceVararg2ToOptionalVararg                                        = _ReduceType(139)
	_ReduceNilToOptionalVararg                                            = _ReduceType(140)
	_ReduceNonEmptyParametersToParameters                                 = _ReduceType(141)
	_ReduceMixedToParameters                                              = _ReduceType(142)
	_ReduceOptionalVarargToParameters                                     = _ReduceType(143)
	_ReduceTypedToImplicitFuncType                                        = _ReduceType(144)
	_ReduceInferredToImplicitFuncType                                     = _ReduceType(145)
	_ReduceImplicitUnitToImplicitFuncType                                 = _ReduceType(146)
	_ReduceToExplicitFuncType                                             = _ReduceType(147)
	_ReduceReceiverToOptionalReceiver                                     = _ReduceType(148)
	_ReduceNilToOptionalReceiver                                          = _ReduceType(149)
	_ReduceToFuncDecl                                                     = _ReduceType(150)
	_ReduceToFuncDef                                                      = _ReduceType(151)
	_ReduceNoSpecToPackageDecl                                            = _ReduceType(152)
	_ReduceWithSpecToPackageDecl                                          = _ReduceType(153)
	_ReduceToPackageStatementBody                                         = _ReduceType(154)
	_ReduceImplicitToPackageStatement                                     = _ReduceType(155)
	_ReduceExplicitToPackageStatement                                     = _ReduceType(156)
	_ReduceEmptyListToPackageStatements                                   = _ReduceType(157)
	_ReduceAddToPackageStatements                                         = _ReduceType(158)
	_ReduceSpacesToLexInternalTokens                                      = _ReduceType(159)
	_ReduceCommentToLexInternalTokens                                     = _ReduceType(160)
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
	case _ReduceToGenericArguments:
		return "ToGenericArguments"
	case _ReduceToArguments:
		return "ToArguments"
	case _ReduceConcreteToCallExpr:
		return "ConcreteToCallExpr"
	case _ReduceGenericToCallExpr:
		return "GenericToCallExpr"
	case _ReduceAtomExprToAccessExpr:
		return "AtomExprToAccessExpr"
	case _ReduceAccessToAccessExpr:
		return "AccessToAccessExpr"
	case _ReduceCallExprToAccessExpr:
		return "CallExprToAccessExpr"
	case _ReduceIndexToAccessExpr:
		return "IndexToAccessExpr"
	case _ReduceNotToUnaryOp:
		return "NotToUnaryOp"
	case _ReduceBitNegToUnaryOp:
		return "BitNegToUnaryOp"
	case _ReduceSubToUnaryOp:
		return "SubToUnaryOp"
	case _ReduceMulToUnaryOp:
		return "MulToUnaryOp"
	case _ReduceBitAndToUnaryOp:
		return "BitAndToUnaryOp"
	case _ReduceAccessExprToUnaryExpr:
		return "AccessExprToUnaryExpr"
	case _ReduceOpToUnaryExpr:
		return "OpToUnaryExpr"
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
	case _ReduceUnaryExprToMulExpr:
		return "UnaryExprToMulExpr"
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
	case _ReduceToGenericParameters:
		return "ToGenericParameters"
	case _ReduceGenericToOptionalGenericParameters:
		return "GenericToOptionalGenericParameters"
	case _ReduceNilToOptionalGenericParameters:
		return "NilToOptionalGenericParameters"
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
	case _ReduceTraitableTypeToTraitIntersectType:
		return "TraitableTypeToTraitIntersectType"
	case _ReduceIntersectToTraitIntersectType:
		return "IntersectToTraitIntersectType"
	case _ReduceDifferenceToTraitIntersectType:
		return "DifferenceToTraitIntersectType"
	case _ReduceTraitIntersectTypeToTraitUnionType:
		return "TraitIntersectTypeToTraitUnionType"
	case _ReduceUnionToTraitUnionType:
		return "UnionToTraitUnionType"
	case _ReduceToImplicitStructType:
		return "ToImplicitStructType"
	case _ReduceToExplicitStructType:
		return "ToExplicitStructType"
	case _ReduceToTraitType:
		return "ToTraitType"
	case _ReduceToExplicitEnumType:
		return "ToExplicitEnumType"
	case _ReduceTraitUnionTypeToValueType:
		return "TraitUnionTypeToValueType"
	case _ReduceImplicitEnumToValueType:
		return "ImplicitEnumToValueType"
	case _ReduceDefinitionToTypeDecl:
		return "DefinitionToTypeDecl"
	case _ReduceAliasToTypeDecl:
		return "AliasToTypeDecl"
	case _ReduceExplicitToFieldDecl:
		return "ExplicitToFieldDecl"
	case _ReduceImplicitToFieldDecl:
		return "ImplicitToFieldDecl"
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
	case _ReduceToFuncDecl:
		return "ToFuncDecl"
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
	case _EndMarker, SpacesToken, NewlinesToken, CommentToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, ReturnToken, BreakToken, ContinueToken, PackageToken, UnsafeToken, TypeToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, DollarLbraceToken, DotdotdotToken, ExclaimToken, ExclaimExclaimToken, LabelDeclToken, JumpLabelToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, NotToken, AndToken, OrToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, LexErrorToken:
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
	case _ReduceToGenericArguments:
		symbol.SymbolId_ = GenericArgumentsType
		symbol.Generic_, err = reducer.ToGenericArguments()
	case _ReduceToArguments:
		symbol.SymbolId_ = ArgumentsType
		symbol.Generic_, err = reducer.ToArguments()
	case _ReduceConcreteToCallExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CallExprType
		symbol.Generic_, err = reducer.ConcreteToCallExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceGenericToCallExpr:
		args := stack[len(stack)-7:]
		stack = stack[:len(stack)-7]
		symbol.SymbolId_ = CallExprType
		symbol.Generic_, err = reducer.GenericToCallExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_, args[6].Generic_)
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
	case _ReduceNotToUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpType
		symbol.Generic_, err = reducer.NotToUnaryOp(args[0].Generic_)
	case _ReduceBitNegToUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpType
		symbol.Generic_, err = reducer.BitNegToUnaryOp(args[0].Generic_)
	case _ReduceSubToUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpType
		symbol.Generic_, err = reducer.SubToUnaryOp(args[0].Generic_)
	case _ReduceMulToUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpType
		symbol.Generic_, err = reducer.MulToUnaryOp(args[0].Generic_)
	case _ReduceBitAndToUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpType
		symbol.Generic_, err = reducer.BitAndToUnaryOp(args[0].Generic_)
	case _ReduceAccessExprToUnaryExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryExprType
		symbol.Generic_, err = reducer.AccessExprToUnaryExpr(args[0].Generic_)
	case _ReduceOpToUnaryExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = UnaryExprType
		symbol.Generic_, err = reducer.OpToUnaryExpr(args[0].Generic_, args[1].Generic_)
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
	case _ReduceUnaryExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		symbol.Generic_, err = reducer.UnaryExprToMulExpr(args[0].Generic_)
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
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = SwitchExprType
		symbol.Generic_, err = reducer.ToSwitchExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	case _ReduceToGenericParameters:
		symbol.SymbolId_ = GenericParametersType
		symbol.Generic_, err = reducer.ToGenericParameters()
	case _ReduceGenericToOptionalGenericParameters:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OptionalGenericParametersType
		symbol.Generic_, err = reducer.GenericToOptionalGenericParameters(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceNilToOptionalGenericParameters:
		symbol.SymbolId_ = OptionalGenericParametersType
		symbol.Generic_, err = reducer.NilToOptionalGenericParameters()
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
	case _ReduceTraitableTypeToTraitIntersectType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitIntersectTypeType
		symbol.Generic_, err = reducer.TraitableTypeToTraitIntersectType(args[0].Generic_)
	case _ReduceIntersectToTraitIntersectType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TraitIntersectTypeType
		symbol.Generic_, err = reducer.IntersectToTraitIntersectType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceDifferenceToTraitIntersectType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TraitIntersectTypeType
		symbol.Generic_, err = reducer.DifferenceToTraitIntersectType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceTraitIntersectTypeToTraitUnionType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitUnionTypeType
		symbol.Generic_, err = reducer.TraitIntersectTypeToTraitUnionType(args[0].Generic_)
	case _ReduceUnionToTraitUnionType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TraitUnionTypeType
		symbol.Generic_, err = reducer.UnionToTraitUnionType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceToImplicitStructType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImplicitStructTypeType
		symbol.Generic_, err = reducer.ToImplicitStructType(args[0].Generic_, args[1].Generic_)
	case _ReduceToExplicitStructType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExplicitStructTypeType
		symbol.Generic_, err = reducer.ToExplicitStructType(args[0].Generic_, args[1].Generic_)
	case _ReduceToTraitType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = TraitTypeType
		symbol.Generic_, err = reducer.ToTraitType(args[0].Generic_, args[1].Generic_)
	case _ReduceToExplicitEnumType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumTypeType
		symbol.Generic_, err = reducer.ToExplicitEnumType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceTraitUnionTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.TraitUnionTypeToValueType(args[0].Generic_)
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
	case _ReduceToFuncDecl:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = FuncDeclType
		symbol.Generic_, err = reducer.ToFuncDecl(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
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
	_ReduceToGenericArgumentsAction                                             = &_Action{_ReduceAction, 0, _ReduceToGenericArguments}
	_ReduceToArgumentsAction                                                    = &_Action{_ReduceAction, 0, _ReduceToArguments}
	_ReduceConcreteToCallExprAction                                             = &_Action{_ReduceAction, 0, _ReduceConcreteToCallExpr}
	_ReduceGenericToCallExprAction                                              = &_Action{_ReduceAction, 0, _ReduceGenericToCallExpr}
	_ReduceAtomExprToAccessExprAction                                           = &_Action{_ReduceAction, 0, _ReduceAtomExprToAccessExpr}
	_ReduceAccessToAccessExprAction                                             = &_Action{_ReduceAction, 0, _ReduceAccessToAccessExpr}
	_ReduceCallExprToAccessExprAction                                           = &_Action{_ReduceAction, 0, _ReduceCallExprToAccessExpr}
	_ReduceIndexToAccessExprAction                                              = &_Action{_ReduceAction, 0, _ReduceIndexToAccessExpr}
	_ReduceNotToUnaryOpAction                                                   = &_Action{_ReduceAction, 0, _ReduceNotToUnaryOp}
	_ReduceBitNegToUnaryOpAction                                                = &_Action{_ReduceAction, 0, _ReduceBitNegToUnaryOp}
	_ReduceSubToUnaryOpAction                                                   = &_Action{_ReduceAction, 0, _ReduceSubToUnaryOp}
	_ReduceMulToUnaryOpAction                                                   = &_Action{_ReduceAction, 0, _ReduceMulToUnaryOp}
	_ReduceBitAndToUnaryOpAction                                                = &_Action{_ReduceAction, 0, _ReduceBitAndToUnaryOp}
	_ReduceAccessExprToUnaryExprAction                                          = &_Action{_ReduceAction, 0, _ReduceAccessExprToUnaryExpr}
	_ReduceOpToUnaryExprAction                                                  = &_Action{_ReduceAction, 0, _ReduceOpToUnaryExpr}
	_ReduceMulToMulOpAction                                                     = &_Action{_ReduceAction, 0, _ReduceMulToMulOp}
	_ReduceDivToMulOpAction                                                     = &_Action{_ReduceAction, 0, _ReduceDivToMulOp}
	_ReduceModToMulOpAction                                                     = &_Action{_ReduceAction, 0, _ReduceModToMulOp}
	_ReduceBitAndToMulOpAction                                                  = &_Action{_ReduceAction, 0, _ReduceBitAndToMulOp}
	_ReduceBitLshiftToMulOpAction                                               = &_Action{_ReduceAction, 0, _ReduceBitLshiftToMulOp}
	_ReduceBitRshiftToMulOpAction                                               = &_Action{_ReduceAction, 0, _ReduceBitRshiftToMulOp}
	_ReduceUnaryExprToMulExprAction                                             = &_Action{_ReduceAction, 0, _ReduceUnaryExprToMulExpr}
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
	_ReduceToGenericParametersAction                                            = &_Action{_ReduceAction, 0, _ReduceToGenericParameters}
	_ReduceGenericToOptionalGenericParametersAction                             = &_Action{_ReduceAction, 0, _ReduceGenericToOptionalGenericParameters}
	_ReduceNilToOptionalGenericParametersAction                                 = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericParameters}
	_ReduceNamedToAtomTypeAction                                                = &_Action{_ReduceAction, 0, _ReduceNamedToAtomType}
	_ReduceExplicitStructTypeToAtomTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceExplicitStructTypeToAtomType}
	_ReduceImplicitStructTypeToAtomTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceImplicitStructTypeToAtomType}
	_ReduceExplicitEnumTypeToAtomTypeAction                                     = &_Action{_ReduceAction, 0, _ReduceExplicitEnumTypeToAtomType}
	_ReduceTraitTypeToAtomTypeAction                                            = &_Action{_ReduceAction, 0, _ReduceTraitTypeToAtomType}
	_ReduceAtomTypeToTraitableTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceAtomTypeToTraitableType}
	_ReduceMethodInterfaceToTraitableTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceMethodInterfaceToTraitableType}
	_ReduceTraitToTraitableTypeAction                                           = &_Action{_ReduceAction, 0, _ReduceTraitToTraitableType}
	_ReduceTraitableTypeToTraitIntersectTypeAction                              = &_Action{_ReduceAction, 0, _ReduceTraitableTypeToTraitIntersectType}
	_ReduceIntersectToTraitIntersectTypeAction                                  = &_Action{_ReduceAction, 0, _ReduceIntersectToTraitIntersectType}
	_ReduceDifferenceToTraitIntersectTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceDifferenceToTraitIntersectType}
	_ReduceTraitIntersectTypeToTraitUnionTypeAction                             = &_Action{_ReduceAction, 0, _ReduceTraitIntersectTypeToTraitUnionType}
	_ReduceUnionToTraitUnionTypeAction                                          = &_Action{_ReduceAction, 0, _ReduceUnionToTraitUnionType}
	_ReduceToImplicitStructTypeAction                                           = &_Action{_ReduceAction, 0, _ReduceToImplicitStructType}
	_ReduceToExplicitStructTypeAction                                           = &_Action{_ReduceAction, 0, _ReduceToExplicitStructType}
	_ReduceToTraitTypeAction                                                    = &_Action{_ReduceAction, 0, _ReduceToTraitType}
	_ReduceToExplicitEnumTypeAction                                             = &_Action{_ReduceAction, 0, _ReduceToExplicitEnumType}
	_ReduceTraitUnionTypeToValueTypeAction                                      = &_Action{_ReduceAction, 0, _ReduceTraitUnionTypeToValueType}
	_ReduceImplicitEnumToValueTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceImplicitEnumToValueType}
	_ReduceDefinitionToTypeDeclAction                                           = &_Action{_ReduceAction, 0, _ReduceDefinitionToTypeDecl}
	_ReduceAliasToTypeDeclAction                                                = &_Action{_ReduceAction, 0, _ReduceAliasToTypeDecl}
	_ReduceExplicitToFieldDeclAction                                            = &_Action{_ReduceAction, 0, _ReduceExplicitToFieldDecl}
	_ReduceImplicitToFieldDeclAction                                            = &_Action{_ReduceAction, 0, _ReduceImplicitToFieldDecl}
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
	_ReduceToFuncDeclAction                                                     = &_Action{_ReduceAction, 0, _ReduceToFuncDecl}
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
	{_State3, FuncDeclType}:                             _GotoState14Action,
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
	{_State4, UnaryOpType}:                              _GotoState49Action,
	{_State4, UnaryExprType}:                            _GotoState48Action,
	{_State4, MulExprType}:                              _GotoState44Action,
	{_State4, AddExprType}:                              _GotoState33Action,
	{_State4, CmpExprType}:                              _GotoState40Action,
	{_State4, AndExprType}:                              _GotoState34Action,
	{_State4, OrExprType}:                               _GotoState46Action,
	{_State4, SequenceExprType}:                         _GotoState47Action,
	{_State4, OptionalLabelDeclType}:                    _GotoState45Action,
	{_State4, BlockExprType}:                            _GotoState38Action,
	{_State4, ExpressionType}:                           _GotoState9Action,
	{_State4, ExplicitStructTypeType}:                   _GotoState42Action,
	{_State4, ExplicitFuncTypeType}:                     _GotoState41Action,
	{_State5, SpacesToken}:                              _GotoState51Action,
	{_State5, CommentToken}:                             _GotoState50Action,
	{_State5, LexInternalTokensType}:                    _GotoState10Action,
	{_State11, IdentifierToken}:                         _GotoState52Action,
	{_State12, IdentifierToken}:                         _GotoState53Action,
	{_State13, LparenToken}:                             _GotoState54Action,
	{_State13, OptionalReceiverType}:                    _GotoState55Action,
	{_State14, LbraceToken}:                             _GotoState56Action,
	{_State14, BlockBodyType}:                           _GotoState57Action,
	{_State19, LparenToken}:                             _GotoState58Action,
	{_State19, ImplicitFuncTypeType}:                    _GotoState59Action,
	{_State24, ArgumentsType}:                           _GotoState60Action,
	{_State29, LparenToken}:                             _GotoState61Action,
	{_State29, ImplicitStructTypeType}:                  _GotoState62Action,
	{_State32, LparenToken}:                             _GotoState66Action,
	{_State32, LbracketToken}:                           _GotoState65Action,
	{_State32, DotToken}:                                _GotoState64Action,
	{_State32, DollarLbraceToken}:                       _GotoState63Action,
	{_State33, AddToken}:                                _GotoState67Action,
	{_State33, SubToken}:                                _GotoState70Action,
	{_State33, BitXorToken}:                             _GotoState69Action,
	{_State33, BitOrToken}:                              _GotoState68Action,
	{_State33, AddOpType}:                               _GotoState71Action,
	{_State34, AndToken}:                                _GotoState72Action,
	{_State40, EqualToken}:                              _GotoState73Action,
	{_State40, NotEqualToken}:                           _GotoState78Action,
	{_State40, LessToken}:                               _GotoState76Action,
	{_State40, LessOrEqualToken}:                        _GotoState77Action,
	{_State40, GreaterToken}:                            _GotoState74Action,
	{_State40, GreaterOrEqualToken}:                     _GotoState75Action,
	{_State40, CmpOpType}:                               _GotoState79Action,
	{_State41, LbraceToken}:                             _GotoState56Action,
	{_State41, BlockBodyType}:                           _GotoState80Action,
	{_State42, LparenToken}:                             _GotoState81Action,
	{_State44, MulToken}:                                _GotoState87Action,
	{_State44, DivToken}:                                _GotoState85Action,
	{_State44, ModToken}:                                _GotoState86Action,
	{_State44, BitAndToken}:                             _GotoState82Action,
	{_State44, BitLshiftToken}:                          _GotoState83Action,
	{_State44, BitRshiftToken}:                          _GotoState84Action,
	{_State44, MulOpType}:                               _GotoState88Action,
	{_State45, IfToken}:                                 _GotoState90Action,
	{_State45, SwitchToken}:                             _GotoState91Action,
	{_State45, ForToken}:                                _GotoState89Action,
	{_State45, LbraceToken}:                             _GotoState56Action,
	{_State45, BlockBodyType}:                           _GotoState92Action,
	{_State45, IfExprType}:                              _GotoState93Action,
	{_State45, SwitchExprType}:                          _GotoState95Action,
	{_State45, LoopExprType}:                            _GotoState94Action,
	{_State46, OrToken}:                                 _GotoState96Action,
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
	{_State49, UnaryOpType}:                             _GotoState49Action,
	{_State49, UnaryExprType}:                           _GotoState98Action,
	{_State49, OptionalLabelDeclType}:                   _GotoState97Action,
	{_State49, BlockExprType}:                           _GotoState38Action,
	{_State49, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State49, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State52, LparenToken}:                             _GotoState99Action,
	{_State53, DollarLbraceToken}:                       _GotoState100Action,
	{_State53, EqualToken}:                              _GotoState101Action,
	{_State53, OptionalGenericParametersType}:           _GotoState102Action,
	{_State54, IdentifierToken}:                         _GotoState106Action,
	{_State54, StructToken}:                             _GotoState29Action,
	{_State54, EnumToken}:                               _GotoState103Action,
	{_State54, TraitToken}:                              _GotoState107Action,
	{_State54, LparenToken}:                             _GotoState61Action,
	{_State54, ExclaimToken}:                            _GotoState104Action,
	{_State54, ExclaimExclaimToken}:                     _GotoState105Action,
	{_State54, AtomTypeType}:                            _GotoState108Action,
	{_State54, TraitableTypeType}:                       _GotoState117Action,
	{_State54, TraitIntersectTypeType}:                  _GotoState114Action,
	{_State54, TraitUnionTypeType}:                      _GotoState116Action,
	{_State54, ImplicitStructTypeType}:                  _GotoState112Action,
	{_State54, ExplicitStructTypeType}:                  _GotoState110Action,
	{_State54, TraitTypeType}:                           _GotoState115Action,
	{_State54, ExplicitEnumTypeType}:                    _GotoState109Action,
	{_State54, ValueTypeType}:                           _GotoState118Action,
	{_State54, FieldDeclType}:                           _GotoState111Action,
	{_State54, ParameterDeclType}:                       _GotoState113Action,
	{_State55, IdentifierToken}:                         _GotoState119Action,
	{_State56, StatementsType}:                          _GotoState120Action,
	{_State58, IdentifierToken}:                         _GotoState122Action,
	{_State58, StructToken}:                             _GotoState29Action,
	{_State58, EnumToken}:                               _GotoState103Action,
	{_State58, TraitToken}:                              _GotoState107Action,
	{_State58, LparenToken}:                             _GotoState61Action,
	{_State58, DotdotdotToken}:                          _GotoState121Action,
	{_State58, ExclaimToken}:                            _GotoState104Action,
	{_State58, ExclaimExclaimToken}:                     _GotoState105Action,
	{_State58, AtomTypeType}:                            _GotoState108Action,
	{_State58, TraitableTypeType}:                       _GotoState117Action,
	{_State58, TraitIntersectTypeType}:                  _GotoState114Action,
	{_State58, TraitUnionTypeType}:                      _GotoState116Action,
	{_State58, ImplicitStructTypeType}:                  _GotoState112Action,
	{_State58, ExplicitStructTypeType}:                  _GotoState110Action,
	{_State58, TraitTypeType}:                           _GotoState115Action,
	{_State58, ExplicitEnumTypeType}:                    _GotoState109Action,
	{_State58, ValueTypeType}:                           _GotoState118Action,
	{_State58, FieldDeclType}:                           _GotoState111Action,
	{_State58, ParameterDeclType}:                       _GotoState125Action,
	{_State58, NonEmptyParametersType}:                  _GotoState123Action,
	{_State58, VarargType}:                              _GotoState127Action,
	{_State58, OptionalVarargType}:                      _GotoState124Action,
	{_State58, ParametersType}:                          _GotoState126Action,
	{_State60, RparenToken}:                             _GotoState128Action,
	{_State61, RparenToken}:                             _GotoState129Action,
	{_State63, GenericArgumentsType}:                    _GotoState130Action,
	{_State64, IdentifierToken}:                         _GotoState131Action,
	{_State65, IntegerLiteralToken}:                     _GotoState21Action,
	{_State65, FloatLiteralToken}:                       _GotoState18Action,
	{_State65, RuneLiteralToken}:                        _GotoState27Action,
	{_State65, StringLiteralToken}:                      _GotoState28Action,
	{_State65, IdentifierToken}:                         _GotoState20Action,
	{_State65, TrueToken}:                               _GotoState31Action,
	{_State65, FalseToken}:                              _GotoState17Action,
	{_State65, StructToken}:                             _GotoState29Action,
	{_State65, FuncToken}:                               _GotoState19Action,
	{_State65, LparenToken}:                             _GotoState24Action,
	{_State65, LabelDeclToken}:                          _GotoState22Action,
	{_State65, NotToken}:                                _GotoState26Action,
	{_State65, SubToken}:                                _GotoState30Action,
	{_State65, MulToken}:                                _GotoState25Action,
	{_State65, BitNegToken}:                             _GotoState16Action,
	{_State65, BitAndToken}:                             _GotoState15Action,
	{_State65, LexErrorToken}:                           _GotoState23Action,
	{_State65, LiteralType}:                             _GotoState43Action,
	{_State65, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State65, AnonymousStructExprType}:                 _GotoState36Action,
	{_State65, AtomExprType}:                            _GotoState37Action,
	{_State65, CallExprType}:                            _GotoState39Action,
	{_State65, AccessExprType}:                          _GotoState32Action,
	{_State65, UnaryOpType}:                             _GotoState49Action,
	{_State65, UnaryExprType}:                           _GotoState48Action,
	{_State65, MulExprType}:                             _GotoState44Action,
	{_State65, AddExprType}:                             _GotoState33Action,
	{_State65, CmpExprType}:                             _GotoState40Action,
	{_State65, AndExprType}:                             _GotoState34Action,
	{_State65, OrExprType}:                              _GotoState46Action,
	{_State65, SequenceExprType}:                        _GotoState47Action,
	{_State65, OptionalLabelDeclType}:                   _GotoState45Action,
	{_State65, BlockExprType}:                           _GotoState38Action,
	{_State65, ExpressionType}:                          _GotoState132Action,
	{_State65, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State65, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State66, ArgumentsType}:                           _GotoState133Action,
	{_State71, IntegerLiteralToken}:                     _GotoState21Action,
	{_State71, FloatLiteralToken}:                       _GotoState18Action,
	{_State71, RuneLiteralToken}:                        _GotoState27Action,
	{_State71, StringLiteralToken}:                      _GotoState28Action,
	{_State71, IdentifierToken}:                         _GotoState20Action,
	{_State71, TrueToken}:                               _GotoState31Action,
	{_State71, FalseToken}:                              _GotoState17Action,
	{_State71, StructToken}:                             _GotoState29Action,
	{_State71, FuncToken}:                               _GotoState19Action,
	{_State71, LparenToken}:                             _GotoState24Action,
	{_State71, LabelDeclToken}:                          _GotoState22Action,
	{_State71, NotToken}:                                _GotoState26Action,
	{_State71, SubToken}:                                _GotoState30Action,
	{_State71, MulToken}:                                _GotoState25Action,
	{_State71, BitNegToken}:                             _GotoState16Action,
	{_State71, BitAndToken}:                             _GotoState15Action,
	{_State71, LexErrorToken}:                           _GotoState23Action,
	{_State71, LiteralType}:                             _GotoState43Action,
	{_State71, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State71, AnonymousStructExprType}:                 _GotoState36Action,
	{_State71, AtomExprType}:                            _GotoState37Action,
	{_State71, CallExprType}:                            _GotoState39Action,
	{_State71, AccessExprType}:                          _GotoState32Action,
	{_State71, UnaryOpType}:                             _GotoState49Action,
	{_State71, UnaryExprType}:                           _GotoState48Action,
	{_State71, MulExprType}:                             _GotoState134Action,
	{_State71, OptionalLabelDeclType}:                   _GotoState97Action,
	{_State71, BlockExprType}:                           _GotoState38Action,
	{_State71, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State71, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State72, IntegerLiteralToken}:                     _GotoState21Action,
	{_State72, FloatLiteralToken}:                       _GotoState18Action,
	{_State72, RuneLiteralToken}:                        _GotoState27Action,
	{_State72, StringLiteralToken}:                      _GotoState28Action,
	{_State72, IdentifierToken}:                         _GotoState20Action,
	{_State72, TrueToken}:                               _GotoState31Action,
	{_State72, FalseToken}:                              _GotoState17Action,
	{_State72, StructToken}:                             _GotoState29Action,
	{_State72, FuncToken}:                               _GotoState19Action,
	{_State72, LparenToken}:                             _GotoState24Action,
	{_State72, LabelDeclToken}:                          _GotoState22Action,
	{_State72, NotToken}:                                _GotoState26Action,
	{_State72, SubToken}:                                _GotoState30Action,
	{_State72, MulToken}:                                _GotoState25Action,
	{_State72, BitNegToken}:                             _GotoState16Action,
	{_State72, BitAndToken}:                             _GotoState15Action,
	{_State72, LexErrorToken}:                           _GotoState23Action,
	{_State72, LiteralType}:                             _GotoState43Action,
	{_State72, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State72, AnonymousStructExprType}:                 _GotoState36Action,
	{_State72, AtomExprType}:                            _GotoState37Action,
	{_State72, CallExprType}:                            _GotoState39Action,
	{_State72, AccessExprType}:                          _GotoState32Action,
	{_State72, UnaryOpType}:                             _GotoState49Action,
	{_State72, UnaryExprType}:                           _GotoState48Action,
	{_State72, MulExprType}:                             _GotoState44Action,
	{_State72, AddExprType}:                             _GotoState33Action,
	{_State72, CmpExprType}:                             _GotoState135Action,
	{_State72, OptionalLabelDeclType}:                   _GotoState97Action,
	{_State72, BlockExprType}:                           _GotoState38Action,
	{_State72, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State72, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State79, IntegerLiteralToken}:                     _GotoState21Action,
	{_State79, FloatLiteralToken}:                       _GotoState18Action,
	{_State79, RuneLiteralToken}:                        _GotoState27Action,
	{_State79, StringLiteralToken}:                      _GotoState28Action,
	{_State79, IdentifierToken}:                         _GotoState20Action,
	{_State79, TrueToken}:                               _GotoState31Action,
	{_State79, FalseToken}:                              _GotoState17Action,
	{_State79, StructToken}:                             _GotoState29Action,
	{_State79, FuncToken}:                               _GotoState19Action,
	{_State79, LparenToken}:                             _GotoState24Action,
	{_State79, LabelDeclToken}:                          _GotoState22Action,
	{_State79, NotToken}:                                _GotoState26Action,
	{_State79, SubToken}:                                _GotoState30Action,
	{_State79, MulToken}:                                _GotoState25Action,
	{_State79, BitNegToken}:                             _GotoState16Action,
	{_State79, BitAndToken}:                             _GotoState15Action,
	{_State79, LexErrorToken}:                           _GotoState23Action,
	{_State79, LiteralType}:                             _GotoState43Action,
	{_State79, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State79, AnonymousStructExprType}:                 _GotoState36Action,
	{_State79, AtomExprType}:                            _GotoState37Action,
	{_State79, CallExprType}:                            _GotoState39Action,
	{_State79, AccessExprType}:                          _GotoState32Action,
	{_State79, UnaryOpType}:                             _GotoState49Action,
	{_State79, UnaryExprType}:                           _GotoState48Action,
	{_State79, MulExprType}:                             _GotoState44Action,
	{_State79, AddExprType}:                             _GotoState136Action,
	{_State79, OptionalLabelDeclType}:                   _GotoState97Action,
	{_State79, BlockExprType}:                           _GotoState38Action,
	{_State79, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State79, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State81, ArgumentsType}:                           _GotoState137Action,
	{_State88, IntegerLiteralToken}:                     _GotoState21Action,
	{_State88, FloatLiteralToken}:                       _GotoState18Action,
	{_State88, RuneLiteralToken}:                        _GotoState27Action,
	{_State88, StringLiteralToken}:                      _GotoState28Action,
	{_State88, IdentifierToken}:                         _GotoState20Action,
	{_State88, TrueToken}:                               _GotoState31Action,
	{_State88, FalseToken}:                              _GotoState17Action,
	{_State88, StructToken}:                             _GotoState29Action,
	{_State88, FuncToken}:                               _GotoState19Action,
	{_State88, LparenToken}:                             _GotoState24Action,
	{_State88, LabelDeclToken}:                          _GotoState22Action,
	{_State88, NotToken}:                                _GotoState26Action,
	{_State88, SubToken}:                                _GotoState30Action,
	{_State88, MulToken}:                                _GotoState25Action,
	{_State88, BitNegToken}:                             _GotoState16Action,
	{_State88, BitAndToken}:                             _GotoState15Action,
	{_State88, LexErrorToken}:                           _GotoState23Action,
	{_State88, LiteralType}:                             _GotoState43Action,
	{_State88, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State88, AnonymousStructExprType}:                 _GotoState36Action,
	{_State88, AtomExprType}:                            _GotoState37Action,
	{_State88, CallExprType}:                            _GotoState39Action,
	{_State88, AccessExprType}:                          _GotoState32Action,
	{_State88, UnaryOpType}:                             _GotoState49Action,
	{_State88, UnaryExprType}:                           _GotoState138Action,
	{_State88, OptionalLabelDeclType}:                   _GotoState97Action,
	{_State88, BlockExprType}:                           _GotoState38Action,
	{_State88, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State88, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State89, IntegerLiteralToken}:                     _GotoState21Action,
	{_State89, FloatLiteralToken}:                       _GotoState18Action,
	{_State89, RuneLiteralToken}:                        _GotoState27Action,
	{_State89, StringLiteralToken}:                      _GotoState28Action,
	{_State89, IdentifierToken}:                         _GotoState20Action,
	{_State89, TrueToken}:                               _GotoState31Action,
	{_State89, FalseToken}:                              _GotoState17Action,
	{_State89, StructToken}:                             _GotoState29Action,
	{_State89, FuncToken}:                               _GotoState19Action,
	{_State89, LparenToken}:                             _GotoState24Action,
	{_State89, LabelDeclToken}:                          _GotoState22Action,
	{_State89, NotToken}:                                _GotoState26Action,
	{_State89, SubToken}:                                _GotoState30Action,
	{_State89, MulToken}:                                _GotoState25Action,
	{_State89, BitNegToken}:                             _GotoState16Action,
	{_State89, BitAndToken}:                             _GotoState15Action,
	{_State89, LexErrorToken}:                           _GotoState23Action,
	{_State89, LiteralType}:                             _GotoState43Action,
	{_State89, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State89, AnonymousStructExprType}:                 _GotoState36Action,
	{_State89, AtomExprType}:                            _GotoState37Action,
	{_State89, CallExprType}:                            _GotoState39Action,
	{_State89, AccessExprType}:                          _GotoState32Action,
	{_State89, UnaryOpType}:                             _GotoState49Action,
	{_State89, UnaryExprType}:                           _GotoState48Action,
	{_State89, MulExprType}:                             _GotoState44Action,
	{_State89, AddExprType}:                             _GotoState33Action,
	{_State89, CmpExprType}:                             _GotoState40Action,
	{_State89, AndExprType}:                             _GotoState34Action,
	{_State89, OrExprType}:                              _GotoState46Action,
	{_State89, SequenceExprType}:                        _GotoState140Action,
	{_State89, OptionalLabelDeclType}:                   _GotoState97Action,
	{_State89, BlockExprType}:                           _GotoState139Action,
	{_State89, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State89, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State90, IntegerLiteralToken}:                     _GotoState21Action,
	{_State90, FloatLiteralToken}:                       _GotoState18Action,
	{_State90, RuneLiteralToken}:                        _GotoState27Action,
	{_State90, StringLiteralToken}:                      _GotoState28Action,
	{_State90, IdentifierToken}:                         _GotoState20Action,
	{_State90, TrueToken}:                               _GotoState31Action,
	{_State90, FalseToken}:                              _GotoState17Action,
	{_State90, StructToken}:                             _GotoState29Action,
	{_State90, FuncToken}:                               _GotoState19Action,
	{_State90, LparenToken}:                             _GotoState24Action,
	{_State90, LabelDeclToken}:                          _GotoState22Action,
	{_State90, NotToken}:                                _GotoState26Action,
	{_State90, SubToken}:                                _GotoState30Action,
	{_State90, MulToken}:                                _GotoState25Action,
	{_State90, BitNegToken}:                             _GotoState16Action,
	{_State90, BitAndToken}:                             _GotoState15Action,
	{_State90, LexErrorToken}:                           _GotoState23Action,
	{_State90, LiteralType}:                             _GotoState43Action,
	{_State90, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State90, AnonymousStructExprType}:                 _GotoState36Action,
	{_State90, AtomExprType}:                            _GotoState37Action,
	{_State90, CallExprType}:                            _GotoState39Action,
	{_State90, AccessExprType}:                          _GotoState32Action,
	{_State90, UnaryOpType}:                             _GotoState49Action,
	{_State90, UnaryExprType}:                           _GotoState48Action,
	{_State90, MulExprType}:                             _GotoState44Action,
	{_State90, AddExprType}:                             _GotoState33Action,
	{_State90, CmpExprType}:                             _GotoState40Action,
	{_State90, AndExprType}:                             _GotoState34Action,
	{_State90, OrExprType}:                              _GotoState46Action,
	{_State90, SequenceExprType}:                        _GotoState141Action,
	{_State90, OptionalLabelDeclType}:                   _GotoState97Action,
	{_State90, BlockExprType}:                           _GotoState38Action,
	{_State90, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State90, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State91, CaseToken}:                               _GotoState142Action,
	{_State96, IntegerLiteralToken}:                     _GotoState21Action,
	{_State96, FloatLiteralToken}:                       _GotoState18Action,
	{_State96, RuneLiteralToken}:                        _GotoState27Action,
	{_State96, StringLiteralToken}:                      _GotoState28Action,
	{_State96, IdentifierToken}:                         _GotoState20Action,
	{_State96, TrueToken}:                               _GotoState31Action,
	{_State96, FalseToken}:                              _GotoState17Action,
	{_State96, StructToken}:                             _GotoState29Action,
	{_State96, FuncToken}:                               _GotoState19Action,
	{_State96, LparenToken}:                             _GotoState24Action,
	{_State96, LabelDeclToken}:                          _GotoState22Action,
	{_State96, NotToken}:                                _GotoState26Action,
	{_State96, SubToken}:                                _GotoState30Action,
	{_State96, MulToken}:                                _GotoState25Action,
	{_State96, BitNegToken}:                             _GotoState16Action,
	{_State96, BitAndToken}:                             _GotoState15Action,
	{_State96, LexErrorToken}:                           _GotoState23Action,
	{_State96, LiteralType}:                             _GotoState43Action,
	{_State96, AnonymousFuncExprType}:                   _GotoState35Action,
	{_State96, AnonymousStructExprType}:                 _GotoState36Action,
	{_State96, AtomExprType}:                            _GotoState37Action,
	{_State96, CallExprType}:                            _GotoState39Action,
	{_State96, AccessExprType}:                          _GotoState32Action,
	{_State96, UnaryOpType}:                             _GotoState49Action,
	{_State96, UnaryExprType}:                           _GotoState48Action,
	{_State96, MulExprType}:                             _GotoState44Action,
	{_State96, AddExprType}:                             _GotoState33Action,
	{_State96, CmpExprType}:                             _GotoState40Action,
	{_State96, AndExprType}:                             _GotoState143Action,
	{_State96, OptionalLabelDeclType}:                   _GotoState97Action,
	{_State96, BlockExprType}:                           _GotoState38Action,
	{_State96, ExplicitStructTypeType}:                  _GotoState42Action,
	{_State96, ExplicitFuncTypeType}:                    _GotoState41Action,
	{_State97, LbraceToken}:                             _GotoState56Action,
	{_State97, BlockBodyType}:                           _GotoState92Action,
	{_State99, PackageStatementsType}:                   _GotoState144Action,
	{_State100, GenericParametersType}:                  _GotoState145Action,
	{_State101, IdentifierToken}:                        _GotoState146Action,
	{_State101, StructToken}:                            _GotoState29Action,
	{_State101, EnumToken}:                              _GotoState103Action,
	{_State101, TraitToken}:                             _GotoState107Action,
	{_State101, LparenToken}:                            _GotoState61Action,
	{_State101, ExclaimToken}:                           _GotoState104Action,
	{_State101, ExclaimExclaimToken}:                    _GotoState105Action,
	{_State101, AtomTypeType}:                           _GotoState108Action,
	{_State101, TraitableTypeType}:                      _GotoState117Action,
	{_State101, TraitIntersectTypeType}:                 _GotoState114Action,
	{_State101, TraitUnionTypeType}:                     _GotoState116Action,
	{_State101, ImplicitStructTypeType}:                 _GotoState112Action,
	{_State101, ExplicitStructTypeType}:                 _GotoState110Action,
	{_State101, TraitTypeType}:                          _GotoState115Action,
	{_State101, ExplicitEnumTypeType}:                   _GotoState109Action,
	{_State101, ValueTypeType}:                          _GotoState147Action,
	{_State102, IdentifierToken}:                        _GotoState146Action,
	{_State102, StructToken}:                            _GotoState29Action,
	{_State102, EnumToken}:                              _GotoState103Action,
	{_State102, TraitToken}:                             _GotoState107Action,
	{_State102, LparenToken}:                            _GotoState61Action,
	{_State102, ExclaimToken}:                           _GotoState104Action,
	{_State102, ExclaimExclaimToken}:                    _GotoState105Action,
	{_State102, AtomTypeType}:                           _GotoState108Action,
	{_State102, TraitableTypeType}:                      _GotoState117Action,
	{_State102, TraitIntersectTypeType}:                 _GotoState114Action,
	{_State102, TraitUnionTypeType}:                     _GotoState116Action,
	{_State102, ImplicitStructTypeType}:                 _GotoState112Action,
	{_State102, ExplicitStructTypeType}:                 _GotoState110Action,
	{_State102, TraitTypeType}:                          _GotoState115Action,
	{_State102, ExplicitEnumTypeType}:                   _GotoState109Action,
	{_State102, ValueTypeType}:                          _GotoState148Action,
	{_State103, LparenToken}:                            _GotoState149Action,
	{_State104, IdentifierToken}:                        _GotoState146Action,
	{_State104, StructToken}:                            _GotoState29Action,
	{_State104, EnumToken}:                              _GotoState103Action,
	{_State104, TraitToken}:                             _GotoState107Action,
	{_State104, LparenToken}:                            _GotoState61Action,
	{_State104, AtomTypeType}:                           _GotoState150Action,
	{_State104, ImplicitStructTypeType}:                 _GotoState112Action,
	{_State104, ExplicitStructTypeType}:                 _GotoState110Action,
	{_State104, TraitTypeType}:                          _GotoState115Action,
	{_State104, ExplicitEnumTypeType}:                   _GotoState109Action,
	{_State105, IdentifierToken}:                        _GotoState146Action,
	{_State105, StructToken}:                            _GotoState29Action,
	{_State105, EnumToken}:                              _GotoState103Action,
	{_State105, TraitToken}:                             _GotoState107Action,
	{_State105, LparenToken}:                            _GotoState61Action,
	{_State105, AtomTypeType}:                           _GotoState151Action,
	{_State105, ImplicitStructTypeType}:                 _GotoState112Action,
	{_State105, ExplicitStructTypeType}:                 _GotoState110Action,
	{_State105, TraitTypeType}:                          _GotoState115Action,
	{_State105, ExplicitEnumTypeType}:                   _GotoState109Action,
	{_State106, IdentifierToken}:                        _GotoState146Action,
	{_State106, StructToken}:                            _GotoState29Action,
	{_State106, EnumToken}:                              _GotoState103Action,
	{_State106, TraitToken}:                             _GotoState107Action,
	{_State106, LparenToken}:                            _GotoState61Action,
	{_State106, QuestionToken}:                          _GotoState152Action,
	{_State106, DollarLbraceToken}:                      _GotoState100Action,
	{_State106, ExclaimToken}:                           _GotoState104Action,
	{_State106, ExclaimExclaimToken}:                    _GotoState105Action,
	{_State106, OptionalGenericParametersType}:          _GotoState153Action,
	{_State106, AtomTypeType}:                           _GotoState108Action,
	{_State106, TraitableTypeType}:                      _GotoState117Action,
	{_State106, TraitIntersectTypeType}:                 _GotoState114Action,
	{_State106, TraitUnionTypeType}:                     _GotoState116Action,
	{_State106, ImplicitStructTypeType}:                 _GotoState112Action,
	{_State106, ExplicitStructTypeType}:                 _GotoState110Action,
	{_State106, TraitTypeType}:                          _GotoState115Action,
	{_State106, ExplicitEnumTypeType}:                   _GotoState109Action,
	{_State106, ValueTypeType}:                          _GotoState154Action,
	{_State107, LparenToken}:                            _GotoState61Action,
	{_State107, ImplicitStructTypeType}:                 _GotoState155Action,
	{_State113, RparenToken}:                            _GotoState156Action,
	{_State114, SubToken}:                               _GotoState158Action,
	{_State114, BitAndToken}:                            _GotoState157Action,
	{_State116, BitOrToken}:                             _GotoState159Action,
	{_State118, OrToken}:                                _GotoState160Action,
	{_State119, DollarLbraceToken}:                      _GotoState100Action,
	{_State119, OptionalGenericParametersType}:          _GotoState161Action,
	{_State120, IntegerLiteralToken}:                    _GotoState21Action,
	{_State120, FloatLiteralToken}:                      _GotoState18Action,
	{_State120, RuneLiteralToken}:                       _GotoState27Action,
	{_State120, StringLiteralToken}:                     _GotoState28Action,
	{_State120, IdentifierToken}:                        _GotoState20Action,
	{_State120, TrueToken}:                              _GotoState31Action,
	{_State120, FalseToken}:                             _GotoState17Action,
	{_State120, ReturnToken}:                            _GotoState166Action,
	{_State120, BreakToken}:                             _GotoState163Action,
	{_State120, ContinueToken}:                          _GotoState164Action,
	{_State120, UnsafeToken}:                            _GotoState167Action,
	{_State120, StructToken}:                            _GotoState29Action,
	{_State120, FuncToken}:                              _GotoState19Action,
	{_State120, AsyncToken}:                             _GotoState162Action,
	{_State120, RbraceToken}:                            _GotoState165Action,
	{_State120, LparenToken}:                            _GotoState24Action,
	{_State120, LabelDeclToken}:                         _GotoState22Action,
	{_State120, NotToken}:                               _GotoState26Action,
	{_State120, SubToken}:                               _GotoState30Action,
	{_State120, MulToken}:                               _GotoState25Action,
	{_State120, BitNegToken}:                            _GotoState16Action,
	{_State120, BitAndToken}:                            _GotoState15Action,
	{_State120, LexErrorToken}:                          _GotoState23Action,
	{_State120, LiteralType}:                            _GotoState43Action,
	{_State120, AnonymousFuncExprType}:                  _GotoState35Action,
	{_State120, AnonymousStructExprType}:                _GotoState36Action,
	{_State120, AtomExprType}:                           _GotoState37Action,
	{_State120, CallExprType}:                           _GotoState39Action,
	{_State120, AccessExprType}:                         _GotoState168Action,
	{_State120, UnaryOpType}:                            _GotoState49Action,
	{_State120, UnaryExprType}:                          _GotoState48Action,
	{_State120, MulExprType}:                            _GotoState44Action,
	{_State120, AddExprType}:                            _GotoState33Action,
	{_State120, CmpExprType}:                            _GotoState40Action,
	{_State120, AndExprType}:                            _GotoState34Action,
	{_State120, OrExprType}:                             _GotoState46Action,
	{_State120, SequenceExprType}:                       _GotoState47Action,
	{_State120, JumpTypeType}:                           _GotoState171Action,
	{_State120, ExpressionOrImplicitStructType}:         _GotoState170Action,
	{_State120, UnsafeStatementType}:                    _GotoState174Action,
	{_State120, StatementBodyType}:                      _GotoState173Action,
	{_State120, StatementType}:                          _GotoState172Action,
	{_State120, OptionalLabelDeclType}:                  _GotoState45Action,
	{_State120, BlockExprType}:                          _GotoState38Action,
	{_State120, ExpressionType}:                         _GotoState169Action,
	{_State120, ExplicitStructTypeType}:                 _GotoState42Action,
	{_State120, ExplicitFuncTypeType}:                   _GotoState41Action,
	{_State121, IdentifierToken}:                        _GotoState146Action,
	{_State121, StructToken}:                            _GotoState29Action,
	{_State121, EnumToken}:                              _GotoState103Action,
	{_State121, TraitToken}:                             _GotoState107Action,
	{_State121, LparenToken}:                            _GotoState61Action,
	{_State121, ExclaimToken}:                           _GotoState104Action,
	{_State121, ExclaimExclaimToken}:                    _GotoState105Action,
	{_State121, AtomTypeType}:                           _GotoState108Action,
	{_State121, TraitableTypeType}:                      _GotoState117Action,
	{_State121, TraitIntersectTypeType}:                 _GotoState114Action,
	{_State121, TraitUnionTypeType}:                     _GotoState116Action,
	{_State121, ImplicitStructTypeType}:                 _GotoState112Action,
	{_State121, ExplicitStructTypeType}:                 _GotoState110Action,
	{_State121, TraitTypeType}:                          _GotoState115Action,
	{_State121, ExplicitEnumTypeType}:                   _GotoState109Action,
	{_State121, ValueTypeType}:                          _GotoState175Action,
	{_State122, IdentifierToken}:                        _GotoState146Action,
	{_State122, StructToken}:                            _GotoState29Action,
	{_State122, EnumToken}:                              _GotoState103Action,
	{_State122, TraitToken}:                             _GotoState107Action,
	{_State122, LparenToken}:                            _GotoState61Action,
	{_State122, QuestionToken}:                          _GotoState152Action,
	{_State122, DollarLbraceToken}:                      _GotoState100Action,
	{_State122, DotdotdotToken}:                         _GotoState176Action,
	{_State122, ExclaimToken}:                           _GotoState104Action,
	{_State122, ExclaimExclaimToken}:                    _GotoState105Action,
	{_State122, OptionalGenericParametersType}:          _GotoState153Action,
	{_State122, AtomTypeType}:                           _GotoState108Action,
	{_State122, TraitableTypeType}:                      _GotoState117Action,
	{_State122, TraitIntersectTypeType}:                 _GotoState114Action,
	{_State122, TraitUnionTypeType}:                     _GotoState116Action,
	{_State122, ImplicitStructTypeType}:                 _GotoState112Action,
	{_State122, ExplicitStructTypeType}:                 _GotoState110Action,
	{_State122, TraitTypeType}:                          _GotoState115Action,
	{_State122, ExplicitEnumTypeType}:                   _GotoState109Action,
	{_State122, ValueTypeType}:                          _GotoState154Action,
	{_State123, CommaToken}:                             _GotoState177Action,
	{_State126, RparenToken}:                            _GotoState178Action,
	{_State127, CommaToken}:                             _GotoState179Action,
	{_State130, RbraceToken}:                            _GotoState180Action,
	{_State132, RbracketToken}:                          _GotoState181Action,
	{_State133, RparenToken}:                            _GotoState182Action,
	{_State134, MulToken}:                               _GotoState87Action,
	{_State134, DivToken}:                               _GotoState85Action,
	{_State134, ModToken}:                               _GotoState86Action,
	{_State134, BitAndToken}:                            _GotoState82Action,
	{_State134, BitLshiftToken}:                         _GotoState83Action,
	{_State134, BitRshiftToken}:                         _GotoState84Action,
	{_State134, MulOpType}:                              _GotoState88Action,
	{_State135, EqualToken}:                             _GotoState73Action,
	{_State135, NotEqualToken}:                          _GotoState78Action,
	{_State135, LessToken}:                              _GotoState76Action,
	{_State135, LessOrEqualToken}:                       _GotoState77Action,
	{_State135, GreaterToken}:                           _GotoState74Action,
	{_State135, GreaterOrEqualToken}:                    _GotoState75Action,
	{_State135, CmpOpType}:                              _GotoState79Action,
	{_State136, AddToken}:                               _GotoState67Action,
	{_State136, SubToken}:                               _GotoState70Action,
	{_State136, BitXorToken}:                            _GotoState69Action,
	{_State136, BitOrToken}:                             _GotoState68Action,
	{_State136, AddOpType}:                              _GotoState71Action,
	{_State137, RparenToken}:                            _GotoState183Action,
	{_State140, LabelDeclToken}:                         _GotoState22Action,
	{_State140, OptionalLabelDeclType}:                  _GotoState97Action,
	{_State140, BlockExprType}:                          _GotoState184Action,
	{_State141, LbraceToken}:                            _GotoState56Action,
	{_State141, BlockBodyType}:                          _GotoState185Action,
	{_State142, DefaultToken}:                           _GotoState186Action,
	{_State143, AndToken}:                               _GotoState72Action,
	{_State144, UnsafeToken}:                            _GotoState167Action,
	{_State144, RparenToken}:                            _GotoState187Action,
	{_State144, UnsafeStatementType}:                    _GotoState190Action,
	{_State144, PackageStatementBodyType}:               _GotoState189Action,
	{_State144, PackageStatementType}:                   _GotoState188Action,
	{_State145, RbraceToken}:                            _GotoState191Action,
	{_State146, DollarLbraceToken}:                      _GotoState100Action,
	{_State146, OptionalGenericParametersType}:          _GotoState153Action,
	{_State147, OrToken}:                                _GotoState160Action,
	{_State148, OrToken}:                                _GotoState160Action,
	{_State149, RparenToken}:                            _GotoState192Action,
	{_State154, OrToken}:                                _GotoState160Action,
	{_State157, IdentifierToken}:                        _GotoState146Action,
	{_State157, StructToken}:                            _GotoState29Action,
	{_State157, EnumToken}:                              _GotoState103Action,
	{_State157, TraitToken}:                             _GotoState107Action,
	{_State157, LparenToken}:                            _GotoState61Action,
	{_State157, ExclaimToken}:                           _GotoState104Action,
	{_State157, ExclaimExclaimToken}:                    _GotoState105Action,
	{_State157, AtomTypeType}:                           _GotoState108Action,
	{_State157, TraitableTypeType}:                      _GotoState193Action,
	{_State157, ImplicitStructTypeType}:                 _GotoState112Action,
	{_State157, ExplicitStructTypeType}:                 _GotoState110Action,
	{_State157, TraitTypeType}:                          _GotoState115Action,
	{_State157, ExplicitEnumTypeType}:                   _GotoState109Action,
	{_State158, IdentifierToken}:                        _GotoState146Action,
	{_State158, StructToken}:                            _GotoState29Action,
	{_State158, EnumToken}:                              _GotoState103Action,
	{_State158, TraitToken}:                             _GotoState107Action,
	{_State158, LparenToken}:                            _GotoState61Action,
	{_State158, ExclaimToken}:                           _GotoState104Action,
	{_State158, ExclaimExclaimToken}:                    _GotoState105Action,
	{_State158, AtomTypeType}:                           _GotoState108Action,
	{_State158, TraitableTypeType}:                      _GotoState194Action,
	{_State158, ImplicitStructTypeType}:                 _GotoState112Action,
	{_State158, ExplicitStructTypeType}:                 _GotoState110Action,
	{_State158, TraitTypeType}:                          _GotoState115Action,
	{_State158, ExplicitEnumTypeType}:                   _GotoState109Action,
	{_State159, IdentifierToken}:                        _GotoState146Action,
	{_State159, StructToken}:                            _GotoState29Action,
	{_State159, EnumToken}:                              _GotoState103Action,
	{_State159, TraitToken}:                             _GotoState107Action,
	{_State159, LparenToken}:                            _GotoState61Action,
	{_State159, ExclaimToken}:                           _GotoState104Action,
	{_State159, ExclaimExclaimToken}:                    _GotoState105Action,
	{_State159, AtomTypeType}:                           _GotoState108Action,
	{_State159, TraitableTypeType}:                      _GotoState117Action,
	{_State159, TraitIntersectTypeType}:                 _GotoState195Action,
	{_State159, ImplicitStructTypeType}:                 _GotoState112Action,
	{_State159, ExplicitStructTypeType}:                 _GotoState110Action,
	{_State159, TraitTypeType}:                          _GotoState115Action,
	{_State159, ExplicitEnumTypeType}:                   _GotoState109Action,
	{_State160, IdentifierToken}:                        _GotoState146Action,
	{_State160, StructToken}:                            _GotoState29Action,
	{_State160, EnumToken}:                              _GotoState103Action,
	{_State160, TraitToken}:                             _GotoState107Action,
	{_State160, LparenToken}:                            _GotoState61Action,
	{_State160, ExclaimToken}:                           _GotoState104Action,
	{_State160, ExclaimExclaimToken}:                    _GotoState105Action,
	{_State160, AtomTypeType}:                           _GotoState108Action,
	{_State160, TraitableTypeType}:                      _GotoState117Action,
	{_State160, TraitIntersectTypeType}:                 _GotoState114Action,
	{_State160, TraitUnionTypeType}:                     _GotoState196Action,
	{_State160, ImplicitStructTypeType}:                 _GotoState112Action,
	{_State160, ExplicitStructTypeType}:                 _GotoState110Action,
	{_State160, TraitTypeType}:                          _GotoState115Action,
	{_State160, ExplicitEnumTypeType}:                   _GotoState109Action,
	{_State161, LparenToken}:                            _GotoState58Action,
	{_State161, ImplicitFuncTypeType}:                   _GotoState197Action,
	{_State162, IntegerLiteralToken}:                    _GotoState21Action,
	{_State162, FloatLiteralToken}:                      _GotoState18Action,
	{_State162, RuneLiteralToken}:                       _GotoState27Action,
	{_State162, StringLiteralToken}:                     _GotoState28Action,
	{_State162, IdentifierToken}:                        _GotoState20Action,
	{_State162, TrueToken}:                              _GotoState31Action,
	{_State162, FalseToken}:                             _GotoState17Action,
	{_State162, StructToken}:                            _GotoState29Action,
	{_State162, FuncToken}:                              _GotoState19Action,
	{_State162, LparenToken}:                            _GotoState24Action,
	{_State162, LabelDeclToken}:                         _GotoState22Action,
	{_State162, LexErrorToken}:                          _GotoState23Action,
	{_State162, LiteralType}:                            _GotoState43Action,
	{_State162, AnonymousFuncExprType}:                  _GotoState35Action,
	{_State162, AnonymousStructExprType}:                _GotoState36Action,
	{_State162, AtomExprType}:                           _GotoState37Action,
	{_State162, CallExprType}:                           _GotoState199Action,
	{_State162, AccessExprType}:                         _GotoState198Action,
	{_State162, OptionalLabelDeclType}:                  _GotoState97Action,
	{_State162, BlockExprType}:                          _GotoState38Action,
	{_State162, ExplicitStructTypeType}:                 _GotoState42Action,
	{_State162, ExplicitFuncTypeType}:                   _GotoState41Action,
	{_State167, LessToken}:                              _GotoState200Action,
	{_State168, LparenToken}:                            _GotoState66Action,
	{_State168, LbracketToken}:                          _GotoState65Action,
	{_State168, DotToken}:                               _GotoState64Action,
	{_State168, DollarLbraceToken}:                      _GotoState63Action,
	{_State168, AddAssignToken}:                         _GotoState201Action,
	{_State168, SubAssignToken}:                         _GotoState212Action,
	{_State168, MulAssignToken}:                         _GotoState211Action,
	{_State168, DivAssignToken}:                         _GotoState209Action,
	{_State168, ModAssignToken}:                         _GotoState210Action,
	{_State168, AddOneAssignToken}:                      _GotoState202Action,
	{_State168, SubOneAssignToken}:                      _GotoState213Action,
	{_State168, BitNegAssignToken}:                      _GotoState205Action,
	{_State168, BitAndAssignToken}:                      _GotoState203Action,
	{_State168, BitOrAssignToken}:                       _GotoState206Action,
	{_State168, BitXorAssignToken}:                      _GotoState208Action,
	{_State168, BitLshiftAssignToken}:                   _GotoState204Action,
	{_State168, BitRshiftAssignToken}:                   _GotoState207Action,
	{_State168, OpOneAssignType}:                        _GotoState215Action,
	{_State168, BinaryOpAssignType}:                     _GotoState214Action,
	{_State170, CommaToken}:                             _GotoState216Action,
	{_State171, JumpLabelToken}:                         _GotoState217Action,
	{_State171, OptionalJumpLabelType}:                  _GotoState218Action,
	{_State173, NewlinesToken}:                          _GotoState219Action,
	{_State173, SemicolonToken}:                         _GotoState220Action,
	{_State175, OrToken}:                                _GotoState160Action,
	{_State176, IdentifierToken}:                        _GotoState146Action,
	{_State176, StructToken}:                            _GotoState29Action,
	{_State176, EnumToken}:                              _GotoState103Action,
	{_State176, TraitToken}:                             _GotoState107Action,
	{_State176, LparenToken}:                            _GotoState61Action,
	{_State176, QuestionToken}:                          _GotoState221Action,
	{_State176, ExclaimToken}:                           _GotoState104Action,
	{_State176, ExclaimExclaimToken}:                    _GotoState105Action,
	{_State176, AtomTypeType}:                           _GotoState108Action,
	{_State176, TraitableTypeType}:                      _GotoState117Action,
	{_State176, TraitIntersectTypeType}:                 _GotoState114Action,
	{_State176, TraitUnionTypeType}:                     _GotoState116Action,
	{_State176, ImplicitStructTypeType}:                 _GotoState112Action,
	{_State176, ExplicitStructTypeType}:                 _GotoState110Action,
	{_State176, TraitTypeType}:                          _GotoState115Action,
	{_State176, ExplicitEnumTypeType}:                   _GotoState109Action,
	{_State176, ValueTypeType}:                          _GotoState222Action,
	{_State177, IdentifierToken}:                        _GotoState122Action,
	{_State177, StructToken}:                            _GotoState29Action,
	{_State177, EnumToken}:                              _GotoState103Action,
	{_State177, TraitToken}:                             _GotoState107Action,
	{_State177, LparenToken}:                            _GotoState61Action,
	{_State177, DotdotdotToken}:                         _GotoState121Action,
	{_State177, ExclaimToken}:                           _GotoState104Action,
	{_State177, ExclaimExclaimToken}:                    _GotoState105Action,
	{_State177, AtomTypeType}:                           _GotoState108Action,
	{_State177, TraitableTypeType}:                      _GotoState117Action,
	{_State177, TraitIntersectTypeType}:                 _GotoState114Action,
	{_State177, TraitUnionTypeType}:                     _GotoState116Action,
	{_State177, ImplicitStructTypeType}:                 _GotoState112Action,
	{_State177, ExplicitStructTypeType}:                 _GotoState110Action,
	{_State177, TraitTypeType}:                          _GotoState115Action,
	{_State177, ExplicitEnumTypeType}:                   _GotoState109Action,
	{_State177, ValueTypeType}:                          _GotoState118Action,
	{_State177, FieldDeclType}:                          _GotoState111Action,
	{_State177, ParameterDeclType}:                      _GotoState224Action,
	{_State177, VarargType}:                             _GotoState127Action,
	{_State177, OptionalVarargType}:                     _GotoState223Action,
	{_State178, IdentifierToken}:                        _GotoState146Action,
	{_State178, StructToken}:                            _GotoState29Action,
	{_State178, EnumToken}:                              _GotoState103Action,
	{_State178, TraitToken}:                             _GotoState107Action,
	{_State178, LparenToken}:                            _GotoState61Action,
	{_State178, QuestionToken}:                          _GotoState225Action,
	{_State178, ExclaimToken}:                           _GotoState104Action,
	{_State178, ExclaimExclaimToken}:                    _GotoState105Action,
	{_State178, AtomTypeType}:                           _GotoState108Action,
	{_State178, TraitableTypeType}:                      _GotoState117Action,
	{_State178, TraitIntersectTypeType}:                 _GotoState114Action,
	{_State178, TraitUnionTypeType}:                     _GotoState116Action,
	{_State178, ImplicitStructTypeType}:                 _GotoState112Action,
	{_State178, ExplicitStructTypeType}:                 _GotoState110Action,
	{_State178, TraitTypeType}:                          _GotoState115Action,
	{_State178, ExplicitEnumTypeType}:                   _GotoState109Action,
	{_State178, ValueTypeType}:                          _GotoState226Action,
	{_State180, LparenToken}:                            _GotoState227Action,
	{_State185, ElseToken}:                              _GotoState228Action,
	{_State189, NewlinesToken}:                          _GotoState229Action,
	{_State189, SemicolonToken}:                         _GotoState230Action,
	{_State195, SubToken}:                               _GotoState158Action,
	{_State195, BitAndToken}:                            _GotoState157Action,
	{_State196, BitOrToken}:                             _GotoState159Action,
	{_State198, LparenToken}:                            _GotoState66Action,
	{_State198, LbracketToken}:                          _GotoState65Action,
	{_State198, DotToken}:                               _GotoState64Action,
	{_State198, DollarLbraceToken}:                      _GotoState63Action,
	{_State200, IdentifierToken}:                        _GotoState231Action,
	{_State214, IntegerLiteralToken}:                    _GotoState21Action,
	{_State214, FloatLiteralToken}:                      _GotoState18Action,
	{_State214, RuneLiteralToken}:                       _GotoState27Action,
	{_State214, StringLiteralToken}:                     _GotoState28Action,
	{_State214, IdentifierToken}:                        _GotoState20Action,
	{_State214, TrueToken}:                              _GotoState31Action,
	{_State214, FalseToken}:                             _GotoState17Action,
	{_State214, StructToken}:                            _GotoState29Action,
	{_State214, FuncToken}:                              _GotoState19Action,
	{_State214, LparenToken}:                            _GotoState24Action,
	{_State214, LabelDeclToken}:                         _GotoState22Action,
	{_State214, NotToken}:                               _GotoState26Action,
	{_State214, SubToken}:                               _GotoState30Action,
	{_State214, MulToken}:                               _GotoState25Action,
	{_State214, BitNegToken}:                            _GotoState16Action,
	{_State214, BitAndToken}:                            _GotoState15Action,
	{_State214, LexErrorToken}:                          _GotoState23Action,
	{_State214, LiteralType}:                            _GotoState43Action,
	{_State214, AnonymousFuncExprType}:                  _GotoState35Action,
	{_State214, AnonymousStructExprType}:                _GotoState36Action,
	{_State214, AtomExprType}:                           _GotoState37Action,
	{_State214, CallExprType}:                           _GotoState39Action,
	{_State214, AccessExprType}:                         _GotoState32Action,
	{_State214, UnaryOpType}:                            _GotoState49Action,
	{_State214, UnaryExprType}:                          _GotoState48Action,
	{_State214, MulExprType}:                            _GotoState44Action,
	{_State214, AddExprType}:                            _GotoState33Action,
	{_State214, CmpExprType}:                            _GotoState40Action,
	{_State214, AndExprType}:                            _GotoState34Action,
	{_State214, OrExprType}:                             _GotoState46Action,
	{_State214, SequenceExprType}:                       _GotoState47Action,
	{_State214, OptionalLabelDeclType}:                  _GotoState45Action,
	{_State214, BlockExprType}:                          _GotoState38Action,
	{_State214, ExpressionType}:                         _GotoState232Action,
	{_State214, ExplicitStructTypeType}:                 _GotoState42Action,
	{_State214, ExplicitFuncTypeType}:                   _GotoState41Action,
	{_State216, IntegerLiteralToken}:                    _GotoState21Action,
	{_State216, FloatLiteralToken}:                      _GotoState18Action,
	{_State216, RuneLiteralToken}:                       _GotoState27Action,
	{_State216, StringLiteralToken}:                     _GotoState28Action,
	{_State216, IdentifierToken}:                        _GotoState20Action,
	{_State216, TrueToken}:                              _GotoState31Action,
	{_State216, FalseToken}:                             _GotoState17Action,
	{_State216, StructToken}:                            _GotoState29Action,
	{_State216, FuncToken}:                              _GotoState19Action,
	{_State216, LparenToken}:                            _GotoState24Action,
	{_State216, LabelDeclToken}:                         _GotoState22Action,
	{_State216, NotToken}:                               _GotoState26Action,
	{_State216, SubToken}:                               _GotoState30Action,
	{_State216, MulToken}:                               _GotoState25Action,
	{_State216, BitNegToken}:                            _GotoState16Action,
	{_State216, BitAndToken}:                            _GotoState15Action,
	{_State216, LexErrorToken}:                          _GotoState23Action,
	{_State216, LiteralType}:                            _GotoState43Action,
	{_State216, AnonymousFuncExprType}:                  _GotoState35Action,
	{_State216, AnonymousStructExprType}:                _GotoState36Action,
	{_State216, AtomExprType}:                           _GotoState37Action,
	{_State216, CallExprType}:                           _GotoState39Action,
	{_State216, AccessExprType}:                         _GotoState32Action,
	{_State216, UnaryOpType}:                            _GotoState49Action,
	{_State216, UnaryExprType}:                          _GotoState48Action,
	{_State216, MulExprType}:                            _GotoState44Action,
	{_State216, AddExprType}:                            _GotoState33Action,
	{_State216, CmpExprType}:                            _GotoState40Action,
	{_State216, AndExprType}:                            _GotoState34Action,
	{_State216, OrExprType}:                             _GotoState46Action,
	{_State216, SequenceExprType}:                       _GotoState47Action,
	{_State216, OptionalLabelDeclType}:                  _GotoState45Action,
	{_State216, BlockExprType}:                          _GotoState38Action,
	{_State216, ExpressionType}:                         _GotoState233Action,
	{_State216, ExplicitStructTypeType}:                 _GotoState42Action,
	{_State216, ExplicitFuncTypeType}:                   _GotoState41Action,
	{_State218, IntegerLiteralToken}:                    _GotoState21Action,
	{_State218, FloatLiteralToken}:                      _GotoState18Action,
	{_State218, RuneLiteralToken}:                       _GotoState27Action,
	{_State218, StringLiteralToken}:                     _GotoState28Action,
	{_State218, IdentifierToken}:                        _GotoState20Action,
	{_State218, TrueToken}:                              _GotoState31Action,
	{_State218, FalseToken}:                             _GotoState17Action,
	{_State218, StructToken}:                            _GotoState29Action,
	{_State218, FuncToken}:                              _GotoState19Action,
	{_State218, LparenToken}:                            _GotoState24Action,
	{_State218, LabelDeclToken}:                         _GotoState22Action,
	{_State218, NotToken}:                               _GotoState26Action,
	{_State218, SubToken}:                               _GotoState30Action,
	{_State218, MulToken}:                               _GotoState25Action,
	{_State218, BitNegToken}:                            _GotoState16Action,
	{_State218, BitAndToken}:                            _GotoState15Action,
	{_State218, LexErrorToken}:                          _GotoState23Action,
	{_State218, LiteralType}:                            _GotoState43Action,
	{_State218, AnonymousFuncExprType}:                  _GotoState35Action,
	{_State218, AnonymousStructExprType}:                _GotoState36Action,
	{_State218, AtomExprType}:                           _GotoState37Action,
	{_State218, CallExprType}:                           _GotoState39Action,
	{_State218, AccessExprType}:                         _GotoState32Action,
	{_State218, UnaryOpType}:                            _GotoState49Action,
	{_State218, UnaryExprType}:                          _GotoState48Action,
	{_State218, MulExprType}:                            _GotoState44Action,
	{_State218, AddExprType}:                            _GotoState33Action,
	{_State218, CmpExprType}:                            _GotoState40Action,
	{_State218, AndExprType}:                            _GotoState34Action,
	{_State218, OrExprType}:                             _GotoState46Action,
	{_State218, SequenceExprType}:                       _GotoState47Action,
	{_State218, OptionalExpressionOrImplicitStructType}: _GotoState235Action,
	{_State218, ExpressionOrImplicitStructType}:         _GotoState234Action,
	{_State218, OptionalLabelDeclType}:                  _GotoState45Action,
	{_State218, BlockExprType}:                          _GotoState38Action,
	{_State218, ExpressionType}:                         _GotoState169Action,
	{_State218, ExplicitStructTypeType}:                 _GotoState42Action,
	{_State218, ExplicitFuncTypeType}:                   _GotoState41Action,
	{_State222, OrToken}:                                _GotoState160Action,
	{_State226, OrToken}:                                _GotoState160Action,
	{_State227, ArgumentsType}:                          _GotoState236Action,
	{_State228, IfToken}:                                _GotoState90Action,
	{_State228, LbraceToken}:                            _GotoState56Action,
	{_State228, BlockBodyType}:                          _GotoState237Action,
	{_State228, IfExprType}:                             _GotoState238Action,
	{_State231, GreaterToken}:                           _GotoState239Action,
	{_State234, CommaToken}:                             _GotoState216Action,
	{_State236, RparenToken}:                            _GotoState240Action,
	{_State239, StringLiteralToken}:                     _GotoState241Action,
	{_State4, _WildcardMarker}:                          _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State13, IdentifierToken}:                         _ReduceNilToOptionalReceiverAction,
	{_State15, _WildcardMarker}:                         _ReduceBitAndToUnaryOpAction,
	{_State16, _WildcardMarker}:                         _ReduceBitNegToUnaryOpAction,
	{_State17, _WildcardMarker}:                         _ReduceFalseToLiteralAction,
	{_State18, _WildcardMarker}:                         _ReduceFloatLiteralToLiteralAction,
	{_State20, _WildcardMarker}:                         _ReduceIdentifierToAtomExprAction,
	{_State21, _WildcardMarker}:                         _ReduceIntegerLiteralToLiteralAction,
	{_State22, _WildcardMarker}:                         _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State23, _WildcardMarker}:                         _ReduceLexErrorToAtomExprAction,
	{_State24, RparenToken}:                             _ReduceToArgumentsAction,
	{_State25, _WildcardMarker}:                         _ReduceMulToUnaryOpAction,
	{_State26, _WildcardMarker}:                         _ReduceNotToUnaryOpAction,
	{_State27, _WildcardMarker}:                         _ReduceRuneLiteralToLiteralAction,
	{_State28, _WildcardMarker}:                         _ReduceStringLiteralToLiteralAction,
	{_State30, _WildcardMarker}:                         _ReduceSubToUnaryOpAction,
	{_State31, _WildcardMarker}:                         _ReduceTrueToLiteralAction,
	{_State32, _WildcardMarker}:                         _ReduceAccessExprToUnaryExprAction,
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
	{_State47, _EndMarker}:                              _ReduceSequenceExprToExpressionAction,
	{_State48, _WildcardMarker}:                         _ReduceUnaryExprToMulExprAction,
	{_State49, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State50, _EndMarker}:                              _ReduceCommentToLexInternalTokensAction,
	{_State51, _EndMarker}:                              _ReduceSpacesToLexInternalTokensAction,
	{_State52, _EndMarker}:                              _ReduceNoSpecToPackageDeclAction,
	{_State53, _WildcardMarker}:                         _ReduceNilToOptionalGenericParametersAction,
	{_State56, _WildcardMarker}:                         _ReduceEmptyListToStatementsAction,
	{_State57, _EndMarker}:                              _ReduceToFuncDefAction,
	{_State58, RparenToken}:                             _ReduceNilToOptionalVarargAction,
	{_State59, LbraceToken}:                             _ReduceToExplicitFuncTypeAction,
	{_State62, _WildcardMarker}:                         _ReduceToExplicitStructTypeAction,
	{_State63, RbraceToken}:                             _ReduceToGenericArgumentsAction,
	{_State65, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State66, RparenToken}:                             _ReduceToArgumentsAction,
	{_State67, _WildcardMarker}:                         _ReduceAddToAddOpAction,
	{_State68, _WildcardMarker}:                         _ReduceBitOrToAddOpAction,
	{_State69, _WildcardMarker}:                         _ReduceBitXorToAddOpAction,
	{_State70, _WildcardMarker}:                         _ReduceSubToAddOpAction,
	{_State71, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State72, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State73, _WildcardMarker}:                         _ReduceEqualToCmpOpAction,
	{_State74, _WildcardMarker}:                         _ReduceGreaterToCmpOpAction,
	{_State75, _WildcardMarker}:                         _ReduceGreaterOrEqualToCmpOpAction,
	{_State76, _WildcardMarker}:                         _ReduceLessToCmpOpAction,
	{_State77, _WildcardMarker}:                         _ReduceLessOrEqualToCmpOpAction,
	{_State78, _WildcardMarker}:                         _ReduceNotEqualToCmpOpAction,
	{_State79, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State80, _WildcardMarker}:                         _ReduceToAnonymousFuncExprAction,
	{_State81, RparenToken}:                             _ReduceToArgumentsAction,
	{_State82, _WildcardMarker}:                         _ReduceBitAndToMulOpAction,
	{_State83, _WildcardMarker}:                         _ReduceBitLshiftToMulOpAction,
	{_State84, _WildcardMarker}:                         _ReduceBitRshiftToMulOpAction,
	{_State85, _WildcardMarker}:                         _ReduceDivToMulOpAction,
	{_State86, _WildcardMarker}:                         _ReduceModToMulOpAction,
	{_State87, _WildcardMarker}:                         _ReduceMulToMulOpAction,
	{_State88, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State89, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State90, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State92, _EndMarker}:                              _ReduceToBlockExprAction,
	{_State93, _EndMarker}:                              _ReduceIfExprToExpressionAction,
	{_State94, _EndMarker}:                              _ReduceLoopExprToExpressionAction,
	{_State95, _EndMarker}:                              _ReduceSwitchExprToExpressionAction,
	{_State96, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State98, _WildcardMarker}:                         _ReduceOpToUnaryExprAction,
	{_State99, _WildcardMarker}:                         _ReduceEmptyListToPackageStatementsAction,
	{_State100, RbraceToken}:                            _ReduceToGenericParametersAction,
	{_State106, _WildcardMarker}:                        _ReduceNilToOptionalGenericParametersAction,
	{_State108, _WildcardMarker}:                        _ReduceAtomTypeToTraitableTypeAction,
	{_State109, _WildcardMarker}:                        _ReduceExplicitEnumTypeToAtomTypeAction,
	{_State110, _WildcardMarker}:                        _ReduceExplicitStructTypeToAtomTypeAction,
	{_State111, _WildcardMarker}:                        _ReduceFieldDeclToParameterDeclAction,
	{_State112, _WildcardMarker}:                        _ReduceImplicitStructTypeToAtomTypeAction,
	{_State114, _WildcardMarker}:                        _ReduceTraitIntersectTypeToTraitUnionTypeAction,
	{_State115, _WildcardMarker}:                        _ReduceTraitTypeToAtomTypeAction,
	{_State116, _WildcardMarker}:                        _ReduceTraitUnionTypeToValueTypeAction,
	{_State117, _WildcardMarker}:                        _ReduceTraitableTypeToTraitIntersectTypeAction,
	{_State118, _WildcardMarker}:                        _ReduceImplicitToFieldDeclAction,
	{_State119, LparenToken}:                            _ReduceNilToOptionalGenericParametersAction,
	{_State120, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State122, _WildcardMarker}:                        _ReduceNilToOptionalGenericParametersAction,
	{_State123, RparenToken}:                            _ReduceNonEmptyParametersToParametersAction,
	{_State124, RparenToken}:                            _ReduceOptionalVarargToParametersAction,
	{_State125, _WildcardMarker}:                        _ReduceParameterDeclToNonEmptyParametersAction,
	{_State127, RparenToken}:                            _ReduceVarargToOptionalVarargAction,
	{_State128, _WildcardMarker}:                        _ReduceImplicitToAnonymousStructExprAction,
	{_State129, _WildcardMarker}:                        _ReduceToImplicitStructTypeAction,
	{_State131, _WildcardMarker}:                        _ReduceAccessToAccessExprAction,
	{_State134, _WildcardMarker}:                        _ReduceOpToAddExprAction,
	{_State135, _WildcardMarker}:                        _ReduceOpToAndExprAction,
	{_State136, _WildcardMarker}:                        _ReduceOpToCmpExprAction,
	{_State138, _WildcardMarker}:                        _ReduceOpToMulExprAction,
	{_State139, _WildcardMarker}:                        _ReduceBlockExprToAtomExprAction,
	{_State139, _EndMarker}:                             _ReduceInfiniteToLoopExprAction,
	{_State140, LbraceToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State143, _WildcardMarker}:                        _ReduceOpToOrExprAction,
	{_State146, _WildcardMarker}:                        _ReduceNilToOptionalGenericParametersAction,
	{_State147, _EndMarker}:                             _ReduceAliasToTypeDeclAction,
	{_State148, _EndMarker}:                             _ReduceDefinitionToTypeDeclAction,
	{_State150, _WildcardMarker}:                        _ReduceMethodInterfaceToTraitableTypeAction,
	{_State151, _WildcardMarker}:                        _ReduceTraitToTraitableTypeAction,
	{_State152, _WildcardMarker}:                        _ReduceInferredToParameterDeclAction,
	{_State153, _WildcardMarker}:                        _ReduceNamedToAtomTypeAction,
	{_State154, _WildcardMarker}:                        _ReduceExplicitToFieldDeclAction,
	{_State155, _WildcardMarker}:                        _ReduceToTraitTypeAction,
	{_State156, IdentifierToken}:                        _ReduceReceiverToOptionalReceiverAction,
	{_State162, LbraceToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State163, _WildcardMarker}:                        _ReduceBreakToJumpTypeAction,
	{_State164, _WildcardMarker}:                        _ReduceContinueToJumpTypeAction,
	{_State165, _EndMarker}:                             _ReduceToBlockBodyAction,
	{_State166, _WildcardMarker}:                        _ReduceReturnToJumpTypeAction,
	{_State168, _WildcardMarker}:                        _ReduceAccessExprToUnaryExprAction,
	{_State169, _WildcardMarker}:                        _ReduceExpressionToExpressionOrImplicitStructAction,
	{_State170, _WildcardMarker}:                        _ReduceExpressionOrImplicitStructToStatementBodyAction,
	{_State171, _WildcardMarker}:                        _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State172, _WildcardMarker}:                        _ReduceAddToStatementsAction,
	{_State174, _WildcardMarker}:                        _ReduceUnsafeStatementToStatementBodyAction,
	{_State175, _WildcardMarker}:                        _ReduceImplicitToVarargAction,
	{_State177, RparenToken}:                            _ReduceNilToOptionalVarargAction,
	{_State178, LbraceToken}:                            _ReduceImplicitUnitToImplicitFuncTypeAction,
	{_State179, RparenToken}:                            _ReduceVararg2ToOptionalVarargAction,
	{_State181, _WildcardMarker}:                        _ReduceIndexToAccessExprAction,
	{_State182, _WildcardMarker}:                        _ReduceConcreteToCallExprAction,
	{_State183, _WildcardMarker}:                        _ReduceExplicitToAnonymousStructExprAction,
	{_State184, _EndMarker}:                             _ReduceWhileToLoopExprAction,
	{_State185, _WildcardMarker}:                        _ReduceNoElseToIfExprAction,
	{_State186, _EndMarker}:                             _ReduceToSwitchExprAction,
	{_State187, _EndMarker}:                             _ReduceWithSpecToPackageDeclAction,
	{_State188, _WildcardMarker}:                        _ReduceAddToPackageStatementsAction,
	{_State190, _WildcardMarker}:                        _ReduceToPackageStatementBodyAction,
	{_State191, _WildcardMarker}:                        _ReduceGenericToOptionalGenericParametersAction,
	{_State192, _WildcardMarker}:                        _ReduceToExplicitEnumTypeAction,
	{_State193, _WildcardMarker}:                        _ReduceIntersectToTraitIntersectTypeAction,
	{_State194, _WildcardMarker}:                        _ReduceDifferenceToTraitIntersectTypeAction,
	{_State195, _WildcardMarker}:                        _ReduceUnionToTraitUnionTypeAction,
	{_State196, _WildcardMarker}:                        _ReduceImplicitEnumToValueTypeAction,
	{_State197, LbraceToken}:                            _ReduceToFuncDeclAction,
	{_State199, _WildcardMarker}:                        _ReduceCallExprToAccessExprAction,
	{_State199, NewlinesToken}:                          _ReduceAsyncToStatementBodyAction,
	{_State199, SemicolonToken}:                         _ReduceAsyncToStatementBodyAction,
	{_State201, _WildcardMarker}:                        _ReduceAddAssignToBinaryOpAssignAction,
	{_State202, _WildcardMarker}:                        _ReduceAddOneAssignToOpOneAssignAction,
	{_State203, _WildcardMarker}:                        _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State204, _WildcardMarker}:                        _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State205, _WildcardMarker}:                        _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State206, _WildcardMarker}:                        _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State207, _WildcardMarker}:                        _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State208, _WildcardMarker}:                        _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State209, _WildcardMarker}:                        _ReduceDivAssignToBinaryOpAssignAction,
	{_State210, _WildcardMarker}:                        _ReduceModAssignToBinaryOpAssignAction,
	{_State211, _WildcardMarker}:                        _ReduceMulAssignToBinaryOpAssignAction,
	{_State212, _WildcardMarker}:                        _ReduceSubAssignToBinaryOpAssignAction,
	{_State213, _WildcardMarker}:                        _ReduceSubOneAssignToOpOneAssignAction,
	{_State214, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State215, _WildcardMarker}:                        _ReduceOpOneAssignToStatementBodyAction,
	{_State216, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State217, _WildcardMarker}:                        _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State218, NewlinesToken}:                          _ReduceNilToOptionalExpressionOrImplicitStructAction,
	{_State218, SemicolonToken}:                         _ReduceNilToOptionalExpressionOrImplicitStructAction,
	{_State218, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State219, _WildcardMarker}:                        _ReduceImplicitToStatementAction,
	{_State220, _WildcardMarker}:                        _ReduceExplicitToStatementAction,
	{_State221, _WildcardMarker}:                        _ReduceInferredToVarargAction,
	{_State222, _WildcardMarker}:                        _ReduceExplicitToVarargAction,
	{_State223, RparenToken}:                            _ReduceMixedToParametersAction,
	{_State224, _WildcardMarker}:                        _ReduceAddToNonEmptyParametersAction,
	{_State225, LbraceToken}:                            _ReduceInferredToImplicitFuncTypeAction,
	{_State226, LbraceToken}:                            _ReduceTypedToImplicitFuncTypeAction,
	{_State227, RparenToken}:                            _ReduceToArgumentsAction,
	{_State229, _WildcardMarker}:                        _ReduceImplicitToPackageStatementAction,
	{_State230, _WildcardMarker}:                        _ReduceExplicitToPackageStatementAction,
	{_State232, _WildcardMarker}:                        _ReduceBinaryOpAssignToStatementBodyAction,
	{_State233, _WildcardMarker}:                        _ReduceImplicitStructToExpressionOrImplicitStructAction,
	{_State234, _WildcardMarker}:                        _ReduceExpressionOrImplicitStructToOptionalExpressionOrImplicitStructAction,
	{_State235, _WildcardMarker}:                        _ReduceJumpToStatementBodyAction,
	{_State237, _EndMarker}:                             _ReduceIfElseToIfExprAction,
	{_State238, _EndMarker}:                             _ReduceMultiIfElseToIfExprAction,
	{_State240, _WildcardMarker}:                        _ReduceGenericToCallExprAction,
	{_State241, _WildcardMarker}:                        _ReduceToUnsafeStatementAction,
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
      func_decl -> State 14
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
      unary_op -> State 49
      unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 47
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
      SPACES -> State 51
      COMMENT -> State 50
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
      IDENTIFIER -> State 52

  State 12:
    Kernel Items:
      type_decl: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_decl: TYPE.IDENTIFIER EQUAL value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 53

  State 13:
    Kernel Items:
      func_decl: FUNC.optional_receiver IDENTIFIER optional_generic_parameters implicit_func_type
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      LPAREN -> State 54
      optional_receiver -> State 55

  State 14:
    Kernel Items:
      func_def: func_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 56
      block_body -> State 57

  State 15:
    Kernel Items:
      unary_op: BIT_AND., *
    Reduce:
      * -> [unary_op]
    Goto:
      (nil)

  State 16:
    Kernel Items:
      unary_op: BIT_NEG., *
    Reduce:
      * -> [unary_op]
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
      LPAREN -> State 58
      implicit_func_type -> State 59

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
      RPAREN -> [arguments]
    Goto:
      arguments -> State 60

  State 25:
    Kernel Items:
      unary_op: MUL., *
    Reduce:
      * -> [unary_op]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      unary_op: NOT., *
    Reduce:
      * -> [unary_op]
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
      LPAREN -> State 61
      implicit_struct_type -> State 62

  State 30:
    Kernel Items:
      unary_op: SUB., *
    Reduce:
      * -> [unary_op]
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
      call_expr: access_expr.LPAREN arguments RPAREN
      call_expr: access_expr.DOLLAR_LBRACE generic_arguments RBRACE LPAREN arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET expression RBRACKET
      unary_expr: access_expr., *
    Reduce:
      * -> [unary_expr]
    Goto:
      LPAREN -> State 66
      LBRACKET -> State 65
      DOT -> State 64
      DOLLAR_LBRACE -> State 63

  State 33:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 67
      SUB -> State 70
      BIT_XOR -> State 69
      BIT_OR -> State 68
      add_op -> State 71

  State 34:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 72

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
      EQUAL -> State 73
      NOT_EQUAL -> State 78
      LESS -> State 76
      LESS_OR_EQUAL -> State 77
      GREATER -> State 74
      GREATER_OR_EQUAL -> State 75
      cmp_op -> State 79

  State 41:
    Kernel Items:
      anonymous_func_expr: explicit_func_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 56
      block_body -> State 80

  State 42:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 81

  State 43:
    Kernel Items:
      atom_expr: literal., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 44:
    Kernel Items:
      mul_expr: mul_expr.mul_op unary_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 87
      DIV -> State 85
      MOD -> State 86
      BIT_AND -> State 82
      BIT_LSHIFT -> State 83
      BIT_RSHIFT -> State 84
      mul_op -> State 88

  State 45:
    Kernel Items:
      block_expr: optional_label_decl.block_body
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 90
      SWITCH -> State 91
      FOR -> State 89
      LBRACE -> State 56
      block_body -> State 92
      if_expr -> State 93
      switch_expr -> State 95
      loop_expr -> State 94

  State 46:
    Kernel Items:
      or_expr: or_expr.OR and_expr
      sequence_expr: or_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      OR -> State 96

  State 47:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      mul_expr: unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 49:
    Kernel Items:
      unary_expr: unary_op.unary_expr
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
      unary_op -> State 49
      unary_expr -> State 98
      optional_label_decl -> State 97
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 50:
    Kernel Items:
      lex_internal_tokens: COMMENT., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 51:
    Kernel Items:
      lex_internal_tokens: SPACES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 52:
    Kernel Items:
      package_decl: PACKAGE IDENTIFIER., $
      package_decl: PACKAGE IDENTIFIER.LPAREN package_statements RPAREN
    Reduce:
      $ -> [package_decl]
    Goto:
      LPAREN -> State 99

  State 53:
    Kernel Items:
      type_decl: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_decl: TYPE IDENTIFIER.EQUAL value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACE -> State 100
      EQUAL -> State 101
      optional_generic_parameters -> State 102

  State 54:
    Kernel Items:
      optional_receiver: LPAREN.parameter_decl RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 106
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      EXCLAIM -> State 104
      EXCLAIM_EXCLAIM -> State 105
      atom_type -> State 108
      traitable_type -> State 117
      trait_intersect_type -> State 114
      trait_union_type -> State 116
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109
      value_type -> State 118
      field_decl -> State 111
      parameter_decl -> State 113

  State 55:
    Kernel Items:
      func_decl: FUNC optional_receiver.IDENTIFIER optional_generic_parameters implicit_func_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 119

  State 56:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 120

  State 57:
    Kernel Items:
      func_def: func_decl block_body., $
    Reduce:
      $ -> [func_def]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      implicit_func_type: LPAREN.parameters RPAREN value_type
      implicit_func_type: LPAREN.parameters RPAREN QUESTION
      implicit_func_type: LPAREN.parameters RPAREN
    Reduce:
      RPAREN -> [optional_vararg]
    Goto:
      IDENTIFIER -> State 122
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      DOTDOTDOT -> State 121
      EXCLAIM -> State 104
      EXCLAIM_EXCLAIM -> State 105
      atom_type -> State 108
      traitable_type -> State 117
      trait_intersect_type -> State 114
      trait_union_type -> State 116
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109
      value_type -> State 118
      field_decl -> State 111
      parameter_decl -> State 125
      non_empty_parameters -> State 123
      vararg -> State 127
      optional_vararg -> State 124
      parameters -> State 126

  State 59:
    Kernel Items:
      explicit_func_type: FUNC implicit_func_type., LBRACE
    Reduce:
      LBRACE -> [explicit_func_type]
    Goto:
      (nil)

  State 60:
    Kernel Items:
      anonymous_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 128

  State 61:
    Kernel Items:
      implicit_struct_type: LPAREN.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 129

  State 62:
    Kernel Items:
      explicit_struct_type: STRUCT implicit_struct_type., *
    Reduce:
      * -> [explicit_struct_type]
    Goto:
      (nil)

  State 63:
    Kernel Items:
      call_expr: access_expr DOLLAR_LBRACE.generic_arguments RBRACE LPAREN arguments RPAREN
    Reduce:
      RBRACE -> [generic_arguments]
    Goto:
      generic_arguments -> State 130

  State 64:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131

  State 65:
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
      unary_op -> State 49
      unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 47
      optional_label_decl -> State 45
      block_expr -> State 38
      expression -> State 132
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 66:
    Kernel Items:
      call_expr: access_expr LPAREN.arguments RPAREN
    Reduce:
      RPAREN -> [arguments]
    Goto:
      arguments -> State 133

  State 67:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 68:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 69:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 70:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 71:
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
      unary_op -> State 49
      unary_expr -> State 48
      mul_expr -> State 134
      optional_label_decl -> State 97
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 72:
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
      unary_op -> State 49
      unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 135
      optional_label_decl -> State 97
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 73:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 74:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 75:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 76:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 77:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 78:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 79:
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
      unary_op -> State 49
      unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 136
      optional_label_decl -> State 97
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 80:
    Kernel Items:
      anonymous_func_expr: explicit_func_type block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 81:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type LPAREN.arguments RPAREN
    Reduce:
      RPAREN -> [arguments]
    Goto:
      arguments -> State 137

  State 82:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 84:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 85:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 86:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 87:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      mul_expr: mul_expr mul_op.unary_expr
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
      unary_op -> State 49
      unary_expr -> State 138
      optional_label_decl -> State 97
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 89:
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
      unary_op -> State 49
      unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 140
      optional_label_decl -> State 97
      block_expr -> State 139
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 90:
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
      unary_op -> State 49
      unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 141
      optional_label_decl -> State 97
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 91:
    Kernel Items:
      switch_expr: SWITCH.CASE DEFAULT
    Reduce:
      (nil)
    Goto:
      CASE -> State 142

  State 92:
    Kernel Items:
      block_expr: optional_label_decl block_body., $
    Reduce:
      $ -> [block_expr]
    Goto:
      (nil)

  State 93:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 94:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 95:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 96:
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
      unary_op -> State 49
      unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 143
      optional_label_decl -> State 97
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 97:
    Kernel Items:
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 56
      block_body -> State 92

  State 98:
    Kernel Items:
      unary_expr: unary_op unary_expr., *
    Reduce:
      * -> [unary_expr]
    Goto:
      (nil)

  State 99:
    Kernel Items:
      package_decl: PACKAGE IDENTIFIER LPAREN.package_statements RPAREN
    Reduce:
      * -> [package_statements]
    Goto:
      package_statements -> State 144

  State 100:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACE.generic_parameters RBRACE
    Reduce:
      RBRACE -> [generic_parameters]
    Goto:
      generic_parameters -> State 145

  State 101:
    Kernel Items:
      type_decl: TYPE IDENTIFIER EQUAL.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      EXCLAIM -> State 104
      EXCLAIM_EXCLAIM -> State 105
      atom_type -> State 108
      traitable_type -> State 117
      trait_intersect_type -> State 114
      trait_union_type -> State 116
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109
      value_type -> State 147

  State 102:
    Kernel Items:
      type_decl: TYPE IDENTIFIER optional_generic_parameters.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      EXCLAIM -> State 104
      EXCLAIM_EXCLAIM -> State 105
      atom_type -> State 108
      traitable_type -> State 117
      trait_intersect_type -> State 114
      trait_union_type -> State 116
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109
      value_type -> State 148

  State 103:
    Kernel Items:
      explicit_enum_type: ENUM.LPAREN RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 149

  State 104:
    Kernel Items:
      traitable_type: EXCLAIM.atom_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      atom_type -> State 150
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109

  State 105:
    Kernel Items:
      traitable_type: EXCLAIM_EXCLAIM.atom_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      atom_type -> State 151
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109

  State 106:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_parameters
      field_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.QUESTION
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      QUESTION -> State 152
      DOLLAR_LBRACE -> State 100
      EXCLAIM -> State 104
      EXCLAIM_EXCLAIM -> State 105
      optional_generic_parameters -> State 153
      atom_type -> State 108
      traitable_type -> State 117
      trait_intersect_type -> State 114
      trait_union_type -> State 116
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109
      value_type -> State 154

  State 107:
    Kernel Items:
      trait_type: TRAIT.implicit_struct_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 61
      implicit_struct_type -> State 155

  State 108:
    Kernel Items:
      traitable_type: atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 109:
    Kernel Items:
      atom_type: explicit_enum_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 110:
    Kernel Items:
      atom_type: explicit_struct_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 111:
    Kernel Items:
      parameter_decl: field_decl., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 112:
    Kernel Items:
      atom_type: implicit_struct_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 113:
    Kernel Items:
      optional_receiver: LPAREN parameter_decl.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 156

  State 114:
    Kernel Items:
      trait_intersect_type: trait_intersect_type.BIT_AND traitable_type
      trait_intersect_type: trait_intersect_type.SUB traitable_type
      trait_union_type: trait_intersect_type., *
    Reduce:
      * -> [trait_union_type]
    Goto:
      SUB -> State 158
      BIT_AND -> State 157

  State 115:
    Kernel Items:
      atom_type: trait_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      trait_union_type: trait_union_type.BIT_OR trait_intersect_type
      value_type: trait_union_type., *
    Reduce:
      * -> [value_type]
    Goto:
      BIT_OR -> State 159

  State 117:
    Kernel Items:
      trait_intersect_type: traitable_type., *
    Reduce:
      * -> [trait_intersect_type]
    Goto:
      (nil)

  State 118:
    Kernel Items:
      value_type: value_type.OR trait_union_type
      field_decl: value_type., *
    Reduce:
      * -> [field_decl]
    Goto:
      OR -> State 160

  State 119:
    Kernel Items:
      func_decl: FUNC optional_receiver IDENTIFIER.optional_generic_parameters implicit_func_type
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACE -> State 100
      optional_generic_parameters -> State 161

  State 120:
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
      RETURN -> State 166
      BREAK -> State 163
      CONTINUE -> State 164
      UNSAFE -> State 167
      STRUCT -> State 29
      FUNC -> State 19
      ASYNC -> State 162
      RBRACE -> State 165
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
      access_expr -> State 168
      unary_op -> State 49
      unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 47
      jump_type -> State 171
      expression_or_implicit_struct -> State 170
      unsafe_statement -> State 174
      statement_body -> State 173
      statement -> State 172
      optional_label_decl -> State 45
      block_expr -> State 38
      expression -> State 169
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 121:
    Kernel Items:
      vararg: DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      EXCLAIM -> State 104
      EXCLAIM_EXCLAIM -> State 105
      atom_type -> State 108
      traitable_type -> State 117
      trait_intersect_type -> State 114
      trait_union_type -> State 116
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109
      value_type -> State 175

  State 122:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_parameters
      field_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.QUESTION
      vararg: IDENTIFIER.DOTDOTDOT value_type
      vararg: IDENTIFIER.DOTDOTDOT QUESTION
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      QUESTION -> State 152
      DOLLAR_LBRACE -> State 100
      DOTDOTDOT -> State 176
      EXCLAIM -> State 104
      EXCLAIM_EXCLAIM -> State 105
      optional_generic_parameters -> State 153
      atom_type -> State 108
      traitable_type -> State 117
      trait_intersect_type -> State 114
      trait_union_type -> State 116
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109
      value_type -> State 154

  State 123:
    Kernel Items:
      non_empty_parameters: non_empty_parameters.COMMA parameter_decl
      parameters: non_empty_parameters., RPAREN
      parameters: non_empty_parameters.COMMA optional_vararg
    Reduce:
      RPAREN -> [parameters]
    Goto:
      COMMA -> State 177

  State 124:
    Kernel Items:
      parameters: optional_vararg., RPAREN
    Reduce:
      RPAREN -> [parameters]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      non_empty_parameters: parameter_decl., *
    Reduce:
      * -> [non_empty_parameters]
    Goto:
      (nil)

  State 126:
    Kernel Items:
      implicit_func_type: LPAREN parameters.RPAREN value_type
      implicit_func_type: LPAREN parameters.RPAREN QUESTION
      implicit_func_type: LPAREN parameters.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 178

  State 127:
    Kernel Items:
      optional_vararg: vararg., RPAREN
      optional_vararg: vararg.COMMA
    Reduce:
      RPAREN -> [optional_vararg]
    Goto:
      COMMA -> State 179

  State 128:
    Kernel Items:
      anonymous_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 129:
    Kernel Items:
      implicit_struct_type: LPAREN RPAREN., *
    Reduce:
      * -> [implicit_struct_type]
    Goto:
      (nil)

  State 130:
    Kernel Items:
      call_expr: access_expr DOLLAR_LBRACE generic_arguments.RBRACE LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 180

  State 131:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 132:
    Kernel Items:
      access_expr: access_expr LBRACKET expression.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 181

  State 133:
    Kernel Items:
      call_expr: access_expr LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 182

  State 134:
    Kernel Items:
      mul_expr: mul_expr.mul_op unary_expr
      add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 87
      DIV -> State 85
      MOD -> State 86
      BIT_AND -> State 82
      BIT_LSHIFT -> State 83
      BIT_RSHIFT -> State 84
      mul_op -> State 88

  State 135:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 73
      NOT_EQUAL -> State 78
      LESS -> State 76
      LESS_OR_EQUAL -> State 77
      GREATER -> State 74
      GREATER_OR_EQUAL -> State 75
      cmp_op -> State 79

  State 136:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 67
      SUB -> State 70
      BIT_XOR -> State 69
      BIT_OR -> State 68
      add_op -> State 71

  State 137:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 183

  State 138:
    Kernel Items:
      mul_expr: mul_expr mul_op unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 139:
    Kernel Items:
      atom_expr: block_expr., *
      loop_expr: FOR block_expr., $
    Reduce:
      * -> [atom_expr]
      $ -> [loop_expr]
    Goto:
      (nil)

  State 140:
    Kernel Items:
      loop_expr: FOR sequence_expr.block_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      LABEL_DECL -> State 22
      optional_label_decl -> State 97
      block_expr -> State 184

  State 141:
    Kernel Items:
      if_expr: IF sequence_expr.block_body
      if_expr: IF sequence_expr.block_body ELSE block_body
      if_expr: IF sequence_expr.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 56
      block_body -> State 185

  State 142:
    Kernel Items:
      switch_expr: SWITCH CASE.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 186

  State 143:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 72

  State 144:
    Kernel Items:
      package_decl: PACKAGE IDENTIFIER LPAREN package_statements.RPAREN
      package_statements: package_statements.package_statement
    Reduce:
      (nil)
    Goto:
      UNSAFE -> State 167
      RPAREN -> State 187
      unsafe_statement -> State 190
      package_statement_body -> State 189
      package_statement -> State 188

  State 145:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACE generic_parameters.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 191

  State 146:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_parameters
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACE -> State 100
      optional_generic_parameters -> State 153

  State 147:
    Kernel Items:
      value_type: value_type.OR trait_union_type
      type_decl: TYPE IDENTIFIER EQUAL value_type., $
    Reduce:
      $ -> [type_decl]
    Goto:
      OR -> State 160

  State 148:
    Kernel Items:
      value_type: value_type.OR trait_union_type
      type_decl: TYPE IDENTIFIER optional_generic_parameters value_type., $
    Reduce:
      $ -> [type_decl]
    Goto:
      OR -> State 160

  State 149:
    Kernel Items:
      explicit_enum_type: ENUM LPAREN.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 192

  State 150:
    Kernel Items:
      traitable_type: EXCLAIM atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 151:
    Kernel Items:
      traitable_type: EXCLAIM_EXCLAIM atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 152:
    Kernel Items:
      parameter_decl: IDENTIFIER QUESTION., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 153:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_parameters., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 154:
    Kernel Items:
      value_type: value_type.OR trait_union_type
      field_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [field_decl]
    Goto:
      OR -> State 160

  State 155:
    Kernel Items:
      trait_type: TRAIT implicit_struct_type., *
    Reduce:
      * -> [trait_type]
    Goto:
      (nil)

  State 156:
    Kernel Items:
      optional_receiver: LPAREN parameter_decl RPAREN., IDENTIFIER
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      (nil)

  State 157:
    Kernel Items:
      trait_intersect_type: trait_intersect_type BIT_AND.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      EXCLAIM -> State 104
      EXCLAIM_EXCLAIM -> State 105
      atom_type -> State 108
      traitable_type -> State 193
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109

  State 158:
    Kernel Items:
      trait_intersect_type: trait_intersect_type SUB.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      EXCLAIM -> State 104
      EXCLAIM_EXCLAIM -> State 105
      atom_type -> State 108
      traitable_type -> State 194
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109

  State 159:
    Kernel Items:
      trait_union_type: trait_union_type BIT_OR.trait_intersect_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      EXCLAIM -> State 104
      EXCLAIM_EXCLAIM -> State 105
      atom_type -> State 108
      traitable_type -> State 117
      trait_intersect_type -> State 195
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109

  State 160:
    Kernel Items:
      value_type: value_type OR.trait_union_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      EXCLAIM -> State 104
      EXCLAIM_EXCLAIM -> State 105
      atom_type -> State 108
      traitable_type -> State 117
      trait_intersect_type -> State 114
      trait_union_type -> State 196
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109

  State 161:
    Kernel Items:
      func_decl: FUNC optional_receiver IDENTIFIER optional_generic_parameters.implicit_func_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 58
      implicit_func_type -> State 197

  State 162:
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
      call_expr -> State 199
      access_expr -> State 198
      optional_label_decl -> State 97
      block_expr -> State 38
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 163:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 164:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 165:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 166:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 167:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 200

  State 168:
    Kernel Items:
      call_expr: access_expr.LPAREN arguments RPAREN
      call_expr: access_expr.DOLLAR_LBRACE generic_arguments RBRACE LPAREN arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET expression RBRACKET
      unary_expr: access_expr., *
      statement_body: access_expr.op_one_assign
      statement_body: access_expr.binary_op_assign expression
    Reduce:
      * -> [unary_expr]
    Goto:
      LPAREN -> State 66
      LBRACKET -> State 65
      DOT -> State 64
      DOLLAR_LBRACE -> State 63
      ADD_ASSIGN -> State 201
      SUB_ASSIGN -> State 212
      MUL_ASSIGN -> State 211
      DIV_ASSIGN -> State 209
      MOD_ASSIGN -> State 210
      ADD_ONE_ASSIGN -> State 202
      SUB_ONE_ASSIGN -> State 213
      BIT_NEG_ASSIGN -> State 205
      BIT_AND_ASSIGN -> State 203
      BIT_OR_ASSIGN -> State 206
      BIT_XOR_ASSIGN -> State 208
      BIT_LSHIFT_ASSIGN -> State 204
      BIT_RSHIFT_ASSIGN -> State 207
      op_one_assign -> State 215
      binary_op_assign -> State 214

  State 169:
    Kernel Items:
      expression_or_implicit_struct: expression., *
    Reduce:
      * -> [expression_or_implicit_struct]
    Goto:
      (nil)

  State 170:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct.COMMA expression
      statement_body: expression_or_implicit_struct., *
    Reduce:
      * -> [statement_body]
    Goto:
      COMMA -> State 216

  State 171:
    Kernel Items:
      statement_body: jump_type.optional_jump_label optional_expression_or_implicit_struct
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 217
      optional_jump_label -> State 218

  State 172:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 173:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 219
      SEMICOLON -> State 220

  State 174:
    Kernel Items:
      statement_body: unsafe_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 175:
    Kernel Items:
      value_type: value_type.OR trait_union_type
      vararg: DOTDOTDOT value_type., *
    Reduce:
      * -> [vararg]
    Goto:
      OR -> State 160

  State 176:
    Kernel Items:
      vararg: IDENTIFIER DOTDOTDOT.value_type
      vararg: IDENTIFIER DOTDOTDOT.QUESTION
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      QUESTION -> State 221
      EXCLAIM -> State 104
      EXCLAIM_EXCLAIM -> State 105
      atom_type -> State 108
      traitable_type -> State 117
      trait_intersect_type -> State 114
      trait_union_type -> State 116
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109
      value_type -> State 222

  State 177:
    Kernel Items:
      non_empty_parameters: non_empty_parameters COMMA.parameter_decl
      parameters: non_empty_parameters COMMA.optional_vararg
    Reduce:
      RPAREN -> [optional_vararg]
    Goto:
      IDENTIFIER -> State 122
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      DOTDOTDOT -> State 121
      EXCLAIM -> State 104
      EXCLAIM_EXCLAIM -> State 105
      atom_type -> State 108
      traitable_type -> State 117
      trait_intersect_type -> State 114
      trait_union_type -> State 116
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109
      value_type -> State 118
      field_decl -> State 111
      parameter_decl -> State 224
      vararg -> State 127
      optional_vararg -> State 223

  State 178:
    Kernel Items:
      implicit_func_type: LPAREN parameters RPAREN.value_type
      implicit_func_type: LPAREN parameters RPAREN.QUESTION
      implicit_func_type: LPAREN parameters RPAREN., LBRACE
    Reduce:
      LBRACE -> [implicit_func_type]
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 29
      ENUM -> State 103
      TRAIT -> State 107
      LPAREN -> State 61
      QUESTION -> State 225
      EXCLAIM -> State 104
      EXCLAIM_EXCLAIM -> State 105
      atom_type -> State 108
      traitable_type -> State 117
      trait_intersect_type -> State 114
      trait_union_type -> State 116
      implicit_struct_type -> State 112
      explicit_struct_type -> State 110
      trait_type -> State 115
      explicit_enum_type -> State 109
      value_type -> State 226

  State 179:
    Kernel Items:
      optional_vararg: vararg COMMA., RPAREN
    Reduce:
      RPAREN -> [optional_vararg]
    Goto:
      (nil)

  State 180:
    Kernel Items:
      call_expr: access_expr DOLLAR_LBRACE generic_arguments RBRACE.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 227

  State 181:
    Kernel Items:
      access_expr: access_expr LBRACKET expression RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 182:
    Kernel Items:
      call_expr: access_expr LPAREN arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 183:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 184:
    Kernel Items:
      loop_expr: FOR sequence_expr block_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 185:
    Kernel Items:
      if_expr: IF sequence_expr block_body., *
      if_expr: IF sequence_expr block_body.ELSE block_body
      if_expr: IF sequence_expr block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 228

  State 186:
    Kernel Items:
      switch_expr: SWITCH CASE DEFAULT., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 187:
    Kernel Items:
      package_decl: PACKAGE IDENTIFIER LPAREN package_statements RPAREN., $
    Reduce:
      $ -> [package_decl]
    Goto:
      (nil)

  State 188:
    Kernel Items:
      package_statements: package_statements package_statement., *
    Reduce:
      * -> [package_statements]
    Goto:
      (nil)

  State 189:
    Kernel Items:
      package_statement: package_statement_body.NEWLINES
      package_statement: package_statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 229
      SEMICOLON -> State 230

  State 190:
    Kernel Items:
      package_statement_body: unsafe_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 191:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACE generic_parameters RBRACE., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 192:
    Kernel Items:
      explicit_enum_type: ENUM LPAREN RPAREN., *
    Reduce:
      * -> [explicit_enum_type]
    Goto:
      (nil)

  State 193:
    Kernel Items:
      trait_intersect_type: trait_intersect_type BIT_AND traitable_type., *
    Reduce:
      * -> [trait_intersect_type]
    Goto:
      (nil)

  State 194:
    Kernel Items:
      trait_intersect_type: trait_intersect_type SUB traitable_type., *
    Reduce:
      * -> [trait_intersect_type]
    Goto:
      (nil)

  State 195:
    Kernel Items:
      trait_intersect_type: trait_intersect_type.BIT_AND traitable_type
      trait_intersect_type: trait_intersect_type.SUB traitable_type
      trait_union_type: trait_union_type BIT_OR trait_intersect_type., *
    Reduce:
      * -> [trait_union_type]
    Goto:
      SUB -> State 158
      BIT_AND -> State 157

  State 196:
    Kernel Items:
      trait_union_type: trait_union_type.BIT_OR trait_intersect_type
      value_type: value_type OR trait_union_type., *
    Reduce:
      * -> [value_type]
    Goto:
      BIT_OR -> State 159

  State 197:
    Kernel Items:
      func_decl: FUNC optional_receiver IDENTIFIER optional_generic_parameters implicit_func_type., LBRACE
    Reduce:
      LBRACE -> [func_decl]
    Goto:
      (nil)

  State 198:
    Kernel Items:
      call_expr: access_expr.LPAREN arguments RPAREN
      call_expr: access_expr.DOLLAR_LBRACE generic_arguments RBRACE LPAREN arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET expression RBRACKET
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 66
      LBRACKET -> State 65
      DOT -> State 64
      DOLLAR_LBRACE -> State 63

  State 199:
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

  State 200:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 231

  State 201:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 202:
    Kernel Items:
      op_one_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 203:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 204:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 205:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 206:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 207:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 208:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 209:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 211:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 212:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      op_one_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 214:
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
      unary_op -> State 49
      unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 47
      optional_label_decl -> State 45
      block_expr -> State 38
      expression -> State 232
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 215:
    Kernel Items:
      statement_body: access_expr op_one_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 216:
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
      unary_op -> State 49
      unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 47
      optional_label_decl -> State 45
      block_expr -> State 38
      expression -> State 233
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 217:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 218:
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
      unary_op -> State 49
      unary_expr -> State 48
      mul_expr -> State 44
      add_expr -> State 33
      cmp_expr -> State 40
      and_expr -> State 34
      or_expr -> State 46
      sequence_expr -> State 47
      optional_expression_or_implicit_struct -> State 235
      expression_or_implicit_struct -> State 234
      optional_label_decl -> State 45
      block_expr -> State 38
      expression -> State 169
      explicit_struct_type -> State 42
      explicit_func_type -> State 41

  State 219:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 220:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 221:
    Kernel Items:
      vararg: IDENTIFIER DOTDOTDOT QUESTION., *
    Reduce:
      * -> [vararg]
    Goto:
      (nil)

  State 222:
    Kernel Items:
      value_type: value_type.OR trait_union_type
      vararg: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [vararg]
    Goto:
      OR -> State 160

  State 223:
    Kernel Items:
      parameters: non_empty_parameters COMMA optional_vararg., RPAREN
    Reduce:
      RPAREN -> [parameters]
    Goto:
      (nil)

  State 224:
    Kernel Items:
      non_empty_parameters: non_empty_parameters COMMA parameter_decl., *
    Reduce:
      * -> [non_empty_parameters]
    Goto:
      (nil)

  State 225:
    Kernel Items:
      implicit_func_type: LPAREN parameters RPAREN QUESTION., LBRACE
    Reduce:
      LBRACE -> [implicit_func_type]
    Goto:
      (nil)

  State 226:
    Kernel Items:
      value_type: value_type.OR trait_union_type
      implicit_func_type: LPAREN parameters RPAREN value_type., LBRACE
    Reduce:
      LBRACE -> [implicit_func_type]
    Goto:
      OR -> State 160

  State 227:
    Kernel Items:
      call_expr: access_expr DOLLAR_LBRACE generic_arguments RBRACE LPAREN.arguments RPAREN
    Reduce:
      RPAREN -> [arguments]
    Goto:
      arguments -> State 236

  State 228:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE.block_body
      if_expr: IF sequence_expr block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 90
      LBRACE -> State 56
      block_body -> State 237
      if_expr -> State 238

  State 229:
    Kernel Items:
      package_statement: package_statement_body NEWLINES., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 230:
    Kernel Items:
      package_statement: package_statement_body SEMICOLON., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 231:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 239

  State 232:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 233:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct COMMA expression., *
    Reduce:
      * -> [expression_or_implicit_struct]
    Goto:
      (nil)

  State 234:
    Kernel Items:
      optional_expression_or_implicit_struct: expression_or_implicit_struct., *
      expression_or_implicit_struct: expression_or_implicit_struct.COMMA expression
    Reduce:
      * -> [optional_expression_or_implicit_struct]
    Goto:
      COMMA -> State 216

  State 235:
    Kernel Items:
      statement_body: jump_type optional_jump_label optional_expression_or_implicit_struct., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 236:
    Kernel Items:
      call_expr: access_expr DOLLAR_LBRACE generic_arguments RBRACE LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 240

  State 237:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 238:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 239:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 241

  State 240:
    Kernel Items:
      call_expr: access_expr DOLLAR_LBRACE generic_arguments RBRACE LPAREN arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

Number of states: 241
Number of shift actions: 943
Number of reduce actions: 193
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
