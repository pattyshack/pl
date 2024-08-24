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
	BoolLiteralToken     = SymbolId(259)
	IntegerLiteralToken  = SymbolId(260)
	FloatLiteralToken    = SymbolId(261)
	RuneLiteralToken     = SymbolId(262)
	StringLiteralToken   = SymbolId(263)
	LbraceToken          = SymbolId(264)
	RbraceToken          = SymbolId(265)
	LparenToken          = SymbolId(266)
	RparenToken          = SymbolId(267)
	LbracketToken        = SymbolId(268)
	RbracketToken        = SymbolId(269)
	GenericLbraceToken   = SymbolId(270)
	DotToken             = SymbolId(271)
	CommaToken           = SymbolId(272)
	QuestionToken        = SymbolId(273)
	SemicolonToken       = SymbolId(274)
	DotdotdotToken       = SymbolId(275)
	IdentifierToken      = SymbolId(276)
	IfToken              = SymbolId(277)
	ElseToken            = SymbolId(278)
	MatchToken           = SymbolId(279)
	CaseToken            = SymbolId(280)
	DefaultToken         = SymbolId(281)
	ForToken             = SymbolId(282)
	ReturnToken          = SymbolId(283)
	BreakToken           = SymbolId(284)
	ContinueToken        = SymbolId(285)
	TypeToken            = SymbolId(286)
	StructToken          = SymbolId(287)
	EnumToken            = SymbolId(288)
	FuncToken            = SymbolId(289)
	AsyncToken           = SymbolId(290)
	LabelDeclToken       = SymbolId(291)
	JumpLabelToken       = SymbolId(292)
	AddAssignToken       = SymbolId(293)
	SubAssignToken       = SymbolId(294)
	MulAssignToken       = SymbolId(295)
	DivAssignToken       = SymbolId(296)
	ModAssignToken       = SymbolId(297)
	AddOneAssignToken    = SymbolId(298)
	SubOneAssignToken    = SymbolId(299)
	BitNegAssignToken    = SymbolId(300)
	BitAndAssignToken    = SymbolId(301)
	BitOrAssignToken     = SymbolId(302)
	BitXorAssignToken    = SymbolId(303)
	BitLshiftAssignToken = SymbolId(304)
	BitRshiftAssignToken = SymbolId(305)
	NotToken             = SymbolId(306)
	AndToken             = SymbolId(307)
	OrToken              = SymbolId(308)
	AddToken             = SymbolId(309)
	SubToken             = SymbolId(310)
	MulToken             = SymbolId(311)
	DivToken             = SymbolId(312)
	ModToken             = SymbolId(313)
	BitNegToken          = SymbolId(314)
	BitAndToken          = SymbolId(315)
	BitXorToken          = SymbolId(316)
	BitOrToken           = SymbolId(317)
	BitLshiftToken       = SymbolId(318)
	BitRshiftToken       = SymbolId(319)
	EqualToken           = SymbolId(320)
	NotEqualToken        = SymbolId(321)
	LessToken            = SymbolId(322)
	LessOrEqualToken     = SymbolId(323)
	GreaterToken         = SymbolId(324)
	GreaterOrEqualToken  = SymbolId(325)
	LexErrorToken        = SymbolId(326)
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
	// 83:2: literal -> BOOL_LITERAL: ...
	BoolLiteralToLiteral(BoolLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 84:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 85:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 86:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 87:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 89:23: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(ExplicitFuncType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 94:2: anonymous_struct_expr -> explicit: ...
	ExplicitToAnonymousStructExpr(ExplicitStructType_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 95:2: anonymous_struct_expr -> implicit: ...
	ImplicitToAnonymousStructExpr(Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 101:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 102:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 103:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 105:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 106:2: atom_expr -> anonymous_struct_expr: ...
	AnonymousStructExprToAtomExpr(AnonymousStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 107:2: atom_expr -> LEX_ERROR: ...
	LexErrorToAtomExpr(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 110:0: generic_arguments -> ...
	ToGenericArguments() (*GenericSymbol, error)

	// 113:0: arguments -> ...
	ToArguments() (*GenericSymbol, error)

	// 116:2: call_expr -> concrete: ...
	ConcreteToCallExpr(AccessExpr_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 117:2: call_expr -> generic: ...
	GenericToCallExpr(AccessExpr_ *GenericSymbol, GenericLbrace_ *GenericSymbol, GenericArguments_ *GenericSymbol, Rbrace_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 120:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 121:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 122:2: access_expr -> call_expr: ...
	CallExprToAccessExpr(CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 124:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Expression_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 127:2: unary_op -> NOT: ...
	NotToUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 128:2: unary_op -> BIT_NEG: ...
	BitNegToUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 129:2: unary_op -> SUB: ...
	SubToUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 132:2: unary_op -> MUL: ...
	MulToUnaryOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 135:2: unary_op -> BIT_AND: ...
	BitAndToUnaryOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 138:2: unary_expr -> access_expr: ...
	AccessExprToUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 139:2: unary_expr -> op: ...
	OpToUnaryExpr(UnaryOp_ *GenericSymbol, UnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 142:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 143:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 144:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 145:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 146:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 147:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 150:2: mul_expr -> unary_expr: ...
	UnaryExprToMulExpr(UnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 151:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, UnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 154:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 155:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 156:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 157:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 160:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 161:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 164:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 165:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 166:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 167:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 168:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 169:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 172:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 173:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 176:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 177:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 180:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 181:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 183:17: sequence_expr -> ...
	ToSequenceExpr(OrExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 186:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 187:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 190:2: optional_expression_or_implicit_struct -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct(ExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 191:2: optional_expression_or_implicit_struct -> nil: ...
	NilToOptionalExpressionOrImplicitStruct() (*GenericSymbol, error)

	// 194:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 195:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 196:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 199:2: op_one_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToOpOneAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 200:2: op_one_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToOpOneAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 203:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 204:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 205:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 206:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 207:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 208:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 209:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 210:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 211:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 212:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 213:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 216:2: expression_or_implicit_struct -> expression: ...
	ExpressionToExpressionOrImplicitStruct(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 217:2: expression_or_implicit_struct -> implicit_struct: ...
	ImplicitStructToExpressionOrImplicitStruct(ExpressionOrImplicitStruct_ *GenericSymbol, Comma_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 220:2: statement_body -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToStatementBody(ExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 222:2: statement_body -> async: ...
	AsyncToStatementBody(Async_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 225:2: statement_body -> jump: ...
	JumpToStatementBody(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 238:2: statement_body -> op_one_assign: ...
	OpOneAssignToStatementBody(AccessExpr_ *GenericSymbol, OpOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 239:2: statement_body -> binary_op_assign: ...
	BinaryOpAssignToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 242:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 243:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 246:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 247:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 250:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 251:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (*GenericSymbol, error)

	// 253:14: block_body -> ...
	ToBlockBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 254:14: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 259:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 260:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, BlockBody_2 *GenericSymbol) (*GenericSymbol, error)

	// 261:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 264:14: match_expr -> ...
	ToMatchExpr(Match_ *GenericSymbol, Case_ *GenericSymbol, Default_ *GenericSymbol) (*GenericSymbol, error)

	// 267:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(For_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 268:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 274:2: expression -> sequence_expr: ...
	SequenceExprToExpression(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 275:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 276:2: expression -> match_expr: ...
	MatchExprToExpression(OptionalLabelDecl_ *GenericSymbol, MatchExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 277:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ *GenericSymbol, LoopExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 280:0: generic_parameters -> ...
	ToGenericParameters() (*GenericSymbol, error)

	// 283:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(GenericLbrace_ *GenericSymbol, GenericParameters_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 284:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (*GenericSymbol, error)

	// 287:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol) (*GenericSymbol, error)

	// 288:2: atom_type -> explicit_struct_type: ...
	ExplicitStructTypeToAtomType(ExplicitStructType_ *GenericSymbol) (*GenericSymbol, error)

	// 289:2: atom_type -> implicit_struct_type: ...
	ImplicitStructTypeToAtomType(ImplicitStructType_ *GenericSymbol) (*GenericSymbol, error)

	// 290:2: atom_type -> explicit_enum_type: ...
	ExplicitEnumTypeToAtomType(ExplicitEnumType_ *GenericSymbol) (*GenericSymbol, error)

	// 293:24: implicit_struct_type -> ...
	ToImplicitStructType(Lparen_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 296:2: explicit_struct_type -> ...
	ToExplicitStructType(Struct_ *GenericSymbol, ImplicitStructType_ *GenericSymbol) (*GenericSymbol, error)

	// 299:22: explicit_enum_type -> ...
	ToExplicitEnumType(Enum_ *GenericSymbol, Lparen_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 302:2: implicit_enum_type -> pair: ...
	PairToImplicitEnumType(AtomType_ *GenericSymbol, BitOr_ *GenericSymbol, AtomType_2 *GenericSymbol) (*GenericSymbol, error)

	// 303:2: implicit_enum_type -> add: ...
	AddToImplicitEnumType(ImplicitEnumType_ *GenericSymbol, BitOr_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 306:2: value_type -> atom_type: ...
	AtomTypeToValueType(AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 307:2: value_type -> implicit_enum_type: ...
	ImplicitEnumTypeToValueType(ImplicitEnumType_ *GenericSymbol) (*GenericSymbol, error)

	// 310:2: type_decl -> definition: ...
	DefinitionToTypeDecl(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 311:2: type_decl -> alias: ...
	AliasToTypeDecl(Type_ *GenericSymbol, Identifier_ *GenericSymbol, Equal_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 314:2: parameter_decl -> explicit: ...
	ExplicitToParameterDecl(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 315:2: parameter_decl -> implicit: ...
	ImplicitToParameterDecl(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 316:2: parameter_decl -> inferred: ...
	InferredToParameterDecl(Identifier_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 319:2: non_empty_parameters -> parameter_decl: ...
	ParameterDeclToNonEmptyParameters(ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 320:2: non_empty_parameters -> add: ...
	AddToNonEmptyParameters(NonEmptyParameters_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 323:2: vararg -> explicit: ...
	ExplicitToVararg(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 324:2: vararg -> implicit: ...
	ImplicitToVararg(Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 325:2: vararg -> inferred: ...
	InferredToVararg(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 328:2: optional_vararg -> vararg: ...
	VarargToOptionalVararg(Vararg_ *GenericSymbol) (*GenericSymbol, error)

	// 329:2: optional_vararg -> vararg2: ...
	Vararg2ToOptionalVararg(Vararg_ *GenericSymbol, Comma_ *GenericSymbol) (*GenericSymbol, error)

	// 330:2: optional_vararg -> nil: ...
	NilToOptionalVararg() (*GenericSymbol, error)

	// 333:2: parameters -> non_empty_parameters: ...
	NonEmptyParametersToParameters(NonEmptyParameters_ *GenericSymbol) (*GenericSymbol, error)

	// 334:2: parameters -> mixed: ...
	MixedToParameters(NonEmptyParameters_ *GenericSymbol, Comma_ *GenericSymbol, OptionalVararg_ *GenericSymbol) (*GenericSymbol, error)

	// 335:2: parameters -> optional_vararg: ...
	OptionalVarargToParameters(OptionalVararg_ *GenericSymbol) (*GenericSymbol, error)

	// 338:2: implicit_func_type -> typed: ...
	TypedToImplicitFuncType(Lparen_ *GenericSymbol, Parameters_ *GenericSymbol, Rparen_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 339:2: implicit_func_type -> inferred: ...
	InferredToImplicitFuncType(Lparen_ *GenericSymbol, Parameters_ *GenericSymbol, Rparen_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 340:2: implicit_func_type -> implicit_unit: ...
	ImplicitUnitToImplicitFuncType(Lparen_ *GenericSymbol, Parameters_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 342:22: explicit_func_type -> ...
	ToExplicitFuncType(Func_ *GenericSymbol, ImplicitFuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 345:2: optional_receiver -> receiver: ...
	ReceiverToOptionalReceiver(Lparen_ *GenericSymbol, ParameterDecl_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 346:2: optional_receiver -> nil: ...
	NilToOptionalReceiver() (*GenericSymbol, error)

	// 349:2: func_decl -> ...
	ToFuncDecl(Func_ *GenericSymbol, OptionalReceiver_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ImplicitFuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 351:12: func_def -> ...
	ToFuncDef(FuncDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 355:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 356:2: lex_internal_tokens -> COMMENT: ...
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

	item, err := _Parse(lexer, reducer, errHandler, _State1)
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

	item, err := _Parse(lexer, reducer, errHandler, _State2)
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

	item, err := _Parse(lexer, reducer, errHandler, _State3)
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

	item, err := _Parse(lexer, reducer, errHandler, _State4)
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
	case BoolLiteralToken:
		return "BOOL_LITERAL"
	case IntegerLiteralToken:
		return "INTEGER_LITERAL"
	case FloatLiteralToken:
		return "FLOAT_LITERAL"
	case RuneLiteralToken:
		return "RUNE_LITERAL"
	case StringLiteralToken:
		return "STRING_LITERAL"
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
	case GenericLbraceToken:
		return "GENERIC_LBRACE"
	case DotToken:
		return "DOT"
	case CommaToken:
		return "COMMA"
	case QuestionToken:
		return "QUESTION"
	case SemicolonToken:
		return "SEMICOLON"
	case DotdotdotToken:
		return "DOTDOTDOT"
	case IdentifierToken:
		return "IDENTIFIER"
	case IfToken:
		return "IF"
	case ElseToken:
		return "ELSE"
	case MatchToken:
		return "MATCH"
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
	case TypeToken:
		return "TYPE"
	case StructToken:
		return "STRUCT"
	case EnumToken:
		return "ENUM"
	case FuncToken:
		return "FUNC"
	case AsyncToken:
		return "ASYNC"
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
	case MatchExprType:
		return "match_expr"
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
	case ImplicitStructTypeType:
		return "implicit_struct_type"
	case ExplicitStructTypeType:
		return "explicit_struct_type"
	case ExplicitEnumTypeType:
		return "explicit_enum_type"
	case ImplicitEnumTypeType:
		return "implicit_enum_type"
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
	case FuncDeclType:
		return "func_decl"
	case FuncDefType:
		return "func_def"
	case LexInternalTokensType:
		return "lex_internal_tokens"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_EndMarker      = SymbolId(0)
	_WildcardMarker = SymbolId(-1)

	LiteralType                            = SymbolId(327)
	AnonymousFuncExprType                  = SymbolId(328)
	AnonymousStructExprType                = SymbolId(329)
	AtomExprType                           = SymbolId(330)
	GenericArgumentsType                   = SymbolId(331)
	ArgumentsType                          = SymbolId(332)
	CallExprType                           = SymbolId(333)
	AccessExprType                         = SymbolId(334)
	UnaryOpType                            = SymbolId(335)
	UnaryExprType                          = SymbolId(336)
	MulOpType                              = SymbolId(337)
	MulExprType                            = SymbolId(338)
	AddOpType                              = SymbolId(339)
	AddExprType                            = SymbolId(340)
	CmpOpType                              = SymbolId(341)
	CmpExprType                            = SymbolId(342)
	AndExprType                            = SymbolId(343)
	OrExprType                             = SymbolId(344)
	SequenceExprType                       = SymbolId(345)
	OptionalJumpLabelType                  = SymbolId(346)
	OptionalExpressionOrImplicitStructType = SymbolId(347)
	JumpTypeType                           = SymbolId(348)
	OpOneAssignType                        = SymbolId(349)
	BinaryOpAssignType                     = SymbolId(350)
	ExpressionOrImplicitStructType         = SymbolId(351)
	StatementBodyType                      = SymbolId(352)
	StatementType                          = SymbolId(353)
	StatementsType                         = SymbolId(354)
	OptionalLabelDeclType                  = SymbolId(355)
	BlockBodyType                          = SymbolId(356)
	BlockExprType                          = SymbolId(357)
	IfExprType                             = SymbolId(358)
	MatchExprType                          = SymbolId(359)
	LoopExprType                           = SymbolId(360)
	ExpressionType                         = SymbolId(361)
	GenericParametersType                  = SymbolId(362)
	OptionalGenericParametersType          = SymbolId(363)
	AtomTypeType                           = SymbolId(364)
	ImplicitStructTypeType                 = SymbolId(365)
	ExplicitStructTypeType                 = SymbolId(366)
	ExplicitEnumTypeType                   = SymbolId(367)
	ImplicitEnumTypeType                   = SymbolId(368)
	ValueTypeType                          = SymbolId(369)
	TypeDeclType                           = SymbolId(370)
	ParameterDeclType                      = SymbolId(371)
	NonEmptyParametersType                 = SymbolId(372)
	VarargType                             = SymbolId(373)
	OptionalVarargType                     = SymbolId(374)
	ParametersType                         = SymbolId(375)
	ImplicitFuncTypeType                   = SymbolId(376)
	ExplicitFuncTypeType                   = SymbolId(377)
	OptionalReceiverType                   = SymbolId(378)
	FuncDeclType                           = SymbolId(379)
	FuncDefType                            = SymbolId(380)
	LexInternalTokensType                  = SymbolId(381)
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
	_ReduceBoolLiteralToLiteral                                           = _ReduceType(1)
	_ReduceIntegerLiteralToLiteral                                        = _ReduceType(2)
	_ReduceFloatLiteralToLiteral                                          = _ReduceType(3)
	_ReduceRuneLiteralToLiteral                                           = _ReduceType(4)
	_ReduceStringLiteralToLiteral                                         = _ReduceType(5)
	_ReduceToAnonymousFuncExpr                                            = _ReduceType(6)
	_ReduceExplicitToAnonymousStructExpr                                  = _ReduceType(7)
	_ReduceImplicitToAnonymousStructExpr                                  = _ReduceType(8)
	_ReduceLiteralToAtomExpr                                              = _ReduceType(9)
	_ReduceIdentifierToAtomExpr                                           = _ReduceType(10)
	_ReduceBlockExprToAtomExpr                                            = _ReduceType(11)
	_ReduceAnonymousFuncExprToAtomExpr                                    = _ReduceType(12)
	_ReduceAnonymousStructExprToAtomExpr                                  = _ReduceType(13)
	_ReduceLexErrorToAtomExpr                                             = _ReduceType(14)
	_ReduceToGenericArguments                                             = _ReduceType(15)
	_ReduceToArguments                                                    = _ReduceType(16)
	_ReduceConcreteToCallExpr                                             = _ReduceType(17)
	_ReduceGenericToCallExpr                                              = _ReduceType(18)
	_ReduceAtomExprToAccessExpr                                           = _ReduceType(19)
	_ReduceAccessToAccessExpr                                             = _ReduceType(20)
	_ReduceCallExprToAccessExpr                                           = _ReduceType(21)
	_ReduceIndexToAccessExpr                                              = _ReduceType(22)
	_ReduceNotToUnaryOp                                                   = _ReduceType(23)
	_ReduceBitNegToUnaryOp                                                = _ReduceType(24)
	_ReduceSubToUnaryOp                                                   = _ReduceType(25)
	_ReduceMulToUnaryOp                                                   = _ReduceType(26)
	_ReduceBitAndToUnaryOp                                                = _ReduceType(27)
	_ReduceAccessExprToUnaryExpr                                          = _ReduceType(28)
	_ReduceOpToUnaryExpr                                                  = _ReduceType(29)
	_ReduceMulToMulOp                                                     = _ReduceType(30)
	_ReduceDivToMulOp                                                     = _ReduceType(31)
	_ReduceModToMulOp                                                     = _ReduceType(32)
	_ReduceBitAndToMulOp                                                  = _ReduceType(33)
	_ReduceBitLshiftToMulOp                                               = _ReduceType(34)
	_ReduceBitRshiftToMulOp                                               = _ReduceType(35)
	_ReduceUnaryExprToMulExpr                                             = _ReduceType(36)
	_ReduceOpToMulExpr                                                    = _ReduceType(37)
	_ReduceAddToAddOp                                                     = _ReduceType(38)
	_ReduceSubToAddOp                                                     = _ReduceType(39)
	_ReduceBitOrToAddOp                                                   = _ReduceType(40)
	_ReduceBitXorToAddOp                                                  = _ReduceType(41)
	_ReduceMulExprToAddExpr                                               = _ReduceType(42)
	_ReduceOpToAddExpr                                                    = _ReduceType(43)
	_ReduceEqualToCmpOp                                                   = _ReduceType(44)
	_ReduceNotEqualToCmpOp                                                = _ReduceType(45)
	_ReduceLessToCmpOp                                                    = _ReduceType(46)
	_ReduceLessOrEqualToCmpOp                                             = _ReduceType(47)
	_ReduceGreaterToCmpOp                                                 = _ReduceType(48)
	_ReduceGreaterOrEqualToCmpOp                                          = _ReduceType(49)
	_ReduceAddExprToCmpExpr                                               = _ReduceType(50)
	_ReduceOpToCmpExpr                                                    = _ReduceType(51)
	_ReduceCmpExprToAndExpr                                               = _ReduceType(52)
	_ReduceOpToAndExpr                                                    = _ReduceType(53)
	_ReduceAndExprToOrExpr                                                = _ReduceType(54)
	_ReduceOpToOrExpr                                                     = _ReduceType(55)
	_ReduceToSequenceExpr                                                 = _ReduceType(56)
	_ReduceJumpLabelToOptionalJumpLabel                                   = _ReduceType(57)
	_ReduceUnlabelledToOptionalJumpLabel                                  = _ReduceType(58)
	_ReduceExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct = _ReduceType(59)
	_ReduceNilToOptionalExpressionOrImplicitStruct                        = _ReduceType(60)
	_ReduceReturnToJumpType                                               = _ReduceType(61)
	_ReduceBreakToJumpType                                                = _ReduceType(62)
	_ReduceContinueToJumpType                                             = _ReduceType(63)
	_ReduceAddOneAssignToOpOneAssign                                      = _ReduceType(64)
	_ReduceSubOneAssignToOpOneAssign                                      = _ReduceType(65)
	_ReduceAddAssignToBinaryOpAssign                                      = _ReduceType(66)
	_ReduceSubAssignToBinaryOpAssign                                      = _ReduceType(67)
	_ReduceMulAssignToBinaryOpAssign                                      = _ReduceType(68)
	_ReduceDivAssignToBinaryOpAssign                                      = _ReduceType(69)
	_ReduceModAssignToBinaryOpAssign                                      = _ReduceType(70)
	_ReduceBitNegAssignToBinaryOpAssign                                   = _ReduceType(71)
	_ReduceBitAndAssignToBinaryOpAssign                                   = _ReduceType(72)
	_ReduceBitOrAssignToBinaryOpAssign                                    = _ReduceType(73)
	_ReduceBitXorAssignToBinaryOpAssign                                   = _ReduceType(74)
	_ReduceBitLshiftAssignToBinaryOpAssign                                = _ReduceType(75)
	_ReduceBitRshiftAssignToBinaryOpAssign                                = _ReduceType(76)
	_ReduceExpressionToExpressionOrImplicitStruct                         = _ReduceType(77)
	_ReduceImplicitStructToExpressionOrImplicitStruct                     = _ReduceType(78)
	_ReduceExpressionOrImplicitStructToStatementBody                      = _ReduceType(79)
	_ReduceAsyncToStatementBody                                           = _ReduceType(80)
	_ReduceJumpToStatementBody                                            = _ReduceType(81)
	_ReduceOpOneAssignToStatementBody                                     = _ReduceType(82)
	_ReduceBinaryOpAssignToStatementBody                                  = _ReduceType(83)
	_ReduceImplicitToStatement                                            = _ReduceType(84)
	_ReduceExplicitToStatement                                            = _ReduceType(85)
	_ReduceEmptyListToStatements                                          = _ReduceType(86)
	_ReduceAddToStatements                                                = _ReduceType(87)
	_ReduceLabelDeclToOptionalLabelDecl                                   = _ReduceType(88)
	_ReduceUnlabelledToOptionalLabelDecl                                  = _ReduceType(89)
	_ReduceToBlockBody                                                    = _ReduceType(90)
	_ReduceToBlockExpr                                                    = _ReduceType(91)
	_ReduceNoElseToIfExpr                                                 = _ReduceType(92)
	_ReduceIfElseToIfExpr                                                 = _ReduceType(93)
	_ReduceMultiIfElseToIfExpr                                            = _ReduceType(94)
	_ReduceToMatchExpr                                                    = _ReduceType(95)
	_ReduceInfiniteToLoopExpr                                             = _ReduceType(96)
	_ReduceWhileToLoopExpr                                                = _ReduceType(97)
	_ReduceSequenceExprToExpression                                       = _ReduceType(98)
	_ReduceIfExprToExpression                                             = _ReduceType(99)
	_ReduceMatchExprToExpression                                          = _ReduceType(100)
	_ReduceLoopExprToExpression                                           = _ReduceType(101)
	_ReduceToGenericParameters                                            = _ReduceType(102)
	_ReduceGenericToOptionalGenericParameters                             = _ReduceType(103)
	_ReduceNilToOptionalGenericParameters                                 = _ReduceType(104)
	_ReduceNamedToAtomType                                                = _ReduceType(105)
	_ReduceExplicitStructTypeToAtomType                                   = _ReduceType(106)
	_ReduceImplicitStructTypeToAtomType                                   = _ReduceType(107)
	_ReduceExplicitEnumTypeToAtomType                                     = _ReduceType(108)
	_ReduceToImplicitStructType                                           = _ReduceType(109)
	_ReduceToExplicitStructType                                           = _ReduceType(110)
	_ReduceToExplicitEnumType                                             = _ReduceType(111)
	_ReducePairToImplicitEnumType                                         = _ReduceType(112)
	_ReduceAddToImplicitEnumType                                          = _ReduceType(113)
	_ReduceAtomTypeToValueType                                            = _ReduceType(114)
	_ReduceImplicitEnumTypeToValueType                                    = _ReduceType(115)
	_ReduceDefinitionToTypeDecl                                           = _ReduceType(116)
	_ReduceAliasToTypeDecl                                                = _ReduceType(117)
	_ReduceExplicitToParameterDecl                                        = _ReduceType(118)
	_ReduceImplicitToParameterDecl                                        = _ReduceType(119)
	_ReduceInferredToParameterDecl                                        = _ReduceType(120)
	_ReduceParameterDeclToNonEmptyParameters                              = _ReduceType(121)
	_ReduceAddToNonEmptyParameters                                        = _ReduceType(122)
	_ReduceExplicitToVararg                                               = _ReduceType(123)
	_ReduceImplicitToVararg                                               = _ReduceType(124)
	_ReduceInferredToVararg                                               = _ReduceType(125)
	_ReduceVarargToOptionalVararg                                         = _ReduceType(126)
	_ReduceVararg2ToOptionalVararg                                        = _ReduceType(127)
	_ReduceNilToOptionalVararg                                            = _ReduceType(128)
	_ReduceNonEmptyParametersToParameters                                 = _ReduceType(129)
	_ReduceMixedToParameters                                              = _ReduceType(130)
	_ReduceOptionalVarargToParameters                                     = _ReduceType(131)
	_ReduceTypedToImplicitFuncType                                        = _ReduceType(132)
	_ReduceInferredToImplicitFuncType                                     = _ReduceType(133)
	_ReduceImplicitUnitToImplicitFuncType                                 = _ReduceType(134)
	_ReduceToExplicitFuncType                                             = _ReduceType(135)
	_ReduceReceiverToOptionalReceiver                                     = _ReduceType(136)
	_ReduceNilToOptionalReceiver                                          = _ReduceType(137)
	_ReduceToFuncDecl                                                     = _ReduceType(138)
	_ReduceToFuncDef                                                      = _ReduceType(139)
	_ReduceSpacesToLexInternalTokens                                      = _ReduceType(140)
	_ReduceCommentToLexInternalTokens                                     = _ReduceType(141)
)

func (i _ReduceType) String() string {
	switch i {
	case _ReduceBoolLiteralToLiteral:
		return "BoolLiteralToLiteral"
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
	case _ReduceToMatchExpr:
		return "ToMatchExpr"
	case _ReduceInfiniteToLoopExpr:
		return "InfiniteToLoopExpr"
	case _ReduceWhileToLoopExpr:
		return "WhileToLoopExpr"
	case _ReduceSequenceExprToExpression:
		return "SequenceExprToExpression"
	case _ReduceIfExprToExpression:
		return "IfExprToExpression"
	case _ReduceMatchExprToExpression:
		return "MatchExprToExpression"
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
	case _ReduceToImplicitStructType:
		return "ToImplicitStructType"
	case _ReduceToExplicitStructType:
		return "ToExplicitStructType"
	case _ReduceToExplicitEnumType:
		return "ToExplicitEnumType"
	case _ReducePairToImplicitEnumType:
		return "PairToImplicitEnumType"
	case _ReduceAddToImplicitEnumType:
		return "AddToImplicitEnumType"
	case _ReduceAtomTypeToValueType:
		return "AtomTypeToValueType"
	case _ReduceImplicitEnumTypeToValueType:
		return "ImplicitEnumTypeToValueType"
	case _ReduceDefinitionToTypeDecl:
		return "DefinitionToTypeDecl"
	case _ReduceAliasToTypeDecl:
		return "AliasToTypeDecl"
	case _ReduceExplicitToParameterDecl:
		return "ExplicitToParameterDecl"
	case _ReduceImplicitToParameterDecl:
		return "ImplicitToParameterDecl"
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
	case _EndMarker, SpacesToken, NewlinesToken, CommentToken, BoolLiteralToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, GenericLbraceToken, DotToken, CommaToken, QuestionToken, SemicolonToken, DotdotdotToken, IdentifierToken, IfToken, ElseToken, MatchToken, CaseToken, DefaultToken, ForToken, ReturnToken, BreakToken, ContinueToken, TypeToken, StructToken, EnumToken, FuncToken, AsyncToken, LabelDeclToken, JumpLabelToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, NotToken, AndToken, OrToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, LexErrorToken:
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
	case _ReduceBoolLiteralToLiteral:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralType
		symbol.Generic_, err = reducer.BoolLiteralToLiteral(args[0].Generic_)
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
	case _ReduceToMatchExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = MatchExprType
		symbol.Generic_, err = reducer.ToMatchExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	case _ReduceMatchExprToExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.MatchExprToExpression(args[0].Generic_, args[1].Generic_)
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
	case _ReduceToExplicitEnumType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumTypeType
		symbol.Generic_, err = reducer.ToExplicitEnumType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReducePairToImplicitEnumType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitEnumTypeType
		symbol.Generic_, err = reducer.PairToImplicitEnumType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceAddToImplicitEnumType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitEnumTypeType
		symbol.Generic_, err = reducer.AddToImplicitEnumType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceAtomTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.AtomTypeToValueType(args[0].Generic_)
	case _ReduceImplicitEnumTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.ImplicitEnumTypeToValueType(args[0].Generic_)
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
	case _ReduceExplicitToParameterDecl:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.ExplicitToParameterDecl(args[0].Generic_, args[1].Generic_)
	case _ReduceImplicitToParameterDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.ImplicitToParameterDecl(args[0].Generic_)
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
	_ReduceBoolLiteralToLiteralAction                                           = &_Action{_ReduceAction, 0, _ReduceBoolLiteralToLiteral}
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
	_ReduceToMatchExprAction                                                    = &_Action{_ReduceAction, 0, _ReduceToMatchExpr}
	_ReduceInfiniteToLoopExprAction                                             = &_Action{_ReduceAction, 0, _ReduceInfiniteToLoopExpr}
	_ReduceWhileToLoopExprAction                                                = &_Action{_ReduceAction, 0, _ReduceWhileToLoopExpr}
	_ReduceSequenceExprToExpressionAction                                       = &_Action{_ReduceAction, 0, _ReduceSequenceExprToExpression}
	_ReduceIfExprToExpressionAction                                             = &_Action{_ReduceAction, 0, _ReduceIfExprToExpression}
	_ReduceMatchExprToExpressionAction                                          = &_Action{_ReduceAction, 0, _ReduceMatchExprToExpression}
	_ReduceLoopExprToExpressionAction                                           = &_Action{_ReduceAction, 0, _ReduceLoopExprToExpression}
	_ReduceToGenericParametersAction                                            = &_Action{_ReduceAction, 0, _ReduceToGenericParameters}
	_ReduceGenericToOptionalGenericParametersAction                             = &_Action{_ReduceAction, 0, _ReduceGenericToOptionalGenericParameters}
	_ReduceNilToOptionalGenericParametersAction                                 = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericParameters}
	_ReduceNamedToAtomTypeAction                                                = &_Action{_ReduceAction, 0, _ReduceNamedToAtomType}
	_ReduceExplicitStructTypeToAtomTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceExplicitStructTypeToAtomType}
	_ReduceImplicitStructTypeToAtomTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceImplicitStructTypeToAtomType}
	_ReduceExplicitEnumTypeToAtomTypeAction                                     = &_Action{_ReduceAction, 0, _ReduceExplicitEnumTypeToAtomType}
	_ReduceToImplicitStructTypeAction                                           = &_Action{_ReduceAction, 0, _ReduceToImplicitStructType}
	_ReduceToExplicitStructTypeAction                                           = &_Action{_ReduceAction, 0, _ReduceToExplicitStructType}
	_ReduceToExplicitEnumTypeAction                                             = &_Action{_ReduceAction, 0, _ReduceToExplicitEnumType}
	_ReducePairToImplicitEnumTypeAction                                         = &_Action{_ReduceAction, 0, _ReducePairToImplicitEnumType}
	_ReduceAddToImplicitEnumTypeAction                                          = &_Action{_ReduceAction, 0, _ReduceAddToImplicitEnumType}
	_ReduceAtomTypeToValueTypeAction                                            = &_Action{_ReduceAction, 0, _ReduceAtomTypeToValueType}
	_ReduceImplicitEnumTypeToValueTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceImplicitEnumTypeToValueType}
	_ReduceDefinitionToTypeDeclAction                                           = &_Action{_ReduceAction, 0, _ReduceDefinitionToTypeDecl}
	_ReduceAliasToTypeDeclAction                                                = &_Action{_ReduceAction, 0, _ReduceAliasToTypeDecl}
	_ReduceExplicitToParameterDeclAction                                        = &_Action{_ReduceAction, 0, _ReduceExplicitToParameterDecl}
	_ReduceImplicitToParameterDeclAction                                        = &_Action{_ReduceAction, 0, _ReduceImplicitToParameterDecl}
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
	_ReduceSpacesToLexInternalTokensAction                                      = &_Action{_ReduceAction, 0, _ReduceSpacesToLexInternalTokens}
	_ReduceCommentToLexInternalTokensAction                                     = &_Action{_ReduceAction, 0, _ReduceCommentToLexInternalTokens}
)

var _ActionTable = _ActionTableType{
	{_State5, _EndMarker}:                               &_Action{_AcceptAction, 0, 0},
	{_State6, _EndMarker}:                               &_Action{_AcceptAction, 0, 0},
	{_State7, _EndMarker}:                               &_Action{_AcceptAction, 0, 0},
	{_State8, _EndMarker}:                               &_Action{_AcceptAction, 0, 0},
	{_State1, TypeToken}:                                _GotoState9Action,
	{_State1, TypeDeclType}:                             _GotoState5Action,
	{_State2, FuncToken}:                                _GotoState10Action,
	{_State2, FuncDeclType}:                             _GotoState11Action,
	{_State2, FuncDefType}:                              _GotoState6Action,
	{_State3, BoolLiteralToken}:                         _GotoState14Action,
	{_State3, IntegerLiteralToken}:                      _GotoState18Action,
	{_State3, FloatLiteralToken}:                        _GotoState15Action,
	{_State3, RuneLiteralToken}:                         _GotoState24Action,
	{_State3, StringLiteralToken}:                       _GotoState25Action,
	{_State3, LparenToken}:                              _GotoState21Action,
	{_State3, IdentifierToken}:                          _GotoState17Action,
	{_State3, StructToken}:                              _GotoState26Action,
	{_State3, FuncToken}:                                _GotoState16Action,
	{_State3, LabelDeclToken}:                           _GotoState19Action,
	{_State3, NotToken}:                                 _GotoState23Action,
	{_State3, SubToken}:                                 _GotoState27Action,
	{_State3, MulToken}:                                 _GotoState22Action,
	{_State3, BitNegToken}:                              _GotoState13Action,
	{_State3, BitAndToken}:                              _GotoState12Action,
	{_State3, LexErrorToken}:                            _GotoState20Action,
	{_State3, LiteralType}:                              _GotoState39Action,
	{_State3, AnonymousFuncExprType}:                    _GotoState31Action,
	{_State3, AnonymousStructExprType}:                  _GotoState32Action,
	{_State3, AtomExprType}:                             _GotoState33Action,
	{_State3, CallExprType}:                             _GotoState35Action,
	{_State3, AccessExprType}:                           _GotoState28Action,
	{_State3, UnaryOpType}:                              _GotoState45Action,
	{_State3, UnaryExprType}:                            _GotoState44Action,
	{_State3, MulExprType}:                              _GotoState40Action,
	{_State3, AddExprType}:                              _GotoState29Action,
	{_State3, CmpExprType}:                              _GotoState36Action,
	{_State3, AndExprType}:                              _GotoState30Action,
	{_State3, OrExprType}:                               _GotoState42Action,
	{_State3, SequenceExprType}:                         _GotoState43Action,
	{_State3, OptionalLabelDeclType}:                    _GotoState41Action,
	{_State3, BlockExprType}:                            _GotoState34Action,
	{_State3, ExpressionType}:                           _GotoState7Action,
	{_State3, ExplicitStructTypeType}:                   _GotoState38Action,
	{_State3, ExplicitFuncTypeType}:                     _GotoState37Action,
	{_State4, SpacesToken}:                              _GotoState47Action,
	{_State4, CommentToken}:                             _GotoState46Action,
	{_State4, LexInternalTokensType}:                    _GotoState8Action,
	{_State9, IdentifierToken}:                          _GotoState48Action,
	{_State10, LparenToken}:                             _GotoState49Action,
	{_State10, OptionalReceiverType}:                    _GotoState50Action,
	{_State11, LbraceToken}:                             _GotoState51Action,
	{_State11, BlockBodyType}:                           _GotoState52Action,
	{_State16, LparenToken}:                             _GotoState53Action,
	{_State16, ImplicitFuncTypeType}:                    _GotoState54Action,
	{_State21, ArgumentsType}:                           _GotoState55Action,
	{_State26, LparenToken}:                             _GotoState56Action,
	{_State26, ImplicitStructTypeType}:                  _GotoState57Action,
	{_State28, LparenToken}:                             _GotoState61Action,
	{_State28, LbracketToken}:                           _GotoState60Action,
	{_State28, GenericLbraceToken}:                      _GotoState59Action,
	{_State28, DotToken}:                                _GotoState58Action,
	{_State29, AddToken}:                                _GotoState62Action,
	{_State29, SubToken}:                                _GotoState65Action,
	{_State29, BitXorToken}:                             _GotoState64Action,
	{_State29, BitOrToken}:                              _GotoState63Action,
	{_State29, AddOpType}:                               _GotoState66Action,
	{_State30, AndToken}:                                _GotoState67Action,
	{_State36, EqualToken}:                              _GotoState68Action,
	{_State36, NotEqualToken}:                           _GotoState73Action,
	{_State36, LessToken}:                               _GotoState71Action,
	{_State36, LessOrEqualToken}:                        _GotoState72Action,
	{_State36, GreaterToken}:                            _GotoState69Action,
	{_State36, GreaterOrEqualToken}:                     _GotoState70Action,
	{_State36, CmpOpType}:                               _GotoState74Action,
	{_State37, LbraceToken}:                             _GotoState51Action,
	{_State37, BlockBodyType}:                           _GotoState75Action,
	{_State38, LparenToken}:                             _GotoState76Action,
	{_State40, MulToken}:                                _GotoState82Action,
	{_State40, DivToken}:                                _GotoState80Action,
	{_State40, ModToken}:                                _GotoState81Action,
	{_State40, BitAndToken}:                             _GotoState77Action,
	{_State40, BitLshiftToken}:                          _GotoState78Action,
	{_State40, BitRshiftToken}:                          _GotoState79Action,
	{_State40, MulOpType}:                               _GotoState83Action,
	{_State41, LbraceToken}:                             _GotoState51Action,
	{_State41, IfToken}:                                 _GotoState85Action,
	{_State41, MatchToken}:                              _GotoState86Action,
	{_State41, ForToken}:                                _GotoState84Action,
	{_State41, BlockBodyType}:                           _GotoState87Action,
	{_State41, IfExprType}:                              _GotoState88Action,
	{_State41, MatchExprType}:                           _GotoState90Action,
	{_State41, LoopExprType}:                            _GotoState89Action,
	{_State42, OrToken}:                                 _GotoState91Action,
	{_State45, BoolLiteralToken}:                        _GotoState14Action,
	{_State45, IntegerLiteralToken}:                     _GotoState18Action,
	{_State45, FloatLiteralToken}:                       _GotoState15Action,
	{_State45, RuneLiteralToken}:                        _GotoState24Action,
	{_State45, StringLiteralToken}:                      _GotoState25Action,
	{_State45, LparenToken}:                             _GotoState21Action,
	{_State45, IdentifierToken}:                         _GotoState17Action,
	{_State45, StructToken}:                             _GotoState26Action,
	{_State45, FuncToken}:                               _GotoState16Action,
	{_State45, LabelDeclToken}:                          _GotoState19Action,
	{_State45, NotToken}:                                _GotoState23Action,
	{_State45, SubToken}:                                _GotoState27Action,
	{_State45, MulToken}:                                _GotoState22Action,
	{_State45, BitNegToken}:                             _GotoState13Action,
	{_State45, BitAndToken}:                             _GotoState12Action,
	{_State45, LexErrorToken}:                           _GotoState20Action,
	{_State45, LiteralType}:                             _GotoState39Action,
	{_State45, AnonymousFuncExprType}:                   _GotoState31Action,
	{_State45, AnonymousStructExprType}:                 _GotoState32Action,
	{_State45, AtomExprType}:                            _GotoState33Action,
	{_State45, CallExprType}:                            _GotoState35Action,
	{_State45, AccessExprType}:                          _GotoState28Action,
	{_State45, UnaryOpType}:                             _GotoState45Action,
	{_State45, UnaryExprType}:                           _GotoState93Action,
	{_State45, OptionalLabelDeclType}:                   _GotoState92Action,
	{_State45, BlockExprType}:                           _GotoState34Action,
	{_State45, ExplicitStructTypeType}:                  _GotoState38Action,
	{_State45, ExplicitFuncTypeType}:                    _GotoState37Action,
	{_State48, GenericLbraceToken}:                      _GotoState95Action,
	{_State48, EqualToken}:                              _GotoState94Action,
	{_State48, OptionalGenericParametersType}:           _GotoState96Action,
	{_State49, LparenToken}:                             _GotoState56Action,
	{_State49, IdentifierToken}:                         _GotoState98Action,
	{_State49, StructToken}:                             _GotoState26Action,
	{_State49, EnumToken}:                               _GotoState97Action,
	{_State49, AtomTypeType}:                            _GotoState99Action,
	{_State49, ImplicitStructTypeType}:                  _GotoState103Action,
	{_State49, ExplicitStructTypeType}:                  _GotoState101Action,
	{_State49, ExplicitEnumTypeType}:                    _GotoState100Action,
	{_State49, ImplicitEnumTypeType}:                    _GotoState102Action,
	{_State49, ValueTypeType}:                           _GotoState105Action,
	{_State49, ParameterDeclType}:                       _GotoState104Action,
	{_State50, IdentifierToken}:                         _GotoState106Action,
	{_State51, StatementsType}:                          _GotoState107Action,
	{_State53, LparenToken}:                             _GotoState56Action,
	{_State53, DotdotdotToken}:                          _GotoState108Action,
	{_State53, IdentifierToken}:                         _GotoState109Action,
	{_State53, StructToken}:                             _GotoState26Action,
	{_State53, EnumToken}:                               _GotoState97Action,
	{_State53, AtomTypeType}:                            _GotoState99Action,
	{_State53, ImplicitStructTypeType}:                  _GotoState103Action,
	{_State53, ExplicitStructTypeType}:                  _GotoState101Action,
	{_State53, ExplicitEnumTypeType}:                    _GotoState100Action,
	{_State53, ImplicitEnumTypeType}:                    _GotoState102Action,
	{_State53, ValueTypeType}:                           _GotoState105Action,
	{_State53, ParameterDeclType}:                       _GotoState112Action,
	{_State53, NonEmptyParametersType}:                  _GotoState110Action,
	{_State53, VarargType}:                              _GotoState114Action,
	{_State53, OptionalVarargType}:                      _GotoState111Action,
	{_State53, ParametersType}:                          _GotoState113Action,
	{_State55, RparenToken}:                             _GotoState115Action,
	{_State56, RparenToken}:                             _GotoState116Action,
	{_State58, IdentifierToken}:                         _GotoState117Action,
	{_State59, GenericArgumentsType}:                    _GotoState118Action,
	{_State60, BoolLiteralToken}:                        _GotoState14Action,
	{_State60, IntegerLiteralToken}:                     _GotoState18Action,
	{_State60, FloatLiteralToken}:                       _GotoState15Action,
	{_State60, RuneLiteralToken}:                        _GotoState24Action,
	{_State60, StringLiteralToken}:                      _GotoState25Action,
	{_State60, LparenToken}:                             _GotoState21Action,
	{_State60, IdentifierToken}:                         _GotoState17Action,
	{_State60, StructToken}:                             _GotoState26Action,
	{_State60, FuncToken}:                               _GotoState16Action,
	{_State60, LabelDeclToken}:                          _GotoState19Action,
	{_State60, NotToken}:                                _GotoState23Action,
	{_State60, SubToken}:                                _GotoState27Action,
	{_State60, MulToken}:                                _GotoState22Action,
	{_State60, BitNegToken}:                             _GotoState13Action,
	{_State60, BitAndToken}:                             _GotoState12Action,
	{_State60, LexErrorToken}:                           _GotoState20Action,
	{_State60, LiteralType}:                             _GotoState39Action,
	{_State60, AnonymousFuncExprType}:                   _GotoState31Action,
	{_State60, AnonymousStructExprType}:                 _GotoState32Action,
	{_State60, AtomExprType}:                            _GotoState33Action,
	{_State60, CallExprType}:                            _GotoState35Action,
	{_State60, AccessExprType}:                          _GotoState28Action,
	{_State60, UnaryOpType}:                             _GotoState45Action,
	{_State60, UnaryExprType}:                           _GotoState44Action,
	{_State60, MulExprType}:                             _GotoState40Action,
	{_State60, AddExprType}:                             _GotoState29Action,
	{_State60, CmpExprType}:                             _GotoState36Action,
	{_State60, AndExprType}:                             _GotoState30Action,
	{_State60, OrExprType}:                              _GotoState42Action,
	{_State60, SequenceExprType}:                        _GotoState43Action,
	{_State60, OptionalLabelDeclType}:                   _GotoState41Action,
	{_State60, BlockExprType}:                           _GotoState34Action,
	{_State60, ExpressionType}:                          _GotoState119Action,
	{_State60, ExplicitStructTypeType}:                  _GotoState38Action,
	{_State60, ExplicitFuncTypeType}:                    _GotoState37Action,
	{_State61, ArgumentsType}:                           _GotoState120Action,
	{_State66, BoolLiteralToken}:                        _GotoState14Action,
	{_State66, IntegerLiteralToken}:                     _GotoState18Action,
	{_State66, FloatLiteralToken}:                       _GotoState15Action,
	{_State66, RuneLiteralToken}:                        _GotoState24Action,
	{_State66, StringLiteralToken}:                      _GotoState25Action,
	{_State66, LparenToken}:                             _GotoState21Action,
	{_State66, IdentifierToken}:                         _GotoState17Action,
	{_State66, StructToken}:                             _GotoState26Action,
	{_State66, FuncToken}:                               _GotoState16Action,
	{_State66, LabelDeclToken}:                          _GotoState19Action,
	{_State66, NotToken}:                                _GotoState23Action,
	{_State66, SubToken}:                                _GotoState27Action,
	{_State66, MulToken}:                                _GotoState22Action,
	{_State66, BitNegToken}:                             _GotoState13Action,
	{_State66, BitAndToken}:                             _GotoState12Action,
	{_State66, LexErrorToken}:                           _GotoState20Action,
	{_State66, LiteralType}:                             _GotoState39Action,
	{_State66, AnonymousFuncExprType}:                   _GotoState31Action,
	{_State66, AnonymousStructExprType}:                 _GotoState32Action,
	{_State66, AtomExprType}:                            _GotoState33Action,
	{_State66, CallExprType}:                            _GotoState35Action,
	{_State66, AccessExprType}:                          _GotoState28Action,
	{_State66, UnaryOpType}:                             _GotoState45Action,
	{_State66, UnaryExprType}:                           _GotoState44Action,
	{_State66, MulExprType}:                             _GotoState121Action,
	{_State66, OptionalLabelDeclType}:                   _GotoState92Action,
	{_State66, BlockExprType}:                           _GotoState34Action,
	{_State66, ExplicitStructTypeType}:                  _GotoState38Action,
	{_State66, ExplicitFuncTypeType}:                    _GotoState37Action,
	{_State67, BoolLiteralToken}:                        _GotoState14Action,
	{_State67, IntegerLiteralToken}:                     _GotoState18Action,
	{_State67, FloatLiteralToken}:                       _GotoState15Action,
	{_State67, RuneLiteralToken}:                        _GotoState24Action,
	{_State67, StringLiteralToken}:                      _GotoState25Action,
	{_State67, LparenToken}:                             _GotoState21Action,
	{_State67, IdentifierToken}:                         _GotoState17Action,
	{_State67, StructToken}:                             _GotoState26Action,
	{_State67, FuncToken}:                               _GotoState16Action,
	{_State67, LabelDeclToken}:                          _GotoState19Action,
	{_State67, NotToken}:                                _GotoState23Action,
	{_State67, SubToken}:                                _GotoState27Action,
	{_State67, MulToken}:                                _GotoState22Action,
	{_State67, BitNegToken}:                             _GotoState13Action,
	{_State67, BitAndToken}:                             _GotoState12Action,
	{_State67, LexErrorToken}:                           _GotoState20Action,
	{_State67, LiteralType}:                             _GotoState39Action,
	{_State67, AnonymousFuncExprType}:                   _GotoState31Action,
	{_State67, AnonymousStructExprType}:                 _GotoState32Action,
	{_State67, AtomExprType}:                            _GotoState33Action,
	{_State67, CallExprType}:                            _GotoState35Action,
	{_State67, AccessExprType}:                          _GotoState28Action,
	{_State67, UnaryOpType}:                             _GotoState45Action,
	{_State67, UnaryExprType}:                           _GotoState44Action,
	{_State67, MulExprType}:                             _GotoState40Action,
	{_State67, AddExprType}:                             _GotoState29Action,
	{_State67, CmpExprType}:                             _GotoState122Action,
	{_State67, OptionalLabelDeclType}:                   _GotoState92Action,
	{_State67, BlockExprType}:                           _GotoState34Action,
	{_State67, ExplicitStructTypeType}:                  _GotoState38Action,
	{_State67, ExplicitFuncTypeType}:                    _GotoState37Action,
	{_State74, BoolLiteralToken}:                        _GotoState14Action,
	{_State74, IntegerLiteralToken}:                     _GotoState18Action,
	{_State74, FloatLiteralToken}:                       _GotoState15Action,
	{_State74, RuneLiteralToken}:                        _GotoState24Action,
	{_State74, StringLiteralToken}:                      _GotoState25Action,
	{_State74, LparenToken}:                             _GotoState21Action,
	{_State74, IdentifierToken}:                         _GotoState17Action,
	{_State74, StructToken}:                             _GotoState26Action,
	{_State74, FuncToken}:                               _GotoState16Action,
	{_State74, LabelDeclToken}:                          _GotoState19Action,
	{_State74, NotToken}:                                _GotoState23Action,
	{_State74, SubToken}:                                _GotoState27Action,
	{_State74, MulToken}:                                _GotoState22Action,
	{_State74, BitNegToken}:                             _GotoState13Action,
	{_State74, BitAndToken}:                             _GotoState12Action,
	{_State74, LexErrorToken}:                           _GotoState20Action,
	{_State74, LiteralType}:                             _GotoState39Action,
	{_State74, AnonymousFuncExprType}:                   _GotoState31Action,
	{_State74, AnonymousStructExprType}:                 _GotoState32Action,
	{_State74, AtomExprType}:                            _GotoState33Action,
	{_State74, CallExprType}:                            _GotoState35Action,
	{_State74, AccessExprType}:                          _GotoState28Action,
	{_State74, UnaryOpType}:                             _GotoState45Action,
	{_State74, UnaryExprType}:                           _GotoState44Action,
	{_State74, MulExprType}:                             _GotoState40Action,
	{_State74, AddExprType}:                             _GotoState123Action,
	{_State74, OptionalLabelDeclType}:                   _GotoState92Action,
	{_State74, BlockExprType}:                           _GotoState34Action,
	{_State74, ExplicitStructTypeType}:                  _GotoState38Action,
	{_State74, ExplicitFuncTypeType}:                    _GotoState37Action,
	{_State76, ArgumentsType}:                           _GotoState124Action,
	{_State83, BoolLiteralToken}:                        _GotoState14Action,
	{_State83, IntegerLiteralToken}:                     _GotoState18Action,
	{_State83, FloatLiteralToken}:                       _GotoState15Action,
	{_State83, RuneLiteralToken}:                        _GotoState24Action,
	{_State83, StringLiteralToken}:                      _GotoState25Action,
	{_State83, LparenToken}:                             _GotoState21Action,
	{_State83, IdentifierToken}:                         _GotoState17Action,
	{_State83, StructToken}:                             _GotoState26Action,
	{_State83, FuncToken}:                               _GotoState16Action,
	{_State83, LabelDeclToken}:                          _GotoState19Action,
	{_State83, NotToken}:                                _GotoState23Action,
	{_State83, SubToken}:                                _GotoState27Action,
	{_State83, MulToken}:                                _GotoState22Action,
	{_State83, BitNegToken}:                             _GotoState13Action,
	{_State83, BitAndToken}:                             _GotoState12Action,
	{_State83, LexErrorToken}:                           _GotoState20Action,
	{_State83, LiteralType}:                             _GotoState39Action,
	{_State83, AnonymousFuncExprType}:                   _GotoState31Action,
	{_State83, AnonymousStructExprType}:                 _GotoState32Action,
	{_State83, AtomExprType}:                            _GotoState33Action,
	{_State83, CallExprType}:                            _GotoState35Action,
	{_State83, AccessExprType}:                          _GotoState28Action,
	{_State83, UnaryOpType}:                             _GotoState45Action,
	{_State83, UnaryExprType}:                           _GotoState125Action,
	{_State83, OptionalLabelDeclType}:                   _GotoState92Action,
	{_State83, BlockExprType}:                           _GotoState34Action,
	{_State83, ExplicitStructTypeType}:                  _GotoState38Action,
	{_State83, ExplicitFuncTypeType}:                    _GotoState37Action,
	{_State84, BoolLiteralToken}:                        _GotoState14Action,
	{_State84, IntegerLiteralToken}:                     _GotoState18Action,
	{_State84, FloatLiteralToken}:                       _GotoState15Action,
	{_State84, RuneLiteralToken}:                        _GotoState24Action,
	{_State84, StringLiteralToken}:                      _GotoState25Action,
	{_State84, LparenToken}:                             _GotoState21Action,
	{_State84, IdentifierToken}:                         _GotoState17Action,
	{_State84, StructToken}:                             _GotoState26Action,
	{_State84, FuncToken}:                               _GotoState16Action,
	{_State84, LabelDeclToken}:                          _GotoState19Action,
	{_State84, NotToken}:                                _GotoState23Action,
	{_State84, SubToken}:                                _GotoState27Action,
	{_State84, MulToken}:                                _GotoState22Action,
	{_State84, BitNegToken}:                             _GotoState13Action,
	{_State84, BitAndToken}:                             _GotoState12Action,
	{_State84, LexErrorToken}:                           _GotoState20Action,
	{_State84, LiteralType}:                             _GotoState39Action,
	{_State84, AnonymousFuncExprType}:                   _GotoState31Action,
	{_State84, AnonymousStructExprType}:                 _GotoState32Action,
	{_State84, AtomExprType}:                            _GotoState33Action,
	{_State84, CallExprType}:                            _GotoState35Action,
	{_State84, AccessExprType}:                          _GotoState28Action,
	{_State84, UnaryOpType}:                             _GotoState45Action,
	{_State84, UnaryExprType}:                           _GotoState44Action,
	{_State84, MulExprType}:                             _GotoState40Action,
	{_State84, AddExprType}:                             _GotoState29Action,
	{_State84, CmpExprType}:                             _GotoState36Action,
	{_State84, AndExprType}:                             _GotoState30Action,
	{_State84, OrExprType}:                              _GotoState42Action,
	{_State84, SequenceExprType}:                        _GotoState127Action,
	{_State84, OptionalLabelDeclType}:                   _GotoState92Action,
	{_State84, BlockExprType}:                           _GotoState126Action,
	{_State84, ExplicitStructTypeType}:                  _GotoState38Action,
	{_State84, ExplicitFuncTypeType}:                    _GotoState37Action,
	{_State85, BoolLiteralToken}:                        _GotoState14Action,
	{_State85, IntegerLiteralToken}:                     _GotoState18Action,
	{_State85, FloatLiteralToken}:                       _GotoState15Action,
	{_State85, RuneLiteralToken}:                        _GotoState24Action,
	{_State85, StringLiteralToken}:                      _GotoState25Action,
	{_State85, LparenToken}:                             _GotoState21Action,
	{_State85, IdentifierToken}:                         _GotoState17Action,
	{_State85, StructToken}:                             _GotoState26Action,
	{_State85, FuncToken}:                               _GotoState16Action,
	{_State85, LabelDeclToken}:                          _GotoState19Action,
	{_State85, NotToken}:                                _GotoState23Action,
	{_State85, SubToken}:                                _GotoState27Action,
	{_State85, MulToken}:                                _GotoState22Action,
	{_State85, BitNegToken}:                             _GotoState13Action,
	{_State85, BitAndToken}:                             _GotoState12Action,
	{_State85, LexErrorToken}:                           _GotoState20Action,
	{_State85, LiteralType}:                             _GotoState39Action,
	{_State85, AnonymousFuncExprType}:                   _GotoState31Action,
	{_State85, AnonymousStructExprType}:                 _GotoState32Action,
	{_State85, AtomExprType}:                            _GotoState33Action,
	{_State85, CallExprType}:                            _GotoState35Action,
	{_State85, AccessExprType}:                          _GotoState28Action,
	{_State85, UnaryOpType}:                             _GotoState45Action,
	{_State85, UnaryExprType}:                           _GotoState44Action,
	{_State85, MulExprType}:                             _GotoState40Action,
	{_State85, AddExprType}:                             _GotoState29Action,
	{_State85, CmpExprType}:                             _GotoState36Action,
	{_State85, AndExprType}:                             _GotoState30Action,
	{_State85, OrExprType}:                              _GotoState42Action,
	{_State85, SequenceExprType}:                        _GotoState128Action,
	{_State85, OptionalLabelDeclType}:                   _GotoState92Action,
	{_State85, BlockExprType}:                           _GotoState34Action,
	{_State85, ExplicitStructTypeType}:                  _GotoState38Action,
	{_State85, ExplicitFuncTypeType}:                    _GotoState37Action,
	{_State86, CaseToken}:                               _GotoState129Action,
	{_State91, BoolLiteralToken}:                        _GotoState14Action,
	{_State91, IntegerLiteralToken}:                     _GotoState18Action,
	{_State91, FloatLiteralToken}:                       _GotoState15Action,
	{_State91, RuneLiteralToken}:                        _GotoState24Action,
	{_State91, StringLiteralToken}:                      _GotoState25Action,
	{_State91, LparenToken}:                             _GotoState21Action,
	{_State91, IdentifierToken}:                         _GotoState17Action,
	{_State91, StructToken}:                             _GotoState26Action,
	{_State91, FuncToken}:                               _GotoState16Action,
	{_State91, LabelDeclToken}:                          _GotoState19Action,
	{_State91, NotToken}:                                _GotoState23Action,
	{_State91, SubToken}:                                _GotoState27Action,
	{_State91, MulToken}:                                _GotoState22Action,
	{_State91, BitNegToken}:                             _GotoState13Action,
	{_State91, BitAndToken}:                             _GotoState12Action,
	{_State91, LexErrorToken}:                           _GotoState20Action,
	{_State91, LiteralType}:                             _GotoState39Action,
	{_State91, AnonymousFuncExprType}:                   _GotoState31Action,
	{_State91, AnonymousStructExprType}:                 _GotoState32Action,
	{_State91, AtomExprType}:                            _GotoState33Action,
	{_State91, CallExprType}:                            _GotoState35Action,
	{_State91, AccessExprType}:                          _GotoState28Action,
	{_State91, UnaryOpType}:                             _GotoState45Action,
	{_State91, UnaryExprType}:                           _GotoState44Action,
	{_State91, MulExprType}:                             _GotoState40Action,
	{_State91, AddExprType}:                             _GotoState29Action,
	{_State91, CmpExprType}:                             _GotoState36Action,
	{_State91, AndExprType}:                             _GotoState130Action,
	{_State91, OptionalLabelDeclType}:                   _GotoState92Action,
	{_State91, BlockExprType}:                           _GotoState34Action,
	{_State91, ExplicitStructTypeType}:                  _GotoState38Action,
	{_State91, ExplicitFuncTypeType}:                    _GotoState37Action,
	{_State92, LbraceToken}:                             _GotoState51Action,
	{_State92, BlockBodyType}:                           _GotoState87Action,
	{_State94, LparenToken}:                             _GotoState56Action,
	{_State94, IdentifierToken}:                         _GotoState131Action,
	{_State94, StructToken}:                             _GotoState26Action,
	{_State94, EnumToken}:                               _GotoState97Action,
	{_State94, AtomTypeType}:                            _GotoState99Action,
	{_State94, ImplicitStructTypeType}:                  _GotoState103Action,
	{_State94, ExplicitStructTypeType}:                  _GotoState101Action,
	{_State94, ExplicitEnumTypeType}:                    _GotoState100Action,
	{_State94, ImplicitEnumTypeType}:                    _GotoState102Action,
	{_State94, ValueTypeType}:                           _GotoState132Action,
	{_State95, GenericParametersType}:                   _GotoState133Action,
	{_State96, LparenToken}:                             _GotoState56Action,
	{_State96, IdentifierToken}:                         _GotoState131Action,
	{_State96, StructToken}:                             _GotoState26Action,
	{_State96, EnumToken}:                               _GotoState97Action,
	{_State96, AtomTypeType}:                            _GotoState99Action,
	{_State96, ImplicitStructTypeType}:                  _GotoState103Action,
	{_State96, ExplicitStructTypeType}:                  _GotoState101Action,
	{_State96, ExplicitEnumTypeType}:                    _GotoState100Action,
	{_State96, ImplicitEnumTypeType}:                    _GotoState102Action,
	{_State96, ValueTypeType}:                           _GotoState134Action,
	{_State97, LparenToken}:                             _GotoState135Action,
	{_State98, LparenToken}:                             _GotoState56Action,
	{_State98, GenericLbraceToken}:                      _GotoState95Action,
	{_State98, QuestionToken}:                           _GotoState136Action,
	{_State98, IdentifierToken}:                         _GotoState131Action,
	{_State98, StructToken}:                             _GotoState26Action,
	{_State98, EnumToken}:                               _GotoState97Action,
	{_State98, OptionalGenericParametersType}:           _GotoState137Action,
	{_State98, AtomTypeType}:                            _GotoState99Action,
	{_State98, ImplicitStructTypeType}:                  _GotoState103Action,
	{_State98, ExplicitStructTypeType}:                  _GotoState101Action,
	{_State98, ExplicitEnumTypeType}:                    _GotoState100Action,
	{_State98, ImplicitEnumTypeType}:                    _GotoState102Action,
	{_State98, ValueTypeType}:                           _GotoState138Action,
	{_State99, BitOrToken}:                              _GotoState139Action,
	{_State102, BitOrToken}:                             _GotoState140Action,
	{_State104, RparenToken}:                            _GotoState141Action,
	{_State106, GenericLbraceToken}:                     _GotoState95Action,
	{_State106, OptionalGenericParametersType}:          _GotoState142Action,
	{_State107, BoolLiteralToken}:                       _GotoState14Action,
	{_State107, IntegerLiteralToken}:                    _GotoState18Action,
	{_State107, FloatLiteralToken}:                      _GotoState15Action,
	{_State107, RuneLiteralToken}:                       _GotoState24Action,
	{_State107, StringLiteralToken}:                     _GotoState25Action,
	{_State107, RbraceToken}:                            _GotoState146Action,
	{_State107, LparenToken}:                            _GotoState21Action,
	{_State107, IdentifierToken}:                        _GotoState17Action,
	{_State107, ReturnToken}:                            _GotoState147Action,
	{_State107, BreakToken}:                             _GotoState144Action,
	{_State107, ContinueToken}:                          _GotoState145Action,
	{_State107, StructToken}:                            _GotoState26Action,
	{_State107, FuncToken}:                              _GotoState16Action,
	{_State107, AsyncToken}:                             _GotoState143Action,
	{_State107, LabelDeclToken}:                         _GotoState19Action,
	{_State107, NotToken}:                               _GotoState23Action,
	{_State107, SubToken}:                               _GotoState27Action,
	{_State107, MulToken}:                               _GotoState22Action,
	{_State107, BitNegToken}:                            _GotoState13Action,
	{_State107, BitAndToken}:                            _GotoState12Action,
	{_State107, LexErrorToken}:                          _GotoState20Action,
	{_State107, LiteralType}:                            _GotoState39Action,
	{_State107, AnonymousFuncExprType}:                  _GotoState31Action,
	{_State107, AnonymousStructExprType}:                _GotoState32Action,
	{_State107, AtomExprType}:                           _GotoState33Action,
	{_State107, CallExprType}:                           _GotoState35Action,
	{_State107, AccessExprType}:                         _GotoState148Action,
	{_State107, UnaryOpType}:                            _GotoState45Action,
	{_State107, UnaryExprType}:                          _GotoState44Action,
	{_State107, MulExprType}:                            _GotoState40Action,
	{_State107, AddExprType}:                            _GotoState29Action,
	{_State107, CmpExprType}:                            _GotoState36Action,
	{_State107, AndExprType}:                            _GotoState30Action,
	{_State107, OrExprType}:                             _GotoState42Action,
	{_State107, SequenceExprType}:                       _GotoState43Action,
	{_State107, JumpTypeType}:                           _GotoState151Action,
	{_State107, ExpressionOrImplicitStructType}:         _GotoState150Action,
	{_State107, StatementBodyType}:                      _GotoState153Action,
	{_State107, StatementType}:                          _GotoState152Action,
	{_State107, OptionalLabelDeclType}:                  _GotoState41Action,
	{_State107, BlockExprType}:                          _GotoState34Action,
	{_State107, ExpressionType}:                         _GotoState149Action,
	{_State107, ExplicitStructTypeType}:                 _GotoState38Action,
	{_State107, ExplicitFuncTypeType}:                   _GotoState37Action,
	{_State108, LparenToken}:                            _GotoState56Action,
	{_State108, IdentifierToken}:                        _GotoState131Action,
	{_State108, StructToken}:                            _GotoState26Action,
	{_State108, EnumToken}:                              _GotoState97Action,
	{_State108, AtomTypeType}:                           _GotoState99Action,
	{_State108, ImplicitStructTypeType}:                 _GotoState103Action,
	{_State108, ExplicitStructTypeType}:                 _GotoState101Action,
	{_State108, ExplicitEnumTypeType}:                   _GotoState100Action,
	{_State108, ImplicitEnumTypeType}:                   _GotoState102Action,
	{_State108, ValueTypeType}:                          _GotoState154Action,
	{_State109, LparenToken}:                            _GotoState56Action,
	{_State109, GenericLbraceToken}:                     _GotoState95Action,
	{_State109, QuestionToken}:                          _GotoState136Action,
	{_State109, DotdotdotToken}:                         _GotoState155Action,
	{_State109, IdentifierToken}:                        _GotoState131Action,
	{_State109, StructToken}:                            _GotoState26Action,
	{_State109, EnumToken}:                              _GotoState97Action,
	{_State109, OptionalGenericParametersType}:          _GotoState137Action,
	{_State109, AtomTypeType}:                           _GotoState99Action,
	{_State109, ImplicitStructTypeType}:                 _GotoState103Action,
	{_State109, ExplicitStructTypeType}:                 _GotoState101Action,
	{_State109, ExplicitEnumTypeType}:                   _GotoState100Action,
	{_State109, ImplicitEnumTypeType}:                   _GotoState102Action,
	{_State109, ValueTypeType}:                          _GotoState138Action,
	{_State110, CommaToken}:                             _GotoState156Action,
	{_State113, RparenToken}:                            _GotoState157Action,
	{_State114, CommaToken}:                             _GotoState158Action,
	{_State118, RbraceToken}:                            _GotoState159Action,
	{_State119, RbracketToken}:                          _GotoState160Action,
	{_State120, RparenToken}:                            _GotoState161Action,
	{_State121, MulToken}:                               _GotoState82Action,
	{_State121, DivToken}:                               _GotoState80Action,
	{_State121, ModToken}:                               _GotoState81Action,
	{_State121, BitAndToken}:                            _GotoState77Action,
	{_State121, BitLshiftToken}:                         _GotoState78Action,
	{_State121, BitRshiftToken}:                         _GotoState79Action,
	{_State121, MulOpType}:                              _GotoState83Action,
	{_State122, EqualToken}:                             _GotoState68Action,
	{_State122, NotEqualToken}:                          _GotoState73Action,
	{_State122, LessToken}:                              _GotoState71Action,
	{_State122, LessOrEqualToken}:                       _GotoState72Action,
	{_State122, GreaterToken}:                           _GotoState69Action,
	{_State122, GreaterOrEqualToken}:                    _GotoState70Action,
	{_State122, CmpOpType}:                              _GotoState74Action,
	{_State123, AddToken}:                               _GotoState62Action,
	{_State123, SubToken}:                               _GotoState65Action,
	{_State123, BitXorToken}:                            _GotoState64Action,
	{_State123, BitOrToken}:                             _GotoState63Action,
	{_State123, AddOpType}:                              _GotoState66Action,
	{_State124, RparenToken}:                            _GotoState162Action,
	{_State127, LabelDeclToken}:                         _GotoState19Action,
	{_State127, OptionalLabelDeclType}:                  _GotoState92Action,
	{_State127, BlockExprType}:                          _GotoState163Action,
	{_State128, LbraceToken}:                            _GotoState51Action,
	{_State128, BlockBodyType}:                          _GotoState164Action,
	{_State129, DefaultToken}:                           _GotoState165Action,
	{_State130, AndToken}:                               _GotoState67Action,
	{_State131, GenericLbraceToken}:                     _GotoState95Action,
	{_State131, OptionalGenericParametersType}:          _GotoState137Action,
	{_State133, RbraceToken}:                            _GotoState166Action,
	{_State135, RparenToken}:                            _GotoState167Action,
	{_State139, LparenToken}:                            _GotoState56Action,
	{_State139, IdentifierToken}:                        _GotoState131Action,
	{_State139, StructToken}:                            _GotoState26Action,
	{_State139, EnumToken}:                              _GotoState97Action,
	{_State139, AtomTypeType}:                           _GotoState168Action,
	{_State139, ImplicitStructTypeType}:                 _GotoState103Action,
	{_State139, ExplicitStructTypeType}:                 _GotoState101Action,
	{_State139, ExplicitEnumTypeType}:                   _GotoState100Action,
	{_State140, LparenToken}:                            _GotoState56Action,
	{_State140, IdentifierToken}:                        _GotoState131Action,
	{_State140, StructToken}:                            _GotoState26Action,
	{_State140, EnumToken}:                              _GotoState97Action,
	{_State140, AtomTypeType}:                           _GotoState169Action,
	{_State140, ImplicitStructTypeType}:                 _GotoState103Action,
	{_State140, ExplicitStructTypeType}:                 _GotoState101Action,
	{_State140, ExplicitEnumTypeType}:                   _GotoState100Action,
	{_State142, LparenToken}:                            _GotoState53Action,
	{_State142, ImplicitFuncTypeType}:                   _GotoState170Action,
	{_State143, BoolLiteralToken}:                       _GotoState14Action,
	{_State143, IntegerLiteralToken}:                    _GotoState18Action,
	{_State143, FloatLiteralToken}:                      _GotoState15Action,
	{_State143, RuneLiteralToken}:                       _GotoState24Action,
	{_State143, StringLiteralToken}:                     _GotoState25Action,
	{_State143, LparenToken}:                            _GotoState21Action,
	{_State143, IdentifierToken}:                        _GotoState17Action,
	{_State143, StructToken}:                            _GotoState26Action,
	{_State143, FuncToken}:                              _GotoState16Action,
	{_State143, LabelDeclToken}:                         _GotoState19Action,
	{_State143, LexErrorToken}:                          _GotoState20Action,
	{_State143, LiteralType}:                            _GotoState39Action,
	{_State143, AnonymousFuncExprType}:                  _GotoState31Action,
	{_State143, AnonymousStructExprType}:                _GotoState32Action,
	{_State143, AtomExprType}:                           _GotoState33Action,
	{_State143, CallExprType}:                           _GotoState172Action,
	{_State143, AccessExprType}:                         _GotoState171Action,
	{_State143, OptionalLabelDeclType}:                  _GotoState92Action,
	{_State143, BlockExprType}:                          _GotoState34Action,
	{_State143, ExplicitStructTypeType}:                 _GotoState38Action,
	{_State143, ExplicitFuncTypeType}:                   _GotoState37Action,
	{_State148, LparenToken}:                            _GotoState61Action,
	{_State148, LbracketToken}:                          _GotoState60Action,
	{_State148, GenericLbraceToken}:                     _GotoState59Action,
	{_State148, DotToken}:                               _GotoState58Action,
	{_State148, AddAssignToken}:                         _GotoState173Action,
	{_State148, SubAssignToken}:                         _GotoState184Action,
	{_State148, MulAssignToken}:                         _GotoState183Action,
	{_State148, DivAssignToken}:                         _GotoState181Action,
	{_State148, ModAssignToken}:                         _GotoState182Action,
	{_State148, AddOneAssignToken}:                      _GotoState174Action,
	{_State148, SubOneAssignToken}:                      _GotoState185Action,
	{_State148, BitNegAssignToken}:                      _GotoState177Action,
	{_State148, BitAndAssignToken}:                      _GotoState175Action,
	{_State148, BitOrAssignToken}:                       _GotoState178Action,
	{_State148, BitXorAssignToken}:                      _GotoState180Action,
	{_State148, BitLshiftAssignToken}:                   _GotoState176Action,
	{_State148, BitRshiftAssignToken}:                   _GotoState179Action,
	{_State148, OpOneAssignType}:                        _GotoState187Action,
	{_State148, BinaryOpAssignType}:                     _GotoState186Action,
	{_State150, CommaToken}:                             _GotoState188Action,
	{_State151, JumpLabelToken}:                         _GotoState189Action,
	{_State151, OptionalJumpLabelType}:                  _GotoState190Action,
	{_State153, NewlinesToken}:                          _GotoState191Action,
	{_State153, SemicolonToken}:                         _GotoState192Action,
	{_State155, LparenToken}:                            _GotoState56Action,
	{_State155, QuestionToken}:                          _GotoState193Action,
	{_State155, IdentifierToken}:                        _GotoState131Action,
	{_State155, StructToken}:                            _GotoState26Action,
	{_State155, EnumToken}:                              _GotoState97Action,
	{_State155, AtomTypeType}:                           _GotoState99Action,
	{_State155, ImplicitStructTypeType}:                 _GotoState103Action,
	{_State155, ExplicitStructTypeType}:                 _GotoState101Action,
	{_State155, ExplicitEnumTypeType}:                   _GotoState100Action,
	{_State155, ImplicitEnumTypeType}:                   _GotoState102Action,
	{_State155, ValueTypeType}:                          _GotoState194Action,
	{_State156, LparenToken}:                            _GotoState56Action,
	{_State156, DotdotdotToken}:                         _GotoState108Action,
	{_State156, IdentifierToken}:                        _GotoState109Action,
	{_State156, StructToken}:                            _GotoState26Action,
	{_State156, EnumToken}:                              _GotoState97Action,
	{_State156, AtomTypeType}:                           _GotoState99Action,
	{_State156, ImplicitStructTypeType}:                 _GotoState103Action,
	{_State156, ExplicitStructTypeType}:                 _GotoState101Action,
	{_State156, ExplicitEnumTypeType}:                   _GotoState100Action,
	{_State156, ImplicitEnumTypeType}:                   _GotoState102Action,
	{_State156, ValueTypeType}:                          _GotoState105Action,
	{_State156, ParameterDeclType}:                      _GotoState196Action,
	{_State156, VarargType}:                             _GotoState114Action,
	{_State156, OptionalVarargType}:                     _GotoState195Action,
	{_State157, LparenToken}:                            _GotoState56Action,
	{_State157, QuestionToken}:                          _GotoState197Action,
	{_State157, IdentifierToken}:                        _GotoState131Action,
	{_State157, StructToken}:                            _GotoState26Action,
	{_State157, EnumToken}:                              _GotoState97Action,
	{_State157, AtomTypeType}:                           _GotoState99Action,
	{_State157, ImplicitStructTypeType}:                 _GotoState103Action,
	{_State157, ExplicitStructTypeType}:                 _GotoState101Action,
	{_State157, ExplicitEnumTypeType}:                   _GotoState100Action,
	{_State157, ImplicitEnumTypeType}:                   _GotoState102Action,
	{_State157, ValueTypeType}:                          _GotoState198Action,
	{_State159, LparenToken}:                            _GotoState199Action,
	{_State164, ElseToken}:                              _GotoState200Action,
	{_State171, LparenToken}:                            _GotoState61Action,
	{_State171, LbracketToken}:                          _GotoState60Action,
	{_State171, GenericLbraceToken}:                     _GotoState59Action,
	{_State171, DotToken}:                               _GotoState58Action,
	{_State186, BoolLiteralToken}:                       _GotoState14Action,
	{_State186, IntegerLiteralToken}:                    _GotoState18Action,
	{_State186, FloatLiteralToken}:                      _GotoState15Action,
	{_State186, RuneLiteralToken}:                       _GotoState24Action,
	{_State186, StringLiteralToken}:                     _GotoState25Action,
	{_State186, LparenToken}:                            _GotoState21Action,
	{_State186, IdentifierToken}:                        _GotoState17Action,
	{_State186, StructToken}:                            _GotoState26Action,
	{_State186, FuncToken}:                              _GotoState16Action,
	{_State186, LabelDeclToken}:                         _GotoState19Action,
	{_State186, NotToken}:                               _GotoState23Action,
	{_State186, SubToken}:                               _GotoState27Action,
	{_State186, MulToken}:                               _GotoState22Action,
	{_State186, BitNegToken}:                            _GotoState13Action,
	{_State186, BitAndToken}:                            _GotoState12Action,
	{_State186, LexErrorToken}:                          _GotoState20Action,
	{_State186, LiteralType}:                            _GotoState39Action,
	{_State186, AnonymousFuncExprType}:                  _GotoState31Action,
	{_State186, AnonymousStructExprType}:                _GotoState32Action,
	{_State186, AtomExprType}:                           _GotoState33Action,
	{_State186, CallExprType}:                           _GotoState35Action,
	{_State186, AccessExprType}:                         _GotoState28Action,
	{_State186, UnaryOpType}:                            _GotoState45Action,
	{_State186, UnaryExprType}:                          _GotoState44Action,
	{_State186, MulExprType}:                            _GotoState40Action,
	{_State186, AddExprType}:                            _GotoState29Action,
	{_State186, CmpExprType}:                            _GotoState36Action,
	{_State186, AndExprType}:                            _GotoState30Action,
	{_State186, OrExprType}:                             _GotoState42Action,
	{_State186, SequenceExprType}:                       _GotoState43Action,
	{_State186, OptionalLabelDeclType}:                  _GotoState41Action,
	{_State186, BlockExprType}:                          _GotoState34Action,
	{_State186, ExpressionType}:                         _GotoState201Action,
	{_State186, ExplicitStructTypeType}:                 _GotoState38Action,
	{_State186, ExplicitFuncTypeType}:                   _GotoState37Action,
	{_State188, BoolLiteralToken}:                       _GotoState14Action,
	{_State188, IntegerLiteralToken}:                    _GotoState18Action,
	{_State188, FloatLiteralToken}:                      _GotoState15Action,
	{_State188, RuneLiteralToken}:                       _GotoState24Action,
	{_State188, StringLiteralToken}:                     _GotoState25Action,
	{_State188, LparenToken}:                            _GotoState21Action,
	{_State188, IdentifierToken}:                        _GotoState17Action,
	{_State188, StructToken}:                            _GotoState26Action,
	{_State188, FuncToken}:                              _GotoState16Action,
	{_State188, LabelDeclToken}:                         _GotoState19Action,
	{_State188, NotToken}:                               _GotoState23Action,
	{_State188, SubToken}:                               _GotoState27Action,
	{_State188, MulToken}:                               _GotoState22Action,
	{_State188, BitNegToken}:                            _GotoState13Action,
	{_State188, BitAndToken}:                            _GotoState12Action,
	{_State188, LexErrorToken}:                          _GotoState20Action,
	{_State188, LiteralType}:                            _GotoState39Action,
	{_State188, AnonymousFuncExprType}:                  _GotoState31Action,
	{_State188, AnonymousStructExprType}:                _GotoState32Action,
	{_State188, AtomExprType}:                           _GotoState33Action,
	{_State188, CallExprType}:                           _GotoState35Action,
	{_State188, AccessExprType}:                         _GotoState28Action,
	{_State188, UnaryOpType}:                            _GotoState45Action,
	{_State188, UnaryExprType}:                          _GotoState44Action,
	{_State188, MulExprType}:                            _GotoState40Action,
	{_State188, AddExprType}:                            _GotoState29Action,
	{_State188, CmpExprType}:                            _GotoState36Action,
	{_State188, AndExprType}:                            _GotoState30Action,
	{_State188, OrExprType}:                             _GotoState42Action,
	{_State188, SequenceExprType}:                       _GotoState43Action,
	{_State188, OptionalLabelDeclType}:                  _GotoState41Action,
	{_State188, BlockExprType}:                          _GotoState34Action,
	{_State188, ExpressionType}:                         _GotoState202Action,
	{_State188, ExplicitStructTypeType}:                 _GotoState38Action,
	{_State188, ExplicitFuncTypeType}:                   _GotoState37Action,
	{_State190, BoolLiteralToken}:                       _GotoState14Action,
	{_State190, IntegerLiteralToken}:                    _GotoState18Action,
	{_State190, FloatLiteralToken}:                      _GotoState15Action,
	{_State190, RuneLiteralToken}:                       _GotoState24Action,
	{_State190, StringLiteralToken}:                     _GotoState25Action,
	{_State190, LparenToken}:                            _GotoState21Action,
	{_State190, IdentifierToken}:                        _GotoState17Action,
	{_State190, StructToken}:                            _GotoState26Action,
	{_State190, FuncToken}:                              _GotoState16Action,
	{_State190, LabelDeclToken}:                         _GotoState19Action,
	{_State190, NotToken}:                               _GotoState23Action,
	{_State190, SubToken}:                               _GotoState27Action,
	{_State190, MulToken}:                               _GotoState22Action,
	{_State190, BitNegToken}:                            _GotoState13Action,
	{_State190, BitAndToken}:                            _GotoState12Action,
	{_State190, LexErrorToken}:                          _GotoState20Action,
	{_State190, LiteralType}:                            _GotoState39Action,
	{_State190, AnonymousFuncExprType}:                  _GotoState31Action,
	{_State190, AnonymousStructExprType}:                _GotoState32Action,
	{_State190, AtomExprType}:                           _GotoState33Action,
	{_State190, CallExprType}:                           _GotoState35Action,
	{_State190, AccessExprType}:                         _GotoState28Action,
	{_State190, UnaryOpType}:                            _GotoState45Action,
	{_State190, UnaryExprType}:                          _GotoState44Action,
	{_State190, MulExprType}:                            _GotoState40Action,
	{_State190, AddExprType}:                            _GotoState29Action,
	{_State190, CmpExprType}:                            _GotoState36Action,
	{_State190, AndExprType}:                            _GotoState30Action,
	{_State190, OrExprType}:                             _GotoState42Action,
	{_State190, SequenceExprType}:                       _GotoState43Action,
	{_State190, OptionalExpressionOrImplicitStructType}: _GotoState204Action,
	{_State190, ExpressionOrImplicitStructType}:         _GotoState203Action,
	{_State190, OptionalLabelDeclType}:                  _GotoState41Action,
	{_State190, BlockExprType}:                          _GotoState34Action,
	{_State190, ExpressionType}:                         _GotoState149Action,
	{_State190, ExplicitStructTypeType}:                 _GotoState38Action,
	{_State190, ExplicitFuncTypeType}:                   _GotoState37Action,
	{_State199, ArgumentsType}:                          _GotoState205Action,
	{_State200, LbraceToken}:                            _GotoState51Action,
	{_State200, IfToken}:                                _GotoState85Action,
	{_State200, BlockBodyType}:                          _GotoState206Action,
	{_State200, IfExprType}:                             _GotoState207Action,
	{_State203, CommaToken}:                             _GotoState188Action,
	{_State205, RparenToken}:                            _GotoState208Action,
	{_State3, _WildcardMarker}:                          _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State10, IdentifierToken}:                         _ReduceNilToOptionalReceiverAction,
	{_State12, _WildcardMarker}:                         _ReduceBitAndToUnaryOpAction,
	{_State13, _WildcardMarker}:                         _ReduceBitNegToUnaryOpAction,
	{_State14, _WildcardMarker}:                         _ReduceBoolLiteralToLiteralAction,
	{_State15, _WildcardMarker}:                         _ReduceFloatLiteralToLiteralAction,
	{_State17, _WildcardMarker}:                         _ReduceIdentifierToAtomExprAction,
	{_State18, _WildcardMarker}:                         _ReduceIntegerLiteralToLiteralAction,
	{_State19, _WildcardMarker}:                         _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State20, _WildcardMarker}:                         _ReduceLexErrorToAtomExprAction,
	{_State21, RparenToken}:                             _ReduceToArgumentsAction,
	{_State22, _WildcardMarker}:                         _ReduceMulToUnaryOpAction,
	{_State23, _WildcardMarker}:                         _ReduceNotToUnaryOpAction,
	{_State24, _WildcardMarker}:                         _ReduceRuneLiteralToLiteralAction,
	{_State25, _WildcardMarker}:                         _ReduceStringLiteralToLiteralAction,
	{_State27, _WildcardMarker}:                         _ReduceSubToUnaryOpAction,
	{_State28, _WildcardMarker}:                         _ReduceAccessExprToUnaryExprAction,
	{_State29, _WildcardMarker}:                         _ReduceAddExprToCmpExprAction,
	{_State30, _WildcardMarker}:                         _ReduceAndExprToOrExprAction,
	{_State31, _WildcardMarker}:                         _ReduceAnonymousFuncExprToAtomExprAction,
	{_State32, _WildcardMarker}:                         _ReduceAnonymousStructExprToAtomExprAction,
	{_State33, _WildcardMarker}:                         _ReduceAtomExprToAccessExprAction,
	{_State34, _WildcardMarker}:                         _ReduceBlockExprToAtomExprAction,
	{_State35, _WildcardMarker}:                         _ReduceCallExprToAccessExprAction,
	{_State36, _WildcardMarker}:                         _ReduceCmpExprToAndExprAction,
	{_State39, _WildcardMarker}:                         _ReduceLiteralToAtomExprAction,
	{_State40, _WildcardMarker}:                         _ReduceMulExprToAddExprAction,
	{_State42, _EndMarker}:                              _ReduceToSequenceExprAction,
	{_State43, _EndMarker}:                              _ReduceSequenceExprToExpressionAction,
	{_State44, _WildcardMarker}:                         _ReduceUnaryExprToMulExprAction,
	{_State45, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State46, _EndMarker}:                              _ReduceCommentToLexInternalTokensAction,
	{_State47, _EndMarker}:                              _ReduceSpacesToLexInternalTokensAction,
	{_State48, _WildcardMarker}:                         _ReduceNilToOptionalGenericParametersAction,
	{_State51, _WildcardMarker}:                         _ReduceEmptyListToStatementsAction,
	{_State52, _EndMarker}:                              _ReduceToFuncDefAction,
	{_State53, RparenToken}:                             _ReduceNilToOptionalVarargAction,
	{_State54, LbraceToken}:                             _ReduceToExplicitFuncTypeAction,
	{_State57, _WildcardMarker}:                         _ReduceToExplicitStructTypeAction,
	{_State59, RbraceToken}:                             _ReduceToGenericArgumentsAction,
	{_State60, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State61, RparenToken}:                             _ReduceToArgumentsAction,
	{_State62, _WildcardMarker}:                         _ReduceAddToAddOpAction,
	{_State63, _WildcardMarker}:                         _ReduceBitOrToAddOpAction,
	{_State64, _WildcardMarker}:                         _ReduceBitXorToAddOpAction,
	{_State65, _WildcardMarker}:                         _ReduceSubToAddOpAction,
	{_State66, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State67, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State68, _WildcardMarker}:                         _ReduceEqualToCmpOpAction,
	{_State69, _WildcardMarker}:                         _ReduceGreaterToCmpOpAction,
	{_State70, _WildcardMarker}:                         _ReduceGreaterOrEqualToCmpOpAction,
	{_State71, _WildcardMarker}:                         _ReduceLessToCmpOpAction,
	{_State72, _WildcardMarker}:                         _ReduceLessOrEqualToCmpOpAction,
	{_State73, _WildcardMarker}:                         _ReduceNotEqualToCmpOpAction,
	{_State74, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State75, _WildcardMarker}:                         _ReduceToAnonymousFuncExprAction,
	{_State76, RparenToken}:                             _ReduceToArgumentsAction,
	{_State77, _WildcardMarker}:                         _ReduceBitAndToMulOpAction,
	{_State78, _WildcardMarker}:                         _ReduceBitLshiftToMulOpAction,
	{_State79, _WildcardMarker}:                         _ReduceBitRshiftToMulOpAction,
	{_State80, _WildcardMarker}:                         _ReduceDivToMulOpAction,
	{_State81, _WildcardMarker}:                         _ReduceModToMulOpAction,
	{_State82, _WildcardMarker}:                         _ReduceMulToMulOpAction,
	{_State83, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State84, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State85, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State87, _EndMarker}:                              _ReduceToBlockExprAction,
	{_State88, _EndMarker}:                              _ReduceIfExprToExpressionAction,
	{_State89, _EndMarker}:                              _ReduceLoopExprToExpressionAction,
	{_State90, _EndMarker}:                              _ReduceMatchExprToExpressionAction,
	{_State91, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State93, _WildcardMarker}:                         _ReduceOpToUnaryExprAction,
	{_State95, RbraceToken}:                             _ReduceToGenericParametersAction,
	{_State98, _WildcardMarker}:                         _ReduceNilToOptionalGenericParametersAction,
	{_State99, _EndMarker}:                              _ReduceAtomTypeToValueTypeAction,
	{_State100, _WildcardMarker}:                        _ReduceExplicitEnumTypeToAtomTypeAction,
	{_State101, _WildcardMarker}:                        _ReduceExplicitStructTypeToAtomTypeAction,
	{_State102, _EndMarker}:                             _ReduceImplicitEnumTypeToValueTypeAction,
	{_State103, _WildcardMarker}:                        _ReduceImplicitStructTypeToAtomTypeAction,
	{_State105, _WildcardMarker}:                        _ReduceImplicitToParameterDeclAction,
	{_State106, LparenToken}:                            _ReduceNilToOptionalGenericParametersAction,
	{_State107, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State109, _WildcardMarker}:                        _ReduceNilToOptionalGenericParametersAction,
	{_State110, RparenToken}:                            _ReduceNonEmptyParametersToParametersAction,
	{_State111, RparenToken}:                            _ReduceOptionalVarargToParametersAction,
	{_State112, _WildcardMarker}:                        _ReduceParameterDeclToNonEmptyParametersAction,
	{_State114, RparenToken}:                            _ReduceVarargToOptionalVarargAction,
	{_State115, _WildcardMarker}:                        _ReduceImplicitToAnonymousStructExprAction,
	{_State116, _WildcardMarker}:                        _ReduceToImplicitStructTypeAction,
	{_State117, _WildcardMarker}:                        _ReduceAccessToAccessExprAction,
	{_State121, _WildcardMarker}:                        _ReduceOpToAddExprAction,
	{_State122, _WildcardMarker}:                        _ReduceOpToAndExprAction,
	{_State123, _WildcardMarker}:                        _ReduceOpToCmpExprAction,
	{_State125, _WildcardMarker}:                        _ReduceOpToMulExprAction,
	{_State126, _WildcardMarker}:                        _ReduceBlockExprToAtomExprAction,
	{_State126, _EndMarker}:                             _ReduceInfiniteToLoopExprAction,
	{_State127, LbraceToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State130, _WildcardMarker}:                        _ReduceOpToOrExprAction,
	{_State131, _WildcardMarker}:                        _ReduceNilToOptionalGenericParametersAction,
	{_State132, _EndMarker}:                             _ReduceAliasToTypeDeclAction,
	{_State134, _EndMarker}:                             _ReduceDefinitionToTypeDeclAction,
	{_State136, _WildcardMarker}:                        _ReduceInferredToParameterDeclAction,
	{_State137, _WildcardMarker}:                        _ReduceNamedToAtomTypeAction,
	{_State138, _WildcardMarker}:                        _ReduceExplicitToParameterDeclAction,
	{_State141, IdentifierToken}:                        _ReduceReceiverToOptionalReceiverAction,
	{_State143, LbraceToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State144, _WildcardMarker}:                        _ReduceBreakToJumpTypeAction,
	{_State145, _WildcardMarker}:                        _ReduceContinueToJumpTypeAction,
	{_State146, _EndMarker}:                             _ReduceToBlockBodyAction,
	{_State147, _WildcardMarker}:                        _ReduceReturnToJumpTypeAction,
	{_State148, _WildcardMarker}:                        _ReduceAccessExprToUnaryExprAction,
	{_State149, _WildcardMarker}:                        _ReduceExpressionToExpressionOrImplicitStructAction,
	{_State150, _WildcardMarker}:                        _ReduceExpressionOrImplicitStructToStatementBodyAction,
	{_State151, _WildcardMarker}:                        _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State152, _WildcardMarker}:                        _ReduceAddToStatementsAction,
	{_State154, _WildcardMarker}:                        _ReduceImplicitToVarargAction,
	{_State156, RparenToken}:                            _ReduceNilToOptionalVarargAction,
	{_State157, LbraceToken}:                            _ReduceImplicitUnitToImplicitFuncTypeAction,
	{_State158, RparenToken}:                            _ReduceVararg2ToOptionalVarargAction,
	{_State160, _WildcardMarker}:                        _ReduceIndexToAccessExprAction,
	{_State161, _WildcardMarker}:                        _ReduceConcreteToCallExprAction,
	{_State162, _WildcardMarker}:                        _ReduceExplicitToAnonymousStructExprAction,
	{_State163, _EndMarker}:                             _ReduceWhileToLoopExprAction,
	{_State164, _WildcardMarker}:                        _ReduceNoElseToIfExprAction,
	{_State165, _EndMarker}:                             _ReduceToMatchExprAction,
	{_State166, _WildcardMarker}:                        _ReduceGenericToOptionalGenericParametersAction,
	{_State167, _WildcardMarker}:                        _ReduceToExplicitEnumTypeAction,
	{_State168, _WildcardMarker}:                        _ReducePairToImplicitEnumTypeAction,
	{_State169, _WildcardMarker}:                        _ReduceAddToImplicitEnumTypeAction,
	{_State170, LbraceToken}:                            _ReduceToFuncDeclAction,
	{_State172, _WildcardMarker}:                        _ReduceCallExprToAccessExprAction,
	{_State172, NewlinesToken}:                          _ReduceAsyncToStatementBodyAction,
	{_State172, SemicolonToken}:                         _ReduceAsyncToStatementBodyAction,
	{_State173, _WildcardMarker}:                        _ReduceAddAssignToBinaryOpAssignAction,
	{_State174, _WildcardMarker}:                        _ReduceAddOneAssignToOpOneAssignAction,
	{_State175, _WildcardMarker}:                        _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State176, _WildcardMarker}:                        _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State177, _WildcardMarker}:                        _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State178, _WildcardMarker}:                        _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State179, _WildcardMarker}:                        _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State180, _WildcardMarker}:                        _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State181, _WildcardMarker}:                        _ReduceDivAssignToBinaryOpAssignAction,
	{_State182, _WildcardMarker}:                        _ReduceModAssignToBinaryOpAssignAction,
	{_State183, _WildcardMarker}:                        _ReduceMulAssignToBinaryOpAssignAction,
	{_State184, _WildcardMarker}:                        _ReduceSubAssignToBinaryOpAssignAction,
	{_State185, _WildcardMarker}:                        _ReduceSubOneAssignToOpOneAssignAction,
	{_State186, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State187, _WildcardMarker}:                        _ReduceOpOneAssignToStatementBodyAction,
	{_State188, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State189, _WildcardMarker}:                        _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State190, NewlinesToken}:                          _ReduceNilToOptionalExpressionOrImplicitStructAction,
	{_State190, SemicolonToken}:                         _ReduceNilToOptionalExpressionOrImplicitStructAction,
	{_State190, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State191, _WildcardMarker}:                        _ReduceImplicitToStatementAction,
	{_State192, _WildcardMarker}:                        _ReduceExplicitToStatementAction,
	{_State193, _WildcardMarker}:                        _ReduceInferredToVarargAction,
	{_State194, _WildcardMarker}:                        _ReduceExplicitToVarargAction,
	{_State195, RparenToken}:                            _ReduceMixedToParametersAction,
	{_State196, _WildcardMarker}:                        _ReduceAddToNonEmptyParametersAction,
	{_State197, LbraceToken}:                            _ReduceInferredToImplicitFuncTypeAction,
	{_State198, LbraceToken}:                            _ReduceTypedToImplicitFuncTypeAction,
	{_State199, RparenToken}:                            _ReduceToArgumentsAction,
	{_State201, _WildcardMarker}:                        _ReduceBinaryOpAssignToStatementBodyAction,
	{_State202, _WildcardMarker}:                        _ReduceImplicitStructToExpressionOrImplicitStructAction,
	{_State203, _WildcardMarker}:                        _ReduceExpressionOrImplicitStructToOptionalExpressionOrImplicitStructAction,
	{_State204, _WildcardMarker}:                        _ReduceJumpToStatementBodyAction,
	{_State206, _EndMarker}:                             _ReduceIfElseToIfExprAction,
	{_State207, _EndMarker}:                             _ReduceMultiIfElseToIfExprAction,
	{_State208, _WildcardMarker}:                        _ReduceGenericToCallExprAction,
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.type_decl
    Reduce:
      (nil)
    Goto:
      TYPE -> State 9
      type_decl -> State 5

  State 2:
    Kernel Items:
      #accept: ^.func_def
    Reduce:
      (nil)
    Goto:
      FUNC -> State 10
      func_decl -> State 11
      func_def -> State 6

  State 3:
    Kernel Items:
      #accept: ^.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 14
      INTEGER_LITERAL -> State 18
      FLOAT_LITERAL -> State 15
      RUNE_LITERAL -> State 24
      STRING_LITERAL -> State 25
      LPAREN -> State 21
      IDENTIFIER -> State 17
      STRUCT -> State 26
      FUNC -> State 16
      LABEL_DECL -> State 19
      NOT -> State 23
      SUB -> State 27
      MUL -> State 22
      BIT_NEG -> State 13
      BIT_AND -> State 12
      LEX_ERROR -> State 20
      literal -> State 39
      anonymous_func_expr -> State 31
      anonymous_struct_expr -> State 32
      atom_expr -> State 33
      call_expr -> State 35
      access_expr -> State 28
      unary_op -> State 45
      unary_expr -> State 44
      mul_expr -> State 40
      add_expr -> State 29
      cmp_expr -> State 36
      and_expr -> State 30
      or_expr -> State 42
      sequence_expr -> State 43
      optional_label_decl -> State 41
      block_expr -> State 34
      expression -> State 7
      explicit_struct_type -> State 38
      explicit_func_type -> State 37

  State 4:
    Kernel Items:
      #accept: ^.lex_internal_tokens
    Reduce:
      (nil)
    Goto:
      SPACES -> State 47
      COMMENT -> State 46
      lex_internal_tokens -> State 8

  State 5:
    Kernel Items:
      #accept: ^ type_decl., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 6:
    Kernel Items:
      #accept: ^ func_def., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 7:
    Kernel Items:
      #accept: ^ expression., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      #accept: ^ lex_internal_tokens., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      type_decl: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_decl: TYPE.IDENTIFIER EQUAL value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 48

  State 10:
    Kernel Items:
      func_decl: FUNC.optional_receiver IDENTIFIER optional_generic_parameters implicit_func_type
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      LPAREN -> State 49
      optional_receiver -> State 50

  State 11:
    Kernel Items:
      func_def: func_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 51
      block_body -> State 52

  State 12:
    Kernel Items:
      unary_op: BIT_AND., *
    Reduce:
      * -> [unary_op]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      unary_op: BIT_NEG., *
    Reduce:
      * -> [unary_op]
    Goto:
      (nil)

  State 14:
    Kernel Items:
      literal: BOOL_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 15:
    Kernel Items:
      literal: FLOAT_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 16:
    Kernel Items:
      explicit_func_type: FUNC.implicit_func_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 53
      implicit_func_type -> State 54

  State 17:
    Kernel Items:
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 18:
    Kernel Items:
      literal: INTEGER_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 19:
    Kernel Items:
      optional_label_decl: LABEL_DECL., *
    Reduce:
      * -> [optional_label_decl]
    Goto:
      (nil)

  State 20:
    Kernel Items:
      atom_expr: LEX_ERROR., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      anonymous_struct_expr: LPAREN.arguments RPAREN
    Reduce:
      RPAREN -> [arguments]
    Goto:
      arguments -> State 55

  State 22:
    Kernel Items:
      unary_op: MUL., *
    Reduce:
      * -> [unary_op]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      unary_op: NOT., *
    Reduce:
      * -> [unary_op]
    Goto:
      (nil)

  State 24:
    Kernel Items:
      literal: RUNE_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 25:
    Kernel Items:
      literal: STRING_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      explicit_struct_type: STRUCT.implicit_struct_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 56
      implicit_struct_type -> State 57

  State 27:
    Kernel Items:
      unary_op: SUB., *
    Reduce:
      * -> [unary_op]
    Goto:
      (nil)

  State 28:
    Kernel Items:
      call_expr: access_expr.LPAREN arguments RPAREN
      call_expr: access_expr.GENERIC_LBRACE generic_arguments RBRACE LPAREN arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET expression RBRACKET
      unary_expr: access_expr., *
    Reduce:
      * -> [unary_expr]
    Goto:
      LPAREN -> State 61
      LBRACKET -> State 60
      GENERIC_LBRACE -> State 59
      DOT -> State 58

  State 29:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 62
      SUB -> State 65
      BIT_XOR -> State 64
      BIT_OR -> State 63
      add_op -> State 66

  State 30:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 67

  State 31:
    Kernel Items:
      atom_expr: anonymous_func_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      atom_expr: anonymous_struct_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 33:
    Kernel Items:
      access_expr: atom_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 34:
    Kernel Items:
      atom_expr: block_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 35:
    Kernel Items:
      access_expr: call_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 36:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 68
      NOT_EQUAL -> State 73
      LESS -> State 71
      LESS_OR_EQUAL -> State 72
      GREATER -> State 69
      GREATER_OR_EQUAL -> State 70
      cmp_op -> State 74

  State 37:
    Kernel Items:
      anonymous_func_expr: explicit_func_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 51
      block_body -> State 75

  State 38:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 76

  State 39:
    Kernel Items:
      atom_expr: literal., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 40:
    Kernel Items:
      mul_expr: mul_expr.mul_op unary_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 82
      DIV -> State 80
      MOD -> State 81
      BIT_AND -> State 77
      BIT_LSHIFT -> State 78
      BIT_RSHIFT -> State 79
      mul_op -> State 83

  State 41:
    Kernel Items:
      block_expr: optional_label_decl.block_body
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.match_expr
      expression: optional_label_decl.loop_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 51
      IF -> State 85
      MATCH -> State 86
      FOR -> State 84
      block_body -> State 87
      if_expr -> State 88
      match_expr -> State 90
      loop_expr -> State 89

  State 42:
    Kernel Items:
      or_expr: or_expr.OR and_expr
      sequence_expr: or_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      OR -> State 91

  State 43:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 44:
    Kernel Items:
      mul_expr: unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 45:
    Kernel Items:
      unary_expr: unary_op.unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 14
      INTEGER_LITERAL -> State 18
      FLOAT_LITERAL -> State 15
      RUNE_LITERAL -> State 24
      STRING_LITERAL -> State 25
      LPAREN -> State 21
      IDENTIFIER -> State 17
      STRUCT -> State 26
      FUNC -> State 16
      LABEL_DECL -> State 19
      NOT -> State 23
      SUB -> State 27
      MUL -> State 22
      BIT_NEG -> State 13
      BIT_AND -> State 12
      LEX_ERROR -> State 20
      literal -> State 39
      anonymous_func_expr -> State 31
      anonymous_struct_expr -> State 32
      atom_expr -> State 33
      call_expr -> State 35
      access_expr -> State 28
      unary_op -> State 45
      unary_expr -> State 93
      optional_label_decl -> State 92
      block_expr -> State 34
      explicit_struct_type -> State 38
      explicit_func_type -> State 37

  State 46:
    Kernel Items:
      lex_internal_tokens: COMMENT., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 47:
    Kernel Items:
      lex_internal_tokens: SPACES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      type_decl: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_decl: TYPE IDENTIFIER.EQUAL value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      GENERIC_LBRACE -> State 95
      EQUAL -> State 94
      optional_generic_parameters -> State 96

  State 49:
    Kernel Items:
      optional_receiver: LPAREN.parameter_decl RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 56
      IDENTIFIER -> State 98
      STRUCT -> State 26
      ENUM -> State 97
      atom_type -> State 99
      implicit_struct_type -> State 103
      explicit_struct_type -> State 101
      explicit_enum_type -> State 100
      implicit_enum_type -> State 102
      value_type -> State 105
      parameter_decl -> State 104

  State 50:
    Kernel Items:
      func_decl: FUNC optional_receiver.IDENTIFIER optional_generic_parameters implicit_func_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 106

  State 51:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 107

  State 52:
    Kernel Items:
      func_def: func_decl block_body., $
    Reduce:
      $ -> [func_def]
    Goto:
      (nil)

  State 53:
    Kernel Items:
      implicit_func_type: LPAREN.parameters RPAREN value_type
      implicit_func_type: LPAREN.parameters RPAREN QUESTION
      implicit_func_type: LPAREN.parameters RPAREN
    Reduce:
      RPAREN -> [optional_vararg]
    Goto:
      LPAREN -> State 56
      DOTDOTDOT -> State 108
      IDENTIFIER -> State 109
      STRUCT -> State 26
      ENUM -> State 97
      atom_type -> State 99
      implicit_struct_type -> State 103
      explicit_struct_type -> State 101
      explicit_enum_type -> State 100
      implicit_enum_type -> State 102
      value_type -> State 105
      parameter_decl -> State 112
      non_empty_parameters -> State 110
      vararg -> State 114
      optional_vararg -> State 111
      parameters -> State 113

  State 54:
    Kernel Items:
      explicit_func_type: FUNC implicit_func_type., LBRACE
    Reduce:
      LBRACE -> [explicit_func_type]
    Goto:
      (nil)

  State 55:
    Kernel Items:
      anonymous_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 115

  State 56:
    Kernel Items:
      implicit_struct_type: LPAREN.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 116

  State 57:
    Kernel Items:
      explicit_struct_type: STRUCT implicit_struct_type., *
    Reduce:
      * -> [explicit_struct_type]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 117

  State 59:
    Kernel Items:
      call_expr: access_expr GENERIC_LBRACE.generic_arguments RBRACE LPAREN arguments RPAREN
    Reduce:
      RBRACE -> [generic_arguments]
    Goto:
      generic_arguments -> State 118

  State 60:
    Kernel Items:
      access_expr: access_expr LBRACKET.expression RBRACKET
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 14
      INTEGER_LITERAL -> State 18
      FLOAT_LITERAL -> State 15
      RUNE_LITERAL -> State 24
      STRING_LITERAL -> State 25
      LPAREN -> State 21
      IDENTIFIER -> State 17
      STRUCT -> State 26
      FUNC -> State 16
      LABEL_DECL -> State 19
      NOT -> State 23
      SUB -> State 27
      MUL -> State 22
      BIT_NEG -> State 13
      BIT_AND -> State 12
      LEX_ERROR -> State 20
      literal -> State 39
      anonymous_func_expr -> State 31
      anonymous_struct_expr -> State 32
      atom_expr -> State 33
      call_expr -> State 35
      access_expr -> State 28
      unary_op -> State 45
      unary_expr -> State 44
      mul_expr -> State 40
      add_expr -> State 29
      cmp_expr -> State 36
      and_expr -> State 30
      or_expr -> State 42
      sequence_expr -> State 43
      optional_label_decl -> State 41
      block_expr -> State 34
      expression -> State 119
      explicit_struct_type -> State 38
      explicit_func_type -> State 37

  State 61:
    Kernel Items:
      call_expr: access_expr LPAREN.arguments RPAREN
    Reduce:
      RPAREN -> [arguments]
    Goto:
      arguments -> State 120

  State 62:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 63:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 64:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 66:
    Kernel Items:
      add_expr: add_expr add_op.mul_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 14
      INTEGER_LITERAL -> State 18
      FLOAT_LITERAL -> State 15
      RUNE_LITERAL -> State 24
      STRING_LITERAL -> State 25
      LPAREN -> State 21
      IDENTIFIER -> State 17
      STRUCT -> State 26
      FUNC -> State 16
      LABEL_DECL -> State 19
      NOT -> State 23
      SUB -> State 27
      MUL -> State 22
      BIT_NEG -> State 13
      BIT_AND -> State 12
      LEX_ERROR -> State 20
      literal -> State 39
      anonymous_func_expr -> State 31
      anonymous_struct_expr -> State 32
      atom_expr -> State 33
      call_expr -> State 35
      access_expr -> State 28
      unary_op -> State 45
      unary_expr -> State 44
      mul_expr -> State 121
      optional_label_decl -> State 92
      block_expr -> State 34
      explicit_struct_type -> State 38
      explicit_func_type -> State 37

  State 67:
    Kernel Items:
      and_expr: and_expr AND.cmp_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 14
      INTEGER_LITERAL -> State 18
      FLOAT_LITERAL -> State 15
      RUNE_LITERAL -> State 24
      STRING_LITERAL -> State 25
      LPAREN -> State 21
      IDENTIFIER -> State 17
      STRUCT -> State 26
      FUNC -> State 16
      LABEL_DECL -> State 19
      NOT -> State 23
      SUB -> State 27
      MUL -> State 22
      BIT_NEG -> State 13
      BIT_AND -> State 12
      LEX_ERROR -> State 20
      literal -> State 39
      anonymous_func_expr -> State 31
      anonymous_struct_expr -> State 32
      atom_expr -> State 33
      call_expr -> State 35
      access_expr -> State 28
      unary_op -> State 45
      unary_expr -> State 44
      mul_expr -> State 40
      add_expr -> State 29
      cmp_expr -> State 122
      optional_label_decl -> State 92
      block_expr -> State 34
      explicit_struct_type -> State 38
      explicit_func_type -> State 37

  State 68:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 69:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 70:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 71:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 72:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 73:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 74:
    Kernel Items:
      cmp_expr: cmp_expr cmp_op.add_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 14
      INTEGER_LITERAL -> State 18
      FLOAT_LITERAL -> State 15
      RUNE_LITERAL -> State 24
      STRING_LITERAL -> State 25
      LPAREN -> State 21
      IDENTIFIER -> State 17
      STRUCT -> State 26
      FUNC -> State 16
      LABEL_DECL -> State 19
      NOT -> State 23
      SUB -> State 27
      MUL -> State 22
      BIT_NEG -> State 13
      BIT_AND -> State 12
      LEX_ERROR -> State 20
      literal -> State 39
      anonymous_func_expr -> State 31
      anonymous_struct_expr -> State 32
      atom_expr -> State 33
      call_expr -> State 35
      access_expr -> State 28
      unary_op -> State 45
      unary_expr -> State 44
      mul_expr -> State 40
      add_expr -> State 123
      optional_label_decl -> State 92
      block_expr -> State 34
      explicit_struct_type -> State 38
      explicit_func_type -> State 37

  State 75:
    Kernel Items:
      anonymous_func_expr: explicit_func_type block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 76:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type LPAREN.arguments RPAREN
    Reduce:
      RPAREN -> [arguments]
    Goto:
      arguments -> State 124

  State 77:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 78:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 79:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 80:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 81:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 82:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      mul_expr: mul_expr mul_op.unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 14
      INTEGER_LITERAL -> State 18
      FLOAT_LITERAL -> State 15
      RUNE_LITERAL -> State 24
      STRING_LITERAL -> State 25
      LPAREN -> State 21
      IDENTIFIER -> State 17
      STRUCT -> State 26
      FUNC -> State 16
      LABEL_DECL -> State 19
      NOT -> State 23
      SUB -> State 27
      MUL -> State 22
      BIT_NEG -> State 13
      BIT_AND -> State 12
      LEX_ERROR -> State 20
      literal -> State 39
      anonymous_func_expr -> State 31
      anonymous_struct_expr -> State 32
      atom_expr -> State 33
      call_expr -> State 35
      access_expr -> State 28
      unary_op -> State 45
      unary_expr -> State 125
      optional_label_decl -> State 92
      block_expr -> State 34
      explicit_struct_type -> State 38
      explicit_func_type -> State 37

  State 84:
    Kernel Items:
      loop_expr: FOR.block_expr
      loop_expr: FOR.sequence_expr block_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 14
      INTEGER_LITERAL -> State 18
      FLOAT_LITERAL -> State 15
      RUNE_LITERAL -> State 24
      STRING_LITERAL -> State 25
      LPAREN -> State 21
      IDENTIFIER -> State 17
      STRUCT -> State 26
      FUNC -> State 16
      LABEL_DECL -> State 19
      NOT -> State 23
      SUB -> State 27
      MUL -> State 22
      BIT_NEG -> State 13
      BIT_AND -> State 12
      LEX_ERROR -> State 20
      literal -> State 39
      anonymous_func_expr -> State 31
      anonymous_struct_expr -> State 32
      atom_expr -> State 33
      call_expr -> State 35
      access_expr -> State 28
      unary_op -> State 45
      unary_expr -> State 44
      mul_expr -> State 40
      add_expr -> State 29
      cmp_expr -> State 36
      and_expr -> State 30
      or_expr -> State 42
      sequence_expr -> State 127
      optional_label_decl -> State 92
      block_expr -> State 126
      explicit_struct_type -> State 38
      explicit_func_type -> State 37

  State 85:
    Kernel Items:
      if_expr: IF.sequence_expr block_body
      if_expr: IF.sequence_expr block_body ELSE block_body
      if_expr: IF.sequence_expr block_body ELSE if_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 14
      INTEGER_LITERAL -> State 18
      FLOAT_LITERAL -> State 15
      RUNE_LITERAL -> State 24
      STRING_LITERAL -> State 25
      LPAREN -> State 21
      IDENTIFIER -> State 17
      STRUCT -> State 26
      FUNC -> State 16
      LABEL_DECL -> State 19
      NOT -> State 23
      SUB -> State 27
      MUL -> State 22
      BIT_NEG -> State 13
      BIT_AND -> State 12
      LEX_ERROR -> State 20
      literal -> State 39
      anonymous_func_expr -> State 31
      anonymous_struct_expr -> State 32
      atom_expr -> State 33
      call_expr -> State 35
      access_expr -> State 28
      unary_op -> State 45
      unary_expr -> State 44
      mul_expr -> State 40
      add_expr -> State 29
      cmp_expr -> State 36
      and_expr -> State 30
      or_expr -> State 42
      sequence_expr -> State 128
      optional_label_decl -> State 92
      block_expr -> State 34
      explicit_struct_type -> State 38
      explicit_func_type -> State 37

  State 86:
    Kernel Items:
      match_expr: MATCH.CASE DEFAULT
    Reduce:
      (nil)
    Goto:
      CASE -> State 129

  State 87:
    Kernel Items:
      block_expr: optional_label_decl block_body., $
    Reduce:
      $ -> [block_expr]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 90:
    Kernel Items:
      expression: optional_label_decl match_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 91:
    Kernel Items:
      or_expr: or_expr OR.and_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 14
      INTEGER_LITERAL -> State 18
      FLOAT_LITERAL -> State 15
      RUNE_LITERAL -> State 24
      STRING_LITERAL -> State 25
      LPAREN -> State 21
      IDENTIFIER -> State 17
      STRUCT -> State 26
      FUNC -> State 16
      LABEL_DECL -> State 19
      NOT -> State 23
      SUB -> State 27
      MUL -> State 22
      BIT_NEG -> State 13
      BIT_AND -> State 12
      LEX_ERROR -> State 20
      literal -> State 39
      anonymous_func_expr -> State 31
      anonymous_struct_expr -> State 32
      atom_expr -> State 33
      call_expr -> State 35
      access_expr -> State 28
      unary_op -> State 45
      unary_expr -> State 44
      mul_expr -> State 40
      add_expr -> State 29
      cmp_expr -> State 36
      and_expr -> State 130
      optional_label_decl -> State 92
      block_expr -> State 34
      explicit_struct_type -> State 38
      explicit_func_type -> State 37

  State 92:
    Kernel Items:
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 51
      block_body -> State 87

  State 93:
    Kernel Items:
      unary_expr: unary_op unary_expr., *
    Reduce:
      * -> [unary_expr]
    Goto:
      (nil)

  State 94:
    Kernel Items:
      type_decl: TYPE IDENTIFIER EQUAL.value_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 56
      IDENTIFIER -> State 131
      STRUCT -> State 26
      ENUM -> State 97
      atom_type -> State 99
      implicit_struct_type -> State 103
      explicit_struct_type -> State 101
      explicit_enum_type -> State 100
      implicit_enum_type -> State 102
      value_type -> State 132

  State 95:
    Kernel Items:
      optional_generic_parameters: GENERIC_LBRACE.generic_parameters RBRACE
    Reduce:
      RBRACE -> [generic_parameters]
    Goto:
      generic_parameters -> State 133

  State 96:
    Kernel Items:
      type_decl: TYPE IDENTIFIER optional_generic_parameters.value_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 56
      IDENTIFIER -> State 131
      STRUCT -> State 26
      ENUM -> State 97
      atom_type -> State 99
      implicit_struct_type -> State 103
      explicit_struct_type -> State 101
      explicit_enum_type -> State 100
      implicit_enum_type -> State 102
      value_type -> State 134

  State 97:
    Kernel Items:
      explicit_enum_type: ENUM.LPAREN RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 135

  State 98:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_parameters
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.QUESTION
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      LPAREN -> State 56
      GENERIC_LBRACE -> State 95
      QUESTION -> State 136
      IDENTIFIER -> State 131
      STRUCT -> State 26
      ENUM -> State 97
      optional_generic_parameters -> State 137
      atom_type -> State 99
      implicit_struct_type -> State 103
      explicit_struct_type -> State 101
      explicit_enum_type -> State 100
      implicit_enum_type -> State 102
      value_type -> State 138

  State 99:
    Kernel Items:
      implicit_enum_type: atom_type.BIT_OR atom_type
      value_type: atom_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      BIT_OR -> State 139

  State 100:
    Kernel Items:
      atom_type: explicit_enum_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 101:
    Kernel Items:
      atom_type: explicit_struct_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 102:
    Kernel Items:
      implicit_enum_type: implicit_enum_type.BIT_OR atom_type
      value_type: implicit_enum_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      BIT_OR -> State 140

  State 103:
    Kernel Items:
      atom_type: implicit_struct_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 104:
    Kernel Items:
      optional_receiver: LPAREN parameter_decl.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 141

  State 105:
    Kernel Items:
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 106:
    Kernel Items:
      func_decl: FUNC optional_receiver IDENTIFIER.optional_generic_parameters implicit_func_type
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      GENERIC_LBRACE -> State 95
      optional_generic_parameters -> State 142

  State 107:
    Kernel Items:
      statements: statements.statement
      block_body: LBRACE statements.RBRACE
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 14
      INTEGER_LITERAL -> State 18
      FLOAT_LITERAL -> State 15
      RUNE_LITERAL -> State 24
      STRING_LITERAL -> State 25
      RBRACE -> State 146
      LPAREN -> State 21
      IDENTIFIER -> State 17
      RETURN -> State 147
      BREAK -> State 144
      CONTINUE -> State 145
      STRUCT -> State 26
      FUNC -> State 16
      ASYNC -> State 143
      LABEL_DECL -> State 19
      NOT -> State 23
      SUB -> State 27
      MUL -> State 22
      BIT_NEG -> State 13
      BIT_AND -> State 12
      LEX_ERROR -> State 20
      literal -> State 39
      anonymous_func_expr -> State 31
      anonymous_struct_expr -> State 32
      atom_expr -> State 33
      call_expr -> State 35
      access_expr -> State 148
      unary_op -> State 45
      unary_expr -> State 44
      mul_expr -> State 40
      add_expr -> State 29
      cmp_expr -> State 36
      and_expr -> State 30
      or_expr -> State 42
      sequence_expr -> State 43
      jump_type -> State 151
      expression_or_implicit_struct -> State 150
      statement_body -> State 153
      statement -> State 152
      optional_label_decl -> State 41
      block_expr -> State 34
      expression -> State 149
      explicit_struct_type -> State 38
      explicit_func_type -> State 37

  State 108:
    Kernel Items:
      vararg: DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 56
      IDENTIFIER -> State 131
      STRUCT -> State 26
      ENUM -> State 97
      atom_type -> State 99
      implicit_struct_type -> State 103
      explicit_struct_type -> State 101
      explicit_enum_type -> State 100
      implicit_enum_type -> State 102
      value_type -> State 154

  State 109:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_parameters
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.QUESTION
      vararg: IDENTIFIER.DOTDOTDOT value_type
      vararg: IDENTIFIER.DOTDOTDOT QUESTION
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      LPAREN -> State 56
      GENERIC_LBRACE -> State 95
      QUESTION -> State 136
      DOTDOTDOT -> State 155
      IDENTIFIER -> State 131
      STRUCT -> State 26
      ENUM -> State 97
      optional_generic_parameters -> State 137
      atom_type -> State 99
      implicit_struct_type -> State 103
      explicit_struct_type -> State 101
      explicit_enum_type -> State 100
      implicit_enum_type -> State 102
      value_type -> State 138

  State 110:
    Kernel Items:
      non_empty_parameters: non_empty_parameters.COMMA parameter_decl
      parameters: non_empty_parameters., RPAREN
      parameters: non_empty_parameters.COMMA optional_vararg
    Reduce:
      RPAREN -> [parameters]
    Goto:
      COMMA -> State 156

  State 111:
    Kernel Items:
      parameters: optional_vararg., RPAREN
    Reduce:
      RPAREN -> [parameters]
    Goto:
      (nil)

  State 112:
    Kernel Items:
      non_empty_parameters: parameter_decl., *
    Reduce:
      * -> [non_empty_parameters]
    Goto:
      (nil)

  State 113:
    Kernel Items:
      implicit_func_type: LPAREN parameters.RPAREN value_type
      implicit_func_type: LPAREN parameters.RPAREN QUESTION
      implicit_func_type: LPAREN parameters.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 157

  State 114:
    Kernel Items:
      optional_vararg: vararg., RPAREN
      optional_vararg: vararg.COMMA
    Reduce:
      RPAREN -> [optional_vararg]
    Goto:
      COMMA -> State 158

  State 115:
    Kernel Items:
      anonymous_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      implicit_struct_type: LPAREN RPAREN., *
    Reduce:
      * -> [implicit_struct_type]
    Goto:
      (nil)

  State 117:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 118:
    Kernel Items:
      call_expr: access_expr GENERIC_LBRACE generic_arguments.RBRACE LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 159

  State 119:
    Kernel Items:
      access_expr: access_expr LBRACKET expression.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 160

  State 120:
    Kernel Items:
      call_expr: access_expr LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 161

  State 121:
    Kernel Items:
      mul_expr: mul_expr.mul_op unary_expr
      add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 82
      DIV -> State 80
      MOD -> State 81
      BIT_AND -> State 77
      BIT_LSHIFT -> State 78
      BIT_RSHIFT -> State 79
      mul_op -> State 83

  State 122:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 68
      NOT_EQUAL -> State 73
      LESS -> State 71
      LESS_OR_EQUAL -> State 72
      GREATER -> State 69
      GREATER_OR_EQUAL -> State 70
      cmp_op -> State 74

  State 123:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 62
      SUB -> State 65
      BIT_XOR -> State 64
      BIT_OR -> State 63
      add_op -> State 66

  State 124:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 162

  State 125:
    Kernel Items:
      mul_expr: mul_expr mul_op unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 126:
    Kernel Items:
      atom_expr: block_expr., *
      loop_expr: FOR block_expr., $
    Reduce:
      * -> [atom_expr]
      $ -> [loop_expr]
    Goto:
      (nil)

  State 127:
    Kernel Items:
      loop_expr: FOR sequence_expr.block_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      LABEL_DECL -> State 19
      optional_label_decl -> State 92
      block_expr -> State 163

  State 128:
    Kernel Items:
      if_expr: IF sequence_expr.block_body
      if_expr: IF sequence_expr.block_body ELSE block_body
      if_expr: IF sequence_expr.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 51
      block_body -> State 164

  State 129:
    Kernel Items:
      match_expr: MATCH CASE.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 165

  State 130:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 67

  State 131:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_parameters
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      GENERIC_LBRACE -> State 95
      optional_generic_parameters -> State 137

  State 132:
    Kernel Items:
      type_decl: TYPE IDENTIFIER EQUAL value_type., $
    Reduce:
      $ -> [type_decl]
    Goto:
      (nil)

  State 133:
    Kernel Items:
      optional_generic_parameters: GENERIC_LBRACE generic_parameters.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 166

  State 134:
    Kernel Items:
      type_decl: TYPE IDENTIFIER optional_generic_parameters value_type., $
    Reduce:
      $ -> [type_decl]
    Goto:
      (nil)

  State 135:
    Kernel Items:
      explicit_enum_type: ENUM LPAREN.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 167

  State 136:
    Kernel Items:
      parameter_decl: IDENTIFIER QUESTION., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 137:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_parameters., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 138:
    Kernel Items:
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 139:
    Kernel Items:
      implicit_enum_type: atom_type BIT_OR.atom_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 56
      IDENTIFIER -> State 131
      STRUCT -> State 26
      ENUM -> State 97
      atom_type -> State 168
      implicit_struct_type -> State 103
      explicit_struct_type -> State 101
      explicit_enum_type -> State 100

  State 140:
    Kernel Items:
      implicit_enum_type: implicit_enum_type BIT_OR.atom_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 56
      IDENTIFIER -> State 131
      STRUCT -> State 26
      ENUM -> State 97
      atom_type -> State 169
      implicit_struct_type -> State 103
      explicit_struct_type -> State 101
      explicit_enum_type -> State 100

  State 141:
    Kernel Items:
      optional_receiver: LPAREN parameter_decl RPAREN., IDENTIFIER
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      (nil)

  State 142:
    Kernel Items:
      func_decl: FUNC optional_receiver IDENTIFIER optional_generic_parameters.implicit_func_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 53
      implicit_func_type -> State 170

  State 143:
    Kernel Items:
      statement_body: ASYNC.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 14
      INTEGER_LITERAL -> State 18
      FLOAT_LITERAL -> State 15
      RUNE_LITERAL -> State 24
      STRING_LITERAL -> State 25
      LPAREN -> State 21
      IDENTIFIER -> State 17
      STRUCT -> State 26
      FUNC -> State 16
      LABEL_DECL -> State 19
      LEX_ERROR -> State 20
      literal -> State 39
      anonymous_func_expr -> State 31
      anonymous_struct_expr -> State 32
      atom_expr -> State 33
      call_expr -> State 172
      access_expr -> State 171
      optional_label_decl -> State 92
      block_expr -> State 34
      explicit_struct_type -> State 38
      explicit_func_type -> State 37

  State 144:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 145:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 146:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 147:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 148:
    Kernel Items:
      call_expr: access_expr.LPAREN arguments RPAREN
      call_expr: access_expr.GENERIC_LBRACE generic_arguments RBRACE LPAREN arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET expression RBRACKET
      unary_expr: access_expr., *
      statement_body: access_expr.op_one_assign
      statement_body: access_expr.binary_op_assign expression
    Reduce:
      * -> [unary_expr]
    Goto:
      LPAREN -> State 61
      LBRACKET -> State 60
      GENERIC_LBRACE -> State 59
      DOT -> State 58
      ADD_ASSIGN -> State 173
      SUB_ASSIGN -> State 184
      MUL_ASSIGN -> State 183
      DIV_ASSIGN -> State 181
      MOD_ASSIGN -> State 182
      ADD_ONE_ASSIGN -> State 174
      SUB_ONE_ASSIGN -> State 185
      BIT_NEG_ASSIGN -> State 177
      BIT_AND_ASSIGN -> State 175
      BIT_OR_ASSIGN -> State 178
      BIT_XOR_ASSIGN -> State 180
      BIT_LSHIFT_ASSIGN -> State 176
      BIT_RSHIFT_ASSIGN -> State 179
      op_one_assign -> State 187
      binary_op_assign -> State 186

  State 149:
    Kernel Items:
      expression_or_implicit_struct: expression., *
    Reduce:
      * -> [expression_or_implicit_struct]
    Goto:
      (nil)

  State 150:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct.COMMA expression
      statement_body: expression_or_implicit_struct., *
    Reduce:
      * -> [statement_body]
    Goto:
      COMMA -> State 188

  State 151:
    Kernel Items:
      statement_body: jump_type.optional_jump_label optional_expression_or_implicit_struct
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 189
      optional_jump_label -> State 190

  State 152:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 153:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 191
      SEMICOLON -> State 192

  State 154:
    Kernel Items:
      vararg: DOTDOTDOT value_type., *
    Reduce:
      * -> [vararg]
    Goto:
      (nil)

  State 155:
    Kernel Items:
      vararg: IDENTIFIER DOTDOTDOT.value_type
      vararg: IDENTIFIER DOTDOTDOT.QUESTION
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 56
      QUESTION -> State 193
      IDENTIFIER -> State 131
      STRUCT -> State 26
      ENUM -> State 97
      atom_type -> State 99
      implicit_struct_type -> State 103
      explicit_struct_type -> State 101
      explicit_enum_type -> State 100
      implicit_enum_type -> State 102
      value_type -> State 194

  State 156:
    Kernel Items:
      non_empty_parameters: non_empty_parameters COMMA.parameter_decl
      parameters: non_empty_parameters COMMA.optional_vararg
    Reduce:
      RPAREN -> [optional_vararg]
    Goto:
      LPAREN -> State 56
      DOTDOTDOT -> State 108
      IDENTIFIER -> State 109
      STRUCT -> State 26
      ENUM -> State 97
      atom_type -> State 99
      implicit_struct_type -> State 103
      explicit_struct_type -> State 101
      explicit_enum_type -> State 100
      implicit_enum_type -> State 102
      value_type -> State 105
      parameter_decl -> State 196
      vararg -> State 114
      optional_vararg -> State 195

  State 157:
    Kernel Items:
      implicit_func_type: LPAREN parameters RPAREN.value_type
      implicit_func_type: LPAREN parameters RPAREN.QUESTION
      implicit_func_type: LPAREN parameters RPAREN., LBRACE
    Reduce:
      LBRACE -> [implicit_func_type]
    Goto:
      LPAREN -> State 56
      QUESTION -> State 197
      IDENTIFIER -> State 131
      STRUCT -> State 26
      ENUM -> State 97
      atom_type -> State 99
      implicit_struct_type -> State 103
      explicit_struct_type -> State 101
      explicit_enum_type -> State 100
      implicit_enum_type -> State 102
      value_type -> State 198

  State 158:
    Kernel Items:
      optional_vararg: vararg COMMA., RPAREN
    Reduce:
      RPAREN -> [optional_vararg]
    Goto:
      (nil)

  State 159:
    Kernel Items:
      call_expr: access_expr GENERIC_LBRACE generic_arguments RBRACE.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 199

  State 160:
    Kernel Items:
      access_expr: access_expr LBRACKET expression RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 161:
    Kernel Items:
      call_expr: access_expr LPAREN arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 162:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_type LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 163:
    Kernel Items:
      loop_expr: FOR sequence_expr block_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 164:
    Kernel Items:
      if_expr: IF sequence_expr block_body., *
      if_expr: IF sequence_expr block_body.ELSE block_body
      if_expr: IF sequence_expr block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 200

  State 165:
    Kernel Items:
      match_expr: MATCH CASE DEFAULT., $
    Reduce:
      $ -> [match_expr]
    Goto:
      (nil)

  State 166:
    Kernel Items:
      optional_generic_parameters: GENERIC_LBRACE generic_parameters RBRACE., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 167:
    Kernel Items:
      explicit_enum_type: ENUM LPAREN RPAREN., *
    Reduce:
      * -> [explicit_enum_type]
    Goto:
      (nil)

  State 168:
    Kernel Items:
      implicit_enum_type: atom_type BIT_OR atom_type., *
    Reduce:
      * -> [implicit_enum_type]
    Goto:
      (nil)

  State 169:
    Kernel Items:
      implicit_enum_type: implicit_enum_type BIT_OR atom_type., *
    Reduce:
      * -> [implicit_enum_type]
    Goto:
      (nil)

  State 170:
    Kernel Items:
      func_decl: FUNC optional_receiver IDENTIFIER optional_generic_parameters implicit_func_type., LBRACE
    Reduce:
      LBRACE -> [func_decl]
    Goto:
      (nil)

  State 171:
    Kernel Items:
      call_expr: access_expr.LPAREN arguments RPAREN
      call_expr: access_expr.GENERIC_LBRACE generic_arguments RBRACE LPAREN arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET expression RBRACKET
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 61
      LBRACKET -> State 60
      GENERIC_LBRACE -> State 59
      DOT -> State 58

  State 172:
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

  State 173:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 174:
    Kernel Items:
      op_one_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 175:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 176:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 177:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 178:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 179:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 180:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 181:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 182:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 183:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 184:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 185:
    Kernel Items:
      op_one_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 186:
    Kernel Items:
      statement_body: access_expr binary_op_assign.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 14
      INTEGER_LITERAL -> State 18
      FLOAT_LITERAL -> State 15
      RUNE_LITERAL -> State 24
      STRING_LITERAL -> State 25
      LPAREN -> State 21
      IDENTIFIER -> State 17
      STRUCT -> State 26
      FUNC -> State 16
      LABEL_DECL -> State 19
      NOT -> State 23
      SUB -> State 27
      MUL -> State 22
      BIT_NEG -> State 13
      BIT_AND -> State 12
      LEX_ERROR -> State 20
      literal -> State 39
      anonymous_func_expr -> State 31
      anonymous_struct_expr -> State 32
      atom_expr -> State 33
      call_expr -> State 35
      access_expr -> State 28
      unary_op -> State 45
      unary_expr -> State 44
      mul_expr -> State 40
      add_expr -> State 29
      cmp_expr -> State 36
      and_expr -> State 30
      or_expr -> State 42
      sequence_expr -> State 43
      optional_label_decl -> State 41
      block_expr -> State 34
      expression -> State 201
      explicit_struct_type -> State 38
      explicit_func_type -> State 37

  State 187:
    Kernel Items:
      statement_body: access_expr op_one_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 188:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct COMMA.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      BOOL_LITERAL -> State 14
      INTEGER_LITERAL -> State 18
      FLOAT_LITERAL -> State 15
      RUNE_LITERAL -> State 24
      STRING_LITERAL -> State 25
      LPAREN -> State 21
      IDENTIFIER -> State 17
      STRUCT -> State 26
      FUNC -> State 16
      LABEL_DECL -> State 19
      NOT -> State 23
      SUB -> State 27
      MUL -> State 22
      BIT_NEG -> State 13
      BIT_AND -> State 12
      LEX_ERROR -> State 20
      literal -> State 39
      anonymous_func_expr -> State 31
      anonymous_struct_expr -> State 32
      atom_expr -> State 33
      call_expr -> State 35
      access_expr -> State 28
      unary_op -> State 45
      unary_expr -> State 44
      mul_expr -> State 40
      add_expr -> State 29
      cmp_expr -> State 36
      and_expr -> State 30
      or_expr -> State 42
      sequence_expr -> State 43
      optional_label_decl -> State 41
      block_expr -> State 34
      expression -> State 202
      explicit_struct_type -> State 38
      explicit_func_type -> State 37

  State 189:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 190:
    Kernel Items:
      statement_body: jump_type optional_jump_label.optional_expression_or_implicit_struct
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_expression_or_implicit_struct]
      SEMICOLON -> [optional_expression_or_implicit_struct]
    Goto:
      BOOL_LITERAL -> State 14
      INTEGER_LITERAL -> State 18
      FLOAT_LITERAL -> State 15
      RUNE_LITERAL -> State 24
      STRING_LITERAL -> State 25
      LPAREN -> State 21
      IDENTIFIER -> State 17
      STRUCT -> State 26
      FUNC -> State 16
      LABEL_DECL -> State 19
      NOT -> State 23
      SUB -> State 27
      MUL -> State 22
      BIT_NEG -> State 13
      BIT_AND -> State 12
      LEX_ERROR -> State 20
      literal -> State 39
      anonymous_func_expr -> State 31
      anonymous_struct_expr -> State 32
      atom_expr -> State 33
      call_expr -> State 35
      access_expr -> State 28
      unary_op -> State 45
      unary_expr -> State 44
      mul_expr -> State 40
      add_expr -> State 29
      cmp_expr -> State 36
      and_expr -> State 30
      or_expr -> State 42
      sequence_expr -> State 43
      optional_expression_or_implicit_struct -> State 204
      expression_or_implicit_struct -> State 203
      optional_label_decl -> State 41
      block_expr -> State 34
      expression -> State 149
      explicit_struct_type -> State 38
      explicit_func_type -> State 37

  State 191:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 192:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 193:
    Kernel Items:
      vararg: IDENTIFIER DOTDOTDOT QUESTION., *
    Reduce:
      * -> [vararg]
    Goto:
      (nil)

  State 194:
    Kernel Items:
      vararg: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [vararg]
    Goto:
      (nil)

  State 195:
    Kernel Items:
      parameters: non_empty_parameters COMMA optional_vararg., RPAREN
    Reduce:
      RPAREN -> [parameters]
    Goto:
      (nil)

  State 196:
    Kernel Items:
      non_empty_parameters: non_empty_parameters COMMA parameter_decl., *
    Reduce:
      * -> [non_empty_parameters]
    Goto:
      (nil)

  State 197:
    Kernel Items:
      implicit_func_type: LPAREN parameters RPAREN QUESTION., LBRACE
    Reduce:
      LBRACE -> [implicit_func_type]
    Goto:
      (nil)

  State 198:
    Kernel Items:
      implicit_func_type: LPAREN parameters RPAREN value_type., LBRACE
    Reduce:
      LBRACE -> [implicit_func_type]
    Goto:
      (nil)

  State 199:
    Kernel Items:
      call_expr: access_expr GENERIC_LBRACE generic_arguments RBRACE LPAREN.arguments RPAREN
    Reduce:
      RPAREN -> [arguments]
    Goto:
      arguments -> State 205

  State 200:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE.block_body
      if_expr: IF sequence_expr block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 51
      IF -> State 85
      block_body -> State 206
      if_expr -> State 207

  State 201:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 202:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct COMMA expression., *
    Reduce:
      * -> [expression_or_implicit_struct]
    Goto:
      (nil)

  State 203:
    Kernel Items:
      optional_expression_or_implicit_struct: expression_or_implicit_struct., *
      expression_or_implicit_struct: expression_or_implicit_struct.COMMA expression
    Reduce:
      * -> [optional_expression_or_implicit_struct]
    Goto:
      COMMA -> State 188

  State 204:
    Kernel Items:
      statement_body: jump_type optional_jump_label optional_expression_or_implicit_struct., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 205:
    Kernel Items:
      call_expr: access_expr GENERIC_LBRACE generic_arguments RBRACE LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 208

  State 206:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 207:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 208:
    Kernel Items:
      call_expr: access_expr GENERIC_LBRACE generic_arguments RBRACE LPAREN arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

Number of states: 208
Number of shift actions: 775
Number of reduce actions: 173
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
