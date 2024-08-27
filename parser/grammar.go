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
	ExclaimToken         = SymbolId(300)
	DollarLbracketToken  = SymbolId(301)
	DotdotdotToken       = SymbolId(302)
	TildeTildeToken      = SymbolId(303)
	AssignToken          = SymbolId(304)
	AddAssignToken       = SymbolId(305)
	SubAssignToken       = SymbolId(306)
	MulAssignToken       = SymbolId(307)
	DivAssignToken       = SymbolId(308)
	ModAssignToken       = SymbolId(309)
	AddOneAssignToken    = SymbolId(310)
	SubOneAssignToken    = SymbolId(311)
	BitNegAssignToken    = SymbolId(312)
	BitAndAssignToken    = SymbolId(313)
	BitOrAssignToken     = SymbolId(314)
	BitXorAssignToken    = SymbolId(315)
	BitLshiftAssignToken = SymbolId(316)
	BitRshiftAssignToken = SymbolId(317)
	NotToken             = SymbolId(318)
	AndToken             = SymbolId(319)
	OrToken              = SymbolId(320)
	AddToken             = SymbolId(321)
	SubToken             = SymbolId(322)
	MulToken             = SymbolId(323)
	DivToken             = SymbolId(324)
	ModToken             = SymbolId(325)
	BitNegToken          = SymbolId(326)
	BitAndToken          = SymbolId(327)
	BitXorToken          = SymbolId(328)
	BitOrToken           = SymbolId(329)
	BitLshiftToken       = SymbolId(330)
	BitRshiftToken       = SymbolId(331)
	EqualToken           = SymbolId(332)
	NotEqualToken        = SymbolId(333)
	LessToken            = SymbolId(334)
	LessOrEqualToken     = SymbolId(335)
	GreaterToken         = SymbolId(336)
	GreaterOrEqualToken  = SymbolId(337)
	LexErrorToken        = SymbolId(338)
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

	// 281:2: atom_expr -> initialize_expr: ...
	InitializeExprToAtomExpr(InitializeExpr_ *GenericSymbol) (*GenericSymbol, error)

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

	// 293:2: initialize_expr -> explicit: ...
	ExplicitToInitializeExpr(InitializableType_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 294:2: initialize_expr -> implicit_struct: ...
	ImplicitStructToInitializeExpr(Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

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

	// 369:2: initializable_type -> explicit_struct_def: ...
	ExplicitStructDefToInitializableType(ExplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 370:2: initializable_type -> explicit_enum_def: ...
	ExplicitEnumDefToInitializableType(ExplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 371:2: initializable_type -> slice: ...
	SliceToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 372:2: initializable_type -> array: ...
	ArrayToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Comma_ *GenericSymbol, IntegerLiteral_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 373:2: initializable_type -> map: ...
	MapToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Colon_ *GenericSymbol, ValueType_2 *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 376:2: atom_type -> initializable_type: ...
	InitializableTypeToAtomType(InitializableType_ *GenericSymbol) (*GenericSymbol, error)

	// 377:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 378:2: atom_type -> extern_named: ...
	ExternNamedToAtomType(Identifier_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_2 *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 379:2: atom_type -> inferred: ...
	InferredToAtomType(Dot_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 380:2: atom_type -> implicit_struct_def: ...
	ImplicitStructDefToAtomType(ImplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 381:2: atom_type -> implicit_enum_def: ...
	ImplicitEnumDefToAtomType(ImplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 382:2: atom_type -> trait_def: ...
	TraitDefToAtomType(TraitDef_ *GenericSymbol) (*GenericSymbol, error)

	// 383:2: atom_type -> func_type: ...
	FuncTypeToAtomType(FuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 384:2: atom_type -> LEX_ERROR: ...
	LexErrorToAtomType(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 387:2: value_type -> atom_type: ...
	AtomTypeToValueType(AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 388:2: value_type -> optional: ...
	OptionalToValueType(Question_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 389:2: value_type -> result: ...
	ResultToValueType(Exclaim_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 390:2: value_type -> reference: ...
	ReferenceToValueType(BitAnd_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 391:2: value_type -> public_methods_trait: ...
	PublicMethodsTraitToValueType(BitNeg_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 392:2: value_type -> public_trait: ...
	PublicTraitToValueType(TildeTilde_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 395:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 396:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol, Implements_ *GenericSymbol, ValueType_2 *GenericSymbol) (*GenericSymbol, error)

	// 397:2: type_def -> alias: ...
	AliasToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, Equal_ *GenericSymbol, TraitAlgebraType_ *GenericSymbol) (*GenericSymbol, error)

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

	// 482:2: trait_algebra_type -> value_type: ...
	ValueTypeToTraitAlgebraType(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 483:2: trait_algebra_type -> intersect: ...
	IntersectToTraitAlgebraType(TraitAlgebraType_ *GenericSymbol, Mul_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 484:2: trait_algebra_type -> union: ...
	UnionToTraitAlgebraType(TraitAlgebraType_ *GenericSymbol, Add_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 485:2: trait_algebra_type -> difference: ...
	DifferenceToTraitAlgebraType(TraitAlgebraType_ *GenericSymbol, Sub_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 488:2: trait_property -> explicit: ...
	ExplicitToTraitProperty(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 489:2: trait_property -> implicit: ...
	ImplicitToTraitProperty(TraitAlgebraType_ *GenericSymbol) (*GenericSymbol, error)

	// 490:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ *GenericSymbol) (*GenericSymbol, error)

	// 491:2: trait_property -> unsafe_statement: ...
	UnsafeStatementToTraitProperty(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 494:2: trait_properties -> trait_property: ...
	TraitPropertyToTraitProperties(TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 495:2: trait_properties -> implicit: ...
	ImplicitToTraitProperties(TraitProperties_ *GenericSymbol, Newlines_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 496:2: trait_properties -> explicit: ...
	ExplicitToTraitProperties(TraitProperties_ *GenericSymbol, Comma_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 499:2: optional_trait_properties -> trait_properties: ...
	TraitPropertiesToOptionalTraitProperties(TraitProperties_ *GenericSymbol) (*GenericSymbol, error)

	// 500:2: optional_trait_properties -> nil: ...
	NilToOptionalTraitProperties() (*GenericSymbol, error)

	// 502:13: trait_def -> ...
	ToTraitDef(Trait_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalTraitProperties_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 510:2: return_type -> value_type: ...
	ValueTypeToReturnType(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 511:2: return_type -> nil: ...
	NilToReturnType() (*GenericSymbol, error)

	// 514:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 515:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 516:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 517:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 520:2: parameter_decls -> parameter_decl: ...
	ParameterDeclToParameterDecls(ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 521:2: parameter_decls -> add: ...
	AddToParameterDecls(ParameterDecls_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 524:2: optional_parameter_decls -> parameter_decls: ...
	ParameterDeclsToOptionalParameterDecls(ParameterDecls_ *GenericSymbol) (*GenericSymbol, error)

	// 525:2: optional_parameter_decls -> nil: ...
	NilToOptionalParameterDecls() (*GenericSymbol, error)

	// 527:13: func_type -> ...
	ToFuncType(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 535:20: method_signature -> ...
	ToMethodSignature(Func_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 538:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 539:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 542:2: parameter_defs -> parameter_def: ...
	ParameterDefToParameterDefs(ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 543:2: parameter_defs -> add: ...
	AddToParameterDefs(ParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 546:2: optional_parameter_defs -> parameter_defs: ...
	ParameterDefsToOptionalParameterDefs(ParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 547:2: optional_parameter_defs -> nil: ...
	NilToOptionalParameterDefs() (*GenericSymbol, error)

	// 550:2: optional_receiver -> receiver: ...
	ReceiverToOptionalReceiver(Lparen_ *GenericSymbol, ParameterDef_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 551:2: optional_receiver -> nil: ...
	NilToOptionalReceiver() (*GenericSymbol, error)

	// 554:2: named_func_def -> ...
	ToNamedFuncDef(Func_ *GenericSymbol, OptionalReceiver_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 557:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 564:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 565:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, PackageStatements_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 567:26: package_statement_body -> ...
	ToPackageStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 570:2: package_statement -> implicit: ...
	ImplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 571:2: package_statement -> explicit: ...
	ExplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 574:2: package_statements -> empty_list: ...
	EmptyListToPackageStatements() (*GenericSymbol, error)

	// 575:2: package_statements -> add: ...
	AddToPackageStatements(PackageStatements_ *GenericSymbol, PackageStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 579:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 580:2: lex_internal_tokens -> COMMENT: ...
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
	case ExclaimToken:
		return "EXCLAIM"
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
	case InitializeExprType:
		return "initialize_expr"
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
	case InitializableTypeType:
		return "initializable_type"
	case AtomTypeType:
		return "atom_type"
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
	case TraitAlgebraTypeType:
		return "trait_algebra_type"
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

	SourceType                       = SymbolId(339)
	OptionalDefinitionsType          = SymbolId(340)
	OptionalNewlinesType             = SymbolId(341)
	DefinitionsType                  = SymbolId(342)
	DefinitionType                   = SymbolId(343)
	ExpressionType                   = SymbolId(344)
	OptionalLabelDeclType            = SymbolId(345)
	IfExprType                       = SymbolId(346)
	SwitchExprType                   = SymbolId(347)
	LoopExprType                     = SymbolId(348)
	OptionalSequenceExprType         = SymbolId(349)
	SequenceExprType                 = SymbolId(350)
	BlockExprType                    = SymbolId(351)
	BlockBodyType                    = SymbolId(352)
	StatementsType                   = SymbolId(353)
	StatementType                    = SymbolId(354)
	StatementBodyType                = SymbolId(355)
	UnaryOpAssignType                = SymbolId(356)
	BinaryOpAssignType               = SymbolId(357)
	UnsafeStatementType              = SymbolId(358)
	JumpStatementType                = SymbolId(359)
	JumpTypeType                     = SymbolId(360)
	OptionalJumpLabelType            = SymbolId(361)
	ExpressionsType                  = SymbolId(362)
	OptionalExpressionsType          = SymbolId(363)
	CallExprType                     = SymbolId(364)
	OptionalGenericBindingType       = SymbolId(365)
	OptionalGenericArgumentsType     = SymbolId(366)
	GenericArgumentsType             = SymbolId(367)
	OptionalArgumentsType            = SymbolId(368)
	ArgumentsType                    = SymbolId(369)
	ArgumentType                     = SymbolId(370)
	ColonExpressionsType             = SymbolId(371)
	OptionalExpressionType           = SymbolId(372)
	AtomExprType                     = SymbolId(373)
	LiteralType                      = SymbolId(374)
	InitializeExprType               = SymbolId(375)
	AccessExprType                   = SymbolId(376)
	PostfixUnaryExprType             = SymbolId(377)
	PrefixUnaryOpType                = SymbolId(378)
	PrefixUnaryExprType              = SymbolId(379)
	MulOpType                        = SymbolId(380)
	MulExprType                      = SymbolId(381)
	AddOpType                        = SymbolId(382)
	AddExprType                      = SymbolId(383)
	CmpOpType                        = SymbolId(384)
	CmpExprType                      = SymbolId(385)
	AndExprType                      = SymbolId(386)
	OrExprType                       = SymbolId(387)
	InitializableTypeType            = SymbolId(388)
	AtomTypeType                     = SymbolId(389)
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
	TraitAlgebraTypeType             = SymbolId(408)
	TraitPropertyType                = SymbolId(409)
	TraitPropertiesType              = SymbolId(410)
	OptionalTraitPropertiesType      = SymbolId(411)
	TraitDefType                     = SymbolId(412)
	ReturnTypeType                   = SymbolId(413)
	ParameterDeclType                = SymbolId(414)
	ParameterDeclsType               = SymbolId(415)
	OptionalParameterDeclsType       = SymbolId(416)
	FuncTypeType                     = SymbolId(417)
	MethodSignatureType              = SymbolId(418)
	ParameterDefType                 = SymbolId(419)
	ParameterDefsType                = SymbolId(420)
	OptionalParameterDefsType        = SymbolId(421)
	OptionalReceiverType             = SymbolId(422)
	NamedFuncDefType                 = SymbolId(423)
	AnonymousFuncExprType            = SymbolId(424)
	PackageDefType                   = SymbolId(425)
	PackageStatementBodyType         = SymbolId(426)
	PackageStatementType             = SymbolId(427)
	PackageStatementsType            = SymbolId(428)
	LexInternalTokensType            = SymbolId(429)
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
	_ReduceInitializeExprToAtomExpr                           = _ReduceType(90)
	_ReduceLexErrorToAtomExpr                                 = _ReduceType(91)
	_ReduceTrueToLiteral                                      = _ReduceType(92)
	_ReduceFalseToLiteral                                     = _ReduceType(93)
	_ReduceIntegerLiteralToLiteral                            = _ReduceType(94)
	_ReduceFloatLiteralToLiteral                              = _ReduceType(95)
	_ReduceRuneLiteralToLiteral                               = _ReduceType(96)
	_ReduceStringLiteralToLiteral                             = _ReduceType(97)
	_ReduceExplicitToInitializeExpr                           = _ReduceType(98)
	_ReduceImplicitStructToInitializeExpr                     = _ReduceType(99)
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
	_ReduceExplicitStructDefToInitializableType               = _ReduceType(139)
	_ReduceExplicitEnumDefToInitializableType                 = _ReduceType(140)
	_ReduceSliceToInitializableType                           = _ReduceType(141)
	_ReduceArrayToInitializableType                           = _ReduceType(142)
	_ReduceMapToInitializableType                             = _ReduceType(143)
	_ReduceInitializableTypeToAtomType                        = _ReduceType(144)
	_ReduceNamedToAtomType                                    = _ReduceType(145)
	_ReduceExternNamedToAtomType                              = _ReduceType(146)
	_ReduceInferredToAtomType                                 = _ReduceType(147)
	_ReduceImplicitStructDefToAtomType                        = _ReduceType(148)
	_ReduceImplicitEnumDefToAtomType                          = _ReduceType(149)
	_ReduceTraitDefToAtomType                                 = _ReduceType(150)
	_ReduceFuncTypeToAtomType                                 = _ReduceType(151)
	_ReduceLexErrorToAtomType                                 = _ReduceType(152)
	_ReduceAtomTypeToValueType                                = _ReduceType(153)
	_ReduceOptionalToValueType                                = _ReduceType(154)
	_ReduceResultToValueType                                  = _ReduceType(155)
	_ReduceReferenceToValueType                               = _ReduceType(156)
	_ReducePublicMethodsTraitToValueType                      = _ReduceType(157)
	_ReducePublicTraitToValueType                             = _ReduceType(158)
	_ReduceDefinitionToTypeDef                                = _ReduceType(159)
	_ReduceConstrainedDefToTypeDef                            = _ReduceType(160)
	_ReduceAliasToTypeDef                                     = _ReduceType(161)
	_ReduceUnconstrainedToGenericParameterDef                 = _ReduceType(162)
	_ReduceConstrainedToGenericParameterDef                   = _ReduceType(163)
	_ReduceGenericParameterDefToGenericParameterDefs          = _ReduceType(164)
	_ReduceAddToGenericParameterDefs                          = _ReduceType(165)
	_ReduceGenericParameterDefsToOptionalGenericParameterDefs = _ReduceType(166)
	_ReduceNilToOptionalGenericParameterDefs                  = _ReduceType(167)
	_ReduceGenericToOptionalGenericParameters                 = _ReduceType(168)
	_ReduceNilToOptionalGenericParameters                     = _ReduceType(169)
	_ReduceExplicitToFieldDef                                 = _ReduceType(170)
	_ReduceImplicitToFieldDef                                 = _ReduceType(171)
	_ReduceUnsafeStatementToFieldDef                          = _ReduceType(172)
	_ReduceFieldDefToImplicitFieldDefs                        = _ReduceType(173)
	_ReduceAddToImplicitFieldDefs                             = _ReduceType(174)
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefs       = _ReduceType(175)
	_ReduceNilToOptionalImplicitFieldDefs                     = _ReduceType(176)
	_ReduceToImplicitStructDef                                = _ReduceType(177)
	_ReduceFieldDefToExplicitFieldDefs                        = _ReduceType(178)
	_ReduceImplicitToExplicitFieldDefs                        = _ReduceType(179)
	_ReduceExplicitToExplicitFieldDefs                        = _ReduceType(180)
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefs       = _ReduceType(181)
	_ReduceNilToOptionalExplicitFieldDefs                     = _ReduceType(182)
	_ReduceToExplicitStructDef                                = _ReduceType(183)
	_ReduceFieldDefToEnumValueDef                             = _ReduceType(184)
	_ReduceDefaultToEnumValueDef                              = _ReduceType(185)
	_ReducePairToImplicitEnumValueDefs                        = _ReduceType(186)
	_ReduceAddToImplicitEnumValueDefs                         = _ReduceType(187)
	_ReduceToImplicitEnumDef                                  = _ReduceType(188)
	_ReduceExplicitPairToExplicitEnumValueDefs                = _ReduceType(189)
	_ReduceImplicitPairToExplicitEnumValueDefs                = _ReduceType(190)
	_ReduceExplicitAddToExplicitEnumValueDefs                 = _ReduceType(191)
	_ReduceImplicitAddToExplicitEnumValueDefs                 = _ReduceType(192)
	_ReduceToExplicitEnumDef                                  = _ReduceType(193)
	_ReduceValueTypeToTraitAlgebraType                        = _ReduceType(194)
	_ReduceIntersectToTraitAlgebraType                        = _ReduceType(195)
	_ReduceUnionToTraitAlgebraType                            = _ReduceType(196)
	_ReduceDifferenceToTraitAlgebraType                       = _ReduceType(197)
	_ReduceExplicitToTraitProperty                            = _ReduceType(198)
	_ReduceImplicitToTraitProperty                            = _ReduceType(199)
	_ReduceMethodSignatureToTraitProperty                     = _ReduceType(200)
	_ReduceUnsafeStatementToTraitProperty                     = _ReduceType(201)
	_ReduceTraitPropertyToTraitProperties                     = _ReduceType(202)
	_ReduceImplicitToTraitProperties                          = _ReduceType(203)
	_ReduceExplicitToTraitProperties                          = _ReduceType(204)
	_ReduceTraitPropertiesToOptionalTraitProperties           = _ReduceType(205)
	_ReduceNilToOptionalTraitProperties                       = _ReduceType(206)
	_ReduceToTraitDef                                         = _ReduceType(207)
	_ReduceValueTypeToReturnType                              = _ReduceType(208)
	_ReduceNilToReturnType                                    = _ReduceType(209)
	_ReduceArgToParameterDecl                                 = _ReduceType(210)
	_ReduceVarargToParameterDecl                              = _ReduceType(211)
	_ReduceUnamedToParameterDecl                              = _ReduceType(212)
	_ReduceUnnamedVarargToParameterDecl                       = _ReduceType(213)
	_ReduceParameterDeclToParameterDecls                      = _ReduceType(214)
	_ReduceAddToParameterDecls                                = _ReduceType(215)
	_ReduceParameterDeclsToOptionalParameterDecls             = _ReduceType(216)
	_ReduceNilToOptionalParameterDecls                        = _ReduceType(217)
	_ReduceToFuncType                                         = _ReduceType(218)
	_ReduceToMethodSignature                                  = _ReduceType(219)
	_ReduceArgToParameterDef                                  = _ReduceType(220)
	_ReduceVarargToParameterDef                               = _ReduceType(221)
	_ReduceParameterDefToParameterDefs                        = _ReduceType(222)
	_ReduceAddToParameterDefs                                 = _ReduceType(223)
	_ReduceParameterDefsToOptionalParameterDefs               = _ReduceType(224)
	_ReduceNilToOptionalParameterDefs                         = _ReduceType(225)
	_ReduceReceiverToOptionalReceiver                         = _ReduceType(226)
	_ReduceNilToOptionalReceiver                              = _ReduceType(227)
	_ReduceToNamedFuncDef                                     = _ReduceType(228)
	_ReduceToAnonymousFuncExpr                                = _ReduceType(229)
	_ReduceNoSpecToPackageDef                                 = _ReduceType(230)
	_ReduceWithSpecToPackageDef                               = _ReduceType(231)
	_ReduceToPackageStatementBody                             = _ReduceType(232)
	_ReduceImplicitToPackageStatement                         = _ReduceType(233)
	_ReduceExplicitToPackageStatement                         = _ReduceType(234)
	_ReduceEmptyListToPackageStatements                       = _ReduceType(235)
	_ReduceAddToPackageStatements                             = _ReduceType(236)
	_ReduceSpacesToLexInternalTokens                          = _ReduceType(237)
	_ReduceCommentToLexInternalTokens                         = _ReduceType(238)
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
	case _ReduceInitializeExprToAtomExpr:
		return "InitializeExprToAtomExpr"
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
	case _ReduceExplicitToInitializeExpr:
		return "ExplicitToInitializeExpr"
	case _ReduceImplicitStructToInitializeExpr:
		return "ImplicitStructToInitializeExpr"
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
	case _ReduceExplicitStructDefToInitializableType:
		return "ExplicitStructDefToInitializableType"
	case _ReduceExplicitEnumDefToInitializableType:
		return "ExplicitEnumDefToInitializableType"
	case _ReduceSliceToInitializableType:
		return "SliceToInitializableType"
	case _ReduceArrayToInitializableType:
		return "ArrayToInitializableType"
	case _ReduceMapToInitializableType:
		return "MapToInitializableType"
	case _ReduceInitializableTypeToAtomType:
		return "InitializableTypeToAtomType"
	case _ReduceNamedToAtomType:
		return "NamedToAtomType"
	case _ReduceExternNamedToAtomType:
		return "ExternNamedToAtomType"
	case _ReduceInferredToAtomType:
		return "InferredToAtomType"
	case _ReduceImplicitStructDefToAtomType:
		return "ImplicitStructDefToAtomType"
	case _ReduceImplicitEnumDefToAtomType:
		return "ImplicitEnumDefToAtomType"
	case _ReduceTraitDefToAtomType:
		return "TraitDefToAtomType"
	case _ReduceFuncTypeToAtomType:
		return "FuncTypeToAtomType"
	case _ReduceLexErrorToAtomType:
		return "LexErrorToAtomType"
	case _ReduceAtomTypeToValueType:
		return "AtomTypeToValueType"
	case _ReduceOptionalToValueType:
		return "OptionalToValueType"
	case _ReduceResultToValueType:
		return "ResultToValueType"
	case _ReduceReferenceToValueType:
		return "ReferenceToValueType"
	case _ReducePublicMethodsTraitToValueType:
		return "PublicMethodsTraitToValueType"
	case _ReducePublicTraitToValueType:
		return "PublicTraitToValueType"
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
	case _ReduceValueTypeToTraitAlgebraType:
		return "ValueTypeToTraitAlgebraType"
	case _ReduceIntersectToTraitAlgebraType:
		return "IntersectToTraitAlgebraType"
	case _ReduceUnionToTraitAlgebraType:
		return "UnionToTraitAlgebraType"
	case _ReduceDifferenceToTraitAlgebraType:
		return "DifferenceToTraitAlgebraType"
	case _ReduceExplicitToTraitProperty:
		return "ExplicitToTraitProperty"
	case _ReduceImplicitToTraitProperty:
		return "ImplicitToTraitProperty"
	case _ReduceMethodSignatureToTraitProperty:
		return "MethodSignatureToTraitProperty"
	case _ReduceUnsafeStatementToTraitProperty:
		return "UnsafeStatementToTraitProperty"
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
	_State366 = _StateId(366)
	_State367 = _StateId(367)
	_State368 = _StateId(368)
	_State369 = _StateId(369)
	_State370 = _StateId(370)
	_State371 = _StateId(371)
	_State372 = _StateId(372)
	_State373 = _StateId(373)
	_State374 = _StateId(374)
	_State375 = _StateId(375)
	_State376 = _StateId(376)
	_State377 = _StateId(377)
	_State378 = _StateId(378)
	_State379 = _StateId(379)
	_State380 = _StateId(380)
	_State381 = _StateId(381)
	_State382 = _StateId(382)
	_State383 = _StateId(383)
	_State384 = _StateId(384)
	_State385 = _StateId(385)
	_State386 = _StateId(386)
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
	case _EndMarker, SpacesToken, NewlinesToken, CommentToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, PackageToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, DotdotdotToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, NotToken, AndToken, OrToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, LexErrorToken:
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
	case _ReduceInitializeExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.InitializeExprToAtomExpr(args[0].Generic_)
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
	case _ReduceExplicitToInitializeExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = InitializeExprType
		symbol.Generic_, err = reducer.ExplicitToInitializeExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceImplicitStructToInitializeExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = InitializeExprType
		symbol.Generic_, err = reducer.ImplicitStructToInitializeExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	case _ReduceExplicitStructDefToInitializableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeType
		symbol.Generic_, err = reducer.ExplicitStructDefToInitializableType(args[0].Generic_)
	case _ReduceExplicitEnumDefToInitializableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeType
		symbol.Generic_, err = reducer.ExplicitEnumDefToInitializableType(args[0].Generic_)
	case _ReduceSliceToInitializableType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = InitializableTypeType
		symbol.Generic_, err = reducer.SliceToInitializableType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceArrayToInitializableType:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = InitializableTypeType
		symbol.Generic_, err = reducer.ArrayToInitializableType(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceMapToInitializableType:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = InitializableTypeType
		symbol.Generic_, err = reducer.MapToInitializableType(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_)
	case _ReduceInitializableTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.InitializableTypeToAtomType(args[0].Generic_)
	case _ReduceNamedToAtomType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.NamedToAtomType(args[0].Generic_, args[1].Generic_)
	case _ReduceExternNamedToAtomType:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ExternNamedToAtomType(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceInferredToAtomType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.InferredToAtomType(args[0].Generic_, args[1].Generic_)
	case _ReduceImplicitStructDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.ImplicitStructDefToAtomType(args[0].Generic_)
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
	case _ReduceFuncTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.FuncTypeToAtomType(args[0].Generic_)
	case _ReduceLexErrorToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.Generic_, err = reducer.LexErrorToAtomType(args[0].Generic_)
	case _ReduceAtomTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.AtomTypeToValueType(args[0].Generic_)
	case _ReduceOptionalToValueType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.OptionalToValueType(args[0].Generic_, args[1].Generic_)
	case _ReduceResultToValueType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.ResultToValueType(args[0].Generic_, args[1].Generic_)
	case _ReduceReferenceToValueType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.ReferenceToValueType(args[0].Generic_, args[1].Generic_)
	case _ReducePublicMethodsTraitToValueType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.PublicMethodsTraitToValueType(args[0].Generic_, args[1].Generic_)
	case _ReducePublicTraitToValueType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.PublicTraitToValueType(args[0].Generic_, args[1].Generic_)
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
	case _ReduceValueTypeToTraitAlgebraType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitAlgebraTypeType
		symbol.Generic_, err = reducer.ValueTypeToTraitAlgebraType(args[0].Generic_)
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
	case _ReduceExplicitToTraitProperty:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = TraitPropertyType
		symbol.Generic_, err = reducer.ExplicitToTraitProperty(args[0].Generic_, args[1].Generic_)
	case _ReduceImplicitToTraitProperty:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitPropertyType
		symbol.Generic_, err = reducer.ImplicitToTraitProperty(args[0].Generic_)
	case _ReduceMethodSignatureToTraitProperty:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitPropertyType
		symbol.Generic_, err = reducer.MethodSignatureToTraitProperty(args[0].Generic_)
	case _ReduceUnsafeStatementToTraitProperty:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitPropertyType
		symbol.Generic_, err = reducer.UnsafeStatementToTraitProperty(args[0].Generic_)
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
	_GotoState366Action                                             = &_Action{_ShiftAction, _State366, 0}
	_GotoState367Action                                             = &_Action{_ShiftAction, _State367, 0}
	_GotoState368Action                                             = &_Action{_ShiftAction, _State368, 0}
	_GotoState369Action                                             = &_Action{_ShiftAction, _State369, 0}
	_GotoState370Action                                             = &_Action{_ShiftAction, _State370, 0}
	_GotoState371Action                                             = &_Action{_ShiftAction, _State371, 0}
	_GotoState372Action                                             = &_Action{_ShiftAction, _State372, 0}
	_GotoState373Action                                             = &_Action{_ShiftAction, _State373, 0}
	_GotoState374Action                                             = &_Action{_ShiftAction, _State374, 0}
	_GotoState375Action                                             = &_Action{_ShiftAction, _State375, 0}
	_GotoState376Action                                             = &_Action{_ShiftAction, _State376, 0}
	_GotoState377Action                                             = &_Action{_ShiftAction, _State377, 0}
	_GotoState378Action                                             = &_Action{_ShiftAction, _State378, 0}
	_GotoState379Action                                             = &_Action{_ShiftAction, _State379, 0}
	_GotoState380Action                                             = &_Action{_ShiftAction, _State380, 0}
	_GotoState381Action                                             = &_Action{_ShiftAction, _State381, 0}
	_GotoState382Action                                             = &_Action{_ShiftAction, _State382, 0}
	_GotoState383Action                                             = &_Action{_ShiftAction, _State383, 0}
	_GotoState384Action                                             = &_Action{_ShiftAction, _State384, 0}
	_GotoState385Action                                             = &_Action{_ShiftAction, _State385, 0}
	_GotoState386Action                                             = &_Action{_ShiftAction, _State386, 0}
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
	_ReduceInitializeExprToAtomExprAction                           = &_Action{_ReduceAction, 0, _ReduceInitializeExprToAtomExpr}
	_ReduceLexErrorToAtomExprAction                                 = &_Action{_ReduceAction, 0, _ReduceLexErrorToAtomExpr}
	_ReduceTrueToLiteralAction                                      = &_Action{_ReduceAction, 0, _ReduceTrueToLiteral}
	_ReduceFalseToLiteralAction                                     = &_Action{_ReduceAction, 0, _ReduceFalseToLiteral}
	_ReduceIntegerLiteralToLiteralAction                            = &_Action{_ReduceAction, 0, _ReduceIntegerLiteralToLiteral}
	_ReduceFloatLiteralToLiteralAction                              = &_Action{_ReduceAction, 0, _ReduceFloatLiteralToLiteral}
	_ReduceRuneLiteralToLiteralAction                               = &_Action{_ReduceAction, 0, _ReduceRuneLiteralToLiteral}
	_ReduceStringLiteralToLiteralAction                             = &_Action{_ReduceAction, 0, _ReduceStringLiteralToLiteral}
	_ReduceExplicitToInitializeExprAction                           = &_Action{_ReduceAction, 0, _ReduceExplicitToInitializeExpr}
	_ReduceImplicitStructToInitializeExprAction                     = &_Action{_ReduceAction, 0, _ReduceImplicitStructToInitializeExpr}
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
	_ReduceExplicitStructDefToInitializableTypeAction               = &_Action{_ReduceAction, 0, _ReduceExplicitStructDefToInitializableType}
	_ReduceExplicitEnumDefToInitializableTypeAction                 = &_Action{_ReduceAction, 0, _ReduceExplicitEnumDefToInitializableType}
	_ReduceSliceToInitializableTypeAction                           = &_Action{_ReduceAction, 0, _ReduceSliceToInitializableType}
	_ReduceArrayToInitializableTypeAction                           = &_Action{_ReduceAction, 0, _ReduceArrayToInitializableType}
	_ReduceMapToInitializableTypeAction                             = &_Action{_ReduceAction, 0, _ReduceMapToInitializableType}
	_ReduceInitializableTypeToAtomTypeAction                        = &_Action{_ReduceAction, 0, _ReduceInitializableTypeToAtomType}
	_ReduceNamedToAtomTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceNamedToAtomType}
	_ReduceExternNamedToAtomTypeAction                              = &_Action{_ReduceAction, 0, _ReduceExternNamedToAtomType}
	_ReduceInferredToAtomTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceInferredToAtomType}
	_ReduceImplicitStructDefToAtomTypeAction                        = &_Action{_ReduceAction, 0, _ReduceImplicitStructDefToAtomType}
	_ReduceImplicitEnumDefToAtomTypeAction                          = &_Action{_ReduceAction, 0, _ReduceImplicitEnumDefToAtomType}
	_ReduceTraitDefToAtomTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceTraitDefToAtomType}
	_ReduceFuncTypeToAtomTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceFuncTypeToAtomType}
	_ReduceLexErrorToAtomTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceLexErrorToAtomType}
	_ReduceAtomTypeToValueTypeAction                                = &_Action{_ReduceAction, 0, _ReduceAtomTypeToValueType}
	_ReduceOptionalToValueTypeAction                                = &_Action{_ReduceAction, 0, _ReduceOptionalToValueType}
	_ReduceResultToValueTypeAction                                  = &_Action{_ReduceAction, 0, _ReduceResultToValueType}
	_ReduceReferenceToValueTypeAction                               = &_Action{_ReduceAction, 0, _ReduceReferenceToValueType}
	_ReducePublicMethodsTraitToValueTypeAction                      = &_Action{_ReduceAction, 0, _ReducePublicMethodsTraitToValueType}
	_ReducePublicTraitToValueTypeAction                             = &_Action{_ReduceAction, 0, _ReducePublicTraitToValueType}
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
	_ReduceValueTypeToTraitAlgebraTypeAction                        = &_Action{_ReduceAction, 0, _ReduceValueTypeToTraitAlgebraType}
	_ReduceIntersectToTraitAlgebraTypeAction                        = &_Action{_ReduceAction, 0, _ReduceIntersectToTraitAlgebraType}
	_ReduceUnionToTraitAlgebraTypeAction                            = &_Action{_ReduceAction, 0, _ReduceUnionToTraitAlgebraType}
	_ReduceDifferenceToTraitAlgebraTypeAction                       = &_Action{_ReduceAction, 0, _ReduceDifferenceToTraitAlgebraType}
	_ReduceExplicitToTraitPropertyAction                            = &_Action{_ReduceAction, 0, _ReduceExplicitToTraitProperty}
	_ReduceImplicitToTraitPropertyAction                            = &_Action{_ReduceAction, 0, _ReduceImplicitToTraitProperty}
	_ReduceMethodSignatureToTraitPropertyAction                     = &_Action{_ReduceAction, 0, _ReduceMethodSignatureToTraitProperty}
	_ReduceUnsafeStatementToTraitPropertyAction                     = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToTraitProperty}
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
	{_State5, IntegerLiteralToken}:                _GotoState26Action,
	{_State5, FloatLiteralToken}:                  _GotoState23Action,
	{_State5, RuneLiteralToken}:                   _GotoState33Action,
	{_State5, StringLiteralToken}:                 _GotoState34Action,
	{_State5, IdentifierToken}:                    _GotoState25Action,
	{_State5, TrueToken}:                          _GotoState37Action,
	{_State5, FalseToken}:                         _GotoState22Action,
	{_State5, StructToken}:                        _GotoState35Action,
	{_State5, EnumToken}:                          _GotoState21Action,
	{_State5, FuncToken}:                          _GotoState24Action,
	{_State5, LabelDeclToken}:                     _GotoState27Action,
	{_State5, LparenToken}:                        _GotoState30Action,
	{_State5, LbracketToken}:                      _GotoState28Action,
	{_State5, NotToken}:                           _GotoState32Action,
	{_State5, SubToken}:                           _GotoState36Action,
	{_State5, MulToken}:                           _GotoState31Action,
	{_State5, BitNegToken}:                        _GotoState20Action,
	{_State5, BitAndToken}:                        _GotoState19Action,
	{_State5, LexErrorToken}:                      _GotoState29Action,
	{_State5, ExpressionType}:                     _GotoState11Action,
	{_State5, OptionalLabelDeclType}:              _GotoState52Action,
	{_State5, SequenceExprType}:                   _GotoState57Action,
	{_State5, BlockExprType}:                      _GotoState43Action,
	{_State5, CallExprType}:                       _GotoState44Action,
	{_State5, AtomExprType}:                       _GotoState42Action,
	{_State5, LiteralType}:                        _GotoState50Action,
	{_State5, InitializeExprType}:                 _GotoState49Action,
	{_State5, AccessExprType}:                     _GotoState38Action,
	{_State5, PostfixUnaryExprType}:               _GotoState54Action,
	{_State5, PrefixUnaryOpType}:                  _GotoState56Action,
	{_State5, PrefixUnaryExprType}:                _GotoState55Action,
	{_State5, MulExprType}:                        _GotoState51Action,
	{_State5, AddExprType}:                        _GotoState39Action,
	{_State5, CmpExprType}:                        _GotoState45Action,
	{_State5, AndExprType}:                        _GotoState40Action,
	{_State5, OrExprType}:                         _GotoState53Action,
	{_State5, InitializableTypeType}:              _GotoState48Action,
	{_State5, ExplicitStructDefType}:              _GotoState47Action,
	{_State5, ExplicitEnumDefType}:                _GotoState46Action,
	{_State5, AnonymousFuncExprType}:              _GotoState41Action,
	{_State6, SpacesToken}:                        _GotoState59Action,
	{_State6, CommentToken}:                       _GotoState58Action,
	{_State6, LexInternalTokensType}:              _GotoState12Action,
	{_State15, PackageToken}:                      _GotoState16Action,
	{_State15, UnsafeToken}:                       _GotoState60Action,
	{_State15, TypeToken}:                         _GotoState17Action,
	{_State15, FuncToken}:                         _GotoState18Action,
	{_State15, DefinitionsType}:                   _GotoState62Action,
	{_State15, DefinitionType}:                    _GotoState61Action,
	{_State15, UnsafeStatementType}:               _GotoState66Action,
	{_State15, TypeDefType}:                       _GotoState65Action,
	{_State15, NamedFuncDefType}:                  _GotoState63Action,
	{_State15, PackageDefType}:                    _GotoState64Action,
	{_State16, IdentifierToken}:                   _GotoState67Action,
	{_State17, IdentifierToken}:                   _GotoState68Action,
	{_State18, LparenToken}:                       _GotoState69Action,
	{_State18, OptionalReceiverType}:              _GotoState70Action,
	{_State21, LparenToken}:                       _GotoState71Action,
	{_State24, LparenToken}:                       _GotoState72Action,
	{_State28, IdentifierToken}:                   _GotoState78Action,
	{_State28, StructToken}:                       _GotoState35Action,
	{_State28, EnumToken}:                         _GotoState21Action,
	{_State28, TraitToken}:                        _GotoState83Action,
	{_State28, FuncToken}:                         _GotoState77Action,
	{_State28, LparenToken}:                       _GotoState80Action,
	{_State28, LbracketToken}:                     _GotoState28Action,
	{_State28, DotToken}:                          _GotoState75Action,
	{_State28, QuestionToken}:                     _GotoState81Action,
	{_State28, ExclaimToken}:                      _GotoState76Action,
	{_State28, TildeTildeToken}:                   _GotoState82Action,
	{_State28, BitNegToken}:                       _GotoState74Action,
	{_State28, BitAndToken}:                       _GotoState73Action,
	{_State28, LexErrorToken}:                     _GotoState79Action,
	{_State28, InitializableTypeType}:             _GotoState88Action,
	{_State28, AtomTypeType}:                      _GotoState84Action,
	{_State28, ValueTypeType}:                     _GotoState90Action,
	{_State28, ImplicitStructDefType}:             _GotoState87Action,
	{_State28, ExplicitStructDefType}:             _GotoState47Action,
	{_State28, ImplicitEnumDefType}:               _GotoState86Action,
	{_State28, ExplicitEnumDefType}:               _GotoState46Action,
	{_State28, TraitDefType}:                      _GotoState89Action,
	{_State28, FuncTypeType}:                      _GotoState85Action,
	{_State30, IntegerLiteralToken}:               _GotoState26Action,
	{_State30, FloatLiteralToken}:                 _GotoState23Action,
	{_State30, RuneLiteralToken}:                  _GotoState33Action,
	{_State30, StringLiteralToken}:                _GotoState34Action,
	{_State30, IdentifierToken}:                   _GotoState91Action,
	{_State30, TrueToken}:                         _GotoState37Action,
	{_State30, FalseToken}:                        _GotoState22Action,
	{_State30, StructToken}:                       _GotoState35Action,
	{_State30, EnumToken}:                         _GotoState21Action,
	{_State30, FuncToken}:                         _GotoState24Action,
	{_State30, LabelDeclToken}:                    _GotoState27Action,
	{_State30, LparenToken}:                       _GotoState30Action,
	{_State30, LbracketToken}:                     _GotoState28Action,
	{_State30, NotToken}:                          _GotoState32Action,
	{_State30, SubToken}:                          _GotoState36Action,
	{_State30, MulToken}:                          _GotoState31Action,
	{_State30, BitNegToken}:                       _GotoState20Action,
	{_State30, BitAndToken}:                       _GotoState19Action,
	{_State30, LexErrorToken}:                     _GotoState29Action,
	{_State30, ExpressionType}:                    _GotoState95Action,
	{_State30, OptionalLabelDeclType}:             _GotoState52Action,
	{_State30, SequenceExprType}:                  _GotoState57Action,
	{_State30, BlockExprType}:                     _GotoState43Action,
	{_State30, CallExprType}:                      _GotoState44Action,
	{_State30, ArgumentsType}:                     _GotoState93Action,
	{_State30, ArgumentType}:                      _GotoState92Action,
	{_State30, ColonExpressionsType}:              _GotoState94Action,
	{_State30, OptionalExpressionType}:            _GotoState96Action,
	{_State30, AtomExprType}:                      _GotoState42Action,
	{_State30, LiteralType}:                       _GotoState50Action,
	{_State30, InitializeExprType}:                _GotoState49Action,
	{_State30, AccessExprType}:                    _GotoState38Action,
	{_State30, PostfixUnaryExprType}:              _GotoState54Action,
	{_State30, PrefixUnaryOpType}:                 _GotoState56Action,
	{_State30, PrefixUnaryExprType}:               _GotoState55Action,
	{_State30, MulExprType}:                       _GotoState51Action,
	{_State30, AddExprType}:                       _GotoState39Action,
	{_State30, CmpExprType}:                       _GotoState45Action,
	{_State30, AndExprType}:                       _GotoState40Action,
	{_State30, OrExprType}:                        _GotoState53Action,
	{_State30, InitializableTypeType}:             _GotoState48Action,
	{_State30, ExplicitStructDefType}:             _GotoState47Action,
	{_State30, ExplicitEnumDefType}:               _GotoState46Action,
	{_State30, AnonymousFuncExprType}:             _GotoState41Action,
	{_State35, LparenToken}:                       _GotoState97Action,
	{_State38, LbracketToken}:                     _GotoState100Action,
	{_State38, DotToken}:                          _GotoState99Action,
	{_State38, QuestionToken}:                     _GotoState101Action,
	{_State38, DollarLbracketToken}:               _GotoState98Action,
	{_State38, OptionalGenericBindingType}:        _GotoState102Action,
	{_State39, AddToken}:                          _GotoState103Action,
	{_State39, SubToken}:                          _GotoState106Action,
	{_State39, BitXorToken}:                       _GotoState105Action,
	{_State39, BitOrToken}:                        _GotoState104Action,
	{_State39, AddOpType}:                         _GotoState107Action,
	{_State40, AndToken}:                          _GotoState108Action,
	{_State45, EqualToken}:                        _GotoState109Action,
	{_State45, NotEqualToken}:                     _GotoState114Action,
	{_State45, LessToken}:                         _GotoState112Action,
	{_State45, LessOrEqualToken}:                  _GotoState113Action,
	{_State45, GreaterToken}:                      _GotoState110Action,
	{_State45, GreaterOrEqualToken}:               _GotoState111Action,
	{_State45, CmpOpType}:                         _GotoState115Action,
	{_State48, LparenToken}:                       _GotoState116Action,
	{_State51, MulToken}:                          _GotoState122Action,
	{_State51, DivToken}:                          _GotoState120Action,
	{_State51, ModToken}:                          _GotoState121Action,
	{_State51, BitAndToken}:                       _GotoState117Action,
	{_State51, BitLshiftToken}:                    _GotoState118Action,
	{_State51, BitRshiftToken}:                    _GotoState119Action,
	{_State51, MulOpType}:                         _GotoState123Action,
	{_State52, IfToken}:                           _GotoState126Action,
	{_State52, SwitchToken}:                       _GotoState128Action,
	{_State52, ForToken}:                          _GotoState125Action,
	{_State52, DoToken}:                           _GotoState124Action,
	{_State52, LbraceToken}:                       _GotoState127Action,
	{_State52, IfExprType}:                        _GotoState130Action,
	{_State52, SwitchExprType}:                    _GotoState132Action,
	{_State52, LoopExprType}:                      _GotoState131Action,
	{_State52, BlockBodyType}:                     _GotoState129Action,
	{_State53, OrToken}:                           _GotoState133Action,
	{_State56, IntegerLiteralToken}:               _GotoState26Action,
	{_State56, FloatLiteralToken}:                 _GotoState23Action,
	{_State56, RuneLiteralToken}:                  _GotoState33Action,
	{_State56, StringLiteralToken}:                _GotoState34Action,
	{_State56, IdentifierToken}:                   _GotoState25Action,
	{_State56, TrueToken}:                         _GotoState37Action,
	{_State56, FalseToken}:                        _GotoState22Action,
	{_State56, StructToken}:                       _GotoState35Action,
	{_State56, EnumToken}:                         _GotoState21Action,
	{_State56, FuncToken}:                         _GotoState24Action,
	{_State56, LabelDeclToken}:                    _GotoState27Action,
	{_State56, LparenToken}:                       _GotoState30Action,
	{_State56, LbracketToken}:                     _GotoState28Action,
	{_State56, NotToken}:                          _GotoState32Action,
	{_State56, SubToken}:                          _GotoState36Action,
	{_State56, MulToken}:                          _GotoState31Action,
	{_State56, BitNegToken}:                       _GotoState20Action,
	{_State56, BitAndToken}:                       _GotoState19Action,
	{_State56, LexErrorToken}:                     _GotoState29Action,
	{_State56, OptionalLabelDeclType}:             _GotoState134Action,
	{_State56, BlockExprType}:                     _GotoState43Action,
	{_State56, CallExprType}:                      _GotoState44Action,
	{_State56, AtomExprType}:                      _GotoState42Action,
	{_State56, LiteralType}:                       _GotoState50Action,
	{_State56, InitializeExprType}:                _GotoState49Action,
	{_State56, AccessExprType}:                    _GotoState38Action,
	{_State56, PostfixUnaryExprType}:              _GotoState54Action,
	{_State56, PrefixUnaryOpType}:                 _GotoState56Action,
	{_State56, PrefixUnaryExprType}:               _GotoState135Action,
	{_State56, InitializableTypeType}:             _GotoState48Action,
	{_State56, ExplicitStructDefType}:             _GotoState47Action,
	{_State56, ExplicitEnumDefType}:               _GotoState46Action,
	{_State56, AnonymousFuncExprType}:             _GotoState41Action,
	{_State60, LessToken}:                         _GotoState136Action,
	{_State62, NewlinesToken}:                     _GotoState137Action,
	{_State62, OptionalNewlinesType}:              _GotoState138Action,
	{_State67, LparenToken}:                       _GotoState139Action,
	{_State68, DollarLbracketToken}:               _GotoState140Action,
	{_State68, EqualToken}:                        _GotoState141Action,
	{_State68, OptionalGenericParametersType}:     _GotoState142Action,
	{_State69, IdentifierToken}:                   _GotoState143Action,
	{_State69, ParameterDefType}:                  _GotoState144Action,
	{_State70, IdentifierToken}:                   _GotoState145Action,
	{_State71, IdentifierToken}:                   _GotoState146Action,
	{_State71, UnsafeToken}:                       _GotoState60Action,
	{_State71, StructToken}:                       _GotoState35Action,
	{_State71, EnumToken}:                         _GotoState21Action,
	{_State71, TraitToken}:                        _GotoState83Action,
	{_State71, FuncToken}:                         _GotoState77Action,
	{_State71, LparenToken}:                       _GotoState80Action,
	{_State71, LbracketToken}:                     _GotoState28Action,
	{_State71, DotToken}:                          _GotoState75Action,
	{_State71, QuestionToken}:                     _GotoState81Action,
	{_State71, ExclaimToken}:                      _GotoState76Action,
	{_State71, TildeTildeToken}:                   _GotoState82Action,
	{_State71, BitNegToken}:                       _GotoState74Action,
	{_State71, BitAndToken}:                       _GotoState73Action,
	{_State71, LexErrorToken}:                     _GotoState79Action,
	{_State71, UnsafeStatementType}:               _GotoState151Action,
	{_State71, InitializableTypeType}:             _GotoState88Action,
	{_State71, AtomTypeType}:                      _GotoState84Action,
	{_State71, ValueTypeType}:                     _GotoState152Action,
	{_State71, FieldDefType}:                      _GotoState149Action,
	{_State71, ImplicitStructDefType}:             _GotoState87Action,
	{_State71, ExplicitStructDefType}:             _GotoState47Action,
	{_State71, EnumValueDefType}:                  _GotoState147Action,
	{_State71, ImplicitEnumValueDefsType}:         _GotoState150Action,
	{_State71, ImplicitEnumDefType}:               _GotoState86Action,
	{_State71, ExplicitEnumValueDefsType}:         _GotoState148Action,
	{_State71, ExplicitEnumDefType}:               _GotoState46Action,
	{_State71, TraitDefType}:                      _GotoState89Action,
	{_State71, FuncTypeType}:                      _GotoState85Action,
	{_State72, IdentifierToken}:                   _GotoState143Action,
	{_State72, ParameterDefType}:                  _GotoState154Action,
	{_State72, ParameterDefsType}:                 _GotoState155Action,
	{_State72, OptionalParameterDefsType}:         _GotoState153Action,
	{_State73, IdentifierToken}:                   _GotoState78Action,
	{_State73, StructToken}:                       _GotoState35Action,
	{_State73, EnumToken}:                         _GotoState21Action,
	{_State73, TraitToken}:                        _GotoState83Action,
	{_State73, FuncToken}:                         _GotoState77Action,
	{_State73, LparenToken}:                       _GotoState80Action,
	{_State73, LbracketToken}:                     _GotoState28Action,
	{_State73, DotToken}:                          _GotoState75Action,
	{_State73, QuestionToken}:                     _GotoState81Action,
	{_State73, ExclaimToken}:                      _GotoState76Action,
	{_State73, TildeTildeToken}:                   _GotoState82Action,
	{_State73, BitNegToken}:                       _GotoState74Action,
	{_State73, BitAndToken}:                       _GotoState73Action,
	{_State73, LexErrorToken}:                     _GotoState79Action,
	{_State73, InitializableTypeType}:             _GotoState88Action,
	{_State73, AtomTypeType}:                      _GotoState84Action,
	{_State73, ValueTypeType}:                     _GotoState156Action,
	{_State73, ImplicitStructDefType}:             _GotoState87Action,
	{_State73, ExplicitStructDefType}:             _GotoState47Action,
	{_State73, ImplicitEnumDefType}:               _GotoState86Action,
	{_State73, ExplicitEnumDefType}:               _GotoState46Action,
	{_State73, TraitDefType}:                      _GotoState89Action,
	{_State73, FuncTypeType}:                      _GotoState85Action,
	{_State74, IdentifierToken}:                   _GotoState78Action,
	{_State74, StructToken}:                       _GotoState35Action,
	{_State74, EnumToken}:                         _GotoState21Action,
	{_State74, TraitToken}:                        _GotoState83Action,
	{_State74, FuncToken}:                         _GotoState77Action,
	{_State74, LparenToken}:                       _GotoState80Action,
	{_State74, LbracketToken}:                     _GotoState28Action,
	{_State74, DotToken}:                          _GotoState75Action,
	{_State74, QuestionToken}:                     _GotoState81Action,
	{_State74, ExclaimToken}:                      _GotoState76Action,
	{_State74, TildeTildeToken}:                   _GotoState82Action,
	{_State74, BitNegToken}:                       _GotoState74Action,
	{_State74, BitAndToken}:                       _GotoState73Action,
	{_State74, LexErrorToken}:                     _GotoState79Action,
	{_State74, InitializableTypeType}:             _GotoState88Action,
	{_State74, AtomTypeType}:                      _GotoState84Action,
	{_State74, ValueTypeType}:                     _GotoState157Action,
	{_State74, ImplicitStructDefType}:             _GotoState87Action,
	{_State74, ExplicitStructDefType}:             _GotoState47Action,
	{_State74, ImplicitEnumDefType}:               _GotoState86Action,
	{_State74, ExplicitEnumDefType}:               _GotoState46Action,
	{_State74, TraitDefType}:                      _GotoState89Action,
	{_State74, FuncTypeType}:                      _GotoState85Action,
	{_State75, DollarLbracketToken}:               _GotoState98Action,
	{_State75, OptionalGenericBindingType}:        _GotoState158Action,
	{_State76, IdentifierToken}:                   _GotoState78Action,
	{_State76, StructToken}:                       _GotoState35Action,
	{_State76, EnumToken}:                         _GotoState21Action,
	{_State76, TraitToken}:                        _GotoState83Action,
	{_State76, FuncToken}:                         _GotoState77Action,
	{_State76, LparenToken}:                       _GotoState80Action,
	{_State76, LbracketToken}:                     _GotoState28Action,
	{_State76, DotToken}:                          _GotoState75Action,
	{_State76, QuestionToken}:                     _GotoState81Action,
	{_State76, ExclaimToken}:                      _GotoState76Action,
	{_State76, TildeTildeToken}:                   _GotoState82Action,
	{_State76, BitNegToken}:                       _GotoState74Action,
	{_State76, BitAndToken}:                       _GotoState73Action,
	{_State76, LexErrorToken}:                     _GotoState79Action,
	{_State76, InitializableTypeType}:             _GotoState88Action,
	{_State76, AtomTypeType}:                      _GotoState84Action,
	{_State76, ValueTypeType}:                     _GotoState159Action,
	{_State76, ImplicitStructDefType}:             _GotoState87Action,
	{_State76, ExplicitStructDefType}:             _GotoState47Action,
	{_State76, ImplicitEnumDefType}:               _GotoState86Action,
	{_State76, ExplicitEnumDefType}:               _GotoState46Action,
	{_State76, TraitDefType}:                      _GotoState89Action,
	{_State76, FuncTypeType}:                      _GotoState85Action,
	{_State77, LparenToken}:                       _GotoState160Action,
	{_State78, DotToken}:                          _GotoState161Action,
	{_State78, DollarLbracketToken}:               _GotoState98Action,
	{_State78, OptionalGenericBindingType}:        _GotoState162Action,
	{_State80, IdentifierToken}:                   _GotoState146Action,
	{_State80, UnsafeToken}:                       _GotoState60Action,
	{_State80, StructToken}:                       _GotoState35Action,
	{_State80, EnumToken}:                         _GotoState21Action,
	{_State80, TraitToken}:                        _GotoState83Action,
	{_State80, FuncToken}:                         _GotoState77Action,
	{_State80, LparenToken}:                       _GotoState80Action,
	{_State80, LbracketToken}:                     _GotoState28Action,
	{_State80, DotToken}:                          _GotoState75Action,
	{_State80, QuestionToken}:                     _GotoState81Action,
	{_State80, ExclaimToken}:                      _GotoState76Action,
	{_State80, TildeTildeToken}:                   _GotoState82Action,
	{_State80, BitNegToken}:                       _GotoState74Action,
	{_State80, BitAndToken}:                       _GotoState73Action,
	{_State80, LexErrorToken}:                     _GotoState79Action,
	{_State80, UnsafeStatementType}:               _GotoState151Action,
	{_State80, InitializableTypeType}:             _GotoState88Action,
	{_State80, AtomTypeType}:                      _GotoState84Action,
	{_State80, ValueTypeType}:                     _GotoState152Action,
	{_State80, FieldDefType}:                      _GotoState164Action,
	{_State80, ImplicitFieldDefsType}:             _GotoState166Action,
	{_State80, OptionalImplicitFieldDefsType}:     _GotoState167Action,
	{_State80, ImplicitStructDefType}:             _GotoState87Action,
	{_State80, ExplicitStructDefType}:             _GotoState47Action,
	{_State80, EnumValueDefType}:                  _GotoState163Action,
	{_State80, ImplicitEnumValueDefsType}:         _GotoState165Action,
	{_State80, ImplicitEnumDefType}:               _GotoState86Action,
	{_State80, ExplicitEnumDefType}:               _GotoState46Action,
	{_State80, TraitDefType}:                      _GotoState89Action,
	{_State80, FuncTypeType}:                      _GotoState85Action,
	{_State81, IdentifierToken}:                   _GotoState78Action,
	{_State81, StructToken}:                       _GotoState35Action,
	{_State81, EnumToken}:                         _GotoState21Action,
	{_State81, TraitToken}:                        _GotoState83Action,
	{_State81, FuncToken}:                         _GotoState77Action,
	{_State81, LparenToken}:                       _GotoState80Action,
	{_State81, LbracketToken}:                     _GotoState28Action,
	{_State81, DotToken}:                          _GotoState75Action,
	{_State81, QuestionToken}:                     _GotoState81Action,
	{_State81, ExclaimToken}:                      _GotoState76Action,
	{_State81, TildeTildeToken}:                   _GotoState82Action,
	{_State81, BitNegToken}:                       _GotoState74Action,
	{_State81, BitAndToken}:                       _GotoState73Action,
	{_State81, LexErrorToken}:                     _GotoState79Action,
	{_State81, InitializableTypeType}:             _GotoState88Action,
	{_State81, AtomTypeType}:                      _GotoState84Action,
	{_State81, ValueTypeType}:                     _GotoState168Action,
	{_State81, ImplicitStructDefType}:             _GotoState87Action,
	{_State81, ExplicitStructDefType}:             _GotoState47Action,
	{_State81, ImplicitEnumDefType}:               _GotoState86Action,
	{_State81, ExplicitEnumDefType}:               _GotoState46Action,
	{_State81, TraitDefType}:                      _GotoState89Action,
	{_State81, FuncTypeType}:                      _GotoState85Action,
	{_State82, IdentifierToken}:                   _GotoState78Action,
	{_State82, StructToken}:                       _GotoState35Action,
	{_State82, EnumToken}:                         _GotoState21Action,
	{_State82, TraitToken}:                        _GotoState83Action,
	{_State82, FuncToken}:                         _GotoState77Action,
	{_State82, LparenToken}:                       _GotoState80Action,
	{_State82, LbracketToken}:                     _GotoState28Action,
	{_State82, DotToken}:                          _GotoState75Action,
	{_State82, QuestionToken}:                     _GotoState81Action,
	{_State82, ExclaimToken}:                      _GotoState76Action,
	{_State82, TildeTildeToken}:                   _GotoState82Action,
	{_State82, BitNegToken}:                       _GotoState74Action,
	{_State82, BitAndToken}:                       _GotoState73Action,
	{_State82, LexErrorToken}:                     _GotoState79Action,
	{_State82, InitializableTypeType}:             _GotoState88Action,
	{_State82, AtomTypeType}:                      _GotoState84Action,
	{_State82, ValueTypeType}:                     _GotoState169Action,
	{_State82, ImplicitStructDefType}:             _GotoState87Action,
	{_State82, ExplicitStructDefType}:             _GotoState47Action,
	{_State82, ImplicitEnumDefType}:               _GotoState86Action,
	{_State82, ExplicitEnumDefType}:               _GotoState46Action,
	{_State82, TraitDefType}:                      _GotoState89Action,
	{_State82, FuncTypeType}:                      _GotoState85Action,
	{_State83, LparenToken}:                       _GotoState170Action,
	{_State90, RbracketToken}:                     _GotoState173Action,
	{_State90, CommaToken}:                        _GotoState172Action,
	{_State90, ColonToken}:                        _GotoState171Action,
	{_State91, AssignToken}:                       _GotoState174Action,
	{_State93, RparenToken}:                       _GotoState176Action,
	{_State93, CommaToken}:                        _GotoState175Action,
	{_State94, ColonToken}:                        _GotoState177Action,
	{_State96, ColonToken}:                        _GotoState178Action,
	{_State97, IdentifierToken}:                   _GotoState146Action,
	{_State97, UnsafeToken}:                       _GotoState60Action,
	{_State97, StructToken}:                       _GotoState35Action,
	{_State97, EnumToken}:                         _GotoState21Action,
	{_State97, TraitToken}:                        _GotoState83Action,
	{_State97, FuncToken}:                         _GotoState77Action,
	{_State97, LparenToken}:                       _GotoState80Action,
	{_State97, LbracketToken}:                     _GotoState28Action,
	{_State97, DotToken}:                          _GotoState75Action,
	{_State97, QuestionToken}:                     _GotoState81Action,
	{_State97, ExclaimToken}:                      _GotoState76Action,
	{_State97, TildeTildeToken}:                   _GotoState82Action,
	{_State97, BitNegToken}:                       _GotoState74Action,
	{_State97, BitAndToken}:                       _GotoState73Action,
	{_State97, LexErrorToken}:                     _GotoState79Action,
	{_State97, UnsafeStatementType}:               _GotoState151Action,
	{_State97, InitializableTypeType}:             _GotoState88Action,
	{_State97, AtomTypeType}:                      _GotoState84Action,
	{_State97, ValueTypeType}:                     _GotoState152Action,
	{_State97, FieldDefType}:                      _GotoState180Action,
	{_State97, ImplicitStructDefType}:             _GotoState87Action,
	{_State97, ExplicitFieldDefsType}:             _GotoState179Action,
	{_State97, OptionalExplicitFieldDefsType}:     _GotoState181Action,
	{_State97, ExplicitStructDefType}:             _GotoState47Action,
	{_State97, ImplicitEnumDefType}:               _GotoState86Action,
	{_State97, ExplicitEnumDefType}:               _GotoState46Action,
	{_State97, TraitDefType}:                      _GotoState89Action,
	{_State97, FuncTypeType}:                      _GotoState85Action,
	{_State98, IdentifierToken}:                   _GotoState78Action,
	{_State98, StructToken}:                       _GotoState35Action,
	{_State98, EnumToken}:                         _GotoState21Action,
	{_State98, TraitToken}:                        _GotoState83Action,
	{_State98, FuncToken}:                         _GotoState77Action,
	{_State98, LparenToken}:                       _GotoState80Action,
	{_State98, LbracketToken}:                     _GotoState28Action,
	{_State98, DotToken}:                          _GotoState75Action,
	{_State98, QuestionToken}:                     _GotoState81Action,
	{_State98, ExclaimToken}:                      _GotoState76Action,
	{_State98, TildeTildeToken}:                   _GotoState82Action,
	{_State98, BitNegToken}:                       _GotoState74Action,
	{_State98, BitAndToken}:                       _GotoState73Action,
	{_State98, LexErrorToken}:                     _GotoState79Action,
	{_State98, OptionalGenericArgumentsType}:      _GotoState183Action,
	{_State98, GenericArgumentsType}:              _GotoState182Action,
	{_State98, InitializableTypeType}:             _GotoState88Action,
	{_State98, AtomTypeType}:                      _GotoState84Action,
	{_State98, ValueTypeType}:                     _GotoState184Action,
	{_State98, ImplicitStructDefType}:             _GotoState87Action,
	{_State98, ExplicitStructDefType}:             _GotoState47Action,
	{_State98, ImplicitEnumDefType}:               _GotoState86Action,
	{_State98, ExplicitEnumDefType}:               _GotoState46Action,
	{_State98, TraitDefType}:                      _GotoState89Action,
	{_State98, FuncTypeType}:                      _GotoState85Action,
	{_State99, IdentifierToken}:                   _GotoState185Action,
	{_State100, IntegerLiteralToken}:              _GotoState26Action,
	{_State100, FloatLiteralToken}:                _GotoState23Action,
	{_State100, RuneLiteralToken}:                 _GotoState33Action,
	{_State100, StringLiteralToken}:               _GotoState34Action,
	{_State100, IdentifierToken}:                  _GotoState91Action,
	{_State100, TrueToken}:                        _GotoState37Action,
	{_State100, FalseToken}:                       _GotoState22Action,
	{_State100, StructToken}:                      _GotoState35Action,
	{_State100, EnumToken}:                        _GotoState21Action,
	{_State100, FuncToken}:                        _GotoState24Action,
	{_State100, LabelDeclToken}:                   _GotoState27Action,
	{_State100, LparenToken}:                      _GotoState30Action,
	{_State100, LbracketToken}:                    _GotoState28Action,
	{_State100, NotToken}:                         _GotoState32Action,
	{_State100, SubToken}:                         _GotoState36Action,
	{_State100, MulToken}:                         _GotoState31Action,
	{_State100, BitNegToken}:                      _GotoState20Action,
	{_State100, BitAndToken}:                      _GotoState19Action,
	{_State100, LexErrorToken}:                    _GotoState29Action,
	{_State100, ExpressionType}:                   _GotoState95Action,
	{_State100, OptionalLabelDeclType}:            _GotoState52Action,
	{_State100, SequenceExprType}:                 _GotoState57Action,
	{_State100, BlockExprType}:                    _GotoState43Action,
	{_State100, CallExprType}:                     _GotoState44Action,
	{_State100, ArgumentType}:                     _GotoState186Action,
	{_State100, ColonExpressionsType}:             _GotoState94Action,
	{_State100, OptionalExpressionType}:           _GotoState96Action,
	{_State100, AtomExprType}:                     _GotoState42Action,
	{_State100, LiteralType}:                      _GotoState50Action,
	{_State100, InitializeExprType}:               _GotoState49Action,
	{_State100, AccessExprType}:                   _GotoState38Action,
	{_State100, PostfixUnaryExprType}:             _GotoState54Action,
	{_State100, PrefixUnaryOpType}:                _GotoState56Action,
	{_State100, PrefixUnaryExprType}:              _GotoState55Action,
	{_State100, MulExprType}:                      _GotoState51Action,
	{_State100, AddExprType}:                      _GotoState39Action,
	{_State100, CmpExprType}:                      _GotoState45Action,
	{_State100, AndExprType}:                      _GotoState40Action,
	{_State100, OrExprType}:                       _GotoState53Action,
	{_State100, InitializableTypeType}:            _GotoState48Action,
	{_State100, ExplicitStructDefType}:            _GotoState47Action,
	{_State100, ExplicitEnumDefType}:              _GotoState46Action,
	{_State100, AnonymousFuncExprType}:            _GotoState41Action,
	{_State102, LparenToken}:                      _GotoState187Action,
	{_State107, IntegerLiteralToken}:              _GotoState26Action,
	{_State107, FloatLiteralToken}:                _GotoState23Action,
	{_State107, RuneLiteralToken}:                 _GotoState33Action,
	{_State107, StringLiteralToken}:               _GotoState34Action,
	{_State107, IdentifierToken}:                  _GotoState25Action,
	{_State107, TrueToken}:                        _GotoState37Action,
	{_State107, FalseToken}:                       _GotoState22Action,
	{_State107, StructToken}:                      _GotoState35Action,
	{_State107, EnumToken}:                        _GotoState21Action,
	{_State107, FuncToken}:                        _GotoState24Action,
	{_State107, LabelDeclToken}:                   _GotoState27Action,
	{_State107, LparenToken}:                      _GotoState30Action,
	{_State107, LbracketToken}:                    _GotoState28Action,
	{_State107, NotToken}:                         _GotoState32Action,
	{_State107, SubToken}:                         _GotoState36Action,
	{_State107, MulToken}:                         _GotoState31Action,
	{_State107, BitNegToken}:                      _GotoState20Action,
	{_State107, BitAndToken}:                      _GotoState19Action,
	{_State107, LexErrorToken}:                    _GotoState29Action,
	{_State107, OptionalLabelDeclType}:            _GotoState134Action,
	{_State107, BlockExprType}:                    _GotoState43Action,
	{_State107, CallExprType}:                     _GotoState44Action,
	{_State107, AtomExprType}:                     _GotoState42Action,
	{_State107, LiteralType}:                      _GotoState50Action,
	{_State107, InitializeExprType}:               _GotoState49Action,
	{_State107, AccessExprType}:                   _GotoState38Action,
	{_State107, PostfixUnaryExprType}:             _GotoState54Action,
	{_State107, PrefixUnaryOpType}:                _GotoState56Action,
	{_State107, PrefixUnaryExprType}:              _GotoState55Action,
	{_State107, MulExprType}:                      _GotoState188Action,
	{_State107, InitializableTypeType}:            _GotoState48Action,
	{_State107, ExplicitStructDefType}:            _GotoState47Action,
	{_State107, ExplicitEnumDefType}:              _GotoState46Action,
	{_State107, AnonymousFuncExprType}:            _GotoState41Action,
	{_State108, IntegerLiteralToken}:              _GotoState26Action,
	{_State108, FloatLiteralToken}:                _GotoState23Action,
	{_State108, RuneLiteralToken}:                 _GotoState33Action,
	{_State108, StringLiteralToken}:               _GotoState34Action,
	{_State108, IdentifierToken}:                  _GotoState25Action,
	{_State108, TrueToken}:                        _GotoState37Action,
	{_State108, FalseToken}:                       _GotoState22Action,
	{_State108, StructToken}:                      _GotoState35Action,
	{_State108, EnumToken}:                        _GotoState21Action,
	{_State108, FuncToken}:                        _GotoState24Action,
	{_State108, LabelDeclToken}:                   _GotoState27Action,
	{_State108, LparenToken}:                      _GotoState30Action,
	{_State108, LbracketToken}:                    _GotoState28Action,
	{_State108, NotToken}:                         _GotoState32Action,
	{_State108, SubToken}:                         _GotoState36Action,
	{_State108, MulToken}:                         _GotoState31Action,
	{_State108, BitNegToken}:                      _GotoState20Action,
	{_State108, BitAndToken}:                      _GotoState19Action,
	{_State108, LexErrorToken}:                    _GotoState29Action,
	{_State108, OptionalLabelDeclType}:            _GotoState134Action,
	{_State108, BlockExprType}:                    _GotoState43Action,
	{_State108, CallExprType}:                     _GotoState44Action,
	{_State108, AtomExprType}:                     _GotoState42Action,
	{_State108, LiteralType}:                      _GotoState50Action,
	{_State108, InitializeExprType}:               _GotoState49Action,
	{_State108, AccessExprType}:                   _GotoState38Action,
	{_State108, PostfixUnaryExprType}:             _GotoState54Action,
	{_State108, PrefixUnaryOpType}:                _GotoState56Action,
	{_State108, PrefixUnaryExprType}:              _GotoState55Action,
	{_State108, MulExprType}:                      _GotoState51Action,
	{_State108, AddExprType}:                      _GotoState39Action,
	{_State108, CmpExprType}:                      _GotoState189Action,
	{_State108, InitializableTypeType}:            _GotoState48Action,
	{_State108, ExplicitStructDefType}:            _GotoState47Action,
	{_State108, ExplicitEnumDefType}:              _GotoState46Action,
	{_State108, AnonymousFuncExprType}:            _GotoState41Action,
	{_State115, IntegerLiteralToken}:              _GotoState26Action,
	{_State115, FloatLiteralToken}:                _GotoState23Action,
	{_State115, RuneLiteralToken}:                 _GotoState33Action,
	{_State115, StringLiteralToken}:               _GotoState34Action,
	{_State115, IdentifierToken}:                  _GotoState25Action,
	{_State115, TrueToken}:                        _GotoState37Action,
	{_State115, FalseToken}:                       _GotoState22Action,
	{_State115, StructToken}:                      _GotoState35Action,
	{_State115, EnumToken}:                        _GotoState21Action,
	{_State115, FuncToken}:                        _GotoState24Action,
	{_State115, LabelDeclToken}:                   _GotoState27Action,
	{_State115, LparenToken}:                      _GotoState30Action,
	{_State115, LbracketToken}:                    _GotoState28Action,
	{_State115, NotToken}:                         _GotoState32Action,
	{_State115, SubToken}:                         _GotoState36Action,
	{_State115, MulToken}:                         _GotoState31Action,
	{_State115, BitNegToken}:                      _GotoState20Action,
	{_State115, BitAndToken}:                      _GotoState19Action,
	{_State115, LexErrorToken}:                    _GotoState29Action,
	{_State115, OptionalLabelDeclType}:            _GotoState134Action,
	{_State115, BlockExprType}:                    _GotoState43Action,
	{_State115, CallExprType}:                     _GotoState44Action,
	{_State115, AtomExprType}:                     _GotoState42Action,
	{_State115, LiteralType}:                      _GotoState50Action,
	{_State115, InitializeExprType}:               _GotoState49Action,
	{_State115, AccessExprType}:                   _GotoState38Action,
	{_State115, PostfixUnaryExprType}:             _GotoState54Action,
	{_State115, PrefixUnaryOpType}:                _GotoState56Action,
	{_State115, PrefixUnaryExprType}:              _GotoState55Action,
	{_State115, MulExprType}:                      _GotoState51Action,
	{_State115, AddExprType}:                      _GotoState190Action,
	{_State115, InitializableTypeType}:            _GotoState48Action,
	{_State115, ExplicitStructDefType}:            _GotoState47Action,
	{_State115, ExplicitEnumDefType}:              _GotoState46Action,
	{_State115, AnonymousFuncExprType}:            _GotoState41Action,
	{_State116, IntegerLiteralToken}:              _GotoState26Action,
	{_State116, FloatLiteralToken}:                _GotoState23Action,
	{_State116, RuneLiteralToken}:                 _GotoState33Action,
	{_State116, StringLiteralToken}:               _GotoState34Action,
	{_State116, IdentifierToken}:                  _GotoState91Action,
	{_State116, TrueToken}:                        _GotoState37Action,
	{_State116, FalseToken}:                       _GotoState22Action,
	{_State116, StructToken}:                      _GotoState35Action,
	{_State116, EnumToken}:                        _GotoState21Action,
	{_State116, FuncToken}:                        _GotoState24Action,
	{_State116, LabelDeclToken}:                   _GotoState27Action,
	{_State116, LparenToken}:                      _GotoState30Action,
	{_State116, LbracketToken}:                    _GotoState28Action,
	{_State116, NotToken}:                         _GotoState32Action,
	{_State116, SubToken}:                         _GotoState36Action,
	{_State116, MulToken}:                         _GotoState31Action,
	{_State116, BitNegToken}:                      _GotoState20Action,
	{_State116, BitAndToken}:                      _GotoState19Action,
	{_State116, LexErrorToken}:                    _GotoState29Action,
	{_State116, ExpressionType}:                   _GotoState95Action,
	{_State116, OptionalLabelDeclType}:            _GotoState52Action,
	{_State116, SequenceExprType}:                 _GotoState57Action,
	{_State116, BlockExprType}:                    _GotoState43Action,
	{_State116, CallExprType}:                     _GotoState44Action,
	{_State116, ArgumentsType}:                    _GotoState191Action,
	{_State116, ArgumentType}:                     _GotoState92Action,
	{_State116, ColonExpressionsType}:             _GotoState94Action,
	{_State116, OptionalExpressionType}:           _GotoState96Action,
	{_State116, AtomExprType}:                     _GotoState42Action,
	{_State116, LiteralType}:                      _GotoState50Action,
	{_State116, InitializeExprType}:               _GotoState49Action,
	{_State116, AccessExprType}:                   _GotoState38Action,
	{_State116, PostfixUnaryExprType}:             _GotoState54Action,
	{_State116, PrefixUnaryOpType}:                _GotoState56Action,
	{_State116, PrefixUnaryExprType}:              _GotoState55Action,
	{_State116, MulExprType}:                      _GotoState51Action,
	{_State116, AddExprType}:                      _GotoState39Action,
	{_State116, CmpExprType}:                      _GotoState45Action,
	{_State116, AndExprType}:                      _GotoState40Action,
	{_State116, OrExprType}:                       _GotoState53Action,
	{_State116, InitializableTypeType}:            _GotoState48Action,
	{_State116, ExplicitStructDefType}:            _GotoState47Action,
	{_State116, ExplicitEnumDefType}:              _GotoState46Action,
	{_State116, AnonymousFuncExprType}:            _GotoState41Action,
	{_State123, IntegerLiteralToken}:              _GotoState26Action,
	{_State123, FloatLiteralToken}:                _GotoState23Action,
	{_State123, RuneLiteralToken}:                 _GotoState33Action,
	{_State123, StringLiteralToken}:               _GotoState34Action,
	{_State123, IdentifierToken}:                  _GotoState25Action,
	{_State123, TrueToken}:                        _GotoState37Action,
	{_State123, FalseToken}:                       _GotoState22Action,
	{_State123, StructToken}:                      _GotoState35Action,
	{_State123, EnumToken}:                        _GotoState21Action,
	{_State123, FuncToken}:                        _GotoState24Action,
	{_State123, LabelDeclToken}:                   _GotoState27Action,
	{_State123, LparenToken}:                      _GotoState30Action,
	{_State123, LbracketToken}:                    _GotoState28Action,
	{_State123, NotToken}:                         _GotoState32Action,
	{_State123, SubToken}:                         _GotoState36Action,
	{_State123, MulToken}:                         _GotoState31Action,
	{_State123, BitNegToken}:                      _GotoState20Action,
	{_State123, BitAndToken}:                      _GotoState19Action,
	{_State123, LexErrorToken}:                    _GotoState29Action,
	{_State123, OptionalLabelDeclType}:            _GotoState134Action,
	{_State123, BlockExprType}:                    _GotoState43Action,
	{_State123, CallExprType}:                     _GotoState44Action,
	{_State123, AtomExprType}:                     _GotoState42Action,
	{_State123, LiteralType}:                      _GotoState50Action,
	{_State123, InitializeExprType}:               _GotoState49Action,
	{_State123, AccessExprType}:                   _GotoState38Action,
	{_State123, PostfixUnaryExprType}:             _GotoState54Action,
	{_State123, PrefixUnaryOpType}:                _GotoState56Action,
	{_State123, PrefixUnaryExprType}:              _GotoState192Action,
	{_State123, InitializableTypeType}:            _GotoState48Action,
	{_State123, ExplicitStructDefType}:            _GotoState47Action,
	{_State123, ExplicitEnumDefType}:              _GotoState46Action,
	{_State123, AnonymousFuncExprType}:            _GotoState41Action,
	{_State124, LbraceToken}:                      _GotoState127Action,
	{_State124, BlockBodyType}:                    _GotoState193Action,
	{_State125, IntegerLiteralToken}:              _GotoState26Action,
	{_State125, FloatLiteralToken}:                _GotoState23Action,
	{_State125, RuneLiteralToken}:                 _GotoState33Action,
	{_State125, StringLiteralToken}:               _GotoState34Action,
	{_State125, IdentifierToken}:                  _GotoState25Action,
	{_State125, TrueToken}:                        _GotoState37Action,
	{_State125, FalseToken}:                       _GotoState22Action,
	{_State125, InToken}:                          _GotoState194Action,
	{_State125, StructToken}:                      _GotoState35Action,
	{_State125, EnumToken}:                        _GotoState21Action,
	{_State125, FuncToken}:                        _GotoState24Action,
	{_State125, LabelDeclToken}:                   _GotoState27Action,
	{_State125, LparenToken}:                      _GotoState30Action,
	{_State125, LbracketToken}:                    _GotoState28Action,
	{_State125, SemicolonToken}:                   _GotoState195Action,
	{_State125, NotToken}:                         _GotoState32Action,
	{_State125, SubToken}:                         _GotoState36Action,
	{_State125, MulToken}:                         _GotoState31Action,
	{_State125, BitNegToken}:                      _GotoState20Action,
	{_State125, BitAndToken}:                      _GotoState19Action,
	{_State125, LexErrorToken}:                    _GotoState29Action,
	{_State125, OptionalLabelDeclType}:            _GotoState134Action,
	{_State125, SequenceExprType}:                 _GotoState196Action,
	{_State125, BlockExprType}:                    _GotoState43Action,
	{_State125, CallExprType}:                     _GotoState44Action,
	{_State125, AtomExprType}:                     _GotoState42Action,
	{_State125, LiteralType}:                      _GotoState50Action,
	{_State125, InitializeExprType}:               _GotoState49Action,
	{_State125, AccessExprType}:                   _GotoState38Action,
	{_State125, PostfixUnaryExprType}:             _GotoState54Action,
	{_State125, PrefixUnaryOpType}:                _GotoState56Action,
	{_State125, PrefixUnaryExprType}:              _GotoState55Action,
	{_State125, MulExprType}:                      _GotoState51Action,
	{_State125, AddExprType}:                      _GotoState39Action,
	{_State125, CmpExprType}:                      _GotoState45Action,
	{_State125, AndExprType}:                      _GotoState40Action,
	{_State125, OrExprType}:                       _GotoState53Action,
	{_State125, InitializableTypeType}:            _GotoState48Action,
	{_State125, ExplicitStructDefType}:            _GotoState47Action,
	{_State125, ExplicitEnumDefType}:              _GotoState46Action,
	{_State125, AnonymousFuncExprType}:            _GotoState41Action,
	{_State126, IntegerLiteralToken}:              _GotoState26Action,
	{_State126, FloatLiteralToken}:                _GotoState23Action,
	{_State126, RuneLiteralToken}:                 _GotoState33Action,
	{_State126, StringLiteralToken}:               _GotoState34Action,
	{_State126, IdentifierToken}:                  _GotoState25Action,
	{_State126, TrueToken}:                        _GotoState37Action,
	{_State126, FalseToken}:                       _GotoState22Action,
	{_State126, StructToken}:                      _GotoState35Action,
	{_State126, EnumToken}:                        _GotoState21Action,
	{_State126, FuncToken}:                        _GotoState24Action,
	{_State126, LabelDeclToken}:                   _GotoState27Action,
	{_State126, LparenToken}:                      _GotoState30Action,
	{_State126, LbracketToken}:                    _GotoState28Action,
	{_State126, NotToken}:                         _GotoState32Action,
	{_State126, SubToken}:                         _GotoState36Action,
	{_State126, MulToken}:                         _GotoState31Action,
	{_State126, BitNegToken}:                      _GotoState20Action,
	{_State126, BitAndToken}:                      _GotoState19Action,
	{_State126, LexErrorToken}:                    _GotoState29Action,
	{_State126, OptionalLabelDeclType}:            _GotoState134Action,
	{_State126, SequenceExprType}:                 _GotoState197Action,
	{_State126, BlockExprType}:                    _GotoState43Action,
	{_State126, CallExprType}:                     _GotoState44Action,
	{_State126, AtomExprType}:                     _GotoState42Action,
	{_State126, LiteralType}:                      _GotoState50Action,
	{_State126, InitializeExprType}:               _GotoState49Action,
	{_State126, AccessExprType}:                   _GotoState38Action,
	{_State126, PostfixUnaryExprType}:             _GotoState54Action,
	{_State126, PrefixUnaryOpType}:                _GotoState56Action,
	{_State126, PrefixUnaryExprType}:              _GotoState55Action,
	{_State126, MulExprType}:                      _GotoState51Action,
	{_State126, AddExprType}:                      _GotoState39Action,
	{_State126, CmpExprType}:                      _GotoState45Action,
	{_State126, AndExprType}:                      _GotoState40Action,
	{_State126, OrExprType}:                       _GotoState53Action,
	{_State126, InitializableTypeType}:            _GotoState48Action,
	{_State126, ExplicitStructDefType}:            _GotoState47Action,
	{_State126, ExplicitEnumDefType}:              _GotoState46Action,
	{_State126, AnonymousFuncExprType}:            _GotoState41Action,
	{_State127, StatementsType}:                   _GotoState198Action,
	{_State128, IntegerLiteralToken}:              _GotoState26Action,
	{_State128, FloatLiteralToken}:                _GotoState23Action,
	{_State128, RuneLiteralToken}:                 _GotoState33Action,
	{_State128, StringLiteralToken}:               _GotoState34Action,
	{_State128, IdentifierToken}:                  _GotoState25Action,
	{_State128, TrueToken}:                        _GotoState37Action,
	{_State128, FalseToken}:                       _GotoState22Action,
	{_State128, StructToken}:                      _GotoState35Action,
	{_State128, EnumToken}:                        _GotoState21Action,
	{_State128, FuncToken}:                        _GotoState24Action,
	{_State128, LabelDeclToken}:                   _GotoState27Action,
	{_State128, LparenToken}:                      _GotoState30Action,
	{_State128, LbracketToken}:                    _GotoState28Action,
	{_State128, NotToken}:                         _GotoState32Action,
	{_State128, SubToken}:                         _GotoState36Action,
	{_State128, MulToken}:                         _GotoState31Action,
	{_State128, BitNegToken}:                      _GotoState20Action,
	{_State128, BitAndToken}:                      _GotoState19Action,
	{_State128, LexErrorToken}:                    _GotoState29Action,
	{_State128, OptionalLabelDeclType}:            _GotoState134Action,
	{_State128, SequenceExprType}:                 _GotoState199Action,
	{_State128, BlockExprType}:                    _GotoState43Action,
	{_State128, CallExprType}:                     _GotoState44Action,
	{_State128, AtomExprType}:                     _GotoState42Action,
	{_State128, LiteralType}:                      _GotoState50Action,
	{_State128, InitializeExprType}:               _GotoState49Action,
	{_State128, AccessExprType}:                   _GotoState38Action,
	{_State128, PostfixUnaryExprType}:             _GotoState54Action,
	{_State128, PrefixUnaryOpType}:                _GotoState56Action,
	{_State128, PrefixUnaryExprType}:              _GotoState55Action,
	{_State128, MulExprType}:                      _GotoState51Action,
	{_State128, AddExprType}:                      _GotoState39Action,
	{_State128, CmpExprType}:                      _GotoState45Action,
	{_State128, AndExprType}:                      _GotoState40Action,
	{_State128, OrExprType}:                       _GotoState53Action,
	{_State128, InitializableTypeType}:            _GotoState48Action,
	{_State128, ExplicitStructDefType}:            _GotoState47Action,
	{_State128, ExplicitEnumDefType}:              _GotoState46Action,
	{_State128, AnonymousFuncExprType}:            _GotoState41Action,
	{_State133, IntegerLiteralToken}:              _GotoState26Action,
	{_State133, FloatLiteralToken}:                _GotoState23Action,
	{_State133, RuneLiteralToken}:                 _GotoState33Action,
	{_State133, StringLiteralToken}:               _GotoState34Action,
	{_State133, IdentifierToken}:                  _GotoState25Action,
	{_State133, TrueToken}:                        _GotoState37Action,
	{_State133, FalseToken}:                       _GotoState22Action,
	{_State133, StructToken}:                      _GotoState35Action,
	{_State133, EnumToken}:                        _GotoState21Action,
	{_State133, FuncToken}:                        _GotoState24Action,
	{_State133, LabelDeclToken}:                   _GotoState27Action,
	{_State133, LparenToken}:                      _GotoState30Action,
	{_State133, LbracketToken}:                    _GotoState28Action,
	{_State133, NotToken}:                         _GotoState32Action,
	{_State133, SubToken}:                         _GotoState36Action,
	{_State133, MulToken}:                         _GotoState31Action,
	{_State133, BitNegToken}:                      _GotoState20Action,
	{_State133, BitAndToken}:                      _GotoState19Action,
	{_State133, LexErrorToken}:                    _GotoState29Action,
	{_State133, OptionalLabelDeclType}:            _GotoState134Action,
	{_State133, BlockExprType}:                    _GotoState43Action,
	{_State133, CallExprType}:                     _GotoState44Action,
	{_State133, AtomExprType}:                     _GotoState42Action,
	{_State133, LiteralType}:                      _GotoState50Action,
	{_State133, InitializeExprType}:               _GotoState49Action,
	{_State133, AccessExprType}:                   _GotoState38Action,
	{_State133, PostfixUnaryExprType}:             _GotoState54Action,
	{_State133, PrefixUnaryOpType}:                _GotoState56Action,
	{_State133, PrefixUnaryExprType}:              _GotoState55Action,
	{_State133, MulExprType}:                      _GotoState51Action,
	{_State133, AddExprType}:                      _GotoState39Action,
	{_State133, CmpExprType}:                      _GotoState45Action,
	{_State133, AndExprType}:                      _GotoState200Action,
	{_State133, InitializableTypeType}:            _GotoState48Action,
	{_State133, ExplicitStructDefType}:            _GotoState47Action,
	{_State133, ExplicitEnumDefType}:              _GotoState46Action,
	{_State133, AnonymousFuncExprType}:            _GotoState41Action,
	{_State134, LbraceToken}:                      _GotoState127Action,
	{_State134, BlockBodyType}:                    _GotoState129Action,
	{_State136, IdentifierToken}:                  _GotoState201Action,
	{_State137, PackageToken}:                     _GotoState16Action,
	{_State137, UnsafeToken}:                      _GotoState60Action,
	{_State137, TypeToken}:                        _GotoState17Action,
	{_State137, FuncToken}:                        _GotoState18Action,
	{_State137, DefinitionType}:                   _GotoState202Action,
	{_State137, UnsafeStatementType}:              _GotoState66Action,
	{_State137, TypeDefType}:                      _GotoState65Action,
	{_State137, NamedFuncDefType}:                 _GotoState63Action,
	{_State137, PackageDefType}:                   _GotoState64Action,
	{_State139, PackageStatementsType}:            _GotoState203Action,
	{_State140, IdentifierToken}:                  _GotoState204Action,
	{_State140, GenericParameterDefType}:          _GotoState205Action,
	{_State140, GenericParameterDefsType}:         _GotoState206Action,
	{_State140, OptionalGenericParameterDefsType}: _GotoState207Action,
	{_State141, IdentifierToken}:                  _GotoState78Action,
	{_State141, StructToken}:                      _GotoState35Action,
	{_State141, EnumToken}:                        _GotoState21Action,
	{_State141, TraitToken}:                       _GotoState83Action,
	{_State141, FuncToken}:                        _GotoState77Action,
	{_State141, LparenToken}:                      _GotoState80Action,
	{_State141, LbracketToken}:                    _GotoState28Action,
	{_State141, DotToken}:                         _GotoState75Action,
	{_State141, QuestionToken}:                    _GotoState81Action,
	{_State141, ExclaimToken}:                     _GotoState76Action,
	{_State141, TildeTildeToken}:                  _GotoState82Action,
	{_State141, BitNegToken}:                      _GotoState74Action,
	{_State141, BitAndToken}:                      _GotoState73Action,
	{_State141, LexErrorToken}:                    _GotoState79Action,
	{_State141, InitializableTypeType}:            _GotoState88Action,
	{_State141, AtomTypeType}:                     _GotoState84Action,
	{_State141, ValueTypeType}:                    _GotoState209Action,
	{_State141, ImplicitStructDefType}:            _GotoState87Action,
	{_State141, ExplicitStructDefType}:            _GotoState47Action,
	{_State141, ImplicitEnumDefType}:              _GotoState86Action,
	{_State141, ExplicitEnumDefType}:              _GotoState46Action,
	{_State141, TraitAlgebraTypeType}:             _GotoState208Action,
	{_State141, TraitDefType}:                     _GotoState89Action,
	{_State141, FuncTypeType}:                     _GotoState85Action,
	{_State142, IdentifierToken}:                  _GotoState78Action,
	{_State142, StructToken}:                      _GotoState35Action,
	{_State142, EnumToken}:                        _GotoState21Action,
	{_State142, TraitToken}:                       _GotoState83Action,
	{_State142, FuncToken}:                        _GotoState77Action,
	{_State142, LparenToken}:                      _GotoState80Action,
	{_State142, LbracketToken}:                    _GotoState28Action,
	{_State142, DotToken}:                         _GotoState75Action,
	{_State142, QuestionToken}:                    _GotoState81Action,
	{_State142, ExclaimToken}:                     _GotoState76Action,
	{_State142, TildeTildeToken}:                  _GotoState82Action,
	{_State142, BitNegToken}:                      _GotoState74Action,
	{_State142, BitAndToken}:                      _GotoState73Action,
	{_State142, LexErrorToken}:                    _GotoState79Action,
	{_State142, InitializableTypeType}:            _GotoState88Action,
	{_State142, AtomTypeType}:                     _GotoState84Action,
	{_State142, ValueTypeType}:                    _GotoState210Action,
	{_State142, ImplicitStructDefType}:            _GotoState87Action,
	{_State142, ExplicitStructDefType}:            _GotoState47Action,
	{_State142, ImplicitEnumDefType}:              _GotoState86Action,
	{_State142, ExplicitEnumDefType}:              _GotoState46Action,
	{_State142, TraitDefType}:                     _GotoState89Action,
	{_State142, FuncTypeType}:                     _GotoState85Action,
	{_State143, IdentifierToken}:                  _GotoState78Action,
	{_State143, StructToken}:                      _GotoState35Action,
	{_State143, EnumToken}:                        _GotoState21Action,
	{_State143, TraitToken}:                       _GotoState83Action,
	{_State143, FuncToken}:                        _GotoState77Action,
	{_State143, LparenToken}:                      _GotoState80Action,
	{_State143, LbracketToken}:                    _GotoState28Action,
	{_State143, DotToken}:                         _GotoState75Action,
	{_State143, QuestionToken}:                    _GotoState81Action,
	{_State143, ExclaimToken}:                     _GotoState76Action,
	{_State143, DotdotdotToken}:                   _GotoState211Action,
	{_State143, TildeTildeToken}:                  _GotoState82Action,
	{_State143, BitNegToken}:                      _GotoState74Action,
	{_State143, BitAndToken}:                      _GotoState73Action,
	{_State143, LexErrorToken}:                    _GotoState79Action,
	{_State143, InitializableTypeType}:            _GotoState88Action,
	{_State143, AtomTypeType}:                     _GotoState84Action,
	{_State143, ValueTypeType}:                    _GotoState212Action,
	{_State143, ImplicitStructDefType}:            _GotoState87Action,
	{_State143, ExplicitStructDefType}:            _GotoState47Action,
	{_State143, ImplicitEnumDefType}:              _GotoState86Action,
	{_State143, ExplicitEnumDefType}:              _GotoState46Action,
	{_State143, TraitDefType}:                     _GotoState89Action,
	{_State143, FuncTypeType}:                     _GotoState85Action,
	{_State144, RparenToken}:                      _GotoState213Action,
	{_State145, DollarLbracketToken}:              _GotoState140Action,
	{_State145, OptionalGenericParametersType}:    _GotoState214Action,
	{_State146, IdentifierToken}:                  _GotoState78Action,
	{_State146, StructToken}:                      _GotoState35Action,
	{_State146, EnumToken}:                        _GotoState21Action,
	{_State146, TraitToken}:                       _GotoState83Action,
	{_State146, FuncToken}:                        _GotoState77Action,
	{_State146, LparenToken}:                      _GotoState80Action,
	{_State146, LbracketToken}:                    _GotoState28Action,
	{_State146, DotToken}:                         _GotoState215Action,
	{_State146, QuestionToken}:                    _GotoState81Action,
	{_State146, ExclaimToken}:                     _GotoState76Action,
	{_State146, DollarLbracketToken}:              _GotoState98Action,
	{_State146, TildeTildeToken}:                  _GotoState82Action,
	{_State146, BitNegToken}:                      _GotoState74Action,
	{_State146, BitAndToken}:                      _GotoState73Action,
	{_State146, LexErrorToken}:                    _GotoState79Action,
	{_State146, OptionalGenericBindingType}:       _GotoState162Action,
	{_State146, InitializableTypeType}:            _GotoState88Action,
	{_State146, AtomTypeType}:                     _GotoState84Action,
	{_State146, ValueTypeType}:                    _GotoState216Action,
	{_State146, ImplicitStructDefType}:            _GotoState87Action,
	{_State146, ExplicitStructDefType}:            _GotoState47Action,
	{_State146, ImplicitEnumDefType}:              _GotoState86Action,
	{_State146, ExplicitEnumDefType}:              _GotoState46Action,
	{_State146, TraitDefType}:                     _GotoState89Action,
	{_State146, FuncTypeType}:                     _GotoState85Action,
	{_State147, NewlinesToken}:                    _GotoState217Action,
	{_State147, OrToken}:                          _GotoState218Action,
	{_State148, RparenToken}:                      _GotoState219Action,
	{_State149, AssignToken}:                      _GotoState220Action,
	{_State150, NewlinesToken}:                    _GotoState221Action,
	{_State150, OrToken}:                          _GotoState222Action,
	{_State153, RparenToken}:                      _GotoState223Action,
	{_State155, CommaToken}:                       _GotoState224Action,
	{_State160, IdentifierToken}:                  _GotoState226Action,
	{_State160, StructToken}:                      _GotoState35Action,
	{_State160, EnumToken}:                        _GotoState21Action,
	{_State160, TraitToken}:                       _GotoState83Action,
	{_State160, FuncToken}:                        _GotoState77Action,
	{_State160, LparenToken}:                      _GotoState80Action,
	{_State160, LbracketToken}:                    _GotoState28Action,
	{_State160, DotToken}:                         _GotoState75Action,
	{_State160, QuestionToken}:                    _GotoState81Action,
	{_State160, ExclaimToken}:                     _GotoState76Action,
	{_State160, DotdotdotToken}:                   _GotoState225Action,
	{_State160, TildeTildeToken}:                  _GotoState82Action,
	{_State160, BitNegToken}:                      _GotoState74Action,
	{_State160, BitAndToken}:                      _GotoState73Action,
	{_State160, LexErrorToken}:                    _GotoState79Action,
	{_State160, InitializableTypeType}:            _GotoState88Action,
	{_State160, AtomTypeType}:                     _GotoState84Action,
	{_State160, ValueTypeType}:                    _GotoState230Action,
	{_State160, ImplicitStructDefType}:            _GotoState87Action,
	{_State160, ExplicitStructDefType}:            _GotoState47Action,
	{_State160, ImplicitEnumDefType}:              _GotoState86Action,
	{_State160, ExplicitEnumDefType}:              _GotoState46Action,
	{_State160, TraitDefType}:                     _GotoState89Action,
	{_State160, ParameterDeclType}:                _GotoState228Action,
	{_State160, ParameterDeclsType}:               _GotoState229Action,
	{_State160, OptionalParameterDeclsType}:       _GotoState227Action,
	{_State160, FuncTypeType}:                     _GotoState85Action,
	{_State161, IdentifierToken}:                  _GotoState231Action,
	{_State163, OrToken}:                          _GotoState232Action,
	{_State164, AssignToken}:                      _GotoState220Action,
	{_State165, RparenToken}:                      _GotoState234Action,
	{_State165, OrToken}:                          _GotoState233Action,
	{_State166, CommaToken}:                       _GotoState235Action,
	{_State167, RparenToken}:                      _GotoState236Action,
	{_State170, IdentifierToken}:                  _GotoState238Action,
	{_State170, UnsafeToken}:                      _GotoState60Action,
	{_State170, StructToken}:                      _GotoState35Action,
	{_State170, EnumToken}:                        _GotoState21Action,
	{_State170, TraitToken}:                       _GotoState83Action,
	{_State170, FuncToken}:                        _GotoState237Action,
	{_State170, LparenToken}:                      _GotoState80Action,
	{_State170, LbracketToken}:                    _GotoState28Action,
	{_State170, DotToken}:                         _GotoState75Action,
	{_State170, QuestionToken}:                    _GotoState81Action,
	{_State170, ExclaimToken}:                     _GotoState76Action,
	{_State170, TildeTildeToken}:                  _GotoState82Action,
	{_State170, BitNegToken}:                      _GotoState74Action,
	{_State170, BitAndToken}:                      _GotoState73Action,
	{_State170, LexErrorToken}:                    _GotoState79Action,
	{_State170, UnsafeStatementType}:              _GotoState244Action,
	{_State170, InitializableTypeType}:            _GotoState88Action,
	{_State170, AtomTypeType}:                     _GotoState84Action,
	{_State170, ValueTypeType}:                    _GotoState209Action,
	{_State170, ImplicitStructDefType}:            _GotoState87Action,
	{_State170, ExplicitStructDefType}:            _GotoState47Action,
	{_State170, ImplicitEnumDefType}:              _GotoState86Action,
	{_State170, ExplicitEnumDefType}:              _GotoState46Action,
	{_State170, TraitAlgebraTypeType}:             _GotoState241Action,
	{_State170, TraitPropertyType}:                _GotoState243Action,
	{_State170, TraitPropertiesType}:              _GotoState242Action,
	{_State170, OptionalTraitPropertiesType}:      _GotoState240Action,
	{_State170, TraitDefType}:                     _GotoState89Action,
	{_State170, FuncTypeType}:                     _GotoState85Action,
	{_State170, MethodSignatureType}:              _GotoState239Action,
	{_State171, IdentifierToken}:                  _GotoState78Action,
	{_State171, StructToken}:                      _GotoState35Action,
	{_State171, EnumToken}:                        _GotoState21Action,
	{_State171, TraitToken}:                       _GotoState83Action,
	{_State171, FuncToken}:                        _GotoState77Action,
	{_State171, LparenToken}:                      _GotoState80Action,
	{_State171, LbracketToken}:                    _GotoState28Action,
	{_State171, DotToken}:                         _GotoState75Action,
	{_State171, QuestionToken}:                    _GotoState81Action,
	{_State171, ExclaimToken}:                     _GotoState76Action,
	{_State171, TildeTildeToken}:                  _GotoState82Action,
	{_State171, BitNegToken}:                      _GotoState74Action,
	{_State171, BitAndToken}:                      _GotoState73Action,
	{_State171, LexErrorToken}:                    _GotoState79Action,
	{_State171, InitializableTypeType}:            _GotoState88Action,
	{_State171, AtomTypeType}:                     _GotoState84Action,
	{_State171, ValueTypeType}:                    _GotoState245Action,
	{_State171, ImplicitStructDefType}:            _GotoState87Action,
	{_State171, ExplicitStructDefType}:            _GotoState47Action,
	{_State171, ImplicitEnumDefType}:              _GotoState86Action,
	{_State171, ExplicitEnumDefType}:              _GotoState46Action,
	{_State171, TraitDefType}:                     _GotoState89Action,
	{_State171, FuncTypeType}:                     _GotoState85Action,
	{_State172, IntegerLiteralToken}:              _GotoState246Action,
	{_State174, IntegerLiteralToken}:              _GotoState26Action,
	{_State174, FloatLiteralToken}:                _GotoState23Action,
	{_State174, RuneLiteralToken}:                 _GotoState33Action,
	{_State174, StringLiteralToken}:               _GotoState34Action,
	{_State174, IdentifierToken}:                  _GotoState25Action,
	{_State174, TrueToken}:                        _GotoState37Action,
	{_State174, FalseToken}:                       _GotoState22Action,
	{_State174, StructToken}:                      _GotoState35Action,
	{_State174, EnumToken}:                        _GotoState21Action,
	{_State174, FuncToken}:                        _GotoState24Action,
	{_State174, LabelDeclToken}:                   _GotoState27Action,
	{_State174, LparenToken}:                      _GotoState30Action,
	{_State174, LbracketToken}:                    _GotoState28Action,
	{_State174, NotToken}:                         _GotoState32Action,
	{_State174, SubToken}:                         _GotoState36Action,
	{_State174, MulToken}:                         _GotoState31Action,
	{_State174, BitNegToken}:                      _GotoState20Action,
	{_State174, BitAndToken}:                      _GotoState19Action,
	{_State174, LexErrorToken}:                    _GotoState29Action,
	{_State174, ExpressionType}:                   _GotoState247Action,
	{_State174, OptionalLabelDeclType}:            _GotoState52Action,
	{_State174, SequenceExprType}:                 _GotoState57Action,
	{_State174, BlockExprType}:                    _GotoState43Action,
	{_State174, CallExprType}:                     _GotoState44Action,
	{_State174, AtomExprType}:                     _GotoState42Action,
	{_State174, LiteralType}:                      _GotoState50Action,
	{_State174, InitializeExprType}:               _GotoState49Action,
	{_State174, AccessExprType}:                   _GotoState38Action,
	{_State174, PostfixUnaryExprType}:             _GotoState54Action,
	{_State174, PrefixUnaryOpType}:                _GotoState56Action,
	{_State174, PrefixUnaryExprType}:              _GotoState55Action,
	{_State174, MulExprType}:                      _GotoState51Action,
	{_State174, AddExprType}:                      _GotoState39Action,
	{_State174, CmpExprType}:                      _GotoState45Action,
	{_State174, AndExprType}:                      _GotoState40Action,
	{_State174, OrExprType}:                       _GotoState53Action,
	{_State174, InitializableTypeType}:            _GotoState48Action,
	{_State174, ExplicitStructDefType}:            _GotoState47Action,
	{_State174, ExplicitEnumDefType}:              _GotoState46Action,
	{_State174, AnonymousFuncExprType}:            _GotoState41Action,
	{_State175, IntegerLiteralToken}:              _GotoState26Action,
	{_State175, FloatLiteralToken}:                _GotoState23Action,
	{_State175, RuneLiteralToken}:                 _GotoState33Action,
	{_State175, StringLiteralToken}:               _GotoState34Action,
	{_State175, IdentifierToken}:                  _GotoState91Action,
	{_State175, TrueToken}:                        _GotoState37Action,
	{_State175, FalseToken}:                       _GotoState22Action,
	{_State175, StructToken}:                      _GotoState35Action,
	{_State175, EnumToken}:                        _GotoState21Action,
	{_State175, FuncToken}:                        _GotoState24Action,
	{_State175, LabelDeclToken}:                   _GotoState27Action,
	{_State175, LparenToken}:                      _GotoState30Action,
	{_State175, LbracketToken}:                    _GotoState28Action,
	{_State175, NotToken}:                         _GotoState32Action,
	{_State175, SubToken}:                         _GotoState36Action,
	{_State175, MulToken}:                         _GotoState31Action,
	{_State175, BitNegToken}:                      _GotoState20Action,
	{_State175, BitAndToken}:                      _GotoState19Action,
	{_State175, LexErrorToken}:                    _GotoState29Action,
	{_State175, ExpressionType}:                   _GotoState95Action,
	{_State175, OptionalLabelDeclType}:            _GotoState52Action,
	{_State175, SequenceExprType}:                 _GotoState57Action,
	{_State175, BlockExprType}:                    _GotoState43Action,
	{_State175, CallExprType}:                     _GotoState44Action,
	{_State175, ArgumentType}:                     _GotoState248Action,
	{_State175, ColonExpressionsType}:             _GotoState94Action,
	{_State175, OptionalExpressionType}:           _GotoState96Action,
	{_State175, AtomExprType}:                     _GotoState42Action,
	{_State175, LiteralType}:                      _GotoState50Action,
	{_State175, InitializeExprType}:               _GotoState49Action,
	{_State175, AccessExprType}:                   _GotoState38Action,
	{_State175, PostfixUnaryExprType}:             _GotoState54Action,
	{_State175, PrefixUnaryOpType}:                _GotoState56Action,
	{_State175, PrefixUnaryExprType}:              _GotoState55Action,
	{_State175, MulExprType}:                      _GotoState51Action,
	{_State175, AddExprType}:                      _GotoState39Action,
	{_State175, CmpExprType}:                      _GotoState45Action,
	{_State175, AndExprType}:                      _GotoState40Action,
	{_State175, OrExprType}:                       _GotoState53Action,
	{_State175, InitializableTypeType}:            _GotoState48Action,
	{_State175, ExplicitStructDefType}:            _GotoState47Action,
	{_State175, ExplicitEnumDefType}:              _GotoState46Action,
	{_State175, AnonymousFuncExprType}:            _GotoState41Action,
	{_State177, IntegerLiteralToken}:              _GotoState26Action,
	{_State177, FloatLiteralToken}:                _GotoState23Action,
	{_State177, RuneLiteralToken}:                 _GotoState33Action,
	{_State177, StringLiteralToken}:               _GotoState34Action,
	{_State177, IdentifierToken}:                  _GotoState25Action,
	{_State177, TrueToken}:                        _GotoState37Action,
	{_State177, FalseToken}:                       _GotoState22Action,
	{_State177, StructToken}:                      _GotoState35Action,
	{_State177, EnumToken}:                        _GotoState21Action,
	{_State177, FuncToken}:                        _GotoState24Action,
	{_State177, LabelDeclToken}:                   _GotoState27Action,
	{_State177, LparenToken}:                      _GotoState30Action,
	{_State177, LbracketToken}:                    _GotoState28Action,
	{_State177, NotToken}:                         _GotoState32Action,
	{_State177, SubToken}:                         _GotoState36Action,
	{_State177, MulToken}:                         _GotoState31Action,
	{_State177, BitNegToken}:                      _GotoState20Action,
	{_State177, BitAndToken}:                      _GotoState19Action,
	{_State177, LexErrorToken}:                    _GotoState29Action,
	{_State177, ExpressionType}:                   _GotoState249Action,
	{_State177, OptionalLabelDeclType}:            _GotoState52Action,
	{_State177, SequenceExprType}:                 _GotoState57Action,
	{_State177, BlockExprType}:                    _GotoState43Action,
	{_State177, CallExprType}:                     _GotoState44Action,
	{_State177, OptionalExpressionType}:           _GotoState250Action,
	{_State177, AtomExprType}:                     _GotoState42Action,
	{_State177, LiteralType}:                      _GotoState50Action,
	{_State177, InitializeExprType}:               _GotoState49Action,
	{_State177, AccessExprType}:                   _GotoState38Action,
	{_State177, PostfixUnaryExprType}:             _GotoState54Action,
	{_State177, PrefixUnaryOpType}:                _GotoState56Action,
	{_State177, PrefixUnaryExprType}:              _GotoState55Action,
	{_State177, MulExprType}:                      _GotoState51Action,
	{_State177, AddExprType}:                      _GotoState39Action,
	{_State177, CmpExprType}:                      _GotoState45Action,
	{_State177, AndExprType}:                      _GotoState40Action,
	{_State177, OrExprType}:                       _GotoState53Action,
	{_State177, InitializableTypeType}:            _GotoState48Action,
	{_State177, ExplicitStructDefType}:            _GotoState47Action,
	{_State177, ExplicitEnumDefType}:              _GotoState46Action,
	{_State177, AnonymousFuncExprType}:            _GotoState41Action,
	{_State178, IntegerLiteralToken}:              _GotoState26Action,
	{_State178, FloatLiteralToken}:                _GotoState23Action,
	{_State178, RuneLiteralToken}:                 _GotoState33Action,
	{_State178, StringLiteralToken}:               _GotoState34Action,
	{_State178, IdentifierToken}:                  _GotoState25Action,
	{_State178, TrueToken}:                        _GotoState37Action,
	{_State178, FalseToken}:                       _GotoState22Action,
	{_State178, StructToken}:                      _GotoState35Action,
	{_State178, EnumToken}:                        _GotoState21Action,
	{_State178, FuncToken}:                        _GotoState24Action,
	{_State178, LabelDeclToken}:                   _GotoState27Action,
	{_State178, LparenToken}:                      _GotoState30Action,
	{_State178, LbracketToken}:                    _GotoState28Action,
	{_State178, NotToken}:                         _GotoState32Action,
	{_State178, SubToken}:                         _GotoState36Action,
	{_State178, MulToken}:                         _GotoState31Action,
	{_State178, BitNegToken}:                      _GotoState20Action,
	{_State178, BitAndToken}:                      _GotoState19Action,
	{_State178, LexErrorToken}:                    _GotoState29Action,
	{_State178, ExpressionType}:                   _GotoState249Action,
	{_State178, OptionalLabelDeclType}:            _GotoState52Action,
	{_State178, SequenceExprType}:                 _GotoState57Action,
	{_State178, BlockExprType}:                    _GotoState43Action,
	{_State178, CallExprType}:                     _GotoState44Action,
	{_State178, OptionalExpressionType}:           _GotoState251Action,
	{_State178, AtomExprType}:                     _GotoState42Action,
	{_State178, LiteralType}:                      _GotoState50Action,
	{_State178, InitializeExprType}:               _GotoState49Action,
	{_State178, AccessExprType}:                   _GotoState38Action,
	{_State178, PostfixUnaryExprType}:             _GotoState54Action,
	{_State178, PrefixUnaryOpType}:                _GotoState56Action,
	{_State178, PrefixUnaryExprType}:              _GotoState55Action,
	{_State178, MulExprType}:                      _GotoState51Action,
	{_State178, AddExprType}:                      _GotoState39Action,
	{_State178, CmpExprType}:                      _GotoState45Action,
	{_State178, AndExprType}:                      _GotoState40Action,
	{_State178, OrExprType}:                       _GotoState53Action,
	{_State178, InitializableTypeType}:            _GotoState48Action,
	{_State178, ExplicitStructDefType}:            _GotoState47Action,
	{_State178, ExplicitEnumDefType}:              _GotoState46Action,
	{_State178, AnonymousFuncExprType}:            _GotoState41Action,
	{_State179, NewlinesToken}:                    _GotoState253Action,
	{_State179, CommaToken}:                       _GotoState252Action,
	{_State181, RparenToken}:                      _GotoState254Action,
	{_State182, CommaToken}:                       _GotoState255Action,
	{_State183, RbracketToken}:                    _GotoState256Action,
	{_State186, RbracketToken}:                    _GotoState257Action,
	{_State187, IntegerLiteralToken}:              _GotoState26Action,
	{_State187, FloatLiteralToken}:                _GotoState23Action,
	{_State187, RuneLiteralToken}:                 _GotoState33Action,
	{_State187, StringLiteralToken}:               _GotoState34Action,
	{_State187, IdentifierToken}:                  _GotoState91Action,
	{_State187, TrueToken}:                        _GotoState37Action,
	{_State187, FalseToken}:                       _GotoState22Action,
	{_State187, StructToken}:                      _GotoState35Action,
	{_State187, EnumToken}:                        _GotoState21Action,
	{_State187, FuncToken}:                        _GotoState24Action,
	{_State187, LabelDeclToken}:                   _GotoState27Action,
	{_State187, LparenToken}:                      _GotoState30Action,
	{_State187, LbracketToken}:                    _GotoState28Action,
	{_State187, NotToken}:                         _GotoState32Action,
	{_State187, SubToken}:                         _GotoState36Action,
	{_State187, MulToken}:                         _GotoState31Action,
	{_State187, BitNegToken}:                      _GotoState20Action,
	{_State187, BitAndToken}:                      _GotoState19Action,
	{_State187, LexErrorToken}:                    _GotoState29Action,
	{_State187, ExpressionType}:                   _GotoState95Action,
	{_State187, OptionalLabelDeclType}:            _GotoState52Action,
	{_State187, SequenceExprType}:                 _GotoState57Action,
	{_State187, BlockExprType}:                    _GotoState43Action,
	{_State187, CallExprType}:                     _GotoState44Action,
	{_State187, OptionalArgumentsType}:            _GotoState259Action,
	{_State187, ArgumentsType}:                    _GotoState258Action,
	{_State187, ArgumentType}:                     _GotoState92Action,
	{_State187, ColonExpressionsType}:             _GotoState94Action,
	{_State187, OptionalExpressionType}:           _GotoState96Action,
	{_State187, AtomExprType}:                     _GotoState42Action,
	{_State187, LiteralType}:                      _GotoState50Action,
	{_State187, InitializeExprType}:               _GotoState49Action,
	{_State187, AccessExprType}:                   _GotoState38Action,
	{_State187, PostfixUnaryExprType}:             _GotoState54Action,
	{_State187, PrefixUnaryOpType}:                _GotoState56Action,
	{_State187, PrefixUnaryExprType}:              _GotoState55Action,
	{_State187, MulExprType}:                      _GotoState51Action,
	{_State187, AddExprType}:                      _GotoState39Action,
	{_State187, CmpExprType}:                      _GotoState45Action,
	{_State187, AndExprType}:                      _GotoState40Action,
	{_State187, OrExprType}:                       _GotoState53Action,
	{_State187, InitializableTypeType}:            _GotoState48Action,
	{_State187, ExplicitStructDefType}:            _GotoState47Action,
	{_State187, ExplicitEnumDefType}:              _GotoState46Action,
	{_State187, AnonymousFuncExprType}:            _GotoState41Action,
	{_State188, MulToken}:                         _GotoState122Action,
	{_State188, DivToken}:                         _GotoState120Action,
	{_State188, ModToken}:                         _GotoState121Action,
	{_State188, BitAndToken}:                      _GotoState117Action,
	{_State188, BitLshiftToken}:                   _GotoState118Action,
	{_State188, BitRshiftToken}:                   _GotoState119Action,
	{_State188, MulOpType}:                        _GotoState123Action,
	{_State189, EqualToken}:                       _GotoState109Action,
	{_State189, NotEqualToken}:                    _GotoState114Action,
	{_State189, LessToken}:                        _GotoState112Action,
	{_State189, LessOrEqualToken}:                 _GotoState113Action,
	{_State189, GreaterToken}:                     _GotoState110Action,
	{_State189, GreaterOrEqualToken}:              _GotoState111Action,
	{_State189, CmpOpType}:                        _GotoState115Action,
	{_State190, AddToken}:                         _GotoState103Action,
	{_State190, SubToken}:                         _GotoState106Action,
	{_State190, BitXorToken}:                      _GotoState105Action,
	{_State190, BitOrToken}:                       _GotoState104Action,
	{_State190, AddOpType}:                        _GotoState107Action,
	{_State191, RparenToken}:                      _GotoState260Action,
	{_State191, CommaToken}:                       _GotoState175Action,
	{_State193, ForToken}:                         _GotoState261Action,
	{_State194, IntegerLiteralToken}:              _GotoState26Action,
	{_State194, FloatLiteralToken}:                _GotoState23Action,
	{_State194, RuneLiteralToken}:                 _GotoState33Action,
	{_State194, StringLiteralToken}:               _GotoState34Action,
	{_State194, IdentifierToken}:                  _GotoState25Action,
	{_State194, TrueToken}:                        _GotoState37Action,
	{_State194, FalseToken}:                       _GotoState22Action,
	{_State194, StructToken}:                      _GotoState35Action,
	{_State194, EnumToken}:                        _GotoState21Action,
	{_State194, FuncToken}:                        _GotoState24Action,
	{_State194, LabelDeclToken}:                   _GotoState27Action,
	{_State194, LparenToken}:                      _GotoState30Action,
	{_State194, LbracketToken}:                    _GotoState28Action,
	{_State194, NotToken}:                         _GotoState32Action,
	{_State194, SubToken}:                         _GotoState36Action,
	{_State194, MulToken}:                         _GotoState31Action,
	{_State194, BitNegToken}:                      _GotoState20Action,
	{_State194, BitAndToken}:                      _GotoState19Action,
	{_State194, LexErrorToken}:                    _GotoState29Action,
	{_State194, OptionalLabelDeclType}:            _GotoState134Action,
	{_State194, SequenceExprType}:                 _GotoState262Action,
	{_State194, BlockExprType}:                    _GotoState43Action,
	{_State194, CallExprType}:                     _GotoState44Action,
	{_State194, AtomExprType}:                     _GotoState42Action,
	{_State194, LiteralType}:                      _GotoState50Action,
	{_State194, InitializeExprType}:               _GotoState49Action,
	{_State194, AccessExprType}:                   _GotoState38Action,
	{_State194, PostfixUnaryExprType}:             _GotoState54Action,
	{_State194, PrefixUnaryOpType}:                _GotoState56Action,
	{_State194, PrefixUnaryExprType}:              _GotoState55Action,
	{_State194, MulExprType}:                      _GotoState51Action,
	{_State194, AddExprType}:                      _GotoState39Action,
	{_State194, CmpExprType}:                      _GotoState45Action,
	{_State194, AndExprType}:                      _GotoState40Action,
	{_State194, OrExprType}:                       _GotoState53Action,
	{_State194, InitializableTypeType}:            _GotoState48Action,
	{_State194, ExplicitStructDefType}:            _GotoState47Action,
	{_State194, ExplicitEnumDefType}:              _GotoState46Action,
	{_State194, AnonymousFuncExprType}:            _GotoState41Action,
	{_State195, IntegerLiteralToken}:              _GotoState26Action,
	{_State195, FloatLiteralToken}:                _GotoState23Action,
	{_State195, RuneLiteralToken}:                 _GotoState33Action,
	{_State195, StringLiteralToken}:               _GotoState34Action,
	{_State195, IdentifierToken}:                  _GotoState25Action,
	{_State195, TrueToken}:                        _GotoState37Action,
	{_State195, FalseToken}:                       _GotoState22Action,
	{_State195, StructToken}:                      _GotoState35Action,
	{_State195, EnumToken}:                        _GotoState21Action,
	{_State195, FuncToken}:                        _GotoState24Action,
	{_State195, LabelDeclToken}:                   _GotoState27Action,
	{_State195, LparenToken}:                      _GotoState30Action,
	{_State195, LbracketToken}:                    _GotoState28Action,
	{_State195, NotToken}:                         _GotoState32Action,
	{_State195, SubToken}:                         _GotoState36Action,
	{_State195, MulToken}:                         _GotoState31Action,
	{_State195, BitNegToken}:                      _GotoState20Action,
	{_State195, BitAndToken}:                      _GotoState19Action,
	{_State195, LexErrorToken}:                    _GotoState29Action,
	{_State195, OptionalLabelDeclType}:            _GotoState134Action,
	{_State195, OptionalSequenceExprType}:         _GotoState263Action,
	{_State195, SequenceExprType}:                 _GotoState264Action,
	{_State195, BlockExprType}:                    _GotoState43Action,
	{_State195, CallExprType}:                     _GotoState44Action,
	{_State195, AtomExprType}:                     _GotoState42Action,
	{_State195, LiteralType}:                      _GotoState50Action,
	{_State195, InitializeExprType}:               _GotoState49Action,
	{_State195, AccessExprType}:                   _GotoState38Action,
	{_State195, PostfixUnaryExprType}:             _GotoState54Action,
	{_State195, PrefixUnaryOpType}:                _GotoState56Action,
	{_State195, PrefixUnaryExprType}:              _GotoState55Action,
	{_State195, MulExprType}:                      _GotoState51Action,
	{_State195, AddExprType}:                      _GotoState39Action,
	{_State195, CmpExprType}:                      _GotoState45Action,
	{_State195, AndExprType}:                      _GotoState40Action,
	{_State195, OrExprType}:                       _GotoState53Action,
	{_State195, InitializableTypeType}:            _GotoState48Action,
	{_State195, ExplicitStructDefType}:            _GotoState47Action,
	{_State195, ExplicitEnumDefType}:              _GotoState46Action,
	{_State195, AnonymousFuncExprType}:            _GotoState41Action,
	{_State196, DoToken}:                          _GotoState265Action,
	{_State197, LbraceToken}:                      _GotoState127Action,
	{_State197, BlockBodyType}:                    _GotoState266Action,
	{_State198, IntegerLiteralToken}:              _GotoState26Action,
	{_State198, FloatLiteralToken}:                _GotoState23Action,
	{_State198, RuneLiteralToken}:                 _GotoState33Action,
	{_State198, StringLiteralToken}:               _GotoState34Action,
	{_State198, IdentifierToken}:                  _GotoState25Action,
	{_State198, TrueToken}:                        _GotoState37Action,
	{_State198, FalseToken}:                       _GotoState22Action,
	{_State198, ReturnToken}:                      _GotoState272Action,
	{_State198, BreakToken}:                       _GotoState268Action,
	{_State198, ContinueToken}:                    _GotoState269Action,
	{_State198, UnsafeToken}:                      _GotoState60Action,
	{_State198, StructToken}:                      _GotoState35Action,
	{_State198, EnumToken}:                        _GotoState21Action,
	{_State198, FuncToken}:                        _GotoState24Action,
	{_State198, AsyncToken}:                       _GotoState267Action,
	{_State198, DeferToken}:                       _GotoState270Action,
	{_State198, LabelDeclToken}:                   _GotoState27Action,
	{_State198, RbraceToken}:                      _GotoState271Action,
	{_State198, LparenToken}:                      _GotoState30Action,
	{_State198, LbracketToken}:                    _GotoState28Action,
	{_State198, NotToken}:                         _GotoState32Action,
	{_State198, SubToken}:                         _GotoState36Action,
	{_State198, MulToken}:                         _GotoState31Action,
	{_State198, BitNegToken}:                      _GotoState20Action,
	{_State198, BitAndToken}:                      _GotoState19Action,
	{_State198, LexErrorToken}:                    _GotoState29Action,
	{_State198, ExpressionType}:                   _GotoState274Action,
	{_State198, OptionalLabelDeclType}:            _GotoState52Action,
	{_State198, SequenceExprType}:                 _GotoState57Action,
	{_State198, BlockExprType}:                    _GotoState43Action,
	{_State198, StatementType}:                    _GotoState278Action,
	{_State198, StatementBodyType}:                _GotoState279Action,
	{_State198, UnsafeStatementType}:              _GotoState280Action,
	{_State198, JumpStatementType}:                _GotoState276Action,
	{_State198, JumpTypeType}:                     _GotoState277Action,
	{_State198, ExpressionsType}:                  _GotoState275Action,
	{_State198, CallExprType}:                     _GotoState44Action,
	{_State198, AtomExprType}:                     _GotoState42Action,
	{_State198, LiteralType}:                      _GotoState50Action,
	{_State198, InitializeExprType}:               _GotoState49Action,
	{_State198, AccessExprType}:                   _GotoState273Action,
	{_State198, PostfixUnaryExprType}:             _GotoState54Action,
	{_State198, PrefixUnaryOpType}:                _GotoState56Action,
	{_State198, PrefixUnaryExprType}:              _GotoState55Action,
	{_State198, MulExprType}:                      _GotoState51Action,
	{_State198, AddExprType}:                      _GotoState39Action,
	{_State198, CmpExprType}:                      _GotoState45Action,
	{_State198, AndExprType}:                      _GotoState40Action,
	{_State198, OrExprType}:                       _GotoState53Action,
	{_State198, InitializableTypeType}:            _GotoState48Action,
	{_State198, ExplicitStructDefType}:            _GotoState47Action,
	{_State198, ExplicitEnumDefType}:              _GotoState46Action,
	{_State198, AnonymousFuncExprType}:            _GotoState41Action,
	{_State199, LbraceToken}:                      _GotoState281Action,
	{_State200, AndToken}:                         _GotoState108Action,
	{_State201, GreaterToken}:                     _GotoState282Action,
	{_State203, UnsafeToken}:                      _GotoState60Action,
	{_State203, RparenToken}:                      _GotoState283Action,
	{_State203, UnsafeStatementType}:              _GotoState286Action,
	{_State203, PackageStatementBodyType}:         _GotoState285Action,
	{_State203, PackageStatementType}:             _GotoState284Action,
	{_State204, IdentifierToken}:                  _GotoState78Action,
	{_State204, StructToken}:                      _GotoState35Action,
	{_State204, EnumToken}:                        _GotoState21Action,
	{_State204, TraitToken}:                       _GotoState83Action,
	{_State204, FuncToken}:                        _GotoState77Action,
	{_State204, LparenToken}:                      _GotoState80Action,
	{_State204, LbracketToken}:                    _GotoState28Action,
	{_State204, DotToken}:                         _GotoState75Action,
	{_State204, QuestionToken}:                    _GotoState81Action,
	{_State204, ExclaimToken}:                     _GotoState76Action,
	{_State204, TildeTildeToken}:                  _GotoState82Action,
	{_State204, BitNegToken}:                      _GotoState74Action,
	{_State204, BitAndToken}:                      _GotoState73Action,
	{_State204, LexErrorToken}:                    _GotoState79Action,
	{_State204, InitializableTypeType}:            _GotoState88Action,
	{_State204, AtomTypeType}:                     _GotoState84Action,
	{_State204, ValueTypeType}:                    _GotoState287Action,
	{_State204, ImplicitStructDefType}:            _GotoState87Action,
	{_State204, ExplicitStructDefType}:            _GotoState47Action,
	{_State204, ImplicitEnumDefType}:              _GotoState86Action,
	{_State204, ExplicitEnumDefType}:              _GotoState46Action,
	{_State204, TraitDefType}:                     _GotoState89Action,
	{_State204, FuncTypeType}:                     _GotoState85Action,
	{_State206, CommaToken}:                       _GotoState288Action,
	{_State207, RbracketToken}:                    _GotoState289Action,
	{_State208, AddToken}:                         _GotoState290Action,
	{_State208, SubToken}:                         _GotoState292Action,
	{_State208, MulToken}:                         _GotoState291Action,
	{_State210, ImplementsToken}:                  _GotoState293Action,
	{_State211, IdentifierToken}:                  _GotoState78Action,
	{_State211, StructToken}:                      _GotoState35Action,
	{_State211, EnumToken}:                        _GotoState21Action,
	{_State211, TraitToken}:                       _GotoState83Action,
	{_State211, FuncToken}:                        _GotoState77Action,
	{_State211, LparenToken}:                      _GotoState80Action,
	{_State211, LbracketToken}:                    _GotoState28Action,
	{_State211, DotToken}:                         _GotoState75Action,
	{_State211, QuestionToken}:                    _GotoState81Action,
	{_State211, ExclaimToken}:                     _GotoState76Action,
	{_State211, TildeTildeToken}:                  _GotoState82Action,
	{_State211, BitNegToken}:                      _GotoState74Action,
	{_State211, BitAndToken}:                      _GotoState73Action,
	{_State211, LexErrorToken}:                    _GotoState79Action,
	{_State211, InitializableTypeType}:            _GotoState88Action,
	{_State211, AtomTypeType}:                     _GotoState84Action,
	{_State211, ValueTypeType}:                    _GotoState294Action,
	{_State211, ImplicitStructDefType}:            _GotoState87Action,
	{_State211, ExplicitStructDefType}:            _GotoState47Action,
	{_State211, ImplicitEnumDefType}:              _GotoState86Action,
	{_State211, ExplicitEnumDefType}:              _GotoState46Action,
	{_State211, TraitDefType}:                     _GotoState89Action,
	{_State211, FuncTypeType}:                     _GotoState85Action,
	{_State214, LparenToken}:                      _GotoState295Action,
	{_State215, IdentifierToken}:                  _GotoState231Action,
	{_State215, DollarLbracketToken}:              _GotoState98Action,
	{_State215, OptionalGenericBindingType}:       _GotoState158Action,
	{_State217, IdentifierToken}:                  _GotoState146Action,
	{_State217, UnsafeToken}:                      _GotoState60Action,
	{_State217, StructToken}:                      _GotoState35Action,
	{_State217, EnumToken}:                        _GotoState21Action,
	{_State217, TraitToken}:                       _GotoState83Action,
	{_State217, FuncToken}:                        _GotoState77Action,
	{_State217, LparenToken}:                      _GotoState80Action,
	{_State217, LbracketToken}:                    _GotoState28Action,
	{_State217, DotToken}:                         _GotoState75Action,
	{_State217, QuestionToken}:                    _GotoState81Action,
	{_State217, ExclaimToken}:                     _GotoState76Action,
	{_State217, TildeTildeToken}:                  _GotoState82Action,
	{_State217, BitNegToken}:                      _GotoState74Action,
	{_State217, BitAndToken}:                      _GotoState73Action,
	{_State217, LexErrorToken}:                    _GotoState79Action,
	{_State217, UnsafeStatementType}:              _GotoState151Action,
	{_State217, InitializableTypeType}:            _GotoState88Action,
	{_State217, AtomTypeType}:                     _GotoState84Action,
	{_State217, ValueTypeType}:                    _GotoState152Action,
	{_State217, FieldDefType}:                     _GotoState149Action,
	{_State217, ImplicitStructDefType}:            _GotoState87Action,
	{_State217, ExplicitStructDefType}:            _GotoState47Action,
	{_State217, EnumValueDefType}:                 _GotoState296Action,
	{_State217, ImplicitEnumDefType}:              _GotoState86Action,
	{_State217, ExplicitEnumDefType}:              _GotoState46Action,
	{_State217, TraitDefType}:                     _GotoState89Action,
	{_State217, FuncTypeType}:                     _GotoState85Action,
	{_State218, IdentifierToken}:                  _GotoState146Action,
	{_State218, UnsafeToken}:                      _GotoState60Action,
	{_State218, StructToken}:                      _GotoState35Action,
	{_State218, EnumToken}:                        _GotoState21Action,
	{_State218, TraitToken}:                       _GotoState83Action,
	{_State218, FuncToken}:                        _GotoState77Action,
	{_State218, LparenToken}:                      _GotoState80Action,
	{_State218, LbracketToken}:                    _GotoState28Action,
	{_State218, DotToken}:                         _GotoState75Action,
	{_State218, QuestionToken}:                    _GotoState81Action,
	{_State218, ExclaimToken}:                     _GotoState76Action,
	{_State218, TildeTildeToken}:                  _GotoState82Action,
	{_State218, BitNegToken}:                      _GotoState74Action,
	{_State218, BitAndToken}:                      _GotoState73Action,
	{_State218, LexErrorToken}:                    _GotoState79Action,
	{_State218, UnsafeStatementType}:              _GotoState151Action,
	{_State218, InitializableTypeType}:            _GotoState88Action,
	{_State218, AtomTypeType}:                     _GotoState84Action,
	{_State218, ValueTypeType}:                    _GotoState152Action,
	{_State218, FieldDefType}:                     _GotoState149Action,
	{_State218, ImplicitStructDefType}:            _GotoState87Action,
	{_State218, ExplicitStructDefType}:            _GotoState47Action,
	{_State218, EnumValueDefType}:                 _GotoState297Action,
	{_State218, ImplicitEnumDefType}:              _GotoState86Action,
	{_State218, ExplicitEnumDefType}:              _GotoState46Action,
	{_State218, TraitDefType}:                     _GotoState89Action,
	{_State218, FuncTypeType}:                     _GotoState85Action,
	{_State220, DefaultToken}:                     _GotoState298Action,
	{_State221, IdentifierToken}:                  _GotoState146Action,
	{_State221, UnsafeToken}:                      _GotoState60Action,
	{_State221, StructToken}:                      _GotoState35Action,
	{_State221, EnumToken}:                        _GotoState21Action,
	{_State221, TraitToken}:                       _GotoState83Action,
	{_State221, FuncToken}:                        _GotoState77Action,
	{_State221, LparenToken}:                      _GotoState80Action,
	{_State221, LbracketToken}:                    _GotoState28Action,
	{_State221, DotToken}:                         _GotoState75Action,
	{_State221, QuestionToken}:                    _GotoState81Action,
	{_State221, ExclaimToken}:                     _GotoState76Action,
	{_State221, TildeTildeToken}:                  _GotoState82Action,
	{_State221, BitNegToken}:                      _GotoState74Action,
	{_State221, BitAndToken}:                      _GotoState73Action,
	{_State221, LexErrorToken}:                    _GotoState79Action,
	{_State221, UnsafeStatementType}:              _GotoState151Action,
	{_State221, InitializableTypeType}:            _GotoState88Action,
	{_State221, AtomTypeType}:                     _GotoState84Action,
	{_State221, ValueTypeType}:                    _GotoState152Action,
	{_State221, FieldDefType}:                     _GotoState149Action,
	{_State221, ImplicitStructDefType}:            _GotoState87Action,
	{_State221, ExplicitStructDefType}:            _GotoState47Action,
	{_State221, EnumValueDefType}:                 _GotoState299Action,
	{_State221, ImplicitEnumDefType}:              _GotoState86Action,
	{_State221, ExplicitEnumDefType}:              _GotoState46Action,
	{_State221, TraitDefType}:                     _GotoState89Action,
	{_State221, FuncTypeType}:                     _GotoState85Action,
	{_State222, IdentifierToken}:                  _GotoState146Action,
	{_State222, UnsafeToken}:                      _GotoState60Action,
	{_State222, StructToken}:                      _GotoState35Action,
	{_State222, EnumToken}:                        _GotoState21Action,
	{_State222, TraitToken}:                       _GotoState83Action,
	{_State222, FuncToken}:                        _GotoState77Action,
	{_State222, LparenToken}:                      _GotoState80Action,
	{_State222, LbracketToken}:                    _GotoState28Action,
	{_State222, DotToken}:                         _GotoState75Action,
	{_State222, QuestionToken}:                    _GotoState81Action,
	{_State222, ExclaimToken}:                     _GotoState76Action,
	{_State222, TildeTildeToken}:                  _GotoState82Action,
	{_State222, BitNegToken}:                      _GotoState74Action,
	{_State222, BitAndToken}:                      _GotoState73Action,
	{_State222, LexErrorToken}:                    _GotoState79Action,
	{_State222, UnsafeStatementType}:              _GotoState151Action,
	{_State222, InitializableTypeType}:            _GotoState88Action,
	{_State222, AtomTypeType}:                     _GotoState84Action,
	{_State222, ValueTypeType}:                    _GotoState152Action,
	{_State222, FieldDefType}:                     _GotoState149Action,
	{_State222, ImplicitStructDefType}:            _GotoState87Action,
	{_State222, ExplicitStructDefType}:            _GotoState47Action,
	{_State222, EnumValueDefType}:                 _GotoState300Action,
	{_State222, ImplicitEnumDefType}:              _GotoState86Action,
	{_State222, ExplicitEnumDefType}:              _GotoState46Action,
	{_State222, TraitDefType}:                     _GotoState89Action,
	{_State222, FuncTypeType}:                     _GotoState85Action,
	{_State223, IdentifierToken}:                  _GotoState78Action,
	{_State223, StructToken}:                      _GotoState35Action,
	{_State223, EnumToken}:                        _GotoState21Action,
	{_State223, TraitToken}:                       _GotoState83Action,
	{_State223, FuncToken}:                        _GotoState77Action,
	{_State223, LparenToken}:                      _GotoState80Action,
	{_State223, LbracketToken}:                    _GotoState28Action,
	{_State223, DotToken}:                         _GotoState75Action,
	{_State223, QuestionToken}:                    _GotoState81Action,
	{_State223, ExclaimToken}:                     _GotoState76Action,
	{_State223, TildeTildeToken}:                  _GotoState82Action,
	{_State223, BitNegToken}:                      _GotoState74Action,
	{_State223, BitAndToken}:                      _GotoState73Action,
	{_State223, LexErrorToken}:                    _GotoState79Action,
	{_State223, InitializableTypeType}:            _GotoState88Action,
	{_State223, AtomTypeType}:                     _GotoState84Action,
	{_State223, ValueTypeType}:                    _GotoState302Action,
	{_State223, ImplicitStructDefType}:            _GotoState87Action,
	{_State223, ExplicitStructDefType}:            _GotoState47Action,
	{_State223, ImplicitEnumDefType}:              _GotoState86Action,
	{_State223, ExplicitEnumDefType}:              _GotoState46Action,
	{_State223, TraitDefType}:                     _GotoState89Action,
	{_State223, ReturnTypeType}:                   _GotoState301Action,
	{_State223, FuncTypeType}:                     _GotoState85Action,
	{_State224, IdentifierToken}:                  _GotoState143Action,
	{_State224, ParameterDefType}:                 _GotoState303Action,
	{_State225, IdentifierToken}:                  _GotoState78Action,
	{_State225, StructToken}:                      _GotoState35Action,
	{_State225, EnumToken}:                        _GotoState21Action,
	{_State225, TraitToken}:                       _GotoState83Action,
	{_State225, FuncToken}:                        _GotoState77Action,
	{_State225, LparenToken}:                      _GotoState80Action,
	{_State225, LbracketToken}:                    _GotoState28Action,
	{_State225, DotToken}:                         _GotoState75Action,
	{_State225, QuestionToken}:                    _GotoState81Action,
	{_State225, ExclaimToken}:                     _GotoState76Action,
	{_State225, TildeTildeToken}:                  _GotoState82Action,
	{_State225, BitNegToken}:                      _GotoState74Action,
	{_State225, BitAndToken}:                      _GotoState73Action,
	{_State225, LexErrorToken}:                    _GotoState79Action,
	{_State225, InitializableTypeType}:            _GotoState88Action,
	{_State225, AtomTypeType}:                     _GotoState84Action,
	{_State225, ValueTypeType}:                    _GotoState304Action,
	{_State225, ImplicitStructDefType}:            _GotoState87Action,
	{_State225, ExplicitStructDefType}:            _GotoState47Action,
	{_State225, ImplicitEnumDefType}:              _GotoState86Action,
	{_State225, ExplicitEnumDefType}:              _GotoState46Action,
	{_State225, TraitDefType}:                     _GotoState89Action,
	{_State225, FuncTypeType}:                     _GotoState85Action,
	{_State226, IdentifierToken}:                  _GotoState78Action,
	{_State226, StructToken}:                      _GotoState35Action,
	{_State226, EnumToken}:                        _GotoState21Action,
	{_State226, TraitToken}:                       _GotoState83Action,
	{_State226, FuncToken}:                        _GotoState77Action,
	{_State226, LparenToken}:                      _GotoState80Action,
	{_State226, LbracketToken}:                    _GotoState28Action,
	{_State226, DotToken}:                         _GotoState215Action,
	{_State226, QuestionToken}:                    _GotoState81Action,
	{_State226, ExclaimToken}:                     _GotoState76Action,
	{_State226, DollarLbracketToken}:              _GotoState98Action,
	{_State226, DotdotdotToken}:                   _GotoState305Action,
	{_State226, TildeTildeToken}:                  _GotoState82Action,
	{_State226, BitNegToken}:                      _GotoState74Action,
	{_State226, BitAndToken}:                      _GotoState73Action,
	{_State226, LexErrorToken}:                    _GotoState79Action,
	{_State226, OptionalGenericBindingType}:       _GotoState162Action,
	{_State226, InitializableTypeType}:            _GotoState88Action,
	{_State226, AtomTypeType}:                     _GotoState84Action,
	{_State226, ValueTypeType}:                    _GotoState306Action,
	{_State226, ImplicitStructDefType}:            _GotoState87Action,
	{_State226, ExplicitStructDefType}:            _GotoState47Action,
	{_State226, ImplicitEnumDefType}:              _GotoState86Action,
	{_State226, ExplicitEnumDefType}:              _GotoState46Action,
	{_State226, TraitDefType}:                     _GotoState89Action,
	{_State226, FuncTypeType}:                     _GotoState85Action,
	{_State227, RparenToken}:                      _GotoState307Action,
	{_State229, CommaToken}:                       _GotoState308Action,
	{_State231, DollarLbracketToken}:              _GotoState98Action,
	{_State231, OptionalGenericBindingType}:       _GotoState309Action,
	{_State232, IdentifierToken}:                  _GotoState146Action,
	{_State232, UnsafeToken}:                      _GotoState60Action,
	{_State232, StructToken}:                      _GotoState35Action,
	{_State232, EnumToken}:                        _GotoState21Action,
	{_State232, TraitToken}:                       _GotoState83Action,
	{_State232, FuncToken}:                        _GotoState77Action,
	{_State232, LparenToken}:                      _GotoState80Action,
	{_State232, LbracketToken}:                    _GotoState28Action,
	{_State232, DotToken}:                         _GotoState75Action,
	{_State232, QuestionToken}:                    _GotoState81Action,
	{_State232, ExclaimToken}:                     _GotoState76Action,
	{_State232, TildeTildeToken}:                  _GotoState82Action,
	{_State232, BitNegToken}:                      _GotoState74Action,
	{_State232, BitAndToken}:                      _GotoState73Action,
	{_State232, LexErrorToken}:                    _GotoState79Action,
	{_State232, UnsafeStatementType}:              _GotoState151Action,
	{_State232, InitializableTypeType}:            _GotoState88Action,
	{_State232, AtomTypeType}:                     _GotoState84Action,
	{_State232, ValueTypeType}:                    _GotoState152Action,
	{_State232, FieldDefType}:                     _GotoState149Action,
	{_State232, ImplicitStructDefType}:            _GotoState87Action,
	{_State232, ExplicitStructDefType}:            _GotoState47Action,
	{_State232, EnumValueDefType}:                 _GotoState310Action,
	{_State232, ImplicitEnumDefType}:              _GotoState86Action,
	{_State232, ExplicitEnumDefType}:              _GotoState46Action,
	{_State232, TraitDefType}:                     _GotoState89Action,
	{_State232, FuncTypeType}:                     _GotoState85Action,
	{_State233, IdentifierToken}:                  _GotoState146Action,
	{_State233, UnsafeToken}:                      _GotoState60Action,
	{_State233, StructToken}:                      _GotoState35Action,
	{_State233, EnumToken}:                        _GotoState21Action,
	{_State233, TraitToken}:                       _GotoState83Action,
	{_State233, FuncToken}:                        _GotoState77Action,
	{_State233, LparenToken}:                      _GotoState80Action,
	{_State233, LbracketToken}:                    _GotoState28Action,
	{_State233, DotToken}:                         _GotoState75Action,
	{_State233, QuestionToken}:                    _GotoState81Action,
	{_State233, ExclaimToken}:                     _GotoState76Action,
	{_State233, TildeTildeToken}:                  _GotoState82Action,
	{_State233, BitNegToken}:                      _GotoState74Action,
	{_State233, BitAndToken}:                      _GotoState73Action,
	{_State233, LexErrorToken}:                    _GotoState79Action,
	{_State233, UnsafeStatementType}:              _GotoState151Action,
	{_State233, InitializableTypeType}:            _GotoState88Action,
	{_State233, AtomTypeType}:                     _GotoState84Action,
	{_State233, ValueTypeType}:                    _GotoState152Action,
	{_State233, FieldDefType}:                     _GotoState149Action,
	{_State233, ImplicitStructDefType}:            _GotoState87Action,
	{_State233, ExplicitStructDefType}:            _GotoState47Action,
	{_State233, EnumValueDefType}:                 _GotoState311Action,
	{_State233, ImplicitEnumDefType}:              _GotoState86Action,
	{_State233, ExplicitEnumDefType}:              _GotoState46Action,
	{_State233, TraitDefType}:                     _GotoState89Action,
	{_State233, FuncTypeType}:                     _GotoState85Action,
	{_State235, IdentifierToken}:                  _GotoState146Action,
	{_State235, UnsafeToken}:                      _GotoState60Action,
	{_State235, StructToken}:                      _GotoState35Action,
	{_State235, EnumToken}:                        _GotoState21Action,
	{_State235, TraitToken}:                       _GotoState83Action,
	{_State235, FuncToken}:                        _GotoState77Action,
	{_State235, LparenToken}:                      _GotoState80Action,
	{_State235, LbracketToken}:                    _GotoState28Action,
	{_State235, DotToken}:                         _GotoState75Action,
	{_State235, QuestionToken}:                    _GotoState81Action,
	{_State235, ExclaimToken}:                     _GotoState76Action,
	{_State235, TildeTildeToken}:                  _GotoState82Action,
	{_State235, BitNegToken}:                      _GotoState74Action,
	{_State235, BitAndToken}:                      _GotoState73Action,
	{_State235, LexErrorToken}:                    _GotoState79Action,
	{_State235, UnsafeStatementType}:              _GotoState151Action,
	{_State235, InitializableTypeType}:            _GotoState88Action,
	{_State235, AtomTypeType}:                     _GotoState84Action,
	{_State235, ValueTypeType}:                    _GotoState152Action,
	{_State235, FieldDefType}:                     _GotoState312Action,
	{_State235, ImplicitStructDefType}:            _GotoState87Action,
	{_State235, ExplicitStructDefType}:            _GotoState47Action,
	{_State235, ImplicitEnumDefType}:              _GotoState86Action,
	{_State235, ExplicitEnumDefType}:              _GotoState46Action,
	{_State235, TraitDefType}:                     _GotoState89Action,
	{_State235, FuncTypeType}:                     _GotoState85Action,
	{_State237, IdentifierToken}:                  _GotoState313Action,
	{_State237, LparenToken}:                      _GotoState160Action,
	{_State238, IdentifierToken}:                  _GotoState78Action,
	{_State238, StructToken}:                      _GotoState35Action,
	{_State238, EnumToken}:                        _GotoState21Action,
	{_State238, TraitToken}:                       _GotoState83Action,
	{_State238, FuncToken}:                        _GotoState77Action,
	{_State238, LparenToken}:                      _GotoState80Action,
	{_State238, LbracketToken}:                    _GotoState28Action,
	{_State238, DotToken}:                         _GotoState215Action,
	{_State238, QuestionToken}:                    _GotoState81Action,
	{_State238, ExclaimToken}:                     _GotoState76Action,
	{_State238, DollarLbracketToken}:              _GotoState98Action,
	{_State238, TildeTildeToken}:                  _GotoState82Action,
	{_State238, BitNegToken}:                      _GotoState74Action,
	{_State238, BitAndToken}:                      _GotoState73Action,
	{_State238, LexErrorToken}:                    _GotoState79Action,
	{_State238, OptionalGenericBindingType}:       _GotoState162Action,
	{_State238, InitializableTypeType}:            _GotoState88Action,
	{_State238, AtomTypeType}:                     _GotoState84Action,
	{_State238, ValueTypeType}:                    _GotoState314Action,
	{_State238, ImplicitStructDefType}:            _GotoState87Action,
	{_State238, ExplicitStructDefType}:            _GotoState47Action,
	{_State238, ImplicitEnumDefType}:              _GotoState86Action,
	{_State238, ExplicitEnumDefType}:              _GotoState46Action,
	{_State238, TraitDefType}:                     _GotoState89Action,
	{_State238, FuncTypeType}:                     _GotoState85Action,
	{_State240, RparenToken}:                      _GotoState315Action,
	{_State241, AddToken}:                         _GotoState290Action,
	{_State241, SubToken}:                         _GotoState292Action,
	{_State241, MulToken}:                         _GotoState291Action,
	{_State242, NewlinesToken}:                    _GotoState317Action,
	{_State242, CommaToken}:                       _GotoState316Action,
	{_State245, RbracketToken}:                    _GotoState318Action,
	{_State246, RbracketToken}:                    _GotoState319Action,
	{_State252, IdentifierToken}:                  _GotoState146Action,
	{_State252, UnsafeToken}:                      _GotoState60Action,
	{_State252, StructToken}:                      _GotoState35Action,
	{_State252, EnumToken}:                        _GotoState21Action,
	{_State252, TraitToken}:                       _GotoState83Action,
	{_State252, FuncToken}:                        _GotoState77Action,
	{_State252, LparenToken}:                      _GotoState80Action,
	{_State252, LbracketToken}:                    _GotoState28Action,
	{_State252, DotToken}:                         _GotoState75Action,
	{_State252, QuestionToken}:                    _GotoState81Action,
	{_State252, ExclaimToken}:                     _GotoState76Action,
	{_State252, TildeTildeToken}:                  _GotoState82Action,
	{_State252, BitNegToken}:                      _GotoState74Action,
	{_State252, BitAndToken}:                      _GotoState73Action,
	{_State252, LexErrorToken}:                    _GotoState79Action,
	{_State252, UnsafeStatementType}:              _GotoState151Action,
	{_State252, InitializableTypeType}:            _GotoState88Action,
	{_State252, AtomTypeType}:                     _GotoState84Action,
	{_State252, ValueTypeType}:                    _GotoState152Action,
	{_State252, FieldDefType}:                     _GotoState320Action,
	{_State252, ImplicitStructDefType}:            _GotoState87Action,
	{_State252, ExplicitStructDefType}:            _GotoState47Action,
	{_State252, ImplicitEnumDefType}:              _GotoState86Action,
	{_State252, ExplicitEnumDefType}:              _GotoState46Action,
	{_State252, TraitDefType}:                     _GotoState89Action,
	{_State252, FuncTypeType}:                     _GotoState85Action,
	{_State253, IdentifierToken}:                  _GotoState146Action,
	{_State253, UnsafeToken}:                      _GotoState60Action,
	{_State253, StructToken}:                      _GotoState35Action,
	{_State253, EnumToken}:                        _GotoState21Action,
	{_State253, TraitToken}:                       _GotoState83Action,
	{_State253, FuncToken}:                        _GotoState77Action,
	{_State253, LparenToken}:                      _GotoState80Action,
	{_State253, LbracketToken}:                    _GotoState28Action,
	{_State253, DotToken}:                         _GotoState75Action,
	{_State253, QuestionToken}:                    _GotoState81Action,
	{_State253, ExclaimToken}:                     _GotoState76Action,
	{_State253, TildeTildeToken}:                  _GotoState82Action,
	{_State253, BitNegToken}:                      _GotoState74Action,
	{_State253, BitAndToken}:                      _GotoState73Action,
	{_State253, LexErrorToken}:                    _GotoState79Action,
	{_State253, UnsafeStatementType}:              _GotoState151Action,
	{_State253, InitializableTypeType}:            _GotoState88Action,
	{_State253, AtomTypeType}:                     _GotoState84Action,
	{_State253, ValueTypeType}:                    _GotoState152Action,
	{_State253, FieldDefType}:                     _GotoState321Action,
	{_State253, ImplicitStructDefType}:            _GotoState87Action,
	{_State253, ExplicitStructDefType}:            _GotoState47Action,
	{_State253, ImplicitEnumDefType}:              _GotoState86Action,
	{_State253, ExplicitEnumDefType}:              _GotoState46Action,
	{_State253, TraitDefType}:                     _GotoState89Action,
	{_State253, FuncTypeType}:                     _GotoState85Action,
	{_State255, IdentifierToken}:                  _GotoState78Action,
	{_State255, StructToken}:                      _GotoState35Action,
	{_State255, EnumToken}:                        _GotoState21Action,
	{_State255, TraitToken}:                       _GotoState83Action,
	{_State255, FuncToken}:                        _GotoState77Action,
	{_State255, LparenToken}:                      _GotoState80Action,
	{_State255, LbracketToken}:                    _GotoState28Action,
	{_State255, DotToken}:                         _GotoState75Action,
	{_State255, QuestionToken}:                    _GotoState81Action,
	{_State255, ExclaimToken}:                     _GotoState76Action,
	{_State255, TildeTildeToken}:                  _GotoState82Action,
	{_State255, BitNegToken}:                      _GotoState74Action,
	{_State255, BitAndToken}:                      _GotoState73Action,
	{_State255, LexErrorToken}:                    _GotoState79Action,
	{_State255, InitializableTypeType}:            _GotoState88Action,
	{_State255, AtomTypeType}:                     _GotoState84Action,
	{_State255, ValueTypeType}:                    _GotoState322Action,
	{_State255, ImplicitStructDefType}:            _GotoState87Action,
	{_State255, ExplicitStructDefType}:            _GotoState47Action,
	{_State255, ImplicitEnumDefType}:              _GotoState86Action,
	{_State255, ExplicitEnumDefType}:              _GotoState46Action,
	{_State255, TraitDefType}:                     _GotoState89Action,
	{_State255, FuncTypeType}:                     _GotoState85Action,
	{_State258, CommaToken}:                       _GotoState175Action,
	{_State259, RparenToken}:                      _GotoState323Action,
	{_State261, IntegerLiteralToken}:              _GotoState26Action,
	{_State261, FloatLiteralToken}:                _GotoState23Action,
	{_State261, RuneLiteralToken}:                 _GotoState33Action,
	{_State261, StringLiteralToken}:               _GotoState34Action,
	{_State261, IdentifierToken}:                  _GotoState25Action,
	{_State261, TrueToken}:                        _GotoState37Action,
	{_State261, FalseToken}:                       _GotoState22Action,
	{_State261, StructToken}:                      _GotoState35Action,
	{_State261, EnumToken}:                        _GotoState21Action,
	{_State261, FuncToken}:                        _GotoState24Action,
	{_State261, LabelDeclToken}:                   _GotoState27Action,
	{_State261, LparenToken}:                      _GotoState30Action,
	{_State261, LbracketToken}:                    _GotoState28Action,
	{_State261, NotToken}:                         _GotoState32Action,
	{_State261, SubToken}:                         _GotoState36Action,
	{_State261, MulToken}:                         _GotoState31Action,
	{_State261, BitNegToken}:                      _GotoState20Action,
	{_State261, BitAndToken}:                      _GotoState19Action,
	{_State261, LexErrorToken}:                    _GotoState29Action,
	{_State261, OptionalLabelDeclType}:            _GotoState134Action,
	{_State261, SequenceExprType}:                 _GotoState324Action,
	{_State261, BlockExprType}:                    _GotoState43Action,
	{_State261, CallExprType}:                     _GotoState44Action,
	{_State261, AtomExprType}:                     _GotoState42Action,
	{_State261, LiteralType}:                      _GotoState50Action,
	{_State261, InitializeExprType}:               _GotoState49Action,
	{_State261, AccessExprType}:                   _GotoState38Action,
	{_State261, PostfixUnaryExprType}:             _GotoState54Action,
	{_State261, PrefixUnaryOpType}:                _GotoState56Action,
	{_State261, PrefixUnaryExprType}:              _GotoState55Action,
	{_State261, MulExprType}:                      _GotoState51Action,
	{_State261, AddExprType}:                      _GotoState39Action,
	{_State261, CmpExprType}:                      _GotoState45Action,
	{_State261, AndExprType}:                      _GotoState40Action,
	{_State261, OrExprType}:                       _GotoState53Action,
	{_State261, InitializableTypeType}:            _GotoState48Action,
	{_State261, ExplicitStructDefType}:            _GotoState47Action,
	{_State261, ExplicitEnumDefType}:              _GotoState46Action,
	{_State261, AnonymousFuncExprType}:            _GotoState41Action,
	{_State262, DoToken}:                          _GotoState325Action,
	{_State263, SemicolonToken}:                   _GotoState326Action,
	{_State265, LbraceToken}:                      _GotoState127Action,
	{_State265, BlockBodyType}:                    _GotoState327Action,
	{_State266, ElseToken}:                        _GotoState328Action,
	{_State267, IntegerLiteralToken}:              _GotoState26Action,
	{_State267, FloatLiteralToken}:                _GotoState23Action,
	{_State267, RuneLiteralToken}:                 _GotoState33Action,
	{_State267, StringLiteralToken}:               _GotoState34Action,
	{_State267, IdentifierToken}:                  _GotoState25Action,
	{_State267, TrueToken}:                        _GotoState37Action,
	{_State267, FalseToken}:                       _GotoState22Action,
	{_State267, StructToken}:                      _GotoState35Action,
	{_State267, EnumToken}:                        _GotoState21Action,
	{_State267, FuncToken}:                        _GotoState24Action,
	{_State267, LabelDeclToken}:                   _GotoState27Action,
	{_State267, LparenToken}:                      _GotoState30Action,
	{_State267, LbracketToken}:                    _GotoState28Action,
	{_State267, LexErrorToken}:                    _GotoState29Action,
	{_State267, OptionalLabelDeclType}:            _GotoState134Action,
	{_State267, BlockExprType}:                    _GotoState43Action,
	{_State267, CallExprType}:                     _GotoState330Action,
	{_State267, AtomExprType}:                     _GotoState42Action,
	{_State267, LiteralType}:                      _GotoState50Action,
	{_State267, InitializeExprType}:               _GotoState49Action,
	{_State267, AccessExprType}:                   _GotoState329Action,
	{_State267, InitializableTypeType}:            _GotoState48Action,
	{_State267, ExplicitStructDefType}:            _GotoState47Action,
	{_State267, ExplicitEnumDefType}:              _GotoState46Action,
	{_State267, AnonymousFuncExprType}:            _GotoState41Action,
	{_State270, IntegerLiteralToken}:              _GotoState26Action,
	{_State270, FloatLiteralToken}:                _GotoState23Action,
	{_State270, RuneLiteralToken}:                 _GotoState33Action,
	{_State270, StringLiteralToken}:               _GotoState34Action,
	{_State270, IdentifierToken}:                  _GotoState25Action,
	{_State270, TrueToken}:                        _GotoState37Action,
	{_State270, FalseToken}:                       _GotoState22Action,
	{_State270, StructToken}:                      _GotoState35Action,
	{_State270, EnumToken}:                        _GotoState21Action,
	{_State270, FuncToken}:                        _GotoState24Action,
	{_State270, LabelDeclToken}:                   _GotoState27Action,
	{_State270, LparenToken}:                      _GotoState30Action,
	{_State270, LbracketToken}:                    _GotoState28Action,
	{_State270, LexErrorToken}:                    _GotoState29Action,
	{_State270, OptionalLabelDeclType}:            _GotoState134Action,
	{_State270, BlockExprType}:                    _GotoState43Action,
	{_State270, CallExprType}:                     _GotoState331Action,
	{_State270, AtomExprType}:                     _GotoState42Action,
	{_State270, LiteralType}:                      _GotoState50Action,
	{_State270, InitializeExprType}:               _GotoState49Action,
	{_State270, AccessExprType}:                   _GotoState329Action,
	{_State270, InitializableTypeType}:            _GotoState48Action,
	{_State270, ExplicitStructDefType}:            _GotoState47Action,
	{_State270, ExplicitEnumDefType}:              _GotoState46Action,
	{_State270, AnonymousFuncExprType}:            _GotoState41Action,
	{_State273, LbracketToken}:                    _GotoState100Action,
	{_State273, DotToken}:                         _GotoState99Action,
	{_State273, QuestionToken}:                    _GotoState101Action,
	{_State273, DollarLbracketToken}:              _GotoState98Action,
	{_State273, AddAssignToken}:                   _GotoState332Action,
	{_State273, SubAssignToken}:                   _GotoState343Action,
	{_State273, MulAssignToken}:                   _GotoState342Action,
	{_State273, DivAssignToken}:                   _GotoState340Action,
	{_State273, ModAssignToken}:                   _GotoState341Action,
	{_State273, AddOneAssignToken}:                _GotoState333Action,
	{_State273, SubOneAssignToken}:                _GotoState344Action,
	{_State273, BitNegAssignToken}:                _GotoState336Action,
	{_State273, BitAndAssignToken}:                _GotoState334Action,
	{_State273, BitOrAssignToken}:                 _GotoState337Action,
	{_State273, BitXorAssignToken}:                _GotoState339Action,
	{_State273, BitLshiftAssignToken}:             _GotoState335Action,
	{_State273, BitRshiftAssignToken}:             _GotoState338Action,
	{_State273, UnaryOpAssignType}:                _GotoState346Action,
	{_State273, BinaryOpAssignType}:               _GotoState345Action,
	{_State273, OptionalGenericBindingType}:       _GotoState102Action,
	{_State275, CommaToken}:                       _GotoState347Action,
	{_State277, JumpLabelToken}:                   _GotoState348Action,
	{_State277, OptionalJumpLabelType}:            _GotoState349Action,
	{_State279, NewlinesToken}:                    _GotoState350Action,
	{_State279, SemicolonToken}:                   _GotoState351Action,
	{_State281, CaseToken}:                        _GotoState352Action,
	{_State282, StringLiteralToken}:               _GotoState353Action,
	{_State285, NewlinesToken}:                    _GotoState354Action,
	{_State285, SemicolonToken}:                   _GotoState355Action,
	{_State288, IdentifierToken}:                  _GotoState204Action,
	{_State288, GenericParameterDefType}:          _GotoState356Action,
	{_State290, IdentifierToken}:                  _GotoState78Action,
	{_State290, StructToken}:                      _GotoState35Action,
	{_State290, EnumToken}:                        _GotoState21Action,
	{_State290, TraitToken}:                       _GotoState83Action,
	{_State290, FuncToken}:                        _GotoState77Action,
	{_State290, LparenToken}:                      _GotoState80Action,
	{_State290, LbracketToken}:                    _GotoState28Action,
	{_State290, DotToken}:                         _GotoState75Action,
	{_State290, QuestionToken}:                    _GotoState81Action,
	{_State290, ExclaimToken}:                     _GotoState76Action,
	{_State290, TildeTildeToken}:                  _GotoState82Action,
	{_State290, BitNegToken}:                      _GotoState74Action,
	{_State290, BitAndToken}:                      _GotoState73Action,
	{_State290, LexErrorToken}:                    _GotoState79Action,
	{_State290, InitializableTypeType}:            _GotoState88Action,
	{_State290, AtomTypeType}:                     _GotoState84Action,
	{_State290, ValueTypeType}:                    _GotoState357Action,
	{_State290, ImplicitStructDefType}:            _GotoState87Action,
	{_State290, ExplicitStructDefType}:            _GotoState47Action,
	{_State290, ImplicitEnumDefType}:              _GotoState86Action,
	{_State290, ExplicitEnumDefType}:              _GotoState46Action,
	{_State290, TraitDefType}:                     _GotoState89Action,
	{_State290, FuncTypeType}:                     _GotoState85Action,
	{_State291, IdentifierToken}:                  _GotoState78Action,
	{_State291, StructToken}:                      _GotoState35Action,
	{_State291, EnumToken}:                        _GotoState21Action,
	{_State291, TraitToken}:                       _GotoState83Action,
	{_State291, FuncToken}:                        _GotoState77Action,
	{_State291, LparenToken}:                      _GotoState80Action,
	{_State291, LbracketToken}:                    _GotoState28Action,
	{_State291, DotToken}:                         _GotoState75Action,
	{_State291, QuestionToken}:                    _GotoState81Action,
	{_State291, ExclaimToken}:                     _GotoState76Action,
	{_State291, TildeTildeToken}:                  _GotoState82Action,
	{_State291, BitNegToken}:                      _GotoState74Action,
	{_State291, BitAndToken}:                      _GotoState73Action,
	{_State291, LexErrorToken}:                    _GotoState79Action,
	{_State291, InitializableTypeType}:            _GotoState88Action,
	{_State291, AtomTypeType}:                     _GotoState84Action,
	{_State291, ValueTypeType}:                    _GotoState358Action,
	{_State291, ImplicitStructDefType}:            _GotoState87Action,
	{_State291, ExplicitStructDefType}:            _GotoState47Action,
	{_State291, ImplicitEnumDefType}:              _GotoState86Action,
	{_State291, ExplicitEnumDefType}:              _GotoState46Action,
	{_State291, TraitDefType}:                     _GotoState89Action,
	{_State291, FuncTypeType}:                     _GotoState85Action,
	{_State292, IdentifierToken}:                  _GotoState78Action,
	{_State292, StructToken}:                      _GotoState35Action,
	{_State292, EnumToken}:                        _GotoState21Action,
	{_State292, TraitToken}:                       _GotoState83Action,
	{_State292, FuncToken}:                        _GotoState77Action,
	{_State292, LparenToken}:                      _GotoState80Action,
	{_State292, LbracketToken}:                    _GotoState28Action,
	{_State292, DotToken}:                         _GotoState75Action,
	{_State292, QuestionToken}:                    _GotoState81Action,
	{_State292, ExclaimToken}:                     _GotoState76Action,
	{_State292, TildeTildeToken}:                  _GotoState82Action,
	{_State292, BitNegToken}:                      _GotoState74Action,
	{_State292, BitAndToken}:                      _GotoState73Action,
	{_State292, LexErrorToken}:                    _GotoState79Action,
	{_State292, InitializableTypeType}:            _GotoState88Action,
	{_State292, AtomTypeType}:                     _GotoState84Action,
	{_State292, ValueTypeType}:                    _GotoState359Action,
	{_State292, ImplicitStructDefType}:            _GotoState87Action,
	{_State292, ExplicitStructDefType}:            _GotoState47Action,
	{_State292, ImplicitEnumDefType}:              _GotoState86Action,
	{_State292, ExplicitEnumDefType}:              _GotoState46Action,
	{_State292, TraitDefType}:                     _GotoState89Action,
	{_State292, FuncTypeType}:                     _GotoState85Action,
	{_State293, IdentifierToken}:                  _GotoState78Action,
	{_State293, StructToken}:                      _GotoState35Action,
	{_State293, EnumToken}:                        _GotoState21Action,
	{_State293, TraitToken}:                       _GotoState83Action,
	{_State293, FuncToken}:                        _GotoState77Action,
	{_State293, LparenToken}:                      _GotoState80Action,
	{_State293, LbracketToken}:                    _GotoState28Action,
	{_State293, DotToken}:                         _GotoState75Action,
	{_State293, QuestionToken}:                    _GotoState81Action,
	{_State293, ExclaimToken}:                     _GotoState76Action,
	{_State293, TildeTildeToken}:                  _GotoState82Action,
	{_State293, BitNegToken}:                      _GotoState74Action,
	{_State293, BitAndToken}:                      _GotoState73Action,
	{_State293, LexErrorToken}:                    _GotoState79Action,
	{_State293, InitializableTypeType}:            _GotoState88Action,
	{_State293, AtomTypeType}:                     _GotoState84Action,
	{_State293, ValueTypeType}:                    _GotoState360Action,
	{_State293, ImplicitStructDefType}:            _GotoState87Action,
	{_State293, ExplicitStructDefType}:            _GotoState47Action,
	{_State293, ImplicitEnumDefType}:              _GotoState86Action,
	{_State293, ExplicitEnumDefType}:              _GotoState46Action,
	{_State293, TraitDefType}:                     _GotoState89Action,
	{_State293, FuncTypeType}:                     _GotoState85Action,
	{_State295, IdentifierToken}:                  _GotoState143Action,
	{_State295, ParameterDefType}:                 _GotoState154Action,
	{_State295, ParameterDefsType}:                _GotoState155Action,
	{_State295, OptionalParameterDefsType}:        _GotoState361Action,
	{_State301, LbraceToken}:                      _GotoState127Action,
	{_State301, BlockBodyType}:                    _GotoState362Action,
	{_State305, IdentifierToken}:                  _GotoState78Action,
	{_State305, StructToken}:                      _GotoState35Action,
	{_State305, EnumToken}:                        _GotoState21Action,
	{_State305, TraitToken}:                       _GotoState83Action,
	{_State305, FuncToken}:                        _GotoState77Action,
	{_State305, LparenToken}:                      _GotoState80Action,
	{_State305, LbracketToken}:                    _GotoState28Action,
	{_State305, DotToken}:                         _GotoState75Action,
	{_State305, QuestionToken}:                    _GotoState81Action,
	{_State305, ExclaimToken}:                     _GotoState76Action,
	{_State305, TildeTildeToken}:                  _GotoState82Action,
	{_State305, BitNegToken}:                      _GotoState74Action,
	{_State305, BitAndToken}:                      _GotoState73Action,
	{_State305, LexErrorToken}:                    _GotoState79Action,
	{_State305, InitializableTypeType}:            _GotoState88Action,
	{_State305, AtomTypeType}:                     _GotoState84Action,
	{_State305, ValueTypeType}:                    _GotoState363Action,
	{_State305, ImplicitStructDefType}:            _GotoState87Action,
	{_State305, ExplicitStructDefType}:            _GotoState47Action,
	{_State305, ImplicitEnumDefType}:              _GotoState86Action,
	{_State305, ExplicitEnumDefType}:              _GotoState46Action,
	{_State305, TraitDefType}:                     _GotoState89Action,
	{_State305, FuncTypeType}:                     _GotoState85Action,
	{_State307, IdentifierToken}:                  _GotoState78Action,
	{_State307, StructToken}:                      _GotoState35Action,
	{_State307, EnumToken}:                        _GotoState21Action,
	{_State307, TraitToken}:                       _GotoState83Action,
	{_State307, FuncToken}:                        _GotoState77Action,
	{_State307, LparenToken}:                      _GotoState80Action,
	{_State307, LbracketToken}:                    _GotoState28Action,
	{_State307, DotToken}:                         _GotoState75Action,
	{_State307, QuestionToken}:                    _GotoState81Action,
	{_State307, ExclaimToken}:                     _GotoState76Action,
	{_State307, TildeTildeToken}:                  _GotoState82Action,
	{_State307, BitNegToken}:                      _GotoState74Action,
	{_State307, BitAndToken}:                      _GotoState73Action,
	{_State307, LexErrorToken}:                    _GotoState79Action,
	{_State307, InitializableTypeType}:            _GotoState88Action,
	{_State307, AtomTypeType}:                     _GotoState84Action,
	{_State307, ValueTypeType}:                    _GotoState302Action,
	{_State307, ImplicitStructDefType}:            _GotoState87Action,
	{_State307, ExplicitStructDefType}:            _GotoState47Action,
	{_State307, ImplicitEnumDefType}:              _GotoState86Action,
	{_State307, ExplicitEnumDefType}:              _GotoState46Action,
	{_State307, TraitDefType}:                     _GotoState89Action,
	{_State307, ReturnTypeType}:                   _GotoState364Action,
	{_State307, FuncTypeType}:                     _GotoState85Action,
	{_State308, IdentifierToken}:                  _GotoState226Action,
	{_State308, StructToken}:                      _GotoState35Action,
	{_State308, EnumToken}:                        _GotoState21Action,
	{_State308, TraitToken}:                       _GotoState83Action,
	{_State308, FuncToken}:                        _GotoState77Action,
	{_State308, LparenToken}:                      _GotoState80Action,
	{_State308, LbracketToken}:                    _GotoState28Action,
	{_State308, DotToken}:                         _GotoState75Action,
	{_State308, QuestionToken}:                    _GotoState81Action,
	{_State308, ExclaimToken}:                     _GotoState76Action,
	{_State308, DotdotdotToken}:                   _GotoState225Action,
	{_State308, TildeTildeToken}:                  _GotoState82Action,
	{_State308, BitNegToken}:                      _GotoState74Action,
	{_State308, BitAndToken}:                      _GotoState73Action,
	{_State308, LexErrorToken}:                    _GotoState79Action,
	{_State308, InitializableTypeType}:            _GotoState88Action,
	{_State308, AtomTypeType}:                     _GotoState84Action,
	{_State308, ValueTypeType}:                    _GotoState230Action,
	{_State308, ImplicitStructDefType}:            _GotoState87Action,
	{_State308, ExplicitStructDefType}:            _GotoState47Action,
	{_State308, ImplicitEnumDefType}:              _GotoState86Action,
	{_State308, ExplicitEnumDefType}:              _GotoState46Action,
	{_State308, TraitDefType}:                     _GotoState89Action,
	{_State308, ParameterDeclType}:                _GotoState365Action,
	{_State308, FuncTypeType}:                     _GotoState85Action,
	{_State313, LparenToken}:                      _GotoState366Action,
	{_State316, IdentifierToken}:                  _GotoState238Action,
	{_State316, UnsafeToken}:                      _GotoState60Action,
	{_State316, StructToken}:                      _GotoState35Action,
	{_State316, EnumToken}:                        _GotoState21Action,
	{_State316, TraitToken}:                       _GotoState83Action,
	{_State316, FuncToken}:                        _GotoState237Action,
	{_State316, LparenToken}:                      _GotoState80Action,
	{_State316, LbracketToken}:                    _GotoState28Action,
	{_State316, DotToken}:                         _GotoState75Action,
	{_State316, QuestionToken}:                    _GotoState81Action,
	{_State316, ExclaimToken}:                     _GotoState76Action,
	{_State316, TildeTildeToken}:                  _GotoState82Action,
	{_State316, BitNegToken}:                      _GotoState74Action,
	{_State316, BitAndToken}:                      _GotoState73Action,
	{_State316, LexErrorToken}:                    _GotoState79Action,
	{_State316, UnsafeStatementType}:              _GotoState244Action,
	{_State316, InitializableTypeType}:            _GotoState88Action,
	{_State316, AtomTypeType}:                     _GotoState84Action,
	{_State316, ValueTypeType}:                    _GotoState209Action,
	{_State316, ImplicitStructDefType}:            _GotoState87Action,
	{_State316, ExplicitStructDefType}:            _GotoState47Action,
	{_State316, ImplicitEnumDefType}:              _GotoState86Action,
	{_State316, ExplicitEnumDefType}:              _GotoState46Action,
	{_State316, TraitAlgebraTypeType}:             _GotoState241Action,
	{_State316, TraitPropertyType}:                _GotoState367Action,
	{_State316, TraitDefType}:                     _GotoState89Action,
	{_State316, FuncTypeType}:                     _GotoState85Action,
	{_State316, MethodSignatureType}:              _GotoState239Action,
	{_State317, IdentifierToken}:                  _GotoState238Action,
	{_State317, UnsafeToken}:                      _GotoState60Action,
	{_State317, StructToken}:                      _GotoState35Action,
	{_State317, EnumToken}:                        _GotoState21Action,
	{_State317, TraitToken}:                       _GotoState83Action,
	{_State317, FuncToken}:                        _GotoState237Action,
	{_State317, LparenToken}:                      _GotoState80Action,
	{_State317, LbracketToken}:                    _GotoState28Action,
	{_State317, DotToken}:                         _GotoState75Action,
	{_State317, QuestionToken}:                    _GotoState81Action,
	{_State317, ExclaimToken}:                     _GotoState76Action,
	{_State317, TildeTildeToken}:                  _GotoState82Action,
	{_State317, BitNegToken}:                      _GotoState74Action,
	{_State317, BitAndToken}:                      _GotoState73Action,
	{_State317, LexErrorToken}:                    _GotoState79Action,
	{_State317, UnsafeStatementType}:              _GotoState244Action,
	{_State317, InitializableTypeType}:            _GotoState88Action,
	{_State317, AtomTypeType}:                     _GotoState84Action,
	{_State317, ValueTypeType}:                    _GotoState209Action,
	{_State317, ImplicitStructDefType}:            _GotoState87Action,
	{_State317, ExplicitStructDefType}:            _GotoState47Action,
	{_State317, ImplicitEnumDefType}:              _GotoState86Action,
	{_State317, ExplicitEnumDefType}:              _GotoState46Action,
	{_State317, TraitAlgebraTypeType}:             _GotoState241Action,
	{_State317, TraitPropertyType}:                _GotoState368Action,
	{_State317, TraitDefType}:                     _GotoState89Action,
	{_State317, FuncTypeType}:                     _GotoState85Action,
	{_State317, MethodSignatureType}:              _GotoState239Action,
	{_State325, LbraceToken}:                      _GotoState127Action,
	{_State325, BlockBodyType}:                    _GotoState369Action,
	{_State326, IntegerLiteralToken}:              _GotoState26Action,
	{_State326, FloatLiteralToken}:                _GotoState23Action,
	{_State326, RuneLiteralToken}:                 _GotoState33Action,
	{_State326, StringLiteralToken}:               _GotoState34Action,
	{_State326, IdentifierToken}:                  _GotoState25Action,
	{_State326, TrueToken}:                        _GotoState37Action,
	{_State326, FalseToken}:                       _GotoState22Action,
	{_State326, StructToken}:                      _GotoState35Action,
	{_State326, EnumToken}:                        _GotoState21Action,
	{_State326, FuncToken}:                        _GotoState24Action,
	{_State326, LabelDeclToken}:                   _GotoState27Action,
	{_State326, LparenToken}:                      _GotoState30Action,
	{_State326, LbracketToken}:                    _GotoState28Action,
	{_State326, NotToken}:                         _GotoState32Action,
	{_State326, SubToken}:                         _GotoState36Action,
	{_State326, MulToken}:                         _GotoState31Action,
	{_State326, BitNegToken}:                      _GotoState20Action,
	{_State326, BitAndToken}:                      _GotoState19Action,
	{_State326, LexErrorToken}:                    _GotoState29Action,
	{_State326, OptionalLabelDeclType}:            _GotoState134Action,
	{_State326, OptionalSequenceExprType}:         _GotoState370Action,
	{_State326, SequenceExprType}:                 _GotoState264Action,
	{_State326, BlockExprType}:                    _GotoState43Action,
	{_State326, CallExprType}:                     _GotoState44Action,
	{_State326, AtomExprType}:                     _GotoState42Action,
	{_State326, LiteralType}:                      _GotoState50Action,
	{_State326, InitializeExprType}:               _GotoState49Action,
	{_State326, AccessExprType}:                   _GotoState38Action,
	{_State326, PostfixUnaryExprType}:             _GotoState54Action,
	{_State326, PrefixUnaryOpType}:                _GotoState56Action,
	{_State326, PrefixUnaryExprType}:              _GotoState55Action,
	{_State326, MulExprType}:                      _GotoState51Action,
	{_State326, AddExprType}:                      _GotoState39Action,
	{_State326, CmpExprType}:                      _GotoState45Action,
	{_State326, AndExprType}:                      _GotoState40Action,
	{_State326, OrExprType}:                       _GotoState53Action,
	{_State326, InitializableTypeType}:            _GotoState48Action,
	{_State326, ExplicitStructDefType}:            _GotoState47Action,
	{_State326, ExplicitEnumDefType}:              _GotoState46Action,
	{_State326, AnonymousFuncExprType}:            _GotoState41Action,
	{_State328, IfToken}:                          _GotoState126Action,
	{_State328, LbraceToken}:                      _GotoState127Action,
	{_State328, IfExprType}:                       _GotoState372Action,
	{_State328, BlockBodyType}:                    _GotoState371Action,
	{_State329, LbracketToken}:                    _GotoState100Action,
	{_State329, DotToken}:                         _GotoState99Action,
	{_State329, DollarLbracketToken}:              _GotoState98Action,
	{_State329, OptionalGenericBindingType}:       _GotoState102Action,
	{_State345, IntegerLiteralToken}:              _GotoState26Action,
	{_State345, FloatLiteralToken}:                _GotoState23Action,
	{_State345, RuneLiteralToken}:                 _GotoState33Action,
	{_State345, StringLiteralToken}:               _GotoState34Action,
	{_State345, IdentifierToken}:                  _GotoState25Action,
	{_State345, TrueToken}:                        _GotoState37Action,
	{_State345, FalseToken}:                       _GotoState22Action,
	{_State345, StructToken}:                      _GotoState35Action,
	{_State345, EnumToken}:                        _GotoState21Action,
	{_State345, FuncToken}:                        _GotoState24Action,
	{_State345, LabelDeclToken}:                   _GotoState27Action,
	{_State345, LparenToken}:                      _GotoState30Action,
	{_State345, LbracketToken}:                    _GotoState28Action,
	{_State345, NotToken}:                         _GotoState32Action,
	{_State345, SubToken}:                         _GotoState36Action,
	{_State345, MulToken}:                         _GotoState31Action,
	{_State345, BitNegToken}:                      _GotoState20Action,
	{_State345, BitAndToken}:                      _GotoState19Action,
	{_State345, LexErrorToken}:                    _GotoState29Action,
	{_State345, ExpressionType}:                   _GotoState373Action,
	{_State345, OptionalLabelDeclType}:            _GotoState52Action,
	{_State345, SequenceExprType}:                 _GotoState57Action,
	{_State345, BlockExprType}:                    _GotoState43Action,
	{_State345, CallExprType}:                     _GotoState44Action,
	{_State345, AtomExprType}:                     _GotoState42Action,
	{_State345, LiteralType}:                      _GotoState50Action,
	{_State345, InitializeExprType}:               _GotoState49Action,
	{_State345, AccessExprType}:                   _GotoState38Action,
	{_State345, PostfixUnaryExprType}:             _GotoState54Action,
	{_State345, PrefixUnaryOpType}:                _GotoState56Action,
	{_State345, PrefixUnaryExprType}:              _GotoState55Action,
	{_State345, MulExprType}:                      _GotoState51Action,
	{_State345, AddExprType}:                      _GotoState39Action,
	{_State345, CmpExprType}:                      _GotoState45Action,
	{_State345, AndExprType}:                      _GotoState40Action,
	{_State345, OrExprType}:                       _GotoState53Action,
	{_State345, InitializableTypeType}:            _GotoState48Action,
	{_State345, ExplicitStructDefType}:            _GotoState47Action,
	{_State345, ExplicitEnumDefType}:              _GotoState46Action,
	{_State345, AnonymousFuncExprType}:            _GotoState41Action,
	{_State347, IntegerLiteralToken}:              _GotoState26Action,
	{_State347, FloatLiteralToken}:                _GotoState23Action,
	{_State347, RuneLiteralToken}:                 _GotoState33Action,
	{_State347, StringLiteralToken}:               _GotoState34Action,
	{_State347, IdentifierToken}:                  _GotoState25Action,
	{_State347, TrueToken}:                        _GotoState37Action,
	{_State347, FalseToken}:                       _GotoState22Action,
	{_State347, StructToken}:                      _GotoState35Action,
	{_State347, EnumToken}:                        _GotoState21Action,
	{_State347, FuncToken}:                        _GotoState24Action,
	{_State347, LabelDeclToken}:                   _GotoState27Action,
	{_State347, LparenToken}:                      _GotoState30Action,
	{_State347, LbracketToken}:                    _GotoState28Action,
	{_State347, NotToken}:                         _GotoState32Action,
	{_State347, SubToken}:                         _GotoState36Action,
	{_State347, MulToken}:                         _GotoState31Action,
	{_State347, BitNegToken}:                      _GotoState20Action,
	{_State347, BitAndToken}:                      _GotoState19Action,
	{_State347, LexErrorToken}:                    _GotoState29Action,
	{_State347, ExpressionType}:                   _GotoState374Action,
	{_State347, OptionalLabelDeclType}:            _GotoState52Action,
	{_State347, SequenceExprType}:                 _GotoState57Action,
	{_State347, BlockExprType}:                    _GotoState43Action,
	{_State347, CallExprType}:                     _GotoState44Action,
	{_State347, AtomExprType}:                     _GotoState42Action,
	{_State347, LiteralType}:                      _GotoState50Action,
	{_State347, InitializeExprType}:               _GotoState49Action,
	{_State347, AccessExprType}:                   _GotoState38Action,
	{_State347, PostfixUnaryExprType}:             _GotoState54Action,
	{_State347, PrefixUnaryOpType}:                _GotoState56Action,
	{_State347, PrefixUnaryExprType}:              _GotoState55Action,
	{_State347, MulExprType}:                      _GotoState51Action,
	{_State347, AddExprType}:                      _GotoState39Action,
	{_State347, CmpExprType}:                      _GotoState45Action,
	{_State347, AndExprType}:                      _GotoState40Action,
	{_State347, OrExprType}:                       _GotoState53Action,
	{_State347, InitializableTypeType}:            _GotoState48Action,
	{_State347, ExplicitStructDefType}:            _GotoState47Action,
	{_State347, ExplicitEnumDefType}:              _GotoState46Action,
	{_State347, AnonymousFuncExprType}:            _GotoState41Action,
	{_State349, IntegerLiteralToken}:              _GotoState26Action,
	{_State349, FloatLiteralToken}:                _GotoState23Action,
	{_State349, RuneLiteralToken}:                 _GotoState33Action,
	{_State349, StringLiteralToken}:               _GotoState34Action,
	{_State349, IdentifierToken}:                  _GotoState25Action,
	{_State349, TrueToken}:                        _GotoState37Action,
	{_State349, FalseToken}:                       _GotoState22Action,
	{_State349, StructToken}:                      _GotoState35Action,
	{_State349, EnumToken}:                        _GotoState21Action,
	{_State349, FuncToken}:                        _GotoState24Action,
	{_State349, LabelDeclToken}:                   _GotoState27Action,
	{_State349, LparenToken}:                      _GotoState30Action,
	{_State349, LbracketToken}:                    _GotoState28Action,
	{_State349, NotToken}:                         _GotoState32Action,
	{_State349, SubToken}:                         _GotoState36Action,
	{_State349, MulToken}:                         _GotoState31Action,
	{_State349, BitNegToken}:                      _GotoState20Action,
	{_State349, BitAndToken}:                      _GotoState19Action,
	{_State349, LexErrorToken}:                    _GotoState29Action,
	{_State349, ExpressionType}:                   _GotoState274Action,
	{_State349, OptionalLabelDeclType}:            _GotoState52Action,
	{_State349, SequenceExprType}:                 _GotoState57Action,
	{_State349, BlockExprType}:                    _GotoState43Action,
	{_State349, ExpressionsType}:                  _GotoState375Action,
	{_State349, OptionalExpressionsType}:          _GotoState376Action,
	{_State349, CallExprType}:                     _GotoState44Action,
	{_State349, AtomExprType}:                     _GotoState42Action,
	{_State349, LiteralType}:                      _GotoState50Action,
	{_State349, InitializeExprType}:               _GotoState49Action,
	{_State349, AccessExprType}:                   _GotoState38Action,
	{_State349, PostfixUnaryExprType}:             _GotoState54Action,
	{_State349, PrefixUnaryOpType}:                _GotoState56Action,
	{_State349, PrefixUnaryExprType}:              _GotoState55Action,
	{_State349, MulExprType}:                      _GotoState51Action,
	{_State349, AddExprType}:                      _GotoState39Action,
	{_State349, CmpExprType}:                      _GotoState45Action,
	{_State349, AndExprType}:                      _GotoState40Action,
	{_State349, OrExprType}:                       _GotoState53Action,
	{_State349, InitializableTypeType}:            _GotoState48Action,
	{_State349, ExplicitStructDefType}:            _GotoState47Action,
	{_State349, ExplicitEnumDefType}:              _GotoState46Action,
	{_State349, AnonymousFuncExprType}:            _GotoState41Action,
	{_State352, DefaultToken}:                     _GotoState377Action,
	{_State361, RparenToken}:                      _GotoState378Action,
	{_State366, IdentifierToken}:                  _GotoState226Action,
	{_State366, StructToken}:                      _GotoState35Action,
	{_State366, EnumToken}:                        _GotoState21Action,
	{_State366, TraitToken}:                       _GotoState83Action,
	{_State366, FuncToken}:                        _GotoState77Action,
	{_State366, LparenToken}:                      _GotoState80Action,
	{_State366, LbracketToken}:                    _GotoState28Action,
	{_State366, DotToken}:                         _GotoState75Action,
	{_State366, QuestionToken}:                    _GotoState81Action,
	{_State366, ExclaimToken}:                     _GotoState76Action,
	{_State366, DotdotdotToken}:                   _GotoState225Action,
	{_State366, TildeTildeToken}:                  _GotoState82Action,
	{_State366, BitNegToken}:                      _GotoState74Action,
	{_State366, BitAndToken}:                      _GotoState73Action,
	{_State366, LexErrorToken}:                    _GotoState79Action,
	{_State366, InitializableTypeType}:            _GotoState88Action,
	{_State366, AtomTypeType}:                     _GotoState84Action,
	{_State366, ValueTypeType}:                    _GotoState230Action,
	{_State366, ImplicitStructDefType}:            _GotoState87Action,
	{_State366, ExplicitStructDefType}:            _GotoState47Action,
	{_State366, ImplicitEnumDefType}:              _GotoState86Action,
	{_State366, ExplicitEnumDefType}:              _GotoState46Action,
	{_State366, TraitDefType}:                     _GotoState89Action,
	{_State366, ParameterDeclType}:                _GotoState228Action,
	{_State366, ParameterDeclsType}:               _GotoState229Action,
	{_State366, OptionalParameterDeclsType}:       _GotoState379Action,
	{_State366, FuncTypeType}:                     _GotoState85Action,
	{_State370, DoToken}:                          _GotoState380Action,
	{_State375, CommaToken}:                       _GotoState347Action,
	{_State377, RbraceToken}:                      _GotoState381Action,
	{_State378, IdentifierToken}:                  _GotoState78Action,
	{_State378, StructToken}:                      _GotoState35Action,
	{_State378, EnumToken}:                        _GotoState21Action,
	{_State378, TraitToken}:                       _GotoState83Action,
	{_State378, FuncToken}:                        _GotoState77Action,
	{_State378, LparenToken}:                      _GotoState80Action,
	{_State378, LbracketToken}:                    _GotoState28Action,
	{_State378, DotToken}:                         _GotoState75Action,
	{_State378, QuestionToken}:                    _GotoState81Action,
	{_State378, ExclaimToken}:                     _GotoState76Action,
	{_State378, TildeTildeToken}:                  _GotoState82Action,
	{_State378, BitNegToken}:                      _GotoState74Action,
	{_State378, BitAndToken}:                      _GotoState73Action,
	{_State378, LexErrorToken}:                    _GotoState79Action,
	{_State378, InitializableTypeType}:            _GotoState88Action,
	{_State378, AtomTypeType}:                     _GotoState84Action,
	{_State378, ValueTypeType}:                    _GotoState302Action,
	{_State378, ImplicitStructDefType}:            _GotoState87Action,
	{_State378, ExplicitStructDefType}:            _GotoState47Action,
	{_State378, ImplicitEnumDefType}:              _GotoState86Action,
	{_State378, ExplicitEnumDefType}:              _GotoState46Action,
	{_State378, TraitDefType}:                     _GotoState89Action,
	{_State378, ReturnTypeType}:                   _GotoState382Action,
	{_State378, FuncTypeType}:                     _GotoState85Action,
	{_State379, RparenToken}:                      _GotoState383Action,
	{_State380, LbraceToken}:                      _GotoState127Action,
	{_State380, BlockBodyType}:                    _GotoState384Action,
	{_State382, LbraceToken}:                      _GotoState127Action,
	{_State382, BlockBodyType}:                    _GotoState385Action,
	{_State383, IdentifierToken}:                  _GotoState78Action,
	{_State383, StructToken}:                      _GotoState35Action,
	{_State383, EnumToken}:                        _GotoState21Action,
	{_State383, TraitToken}:                       _GotoState83Action,
	{_State383, FuncToken}:                        _GotoState77Action,
	{_State383, LparenToken}:                      _GotoState80Action,
	{_State383, LbracketToken}:                    _GotoState28Action,
	{_State383, DotToken}:                         _GotoState75Action,
	{_State383, QuestionToken}:                    _GotoState81Action,
	{_State383, ExclaimToken}:                     _GotoState76Action,
	{_State383, TildeTildeToken}:                  _GotoState82Action,
	{_State383, BitNegToken}:                      _GotoState74Action,
	{_State383, BitAndToken}:                      _GotoState73Action,
	{_State383, LexErrorToken}:                    _GotoState79Action,
	{_State383, InitializableTypeType}:            _GotoState88Action,
	{_State383, AtomTypeType}:                     _GotoState84Action,
	{_State383, ValueTypeType}:                    _GotoState302Action,
	{_State383, ImplicitStructDefType}:            _GotoState87Action,
	{_State383, ExplicitStructDefType}:            _GotoState47Action,
	{_State383, ImplicitEnumDefType}:              _GotoState86Action,
	{_State383, ExplicitEnumDefType}:              _GotoState46Action,
	{_State383, TraitDefType}:                     _GotoState89Action,
	{_State383, ReturnTypeType}:                   _GotoState386Action,
	{_State383, FuncTypeType}:                     _GotoState85Action,
	{_State1, _EndMarker}:                         _ReduceNilToOptionalDefinitionsAction,
	{_State1, _WildcardMarker}:                    _ReduceNilToOptionalNewlinesAction,
	{_State5, _WildcardMarker}:                    _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State13, _EndMarker}:                        _ReduceNewlinesToOptionalDefinitionsAction,
	{_State13, _WildcardMarker}:                   _ReduceNewlinesToOptionalNewlinesAction,
	{_State14, _EndMarker}:                        _ReduceToSourceAction,
	{_State18, IdentifierToken}:                   _ReduceNilToOptionalReceiverAction,
	{_State19, _WildcardMarker}:                   _ReduceBitAndToPrefixUnaryOpAction,
	{_State20, _WildcardMarker}:                   _ReduceBitNegToPrefixUnaryOpAction,
	{_State22, _WildcardMarker}:                   _ReduceFalseToLiteralAction,
	{_State23, _WildcardMarker}:                   _ReduceFloatLiteralToLiteralAction,
	{_State25, _WildcardMarker}:                   _ReduceIdentifierToAtomExprAction,
	{_State26, _WildcardMarker}:                   _ReduceIntegerLiteralToLiteralAction,
	{_State27, _WildcardMarker}:                   _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State29, _WildcardMarker}:                   _ReduceLexErrorToAtomExprAction,
	{_State30, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State30, ColonToken}:                        _ReduceNilToOptionalExpressionAction,
	{_State31, _WildcardMarker}:                   _ReduceMulToPrefixUnaryOpAction,
	{_State32, _WildcardMarker}:                   _ReduceNotToPrefixUnaryOpAction,
	{_State33, _WildcardMarker}:                   _ReduceRuneLiteralToLiteralAction,
	{_State34, _WildcardMarker}:                   _ReduceStringLiteralToLiteralAction,
	{_State36, _WildcardMarker}:                   _ReduceSubToPrefixUnaryOpAction,
	{_State37, _WildcardMarker}:                   _ReduceTrueToLiteralAction,
	{_State38, _WildcardMarker}:                   _ReduceAccessExprToPostfixUnaryExprAction,
	{_State38, LparenToken}:                       _ReduceNilToOptionalGenericBindingAction,
	{_State39, _WildcardMarker}:                   _ReduceAddExprToCmpExprAction,
	{_State40, _WildcardMarker}:                   _ReduceAndExprToOrExprAction,
	{_State41, _WildcardMarker}:                   _ReduceAnonymousFuncExprToAtomExprAction,
	{_State42, _WildcardMarker}:                   _ReduceAtomExprToAccessExprAction,
	{_State43, _WildcardMarker}:                   _ReduceBlockExprToAtomExprAction,
	{_State44, _WildcardMarker}:                   _ReduceCallExprToAccessExprAction,
	{_State45, _WildcardMarker}:                   _ReduceCmpExprToAndExprAction,
	{_State46, _EndMarker}:                        _ReduceExplicitEnumDefToInitializableTypeAction,
	{_State47, _EndMarker}:                        _ReduceExplicitStructDefToInitializableTypeAction,
	{_State49, _WildcardMarker}:                   _ReduceInitializeExprToAtomExprAction,
	{_State50, _WildcardMarker}:                   _ReduceLiteralToAtomExprAction,
	{_State51, _WildcardMarker}:                   _ReduceMulExprToAddExprAction,
	{_State53, _WildcardMarker}:                   _ReduceToSequenceExprAction,
	{_State54, _WildcardMarker}:                   _ReducePostfixUnaryExprToPrefixUnaryExprAction,
	{_State55, _WildcardMarker}:                   _ReducePrefixUnaryExprToMulExprAction,
	{_State56, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State57, _EndMarker}:                        _ReduceSequenceExprToExpressionAction,
	{_State58, _EndMarker}:                        _ReduceCommentToLexInternalTokensAction,
	{_State59, _EndMarker}:                        _ReduceSpacesToLexInternalTokensAction,
	{_State61, _WildcardMarker}:                   _ReduceNilToDefinitionsAction,
	{_State62, _EndMarker}:                        _ReduceNilToOptionalNewlinesAction,
	{_State63, _WildcardMarker}:                   _ReduceNamedFuncDefToDefinitionAction,
	{_State64, _WildcardMarker}:                   _ReducePackageDefToDefinitionAction,
	{_State65, _WildcardMarker}:                   _ReduceTypeDefToDefinitionAction,
	{_State66, _WildcardMarker}:                   _ReduceUnsafeStatementToDefinitionAction,
	{_State67, _WildcardMarker}:                   _ReduceNoSpecToPackageDefAction,
	{_State68, _WildcardMarker}:                   _ReduceNilToOptionalGenericParametersAction,
	{_State72, RparenToken}:                       _ReduceNilToOptionalParameterDefsAction,
	{_State75, _EndMarker}:                        _ReduceNilToOptionalGenericBindingAction,
	{_State78, _EndMarker}:                        _ReduceNilToOptionalGenericBindingAction,
	{_State79, _EndMarker}:                        _ReduceLexErrorToAtomTypeAction,
	{_State80, RparenToken}:                       _ReduceNilToOptionalImplicitFieldDefsAction,
	{_State84, _EndMarker}:                        _ReduceAtomTypeToValueTypeAction,
	{_State85, _EndMarker}:                        _ReduceFuncTypeToAtomTypeAction,
	{_State86, _EndMarker}:                        _ReduceImplicitEnumDefToAtomTypeAction,
	{_State87, _EndMarker}:                        _ReduceImplicitStructDefToAtomTypeAction,
	{_State88, _EndMarker}:                        _ReduceInitializableTypeToAtomTypeAction,
	{_State89, _EndMarker}:                        _ReduceTraitDefToAtomTypeAction,
	{_State91, _WildcardMarker}:                   _ReduceIdentifierToAtomExprAction,
	{_State92, _WildcardMarker}:                   _ReduceArgumentToArgumentsAction,
	{_State94, _WildcardMarker}:                   _ReduceColonExpressionsToArgumentAction,
	{_State95, _WildcardMarker}:                   _ReducePositionalToArgumentAction,
	{_State95, ColonToken}:                        _ReduceExpressionToOptionalExpressionAction,
	{_State97, RparenToken}:                       _ReduceNilToOptionalExplicitFieldDefsAction,
	{_State98, RbracketToken}:                     _ReduceNilToOptionalGenericArgumentsAction,
	{_State100, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State100, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State101, _WildcardMarker}:                  _ReduceQuestionToPostfixUnaryExprAction,
	{_State103, _WildcardMarker}:                  _ReduceAddToAddOpAction,
	{_State104, _WildcardMarker}:                  _ReduceBitOrToAddOpAction,
	{_State105, _WildcardMarker}:                  _ReduceBitXorToAddOpAction,
	{_State106, _WildcardMarker}:                  _ReduceSubToAddOpAction,
	{_State107, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State108, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State109, _WildcardMarker}:                  _ReduceEqualToCmpOpAction,
	{_State110, _WildcardMarker}:                  _ReduceGreaterToCmpOpAction,
	{_State111, _WildcardMarker}:                  _ReduceGreaterOrEqualToCmpOpAction,
	{_State112, _WildcardMarker}:                  _ReduceLessToCmpOpAction,
	{_State113, _WildcardMarker}:                  _ReduceLessOrEqualToCmpOpAction,
	{_State114, _WildcardMarker}:                  _ReduceNotEqualToCmpOpAction,
	{_State115, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State116, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State116, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State117, _WildcardMarker}:                  _ReduceBitAndToMulOpAction,
	{_State118, _WildcardMarker}:                  _ReduceBitLshiftToMulOpAction,
	{_State119, _WildcardMarker}:                  _ReduceBitRshiftToMulOpAction,
	{_State120, _WildcardMarker}:                  _ReduceDivToMulOpAction,
	{_State121, _WildcardMarker}:                  _ReduceModToMulOpAction,
	{_State122, _WildcardMarker}:                  _ReduceMulToMulOpAction,
	{_State123, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State125, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State126, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State127, _WildcardMarker}:                  _ReduceEmptyListToStatementsAction,
	{_State128, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State129, _WildcardMarker}:                  _ReduceToBlockExprAction,
	{_State130, _EndMarker}:                       _ReduceIfExprToExpressionAction,
	{_State131, _EndMarker}:                       _ReduceLoopExprToExpressionAction,
	{_State132, _EndMarker}:                       _ReduceSwitchExprToExpressionAction,
	{_State133, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State135, _WildcardMarker}:                  _ReducePrefixOpToPrefixUnaryExprAction,
	{_State137, _EndMarker}:                       _ReduceNewlinesToOptionalNewlinesAction,
	{_State138, _EndMarker}:                       _ReduceDefinitionsToOptionalDefinitionsAction,
	{_State139, _WildcardMarker}:                  _ReduceEmptyListToPackageStatementsAction,
	{_State140, RbracketToken}:                    _ReduceNilToOptionalGenericParameterDefsAction,
	{_State145, LparenToken}:                      _ReduceNilToOptionalGenericParametersAction,
	{_State146, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State149, _WildcardMarker}:                  _ReduceFieldDefToEnumValueDefAction,
	{_State151, _WildcardMarker}:                  _ReduceUnsafeStatementToFieldDefAction,
	{_State152, _WildcardMarker}:                  _ReduceImplicitToFieldDefAction,
	{_State154, _WildcardMarker}:                  _ReduceParameterDefToParameterDefsAction,
	{_State155, RparenToken}:                      _ReduceParameterDefsToOptionalParameterDefsAction,
	{_State156, _EndMarker}:                       _ReduceReferenceToValueTypeAction,
	{_State157, _EndMarker}:                       _ReducePublicMethodsTraitToValueTypeAction,
	{_State158, _EndMarker}:                       _ReduceInferredToAtomTypeAction,
	{_State159, _EndMarker}:                       _ReduceResultToValueTypeAction,
	{_State160, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State162, _EndMarker}:                       _ReduceNamedToAtomTypeAction,
	{_State164, _WildcardMarker}:                  _ReduceFieldDefToImplicitFieldDefsAction,
	{_State164, OrToken}:                          _ReduceFieldDefToEnumValueDefAction,
	{_State166, RparenToken}:                      _ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction,
	{_State168, _EndMarker}:                       _ReduceOptionalToValueTypeAction,
	{_State169, _EndMarker}:                       _ReducePublicTraitToValueTypeAction,
	{_State170, RparenToken}:                      _ReduceNilToOptionalTraitPropertiesAction,
	{_State173, _EndMarker}:                       _ReduceSliceToInitializableTypeAction,
	{_State174, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State175, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State175, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State176, _WildcardMarker}:                  _ReduceImplicitStructToInitializeExprAction,
	{_State177, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State177, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State177, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State177, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State177, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State178, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State178, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State178, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State178, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State178, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State179, RparenToken}:                      _ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction,
	{_State180, _WildcardMarker}:                  _ReduceFieldDefToExplicitFieldDefsAction,
	{_State182, RbracketToken}:                    _ReduceGenericArgumentsToOptionalGenericArgumentsAction,
	{_State184, _WildcardMarker}:                  _ReduceValueTypeToGenericArgumentsAction,
	{_State185, _WildcardMarker}:                  _ReduceAccessToAccessExprAction,
	{_State187, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State187, RparenToken}:                      _ReduceNilToOptionalArgumentsAction,
	{_State187, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State188, _WildcardMarker}:                  _ReduceOpToAddExprAction,
	{_State189, _WildcardMarker}:                  _ReduceOpToAndExprAction,
	{_State190, _WildcardMarker}:                  _ReduceOpToCmpExprAction,
	{_State192, _WildcardMarker}:                  _ReduceOpToMulExprAction,
	{_State193, _WildcardMarker}:                  _ReduceInfiniteToLoopExprAction,
	{_State194, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State195, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State195, SemicolonToken}:                   _ReduceNilToOptionalSequenceExprAction,
	{_State198, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State200, _WildcardMarker}:                  _ReduceOpToOrExprAction,
	{_State202, _WildcardMarker}:                  _ReduceAddToDefinitionsAction,
	{_State204, _WildcardMarker}:                  _ReduceUnconstrainedToGenericParameterDefAction,
	{_State205, _WildcardMarker}:                  _ReduceGenericParameterDefToGenericParameterDefsAction,
	{_State206, RbracketToken}:                    _ReduceGenericParameterDefsToOptionalGenericParameterDefsAction,
	{_State208, _WildcardMarker}:                  _ReduceAliasToTypeDefAction,
	{_State209, _WildcardMarker}:                  _ReduceValueTypeToTraitAlgebraTypeAction,
	{_State210, _WildcardMarker}:                  _ReduceDefinitionToTypeDefAction,
	{_State212, _WildcardMarker}:                  _ReduceArgToParameterDefAction,
	{_State213, IdentifierToken}:                  _ReduceReceiverToOptionalReceiverAction,
	{_State215, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State216, _WildcardMarker}:                  _ReduceExplicitToFieldDefAction,
	{_State219, _EndMarker}:                       _ReduceToExplicitEnumDefAction,
	{_State223, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State226, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State228, _WildcardMarker}:                  _ReduceParameterDeclToParameterDeclsAction,
	{_State229, RparenToken}:                      _ReduceParameterDeclsToOptionalParameterDeclsAction,
	{_State230, _WildcardMarker}:                  _ReduceUnamedToParameterDeclAction,
	{_State231, _EndMarker}:                       _ReduceNilToOptionalGenericBindingAction,
	{_State234, _EndMarker}:                       _ReduceToImplicitEnumDefAction,
	{_State236, _EndMarker}:                       _ReduceToImplicitStructDefAction,
	{_State238, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State239, _WildcardMarker}:                  _ReduceMethodSignatureToTraitPropertyAction,
	{_State241, _WildcardMarker}:                  _ReduceImplicitToTraitPropertyAction,
	{_State242, RparenToken}:                      _ReduceTraitPropertiesToOptionalTraitPropertiesAction,
	{_State243, _WildcardMarker}:                  _ReduceTraitPropertyToTraitPropertiesAction,
	{_State244, _WildcardMarker}:                  _ReduceUnsafeStatementToTraitPropertyAction,
	{_State247, _WildcardMarker}:                  _ReduceNamedToArgumentAction,
	{_State248, _WildcardMarker}:                  _ReduceAddToArgumentsAction,
	{_State249, _WildcardMarker}:                  _ReduceExpressionToOptionalExpressionAction,
	{_State250, _WildcardMarker}:                  _ReduceAddToColonExpressionsAction,
	{_State251, _WildcardMarker}:                  _ReducePairToColonExpressionsAction,
	{_State254, _EndMarker}:                       _ReduceToExplicitStructDefAction,
	{_State256, _EndMarker}:                       _ReduceBindingToOptionalGenericBindingAction,
	{_State257, _WildcardMarker}:                  _ReduceIndexToAccessExprAction,
	{_State258, RparenToken}:                      _ReduceArgumentsToOptionalArgumentsAction,
	{_State260, _WildcardMarker}:                  _ReduceExplicitToInitializeExprAction,
	{_State261, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State264, DoToken}:                          _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State266, _WildcardMarker}:                  _ReduceNoElseToIfExprAction,
	{_State267, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State268, _WildcardMarker}:                  _ReduceBreakToJumpTypeAction,
	{_State269, _WildcardMarker}:                  _ReduceContinueToJumpTypeAction,
	{_State270, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State271, _EndMarker}:                       _ReduceToBlockBodyAction,
	{_State272, _WildcardMarker}:                  _ReduceReturnToJumpTypeAction,
	{_State273, _WildcardMarker}:                  _ReduceAccessExprToPostfixUnaryExprAction,
	{_State273, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State274, _WildcardMarker}:                  _ReduceExpressionToExpressionsAction,
	{_State275, _WildcardMarker}:                  _ReduceExpressionOrImplicitStructToStatementBodyAction,
	{_State276, _WildcardMarker}:                  _ReduceJumpStatementToStatementBodyAction,
	{_State277, _WildcardMarker}:                  _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State278, _WildcardMarker}:                  _ReduceAddToStatementsAction,
	{_State280, _WildcardMarker}:                  _ReduceUnsafeStatementToStatementBodyAction,
	{_State283, _EndMarker}:                       _ReduceWithSpecToPackageDefAction,
	{_State284, _WildcardMarker}:                  _ReduceAddToPackageStatementsAction,
	{_State286, _WildcardMarker}:                  _ReduceToPackageStatementBodyAction,
	{_State287, _WildcardMarker}:                  _ReduceConstrainedToGenericParameterDefAction,
	{_State289, _WildcardMarker}:                  _ReduceGenericToOptionalGenericParametersAction,
	{_State294, _WildcardMarker}:                  _ReduceVarargToParameterDefAction,
	{_State295, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State296, RparenToken}:                      _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State297, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State297, RparenToken}:                      _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State298, _WildcardMarker}:                  _ReduceDefaultToEnumValueDefAction,
	{_State299, RparenToken}:                      _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State300, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State300, RparenToken}:                      _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State302, _EndMarker}:                       _ReduceValueTypeToReturnTypeAction,
	{_State303, _WildcardMarker}:                  _ReduceAddToParameterDefsAction,
	{_State304, _WildcardMarker}:                  _ReduceUnnamedVarargToParameterDeclAction,
	{_State306, _WildcardMarker}:                  _ReduceArgToParameterDeclAction,
	{_State307, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State309, _EndMarker}:                       _ReduceExternNamedToAtomTypeAction,
	{_State310, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State311, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State312, _WildcardMarker}:                  _ReduceAddToImplicitFieldDefsAction,
	{_State314, _WildcardMarker}:                  _ReduceExplicitToTraitPropertyAction,
	{_State315, _EndMarker}:                       _ReduceToTraitDefAction,
	{_State318, _EndMarker}:                       _ReduceMapToInitializableTypeAction,
	{_State319, _EndMarker}:                       _ReduceArrayToInitializableTypeAction,
	{_State320, _WildcardMarker}:                  _ReduceExplicitToExplicitFieldDefsAction,
	{_State321, _WildcardMarker}:                  _ReduceImplicitToExplicitFieldDefsAction,
	{_State322, _WildcardMarker}:                  _ReduceAddToGenericArgumentsAction,
	{_State323, _WildcardMarker}:                  _ReduceToCallExprAction,
	{_State324, _EndMarker}:                       _ReduceDoWhileToLoopExprAction,
	{_State326, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State326, DoToken}:                          _ReduceNilToOptionalSequenceExprAction,
	{_State327, _EndMarker}:                       _ReduceWhileToLoopExprAction,
	{_State329, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State330, NewlinesToken}:                    _ReduceAsyncToStatementBodyAction,
	{_State330, SemicolonToken}:                   _ReduceAsyncToStatementBodyAction,
	{_State330, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State331, NewlinesToken}:                    _ReduceDeferToStatementBodyAction,
	{_State331, SemicolonToken}:                   _ReduceDeferToStatementBodyAction,
	{_State331, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State332, _WildcardMarker}:                  _ReduceAddAssignToBinaryOpAssignAction,
	{_State333, _WildcardMarker}:                  _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State334, _WildcardMarker}:                  _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State335, _WildcardMarker}:                  _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State336, _WildcardMarker}:                  _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State337, _WildcardMarker}:                  _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State338, _WildcardMarker}:                  _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State339, _WildcardMarker}:                  _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State340, _WildcardMarker}:                  _ReduceDivAssignToBinaryOpAssignAction,
	{_State341, _WildcardMarker}:                  _ReduceModAssignToBinaryOpAssignAction,
	{_State342, _WildcardMarker}:                  _ReduceMulAssignToBinaryOpAssignAction,
	{_State343, _WildcardMarker}:                  _ReduceSubAssignToBinaryOpAssignAction,
	{_State344, _WildcardMarker}:                  _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State345, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State346, _WildcardMarker}:                  _ReduceUnaryOpAssignStatementToStatementBodyAction,
	{_State347, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State348, _WildcardMarker}:                  _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State349, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State349, NewlinesToken}:                    _ReduceNilToOptionalExpressionsAction,
	{_State349, SemicolonToken}:                   _ReduceNilToOptionalExpressionsAction,
	{_State350, _WildcardMarker}:                  _ReduceImplicitToStatementAction,
	{_State351, _WildcardMarker}:                  _ReduceExplicitToStatementAction,
	{_State353, _WildcardMarker}:                  _ReduceToUnsafeStatementAction,
	{_State354, _WildcardMarker}:                  _ReduceImplicitToPackageStatementAction,
	{_State355, _WildcardMarker}:                  _ReduceExplicitToPackageStatementAction,
	{_State356, _WildcardMarker}:                  _ReduceAddToGenericParameterDefsAction,
	{_State357, _WildcardMarker}:                  _ReduceUnionToTraitAlgebraTypeAction,
	{_State358, _WildcardMarker}:                  _ReduceIntersectToTraitAlgebraTypeAction,
	{_State359, _WildcardMarker}:                  _ReduceDifferenceToTraitAlgebraTypeAction,
	{_State360, _EndMarker}:                       _ReduceConstrainedDefToTypeDefAction,
	{_State362, _WildcardMarker}:                  _ReduceToAnonymousFuncExprAction,
	{_State363, _WildcardMarker}:                  _ReduceVarargToParameterDeclAction,
	{_State364, _EndMarker}:                       _ReduceToFuncTypeAction,
	{_State365, _WildcardMarker}:                  _ReduceAddToParameterDeclsAction,
	{_State366, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State367, _WildcardMarker}:                  _ReduceExplicitToTraitPropertiesAction,
	{_State368, _WildcardMarker}:                  _ReduceImplicitToTraitPropertiesAction,
	{_State369, _EndMarker}:                       _ReduceIteratorToLoopExprAction,
	{_State371, _EndMarker}:                       _ReduceIfElseToIfExprAction,
	{_State372, _EndMarker}:                       _ReduceMultiIfElseToIfExprAction,
	{_State373, _WildcardMarker}:                  _ReduceBinaryOpAssignStatementToStatementBodyAction,
	{_State374, _WildcardMarker}:                  _ReduceAddToExpressionsAction,
	{_State375, _WildcardMarker}:                  _ReduceExpressionsToOptionalExpressionsAction,
	{_State376, _WildcardMarker}:                  _ReduceToJumpStatementAction,
	{_State378, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State381, _EndMarker}:                       _ReduceToSwitchExprAction,
	{_State383, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State384, _EndMarker}:                       _ReduceForToLoopExprAction,
	{_State385, _EndMarker}:                       _ReduceToNamedFuncDefAction,
	{_State386, _WildcardMarker}:                  _ReduceToMethodSignatureAction,
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
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      expression -> State 11
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 6:
    Kernel Items:
      #accept: ^.lex_internal_tokens
    Reduce:
      (nil)
    Goto:
      SPACES -> State 59
      COMMENT -> State 58
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
      UNSAFE -> State 60
      TYPE -> State 17
      FUNC -> State 18
      definitions -> State 62
      definition -> State 61
      unsafe_statement -> State 66
      type_def -> State 65
      named_func_def -> State 63
      package_def -> State 64

  State 16:
    Kernel Items:
      package_def: PACKAGE.IDENTIFIER
      package_def: PACKAGE.IDENTIFIER LPAREN package_statements RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 67

  State 17:
    Kernel Items:
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE.IDENTIFIER EQUAL trait_algebra_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 68

  State 18:
    Kernel Items:
      named_func_def: FUNC.optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      LPAREN -> State 69
      optional_receiver -> State 70

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
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 71

  State 22:
    Kernel Items:
      literal: FALSE., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      literal: FLOAT_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 24:
    Kernel Items:
      anonymous_func_expr: FUNC.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 72

  State 25:
    Kernel Items:
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      literal: INTEGER_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 27:
    Kernel Items:
      optional_label_decl: LABEL_DECL., *
    Reduce:
      * -> [optional_label_decl]
    Goto:
      (nil)

  State 28:
    Kernel Items:
      initializable_type: LBRACKET.value_type RBRACKET
      initializable_type: LBRACKET.value_type COMMA INTEGER_LITERAL RBRACKET
      initializable_type: LBRACKET.value_type COLON value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 90
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 29:
    Kernel Items:
      atom_expr: LEX_ERROR., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 30:
    Kernel Items:
      initialize_expr: LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 91
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      expression -> State 95
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 43
      call_expr -> State 44
      arguments -> State 93
      argument -> State 92
      colon_expressions -> State 94
      optional_expression -> State 96
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 31:
    Kernel Items:
      prefix_unary_op: MUL., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      prefix_unary_op: NOT., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 33:
    Kernel Items:
      literal: RUNE_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 34:
    Kernel Items:
      literal: STRING_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 35:
    Kernel Items:
      explicit_struct_def: STRUCT.LPAREN optional_explicit_field_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 97

  State 36:
    Kernel Items:
      prefix_unary_op: SUB., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 37:
    Kernel Items:
      literal: TRUE., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 38:
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
      LBRACKET -> State 100
      DOT -> State 99
      QUESTION -> State 101
      DOLLAR_LBRACKET -> State 98
      optional_generic_binding -> State 102

  State 39:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 103
      SUB -> State 106
      BIT_XOR -> State 105
      BIT_OR -> State 104
      add_op -> State 107

  State 40:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 108

  State 41:
    Kernel Items:
      atom_expr: anonymous_func_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 42:
    Kernel Items:
      access_expr: atom_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 43:
    Kernel Items:
      atom_expr: block_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 44:
    Kernel Items:
      access_expr: call_expr., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 45:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 109
      NOT_EQUAL -> State 114
      LESS -> State 112
      LESS_OR_EQUAL -> State 113
      GREATER -> State 110
      GREATER_OR_EQUAL -> State 111
      cmp_op -> State 115

  State 46:
    Kernel Items:
      initializable_type: explicit_enum_def., $
    Reduce:
      $ -> [initializable_type]
    Goto:
      (nil)

  State 47:
    Kernel Items:
      initializable_type: explicit_struct_def., $
    Reduce:
      $ -> [initializable_type]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      initialize_expr: initializable_type.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 116

  State 49:
    Kernel Items:
      atom_expr: initialize_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 50:
    Kernel Items:
      atom_expr: literal., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 51:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 122
      DIV -> State 120
      MOD -> State 121
      BIT_AND -> State 117
      BIT_LSHIFT -> State 118
      BIT_RSHIFT -> State 119
      mul_op -> State 123

  State 52:
    Kernel Items:
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      IF -> State 126
      SWITCH -> State 128
      FOR -> State 125
      DO -> State 124
      LBRACE -> State 127
      if_expr -> State 130
      switch_expr -> State 132
      loop_expr -> State 131
      block_body -> State 129

  State 53:
    Kernel Items:
      sequence_expr: or_expr., *
      or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    Goto:
      OR -> State 133

  State 54:
    Kernel Items:
      prefix_unary_expr: postfix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 55:
    Kernel Items:
      mul_expr: prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 56:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op.prefix_unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 134
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 135
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 57:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      lex_internal_tokens: COMMENT., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 59:
    Kernel Items:
      lex_internal_tokens: SPACES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 60:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 136

  State 61:
    Kernel Items:
      definitions: definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      optional_definitions: optional_newlines definitions.optional_newlines
      definitions: definitions.NEWLINES definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      NEWLINES -> State 137
      optional_newlines -> State 138

  State 63:
    Kernel Items:
      definition: named_func_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 64:
    Kernel Items:
      definition: package_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      definition: type_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 66:
    Kernel Items:
      definition: unsafe_statement., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 67:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER., *
      package_def: PACKAGE IDENTIFIER.LPAREN package_statements RPAREN
    Reduce:
      * -> [package_def]
    Goto:
      LPAREN -> State 139

  State 68:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE IDENTIFIER.EQUAL trait_algebra_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 140
      EQUAL -> State 141
      optional_generic_parameters -> State 142

  State 69:
    Kernel Items:
      optional_receiver: LPAREN.parameter_def RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 143
      parameter_def -> State 144

  State 70:
    Kernel Items:
      named_func_def: FUNC optional_receiver.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 145

  State 71:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      unsafe_statement -> State 151
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 152
      field_def -> State 149
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      enum_value_def -> State 147
      implicit_enum_value_defs -> State 150
      implicit_enum_def -> State 86
      explicit_enum_value_defs -> State 148
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 72:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 143
      parameter_def -> State 154
      parameter_defs -> State 155
      optional_parameter_defs -> State 153

  State 73:
    Kernel Items:
      value_type: BIT_AND.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 156
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 74:
    Kernel Items:
      value_type: BIT_NEG.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 157
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 75:
    Kernel Items:
      atom_type: DOT.optional_generic_binding
    Reduce:
      $ -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 98
      optional_generic_binding -> State 158

  State 76:
    Kernel Items:
      value_type: EXCLAIM.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 159
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 77:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 160

  State 78:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
    Reduce:
      $ -> [optional_generic_binding]
    Goto:
      DOT -> State 161
      DOLLAR_LBRACKET -> State 98
      optional_generic_binding -> State 162

  State 79:
    Kernel Items:
      atom_type: LEX_ERROR., $
    Reduce:
      $ -> [atom_type]
    Goto:
      (nil)

  State 80:
    Kernel Items:
      implicit_struct_def: LPAREN.optional_implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      IDENTIFIER -> State 146
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      unsafe_statement -> State 151
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 152
      field_def -> State 164
      implicit_field_defs -> State 166
      optional_implicit_field_defs -> State 167
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      enum_value_def -> State 163
      implicit_enum_value_defs -> State 165
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 81:
    Kernel Items:
      value_type: QUESTION.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 168
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 82:
    Kernel Items:
      value_type: TILDE_TILDE.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 169
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 83:
    Kernel Items:
      trait_def: TRAIT.LPAREN optional_trait_properties RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 170

  State 84:
    Kernel Items:
      value_type: atom_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      (nil)

  State 85:
    Kernel Items:
      atom_type: func_type., $
    Reduce:
      $ -> [atom_type]
    Goto:
      (nil)

  State 86:
    Kernel Items:
      atom_type: implicit_enum_def., $
    Reduce:
      $ -> [atom_type]
    Goto:
      (nil)

  State 87:
    Kernel Items:
      atom_type: implicit_struct_def., $
    Reduce:
      $ -> [atom_type]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      atom_type: initializable_type., $
    Reduce:
      $ -> [atom_type]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      atom_type: trait_def., $
    Reduce:
      $ -> [atom_type]
    Goto:
      (nil)

  State 90:
    Kernel Items:
      initializable_type: LBRACKET value_type.RBRACKET
      initializable_type: LBRACKET value_type.COMMA INTEGER_LITERAL RBRACKET
      initializable_type: LBRACKET value_type.COLON value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 173
      COMMA -> State 172
      COLON -> State 171

  State 91:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expression
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      ASSIGN -> State 174

  State 92:
    Kernel Items:
      arguments: argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 93:
    Kernel Items:
      arguments: arguments.COMMA argument
      initialize_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 176
      COMMA -> State 175

  State 94:
    Kernel Items:
      argument: colon_expressions., *
      colon_expressions: colon_expressions.COLON optional_expression
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 177

  State 95:
    Kernel Items:
      argument: expression., *
      optional_expression: expression., COLON
    Reduce:
      * -> [argument]
      COLON -> [optional_expression]
    Goto:
      (nil)

  State 96:
    Kernel Items:
      colon_expressions: optional_expression.COLON optional_expression
    Reduce:
      (nil)
    Goto:
      COLON -> State 178

  State 97:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN.optional_explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      IDENTIFIER -> State 146
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      unsafe_statement -> State 151
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 152
      field_def -> State 180
      implicit_struct_def -> State 87
      explicit_field_defs -> State 179
      optional_explicit_field_defs -> State 181
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 98:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      optional_generic_arguments -> State 183
      generic_arguments -> State 182
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 184
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 99:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 185

  State 100:
    Kernel Items:
      access_expr: access_expr LBRACKET.argument RBRACKET
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 91
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      expression -> State 95
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 43
      call_expr -> State 44
      argument -> State 186
      colon_expressions -> State 94
      optional_expression -> State 96
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 101:
    Kernel Items:
      postfix_unary_expr: access_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 102:
    Kernel Items:
      call_expr: access_expr optional_generic_binding.LPAREN optional_arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 187

  State 103:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 104:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 105:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 106:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 107:
    Kernel Items:
      add_expr: add_expr add_op.mul_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 134
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 188
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 108:
    Kernel Items:
      and_expr: and_expr AND.cmp_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 134
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 189
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 109:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 110:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 111:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 112:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 113:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 114:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      cmp_expr: cmp_expr cmp_op.add_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 134
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 190
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 116:
    Kernel Items:
      initialize_expr: initializable_type LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 91
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      expression -> State 95
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 43
      call_expr -> State 44
      arguments -> State 191
      argument -> State 92
      colon_expressions -> State 94
      optional_expression -> State 96
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 117:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 118:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 119:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 121:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 122:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 123:
    Kernel Items:
      mul_expr: mul_expr mul_op.prefix_unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 134
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 192
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 124:
    Kernel Items:
      loop_expr: DO.block_body
      loop_expr: DO.block_body FOR sequence_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 127
      block_body -> State 193

  State 125:
    Kernel Items:
      loop_expr: FOR.sequence_expr DO block_body
      loop_expr: FOR.IN sequence_expr DO block_body
      loop_expr: FOR.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      IN -> State 194
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      SEMICOLON -> State 195
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 134
      sequence_expr -> State 196
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 126:
    Kernel Items:
      if_expr: IF.sequence_expr block_body
      if_expr: IF.sequence_expr block_body ELSE block_body
      if_expr: IF.sequence_expr block_body ELSE if_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 134
      sequence_expr -> State 197
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 127:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 198

  State 128:
    Kernel Items:
      switch_expr: SWITCH.sequence_expr LBRACE CASE DEFAULT RBRACE
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 134
      sequence_expr -> State 199
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 129:
    Kernel Items:
      block_expr: optional_label_decl block_body., *
    Reduce:
      * -> [block_expr]
    Goto:
      (nil)

  State 130:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 131:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 132:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 133:
    Kernel Items:
      or_expr: or_expr OR.and_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 134
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 200
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 134:
    Kernel Items:
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 127
      block_body -> State 129

  State 135:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 136:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 201

  State 137:
    Kernel Items:
      optional_newlines: NEWLINES., $
      definitions: definitions NEWLINES.definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      PACKAGE -> State 16
      UNSAFE -> State 60
      TYPE -> State 17
      FUNC -> State 18
      definition -> State 202
      unsafe_statement -> State 66
      type_def -> State 65
      named_func_def -> State 63
      package_def -> State 64

  State 138:
    Kernel Items:
      optional_definitions: optional_newlines definitions optional_newlines., $
    Reduce:
      $ -> [optional_definitions]
    Goto:
      (nil)

  State 139:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN.package_statements RPAREN
    Reduce:
      * -> [package_statements]
    Goto:
      package_statements -> State 203

  State 140:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.optional_generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 204
      generic_parameter_def -> State 205
      generic_parameter_defs -> State 206
      optional_generic_parameter_defs -> State 207

  State 141:
    Kernel Items:
      type_def: TYPE IDENTIFIER EQUAL.trait_algebra_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 209
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_algebra_type -> State 208
      trait_def -> State 89
      func_type -> State 85

  State 142:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 210
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 143:
    Kernel Items:
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      DOTDOTDOT -> State 211
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 212
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 144:
    Kernel Items:
      optional_receiver: LPAREN parameter_def.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 213

  State 145:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 140
      optional_generic_parameters -> State 214

  State 146:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 215
      QUESTION -> State 81
      EXCLAIM -> State 76
      DOLLAR_LBRACKET -> State 98
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      optional_generic_binding -> State 162
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 216
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 147:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 217
      OR -> State 218

  State 148:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 219

  State 149:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 220

  State 150:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 221
      OR -> State 222

  State 151:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 152:
    Kernel Items:
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 153:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 223

  State 154:
    Kernel Items:
      parameter_defs: parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 155:
    Kernel Items:
      parameter_defs: parameter_defs.COMMA parameter_def
      optional_parameter_defs: parameter_defs., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      COMMA -> State 224

  State 156:
    Kernel Items:
      value_type: BIT_AND value_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      (nil)

  State 157:
    Kernel Items:
      value_type: BIT_NEG value_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      (nil)

  State 158:
    Kernel Items:
      atom_type: DOT optional_generic_binding., $
    Reduce:
      $ -> [atom_type]
    Goto:
      (nil)

  State 159:
    Kernel Items:
      value_type: EXCLAIM value_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      (nil)

  State 160:
    Kernel Items:
      func_type: FUNC LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 226
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      DOTDOTDOT -> State 225
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 230
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      parameter_decl -> State 228
      parameter_decls -> State 229
      optional_parameter_decls -> State 227
      func_type -> State 85

  State 161:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 231

  State 162:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., $
    Reduce:
      $ -> [atom_type]
    Goto:
      (nil)

  State 163:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 232

  State 164:
    Kernel Items:
      implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 220

  State 165:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 234
      OR -> State 233

  State 166:
    Kernel Items:
      implicit_field_defs: implicit_field_defs.COMMA field_def
      optional_implicit_field_defs: implicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      COMMA -> State 235

  State 167:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 236

  State 168:
    Kernel Items:
      value_type: QUESTION value_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      (nil)

  State 169:
    Kernel Items:
      value_type: TILDE_TILDE value_type., $
    Reduce:
      $ -> [value_type]
    Goto:
      (nil)

  State 170:
    Kernel Items:
      trait_def: TRAIT LPAREN.optional_trait_properties RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      IDENTIFIER -> State 238
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 237
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      unsafe_statement -> State 244
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 209
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_algebra_type -> State 241
      trait_property -> State 243
      trait_properties -> State 242
      optional_trait_properties -> State 240
      trait_def -> State 89
      func_type -> State 85
      method_signature -> State 239

  State 171:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON.value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 245
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 172:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 246

  State 173:
    Kernel Items:
      initializable_type: LBRACKET value_type RBRACKET., $
    Reduce:
      $ -> [initializable_type]
    Goto:
      (nil)

  State 174:
    Kernel Items:
      argument: IDENTIFIER ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      expression -> State 247
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 175:
    Kernel Items:
      arguments: arguments COMMA.argument
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 91
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      expression -> State 95
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 43
      call_expr -> State 44
      argument -> State 248
      colon_expressions -> State 94
      optional_expression -> State 96
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 176:
    Kernel Items:
      initialize_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [initialize_expr]
    Goto:
      (nil)

  State 177:
    Kernel Items:
      colon_expressions: colon_expressions COLON.optional_expression
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_expression]
      RBRACKET -> [optional_expression]
      COMMA -> [optional_expression]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      expression -> State 249
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 43
      call_expr -> State 44
      optional_expression -> State 250
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 178:
    Kernel Items:
      colon_expressions: optional_expression COLON.optional_expression
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_expression]
      RBRACKET -> [optional_expression]
      COMMA -> [optional_expression]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      expression -> State 249
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 43
      call_expr -> State 44
      optional_expression -> State 251
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 179:
    Kernel Items:
      explicit_field_defs: explicit_field_defs.NEWLINES field_def
      explicit_field_defs: explicit_field_defs.COMMA field_def
      optional_explicit_field_defs: explicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      NEWLINES -> State 253
      COMMA -> State 252

  State 180:
    Kernel Items:
      explicit_field_defs: field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 181:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 254

  State 182:
    Kernel Items:
      optional_generic_arguments: generic_arguments., RBRACKET
      generic_arguments: generic_arguments.COMMA value_type
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 255

  State 183:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 256

  State 184:
    Kernel Items:
      generic_arguments: value_type., *
    Reduce:
      * -> [generic_arguments]
    Goto:
      (nil)

  State 185:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 186:
    Kernel Items:
      access_expr: access_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 257

  State 187:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN.optional_arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_arguments]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 91
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      expression -> State 95
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 43
      call_expr -> State 44
      optional_arguments -> State 259
      arguments -> State 258
      argument -> State 92
      colon_expressions -> State 94
      optional_expression -> State 96
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 188:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 122
      DIV -> State 120
      MOD -> State 121
      BIT_AND -> State 117
      BIT_LSHIFT -> State 118
      BIT_RSHIFT -> State 119
      mul_op -> State 123

  State 189:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 109
      NOT_EQUAL -> State 114
      LESS -> State 112
      LESS_OR_EQUAL -> State 113
      GREATER -> State 110
      GREATER_OR_EQUAL -> State 111
      cmp_op -> State 115

  State 190:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 103
      SUB -> State 106
      BIT_XOR -> State 105
      BIT_OR -> State 104
      add_op -> State 107

  State 191:
    Kernel Items:
      arguments: arguments.COMMA argument
      initialize_expr: initializable_type LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 260
      COMMA -> State 175

  State 192:
    Kernel Items:
      mul_expr: mul_expr mul_op prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 193:
    Kernel Items:
      loop_expr: DO block_body., *
      loop_expr: DO block_body.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 261

  State 194:
    Kernel Items:
      loop_expr: FOR IN.sequence_expr DO block_body
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 134
      sequence_expr -> State 262
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 195:
    Kernel Items:
      loop_expr: FOR SEMICOLON.optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      LBRACE -> [optional_label_decl]
      SEMICOLON -> [optional_sequence_expr]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 134
      optional_sequence_expr -> State 263
      sequence_expr -> State 264
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 196:
    Kernel Items:
      loop_expr: FOR sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 265

  State 197:
    Kernel Items:
      if_expr: IF sequence_expr.block_body
      if_expr: IF sequence_expr.block_body ELSE block_body
      if_expr: IF sequence_expr.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 127
      block_body -> State 266

  State 198:
    Kernel Items:
      block_body: LBRACE statements.RBRACE
      statements: statements.statement
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      RETURN -> State 272
      BREAK -> State 268
      CONTINUE -> State 269
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      ASYNC -> State 267
      DEFER -> State 270
      LABEL_DECL -> State 27
      RBRACE -> State 271
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      expression -> State 274
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 43
      statement -> State 278
      statement_body -> State 279
      unsafe_statement -> State 280
      jump_statement -> State 276
      jump_type -> State 277
      expressions -> State 275
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 273
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 199:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.LBRACE CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 281

  State 200:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 108

  State 201:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 282

  State 202:
    Kernel Items:
      definitions: definitions NEWLINES definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 203:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements.RPAREN
      package_statements: package_statements.package_statement
    Reduce:
      (nil)
    Goto:
      UNSAFE -> State 60
      RPAREN -> State 283
      unsafe_statement -> State 286
      package_statement_body -> State 285
      package_statement -> State 284

  State 204:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 287
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 205:
    Kernel Items:
      generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 206:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs.COMMA generic_parameter_def
      optional_generic_parameter_defs: generic_parameter_defs., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      COMMA -> State 288

  State 207:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 289

  State 208:
    Kernel Items:
      type_def: TYPE IDENTIFIER EQUAL trait_algebra_type., *
      trait_algebra_type: trait_algebra_type.MUL value_type
      trait_algebra_type: trait_algebra_type.ADD value_type
      trait_algebra_type: trait_algebra_type.SUB value_type
    Reduce:
      * -> [type_def]
    Goto:
      ADD -> State 290
      SUB -> State 292
      MUL -> State 291

  State 209:
    Kernel Items:
      trait_algebra_type: value_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., *
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 293

  State 211:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 294
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 212:
    Kernel Items:
      parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      optional_receiver: LPAREN parameter_def RPAREN., IDENTIFIER
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      (nil)

  State 214:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 295

  State 215:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 231
      DOLLAR_LBRACKET -> State 98
      optional_generic_binding -> State 158

  State 216:
    Kernel Items:
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      unsafe_statement -> State 151
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 152
      field_def -> State 149
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      enum_value_def -> State 296
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 218:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      unsafe_statement -> State 151
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 152
      field_def -> State 149
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      enum_value_def -> State 297
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 219:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., $
    Reduce:
      $ -> [explicit_enum_def]
    Goto:
      (nil)

  State 220:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 298

  State 221:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      unsafe_statement -> State 151
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 152
      field_def -> State 149
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      enum_value_def -> State 299
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 222:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      unsafe_statement -> State 151
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 152
      field_def -> State 149
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      enum_value_def -> State 300
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 223:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 302
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      return_type -> State 301
      func_type -> State 85

  State 224:
    Kernel Items:
      parameter_defs: parameter_defs COMMA.parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 143
      parameter_def -> State 303

  State 225:
    Kernel Items:
      parameter_decl: DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 304
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 226:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 215
      QUESTION -> State 81
      EXCLAIM -> State 76
      DOLLAR_LBRACKET -> State 98
      DOTDOTDOT -> State 305
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      optional_generic_binding -> State 162
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 306
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 227:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 307

  State 228:
    Kernel Items:
      parameter_decls: parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 229:
    Kernel Items:
      parameter_decls: parameter_decls.COMMA parameter_decl
      optional_parameter_decls: parameter_decls., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      COMMA -> State 308

  State 230:
    Kernel Items:
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 231:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER.optional_generic_binding
    Reduce:
      $ -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 98
      optional_generic_binding -> State 309

  State 232:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      unsafe_statement -> State 151
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 152
      field_def -> State 149
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      enum_value_def -> State 310
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 233:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      unsafe_statement -> State 151
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 152
      field_def -> State 149
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      enum_value_def -> State 311
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 234:
    Kernel Items:
      implicit_enum_def: LPAREN implicit_enum_value_defs RPAREN., $
    Reduce:
      $ -> [implicit_enum_def]
    Goto:
      (nil)

  State 235:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      unsafe_statement -> State 151
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 152
      field_def -> State 312
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 236:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs RPAREN., $
    Reduce:
      $ -> [implicit_struct_def]
    Goto:
      (nil)

  State 237:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 313
      LPAREN -> State 160

  State 238:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      trait_property: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 215
      QUESTION -> State 81
      EXCLAIM -> State 76
      DOLLAR_LBRACKET -> State 98
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      optional_generic_binding -> State 162
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 314
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 239:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 315

  State 241:
    Kernel Items:
      trait_algebra_type: trait_algebra_type.MUL value_type
      trait_algebra_type: trait_algebra_type.ADD value_type
      trait_algebra_type: trait_algebra_type.SUB value_type
      trait_property: trait_algebra_type., *
    Reduce:
      * -> [trait_property]
    Goto:
      ADD -> State 290
      SUB -> State 292
      MUL -> State 291

  State 242:
    Kernel Items:
      trait_properties: trait_properties.NEWLINES trait_property
      trait_properties: trait_properties.COMMA trait_property
      optional_trait_properties: trait_properties., RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      NEWLINES -> State 317
      COMMA -> State 316

  State 243:
    Kernel Items:
      trait_properties: trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 244:
    Kernel Items:
      trait_property: unsafe_statement., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 245:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 318

  State 246:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 319

  State 247:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 248:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 249:
    Kernel Items:
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      (nil)

  State 250:
    Kernel Items:
      colon_expressions: colon_expressions COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 251:
    Kernel Items:
      colon_expressions: optional_expression COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 252:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      unsafe_statement -> State 151
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 152
      field_def -> State 320
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 253:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      unsafe_statement -> State 151
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 152
      field_def -> State 321
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 254:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs RPAREN., $
    Reduce:
      $ -> [explicit_struct_def]
    Goto:
      (nil)

  State 255:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 322
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 256:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., $
    Reduce:
      $ -> [optional_generic_binding]
    Goto:
      (nil)

  State 257:
    Kernel Items:
      access_expr: access_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 258:
    Kernel Items:
      optional_arguments: arguments., RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 175

  State 259:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 323

  State 260:
    Kernel Items:
      initialize_expr: initializable_type LPAREN arguments RPAREN., *
    Reduce:
      * -> [initialize_expr]
    Goto:
      (nil)

  State 261:
    Kernel Items:
      loop_expr: DO block_body FOR.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 134
      sequence_expr -> State 324
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 262:
    Kernel Items:
      loop_expr: FOR IN sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 325

  State 263:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 326

  State 264:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 265:
    Kernel Items:
      loop_expr: FOR sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 127
      block_body -> State 327

  State 266:
    Kernel Items:
      if_expr: IF sequence_expr block_body., *
      if_expr: IF sequence_expr block_body.ELSE block_body
      if_expr: IF sequence_expr block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 328

  State 267:
    Kernel Items:
      statement_body: ASYNC.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      LEX_ERROR -> State 29
      optional_label_decl -> State 134
      block_expr -> State 43
      call_expr -> State 330
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 329
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 268:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 269:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 270:
    Kernel Items:
      statement_body: DEFER.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      LEX_ERROR -> State 29
      optional_label_decl -> State 134
      block_expr -> State 43
      call_expr -> State 331
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 329
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 271:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 272:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 273:
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
      LBRACKET -> State 100
      DOT -> State 99
      QUESTION -> State 101
      DOLLAR_LBRACKET -> State 98
      ADD_ASSIGN -> State 332
      SUB_ASSIGN -> State 343
      MUL_ASSIGN -> State 342
      DIV_ASSIGN -> State 340
      MOD_ASSIGN -> State 341
      ADD_ONE_ASSIGN -> State 333
      SUB_ONE_ASSIGN -> State 344
      BIT_NEG_ASSIGN -> State 336
      BIT_AND_ASSIGN -> State 334
      BIT_OR_ASSIGN -> State 337
      BIT_XOR_ASSIGN -> State 339
      BIT_LSHIFT_ASSIGN -> State 335
      BIT_RSHIFT_ASSIGN -> State 338
      unary_op_assign -> State 346
      binary_op_assign -> State 345
      optional_generic_binding -> State 102

  State 274:
    Kernel Items:
      expressions: expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 275:
    Kernel Items:
      statement_body: expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [statement_body]
    Goto:
      COMMA -> State 347

  State 276:
    Kernel Items:
      statement_body: jump_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 277:
    Kernel Items:
      jump_statement: jump_type.optional_jump_label optional_expressions
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 348
      optional_jump_label -> State 349

  State 278:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 350
      SEMICOLON -> State 351

  State 280:
    Kernel Items:
      statement_body: unsafe_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 281:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE.CASE DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      CASE -> State 352

  State 282:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 353

  State 283:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements RPAREN., $
    Reduce:
      $ -> [package_def]
    Goto:
      (nil)

  State 284:
    Kernel Items:
      package_statements: package_statements package_statement., *
    Reduce:
      * -> [package_statements]
    Goto:
      (nil)

  State 285:
    Kernel Items:
      package_statement: package_statement_body.NEWLINES
      package_statement: package_statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 354
      SEMICOLON -> State 355

  State 286:
    Kernel Items:
      package_statement_body: unsafe_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 287:
    Kernel Items:
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      (nil)

  State 288:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA.generic_parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 204
      generic_parameter_def -> State 356

  State 289:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 290:
    Kernel Items:
      trait_algebra_type: trait_algebra_type ADD.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 357
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 291:
    Kernel Items:
      trait_algebra_type: trait_algebra_type MUL.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 358
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 292:
    Kernel Items:
      trait_algebra_type: trait_algebra_type SUB.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 359
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 293:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 360
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 294:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      (nil)

  State 295:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 143
      parameter_def -> State 154
      parameter_defs -> State 155
      optional_parameter_defs -> State 361

  State 296:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 297:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
      explicit_enum_value_defs: enum_value_def OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 298:
    Kernel Items:
      enum_value_def: field_def ASSIGN DEFAULT., *
    Reduce:
      * -> [enum_value_def]
    Goto:
      (nil)

  State 299:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 300:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 301:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 127
      block_body -> State 362

  State 302:
    Kernel Items:
      return_type: value_type., $
    Reduce:
      $ -> [return_type]
    Goto:
      (nil)

  State 303:
    Kernel Items:
      parameter_defs: parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 304:
    Kernel Items:
      parameter_decl: DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 305:
    Kernel Items:
      parameter_decl: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 363
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      func_type -> State 85

  State 306:
    Kernel Items:
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 307:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 302
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      return_type -> State 364
      func_type -> State 85

  State 308:
    Kernel Items:
      parameter_decls: parameter_decls COMMA.parameter_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 226
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      DOTDOTDOT -> State 225
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 230
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      parameter_decl -> State 365
      func_type -> State 85

  State 309:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER optional_generic_binding., $
    Reduce:
      $ -> [atom_type]
    Goto:
      (nil)

  State 310:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 311:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 312:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [implicit_field_defs]
    Goto:
      (nil)

  State 313:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 366

  State 314:
    Kernel Items:
      trait_property: IDENTIFIER value_type., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 315:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties RPAREN., $
    Reduce:
      $ -> [trait_def]
    Goto:
      (nil)

  State 316:
    Kernel Items:
      trait_properties: trait_properties COMMA.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 238
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 237
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      unsafe_statement -> State 244
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 209
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_algebra_type -> State 241
      trait_property -> State 367
      trait_def -> State 89
      func_type -> State 85
      method_signature -> State 239

  State 317:
    Kernel Items:
      trait_properties: trait_properties NEWLINES.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 238
      UNSAFE -> State 60
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 237
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      unsafe_statement -> State 244
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 209
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_algebra_type -> State 241
      trait_property -> State 368
      trait_def -> State 89
      func_type -> State 85
      method_signature -> State 239

  State 318:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type RBRACKET., $
    Reduce:
      $ -> [initializable_type]
    Goto:
      (nil)

  State 319:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL RBRACKET., $
    Reduce:
      $ -> [initializable_type]
    Goto:
      (nil)

  State 320:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 321:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 322:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
    Reduce:
      * -> [generic_arguments]
    Goto:
      (nil)

  State 323:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 324:
    Kernel Items:
      loop_expr: DO block_body FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 325:
    Kernel Items:
      loop_expr: FOR IN sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 127
      block_body -> State 369

  State 326:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON.optional_sequence_expr DO block_body
    Reduce:
      DO -> [optional_sequence_expr]
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      optional_label_decl -> State 134
      optional_sequence_expr -> State 370
      sequence_expr -> State 264
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 327:
    Kernel Items:
      loop_expr: FOR sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 328:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE.block_body
      if_expr: IF sequence_expr block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 126
      LBRACE -> State 127
      if_expr -> State 372
      block_body -> State 371

  State 329:
    Kernel Items:
      call_expr: access_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET argument RBRACKET
    Reduce:
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 100
      DOT -> State 99
      DOLLAR_LBRACKET -> State 98
      optional_generic_binding -> State 102

  State 330:
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

  State 331:
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

  State 332:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 333:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 334:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 335:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 336:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 337:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 338:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 339:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 340:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 342:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 343:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 344:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 345:
    Kernel Items:
      statement_body: access_expr binary_op_assign.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      expression -> State 373
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 346:
    Kernel Items:
      statement_body: access_expr unary_op_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 347:
    Kernel Items:
      expressions: expressions COMMA.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      expression -> State 374
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 348:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 349:
    Kernel Items:
      jump_statement: jump_type optional_jump_label.optional_expressions
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_expressions]
      SEMICOLON -> [optional_expressions]
    Goto:
      INTEGER_LITERAL -> State 26
      FLOAT_LITERAL -> State 23
      RUNE_LITERAL -> State 33
      STRING_LITERAL -> State 34
      IDENTIFIER -> State 25
      TRUE -> State 37
      FALSE -> State 22
      STRUCT -> State 35
      ENUM -> State 21
      FUNC -> State 24
      LABEL_DECL -> State 27
      LPAREN -> State 30
      LBRACKET -> State 28
      NOT -> State 32
      SUB -> State 36
      MUL -> State 31
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 29
      expression -> State 274
      optional_label_decl -> State 52
      sequence_expr -> State 57
      block_expr -> State 43
      expressions -> State 375
      optional_expressions -> State 376
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 50
      initialize_expr -> State 49
      access_expr -> State 38
      postfix_unary_expr -> State 54
      prefix_unary_op -> State 56
      prefix_unary_expr -> State 55
      mul_expr -> State 51
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 53
      initializable_type -> State 48
      explicit_struct_def -> State 47
      explicit_enum_def -> State 46
      anonymous_func_expr -> State 41

  State 350:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 351:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 352:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE CASE.DEFAULT RBRACE
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 377

  State 353:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 354:
    Kernel Items:
      package_statement: package_statement_body NEWLINES., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 355:
    Kernel Items:
      package_statement: package_statement_body SEMICOLON., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 356:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 357:
    Kernel Items:
      trait_algebra_type: trait_algebra_type ADD value_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 358:
    Kernel Items:
      trait_algebra_type: trait_algebra_type MUL value_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 359:
    Kernel Items:
      trait_algebra_type: trait_algebra_type SUB value_type., *
    Reduce:
      * -> [trait_algebra_type]
    Goto:
      (nil)

  State 360:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      (nil)

  State 361:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 378

  State 362:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 363:
    Kernel Items:
      parameter_decl: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 364:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN return_type., $
    Reduce:
      $ -> [func_type]
    Goto:
      (nil)

  State 365:
    Kernel Items:
      parameter_decls: parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 366:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 226
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      DOTDOTDOT -> State 225
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 230
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      parameter_decl -> State 228
      parameter_decls -> State 229
      optional_parameter_decls -> State 379
      func_type -> State 85

  State 367:
    Kernel Items:
      trait_properties: trait_properties COMMA trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 368:
    Kernel Items:
      trait_properties: trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 369:
    Kernel Items:
      loop_expr: FOR IN sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 370:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 380

  State 371:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 372:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 373:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 374:
    Kernel Items:
      expressions: expressions COMMA expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 375:
    Kernel Items:
      expressions: expressions.COMMA expression
      optional_expressions: expressions., *
    Reduce:
      * -> [optional_expressions]
    Goto:
      COMMA -> State 347

  State 376:
    Kernel Items:
      jump_statement: jump_type optional_jump_label optional_expressions., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 377:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE CASE DEFAULT.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 381

  State 378:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 302
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      return_type -> State 382
      func_type -> State 85

  State 379:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 383

  State 380:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 127
      block_body -> State 384

  State 381:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE CASE DEFAULT RBRACE., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 382:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 127
      block_body -> State 385

  State 383:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 78
      STRUCT -> State 35
      ENUM -> State 21
      TRAIT -> State 83
      FUNC -> State 77
      LPAREN -> State 80
      LBRACKET -> State 28
      DOT -> State 75
      QUESTION -> State 81
      EXCLAIM -> State 76
      TILDE_TILDE -> State 82
      BIT_NEG -> State 74
      BIT_AND -> State 73
      LEX_ERROR -> State 79
      initializable_type -> State 88
      atom_type -> State 84
      value_type -> State 302
      implicit_struct_def -> State 87
      explicit_struct_def -> State 47
      implicit_enum_def -> State 86
      explicit_enum_def -> State 46
      trait_def -> State 89
      return_type -> State 386
      func_type -> State 85

  State 384:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 385:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 386:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

Number of states: 386
Number of shift actions: 2486
Number of reduce actions: 312
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
