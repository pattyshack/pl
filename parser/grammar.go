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
	LabelDeclToken       = SymbolId(286)
	JumpLabelToken       = SymbolId(287)
	LbraceToken          = SymbolId(288)
	RbraceToken          = SymbolId(289)
	LparenToken          = SymbolId(290)
	RparenToken          = SymbolId(291)
	LbracketToken        = SymbolId(292)
	RbracketToken        = SymbolId(293)
	DotToken             = SymbolId(294)
	CommaToken           = SymbolId(295)
	QuestionToken        = SymbolId(296)
	SemicolonToken       = SymbolId(297)
	ColonToken           = SymbolId(298)
	DollarLbracketToken  = SymbolId(299)
	DotdotdotToken       = SymbolId(300)
	TildeTildeToken      = SymbolId(301)
	AssignToken          = SymbolId(302)
	AddAssignToken       = SymbolId(303)
	SubAssignToken       = SymbolId(304)
	MulAssignToken       = SymbolId(305)
	DivAssignToken       = SymbolId(306)
	ModAssignToken       = SymbolId(307)
	AddOneAssignToken    = SymbolId(308)
	SubOneAssignToken    = SymbolId(309)
	BitNegAssignToken    = SymbolId(310)
	BitAndAssignToken    = SymbolId(311)
	BitOrAssignToken     = SymbolId(312)
	BitXorAssignToken    = SymbolId(313)
	BitLshiftAssignToken = SymbolId(314)
	BitRshiftAssignToken = SymbolId(315)
	NotToken             = SymbolId(316)
	AndToken             = SymbolId(317)
	OrToken              = SymbolId(318)
	AddToken             = SymbolId(319)
	SubToken             = SymbolId(320)
	MulToken             = SymbolId(321)
	DivToken             = SymbolId(322)
	ModToken             = SymbolId(323)
	BitNegToken          = SymbolId(324)
	BitAndToken          = SymbolId(325)
	BitXorToken          = SymbolId(326)
	BitOrToken           = SymbolId(327)
	BitLshiftToken       = SymbolId(328)
	BitRshiftToken       = SymbolId(329)
	EqualToken           = SymbolId(330)
	NotEqualToken        = SymbolId(331)
	LessToken            = SymbolId(332)
	LessOrEqualToken     = SymbolId(333)
	GreaterToken         = SymbolId(334)
	GreaterOrEqualToken  = SymbolId(335)
	LexErrorToken        = SymbolId(336)
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

	// 126:2: statement_body -> async: ...
	AsyncToStatementBody(Async_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 129:2: statement_body -> jump_statement: ...
	JumpStatementToStatementBody(JumpStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 141:2: statement_body -> unary_op_assign_statement: ...
	UnaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, UnaryOpAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 142:2: statement_body -> binary_op_assign_statement: ...
	BinaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 149:2: unary_op_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToUnaryOpAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 150:2: unary_op_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToUnaryOpAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 153:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 154:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 155:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 156:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 157:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 158:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 159:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 160:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 161:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 162:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 163:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 171:20: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ *GenericSymbol, Less_ *GenericSymbol, Identifier_ *GenericSymbol, Greater_ *GenericSymbol, StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 178:2: jump_statement -> ...
	ToJumpStatement(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 181:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 182:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 183:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 186:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 187:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 190:2: expressions -> expression: ...
	ExpressionToExpressions(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 191:2: expressions -> add: ...
	AddToExpressions(Expressions_ *GenericSymbol, Comma_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 194:2: optional_expressions -> expressions: ...
	ExpressionsToOptionalExpressions(Expressions_ *GenericSymbol) (*GenericSymbol, error)

	// 195:2: optional_expressions -> nil: ...
	NilToOptionalExpressions() (*GenericSymbol, error)

	// 201:13: call_expr -> ...
	ToCallExpr(AccessExpr_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalArguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 204:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ *GenericSymbol, OptionalGenericArguments_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 205:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (*GenericSymbol, error)

	// 208:2: optional_generic_arguments -> generic_arguments: ...
	GenericArgumentsToOptionalGenericArguments(GenericArguments_ *GenericSymbol) (*GenericSymbol, error)

	// 209:2: optional_generic_arguments -> nil: ...
	NilToOptionalGenericArguments() (*GenericSymbol, error)

	// 213:2: generic_arguments -> value_type: ...
	ValueTypeToGenericArguments(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 214:2: generic_arguments -> add: ...
	AddToGenericArguments(GenericArguments_ *GenericSymbol, Comma_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 217:2: optional_arguments -> arguments: ...
	ArgumentsToOptionalArguments(Arguments_ *GenericSymbol) (*GenericSymbol, error)

	// 218:2: optional_arguments -> nil: ...
	NilToOptionalArguments() (*GenericSymbol, error)

	// 221:2: arguments -> argument: ...
	ArgumentToArguments(Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 222:2: arguments -> add: ...
	AddToArguments(Arguments_ *GenericSymbol, Comma_ *GenericSymbol, Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 225:2: argument -> positional: ...
	PositionalToArgument(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 226:2: argument -> named: ...
	NamedToArgument(Identifier_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 227:2: argument -> colon_expressions: ...
	ColonExpressionsToArgument(ColonExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 230:2: colon_expressions -> pair: ...
	PairToColonExpressions(OptionalExpression_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_2 *GenericSymbol) (*GenericSymbol, error)

	// 231:2: colon_expressions -> add: ...
	AddToColonExpressions(ColonExpressions_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_ *GenericSymbol) (*GenericSymbol, error)

	// 234:2: optional_expression -> expression: ...
	ExpressionToOptionalExpression(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 235:2: optional_expression -> nil: ...
	NilToOptionalExpression() (*GenericSymbol, error)

	// 245:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 246:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 247:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 248:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 249:2: atom_expr -> anonymous_struct_expr: ...
	AnonymousStructExprToAtomExpr(AnonymousStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 250:2: atom_expr -> LEX_ERROR: ...
	LexErrorToAtomExpr(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 253:2: literal -> TRUE: ...
	TrueToLiteral(True_ *GenericSymbol) (*GenericSymbol, error)

	// 254:2: literal -> FALSE: ...
	FalseToLiteral(False_ *GenericSymbol) (*GenericSymbol, error)

	// 255:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 256:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 257:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 258:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 261:2: anonymous_struct_expr -> explicit: ...
	ExplicitToAnonymousStructExpr(ExplicitStructDef_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 262:2: anonymous_struct_expr -> implicit: ...
	ImplicitToAnonymousStructExpr(Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 265:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 266:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 267:2: access_expr -> call_expr: ...
	CallExprToAccessExpr(CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 268:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Argument_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 271:2: postfix_unary_expr -> access_expr: ...
	AccessExprToPostfixUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 272:2: postfix_unary_expr -> question: ...
	QuestionToPostfixUnaryExpr(AccessExpr_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 275:2: prefix_unary_op -> NOT: ...
	NotToPrefixUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 276:2: prefix_unary_op -> BIT_NEG: ...
	BitNegToPrefixUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 277:2: prefix_unary_op -> SUB: ...
	SubToPrefixUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 280:2: prefix_unary_op -> MUL: ...
	MulToPrefixUnaryOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 283:2: prefix_unary_op -> BIT_AND: ...
	BitAndToPrefixUnaryOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 286:2: prefix_unary_expr -> postfix_unary_expr: ...
	PostfixUnaryExprToPrefixUnaryExpr(PostfixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 287:2: prefix_unary_expr -> prefix_op: ...
	PrefixOpToPrefixUnaryExpr(PrefixUnaryOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 290:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 291:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 292:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 293:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 294:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 295:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 298:2: mul_expr -> prefix_unary_expr: ...
	PrefixUnaryExprToMulExpr(PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 299:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 302:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 303:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 304:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 305:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 308:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 309:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 312:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 313:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 314:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 315:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 316:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 317:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 320:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 321:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 324:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 325:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 328:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 329:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 336:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 337:2: atom_type -> explicit_struct_def: ...
	ExplicitStructDefToAtomType(ExplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 338:2: atom_type -> implicit_struct_def: ...
	ImplicitStructDefToAtomType(ImplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 339:2: atom_type -> explicit_enum_def: ...
	ExplicitEnumDefToAtomType(ExplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 340:2: atom_type -> implicit_enum_def: ...
	ImplicitEnumDefToAtomType(ImplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 341:2: atom_type -> trait_def: ...
	TraitDefToAtomType(TraitDef_ *GenericSymbol) (*GenericSymbol, error)

	// 344:2: traitable_type -> atom_type: ...
	AtomTypeToTraitableType(AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 345:2: traitable_type -> public_methods_trait: ...
	PublicMethodsTraitToTraitableType(BitNeg_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 346:2: traitable_type -> public_trait: ...
	PublicTraitToTraitableType(TildeTilde_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 350:2: trait_algebra_type -> traitable_type: ...
	TraitableTypeToTraitAlgebraType(TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 351:2: trait_algebra_type -> intersect: ...
	IntersectToTraitAlgebraType(TraitAlgebraType_ *GenericSymbol, Mul_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 352:2: trait_algebra_type -> union: ...
	UnionToTraitAlgebraType(TraitAlgebraType_ *GenericSymbol, Add_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 353:2: trait_algebra_type -> difference: ...
	DifferenceToTraitAlgebraType(TraitAlgebraType_ *GenericSymbol, Sub_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 356:2: value_type -> inferred: ...
	InferredToValueType(Question_ *GenericSymbol) (*GenericSymbol, error)

	// 357:2: value_type -> trait_algebra_type: ...
	TraitAlgebraTypeToValueType(TraitAlgebraType_ *GenericSymbol) (*GenericSymbol, error)

	// 358:2: value_type -> reference: ...
	ReferenceToValueType(BitAnd_ *GenericSymbol, TraitAlgebraType_ *GenericSymbol) (*GenericSymbol, error)

	// 359:2: value_type -> func_type: ...
	FuncTypeToValueType(FuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 362:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 363:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol, Implements_ *GenericSymbol, ValueType_2 *GenericSymbol) (*GenericSymbol, error)

	// 364:2: type_def -> alias: ...
	AliasToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, Equal_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 372:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 373:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 376:2: generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToGenericParameterDefs(GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 377:2: generic_parameter_defs -> add: ...
	AddToGenericParameterDefs(GenericParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 380:2: optional_generic_parameter_defs -> generic_parameter_defs: ...
	GenericParameterDefsToOptionalGenericParameterDefs(GenericParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 381:2: optional_generic_parameter_defs -> nil: ...
	NilToOptionalGenericParameterDefs() (*GenericSymbol, error)

	// 384:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ *GenericSymbol, OptionalGenericParameterDefs_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 385:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (*GenericSymbol, error)

	// 392:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 393:2: field_def -> implicit: ...
	ImplicitToFieldDef(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 396:2: implicit_field_defs -> field_def: ...
	FieldDefToImplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 397:2: implicit_field_defs -> add: ...
	AddToImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 400:2: optional_implicit_field_defs -> implicit_field_defs: ...
	ImplicitFieldDefsToOptionalImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 401:2: optional_implicit_field_defs -> nil: ...
	NilToOptionalImplicitFieldDefs() (*GenericSymbol, error)

	// 403:23: implicit_struct_def -> ...
	ToImplicitStructDef(Lparen_ *GenericSymbol, OptionalImplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 406:2: explicit_field_defs -> field_def: ...
	FieldDefToExplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 407:2: explicit_field_defs -> implicit: ...
	ImplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Newlines_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 408:2: explicit_field_defs -> explicit: ...
	ExplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 411:2: optional_explicit_field_defs -> explicit_field_defs: ...
	ExplicitFieldDefsToOptionalExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 412:2: optional_explicit_field_defs -> nil: ...
	NilToOptionalExplicitFieldDefs() (*GenericSymbol, error)

	// 414:23: explicit_struct_def -> ...
	ToExplicitStructDef(Struct_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalExplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 424:2: enum_value_def -> field_def: ...
	FieldDefToEnumValueDef(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 425:2: enum_value_def -> default: ...
	DefaultToEnumValueDef(FieldDef_ *GenericSymbol, Assign_ *GenericSymbol, Default_ *GenericSymbol) (*GenericSymbol, error)

	// 428:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 429:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 431:21: implicit_enum_def -> ...
	ToImplicitEnumDef(Lparen_ *GenericSymbol, ImplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 434:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 435:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 436:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 437:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 439:21: explicit_enum_def -> ...
	ToExplicitEnumDef(Enum_ *GenericSymbol, Lparen_ *GenericSymbol, ExplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 446:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 447:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ *GenericSymbol) (*GenericSymbol, error)

	// 450:2: trait_properties -> trait_property: ...
	TraitPropertyToTraitProperties(TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 451:2: trait_properties -> implicit: ...
	ImplicitToTraitProperties(TraitProperties_ *GenericSymbol, Newlines_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 452:2: trait_properties -> explicit: ...
	ExplicitToTraitProperties(TraitProperties_ *GenericSymbol, Comma_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 455:2: optional_trait_properties -> trait_properties: ...
	TraitPropertiesToOptionalTraitProperties(TraitProperties_ *GenericSymbol) (*GenericSymbol, error)

	// 456:2: optional_trait_properties -> nil: ...
	NilToOptionalTraitProperties() (*GenericSymbol, error)

	// 458:13: trait_def -> ...
	ToTraitDef(Trait_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalTraitProperties_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 466:2: return_type -> value_type: ...
	ValueTypeToReturnType(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 467:2: return_type -> nil: ...
	NilToReturnType() (*GenericSymbol, error)

	// 470:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 471:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 472:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 473:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 476:2: parameter_decls -> parameter_decl: ...
	ParameterDeclToParameterDecls(ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 477:2: parameter_decls -> add: ...
	AddToParameterDecls(ParameterDecls_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 480:2: optional_parameter_decls -> parameter_decls: ...
	ParameterDeclsToOptionalParameterDecls(ParameterDecls_ *GenericSymbol) (*GenericSymbol, error)

	// 481:2: optional_parameter_decls -> nil: ...
	NilToOptionalParameterDecls() (*GenericSymbol, error)

	// 483:13: func_type -> ...
	ToFuncType(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 491:20: method_signature -> ...
	ToMethodSignature(Func_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 494:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 495:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 498:2: parameter_defs -> parameter_def: ...
	ParameterDefToParameterDefs(ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 499:2: parameter_defs -> add: ...
	AddToParameterDefs(ParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 502:2: optional_parameter_defs -> parameter_defs: ...
	ParameterDefsToOptionalParameterDefs(ParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 503:2: optional_parameter_defs -> nil: ...
	NilToOptionalParameterDefs() (*GenericSymbol, error)

	// 506:2: optional_receiver -> receiver: ...
	ReceiverToOptionalReceiver(Lparen_ *GenericSymbol, ParameterDef_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 507:2: optional_receiver -> nil: ...
	NilToOptionalReceiver() (*GenericSymbol, error)

	// 510:2: named_func_def -> ...
	ToNamedFuncDef(Func_ *GenericSymbol, OptionalReceiver_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 513:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 520:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 521:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, PackageStatements_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 523:26: package_statement_body -> ...
	ToPackageStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 526:2: package_statement -> implicit: ...
	ImplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 527:2: package_statement -> explicit: ...
	ExplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 530:2: package_statements -> empty_list: ...
	EmptyListToPackageStatements() (*GenericSymbol, error)

	// 531:2: package_statements -> add: ...
	AddToPackageStatements(PackageStatements_ *GenericSymbol, PackageStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 535:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 536:2: lex_internal_tokens -> COMMENT: ...
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

	ExpressionType                   = SymbolId(337)
	OptionalLabelDeclType            = SymbolId(338)
	IfExprType                       = SymbolId(339)
	SwitchExprType                   = SymbolId(340)
	LoopExprType                     = SymbolId(341)
	OptionalSequenceExprType         = SymbolId(342)
	SequenceExprType                 = SymbolId(343)
	BlockExprType                    = SymbolId(344)
	BlockBodyType                    = SymbolId(345)
	StatementsType                   = SymbolId(346)
	StatementType                    = SymbolId(347)
	StatementBodyType                = SymbolId(348)
	UnaryOpAssignType                = SymbolId(349)
	BinaryOpAssignType               = SymbolId(350)
	UnsafeStatementType              = SymbolId(351)
	JumpStatementType                = SymbolId(352)
	JumpTypeType                     = SymbolId(353)
	OptionalJumpLabelType            = SymbolId(354)
	ExpressionsType                  = SymbolId(355)
	OptionalExpressionsType          = SymbolId(356)
	CallExprType                     = SymbolId(357)
	OptionalGenericBindingType       = SymbolId(358)
	OptionalGenericArgumentsType     = SymbolId(359)
	GenericArgumentsType             = SymbolId(360)
	OptionalArgumentsType            = SymbolId(361)
	ArgumentsType                    = SymbolId(362)
	ArgumentType                     = SymbolId(363)
	ColonExpressionsType             = SymbolId(364)
	OptionalExpressionType           = SymbolId(365)
	AtomExprType                     = SymbolId(366)
	LiteralType                      = SymbolId(367)
	AnonymousStructExprType          = SymbolId(368)
	AccessExprType                   = SymbolId(369)
	PostfixUnaryExprType             = SymbolId(370)
	PrefixUnaryOpType                = SymbolId(371)
	PrefixUnaryExprType              = SymbolId(372)
	MulOpType                        = SymbolId(373)
	MulExprType                      = SymbolId(374)
	AddOpType                        = SymbolId(375)
	AddExprType                      = SymbolId(376)
	CmpOpType                        = SymbolId(377)
	CmpExprType                      = SymbolId(378)
	AndExprType                      = SymbolId(379)
	OrExprType                       = SymbolId(380)
	AtomTypeType                     = SymbolId(381)
	TraitableTypeType                = SymbolId(382)
	TraitAlgebraTypeType             = SymbolId(383)
	ValueTypeType                    = SymbolId(384)
	TypeDefType                      = SymbolId(385)
	GenericParameterDefType          = SymbolId(386)
	GenericParameterDefsType         = SymbolId(387)
	OptionalGenericParameterDefsType = SymbolId(388)
	OptionalGenericParametersType    = SymbolId(389)
	FieldDefType                     = SymbolId(390)
	ImplicitFieldDefsType            = SymbolId(391)
	OptionalImplicitFieldDefsType    = SymbolId(392)
	ImplicitStructDefType            = SymbolId(393)
	ExplicitFieldDefsType            = SymbolId(394)
	OptionalExplicitFieldDefsType    = SymbolId(395)
	ExplicitStructDefType            = SymbolId(396)
	EnumValueDefType                 = SymbolId(397)
	ImplicitEnumValueDefsType        = SymbolId(398)
	ImplicitEnumDefType              = SymbolId(399)
	ExplicitEnumValueDefsType        = SymbolId(400)
	ExplicitEnumDefType              = SymbolId(401)
	TraitPropertyType                = SymbolId(402)
	TraitPropertiesType              = SymbolId(403)
	OptionalTraitPropertiesType      = SymbolId(404)
	TraitDefType                     = SymbolId(405)
	ReturnTypeType                   = SymbolId(406)
	ParameterDeclType                = SymbolId(407)
	ParameterDeclsType               = SymbolId(408)
	OptionalParameterDeclsType       = SymbolId(409)
	FuncTypeType                     = SymbolId(410)
	MethodSignatureType              = SymbolId(411)
	ParameterDefType                 = SymbolId(412)
	ParameterDefsType                = SymbolId(413)
	OptionalParameterDefsType        = SymbolId(414)
	OptionalReceiverType             = SymbolId(415)
	NamedFuncDefType                 = SymbolId(416)
	AnonymousFuncExprType            = SymbolId(417)
	PackageDefType                   = SymbolId(418)
	PackageStatementBodyType         = SymbolId(419)
	PackageStatementType             = SymbolId(420)
	PackageStatementsType            = SymbolId(421)
	LexInternalTokensType            = SymbolId(422)
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
	_ReduceJumpStatementToStatementBody                       = _ReduceType(28)
	_ReduceUnaryOpAssignStatementToStatementBody              = _ReduceType(29)
	_ReduceBinaryOpAssignStatementToStatementBody             = _ReduceType(30)
	_ReduceAddOneAssignToUnaryOpAssign                        = _ReduceType(31)
	_ReduceSubOneAssignToUnaryOpAssign                        = _ReduceType(32)
	_ReduceAddAssignToBinaryOpAssign                          = _ReduceType(33)
	_ReduceSubAssignToBinaryOpAssign                          = _ReduceType(34)
	_ReduceMulAssignToBinaryOpAssign                          = _ReduceType(35)
	_ReduceDivAssignToBinaryOpAssign                          = _ReduceType(36)
	_ReduceModAssignToBinaryOpAssign                          = _ReduceType(37)
	_ReduceBitNegAssignToBinaryOpAssign                       = _ReduceType(38)
	_ReduceBitAndAssignToBinaryOpAssign                       = _ReduceType(39)
	_ReduceBitOrAssignToBinaryOpAssign                        = _ReduceType(40)
	_ReduceBitXorAssignToBinaryOpAssign                       = _ReduceType(41)
	_ReduceBitLshiftAssignToBinaryOpAssign                    = _ReduceType(42)
	_ReduceBitRshiftAssignToBinaryOpAssign                    = _ReduceType(43)
	_ReduceToUnsafeStatement                                  = _ReduceType(44)
	_ReduceToJumpStatement                                    = _ReduceType(45)
	_ReduceReturnToJumpType                                   = _ReduceType(46)
	_ReduceBreakToJumpType                                    = _ReduceType(47)
	_ReduceContinueToJumpType                                 = _ReduceType(48)
	_ReduceJumpLabelToOptionalJumpLabel                       = _ReduceType(49)
	_ReduceUnlabelledToOptionalJumpLabel                      = _ReduceType(50)
	_ReduceExpressionToExpressions                            = _ReduceType(51)
	_ReduceAddToExpressions                                   = _ReduceType(52)
	_ReduceExpressionsToOptionalExpressions                   = _ReduceType(53)
	_ReduceNilToOptionalExpressions                           = _ReduceType(54)
	_ReduceToCallExpr                                         = _ReduceType(55)
	_ReduceBindingToOptionalGenericBinding                    = _ReduceType(56)
	_ReduceNilToOptionalGenericBinding                        = _ReduceType(57)
	_ReduceGenericArgumentsToOptionalGenericArguments         = _ReduceType(58)
	_ReduceNilToOptionalGenericArguments                      = _ReduceType(59)
	_ReduceValueTypeToGenericArguments                        = _ReduceType(60)
	_ReduceAddToGenericArguments                              = _ReduceType(61)
	_ReduceArgumentsToOptionalArguments                       = _ReduceType(62)
	_ReduceNilToOptionalArguments                             = _ReduceType(63)
	_ReduceArgumentToArguments                                = _ReduceType(64)
	_ReduceAddToArguments                                     = _ReduceType(65)
	_ReducePositionalToArgument                               = _ReduceType(66)
	_ReduceNamedToArgument                                    = _ReduceType(67)
	_ReduceColonExpressionsToArgument                         = _ReduceType(68)
	_ReducePairToColonExpressions                             = _ReduceType(69)
	_ReduceAddToColonExpressions                              = _ReduceType(70)
	_ReduceExpressionToOptionalExpression                     = _ReduceType(71)
	_ReduceNilToOptionalExpression                            = _ReduceType(72)
	_ReduceLiteralToAtomExpr                                  = _ReduceType(73)
	_ReduceIdentifierToAtomExpr                               = _ReduceType(74)
	_ReduceBlockExprToAtomExpr                                = _ReduceType(75)
	_ReduceAnonymousFuncExprToAtomExpr                        = _ReduceType(76)
	_ReduceAnonymousStructExprToAtomExpr                      = _ReduceType(77)
	_ReduceLexErrorToAtomExpr                                 = _ReduceType(78)
	_ReduceTrueToLiteral                                      = _ReduceType(79)
	_ReduceFalseToLiteral                                     = _ReduceType(80)
	_ReduceIntegerLiteralToLiteral                            = _ReduceType(81)
	_ReduceFloatLiteralToLiteral                              = _ReduceType(82)
	_ReduceRuneLiteralToLiteral                               = _ReduceType(83)
	_ReduceStringLiteralToLiteral                             = _ReduceType(84)
	_ReduceExplicitToAnonymousStructExpr                      = _ReduceType(85)
	_ReduceImplicitToAnonymousStructExpr                      = _ReduceType(86)
	_ReduceAtomExprToAccessExpr                               = _ReduceType(87)
	_ReduceAccessToAccessExpr                                 = _ReduceType(88)
	_ReduceCallExprToAccessExpr                               = _ReduceType(89)
	_ReduceIndexToAccessExpr                                  = _ReduceType(90)
	_ReduceAccessExprToPostfixUnaryExpr                       = _ReduceType(91)
	_ReduceQuestionToPostfixUnaryExpr                         = _ReduceType(92)
	_ReduceNotToPrefixUnaryOp                                 = _ReduceType(93)
	_ReduceBitNegToPrefixUnaryOp                              = _ReduceType(94)
	_ReduceSubToPrefixUnaryOp                                 = _ReduceType(95)
	_ReduceMulToPrefixUnaryOp                                 = _ReduceType(96)
	_ReduceBitAndToPrefixUnaryOp                              = _ReduceType(97)
	_ReducePostfixUnaryExprToPrefixUnaryExpr                  = _ReduceType(98)
	_ReducePrefixOpToPrefixUnaryExpr                          = _ReduceType(99)
	_ReduceMulToMulOp                                         = _ReduceType(100)
	_ReduceDivToMulOp                                         = _ReduceType(101)
	_ReduceModToMulOp                                         = _ReduceType(102)
	_ReduceBitAndToMulOp                                      = _ReduceType(103)
	_ReduceBitLshiftToMulOp                                   = _ReduceType(104)
	_ReduceBitRshiftToMulOp                                   = _ReduceType(105)
	_ReducePrefixUnaryExprToMulExpr                           = _ReduceType(106)
	_ReduceOpToMulExpr                                        = _ReduceType(107)
	_ReduceAddToAddOp                                         = _ReduceType(108)
	_ReduceSubToAddOp                                         = _ReduceType(109)
	_ReduceBitOrToAddOp                                       = _ReduceType(110)
	_ReduceBitXorToAddOp                                      = _ReduceType(111)
	_ReduceMulExprToAddExpr                                   = _ReduceType(112)
	_ReduceOpToAddExpr                                        = _ReduceType(113)
	_ReduceEqualToCmpOp                                       = _ReduceType(114)
	_ReduceNotEqualToCmpOp                                    = _ReduceType(115)
	_ReduceLessToCmpOp                                        = _ReduceType(116)
	_ReduceLessOrEqualToCmpOp                                 = _ReduceType(117)
	_ReduceGreaterToCmpOp                                     = _ReduceType(118)
	_ReduceGreaterOrEqualToCmpOp                              = _ReduceType(119)
	_ReduceAddExprToCmpExpr                                   = _ReduceType(120)
	_ReduceOpToCmpExpr                                        = _ReduceType(121)
	_ReduceCmpExprToAndExpr                                   = _ReduceType(122)
	_ReduceOpToAndExpr                                        = _ReduceType(123)
	_ReduceAndExprToOrExpr                                    = _ReduceType(124)
	_ReduceOpToOrExpr                                         = _ReduceType(125)
	_ReduceNamedToAtomType                                    = _ReduceType(126)
	_ReduceExplicitStructDefToAtomType                        = _ReduceType(127)
	_ReduceImplicitStructDefToAtomType                        = _ReduceType(128)
	_ReduceExplicitEnumDefToAtomType                          = _ReduceType(129)
	_ReduceImplicitEnumDefToAtomType                          = _ReduceType(130)
	_ReduceTraitDefToAtomType                                 = _ReduceType(131)
	_ReduceAtomTypeToTraitableType                            = _ReduceType(132)
	_ReducePublicMethodsTraitToTraitableType                  = _ReduceType(133)
	_ReducePublicTraitToTraitableType                         = _ReduceType(134)
	_ReduceTraitableTypeToTraitAlgebraType                    = _ReduceType(135)
	_ReduceIntersectToTraitAlgebraType                        = _ReduceType(136)
	_ReduceUnionToTraitAlgebraType                            = _ReduceType(137)
	_ReduceDifferenceToTraitAlgebraType                       = _ReduceType(138)
	_ReduceInferredToValueType                                = _ReduceType(139)
	_ReduceTraitAlgebraTypeToValueType                        = _ReduceType(140)
	_ReduceReferenceToValueType                               = _ReduceType(141)
	_ReduceFuncTypeToValueType                                = _ReduceType(142)
	_ReduceDefinitionToTypeDef                                = _ReduceType(143)
	_ReduceConstrainedDefToTypeDef                            = _ReduceType(144)
	_ReduceAliasToTypeDef                                     = _ReduceType(145)
	_ReduceUnconstrainedToGenericParameterDef                 = _ReduceType(146)
	_ReduceConstrainedToGenericParameterDef                   = _ReduceType(147)
	_ReduceGenericParameterDefToGenericParameterDefs          = _ReduceType(148)
	_ReduceAddToGenericParameterDefs                          = _ReduceType(149)
	_ReduceGenericParameterDefsToOptionalGenericParameterDefs = _ReduceType(150)
	_ReduceNilToOptionalGenericParameterDefs                  = _ReduceType(151)
	_ReduceGenericToOptionalGenericParameters                 = _ReduceType(152)
	_ReduceNilToOptionalGenericParameters                     = _ReduceType(153)
	_ReduceExplicitToFieldDef                                 = _ReduceType(154)
	_ReduceImplicitToFieldDef                                 = _ReduceType(155)
	_ReduceFieldDefToImplicitFieldDefs                        = _ReduceType(156)
	_ReduceAddToImplicitFieldDefs                             = _ReduceType(157)
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefs       = _ReduceType(158)
	_ReduceNilToOptionalImplicitFieldDefs                     = _ReduceType(159)
	_ReduceToImplicitStructDef                                = _ReduceType(160)
	_ReduceFieldDefToExplicitFieldDefs                        = _ReduceType(161)
	_ReduceImplicitToExplicitFieldDefs                        = _ReduceType(162)
	_ReduceExplicitToExplicitFieldDefs                        = _ReduceType(163)
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefs       = _ReduceType(164)
	_ReduceNilToOptionalExplicitFieldDefs                     = _ReduceType(165)
	_ReduceToExplicitStructDef                                = _ReduceType(166)
	_ReduceFieldDefToEnumValueDef                             = _ReduceType(167)
	_ReduceDefaultToEnumValueDef                              = _ReduceType(168)
	_ReducePairToImplicitEnumValueDefs                        = _ReduceType(169)
	_ReduceAddToImplicitEnumValueDefs                         = _ReduceType(170)
	_ReduceToImplicitEnumDef                                  = _ReduceType(171)
	_ReduceExplicitPairToExplicitEnumValueDefs                = _ReduceType(172)
	_ReduceImplicitPairToExplicitEnumValueDefs                = _ReduceType(173)
	_ReduceExplicitAddToExplicitEnumValueDefs                 = _ReduceType(174)
	_ReduceImplicitAddToExplicitEnumValueDefs                 = _ReduceType(175)
	_ReduceToExplicitEnumDef                                  = _ReduceType(176)
	_ReduceFieldDefToTraitProperty                            = _ReduceType(177)
	_ReduceMethodSignatureToTraitProperty                     = _ReduceType(178)
	_ReduceTraitPropertyToTraitProperties                     = _ReduceType(179)
	_ReduceImplicitToTraitProperties                          = _ReduceType(180)
	_ReduceExplicitToTraitProperties                          = _ReduceType(181)
	_ReduceTraitPropertiesToOptionalTraitProperties           = _ReduceType(182)
	_ReduceNilToOptionalTraitProperties                       = _ReduceType(183)
	_ReduceToTraitDef                                         = _ReduceType(184)
	_ReduceValueTypeToReturnType                              = _ReduceType(185)
	_ReduceNilToReturnType                                    = _ReduceType(186)
	_ReduceArgToParameterDecl                                 = _ReduceType(187)
	_ReduceVarargToParameterDecl                              = _ReduceType(188)
	_ReduceUnamedToParameterDecl                              = _ReduceType(189)
	_ReduceUnnamedVarargToParameterDecl                       = _ReduceType(190)
	_ReduceParameterDeclToParameterDecls                      = _ReduceType(191)
	_ReduceAddToParameterDecls                                = _ReduceType(192)
	_ReduceParameterDeclsToOptionalParameterDecls             = _ReduceType(193)
	_ReduceNilToOptionalParameterDecls                        = _ReduceType(194)
	_ReduceToFuncType                                         = _ReduceType(195)
	_ReduceToMethodSignature                                  = _ReduceType(196)
	_ReduceArgToParameterDef                                  = _ReduceType(197)
	_ReduceVarargToParameterDef                               = _ReduceType(198)
	_ReduceParameterDefToParameterDefs                        = _ReduceType(199)
	_ReduceAddToParameterDefs                                 = _ReduceType(200)
	_ReduceParameterDefsToOptionalParameterDefs               = _ReduceType(201)
	_ReduceNilToOptionalParameterDefs                         = _ReduceType(202)
	_ReduceReceiverToOptionalReceiver                         = _ReduceType(203)
	_ReduceNilToOptionalReceiver                              = _ReduceType(204)
	_ReduceToNamedFuncDef                                     = _ReduceType(205)
	_ReduceToAnonymousFuncExpr                                = _ReduceType(206)
	_ReduceNoSpecToPackageDef                                 = _ReduceType(207)
	_ReduceWithSpecToPackageDef                               = _ReduceType(208)
	_ReduceToPackageStatementBody                             = _ReduceType(209)
	_ReduceImplicitToPackageStatement                         = _ReduceType(210)
	_ReduceExplicitToPackageStatement                         = _ReduceType(211)
	_ReduceEmptyListToPackageStatements                       = _ReduceType(212)
	_ReduceAddToPackageStatements                             = _ReduceType(213)
	_ReduceSpacesToLexInternalTokens                          = _ReduceType(214)
	_ReduceCommentToLexInternalTokens                         = _ReduceType(215)
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
	case _EndMarker, SpacesToken, NewlinesToken, CommentToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, PackageToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, DollarLbracketToken, DotdotdotToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, NotToken, AndToken, OrToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, LexErrorToken:
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
	{_State56, StructToken}:                       _GotoState31Action,
	{_State56, EnumToken}:                         _GotoState111Action,
	{_State56, TraitToken}:                        _GotoState15Action,
	{_State56, FuncToken}:                         _GotoState112Action,
	{_State56, LparenToken}:                       _GotoState114Action,
	{_State56, QuestionToken}:                     _GotoState115Action,
	{_State56, TildeTildeToken}:                   _GotoState116Action,
	{_State56, BitNegToken}:                       _GotoState110Action,
	{_State56, BitAndToken}:                       _GotoState109Action,
	{_State56, AtomTypeType}:                      _GotoState117Action,
	{_State56, TraitableTypeType}:                 _GotoState130Action,
	{_State56, TraitAlgebraTypeType}:              _GotoState126Action,
	{_State56, ValueTypeType}:                     _GotoState131Action,
	{_State56, FieldDefType}:                      _GotoState120Action,
	{_State56, ImplicitStructDefType}:             _GotoState123Action,
	{_State56, ExplicitStructDefType}:             _GotoState119Action,
	{_State56, ImplicitEnumDefType}:               _GotoState122Action,
	{_State56, ExplicitEnumDefType}:               _GotoState118Action,
	{_State56, TraitPropertyType}:                 _GotoState129Action,
	{_State56, TraitPropertiesType}:               _GotoState128Action,
	{_State56, OptionalTraitPropertiesType}:       _GotoState125Action,
	{_State56, TraitDefType}:                      _GotoState127Action,
	{_State56, FuncTypeType}:                      _GotoState121Action,
	{_State56, MethodSignatureType}:               _GotoState124Action,
	{_State57, IdentifierToken}:                   _GotoState132Action,
	{_State57, ParameterDefType}:                  _GotoState133Action,
	{_State58, IdentifierToken}:                   _GotoState134Action,
	{_State59, IdentifierToken}:                   _GotoState132Action,
	{_State59, ParameterDefType}:                  _GotoState136Action,
	{_State59, ParameterDefsType}:                 _GotoState137Action,
	{_State59, OptionalParameterDefsType}:         _GotoState135Action,
	{_State60, AssignToken}:                       _GotoState138Action,
	{_State62, RparenToken}:                       _GotoState140Action,
	{_State62, CommaToken}:                        _GotoState139Action,
	{_State63, ColonToken}:                        _GotoState141Action,
	{_State65, ColonToken}:                        _GotoState142Action,
	{_State66, IdentifierToken}:                   _GotoState113Action,
	{_State66, StructToken}:                       _GotoState31Action,
	{_State66, EnumToken}:                         _GotoState111Action,
	{_State66, TraitToken}:                        _GotoState15Action,
	{_State66, FuncToken}:                         _GotoState143Action,
	{_State66, LparenToken}:                       _GotoState114Action,
	{_State66, QuestionToken}:                     _GotoState115Action,
	{_State66, TildeTildeToken}:                   _GotoState116Action,
	{_State66, BitNegToken}:                       _GotoState110Action,
	{_State66, BitAndToken}:                       _GotoState109Action,
	{_State66, AtomTypeType}:                      _GotoState117Action,
	{_State66, TraitableTypeType}:                 _GotoState130Action,
	{_State66, TraitAlgebraTypeType}:              _GotoState126Action,
	{_State66, ValueTypeType}:                     _GotoState131Action,
	{_State66, FieldDefType}:                      _GotoState145Action,
	{_State66, ImplicitStructDefType}:             _GotoState123Action,
	{_State66, ExplicitFieldDefsType}:             _GotoState144Action,
	{_State66, OptionalExplicitFieldDefsType}:     _GotoState146Action,
	{_State66, ExplicitStructDefType}:             _GotoState119Action,
	{_State66, ImplicitEnumDefType}:               _GotoState122Action,
	{_State66, ExplicitEnumDefType}:               _GotoState118Action,
	{_State66, TraitDefType}:                      _GotoState127Action,
	{_State66, FuncTypeType}:                      _GotoState121Action,
	{_State67, IdentifierToken}:                   _GotoState147Action,
	{_State67, StructToken}:                       _GotoState31Action,
	{_State67, EnumToken}:                         _GotoState111Action,
	{_State67, TraitToken}:                        _GotoState15Action,
	{_State67, FuncToken}:                         _GotoState143Action,
	{_State67, LparenToken}:                       _GotoState114Action,
	{_State67, QuestionToken}:                     _GotoState115Action,
	{_State67, TildeTildeToken}:                   _GotoState116Action,
	{_State67, BitNegToken}:                       _GotoState110Action,
	{_State67, BitAndToken}:                       _GotoState109Action,
	{_State67, OptionalGenericArgumentsType}:      _GotoState149Action,
	{_State67, GenericArgumentsType}:              _GotoState148Action,
	{_State67, AtomTypeType}:                      _GotoState117Action,
	{_State67, TraitableTypeType}:                 _GotoState130Action,
	{_State67, TraitAlgebraTypeType}:              _GotoState126Action,
	{_State67, ValueTypeType}:                     _GotoState150Action,
	{_State67, ImplicitStructDefType}:             _GotoState123Action,
	{_State67, ExplicitStructDefType}:             _GotoState119Action,
	{_State67, ImplicitEnumDefType}:               _GotoState122Action,
	{_State67, ExplicitEnumDefType}:               _GotoState118Action,
	{_State67, TraitDefType}:                      _GotoState127Action,
	{_State67, FuncTypeType}:                      _GotoState121Action,
	{_State68, IdentifierToken}:                   _GotoState151Action,
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
	{_State69, ArgumentType}:                      _GotoState152Action,
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
	{_State71, LparenToken}:                       _GotoState153Action,
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
	{_State76, MulExprType}:                       _GotoState154Action,
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
	{_State77, CmpExprType}:                       _GotoState155Action,
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
	{_State84, AddExprType}:                       _GotoState156Action,
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
	{_State85, ArgumentsType}:                     _GotoState157Action,
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
	{_State92, PrefixUnaryExprType}:               _GotoState158Action,
	{_State92, ExplicitStructDefType}:             _GotoState43Action,
	{_State92, AnonymousFuncExprType}:             _GotoState37Action,
	{_State93, LbraceToken}:                       _GotoState96Action,
	{_State93, BlockBodyType}:                     _GotoState159Action,
	{_State94, IntegerLiteralToken}:               _GotoState23Action,
	{_State94, FloatLiteralToken}:                 _GotoState20Action,
	{_State94, RuneLiteralToken}:                  _GotoState29Action,
	{_State94, StringLiteralToken}:                _GotoState30Action,
	{_State94, IdentifierToken}:                   _GotoState22Action,
	{_State94, TrueToken}:                         _GotoState33Action,
	{_State94, FalseToken}:                        _GotoState19Action,
	{_State94, InToken}:                           _GotoState160Action,
	{_State94, StructToken}:                       _GotoState31Action,
	{_State94, FuncToken}:                         _GotoState21Action,
	{_State94, LabelDeclToken}:                    _GotoState24Action,
	{_State94, LparenToken}:                       _GotoState26Action,
	{_State94, SemicolonToken}:                    _GotoState161Action,
	{_State94, NotToken}:                          _GotoState28Action,
	{_State94, SubToken}:                          _GotoState32Action,
	{_State94, MulToken}:                          _GotoState27Action,
	{_State94, BitNegToken}:                       _GotoState18Action,
	{_State94, BitAndToken}:                       _GotoState17Action,
	{_State94, LexErrorToken}:                     _GotoState25Action,
	{_State94, OptionalLabelDeclType}:             _GotoState103Action,
	{_State94, SequenceExprType}:                  _GotoState162Action,
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
	{_State95, SequenceExprType}:                  _GotoState163Action,
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
	{_State96, StatementsType}:                    _GotoState164Action,
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
	{_State97, SequenceExprType}:                  _GotoState165Action,
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
	{_State102, AndExprType}:                      _GotoState166Action,
	{_State102, ExplicitStructDefType}:            _GotoState43Action,
	{_State102, AnonymousFuncExprType}:            _GotoState37Action,
	{_State103, LbraceToken}:                      _GotoState96Action,
	{_State103, BlockBodyType}:                    _GotoState98Action,
	{_State105, PackageStatementsType}:            _GotoState167Action,
	{_State106, IdentifierToken}:                  _GotoState168Action,
	{_State106, GenericParameterDefType}:          _GotoState169Action,
	{_State106, GenericParameterDefsType}:         _GotoState170Action,
	{_State106, OptionalGenericParameterDefsType}: _GotoState171Action,
	{_State107, IdentifierToken}:                  _GotoState147Action,
	{_State107, StructToken}:                      _GotoState31Action,
	{_State107, EnumToken}:                        _GotoState111Action,
	{_State107, TraitToken}:                       _GotoState15Action,
	{_State107, FuncToken}:                        _GotoState143Action,
	{_State107, LparenToken}:                      _GotoState114Action,
	{_State107, QuestionToken}:                    _GotoState115Action,
	{_State107, TildeTildeToken}:                  _GotoState116Action,
	{_State107, BitNegToken}:                      _GotoState110Action,
	{_State107, BitAndToken}:                      _GotoState109Action,
	{_State107, AtomTypeType}:                     _GotoState117Action,
	{_State107, TraitableTypeType}:                _GotoState130Action,
	{_State107, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State107, ValueTypeType}:                    _GotoState172Action,
	{_State107, ImplicitStructDefType}:            _GotoState123Action,
	{_State107, ExplicitStructDefType}:            _GotoState119Action,
	{_State107, ImplicitEnumDefType}:              _GotoState122Action,
	{_State107, ExplicitEnumDefType}:              _GotoState118Action,
	{_State107, TraitDefType}:                     _GotoState127Action,
	{_State107, FuncTypeType}:                     _GotoState121Action,
	{_State108, IdentifierToken}:                  _GotoState147Action,
	{_State108, StructToken}:                      _GotoState31Action,
	{_State108, EnumToken}:                        _GotoState111Action,
	{_State108, TraitToken}:                       _GotoState15Action,
	{_State108, FuncToken}:                        _GotoState143Action,
	{_State108, LparenToken}:                      _GotoState114Action,
	{_State108, QuestionToken}:                    _GotoState115Action,
	{_State108, TildeTildeToken}:                  _GotoState116Action,
	{_State108, BitNegToken}:                      _GotoState110Action,
	{_State108, BitAndToken}:                      _GotoState109Action,
	{_State108, AtomTypeType}:                     _GotoState117Action,
	{_State108, TraitableTypeType}:                _GotoState130Action,
	{_State108, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State108, ValueTypeType}:                    _GotoState173Action,
	{_State108, ImplicitStructDefType}:            _GotoState123Action,
	{_State108, ExplicitStructDefType}:            _GotoState119Action,
	{_State108, ImplicitEnumDefType}:              _GotoState122Action,
	{_State108, ExplicitEnumDefType}:              _GotoState118Action,
	{_State108, TraitDefType}:                     _GotoState127Action,
	{_State108, FuncTypeType}:                     _GotoState121Action,
	{_State109, IdentifierToken}:                  _GotoState147Action,
	{_State109, StructToken}:                      _GotoState31Action,
	{_State109, EnumToken}:                        _GotoState111Action,
	{_State109, TraitToken}:                       _GotoState15Action,
	{_State109, LparenToken}:                      _GotoState114Action,
	{_State109, TildeTildeToken}:                  _GotoState116Action,
	{_State109, BitNegToken}:                      _GotoState110Action,
	{_State109, AtomTypeType}:                     _GotoState117Action,
	{_State109, TraitableTypeType}:                _GotoState130Action,
	{_State109, TraitAlgebraTypeType}:             _GotoState174Action,
	{_State109, ImplicitStructDefType}:            _GotoState123Action,
	{_State109, ExplicitStructDefType}:            _GotoState119Action,
	{_State109, ImplicitEnumDefType}:              _GotoState122Action,
	{_State109, ExplicitEnumDefType}:              _GotoState118Action,
	{_State109, TraitDefType}:                     _GotoState127Action,
	{_State110, IdentifierToken}:                  _GotoState147Action,
	{_State110, StructToken}:                      _GotoState31Action,
	{_State110, EnumToken}:                        _GotoState111Action,
	{_State110, TraitToken}:                       _GotoState15Action,
	{_State110, LparenToken}:                      _GotoState114Action,
	{_State110, AtomTypeType}:                     _GotoState175Action,
	{_State110, ImplicitStructDefType}:            _GotoState123Action,
	{_State110, ExplicitStructDefType}:            _GotoState119Action,
	{_State110, ImplicitEnumDefType}:              _GotoState122Action,
	{_State110, ExplicitEnumDefType}:              _GotoState118Action,
	{_State110, TraitDefType}:                     _GotoState127Action,
	{_State111, LparenToken}:                      _GotoState176Action,
	{_State112, IdentifierToken}:                  _GotoState177Action,
	{_State112, LparenToken}:                      _GotoState178Action,
	{_State113, IdentifierToken}:                  _GotoState147Action,
	{_State113, StructToken}:                      _GotoState31Action,
	{_State113, EnumToken}:                        _GotoState111Action,
	{_State113, TraitToken}:                       _GotoState15Action,
	{_State113, FuncToken}:                        _GotoState143Action,
	{_State113, LparenToken}:                      _GotoState114Action,
	{_State113, QuestionToken}:                    _GotoState115Action,
	{_State113, DollarLbracketToken}:              _GotoState67Action,
	{_State113, TildeTildeToken}:                  _GotoState116Action,
	{_State113, BitNegToken}:                      _GotoState110Action,
	{_State113, BitAndToken}:                      _GotoState109Action,
	{_State113, OptionalGenericBindingType}:       _GotoState179Action,
	{_State113, AtomTypeType}:                     _GotoState117Action,
	{_State113, TraitableTypeType}:                _GotoState130Action,
	{_State113, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State113, ValueTypeType}:                    _GotoState180Action,
	{_State113, ImplicitStructDefType}:            _GotoState123Action,
	{_State113, ExplicitStructDefType}:            _GotoState119Action,
	{_State113, ImplicitEnumDefType}:              _GotoState122Action,
	{_State113, ExplicitEnumDefType}:              _GotoState118Action,
	{_State113, TraitDefType}:                     _GotoState127Action,
	{_State113, FuncTypeType}:                     _GotoState121Action,
	{_State114, IdentifierToken}:                  _GotoState113Action,
	{_State114, StructToken}:                      _GotoState31Action,
	{_State114, EnumToken}:                        _GotoState111Action,
	{_State114, TraitToken}:                       _GotoState15Action,
	{_State114, FuncToken}:                        _GotoState143Action,
	{_State114, LparenToken}:                      _GotoState114Action,
	{_State114, QuestionToken}:                    _GotoState115Action,
	{_State114, TildeTildeToken}:                  _GotoState116Action,
	{_State114, BitNegToken}:                      _GotoState110Action,
	{_State114, BitAndToken}:                      _GotoState109Action,
	{_State114, AtomTypeType}:                     _GotoState117Action,
	{_State114, TraitableTypeType}:                _GotoState130Action,
	{_State114, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State114, ValueTypeType}:                    _GotoState131Action,
	{_State114, FieldDefType}:                     _GotoState182Action,
	{_State114, ImplicitFieldDefsType}:            _GotoState184Action,
	{_State114, OptionalImplicitFieldDefsType}:    _GotoState185Action,
	{_State114, ImplicitStructDefType}:            _GotoState123Action,
	{_State114, ExplicitStructDefType}:            _GotoState119Action,
	{_State114, EnumValueDefType}:                 _GotoState181Action,
	{_State114, ImplicitEnumValueDefsType}:        _GotoState183Action,
	{_State114, ImplicitEnumDefType}:              _GotoState122Action,
	{_State114, ExplicitEnumDefType}:              _GotoState118Action,
	{_State114, TraitDefType}:                     _GotoState127Action,
	{_State114, FuncTypeType}:                     _GotoState121Action,
	{_State116, IdentifierToken}:                  _GotoState147Action,
	{_State116, StructToken}:                      _GotoState31Action,
	{_State116, EnumToken}:                        _GotoState111Action,
	{_State116, TraitToken}:                       _GotoState15Action,
	{_State116, LparenToken}:                      _GotoState114Action,
	{_State116, AtomTypeType}:                     _GotoState186Action,
	{_State116, ImplicitStructDefType}:            _GotoState123Action,
	{_State116, ExplicitStructDefType}:            _GotoState119Action,
	{_State116, ImplicitEnumDefType}:              _GotoState122Action,
	{_State116, ExplicitEnumDefType}:              _GotoState118Action,
	{_State116, TraitDefType}:                     _GotoState127Action,
	{_State125, RparenToken}:                      _GotoState187Action,
	{_State126, AddToken}:                         _GotoState188Action,
	{_State126, SubToken}:                         _GotoState190Action,
	{_State126, MulToken}:                         _GotoState189Action,
	{_State128, NewlinesToken}:                    _GotoState192Action,
	{_State128, CommaToken}:                       _GotoState191Action,
	{_State132, IdentifierToken}:                  _GotoState147Action,
	{_State132, StructToken}:                      _GotoState31Action,
	{_State132, EnumToken}:                        _GotoState111Action,
	{_State132, TraitToken}:                       _GotoState15Action,
	{_State132, FuncToken}:                        _GotoState143Action,
	{_State132, LparenToken}:                      _GotoState114Action,
	{_State132, QuestionToken}:                    _GotoState115Action,
	{_State132, DotdotdotToken}:                   _GotoState193Action,
	{_State132, TildeTildeToken}:                  _GotoState116Action,
	{_State132, BitNegToken}:                      _GotoState110Action,
	{_State132, BitAndToken}:                      _GotoState109Action,
	{_State132, AtomTypeType}:                     _GotoState117Action,
	{_State132, TraitableTypeType}:                _GotoState130Action,
	{_State132, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State132, ValueTypeType}:                    _GotoState194Action,
	{_State132, ImplicitStructDefType}:            _GotoState123Action,
	{_State132, ExplicitStructDefType}:            _GotoState119Action,
	{_State132, ImplicitEnumDefType}:              _GotoState122Action,
	{_State132, ExplicitEnumDefType}:              _GotoState118Action,
	{_State132, TraitDefType}:                     _GotoState127Action,
	{_State132, FuncTypeType}:                     _GotoState121Action,
	{_State133, RparenToken}:                      _GotoState195Action,
	{_State134, DollarLbracketToken}:              _GotoState106Action,
	{_State134, OptionalGenericParametersType}:    _GotoState196Action,
	{_State135, RparenToken}:                      _GotoState197Action,
	{_State137, CommaToken}:                       _GotoState198Action,
	{_State138, IntegerLiteralToken}:              _GotoState23Action,
	{_State138, FloatLiteralToken}:                _GotoState20Action,
	{_State138, RuneLiteralToken}:                 _GotoState29Action,
	{_State138, StringLiteralToken}:               _GotoState30Action,
	{_State138, IdentifierToken}:                  _GotoState22Action,
	{_State138, TrueToken}:                        _GotoState33Action,
	{_State138, FalseToken}:                       _GotoState19Action,
	{_State138, StructToken}:                      _GotoState31Action,
	{_State138, FuncToken}:                        _GotoState21Action,
	{_State138, LabelDeclToken}:                   _GotoState24Action,
	{_State138, LparenToken}:                      _GotoState26Action,
	{_State138, NotToken}:                         _GotoState28Action,
	{_State138, SubToken}:                         _GotoState32Action,
	{_State138, MulToken}:                         _GotoState27Action,
	{_State138, BitNegToken}:                      _GotoState18Action,
	{_State138, BitAndToken}:                      _GotoState17Action,
	{_State138, LexErrorToken}:                    _GotoState25Action,
	{_State138, ExpressionType}:                   _GotoState199Action,
	{_State138, OptionalLabelDeclType}:            _GotoState46Action,
	{_State138, SequenceExprType}:                 _GotoState51Action,
	{_State138, BlockExprType}:                    _GotoState40Action,
	{_State138, CallExprType}:                     _GotoState41Action,
	{_State138, AtomExprType}:                     _GotoState39Action,
	{_State138, LiteralType}:                      _GotoState44Action,
	{_State138, AnonymousStructExprType}:          _GotoState38Action,
	{_State138, AccessExprType}:                   _GotoState34Action,
	{_State138, PostfixUnaryExprType}:             _GotoState48Action,
	{_State138, PrefixUnaryOpType}:                _GotoState50Action,
	{_State138, PrefixUnaryExprType}:              _GotoState49Action,
	{_State138, MulExprType}:                      _GotoState45Action,
	{_State138, AddExprType}:                      _GotoState35Action,
	{_State138, CmpExprType}:                      _GotoState42Action,
	{_State138, AndExprType}:                      _GotoState36Action,
	{_State138, OrExprType}:                       _GotoState47Action,
	{_State138, ExplicitStructDefType}:            _GotoState43Action,
	{_State138, AnonymousFuncExprType}:            _GotoState37Action,
	{_State139, IntegerLiteralToken}:              _GotoState23Action,
	{_State139, FloatLiteralToken}:                _GotoState20Action,
	{_State139, RuneLiteralToken}:                 _GotoState29Action,
	{_State139, StringLiteralToken}:               _GotoState30Action,
	{_State139, IdentifierToken}:                  _GotoState60Action,
	{_State139, TrueToken}:                        _GotoState33Action,
	{_State139, FalseToken}:                       _GotoState19Action,
	{_State139, StructToken}:                      _GotoState31Action,
	{_State139, FuncToken}:                        _GotoState21Action,
	{_State139, LabelDeclToken}:                   _GotoState24Action,
	{_State139, LparenToken}:                      _GotoState26Action,
	{_State139, NotToken}:                         _GotoState28Action,
	{_State139, SubToken}:                         _GotoState32Action,
	{_State139, MulToken}:                         _GotoState27Action,
	{_State139, BitNegToken}:                      _GotoState18Action,
	{_State139, BitAndToken}:                      _GotoState17Action,
	{_State139, LexErrorToken}:                    _GotoState25Action,
	{_State139, ExpressionType}:                   _GotoState64Action,
	{_State139, OptionalLabelDeclType}:            _GotoState46Action,
	{_State139, SequenceExprType}:                 _GotoState51Action,
	{_State139, BlockExprType}:                    _GotoState40Action,
	{_State139, CallExprType}:                     _GotoState41Action,
	{_State139, ArgumentType}:                     _GotoState200Action,
	{_State139, ColonExpressionsType}:             _GotoState63Action,
	{_State139, OptionalExpressionType}:           _GotoState65Action,
	{_State139, AtomExprType}:                     _GotoState39Action,
	{_State139, LiteralType}:                      _GotoState44Action,
	{_State139, AnonymousStructExprType}:          _GotoState38Action,
	{_State139, AccessExprType}:                   _GotoState34Action,
	{_State139, PostfixUnaryExprType}:             _GotoState48Action,
	{_State139, PrefixUnaryOpType}:                _GotoState50Action,
	{_State139, PrefixUnaryExprType}:              _GotoState49Action,
	{_State139, MulExprType}:                      _GotoState45Action,
	{_State139, AddExprType}:                      _GotoState35Action,
	{_State139, CmpExprType}:                      _GotoState42Action,
	{_State139, AndExprType}:                      _GotoState36Action,
	{_State139, OrExprType}:                       _GotoState47Action,
	{_State139, ExplicitStructDefType}:            _GotoState43Action,
	{_State139, AnonymousFuncExprType}:            _GotoState37Action,
	{_State141, IntegerLiteralToken}:              _GotoState23Action,
	{_State141, FloatLiteralToken}:                _GotoState20Action,
	{_State141, RuneLiteralToken}:                 _GotoState29Action,
	{_State141, StringLiteralToken}:               _GotoState30Action,
	{_State141, IdentifierToken}:                  _GotoState22Action,
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
	{_State141, ExpressionType}:                   _GotoState201Action,
	{_State141, OptionalLabelDeclType}:            _GotoState46Action,
	{_State141, SequenceExprType}:                 _GotoState51Action,
	{_State141, BlockExprType}:                    _GotoState40Action,
	{_State141, CallExprType}:                     _GotoState41Action,
	{_State141, OptionalExpressionType}:           _GotoState202Action,
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
	{_State142, IntegerLiteralToken}:              _GotoState23Action,
	{_State142, FloatLiteralToken}:                _GotoState20Action,
	{_State142, RuneLiteralToken}:                 _GotoState29Action,
	{_State142, StringLiteralToken}:               _GotoState30Action,
	{_State142, IdentifierToken}:                  _GotoState22Action,
	{_State142, TrueToken}:                        _GotoState33Action,
	{_State142, FalseToken}:                       _GotoState19Action,
	{_State142, StructToken}:                      _GotoState31Action,
	{_State142, FuncToken}:                        _GotoState21Action,
	{_State142, LabelDeclToken}:                   _GotoState24Action,
	{_State142, LparenToken}:                      _GotoState26Action,
	{_State142, NotToken}:                         _GotoState28Action,
	{_State142, SubToken}:                         _GotoState32Action,
	{_State142, MulToken}:                         _GotoState27Action,
	{_State142, BitNegToken}:                      _GotoState18Action,
	{_State142, BitAndToken}:                      _GotoState17Action,
	{_State142, LexErrorToken}:                    _GotoState25Action,
	{_State142, ExpressionType}:                   _GotoState201Action,
	{_State142, OptionalLabelDeclType}:            _GotoState46Action,
	{_State142, SequenceExprType}:                 _GotoState51Action,
	{_State142, BlockExprType}:                    _GotoState40Action,
	{_State142, CallExprType}:                     _GotoState41Action,
	{_State142, OptionalExpressionType}:           _GotoState203Action,
	{_State142, AtomExprType}:                     _GotoState39Action,
	{_State142, LiteralType}:                      _GotoState44Action,
	{_State142, AnonymousStructExprType}:          _GotoState38Action,
	{_State142, AccessExprType}:                   _GotoState34Action,
	{_State142, PostfixUnaryExprType}:             _GotoState48Action,
	{_State142, PrefixUnaryOpType}:                _GotoState50Action,
	{_State142, PrefixUnaryExprType}:              _GotoState49Action,
	{_State142, MulExprType}:                      _GotoState45Action,
	{_State142, AddExprType}:                      _GotoState35Action,
	{_State142, CmpExprType}:                      _GotoState42Action,
	{_State142, AndExprType}:                      _GotoState36Action,
	{_State142, OrExprType}:                       _GotoState47Action,
	{_State142, ExplicitStructDefType}:            _GotoState43Action,
	{_State142, AnonymousFuncExprType}:            _GotoState37Action,
	{_State143, LparenToken}:                      _GotoState178Action,
	{_State144, NewlinesToken}:                    _GotoState205Action,
	{_State144, CommaToken}:                       _GotoState204Action,
	{_State146, RparenToken}:                      _GotoState206Action,
	{_State147, DollarLbracketToken}:              _GotoState67Action,
	{_State147, OptionalGenericBindingType}:       _GotoState179Action,
	{_State148, CommaToken}:                       _GotoState207Action,
	{_State149, RbracketToken}:                    _GotoState208Action,
	{_State152, RbracketToken}:                    _GotoState209Action,
	{_State153, IntegerLiteralToken}:              _GotoState23Action,
	{_State153, FloatLiteralToken}:                _GotoState20Action,
	{_State153, RuneLiteralToken}:                 _GotoState29Action,
	{_State153, StringLiteralToken}:               _GotoState30Action,
	{_State153, IdentifierToken}:                  _GotoState60Action,
	{_State153, TrueToken}:                        _GotoState33Action,
	{_State153, FalseToken}:                       _GotoState19Action,
	{_State153, StructToken}:                      _GotoState31Action,
	{_State153, FuncToken}:                        _GotoState21Action,
	{_State153, LabelDeclToken}:                   _GotoState24Action,
	{_State153, LparenToken}:                      _GotoState26Action,
	{_State153, NotToken}:                         _GotoState28Action,
	{_State153, SubToken}:                         _GotoState32Action,
	{_State153, MulToken}:                         _GotoState27Action,
	{_State153, BitNegToken}:                      _GotoState18Action,
	{_State153, BitAndToken}:                      _GotoState17Action,
	{_State153, LexErrorToken}:                    _GotoState25Action,
	{_State153, ExpressionType}:                   _GotoState64Action,
	{_State153, OptionalLabelDeclType}:            _GotoState46Action,
	{_State153, SequenceExprType}:                 _GotoState51Action,
	{_State153, BlockExprType}:                    _GotoState40Action,
	{_State153, CallExprType}:                     _GotoState41Action,
	{_State153, OptionalArgumentsType}:            _GotoState211Action,
	{_State153, ArgumentsType}:                    _GotoState210Action,
	{_State153, ArgumentType}:                     _GotoState61Action,
	{_State153, ColonExpressionsType}:             _GotoState63Action,
	{_State153, OptionalExpressionType}:           _GotoState65Action,
	{_State153, AtomExprType}:                     _GotoState39Action,
	{_State153, LiteralType}:                      _GotoState44Action,
	{_State153, AnonymousStructExprType}:          _GotoState38Action,
	{_State153, AccessExprType}:                   _GotoState34Action,
	{_State153, PostfixUnaryExprType}:             _GotoState48Action,
	{_State153, PrefixUnaryOpType}:                _GotoState50Action,
	{_State153, PrefixUnaryExprType}:              _GotoState49Action,
	{_State153, MulExprType}:                      _GotoState45Action,
	{_State153, AddExprType}:                      _GotoState35Action,
	{_State153, CmpExprType}:                      _GotoState42Action,
	{_State153, AndExprType}:                      _GotoState36Action,
	{_State153, OrExprType}:                       _GotoState47Action,
	{_State153, ExplicitStructDefType}:            _GotoState43Action,
	{_State153, AnonymousFuncExprType}:            _GotoState37Action,
	{_State154, MulToken}:                         _GotoState91Action,
	{_State154, DivToken}:                         _GotoState89Action,
	{_State154, ModToken}:                         _GotoState90Action,
	{_State154, BitAndToken}:                      _GotoState86Action,
	{_State154, BitLshiftToken}:                   _GotoState87Action,
	{_State154, BitRshiftToken}:                   _GotoState88Action,
	{_State154, MulOpType}:                        _GotoState92Action,
	{_State155, EqualToken}:                       _GotoState78Action,
	{_State155, NotEqualToken}:                    _GotoState83Action,
	{_State155, LessToken}:                        _GotoState81Action,
	{_State155, LessOrEqualToken}:                 _GotoState82Action,
	{_State155, GreaterToken}:                     _GotoState79Action,
	{_State155, GreaterOrEqualToken}:              _GotoState80Action,
	{_State155, CmpOpType}:                        _GotoState84Action,
	{_State156, AddToken}:                         _GotoState72Action,
	{_State156, SubToken}:                         _GotoState75Action,
	{_State156, BitXorToken}:                      _GotoState74Action,
	{_State156, BitOrToken}:                       _GotoState73Action,
	{_State156, AddOpType}:                        _GotoState76Action,
	{_State157, RparenToken}:                      _GotoState212Action,
	{_State157, CommaToken}:                       _GotoState139Action,
	{_State159, ForToken}:                         _GotoState213Action,
	{_State160, IntegerLiteralToken}:              _GotoState23Action,
	{_State160, FloatLiteralToken}:                _GotoState20Action,
	{_State160, RuneLiteralToken}:                 _GotoState29Action,
	{_State160, StringLiteralToken}:               _GotoState30Action,
	{_State160, IdentifierToken}:                  _GotoState22Action,
	{_State160, TrueToken}:                        _GotoState33Action,
	{_State160, FalseToken}:                       _GotoState19Action,
	{_State160, StructToken}:                      _GotoState31Action,
	{_State160, FuncToken}:                        _GotoState21Action,
	{_State160, LabelDeclToken}:                   _GotoState24Action,
	{_State160, LparenToken}:                      _GotoState26Action,
	{_State160, NotToken}:                         _GotoState28Action,
	{_State160, SubToken}:                         _GotoState32Action,
	{_State160, MulToken}:                         _GotoState27Action,
	{_State160, BitNegToken}:                      _GotoState18Action,
	{_State160, BitAndToken}:                      _GotoState17Action,
	{_State160, LexErrorToken}:                    _GotoState25Action,
	{_State160, OptionalLabelDeclType}:            _GotoState103Action,
	{_State160, SequenceExprType}:                 _GotoState214Action,
	{_State160, BlockExprType}:                    _GotoState40Action,
	{_State160, CallExprType}:                     _GotoState41Action,
	{_State160, AtomExprType}:                     _GotoState39Action,
	{_State160, LiteralType}:                      _GotoState44Action,
	{_State160, AnonymousStructExprType}:          _GotoState38Action,
	{_State160, AccessExprType}:                   _GotoState34Action,
	{_State160, PostfixUnaryExprType}:             _GotoState48Action,
	{_State160, PrefixUnaryOpType}:                _GotoState50Action,
	{_State160, PrefixUnaryExprType}:              _GotoState49Action,
	{_State160, MulExprType}:                      _GotoState45Action,
	{_State160, AddExprType}:                      _GotoState35Action,
	{_State160, CmpExprType}:                      _GotoState42Action,
	{_State160, AndExprType}:                      _GotoState36Action,
	{_State160, OrExprType}:                       _GotoState47Action,
	{_State160, ExplicitStructDefType}:            _GotoState43Action,
	{_State160, AnonymousFuncExprType}:            _GotoState37Action,
	{_State161, IntegerLiteralToken}:              _GotoState23Action,
	{_State161, FloatLiteralToken}:                _GotoState20Action,
	{_State161, RuneLiteralToken}:                 _GotoState29Action,
	{_State161, StringLiteralToken}:               _GotoState30Action,
	{_State161, IdentifierToken}:                  _GotoState22Action,
	{_State161, TrueToken}:                        _GotoState33Action,
	{_State161, FalseToken}:                       _GotoState19Action,
	{_State161, StructToken}:                      _GotoState31Action,
	{_State161, FuncToken}:                        _GotoState21Action,
	{_State161, LabelDeclToken}:                   _GotoState24Action,
	{_State161, LparenToken}:                      _GotoState26Action,
	{_State161, NotToken}:                         _GotoState28Action,
	{_State161, SubToken}:                         _GotoState32Action,
	{_State161, MulToken}:                         _GotoState27Action,
	{_State161, BitNegToken}:                      _GotoState18Action,
	{_State161, BitAndToken}:                      _GotoState17Action,
	{_State161, LexErrorToken}:                    _GotoState25Action,
	{_State161, OptionalLabelDeclType}:            _GotoState103Action,
	{_State161, OptionalSequenceExprType}:         _GotoState215Action,
	{_State161, SequenceExprType}:                 _GotoState216Action,
	{_State161, BlockExprType}:                    _GotoState40Action,
	{_State161, CallExprType}:                     _GotoState41Action,
	{_State161, AtomExprType}:                     _GotoState39Action,
	{_State161, LiteralType}:                      _GotoState44Action,
	{_State161, AnonymousStructExprType}:          _GotoState38Action,
	{_State161, AccessExprType}:                   _GotoState34Action,
	{_State161, PostfixUnaryExprType}:             _GotoState48Action,
	{_State161, PrefixUnaryOpType}:                _GotoState50Action,
	{_State161, PrefixUnaryExprType}:              _GotoState49Action,
	{_State161, MulExprType}:                      _GotoState45Action,
	{_State161, AddExprType}:                      _GotoState35Action,
	{_State161, CmpExprType}:                      _GotoState42Action,
	{_State161, AndExprType}:                      _GotoState36Action,
	{_State161, OrExprType}:                       _GotoState47Action,
	{_State161, ExplicitStructDefType}:            _GotoState43Action,
	{_State161, AnonymousFuncExprType}:            _GotoState37Action,
	{_State162, DoToken}:                          _GotoState217Action,
	{_State163, LbraceToken}:                      _GotoState96Action,
	{_State163, BlockBodyType}:                    _GotoState218Action,
	{_State164, IntegerLiteralToken}:              _GotoState23Action,
	{_State164, FloatLiteralToken}:                _GotoState20Action,
	{_State164, RuneLiteralToken}:                 _GotoState29Action,
	{_State164, StringLiteralToken}:               _GotoState30Action,
	{_State164, IdentifierToken}:                  _GotoState22Action,
	{_State164, TrueToken}:                        _GotoState33Action,
	{_State164, FalseToken}:                       _GotoState19Action,
	{_State164, ReturnToken}:                      _GotoState223Action,
	{_State164, BreakToken}:                       _GotoState220Action,
	{_State164, ContinueToken}:                    _GotoState221Action,
	{_State164, UnsafeToken}:                      _GotoState224Action,
	{_State164, StructToken}:                      _GotoState31Action,
	{_State164, FuncToken}:                        _GotoState21Action,
	{_State164, AsyncToken}:                       _GotoState219Action,
	{_State164, LabelDeclToken}:                   _GotoState24Action,
	{_State164, RbraceToken}:                      _GotoState222Action,
	{_State164, LparenToken}:                      _GotoState26Action,
	{_State164, NotToken}:                         _GotoState28Action,
	{_State164, SubToken}:                         _GotoState32Action,
	{_State164, MulToken}:                         _GotoState27Action,
	{_State164, BitNegToken}:                      _GotoState18Action,
	{_State164, BitAndToken}:                      _GotoState17Action,
	{_State164, LexErrorToken}:                    _GotoState25Action,
	{_State164, ExpressionType}:                   _GotoState226Action,
	{_State164, OptionalLabelDeclType}:            _GotoState46Action,
	{_State164, SequenceExprType}:                 _GotoState51Action,
	{_State164, BlockExprType}:                    _GotoState40Action,
	{_State164, StatementType}:                    _GotoState230Action,
	{_State164, StatementBodyType}:                _GotoState231Action,
	{_State164, UnsafeStatementType}:              _GotoState232Action,
	{_State164, JumpStatementType}:                _GotoState228Action,
	{_State164, JumpTypeType}:                     _GotoState229Action,
	{_State164, ExpressionsType}:                  _GotoState227Action,
	{_State164, CallExprType}:                     _GotoState41Action,
	{_State164, AtomExprType}:                     _GotoState39Action,
	{_State164, LiteralType}:                      _GotoState44Action,
	{_State164, AnonymousStructExprType}:          _GotoState38Action,
	{_State164, AccessExprType}:                   _GotoState225Action,
	{_State164, PostfixUnaryExprType}:             _GotoState48Action,
	{_State164, PrefixUnaryOpType}:                _GotoState50Action,
	{_State164, PrefixUnaryExprType}:              _GotoState49Action,
	{_State164, MulExprType}:                      _GotoState45Action,
	{_State164, AddExprType}:                      _GotoState35Action,
	{_State164, CmpExprType}:                      _GotoState42Action,
	{_State164, AndExprType}:                      _GotoState36Action,
	{_State164, OrExprType}:                       _GotoState47Action,
	{_State164, ExplicitStructDefType}:            _GotoState43Action,
	{_State164, AnonymousFuncExprType}:            _GotoState37Action,
	{_State165, LbraceToken}:                      _GotoState233Action,
	{_State166, AndToken}:                         _GotoState77Action,
	{_State167, UnsafeToken}:                      _GotoState224Action,
	{_State167, RparenToken}:                      _GotoState234Action,
	{_State167, UnsafeStatementType}:              _GotoState237Action,
	{_State167, PackageStatementBodyType}:         _GotoState236Action,
	{_State167, PackageStatementType}:             _GotoState235Action,
	{_State168, IdentifierToken}:                  _GotoState147Action,
	{_State168, StructToken}:                      _GotoState31Action,
	{_State168, EnumToken}:                        _GotoState111Action,
	{_State168, TraitToken}:                       _GotoState15Action,
	{_State168, FuncToken}:                        _GotoState143Action,
	{_State168, LparenToken}:                      _GotoState114Action,
	{_State168, QuestionToken}:                    _GotoState115Action,
	{_State168, TildeTildeToken}:                  _GotoState116Action,
	{_State168, BitNegToken}:                      _GotoState110Action,
	{_State168, BitAndToken}:                      _GotoState109Action,
	{_State168, AtomTypeType}:                     _GotoState117Action,
	{_State168, TraitableTypeType}:                _GotoState130Action,
	{_State168, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State168, ValueTypeType}:                    _GotoState238Action,
	{_State168, ImplicitStructDefType}:            _GotoState123Action,
	{_State168, ExplicitStructDefType}:            _GotoState119Action,
	{_State168, ImplicitEnumDefType}:              _GotoState122Action,
	{_State168, ExplicitEnumDefType}:              _GotoState118Action,
	{_State168, TraitDefType}:                     _GotoState127Action,
	{_State168, FuncTypeType}:                     _GotoState121Action,
	{_State170, CommaToken}:                       _GotoState239Action,
	{_State171, RbracketToken}:                    _GotoState240Action,
	{_State173, ImplementsToken}:                  _GotoState241Action,
	{_State174, AddToken}:                         _GotoState188Action,
	{_State174, SubToken}:                         _GotoState190Action,
	{_State174, MulToken}:                         _GotoState189Action,
	{_State176, IdentifierToken}:                  _GotoState113Action,
	{_State176, StructToken}:                      _GotoState31Action,
	{_State176, EnumToken}:                        _GotoState111Action,
	{_State176, TraitToken}:                       _GotoState15Action,
	{_State176, FuncToken}:                        _GotoState143Action,
	{_State176, LparenToken}:                      _GotoState114Action,
	{_State176, QuestionToken}:                    _GotoState115Action,
	{_State176, TildeTildeToken}:                  _GotoState116Action,
	{_State176, BitNegToken}:                      _GotoState110Action,
	{_State176, BitAndToken}:                      _GotoState109Action,
	{_State176, AtomTypeType}:                     _GotoState117Action,
	{_State176, TraitableTypeType}:                _GotoState130Action,
	{_State176, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State176, ValueTypeType}:                    _GotoState131Action,
	{_State176, FieldDefType}:                     _GotoState244Action,
	{_State176, ImplicitStructDefType}:            _GotoState123Action,
	{_State176, ExplicitStructDefType}:            _GotoState119Action,
	{_State176, EnumValueDefType}:                 _GotoState242Action,
	{_State176, ImplicitEnumValueDefsType}:        _GotoState245Action,
	{_State176, ImplicitEnumDefType}:              _GotoState122Action,
	{_State176, ExplicitEnumValueDefsType}:        _GotoState243Action,
	{_State176, ExplicitEnumDefType}:              _GotoState118Action,
	{_State176, TraitDefType}:                     _GotoState127Action,
	{_State176, FuncTypeType}:                     _GotoState121Action,
	{_State177, LparenToken}:                      _GotoState246Action,
	{_State178, IdentifierToken}:                  _GotoState248Action,
	{_State178, StructToken}:                      _GotoState31Action,
	{_State178, EnumToken}:                        _GotoState111Action,
	{_State178, TraitToken}:                       _GotoState15Action,
	{_State178, FuncToken}:                        _GotoState143Action,
	{_State178, LparenToken}:                      _GotoState114Action,
	{_State178, QuestionToken}:                    _GotoState115Action,
	{_State178, DotdotdotToken}:                   _GotoState247Action,
	{_State178, TildeTildeToken}:                  _GotoState116Action,
	{_State178, BitNegToken}:                      _GotoState110Action,
	{_State178, BitAndToken}:                      _GotoState109Action,
	{_State178, AtomTypeType}:                     _GotoState117Action,
	{_State178, TraitableTypeType}:                _GotoState130Action,
	{_State178, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State178, ValueTypeType}:                    _GotoState252Action,
	{_State178, ImplicitStructDefType}:            _GotoState123Action,
	{_State178, ExplicitStructDefType}:            _GotoState119Action,
	{_State178, ImplicitEnumDefType}:              _GotoState122Action,
	{_State178, ExplicitEnumDefType}:              _GotoState118Action,
	{_State178, TraitDefType}:                     _GotoState127Action,
	{_State178, ParameterDeclType}:                _GotoState250Action,
	{_State178, ParameterDeclsType}:               _GotoState251Action,
	{_State178, OptionalParameterDeclsType}:       _GotoState249Action,
	{_State178, FuncTypeType}:                     _GotoState121Action,
	{_State181, OrToken}:                          _GotoState253Action,
	{_State182, AssignToken}:                      _GotoState254Action,
	{_State183, RparenToken}:                      _GotoState256Action,
	{_State183, OrToken}:                          _GotoState255Action,
	{_State184, CommaToken}:                       _GotoState257Action,
	{_State185, RparenToken}:                      _GotoState258Action,
	{_State188, IdentifierToken}:                  _GotoState147Action,
	{_State188, StructToken}:                      _GotoState31Action,
	{_State188, EnumToken}:                        _GotoState111Action,
	{_State188, TraitToken}:                       _GotoState15Action,
	{_State188, LparenToken}:                      _GotoState114Action,
	{_State188, TildeTildeToken}:                  _GotoState116Action,
	{_State188, BitNegToken}:                      _GotoState110Action,
	{_State188, AtomTypeType}:                     _GotoState117Action,
	{_State188, TraitableTypeType}:                _GotoState259Action,
	{_State188, ImplicitStructDefType}:            _GotoState123Action,
	{_State188, ExplicitStructDefType}:            _GotoState119Action,
	{_State188, ImplicitEnumDefType}:              _GotoState122Action,
	{_State188, ExplicitEnumDefType}:              _GotoState118Action,
	{_State188, TraitDefType}:                     _GotoState127Action,
	{_State189, IdentifierToken}:                  _GotoState147Action,
	{_State189, StructToken}:                      _GotoState31Action,
	{_State189, EnumToken}:                        _GotoState111Action,
	{_State189, TraitToken}:                       _GotoState15Action,
	{_State189, LparenToken}:                      _GotoState114Action,
	{_State189, TildeTildeToken}:                  _GotoState116Action,
	{_State189, BitNegToken}:                      _GotoState110Action,
	{_State189, AtomTypeType}:                     _GotoState117Action,
	{_State189, TraitableTypeType}:                _GotoState260Action,
	{_State189, ImplicitStructDefType}:            _GotoState123Action,
	{_State189, ExplicitStructDefType}:            _GotoState119Action,
	{_State189, ImplicitEnumDefType}:              _GotoState122Action,
	{_State189, ExplicitEnumDefType}:              _GotoState118Action,
	{_State189, TraitDefType}:                     _GotoState127Action,
	{_State190, IdentifierToken}:                  _GotoState147Action,
	{_State190, StructToken}:                      _GotoState31Action,
	{_State190, EnumToken}:                        _GotoState111Action,
	{_State190, TraitToken}:                       _GotoState15Action,
	{_State190, LparenToken}:                      _GotoState114Action,
	{_State190, TildeTildeToken}:                  _GotoState116Action,
	{_State190, BitNegToken}:                      _GotoState110Action,
	{_State190, AtomTypeType}:                     _GotoState117Action,
	{_State190, TraitableTypeType}:                _GotoState261Action,
	{_State190, ImplicitStructDefType}:            _GotoState123Action,
	{_State190, ExplicitStructDefType}:            _GotoState119Action,
	{_State190, ImplicitEnumDefType}:              _GotoState122Action,
	{_State190, ExplicitEnumDefType}:              _GotoState118Action,
	{_State190, TraitDefType}:                     _GotoState127Action,
	{_State191, IdentifierToken}:                  _GotoState113Action,
	{_State191, StructToken}:                      _GotoState31Action,
	{_State191, EnumToken}:                        _GotoState111Action,
	{_State191, TraitToken}:                       _GotoState15Action,
	{_State191, FuncToken}:                        _GotoState112Action,
	{_State191, LparenToken}:                      _GotoState114Action,
	{_State191, QuestionToken}:                    _GotoState115Action,
	{_State191, TildeTildeToken}:                  _GotoState116Action,
	{_State191, BitNegToken}:                      _GotoState110Action,
	{_State191, BitAndToken}:                      _GotoState109Action,
	{_State191, AtomTypeType}:                     _GotoState117Action,
	{_State191, TraitableTypeType}:                _GotoState130Action,
	{_State191, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State191, ValueTypeType}:                    _GotoState131Action,
	{_State191, FieldDefType}:                     _GotoState120Action,
	{_State191, ImplicitStructDefType}:            _GotoState123Action,
	{_State191, ExplicitStructDefType}:            _GotoState119Action,
	{_State191, ImplicitEnumDefType}:              _GotoState122Action,
	{_State191, ExplicitEnumDefType}:              _GotoState118Action,
	{_State191, TraitPropertyType}:                _GotoState262Action,
	{_State191, TraitDefType}:                     _GotoState127Action,
	{_State191, FuncTypeType}:                     _GotoState121Action,
	{_State191, MethodSignatureType}:              _GotoState124Action,
	{_State192, IdentifierToken}:                  _GotoState113Action,
	{_State192, StructToken}:                      _GotoState31Action,
	{_State192, EnumToken}:                        _GotoState111Action,
	{_State192, TraitToken}:                       _GotoState15Action,
	{_State192, FuncToken}:                        _GotoState112Action,
	{_State192, LparenToken}:                      _GotoState114Action,
	{_State192, QuestionToken}:                    _GotoState115Action,
	{_State192, TildeTildeToken}:                  _GotoState116Action,
	{_State192, BitNegToken}:                      _GotoState110Action,
	{_State192, BitAndToken}:                      _GotoState109Action,
	{_State192, AtomTypeType}:                     _GotoState117Action,
	{_State192, TraitableTypeType}:                _GotoState130Action,
	{_State192, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State192, ValueTypeType}:                    _GotoState131Action,
	{_State192, FieldDefType}:                     _GotoState120Action,
	{_State192, ImplicitStructDefType}:            _GotoState123Action,
	{_State192, ExplicitStructDefType}:            _GotoState119Action,
	{_State192, ImplicitEnumDefType}:              _GotoState122Action,
	{_State192, ExplicitEnumDefType}:              _GotoState118Action,
	{_State192, TraitPropertyType}:                _GotoState263Action,
	{_State192, TraitDefType}:                     _GotoState127Action,
	{_State192, FuncTypeType}:                     _GotoState121Action,
	{_State192, MethodSignatureType}:              _GotoState124Action,
	{_State193, IdentifierToken}:                  _GotoState147Action,
	{_State193, StructToken}:                      _GotoState31Action,
	{_State193, EnumToken}:                        _GotoState111Action,
	{_State193, TraitToken}:                       _GotoState15Action,
	{_State193, FuncToken}:                        _GotoState143Action,
	{_State193, LparenToken}:                      _GotoState114Action,
	{_State193, QuestionToken}:                    _GotoState115Action,
	{_State193, TildeTildeToken}:                  _GotoState116Action,
	{_State193, BitNegToken}:                      _GotoState110Action,
	{_State193, BitAndToken}:                      _GotoState109Action,
	{_State193, AtomTypeType}:                     _GotoState117Action,
	{_State193, TraitableTypeType}:                _GotoState130Action,
	{_State193, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State193, ValueTypeType}:                    _GotoState264Action,
	{_State193, ImplicitStructDefType}:            _GotoState123Action,
	{_State193, ExplicitStructDefType}:            _GotoState119Action,
	{_State193, ImplicitEnumDefType}:              _GotoState122Action,
	{_State193, ExplicitEnumDefType}:              _GotoState118Action,
	{_State193, TraitDefType}:                     _GotoState127Action,
	{_State193, FuncTypeType}:                     _GotoState121Action,
	{_State196, LparenToken}:                      _GotoState265Action,
	{_State197, IdentifierToken}:                  _GotoState147Action,
	{_State197, StructToken}:                      _GotoState31Action,
	{_State197, EnumToken}:                        _GotoState111Action,
	{_State197, TraitToken}:                       _GotoState15Action,
	{_State197, FuncToken}:                        _GotoState143Action,
	{_State197, LparenToken}:                      _GotoState114Action,
	{_State197, QuestionToken}:                    _GotoState115Action,
	{_State197, TildeTildeToken}:                  _GotoState116Action,
	{_State197, BitNegToken}:                      _GotoState110Action,
	{_State197, BitAndToken}:                      _GotoState109Action,
	{_State197, AtomTypeType}:                     _GotoState117Action,
	{_State197, TraitableTypeType}:                _GotoState130Action,
	{_State197, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State197, ValueTypeType}:                    _GotoState267Action,
	{_State197, ImplicitStructDefType}:            _GotoState123Action,
	{_State197, ExplicitStructDefType}:            _GotoState119Action,
	{_State197, ImplicitEnumDefType}:              _GotoState122Action,
	{_State197, ExplicitEnumDefType}:              _GotoState118Action,
	{_State197, TraitDefType}:                     _GotoState127Action,
	{_State197, ReturnTypeType}:                   _GotoState266Action,
	{_State197, FuncTypeType}:                     _GotoState121Action,
	{_State198, IdentifierToken}:                  _GotoState132Action,
	{_State198, ParameterDefType}:                 _GotoState268Action,
	{_State204, IdentifierToken}:                  _GotoState113Action,
	{_State204, StructToken}:                      _GotoState31Action,
	{_State204, EnumToken}:                        _GotoState111Action,
	{_State204, TraitToken}:                       _GotoState15Action,
	{_State204, FuncToken}:                        _GotoState143Action,
	{_State204, LparenToken}:                      _GotoState114Action,
	{_State204, QuestionToken}:                    _GotoState115Action,
	{_State204, TildeTildeToken}:                  _GotoState116Action,
	{_State204, BitNegToken}:                      _GotoState110Action,
	{_State204, BitAndToken}:                      _GotoState109Action,
	{_State204, AtomTypeType}:                     _GotoState117Action,
	{_State204, TraitableTypeType}:                _GotoState130Action,
	{_State204, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State204, ValueTypeType}:                    _GotoState131Action,
	{_State204, FieldDefType}:                     _GotoState269Action,
	{_State204, ImplicitStructDefType}:            _GotoState123Action,
	{_State204, ExplicitStructDefType}:            _GotoState119Action,
	{_State204, ImplicitEnumDefType}:              _GotoState122Action,
	{_State204, ExplicitEnumDefType}:              _GotoState118Action,
	{_State204, TraitDefType}:                     _GotoState127Action,
	{_State204, FuncTypeType}:                     _GotoState121Action,
	{_State205, IdentifierToken}:                  _GotoState113Action,
	{_State205, StructToken}:                      _GotoState31Action,
	{_State205, EnumToken}:                        _GotoState111Action,
	{_State205, TraitToken}:                       _GotoState15Action,
	{_State205, FuncToken}:                        _GotoState143Action,
	{_State205, LparenToken}:                      _GotoState114Action,
	{_State205, QuestionToken}:                    _GotoState115Action,
	{_State205, TildeTildeToken}:                  _GotoState116Action,
	{_State205, BitNegToken}:                      _GotoState110Action,
	{_State205, BitAndToken}:                      _GotoState109Action,
	{_State205, AtomTypeType}:                     _GotoState117Action,
	{_State205, TraitableTypeType}:                _GotoState130Action,
	{_State205, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State205, ValueTypeType}:                    _GotoState131Action,
	{_State205, FieldDefType}:                     _GotoState270Action,
	{_State205, ImplicitStructDefType}:            _GotoState123Action,
	{_State205, ExplicitStructDefType}:            _GotoState119Action,
	{_State205, ImplicitEnumDefType}:              _GotoState122Action,
	{_State205, ExplicitEnumDefType}:              _GotoState118Action,
	{_State205, TraitDefType}:                     _GotoState127Action,
	{_State205, FuncTypeType}:                     _GotoState121Action,
	{_State207, IdentifierToken}:                  _GotoState147Action,
	{_State207, StructToken}:                      _GotoState31Action,
	{_State207, EnumToken}:                        _GotoState111Action,
	{_State207, TraitToken}:                       _GotoState15Action,
	{_State207, FuncToken}:                        _GotoState143Action,
	{_State207, LparenToken}:                      _GotoState114Action,
	{_State207, QuestionToken}:                    _GotoState115Action,
	{_State207, TildeTildeToken}:                  _GotoState116Action,
	{_State207, BitNegToken}:                      _GotoState110Action,
	{_State207, BitAndToken}:                      _GotoState109Action,
	{_State207, AtomTypeType}:                     _GotoState117Action,
	{_State207, TraitableTypeType}:                _GotoState130Action,
	{_State207, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State207, ValueTypeType}:                    _GotoState271Action,
	{_State207, ImplicitStructDefType}:            _GotoState123Action,
	{_State207, ExplicitStructDefType}:            _GotoState119Action,
	{_State207, ImplicitEnumDefType}:              _GotoState122Action,
	{_State207, ExplicitEnumDefType}:              _GotoState118Action,
	{_State207, TraitDefType}:                     _GotoState127Action,
	{_State207, FuncTypeType}:                     _GotoState121Action,
	{_State210, CommaToken}:                       _GotoState139Action,
	{_State211, RparenToken}:                      _GotoState272Action,
	{_State213, IntegerLiteralToken}:              _GotoState23Action,
	{_State213, FloatLiteralToken}:                _GotoState20Action,
	{_State213, RuneLiteralToken}:                 _GotoState29Action,
	{_State213, StringLiteralToken}:               _GotoState30Action,
	{_State213, IdentifierToken}:                  _GotoState22Action,
	{_State213, TrueToken}:                        _GotoState33Action,
	{_State213, FalseToken}:                       _GotoState19Action,
	{_State213, StructToken}:                      _GotoState31Action,
	{_State213, FuncToken}:                        _GotoState21Action,
	{_State213, LabelDeclToken}:                   _GotoState24Action,
	{_State213, LparenToken}:                      _GotoState26Action,
	{_State213, NotToken}:                         _GotoState28Action,
	{_State213, SubToken}:                         _GotoState32Action,
	{_State213, MulToken}:                         _GotoState27Action,
	{_State213, BitNegToken}:                      _GotoState18Action,
	{_State213, BitAndToken}:                      _GotoState17Action,
	{_State213, LexErrorToken}:                    _GotoState25Action,
	{_State213, OptionalLabelDeclType}:            _GotoState103Action,
	{_State213, SequenceExprType}:                 _GotoState273Action,
	{_State213, BlockExprType}:                    _GotoState40Action,
	{_State213, CallExprType}:                     _GotoState41Action,
	{_State213, AtomExprType}:                     _GotoState39Action,
	{_State213, LiteralType}:                      _GotoState44Action,
	{_State213, AnonymousStructExprType}:          _GotoState38Action,
	{_State213, AccessExprType}:                   _GotoState34Action,
	{_State213, PostfixUnaryExprType}:             _GotoState48Action,
	{_State213, PrefixUnaryOpType}:                _GotoState50Action,
	{_State213, PrefixUnaryExprType}:              _GotoState49Action,
	{_State213, MulExprType}:                      _GotoState45Action,
	{_State213, AddExprType}:                      _GotoState35Action,
	{_State213, CmpExprType}:                      _GotoState42Action,
	{_State213, AndExprType}:                      _GotoState36Action,
	{_State213, OrExprType}:                       _GotoState47Action,
	{_State213, ExplicitStructDefType}:            _GotoState43Action,
	{_State213, AnonymousFuncExprType}:            _GotoState37Action,
	{_State214, DoToken}:                          _GotoState274Action,
	{_State215, SemicolonToken}:                   _GotoState275Action,
	{_State217, LbraceToken}:                      _GotoState96Action,
	{_State217, BlockBodyType}:                    _GotoState276Action,
	{_State218, ElseToken}:                        _GotoState277Action,
	{_State219, IntegerLiteralToken}:              _GotoState23Action,
	{_State219, FloatLiteralToken}:                _GotoState20Action,
	{_State219, RuneLiteralToken}:                 _GotoState29Action,
	{_State219, StringLiteralToken}:               _GotoState30Action,
	{_State219, IdentifierToken}:                  _GotoState22Action,
	{_State219, TrueToken}:                        _GotoState33Action,
	{_State219, FalseToken}:                       _GotoState19Action,
	{_State219, StructToken}:                      _GotoState31Action,
	{_State219, FuncToken}:                        _GotoState21Action,
	{_State219, LabelDeclToken}:                   _GotoState24Action,
	{_State219, LparenToken}:                      _GotoState26Action,
	{_State219, LexErrorToken}:                    _GotoState25Action,
	{_State219, OptionalLabelDeclType}:            _GotoState103Action,
	{_State219, BlockExprType}:                    _GotoState40Action,
	{_State219, CallExprType}:                     _GotoState279Action,
	{_State219, AtomExprType}:                     _GotoState39Action,
	{_State219, LiteralType}:                      _GotoState44Action,
	{_State219, AnonymousStructExprType}:          _GotoState38Action,
	{_State219, AccessExprType}:                   _GotoState278Action,
	{_State219, ExplicitStructDefType}:            _GotoState43Action,
	{_State219, AnonymousFuncExprType}:            _GotoState37Action,
	{_State224, LessToken}:                        _GotoState280Action,
	{_State225, LbracketToken}:                    _GotoState69Action,
	{_State225, DotToken}:                         _GotoState68Action,
	{_State225, QuestionToken}:                    _GotoState70Action,
	{_State225, DollarLbracketToken}:              _GotoState67Action,
	{_State225, AddAssignToken}:                   _GotoState281Action,
	{_State225, SubAssignToken}:                   _GotoState292Action,
	{_State225, MulAssignToken}:                   _GotoState291Action,
	{_State225, DivAssignToken}:                   _GotoState289Action,
	{_State225, ModAssignToken}:                   _GotoState290Action,
	{_State225, AddOneAssignToken}:                _GotoState282Action,
	{_State225, SubOneAssignToken}:                _GotoState293Action,
	{_State225, BitNegAssignToken}:                _GotoState285Action,
	{_State225, BitAndAssignToken}:                _GotoState283Action,
	{_State225, BitOrAssignToken}:                 _GotoState286Action,
	{_State225, BitXorAssignToken}:                _GotoState288Action,
	{_State225, BitLshiftAssignToken}:             _GotoState284Action,
	{_State225, BitRshiftAssignToken}:             _GotoState287Action,
	{_State225, UnaryOpAssignType}:                _GotoState295Action,
	{_State225, BinaryOpAssignType}:               _GotoState294Action,
	{_State225, OptionalGenericBindingType}:       _GotoState71Action,
	{_State227, CommaToken}:                       _GotoState296Action,
	{_State229, JumpLabelToken}:                   _GotoState297Action,
	{_State229, OptionalJumpLabelType}:            _GotoState298Action,
	{_State231, NewlinesToken}:                    _GotoState299Action,
	{_State231, SemicolonToken}:                   _GotoState300Action,
	{_State233, CaseToken}:                        _GotoState301Action,
	{_State236, NewlinesToken}:                    _GotoState302Action,
	{_State236, SemicolonToken}:                   _GotoState303Action,
	{_State239, IdentifierToken}:                  _GotoState168Action,
	{_State239, GenericParameterDefType}:          _GotoState304Action,
	{_State241, IdentifierToken}:                  _GotoState147Action,
	{_State241, StructToken}:                      _GotoState31Action,
	{_State241, EnumToken}:                        _GotoState111Action,
	{_State241, TraitToken}:                       _GotoState15Action,
	{_State241, FuncToken}:                        _GotoState143Action,
	{_State241, LparenToken}:                      _GotoState114Action,
	{_State241, QuestionToken}:                    _GotoState115Action,
	{_State241, TildeTildeToken}:                  _GotoState116Action,
	{_State241, BitNegToken}:                      _GotoState110Action,
	{_State241, BitAndToken}:                      _GotoState109Action,
	{_State241, AtomTypeType}:                     _GotoState117Action,
	{_State241, TraitableTypeType}:                _GotoState130Action,
	{_State241, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State241, ValueTypeType}:                    _GotoState305Action,
	{_State241, ImplicitStructDefType}:            _GotoState123Action,
	{_State241, ExplicitStructDefType}:            _GotoState119Action,
	{_State241, ImplicitEnumDefType}:              _GotoState122Action,
	{_State241, ExplicitEnumDefType}:              _GotoState118Action,
	{_State241, TraitDefType}:                     _GotoState127Action,
	{_State241, FuncTypeType}:                     _GotoState121Action,
	{_State242, NewlinesToken}:                    _GotoState306Action,
	{_State242, OrToken}:                          _GotoState307Action,
	{_State243, RparenToken}:                      _GotoState308Action,
	{_State244, AssignToken}:                      _GotoState254Action,
	{_State245, NewlinesToken}:                    _GotoState309Action,
	{_State245, OrToken}:                          _GotoState310Action,
	{_State246, IdentifierToken}:                  _GotoState248Action,
	{_State246, StructToken}:                      _GotoState31Action,
	{_State246, EnumToken}:                        _GotoState111Action,
	{_State246, TraitToken}:                       _GotoState15Action,
	{_State246, FuncToken}:                        _GotoState143Action,
	{_State246, LparenToken}:                      _GotoState114Action,
	{_State246, QuestionToken}:                    _GotoState115Action,
	{_State246, DotdotdotToken}:                   _GotoState247Action,
	{_State246, TildeTildeToken}:                  _GotoState116Action,
	{_State246, BitNegToken}:                      _GotoState110Action,
	{_State246, BitAndToken}:                      _GotoState109Action,
	{_State246, AtomTypeType}:                     _GotoState117Action,
	{_State246, TraitableTypeType}:                _GotoState130Action,
	{_State246, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State246, ValueTypeType}:                    _GotoState252Action,
	{_State246, ImplicitStructDefType}:            _GotoState123Action,
	{_State246, ExplicitStructDefType}:            _GotoState119Action,
	{_State246, ImplicitEnumDefType}:              _GotoState122Action,
	{_State246, ExplicitEnumDefType}:              _GotoState118Action,
	{_State246, TraitDefType}:                     _GotoState127Action,
	{_State246, ParameterDeclType}:                _GotoState250Action,
	{_State246, ParameterDeclsType}:               _GotoState251Action,
	{_State246, OptionalParameterDeclsType}:       _GotoState311Action,
	{_State246, FuncTypeType}:                     _GotoState121Action,
	{_State247, IdentifierToken}:                  _GotoState147Action,
	{_State247, StructToken}:                      _GotoState31Action,
	{_State247, EnumToken}:                        _GotoState111Action,
	{_State247, TraitToken}:                       _GotoState15Action,
	{_State247, FuncToken}:                        _GotoState143Action,
	{_State247, LparenToken}:                      _GotoState114Action,
	{_State247, QuestionToken}:                    _GotoState115Action,
	{_State247, TildeTildeToken}:                  _GotoState116Action,
	{_State247, BitNegToken}:                      _GotoState110Action,
	{_State247, BitAndToken}:                      _GotoState109Action,
	{_State247, AtomTypeType}:                     _GotoState117Action,
	{_State247, TraitableTypeType}:                _GotoState130Action,
	{_State247, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State247, ValueTypeType}:                    _GotoState312Action,
	{_State247, ImplicitStructDefType}:            _GotoState123Action,
	{_State247, ExplicitStructDefType}:            _GotoState119Action,
	{_State247, ImplicitEnumDefType}:              _GotoState122Action,
	{_State247, ExplicitEnumDefType}:              _GotoState118Action,
	{_State247, TraitDefType}:                     _GotoState127Action,
	{_State247, FuncTypeType}:                     _GotoState121Action,
	{_State248, IdentifierToken}:                  _GotoState147Action,
	{_State248, StructToken}:                      _GotoState31Action,
	{_State248, EnumToken}:                        _GotoState111Action,
	{_State248, TraitToken}:                       _GotoState15Action,
	{_State248, FuncToken}:                        _GotoState143Action,
	{_State248, LparenToken}:                      _GotoState114Action,
	{_State248, QuestionToken}:                    _GotoState115Action,
	{_State248, DollarLbracketToken}:              _GotoState67Action,
	{_State248, DotdotdotToken}:                   _GotoState313Action,
	{_State248, TildeTildeToken}:                  _GotoState116Action,
	{_State248, BitNegToken}:                      _GotoState110Action,
	{_State248, BitAndToken}:                      _GotoState109Action,
	{_State248, OptionalGenericBindingType}:       _GotoState179Action,
	{_State248, AtomTypeType}:                     _GotoState117Action,
	{_State248, TraitableTypeType}:                _GotoState130Action,
	{_State248, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State248, ValueTypeType}:                    _GotoState314Action,
	{_State248, ImplicitStructDefType}:            _GotoState123Action,
	{_State248, ExplicitStructDefType}:            _GotoState119Action,
	{_State248, ImplicitEnumDefType}:              _GotoState122Action,
	{_State248, ExplicitEnumDefType}:              _GotoState118Action,
	{_State248, TraitDefType}:                     _GotoState127Action,
	{_State248, FuncTypeType}:                     _GotoState121Action,
	{_State249, RparenToken}:                      _GotoState315Action,
	{_State251, CommaToken}:                       _GotoState316Action,
	{_State253, IdentifierToken}:                  _GotoState113Action,
	{_State253, StructToken}:                      _GotoState31Action,
	{_State253, EnumToken}:                        _GotoState111Action,
	{_State253, TraitToken}:                       _GotoState15Action,
	{_State253, FuncToken}:                        _GotoState143Action,
	{_State253, LparenToken}:                      _GotoState114Action,
	{_State253, QuestionToken}:                    _GotoState115Action,
	{_State253, TildeTildeToken}:                  _GotoState116Action,
	{_State253, BitNegToken}:                      _GotoState110Action,
	{_State253, BitAndToken}:                      _GotoState109Action,
	{_State253, AtomTypeType}:                     _GotoState117Action,
	{_State253, TraitableTypeType}:                _GotoState130Action,
	{_State253, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State253, ValueTypeType}:                    _GotoState131Action,
	{_State253, FieldDefType}:                     _GotoState244Action,
	{_State253, ImplicitStructDefType}:            _GotoState123Action,
	{_State253, ExplicitStructDefType}:            _GotoState119Action,
	{_State253, EnumValueDefType}:                 _GotoState317Action,
	{_State253, ImplicitEnumDefType}:              _GotoState122Action,
	{_State253, ExplicitEnumDefType}:              _GotoState118Action,
	{_State253, TraitDefType}:                     _GotoState127Action,
	{_State253, FuncTypeType}:                     _GotoState121Action,
	{_State254, DefaultToken}:                     _GotoState318Action,
	{_State255, IdentifierToken}:                  _GotoState113Action,
	{_State255, StructToken}:                      _GotoState31Action,
	{_State255, EnumToken}:                        _GotoState111Action,
	{_State255, TraitToken}:                       _GotoState15Action,
	{_State255, FuncToken}:                        _GotoState143Action,
	{_State255, LparenToken}:                      _GotoState114Action,
	{_State255, QuestionToken}:                    _GotoState115Action,
	{_State255, TildeTildeToken}:                  _GotoState116Action,
	{_State255, BitNegToken}:                      _GotoState110Action,
	{_State255, BitAndToken}:                      _GotoState109Action,
	{_State255, AtomTypeType}:                     _GotoState117Action,
	{_State255, TraitableTypeType}:                _GotoState130Action,
	{_State255, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State255, ValueTypeType}:                    _GotoState131Action,
	{_State255, FieldDefType}:                     _GotoState244Action,
	{_State255, ImplicitStructDefType}:            _GotoState123Action,
	{_State255, ExplicitStructDefType}:            _GotoState119Action,
	{_State255, EnumValueDefType}:                 _GotoState319Action,
	{_State255, ImplicitEnumDefType}:              _GotoState122Action,
	{_State255, ExplicitEnumDefType}:              _GotoState118Action,
	{_State255, TraitDefType}:                     _GotoState127Action,
	{_State255, FuncTypeType}:                     _GotoState121Action,
	{_State257, IdentifierToken}:                  _GotoState113Action,
	{_State257, StructToken}:                      _GotoState31Action,
	{_State257, EnumToken}:                        _GotoState111Action,
	{_State257, TraitToken}:                       _GotoState15Action,
	{_State257, FuncToken}:                        _GotoState143Action,
	{_State257, LparenToken}:                      _GotoState114Action,
	{_State257, QuestionToken}:                    _GotoState115Action,
	{_State257, TildeTildeToken}:                  _GotoState116Action,
	{_State257, BitNegToken}:                      _GotoState110Action,
	{_State257, BitAndToken}:                      _GotoState109Action,
	{_State257, AtomTypeType}:                     _GotoState117Action,
	{_State257, TraitableTypeType}:                _GotoState130Action,
	{_State257, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State257, ValueTypeType}:                    _GotoState131Action,
	{_State257, FieldDefType}:                     _GotoState320Action,
	{_State257, ImplicitStructDefType}:            _GotoState123Action,
	{_State257, ExplicitStructDefType}:            _GotoState119Action,
	{_State257, ImplicitEnumDefType}:              _GotoState122Action,
	{_State257, ExplicitEnumDefType}:              _GotoState118Action,
	{_State257, TraitDefType}:                     _GotoState127Action,
	{_State257, FuncTypeType}:                     _GotoState121Action,
	{_State265, IdentifierToken}:                  _GotoState132Action,
	{_State265, ParameterDefType}:                 _GotoState136Action,
	{_State265, ParameterDefsType}:                _GotoState137Action,
	{_State265, OptionalParameterDefsType}:        _GotoState321Action,
	{_State266, LbraceToken}:                      _GotoState96Action,
	{_State266, BlockBodyType}:                    _GotoState322Action,
	{_State274, LbraceToken}:                      _GotoState96Action,
	{_State274, BlockBodyType}:                    _GotoState323Action,
	{_State275, IntegerLiteralToken}:              _GotoState23Action,
	{_State275, FloatLiteralToken}:                _GotoState20Action,
	{_State275, RuneLiteralToken}:                 _GotoState29Action,
	{_State275, StringLiteralToken}:               _GotoState30Action,
	{_State275, IdentifierToken}:                  _GotoState22Action,
	{_State275, TrueToken}:                        _GotoState33Action,
	{_State275, FalseToken}:                       _GotoState19Action,
	{_State275, StructToken}:                      _GotoState31Action,
	{_State275, FuncToken}:                        _GotoState21Action,
	{_State275, LabelDeclToken}:                   _GotoState24Action,
	{_State275, LparenToken}:                      _GotoState26Action,
	{_State275, NotToken}:                         _GotoState28Action,
	{_State275, SubToken}:                         _GotoState32Action,
	{_State275, MulToken}:                         _GotoState27Action,
	{_State275, BitNegToken}:                      _GotoState18Action,
	{_State275, BitAndToken}:                      _GotoState17Action,
	{_State275, LexErrorToken}:                    _GotoState25Action,
	{_State275, OptionalLabelDeclType}:            _GotoState103Action,
	{_State275, OptionalSequenceExprType}:         _GotoState324Action,
	{_State275, SequenceExprType}:                 _GotoState216Action,
	{_State275, BlockExprType}:                    _GotoState40Action,
	{_State275, CallExprType}:                     _GotoState41Action,
	{_State275, AtomExprType}:                     _GotoState39Action,
	{_State275, LiteralType}:                      _GotoState44Action,
	{_State275, AnonymousStructExprType}:          _GotoState38Action,
	{_State275, AccessExprType}:                   _GotoState34Action,
	{_State275, PostfixUnaryExprType}:             _GotoState48Action,
	{_State275, PrefixUnaryOpType}:                _GotoState50Action,
	{_State275, PrefixUnaryExprType}:              _GotoState49Action,
	{_State275, MulExprType}:                      _GotoState45Action,
	{_State275, AddExprType}:                      _GotoState35Action,
	{_State275, CmpExprType}:                      _GotoState42Action,
	{_State275, AndExprType}:                      _GotoState36Action,
	{_State275, OrExprType}:                       _GotoState47Action,
	{_State275, ExplicitStructDefType}:            _GotoState43Action,
	{_State275, AnonymousFuncExprType}:            _GotoState37Action,
	{_State277, IfToken}:                          _GotoState95Action,
	{_State277, LbraceToken}:                      _GotoState96Action,
	{_State277, IfExprType}:                       _GotoState326Action,
	{_State277, BlockBodyType}:                    _GotoState325Action,
	{_State278, LbracketToken}:                    _GotoState69Action,
	{_State278, DotToken}:                         _GotoState68Action,
	{_State278, DollarLbracketToken}:              _GotoState67Action,
	{_State278, OptionalGenericBindingType}:       _GotoState71Action,
	{_State280, IdentifierToken}:                  _GotoState327Action,
	{_State294, IntegerLiteralToken}:              _GotoState23Action,
	{_State294, FloatLiteralToken}:                _GotoState20Action,
	{_State294, RuneLiteralToken}:                 _GotoState29Action,
	{_State294, StringLiteralToken}:               _GotoState30Action,
	{_State294, IdentifierToken}:                  _GotoState22Action,
	{_State294, TrueToken}:                        _GotoState33Action,
	{_State294, FalseToken}:                       _GotoState19Action,
	{_State294, StructToken}:                      _GotoState31Action,
	{_State294, FuncToken}:                        _GotoState21Action,
	{_State294, LabelDeclToken}:                   _GotoState24Action,
	{_State294, LparenToken}:                      _GotoState26Action,
	{_State294, NotToken}:                         _GotoState28Action,
	{_State294, SubToken}:                         _GotoState32Action,
	{_State294, MulToken}:                         _GotoState27Action,
	{_State294, BitNegToken}:                      _GotoState18Action,
	{_State294, BitAndToken}:                      _GotoState17Action,
	{_State294, LexErrorToken}:                    _GotoState25Action,
	{_State294, ExpressionType}:                   _GotoState328Action,
	{_State294, OptionalLabelDeclType}:            _GotoState46Action,
	{_State294, SequenceExprType}:                 _GotoState51Action,
	{_State294, BlockExprType}:                    _GotoState40Action,
	{_State294, CallExprType}:                     _GotoState41Action,
	{_State294, AtomExprType}:                     _GotoState39Action,
	{_State294, LiteralType}:                      _GotoState44Action,
	{_State294, AnonymousStructExprType}:          _GotoState38Action,
	{_State294, AccessExprType}:                   _GotoState34Action,
	{_State294, PostfixUnaryExprType}:             _GotoState48Action,
	{_State294, PrefixUnaryOpType}:                _GotoState50Action,
	{_State294, PrefixUnaryExprType}:              _GotoState49Action,
	{_State294, MulExprType}:                      _GotoState45Action,
	{_State294, AddExprType}:                      _GotoState35Action,
	{_State294, CmpExprType}:                      _GotoState42Action,
	{_State294, AndExprType}:                      _GotoState36Action,
	{_State294, OrExprType}:                       _GotoState47Action,
	{_State294, ExplicitStructDefType}:            _GotoState43Action,
	{_State294, AnonymousFuncExprType}:            _GotoState37Action,
	{_State296, IntegerLiteralToken}:              _GotoState23Action,
	{_State296, FloatLiteralToken}:                _GotoState20Action,
	{_State296, RuneLiteralToken}:                 _GotoState29Action,
	{_State296, StringLiteralToken}:               _GotoState30Action,
	{_State296, IdentifierToken}:                  _GotoState22Action,
	{_State296, TrueToken}:                        _GotoState33Action,
	{_State296, FalseToken}:                       _GotoState19Action,
	{_State296, StructToken}:                      _GotoState31Action,
	{_State296, FuncToken}:                        _GotoState21Action,
	{_State296, LabelDeclToken}:                   _GotoState24Action,
	{_State296, LparenToken}:                      _GotoState26Action,
	{_State296, NotToken}:                         _GotoState28Action,
	{_State296, SubToken}:                         _GotoState32Action,
	{_State296, MulToken}:                         _GotoState27Action,
	{_State296, BitNegToken}:                      _GotoState18Action,
	{_State296, BitAndToken}:                      _GotoState17Action,
	{_State296, LexErrorToken}:                    _GotoState25Action,
	{_State296, ExpressionType}:                   _GotoState329Action,
	{_State296, OptionalLabelDeclType}:            _GotoState46Action,
	{_State296, SequenceExprType}:                 _GotoState51Action,
	{_State296, BlockExprType}:                    _GotoState40Action,
	{_State296, CallExprType}:                     _GotoState41Action,
	{_State296, AtomExprType}:                     _GotoState39Action,
	{_State296, LiteralType}:                      _GotoState44Action,
	{_State296, AnonymousStructExprType}:          _GotoState38Action,
	{_State296, AccessExprType}:                   _GotoState34Action,
	{_State296, PostfixUnaryExprType}:             _GotoState48Action,
	{_State296, PrefixUnaryOpType}:                _GotoState50Action,
	{_State296, PrefixUnaryExprType}:              _GotoState49Action,
	{_State296, MulExprType}:                      _GotoState45Action,
	{_State296, AddExprType}:                      _GotoState35Action,
	{_State296, CmpExprType}:                      _GotoState42Action,
	{_State296, AndExprType}:                      _GotoState36Action,
	{_State296, OrExprType}:                       _GotoState47Action,
	{_State296, ExplicitStructDefType}:            _GotoState43Action,
	{_State296, AnonymousFuncExprType}:            _GotoState37Action,
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
	{_State298, ExpressionType}:                   _GotoState226Action,
	{_State298, OptionalLabelDeclType}:            _GotoState46Action,
	{_State298, SequenceExprType}:                 _GotoState51Action,
	{_State298, BlockExprType}:                    _GotoState40Action,
	{_State298, ExpressionsType}:                  _GotoState330Action,
	{_State298, OptionalExpressionsType}:          _GotoState331Action,
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
	{_State301, DefaultToken}:                     _GotoState332Action,
	{_State306, IdentifierToken}:                  _GotoState113Action,
	{_State306, StructToken}:                      _GotoState31Action,
	{_State306, EnumToken}:                        _GotoState111Action,
	{_State306, TraitToken}:                       _GotoState15Action,
	{_State306, FuncToken}:                        _GotoState143Action,
	{_State306, LparenToken}:                      _GotoState114Action,
	{_State306, QuestionToken}:                    _GotoState115Action,
	{_State306, TildeTildeToken}:                  _GotoState116Action,
	{_State306, BitNegToken}:                      _GotoState110Action,
	{_State306, BitAndToken}:                      _GotoState109Action,
	{_State306, AtomTypeType}:                     _GotoState117Action,
	{_State306, TraitableTypeType}:                _GotoState130Action,
	{_State306, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State306, ValueTypeType}:                    _GotoState131Action,
	{_State306, FieldDefType}:                     _GotoState244Action,
	{_State306, ImplicitStructDefType}:            _GotoState123Action,
	{_State306, ExplicitStructDefType}:            _GotoState119Action,
	{_State306, EnumValueDefType}:                 _GotoState333Action,
	{_State306, ImplicitEnumDefType}:              _GotoState122Action,
	{_State306, ExplicitEnumDefType}:              _GotoState118Action,
	{_State306, TraitDefType}:                     _GotoState127Action,
	{_State306, FuncTypeType}:                     _GotoState121Action,
	{_State307, IdentifierToken}:                  _GotoState113Action,
	{_State307, StructToken}:                      _GotoState31Action,
	{_State307, EnumToken}:                        _GotoState111Action,
	{_State307, TraitToken}:                       _GotoState15Action,
	{_State307, FuncToken}:                        _GotoState143Action,
	{_State307, LparenToken}:                      _GotoState114Action,
	{_State307, QuestionToken}:                    _GotoState115Action,
	{_State307, TildeTildeToken}:                  _GotoState116Action,
	{_State307, BitNegToken}:                      _GotoState110Action,
	{_State307, BitAndToken}:                      _GotoState109Action,
	{_State307, AtomTypeType}:                     _GotoState117Action,
	{_State307, TraitableTypeType}:                _GotoState130Action,
	{_State307, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State307, ValueTypeType}:                    _GotoState131Action,
	{_State307, FieldDefType}:                     _GotoState244Action,
	{_State307, ImplicitStructDefType}:            _GotoState123Action,
	{_State307, ExplicitStructDefType}:            _GotoState119Action,
	{_State307, EnumValueDefType}:                 _GotoState334Action,
	{_State307, ImplicitEnumDefType}:              _GotoState122Action,
	{_State307, ExplicitEnumDefType}:              _GotoState118Action,
	{_State307, TraitDefType}:                     _GotoState127Action,
	{_State307, FuncTypeType}:                     _GotoState121Action,
	{_State309, IdentifierToken}:                  _GotoState113Action,
	{_State309, StructToken}:                      _GotoState31Action,
	{_State309, EnumToken}:                        _GotoState111Action,
	{_State309, TraitToken}:                       _GotoState15Action,
	{_State309, FuncToken}:                        _GotoState143Action,
	{_State309, LparenToken}:                      _GotoState114Action,
	{_State309, QuestionToken}:                    _GotoState115Action,
	{_State309, TildeTildeToken}:                  _GotoState116Action,
	{_State309, BitNegToken}:                      _GotoState110Action,
	{_State309, BitAndToken}:                      _GotoState109Action,
	{_State309, AtomTypeType}:                     _GotoState117Action,
	{_State309, TraitableTypeType}:                _GotoState130Action,
	{_State309, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State309, ValueTypeType}:                    _GotoState131Action,
	{_State309, FieldDefType}:                     _GotoState244Action,
	{_State309, ImplicitStructDefType}:            _GotoState123Action,
	{_State309, ExplicitStructDefType}:            _GotoState119Action,
	{_State309, EnumValueDefType}:                 _GotoState335Action,
	{_State309, ImplicitEnumDefType}:              _GotoState122Action,
	{_State309, ExplicitEnumDefType}:              _GotoState118Action,
	{_State309, TraitDefType}:                     _GotoState127Action,
	{_State309, FuncTypeType}:                     _GotoState121Action,
	{_State310, IdentifierToken}:                  _GotoState113Action,
	{_State310, StructToken}:                      _GotoState31Action,
	{_State310, EnumToken}:                        _GotoState111Action,
	{_State310, TraitToken}:                       _GotoState15Action,
	{_State310, FuncToken}:                        _GotoState143Action,
	{_State310, LparenToken}:                      _GotoState114Action,
	{_State310, QuestionToken}:                    _GotoState115Action,
	{_State310, TildeTildeToken}:                  _GotoState116Action,
	{_State310, BitNegToken}:                      _GotoState110Action,
	{_State310, BitAndToken}:                      _GotoState109Action,
	{_State310, AtomTypeType}:                     _GotoState117Action,
	{_State310, TraitableTypeType}:                _GotoState130Action,
	{_State310, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State310, ValueTypeType}:                    _GotoState131Action,
	{_State310, FieldDefType}:                     _GotoState244Action,
	{_State310, ImplicitStructDefType}:            _GotoState123Action,
	{_State310, ExplicitStructDefType}:            _GotoState119Action,
	{_State310, EnumValueDefType}:                 _GotoState336Action,
	{_State310, ImplicitEnumDefType}:              _GotoState122Action,
	{_State310, ExplicitEnumDefType}:              _GotoState118Action,
	{_State310, TraitDefType}:                     _GotoState127Action,
	{_State310, FuncTypeType}:                     _GotoState121Action,
	{_State311, RparenToken}:                      _GotoState337Action,
	{_State313, IdentifierToken}:                  _GotoState147Action,
	{_State313, StructToken}:                      _GotoState31Action,
	{_State313, EnumToken}:                        _GotoState111Action,
	{_State313, TraitToken}:                       _GotoState15Action,
	{_State313, FuncToken}:                        _GotoState143Action,
	{_State313, LparenToken}:                      _GotoState114Action,
	{_State313, QuestionToken}:                    _GotoState115Action,
	{_State313, TildeTildeToken}:                  _GotoState116Action,
	{_State313, BitNegToken}:                      _GotoState110Action,
	{_State313, BitAndToken}:                      _GotoState109Action,
	{_State313, AtomTypeType}:                     _GotoState117Action,
	{_State313, TraitableTypeType}:                _GotoState130Action,
	{_State313, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State313, ValueTypeType}:                    _GotoState338Action,
	{_State313, ImplicitStructDefType}:            _GotoState123Action,
	{_State313, ExplicitStructDefType}:            _GotoState119Action,
	{_State313, ImplicitEnumDefType}:              _GotoState122Action,
	{_State313, ExplicitEnumDefType}:              _GotoState118Action,
	{_State313, TraitDefType}:                     _GotoState127Action,
	{_State313, FuncTypeType}:                     _GotoState121Action,
	{_State315, IdentifierToken}:                  _GotoState147Action,
	{_State315, StructToken}:                      _GotoState31Action,
	{_State315, EnumToken}:                        _GotoState111Action,
	{_State315, TraitToken}:                       _GotoState15Action,
	{_State315, FuncToken}:                        _GotoState143Action,
	{_State315, LparenToken}:                      _GotoState114Action,
	{_State315, QuestionToken}:                    _GotoState115Action,
	{_State315, TildeTildeToken}:                  _GotoState116Action,
	{_State315, BitNegToken}:                      _GotoState110Action,
	{_State315, BitAndToken}:                      _GotoState109Action,
	{_State315, AtomTypeType}:                     _GotoState117Action,
	{_State315, TraitableTypeType}:                _GotoState130Action,
	{_State315, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State315, ValueTypeType}:                    _GotoState267Action,
	{_State315, ImplicitStructDefType}:            _GotoState123Action,
	{_State315, ExplicitStructDefType}:            _GotoState119Action,
	{_State315, ImplicitEnumDefType}:              _GotoState122Action,
	{_State315, ExplicitEnumDefType}:              _GotoState118Action,
	{_State315, TraitDefType}:                     _GotoState127Action,
	{_State315, ReturnTypeType}:                   _GotoState339Action,
	{_State315, FuncTypeType}:                     _GotoState121Action,
	{_State316, IdentifierToken}:                  _GotoState248Action,
	{_State316, StructToken}:                      _GotoState31Action,
	{_State316, EnumToken}:                        _GotoState111Action,
	{_State316, TraitToken}:                       _GotoState15Action,
	{_State316, FuncToken}:                        _GotoState143Action,
	{_State316, LparenToken}:                      _GotoState114Action,
	{_State316, QuestionToken}:                    _GotoState115Action,
	{_State316, DotdotdotToken}:                   _GotoState247Action,
	{_State316, TildeTildeToken}:                  _GotoState116Action,
	{_State316, BitNegToken}:                      _GotoState110Action,
	{_State316, BitAndToken}:                      _GotoState109Action,
	{_State316, AtomTypeType}:                     _GotoState117Action,
	{_State316, TraitableTypeType}:                _GotoState130Action,
	{_State316, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State316, ValueTypeType}:                    _GotoState252Action,
	{_State316, ImplicitStructDefType}:            _GotoState123Action,
	{_State316, ExplicitStructDefType}:            _GotoState119Action,
	{_State316, ImplicitEnumDefType}:              _GotoState122Action,
	{_State316, ExplicitEnumDefType}:              _GotoState118Action,
	{_State316, TraitDefType}:                     _GotoState127Action,
	{_State316, ParameterDeclType}:                _GotoState340Action,
	{_State316, FuncTypeType}:                     _GotoState121Action,
	{_State321, RparenToken}:                      _GotoState341Action,
	{_State324, DoToken}:                          _GotoState342Action,
	{_State327, GreaterToken}:                     _GotoState343Action,
	{_State330, CommaToken}:                       _GotoState296Action,
	{_State332, RbraceToken}:                      _GotoState344Action,
	{_State337, IdentifierToken}:                  _GotoState147Action,
	{_State337, StructToken}:                      _GotoState31Action,
	{_State337, EnumToken}:                        _GotoState111Action,
	{_State337, TraitToken}:                       _GotoState15Action,
	{_State337, FuncToken}:                        _GotoState143Action,
	{_State337, LparenToken}:                      _GotoState114Action,
	{_State337, QuestionToken}:                    _GotoState115Action,
	{_State337, TildeTildeToken}:                  _GotoState116Action,
	{_State337, BitNegToken}:                      _GotoState110Action,
	{_State337, BitAndToken}:                      _GotoState109Action,
	{_State337, AtomTypeType}:                     _GotoState117Action,
	{_State337, TraitableTypeType}:                _GotoState130Action,
	{_State337, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State337, ValueTypeType}:                    _GotoState267Action,
	{_State337, ImplicitStructDefType}:            _GotoState123Action,
	{_State337, ExplicitStructDefType}:            _GotoState119Action,
	{_State337, ImplicitEnumDefType}:              _GotoState122Action,
	{_State337, ExplicitEnumDefType}:              _GotoState118Action,
	{_State337, TraitDefType}:                     _GotoState127Action,
	{_State337, ReturnTypeType}:                   _GotoState345Action,
	{_State337, FuncTypeType}:                     _GotoState121Action,
	{_State341, IdentifierToken}:                  _GotoState147Action,
	{_State341, StructToken}:                      _GotoState31Action,
	{_State341, EnumToken}:                        _GotoState111Action,
	{_State341, TraitToken}:                       _GotoState15Action,
	{_State341, FuncToken}:                        _GotoState143Action,
	{_State341, LparenToken}:                      _GotoState114Action,
	{_State341, QuestionToken}:                    _GotoState115Action,
	{_State341, TildeTildeToken}:                  _GotoState116Action,
	{_State341, BitNegToken}:                      _GotoState110Action,
	{_State341, BitAndToken}:                      _GotoState109Action,
	{_State341, AtomTypeType}:                     _GotoState117Action,
	{_State341, TraitableTypeType}:                _GotoState130Action,
	{_State341, TraitAlgebraTypeType}:             _GotoState126Action,
	{_State341, ValueTypeType}:                    _GotoState267Action,
	{_State341, ImplicitStructDefType}:            _GotoState123Action,
	{_State341, ExplicitStructDefType}:            _GotoState119Action,
	{_State341, ImplicitEnumDefType}:              _GotoState122Action,
	{_State341, ExplicitEnumDefType}:              _GotoState118Action,
	{_State341, TraitDefType}:                     _GotoState127Action,
	{_State341, ReturnTypeType}:                   _GotoState346Action,
	{_State341, FuncTypeType}:                     _GotoState121Action,
	{_State342, LbraceToken}:                      _GotoState96Action,
	{_State342, BlockBodyType}:                    _GotoState347Action,
	{_State343, StringLiteralToken}:               _GotoState348Action,
	{_State346, LbraceToken}:                      _GotoState96Action,
	{_State346, BlockBodyType}:                    _GotoState349Action,
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
	{_State117, _WildcardMarker}:                  _ReduceAtomTypeToTraitableTypeAction,
	{_State118, _WildcardMarker}:                  _ReduceExplicitEnumDefToAtomTypeAction,
	{_State119, _WildcardMarker}:                  _ReduceExplicitStructDefToAtomTypeAction,
	{_State120, _WildcardMarker}:                  _ReduceFieldDefToTraitPropertyAction,
	{_State121, _EndMarker}:                       _ReduceFuncTypeToValueTypeAction,
	{_State122, _WildcardMarker}:                  _ReduceImplicitEnumDefToAtomTypeAction,
	{_State123, _WildcardMarker}:                  _ReduceImplicitStructDefToAtomTypeAction,
	{_State124, _WildcardMarker}:                  _ReduceMethodSignatureToTraitPropertyAction,
	{_State126, _EndMarker}:                       _ReduceTraitAlgebraTypeToValueTypeAction,
	{_State127, _WildcardMarker}:                  _ReduceTraitDefToAtomTypeAction,
	{_State128, RparenToken}:                      _ReduceTraitPropertiesToOptionalTraitPropertiesAction,
	{_State129, _WildcardMarker}:                  _ReduceTraitPropertyToTraitPropertiesAction,
	{_State130, _WildcardMarker}:                  _ReduceTraitableTypeToTraitAlgebraTypeAction,
	{_State131, _WildcardMarker}:                  _ReduceImplicitToFieldDefAction,
	{_State134, LparenToken}:                      _ReduceNilToOptionalGenericParametersAction,
	{_State136, _WildcardMarker}:                  _ReduceParameterDefToParameterDefsAction,
	{_State137, RparenToken}:                      _ReduceParameterDefsToOptionalParameterDefsAction,
	{_State138, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State139, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State139, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State140, _WildcardMarker}:                  _ReduceImplicitToAnonymousStructExprAction,
	{_State141, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State141, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State141, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State141, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State141, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State142, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State142, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State142, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State142, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State142, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State144, RparenToken}:                      _ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction,
	{_State145, _WildcardMarker}:                  _ReduceFieldDefToExplicitFieldDefsAction,
	{_State147, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State148, RbracketToken}:                    _ReduceGenericArgumentsToOptionalGenericArgumentsAction,
	{_State150, _WildcardMarker}:                  _ReduceValueTypeToGenericArgumentsAction,
	{_State151, _WildcardMarker}:                  _ReduceAccessToAccessExprAction,
	{_State153, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State153, RparenToken}:                      _ReduceNilToOptionalArgumentsAction,
	{_State153, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State154, _WildcardMarker}:                  _ReduceOpToAddExprAction,
	{_State155, _WildcardMarker}:                  _ReduceOpToAndExprAction,
	{_State156, _WildcardMarker}:                  _ReduceOpToCmpExprAction,
	{_State158, _WildcardMarker}:                  _ReduceOpToMulExprAction,
	{_State159, _WildcardMarker}:                  _ReduceInfiniteToLoopExprAction,
	{_State160, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State161, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State161, SemicolonToken}:                   _ReduceNilToOptionalSequenceExprAction,
	{_State164, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State166, _WildcardMarker}:                  _ReduceOpToOrExprAction,
	{_State168, _WildcardMarker}:                  _ReduceUnconstrainedToGenericParameterDefAction,
	{_State169, _WildcardMarker}:                  _ReduceGenericParameterDefToGenericParameterDefsAction,
	{_State170, RbracketToken}:                    _ReduceGenericParameterDefsToOptionalGenericParameterDefsAction,
	{_State172, _EndMarker}:                       _ReduceAliasToTypeDefAction,
	{_State173, _EndMarker}:                       _ReduceDefinitionToTypeDefAction,
	{_State174, _EndMarker}:                       _ReduceReferenceToValueTypeAction,
	{_State175, _WildcardMarker}:                  _ReducePublicMethodsTraitToTraitableTypeAction,
	{_State178, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State179, _WildcardMarker}:                  _ReduceNamedToAtomTypeAction,
	{_State180, _WildcardMarker}:                  _ReduceExplicitToFieldDefAction,
	{_State182, _WildcardMarker}:                  _ReduceFieldDefToImplicitFieldDefsAction,
	{_State182, OrToken}:                          _ReduceFieldDefToEnumValueDefAction,
	{_State184, RparenToken}:                      _ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction,
	{_State186, _WildcardMarker}:                  _ReducePublicTraitToTraitableTypeAction,
	{_State187, _EndMarker}:                       _ReduceToTraitDefAction,
	{_State194, _WildcardMarker}:                  _ReduceArgToParameterDefAction,
	{_State195, IdentifierToken}:                  _ReduceReceiverToOptionalReceiverAction,
	{_State197, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State199, _WildcardMarker}:                  _ReduceNamedToArgumentAction,
	{_State200, _WildcardMarker}:                  _ReduceAddToArgumentsAction,
	{_State201, _WildcardMarker}:                  _ReduceExpressionToOptionalExpressionAction,
	{_State202, _WildcardMarker}:                  _ReduceAddToColonExpressionsAction,
	{_State203, _WildcardMarker}:                  _ReducePairToColonExpressionsAction,
	{_State206, _WildcardMarker}:                  _ReduceToExplicitStructDefAction,
	{_State208, _WildcardMarker}:                  _ReduceBindingToOptionalGenericBindingAction,
	{_State209, _WildcardMarker}:                  _ReduceIndexToAccessExprAction,
	{_State210, RparenToken}:                      _ReduceArgumentsToOptionalArgumentsAction,
	{_State212, _WildcardMarker}:                  _ReduceExplicitToAnonymousStructExprAction,
	{_State213, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State216, DoToken}:                          _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State218, _WildcardMarker}:                  _ReduceNoElseToIfExprAction,
	{_State219, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State220, _WildcardMarker}:                  _ReduceBreakToJumpTypeAction,
	{_State221, _WildcardMarker}:                  _ReduceContinueToJumpTypeAction,
	{_State222, _EndMarker}:                       _ReduceToBlockBodyAction,
	{_State223, _WildcardMarker}:                  _ReduceReturnToJumpTypeAction,
	{_State225, _WildcardMarker}:                  _ReduceAccessExprToPostfixUnaryExprAction,
	{_State225, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State226, _WildcardMarker}:                  _ReduceExpressionToExpressionsAction,
	{_State227, _WildcardMarker}:                  _ReduceExpressionOrImplicitStructToStatementBodyAction,
	{_State228, _WildcardMarker}:                  _ReduceJumpStatementToStatementBodyAction,
	{_State229, _WildcardMarker}:                  _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State230, _WildcardMarker}:                  _ReduceAddToStatementsAction,
	{_State232, _WildcardMarker}:                  _ReduceUnsafeStatementToStatementBodyAction,
	{_State234, _EndMarker}:                       _ReduceWithSpecToPackageDefAction,
	{_State235, _WildcardMarker}:                  _ReduceAddToPackageStatementsAction,
	{_State237, _WildcardMarker}:                  _ReduceToPackageStatementBodyAction,
	{_State238, _WildcardMarker}:                  _ReduceConstrainedToGenericParameterDefAction,
	{_State240, _WildcardMarker}:                  _ReduceGenericToOptionalGenericParametersAction,
	{_State244, _WildcardMarker}:                  _ReduceFieldDefToEnumValueDefAction,
	{_State246, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State248, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State250, _WildcardMarker}:                  _ReduceParameterDeclToParameterDeclsAction,
	{_State251, RparenToken}:                      _ReduceParameterDeclsToOptionalParameterDeclsAction,
	{_State252, _WildcardMarker}:                  _ReduceUnamedToParameterDeclAction,
	{_State256, _WildcardMarker}:                  _ReduceToImplicitEnumDefAction,
	{_State258, _WildcardMarker}:                  _ReduceToImplicitStructDefAction,
	{_State259, _WildcardMarker}:                  _ReduceUnionToTraitAlgebraTypeAction,
	{_State260, _WildcardMarker}:                  _ReduceIntersectToTraitAlgebraTypeAction,
	{_State261, _WildcardMarker}:                  _ReduceDifferenceToTraitAlgebraTypeAction,
	{_State262, _WildcardMarker}:                  _ReduceExplicitToTraitPropertiesAction,
	{_State263, _WildcardMarker}:                  _ReduceImplicitToTraitPropertiesAction,
	{_State264, _WildcardMarker}:                  _ReduceVarargToParameterDefAction,
	{_State265, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State267, _EndMarker}:                       _ReduceValueTypeToReturnTypeAction,
	{_State268, _WildcardMarker}:                  _ReduceAddToParameterDefsAction,
	{_State269, _WildcardMarker}:                  _ReduceExplicitToExplicitFieldDefsAction,
	{_State270, _WildcardMarker}:                  _ReduceImplicitToExplicitFieldDefsAction,
	{_State271, _WildcardMarker}:                  _ReduceAddToGenericArgumentsAction,
	{_State272, _WildcardMarker}:                  _ReduceToCallExprAction,
	{_State273, _EndMarker}:                       _ReduceDoWhileToLoopExprAction,
	{_State275, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State275, DoToken}:                          _ReduceNilToOptionalSequenceExprAction,
	{_State276, _EndMarker}:                       _ReduceWhileToLoopExprAction,
	{_State278, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State279, NewlinesToken}:                    _ReduceAsyncToStatementBodyAction,
	{_State279, SemicolonToken}:                   _ReduceAsyncToStatementBodyAction,
	{_State279, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State281, _WildcardMarker}:                  _ReduceAddAssignToBinaryOpAssignAction,
	{_State282, _WildcardMarker}:                  _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State283, _WildcardMarker}:                  _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State284, _WildcardMarker}:                  _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State285, _WildcardMarker}:                  _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State286, _WildcardMarker}:                  _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State287, _WildcardMarker}:                  _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State288, _WildcardMarker}:                  _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State289, _WildcardMarker}:                  _ReduceDivAssignToBinaryOpAssignAction,
	{_State290, _WildcardMarker}:                  _ReduceModAssignToBinaryOpAssignAction,
	{_State291, _WildcardMarker}:                  _ReduceMulAssignToBinaryOpAssignAction,
	{_State292, _WildcardMarker}:                  _ReduceSubAssignToBinaryOpAssignAction,
	{_State293, _WildcardMarker}:                  _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State294, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State295, _WildcardMarker}:                  _ReduceUnaryOpAssignStatementToStatementBodyAction,
	{_State296, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State297, _WildcardMarker}:                  _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State298, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State298, NewlinesToken}:                    _ReduceNilToOptionalExpressionsAction,
	{_State298, SemicolonToken}:                   _ReduceNilToOptionalExpressionsAction,
	{_State299, _WildcardMarker}:                  _ReduceImplicitToStatementAction,
	{_State300, _WildcardMarker}:                  _ReduceExplicitToStatementAction,
	{_State302, _WildcardMarker}:                  _ReduceImplicitToPackageStatementAction,
	{_State303, _WildcardMarker}:                  _ReduceExplicitToPackageStatementAction,
	{_State304, _WildcardMarker}:                  _ReduceAddToGenericParameterDefsAction,
	{_State305, _EndMarker}:                       _ReduceConstrainedDefToTypeDefAction,
	{_State308, _WildcardMarker}:                  _ReduceToExplicitEnumDefAction,
	{_State312, _WildcardMarker}:                  _ReduceUnnamedVarargToParameterDeclAction,
	{_State314, _WildcardMarker}:                  _ReduceArgToParameterDeclAction,
	{_State315, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State317, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State318, _WildcardMarker}:                  _ReduceDefaultToEnumValueDefAction,
	{_State319, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State320, _WildcardMarker}:                  _ReduceAddToImplicitFieldDefsAction,
	{_State322, _WildcardMarker}:                  _ReduceToAnonymousFuncExprAction,
	{_State323, _EndMarker}:                       _ReduceIteratorToLoopExprAction,
	{_State325, _EndMarker}:                       _ReduceIfElseToIfExprAction,
	{_State326, _EndMarker}:                       _ReduceMultiIfElseToIfExprAction,
	{_State328, _WildcardMarker}:                  _ReduceBinaryOpAssignStatementToStatementBodyAction,
	{_State329, _WildcardMarker}:                  _ReduceAddToExpressionsAction,
	{_State330, _WildcardMarker}:                  _ReduceExpressionsToOptionalExpressionsAction,
	{_State331, _WildcardMarker}:                  _ReduceToJumpStatementAction,
	{_State333, RparenToken}:                      _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State334, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State334, RparenToken}:                      _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State335, RparenToken}:                      _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State336, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State336, RparenToken}:                      _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State337, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State338, _WildcardMarker}:                  _ReduceVarargToParameterDeclAction,
	{_State339, _EndMarker}:                       _ReduceToFuncTypeAction,
	{_State340, _WildcardMarker}:                  _ReduceAddToParameterDeclsAction,
	{_State341, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State344, _EndMarker}:                       _ReduceToSwitchExprAction,
	{_State345, _WildcardMarker}:                  _ReduceToMethodSignatureAction,
	{_State347, _EndMarker}:                       _ReduceForToLoopExprAction,
	{_State348, _WildcardMarker}:                  _ReduceToUnsafeStatementAction,
	{_State349, _EndMarker}:                       _ReduceToNamedFuncDefAction,
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
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 112
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 131
      field_def -> State 120
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_property -> State 129
      trait_properties -> State 128
      optional_trait_properties -> State 125
      trait_def -> State 127
      func_type -> State 121
      method_signature -> State 124

  State 57:
    Kernel Items:
      optional_receiver: LPAREN.parameter_def RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 132
      parameter_def -> State 133

  State 58:
    Kernel Items:
      named_func_def: FUNC optional_receiver.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 134

  State 59:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 132
      parameter_def -> State 136
      parameter_defs -> State 137
      optional_parameter_defs -> State 135

  State 60:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expression
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      ASSIGN -> State 138

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
      RPAREN -> State 140
      COMMA -> State 139

  State 63:
    Kernel Items:
      argument: colon_expressions., *
      colon_expressions: colon_expressions.COLON optional_expression
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 141

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
      COLON -> State 142

  State 66:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN.optional_explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      IDENTIFIER -> State 113
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 131
      field_def -> State 145
      implicit_struct_def -> State 123
      explicit_field_defs -> State 144
      optional_explicit_field_defs -> State 146
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 67:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      optional_generic_arguments -> State 149
      generic_arguments -> State 148
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 150
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 68:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 151

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
      argument -> State 152
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
      LPAREN -> State 153

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
      mul_expr -> State 154
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
      cmp_expr -> State 155
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
      add_expr -> State 156
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
      arguments -> State 157
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
      prefix_unary_expr -> State 158
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
      block_body -> State 159

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
      IN -> State 160
      STRUCT -> State 31
      FUNC -> State 21
      LABEL_DECL -> State 24
      LPAREN -> State 26
      SEMICOLON -> State 161
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      optional_label_decl -> State 103
      sequence_expr -> State 162
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
      sequence_expr -> State 163
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
      statements -> State 164

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
      and_expr -> State 166
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
      package_statements -> State 167

  State 106:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.optional_generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 168
      generic_parameter_def -> State 169
      generic_parameter_defs -> State 170
      optional_generic_parameter_defs -> State 171

  State 107:
    Kernel Items:
      type_def: TYPE IDENTIFIER EQUAL.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 172
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 108:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 173
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 109:
    Kernel Items:
      value_type: BIT_AND.trait_algebra_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      LPAREN -> State 114
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 174
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127

  State 110:
    Kernel Items:
      traitable_type: BIT_NEG.atom_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      LPAREN -> State 114
      atom_type -> State 175
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127

  State 111:
    Kernel Items:
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 176

  State 112:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 177
      LPAREN -> State 178

  State 113:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      DOLLAR_LBRACKET -> State 67
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      optional_generic_binding -> State 179
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 180
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 114:
    Kernel Items:
      implicit_struct_def: LPAREN.optional_implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      IDENTIFIER -> State 113
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 131
      field_def -> State 182
      implicit_field_defs -> State 184
      optional_implicit_field_defs -> State 185
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      enum_value_def -> State 181
      implicit_enum_value_defs -> State 183
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

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
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      LPAREN -> State 114
      atom_type -> State 186
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127

  State 117:
    Kernel Items:
      traitable_type: atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 118:
    Kernel Items:
      atom_type: explicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 119:
    Kernel Items:
      atom_type: explicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 121:
    Kernel Items:
      value_type: func_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      (nil)

  State 122:
    Kernel Items:
      atom_type: implicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 123:
    Kernel Items:
      atom_type: implicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 187

  State 126:
    Kernel Items:
      trait_algebra_type: trait_algebra_type.MUL traitable_type
      trait_algebra_type: trait_algebra_type.ADD traitable_type
      trait_algebra_type: trait_algebra_type.SUB traitable_type
      value_type: trait_algebra_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      ADD -> State 188
      SUB -> State 190
      MUL -> State 189

  State 127:
    Kernel Items:
      atom_type: trait_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 128:
    Kernel Items:
      trait_properties: trait_properties.NEWLINES trait_property
      trait_properties: trait_properties.COMMA trait_property
      optional_trait_properties: trait_properties., RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      NEWLINES -> State 192
      COMMA -> State 191

  State 129:
    Kernel Items:
      trait_properties: trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 130:
    Kernel Items:
      trait_algebra_type: traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 131:
    Kernel Items:
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 132:
    Kernel Items:
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      DOTDOTDOT -> State 193
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 194
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 133:
    Kernel Items:
      optional_receiver: LPAREN parameter_def.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 195

  State 134:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 106
      optional_generic_parameters -> State 196

  State 135:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 197

  State 136:
    Kernel Items:
      parameter_defs: parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 137:
    Kernel Items:
      parameter_defs: parameter_defs.COMMA parameter_def
      optional_parameter_defs: parameter_defs., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      COMMA -> State 198

  State 138:
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
      expression -> State 199
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

  State 139:
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
      argument -> State 200
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

  State 140:
    Kernel Items:
      anonymous_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 141:
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
      expression -> State 201
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      call_expr -> State 41
      optional_expression -> State 202
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
      expression -> State 201
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      call_expr -> State 41
      optional_expression -> State 203
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

  State 143:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 178

  State 144:
    Kernel Items:
      explicit_field_defs: explicit_field_defs.NEWLINES field_def
      explicit_field_defs: explicit_field_defs.COMMA field_def
      optional_explicit_field_defs: explicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      NEWLINES -> State 205
      COMMA -> State 204

  State 145:
    Kernel Items:
      explicit_field_defs: field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 146:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 206

  State 147:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 67
      optional_generic_binding -> State 179

  State 148:
    Kernel Items:
      optional_generic_arguments: generic_arguments., RBRACKET
      generic_arguments: generic_arguments.COMMA value_type
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 207

  State 149:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 208

  State 150:
    Kernel Items:
      generic_arguments: value_type., *
    Reduce:
      * -> [generic_arguments]
    Goto:
      (nil)

  State 151:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 152:
    Kernel Items:
      access_expr: access_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 209

  State 153:
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
      optional_arguments -> State 211
      arguments -> State 210
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

  State 154:
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

  State 155:
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

  State 156:
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

  State 157:
    Kernel Items:
      arguments: arguments.COMMA argument
      anonymous_struct_expr: explicit_struct_def LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 212
      COMMA -> State 139

  State 158:
    Kernel Items:
      mul_expr: mul_expr mul_op prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 159:
    Kernel Items:
      loop_expr: DO block_body., *
      loop_expr: DO block_body.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 213

  State 160:
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
      sequence_expr -> State 214
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

  State 161:
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
      optional_sequence_expr -> State 215
      sequence_expr -> State 216
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

  State 162:
    Kernel Items:
      loop_expr: FOR sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 217

  State 163:
    Kernel Items:
      if_expr: IF sequence_expr.block_body
      if_expr: IF sequence_expr.block_body ELSE block_body
      if_expr: IF sequence_expr.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 96
      block_body -> State 218

  State 164:
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
      RETURN -> State 223
      BREAK -> State 220
      CONTINUE -> State 221
      UNSAFE -> State 224
      STRUCT -> State 31
      FUNC -> State 21
      ASYNC -> State 219
      LABEL_DECL -> State 24
      RBRACE -> State 222
      LPAREN -> State 26
      NOT -> State 28
      SUB -> State 32
      MUL -> State 27
      BIT_NEG -> State 18
      BIT_AND -> State 17
      LEX_ERROR -> State 25
      expression -> State 226
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      statement -> State 230
      statement_body -> State 231
      unsafe_statement -> State 232
      jump_statement -> State 228
      jump_type -> State 229
      expressions -> State 227
      call_expr -> State 41
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 225
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

  State 165:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.LBRACE CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 233

  State 166:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 77

  State 167:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements.RPAREN
      package_statements: package_statements.package_statement
    Reduce:
      (nil)
    Goto:
      UNSAFE -> State 224
      RPAREN -> State 234
      unsafe_statement -> State 237
      package_statement_body -> State 236
      package_statement -> State 235

  State 168:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 238
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 169:
    Kernel Items:
      generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 170:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs.COMMA generic_parameter_def
      optional_generic_parameter_defs: generic_parameter_defs., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      COMMA -> State 239

  State 171:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 240

  State 172:
    Kernel Items:
      type_def: TYPE IDENTIFIER EQUAL value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      (nil)

  State 173:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., $
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      $ -> [type_def]
    Goto:
      IMPLEMENTS -> State 241

  State 174:
    Kernel Items:
      trait_algebra_type: trait_algebra_type.MUL traitable_type
      trait_algebra_type: trait_algebra_type.ADD traitable_type
      trait_algebra_type: trait_algebra_type.SUB traitable_type
      value_type: BIT_AND trait_algebra_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      ADD -> State 188
      SUB -> State 190
      MUL -> State 189

  State 175:
    Kernel Items:
      traitable_type: BIT_NEG atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 176:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 131
      field_def -> State 244
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      enum_value_def -> State 242
      implicit_enum_value_defs -> State 245
      implicit_enum_def -> State 122
      explicit_enum_value_defs -> State 243
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 177:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 246

  State 178:
    Kernel Items:
      func_type: FUNC LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 248
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      DOTDOTDOT -> State 247
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 252
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      parameter_decl -> State 250
      parameter_decls -> State 251
      optional_parameter_decls -> State 249
      func_type -> State 121

  State 179:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 180:
    Kernel Items:
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 181:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 253

  State 182:
    Kernel Items:
      implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 254

  State 183:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 256
      OR -> State 255

  State 184:
    Kernel Items:
      implicit_field_defs: implicit_field_defs.COMMA field_def
      optional_implicit_field_defs: implicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      COMMA -> State 257

  State 185:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 258

  State 186:
    Kernel Items:
      traitable_type: TILDE_TILDE atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 187:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties RPAREN., $
    Reduce:
      $ -> [trait_def]
    Goto:
      (nil)

  State 188:
    Kernel Items:
      trait_algebra_type: trait_algebra_type ADD.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      LPAREN -> State 114
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      atom_type -> State 117
      traitable_type -> State 259
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127

  State 189:
    Kernel Items:
      trait_algebra_type: trait_algebra_type MUL.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      LPAREN -> State 114
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      atom_type -> State 117
      traitable_type -> State 260
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127

  State 190:
    Kernel Items:
      trait_algebra_type: trait_algebra_type SUB.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      LPAREN -> State 114
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      atom_type -> State 117
      traitable_type -> State 261
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127

  State 191:
    Kernel Items:
      trait_properties: trait_properties COMMA.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 112
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 131
      field_def -> State 120
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_property -> State 262
      trait_def -> State 127
      func_type -> State 121
      method_signature -> State 124

  State 192:
    Kernel Items:
      trait_properties: trait_properties NEWLINES.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 112
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 131
      field_def -> State 120
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_property -> State 263
      trait_def -> State 127
      func_type -> State 121
      method_signature -> State 124

  State 193:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 264
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 194:
    Kernel Items:
      parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      (nil)

  State 195:
    Kernel Items:
      optional_receiver: LPAREN parameter_def RPAREN., IDENTIFIER
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      (nil)

  State 196:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 265

  State 197:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 267
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      return_type -> State 266
      func_type -> State 121

  State 198:
    Kernel Items:
      parameter_defs: parameter_defs COMMA.parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 132
      parameter_def -> State 268

  State 199:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 200:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 201:
    Kernel Items:
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      (nil)

  State 202:
    Kernel Items:
      colon_expressions: colon_expressions COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 203:
    Kernel Items:
      colon_expressions: optional_expression COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 204:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 131
      field_def -> State 269
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 205:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 131
      field_def -> State 270
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 206:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_def]
    Goto:
      (nil)

  State 207:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 271
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 208:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 209:
    Kernel Items:
      access_expr: access_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      optional_arguments: arguments., RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 139

  State 211:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 272

  State 212:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_def LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 213:
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
      sequence_expr -> State 273
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

  State 214:
    Kernel Items:
      loop_expr: FOR IN sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 274

  State 215:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 275

  State 216:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      loop_expr: FOR sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 96
      block_body -> State 276

  State 218:
    Kernel Items:
      if_expr: IF sequence_expr block_body., *
      if_expr: IF sequence_expr block_body.ELSE block_body
      if_expr: IF sequence_expr block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 277

  State 219:
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
      call_expr -> State 279
      atom_expr -> State 39
      literal -> State 44
      anonymous_struct_expr -> State 38
      access_expr -> State 278
      explicit_struct_def -> State 43
      anonymous_func_expr -> State 37

  State 220:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 221:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 222:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 223:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 224:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 280

  State 225:
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
      ADD_ASSIGN -> State 281
      SUB_ASSIGN -> State 292
      MUL_ASSIGN -> State 291
      DIV_ASSIGN -> State 289
      MOD_ASSIGN -> State 290
      ADD_ONE_ASSIGN -> State 282
      SUB_ONE_ASSIGN -> State 293
      BIT_NEG_ASSIGN -> State 285
      BIT_AND_ASSIGN -> State 283
      BIT_OR_ASSIGN -> State 286
      BIT_XOR_ASSIGN -> State 288
      BIT_LSHIFT_ASSIGN -> State 284
      BIT_RSHIFT_ASSIGN -> State 287
      unary_op_assign -> State 295
      binary_op_assign -> State 294
      optional_generic_binding -> State 71

  State 226:
    Kernel Items:
      expressions: expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 227:
    Kernel Items:
      statement_body: expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [statement_body]
    Goto:
      COMMA -> State 296

  State 228:
    Kernel Items:
      statement_body: jump_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 229:
    Kernel Items:
      jump_statement: jump_type.optional_jump_label optional_expressions
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 297
      optional_jump_label -> State 298

  State 230:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 231:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 299
      SEMICOLON -> State 300

  State 232:
    Kernel Items:
      statement_body: unsafe_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 233:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE.CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      CASE -> State 301

  State 234:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements RPAREN., $
    Reduce:
      $ -> [package_def]
    Goto:
      (nil)

  State 235:
    Kernel Items:
      package_statements: package_statements package_statement., *
    Reduce:
      * -> [package_statements]
    Goto:
      (nil)

  State 236:
    Kernel Items:
      package_statement: package_statement_body.NEWLINES
      package_statement: package_statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 302
      SEMICOLON -> State 303

  State 237:
    Kernel Items:
      package_statement_body: unsafe_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 238:
    Kernel Items:
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      (nil)

  State 239:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA.generic_parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 168
      generic_parameter_def -> State 304

  State 240:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 305
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 242:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 306
      OR -> State 307

  State 243:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 308

  State 244:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 254

  State 245:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 309
      OR -> State 310

  State 246:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 248
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      DOTDOTDOT -> State 247
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 252
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      parameter_decl -> State 250
      parameter_decls -> State 251
      optional_parameter_decls -> State 311
      func_type -> State 121

  State 247:
    Kernel Items:
      parameter_decl: DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 312
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 248:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      DOLLAR_LBRACKET -> State 67
      DOTDOTDOT -> State 313
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      optional_generic_binding -> State 179
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 314
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 249:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 315

  State 250:
    Kernel Items:
      parameter_decls: parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 251:
    Kernel Items:
      parameter_decls: parameter_decls.COMMA parameter_decl
      optional_parameter_decls: parameter_decls., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      COMMA -> State 316

  State 252:
    Kernel Items:
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 253:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 131
      field_def -> State 244
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      enum_value_def -> State 317
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 254:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 318

  State 255:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 131
      field_def -> State 244
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      enum_value_def -> State 319
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 256:
    Kernel Items:
      implicit_enum_def: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_def]
    Goto:
      (nil)

  State 257:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 131
      field_def -> State 320
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 258:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_def]
    Goto:
      (nil)

  State 259:
    Kernel Items:
      trait_algebra_type: trait_algebra_type ADD traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 260:
    Kernel Items:
      trait_algebra_type: trait_algebra_type MUL traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 261:
    Kernel Items:
      trait_algebra_type: trait_algebra_type SUB traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 262:
    Kernel Items:
      trait_properties: trait_properties COMMA trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 263:
    Kernel Items:
      trait_properties: trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 264:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      (nil)

  State 265:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 132
      parameter_def -> State 136
      parameter_defs -> State 137
      optional_parameter_defs -> State 321

  State 266:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 96
      block_body -> State 322

  State 267:
    Kernel Items:
      return_type: value_type., $
    Reduce:
      $ -> [return_type]
    Goto:
      (nil)

  State 268:
    Kernel Items:
      parameter_defs: parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 269:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 270:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 271:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
    Reduce:
      * -> [generic_arguments]
    Goto:
      (nil)

  State 272:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 273:
    Kernel Items:
      loop_expr: DO block_body FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 274:
    Kernel Items:
      loop_expr: FOR IN sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 96
      block_body -> State 323

  State 275:
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
      optional_sequence_expr -> State 324
      sequence_expr -> State 216
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

  State 276:
    Kernel Items:
      loop_expr: FOR sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 277:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE.block_body
      if_expr: IF sequence_expr block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 95
      LBRACE -> State 96
      if_expr -> State 326
      block_body -> State 325

  State 278:
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

  State 279:
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

  State 280:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 327

  State 281:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 282:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 283:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 284:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 285:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 286:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 287:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 288:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 289:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 290:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 291:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 292:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 293:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 294:
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
      expression -> State 328
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

  State 295:
    Kernel Items:
      statement_body: access_expr unary_op_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 296:
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
      expression -> State 329
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

  State 297:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 298:
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
      expression -> State 226
      optional_label_decl -> State 46
      sequence_expr -> State 51
      block_expr -> State 40
      expressions -> State 330
      optional_expressions -> State 331
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
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 300:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 301:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE CASE.DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 332

  State 302:
    Kernel Items:
      package_statement: package_statement_body NEWLINES., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 303:
    Kernel Items:
      package_statement: package_statement_body SEMICOLON., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 304:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 305:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      (nil)

  State 306:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 131
      field_def -> State 244
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      enum_value_def -> State 333
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 307:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 131
      field_def -> State 244
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      enum_value_def -> State 334
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 308:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_def]
    Goto:
      (nil)

  State 309:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 131
      field_def -> State 244
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      enum_value_def -> State 335
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 310:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 131
      field_def -> State 244
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      enum_value_def -> State 336
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 311:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 337

  State 312:
    Kernel Items:
      parameter_decl: DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 313:
    Kernel Items:
      parameter_decl: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 338
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      func_type -> State 121

  State 314:
    Kernel Items:
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 315:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 267
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      return_type -> State 339
      func_type -> State 121

  State 316:
    Kernel Items:
      parameter_decls: parameter_decls COMMA.parameter_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 248
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      DOTDOTDOT -> State 247
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 252
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      parameter_decl -> State 340
      func_type -> State 121

  State 317:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 318:
    Kernel Items:
      enum_value_def: field_def ASSIGN DEFAULT., *
    Reduce:
      * -> [enum_value_def]
    Goto:
      (nil)

  State 319:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 320:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [implicit_field_defs]
    Goto:
      (nil)

  State 321:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 341

  State 322:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 323:
    Kernel Items:
      loop_expr: FOR IN sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 324:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 342

  State 325:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 326:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 327:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 343

  State 328:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 329:
    Kernel Items:
      expressions: expressions COMMA expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 330:
    Kernel Items:
      expressions: expressions.COMMA expression
      optional_expressions: expressions., *
    Reduce:
      * -> [optional_expressions]
    Goto:
      COMMA -> State 296

  State 331:
    Kernel Items:
      jump_statement: jump_type optional_jump_label optional_expressions., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE CASE DEFAULT.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 344

  State 333:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 334:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
      explicit_enum_value_defs: enum_value_def OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 335:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 336:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 337:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 267
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      return_type -> State 345
      func_type -> State 121

  State 338:
    Kernel Items:
      parameter_decl: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 339:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN return_type., $
    Reduce:
      $ -> [func_type]
    Goto:
      (nil)

  State 340:
    Kernel Items:
      parameter_decls: parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 147
      STRUCT -> State 31
      ENUM -> State 111
      TRAIT -> State 15
      FUNC -> State 143
      LPAREN -> State 114
      QUESTION -> State 115
      TILDE_TILDE -> State 116
      BIT_NEG -> State 110
      BIT_AND -> State 109
      atom_type -> State 117
      traitable_type -> State 130
      trait_algebra_type -> State 126
      value_type -> State 267
      implicit_struct_def -> State 123
      explicit_struct_def -> State 119
      implicit_enum_def -> State 122
      explicit_enum_def -> State 118
      trait_def -> State 127
      return_type -> State 346
      func_type -> State 121

  State 342:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 96
      block_body -> State 347

  State 343:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 348

  State 344:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE CASE DEFAULT RBRACE., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 345:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 346:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 96
      block_body -> State 349

  State 347:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 348:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 349:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

Number of states: 349
Number of shift actions: 2005
Number of reduce actions: 280
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
