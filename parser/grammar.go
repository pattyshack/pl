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
	DoToken              = SymbolId(272)
	InToken              = SymbolId(273)
	ReturnToken          = SymbolId(274)
	BreakToken           = SymbolId(275)
	ContinueToken        = SymbolId(276)
	PackageToken         = SymbolId(277)
	UnsafeToken          = SymbolId(278)
	TypeToken            = SymbolId(279)
	ImplementsToken      = SymbolId(280)
	StructToken          = SymbolId(281)
	EnumToken            = SymbolId(282)
	TraitToken           = SymbolId(283)
	FuncToken            = SymbolId(284)
	AsyncToken           = SymbolId(285)
	DeferToken           = SymbolId(286)
	LabelDeclToken       = SymbolId(287)
	JumpLabelToken       = SymbolId(288)
	LbraceToken          = SymbolId(289)
	RbraceToken          = SymbolId(290)
	LparenToken          = SymbolId(291)
	RparenToken          = SymbolId(292)
	LbracketToken        = SymbolId(293)
	RbracketToken        = SymbolId(294)
	DotToken             = SymbolId(295)
	CommaToken           = SymbolId(296)
	QuestionToken        = SymbolId(297)
	SemicolonToken       = SymbolId(298)
	ColonToken           = SymbolId(299)
	DollarLbracketToken  = SymbolId(300)
	DotdotdotToken       = SymbolId(301)
	TildeTildeToken      = SymbolId(302)
	AssignToken          = SymbolId(303)
	AddAssignToken       = SymbolId(304)
	SubAssignToken       = SymbolId(305)
	MulAssignToken       = SymbolId(306)
	DivAssignToken       = SymbolId(307)
	ModAssignToken       = SymbolId(308)
	AddOneAssignToken    = SymbolId(309)
	SubOneAssignToken    = SymbolId(310)
	BitNegAssignToken    = SymbolId(311)
	BitAndAssignToken    = SymbolId(312)
	BitOrAssignToken     = SymbolId(313)
	BitXorAssignToken    = SymbolId(314)
	BitLshiftAssignToken = SymbolId(315)
	BitRshiftAssignToken = SymbolId(316)
	NotToken             = SymbolId(317)
	AndToken             = SymbolId(318)
	OrToken              = SymbolId(319)
	AddToken             = SymbolId(320)
	SubToken             = SymbolId(321)
	MulToken             = SymbolId(322)
	DivToken             = SymbolId(323)
	ModToken             = SymbolId(324)
	BitNegToken          = SymbolId(325)
	BitAndToken          = SymbolId(326)
	BitXorToken          = SymbolId(327)
	BitOrToken           = SymbolId(328)
	BitLshiftToken       = SymbolId(329)
	BitRshiftToken       = SymbolId(330)
	EqualToken           = SymbolId(331)
	NotEqualToken        = SymbolId(332)
	LessToken            = SymbolId(333)
	LessOrEqualToken     = SymbolId(334)
	GreaterToken         = SymbolId(335)
	GreaterOrEqualToken  = SymbolId(336)
	LexErrorToken        = SymbolId(337)
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
	// 50:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 51:2: expression -> switch_expr: ...
	SwitchExprToExpression(OptionalLabelDecl_ *GenericSymbol, SwitchExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 52:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ *GenericSymbol, LoopExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 53:2: expression -> sequence_expr: ...
	SequenceExprToExpression(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 56:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 57:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (*GenericSymbol, error)

	// 66:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 67:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, BlockBody_2 *GenericSymbol) (*GenericSymbol, error)

	// 68:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 75:15: switch_expr -> ...
	ToSwitchExpr(Switch_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Lbrace_ *GenericSymbol, Case_ *GenericSymbol, Default_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 89:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 90:2: loop_expr -> do_while: ...
	DoWhileToLoopExpr(Do_ *GenericSymbol, BlockBody_ *GenericSymbol, For_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 91:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 92:2: loop_expr -> iterator: ...
	IteratorToLoopExpr(For_ *GenericSymbol, In_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 93:2: loop_expr -> for: ...
	ForToLoopExpr(For_ *GenericSymbol, Semicolon_ *GenericSymbol, OptionalSequenceExpr_ *GenericSymbol, Semicolon_2 *GenericSymbol, OptionalSequenceExpr_2 *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 96:2: optional_sequence_expr -> sequence_expr: ...
	SequenceExprToOptionalSequenceExpr(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 97:2: optional_sequence_expr -> nil: ...
	NilToOptionalSequenceExpr() (*GenericSymbol, error)

	// 107:17: sequence_expr -> ...
	ToSequenceExpr(OrExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 109:14: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 111:14: block_body -> ...
	ToBlockBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 114:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 115:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 118:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 119:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 122:2: statement_body -> unsafe_statement: ...
	UnsafeStatementToStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 124:2: statement_body -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToStatementBody(Expressions_ *GenericSymbol) (*GenericSymbol, error)

	// 127:2: statement_body -> async: ...
	AsyncToStatementBody(Async_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 136:2: statement_body -> defer: ...
	DeferToStatementBody(Defer_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 140:2: statement_body -> jump_statement: ...
	JumpStatementToStatementBody(JumpStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 152:2: statement_body -> unary_op_assign_statement: ...
	UnaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, UnaryOpAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 153:2: statement_body -> binary_op_assign_statement: ...
	BinaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 160:2: unary_op_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToUnaryOpAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 161:2: unary_op_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToUnaryOpAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 164:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 165:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 166:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 167:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 168:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 169:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 170:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 171:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 172:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 173:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 174:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 182:20: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ *GenericSymbol, Less_ *GenericSymbol, Identifier_ *GenericSymbol, Greater_ *GenericSymbol, StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 189:2: jump_statement -> ...
	ToJumpStatement(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 192:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 193:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 194:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 197:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 198:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 201:2: expressions -> expression: ...
	ExpressionToExpressions(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 202:2: expressions -> add: ...
	AddToExpressions(Expressions_ *GenericSymbol, Comma_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 205:2: optional_expressions -> expressions: ...
	ExpressionsToOptionalExpressions(Expressions_ *GenericSymbol) (*GenericSymbol, error)

	// 206:2: optional_expressions -> nil: ...
	NilToOptionalExpressions() (*GenericSymbol, error)

	// 212:13: call_expr -> ...
	ToCallExpr(AccessExpr_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalArguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 215:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ *GenericSymbol, OptionalGenericArguments_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 216:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (*GenericSymbol, error)

	// 219:2: optional_generic_arguments -> generic_arguments: ...
	GenericArgumentsToOptionalGenericArguments(GenericArguments_ *GenericSymbol) (*GenericSymbol, error)

	// 220:2: optional_generic_arguments -> nil: ...
	NilToOptionalGenericArguments() (*GenericSymbol, error)

	// 224:2: generic_arguments -> value_type: ...
	ValueTypeToGenericArguments(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 225:2: generic_arguments -> add: ...
	AddToGenericArguments(GenericArguments_ *GenericSymbol, Comma_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 228:2: optional_arguments -> arguments: ...
	ArgumentsToOptionalArguments(Arguments_ *GenericSymbol) (*GenericSymbol, error)

	// 229:2: optional_arguments -> nil: ...
	NilToOptionalArguments() (*GenericSymbol, error)

	// 232:2: arguments -> argument: ...
	ArgumentToArguments(Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 233:2: arguments -> add: ...
	AddToArguments(Arguments_ *GenericSymbol, Comma_ *GenericSymbol, Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 236:2: argument -> positional: ...
	PositionalToArgument(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 237:2: argument -> named: ...
	NamedToArgument(Identifier_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 238:2: argument -> colon_expressions: ...
	ColonExpressionsToArgument(ColonExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 241:2: colon_expressions -> pair: ...
	PairToColonExpressions(OptionalExpression_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_2 *GenericSymbol) (*GenericSymbol, error)

	// 242:2: colon_expressions -> add: ...
	AddToColonExpressions(ColonExpressions_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_ *GenericSymbol) (*GenericSymbol, error)

	// 245:2: optional_expression -> expression: ...
	ExpressionToOptionalExpression(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 246:2: optional_expression -> nil: ...
	NilToOptionalExpression() (*GenericSymbol, error)

	// 256:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 257:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 258:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 259:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 260:2: atom_expr -> anonymous_struct_expr: ...
	AnonymousStructExprToAtomExpr(AnonymousStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 261:2: atom_expr -> LEX_ERROR: ...
	LexErrorToAtomExpr(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 264:2: literal -> TRUE: ...
	TrueToLiteral(True_ *GenericSymbol) (*GenericSymbol, error)

	// 265:2: literal -> FALSE: ...
	FalseToLiteral(False_ *GenericSymbol) (*GenericSymbol, error)

	// 266:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 267:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 268:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 269:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 272:2: anonymous_struct_expr -> explicit: ...
	ExplicitToAnonymousStructExpr(ExplicitStructDef_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 273:2: anonymous_struct_expr -> implicit: ...
	ImplicitToAnonymousStructExpr(Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 276:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 277:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 278:2: access_expr -> call_expr: ...
	CallExprToAccessExpr(CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 279:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Argument_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 282:2: postfix_unary_expr -> access_expr: ...
	AccessExprToPostfixUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 283:2: postfix_unary_expr -> question: ...
	QuestionToPostfixUnaryExpr(AccessExpr_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 286:2: prefix_unary_op -> NOT: ...
	NotToPrefixUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 287:2: prefix_unary_op -> BIT_NEG: ...
	BitNegToPrefixUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 288:2: prefix_unary_op -> SUB: ...
	SubToPrefixUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 291:2: prefix_unary_op -> MUL: ...
	MulToPrefixUnaryOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 294:2: prefix_unary_op -> BIT_AND: ...
	BitAndToPrefixUnaryOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 297:2: prefix_unary_expr -> postfix_unary_expr: ...
	PostfixUnaryExprToPrefixUnaryExpr(PostfixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 298:2: prefix_unary_expr -> prefix_op: ...
	PrefixOpToPrefixUnaryExpr(PrefixUnaryOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 301:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 302:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 303:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 304:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 305:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 306:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 309:2: mul_expr -> prefix_unary_expr: ...
	PrefixUnaryExprToMulExpr(PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 310:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 313:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 314:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 315:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 316:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 319:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 320:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 323:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 324:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 325:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 326:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 327:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 328:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 331:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 332:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 335:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 336:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 339:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 340:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 347:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 348:2: atom_type -> explicit_struct_def: ...
	ExplicitStructDefToAtomType(ExplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 349:2: atom_type -> implicit_struct_def: ...
	ImplicitStructDefToAtomType(ImplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 350:2: atom_type -> explicit_enum_def: ...
	ExplicitEnumDefToAtomType(ExplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 351:2: atom_type -> implicit_enum_def: ...
	ImplicitEnumDefToAtomType(ImplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 352:2: atom_type -> trait_def: ...
	TraitDefToAtomType(TraitDef_ *GenericSymbol) (*GenericSymbol, error)

	// 355:2: traitable_type -> atom_type: ...
	AtomTypeToTraitableType(AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 356:2: traitable_type -> public_methods_trait: ...
	PublicMethodsTraitToTraitableType(BitNeg_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 357:2: traitable_type -> public_trait: ...
	PublicTraitToTraitableType(TildeTilde_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 361:2: trait_algebra_type -> traitable_type: ...
	TraitableTypeToTraitAlgebraType(TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 362:2: trait_algebra_type -> intersect: ...
	IntersectToTraitAlgebraType(TraitAlgebraType_ *GenericSymbol, Mul_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 363:2: trait_algebra_type -> union: ...
	UnionToTraitAlgebraType(TraitAlgebraType_ *GenericSymbol, Add_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 364:2: trait_algebra_type -> difference: ...
	DifferenceToTraitAlgebraType(TraitAlgebraType_ *GenericSymbol, Sub_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 367:2: value_type -> inferred: ...
	InferredToValueType(Question_ *GenericSymbol) (*GenericSymbol, error)

	// 368:2: value_type -> trait_algebra_type: ...
	TraitAlgebraTypeToValueType(TraitAlgebraType_ *GenericSymbol) (*GenericSymbol, error)

	// 369:2: value_type -> reference: ...
	ReferenceToValueType(BitAnd_ *GenericSymbol, TraitAlgebraType_ *GenericSymbol) (*GenericSymbol, error)

	// 370:2: value_type -> func_type: ...
	FuncTypeToValueType(FuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 373:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 374:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol, Implements_ *GenericSymbol, ValueType_2 *GenericSymbol) (*GenericSymbol, error)

	// 375:2: type_def -> alias: ...
	AliasToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, Equal_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 383:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 384:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 387:2: generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToGenericParameterDefs(GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 388:2: generic_parameter_defs -> add: ...
	AddToGenericParameterDefs(GenericParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 391:2: optional_generic_parameter_defs -> generic_parameter_defs: ...
	GenericParameterDefsToOptionalGenericParameterDefs(GenericParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 392:2: optional_generic_parameter_defs -> nil: ...
	NilToOptionalGenericParameterDefs() (*GenericSymbol, error)

	// 395:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ *GenericSymbol, OptionalGenericParameterDefs_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 396:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (*GenericSymbol, error)

	// 403:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 404:2: field_def -> implicit: ...
	ImplicitToFieldDef(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 405:2: field_def -> unsafe_statement: ...
	UnsafeStatementToFieldDef(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 408:2: implicit_field_defs -> field_def: ...
	FieldDefToImplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 409:2: implicit_field_defs -> add: ...
	AddToImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 412:2: optional_implicit_field_defs -> implicit_field_defs: ...
	ImplicitFieldDefsToOptionalImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 413:2: optional_implicit_field_defs -> nil: ...
	NilToOptionalImplicitFieldDefs() (*GenericSymbol, error)

	// 415:23: implicit_struct_def -> ...
	ToImplicitStructDef(Lparen_ *GenericSymbol, OptionalImplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 418:2: explicit_field_defs -> field_def: ...
	FieldDefToExplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 419:2: explicit_field_defs -> implicit: ...
	ImplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Newlines_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 420:2: explicit_field_defs -> explicit: ...
	ExplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 423:2: optional_explicit_field_defs -> explicit_field_defs: ...
	ExplicitFieldDefsToOptionalExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 424:2: optional_explicit_field_defs -> nil: ...
	NilToOptionalExplicitFieldDefs() (*GenericSymbol, error)

	// 426:23: explicit_struct_def -> ...
	ToExplicitStructDef(Struct_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalExplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 436:2: enum_value_def -> field_def: ...
	FieldDefToEnumValueDef(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 437:2: enum_value_def -> default: ...
	DefaultToEnumValueDef(FieldDef_ *GenericSymbol, Assign_ *GenericSymbol, Default_ *GenericSymbol) (*GenericSymbol, error)

	// 440:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 441:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 443:21: implicit_enum_def -> ...
	ToImplicitEnumDef(Lparen_ *GenericSymbol, ImplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 446:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 447:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 448:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 449:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 451:21: explicit_enum_def -> ...
	ToExplicitEnumDef(Enum_ *GenericSymbol, Lparen_ *GenericSymbol, ExplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 458:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 459:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ *GenericSymbol) (*GenericSymbol, error)

	// 462:2: trait_properties -> trait_property: ...
	TraitPropertyToTraitProperties(TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 463:2: trait_properties -> implicit: ...
	ImplicitToTraitProperties(TraitProperties_ *GenericSymbol, Newlines_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 464:2: trait_properties -> explicit: ...
	ExplicitToTraitProperties(TraitProperties_ *GenericSymbol, Comma_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 467:2: optional_trait_properties -> trait_properties: ...
	TraitPropertiesToOptionalTraitProperties(TraitProperties_ *GenericSymbol) (*GenericSymbol, error)

	// 468:2: optional_trait_properties -> nil: ...
	NilToOptionalTraitProperties() (*GenericSymbol, error)

	// 470:13: trait_def -> ...
	ToTraitDef(Trait_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalTraitProperties_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 478:2: return_type -> value_type: ...
	ValueTypeToReturnType(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 479:2: return_type -> nil: ...
	NilToReturnType() (*GenericSymbol, error)

	// 482:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 483:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 484:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 485:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 488:2: parameter_decls -> parameter_decl: ...
	ParameterDeclToParameterDecls(ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 489:2: parameter_decls -> add: ...
	AddToParameterDecls(ParameterDecls_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 492:2: optional_parameter_decls -> parameter_decls: ...
	ParameterDeclsToOptionalParameterDecls(ParameterDecls_ *GenericSymbol) (*GenericSymbol, error)

	// 493:2: optional_parameter_decls -> nil: ...
	NilToOptionalParameterDecls() (*GenericSymbol, error)

	// 495:13: func_type -> ...
	ToFuncType(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 503:20: method_signature -> ...
	ToMethodSignature(Func_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 506:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 507:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 510:2: parameter_defs -> parameter_def: ...
	ParameterDefToParameterDefs(ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 511:2: parameter_defs -> add: ...
	AddToParameterDefs(ParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 514:2: optional_parameter_defs -> parameter_defs: ...
	ParameterDefsToOptionalParameterDefs(ParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 515:2: optional_parameter_defs -> nil: ...
	NilToOptionalParameterDefs() (*GenericSymbol, error)

	// 518:2: optional_receiver -> receiver: ...
	ReceiverToOptionalReceiver(Lparen_ *GenericSymbol, ParameterDef_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 519:2: optional_receiver -> nil: ...
	NilToOptionalReceiver() (*GenericSymbol, error)

	// 522:2: named_func_def -> ...
	ToNamedFuncDef(Func_ *GenericSymbol, OptionalReceiver_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 525:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 532:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 533:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, PackageStatements_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 535:26: package_statement_body -> ...
	ToPackageStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 538:2: package_statement -> implicit: ...
	ImplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 539:2: package_statement -> explicit: ...
	ExplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 542:2: package_statements -> empty_list: ...
	EmptyListToPackageStatements() (*GenericSymbol, error)

	// 543:2: package_statements -> add: ...
	AddToPackageStatements(PackageStatements_ *GenericSymbol, PackageStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 547:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 548:2: lex_internal_tokens -> COMMENT: ...
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
	case DoToken:
		return "DO"
	case InToken:
		return "IN"
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
	case ImplementsToken:
		return "IMPLEMENTS"
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
	case DeferToken:
		return "DEFER"
	case LabelDeclToken:
		return "LABEL_DECL"
	case JumpLabelToken:
		return "JUMP_LABEL"
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
	case ExpressionType:
		return "expression"
	case OptionalLabelDeclType:
		return "optional_label_decl"
	case IfExprType:
		return "if_expr"
	case SwitchExprType:
		return "switch_expr"
	case LoopExprType:
		return "loop_expr"
	case OptionalSequenceExprType:
		return "optional_sequence_expr"
	case SequenceExprType:
		return "sequence_expr"
	case BlockExprType:
		return "block_expr"
	case BlockBodyType:
		return "block_body"
	case StatementsType:
		return "statements"
	case StatementType:
		return "statement"
	case StatementBodyType:
		return "statement_body"
	case UnaryOpAssignType:
		return "unary_op_assign"
	case BinaryOpAssignType:
		return "binary_op_assign"
	case UnsafeStatementType:
		return "unsafe_statement"
	case JumpStatementType:
		return "jump_statement"
	case JumpTypeType:
		return "jump_type"
	case OptionalJumpLabelType:
		return "optional_jump_label"
	case ExpressionsType:
		return "expressions"
	case OptionalExpressionsType:
		return "optional_expressions"
	case CallExprType:
		return "call_expr"
	case OptionalGenericBindingType:
		return "optional_generic_binding"
	case OptionalGenericArgumentsType:
		return "optional_generic_arguments"
	case GenericArgumentsType:
		return "generic_arguments"
	case OptionalArgumentsType:
		return "optional_arguments"
	case ArgumentsType:
		return "arguments"
	case ArgumentType:
		return "argument"
	case ColonExpressionsType:
		return "colon_expressions"
	case OptionalExpressionType:
		return "optional_expression"
	case AtomExprType:
		return "atom_expr"
	case LiteralType:
		return "literal"
	case AnonymousStructExprType:
		return "anonymous_struct_expr"
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

	ExpressionType                   = SymbolId(338)
	OptionalLabelDeclType            = SymbolId(339)
	IfExprType                       = SymbolId(340)
	SwitchExprType                   = SymbolId(341)
	LoopExprType                     = SymbolId(342)
	OptionalSequenceExprType         = SymbolId(343)
	SequenceExprType                 = SymbolId(344)
	BlockExprType                    = SymbolId(345)
	BlockBodyType                    = SymbolId(346)
	StatementsType                   = SymbolId(347)
	StatementType                    = SymbolId(348)
	StatementBodyType                = SymbolId(349)
	UnaryOpAssignType                = SymbolId(350)
	BinaryOpAssignType               = SymbolId(351)
	UnsafeStatementType              = SymbolId(352)
	JumpStatementType                = SymbolId(353)
	JumpTypeType                     = SymbolId(354)
	OptionalJumpLabelType            = SymbolId(355)
	ExpressionsType                  = SymbolId(356)
	OptionalExpressionsType          = SymbolId(357)
	CallExprType                     = SymbolId(358)
	OptionalGenericBindingType       = SymbolId(359)
	OptionalGenericArgumentsType     = SymbolId(360)
	GenericArgumentsType             = SymbolId(361)
	OptionalArgumentsType            = SymbolId(362)
	ArgumentsType                    = SymbolId(363)
	ArgumentType                     = SymbolId(364)
	ColonExpressionsType             = SymbolId(365)
	OptionalExpressionType           = SymbolId(366)
	AtomExprType                     = SymbolId(367)
	LiteralType                      = SymbolId(368)
	AnonymousStructExprType          = SymbolId(369)
	AccessExprType                   = SymbolId(370)
	PostfixUnaryExprType             = SymbolId(371)
	PrefixUnaryOpType                = SymbolId(372)
	PrefixUnaryExprType              = SymbolId(373)
	MulOpType                        = SymbolId(374)
	MulExprType                      = SymbolId(375)
	AddOpType                        = SymbolId(376)
	AddExprType                      = SymbolId(377)
	CmpOpType                        = SymbolId(378)
	CmpExprType                      = SymbolId(379)
	AndExprType                      = SymbolId(380)
	OrExprType                       = SymbolId(381)
	AtomTypeType                     = SymbolId(382)
	TraitableTypeType                = SymbolId(383)
	TraitAlgebraTypeType             = SymbolId(384)
	ValueTypeType                    = SymbolId(385)
	TypeDefType                      = SymbolId(386)
	GenericParameterDefType          = SymbolId(387)
	GenericParameterDefsType         = SymbolId(388)
	OptionalGenericParameterDefsType = SymbolId(389)
	OptionalGenericParametersType    = SymbolId(390)
	FieldDefType                     = SymbolId(391)
	ImplicitFieldDefsType            = SymbolId(392)
	OptionalImplicitFieldDefsType    = SymbolId(393)
	ImplicitStructDefType            = SymbolId(394)
	ExplicitFieldDefsType            = SymbolId(395)
	OptionalExplicitFieldDefsType    = SymbolId(396)
	ExplicitStructDefType            = SymbolId(397)
	EnumValueDefType                 = SymbolId(398)
	ImplicitEnumValueDefsType        = SymbolId(399)
	ImplicitEnumDefType              = SymbolId(400)
	ExplicitEnumValueDefsType        = SymbolId(401)
	ExplicitEnumDefType              = SymbolId(402)
	TraitPropertyType                = SymbolId(403)
	TraitPropertiesType              = SymbolId(404)
	OptionalTraitPropertiesType      = SymbolId(405)
	TraitDefType                     = SymbolId(406)
	ReturnTypeType                   = SymbolId(407)
	ParameterDeclType                = SymbolId(408)
	ParameterDeclsType               = SymbolId(409)
	OptionalParameterDeclsType       = SymbolId(410)
	FuncTypeType                     = SymbolId(411)
	MethodSignatureType              = SymbolId(412)
	ParameterDefType                 = SymbolId(413)
	ParameterDefsType                = SymbolId(414)
	OptionalParameterDefsType        = SymbolId(415)
	OptionalReceiverType             = SymbolId(416)
	NamedFuncDefType                 = SymbolId(417)
	AnonymousFuncExprType            = SymbolId(418)
	PackageDefType                   = SymbolId(419)
	PackageStatementBodyType         = SymbolId(420)
	PackageStatementType             = SymbolId(421)
	PackageStatementsType            = SymbolId(422)
	LexInternalTokensType            = SymbolId(423)
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
	_ReduceIfExprToExpression                                 = _ReduceType(1)
	_ReduceSwitchExprToExpression                             = _ReduceType(2)
	_ReduceLoopExprToExpression                               = _ReduceType(3)
	_ReduceSequenceExprToExpression                           = _ReduceType(4)
	_ReduceLabelDeclToOptionalLabelDecl                       = _ReduceType(5)
	_ReduceUnlabelledToOptionalLabelDecl                      = _ReduceType(6)
	_ReduceNoElseToIfExpr                                     = _ReduceType(7)
	_ReduceIfElseToIfExpr                                     = _ReduceType(8)
	_ReduceMultiIfElseToIfExpr                                = _ReduceType(9)
	_ReduceToSwitchExpr                                       = _ReduceType(10)
	_ReduceInfiniteToLoopExpr                                 = _ReduceType(11)
	_ReduceDoWhileToLoopExpr                                  = _ReduceType(12)
	_ReduceWhileToLoopExpr                                    = _ReduceType(13)
	_ReduceIteratorToLoopExpr                                 = _ReduceType(14)
	_ReduceForToLoopExpr                                      = _ReduceType(15)
	_ReduceSequenceExprToOptionalSequenceExpr                 = _ReduceType(16)
	_ReduceNilToOptionalSequenceExpr                          = _ReduceType(17)
	_ReduceToSequenceExpr                                     = _ReduceType(18)
	_ReduceToBlockExpr                                        = _ReduceType(19)
	_ReduceToBlockBody                                        = _ReduceType(20)
	_ReduceEmptyListToStatements                              = _ReduceType(21)
	_ReduceAddToStatements                                    = _ReduceType(22)
	_ReduceImplicitToStatement                                = _ReduceType(23)
	_ReduceExplicitToStatement                                = _ReduceType(24)
	_ReduceUnsafeStatementToStatementBody                     = _ReduceType(25)
	_ReduceExpressionOrImplicitStructToStatementBody          = _ReduceType(26)
	_ReduceAsyncToStatementBody                               = _ReduceType(27)
	_ReduceDeferToStatementBody                               = _ReduceType(28)
	_ReduceJumpStatementToStatementBody                       = _ReduceType(29)
	_ReduceUnaryOpAssignStatementToStatementBody              = _ReduceType(30)
	_ReduceBinaryOpAssignStatementToStatementBody             = _ReduceType(31)
	_ReduceAddOneAssignToUnaryOpAssign                        = _ReduceType(32)
	_ReduceSubOneAssignToUnaryOpAssign                        = _ReduceType(33)
	_ReduceAddAssignToBinaryOpAssign                          = _ReduceType(34)
	_ReduceSubAssignToBinaryOpAssign                          = _ReduceType(35)
	_ReduceMulAssignToBinaryOpAssign                          = _ReduceType(36)
	_ReduceDivAssignToBinaryOpAssign                          = _ReduceType(37)
	_ReduceModAssignToBinaryOpAssign                          = _ReduceType(38)
	_ReduceBitNegAssignToBinaryOpAssign                       = _ReduceType(39)
	_ReduceBitAndAssignToBinaryOpAssign                       = _ReduceType(40)
	_ReduceBitOrAssignToBinaryOpAssign                        = _ReduceType(41)
	_ReduceBitXorAssignToBinaryOpAssign                       = _ReduceType(42)
	_ReduceBitLshiftAssignToBinaryOpAssign                    = _ReduceType(43)
	_ReduceBitRshiftAssignToBinaryOpAssign                    = _ReduceType(44)
	_ReduceToUnsafeStatement                                  = _ReduceType(45)
	_ReduceToJumpStatement                                    = _ReduceType(46)
	_ReduceReturnToJumpType                                   = _ReduceType(47)
	_ReduceBreakToJumpType                                    = _ReduceType(48)
	_ReduceContinueToJumpType                                 = _ReduceType(49)
	_ReduceJumpLabelToOptionalJumpLabel                       = _ReduceType(50)
	_ReduceUnlabelledToOptionalJumpLabel                      = _ReduceType(51)
	_ReduceExpressionToExpressions                            = _ReduceType(52)
	_ReduceAddToExpressions                                   = _ReduceType(53)
	_ReduceExpressionsToOptionalExpressions                   = _ReduceType(54)
	_ReduceNilToOptionalExpressions                           = _ReduceType(55)
	_ReduceToCallExpr                                         = _ReduceType(56)
	_ReduceBindingToOptionalGenericBinding                    = _ReduceType(57)
	_ReduceNilToOptionalGenericBinding                        = _ReduceType(58)
	_ReduceGenericArgumentsToOptionalGenericArguments         = _ReduceType(59)
	_ReduceNilToOptionalGenericArguments                      = _ReduceType(60)
	_ReduceValueTypeToGenericArguments                        = _ReduceType(61)
	_ReduceAddToGenericArguments                              = _ReduceType(62)
	_ReduceArgumentsToOptionalArguments                       = _ReduceType(63)
	_ReduceNilToOptionalArguments                             = _ReduceType(64)
	_ReduceArgumentToArguments                                = _ReduceType(65)
	_ReduceAddToArguments                                     = _ReduceType(66)
	_ReducePositionalToArgument                               = _ReduceType(67)
	_ReduceNamedToArgument                                    = _ReduceType(68)
	_ReduceColonExpressionsToArgument                         = _ReduceType(69)
	_ReducePairToColonExpressions                             = _ReduceType(70)
	_ReduceAddToColonExpressions                              = _ReduceType(71)
	_ReduceExpressionToOptionalExpression                     = _ReduceType(72)
	_ReduceNilToOptionalExpression                            = _ReduceType(73)
	_ReduceLiteralToAtomExpr                                  = _ReduceType(74)
	_ReduceIdentifierToAtomExpr                               = _ReduceType(75)
	_ReduceBlockExprToAtomExpr                                = _ReduceType(76)
	_ReduceAnonymousFuncExprToAtomExpr                        = _ReduceType(77)
	_ReduceAnonymousStructExprToAtomExpr                      = _ReduceType(78)
	_ReduceLexErrorToAtomExpr                                 = _ReduceType(79)
	_ReduceTrueToLiteral                                      = _ReduceType(80)
	_ReduceFalseToLiteral                                     = _ReduceType(81)
	_ReduceIntegerLiteralToLiteral                            = _ReduceType(82)
	_ReduceFloatLiteralToLiteral                              = _ReduceType(83)
	_ReduceRuneLiteralToLiteral                               = _ReduceType(84)
	_ReduceStringLiteralToLiteral                             = _ReduceType(85)
	_ReduceExplicitToAnonymousStructExpr                      = _ReduceType(86)
	_ReduceImplicitToAnonymousStructExpr                      = _ReduceType(87)
	_ReduceAtomExprToAccessExpr                               = _ReduceType(88)
	_ReduceAccessToAccessExpr                                 = _ReduceType(89)
	_ReduceCallExprToAccessExpr                               = _ReduceType(90)
	_ReduceIndexToAccessExpr                                  = _ReduceType(91)
	_ReduceAccessExprToPostfixUnaryExpr                       = _ReduceType(92)
	_ReduceQuestionToPostfixUnaryExpr                         = _ReduceType(93)
	_ReduceNotToPrefixUnaryOp                                 = _ReduceType(94)
	_ReduceBitNegToPrefixUnaryOp                              = _ReduceType(95)
	_ReduceSubToPrefixUnaryOp                                 = _ReduceType(96)
	_ReduceMulToPrefixUnaryOp                                 = _ReduceType(97)
	_ReduceBitAndToPrefixUnaryOp                              = _ReduceType(98)
	_ReducePostfixUnaryExprToPrefixUnaryExpr                  = _ReduceType(99)
	_ReducePrefixOpToPrefixUnaryExpr                          = _ReduceType(100)
	_ReduceMulToMulOp                                         = _ReduceType(101)
	_ReduceDivToMulOp                                         = _ReduceType(102)
	_ReduceModToMulOp                                         = _ReduceType(103)
	_ReduceBitAndToMulOp                                      = _ReduceType(104)
	_ReduceBitLshiftToMulOp                                   = _ReduceType(105)
	_ReduceBitRshiftToMulOp                                   = _ReduceType(106)
	_ReducePrefixUnaryExprToMulExpr                           = _ReduceType(107)
	_ReduceOpToMulExpr                                        = _ReduceType(108)
	_ReduceAddToAddOp                                         = _ReduceType(109)
	_ReduceSubToAddOp                                         = _ReduceType(110)
	_ReduceBitOrToAddOp                                       = _ReduceType(111)
	_ReduceBitXorToAddOp                                      = _ReduceType(112)
	_ReduceMulExprToAddExpr                                   = _ReduceType(113)
	_ReduceOpToAddExpr                                        = _ReduceType(114)
	_ReduceEqualToCmpOp                                       = _ReduceType(115)
	_ReduceNotEqualToCmpOp                                    = _ReduceType(116)
	_ReduceLessToCmpOp                                        = _ReduceType(117)
	_ReduceLessOrEqualToCmpOp                                 = _ReduceType(118)
	_ReduceGreaterToCmpOp                                     = _ReduceType(119)
	_ReduceGreaterOrEqualToCmpOp                              = _ReduceType(120)
	_ReduceAddExprToCmpExpr                                   = _ReduceType(121)
	_ReduceOpToCmpExpr                                        = _ReduceType(122)
	_ReduceCmpExprToAndExpr                                   = _ReduceType(123)
	_ReduceOpToAndExpr                                        = _ReduceType(124)
	_ReduceAndExprToOrExpr                                    = _ReduceType(125)
	_ReduceOpToOrExpr                                         = _ReduceType(126)
	_ReduceNamedToAtomType                                    = _ReduceType(127)
	_ReduceExplicitStructDefToAtomType                        = _ReduceType(128)
	_ReduceImplicitStructDefToAtomType                        = _ReduceType(129)
	_ReduceExplicitEnumDefToAtomType                          = _ReduceType(130)
	_ReduceImplicitEnumDefToAtomType                          = _ReduceType(131)
	_ReduceTraitDefToAtomType                                 = _ReduceType(132)
	_ReduceAtomTypeToTraitableType                            = _ReduceType(133)
	_ReducePublicMethodsTraitToTraitableType                  = _ReduceType(134)
	_ReducePublicTraitToTraitableType                         = _ReduceType(135)
	_ReduceTraitableTypeToTraitAlgebraType                    = _ReduceType(136)
	_ReduceIntersectToTraitAlgebraType                        = _ReduceType(137)
	_ReduceUnionToTraitAlgebraType                            = _ReduceType(138)
	_ReduceDifferenceToTraitAlgebraType                       = _ReduceType(139)
	_ReduceInferredToValueType                                = _ReduceType(140)
	_ReduceTraitAlgebraTypeToValueType                        = _ReduceType(141)
	_ReduceReferenceToValueType                               = _ReduceType(142)
	_ReduceFuncTypeToValueType                                = _ReduceType(143)
	_ReduceDefinitionToTypeDef                                = _ReduceType(144)
	_ReduceConstrainedDefToTypeDef                            = _ReduceType(145)
	_ReduceAliasToTypeDef                                     = _ReduceType(146)
	_ReduceUnconstrainedToGenericParameterDef                 = _ReduceType(147)
	_ReduceConstrainedToGenericParameterDef                   = _ReduceType(148)
	_ReduceGenericParameterDefToGenericParameterDefs          = _ReduceType(149)
	_ReduceAddToGenericParameterDefs                          = _ReduceType(150)
	_ReduceGenericParameterDefsToOptionalGenericParameterDefs = _ReduceType(151)
	_ReduceNilToOptionalGenericParameterDefs                  = _ReduceType(152)
	_ReduceGenericToOptionalGenericParameters                 = _ReduceType(153)
	_ReduceNilToOptionalGenericParameters                     = _ReduceType(154)
	_ReduceExplicitToFieldDef                                 = _ReduceType(155)
	_ReduceImplicitToFieldDef                                 = _ReduceType(156)
	_ReduceUnsafeStatementToFieldDef                          = _ReduceType(157)
	_ReduceFieldDefToImplicitFieldDefs                        = _ReduceType(158)
	_ReduceAddToImplicitFieldDefs                             = _ReduceType(159)
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefs       = _ReduceType(160)
	_ReduceNilToOptionalImplicitFieldDefs                     = _ReduceType(161)
	_ReduceToImplicitStructDef                                = _ReduceType(162)
	_ReduceFieldDefToExplicitFieldDefs                        = _ReduceType(163)
	_ReduceImplicitToExplicitFieldDefs                        = _ReduceType(164)
	_ReduceExplicitToExplicitFieldDefs                        = _ReduceType(165)
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefs       = _ReduceType(166)
	_ReduceNilToOptionalExplicitFieldDefs                     = _ReduceType(167)
	_ReduceToExplicitStructDef                                = _ReduceType(168)
	_ReduceFieldDefToEnumValueDef                             = _ReduceType(169)
	_ReduceDefaultToEnumValueDef                              = _ReduceType(170)
	_ReducePairToImplicitEnumValueDefs                        = _ReduceType(171)
	_ReduceAddToImplicitEnumValueDefs                         = _ReduceType(172)
	_ReduceToImplicitEnumDef                                  = _ReduceType(173)
	_ReduceExplicitPairToExplicitEnumValueDefs                = _ReduceType(174)
	_ReduceImplicitPairToExplicitEnumValueDefs                = _ReduceType(175)
	_ReduceExplicitAddToExplicitEnumValueDefs                 = _ReduceType(176)
	_ReduceImplicitAddToExplicitEnumValueDefs                 = _ReduceType(177)
	_ReduceToExplicitEnumDef                                  = _ReduceType(178)
	_ReduceFieldDefToTraitProperty                            = _ReduceType(179)
	_ReduceMethodSignatureToTraitProperty                     = _ReduceType(180)
	_ReduceTraitPropertyToTraitProperties                     = _ReduceType(181)
	_ReduceImplicitToTraitProperties                          = _ReduceType(182)
	_ReduceExplicitToTraitProperties                          = _ReduceType(183)
	_ReduceTraitPropertiesToOptionalTraitProperties           = _ReduceType(184)
	_ReduceNilToOptionalTraitProperties                       = _ReduceType(185)
	_ReduceToTraitDef                                         = _ReduceType(186)
	_ReduceValueTypeToReturnType                              = _ReduceType(187)
	_ReduceNilToReturnType                                    = _ReduceType(188)
	_ReduceArgToParameterDecl                                 = _ReduceType(189)
	_ReduceVarargToParameterDecl                              = _ReduceType(190)
	_ReduceUnamedToParameterDecl                              = _ReduceType(191)
	_ReduceUnnamedVarargToParameterDecl                       = _ReduceType(192)
	_ReduceParameterDeclToParameterDecls                      = _ReduceType(193)
	_ReduceAddToParameterDecls                                = _ReduceType(194)
	_ReduceParameterDeclsToOptionalParameterDecls             = _ReduceType(195)
	_ReduceNilToOptionalParameterDecls                        = _ReduceType(196)
	_ReduceToFuncType                                         = _ReduceType(197)
	_ReduceToMethodSignature                                  = _ReduceType(198)
	_ReduceArgToParameterDef                                  = _ReduceType(199)
	_ReduceVarargToParameterDef                               = _ReduceType(200)
	_ReduceParameterDefToParameterDefs                        = _ReduceType(201)
	_ReduceAddToParameterDefs                                 = _ReduceType(202)
	_ReduceParameterDefsToOptionalParameterDefs               = _ReduceType(203)
	_ReduceNilToOptionalParameterDefs                         = _ReduceType(204)
	_ReduceReceiverToOptionalReceiver                         = _ReduceType(205)
	_ReduceNilToOptionalReceiver                              = _ReduceType(206)
	_ReduceToNamedFuncDef                                     = _ReduceType(207)
	_ReduceToAnonymousFuncExpr                                = _ReduceType(208)
	_ReduceNoSpecToPackageDef                                 = _ReduceType(209)
	_ReduceWithSpecToPackageDef                               = _ReduceType(210)
	_ReduceToPackageStatementBody                             = _ReduceType(211)
	_ReduceImplicitToPackageStatement                         = _ReduceType(212)
	_ReduceExplicitToPackageStatement                         = _ReduceType(213)
	_ReduceEmptyListToPackageStatements                       = _ReduceType(214)
	_ReduceAddToPackageStatements                             = _ReduceType(215)
	_ReduceSpacesToLexInternalTokens                          = _ReduceType(216)
	_ReduceCommentToLexInternalTokens                         = _ReduceType(217)
)

func (i _ReduceType) String() string {
	switch i {
	case _ReduceIfExprToExpression:
		return "IfExprToExpression"
	case _ReduceSwitchExprToExpression:
		return "SwitchExprToExpression"
	case _ReduceLoopExprToExpression:
		return "LoopExprToExpression"
	case _ReduceSequenceExprToExpression:
		return "SequenceExprToExpression"
	case _ReduceLabelDeclToOptionalLabelDecl:
		return "LabelDeclToOptionalLabelDecl"
	case _ReduceUnlabelledToOptionalLabelDecl:
		return "UnlabelledToOptionalLabelDecl"
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
	case _ReduceDoWhileToLoopExpr:
		return "DoWhileToLoopExpr"
	case _ReduceWhileToLoopExpr:
		return "WhileToLoopExpr"
	case _ReduceIteratorToLoopExpr:
		return "IteratorToLoopExpr"
	case _ReduceForToLoopExpr:
		return "ForToLoopExpr"
	case _ReduceSequenceExprToOptionalSequenceExpr:
		return "SequenceExprToOptionalSequenceExpr"
	case _ReduceNilToOptionalSequenceExpr:
		return "NilToOptionalSequenceExpr"
	case _ReduceToSequenceExpr:
		return "ToSequenceExpr"
	case _ReduceToBlockExpr:
		return "ToBlockExpr"
	case _ReduceToBlockBody:
		return "ToBlockBody"
	case _ReduceEmptyListToStatements:
		return "EmptyListToStatements"
	case _ReduceAddToStatements:
		return "AddToStatements"
	case _ReduceImplicitToStatement:
		return "ImplicitToStatement"
	case _ReduceExplicitToStatement:
		return "ExplicitToStatement"
	case _ReduceUnsafeStatementToStatementBody:
		return "UnsafeStatementToStatementBody"
	case _ReduceExpressionOrImplicitStructToStatementBody:
		return "ExpressionOrImplicitStructToStatementBody"
	case _ReduceAsyncToStatementBody:
		return "AsyncToStatementBody"
	case _ReduceDeferToStatementBody:
		return "DeferToStatementBody"
	case _ReduceJumpStatementToStatementBody:
		return "JumpStatementToStatementBody"
	case _ReduceUnaryOpAssignStatementToStatementBody:
		return "UnaryOpAssignStatementToStatementBody"
	case _ReduceBinaryOpAssignStatementToStatementBody:
		return "BinaryOpAssignStatementToStatementBody"
	case _ReduceAddOneAssignToUnaryOpAssign:
		return "AddOneAssignToUnaryOpAssign"
	case _ReduceSubOneAssignToUnaryOpAssign:
		return "SubOneAssignToUnaryOpAssign"
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
	case _ReduceToUnsafeStatement:
		return "ToUnsafeStatement"
	case _ReduceToJumpStatement:
		return "ToJumpStatement"
	case _ReduceReturnToJumpType:
		return "ReturnToJumpType"
	case _ReduceBreakToJumpType:
		return "BreakToJumpType"
	case _ReduceContinueToJumpType:
		return "ContinueToJumpType"
	case _ReduceJumpLabelToOptionalJumpLabel:
		return "JumpLabelToOptionalJumpLabel"
	case _ReduceUnlabelledToOptionalJumpLabel:
		return "UnlabelledToOptionalJumpLabel"
	case _ReduceExpressionToExpressions:
		return "ExpressionToExpressions"
	case _ReduceAddToExpressions:
		return "AddToExpressions"
	case _ReduceExpressionsToOptionalExpressions:
		return "ExpressionsToOptionalExpressions"
	case _ReduceNilToOptionalExpressions:
		return "NilToOptionalExpressions"
	case _ReduceToCallExpr:
		return "ToCallExpr"
	case _ReduceBindingToOptionalGenericBinding:
		return "BindingToOptionalGenericBinding"
	case _ReduceNilToOptionalGenericBinding:
		return "NilToOptionalGenericBinding"
	case _ReduceGenericArgumentsToOptionalGenericArguments:
		return "GenericArgumentsToOptionalGenericArguments"
	case _ReduceNilToOptionalGenericArguments:
		return "NilToOptionalGenericArguments"
	case _ReduceValueTypeToGenericArguments:
		return "ValueTypeToGenericArguments"
	case _ReduceAddToGenericArguments:
		return "AddToGenericArguments"
	case _ReduceArgumentsToOptionalArguments:
		return "ArgumentsToOptionalArguments"
	case _ReduceNilToOptionalArguments:
		return "NilToOptionalArguments"
	case _ReduceArgumentToArguments:
		return "ArgumentToArguments"
	case _ReduceAddToArguments:
		return "AddToArguments"
	case _ReducePositionalToArgument:
		return "PositionalToArgument"
	case _ReduceNamedToArgument:
		return "NamedToArgument"
	case _ReduceColonExpressionsToArgument:
		return "ColonExpressionsToArgument"
	case _ReducePairToColonExpressions:
		return "PairToColonExpressions"
	case _ReduceAddToColonExpressions:
		return "AddToColonExpressions"
	case _ReduceExpressionToOptionalExpression:
		return "ExpressionToOptionalExpression"
	case _ReduceNilToOptionalExpression:
		return "NilToOptionalExpression"
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
	case _ReduceUnsafeStatementToFieldDef:
		return "UnsafeStatementToFieldDef"
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
	_State334 = _StateId(334)
	_State335 = _StateId(335)
	_State336 = _StateId(336)
	_State337 = _StateId(337)
	_State338 = _StateId(338)
	_State339 = _StateId(339)
	_State340 = _StateId(340)
	_State341 = _StateId(341)
	_State342 = _StateId(342)
	_State343 = _StateId(343)
	_State344 = _StateId(344)
	_State345 = _StateId(345)
	_State346 = _StateId(346)
	_State347 = _StateId(347)
	_State348 = _StateId(348)
	_State349 = _StateId(349)
	_State350 = _StateId(350)
	_State351 = _StateId(351)
	_State352 = _StateId(352)
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
	case _EndMarker, SpacesToken, NewlinesToken, CommentToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, PackageToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, DollarLbracketToken, DotdotdotToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, NotToken, AndToken, OrToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, LexErrorToken:
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
	case _ReduceSequenceExprToExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.SequenceExprToExpression(args[0].Generic_)
	case _ReduceLabelDeclToOptionalLabelDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalLabelDeclType
		symbol.Generic_, err = reducer.LabelDeclToOptionalLabelDecl(args[0].Generic_)
	case _ReduceUnlabelledToOptionalLabelDecl:
		symbol.SymbolId_ = OptionalLabelDeclType
		symbol.Generic_, err = reducer.UnlabelledToOptionalLabelDecl()
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
	case _ReduceDoWhileToLoopExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = LoopExprType
		symbol.Generic_, err = reducer.DoWhileToLoopExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceWhileToLoopExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = LoopExprType
		symbol.Generic_, err = reducer.WhileToLoopExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceIteratorToLoopExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = LoopExprType
		symbol.Generic_, err = reducer.IteratorToLoopExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceForToLoopExpr:
		args := stack[len(stack)-7:]
		stack = stack[:len(stack)-7]
		symbol.SymbolId_ = LoopExprType
		symbol.Generic_, err = reducer.ForToLoopExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_, args[6].Generic_)
	case _ReduceSequenceExprToOptionalSequenceExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalSequenceExprType
		symbol.Generic_, err = reducer.SequenceExprToOptionalSequenceExpr(args[0].Generic_)
	case _ReduceNilToOptionalSequenceExpr:
		symbol.SymbolId_ = OptionalSequenceExprType
		symbol.Generic_, err = reducer.NilToOptionalSequenceExpr()
	case _ReduceToSequenceExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SequenceExprType
		symbol.Generic_, err = reducer.ToSequenceExpr(args[0].Generic_)
	case _ReduceToBlockExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = BlockExprType
		symbol.Generic_, err = reducer.ToBlockExpr(args[0].Generic_, args[1].Generic_)
	case _ReduceToBlockBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BlockBodyType
		symbol.Generic_, err = reducer.ToBlockBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceEmptyListToStatements:
		symbol.SymbolId_ = StatementsType
		symbol.Generic_, err = reducer.EmptyListToStatements()
	case _ReduceAddToStatements:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementsType
		symbol.Generic_, err = reducer.AddToStatements(args[0].Generic_, args[1].Generic_)
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
	case _ReduceDeferToStatementBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.DeferToStatementBody(args[0].Generic_, args[1].Generic_)
	case _ReduceJumpStatementToStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.JumpStatementToStatementBody(args[0].Generic_)
	case _ReduceUnaryOpAssignStatementToStatementBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.UnaryOpAssignStatementToStatementBody(args[0].Generic_, args[1].Generic_)
	case _ReduceBinaryOpAssignStatementToStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.BinaryOpAssignStatementToStatementBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceAddOneAssignToUnaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpAssignType
		symbol.Generic_, err = reducer.AddOneAssignToUnaryOpAssign(args[0].Generic_)
	case _ReduceSubOneAssignToUnaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpAssignType
		symbol.Generic_, err = reducer.SubOneAssignToUnaryOpAssign(args[0].Generic_)
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
	case _ReduceToUnsafeStatement:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = UnsafeStatementType
		symbol.Generic_, err = reducer.ToUnsafeStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceToJumpStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = JumpStatementType
		symbol.Generic_, err = reducer.ToJumpStatement(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	case _ReduceJumpLabelToOptionalJumpLabel:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalJumpLabelType
		symbol.Generic_, err = reducer.JumpLabelToOptionalJumpLabel(args[0].Generic_)
	case _ReduceUnlabelledToOptionalJumpLabel:
		symbol.SymbolId_ = OptionalJumpLabelType
		symbol.Generic_, err = reducer.UnlabelledToOptionalJumpLabel()
	case _ReduceExpressionToExpressions:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionsType
		symbol.Generic_, err = reducer.ExpressionToExpressions(args[0].Generic_)
	case _ReduceAddToExpressions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExpressionsType
		symbol.Generic_, err = reducer.AddToExpressions(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceExpressionsToOptionalExpressions:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalExpressionsType
		symbol.Generic_, err = reducer.ExpressionsToOptionalExpressions(args[0].Generic_)
	case _ReduceNilToOptionalExpressions:
		symbol.SymbolId_ = OptionalExpressionsType
		symbol.Generic_, err = reducer.NilToOptionalExpressions()
	case _ReduceToCallExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CallExprType
		symbol.Generic_, err = reducer.ToCallExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceBindingToOptionalGenericBinding:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OptionalGenericBindingType
		symbol.Generic_, err = reducer.BindingToOptionalGenericBinding(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceNilToOptionalGenericBinding:
		symbol.SymbolId_ = OptionalGenericBindingType
		symbol.Generic_, err = reducer.NilToOptionalGenericBinding()
	case _ReduceGenericArgumentsToOptionalGenericArguments:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalGenericArgumentsType
		symbol.Generic_, err = reducer.GenericArgumentsToOptionalGenericArguments(args[0].Generic_)
	case _ReduceNilToOptionalGenericArguments:
		symbol.SymbolId_ = OptionalGenericArgumentsType
		symbol.Generic_, err = reducer.NilToOptionalGenericArguments()
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
	case _ReduceArgumentsToOptionalArguments:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalArgumentsType
		symbol.Generic_, err = reducer.ArgumentsToOptionalArguments(args[0].Generic_)
	case _ReduceNilToOptionalArguments:
		symbol.SymbolId_ = OptionalArgumentsType
		symbol.Generic_, err = reducer.NilToOptionalArguments()
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
	case _ReduceExpressionToOptionalExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalExpressionType
		symbol.Generic_, err = reducer.ExpressionToOptionalExpression(args[0].Generic_)
	case _ReduceNilToOptionalExpression:
		symbol.SymbolId_ = OptionalExpressionType
		symbol.Generic_, err = reducer.NilToOptionalExpression()
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
	case _ReduceUnsafeStatementToFieldDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldDefType
		symbol.Generic_, err = reducer.UnsafeStatementToFieldDef(args[0].Generic_)
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
	_GotoState1Action                                               = &_Action{_ShiftAction, _State1, 0}
	_GotoState2Action                                               = &_Action{_ShiftAction, _State2, 0}
	_GotoState3Action                                               = &_Action{_ShiftAction, _State3, 0}
	_GotoState4Action                                               = &_Action{_ShiftAction, _State4, 0}
	_GotoState5Action                                               = &_Action{_ShiftAction, _State5, 0}
	_GotoState6Action                                               = &_Action{_ShiftAction, _State6, 0}
	_GotoState7Action                                               = &_Action{_ShiftAction, _State7, 0}
	_GotoState8Action                                               = &_Action{_ShiftAction, _State8, 0}
	_GotoState9Action                                               = &_Action{_ShiftAction, _State9, 0}
	_GotoState10Action                                              = &_Action{_ShiftAction, _State10, 0}
	_GotoState11Action                                              = &_Action{_ShiftAction, _State11, 0}
	_GotoState12Action                                              = &_Action{_ShiftAction, _State12, 0}
	_GotoState13Action                                              = &_Action{_ShiftAction, _State13, 0}
	_GotoState14Action                                              = &_Action{_ShiftAction, _State14, 0}
	_GotoState15Action                                              = &_Action{_ShiftAction, _State15, 0}
	_GotoState16Action                                              = &_Action{_ShiftAction, _State16, 0}
	_GotoState17Action                                              = &_Action{_ShiftAction, _State17, 0}
	_GotoState18Action                                              = &_Action{_ShiftAction, _State18, 0}
	_GotoState19Action                                              = &_Action{_ShiftAction, _State19, 0}
	_GotoState20Action                                              = &_Action{_ShiftAction, _State20, 0}
	_GotoState21Action                                              = &_Action{_ShiftAction, _State21, 0}
	_GotoState22Action                                              = &_Action{_ShiftAction, _State22, 0}
	_GotoState23Action                                              = &_Action{_ShiftAction, _State23, 0}
	_GotoState24Action                                              = &_Action{_ShiftAction, _State24, 0}
	_GotoState25Action                                              = &_Action{_ShiftAction, _State25, 0}
	_GotoState26Action                                              = &_Action{_ShiftAction, _State26, 0}
	_GotoState27Action                                              = &_Action{_ShiftAction, _State27, 0}
	_GotoState28Action                                              = &_Action{_ShiftAction, _State28, 0}
	_GotoState29Action                                              = &_Action{_ShiftAction, _State29, 0}
	_GotoState30Action                                              = &_Action{_ShiftAction, _State30, 0}
	_GotoState31Action                                              = &_Action{_ShiftAction, _State31, 0}
	_GotoState32Action                                              = &_Action{_ShiftAction, _State32, 0}
	_GotoState33Action                                              = &_Action{_ShiftAction, _State33, 0}
	_GotoState34Action                                              = &_Action{_ShiftAction, _State34, 0}
	_GotoState35Action                                              = &_Action{_ShiftAction, _State35, 0}
	_GotoState36Action                                              = &_Action{_ShiftAction, _State36, 0}
	_GotoState37Action                                              = &_Action{_ShiftAction, _State37, 0}
	_GotoState38Action                                              = &_Action{_ShiftAction, _State38, 0}
	_GotoState39Action                                              = &_Action{_ShiftAction, _State39, 0}
	_GotoState40Action                                              = &_Action{_ShiftAction, _State40, 0}
	_GotoState41Action                                              = &_Action{_ShiftAction, _State41, 0}
	_GotoState42Action                                              = &_Action{_ShiftAction, _State42, 0}
	_GotoState43Action                                              = &_Action{_ShiftAction, _State43, 0}
	_GotoState44Action                                              = &_Action{_ShiftAction, _State44, 0}
	_GotoState45Action                                              = &_Action{_ShiftAction, _State45, 0}
	_GotoState46Action                                              = &_Action{_ShiftAction, _State46, 0}
	_GotoState47Action                                              = &_Action{_ShiftAction, _State47, 0}
	_GotoState48Action                                              = &_Action{_ShiftAction, _State48, 0}
	_GotoState49Action                                              = &_Action{_ShiftAction, _State49, 0}
	_GotoState50Action                                              = &_Action{_ShiftAction, _State50, 0}
	_GotoState51Action                                              = &_Action{_ShiftAction, _State51, 0}
	_GotoState52Action                                              = &_Action{_ShiftAction, _State52, 0}
	_GotoState53Action                                              = &_Action{_ShiftAction, _State53, 0}
	_GotoState54Action                                              = &_Action{_ShiftAction, _State54, 0}
	_GotoState55Action                                              = &_Action{_ShiftAction, _State55, 0}
	_GotoState56Action                                              = &_Action{_ShiftAction, _State56, 0}
	_GotoState57Action                                              = &_Action{_ShiftAction, _State57, 0}
	_GotoState58Action                                              = &_Action{_ShiftAction, _State58, 0}
	_GotoState59Action                                              = &_Action{_ShiftAction, _State59, 0}
	_GotoState60Action                                              = &_Action{_ShiftAction, _State60, 0}
	_GotoState61Action                                              = &_Action{_ShiftAction, _State61, 0}
	_GotoState62Action                                              = &_Action{_ShiftAction, _State62, 0}
	_GotoState63Action                                              = &_Action{_ShiftAction, _State63, 0}
	_GotoState64Action                                              = &_Action{_ShiftAction, _State64, 0}
	_GotoState65Action                                              = &_Action{_ShiftAction, _State65, 0}
	_GotoState66Action                                              = &_Action{_ShiftAction, _State66, 0}
	_GotoState67Action                                              = &_Action{_ShiftAction, _State67, 0}
	_GotoState68Action                                              = &_Action{_ShiftAction, _State68, 0}
	_GotoState69Action                                              = &_Action{_ShiftAction, _State69, 0}
	_GotoState70Action                                              = &_Action{_ShiftAction, _State70, 0}
	_GotoState71Action                                              = &_Action{_ShiftAction, _State71, 0}
	_GotoState72Action                                              = &_Action{_ShiftAction, _State72, 0}
	_GotoState73Action                                              = &_Action{_ShiftAction, _State73, 0}
	_GotoState74Action                                              = &_Action{_ShiftAction, _State74, 0}
	_GotoState75Action                                              = &_Action{_ShiftAction, _State75, 0}
	_GotoState76Action                                              = &_Action{_ShiftAction, _State76, 0}
	_GotoState77Action                                              = &_Action{_ShiftAction, _State77, 0}
	_GotoState78Action                                              = &_Action{_ShiftAction, _State78, 0}
	_GotoState79Action                                              = &_Action{_ShiftAction, _State79, 0}
	_GotoState80Action                                              = &_Action{_ShiftAction, _State80, 0}
	_GotoState81Action                                              = &_Action{_ShiftAction, _State81, 0}
	_GotoState82Action                                              = &_Action{_ShiftAction, _State82, 0}
	_GotoState83Action                                              = &_Action{_ShiftAction, _State83, 0}
	_GotoState84Action                                              = &_Action{_ShiftAction, _State84, 0}
	_GotoState85Action                                              = &_Action{_ShiftAction, _State85, 0}
	_GotoState86Action                                              = &_Action{_ShiftAction, _State86, 0}
	_GotoState87Action                                              = &_Action{_ShiftAction, _State87, 0}
	_GotoState88Action                                              = &_Action{_ShiftAction, _State88, 0}
	_GotoState89Action                                              = &_Action{_ShiftAction, _State89, 0}
	_GotoState90Action                                              = &_Action{_ShiftAction, _State90, 0}
	_GotoState91Action                                              = &_Action{_ShiftAction, _State91, 0}
	_GotoState92Action                                              = &_Action{_ShiftAction, _State92, 0}
	_GotoState93Action                                              = &_Action{_ShiftAction, _State93, 0}
	_GotoState94Action                                              = &_Action{_ShiftAction, _State94, 0}
	_GotoState95Action                                              = &_Action{_ShiftAction, _State95, 0}
	_GotoState96Action                                              = &_Action{_ShiftAction, _State96, 0}
	_GotoState97Action                                              = &_Action{_ShiftAction, _State97, 0}
	_GotoState98Action                                              = &_Action{_ShiftAction, _State98, 0}
	_GotoState99Action                                              = &_Action{_ShiftAction, _State99, 0}
	_GotoState100Action                                             = &_Action{_ShiftAction, _State100, 0}
	_GotoState101Action                                             = &_Action{_ShiftAction, _State101, 0}
	_GotoState102Action                                             = &_Action{_ShiftAction, _State102, 0}
	_GotoState103Action                                             = &_Action{_ShiftAction, _State103, 0}
	_GotoState104Action                                             = &_Action{_ShiftAction, _State104, 0}
	_GotoState105Action                                             = &_Action{_ShiftAction, _State105, 0}
	_GotoState106Action                                             = &_Action{_ShiftAction, _State106, 0}
	_GotoState107Action                                             = &_Action{_ShiftAction, _State107, 0}
	_GotoState108Action                                             = &_Action{_ShiftAction, _State108, 0}
	_GotoState109Action                                             = &_Action{_ShiftAction, _State109, 0}
	_GotoState110Action                                             = &_Action{_ShiftAction, _State110, 0}
	_GotoState111Action                                             = &_Action{_ShiftAction, _State111, 0}
	_GotoState112Action                                             = &_Action{_ShiftAction, _State112, 0}
	_GotoState113Action                                             = &_Action{_ShiftAction, _State113, 0}
	_GotoState114Action                                             = &_Action{_ShiftAction, _State114, 0}
	_GotoState115Action                                             = &_Action{_ShiftAction, _State115, 0}
	_GotoState116Action                                             = &_Action{_ShiftAction, _State116, 0}
	_GotoState117Action                                             = &_Action{_ShiftAction, _State117, 0}
	_GotoState118Action                                             = &_Action{_ShiftAction, _State118, 0}
	_GotoState119Action                                             = &_Action{_ShiftAction, _State119, 0}
	_GotoState120Action                                             = &_Action{_ShiftAction, _State120, 0}
	_GotoState121Action                                             = &_Action{_ShiftAction, _State121, 0}
	_GotoState122Action                                             = &_Action{_ShiftAction, _State122, 0}
	_GotoState123Action                                             = &_Action{_ShiftAction, _State123, 0}
	_GotoState124Action                                             = &_Action{_ShiftAction, _State124, 0}
	_GotoState125Action                                             = &_Action{_ShiftAction, _State125, 0}
	_GotoState126Action                                             = &_Action{_ShiftAction, _State126, 0}
	_GotoState127Action                                             = &_Action{_ShiftAction, _State127, 0}
	_GotoState128Action                                             = &_Action{_ShiftAction, _State128, 0}
	_GotoState129Action                                             = &_Action{_ShiftAction, _State129, 0}
	_GotoState130Action                                             = &_Action{_ShiftAction, _State130, 0}
	_GotoState131Action                                             = &_Action{_ShiftAction, _State131, 0}
	_GotoState132Action                                             = &_Action{_ShiftAction, _State132, 0}
	_GotoState133Action                                             = &_Action{_ShiftAction, _State133, 0}
	_GotoState134Action                                             = &_Action{_ShiftAction, _State134, 0}
	_GotoState135Action                                             = &_Action{_ShiftAction, _State135, 0}
	_GotoState136Action                                             = &_Action{_ShiftAction, _State136, 0}
	_GotoState137Action                                             = &_Action{_ShiftAction, _State137, 0}
	_GotoState138Action                                             = &_Action{_ShiftAction, _State138, 0}
	_GotoState139Action                                             = &_Action{_ShiftAction, _State139, 0}
	_GotoState140Action                                             = &_Action{_ShiftAction, _State140, 0}
	_GotoState141Action                                             = &_Action{_ShiftAction, _State141, 0}
	_GotoState142Action                                             = &_Action{_ShiftAction, _State142, 0}
	_GotoState143Action                                             = &_Action{_ShiftAction, _State143, 0}
	_GotoState144Action                                             = &_Action{_ShiftAction, _State144, 0}
	_GotoState145Action                                             = &_Action{_ShiftAction, _State145, 0}
	_GotoState146Action                                             = &_Action{_ShiftAction, _State146, 0}
	_GotoState147Action                                             = &_Action{_ShiftAction, _State147, 0}
	_GotoState148Action                                             = &_Action{_ShiftAction, _State148, 0}
	_GotoState149Action                                             = &_Action{_ShiftAction, _State149, 0}
	_GotoState150Action                                             = &_Action{_ShiftAction, _State150, 0}
	_GotoState151Action                                             = &_Action{_ShiftAction, _State151, 0}
	_GotoState152Action                                             = &_Action{_ShiftAction, _State152, 0}
	_GotoState153Action                                             = &_Action{_ShiftAction, _State153, 0}
	_GotoState154Action                                             = &_Action{_ShiftAction, _State154, 0}
	_GotoState155Action                                             = &_Action{_ShiftAction, _State155, 0}
	_GotoState156Action                                             = &_Action{_ShiftAction, _State156, 0}
	_GotoState157Action                                             = &_Action{_ShiftAction, _State157, 0}
	_GotoState158Action                                             = &_Action{_ShiftAction, _State158, 0}
	_GotoState159Action                                             = &_Action{_ShiftAction, _State159, 0}
	_GotoState160Action                                             = &_Action{_ShiftAction, _State160, 0}
	_GotoState161Action                                             = &_Action{_ShiftAction, _State161, 0}
	_GotoState162Action                                             = &_Action{_ShiftAction, _State162, 0}
	_GotoState163Action                                             = &_Action{_ShiftAction, _State163, 0}
	_GotoState164Action                                             = &_Action{_ShiftAction, _State164, 0}
	_GotoState165Action                                             = &_Action{_ShiftAction, _State165, 0}
	_GotoState166Action                                             = &_Action{_ShiftAction, _State166, 0}
	_GotoState167Action                                             = &_Action{_ShiftAction, _State167, 0}
	_GotoState168Action                                             = &_Action{_ShiftAction, _State168, 0}
	_GotoState169Action                                             = &_Action{_ShiftAction, _State169, 0}
	_GotoState170Action                                             = &_Action{_ShiftAction, _State170, 0}
	_GotoState171Action                                             = &_Action{_ShiftAction, _State171, 0}
	_GotoState172Action                                             = &_Action{_ShiftAction, _State172, 0}
	_GotoState173Action                                             = &_Action{_ShiftAction, _State173, 0}
	_GotoState174Action                                             = &_Action{_ShiftAction, _State174, 0}
	_GotoState175Action                                             = &_Action{_ShiftAction, _State175, 0}
	_GotoState176Action                                             = &_Action{_ShiftAction, _State176, 0}
	_GotoState177Action                                             = &_Action{_ShiftAction, _State177, 0}
	_GotoState178Action                                             = &_Action{_ShiftAction, _State178, 0}
	_GotoState179Action                                             = &_Action{_ShiftAction, _State179, 0}
	_GotoState180Action                                             = &_Action{_ShiftAction, _State180, 0}
	_GotoState181Action                                             = &_Action{_ShiftAction, _State181, 0}
	_GotoState182Action                                             = &_Action{_ShiftAction, _State182, 0}
	_GotoState183Action                                             = &_Action{_ShiftAction, _State183, 0}
	_GotoState184Action                                             = &_Action{_ShiftAction, _State184, 0}
	_GotoState185Action                                             = &_Action{_ShiftAction, _State185, 0}
	_GotoState186Action                                             = &_Action{_ShiftAction, _State186, 0}
	_GotoState187Action                                             = &_Action{_ShiftAction, _State187, 0}
	_GotoState188Action                                             = &_Action{_ShiftAction, _State188, 0}
	_GotoState189Action                                             = &_Action{_ShiftAction, _State189, 0}
	_GotoState190Action                                             = &_Action{_ShiftAction, _State190, 0}
	_GotoState191Action                                             = &_Action{_ShiftAction, _State191, 0}
	_GotoState192Action                                             = &_Action{_ShiftAction, _State192, 0}
	_GotoState193Action                                             = &_Action{_ShiftAction, _State193, 0}
	_GotoState194Action                                             = &_Action{_ShiftAction, _State194, 0}
	_GotoState195Action                                             = &_Action{_ShiftAction, _State195, 0}
	_GotoState196Action                                             = &_Action{_ShiftAction, _State196, 0}
	_GotoState197Action                                             = &_Action{_ShiftAction, _State197, 0}
	_GotoState198Action                                             = &_Action{_ShiftAction, _State198, 0}
	_GotoState199Action                                             = &_Action{_ShiftAction, _State199, 0}
	_GotoState200Action                                             = &_Action{_ShiftAction, _State200, 0}
	_GotoState201Action                                             = &_Action{_ShiftAction, _State201, 0}
	_GotoState202Action                                             = &_Action{_ShiftAction, _State202, 0}
	_GotoState203Action                                             = &_Action{_ShiftAction, _State203, 0}
	_GotoState204Action                                             = &_Action{_ShiftAction, _State204, 0}
	_GotoState205Action                                             = &_Action{_ShiftAction, _State205, 0}
	_GotoState206Action                                             = &_Action{_ShiftAction, _State206, 0}
	_GotoState207Action                                             = &_Action{_ShiftAction, _State207, 0}
	_GotoState208Action                                             = &_Action{_ShiftAction, _State208, 0}
	_GotoState209Action                                             = &_Action{_ShiftAction, _State209, 0}
	_GotoState210Action                                             = &_Action{_ShiftAction, _State210, 0}
	_GotoState211Action                                             = &_Action{_ShiftAction, _State211, 0}
	_GotoState212Action                                             = &_Action{_ShiftAction, _State212, 0}
	_GotoState213Action                                             = &_Action{_ShiftAction, _State213, 0}
	_GotoState214Action                                             = &_Action{_ShiftAction, _State214, 0}
	_GotoState215Action                                             = &_Action{_ShiftAction, _State215, 0}
	_GotoState216Action                                             = &_Action{_ShiftAction, _State216, 0}
	_GotoState217Action                                             = &_Action{_ShiftAction, _State217, 0}
	_GotoState218Action                                             = &_Action{_ShiftAction, _State218, 0}
	_GotoState219Action                                             = &_Action{_ShiftAction, _State219, 0}
	_GotoState220Action                                             = &_Action{_ShiftAction, _State220, 0}
	_GotoState221Action                                             = &_Action{_ShiftAction, _State221, 0}
	_GotoState222Action                                             = &_Action{_ShiftAction, _State222, 0}
	_GotoState223Action                                             = &_Action{_ShiftAction, _State223, 0}
	_GotoState224Action                                             = &_Action{_ShiftAction, _State224, 0}
	_GotoState225Action                                             = &_Action{_ShiftAction, _State225, 0}
	_GotoState226Action                                             = &_Action{_ShiftAction, _State226, 0}
	_GotoState227Action                                             = &_Action{_ShiftAction, _State227, 0}
	_GotoState228Action                                             = &_Action{_ShiftAction, _State228, 0}
	_GotoState229Action                                             = &_Action{_ShiftAction, _State229, 0}
	_GotoState230Action                                             = &_Action{_ShiftAction, _State230, 0}
	_GotoState231Action                                             = &_Action{_ShiftAction, _State231, 0}
	_GotoState232Action                                             = &_Action{_ShiftAction, _State232, 0}
	_GotoState233Action                                             = &_Action{_ShiftAction, _State233, 0}
	_GotoState234Action                                             = &_Action{_ShiftAction, _State234, 0}
	_GotoState235Action                                             = &_Action{_ShiftAction, _State235, 0}
	_GotoState236Action                                             = &_Action{_ShiftAction, _State236, 0}
	_GotoState237Action                                             = &_Action{_ShiftAction, _State237, 0}
	_GotoState238Action                                             = &_Action{_ShiftAction, _State238, 0}
	_GotoState239Action                                             = &_Action{_ShiftAction, _State239, 0}
	_GotoState240Action                                             = &_Action{_ShiftAction, _State240, 0}
	_GotoState241Action                                             = &_Action{_ShiftAction, _State241, 0}
	_GotoState242Action                                             = &_Action{_ShiftAction, _State242, 0}
	_GotoState243Action                                             = &_Action{_ShiftAction, _State243, 0}
	_GotoState244Action                                             = &_Action{_ShiftAction, _State244, 0}
	_GotoState245Action                                             = &_Action{_ShiftAction, _State245, 0}
	_GotoState246Action                                             = &_Action{_ShiftAction, _State246, 0}
	_GotoState247Action                                             = &_Action{_ShiftAction, _State247, 0}
	_GotoState248Action                                             = &_Action{_ShiftAction, _State248, 0}
	_GotoState249Action                                             = &_Action{_ShiftAction, _State249, 0}
	_GotoState250Action                                             = &_Action{_ShiftAction, _State250, 0}
	_GotoState251Action                                             = &_Action{_ShiftAction, _State251, 0}
	_GotoState252Action                                             = &_Action{_ShiftAction, _State252, 0}
	_GotoState253Action                                             = &_Action{_ShiftAction, _State253, 0}
	_GotoState254Action                                             = &_Action{_ShiftAction, _State254, 0}
	_GotoState255Action                                             = &_Action{_ShiftAction, _State255, 0}
	_GotoState256Action                                             = &_Action{_ShiftAction, _State256, 0}
	_GotoState257Action                                             = &_Action{_ShiftAction, _State257, 0}
	_GotoState258Action                                             = &_Action{_ShiftAction, _State258, 0}
	_GotoState259Action                                             = &_Action{_ShiftAction, _State259, 0}
	_GotoState260Action                                             = &_Action{_ShiftAction, _State260, 0}
	_GotoState261Action                                             = &_Action{_ShiftAction, _State261, 0}
	_GotoState262Action                                             = &_Action{_ShiftAction, _State262, 0}
	_GotoState263Action                                             = &_Action{_ShiftAction, _State263, 0}
	_GotoState264Action                                             = &_Action{_ShiftAction, _State264, 0}
	_GotoState265Action                                             = &_Action{_ShiftAction, _State265, 0}
	_GotoState266Action                                             = &_Action{_ShiftAction, _State266, 0}
	_GotoState267Action                                             = &_Action{_ShiftAction, _State267, 0}
	_GotoState268Action                                             = &_Action{_ShiftAction, _State268, 0}
	_GotoState269Action                                             = &_Action{_ShiftAction, _State269, 0}
	_GotoState270Action                                             = &_Action{_ShiftAction, _State270, 0}
	_GotoState271Action                                             = &_Action{_ShiftAction, _State271, 0}
	_GotoState272Action                                             = &_Action{_ShiftAction, _State272, 0}
	_GotoState273Action                                             = &_Action{_ShiftAction, _State273, 0}
	_GotoState274Action                                             = &_Action{_ShiftAction, _State274, 0}
	_GotoState275Action                                             = &_Action{_ShiftAction, _State275, 0}
	_GotoState276Action                                             = &_Action{_ShiftAction, _State276, 0}
	_GotoState277Action                                             = &_Action{_ShiftAction, _State277, 0}
	_GotoState278Action                                             = &_Action{_ShiftAction, _State278, 0}
	_GotoState279Action                                             = &_Action{_ShiftAction, _State279, 0}
	_GotoState280Action                                             = &_Action{_ShiftAction, _State280, 0}
	_GotoState281Action                                             = &_Action{_ShiftAction, _State281, 0}
	_GotoState282Action                                             = &_Action{_ShiftAction, _State282, 0}
	_GotoState283Action                                             = &_Action{_ShiftAction, _State283, 0}
	_GotoState284Action                                             = &_Action{_ShiftAction, _State284, 0}
	_GotoState285Action                                             = &_Action{_ShiftAction, _State285, 0}
	_GotoState286Action                                             = &_Action{_ShiftAction, _State286, 0}
	_GotoState287Action                                             = &_Action{_ShiftAction, _State287, 0}
	_GotoState288Action                                             = &_Action{_ShiftAction, _State288, 0}
	_GotoState289Action                                             = &_Action{_ShiftAction, _State289, 0}
	_GotoState290Action                                             = &_Action{_ShiftAction, _State290, 0}
	_GotoState291Action                                             = &_Action{_ShiftAction, _State291, 0}
	_GotoState292Action                                             = &_Action{_ShiftAction, _State292, 0}
	_GotoState293Action                                             = &_Action{_ShiftAction, _State293, 0}
	_GotoState294Action                                             = &_Action{_ShiftAction, _State294, 0}
	_GotoState295Action                                             = &_Action{_ShiftAction, _State295, 0}
	_GotoState296Action                                             = &_Action{_ShiftAction, _State296, 0}
	_GotoState297Action                                             = &_Action{_ShiftAction, _State297, 0}
	_GotoState298Action                                             = &_Action{_ShiftAction, _State298, 0}
	_GotoState299Action                                             = &_Action{_ShiftAction, _State299, 0}
	_GotoState300Action                                             = &_Action{_ShiftAction, _State300, 0}
	_GotoState301Action                                             = &_Action{_ShiftAction, _State301, 0}
	_GotoState302Action                                             = &_Action{_ShiftAction, _State302, 0}
	_GotoState303Action                                             = &_Action{_ShiftAction, _State303, 0}
	_GotoState304Action                                             = &_Action{_ShiftAction, _State304, 0}
	_GotoState305Action                                             = &_Action{_ShiftAction, _State305, 0}
	_GotoState306Action                                             = &_Action{_ShiftAction, _State306, 0}
	_GotoState307Action                                             = &_Action{_ShiftAction, _State307, 0}
	_GotoState308Action                                             = &_Action{_ShiftAction, _State308, 0}
	_GotoState309Action                                             = &_Action{_ShiftAction, _State309, 0}
	_GotoState310Action                                             = &_Action{_ShiftAction, _State310, 0}
	_GotoState311Action                                             = &_Action{_ShiftAction, _State311, 0}
	_GotoState312Action                                             = &_Action{_ShiftAction, _State312, 0}
	_GotoState313Action                                             = &_Action{_ShiftAction, _State313, 0}
	_GotoState314Action                                             = &_Action{_ShiftAction, _State314, 0}
	_GotoState315Action                                             = &_Action{_ShiftAction, _State315, 0}
	_GotoState316Action                                             = &_Action{_ShiftAction, _State316, 0}
	_GotoState317Action                                             = &_Action{_ShiftAction, _State317, 0}
	_GotoState318Action                                             = &_Action{_ShiftAction, _State318, 0}
	_GotoState319Action                                             = &_Action{_ShiftAction, _State319, 0}
	_GotoState320Action                                             = &_Action{_ShiftAction, _State320, 0}
	_GotoState321Action                                             = &_Action{_ShiftAction, _State321, 0}
	_GotoState322Action                                             = &_Action{_ShiftAction, _State322, 0}
	_GotoState323Action                                             = &_Action{_ShiftAction, _State323, 0}
	_GotoState324Action                                             = &_Action{_ShiftAction, _State324, 0}
	_GotoState325Action                                             = &_Action{_ShiftAction, _State325, 0}
	_GotoState326Action                                             = &_Action{_ShiftAction, _State326, 0}
	_GotoState327Action                                             = &_Action{_ShiftAction, _State327, 0}
	_GotoState328Action                                             = &_Action{_ShiftAction, _State328, 0}
	_GotoState329Action                                             = &_Action{_ShiftAction, _State329, 0}
	_GotoState330Action                                             = &_Action{_ShiftAction, _State330, 0}
	_GotoState331Action                                             = &_Action{_ShiftAction, _State331, 0}
	_GotoState332Action                                             = &_Action{_ShiftAction, _State332, 0}
	_GotoState333Action                                             = &_Action{_ShiftAction, _State333, 0}
	_GotoState334Action                                             = &_Action{_ShiftAction, _State334, 0}
	_GotoState335Action                                             = &_Action{_ShiftAction, _State335, 0}
	_GotoState336Action                                             = &_Action{_ShiftAction, _State336, 0}
	_GotoState337Action                                             = &_Action{_ShiftAction, _State337, 0}
	_GotoState338Action                                             = &_Action{_ShiftAction, _State338, 0}
	_GotoState339Action                                             = &_Action{_ShiftAction, _State339, 0}
	_GotoState340Action                                             = &_Action{_ShiftAction, _State340, 0}
	_GotoState341Action                                             = &_Action{_ShiftAction, _State341, 0}
	_GotoState342Action                                             = &_Action{_ShiftAction, _State342, 0}
	_GotoState343Action                                             = &_Action{_ShiftAction, _State343, 0}
	_GotoState344Action                                             = &_Action{_ShiftAction, _State344, 0}
	_GotoState345Action                                             = &_Action{_ShiftAction, _State345, 0}
	_GotoState346Action                                             = &_Action{_ShiftAction, _State346, 0}
	_GotoState347Action                                             = &_Action{_ShiftAction, _State347, 0}
	_GotoState348Action                                             = &_Action{_ShiftAction, _State348, 0}
	_GotoState349Action                                             = &_Action{_ShiftAction, _State349, 0}
	_GotoState350Action                                             = &_Action{_ShiftAction, _State350, 0}
	_GotoState351Action                                             = &_Action{_ShiftAction, _State351, 0}
	_GotoState352Action                                             = &_Action{_ShiftAction, _State352, 0}
	_ReduceIfExprToExpressionAction                                 = &_Action{_ReduceAction, 0, _ReduceIfExprToExpression}
	_ReduceSwitchExprToExpressionAction                             = &_Action{_ReduceAction, 0, _ReduceSwitchExprToExpression}
	_ReduceLoopExprToExpressionAction                               = &_Action{_ReduceAction, 0, _ReduceLoopExprToExpression}
	_ReduceSequenceExprToExpressionAction                           = &_Action{_ReduceAction, 0, _ReduceSequenceExprToExpression}
	_ReduceLabelDeclToOptionalLabelDeclAction                       = &_Action{_ReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}
	_ReduceUnlabelledToOptionalLabelDeclAction                      = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}
	_ReduceNoElseToIfExprAction                                     = &_Action{_ReduceAction, 0, _ReduceNoElseToIfExpr}
	_ReduceIfElseToIfExprAction                                     = &_Action{_ReduceAction, 0, _ReduceIfElseToIfExpr}
	_ReduceMultiIfElseToIfExprAction                                = &_Action{_ReduceAction, 0, _ReduceMultiIfElseToIfExpr}
	_ReduceToSwitchExprAction                                       = &_Action{_ReduceAction, 0, _ReduceToSwitchExpr}
	_ReduceInfiniteToLoopExprAction                                 = &_Action{_ReduceAction, 0, _ReduceInfiniteToLoopExpr}
	_ReduceDoWhileToLoopExprAction                                  = &_Action{_ReduceAction, 0, _ReduceDoWhileToLoopExpr}
	_ReduceWhileToLoopExprAction                                    = &_Action{_ReduceAction, 0, _ReduceWhileToLoopExpr}
	_ReduceIteratorToLoopExprAction                                 = &_Action{_ReduceAction, 0, _ReduceIteratorToLoopExpr}
	_ReduceForToLoopExprAction                                      = &_Action{_ReduceAction, 0, _ReduceForToLoopExpr}
	_ReduceSequenceExprToOptionalSequenceExprAction                 = &_Action{_ReduceAction, 0, _ReduceSequenceExprToOptionalSequenceExpr}
	_ReduceNilToOptionalSequenceExprAction                          = &_Action{_ReduceAction, 0, _ReduceNilToOptionalSequenceExpr}
	_ReduceToSequenceExprAction                                     = &_Action{_ReduceAction, 0, _ReduceToSequenceExpr}
	_ReduceToBlockExprAction                                        = &_Action{_ReduceAction, 0, _ReduceToBlockExpr}
	_ReduceToBlockBodyAction                                        = &_Action{_ReduceAction, 0, _ReduceToBlockBody}
	_ReduceEmptyListToStatementsAction                              = &_Action{_ReduceAction, 0, _ReduceEmptyListToStatements}
	_ReduceAddToStatementsAction                                    = &_Action{_ReduceAction, 0, _ReduceAddToStatements}
	_ReduceImplicitToStatementAction                                = &_Action{_ReduceAction, 0, _ReduceImplicitToStatement}
	_ReduceExplicitToStatementAction                                = &_Action{_ReduceAction, 0, _ReduceExplicitToStatement}
	_ReduceUnsafeStatementToStatementBodyAction                     = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToStatementBody}
	_ReduceExpressionOrImplicitStructToStatementBodyAction          = &_Action{_ReduceAction, 0, _ReduceExpressionOrImplicitStructToStatementBody}
	_ReduceAsyncToStatementBodyAction                               = &_Action{_ReduceAction, 0, _ReduceAsyncToStatementBody}
	_ReduceDeferToStatementBodyAction                               = &_Action{_ReduceAction, 0, _ReduceDeferToStatementBody}
	_ReduceJumpStatementToStatementBodyAction                       = &_Action{_ReduceAction, 0, _ReduceJumpStatementToStatementBody}
	_ReduceUnaryOpAssignStatementToStatementBodyAction              = &_Action{_ReduceAction, 0, _ReduceUnaryOpAssignStatementToStatementBody}
	_ReduceBinaryOpAssignStatementToStatementBodyAction             = &_Action{_ReduceAction, 0, _ReduceBinaryOpAssignStatementToStatementBody}
	_ReduceAddOneAssignToUnaryOpAssignAction                        = &_Action{_ReduceAction, 0, _ReduceAddOneAssignToUnaryOpAssign}
	_ReduceSubOneAssignToUnaryOpAssignAction                        = &_Action{_ReduceAction, 0, _ReduceSubOneAssignToUnaryOpAssign}
	_ReduceAddAssignToBinaryOpAssignAction                          = &_Action{_ReduceAction, 0, _ReduceAddAssignToBinaryOpAssign}
	_ReduceSubAssignToBinaryOpAssignAction                          = &_Action{_ReduceAction, 0, _ReduceSubAssignToBinaryOpAssign}
	_ReduceMulAssignToBinaryOpAssignAction                          = &_Action{_ReduceAction, 0, _ReduceMulAssignToBinaryOpAssign}
	_ReduceDivAssignToBinaryOpAssignAction                          = &_Action{_ReduceAction, 0, _ReduceDivAssignToBinaryOpAssign}
	_ReduceModAssignToBinaryOpAssignAction                          = &_Action{_ReduceAction, 0, _ReduceModAssignToBinaryOpAssign}
	_ReduceBitNegAssignToBinaryOpAssignAction                       = &_Action{_ReduceAction, 0, _ReduceBitNegAssignToBinaryOpAssign}
	_ReduceBitAndAssignToBinaryOpAssignAction                       = &_Action{_ReduceAction, 0, _ReduceBitAndAssignToBinaryOpAssign}
	_ReduceBitOrAssignToBinaryOpAssignAction                        = &_Action{_ReduceAction, 0, _ReduceBitOrAssignToBinaryOpAssign}
	_ReduceBitXorAssignToBinaryOpAssignAction                       = &_Action{_ReduceAction, 0, _ReduceBitXorAssignToBinaryOpAssign}
	_ReduceBitLshiftAssignToBinaryOpAssignAction                    = &_Action{_ReduceAction, 0, _ReduceBitLshiftAssignToBinaryOpAssign}
	_ReduceBitRshiftAssignToBinaryOpAssignAction                    = &_Action{_ReduceAction, 0, _ReduceBitRshiftAssignToBinaryOpAssign}
	_ReduceToUnsafeStatementAction                                  = &_Action{_ReduceAction, 0, _ReduceToUnsafeStatement}
	_ReduceToJumpStatementAction                                    = &_Action{_ReduceAction, 0, _ReduceToJumpStatement}
	_ReduceReturnToJumpTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceReturnToJumpType}
	_ReduceBreakToJumpTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceBreakToJumpType}
	_ReduceContinueToJumpTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceContinueToJumpType}
	_ReduceJumpLabelToOptionalJumpLabelAction                       = &_Action{_ReduceAction, 0, _ReduceJumpLabelToOptionalJumpLabel}
	_ReduceUnlabelledToOptionalJumpLabelAction                      = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalJumpLabel}
	_ReduceExpressionToExpressionsAction                            = &_Action{_ReduceAction, 0, _ReduceExpressionToExpressions}
	_ReduceAddToExpressionsAction                                   = &_Action{_ReduceAction, 0, _ReduceAddToExpressions}
	_ReduceExpressionsToOptionalExpressionsAction                   = &_Action{_ReduceAction, 0, _ReduceExpressionsToOptionalExpressions}
	_ReduceNilToOptionalExpressionsAction                           = &_Action{_ReduceAction, 0, _ReduceNilToOptionalExpressions}
	_ReduceToCallExprAction                                         = &_Action{_ReduceAction, 0, _ReduceToCallExpr}
	_ReduceBindingToOptionalGenericBindingAction                    = &_Action{_ReduceAction, 0, _ReduceBindingToOptionalGenericBinding}
	_ReduceNilToOptionalGenericBindingAction                        = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericBinding}
	_ReduceGenericArgumentsToOptionalGenericArgumentsAction         = &_Action{_ReduceAction, 0, _ReduceGenericArgumentsToOptionalGenericArguments}
	_ReduceNilToOptionalGenericArgumentsAction                      = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericArguments}
	_ReduceValueTypeToGenericArgumentsAction                        = &_Action{_ReduceAction, 0, _ReduceValueTypeToGenericArguments}
	_ReduceAddToGenericArgumentsAction                              = &_Action{_ReduceAction, 0, _ReduceAddToGenericArguments}
	_ReduceArgumentsToOptionalArgumentsAction                       = &_Action{_ReduceAction, 0, _ReduceArgumentsToOptionalArguments}
	_ReduceNilToOptionalArgumentsAction                             = &_Action{_ReduceAction, 0, _ReduceNilToOptionalArguments}
	_ReduceArgumentToArgumentsAction                                = &_Action{_ReduceAction, 0, _ReduceArgumentToArguments}
	_ReduceAddToArgumentsAction                                     = &_Action{_ReduceAction, 0, _ReduceAddToArguments}
	_ReducePositionalToArgumentAction                               = &_Action{_ReduceAction, 0, _ReducePositionalToArgument}
	_ReduceNamedToArgumentAction                                    = &_Action{_ReduceAction, 0, _ReduceNamedToArgument}
	_ReduceColonExpressionsToArgumentAction                         = &_Action{_ReduceAction, 0, _ReduceColonExpressionsToArgument}
	_ReducePairToColonExpressionsAction                             = &_Action{_ReduceAction, 0, _ReducePairToColonExpressions}
	_ReduceAddToColonExpressionsAction                              = &_Action{_ReduceAction, 0, _ReduceAddToColonExpressions}
	_ReduceExpressionToOptionalExpressionAction                     = &_Action{_ReduceAction, 0, _ReduceExpressionToOptionalExpression}
	_ReduceNilToOptionalExpressionAction                            = &_Action{_ReduceAction, 0, _ReduceNilToOptionalExpression}
	_ReduceLiteralToAtomExprAction                                  = &_Action{_ReduceAction, 0, _ReduceLiteralToAtomExpr}
	_ReduceIdentifierToAtomExprAction                               = &_Action{_ReduceAction, 0, _ReduceIdentifierToAtomExpr}
	_ReduceBlockExprToAtomExprAction                                = &_Action{_ReduceAction, 0, _ReduceBlockExprToAtomExpr}
	_ReduceAnonymousFuncExprToAtomExprAction                        = &_Action{_ReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}
	_ReduceAnonymousStructExprToAtomExprAction                      = &_Action{_ReduceAction, 0, _ReduceAnonymousStructExprToAtomExpr}
	_ReduceLexErrorToAtomExprAction                                 = &_Action{_ReduceAction, 0, _ReduceLexErrorToAtomExpr}
	_ReduceTrueToLiteralAction                                      = &_Action{_ReduceAction, 0, _ReduceTrueToLiteral}
	_ReduceFalseToLiteralAction                                     = &_Action{_ReduceAction, 0, _ReduceFalseToLiteral}
	_ReduceIntegerLiteralToLiteralAction                            = &_Action{_ReduceAction, 0, _ReduceIntegerLiteralToLiteral}
	_ReduceFloatLiteralToLiteralAction                              = &_Action{_ReduceAction, 0, _ReduceFloatLiteralToLiteral}
	_ReduceRuneLiteralToLiteralAction                               = &_Action{_ReduceAction, 0, _ReduceRuneLiteralToLiteral}
	_ReduceStringLiteralToLiteralAction                             = &_Action{_ReduceAction, 0, _ReduceStringLiteralToLiteral}
	_ReduceExplicitToAnonymousStructExprAction                      = &_Action{_ReduceAction, 0, _ReduceExplicitToAnonymousStructExpr}
	_ReduceImplicitToAnonymousStructExprAction                      = &_Action{_ReduceAction, 0, _ReduceImplicitToAnonymousStructExpr}
	_ReduceAtomExprToAccessExprAction                               = &_Action{_ReduceAction, 0, _ReduceAtomExprToAccessExpr}
	_ReduceAccessToAccessExprAction                                 = &_Action{_ReduceAction, 0, _ReduceAccessToAccessExpr}
	_ReduceCallExprToAccessExprAction                               = &_Action{_ReduceAction, 0, _ReduceCallExprToAccessExpr}
	_ReduceIndexToAccessExprAction                                  = &_Action{_ReduceAction, 0, _ReduceIndexToAccessExpr}
	_ReduceAccessExprToPostfixUnaryExprAction                       = &_Action{_ReduceAction, 0, _ReduceAccessExprToPostfixUnaryExpr}
	_ReduceQuestionToPostfixUnaryExprAction                         = &_Action{_ReduceAction, 0, _ReduceQuestionToPostfixUnaryExpr}
	_ReduceNotToPrefixUnaryOpAction                                 = &_Action{_ReduceAction, 0, _ReduceNotToPrefixUnaryOp}
	_ReduceBitNegToPrefixUnaryOpAction                              = &_Action{_ReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}
	_ReduceSubToPrefixUnaryOpAction                                 = &_Action{_ReduceAction, 0, _ReduceSubToPrefixUnaryOp}
	_ReduceMulToPrefixUnaryOpAction                                 = &_Action{_ReduceAction, 0, _ReduceMulToPrefixUnaryOp}
	_ReduceBitAndToPrefixUnaryOpAction                              = &_Action{_ReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}
	_ReducePostfixUnaryExprToPrefixUnaryExprAction                  = &_Action{_ReduceAction, 0, _ReducePostfixUnaryExprToPrefixUnaryExpr}
	_ReducePrefixOpToPrefixUnaryExprAction                          = &_Action{_ReduceAction, 0, _ReducePrefixOpToPrefixUnaryExpr}
	_ReduceMulToMulOpAction                                         = &_Action{_ReduceAction, 0, _ReduceMulToMulOp}
	_ReduceDivToMulOpAction                                         = &_Action{_ReduceAction, 0, _ReduceDivToMulOp}
	_ReduceModToMulOpAction                                         = &_Action{_ReduceAction, 0, _ReduceModToMulOp}
	_ReduceBitAndToMulOpAction                                      = &_Action{_ReduceAction, 0, _ReduceBitAndToMulOp}
	_ReduceBitLshiftToMulOpAction                                   = &_Action{_ReduceAction, 0, _ReduceBitLshiftToMulOp}
	_ReduceBitRshiftToMulOpAction                                   = &_Action{_ReduceAction, 0, _ReduceBitRshiftToMulOp}
	_ReducePrefixUnaryExprToMulExprAction                           = &_Action{_ReduceAction, 0, _ReducePrefixUnaryExprToMulExpr}
	_ReduceOpToMulExprAction                                        = &_Action{_ReduceAction, 0, _ReduceOpToMulExpr}
	_ReduceAddToAddOpAction                                         = &_Action{_ReduceAction, 0, _ReduceAddToAddOp}
	_ReduceSubToAddOpAction                                         = &_Action{_ReduceAction, 0, _ReduceSubToAddOp}
	_ReduceBitOrToAddOpAction                                       = &_Action{_ReduceAction, 0, _ReduceBitOrToAddOp}
	_ReduceBitXorToAddOpAction                                      = &_Action{_ReduceAction, 0, _ReduceBitXorToAddOp}
	_ReduceMulExprToAddExprAction                                   = &_Action{_ReduceAction, 0, _ReduceMulExprToAddExpr}
	_ReduceOpToAddExprAction                                        = &_Action{_ReduceAction, 0, _ReduceOpToAddExpr}
	_ReduceEqualToCmpOpAction                                       = &_Action{_ReduceAction, 0, _ReduceEqualToCmpOp}
	_ReduceNotEqualToCmpOpAction                                    = &_Action{_ReduceAction, 0, _ReduceNotEqualToCmpOp}
	_ReduceLessToCmpOpAction                                        = &_Action{_ReduceAction, 0, _ReduceLessToCmpOp}
	_ReduceLessOrEqualToCmpOpAction                                 = &_Action{_ReduceAction, 0, _ReduceLessOrEqualToCmpOp}
	_ReduceGreaterToCmpOpAction                                     = &_Action{_ReduceAction, 0, _ReduceGreaterToCmpOp}
	_ReduceGreaterOrEqualToCmpOpAction                              = &_Action{_ReduceAction, 0, _ReduceGreaterOrEqualToCmpOp}
	_ReduceAddExprToCmpExprAction                                   = &_Action{_ReduceAction, 0, _ReduceAddExprToCmpExpr}
	_ReduceOpToCmpExprAction                                        = &_Action{_ReduceAction, 0, _ReduceOpToCmpExpr}
	_ReduceCmpExprToAndExprAction                                   = &_Action{_ReduceAction, 0, _ReduceCmpExprToAndExpr}
	_ReduceOpToAndExprAction                                        = &_Action{_ReduceAction, 0, _ReduceOpToAndExpr}
	_ReduceAndExprToOrExprAction                                    = &_Action{_ReduceAction, 0, _ReduceAndExprToOrExpr}
	_ReduceOpToOrExprAction                                         = &_Action{_ReduceAction, 0, _ReduceOpToOrExpr}
	_ReduceNamedToAtomTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceNamedToAtomType}
	_ReduceExplicitStructDefToAtomTypeAction                        = &_Action{_ReduceAction, 0, _ReduceExplicitStructDefToAtomType}
	_ReduceImplicitStructDefToAtomTypeAction                        = &_Action{_ReduceAction, 0, _ReduceImplicitStructDefToAtomType}
	_ReduceExplicitEnumDefToAtomTypeAction                          = &_Action{_ReduceAction, 0, _ReduceExplicitEnumDefToAtomType}
	_ReduceImplicitEnumDefToAtomTypeAction                          = &_Action{_ReduceAction, 0, _ReduceImplicitEnumDefToAtomType}
	_ReduceTraitDefToAtomTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceTraitDefToAtomType}
	_ReduceAtomTypeToTraitableTypeAction                            = &_Action{_ReduceAction, 0, _ReduceAtomTypeToTraitableType}
	_ReducePublicMethodsTraitToTraitableTypeAction                  = &_Action{_ReduceAction, 0, _ReducePublicMethodsTraitToTraitableType}
	_ReducePublicTraitToTraitableTypeAction                         = &_Action{_ReduceAction, 0, _ReducePublicTraitToTraitableType}
	_ReduceTraitableTypeToTraitAlgebraTypeAction                    = &_Action{_ReduceAction, 0, _ReduceTraitableTypeToTraitAlgebraType}
	_ReduceIntersectToTraitAlgebraTypeAction                        = &_Action{_ReduceAction, 0, _ReduceIntersectToTraitAlgebraType}
	_ReduceUnionToTraitAlgebraTypeAction                            = &_Action{_ReduceAction, 0, _ReduceUnionToTraitAlgebraType}
	_ReduceDifferenceToTraitAlgebraTypeAction                       = &_Action{_ReduceAction, 0, _ReduceDifferenceToTraitAlgebraType}
	_ReduceInferredToValueTypeAction                                = &_Action{_ReduceAction, 0, _ReduceInferredToValueType}
	_ReduceTraitAlgebraTypeToValueTypeAction                        = &_Action{_ReduceAction, 0, _ReduceTraitAlgebraTypeToValueType}
	_ReduceReferenceToValueTypeAction                               = &_Action{_ReduceAction, 0, _ReduceReferenceToValueType}
	_ReduceFuncTypeToValueTypeAction                                = &_Action{_ReduceAction, 0, _ReduceFuncTypeToValueType}
	_ReduceDefinitionToTypeDefAction                                = &_Action{_ReduceAction, 0, _ReduceDefinitionToTypeDef}
	_ReduceConstrainedDefToTypeDefAction                            = &_Action{_ReduceAction, 0, _ReduceConstrainedDefToTypeDef}
	_ReduceAliasToTypeDefAction                                     = &_Action{_ReduceAction, 0, _ReduceAliasToTypeDef}
	_ReduceUnconstrainedToGenericParameterDefAction                 = &_Action{_ReduceAction, 0, _ReduceUnconstrainedToGenericParameterDef}
	_ReduceConstrainedToGenericParameterDefAction                   = &_Action{_ReduceAction, 0, _ReduceConstrainedToGenericParameterDef}
	_ReduceGenericParameterDefToGenericParameterDefsAction          = &_Action{_ReduceAction, 0, _ReduceGenericParameterDefToGenericParameterDefs}
	_ReduceAddToGenericParameterDefsAction                          = &_Action{_ReduceAction, 0, _ReduceAddToGenericParameterDefs}
	_ReduceGenericParameterDefsToOptionalGenericParameterDefsAction = &_Action{_ReduceAction, 0, _ReduceGenericParameterDefsToOptionalGenericParameterDefs}
	_ReduceNilToOptionalGenericParameterDefsAction                  = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericParameterDefs}
	_ReduceGenericToOptionalGenericParametersAction                 = &_Action{_ReduceAction, 0, _ReduceGenericToOptionalGenericParameters}
	_ReduceNilToOptionalGenericParametersAction                     = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericParameters}
	_ReduceExplicitToFieldDefAction                                 = &_Action{_ReduceAction, 0, _ReduceExplicitToFieldDef}
	_ReduceImplicitToFieldDefAction                                 = &_Action{_ReduceAction, 0, _ReduceImplicitToFieldDef}
	_ReduceUnsafeStatementToFieldDefAction                          = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToFieldDef}
	_ReduceFieldDefToImplicitFieldDefsAction                        = &_Action{_ReduceAction, 0, _ReduceFieldDefToImplicitFieldDefs}
	_ReduceAddToImplicitFieldDefsAction                             = &_Action{_ReduceAction, 0, _ReduceAddToImplicitFieldDefs}
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction       = &_Action{_ReduceAction, 0, _ReduceImplicitFieldDefsToOptionalImplicitFieldDefs}
	_ReduceNilToOptionalImplicitFieldDefsAction                     = &_Action{_ReduceAction, 0, _ReduceNilToOptionalImplicitFieldDefs}
	_ReduceToImplicitStructDefAction                                = &_Action{_ReduceAction, 0, _ReduceToImplicitStructDef}
	_ReduceFieldDefToExplicitFieldDefsAction                        = &_Action{_ReduceAction, 0, _ReduceFieldDefToExplicitFieldDefs}
	_ReduceImplicitToExplicitFieldDefsAction                        = &_Action{_ReduceAction, 0, _ReduceImplicitToExplicitFieldDefs}
	_ReduceExplicitToExplicitFieldDefsAction                        = &_Action{_ReduceAction, 0, _ReduceExplicitToExplicitFieldDefs}
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction       = &_Action{_ReduceAction, 0, _ReduceExplicitFieldDefsToOptionalExplicitFieldDefs}
	_ReduceNilToOptionalExplicitFieldDefsAction                     = &_Action{_ReduceAction, 0, _ReduceNilToOptionalExplicitFieldDefs}
	_ReduceToExplicitStructDefAction                                = &_Action{_ReduceAction, 0, _ReduceToExplicitStructDef}
	_ReduceFieldDefToEnumValueDefAction                             = &_Action{_ReduceAction, 0, _ReduceFieldDefToEnumValueDef}
	_ReduceDefaultToEnumValueDefAction                              = &_Action{_ReduceAction, 0, _ReduceDefaultToEnumValueDef}
	_ReducePairToImplicitEnumValueDefsAction                        = &_Action{_ReduceAction, 0, _ReducePairToImplicitEnumValueDefs}
	_ReduceAddToImplicitEnumValueDefsAction                         = &_Action{_ReduceAction, 0, _ReduceAddToImplicitEnumValueDefs}
	_ReduceToImplicitEnumDefAction                                  = &_Action{_ReduceAction, 0, _ReduceToImplicitEnumDef}
	_ReduceExplicitPairToExplicitEnumValueDefsAction                = &_Action{_ReduceAction, 0, _ReduceExplicitPairToExplicitEnumValueDefs}
	_ReduceImplicitPairToExplicitEnumValueDefsAction                = &_Action{_ReduceAction, 0, _ReduceImplicitPairToExplicitEnumValueDefs}
	_ReduceExplicitAddToExplicitEnumValueDefsAction                 = &_Action{_ReduceAction, 0, _ReduceExplicitAddToExplicitEnumValueDefs}
	_ReduceImplicitAddToExplicitEnumValueDefsAction                 = &_Action{_ReduceAction, 0, _ReduceImplicitAddToExplicitEnumValueDefs}
	_ReduceToExplicitEnumDefAction                                  = &_Action{_ReduceAction, 0, _ReduceToExplicitEnumDef}
	_ReduceFieldDefToTraitPropertyAction                            = &_Action{_ReduceAction, 0, _ReduceFieldDefToTraitProperty}
	_ReduceMethodSignatureToTraitPropertyAction                     = &_Action{_ReduceAction, 0, _ReduceMethodSignatureToTraitProperty}
	_ReduceTraitPropertyToTraitPropertiesAction                     = &_Action{_ReduceAction, 0, _ReduceTraitPropertyToTraitProperties}
	_ReduceImplicitToTraitPropertiesAction                          = &_Action{_ReduceAction, 0, _ReduceImplicitToTraitProperties}
	_ReduceExplicitToTraitPropertiesAction                          = &_Action{_ReduceAction, 0, _ReduceExplicitToTraitProperties}
	_ReduceTraitPropertiesToOptionalTraitPropertiesAction           = &_Action{_ReduceAction, 0, _ReduceTraitPropertiesToOptionalTraitProperties}
	_ReduceNilToOptionalTraitPropertiesAction                       = &_Action{_ReduceAction, 0, _ReduceNilToOptionalTraitProperties}
	_ReduceToTraitDefAction                                         = &_Action{_ReduceAction, 0, _ReduceToTraitDef}
	_ReduceValueTypeToReturnTypeAction                              = &_Action{_ReduceAction, 0, _ReduceValueTypeToReturnType}
	_ReduceNilToReturnTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceNilToReturnType}
	_ReduceArgToParameterDeclAction                                 = &_Action{_ReduceAction, 0, _ReduceArgToParameterDecl}
	_ReduceVarargToParameterDeclAction                              = &_Action{_ReduceAction, 0, _ReduceVarargToParameterDecl}
	_ReduceUnamedToParameterDeclAction                              = &_Action{_ReduceAction, 0, _ReduceUnamedToParameterDecl}
	_ReduceUnnamedVarargToParameterDeclAction                       = &_Action{_ReduceAction, 0, _ReduceUnnamedVarargToParameterDecl}
	_ReduceParameterDeclToParameterDeclsAction                      = &_Action{_ReduceAction, 0, _ReduceParameterDeclToParameterDecls}
	_ReduceAddToParameterDeclsAction                                = &_Action{_ReduceAction, 0, _ReduceAddToParameterDecls}
	_ReduceParameterDeclsToOptionalParameterDeclsAction             = &_Action{_ReduceAction, 0, _ReduceParameterDeclsToOptionalParameterDecls}
	_ReduceNilToOptionalParameterDeclsAction                        = &_Action{_ReduceAction, 0, _ReduceNilToOptionalParameterDecls}
	_ReduceToFuncTypeAction                                         = &_Action{_ReduceAction, 0, _ReduceToFuncType}
	_ReduceToMethodSignatureAction                                  = &_Action{_ReduceAction, 0, _ReduceToMethodSignature}
	_ReduceArgToParameterDefAction                                  = &_Action{_ReduceAction, 0, _ReduceArgToParameterDef}
	_ReduceVarargToParameterDefAction                               = &_Action{_ReduceAction, 0, _ReduceVarargToParameterDef}
	_ReduceParameterDefToParameterDefsAction                        = &_Action{_ReduceAction, 0, _ReduceParameterDefToParameterDefs}
	_ReduceAddToParameterDefsAction                                 = &_Action{_ReduceAction, 0, _ReduceAddToParameterDefs}
	_ReduceParameterDefsToOptionalParameterDefsAction               = &_Action{_ReduceAction, 0, _ReduceParameterDefsToOptionalParameterDefs}
	_ReduceNilToOptionalParameterDefsAction                         = &_Action{_ReduceAction, 0, _ReduceNilToOptionalParameterDefs}
	_ReduceReceiverToOptionalReceiverAction                         = &_Action{_ReduceAction, 0, _ReduceReceiverToOptionalReceiver}
	_ReduceNilToOptionalReceiverAction                              = &_Action{_ReduceAction, 0, _ReduceNilToOptionalReceiver}
	_ReduceToNamedFuncDefAction                                     = &_Action{_ReduceAction, 0, _ReduceToNamedFuncDef}
	_ReduceToAnonymousFuncExprAction                                = &_Action{_ReduceAction, 0, _ReduceToAnonymousFuncExpr}
	_ReduceNoSpecToPackageDefAction                                 = &_Action{_ReduceAction, 0, _ReduceNoSpecToPackageDef}
	_ReduceWithSpecToPackageDefAction                               = &_Action{_ReduceAction, 0, _ReduceWithSpecToPackageDef}
	_ReduceToPackageStatementBodyAction                             = &_Action{_ReduceAction, 0, _ReduceToPackageStatementBody}
	_ReduceImplicitToPackageStatementAction                         = &_Action{_ReduceAction, 0, _ReduceImplicitToPackageStatement}
	_ReduceExplicitToPackageStatementAction                         = &_Action{_ReduceAction, 0, _ReduceExplicitToPackageStatement}
	_ReduceEmptyListToPackageStatementsAction                       = &_Action{_ReduceAction, 0, _ReduceEmptyListToPackageStatements}
	_ReduceAddToPackageStatementsAction                             = &_Action{_ReduceAction, 0, _ReduceAddToPackageStatements}
	_ReduceSpacesToLexInternalTokensAction                          = &_Action{_ReduceAction, 0, _ReduceSpacesToLexInternalTokens}
	_ReduceCommentToLexInternalTokensAction                         = &_Action{_ReduceAction, 0, _ReduceCommentToLexInternalTokens}
)

var _ActionTable = _ActionTableType{
	{_State7, _EndMarker}:                         &_Action{_AcceptAction, 0, 0},
	{_State8, _EndMarker}:                         &_Action{_AcceptAction, 0, 0},
	{_State9, _EndMarker}:                         &_Action{_AcceptAction, 0, 0},
	{_State10, _EndMarker}:                        &_Action{_AcceptAction, 0, 0},
	{_State11, _EndMarker}:                        &_Action{_AcceptAction, 0, 0},
	{_State12, _EndMarker}:                        &_Action{_AcceptAction, 0, 0},
	{_State1, PackageToken}:                       _GotoState13Action,
	{_State1, PackageDefType}:                     _GotoState7Action,
	{_State2, TypeToken}:                          _GotoState14Action,
	{_State2, TypeDefType}:                        _GotoState8Action,
	{_State3, TraitToken}:                         _GotoState15Action,
	{_State3, TraitDefType}:                       _GotoState9Action,
	{_State4, FuncToken}:                          _GotoState16Action,
	{_State4, NamedFuncDefType}:                   _GotoState10Action,
	{_State5, IntegerLiteralToken}:                _GotoState23Action,
	{_State5, FloatLiteralToken}:                  _GotoState20Action,
	{_State5, RuneLiteralToken}:                   _GotoState29Action,
	{_State5, StringLiteralToken}:                 _GotoState30Action,
	{_State5, IdentifierToken}:                    _GotoState22Action,
	{_State5, TrueToken}:                          _GotoState33Action,
	{_State5, FalseToken}:                         _GotoState19Action,
	{_State5, StructToken}:                        _GotoState31Action,
	{_State5, FuncToken}:                          _GotoState21Action,
	{_State5, LabelDeclToken}:                     _GotoState24Action,
	{_State5, LparenToken}:                        _GotoState26Action,
	{_State5, NotToken}:                           _GotoState28Action,
	{_State5, SubToken}:                           _GotoState32Action,
	{_State5, MulToken}:                           _GotoState27Action,
	{_State5, BitNegToken}:                        _GotoState18Action,
	{_State5, BitAndToken}:                        _GotoState17Action,
	{_State5, LexErrorToken}:                      _GotoState25Action,
	{_State5, ExpressionType}:                     _GotoState11Action,
	{_State5, OptionalLabelDeclType}:              _GotoState46Action,
	{_State5, SequenceExprType}:                   _GotoState51Action,
	{_State5, BlockExprType}:                      _GotoState40Action,
	{_State5, CallExprType}:                       _GotoState41Action,
	{_State5, AtomExprType}:                       _GotoState39Action,
	{_State5, LiteralType}:                        _GotoState44Action,
	{_State5, AnonymousStructExprType}:            _GotoState38Action,
	{_State5, AccessExprType}:                     _GotoState34Action,
	{_State5, PostfixUnaryExprType}:               _GotoState48Action,
	{_State5, PrefixUnaryOpType}:                  _GotoState50Action,
	{_State5, PrefixUnaryExprType}:                _GotoState49Action,
	{_State5, MulExprType}:                        _GotoState45Action,
	{_State5, AddExprType}:                        _GotoState35Action,
	{_State5, CmpExprType}:                        _GotoState42Action,
	{_State5, AndExprType}:                        _GotoState36Action,
	{_State5, OrExprType}:                         _GotoState47Action,
	{_State5, ExplicitStructDefType}:              _GotoState43Action,
	{_State5, AnonymousFuncExprType}:              _GotoState37Action,
	{_State6, SpacesToken}:                        _GotoState53Action,
	{_State6, CommentToken}:                       _GotoState52Action,
	{_State6, LexInternalTokensType}:              _GotoState12Action,
	{_State13, IdentifierToken}:                   _GotoState54Action,
	{_State14, IdentifierToken}:                   _GotoState55Action,
	{_State15, LparenToken}:                       _GotoState56Action,
	{_State16, LparenToken}:                       _GotoState57Action,
	{_State16, OptionalReceiverType}:              _GotoState58Action,
	{_State21, LparenToken}:                       _GotoState59Action,
	{_State26, IntegerLiteralToken}:               _GotoState23Action,
	{_State26, FloatLiteralToken}:                 _GotoState20Action,
	{_State26, RuneLiteralToken}:                  _GotoState29Action,
	{_State26, StringLiteralToken}:                _GotoState30Action,
	{_State26, IdentifierToken}:                   _GotoState60Action,
	{_State26, TrueToken}:                         _GotoState33Action,
	{_State26, FalseToken}:                        _GotoState19Action,
	{_State26, StructToken}:                       _GotoState31Action,
	{_State26, FuncToken}:                         _GotoState21Action,
	{_State26, LabelDeclToken}:                    _GotoState24Action,
	{_State26, LparenToken}:                       _GotoState26Action,
	{_State26, NotToken}:                          _GotoState28Action,
	{_State26, SubToken}:                          _GotoState32Action,
	{_State26, MulToken}:                          _GotoState27Action,
	{_State26, BitNegToken}:                       _GotoState18Action,
	{_State26, BitAndToken}:                       _GotoState17Action,
	{_State26, LexErrorToken}:                     _GotoState25Action,
	{_State26, ExpressionType}:                    _GotoState64Action,
	{_State26, OptionalLabelDeclType}:             _GotoState46Action,
	{_State26, SequenceExprType}:                  _GotoState51Action,
	{_State26, BlockExprType}:                     _GotoState40Action,
	{_State26, CallExprType}:                      _GotoState41Action,
	{_State26, ArgumentsType}:                     _GotoState62Action,
	{_State26, ArgumentType}:                      _GotoState61Action,
	{_State26, ColonExpressionsType}:              _GotoState63Action,
	{_State26, OptionalExpressionType}:            _GotoState65Action,
	{_State26, AtomExprType}:                      _GotoState39Action,
	{_State26, LiteralType}:                       _GotoState44Action,
	{_State26, AnonymousStructExprType}:           _GotoState38Action,
	{_State26, AccessExprType}:                    _GotoState34Action,
	{_State26, PostfixUnaryExprType}:              _GotoState48Action,
	{_State26, PrefixUnaryOpType}:                 _GotoState50Action,
	{_State26, PrefixUnaryExprType}:               _GotoState49Action,
	{_State26, MulExprType}:                       _GotoState45Action,
	{_State26, AddExprType}:                       _GotoState35Action,
	{_State26, CmpExprType}:                       _GotoState42Action,
	{_State26, AndExprType}:                       _GotoState36Action,
	{_State26, OrExprType}:                        _GotoState47Action,
	{_State26, ExplicitStructDefType}:             _GotoState43Action,
	{_State26, AnonymousFuncExprType}:             _GotoState37Action,
	{_State31, LparenToken}:                       _GotoState66Action,
	{_State34, LbracketToken}:                     _GotoState69Action,
	{_State34, DotToken}:                          _GotoState68Action,
	{_State34, QuestionToken}:                     _GotoState70Action,
	{_State34, DollarLbracketToken}:               _GotoState67Action,
	{_State34, OptionalGenericBindingType}:        _GotoState71Action,
	{_State35, AddToken}:                          _GotoState72Action,
	{_State35, SubToken}:                          _GotoState75Action,
	{_State35, BitXorToken}:                       _GotoState74Action,
	{_State35, BitOrToken}:                        _GotoState73Action,
	{_State35, AddOpType}:                         _GotoState76Action,
	{_State36, AndToken}:                          _GotoState77Action,
	{_State42, EqualToken}:                        _GotoState78Action,
	{_State42, NotEqualToken}:                     _GotoState83Action,
	{_State42, LessToken}:                         _GotoState81Action,
	{_State42, LessOrEqualToken}:                  _GotoState82Action,
	{_State42, GreaterToken}:                      _GotoState79Action,
	{_State42, GreaterOrEqualToken}:               _GotoState80Action,
	{_State42, CmpOpType}:                         _GotoState84Action,
	{_State43, LparenToken}:                       _GotoState85Action,
	{_State45, MulToken}:                          _GotoState91Action,
	{_State45, DivToken}:                          _GotoState89Action,
	{_State45, ModToken}:                          _GotoState90Action,
	{_State45, BitAndToken}:                       _GotoState86Action,
	{_State45, BitLshiftToken}:                    _GotoState87Action,
	{_State45, BitRshiftToken}:                    _GotoState88Action,
	{_State45, MulOpType}:                         _GotoState92Action,
	{_State46, IfToken}:                           _GotoState95Action,
	{_State46, SwitchToken}:                       _GotoState97Action,
	{_State46, ForToken}:                          _GotoState94Action,
	{_State46, DoToken}:                           _GotoState93Action,
	{_State46, LbraceToken}:                       _GotoState96Action,
	{_State46, IfExprType}:                        _GotoState99Action,
	{_State46, SwitchExprType}:                    _GotoState101Action,
	{_State46, LoopExprType}:                      _GotoState100Action,
	{_State46, BlockBodyType}:                     _GotoState98Action,
	{_State47, OrToken}:                           _GotoState102Action,
	{_State50, IntegerLiteralToken}:               _GotoState23Action,
	{_State50, FloatLiteralToken}:                 _GotoState20Action,
	{_State50, RuneLiteralToken}:                  _GotoState29Action,
	{_State50, StringLiteralToken}:                _GotoState30Action,
	{_State50, IdentifierToken}:                   _GotoState22Action,
	{_State50, TrueToken}:                         _GotoState33Action,
	{_State50, FalseToken}:                        _GotoState19Action,
	{_State50, StructToken}:                       _GotoState31Action,
	{_State50, FuncToken}:                         _GotoState21Action,
	{_State50, LabelDeclToken}:                    _GotoState24Action,
	{_State50, LparenToken}:                       _GotoState26Action,
	{_State50, NotToken}:                          _GotoState28Action,
	{_State50, SubToken}:                          _GotoState32Action,
	{_State50, MulToken}:                          _GotoState27Action,
	{_State50, BitNegToken}:                       _GotoState18Action,
	{_State50, BitAndToken}:                       _GotoState17Action,
	{_State50, LexErrorToken}:                     _GotoState25Action,
	{_State50, OptionalLabelDeclType}:             _GotoState103Action,
	{_State50, BlockExprType}:                     _GotoState40Action,
	{_State50, CallExprType}:                      _GotoState41Action,
	{_State50, AtomExprType}:                      _GotoState39Action,
	{_State50, LiteralType}:                       _GotoState44Action,
	{_State50, AnonymousStructExprType}:           _GotoState38Action,
	{_State50, AccessExprType}:                    _GotoState34Action,
	{_State50, PostfixUnaryExprType}:              _GotoState48Action,
	{_State50, PrefixUnaryOpType}:                 _GotoState50Action,
	{_State50, PrefixUnaryExprType}:               _GotoState104Action,
	{_State50, ExplicitStructDefType}:             _GotoState43Action,
	{_State50, AnonymousFuncExprType}:             _GotoState37Action,
	{_State54, LparenToken}:                       _GotoState105Action,
	{_State55, DollarLbracketToken}:               _GotoState106Action,
	{_State55, EqualToken}:                        _GotoState107Action,
	{_State55, OptionalGenericParametersType}:     _GotoState108Action,
	{_State56, IdentifierToken}:                   _GotoState113Action,
	{_State56, UnsafeToken}:                       _GotoState117Action,
	{_State56, StructToken}:                       _GotoState31Action,
	{_State56, EnumToken}:                         _GotoState111Action,
	{_State56, TraitToken}:                        _GotoState15Action,
	{_State56, FuncToken}:                         _GotoState112Action,
	{_State56, LparenToken}:                       _GotoState114Action,
	{_State56, QuestionToken}:                     _GotoState115Action,
	{_State56, TildeTildeToken}:                   _GotoState116Action,
	{_State56, BitNegToken}:                       _GotoState110Action,
	{_State56, BitAndToken}:                       _GotoState109Action,
	{_State56, UnsafeStatementType}:               _GotoState132Action,
	{_State56, AtomTypeType}:                      _GotoState118Action,
	{_State56, TraitableTypeType}:                 _GotoState131Action,
	{_State56, TraitAlgebraTypeType}:              _GotoState127Action,
	{_State56, ValueTypeType}:                     _GotoState133Action,
	{_State56, FieldDefType}:                      _GotoState121Action,
	{_State56, ImplicitStructDefType}:             _GotoState124Action,
	{_State56, ExplicitStructDefType}:             _GotoState120Action,
	{_State56, ImplicitEnumDefType}:               _GotoState123Action,
	{_State56, ExplicitEnumDefType}:               _GotoState119Action,
	{_State56, TraitPropertyType}:                 _GotoState130Action,
	{_State56, TraitPropertiesType}:               _GotoState129Action,
	{_State56, OptionalTraitPropertiesType}:       _GotoState126Action,
	{_State56, TraitDefType}:                      _GotoState128Action,
	{_State56, FuncTypeType}:                      _GotoState122Action,
	{_State56, MethodSignatureType}:               _GotoState125Action,
	{_State57, IdentifierToken}:                   _GotoState134Action,
	{_State57, ParameterDefType}:                  _GotoState135Action,
	{_State58, IdentifierToken}:                   _GotoState136Action,
	{_State59, IdentifierToken}:                   _GotoState134Action,
	{_State59, ParameterDefType}:                  _GotoState138Action,
	{_State59, ParameterDefsType}:                 _GotoState139Action,
	{_State59, OptionalParameterDefsType}:         _GotoState137Action,
	{_State60, AssignToken}:                       _GotoState140Action,
	{_State62, RparenToken}:                       _GotoState142Action,
	{_State62, CommaToken}:                        _GotoState141Action,
	{_State63, ColonToken}:                        _GotoState143Action,
	{_State65, ColonToken}:                        _GotoState144Action,
	{_State66, IdentifierToken}:                   _GotoState113Action,
	{_State66, UnsafeToken}:                       _GotoState117Action,
	{_State66, StructToken}:                       _GotoState31Action,
	{_State66, EnumToken}:                         _GotoState111Action,
	{_State66, TraitToken}:                        _GotoState15Action,
	{_State66, FuncToken}:                         _GotoState145Action,
	{_State66, LparenToken}:                       _GotoState114Action,
	{_State66, QuestionToken}:                     _GotoState115Action,
	{_State66, TildeTildeToken}:                   _GotoState116Action,
	{_State66, BitNegToken}:                       _GotoState110Action,
	{_State66, BitAndToken}:                       _GotoState109Action,
	{_State66, UnsafeStatementType}:               _GotoState132Action,
	{_State66, AtomTypeType}:                      _GotoState118Action,
	{_State66, TraitableTypeType}:                 _GotoState131Action,
	{_State66, TraitAlgebraTypeType}:              _GotoState127Action,
	{_State66, ValueTypeType}:                     _GotoState133Action,
	{_State66, FieldDefType}:                      _GotoState147Action,
	{_State66, ImplicitStructDefType}:             _GotoState124Action,
	{_State66, ExplicitFieldDefsType}:             _GotoState146Action,
	{_State66, OptionalExplicitFieldDefsType}:     _GotoState148Action,
	{_State66, ExplicitStructDefType}:             _GotoState120Action,
	{_State66, ImplicitEnumDefType}:               _GotoState123Action,
	{_State66, ExplicitEnumDefType}:               _GotoState119Action,
	{_State66, TraitDefType}:                      _GotoState128Action,
	{_State66, FuncTypeType}:                      _GotoState122Action,
	{_State67, IdentifierToken}:                   _GotoState149Action,
	{_State67, StructToken}:                       _GotoState31Action,
	{_State67, EnumToken}:                         _GotoState111Action,
	{_State67, TraitToken}:                        _GotoState15Action,
	{_State67, FuncToken}:                         _GotoState145Action,
	{_State67, LparenToken}:                       _GotoState114Action,
	{_State67, QuestionToken}:                     _GotoState115Action,
	{_State67, TildeTildeToken}:                   _GotoState116Action,
	{_State67, BitNegToken}:                       _GotoState110Action,
	{_State67, BitAndToken}:                       _GotoState109Action,
	{_State67, OptionalGenericArgumentsType}:      _GotoState151Action,
	{_State67, GenericArgumentsType}:              _GotoState150Action,
	{_State67, AtomTypeType}:                      _GotoState118Action,
	{_State67, TraitableTypeType}:                 _GotoState131Action,
	{_State67, TraitAlgebraTypeType}:              _GotoState127Action,
	{_State67, ValueTypeType}:                     _GotoState152Action,
	{_State67, ImplicitStructDefType}:             _GotoState124Action,
	{_State67, ExplicitStructDefType}:             _GotoState120Action,
	{_State67, ImplicitEnumDefType}:               _GotoState123Action,
	{_State67, ExplicitEnumDefType}:               _GotoState119Action,
	{_State67, TraitDefType}:                      _GotoState128Action,
	{_State67, FuncTypeType}:                      _GotoState122Action,
	{_State68, IdentifierToken}:                   _GotoState153Action,
	{_State69, IntegerLiteralToken}:               _GotoState23Action,
	{_State69, FloatLiteralToken}:                 _GotoState20Action,
	{_State69, RuneLiteralToken}:                  _GotoState29Action,
	{_State69, StringLiteralToken}:                _GotoState30Action,
	{_State69, IdentifierToken}:                   _GotoState60Action,
	{_State69, TrueToken}:                         _GotoState33Action,
	{_State69, FalseToken}:                        _GotoState19Action,
	{_State69, StructToken}:                       _GotoState31Action,
	{_State69, FuncToken}:                         _GotoState21Action,
	{_State69, LabelDeclToken}:                    _GotoState24Action,
	{_State69, LparenToken}:                       _GotoState26Action,
	{_State69, NotToken}:                          _GotoState28Action,
	{_State69, SubToken}:                          _GotoState32Action,
	{_State69, MulToken}:                          _GotoState27Action,
	{_State69, BitNegToken}:                       _GotoState18Action,
	{_State69, BitAndToken}:                       _GotoState17Action,
	{_State69, LexErrorToken}:                     _GotoState25Action,
	{_State69, ExpressionType}:                    _GotoState64Action,
	{_State69, OptionalLabelDeclType}:             _GotoState46Action,
	{_State69, SequenceExprType}:                  _GotoState51Action,
	{_State69, BlockExprType}:                     _GotoState40Action,
	{_State69, CallExprType}:                      _GotoState41Action,
	{_State69, ArgumentType}:                      _GotoState154Action,
	{_State69, ColonExpressionsType}:              _GotoState63Action,
	{_State69, OptionalExpressionType}:            _GotoState65Action,
	{_State69, AtomExprType}:                      _GotoState39Action,
	{_State69, LiteralType}:                       _GotoState44Action,
	{_State69, AnonymousStructExprType}:           _GotoState38Action,
	{_State69, AccessExprType}:                    _GotoState34Action,
	{_State69, PostfixUnaryExprType}:              _GotoState48Action,
	{_State69, PrefixUnaryOpType}:                 _GotoState50Action,
	{_State69, PrefixUnaryExprType}:               _GotoState49Action,
	{_State69, MulExprType}:                       _GotoState45Action,
	{_State69, AddExprType}:                       _GotoState35Action,
	{_State69, CmpExprType}:                       _GotoState42Action,
	{_State69, AndExprType}:                       _GotoState36Action,
	{_State69, OrExprType}:                        _GotoState47Action,
	{_State69, ExplicitStructDefType}:             _GotoState43Action,
	{_State69, AnonymousFuncExprType}:             _GotoState37Action,
	{_State71, LparenToken}:                       _GotoState155Action,
	{_State76, IntegerLiteralToken}:               _GotoState23Action,
	{_State76, FloatLiteralToken}:                 _GotoState20Action,
	{_State76, RuneLiteralToken}:                  _GotoState29Action,
	{_State76, StringLiteralToken}:                _GotoState30Action,
	{_State76, IdentifierToken}:                   _GotoState22Action,
	{_State76, TrueToken}:                         _GotoState33Action,
	{_State76, FalseToken}:                        _GotoState19Action,
	{_State76, StructToken}:                       _GotoState31Action,
	{_State76, FuncToken}:                         _GotoState21Action,
	{_State76, LabelDeclToken}:                    _GotoState24Action,
	{_State76, LparenToken}:                       _GotoState26Action,
	{_State76, NotToken}:                          _GotoState28Action,
	{_State76, SubToken}:                          _GotoState32Action,
	{_State76, MulToken}:                          _GotoState27Action,
	{_State76, BitNegToken}:                       _GotoState18Action,
	{_State76, BitAndToken}:                       _GotoState17Action,
	{_State76, LexErrorToken}:                     _GotoState25Action,
	{_State76, OptionalLabelDeclType}:             _GotoState103Action,
	{_State76, BlockExprType}:                     _GotoState40Action,
	{_State76, CallExprType}:                      _GotoState41Action,
	{_State76, AtomExprType}:                      _GotoState39Action,
	{_State76, LiteralType}:                       _GotoState44Action,
	{_State76, AnonymousStructExprType}:           _GotoState38Action,
	{_State76, AccessExprType}:                    _GotoState34Action,
	{_State76, PostfixUnaryExprType}:              _GotoState48Action,
	{_State76, PrefixUnaryOpType}:                 _GotoState50Action,
	{_State76, PrefixUnaryExprType}:               _GotoState49Action,
	{_State76, MulExprType}:                       _GotoState156Action,
	{_State76, ExplicitStructDefType}:             _GotoState43Action,
	{_State76, AnonymousFuncExprType}:             _GotoState37Action,
	{_State77, IntegerLiteralToken}:               _GotoState23Action,
	{_State77, FloatLiteralToken}:                 _GotoState20Action,
	{_State77, RuneLiteralToken}:                  _GotoState29Action,
	{_State77, StringLiteralToken}:                _GotoState30Action,
	{_State77, IdentifierToken}:                   _GotoState22Action,
	{_State77, TrueToken}:                         _GotoState33Action,
	{_State77, FalseToken}:                        _GotoState19Action,
	{_State77, StructToken}:                       _GotoState31Action,
	{_State77, FuncToken}:                         _GotoState21Action,
	{_State77, LabelDeclToken}:                    _GotoState24Action,
	{_State77, LparenToken}:                       _GotoState26Action,
	{_State77, NotToken}:                          _GotoState28Action,
	{_State77, SubToken}:                          _GotoState32Action,
	{_State77, MulToken}:                          _GotoState27Action,
	{_State77, BitNegToken}:                       _GotoState18Action,
	{_State77, BitAndToken}:                       _GotoState17Action,
	{_State77, LexErrorToken}:                     _GotoState25Action,
	{_State77, OptionalLabelDeclType}:             _GotoState103Action,
	{_State77, BlockExprType}:                     _GotoState40Action,
	{_State77, CallExprType}:                      _GotoState41Action,
	{_State77, AtomExprType}:                      _GotoState39Action,
	{_State77, LiteralType}:                       _GotoState44Action,
	{_State77, AnonymousStructExprType}:           _GotoState38Action,
	{_State77, AccessExprType}:                    _GotoState34Action,
	{_State77, PostfixUnaryExprType}:              _GotoState48Action,
	{_State77, PrefixUnaryOpType}:                 _GotoState50Action,
	{_State77, PrefixUnaryExprType}:               _GotoState49Action,
	{_State77, MulExprType}:                       _GotoState45Action,
	{_State77, AddExprType}:                       _GotoState35Action,
	{_State77, CmpExprType}:                       _GotoState157Action,
	{_State77, ExplicitStructDefType}:             _GotoState43Action,
	{_State77, AnonymousFuncExprType}:             _GotoState37Action,
	{_State84, IntegerLiteralToken}:               _GotoState23Action,
	{_State84, FloatLiteralToken}:                 _GotoState20Action,
	{_State84, RuneLiteralToken}:                  _GotoState29Action,
	{_State84, StringLiteralToken}:                _GotoState30Action,
	{_State84, IdentifierToken}:                   _GotoState22Action,
	{_State84, TrueToken}:                         _GotoState33Action,
	{_State84, FalseToken}:                        _GotoState19Action,
	{_State84, StructToken}:                       _GotoState31Action,
	{_State84, FuncToken}:                         _GotoState21Action,
	{_State84, LabelDeclToken}:                    _GotoState24Action,
	{_State84, LparenToken}:                       _GotoState26Action,
	{_State84, NotToken}:                          _GotoState28Action,
	{_State84, SubToken}:                          _GotoState32Action,
	{_State84, MulToken}:                          _GotoState27Action,
	{_State84, BitNegToken}:                       _GotoState18Action,
	{_State84, BitAndToken}:                       _GotoState17Action,
	{_State84, LexErrorToken}:                     _GotoState25Action,
	{_State84, OptionalLabelDeclType}:             _GotoState103Action,
	{_State84, BlockExprType}:                     _GotoState40Action,
	{_State84, CallExprType}:                      _GotoState41Action,
	{_State84, AtomExprType}:                      _GotoState39Action,
	{_State84, LiteralType}:                       _GotoState44Action,
	{_State84, AnonymousStructExprType}:           _GotoState38Action,
	{_State84, AccessExprType}:                    _GotoState34Action,
	{_State84, PostfixUnaryExprType}:              _GotoState48Action,
	{_State84, PrefixUnaryOpType}:                 _GotoState50Action,
	{_State84, PrefixUnaryExprType}:               _GotoState49Action,
	{_State84, MulExprType}:                       _GotoState45Action,
	{_State84, AddExprType}:                       _GotoState158Action,
	{_State84, ExplicitStructDefType}:             _GotoState43Action,
	{_State84, AnonymousFuncExprType}:             _GotoState37Action,
	{_State85, IntegerLiteralToken}:               _GotoState23Action,
	{_State85, FloatLiteralToken}:                 _GotoState20Action,
	{_State85, RuneLiteralToken}:                  _GotoState29Action,
	{_State85, StringLiteralToken}:                _GotoState30Action,
	{_State85, IdentifierToken}:                   _GotoState60Action,
	{_State85, TrueToken}:                         _GotoState33Action,
	{_State85, FalseToken}:                        _GotoState19Action,
	{_State85, StructToken}:                       _GotoState31Action,
	{_State85, FuncToken}:                         _GotoState21Action,
	{_State85, LabelDeclToken}:                    _GotoState24Action,
	{_State85, LparenToken}:                       _GotoState26Action,
	{_State85, NotToken}:                          _GotoState28Action,
	{_State85, SubToken}:                          _GotoState32Action,
	{_State85, MulToken}:                          _GotoState27Action,
	{_State85, BitNegToken}:                       _GotoState18Action,
	{_State85, BitAndToken}:                       _GotoState17Action,
	{_State85, LexErrorToken}:                     _GotoState25Action,
	{_State85, ExpressionType}:                    _GotoState64Action,
	{_State85, OptionalLabelDeclType}:             _GotoState46Action,
	{_State85, SequenceExprType}:                  _GotoState51Action,
	{_State85, BlockExprType}:                     _GotoState40Action,
	{_State85, CallExprType}:                      _GotoState41Action,
	{_State85, ArgumentsType}:                     _GotoState159Action,
	{_State85, ArgumentType}:                      _GotoState61Action,
	{_State85, ColonExpressionsType}:              _GotoState63Action,
	{_State85, OptionalExpressionType}:            _GotoState65Action,
	{_State85, AtomExprType}:                      _GotoState39Action,
	{_State85, LiteralType}:                       _GotoState44Action,
	{_State85, AnonymousStructExprType}:           _GotoState38Action,
	{_State85, AccessExprType}:                    _GotoState34Action,
	{_State85, PostfixUnaryExprType}:              _GotoState48Action,
	{_State85, PrefixUnaryOpType}:                 _GotoState50Action,
	{_State85, PrefixUnaryExprType}:               _GotoState49Action,
	{_State85, MulExprType}:                       _GotoState45Action,
	{_State85, AddExprType}:                       _GotoState35Action,
	{_State85, CmpExprType}:                       _GotoState42Action,
	{_State85, AndExprType}:                       _GotoState36Action,
	{_State85, OrExprType}:                        _GotoState47Action,
	{_State85, ExplicitStructDefType}:             _GotoState43Action,
	{_State85, AnonymousFuncExprType}:             _GotoState37Action,
	{_State92, IntegerLiteralToken}:               _GotoState23Action,
	{_State92, FloatLiteralToken}:                 _GotoState20Action,
	{_State92, RuneLiteralToken}:                  _GotoState29Action,
	{_State92, StringLiteralToken}:                _GotoState30Action,
	{_State92, IdentifierToken}:                   _GotoState22Action,
	{_State92, TrueToken}:                         _GotoState33Action,
	{_State92, FalseToken}:                        _GotoState19Action,
	{_State92, StructToken}:                       _GotoState31Action,
	{_State92, FuncToken}:                         _GotoState21Action,
	{_State92, LabelDeclToken}:                    _GotoState24Action,
	{_State92, LparenToken}:                       _GotoState26Action,
	{_State92, NotToken}:                          _GotoState28Action,
	{_State92, SubToken}:                          _GotoState32Action,
	{_State92, MulToken}:                          _GotoState27Action,
	{_State92, BitNegToken}:                       _GotoState18Action,
	{_State92, BitAndToken}:                       _GotoState17Action,
	{_State92, LexErrorToken}:                     _GotoState25Action,
	{_State92, OptionalLabelDeclType}:             _GotoState103Action,
	{_State92, BlockExprType}:                     _GotoState40Action,
	{_State92, CallExprType}:                      _GotoState41Action,
	{_State92, AtomExprType}:                      _GotoState39Action,
	{_State92, LiteralType}:                       _GotoState44Action,
	{_State92, AnonymousStructExprType}:           _GotoState38Action,
	{_State92, AccessExprType}:                    _GotoState34Action,
	{_State92, PostfixUnaryExprType}:              _GotoState48Action,
	{_State92, PrefixUnaryOpType}:                 _GotoState50Action,
	{_State92, PrefixUnaryExprType}:               _GotoState160Action,
	{_State92, ExplicitStructDefType}:             _GotoState43Action,
	{_State92, AnonymousFuncExprType}:             _GotoState37Action,
	{_State93, LbraceToken}:                       _GotoState96Action,
	{_State93, BlockBodyType}:                     _GotoState161Action,
	{_State94, IntegerLiteralToken}:               _GotoState23Action,
	{_State94, FloatLiteralToken}:                 _GotoState20Action,
	{_State94, RuneLiteralToken}:                  _GotoState29Action,
	{_State94, StringLiteralToken}:                _GotoState30Action,
	{_State94, IdentifierToken}:                   _GotoState22Action,
	{_State94, TrueToken}:                         _GotoState33Action,
	{_State94, FalseToken}:                        _GotoState19Action,
	{_State94, InToken}:                           _GotoState162Action,
	{_State94, StructToken}:                       _GotoState31Action,
	{_State94, FuncToken}:                         _GotoState21Action,
	{_State94, LabelDeclToken}:                    _GotoState24Action,
	{_State94, LparenToken}:                       _GotoState26Action,
	{_State94, SemicolonToken}:                    _GotoState163Action,
	{_State94, NotToken}:                          _GotoState28Action,
	{_State94, SubToken}:                          _GotoState32Action,
	{_State94, MulToken}:                          _GotoState27Action,
	{_State94, BitNegToken}:                       _GotoState18Action,
	{_State94, BitAndToken}:                       _GotoState17Action,
	{_State94, LexErrorToken}:                     _GotoState25Action,
	{_State94, OptionalLabelDeclType}:             _GotoState103Action,
	{_State94, SequenceExprType}:                  _GotoState164Action,
	{_State94, BlockExprType}:                     _GotoState40Action,
	{_State94, CallExprType}:                      _GotoState41Action,
	{_State94, AtomExprType}:                      _GotoState39Action,
	{_State94, LiteralType}:                       _GotoState44Action,
	{_State94, AnonymousStructExprType}:           _GotoState38Action,
	{_State94, AccessExprType}:                    _GotoState34Action,
	{_State94, PostfixUnaryExprType}:              _GotoState48Action,
	{_State94, PrefixUnaryOpType}:                 _GotoState50Action,
	{_State94, PrefixUnaryExprType}:               _GotoState49Action,
	{_State94, MulExprType}:                       _GotoState45Action,
	{_State94, AddExprType}:                       _GotoState35Action,
	{_State94, CmpExprType}:                       _GotoState42Action,
	{_State94, AndExprType}:                       _GotoState36Action,
	{_State94, OrExprType}:                        _GotoState47Action,
	{_State94, ExplicitStructDefType}:             _GotoState43Action,
	{_State94, AnonymousFuncExprType}:             _GotoState37Action,
	{_State95, IntegerLiteralToken}:               _GotoState23Action,
	{_State95, FloatLiteralToken}:                 _GotoState20Action,
	{_State95, RuneLiteralToken}:                  _GotoState29Action,
	{_State95, StringLiteralToken}:                _GotoState30Action,
	{_State95, IdentifierToken}:                   _GotoState22Action,
	{_State95, TrueToken}:                         _GotoState33Action,
	{_State95, FalseToken}:                        _GotoState19Action,
	{_State95, StructToken}:                       _GotoState31Action,
	{_State95, FuncToken}:                         _GotoState21Action,
	{_State95, LabelDeclToken}:                    _GotoState24Action,
	{_State95, LparenToken}:                       _GotoState26Action,
	{_State95, NotToken}:                          _GotoState28Action,
	{_State95, SubToken}:                          _GotoState32Action,
	{_State95, MulToken}:                          _GotoState27Action,
	{_State95, BitNegToken}:                       _GotoState18Action,
	{_State95, BitAndToken}:                       _GotoState17Action,
	{_State95, LexErrorToken}:                     _GotoState25Action,
	{_State95, OptionalLabelDeclType}:             _GotoState103Action,
	{_State95, SequenceExprType}:                  _GotoState165Action,
	{_State95, BlockExprType}:                     _GotoState40Action,
	{_State95, CallExprType}:                      _GotoState41Action,
	{_State95, AtomExprType}:                      _GotoState39Action,
	{_State95, LiteralType}:                       _GotoState44Action,
	{_State95, AnonymousStructExprType}:           _GotoState38Action,
	{_State95, AccessExprType}:                    _GotoState34Action,
	{_State95, PostfixUnaryExprType}:              _GotoState48Action,
	{_State95, PrefixUnaryOpType}:                 _GotoState50Action,
	{_State95, PrefixUnaryExprType}:               _GotoState49Action,
	{_State95, MulExprType}:                       _GotoState45Action,
	{_State95, AddExprType}:                       _GotoState35Action,
	{_State95, CmpExprType}:                       _GotoState42Action,
	{_State95, AndExprType}:                       _GotoState36Action,
	{_State95, OrExprType}:                        _GotoState47Action,
	{_State95, ExplicitStructDefType}:             _GotoState43Action,
	{_State95, AnonymousFuncExprType}:             _GotoState37Action,
	{_State96, StatementsType}:                    _GotoState166Action,
	{_State97, IntegerLiteralToken}:               _GotoState23Action,
	{_State97, FloatLiteralToken}:                 _GotoState20Action,
	{_State97, RuneLiteralToken}:                  _GotoState29Action,
	{_State97, StringLiteralToken}:                _GotoState30Action,
	{_State97, IdentifierToken}:                   _GotoState22Action,
	{_State97, TrueToken}:                         _GotoState33Action,
	{_State97, FalseToken}:                        _GotoState19Action,
	{_State97, StructToken}:                       _GotoState31Action,
	{_State97, FuncToken}:                         _GotoState21Action,
	{_State97, LabelDeclToken}:                    _GotoState24Action,
	{_State97, LparenToken}:                       _GotoState26Action,
	{_State97, NotToken}:                          _GotoState28Action,
	{_State97, SubToken}:                          _GotoState32Action,
	{_State97, MulToken}:                          _GotoState27Action,
	{_State97, BitNegToken}:                       _GotoState18Action,
	{_State97, BitAndToken}:                       _GotoState17Action,
	{_State97, LexErrorToken}:                     _GotoState25Action,
	{_State97, OptionalLabelDeclType}:             _GotoState103Action,
	{_State97, SequenceExprType}:                  _GotoState167Action,
	{_State97, BlockExprType}:                     _GotoState40Action,
	{_State97, CallExprType}:                      _GotoState41Action,
	{_State97, AtomExprType}:                      _GotoState39Action,
	{_State97, LiteralType}:                       _GotoState44Action,
	{_State97, AnonymousStructExprType}:           _GotoState38Action,
	{_State97, AccessExprType}:                    _GotoState34Action,
	{_State97, PostfixUnaryExprType}:              _GotoState48Action,
	{_State97, PrefixUnaryOpType}:                 _GotoState50Action,
	{_State97, PrefixUnaryExprType}:               _GotoState49Action,
	{_State97, MulExprType}:                       _GotoState45Action,
	{_State97, AddExprType}:                       _GotoState35Action,
	{_State97, CmpExprType}:                       _GotoState42Action,
	{_State97, AndExprType}:                       _GotoState36Action,
	{_State97, OrExprType}:                        _GotoState47Action,
	{_State97, ExplicitStructDefType}:             _GotoState43Action,
	{_State97, AnonymousFuncExprType}:             _GotoState37Action,
	{_State102, IntegerLiteralToken}:              _GotoState23Action,
	{_State102, FloatLiteralToken}:                _GotoState20Action,
	{_State102, RuneLiteralToken}:                 _GotoState29Action,
	{_State102, StringLiteralToken}:               _GotoState30Action,
	{_State102, IdentifierToken}:                  _GotoState22Action,
	{_State102, TrueToken}:                        _GotoState33Action,
	{_State102, FalseToken}:                       _GotoState19Action,
	{_State102, StructToken}:                      _GotoState31Action,
	{_State102, FuncToken}:                        _GotoState21Action,
	{_State102, LabelDeclToken}:                   _GotoState24Action,
	{_State102, LparenToken}:                      _GotoState26Action,
	{_State102, NotToken}:                         _GotoState28Action,
	{_State102, SubToken}:                         _GotoState32Action,
	{_State102, MulToken}:                         _GotoState27Action,
	{_State102, BitNegToken}:                      _GotoState18Action,
	{_State102, BitAndToken}:                      _GotoState17Action,
	{_State102, LexErrorToken}:                    _GotoState25Action,
	{_State102, OptionalLabelDeclType}:            _GotoState103Action,
	{_State102, BlockExprType}:                    _GotoState40Action,
	{_State102, CallExprType}:                     _GotoState41Action,
	{_State102, AtomExprType}:                     _GotoState39Action,
	{_State102, LiteralType}:                      _GotoState44Action,
	{_State102, AnonymousStructExprType}:          _GotoState38Action,
	{_State102, AccessExprType}:                   _GotoState34Action,
	{_State102, PostfixUnaryExprType}:             _GotoState48Action,
	{_State102, PrefixUnaryOpType}:                _GotoState50Action,
	{_State102, PrefixUnaryExprType}:              _GotoState49Action,
	{_State102, MulExprType}:                      _GotoState45Action,
	{_State102, AddExprType}:                      _GotoState35Action,
	{_State102, CmpExprType}:                      _GotoState42Action,
	{_State102, AndExprType}:                      _GotoState168Action,
	{_State102, ExplicitStructDefType}:            _GotoState43Action,
	{_State102, AnonymousFuncExprType}:            _GotoState37Action,
	{_State103, LbraceToken}:                      _GotoState96Action,
	{_State103, BlockBodyType}:                    _GotoState98Action,
	{_State105, PackageStatementsType}:            _GotoState169Action,
	{_State106, IdentifierToken}:                  _GotoState170Action,
	{_State106, GenericParameterDefType}:          _GotoState171Action,
	{_State106, GenericParameterDefsType}:         _GotoState172Action,
	{_State106, OptionalGenericParameterDefsType}: _GotoState173Action,
	{_State107, IdentifierToken}:                  _GotoState149Action,
	{_State107, StructToken}:                      _GotoState31Action,
	{_State107, EnumToken}:                        _GotoState111Action,
	{_State107, TraitToken}:                       _GotoState15Action,
	{_State107, FuncToken}:                        _GotoState145Action,
	{_State107, LparenToken}:                      _GotoState114Action,
	{_State107, QuestionToken}:                    _GotoState115Action,
	{_State107, TildeTildeToken}:                  _GotoState116Action,
	{_State107, BitNegToken}:                      _GotoState110Action,
	{_State107, BitAndToken}:                      _GotoState109Action,
	{_State107, AtomTypeType}:                     _GotoState118Action,
	{_State107, TraitableTypeType}:                _GotoState131Action,
	{_State107, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State107, ValueTypeType}:                    _GotoState174Action,
	{_State107, ImplicitStructDefType}:            _GotoState124Action,
	{_State107, ExplicitStructDefType}:            _GotoState120Action,
	{_State107, ImplicitEnumDefType}:              _GotoState123Action,
	{_State107, ExplicitEnumDefType}:              _GotoState119Action,
	{_State107, TraitDefType}:                     _GotoState128Action,
	{_State107, FuncTypeType}:                     _GotoState122Action,
	{_State108, IdentifierToken}:                  _GotoState149Action,
	{_State108, StructToken}:                      _GotoState31Action,
	{_State108, EnumToken}:                        _GotoState111Action,
	{_State108, TraitToken}:                       _GotoState15Action,
	{_State108, FuncToken}:                        _GotoState145Action,
	{_State108, LparenToken}:                      _GotoState114Action,
	{_State108, QuestionToken}:                    _GotoState115Action,
	{_State108, TildeTildeToken}:                  _GotoState116Action,
	{_State108, BitNegToken}:                      _GotoState110Action,
	{_State108, BitAndToken}:                      _GotoState109Action,
	{_State108, AtomTypeType}:                     _GotoState118Action,
	{_State108, TraitableTypeType}:                _GotoState131Action,
	{_State108, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State108, ValueTypeType}:                    _GotoState175Action,
	{_State108, ImplicitStructDefType}:            _GotoState124Action,
	{_State108, ExplicitStructDefType}:            _GotoState120Action,
	{_State108, ImplicitEnumDefType}:              _GotoState123Action,
	{_State108, ExplicitEnumDefType}:              _GotoState119Action,
	{_State108, TraitDefType}:                     _GotoState128Action,
	{_State108, FuncTypeType}:                     _GotoState122Action,
	{_State109, IdentifierToken}:                  _GotoState149Action,
	{_State109, StructToken}:                      _GotoState31Action,
	{_State109, EnumToken}:                        _GotoState111Action,
	{_State109, TraitToken}:                       _GotoState15Action,
	{_State109, LparenToken}:                      _GotoState114Action,
	{_State109, TildeTildeToken}:                  _GotoState116Action,
	{_State109, BitNegToken}:                      _GotoState110Action,
	{_State109, AtomTypeType}:                     _GotoState118Action,
	{_State109, TraitableTypeType}:                _GotoState131Action,
	{_State109, TraitAlgebraTypeType}:             _GotoState176Action,
	{_State109, ImplicitStructDefType}:            _GotoState124Action,
	{_State109, ExplicitStructDefType}:            _GotoState120Action,
	{_State109, ImplicitEnumDefType}:              _GotoState123Action,
	{_State109, ExplicitEnumDefType}:              _GotoState119Action,
	{_State109, TraitDefType}:                     _GotoState128Action,
	{_State110, IdentifierToken}:                  _GotoState149Action,
	{_State110, StructToken}:                      _GotoState31Action,
	{_State110, EnumToken}:                        _GotoState111Action,
	{_State110, TraitToken}:                       _GotoState15Action,
	{_State110, LparenToken}:                      _GotoState114Action,
	{_State110, AtomTypeType}:                     _GotoState177Action,
	{_State110, ImplicitStructDefType}:            _GotoState124Action,
	{_State110, ExplicitStructDefType}:            _GotoState120Action,
	{_State110, ImplicitEnumDefType}:              _GotoState123Action,
	{_State110, ExplicitEnumDefType}:              _GotoState119Action,
	{_State110, TraitDefType}:                     _GotoState128Action,
	{_State111, LparenToken}:                      _GotoState178Action,
	{_State112, IdentifierToken}:                  _GotoState179Action,
	{_State112, LparenToken}:                      _GotoState180Action,
	{_State113, IdentifierToken}:                  _GotoState149Action,
	{_State113, StructToken}:                      _GotoState31Action,
	{_State113, EnumToken}:                        _GotoState111Action,
	{_State113, TraitToken}:                       _GotoState15Action,
	{_State113, FuncToken}:                        _GotoState145Action,
	{_State113, LparenToken}:                      _GotoState114Action,
	{_State113, QuestionToken}:                    _GotoState115Action,
	{_State113, DollarLbracketToken}:              _GotoState67Action,
	{_State113, TildeTildeToken}:                  _GotoState116Action,
	{_State113, BitNegToken}:                      _GotoState110Action,
	{_State113, BitAndToken}:                      _GotoState109Action,
	{_State113, OptionalGenericBindingType}:       _GotoState181Action,
	{_State113, AtomTypeType}:                     _GotoState118Action,
	{_State113, TraitableTypeType}:                _GotoState131Action,
	{_State113, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State113, ValueTypeType}:                    _GotoState182Action,
	{_State113, ImplicitStructDefType}:            _GotoState124Action,
	{_State113, ExplicitStructDefType}:            _GotoState120Action,
	{_State113, ImplicitEnumDefType}:              _GotoState123Action,
	{_State113, ExplicitEnumDefType}:              _GotoState119Action,
	{_State113, TraitDefType}:                     _GotoState128Action,
	{_State113, FuncTypeType}:                     _GotoState122Action,
	{_State114, IdentifierToken}:                  _GotoState113Action,
	{_State114, UnsafeToken}:                      _GotoState117Action,
	{_State114, StructToken}:                      _GotoState31Action,
	{_State114, EnumToken}:                        _GotoState111Action,
	{_State114, TraitToken}:                       _GotoState15Action,
	{_State114, FuncToken}:                        _GotoState145Action,
	{_State114, LparenToken}:                      _GotoState114Action,
	{_State114, QuestionToken}:                    _GotoState115Action,
	{_State114, TildeTildeToken}:                  _GotoState116Action,
	{_State114, BitNegToken}:                      _GotoState110Action,
	{_State114, BitAndToken}:                      _GotoState109Action,
	{_State114, UnsafeStatementType}:              _GotoState132Action,
	{_State114, AtomTypeType}:                     _GotoState118Action,
	{_State114, TraitableTypeType}:                _GotoState131Action,
	{_State114, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State114, ValueTypeType}:                    _GotoState133Action,
	{_State114, FieldDefType}:                     _GotoState184Action,
	{_State114, ImplicitFieldDefsType}:            _GotoState186Action,
	{_State114, OptionalImplicitFieldDefsType}:    _GotoState187Action,
	{_State114, ImplicitStructDefType}:            _GotoState124Action,
	{_State114, ExplicitStructDefType}:            _GotoState120Action,
	{_State114, EnumValueDefType}:                 _GotoState183Action,
	{_State114, ImplicitEnumValueDefsType}:        _GotoState185Action,
	{_State114, ImplicitEnumDefType}:              _GotoState123Action,
	{_State114, ExplicitEnumDefType}:              _GotoState119Action,
	{_State114, TraitDefType}:                     _GotoState128Action,
	{_State114, FuncTypeType}:                     _GotoState122Action,
	{_State116, IdentifierToken}:                  _GotoState149Action,
	{_State116, StructToken}:                      _GotoState31Action,
	{_State116, EnumToken}:                        _GotoState111Action,
	{_State116, TraitToken}:                       _GotoState15Action,
	{_State116, LparenToken}:                      _GotoState114Action,
	{_State116, AtomTypeType}:                     _GotoState188Action,
	{_State116, ImplicitStructDefType}:            _GotoState124Action,
	{_State116, ExplicitStructDefType}:            _GotoState120Action,
	{_State116, ImplicitEnumDefType}:              _GotoState123Action,
	{_State116, ExplicitEnumDefType}:              _GotoState119Action,
	{_State116, TraitDefType}:                     _GotoState128Action,
	{_State117, LessToken}:                        _GotoState189Action,
	{_State126, RparenToken}:                      _GotoState190Action,
	{_State127, AddToken}:                         _GotoState191Action,
	{_State127, SubToken}:                         _GotoState193Action,
	{_State127, MulToken}:                         _GotoState192Action,
	{_State129, NewlinesToken}:                    _GotoState195Action,
	{_State129, CommaToken}:                       _GotoState194Action,
	{_State134, IdentifierToken}:                  _GotoState149Action,
	{_State134, StructToken}:                      _GotoState31Action,
	{_State134, EnumToken}:                        _GotoState111Action,
	{_State134, TraitToken}:                       _GotoState15Action,
	{_State134, FuncToken}:                        _GotoState145Action,
	{_State134, LparenToken}:                      _GotoState114Action,
	{_State134, QuestionToken}:                    _GotoState115Action,
	{_State134, DotdotdotToken}:                   _GotoState196Action,
	{_State134, TildeTildeToken}:                  _GotoState116Action,
	{_State134, BitNegToken}:                      _GotoState110Action,
	{_State134, BitAndToken}:                      _GotoState109Action,
	{_State134, AtomTypeType}:                     _GotoState118Action,
	{_State134, TraitableTypeType}:                _GotoState131Action,
	{_State134, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State134, ValueTypeType}:                    _GotoState197Action,
	{_State134, ImplicitStructDefType}:            _GotoState124Action,
	{_State134, ExplicitStructDefType}:            _GotoState120Action,
	{_State134, ImplicitEnumDefType}:              _GotoState123Action,
	{_State134, ExplicitEnumDefType}:              _GotoState119Action,
	{_State134, TraitDefType}:                     _GotoState128Action,
	{_State134, FuncTypeType}:                     _GotoState122Action,
	{_State135, RparenToken}:                      _GotoState198Action,
	{_State136, DollarLbracketToken}:              _GotoState106Action,
	{_State136, OptionalGenericParametersType}:    _GotoState199Action,
	{_State137, RparenToken}:                      _GotoState200Action,
	{_State139, CommaToken}:                       _GotoState201Action,
	{_State140, IntegerLiteralToken}:              _GotoState23Action,
	{_State140, FloatLiteralToken}:                _GotoState20Action,
	{_State140, RuneLiteralToken}:                 _GotoState29Action,
	{_State140, StringLiteralToken}:               _GotoState30Action,
	{_State140, IdentifierToken}:                  _GotoState22Action,
	{_State140, TrueToken}:                        _GotoState33Action,
	{_State140, FalseToken}:                       _GotoState19Action,
	{_State140, StructToken}:                      _GotoState31Action,
	{_State140, FuncToken}:                        _GotoState21Action,
	{_State140, LabelDeclToken}:                   _GotoState24Action,
	{_State140, LparenToken}:                      _GotoState26Action,
	{_State140, NotToken}:                         _GotoState28Action,
	{_State140, SubToken}:                         _GotoState32Action,
	{_State140, MulToken}:                         _GotoState27Action,
	{_State140, BitNegToken}:                      _GotoState18Action,
	{_State140, BitAndToken}:                      _GotoState17Action,
	{_State140, LexErrorToken}:                    _GotoState25Action,
	{_State140, ExpressionType}:                   _GotoState202Action,
	{_State140, OptionalLabelDeclType}:            _GotoState46Action,
	{_State140, SequenceExprType}:                 _GotoState51Action,
	{_State140, BlockExprType}:                    _GotoState40Action,
	{_State140, CallExprType}:                     _GotoState41Action,
	{_State140, AtomExprType}:                     _GotoState39Action,
	{_State140, LiteralType}:                      _GotoState44Action,
	{_State140, AnonymousStructExprType}:          _GotoState38Action,
	{_State140, AccessExprType}:                   _GotoState34Action,
	{_State140, PostfixUnaryExprType}:             _GotoState48Action,
	{_State140, PrefixUnaryOpType}:                _GotoState50Action,
	{_State140, PrefixUnaryExprType}:              _GotoState49Action,
	{_State140, MulExprType}:                      _GotoState45Action,
	{_State140, AddExprType}:                      _GotoState35Action,
	{_State140, CmpExprType}:                      _GotoState42Action,
	{_State140, AndExprType}:                      _GotoState36Action,
	{_State140, OrExprType}:                       _GotoState47Action,
	{_State140, ExplicitStructDefType}:            _GotoState43Action,
	{_State140, AnonymousFuncExprType}:            _GotoState37Action,
	{_State141, IntegerLiteralToken}:              _GotoState23Action,
	{_State141, FloatLiteralToken}:                _GotoState20Action,
	{_State141, RuneLiteralToken}:                 _GotoState29Action,
	{_State141, StringLiteralToken}:               _GotoState30Action,
	{_State141, IdentifierToken}:                  _GotoState60Action,
	{_State141, TrueToken}:                        _GotoState33Action,
	{_State141, FalseToken}:                       _GotoState19Action,
	{_State141, StructToken}:                      _GotoState31Action,
	{_State141, FuncToken}:                        _GotoState21Action,
	{_State141, LabelDeclToken}:                   _GotoState24Action,
	{_State141, LparenToken}:                      _GotoState26Action,
	{_State141, NotToken}:                         _GotoState28Action,
	{_State141, SubToken}:                         _GotoState32Action,
	{_State141, MulToken}:                         _GotoState27Action,
	{_State141, BitNegToken}:                      _GotoState18Action,
	{_State141, BitAndToken}:                      _GotoState17Action,
	{_State141, LexErrorToken}:                    _GotoState25Action,
	{_State141, ExpressionType}:                   _GotoState64Action,
	{_State141, OptionalLabelDeclType}:            _GotoState46Action,
	{_State141, SequenceExprType}:                 _GotoState51Action,
	{_State141, BlockExprType}:                    _GotoState40Action,
	{_State141, CallExprType}:                     _GotoState41Action,
	{_State141, ArgumentType}:                     _GotoState203Action,
	{_State141, ColonExpressionsType}:             _GotoState63Action,
	{_State141, OptionalExpressionType}:           _GotoState65Action,
	{_State141, AtomExprType}:                     _GotoState39Action,
	{_State141, LiteralType}:                      _GotoState44Action,
	{_State141, AnonymousStructExprType}:          _GotoState38Action,
	{_State141, AccessExprType}:                   _GotoState34Action,
	{_State141, PostfixUnaryExprType}:             _GotoState48Action,
	{_State141, PrefixUnaryOpType}:                _GotoState50Action,
	{_State141, PrefixUnaryExprType}:              _GotoState49Action,
	{_State141, MulExprType}:                      _GotoState45Action,
	{_State141, AddExprType}:                      _GotoState35Action,
	{_State141, CmpExprType}:                      _GotoState42Action,
	{_State141, AndExprType}:                      _GotoState36Action,
	{_State141, OrExprType}:                       _GotoState47Action,
	{_State141, ExplicitStructDefType}:            _GotoState43Action,
	{_State141, AnonymousFuncExprType}:            _GotoState37Action,
	{_State143, IntegerLiteralToken}:              _GotoState23Action,
	{_State143, FloatLiteralToken}:                _GotoState20Action,
	{_State143, RuneLiteralToken}:                 _GotoState29Action,
	{_State143, StringLiteralToken}:               _GotoState30Action,
	{_State143, IdentifierToken}:                  _GotoState22Action,
	{_State143, TrueToken}:                        _GotoState33Action,
	{_State143, FalseToken}:                       _GotoState19Action,
	{_State143, StructToken}:                      _GotoState31Action,
	{_State143, FuncToken}:                        _GotoState21Action,
	{_State143, LabelDeclToken}:                   _GotoState24Action,
	{_State143, LparenToken}:                      _GotoState26Action,
	{_State143, NotToken}:                         _GotoState28Action,
	{_State143, SubToken}:                         _GotoState32Action,
	{_State143, MulToken}:                         _GotoState27Action,
	{_State143, BitNegToken}:                      _GotoState18Action,
	{_State143, BitAndToken}:                      _GotoState17Action,
	{_State143, LexErrorToken}:                    _GotoState25Action,
	{_State143, ExpressionType}:                   _GotoState204Action,
	{_State143, OptionalLabelDeclType}:            _GotoState46Action,
	{_State143, SequenceExprType}:                 _GotoState51Action,
	{_State143, BlockExprType}:                    _GotoState40Action,
	{_State143, CallExprType}:                     _GotoState41Action,
	{_State143, OptionalExpressionType}:           _GotoState205Action,
	{_State143, AtomExprType}:                     _GotoState39Action,
	{_State143, LiteralType}:                      _GotoState44Action,
	{_State143, AnonymousStructExprType}:          _GotoState38Action,
	{_State143, AccessExprType}:                   _GotoState34Action,
	{_State143, PostfixUnaryExprType}:             _GotoState48Action,
	{_State143, PrefixUnaryOpType}:                _GotoState50Action,
	{_State143, PrefixUnaryExprType}:              _GotoState49Action,
	{_State143, MulExprType}:                      _GotoState45Action,
	{_State143, AddExprType}:                      _GotoState35Action,
	{_State143, CmpExprType}:                      _GotoState42Action,
	{_State143, AndExprType}:                      _GotoState36Action,
	{_State143, OrExprType}:                       _GotoState47Action,
	{_State143, ExplicitStructDefType}:            _GotoState43Action,
	{_State143, AnonymousFuncExprType}:            _GotoState37Action,
	{_State144, IntegerLiteralToken}:              _GotoState23Action,
	{_State144, FloatLiteralToken}:                _GotoState20Action,
	{_State144, RuneLiteralToken}:                 _GotoState29Action,
	{_State144, StringLiteralToken}:               _GotoState30Action,
	{_State144, IdentifierToken}:                  _GotoState22Action,
	{_State144, TrueToken}:                        _GotoState33Action,
	{_State144, FalseToken}:                       _GotoState19Action,
	{_State144, StructToken}:                      _GotoState31Action,
	{_State144, FuncToken}:                        _GotoState21Action,
	{_State144, LabelDeclToken}:                   _GotoState24Action,
	{_State144, LparenToken}:                      _GotoState26Action,
	{_State144, NotToken}:                         _GotoState28Action,
	{_State144, SubToken}:                         _GotoState32Action,
	{_State144, MulToken}:                         _GotoState27Action,
	{_State144, BitNegToken}:                      _GotoState18Action,
	{_State144, BitAndToken}:                      _GotoState17Action,
	{_State144, LexErrorToken}:                    _GotoState25Action,
	{_State144, ExpressionType}:                   _GotoState204Action,
	{_State144, OptionalLabelDeclType}:            _GotoState46Action,
	{_State144, SequenceExprType}:                 _GotoState51Action,
	{_State144, BlockExprType}:                    _GotoState40Action,
	{_State144, CallExprType}:                     _GotoState41Action,
	{_State144, OptionalExpressionType}:           _GotoState206Action,
	{_State144, AtomExprType}:                     _GotoState39Action,
	{_State144, LiteralType}:                      _GotoState44Action,
	{_State144, AnonymousStructExprType}:          _GotoState38Action,
	{_State144, AccessExprType}:                   _GotoState34Action,
	{_State144, PostfixUnaryExprType}:             _GotoState48Action,
	{_State144, PrefixUnaryOpType}:                _GotoState50Action,
	{_State144, PrefixUnaryExprType}:              _GotoState49Action,
	{_State144, MulExprType}:                      _GotoState45Action,
	{_State144, AddExprType}:                      _GotoState35Action,
	{_State144, CmpExprType}:                      _GotoState42Action,
	{_State144, AndExprType}:                      _GotoState36Action,
	{_State144, OrExprType}:                       _GotoState47Action,
	{_State144, ExplicitStructDefType}:            _GotoState43Action,
	{_State144, AnonymousFuncExprType}:            _GotoState37Action,
	{_State145, LparenToken}:                      _GotoState180Action,
	{_State146, NewlinesToken}:                    _GotoState208Action,
	{_State146, CommaToken}:                       _GotoState207Action,
	{_State148, RparenToken}:                      _GotoState209Action,
	{_State149, DollarLbracketToken}:              _GotoState67Action,
	{_State149, OptionalGenericBindingType}:       _GotoState181Action,
	{_State150, CommaToken}:                       _GotoState210Action,
	{_State151, RbracketToken}:                    _GotoState211Action,
	{_State154, RbracketToken}:                    _GotoState212Action,
	{_State155, IntegerLiteralToken}:              _GotoState23Action,
	{_State155, FloatLiteralToken}:                _GotoState20Action,
	{_State155, RuneLiteralToken}:                 _GotoState29Action,
	{_State155, StringLiteralToken}:               _GotoState30Action,
	{_State155, IdentifierToken}:                  _GotoState60Action,
	{_State155, TrueToken}:                        _GotoState33Action,
	{_State155, FalseToken}:                       _GotoState19Action,
	{_State155, StructToken}:                      _GotoState31Action,
	{_State155, FuncToken}:                        _GotoState21Action,
	{_State155, LabelDeclToken}:                   _GotoState24Action,
	{_State155, LparenToken}:                      _GotoState26Action,
	{_State155, NotToken}:                         _GotoState28Action,
	{_State155, SubToken}:                         _GotoState32Action,
	{_State155, MulToken}:                         _GotoState27Action,
	{_State155, BitNegToken}:                      _GotoState18Action,
	{_State155, BitAndToken}:                      _GotoState17Action,
	{_State155, LexErrorToken}:                    _GotoState25Action,
	{_State155, ExpressionType}:                   _GotoState64Action,
	{_State155, OptionalLabelDeclType}:            _GotoState46Action,
	{_State155, SequenceExprType}:                 _GotoState51Action,
	{_State155, BlockExprType}:                    _GotoState40Action,
	{_State155, CallExprType}:                     _GotoState41Action,
	{_State155, OptionalArgumentsType}:            _GotoState214Action,
	{_State155, ArgumentsType}:                    _GotoState213Action,
	{_State155, ArgumentType}:                     _GotoState61Action,
	{_State155, ColonExpressionsType}:             _GotoState63Action,
	{_State155, OptionalExpressionType}:           _GotoState65Action,
	{_State155, AtomExprType}:                     _GotoState39Action,
	{_State155, LiteralType}:                      _GotoState44Action,
	{_State155, AnonymousStructExprType}:          _GotoState38Action,
	{_State155, AccessExprType}:                   _GotoState34Action,
	{_State155, PostfixUnaryExprType}:             _GotoState48Action,
	{_State155, PrefixUnaryOpType}:                _GotoState50Action,
	{_State155, PrefixUnaryExprType}:              _GotoState49Action,
	{_State155, MulExprType}:                      _GotoState45Action,
	{_State155, AddExprType}:                      _GotoState35Action,
	{_State155, CmpExprType}:                      _GotoState42Action,
	{_State155, AndExprType}:                      _GotoState36Action,
	{_State155, OrExprType}:                       _GotoState47Action,
	{_State155, ExplicitStructDefType}:            _GotoState43Action,
	{_State155, AnonymousFuncExprType}:            _GotoState37Action,
	{_State156, MulToken}:                         _GotoState91Action,
	{_State156, DivToken}:                         _GotoState89Action,
	{_State156, ModToken}:                         _GotoState90Action,
	{_State156, BitAndToken}:                      _GotoState86Action,
	{_State156, BitLshiftToken}:                   _GotoState87Action,
	{_State156, BitRshiftToken}:                   _GotoState88Action,
	{_State156, MulOpType}:                        _GotoState92Action,
	{_State157, EqualToken}:                       _GotoState78Action,
	{_State157, NotEqualToken}:                    _GotoState83Action,
	{_State157, LessToken}:                        _GotoState81Action,
	{_State157, LessOrEqualToken}:                 _GotoState82Action,
	{_State157, GreaterToken}:                     _GotoState79Action,
	{_State157, GreaterOrEqualToken}:              _GotoState80Action,
	{_State157, CmpOpType}:                        _GotoState84Action,
	{_State158, AddToken}:                         _GotoState72Action,
	{_State158, SubToken}:                         _GotoState75Action,
	{_State158, BitXorToken}:                      _GotoState74Action,
	{_State158, BitOrToken}:                       _GotoState73Action,
	{_State158, AddOpType}:                        _GotoState76Action,
	{_State159, RparenToken}:                      _GotoState215Action,
	{_State159, CommaToken}:                       _GotoState141Action,
	{_State161, ForToken}:                         _GotoState216Action,
	{_State162, IntegerLiteralToken}:              _GotoState23Action,
	{_State162, FloatLiteralToken}:                _GotoState20Action,
	{_State162, RuneLiteralToken}:                 _GotoState29Action,
	{_State162, StringLiteralToken}:               _GotoState30Action,
	{_State162, IdentifierToken}:                  _GotoState22Action,
	{_State162, TrueToken}:                        _GotoState33Action,
	{_State162, FalseToken}:                       _GotoState19Action,
	{_State162, StructToken}:                      _GotoState31Action,
	{_State162, FuncToken}:                        _GotoState21Action,
	{_State162, LabelDeclToken}:                   _GotoState24Action,
	{_State162, LparenToken}:                      _GotoState26Action,
	{_State162, NotToken}:                         _GotoState28Action,
	{_State162, SubToken}:                         _GotoState32Action,
	{_State162, MulToken}:                         _GotoState27Action,
	{_State162, BitNegToken}:                      _GotoState18Action,
	{_State162, BitAndToken}:                      _GotoState17Action,
	{_State162, LexErrorToken}:                    _GotoState25Action,
	{_State162, OptionalLabelDeclType}:            _GotoState103Action,
	{_State162, SequenceExprType}:                 _GotoState217Action,
	{_State162, BlockExprType}:                    _GotoState40Action,
	{_State162, CallExprType}:                     _GotoState41Action,
	{_State162, AtomExprType}:                     _GotoState39Action,
	{_State162, LiteralType}:                      _GotoState44Action,
	{_State162, AnonymousStructExprType}:          _GotoState38Action,
	{_State162, AccessExprType}:                   _GotoState34Action,
	{_State162, PostfixUnaryExprType}:             _GotoState48Action,
	{_State162, PrefixUnaryOpType}:                _GotoState50Action,
	{_State162, PrefixUnaryExprType}:              _GotoState49Action,
	{_State162, MulExprType}:                      _GotoState45Action,
	{_State162, AddExprType}:                      _GotoState35Action,
	{_State162, CmpExprType}:                      _GotoState42Action,
	{_State162, AndExprType}:                      _GotoState36Action,
	{_State162, OrExprType}:                       _GotoState47Action,
	{_State162, ExplicitStructDefType}:            _GotoState43Action,
	{_State162, AnonymousFuncExprType}:            _GotoState37Action,
	{_State163, IntegerLiteralToken}:              _GotoState23Action,
	{_State163, FloatLiteralToken}:                _GotoState20Action,
	{_State163, RuneLiteralToken}:                 _GotoState29Action,
	{_State163, StringLiteralToken}:               _GotoState30Action,
	{_State163, IdentifierToken}:                  _GotoState22Action,
	{_State163, TrueToken}:                        _GotoState33Action,
	{_State163, FalseToken}:                       _GotoState19Action,
	{_State163, StructToken}:                      _GotoState31Action,
	{_State163, FuncToken}:                        _GotoState21Action,
	{_State163, LabelDeclToken}:                   _GotoState24Action,
	{_State163, LparenToken}:                      _GotoState26Action,
	{_State163, NotToken}:                         _GotoState28Action,
	{_State163, SubToken}:                         _GotoState32Action,
	{_State163, MulToken}:                         _GotoState27Action,
	{_State163, BitNegToken}:                      _GotoState18Action,
	{_State163, BitAndToken}:                      _GotoState17Action,
	{_State163, LexErrorToken}:                    _GotoState25Action,
	{_State163, OptionalLabelDeclType}:            _GotoState103Action,
	{_State163, OptionalSequenceExprType}:         _GotoState218Action,
	{_State163, SequenceExprType}:                 _GotoState219Action,
	{_State163, BlockExprType}:                    _GotoState40Action,
	{_State163, CallExprType}:                     _GotoState41Action,
	{_State163, AtomExprType}:                     _GotoState39Action,
	{_State163, LiteralType}:                      _GotoState44Action,
	{_State163, AnonymousStructExprType}:          _GotoState38Action,
	{_State163, AccessExprType}:                   _GotoState34Action,
	{_State163, PostfixUnaryExprType}:             _GotoState48Action,
	{_State163, PrefixUnaryOpType}:                _GotoState50Action,
	{_State163, PrefixUnaryExprType}:              _GotoState49Action,
	{_State163, MulExprType}:                      _GotoState45Action,
	{_State163, AddExprType}:                      _GotoState35Action,
	{_State163, CmpExprType}:                      _GotoState42Action,
	{_State163, AndExprType}:                      _GotoState36Action,
	{_State163, OrExprType}:                       _GotoState47Action,
	{_State163, ExplicitStructDefType}:            _GotoState43Action,
	{_State163, AnonymousFuncExprType}:            _GotoState37Action,
	{_State164, DoToken}:                          _GotoState220Action,
	{_State165, LbraceToken}:                      _GotoState96Action,
	{_State165, BlockBodyType}:                    _GotoState221Action,
	{_State166, IntegerLiteralToken}:              _GotoState23Action,
	{_State166, FloatLiteralToken}:                _GotoState20Action,
	{_State166, RuneLiteralToken}:                 _GotoState29Action,
	{_State166, StringLiteralToken}:               _GotoState30Action,
	{_State166, IdentifierToken}:                  _GotoState22Action,
	{_State166, TrueToken}:                        _GotoState33Action,
	{_State166, FalseToken}:                       _GotoState19Action,
	{_State166, ReturnToken}:                      _GotoState227Action,
	{_State166, BreakToken}:                       _GotoState223Action,
	{_State166, ContinueToken}:                    _GotoState224Action,
	{_State166, UnsafeToken}:                      _GotoState117Action,
	{_State166, StructToken}:                      _GotoState31Action,
	{_State166, FuncToken}:                        _GotoState21Action,
	{_State166, AsyncToken}:                       _GotoState222Action,
	{_State166, DeferToken}:                       _GotoState225Action,
	{_State166, LabelDeclToken}:                   _GotoState24Action,
	{_State166, RbraceToken}:                      _GotoState226Action,
	{_State166, LparenToken}:                      _GotoState26Action,
	{_State166, NotToken}:                         _GotoState28Action,
	{_State166, SubToken}:                         _GotoState32Action,
	{_State166, MulToken}:                         _GotoState27Action,
	{_State166, BitNegToken}:                      _GotoState18Action,
	{_State166, BitAndToken}:                      _GotoState17Action,
	{_State166, LexErrorToken}:                    _GotoState25Action,
	{_State166, ExpressionType}:                   _GotoState229Action,
	{_State166, OptionalLabelDeclType}:            _GotoState46Action,
	{_State166, SequenceExprType}:                 _GotoState51Action,
	{_State166, BlockExprType}:                    _GotoState40Action,
	{_State166, StatementType}:                    _GotoState233Action,
	{_State166, StatementBodyType}:                _GotoState234Action,
	{_State166, UnsafeStatementType}:              _GotoState235Action,
	{_State166, JumpStatementType}:                _GotoState231Action,
	{_State166, JumpTypeType}:                     _GotoState232Action,
	{_State166, ExpressionsType}:                  _GotoState230Action,
	{_State166, CallExprType}:                     _GotoState41Action,
	{_State166, AtomExprType}:                     _GotoState39Action,
	{_State166, LiteralType}:                      _GotoState44Action,
	{_State166, AnonymousStructExprType}:          _GotoState38Action,
	{_State166, AccessExprType}:                   _GotoState228Action,
	{_State166, PostfixUnaryExprType}:             _GotoState48Action,
	{_State166, PrefixUnaryOpType}:                _GotoState50Action,
	{_State166, PrefixUnaryExprType}:              _GotoState49Action,
	{_State166, MulExprType}:                      _GotoState45Action,
	{_State166, AddExprType}:                      _GotoState35Action,
	{_State166, CmpExprType}:                      _GotoState42Action,
	{_State166, AndExprType}:                      _GotoState36Action,
	{_State166, OrExprType}:                       _GotoState47Action,
	{_State166, ExplicitStructDefType}:            _GotoState43Action,
	{_State166, AnonymousFuncExprType}:            _GotoState37Action,
	{_State167, LbraceToken}:                      _GotoState236Action,
	{_State168, AndToken}:                         _GotoState77Action,
	{_State169, UnsafeToken}:                      _GotoState117Action,
	{_State169, RparenToken}:                      _GotoState237Action,
	{_State169, UnsafeStatementType}:              _GotoState240Action,
	{_State169, PackageStatementBodyType}:         _GotoState239Action,
	{_State169, PackageStatementType}:             _GotoState238Action,
	{_State170, IdentifierToken}:                  _GotoState149Action,
	{_State170, StructToken}:                      _GotoState31Action,
	{_State170, EnumToken}:                        _GotoState111Action,
	{_State170, TraitToken}:                       _GotoState15Action,
	{_State170, FuncToken}:                        _GotoState145Action,
	{_State170, LparenToken}:                      _GotoState114Action,
	{_State170, QuestionToken}:                    _GotoState115Action,
	{_State170, TildeTildeToken}:                  _GotoState116Action,
	{_State170, BitNegToken}:                      _GotoState110Action,
	{_State170, BitAndToken}:                      _GotoState109Action,
	{_State170, AtomTypeType}:                     _GotoState118Action,
	{_State170, TraitableTypeType}:                _GotoState131Action,
	{_State170, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State170, ValueTypeType}:                    _GotoState241Action,
	{_State170, ImplicitStructDefType}:            _GotoState124Action,
	{_State170, ExplicitStructDefType}:            _GotoState120Action,
	{_State170, ImplicitEnumDefType}:              _GotoState123Action,
	{_State170, ExplicitEnumDefType}:              _GotoState119Action,
	{_State170, TraitDefType}:                     _GotoState128Action,
	{_State170, FuncTypeType}:                     _GotoState122Action,
	{_State172, CommaToken}:                       _GotoState242Action,
	{_State173, RbracketToken}:                    _GotoState243Action,
	{_State175, ImplementsToken}:                  _GotoState244Action,
	{_State176, AddToken}:                         _GotoState191Action,
	{_State176, SubToken}:                         _GotoState193Action,
	{_State176, MulToken}:                         _GotoState192Action,
	{_State178, IdentifierToken}:                  _GotoState113Action,
	{_State178, UnsafeToken}:                      _GotoState117Action,
	{_State178, StructToken}:                      _GotoState31Action,
	{_State178, EnumToken}:                        _GotoState111Action,
	{_State178, TraitToken}:                       _GotoState15Action,
	{_State178, FuncToken}:                        _GotoState145Action,
	{_State178, LparenToken}:                      _GotoState114Action,
	{_State178, QuestionToken}:                    _GotoState115Action,
	{_State178, TildeTildeToken}:                  _GotoState116Action,
	{_State178, BitNegToken}:                      _GotoState110Action,
	{_State178, BitAndToken}:                      _GotoState109Action,
	{_State178, UnsafeStatementType}:              _GotoState132Action,
	{_State178, AtomTypeType}:                     _GotoState118Action,
	{_State178, TraitableTypeType}:                _GotoState131Action,
	{_State178, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State178, ValueTypeType}:                    _GotoState133Action,
	{_State178, FieldDefType}:                     _GotoState247Action,
	{_State178, ImplicitStructDefType}:            _GotoState124Action,
	{_State178, ExplicitStructDefType}:            _GotoState120Action,
	{_State178, EnumValueDefType}:                 _GotoState245Action,
	{_State178, ImplicitEnumValueDefsType}:        _GotoState248Action,
	{_State178, ImplicitEnumDefType}:              _GotoState123Action,
	{_State178, ExplicitEnumValueDefsType}:        _GotoState246Action,
	{_State178, ExplicitEnumDefType}:              _GotoState119Action,
	{_State178, TraitDefType}:                     _GotoState128Action,
	{_State178, FuncTypeType}:                     _GotoState122Action,
	{_State179, LparenToken}:                      _GotoState249Action,
	{_State180, IdentifierToken}:                  _GotoState251Action,
	{_State180, StructToken}:                      _GotoState31Action,
	{_State180, EnumToken}:                        _GotoState111Action,
	{_State180, TraitToken}:                       _GotoState15Action,
	{_State180, FuncToken}:                        _GotoState145Action,
	{_State180, LparenToken}:                      _GotoState114Action,
	{_State180, QuestionToken}:                    _GotoState115Action,
	{_State180, DotdotdotToken}:                   _GotoState250Action,
	{_State180, TildeTildeToken}:                  _GotoState116Action,
	{_State180, BitNegToken}:                      _GotoState110Action,
	{_State180, BitAndToken}:                      _GotoState109Action,
	{_State180, AtomTypeType}:                     _GotoState118Action,
	{_State180, TraitableTypeType}:                _GotoState131Action,
	{_State180, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State180, ValueTypeType}:                    _GotoState255Action,
	{_State180, ImplicitStructDefType}:            _GotoState124Action,
	{_State180, ExplicitStructDefType}:            _GotoState120Action,
	{_State180, ImplicitEnumDefType}:              _GotoState123Action,
	{_State180, ExplicitEnumDefType}:              _GotoState119Action,
	{_State180, TraitDefType}:                     _GotoState128Action,
	{_State180, ParameterDeclType}:                _GotoState253Action,
	{_State180, ParameterDeclsType}:               _GotoState254Action,
	{_State180, OptionalParameterDeclsType}:       _GotoState252Action,
	{_State180, FuncTypeType}:                     _GotoState122Action,
	{_State183, OrToken}:                          _GotoState256Action,
	{_State184, AssignToken}:                      _GotoState257Action,
	{_State185, RparenToken}:                      _GotoState259Action,
	{_State185, OrToken}:                          _GotoState258Action,
	{_State186, CommaToken}:                       _GotoState260Action,
	{_State187, RparenToken}:                      _GotoState261Action,
	{_State189, IdentifierToken}:                  _GotoState262Action,
	{_State191, IdentifierToken}:                  _GotoState149Action,
	{_State191, StructToken}:                      _GotoState31Action,
	{_State191, EnumToken}:                        _GotoState111Action,
	{_State191, TraitToken}:                       _GotoState15Action,
	{_State191, LparenToken}:                      _GotoState114Action,
	{_State191, TildeTildeToken}:                  _GotoState116Action,
	{_State191, BitNegToken}:                      _GotoState110Action,
	{_State191, AtomTypeType}:                     _GotoState118Action,
	{_State191, TraitableTypeType}:                _GotoState263Action,
	{_State191, ImplicitStructDefType}:            _GotoState124Action,
	{_State191, ExplicitStructDefType}:            _GotoState120Action,
	{_State191, ImplicitEnumDefType}:              _GotoState123Action,
	{_State191, ExplicitEnumDefType}:              _GotoState119Action,
	{_State191, TraitDefType}:                     _GotoState128Action,
	{_State192, IdentifierToken}:                  _GotoState149Action,
	{_State192, StructToken}:                      _GotoState31Action,
	{_State192, EnumToken}:                        _GotoState111Action,
	{_State192, TraitToken}:                       _GotoState15Action,
	{_State192, LparenToken}:                      _GotoState114Action,
	{_State192, TildeTildeToken}:                  _GotoState116Action,
	{_State192, BitNegToken}:                      _GotoState110Action,
	{_State192, AtomTypeType}:                     _GotoState118Action,
	{_State192, TraitableTypeType}:                _GotoState264Action,
	{_State192, ImplicitStructDefType}:            _GotoState124Action,
	{_State192, ExplicitStructDefType}:            _GotoState120Action,
	{_State192, ImplicitEnumDefType}:              _GotoState123Action,
	{_State192, ExplicitEnumDefType}:              _GotoState119Action,
	{_State192, TraitDefType}:                     _GotoState128Action,
	{_State193, IdentifierToken}:                  _GotoState149Action,
	{_State193, StructToken}:                      _GotoState31Action,
	{_State193, EnumToken}:                        _GotoState111Action,
	{_State193, TraitToken}:                       _GotoState15Action,
	{_State193, LparenToken}:                      _GotoState114Action,
	{_State193, TildeTildeToken}:                  _GotoState116Action,
	{_State193, BitNegToken}:                      _GotoState110Action,
	{_State193, AtomTypeType}:                     _GotoState118Action,
	{_State193, TraitableTypeType}:                _GotoState265Action,
	{_State193, ImplicitStructDefType}:            _GotoState124Action,
	{_State193, ExplicitStructDefType}:            _GotoState120Action,
	{_State193, ImplicitEnumDefType}:              _GotoState123Action,
	{_State193, ExplicitEnumDefType}:              _GotoState119Action,
	{_State193, TraitDefType}:                     _GotoState128Action,
	{_State194, IdentifierToken}:                  _GotoState113Action,
	{_State194, UnsafeToken}:                      _GotoState117Action,
	{_State194, StructToken}:                      _GotoState31Action,
	{_State194, EnumToken}:                        _GotoState111Action,
	{_State194, TraitToken}:                       _GotoState15Action,
	{_State194, FuncToken}:                        _GotoState112Action,
	{_State194, LparenToken}:                      _GotoState114Action,
	{_State194, QuestionToken}:                    _GotoState115Action,
	{_State194, TildeTildeToken}:                  _GotoState116Action,
	{_State194, BitNegToken}:                      _GotoState110Action,
	{_State194, BitAndToken}:                      _GotoState109Action,
	{_State194, UnsafeStatementType}:              _GotoState132Action,
	{_State194, AtomTypeType}:                     _GotoState118Action,
	{_State194, TraitableTypeType}:                _GotoState131Action,
	{_State194, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State194, ValueTypeType}:                    _GotoState133Action,
	{_State194, FieldDefType}:                     _GotoState121Action,
	{_State194, ImplicitStructDefType}:            _GotoState124Action,
	{_State194, ExplicitStructDefType}:            _GotoState120Action,
	{_State194, ImplicitEnumDefType}:              _GotoState123Action,
	{_State194, ExplicitEnumDefType}:              _GotoState119Action,
	{_State194, TraitPropertyType}:                _GotoState266Action,
	{_State194, TraitDefType}:                     _GotoState128Action,
	{_State194, FuncTypeType}:                     _GotoState122Action,
	{_State194, MethodSignatureType}:              _GotoState125Action,
	{_State195, IdentifierToken}:                  _GotoState113Action,
	{_State195, UnsafeToken}:                      _GotoState117Action,
	{_State195, StructToken}:                      _GotoState31Action,
	{_State195, EnumToken}:                        _GotoState111Action,
	{_State195, TraitToken}:                       _GotoState15Action,
	{_State195, FuncToken}:                        _GotoState112Action,
	{_State195, LparenToken}:                      _GotoState114Action,
	{_State195, QuestionToken}:                    _GotoState115Action,
	{_State195, TildeTildeToken}:                  _GotoState116Action,
	{_State195, BitNegToken}:                      _GotoState110Action,
	{_State195, BitAndToken}:                      _GotoState109Action,
	{_State195, UnsafeStatementType}:              _GotoState132Action,
	{_State195, AtomTypeType}:                     _GotoState118Action,
	{_State195, TraitableTypeType}:                _GotoState131Action,
	{_State195, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State195, ValueTypeType}:                    _GotoState133Action,
	{_State195, FieldDefType}:                     _GotoState121Action,
	{_State195, ImplicitStructDefType}:            _GotoState124Action,
	{_State195, ExplicitStructDefType}:            _GotoState120Action,
	{_State195, ImplicitEnumDefType}:              _GotoState123Action,
	{_State195, ExplicitEnumDefType}:              _GotoState119Action,
	{_State195, TraitPropertyType}:                _GotoState267Action,
	{_State195, TraitDefType}:                     _GotoState128Action,
	{_State195, FuncTypeType}:                     _GotoState122Action,
	{_State195, MethodSignatureType}:              _GotoState125Action,
	{_State196, IdentifierToken}:                  _GotoState149Action,
	{_State196, StructToken}:                      _GotoState31Action,
	{_State196, EnumToken}:                        _GotoState111Action,
	{_State196, TraitToken}:                       _GotoState15Action,
	{_State196, FuncToken}:                        _GotoState145Action,
	{_State196, LparenToken}:                      _GotoState114Action,
	{_State196, QuestionToken}:                    _GotoState115Action,
	{_State196, TildeTildeToken}:                  _GotoState116Action,
	{_State196, BitNegToken}:                      _GotoState110Action,
	{_State196, BitAndToken}:                      _GotoState109Action,
	{_State196, AtomTypeType}:                     _GotoState118Action,
	{_State196, TraitableTypeType}:                _GotoState131Action,
	{_State196, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State196, ValueTypeType}:                    _GotoState268Action,
	{_State196, ImplicitStructDefType}:            _GotoState124Action,
	{_State196, ExplicitStructDefType}:            _GotoState120Action,
	{_State196, ImplicitEnumDefType}:              _GotoState123Action,
	{_State196, ExplicitEnumDefType}:              _GotoState119Action,
	{_State196, TraitDefType}:                     _GotoState128Action,
	{_State196, FuncTypeType}:                     _GotoState122Action,
	{_State199, LparenToken}:                      _GotoState269Action,
	{_State200, IdentifierToken}:                  _GotoState149Action,
	{_State200, StructToken}:                      _GotoState31Action,
	{_State200, EnumToken}:                        _GotoState111Action,
	{_State200, TraitToken}:                       _GotoState15Action,
	{_State200, FuncToken}:                        _GotoState145Action,
	{_State200, LparenToken}:                      _GotoState114Action,
	{_State200, QuestionToken}:                    _GotoState115Action,
	{_State200, TildeTildeToken}:                  _GotoState116Action,
	{_State200, BitNegToken}:                      _GotoState110Action,
	{_State200, BitAndToken}:                      _GotoState109Action,
	{_State200, AtomTypeType}:                     _GotoState118Action,
	{_State200, TraitableTypeType}:                _GotoState131Action,
	{_State200, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State200, ValueTypeType}:                    _GotoState271Action,
	{_State200, ImplicitStructDefType}:            _GotoState124Action,
	{_State200, ExplicitStructDefType}:            _GotoState120Action,
	{_State200, ImplicitEnumDefType}:              _GotoState123Action,
	{_State200, ExplicitEnumDefType}:              _GotoState119Action,
	{_State200, TraitDefType}:                     _GotoState128Action,
	{_State200, ReturnTypeType}:                   _GotoState270Action,
	{_State200, FuncTypeType}:                     _GotoState122Action,
	{_State201, IdentifierToken}:                  _GotoState134Action,
	{_State201, ParameterDefType}:                 _GotoState272Action,
	{_State207, IdentifierToken}:                  _GotoState113Action,
	{_State207, UnsafeToken}:                      _GotoState117Action,
	{_State207, StructToken}:                      _GotoState31Action,
	{_State207, EnumToken}:                        _GotoState111Action,
	{_State207, TraitToken}:                       _GotoState15Action,
	{_State207, FuncToken}:                        _GotoState145Action,
	{_State207, LparenToken}:                      _GotoState114Action,
	{_State207, QuestionToken}:                    _GotoState115Action,
	{_State207, TildeTildeToken}:                  _GotoState116Action,
	{_State207, BitNegToken}:                      _GotoState110Action,
	{_State207, BitAndToken}:                      _GotoState109Action,
	{_State207, UnsafeStatementType}:              _GotoState132Action,
	{_State207, AtomTypeType}:                     _GotoState118Action,
	{_State207, TraitableTypeType}:                _GotoState131Action,
	{_State207, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State207, ValueTypeType}:                    _GotoState133Action,
	{_State207, FieldDefType}:                     _GotoState273Action,
	{_State207, ImplicitStructDefType}:            _GotoState124Action,
	{_State207, ExplicitStructDefType}:            _GotoState120Action,
	{_State207, ImplicitEnumDefType}:              _GotoState123Action,
	{_State207, ExplicitEnumDefType}:              _GotoState119Action,
	{_State207, TraitDefType}:                     _GotoState128Action,
	{_State207, FuncTypeType}:                     _GotoState122Action,
	{_State208, IdentifierToken}:                  _GotoState113Action,
	{_State208, UnsafeToken}:                      _GotoState117Action,
	{_State208, StructToken}:                      _GotoState31Action,
	{_State208, EnumToken}:                        _GotoState111Action,
	{_State208, TraitToken}:                       _GotoState15Action,
	{_State208, FuncToken}:                        _GotoState145Action,
	{_State208, LparenToken}:                      _GotoState114Action,
	{_State208, QuestionToken}:                    _GotoState115Action,
	{_State208, TildeTildeToken}:                  _GotoState116Action,
	{_State208, BitNegToken}:                      _GotoState110Action,
	{_State208, BitAndToken}:                      _GotoState109Action,
	{_State208, UnsafeStatementType}:              _GotoState132Action,
	{_State208, AtomTypeType}:                     _GotoState118Action,
	{_State208, TraitableTypeType}:                _GotoState131Action,
	{_State208, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State208, ValueTypeType}:                    _GotoState133Action,
	{_State208, FieldDefType}:                     _GotoState274Action,
	{_State208, ImplicitStructDefType}:            _GotoState124Action,
	{_State208, ExplicitStructDefType}:            _GotoState120Action,
	{_State208, ImplicitEnumDefType}:              _GotoState123Action,
	{_State208, ExplicitEnumDefType}:              _GotoState119Action,
	{_State208, TraitDefType}:                     _GotoState128Action,
	{_State208, FuncTypeType}:                     _GotoState122Action,
	{_State210, IdentifierToken}:                  _GotoState149Action,
	{_State210, StructToken}:                      _GotoState31Action,
	{_State210, EnumToken}:                        _GotoState111Action,
	{_State210, TraitToken}:                       _GotoState15Action,
	{_State210, FuncToken}:                        _GotoState145Action,
	{_State210, LparenToken}:                      _GotoState114Action,
	{_State210, QuestionToken}:                    _GotoState115Action,
	{_State210, TildeTildeToken}:                  _GotoState116Action,
	{_State210, BitNegToken}:                      _GotoState110Action,
	{_State210, BitAndToken}:                      _GotoState109Action,
	{_State210, AtomTypeType}:                     _GotoState118Action,
	{_State210, TraitableTypeType}:                _GotoState131Action,
	{_State210, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State210, ValueTypeType}:                    _GotoState275Action,
	{_State210, ImplicitStructDefType}:            _GotoState124Action,
	{_State210, ExplicitStructDefType}:            _GotoState120Action,
	{_State210, ImplicitEnumDefType}:              _GotoState123Action,
	{_State210, ExplicitEnumDefType}:              _GotoState119Action,
	{_State210, TraitDefType}:                     _GotoState128Action,
	{_State210, FuncTypeType}:                     _GotoState122Action,
	{_State213, CommaToken}:                       _GotoState141Action,
	{_State214, RparenToken}:                      _GotoState276Action,
	{_State216, IntegerLiteralToken}:              _GotoState23Action,
	{_State216, FloatLiteralToken}:                _GotoState20Action,
	{_State216, RuneLiteralToken}:                 _GotoState29Action,
	{_State216, StringLiteralToken}:               _GotoState30Action,
	{_State216, IdentifierToken}:                  _GotoState22Action,
	{_State216, TrueToken}:                        _GotoState33Action,
	{_State216, FalseToken}:                       _GotoState19Action,
	{_State216, StructToken}:                      _GotoState31Action,
	{_State216, FuncToken}:                        _GotoState21Action,
	{_State216, LabelDeclToken}:                   _GotoState24Action,
	{_State216, LparenToken}:                      _GotoState26Action,
	{_State216, NotToken}:                         _GotoState28Action,
	{_State216, SubToken}:                         _GotoState32Action,
	{_State216, MulToken}:                         _GotoState27Action,
	{_State216, BitNegToken}:                      _GotoState18Action,
	{_State216, BitAndToken}:                      _GotoState17Action,
	{_State216, LexErrorToken}:                    _GotoState25Action,
	{_State216, OptionalLabelDeclType}:            _GotoState103Action,
	{_State216, SequenceExprType}:                 _GotoState277Action,
	{_State216, BlockExprType}:                    _GotoState40Action,
	{_State216, CallExprType}:                     _GotoState41Action,
	{_State216, AtomExprType}:                     _GotoState39Action,
	{_State216, LiteralType}:                      _GotoState44Action,
	{_State216, AnonymousStructExprType}:          _GotoState38Action,
	{_State216, AccessExprType}:                   _GotoState34Action,
	{_State216, PostfixUnaryExprType}:             _GotoState48Action,
	{_State216, PrefixUnaryOpType}:                _GotoState50Action,
	{_State216, PrefixUnaryExprType}:              _GotoState49Action,
	{_State216, MulExprType}:                      _GotoState45Action,
	{_State216, AddExprType}:                      _GotoState35Action,
	{_State216, CmpExprType}:                      _GotoState42Action,
	{_State216, AndExprType}:                      _GotoState36Action,
	{_State216, OrExprType}:                       _GotoState47Action,
	{_State216, ExplicitStructDefType}:            _GotoState43Action,
	{_State216, AnonymousFuncExprType}:            _GotoState37Action,
	{_State217, DoToken}:                          _GotoState278Action,
	{_State218, SemicolonToken}:                   _GotoState279Action,
	{_State220, LbraceToken}:                      _GotoState96Action,
	{_State220, BlockBodyType}:                    _GotoState280Action,
	{_State221, ElseToken}:                        _GotoState281Action,
	{_State222, IntegerLiteralToken}:              _GotoState23Action,
	{_State222, FloatLiteralToken}:                _GotoState20Action,
	{_State222, RuneLiteralToken}:                 _GotoState29Action,
	{_State222, StringLiteralToken}:               _GotoState30Action,
	{_State222, IdentifierToken}:                  _GotoState22Action,
	{_State222, TrueToken}:                        _GotoState33Action,
	{_State222, FalseToken}:                       _GotoState19Action,
	{_State222, StructToken}:                      _GotoState31Action,
	{_State222, FuncToken}:                        _GotoState21Action,
	{_State222, LabelDeclToken}:                   _GotoState24Action,
	{_State222, LparenToken}:                      _GotoState26Action,
	{_State222, LexErrorToken}:                    _GotoState25Action,
	{_State222, OptionalLabelDeclType}:            _GotoState103Action,
	{_State222, BlockExprType}:                    _GotoState40Action,
	{_State222, CallExprType}:                     _GotoState283Action,
	{_State222, AtomExprType}:                     _GotoState39Action,
	{_State222, LiteralType}:                      _GotoState44Action,
	{_State222, AnonymousStructExprType}:          _GotoState38Action,
	{_State222, AccessExprType}:                   _GotoState282Action,
	{_State222, ExplicitStructDefType}:            _GotoState43Action,
	{_State222, AnonymousFuncExprType}:            _GotoState37Action,
	{_State225, IntegerLiteralToken}:              _GotoState23Action,
	{_State225, FloatLiteralToken}:                _GotoState20Action,
	{_State225, RuneLiteralToken}:                 _GotoState29Action,
	{_State225, StringLiteralToken}:               _GotoState30Action,
	{_State225, IdentifierToken}:                  _GotoState22Action,
	{_State225, TrueToken}:                        _GotoState33Action,
	{_State225, FalseToken}:                       _GotoState19Action,
	{_State225, StructToken}:                      _GotoState31Action,
	{_State225, FuncToken}:                        _GotoState21Action,
	{_State225, LabelDeclToken}:                   _GotoState24Action,
	{_State225, LparenToken}:                      _GotoState26Action,
	{_State225, LexErrorToken}:                    _GotoState25Action,
	{_State225, OptionalLabelDeclType}:            _GotoState103Action,
	{_State225, BlockExprType}:                    _GotoState40Action,
	{_State225, CallExprType}:                     _GotoState284Action,
	{_State225, AtomExprType}:                     _GotoState39Action,
	{_State225, LiteralType}:                      _GotoState44Action,
	{_State225, AnonymousStructExprType}:          _GotoState38Action,
	{_State225, AccessExprType}:                   _GotoState282Action,
	{_State225, ExplicitStructDefType}:            _GotoState43Action,
	{_State225, AnonymousFuncExprType}:            _GotoState37Action,
	{_State228, LbracketToken}:                    _GotoState69Action,
	{_State228, DotToken}:                         _GotoState68Action,
	{_State228, QuestionToken}:                    _GotoState70Action,
	{_State228, DollarLbracketToken}:              _GotoState67Action,
	{_State228, AddAssignToken}:                   _GotoState285Action,
	{_State228, SubAssignToken}:                   _GotoState296Action,
	{_State228, MulAssignToken}:                   _GotoState295Action,
	{_State228, DivAssignToken}:                   _GotoState293Action,
	{_State228, ModAssignToken}:                   _GotoState294Action,
	{_State228, AddOneAssignToken}:                _GotoState286Action,
	{_State228, SubOneAssignToken}:                _GotoState297Action,
	{_State228, BitNegAssignToken}:                _GotoState289Action,
	{_State228, BitAndAssignToken}:                _GotoState287Action,
	{_State228, BitOrAssignToken}:                 _GotoState290Action,
	{_State228, BitXorAssignToken}:                _GotoState292Action,
	{_State228, BitLshiftAssignToken}:             _GotoState288Action,
	{_State228, BitRshiftAssignToken}:             _GotoState291Action,
	{_State228, UnaryOpAssignType}:                _GotoState299Action,
	{_State228, BinaryOpAssignType}:               _GotoState298Action,
	{_State228, OptionalGenericBindingType}:       _GotoState71Action,
	{_State230, CommaToken}:                       _GotoState300Action,
	{_State232, JumpLabelToken}:                   _GotoState301Action,
	{_State232, OptionalJumpLabelType}:            _GotoState302Action,
	{_State234, NewlinesToken}:                    _GotoState303Action,
	{_State234, SemicolonToken}:                   _GotoState304Action,
	{_State236, CaseToken}:                        _GotoState305Action,
	{_State239, NewlinesToken}:                    _GotoState306Action,
	{_State239, SemicolonToken}:                   _GotoState307Action,
	{_State242, IdentifierToken}:                  _GotoState170Action,
	{_State242, GenericParameterDefType}:          _GotoState308Action,
	{_State244, IdentifierToken}:                  _GotoState149Action,
	{_State244, StructToken}:                      _GotoState31Action,
	{_State244, EnumToken}:                        _GotoState111Action,
	{_State244, TraitToken}:                       _GotoState15Action,
	{_State244, FuncToken}:                        _GotoState145Action,
	{_State244, LparenToken}:                      _GotoState114Action,
	{_State244, QuestionToken}:                    _GotoState115Action,
	{_State244, TildeTildeToken}:                  _GotoState116Action,
	{_State244, BitNegToken}:                      _GotoState110Action,
	{_State244, BitAndToken}:                      _GotoState109Action,
	{_State244, AtomTypeType}:                     _GotoState118Action,
	{_State244, TraitableTypeType}:                _GotoState131Action,
	{_State244, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State244, ValueTypeType}:                    _GotoState309Action,
	{_State244, ImplicitStructDefType}:            _GotoState124Action,
	{_State244, ExplicitStructDefType}:            _GotoState120Action,
	{_State244, ImplicitEnumDefType}:              _GotoState123Action,
	{_State244, ExplicitEnumDefType}:              _GotoState119Action,
	{_State244, TraitDefType}:                     _GotoState128Action,
	{_State244, FuncTypeType}:                     _GotoState122Action,
	{_State245, NewlinesToken}:                    _GotoState310Action,
	{_State245, OrToken}:                          _GotoState311Action,
	{_State246, RparenToken}:                      _GotoState312Action,
	{_State247, AssignToken}:                      _GotoState257Action,
	{_State248, NewlinesToken}:                    _GotoState313Action,
	{_State248, OrToken}:                          _GotoState314Action,
	{_State249, IdentifierToken}:                  _GotoState251Action,
	{_State249, StructToken}:                      _GotoState31Action,
	{_State249, EnumToken}:                        _GotoState111Action,
	{_State249, TraitToken}:                       _GotoState15Action,
	{_State249, FuncToken}:                        _GotoState145Action,
	{_State249, LparenToken}:                      _GotoState114Action,
	{_State249, QuestionToken}:                    _GotoState115Action,
	{_State249, DotdotdotToken}:                   _GotoState250Action,
	{_State249, TildeTildeToken}:                  _GotoState116Action,
	{_State249, BitNegToken}:                      _GotoState110Action,
	{_State249, BitAndToken}:                      _GotoState109Action,
	{_State249, AtomTypeType}:                     _GotoState118Action,
	{_State249, TraitableTypeType}:                _GotoState131Action,
	{_State249, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State249, ValueTypeType}:                    _GotoState255Action,
	{_State249, ImplicitStructDefType}:            _GotoState124Action,
	{_State249, ExplicitStructDefType}:            _GotoState120Action,
	{_State249, ImplicitEnumDefType}:              _GotoState123Action,
	{_State249, ExplicitEnumDefType}:              _GotoState119Action,
	{_State249, TraitDefType}:                     _GotoState128Action,
	{_State249, ParameterDeclType}:                _GotoState253Action,
	{_State249, ParameterDeclsType}:               _GotoState254Action,
	{_State249, OptionalParameterDeclsType}:       _GotoState315Action,
	{_State249, FuncTypeType}:                     _GotoState122Action,
	{_State250, IdentifierToken}:                  _GotoState149Action,
	{_State250, StructToken}:                      _GotoState31Action,
	{_State250, EnumToken}:                        _GotoState111Action,
	{_State250, TraitToken}:                       _GotoState15Action,
	{_State250, FuncToken}:                        _GotoState145Action,
	{_State250, LparenToken}:                      _GotoState114Action,
	{_State250, QuestionToken}:                    _GotoState115Action,
	{_State250, TildeTildeToken}:                  _GotoState116Action,
	{_State250, BitNegToken}:                      _GotoState110Action,
	{_State250, BitAndToken}:                      _GotoState109Action,
	{_State250, AtomTypeType}:                     _GotoState118Action,
	{_State250, TraitableTypeType}:                _GotoState131Action,
	{_State250, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State250, ValueTypeType}:                    _GotoState316Action,
	{_State250, ImplicitStructDefType}:            _GotoState124Action,
	{_State250, ExplicitStructDefType}:            _GotoState120Action,
	{_State250, ImplicitEnumDefType}:              _GotoState123Action,
	{_State250, ExplicitEnumDefType}:              _GotoState119Action,
	{_State250, TraitDefType}:                     _GotoState128Action,
	{_State250, FuncTypeType}:                     _GotoState122Action,
	{_State251, IdentifierToken}:                  _GotoState149Action,
	{_State251, StructToken}:                      _GotoState31Action,
	{_State251, EnumToken}:                        _GotoState111Action,
	{_State251, TraitToken}:                       _GotoState15Action,
	{_State251, FuncToken}:                        _GotoState145Action,
	{_State251, LparenToken}:                      _GotoState114Action,
	{_State251, QuestionToken}:                    _GotoState115Action,
	{_State251, DollarLbracketToken}:              _GotoState67Action,
	{_State251, DotdotdotToken}:                   _GotoState317Action,
	{_State251, TildeTildeToken}:                  _GotoState116Action,
	{_State251, BitNegToken}:                      _GotoState110Action,
	{_State251, BitAndToken}:                      _GotoState109Action,
	{_State251, OptionalGenericBindingType}:       _GotoState181Action,
	{_State251, AtomTypeType}:                     _GotoState118Action,
	{_State251, TraitableTypeType}:                _GotoState131Action,
	{_State251, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State251, ValueTypeType}:                    _GotoState318Action,
	{_State251, ImplicitStructDefType}:            _GotoState124Action,
	{_State251, ExplicitStructDefType}:            _GotoState120Action,
	{_State251, ImplicitEnumDefType}:              _GotoState123Action,
	{_State251, ExplicitEnumDefType}:              _GotoState119Action,
	{_State251, TraitDefType}:                     _GotoState128Action,
	{_State251, FuncTypeType}:                     _GotoState122Action,
	{_State252, RparenToken}:                      _GotoState319Action,
	{_State254, CommaToken}:                       _GotoState320Action,
	{_State256, IdentifierToken}:                  _GotoState113Action,
	{_State256, UnsafeToken}:                      _GotoState117Action,
	{_State256, StructToken}:                      _GotoState31Action,
	{_State256, EnumToken}:                        _GotoState111Action,
	{_State256, TraitToken}:                       _GotoState15Action,
	{_State256, FuncToken}:                        _GotoState145Action,
	{_State256, LparenToken}:                      _GotoState114Action,
	{_State256, QuestionToken}:                    _GotoState115Action,
	{_State256, TildeTildeToken}:                  _GotoState116Action,
	{_State256, BitNegToken}:                      _GotoState110Action,
	{_State256, BitAndToken}:                      _GotoState109Action,
	{_State256, UnsafeStatementType}:              _GotoState132Action,
	{_State256, AtomTypeType}:                     _GotoState118Action,
	{_State256, TraitableTypeType}:                _GotoState131Action,
	{_State256, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State256, ValueTypeType}:                    _GotoState133Action,
	{_State256, FieldDefType}:                     _GotoState247Action,
	{_State256, ImplicitStructDefType}:            _GotoState124Action,
	{_State256, ExplicitStructDefType}:            _GotoState120Action,
	{_State256, EnumValueDefType}:                 _GotoState321Action,
	{_State256, ImplicitEnumDefType}:              _GotoState123Action,
	{_State256, ExplicitEnumDefType}:              _GotoState119Action,
	{_State256, TraitDefType}:                     _GotoState128Action,
	{_State256, FuncTypeType}:                     _GotoState122Action,
	{_State257, DefaultToken}:                     _GotoState322Action,
	{_State258, IdentifierToken}:                  _GotoState113Action,
	{_State258, UnsafeToken}:                      _GotoState117Action,
	{_State258, StructToken}:                      _GotoState31Action,
	{_State258, EnumToken}:                        _GotoState111Action,
	{_State258, TraitToken}:                       _GotoState15Action,
	{_State258, FuncToken}:                        _GotoState145Action,
	{_State258, LparenToken}:                      _GotoState114Action,
	{_State258, QuestionToken}:                    _GotoState115Action,
	{_State258, TildeTildeToken}:                  _GotoState116Action,
	{_State258, BitNegToken}:                      _GotoState110Action,
	{_State258, BitAndToken}:                      _GotoState109Action,
	{_State258, UnsafeStatementType}:              _GotoState132Action,
	{_State258, AtomTypeType}:                     _GotoState118Action,
	{_State258, TraitableTypeType}:                _GotoState131Action,
	{_State258, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State258, ValueTypeType}:                    _GotoState133Action,
	{_State258, FieldDefType}:                     _GotoState247Action,
	{_State258, ImplicitStructDefType}:            _GotoState124Action,
	{_State258, ExplicitStructDefType}:            _GotoState120Action,
	{_State258, EnumValueDefType}:                 _GotoState323Action,
	{_State258, ImplicitEnumDefType}:              _GotoState123Action,
	{_State258, ExplicitEnumDefType}:              _GotoState119Action,
	{_State258, TraitDefType}:                     _GotoState128Action,
	{_State258, FuncTypeType}:                     _GotoState122Action,
	{_State260, IdentifierToken}:                  _GotoState113Action,
	{_State260, UnsafeToken}:                      _GotoState117Action,
	{_State260, StructToken}:                      _GotoState31Action,
	{_State260, EnumToken}:                        _GotoState111Action,
	{_State260, TraitToken}:                       _GotoState15Action,
	{_State260, FuncToken}:                        _GotoState145Action,
	{_State260, LparenToken}:                      _GotoState114Action,
	{_State260, QuestionToken}:                    _GotoState115Action,
	{_State260, TildeTildeToken}:                  _GotoState116Action,
	{_State260, BitNegToken}:                      _GotoState110Action,
	{_State260, BitAndToken}:                      _GotoState109Action,
	{_State260, UnsafeStatementType}:              _GotoState132Action,
	{_State260, AtomTypeType}:                     _GotoState118Action,
	{_State260, TraitableTypeType}:                _GotoState131Action,
	{_State260, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State260, ValueTypeType}:                    _GotoState133Action,
	{_State260, FieldDefType}:                     _GotoState324Action,
	{_State260, ImplicitStructDefType}:            _GotoState124Action,
	{_State260, ExplicitStructDefType}:            _GotoState120Action,
	{_State260, ImplicitEnumDefType}:              _GotoState123Action,
	{_State260, ExplicitEnumDefType}:              _GotoState119Action,
	{_State260, TraitDefType}:                     _GotoState128Action,
	{_State260, FuncTypeType}:                     _GotoState122Action,
	{_State262, GreaterToken}:                     _GotoState325Action,
	{_State269, IdentifierToken}:                  _GotoState134Action,
	{_State269, ParameterDefType}:                 _GotoState138Action,
	{_State269, ParameterDefsType}:                _GotoState139Action,
	{_State269, OptionalParameterDefsType}:        _GotoState326Action,
	{_State270, LbraceToken}:                      _GotoState96Action,
	{_State270, BlockBodyType}:                    _GotoState327Action,
	{_State278, LbraceToken}:                      _GotoState96Action,
	{_State278, BlockBodyType}:                    _GotoState328Action,
	{_State279, IntegerLiteralToken}:              _GotoState23Action,
	{_State279, FloatLiteralToken}:                _GotoState20Action,
	{_State279, RuneLiteralToken}:                 _GotoState29Action,
	{_State279, StringLiteralToken}:               _GotoState30Action,
	{_State279, IdentifierToken}:                  _GotoState22Action,
	{_State279, TrueToken}:                        _GotoState33Action,
	{_State279, FalseToken}:                       _GotoState19Action,
	{_State279, StructToken}:                      _GotoState31Action,
	{_State279, FuncToken}:                        _GotoState21Action,
	{_State279, LabelDeclToken}:                   _GotoState24Action,
	{_State279, LparenToken}:                      _GotoState26Action,
	{_State279, NotToken}:                         _GotoState28Action,
	{_State279, SubToken}:                         _GotoState32Action,
	{_State279, MulToken}:                         _GotoState27Action,
	{_State279, BitNegToken}:                      _GotoState18Action,
	{_State279, BitAndToken}:                      _GotoState17Action,
	{_State279, LexErrorToken}:                    _GotoState25Action,
	{_State279, OptionalLabelDeclType}:            _GotoState103Action,
	{_State279, OptionalSequenceExprType}:         _GotoState329Action,
	{_State279, SequenceExprType}:                 _GotoState219Action,
	{_State279, BlockExprType}:                    _GotoState40Action,
	{_State279, CallExprType}:                     _GotoState41Action,
	{_State279, AtomExprType}:                     _GotoState39Action,
	{_State279, LiteralType}:                      _GotoState44Action,
	{_State279, AnonymousStructExprType}:          _GotoState38Action,
	{_State279, AccessExprType}:                   _GotoState34Action,
	{_State279, PostfixUnaryExprType}:             _GotoState48Action,
	{_State279, PrefixUnaryOpType}:                _GotoState50Action,
	{_State279, PrefixUnaryExprType}:              _GotoState49Action,
	{_State279, MulExprType}:                      _GotoState45Action,
	{_State279, AddExprType}:                      _GotoState35Action,
	{_State279, CmpExprType}:                      _GotoState42Action,
	{_State279, AndExprType}:                      _GotoState36Action,
	{_State279, OrExprType}:                       _GotoState47Action,
	{_State279, ExplicitStructDefType}:            _GotoState43Action,
	{_State279, AnonymousFuncExprType}:            _GotoState37Action,
	{_State281, IfToken}:                          _GotoState95Action,
	{_State281, LbraceToken}:                      _GotoState96Action,
	{_State281, IfExprType}:                       _GotoState331Action,
	{_State281, BlockBodyType}:                    _GotoState330Action,
	{_State282, LbracketToken}:                    _GotoState69Action,
	{_State282, DotToken}:                         _GotoState68Action,
	{_State282, DollarLbracketToken}:              _GotoState67Action,
	{_State282, OptionalGenericBindingType}:       _GotoState71Action,
	{_State298, IntegerLiteralToken}:              _GotoState23Action,
	{_State298, FloatLiteralToken}:                _GotoState20Action,
	{_State298, RuneLiteralToken}:                 _GotoState29Action,
	{_State298, StringLiteralToken}:               _GotoState30Action,
	{_State298, IdentifierToken}:                  _GotoState22Action,
	{_State298, TrueToken}:                        _GotoState33Action,
	{_State298, FalseToken}:                       _GotoState19Action,
	{_State298, StructToken}:                      _GotoState31Action,
	{_State298, FuncToken}:                        _GotoState21Action,
	{_State298, LabelDeclToken}:                   _GotoState24Action,
	{_State298, LparenToken}:                      _GotoState26Action,
	{_State298, NotToken}:                         _GotoState28Action,
	{_State298, SubToken}:                         _GotoState32Action,
	{_State298, MulToken}:                         _GotoState27Action,
	{_State298, BitNegToken}:                      _GotoState18Action,
	{_State298, BitAndToken}:                      _GotoState17Action,
	{_State298, LexErrorToken}:                    _GotoState25Action,
	{_State298, ExpressionType}:                   _GotoState332Action,
	{_State298, OptionalLabelDeclType}:            _GotoState46Action,
	{_State298, SequenceExprType}:                 _GotoState51Action,
	{_State298, BlockExprType}:                    _GotoState40Action,
	{_State298, CallExprType}:                     _GotoState41Action,
	{_State298, AtomExprType}:                     _GotoState39Action,
	{_State298, LiteralType}:                      _GotoState44Action,
	{_State298, AnonymousStructExprType}:          _GotoState38Action,
	{_State298, AccessExprType}:                   _GotoState34Action,
	{_State298, PostfixUnaryExprType}:             _GotoState48Action,
	{_State298, PrefixUnaryOpType}:                _GotoState50Action,
	{_State298, PrefixUnaryExprType}:              _GotoState49Action,
	{_State298, MulExprType}:                      _GotoState45Action,
	{_State298, AddExprType}:                      _GotoState35Action,
	{_State298, CmpExprType}:                      _GotoState42Action,
	{_State298, AndExprType}:                      _GotoState36Action,
	{_State298, OrExprType}:                       _GotoState47Action,
	{_State298, ExplicitStructDefType}:            _GotoState43Action,
	{_State298, AnonymousFuncExprType}:            _GotoState37Action,
	{_State300, IntegerLiteralToken}:              _GotoState23Action,
	{_State300, FloatLiteralToken}:                _GotoState20Action,
	{_State300, RuneLiteralToken}:                 _GotoState29Action,
	{_State300, StringLiteralToken}:               _GotoState30Action,
	{_State300, IdentifierToken}:                  _GotoState22Action,
	{_State300, TrueToken}:                        _GotoState33Action,
	{_State300, FalseToken}:                       _GotoState19Action,
	{_State300, StructToken}:                      _GotoState31Action,
	{_State300, FuncToken}:                        _GotoState21Action,
	{_State300, LabelDeclToken}:                   _GotoState24Action,
	{_State300, LparenToken}:                      _GotoState26Action,
	{_State300, NotToken}:                         _GotoState28Action,
	{_State300, SubToken}:                         _GotoState32Action,
	{_State300, MulToken}:                         _GotoState27Action,
	{_State300, BitNegToken}:                      _GotoState18Action,
	{_State300, BitAndToken}:                      _GotoState17Action,
	{_State300, LexErrorToken}:                    _GotoState25Action,
	{_State300, ExpressionType}:                   _GotoState333Action,
	{_State300, OptionalLabelDeclType}:            _GotoState46Action,
	{_State300, SequenceExprType}:                 _GotoState51Action,
	{_State300, BlockExprType}:                    _GotoState40Action,
	{_State300, CallExprType}:                     _GotoState41Action,
	{_State300, AtomExprType}:                     _GotoState39Action,
	{_State300, LiteralType}:                      _GotoState44Action,
	{_State300, AnonymousStructExprType}:          _GotoState38Action,
	{_State300, AccessExprType}:                   _GotoState34Action,
	{_State300, PostfixUnaryExprType}:             _GotoState48Action,
	{_State300, PrefixUnaryOpType}:                _GotoState50Action,
	{_State300, PrefixUnaryExprType}:              _GotoState49Action,
	{_State300, MulExprType}:                      _GotoState45Action,
	{_State300, AddExprType}:                      _GotoState35Action,
	{_State300, CmpExprType}:                      _GotoState42Action,
	{_State300, AndExprType}:                      _GotoState36Action,
	{_State300, OrExprType}:                       _GotoState47Action,
	{_State300, ExplicitStructDefType}:            _GotoState43Action,
	{_State300, AnonymousFuncExprType}:            _GotoState37Action,
	{_State302, IntegerLiteralToken}:              _GotoState23Action,
	{_State302, FloatLiteralToken}:                _GotoState20Action,
	{_State302, RuneLiteralToken}:                 _GotoState29Action,
	{_State302, StringLiteralToken}:               _GotoState30Action,
	{_State302, IdentifierToken}:                  _GotoState22Action,
	{_State302, TrueToken}:                        _GotoState33Action,
	{_State302, FalseToken}:                       _GotoState19Action,
	{_State302, StructToken}:                      _GotoState31Action,
	{_State302, FuncToken}:                        _GotoState21Action,
	{_State302, LabelDeclToken}:                   _GotoState24Action,
	{_State302, LparenToken}:                      _GotoState26Action,
	{_State302, NotToken}:                         _GotoState28Action,
	{_State302, SubToken}:                         _GotoState32Action,
	{_State302, MulToken}:                         _GotoState27Action,
	{_State302, BitNegToken}:                      _GotoState18Action,
	{_State302, BitAndToken}:                      _GotoState17Action,
	{_State302, LexErrorToken}:                    _GotoState25Action,
	{_State302, ExpressionType}:                   _GotoState229Action,
	{_State302, OptionalLabelDeclType}:            _GotoState46Action,
	{_State302, SequenceExprType}:                 _GotoState51Action,
	{_State302, BlockExprType}:                    _GotoState40Action,
	{_State302, ExpressionsType}:                  _GotoState334Action,
	{_State302, OptionalExpressionsType}:          _GotoState335Action,
	{_State302, CallExprType}:                     _GotoState41Action,
	{_State302, AtomExprType}:                     _GotoState39Action,
	{_State302, LiteralType}:                      _GotoState44Action,
	{_State302, AnonymousStructExprType}:          _GotoState38Action,
	{_State302, AccessExprType}:                   _GotoState34Action,
	{_State302, PostfixUnaryExprType}:             _GotoState48Action,
	{_State302, PrefixUnaryOpType}:                _GotoState50Action,
	{_State302, PrefixUnaryExprType}:              _GotoState49Action,
	{_State302, MulExprType}:                      _GotoState45Action,
	{_State302, AddExprType}:                      _GotoState35Action,
	{_State302, CmpExprType}:                      _GotoState42Action,
	{_State302, AndExprType}:                      _GotoState36Action,
	{_State302, OrExprType}:                       _GotoState47Action,
	{_State302, ExplicitStructDefType}:            _GotoState43Action,
	{_State302, AnonymousFuncExprType}:            _GotoState37Action,
	{_State305, DefaultToken}:                     _GotoState336Action,
	{_State310, IdentifierToken}:                  _GotoState113Action,
	{_State310, UnsafeToken}:                      _GotoState117Action,
	{_State310, StructToken}:                      _GotoState31Action,
	{_State310, EnumToken}:                        _GotoState111Action,
	{_State310, TraitToken}:                       _GotoState15Action,
	{_State310, FuncToken}:                        _GotoState145Action,
	{_State310, LparenToken}:                      _GotoState114Action,
	{_State310, QuestionToken}:                    _GotoState115Action,
	{_State310, TildeTildeToken}:                  _GotoState116Action,
	{_State310, BitNegToken}:                      _GotoState110Action,
	{_State310, BitAndToken}:                      _GotoState109Action,
	{_State310, UnsafeStatementType}:              _GotoState132Action,
	{_State310, AtomTypeType}:                     _GotoState118Action,
	{_State310, TraitableTypeType}:                _GotoState131Action,
	{_State310, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State310, ValueTypeType}:                    _GotoState133Action,
	{_State310, FieldDefType}:                     _GotoState247Action,
	{_State310, ImplicitStructDefType}:            _GotoState124Action,
	{_State310, ExplicitStructDefType}:            _GotoState120Action,
	{_State310, EnumValueDefType}:                 _GotoState337Action,
	{_State310, ImplicitEnumDefType}:              _GotoState123Action,
	{_State310, ExplicitEnumDefType}:              _GotoState119Action,
	{_State310, TraitDefType}:                     _GotoState128Action,
	{_State310, FuncTypeType}:                     _GotoState122Action,
	{_State311, IdentifierToken}:                  _GotoState113Action,
	{_State311, UnsafeToken}:                      _GotoState117Action,
	{_State311, StructToken}:                      _GotoState31Action,
	{_State311, EnumToken}:                        _GotoState111Action,
	{_State311, TraitToken}:                       _GotoState15Action,
	{_State311, FuncToken}:                        _GotoState145Action,
	{_State311, LparenToken}:                      _GotoState114Action,
	{_State311, QuestionToken}:                    _GotoState115Action,
	{_State311, TildeTildeToken}:                  _GotoState116Action,
	{_State311, BitNegToken}:                      _GotoState110Action,
	{_State311, BitAndToken}:                      _GotoState109Action,
	{_State311, UnsafeStatementType}:              _GotoState132Action,
	{_State311, AtomTypeType}:                     _GotoState118Action,
	{_State311, TraitableTypeType}:                _GotoState131Action,
	{_State311, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State311, ValueTypeType}:                    _GotoState133Action,
	{_State311, FieldDefType}:                     _GotoState247Action,
	{_State311, ImplicitStructDefType}:            _GotoState124Action,
	{_State311, ExplicitStructDefType}:            _GotoState120Action,
	{_State311, EnumValueDefType}:                 _GotoState338Action,
	{_State311, ImplicitEnumDefType}:              _GotoState123Action,
	{_State311, ExplicitEnumDefType}:              _GotoState119Action,
	{_State311, TraitDefType}:                     _GotoState128Action,
	{_State311, FuncTypeType}:                     _GotoState122Action,
	{_State313, IdentifierToken}:                  _GotoState113Action,
	{_State313, UnsafeToken}:                      _GotoState117Action,
	{_State313, StructToken}:                      _GotoState31Action,
	{_State313, EnumToken}:                        _GotoState111Action,
	{_State313, TraitToken}:                       _GotoState15Action,
	{_State313, FuncToken}:                        _GotoState145Action,
	{_State313, LparenToken}:                      _GotoState114Action,
	{_State313, QuestionToken}:                    _GotoState115Action,
	{_State313, TildeTildeToken}:                  _GotoState116Action,
	{_State313, BitNegToken}:                      _GotoState110Action,
	{_State313, BitAndToken}:                      _GotoState109Action,
	{_State313, UnsafeStatementType}:              _GotoState132Action,
	{_State313, AtomTypeType}:                     _GotoState118Action,
	{_State313, TraitableTypeType}:                _GotoState131Action,
	{_State313, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State313, ValueTypeType}:                    _GotoState133Action,
	{_State313, FieldDefType}:                     _GotoState247Action,
	{_State313, ImplicitStructDefType}:            _GotoState124Action,
	{_State313, ExplicitStructDefType}:            _GotoState120Action,
	{_State313, EnumValueDefType}:                 _GotoState339Action,
	{_State313, ImplicitEnumDefType}:              _GotoState123Action,
	{_State313, ExplicitEnumDefType}:              _GotoState119Action,
	{_State313, TraitDefType}:                     _GotoState128Action,
	{_State313, FuncTypeType}:                     _GotoState122Action,
	{_State314, IdentifierToken}:                  _GotoState113Action,
	{_State314, UnsafeToken}:                      _GotoState117Action,
	{_State314, StructToken}:                      _GotoState31Action,
	{_State314, EnumToken}:                        _GotoState111Action,
	{_State314, TraitToken}:                       _GotoState15Action,
	{_State314, FuncToken}:                        _GotoState145Action,
	{_State314, LparenToken}:                      _GotoState114Action,
	{_State314, QuestionToken}:                    _GotoState115Action,
	{_State314, TildeTildeToken}:                  _GotoState116Action,
	{_State314, BitNegToken}:                      _GotoState110Action,
	{_State314, BitAndToken}:                      _GotoState109Action,
	{_State314, UnsafeStatementType}:              _GotoState132Action,
	{_State314, AtomTypeType}:                     _GotoState118Action,
	{_State314, TraitableTypeType}:                _GotoState131Action,
	{_State314, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State314, ValueTypeType}:                    _GotoState133Action,
	{_State314, FieldDefType}:                     _GotoState247Action,
	{_State314, ImplicitStructDefType}:            _GotoState124Action,
	{_State314, ExplicitStructDefType}:            _GotoState120Action,
	{_State314, EnumValueDefType}:                 _GotoState340Action,
	{_State314, ImplicitEnumDefType}:              _GotoState123Action,
	{_State314, ExplicitEnumDefType}:              _GotoState119Action,
	{_State314, TraitDefType}:                     _GotoState128Action,
	{_State314, FuncTypeType}:                     _GotoState122Action,
	{_State315, RparenToken}:                      _GotoState341Action,
	{_State317, IdentifierToken}:                  _GotoState149Action,
	{_State317, StructToken}:                      _GotoState31Action,
	{_State317, EnumToken}:                        _GotoState111Action,
	{_State317, TraitToken}:                       _GotoState15Action,
	{_State317, FuncToken}:                        _GotoState145Action,
	{_State317, LparenToken}:                      _GotoState114Action,
	{_State317, QuestionToken}:                    _GotoState115Action,
	{_State317, TildeTildeToken}:                  _GotoState116Action,
	{_State317, BitNegToken}:                      _GotoState110Action,
	{_State317, BitAndToken}:                      _GotoState109Action,
	{_State317, AtomTypeType}:                     _GotoState118Action,
	{_State317, TraitableTypeType}:                _GotoState131Action,
	{_State317, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State317, ValueTypeType}:                    _GotoState342Action,
	{_State317, ImplicitStructDefType}:            _GotoState124Action,
	{_State317, ExplicitStructDefType}:            _GotoState120Action,
	{_State317, ImplicitEnumDefType}:              _GotoState123Action,
	{_State317, ExplicitEnumDefType}:              _GotoState119Action,
	{_State317, TraitDefType}:                     _GotoState128Action,
	{_State317, FuncTypeType}:                     _GotoState122Action,
	{_State319, IdentifierToken}:                  _GotoState149Action,
	{_State319, StructToken}:                      _GotoState31Action,
	{_State319, EnumToken}:                        _GotoState111Action,
	{_State319, TraitToken}:                       _GotoState15Action,
	{_State319, FuncToken}:                        _GotoState145Action,
	{_State319, LparenToken}:                      _GotoState114Action,
	{_State319, QuestionToken}:                    _GotoState115Action,
	{_State319, TildeTildeToken}:                  _GotoState116Action,
	{_State319, BitNegToken}:                      _GotoState110Action,
	{_State319, BitAndToken}:                      _GotoState109Action,
	{_State319, AtomTypeType}:                     _GotoState118Action,
	{_State319, TraitableTypeType}:                _GotoState131Action,
	{_State319, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State319, ValueTypeType}:                    _GotoState271Action,
	{_State319, ImplicitStructDefType}:            _GotoState124Action,
	{_State319, ExplicitStructDefType}:            _GotoState120Action,
	{_State319, ImplicitEnumDefType}:              _GotoState123Action,
	{_State319, ExplicitEnumDefType}:              _GotoState119Action,
	{_State319, TraitDefType}:                     _GotoState128Action,
	{_State319, ReturnTypeType}:                   _GotoState343Action,
	{_State319, FuncTypeType}:                     _GotoState122Action,
	{_State320, IdentifierToken}:                  _GotoState251Action,
	{_State320, StructToken}:                      _GotoState31Action,
	{_State320, EnumToken}:                        _GotoState111Action,
	{_State320, TraitToken}:                       _GotoState15Action,
	{_State320, FuncToken}:                        _GotoState145Action,
	{_State320, LparenToken}:                      _GotoState114Action,
	{_State320, QuestionToken}:                    _GotoState115Action,
	{_State320, DotdotdotToken}:                   _GotoState250Action,
	{_State320, TildeTildeToken}:                  _GotoState116Action,
	{_State320, BitNegToken}:                      _GotoState110Action,
	{_State320, BitAndToken}:                      _GotoState109Action,
	{_State320, AtomTypeType}:                     _GotoState118Action,
	{_State320, TraitableTypeType}:                _GotoState131Action,
	{_State320, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State320, ValueTypeType}:                    _GotoState255Action,
	{_State320, ImplicitStructDefType}:            _GotoState124Action,
	{_State320, ExplicitStructDefType}:            _GotoState120Action,
	{_State320, ImplicitEnumDefType}:              _GotoState123Action,
	{_State320, ExplicitEnumDefType}:              _GotoState119Action,
	{_State320, TraitDefType}:                     _GotoState128Action,
	{_State320, ParameterDeclType}:                _GotoState344Action,
	{_State320, FuncTypeType}:                     _GotoState122Action,
	{_State325, StringLiteralToken}:               _GotoState345Action,
	{_State326, RparenToken}:                      _GotoState346Action,
	{_State329, DoToken}:                          _GotoState347Action,
	{_State334, CommaToken}:                       _GotoState300Action,
	{_State336, RbraceToken}:                      _GotoState348Action,
	{_State341, IdentifierToken}:                  _GotoState149Action,
	{_State341, StructToken}:                      _GotoState31Action,
	{_State341, EnumToken}:                        _GotoState111Action,
	{_State341, TraitToken}:                       _GotoState15Action,
	{_State341, FuncToken}:                        _GotoState145Action,
	{_State341, LparenToken}:                      _GotoState114Action,
	{_State341, QuestionToken}:                    _GotoState115Action,
	{_State341, TildeTildeToken}:                  _GotoState116Action,
	{_State341, BitNegToken}:                      _GotoState110Action,
	{_State341, BitAndToken}:                      _GotoState109Action,
	{_State341, AtomTypeType}:                     _GotoState118Action,
	{_State341, TraitableTypeType}:                _GotoState131Action,
	{_State341, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State341, ValueTypeType}:                    _GotoState271Action,
	{_State341, ImplicitStructDefType}:            _GotoState124Action,
	{_State341, ExplicitStructDefType}:            _GotoState120Action,
	{_State341, ImplicitEnumDefType}:              _GotoState123Action,
	{_State341, ExplicitEnumDefType}:              _GotoState119Action,
	{_State341, TraitDefType}:                     _GotoState128Action,
	{_State341, ReturnTypeType}:                   _GotoState349Action,
	{_State341, FuncTypeType}:                     _GotoState122Action,
	{_State346, IdentifierToken}:                  _GotoState149Action,
	{_State346, StructToken}:                      _GotoState31Action,
	{_State346, EnumToken}:                        _GotoState111Action,
	{_State346, TraitToken}:                       _GotoState15Action,
	{_State346, FuncToken}:                        _GotoState145Action,
	{_State346, LparenToken}:                      _GotoState114Action,
	{_State346, QuestionToken}:                    _GotoState115Action,
	{_State346, TildeTildeToken}:                  _GotoState116Action,
	{_State346, BitNegToken}:                      _GotoState110Action,
	{_State346, BitAndToken}:                      _GotoState109Action,
	{_State346, AtomTypeType}:                     _GotoState118Action,
	{_State346, TraitableTypeType}:                _GotoState131Action,
	{_State346, TraitAlgebraTypeType}:             _GotoState127Action,
	{_State346, ValueTypeType}:                    _GotoState271Action,
	{_State346, ImplicitStructDefType}:            _GotoState124Action,
	{_State346, ExplicitStructDefType}:            _GotoState120Action,
	{_State346, ImplicitEnumDefType}:              _GotoState123Action,
	{_State346, ExplicitEnumDefType}:              _GotoState119Action,
	{_State346, TraitDefType}:                     _GotoState128Action,
	{_State346, ReturnTypeType}:                   _GotoState350Action,
	{_State346, FuncTypeType}:                     _GotoState122Action,
	{_State347, LbraceToken}:                      _GotoState96Action,
	{_State347, BlockBodyType}:                    _GotoState351Action,
	{_State350, LbraceToken}:                      _GotoState96Action,
	{_State350, BlockBodyType}:                    _GotoState352Action,
	{_State5, _WildcardMarker}:                    _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State16, IdentifierToken}:                   _ReduceNilToOptionalReceiverAction,
	{_State17, _WildcardMarker}:                   _ReduceBitAndToPrefixUnaryOpAction,
	{_State18, _WildcardMarker}:                   _ReduceBitNegToPrefixUnaryOpAction,
	{_State19, _WildcardMarker}:                   _ReduceFalseToLiteralAction,
	{_State20, _WildcardMarker}:                   _ReduceFloatLiteralToLiteralAction,
	{_State22, _WildcardMarker}:                   _ReduceIdentifierToAtomExprAction,
	{_State23, _WildcardMarker}:                   _ReduceIntegerLiteralToLiteralAction,
	{_State24, _WildcardMarker}:                   _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State25, _WildcardMarker}:                   _ReduceLexErrorToAtomExprAction,
	{_State26, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State26, ColonToken}:                        _ReduceNilToOptionalExpressionAction,
	{_State27, _WildcardMarker}:                   _ReduceMulToPrefixUnaryOpAction,
	{_State28, _WildcardMarker}:                   _ReduceNotToPrefixUnaryOpAction,
	{_State29, _WildcardMarker}:                   _ReduceRuneLiteralToLiteralAction,
	{_State30, _WildcardMarker}:                   _ReduceStringLiteralToLiteralAction,
	{_State32, _WildcardMarker}:                   _ReduceSubToPrefixUnaryOpAction,
	{_State33, _WildcardMarker}:                   _ReduceTrueToLiteralAction,
	{_State34, _WildcardMarker}:                   _ReduceAccessExprToPostfixUnaryExprAction,
	{_State34, LparenToken}:                       _ReduceNilToOptionalGenericBindingAction,
	{_State35, _WildcardMarker}:                   _ReduceAddExprToCmpExprAction,
	{_State36, _WildcardMarker}:                   _ReduceAndExprToOrExprAction,
	{_State37, _WildcardMarker}:                   _ReduceAnonymousFuncExprToAtomExprAction,
	{_State38, _WildcardMarker}:                   _ReduceAnonymousStructExprToAtomExprAction,
	{_State39, _WildcardMarker}:                   _ReduceAtomExprToAccessExprAction,
	{_State40, _WildcardMarker}:                   _ReduceBlockExprToAtomExprAction,
	{_State41, _WildcardMarker}:                   _ReduceCallExprToAccessExprAction,
	{_State42, _WildcardMarker}:                   _ReduceCmpExprToAndExprAction,
	{_State44, _WildcardMarker}:                   _ReduceLiteralToAtomExprAction,
	{_State45, _WildcardMarker}:                   _ReduceMulExprToAddExprAction,
	{_State47, _WildcardMarker}:                   _ReduceToSequenceExprAction,
	{_State48, _WildcardMarker}:                   _ReducePostfixUnaryExprToPrefixUnaryExprAction,
	{_State49, _WildcardMarker}:                   _ReducePrefixUnaryExprToMulExprAction,
	{_State50, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State51, _EndMarker}:                        _ReduceSequenceExprToExpressionAction,
	{_State52, _EndMarker}:                        _ReduceCommentToLexInternalTokensAction,
	{_State53, _EndMarker}:                        _ReduceSpacesToLexInternalTokensAction,
	{_State54, _EndMarker}:                        _ReduceNoSpecToPackageDefAction,
	{_State55, _WildcardMarker}:                   _ReduceNilToOptionalGenericParametersAction,
	{_State56, RparenToken}:                       _ReduceNilToOptionalTraitPropertiesAction,
	{_State59, RparenToken}:                       _ReduceNilToOptionalParameterDefsAction,
	{_State60, _WildcardMarker}:                   _ReduceIdentifierToAtomExprAction,
	{_State61, _WildcardMarker}:                   _ReduceArgumentToArgumentsAction,
	{_State63, _WildcardMarker}:                   _ReduceColonExpressionsToArgumentAction,
	{_State64, _WildcardMarker}:                   _ReducePositionalToArgumentAction,
	{_State64, ColonToken}:                        _ReduceExpressionToOptionalExpressionAction,
	{_State66, RparenToken}:                       _ReduceNilToOptionalExplicitFieldDefsAction,
	{_State67, RbracketToken}:                     _ReduceNilToOptionalGenericArgumentsAction,
	{_State69, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State69, ColonToken}:                        _ReduceNilToOptionalExpressionAction,
	{_State70, _WildcardMarker}:                   _ReduceQuestionToPostfixUnaryExprAction,
	{_State72, _WildcardMarker}:                   _ReduceAddToAddOpAction,
	{_State73, _WildcardMarker}:                   _ReduceBitOrToAddOpAction,
	{_State74, _WildcardMarker}:                   _ReduceBitXorToAddOpAction,
	{_State75, _WildcardMarker}:                   _ReduceSubToAddOpAction,
	{_State76, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State77, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State78, _WildcardMarker}:                   _ReduceEqualToCmpOpAction,
	{_State79, _WildcardMarker}:                   _ReduceGreaterToCmpOpAction,
	{_State80, _WildcardMarker}:                   _ReduceGreaterOrEqualToCmpOpAction,
	{_State81, _WildcardMarker}:                   _ReduceLessToCmpOpAction,
	{_State82, _WildcardMarker}:                   _ReduceLessOrEqualToCmpOpAction,
	{_State83, _WildcardMarker}:                   _ReduceNotEqualToCmpOpAction,
	{_State84, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State85, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State85, ColonToken}:                        _ReduceNilToOptionalExpressionAction,
	{_State86, _WildcardMarker}:                   _ReduceBitAndToMulOpAction,
	{_State87, _WildcardMarker}:                   _ReduceBitLshiftToMulOpAction,
	{_State88, _WildcardMarker}:                   _ReduceBitRshiftToMulOpAction,
	{_State89, _WildcardMarker}:                   _ReduceDivToMulOpAction,
	{_State90, _WildcardMarker}:                   _ReduceModToMulOpAction,
	{_State91, _WildcardMarker}:                   _ReduceMulToMulOpAction,
	{_State92, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State94, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State95, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State96, _WildcardMarker}:                   _ReduceEmptyListToStatementsAction,
	{_State97, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State98, _WildcardMarker}:                   _ReduceToBlockExprAction,
	{_State99, _EndMarker}:                        _ReduceIfExprToExpressionAction,
	{_State100, _EndMarker}:                       _ReduceLoopExprToExpressionAction,
	{_State101, _EndMarker}:                       _ReduceSwitchExprToExpressionAction,
	{_State102, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State104, _WildcardMarker}:                  _ReducePrefixOpToPrefixUnaryExprAction,
	{_State105, _WildcardMarker}:                  _ReduceEmptyListToPackageStatementsAction,
	{_State106, RbracketToken}:                    _ReduceNilToOptionalGenericParameterDefsAction,
	{_State113, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State114, RparenToken}:                      _ReduceNilToOptionalImplicitFieldDefsAction,
	{_State115, _EndMarker}:                       _ReduceInferredToValueTypeAction,
	{_State118, _WildcardMarker}:                  _ReduceAtomTypeToTraitableTypeAction,
	{_State119, _WildcardMarker}:                  _ReduceExplicitEnumDefToAtomTypeAction,
	{_State120, _WildcardMarker}:                  _ReduceExplicitStructDefToAtomTypeAction,
	{_State121, _WildcardMarker}:                  _ReduceFieldDefToTraitPropertyAction,
	{_State122, _EndMarker}:                       _ReduceFuncTypeToValueTypeAction,
	{_State123, _WildcardMarker}:                  _ReduceImplicitEnumDefToAtomTypeAction,
	{_State124, _WildcardMarker}:                  _ReduceImplicitStructDefToAtomTypeAction,
	{_State125, _WildcardMarker}:                  _ReduceMethodSignatureToTraitPropertyAction,
	{_State127, _EndMarker}:                       _ReduceTraitAlgebraTypeToValueTypeAction,
	{_State128, _WildcardMarker}:                  _ReduceTraitDefToAtomTypeAction,
	{_State129, RparenToken}:                      _ReduceTraitPropertiesToOptionalTraitPropertiesAction,
	{_State130, _WildcardMarker}:                  _ReduceTraitPropertyToTraitPropertiesAction,
	{_State131, _WildcardMarker}:                  _ReduceTraitableTypeToTraitAlgebraTypeAction,
	{_State132, _WildcardMarker}:                  _ReduceUnsafeStatementToFieldDefAction,
	{_State133, _WildcardMarker}:                  _ReduceImplicitToFieldDefAction,
	{_State136, LparenToken}:                      _ReduceNilToOptionalGenericParametersAction,
	{_State138, _WildcardMarker}:                  _ReduceParameterDefToParameterDefsAction,
	{_State139, RparenToken}:                      _ReduceParameterDefsToOptionalParameterDefsAction,
	{_State140, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State141, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State141, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State142, _WildcardMarker}:                  _ReduceImplicitToAnonymousStructExprAction,
	{_State143, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State143, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State143, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State143, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State143, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State144, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State144, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State144, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State144, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State144, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State146, RparenToken}:                      _ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction,
	{_State147, _WildcardMarker}:                  _ReduceFieldDefToExplicitFieldDefsAction,
	{_State149, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State150, RbracketToken}:                    _ReduceGenericArgumentsToOptionalGenericArgumentsAction,
	{_State152, _WildcardMarker}:                  _ReduceValueTypeToGenericArgumentsAction,
	{_State153, _WildcardMarker}:                  _ReduceAccessToAccessExprAction,
	{_State155, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State155, RparenToken}:                      _ReduceNilToOptionalArgumentsAction,
	{_State155, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State156, _WildcardMarker}:                  _ReduceOpToAddExprAction,
	{_State157, _WildcardMarker}:                  _ReduceOpToAndExprAction,
	{_State158, _WildcardMarker}:                  _ReduceOpToCmpExprAction,
	{_State160, _WildcardMarker}:                  _ReduceOpToMulExprAction,
	{_State161, _WildcardMarker}:                  _ReduceInfiniteToLoopExprAction,
	{_State162, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State163, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State163, SemicolonToken}:                   _ReduceNilToOptionalSequenceExprAction,
	{_State166, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State168, _WildcardMarker}:                  _ReduceOpToOrExprAction,
	{_State170, _WildcardMarker}:                  _ReduceUnconstrainedToGenericParameterDefAction,
	{_State171, _WildcardMarker}:                  _ReduceGenericParameterDefToGenericParameterDefsAction,
	{_State172, RbracketToken}:                    _ReduceGenericParameterDefsToOptionalGenericParameterDefsAction,
	{_State174, _EndMarker}:                       _ReduceAliasToTypeDefAction,
	{_State175, _EndMarker}:                       _ReduceDefinitionToTypeDefAction,
	{_State176, _EndMarker}:                       _ReduceReferenceToValueTypeAction,
	{_State177, _WildcardMarker}:                  _ReducePublicMethodsTraitToTraitableTypeAction,
	{_State180, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State181, _WildcardMarker}:                  _ReduceNamedToAtomTypeAction,
	{_State182, _WildcardMarker}:                  _ReduceExplicitToFieldDefAction,
	{_State184, _WildcardMarker}:                  _ReduceFieldDefToImplicitFieldDefsAction,
	{_State184, OrToken}:                          _ReduceFieldDefToEnumValueDefAction,
	{_State186, RparenToken}:                      _ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction,
	{_State188, _WildcardMarker}:                  _ReducePublicTraitToTraitableTypeAction,
	{_State190, _EndMarker}:                       _ReduceToTraitDefAction,
	{_State197, _WildcardMarker}:                  _ReduceArgToParameterDefAction,
	{_State198, IdentifierToken}:                  _ReduceReceiverToOptionalReceiverAction,
	{_State200, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State202, _WildcardMarker}:                  _ReduceNamedToArgumentAction,
	{_State203, _WildcardMarker}:                  _ReduceAddToArgumentsAction,
	{_State204, _WildcardMarker}:                  _ReduceExpressionToOptionalExpressionAction,
	{_State205, _WildcardMarker}:                  _ReduceAddToColonExpressionsAction,
	{_State206, _WildcardMarker}:                  _ReducePairToColonExpressionsAction,
	{_State209, _WildcardMarker}:                  _ReduceToExplicitStructDefAction,
	{_State211, _WildcardMarker}:                  _ReduceBindingToOptionalGenericBindingAction,
	{_State212, _WildcardMarker}:                  _ReduceIndexToAccessExprAction,
	{_State213, RparenToken}:                      _ReduceArgumentsToOptionalArgumentsAction,
	{_State215, _WildcardMarker}:                  _ReduceExplicitToAnonymousStructExprAction,
	{_State216, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State219, DoToken}:                          _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State221, _WildcardMarker}:                  _ReduceNoElseToIfExprAction,
	{_State222, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State223, _WildcardMarker}:                  _ReduceBreakToJumpTypeAction,
	{_State224, _WildcardMarker}:                  _ReduceContinueToJumpTypeAction,
	{_State225, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State226, _EndMarker}:                       _ReduceToBlockBodyAction,
	{_State227, _WildcardMarker}:                  _ReduceReturnToJumpTypeAction,
	{_State228, _WildcardMarker}:                  _ReduceAccessExprToPostfixUnaryExprAction,
	{_State228, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State229, _WildcardMarker}:                  _ReduceExpressionToExpressionsAction,
	{_State230, _WildcardMarker}:                  _ReduceExpressionOrImplicitStructToStatementBodyAction,
	{_State231, _WildcardMarker}:                  _ReduceJumpStatementToStatementBodyAction,
	{_State232, _WildcardMarker}:                  _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State233, _WildcardMarker}:                  _ReduceAddToStatementsAction,
	{_State235, _WildcardMarker}:                  _ReduceUnsafeStatementToStatementBodyAction,
	{_State237, _EndMarker}:                       _ReduceWithSpecToPackageDefAction,
	{_State238, _WildcardMarker}:                  _ReduceAddToPackageStatementsAction,
	{_State240, _WildcardMarker}:                  _ReduceToPackageStatementBodyAction,
	{_State241, _WildcardMarker}:                  _ReduceConstrainedToGenericParameterDefAction,
	{_State243, _WildcardMarker}:                  _ReduceGenericToOptionalGenericParametersAction,
	{_State247, _WildcardMarker}:                  _ReduceFieldDefToEnumValueDefAction,
	{_State249, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State251, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State253, _WildcardMarker}:                  _ReduceParameterDeclToParameterDeclsAction,
	{_State254, RparenToken}:                      _ReduceParameterDeclsToOptionalParameterDeclsAction,
	{_State255, _WildcardMarker}:                  _ReduceUnamedToParameterDeclAction,
	{_State259, _WildcardMarker}:                  _ReduceToImplicitEnumDefAction,
	{_State261, _WildcardMarker}:                  _ReduceToImplicitStructDefAction,
	{_State263, _WildcardMarker}:                  _ReduceUnionToTraitAlgebraTypeAction,
	{_State264, _WildcardMarker}:                  _ReduceIntersectToTraitAlgebraTypeAction,
	{_State265, _WildcardMarker}:                  _ReduceDifferenceToTraitAlgebraTypeAction,
	{_State266, _WildcardMarker}:                  _ReduceExplicitToTraitPropertiesAction,
	{_State267, _WildcardMarker}:                  _ReduceImplicitToTraitPropertiesAction,
	{_State268, _WildcardMarker}:                  _ReduceVarargToParameterDefAction,
	{_State269, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State271, _EndMarker}:                       _ReduceValueTypeToReturnTypeAction,
	{_State272, _WildcardMarker}:                  _ReduceAddToParameterDefsAction,
	{_State273, _WildcardMarker}:                  _ReduceExplicitToExplicitFieldDefsAction,
	{_State274, _WildcardMarker}:                  _ReduceImplicitToExplicitFieldDefsAction,
	{_State275, _WildcardMarker}:                  _ReduceAddToGenericArgumentsAction,
	{_State276, _WildcardMarker}:                  _ReduceToCallExprAction,
	{_State277, _EndMarker}:                       _ReduceDoWhileToLoopExprAction,
	{_State279, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State279, DoToken}:                          _ReduceNilToOptionalSequenceExprAction,
	{_State280, _EndMarker}:                       _ReduceWhileToLoopExprAction,
	{_State282, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State283, NewlinesToken}:                    _ReduceAsyncToStatementBodyAction,
	{_State283, SemicolonToken}:                   _ReduceAsyncToStatementBodyAction,
	{_State283, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State284, NewlinesToken}:                    _ReduceDeferToStatementBodyAction,
	{_State284, SemicolonToken}:                   _ReduceDeferToStatementBodyAction,
	{_State284, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State285, _WildcardMarker}:                  _ReduceAddAssignToBinaryOpAssignAction,
	{_State286, _WildcardMarker}:                  _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State287, _WildcardMarker}:                  _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State288, _WildcardMarker}:                  _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State289, _WildcardMarker}:                  _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State290, _WildcardMarker}:                  _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State291, _WildcardMarker}:                  _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State292, _WildcardMarker}:                  _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State293, _WildcardMarker}:                  _ReduceDivAssignToBinaryOpAssignAction,
	{_State294, _WildcardMarker}:                  _ReduceModAssignToBinaryOpAssignAction,
	{_State295, _WildcardMarker}:                  _ReduceMulAssignToBinaryOpAssignAction,
	{_State296, _WildcardMarker}:                  _ReduceSubAssignToBinaryOpAssignAction,
	{_State297, _WildcardMarker}:                  _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State298, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State299, _WildcardMarker}:                  _ReduceUnaryOpAssignStatementToStatementBodyAction,
	{_State300, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State301, _WildcardMarker}:                  _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State302, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State302, NewlinesToken}:                    _ReduceNilToOptionalExpressionsAction,
	{_State302, SemicolonToken}:                   _ReduceNilToOptionalExpressionsAction,
	{_State303, _WildcardMarker}:                  _ReduceImplicitToStatementAction,
	{_State304, _WildcardMarker}:                  _ReduceExplicitToStatementAction,
	{_State306, _WildcardMarker}:                  _ReduceImplicitToPackageStatementAction,
	{_State307, _WildcardMarker}:                  _ReduceExplicitToPackageStatementAction,
	{_State308, _WildcardMarker}:                  _ReduceAddToGenericParameterDefsAction,
	{_State309, _EndMarker}:                       _ReduceConstrainedDefToTypeDefAction,
	{_State312, _WildcardMarker}:                  _ReduceToExplicitEnumDefAction,
	{_State316, _WildcardMarker}:                  _ReduceUnnamedVarargToParameterDeclAction,
	{_State318, _WildcardMarker}:                  _ReduceArgToParameterDeclAction,
	{_State319, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State321, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State322, _WildcardMarker}:                  _ReduceDefaultToEnumValueDefAction,
	{_State323, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State324, _WildcardMarker}:                  _ReduceAddToImplicitFieldDefsAction,
	{_State327, _WildcardMarker}:                  _ReduceToAnonymousFuncExprAction,
	{_State328, _EndMarker}:                       _ReduceIteratorToLoopExprAction,
	{_State330, _EndMarker}:                       _ReduceIfElseToIfExprAction,
	{_State331, _EndMarker}:                       _ReduceMultiIfElseToIfExprAction,
	{_State332, _WildcardMarker}:                  _ReduceBinaryOpAssignStatementToStatementBodyAction,
	{_State333, _WildcardMarker}:                  _ReduceAddToExpressionsAction,
	{_State334, _WildcardMarker}:                  _ReduceExpressionsToOptionalExpressionsAction,
	{_State335, _WildcardMarker}:                  _ReduceToJumpStatementAction,
	{_State337, RparenToken}:                      _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State338, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State338, RparenToken}:                      _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State339, RparenToken}:                      _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State340, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State340, RparenToken}:                      _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State341, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State342, _WildcardMarker}:                  _ReduceVarargToParameterDeclAction,
	{_State343, _EndMarker}:                       _ReduceToFuncTypeAction,
	{_State344, _WildcardMarker}:                  _ReduceAddToParameterDeclsAction,
	{_State345, _WildcardMarker}:                  _ReduceToUnsafeStatementAction,
	{_State346, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State348, _EndMarker}:                       _ReduceToSwitchExprAction,
	{_State349, _WildcardMarker}:                  _ReduceToMethodSignatureAction,
	{_State351, _EndMarker}:                       _ReduceForToLoopExprAction,
	{_State352, _EndMarker}:                       _ReduceToNamedFuncDefAction,
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      expression -> State 11
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      expression -> State 64
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      call_expr -> State 41
      arguments -> State 62
      argument -> State 61
      colon_expressions -> State 63
      optional_expression -> State 65
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
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
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      IF -> State 95
      SWITCH -> State 97
      FOR -> State 94
      DO -> State 93
      LBRACE -> State 96
      if_expr -> State 99
      switch_expr -> State 101
      loop_expr -> State 100
      block_body -> State 98

  State 47:
    Kernel Items:
      sequence_expr: or_expr., *
      or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    Goto:
      OR -> State 102

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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 104
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
      LPAREN -> State 105

  State 55:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE IDENTIFIER.EQUAL value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 106
      EQUAL -> State 107
      optional_generic_parameters -> State 108

  State 56:
    Kernel Items:
      trait_def: TRAIT LPAREN.optional_trait_properties RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      IDENTIFIER -> State 113
      UNSAFE -> State 117
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 112
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      unsafe_statement -> State 132
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 133
      field_def -> State 121
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_property -> State 130
      trait_properties -> State 129
      optional_trait_properties -> State 126
      trait_def -> State 128
      func_type -> State 122
      method_signature -> State 125

  State 57:
    Kernel Items:
      optional_receiver: LPAREN.parameter_def RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 134
      parameter_def -> State 135

  State 58:
    Kernel Items:
      named_func_def: FUNC optional_receiver.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 136

  State 59:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 134
      parameter_def -> State 138
      parameter_defs -> State 139
      optional_parameter_defs -> State 137

  State 60:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expression
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      ASSIGN -> State 140

  State 61:
    Kernel Items:
      arguments: argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      arguments: arguments.COMMA argument
      anonymous_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 142
      COMMA -> State 141

  State 63:
    Kernel Items:
      argument: colon_expressions., *
      colon_expressions: colon_expressions.COLON optional_expression
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 143

  State 64:
    Kernel Items:
      argument: expression., *
      optional_expression: expression., COLON
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
      COLON -> State 144

  State 66:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN.optional_explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      IDENTIFIER -> State 113
      UNSAFE -> State 117
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      unsafe_statement -> State 132
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 133
      field_def -> State 147
      implicit_struct_def -> State 124
      explicit_field_defs -> State 146
      optional_explicit_field_defs -> State 148
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 67:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      optional_generic_arguments -> State 151
      generic_arguments -> State 150
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 152
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 68:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 153

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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      expression -> State 64
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      call_expr -> State 41
      argument -> State 154
      colon_expressions -> State 63
      optional_expression -> State 65
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
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
      LPAREN -> State 155

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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 156
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 157
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 158
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      expression -> State 64
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      call_expr -> State 41
      arguments -> State 159
      argument -> State 61
      colon_expressions -> State 63
      optional_expression -> State 65
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 160
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 93:
    Kernel Items:
      loop_expr: DO.block_body
      loop_expr: DO.block_body FOR sequence_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 96
      block_body -> State 161

  State 94:
    Kernel Items:
      loop_expr: FOR.sequence_expr DO block_body
      loop_expr: FOR.IN sequence_expr DO block_body
      loop_expr: FOR.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
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
      IN -> State 162
      STRUCT -> State 31
      FUNC -> State 21
      LABEL_DECL -> State 24
      LPAREN -> State 26
      SEMICOLON -> State 163
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      sequence_expr -> State 164
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 95:
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      sequence_expr -> State 165
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 96:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 166

  State 97:
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      sequence_expr -> State 167
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 98:
    Kernel Items:
      block_expr: optional_label_decl block_body., *
    Reduce:
      * -> [block_expr]
    Goto:
      (nil)

  State 99:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 100:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 101:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 102:
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 168
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 103:
    Kernel Items:
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 96
      block_body -> State 98

  State 104:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 105:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN.package_statements RPAREN
    Reduce:
      * -> [package_statements]
    Goto:
      package_statements -> State 169

  State 106:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.optional_generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 170
      generic_parameter_def -> State 171
      generic_parameter_defs -> State 172
      optional_generic_parameter_defs -> State 173

  State 107:
    Kernel Items:
      type_def: TYPE IDENTIFIER EQUAL.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 174
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 108:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 175
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 109:
    Kernel Items:
      value_type: BIT_AND.trait_algebra_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      LPAREN -> State 114
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 176
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128

  State 110:
    Kernel Items:
      traitable_type: BIT_NEG.atom_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      LPAREN -> State 114
      atom_type -> State 177
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128

  State 111:
    Kernel Items:
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 178

  State 112:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 179
      LPAREN -> State 180

  State 113:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      DOLLAR_LBRACKET -> State 67
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      optional_generic_binding -> State 181
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 182
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 114:
    Kernel Items:
      implicit_struct_def: LPAREN.optional_implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      IDENTIFIER -> State 113
      UNSAFE -> State 117
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      unsafe_statement -> State 132
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 133
      field_def -> State 184
      implicit_field_defs -> State 186
      optional_implicit_field_defs -> State 187
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      enum_value_def -> State 183
      implicit_enum_value_defs -> State 185
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 115:
    Kernel Items:
      value_type: QUESTION., $
    Reduce:
      $ -> [value_type]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      traitable_type: TILDE_TILDE.atom_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      LPAREN -> State 114
      atom_type -> State 188
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128

  State 117:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 189

  State 118:
    Kernel Items:
      traitable_type: atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 119:
    Kernel Items:
      atom_type: explicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      atom_type: explicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 121:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 122:
    Kernel Items:
      value_type: func_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      (nil)

  State 123:
    Kernel Items:
      atom_type: implicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      atom_type: implicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 126:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 190

  State 127:
    Kernel Items:
      trait_algebra_type: trait_algebra_type.MUL traitable_type
      trait_algebra_type: trait_algebra_type.ADD traitable_type
      trait_algebra_type: trait_algebra_type.SUB traitable_type
      value_type: trait_algebra_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      ADD -> State 191
      SUB -> State 193
      MUL -> State 192

  State 128:
    Kernel Items:
      atom_type: trait_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 129:
    Kernel Items:
      trait_properties: trait_properties.NEWLINES trait_property
      trait_properties: trait_properties.COMMA trait_property
      optional_trait_properties: trait_properties., RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      NEWLINES -> State 195
      COMMA -> State 194

  State 130:
    Kernel Items:
      trait_properties: trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 131:
    Kernel Items:
      trait_algebra_type: traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 132:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 133:
    Kernel Items:
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 134:
    Kernel Items:
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      DOTDOTDOT -> State 196
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 197
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 135:
    Kernel Items:
      optional_receiver: LPAREN parameter_def.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 198

  State 136:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 106
      optional_generic_parameters -> State 199

  State 137:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 200

  State 138:
    Kernel Items:
      parameter_defs: parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 139:
    Kernel Items:
      parameter_defs: parameter_defs.COMMA parameter_def
      optional_parameter_defs: parameter_defs., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      COMMA -> State 201

  State 140:
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      expression -> State 202
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 141:
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      expression -> State 64
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      call_expr -> State 41
      argument -> State 203
      colon_expressions -> State 63
      optional_expression -> State 65
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 142:
    Kernel Items:
      anonymous_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 143:
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      expression -> State 204
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      call_expr -> State 41
      optional_expression -> State 205
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 144:
    Kernel Items:
      colon_expressions: optional_expression COLON.optional_expression
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      expression -> State 204
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      call_expr -> State 41
      optional_expression -> State 206
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 145:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 180

  State 146:
    Kernel Items:
      explicit_field_defs: explicit_field_defs.NEWLINES field_def
      explicit_field_defs: explicit_field_defs.COMMA field_def
      optional_explicit_field_defs: explicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      NEWLINES -> State 208
      COMMA -> State 207

  State 147:
    Kernel Items:
      explicit_field_defs: field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 148:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 209

  State 149:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 67
      optional_generic_binding -> State 181

  State 150:
    Kernel Items:
      optional_generic_arguments: generic_arguments., RBRACKET
      generic_arguments: generic_arguments.COMMA value_type
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 210

  State 151:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 211

  State 152:
    Kernel Items:
      generic_arguments: value_type., *
    Reduce:
      * -> [generic_arguments]
    Goto:
      (nil)

  State 153:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 154:
    Kernel Items:
      access_expr: access_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 212

  State 155:
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      expression -> State 64
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      call_expr -> State 41
      optional_arguments -> State 214
      arguments -> State 213
      argument -> State 61
      colon_expressions -> State 63
      optional_expression -> State 65
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 156:
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

  State 157:
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

  State 158:
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

  State 159:
    Kernel Items:
      arguments: arguments.COMMA argument
      anonymous_struct_expr: explicit_struct_def LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 215
      COMMA -> State 141

  State 160:
    Kernel Items:
      mul_expr: mul_expr mul_op prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 161:
    Kernel Items:
      loop_expr: DO block_body., *
      loop_expr: DO block_body.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 216

  State 162:
    Kernel Items:
      loop_expr: FOR IN.sequence_expr DO block_body
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      sequence_expr -> State 217
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 163:
    Kernel Items:
      loop_expr: FOR SEMICOLON.optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      LBRACE -> [optional_label_decl]
      SEMICOLON -> [optional_sequence_expr]
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      optional_sequence_expr -> State 218
      sequence_expr -> State 219
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 164:
    Kernel Items:
      loop_expr: FOR sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 220

  State 165:
    Kernel Items:
      if_expr: IF sequence_expr.block_body
      if_expr: IF sequence_expr.block_body ELSE block_body
      if_expr: IF sequence_expr.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 96
      block_body -> State 221

  State 166:
    Kernel Items:
      block_body: LBRACE statements.RBRACE
      statements: statements.statement
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
      RETURN -> State 227
      BREAK -> State 223
      CONTINUE -> State 224
      UNSAFE -> State 117
      STRUCT -> State 31
      FUNC -> State 21
      ASYNC -> State 222
      DEFER -> State 225
      LABEL_DECL -> State 24
      RBRACE -> State 226
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      expression -> State 229
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      statement -> State 233
      statement_body -> State 234
      unsafe_statement -> State 235
      jump_statement -> State 231
      jump_type -> State 232
      expressions -> State 230
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 228
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 167:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.LBRACE CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 236

  State 168:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 77

  State 169:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements.RPAREN
      package_statements: package_statements.package_statement
    Reduce:
      (nil)
    Goto:
      UNSAFE -> State 117
      RPAREN -> State 237
      unsafe_statement -> State 240
      package_statement_body -> State 239
      package_statement -> State 238

  State 170:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 241
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 171:
    Kernel Items:
      generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 172:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs.COMMA generic_parameter_def
      optional_generic_parameter_defs: generic_parameter_defs., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      COMMA -> State 242

  State 173:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 243

  State 174:
    Kernel Items:
      type_def: TYPE IDENTIFIER EQUAL value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      (nil)

  State 175:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., $
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      $ -> [type_def]
    Goto:
      IMPLEMENTS -> State 244

  State 176:
    Kernel Items:
      trait_algebra_type: trait_algebra_type.MUL traitable_type
      trait_algebra_type: trait_algebra_type.ADD traitable_type
      trait_algebra_type: trait_algebra_type.SUB traitable_type
      value_type: BIT_AND trait_algebra_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      ADD -> State 191
      SUB -> State 193
      MUL -> State 192

  State 177:
    Kernel Items:
      traitable_type: BIT_NEG atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 178:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNSAFE -> State 117
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      unsafe_statement -> State 132
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 133
      field_def -> State 247
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      enum_value_def -> State 245
      implicit_enum_value_defs -> State 248
      implicit_enum_def -> State 123
      explicit_enum_value_defs -> State 246
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 179:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 249

  State 180:
    Kernel Items:
      func_type: FUNC LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 251
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      DOTDOTDOT -> State 250
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 255
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      parameter_decl -> State 253
      parameter_decls -> State 254
      optional_parameter_decls -> State 252
      func_type -> State 122

  State 181:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 182:
    Kernel Items:
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 183:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 256

  State 184:
    Kernel Items:
      implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 257

  State 185:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 259
      OR -> State 258

  State 186:
    Kernel Items:
      implicit_field_defs: implicit_field_defs.COMMA field_def
      optional_implicit_field_defs: implicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      COMMA -> State 260

  State 187:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 261

  State 188:
    Kernel Items:
      traitable_type: TILDE_TILDE atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 189:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 262

  State 190:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties RPAREN., $
    Reduce:
      $ -> [trait_def]
    Goto:
      (nil)

  State 191:
    Kernel Items:
      trait_algebra_type: trait_algebra_type ADD.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      LPAREN -> State 114
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      atom_type -> State 118
      traitable_type -> State 263
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128

  State 192:
    Kernel Items:
      trait_algebra_type: trait_algebra_type MUL.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      LPAREN -> State 114
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      atom_type -> State 118
      traitable_type -> State 264
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128

  State 193:
    Kernel Items:
      trait_algebra_type: trait_algebra_type SUB.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      LPAREN -> State 114
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      atom_type -> State 118
      traitable_type -> State 265
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128

  State 194:
    Kernel Items:
      trait_properties: trait_properties COMMA.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNSAFE -> State 117
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 112
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      unsafe_statement -> State 132
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 133
      field_def -> State 121
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_property -> State 266
      trait_def -> State 128
      func_type -> State 122
      method_signature -> State 125

  State 195:
    Kernel Items:
      trait_properties: trait_properties NEWLINES.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNSAFE -> State 117
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 112
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      unsafe_statement -> State 132
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 133
      field_def -> State 121
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_property -> State 267
      trait_def -> State 128
      func_type -> State 122
      method_signature -> State 125

  State 196:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 268
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 197:
    Kernel Items:
      parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      (nil)

  State 198:
    Kernel Items:
      optional_receiver: LPAREN parameter_def RPAREN., IDENTIFIER
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      (nil)

  State 199:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 269

  State 200:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 271
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      return_type -> State 270
      func_type -> State 122

  State 201:
    Kernel Items:
      parameter_defs: parameter_defs COMMA.parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 134
      parameter_def -> State 272

  State 202:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 203:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 204:
    Kernel Items:
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      (nil)

  State 205:
    Kernel Items:
      colon_expressions: colon_expressions COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 206:
    Kernel Items:
      colon_expressions: optional_expression COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 207:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNSAFE -> State 117
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      unsafe_statement -> State 132
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 133
      field_def -> State 273
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 208:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNSAFE -> State 117
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      unsafe_statement -> State 132
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 133
      field_def -> State 274
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 209:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_def]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 275
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 211:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 212:
    Kernel Items:
      access_expr: access_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      optional_arguments: arguments., RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 141

  State 214:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 276

  State 215:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_def LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 216:
    Kernel Items:
      loop_expr: DO block_body FOR.sequence_expr
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      sequence_expr -> State 277
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 217:
    Kernel Items:
      loop_expr: FOR IN sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 278

  State 218:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 279

  State 219:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 220:
    Kernel Items:
      loop_expr: FOR sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 96
      block_body -> State 280

  State 221:
    Kernel Items:
      if_expr: IF sequence_expr block_body., *
      if_expr: IF sequence_expr block_body.ELSE block_body
      if_expr: IF sequence_expr block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 281

  State 222:
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      block_expr -> State 40
      call_expr -> State 283
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 282
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 223:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 224:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 225:
    Kernel Items:
      statement_body: DEFER.call_expr
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      block_expr -> State 40
      call_expr -> State 284
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 282
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 226:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 227:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 228:
    Kernel Items:
      statement_body: access_expr.unary_op_assign
      statement_body: access_expr.binary_op_assign expression
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
      ADD_ASSIGN -> State 285
      SUB_ASSIGN -> State 296
      MUL_ASSIGN -> State 295
      DIV_ASSIGN -> State 293
      MOD_ASSIGN -> State 294
      ADD_ONE_ASSIGN -> State 286
      SUB_ONE_ASSIGN -> State 297
      BIT_NEG_ASSIGN -> State 289
      BIT_AND_ASSIGN -> State 287
      BIT_OR_ASSIGN -> State 290
      BIT_XOR_ASSIGN -> State 292
      BIT_LSHIFT_ASSIGN -> State 288
      BIT_RSHIFT_ASSIGN -> State 291
      unary_op_assign -> State 299
      binary_op_assign -> State 298
      optional_generic_binding -> State 71

  State 229:
    Kernel Items:
      expressions: expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 230:
    Kernel Items:
      statement_body: expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [statement_body]
    Goto:
      COMMA -> State 300

  State 231:
    Kernel Items:
      statement_body: jump_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 232:
    Kernel Items:
      jump_statement: jump_type.optional_jump_label optional_expressions
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 301
      optional_jump_label -> State 302

  State 233:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 234:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 303
      SEMICOLON -> State 304

  State 235:
    Kernel Items:
      statement_body: unsafe_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 236:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE.CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      CASE -> State 305

  State 237:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements RPAREN., $
    Reduce:
      $ -> [package_def]
    Goto:
      (nil)

  State 238:
    Kernel Items:
      package_statements: package_statements package_statement., *
    Reduce:
      * -> [package_statements]
    Goto:
      (nil)

  State 239:
    Kernel Items:
      package_statement: package_statement_body.NEWLINES
      package_statement: package_statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 306
      SEMICOLON -> State 307

  State 240:
    Kernel Items:
      package_statement_body: unsafe_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      (nil)

  State 242:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA.generic_parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 170
      generic_parameter_def -> State 308

  State 243:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 244:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 309
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 245:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 310
      OR -> State 311

  State 246:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 312

  State 247:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 257

  State 248:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 313
      OR -> State 314

  State 249:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 251
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      DOTDOTDOT -> State 250
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 255
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      parameter_decl -> State 253
      parameter_decls -> State 254
      optional_parameter_decls -> State 315
      func_type -> State 122

  State 250:
    Kernel Items:
      parameter_decl: DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 316
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 251:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      DOLLAR_LBRACKET -> State 67
      DOTDOTDOT -> State 317
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      optional_generic_binding -> State 181
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 318
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 252:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 319

  State 253:
    Kernel Items:
      parameter_decls: parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 254:
    Kernel Items:
      parameter_decls: parameter_decls.COMMA parameter_decl
      optional_parameter_decls: parameter_decls., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      COMMA -> State 320

  State 255:
    Kernel Items:
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 256:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNSAFE -> State 117
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      unsafe_statement -> State 132
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 133
      field_def -> State 247
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      enum_value_def -> State 321
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 257:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 322

  State 258:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNSAFE -> State 117
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      unsafe_statement -> State 132
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 133
      field_def -> State 247
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      enum_value_def -> State 323
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 259:
    Kernel Items:
      implicit_enum_def: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_def]
    Goto:
      (nil)

  State 260:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNSAFE -> State 117
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      unsafe_statement -> State 132
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 133
      field_def -> State 324
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 261:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_def]
    Goto:
      (nil)

  State 262:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 325

  State 263:
    Kernel Items:
      trait_algebra_type: trait_algebra_type ADD traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 264:
    Kernel Items:
      trait_algebra_type: trait_algebra_type MUL traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 265:
    Kernel Items:
      trait_algebra_type: trait_algebra_type SUB traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 266:
    Kernel Items:
      trait_properties: trait_properties COMMA trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 267:
    Kernel Items:
      trait_properties: trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 268:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      (nil)

  State 269:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 134
      parameter_def -> State 138
      parameter_defs -> State 139
      optional_parameter_defs -> State 326

  State 270:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 96
      block_body -> State 327

  State 271:
    Kernel Items:
      return_type: value_type., $
    Reduce:
      $ -> [return_type]
    Goto:
      (nil)

  State 272:
    Kernel Items:
      parameter_defs: parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 273:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 274:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 275:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
    Reduce:
      * -> [generic_arguments]
    Goto:
      (nil)

  State 276:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 277:
    Kernel Items:
      loop_expr: DO block_body FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      loop_expr: FOR IN sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 96
      block_body -> State 328

  State 279:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON.optional_sequence_expr DO block_body
    Reduce:
      DO -> [optional_sequence_expr]
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      optional_sequence_expr -> State 329
      sequence_expr -> State 219
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 280:
    Kernel Items:
      loop_expr: FOR sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 281:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE.block_body
      if_expr: IF sequence_expr block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 95
      LBRACE -> State 96
      if_expr -> State 331
      block_body -> State 330

  State 282:
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

  State 283:
    Kernel Items:
      statement_body: ASYNC call_expr., NEWLINES
      statement_body: ASYNC call_expr., SEMICOLON
      access_expr: call_expr., *
    Reduce:
      * -> [access_expr]
      NEWLINES -> [statement_body]
      SEMICOLON -> [statement_body]
    Goto:
      (nil)

  State 284:
    Kernel Items:
      statement_body: DEFER call_expr., NEWLINES
      statement_body: DEFER call_expr., SEMICOLON
      access_expr: call_expr., *
    Reduce:
      * -> [access_expr]
      NEWLINES -> [statement_body]
      SEMICOLON -> [statement_body]
    Goto:
      (nil)

  State 285:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 286:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 287:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 288:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 289:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 290:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 291:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 292:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 293:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 294:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 295:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 296:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 297:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 298:
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      expression -> State 332
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 299:
    Kernel Items:
      statement_body: access_expr unary_op_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 300:
    Kernel Items:
      expressions: expressions COMMA.expression
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      expression -> State 333
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 301:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 302:
    Kernel Items:
      jump_statement: jump_type optional_jump_label.optional_expressions
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_expressions]
      SEMICOLON -> [optional_expressions]
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
      LABEL_DECL -> State 24
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      expression -> State 229
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      expressions -> State 334
      optional_expressions -> State 335
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 34
      postfix_unary_expr -> State 48
      prefix_unary_op -> State 50
      prefix_unary_expr -> State 49
      mul_expr -> State 45
      add_expr -> State 35
      cmp_expr -> State 42
      and_expr -> State 36
      or_expr -> State 47
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 303:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 304:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 305:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE CASE.DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 336

  State 306:
    Kernel Items:
      package_statement: package_statement_body NEWLINES., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 307:
    Kernel Items:
      package_statement: package_statement_body SEMICOLON., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 308:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 309:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      (nil)

  State 310:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNSAFE -> State 117
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      unsafe_statement -> State 132
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 133
      field_def -> State 247
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      enum_value_def -> State 337
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 311:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNSAFE -> State 117
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      unsafe_statement -> State 132
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 133
      field_def -> State 247
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      enum_value_def -> State 338
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 312:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_def]
    Goto:
      (nil)

  State 313:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNSAFE -> State 117
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      unsafe_statement -> State 132
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 133
      field_def -> State 247
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      enum_value_def -> State 339
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 314:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNSAFE -> State 117
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      unsafe_statement -> State 132
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 133
      field_def -> State 247
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      enum_value_def -> State 340
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 315:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 341

  State 316:
    Kernel Items:
      parameter_decl: DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 317:
    Kernel Items:
      parameter_decl: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 342
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 122

  State 318:
    Kernel Items:
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 319:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 271
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      return_type -> State 343
      func_type -> State 122

  State 320:
    Kernel Items:
      parameter_decls: parameter_decls COMMA.parameter_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 251
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      DOTDOTDOT -> State 250
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 255
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      parameter_decl -> State 344
      func_type -> State 122

  State 321:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 322:
    Kernel Items:
      enum_value_def: field_def ASSIGN DEFAULT., *
    Reduce:
      * -> [enum_value_def]
    Goto:
      (nil)

  State 323:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 324:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [implicit_field_defs]
    Goto:
      (nil)

  State 325:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 345

  State 326:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 346

  State 327:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 328:
    Kernel Items:
      loop_expr: FOR IN sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 329:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 347

  State 330:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 331:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 333:
    Kernel Items:
      expressions: expressions COMMA expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 334:
    Kernel Items:
      expressions: expressions.COMMA expression
      optional_expressions: expressions., *
    Reduce:
      * -> [optional_expressions]
    Goto:
      COMMA -> State 300

  State 335:
    Kernel Items:
      jump_statement: jump_type optional_jump_label optional_expressions., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 336:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE CASE DEFAULT.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 348

  State 337:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 338:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
      explicit_enum_value_defs: enum_value_def OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 339:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 340:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 271
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      return_type -> State 349
      func_type -> State 122

  State 342:
    Kernel Items:
      parameter_decl: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 343:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN return_type., $
    Reduce:
      $ -> [func_type]
    Goto:
      (nil)

  State 344:
    Kernel Items:
      parameter_decls: parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 345:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 346:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 149
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 145
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 118
      traitable_type -> State 131
      trait_algebra_type -> State 127
      value_type -> State 271
      implicit_struct_def -> State 124
      explicit_struct_def -> State 120
      implicit_enum_def -> State 123
      explicit_enum_def -> State 119
      trait_def -> State 128
      return_type -> State 350
      func_type -> State 122

  State 347:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 96
      block_body -> State 351

  State 348:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE CASE DEFAULT RBRACE., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 349:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 350:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 96
      block_body -> State 352

  State 351:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 352:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

Number of states: 352
Number of shift actions: 2057
Number of reduce actions: 285
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
