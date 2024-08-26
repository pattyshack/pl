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
	ColonToken           = SymbolId(294)
	DollarLbracketToken  = SymbolId(295)
	DotdotdotToken       = SymbolId(296)
	TildeTildeToken      = SymbolId(297)
	LabelDeclToken       = SymbolId(298)
	JumpLabelToken       = SymbolId(299)
	AssignToken          = SymbolId(300)
	AddAssignToken       = SymbolId(301)
	SubAssignToken       = SymbolId(302)
	MulAssignToken       = SymbolId(303)
	DivAssignToken       = SymbolId(304)
	ModAssignToken       = SymbolId(305)
	AddOneAssignToken    = SymbolId(306)
	SubOneAssignToken    = SymbolId(307)
	BitNegAssignToken    = SymbolId(308)
	BitAndAssignToken    = SymbolId(309)
	BitOrAssignToken     = SymbolId(310)
	BitXorAssignToken    = SymbolId(311)
	BitLshiftAssignToken = SymbolId(312)
	BitRshiftAssignToken = SymbolId(313)
	NotToken             = SymbolId(314)
	AndToken             = SymbolId(315)
	OrToken              = SymbolId(316)
	AddToken             = SymbolId(317)
	SubToken             = SymbolId(318)
	MulToken             = SymbolId(319)
	DivToken             = SymbolId(320)
	ModToken             = SymbolId(321)
	BitNegToken          = SymbolId(322)
	BitAndToken          = SymbolId(323)
	BitXorToken          = SymbolId(324)
	BitOrToken           = SymbolId(325)
	BitLshiftToken       = SymbolId(326)
	BitRshiftToken       = SymbolId(327)
	EqualToken           = SymbolId(328)
	NotEqualToken        = SymbolId(329)
	LessToken            = SymbolId(330)
	LessOrEqualToken     = SymbolId(331)
	GreaterToken         = SymbolId(332)
	GreaterOrEqualToken  = SymbolId(333)
	LexErrorToken        = SymbolId(334)
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
	// 114:2: literal -> TRUE: ...
	TrueToLiteral(True_ *GenericSymbol) (*GenericSymbol, error)

	// 115:2: literal -> FALSE: ...
	FalseToLiteral(False_ *GenericSymbol) (*GenericSymbol, error)

	// 116:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 117:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 118:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 119:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 122:2: anonymous_struct_expr -> explicit: ...
	ExplicitToAnonymousStructExpr(ExplicitStructDef_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 123:2: anonymous_struct_expr -> implicit: ...
	ImplicitToAnonymousStructExpr(Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 129:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 130:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 131:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 132:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 133:2: atom_expr -> anonymous_struct_expr: ...
	AnonymousStructExprToAtomExpr(AnonymousStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 134:2: atom_expr -> LEX_ERROR: ...
	LexErrorToAtomExpr(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 138:2: generic_arguments -> value_type: ...
	ValueTypeToGenericArguments(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 139:2: generic_arguments -> add: ...
	AddToGenericArguments(GenericArguments_ *GenericSymbol, Comma_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 142:2: optional_generic_arguments -> generic_arguments: ...
	GenericArgumentsToOptionalGenericArguments(GenericArguments_ *GenericSymbol) (*GenericSymbol, error)

	// 143:2: optional_generic_arguments -> nil: ...
	NilToOptionalGenericArguments() (*GenericSymbol, error)

	// 146:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ *GenericSymbol, OptionalGenericArguments_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 147:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (*GenericSymbol, error)

	// 150:2: optional_expression -> expression: ...
	ExpressionToOptionalExpression(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 151:2: optional_expression -> nil: ...
	NilToOptionalExpression() (*GenericSymbol, error)

	// 154:2: colon_expressions -> pair: ...
	PairToColonExpressions(OptionalExpression_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_2 *GenericSymbol) (*GenericSymbol, error)

	// 155:2: colon_expressions -> add: ...
	AddToColonExpressions(ColonExpressions_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_ *GenericSymbol) (*GenericSymbol, error)

	// 158:2: argument -> positional: ...
	PositionalToArgument(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 159:2: argument -> named: ...
	NamedToArgument(Identifier_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 160:2: argument -> colon_expressions: ...
	ColonExpressionsToArgument(ColonExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 163:2: arguments -> argument: ...
	ArgumentToArguments(Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 164:2: arguments -> add: ...
	AddToArguments(Arguments_ *GenericSymbol, Comma_ *GenericSymbol, Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 167:2: optional_arguments -> arguments: ...
	ArgumentsToOptionalArguments(Arguments_ *GenericSymbol) (*GenericSymbol, error)

	// 168:2: optional_arguments -> nil: ...
	NilToOptionalArguments() (*GenericSymbol, error)

	// 170:13: call_expr -> ...
	ToCallExpr(AccessExpr_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalArguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 173:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 174:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 175:2: access_expr -> call_expr: ...
	CallExprToAccessExpr(CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 176:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Argument_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 179:2: postfix_unary_expr -> access_expr: ...
	AccessExprToPostfixUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 180:2: postfix_unary_expr -> question: ...
	QuestionToPostfixUnaryExpr(AccessExpr_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 183:2: prefix_unary_op -> NOT: ...
	NotToPrefixUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 184:2: prefix_unary_op -> BIT_NEG: ...
	BitNegToPrefixUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 185:2: prefix_unary_op -> SUB: ...
	SubToPrefixUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 188:2: prefix_unary_op -> MUL: ...
	MulToPrefixUnaryOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 191:2: prefix_unary_op -> BIT_AND: ...
	BitAndToPrefixUnaryOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 194:2: prefix_unary_expr -> postfix_unary_expr: ...
	PostfixUnaryExprToPrefixUnaryExpr(PostfixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 195:2: prefix_unary_expr -> prefix_op: ...
	PrefixOpToPrefixUnaryExpr(PrefixUnaryOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 198:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 199:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 200:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 201:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 202:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 203:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 206:2: mul_expr -> prefix_unary_expr: ...
	PrefixUnaryExprToMulExpr(PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 207:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 210:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 211:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 212:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 213:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 216:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 217:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 220:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 221:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 222:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 223:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 224:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 225:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 228:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 229:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 232:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 233:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 236:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 237:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 239:17: sequence_expr -> ...
	ToSequenceExpr(OrExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 242:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 243:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 246:2: optional_expression_or_implicit_struct -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct(ExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 247:2: optional_expression_or_implicit_struct -> nil: ...
	NilToOptionalExpressionOrImplicitStruct() (*GenericSymbol, error)

	// 250:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 251:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 252:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 255:2: op_one_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToOpOneAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 256:2: op_one_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToOpOneAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 259:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 260:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 261:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 262:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 263:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 264:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 265:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 266:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 267:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 268:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 269:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 272:2: expression_or_implicit_struct -> expression: ...
	ExpressionToExpressionOrImplicitStruct(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 273:2: expression_or_implicit_struct -> implicit_struct: ...
	ImplicitStructToExpressionOrImplicitStruct(ExpressionOrImplicitStruct_ *GenericSymbol, Comma_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 277:20: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ *GenericSymbol, Less_ *GenericSymbol, Identifier_ *GenericSymbol, Greater_ *GenericSymbol, StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 280:2: statement_body -> unsafe_statement: ...
	UnsafeStatementToStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 282:2: statement_body -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToStatementBody(ExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 284:2: statement_body -> async: ...
	AsyncToStatementBody(Async_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 287:2: statement_body -> jump: ...
	JumpToStatementBody(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpressionOrImplicitStruct_ *GenericSymbol) (*GenericSymbol, error)

	// 300:2: statement_body -> op_one_assign: ...
	OpOneAssignToStatementBody(AccessExpr_ *GenericSymbol, OpOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 301:2: statement_body -> binary_op_assign: ...
	BinaryOpAssignToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 304:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 305:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 308:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 309:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 312:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 313:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (*GenericSymbol, error)

	// 315:14: block_body -> ...
	ToBlockBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 316:14: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 321:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 322:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, BlockBody_2 *GenericSymbol) (*GenericSymbol, error)

	// 323:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 326:15: switch_expr -> ...
	ToSwitchExpr(Switch_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Lbrace_ *GenericSymbol, Case_ *GenericSymbol, Default_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 329:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(For_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 330:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 335:2: expression -> sequence_expr: ...
	SequenceExprToExpression(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 336:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 337:2: expression -> switch_expr: ...
	SwitchExprToExpression(OptionalLabelDecl_ *GenericSymbol, SwitchExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 338:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ *GenericSymbol, LoopExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 345:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 346:2: atom_type -> explicit_struct_def: ...
	ExplicitStructDefToAtomType(ExplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 347:2: atom_type -> implicit_struct_def: ...
	ImplicitStructDefToAtomType(ImplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 348:2: atom_type -> explicit_enum_def: ...
	ExplicitEnumDefToAtomType(ExplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 349:2: atom_type -> implicit_enum_def: ...
	ImplicitEnumDefToAtomType(ImplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 350:2: atom_type -> trait_def: ...
	TraitDefToAtomType(TraitDef_ *GenericSymbol) (*GenericSymbol, error)

	// 353:2: traitable_type -> atom_type: ...
	AtomTypeToTraitableType(AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 354:2: traitable_type -> public_methods_trait: ...
	PublicMethodsTraitToTraitableType(BitNeg_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 355:2: traitable_type -> public_trait: ...
	PublicTraitToTraitableType(TildeTilde_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 359:2: trait_algebra_type -> traitable_type: ...
	TraitableTypeToTraitAlgebraType(TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 360:2: trait_algebra_type -> intersect: ...
	IntersectToTraitAlgebraType(TraitAlgebraType_ *GenericSymbol, Mul_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 361:2: trait_algebra_type -> union: ...
	UnionToTraitAlgebraType(TraitAlgebraType_ *GenericSymbol, Add_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 362:2: trait_algebra_type -> difference: ...
	DifferenceToTraitAlgebraType(TraitAlgebraType_ *GenericSymbol, Sub_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 365:2: value_type -> inferred: ...
	InferredToValueType(Question_ *GenericSymbol) (*GenericSymbol, error)

	// 366:2: value_type -> trait_algebra_type: ...
	TraitAlgebraTypeToValueType(TraitAlgebraType_ *GenericSymbol) (*GenericSymbol, error)

	// 367:2: value_type -> reference: ...
	ReferenceToValueType(BitAnd_ *GenericSymbol, TraitAlgebraType_ *GenericSymbol) (*GenericSymbol, error)

	// 368:2: value_type -> func_type: ...
	FuncTypeToValueType(FuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 371:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 372:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol, Implements_ *GenericSymbol, ValueType_2 *GenericSymbol) (*GenericSymbol, error)

	// 373:2: type_def -> alias: ...
	AliasToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, Equal_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 380:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 381:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 384:2: generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToGenericParameterDefs(GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 385:2: generic_parameter_defs -> add: ...
	AddToGenericParameterDefs(GenericParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 388:2: optional_generic_parameter_defs -> generic_parameter_defs: ...
	GenericParameterDefsToOptionalGenericParameterDefs(GenericParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 389:2: optional_generic_parameter_defs -> nil: ...
	NilToOptionalGenericParameterDefs() (*GenericSymbol, error)

	// 392:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ *GenericSymbol, OptionalGenericParameterDefs_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 393:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (*GenericSymbol, error)

	// 400:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 401:2: field_def -> implicit: ...
	ImplicitToFieldDef(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 404:2: implicit_field_defs -> field_def: ...
	FieldDefToImplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 405:2: implicit_field_defs -> add: ...
	AddToImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 408:2: optional_implicit_field_defs -> implicit_field_defs: ...
	ImplicitFieldDefsToOptionalImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 409:2: optional_implicit_field_defs -> nil: ...
	NilToOptionalImplicitFieldDefs() (*GenericSymbol, error)

	// 411:23: implicit_struct_def -> ...
	ToImplicitStructDef(Lparen_ *GenericSymbol, OptionalImplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 414:2: explicit_field_defs -> field_def: ...
	FieldDefToExplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 415:2: explicit_field_defs -> implicit: ...
	ImplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Newlines_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 416:2: explicit_field_defs -> explicit: ...
	ExplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 419:2: optional_explicit_field_defs -> explicit_field_defs: ...
	ExplicitFieldDefsToOptionalExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 420:2: optional_explicit_field_defs -> nil: ...
	NilToOptionalExplicitFieldDefs() (*GenericSymbol, error)

	// 422:23: explicit_struct_def -> ...
	ToExplicitStructDef(Struct_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalExplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 432:2: enum_value_def -> field_def: ...
	FieldDefToEnumValueDef(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 433:2: enum_value_def -> default: ...
	DefaultToEnumValueDef(FieldDef_ *GenericSymbol, Assign_ *GenericSymbol, Default_ *GenericSymbol) (*GenericSymbol, error)

	// 436:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 437:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 439:21: implicit_enum_def -> ...
	ToImplicitEnumDef(Lparen_ *GenericSymbol, ImplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 442:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 443:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 444:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 445:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 447:21: explicit_enum_def -> ...
	ToExplicitEnumDef(Enum_ *GenericSymbol, Lparen_ *GenericSymbol, ExplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 454:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 455:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ *GenericSymbol) (*GenericSymbol, error)

	// 458:2: trait_properties -> trait_property: ...
	TraitPropertyToTraitProperties(TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 459:2: trait_properties -> implicit: ...
	ImplicitToTraitProperties(TraitProperties_ *GenericSymbol, Newlines_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 460:2: trait_properties -> explicit: ...
	ExplicitToTraitProperties(TraitProperties_ *GenericSymbol, Comma_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 463:2: optional_trait_properties -> trait_properties: ...
	TraitPropertiesToOptionalTraitProperties(TraitProperties_ *GenericSymbol) (*GenericSymbol, error)

	// 464:2: optional_trait_properties -> nil: ...
	NilToOptionalTraitProperties() (*GenericSymbol, error)

	// 466:13: trait_def -> ...
	ToTraitDef(Trait_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalTraitProperties_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 474:2: return_type -> value_type: ...
	ValueTypeToReturnType(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 475:2: return_type -> nil: ...
	NilToReturnType() (*GenericSymbol, error)

	// 478:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 479:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 480:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 481:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 484:2: parameter_decls -> parameter_decl: ...
	ParameterDeclToParameterDecls(ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 485:2: parameter_decls -> add: ...
	AddToParameterDecls(ParameterDecls_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 488:2: optional_parameter_decls -> parameter_decls: ...
	ParameterDeclsToOptionalParameterDecls(ParameterDecls_ *GenericSymbol) (*GenericSymbol, error)

	// 489:2: optional_parameter_decls -> nil: ...
	NilToOptionalParameterDecls() (*GenericSymbol, error)

	// 491:13: func_type -> ...
	ToFuncType(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 499:20: method_signature -> ...
	ToMethodSignature(Func_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 502:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 503:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 506:2: parameter_defs -> parameter_def: ...
	ParameterDefToParameterDefs(ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 507:2: parameter_defs -> add: ...
	AddToParameterDefs(ParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 510:2: optional_parameter_defs -> parameter_defs: ...
	ParameterDefsToOptionalParameterDefs(ParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 511:2: optional_parameter_defs -> nil: ...
	NilToOptionalParameterDefs() (*GenericSymbol, error)

	// 514:2: optional_receiver -> receiver: ...
	ReceiverToOptionalReceiver(Lparen_ *GenericSymbol, ParameterDef_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 515:2: optional_receiver -> nil: ...
	NilToOptionalReceiver() (*GenericSymbol, error)

	// 518:2: named_func_def -> ...
	ToNamedFuncDef(Func_ *GenericSymbol, OptionalReceiver_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 521:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 528:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 529:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, PackageStatements_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 531:26: package_statement_body -> ...
	ToPackageStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 534:2: package_statement -> implicit: ...
	ImplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 535:2: package_statement -> explicit: ...
	ExplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 538:2: package_statements -> empty_list: ...
	EmptyListToPackageStatements() (*GenericSymbol, error)

	// 539:2: package_statements -> add: ...
	AddToPackageStatements(PackageStatements_ *GenericSymbol, PackageStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 543:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 544:2: lex_internal_tokens -> COMMENT: ...
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
	case ColonToken:
		return "COLON"
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
	case AssignToken:
		return "ASSIGN"
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
	case OptionalExpressionType:
		return "optional_expression"
	case ColonExpressionsType:
		return "colon_expressions"
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
	case TraitAlgebraTypeType:
		return "trait_algebra_type"
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
	case EnumValueDefType:
		return "enum_value_def"
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

	LiteralType                            = SymbolId(335)
	AnonymousStructExprType                = SymbolId(336)
	AtomExprType                           = SymbolId(337)
	GenericArgumentsType                   = SymbolId(338)
	OptionalGenericArgumentsType           = SymbolId(339)
	OptionalGenericBindingType             = SymbolId(340)
	OptionalExpressionType                 = SymbolId(341)
	ColonExpressionsType                   = SymbolId(342)
	ArgumentType                           = SymbolId(343)
	ArgumentsType                          = SymbolId(344)
	OptionalArgumentsType                  = SymbolId(345)
	CallExprType                           = SymbolId(346)
	AccessExprType                         = SymbolId(347)
	PostfixUnaryExprType                   = SymbolId(348)
	PrefixUnaryOpType                      = SymbolId(349)
	PrefixUnaryExprType                    = SymbolId(350)
	MulOpType                              = SymbolId(351)
	MulExprType                            = SymbolId(352)
	AddOpType                              = SymbolId(353)
	AddExprType                            = SymbolId(354)
	CmpOpType                              = SymbolId(355)
	CmpExprType                            = SymbolId(356)
	AndExprType                            = SymbolId(357)
	OrExprType                             = SymbolId(358)
	SequenceExprType                       = SymbolId(359)
	OptionalJumpLabelType                  = SymbolId(360)
	OptionalExpressionOrImplicitStructType = SymbolId(361)
	JumpTypeType                           = SymbolId(362)
	OpOneAssignType                        = SymbolId(363)
	BinaryOpAssignType                     = SymbolId(364)
	ExpressionOrImplicitStructType         = SymbolId(365)
	UnsafeStatementType                    = SymbolId(366)
	StatementBodyType                      = SymbolId(367)
	StatementType                          = SymbolId(368)
	StatementsType                         = SymbolId(369)
	OptionalLabelDeclType                  = SymbolId(370)
	BlockBodyType                          = SymbolId(371)
	BlockExprType                          = SymbolId(372)
	IfExprType                             = SymbolId(373)
	SwitchExprType                         = SymbolId(374)
	LoopExprType                           = SymbolId(375)
	ExpressionType                         = SymbolId(376)
	AtomTypeType                           = SymbolId(377)
	TraitableTypeType                      = SymbolId(378)
	TraitAlgebraTypeType                   = SymbolId(379)
	ValueTypeType                          = SymbolId(380)
	TypeDefType                            = SymbolId(381)
	GenericParameterDefType                = SymbolId(382)
	GenericParameterDefsType               = SymbolId(383)
	OptionalGenericParameterDefsType       = SymbolId(384)
	OptionalGenericParametersType          = SymbolId(385)
	FieldDefType                           = SymbolId(386)
	ImplicitFieldDefsType                  = SymbolId(387)
	OptionalImplicitFieldDefsType          = SymbolId(388)
	ImplicitStructDefType                  = SymbolId(389)
	ExplicitFieldDefsType                  = SymbolId(390)
	OptionalExplicitFieldDefsType          = SymbolId(391)
	ExplicitStructDefType                  = SymbolId(392)
	EnumValueDefType                       = SymbolId(393)
	ImplicitEnumValueDefsType              = SymbolId(394)
	ImplicitEnumDefType                    = SymbolId(395)
	ExplicitEnumValueDefsType              = SymbolId(396)
	ExplicitEnumDefType                    = SymbolId(397)
	TraitPropertyType                      = SymbolId(398)
	TraitPropertiesType                    = SymbolId(399)
	OptionalTraitPropertiesType            = SymbolId(400)
	TraitDefType                           = SymbolId(401)
	ReturnTypeType                         = SymbolId(402)
	ParameterDeclType                      = SymbolId(403)
	ParameterDeclsType                     = SymbolId(404)
	OptionalParameterDeclsType             = SymbolId(405)
	FuncTypeType                           = SymbolId(406)
	MethodSignatureType                    = SymbolId(407)
	ParameterDefType                       = SymbolId(408)
	ParameterDefsType                      = SymbolId(409)
	OptionalParameterDefsType              = SymbolId(410)
	OptionalReceiverType                   = SymbolId(411)
	NamedFuncDefType                       = SymbolId(412)
	AnonymousFuncExprType                  = SymbolId(413)
	PackageDefType                         = SymbolId(414)
	PackageStatementBodyType               = SymbolId(415)
	PackageStatementType                   = SymbolId(416)
	PackageStatementsType                  = SymbolId(417)
	LexInternalTokensType                  = SymbolId(418)
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
	_ReduceExpressionToOptionalExpression                                 = _ReduceType(21)
	_ReduceNilToOptionalExpression                                        = _ReduceType(22)
	_ReducePairToColonExpressions                                         = _ReduceType(23)
	_ReduceAddToColonExpressions                                          = _ReduceType(24)
	_ReducePositionalToArgument                                           = _ReduceType(25)
	_ReduceNamedToArgument                                                = _ReduceType(26)
	_ReduceColonExpressionsToArgument                                     = _ReduceType(27)
	_ReduceArgumentToArguments                                            = _ReduceType(28)
	_ReduceAddToArguments                                                 = _ReduceType(29)
	_ReduceArgumentsToOptionalArguments                                   = _ReduceType(30)
	_ReduceNilToOptionalArguments                                         = _ReduceType(31)
	_ReduceToCallExpr                                                     = _ReduceType(32)
	_ReduceAtomExprToAccessExpr                                           = _ReduceType(33)
	_ReduceAccessToAccessExpr                                             = _ReduceType(34)
	_ReduceCallExprToAccessExpr                                           = _ReduceType(35)
	_ReduceIndexToAccessExpr                                              = _ReduceType(36)
	_ReduceAccessExprToPostfixUnaryExpr                                   = _ReduceType(37)
	_ReduceQuestionToPostfixUnaryExpr                                     = _ReduceType(38)
	_ReduceNotToPrefixUnaryOp                                             = _ReduceType(39)
	_ReduceBitNegToPrefixUnaryOp                                          = _ReduceType(40)
	_ReduceSubToPrefixUnaryOp                                             = _ReduceType(41)
	_ReduceMulToPrefixUnaryOp                                             = _ReduceType(42)
	_ReduceBitAndToPrefixUnaryOp                                          = _ReduceType(43)
	_ReducePostfixUnaryExprToPrefixUnaryExpr                              = _ReduceType(44)
	_ReducePrefixOpToPrefixUnaryExpr                                      = _ReduceType(45)
	_ReduceMulToMulOp                                                     = _ReduceType(46)
	_ReduceDivToMulOp                                                     = _ReduceType(47)
	_ReduceModToMulOp                                                     = _ReduceType(48)
	_ReduceBitAndToMulOp                                                  = _ReduceType(49)
	_ReduceBitLshiftToMulOp                                               = _ReduceType(50)
	_ReduceBitRshiftToMulOp                                               = _ReduceType(51)
	_ReducePrefixUnaryExprToMulExpr                                       = _ReduceType(52)
	_ReduceOpToMulExpr                                                    = _ReduceType(53)
	_ReduceAddToAddOp                                                     = _ReduceType(54)
	_ReduceSubToAddOp                                                     = _ReduceType(55)
	_ReduceBitOrToAddOp                                                   = _ReduceType(56)
	_ReduceBitXorToAddOp                                                  = _ReduceType(57)
	_ReduceMulExprToAddExpr                                               = _ReduceType(58)
	_ReduceOpToAddExpr                                                    = _ReduceType(59)
	_ReduceEqualToCmpOp                                                   = _ReduceType(60)
	_ReduceNotEqualToCmpOp                                                = _ReduceType(61)
	_ReduceLessToCmpOp                                                    = _ReduceType(62)
	_ReduceLessOrEqualToCmpOp                                             = _ReduceType(63)
	_ReduceGreaterToCmpOp                                                 = _ReduceType(64)
	_ReduceGreaterOrEqualToCmpOp                                          = _ReduceType(65)
	_ReduceAddExprToCmpExpr                                               = _ReduceType(66)
	_ReduceOpToCmpExpr                                                    = _ReduceType(67)
	_ReduceCmpExprToAndExpr                                               = _ReduceType(68)
	_ReduceOpToAndExpr                                                    = _ReduceType(69)
	_ReduceAndExprToOrExpr                                                = _ReduceType(70)
	_ReduceOpToOrExpr                                                     = _ReduceType(71)
	_ReduceToSequenceExpr                                                 = _ReduceType(72)
	_ReduceJumpLabelToOptionalJumpLabel                                   = _ReduceType(73)
	_ReduceUnlabelledToOptionalJumpLabel                                  = _ReduceType(74)
	_ReduceExpressionOrImplicitStructToOptionalExpressionOrImplicitStruct = _ReduceType(75)
	_ReduceNilToOptionalExpressionOrImplicitStruct                        = _ReduceType(76)
	_ReduceReturnToJumpType                                               = _ReduceType(77)
	_ReduceBreakToJumpType                                                = _ReduceType(78)
	_ReduceContinueToJumpType                                             = _ReduceType(79)
	_ReduceAddOneAssignToOpOneAssign                                      = _ReduceType(80)
	_ReduceSubOneAssignToOpOneAssign                                      = _ReduceType(81)
	_ReduceAddAssignToBinaryOpAssign                                      = _ReduceType(82)
	_ReduceSubAssignToBinaryOpAssign                                      = _ReduceType(83)
	_ReduceMulAssignToBinaryOpAssign                                      = _ReduceType(84)
	_ReduceDivAssignToBinaryOpAssign                                      = _ReduceType(85)
	_ReduceModAssignToBinaryOpAssign                                      = _ReduceType(86)
	_ReduceBitNegAssignToBinaryOpAssign                                   = _ReduceType(87)
	_ReduceBitAndAssignToBinaryOpAssign                                   = _ReduceType(88)
	_ReduceBitOrAssignToBinaryOpAssign                                    = _ReduceType(89)
	_ReduceBitXorAssignToBinaryOpAssign                                   = _ReduceType(90)
	_ReduceBitLshiftAssignToBinaryOpAssign                                = _ReduceType(91)
	_ReduceBitRshiftAssignToBinaryOpAssign                                = _ReduceType(92)
	_ReduceExpressionToExpressionOrImplicitStruct                         = _ReduceType(93)
	_ReduceImplicitStructToExpressionOrImplicitStruct                     = _ReduceType(94)
	_ReduceToUnsafeStatement                                              = _ReduceType(95)
	_ReduceUnsafeStatementToStatementBody                                 = _ReduceType(96)
	_ReduceExpressionOrImplicitStructToStatementBody                      = _ReduceType(97)
	_ReduceAsyncToStatementBody                                           = _ReduceType(98)
	_ReduceJumpToStatementBody                                            = _ReduceType(99)
	_ReduceOpOneAssignToStatementBody                                     = _ReduceType(100)
	_ReduceBinaryOpAssignToStatementBody                                  = _ReduceType(101)
	_ReduceImplicitToStatement                                            = _ReduceType(102)
	_ReduceExplicitToStatement                                            = _ReduceType(103)
	_ReduceEmptyListToStatements                                          = _ReduceType(104)
	_ReduceAddToStatements                                                = _ReduceType(105)
	_ReduceLabelDeclToOptionalLabelDecl                                   = _ReduceType(106)
	_ReduceUnlabelledToOptionalLabelDecl                                  = _ReduceType(107)
	_ReduceToBlockBody                                                    = _ReduceType(108)
	_ReduceToBlockExpr                                                    = _ReduceType(109)
	_ReduceNoElseToIfExpr                                                 = _ReduceType(110)
	_ReduceIfElseToIfExpr                                                 = _ReduceType(111)
	_ReduceMultiIfElseToIfExpr                                            = _ReduceType(112)
	_ReduceToSwitchExpr                                                   = _ReduceType(113)
	_ReduceInfiniteToLoopExpr                                             = _ReduceType(114)
	_ReduceWhileToLoopExpr                                                = _ReduceType(115)
	_ReduceSequenceExprToExpression                                       = _ReduceType(116)
	_ReduceIfExprToExpression                                             = _ReduceType(117)
	_ReduceSwitchExprToExpression                                         = _ReduceType(118)
	_ReduceLoopExprToExpression                                           = _ReduceType(119)
	_ReduceNamedToAtomType                                                = _ReduceType(120)
	_ReduceExplicitStructDefToAtomType                                    = _ReduceType(121)
	_ReduceImplicitStructDefToAtomType                                    = _ReduceType(122)
	_ReduceExplicitEnumDefToAtomType                                      = _ReduceType(123)
	_ReduceImplicitEnumDefToAtomType                                      = _ReduceType(124)
	_ReduceTraitDefToAtomType                                             = _ReduceType(125)
	_ReduceAtomTypeToTraitableType                                        = _ReduceType(126)
	_ReducePublicMethodsTraitToTraitableType                              = _ReduceType(127)
	_ReducePublicTraitToTraitableType                                     = _ReduceType(128)
	_ReduceTraitableTypeToTraitAlgebraType                                = _ReduceType(129)
	_ReduceIntersectToTraitAlgebraType                                    = _ReduceType(130)
	_ReduceUnionToTraitAlgebraType                                        = _ReduceType(131)
	_ReduceDifferenceToTraitAlgebraType                                   = _ReduceType(132)
	_ReduceInferredToValueType                                            = _ReduceType(133)
	_ReduceTraitAlgebraTypeToValueType                                    = _ReduceType(134)
	_ReduceReferenceToValueType                                           = _ReduceType(135)
	_ReduceFuncTypeToValueType                                            = _ReduceType(136)
	_ReduceDefinitionToTypeDef                                            = _ReduceType(137)
	_ReduceConstrainedDefToTypeDef                                        = _ReduceType(138)
	_ReduceAliasToTypeDef                                                 = _ReduceType(139)
	_ReduceUnconstrainedToGenericParameterDef                             = _ReduceType(140)
	_ReduceConstrainedToGenericParameterDef                               = _ReduceType(141)
	_ReduceGenericParameterDefToGenericParameterDefs                      = _ReduceType(142)
	_ReduceAddToGenericParameterDefs                                      = _ReduceType(143)
	_ReduceGenericParameterDefsToOptionalGenericParameterDefs             = _ReduceType(144)
	_ReduceNilToOptionalGenericParameterDefs                              = _ReduceType(145)
	_ReduceGenericToOptionalGenericParameters                             = _ReduceType(146)
	_ReduceNilToOptionalGenericParameters                                 = _ReduceType(147)
	_ReduceExplicitToFieldDef                                             = _ReduceType(148)
	_ReduceImplicitToFieldDef                                             = _ReduceType(149)
	_ReduceFieldDefToImplicitFieldDefs                                    = _ReduceType(150)
	_ReduceAddToImplicitFieldDefs                                         = _ReduceType(151)
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefs                   = _ReduceType(152)
	_ReduceNilToOptionalImplicitFieldDefs                                 = _ReduceType(153)
	_ReduceToImplicitStructDef                                            = _ReduceType(154)
	_ReduceFieldDefToExplicitFieldDefs                                    = _ReduceType(155)
	_ReduceImplicitToExplicitFieldDefs                                    = _ReduceType(156)
	_ReduceExplicitToExplicitFieldDefs                                    = _ReduceType(157)
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefs                   = _ReduceType(158)
	_ReduceNilToOptionalExplicitFieldDefs                                 = _ReduceType(159)
	_ReduceToExplicitStructDef                                            = _ReduceType(160)
	_ReduceFieldDefToEnumValueDef                                         = _ReduceType(161)
	_ReduceDefaultToEnumValueDef                                          = _ReduceType(162)
	_ReducePairToImplicitEnumValueDefs                                    = _ReduceType(163)
	_ReduceAddToImplicitEnumValueDefs                                     = _ReduceType(164)
	_ReduceToImplicitEnumDef                                              = _ReduceType(165)
	_ReduceExplicitPairToExplicitEnumValueDefs                            = _ReduceType(166)
	_ReduceImplicitPairToExplicitEnumValueDefs                            = _ReduceType(167)
	_ReduceExplicitAddToExplicitEnumValueDefs                             = _ReduceType(168)
	_ReduceImplicitAddToExplicitEnumValueDefs                             = _ReduceType(169)
	_ReduceToExplicitEnumDef                                              = _ReduceType(170)
	_ReduceFieldDefToTraitProperty                                        = _ReduceType(171)
	_ReduceMethodSignatureToTraitProperty                                 = _ReduceType(172)
	_ReduceTraitPropertyToTraitProperties                                 = _ReduceType(173)
	_ReduceImplicitToTraitProperties                                      = _ReduceType(174)
	_ReduceExplicitToTraitProperties                                      = _ReduceType(175)
	_ReduceTraitPropertiesToOptionalTraitProperties                       = _ReduceType(176)
	_ReduceNilToOptionalTraitProperties                                   = _ReduceType(177)
	_ReduceToTraitDef                                                     = _ReduceType(178)
	_ReduceValueTypeToReturnType                                          = _ReduceType(179)
	_ReduceNilToReturnType                                                = _ReduceType(180)
	_ReduceArgToParameterDecl                                             = _ReduceType(181)
	_ReduceVarargToParameterDecl                                          = _ReduceType(182)
	_ReduceUnamedToParameterDecl                                          = _ReduceType(183)
	_ReduceUnnamedVarargToParameterDecl                                   = _ReduceType(184)
	_ReduceParameterDeclToParameterDecls                                  = _ReduceType(185)
	_ReduceAddToParameterDecls                                            = _ReduceType(186)
	_ReduceParameterDeclsToOptionalParameterDecls                         = _ReduceType(187)
	_ReduceNilToOptionalParameterDecls                                    = _ReduceType(188)
	_ReduceToFuncType                                                     = _ReduceType(189)
	_ReduceToMethodSignature                                              = _ReduceType(190)
	_ReduceArgToParameterDef                                              = _ReduceType(191)
	_ReduceVarargToParameterDef                                           = _ReduceType(192)
	_ReduceParameterDefToParameterDefs                                    = _ReduceType(193)
	_ReduceAddToParameterDefs                                             = _ReduceType(194)
	_ReduceParameterDefsToOptionalParameterDefs                           = _ReduceType(195)
	_ReduceNilToOptionalParameterDefs                                     = _ReduceType(196)
	_ReduceReceiverToOptionalReceiver                                     = _ReduceType(197)
	_ReduceNilToOptionalReceiver                                          = _ReduceType(198)
	_ReduceToNamedFuncDef                                                 = _ReduceType(199)
	_ReduceToAnonymousFuncExpr                                            = _ReduceType(200)
	_ReduceNoSpecToPackageDef                                             = _ReduceType(201)
	_ReduceWithSpecToPackageDef                                           = _ReduceType(202)
	_ReduceToPackageStatementBody                                         = _ReduceType(203)
	_ReduceImplicitToPackageStatement                                     = _ReduceType(204)
	_ReduceExplicitToPackageStatement                                     = _ReduceType(205)
	_ReduceEmptyListToPackageStatements                                   = _ReduceType(206)
	_ReduceAddToPackageStatements                                         = _ReduceType(207)
	_ReduceSpacesToLexInternalTokens                                      = _ReduceType(208)
	_ReduceCommentToLexInternalTokens                                     = _ReduceType(209)
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
	case _ReduceExpressionToOptionalExpression:
		return "ExpressionToOptionalExpression"
	case _ReduceNilToOptionalExpression:
		return "NilToOptionalExpression"
	case _ReducePairToColonExpressions:
		return "PairToColonExpressions"
	case _ReduceAddToColonExpressions:
		return "AddToColonExpressions"
	case _ReducePositionalToArgument:
		return "PositionalToArgument"
	case _ReduceNamedToArgument:
		return "NamedToArgument"
	case _ReduceColonExpressionsToArgument:
		return "ColonExpressionsToArgument"
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
	case _ReduceAtomTypeToTraitableType:
		return "AtomTypeToTraitableType"
	case _ReducePublicMethodsTraitToTraitableType:
		return "PublicMethodsTraitToTraitableType"
	case _ReducePublicTraitToTraitableType:
		return "PublicTraitToTraitableType"
	case _ReduceTraitableTypeToTraitAlgebraType:
		return "TraitableTypeToTraitAlgebraType"
	case _ReduceIntersectToTraitAlgebraType:
		return "IntersectToTraitAlgebraType"
	case _ReduceUnionToTraitAlgebraType:
		return "UnionToTraitAlgebraType"
	case _ReduceDifferenceToTraitAlgebraType:
		return "DifferenceToTraitAlgebraType"
	case _ReduceInferredToValueType:
		return "InferredToValueType"
	case _ReduceTraitAlgebraTypeToValueType:
		return "TraitAlgebraTypeToValueType"
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
	case _ReduceFieldDefToEnumValueDef:
		return "FieldDefToEnumValueDef"
	case _ReduceDefaultToEnumValueDef:
		return "DefaultToEnumValueDef"
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
	_State320 = _StateId(320)
	_State321 = _StateId(321)
	_State322 = _StateId(322)
	_State323 = _StateId(323)
	_State324 = _StateId(324)
	_State325 = _StateId(325)
	_State326 = _StateId(326)
	_State327 = _StateId(327)
	_State328 = _StateId(328)
	_State329 = _StateId(329)
	_State330 = _StateId(330)
	_State331 = _StateId(331)
	_State332 = _StateId(332)
	_State333 = _StateId(333)
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
	case _EndMarker, SpacesToken, NewlinesToken, CommentToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, ReturnToken, BreakToken, ContinueToken, PackageToken, UnsafeToken, TypeToken, StructToken, EnumToken, TraitToken, ImplementsToken, FuncToken, AsyncToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, DollarLbracketToken, DotdotdotToken, TildeTildeToken, LabelDeclToken, JumpLabelToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, NotToken, AndToken, OrToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, LexErrorToken:
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
	case _ReduceExpressionToOptionalExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalExpressionType
		symbol.Generic_, err = reducer.ExpressionToOptionalExpression(args[0].Generic_)
	case _ReduceNilToOptionalExpression:
		symbol.SymbolId_ = OptionalExpressionType
		symbol.Generic_, err = reducer.NilToOptionalExpression()
	case _ReducePairToColonExpressions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ColonExpressionsType
		symbol.Generic_, err = reducer.PairToColonExpressions(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceAddToColonExpressions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ColonExpressionsType
		symbol.Generic_, err = reducer.AddToColonExpressions(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReducePositionalToArgument:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentType
		symbol.Generic_, err = reducer.PositionalToArgument(args[0].Generic_)
	case _ReduceNamedToArgument:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ArgumentType
		symbol.Generic_, err = reducer.NamedToArgument(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceColonExpressionsToArgument:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentType
		symbol.Generic_, err = reducer.ColonExpressionsToArgument(args[0].Generic_)
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
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = SwitchExprType
		symbol.Generic_, err = reducer.ToSwitchExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_)
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
	case _ReduceAtomTypeToTraitableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitableTypeType
		symbol.Generic_, err = reducer.AtomTypeToTraitableType(args[0].Generic_)
	case _ReducePublicMethodsTraitToTraitableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = TraitableTypeType
		symbol.Generic_, err = reducer.PublicMethodsTraitToTraitableType(args[0].Generic_, args[1].Generic_)
	case _ReducePublicTraitToTraitableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = TraitableTypeType
		symbol.Generic_, err = reducer.PublicTraitToTraitableType(args[0].Generic_, args[1].Generic_)
	case _ReduceTraitableTypeToTraitAlgebraType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitAlgebraTypeType
		symbol.Generic_, err = reducer.TraitableTypeToTraitAlgebraType(args[0].Generic_)
	case _ReduceIntersectToTraitAlgebraType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TraitAlgebraTypeType
		symbol.Generic_, err = reducer.IntersectToTraitAlgebraType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceUnionToTraitAlgebraType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TraitAlgebraTypeType
		symbol.Generic_, err = reducer.UnionToTraitAlgebraType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceDifferenceToTraitAlgebraType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TraitAlgebraTypeType
		symbol.Generic_, err = reducer.DifferenceToTraitAlgebraType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceInferredToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.InferredToValueType(args[0].Generic_)
	case _ReduceTraitAlgebraTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.TraitAlgebraTypeToValueType(args[0].Generic_)
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
	case _ReduceFieldDefToEnumValueDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = EnumValueDefType
		symbol.Generic_, err = reducer.FieldDefToEnumValueDef(args[0].Generic_)
	case _ReduceDefaultToEnumValueDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = EnumValueDefType
		symbol.Generic_, err = reducer.DefaultToEnumValueDef(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	_GotoState320Action                                                         = &_Action{_ShiftAction, _State320, 0}
	_GotoState321Action                                                         = &_Action{_ShiftAction, _State321, 0}
	_GotoState322Action                                                         = &_Action{_ShiftAction, _State322, 0}
	_GotoState323Action                                                         = &_Action{_ShiftAction, _State323, 0}
	_GotoState324Action                                                         = &_Action{_ShiftAction, _State324, 0}
	_GotoState325Action                                                         = &_Action{_ShiftAction, _State325, 0}
	_GotoState326Action                                                         = &_Action{_ShiftAction, _State326, 0}
	_GotoState327Action                                                         = &_Action{_ShiftAction, _State327, 0}
	_GotoState328Action                                                         = &_Action{_ShiftAction, _State328, 0}
	_GotoState329Action                                                         = &_Action{_ShiftAction, _State329, 0}
	_GotoState330Action                                                         = &_Action{_ShiftAction, _State330, 0}
	_GotoState331Action                                                         = &_Action{_ShiftAction, _State331, 0}
	_GotoState332Action                                                         = &_Action{_ShiftAction, _State332, 0}
	_GotoState333Action                                                         = &_Action{_ShiftAction, _State333, 0}
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
	_ReduceExpressionToOptionalExpressionAction                                 = &_Action{_ReduceAction, 0, _ReduceExpressionToOptionalExpression}
	_ReduceNilToOptionalExpressionAction                                        = &_Action{_ReduceAction, 0, _ReduceNilToOptionalExpression}
	_ReducePairToColonExpressionsAction                                         = &_Action{_ReduceAction, 0, _ReducePairToColonExpressions}
	_ReduceAddToColonExpressionsAction                                          = &_Action{_ReduceAction, 0, _ReduceAddToColonExpressions}
	_ReducePositionalToArgumentAction                                           = &_Action{_ReduceAction, 0, _ReducePositionalToArgument}
	_ReduceNamedToArgumentAction                                                = &_Action{_ReduceAction, 0, _ReduceNamedToArgument}
	_ReduceColonExpressionsToArgumentAction                                     = &_Action{_ReduceAction, 0, _ReduceColonExpressionsToArgument}
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
	_ReduceAtomTypeToTraitableTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceAtomTypeToTraitableType}
	_ReducePublicMethodsTraitToTraitableTypeAction                              = &_Action{_ReduceAction, 0, _ReducePublicMethodsTraitToTraitableType}
	_ReducePublicTraitToTraitableTypeAction                                     = &_Action{_ReduceAction, 0, _ReducePublicTraitToTraitableType}
	_ReduceTraitableTypeToTraitAlgebraTypeAction                                = &_Action{_ReduceAction, 0, _ReduceTraitableTypeToTraitAlgebraType}
	_ReduceIntersectToTraitAlgebraTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceIntersectToTraitAlgebraType}
	_ReduceUnionToTraitAlgebraTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceUnionToTraitAlgebraType}
	_ReduceDifferenceToTraitAlgebraTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceDifferenceToTraitAlgebraType}
	_ReduceInferredToValueTypeAction                                            = &_Action{_ReduceAction, 0, _ReduceInferredToValueType}
	_ReduceTraitAlgebraTypeToValueTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceTraitAlgebraTypeToValueType}
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
	_ReduceFieldDefToEnumValueDefAction                                         = &_Action{_ReduceAction, 0, _ReduceFieldDefToEnumValueDef}
	_ReduceDefaultToEnumValueDefAction                                          = &_Action{_ReduceAction, 0, _ReduceDefaultToEnumValueDef}
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
	{_State26, IdentifierToken}:                         _GotoState60Action,
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
	{_State26, OptionalExpressionType}:                  _GotoState65Action,
	{_State26, ColonExpressionsType}:                    _GotoState63Action,
	{_State26, ArgumentType}:                            _GotoState61Action,
	{_State26, ArgumentsType}:                           _GotoState62Action,
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
	{_State26, ExpressionType}:                          _GotoState64Action,
	{_State26, ExplicitStructDefType}:                   _GotoState43Action,
	{_State26, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State31, LparenToken}:                             _GotoState66Action,
	{_State34, LbracketToken}:                           _GotoState69Action,
	{_State34, DotToken}:                                _GotoState68Action,
	{_State34, QuestionToken}:                           _GotoState70Action,
	{_State34, DollarLbracketToken}:                     _GotoState67Action,
	{_State34, OptionalGenericBindingType}:              _GotoState71Action,
	{_State35, AddToken}:                                _GotoState72Action,
	{_State35, SubToken}:                                _GotoState75Action,
	{_State35, BitXorToken}:                             _GotoState74Action,
	{_State35, BitOrToken}:                              _GotoState73Action,
	{_State35, AddOpType}:                               _GotoState76Action,
	{_State36, AndToken}:                                _GotoState77Action,
	{_State42, EqualToken}:                              _GotoState78Action,
	{_State42, NotEqualToken}:                           _GotoState83Action,
	{_State42, LessToken}:                               _GotoState81Action,
	{_State42, LessOrEqualToken}:                        _GotoState82Action,
	{_State42, GreaterToken}:                            _GotoState79Action,
	{_State42, GreaterOrEqualToken}:                     _GotoState80Action,
	{_State42, CmpOpType}:                               _GotoState84Action,
	{_State43, LparenToken}:                             _GotoState85Action,
	{_State45, MulToken}:                                _GotoState91Action,
	{_State45, DivToken}:                                _GotoState89Action,
	{_State45, ModToken}:                                _GotoState90Action,
	{_State45, BitAndToken}:                             _GotoState86Action,
	{_State45, BitLshiftToken}:                          _GotoState87Action,
	{_State45, BitRshiftToken}:                          _GotoState88Action,
	{_State45, MulOpType}:                               _GotoState92Action,
	{_State46, IfToken}:                                 _GotoState94Action,
	{_State46, SwitchToken}:                             _GotoState96Action,
	{_State46, ForToken}:                                _GotoState93Action,
	{_State46, LbraceToken}:                             _GotoState95Action,
	{_State46, BlockBodyType}:                           _GotoState97Action,
	{_State46, IfExprType}:                              _GotoState98Action,
	{_State46, SwitchExprType}:                          _GotoState100Action,
	{_State46, LoopExprType}:                            _GotoState99Action,
	{_State47, OrToken}:                                 _GotoState101Action,
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
	{_State50, PrefixUnaryExprType}:                     _GotoState103Action,
	{_State50, OptionalLabelDeclType}:                   _GotoState102Action,
	{_State50, BlockExprType}:                           _GotoState40Action,
	{_State50, ExplicitStructDefType}:                   _GotoState43Action,
	{_State50, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State54, LparenToken}:                             _GotoState104Action,
	{_State55, DollarLbracketToken}:                     _GotoState105Action,
	{_State55, EqualToken}:                              _GotoState106Action,
	{_State55, OptionalGenericParametersType}:           _GotoState107Action,
	{_State56, IdentifierToken}:                         _GotoState112Action,
	{_State56, StructToken}:                             _GotoState31Action,
	{_State56, EnumToken}:                               _GotoState110Action,
	{_State56, TraitToken}:                              _GotoState15Action,
	{_State56, FuncToken}:                               _GotoState111Action,
	{_State56, LparenToken}:                             _GotoState113Action,
	{_State56, QuestionToken}:                           _GotoState114Action,
	{_State56, TildeTildeToken}:                         _GotoState115Action,
	{_State56, BitNegToken}:                             _GotoState109Action,
	{_State56, BitAndToken}:                             _GotoState108Action,
	{_State56, AtomTypeType}:                            _GotoState116Action,
	{_State56, TraitableTypeType}:                       _GotoState129Action,
	{_State56, TraitAlgebraTypeType}:                    _GotoState125Action,
	{_State56, ValueTypeType}:                           _GotoState130Action,
	{_State56, FieldDefType}:                            _GotoState119Action,
	{_State56, ImplicitStructDefType}:                   _GotoState122Action,
	{_State56, ExplicitStructDefType}:                   _GotoState118Action,
	{_State56, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State56, ExplicitEnumDefType}:                     _GotoState117Action,
	{_State56, TraitPropertyType}:                       _GotoState128Action,
	{_State56, TraitPropertiesType}:                     _GotoState127Action,
	{_State56, OptionalTraitPropertiesType}:             _GotoState124Action,
	{_State56, TraitDefType}:                            _GotoState126Action,
	{_State56, FuncTypeType}:                            _GotoState120Action,
	{_State56, MethodSignatureType}:                     _GotoState123Action,
	{_State57, IdentifierToken}:                         _GotoState131Action,
	{_State57, ParameterDefType}:                        _GotoState132Action,
	{_State58, IdentifierToken}:                         _GotoState133Action,
	{_State59, IdentifierToken}:                         _GotoState131Action,
	{_State59, ParameterDefType}:                        _GotoState135Action,
	{_State59, ParameterDefsType}:                       _GotoState136Action,
	{_State59, OptionalParameterDefsType}:               _GotoState134Action,
	{_State60, AssignToken}:                             _GotoState137Action,
	{_State62, RparenToken}:                             _GotoState139Action,
	{_State62, CommaToken}:                              _GotoState138Action,
	{_State63, ColonToken}:                              _GotoState140Action,
	{_State65, ColonToken}:                              _GotoState141Action,
	{_State66, IdentifierToken}:                         _GotoState112Action,
	{_State66, StructToken}:                             _GotoState31Action,
	{_State66, EnumToken}:                               _GotoState110Action,
	{_State66, TraitToken}:                              _GotoState15Action,
	{_State66, FuncToken}:                               _GotoState142Action,
	{_State66, LparenToken}:                             _GotoState113Action,
	{_State66, QuestionToken}:                           _GotoState114Action,
	{_State66, TildeTildeToken}:                         _GotoState115Action,
	{_State66, BitNegToken}:                             _GotoState109Action,
	{_State66, BitAndToken}:                             _GotoState108Action,
	{_State66, AtomTypeType}:                            _GotoState116Action,
	{_State66, TraitableTypeType}:                       _GotoState129Action,
	{_State66, TraitAlgebraTypeType}:                    _GotoState125Action,
	{_State66, ValueTypeType}:                           _GotoState130Action,
	{_State66, FieldDefType}:                            _GotoState144Action,
	{_State66, ImplicitStructDefType}:                   _GotoState122Action,
	{_State66, ExplicitFieldDefsType}:                   _GotoState143Action,
	{_State66, OptionalExplicitFieldDefsType}:           _GotoState145Action,
	{_State66, ExplicitStructDefType}:                   _GotoState118Action,
	{_State66, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State66, ExplicitEnumDefType}:                     _GotoState117Action,
	{_State66, TraitDefType}:                            _GotoState126Action,
	{_State66, FuncTypeType}:                            _GotoState120Action,
	{_State67, IdentifierToken}:                         _GotoState146Action,
	{_State67, StructToken}:                             _GotoState31Action,
	{_State67, EnumToken}:                               _GotoState110Action,
	{_State67, TraitToken}:                              _GotoState15Action,
	{_State67, FuncToken}:                               _GotoState142Action,
	{_State67, LparenToken}:                             _GotoState113Action,
	{_State67, QuestionToken}:                           _GotoState114Action,
	{_State67, TildeTildeToken}:                         _GotoState115Action,
	{_State67, BitNegToken}:                             _GotoState109Action,
	{_State67, BitAndToken}:                             _GotoState108Action,
	{_State67, GenericArgumentsType}:                    _GotoState147Action,
	{_State67, OptionalGenericArgumentsType}:            _GotoState148Action,
	{_State67, AtomTypeType}:                            _GotoState116Action,
	{_State67, TraitableTypeType}:                       _GotoState129Action,
	{_State67, TraitAlgebraTypeType}:                    _GotoState125Action,
	{_State67, ValueTypeType}:                           _GotoState149Action,
	{_State67, ImplicitStructDefType}:                   _GotoState122Action,
	{_State67, ExplicitStructDefType}:                   _GotoState118Action,
	{_State67, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State67, ExplicitEnumDefType}:                     _GotoState117Action,
	{_State67, TraitDefType}:                            _GotoState126Action,
	{_State67, FuncTypeType}:                            _GotoState120Action,
	{_State68, IdentifierToken}:                         _GotoState150Action,
	{_State69, IntegerLiteralToken}:                     _GotoState23Action,
	{_State69, FloatLiteralToken}:                       _GotoState20Action,
	{_State69, RuneLiteralToken}:                        _GotoState29Action,
	{_State69, StringLiteralToken}:                      _GotoState30Action,
	{_State69, IdentifierToken}:                         _GotoState60Action,
	{_State69, TrueToken}:                               _GotoState33Action,
	{_State69, FalseToken}:                              _GotoState19Action,
	{_State69, StructToken}:                             _GotoState31Action,
	{_State69, FuncToken}:                               _GotoState21Action,
	{_State69, LparenToken}:                             _GotoState26Action,
	{_State69, LabelDeclToken}:                          _GotoState24Action,
	{_State69, NotToken}:                                _GotoState28Action,
	{_State69, SubToken}:                                _GotoState32Action,
	{_State69, MulToken}:                                _GotoState27Action,
	{_State69, BitNegToken}:                             _GotoState18Action,
	{_State69, BitAndToken}:                             _GotoState17Action,
	{_State69, LexErrorToken}:                           _GotoState25Action,
	{_State69, LiteralType}:                             _GotoState44Action,
	{_State69, AnonymousStructExprType}:                 _GotoState38Action,
	{_State69, AtomExprType}:                            _GotoState39Action,
	{_State69, OptionalExpressionType}:                  _GotoState65Action,
	{_State69, ColonExpressionsType}:                    _GotoState63Action,
	{_State69, ArgumentType}:                            _GotoState151Action,
	{_State69, CallExprType}:                            _GotoState41Action,
	{_State69, AccessExprType}:                          _GotoState34Action,
	{_State69, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State69, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State69, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State69, MulExprType}:                             _GotoState45Action,
	{_State69, AddExprType}:                             _GotoState35Action,
	{_State69, CmpExprType}:                             _GotoState42Action,
	{_State69, AndExprType}:                             _GotoState36Action,
	{_State69, OrExprType}:                              _GotoState47Action,
	{_State69, SequenceExprType}:                        _GotoState51Action,
	{_State69, OptionalLabelDeclType}:                   _GotoState46Action,
	{_State69, BlockExprType}:                           _GotoState40Action,
	{_State69, ExpressionType}:                          _GotoState64Action,
	{_State69, ExplicitStructDefType}:                   _GotoState43Action,
	{_State69, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State71, LparenToken}:                             _GotoState152Action,
	{_State76, IntegerLiteralToken}:                     _GotoState23Action,
	{_State76, FloatLiteralToken}:                       _GotoState20Action,
	{_State76, RuneLiteralToken}:                        _GotoState29Action,
	{_State76, StringLiteralToken}:                      _GotoState30Action,
	{_State76, IdentifierToken}:                         _GotoState22Action,
	{_State76, TrueToken}:                               _GotoState33Action,
	{_State76, FalseToken}:                              _GotoState19Action,
	{_State76, StructToken}:                             _GotoState31Action,
	{_State76, FuncToken}:                               _GotoState21Action,
	{_State76, LparenToken}:                             _GotoState26Action,
	{_State76, LabelDeclToken}:                          _GotoState24Action,
	{_State76, NotToken}:                                _GotoState28Action,
	{_State76, SubToken}:                                _GotoState32Action,
	{_State76, MulToken}:                                _GotoState27Action,
	{_State76, BitNegToken}:                             _GotoState18Action,
	{_State76, BitAndToken}:                             _GotoState17Action,
	{_State76, LexErrorToken}:                           _GotoState25Action,
	{_State76, LiteralType}:                             _GotoState44Action,
	{_State76, AnonymousStructExprType}:                 _GotoState38Action,
	{_State76, AtomExprType}:                            _GotoState39Action,
	{_State76, CallExprType}:                            _GotoState41Action,
	{_State76, AccessExprType}:                          _GotoState34Action,
	{_State76, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State76, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State76, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State76, MulExprType}:                             _GotoState153Action,
	{_State76, OptionalLabelDeclType}:                   _GotoState102Action,
	{_State76, BlockExprType}:                           _GotoState40Action,
	{_State76, ExplicitStructDefType}:                   _GotoState43Action,
	{_State76, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State77, IntegerLiteralToken}:                     _GotoState23Action,
	{_State77, FloatLiteralToken}:                       _GotoState20Action,
	{_State77, RuneLiteralToken}:                        _GotoState29Action,
	{_State77, StringLiteralToken}:                      _GotoState30Action,
	{_State77, IdentifierToken}:                         _GotoState22Action,
	{_State77, TrueToken}:                               _GotoState33Action,
	{_State77, FalseToken}:                              _GotoState19Action,
	{_State77, StructToken}:                             _GotoState31Action,
	{_State77, FuncToken}:                               _GotoState21Action,
	{_State77, LparenToken}:                             _GotoState26Action,
	{_State77, LabelDeclToken}:                          _GotoState24Action,
	{_State77, NotToken}:                                _GotoState28Action,
	{_State77, SubToken}:                                _GotoState32Action,
	{_State77, MulToken}:                                _GotoState27Action,
	{_State77, BitNegToken}:                             _GotoState18Action,
	{_State77, BitAndToken}:                             _GotoState17Action,
	{_State77, LexErrorToken}:                           _GotoState25Action,
	{_State77, LiteralType}:                             _GotoState44Action,
	{_State77, AnonymousStructExprType}:                 _GotoState38Action,
	{_State77, AtomExprType}:                            _GotoState39Action,
	{_State77, CallExprType}:                            _GotoState41Action,
	{_State77, AccessExprType}:                          _GotoState34Action,
	{_State77, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State77, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State77, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State77, MulExprType}:                             _GotoState45Action,
	{_State77, AddExprType}:                             _GotoState35Action,
	{_State77, CmpExprType}:                             _GotoState154Action,
	{_State77, OptionalLabelDeclType}:                   _GotoState102Action,
	{_State77, BlockExprType}:                           _GotoState40Action,
	{_State77, ExplicitStructDefType}:                   _GotoState43Action,
	{_State77, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State84, IntegerLiteralToken}:                     _GotoState23Action,
	{_State84, FloatLiteralToken}:                       _GotoState20Action,
	{_State84, RuneLiteralToken}:                        _GotoState29Action,
	{_State84, StringLiteralToken}:                      _GotoState30Action,
	{_State84, IdentifierToken}:                         _GotoState22Action,
	{_State84, TrueToken}:                               _GotoState33Action,
	{_State84, FalseToken}:                              _GotoState19Action,
	{_State84, StructToken}:                             _GotoState31Action,
	{_State84, FuncToken}:                               _GotoState21Action,
	{_State84, LparenToken}:                             _GotoState26Action,
	{_State84, LabelDeclToken}:                          _GotoState24Action,
	{_State84, NotToken}:                                _GotoState28Action,
	{_State84, SubToken}:                                _GotoState32Action,
	{_State84, MulToken}:                                _GotoState27Action,
	{_State84, BitNegToken}:                             _GotoState18Action,
	{_State84, BitAndToken}:                             _GotoState17Action,
	{_State84, LexErrorToken}:                           _GotoState25Action,
	{_State84, LiteralType}:                             _GotoState44Action,
	{_State84, AnonymousStructExprType}:                 _GotoState38Action,
	{_State84, AtomExprType}:                            _GotoState39Action,
	{_State84, CallExprType}:                            _GotoState41Action,
	{_State84, AccessExprType}:                          _GotoState34Action,
	{_State84, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State84, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State84, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State84, MulExprType}:                             _GotoState45Action,
	{_State84, AddExprType}:                             _GotoState155Action,
	{_State84, OptionalLabelDeclType}:                   _GotoState102Action,
	{_State84, BlockExprType}:                           _GotoState40Action,
	{_State84, ExplicitStructDefType}:                   _GotoState43Action,
	{_State84, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State85, IntegerLiteralToken}:                     _GotoState23Action,
	{_State85, FloatLiteralToken}:                       _GotoState20Action,
	{_State85, RuneLiteralToken}:                        _GotoState29Action,
	{_State85, StringLiteralToken}:                      _GotoState30Action,
	{_State85, IdentifierToken}:                         _GotoState60Action,
	{_State85, TrueToken}:                               _GotoState33Action,
	{_State85, FalseToken}:                              _GotoState19Action,
	{_State85, StructToken}:                             _GotoState31Action,
	{_State85, FuncToken}:                               _GotoState21Action,
	{_State85, LparenToken}:                             _GotoState26Action,
	{_State85, LabelDeclToken}:                          _GotoState24Action,
	{_State85, NotToken}:                                _GotoState28Action,
	{_State85, SubToken}:                                _GotoState32Action,
	{_State85, MulToken}:                                _GotoState27Action,
	{_State85, BitNegToken}:                             _GotoState18Action,
	{_State85, BitAndToken}:                             _GotoState17Action,
	{_State85, LexErrorToken}:                           _GotoState25Action,
	{_State85, LiteralType}:                             _GotoState44Action,
	{_State85, AnonymousStructExprType}:                 _GotoState38Action,
	{_State85, AtomExprType}:                            _GotoState39Action,
	{_State85, OptionalExpressionType}:                  _GotoState65Action,
	{_State85, ColonExpressionsType}:                    _GotoState63Action,
	{_State85, ArgumentType}:                            _GotoState61Action,
	{_State85, ArgumentsType}:                           _GotoState156Action,
	{_State85, CallExprType}:                            _GotoState41Action,
	{_State85, AccessExprType}:                          _GotoState34Action,
	{_State85, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State85, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State85, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State85, MulExprType}:                             _GotoState45Action,
	{_State85, AddExprType}:                             _GotoState35Action,
	{_State85, CmpExprType}:                             _GotoState42Action,
	{_State85, AndExprType}:                             _GotoState36Action,
	{_State85, OrExprType}:                              _GotoState47Action,
	{_State85, SequenceExprType}:                        _GotoState51Action,
	{_State85, OptionalLabelDeclType}:                   _GotoState46Action,
	{_State85, BlockExprType}:                           _GotoState40Action,
	{_State85, ExpressionType}:                          _GotoState64Action,
	{_State85, ExplicitStructDefType}:                   _GotoState43Action,
	{_State85, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State92, IntegerLiteralToken}:                     _GotoState23Action,
	{_State92, FloatLiteralToken}:                       _GotoState20Action,
	{_State92, RuneLiteralToken}:                        _GotoState29Action,
	{_State92, StringLiteralToken}:                      _GotoState30Action,
	{_State92, IdentifierToken}:                         _GotoState22Action,
	{_State92, TrueToken}:                               _GotoState33Action,
	{_State92, FalseToken}:                              _GotoState19Action,
	{_State92, StructToken}:                             _GotoState31Action,
	{_State92, FuncToken}:                               _GotoState21Action,
	{_State92, LparenToken}:                             _GotoState26Action,
	{_State92, LabelDeclToken}:                          _GotoState24Action,
	{_State92, NotToken}:                                _GotoState28Action,
	{_State92, SubToken}:                                _GotoState32Action,
	{_State92, MulToken}:                                _GotoState27Action,
	{_State92, BitNegToken}:                             _GotoState18Action,
	{_State92, BitAndToken}:                             _GotoState17Action,
	{_State92, LexErrorToken}:                           _GotoState25Action,
	{_State92, LiteralType}:                             _GotoState44Action,
	{_State92, AnonymousStructExprType}:                 _GotoState38Action,
	{_State92, AtomExprType}:                            _GotoState39Action,
	{_State92, CallExprType}:                            _GotoState41Action,
	{_State92, AccessExprType}:                          _GotoState34Action,
	{_State92, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State92, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State92, PrefixUnaryExprType}:                     _GotoState157Action,
	{_State92, OptionalLabelDeclType}:                   _GotoState102Action,
	{_State92, BlockExprType}:                           _GotoState40Action,
	{_State92, ExplicitStructDefType}:                   _GotoState43Action,
	{_State92, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State93, IntegerLiteralToken}:                     _GotoState23Action,
	{_State93, FloatLiteralToken}:                       _GotoState20Action,
	{_State93, RuneLiteralToken}:                        _GotoState29Action,
	{_State93, StringLiteralToken}:                      _GotoState30Action,
	{_State93, IdentifierToken}:                         _GotoState22Action,
	{_State93, TrueToken}:                               _GotoState33Action,
	{_State93, FalseToken}:                              _GotoState19Action,
	{_State93, StructToken}:                             _GotoState31Action,
	{_State93, FuncToken}:                               _GotoState21Action,
	{_State93, LparenToken}:                             _GotoState26Action,
	{_State93, LabelDeclToken}:                          _GotoState24Action,
	{_State93, NotToken}:                                _GotoState28Action,
	{_State93, SubToken}:                                _GotoState32Action,
	{_State93, MulToken}:                                _GotoState27Action,
	{_State93, BitNegToken}:                             _GotoState18Action,
	{_State93, BitAndToken}:                             _GotoState17Action,
	{_State93, LexErrorToken}:                           _GotoState25Action,
	{_State93, LiteralType}:                             _GotoState44Action,
	{_State93, AnonymousStructExprType}:                 _GotoState38Action,
	{_State93, AtomExprType}:                            _GotoState39Action,
	{_State93, CallExprType}:                            _GotoState41Action,
	{_State93, AccessExprType}:                          _GotoState34Action,
	{_State93, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State93, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State93, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State93, MulExprType}:                             _GotoState45Action,
	{_State93, AddExprType}:                             _GotoState35Action,
	{_State93, CmpExprType}:                             _GotoState42Action,
	{_State93, AndExprType}:                             _GotoState36Action,
	{_State93, OrExprType}:                              _GotoState47Action,
	{_State93, SequenceExprType}:                        _GotoState159Action,
	{_State93, OptionalLabelDeclType}:                   _GotoState102Action,
	{_State93, BlockExprType}:                           _GotoState158Action,
	{_State93, ExplicitStructDefType}:                   _GotoState43Action,
	{_State93, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State94, IntegerLiteralToken}:                     _GotoState23Action,
	{_State94, FloatLiteralToken}:                       _GotoState20Action,
	{_State94, RuneLiteralToken}:                        _GotoState29Action,
	{_State94, StringLiteralToken}:                      _GotoState30Action,
	{_State94, IdentifierToken}:                         _GotoState22Action,
	{_State94, TrueToken}:                               _GotoState33Action,
	{_State94, FalseToken}:                              _GotoState19Action,
	{_State94, StructToken}:                             _GotoState31Action,
	{_State94, FuncToken}:                               _GotoState21Action,
	{_State94, LparenToken}:                             _GotoState26Action,
	{_State94, LabelDeclToken}:                          _GotoState24Action,
	{_State94, NotToken}:                                _GotoState28Action,
	{_State94, SubToken}:                                _GotoState32Action,
	{_State94, MulToken}:                                _GotoState27Action,
	{_State94, BitNegToken}:                             _GotoState18Action,
	{_State94, BitAndToken}:                             _GotoState17Action,
	{_State94, LexErrorToken}:                           _GotoState25Action,
	{_State94, LiteralType}:                             _GotoState44Action,
	{_State94, AnonymousStructExprType}:                 _GotoState38Action,
	{_State94, AtomExprType}:                            _GotoState39Action,
	{_State94, CallExprType}:                            _GotoState41Action,
	{_State94, AccessExprType}:                          _GotoState34Action,
	{_State94, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State94, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State94, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State94, MulExprType}:                             _GotoState45Action,
	{_State94, AddExprType}:                             _GotoState35Action,
	{_State94, CmpExprType}:                             _GotoState42Action,
	{_State94, AndExprType}:                             _GotoState36Action,
	{_State94, OrExprType}:                              _GotoState47Action,
	{_State94, SequenceExprType}:                        _GotoState160Action,
	{_State94, OptionalLabelDeclType}:                   _GotoState102Action,
	{_State94, BlockExprType}:                           _GotoState40Action,
	{_State94, ExplicitStructDefType}:                   _GotoState43Action,
	{_State94, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State95, StatementsType}:                          _GotoState161Action,
	{_State96, IntegerLiteralToken}:                     _GotoState23Action,
	{_State96, FloatLiteralToken}:                       _GotoState20Action,
	{_State96, RuneLiteralToken}:                        _GotoState29Action,
	{_State96, StringLiteralToken}:                      _GotoState30Action,
	{_State96, IdentifierToken}:                         _GotoState22Action,
	{_State96, TrueToken}:                               _GotoState33Action,
	{_State96, FalseToken}:                              _GotoState19Action,
	{_State96, StructToken}:                             _GotoState31Action,
	{_State96, FuncToken}:                               _GotoState21Action,
	{_State96, LparenToken}:                             _GotoState26Action,
	{_State96, LabelDeclToken}:                          _GotoState24Action,
	{_State96, NotToken}:                                _GotoState28Action,
	{_State96, SubToken}:                                _GotoState32Action,
	{_State96, MulToken}:                                _GotoState27Action,
	{_State96, BitNegToken}:                             _GotoState18Action,
	{_State96, BitAndToken}:                             _GotoState17Action,
	{_State96, LexErrorToken}:                           _GotoState25Action,
	{_State96, LiteralType}:                             _GotoState44Action,
	{_State96, AnonymousStructExprType}:                 _GotoState38Action,
	{_State96, AtomExprType}:                            _GotoState39Action,
	{_State96, CallExprType}:                            _GotoState41Action,
	{_State96, AccessExprType}:                          _GotoState34Action,
	{_State96, PostfixUnaryExprType}:                    _GotoState48Action,
	{_State96, PrefixUnaryOpType}:                       _GotoState50Action,
	{_State96, PrefixUnaryExprType}:                     _GotoState49Action,
	{_State96, MulExprType}:                             _GotoState45Action,
	{_State96, AddExprType}:                             _GotoState35Action,
	{_State96, CmpExprType}:                             _GotoState42Action,
	{_State96, AndExprType}:                             _GotoState36Action,
	{_State96, OrExprType}:                              _GotoState47Action,
	{_State96, SequenceExprType}:                        _GotoState162Action,
	{_State96, OptionalLabelDeclType}:                   _GotoState102Action,
	{_State96, BlockExprType}:                           _GotoState40Action,
	{_State96, ExplicitStructDefType}:                   _GotoState43Action,
	{_State96, AnonymousFuncExprType}:                   _GotoState37Action,
	{_State101, IntegerLiteralToken}:                    _GotoState23Action,
	{_State101, FloatLiteralToken}:                      _GotoState20Action,
	{_State101, RuneLiteralToken}:                       _GotoState29Action,
	{_State101, StringLiteralToken}:                     _GotoState30Action,
	{_State101, IdentifierToken}:                        _GotoState22Action,
	{_State101, TrueToken}:                              _GotoState33Action,
	{_State101, FalseToken}:                             _GotoState19Action,
	{_State101, StructToken}:                            _GotoState31Action,
	{_State101, FuncToken}:                              _GotoState21Action,
	{_State101, LparenToken}:                            _GotoState26Action,
	{_State101, LabelDeclToken}:                         _GotoState24Action,
	{_State101, NotToken}:                               _GotoState28Action,
	{_State101, SubToken}:                               _GotoState32Action,
	{_State101, MulToken}:                               _GotoState27Action,
	{_State101, BitNegToken}:                            _GotoState18Action,
	{_State101, BitAndToken}:                            _GotoState17Action,
	{_State101, LexErrorToken}:                          _GotoState25Action,
	{_State101, LiteralType}:                            _GotoState44Action,
	{_State101, AnonymousStructExprType}:                _GotoState38Action,
	{_State101, AtomExprType}:                           _GotoState39Action,
	{_State101, CallExprType}:                           _GotoState41Action,
	{_State101, AccessExprType}:                         _GotoState34Action,
	{_State101, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State101, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State101, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State101, MulExprType}:                            _GotoState45Action,
	{_State101, AddExprType}:                            _GotoState35Action,
	{_State101, CmpExprType}:                            _GotoState42Action,
	{_State101, AndExprType}:                            _GotoState163Action,
	{_State101, OptionalLabelDeclType}:                  _GotoState102Action,
	{_State101, BlockExprType}:                          _GotoState40Action,
	{_State101, ExplicitStructDefType}:                  _GotoState43Action,
	{_State101, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State102, LbraceToken}:                            _GotoState95Action,
	{_State102, BlockBodyType}:                          _GotoState97Action,
	{_State104, PackageStatementsType}:                  _GotoState164Action,
	{_State105, IdentifierToken}:                        _GotoState165Action,
	{_State105, GenericParameterDefType}:                _GotoState166Action,
	{_State105, GenericParameterDefsType}:               _GotoState167Action,
	{_State105, OptionalGenericParameterDefsType}:       _GotoState168Action,
	{_State106, IdentifierToken}:                        _GotoState146Action,
	{_State106, StructToken}:                            _GotoState31Action,
	{_State106, EnumToken}:                              _GotoState110Action,
	{_State106, TraitToken}:                             _GotoState15Action,
	{_State106, FuncToken}:                              _GotoState142Action,
	{_State106, LparenToken}:                            _GotoState113Action,
	{_State106, QuestionToken}:                          _GotoState114Action,
	{_State106, TildeTildeToken}:                        _GotoState115Action,
	{_State106, BitNegToken}:                            _GotoState109Action,
	{_State106, BitAndToken}:                            _GotoState108Action,
	{_State106, AtomTypeType}:                           _GotoState116Action,
	{_State106, TraitableTypeType}:                      _GotoState129Action,
	{_State106, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State106, ValueTypeType}:                          _GotoState169Action,
	{_State106, ImplicitStructDefType}:                  _GotoState122Action,
	{_State106, ExplicitStructDefType}:                  _GotoState118Action,
	{_State106, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State106, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State106, TraitDefType}:                           _GotoState126Action,
	{_State106, FuncTypeType}:                           _GotoState120Action,
	{_State107, IdentifierToken}:                        _GotoState146Action,
	{_State107, StructToken}:                            _GotoState31Action,
	{_State107, EnumToken}:                              _GotoState110Action,
	{_State107, TraitToken}:                             _GotoState15Action,
	{_State107, FuncToken}:                              _GotoState142Action,
	{_State107, LparenToken}:                            _GotoState113Action,
	{_State107, QuestionToken}:                          _GotoState114Action,
	{_State107, TildeTildeToken}:                        _GotoState115Action,
	{_State107, BitNegToken}:                            _GotoState109Action,
	{_State107, BitAndToken}:                            _GotoState108Action,
	{_State107, AtomTypeType}:                           _GotoState116Action,
	{_State107, TraitableTypeType}:                      _GotoState129Action,
	{_State107, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State107, ValueTypeType}:                          _GotoState170Action,
	{_State107, ImplicitStructDefType}:                  _GotoState122Action,
	{_State107, ExplicitStructDefType}:                  _GotoState118Action,
	{_State107, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State107, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State107, TraitDefType}:                           _GotoState126Action,
	{_State107, FuncTypeType}:                           _GotoState120Action,
	{_State108, IdentifierToken}:                        _GotoState146Action,
	{_State108, StructToken}:                            _GotoState31Action,
	{_State108, EnumToken}:                              _GotoState110Action,
	{_State108, TraitToken}:                             _GotoState15Action,
	{_State108, LparenToken}:                            _GotoState113Action,
	{_State108, TildeTildeToken}:                        _GotoState115Action,
	{_State108, BitNegToken}:                            _GotoState109Action,
	{_State108, AtomTypeType}:                           _GotoState116Action,
	{_State108, TraitableTypeType}:                      _GotoState129Action,
	{_State108, TraitAlgebraTypeType}:                   _GotoState171Action,
	{_State108, ImplicitStructDefType}:                  _GotoState122Action,
	{_State108, ExplicitStructDefType}:                  _GotoState118Action,
	{_State108, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State108, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State108, TraitDefType}:                           _GotoState126Action,
	{_State109, IdentifierToken}:                        _GotoState146Action,
	{_State109, StructToken}:                            _GotoState31Action,
	{_State109, EnumToken}:                              _GotoState110Action,
	{_State109, TraitToken}:                             _GotoState15Action,
	{_State109, LparenToken}:                            _GotoState113Action,
	{_State109, AtomTypeType}:                           _GotoState172Action,
	{_State109, ImplicitStructDefType}:                  _GotoState122Action,
	{_State109, ExplicitStructDefType}:                  _GotoState118Action,
	{_State109, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State109, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State109, TraitDefType}:                           _GotoState126Action,
	{_State110, LparenToken}:                            _GotoState173Action,
	{_State111, IdentifierToken}:                        _GotoState174Action,
	{_State111, LparenToken}:                            _GotoState175Action,
	{_State112, IdentifierToken}:                        _GotoState146Action,
	{_State112, StructToken}:                            _GotoState31Action,
	{_State112, EnumToken}:                              _GotoState110Action,
	{_State112, TraitToken}:                             _GotoState15Action,
	{_State112, FuncToken}:                              _GotoState142Action,
	{_State112, LparenToken}:                            _GotoState113Action,
	{_State112, QuestionToken}:                          _GotoState114Action,
	{_State112, DollarLbracketToken}:                    _GotoState67Action,
	{_State112, TildeTildeToken}:                        _GotoState115Action,
	{_State112, BitNegToken}:                            _GotoState109Action,
	{_State112, BitAndToken}:                            _GotoState108Action,
	{_State112, OptionalGenericBindingType}:             _GotoState176Action,
	{_State112, AtomTypeType}:                           _GotoState116Action,
	{_State112, TraitableTypeType}:                      _GotoState129Action,
	{_State112, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State112, ValueTypeType}:                          _GotoState177Action,
	{_State112, ImplicitStructDefType}:                  _GotoState122Action,
	{_State112, ExplicitStructDefType}:                  _GotoState118Action,
	{_State112, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State112, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State112, TraitDefType}:                           _GotoState126Action,
	{_State112, FuncTypeType}:                           _GotoState120Action,
	{_State113, IdentifierToken}:                        _GotoState112Action,
	{_State113, StructToken}:                            _GotoState31Action,
	{_State113, EnumToken}:                              _GotoState110Action,
	{_State113, TraitToken}:                             _GotoState15Action,
	{_State113, FuncToken}:                              _GotoState142Action,
	{_State113, LparenToken}:                            _GotoState113Action,
	{_State113, QuestionToken}:                          _GotoState114Action,
	{_State113, TildeTildeToken}:                        _GotoState115Action,
	{_State113, BitNegToken}:                            _GotoState109Action,
	{_State113, BitAndToken}:                            _GotoState108Action,
	{_State113, AtomTypeType}:                           _GotoState116Action,
	{_State113, TraitableTypeType}:                      _GotoState129Action,
	{_State113, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State113, ValueTypeType}:                          _GotoState130Action,
	{_State113, FieldDefType}:                           _GotoState179Action,
	{_State113, ImplicitFieldDefsType}:                  _GotoState181Action,
	{_State113, OptionalImplicitFieldDefsType}:          _GotoState182Action,
	{_State113, ImplicitStructDefType}:                  _GotoState122Action,
	{_State113, ExplicitStructDefType}:                  _GotoState118Action,
	{_State113, EnumValueDefType}:                       _GotoState178Action,
	{_State113, ImplicitEnumValueDefsType}:              _GotoState180Action,
	{_State113, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State113, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State113, TraitDefType}:                           _GotoState126Action,
	{_State113, FuncTypeType}:                           _GotoState120Action,
	{_State115, IdentifierToken}:                        _GotoState146Action,
	{_State115, StructToken}:                            _GotoState31Action,
	{_State115, EnumToken}:                              _GotoState110Action,
	{_State115, TraitToken}:                             _GotoState15Action,
	{_State115, LparenToken}:                            _GotoState113Action,
	{_State115, AtomTypeType}:                           _GotoState183Action,
	{_State115, ImplicitStructDefType}:                  _GotoState122Action,
	{_State115, ExplicitStructDefType}:                  _GotoState118Action,
	{_State115, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State115, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State115, TraitDefType}:                           _GotoState126Action,
	{_State124, RparenToken}:                            _GotoState184Action,
	{_State125, AddToken}:                               _GotoState185Action,
	{_State125, SubToken}:                               _GotoState187Action,
	{_State125, MulToken}:                               _GotoState186Action,
	{_State127, NewlinesToken}:                          _GotoState189Action,
	{_State127, CommaToken}:                             _GotoState188Action,
	{_State131, IdentifierToken}:                        _GotoState146Action,
	{_State131, StructToken}:                            _GotoState31Action,
	{_State131, EnumToken}:                              _GotoState110Action,
	{_State131, TraitToken}:                             _GotoState15Action,
	{_State131, FuncToken}:                              _GotoState142Action,
	{_State131, LparenToken}:                            _GotoState113Action,
	{_State131, QuestionToken}:                          _GotoState114Action,
	{_State131, DotdotdotToken}:                         _GotoState190Action,
	{_State131, TildeTildeToken}:                        _GotoState115Action,
	{_State131, BitNegToken}:                            _GotoState109Action,
	{_State131, BitAndToken}:                            _GotoState108Action,
	{_State131, AtomTypeType}:                           _GotoState116Action,
	{_State131, TraitableTypeType}:                      _GotoState129Action,
	{_State131, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State131, ValueTypeType}:                          _GotoState191Action,
	{_State131, ImplicitStructDefType}:                  _GotoState122Action,
	{_State131, ExplicitStructDefType}:                  _GotoState118Action,
	{_State131, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State131, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State131, TraitDefType}:                           _GotoState126Action,
	{_State131, FuncTypeType}:                           _GotoState120Action,
	{_State132, RparenToken}:                            _GotoState192Action,
	{_State133, DollarLbracketToken}:                    _GotoState105Action,
	{_State133, OptionalGenericParametersType}:          _GotoState193Action,
	{_State134, RparenToken}:                            _GotoState194Action,
	{_State136, CommaToken}:                             _GotoState195Action,
	{_State137, IntegerLiteralToken}:                    _GotoState23Action,
	{_State137, FloatLiteralToken}:                      _GotoState20Action,
	{_State137, RuneLiteralToken}:                       _GotoState29Action,
	{_State137, StringLiteralToken}:                     _GotoState30Action,
	{_State137, IdentifierToken}:                        _GotoState22Action,
	{_State137, TrueToken}:                              _GotoState33Action,
	{_State137, FalseToken}:                             _GotoState19Action,
	{_State137, StructToken}:                            _GotoState31Action,
	{_State137, FuncToken}:                              _GotoState21Action,
	{_State137, LparenToken}:                            _GotoState26Action,
	{_State137, LabelDeclToken}:                         _GotoState24Action,
	{_State137, NotToken}:                               _GotoState28Action,
	{_State137, SubToken}:                               _GotoState32Action,
	{_State137, MulToken}:                               _GotoState27Action,
	{_State137, BitNegToken}:                            _GotoState18Action,
	{_State137, BitAndToken}:                            _GotoState17Action,
	{_State137, LexErrorToken}:                          _GotoState25Action,
	{_State137, LiteralType}:                            _GotoState44Action,
	{_State137, AnonymousStructExprType}:                _GotoState38Action,
	{_State137, AtomExprType}:                           _GotoState39Action,
	{_State137, CallExprType}:                           _GotoState41Action,
	{_State137, AccessExprType}:                         _GotoState34Action,
	{_State137, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State137, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State137, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State137, MulExprType}:                            _GotoState45Action,
	{_State137, AddExprType}:                            _GotoState35Action,
	{_State137, CmpExprType}:                            _GotoState42Action,
	{_State137, AndExprType}:                            _GotoState36Action,
	{_State137, OrExprType}:                             _GotoState47Action,
	{_State137, SequenceExprType}:                       _GotoState51Action,
	{_State137, OptionalLabelDeclType}:                  _GotoState46Action,
	{_State137, BlockExprType}:                          _GotoState40Action,
	{_State137, ExpressionType}:                         _GotoState196Action,
	{_State137, ExplicitStructDefType}:                  _GotoState43Action,
	{_State137, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State138, IntegerLiteralToken}:                    _GotoState23Action,
	{_State138, FloatLiteralToken}:                      _GotoState20Action,
	{_State138, RuneLiteralToken}:                       _GotoState29Action,
	{_State138, StringLiteralToken}:                     _GotoState30Action,
	{_State138, IdentifierToken}:                        _GotoState60Action,
	{_State138, TrueToken}:                              _GotoState33Action,
	{_State138, FalseToken}:                             _GotoState19Action,
	{_State138, StructToken}:                            _GotoState31Action,
	{_State138, FuncToken}:                              _GotoState21Action,
	{_State138, LparenToken}:                            _GotoState26Action,
	{_State138, LabelDeclToken}:                         _GotoState24Action,
	{_State138, NotToken}:                               _GotoState28Action,
	{_State138, SubToken}:                               _GotoState32Action,
	{_State138, MulToken}:                               _GotoState27Action,
	{_State138, BitNegToken}:                            _GotoState18Action,
	{_State138, BitAndToken}:                            _GotoState17Action,
	{_State138, LexErrorToken}:                          _GotoState25Action,
	{_State138, LiteralType}:                            _GotoState44Action,
	{_State138, AnonymousStructExprType}:                _GotoState38Action,
	{_State138, AtomExprType}:                           _GotoState39Action,
	{_State138, OptionalExpressionType}:                 _GotoState65Action,
	{_State138, ColonExpressionsType}:                   _GotoState63Action,
	{_State138, ArgumentType}:                           _GotoState197Action,
	{_State138, CallExprType}:                           _GotoState41Action,
	{_State138, AccessExprType}:                         _GotoState34Action,
	{_State138, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State138, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State138, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State138, MulExprType}:                            _GotoState45Action,
	{_State138, AddExprType}:                            _GotoState35Action,
	{_State138, CmpExprType}:                            _GotoState42Action,
	{_State138, AndExprType}:                            _GotoState36Action,
	{_State138, OrExprType}:                             _GotoState47Action,
	{_State138, SequenceExprType}:                       _GotoState51Action,
	{_State138, OptionalLabelDeclType}:                  _GotoState46Action,
	{_State138, BlockExprType}:                          _GotoState40Action,
	{_State138, ExpressionType}:                         _GotoState64Action,
	{_State138, ExplicitStructDefType}:                  _GotoState43Action,
	{_State138, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State140, IntegerLiteralToken}:                    _GotoState23Action,
	{_State140, FloatLiteralToken}:                      _GotoState20Action,
	{_State140, RuneLiteralToken}:                       _GotoState29Action,
	{_State140, StringLiteralToken}:                     _GotoState30Action,
	{_State140, IdentifierToken}:                        _GotoState22Action,
	{_State140, TrueToken}:                              _GotoState33Action,
	{_State140, FalseToken}:                             _GotoState19Action,
	{_State140, StructToken}:                            _GotoState31Action,
	{_State140, FuncToken}:                              _GotoState21Action,
	{_State140, LparenToken}:                            _GotoState26Action,
	{_State140, LabelDeclToken}:                         _GotoState24Action,
	{_State140, NotToken}:                               _GotoState28Action,
	{_State140, SubToken}:                               _GotoState32Action,
	{_State140, MulToken}:                               _GotoState27Action,
	{_State140, BitNegToken}:                            _GotoState18Action,
	{_State140, BitAndToken}:                            _GotoState17Action,
	{_State140, LexErrorToken}:                          _GotoState25Action,
	{_State140, LiteralType}:                            _GotoState44Action,
	{_State140, AnonymousStructExprType}:                _GotoState38Action,
	{_State140, AtomExprType}:                           _GotoState39Action,
	{_State140, OptionalExpressionType}:                 _GotoState199Action,
	{_State140, CallExprType}:                           _GotoState41Action,
	{_State140, AccessExprType}:                         _GotoState34Action,
	{_State140, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State140, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State140, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State140, MulExprType}:                            _GotoState45Action,
	{_State140, AddExprType}:                            _GotoState35Action,
	{_State140, CmpExprType}:                            _GotoState42Action,
	{_State140, AndExprType}:                            _GotoState36Action,
	{_State140, OrExprType}:                             _GotoState47Action,
	{_State140, SequenceExprType}:                       _GotoState51Action,
	{_State140, OptionalLabelDeclType}:                  _GotoState46Action,
	{_State140, BlockExprType}:                          _GotoState40Action,
	{_State140, ExpressionType}:                         _GotoState198Action,
	{_State140, ExplicitStructDefType}:                  _GotoState43Action,
	{_State140, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State141, IntegerLiteralToken}:                    _GotoState23Action,
	{_State141, FloatLiteralToken}:                      _GotoState20Action,
	{_State141, RuneLiteralToken}:                       _GotoState29Action,
	{_State141, StringLiteralToken}:                     _GotoState30Action,
	{_State141, IdentifierToken}:                        _GotoState22Action,
	{_State141, TrueToken}:                              _GotoState33Action,
	{_State141, FalseToken}:                             _GotoState19Action,
	{_State141, StructToken}:                            _GotoState31Action,
	{_State141, FuncToken}:                              _GotoState21Action,
	{_State141, LparenToken}:                            _GotoState26Action,
	{_State141, LabelDeclToken}:                         _GotoState24Action,
	{_State141, NotToken}:                               _GotoState28Action,
	{_State141, SubToken}:                               _GotoState32Action,
	{_State141, MulToken}:                               _GotoState27Action,
	{_State141, BitNegToken}:                            _GotoState18Action,
	{_State141, BitAndToken}:                            _GotoState17Action,
	{_State141, LexErrorToken}:                          _GotoState25Action,
	{_State141, LiteralType}:                            _GotoState44Action,
	{_State141, AnonymousStructExprType}:                _GotoState38Action,
	{_State141, AtomExprType}:                           _GotoState39Action,
	{_State141, OptionalExpressionType}:                 _GotoState200Action,
	{_State141, CallExprType}:                           _GotoState41Action,
	{_State141, AccessExprType}:                         _GotoState34Action,
	{_State141, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State141, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State141, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State141, MulExprType}:                            _GotoState45Action,
	{_State141, AddExprType}:                            _GotoState35Action,
	{_State141, CmpExprType}:                            _GotoState42Action,
	{_State141, AndExprType}:                            _GotoState36Action,
	{_State141, OrExprType}:                             _GotoState47Action,
	{_State141, SequenceExprType}:                       _GotoState51Action,
	{_State141, OptionalLabelDeclType}:                  _GotoState46Action,
	{_State141, BlockExprType}:                          _GotoState40Action,
	{_State141, ExpressionType}:                         _GotoState198Action,
	{_State141, ExplicitStructDefType}:                  _GotoState43Action,
	{_State141, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State142, LparenToken}:                            _GotoState175Action,
	{_State143, NewlinesToken}:                          _GotoState202Action,
	{_State143, CommaToken}:                             _GotoState201Action,
	{_State145, RparenToken}:                            _GotoState203Action,
	{_State146, DollarLbracketToken}:                    _GotoState67Action,
	{_State146, OptionalGenericBindingType}:             _GotoState176Action,
	{_State147, CommaToken}:                             _GotoState204Action,
	{_State148, RbracketToken}:                          _GotoState205Action,
	{_State151, RbracketToken}:                          _GotoState206Action,
	{_State152, IntegerLiteralToken}:                    _GotoState23Action,
	{_State152, FloatLiteralToken}:                      _GotoState20Action,
	{_State152, RuneLiteralToken}:                       _GotoState29Action,
	{_State152, StringLiteralToken}:                     _GotoState30Action,
	{_State152, IdentifierToken}:                        _GotoState60Action,
	{_State152, TrueToken}:                              _GotoState33Action,
	{_State152, FalseToken}:                             _GotoState19Action,
	{_State152, StructToken}:                            _GotoState31Action,
	{_State152, FuncToken}:                              _GotoState21Action,
	{_State152, LparenToken}:                            _GotoState26Action,
	{_State152, LabelDeclToken}:                         _GotoState24Action,
	{_State152, NotToken}:                               _GotoState28Action,
	{_State152, SubToken}:                               _GotoState32Action,
	{_State152, MulToken}:                               _GotoState27Action,
	{_State152, BitNegToken}:                            _GotoState18Action,
	{_State152, BitAndToken}:                            _GotoState17Action,
	{_State152, LexErrorToken}:                          _GotoState25Action,
	{_State152, LiteralType}:                            _GotoState44Action,
	{_State152, AnonymousStructExprType}:                _GotoState38Action,
	{_State152, AtomExprType}:                           _GotoState39Action,
	{_State152, OptionalExpressionType}:                 _GotoState65Action,
	{_State152, ColonExpressionsType}:                   _GotoState63Action,
	{_State152, ArgumentType}:                           _GotoState61Action,
	{_State152, ArgumentsType}:                          _GotoState207Action,
	{_State152, OptionalArgumentsType}:                  _GotoState208Action,
	{_State152, CallExprType}:                           _GotoState41Action,
	{_State152, AccessExprType}:                         _GotoState34Action,
	{_State152, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State152, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State152, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State152, MulExprType}:                            _GotoState45Action,
	{_State152, AddExprType}:                            _GotoState35Action,
	{_State152, CmpExprType}:                            _GotoState42Action,
	{_State152, AndExprType}:                            _GotoState36Action,
	{_State152, OrExprType}:                             _GotoState47Action,
	{_State152, SequenceExprType}:                       _GotoState51Action,
	{_State152, OptionalLabelDeclType}:                  _GotoState46Action,
	{_State152, BlockExprType}:                          _GotoState40Action,
	{_State152, ExpressionType}:                         _GotoState64Action,
	{_State152, ExplicitStructDefType}:                  _GotoState43Action,
	{_State152, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State153, MulToken}:                               _GotoState91Action,
	{_State153, DivToken}:                               _GotoState89Action,
	{_State153, ModToken}:                               _GotoState90Action,
	{_State153, BitAndToken}:                            _GotoState86Action,
	{_State153, BitLshiftToken}:                         _GotoState87Action,
	{_State153, BitRshiftToken}:                         _GotoState88Action,
	{_State153, MulOpType}:                              _GotoState92Action,
	{_State154, EqualToken}:                             _GotoState78Action,
	{_State154, NotEqualToken}:                          _GotoState83Action,
	{_State154, LessToken}:                              _GotoState81Action,
	{_State154, LessOrEqualToken}:                       _GotoState82Action,
	{_State154, GreaterToken}:                           _GotoState79Action,
	{_State154, GreaterOrEqualToken}:                    _GotoState80Action,
	{_State154, CmpOpType}:                              _GotoState84Action,
	{_State155, AddToken}:                               _GotoState72Action,
	{_State155, SubToken}:                               _GotoState75Action,
	{_State155, BitXorToken}:                            _GotoState74Action,
	{_State155, BitOrToken}:                             _GotoState73Action,
	{_State155, AddOpType}:                              _GotoState76Action,
	{_State156, RparenToken}:                            _GotoState209Action,
	{_State156, CommaToken}:                             _GotoState138Action,
	{_State159, LabelDeclToken}:                         _GotoState24Action,
	{_State159, OptionalLabelDeclType}:                  _GotoState102Action,
	{_State159, BlockExprType}:                          _GotoState210Action,
	{_State160, LbraceToken}:                            _GotoState95Action,
	{_State160, BlockBodyType}:                          _GotoState211Action,
	{_State161, IntegerLiteralToken}:                    _GotoState23Action,
	{_State161, FloatLiteralToken}:                      _GotoState20Action,
	{_State161, RuneLiteralToken}:                       _GotoState29Action,
	{_State161, StringLiteralToken}:                     _GotoState30Action,
	{_State161, IdentifierToken}:                        _GotoState22Action,
	{_State161, TrueToken}:                              _GotoState33Action,
	{_State161, FalseToken}:                             _GotoState19Action,
	{_State161, ReturnToken}:                            _GotoState216Action,
	{_State161, BreakToken}:                             _GotoState213Action,
	{_State161, ContinueToken}:                          _GotoState214Action,
	{_State161, UnsafeToken}:                            _GotoState217Action,
	{_State161, StructToken}:                            _GotoState31Action,
	{_State161, FuncToken}:                              _GotoState21Action,
	{_State161, AsyncToken}:                             _GotoState212Action,
	{_State161, RbraceToken}:                            _GotoState215Action,
	{_State161, LparenToken}:                            _GotoState26Action,
	{_State161, LabelDeclToken}:                         _GotoState24Action,
	{_State161, NotToken}:                               _GotoState28Action,
	{_State161, SubToken}:                               _GotoState32Action,
	{_State161, MulToken}:                               _GotoState27Action,
	{_State161, BitNegToken}:                            _GotoState18Action,
	{_State161, BitAndToken}:                            _GotoState17Action,
	{_State161, LexErrorToken}:                          _GotoState25Action,
	{_State161, LiteralType}:                            _GotoState44Action,
	{_State161, AnonymousStructExprType}:                _GotoState38Action,
	{_State161, AtomExprType}:                           _GotoState39Action,
	{_State161, CallExprType}:                           _GotoState41Action,
	{_State161, AccessExprType}:                         _GotoState218Action,
	{_State161, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State161, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State161, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State161, MulExprType}:                            _GotoState45Action,
	{_State161, AddExprType}:                            _GotoState35Action,
	{_State161, CmpExprType}:                            _GotoState42Action,
	{_State161, AndExprType}:                            _GotoState36Action,
	{_State161, OrExprType}:                             _GotoState47Action,
	{_State161, SequenceExprType}:                       _GotoState51Action,
	{_State161, JumpTypeType}:                           _GotoState221Action,
	{_State161, ExpressionOrImplicitStructType}:         _GotoState220Action,
	{_State161, UnsafeStatementType}:                    _GotoState224Action,
	{_State161, StatementBodyType}:                      _GotoState223Action,
	{_State161, StatementType}:                          _GotoState222Action,
	{_State161, OptionalLabelDeclType}:                  _GotoState46Action,
	{_State161, BlockExprType}:                          _GotoState40Action,
	{_State161, ExpressionType}:                         _GotoState219Action,
	{_State161, ExplicitStructDefType}:                  _GotoState43Action,
	{_State161, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State162, LbraceToken}:                            _GotoState225Action,
	{_State163, AndToken}:                               _GotoState77Action,
	{_State164, UnsafeToken}:                            _GotoState217Action,
	{_State164, RparenToken}:                            _GotoState226Action,
	{_State164, UnsafeStatementType}:                    _GotoState229Action,
	{_State164, PackageStatementBodyType}:               _GotoState228Action,
	{_State164, PackageStatementType}:                   _GotoState227Action,
	{_State165, IdentifierToken}:                        _GotoState146Action,
	{_State165, StructToken}:                            _GotoState31Action,
	{_State165, EnumToken}:                              _GotoState110Action,
	{_State165, TraitToken}:                             _GotoState15Action,
	{_State165, FuncToken}:                              _GotoState142Action,
	{_State165, LparenToken}:                            _GotoState113Action,
	{_State165, QuestionToken}:                          _GotoState114Action,
	{_State165, TildeTildeToken}:                        _GotoState115Action,
	{_State165, BitNegToken}:                            _GotoState109Action,
	{_State165, BitAndToken}:                            _GotoState108Action,
	{_State165, AtomTypeType}:                           _GotoState116Action,
	{_State165, TraitableTypeType}:                      _GotoState129Action,
	{_State165, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State165, ValueTypeType}:                          _GotoState230Action,
	{_State165, ImplicitStructDefType}:                  _GotoState122Action,
	{_State165, ExplicitStructDefType}:                  _GotoState118Action,
	{_State165, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State165, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State165, TraitDefType}:                           _GotoState126Action,
	{_State165, FuncTypeType}:                           _GotoState120Action,
	{_State167, CommaToken}:                             _GotoState231Action,
	{_State168, RbracketToken}:                          _GotoState232Action,
	{_State170, ImplementsToken}:                        _GotoState233Action,
	{_State171, AddToken}:                               _GotoState185Action,
	{_State171, SubToken}:                               _GotoState187Action,
	{_State171, MulToken}:                               _GotoState186Action,
	{_State173, IdentifierToken}:                        _GotoState112Action,
	{_State173, StructToken}:                            _GotoState31Action,
	{_State173, EnumToken}:                              _GotoState110Action,
	{_State173, TraitToken}:                             _GotoState15Action,
	{_State173, FuncToken}:                              _GotoState142Action,
	{_State173, LparenToken}:                            _GotoState113Action,
	{_State173, QuestionToken}:                          _GotoState114Action,
	{_State173, TildeTildeToken}:                        _GotoState115Action,
	{_State173, BitNegToken}:                            _GotoState109Action,
	{_State173, BitAndToken}:                            _GotoState108Action,
	{_State173, AtomTypeType}:                           _GotoState116Action,
	{_State173, TraitableTypeType}:                      _GotoState129Action,
	{_State173, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State173, ValueTypeType}:                          _GotoState130Action,
	{_State173, FieldDefType}:                           _GotoState236Action,
	{_State173, ImplicitStructDefType}:                  _GotoState122Action,
	{_State173, ExplicitStructDefType}:                  _GotoState118Action,
	{_State173, EnumValueDefType}:                       _GotoState234Action,
	{_State173, ImplicitEnumValueDefsType}:              _GotoState237Action,
	{_State173, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State173, ExplicitEnumValueDefsType}:              _GotoState235Action,
	{_State173, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State173, TraitDefType}:                           _GotoState126Action,
	{_State173, FuncTypeType}:                           _GotoState120Action,
	{_State174, LparenToken}:                            _GotoState238Action,
	{_State175, IdentifierToken}:                        _GotoState240Action,
	{_State175, StructToken}:                            _GotoState31Action,
	{_State175, EnumToken}:                              _GotoState110Action,
	{_State175, TraitToken}:                             _GotoState15Action,
	{_State175, FuncToken}:                              _GotoState142Action,
	{_State175, LparenToken}:                            _GotoState113Action,
	{_State175, QuestionToken}:                          _GotoState114Action,
	{_State175, DotdotdotToken}:                         _GotoState239Action,
	{_State175, TildeTildeToken}:                        _GotoState115Action,
	{_State175, BitNegToken}:                            _GotoState109Action,
	{_State175, BitAndToken}:                            _GotoState108Action,
	{_State175, AtomTypeType}:                           _GotoState116Action,
	{_State175, TraitableTypeType}:                      _GotoState129Action,
	{_State175, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State175, ValueTypeType}:                          _GotoState244Action,
	{_State175, ImplicitStructDefType}:                  _GotoState122Action,
	{_State175, ExplicitStructDefType}:                  _GotoState118Action,
	{_State175, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State175, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State175, TraitDefType}:                           _GotoState126Action,
	{_State175, ParameterDeclType}:                      _GotoState242Action,
	{_State175, ParameterDeclsType}:                     _GotoState243Action,
	{_State175, OptionalParameterDeclsType}:             _GotoState241Action,
	{_State175, FuncTypeType}:                           _GotoState120Action,
	{_State178, OrToken}:                                _GotoState245Action,
	{_State179, AssignToken}:                            _GotoState246Action,
	{_State180, RparenToken}:                            _GotoState248Action,
	{_State180, OrToken}:                                _GotoState247Action,
	{_State181, CommaToken}:                             _GotoState249Action,
	{_State182, RparenToken}:                            _GotoState250Action,
	{_State185, IdentifierToken}:                        _GotoState146Action,
	{_State185, StructToken}:                            _GotoState31Action,
	{_State185, EnumToken}:                              _GotoState110Action,
	{_State185, TraitToken}:                             _GotoState15Action,
	{_State185, LparenToken}:                            _GotoState113Action,
	{_State185, TildeTildeToken}:                        _GotoState115Action,
	{_State185, BitNegToken}:                            _GotoState109Action,
	{_State185, AtomTypeType}:                           _GotoState116Action,
	{_State185, TraitableTypeType}:                      _GotoState251Action,
	{_State185, ImplicitStructDefType}:                  _GotoState122Action,
	{_State185, ExplicitStructDefType}:                  _GotoState118Action,
	{_State185, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State185, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State185, TraitDefType}:                           _GotoState126Action,
	{_State186, IdentifierToken}:                        _GotoState146Action,
	{_State186, StructToken}:                            _GotoState31Action,
	{_State186, EnumToken}:                              _GotoState110Action,
	{_State186, TraitToken}:                             _GotoState15Action,
	{_State186, LparenToken}:                            _GotoState113Action,
	{_State186, TildeTildeToken}:                        _GotoState115Action,
	{_State186, BitNegToken}:                            _GotoState109Action,
	{_State186, AtomTypeType}:                           _GotoState116Action,
	{_State186, TraitableTypeType}:                      _GotoState252Action,
	{_State186, ImplicitStructDefType}:                  _GotoState122Action,
	{_State186, ExplicitStructDefType}:                  _GotoState118Action,
	{_State186, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State186, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State186, TraitDefType}:                           _GotoState126Action,
	{_State187, IdentifierToken}:                        _GotoState146Action,
	{_State187, StructToken}:                            _GotoState31Action,
	{_State187, EnumToken}:                              _GotoState110Action,
	{_State187, TraitToken}:                             _GotoState15Action,
	{_State187, LparenToken}:                            _GotoState113Action,
	{_State187, TildeTildeToken}:                        _GotoState115Action,
	{_State187, BitNegToken}:                            _GotoState109Action,
	{_State187, AtomTypeType}:                           _GotoState116Action,
	{_State187, TraitableTypeType}:                      _GotoState253Action,
	{_State187, ImplicitStructDefType}:                  _GotoState122Action,
	{_State187, ExplicitStructDefType}:                  _GotoState118Action,
	{_State187, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State187, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State187, TraitDefType}:                           _GotoState126Action,
	{_State188, IdentifierToken}:                        _GotoState112Action,
	{_State188, StructToken}:                            _GotoState31Action,
	{_State188, EnumToken}:                              _GotoState110Action,
	{_State188, TraitToken}:                             _GotoState15Action,
	{_State188, FuncToken}:                              _GotoState111Action,
	{_State188, LparenToken}:                            _GotoState113Action,
	{_State188, QuestionToken}:                          _GotoState114Action,
	{_State188, TildeTildeToken}:                        _GotoState115Action,
	{_State188, BitNegToken}:                            _GotoState109Action,
	{_State188, BitAndToken}:                            _GotoState108Action,
	{_State188, AtomTypeType}:                           _GotoState116Action,
	{_State188, TraitableTypeType}:                      _GotoState129Action,
	{_State188, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State188, ValueTypeType}:                          _GotoState130Action,
	{_State188, FieldDefType}:                           _GotoState119Action,
	{_State188, ImplicitStructDefType}:                  _GotoState122Action,
	{_State188, ExplicitStructDefType}:                  _GotoState118Action,
	{_State188, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State188, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State188, TraitPropertyType}:                      _GotoState254Action,
	{_State188, TraitDefType}:                           _GotoState126Action,
	{_State188, FuncTypeType}:                           _GotoState120Action,
	{_State188, MethodSignatureType}:                    _GotoState123Action,
	{_State189, IdentifierToken}:                        _GotoState112Action,
	{_State189, StructToken}:                            _GotoState31Action,
	{_State189, EnumToken}:                              _GotoState110Action,
	{_State189, TraitToken}:                             _GotoState15Action,
	{_State189, FuncToken}:                              _GotoState111Action,
	{_State189, LparenToken}:                            _GotoState113Action,
	{_State189, QuestionToken}:                          _GotoState114Action,
	{_State189, TildeTildeToken}:                        _GotoState115Action,
	{_State189, BitNegToken}:                            _GotoState109Action,
	{_State189, BitAndToken}:                            _GotoState108Action,
	{_State189, AtomTypeType}:                           _GotoState116Action,
	{_State189, TraitableTypeType}:                      _GotoState129Action,
	{_State189, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State189, ValueTypeType}:                          _GotoState130Action,
	{_State189, FieldDefType}:                           _GotoState119Action,
	{_State189, ImplicitStructDefType}:                  _GotoState122Action,
	{_State189, ExplicitStructDefType}:                  _GotoState118Action,
	{_State189, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State189, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State189, TraitPropertyType}:                      _GotoState255Action,
	{_State189, TraitDefType}:                           _GotoState126Action,
	{_State189, FuncTypeType}:                           _GotoState120Action,
	{_State189, MethodSignatureType}:                    _GotoState123Action,
	{_State190, IdentifierToken}:                        _GotoState146Action,
	{_State190, StructToken}:                            _GotoState31Action,
	{_State190, EnumToken}:                              _GotoState110Action,
	{_State190, TraitToken}:                             _GotoState15Action,
	{_State190, FuncToken}:                              _GotoState142Action,
	{_State190, LparenToken}:                            _GotoState113Action,
	{_State190, QuestionToken}:                          _GotoState114Action,
	{_State190, TildeTildeToken}:                        _GotoState115Action,
	{_State190, BitNegToken}:                            _GotoState109Action,
	{_State190, BitAndToken}:                            _GotoState108Action,
	{_State190, AtomTypeType}:                           _GotoState116Action,
	{_State190, TraitableTypeType}:                      _GotoState129Action,
	{_State190, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State190, ValueTypeType}:                          _GotoState256Action,
	{_State190, ImplicitStructDefType}:                  _GotoState122Action,
	{_State190, ExplicitStructDefType}:                  _GotoState118Action,
	{_State190, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State190, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State190, TraitDefType}:                           _GotoState126Action,
	{_State190, FuncTypeType}:                           _GotoState120Action,
	{_State193, LparenToken}:                            _GotoState257Action,
	{_State194, IdentifierToken}:                        _GotoState146Action,
	{_State194, StructToken}:                            _GotoState31Action,
	{_State194, EnumToken}:                              _GotoState110Action,
	{_State194, TraitToken}:                             _GotoState15Action,
	{_State194, FuncToken}:                              _GotoState142Action,
	{_State194, LparenToken}:                            _GotoState113Action,
	{_State194, QuestionToken}:                          _GotoState114Action,
	{_State194, TildeTildeToken}:                        _GotoState115Action,
	{_State194, BitNegToken}:                            _GotoState109Action,
	{_State194, BitAndToken}:                            _GotoState108Action,
	{_State194, AtomTypeType}:                           _GotoState116Action,
	{_State194, TraitableTypeType}:                      _GotoState129Action,
	{_State194, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State194, ValueTypeType}:                          _GotoState259Action,
	{_State194, ImplicitStructDefType}:                  _GotoState122Action,
	{_State194, ExplicitStructDefType}:                  _GotoState118Action,
	{_State194, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State194, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State194, TraitDefType}:                           _GotoState126Action,
	{_State194, ReturnTypeType}:                         _GotoState258Action,
	{_State194, FuncTypeType}:                           _GotoState120Action,
	{_State195, IdentifierToken}:                        _GotoState131Action,
	{_State195, ParameterDefType}:                       _GotoState260Action,
	{_State201, IdentifierToken}:                        _GotoState112Action,
	{_State201, StructToken}:                            _GotoState31Action,
	{_State201, EnumToken}:                              _GotoState110Action,
	{_State201, TraitToken}:                             _GotoState15Action,
	{_State201, FuncToken}:                              _GotoState142Action,
	{_State201, LparenToken}:                            _GotoState113Action,
	{_State201, QuestionToken}:                          _GotoState114Action,
	{_State201, TildeTildeToken}:                        _GotoState115Action,
	{_State201, BitNegToken}:                            _GotoState109Action,
	{_State201, BitAndToken}:                            _GotoState108Action,
	{_State201, AtomTypeType}:                           _GotoState116Action,
	{_State201, TraitableTypeType}:                      _GotoState129Action,
	{_State201, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State201, ValueTypeType}:                          _GotoState130Action,
	{_State201, FieldDefType}:                           _GotoState261Action,
	{_State201, ImplicitStructDefType}:                  _GotoState122Action,
	{_State201, ExplicitStructDefType}:                  _GotoState118Action,
	{_State201, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State201, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State201, TraitDefType}:                           _GotoState126Action,
	{_State201, FuncTypeType}:                           _GotoState120Action,
	{_State202, IdentifierToken}:                        _GotoState112Action,
	{_State202, StructToken}:                            _GotoState31Action,
	{_State202, EnumToken}:                              _GotoState110Action,
	{_State202, TraitToken}:                             _GotoState15Action,
	{_State202, FuncToken}:                              _GotoState142Action,
	{_State202, LparenToken}:                            _GotoState113Action,
	{_State202, QuestionToken}:                          _GotoState114Action,
	{_State202, TildeTildeToken}:                        _GotoState115Action,
	{_State202, BitNegToken}:                            _GotoState109Action,
	{_State202, BitAndToken}:                            _GotoState108Action,
	{_State202, AtomTypeType}:                           _GotoState116Action,
	{_State202, TraitableTypeType}:                      _GotoState129Action,
	{_State202, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State202, ValueTypeType}:                          _GotoState130Action,
	{_State202, FieldDefType}:                           _GotoState262Action,
	{_State202, ImplicitStructDefType}:                  _GotoState122Action,
	{_State202, ExplicitStructDefType}:                  _GotoState118Action,
	{_State202, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State202, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State202, TraitDefType}:                           _GotoState126Action,
	{_State202, FuncTypeType}:                           _GotoState120Action,
	{_State204, IdentifierToken}:                        _GotoState146Action,
	{_State204, StructToken}:                            _GotoState31Action,
	{_State204, EnumToken}:                              _GotoState110Action,
	{_State204, TraitToken}:                             _GotoState15Action,
	{_State204, FuncToken}:                              _GotoState142Action,
	{_State204, LparenToken}:                            _GotoState113Action,
	{_State204, QuestionToken}:                          _GotoState114Action,
	{_State204, TildeTildeToken}:                        _GotoState115Action,
	{_State204, BitNegToken}:                            _GotoState109Action,
	{_State204, BitAndToken}:                            _GotoState108Action,
	{_State204, AtomTypeType}:                           _GotoState116Action,
	{_State204, TraitableTypeType}:                      _GotoState129Action,
	{_State204, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State204, ValueTypeType}:                          _GotoState263Action,
	{_State204, ImplicitStructDefType}:                  _GotoState122Action,
	{_State204, ExplicitStructDefType}:                  _GotoState118Action,
	{_State204, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State204, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State204, TraitDefType}:                           _GotoState126Action,
	{_State204, FuncTypeType}:                           _GotoState120Action,
	{_State207, CommaToken}:                             _GotoState138Action,
	{_State208, RparenToken}:                            _GotoState264Action,
	{_State211, ElseToken}:                              _GotoState265Action,
	{_State212, IntegerLiteralToken}:                    _GotoState23Action,
	{_State212, FloatLiteralToken}:                      _GotoState20Action,
	{_State212, RuneLiteralToken}:                       _GotoState29Action,
	{_State212, StringLiteralToken}:                     _GotoState30Action,
	{_State212, IdentifierToken}:                        _GotoState22Action,
	{_State212, TrueToken}:                              _GotoState33Action,
	{_State212, FalseToken}:                             _GotoState19Action,
	{_State212, StructToken}:                            _GotoState31Action,
	{_State212, FuncToken}:                              _GotoState21Action,
	{_State212, LparenToken}:                            _GotoState26Action,
	{_State212, LabelDeclToken}:                         _GotoState24Action,
	{_State212, LexErrorToken}:                          _GotoState25Action,
	{_State212, LiteralType}:                            _GotoState44Action,
	{_State212, AnonymousStructExprType}:                _GotoState38Action,
	{_State212, AtomExprType}:                           _GotoState39Action,
	{_State212, CallExprType}:                           _GotoState267Action,
	{_State212, AccessExprType}:                         _GotoState266Action,
	{_State212, OptionalLabelDeclType}:                  _GotoState102Action,
	{_State212, BlockExprType}:                          _GotoState40Action,
	{_State212, ExplicitStructDefType}:                  _GotoState43Action,
	{_State212, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State217, LessToken}:                              _GotoState268Action,
	{_State218, LbracketToken}:                          _GotoState69Action,
	{_State218, DotToken}:                               _GotoState68Action,
	{_State218, QuestionToken}:                          _GotoState70Action,
	{_State218, DollarLbracketToken}:                    _GotoState67Action,
	{_State218, AddAssignToken}:                         _GotoState269Action,
	{_State218, SubAssignToken}:                         _GotoState280Action,
	{_State218, MulAssignToken}:                         _GotoState279Action,
	{_State218, DivAssignToken}:                         _GotoState277Action,
	{_State218, ModAssignToken}:                         _GotoState278Action,
	{_State218, AddOneAssignToken}:                      _GotoState270Action,
	{_State218, SubOneAssignToken}:                      _GotoState281Action,
	{_State218, BitNegAssignToken}:                      _GotoState273Action,
	{_State218, BitAndAssignToken}:                      _GotoState271Action,
	{_State218, BitOrAssignToken}:                       _GotoState274Action,
	{_State218, BitXorAssignToken}:                      _GotoState276Action,
	{_State218, BitLshiftAssignToken}:                   _GotoState272Action,
	{_State218, BitRshiftAssignToken}:                   _GotoState275Action,
	{_State218, OptionalGenericBindingType}:             _GotoState71Action,
	{_State218, OpOneAssignType}:                        _GotoState283Action,
	{_State218, BinaryOpAssignType}:                     _GotoState282Action,
	{_State220, CommaToken}:                             _GotoState284Action,
	{_State221, JumpLabelToken}:                         _GotoState285Action,
	{_State221, OptionalJumpLabelType}:                  _GotoState286Action,
	{_State223, NewlinesToken}:                          _GotoState287Action,
	{_State223, SemicolonToken}:                         _GotoState288Action,
	{_State225, CaseToken}:                              _GotoState289Action,
	{_State228, NewlinesToken}:                          _GotoState290Action,
	{_State228, SemicolonToken}:                         _GotoState291Action,
	{_State231, IdentifierToken}:                        _GotoState165Action,
	{_State231, GenericParameterDefType}:                _GotoState292Action,
	{_State233, IdentifierToken}:                        _GotoState146Action,
	{_State233, StructToken}:                            _GotoState31Action,
	{_State233, EnumToken}:                              _GotoState110Action,
	{_State233, TraitToken}:                             _GotoState15Action,
	{_State233, FuncToken}:                              _GotoState142Action,
	{_State233, LparenToken}:                            _GotoState113Action,
	{_State233, QuestionToken}:                          _GotoState114Action,
	{_State233, TildeTildeToken}:                        _GotoState115Action,
	{_State233, BitNegToken}:                            _GotoState109Action,
	{_State233, BitAndToken}:                            _GotoState108Action,
	{_State233, AtomTypeType}:                           _GotoState116Action,
	{_State233, TraitableTypeType}:                      _GotoState129Action,
	{_State233, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State233, ValueTypeType}:                          _GotoState293Action,
	{_State233, ImplicitStructDefType}:                  _GotoState122Action,
	{_State233, ExplicitStructDefType}:                  _GotoState118Action,
	{_State233, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State233, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State233, TraitDefType}:                           _GotoState126Action,
	{_State233, FuncTypeType}:                           _GotoState120Action,
	{_State234, NewlinesToken}:                          _GotoState294Action,
	{_State234, OrToken}:                                _GotoState295Action,
	{_State235, RparenToken}:                            _GotoState296Action,
	{_State236, AssignToken}:                            _GotoState246Action,
	{_State237, NewlinesToken}:                          _GotoState297Action,
	{_State237, OrToken}:                                _GotoState298Action,
	{_State238, IdentifierToken}:                        _GotoState240Action,
	{_State238, StructToken}:                            _GotoState31Action,
	{_State238, EnumToken}:                              _GotoState110Action,
	{_State238, TraitToken}:                             _GotoState15Action,
	{_State238, FuncToken}:                              _GotoState142Action,
	{_State238, LparenToken}:                            _GotoState113Action,
	{_State238, QuestionToken}:                          _GotoState114Action,
	{_State238, DotdotdotToken}:                         _GotoState239Action,
	{_State238, TildeTildeToken}:                        _GotoState115Action,
	{_State238, BitNegToken}:                            _GotoState109Action,
	{_State238, BitAndToken}:                            _GotoState108Action,
	{_State238, AtomTypeType}:                           _GotoState116Action,
	{_State238, TraitableTypeType}:                      _GotoState129Action,
	{_State238, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State238, ValueTypeType}:                          _GotoState244Action,
	{_State238, ImplicitStructDefType}:                  _GotoState122Action,
	{_State238, ExplicitStructDefType}:                  _GotoState118Action,
	{_State238, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State238, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State238, TraitDefType}:                           _GotoState126Action,
	{_State238, ParameterDeclType}:                      _GotoState242Action,
	{_State238, ParameterDeclsType}:                     _GotoState243Action,
	{_State238, OptionalParameterDeclsType}:             _GotoState299Action,
	{_State238, FuncTypeType}:                           _GotoState120Action,
	{_State239, IdentifierToken}:                        _GotoState146Action,
	{_State239, StructToken}:                            _GotoState31Action,
	{_State239, EnumToken}:                              _GotoState110Action,
	{_State239, TraitToken}:                             _GotoState15Action,
	{_State239, FuncToken}:                              _GotoState142Action,
	{_State239, LparenToken}:                            _GotoState113Action,
	{_State239, QuestionToken}:                          _GotoState114Action,
	{_State239, TildeTildeToken}:                        _GotoState115Action,
	{_State239, BitNegToken}:                            _GotoState109Action,
	{_State239, BitAndToken}:                            _GotoState108Action,
	{_State239, AtomTypeType}:                           _GotoState116Action,
	{_State239, TraitableTypeType}:                      _GotoState129Action,
	{_State239, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State239, ValueTypeType}:                          _GotoState300Action,
	{_State239, ImplicitStructDefType}:                  _GotoState122Action,
	{_State239, ExplicitStructDefType}:                  _GotoState118Action,
	{_State239, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State239, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State239, TraitDefType}:                           _GotoState126Action,
	{_State239, FuncTypeType}:                           _GotoState120Action,
	{_State240, IdentifierToken}:                        _GotoState146Action,
	{_State240, StructToken}:                            _GotoState31Action,
	{_State240, EnumToken}:                              _GotoState110Action,
	{_State240, TraitToken}:                             _GotoState15Action,
	{_State240, FuncToken}:                              _GotoState142Action,
	{_State240, LparenToken}:                            _GotoState113Action,
	{_State240, QuestionToken}:                          _GotoState114Action,
	{_State240, DollarLbracketToken}:                    _GotoState67Action,
	{_State240, DotdotdotToken}:                         _GotoState301Action,
	{_State240, TildeTildeToken}:                        _GotoState115Action,
	{_State240, BitNegToken}:                            _GotoState109Action,
	{_State240, BitAndToken}:                            _GotoState108Action,
	{_State240, OptionalGenericBindingType}:             _GotoState176Action,
	{_State240, AtomTypeType}:                           _GotoState116Action,
	{_State240, TraitableTypeType}:                      _GotoState129Action,
	{_State240, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State240, ValueTypeType}:                          _GotoState302Action,
	{_State240, ImplicitStructDefType}:                  _GotoState122Action,
	{_State240, ExplicitStructDefType}:                  _GotoState118Action,
	{_State240, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State240, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State240, TraitDefType}:                           _GotoState126Action,
	{_State240, FuncTypeType}:                           _GotoState120Action,
	{_State241, RparenToken}:                            _GotoState303Action,
	{_State243, CommaToken}:                             _GotoState304Action,
	{_State245, IdentifierToken}:                        _GotoState112Action,
	{_State245, StructToken}:                            _GotoState31Action,
	{_State245, EnumToken}:                              _GotoState110Action,
	{_State245, TraitToken}:                             _GotoState15Action,
	{_State245, FuncToken}:                              _GotoState142Action,
	{_State245, LparenToken}:                            _GotoState113Action,
	{_State245, QuestionToken}:                          _GotoState114Action,
	{_State245, TildeTildeToken}:                        _GotoState115Action,
	{_State245, BitNegToken}:                            _GotoState109Action,
	{_State245, BitAndToken}:                            _GotoState108Action,
	{_State245, AtomTypeType}:                           _GotoState116Action,
	{_State245, TraitableTypeType}:                      _GotoState129Action,
	{_State245, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State245, ValueTypeType}:                          _GotoState130Action,
	{_State245, FieldDefType}:                           _GotoState236Action,
	{_State245, ImplicitStructDefType}:                  _GotoState122Action,
	{_State245, ExplicitStructDefType}:                  _GotoState118Action,
	{_State245, EnumValueDefType}:                       _GotoState305Action,
	{_State245, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State245, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State245, TraitDefType}:                           _GotoState126Action,
	{_State245, FuncTypeType}:                           _GotoState120Action,
	{_State246, DefaultToken}:                           _GotoState306Action,
	{_State247, IdentifierToken}:                        _GotoState112Action,
	{_State247, StructToken}:                            _GotoState31Action,
	{_State247, EnumToken}:                              _GotoState110Action,
	{_State247, TraitToken}:                             _GotoState15Action,
	{_State247, FuncToken}:                              _GotoState142Action,
	{_State247, LparenToken}:                            _GotoState113Action,
	{_State247, QuestionToken}:                          _GotoState114Action,
	{_State247, TildeTildeToken}:                        _GotoState115Action,
	{_State247, BitNegToken}:                            _GotoState109Action,
	{_State247, BitAndToken}:                            _GotoState108Action,
	{_State247, AtomTypeType}:                           _GotoState116Action,
	{_State247, TraitableTypeType}:                      _GotoState129Action,
	{_State247, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State247, ValueTypeType}:                          _GotoState130Action,
	{_State247, FieldDefType}:                           _GotoState236Action,
	{_State247, ImplicitStructDefType}:                  _GotoState122Action,
	{_State247, ExplicitStructDefType}:                  _GotoState118Action,
	{_State247, EnumValueDefType}:                       _GotoState307Action,
	{_State247, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State247, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State247, TraitDefType}:                           _GotoState126Action,
	{_State247, FuncTypeType}:                           _GotoState120Action,
	{_State249, IdentifierToken}:                        _GotoState112Action,
	{_State249, StructToken}:                            _GotoState31Action,
	{_State249, EnumToken}:                              _GotoState110Action,
	{_State249, TraitToken}:                             _GotoState15Action,
	{_State249, FuncToken}:                              _GotoState142Action,
	{_State249, LparenToken}:                            _GotoState113Action,
	{_State249, QuestionToken}:                          _GotoState114Action,
	{_State249, TildeTildeToken}:                        _GotoState115Action,
	{_State249, BitNegToken}:                            _GotoState109Action,
	{_State249, BitAndToken}:                            _GotoState108Action,
	{_State249, AtomTypeType}:                           _GotoState116Action,
	{_State249, TraitableTypeType}:                      _GotoState129Action,
	{_State249, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State249, ValueTypeType}:                          _GotoState130Action,
	{_State249, FieldDefType}:                           _GotoState308Action,
	{_State249, ImplicitStructDefType}:                  _GotoState122Action,
	{_State249, ExplicitStructDefType}:                  _GotoState118Action,
	{_State249, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State249, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State249, TraitDefType}:                           _GotoState126Action,
	{_State249, FuncTypeType}:                           _GotoState120Action,
	{_State257, IdentifierToken}:                        _GotoState131Action,
	{_State257, ParameterDefType}:                       _GotoState135Action,
	{_State257, ParameterDefsType}:                      _GotoState136Action,
	{_State257, OptionalParameterDefsType}:              _GotoState309Action,
	{_State258, LbraceToken}:                            _GotoState95Action,
	{_State258, BlockBodyType}:                          _GotoState310Action,
	{_State265, IfToken}:                                _GotoState94Action,
	{_State265, LbraceToken}:                            _GotoState95Action,
	{_State265, BlockBodyType}:                          _GotoState311Action,
	{_State265, IfExprType}:                             _GotoState312Action,
	{_State266, LbracketToken}:                          _GotoState69Action,
	{_State266, DotToken}:                               _GotoState68Action,
	{_State266, DollarLbracketToken}:                    _GotoState67Action,
	{_State266, OptionalGenericBindingType}:             _GotoState71Action,
	{_State268, IdentifierToken}:                        _GotoState313Action,
	{_State282, IntegerLiteralToken}:                    _GotoState23Action,
	{_State282, FloatLiteralToken}:                      _GotoState20Action,
	{_State282, RuneLiteralToken}:                       _GotoState29Action,
	{_State282, StringLiteralToken}:                     _GotoState30Action,
	{_State282, IdentifierToken}:                        _GotoState22Action,
	{_State282, TrueToken}:                              _GotoState33Action,
	{_State282, FalseToken}:                             _GotoState19Action,
	{_State282, StructToken}:                            _GotoState31Action,
	{_State282, FuncToken}:                              _GotoState21Action,
	{_State282, LparenToken}:                            _GotoState26Action,
	{_State282, LabelDeclToken}:                         _GotoState24Action,
	{_State282, NotToken}:                               _GotoState28Action,
	{_State282, SubToken}:                               _GotoState32Action,
	{_State282, MulToken}:                               _GotoState27Action,
	{_State282, BitNegToken}:                            _GotoState18Action,
	{_State282, BitAndToken}:                            _GotoState17Action,
	{_State282, LexErrorToken}:                          _GotoState25Action,
	{_State282, LiteralType}:                            _GotoState44Action,
	{_State282, AnonymousStructExprType}:                _GotoState38Action,
	{_State282, AtomExprType}:                           _GotoState39Action,
	{_State282, CallExprType}:                           _GotoState41Action,
	{_State282, AccessExprType}:                         _GotoState34Action,
	{_State282, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State282, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State282, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State282, MulExprType}:                            _GotoState45Action,
	{_State282, AddExprType}:                            _GotoState35Action,
	{_State282, CmpExprType}:                            _GotoState42Action,
	{_State282, AndExprType}:                            _GotoState36Action,
	{_State282, OrExprType}:                             _GotoState47Action,
	{_State282, SequenceExprType}:                       _GotoState51Action,
	{_State282, OptionalLabelDeclType}:                  _GotoState46Action,
	{_State282, BlockExprType}:                          _GotoState40Action,
	{_State282, ExpressionType}:                         _GotoState314Action,
	{_State282, ExplicitStructDefType}:                  _GotoState43Action,
	{_State282, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State284, IntegerLiteralToken}:                    _GotoState23Action,
	{_State284, FloatLiteralToken}:                      _GotoState20Action,
	{_State284, RuneLiteralToken}:                       _GotoState29Action,
	{_State284, StringLiteralToken}:                     _GotoState30Action,
	{_State284, IdentifierToken}:                        _GotoState22Action,
	{_State284, TrueToken}:                              _GotoState33Action,
	{_State284, FalseToken}:                             _GotoState19Action,
	{_State284, StructToken}:                            _GotoState31Action,
	{_State284, FuncToken}:                              _GotoState21Action,
	{_State284, LparenToken}:                            _GotoState26Action,
	{_State284, LabelDeclToken}:                         _GotoState24Action,
	{_State284, NotToken}:                               _GotoState28Action,
	{_State284, SubToken}:                               _GotoState32Action,
	{_State284, MulToken}:                               _GotoState27Action,
	{_State284, BitNegToken}:                            _GotoState18Action,
	{_State284, BitAndToken}:                            _GotoState17Action,
	{_State284, LexErrorToken}:                          _GotoState25Action,
	{_State284, LiteralType}:                            _GotoState44Action,
	{_State284, AnonymousStructExprType}:                _GotoState38Action,
	{_State284, AtomExprType}:                           _GotoState39Action,
	{_State284, CallExprType}:                           _GotoState41Action,
	{_State284, AccessExprType}:                         _GotoState34Action,
	{_State284, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State284, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State284, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State284, MulExprType}:                            _GotoState45Action,
	{_State284, AddExprType}:                            _GotoState35Action,
	{_State284, CmpExprType}:                            _GotoState42Action,
	{_State284, AndExprType}:                            _GotoState36Action,
	{_State284, OrExprType}:                             _GotoState47Action,
	{_State284, SequenceExprType}:                       _GotoState51Action,
	{_State284, OptionalLabelDeclType}:                  _GotoState46Action,
	{_State284, BlockExprType}:                          _GotoState40Action,
	{_State284, ExpressionType}:                         _GotoState315Action,
	{_State284, ExplicitStructDefType}:                  _GotoState43Action,
	{_State284, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State286, IntegerLiteralToken}:                    _GotoState23Action,
	{_State286, FloatLiteralToken}:                      _GotoState20Action,
	{_State286, RuneLiteralToken}:                       _GotoState29Action,
	{_State286, StringLiteralToken}:                     _GotoState30Action,
	{_State286, IdentifierToken}:                        _GotoState22Action,
	{_State286, TrueToken}:                              _GotoState33Action,
	{_State286, FalseToken}:                             _GotoState19Action,
	{_State286, StructToken}:                            _GotoState31Action,
	{_State286, FuncToken}:                              _GotoState21Action,
	{_State286, LparenToken}:                            _GotoState26Action,
	{_State286, LabelDeclToken}:                         _GotoState24Action,
	{_State286, NotToken}:                               _GotoState28Action,
	{_State286, SubToken}:                               _GotoState32Action,
	{_State286, MulToken}:                               _GotoState27Action,
	{_State286, BitNegToken}:                            _GotoState18Action,
	{_State286, BitAndToken}:                            _GotoState17Action,
	{_State286, LexErrorToken}:                          _GotoState25Action,
	{_State286, LiteralType}:                            _GotoState44Action,
	{_State286, AnonymousStructExprType}:                _GotoState38Action,
	{_State286, AtomExprType}:                           _GotoState39Action,
	{_State286, CallExprType}:                           _GotoState41Action,
	{_State286, AccessExprType}:                         _GotoState34Action,
	{_State286, PostfixUnaryExprType}:                   _GotoState48Action,
	{_State286, PrefixUnaryOpType}:                      _GotoState50Action,
	{_State286, PrefixUnaryExprType}:                    _GotoState49Action,
	{_State286, MulExprType}:                            _GotoState45Action,
	{_State286, AddExprType}:                            _GotoState35Action,
	{_State286, CmpExprType}:                            _GotoState42Action,
	{_State286, AndExprType}:                            _GotoState36Action,
	{_State286, OrExprType}:                             _GotoState47Action,
	{_State286, SequenceExprType}:                       _GotoState51Action,
	{_State286, OptionalExpressionOrImplicitStructType}: _GotoState317Action,
	{_State286, ExpressionOrImplicitStructType}:         _GotoState316Action,
	{_State286, OptionalLabelDeclType}:                  _GotoState46Action,
	{_State286, BlockExprType}:                          _GotoState40Action,
	{_State286, ExpressionType}:                         _GotoState219Action,
	{_State286, ExplicitStructDefType}:                  _GotoState43Action,
	{_State286, AnonymousFuncExprType}:                  _GotoState37Action,
	{_State289, DefaultToken}:                           _GotoState318Action,
	{_State294, IdentifierToken}:                        _GotoState112Action,
	{_State294, StructToken}:                            _GotoState31Action,
	{_State294, EnumToken}:                              _GotoState110Action,
	{_State294, TraitToken}:                             _GotoState15Action,
	{_State294, FuncToken}:                              _GotoState142Action,
	{_State294, LparenToken}:                            _GotoState113Action,
	{_State294, QuestionToken}:                          _GotoState114Action,
	{_State294, TildeTildeToken}:                        _GotoState115Action,
	{_State294, BitNegToken}:                            _GotoState109Action,
	{_State294, BitAndToken}:                            _GotoState108Action,
	{_State294, AtomTypeType}:                           _GotoState116Action,
	{_State294, TraitableTypeType}:                      _GotoState129Action,
	{_State294, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State294, ValueTypeType}:                          _GotoState130Action,
	{_State294, FieldDefType}:                           _GotoState236Action,
	{_State294, ImplicitStructDefType}:                  _GotoState122Action,
	{_State294, ExplicitStructDefType}:                  _GotoState118Action,
	{_State294, EnumValueDefType}:                       _GotoState319Action,
	{_State294, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State294, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State294, TraitDefType}:                           _GotoState126Action,
	{_State294, FuncTypeType}:                           _GotoState120Action,
	{_State295, IdentifierToken}:                        _GotoState112Action,
	{_State295, StructToken}:                            _GotoState31Action,
	{_State295, EnumToken}:                              _GotoState110Action,
	{_State295, TraitToken}:                             _GotoState15Action,
	{_State295, FuncToken}:                              _GotoState142Action,
	{_State295, LparenToken}:                            _GotoState113Action,
	{_State295, QuestionToken}:                          _GotoState114Action,
	{_State295, TildeTildeToken}:                        _GotoState115Action,
	{_State295, BitNegToken}:                            _GotoState109Action,
	{_State295, BitAndToken}:                            _GotoState108Action,
	{_State295, AtomTypeType}:                           _GotoState116Action,
	{_State295, TraitableTypeType}:                      _GotoState129Action,
	{_State295, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State295, ValueTypeType}:                          _GotoState130Action,
	{_State295, FieldDefType}:                           _GotoState236Action,
	{_State295, ImplicitStructDefType}:                  _GotoState122Action,
	{_State295, ExplicitStructDefType}:                  _GotoState118Action,
	{_State295, EnumValueDefType}:                       _GotoState320Action,
	{_State295, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State295, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State295, TraitDefType}:                           _GotoState126Action,
	{_State295, FuncTypeType}:                           _GotoState120Action,
	{_State297, IdentifierToken}:                        _GotoState112Action,
	{_State297, StructToken}:                            _GotoState31Action,
	{_State297, EnumToken}:                              _GotoState110Action,
	{_State297, TraitToken}:                             _GotoState15Action,
	{_State297, FuncToken}:                              _GotoState142Action,
	{_State297, LparenToken}:                            _GotoState113Action,
	{_State297, QuestionToken}:                          _GotoState114Action,
	{_State297, TildeTildeToken}:                        _GotoState115Action,
	{_State297, BitNegToken}:                            _GotoState109Action,
	{_State297, BitAndToken}:                            _GotoState108Action,
	{_State297, AtomTypeType}:                           _GotoState116Action,
	{_State297, TraitableTypeType}:                      _GotoState129Action,
	{_State297, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State297, ValueTypeType}:                          _GotoState130Action,
	{_State297, FieldDefType}:                           _GotoState236Action,
	{_State297, ImplicitStructDefType}:                  _GotoState122Action,
	{_State297, ExplicitStructDefType}:                  _GotoState118Action,
	{_State297, EnumValueDefType}:                       _GotoState321Action,
	{_State297, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State297, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State297, TraitDefType}:                           _GotoState126Action,
	{_State297, FuncTypeType}:                           _GotoState120Action,
	{_State298, IdentifierToken}:                        _GotoState112Action,
	{_State298, StructToken}:                            _GotoState31Action,
	{_State298, EnumToken}:                              _GotoState110Action,
	{_State298, TraitToken}:                             _GotoState15Action,
	{_State298, FuncToken}:                              _GotoState142Action,
	{_State298, LparenToken}:                            _GotoState113Action,
	{_State298, QuestionToken}:                          _GotoState114Action,
	{_State298, TildeTildeToken}:                        _GotoState115Action,
	{_State298, BitNegToken}:                            _GotoState109Action,
	{_State298, BitAndToken}:                            _GotoState108Action,
	{_State298, AtomTypeType}:                           _GotoState116Action,
	{_State298, TraitableTypeType}:                      _GotoState129Action,
	{_State298, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State298, ValueTypeType}:                          _GotoState130Action,
	{_State298, FieldDefType}:                           _GotoState236Action,
	{_State298, ImplicitStructDefType}:                  _GotoState122Action,
	{_State298, ExplicitStructDefType}:                  _GotoState118Action,
	{_State298, EnumValueDefType}:                       _GotoState322Action,
	{_State298, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State298, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State298, TraitDefType}:                           _GotoState126Action,
	{_State298, FuncTypeType}:                           _GotoState120Action,
	{_State299, RparenToken}:                            _GotoState323Action,
	{_State301, IdentifierToken}:                        _GotoState146Action,
	{_State301, StructToken}:                            _GotoState31Action,
	{_State301, EnumToken}:                              _GotoState110Action,
	{_State301, TraitToken}:                             _GotoState15Action,
	{_State301, FuncToken}:                              _GotoState142Action,
	{_State301, LparenToken}:                            _GotoState113Action,
	{_State301, QuestionToken}:                          _GotoState114Action,
	{_State301, TildeTildeToken}:                        _GotoState115Action,
	{_State301, BitNegToken}:                            _GotoState109Action,
	{_State301, BitAndToken}:                            _GotoState108Action,
	{_State301, AtomTypeType}:                           _GotoState116Action,
	{_State301, TraitableTypeType}:                      _GotoState129Action,
	{_State301, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State301, ValueTypeType}:                          _GotoState324Action,
	{_State301, ImplicitStructDefType}:                  _GotoState122Action,
	{_State301, ExplicitStructDefType}:                  _GotoState118Action,
	{_State301, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State301, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State301, TraitDefType}:                           _GotoState126Action,
	{_State301, FuncTypeType}:                           _GotoState120Action,
	{_State303, IdentifierToken}:                        _GotoState146Action,
	{_State303, StructToken}:                            _GotoState31Action,
	{_State303, EnumToken}:                              _GotoState110Action,
	{_State303, TraitToken}:                             _GotoState15Action,
	{_State303, FuncToken}:                              _GotoState142Action,
	{_State303, LparenToken}:                            _GotoState113Action,
	{_State303, QuestionToken}:                          _GotoState114Action,
	{_State303, TildeTildeToken}:                        _GotoState115Action,
	{_State303, BitNegToken}:                            _GotoState109Action,
	{_State303, BitAndToken}:                            _GotoState108Action,
	{_State303, AtomTypeType}:                           _GotoState116Action,
	{_State303, TraitableTypeType}:                      _GotoState129Action,
	{_State303, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State303, ValueTypeType}:                          _GotoState259Action,
	{_State303, ImplicitStructDefType}:                  _GotoState122Action,
	{_State303, ExplicitStructDefType}:                  _GotoState118Action,
	{_State303, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State303, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State303, TraitDefType}:                           _GotoState126Action,
	{_State303, ReturnTypeType}:                         _GotoState325Action,
	{_State303, FuncTypeType}:                           _GotoState120Action,
	{_State304, IdentifierToken}:                        _GotoState240Action,
	{_State304, StructToken}:                            _GotoState31Action,
	{_State304, EnumToken}:                              _GotoState110Action,
	{_State304, TraitToken}:                             _GotoState15Action,
	{_State304, FuncToken}:                              _GotoState142Action,
	{_State304, LparenToken}:                            _GotoState113Action,
	{_State304, QuestionToken}:                          _GotoState114Action,
	{_State304, DotdotdotToken}:                         _GotoState239Action,
	{_State304, TildeTildeToken}:                        _GotoState115Action,
	{_State304, BitNegToken}:                            _GotoState109Action,
	{_State304, BitAndToken}:                            _GotoState108Action,
	{_State304, AtomTypeType}:                           _GotoState116Action,
	{_State304, TraitableTypeType}:                      _GotoState129Action,
	{_State304, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State304, ValueTypeType}:                          _GotoState244Action,
	{_State304, ImplicitStructDefType}:                  _GotoState122Action,
	{_State304, ExplicitStructDefType}:                  _GotoState118Action,
	{_State304, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State304, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State304, TraitDefType}:                           _GotoState126Action,
	{_State304, ParameterDeclType}:                      _GotoState326Action,
	{_State304, FuncTypeType}:                           _GotoState120Action,
	{_State309, RparenToken}:                            _GotoState327Action,
	{_State313, GreaterToken}:                           _GotoState328Action,
	{_State316, CommaToken}:                             _GotoState284Action,
	{_State318, RbraceToken}:                            _GotoState329Action,
	{_State323, IdentifierToken}:                        _GotoState146Action,
	{_State323, StructToken}:                            _GotoState31Action,
	{_State323, EnumToken}:                              _GotoState110Action,
	{_State323, TraitToken}:                             _GotoState15Action,
	{_State323, FuncToken}:                              _GotoState142Action,
	{_State323, LparenToken}:                            _GotoState113Action,
	{_State323, QuestionToken}:                          _GotoState114Action,
	{_State323, TildeTildeToken}:                        _GotoState115Action,
	{_State323, BitNegToken}:                            _GotoState109Action,
	{_State323, BitAndToken}:                            _GotoState108Action,
	{_State323, AtomTypeType}:                           _GotoState116Action,
	{_State323, TraitableTypeType}:                      _GotoState129Action,
	{_State323, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State323, ValueTypeType}:                          _GotoState259Action,
	{_State323, ImplicitStructDefType}:                  _GotoState122Action,
	{_State323, ExplicitStructDefType}:                  _GotoState118Action,
	{_State323, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State323, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State323, TraitDefType}:                           _GotoState126Action,
	{_State323, ReturnTypeType}:                         _GotoState330Action,
	{_State323, FuncTypeType}:                           _GotoState120Action,
	{_State327, IdentifierToken}:                        _GotoState146Action,
	{_State327, StructToken}:                            _GotoState31Action,
	{_State327, EnumToken}:                              _GotoState110Action,
	{_State327, TraitToken}:                             _GotoState15Action,
	{_State327, FuncToken}:                              _GotoState142Action,
	{_State327, LparenToken}:                            _GotoState113Action,
	{_State327, QuestionToken}:                          _GotoState114Action,
	{_State327, TildeTildeToken}:                        _GotoState115Action,
	{_State327, BitNegToken}:                            _GotoState109Action,
	{_State327, BitAndToken}:                            _GotoState108Action,
	{_State327, AtomTypeType}:                           _GotoState116Action,
	{_State327, TraitableTypeType}:                      _GotoState129Action,
	{_State327, TraitAlgebraTypeType}:                   _GotoState125Action,
	{_State327, ValueTypeType}:                          _GotoState259Action,
	{_State327, ImplicitStructDefType}:                  _GotoState122Action,
	{_State327, ExplicitStructDefType}:                  _GotoState118Action,
	{_State327, ImplicitEnumDefType}:                    _GotoState121Action,
	{_State327, ExplicitEnumDefType}:                    _GotoState117Action,
	{_State327, TraitDefType}:                           _GotoState126Action,
	{_State327, ReturnTypeType}:                         _GotoState331Action,
	{_State327, FuncTypeType}:                           _GotoState120Action,
	{_State328, StringLiteralToken}:                     _GotoState332Action,
	{_State331, LbraceToken}:                            _GotoState95Action,
	{_State331, BlockBodyType}:                          _GotoState333Action,
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
	{_State26, ColonToken}:                              _ReduceNilToOptionalExpressionAction,
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
	{_State60, _WildcardMarker}:                         _ReduceIdentifierToAtomExprAction,
	{_State61, _WildcardMarker}:                         _ReduceArgumentToArgumentsAction,
	{_State63, _WildcardMarker}:                         _ReduceColonExpressionsToArgumentAction,
	{_State64, ColonToken}:                              _ReduceExpressionToOptionalExpressionAction,
	{_State64, _WildcardMarker}:                         _ReducePositionalToArgumentAction,
	{_State66, RparenToken}:                             _ReduceNilToOptionalExplicitFieldDefsAction,
	{_State67, RbracketToken}:                           _ReduceNilToOptionalGenericArgumentsAction,
	{_State69, ColonToken}:                              _ReduceNilToOptionalExpressionAction,
	{_State69, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State70, _WildcardMarker}:                         _ReduceQuestionToPostfixUnaryExprAction,
	{_State72, _WildcardMarker}:                         _ReduceAddToAddOpAction,
	{_State73, _WildcardMarker}:                         _ReduceBitOrToAddOpAction,
	{_State74, _WildcardMarker}:                         _ReduceBitXorToAddOpAction,
	{_State75, _WildcardMarker}:                         _ReduceSubToAddOpAction,
	{_State76, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State77, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State78, _WildcardMarker}:                         _ReduceEqualToCmpOpAction,
	{_State79, _WildcardMarker}:                         _ReduceGreaterToCmpOpAction,
	{_State80, _WildcardMarker}:                         _ReduceGreaterOrEqualToCmpOpAction,
	{_State81, _WildcardMarker}:                         _ReduceLessToCmpOpAction,
	{_State82, _WildcardMarker}:                         _ReduceLessOrEqualToCmpOpAction,
	{_State83, _WildcardMarker}:                         _ReduceNotEqualToCmpOpAction,
	{_State84, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State85, ColonToken}:                              _ReduceNilToOptionalExpressionAction,
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
	{_State95, _WildcardMarker}:                         _ReduceEmptyListToStatementsAction,
	{_State96, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State97, _EndMarker}:                              _ReduceToBlockExprAction,
	{_State98, _EndMarker}:                              _ReduceIfExprToExpressionAction,
	{_State99, _EndMarker}:                              _ReduceLoopExprToExpressionAction,
	{_State100, _EndMarker}:                             _ReduceSwitchExprToExpressionAction,
	{_State101, LbraceToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State103, _WildcardMarker}:                        _ReducePrefixOpToPrefixUnaryExprAction,
	{_State104, _WildcardMarker}:                        _ReduceEmptyListToPackageStatementsAction,
	{_State105, RbracketToken}:                          _ReduceNilToOptionalGenericParameterDefsAction,
	{_State112, _WildcardMarker}:                        _ReduceNilToOptionalGenericBindingAction,
	{_State113, RparenToken}:                            _ReduceNilToOptionalImplicitFieldDefsAction,
	{_State114, _EndMarker}:                             _ReduceInferredToValueTypeAction,
	{_State116, _WildcardMarker}:                        _ReduceAtomTypeToTraitableTypeAction,
	{_State117, _WildcardMarker}:                        _ReduceExplicitEnumDefToAtomTypeAction,
	{_State118, _WildcardMarker}:                        _ReduceExplicitStructDefToAtomTypeAction,
	{_State119, _WildcardMarker}:                        _ReduceFieldDefToTraitPropertyAction,
	{_State120, _EndMarker}:                             _ReduceFuncTypeToValueTypeAction,
	{_State121, _WildcardMarker}:                        _ReduceImplicitEnumDefToAtomTypeAction,
	{_State122, _WildcardMarker}:                        _ReduceImplicitStructDefToAtomTypeAction,
	{_State123, _WildcardMarker}:                        _ReduceMethodSignatureToTraitPropertyAction,
	{_State125, _EndMarker}:                             _ReduceTraitAlgebraTypeToValueTypeAction,
	{_State126, _WildcardMarker}:                        _ReduceTraitDefToAtomTypeAction,
	{_State127, RparenToken}:                            _ReduceTraitPropertiesToOptionalTraitPropertiesAction,
	{_State128, _WildcardMarker}:                        _ReduceTraitPropertyToTraitPropertiesAction,
	{_State129, _WildcardMarker}:                        _ReduceTraitableTypeToTraitAlgebraTypeAction,
	{_State130, _WildcardMarker}:                        _ReduceImplicitToFieldDefAction,
	{_State133, LparenToken}:                            _ReduceNilToOptionalGenericParametersAction,
	{_State135, _WildcardMarker}:                        _ReduceParameterDefToParameterDefsAction,
	{_State136, RparenToken}:                            _ReduceParameterDefsToOptionalParameterDefsAction,
	{_State137, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State138, ColonToken}:                             _ReduceNilToOptionalExpressionAction,
	{_State138, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State139, _WildcardMarker}:                        _ReduceImplicitToAnonymousStructExprAction,
	{_State140, ColonToken}:                             _ReduceNilToOptionalExpressionAction,
	{_State140, CommaToken}:                             _ReduceNilToOptionalExpressionAction,
	{_State140, RbracketToken}:                          _ReduceNilToOptionalExpressionAction,
	{_State140, RparenToken}:                            _ReduceNilToOptionalExpressionAction,
	{_State140, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State141, _WildcardMarker}:                        _ReduceNilToOptionalExpressionAction,
	{_State141, ForToken}:                               _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State141, IfToken}:                                _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State141, LbraceToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State141, SwitchToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State143, RparenToken}:                            _ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction,
	{_State144, _WildcardMarker}:                        _ReduceFieldDefToExplicitFieldDefsAction,
	{_State146, _WildcardMarker}:                        _ReduceNilToOptionalGenericBindingAction,
	{_State147, RbracketToken}:                          _ReduceGenericArgumentsToOptionalGenericArgumentsAction,
	{_State149, _WildcardMarker}:                        _ReduceValueTypeToGenericArgumentsAction,
	{_State150, _WildcardMarker}:                        _ReduceAccessToAccessExprAction,
	{_State152, ColonToken}:                             _ReduceNilToOptionalExpressionAction,
	{_State152, RparenToken}:                            _ReduceNilToOptionalArgumentsAction,
	{_State152, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State153, _WildcardMarker}:                        _ReduceOpToAddExprAction,
	{_State154, _WildcardMarker}:                        _ReduceOpToAndExprAction,
	{_State155, _WildcardMarker}:                        _ReduceOpToCmpExprAction,
	{_State157, _WildcardMarker}:                        _ReduceOpToMulExprAction,
	{_State158, _WildcardMarker}:                        _ReduceBlockExprToAtomExprAction,
	{_State158, _EndMarker}:                             _ReduceInfiniteToLoopExprAction,
	{_State159, LbraceToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State161, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State163, _WildcardMarker}:                        _ReduceOpToOrExprAction,
	{_State165, _WildcardMarker}:                        _ReduceUnconstrainedToGenericParameterDefAction,
	{_State166, _WildcardMarker}:                        _ReduceGenericParameterDefToGenericParameterDefsAction,
	{_State167, RbracketToken}:                          _ReduceGenericParameterDefsToOptionalGenericParameterDefsAction,
	{_State169, _EndMarker}:                             _ReduceAliasToTypeDefAction,
	{_State170, _EndMarker}:                             _ReduceDefinitionToTypeDefAction,
	{_State171, _EndMarker}:                             _ReduceReferenceToValueTypeAction,
	{_State172, _WildcardMarker}:                        _ReducePublicMethodsTraitToTraitableTypeAction,
	{_State175, RparenToken}:                            _ReduceNilToOptionalParameterDeclsAction,
	{_State176, _WildcardMarker}:                        _ReduceNamedToAtomTypeAction,
	{_State177, _WildcardMarker}:                        _ReduceExplicitToFieldDefAction,
	{_State179, _WildcardMarker}:                        _ReduceFieldDefToImplicitFieldDefsAction,
	{_State179, OrToken}:                                _ReduceFieldDefToEnumValueDefAction,
	{_State181, RparenToken}:                            _ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction,
	{_State183, _WildcardMarker}:                        _ReducePublicTraitToTraitableTypeAction,
	{_State184, _EndMarker}:                             _ReduceToTraitDefAction,
	{_State191, _WildcardMarker}:                        _ReduceArgToParameterDefAction,
	{_State192, IdentifierToken}:                        _ReduceReceiverToOptionalReceiverAction,
	{_State194, LbraceToken}:                            _ReduceNilToReturnTypeAction,
	{_State196, _WildcardMarker}:                        _ReduceNamedToArgumentAction,
	{_State197, _WildcardMarker}:                        _ReduceAddToArgumentsAction,
	{_State198, _WildcardMarker}:                        _ReduceExpressionToOptionalExpressionAction,
	{_State199, _WildcardMarker}:                        _ReduceAddToColonExpressionsAction,
	{_State200, _WildcardMarker}:                        _ReducePairToColonExpressionsAction,
	{_State203, _WildcardMarker}:                        _ReduceToExplicitStructDefAction,
	{_State205, _WildcardMarker}:                        _ReduceBindingToOptionalGenericBindingAction,
	{_State206, _WildcardMarker}:                        _ReduceIndexToAccessExprAction,
	{_State207, RparenToken}:                            _ReduceArgumentsToOptionalArgumentsAction,
	{_State209, _WildcardMarker}:                        _ReduceExplicitToAnonymousStructExprAction,
	{_State210, _EndMarker}:                             _ReduceWhileToLoopExprAction,
	{_State211, _WildcardMarker}:                        _ReduceNoElseToIfExprAction,
	{_State212, LbraceToken}:                            _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State213, _WildcardMarker}:                        _ReduceBreakToJumpTypeAction,
	{_State214, _WildcardMarker}:                        _ReduceContinueToJumpTypeAction,
	{_State215, _EndMarker}:                             _ReduceToBlockBodyAction,
	{_State216, _WildcardMarker}:                        _ReduceReturnToJumpTypeAction,
	{_State218, _WildcardMarker}:                        _ReduceAccessExprToPostfixUnaryExprAction,
	{_State218, LparenToken}:                            _ReduceNilToOptionalGenericBindingAction,
	{_State219, _WildcardMarker}:                        _ReduceExpressionToExpressionOrImplicitStructAction,
	{_State220, _WildcardMarker}:                        _ReduceExpressionOrImplicitStructToStatementBodyAction,
	{_State221, _WildcardMarker}:                        _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State222, _WildcardMarker}:                        _ReduceAddToStatementsAction,
	{_State224, _WildcardMarker}:                        _ReduceUnsafeStatementToStatementBodyAction,
	{_State226, _EndMarker}:                             _ReduceWithSpecToPackageDefAction,
	{_State227, _WildcardMarker}:                        _ReduceAddToPackageStatementsAction,
	{_State229, _WildcardMarker}:                        _ReduceToPackageStatementBodyAction,
	{_State230, _WildcardMarker}:                        _ReduceConstrainedToGenericParameterDefAction,
	{_State232, _WildcardMarker}:                        _ReduceGenericToOptionalGenericParametersAction,
	{_State236, _WildcardMarker}:                        _ReduceFieldDefToEnumValueDefAction,
	{_State238, RparenToken}:                            _ReduceNilToOptionalParameterDeclsAction,
	{_State240, _WildcardMarker}:                        _ReduceNilToOptionalGenericBindingAction,
	{_State242, _WildcardMarker}:                        _ReduceParameterDeclToParameterDeclsAction,
	{_State243, RparenToken}:                            _ReduceParameterDeclsToOptionalParameterDeclsAction,
	{_State244, _WildcardMarker}:                        _ReduceUnamedToParameterDeclAction,
	{_State248, _WildcardMarker}:                        _ReduceToImplicitEnumDefAction,
	{_State250, _WildcardMarker}:                        _ReduceToImplicitStructDefAction,
	{_State251, _WildcardMarker}:                        _ReduceUnionToTraitAlgebraTypeAction,
	{_State252, _WildcardMarker}:                        _ReduceIntersectToTraitAlgebraTypeAction,
	{_State253, _WildcardMarker}:                        _ReduceDifferenceToTraitAlgebraTypeAction,
	{_State254, _WildcardMarker}:                        _ReduceExplicitToTraitPropertiesAction,
	{_State255, _WildcardMarker}:                        _ReduceImplicitToTraitPropertiesAction,
	{_State256, _WildcardMarker}:                        _ReduceVarargToParameterDefAction,
	{_State257, RparenToken}:                            _ReduceNilToOptionalParameterDefsAction,
	{_State259, _EndMarker}:                             _ReduceValueTypeToReturnTypeAction,
	{_State260, _WildcardMarker}:                        _ReduceAddToParameterDefsAction,
	{_State261, _WildcardMarker}:                        _ReduceExplicitToExplicitFieldDefsAction,
	{_State262, _WildcardMarker}:                        _ReduceImplicitToExplicitFieldDefsAction,
	{_State263, _WildcardMarker}:                        _ReduceAddToGenericArgumentsAction,
	{_State264, _WildcardMarker}:                        _ReduceToCallExprAction,
	{_State266, LparenToken}:                            _ReduceNilToOptionalGenericBindingAction,
	{_State267, _WildcardMarker}:                        _ReduceCallExprToAccessExprAction,
	{_State267, NewlinesToken}:                          _ReduceAsyncToStatementBodyAction,
	{_State267, SemicolonToken}:                         _ReduceAsyncToStatementBodyAction,
	{_State269, _WildcardMarker}:                        _ReduceAddAssignToBinaryOpAssignAction,
	{_State270, _WildcardMarker}:                        _ReduceAddOneAssignToOpOneAssignAction,
	{_State271, _WildcardMarker}:                        _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State272, _WildcardMarker}:                        _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State273, _WildcardMarker}:                        _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State274, _WildcardMarker}:                        _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State275, _WildcardMarker}:                        _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State276, _WildcardMarker}:                        _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State277, _WildcardMarker}:                        _ReduceDivAssignToBinaryOpAssignAction,
	{_State278, _WildcardMarker}:                        _ReduceModAssignToBinaryOpAssignAction,
	{_State279, _WildcardMarker}:                        _ReduceMulAssignToBinaryOpAssignAction,
	{_State280, _WildcardMarker}:                        _ReduceSubAssignToBinaryOpAssignAction,
	{_State281, _WildcardMarker}:                        _ReduceSubOneAssignToOpOneAssignAction,
	{_State282, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State283, _WildcardMarker}:                        _ReduceOpOneAssignToStatementBodyAction,
	{_State284, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State285, _WildcardMarker}:                        _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State286, NewlinesToken}:                          _ReduceNilToOptionalExpressionOrImplicitStructAction,
	{_State286, SemicolonToken}:                         _ReduceNilToOptionalExpressionOrImplicitStructAction,
	{_State286, _WildcardMarker}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State287, _WildcardMarker}:                        _ReduceImplicitToStatementAction,
	{_State288, _WildcardMarker}:                        _ReduceExplicitToStatementAction,
	{_State290, _WildcardMarker}:                        _ReduceImplicitToPackageStatementAction,
	{_State291, _WildcardMarker}:                        _ReduceExplicitToPackageStatementAction,
	{_State292, _WildcardMarker}:                        _ReduceAddToGenericParameterDefsAction,
	{_State293, _EndMarker}:                             _ReduceConstrainedDefToTypeDefAction,
	{_State296, _WildcardMarker}:                        _ReduceToExplicitEnumDefAction,
	{_State300, _WildcardMarker}:                        _ReduceUnnamedVarargToParameterDeclAction,
	{_State302, _WildcardMarker}:                        _ReduceArgToParameterDeclAction,
	{_State303, _WildcardMarker}:                        _ReduceNilToReturnTypeAction,
	{_State305, _WildcardMarker}:                        _ReducePairToImplicitEnumValueDefsAction,
	{_State306, _WildcardMarker}:                        _ReduceDefaultToEnumValueDefAction,
	{_State307, _WildcardMarker}:                        _ReduceAddToImplicitEnumValueDefsAction,
	{_State308, _WildcardMarker}:                        _ReduceAddToImplicitFieldDefsAction,
	{_State310, _WildcardMarker}:                        _ReduceToAnonymousFuncExprAction,
	{_State311, _EndMarker}:                             _ReduceIfElseToIfExprAction,
	{_State312, _EndMarker}:                             _ReduceMultiIfElseToIfExprAction,
	{_State314, _WildcardMarker}:                        _ReduceBinaryOpAssignToStatementBodyAction,
	{_State315, _WildcardMarker}:                        _ReduceImplicitStructToExpressionOrImplicitStructAction,
	{_State316, _WildcardMarker}:                        _ReduceExpressionOrImplicitStructToOptionalExpressionOrImplicitStructAction,
	{_State317, _WildcardMarker}:                        _ReduceJumpToStatementBodyAction,
	{_State319, RparenToken}:                            _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State320, _WildcardMarker}:                        _ReducePairToImplicitEnumValueDefsAction,
	{_State320, RparenToken}:                            _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State321, RparenToken}:                            _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State322, _WildcardMarker}:                        _ReduceAddToImplicitEnumValueDefsAction,
	{_State322, RparenToken}:                            _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State323, _WildcardMarker}:                        _ReduceNilToReturnTypeAction,
	{_State324, _WildcardMarker}:                        _ReduceVarargToParameterDeclAction,
	{_State325, _EndMarker}:                             _ReduceToFuncTypeAction,
	{_State326, _WildcardMarker}:                        _ReduceAddToParameterDeclsAction,
	{_State327, LbraceToken}:                            _ReduceNilToReturnTypeAction,
	{_State329, _EndMarker}:                             _ReduceToSwitchExprAction,
	{_State330, _WildcardMarker}:                        _ReduceToMethodSignatureAction,
	{_State332, _WildcardMarker}:                        _ReduceToUnsafeStatementAction,
	{_State333, _EndMarker}:                             _ReduceToNamedFuncDefAction,
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
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 60
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
      optional_expression -> State 65
      colon_expressions -> State 63
      argument -> State 61
      arguments -> State 62
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
      expression -> State 64
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
      LPAREN -> State 66

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
      access_expr: access_expr.LBRACKET argument RBRACKET
      postfix_unary_expr: access_expr., *
      postfix_unary_expr: access_expr.QUESTION
    Reduce:
      * -> [postfix_unary_expr]
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 69
      DOT -> State 68
      QUESTION -> State 70
      DOLLAR_LBRACKET -> State 67
      optional_generic_binding -> State 71

  State 35:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 72
      SUB -> State 75
      BIT_XOR -> State 74
      BIT_OR -> State 73
      add_op -> State 76

  State 36:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 77

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
      EQUAL -> State 78
      NOT_EQUAL -> State 83
      LESS -> State 81
      LESS_OR_EQUAL -> State 82
      GREATER -> State 79
      GREATER_OR_EQUAL -> State 80
      cmp_op -> State 84

  State 43:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_def.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 85

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
      MUL -> State 91
      DIV -> State 89
      MOD -> State 90
      BIT_AND -> State 86
      BIT_LSHIFT -> State 87
      BIT_RSHIFT -> State 88
      mul_op -> State 92

  State 46:
    Kernel Items:
      block_expr: optional_label_decl.block_body
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 94
      SWITCH -> State 96
      FOR -> State 93
      LBRACE -> State 95
      block_body -> State 97
      if_expr -> State 98
      switch_expr -> State 100
      loop_expr -> State 99

  State 47:
    Kernel Items:
      or_expr: or_expr.OR and_expr
      sequence_expr: or_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      OR -> State 101

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
      prefix_unary_expr -> State 103
      optional_label_decl -> State 102
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
      LPAREN -> State 104

  State 55:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE IDENTIFIER.EQUAL value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 105
      EQUAL -> State 106
      optional_generic_parameters -> State 107

  State 56:
    Kernel Items:
      trait_def: TRAIT LPAREN.optional_trait_properties RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 111
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 130
      field_def -> State 119
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_property -> State 128
      trait_properties -> State 127
      optional_trait_properties -> State 124
      trait_def -> State 126
      func_type -> State 120
      method_signature -> State 123

  State 57:
    Kernel Items:
      optional_receiver: LPAREN.parameter_def RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      parameter_def -> State 132

  State 58:
    Kernel Items:
      named_func_def: FUNC optional_receiver.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 133

  State 59:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 131
      parameter_def -> State 135
      parameter_defs -> State 136
      optional_parameter_defs -> State 134

  State 60:
    Kernel Items:
      atom_expr: IDENTIFIER., *
      argument: IDENTIFIER.ASSIGN expression
    Reduce:
      * -> [atom_expr]
    Goto:
      ASSIGN -> State 137

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
      RPAREN -> State 139
      COMMA -> State 138

  State 63:
    Kernel Items:
      colon_expressions: colon_expressions.COLON optional_expression
      argument: colon_expressions., *
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 140

  State 64:
    Kernel Items:
      optional_expression: expression., COLON
      argument: expression., *
    Reduce:
      * -> [argument]
      COLON -> [optional_expression]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      colon_expressions: optional_expression.COLON optional_expression
    Reduce:
      (nil)
    Goto:
      COLON -> State 141

  State 66:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN.optional_explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 130
      field_def -> State 144
      implicit_struct_def -> State 122
      explicit_field_defs -> State 143
      optional_explicit_field_defs -> State 145
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 67:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      generic_arguments -> State 147
      optional_generic_arguments -> State 148
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 149
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 68:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 150

  State 69:
    Kernel Items:
      access_expr: access_expr LBRACKET.argument RBRACKET
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 60
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
      optional_expression -> State 65
      colon_expressions -> State 63
      argument -> State 151
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
      expression -> State 64
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 70:
    Kernel Items:
      postfix_unary_expr: access_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 71:
    Kernel Items:
      call_expr: access_expr optional_generic_binding.LPAREN optional_arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 152

  State 72:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 73:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 74:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 75:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 76:
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
      mul_expr -> State 153
      optional_label_decl -> State 102
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 77:
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
      cmp_expr -> State 154
      optional_label_decl -> State 102
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 78:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 79:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 80:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 81:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 82:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 84:
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
      add_expr -> State 155
      optional_label_decl -> State 102
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 85:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_def LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 60
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
      optional_expression -> State 65
      colon_expressions -> State 63
      argument -> State 61
      arguments -> State 156
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
      expression -> State 64
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

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
      prefix_unary_expr -> State 157
      optional_label_decl -> State 102
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 93:
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
      sequence_expr -> State 159
      optional_label_decl -> State 102
      block_expr -> State 158
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 94:
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
      sequence_expr -> State 160
      optional_label_decl -> State 102
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 95:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 161

  State 96:
    Kernel Items:
      switch_expr: SWITCH.sequence_expr LBRACE CASE DEFAULT RBRACE
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
      sequence_expr -> State 162
      optional_label_decl -> State 102
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 97:
    Kernel Items:
      block_expr: optional_label_decl block_body., $
    Reduce:
      $ -> [block_expr]
    Goto:
      (nil)

  State 98:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 99:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 100:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 101:
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
      and_expr -> State 163
      optional_label_decl -> State 102
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 102:
    Kernel Items:
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 95
      block_body -> State 97

  State 103:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 104:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN.package_statements RPAREN
    Reduce:
      * -> [package_statements]
    Goto:
      package_statements -> State 164

  State 105:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.optional_generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 165
      generic_parameter_def -> State 166
      generic_parameter_defs -> State 167
      optional_generic_parameter_defs -> State 168

  State 106:
    Kernel Items:
      type_def: TYPE IDENTIFIER EQUAL.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 169
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 107:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 170
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 108:
    Kernel Items:
      value_type: BIT_AND.trait_algebra_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      LPAREN -> State 113
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 171
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126

  State 109:
    Kernel Items:
      traitable_type: BIT_NEG.atom_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      LPAREN -> State 113
      atom_type -> State 172
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126

  State 110:
    Kernel Items:
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 173

  State 111:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 174
      LPAREN -> State 175

  State 112:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      DOLLAR_LBRACKET -> State 67
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      optional_generic_binding -> State 176
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 177
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 113:
    Kernel Items:
      implicit_struct_def: LPAREN.optional_implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 130
      field_def -> State 179
      implicit_field_defs -> State 181
      optional_implicit_field_defs -> State 182
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      enum_value_def -> State 178
      implicit_enum_value_defs -> State 180
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 114:
    Kernel Items:
      value_type: QUESTION., $
    Reduce:
      $ -> [value_type]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      traitable_type: TILDE_TILDE.atom_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      LPAREN -> State 113
      atom_type -> State 183
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126

  State 116:
    Kernel Items:
      traitable_type: atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 117:
    Kernel Items:
      atom_type: explicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 118:
    Kernel Items:
      atom_type: explicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 119:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      value_type: func_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      (nil)

  State 121:
    Kernel Items:
      atom_type: implicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 122:
    Kernel Items:
      atom_type: implicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 123:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 184

  State 125:
    Kernel Items:
      trait_algebra_type: trait_algebra_type.MUL traitable_type
      trait_algebra_type: trait_algebra_type.ADD traitable_type
      trait_algebra_type: trait_algebra_type.SUB traitable_type
      value_type: trait_algebra_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      ADD -> State 185
      SUB -> State 187
      MUL -> State 186

  State 126:
    Kernel Items:
      atom_type: trait_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 127:
    Kernel Items:
      trait_properties: trait_properties.NEWLINES trait_property
      trait_properties: trait_properties.COMMA trait_property
      optional_trait_properties: trait_properties., RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      NEWLINES -> State 189
      COMMA -> State 188

  State 128:
    Kernel Items:
      trait_properties: trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 129:
    Kernel Items:
      trait_algebra_type: traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 130:
    Kernel Items:
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 131:
    Kernel Items:
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      DOTDOTDOT -> State 190
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 191
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 132:
    Kernel Items:
      optional_receiver: LPAREN parameter_def.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 192

  State 133:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 105
      optional_generic_parameters -> State 193

  State 134:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 194

  State 135:
    Kernel Items:
      parameter_defs: parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 136:
    Kernel Items:
      parameter_defs: parameter_defs.COMMA parameter_def
      optional_parameter_defs: parameter_defs., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      COMMA -> State 195

  State 137:
    Kernel Items:
      argument: IDENTIFIER ASSIGN.expression
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
      expression -> State 196
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 138:
    Kernel Items:
      arguments: arguments COMMA.argument
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 60
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
      optional_expression -> State 65
      colon_expressions -> State 63
      argument -> State 197
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
      expression -> State 64
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 139:
    Kernel Items:
      anonymous_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 140:
    Kernel Items:
      colon_expressions: colon_expressions COLON.optional_expression
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_expression]
      RBRACKET -> [optional_expression]
      COMMA -> [optional_expression]
      COLON -> [optional_expression]
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
      optional_expression -> State 199
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
      expression -> State 198
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 141:
    Kernel Items:
      colon_expressions: optional_expression COLON.optional_expression
    Reduce:
      * -> [optional_expression]
      IF -> [optional_label_decl]
      SWITCH -> [optional_label_decl]
      FOR -> [optional_label_decl]
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
      optional_expression -> State 200
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
      expression -> State 198
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 142:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 175

  State 143:
    Kernel Items:
      explicit_field_defs: explicit_field_defs.NEWLINES field_def
      explicit_field_defs: explicit_field_defs.COMMA field_def
      optional_explicit_field_defs: explicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      NEWLINES -> State 202
      COMMA -> State 201

  State 144:
    Kernel Items:
      explicit_field_defs: field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 145:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 203

  State 146:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 67
      optional_generic_binding -> State 176

  State 147:
    Kernel Items:
      generic_arguments: generic_arguments.COMMA value_type
      optional_generic_arguments: generic_arguments., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 204

  State 148:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 205

  State 149:
    Kernel Items:
      generic_arguments: value_type., *
    Reduce:
      * -> [generic_arguments]
    Goto:
      (nil)

  State 150:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 151:
    Kernel Items:
      access_expr: access_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 206

  State 152:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN.optional_arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_arguments]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 23
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 29
      STRING_LITERAL -> State 30
      IDENTIFIER -> State 60
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
      optional_expression -> State 65
      colon_expressions -> State 63
      argument -> State 61
      arguments -> State 207
      optional_arguments -> State 208
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
      expression -> State 64
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 153:
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

  State 154:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 78
      NOT_EQUAL -> State 83
      LESS -> State 81
      LESS_OR_EQUAL -> State 82
      GREATER -> State 79
      GREATER_OR_EQUAL -> State 80
      cmp_op -> State 84

  State 155:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 72
      SUB -> State 75
      BIT_XOR -> State 74
      BIT_OR -> State 73
      add_op -> State 76

  State 156:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_def LPAREN arguments.RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 209
      COMMA -> State 138

  State 157:
    Kernel Items:
      mul_expr: mul_expr mul_op prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 158:
    Kernel Items:
      atom_expr: block_expr., *
      loop_expr: FOR block_expr., $
    Reduce:
      * -> [atom_expr]
      $ -> [loop_expr]
    Goto:
      (nil)

  State 159:
    Kernel Items:
      loop_expr: FOR sequence_expr.block_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      LABEL_DECL -> State 24
      optional_label_decl -> State 102
      block_expr -> State 210

  State 160:
    Kernel Items:
      if_expr: IF sequence_expr.block_body
      if_expr: IF sequence_expr.block_body ELSE block_body
      if_expr: IF sequence_expr.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 95
      block_body -> State 211

  State 161:
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
      RETURN -> State 216
      BREAK -> State 213
      CONTINUE -> State 214
      UNSAFE -> State 217
      STRUCT -> State 31
      FUNC -> State 21
      ASYNC -> State 212
      RBRACE -> State 215
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
      access_expr -> State 218
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      sequence_expr -> State 51
      jump_type -> State 221
      expression_or_implicit_struct -> State 220
      unsafe_statement -> State 224
      statement_body -> State 223
      statement -> State 222
      optional_label_decl -> State 46
      block_expr -> State 40
      expression -> State 219
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 162:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.LBRACE CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 225

  State 163:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 77

  State 164:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements.RPAREN
      package_statements: package_statements.package_statement
    Reduce:
      (nil)
    Goto:
      UNSAFE -> State 217
      RPAREN -> State 226
      unsafe_statement -> State 229
      package_statement_body -> State 228
      package_statement -> State 227

  State 165:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 230
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 166:
    Kernel Items:
      generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 167:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs.COMMA generic_parameter_def
      optional_generic_parameter_defs: generic_parameter_defs., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      COMMA -> State 231

  State 168:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 232

  State 169:
    Kernel Items:
      type_def: TYPE IDENTIFIER EQUAL value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      (nil)

  State 170:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., $
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      $ -> [type_def]
    Goto:
      IMPLEMENTS -> State 233

  State 171:
    Kernel Items:
      trait_algebra_type: trait_algebra_type.MUL traitable_type
      trait_algebra_type: trait_algebra_type.ADD traitable_type
      trait_algebra_type: trait_algebra_type.SUB traitable_type
      value_type: BIT_AND trait_algebra_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      ADD -> State 185
      SUB -> State 187
      MUL -> State 186

  State 172:
    Kernel Items:
      traitable_type: BIT_NEG atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 173:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 130
      field_def -> State 236
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      enum_value_def -> State 234
      implicit_enum_value_defs -> State 237
      implicit_enum_def -> State 121
      explicit_enum_value_defs -> State 235
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 174:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 238

  State 175:
    Kernel Items:
      func_type: FUNC LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 240
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      DOTDOTDOT -> State 239
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 244
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      parameter_decl -> State 242
      parameter_decls -> State 243
      optional_parameter_decls -> State 241
      func_type -> State 120

  State 176:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 177:
    Kernel Items:
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 178:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 245

  State 179:
    Kernel Items:
      implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 246

  State 180:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 248
      OR -> State 247

  State 181:
    Kernel Items:
      implicit_field_defs: implicit_field_defs.COMMA field_def
      optional_implicit_field_defs: implicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      COMMA -> State 249

  State 182:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 250

  State 183:
    Kernel Items:
      traitable_type: TILDE_TILDE atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 184:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties RPAREN., $
    Reduce:
      $ -> [trait_def]
    Goto:
      (nil)

  State 185:
    Kernel Items:
      trait_algebra_type: trait_algebra_type ADD.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      LPAREN -> State 113
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      atom_type -> State 116
      traitable_type -> State 251
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126

  State 186:
    Kernel Items:
      trait_algebra_type: trait_algebra_type MUL.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      LPAREN -> State 113
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      atom_type -> State 116
      traitable_type -> State 252
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126

  State 187:
    Kernel Items:
      trait_algebra_type: trait_algebra_type SUB.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      LPAREN -> State 113
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      atom_type -> State 116
      traitable_type -> State 253
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126

  State 188:
    Kernel Items:
      trait_properties: trait_properties COMMA.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 111
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 130
      field_def -> State 119
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_property -> State 254
      trait_def -> State 126
      func_type -> State 120
      method_signature -> State 123

  State 189:
    Kernel Items:
      trait_properties: trait_properties NEWLINES.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 111
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 130
      field_def -> State 119
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_property -> State 255
      trait_def -> State 126
      func_type -> State 120
      method_signature -> State 123

  State 190:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 256
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 191:
    Kernel Items:
      parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      (nil)

  State 192:
    Kernel Items:
      optional_receiver: LPAREN parameter_def RPAREN., IDENTIFIER
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      (nil)

  State 193:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 257

  State 194:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 259
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      return_type -> State 258
      func_type -> State 120

  State 195:
    Kernel Items:
      parameter_defs: parameter_defs COMMA.parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 131
      parameter_def -> State 260

  State 196:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 197:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 198:
    Kernel Items:
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      (nil)

  State 199:
    Kernel Items:
      colon_expressions: colon_expressions COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 200:
    Kernel Items:
      colon_expressions: optional_expression COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 201:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 130
      field_def -> State 261
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 202:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 130
      field_def -> State 262
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 203:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_def]
    Goto:
      (nil)

  State 204:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 263
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 205:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 206:
    Kernel Items:
      access_expr: access_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 207:
    Kernel Items:
      arguments: arguments.COMMA argument
      optional_arguments: arguments., RPAREN
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 138

  State 208:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 264

  State 209:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_def LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      loop_expr: FOR sequence_expr block_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 211:
    Kernel Items:
      if_expr: IF sequence_expr block_body., *
      if_expr: IF sequence_expr block_body.ELSE block_body
      if_expr: IF sequence_expr block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 265

  State 212:
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
      call_expr -> State 267
      access_expr -> State 266
      optional_label_decl -> State 102
      block_expr -> State 40
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 213:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 214:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 215:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 216:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 268

  State 218:
    Kernel Items:
      call_expr: access_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET argument RBRACKET
      postfix_unary_expr: access_expr., *
      postfix_unary_expr: access_expr.QUESTION
      statement_body: access_expr.op_one_assign
      statement_body: access_expr.binary_op_assign expression
    Reduce:
      * -> [postfix_unary_expr]
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 69
      DOT -> State 68
      QUESTION -> State 70
      DOLLAR_LBRACKET -> State 67
      ADD_ASSIGN -> State 269
      SUB_ASSIGN -> State 280
      MUL_ASSIGN -> State 279
      DIV_ASSIGN -> State 277
      MOD_ASSIGN -> State 278
      ADD_ONE_ASSIGN -> State 270
      SUB_ONE_ASSIGN -> State 281
      BIT_NEG_ASSIGN -> State 273
      BIT_AND_ASSIGN -> State 271
      BIT_OR_ASSIGN -> State 274
      BIT_XOR_ASSIGN -> State 276
      BIT_LSHIFT_ASSIGN -> State 272
      BIT_RSHIFT_ASSIGN -> State 275
      optional_generic_binding -> State 71
      op_one_assign -> State 283
      binary_op_assign -> State 282

  State 219:
    Kernel Items:
      expression_or_implicit_struct: expression., *
    Reduce:
      * -> [expression_or_implicit_struct]
    Goto:
      (nil)

  State 220:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct.COMMA expression
      statement_body: expression_or_implicit_struct., *
    Reduce:
      * -> [statement_body]
    Goto:
      COMMA -> State 284

  State 221:
    Kernel Items:
      statement_body: jump_type.optional_jump_label optional_expression_or_implicit_struct
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 285
      optional_jump_label -> State 286

  State 222:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 223:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 287
      SEMICOLON -> State 288

  State 224:
    Kernel Items:
      statement_body: unsafe_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 225:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE.CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      CASE -> State 289

  State 226:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements RPAREN., $
    Reduce:
      $ -> [package_def]
    Goto:
      (nil)

  State 227:
    Kernel Items:
      package_statements: package_statements package_statement., *
    Reduce:
      * -> [package_statements]
    Goto:
      (nil)

  State 228:
    Kernel Items:
      package_statement: package_statement_body.NEWLINES
      package_statement: package_statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 290
      SEMICOLON -> State 291

  State 229:
    Kernel Items:
      package_statement_body: unsafe_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 230:
    Kernel Items:
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      (nil)

  State 231:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA.generic_parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 165
      generic_parameter_def -> State 292

  State 232:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 233:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 293
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 234:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 294
      OR -> State 295

  State 235:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 296

  State 236:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 246

  State 237:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 297
      OR -> State 298

  State 238:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 240
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      DOTDOTDOT -> State 239
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 244
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      parameter_decl -> State 242
      parameter_decls -> State 243
      optional_parameter_decls -> State 299
      func_type -> State 120

  State 239:
    Kernel Items:
      parameter_decl: DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 300
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 240:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      DOLLAR_LBRACKET -> State 67
      DOTDOTDOT -> State 301
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      optional_generic_binding -> State 176
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 302
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 241:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 303

  State 242:
    Kernel Items:
      parameter_decls: parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 243:
    Kernel Items:
      parameter_decls: parameter_decls.COMMA parameter_decl
      optional_parameter_decls: parameter_decls., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      COMMA -> State 304

  State 244:
    Kernel Items:
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 245:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 130
      field_def -> State 236
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      enum_value_def -> State 305
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 246:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 306

  State 247:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 130
      field_def -> State 236
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      enum_value_def -> State 307
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 248:
    Kernel Items:
      implicit_enum_def: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_def]
    Goto:
      (nil)

  State 249:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 130
      field_def -> State 308
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 250:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_def]
    Goto:
      (nil)

  State 251:
    Kernel Items:
      trait_algebra_type: trait_algebra_type ADD traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 252:
    Kernel Items:
      trait_algebra_type: trait_algebra_type MUL traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 253:
    Kernel Items:
      trait_algebra_type: trait_algebra_type SUB traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 254:
    Kernel Items:
      trait_properties: trait_properties COMMA trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 255:
    Kernel Items:
      trait_properties: trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 256:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      (nil)

  State 257:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 131
      parameter_def -> State 135
      parameter_defs -> State 136
      optional_parameter_defs -> State 309

  State 258:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 95
      block_body -> State 310

  State 259:
    Kernel Items:
      return_type: value_type., $
    Reduce:
      $ -> [return_type]
    Goto:
      (nil)

  State 260:
    Kernel Items:
      parameter_defs: parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 261:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 262:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 263:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
    Reduce:
      * -> [generic_arguments]
    Goto:
      (nil)

  State 264:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 265:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE.block_body
      if_expr: IF sequence_expr block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 94
      LBRACE -> State 95
      block_body -> State 311
      if_expr -> State 312

  State 266:
    Kernel Items:
      call_expr: access_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET argument RBRACKET
    Reduce:
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 69
      DOT -> State 68
      DOLLAR_LBRACKET -> State 67
      optional_generic_binding -> State 71

  State 267:
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

  State 268:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 313

  State 269:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 270:
    Kernel Items:
      op_one_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 271:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 272:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 273:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 274:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 275:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 276:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 277:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 280:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 281:
    Kernel Items:
      op_one_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [op_one_assign]
    Goto:
      (nil)

  State 282:
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
      expression -> State 314
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 283:
    Kernel Items:
      statement_body: access_expr op_one_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 284:
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
      expression -> State 315
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 285:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 286:
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
      optional_expression_or_implicit_struct -> State 317
      expression_or_implicit_struct -> State 316
      optional_label_decl -> State 46
      block_expr -> State 40
      expression -> State 219
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 287:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 288:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 289:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE CASE.DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 318

  State 290:
    Kernel Items:
      package_statement: package_statement_body NEWLINES., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 291:
    Kernel Items:
      package_statement: package_statement_body SEMICOLON., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 292:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 293:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      (nil)

  State 294:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 130
      field_def -> State 236
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      enum_value_def -> State 319
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 295:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 130
      field_def -> State 236
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      enum_value_def -> State 320
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 296:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_def]
    Goto:
      (nil)

  State 297:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 130
      field_def -> State 236
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      enum_value_def -> State 321
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 298:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 130
      field_def -> State 236
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      enum_value_def -> State 322
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 299:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 323

  State 300:
    Kernel Items:
      parameter_decl: DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 301:
    Kernel Items:
      parameter_decl: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 324
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      func_type -> State 120

  State 302:
    Kernel Items:
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 303:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 259
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      return_type -> State 325
      func_type -> State 120

  State 304:
    Kernel Items:
      parameter_decls: parameter_decls COMMA.parameter_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 240
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      DOTDOTDOT -> State 239
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 244
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      parameter_decl -> State 326
      func_type -> State 120

  State 305:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 306:
    Kernel Items:
      enum_value_def: field_def ASSIGN DEFAULT., *
    Reduce:
      * -> [enum_value_def]
    Goto:
      (nil)

  State 307:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 308:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [implicit_field_defs]
    Goto:
      (nil)

  State 309:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 327

  State 310:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 311:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 312:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 313:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 328

  State 314:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 315:
    Kernel Items:
      expression_or_implicit_struct: expression_or_implicit_struct COMMA expression., *
    Reduce:
      * -> [expression_or_implicit_struct]
    Goto:
      (nil)

  State 316:
    Kernel Items:
      optional_expression_or_implicit_struct: expression_or_implicit_struct., *
      expression_or_implicit_struct: expression_or_implicit_struct.COMMA expression
    Reduce:
      * -> [optional_expression_or_implicit_struct]
    Goto:
      COMMA -> State 284

  State 317:
    Kernel Items:
      statement_body: jump_type optional_jump_label optional_expression_or_implicit_struct., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 318:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE CASE DEFAULT.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 329

  State 319:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 320:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
      explicit_enum_value_defs: enum_value_def OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 321:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 322:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 323:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 259
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      return_type -> State 330
      func_type -> State 120

  State 324:
    Kernel Items:
      parameter_decl: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 325:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN return_type., $
    Reduce:
      $ -> [func_type]
    Goto:
      (nil)

  State 326:
    Kernel Items:
      parameter_decls: parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 327:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 146
      STRUCT -> State 31
      ENUM -> State 110
      TRAIT -> State 15
      FUNC -> State 142
      LPAREN -> State 113
      QUESTION -> State 114
      TILDE_TILDE -> State 115
      BIT_NEG -> State 109
      BIT_AND -> State 108
      atom_type -> State 116
      traitable_type -> State 129
      trait_algebra_type -> State 125
      value_type -> State 259
      implicit_struct_def -> State 122
      explicit_struct_def -> State 118
      implicit_enum_def -> State 121
      explicit_enum_def -> State 117
      trait_def -> State 126
      return_type -> State 331
      func_type -> State 120

  State 328:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 332

  State 329:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE CASE DEFAULT RBRACE., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 330:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 331:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 95
      block_body -> State 333

  State 332:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 333:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

Number of states: 333
Number of shift actions: 1849
Number of reduce actions: 271
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
