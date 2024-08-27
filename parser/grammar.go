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
	// 45:10: source -> ...
	ToSource(OptionalDefinitions_ *GenericSymbol) (*GenericSymbol, error)

	// 48:2: optional_definitions -> NEWLINES: ...
	NewlinesToOptionalDefinitions(Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 49:2: optional_definitions -> definitions: ...
	DefinitionsToOptionalDefinitions(OptionalNewlines_ *GenericSymbol, Definitions_ *GenericSymbol, OptionalNewlines_2 *GenericSymbol) (*GenericSymbol, error)

	// 50:2: optional_definitions -> nil: ...
	NilToOptionalDefinitions() (*GenericSymbol, error)

	// 53:2: optional_newlines -> NEWLINES: ...
	NewlinesToOptionalNewlines(Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 54:2: optional_newlines -> nil: ...
	NilToOptionalNewlines() (*GenericSymbol, error)

	// 57:2: definitions -> nil: ...
	NilToDefinitions(Definition_ *GenericSymbol) (*GenericSymbol, error)

	// 58:2: definitions -> add: ...
	AddToDefinitions(Definitions_ *GenericSymbol, Newlines_ *GenericSymbol, Definition_ *GenericSymbol) (*GenericSymbol, error)

	// 61:2: definition -> package_def: ...
	PackageDefToDefinition(PackageDef_ *GenericSymbol) (*GenericSymbol, error)

	// 62:2: definition -> type_def: ...
	TypeDefToDefinition(TypeDef_ *GenericSymbol) (*GenericSymbol, error)

	// 63:2: definition -> named_func_def: ...
	NamedFuncDefToDefinition(NamedFuncDef_ *GenericSymbol) (*GenericSymbol, error)

	// 64:2: definition -> unsafe_statement: ...
	UnsafeStatementToDefinition(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 71:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 72:2: expression -> switch_expr: ...
	SwitchExprToExpression(OptionalLabelDecl_ *GenericSymbol, SwitchExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 73:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ *GenericSymbol, LoopExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 74:2: expression -> sequence_expr: ...
	SequenceExprToExpression(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 77:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 78:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (*GenericSymbol, error)

	// 87:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 88:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, BlockBody_2 *GenericSymbol) (*GenericSymbol, error)

	// 89:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 96:15: switch_expr -> ...
	ToSwitchExpr(Switch_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Lbrace_ *GenericSymbol, Case_ *GenericSymbol, Default_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 110:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 111:2: loop_expr -> do_while: ...
	DoWhileToLoopExpr(Do_ *GenericSymbol, BlockBody_ *GenericSymbol, For_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 112:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 113:2: loop_expr -> iterator: ...
	IteratorToLoopExpr(For_ *GenericSymbol, In_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 114:2: loop_expr -> for: ...
	ForToLoopExpr(For_ *GenericSymbol, Semicolon_ *GenericSymbol, OptionalSequenceExpr_ *GenericSymbol, Semicolon_2 *GenericSymbol, OptionalSequenceExpr_2 *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 117:2: optional_sequence_expr -> sequence_expr: ...
	SequenceExprToOptionalSequenceExpr(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 118:2: optional_sequence_expr -> nil: ...
	NilToOptionalSequenceExpr() (*GenericSymbol, error)

	// 128:17: sequence_expr -> ...
	ToSequenceExpr(OrExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 130:14: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 132:14: block_body -> ...
	ToBlockBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 135:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 136:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 139:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 140:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 143:2: statement_body -> unsafe_statement: ...
	UnsafeStatementToStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 145:2: statement_body -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToStatementBody(Expressions_ *GenericSymbol) (*GenericSymbol, error)

	// 148:2: statement_body -> async: ...
	AsyncToStatementBody(Async_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 157:2: statement_body -> defer: ...
	DeferToStatementBody(Defer_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 161:2: statement_body -> jump_statement: ...
	JumpStatementToStatementBody(JumpStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 173:2: statement_body -> unary_op_assign_statement: ...
	UnaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, UnaryOpAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 174:2: statement_body -> binary_op_assign_statement: ...
	BinaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 181:2: unary_op_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToUnaryOpAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 182:2: unary_op_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToUnaryOpAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 185:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 186:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 187:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 188:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 189:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 190:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 191:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 192:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 193:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 194:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 195:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 203:20: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ *GenericSymbol, Less_ *GenericSymbol, Identifier_ *GenericSymbol, Greater_ *GenericSymbol, StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 210:2: jump_statement -> ...
	ToJumpStatement(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 213:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 214:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 215:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 218:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 219:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 222:2: expressions -> expression: ...
	ExpressionToExpressions(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 223:2: expressions -> add: ...
	AddToExpressions(Expressions_ *GenericSymbol, Comma_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 226:2: optional_expressions -> expressions: ...
	ExpressionsToOptionalExpressions(Expressions_ *GenericSymbol) (*GenericSymbol, error)

	// 227:2: optional_expressions -> nil: ...
	NilToOptionalExpressions() (*GenericSymbol, error)

	// 233:13: call_expr -> ...
	ToCallExpr(AccessExpr_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalArguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 236:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ *GenericSymbol, OptionalGenericArguments_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 237:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (*GenericSymbol, error)

	// 240:2: optional_generic_arguments -> generic_arguments: ...
	GenericArgumentsToOptionalGenericArguments(GenericArguments_ *GenericSymbol) (*GenericSymbol, error)

	// 241:2: optional_generic_arguments -> nil: ...
	NilToOptionalGenericArguments() (*GenericSymbol, error)

	// 245:2: generic_arguments -> value_type: ...
	ValueTypeToGenericArguments(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 246:2: generic_arguments -> add: ...
	AddToGenericArguments(GenericArguments_ *GenericSymbol, Comma_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 249:2: optional_arguments -> arguments: ...
	ArgumentsToOptionalArguments(Arguments_ *GenericSymbol) (*GenericSymbol, error)

	// 250:2: optional_arguments -> nil: ...
	NilToOptionalArguments() (*GenericSymbol, error)

	// 253:2: arguments -> argument: ...
	ArgumentToArguments(Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 254:2: arguments -> add: ...
	AddToArguments(Arguments_ *GenericSymbol, Comma_ *GenericSymbol, Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 257:2: argument -> positional: ...
	PositionalToArgument(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 258:2: argument -> named: ...
	NamedToArgument(Identifier_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 259:2: argument -> colon_expressions: ...
	ColonExpressionsToArgument(ColonExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 262:2: colon_expressions -> pair: ...
	PairToColonExpressions(OptionalExpression_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_2 *GenericSymbol) (*GenericSymbol, error)

	// 263:2: colon_expressions -> add: ...
	AddToColonExpressions(ColonExpressions_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_ *GenericSymbol) (*GenericSymbol, error)

	// 266:2: optional_expression -> expression: ...
	ExpressionToOptionalExpression(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 267:2: optional_expression -> nil: ...
	NilToOptionalExpression() (*GenericSymbol, error)

	// 277:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 278:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 279:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 280:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 281:2: atom_expr -> anonymous_struct_expr: ...
	AnonymousStructExprToAtomExpr(AnonymousStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 282:2: atom_expr -> LEX_ERROR: ...
	LexErrorToAtomExpr(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 285:2: literal -> TRUE: ...
	TrueToLiteral(True_ *GenericSymbol) (*GenericSymbol, error)

	// 286:2: literal -> FALSE: ...
	FalseToLiteral(False_ *GenericSymbol) (*GenericSymbol, error)

	// 287:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 288:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 289:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 290:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 293:2: anonymous_struct_expr -> explicit: ...
	ExplicitToAnonymousStructExpr(ExplicitStructDef_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 294:2: anonymous_struct_expr -> implicit: ...
	ImplicitToAnonymousStructExpr(Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 297:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 298:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 299:2: access_expr -> call_expr: ...
	CallExprToAccessExpr(CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 300:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Argument_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 303:2: postfix_unary_expr -> access_expr: ...
	AccessExprToPostfixUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 304:2: postfix_unary_expr -> question: ...
	QuestionToPostfixUnaryExpr(AccessExpr_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 307:2: prefix_unary_op -> NOT: ...
	NotToPrefixUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 308:2: prefix_unary_op -> BIT_NEG: ...
	BitNegToPrefixUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 309:2: prefix_unary_op -> SUB: ...
	SubToPrefixUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 312:2: prefix_unary_op -> MUL: ...
	MulToPrefixUnaryOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 315:2: prefix_unary_op -> BIT_AND: ...
	BitAndToPrefixUnaryOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 318:2: prefix_unary_expr -> postfix_unary_expr: ...
	PostfixUnaryExprToPrefixUnaryExpr(PostfixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 319:2: prefix_unary_expr -> prefix_op: ...
	PrefixOpToPrefixUnaryExpr(PrefixUnaryOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 322:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 323:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 324:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 325:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 326:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 327:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 330:2: mul_expr -> prefix_unary_expr: ...
	PrefixUnaryExprToMulExpr(PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 331:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 334:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 335:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 336:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 337:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 340:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 341:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 344:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 345:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 346:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 347:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 348:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 349:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 352:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 353:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 356:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 357:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 360:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 361:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 368:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 369:2: atom_type -> explicit_struct_def: ...
	ExplicitStructDefToAtomType(ExplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 370:2: atom_type -> implicit_struct_def: ...
	ImplicitStructDefToAtomType(ImplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 371:2: atom_type -> explicit_enum_def: ...
	ExplicitEnumDefToAtomType(ExplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 372:2: atom_type -> implicit_enum_def: ...
	ImplicitEnumDefToAtomType(ImplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 373:2: atom_type -> trait_def: ...
	TraitDefToAtomType(TraitDef_ *GenericSymbol) (*GenericSymbol, error)

	// 374:2: atom_type -> LEX_ERROR: ...
	LexErrorToAtomType(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 377:2: traitable_type -> atom_type: ...
	AtomTypeToTraitableType(AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 378:2: traitable_type -> public_methods_trait: ...
	PublicMethodsTraitToTraitableType(BitNeg_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 379:2: traitable_type -> public_trait: ...
	PublicTraitToTraitableType(TildeTilde_ *GenericSymbol, AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 383:2: trait_algebra_type -> traitable_type: ...
	TraitableTypeToTraitAlgebraType(TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 384:2: trait_algebra_type -> intersect: ...
	IntersectToTraitAlgebraType(TraitAlgebraType_ *GenericSymbol, Mul_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 385:2: trait_algebra_type -> union: ...
	UnionToTraitAlgebraType(TraitAlgebraType_ *GenericSymbol, Add_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 386:2: trait_algebra_type -> difference: ...
	DifferenceToTraitAlgebraType(TraitAlgebraType_ *GenericSymbol, Sub_ *GenericSymbol, TraitableType_ *GenericSymbol) (*GenericSymbol, error)

	// 389:2: value_type -> inferred: ...
	InferredToValueType(Question_ *GenericSymbol) (*GenericSymbol, error)

	// 390:2: value_type -> trait_algebra_type: ...
	TraitAlgebraTypeToValueType(TraitAlgebraType_ *GenericSymbol) (*GenericSymbol, error)

	// 391:2: value_type -> reference: ...
	ReferenceToValueType(BitAnd_ *GenericSymbol, TraitAlgebraType_ *GenericSymbol) (*GenericSymbol, error)

	// 392:2: value_type -> func_type: ...
	FuncTypeToValueType(FuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 395:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 396:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol, Implements_ *GenericSymbol, ValueType_2 *GenericSymbol) (*GenericSymbol, error)

	// 397:2: type_def -> alias: ...
	AliasToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, Equal_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 405:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 406:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 409:2: generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToGenericParameterDefs(GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 410:2: generic_parameter_defs -> add: ...
	AddToGenericParameterDefs(GenericParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 413:2: optional_generic_parameter_defs -> generic_parameter_defs: ...
	GenericParameterDefsToOptionalGenericParameterDefs(GenericParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 414:2: optional_generic_parameter_defs -> nil: ...
	NilToOptionalGenericParameterDefs() (*GenericSymbol, error)

	// 417:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ *GenericSymbol, OptionalGenericParameterDefs_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 418:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (*GenericSymbol, error)

	// 425:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 426:2: field_def -> implicit: ...
	ImplicitToFieldDef(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 427:2: field_def -> unsafe_statement: ...
	UnsafeStatementToFieldDef(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 430:2: implicit_field_defs -> field_def: ...
	FieldDefToImplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 431:2: implicit_field_defs -> add: ...
	AddToImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 434:2: optional_implicit_field_defs -> implicit_field_defs: ...
	ImplicitFieldDefsToOptionalImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 435:2: optional_implicit_field_defs -> nil: ...
	NilToOptionalImplicitFieldDefs() (*GenericSymbol, error)

	// 437:23: implicit_struct_def -> ...
	ToImplicitStructDef(Lparen_ *GenericSymbol, OptionalImplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 440:2: explicit_field_defs -> field_def: ...
	FieldDefToExplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 441:2: explicit_field_defs -> implicit: ...
	ImplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Newlines_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 442:2: explicit_field_defs -> explicit: ...
	ExplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 445:2: optional_explicit_field_defs -> explicit_field_defs: ...
	ExplicitFieldDefsToOptionalExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 446:2: optional_explicit_field_defs -> nil: ...
	NilToOptionalExplicitFieldDefs() (*GenericSymbol, error)

	// 448:23: explicit_struct_def -> ...
	ToExplicitStructDef(Struct_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalExplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 458:2: enum_value_def -> field_def: ...
	FieldDefToEnumValueDef(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 459:2: enum_value_def -> default: ...
	DefaultToEnumValueDef(FieldDef_ *GenericSymbol, Assign_ *GenericSymbol, Default_ *GenericSymbol) (*GenericSymbol, error)

	// 462:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 463:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 465:21: implicit_enum_def -> ...
	ToImplicitEnumDef(Lparen_ *GenericSymbol, ImplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 468:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 469:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 470:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 471:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 473:21: explicit_enum_def -> ...
	ToExplicitEnumDef(Enum_ *GenericSymbol, Lparen_ *GenericSymbol, ExplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 480:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 481:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ *GenericSymbol) (*GenericSymbol, error)

	// 484:2: trait_properties -> trait_property: ...
	TraitPropertyToTraitProperties(TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 485:2: trait_properties -> implicit: ...
	ImplicitToTraitProperties(TraitProperties_ *GenericSymbol, Newlines_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 486:2: trait_properties -> explicit: ...
	ExplicitToTraitProperties(TraitProperties_ *GenericSymbol, Comma_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 489:2: optional_trait_properties -> trait_properties: ...
	TraitPropertiesToOptionalTraitProperties(TraitProperties_ *GenericSymbol) (*GenericSymbol, error)

	// 490:2: optional_trait_properties -> nil: ...
	NilToOptionalTraitProperties() (*GenericSymbol, error)

	// 492:13: trait_def -> ...
	ToTraitDef(Trait_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalTraitProperties_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 500:2: return_type -> value_type: ...
	ValueTypeToReturnType(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 501:2: return_type -> nil: ...
	NilToReturnType() (*GenericSymbol, error)

	// 504:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 505:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 506:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 507:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 510:2: parameter_decls -> parameter_decl: ...
	ParameterDeclToParameterDecls(ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 511:2: parameter_decls -> add: ...
	AddToParameterDecls(ParameterDecls_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 514:2: optional_parameter_decls -> parameter_decls: ...
	ParameterDeclsToOptionalParameterDecls(ParameterDecls_ *GenericSymbol) (*GenericSymbol, error)

	// 515:2: optional_parameter_decls -> nil: ...
	NilToOptionalParameterDecls() (*GenericSymbol, error)

	// 517:13: func_type -> ...
	ToFuncType(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 525:20: method_signature -> ...
	ToMethodSignature(Func_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 528:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 529:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 532:2: parameter_defs -> parameter_def: ...
	ParameterDefToParameterDefs(ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 533:2: parameter_defs -> add: ...
	AddToParameterDefs(ParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 536:2: optional_parameter_defs -> parameter_defs: ...
	ParameterDefsToOptionalParameterDefs(ParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 537:2: optional_parameter_defs -> nil: ...
	NilToOptionalParameterDefs() (*GenericSymbol, error)

	// 540:2: optional_receiver -> receiver: ...
	ReceiverToOptionalReceiver(Lparen_ *GenericSymbol, ParameterDef_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 541:2: optional_receiver -> nil: ...
	NilToOptionalReceiver() (*GenericSymbol, error)

	// 544:2: named_func_def -> ...
	ToNamedFuncDef(Func_ *GenericSymbol, OptionalReceiver_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 547:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 554:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 555:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, PackageStatements_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 557:26: package_statement_body -> ...
	ToPackageStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 560:2: package_statement -> implicit: ...
	ImplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 561:2: package_statement -> explicit: ...
	ExplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 564:2: package_statements -> empty_list: ...
	EmptyListToPackageStatements() (*GenericSymbol, error)

	// 565:2: package_statements -> add: ...
	AddToPackageStatements(PackageStatements_ *GenericSymbol, PackageStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 569:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 570:2: lex_internal_tokens -> COMMENT: ...
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

func ParseSource(lexer Lexer, reducer Reducer) (*GenericSymbol, error) {

	return ParseSourceWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseSourceWithCustomErrorHandler(
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

	item, err := _Parse(lexer, reducer, errHandler, _State2)
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
	case SourceType:
		return "source"
	case OptionalDefinitionsType:
		return "optional_definitions"
	case OptionalNewlinesType:
		return "optional_newlines"
	case DefinitionsType:
		return "definitions"
	case DefinitionType:
		return "definition"
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

	SourceType                       = SymbolId(338)
	OptionalDefinitionsType          = SymbolId(339)
	OptionalNewlinesType             = SymbolId(340)
	DefinitionsType                  = SymbolId(341)
	DefinitionType                   = SymbolId(342)
	ExpressionType                   = SymbolId(343)
	OptionalLabelDeclType            = SymbolId(344)
	IfExprType                       = SymbolId(345)
	SwitchExprType                   = SymbolId(346)
	LoopExprType                     = SymbolId(347)
	OptionalSequenceExprType         = SymbolId(348)
	SequenceExprType                 = SymbolId(349)
	BlockExprType                    = SymbolId(350)
	BlockBodyType                    = SymbolId(351)
	StatementsType                   = SymbolId(352)
	StatementType                    = SymbolId(353)
	StatementBodyType                = SymbolId(354)
	UnaryOpAssignType                = SymbolId(355)
	BinaryOpAssignType               = SymbolId(356)
	UnsafeStatementType              = SymbolId(357)
	JumpStatementType                = SymbolId(358)
	JumpTypeType                     = SymbolId(359)
	OptionalJumpLabelType            = SymbolId(360)
	ExpressionsType                  = SymbolId(361)
	OptionalExpressionsType          = SymbolId(362)
	CallExprType                     = SymbolId(363)
	OptionalGenericBindingType       = SymbolId(364)
	OptionalGenericArgumentsType     = SymbolId(365)
	GenericArgumentsType             = SymbolId(366)
	OptionalArgumentsType            = SymbolId(367)
	ArgumentsType                    = SymbolId(368)
	ArgumentType                     = SymbolId(369)
	ColonExpressionsType             = SymbolId(370)
	OptionalExpressionType           = SymbolId(371)
	AtomExprType                     = SymbolId(372)
	LiteralType                      = SymbolId(373)
	AnonymousStructExprType          = SymbolId(374)
	AccessExprType                   = SymbolId(375)
	PostfixUnaryExprType             = SymbolId(376)
	PrefixUnaryOpType                = SymbolId(377)
	PrefixUnaryExprType              = SymbolId(378)
	MulOpType                        = SymbolId(379)
	MulExprType                      = SymbolId(380)
	AddOpType                        = SymbolId(381)
	AddExprType                      = SymbolId(382)
	CmpOpType                        = SymbolId(383)
	CmpExprType                      = SymbolId(384)
	AndExprType                      = SymbolId(385)
	OrExprType                       = SymbolId(386)
	AtomTypeType                     = SymbolId(387)
	TraitableTypeType                = SymbolId(388)
	TraitAlgebraTypeType             = SymbolId(389)
	ValueTypeType                    = SymbolId(390)
	TypeDefType                      = SymbolId(391)
	GenericParameterDefType          = SymbolId(392)
	GenericParameterDefsType         = SymbolId(393)
	OptionalGenericParameterDefsType = SymbolId(394)
	OptionalGenericParametersType    = SymbolId(395)
	FieldDefType                     = SymbolId(396)
	ImplicitFieldDefsType            = SymbolId(397)
	OptionalImplicitFieldDefsType    = SymbolId(398)
	ImplicitStructDefType            = SymbolId(399)
	ExplicitFieldDefsType            = SymbolId(400)
	OptionalExplicitFieldDefsType    = SymbolId(401)
	ExplicitStructDefType            = SymbolId(402)
	EnumValueDefType                 = SymbolId(403)
	ImplicitEnumValueDefsType        = SymbolId(404)
	ImplicitEnumDefType              = SymbolId(405)
	ExplicitEnumValueDefsType        = SymbolId(406)
	ExplicitEnumDefType              = SymbolId(407)
	TraitPropertyType                = SymbolId(408)
	TraitPropertiesType              = SymbolId(409)
	OptionalTraitPropertiesType      = SymbolId(410)
	TraitDefType                     = SymbolId(411)
	ReturnTypeType                   = SymbolId(412)
	ParameterDeclType                = SymbolId(413)
	ParameterDeclsType               = SymbolId(414)
	OptionalParameterDeclsType       = SymbolId(415)
	FuncTypeType                     = SymbolId(416)
	MethodSignatureType              = SymbolId(417)
	ParameterDefType                 = SymbolId(418)
	ParameterDefsType                = SymbolId(419)
	OptionalParameterDefsType        = SymbolId(420)
	OptionalReceiverType             = SymbolId(421)
	NamedFuncDefType                 = SymbolId(422)
	AnonymousFuncExprType            = SymbolId(423)
	PackageDefType                   = SymbolId(424)
	PackageStatementBodyType         = SymbolId(425)
	PackageStatementType             = SymbolId(426)
	PackageStatementsType            = SymbolId(427)
	LexInternalTokensType            = SymbolId(428)
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
	_ReduceToSource                                           = _ReduceType(1)
	_ReduceNewlinesToOptionalDefinitions                      = _ReduceType(2)
	_ReduceDefinitionsToOptionalDefinitions                   = _ReduceType(3)
	_ReduceNilToOptionalDefinitions                           = _ReduceType(4)
	_ReduceNewlinesToOptionalNewlines                         = _ReduceType(5)
	_ReduceNilToOptionalNewlines                              = _ReduceType(6)
	_ReduceNilToDefinitions                                   = _ReduceType(7)
	_ReduceAddToDefinitions                                   = _ReduceType(8)
	_ReducePackageDefToDefinition                             = _ReduceType(9)
	_ReduceTypeDefToDefinition                                = _ReduceType(10)
	_ReduceNamedFuncDefToDefinition                           = _ReduceType(11)
	_ReduceUnsafeStatementToDefinition                        = _ReduceType(12)
	_ReduceIfExprToExpression                                 = _ReduceType(13)
	_ReduceSwitchExprToExpression                             = _ReduceType(14)
	_ReduceLoopExprToExpression                               = _ReduceType(15)
	_ReduceSequenceExprToExpression                           = _ReduceType(16)
	_ReduceLabelDeclToOptionalLabelDecl                       = _ReduceType(17)
	_ReduceUnlabelledToOptionalLabelDecl                      = _ReduceType(18)
	_ReduceNoElseToIfExpr                                     = _ReduceType(19)
	_ReduceIfElseToIfExpr                                     = _ReduceType(20)
	_ReduceMultiIfElseToIfExpr                                = _ReduceType(21)
	_ReduceToSwitchExpr                                       = _ReduceType(22)
	_ReduceInfiniteToLoopExpr                                 = _ReduceType(23)
	_ReduceDoWhileToLoopExpr                                  = _ReduceType(24)
	_ReduceWhileToLoopExpr                                    = _ReduceType(25)
	_ReduceIteratorToLoopExpr                                 = _ReduceType(26)
	_ReduceForToLoopExpr                                      = _ReduceType(27)
	_ReduceSequenceExprToOptionalSequenceExpr                 = _ReduceType(28)
	_ReduceNilToOptionalSequenceExpr                          = _ReduceType(29)
	_ReduceToSequenceExpr                                     = _ReduceType(30)
	_ReduceToBlockExpr                                        = _ReduceType(31)
	_ReduceToBlockBody                                        = _ReduceType(32)
	_ReduceEmptyListToStatements                              = _ReduceType(33)
	_ReduceAddToStatements                                    = _ReduceType(34)
	_ReduceImplicitToStatement                                = _ReduceType(35)
	_ReduceExplicitToStatement                                = _ReduceType(36)
	_ReduceUnsafeStatementToStatementBody                     = _ReduceType(37)
	_ReduceExpressionOrImplicitStructToStatementBody          = _ReduceType(38)
	_ReduceAsyncToStatementBody                               = _ReduceType(39)
	_ReduceDeferToStatementBody                               = _ReduceType(40)
	_ReduceJumpStatementToStatementBody                       = _ReduceType(41)
	_ReduceUnaryOpAssignStatementToStatementBody              = _ReduceType(42)
	_ReduceBinaryOpAssignStatementToStatementBody             = _ReduceType(43)
	_ReduceAddOneAssignToUnaryOpAssign                        = _ReduceType(44)
	_ReduceSubOneAssignToUnaryOpAssign                        = _ReduceType(45)
	_ReduceAddAssignToBinaryOpAssign                          = _ReduceType(46)
	_ReduceSubAssignToBinaryOpAssign                          = _ReduceType(47)
	_ReduceMulAssignToBinaryOpAssign                          = _ReduceType(48)
	_ReduceDivAssignToBinaryOpAssign                          = _ReduceType(49)
	_ReduceModAssignToBinaryOpAssign                          = _ReduceType(50)
	_ReduceBitNegAssignToBinaryOpAssign                       = _ReduceType(51)
	_ReduceBitAndAssignToBinaryOpAssign                       = _ReduceType(52)
	_ReduceBitOrAssignToBinaryOpAssign                        = _ReduceType(53)
	_ReduceBitXorAssignToBinaryOpAssign                       = _ReduceType(54)
	_ReduceBitLshiftAssignToBinaryOpAssign                    = _ReduceType(55)
	_ReduceBitRshiftAssignToBinaryOpAssign                    = _ReduceType(56)
	_ReduceToUnsafeStatement                                  = _ReduceType(57)
	_ReduceToJumpStatement                                    = _ReduceType(58)
	_ReduceReturnToJumpType                                   = _ReduceType(59)
	_ReduceBreakToJumpType                                    = _ReduceType(60)
	_ReduceContinueToJumpType                                 = _ReduceType(61)
	_ReduceJumpLabelToOptionalJumpLabel                       = _ReduceType(62)
	_ReduceUnlabelledToOptionalJumpLabel                      = _ReduceType(63)
	_ReduceExpressionToExpressions                            = _ReduceType(64)
	_ReduceAddToExpressions                                   = _ReduceType(65)
	_ReduceExpressionsToOptionalExpressions                   = _ReduceType(66)
	_ReduceNilToOptionalExpressions                           = _ReduceType(67)
	_ReduceToCallExpr                                         = _ReduceType(68)
	_ReduceBindingToOptionalGenericBinding                    = _ReduceType(69)
	_ReduceNilToOptionalGenericBinding                        = _ReduceType(70)
	_ReduceGenericArgumentsToOptionalGenericArguments         = _ReduceType(71)
	_ReduceNilToOptionalGenericArguments                      = _ReduceType(72)
	_ReduceValueTypeToGenericArguments                        = _ReduceType(73)
	_ReduceAddToGenericArguments                              = _ReduceType(74)
	_ReduceArgumentsToOptionalArguments                       = _ReduceType(75)
	_ReduceNilToOptionalArguments                             = _ReduceType(76)
	_ReduceArgumentToArguments                                = _ReduceType(77)
	_ReduceAddToArguments                                     = _ReduceType(78)
	_ReducePositionalToArgument                               = _ReduceType(79)
	_ReduceNamedToArgument                                    = _ReduceType(80)
	_ReduceColonExpressionsToArgument                         = _ReduceType(81)
	_ReducePairToColonExpressions                             = _ReduceType(82)
	_ReduceAddToColonExpressions                              = _ReduceType(83)
	_ReduceExpressionToOptionalExpression                     = _ReduceType(84)
	_ReduceNilToOptionalExpression                            = _ReduceType(85)
	_ReduceLiteralToAtomExpr                                  = _ReduceType(86)
	_ReduceIdentifierToAtomExpr                               = _ReduceType(87)
	_ReduceBlockExprToAtomExpr                                = _ReduceType(88)
	_ReduceAnonymousFuncExprToAtomExpr                        = _ReduceType(89)
	_ReduceAnonymousStructExprToAtomExpr                      = _ReduceType(90)
	_ReduceLexErrorToAtomExpr                                 = _ReduceType(91)
	_ReduceTrueToLiteral                                      = _ReduceType(92)
	_ReduceFalseToLiteral                                     = _ReduceType(93)
	_ReduceIntegerLiteralToLiteral                            = _ReduceType(94)
	_ReduceFloatLiteralToLiteral                              = _ReduceType(95)
	_ReduceRuneLiteralToLiteral                               = _ReduceType(96)
	_ReduceStringLiteralToLiteral                             = _ReduceType(97)
	_ReduceExplicitToAnonymousStructExpr                      = _ReduceType(98)
	_ReduceImplicitToAnonymousStructExpr                      = _ReduceType(99)
	_ReduceAtomExprToAccessExpr                               = _ReduceType(100)
	_ReduceAccessToAccessExpr                                 = _ReduceType(101)
	_ReduceCallExprToAccessExpr                               = _ReduceType(102)
	_ReduceIndexToAccessExpr                                  = _ReduceType(103)
	_ReduceAccessExprToPostfixUnaryExpr                       = _ReduceType(104)
	_ReduceQuestionToPostfixUnaryExpr                         = _ReduceType(105)
	_ReduceNotToPrefixUnaryOp                                 = _ReduceType(106)
	_ReduceBitNegToPrefixUnaryOp                              = _ReduceType(107)
	_ReduceSubToPrefixUnaryOp                                 = _ReduceType(108)
	_ReduceMulToPrefixUnaryOp                                 = _ReduceType(109)
	_ReduceBitAndToPrefixUnaryOp                              = _ReduceType(110)
	_ReducePostfixUnaryExprToPrefixUnaryExpr                  = _ReduceType(111)
	_ReducePrefixOpToPrefixUnaryExpr                          = _ReduceType(112)
	_ReduceMulToMulOp                                         = _ReduceType(113)
	_ReduceDivToMulOp                                         = _ReduceType(114)
	_ReduceModToMulOp                                         = _ReduceType(115)
	_ReduceBitAndToMulOp                                      = _ReduceType(116)
	_ReduceBitLshiftToMulOp                                   = _ReduceType(117)
	_ReduceBitRshiftToMulOp                                   = _ReduceType(118)
	_ReducePrefixUnaryExprToMulExpr                           = _ReduceType(119)
	_ReduceOpToMulExpr                                        = _ReduceType(120)
	_ReduceAddToAddOp                                         = _ReduceType(121)
	_ReduceSubToAddOp                                         = _ReduceType(122)
	_ReduceBitOrToAddOp                                       = _ReduceType(123)
	_ReduceBitXorToAddOp                                      = _ReduceType(124)
	_ReduceMulExprToAddExpr                                   = _ReduceType(125)
	_ReduceOpToAddExpr                                        = _ReduceType(126)
	_ReduceEqualToCmpOp                                       = _ReduceType(127)
	_ReduceNotEqualToCmpOp                                    = _ReduceType(128)
	_ReduceLessToCmpOp                                        = _ReduceType(129)
	_ReduceLessOrEqualToCmpOp                                 = _ReduceType(130)
	_ReduceGreaterToCmpOp                                     = _ReduceType(131)
	_ReduceGreaterOrEqualToCmpOp                              = _ReduceType(132)
	_ReduceAddExprToCmpExpr                                   = _ReduceType(133)
	_ReduceOpToCmpExpr                                        = _ReduceType(134)
	_ReduceCmpExprToAndExpr                                   = _ReduceType(135)
	_ReduceOpToAndExpr                                        = _ReduceType(136)
	_ReduceAndExprToOrExpr                                    = _ReduceType(137)
	_ReduceOpToOrExpr                                         = _ReduceType(138)
	_ReduceNamedToAtomType                                    = _ReduceType(139)
	_ReduceExplicitStructDefToAtomType                        = _ReduceType(140)
	_ReduceImplicitStructDefToAtomType                        = _ReduceType(141)
	_ReduceExplicitEnumDefToAtomType                          = _ReduceType(142)
	_ReduceImplicitEnumDefToAtomType                          = _ReduceType(143)
	_ReduceTraitDefToAtomType                                 = _ReduceType(144)
	_ReduceLexErrorToAtomType                                 = _ReduceType(145)
	_ReduceAtomTypeToTraitableType                            = _ReduceType(146)
	_ReducePublicMethodsTraitToTraitableType                  = _ReduceType(147)
	_ReducePublicTraitToTraitableType                         = _ReduceType(148)
	_ReduceTraitableTypeToTraitAlgebraType                    = _ReduceType(149)
	_ReduceIntersectToTraitAlgebraType                        = _ReduceType(150)
	_ReduceUnionToTraitAlgebraType                            = _ReduceType(151)
	_ReduceDifferenceToTraitAlgebraType                       = _ReduceType(152)
	_ReduceInferredToValueType                                = _ReduceType(153)
	_ReduceTraitAlgebraTypeToValueType                        = _ReduceType(154)
	_ReduceReferenceToValueType                               = _ReduceType(155)
	_ReduceFuncTypeToValueType                                = _ReduceType(156)
	_ReduceDefinitionToTypeDef                                = _ReduceType(157)
	_ReduceConstrainedDefToTypeDef                            = _ReduceType(158)
	_ReduceAliasToTypeDef                                     = _ReduceType(159)
	_ReduceUnconstrainedToGenericParameterDef                 = _ReduceType(160)
	_ReduceConstrainedToGenericParameterDef                   = _ReduceType(161)
	_ReduceGenericParameterDefToGenericParameterDefs          = _ReduceType(162)
	_ReduceAddToGenericParameterDefs                          = _ReduceType(163)
	_ReduceGenericParameterDefsToOptionalGenericParameterDefs = _ReduceType(164)
	_ReduceNilToOptionalGenericParameterDefs                  = _ReduceType(165)
	_ReduceGenericToOptionalGenericParameters                 = _ReduceType(166)
	_ReduceNilToOptionalGenericParameters                     = _ReduceType(167)
	_ReduceExplicitToFieldDef                                 = _ReduceType(168)
	_ReduceImplicitToFieldDef                                 = _ReduceType(169)
	_ReduceUnsafeStatementToFieldDef                          = _ReduceType(170)
	_ReduceFieldDefToImplicitFieldDefs                        = _ReduceType(171)
	_ReduceAddToImplicitFieldDefs                             = _ReduceType(172)
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefs       = _ReduceType(173)
	_ReduceNilToOptionalImplicitFieldDefs                     = _ReduceType(174)
	_ReduceToImplicitStructDef                                = _ReduceType(175)
	_ReduceFieldDefToExplicitFieldDefs                        = _ReduceType(176)
	_ReduceImplicitToExplicitFieldDefs                        = _ReduceType(177)
	_ReduceExplicitToExplicitFieldDefs                        = _ReduceType(178)
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefs       = _ReduceType(179)
	_ReduceNilToOptionalExplicitFieldDefs                     = _ReduceType(180)
	_ReduceToExplicitStructDef                                = _ReduceType(181)
	_ReduceFieldDefToEnumValueDef                             = _ReduceType(182)
	_ReduceDefaultToEnumValueDef                              = _ReduceType(183)
	_ReducePairToImplicitEnumValueDefs                        = _ReduceType(184)
	_ReduceAddToImplicitEnumValueDefs                         = _ReduceType(185)
	_ReduceToImplicitEnumDef                                  = _ReduceType(186)
	_ReduceExplicitPairToExplicitEnumValueDefs                = _ReduceType(187)
	_ReduceImplicitPairToExplicitEnumValueDefs                = _ReduceType(188)
	_ReduceExplicitAddToExplicitEnumValueDefs                 = _ReduceType(189)
	_ReduceImplicitAddToExplicitEnumValueDefs                 = _ReduceType(190)
	_ReduceToExplicitEnumDef                                  = _ReduceType(191)
	_ReduceFieldDefToTraitProperty                            = _ReduceType(192)
	_ReduceMethodSignatureToTraitProperty                     = _ReduceType(193)
	_ReduceTraitPropertyToTraitProperties                     = _ReduceType(194)
	_ReduceImplicitToTraitProperties                          = _ReduceType(195)
	_ReduceExplicitToTraitProperties                          = _ReduceType(196)
	_ReduceTraitPropertiesToOptionalTraitProperties           = _ReduceType(197)
	_ReduceNilToOptionalTraitProperties                       = _ReduceType(198)
	_ReduceToTraitDef                                         = _ReduceType(199)
	_ReduceValueTypeToReturnType                              = _ReduceType(200)
	_ReduceNilToReturnType                                    = _ReduceType(201)
	_ReduceArgToParameterDecl                                 = _ReduceType(202)
	_ReduceVarargToParameterDecl                              = _ReduceType(203)
	_ReduceUnamedToParameterDecl                              = _ReduceType(204)
	_ReduceUnnamedVarargToParameterDecl                       = _ReduceType(205)
	_ReduceParameterDeclToParameterDecls                      = _ReduceType(206)
	_ReduceAddToParameterDecls                                = _ReduceType(207)
	_ReduceParameterDeclsToOptionalParameterDecls             = _ReduceType(208)
	_ReduceNilToOptionalParameterDecls                        = _ReduceType(209)
	_ReduceToFuncType                                         = _ReduceType(210)
	_ReduceToMethodSignature                                  = _ReduceType(211)
	_ReduceArgToParameterDef                                  = _ReduceType(212)
	_ReduceVarargToParameterDef                               = _ReduceType(213)
	_ReduceParameterDefToParameterDefs                        = _ReduceType(214)
	_ReduceAddToParameterDefs                                 = _ReduceType(215)
	_ReduceParameterDefsToOptionalParameterDefs               = _ReduceType(216)
	_ReduceNilToOptionalParameterDefs                         = _ReduceType(217)
	_ReduceReceiverToOptionalReceiver                         = _ReduceType(218)
	_ReduceNilToOptionalReceiver                              = _ReduceType(219)
	_ReduceToNamedFuncDef                                     = _ReduceType(220)
	_ReduceToAnonymousFuncExpr                                = _ReduceType(221)
	_ReduceNoSpecToPackageDef                                 = _ReduceType(222)
	_ReduceWithSpecToPackageDef                               = _ReduceType(223)
	_ReduceToPackageStatementBody                             = _ReduceType(224)
	_ReduceImplicitToPackageStatement                         = _ReduceType(225)
	_ReduceExplicitToPackageStatement                         = _ReduceType(226)
	_ReduceEmptyListToPackageStatements                       = _ReduceType(227)
	_ReduceAddToPackageStatements                             = _ReduceType(228)
	_ReduceSpacesToLexInternalTokens                          = _ReduceType(229)
	_ReduceCommentToLexInternalTokens                         = _ReduceType(230)
)

func (i _ReduceType) String() string {
	switch i {
	case _ReduceToSource:
		return "ToSource"
	case _ReduceNewlinesToOptionalDefinitions:
		return "NewlinesToOptionalDefinitions"
	case _ReduceDefinitionsToOptionalDefinitions:
		return "DefinitionsToOptionalDefinitions"
	case _ReduceNilToOptionalDefinitions:
		return "NilToOptionalDefinitions"
	case _ReduceNewlinesToOptionalNewlines:
		return "NewlinesToOptionalNewlines"
	case _ReduceNilToOptionalNewlines:
		return "NilToOptionalNewlines"
	case _ReduceNilToDefinitions:
		return "NilToDefinitions"
	case _ReduceAddToDefinitions:
		return "AddToDefinitions"
	case _ReducePackageDefToDefinition:
		return "PackageDefToDefinition"
	case _ReduceTypeDefToDefinition:
		return "TypeDefToDefinition"
	case _ReduceNamedFuncDefToDefinition:
		return "NamedFuncDefToDefinition"
	case _ReduceUnsafeStatementToDefinition:
		return "UnsafeStatementToDefinition"
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
	case _ReduceLexErrorToAtomType:
		return "LexErrorToAtomType"
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
	_State353 = _StateId(353)
	_State354 = _StateId(354)
	_State355 = _StateId(355)
	_State356 = _StateId(356)
	_State357 = _StateId(357)
	_State358 = _StateId(358)
	_State359 = _StateId(359)
	_State360 = _StateId(360)
	_State361 = _StateId(361)
	_State362 = _StateId(362)
	_State363 = _StateId(363)
	_State364 = _StateId(364)
	_State365 = _StateId(365)
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
	case _ReduceToSource:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SourceType
		symbol.Generic_, err = reducer.ToSource(args[0].Generic_)
	case _ReduceNewlinesToOptionalDefinitions:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalDefinitionsType
		symbol.Generic_, err = reducer.NewlinesToOptionalDefinitions(args[0].Generic_)
	case _ReduceDefinitionsToOptionalDefinitions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OptionalDefinitionsType
		symbol.Generic_, err = reducer.DefinitionsToOptionalDefinitions(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceNilToOptionalDefinitions:
		symbol.SymbolId_ = OptionalDefinitionsType
		symbol.Generic_, err = reducer.NilToOptionalDefinitions()
	case _ReduceNewlinesToOptionalNewlines:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalNewlinesType
		symbol.Generic_, err = reducer.NewlinesToOptionalNewlines(args[0].Generic_)
	case _ReduceNilToOptionalNewlines:
		symbol.SymbolId_ = OptionalNewlinesType
		symbol.Generic_, err = reducer.NilToOptionalNewlines()
	case _ReduceNilToDefinitions:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionsType
		symbol.Generic_, err = reducer.NilToDefinitions(args[0].Generic_)
	case _ReduceAddToDefinitions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = DefinitionsType
		symbol.Generic_, err = reducer.AddToDefinitions(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReducePackageDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.Generic_, err = reducer.PackageDefToDefinition(args[0].Generic_)
	case _ReduceTypeDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.Generic_, err = reducer.TypeDefToDefinition(args[0].Generic_)
	case _ReduceNamedFuncDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.Generic_, err = reducer.NamedFuncDefToDefinition(args[0].Generic_)
	case _ReduceUnsafeStatementToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.Generic_, err = reducer.UnsafeStatementToDefinition(args[0].Generic_)
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
	case _ReduceLexErrorToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.LexErrorToAtomType(args[0].Generic_)
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
	_GotoState353Action                                             = &_Action{_ShiftAction, _State353, 0}
	_GotoState354Action                                             = &_Action{_ShiftAction, _State354, 0}
	_GotoState355Action                                             = &_Action{_ShiftAction, _State355, 0}
	_GotoState356Action                                             = &_Action{_ShiftAction, _State356, 0}
	_GotoState357Action                                             = &_Action{_ShiftAction, _State357, 0}
	_GotoState358Action                                             = &_Action{_ShiftAction, _State358, 0}
	_GotoState359Action                                             = &_Action{_ShiftAction, _State359, 0}
	_GotoState360Action                                             = &_Action{_ShiftAction, _State360, 0}
	_GotoState361Action                                             = &_Action{_ShiftAction, _State361, 0}
	_GotoState362Action                                             = &_Action{_ShiftAction, _State362, 0}
	_GotoState363Action                                             = &_Action{_ShiftAction, _State363, 0}
	_GotoState364Action                                             = &_Action{_ShiftAction, _State364, 0}
	_GotoState365Action                                             = &_Action{_ShiftAction, _State365, 0}
	_ReduceToSourceAction                                           = &_Action{_ReduceAction, 0, _ReduceToSource}
	_ReduceNewlinesToOptionalDefinitionsAction                      = &_Action{_ReduceAction, 0, _ReduceNewlinesToOptionalDefinitions}
	_ReduceDefinitionsToOptionalDefinitionsAction                   = &_Action{_ReduceAction, 0, _ReduceDefinitionsToOptionalDefinitions}
	_ReduceNilToOptionalDefinitionsAction                           = &_Action{_ReduceAction, 0, _ReduceNilToOptionalDefinitions}
	_ReduceNewlinesToOptionalNewlinesAction                         = &_Action{_ReduceAction, 0, _ReduceNewlinesToOptionalNewlines}
	_ReduceNilToOptionalNewlinesAction                              = &_Action{_ReduceAction, 0, _ReduceNilToOptionalNewlines}
	_ReduceNilToDefinitionsAction                                   = &_Action{_ReduceAction, 0, _ReduceNilToDefinitions}
	_ReduceAddToDefinitionsAction                                   = &_Action{_ReduceAction, 0, _ReduceAddToDefinitions}
	_ReducePackageDefToDefinitionAction                             = &_Action{_ReduceAction, 0, _ReducePackageDefToDefinition}
	_ReduceTypeDefToDefinitionAction                                = &_Action{_ReduceAction, 0, _ReduceTypeDefToDefinition}
	_ReduceNamedFuncDefToDefinitionAction                           = &_Action{_ReduceAction, 0, _ReduceNamedFuncDefToDefinition}
	_ReduceUnsafeStatementToDefinitionAction                        = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToDefinition}
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
	_ReduceLexErrorToAtomTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceLexErrorToAtomType}
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
	{_State1, NewlinesToken}:                      _GotoState13Action,
	{_State1, SourceType}:                         _GotoState7Action,
	{_State1, OptionalDefinitionsType}:            _GotoState14Action,
	{_State1, OptionalNewlinesType}:               _GotoState15Action,
	{_State2, PackageToken}:                       _GotoState16Action,
	{_State2, PackageDefType}:                     _GotoState8Action,
	{_State3, TypeToken}:                          _GotoState17Action,
	{_State3, TypeDefType}:                        _GotoState9Action,
	{_State4, FuncToken}:                          _GotoState18Action,
	{_State4, NamedFuncDefType}:                   _GotoState10Action,
	{_State5, IntegerLiteralToken}:                _GotoState25Action,
	{_State5, FloatLiteralToken}:                  _GotoState22Action,
	{_State5, RuneLiteralToken}:                   _GotoState31Action,
	{_State5, StringLiteralToken}:                 _GotoState32Action,
	{_State5, IdentifierToken}:                    _GotoState24Action,
	{_State5, TrueToken}:                          _GotoState35Action,
	{_State5, FalseToken}:                         _GotoState21Action,
	{_State5, StructToken}:                        _GotoState33Action,
	{_State5, FuncToken}:                          _GotoState23Action,
	{_State5, LabelDeclToken}:                     _GotoState26Action,
	{_State5, LparenToken}:                        _GotoState28Action,
	{_State5, NotToken}:                           _GotoState30Action,
	{_State5, SubToken}:                           _GotoState34Action,
	{_State5, MulToken}:                           _GotoState29Action,
	{_State5, BitNegToken}:                        _GotoState20Action,
	{_State5, BitAndToken}:                        _GotoState19Action,
	{_State5, LexErrorToken}:                      _GotoState27Action,
	{_State5, ExpressionType}:                     _GotoState11Action,
	{_State5, OptionalLabelDeclType}:              _GotoState48Action,
	{_State5, SequenceExprType}:                   _GotoState53Action,
	{_State5, BlockExprType}:                      _GotoState42Action,
	{_State5, CallExprType}:                       _GotoState43Action,
	{_State5, AtomExprType}:                       _GotoState41Action,
	{_State5, LiteralType}:                        _GotoState46Action,
	{_State5, AnonymousStructExprType}:            _GotoState40Action,
	{_State5, AccessExprType}:                     _GotoState36Action,
	{_State5, PostfixUnaryExprType}:               _GotoState50Action,
	{_State5, PrefixUnaryOpType}:                  _GotoState52Action,
	{_State5, PrefixUnaryExprType}:                _GotoState51Action,
	{_State5, MulExprType}:                        _GotoState47Action,
	{_State5, AddExprType}:                        _GotoState37Action,
	{_State5, CmpExprType}:                        _GotoState44Action,
	{_State5, AndExprType}:                        _GotoState38Action,
	{_State5, OrExprType}:                         _GotoState49Action,
	{_State5, ExplicitStructDefType}:              _GotoState45Action,
	{_State5, AnonymousFuncExprType}:              _GotoState39Action,
	{_State6, SpacesToken}:                        _GotoState55Action,
	{_State6, CommentToken}:                       _GotoState54Action,
	{_State6, LexInternalTokensType}:              _GotoState12Action,
	{_State15, PackageToken}:                      _GotoState16Action,
	{_State15, UnsafeToken}:                       _GotoState56Action,
	{_State15, TypeToken}:                         _GotoState17Action,
	{_State15, FuncToken}:                         _GotoState18Action,
	{_State15, DefinitionsType}:                   _GotoState58Action,
	{_State15, DefinitionType}:                    _GotoState57Action,
	{_State15, UnsafeStatementType}:               _GotoState62Action,
	{_State15, TypeDefType}:                       _GotoState61Action,
	{_State15, NamedFuncDefType}:                  _GotoState59Action,
	{_State15, PackageDefType}:                    _GotoState60Action,
	{_State16, IdentifierToken}:                   _GotoState63Action,
	{_State17, IdentifierToken}:                   _GotoState64Action,
	{_State18, LparenToken}:                       _GotoState65Action,
	{_State18, OptionalReceiverType}:              _GotoState66Action,
	{_State23, LparenToken}:                       _GotoState67Action,
	{_State28, IntegerLiteralToken}:               _GotoState25Action,
	{_State28, FloatLiteralToken}:                 _GotoState22Action,
	{_State28, RuneLiteralToken}:                  _GotoState31Action,
	{_State28, StringLiteralToken}:                _GotoState32Action,
	{_State28, IdentifierToken}:                   _GotoState68Action,
	{_State28, TrueToken}:                         _GotoState35Action,
	{_State28, FalseToken}:                        _GotoState21Action,
	{_State28, StructToken}:                       _GotoState33Action,
	{_State28, FuncToken}:                         _GotoState23Action,
	{_State28, LabelDeclToken}:                    _GotoState26Action,
	{_State28, LparenToken}:                       _GotoState28Action,
	{_State28, NotToken}:                          _GotoState30Action,
	{_State28, SubToken}:                          _GotoState34Action,
	{_State28, MulToken}:                          _GotoState29Action,
	{_State28, BitNegToken}:                       _GotoState20Action,
	{_State28, BitAndToken}:                       _GotoState19Action,
	{_State28, LexErrorToken}:                     _GotoState27Action,
	{_State28, ExpressionType}:                    _GotoState72Action,
	{_State28, OptionalLabelDeclType}:             _GotoState48Action,
	{_State28, SequenceExprType}:                  _GotoState53Action,
	{_State28, BlockExprType}:                     _GotoState42Action,
	{_State28, CallExprType}:                      _GotoState43Action,
	{_State28, ArgumentsType}:                     _GotoState70Action,
	{_State28, ArgumentType}:                      _GotoState69Action,
	{_State28, ColonExpressionsType}:              _GotoState71Action,
	{_State28, OptionalExpressionType}:            _GotoState73Action,
	{_State28, AtomExprType}:                      _GotoState41Action,
	{_State28, LiteralType}:                       _GotoState46Action,
	{_State28, AnonymousStructExprType}:           _GotoState40Action,
	{_State28, AccessExprType}:                    _GotoState36Action,
	{_State28, PostfixUnaryExprType}:              _GotoState50Action,
	{_State28, PrefixUnaryOpType}:                 _GotoState52Action,
	{_State28, PrefixUnaryExprType}:               _GotoState51Action,
	{_State28, MulExprType}:                       _GotoState47Action,
	{_State28, AddExprType}:                       _GotoState37Action,
	{_State28, CmpExprType}:                       _GotoState44Action,
	{_State28, AndExprType}:                       _GotoState38Action,
	{_State28, OrExprType}:                        _GotoState49Action,
	{_State28, ExplicitStructDefType}:             _GotoState45Action,
	{_State28, AnonymousFuncExprType}:             _GotoState39Action,
	{_State33, LparenToken}:                       _GotoState74Action,
	{_State36, LbracketToken}:                     _GotoState77Action,
	{_State36, DotToken}:                          _GotoState76Action,
	{_State36, QuestionToken}:                     _GotoState78Action,
	{_State36, DollarLbracketToken}:               _GotoState75Action,
	{_State36, OptionalGenericBindingType}:        _GotoState79Action,
	{_State37, AddToken}:                          _GotoState80Action,
	{_State37, SubToken}:                          _GotoState83Action,
	{_State37, BitXorToken}:                       _GotoState82Action,
	{_State37, BitOrToken}:                        _GotoState81Action,
	{_State37, AddOpType}:                         _GotoState84Action,
	{_State38, AndToken}:                          _GotoState85Action,
	{_State44, EqualToken}:                        _GotoState86Action,
	{_State44, NotEqualToken}:                     _GotoState91Action,
	{_State44, LessToken}:                         _GotoState89Action,
	{_State44, LessOrEqualToken}:                  _GotoState90Action,
	{_State44, GreaterToken}:                      _GotoState87Action,
	{_State44, GreaterOrEqualToken}:               _GotoState88Action,
	{_State44, CmpOpType}:                         _GotoState92Action,
	{_State45, LparenToken}:                       _GotoState93Action,
	{_State47, MulToken}:                          _GotoState99Action,
	{_State47, DivToken}:                          _GotoState97Action,
	{_State47, ModToken}:                          _GotoState98Action,
	{_State47, BitAndToken}:                       _GotoState94Action,
	{_State47, BitLshiftToken}:                    _GotoState95Action,
	{_State47, BitRshiftToken}:                    _GotoState96Action,
	{_State47, MulOpType}:                         _GotoState100Action,
	{_State48, IfToken}:                           _GotoState103Action,
	{_State48, SwitchToken}:                       _GotoState105Action,
	{_State48, ForToken}:                          _GotoState102Action,
	{_State48, DoToken}:                           _GotoState101Action,
	{_State48, LbraceToken}:                       _GotoState104Action,
	{_State48, IfExprType}:                        _GotoState107Action,
	{_State48, SwitchExprType}:                    _GotoState109Action,
	{_State48, LoopExprType}:                      _GotoState108Action,
	{_State48, BlockBodyType}:                     _GotoState106Action,
	{_State49, OrToken}:                           _GotoState110Action,
	{_State52, IntegerLiteralToken}:               _GotoState25Action,
	{_State52, FloatLiteralToken}:                 _GotoState22Action,
	{_State52, RuneLiteralToken}:                  _GotoState31Action,
	{_State52, StringLiteralToken}:                _GotoState32Action,
	{_State52, IdentifierToken}:                   _GotoState24Action,
	{_State52, TrueToken}:                         _GotoState35Action,
	{_State52, FalseToken}:                        _GotoState21Action,
	{_State52, StructToken}:                       _GotoState33Action,
	{_State52, FuncToken}:                         _GotoState23Action,
	{_State52, LabelDeclToken}:                    _GotoState26Action,
	{_State52, LparenToken}:                       _GotoState28Action,
	{_State52, NotToken}:                          _GotoState30Action,
	{_State52, SubToken}:                          _GotoState34Action,
	{_State52, MulToken}:                          _GotoState29Action,
	{_State52, BitNegToken}:                       _GotoState20Action,
	{_State52, BitAndToken}:                       _GotoState19Action,
	{_State52, LexErrorToken}:                     _GotoState27Action,
	{_State52, OptionalLabelDeclType}:             _GotoState111Action,
	{_State52, BlockExprType}:                     _GotoState42Action,
	{_State52, CallExprType}:                      _GotoState43Action,
	{_State52, AtomExprType}:                      _GotoState41Action,
	{_State52, LiteralType}:                       _GotoState46Action,
	{_State52, AnonymousStructExprType}:           _GotoState40Action,
	{_State52, AccessExprType}:                    _GotoState36Action,
	{_State52, PostfixUnaryExprType}:              _GotoState50Action,
	{_State52, PrefixUnaryOpType}:                 _GotoState52Action,
	{_State52, PrefixUnaryExprType}:               _GotoState112Action,
	{_State52, ExplicitStructDefType}:             _GotoState45Action,
	{_State52, AnonymousFuncExprType}:             _GotoState39Action,
	{_State56, LessToken}:                         _GotoState113Action,
	{_State58, NewlinesToken}:                     _GotoState114Action,
	{_State58, OptionalNewlinesType}:              _GotoState115Action,
	{_State63, LparenToken}:                       _GotoState116Action,
	{_State64, DollarLbracketToken}:               _GotoState117Action,
	{_State64, EqualToken}:                        _GotoState118Action,
	{_State64, OptionalGenericParametersType}:     _GotoState119Action,
	{_State65, IdentifierToken}:                   _GotoState120Action,
	{_State65, ParameterDefType}:                  _GotoState121Action,
	{_State66, IdentifierToken}:                   _GotoState122Action,
	{_State67, IdentifierToken}:                   _GotoState120Action,
	{_State67, ParameterDefType}:                  _GotoState124Action,
	{_State67, ParameterDefsType}:                 _GotoState125Action,
	{_State67, OptionalParameterDefsType}:         _GotoState123Action,
	{_State68, AssignToken}:                       _GotoState126Action,
	{_State70, RparenToken}:                       _GotoState128Action,
	{_State70, CommaToken}:                        _GotoState127Action,
	{_State71, ColonToken}:                        _GotoState129Action,
	{_State73, ColonToken}:                        _GotoState130Action,
	{_State74, IdentifierToken}:                   _GotoState135Action,
	{_State74, UnsafeToken}:                       _GotoState56Action,
	{_State74, StructToken}:                       _GotoState33Action,
	{_State74, EnumToken}:                         _GotoState133Action,
	{_State74, TraitToken}:                        _GotoState140Action,
	{_State74, FuncToken}:                         _GotoState134Action,
	{_State74, LparenToken}:                       _GotoState137Action,
	{_State74, QuestionToken}:                     _GotoState138Action,
	{_State74, TildeTildeToken}:                   _GotoState139Action,
	{_State74, BitNegToken}:                       _GotoState132Action,
	{_State74, BitAndToken}:                       _GotoState131Action,
	{_State74, LexErrorToken}:                     _GotoState136Action,
	{_State74, UnsafeStatementType}:               _GotoState153Action,
	{_State74, AtomTypeType}:                      _GotoState141Action,
	{_State74, TraitableTypeType}:                 _GotoState152Action,
	{_State74, TraitAlgebraTypeType}:              _GotoState150Action,
	{_State74, ValueTypeType}:                     _GotoState154Action,
	{_State74, FieldDefType}:                      _GotoState145Action,
	{_State74, ImplicitStructDefType}:             _GotoState148Action,
	{_State74, ExplicitFieldDefsType}:             _GotoState143Action,
	{_State74, OptionalExplicitFieldDefsType}:     _GotoState149Action,
	{_State74, ExplicitStructDefType}:             _GotoState144Action,
	{_State74, ImplicitEnumDefType}:               _GotoState147Action,
	{_State74, ExplicitEnumDefType}:               _GotoState142Action,
	{_State74, TraitDefType}:                      _GotoState151Action,
	{_State74, FuncTypeType}:                      _GotoState146Action,
	{_State75, IdentifierToken}:                   _GotoState155Action,
	{_State75, StructToken}:                       _GotoState33Action,
	{_State75, EnumToken}:                         _GotoState133Action,
	{_State75, TraitToken}:                        _GotoState140Action,
	{_State75, FuncToken}:                         _GotoState134Action,
	{_State75, LparenToken}:                       _GotoState137Action,
	{_State75, QuestionToken}:                     _GotoState138Action,
	{_State75, TildeTildeToken}:                   _GotoState139Action,
	{_State75, BitNegToken}:                       _GotoState132Action,
	{_State75, BitAndToken}:                       _GotoState131Action,
	{_State75, LexErrorToken}:                     _GotoState136Action,
	{_State75, OptionalGenericArgumentsType}:      _GotoState157Action,
	{_State75, GenericArgumentsType}:              _GotoState156Action,
	{_State75, AtomTypeType}:                      _GotoState141Action,
	{_State75, TraitableTypeType}:                 _GotoState152Action,
	{_State75, TraitAlgebraTypeType}:              _GotoState150Action,
	{_State75, ValueTypeType}:                     _GotoState158Action,
	{_State75, ImplicitStructDefType}:             _GotoState148Action,
	{_State75, ExplicitStructDefType}:             _GotoState144Action,
	{_State75, ImplicitEnumDefType}:               _GotoState147Action,
	{_State75, ExplicitEnumDefType}:               _GotoState142Action,
	{_State75, TraitDefType}:                      _GotoState151Action,
	{_State75, FuncTypeType}:                      _GotoState146Action,
	{_State76, IdentifierToken}:                   _GotoState159Action,
	{_State77, IntegerLiteralToken}:               _GotoState25Action,
	{_State77, FloatLiteralToken}:                 _GotoState22Action,
	{_State77, RuneLiteralToken}:                  _GotoState31Action,
	{_State77, StringLiteralToken}:                _GotoState32Action,
	{_State77, IdentifierToken}:                   _GotoState68Action,
	{_State77, TrueToken}:                         _GotoState35Action,
	{_State77, FalseToken}:                        _GotoState21Action,
	{_State77, StructToken}:                       _GotoState33Action,
	{_State77, FuncToken}:                         _GotoState23Action,
	{_State77, LabelDeclToken}:                    _GotoState26Action,
	{_State77, LparenToken}:                       _GotoState28Action,
	{_State77, NotToken}:                          _GotoState30Action,
	{_State77, SubToken}:                          _GotoState34Action,
	{_State77, MulToken}:                          _GotoState29Action,
	{_State77, BitNegToken}:                       _GotoState20Action,
	{_State77, BitAndToken}:                       _GotoState19Action,
	{_State77, LexErrorToken}:                     _GotoState27Action,
	{_State77, ExpressionType}:                    _GotoState72Action,
	{_State77, OptionalLabelDeclType}:             _GotoState48Action,
	{_State77, SequenceExprType}:                  _GotoState53Action,
	{_State77, BlockExprType}:                     _GotoState42Action,
	{_State77, CallExprType}:                      _GotoState43Action,
	{_State77, ArgumentType}:                      _GotoState160Action,
	{_State77, ColonExpressionsType}:              _GotoState71Action,
	{_State77, OptionalExpressionType}:            _GotoState73Action,
	{_State77, AtomExprType}:                      _GotoState41Action,
	{_State77, LiteralType}:                       _GotoState46Action,
	{_State77, AnonymousStructExprType}:           _GotoState40Action,
	{_State77, AccessExprType}:                    _GotoState36Action,
	{_State77, PostfixUnaryExprType}:              _GotoState50Action,
	{_State77, PrefixUnaryOpType}:                 _GotoState52Action,
	{_State77, PrefixUnaryExprType}:               _GotoState51Action,
	{_State77, MulExprType}:                       _GotoState47Action,
	{_State77, AddExprType}:                       _GotoState37Action,
	{_State77, CmpExprType}:                       _GotoState44Action,
	{_State77, AndExprType}:                       _GotoState38Action,
	{_State77, OrExprType}:                        _GotoState49Action,
	{_State77, ExplicitStructDefType}:             _GotoState45Action,
	{_State77, AnonymousFuncExprType}:             _GotoState39Action,
	{_State79, LparenToken}:                       _GotoState161Action,
	{_State84, IntegerLiteralToken}:               _GotoState25Action,
	{_State84, FloatLiteralToken}:                 _GotoState22Action,
	{_State84, RuneLiteralToken}:                  _GotoState31Action,
	{_State84, StringLiteralToken}:                _GotoState32Action,
	{_State84, IdentifierToken}:                   _GotoState24Action,
	{_State84, TrueToken}:                         _GotoState35Action,
	{_State84, FalseToken}:                        _GotoState21Action,
	{_State84, StructToken}:                       _GotoState33Action,
	{_State84, FuncToken}:                         _GotoState23Action,
	{_State84, LabelDeclToken}:                    _GotoState26Action,
	{_State84, LparenToken}:                       _GotoState28Action,
	{_State84, NotToken}:                          _GotoState30Action,
	{_State84, SubToken}:                          _GotoState34Action,
	{_State84, MulToken}:                          _GotoState29Action,
	{_State84, BitNegToken}:                       _GotoState20Action,
	{_State84, BitAndToken}:                       _GotoState19Action,
	{_State84, LexErrorToken}:                     _GotoState27Action,
	{_State84, OptionalLabelDeclType}:             _GotoState111Action,
	{_State84, BlockExprType}:                     _GotoState42Action,
	{_State84, CallExprType}:                      _GotoState43Action,
	{_State84, AtomExprType}:                      _GotoState41Action,
	{_State84, LiteralType}:                       _GotoState46Action,
	{_State84, AnonymousStructExprType}:           _GotoState40Action,
	{_State84, AccessExprType}:                    _GotoState36Action,
	{_State84, PostfixUnaryExprType}:              _GotoState50Action,
	{_State84, PrefixUnaryOpType}:                 _GotoState52Action,
	{_State84, PrefixUnaryExprType}:               _GotoState51Action,
	{_State84, MulExprType}:                       _GotoState162Action,
	{_State84, ExplicitStructDefType}:             _GotoState45Action,
	{_State84, AnonymousFuncExprType}:             _GotoState39Action,
	{_State85, IntegerLiteralToken}:               _GotoState25Action,
	{_State85, FloatLiteralToken}:                 _GotoState22Action,
	{_State85, RuneLiteralToken}:                  _GotoState31Action,
	{_State85, StringLiteralToken}:                _GotoState32Action,
	{_State85, IdentifierToken}:                   _GotoState24Action,
	{_State85, TrueToken}:                         _GotoState35Action,
	{_State85, FalseToken}:                        _GotoState21Action,
	{_State85, StructToken}:                       _GotoState33Action,
	{_State85, FuncToken}:                         _GotoState23Action,
	{_State85, LabelDeclToken}:                    _GotoState26Action,
	{_State85, LparenToken}:                       _GotoState28Action,
	{_State85, NotToken}:                          _GotoState30Action,
	{_State85, SubToken}:                          _GotoState34Action,
	{_State85, MulToken}:                          _GotoState29Action,
	{_State85, BitNegToken}:                       _GotoState20Action,
	{_State85, BitAndToken}:                       _GotoState19Action,
	{_State85, LexErrorToken}:                     _GotoState27Action,
	{_State85, OptionalLabelDeclType}:             _GotoState111Action,
	{_State85, BlockExprType}:                     _GotoState42Action,
	{_State85, CallExprType}:                      _GotoState43Action,
	{_State85, AtomExprType}:                      _GotoState41Action,
	{_State85, LiteralType}:                       _GotoState46Action,
	{_State85, AnonymousStructExprType}:           _GotoState40Action,
	{_State85, AccessExprType}:                    _GotoState36Action,
	{_State85, PostfixUnaryExprType}:              _GotoState50Action,
	{_State85, PrefixUnaryOpType}:                 _GotoState52Action,
	{_State85, PrefixUnaryExprType}:               _GotoState51Action,
	{_State85, MulExprType}:                       _GotoState47Action,
	{_State85, AddExprType}:                       _GotoState37Action,
	{_State85, CmpExprType}:                       _GotoState163Action,
	{_State85, ExplicitStructDefType}:             _GotoState45Action,
	{_State85, AnonymousFuncExprType}:             _GotoState39Action,
	{_State92, IntegerLiteralToken}:               _GotoState25Action,
	{_State92, FloatLiteralToken}:                 _GotoState22Action,
	{_State92, RuneLiteralToken}:                  _GotoState31Action,
	{_State92, StringLiteralToken}:                _GotoState32Action,
	{_State92, IdentifierToken}:                   _GotoState24Action,
	{_State92, TrueToken}:                         _GotoState35Action,
	{_State92, FalseToken}:                        _GotoState21Action,
	{_State92, StructToken}:                       _GotoState33Action,
	{_State92, FuncToken}:                         _GotoState23Action,
	{_State92, LabelDeclToken}:                    _GotoState26Action,
	{_State92, LparenToken}:                       _GotoState28Action,
	{_State92, NotToken}:                          _GotoState30Action,
	{_State92, SubToken}:                          _GotoState34Action,
	{_State92, MulToken}:                          _GotoState29Action,
	{_State92, BitNegToken}:                       _GotoState20Action,
	{_State92, BitAndToken}:                       _GotoState19Action,
	{_State92, LexErrorToken}:                     _GotoState27Action,
	{_State92, OptionalLabelDeclType}:             _GotoState111Action,
	{_State92, BlockExprType}:                     _GotoState42Action,
	{_State92, CallExprType}:                      _GotoState43Action,
	{_State92, AtomExprType}:                      _GotoState41Action,
	{_State92, LiteralType}:                       _GotoState46Action,
	{_State92, AnonymousStructExprType}:           _GotoState40Action,
	{_State92, AccessExprType}:                    _GotoState36Action,
	{_State92, PostfixUnaryExprType}:              _GotoState50Action,
	{_State92, PrefixUnaryOpType}:                 _GotoState52Action,
	{_State92, PrefixUnaryExprType}:               _GotoState51Action,
	{_State92, MulExprType}:                       _GotoState47Action,
	{_State92, AddExprType}:                       _GotoState164Action,
	{_State92, ExplicitStructDefType}:             _GotoState45Action,
	{_State92, AnonymousFuncExprType}:             _GotoState39Action,
	{_State93, IntegerLiteralToken}:               _GotoState25Action,
	{_State93, FloatLiteralToken}:                 _GotoState22Action,
	{_State93, RuneLiteralToken}:                  _GotoState31Action,
	{_State93, StringLiteralToken}:                _GotoState32Action,
	{_State93, IdentifierToken}:                   _GotoState68Action,
	{_State93, TrueToken}:                         _GotoState35Action,
	{_State93, FalseToken}:                        _GotoState21Action,
	{_State93, StructToken}:                       _GotoState33Action,
	{_State93, FuncToken}:                         _GotoState23Action,
	{_State93, LabelDeclToken}:                    _GotoState26Action,
	{_State93, LparenToken}:                       _GotoState28Action,
	{_State93, NotToken}:                          _GotoState30Action,
	{_State93, SubToken}:                          _GotoState34Action,
	{_State93, MulToken}:                          _GotoState29Action,
	{_State93, BitNegToken}:                       _GotoState20Action,
	{_State93, BitAndToken}:                       _GotoState19Action,
	{_State93, LexErrorToken}:                     _GotoState27Action,
	{_State93, ExpressionType}:                    _GotoState72Action,
	{_State93, OptionalLabelDeclType}:             _GotoState48Action,
	{_State93, SequenceExprType}:                  _GotoState53Action,
	{_State93, BlockExprType}:                     _GotoState42Action,
	{_State93, CallExprType}:                      _GotoState43Action,
	{_State93, ArgumentsType}:                     _GotoState165Action,
	{_State93, ArgumentType}:                      _GotoState69Action,
	{_State93, ColonExpressionsType}:              _GotoState71Action,
	{_State93, OptionalExpressionType}:            _GotoState73Action,
	{_State93, AtomExprType}:                      _GotoState41Action,
	{_State93, LiteralType}:                       _GotoState46Action,
	{_State93, AnonymousStructExprType}:           _GotoState40Action,
	{_State93, AccessExprType}:                    _GotoState36Action,
	{_State93, PostfixUnaryExprType}:              _GotoState50Action,
	{_State93, PrefixUnaryOpType}:                 _GotoState52Action,
	{_State93, PrefixUnaryExprType}:               _GotoState51Action,
	{_State93, MulExprType}:                       _GotoState47Action,
	{_State93, AddExprType}:                       _GotoState37Action,
	{_State93, CmpExprType}:                       _GotoState44Action,
	{_State93, AndExprType}:                       _GotoState38Action,
	{_State93, OrExprType}:                        _GotoState49Action,
	{_State93, ExplicitStructDefType}:             _GotoState45Action,
	{_State93, AnonymousFuncExprType}:             _GotoState39Action,
	{_State100, IntegerLiteralToken}:              _GotoState25Action,
	{_State100, FloatLiteralToken}:                _GotoState22Action,
	{_State100, RuneLiteralToken}:                 _GotoState31Action,
	{_State100, StringLiteralToken}:               _GotoState32Action,
	{_State100, IdentifierToken}:                  _GotoState24Action,
	{_State100, TrueToken}:                        _GotoState35Action,
	{_State100, FalseToken}:                       _GotoState21Action,
	{_State100, StructToken}:                      _GotoState33Action,
	{_State100, FuncToken}:                        _GotoState23Action,
	{_State100, LabelDeclToken}:                   _GotoState26Action,
	{_State100, LparenToken}:                      _GotoState28Action,
	{_State100, NotToken}:                         _GotoState30Action,
	{_State100, SubToken}:                         _GotoState34Action,
	{_State100, MulToken}:                         _GotoState29Action,
	{_State100, BitNegToken}:                      _GotoState20Action,
	{_State100, BitAndToken}:                      _GotoState19Action,
	{_State100, LexErrorToken}:                    _GotoState27Action,
	{_State100, OptionalLabelDeclType}:            _GotoState111Action,
	{_State100, BlockExprType}:                    _GotoState42Action,
	{_State100, CallExprType}:                     _GotoState43Action,
	{_State100, AtomExprType}:                     _GotoState41Action,
	{_State100, LiteralType}:                      _GotoState46Action,
	{_State100, AnonymousStructExprType}:          _GotoState40Action,
	{_State100, AccessExprType}:                   _GotoState36Action,
	{_State100, PostfixUnaryExprType}:             _GotoState50Action,
	{_State100, PrefixUnaryOpType}:                _GotoState52Action,
	{_State100, PrefixUnaryExprType}:              _GotoState166Action,
	{_State100, ExplicitStructDefType}:            _GotoState45Action,
	{_State100, AnonymousFuncExprType}:            _GotoState39Action,
	{_State101, LbraceToken}:                      _GotoState104Action,
	{_State101, BlockBodyType}:                    _GotoState167Action,
	{_State102, IntegerLiteralToken}:              _GotoState25Action,
	{_State102, FloatLiteralToken}:                _GotoState22Action,
	{_State102, RuneLiteralToken}:                 _GotoState31Action,
	{_State102, StringLiteralToken}:               _GotoState32Action,
	{_State102, IdentifierToken}:                  _GotoState24Action,
	{_State102, TrueToken}:                        _GotoState35Action,
	{_State102, FalseToken}:                       _GotoState21Action,
	{_State102, InToken}:                          _GotoState168Action,
	{_State102, StructToken}:                      _GotoState33Action,
	{_State102, FuncToken}:                        _GotoState23Action,
	{_State102, LabelDeclToken}:                   _GotoState26Action,
	{_State102, LparenToken}:                      _GotoState28Action,
	{_State102, SemicolonToken}:                   _GotoState169Action,
	{_State102, NotToken}:                         _GotoState30Action,
	{_State102, SubToken}:                         _GotoState34Action,
	{_State102, MulToken}:                         _GotoState29Action,
	{_State102, BitNegToken}:                      _GotoState20Action,
	{_State102, BitAndToken}:                      _GotoState19Action,
	{_State102, LexErrorToken}:                    _GotoState27Action,
	{_State102, OptionalLabelDeclType}:            _GotoState111Action,
	{_State102, SequenceExprType}:                 _GotoState170Action,
	{_State102, BlockExprType}:                    _GotoState42Action,
	{_State102, CallExprType}:                     _GotoState43Action,
	{_State102, AtomExprType}:                     _GotoState41Action,
	{_State102, LiteralType}:                      _GotoState46Action,
	{_State102, AnonymousStructExprType}:          _GotoState40Action,
	{_State102, AccessExprType}:                   _GotoState36Action,
	{_State102, PostfixUnaryExprType}:             _GotoState50Action,
	{_State102, PrefixUnaryOpType}:                _GotoState52Action,
	{_State102, PrefixUnaryExprType}:              _GotoState51Action,
	{_State102, MulExprType}:                      _GotoState47Action,
	{_State102, AddExprType}:                      _GotoState37Action,
	{_State102, CmpExprType}:                      _GotoState44Action,
	{_State102, AndExprType}:                      _GotoState38Action,
	{_State102, OrExprType}:                       _GotoState49Action,
	{_State102, ExplicitStructDefType}:            _GotoState45Action,
	{_State102, AnonymousFuncExprType}:            _GotoState39Action,
	{_State103, IntegerLiteralToken}:              _GotoState25Action,
	{_State103, FloatLiteralToken}:                _GotoState22Action,
	{_State103, RuneLiteralToken}:                 _GotoState31Action,
	{_State103, StringLiteralToken}:               _GotoState32Action,
	{_State103, IdentifierToken}:                  _GotoState24Action,
	{_State103, TrueToken}:                        _GotoState35Action,
	{_State103, FalseToken}:                       _GotoState21Action,
	{_State103, StructToken}:                      _GotoState33Action,
	{_State103, FuncToken}:                        _GotoState23Action,
	{_State103, LabelDeclToken}:                   _GotoState26Action,
	{_State103, LparenToken}:                      _GotoState28Action,
	{_State103, NotToken}:                         _GotoState30Action,
	{_State103, SubToken}:                         _GotoState34Action,
	{_State103, MulToken}:                         _GotoState29Action,
	{_State103, BitNegToken}:                      _GotoState20Action,
	{_State103, BitAndToken}:                      _GotoState19Action,
	{_State103, LexErrorToken}:                    _GotoState27Action,
	{_State103, OptionalLabelDeclType}:            _GotoState111Action,
	{_State103, SequenceExprType}:                 _GotoState171Action,
	{_State103, BlockExprType}:                    _GotoState42Action,
	{_State103, CallExprType}:                     _GotoState43Action,
	{_State103, AtomExprType}:                     _GotoState41Action,
	{_State103, LiteralType}:                      _GotoState46Action,
	{_State103, AnonymousStructExprType}:          _GotoState40Action,
	{_State103, AccessExprType}:                   _GotoState36Action,
	{_State103, PostfixUnaryExprType}:             _GotoState50Action,
	{_State103, PrefixUnaryOpType}:                _GotoState52Action,
	{_State103, PrefixUnaryExprType}:              _GotoState51Action,
	{_State103, MulExprType}:                      _GotoState47Action,
	{_State103, AddExprType}:                      _GotoState37Action,
	{_State103, CmpExprType}:                      _GotoState44Action,
	{_State103, AndExprType}:                      _GotoState38Action,
	{_State103, OrExprType}:                       _GotoState49Action,
	{_State103, ExplicitStructDefType}:            _GotoState45Action,
	{_State103, AnonymousFuncExprType}:            _GotoState39Action,
	{_State104, StatementsType}:                   _GotoState172Action,
	{_State105, IntegerLiteralToken}:              _GotoState25Action,
	{_State105, FloatLiteralToken}:                _GotoState22Action,
	{_State105, RuneLiteralToken}:                 _GotoState31Action,
	{_State105, StringLiteralToken}:               _GotoState32Action,
	{_State105, IdentifierToken}:                  _GotoState24Action,
	{_State105, TrueToken}:                        _GotoState35Action,
	{_State105, FalseToken}:                       _GotoState21Action,
	{_State105, StructToken}:                      _GotoState33Action,
	{_State105, FuncToken}:                        _GotoState23Action,
	{_State105, LabelDeclToken}:                   _GotoState26Action,
	{_State105, LparenToken}:                      _GotoState28Action,
	{_State105, NotToken}:                         _GotoState30Action,
	{_State105, SubToken}:                         _GotoState34Action,
	{_State105, MulToken}:                         _GotoState29Action,
	{_State105, BitNegToken}:                      _GotoState20Action,
	{_State105, BitAndToken}:                      _GotoState19Action,
	{_State105, LexErrorToken}:                    _GotoState27Action,
	{_State105, OptionalLabelDeclType}:            _GotoState111Action,
	{_State105, SequenceExprType}:                 _GotoState173Action,
	{_State105, BlockExprType}:                    _GotoState42Action,
	{_State105, CallExprType}:                     _GotoState43Action,
	{_State105, AtomExprType}:                     _GotoState41Action,
	{_State105, LiteralType}:                      _GotoState46Action,
	{_State105, AnonymousStructExprType}:          _GotoState40Action,
	{_State105, AccessExprType}:                   _GotoState36Action,
	{_State105, PostfixUnaryExprType}:             _GotoState50Action,
	{_State105, PrefixUnaryOpType}:                _GotoState52Action,
	{_State105, PrefixUnaryExprType}:              _GotoState51Action,
	{_State105, MulExprType}:                      _GotoState47Action,
	{_State105, AddExprType}:                      _GotoState37Action,
	{_State105, CmpExprType}:                      _GotoState44Action,
	{_State105, AndExprType}:                      _GotoState38Action,
	{_State105, OrExprType}:                       _GotoState49Action,
	{_State105, ExplicitStructDefType}:            _GotoState45Action,
	{_State105, AnonymousFuncExprType}:            _GotoState39Action,
	{_State110, IntegerLiteralToken}:              _GotoState25Action,
	{_State110, FloatLiteralToken}:                _GotoState22Action,
	{_State110, RuneLiteralToken}:                 _GotoState31Action,
	{_State110, StringLiteralToken}:               _GotoState32Action,
	{_State110, IdentifierToken}:                  _GotoState24Action,
	{_State110, TrueToken}:                        _GotoState35Action,
	{_State110, FalseToken}:                       _GotoState21Action,
	{_State110, StructToken}:                      _GotoState33Action,
	{_State110, FuncToken}:                        _GotoState23Action,
	{_State110, LabelDeclToken}:                   _GotoState26Action,
	{_State110, LparenToken}:                      _GotoState28Action,
	{_State110, NotToken}:                         _GotoState30Action,
	{_State110, SubToken}:                         _GotoState34Action,
	{_State110, MulToken}:                         _GotoState29Action,
	{_State110, BitNegToken}:                      _GotoState20Action,
	{_State110, BitAndToken}:                      _GotoState19Action,
	{_State110, LexErrorToken}:                    _GotoState27Action,
	{_State110, OptionalLabelDeclType}:            _GotoState111Action,
	{_State110, BlockExprType}:                    _GotoState42Action,
	{_State110, CallExprType}:                     _GotoState43Action,
	{_State110, AtomExprType}:                     _GotoState41Action,
	{_State110, LiteralType}:                      _GotoState46Action,
	{_State110, AnonymousStructExprType}:          _GotoState40Action,
	{_State110, AccessExprType}:                   _GotoState36Action,
	{_State110, PostfixUnaryExprType}:             _GotoState50Action,
	{_State110, PrefixUnaryOpType}:                _GotoState52Action,
	{_State110, PrefixUnaryExprType}:              _GotoState51Action,
	{_State110, MulExprType}:                      _GotoState47Action,
	{_State110, AddExprType}:                      _GotoState37Action,
	{_State110, CmpExprType}:                      _GotoState44Action,
	{_State110, AndExprType}:                      _GotoState174Action,
	{_State110, ExplicitStructDefType}:            _GotoState45Action,
	{_State110, AnonymousFuncExprType}:            _GotoState39Action,
	{_State111, LbraceToken}:                      _GotoState104Action,
	{_State111, BlockBodyType}:                    _GotoState106Action,
	{_State113, IdentifierToken}:                  _GotoState175Action,
	{_State114, PackageToken}:                     _GotoState16Action,
	{_State114, UnsafeToken}:                      _GotoState56Action,
	{_State114, TypeToken}:                        _GotoState17Action,
	{_State114, FuncToken}:                        _GotoState18Action,
	{_State114, DefinitionType}:                   _GotoState176Action,
	{_State114, UnsafeStatementType}:              _GotoState62Action,
	{_State114, TypeDefType}:                      _GotoState61Action,
	{_State114, NamedFuncDefType}:                 _GotoState59Action,
	{_State114, PackageDefType}:                   _GotoState60Action,
	{_State116, PackageStatementsType}:            _GotoState177Action,
	{_State117, IdentifierToken}:                  _GotoState178Action,
	{_State117, GenericParameterDefType}:          _GotoState179Action,
	{_State117, GenericParameterDefsType}:         _GotoState180Action,
	{_State117, OptionalGenericParameterDefsType}: _GotoState181Action,
	{_State118, IdentifierToken}:                  _GotoState155Action,
	{_State118, StructToken}:                      _GotoState33Action,
	{_State118, EnumToken}:                        _GotoState133Action,
	{_State118, TraitToken}:                       _GotoState140Action,
	{_State118, FuncToken}:                        _GotoState134Action,
	{_State118, LparenToken}:                      _GotoState137Action,
	{_State118, QuestionToken}:                    _GotoState138Action,
	{_State118, TildeTildeToken}:                  _GotoState139Action,
	{_State118, BitNegToken}:                      _GotoState132Action,
	{_State118, BitAndToken}:                      _GotoState131Action,
	{_State118, LexErrorToken}:                    _GotoState136Action,
	{_State118, AtomTypeType}:                     _GotoState141Action,
	{_State118, TraitableTypeType}:                _GotoState152Action,
	{_State118, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State118, ValueTypeType}:                    _GotoState182Action,
	{_State118, ImplicitStructDefType}:            _GotoState148Action,
	{_State118, ExplicitStructDefType}:            _GotoState144Action,
	{_State118, ImplicitEnumDefType}:              _GotoState147Action,
	{_State118, ExplicitEnumDefType}:              _GotoState142Action,
	{_State118, TraitDefType}:                     _GotoState151Action,
	{_State118, FuncTypeType}:                     _GotoState146Action,
	{_State119, IdentifierToken}:                  _GotoState155Action,
	{_State119, StructToken}:                      _GotoState33Action,
	{_State119, EnumToken}:                        _GotoState133Action,
	{_State119, TraitToken}:                       _GotoState140Action,
	{_State119, FuncToken}:                        _GotoState134Action,
	{_State119, LparenToken}:                      _GotoState137Action,
	{_State119, QuestionToken}:                    _GotoState138Action,
	{_State119, TildeTildeToken}:                  _GotoState139Action,
	{_State119, BitNegToken}:                      _GotoState132Action,
	{_State119, BitAndToken}:                      _GotoState131Action,
	{_State119, LexErrorToken}:                    _GotoState136Action,
	{_State119, AtomTypeType}:                     _GotoState141Action,
	{_State119, TraitableTypeType}:                _GotoState152Action,
	{_State119, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State119, ValueTypeType}:                    _GotoState183Action,
	{_State119, ImplicitStructDefType}:            _GotoState148Action,
	{_State119, ExplicitStructDefType}:            _GotoState144Action,
	{_State119, ImplicitEnumDefType}:              _GotoState147Action,
	{_State119, ExplicitEnumDefType}:              _GotoState142Action,
	{_State119, TraitDefType}:                     _GotoState151Action,
	{_State119, FuncTypeType}:                     _GotoState146Action,
	{_State120, IdentifierToken}:                  _GotoState155Action,
	{_State120, StructToken}:                      _GotoState33Action,
	{_State120, EnumToken}:                        _GotoState133Action,
	{_State120, TraitToken}:                       _GotoState140Action,
	{_State120, FuncToken}:                        _GotoState134Action,
	{_State120, LparenToken}:                      _GotoState137Action,
	{_State120, QuestionToken}:                    _GotoState138Action,
	{_State120, DotdotdotToken}:                   _GotoState184Action,
	{_State120, TildeTildeToken}:                  _GotoState139Action,
	{_State120, BitNegToken}:                      _GotoState132Action,
	{_State120, BitAndToken}:                      _GotoState131Action,
	{_State120, LexErrorToken}:                    _GotoState136Action,
	{_State120, AtomTypeType}:                     _GotoState141Action,
	{_State120, TraitableTypeType}:                _GotoState152Action,
	{_State120, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State120, ValueTypeType}:                    _GotoState185Action,
	{_State120, ImplicitStructDefType}:            _GotoState148Action,
	{_State120, ExplicitStructDefType}:            _GotoState144Action,
	{_State120, ImplicitEnumDefType}:              _GotoState147Action,
	{_State120, ExplicitEnumDefType}:              _GotoState142Action,
	{_State120, TraitDefType}:                     _GotoState151Action,
	{_State120, FuncTypeType}:                     _GotoState146Action,
	{_State121, RparenToken}:                      _GotoState186Action,
	{_State122, DollarLbracketToken}:              _GotoState117Action,
	{_State122, OptionalGenericParametersType}:    _GotoState187Action,
	{_State123, RparenToken}:                      _GotoState188Action,
	{_State125, CommaToken}:                       _GotoState189Action,
	{_State126, IntegerLiteralToken}:              _GotoState25Action,
	{_State126, FloatLiteralToken}:                _GotoState22Action,
	{_State126, RuneLiteralToken}:                 _GotoState31Action,
	{_State126, StringLiteralToken}:               _GotoState32Action,
	{_State126, IdentifierToken}:                  _GotoState24Action,
	{_State126, TrueToken}:                        _GotoState35Action,
	{_State126, FalseToken}:                       _GotoState21Action,
	{_State126, StructToken}:                      _GotoState33Action,
	{_State126, FuncToken}:                        _GotoState23Action,
	{_State126, LabelDeclToken}:                   _GotoState26Action,
	{_State126, LparenToken}:                      _GotoState28Action,
	{_State126, NotToken}:                         _GotoState30Action,
	{_State126, SubToken}:                         _GotoState34Action,
	{_State126, MulToken}:                         _GotoState29Action,
	{_State126, BitNegToken}:                      _GotoState20Action,
	{_State126, BitAndToken}:                      _GotoState19Action,
	{_State126, LexErrorToken}:                    _GotoState27Action,
	{_State126, ExpressionType}:                   _GotoState190Action,
	{_State126, OptionalLabelDeclType}:            _GotoState48Action,
	{_State126, SequenceExprType}:                 _GotoState53Action,
	{_State126, BlockExprType}:                    _GotoState42Action,
	{_State126, CallExprType}:                     _GotoState43Action,
	{_State126, AtomExprType}:                     _GotoState41Action,
	{_State126, LiteralType}:                      _GotoState46Action,
	{_State126, AnonymousStructExprType}:          _GotoState40Action,
	{_State126, AccessExprType}:                   _GotoState36Action,
	{_State126, PostfixUnaryExprType}:             _GotoState50Action,
	{_State126, PrefixUnaryOpType}:                _GotoState52Action,
	{_State126, PrefixUnaryExprType}:              _GotoState51Action,
	{_State126, MulExprType}:                      _GotoState47Action,
	{_State126, AddExprType}:                      _GotoState37Action,
	{_State126, CmpExprType}:                      _GotoState44Action,
	{_State126, AndExprType}:                      _GotoState38Action,
	{_State126, OrExprType}:                       _GotoState49Action,
	{_State126, ExplicitStructDefType}:            _GotoState45Action,
	{_State126, AnonymousFuncExprType}:            _GotoState39Action,
	{_State127, IntegerLiteralToken}:              _GotoState25Action,
	{_State127, FloatLiteralToken}:                _GotoState22Action,
	{_State127, RuneLiteralToken}:                 _GotoState31Action,
	{_State127, StringLiteralToken}:               _GotoState32Action,
	{_State127, IdentifierToken}:                  _GotoState68Action,
	{_State127, TrueToken}:                        _GotoState35Action,
	{_State127, FalseToken}:                       _GotoState21Action,
	{_State127, StructToken}:                      _GotoState33Action,
	{_State127, FuncToken}:                        _GotoState23Action,
	{_State127, LabelDeclToken}:                   _GotoState26Action,
	{_State127, LparenToken}:                      _GotoState28Action,
	{_State127, NotToken}:                         _GotoState30Action,
	{_State127, SubToken}:                         _GotoState34Action,
	{_State127, MulToken}:                         _GotoState29Action,
	{_State127, BitNegToken}:                      _GotoState20Action,
	{_State127, BitAndToken}:                      _GotoState19Action,
	{_State127, LexErrorToken}:                    _GotoState27Action,
	{_State127, ExpressionType}:                   _GotoState72Action,
	{_State127, OptionalLabelDeclType}:            _GotoState48Action,
	{_State127, SequenceExprType}:                 _GotoState53Action,
	{_State127, BlockExprType}:                    _GotoState42Action,
	{_State127, CallExprType}:                     _GotoState43Action,
	{_State127, ArgumentType}:                     _GotoState191Action,
	{_State127, ColonExpressionsType}:             _GotoState71Action,
	{_State127, OptionalExpressionType}:           _GotoState73Action,
	{_State127, AtomExprType}:                     _GotoState41Action,
	{_State127, LiteralType}:                      _GotoState46Action,
	{_State127, AnonymousStructExprType}:          _GotoState40Action,
	{_State127, AccessExprType}:                   _GotoState36Action,
	{_State127, PostfixUnaryExprType}:             _GotoState50Action,
	{_State127, PrefixUnaryOpType}:                _GotoState52Action,
	{_State127, PrefixUnaryExprType}:              _GotoState51Action,
	{_State127, MulExprType}:                      _GotoState47Action,
	{_State127, AddExprType}:                      _GotoState37Action,
	{_State127, CmpExprType}:                      _GotoState44Action,
	{_State127, AndExprType}:                      _GotoState38Action,
	{_State127, OrExprType}:                       _GotoState49Action,
	{_State127, ExplicitStructDefType}:            _GotoState45Action,
	{_State127, AnonymousFuncExprType}:            _GotoState39Action,
	{_State129, IntegerLiteralToken}:              _GotoState25Action,
	{_State129, FloatLiteralToken}:                _GotoState22Action,
	{_State129, RuneLiteralToken}:                 _GotoState31Action,
	{_State129, StringLiteralToken}:               _GotoState32Action,
	{_State129, IdentifierToken}:                  _GotoState24Action,
	{_State129, TrueToken}:                        _GotoState35Action,
	{_State129, FalseToken}:                       _GotoState21Action,
	{_State129, StructToken}:                      _GotoState33Action,
	{_State129, FuncToken}:                        _GotoState23Action,
	{_State129, LabelDeclToken}:                   _GotoState26Action,
	{_State129, LparenToken}:                      _GotoState28Action,
	{_State129, NotToken}:                         _GotoState30Action,
	{_State129, SubToken}:                         _GotoState34Action,
	{_State129, MulToken}:                         _GotoState29Action,
	{_State129, BitNegToken}:                      _GotoState20Action,
	{_State129, BitAndToken}:                      _GotoState19Action,
	{_State129, LexErrorToken}:                    _GotoState27Action,
	{_State129, ExpressionType}:                   _GotoState192Action,
	{_State129, OptionalLabelDeclType}:            _GotoState48Action,
	{_State129, SequenceExprType}:                 _GotoState53Action,
	{_State129, BlockExprType}:                    _GotoState42Action,
	{_State129, CallExprType}:                     _GotoState43Action,
	{_State129, OptionalExpressionType}:           _GotoState193Action,
	{_State129, AtomExprType}:                     _GotoState41Action,
	{_State129, LiteralType}:                      _GotoState46Action,
	{_State129, AnonymousStructExprType}:          _GotoState40Action,
	{_State129, AccessExprType}:                   _GotoState36Action,
	{_State129, PostfixUnaryExprType}:             _GotoState50Action,
	{_State129, PrefixUnaryOpType}:                _GotoState52Action,
	{_State129, PrefixUnaryExprType}:              _GotoState51Action,
	{_State129, MulExprType}:                      _GotoState47Action,
	{_State129, AddExprType}:                      _GotoState37Action,
	{_State129, CmpExprType}:                      _GotoState44Action,
	{_State129, AndExprType}:                      _GotoState38Action,
	{_State129, OrExprType}:                       _GotoState49Action,
	{_State129, ExplicitStructDefType}:            _GotoState45Action,
	{_State129, AnonymousFuncExprType}:            _GotoState39Action,
	{_State130, IntegerLiteralToken}:              _GotoState25Action,
	{_State130, FloatLiteralToken}:                _GotoState22Action,
	{_State130, RuneLiteralToken}:                 _GotoState31Action,
	{_State130, StringLiteralToken}:               _GotoState32Action,
	{_State130, IdentifierToken}:                  _GotoState24Action,
	{_State130, TrueToken}:                        _GotoState35Action,
	{_State130, FalseToken}:                       _GotoState21Action,
	{_State130, StructToken}:                      _GotoState33Action,
	{_State130, FuncToken}:                        _GotoState23Action,
	{_State130, LabelDeclToken}:                   _GotoState26Action,
	{_State130, LparenToken}:                      _GotoState28Action,
	{_State130, NotToken}:                         _GotoState30Action,
	{_State130, SubToken}:                         _GotoState34Action,
	{_State130, MulToken}:                         _GotoState29Action,
	{_State130, BitNegToken}:                      _GotoState20Action,
	{_State130, BitAndToken}:                      _GotoState19Action,
	{_State130, LexErrorToken}:                    _GotoState27Action,
	{_State130, ExpressionType}:                   _GotoState192Action,
	{_State130, OptionalLabelDeclType}:            _GotoState48Action,
	{_State130, SequenceExprType}:                 _GotoState53Action,
	{_State130, BlockExprType}:                    _GotoState42Action,
	{_State130, CallExprType}:                     _GotoState43Action,
	{_State130, OptionalExpressionType}:           _GotoState194Action,
	{_State130, AtomExprType}:                     _GotoState41Action,
	{_State130, LiteralType}:                      _GotoState46Action,
	{_State130, AnonymousStructExprType}:          _GotoState40Action,
	{_State130, AccessExprType}:                   _GotoState36Action,
	{_State130, PostfixUnaryExprType}:             _GotoState50Action,
	{_State130, PrefixUnaryOpType}:                _GotoState52Action,
	{_State130, PrefixUnaryExprType}:              _GotoState51Action,
	{_State130, MulExprType}:                      _GotoState47Action,
	{_State130, AddExprType}:                      _GotoState37Action,
	{_State130, CmpExprType}:                      _GotoState44Action,
	{_State130, AndExprType}:                      _GotoState38Action,
	{_State130, OrExprType}:                       _GotoState49Action,
	{_State130, ExplicitStructDefType}:            _GotoState45Action,
	{_State130, AnonymousFuncExprType}:            _GotoState39Action,
	{_State131, IdentifierToken}:                  _GotoState155Action,
	{_State131, StructToken}:                      _GotoState33Action,
	{_State131, EnumToken}:                        _GotoState133Action,
	{_State131, TraitToken}:                       _GotoState140Action,
	{_State131, LparenToken}:                      _GotoState137Action,
	{_State131, TildeTildeToken}:                  _GotoState139Action,
	{_State131, BitNegToken}:                      _GotoState132Action,
	{_State131, LexErrorToken}:                    _GotoState136Action,
	{_State131, AtomTypeType}:                     _GotoState141Action,
	{_State131, TraitableTypeType}:                _GotoState152Action,
	{_State131, TraitAlgebraTypeType}:             _GotoState195Action,
	{_State131, ImplicitStructDefType}:            _GotoState148Action,
	{_State131, ExplicitStructDefType}:            _GotoState144Action,
	{_State131, ImplicitEnumDefType}:              _GotoState147Action,
	{_State131, ExplicitEnumDefType}:              _GotoState142Action,
	{_State131, TraitDefType}:                     _GotoState151Action,
	{_State132, IdentifierToken}:                  _GotoState155Action,
	{_State132, StructToken}:                      _GotoState33Action,
	{_State132, EnumToken}:                        _GotoState133Action,
	{_State132, TraitToken}:                       _GotoState140Action,
	{_State132, LparenToken}:                      _GotoState137Action,
	{_State132, LexErrorToken}:                    _GotoState136Action,
	{_State132, AtomTypeType}:                     _GotoState196Action,
	{_State132, ImplicitStructDefType}:            _GotoState148Action,
	{_State132, ExplicitStructDefType}:            _GotoState144Action,
	{_State132, ImplicitEnumDefType}:              _GotoState147Action,
	{_State132, ExplicitEnumDefType}:              _GotoState142Action,
	{_State132, TraitDefType}:                     _GotoState151Action,
	{_State133, LparenToken}:                      _GotoState197Action,
	{_State134, LparenToken}:                      _GotoState198Action,
	{_State135, IdentifierToken}:                  _GotoState155Action,
	{_State135, StructToken}:                      _GotoState33Action,
	{_State135, EnumToken}:                        _GotoState133Action,
	{_State135, TraitToken}:                       _GotoState140Action,
	{_State135, FuncToken}:                        _GotoState134Action,
	{_State135, LparenToken}:                      _GotoState137Action,
	{_State135, QuestionToken}:                    _GotoState138Action,
	{_State135, DollarLbracketToken}:              _GotoState75Action,
	{_State135, TildeTildeToken}:                  _GotoState139Action,
	{_State135, BitNegToken}:                      _GotoState132Action,
	{_State135, BitAndToken}:                      _GotoState131Action,
	{_State135, LexErrorToken}:                    _GotoState136Action,
	{_State135, OptionalGenericBindingType}:       _GotoState199Action,
	{_State135, AtomTypeType}:                     _GotoState141Action,
	{_State135, TraitableTypeType}:                _GotoState152Action,
	{_State135, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State135, ValueTypeType}:                    _GotoState200Action,
	{_State135, ImplicitStructDefType}:            _GotoState148Action,
	{_State135, ExplicitStructDefType}:            _GotoState144Action,
	{_State135, ImplicitEnumDefType}:              _GotoState147Action,
	{_State135, ExplicitEnumDefType}:              _GotoState142Action,
	{_State135, TraitDefType}:                     _GotoState151Action,
	{_State135, FuncTypeType}:                     _GotoState146Action,
	{_State137, IdentifierToken}:                  _GotoState135Action,
	{_State137, UnsafeToken}:                      _GotoState56Action,
	{_State137, StructToken}:                      _GotoState33Action,
	{_State137, EnumToken}:                        _GotoState133Action,
	{_State137, TraitToken}:                       _GotoState140Action,
	{_State137, FuncToken}:                        _GotoState134Action,
	{_State137, LparenToken}:                      _GotoState137Action,
	{_State137, QuestionToken}:                    _GotoState138Action,
	{_State137, TildeTildeToken}:                  _GotoState139Action,
	{_State137, BitNegToken}:                      _GotoState132Action,
	{_State137, BitAndToken}:                      _GotoState131Action,
	{_State137, LexErrorToken}:                    _GotoState136Action,
	{_State137, UnsafeStatementType}:              _GotoState153Action,
	{_State137, AtomTypeType}:                     _GotoState141Action,
	{_State137, TraitableTypeType}:                _GotoState152Action,
	{_State137, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State137, ValueTypeType}:                    _GotoState154Action,
	{_State137, FieldDefType}:                     _GotoState202Action,
	{_State137, ImplicitFieldDefsType}:            _GotoState204Action,
	{_State137, OptionalImplicitFieldDefsType}:    _GotoState205Action,
	{_State137, ImplicitStructDefType}:            _GotoState148Action,
	{_State137, ExplicitStructDefType}:            _GotoState144Action,
	{_State137, EnumValueDefType}:                 _GotoState201Action,
	{_State137, ImplicitEnumValueDefsType}:        _GotoState203Action,
	{_State137, ImplicitEnumDefType}:              _GotoState147Action,
	{_State137, ExplicitEnumDefType}:              _GotoState142Action,
	{_State137, TraitDefType}:                     _GotoState151Action,
	{_State137, FuncTypeType}:                     _GotoState146Action,
	{_State139, IdentifierToken}:                  _GotoState155Action,
	{_State139, StructToken}:                      _GotoState33Action,
	{_State139, EnumToken}:                        _GotoState133Action,
	{_State139, TraitToken}:                       _GotoState140Action,
	{_State139, LparenToken}:                      _GotoState137Action,
	{_State139, LexErrorToken}:                    _GotoState136Action,
	{_State139, AtomTypeType}:                     _GotoState206Action,
	{_State139, ImplicitStructDefType}:            _GotoState148Action,
	{_State139, ExplicitStructDefType}:            _GotoState144Action,
	{_State139, ImplicitEnumDefType}:              _GotoState147Action,
	{_State139, ExplicitEnumDefType}:              _GotoState142Action,
	{_State139, TraitDefType}:                     _GotoState151Action,
	{_State140, LparenToken}:                      _GotoState207Action,
	{_State143, NewlinesToken}:                    _GotoState209Action,
	{_State143, CommaToken}:                       _GotoState208Action,
	{_State149, RparenToken}:                      _GotoState210Action,
	{_State150, AddToken}:                         _GotoState211Action,
	{_State150, SubToken}:                         _GotoState213Action,
	{_State150, MulToken}:                         _GotoState212Action,
	{_State155, DollarLbracketToken}:              _GotoState75Action,
	{_State155, OptionalGenericBindingType}:       _GotoState199Action,
	{_State156, CommaToken}:                       _GotoState214Action,
	{_State157, RbracketToken}:                    _GotoState215Action,
	{_State160, RbracketToken}:                    _GotoState216Action,
	{_State161, IntegerLiteralToken}:              _GotoState25Action,
	{_State161, FloatLiteralToken}:                _GotoState22Action,
	{_State161, RuneLiteralToken}:                 _GotoState31Action,
	{_State161, StringLiteralToken}:               _GotoState32Action,
	{_State161, IdentifierToken}:                  _GotoState68Action,
	{_State161, TrueToken}:                        _GotoState35Action,
	{_State161, FalseToken}:                       _GotoState21Action,
	{_State161, StructToken}:                      _GotoState33Action,
	{_State161, FuncToken}:                        _GotoState23Action,
	{_State161, LabelDeclToken}:                   _GotoState26Action,
	{_State161, LparenToken}:                      _GotoState28Action,
	{_State161, NotToken}:                         _GotoState30Action,
	{_State161, SubToken}:                         _GotoState34Action,
	{_State161, MulToken}:                         _GotoState29Action,
	{_State161, BitNegToken}:                      _GotoState20Action,
	{_State161, BitAndToken}:                      _GotoState19Action,
	{_State161, LexErrorToken}:                    _GotoState27Action,
	{_State161, ExpressionType}:                   _GotoState72Action,
	{_State161, OptionalLabelDeclType}:            _GotoState48Action,
	{_State161, SequenceExprType}:                 _GotoState53Action,
	{_State161, BlockExprType}:                    _GotoState42Action,
	{_State161, CallExprType}:                     _GotoState43Action,
	{_State161, OptionalArgumentsType}:            _GotoState218Action,
	{_State161, ArgumentsType}:                    _GotoState217Action,
	{_State161, ArgumentType}:                     _GotoState69Action,
	{_State161, ColonExpressionsType}:             _GotoState71Action,
	{_State161, OptionalExpressionType}:           _GotoState73Action,
	{_State161, AtomExprType}:                     _GotoState41Action,
	{_State161, LiteralType}:                      _GotoState46Action,
	{_State161, AnonymousStructExprType}:          _GotoState40Action,
	{_State161, AccessExprType}:                   _GotoState36Action,
	{_State161, PostfixUnaryExprType}:             _GotoState50Action,
	{_State161, PrefixUnaryOpType}:                _GotoState52Action,
	{_State161, PrefixUnaryExprType}:              _GotoState51Action,
	{_State161, MulExprType}:                      _GotoState47Action,
	{_State161, AddExprType}:                      _GotoState37Action,
	{_State161, CmpExprType}:                      _GotoState44Action,
	{_State161, AndExprType}:                      _GotoState38Action,
	{_State161, OrExprType}:                       _GotoState49Action,
	{_State161, ExplicitStructDefType}:            _GotoState45Action,
	{_State161, AnonymousFuncExprType}:            _GotoState39Action,
	{_State162, MulToken}:                         _GotoState99Action,
	{_State162, DivToken}:                         _GotoState97Action,
	{_State162, ModToken}:                         _GotoState98Action,
	{_State162, BitAndToken}:                      _GotoState94Action,
	{_State162, BitLshiftToken}:                   _GotoState95Action,
	{_State162, BitRshiftToken}:                   _GotoState96Action,
	{_State162, MulOpType}:                        _GotoState100Action,
	{_State163, EqualToken}:                       _GotoState86Action,
	{_State163, NotEqualToken}:                    _GotoState91Action,
	{_State163, LessToken}:                        _GotoState89Action,
	{_State163, LessOrEqualToken}:                 _GotoState90Action,
	{_State163, GreaterToken}:                     _GotoState87Action,
	{_State163, GreaterOrEqualToken}:              _GotoState88Action,
	{_State163, CmpOpType}:                        _GotoState92Action,
	{_State164, AddToken}:                         _GotoState80Action,
	{_State164, SubToken}:                         _GotoState83Action,
	{_State164, BitXorToken}:                      _GotoState82Action,
	{_State164, BitOrToken}:                       _GotoState81Action,
	{_State164, AddOpType}:                        _GotoState84Action,
	{_State165, RparenToken}:                      _GotoState219Action,
	{_State165, CommaToken}:                       _GotoState127Action,
	{_State167, ForToken}:                         _GotoState220Action,
	{_State168, IntegerLiteralToken}:              _GotoState25Action,
	{_State168, FloatLiteralToken}:                _GotoState22Action,
	{_State168, RuneLiteralToken}:                 _GotoState31Action,
	{_State168, StringLiteralToken}:               _GotoState32Action,
	{_State168, IdentifierToken}:                  _GotoState24Action,
	{_State168, TrueToken}:                        _GotoState35Action,
	{_State168, FalseToken}:                       _GotoState21Action,
	{_State168, StructToken}:                      _GotoState33Action,
	{_State168, FuncToken}:                        _GotoState23Action,
	{_State168, LabelDeclToken}:                   _GotoState26Action,
	{_State168, LparenToken}:                      _GotoState28Action,
	{_State168, NotToken}:                         _GotoState30Action,
	{_State168, SubToken}:                         _GotoState34Action,
	{_State168, MulToken}:                         _GotoState29Action,
	{_State168, BitNegToken}:                      _GotoState20Action,
	{_State168, BitAndToken}:                      _GotoState19Action,
	{_State168, LexErrorToken}:                    _GotoState27Action,
	{_State168, OptionalLabelDeclType}:            _GotoState111Action,
	{_State168, SequenceExprType}:                 _GotoState221Action,
	{_State168, BlockExprType}:                    _GotoState42Action,
	{_State168, CallExprType}:                     _GotoState43Action,
	{_State168, AtomExprType}:                     _GotoState41Action,
	{_State168, LiteralType}:                      _GotoState46Action,
	{_State168, AnonymousStructExprType}:          _GotoState40Action,
	{_State168, AccessExprType}:                   _GotoState36Action,
	{_State168, PostfixUnaryExprType}:             _GotoState50Action,
	{_State168, PrefixUnaryOpType}:                _GotoState52Action,
	{_State168, PrefixUnaryExprType}:              _GotoState51Action,
	{_State168, MulExprType}:                      _GotoState47Action,
	{_State168, AddExprType}:                      _GotoState37Action,
	{_State168, CmpExprType}:                      _GotoState44Action,
	{_State168, AndExprType}:                      _GotoState38Action,
	{_State168, OrExprType}:                       _GotoState49Action,
	{_State168, ExplicitStructDefType}:            _GotoState45Action,
	{_State168, AnonymousFuncExprType}:            _GotoState39Action,
	{_State169, IntegerLiteralToken}:              _GotoState25Action,
	{_State169, FloatLiteralToken}:                _GotoState22Action,
	{_State169, RuneLiteralToken}:                 _GotoState31Action,
	{_State169, StringLiteralToken}:               _GotoState32Action,
	{_State169, IdentifierToken}:                  _GotoState24Action,
	{_State169, TrueToken}:                        _GotoState35Action,
	{_State169, FalseToken}:                       _GotoState21Action,
	{_State169, StructToken}:                      _GotoState33Action,
	{_State169, FuncToken}:                        _GotoState23Action,
	{_State169, LabelDeclToken}:                   _GotoState26Action,
	{_State169, LparenToken}:                      _GotoState28Action,
	{_State169, NotToken}:                         _GotoState30Action,
	{_State169, SubToken}:                         _GotoState34Action,
	{_State169, MulToken}:                         _GotoState29Action,
	{_State169, BitNegToken}:                      _GotoState20Action,
	{_State169, BitAndToken}:                      _GotoState19Action,
	{_State169, LexErrorToken}:                    _GotoState27Action,
	{_State169, OptionalLabelDeclType}:            _GotoState111Action,
	{_State169, OptionalSequenceExprType}:         _GotoState222Action,
	{_State169, SequenceExprType}:                 _GotoState223Action,
	{_State169, BlockExprType}:                    _GotoState42Action,
	{_State169, CallExprType}:                     _GotoState43Action,
	{_State169, AtomExprType}:                     _GotoState41Action,
	{_State169, LiteralType}:                      _GotoState46Action,
	{_State169, AnonymousStructExprType}:          _GotoState40Action,
	{_State169, AccessExprType}:                   _GotoState36Action,
	{_State169, PostfixUnaryExprType}:             _GotoState50Action,
	{_State169, PrefixUnaryOpType}:                _GotoState52Action,
	{_State169, PrefixUnaryExprType}:              _GotoState51Action,
	{_State169, MulExprType}:                      _GotoState47Action,
	{_State169, AddExprType}:                      _GotoState37Action,
	{_State169, CmpExprType}:                      _GotoState44Action,
	{_State169, AndExprType}:                      _GotoState38Action,
	{_State169, OrExprType}:                       _GotoState49Action,
	{_State169, ExplicitStructDefType}:            _GotoState45Action,
	{_State169, AnonymousFuncExprType}:            _GotoState39Action,
	{_State170, DoToken}:                          _GotoState224Action,
	{_State171, LbraceToken}:                      _GotoState104Action,
	{_State171, BlockBodyType}:                    _GotoState225Action,
	{_State172, IntegerLiteralToken}:              _GotoState25Action,
	{_State172, FloatLiteralToken}:                _GotoState22Action,
	{_State172, RuneLiteralToken}:                 _GotoState31Action,
	{_State172, StringLiteralToken}:               _GotoState32Action,
	{_State172, IdentifierToken}:                  _GotoState24Action,
	{_State172, TrueToken}:                        _GotoState35Action,
	{_State172, FalseToken}:                       _GotoState21Action,
	{_State172, ReturnToken}:                      _GotoState231Action,
	{_State172, BreakToken}:                       _GotoState227Action,
	{_State172, ContinueToken}:                    _GotoState228Action,
	{_State172, UnsafeToken}:                      _GotoState56Action,
	{_State172, StructToken}:                      _GotoState33Action,
	{_State172, FuncToken}:                        _GotoState23Action,
	{_State172, AsyncToken}:                       _GotoState226Action,
	{_State172, DeferToken}:                       _GotoState229Action,
	{_State172, LabelDeclToken}:                   _GotoState26Action,
	{_State172, RbraceToken}:                      _GotoState230Action,
	{_State172, LparenToken}:                      _GotoState28Action,
	{_State172, NotToken}:                         _GotoState30Action,
	{_State172, SubToken}:                         _GotoState34Action,
	{_State172, MulToken}:                         _GotoState29Action,
	{_State172, BitNegToken}:                      _GotoState20Action,
	{_State172, BitAndToken}:                      _GotoState19Action,
	{_State172, LexErrorToken}:                    _GotoState27Action,
	{_State172, ExpressionType}:                   _GotoState233Action,
	{_State172, OptionalLabelDeclType}:            _GotoState48Action,
	{_State172, SequenceExprType}:                 _GotoState53Action,
	{_State172, BlockExprType}:                    _GotoState42Action,
	{_State172, StatementType}:                    _GotoState237Action,
	{_State172, StatementBodyType}:                _GotoState238Action,
	{_State172, UnsafeStatementType}:              _GotoState239Action,
	{_State172, JumpStatementType}:                _GotoState235Action,
	{_State172, JumpTypeType}:                     _GotoState236Action,
	{_State172, ExpressionsType}:                  _GotoState234Action,
	{_State172, CallExprType}:                     _GotoState43Action,
	{_State172, AtomExprType}:                     _GotoState41Action,
	{_State172, LiteralType}:                      _GotoState46Action,
	{_State172, AnonymousStructExprType}:          _GotoState40Action,
	{_State172, AccessExprType}:                   _GotoState232Action,
	{_State172, PostfixUnaryExprType}:             _GotoState50Action,
	{_State172, PrefixUnaryOpType}:                _GotoState52Action,
	{_State172, PrefixUnaryExprType}:              _GotoState51Action,
	{_State172, MulExprType}:                      _GotoState47Action,
	{_State172, AddExprType}:                      _GotoState37Action,
	{_State172, CmpExprType}:                      _GotoState44Action,
	{_State172, AndExprType}:                      _GotoState38Action,
	{_State172, OrExprType}:                       _GotoState49Action,
	{_State172, ExplicitStructDefType}:            _GotoState45Action,
	{_State172, AnonymousFuncExprType}:            _GotoState39Action,
	{_State173, LbraceToken}:                      _GotoState240Action,
	{_State174, AndToken}:                         _GotoState85Action,
	{_State175, GreaterToken}:                     _GotoState241Action,
	{_State177, UnsafeToken}:                      _GotoState56Action,
	{_State177, RparenToken}:                      _GotoState242Action,
	{_State177, UnsafeStatementType}:              _GotoState245Action,
	{_State177, PackageStatementBodyType}:         _GotoState244Action,
	{_State177, PackageStatementType}:             _GotoState243Action,
	{_State178, IdentifierToken}:                  _GotoState155Action,
	{_State178, StructToken}:                      _GotoState33Action,
	{_State178, EnumToken}:                        _GotoState133Action,
	{_State178, TraitToken}:                       _GotoState140Action,
	{_State178, FuncToken}:                        _GotoState134Action,
	{_State178, LparenToken}:                      _GotoState137Action,
	{_State178, QuestionToken}:                    _GotoState138Action,
	{_State178, TildeTildeToken}:                  _GotoState139Action,
	{_State178, BitNegToken}:                      _GotoState132Action,
	{_State178, BitAndToken}:                      _GotoState131Action,
	{_State178, LexErrorToken}:                    _GotoState136Action,
	{_State178, AtomTypeType}:                     _GotoState141Action,
	{_State178, TraitableTypeType}:                _GotoState152Action,
	{_State178, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State178, ValueTypeType}:                    _GotoState246Action,
	{_State178, ImplicitStructDefType}:            _GotoState148Action,
	{_State178, ExplicitStructDefType}:            _GotoState144Action,
	{_State178, ImplicitEnumDefType}:              _GotoState147Action,
	{_State178, ExplicitEnumDefType}:              _GotoState142Action,
	{_State178, TraitDefType}:                     _GotoState151Action,
	{_State178, FuncTypeType}:                     _GotoState146Action,
	{_State180, CommaToken}:                       _GotoState247Action,
	{_State181, RbracketToken}:                    _GotoState248Action,
	{_State183, ImplementsToken}:                  _GotoState249Action,
	{_State184, IdentifierToken}:                  _GotoState155Action,
	{_State184, StructToken}:                      _GotoState33Action,
	{_State184, EnumToken}:                        _GotoState133Action,
	{_State184, TraitToken}:                       _GotoState140Action,
	{_State184, FuncToken}:                        _GotoState134Action,
	{_State184, LparenToken}:                      _GotoState137Action,
	{_State184, QuestionToken}:                    _GotoState138Action,
	{_State184, TildeTildeToken}:                  _GotoState139Action,
	{_State184, BitNegToken}:                      _GotoState132Action,
	{_State184, BitAndToken}:                      _GotoState131Action,
	{_State184, LexErrorToken}:                    _GotoState136Action,
	{_State184, AtomTypeType}:                     _GotoState141Action,
	{_State184, TraitableTypeType}:                _GotoState152Action,
	{_State184, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State184, ValueTypeType}:                    _GotoState250Action,
	{_State184, ImplicitStructDefType}:            _GotoState148Action,
	{_State184, ExplicitStructDefType}:            _GotoState144Action,
	{_State184, ImplicitEnumDefType}:              _GotoState147Action,
	{_State184, ExplicitEnumDefType}:              _GotoState142Action,
	{_State184, TraitDefType}:                     _GotoState151Action,
	{_State184, FuncTypeType}:                     _GotoState146Action,
	{_State187, LparenToken}:                      _GotoState251Action,
	{_State188, IdentifierToken}:                  _GotoState155Action,
	{_State188, StructToken}:                      _GotoState33Action,
	{_State188, EnumToken}:                        _GotoState133Action,
	{_State188, TraitToken}:                       _GotoState140Action,
	{_State188, FuncToken}:                        _GotoState134Action,
	{_State188, LparenToken}:                      _GotoState137Action,
	{_State188, QuestionToken}:                    _GotoState138Action,
	{_State188, TildeTildeToken}:                  _GotoState139Action,
	{_State188, BitNegToken}:                      _GotoState132Action,
	{_State188, BitAndToken}:                      _GotoState131Action,
	{_State188, LexErrorToken}:                    _GotoState136Action,
	{_State188, AtomTypeType}:                     _GotoState141Action,
	{_State188, TraitableTypeType}:                _GotoState152Action,
	{_State188, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State188, ValueTypeType}:                    _GotoState253Action,
	{_State188, ImplicitStructDefType}:            _GotoState148Action,
	{_State188, ExplicitStructDefType}:            _GotoState144Action,
	{_State188, ImplicitEnumDefType}:              _GotoState147Action,
	{_State188, ExplicitEnumDefType}:              _GotoState142Action,
	{_State188, TraitDefType}:                     _GotoState151Action,
	{_State188, ReturnTypeType}:                   _GotoState252Action,
	{_State188, FuncTypeType}:                     _GotoState146Action,
	{_State189, IdentifierToken}:                  _GotoState120Action,
	{_State189, ParameterDefType}:                 _GotoState254Action,
	{_State195, AddToken}:                         _GotoState211Action,
	{_State195, SubToken}:                         _GotoState213Action,
	{_State195, MulToken}:                         _GotoState212Action,
	{_State197, IdentifierToken}:                  _GotoState135Action,
	{_State197, UnsafeToken}:                      _GotoState56Action,
	{_State197, StructToken}:                      _GotoState33Action,
	{_State197, EnumToken}:                        _GotoState133Action,
	{_State197, TraitToken}:                       _GotoState140Action,
	{_State197, FuncToken}:                        _GotoState134Action,
	{_State197, LparenToken}:                      _GotoState137Action,
	{_State197, QuestionToken}:                    _GotoState138Action,
	{_State197, TildeTildeToken}:                  _GotoState139Action,
	{_State197, BitNegToken}:                      _GotoState132Action,
	{_State197, BitAndToken}:                      _GotoState131Action,
	{_State197, LexErrorToken}:                    _GotoState136Action,
	{_State197, UnsafeStatementType}:              _GotoState153Action,
	{_State197, AtomTypeType}:                     _GotoState141Action,
	{_State197, TraitableTypeType}:                _GotoState152Action,
	{_State197, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State197, ValueTypeType}:                    _GotoState154Action,
	{_State197, FieldDefType}:                     _GotoState257Action,
	{_State197, ImplicitStructDefType}:            _GotoState148Action,
	{_State197, ExplicitStructDefType}:            _GotoState144Action,
	{_State197, EnumValueDefType}:                 _GotoState255Action,
	{_State197, ImplicitEnumValueDefsType}:        _GotoState258Action,
	{_State197, ImplicitEnumDefType}:              _GotoState147Action,
	{_State197, ExplicitEnumValueDefsType}:        _GotoState256Action,
	{_State197, ExplicitEnumDefType}:              _GotoState142Action,
	{_State197, TraitDefType}:                     _GotoState151Action,
	{_State197, FuncTypeType}:                     _GotoState146Action,
	{_State198, IdentifierToken}:                  _GotoState260Action,
	{_State198, StructToken}:                      _GotoState33Action,
	{_State198, EnumToken}:                        _GotoState133Action,
	{_State198, TraitToken}:                       _GotoState140Action,
	{_State198, FuncToken}:                        _GotoState134Action,
	{_State198, LparenToken}:                      _GotoState137Action,
	{_State198, QuestionToken}:                    _GotoState138Action,
	{_State198, DotdotdotToken}:                   _GotoState259Action,
	{_State198, TildeTildeToken}:                  _GotoState139Action,
	{_State198, BitNegToken}:                      _GotoState132Action,
	{_State198, BitAndToken}:                      _GotoState131Action,
	{_State198, LexErrorToken}:                    _GotoState136Action,
	{_State198, AtomTypeType}:                     _GotoState141Action,
	{_State198, TraitableTypeType}:                _GotoState152Action,
	{_State198, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State198, ValueTypeType}:                    _GotoState264Action,
	{_State198, ImplicitStructDefType}:            _GotoState148Action,
	{_State198, ExplicitStructDefType}:            _GotoState144Action,
	{_State198, ImplicitEnumDefType}:              _GotoState147Action,
	{_State198, ExplicitEnumDefType}:              _GotoState142Action,
	{_State198, TraitDefType}:                     _GotoState151Action,
	{_State198, ParameterDeclType}:                _GotoState262Action,
	{_State198, ParameterDeclsType}:               _GotoState263Action,
	{_State198, OptionalParameterDeclsType}:       _GotoState261Action,
	{_State198, FuncTypeType}:                     _GotoState146Action,
	{_State201, OrToken}:                          _GotoState265Action,
	{_State202, AssignToken}:                      _GotoState266Action,
	{_State203, RparenToken}:                      _GotoState268Action,
	{_State203, OrToken}:                          _GotoState267Action,
	{_State204, CommaToken}:                       _GotoState269Action,
	{_State205, RparenToken}:                      _GotoState270Action,
	{_State207, IdentifierToken}:                  _GotoState135Action,
	{_State207, UnsafeToken}:                      _GotoState56Action,
	{_State207, StructToken}:                      _GotoState33Action,
	{_State207, EnumToken}:                        _GotoState133Action,
	{_State207, TraitToken}:                       _GotoState140Action,
	{_State207, FuncToken}:                        _GotoState271Action,
	{_State207, LparenToken}:                      _GotoState137Action,
	{_State207, QuestionToken}:                    _GotoState138Action,
	{_State207, TildeTildeToken}:                  _GotoState139Action,
	{_State207, BitNegToken}:                      _GotoState132Action,
	{_State207, BitAndToken}:                      _GotoState131Action,
	{_State207, LexErrorToken}:                    _GotoState136Action,
	{_State207, UnsafeStatementType}:              _GotoState153Action,
	{_State207, AtomTypeType}:                     _GotoState141Action,
	{_State207, TraitableTypeType}:                _GotoState152Action,
	{_State207, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State207, ValueTypeType}:                    _GotoState154Action,
	{_State207, FieldDefType}:                     _GotoState272Action,
	{_State207, ImplicitStructDefType}:            _GotoState148Action,
	{_State207, ExplicitStructDefType}:            _GotoState144Action,
	{_State207, ImplicitEnumDefType}:              _GotoState147Action,
	{_State207, ExplicitEnumDefType}:              _GotoState142Action,
	{_State207, TraitPropertyType}:                _GotoState276Action,
	{_State207, TraitPropertiesType}:              _GotoState275Action,
	{_State207, OptionalTraitPropertiesType}:      _GotoState274Action,
	{_State207, TraitDefType}:                     _GotoState151Action,
	{_State207, FuncTypeType}:                     _GotoState146Action,
	{_State207, MethodSignatureType}:              _GotoState273Action,
	{_State208, IdentifierToken}:                  _GotoState135Action,
	{_State208, UnsafeToken}:                      _GotoState56Action,
	{_State208, StructToken}:                      _GotoState33Action,
	{_State208, EnumToken}:                        _GotoState133Action,
	{_State208, TraitToken}:                       _GotoState140Action,
	{_State208, FuncToken}:                        _GotoState134Action,
	{_State208, LparenToken}:                      _GotoState137Action,
	{_State208, QuestionToken}:                    _GotoState138Action,
	{_State208, TildeTildeToken}:                  _GotoState139Action,
	{_State208, BitNegToken}:                      _GotoState132Action,
	{_State208, BitAndToken}:                      _GotoState131Action,
	{_State208, LexErrorToken}:                    _GotoState136Action,
	{_State208, UnsafeStatementType}:              _GotoState153Action,
	{_State208, AtomTypeType}:                     _GotoState141Action,
	{_State208, TraitableTypeType}:                _GotoState152Action,
	{_State208, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State208, ValueTypeType}:                    _GotoState154Action,
	{_State208, FieldDefType}:                     _GotoState277Action,
	{_State208, ImplicitStructDefType}:            _GotoState148Action,
	{_State208, ExplicitStructDefType}:            _GotoState144Action,
	{_State208, ImplicitEnumDefType}:              _GotoState147Action,
	{_State208, ExplicitEnumDefType}:              _GotoState142Action,
	{_State208, TraitDefType}:                     _GotoState151Action,
	{_State208, FuncTypeType}:                     _GotoState146Action,
	{_State209, IdentifierToken}:                  _GotoState135Action,
	{_State209, UnsafeToken}:                      _GotoState56Action,
	{_State209, StructToken}:                      _GotoState33Action,
	{_State209, EnumToken}:                        _GotoState133Action,
	{_State209, TraitToken}:                       _GotoState140Action,
	{_State209, FuncToken}:                        _GotoState134Action,
	{_State209, LparenToken}:                      _GotoState137Action,
	{_State209, QuestionToken}:                    _GotoState138Action,
	{_State209, TildeTildeToken}:                  _GotoState139Action,
	{_State209, BitNegToken}:                      _GotoState132Action,
	{_State209, BitAndToken}:                      _GotoState131Action,
	{_State209, LexErrorToken}:                    _GotoState136Action,
	{_State209, UnsafeStatementType}:              _GotoState153Action,
	{_State209, AtomTypeType}:                     _GotoState141Action,
	{_State209, TraitableTypeType}:                _GotoState152Action,
	{_State209, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State209, ValueTypeType}:                    _GotoState154Action,
	{_State209, FieldDefType}:                     _GotoState278Action,
	{_State209, ImplicitStructDefType}:            _GotoState148Action,
	{_State209, ExplicitStructDefType}:            _GotoState144Action,
	{_State209, ImplicitEnumDefType}:              _GotoState147Action,
	{_State209, ExplicitEnumDefType}:              _GotoState142Action,
	{_State209, TraitDefType}:                     _GotoState151Action,
	{_State209, FuncTypeType}:                     _GotoState146Action,
	{_State211, IdentifierToken}:                  _GotoState155Action,
	{_State211, StructToken}:                      _GotoState33Action,
	{_State211, EnumToken}:                        _GotoState133Action,
	{_State211, TraitToken}:                       _GotoState140Action,
	{_State211, LparenToken}:                      _GotoState137Action,
	{_State211, TildeTildeToken}:                  _GotoState139Action,
	{_State211, BitNegToken}:                      _GotoState132Action,
	{_State211, LexErrorToken}:                    _GotoState136Action,
	{_State211, AtomTypeType}:                     _GotoState141Action,
	{_State211, TraitableTypeType}:                _GotoState279Action,
	{_State211, ImplicitStructDefType}:            _GotoState148Action,
	{_State211, ExplicitStructDefType}:            _GotoState144Action,
	{_State211, ImplicitEnumDefType}:              _GotoState147Action,
	{_State211, ExplicitEnumDefType}:              _GotoState142Action,
	{_State211, TraitDefType}:                     _GotoState151Action,
	{_State212, IdentifierToken}:                  _GotoState155Action,
	{_State212, StructToken}:                      _GotoState33Action,
	{_State212, EnumToken}:                        _GotoState133Action,
	{_State212, TraitToken}:                       _GotoState140Action,
	{_State212, LparenToken}:                      _GotoState137Action,
	{_State212, TildeTildeToken}:                  _GotoState139Action,
	{_State212, BitNegToken}:                      _GotoState132Action,
	{_State212, LexErrorToken}:                    _GotoState136Action,
	{_State212, AtomTypeType}:                     _GotoState141Action,
	{_State212, TraitableTypeType}:                _GotoState280Action,
	{_State212, ImplicitStructDefType}:            _GotoState148Action,
	{_State212, ExplicitStructDefType}:            _GotoState144Action,
	{_State212, ImplicitEnumDefType}:              _GotoState147Action,
	{_State212, ExplicitEnumDefType}:              _GotoState142Action,
	{_State212, TraitDefType}:                     _GotoState151Action,
	{_State213, IdentifierToken}:                  _GotoState155Action,
	{_State213, StructToken}:                      _GotoState33Action,
	{_State213, EnumToken}:                        _GotoState133Action,
	{_State213, TraitToken}:                       _GotoState140Action,
	{_State213, LparenToken}:                      _GotoState137Action,
	{_State213, TildeTildeToken}:                  _GotoState139Action,
	{_State213, BitNegToken}:                      _GotoState132Action,
	{_State213, LexErrorToken}:                    _GotoState136Action,
	{_State213, AtomTypeType}:                     _GotoState141Action,
	{_State213, TraitableTypeType}:                _GotoState281Action,
	{_State213, ImplicitStructDefType}:            _GotoState148Action,
	{_State213, ExplicitStructDefType}:            _GotoState144Action,
	{_State213, ImplicitEnumDefType}:              _GotoState147Action,
	{_State213, ExplicitEnumDefType}:              _GotoState142Action,
	{_State213, TraitDefType}:                     _GotoState151Action,
	{_State214, IdentifierToken}:                  _GotoState155Action,
	{_State214, StructToken}:                      _GotoState33Action,
	{_State214, EnumToken}:                        _GotoState133Action,
	{_State214, TraitToken}:                       _GotoState140Action,
	{_State214, FuncToken}:                        _GotoState134Action,
	{_State214, LparenToken}:                      _GotoState137Action,
	{_State214, QuestionToken}:                    _GotoState138Action,
	{_State214, TildeTildeToken}:                  _GotoState139Action,
	{_State214, BitNegToken}:                      _GotoState132Action,
	{_State214, BitAndToken}:                      _GotoState131Action,
	{_State214, LexErrorToken}:                    _GotoState136Action,
	{_State214, AtomTypeType}:                     _GotoState141Action,
	{_State214, TraitableTypeType}:                _GotoState152Action,
	{_State214, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State214, ValueTypeType}:                    _GotoState282Action,
	{_State214, ImplicitStructDefType}:            _GotoState148Action,
	{_State214, ExplicitStructDefType}:            _GotoState144Action,
	{_State214, ImplicitEnumDefType}:              _GotoState147Action,
	{_State214, ExplicitEnumDefType}:              _GotoState142Action,
	{_State214, TraitDefType}:                     _GotoState151Action,
	{_State214, FuncTypeType}:                     _GotoState146Action,
	{_State217, CommaToken}:                       _GotoState127Action,
	{_State218, RparenToken}:                      _GotoState283Action,
	{_State220, IntegerLiteralToken}:              _GotoState25Action,
	{_State220, FloatLiteralToken}:                _GotoState22Action,
	{_State220, RuneLiteralToken}:                 _GotoState31Action,
	{_State220, StringLiteralToken}:               _GotoState32Action,
	{_State220, IdentifierToken}:                  _GotoState24Action,
	{_State220, TrueToken}:                        _GotoState35Action,
	{_State220, FalseToken}:                       _GotoState21Action,
	{_State220, StructToken}:                      _GotoState33Action,
	{_State220, FuncToken}:                        _GotoState23Action,
	{_State220, LabelDeclToken}:                   _GotoState26Action,
	{_State220, LparenToken}:                      _GotoState28Action,
	{_State220, NotToken}:                         _GotoState30Action,
	{_State220, SubToken}:                         _GotoState34Action,
	{_State220, MulToken}:                         _GotoState29Action,
	{_State220, BitNegToken}:                      _GotoState20Action,
	{_State220, BitAndToken}:                      _GotoState19Action,
	{_State220, LexErrorToken}:                    _GotoState27Action,
	{_State220, OptionalLabelDeclType}:            _GotoState111Action,
	{_State220, SequenceExprType}:                 _GotoState284Action,
	{_State220, BlockExprType}:                    _GotoState42Action,
	{_State220, CallExprType}:                     _GotoState43Action,
	{_State220, AtomExprType}:                     _GotoState41Action,
	{_State220, LiteralType}:                      _GotoState46Action,
	{_State220, AnonymousStructExprType}:          _GotoState40Action,
	{_State220, AccessExprType}:                   _GotoState36Action,
	{_State220, PostfixUnaryExprType}:             _GotoState50Action,
	{_State220, PrefixUnaryOpType}:                _GotoState52Action,
	{_State220, PrefixUnaryExprType}:              _GotoState51Action,
	{_State220, MulExprType}:                      _GotoState47Action,
	{_State220, AddExprType}:                      _GotoState37Action,
	{_State220, CmpExprType}:                      _GotoState44Action,
	{_State220, AndExprType}:                      _GotoState38Action,
	{_State220, OrExprType}:                       _GotoState49Action,
	{_State220, ExplicitStructDefType}:            _GotoState45Action,
	{_State220, AnonymousFuncExprType}:            _GotoState39Action,
	{_State221, DoToken}:                          _GotoState285Action,
	{_State222, SemicolonToken}:                   _GotoState286Action,
	{_State224, LbraceToken}:                      _GotoState104Action,
	{_State224, BlockBodyType}:                    _GotoState287Action,
	{_State225, ElseToken}:                        _GotoState288Action,
	{_State226, IntegerLiteralToken}:              _GotoState25Action,
	{_State226, FloatLiteralToken}:                _GotoState22Action,
	{_State226, RuneLiteralToken}:                 _GotoState31Action,
	{_State226, StringLiteralToken}:               _GotoState32Action,
	{_State226, IdentifierToken}:                  _GotoState24Action,
	{_State226, TrueToken}:                        _GotoState35Action,
	{_State226, FalseToken}:                       _GotoState21Action,
	{_State226, StructToken}:                      _GotoState33Action,
	{_State226, FuncToken}:                        _GotoState23Action,
	{_State226, LabelDeclToken}:                   _GotoState26Action,
	{_State226, LparenToken}:                      _GotoState28Action,
	{_State226, LexErrorToken}:                    _GotoState27Action,
	{_State226, OptionalLabelDeclType}:            _GotoState111Action,
	{_State226, BlockExprType}:                    _GotoState42Action,
	{_State226, CallExprType}:                     _GotoState290Action,
	{_State226, AtomExprType}:                     _GotoState41Action,
	{_State226, LiteralType}:                      _GotoState46Action,
	{_State226, AnonymousStructExprType}:          _GotoState40Action,
	{_State226, AccessExprType}:                   _GotoState289Action,
	{_State226, ExplicitStructDefType}:            _GotoState45Action,
	{_State226, AnonymousFuncExprType}:            _GotoState39Action,
	{_State229, IntegerLiteralToken}:              _GotoState25Action,
	{_State229, FloatLiteralToken}:                _GotoState22Action,
	{_State229, RuneLiteralToken}:                 _GotoState31Action,
	{_State229, StringLiteralToken}:               _GotoState32Action,
	{_State229, IdentifierToken}:                  _GotoState24Action,
	{_State229, TrueToken}:                        _GotoState35Action,
	{_State229, FalseToken}:                       _GotoState21Action,
	{_State229, StructToken}:                      _GotoState33Action,
	{_State229, FuncToken}:                        _GotoState23Action,
	{_State229, LabelDeclToken}:                   _GotoState26Action,
	{_State229, LparenToken}:                      _GotoState28Action,
	{_State229, LexErrorToken}:                    _GotoState27Action,
	{_State229, OptionalLabelDeclType}:            _GotoState111Action,
	{_State229, BlockExprType}:                    _GotoState42Action,
	{_State229, CallExprType}:                     _GotoState291Action,
	{_State229, AtomExprType}:                     _GotoState41Action,
	{_State229, LiteralType}:                      _GotoState46Action,
	{_State229, AnonymousStructExprType}:          _GotoState40Action,
	{_State229, AccessExprType}:                   _GotoState289Action,
	{_State229, ExplicitStructDefType}:            _GotoState45Action,
	{_State229, AnonymousFuncExprType}:            _GotoState39Action,
	{_State232, LbracketToken}:                    _GotoState77Action,
	{_State232, DotToken}:                         _GotoState76Action,
	{_State232, QuestionToken}:                    _GotoState78Action,
	{_State232, DollarLbracketToken}:              _GotoState75Action,
	{_State232, AddAssignToken}:                   _GotoState292Action,
	{_State232, SubAssignToken}:                   _GotoState303Action,
	{_State232, MulAssignToken}:                   _GotoState302Action,
	{_State232, DivAssignToken}:                   _GotoState300Action,
	{_State232, ModAssignToken}:                   _GotoState301Action,
	{_State232, AddOneAssignToken}:                _GotoState293Action,
	{_State232, SubOneAssignToken}:                _GotoState304Action,
	{_State232, BitNegAssignToken}:                _GotoState296Action,
	{_State232, BitAndAssignToken}:                _GotoState294Action,
	{_State232, BitOrAssignToken}:                 _GotoState297Action,
	{_State232, BitXorAssignToken}:                _GotoState299Action,
	{_State232, BitLshiftAssignToken}:             _GotoState295Action,
	{_State232, BitRshiftAssignToken}:             _GotoState298Action,
	{_State232, UnaryOpAssignType}:                _GotoState306Action,
	{_State232, BinaryOpAssignType}:               _GotoState305Action,
	{_State232, OptionalGenericBindingType}:       _GotoState79Action,
	{_State234, CommaToken}:                       _GotoState307Action,
	{_State236, JumpLabelToken}:                   _GotoState308Action,
	{_State236, OptionalJumpLabelType}:            _GotoState309Action,
	{_State238, NewlinesToken}:                    _GotoState310Action,
	{_State238, SemicolonToken}:                   _GotoState311Action,
	{_State240, CaseToken}:                        _GotoState312Action,
	{_State241, StringLiteralToken}:               _GotoState313Action,
	{_State244, NewlinesToken}:                    _GotoState314Action,
	{_State244, SemicolonToken}:                   _GotoState315Action,
	{_State247, IdentifierToken}:                  _GotoState178Action,
	{_State247, GenericParameterDefType}:          _GotoState316Action,
	{_State249, IdentifierToken}:                  _GotoState155Action,
	{_State249, StructToken}:                      _GotoState33Action,
	{_State249, EnumToken}:                        _GotoState133Action,
	{_State249, TraitToken}:                       _GotoState140Action,
	{_State249, FuncToken}:                        _GotoState134Action,
	{_State249, LparenToken}:                      _GotoState137Action,
	{_State249, QuestionToken}:                    _GotoState138Action,
	{_State249, TildeTildeToken}:                  _GotoState139Action,
	{_State249, BitNegToken}:                      _GotoState132Action,
	{_State249, BitAndToken}:                      _GotoState131Action,
	{_State249, LexErrorToken}:                    _GotoState136Action,
	{_State249, AtomTypeType}:                     _GotoState141Action,
	{_State249, TraitableTypeType}:                _GotoState152Action,
	{_State249, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State249, ValueTypeType}:                    _GotoState317Action,
	{_State249, ImplicitStructDefType}:            _GotoState148Action,
	{_State249, ExplicitStructDefType}:            _GotoState144Action,
	{_State249, ImplicitEnumDefType}:              _GotoState147Action,
	{_State249, ExplicitEnumDefType}:              _GotoState142Action,
	{_State249, TraitDefType}:                     _GotoState151Action,
	{_State249, FuncTypeType}:                     _GotoState146Action,
	{_State251, IdentifierToken}:                  _GotoState120Action,
	{_State251, ParameterDefType}:                 _GotoState124Action,
	{_State251, ParameterDefsType}:                _GotoState125Action,
	{_State251, OptionalParameterDefsType}:        _GotoState318Action,
	{_State252, LbraceToken}:                      _GotoState104Action,
	{_State252, BlockBodyType}:                    _GotoState319Action,
	{_State255, NewlinesToken}:                    _GotoState320Action,
	{_State255, OrToken}:                          _GotoState321Action,
	{_State256, RparenToken}:                      _GotoState322Action,
	{_State257, AssignToken}:                      _GotoState266Action,
	{_State258, NewlinesToken}:                    _GotoState323Action,
	{_State258, OrToken}:                          _GotoState324Action,
	{_State259, IdentifierToken}:                  _GotoState155Action,
	{_State259, StructToken}:                      _GotoState33Action,
	{_State259, EnumToken}:                        _GotoState133Action,
	{_State259, TraitToken}:                       _GotoState140Action,
	{_State259, FuncToken}:                        _GotoState134Action,
	{_State259, LparenToken}:                      _GotoState137Action,
	{_State259, QuestionToken}:                    _GotoState138Action,
	{_State259, TildeTildeToken}:                  _GotoState139Action,
	{_State259, BitNegToken}:                      _GotoState132Action,
	{_State259, BitAndToken}:                      _GotoState131Action,
	{_State259, LexErrorToken}:                    _GotoState136Action,
	{_State259, AtomTypeType}:                     _GotoState141Action,
	{_State259, TraitableTypeType}:                _GotoState152Action,
	{_State259, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State259, ValueTypeType}:                    _GotoState325Action,
	{_State259, ImplicitStructDefType}:            _GotoState148Action,
	{_State259, ExplicitStructDefType}:            _GotoState144Action,
	{_State259, ImplicitEnumDefType}:              _GotoState147Action,
	{_State259, ExplicitEnumDefType}:              _GotoState142Action,
	{_State259, TraitDefType}:                     _GotoState151Action,
	{_State259, FuncTypeType}:                     _GotoState146Action,
	{_State260, IdentifierToken}:                  _GotoState155Action,
	{_State260, StructToken}:                      _GotoState33Action,
	{_State260, EnumToken}:                        _GotoState133Action,
	{_State260, TraitToken}:                       _GotoState140Action,
	{_State260, FuncToken}:                        _GotoState134Action,
	{_State260, LparenToken}:                      _GotoState137Action,
	{_State260, QuestionToken}:                    _GotoState138Action,
	{_State260, DollarLbracketToken}:              _GotoState75Action,
	{_State260, DotdotdotToken}:                   _GotoState326Action,
	{_State260, TildeTildeToken}:                  _GotoState139Action,
	{_State260, BitNegToken}:                      _GotoState132Action,
	{_State260, BitAndToken}:                      _GotoState131Action,
	{_State260, LexErrorToken}:                    _GotoState136Action,
	{_State260, OptionalGenericBindingType}:       _GotoState199Action,
	{_State260, AtomTypeType}:                     _GotoState141Action,
	{_State260, TraitableTypeType}:                _GotoState152Action,
	{_State260, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State260, ValueTypeType}:                    _GotoState327Action,
	{_State260, ImplicitStructDefType}:            _GotoState148Action,
	{_State260, ExplicitStructDefType}:            _GotoState144Action,
	{_State260, ImplicitEnumDefType}:              _GotoState147Action,
	{_State260, ExplicitEnumDefType}:              _GotoState142Action,
	{_State260, TraitDefType}:                     _GotoState151Action,
	{_State260, FuncTypeType}:                     _GotoState146Action,
	{_State261, RparenToken}:                      _GotoState328Action,
	{_State263, CommaToken}:                       _GotoState329Action,
	{_State265, IdentifierToken}:                  _GotoState135Action,
	{_State265, UnsafeToken}:                      _GotoState56Action,
	{_State265, StructToken}:                      _GotoState33Action,
	{_State265, EnumToken}:                        _GotoState133Action,
	{_State265, TraitToken}:                       _GotoState140Action,
	{_State265, FuncToken}:                        _GotoState134Action,
	{_State265, LparenToken}:                      _GotoState137Action,
	{_State265, QuestionToken}:                    _GotoState138Action,
	{_State265, TildeTildeToken}:                  _GotoState139Action,
	{_State265, BitNegToken}:                      _GotoState132Action,
	{_State265, BitAndToken}:                      _GotoState131Action,
	{_State265, LexErrorToken}:                    _GotoState136Action,
	{_State265, UnsafeStatementType}:              _GotoState153Action,
	{_State265, AtomTypeType}:                     _GotoState141Action,
	{_State265, TraitableTypeType}:                _GotoState152Action,
	{_State265, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State265, ValueTypeType}:                    _GotoState154Action,
	{_State265, FieldDefType}:                     _GotoState257Action,
	{_State265, ImplicitStructDefType}:            _GotoState148Action,
	{_State265, ExplicitStructDefType}:            _GotoState144Action,
	{_State265, EnumValueDefType}:                 _GotoState330Action,
	{_State265, ImplicitEnumDefType}:              _GotoState147Action,
	{_State265, ExplicitEnumDefType}:              _GotoState142Action,
	{_State265, TraitDefType}:                     _GotoState151Action,
	{_State265, FuncTypeType}:                     _GotoState146Action,
	{_State266, DefaultToken}:                     _GotoState331Action,
	{_State267, IdentifierToken}:                  _GotoState135Action,
	{_State267, UnsafeToken}:                      _GotoState56Action,
	{_State267, StructToken}:                      _GotoState33Action,
	{_State267, EnumToken}:                        _GotoState133Action,
	{_State267, TraitToken}:                       _GotoState140Action,
	{_State267, FuncToken}:                        _GotoState134Action,
	{_State267, LparenToken}:                      _GotoState137Action,
	{_State267, QuestionToken}:                    _GotoState138Action,
	{_State267, TildeTildeToken}:                  _GotoState139Action,
	{_State267, BitNegToken}:                      _GotoState132Action,
	{_State267, BitAndToken}:                      _GotoState131Action,
	{_State267, LexErrorToken}:                    _GotoState136Action,
	{_State267, UnsafeStatementType}:              _GotoState153Action,
	{_State267, AtomTypeType}:                     _GotoState141Action,
	{_State267, TraitableTypeType}:                _GotoState152Action,
	{_State267, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State267, ValueTypeType}:                    _GotoState154Action,
	{_State267, FieldDefType}:                     _GotoState257Action,
	{_State267, ImplicitStructDefType}:            _GotoState148Action,
	{_State267, ExplicitStructDefType}:            _GotoState144Action,
	{_State267, EnumValueDefType}:                 _GotoState332Action,
	{_State267, ImplicitEnumDefType}:              _GotoState147Action,
	{_State267, ExplicitEnumDefType}:              _GotoState142Action,
	{_State267, TraitDefType}:                     _GotoState151Action,
	{_State267, FuncTypeType}:                     _GotoState146Action,
	{_State269, IdentifierToken}:                  _GotoState135Action,
	{_State269, UnsafeToken}:                      _GotoState56Action,
	{_State269, StructToken}:                      _GotoState33Action,
	{_State269, EnumToken}:                        _GotoState133Action,
	{_State269, TraitToken}:                       _GotoState140Action,
	{_State269, FuncToken}:                        _GotoState134Action,
	{_State269, LparenToken}:                      _GotoState137Action,
	{_State269, QuestionToken}:                    _GotoState138Action,
	{_State269, TildeTildeToken}:                  _GotoState139Action,
	{_State269, BitNegToken}:                      _GotoState132Action,
	{_State269, BitAndToken}:                      _GotoState131Action,
	{_State269, LexErrorToken}:                    _GotoState136Action,
	{_State269, UnsafeStatementType}:              _GotoState153Action,
	{_State269, AtomTypeType}:                     _GotoState141Action,
	{_State269, TraitableTypeType}:                _GotoState152Action,
	{_State269, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State269, ValueTypeType}:                    _GotoState154Action,
	{_State269, FieldDefType}:                     _GotoState333Action,
	{_State269, ImplicitStructDefType}:            _GotoState148Action,
	{_State269, ExplicitStructDefType}:            _GotoState144Action,
	{_State269, ImplicitEnumDefType}:              _GotoState147Action,
	{_State269, ExplicitEnumDefType}:              _GotoState142Action,
	{_State269, TraitDefType}:                     _GotoState151Action,
	{_State269, FuncTypeType}:                     _GotoState146Action,
	{_State271, IdentifierToken}:                  _GotoState334Action,
	{_State271, LparenToken}:                      _GotoState198Action,
	{_State274, RparenToken}:                      _GotoState335Action,
	{_State275, NewlinesToken}:                    _GotoState337Action,
	{_State275, CommaToken}:                       _GotoState336Action,
	{_State285, LbraceToken}:                      _GotoState104Action,
	{_State285, BlockBodyType}:                    _GotoState338Action,
	{_State286, IntegerLiteralToken}:              _GotoState25Action,
	{_State286, FloatLiteralToken}:                _GotoState22Action,
	{_State286, RuneLiteralToken}:                 _GotoState31Action,
	{_State286, StringLiteralToken}:               _GotoState32Action,
	{_State286, IdentifierToken}:                  _GotoState24Action,
	{_State286, TrueToken}:                        _GotoState35Action,
	{_State286, FalseToken}:                       _GotoState21Action,
	{_State286, StructToken}:                      _GotoState33Action,
	{_State286, FuncToken}:                        _GotoState23Action,
	{_State286, LabelDeclToken}:                   _GotoState26Action,
	{_State286, LparenToken}:                      _GotoState28Action,
	{_State286, NotToken}:                         _GotoState30Action,
	{_State286, SubToken}:                         _GotoState34Action,
	{_State286, MulToken}:                         _GotoState29Action,
	{_State286, BitNegToken}:                      _GotoState20Action,
	{_State286, BitAndToken}:                      _GotoState19Action,
	{_State286, LexErrorToken}:                    _GotoState27Action,
	{_State286, OptionalLabelDeclType}:            _GotoState111Action,
	{_State286, OptionalSequenceExprType}:         _GotoState339Action,
	{_State286, SequenceExprType}:                 _GotoState223Action,
	{_State286, BlockExprType}:                    _GotoState42Action,
	{_State286, CallExprType}:                     _GotoState43Action,
	{_State286, AtomExprType}:                     _GotoState41Action,
	{_State286, LiteralType}:                      _GotoState46Action,
	{_State286, AnonymousStructExprType}:          _GotoState40Action,
	{_State286, AccessExprType}:                   _GotoState36Action,
	{_State286, PostfixUnaryExprType}:             _GotoState50Action,
	{_State286, PrefixUnaryOpType}:                _GotoState52Action,
	{_State286, PrefixUnaryExprType}:              _GotoState51Action,
	{_State286, MulExprType}:                      _GotoState47Action,
	{_State286, AddExprType}:                      _GotoState37Action,
	{_State286, CmpExprType}:                      _GotoState44Action,
	{_State286, AndExprType}:                      _GotoState38Action,
	{_State286, OrExprType}:                       _GotoState49Action,
	{_State286, ExplicitStructDefType}:            _GotoState45Action,
	{_State286, AnonymousFuncExprType}:            _GotoState39Action,
	{_State288, IfToken}:                          _GotoState103Action,
	{_State288, LbraceToken}:                      _GotoState104Action,
	{_State288, IfExprType}:                       _GotoState341Action,
	{_State288, BlockBodyType}:                    _GotoState340Action,
	{_State289, LbracketToken}:                    _GotoState77Action,
	{_State289, DotToken}:                         _GotoState76Action,
	{_State289, DollarLbracketToken}:              _GotoState75Action,
	{_State289, OptionalGenericBindingType}:       _GotoState79Action,
	{_State305, IntegerLiteralToken}:              _GotoState25Action,
	{_State305, FloatLiteralToken}:                _GotoState22Action,
	{_State305, RuneLiteralToken}:                 _GotoState31Action,
	{_State305, StringLiteralToken}:               _GotoState32Action,
	{_State305, IdentifierToken}:                  _GotoState24Action,
	{_State305, TrueToken}:                        _GotoState35Action,
	{_State305, FalseToken}:                       _GotoState21Action,
	{_State305, StructToken}:                      _GotoState33Action,
	{_State305, FuncToken}:                        _GotoState23Action,
	{_State305, LabelDeclToken}:                   _GotoState26Action,
	{_State305, LparenToken}:                      _GotoState28Action,
	{_State305, NotToken}:                         _GotoState30Action,
	{_State305, SubToken}:                         _GotoState34Action,
	{_State305, MulToken}:                         _GotoState29Action,
	{_State305, BitNegToken}:                      _GotoState20Action,
	{_State305, BitAndToken}:                      _GotoState19Action,
	{_State305, LexErrorToken}:                    _GotoState27Action,
	{_State305, ExpressionType}:                   _GotoState342Action,
	{_State305, OptionalLabelDeclType}:            _GotoState48Action,
	{_State305, SequenceExprType}:                 _GotoState53Action,
	{_State305, BlockExprType}:                    _GotoState42Action,
	{_State305, CallExprType}:                     _GotoState43Action,
	{_State305, AtomExprType}:                     _GotoState41Action,
	{_State305, LiteralType}:                      _GotoState46Action,
	{_State305, AnonymousStructExprType}:          _GotoState40Action,
	{_State305, AccessExprType}:                   _GotoState36Action,
	{_State305, PostfixUnaryExprType}:             _GotoState50Action,
	{_State305, PrefixUnaryOpType}:                _GotoState52Action,
	{_State305, PrefixUnaryExprType}:              _GotoState51Action,
	{_State305, MulExprType}:                      _GotoState47Action,
	{_State305, AddExprType}:                      _GotoState37Action,
	{_State305, CmpExprType}:                      _GotoState44Action,
	{_State305, AndExprType}:                      _GotoState38Action,
	{_State305, OrExprType}:                       _GotoState49Action,
	{_State305, ExplicitStructDefType}:            _GotoState45Action,
	{_State305, AnonymousFuncExprType}:            _GotoState39Action,
	{_State307, IntegerLiteralToken}:              _GotoState25Action,
	{_State307, FloatLiteralToken}:                _GotoState22Action,
	{_State307, RuneLiteralToken}:                 _GotoState31Action,
	{_State307, StringLiteralToken}:               _GotoState32Action,
	{_State307, IdentifierToken}:                  _GotoState24Action,
	{_State307, TrueToken}:                        _GotoState35Action,
	{_State307, FalseToken}:                       _GotoState21Action,
	{_State307, StructToken}:                      _GotoState33Action,
	{_State307, FuncToken}:                        _GotoState23Action,
	{_State307, LabelDeclToken}:                   _GotoState26Action,
	{_State307, LparenToken}:                      _GotoState28Action,
	{_State307, NotToken}:                         _GotoState30Action,
	{_State307, SubToken}:                         _GotoState34Action,
	{_State307, MulToken}:                         _GotoState29Action,
	{_State307, BitNegToken}:                      _GotoState20Action,
	{_State307, BitAndToken}:                      _GotoState19Action,
	{_State307, LexErrorToken}:                    _GotoState27Action,
	{_State307, ExpressionType}:                   _GotoState343Action,
	{_State307, OptionalLabelDeclType}:            _GotoState48Action,
	{_State307, SequenceExprType}:                 _GotoState53Action,
	{_State307, BlockExprType}:                    _GotoState42Action,
	{_State307, CallExprType}:                     _GotoState43Action,
	{_State307, AtomExprType}:                     _GotoState41Action,
	{_State307, LiteralType}:                      _GotoState46Action,
	{_State307, AnonymousStructExprType}:          _GotoState40Action,
	{_State307, AccessExprType}:                   _GotoState36Action,
	{_State307, PostfixUnaryExprType}:             _GotoState50Action,
	{_State307, PrefixUnaryOpType}:                _GotoState52Action,
	{_State307, PrefixUnaryExprType}:              _GotoState51Action,
	{_State307, MulExprType}:                      _GotoState47Action,
	{_State307, AddExprType}:                      _GotoState37Action,
	{_State307, CmpExprType}:                      _GotoState44Action,
	{_State307, AndExprType}:                      _GotoState38Action,
	{_State307, OrExprType}:                       _GotoState49Action,
	{_State307, ExplicitStructDefType}:            _GotoState45Action,
	{_State307, AnonymousFuncExprType}:            _GotoState39Action,
	{_State309, IntegerLiteralToken}:              _GotoState25Action,
	{_State309, FloatLiteralToken}:                _GotoState22Action,
	{_State309, RuneLiteralToken}:                 _GotoState31Action,
	{_State309, StringLiteralToken}:               _GotoState32Action,
	{_State309, IdentifierToken}:                  _GotoState24Action,
	{_State309, TrueToken}:                        _GotoState35Action,
	{_State309, FalseToken}:                       _GotoState21Action,
	{_State309, StructToken}:                      _GotoState33Action,
	{_State309, FuncToken}:                        _GotoState23Action,
	{_State309, LabelDeclToken}:                   _GotoState26Action,
	{_State309, LparenToken}:                      _GotoState28Action,
	{_State309, NotToken}:                         _GotoState30Action,
	{_State309, SubToken}:                         _GotoState34Action,
	{_State309, MulToken}:                         _GotoState29Action,
	{_State309, BitNegToken}:                      _GotoState20Action,
	{_State309, BitAndToken}:                      _GotoState19Action,
	{_State309, LexErrorToken}:                    _GotoState27Action,
	{_State309, ExpressionType}:                   _GotoState233Action,
	{_State309, OptionalLabelDeclType}:            _GotoState48Action,
	{_State309, SequenceExprType}:                 _GotoState53Action,
	{_State309, BlockExprType}:                    _GotoState42Action,
	{_State309, ExpressionsType}:                  _GotoState344Action,
	{_State309, OptionalExpressionsType}:          _GotoState345Action,
	{_State309, CallExprType}:                     _GotoState43Action,
	{_State309, AtomExprType}:                     _GotoState41Action,
	{_State309, LiteralType}:                      _GotoState46Action,
	{_State309, AnonymousStructExprType}:          _GotoState40Action,
	{_State309, AccessExprType}:                   _GotoState36Action,
	{_State309, PostfixUnaryExprType}:             _GotoState50Action,
	{_State309, PrefixUnaryOpType}:                _GotoState52Action,
	{_State309, PrefixUnaryExprType}:              _GotoState51Action,
	{_State309, MulExprType}:                      _GotoState47Action,
	{_State309, AddExprType}:                      _GotoState37Action,
	{_State309, CmpExprType}:                      _GotoState44Action,
	{_State309, AndExprType}:                      _GotoState38Action,
	{_State309, OrExprType}:                       _GotoState49Action,
	{_State309, ExplicitStructDefType}:            _GotoState45Action,
	{_State309, AnonymousFuncExprType}:            _GotoState39Action,
	{_State312, DefaultToken}:                     _GotoState346Action,
	{_State318, RparenToken}:                      _GotoState347Action,
	{_State320, IdentifierToken}:                  _GotoState135Action,
	{_State320, UnsafeToken}:                      _GotoState56Action,
	{_State320, StructToken}:                      _GotoState33Action,
	{_State320, EnumToken}:                        _GotoState133Action,
	{_State320, TraitToken}:                       _GotoState140Action,
	{_State320, FuncToken}:                        _GotoState134Action,
	{_State320, LparenToken}:                      _GotoState137Action,
	{_State320, QuestionToken}:                    _GotoState138Action,
	{_State320, TildeTildeToken}:                  _GotoState139Action,
	{_State320, BitNegToken}:                      _GotoState132Action,
	{_State320, BitAndToken}:                      _GotoState131Action,
	{_State320, LexErrorToken}:                    _GotoState136Action,
	{_State320, UnsafeStatementType}:              _GotoState153Action,
	{_State320, AtomTypeType}:                     _GotoState141Action,
	{_State320, TraitableTypeType}:                _GotoState152Action,
	{_State320, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State320, ValueTypeType}:                    _GotoState154Action,
	{_State320, FieldDefType}:                     _GotoState257Action,
	{_State320, ImplicitStructDefType}:            _GotoState148Action,
	{_State320, ExplicitStructDefType}:            _GotoState144Action,
	{_State320, EnumValueDefType}:                 _GotoState348Action,
	{_State320, ImplicitEnumDefType}:              _GotoState147Action,
	{_State320, ExplicitEnumDefType}:              _GotoState142Action,
	{_State320, TraitDefType}:                     _GotoState151Action,
	{_State320, FuncTypeType}:                     _GotoState146Action,
	{_State321, IdentifierToken}:                  _GotoState135Action,
	{_State321, UnsafeToken}:                      _GotoState56Action,
	{_State321, StructToken}:                      _GotoState33Action,
	{_State321, EnumToken}:                        _GotoState133Action,
	{_State321, TraitToken}:                       _GotoState140Action,
	{_State321, FuncToken}:                        _GotoState134Action,
	{_State321, LparenToken}:                      _GotoState137Action,
	{_State321, QuestionToken}:                    _GotoState138Action,
	{_State321, TildeTildeToken}:                  _GotoState139Action,
	{_State321, BitNegToken}:                      _GotoState132Action,
	{_State321, BitAndToken}:                      _GotoState131Action,
	{_State321, LexErrorToken}:                    _GotoState136Action,
	{_State321, UnsafeStatementType}:              _GotoState153Action,
	{_State321, AtomTypeType}:                     _GotoState141Action,
	{_State321, TraitableTypeType}:                _GotoState152Action,
	{_State321, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State321, ValueTypeType}:                    _GotoState154Action,
	{_State321, FieldDefType}:                     _GotoState257Action,
	{_State321, ImplicitStructDefType}:            _GotoState148Action,
	{_State321, ExplicitStructDefType}:            _GotoState144Action,
	{_State321, EnumValueDefType}:                 _GotoState349Action,
	{_State321, ImplicitEnumDefType}:              _GotoState147Action,
	{_State321, ExplicitEnumDefType}:              _GotoState142Action,
	{_State321, TraitDefType}:                     _GotoState151Action,
	{_State321, FuncTypeType}:                     _GotoState146Action,
	{_State323, IdentifierToken}:                  _GotoState135Action,
	{_State323, UnsafeToken}:                      _GotoState56Action,
	{_State323, StructToken}:                      _GotoState33Action,
	{_State323, EnumToken}:                        _GotoState133Action,
	{_State323, TraitToken}:                       _GotoState140Action,
	{_State323, FuncToken}:                        _GotoState134Action,
	{_State323, LparenToken}:                      _GotoState137Action,
	{_State323, QuestionToken}:                    _GotoState138Action,
	{_State323, TildeTildeToken}:                  _GotoState139Action,
	{_State323, BitNegToken}:                      _GotoState132Action,
	{_State323, BitAndToken}:                      _GotoState131Action,
	{_State323, LexErrorToken}:                    _GotoState136Action,
	{_State323, UnsafeStatementType}:              _GotoState153Action,
	{_State323, AtomTypeType}:                     _GotoState141Action,
	{_State323, TraitableTypeType}:                _GotoState152Action,
	{_State323, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State323, ValueTypeType}:                    _GotoState154Action,
	{_State323, FieldDefType}:                     _GotoState257Action,
	{_State323, ImplicitStructDefType}:            _GotoState148Action,
	{_State323, ExplicitStructDefType}:            _GotoState144Action,
	{_State323, EnumValueDefType}:                 _GotoState350Action,
	{_State323, ImplicitEnumDefType}:              _GotoState147Action,
	{_State323, ExplicitEnumDefType}:              _GotoState142Action,
	{_State323, TraitDefType}:                     _GotoState151Action,
	{_State323, FuncTypeType}:                     _GotoState146Action,
	{_State324, IdentifierToken}:                  _GotoState135Action,
	{_State324, UnsafeToken}:                      _GotoState56Action,
	{_State324, StructToken}:                      _GotoState33Action,
	{_State324, EnumToken}:                        _GotoState133Action,
	{_State324, TraitToken}:                       _GotoState140Action,
	{_State324, FuncToken}:                        _GotoState134Action,
	{_State324, LparenToken}:                      _GotoState137Action,
	{_State324, QuestionToken}:                    _GotoState138Action,
	{_State324, TildeTildeToken}:                  _GotoState139Action,
	{_State324, BitNegToken}:                      _GotoState132Action,
	{_State324, BitAndToken}:                      _GotoState131Action,
	{_State324, LexErrorToken}:                    _GotoState136Action,
	{_State324, UnsafeStatementType}:              _GotoState153Action,
	{_State324, AtomTypeType}:                     _GotoState141Action,
	{_State324, TraitableTypeType}:                _GotoState152Action,
	{_State324, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State324, ValueTypeType}:                    _GotoState154Action,
	{_State324, FieldDefType}:                     _GotoState257Action,
	{_State324, ImplicitStructDefType}:            _GotoState148Action,
	{_State324, ExplicitStructDefType}:            _GotoState144Action,
	{_State324, EnumValueDefType}:                 _GotoState351Action,
	{_State324, ImplicitEnumDefType}:              _GotoState147Action,
	{_State324, ExplicitEnumDefType}:              _GotoState142Action,
	{_State324, TraitDefType}:                     _GotoState151Action,
	{_State324, FuncTypeType}:                     _GotoState146Action,
	{_State326, IdentifierToken}:                  _GotoState155Action,
	{_State326, StructToken}:                      _GotoState33Action,
	{_State326, EnumToken}:                        _GotoState133Action,
	{_State326, TraitToken}:                       _GotoState140Action,
	{_State326, FuncToken}:                        _GotoState134Action,
	{_State326, LparenToken}:                      _GotoState137Action,
	{_State326, QuestionToken}:                    _GotoState138Action,
	{_State326, TildeTildeToken}:                  _GotoState139Action,
	{_State326, BitNegToken}:                      _GotoState132Action,
	{_State326, BitAndToken}:                      _GotoState131Action,
	{_State326, LexErrorToken}:                    _GotoState136Action,
	{_State326, AtomTypeType}:                     _GotoState141Action,
	{_State326, TraitableTypeType}:                _GotoState152Action,
	{_State326, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State326, ValueTypeType}:                    _GotoState352Action,
	{_State326, ImplicitStructDefType}:            _GotoState148Action,
	{_State326, ExplicitStructDefType}:            _GotoState144Action,
	{_State326, ImplicitEnumDefType}:              _GotoState147Action,
	{_State326, ExplicitEnumDefType}:              _GotoState142Action,
	{_State326, TraitDefType}:                     _GotoState151Action,
	{_State326, FuncTypeType}:                     _GotoState146Action,
	{_State328, IdentifierToken}:                  _GotoState155Action,
	{_State328, StructToken}:                      _GotoState33Action,
	{_State328, EnumToken}:                        _GotoState133Action,
	{_State328, TraitToken}:                       _GotoState140Action,
	{_State328, FuncToken}:                        _GotoState134Action,
	{_State328, LparenToken}:                      _GotoState137Action,
	{_State328, QuestionToken}:                    _GotoState138Action,
	{_State328, TildeTildeToken}:                  _GotoState139Action,
	{_State328, BitNegToken}:                      _GotoState132Action,
	{_State328, BitAndToken}:                      _GotoState131Action,
	{_State328, LexErrorToken}:                    _GotoState136Action,
	{_State328, AtomTypeType}:                     _GotoState141Action,
	{_State328, TraitableTypeType}:                _GotoState152Action,
	{_State328, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State328, ValueTypeType}:                    _GotoState253Action,
	{_State328, ImplicitStructDefType}:            _GotoState148Action,
	{_State328, ExplicitStructDefType}:            _GotoState144Action,
	{_State328, ImplicitEnumDefType}:              _GotoState147Action,
	{_State328, ExplicitEnumDefType}:              _GotoState142Action,
	{_State328, TraitDefType}:                     _GotoState151Action,
	{_State328, ReturnTypeType}:                   _GotoState353Action,
	{_State328, FuncTypeType}:                     _GotoState146Action,
	{_State329, IdentifierToken}:                  _GotoState260Action,
	{_State329, StructToken}:                      _GotoState33Action,
	{_State329, EnumToken}:                        _GotoState133Action,
	{_State329, TraitToken}:                       _GotoState140Action,
	{_State329, FuncToken}:                        _GotoState134Action,
	{_State329, LparenToken}:                      _GotoState137Action,
	{_State329, QuestionToken}:                    _GotoState138Action,
	{_State329, DotdotdotToken}:                   _GotoState259Action,
	{_State329, TildeTildeToken}:                  _GotoState139Action,
	{_State329, BitNegToken}:                      _GotoState132Action,
	{_State329, BitAndToken}:                      _GotoState131Action,
	{_State329, LexErrorToken}:                    _GotoState136Action,
	{_State329, AtomTypeType}:                     _GotoState141Action,
	{_State329, TraitableTypeType}:                _GotoState152Action,
	{_State329, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State329, ValueTypeType}:                    _GotoState264Action,
	{_State329, ImplicitStructDefType}:            _GotoState148Action,
	{_State329, ExplicitStructDefType}:            _GotoState144Action,
	{_State329, ImplicitEnumDefType}:              _GotoState147Action,
	{_State329, ExplicitEnumDefType}:              _GotoState142Action,
	{_State329, TraitDefType}:                     _GotoState151Action,
	{_State329, ParameterDeclType}:                _GotoState354Action,
	{_State329, FuncTypeType}:                     _GotoState146Action,
	{_State334, LparenToken}:                      _GotoState355Action,
	{_State336, IdentifierToken}:                  _GotoState135Action,
	{_State336, UnsafeToken}:                      _GotoState56Action,
	{_State336, StructToken}:                      _GotoState33Action,
	{_State336, EnumToken}:                        _GotoState133Action,
	{_State336, TraitToken}:                       _GotoState140Action,
	{_State336, FuncToken}:                        _GotoState271Action,
	{_State336, LparenToken}:                      _GotoState137Action,
	{_State336, QuestionToken}:                    _GotoState138Action,
	{_State336, TildeTildeToken}:                  _GotoState139Action,
	{_State336, BitNegToken}:                      _GotoState132Action,
	{_State336, BitAndToken}:                      _GotoState131Action,
	{_State336, LexErrorToken}:                    _GotoState136Action,
	{_State336, UnsafeStatementType}:              _GotoState153Action,
	{_State336, AtomTypeType}:                     _GotoState141Action,
	{_State336, TraitableTypeType}:                _GotoState152Action,
	{_State336, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State336, ValueTypeType}:                    _GotoState154Action,
	{_State336, FieldDefType}:                     _GotoState272Action,
	{_State336, ImplicitStructDefType}:            _GotoState148Action,
	{_State336, ExplicitStructDefType}:            _GotoState144Action,
	{_State336, ImplicitEnumDefType}:              _GotoState147Action,
	{_State336, ExplicitEnumDefType}:              _GotoState142Action,
	{_State336, TraitPropertyType}:                _GotoState356Action,
	{_State336, TraitDefType}:                     _GotoState151Action,
	{_State336, FuncTypeType}:                     _GotoState146Action,
	{_State336, MethodSignatureType}:              _GotoState273Action,
	{_State337, IdentifierToken}:                  _GotoState135Action,
	{_State337, UnsafeToken}:                      _GotoState56Action,
	{_State337, StructToken}:                      _GotoState33Action,
	{_State337, EnumToken}:                        _GotoState133Action,
	{_State337, TraitToken}:                       _GotoState140Action,
	{_State337, FuncToken}:                        _GotoState271Action,
	{_State337, LparenToken}:                      _GotoState137Action,
	{_State337, QuestionToken}:                    _GotoState138Action,
	{_State337, TildeTildeToken}:                  _GotoState139Action,
	{_State337, BitNegToken}:                      _GotoState132Action,
	{_State337, BitAndToken}:                      _GotoState131Action,
	{_State337, LexErrorToken}:                    _GotoState136Action,
	{_State337, UnsafeStatementType}:              _GotoState153Action,
	{_State337, AtomTypeType}:                     _GotoState141Action,
	{_State337, TraitableTypeType}:                _GotoState152Action,
	{_State337, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State337, ValueTypeType}:                    _GotoState154Action,
	{_State337, FieldDefType}:                     _GotoState272Action,
	{_State337, ImplicitStructDefType}:            _GotoState148Action,
	{_State337, ExplicitStructDefType}:            _GotoState144Action,
	{_State337, ImplicitEnumDefType}:              _GotoState147Action,
	{_State337, ExplicitEnumDefType}:              _GotoState142Action,
	{_State337, TraitPropertyType}:                _GotoState357Action,
	{_State337, TraitDefType}:                     _GotoState151Action,
	{_State337, FuncTypeType}:                     _GotoState146Action,
	{_State337, MethodSignatureType}:              _GotoState273Action,
	{_State339, DoToken}:                          _GotoState358Action,
	{_State344, CommaToken}:                       _GotoState307Action,
	{_State346, RbraceToken}:                      _GotoState359Action,
	{_State347, IdentifierToken}:                  _GotoState155Action,
	{_State347, StructToken}:                      _GotoState33Action,
	{_State347, EnumToken}:                        _GotoState133Action,
	{_State347, TraitToken}:                       _GotoState140Action,
	{_State347, FuncToken}:                        _GotoState134Action,
	{_State347, LparenToken}:                      _GotoState137Action,
	{_State347, QuestionToken}:                    _GotoState138Action,
	{_State347, TildeTildeToken}:                  _GotoState139Action,
	{_State347, BitNegToken}:                      _GotoState132Action,
	{_State347, BitAndToken}:                      _GotoState131Action,
	{_State347, LexErrorToken}:                    _GotoState136Action,
	{_State347, AtomTypeType}:                     _GotoState141Action,
	{_State347, TraitableTypeType}:                _GotoState152Action,
	{_State347, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State347, ValueTypeType}:                    _GotoState253Action,
	{_State347, ImplicitStructDefType}:            _GotoState148Action,
	{_State347, ExplicitStructDefType}:            _GotoState144Action,
	{_State347, ImplicitEnumDefType}:              _GotoState147Action,
	{_State347, ExplicitEnumDefType}:              _GotoState142Action,
	{_State347, TraitDefType}:                     _GotoState151Action,
	{_State347, ReturnTypeType}:                   _GotoState360Action,
	{_State347, FuncTypeType}:                     _GotoState146Action,
	{_State355, IdentifierToken}:                  _GotoState260Action,
	{_State355, StructToken}:                      _GotoState33Action,
	{_State355, EnumToken}:                        _GotoState133Action,
	{_State355, TraitToken}:                       _GotoState140Action,
	{_State355, FuncToken}:                        _GotoState134Action,
	{_State355, LparenToken}:                      _GotoState137Action,
	{_State355, QuestionToken}:                    _GotoState138Action,
	{_State355, DotdotdotToken}:                   _GotoState259Action,
	{_State355, TildeTildeToken}:                  _GotoState139Action,
	{_State355, BitNegToken}:                      _GotoState132Action,
	{_State355, BitAndToken}:                      _GotoState131Action,
	{_State355, LexErrorToken}:                    _GotoState136Action,
	{_State355, AtomTypeType}:                     _GotoState141Action,
	{_State355, TraitableTypeType}:                _GotoState152Action,
	{_State355, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State355, ValueTypeType}:                    _GotoState264Action,
	{_State355, ImplicitStructDefType}:            _GotoState148Action,
	{_State355, ExplicitStructDefType}:            _GotoState144Action,
	{_State355, ImplicitEnumDefType}:              _GotoState147Action,
	{_State355, ExplicitEnumDefType}:              _GotoState142Action,
	{_State355, TraitDefType}:                     _GotoState151Action,
	{_State355, ParameterDeclType}:                _GotoState262Action,
	{_State355, ParameterDeclsType}:               _GotoState263Action,
	{_State355, OptionalParameterDeclsType}:       _GotoState361Action,
	{_State355, FuncTypeType}:                     _GotoState146Action,
	{_State358, LbraceToken}:                      _GotoState104Action,
	{_State358, BlockBodyType}:                    _GotoState362Action,
	{_State360, LbraceToken}:                      _GotoState104Action,
	{_State360, BlockBodyType}:                    _GotoState363Action,
	{_State361, RparenToken}:                      _GotoState364Action,
	{_State364, IdentifierToken}:                  _GotoState155Action,
	{_State364, StructToken}:                      _GotoState33Action,
	{_State364, EnumToken}:                        _GotoState133Action,
	{_State364, TraitToken}:                       _GotoState140Action,
	{_State364, FuncToken}:                        _GotoState134Action,
	{_State364, LparenToken}:                      _GotoState137Action,
	{_State364, QuestionToken}:                    _GotoState138Action,
	{_State364, TildeTildeToken}:                  _GotoState139Action,
	{_State364, BitNegToken}:                      _GotoState132Action,
	{_State364, BitAndToken}:                      _GotoState131Action,
	{_State364, LexErrorToken}:                    _GotoState136Action,
	{_State364, AtomTypeType}:                     _GotoState141Action,
	{_State364, TraitableTypeType}:                _GotoState152Action,
	{_State364, TraitAlgebraTypeType}:             _GotoState150Action,
	{_State364, ValueTypeType}:                    _GotoState253Action,
	{_State364, ImplicitStructDefType}:            _GotoState148Action,
	{_State364, ExplicitStructDefType}:            _GotoState144Action,
	{_State364, ImplicitEnumDefType}:              _GotoState147Action,
	{_State364, ExplicitEnumDefType}:              _GotoState142Action,
	{_State364, TraitDefType}:                     _GotoState151Action,
	{_State364, ReturnTypeType}:                   _GotoState365Action,
	{_State364, FuncTypeType}:                     _GotoState146Action,
	{_State1, _EndMarker}:                         _ReduceNilToOptionalDefinitionsAction,
	{_State1, _WildcardMarker}:                    _ReduceNilToOptionalNewlinesAction,
	{_State5, _WildcardMarker}:                    _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State13, _EndMarker}:                        _ReduceNewlinesToOptionalDefinitionsAction,
	{_State13, _WildcardMarker}:                   _ReduceNewlinesToOptionalNewlinesAction,
	{_State14, _EndMarker}:                        _ReduceToSourceAction,
	{_State18, IdentifierToken}:                   _ReduceNilToOptionalReceiverAction,
	{_State19, _WildcardMarker}:                   _ReduceBitAndToPrefixUnaryOpAction,
	{_State20, _WildcardMarker}:                   _ReduceBitNegToPrefixUnaryOpAction,
	{_State21, _WildcardMarker}:                   _ReduceFalseToLiteralAction,
	{_State22, _WildcardMarker}:                   _ReduceFloatLiteralToLiteralAction,
	{_State24, _WildcardMarker}:                   _ReduceIdentifierToAtomExprAction,
	{_State25, _WildcardMarker}:                   _ReduceIntegerLiteralToLiteralAction,
	{_State26, _WildcardMarker}:                   _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State27, _WildcardMarker}:                   _ReduceLexErrorToAtomExprAction,
	{_State28, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State28, ColonToken}:                        _ReduceNilToOptionalExpressionAction,
	{_State29, _WildcardMarker}:                   _ReduceMulToPrefixUnaryOpAction,
	{_State30, _WildcardMarker}:                   _ReduceNotToPrefixUnaryOpAction,
	{_State31, _WildcardMarker}:                   _ReduceRuneLiteralToLiteralAction,
	{_State32, _WildcardMarker}:                   _ReduceStringLiteralToLiteralAction,
	{_State34, _WildcardMarker}:                   _ReduceSubToPrefixUnaryOpAction,
	{_State35, _WildcardMarker}:                   _ReduceTrueToLiteralAction,
	{_State36, _WildcardMarker}:                   _ReduceAccessExprToPostfixUnaryExprAction,
	{_State36, LparenToken}:                       _ReduceNilToOptionalGenericBindingAction,
	{_State37, _WildcardMarker}:                   _ReduceAddExprToCmpExprAction,
	{_State38, _WildcardMarker}:                   _ReduceAndExprToOrExprAction,
	{_State39, _WildcardMarker}:                   _ReduceAnonymousFuncExprToAtomExprAction,
	{_State40, _WildcardMarker}:                   _ReduceAnonymousStructExprToAtomExprAction,
	{_State41, _WildcardMarker}:                   _ReduceAtomExprToAccessExprAction,
	{_State42, _WildcardMarker}:                   _ReduceBlockExprToAtomExprAction,
	{_State43, _WildcardMarker}:                   _ReduceCallExprToAccessExprAction,
	{_State44, _WildcardMarker}:                   _ReduceCmpExprToAndExprAction,
	{_State46, _WildcardMarker}:                   _ReduceLiteralToAtomExprAction,
	{_State47, _WildcardMarker}:                   _ReduceMulExprToAddExprAction,
	{_State49, _WildcardMarker}:                   _ReduceToSequenceExprAction,
	{_State50, _WildcardMarker}:                   _ReducePostfixUnaryExprToPrefixUnaryExprAction,
	{_State51, _WildcardMarker}:                   _ReducePrefixUnaryExprToMulExprAction,
	{_State52, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State53, _EndMarker}:                        _ReduceSequenceExprToExpressionAction,
	{_State54, _EndMarker}:                        _ReduceCommentToLexInternalTokensAction,
	{_State55, _EndMarker}:                        _ReduceSpacesToLexInternalTokensAction,
	{_State57, _WildcardMarker}:                   _ReduceNilToDefinitionsAction,
	{_State58, _EndMarker}:                        _ReduceNilToOptionalNewlinesAction,
	{_State59, _WildcardMarker}:                   _ReduceNamedFuncDefToDefinitionAction,
	{_State60, _WildcardMarker}:                   _ReducePackageDefToDefinitionAction,
	{_State61, _WildcardMarker}:                   _ReduceTypeDefToDefinitionAction,
	{_State62, _WildcardMarker}:                   _ReduceUnsafeStatementToDefinitionAction,
	{_State63, _WildcardMarker}:                   _ReduceNoSpecToPackageDefAction,
	{_State64, _WildcardMarker}:                   _ReduceNilToOptionalGenericParametersAction,
	{_State67, RparenToken}:                       _ReduceNilToOptionalParameterDefsAction,
	{_State68, _WildcardMarker}:                   _ReduceIdentifierToAtomExprAction,
	{_State69, _WildcardMarker}:                   _ReduceArgumentToArgumentsAction,
	{_State71, _WildcardMarker}:                   _ReduceColonExpressionsToArgumentAction,
	{_State72, _WildcardMarker}:                   _ReducePositionalToArgumentAction,
	{_State72, ColonToken}:                        _ReduceExpressionToOptionalExpressionAction,
	{_State74, RparenToken}:                       _ReduceNilToOptionalExplicitFieldDefsAction,
	{_State75, RbracketToken}:                     _ReduceNilToOptionalGenericArgumentsAction,
	{_State77, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State77, ColonToken}:                        _ReduceNilToOptionalExpressionAction,
	{_State78, _WildcardMarker}:                   _ReduceQuestionToPostfixUnaryExprAction,
	{_State80, _WildcardMarker}:                   _ReduceAddToAddOpAction,
	{_State81, _WildcardMarker}:                   _ReduceBitOrToAddOpAction,
	{_State82, _WildcardMarker}:                   _ReduceBitXorToAddOpAction,
	{_State83, _WildcardMarker}:                   _ReduceSubToAddOpAction,
	{_State84, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State85, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State86, _WildcardMarker}:                   _ReduceEqualToCmpOpAction,
	{_State87, _WildcardMarker}:                   _ReduceGreaterToCmpOpAction,
	{_State88, _WildcardMarker}:                   _ReduceGreaterOrEqualToCmpOpAction,
	{_State89, _WildcardMarker}:                   _ReduceLessToCmpOpAction,
	{_State90, _WildcardMarker}:                   _ReduceLessOrEqualToCmpOpAction,
	{_State91, _WildcardMarker}:                   _ReduceNotEqualToCmpOpAction,
	{_State92, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State93, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State93, ColonToken}:                        _ReduceNilToOptionalExpressionAction,
	{_State94, _WildcardMarker}:                   _ReduceBitAndToMulOpAction,
	{_State95, _WildcardMarker}:                   _ReduceBitLshiftToMulOpAction,
	{_State96, _WildcardMarker}:                   _ReduceBitRshiftToMulOpAction,
	{_State97, _WildcardMarker}:                   _ReduceDivToMulOpAction,
	{_State98, _WildcardMarker}:                   _ReduceModToMulOpAction,
	{_State99, _WildcardMarker}:                   _ReduceMulToMulOpAction,
	{_State100, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State102, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State103, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State104, _WildcardMarker}:                  _ReduceEmptyListToStatementsAction,
	{_State105, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State106, _WildcardMarker}:                  _ReduceToBlockExprAction,
	{_State107, _EndMarker}:                       _ReduceIfExprToExpressionAction,
	{_State108, _EndMarker}:                       _ReduceLoopExprToExpressionAction,
	{_State109, _EndMarker}:                       _ReduceSwitchExprToExpressionAction,
	{_State110, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State112, _WildcardMarker}:                  _ReducePrefixOpToPrefixUnaryExprAction,
	{_State114, _EndMarker}:                       _ReduceNewlinesToOptionalNewlinesAction,
	{_State115, _EndMarker}:                       _ReduceDefinitionsToOptionalDefinitionsAction,
	{_State116, _WildcardMarker}:                  _ReduceEmptyListToPackageStatementsAction,
	{_State117, RbracketToken}:                    _ReduceNilToOptionalGenericParameterDefsAction,
	{_State122, LparenToken}:                      _ReduceNilToOptionalGenericParametersAction,
	{_State124, _WildcardMarker}:                  _ReduceParameterDefToParameterDefsAction,
	{_State125, RparenToken}:                      _ReduceParameterDefsToOptionalParameterDefsAction,
	{_State126, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State127, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State127, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State128, _WildcardMarker}:                  _ReduceImplicitToAnonymousStructExprAction,
	{_State129, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State129, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State129, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State129, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State129, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State130, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State130, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State130, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State130, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State130, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State135, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State136, _WildcardMarker}:                  _ReduceLexErrorToAtomTypeAction,
	{_State137, RparenToken}:                      _ReduceNilToOptionalImplicitFieldDefsAction,
	{_State138, _EndMarker}:                       _ReduceInferredToValueTypeAction,
	{_State141, _WildcardMarker}:                  _ReduceAtomTypeToTraitableTypeAction,
	{_State142, _WildcardMarker}:                  _ReduceExplicitEnumDefToAtomTypeAction,
	{_State143, RparenToken}:                      _ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction,
	{_State144, _WildcardMarker}:                  _ReduceExplicitStructDefToAtomTypeAction,
	{_State145, _WildcardMarker}:                  _ReduceFieldDefToExplicitFieldDefsAction,
	{_State146, _EndMarker}:                       _ReduceFuncTypeToValueTypeAction,
	{_State147, _WildcardMarker}:                  _ReduceImplicitEnumDefToAtomTypeAction,
	{_State148, _WildcardMarker}:                  _ReduceImplicitStructDefToAtomTypeAction,
	{_State150, _EndMarker}:                       _ReduceTraitAlgebraTypeToValueTypeAction,
	{_State151, _WildcardMarker}:                  _ReduceTraitDefToAtomTypeAction,
	{_State152, _WildcardMarker}:                  _ReduceTraitableTypeToTraitAlgebraTypeAction,
	{_State153, _WildcardMarker}:                  _ReduceUnsafeStatementToFieldDefAction,
	{_State154, _WildcardMarker}:                  _ReduceImplicitToFieldDefAction,
	{_State155, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State156, RbracketToken}:                    _ReduceGenericArgumentsToOptionalGenericArgumentsAction,
	{_State158, _WildcardMarker}:                  _ReduceValueTypeToGenericArgumentsAction,
	{_State159, _WildcardMarker}:                  _ReduceAccessToAccessExprAction,
	{_State161, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State161, RparenToken}:                      _ReduceNilToOptionalArgumentsAction,
	{_State161, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State162, _WildcardMarker}:                  _ReduceOpToAddExprAction,
	{_State163, _WildcardMarker}:                  _ReduceOpToAndExprAction,
	{_State164, _WildcardMarker}:                  _ReduceOpToCmpExprAction,
	{_State166, _WildcardMarker}:                  _ReduceOpToMulExprAction,
	{_State167, _WildcardMarker}:                  _ReduceInfiniteToLoopExprAction,
	{_State168, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State169, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State169, SemicolonToken}:                   _ReduceNilToOptionalSequenceExprAction,
	{_State172, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State174, _WildcardMarker}:                  _ReduceOpToOrExprAction,
	{_State176, _WildcardMarker}:                  _ReduceAddToDefinitionsAction,
	{_State178, _WildcardMarker}:                  _ReduceUnconstrainedToGenericParameterDefAction,
	{_State179, _WildcardMarker}:                  _ReduceGenericParameterDefToGenericParameterDefsAction,
	{_State180, RbracketToken}:                    _ReduceGenericParameterDefsToOptionalGenericParameterDefsAction,
	{_State182, _EndMarker}:                       _ReduceAliasToTypeDefAction,
	{_State183, _WildcardMarker}:                  _ReduceDefinitionToTypeDefAction,
	{_State185, _WildcardMarker}:                  _ReduceArgToParameterDefAction,
	{_State186, IdentifierToken}:                  _ReduceReceiverToOptionalReceiverAction,
	{_State188, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State190, _WildcardMarker}:                  _ReduceNamedToArgumentAction,
	{_State191, _WildcardMarker}:                  _ReduceAddToArgumentsAction,
	{_State192, _WildcardMarker}:                  _ReduceExpressionToOptionalExpressionAction,
	{_State193, _WildcardMarker}:                  _ReduceAddToColonExpressionsAction,
	{_State194, _WildcardMarker}:                  _ReducePairToColonExpressionsAction,
	{_State195, _EndMarker}:                       _ReduceReferenceToValueTypeAction,
	{_State196, _WildcardMarker}:                  _ReducePublicMethodsTraitToTraitableTypeAction,
	{_State198, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State199, _WildcardMarker}:                  _ReduceNamedToAtomTypeAction,
	{_State200, _WildcardMarker}:                  _ReduceExplicitToFieldDefAction,
	{_State202, _WildcardMarker}:                  _ReduceFieldDefToImplicitFieldDefsAction,
	{_State202, OrToken}:                          _ReduceFieldDefToEnumValueDefAction,
	{_State204, RparenToken}:                      _ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction,
	{_State206, _WildcardMarker}:                  _ReducePublicTraitToTraitableTypeAction,
	{_State207, RparenToken}:                      _ReduceNilToOptionalTraitPropertiesAction,
	{_State210, _WildcardMarker}:                  _ReduceToExplicitStructDefAction,
	{_State215, _WildcardMarker}:                  _ReduceBindingToOptionalGenericBindingAction,
	{_State216, _WildcardMarker}:                  _ReduceIndexToAccessExprAction,
	{_State217, RparenToken}:                      _ReduceArgumentsToOptionalArgumentsAction,
	{_State219, _WildcardMarker}:                  _ReduceExplicitToAnonymousStructExprAction,
	{_State220, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State223, DoToken}:                          _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State225, _WildcardMarker}:                  _ReduceNoElseToIfExprAction,
	{_State226, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State227, _WildcardMarker}:                  _ReduceBreakToJumpTypeAction,
	{_State228, _WildcardMarker}:                  _ReduceContinueToJumpTypeAction,
	{_State229, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State230, _EndMarker}:                       _ReduceToBlockBodyAction,
	{_State231, _WildcardMarker}:                  _ReduceReturnToJumpTypeAction,
	{_State232, _WildcardMarker}:                  _ReduceAccessExprToPostfixUnaryExprAction,
	{_State232, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State233, _WildcardMarker}:                  _ReduceExpressionToExpressionsAction,
	{_State234, _WildcardMarker}:                  _ReduceExpressionOrImplicitStructToStatementBodyAction,
	{_State235, _WildcardMarker}:                  _ReduceJumpStatementToStatementBodyAction,
	{_State236, _WildcardMarker}:                  _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State237, _WildcardMarker}:                  _ReduceAddToStatementsAction,
	{_State239, _WildcardMarker}:                  _ReduceUnsafeStatementToStatementBodyAction,
	{_State242, _EndMarker}:                       _ReduceWithSpecToPackageDefAction,
	{_State243, _WildcardMarker}:                  _ReduceAddToPackageStatementsAction,
	{_State245, _WildcardMarker}:                  _ReduceToPackageStatementBodyAction,
	{_State246, _WildcardMarker}:                  _ReduceConstrainedToGenericParameterDefAction,
	{_State248, _WildcardMarker}:                  _ReduceGenericToOptionalGenericParametersAction,
	{_State250, _WildcardMarker}:                  _ReduceVarargToParameterDefAction,
	{_State251, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State253, _EndMarker}:                       _ReduceValueTypeToReturnTypeAction,
	{_State254, _WildcardMarker}:                  _ReduceAddToParameterDefsAction,
	{_State257, _WildcardMarker}:                  _ReduceFieldDefToEnumValueDefAction,
	{_State260, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State262, _WildcardMarker}:                  _ReduceParameterDeclToParameterDeclsAction,
	{_State263, RparenToken}:                      _ReduceParameterDeclsToOptionalParameterDeclsAction,
	{_State264, _WildcardMarker}:                  _ReduceUnamedToParameterDeclAction,
	{_State268, _WildcardMarker}:                  _ReduceToImplicitEnumDefAction,
	{_State270, _WildcardMarker}:                  _ReduceToImplicitStructDefAction,
	{_State272, _WildcardMarker}:                  _ReduceFieldDefToTraitPropertyAction,
	{_State273, _WildcardMarker}:                  _ReduceMethodSignatureToTraitPropertyAction,
	{_State275, RparenToken}:                      _ReduceTraitPropertiesToOptionalTraitPropertiesAction,
	{_State276, _WildcardMarker}:                  _ReduceTraitPropertyToTraitPropertiesAction,
	{_State277, _WildcardMarker}:                  _ReduceExplicitToExplicitFieldDefsAction,
	{_State278, _WildcardMarker}:                  _ReduceImplicitToExplicitFieldDefsAction,
	{_State279, _WildcardMarker}:                  _ReduceUnionToTraitAlgebraTypeAction,
	{_State280, _WildcardMarker}:                  _ReduceIntersectToTraitAlgebraTypeAction,
	{_State281, _WildcardMarker}:                  _ReduceDifferenceToTraitAlgebraTypeAction,
	{_State282, _WildcardMarker}:                  _ReduceAddToGenericArgumentsAction,
	{_State283, _WildcardMarker}:                  _ReduceToCallExprAction,
	{_State284, _EndMarker}:                       _ReduceDoWhileToLoopExprAction,
	{_State286, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State286, DoToken}:                          _ReduceNilToOptionalSequenceExprAction,
	{_State287, _EndMarker}:                       _ReduceWhileToLoopExprAction,
	{_State289, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State290, NewlinesToken}:                    _ReduceAsyncToStatementBodyAction,
	{_State290, SemicolonToken}:                   _ReduceAsyncToStatementBodyAction,
	{_State290, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State291, NewlinesToken}:                    _ReduceDeferToStatementBodyAction,
	{_State291, SemicolonToken}:                   _ReduceDeferToStatementBodyAction,
	{_State291, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State292, _WildcardMarker}:                  _ReduceAddAssignToBinaryOpAssignAction,
	{_State293, _WildcardMarker}:                  _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State294, _WildcardMarker}:                  _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State295, _WildcardMarker}:                  _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State296, _WildcardMarker}:                  _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State297, _WildcardMarker}:                  _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State298, _WildcardMarker}:                  _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State299, _WildcardMarker}:                  _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State300, _WildcardMarker}:                  _ReduceDivAssignToBinaryOpAssignAction,
	{_State301, _WildcardMarker}:                  _ReduceModAssignToBinaryOpAssignAction,
	{_State302, _WildcardMarker}:                  _ReduceMulAssignToBinaryOpAssignAction,
	{_State303, _WildcardMarker}:                  _ReduceSubAssignToBinaryOpAssignAction,
	{_State304, _WildcardMarker}:                  _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State305, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State306, _WildcardMarker}:                  _ReduceUnaryOpAssignStatementToStatementBodyAction,
	{_State307, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State308, _WildcardMarker}:                  _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State309, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State309, NewlinesToken}:                    _ReduceNilToOptionalExpressionsAction,
	{_State309, SemicolonToken}:                   _ReduceNilToOptionalExpressionsAction,
	{_State310, _WildcardMarker}:                  _ReduceImplicitToStatementAction,
	{_State311, _WildcardMarker}:                  _ReduceExplicitToStatementAction,
	{_State313, _WildcardMarker}:                  _ReduceToUnsafeStatementAction,
	{_State314, _WildcardMarker}:                  _ReduceImplicitToPackageStatementAction,
	{_State315, _WildcardMarker}:                  _ReduceExplicitToPackageStatementAction,
	{_State316, _WildcardMarker}:                  _ReduceAddToGenericParameterDefsAction,
	{_State317, _EndMarker}:                       _ReduceConstrainedDefToTypeDefAction,
	{_State319, _WildcardMarker}:                  _ReduceToAnonymousFuncExprAction,
	{_State322, _WildcardMarker}:                  _ReduceToExplicitEnumDefAction,
	{_State325, _WildcardMarker}:                  _ReduceUnnamedVarargToParameterDeclAction,
	{_State327, _WildcardMarker}:                  _ReduceArgToParameterDeclAction,
	{_State328, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State330, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State331, _WildcardMarker}:                  _ReduceDefaultToEnumValueDefAction,
	{_State332, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State333, _WildcardMarker}:                  _ReduceAddToImplicitFieldDefsAction,
	{_State335, _WildcardMarker}:                  _ReduceToTraitDefAction,
	{_State338, _EndMarker}:                       _ReduceIteratorToLoopExprAction,
	{_State340, _EndMarker}:                       _ReduceIfElseToIfExprAction,
	{_State341, _EndMarker}:                       _ReduceMultiIfElseToIfExprAction,
	{_State342, _WildcardMarker}:                  _ReduceBinaryOpAssignStatementToStatementBodyAction,
	{_State343, _WildcardMarker}:                  _ReduceAddToExpressionsAction,
	{_State344, _WildcardMarker}:                  _ReduceExpressionsToOptionalExpressionsAction,
	{_State345, _WildcardMarker}:                  _ReduceToJumpStatementAction,
	{_State347, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State348, RparenToken}:                      _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State349, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State349, RparenToken}:                      _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State350, RparenToken}:                      _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State351, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State351, RparenToken}:                      _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State352, _WildcardMarker}:                  _ReduceVarargToParameterDeclAction,
	{_State353, _EndMarker}:                       _ReduceToFuncTypeAction,
	{_State354, _WildcardMarker}:                  _ReduceAddToParameterDeclsAction,
	{_State355, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State356, _WildcardMarker}:                  _ReduceExplicitToTraitPropertiesAction,
	{_State357, _WildcardMarker}:                  _ReduceImplicitToTraitPropertiesAction,
	{_State359, _EndMarker}:                       _ReduceToSwitchExprAction,
	{_State362, _EndMarker}:                       _ReduceForToLoopExprAction,
	{_State363, _EndMarker}:                       _ReduceToNamedFuncDefAction,
	{_State364, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State365, _WildcardMarker}:                  _ReduceToMethodSignatureAction,
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.source
    Reduce:
      * -> [optional_newlines]
      $ -> [optional_definitions]
    Goto:
      NEWLINES -> State 13
      source -> State 7
      optional_definitions -> State 14
      optional_newlines -> State 15

  State 2:
    Kernel Items:
      #accept: ^.package_def
    Reduce:
      (nil)
    Goto:
      PACKAGE -> State 16
      package_def -> State 8

  State 3:
    Kernel Items:
      #accept: ^.type_def
    Reduce:
      (nil)
    Goto:
      TYPE -> State 17
      type_def -> State 9

  State 4:
    Kernel Items:
      #accept: ^.named_func_def
    Reduce:
      (nil)
    Goto:
      FUNC -> State 18
      named_func_def -> State 10

  State 5:
    Kernel Items:
      #accept: ^.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      expression -> State 11
      optional_label_decl -> State 48
      sequence_expr -> State 53
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 6:
    Kernel Items:
      #accept: ^.lex_internal_tokens
    Reduce:
      (nil)
    Goto:
      SPACES -> State 55
      COMMENT -> State 54
      lex_internal_tokens -> State 12

  State 7:
    Kernel Items:
      #accept: ^ source., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      #accept: ^ package_def., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 9:
    Kernel Items:
      #accept: ^ type_def., $
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
      optional_definitions: NEWLINES., $
      optional_newlines: NEWLINES., *
    Reduce:
      * -> [optional_newlines]
      $ -> [optional_definitions]
    Goto:
      (nil)

  State 14:
    Kernel Items:
      source: optional_definitions., $
    Reduce:
      $ -> [source]
    Goto:
      (nil)

  State 15:
    Kernel Items:
      optional_definitions: optional_newlines.definitions optional_newlines
    Reduce:
      (nil)
    Goto:
      PACKAGE -> State 16
      UNSAFE -> State 56
      TYPE -> State 17
      FUNC -> State 18
      definitions -> State 58
      definition -> State 57
      unsafe_statement -> State 62
      type_def -> State 61
      named_func_def -> State 59
      package_def -> State 60

  State 16:
    Kernel Items:
      package_def: PACKAGE.IDENTIFIER
      package_def: PACKAGE.IDENTIFIER LPAREN package_statements RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 63

  State 17:
    Kernel Items:
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE.IDENTIFIER EQUAL value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 64

  State 18:
    Kernel Items:
      named_func_def: FUNC.optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      LPAREN -> State 65
      optional_receiver -> State 66

  State 19:
    Kernel Items:
      prefix_unary_op: BIT_AND., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 20:
    Kernel Items:
      prefix_unary_op: BIT_NEG., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      literal: FALSE., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 22:
    Kernel Items:
      literal: FLOAT_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      anonymous_func_expr: FUNC.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 67

  State 24:
    Kernel Items:
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 25:
    Kernel Items:
      literal: INTEGER_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      optional_label_decl: LABEL_DECL., *
    Reduce:
      * -> [optional_label_decl]
    Goto:
      (nil)

  State 27:
    Kernel Items:
      atom_expr: LEX_ERROR., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 28:
    Kernel Items:
      anonymous_struct_expr: LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 68
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      expression -> State 72
      optional_label_decl -> State 48
      sequence_expr -> State 53
      block_expr -> State 42
      call_expr -> State 43
      arguments -> State 70
      argument -> State 69
      colon_expressions -> State 71
      optional_expression -> State 73
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 29:
    Kernel Items:
      prefix_unary_op: MUL., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 30:
    Kernel Items:
      prefix_unary_op: NOT., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 31:
    Kernel Items:
      literal: RUNE_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      literal: STRING_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 33:
    Kernel Items:
      explicit_struct_def: STRUCT.LPAREN optional_explicit_field_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 74

  State 34:
    Kernel Items:
      prefix_unary_op: SUB., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 35:
    Kernel Items:
      literal: TRUE., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 36:
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
      LBRACKET -> State 77
      DOT -> State 76
      QUESTION -> State 78
      DOLLAR_LBRACKET -> State 75
      optional_generic_binding -> State 79

  State 37:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 80
      SUB -> State 83
      BIT_XOR -> State 82
      BIT_OR -> State 81
      add_op -> State 84

  State 38:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 85

  State 39:
    Kernel Items:
      atom_expr: anonymous_func_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 40:
    Kernel Items:
      atom_expr: anonymous_struct_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 41:
    Kernel Items:
      access_expr: atom_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 42:
    Kernel Items:
      atom_expr: block_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 43:
    Kernel Items:
      access_expr: call_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 44:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 86
      NOT_EQUAL -> State 91
      LESS -> State 89
      LESS_OR_EQUAL -> State 90
      GREATER -> State 87
      GREATER_OR_EQUAL -> State 88
      cmp_op -> State 92

  State 45:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_def.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 93

  State 46:
    Kernel Items:
      atom_expr: literal., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 47:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 99
      DIV -> State 97
      MOD -> State 98
      BIT_AND -> State 94
      BIT_LSHIFT -> State 95
      BIT_RSHIFT -> State 96
      mul_op -> State 100

  State 48:
    Kernel Items:
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      IF -> State 103
      SWITCH -> State 105
      FOR -> State 102
      DO -> State 101
      LBRACE -> State 104
      if_expr -> State 107
      switch_expr -> State 109
      loop_expr -> State 108
      block_body -> State 106

  State 49:
    Kernel Items:
      sequence_expr: or_expr., *
      or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    Goto:
      OR -> State 110

  State 50:
    Kernel Items:
      prefix_unary_expr: postfix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 51:
    Kernel Items:
      mul_expr: prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 52:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op.prefix_unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      optional_label_decl -> State 111
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 112
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 53:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 54:
    Kernel Items:
      lex_internal_tokens: COMMENT., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 55:
    Kernel Items:
      lex_internal_tokens: SPACES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 56:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 113

  State 57:
    Kernel Items:
      definitions: definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      optional_definitions: optional_newlines definitions.optional_newlines
      definitions: definitions.NEWLINES definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      NEWLINES -> State 114
      optional_newlines -> State 115

  State 59:
    Kernel Items:
      definition: named_func_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 60:
    Kernel Items:
      definition: package_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 61:
    Kernel Items:
      definition: type_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      definition: unsafe_statement., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 63:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER., *
      package_def: PACKAGE IDENTIFIER.LPAREN package_statements RPAREN
    Reduce:
      * -> [package_def]
    Goto:
      LPAREN -> State 116

  State 64:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE IDENTIFIER.EQUAL value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 117
      EQUAL -> State 118
      optional_generic_parameters -> State 119

  State 65:
    Kernel Items:
      optional_receiver: LPAREN.parameter_def RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 120
      parameter_def -> State 121

  State 66:
    Kernel Items:
      named_func_def: FUNC optional_receiver.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 122

  State 67:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 120
      parameter_def -> State 124
      parameter_defs -> State 125
      optional_parameter_defs -> State 123

  State 68:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expression
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      ASSIGN -> State 126

  State 69:
    Kernel Items:
      arguments: argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 70:
    Kernel Items:
      arguments: arguments.COMMA argument
      anonymous_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 128
      COMMA -> State 127

  State 71:
    Kernel Items:
      argument: colon_expressions., *
      colon_expressions: colon_expressions.COLON optional_expression
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 129

  State 72:
    Kernel Items:
      argument: expression., *
      optional_expression: expression., COLON
    Reduce:
      * -> [argument]
      COLON -> [optional_expression]
    Goto:
      (nil)

  State 73:
    Kernel Items:
      colon_expressions: optional_expression.COLON optional_expression
    Reduce:
      (nil)
    Goto:
      COLON -> State 130

  State 74:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN.optional_explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      IDENTIFIER -> State 135
      UNSAFE -> State 56
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      unsafe_statement -> State 153
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 154
      field_def -> State 145
      implicit_struct_def -> State 148
      explicit_field_defs -> State 143
      optional_explicit_field_defs -> State 149
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 75:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      optional_generic_arguments -> State 157
      generic_arguments -> State 156
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 158
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 76:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 159

  State 77:
    Kernel Items:
      access_expr: access_expr LBRACKET.argument RBRACKET
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 68
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      expression -> State 72
      optional_label_decl -> State 48
      sequence_expr -> State 53
      block_expr -> State 42
      call_expr -> State 43
      argument -> State 160
      colon_expressions -> State 71
      optional_expression -> State 73
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 78:
    Kernel Items:
      postfix_unary_expr: access_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 79:
    Kernel Items:
      call_expr: access_expr optional_generic_binding.LPAREN optional_arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 161

  State 80:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 81:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 82:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 84:
    Kernel Items:
      add_expr: add_expr add_op.mul_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      optional_label_decl -> State 111
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 162
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 85:
    Kernel Items:
      and_expr: and_expr AND.cmp_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      optional_label_decl -> State 111
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 163
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 86:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 87:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 90:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 91:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 92:
    Kernel Items:
      cmp_expr: cmp_expr cmp_op.add_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      optional_label_decl -> State 111
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 164
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 93:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_def LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 68
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      expression -> State 72
      optional_label_decl -> State 48
      sequence_expr -> State 53
      block_expr -> State 42
      call_expr -> State 43
      arguments -> State 165
      argument -> State 69
      colon_expressions -> State 71
      optional_expression -> State 73
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 94:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 95:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 96:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 97:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 98:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 99:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 100:
    Kernel Items:
      mul_expr: mul_expr mul_op.prefix_unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      optional_label_decl -> State 111
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 166
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 101:
    Kernel Items:
      loop_expr: DO.block_body
      loop_expr: DO.block_body FOR sequence_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 104
      block_body -> State 167

  State 102:
    Kernel Items:
      loop_expr: FOR.sequence_expr DO block_body
      loop_expr: FOR.IN sequence_expr DO block_body
      loop_expr: FOR.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      IN -> State 168
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      SEMICOLON -> State 169
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      optional_label_decl -> State 111
      sequence_expr -> State 170
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 103:
    Kernel Items:
      if_expr: IF.sequence_expr block_body
      if_expr: IF.sequence_expr block_body ELSE block_body
      if_expr: IF.sequence_expr block_body ELSE if_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      optional_label_decl -> State 111
      sequence_expr -> State 171
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 104:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 172

  State 105:
    Kernel Items:
      switch_expr: SWITCH.sequence_expr LBRACE CASE DEFAULT RBRACE
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      optional_label_decl -> State 111
      sequence_expr -> State 173
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 106:
    Kernel Items:
      block_expr: optional_label_decl block_body., *
    Reduce:
      * -> [block_expr]
    Goto:
      (nil)

  State 107:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 108:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 109:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 110:
    Kernel Items:
      or_expr: or_expr OR.and_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      optional_label_decl -> State 111
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 174
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 111:
    Kernel Items:
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 104
      block_body -> State 106

  State 112:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 113:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 175

  State 114:
    Kernel Items:
      optional_newlines: NEWLINES., $
      definitions: definitions NEWLINES.definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      PACKAGE -> State 16
      UNSAFE -> State 56
      TYPE -> State 17
      FUNC -> State 18
      definition -> State 176
      unsafe_statement -> State 62
      type_def -> State 61
      named_func_def -> State 59
      package_def -> State 60

  State 115:
    Kernel Items:
      optional_definitions: optional_newlines definitions optional_newlines., $
    Reduce:
      $ -> [optional_definitions]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN.package_statements RPAREN
    Reduce:
      * -> [package_statements]
    Goto:
      package_statements -> State 177

  State 117:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.optional_generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 178
      generic_parameter_def -> State 179
      generic_parameter_defs -> State 180
      optional_generic_parameter_defs -> State 181

  State 118:
    Kernel Items:
      type_def: TYPE IDENTIFIER EQUAL.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 182
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 119:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 183
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 120:
    Kernel Items:
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      DOTDOTDOT -> State 184
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 185
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 121:
    Kernel Items:
      optional_receiver: LPAREN parameter_def.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 186

  State 122:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 117
      optional_generic_parameters -> State 187

  State 123:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 188

  State 124:
    Kernel Items:
      parameter_defs: parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      parameter_defs: parameter_defs.COMMA parameter_def
      optional_parameter_defs: parameter_defs., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      COMMA -> State 189

  State 126:
    Kernel Items:
      argument: IDENTIFIER ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      expression -> State 190
      optional_label_decl -> State 48
      sequence_expr -> State 53
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 127:
    Kernel Items:
      arguments: arguments COMMA.argument
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 68
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      expression -> State 72
      optional_label_decl -> State 48
      sequence_expr -> State 53
      block_expr -> State 42
      call_expr -> State 43
      argument -> State 191
      colon_expressions -> State 71
      optional_expression -> State 73
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 128:
    Kernel Items:
      anonymous_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 129:
    Kernel Items:
      colon_expressions: colon_expressions COLON.optional_expression
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_expression]
      RBRACKET -> [optional_expression]
      COMMA -> [optional_expression]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      expression -> State 192
      optional_label_decl -> State 48
      sequence_expr -> State 53
      block_expr -> State 42
      call_expr -> State 43
      optional_expression -> State 193
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 130:
    Kernel Items:
      colon_expressions: optional_expression COLON.optional_expression
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_expression]
      RBRACKET -> [optional_expression]
      COMMA -> [optional_expression]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      expression -> State 192
      optional_label_decl -> State 48
      sequence_expr -> State 53
      block_expr -> State 42
      call_expr -> State 43
      optional_expression -> State 194
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 131:
    Kernel Items:
      value_type: BIT_AND.trait_algebra_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      LPAREN -> State 137
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 195
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151

  State 132:
    Kernel Items:
      traitable_type: BIT_NEG.atom_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      LPAREN -> State 137
      LEX_ERROR -> State 136
      atom_type -> State 196
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151

  State 133:
    Kernel Items:
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 197

  State 134:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 198

  State 135:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      DOLLAR_LBRACKET -> State 75
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      optional_generic_binding -> State 199
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 200
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 136:
    Kernel Items:
      atom_type: LEX_ERROR., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 137:
    Kernel Items:
      implicit_struct_def: LPAREN.optional_implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      IDENTIFIER -> State 135
      UNSAFE -> State 56
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      unsafe_statement -> State 153
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 154
      field_def -> State 202
      implicit_field_defs -> State 204
      optional_implicit_field_defs -> State 205
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      enum_value_def -> State 201
      implicit_enum_value_defs -> State 203
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 138:
    Kernel Items:
      value_type: QUESTION., $
    Reduce:
      $ -> [value_type]
    Goto:
      (nil)

  State 139:
    Kernel Items:
      traitable_type: TILDE_TILDE.atom_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      LPAREN -> State 137
      LEX_ERROR -> State 136
      atom_type -> State 206
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151

  State 140:
    Kernel Items:
      trait_def: TRAIT.LPAREN optional_trait_properties RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 207

  State 141:
    Kernel Items:
      traitable_type: atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 142:
    Kernel Items:
      atom_type: explicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 143:
    Kernel Items:
      explicit_field_defs: explicit_field_defs.NEWLINES field_def
      explicit_field_defs: explicit_field_defs.COMMA field_def
      optional_explicit_field_defs: explicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      NEWLINES -> State 209
      COMMA -> State 208

  State 144:
    Kernel Items:
      atom_type: explicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 145:
    Kernel Items:
      explicit_field_defs: field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 146:
    Kernel Items:
      value_type: func_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      (nil)

  State 147:
    Kernel Items:
      atom_type: implicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 148:
    Kernel Items:
      atom_type: implicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 149:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 210

  State 150:
    Kernel Items:
      trait_algebra_type: trait_algebra_type.MUL traitable_type
      trait_algebra_type: trait_algebra_type.ADD traitable_type
      trait_algebra_type: trait_algebra_type.SUB traitable_type
      value_type: trait_algebra_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      ADD -> State 211
      SUB -> State 213
      MUL -> State 212

  State 151:
    Kernel Items:
      atom_type: trait_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 152:
    Kernel Items:
      trait_algebra_type: traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 153:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 154:
    Kernel Items:
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 155:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 75
      optional_generic_binding -> State 199

  State 156:
    Kernel Items:
      optional_generic_arguments: generic_arguments., RBRACKET
      generic_arguments: generic_arguments.COMMA value_type
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 214

  State 157:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 215

  State 158:
    Kernel Items:
      generic_arguments: value_type., *
    Reduce:
      * -> [generic_arguments]
    Goto:
      (nil)

  State 159:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 160:
    Kernel Items:
      access_expr: access_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 216

  State 161:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN.optional_arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_arguments]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 68
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      expression -> State 72
      optional_label_decl -> State 48
      sequence_expr -> State 53
      block_expr -> State 42
      call_expr -> State 43
      optional_arguments -> State 218
      arguments -> State 217
      argument -> State 69
      colon_expressions -> State 71
      optional_expression -> State 73
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 162:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 99
      DIV -> State 97
      MOD -> State 98
      BIT_AND -> State 94
      BIT_LSHIFT -> State 95
      BIT_RSHIFT -> State 96
      mul_op -> State 100

  State 163:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 86
      NOT_EQUAL -> State 91
      LESS -> State 89
      LESS_OR_EQUAL -> State 90
      GREATER -> State 87
      GREATER_OR_EQUAL -> State 88
      cmp_op -> State 92

  State 164:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 80
      SUB -> State 83
      BIT_XOR -> State 82
      BIT_OR -> State 81
      add_op -> State 84

  State 165:
    Kernel Items:
      arguments: arguments.COMMA argument
      anonymous_struct_expr: explicit_struct_def LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 219
      COMMA -> State 127

  State 166:
    Kernel Items:
      mul_expr: mul_expr mul_op prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 167:
    Kernel Items:
      loop_expr: DO block_body., *
      loop_expr: DO block_body.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 220

  State 168:
    Kernel Items:
      loop_expr: FOR IN.sequence_expr DO block_body
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      optional_label_decl -> State 111
      sequence_expr -> State 221
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 169:
    Kernel Items:
      loop_expr: FOR SEMICOLON.optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      LBRACE -> [optional_label_decl]
      SEMICOLON -> [optional_sequence_expr]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      optional_label_decl -> State 111
      optional_sequence_expr -> State 222
      sequence_expr -> State 223
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 170:
    Kernel Items:
      loop_expr: FOR sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 224

  State 171:
    Kernel Items:
      if_expr: IF sequence_expr.block_body
      if_expr: IF sequence_expr.block_body ELSE block_body
      if_expr: IF sequence_expr.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 104
      block_body -> State 225

  State 172:
    Kernel Items:
      block_body: LBRACE statements.RBRACE
      statements: statements.statement
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      RETURN -> State 231
      BREAK -> State 227
      CONTINUE -> State 228
      UNSAFE -> State 56
      STRUCT -> State 33
      FUNC -> State 23
      ASYNC -> State 226
      DEFER -> State 229
      LABEL_DECL -> State 26
      RBRACE -> State 230
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      expression -> State 233
      optional_label_decl -> State 48
      sequence_expr -> State 53
      block_expr -> State 42
      statement -> State 237
      statement_body -> State 238
      unsafe_statement -> State 239
      jump_statement -> State 235
      jump_type -> State 236
      expressions -> State 234
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 232
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 173:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.LBRACE CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 240

  State 174:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 85

  State 175:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 241

  State 176:
    Kernel Items:
      definitions: definitions NEWLINES definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 177:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements.RPAREN
      package_statements: package_statements.package_statement
    Reduce:
      (nil)
    Goto:
      UNSAFE -> State 56
      RPAREN -> State 242
      unsafe_statement -> State 245
      package_statement_body -> State 244
      package_statement -> State 243

  State 178:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 246
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 179:
    Kernel Items:
      generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 180:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs.COMMA generic_parameter_def
      optional_generic_parameter_defs: generic_parameter_defs., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      COMMA -> State 247

  State 181:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 248

  State 182:
    Kernel Items:
      type_def: TYPE IDENTIFIER EQUAL value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      (nil)

  State 183:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., *
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 249

  State 184:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 250
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

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
      LPAREN -> State 251

  State 188:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 253
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      return_type -> State 252
      func_type -> State 146

  State 189:
    Kernel Items:
      parameter_defs: parameter_defs COMMA.parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 120
      parameter_def -> State 254

  State 190:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 191:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 192:
    Kernel Items:
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      (nil)

  State 193:
    Kernel Items:
      colon_expressions: colon_expressions COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 194:
    Kernel Items:
      colon_expressions: optional_expression COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 195:
    Kernel Items:
      trait_algebra_type: trait_algebra_type.MUL traitable_type
      trait_algebra_type: trait_algebra_type.ADD traitable_type
      trait_algebra_type: trait_algebra_type.SUB traitable_type
      value_type: BIT_AND trait_algebra_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      ADD -> State 211
      SUB -> State 213
      MUL -> State 212

  State 196:
    Kernel Items:
      traitable_type: BIT_NEG atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 197:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 135
      UNSAFE -> State 56
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      unsafe_statement -> State 153
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 154
      field_def -> State 257
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      enum_value_def -> State 255
      implicit_enum_value_defs -> State 258
      implicit_enum_def -> State 147
      explicit_enum_value_defs -> State 256
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 198:
    Kernel Items:
      func_type: FUNC LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 260
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      DOTDOTDOT -> State 259
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 264
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      parameter_decl -> State 262
      parameter_decls -> State 263
      optional_parameter_decls -> State 261
      func_type -> State 146

  State 199:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 200:
    Kernel Items:
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 201:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 265

  State 202:
    Kernel Items:
      implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 266

  State 203:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 268
      OR -> State 267

  State 204:
    Kernel Items:
      implicit_field_defs: implicit_field_defs.COMMA field_def
      optional_implicit_field_defs: implicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      COMMA -> State 269

  State 205:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 270

  State 206:
    Kernel Items:
      traitable_type: TILDE_TILDE atom_type., *
    Reduce:
      * -> [traitable_type]
    Goto:
      (nil)

  State 207:
    Kernel Items:
      trait_def: TRAIT LPAREN.optional_trait_properties RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      IDENTIFIER -> State 135
      UNSAFE -> State 56
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 271
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      unsafe_statement -> State 153
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 154
      field_def -> State 272
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_property -> State 276
      trait_properties -> State 275
      optional_trait_properties -> State 274
      trait_def -> State 151
      func_type -> State 146
      method_signature -> State 273

  State 208:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 135
      UNSAFE -> State 56
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      unsafe_statement -> State 153
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 154
      field_def -> State 277
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 209:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 135
      UNSAFE -> State 56
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      unsafe_statement -> State 153
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 154
      field_def -> State 278
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 210:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_def]
    Goto:
      (nil)

  State 211:
    Kernel Items:
      trait_algebra_type: trait_algebra_type ADD.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      LPAREN -> State 137
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 279
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151

  State 212:
    Kernel Items:
      trait_algebra_type: trait_algebra_type MUL.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      LPAREN -> State 137
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 280
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151

  State 213:
    Kernel Items:
      trait_algebra_type: trait_algebra_type SUB.traitable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      LPAREN -> State 137
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 281
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151

  State 214:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 282
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 215:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 216:
    Kernel Items:
      access_expr: access_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      optional_arguments: arguments., RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 127

  State 218:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 283

  State 219:
    Kernel Items:
      anonymous_struct_expr: explicit_struct_def LPAREN arguments RPAREN., *
    Reduce:
      * -> [anonymous_struct_expr]
    Goto:
      (nil)

  State 220:
    Kernel Items:
      loop_expr: DO block_body FOR.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      optional_label_decl -> State 111
      sequence_expr -> State 284
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 221:
    Kernel Items:
      loop_expr: FOR IN sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 285

  State 222:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 286

  State 223:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 224:
    Kernel Items:
      loop_expr: FOR sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 104
      block_body -> State 287

  State 225:
    Kernel Items:
      if_expr: IF sequence_expr block_body., *
      if_expr: IF sequence_expr block_body.ELSE block_body
      if_expr: IF sequence_expr block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 288

  State 226:
    Kernel Items:
      statement_body: ASYNC.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      LEX_ERROR -> State 27
      optional_label_decl -> State 111
      block_expr -> State 42
      call_expr -> State 290
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 289
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 227:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 228:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 229:
    Kernel Items:
      statement_body: DEFER.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      LEX_ERROR -> State 27
      optional_label_decl -> State 111
      block_expr -> State 42
      call_expr -> State 291
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 289
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 230:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 231:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 232:
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
      LBRACKET -> State 77
      DOT -> State 76
      QUESTION -> State 78
      DOLLAR_LBRACKET -> State 75
      ADD_ASSIGN -> State 292
      SUB_ASSIGN -> State 303
      MUL_ASSIGN -> State 302
      DIV_ASSIGN -> State 300
      MOD_ASSIGN -> State 301
      ADD_ONE_ASSIGN -> State 293
      SUB_ONE_ASSIGN -> State 304
      BIT_NEG_ASSIGN -> State 296
      BIT_AND_ASSIGN -> State 294
      BIT_OR_ASSIGN -> State 297
      BIT_XOR_ASSIGN -> State 299
      BIT_LSHIFT_ASSIGN -> State 295
      BIT_RSHIFT_ASSIGN -> State 298
      unary_op_assign -> State 306
      binary_op_assign -> State 305
      optional_generic_binding -> State 79

  State 233:
    Kernel Items:
      expressions: expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 234:
    Kernel Items:
      statement_body: expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [statement_body]
    Goto:
      COMMA -> State 307

  State 235:
    Kernel Items:
      statement_body: jump_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 236:
    Kernel Items:
      jump_statement: jump_type.optional_jump_label optional_expressions
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 308
      optional_jump_label -> State 309

  State 237:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 238:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 310
      SEMICOLON -> State 311

  State 239:
    Kernel Items:
      statement_body: unsafe_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE.CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      CASE -> State 312

  State 241:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 313

  State 242:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements RPAREN., $
    Reduce:
      $ -> [package_def]
    Goto:
      (nil)

  State 243:
    Kernel Items:
      package_statements: package_statements package_statement., *
    Reduce:
      * -> [package_statements]
    Goto:
      (nil)

  State 244:
    Kernel Items:
      package_statement: package_statement_body.NEWLINES
      package_statement: package_statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 314
      SEMICOLON -> State 315

  State 245:
    Kernel Items:
      package_statement_body: unsafe_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 246:
    Kernel Items:
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      (nil)

  State 247:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA.generic_parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 178
      generic_parameter_def -> State 316

  State 248:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 249:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 317
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 250:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      (nil)

  State 251:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 120
      parameter_def -> State 124
      parameter_defs -> State 125
      optional_parameter_defs -> State 318

  State 252:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 104
      block_body -> State 319

  State 253:
    Kernel Items:
      return_type: value_type., $
    Reduce:
      $ -> [return_type]
    Goto:
      (nil)

  State 254:
    Kernel Items:
      parameter_defs: parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 255:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 320
      OR -> State 321

  State 256:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 322

  State 257:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 266

  State 258:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 323
      OR -> State 324

  State 259:
    Kernel Items:
      parameter_decl: DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 325
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 260:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      DOLLAR_LBRACKET -> State 75
      DOTDOTDOT -> State 326
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      optional_generic_binding -> State 199
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 327
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 261:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 328

  State 262:
    Kernel Items:
      parameter_decls: parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 263:
    Kernel Items:
      parameter_decls: parameter_decls.COMMA parameter_decl
      optional_parameter_decls: parameter_decls., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      COMMA -> State 329

  State 264:
    Kernel Items:
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 265:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 135
      UNSAFE -> State 56
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      unsafe_statement -> State 153
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 154
      field_def -> State 257
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      enum_value_def -> State 330
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 266:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 331

  State 267:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 135
      UNSAFE -> State 56
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      unsafe_statement -> State 153
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 154
      field_def -> State 257
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      enum_value_def -> State 332
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 268:
    Kernel Items:
      implicit_enum_def: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_def]
    Goto:
      (nil)

  State 269:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 135
      UNSAFE -> State 56
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      unsafe_statement -> State 153
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 154
      field_def -> State 333
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 270:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_def]
    Goto:
      (nil)

  State 271:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 334
      LPAREN -> State 198

  State 272:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 273:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 274:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 335

  State 275:
    Kernel Items:
      trait_properties: trait_properties.NEWLINES trait_property
      trait_properties: trait_properties.COMMA trait_property
      optional_trait_properties: trait_properties., RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      NEWLINES -> State 337
      COMMA -> State 336

  State 276:
    Kernel Items:
      trait_properties: trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 277:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      trait_algebra_type: trait_algebra_type ADD traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 280:
    Kernel Items:
      trait_algebra_type: trait_algebra_type MUL traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 281:
    Kernel Items:
      trait_algebra_type: trait_algebra_type SUB traitable_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 282:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
    Reduce:
      * -> [generic_arguments]
    Goto:
      (nil)

  State 283:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 284:
    Kernel Items:
      loop_expr: DO block_body FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 285:
    Kernel Items:
      loop_expr: FOR IN sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 104
      block_body -> State 338

  State 286:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON.optional_sequence_expr DO block_body
    Reduce:
      DO -> [optional_sequence_expr]
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      optional_label_decl -> State 111
      optional_sequence_expr -> State 339
      sequence_expr -> State 223
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 287:
    Kernel Items:
      loop_expr: FOR sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 288:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE.block_body
      if_expr: IF sequence_expr block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 103
      LBRACE -> State 104
      if_expr -> State 341
      block_body -> State 340

  State 289:
    Kernel Items:
      call_expr: access_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET argument RBRACKET
    Reduce:
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 77
      DOT -> State 76
      DOLLAR_LBRACKET -> State 75
      optional_generic_binding -> State 79

  State 290:
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

  State 291:
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

  State 292:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 293:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 294:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 295:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 296:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 297:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 298:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 299:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 300:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 301:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 302:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 303:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 304:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 305:
    Kernel Items:
      statement_body: access_expr binary_op_assign.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      expression -> State 342
      optional_label_decl -> State 48
      sequence_expr -> State 53
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 306:
    Kernel Items:
      statement_body: access_expr unary_op_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 307:
    Kernel Items:
      expressions: expressions COMMA.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      expression -> State 343
      optional_label_decl -> State 48
      sequence_expr -> State 53
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 308:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 309:
    Kernel Items:
      jump_statement: jump_type optional_jump_label.optional_expressions
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_expressions]
      SEMICOLON -> [optional_expressions]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 31
      STRING_LITERAL -> State 32
      IDENTIFIER -> State 24
      TRUE -> State 35
      FALSE -> State 21
      STRUCT -> State 33
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 28
      NOT -> State 30
      SUB -> State 34
      MUL -> State 29
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 27
      expression -> State 233
      optional_label_decl -> State 48
      sequence_expr -> State 53
      block_expr -> State 42
      expressions -> State 344
      optional_expressions -> State 345
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 46
      anonymous_struct_expr -> State 40
      access_expr -> State 36
      postfix_unary_expr -> State 50
      prefix_unary_op -> State 52
      prefix_unary_expr -> State 51
      mul_expr -> State 47
      add_expr -> State 37
      cmp_expr -> State 44
      and_expr -> State 38
      or_expr -> State 49
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 39

  State 310:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 311:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 312:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE CASE.DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 346

  State 313:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 314:
    Kernel Items:
      package_statement: package_statement_body NEWLINES., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 315:
    Kernel Items:
      package_statement: package_statement_body SEMICOLON., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 316:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 317:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      (nil)

  State 318:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 347

  State 319:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 320:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 135
      UNSAFE -> State 56
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      unsafe_statement -> State 153
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 154
      field_def -> State 257
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      enum_value_def -> State 348
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 321:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 135
      UNSAFE -> State 56
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      unsafe_statement -> State 153
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 154
      field_def -> State 257
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      enum_value_def -> State 349
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 322:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_def]
    Goto:
      (nil)

  State 323:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 135
      UNSAFE -> State 56
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      unsafe_statement -> State 153
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 154
      field_def -> State 257
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      enum_value_def -> State 350
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 324:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 135
      UNSAFE -> State 56
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      unsafe_statement -> State 153
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 154
      field_def -> State 257
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      enum_value_def -> State 351
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 325:
    Kernel Items:
      parameter_decl: DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 326:
    Kernel Items:
      parameter_decl: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 352
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      func_type -> State 146

  State 327:
    Kernel Items:
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 328:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 253
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      return_type -> State 353
      func_type -> State 146

  State 329:
    Kernel Items:
      parameter_decls: parameter_decls COMMA.parameter_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 260
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      DOTDOTDOT -> State 259
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 264
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      parameter_decl -> State 354
      func_type -> State 146

  State 330:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 331:
    Kernel Items:
      enum_value_def: field_def ASSIGN DEFAULT., *
    Reduce:
      * -> [enum_value_def]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 333:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [implicit_field_defs]
    Goto:
      (nil)

  State 334:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 355

  State 335:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties RPAREN., *
    Reduce:
      * -> [trait_def]
    Goto:
      (nil)

  State 336:
    Kernel Items:
      trait_properties: trait_properties COMMA.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 135
      UNSAFE -> State 56
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 271
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      unsafe_statement -> State 153
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 154
      field_def -> State 272
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_property -> State 356
      trait_def -> State 151
      func_type -> State 146
      method_signature -> State 273

  State 337:
    Kernel Items:
      trait_properties: trait_properties NEWLINES.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 135
      UNSAFE -> State 56
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 271
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      unsafe_statement -> State 153
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 154
      field_def -> State 272
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_property -> State 357
      trait_def -> State 151
      func_type -> State 146
      method_signature -> State 273

  State 338:
    Kernel Items:
      loop_expr: FOR IN sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 339:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 358

  State 340:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 342:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 343:
    Kernel Items:
      expressions: expressions COMMA expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 344:
    Kernel Items:
      expressions: expressions.COMMA expression
      optional_expressions: expressions., *
    Reduce:
      * -> [optional_expressions]
    Goto:
      COMMA -> State 307

  State 345:
    Kernel Items:
      jump_statement: jump_type optional_jump_label optional_expressions., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 346:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE CASE DEFAULT.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 359

  State 347:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 253
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      return_type -> State 360
      func_type -> State 146

  State 348:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 349:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
      explicit_enum_value_defs: enum_value_def OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 350:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 351:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 352:
    Kernel Items:
      parameter_decl: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 353:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN return_type., $
    Reduce:
      $ -> [func_type]
    Goto:
      (nil)

  State 354:
    Kernel Items:
      parameter_decls: parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 355:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 260
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      DOTDOTDOT -> State 259
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 264
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      parameter_decl -> State 262
      parameter_decls -> State 263
      optional_parameter_decls -> State 361
      func_type -> State 146

  State 356:
    Kernel Items:
      trait_properties: trait_properties COMMA trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 357:
    Kernel Items:
      trait_properties: trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 358:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 104
      block_body -> State 362

  State 359:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE CASE DEFAULT RBRACE., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 360:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 104
      block_body -> State 363

  State 361:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 364

  State 362:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 363:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 364:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 155
      STRUCT -> State 33
      ENUM -> State 133
      TRAIT -> State 140
      FUNC -> State 134
      LPAREN -> State 137
      QUESTION -> State 138
      TILDE_TILDE -> State 139
      BIT_NEG -> State 132
      BIT_AND -> State 131
      LEX_ERROR -> State 136
      atom_type -> State 141
      traitable_type -> State 152
      trait_algebra_type -> State 150
      value_type -> State 253
      implicit_struct_def -> State 148
      explicit_struct_def -> State 144
      implicit_enum_def -> State 147
      explicit_enum_def -> State 142
      trait_def -> State 151
      return_type -> State 365
      func_type -> State 146

  State 365:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

Number of states: 365
Number of shift actions: 2120
Number of reduce actions: 300
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
