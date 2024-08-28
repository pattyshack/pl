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

	// 97:15: switch_expr -> ...
	ToSwitchExpr(Switch_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Lbrace_ *GenericSymbol, CaseBranches_ *GenericSymbol, OptionalDefaultBranch_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 100:2: case_branches -> case_branch: ...
	CaseBranchToCaseBranches(CaseBranch_ *GenericSymbol) (*GenericSymbol, error)

	// 101:2: case_branches -> add: ...
	AddToCaseBranches(CaseBranches_ *GenericSymbol, CaseBranch_ *GenericSymbol) (*GenericSymbol, error)

	// 103:15: case_branch -> ...
	ToCaseBranch(Case_ *GenericSymbol, MatchPattern_ *GenericSymbol, Colon_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 106:0: match_pattern -> ...
	ToMatchPattern() (*GenericSymbol, error)

	// 109:2: optional_default_branch -> default_branch: ...
	DefaultBranchToOptionalDefaultBranch(DefaultBranch_ *GenericSymbol) (*GenericSymbol, error)

	// 110:2: optional_default_branch -> nil: ...
	NilToOptionalDefaultBranch() (*GenericSymbol, error)

	// 112:18: default_branch -> ...
	ToDefaultBranch(Default_ *GenericSymbol, Colon_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 126:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 127:2: loop_expr -> do_while: ...
	DoWhileToLoopExpr(Do_ *GenericSymbol, BlockBody_ *GenericSymbol, For_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 128:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 129:2: loop_expr -> iterator: ...
	IteratorToLoopExpr(For_ *GenericSymbol, In_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 130:2: loop_expr -> for: ...
	ForToLoopExpr(For_ *GenericSymbol, Semicolon_ *GenericSymbol, OptionalSequenceExpr_ *GenericSymbol, Semicolon_2 *GenericSymbol, OptionalSequenceExpr_2 *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 133:2: optional_sequence_expr -> sequence_expr: ...
	SequenceExprToOptionalSequenceExpr(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 134:2: optional_sequence_expr -> nil: ...
	NilToOptionalSequenceExpr() (*GenericSymbol, error)

	// 140:17: sequence_expr -> ...
	ToSequenceExpr(OrExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 142:14: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 144:14: block_body -> ...
	ToBlockBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 147:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 148:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 151:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 152:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 155:2: statement_body -> unsafe_statement: ...
	UnsafeStatementToStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 157:2: statement_body -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToStatementBody(Expressions_ *GenericSymbol) (*GenericSymbol, error)

	// 160:2: statement_body -> async: ...
	AsyncToStatementBody(Async_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 169:2: statement_body -> defer: ...
	DeferToStatementBody(Defer_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 173:2: statement_body -> jump_statement: ...
	JumpStatementToStatementBody(JumpStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 185:2: statement_body -> unary_op_assign_statement: ...
	UnaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, UnaryOpAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 186:2: statement_body -> binary_op_assign_statement: ...
	BinaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 193:2: unary_op_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToUnaryOpAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 194:2: unary_op_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToUnaryOpAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 197:2: binary_op_assign -> ADD_ASSIGN: ...
	AddAssignToBinaryOpAssign(AddAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 198:2: binary_op_assign -> SUB_ASSIGN: ...
	SubAssignToBinaryOpAssign(SubAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 199:2: binary_op_assign -> MUL_ASSIGN: ...
	MulAssignToBinaryOpAssign(MulAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 200:2: binary_op_assign -> DIV_ASSIGN: ...
	DivAssignToBinaryOpAssign(DivAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 201:2: binary_op_assign -> MOD_ASSIGN: ...
	ModAssignToBinaryOpAssign(ModAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 202:2: binary_op_assign -> BIT_NEG_ASSIGN: ...
	BitNegAssignToBinaryOpAssign(BitNegAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 203:2: binary_op_assign -> BIT_AND_ASSIGN: ...
	BitAndAssignToBinaryOpAssign(BitAndAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 204:2: binary_op_assign -> BIT_OR_ASSIGN: ...
	BitOrAssignToBinaryOpAssign(BitOrAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 205:2: binary_op_assign -> BIT_XOR_ASSIGN: ...
	BitXorAssignToBinaryOpAssign(BitXorAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 206:2: binary_op_assign -> BIT_LSHIFT_ASSIGN: ...
	BitLshiftAssignToBinaryOpAssign(BitLshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 207:2: binary_op_assign -> BIT_RSHIFT_ASSIGN: ...
	BitRshiftAssignToBinaryOpAssign(BitRshiftAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 215:20: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ *GenericSymbol, Less_ *GenericSymbol, Identifier_ *GenericSymbol, Greater_ *GenericSymbol, StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 222:2: jump_statement -> ...
	ToJumpStatement(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 225:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 226:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 227:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 230:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 231:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 234:2: expressions -> expression: ...
	ExpressionToExpressions(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 235:2: expressions -> add: ...
	AddToExpressions(Expressions_ *GenericSymbol, Comma_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 238:2: optional_expressions -> expressions: ...
	ExpressionsToOptionalExpressions(Expressions_ *GenericSymbol) (*GenericSymbol, error)

	// 239:2: optional_expressions -> nil: ...
	NilToOptionalExpressions() (*GenericSymbol, error)

	// 245:13: call_expr -> ...
	ToCallExpr(AccessExpr_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalArguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 248:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ *GenericSymbol, OptionalGenericArguments_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 249:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (*GenericSymbol, error)

	// 252:2: optional_generic_arguments -> generic_arguments: ...
	GenericArgumentsToOptionalGenericArguments(GenericArguments_ *GenericSymbol) (*GenericSymbol, error)

	// 253:2: optional_generic_arguments -> nil: ...
	NilToOptionalGenericArguments() (*GenericSymbol, error)

	// 257:2: generic_arguments -> value_type: ...
	ValueTypeToGenericArguments(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 258:2: generic_arguments -> add: ...
	AddToGenericArguments(GenericArguments_ *GenericSymbol, Comma_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 261:2: optional_arguments -> arguments: ...
	ArgumentsToOptionalArguments(Arguments_ *GenericSymbol) (*GenericSymbol, error)

	// 262:2: optional_arguments -> nil: ...
	NilToOptionalArguments() (*GenericSymbol, error)

	// 265:2: arguments -> argument: ...
	ArgumentToArguments(Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 266:2: arguments -> add: ...
	AddToArguments(Arguments_ *GenericSymbol, Comma_ *GenericSymbol, Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 269:2: argument -> positional: ...
	PositionalToArgument(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 270:2: argument -> named: ...
	NamedToArgument(Identifier_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 271:2: argument -> colon_expressions: ...
	ColonExpressionsToArgument(ColonExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 274:2: colon_expressions -> pair: ...
	PairToColonExpressions(OptionalExpression_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_2 *GenericSymbol) (*GenericSymbol, error)

	// 275:2: colon_expressions -> add: ...
	AddToColonExpressions(ColonExpressions_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_ *GenericSymbol) (*GenericSymbol, error)

	// 278:2: optional_expression -> expression: ...
	ExpressionToOptionalExpression(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 279:2: optional_expression -> nil: ...
	NilToOptionalExpression() (*GenericSymbol, error)

	// 289:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 290:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 291:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 292:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 293:2: atom_expr -> initialize_expr: ...
	InitializeExprToAtomExpr(InitializeExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 294:2: atom_expr -> LEX_ERROR: ...
	LexErrorToAtomExpr(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 297:2: literal -> TRUE: ...
	TrueToLiteral(True_ *GenericSymbol) (*GenericSymbol, error)

	// 298:2: literal -> FALSE: ...
	FalseToLiteral(False_ *GenericSymbol) (*GenericSymbol, error)

	// 299:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 300:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 301:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 302:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 305:2: initialize_expr -> explicit: ...
	ExplicitToInitializeExpr(InitializableType_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 306:2: initialize_expr -> implicit_struct: ...
	ImplicitStructToInitializeExpr(Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 309:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 310:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 311:2: access_expr -> call_expr: ...
	CallExprToAccessExpr(CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 312:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Argument_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 315:2: postfix_unary_expr -> access_expr: ...
	AccessExprToPostfixUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 316:2: postfix_unary_expr -> question: ...
	QuestionToPostfixUnaryExpr(AccessExpr_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 319:2: prefix_unary_op -> NOT: ...
	NotToPrefixUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 320:2: prefix_unary_op -> BIT_NEG: ...
	BitNegToPrefixUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 321:2: prefix_unary_op -> SUB: ...
	SubToPrefixUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 324:2: prefix_unary_op -> MUL: ...
	MulToPrefixUnaryOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 327:2: prefix_unary_op -> BIT_AND: ...
	BitAndToPrefixUnaryOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 330:2: prefix_unary_expr -> postfix_unary_expr: ...
	PostfixUnaryExprToPrefixUnaryExpr(PostfixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 331:2: prefix_unary_expr -> prefix_op: ...
	PrefixOpToPrefixUnaryExpr(PrefixUnaryOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 334:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 335:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 336:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 337:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 338:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 339:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 342:2: mul_expr -> prefix_unary_expr: ...
	PrefixUnaryExprToMulExpr(PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 343:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 346:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 347:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 348:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 349:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 352:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 353:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 356:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 357:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 358:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 359:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 360:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 361:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 364:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 365:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 368:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 369:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 372:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 373:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 382:2: initializable_type -> explicit_struct_def: ...
	ExplicitStructDefToInitializableType(ExplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 383:2: initializable_type -> slice: ...
	SliceToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 384:2: initializable_type -> array: ...
	ArrayToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Comma_ *GenericSymbol, IntegerLiteral_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 385:2: initializable_type -> map: ...
	MapToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Colon_ *GenericSymbol, ValueType_2 *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 388:2: atom_type -> initializable_type: ...
	InitializableTypeToAtomType(InitializableType_ *GenericSymbol) (*GenericSymbol, error)

	// 389:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 390:2: atom_type -> extern_named: ...
	ExternNamedToAtomType(Identifier_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_2 *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 391:2: atom_type -> inferred: ...
	InferredToAtomType(Dot_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 392:2: atom_type -> implicit_struct_def: ...
	ImplicitStructDefToAtomType(ImplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 393:2: atom_type -> explicit_enum_def: ...
	ExplicitEnumDefToAtomType(ExplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 394:2: atom_type -> implicit_enum_def: ...
	ImplicitEnumDefToAtomType(ImplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 395:2: atom_type -> trait_def: ...
	TraitDefToAtomType(TraitDef_ *GenericSymbol) (*GenericSymbol, error)

	// 396:2: atom_type -> func_type: ...
	FuncTypeToAtomType(FuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 397:2: atom_type -> LEX_ERROR: ...
	LexErrorToAtomType(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 403:2: returnable_type -> atom_type: ...
	AtomTypeToReturnableType(AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 404:2: returnable_type -> optional: ...
	OptionalToReturnableType(Question_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 405:2: returnable_type -> result: ...
	ResultToReturnableType(Exclaim_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 406:2: returnable_type -> reference: ...
	ReferenceToReturnableType(BitAnd_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 407:2: returnable_type -> public_methods_trait: ...
	PublicMethodsTraitToReturnableType(BitNeg_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 408:2: returnable_type -> public_trait: ...
	PublicTraitToReturnableType(TildeTilde_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 413:2: value_type -> returnable_type: ...
	ReturnableTypeToValueType(ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 414:2: value_type -> trait_intersect: ...
	TraitIntersectToValueType(ValueType_ *GenericSymbol, Mul_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 415:2: value_type -> trait_union: ...
	TraitUnionToValueType(ValueType_ *GenericSymbol, Add_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 416:2: value_type -> trait_difference: ...
	TraitDifferenceToValueType(ValueType_ *GenericSymbol, Sub_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 419:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 420:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol, Implements_ *GenericSymbol, ValueType_2 *GenericSymbol) (*GenericSymbol, error)

	// 421:2: type_def -> alias: ...
	AliasToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, Equal_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 429:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 430:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 433:2: generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToGenericParameterDefs(GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 434:2: generic_parameter_defs -> add: ...
	AddToGenericParameterDefs(GenericParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 437:2: optional_generic_parameter_defs -> generic_parameter_defs: ...
	GenericParameterDefsToOptionalGenericParameterDefs(GenericParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 438:2: optional_generic_parameter_defs -> nil: ...
	NilToOptionalGenericParameterDefs() (*GenericSymbol, error)

	// 441:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ *GenericSymbol, OptionalGenericParameterDefs_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 442:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (*GenericSymbol, error)

	// 449:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 450:2: field_def -> implicit: ...
	ImplicitToFieldDef(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 451:2: field_def -> unsafe_statement: ...
	UnsafeStatementToFieldDef(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 454:2: implicit_field_defs -> field_def: ...
	FieldDefToImplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 455:2: implicit_field_defs -> add: ...
	AddToImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 458:2: optional_implicit_field_defs -> implicit_field_defs: ...
	ImplicitFieldDefsToOptionalImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 459:2: optional_implicit_field_defs -> nil: ...
	NilToOptionalImplicitFieldDefs() (*GenericSymbol, error)

	// 461:23: implicit_struct_def -> ...
	ToImplicitStructDef(Lparen_ *GenericSymbol, OptionalImplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 464:2: explicit_field_defs -> field_def: ...
	FieldDefToExplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 465:2: explicit_field_defs -> implicit: ...
	ImplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Newlines_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 466:2: explicit_field_defs -> explicit: ...
	ExplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 469:2: optional_explicit_field_defs -> explicit_field_defs: ...
	ExplicitFieldDefsToOptionalExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 470:2: optional_explicit_field_defs -> nil: ...
	NilToOptionalExplicitFieldDefs() (*GenericSymbol, error)

	// 472:23: explicit_struct_def -> ...
	ToExplicitStructDef(Struct_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalExplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 482:2: enum_value_def -> field_def: ...
	FieldDefToEnumValueDef(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 483:2: enum_value_def -> default: ...
	DefaultToEnumValueDef(FieldDef_ *GenericSymbol, Assign_ *GenericSymbol, Default_ *GenericSymbol) (*GenericSymbol, error)

	// 486:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 487:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 489:21: implicit_enum_def -> ...
	ToImplicitEnumDef(Lparen_ *GenericSymbol, ImplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 492:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 493:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 494:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 495:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 497:21: explicit_enum_def -> ...
	ToExplicitEnumDef(Enum_ *GenericSymbol, Lparen_ *GenericSymbol, ExplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 504:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 505:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ *GenericSymbol) (*GenericSymbol, error)

	// 508:2: trait_properties -> trait_property: ...
	TraitPropertyToTraitProperties(TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 509:2: trait_properties -> implicit: ...
	ImplicitToTraitProperties(TraitProperties_ *GenericSymbol, Newlines_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 510:2: trait_properties -> explicit: ...
	ExplicitToTraitProperties(TraitProperties_ *GenericSymbol, Comma_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 513:2: optional_trait_properties -> trait_properties: ...
	TraitPropertiesToOptionalTraitProperties(TraitProperties_ *GenericSymbol) (*GenericSymbol, error)

	// 514:2: optional_trait_properties -> nil: ...
	NilToOptionalTraitProperties() (*GenericSymbol, error)

	// 516:13: trait_def -> ...
	ToTraitDef(Trait_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalTraitProperties_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 524:2: return_type -> returnable_type: ...
	ReturnableTypeToReturnType(ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 525:2: return_type -> nil: ...
	NilToReturnType() (*GenericSymbol, error)

	// 528:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 529:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 530:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 531:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 534:2: parameter_decls -> parameter_decl: ...
	ParameterDeclToParameterDecls(ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 535:2: parameter_decls -> add: ...
	AddToParameterDecls(ParameterDecls_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 538:2: optional_parameter_decls -> parameter_decls: ...
	ParameterDeclsToOptionalParameterDecls(ParameterDecls_ *GenericSymbol) (*GenericSymbol, error)

	// 539:2: optional_parameter_decls -> nil: ...
	NilToOptionalParameterDecls() (*GenericSymbol, error)

	// 541:13: func_type -> ...
	ToFuncType(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 549:20: method_signature -> ...
	ToMethodSignature(Func_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 552:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 553:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 556:2: parameter_defs -> parameter_def: ...
	ParameterDefToParameterDefs(ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 557:2: parameter_defs -> add: ...
	AddToParameterDefs(ParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 560:2: optional_parameter_defs -> parameter_defs: ...
	ParameterDefsToOptionalParameterDefs(ParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 561:2: optional_parameter_defs -> nil: ...
	NilToOptionalParameterDefs() (*GenericSymbol, error)

	// 564:2: optional_receiver -> receiver: ...
	ReceiverToOptionalReceiver(Lparen_ *GenericSymbol, ParameterDef_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 565:2: optional_receiver -> nil: ...
	NilToOptionalReceiver() (*GenericSymbol, error)

	// 568:2: named_func_def -> ...
	ToNamedFuncDef(Func_ *GenericSymbol, OptionalReceiver_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 571:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 578:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 579:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, PackageStatements_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 581:26: package_statement_body -> ...
	ToPackageStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 584:2: package_statement -> implicit: ...
	ImplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 585:2: package_statement -> explicit: ...
	ExplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 588:2: package_statements -> empty_list: ...
	EmptyListToPackageStatements() (*GenericSymbol, error)

	// 589:2: package_statements -> add: ...
	AddToPackageStatements(PackageStatements_ *GenericSymbol, PackageStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 593:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 594:2: lex_internal_tokens -> COMMENT: ...
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
	case CaseBranchesType:
		return "case_branches"
	case CaseBranchType:
		return "case_branch"
	case MatchPatternType:
		return "match_pattern"
	case OptionalDefaultBranchType:
		return "optional_default_branch"
	case DefaultBranchType:
		return "default_branch"
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
	case ReturnableTypeType:
		return "returnable_type"
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

	SourceType                       = SymbolId(339)
	OptionalDefinitionsType          = SymbolId(340)
	OptionalNewlinesType             = SymbolId(341)
	DefinitionsType                  = SymbolId(342)
	DefinitionType                   = SymbolId(343)
	ExpressionType                   = SymbolId(344)
	OptionalLabelDeclType            = SymbolId(345)
	IfExprType                       = SymbolId(346)
	SwitchExprType                   = SymbolId(347)
	CaseBranchesType                 = SymbolId(348)
	CaseBranchType                   = SymbolId(349)
	MatchPatternType                 = SymbolId(350)
	OptionalDefaultBranchType        = SymbolId(351)
	DefaultBranchType                = SymbolId(352)
	LoopExprType                     = SymbolId(353)
	OptionalSequenceExprType         = SymbolId(354)
	SequenceExprType                 = SymbolId(355)
	BlockExprType                    = SymbolId(356)
	BlockBodyType                    = SymbolId(357)
	StatementsType                   = SymbolId(358)
	StatementType                    = SymbolId(359)
	StatementBodyType                = SymbolId(360)
	UnaryOpAssignType                = SymbolId(361)
	BinaryOpAssignType               = SymbolId(362)
	UnsafeStatementType              = SymbolId(363)
	JumpStatementType                = SymbolId(364)
	JumpTypeType                     = SymbolId(365)
	OptionalJumpLabelType            = SymbolId(366)
	ExpressionsType                  = SymbolId(367)
	OptionalExpressionsType          = SymbolId(368)
	CallExprType                     = SymbolId(369)
	OptionalGenericBindingType       = SymbolId(370)
	OptionalGenericArgumentsType     = SymbolId(371)
	GenericArgumentsType             = SymbolId(372)
	OptionalArgumentsType            = SymbolId(373)
	ArgumentsType                    = SymbolId(374)
	ArgumentType                     = SymbolId(375)
	ColonExpressionsType             = SymbolId(376)
	OptionalExpressionType           = SymbolId(377)
	AtomExprType                     = SymbolId(378)
	LiteralType                      = SymbolId(379)
	InitializeExprType               = SymbolId(380)
	AccessExprType                   = SymbolId(381)
	PostfixUnaryExprType             = SymbolId(382)
	PrefixUnaryOpType                = SymbolId(383)
	PrefixUnaryExprType              = SymbolId(384)
	MulOpType                        = SymbolId(385)
	MulExprType                      = SymbolId(386)
	AddOpType                        = SymbolId(387)
	AddExprType                      = SymbolId(388)
	CmpOpType                        = SymbolId(389)
	CmpExprType                      = SymbolId(390)
	AndExprType                      = SymbolId(391)
	OrExprType                       = SymbolId(392)
	InitializableTypeType            = SymbolId(393)
	AtomTypeType                     = SymbolId(394)
	ReturnableTypeType               = SymbolId(395)
	ValueTypeType                    = SymbolId(396)
	TypeDefType                      = SymbolId(397)
	GenericParameterDefType          = SymbolId(398)
	GenericParameterDefsType         = SymbolId(399)
	OptionalGenericParameterDefsType = SymbolId(400)
	OptionalGenericParametersType    = SymbolId(401)
	FieldDefType                     = SymbolId(402)
	ImplicitFieldDefsType            = SymbolId(403)
	OptionalImplicitFieldDefsType    = SymbolId(404)
	ImplicitStructDefType            = SymbolId(405)
	ExplicitFieldDefsType            = SymbolId(406)
	OptionalExplicitFieldDefsType    = SymbolId(407)
	ExplicitStructDefType            = SymbolId(408)
	EnumValueDefType                 = SymbolId(409)
	ImplicitEnumValueDefsType        = SymbolId(410)
	ImplicitEnumDefType              = SymbolId(411)
	ExplicitEnumValueDefsType        = SymbolId(412)
	ExplicitEnumDefType              = SymbolId(413)
	TraitPropertyType                = SymbolId(414)
	TraitPropertiesType              = SymbolId(415)
	OptionalTraitPropertiesType      = SymbolId(416)
	TraitDefType                     = SymbolId(417)
	ReturnTypeType                   = SymbolId(418)
	ParameterDeclType                = SymbolId(419)
	ParameterDeclsType               = SymbolId(420)
	OptionalParameterDeclsType       = SymbolId(421)
	FuncTypeType                     = SymbolId(422)
	MethodSignatureType              = SymbolId(423)
	ParameterDefType                 = SymbolId(424)
	ParameterDefsType                = SymbolId(425)
	OptionalParameterDefsType        = SymbolId(426)
	OptionalReceiverType             = SymbolId(427)
	NamedFuncDefType                 = SymbolId(428)
	AnonymousFuncExprType            = SymbolId(429)
	PackageDefType                   = SymbolId(430)
	PackageStatementBodyType         = SymbolId(431)
	PackageStatementType             = SymbolId(432)
	PackageStatementsType            = SymbolId(433)
	LexInternalTokensType            = SymbolId(434)
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
	_ReduceCaseBranchToCaseBranches                           = _ReduceType(23)
	_ReduceAddToCaseBranches                                  = _ReduceType(24)
	_ReduceToCaseBranch                                       = _ReduceType(25)
	_ReduceToMatchPattern                                     = _ReduceType(26)
	_ReduceDefaultBranchToOptionalDefaultBranch               = _ReduceType(27)
	_ReduceNilToOptionalDefaultBranch                         = _ReduceType(28)
	_ReduceToDefaultBranch                                    = _ReduceType(29)
	_ReduceInfiniteToLoopExpr                                 = _ReduceType(30)
	_ReduceDoWhileToLoopExpr                                  = _ReduceType(31)
	_ReduceWhileToLoopExpr                                    = _ReduceType(32)
	_ReduceIteratorToLoopExpr                                 = _ReduceType(33)
	_ReduceForToLoopExpr                                      = _ReduceType(34)
	_ReduceSequenceExprToOptionalSequenceExpr                 = _ReduceType(35)
	_ReduceNilToOptionalSequenceExpr                          = _ReduceType(36)
	_ReduceToSequenceExpr                                     = _ReduceType(37)
	_ReduceToBlockExpr                                        = _ReduceType(38)
	_ReduceToBlockBody                                        = _ReduceType(39)
	_ReduceEmptyListToStatements                              = _ReduceType(40)
	_ReduceAddToStatements                                    = _ReduceType(41)
	_ReduceImplicitToStatement                                = _ReduceType(42)
	_ReduceExplicitToStatement                                = _ReduceType(43)
	_ReduceUnsafeStatementToStatementBody                     = _ReduceType(44)
	_ReduceExpressionOrImplicitStructToStatementBody          = _ReduceType(45)
	_ReduceAsyncToStatementBody                               = _ReduceType(46)
	_ReduceDeferToStatementBody                               = _ReduceType(47)
	_ReduceJumpStatementToStatementBody                       = _ReduceType(48)
	_ReduceUnaryOpAssignStatementToStatementBody              = _ReduceType(49)
	_ReduceBinaryOpAssignStatementToStatementBody             = _ReduceType(50)
	_ReduceAddOneAssignToUnaryOpAssign                        = _ReduceType(51)
	_ReduceSubOneAssignToUnaryOpAssign                        = _ReduceType(52)
	_ReduceAddAssignToBinaryOpAssign                          = _ReduceType(53)
	_ReduceSubAssignToBinaryOpAssign                          = _ReduceType(54)
	_ReduceMulAssignToBinaryOpAssign                          = _ReduceType(55)
	_ReduceDivAssignToBinaryOpAssign                          = _ReduceType(56)
	_ReduceModAssignToBinaryOpAssign                          = _ReduceType(57)
	_ReduceBitNegAssignToBinaryOpAssign                       = _ReduceType(58)
	_ReduceBitAndAssignToBinaryOpAssign                       = _ReduceType(59)
	_ReduceBitOrAssignToBinaryOpAssign                        = _ReduceType(60)
	_ReduceBitXorAssignToBinaryOpAssign                       = _ReduceType(61)
	_ReduceBitLshiftAssignToBinaryOpAssign                    = _ReduceType(62)
	_ReduceBitRshiftAssignToBinaryOpAssign                    = _ReduceType(63)
	_ReduceToUnsafeStatement                                  = _ReduceType(64)
	_ReduceToJumpStatement                                    = _ReduceType(65)
	_ReduceReturnToJumpType                                   = _ReduceType(66)
	_ReduceBreakToJumpType                                    = _ReduceType(67)
	_ReduceContinueToJumpType                                 = _ReduceType(68)
	_ReduceJumpLabelToOptionalJumpLabel                       = _ReduceType(69)
	_ReduceUnlabelledToOptionalJumpLabel                      = _ReduceType(70)
	_ReduceExpressionToExpressions                            = _ReduceType(71)
	_ReduceAddToExpressions                                   = _ReduceType(72)
	_ReduceExpressionsToOptionalExpressions                   = _ReduceType(73)
	_ReduceNilToOptionalExpressions                           = _ReduceType(74)
	_ReduceToCallExpr                                         = _ReduceType(75)
	_ReduceBindingToOptionalGenericBinding                    = _ReduceType(76)
	_ReduceNilToOptionalGenericBinding                        = _ReduceType(77)
	_ReduceGenericArgumentsToOptionalGenericArguments         = _ReduceType(78)
	_ReduceNilToOptionalGenericArguments                      = _ReduceType(79)
	_ReduceValueTypeToGenericArguments                        = _ReduceType(80)
	_ReduceAddToGenericArguments                              = _ReduceType(81)
	_ReduceArgumentsToOptionalArguments                       = _ReduceType(82)
	_ReduceNilToOptionalArguments                             = _ReduceType(83)
	_ReduceArgumentToArguments                                = _ReduceType(84)
	_ReduceAddToArguments                                     = _ReduceType(85)
	_ReducePositionalToArgument                               = _ReduceType(86)
	_ReduceNamedToArgument                                    = _ReduceType(87)
	_ReduceColonExpressionsToArgument                         = _ReduceType(88)
	_ReducePairToColonExpressions                             = _ReduceType(89)
	_ReduceAddToColonExpressions                              = _ReduceType(90)
	_ReduceExpressionToOptionalExpression                     = _ReduceType(91)
	_ReduceNilToOptionalExpression                            = _ReduceType(92)
	_ReduceLiteralToAtomExpr                                  = _ReduceType(93)
	_ReduceIdentifierToAtomExpr                               = _ReduceType(94)
	_ReduceBlockExprToAtomExpr                                = _ReduceType(95)
	_ReduceAnonymousFuncExprToAtomExpr                        = _ReduceType(96)
	_ReduceInitializeExprToAtomExpr                           = _ReduceType(97)
	_ReduceLexErrorToAtomExpr                                 = _ReduceType(98)
	_ReduceTrueToLiteral                                      = _ReduceType(99)
	_ReduceFalseToLiteral                                     = _ReduceType(100)
	_ReduceIntegerLiteralToLiteral                            = _ReduceType(101)
	_ReduceFloatLiteralToLiteral                              = _ReduceType(102)
	_ReduceRuneLiteralToLiteral                               = _ReduceType(103)
	_ReduceStringLiteralToLiteral                             = _ReduceType(104)
	_ReduceExplicitToInitializeExpr                           = _ReduceType(105)
	_ReduceImplicitStructToInitializeExpr                     = _ReduceType(106)
	_ReduceAtomExprToAccessExpr                               = _ReduceType(107)
	_ReduceAccessToAccessExpr                                 = _ReduceType(108)
	_ReduceCallExprToAccessExpr                               = _ReduceType(109)
	_ReduceIndexToAccessExpr                                  = _ReduceType(110)
	_ReduceAccessExprToPostfixUnaryExpr                       = _ReduceType(111)
	_ReduceQuestionToPostfixUnaryExpr                         = _ReduceType(112)
	_ReduceNotToPrefixUnaryOp                                 = _ReduceType(113)
	_ReduceBitNegToPrefixUnaryOp                              = _ReduceType(114)
	_ReduceSubToPrefixUnaryOp                                 = _ReduceType(115)
	_ReduceMulToPrefixUnaryOp                                 = _ReduceType(116)
	_ReduceBitAndToPrefixUnaryOp                              = _ReduceType(117)
	_ReducePostfixUnaryExprToPrefixUnaryExpr                  = _ReduceType(118)
	_ReducePrefixOpToPrefixUnaryExpr                          = _ReduceType(119)
	_ReduceMulToMulOp                                         = _ReduceType(120)
	_ReduceDivToMulOp                                         = _ReduceType(121)
	_ReduceModToMulOp                                         = _ReduceType(122)
	_ReduceBitAndToMulOp                                      = _ReduceType(123)
	_ReduceBitLshiftToMulOp                                   = _ReduceType(124)
	_ReduceBitRshiftToMulOp                                   = _ReduceType(125)
	_ReducePrefixUnaryExprToMulExpr                           = _ReduceType(126)
	_ReduceOpToMulExpr                                        = _ReduceType(127)
	_ReduceAddToAddOp                                         = _ReduceType(128)
	_ReduceSubToAddOp                                         = _ReduceType(129)
	_ReduceBitOrToAddOp                                       = _ReduceType(130)
	_ReduceBitXorToAddOp                                      = _ReduceType(131)
	_ReduceMulExprToAddExpr                                   = _ReduceType(132)
	_ReduceOpToAddExpr                                        = _ReduceType(133)
	_ReduceEqualToCmpOp                                       = _ReduceType(134)
	_ReduceNotEqualToCmpOp                                    = _ReduceType(135)
	_ReduceLessToCmpOp                                        = _ReduceType(136)
	_ReduceLessOrEqualToCmpOp                                 = _ReduceType(137)
	_ReduceGreaterToCmpOp                                     = _ReduceType(138)
	_ReduceGreaterOrEqualToCmpOp                              = _ReduceType(139)
	_ReduceAddExprToCmpExpr                                   = _ReduceType(140)
	_ReduceOpToCmpExpr                                        = _ReduceType(141)
	_ReduceCmpExprToAndExpr                                   = _ReduceType(142)
	_ReduceOpToAndExpr                                        = _ReduceType(143)
	_ReduceAndExprToOrExpr                                    = _ReduceType(144)
	_ReduceOpToOrExpr                                         = _ReduceType(145)
	_ReduceExplicitStructDefToInitializableType               = _ReduceType(146)
	_ReduceSliceToInitializableType                           = _ReduceType(147)
	_ReduceArrayToInitializableType                           = _ReduceType(148)
	_ReduceMapToInitializableType                             = _ReduceType(149)
	_ReduceInitializableTypeToAtomType                        = _ReduceType(150)
	_ReduceNamedToAtomType                                    = _ReduceType(151)
	_ReduceExternNamedToAtomType                              = _ReduceType(152)
	_ReduceInferredToAtomType                                 = _ReduceType(153)
	_ReduceImplicitStructDefToAtomType                        = _ReduceType(154)
	_ReduceExplicitEnumDefToAtomType                          = _ReduceType(155)
	_ReduceImplicitEnumDefToAtomType                          = _ReduceType(156)
	_ReduceTraitDefToAtomType                                 = _ReduceType(157)
	_ReduceFuncTypeToAtomType                                 = _ReduceType(158)
	_ReduceLexErrorToAtomType                                 = _ReduceType(159)
	_ReduceAtomTypeToReturnableType                           = _ReduceType(160)
	_ReduceOptionalToReturnableType                           = _ReduceType(161)
	_ReduceResultToReturnableType                             = _ReduceType(162)
	_ReduceReferenceToReturnableType                          = _ReduceType(163)
	_ReducePublicMethodsTraitToReturnableType                 = _ReduceType(164)
	_ReducePublicTraitToReturnableType                        = _ReduceType(165)
	_ReduceReturnableTypeToValueType                          = _ReduceType(166)
	_ReduceTraitIntersectToValueType                          = _ReduceType(167)
	_ReduceTraitUnionToValueType                              = _ReduceType(168)
	_ReduceTraitDifferenceToValueType                         = _ReduceType(169)
	_ReduceDefinitionToTypeDef                                = _ReduceType(170)
	_ReduceConstrainedDefToTypeDef                            = _ReduceType(171)
	_ReduceAliasToTypeDef                                     = _ReduceType(172)
	_ReduceUnconstrainedToGenericParameterDef                 = _ReduceType(173)
	_ReduceConstrainedToGenericParameterDef                   = _ReduceType(174)
	_ReduceGenericParameterDefToGenericParameterDefs          = _ReduceType(175)
	_ReduceAddToGenericParameterDefs                          = _ReduceType(176)
	_ReduceGenericParameterDefsToOptionalGenericParameterDefs = _ReduceType(177)
	_ReduceNilToOptionalGenericParameterDefs                  = _ReduceType(178)
	_ReduceGenericToOptionalGenericParameters                 = _ReduceType(179)
	_ReduceNilToOptionalGenericParameters                     = _ReduceType(180)
	_ReduceExplicitToFieldDef                                 = _ReduceType(181)
	_ReduceImplicitToFieldDef                                 = _ReduceType(182)
	_ReduceUnsafeStatementToFieldDef                          = _ReduceType(183)
	_ReduceFieldDefToImplicitFieldDefs                        = _ReduceType(184)
	_ReduceAddToImplicitFieldDefs                             = _ReduceType(185)
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefs       = _ReduceType(186)
	_ReduceNilToOptionalImplicitFieldDefs                     = _ReduceType(187)
	_ReduceToImplicitStructDef                                = _ReduceType(188)
	_ReduceFieldDefToExplicitFieldDefs                        = _ReduceType(189)
	_ReduceImplicitToExplicitFieldDefs                        = _ReduceType(190)
	_ReduceExplicitToExplicitFieldDefs                        = _ReduceType(191)
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefs       = _ReduceType(192)
	_ReduceNilToOptionalExplicitFieldDefs                     = _ReduceType(193)
	_ReduceToExplicitStructDef                                = _ReduceType(194)
	_ReduceFieldDefToEnumValueDef                             = _ReduceType(195)
	_ReduceDefaultToEnumValueDef                              = _ReduceType(196)
	_ReducePairToImplicitEnumValueDefs                        = _ReduceType(197)
	_ReduceAddToImplicitEnumValueDefs                         = _ReduceType(198)
	_ReduceToImplicitEnumDef                                  = _ReduceType(199)
	_ReduceExplicitPairToExplicitEnumValueDefs                = _ReduceType(200)
	_ReduceImplicitPairToExplicitEnumValueDefs                = _ReduceType(201)
	_ReduceExplicitAddToExplicitEnumValueDefs                 = _ReduceType(202)
	_ReduceImplicitAddToExplicitEnumValueDefs                 = _ReduceType(203)
	_ReduceToExplicitEnumDef                                  = _ReduceType(204)
	_ReduceFieldDefToTraitProperty                            = _ReduceType(205)
	_ReduceMethodSignatureToTraitProperty                     = _ReduceType(206)
	_ReduceTraitPropertyToTraitProperties                     = _ReduceType(207)
	_ReduceImplicitToTraitProperties                          = _ReduceType(208)
	_ReduceExplicitToTraitProperties                          = _ReduceType(209)
	_ReduceTraitPropertiesToOptionalTraitProperties           = _ReduceType(210)
	_ReduceNilToOptionalTraitProperties                       = _ReduceType(211)
	_ReduceToTraitDef                                         = _ReduceType(212)
	_ReduceReturnableTypeToReturnType                         = _ReduceType(213)
	_ReduceNilToReturnType                                    = _ReduceType(214)
	_ReduceArgToParameterDecl                                 = _ReduceType(215)
	_ReduceVarargToParameterDecl                              = _ReduceType(216)
	_ReduceUnamedToParameterDecl                              = _ReduceType(217)
	_ReduceUnnamedVarargToParameterDecl                       = _ReduceType(218)
	_ReduceParameterDeclToParameterDecls                      = _ReduceType(219)
	_ReduceAddToParameterDecls                                = _ReduceType(220)
	_ReduceParameterDeclsToOptionalParameterDecls             = _ReduceType(221)
	_ReduceNilToOptionalParameterDecls                        = _ReduceType(222)
	_ReduceToFuncType                                         = _ReduceType(223)
	_ReduceToMethodSignature                                  = _ReduceType(224)
	_ReduceArgToParameterDef                                  = _ReduceType(225)
	_ReduceVarargToParameterDef                               = _ReduceType(226)
	_ReduceParameterDefToParameterDefs                        = _ReduceType(227)
	_ReduceAddToParameterDefs                                 = _ReduceType(228)
	_ReduceParameterDefsToOptionalParameterDefs               = _ReduceType(229)
	_ReduceNilToOptionalParameterDefs                         = _ReduceType(230)
	_ReduceReceiverToOptionalReceiver                         = _ReduceType(231)
	_ReduceNilToOptionalReceiver                              = _ReduceType(232)
	_ReduceToNamedFuncDef                                     = _ReduceType(233)
	_ReduceToAnonymousFuncExpr                                = _ReduceType(234)
	_ReduceNoSpecToPackageDef                                 = _ReduceType(235)
	_ReduceWithSpecToPackageDef                               = _ReduceType(236)
	_ReduceToPackageStatementBody                             = _ReduceType(237)
	_ReduceImplicitToPackageStatement                         = _ReduceType(238)
	_ReduceExplicitToPackageStatement                         = _ReduceType(239)
	_ReduceEmptyListToPackageStatements                       = _ReduceType(240)
	_ReduceAddToPackageStatements                             = _ReduceType(241)
	_ReduceSpacesToLexInternalTokens                          = _ReduceType(242)
	_ReduceCommentToLexInternalTokens                         = _ReduceType(243)
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
	case _ReduceCaseBranchToCaseBranches:
		return "CaseBranchToCaseBranches"
	case _ReduceAddToCaseBranches:
		return "AddToCaseBranches"
	case _ReduceToCaseBranch:
		return "ToCaseBranch"
	case _ReduceToMatchPattern:
		return "ToMatchPattern"
	case _ReduceDefaultBranchToOptionalDefaultBranch:
		return "DefaultBranchToOptionalDefaultBranch"
	case _ReduceNilToOptionalDefaultBranch:
		return "NilToOptionalDefaultBranch"
	case _ReduceToDefaultBranch:
		return "ToDefaultBranch"
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
	case _ReduceExplicitEnumDefToAtomType:
		return "ExplicitEnumDefToAtomType"
	case _ReduceImplicitEnumDefToAtomType:
		return "ImplicitEnumDefToAtomType"
	case _ReduceTraitDefToAtomType:
		return "TraitDefToAtomType"
	case _ReduceFuncTypeToAtomType:
		return "FuncTypeToAtomType"
	case _ReduceLexErrorToAtomType:
		return "LexErrorToAtomType"
	case _ReduceAtomTypeToReturnableType:
		return "AtomTypeToReturnableType"
	case _ReduceOptionalToReturnableType:
		return "OptionalToReturnableType"
	case _ReduceResultToReturnableType:
		return "ResultToReturnableType"
	case _ReduceReferenceToReturnableType:
		return "ReferenceToReturnableType"
	case _ReducePublicMethodsTraitToReturnableType:
		return "PublicMethodsTraitToReturnableType"
	case _ReducePublicTraitToReturnableType:
		return "PublicTraitToReturnableType"
	case _ReduceReturnableTypeToValueType:
		return "ReturnableTypeToValueType"
	case _ReduceTraitIntersectToValueType:
		return "TraitIntersectToValueType"
	case _ReduceTraitUnionToValueType:
		return "TraitUnionToValueType"
	case _ReduceTraitDifferenceToValueType:
		return "TraitDifferenceToValueType"
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
	case _ReduceReturnableTypeToReturnType:
		return "ReturnableTypeToReturnType"
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
	_State387 = _StateId(387)
	_State388 = _StateId(388)
	_State389 = _StateId(389)
	_State390 = _StateId(390)
	_State391 = _StateId(391)
	_State392 = _StateId(392)
	_State393 = _StateId(393)
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
	case _ReduceCaseBranchToCaseBranches:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CaseBranchesType
		symbol.Generic_, err = reducer.CaseBranchToCaseBranches(args[0].Generic_)
	case _ReduceAddToCaseBranches:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CaseBranchesType
		symbol.Generic_, err = reducer.AddToCaseBranches(args[0].Generic_, args[1].Generic_)
	case _ReduceToCaseBranch:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CaseBranchType
		symbol.Generic_, err = reducer.ToCaseBranch(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceToMatchPattern:
		symbol.SymbolId_ = MatchPatternType
		symbol.Generic_, err = reducer.ToMatchPattern()
	case _ReduceDefaultBranchToOptionalDefaultBranch:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalDefaultBranchType
		symbol.Generic_, err = reducer.DefaultBranchToOptionalDefaultBranch(args[0].Generic_)
	case _ReduceNilToOptionalDefaultBranch:
		symbol.SymbolId_ = OptionalDefaultBranchType
		symbol.Generic_, err = reducer.NilToOptionalDefaultBranch()
	case _ReduceToDefaultBranch:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = DefaultBranchType
		symbol.Generic_, err = reducer.ToDefaultBranch(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	case _ReduceAtomTypeToReturnableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.Generic_, err = reducer.AtomTypeToReturnableType(args[0].Generic_)
	case _ReduceOptionalToReturnableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.Generic_, err = reducer.OptionalToReturnableType(args[0].Generic_, args[1].Generic_)
	case _ReduceResultToReturnableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.Generic_, err = reducer.ResultToReturnableType(args[0].Generic_, args[1].Generic_)
	case _ReduceReferenceToReturnableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.Generic_, err = reducer.ReferenceToReturnableType(args[0].Generic_, args[1].Generic_)
	case _ReducePublicMethodsTraitToReturnableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.Generic_, err = reducer.PublicMethodsTraitToReturnableType(args[0].Generic_, args[1].Generic_)
	case _ReducePublicTraitToReturnableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.Generic_, err = reducer.PublicTraitToReturnableType(args[0].Generic_, args[1].Generic_)
	case _ReduceReturnableTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.ReturnableTypeToValueType(args[0].Generic_)
	case _ReduceTraitIntersectToValueType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.TraitIntersectToValueType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceTraitUnionToValueType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.TraitUnionToValueType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceTraitDifferenceToValueType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ValueTypeType
		symbol.Generic_, err = reducer.TraitDifferenceToValueType(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	case _ReduceReturnableTypeToReturnType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnTypeType
		symbol.Generic_, err = reducer.ReturnableTypeToReturnType(args[0].Generic_)
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
	_GotoState387Action                                             = &_Action{_ShiftAction, _State387, 0}
	_GotoState388Action                                             = &_Action{_ShiftAction, _State388, 0}
	_GotoState389Action                                             = &_Action{_ShiftAction, _State389, 0}
	_GotoState390Action                                             = &_Action{_ShiftAction, _State390, 0}
	_GotoState391Action                                             = &_Action{_ShiftAction, _State391, 0}
	_GotoState392Action                                             = &_Action{_ShiftAction, _State392, 0}
	_GotoState393Action                                             = &_Action{_ShiftAction, _State393, 0}
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
	_ReduceCaseBranchToCaseBranchesAction                           = &_Action{_ReduceAction, 0, _ReduceCaseBranchToCaseBranches}
	_ReduceAddToCaseBranchesAction                                  = &_Action{_ReduceAction, 0, _ReduceAddToCaseBranches}
	_ReduceToCaseBranchAction                                       = &_Action{_ReduceAction, 0, _ReduceToCaseBranch}
	_ReduceToMatchPatternAction                                     = &_Action{_ReduceAction, 0, _ReduceToMatchPattern}
	_ReduceDefaultBranchToOptionalDefaultBranchAction               = &_Action{_ReduceAction, 0, _ReduceDefaultBranchToOptionalDefaultBranch}
	_ReduceNilToOptionalDefaultBranchAction                         = &_Action{_ReduceAction, 0, _ReduceNilToOptionalDefaultBranch}
	_ReduceToDefaultBranchAction                                    = &_Action{_ReduceAction, 0, _ReduceToDefaultBranch}
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
	_ReduceSliceToInitializableTypeAction                           = &_Action{_ReduceAction, 0, _ReduceSliceToInitializableType}
	_ReduceArrayToInitializableTypeAction                           = &_Action{_ReduceAction, 0, _ReduceArrayToInitializableType}
	_ReduceMapToInitializableTypeAction                             = &_Action{_ReduceAction, 0, _ReduceMapToInitializableType}
	_ReduceInitializableTypeToAtomTypeAction                        = &_Action{_ReduceAction, 0, _ReduceInitializableTypeToAtomType}
	_ReduceNamedToAtomTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceNamedToAtomType}
	_ReduceExternNamedToAtomTypeAction                              = &_Action{_ReduceAction, 0, _ReduceExternNamedToAtomType}
	_ReduceInferredToAtomTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceInferredToAtomType}
	_ReduceImplicitStructDefToAtomTypeAction                        = &_Action{_ReduceAction, 0, _ReduceImplicitStructDefToAtomType}
	_ReduceExplicitEnumDefToAtomTypeAction                          = &_Action{_ReduceAction, 0, _ReduceExplicitEnumDefToAtomType}
	_ReduceImplicitEnumDefToAtomTypeAction                          = &_Action{_ReduceAction, 0, _ReduceImplicitEnumDefToAtomType}
	_ReduceTraitDefToAtomTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceTraitDefToAtomType}
	_ReduceFuncTypeToAtomTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceFuncTypeToAtomType}
	_ReduceLexErrorToAtomTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceLexErrorToAtomType}
	_ReduceAtomTypeToReturnableTypeAction                           = &_Action{_ReduceAction, 0, _ReduceAtomTypeToReturnableType}
	_ReduceOptionalToReturnableTypeAction                           = &_Action{_ReduceAction, 0, _ReduceOptionalToReturnableType}
	_ReduceResultToReturnableTypeAction                             = &_Action{_ReduceAction, 0, _ReduceResultToReturnableType}
	_ReduceReferenceToReturnableTypeAction                          = &_Action{_ReduceAction, 0, _ReduceReferenceToReturnableType}
	_ReducePublicMethodsTraitToReturnableTypeAction                 = &_Action{_ReduceAction, 0, _ReducePublicMethodsTraitToReturnableType}
	_ReducePublicTraitToReturnableTypeAction                        = &_Action{_ReduceAction, 0, _ReducePublicTraitToReturnableType}
	_ReduceReturnableTypeToValueTypeAction                          = &_Action{_ReduceAction, 0, _ReduceReturnableTypeToValueType}
	_ReduceTraitIntersectToValueTypeAction                          = &_Action{_ReduceAction, 0, _ReduceTraitIntersectToValueType}
	_ReduceTraitUnionToValueTypeAction                              = &_Action{_ReduceAction, 0, _ReduceTraitUnionToValueType}
	_ReduceTraitDifferenceToValueTypeAction                         = &_Action{_ReduceAction, 0, _ReduceTraitDifferenceToValueType}
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
	_ReduceReturnableTypeToReturnTypeAction                         = &_Action{_ReduceAction, 0, _ReduceReturnableTypeToReturnType}
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
	{_State5, RuneLiteralToken}:                   _GotoState32Action,
	{_State5, StringLiteralToken}:                 _GotoState33Action,
	{_State5, IdentifierToken}:                    _GotoState24Action,
	{_State5, TrueToken}:                          _GotoState36Action,
	{_State5, FalseToken}:                         _GotoState21Action,
	{_State5, StructToken}:                        _GotoState34Action,
	{_State5, FuncToken}:                          _GotoState23Action,
	{_State5, LabelDeclToken}:                     _GotoState26Action,
	{_State5, LparenToken}:                        _GotoState29Action,
	{_State5, LbracketToken}:                      _GotoState27Action,
	{_State5, NotToken}:                           _GotoState31Action,
	{_State5, SubToken}:                           _GotoState35Action,
	{_State5, MulToken}:                           _GotoState30Action,
	{_State5, BitNegToken}:                        _GotoState20Action,
	{_State5, BitAndToken}:                        _GotoState19Action,
	{_State5, LexErrorToken}:                      _GotoState28Action,
	{_State5, ExpressionType}:                     _GotoState11Action,
	{_State5, OptionalLabelDeclType}:              _GotoState50Action,
	{_State5, SequenceExprType}:                   _GotoState55Action,
	{_State5, BlockExprType}:                      _GotoState42Action,
	{_State5, CallExprType}:                       _GotoState43Action,
	{_State5, AtomExprType}:                       _GotoState41Action,
	{_State5, LiteralType}:                        _GotoState48Action,
	{_State5, InitializeExprType}:                 _GotoState47Action,
	{_State5, AccessExprType}:                     _GotoState37Action,
	{_State5, PostfixUnaryExprType}:               _GotoState52Action,
	{_State5, PrefixUnaryOpType}:                  _GotoState54Action,
	{_State5, PrefixUnaryExprType}:                _GotoState53Action,
	{_State5, MulExprType}:                        _GotoState49Action,
	{_State5, AddExprType}:                        _GotoState38Action,
	{_State5, CmpExprType}:                        _GotoState44Action,
	{_State5, AndExprType}:                        _GotoState39Action,
	{_State5, OrExprType}:                         _GotoState51Action,
	{_State5, InitializableTypeType}:              _GotoState46Action,
	{_State5, ExplicitStructDefType}:              _GotoState45Action,
	{_State5, AnonymousFuncExprType}:              _GotoState40Action,
	{_State6, SpacesToken}:                        _GotoState57Action,
	{_State6, CommentToken}:                       _GotoState56Action,
	{_State6, LexInternalTokensType}:              _GotoState12Action,
	{_State15, PackageToken}:                      _GotoState16Action,
	{_State15, UnsafeToken}:                       _GotoState58Action,
	{_State15, TypeToken}:                         _GotoState17Action,
	{_State15, FuncToken}:                         _GotoState18Action,
	{_State15, DefinitionsType}:                   _GotoState60Action,
	{_State15, DefinitionType}:                    _GotoState59Action,
	{_State15, UnsafeStatementType}:               _GotoState64Action,
	{_State15, TypeDefType}:                       _GotoState63Action,
	{_State15, NamedFuncDefType}:                  _GotoState61Action,
	{_State15, PackageDefType}:                    _GotoState62Action,
	{_State16, IdentifierToken}:                   _GotoState65Action,
	{_State17, IdentifierToken}:                   _GotoState66Action,
	{_State18, LparenToken}:                       _GotoState67Action,
	{_State18, OptionalReceiverType}:              _GotoState68Action,
	{_State23, LparenToken}:                       _GotoState69Action,
	{_State27, IdentifierToken}:                   _GotoState76Action,
	{_State27, StructToken}:                       _GotoState34Action,
	{_State27, EnumToken}:                         _GotoState73Action,
	{_State27, TraitToken}:                        _GotoState81Action,
	{_State27, FuncToken}:                         _GotoState75Action,
	{_State27, LparenToken}:                       _GotoState78Action,
	{_State27, LbracketToken}:                     _GotoState27Action,
	{_State27, DotToken}:                          _GotoState72Action,
	{_State27, QuestionToken}:                     _GotoState79Action,
	{_State27, ExclaimToken}:                      _GotoState74Action,
	{_State27, TildeTildeToken}:                   _GotoState80Action,
	{_State27, BitNegToken}:                       _GotoState71Action,
	{_State27, BitAndToken}:                       _GotoState70Action,
	{_State27, LexErrorToken}:                     _GotoState77Action,
	{_State27, InitializableTypeType}:             _GotoState87Action,
	{_State27, AtomTypeType}:                      _GotoState82Action,
	{_State27, ReturnableTypeType}:                _GotoState88Action,
	{_State27, ValueTypeType}:                     _GotoState90Action,
	{_State27, ImplicitStructDefType}:             _GotoState86Action,
	{_State27, ExplicitStructDefType}:             _GotoState45Action,
	{_State27, ImplicitEnumDefType}:               _GotoState85Action,
	{_State27, ExplicitEnumDefType}:               _GotoState83Action,
	{_State27, TraitDefType}:                      _GotoState89Action,
	{_State27, FuncTypeType}:                      _GotoState84Action,
	{_State29, IntegerLiteralToken}:               _GotoState25Action,
	{_State29, FloatLiteralToken}:                 _GotoState22Action,
	{_State29, RuneLiteralToken}:                  _GotoState32Action,
	{_State29, StringLiteralToken}:                _GotoState33Action,
	{_State29, IdentifierToken}:                   _GotoState91Action,
	{_State29, TrueToken}:                         _GotoState36Action,
	{_State29, FalseToken}:                        _GotoState21Action,
	{_State29, StructToken}:                       _GotoState34Action,
	{_State29, FuncToken}:                         _GotoState23Action,
	{_State29, LabelDeclToken}:                    _GotoState26Action,
	{_State29, LparenToken}:                       _GotoState29Action,
	{_State29, LbracketToken}:                     _GotoState27Action,
	{_State29, NotToken}:                          _GotoState31Action,
	{_State29, SubToken}:                          _GotoState35Action,
	{_State29, MulToken}:                          _GotoState30Action,
	{_State29, BitNegToken}:                       _GotoState20Action,
	{_State29, BitAndToken}:                       _GotoState19Action,
	{_State29, LexErrorToken}:                     _GotoState28Action,
	{_State29, ExpressionType}:                    _GotoState95Action,
	{_State29, OptionalLabelDeclType}:             _GotoState50Action,
	{_State29, SequenceExprType}:                  _GotoState55Action,
	{_State29, BlockExprType}:                     _GotoState42Action,
	{_State29, CallExprType}:                      _GotoState43Action,
	{_State29, ArgumentsType}:                     _GotoState93Action,
	{_State29, ArgumentType}:                      _GotoState92Action,
	{_State29, ColonExpressionsType}:              _GotoState94Action,
	{_State29, OptionalExpressionType}:            _GotoState96Action,
	{_State29, AtomExprType}:                      _GotoState41Action,
	{_State29, LiteralType}:                       _GotoState48Action,
	{_State29, InitializeExprType}:                _GotoState47Action,
	{_State29, AccessExprType}:                    _GotoState37Action,
	{_State29, PostfixUnaryExprType}:              _GotoState52Action,
	{_State29, PrefixUnaryOpType}:                 _GotoState54Action,
	{_State29, PrefixUnaryExprType}:               _GotoState53Action,
	{_State29, MulExprType}:                       _GotoState49Action,
	{_State29, AddExprType}:                       _GotoState38Action,
	{_State29, CmpExprType}:                       _GotoState44Action,
	{_State29, AndExprType}:                       _GotoState39Action,
	{_State29, OrExprType}:                        _GotoState51Action,
	{_State29, InitializableTypeType}:             _GotoState46Action,
	{_State29, ExplicitStructDefType}:             _GotoState45Action,
	{_State29, AnonymousFuncExprType}:             _GotoState40Action,
	{_State34, LparenToken}:                       _GotoState97Action,
	{_State37, LbracketToken}:                     _GotoState100Action,
	{_State37, DotToken}:                          _GotoState99Action,
	{_State37, QuestionToken}:                     _GotoState101Action,
	{_State37, DollarLbracketToken}:               _GotoState98Action,
	{_State37, OptionalGenericBindingType}:        _GotoState102Action,
	{_State38, AddToken}:                          _GotoState103Action,
	{_State38, SubToken}:                          _GotoState106Action,
	{_State38, BitXorToken}:                       _GotoState105Action,
	{_State38, BitOrToken}:                        _GotoState104Action,
	{_State38, AddOpType}:                         _GotoState107Action,
	{_State39, AndToken}:                          _GotoState108Action,
	{_State44, EqualToken}:                        _GotoState109Action,
	{_State44, NotEqualToken}:                     _GotoState114Action,
	{_State44, LessToken}:                         _GotoState112Action,
	{_State44, LessOrEqualToken}:                  _GotoState113Action,
	{_State44, GreaterToken}:                      _GotoState110Action,
	{_State44, GreaterOrEqualToken}:               _GotoState111Action,
	{_State44, CmpOpType}:                         _GotoState115Action,
	{_State46, LparenToken}:                       _GotoState116Action,
	{_State49, MulToken}:                          _GotoState122Action,
	{_State49, DivToken}:                          _GotoState120Action,
	{_State49, ModToken}:                          _GotoState121Action,
	{_State49, BitAndToken}:                       _GotoState117Action,
	{_State49, BitLshiftToken}:                    _GotoState118Action,
	{_State49, BitRshiftToken}:                    _GotoState119Action,
	{_State49, MulOpType}:                         _GotoState123Action,
	{_State50, IfToken}:                           _GotoState126Action,
	{_State50, SwitchToken}:                       _GotoState128Action,
	{_State50, ForToken}:                          _GotoState125Action,
	{_State50, DoToken}:                           _GotoState124Action,
	{_State50, LbraceToken}:                       _GotoState127Action,
	{_State50, IfExprType}:                        _GotoState130Action,
	{_State50, SwitchExprType}:                    _GotoState132Action,
	{_State50, LoopExprType}:                      _GotoState131Action,
	{_State50, BlockBodyType}:                     _GotoState129Action,
	{_State51, OrToken}:                           _GotoState133Action,
	{_State54, IntegerLiteralToken}:               _GotoState25Action,
	{_State54, FloatLiteralToken}:                 _GotoState22Action,
	{_State54, RuneLiteralToken}:                  _GotoState32Action,
	{_State54, StringLiteralToken}:                _GotoState33Action,
	{_State54, IdentifierToken}:                   _GotoState24Action,
	{_State54, TrueToken}:                         _GotoState36Action,
	{_State54, FalseToken}:                        _GotoState21Action,
	{_State54, StructToken}:                       _GotoState34Action,
	{_State54, FuncToken}:                         _GotoState23Action,
	{_State54, LabelDeclToken}:                    _GotoState26Action,
	{_State54, LparenToken}:                       _GotoState29Action,
	{_State54, LbracketToken}:                     _GotoState27Action,
	{_State54, NotToken}:                          _GotoState31Action,
	{_State54, SubToken}:                          _GotoState35Action,
	{_State54, MulToken}:                          _GotoState30Action,
	{_State54, BitNegToken}:                       _GotoState20Action,
	{_State54, BitAndToken}:                       _GotoState19Action,
	{_State54, LexErrorToken}:                     _GotoState28Action,
	{_State54, OptionalLabelDeclType}:             _GotoState134Action,
	{_State54, BlockExprType}:                     _GotoState42Action,
	{_State54, CallExprType}:                      _GotoState43Action,
	{_State54, AtomExprType}:                      _GotoState41Action,
	{_State54, LiteralType}:                       _GotoState48Action,
	{_State54, InitializeExprType}:                _GotoState47Action,
	{_State54, AccessExprType}:                    _GotoState37Action,
	{_State54, PostfixUnaryExprType}:              _GotoState52Action,
	{_State54, PrefixUnaryOpType}:                 _GotoState54Action,
	{_State54, PrefixUnaryExprType}:               _GotoState135Action,
	{_State54, InitializableTypeType}:             _GotoState46Action,
	{_State54, ExplicitStructDefType}:             _GotoState45Action,
	{_State54, AnonymousFuncExprType}:             _GotoState40Action,
	{_State58, LessToken}:                         _GotoState136Action,
	{_State60, NewlinesToken}:                     _GotoState137Action,
	{_State60, OptionalNewlinesType}:              _GotoState138Action,
	{_State65, LparenToken}:                       _GotoState139Action,
	{_State66, DollarLbracketToken}:               _GotoState140Action,
	{_State66, EqualToken}:                        _GotoState141Action,
	{_State66, OptionalGenericParametersType}:     _GotoState142Action,
	{_State67, IdentifierToken}:                   _GotoState143Action,
	{_State67, ParameterDefType}:                  _GotoState144Action,
	{_State68, IdentifierToken}:                   _GotoState145Action,
	{_State69, IdentifierToken}:                   _GotoState143Action,
	{_State69, ParameterDefType}:                  _GotoState147Action,
	{_State69, ParameterDefsType}:                 _GotoState148Action,
	{_State69, OptionalParameterDefsType}:         _GotoState146Action,
	{_State70, IdentifierToken}:                   _GotoState76Action,
	{_State70, StructToken}:                       _GotoState34Action,
	{_State70, EnumToken}:                         _GotoState73Action,
	{_State70, TraitToken}:                        _GotoState81Action,
	{_State70, FuncToken}:                         _GotoState75Action,
	{_State70, LparenToken}:                       _GotoState78Action,
	{_State70, LbracketToken}:                     _GotoState27Action,
	{_State70, DotToken}:                          _GotoState72Action,
	{_State70, QuestionToken}:                     _GotoState79Action,
	{_State70, ExclaimToken}:                      _GotoState74Action,
	{_State70, TildeTildeToken}:                   _GotoState80Action,
	{_State70, BitNegToken}:                       _GotoState71Action,
	{_State70, BitAndToken}:                       _GotoState70Action,
	{_State70, LexErrorToken}:                     _GotoState77Action,
	{_State70, InitializableTypeType}:             _GotoState87Action,
	{_State70, AtomTypeType}:                      _GotoState82Action,
	{_State70, ReturnableTypeType}:                _GotoState149Action,
	{_State70, ImplicitStructDefType}:             _GotoState86Action,
	{_State70, ExplicitStructDefType}:             _GotoState45Action,
	{_State70, ImplicitEnumDefType}:               _GotoState85Action,
	{_State70, ExplicitEnumDefType}:               _GotoState83Action,
	{_State70, TraitDefType}:                      _GotoState89Action,
	{_State70, FuncTypeType}:                      _GotoState84Action,
	{_State71, IdentifierToken}:                   _GotoState76Action,
	{_State71, StructToken}:                       _GotoState34Action,
	{_State71, EnumToken}:                         _GotoState73Action,
	{_State71, TraitToken}:                        _GotoState81Action,
	{_State71, FuncToken}:                         _GotoState75Action,
	{_State71, LparenToken}:                       _GotoState78Action,
	{_State71, LbracketToken}:                     _GotoState27Action,
	{_State71, DotToken}:                          _GotoState72Action,
	{_State71, QuestionToken}:                     _GotoState79Action,
	{_State71, ExclaimToken}:                      _GotoState74Action,
	{_State71, TildeTildeToken}:                   _GotoState80Action,
	{_State71, BitNegToken}:                       _GotoState71Action,
	{_State71, BitAndToken}:                       _GotoState70Action,
	{_State71, LexErrorToken}:                     _GotoState77Action,
	{_State71, InitializableTypeType}:             _GotoState87Action,
	{_State71, AtomTypeType}:                      _GotoState82Action,
	{_State71, ReturnableTypeType}:                _GotoState150Action,
	{_State71, ImplicitStructDefType}:             _GotoState86Action,
	{_State71, ExplicitStructDefType}:             _GotoState45Action,
	{_State71, ImplicitEnumDefType}:               _GotoState85Action,
	{_State71, ExplicitEnumDefType}:               _GotoState83Action,
	{_State71, TraitDefType}:                      _GotoState89Action,
	{_State71, FuncTypeType}:                      _GotoState84Action,
	{_State72, DollarLbracketToken}:               _GotoState98Action,
	{_State72, OptionalGenericBindingType}:        _GotoState151Action,
	{_State73, LparenToken}:                       _GotoState152Action,
	{_State74, IdentifierToken}:                   _GotoState76Action,
	{_State74, StructToken}:                       _GotoState34Action,
	{_State74, EnumToken}:                         _GotoState73Action,
	{_State74, TraitToken}:                        _GotoState81Action,
	{_State74, FuncToken}:                         _GotoState75Action,
	{_State74, LparenToken}:                       _GotoState78Action,
	{_State74, LbracketToken}:                     _GotoState27Action,
	{_State74, DotToken}:                          _GotoState72Action,
	{_State74, QuestionToken}:                     _GotoState79Action,
	{_State74, ExclaimToken}:                      _GotoState74Action,
	{_State74, TildeTildeToken}:                   _GotoState80Action,
	{_State74, BitNegToken}:                       _GotoState71Action,
	{_State74, BitAndToken}:                       _GotoState70Action,
	{_State74, LexErrorToken}:                     _GotoState77Action,
	{_State74, InitializableTypeType}:             _GotoState87Action,
	{_State74, AtomTypeType}:                      _GotoState82Action,
	{_State74, ReturnableTypeType}:                _GotoState153Action,
	{_State74, ImplicitStructDefType}:             _GotoState86Action,
	{_State74, ExplicitStructDefType}:             _GotoState45Action,
	{_State74, ImplicitEnumDefType}:               _GotoState85Action,
	{_State74, ExplicitEnumDefType}:               _GotoState83Action,
	{_State74, TraitDefType}:                      _GotoState89Action,
	{_State74, FuncTypeType}:                      _GotoState84Action,
	{_State75, LparenToken}:                       _GotoState154Action,
	{_State76, DotToken}:                          _GotoState155Action,
	{_State76, DollarLbracketToken}:               _GotoState98Action,
	{_State76, OptionalGenericBindingType}:        _GotoState156Action,
	{_State78, IdentifierToken}:                   _GotoState157Action,
	{_State78, UnsafeToken}:                       _GotoState58Action,
	{_State78, StructToken}:                       _GotoState34Action,
	{_State78, EnumToken}:                         _GotoState73Action,
	{_State78, TraitToken}:                        _GotoState81Action,
	{_State78, FuncToken}:                         _GotoState75Action,
	{_State78, LparenToken}:                       _GotoState78Action,
	{_State78, LbracketToken}:                     _GotoState27Action,
	{_State78, DotToken}:                          _GotoState72Action,
	{_State78, QuestionToken}:                     _GotoState79Action,
	{_State78, ExclaimToken}:                      _GotoState74Action,
	{_State78, TildeTildeToken}:                   _GotoState80Action,
	{_State78, BitNegToken}:                       _GotoState71Action,
	{_State78, BitAndToken}:                       _GotoState70Action,
	{_State78, LexErrorToken}:                     _GotoState77Action,
	{_State78, UnsafeStatementType}:               _GotoState163Action,
	{_State78, InitializableTypeType}:             _GotoState87Action,
	{_State78, AtomTypeType}:                      _GotoState82Action,
	{_State78, ReturnableTypeType}:                _GotoState88Action,
	{_State78, ValueTypeType}:                     _GotoState164Action,
	{_State78, FieldDefType}:                      _GotoState159Action,
	{_State78, ImplicitFieldDefsType}:             _GotoState161Action,
	{_State78, OptionalImplicitFieldDefsType}:     _GotoState162Action,
	{_State78, ImplicitStructDefType}:             _GotoState86Action,
	{_State78, ExplicitStructDefType}:             _GotoState45Action,
	{_State78, EnumValueDefType}:                  _GotoState158Action,
	{_State78, ImplicitEnumValueDefsType}:         _GotoState160Action,
	{_State78, ImplicitEnumDefType}:               _GotoState85Action,
	{_State78, ExplicitEnumDefType}:               _GotoState83Action,
	{_State78, TraitDefType}:                      _GotoState89Action,
	{_State78, FuncTypeType}:                      _GotoState84Action,
	{_State79, IdentifierToken}:                   _GotoState76Action,
	{_State79, StructToken}:                       _GotoState34Action,
	{_State79, EnumToken}:                         _GotoState73Action,
	{_State79, TraitToken}:                        _GotoState81Action,
	{_State79, FuncToken}:                         _GotoState75Action,
	{_State79, LparenToken}:                       _GotoState78Action,
	{_State79, LbracketToken}:                     _GotoState27Action,
	{_State79, DotToken}:                          _GotoState72Action,
	{_State79, QuestionToken}:                     _GotoState79Action,
	{_State79, ExclaimToken}:                      _GotoState74Action,
	{_State79, TildeTildeToken}:                   _GotoState80Action,
	{_State79, BitNegToken}:                       _GotoState71Action,
	{_State79, BitAndToken}:                       _GotoState70Action,
	{_State79, LexErrorToken}:                     _GotoState77Action,
	{_State79, InitializableTypeType}:             _GotoState87Action,
	{_State79, AtomTypeType}:                      _GotoState82Action,
	{_State79, ReturnableTypeType}:                _GotoState165Action,
	{_State79, ImplicitStructDefType}:             _GotoState86Action,
	{_State79, ExplicitStructDefType}:             _GotoState45Action,
	{_State79, ImplicitEnumDefType}:               _GotoState85Action,
	{_State79, ExplicitEnumDefType}:               _GotoState83Action,
	{_State79, TraitDefType}:                      _GotoState89Action,
	{_State79, FuncTypeType}:                      _GotoState84Action,
	{_State80, IdentifierToken}:                   _GotoState76Action,
	{_State80, StructToken}:                       _GotoState34Action,
	{_State80, EnumToken}:                         _GotoState73Action,
	{_State80, TraitToken}:                        _GotoState81Action,
	{_State80, FuncToken}:                         _GotoState75Action,
	{_State80, LparenToken}:                       _GotoState78Action,
	{_State80, LbracketToken}:                     _GotoState27Action,
	{_State80, DotToken}:                          _GotoState72Action,
	{_State80, QuestionToken}:                     _GotoState79Action,
	{_State80, ExclaimToken}:                      _GotoState74Action,
	{_State80, TildeTildeToken}:                   _GotoState80Action,
	{_State80, BitNegToken}:                       _GotoState71Action,
	{_State80, BitAndToken}:                       _GotoState70Action,
	{_State80, LexErrorToken}:                     _GotoState77Action,
	{_State80, InitializableTypeType}:             _GotoState87Action,
	{_State80, AtomTypeType}:                      _GotoState82Action,
	{_State80, ReturnableTypeType}:                _GotoState166Action,
	{_State80, ImplicitStructDefType}:             _GotoState86Action,
	{_State80, ExplicitStructDefType}:             _GotoState45Action,
	{_State80, ImplicitEnumDefType}:               _GotoState85Action,
	{_State80, ExplicitEnumDefType}:               _GotoState83Action,
	{_State80, TraitDefType}:                      _GotoState89Action,
	{_State80, FuncTypeType}:                      _GotoState84Action,
	{_State81, LparenToken}:                       _GotoState167Action,
	{_State90, RbracketToken}:                     _GotoState172Action,
	{_State90, CommaToken}:                        _GotoState170Action,
	{_State90, ColonToken}:                        _GotoState169Action,
	{_State90, AddToken}:                          _GotoState168Action,
	{_State90, SubToken}:                          _GotoState173Action,
	{_State90, MulToken}:                          _GotoState171Action,
	{_State91, AssignToken}:                       _GotoState174Action,
	{_State93, RparenToken}:                       _GotoState176Action,
	{_State93, CommaToken}:                        _GotoState175Action,
	{_State94, ColonToken}:                        _GotoState177Action,
	{_State96, ColonToken}:                        _GotoState178Action,
	{_State97, IdentifierToken}:                   _GotoState157Action,
	{_State97, UnsafeToken}:                       _GotoState58Action,
	{_State97, StructToken}:                       _GotoState34Action,
	{_State97, EnumToken}:                         _GotoState73Action,
	{_State97, TraitToken}:                        _GotoState81Action,
	{_State97, FuncToken}:                         _GotoState75Action,
	{_State97, LparenToken}:                       _GotoState78Action,
	{_State97, LbracketToken}:                     _GotoState27Action,
	{_State97, DotToken}:                          _GotoState72Action,
	{_State97, QuestionToken}:                     _GotoState79Action,
	{_State97, ExclaimToken}:                      _GotoState74Action,
	{_State97, TildeTildeToken}:                   _GotoState80Action,
	{_State97, BitNegToken}:                       _GotoState71Action,
	{_State97, BitAndToken}:                       _GotoState70Action,
	{_State97, LexErrorToken}:                     _GotoState77Action,
	{_State97, UnsafeStatementType}:               _GotoState163Action,
	{_State97, InitializableTypeType}:             _GotoState87Action,
	{_State97, AtomTypeType}:                      _GotoState82Action,
	{_State97, ReturnableTypeType}:                _GotoState88Action,
	{_State97, ValueTypeType}:                     _GotoState164Action,
	{_State97, FieldDefType}:                      _GotoState180Action,
	{_State97, ImplicitStructDefType}:             _GotoState86Action,
	{_State97, ExplicitFieldDefsType}:             _GotoState179Action,
	{_State97, OptionalExplicitFieldDefsType}:     _GotoState181Action,
	{_State97, ExplicitStructDefType}:             _GotoState45Action,
	{_State97, ImplicitEnumDefType}:               _GotoState85Action,
	{_State97, ExplicitEnumDefType}:               _GotoState83Action,
	{_State97, TraitDefType}:                      _GotoState89Action,
	{_State97, FuncTypeType}:                      _GotoState84Action,
	{_State98, IdentifierToken}:                   _GotoState76Action,
	{_State98, StructToken}:                       _GotoState34Action,
	{_State98, EnumToken}:                         _GotoState73Action,
	{_State98, TraitToken}:                        _GotoState81Action,
	{_State98, FuncToken}:                         _GotoState75Action,
	{_State98, LparenToken}:                       _GotoState78Action,
	{_State98, LbracketToken}:                     _GotoState27Action,
	{_State98, DotToken}:                          _GotoState72Action,
	{_State98, QuestionToken}:                     _GotoState79Action,
	{_State98, ExclaimToken}:                      _GotoState74Action,
	{_State98, TildeTildeToken}:                   _GotoState80Action,
	{_State98, BitNegToken}:                       _GotoState71Action,
	{_State98, BitAndToken}:                       _GotoState70Action,
	{_State98, LexErrorToken}:                     _GotoState77Action,
	{_State98, OptionalGenericArgumentsType}:      _GotoState183Action,
	{_State98, GenericArgumentsType}:              _GotoState182Action,
	{_State98, InitializableTypeType}:             _GotoState87Action,
	{_State98, AtomTypeType}:                      _GotoState82Action,
	{_State98, ReturnableTypeType}:                _GotoState88Action,
	{_State98, ValueTypeType}:                     _GotoState184Action,
	{_State98, ImplicitStructDefType}:             _GotoState86Action,
	{_State98, ExplicitStructDefType}:             _GotoState45Action,
	{_State98, ImplicitEnumDefType}:               _GotoState85Action,
	{_State98, ExplicitEnumDefType}:               _GotoState83Action,
	{_State98, TraitDefType}:                      _GotoState89Action,
	{_State98, FuncTypeType}:                      _GotoState84Action,
	{_State99, IdentifierToken}:                   _GotoState185Action,
	{_State100, IntegerLiteralToken}:              _GotoState25Action,
	{_State100, FloatLiteralToken}:                _GotoState22Action,
	{_State100, RuneLiteralToken}:                 _GotoState32Action,
	{_State100, StringLiteralToken}:               _GotoState33Action,
	{_State100, IdentifierToken}:                  _GotoState91Action,
	{_State100, TrueToken}:                        _GotoState36Action,
	{_State100, FalseToken}:                       _GotoState21Action,
	{_State100, StructToken}:                      _GotoState34Action,
	{_State100, FuncToken}:                        _GotoState23Action,
	{_State100, LabelDeclToken}:                   _GotoState26Action,
	{_State100, LparenToken}:                      _GotoState29Action,
	{_State100, LbracketToken}:                    _GotoState27Action,
	{_State100, NotToken}:                         _GotoState31Action,
	{_State100, SubToken}:                         _GotoState35Action,
	{_State100, MulToken}:                         _GotoState30Action,
	{_State100, BitNegToken}:                      _GotoState20Action,
	{_State100, BitAndToken}:                      _GotoState19Action,
	{_State100, LexErrorToken}:                    _GotoState28Action,
	{_State100, ExpressionType}:                   _GotoState95Action,
	{_State100, OptionalLabelDeclType}:            _GotoState50Action,
	{_State100, SequenceExprType}:                 _GotoState55Action,
	{_State100, BlockExprType}:                    _GotoState42Action,
	{_State100, CallExprType}:                     _GotoState43Action,
	{_State100, ArgumentType}:                     _GotoState186Action,
	{_State100, ColonExpressionsType}:             _GotoState94Action,
	{_State100, OptionalExpressionType}:           _GotoState96Action,
	{_State100, AtomExprType}:                     _GotoState41Action,
	{_State100, LiteralType}:                      _GotoState48Action,
	{_State100, InitializeExprType}:               _GotoState47Action,
	{_State100, AccessExprType}:                   _GotoState37Action,
	{_State100, PostfixUnaryExprType}:             _GotoState52Action,
	{_State100, PrefixUnaryOpType}:                _GotoState54Action,
	{_State100, PrefixUnaryExprType}:              _GotoState53Action,
	{_State100, MulExprType}:                      _GotoState49Action,
	{_State100, AddExprType}:                      _GotoState38Action,
	{_State100, CmpExprType}:                      _GotoState44Action,
	{_State100, AndExprType}:                      _GotoState39Action,
	{_State100, OrExprType}:                       _GotoState51Action,
	{_State100, InitializableTypeType}:            _GotoState46Action,
	{_State100, ExplicitStructDefType}:            _GotoState45Action,
	{_State100, AnonymousFuncExprType}:            _GotoState40Action,
	{_State102, LparenToken}:                      _GotoState187Action,
	{_State107, IntegerLiteralToken}:              _GotoState25Action,
	{_State107, FloatLiteralToken}:                _GotoState22Action,
	{_State107, RuneLiteralToken}:                 _GotoState32Action,
	{_State107, StringLiteralToken}:               _GotoState33Action,
	{_State107, IdentifierToken}:                  _GotoState24Action,
	{_State107, TrueToken}:                        _GotoState36Action,
	{_State107, FalseToken}:                       _GotoState21Action,
	{_State107, StructToken}:                      _GotoState34Action,
	{_State107, FuncToken}:                        _GotoState23Action,
	{_State107, LabelDeclToken}:                   _GotoState26Action,
	{_State107, LparenToken}:                      _GotoState29Action,
	{_State107, LbracketToken}:                    _GotoState27Action,
	{_State107, NotToken}:                         _GotoState31Action,
	{_State107, SubToken}:                         _GotoState35Action,
	{_State107, MulToken}:                         _GotoState30Action,
	{_State107, BitNegToken}:                      _GotoState20Action,
	{_State107, BitAndToken}:                      _GotoState19Action,
	{_State107, LexErrorToken}:                    _GotoState28Action,
	{_State107, OptionalLabelDeclType}:            _GotoState134Action,
	{_State107, BlockExprType}:                    _GotoState42Action,
	{_State107, CallExprType}:                     _GotoState43Action,
	{_State107, AtomExprType}:                     _GotoState41Action,
	{_State107, LiteralType}:                      _GotoState48Action,
	{_State107, InitializeExprType}:               _GotoState47Action,
	{_State107, AccessExprType}:                   _GotoState37Action,
	{_State107, PostfixUnaryExprType}:             _GotoState52Action,
	{_State107, PrefixUnaryOpType}:                _GotoState54Action,
	{_State107, PrefixUnaryExprType}:              _GotoState53Action,
	{_State107, MulExprType}:                      _GotoState188Action,
	{_State107, InitializableTypeType}:            _GotoState46Action,
	{_State107, ExplicitStructDefType}:            _GotoState45Action,
	{_State107, AnonymousFuncExprType}:            _GotoState40Action,
	{_State108, IntegerLiteralToken}:              _GotoState25Action,
	{_State108, FloatLiteralToken}:                _GotoState22Action,
	{_State108, RuneLiteralToken}:                 _GotoState32Action,
	{_State108, StringLiteralToken}:               _GotoState33Action,
	{_State108, IdentifierToken}:                  _GotoState24Action,
	{_State108, TrueToken}:                        _GotoState36Action,
	{_State108, FalseToken}:                       _GotoState21Action,
	{_State108, StructToken}:                      _GotoState34Action,
	{_State108, FuncToken}:                        _GotoState23Action,
	{_State108, LabelDeclToken}:                   _GotoState26Action,
	{_State108, LparenToken}:                      _GotoState29Action,
	{_State108, LbracketToken}:                    _GotoState27Action,
	{_State108, NotToken}:                         _GotoState31Action,
	{_State108, SubToken}:                         _GotoState35Action,
	{_State108, MulToken}:                         _GotoState30Action,
	{_State108, BitNegToken}:                      _GotoState20Action,
	{_State108, BitAndToken}:                      _GotoState19Action,
	{_State108, LexErrorToken}:                    _GotoState28Action,
	{_State108, OptionalLabelDeclType}:            _GotoState134Action,
	{_State108, BlockExprType}:                    _GotoState42Action,
	{_State108, CallExprType}:                     _GotoState43Action,
	{_State108, AtomExprType}:                     _GotoState41Action,
	{_State108, LiteralType}:                      _GotoState48Action,
	{_State108, InitializeExprType}:               _GotoState47Action,
	{_State108, AccessExprType}:                   _GotoState37Action,
	{_State108, PostfixUnaryExprType}:             _GotoState52Action,
	{_State108, PrefixUnaryOpType}:                _GotoState54Action,
	{_State108, PrefixUnaryExprType}:              _GotoState53Action,
	{_State108, MulExprType}:                      _GotoState49Action,
	{_State108, AddExprType}:                      _GotoState38Action,
	{_State108, CmpExprType}:                      _GotoState189Action,
	{_State108, InitializableTypeType}:            _GotoState46Action,
	{_State108, ExplicitStructDefType}:            _GotoState45Action,
	{_State108, AnonymousFuncExprType}:            _GotoState40Action,
	{_State115, IntegerLiteralToken}:              _GotoState25Action,
	{_State115, FloatLiteralToken}:                _GotoState22Action,
	{_State115, RuneLiteralToken}:                 _GotoState32Action,
	{_State115, StringLiteralToken}:               _GotoState33Action,
	{_State115, IdentifierToken}:                  _GotoState24Action,
	{_State115, TrueToken}:                        _GotoState36Action,
	{_State115, FalseToken}:                       _GotoState21Action,
	{_State115, StructToken}:                      _GotoState34Action,
	{_State115, FuncToken}:                        _GotoState23Action,
	{_State115, LabelDeclToken}:                   _GotoState26Action,
	{_State115, LparenToken}:                      _GotoState29Action,
	{_State115, LbracketToken}:                    _GotoState27Action,
	{_State115, NotToken}:                         _GotoState31Action,
	{_State115, SubToken}:                         _GotoState35Action,
	{_State115, MulToken}:                         _GotoState30Action,
	{_State115, BitNegToken}:                      _GotoState20Action,
	{_State115, BitAndToken}:                      _GotoState19Action,
	{_State115, LexErrorToken}:                    _GotoState28Action,
	{_State115, OptionalLabelDeclType}:            _GotoState134Action,
	{_State115, BlockExprType}:                    _GotoState42Action,
	{_State115, CallExprType}:                     _GotoState43Action,
	{_State115, AtomExprType}:                     _GotoState41Action,
	{_State115, LiteralType}:                      _GotoState48Action,
	{_State115, InitializeExprType}:               _GotoState47Action,
	{_State115, AccessExprType}:                   _GotoState37Action,
	{_State115, PostfixUnaryExprType}:             _GotoState52Action,
	{_State115, PrefixUnaryOpType}:                _GotoState54Action,
	{_State115, PrefixUnaryExprType}:              _GotoState53Action,
	{_State115, MulExprType}:                      _GotoState49Action,
	{_State115, AddExprType}:                      _GotoState190Action,
	{_State115, InitializableTypeType}:            _GotoState46Action,
	{_State115, ExplicitStructDefType}:            _GotoState45Action,
	{_State115, AnonymousFuncExprType}:            _GotoState40Action,
	{_State116, IntegerLiteralToken}:              _GotoState25Action,
	{_State116, FloatLiteralToken}:                _GotoState22Action,
	{_State116, RuneLiteralToken}:                 _GotoState32Action,
	{_State116, StringLiteralToken}:               _GotoState33Action,
	{_State116, IdentifierToken}:                  _GotoState91Action,
	{_State116, TrueToken}:                        _GotoState36Action,
	{_State116, FalseToken}:                       _GotoState21Action,
	{_State116, StructToken}:                      _GotoState34Action,
	{_State116, FuncToken}:                        _GotoState23Action,
	{_State116, LabelDeclToken}:                   _GotoState26Action,
	{_State116, LparenToken}:                      _GotoState29Action,
	{_State116, LbracketToken}:                    _GotoState27Action,
	{_State116, NotToken}:                         _GotoState31Action,
	{_State116, SubToken}:                         _GotoState35Action,
	{_State116, MulToken}:                         _GotoState30Action,
	{_State116, BitNegToken}:                      _GotoState20Action,
	{_State116, BitAndToken}:                      _GotoState19Action,
	{_State116, LexErrorToken}:                    _GotoState28Action,
	{_State116, ExpressionType}:                   _GotoState95Action,
	{_State116, OptionalLabelDeclType}:            _GotoState50Action,
	{_State116, SequenceExprType}:                 _GotoState55Action,
	{_State116, BlockExprType}:                    _GotoState42Action,
	{_State116, CallExprType}:                     _GotoState43Action,
	{_State116, ArgumentsType}:                    _GotoState191Action,
	{_State116, ArgumentType}:                     _GotoState92Action,
	{_State116, ColonExpressionsType}:             _GotoState94Action,
	{_State116, OptionalExpressionType}:           _GotoState96Action,
	{_State116, AtomExprType}:                     _GotoState41Action,
	{_State116, LiteralType}:                      _GotoState48Action,
	{_State116, InitializeExprType}:               _GotoState47Action,
	{_State116, AccessExprType}:                   _GotoState37Action,
	{_State116, PostfixUnaryExprType}:             _GotoState52Action,
	{_State116, PrefixUnaryOpType}:                _GotoState54Action,
	{_State116, PrefixUnaryExprType}:              _GotoState53Action,
	{_State116, MulExprType}:                      _GotoState49Action,
	{_State116, AddExprType}:                      _GotoState38Action,
	{_State116, CmpExprType}:                      _GotoState44Action,
	{_State116, AndExprType}:                      _GotoState39Action,
	{_State116, OrExprType}:                       _GotoState51Action,
	{_State116, InitializableTypeType}:            _GotoState46Action,
	{_State116, ExplicitStructDefType}:            _GotoState45Action,
	{_State116, AnonymousFuncExprType}:            _GotoState40Action,
	{_State123, IntegerLiteralToken}:              _GotoState25Action,
	{_State123, FloatLiteralToken}:                _GotoState22Action,
	{_State123, RuneLiteralToken}:                 _GotoState32Action,
	{_State123, StringLiteralToken}:               _GotoState33Action,
	{_State123, IdentifierToken}:                  _GotoState24Action,
	{_State123, TrueToken}:                        _GotoState36Action,
	{_State123, FalseToken}:                       _GotoState21Action,
	{_State123, StructToken}:                      _GotoState34Action,
	{_State123, FuncToken}:                        _GotoState23Action,
	{_State123, LabelDeclToken}:                   _GotoState26Action,
	{_State123, LparenToken}:                      _GotoState29Action,
	{_State123, LbracketToken}:                    _GotoState27Action,
	{_State123, NotToken}:                         _GotoState31Action,
	{_State123, SubToken}:                         _GotoState35Action,
	{_State123, MulToken}:                         _GotoState30Action,
	{_State123, BitNegToken}:                      _GotoState20Action,
	{_State123, BitAndToken}:                      _GotoState19Action,
	{_State123, LexErrorToken}:                    _GotoState28Action,
	{_State123, OptionalLabelDeclType}:            _GotoState134Action,
	{_State123, BlockExprType}:                    _GotoState42Action,
	{_State123, CallExprType}:                     _GotoState43Action,
	{_State123, AtomExprType}:                     _GotoState41Action,
	{_State123, LiteralType}:                      _GotoState48Action,
	{_State123, InitializeExprType}:               _GotoState47Action,
	{_State123, AccessExprType}:                   _GotoState37Action,
	{_State123, PostfixUnaryExprType}:             _GotoState52Action,
	{_State123, PrefixUnaryOpType}:                _GotoState54Action,
	{_State123, PrefixUnaryExprType}:              _GotoState192Action,
	{_State123, InitializableTypeType}:            _GotoState46Action,
	{_State123, ExplicitStructDefType}:            _GotoState45Action,
	{_State123, AnonymousFuncExprType}:            _GotoState40Action,
	{_State124, LbraceToken}:                      _GotoState127Action,
	{_State124, BlockBodyType}:                    _GotoState193Action,
	{_State125, IntegerLiteralToken}:              _GotoState25Action,
	{_State125, FloatLiteralToken}:                _GotoState22Action,
	{_State125, RuneLiteralToken}:                 _GotoState32Action,
	{_State125, StringLiteralToken}:               _GotoState33Action,
	{_State125, IdentifierToken}:                  _GotoState24Action,
	{_State125, TrueToken}:                        _GotoState36Action,
	{_State125, FalseToken}:                       _GotoState21Action,
	{_State125, InToken}:                          _GotoState194Action,
	{_State125, StructToken}:                      _GotoState34Action,
	{_State125, FuncToken}:                        _GotoState23Action,
	{_State125, LabelDeclToken}:                   _GotoState26Action,
	{_State125, LparenToken}:                      _GotoState29Action,
	{_State125, LbracketToken}:                    _GotoState27Action,
	{_State125, SemicolonToken}:                   _GotoState195Action,
	{_State125, NotToken}:                         _GotoState31Action,
	{_State125, SubToken}:                         _GotoState35Action,
	{_State125, MulToken}:                         _GotoState30Action,
	{_State125, BitNegToken}:                      _GotoState20Action,
	{_State125, BitAndToken}:                      _GotoState19Action,
	{_State125, LexErrorToken}:                    _GotoState28Action,
	{_State125, OptionalLabelDeclType}:            _GotoState134Action,
	{_State125, SequenceExprType}:                 _GotoState196Action,
	{_State125, BlockExprType}:                    _GotoState42Action,
	{_State125, CallExprType}:                     _GotoState43Action,
	{_State125, AtomExprType}:                     _GotoState41Action,
	{_State125, LiteralType}:                      _GotoState48Action,
	{_State125, InitializeExprType}:               _GotoState47Action,
	{_State125, AccessExprType}:                   _GotoState37Action,
	{_State125, PostfixUnaryExprType}:             _GotoState52Action,
	{_State125, PrefixUnaryOpType}:                _GotoState54Action,
	{_State125, PrefixUnaryExprType}:              _GotoState53Action,
	{_State125, MulExprType}:                      _GotoState49Action,
	{_State125, AddExprType}:                      _GotoState38Action,
	{_State125, CmpExprType}:                      _GotoState44Action,
	{_State125, AndExprType}:                      _GotoState39Action,
	{_State125, OrExprType}:                       _GotoState51Action,
	{_State125, InitializableTypeType}:            _GotoState46Action,
	{_State125, ExplicitStructDefType}:            _GotoState45Action,
	{_State125, AnonymousFuncExprType}:            _GotoState40Action,
	{_State126, IntegerLiteralToken}:              _GotoState25Action,
	{_State126, FloatLiteralToken}:                _GotoState22Action,
	{_State126, RuneLiteralToken}:                 _GotoState32Action,
	{_State126, StringLiteralToken}:               _GotoState33Action,
	{_State126, IdentifierToken}:                  _GotoState24Action,
	{_State126, TrueToken}:                        _GotoState36Action,
	{_State126, FalseToken}:                       _GotoState21Action,
	{_State126, StructToken}:                      _GotoState34Action,
	{_State126, FuncToken}:                        _GotoState23Action,
	{_State126, LabelDeclToken}:                   _GotoState26Action,
	{_State126, LparenToken}:                      _GotoState29Action,
	{_State126, LbracketToken}:                    _GotoState27Action,
	{_State126, NotToken}:                         _GotoState31Action,
	{_State126, SubToken}:                         _GotoState35Action,
	{_State126, MulToken}:                         _GotoState30Action,
	{_State126, BitNegToken}:                      _GotoState20Action,
	{_State126, BitAndToken}:                      _GotoState19Action,
	{_State126, LexErrorToken}:                    _GotoState28Action,
	{_State126, OptionalLabelDeclType}:            _GotoState134Action,
	{_State126, SequenceExprType}:                 _GotoState197Action,
	{_State126, BlockExprType}:                    _GotoState42Action,
	{_State126, CallExprType}:                     _GotoState43Action,
	{_State126, AtomExprType}:                     _GotoState41Action,
	{_State126, LiteralType}:                      _GotoState48Action,
	{_State126, InitializeExprType}:               _GotoState47Action,
	{_State126, AccessExprType}:                   _GotoState37Action,
	{_State126, PostfixUnaryExprType}:             _GotoState52Action,
	{_State126, PrefixUnaryOpType}:                _GotoState54Action,
	{_State126, PrefixUnaryExprType}:              _GotoState53Action,
	{_State126, MulExprType}:                      _GotoState49Action,
	{_State126, AddExprType}:                      _GotoState38Action,
	{_State126, CmpExprType}:                      _GotoState44Action,
	{_State126, AndExprType}:                      _GotoState39Action,
	{_State126, OrExprType}:                       _GotoState51Action,
	{_State126, InitializableTypeType}:            _GotoState46Action,
	{_State126, ExplicitStructDefType}:            _GotoState45Action,
	{_State126, AnonymousFuncExprType}:            _GotoState40Action,
	{_State127, StatementsType}:                   _GotoState198Action,
	{_State128, IntegerLiteralToken}:              _GotoState25Action,
	{_State128, FloatLiteralToken}:                _GotoState22Action,
	{_State128, RuneLiteralToken}:                 _GotoState32Action,
	{_State128, StringLiteralToken}:               _GotoState33Action,
	{_State128, IdentifierToken}:                  _GotoState24Action,
	{_State128, TrueToken}:                        _GotoState36Action,
	{_State128, FalseToken}:                       _GotoState21Action,
	{_State128, StructToken}:                      _GotoState34Action,
	{_State128, FuncToken}:                        _GotoState23Action,
	{_State128, LabelDeclToken}:                   _GotoState26Action,
	{_State128, LparenToken}:                      _GotoState29Action,
	{_State128, LbracketToken}:                    _GotoState27Action,
	{_State128, NotToken}:                         _GotoState31Action,
	{_State128, SubToken}:                         _GotoState35Action,
	{_State128, MulToken}:                         _GotoState30Action,
	{_State128, BitNegToken}:                      _GotoState20Action,
	{_State128, BitAndToken}:                      _GotoState19Action,
	{_State128, LexErrorToken}:                    _GotoState28Action,
	{_State128, OptionalLabelDeclType}:            _GotoState134Action,
	{_State128, SequenceExprType}:                 _GotoState199Action,
	{_State128, BlockExprType}:                    _GotoState42Action,
	{_State128, CallExprType}:                     _GotoState43Action,
	{_State128, AtomExprType}:                     _GotoState41Action,
	{_State128, LiteralType}:                      _GotoState48Action,
	{_State128, InitializeExprType}:               _GotoState47Action,
	{_State128, AccessExprType}:                   _GotoState37Action,
	{_State128, PostfixUnaryExprType}:             _GotoState52Action,
	{_State128, PrefixUnaryOpType}:                _GotoState54Action,
	{_State128, PrefixUnaryExprType}:              _GotoState53Action,
	{_State128, MulExprType}:                      _GotoState49Action,
	{_State128, AddExprType}:                      _GotoState38Action,
	{_State128, CmpExprType}:                      _GotoState44Action,
	{_State128, AndExprType}:                      _GotoState39Action,
	{_State128, OrExprType}:                       _GotoState51Action,
	{_State128, InitializableTypeType}:            _GotoState46Action,
	{_State128, ExplicitStructDefType}:            _GotoState45Action,
	{_State128, AnonymousFuncExprType}:            _GotoState40Action,
	{_State133, IntegerLiteralToken}:              _GotoState25Action,
	{_State133, FloatLiteralToken}:                _GotoState22Action,
	{_State133, RuneLiteralToken}:                 _GotoState32Action,
	{_State133, StringLiteralToken}:               _GotoState33Action,
	{_State133, IdentifierToken}:                  _GotoState24Action,
	{_State133, TrueToken}:                        _GotoState36Action,
	{_State133, FalseToken}:                       _GotoState21Action,
	{_State133, StructToken}:                      _GotoState34Action,
	{_State133, FuncToken}:                        _GotoState23Action,
	{_State133, LabelDeclToken}:                   _GotoState26Action,
	{_State133, LparenToken}:                      _GotoState29Action,
	{_State133, LbracketToken}:                    _GotoState27Action,
	{_State133, NotToken}:                         _GotoState31Action,
	{_State133, SubToken}:                         _GotoState35Action,
	{_State133, MulToken}:                         _GotoState30Action,
	{_State133, BitNegToken}:                      _GotoState20Action,
	{_State133, BitAndToken}:                      _GotoState19Action,
	{_State133, LexErrorToken}:                    _GotoState28Action,
	{_State133, OptionalLabelDeclType}:            _GotoState134Action,
	{_State133, BlockExprType}:                    _GotoState42Action,
	{_State133, CallExprType}:                     _GotoState43Action,
	{_State133, AtomExprType}:                     _GotoState41Action,
	{_State133, LiteralType}:                      _GotoState48Action,
	{_State133, InitializeExprType}:               _GotoState47Action,
	{_State133, AccessExprType}:                   _GotoState37Action,
	{_State133, PostfixUnaryExprType}:             _GotoState52Action,
	{_State133, PrefixUnaryOpType}:                _GotoState54Action,
	{_State133, PrefixUnaryExprType}:              _GotoState53Action,
	{_State133, MulExprType}:                      _GotoState49Action,
	{_State133, AddExprType}:                      _GotoState38Action,
	{_State133, CmpExprType}:                      _GotoState44Action,
	{_State133, AndExprType}:                      _GotoState200Action,
	{_State133, InitializableTypeType}:            _GotoState46Action,
	{_State133, ExplicitStructDefType}:            _GotoState45Action,
	{_State133, AnonymousFuncExprType}:            _GotoState40Action,
	{_State134, LbraceToken}:                      _GotoState127Action,
	{_State134, BlockBodyType}:                    _GotoState129Action,
	{_State136, IdentifierToken}:                  _GotoState201Action,
	{_State137, PackageToken}:                     _GotoState16Action,
	{_State137, UnsafeToken}:                      _GotoState58Action,
	{_State137, TypeToken}:                        _GotoState17Action,
	{_State137, FuncToken}:                        _GotoState18Action,
	{_State137, DefinitionType}:                   _GotoState202Action,
	{_State137, UnsafeStatementType}:              _GotoState64Action,
	{_State137, TypeDefType}:                      _GotoState63Action,
	{_State137, NamedFuncDefType}:                 _GotoState61Action,
	{_State137, PackageDefType}:                   _GotoState62Action,
	{_State139, PackageStatementsType}:            _GotoState203Action,
	{_State140, IdentifierToken}:                  _GotoState204Action,
	{_State140, GenericParameterDefType}:          _GotoState205Action,
	{_State140, GenericParameterDefsType}:         _GotoState206Action,
	{_State140, OptionalGenericParameterDefsType}: _GotoState207Action,
	{_State141, IdentifierToken}:                  _GotoState76Action,
	{_State141, StructToken}:                      _GotoState34Action,
	{_State141, EnumToken}:                        _GotoState73Action,
	{_State141, TraitToken}:                       _GotoState81Action,
	{_State141, FuncToken}:                        _GotoState75Action,
	{_State141, LparenToken}:                      _GotoState78Action,
	{_State141, LbracketToken}:                    _GotoState27Action,
	{_State141, DotToken}:                         _GotoState72Action,
	{_State141, QuestionToken}:                    _GotoState79Action,
	{_State141, ExclaimToken}:                     _GotoState74Action,
	{_State141, TildeTildeToken}:                  _GotoState80Action,
	{_State141, BitNegToken}:                      _GotoState71Action,
	{_State141, BitAndToken}:                      _GotoState70Action,
	{_State141, LexErrorToken}:                    _GotoState77Action,
	{_State141, InitializableTypeType}:            _GotoState87Action,
	{_State141, AtomTypeType}:                     _GotoState82Action,
	{_State141, ReturnableTypeType}:               _GotoState88Action,
	{_State141, ValueTypeType}:                    _GotoState208Action,
	{_State141, ImplicitStructDefType}:            _GotoState86Action,
	{_State141, ExplicitStructDefType}:            _GotoState45Action,
	{_State141, ImplicitEnumDefType}:              _GotoState85Action,
	{_State141, ExplicitEnumDefType}:              _GotoState83Action,
	{_State141, TraitDefType}:                     _GotoState89Action,
	{_State141, FuncTypeType}:                     _GotoState84Action,
	{_State142, IdentifierToken}:                  _GotoState76Action,
	{_State142, StructToken}:                      _GotoState34Action,
	{_State142, EnumToken}:                        _GotoState73Action,
	{_State142, TraitToken}:                       _GotoState81Action,
	{_State142, FuncToken}:                        _GotoState75Action,
	{_State142, LparenToken}:                      _GotoState78Action,
	{_State142, LbracketToken}:                    _GotoState27Action,
	{_State142, DotToken}:                         _GotoState72Action,
	{_State142, QuestionToken}:                    _GotoState79Action,
	{_State142, ExclaimToken}:                     _GotoState74Action,
	{_State142, TildeTildeToken}:                  _GotoState80Action,
	{_State142, BitNegToken}:                      _GotoState71Action,
	{_State142, BitAndToken}:                      _GotoState70Action,
	{_State142, LexErrorToken}:                    _GotoState77Action,
	{_State142, InitializableTypeType}:            _GotoState87Action,
	{_State142, AtomTypeType}:                     _GotoState82Action,
	{_State142, ReturnableTypeType}:               _GotoState88Action,
	{_State142, ValueTypeType}:                    _GotoState209Action,
	{_State142, ImplicitStructDefType}:            _GotoState86Action,
	{_State142, ExplicitStructDefType}:            _GotoState45Action,
	{_State142, ImplicitEnumDefType}:              _GotoState85Action,
	{_State142, ExplicitEnumDefType}:              _GotoState83Action,
	{_State142, TraitDefType}:                     _GotoState89Action,
	{_State142, FuncTypeType}:                     _GotoState84Action,
	{_State143, IdentifierToken}:                  _GotoState76Action,
	{_State143, StructToken}:                      _GotoState34Action,
	{_State143, EnumToken}:                        _GotoState73Action,
	{_State143, TraitToken}:                       _GotoState81Action,
	{_State143, FuncToken}:                        _GotoState75Action,
	{_State143, LparenToken}:                      _GotoState78Action,
	{_State143, LbracketToken}:                    _GotoState27Action,
	{_State143, DotToken}:                         _GotoState72Action,
	{_State143, QuestionToken}:                    _GotoState79Action,
	{_State143, ExclaimToken}:                     _GotoState74Action,
	{_State143, DotdotdotToken}:                   _GotoState210Action,
	{_State143, TildeTildeToken}:                  _GotoState80Action,
	{_State143, BitNegToken}:                      _GotoState71Action,
	{_State143, BitAndToken}:                      _GotoState70Action,
	{_State143, LexErrorToken}:                    _GotoState77Action,
	{_State143, InitializableTypeType}:            _GotoState87Action,
	{_State143, AtomTypeType}:                     _GotoState82Action,
	{_State143, ReturnableTypeType}:               _GotoState88Action,
	{_State143, ValueTypeType}:                    _GotoState211Action,
	{_State143, ImplicitStructDefType}:            _GotoState86Action,
	{_State143, ExplicitStructDefType}:            _GotoState45Action,
	{_State143, ImplicitEnumDefType}:              _GotoState85Action,
	{_State143, ExplicitEnumDefType}:              _GotoState83Action,
	{_State143, TraitDefType}:                     _GotoState89Action,
	{_State143, FuncTypeType}:                     _GotoState84Action,
	{_State144, RparenToken}:                      _GotoState212Action,
	{_State145, DollarLbracketToken}:              _GotoState140Action,
	{_State145, OptionalGenericParametersType}:    _GotoState213Action,
	{_State146, RparenToken}:                      _GotoState214Action,
	{_State148, CommaToken}:                       _GotoState215Action,
	{_State152, IdentifierToken}:                  _GotoState157Action,
	{_State152, UnsafeToken}:                      _GotoState58Action,
	{_State152, StructToken}:                      _GotoState34Action,
	{_State152, EnumToken}:                        _GotoState73Action,
	{_State152, TraitToken}:                       _GotoState81Action,
	{_State152, FuncToken}:                        _GotoState75Action,
	{_State152, LparenToken}:                      _GotoState78Action,
	{_State152, LbracketToken}:                    _GotoState27Action,
	{_State152, DotToken}:                         _GotoState72Action,
	{_State152, QuestionToken}:                    _GotoState79Action,
	{_State152, ExclaimToken}:                     _GotoState74Action,
	{_State152, TildeTildeToken}:                  _GotoState80Action,
	{_State152, BitNegToken}:                      _GotoState71Action,
	{_State152, BitAndToken}:                      _GotoState70Action,
	{_State152, LexErrorToken}:                    _GotoState77Action,
	{_State152, UnsafeStatementType}:              _GotoState163Action,
	{_State152, InitializableTypeType}:            _GotoState87Action,
	{_State152, AtomTypeType}:                     _GotoState82Action,
	{_State152, ReturnableTypeType}:               _GotoState88Action,
	{_State152, ValueTypeType}:                    _GotoState164Action,
	{_State152, FieldDefType}:                     _GotoState218Action,
	{_State152, ImplicitStructDefType}:            _GotoState86Action,
	{_State152, ExplicitStructDefType}:            _GotoState45Action,
	{_State152, EnumValueDefType}:                 _GotoState216Action,
	{_State152, ImplicitEnumValueDefsType}:        _GotoState219Action,
	{_State152, ImplicitEnumDefType}:              _GotoState85Action,
	{_State152, ExplicitEnumValueDefsType}:        _GotoState217Action,
	{_State152, ExplicitEnumDefType}:              _GotoState83Action,
	{_State152, TraitDefType}:                     _GotoState89Action,
	{_State152, FuncTypeType}:                     _GotoState84Action,
	{_State154, IdentifierToken}:                  _GotoState221Action,
	{_State154, StructToken}:                      _GotoState34Action,
	{_State154, EnumToken}:                        _GotoState73Action,
	{_State154, TraitToken}:                       _GotoState81Action,
	{_State154, FuncToken}:                        _GotoState75Action,
	{_State154, LparenToken}:                      _GotoState78Action,
	{_State154, LbracketToken}:                    _GotoState27Action,
	{_State154, DotToken}:                         _GotoState72Action,
	{_State154, QuestionToken}:                    _GotoState79Action,
	{_State154, ExclaimToken}:                     _GotoState74Action,
	{_State154, DotdotdotToken}:                   _GotoState220Action,
	{_State154, TildeTildeToken}:                  _GotoState80Action,
	{_State154, BitNegToken}:                      _GotoState71Action,
	{_State154, BitAndToken}:                      _GotoState70Action,
	{_State154, LexErrorToken}:                    _GotoState77Action,
	{_State154, InitializableTypeType}:            _GotoState87Action,
	{_State154, AtomTypeType}:                     _GotoState82Action,
	{_State154, ReturnableTypeType}:               _GotoState88Action,
	{_State154, ValueTypeType}:                    _GotoState225Action,
	{_State154, ImplicitStructDefType}:            _GotoState86Action,
	{_State154, ExplicitStructDefType}:            _GotoState45Action,
	{_State154, ImplicitEnumDefType}:              _GotoState85Action,
	{_State154, ExplicitEnumDefType}:              _GotoState83Action,
	{_State154, TraitDefType}:                     _GotoState89Action,
	{_State154, ParameterDeclType}:                _GotoState223Action,
	{_State154, ParameterDeclsType}:               _GotoState224Action,
	{_State154, OptionalParameterDeclsType}:       _GotoState222Action,
	{_State154, FuncTypeType}:                     _GotoState84Action,
	{_State155, IdentifierToken}:                  _GotoState226Action,
	{_State157, IdentifierToken}:                  _GotoState76Action,
	{_State157, StructToken}:                      _GotoState34Action,
	{_State157, EnumToken}:                        _GotoState73Action,
	{_State157, TraitToken}:                       _GotoState81Action,
	{_State157, FuncToken}:                        _GotoState75Action,
	{_State157, LparenToken}:                      _GotoState78Action,
	{_State157, LbracketToken}:                    _GotoState27Action,
	{_State157, DotToken}:                         _GotoState227Action,
	{_State157, QuestionToken}:                    _GotoState79Action,
	{_State157, ExclaimToken}:                     _GotoState74Action,
	{_State157, DollarLbracketToken}:              _GotoState98Action,
	{_State157, TildeTildeToken}:                  _GotoState80Action,
	{_State157, BitNegToken}:                      _GotoState71Action,
	{_State157, BitAndToken}:                      _GotoState70Action,
	{_State157, LexErrorToken}:                    _GotoState77Action,
	{_State157, OptionalGenericBindingType}:       _GotoState156Action,
	{_State157, InitializableTypeType}:            _GotoState87Action,
	{_State157, AtomTypeType}:                     _GotoState82Action,
	{_State157, ReturnableTypeType}:               _GotoState88Action,
	{_State157, ValueTypeType}:                    _GotoState228Action,
	{_State157, ImplicitStructDefType}:            _GotoState86Action,
	{_State157, ExplicitStructDefType}:            _GotoState45Action,
	{_State157, ImplicitEnumDefType}:              _GotoState85Action,
	{_State157, ExplicitEnumDefType}:              _GotoState83Action,
	{_State157, TraitDefType}:                     _GotoState89Action,
	{_State157, FuncTypeType}:                     _GotoState84Action,
	{_State158, OrToken}:                          _GotoState229Action,
	{_State159, AssignToken}:                      _GotoState230Action,
	{_State160, RparenToken}:                      _GotoState232Action,
	{_State160, OrToken}:                          _GotoState231Action,
	{_State161, CommaToken}:                       _GotoState233Action,
	{_State162, RparenToken}:                      _GotoState234Action,
	{_State164, AddToken}:                         _GotoState168Action,
	{_State164, SubToken}:                         _GotoState173Action,
	{_State164, MulToken}:                         _GotoState171Action,
	{_State167, IdentifierToken}:                  _GotoState157Action,
	{_State167, UnsafeToken}:                      _GotoState58Action,
	{_State167, StructToken}:                      _GotoState34Action,
	{_State167, EnumToken}:                        _GotoState73Action,
	{_State167, TraitToken}:                       _GotoState81Action,
	{_State167, FuncToken}:                        _GotoState235Action,
	{_State167, LparenToken}:                      _GotoState78Action,
	{_State167, LbracketToken}:                    _GotoState27Action,
	{_State167, DotToken}:                         _GotoState72Action,
	{_State167, QuestionToken}:                    _GotoState79Action,
	{_State167, ExclaimToken}:                     _GotoState74Action,
	{_State167, TildeTildeToken}:                  _GotoState80Action,
	{_State167, BitNegToken}:                      _GotoState71Action,
	{_State167, BitAndToken}:                      _GotoState70Action,
	{_State167, LexErrorToken}:                    _GotoState77Action,
	{_State167, UnsafeStatementType}:              _GotoState163Action,
	{_State167, InitializableTypeType}:            _GotoState87Action,
	{_State167, AtomTypeType}:                     _GotoState82Action,
	{_State167, ReturnableTypeType}:               _GotoState88Action,
	{_State167, ValueTypeType}:                    _GotoState164Action,
	{_State167, FieldDefType}:                     _GotoState236Action,
	{_State167, ImplicitStructDefType}:            _GotoState86Action,
	{_State167, ExplicitStructDefType}:            _GotoState45Action,
	{_State167, ImplicitEnumDefType}:              _GotoState85Action,
	{_State167, ExplicitEnumDefType}:              _GotoState83Action,
	{_State167, TraitPropertyType}:                _GotoState240Action,
	{_State167, TraitPropertiesType}:              _GotoState239Action,
	{_State167, OptionalTraitPropertiesType}:      _GotoState238Action,
	{_State167, TraitDefType}:                     _GotoState89Action,
	{_State167, FuncTypeType}:                     _GotoState84Action,
	{_State167, MethodSignatureType}:              _GotoState237Action,
	{_State168, IdentifierToken}:                  _GotoState76Action,
	{_State168, StructToken}:                      _GotoState34Action,
	{_State168, EnumToken}:                        _GotoState73Action,
	{_State168, TraitToken}:                       _GotoState81Action,
	{_State168, FuncToken}:                        _GotoState75Action,
	{_State168, LparenToken}:                      _GotoState78Action,
	{_State168, LbracketToken}:                    _GotoState27Action,
	{_State168, DotToken}:                         _GotoState72Action,
	{_State168, QuestionToken}:                    _GotoState79Action,
	{_State168, ExclaimToken}:                     _GotoState74Action,
	{_State168, TildeTildeToken}:                  _GotoState80Action,
	{_State168, BitNegToken}:                      _GotoState71Action,
	{_State168, BitAndToken}:                      _GotoState70Action,
	{_State168, LexErrorToken}:                    _GotoState77Action,
	{_State168, InitializableTypeType}:            _GotoState87Action,
	{_State168, AtomTypeType}:                     _GotoState82Action,
	{_State168, ReturnableTypeType}:               _GotoState241Action,
	{_State168, ImplicitStructDefType}:            _GotoState86Action,
	{_State168, ExplicitStructDefType}:            _GotoState45Action,
	{_State168, ImplicitEnumDefType}:              _GotoState85Action,
	{_State168, ExplicitEnumDefType}:              _GotoState83Action,
	{_State168, TraitDefType}:                     _GotoState89Action,
	{_State168, FuncTypeType}:                     _GotoState84Action,
	{_State169, IdentifierToken}:                  _GotoState76Action,
	{_State169, StructToken}:                      _GotoState34Action,
	{_State169, EnumToken}:                        _GotoState73Action,
	{_State169, TraitToken}:                       _GotoState81Action,
	{_State169, FuncToken}:                        _GotoState75Action,
	{_State169, LparenToken}:                      _GotoState78Action,
	{_State169, LbracketToken}:                    _GotoState27Action,
	{_State169, DotToken}:                         _GotoState72Action,
	{_State169, QuestionToken}:                    _GotoState79Action,
	{_State169, ExclaimToken}:                     _GotoState74Action,
	{_State169, TildeTildeToken}:                  _GotoState80Action,
	{_State169, BitNegToken}:                      _GotoState71Action,
	{_State169, BitAndToken}:                      _GotoState70Action,
	{_State169, LexErrorToken}:                    _GotoState77Action,
	{_State169, InitializableTypeType}:            _GotoState87Action,
	{_State169, AtomTypeType}:                     _GotoState82Action,
	{_State169, ReturnableTypeType}:               _GotoState88Action,
	{_State169, ValueTypeType}:                    _GotoState242Action,
	{_State169, ImplicitStructDefType}:            _GotoState86Action,
	{_State169, ExplicitStructDefType}:            _GotoState45Action,
	{_State169, ImplicitEnumDefType}:              _GotoState85Action,
	{_State169, ExplicitEnumDefType}:              _GotoState83Action,
	{_State169, TraitDefType}:                     _GotoState89Action,
	{_State169, FuncTypeType}:                     _GotoState84Action,
	{_State170, IntegerLiteralToken}:              _GotoState243Action,
	{_State171, IdentifierToken}:                  _GotoState76Action,
	{_State171, StructToken}:                      _GotoState34Action,
	{_State171, EnumToken}:                        _GotoState73Action,
	{_State171, TraitToken}:                       _GotoState81Action,
	{_State171, FuncToken}:                        _GotoState75Action,
	{_State171, LparenToken}:                      _GotoState78Action,
	{_State171, LbracketToken}:                    _GotoState27Action,
	{_State171, DotToken}:                         _GotoState72Action,
	{_State171, QuestionToken}:                    _GotoState79Action,
	{_State171, ExclaimToken}:                     _GotoState74Action,
	{_State171, TildeTildeToken}:                  _GotoState80Action,
	{_State171, BitNegToken}:                      _GotoState71Action,
	{_State171, BitAndToken}:                      _GotoState70Action,
	{_State171, LexErrorToken}:                    _GotoState77Action,
	{_State171, InitializableTypeType}:            _GotoState87Action,
	{_State171, AtomTypeType}:                     _GotoState82Action,
	{_State171, ReturnableTypeType}:               _GotoState244Action,
	{_State171, ImplicitStructDefType}:            _GotoState86Action,
	{_State171, ExplicitStructDefType}:            _GotoState45Action,
	{_State171, ImplicitEnumDefType}:              _GotoState85Action,
	{_State171, ExplicitEnumDefType}:              _GotoState83Action,
	{_State171, TraitDefType}:                     _GotoState89Action,
	{_State171, FuncTypeType}:                     _GotoState84Action,
	{_State173, IdentifierToken}:                  _GotoState76Action,
	{_State173, StructToken}:                      _GotoState34Action,
	{_State173, EnumToken}:                        _GotoState73Action,
	{_State173, TraitToken}:                       _GotoState81Action,
	{_State173, FuncToken}:                        _GotoState75Action,
	{_State173, LparenToken}:                      _GotoState78Action,
	{_State173, LbracketToken}:                    _GotoState27Action,
	{_State173, DotToken}:                         _GotoState72Action,
	{_State173, QuestionToken}:                    _GotoState79Action,
	{_State173, ExclaimToken}:                     _GotoState74Action,
	{_State173, TildeTildeToken}:                  _GotoState80Action,
	{_State173, BitNegToken}:                      _GotoState71Action,
	{_State173, BitAndToken}:                      _GotoState70Action,
	{_State173, LexErrorToken}:                    _GotoState77Action,
	{_State173, InitializableTypeType}:            _GotoState87Action,
	{_State173, AtomTypeType}:                     _GotoState82Action,
	{_State173, ReturnableTypeType}:               _GotoState245Action,
	{_State173, ImplicitStructDefType}:            _GotoState86Action,
	{_State173, ExplicitStructDefType}:            _GotoState45Action,
	{_State173, ImplicitEnumDefType}:              _GotoState85Action,
	{_State173, ExplicitEnumDefType}:              _GotoState83Action,
	{_State173, TraitDefType}:                     _GotoState89Action,
	{_State173, FuncTypeType}:                     _GotoState84Action,
	{_State174, IntegerLiteralToken}:              _GotoState25Action,
	{_State174, FloatLiteralToken}:                _GotoState22Action,
	{_State174, RuneLiteralToken}:                 _GotoState32Action,
	{_State174, StringLiteralToken}:               _GotoState33Action,
	{_State174, IdentifierToken}:                  _GotoState24Action,
	{_State174, TrueToken}:                        _GotoState36Action,
	{_State174, FalseToken}:                       _GotoState21Action,
	{_State174, StructToken}:                      _GotoState34Action,
	{_State174, FuncToken}:                        _GotoState23Action,
	{_State174, LabelDeclToken}:                   _GotoState26Action,
	{_State174, LparenToken}:                      _GotoState29Action,
	{_State174, LbracketToken}:                    _GotoState27Action,
	{_State174, NotToken}:                         _GotoState31Action,
	{_State174, SubToken}:                         _GotoState35Action,
	{_State174, MulToken}:                         _GotoState30Action,
	{_State174, BitNegToken}:                      _GotoState20Action,
	{_State174, BitAndToken}:                      _GotoState19Action,
	{_State174, LexErrorToken}:                    _GotoState28Action,
	{_State174, ExpressionType}:                   _GotoState246Action,
	{_State174, OptionalLabelDeclType}:            _GotoState50Action,
	{_State174, SequenceExprType}:                 _GotoState55Action,
	{_State174, BlockExprType}:                    _GotoState42Action,
	{_State174, CallExprType}:                     _GotoState43Action,
	{_State174, AtomExprType}:                     _GotoState41Action,
	{_State174, LiteralType}:                      _GotoState48Action,
	{_State174, InitializeExprType}:               _GotoState47Action,
	{_State174, AccessExprType}:                   _GotoState37Action,
	{_State174, PostfixUnaryExprType}:             _GotoState52Action,
	{_State174, PrefixUnaryOpType}:                _GotoState54Action,
	{_State174, PrefixUnaryExprType}:              _GotoState53Action,
	{_State174, MulExprType}:                      _GotoState49Action,
	{_State174, AddExprType}:                      _GotoState38Action,
	{_State174, CmpExprType}:                      _GotoState44Action,
	{_State174, AndExprType}:                      _GotoState39Action,
	{_State174, OrExprType}:                       _GotoState51Action,
	{_State174, InitializableTypeType}:            _GotoState46Action,
	{_State174, ExplicitStructDefType}:            _GotoState45Action,
	{_State174, AnonymousFuncExprType}:            _GotoState40Action,
	{_State175, IntegerLiteralToken}:              _GotoState25Action,
	{_State175, FloatLiteralToken}:                _GotoState22Action,
	{_State175, RuneLiteralToken}:                 _GotoState32Action,
	{_State175, StringLiteralToken}:               _GotoState33Action,
	{_State175, IdentifierToken}:                  _GotoState91Action,
	{_State175, TrueToken}:                        _GotoState36Action,
	{_State175, FalseToken}:                       _GotoState21Action,
	{_State175, StructToken}:                      _GotoState34Action,
	{_State175, FuncToken}:                        _GotoState23Action,
	{_State175, LabelDeclToken}:                   _GotoState26Action,
	{_State175, LparenToken}:                      _GotoState29Action,
	{_State175, LbracketToken}:                    _GotoState27Action,
	{_State175, NotToken}:                         _GotoState31Action,
	{_State175, SubToken}:                         _GotoState35Action,
	{_State175, MulToken}:                         _GotoState30Action,
	{_State175, BitNegToken}:                      _GotoState20Action,
	{_State175, BitAndToken}:                      _GotoState19Action,
	{_State175, LexErrorToken}:                    _GotoState28Action,
	{_State175, ExpressionType}:                   _GotoState95Action,
	{_State175, OptionalLabelDeclType}:            _GotoState50Action,
	{_State175, SequenceExprType}:                 _GotoState55Action,
	{_State175, BlockExprType}:                    _GotoState42Action,
	{_State175, CallExprType}:                     _GotoState43Action,
	{_State175, ArgumentType}:                     _GotoState247Action,
	{_State175, ColonExpressionsType}:             _GotoState94Action,
	{_State175, OptionalExpressionType}:           _GotoState96Action,
	{_State175, AtomExprType}:                     _GotoState41Action,
	{_State175, LiteralType}:                      _GotoState48Action,
	{_State175, InitializeExprType}:               _GotoState47Action,
	{_State175, AccessExprType}:                   _GotoState37Action,
	{_State175, PostfixUnaryExprType}:             _GotoState52Action,
	{_State175, PrefixUnaryOpType}:                _GotoState54Action,
	{_State175, PrefixUnaryExprType}:              _GotoState53Action,
	{_State175, MulExprType}:                      _GotoState49Action,
	{_State175, AddExprType}:                      _GotoState38Action,
	{_State175, CmpExprType}:                      _GotoState44Action,
	{_State175, AndExprType}:                      _GotoState39Action,
	{_State175, OrExprType}:                       _GotoState51Action,
	{_State175, InitializableTypeType}:            _GotoState46Action,
	{_State175, ExplicitStructDefType}:            _GotoState45Action,
	{_State175, AnonymousFuncExprType}:            _GotoState40Action,
	{_State177, IntegerLiteralToken}:              _GotoState25Action,
	{_State177, FloatLiteralToken}:                _GotoState22Action,
	{_State177, RuneLiteralToken}:                 _GotoState32Action,
	{_State177, StringLiteralToken}:               _GotoState33Action,
	{_State177, IdentifierToken}:                  _GotoState24Action,
	{_State177, TrueToken}:                        _GotoState36Action,
	{_State177, FalseToken}:                       _GotoState21Action,
	{_State177, StructToken}:                      _GotoState34Action,
	{_State177, FuncToken}:                        _GotoState23Action,
	{_State177, LabelDeclToken}:                   _GotoState26Action,
	{_State177, LparenToken}:                      _GotoState29Action,
	{_State177, LbracketToken}:                    _GotoState27Action,
	{_State177, NotToken}:                         _GotoState31Action,
	{_State177, SubToken}:                         _GotoState35Action,
	{_State177, MulToken}:                         _GotoState30Action,
	{_State177, BitNegToken}:                      _GotoState20Action,
	{_State177, BitAndToken}:                      _GotoState19Action,
	{_State177, LexErrorToken}:                    _GotoState28Action,
	{_State177, ExpressionType}:                   _GotoState248Action,
	{_State177, OptionalLabelDeclType}:            _GotoState50Action,
	{_State177, SequenceExprType}:                 _GotoState55Action,
	{_State177, BlockExprType}:                    _GotoState42Action,
	{_State177, CallExprType}:                     _GotoState43Action,
	{_State177, OptionalExpressionType}:           _GotoState249Action,
	{_State177, AtomExprType}:                     _GotoState41Action,
	{_State177, LiteralType}:                      _GotoState48Action,
	{_State177, InitializeExprType}:               _GotoState47Action,
	{_State177, AccessExprType}:                   _GotoState37Action,
	{_State177, PostfixUnaryExprType}:             _GotoState52Action,
	{_State177, PrefixUnaryOpType}:                _GotoState54Action,
	{_State177, PrefixUnaryExprType}:              _GotoState53Action,
	{_State177, MulExprType}:                      _GotoState49Action,
	{_State177, AddExprType}:                      _GotoState38Action,
	{_State177, CmpExprType}:                      _GotoState44Action,
	{_State177, AndExprType}:                      _GotoState39Action,
	{_State177, OrExprType}:                       _GotoState51Action,
	{_State177, InitializableTypeType}:            _GotoState46Action,
	{_State177, ExplicitStructDefType}:            _GotoState45Action,
	{_State177, AnonymousFuncExprType}:            _GotoState40Action,
	{_State178, IntegerLiteralToken}:              _GotoState25Action,
	{_State178, FloatLiteralToken}:                _GotoState22Action,
	{_State178, RuneLiteralToken}:                 _GotoState32Action,
	{_State178, StringLiteralToken}:               _GotoState33Action,
	{_State178, IdentifierToken}:                  _GotoState24Action,
	{_State178, TrueToken}:                        _GotoState36Action,
	{_State178, FalseToken}:                       _GotoState21Action,
	{_State178, StructToken}:                      _GotoState34Action,
	{_State178, FuncToken}:                        _GotoState23Action,
	{_State178, LabelDeclToken}:                   _GotoState26Action,
	{_State178, LparenToken}:                      _GotoState29Action,
	{_State178, LbracketToken}:                    _GotoState27Action,
	{_State178, NotToken}:                         _GotoState31Action,
	{_State178, SubToken}:                         _GotoState35Action,
	{_State178, MulToken}:                         _GotoState30Action,
	{_State178, BitNegToken}:                      _GotoState20Action,
	{_State178, BitAndToken}:                      _GotoState19Action,
	{_State178, LexErrorToken}:                    _GotoState28Action,
	{_State178, ExpressionType}:                   _GotoState248Action,
	{_State178, OptionalLabelDeclType}:            _GotoState50Action,
	{_State178, SequenceExprType}:                 _GotoState55Action,
	{_State178, BlockExprType}:                    _GotoState42Action,
	{_State178, CallExprType}:                     _GotoState43Action,
	{_State178, OptionalExpressionType}:           _GotoState250Action,
	{_State178, AtomExprType}:                     _GotoState41Action,
	{_State178, LiteralType}:                      _GotoState48Action,
	{_State178, InitializeExprType}:               _GotoState47Action,
	{_State178, AccessExprType}:                   _GotoState37Action,
	{_State178, PostfixUnaryExprType}:             _GotoState52Action,
	{_State178, PrefixUnaryOpType}:                _GotoState54Action,
	{_State178, PrefixUnaryExprType}:              _GotoState53Action,
	{_State178, MulExprType}:                      _GotoState49Action,
	{_State178, AddExprType}:                      _GotoState38Action,
	{_State178, CmpExprType}:                      _GotoState44Action,
	{_State178, AndExprType}:                      _GotoState39Action,
	{_State178, OrExprType}:                       _GotoState51Action,
	{_State178, InitializableTypeType}:            _GotoState46Action,
	{_State178, ExplicitStructDefType}:            _GotoState45Action,
	{_State178, AnonymousFuncExprType}:            _GotoState40Action,
	{_State179, NewlinesToken}:                    _GotoState252Action,
	{_State179, CommaToken}:                       _GotoState251Action,
	{_State181, RparenToken}:                      _GotoState253Action,
	{_State182, CommaToken}:                       _GotoState254Action,
	{_State183, RbracketToken}:                    _GotoState255Action,
	{_State184, AddToken}:                         _GotoState168Action,
	{_State184, SubToken}:                         _GotoState173Action,
	{_State184, MulToken}:                         _GotoState171Action,
	{_State186, RbracketToken}:                    _GotoState256Action,
	{_State187, IntegerLiteralToken}:              _GotoState25Action,
	{_State187, FloatLiteralToken}:                _GotoState22Action,
	{_State187, RuneLiteralToken}:                 _GotoState32Action,
	{_State187, StringLiteralToken}:               _GotoState33Action,
	{_State187, IdentifierToken}:                  _GotoState91Action,
	{_State187, TrueToken}:                        _GotoState36Action,
	{_State187, FalseToken}:                       _GotoState21Action,
	{_State187, StructToken}:                      _GotoState34Action,
	{_State187, FuncToken}:                        _GotoState23Action,
	{_State187, LabelDeclToken}:                   _GotoState26Action,
	{_State187, LparenToken}:                      _GotoState29Action,
	{_State187, LbracketToken}:                    _GotoState27Action,
	{_State187, NotToken}:                         _GotoState31Action,
	{_State187, SubToken}:                         _GotoState35Action,
	{_State187, MulToken}:                         _GotoState30Action,
	{_State187, BitNegToken}:                      _GotoState20Action,
	{_State187, BitAndToken}:                      _GotoState19Action,
	{_State187, LexErrorToken}:                    _GotoState28Action,
	{_State187, ExpressionType}:                   _GotoState95Action,
	{_State187, OptionalLabelDeclType}:            _GotoState50Action,
	{_State187, SequenceExprType}:                 _GotoState55Action,
	{_State187, BlockExprType}:                    _GotoState42Action,
	{_State187, CallExprType}:                     _GotoState43Action,
	{_State187, OptionalArgumentsType}:            _GotoState258Action,
	{_State187, ArgumentsType}:                    _GotoState257Action,
	{_State187, ArgumentType}:                     _GotoState92Action,
	{_State187, ColonExpressionsType}:             _GotoState94Action,
	{_State187, OptionalExpressionType}:           _GotoState96Action,
	{_State187, AtomExprType}:                     _GotoState41Action,
	{_State187, LiteralType}:                      _GotoState48Action,
	{_State187, InitializeExprType}:               _GotoState47Action,
	{_State187, AccessExprType}:                   _GotoState37Action,
	{_State187, PostfixUnaryExprType}:             _GotoState52Action,
	{_State187, PrefixUnaryOpType}:                _GotoState54Action,
	{_State187, PrefixUnaryExprType}:              _GotoState53Action,
	{_State187, MulExprType}:                      _GotoState49Action,
	{_State187, AddExprType}:                      _GotoState38Action,
	{_State187, CmpExprType}:                      _GotoState44Action,
	{_State187, AndExprType}:                      _GotoState39Action,
	{_State187, OrExprType}:                       _GotoState51Action,
	{_State187, InitializableTypeType}:            _GotoState46Action,
	{_State187, ExplicitStructDefType}:            _GotoState45Action,
	{_State187, AnonymousFuncExprType}:            _GotoState40Action,
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
	{_State191, RparenToken}:                      _GotoState259Action,
	{_State191, CommaToken}:                       _GotoState175Action,
	{_State193, ForToken}:                         _GotoState260Action,
	{_State194, IntegerLiteralToken}:              _GotoState25Action,
	{_State194, FloatLiteralToken}:                _GotoState22Action,
	{_State194, RuneLiteralToken}:                 _GotoState32Action,
	{_State194, StringLiteralToken}:               _GotoState33Action,
	{_State194, IdentifierToken}:                  _GotoState24Action,
	{_State194, TrueToken}:                        _GotoState36Action,
	{_State194, FalseToken}:                       _GotoState21Action,
	{_State194, StructToken}:                      _GotoState34Action,
	{_State194, FuncToken}:                        _GotoState23Action,
	{_State194, LabelDeclToken}:                   _GotoState26Action,
	{_State194, LparenToken}:                      _GotoState29Action,
	{_State194, LbracketToken}:                    _GotoState27Action,
	{_State194, NotToken}:                         _GotoState31Action,
	{_State194, SubToken}:                         _GotoState35Action,
	{_State194, MulToken}:                         _GotoState30Action,
	{_State194, BitNegToken}:                      _GotoState20Action,
	{_State194, BitAndToken}:                      _GotoState19Action,
	{_State194, LexErrorToken}:                    _GotoState28Action,
	{_State194, OptionalLabelDeclType}:            _GotoState134Action,
	{_State194, SequenceExprType}:                 _GotoState261Action,
	{_State194, BlockExprType}:                    _GotoState42Action,
	{_State194, CallExprType}:                     _GotoState43Action,
	{_State194, AtomExprType}:                     _GotoState41Action,
	{_State194, LiteralType}:                      _GotoState48Action,
	{_State194, InitializeExprType}:               _GotoState47Action,
	{_State194, AccessExprType}:                   _GotoState37Action,
	{_State194, PostfixUnaryExprType}:             _GotoState52Action,
	{_State194, PrefixUnaryOpType}:                _GotoState54Action,
	{_State194, PrefixUnaryExprType}:              _GotoState53Action,
	{_State194, MulExprType}:                      _GotoState49Action,
	{_State194, AddExprType}:                      _GotoState38Action,
	{_State194, CmpExprType}:                      _GotoState44Action,
	{_State194, AndExprType}:                      _GotoState39Action,
	{_State194, OrExprType}:                       _GotoState51Action,
	{_State194, InitializableTypeType}:            _GotoState46Action,
	{_State194, ExplicitStructDefType}:            _GotoState45Action,
	{_State194, AnonymousFuncExprType}:            _GotoState40Action,
	{_State195, IntegerLiteralToken}:              _GotoState25Action,
	{_State195, FloatLiteralToken}:                _GotoState22Action,
	{_State195, RuneLiteralToken}:                 _GotoState32Action,
	{_State195, StringLiteralToken}:               _GotoState33Action,
	{_State195, IdentifierToken}:                  _GotoState24Action,
	{_State195, TrueToken}:                        _GotoState36Action,
	{_State195, FalseToken}:                       _GotoState21Action,
	{_State195, StructToken}:                      _GotoState34Action,
	{_State195, FuncToken}:                        _GotoState23Action,
	{_State195, LabelDeclToken}:                   _GotoState26Action,
	{_State195, LparenToken}:                      _GotoState29Action,
	{_State195, LbracketToken}:                    _GotoState27Action,
	{_State195, NotToken}:                         _GotoState31Action,
	{_State195, SubToken}:                         _GotoState35Action,
	{_State195, MulToken}:                         _GotoState30Action,
	{_State195, BitNegToken}:                      _GotoState20Action,
	{_State195, BitAndToken}:                      _GotoState19Action,
	{_State195, LexErrorToken}:                    _GotoState28Action,
	{_State195, OptionalLabelDeclType}:            _GotoState134Action,
	{_State195, OptionalSequenceExprType}:         _GotoState262Action,
	{_State195, SequenceExprType}:                 _GotoState263Action,
	{_State195, BlockExprType}:                    _GotoState42Action,
	{_State195, CallExprType}:                     _GotoState43Action,
	{_State195, AtomExprType}:                     _GotoState41Action,
	{_State195, LiteralType}:                      _GotoState48Action,
	{_State195, InitializeExprType}:               _GotoState47Action,
	{_State195, AccessExprType}:                   _GotoState37Action,
	{_State195, PostfixUnaryExprType}:             _GotoState52Action,
	{_State195, PrefixUnaryOpType}:                _GotoState54Action,
	{_State195, PrefixUnaryExprType}:              _GotoState53Action,
	{_State195, MulExprType}:                      _GotoState49Action,
	{_State195, AddExprType}:                      _GotoState38Action,
	{_State195, CmpExprType}:                      _GotoState44Action,
	{_State195, AndExprType}:                      _GotoState39Action,
	{_State195, OrExprType}:                       _GotoState51Action,
	{_State195, InitializableTypeType}:            _GotoState46Action,
	{_State195, ExplicitStructDefType}:            _GotoState45Action,
	{_State195, AnonymousFuncExprType}:            _GotoState40Action,
	{_State196, DoToken}:                          _GotoState264Action,
	{_State197, LbraceToken}:                      _GotoState127Action,
	{_State197, BlockBodyType}:                    _GotoState265Action,
	{_State198, IntegerLiteralToken}:              _GotoState25Action,
	{_State198, FloatLiteralToken}:                _GotoState22Action,
	{_State198, RuneLiteralToken}:                 _GotoState32Action,
	{_State198, StringLiteralToken}:               _GotoState33Action,
	{_State198, IdentifierToken}:                  _GotoState24Action,
	{_State198, TrueToken}:                        _GotoState36Action,
	{_State198, FalseToken}:                       _GotoState21Action,
	{_State198, ReturnToken}:                      _GotoState271Action,
	{_State198, BreakToken}:                       _GotoState267Action,
	{_State198, ContinueToken}:                    _GotoState268Action,
	{_State198, UnsafeToken}:                      _GotoState58Action,
	{_State198, StructToken}:                      _GotoState34Action,
	{_State198, FuncToken}:                        _GotoState23Action,
	{_State198, AsyncToken}:                       _GotoState266Action,
	{_State198, DeferToken}:                       _GotoState269Action,
	{_State198, LabelDeclToken}:                   _GotoState26Action,
	{_State198, RbraceToken}:                      _GotoState270Action,
	{_State198, LparenToken}:                      _GotoState29Action,
	{_State198, LbracketToken}:                    _GotoState27Action,
	{_State198, NotToken}:                         _GotoState31Action,
	{_State198, SubToken}:                         _GotoState35Action,
	{_State198, MulToken}:                         _GotoState30Action,
	{_State198, BitNegToken}:                      _GotoState20Action,
	{_State198, BitAndToken}:                      _GotoState19Action,
	{_State198, LexErrorToken}:                    _GotoState28Action,
	{_State198, ExpressionType}:                   _GotoState273Action,
	{_State198, OptionalLabelDeclType}:            _GotoState50Action,
	{_State198, SequenceExprType}:                 _GotoState55Action,
	{_State198, BlockExprType}:                    _GotoState42Action,
	{_State198, StatementType}:                    _GotoState277Action,
	{_State198, StatementBodyType}:                _GotoState278Action,
	{_State198, UnsafeStatementType}:              _GotoState279Action,
	{_State198, JumpStatementType}:                _GotoState275Action,
	{_State198, JumpTypeType}:                     _GotoState276Action,
	{_State198, ExpressionsType}:                  _GotoState274Action,
	{_State198, CallExprType}:                     _GotoState43Action,
	{_State198, AtomExprType}:                     _GotoState41Action,
	{_State198, LiteralType}:                      _GotoState48Action,
	{_State198, InitializeExprType}:               _GotoState47Action,
	{_State198, AccessExprType}:                   _GotoState272Action,
	{_State198, PostfixUnaryExprType}:             _GotoState52Action,
	{_State198, PrefixUnaryOpType}:                _GotoState54Action,
	{_State198, PrefixUnaryExprType}:              _GotoState53Action,
	{_State198, MulExprType}:                      _GotoState49Action,
	{_State198, AddExprType}:                      _GotoState38Action,
	{_State198, CmpExprType}:                      _GotoState44Action,
	{_State198, AndExprType}:                      _GotoState39Action,
	{_State198, OrExprType}:                       _GotoState51Action,
	{_State198, InitializableTypeType}:            _GotoState46Action,
	{_State198, ExplicitStructDefType}:            _GotoState45Action,
	{_State198, AnonymousFuncExprType}:            _GotoState40Action,
	{_State199, LbraceToken}:                      _GotoState280Action,
	{_State200, AndToken}:                         _GotoState108Action,
	{_State201, GreaterToken}:                     _GotoState281Action,
	{_State203, UnsafeToken}:                      _GotoState58Action,
	{_State203, RparenToken}:                      _GotoState282Action,
	{_State203, UnsafeStatementType}:              _GotoState285Action,
	{_State203, PackageStatementBodyType}:         _GotoState284Action,
	{_State203, PackageStatementType}:             _GotoState283Action,
	{_State204, IdentifierToken}:                  _GotoState76Action,
	{_State204, StructToken}:                      _GotoState34Action,
	{_State204, EnumToken}:                        _GotoState73Action,
	{_State204, TraitToken}:                       _GotoState81Action,
	{_State204, FuncToken}:                        _GotoState75Action,
	{_State204, LparenToken}:                      _GotoState78Action,
	{_State204, LbracketToken}:                    _GotoState27Action,
	{_State204, DotToken}:                         _GotoState72Action,
	{_State204, QuestionToken}:                    _GotoState79Action,
	{_State204, ExclaimToken}:                     _GotoState74Action,
	{_State204, TildeTildeToken}:                  _GotoState80Action,
	{_State204, BitNegToken}:                      _GotoState71Action,
	{_State204, BitAndToken}:                      _GotoState70Action,
	{_State204, LexErrorToken}:                    _GotoState77Action,
	{_State204, InitializableTypeType}:            _GotoState87Action,
	{_State204, AtomTypeType}:                     _GotoState82Action,
	{_State204, ReturnableTypeType}:               _GotoState88Action,
	{_State204, ValueTypeType}:                    _GotoState286Action,
	{_State204, ImplicitStructDefType}:            _GotoState86Action,
	{_State204, ExplicitStructDefType}:            _GotoState45Action,
	{_State204, ImplicitEnumDefType}:              _GotoState85Action,
	{_State204, ExplicitEnumDefType}:              _GotoState83Action,
	{_State204, TraitDefType}:                     _GotoState89Action,
	{_State204, FuncTypeType}:                     _GotoState84Action,
	{_State206, CommaToken}:                       _GotoState287Action,
	{_State207, RbracketToken}:                    _GotoState288Action,
	{_State208, AddToken}:                         _GotoState168Action,
	{_State208, SubToken}:                         _GotoState173Action,
	{_State208, MulToken}:                         _GotoState171Action,
	{_State209, ImplementsToken}:                  _GotoState289Action,
	{_State209, AddToken}:                         _GotoState168Action,
	{_State209, SubToken}:                         _GotoState173Action,
	{_State209, MulToken}:                         _GotoState171Action,
	{_State210, IdentifierToken}:                  _GotoState76Action,
	{_State210, StructToken}:                      _GotoState34Action,
	{_State210, EnumToken}:                        _GotoState73Action,
	{_State210, TraitToken}:                       _GotoState81Action,
	{_State210, FuncToken}:                        _GotoState75Action,
	{_State210, LparenToken}:                      _GotoState78Action,
	{_State210, LbracketToken}:                    _GotoState27Action,
	{_State210, DotToken}:                         _GotoState72Action,
	{_State210, QuestionToken}:                    _GotoState79Action,
	{_State210, ExclaimToken}:                     _GotoState74Action,
	{_State210, TildeTildeToken}:                  _GotoState80Action,
	{_State210, BitNegToken}:                      _GotoState71Action,
	{_State210, BitAndToken}:                      _GotoState70Action,
	{_State210, LexErrorToken}:                    _GotoState77Action,
	{_State210, InitializableTypeType}:            _GotoState87Action,
	{_State210, AtomTypeType}:                     _GotoState82Action,
	{_State210, ReturnableTypeType}:               _GotoState88Action,
	{_State210, ValueTypeType}:                    _GotoState290Action,
	{_State210, ImplicitStructDefType}:            _GotoState86Action,
	{_State210, ExplicitStructDefType}:            _GotoState45Action,
	{_State210, ImplicitEnumDefType}:              _GotoState85Action,
	{_State210, ExplicitEnumDefType}:              _GotoState83Action,
	{_State210, TraitDefType}:                     _GotoState89Action,
	{_State210, FuncTypeType}:                     _GotoState84Action,
	{_State211, AddToken}:                         _GotoState168Action,
	{_State211, SubToken}:                         _GotoState173Action,
	{_State211, MulToken}:                         _GotoState171Action,
	{_State213, LparenToken}:                      _GotoState291Action,
	{_State214, IdentifierToken}:                  _GotoState76Action,
	{_State214, StructToken}:                      _GotoState34Action,
	{_State214, EnumToken}:                        _GotoState73Action,
	{_State214, TraitToken}:                       _GotoState81Action,
	{_State214, FuncToken}:                        _GotoState75Action,
	{_State214, LparenToken}:                      _GotoState78Action,
	{_State214, LbracketToken}:                    _GotoState27Action,
	{_State214, DotToken}:                         _GotoState72Action,
	{_State214, QuestionToken}:                    _GotoState79Action,
	{_State214, ExclaimToken}:                     _GotoState74Action,
	{_State214, TildeTildeToken}:                  _GotoState80Action,
	{_State214, BitNegToken}:                      _GotoState71Action,
	{_State214, BitAndToken}:                      _GotoState70Action,
	{_State214, LexErrorToken}:                    _GotoState77Action,
	{_State214, InitializableTypeType}:            _GotoState87Action,
	{_State214, AtomTypeType}:                     _GotoState82Action,
	{_State214, ReturnableTypeType}:               _GotoState293Action,
	{_State214, ImplicitStructDefType}:            _GotoState86Action,
	{_State214, ExplicitStructDefType}:            _GotoState45Action,
	{_State214, ImplicitEnumDefType}:              _GotoState85Action,
	{_State214, ExplicitEnumDefType}:              _GotoState83Action,
	{_State214, TraitDefType}:                     _GotoState89Action,
	{_State214, ReturnTypeType}:                   _GotoState292Action,
	{_State214, FuncTypeType}:                     _GotoState84Action,
	{_State215, IdentifierToken}:                  _GotoState143Action,
	{_State215, ParameterDefType}:                 _GotoState294Action,
	{_State216, NewlinesToken}:                    _GotoState295Action,
	{_State216, OrToken}:                          _GotoState296Action,
	{_State217, RparenToken}:                      _GotoState297Action,
	{_State218, AssignToken}:                      _GotoState230Action,
	{_State219, NewlinesToken}:                    _GotoState298Action,
	{_State219, OrToken}:                          _GotoState299Action,
	{_State220, IdentifierToken}:                  _GotoState76Action,
	{_State220, StructToken}:                      _GotoState34Action,
	{_State220, EnumToken}:                        _GotoState73Action,
	{_State220, TraitToken}:                       _GotoState81Action,
	{_State220, FuncToken}:                        _GotoState75Action,
	{_State220, LparenToken}:                      _GotoState78Action,
	{_State220, LbracketToken}:                    _GotoState27Action,
	{_State220, DotToken}:                         _GotoState72Action,
	{_State220, QuestionToken}:                    _GotoState79Action,
	{_State220, ExclaimToken}:                     _GotoState74Action,
	{_State220, TildeTildeToken}:                  _GotoState80Action,
	{_State220, BitNegToken}:                      _GotoState71Action,
	{_State220, BitAndToken}:                      _GotoState70Action,
	{_State220, LexErrorToken}:                    _GotoState77Action,
	{_State220, InitializableTypeType}:            _GotoState87Action,
	{_State220, AtomTypeType}:                     _GotoState82Action,
	{_State220, ReturnableTypeType}:               _GotoState88Action,
	{_State220, ValueTypeType}:                    _GotoState300Action,
	{_State220, ImplicitStructDefType}:            _GotoState86Action,
	{_State220, ExplicitStructDefType}:            _GotoState45Action,
	{_State220, ImplicitEnumDefType}:              _GotoState85Action,
	{_State220, ExplicitEnumDefType}:              _GotoState83Action,
	{_State220, TraitDefType}:                     _GotoState89Action,
	{_State220, FuncTypeType}:                     _GotoState84Action,
	{_State221, IdentifierToken}:                  _GotoState76Action,
	{_State221, StructToken}:                      _GotoState34Action,
	{_State221, EnumToken}:                        _GotoState73Action,
	{_State221, TraitToken}:                       _GotoState81Action,
	{_State221, FuncToken}:                        _GotoState75Action,
	{_State221, LparenToken}:                      _GotoState78Action,
	{_State221, LbracketToken}:                    _GotoState27Action,
	{_State221, DotToken}:                         _GotoState227Action,
	{_State221, QuestionToken}:                    _GotoState79Action,
	{_State221, ExclaimToken}:                     _GotoState74Action,
	{_State221, DollarLbracketToken}:              _GotoState98Action,
	{_State221, DotdotdotToken}:                   _GotoState301Action,
	{_State221, TildeTildeToken}:                  _GotoState80Action,
	{_State221, BitNegToken}:                      _GotoState71Action,
	{_State221, BitAndToken}:                      _GotoState70Action,
	{_State221, LexErrorToken}:                    _GotoState77Action,
	{_State221, OptionalGenericBindingType}:       _GotoState156Action,
	{_State221, InitializableTypeType}:            _GotoState87Action,
	{_State221, AtomTypeType}:                     _GotoState82Action,
	{_State221, ReturnableTypeType}:               _GotoState88Action,
	{_State221, ValueTypeType}:                    _GotoState302Action,
	{_State221, ImplicitStructDefType}:            _GotoState86Action,
	{_State221, ExplicitStructDefType}:            _GotoState45Action,
	{_State221, ImplicitEnumDefType}:              _GotoState85Action,
	{_State221, ExplicitEnumDefType}:              _GotoState83Action,
	{_State221, TraitDefType}:                     _GotoState89Action,
	{_State221, FuncTypeType}:                     _GotoState84Action,
	{_State222, RparenToken}:                      _GotoState303Action,
	{_State224, CommaToken}:                       _GotoState304Action,
	{_State225, AddToken}:                         _GotoState168Action,
	{_State225, SubToken}:                         _GotoState173Action,
	{_State225, MulToken}:                         _GotoState171Action,
	{_State226, DollarLbracketToken}:              _GotoState98Action,
	{_State226, OptionalGenericBindingType}:       _GotoState305Action,
	{_State227, IdentifierToken}:                  _GotoState226Action,
	{_State227, DollarLbracketToken}:              _GotoState98Action,
	{_State227, OptionalGenericBindingType}:       _GotoState151Action,
	{_State228, AddToken}:                         _GotoState168Action,
	{_State228, SubToken}:                         _GotoState173Action,
	{_State228, MulToken}:                         _GotoState171Action,
	{_State229, IdentifierToken}:                  _GotoState157Action,
	{_State229, UnsafeToken}:                      _GotoState58Action,
	{_State229, StructToken}:                      _GotoState34Action,
	{_State229, EnumToken}:                        _GotoState73Action,
	{_State229, TraitToken}:                       _GotoState81Action,
	{_State229, FuncToken}:                        _GotoState75Action,
	{_State229, LparenToken}:                      _GotoState78Action,
	{_State229, LbracketToken}:                    _GotoState27Action,
	{_State229, DotToken}:                         _GotoState72Action,
	{_State229, QuestionToken}:                    _GotoState79Action,
	{_State229, ExclaimToken}:                     _GotoState74Action,
	{_State229, TildeTildeToken}:                  _GotoState80Action,
	{_State229, BitNegToken}:                      _GotoState71Action,
	{_State229, BitAndToken}:                      _GotoState70Action,
	{_State229, LexErrorToken}:                    _GotoState77Action,
	{_State229, UnsafeStatementType}:              _GotoState163Action,
	{_State229, InitializableTypeType}:            _GotoState87Action,
	{_State229, AtomTypeType}:                     _GotoState82Action,
	{_State229, ReturnableTypeType}:               _GotoState88Action,
	{_State229, ValueTypeType}:                    _GotoState164Action,
	{_State229, FieldDefType}:                     _GotoState218Action,
	{_State229, ImplicitStructDefType}:            _GotoState86Action,
	{_State229, ExplicitStructDefType}:            _GotoState45Action,
	{_State229, EnumValueDefType}:                 _GotoState306Action,
	{_State229, ImplicitEnumDefType}:              _GotoState85Action,
	{_State229, ExplicitEnumDefType}:              _GotoState83Action,
	{_State229, TraitDefType}:                     _GotoState89Action,
	{_State229, FuncTypeType}:                     _GotoState84Action,
	{_State230, DefaultToken}:                     _GotoState307Action,
	{_State231, IdentifierToken}:                  _GotoState157Action,
	{_State231, UnsafeToken}:                      _GotoState58Action,
	{_State231, StructToken}:                      _GotoState34Action,
	{_State231, EnumToken}:                        _GotoState73Action,
	{_State231, TraitToken}:                       _GotoState81Action,
	{_State231, FuncToken}:                        _GotoState75Action,
	{_State231, LparenToken}:                      _GotoState78Action,
	{_State231, LbracketToken}:                    _GotoState27Action,
	{_State231, DotToken}:                         _GotoState72Action,
	{_State231, QuestionToken}:                    _GotoState79Action,
	{_State231, ExclaimToken}:                     _GotoState74Action,
	{_State231, TildeTildeToken}:                  _GotoState80Action,
	{_State231, BitNegToken}:                      _GotoState71Action,
	{_State231, BitAndToken}:                      _GotoState70Action,
	{_State231, LexErrorToken}:                    _GotoState77Action,
	{_State231, UnsafeStatementType}:              _GotoState163Action,
	{_State231, InitializableTypeType}:            _GotoState87Action,
	{_State231, AtomTypeType}:                     _GotoState82Action,
	{_State231, ReturnableTypeType}:               _GotoState88Action,
	{_State231, ValueTypeType}:                    _GotoState164Action,
	{_State231, FieldDefType}:                     _GotoState218Action,
	{_State231, ImplicitStructDefType}:            _GotoState86Action,
	{_State231, ExplicitStructDefType}:            _GotoState45Action,
	{_State231, EnumValueDefType}:                 _GotoState308Action,
	{_State231, ImplicitEnumDefType}:              _GotoState85Action,
	{_State231, ExplicitEnumDefType}:              _GotoState83Action,
	{_State231, TraitDefType}:                     _GotoState89Action,
	{_State231, FuncTypeType}:                     _GotoState84Action,
	{_State233, IdentifierToken}:                  _GotoState157Action,
	{_State233, UnsafeToken}:                      _GotoState58Action,
	{_State233, StructToken}:                      _GotoState34Action,
	{_State233, EnumToken}:                        _GotoState73Action,
	{_State233, TraitToken}:                       _GotoState81Action,
	{_State233, FuncToken}:                        _GotoState75Action,
	{_State233, LparenToken}:                      _GotoState78Action,
	{_State233, LbracketToken}:                    _GotoState27Action,
	{_State233, DotToken}:                         _GotoState72Action,
	{_State233, QuestionToken}:                    _GotoState79Action,
	{_State233, ExclaimToken}:                     _GotoState74Action,
	{_State233, TildeTildeToken}:                  _GotoState80Action,
	{_State233, BitNegToken}:                      _GotoState71Action,
	{_State233, BitAndToken}:                      _GotoState70Action,
	{_State233, LexErrorToken}:                    _GotoState77Action,
	{_State233, UnsafeStatementType}:              _GotoState163Action,
	{_State233, InitializableTypeType}:            _GotoState87Action,
	{_State233, AtomTypeType}:                     _GotoState82Action,
	{_State233, ReturnableTypeType}:               _GotoState88Action,
	{_State233, ValueTypeType}:                    _GotoState164Action,
	{_State233, FieldDefType}:                     _GotoState309Action,
	{_State233, ImplicitStructDefType}:            _GotoState86Action,
	{_State233, ExplicitStructDefType}:            _GotoState45Action,
	{_State233, ImplicitEnumDefType}:              _GotoState85Action,
	{_State233, ExplicitEnumDefType}:              _GotoState83Action,
	{_State233, TraitDefType}:                     _GotoState89Action,
	{_State233, FuncTypeType}:                     _GotoState84Action,
	{_State235, IdentifierToken}:                  _GotoState310Action,
	{_State235, LparenToken}:                      _GotoState154Action,
	{_State238, RparenToken}:                      _GotoState311Action,
	{_State239, NewlinesToken}:                    _GotoState313Action,
	{_State239, CommaToken}:                       _GotoState312Action,
	{_State242, RbracketToken}:                    _GotoState314Action,
	{_State242, AddToken}:                         _GotoState168Action,
	{_State242, SubToken}:                         _GotoState173Action,
	{_State242, MulToken}:                         _GotoState171Action,
	{_State243, RbracketToken}:                    _GotoState315Action,
	{_State251, IdentifierToken}:                  _GotoState157Action,
	{_State251, UnsafeToken}:                      _GotoState58Action,
	{_State251, StructToken}:                      _GotoState34Action,
	{_State251, EnumToken}:                        _GotoState73Action,
	{_State251, TraitToken}:                       _GotoState81Action,
	{_State251, FuncToken}:                        _GotoState75Action,
	{_State251, LparenToken}:                      _GotoState78Action,
	{_State251, LbracketToken}:                    _GotoState27Action,
	{_State251, DotToken}:                         _GotoState72Action,
	{_State251, QuestionToken}:                    _GotoState79Action,
	{_State251, ExclaimToken}:                     _GotoState74Action,
	{_State251, TildeTildeToken}:                  _GotoState80Action,
	{_State251, BitNegToken}:                      _GotoState71Action,
	{_State251, BitAndToken}:                      _GotoState70Action,
	{_State251, LexErrorToken}:                    _GotoState77Action,
	{_State251, UnsafeStatementType}:              _GotoState163Action,
	{_State251, InitializableTypeType}:            _GotoState87Action,
	{_State251, AtomTypeType}:                     _GotoState82Action,
	{_State251, ReturnableTypeType}:               _GotoState88Action,
	{_State251, ValueTypeType}:                    _GotoState164Action,
	{_State251, FieldDefType}:                     _GotoState316Action,
	{_State251, ImplicitStructDefType}:            _GotoState86Action,
	{_State251, ExplicitStructDefType}:            _GotoState45Action,
	{_State251, ImplicitEnumDefType}:              _GotoState85Action,
	{_State251, ExplicitEnumDefType}:              _GotoState83Action,
	{_State251, TraitDefType}:                     _GotoState89Action,
	{_State251, FuncTypeType}:                     _GotoState84Action,
	{_State252, IdentifierToken}:                  _GotoState157Action,
	{_State252, UnsafeToken}:                      _GotoState58Action,
	{_State252, StructToken}:                      _GotoState34Action,
	{_State252, EnumToken}:                        _GotoState73Action,
	{_State252, TraitToken}:                       _GotoState81Action,
	{_State252, FuncToken}:                        _GotoState75Action,
	{_State252, LparenToken}:                      _GotoState78Action,
	{_State252, LbracketToken}:                    _GotoState27Action,
	{_State252, DotToken}:                         _GotoState72Action,
	{_State252, QuestionToken}:                    _GotoState79Action,
	{_State252, ExclaimToken}:                     _GotoState74Action,
	{_State252, TildeTildeToken}:                  _GotoState80Action,
	{_State252, BitNegToken}:                      _GotoState71Action,
	{_State252, BitAndToken}:                      _GotoState70Action,
	{_State252, LexErrorToken}:                    _GotoState77Action,
	{_State252, UnsafeStatementType}:              _GotoState163Action,
	{_State252, InitializableTypeType}:            _GotoState87Action,
	{_State252, AtomTypeType}:                     _GotoState82Action,
	{_State252, ReturnableTypeType}:               _GotoState88Action,
	{_State252, ValueTypeType}:                    _GotoState164Action,
	{_State252, FieldDefType}:                     _GotoState317Action,
	{_State252, ImplicitStructDefType}:            _GotoState86Action,
	{_State252, ExplicitStructDefType}:            _GotoState45Action,
	{_State252, ImplicitEnumDefType}:              _GotoState85Action,
	{_State252, ExplicitEnumDefType}:              _GotoState83Action,
	{_State252, TraitDefType}:                     _GotoState89Action,
	{_State252, FuncTypeType}:                     _GotoState84Action,
	{_State254, IdentifierToken}:                  _GotoState76Action,
	{_State254, StructToken}:                      _GotoState34Action,
	{_State254, EnumToken}:                        _GotoState73Action,
	{_State254, TraitToken}:                       _GotoState81Action,
	{_State254, FuncToken}:                        _GotoState75Action,
	{_State254, LparenToken}:                      _GotoState78Action,
	{_State254, LbracketToken}:                    _GotoState27Action,
	{_State254, DotToken}:                         _GotoState72Action,
	{_State254, QuestionToken}:                    _GotoState79Action,
	{_State254, ExclaimToken}:                     _GotoState74Action,
	{_State254, TildeTildeToken}:                  _GotoState80Action,
	{_State254, BitNegToken}:                      _GotoState71Action,
	{_State254, BitAndToken}:                      _GotoState70Action,
	{_State254, LexErrorToken}:                    _GotoState77Action,
	{_State254, InitializableTypeType}:            _GotoState87Action,
	{_State254, AtomTypeType}:                     _GotoState82Action,
	{_State254, ReturnableTypeType}:               _GotoState88Action,
	{_State254, ValueTypeType}:                    _GotoState318Action,
	{_State254, ImplicitStructDefType}:            _GotoState86Action,
	{_State254, ExplicitStructDefType}:            _GotoState45Action,
	{_State254, ImplicitEnumDefType}:              _GotoState85Action,
	{_State254, ExplicitEnumDefType}:              _GotoState83Action,
	{_State254, TraitDefType}:                     _GotoState89Action,
	{_State254, FuncTypeType}:                     _GotoState84Action,
	{_State257, CommaToken}:                       _GotoState175Action,
	{_State258, RparenToken}:                      _GotoState319Action,
	{_State260, IntegerLiteralToken}:              _GotoState25Action,
	{_State260, FloatLiteralToken}:                _GotoState22Action,
	{_State260, RuneLiteralToken}:                 _GotoState32Action,
	{_State260, StringLiteralToken}:               _GotoState33Action,
	{_State260, IdentifierToken}:                  _GotoState24Action,
	{_State260, TrueToken}:                        _GotoState36Action,
	{_State260, FalseToken}:                       _GotoState21Action,
	{_State260, StructToken}:                      _GotoState34Action,
	{_State260, FuncToken}:                        _GotoState23Action,
	{_State260, LabelDeclToken}:                   _GotoState26Action,
	{_State260, LparenToken}:                      _GotoState29Action,
	{_State260, LbracketToken}:                    _GotoState27Action,
	{_State260, NotToken}:                         _GotoState31Action,
	{_State260, SubToken}:                         _GotoState35Action,
	{_State260, MulToken}:                         _GotoState30Action,
	{_State260, BitNegToken}:                      _GotoState20Action,
	{_State260, BitAndToken}:                      _GotoState19Action,
	{_State260, LexErrorToken}:                    _GotoState28Action,
	{_State260, OptionalLabelDeclType}:            _GotoState134Action,
	{_State260, SequenceExprType}:                 _GotoState320Action,
	{_State260, BlockExprType}:                    _GotoState42Action,
	{_State260, CallExprType}:                     _GotoState43Action,
	{_State260, AtomExprType}:                     _GotoState41Action,
	{_State260, LiteralType}:                      _GotoState48Action,
	{_State260, InitializeExprType}:               _GotoState47Action,
	{_State260, AccessExprType}:                   _GotoState37Action,
	{_State260, PostfixUnaryExprType}:             _GotoState52Action,
	{_State260, PrefixUnaryOpType}:                _GotoState54Action,
	{_State260, PrefixUnaryExprType}:              _GotoState53Action,
	{_State260, MulExprType}:                      _GotoState49Action,
	{_State260, AddExprType}:                      _GotoState38Action,
	{_State260, CmpExprType}:                      _GotoState44Action,
	{_State260, AndExprType}:                      _GotoState39Action,
	{_State260, OrExprType}:                       _GotoState51Action,
	{_State260, InitializableTypeType}:            _GotoState46Action,
	{_State260, ExplicitStructDefType}:            _GotoState45Action,
	{_State260, AnonymousFuncExprType}:            _GotoState40Action,
	{_State261, DoToken}:                          _GotoState321Action,
	{_State262, SemicolonToken}:                   _GotoState322Action,
	{_State264, LbraceToken}:                      _GotoState127Action,
	{_State264, BlockBodyType}:                    _GotoState323Action,
	{_State265, ElseToken}:                        _GotoState324Action,
	{_State266, IntegerLiteralToken}:              _GotoState25Action,
	{_State266, FloatLiteralToken}:                _GotoState22Action,
	{_State266, RuneLiteralToken}:                 _GotoState32Action,
	{_State266, StringLiteralToken}:               _GotoState33Action,
	{_State266, IdentifierToken}:                  _GotoState24Action,
	{_State266, TrueToken}:                        _GotoState36Action,
	{_State266, FalseToken}:                       _GotoState21Action,
	{_State266, StructToken}:                      _GotoState34Action,
	{_State266, FuncToken}:                        _GotoState23Action,
	{_State266, LabelDeclToken}:                   _GotoState26Action,
	{_State266, LparenToken}:                      _GotoState29Action,
	{_State266, LbracketToken}:                    _GotoState27Action,
	{_State266, LexErrorToken}:                    _GotoState28Action,
	{_State266, OptionalLabelDeclType}:            _GotoState134Action,
	{_State266, BlockExprType}:                    _GotoState42Action,
	{_State266, CallExprType}:                     _GotoState326Action,
	{_State266, AtomExprType}:                     _GotoState41Action,
	{_State266, LiteralType}:                      _GotoState48Action,
	{_State266, InitializeExprType}:               _GotoState47Action,
	{_State266, AccessExprType}:                   _GotoState325Action,
	{_State266, InitializableTypeType}:            _GotoState46Action,
	{_State266, ExplicitStructDefType}:            _GotoState45Action,
	{_State266, AnonymousFuncExprType}:            _GotoState40Action,
	{_State269, IntegerLiteralToken}:              _GotoState25Action,
	{_State269, FloatLiteralToken}:                _GotoState22Action,
	{_State269, RuneLiteralToken}:                 _GotoState32Action,
	{_State269, StringLiteralToken}:               _GotoState33Action,
	{_State269, IdentifierToken}:                  _GotoState24Action,
	{_State269, TrueToken}:                        _GotoState36Action,
	{_State269, FalseToken}:                       _GotoState21Action,
	{_State269, StructToken}:                      _GotoState34Action,
	{_State269, FuncToken}:                        _GotoState23Action,
	{_State269, LabelDeclToken}:                   _GotoState26Action,
	{_State269, LparenToken}:                      _GotoState29Action,
	{_State269, LbracketToken}:                    _GotoState27Action,
	{_State269, LexErrorToken}:                    _GotoState28Action,
	{_State269, OptionalLabelDeclType}:            _GotoState134Action,
	{_State269, BlockExprType}:                    _GotoState42Action,
	{_State269, CallExprType}:                     _GotoState327Action,
	{_State269, AtomExprType}:                     _GotoState41Action,
	{_State269, LiteralType}:                      _GotoState48Action,
	{_State269, InitializeExprType}:               _GotoState47Action,
	{_State269, AccessExprType}:                   _GotoState325Action,
	{_State269, InitializableTypeType}:            _GotoState46Action,
	{_State269, ExplicitStructDefType}:            _GotoState45Action,
	{_State269, AnonymousFuncExprType}:            _GotoState40Action,
	{_State272, LbracketToken}:                    _GotoState100Action,
	{_State272, DotToken}:                         _GotoState99Action,
	{_State272, QuestionToken}:                    _GotoState101Action,
	{_State272, DollarLbracketToken}:              _GotoState98Action,
	{_State272, AddAssignToken}:                   _GotoState328Action,
	{_State272, SubAssignToken}:                   _GotoState339Action,
	{_State272, MulAssignToken}:                   _GotoState338Action,
	{_State272, DivAssignToken}:                   _GotoState336Action,
	{_State272, ModAssignToken}:                   _GotoState337Action,
	{_State272, AddOneAssignToken}:                _GotoState329Action,
	{_State272, SubOneAssignToken}:                _GotoState340Action,
	{_State272, BitNegAssignToken}:                _GotoState332Action,
	{_State272, BitAndAssignToken}:                _GotoState330Action,
	{_State272, BitOrAssignToken}:                 _GotoState333Action,
	{_State272, BitXorAssignToken}:                _GotoState335Action,
	{_State272, BitLshiftAssignToken}:             _GotoState331Action,
	{_State272, BitRshiftAssignToken}:             _GotoState334Action,
	{_State272, UnaryOpAssignType}:                _GotoState342Action,
	{_State272, BinaryOpAssignType}:               _GotoState341Action,
	{_State272, OptionalGenericBindingType}:       _GotoState102Action,
	{_State274, CommaToken}:                       _GotoState343Action,
	{_State276, JumpLabelToken}:                   _GotoState344Action,
	{_State276, OptionalJumpLabelType}:            _GotoState345Action,
	{_State278, NewlinesToken}:                    _GotoState346Action,
	{_State278, SemicolonToken}:                   _GotoState347Action,
	{_State280, CaseToken}:                        _GotoState348Action,
	{_State280, CaseBranchesType}:                 _GotoState350Action,
	{_State280, CaseBranchType}:                   _GotoState349Action,
	{_State281, StringLiteralToken}:               _GotoState351Action,
	{_State284, NewlinesToken}:                    _GotoState352Action,
	{_State284, SemicolonToken}:                   _GotoState353Action,
	{_State286, AddToken}:                         _GotoState168Action,
	{_State286, SubToken}:                         _GotoState173Action,
	{_State286, MulToken}:                         _GotoState171Action,
	{_State287, IdentifierToken}:                  _GotoState204Action,
	{_State287, GenericParameterDefType}:          _GotoState354Action,
	{_State289, IdentifierToken}:                  _GotoState76Action,
	{_State289, StructToken}:                      _GotoState34Action,
	{_State289, EnumToken}:                        _GotoState73Action,
	{_State289, TraitToken}:                       _GotoState81Action,
	{_State289, FuncToken}:                        _GotoState75Action,
	{_State289, LparenToken}:                      _GotoState78Action,
	{_State289, LbracketToken}:                    _GotoState27Action,
	{_State289, DotToken}:                         _GotoState72Action,
	{_State289, QuestionToken}:                    _GotoState79Action,
	{_State289, ExclaimToken}:                     _GotoState74Action,
	{_State289, TildeTildeToken}:                  _GotoState80Action,
	{_State289, BitNegToken}:                      _GotoState71Action,
	{_State289, BitAndToken}:                      _GotoState70Action,
	{_State289, LexErrorToken}:                    _GotoState77Action,
	{_State289, InitializableTypeType}:            _GotoState87Action,
	{_State289, AtomTypeType}:                     _GotoState82Action,
	{_State289, ReturnableTypeType}:               _GotoState88Action,
	{_State289, ValueTypeType}:                    _GotoState355Action,
	{_State289, ImplicitStructDefType}:            _GotoState86Action,
	{_State289, ExplicitStructDefType}:            _GotoState45Action,
	{_State289, ImplicitEnumDefType}:              _GotoState85Action,
	{_State289, ExplicitEnumDefType}:              _GotoState83Action,
	{_State289, TraitDefType}:                     _GotoState89Action,
	{_State289, FuncTypeType}:                     _GotoState84Action,
	{_State290, AddToken}:                         _GotoState168Action,
	{_State290, SubToken}:                         _GotoState173Action,
	{_State290, MulToken}:                         _GotoState171Action,
	{_State291, IdentifierToken}:                  _GotoState143Action,
	{_State291, ParameterDefType}:                 _GotoState147Action,
	{_State291, ParameterDefsType}:                _GotoState148Action,
	{_State291, OptionalParameterDefsType}:        _GotoState356Action,
	{_State292, LbraceToken}:                      _GotoState127Action,
	{_State292, BlockBodyType}:                    _GotoState357Action,
	{_State295, IdentifierToken}:                  _GotoState157Action,
	{_State295, UnsafeToken}:                      _GotoState58Action,
	{_State295, StructToken}:                      _GotoState34Action,
	{_State295, EnumToken}:                        _GotoState73Action,
	{_State295, TraitToken}:                       _GotoState81Action,
	{_State295, FuncToken}:                        _GotoState75Action,
	{_State295, LparenToken}:                      _GotoState78Action,
	{_State295, LbracketToken}:                    _GotoState27Action,
	{_State295, DotToken}:                         _GotoState72Action,
	{_State295, QuestionToken}:                    _GotoState79Action,
	{_State295, ExclaimToken}:                     _GotoState74Action,
	{_State295, TildeTildeToken}:                  _GotoState80Action,
	{_State295, BitNegToken}:                      _GotoState71Action,
	{_State295, BitAndToken}:                      _GotoState70Action,
	{_State295, LexErrorToken}:                    _GotoState77Action,
	{_State295, UnsafeStatementType}:              _GotoState163Action,
	{_State295, InitializableTypeType}:            _GotoState87Action,
	{_State295, AtomTypeType}:                     _GotoState82Action,
	{_State295, ReturnableTypeType}:               _GotoState88Action,
	{_State295, ValueTypeType}:                    _GotoState164Action,
	{_State295, FieldDefType}:                     _GotoState218Action,
	{_State295, ImplicitStructDefType}:            _GotoState86Action,
	{_State295, ExplicitStructDefType}:            _GotoState45Action,
	{_State295, EnumValueDefType}:                 _GotoState358Action,
	{_State295, ImplicitEnumDefType}:              _GotoState85Action,
	{_State295, ExplicitEnumDefType}:              _GotoState83Action,
	{_State295, TraitDefType}:                     _GotoState89Action,
	{_State295, FuncTypeType}:                     _GotoState84Action,
	{_State296, IdentifierToken}:                  _GotoState157Action,
	{_State296, UnsafeToken}:                      _GotoState58Action,
	{_State296, StructToken}:                      _GotoState34Action,
	{_State296, EnumToken}:                        _GotoState73Action,
	{_State296, TraitToken}:                       _GotoState81Action,
	{_State296, FuncToken}:                        _GotoState75Action,
	{_State296, LparenToken}:                      _GotoState78Action,
	{_State296, LbracketToken}:                    _GotoState27Action,
	{_State296, DotToken}:                         _GotoState72Action,
	{_State296, QuestionToken}:                    _GotoState79Action,
	{_State296, ExclaimToken}:                     _GotoState74Action,
	{_State296, TildeTildeToken}:                  _GotoState80Action,
	{_State296, BitNegToken}:                      _GotoState71Action,
	{_State296, BitAndToken}:                      _GotoState70Action,
	{_State296, LexErrorToken}:                    _GotoState77Action,
	{_State296, UnsafeStatementType}:              _GotoState163Action,
	{_State296, InitializableTypeType}:            _GotoState87Action,
	{_State296, AtomTypeType}:                     _GotoState82Action,
	{_State296, ReturnableTypeType}:               _GotoState88Action,
	{_State296, ValueTypeType}:                    _GotoState164Action,
	{_State296, FieldDefType}:                     _GotoState218Action,
	{_State296, ImplicitStructDefType}:            _GotoState86Action,
	{_State296, ExplicitStructDefType}:            _GotoState45Action,
	{_State296, EnumValueDefType}:                 _GotoState359Action,
	{_State296, ImplicitEnumDefType}:              _GotoState85Action,
	{_State296, ExplicitEnumDefType}:              _GotoState83Action,
	{_State296, TraitDefType}:                     _GotoState89Action,
	{_State296, FuncTypeType}:                     _GotoState84Action,
	{_State298, IdentifierToken}:                  _GotoState157Action,
	{_State298, UnsafeToken}:                      _GotoState58Action,
	{_State298, StructToken}:                      _GotoState34Action,
	{_State298, EnumToken}:                        _GotoState73Action,
	{_State298, TraitToken}:                       _GotoState81Action,
	{_State298, FuncToken}:                        _GotoState75Action,
	{_State298, LparenToken}:                      _GotoState78Action,
	{_State298, LbracketToken}:                    _GotoState27Action,
	{_State298, DotToken}:                         _GotoState72Action,
	{_State298, QuestionToken}:                    _GotoState79Action,
	{_State298, ExclaimToken}:                     _GotoState74Action,
	{_State298, TildeTildeToken}:                  _GotoState80Action,
	{_State298, BitNegToken}:                      _GotoState71Action,
	{_State298, BitAndToken}:                      _GotoState70Action,
	{_State298, LexErrorToken}:                    _GotoState77Action,
	{_State298, UnsafeStatementType}:              _GotoState163Action,
	{_State298, InitializableTypeType}:            _GotoState87Action,
	{_State298, AtomTypeType}:                     _GotoState82Action,
	{_State298, ReturnableTypeType}:               _GotoState88Action,
	{_State298, ValueTypeType}:                    _GotoState164Action,
	{_State298, FieldDefType}:                     _GotoState218Action,
	{_State298, ImplicitStructDefType}:            _GotoState86Action,
	{_State298, ExplicitStructDefType}:            _GotoState45Action,
	{_State298, EnumValueDefType}:                 _GotoState360Action,
	{_State298, ImplicitEnumDefType}:              _GotoState85Action,
	{_State298, ExplicitEnumDefType}:              _GotoState83Action,
	{_State298, TraitDefType}:                     _GotoState89Action,
	{_State298, FuncTypeType}:                     _GotoState84Action,
	{_State299, IdentifierToken}:                  _GotoState157Action,
	{_State299, UnsafeToken}:                      _GotoState58Action,
	{_State299, StructToken}:                      _GotoState34Action,
	{_State299, EnumToken}:                        _GotoState73Action,
	{_State299, TraitToken}:                       _GotoState81Action,
	{_State299, FuncToken}:                        _GotoState75Action,
	{_State299, LparenToken}:                      _GotoState78Action,
	{_State299, LbracketToken}:                    _GotoState27Action,
	{_State299, DotToken}:                         _GotoState72Action,
	{_State299, QuestionToken}:                    _GotoState79Action,
	{_State299, ExclaimToken}:                     _GotoState74Action,
	{_State299, TildeTildeToken}:                  _GotoState80Action,
	{_State299, BitNegToken}:                      _GotoState71Action,
	{_State299, BitAndToken}:                      _GotoState70Action,
	{_State299, LexErrorToken}:                    _GotoState77Action,
	{_State299, UnsafeStatementType}:              _GotoState163Action,
	{_State299, InitializableTypeType}:            _GotoState87Action,
	{_State299, AtomTypeType}:                     _GotoState82Action,
	{_State299, ReturnableTypeType}:               _GotoState88Action,
	{_State299, ValueTypeType}:                    _GotoState164Action,
	{_State299, FieldDefType}:                     _GotoState218Action,
	{_State299, ImplicitStructDefType}:            _GotoState86Action,
	{_State299, ExplicitStructDefType}:            _GotoState45Action,
	{_State299, EnumValueDefType}:                 _GotoState361Action,
	{_State299, ImplicitEnumDefType}:              _GotoState85Action,
	{_State299, ExplicitEnumDefType}:              _GotoState83Action,
	{_State299, TraitDefType}:                     _GotoState89Action,
	{_State299, FuncTypeType}:                     _GotoState84Action,
	{_State300, AddToken}:                         _GotoState168Action,
	{_State300, SubToken}:                         _GotoState173Action,
	{_State300, MulToken}:                         _GotoState171Action,
	{_State301, IdentifierToken}:                  _GotoState76Action,
	{_State301, StructToken}:                      _GotoState34Action,
	{_State301, EnumToken}:                        _GotoState73Action,
	{_State301, TraitToken}:                       _GotoState81Action,
	{_State301, FuncToken}:                        _GotoState75Action,
	{_State301, LparenToken}:                      _GotoState78Action,
	{_State301, LbracketToken}:                    _GotoState27Action,
	{_State301, DotToken}:                         _GotoState72Action,
	{_State301, QuestionToken}:                    _GotoState79Action,
	{_State301, ExclaimToken}:                     _GotoState74Action,
	{_State301, TildeTildeToken}:                  _GotoState80Action,
	{_State301, BitNegToken}:                      _GotoState71Action,
	{_State301, BitAndToken}:                      _GotoState70Action,
	{_State301, LexErrorToken}:                    _GotoState77Action,
	{_State301, InitializableTypeType}:            _GotoState87Action,
	{_State301, AtomTypeType}:                     _GotoState82Action,
	{_State301, ReturnableTypeType}:               _GotoState88Action,
	{_State301, ValueTypeType}:                    _GotoState362Action,
	{_State301, ImplicitStructDefType}:            _GotoState86Action,
	{_State301, ExplicitStructDefType}:            _GotoState45Action,
	{_State301, ImplicitEnumDefType}:              _GotoState85Action,
	{_State301, ExplicitEnumDefType}:              _GotoState83Action,
	{_State301, TraitDefType}:                     _GotoState89Action,
	{_State301, FuncTypeType}:                     _GotoState84Action,
	{_State302, AddToken}:                         _GotoState168Action,
	{_State302, SubToken}:                         _GotoState173Action,
	{_State302, MulToken}:                         _GotoState171Action,
	{_State303, IdentifierToken}:                  _GotoState76Action,
	{_State303, StructToken}:                      _GotoState34Action,
	{_State303, EnumToken}:                        _GotoState73Action,
	{_State303, TraitToken}:                       _GotoState81Action,
	{_State303, FuncToken}:                        _GotoState75Action,
	{_State303, LparenToken}:                      _GotoState78Action,
	{_State303, LbracketToken}:                    _GotoState27Action,
	{_State303, DotToken}:                         _GotoState72Action,
	{_State303, QuestionToken}:                    _GotoState79Action,
	{_State303, ExclaimToken}:                     _GotoState74Action,
	{_State303, TildeTildeToken}:                  _GotoState80Action,
	{_State303, BitNegToken}:                      _GotoState71Action,
	{_State303, BitAndToken}:                      _GotoState70Action,
	{_State303, LexErrorToken}:                    _GotoState77Action,
	{_State303, InitializableTypeType}:            _GotoState87Action,
	{_State303, AtomTypeType}:                     _GotoState82Action,
	{_State303, ReturnableTypeType}:               _GotoState293Action,
	{_State303, ImplicitStructDefType}:            _GotoState86Action,
	{_State303, ExplicitStructDefType}:            _GotoState45Action,
	{_State303, ImplicitEnumDefType}:              _GotoState85Action,
	{_State303, ExplicitEnumDefType}:              _GotoState83Action,
	{_State303, TraitDefType}:                     _GotoState89Action,
	{_State303, ReturnTypeType}:                   _GotoState363Action,
	{_State303, FuncTypeType}:                     _GotoState84Action,
	{_State304, IdentifierToken}:                  _GotoState221Action,
	{_State304, StructToken}:                      _GotoState34Action,
	{_State304, EnumToken}:                        _GotoState73Action,
	{_State304, TraitToken}:                       _GotoState81Action,
	{_State304, FuncToken}:                        _GotoState75Action,
	{_State304, LparenToken}:                      _GotoState78Action,
	{_State304, LbracketToken}:                    _GotoState27Action,
	{_State304, DotToken}:                         _GotoState72Action,
	{_State304, QuestionToken}:                    _GotoState79Action,
	{_State304, ExclaimToken}:                     _GotoState74Action,
	{_State304, DotdotdotToken}:                   _GotoState220Action,
	{_State304, TildeTildeToken}:                  _GotoState80Action,
	{_State304, BitNegToken}:                      _GotoState71Action,
	{_State304, BitAndToken}:                      _GotoState70Action,
	{_State304, LexErrorToken}:                    _GotoState77Action,
	{_State304, InitializableTypeType}:            _GotoState87Action,
	{_State304, AtomTypeType}:                     _GotoState82Action,
	{_State304, ReturnableTypeType}:               _GotoState88Action,
	{_State304, ValueTypeType}:                    _GotoState225Action,
	{_State304, ImplicitStructDefType}:            _GotoState86Action,
	{_State304, ExplicitStructDefType}:            _GotoState45Action,
	{_State304, ImplicitEnumDefType}:              _GotoState85Action,
	{_State304, ExplicitEnumDefType}:              _GotoState83Action,
	{_State304, TraitDefType}:                     _GotoState89Action,
	{_State304, ParameterDeclType}:                _GotoState364Action,
	{_State304, FuncTypeType}:                     _GotoState84Action,
	{_State310, LparenToken}:                      _GotoState365Action,
	{_State312, IdentifierToken}:                  _GotoState157Action,
	{_State312, UnsafeToken}:                      _GotoState58Action,
	{_State312, StructToken}:                      _GotoState34Action,
	{_State312, EnumToken}:                        _GotoState73Action,
	{_State312, TraitToken}:                       _GotoState81Action,
	{_State312, FuncToken}:                        _GotoState235Action,
	{_State312, LparenToken}:                      _GotoState78Action,
	{_State312, LbracketToken}:                    _GotoState27Action,
	{_State312, DotToken}:                         _GotoState72Action,
	{_State312, QuestionToken}:                    _GotoState79Action,
	{_State312, ExclaimToken}:                     _GotoState74Action,
	{_State312, TildeTildeToken}:                  _GotoState80Action,
	{_State312, BitNegToken}:                      _GotoState71Action,
	{_State312, BitAndToken}:                      _GotoState70Action,
	{_State312, LexErrorToken}:                    _GotoState77Action,
	{_State312, UnsafeStatementType}:              _GotoState163Action,
	{_State312, InitializableTypeType}:            _GotoState87Action,
	{_State312, AtomTypeType}:                     _GotoState82Action,
	{_State312, ReturnableTypeType}:               _GotoState88Action,
	{_State312, ValueTypeType}:                    _GotoState164Action,
	{_State312, FieldDefType}:                     _GotoState236Action,
	{_State312, ImplicitStructDefType}:            _GotoState86Action,
	{_State312, ExplicitStructDefType}:            _GotoState45Action,
	{_State312, ImplicitEnumDefType}:              _GotoState85Action,
	{_State312, ExplicitEnumDefType}:              _GotoState83Action,
	{_State312, TraitPropertyType}:                _GotoState366Action,
	{_State312, TraitDefType}:                     _GotoState89Action,
	{_State312, FuncTypeType}:                     _GotoState84Action,
	{_State312, MethodSignatureType}:              _GotoState237Action,
	{_State313, IdentifierToken}:                  _GotoState157Action,
	{_State313, UnsafeToken}:                      _GotoState58Action,
	{_State313, StructToken}:                      _GotoState34Action,
	{_State313, EnumToken}:                        _GotoState73Action,
	{_State313, TraitToken}:                       _GotoState81Action,
	{_State313, FuncToken}:                        _GotoState235Action,
	{_State313, LparenToken}:                      _GotoState78Action,
	{_State313, LbracketToken}:                    _GotoState27Action,
	{_State313, DotToken}:                         _GotoState72Action,
	{_State313, QuestionToken}:                    _GotoState79Action,
	{_State313, ExclaimToken}:                     _GotoState74Action,
	{_State313, TildeTildeToken}:                  _GotoState80Action,
	{_State313, BitNegToken}:                      _GotoState71Action,
	{_State313, BitAndToken}:                      _GotoState70Action,
	{_State313, LexErrorToken}:                    _GotoState77Action,
	{_State313, UnsafeStatementType}:              _GotoState163Action,
	{_State313, InitializableTypeType}:            _GotoState87Action,
	{_State313, AtomTypeType}:                     _GotoState82Action,
	{_State313, ReturnableTypeType}:               _GotoState88Action,
	{_State313, ValueTypeType}:                    _GotoState164Action,
	{_State313, FieldDefType}:                     _GotoState236Action,
	{_State313, ImplicitStructDefType}:            _GotoState86Action,
	{_State313, ExplicitStructDefType}:            _GotoState45Action,
	{_State313, ImplicitEnumDefType}:              _GotoState85Action,
	{_State313, ExplicitEnumDefType}:              _GotoState83Action,
	{_State313, TraitPropertyType}:                _GotoState367Action,
	{_State313, TraitDefType}:                     _GotoState89Action,
	{_State313, FuncTypeType}:                     _GotoState84Action,
	{_State313, MethodSignatureType}:              _GotoState237Action,
	{_State318, AddToken}:                         _GotoState168Action,
	{_State318, SubToken}:                         _GotoState173Action,
	{_State318, MulToken}:                         _GotoState171Action,
	{_State321, LbraceToken}:                      _GotoState127Action,
	{_State321, BlockBodyType}:                    _GotoState368Action,
	{_State322, IntegerLiteralToken}:              _GotoState25Action,
	{_State322, FloatLiteralToken}:                _GotoState22Action,
	{_State322, RuneLiteralToken}:                 _GotoState32Action,
	{_State322, StringLiteralToken}:               _GotoState33Action,
	{_State322, IdentifierToken}:                  _GotoState24Action,
	{_State322, TrueToken}:                        _GotoState36Action,
	{_State322, FalseToken}:                       _GotoState21Action,
	{_State322, StructToken}:                      _GotoState34Action,
	{_State322, FuncToken}:                        _GotoState23Action,
	{_State322, LabelDeclToken}:                   _GotoState26Action,
	{_State322, LparenToken}:                      _GotoState29Action,
	{_State322, LbracketToken}:                    _GotoState27Action,
	{_State322, NotToken}:                         _GotoState31Action,
	{_State322, SubToken}:                         _GotoState35Action,
	{_State322, MulToken}:                         _GotoState30Action,
	{_State322, BitNegToken}:                      _GotoState20Action,
	{_State322, BitAndToken}:                      _GotoState19Action,
	{_State322, LexErrorToken}:                    _GotoState28Action,
	{_State322, OptionalLabelDeclType}:            _GotoState134Action,
	{_State322, OptionalSequenceExprType}:         _GotoState369Action,
	{_State322, SequenceExprType}:                 _GotoState263Action,
	{_State322, BlockExprType}:                    _GotoState42Action,
	{_State322, CallExprType}:                     _GotoState43Action,
	{_State322, AtomExprType}:                     _GotoState41Action,
	{_State322, LiteralType}:                      _GotoState48Action,
	{_State322, InitializeExprType}:               _GotoState47Action,
	{_State322, AccessExprType}:                   _GotoState37Action,
	{_State322, PostfixUnaryExprType}:             _GotoState52Action,
	{_State322, PrefixUnaryOpType}:                _GotoState54Action,
	{_State322, PrefixUnaryExprType}:              _GotoState53Action,
	{_State322, MulExprType}:                      _GotoState49Action,
	{_State322, AddExprType}:                      _GotoState38Action,
	{_State322, CmpExprType}:                      _GotoState44Action,
	{_State322, AndExprType}:                      _GotoState39Action,
	{_State322, OrExprType}:                       _GotoState51Action,
	{_State322, InitializableTypeType}:            _GotoState46Action,
	{_State322, ExplicitStructDefType}:            _GotoState45Action,
	{_State322, AnonymousFuncExprType}:            _GotoState40Action,
	{_State324, IfToken}:                          _GotoState126Action,
	{_State324, LbraceToken}:                      _GotoState127Action,
	{_State324, IfExprType}:                       _GotoState371Action,
	{_State324, BlockBodyType}:                    _GotoState370Action,
	{_State325, LbracketToken}:                    _GotoState100Action,
	{_State325, DotToken}:                         _GotoState99Action,
	{_State325, DollarLbracketToken}:              _GotoState98Action,
	{_State325, OptionalGenericBindingType}:       _GotoState102Action,
	{_State341, IntegerLiteralToken}:              _GotoState25Action,
	{_State341, FloatLiteralToken}:                _GotoState22Action,
	{_State341, RuneLiteralToken}:                 _GotoState32Action,
	{_State341, StringLiteralToken}:               _GotoState33Action,
	{_State341, IdentifierToken}:                  _GotoState24Action,
	{_State341, TrueToken}:                        _GotoState36Action,
	{_State341, FalseToken}:                       _GotoState21Action,
	{_State341, StructToken}:                      _GotoState34Action,
	{_State341, FuncToken}:                        _GotoState23Action,
	{_State341, LabelDeclToken}:                   _GotoState26Action,
	{_State341, LparenToken}:                      _GotoState29Action,
	{_State341, LbracketToken}:                    _GotoState27Action,
	{_State341, NotToken}:                         _GotoState31Action,
	{_State341, SubToken}:                         _GotoState35Action,
	{_State341, MulToken}:                         _GotoState30Action,
	{_State341, BitNegToken}:                      _GotoState20Action,
	{_State341, BitAndToken}:                      _GotoState19Action,
	{_State341, LexErrorToken}:                    _GotoState28Action,
	{_State341, ExpressionType}:                   _GotoState372Action,
	{_State341, OptionalLabelDeclType}:            _GotoState50Action,
	{_State341, SequenceExprType}:                 _GotoState55Action,
	{_State341, BlockExprType}:                    _GotoState42Action,
	{_State341, CallExprType}:                     _GotoState43Action,
	{_State341, AtomExprType}:                     _GotoState41Action,
	{_State341, LiteralType}:                      _GotoState48Action,
	{_State341, InitializeExprType}:               _GotoState47Action,
	{_State341, AccessExprType}:                   _GotoState37Action,
	{_State341, PostfixUnaryExprType}:             _GotoState52Action,
	{_State341, PrefixUnaryOpType}:                _GotoState54Action,
	{_State341, PrefixUnaryExprType}:              _GotoState53Action,
	{_State341, MulExprType}:                      _GotoState49Action,
	{_State341, AddExprType}:                      _GotoState38Action,
	{_State341, CmpExprType}:                      _GotoState44Action,
	{_State341, AndExprType}:                      _GotoState39Action,
	{_State341, OrExprType}:                       _GotoState51Action,
	{_State341, InitializableTypeType}:            _GotoState46Action,
	{_State341, ExplicitStructDefType}:            _GotoState45Action,
	{_State341, AnonymousFuncExprType}:            _GotoState40Action,
	{_State343, IntegerLiteralToken}:              _GotoState25Action,
	{_State343, FloatLiteralToken}:                _GotoState22Action,
	{_State343, RuneLiteralToken}:                 _GotoState32Action,
	{_State343, StringLiteralToken}:               _GotoState33Action,
	{_State343, IdentifierToken}:                  _GotoState24Action,
	{_State343, TrueToken}:                        _GotoState36Action,
	{_State343, FalseToken}:                       _GotoState21Action,
	{_State343, StructToken}:                      _GotoState34Action,
	{_State343, FuncToken}:                        _GotoState23Action,
	{_State343, LabelDeclToken}:                   _GotoState26Action,
	{_State343, LparenToken}:                      _GotoState29Action,
	{_State343, LbracketToken}:                    _GotoState27Action,
	{_State343, NotToken}:                         _GotoState31Action,
	{_State343, SubToken}:                         _GotoState35Action,
	{_State343, MulToken}:                         _GotoState30Action,
	{_State343, BitNegToken}:                      _GotoState20Action,
	{_State343, BitAndToken}:                      _GotoState19Action,
	{_State343, LexErrorToken}:                    _GotoState28Action,
	{_State343, ExpressionType}:                   _GotoState373Action,
	{_State343, OptionalLabelDeclType}:            _GotoState50Action,
	{_State343, SequenceExprType}:                 _GotoState55Action,
	{_State343, BlockExprType}:                    _GotoState42Action,
	{_State343, CallExprType}:                     _GotoState43Action,
	{_State343, AtomExprType}:                     _GotoState41Action,
	{_State343, LiteralType}:                      _GotoState48Action,
	{_State343, InitializeExprType}:               _GotoState47Action,
	{_State343, AccessExprType}:                   _GotoState37Action,
	{_State343, PostfixUnaryExprType}:             _GotoState52Action,
	{_State343, PrefixUnaryOpType}:                _GotoState54Action,
	{_State343, PrefixUnaryExprType}:              _GotoState53Action,
	{_State343, MulExprType}:                      _GotoState49Action,
	{_State343, AddExprType}:                      _GotoState38Action,
	{_State343, CmpExprType}:                      _GotoState44Action,
	{_State343, AndExprType}:                      _GotoState39Action,
	{_State343, OrExprType}:                       _GotoState51Action,
	{_State343, InitializableTypeType}:            _GotoState46Action,
	{_State343, ExplicitStructDefType}:            _GotoState45Action,
	{_State343, AnonymousFuncExprType}:            _GotoState40Action,
	{_State345, IntegerLiteralToken}:              _GotoState25Action,
	{_State345, FloatLiteralToken}:                _GotoState22Action,
	{_State345, RuneLiteralToken}:                 _GotoState32Action,
	{_State345, StringLiteralToken}:               _GotoState33Action,
	{_State345, IdentifierToken}:                  _GotoState24Action,
	{_State345, TrueToken}:                        _GotoState36Action,
	{_State345, FalseToken}:                       _GotoState21Action,
	{_State345, StructToken}:                      _GotoState34Action,
	{_State345, FuncToken}:                        _GotoState23Action,
	{_State345, LabelDeclToken}:                   _GotoState26Action,
	{_State345, LparenToken}:                      _GotoState29Action,
	{_State345, LbracketToken}:                    _GotoState27Action,
	{_State345, NotToken}:                         _GotoState31Action,
	{_State345, SubToken}:                         _GotoState35Action,
	{_State345, MulToken}:                         _GotoState30Action,
	{_State345, BitNegToken}:                      _GotoState20Action,
	{_State345, BitAndToken}:                      _GotoState19Action,
	{_State345, LexErrorToken}:                    _GotoState28Action,
	{_State345, ExpressionType}:                   _GotoState273Action,
	{_State345, OptionalLabelDeclType}:            _GotoState50Action,
	{_State345, SequenceExprType}:                 _GotoState55Action,
	{_State345, BlockExprType}:                    _GotoState42Action,
	{_State345, ExpressionsType}:                  _GotoState374Action,
	{_State345, OptionalExpressionsType}:          _GotoState375Action,
	{_State345, CallExprType}:                     _GotoState43Action,
	{_State345, AtomExprType}:                     _GotoState41Action,
	{_State345, LiteralType}:                      _GotoState48Action,
	{_State345, InitializeExprType}:               _GotoState47Action,
	{_State345, AccessExprType}:                   _GotoState37Action,
	{_State345, PostfixUnaryExprType}:             _GotoState52Action,
	{_State345, PrefixUnaryOpType}:                _GotoState54Action,
	{_State345, PrefixUnaryExprType}:              _GotoState53Action,
	{_State345, MulExprType}:                      _GotoState49Action,
	{_State345, AddExprType}:                      _GotoState38Action,
	{_State345, CmpExprType}:                      _GotoState44Action,
	{_State345, AndExprType}:                      _GotoState39Action,
	{_State345, OrExprType}:                       _GotoState51Action,
	{_State345, InitializableTypeType}:            _GotoState46Action,
	{_State345, ExplicitStructDefType}:            _GotoState45Action,
	{_State345, AnonymousFuncExprType}:            _GotoState40Action,
	{_State348, MatchPatternType}:                 _GotoState376Action,
	{_State350, CaseToken}:                        _GotoState348Action,
	{_State350, DefaultToken}:                     _GotoState377Action,
	{_State350, CaseBranchType}:                   _GotoState378Action,
	{_State350, OptionalDefaultBranchType}:        _GotoState380Action,
	{_State350, DefaultBranchType}:                _GotoState379Action,
	{_State355, AddToken}:                         _GotoState168Action,
	{_State355, SubToken}:                         _GotoState173Action,
	{_State355, MulToken}:                         _GotoState171Action,
	{_State356, RparenToken}:                      _GotoState381Action,
	{_State362, AddToken}:                         _GotoState168Action,
	{_State362, SubToken}:                         _GotoState173Action,
	{_State362, MulToken}:                         _GotoState171Action,
	{_State365, IdentifierToken}:                  _GotoState221Action,
	{_State365, StructToken}:                      _GotoState34Action,
	{_State365, EnumToken}:                        _GotoState73Action,
	{_State365, TraitToken}:                       _GotoState81Action,
	{_State365, FuncToken}:                        _GotoState75Action,
	{_State365, LparenToken}:                      _GotoState78Action,
	{_State365, LbracketToken}:                    _GotoState27Action,
	{_State365, DotToken}:                         _GotoState72Action,
	{_State365, QuestionToken}:                    _GotoState79Action,
	{_State365, ExclaimToken}:                     _GotoState74Action,
	{_State365, DotdotdotToken}:                   _GotoState220Action,
	{_State365, TildeTildeToken}:                  _GotoState80Action,
	{_State365, BitNegToken}:                      _GotoState71Action,
	{_State365, BitAndToken}:                      _GotoState70Action,
	{_State365, LexErrorToken}:                    _GotoState77Action,
	{_State365, InitializableTypeType}:            _GotoState87Action,
	{_State365, AtomTypeType}:                     _GotoState82Action,
	{_State365, ReturnableTypeType}:               _GotoState88Action,
	{_State365, ValueTypeType}:                    _GotoState225Action,
	{_State365, ImplicitStructDefType}:            _GotoState86Action,
	{_State365, ExplicitStructDefType}:            _GotoState45Action,
	{_State365, ImplicitEnumDefType}:              _GotoState85Action,
	{_State365, ExplicitEnumDefType}:              _GotoState83Action,
	{_State365, TraitDefType}:                     _GotoState89Action,
	{_State365, ParameterDeclType}:                _GotoState223Action,
	{_State365, ParameterDeclsType}:               _GotoState224Action,
	{_State365, OptionalParameterDeclsType}:       _GotoState382Action,
	{_State365, FuncTypeType}:                     _GotoState84Action,
	{_State369, DoToken}:                          _GotoState383Action,
	{_State374, CommaToken}:                       _GotoState343Action,
	{_State376, ColonToken}:                       _GotoState384Action,
	{_State377, ColonToken}:                       _GotoState385Action,
	{_State380, RbraceToken}:                      _GotoState386Action,
	{_State381, IdentifierToken}:                  _GotoState76Action,
	{_State381, StructToken}:                      _GotoState34Action,
	{_State381, EnumToken}:                        _GotoState73Action,
	{_State381, TraitToken}:                       _GotoState81Action,
	{_State381, FuncToken}:                        _GotoState75Action,
	{_State381, LparenToken}:                      _GotoState78Action,
	{_State381, LbracketToken}:                    _GotoState27Action,
	{_State381, DotToken}:                         _GotoState72Action,
	{_State381, QuestionToken}:                    _GotoState79Action,
	{_State381, ExclaimToken}:                     _GotoState74Action,
	{_State381, TildeTildeToken}:                  _GotoState80Action,
	{_State381, BitNegToken}:                      _GotoState71Action,
	{_State381, BitAndToken}:                      _GotoState70Action,
	{_State381, LexErrorToken}:                    _GotoState77Action,
	{_State381, InitializableTypeType}:            _GotoState87Action,
	{_State381, AtomTypeType}:                     _GotoState82Action,
	{_State381, ReturnableTypeType}:               _GotoState293Action,
	{_State381, ImplicitStructDefType}:            _GotoState86Action,
	{_State381, ExplicitStructDefType}:            _GotoState45Action,
	{_State381, ImplicitEnumDefType}:              _GotoState85Action,
	{_State381, ExplicitEnumDefType}:              _GotoState83Action,
	{_State381, TraitDefType}:                     _GotoState89Action,
	{_State381, ReturnTypeType}:                   _GotoState387Action,
	{_State381, FuncTypeType}:                     _GotoState84Action,
	{_State382, RparenToken}:                      _GotoState388Action,
	{_State383, LbraceToken}:                      _GotoState127Action,
	{_State383, BlockBodyType}:                    _GotoState389Action,
	{_State384, IntegerLiteralToken}:              _GotoState25Action,
	{_State384, FloatLiteralToken}:                _GotoState22Action,
	{_State384, RuneLiteralToken}:                 _GotoState32Action,
	{_State384, StringLiteralToken}:               _GotoState33Action,
	{_State384, IdentifierToken}:                  _GotoState24Action,
	{_State384, TrueToken}:                        _GotoState36Action,
	{_State384, FalseToken}:                       _GotoState21Action,
	{_State384, StructToken}:                      _GotoState34Action,
	{_State384, FuncToken}:                        _GotoState23Action,
	{_State384, LabelDeclToken}:                   _GotoState26Action,
	{_State384, LparenToken}:                      _GotoState29Action,
	{_State384, LbracketToken}:                    _GotoState27Action,
	{_State384, NotToken}:                         _GotoState31Action,
	{_State384, SubToken}:                         _GotoState35Action,
	{_State384, MulToken}:                         _GotoState30Action,
	{_State384, BitNegToken}:                      _GotoState20Action,
	{_State384, BitAndToken}:                      _GotoState19Action,
	{_State384, LexErrorToken}:                    _GotoState28Action,
	{_State384, OptionalLabelDeclType}:            _GotoState134Action,
	{_State384, SequenceExprType}:                 _GotoState390Action,
	{_State384, BlockExprType}:                    _GotoState42Action,
	{_State384, CallExprType}:                     _GotoState43Action,
	{_State384, AtomExprType}:                     _GotoState41Action,
	{_State384, LiteralType}:                      _GotoState48Action,
	{_State384, InitializeExprType}:               _GotoState47Action,
	{_State384, AccessExprType}:                   _GotoState37Action,
	{_State384, PostfixUnaryExprType}:             _GotoState52Action,
	{_State384, PrefixUnaryOpType}:                _GotoState54Action,
	{_State384, PrefixUnaryExprType}:              _GotoState53Action,
	{_State384, MulExprType}:                      _GotoState49Action,
	{_State384, AddExprType}:                      _GotoState38Action,
	{_State384, CmpExprType}:                      _GotoState44Action,
	{_State384, AndExprType}:                      _GotoState39Action,
	{_State384, OrExprType}:                       _GotoState51Action,
	{_State384, InitializableTypeType}:            _GotoState46Action,
	{_State384, ExplicitStructDefType}:            _GotoState45Action,
	{_State384, AnonymousFuncExprType}:            _GotoState40Action,
	{_State385, IntegerLiteralToken}:              _GotoState25Action,
	{_State385, FloatLiteralToken}:                _GotoState22Action,
	{_State385, RuneLiteralToken}:                 _GotoState32Action,
	{_State385, StringLiteralToken}:               _GotoState33Action,
	{_State385, IdentifierToken}:                  _GotoState24Action,
	{_State385, TrueToken}:                        _GotoState36Action,
	{_State385, FalseToken}:                       _GotoState21Action,
	{_State385, StructToken}:                      _GotoState34Action,
	{_State385, FuncToken}:                        _GotoState23Action,
	{_State385, LabelDeclToken}:                   _GotoState26Action,
	{_State385, LparenToken}:                      _GotoState29Action,
	{_State385, LbracketToken}:                    _GotoState27Action,
	{_State385, NotToken}:                         _GotoState31Action,
	{_State385, SubToken}:                         _GotoState35Action,
	{_State385, MulToken}:                         _GotoState30Action,
	{_State385, BitNegToken}:                      _GotoState20Action,
	{_State385, BitAndToken}:                      _GotoState19Action,
	{_State385, LexErrorToken}:                    _GotoState28Action,
	{_State385, OptionalLabelDeclType}:            _GotoState134Action,
	{_State385, SequenceExprType}:                 _GotoState391Action,
	{_State385, BlockExprType}:                    _GotoState42Action,
	{_State385, CallExprType}:                     _GotoState43Action,
	{_State385, AtomExprType}:                     _GotoState41Action,
	{_State385, LiteralType}:                      _GotoState48Action,
	{_State385, InitializeExprType}:               _GotoState47Action,
	{_State385, AccessExprType}:                   _GotoState37Action,
	{_State385, PostfixUnaryExprType}:             _GotoState52Action,
	{_State385, PrefixUnaryOpType}:                _GotoState54Action,
	{_State385, PrefixUnaryExprType}:              _GotoState53Action,
	{_State385, MulExprType}:                      _GotoState49Action,
	{_State385, AddExprType}:                      _GotoState38Action,
	{_State385, CmpExprType}:                      _GotoState44Action,
	{_State385, AndExprType}:                      _GotoState39Action,
	{_State385, OrExprType}:                       _GotoState51Action,
	{_State385, InitializableTypeType}:            _GotoState46Action,
	{_State385, ExplicitStructDefType}:            _GotoState45Action,
	{_State385, AnonymousFuncExprType}:            _GotoState40Action,
	{_State387, LbraceToken}:                      _GotoState127Action,
	{_State387, BlockBodyType}:                    _GotoState392Action,
	{_State388, IdentifierToken}:                  _GotoState76Action,
	{_State388, StructToken}:                      _GotoState34Action,
	{_State388, EnumToken}:                        _GotoState73Action,
	{_State388, TraitToken}:                       _GotoState81Action,
	{_State388, FuncToken}:                        _GotoState75Action,
	{_State388, LparenToken}:                      _GotoState78Action,
	{_State388, LbracketToken}:                    _GotoState27Action,
	{_State388, DotToken}:                         _GotoState72Action,
	{_State388, QuestionToken}:                    _GotoState79Action,
	{_State388, ExclaimToken}:                     _GotoState74Action,
	{_State388, TildeTildeToken}:                  _GotoState80Action,
	{_State388, BitNegToken}:                      _GotoState71Action,
	{_State388, BitAndToken}:                      _GotoState70Action,
	{_State388, LexErrorToken}:                    _GotoState77Action,
	{_State388, InitializableTypeType}:            _GotoState87Action,
	{_State388, AtomTypeType}:                     _GotoState82Action,
	{_State388, ReturnableTypeType}:               _GotoState293Action,
	{_State388, ImplicitStructDefType}:            _GotoState86Action,
	{_State388, ExplicitStructDefType}:            _GotoState45Action,
	{_State388, ImplicitEnumDefType}:              _GotoState85Action,
	{_State388, ExplicitEnumDefType}:              _GotoState83Action,
	{_State388, TraitDefType}:                     _GotoState89Action,
	{_State388, ReturnTypeType}:                   _GotoState393Action,
	{_State388, FuncTypeType}:                     _GotoState84Action,
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
	{_State28, _WildcardMarker}:                   _ReduceLexErrorToAtomExprAction,
	{_State29, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State29, ColonToken}:                        _ReduceNilToOptionalExpressionAction,
	{_State30, _WildcardMarker}:                   _ReduceMulToPrefixUnaryOpAction,
	{_State31, _WildcardMarker}:                   _ReduceNotToPrefixUnaryOpAction,
	{_State32, _WildcardMarker}:                   _ReduceRuneLiteralToLiteralAction,
	{_State33, _WildcardMarker}:                   _ReduceStringLiteralToLiteralAction,
	{_State35, _WildcardMarker}:                   _ReduceSubToPrefixUnaryOpAction,
	{_State36, _WildcardMarker}:                   _ReduceTrueToLiteralAction,
	{_State37, _WildcardMarker}:                   _ReduceAccessExprToPostfixUnaryExprAction,
	{_State37, LparenToken}:                       _ReduceNilToOptionalGenericBindingAction,
	{_State38, _WildcardMarker}:                   _ReduceAddExprToCmpExprAction,
	{_State39, _WildcardMarker}:                   _ReduceAndExprToOrExprAction,
	{_State40, _WildcardMarker}:                   _ReduceAnonymousFuncExprToAtomExprAction,
	{_State41, _WildcardMarker}:                   _ReduceAtomExprToAccessExprAction,
	{_State42, _WildcardMarker}:                   _ReduceBlockExprToAtomExprAction,
	{_State43, _WildcardMarker}:                   _ReduceCallExprToAccessExprAction,
	{_State44, _WildcardMarker}:                   _ReduceCmpExprToAndExprAction,
	{_State45, _WildcardMarker}:                   _ReduceExplicitStructDefToInitializableTypeAction,
	{_State47, _WildcardMarker}:                   _ReduceInitializeExprToAtomExprAction,
	{_State48, _WildcardMarker}:                   _ReduceLiteralToAtomExprAction,
	{_State49, _WildcardMarker}:                   _ReduceMulExprToAddExprAction,
	{_State51, _WildcardMarker}:                   _ReduceToSequenceExprAction,
	{_State52, _WildcardMarker}:                   _ReducePostfixUnaryExprToPrefixUnaryExprAction,
	{_State53, _WildcardMarker}:                   _ReducePrefixUnaryExprToMulExprAction,
	{_State54, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State55, _EndMarker}:                        _ReduceSequenceExprToExpressionAction,
	{_State56, _EndMarker}:                        _ReduceCommentToLexInternalTokensAction,
	{_State57, _EndMarker}:                        _ReduceSpacesToLexInternalTokensAction,
	{_State59, _WildcardMarker}:                   _ReduceNilToDefinitionsAction,
	{_State60, _EndMarker}:                        _ReduceNilToOptionalNewlinesAction,
	{_State61, _WildcardMarker}:                   _ReduceNamedFuncDefToDefinitionAction,
	{_State62, _WildcardMarker}:                   _ReducePackageDefToDefinitionAction,
	{_State63, _WildcardMarker}:                   _ReduceTypeDefToDefinitionAction,
	{_State64, _WildcardMarker}:                   _ReduceUnsafeStatementToDefinitionAction,
	{_State65, _WildcardMarker}:                   _ReduceNoSpecToPackageDefAction,
	{_State66, _WildcardMarker}:                   _ReduceNilToOptionalGenericParametersAction,
	{_State69, RparenToken}:                       _ReduceNilToOptionalParameterDefsAction,
	{_State72, _WildcardMarker}:                   _ReduceNilToOptionalGenericBindingAction,
	{_State76, _WildcardMarker}:                   _ReduceNilToOptionalGenericBindingAction,
	{_State77, _WildcardMarker}:                   _ReduceLexErrorToAtomTypeAction,
	{_State78, RparenToken}:                       _ReduceNilToOptionalImplicitFieldDefsAction,
	{_State82, _WildcardMarker}:                   _ReduceAtomTypeToReturnableTypeAction,
	{_State83, _WildcardMarker}:                   _ReduceExplicitEnumDefToAtomTypeAction,
	{_State84, _WildcardMarker}:                   _ReduceFuncTypeToAtomTypeAction,
	{_State85, _WildcardMarker}:                   _ReduceImplicitEnumDefToAtomTypeAction,
	{_State86, _WildcardMarker}:                   _ReduceImplicitStructDefToAtomTypeAction,
	{_State87, _WildcardMarker}:                   _ReduceInitializableTypeToAtomTypeAction,
	{_State88, _WildcardMarker}:                   _ReduceReturnableTypeToValueTypeAction,
	{_State89, _WildcardMarker}:                   _ReduceTraitDefToAtomTypeAction,
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
	{_State147, _WildcardMarker}:                  _ReduceParameterDefToParameterDefsAction,
	{_State148, RparenToken}:                      _ReduceParameterDefsToOptionalParameterDefsAction,
	{_State149, _WildcardMarker}:                  _ReduceReferenceToReturnableTypeAction,
	{_State150, _WildcardMarker}:                  _ReducePublicMethodsTraitToReturnableTypeAction,
	{_State151, _WildcardMarker}:                  _ReduceInferredToAtomTypeAction,
	{_State153, _WildcardMarker}:                  _ReduceResultToReturnableTypeAction,
	{_State154, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State156, _WildcardMarker}:                  _ReduceNamedToAtomTypeAction,
	{_State157, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State159, _WildcardMarker}:                  _ReduceFieldDefToImplicitFieldDefsAction,
	{_State159, OrToken}:                          _ReduceFieldDefToEnumValueDefAction,
	{_State161, RparenToken}:                      _ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction,
	{_State163, _WildcardMarker}:                  _ReduceUnsafeStatementToFieldDefAction,
	{_State164, _WildcardMarker}:                  _ReduceImplicitToFieldDefAction,
	{_State165, _WildcardMarker}:                  _ReduceOptionalToReturnableTypeAction,
	{_State166, _WildcardMarker}:                  _ReducePublicTraitToReturnableTypeAction,
	{_State167, RparenToken}:                      _ReduceNilToOptionalTraitPropertiesAction,
	{_State172, _WildcardMarker}:                  _ReduceSliceToInitializableTypeAction,
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
	{_State208, _EndMarker}:                       _ReduceAliasToTypeDefAction,
	{_State209, _WildcardMarker}:                  _ReduceDefinitionToTypeDefAction,
	{_State211, _WildcardMarker}:                  _ReduceArgToParameterDefAction,
	{_State212, IdentifierToken}:                  _ReduceReceiverToOptionalReceiverAction,
	{_State214, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State218, _WildcardMarker}:                  _ReduceFieldDefToEnumValueDefAction,
	{_State221, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State223, _WildcardMarker}:                  _ReduceParameterDeclToParameterDeclsAction,
	{_State224, RparenToken}:                      _ReduceParameterDeclsToOptionalParameterDeclsAction,
	{_State225, _WildcardMarker}:                  _ReduceUnamedToParameterDeclAction,
	{_State226, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State227, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State228, _WildcardMarker}:                  _ReduceExplicitToFieldDefAction,
	{_State232, _WildcardMarker}:                  _ReduceToImplicitEnumDefAction,
	{_State234, _WildcardMarker}:                  _ReduceToImplicitStructDefAction,
	{_State236, _WildcardMarker}:                  _ReduceFieldDefToTraitPropertyAction,
	{_State237, _WildcardMarker}:                  _ReduceMethodSignatureToTraitPropertyAction,
	{_State239, RparenToken}:                      _ReduceTraitPropertiesToOptionalTraitPropertiesAction,
	{_State240, _WildcardMarker}:                  _ReduceTraitPropertyToTraitPropertiesAction,
	{_State241, _WildcardMarker}:                  _ReduceTraitUnionToValueTypeAction,
	{_State244, _WildcardMarker}:                  _ReduceTraitIntersectToValueTypeAction,
	{_State245, _WildcardMarker}:                  _ReduceTraitDifferenceToValueTypeAction,
	{_State246, _WildcardMarker}:                  _ReduceNamedToArgumentAction,
	{_State247, _WildcardMarker}:                  _ReduceAddToArgumentsAction,
	{_State248, _WildcardMarker}:                  _ReduceExpressionToOptionalExpressionAction,
	{_State249, _WildcardMarker}:                  _ReduceAddToColonExpressionsAction,
	{_State250, _WildcardMarker}:                  _ReducePairToColonExpressionsAction,
	{_State253, _WildcardMarker}:                  _ReduceToExplicitStructDefAction,
	{_State255, _WildcardMarker}:                  _ReduceBindingToOptionalGenericBindingAction,
	{_State256, _WildcardMarker}:                  _ReduceIndexToAccessExprAction,
	{_State257, RparenToken}:                      _ReduceArgumentsToOptionalArgumentsAction,
	{_State259, _WildcardMarker}:                  _ReduceExplicitToInitializeExprAction,
	{_State260, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State263, DoToken}:                          _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State265, _WildcardMarker}:                  _ReduceNoElseToIfExprAction,
	{_State266, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State267, _WildcardMarker}:                  _ReduceBreakToJumpTypeAction,
	{_State268, _WildcardMarker}:                  _ReduceContinueToJumpTypeAction,
	{_State269, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State270, _EndMarker}:                       _ReduceToBlockBodyAction,
	{_State271, _WildcardMarker}:                  _ReduceReturnToJumpTypeAction,
	{_State272, _WildcardMarker}:                  _ReduceAccessExprToPostfixUnaryExprAction,
	{_State272, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State273, _WildcardMarker}:                  _ReduceExpressionToExpressionsAction,
	{_State274, _WildcardMarker}:                  _ReduceExpressionOrImplicitStructToStatementBodyAction,
	{_State275, _WildcardMarker}:                  _ReduceJumpStatementToStatementBodyAction,
	{_State276, _WildcardMarker}:                  _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State277, _WildcardMarker}:                  _ReduceAddToStatementsAction,
	{_State279, _WildcardMarker}:                  _ReduceUnsafeStatementToStatementBodyAction,
	{_State282, _EndMarker}:                       _ReduceWithSpecToPackageDefAction,
	{_State283, _WildcardMarker}:                  _ReduceAddToPackageStatementsAction,
	{_State285, _WildcardMarker}:                  _ReduceToPackageStatementBodyAction,
	{_State286, _WildcardMarker}:                  _ReduceConstrainedToGenericParameterDefAction,
	{_State288, _WildcardMarker}:                  _ReduceGenericToOptionalGenericParametersAction,
	{_State290, _WildcardMarker}:                  _ReduceVarargToParameterDefAction,
	{_State291, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State293, _WildcardMarker}:                  _ReduceReturnableTypeToReturnTypeAction,
	{_State294, _WildcardMarker}:                  _ReduceAddToParameterDefsAction,
	{_State297, _WildcardMarker}:                  _ReduceToExplicitEnumDefAction,
	{_State300, _WildcardMarker}:                  _ReduceUnnamedVarargToParameterDeclAction,
	{_State302, _WildcardMarker}:                  _ReduceArgToParameterDeclAction,
	{_State303, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State305, _WildcardMarker}:                  _ReduceExternNamedToAtomTypeAction,
	{_State306, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State307, _WildcardMarker}:                  _ReduceDefaultToEnumValueDefAction,
	{_State308, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State309, _WildcardMarker}:                  _ReduceAddToImplicitFieldDefsAction,
	{_State311, _WildcardMarker}:                  _ReduceToTraitDefAction,
	{_State314, _WildcardMarker}:                  _ReduceMapToInitializableTypeAction,
	{_State315, _WildcardMarker}:                  _ReduceArrayToInitializableTypeAction,
	{_State316, _WildcardMarker}:                  _ReduceExplicitToExplicitFieldDefsAction,
	{_State317, _WildcardMarker}:                  _ReduceImplicitToExplicitFieldDefsAction,
	{_State318, _WildcardMarker}:                  _ReduceAddToGenericArgumentsAction,
	{_State319, _WildcardMarker}:                  _ReduceToCallExprAction,
	{_State320, _EndMarker}:                       _ReduceDoWhileToLoopExprAction,
	{_State322, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State322, DoToken}:                          _ReduceNilToOptionalSequenceExprAction,
	{_State323, _EndMarker}:                       _ReduceWhileToLoopExprAction,
	{_State325, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State326, NewlinesToken}:                    _ReduceAsyncToStatementBodyAction,
	{_State326, SemicolonToken}:                   _ReduceAsyncToStatementBodyAction,
	{_State326, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State327, NewlinesToken}:                    _ReduceDeferToStatementBodyAction,
	{_State327, SemicolonToken}:                   _ReduceDeferToStatementBodyAction,
	{_State327, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State328, _WildcardMarker}:                  _ReduceAddAssignToBinaryOpAssignAction,
	{_State329, _WildcardMarker}:                  _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State330, _WildcardMarker}:                  _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State331, _WildcardMarker}:                  _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State332, _WildcardMarker}:                  _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State333, _WildcardMarker}:                  _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State334, _WildcardMarker}:                  _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State335, _WildcardMarker}:                  _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State336, _WildcardMarker}:                  _ReduceDivAssignToBinaryOpAssignAction,
	{_State337, _WildcardMarker}:                  _ReduceModAssignToBinaryOpAssignAction,
	{_State338, _WildcardMarker}:                  _ReduceMulAssignToBinaryOpAssignAction,
	{_State339, _WildcardMarker}:                  _ReduceSubAssignToBinaryOpAssignAction,
	{_State340, _WildcardMarker}:                  _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State341, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State342, _WildcardMarker}:                  _ReduceUnaryOpAssignStatementToStatementBodyAction,
	{_State343, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State344, _WildcardMarker}:                  _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State345, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State345, NewlinesToken}:                    _ReduceNilToOptionalExpressionsAction,
	{_State345, SemicolonToken}:                   _ReduceNilToOptionalExpressionsAction,
	{_State346, _WildcardMarker}:                  _ReduceImplicitToStatementAction,
	{_State347, _WildcardMarker}:                  _ReduceExplicitToStatementAction,
	{_State348, ColonToken}:                       _ReduceToMatchPatternAction,
	{_State349, _WildcardMarker}:                  _ReduceCaseBranchToCaseBranchesAction,
	{_State350, RbraceToken}:                      _ReduceNilToOptionalDefaultBranchAction,
	{_State351, _WildcardMarker}:                  _ReduceToUnsafeStatementAction,
	{_State352, _WildcardMarker}:                  _ReduceImplicitToPackageStatementAction,
	{_State353, _WildcardMarker}:                  _ReduceExplicitToPackageStatementAction,
	{_State354, _WildcardMarker}:                  _ReduceAddToGenericParameterDefsAction,
	{_State355, _EndMarker}:                       _ReduceConstrainedDefToTypeDefAction,
	{_State357, _WildcardMarker}:                  _ReduceToAnonymousFuncExprAction,
	{_State358, RparenToken}:                      _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State359, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State359, RparenToken}:                      _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State360, RparenToken}:                      _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State361, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State361, RparenToken}:                      _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State362, _WildcardMarker}:                  _ReduceVarargToParameterDeclAction,
	{_State363, _WildcardMarker}:                  _ReduceToFuncTypeAction,
	{_State364, _WildcardMarker}:                  _ReduceAddToParameterDeclsAction,
	{_State365, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State366, _WildcardMarker}:                  _ReduceExplicitToTraitPropertiesAction,
	{_State367, _WildcardMarker}:                  _ReduceImplicitToTraitPropertiesAction,
	{_State368, _EndMarker}:                       _ReduceIteratorToLoopExprAction,
	{_State370, _EndMarker}:                       _ReduceIfElseToIfExprAction,
	{_State371, _EndMarker}:                       _ReduceMultiIfElseToIfExprAction,
	{_State372, _WildcardMarker}:                  _ReduceBinaryOpAssignStatementToStatementBodyAction,
	{_State373, _WildcardMarker}:                  _ReduceAddToExpressionsAction,
	{_State374, _WildcardMarker}:                  _ReduceExpressionsToOptionalExpressionsAction,
	{_State375, _WildcardMarker}:                  _ReduceToJumpStatementAction,
	{_State378, _WildcardMarker}:                  _ReduceAddToCaseBranchesAction,
	{_State379, RbraceToken}:                      _ReduceDefaultBranchToOptionalDefaultBranchAction,
	{_State381, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State384, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State385, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State386, _EndMarker}:                       _ReduceToSwitchExprAction,
	{_State388, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State389, _EndMarker}:                       _ReduceForToLoopExprAction,
	{_State390, _WildcardMarker}:                  _ReduceToCaseBranchAction,
	{_State391, RbraceToken}:                      _ReduceToDefaultBranchAction,
	{_State392, _EndMarker}:                       _ReduceToNamedFuncDefAction,
	{_State393, _WildcardMarker}:                  _ReduceToMethodSignatureAction,
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
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 11
      optional_label_decl -> State 50
      sequence_expr -> State 55
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 6:
    Kernel Items:
      #accept: ^.lex_internal_tokens
    Reduce:
      (nil)
    Goto:
      SPACES -> State 57
      COMMENT -> State 56
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
      UNSAFE -> State 58
      TYPE -> State 17
      FUNC -> State 18
      definitions -> State 60
      definition -> State 59
      unsafe_statement -> State 64
      type_def -> State 63
      named_func_def -> State 61
      package_def -> State 62

  State 16:
    Kernel Items:
      package_def: PACKAGE.IDENTIFIER
      package_def: PACKAGE.IDENTIFIER LPAREN package_statements RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 65

  State 17:
    Kernel Items:
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE.IDENTIFIER EQUAL value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 66

  State 18:
    Kernel Items:
      named_func_def: FUNC.optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      LPAREN -> State 67
      optional_receiver -> State 68

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
      LPAREN -> State 69

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
      initializable_type: LBRACKET.value_type RBRACKET
      initializable_type: LBRACKET.value_type COMMA INTEGER_LITERAL RBRACKET
      initializable_type: LBRACKET.value_type COLON value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 90
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 28:
    Kernel Items:
      atom_expr: LEX_ERROR., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 29:
    Kernel Items:
      initialize_expr: LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 91
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 95
      optional_label_decl -> State 50
      sequence_expr -> State 55
      block_expr -> State 42
      call_expr -> State 43
      arguments -> State 93
      argument -> State 92
      colon_expressions -> State 94
      optional_expression -> State 96
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 30:
    Kernel Items:
      prefix_unary_op: MUL., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 31:
    Kernel Items:
      prefix_unary_op: NOT., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      literal: RUNE_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 33:
    Kernel Items:
      literal: STRING_LITERAL., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 34:
    Kernel Items:
      explicit_struct_def: STRUCT.LPAREN optional_explicit_field_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 97

  State 35:
    Kernel Items:
      prefix_unary_op: SUB., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 36:
    Kernel Items:
      literal: TRUE., *
    Reduce:
      * -> [literal]
    Goto:
      (nil)

  State 37:
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

  State 38:
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

  State 39:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 108

  State 40:
    Kernel Items:
      atom_expr: anonymous_func_expr., *
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
      EQUAL -> State 109
      NOT_EQUAL -> State 114
      LESS -> State 112
      LESS_OR_EQUAL -> State 113
      GREATER -> State 110
      GREATER_OR_EQUAL -> State 111
      cmp_op -> State 115

  State 45:
    Kernel Items:
      initializable_type: explicit_struct_def., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 46:
    Kernel Items:
      initialize_expr: initializable_type.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 116

  State 47:
    Kernel Items:
      atom_expr: initialize_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      atom_expr: literal., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 49:
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

  State 50:
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

  State 51:
    Kernel Items:
      sequence_expr: or_expr., *
      or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    Goto:
      OR -> State 133

  State 52:
    Kernel Items:
      prefix_unary_expr: postfix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 53:
    Kernel Items:
      mul_expr: prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 54:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op.prefix_unary_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 135
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 55:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 56:
    Kernel Items:
      lex_internal_tokens: COMMENT., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 57:
    Kernel Items:
      lex_internal_tokens: SPACES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 136

  State 59:
    Kernel Items:
      definitions: definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 60:
    Kernel Items:
      optional_definitions: optional_newlines definitions.optional_newlines
      definitions: definitions.NEWLINES definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      NEWLINES -> State 137
      optional_newlines -> State 138

  State 61:
    Kernel Items:
      definition: named_func_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      definition: package_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 63:
    Kernel Items:
      definition: type_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 64:
    Kernel Items:
      definition: unsafe_statement., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER., *
      package_def: PACKAGE IDENTIFIER.LPAREN package_statements RPAREN
    Reduce:
      * -> [package_def]
    Goto:
      LPAREN -> State 139

  State 66:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE IDENTIFIER.EQUAL value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 140
      EQUAL -> State 141
      optional_generic_parameters -> State 142

  State 67:
    Kernel Items:
      optional_receiver: LPAREN.parameter_def RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 143
      parameter_def -> State 144

  State 68:
    Kernel Items:
      named_func_def: FUNC optional_receiver.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 145

  State 69:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 143
      parameter_def -> State 147
      parameter_defs -> State 148
      optional_parameter_defs -> State 146

  State 70:
    Kernel Items:
      returnable_type: BIT_AND.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 149
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 71:
    Kernel Items:
      returnable_type: BIT_NEG.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 150
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 72:
    Kernel Items:
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 98
      optional_generic_binding -> State 151

  State 73:
    Kernel Items:
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 152

  State 74:
    Kernel Items:
      returnable_type: EXCLAIM.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 153
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 75:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 154

  State 76:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOT -> State 155
      DOLLAR_LBRACKET -> State 98
      optional_generic_binding -> State 156

  State 77:
    Kernel Items:
      atom_type: LEX_ERROR., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 78:
    Kernel Items:
      implicit_struct_def: LPAREN.optional_implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      IDENTIFIER -> State 157
      UNSAFE -> State 58
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      unsafe_statement -> State 163
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 164
      field_def -> State 159
      implicit_field_defs -> State 161
      optional_implicit_field_defs -> State 162
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      enum_value_def -> State 158
      implicit_enum_value_defs -> State 160
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 79:
    Kernel Items:
      returnable_type: QUESTION.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 165
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 80:
    Kernel Items:
      returnable_type: TILDE_TILDE.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 166
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 81:
    Kernel Items:
      trait_def: TRAIT.LPAREN optional_trait_properties RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 167

  State 82:
    Kernel Items:
      returnable_type: atom_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      atom_type: explicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 84:
    Kernel Items:
      atom_type: func_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 85:
    Kernel Items:
      atom_type: implicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 86:
    Kernel Items:
      atom_type: implicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 87:
    Kernel Items:
      atom_type: initializable_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      value_type: returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      atom_type: trait_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 90:
    Kernel Items:
      initializable_type: LBRACKET value_type.RBRACKET
      initializable_type: LBRACKET value_type.COMMA INTEGER_LITERAL RBRACKET
      initializable_type: LBRACKET value_type.COLON value_type RBRACKET
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 172
      COMMA -> State 170
      COLON -> State 169
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

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
      IDENTIFIER -> State 157
      UNSAFE -> State 58
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      unsafe_statement -> State 163
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 164
      field_def -> State 180
      implicit_struct_def -> State 86
      explicit_field_defs -> State 179
      optional_explicit_field_defs -> State 181
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 98:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      optional_generic_arguments -> State 183
      generic_arguments -> State 182
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 184
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

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
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 91
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 95
      optional_label_decl -> State 50
      sequence_expr -> State 55
      block_expr -> State 42
      call_expr -> State 43
      argument -> State 186
      colon_expressions -> State 94
      optional_expression -> State 96
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

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
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 188
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 108:
    Kernel Items:
      and_expr: and_expr AND.cmp_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 189
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

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
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 190
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 116:
    Kernel Items:
      initialize_expr: initializable_type LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 91
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 95
      optional_label_decl -> State 50
      sequence_expr -> State 55
      block_expr -> State 42
      call_expr -> State 43
      arguments -> State 191
      argument -> State 92
      colon_expressions -> State 94
      optional_expression -> State 96
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

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
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 192
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

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
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      IN -> State 194
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      SEMICOLON -> State 195
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      sequence_expr -> State 196
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 126:
    Kernel Items:
      if_expr: IF.sequence_expr block_body
      if_expr: IF.sequence_expr block_body ELSE block_body
      if_expr: IF.sequence_expr block_body ELSE if_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      sequence_expr -> State 197
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 127:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 198

  State 128:
    Kernel Items:
      switch_expr: SWITCH.sequence_expr LBRACE case_branches optional_default_branch RBRACE
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      sequence_expr -> State 199
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

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
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 200
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

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
      UNSAFE -> State 58
      TYPE -> State 17
      FUNC -> State 18
      definition -> State 202
      unsafe_statement -> State 64
      type_def -> State 63
      named_func_def -> State 61
      package_def -> State 62

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
      type_def: TYPE IDENTIFIER EQUAL.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 208
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 142:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 209
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 143:
    Kernel Items:
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      DOTDOTDOT -> State 210
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 211
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 144:
    Kernel Items:
      optional_receiver: LPAREN parameter_def.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 212

  State 145:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 140
      optional_generic_parameters -> State 213

  State 146:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 214

  State 147:
    Kernel Items:
      parameter_defs: parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 148:
    Kernel Items:
      parameter_defs: parameter_defs.COMMA parameter_def
      optional_parameter_defs: parameter_defs., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      COMMA -> State 215

  State 149:
    Kernel Items:
      returnable_type: BIT_AND returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 150:
    Kernel Items:
      returnable_type: BIT_NEG returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 151:
    Kernel Items:
      atom_type: DOT optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 152:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      UNSAFE -> State 58
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      unsafe_statement -> State 163
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 164
      field_def -> State 218
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      enum_value_def -> State 216
      implicit_enum_value_defs -> State 219
      implicit_enum_def -> State 85
      explicit_enum_value_defs -> State 217
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 153:
    Kernel Items:
      returnable_type: EXCLAIM returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 154:
    Kernel Items:
      func_type: FUNC LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 221
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      DOTDOTDOT -> State 220
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 225
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      parameter_decl -> State 223
      parameter_decls -> State 224
      optional_parameter_decls -> State 222
      func_type -> State 84

  State 155:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 226

  State 156:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 157:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 227
      QUESTION -> State 79
      EXCLAIM -> State 74
      DOLLAR_LBRACKET -> State 98
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      optional_generic_binding -> State 156
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 228
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 158:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 229

  State 159:
    Kernel Items:
      implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 230

  State 160:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 232
      OR -> State 231

  State 161:
    Kernel Items:
      implicit_field_defs: implicit_field_defs.COMMA field_def
      optional_implicit_field_defs: implicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      COMMA -> State 233

  State 162:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 234

  State 163:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 164:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

  State 165:
    Kernel Items:
      returnable_type: QUESTION returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 166:
    Kernel Items:
      returnable_type: TILDE_TILDE returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 167:
    Kernel Items:
      trait_def: TRAIT LPAREN.optional_trait_properties RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      IDENTIFIER -> State 157
      UNSAFE -> State 58
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 235
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      unsafe_statement -> State 163
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 164
      field_def -> State 236
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_property -> State 240
      trait_properties -> State 239
      optional_trait_properties -> State 238
      trait_def -> State 89
      func_type -> State 84
      method_signature -> State 237

  State 168:
    Kernel Items:
      value_type: value_type ADD.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 241
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 169:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON.value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 242
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 170:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 243

  State 171:
    Kernel Items:
      value_type: value_type MUL.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 244
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 172:
    Kernel Items:
      initializable_type: LBRACKET value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 173:
    Kernel Items:
      value_type: value_type SUB.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 245
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 174:
    Kernel Items:
      argument: IDENTIFIER ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 246
      optional_label_decl -> State 50
      sequence_expr -> State 55
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 175:
    Kernel Items:
      arguments: arguments COMMA.argument
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 91
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 95
      optional_label_decl -> State 50
      sequence_expr -> State 55
      block_expr -> State 42
      call_expr -> State 43
      argument -> State 247
      colon_expressions -> State 94
      optional_expression -> State 96
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

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
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 248
      optional_label_decl -> State 50
      sequence_expr -> State 55
      block_expr -> State 42
      call_expr -> State 43
      optional_expression -> State 249
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

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
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 248
      optional_label_decl -> State 50
      sequence_expr -> State 55
      block_expr -> State 42
      call_expr -> State 43
      optional_expression -> State 250
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 179:
    Kernel Items:
      explicit_field_defs: explicit_field_defs.NEWLINES field_def
      explicit_field_defs: explicit_field_defs.COMMA field_def
      optional_explicit_field_defs: explicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      NEWLINES -> State 252
      COMMA -> State 251

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
      RPAREN -> State 253

  State 182:
    Kernel Items:
      optional_generic_arguments: generic_arguments., RBRACKET
      generic_arguments: generic_arguments.COMMA value_type
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 254

  State 183:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 255

  State 184:
    Kernel Items:
      generic_arguments: value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

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
      RBRACKET -> State 256

  State 187:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN.optional_arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_arguments]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 91
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 95
      optional_label_decl -> State 50
      sequence_expr -> State 55
      block_expr -> State 42
      call_expr -> State 43
      optional_arguments -> State 258
      arguments -> State 257
      argument -> State 92
      colon_expressions -> State 94
      optional_expression -> State 96
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

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
      RPAREN -> State 259
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
      FOR -> State 260

  State 194:
    Kernel Items:
      loop_expr: FOR IN.sequence_expr DO block_body
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      sequence_expr -> State 261
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 195:
    Kernel Items:
      loop_expr: FOR SEMICOLON.optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      LBRACE -> [optional_label_decl]
      SEMICOLON -> [optional_sequence_expr]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      optional_sequence_expr -> State 262
      sequence_expr -> State 263
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 196:
    Kernel Items:
      loop_expr: FOR sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 264

  State 197:
    Kernel Items:
      if_expr: IF sequence_expr.block_body
      if_expr: IF sequence_expr.block_body ELSE block_body
      if_expr: IF sequence_expr.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 127
      block_body -> State 265

  State 198:
    Kernel Items:
      block_body: LBRACE statements.RBRACE
      statements: statements.statement
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      RETURN -> State 271
      BREAK -> State 267
      CONTINUE -> State 268
      UNSAFE -> State 58
      STRUCT -> State 34
      FUNC -> State 23
      ASYNC -> State 266
      DEFER -> State 269
      LABEL_DECL -> State 26
      RBRACE -> State 270
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 273
      optional_label_decl -> State 50
      sequence_expr -> State 55
      block_expr -> State 42
      statement -> State 277
      statement_body -> State 278
      unsafe_statement -> State 279
      jump_statement -> State 275
      jump_type -> State 276
      expressions -> State 274
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 272
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 199:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.LBRACE case_branches optional_default_branch RBRACE
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 280

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
      GREATER -> State 281

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
      UNSAFE -> State 58
      RPAREN -> State 282
      unsafe_statement -> State 285
      package_statement_body -> State 284
      package_statement -> State 283

  State 204:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 286
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

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
      COMMA -> State 287

  State 207:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 288

  State 208:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER EQUAL value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

  State 209:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., *
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 289
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

  State 210:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 290
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 211:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

  State 212:
    Kernel Items:
      optional_receiver: LPAREN parameter_def RPAREN., IDENTIFIER
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 291

  State 214:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 293
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      return_type -> State 292
      func_type -> State 84

  State 215:
    Kernel Items:
      parameter_defs: parameter_defs COMMA.parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 143
      parameter_def -> State 294

  State 216:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 295
      OR -> State 296

  State 217:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 297

  State 218:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 230

  State 219:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 298
      OR -> State 299

  State 220:
    Kernel Items:
      parameter_decl: DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 300
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 221:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 227
      QUESTION -> State 79
      EXCLAIM -> State 74
      DOLLAR_LBRACKET -> State 98
      DOTDOTDOT -> State 301
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      optional_generic_binding -> State 156
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 302
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 222:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 303

  State 223:
    Kernel Items:
      parameter_decls: parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 224:
    Kernel Items:
      parameter_decls: parameter_decls.COMMA parameter_decl
      optional_parameter_decls: parameter_decls., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      COMMA -> State 304

  State 225:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

  State 226:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 98
      optional_generic_binding -> State 305

  State 227:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 226
      DOLLAR_LBRACKET -> State 98
      optional_generic_binding -> State 151

  State 228:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

  State 229:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      UNSAFE -> State 58
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      unsafe_statement -> State 163
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 164
      field_def -> State 218
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      enum_value_def -> State 306
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 230:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 307

  State 231:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      UNSAFE -> State 58
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      unsafe_statement -> State 163
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 164
      field_def -> State 218
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      enum_value_def -> State 308
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 232:
    Kernel Items:
      implicit_enum_def: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_def]
    Goto:
      (nil)

  State 233:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      UNSAFE -> State 58
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      unsafe_statement -> State 163
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 164
      field_def -> State 309
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 234:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_def]
    Goto:
      (nil)

  State 235:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 310
      LPAREN -> State 154

  State 236:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 237:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 238:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 311

  State 239:
    Kernel Items:
      trait_properties: trait_properties.NEWLINES trait_property
      trait_properties: trait_properties.COMMA trait_property
      optional_trait_properties: trait_properties., RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      NEWLINES -> State 313
      COMMA -> State 312

  State 240:
    Kernel Items:
      trait_properties: trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      value_type: value_type ADD returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 242:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type.RBRACKET
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 314
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

  State 243:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 315

  State 244:
    Kernel Items:
      value_type: value_type MUL returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 245:
    Kernel Items:
      value_type: value_type SUB returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 246:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 247:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 248:
    Kernel Items:
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      (nil)

  State 249:
    Kernel Items:
      colon_expressions: colon_expressions COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 250:
    Kernel Items:
      colon_expressions: optional_expression COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 251:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      UNSAFE -> State 58
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      unsafe_statement -> State 163
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 164
      field_def -> State 316
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 252:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      UNSAFE -> State 58
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      unsafe_statement -> State 163
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 164
      field_def -> State 317
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 253:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_def]
    Goto:
      (nil)

  State 254:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 318
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 255:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 256:
    Kernel Items:
      access_expr: access_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 257:
    Kernel Items:
      optional_arguments: arguments., RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 175

  State 258:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 319

  State 259:
    Kernel Items:
      initialize_expr: initializable_type LPAREN arguments RPAREN., *
    Reduce:
      * -> [initialize_expr]
    Goto:
      (nil)

  State 260:
    Kernel Items:
      loop_expr: DO block_body FOR.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      sequence_expr -> State 320
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 261:
    Kernel Items:
      loop_expr: FOR IN sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 321

  State 262:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 322

  State 263:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 264:
    Kernel Items:
      loop_expr: FOR sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 127
      block_body -> State 323

  State 265:
    Kernel Items:
      if_expr: IF sequence_expr block_body., *
      if_expr: IF sequence_expr block_body.ELSE block_body
      if_expr: IF sequence_expr block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 324

  State 266:
    Kernel Items:
      statement_body: ASYNC.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      block_expr -> State 42
      call_expr -> State 326
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 325
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 267:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 268:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 269:
    Kernel Items:
      statement_body: DEFER.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      block_expr -> State 42
      call_expr -> State 327
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 325
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 270:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 271:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 272:
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
      ADD_ASSIGN -> State 328
      SUB_ASSIGN -> State 339
      MUL_ASSIGN -> State 338
      DIV_ASSIGN -> State 336
      MOD_ASSIGN -> State 337
      ADD_ONE_ASSIGN -> State 329
      SUB_ONE_ASSIGN -> State 340
      BIT_NEG_ASSIGN -> State 332
      BIT_AND_ASSIGN -> State 330
      BIT_OR_ASSIGN -> State 333
      BIT_XOR_ASSIGN -> State 335
      BIT_LSHIFT_ASSIGN -> State 331
      BIT_RSHIFT_ASSIGN -> State 334
      unary_op_assign -> State 342
      binary_op_assign -> State 341
      optional_generic_binding -> State 102

  State 273:
    Kernel Items:
      expressions: expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 274:
    Kernel Items:
      statement_body: expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [statement_body]
    Goto:
      COMMA -> State 343

  State 275:
    Kernel Items:
      statement_body: jump_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 276:
    Kernel Items:
      jump_statement: jump_type.optional_jump_label optional_expressions
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 344
      optional_jump_label -> State 345

  State 277:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 346
      SEMICOLON -> State 347

  State 279:
    Kernel Items:
      statement_body: unsafe_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 280:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE.case_branches optional_default_branch RBRACE
    Reduce:
      (nil)
    Goto:
      CASE -> State 348
      case_branches -> State 350
      case_branch -> State 349

  State 281:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 351

  State 282:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements RPAREN., $
    Reduce:
      $ -> [package_def]
    Goto:
      (nil)

  State 283:
    Kernel Items:
      package_statements: package_statements package_statement., *
    Reduce:
      * -> [package_statements]
    Goto:
      (nil)

  State 284:
    Kernel Items:
      package_statement: package_statement_body.NEWLINES
      package_statement: package_statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 352
      SEMICOLON -> State 353

  State 285:
    Kernel Items:
      package_statement_body: unsafe_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 286:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

  State 287:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA.generic_parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 204
      generic_parameter_def -> State 354

  State 288:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 289:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 355
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 290:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_def: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

  State 291:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 143
      parameter_def -> State 147
      parameter_defs -> State 148
      optional_parameter_defs -> State 356

  State 292:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 127
      block_body -> State 357

  State 293:
    Kernel Items:
      return_type: returnable_type., *
    Reduce:
      * -> [return_type]
    Goto:
      (nil)

  State 294:
    Kernel Items:
      parameter_defs: parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 295:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      UNSAFE -> State 58
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      unsafe_statement -> State 163
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 164
      field_def -> State 218
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      enum_value_def -> State 358
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 296:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      UNSAFE -> State 58
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      unsafe_statement -> State 163
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 164
      field_def -> State 218
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      enum_value_def -> State 359
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 297:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_def]
    Goto:
      (nil)

  State 298:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      UNSAFE -> State 58
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      unsafe_statement -> State 163
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 164
      field_def -> State 218
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      enum_value_def -> State 360
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 299:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      UNSAFE -> State 58
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      unsafe_statement -> State 163
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 164
      field_def -> State 218
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      enum_value_def -> State 361
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 300:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

  State 301:
    Kernel Items:
      parameter_decl: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 362
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      func_type -> State 84

  State 302:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

  State 303:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 293
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      return_type -> State 363
      func_type -> State 84

  State 304:
    Kernel Items:
      parameter_decls: parameter_decls COMMA.parameter_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 221
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      DOTDOTDOT -> State 220
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 225
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      parameter_decl -> State 364
      func_type -> State 84

  State 305:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 306:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 307:
    Kernel Items:
      enum_value_def: field_def ASSIGN DEFAULT., *
    Reduce:
      * -> [enum_value_def]
    Goto:
      (nil)

  State 308:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 309:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [implicit_field_defs]
    Goto:
      (nil)

  State 310:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 365

  State 311:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties RPAREN., *
    Reduce:
      * -> [trait_def]
    Goto:
      (nil)

  State 312:
    Kernel Items:
      trait_properties: trait_properties COMMA.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      UNSAFE -> State 58
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 235
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      unsafe_statement -> State 163
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 164
      field_def -> State 236
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_property -> State 366
      trait_def -> State 89
      func_type -> State 84
      method_signature -> State 237

  State 313:
    Kernel Items:
      trait_properties: trait_properties NEWLINES.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      UNSAFE -> State 58
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 235
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      unsafe_statement -> State 163
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 164
      field_def -> State 236
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_property -> State 367
      trait_def -> State 89
      func_type -> State 84
      method_signature -> State 237

  State 314:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 315:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 316:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 317:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 318:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

  State 319:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 320:
    Kernel Items:
      loop_expr: DO block_body FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 321:
    Kernel Items:
      loop_expr: FOR IN sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 127
      block_body -> State 368

  State 322:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON.optional_sequence_expr DO block_body
    Reduce:
      DO -> [optional_sequence_expr]
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      optional_sequence_expr -> State 369
      sequence_expr -> State 263
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 323:
    Kernel Items:
      loop_expr: FOR sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 324:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE.block_body
      if_expr: IF sequence_expr block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 126
      LBRACE -> State 127
      if_expr -> State 371
      block_body -> State 370

  State 325:
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

  State 326:
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

  State 327:
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

  State 328:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 329:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 330:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 331:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 333:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 334:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 335:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 336:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 337:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 338:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 339:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 340:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      statement_body: access_expr binary_op_assign.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 372
      optional_label_decl -> State 50
      sequence_expr -> State 55
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 342:
    Kernel Items:
      statement_body: access_expr unary_op_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 343:
    Kernel Items:
      expressions: expressions COMMA.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 373
      optional_label_decl -> State 50
      sequence_expr -> State 55
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 344:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 345:
    Kernel Items:
      jump_statement: jump_type optional_jump_label.optional_expressions
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_expressions]
      SEMICOLON -> [optional_expressions]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 273
      optional_label_decl -> State 50
      sequence_expr -> State 55
      block_expr -> State 42
      expressions -> State 374
      optional_expressions -> State 375
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 346:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 347:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 348:
    Kernel Items:
      case_branch: CASE.match_pattern COLON sequence_expr
    Reduce:
      COLON -> [match_pattern]
    Goto:
      match_pattern -> State 376

  State 349:
    Kernel Items:
      case_branches: case_branch., *
    Reduce:
      * -> [case_branches]
    Goto:
      (nil)

  State 350:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE case_branches.optional_default_branch RBRACE
      case_branches: case_branches.case_branch
    Reduce:
      RBRACE -> [optional_default_branch]
    Goto:
      CASE -> State 348
      DEFAULT -> State 377
      case_branch -> State 378
      optional_default_branch -> State 380
      default_branch -> State 379

  State 351:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 352:
    Kernel Items:
      package_statement: package_statement_body NEWLINES., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 353:
    Kernel Items:
      package_statement: package_statement_body SEMICOLON., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 354:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 355:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

  State 356:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 381

  State 357:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 358:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 359:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
      explicit_enum_value_defs: enum_value_def OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 360:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 361:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 362:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 168
      SUB -> State 173
      MUL -> State 171

  State 363:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [func_type]
    Goto:
      (nil)

  State 364:
    Kernel Items:
      parameter_decls: parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 365:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 221
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      DOTDOTDOT -> State 220
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 88
      value_type -> State 225
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      parameter_decl -> State 223
      parameter_decls -> State 224
      optional_parameter_decls -> State 382
      func_type -> State 84

  State 366:
    Kernel Items:
      trait_properties: trait_properties COMMA trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 367:
    Kernel Items:
      trait_properties: trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 368:
    Kernel Items:
      loop_expr: FOR IN sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 369:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 383

  State 370:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 371:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 372:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 373:
    Kernel Items:
      expressions: expressions COMMA expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 374:
    Kernel Items:
      expressions: expressions.COMMA expression
      optional_expressions: expressions., *
    Reduce:
      * -> [optional_expressions]
    Goto:
      COMMA -> State 343

  State 375:
    Kernel Items:
      jump_statement: jump_type optional_jump_label optional_expressions., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 376:
    Kernel Items:
      case_branch: CASE match_pattern.COLON sequence_expr
    Reduce:
      (nil)
    Goto:
      COLON -> State 384

  State 377:
    Kernel Items:
      default_branch: DEFAULT.COLON sequence_expr
    Reduce:
      (nil)
    Goto:
      COLON -> State 385

  State 378:
    Kernel Items:
      case_branches: case_branches case_branch., *
    Reduce:
      * -> [case_branches]
    Goto:
      (nil)

  State 379:
    Kernel Items:
      optional_default_branch: default_branch., RBRACE
    Reduce:
      RBRACE -> [optional_default_branch]
    Goto:
      (nil)

  State 380:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE case_branches optional_default_branch.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 386

  State 381:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 293
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      return_type -> State 387
      func_type -> State 84

  State 382:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 388

  State 383:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 127
      block_body -> State 389

  State 384:
    Kernel Items:
      case_branch: CASE match_pattern COLON.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      sequence_expr -> State 390
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 385:
    Kernel Items:
      default_branch: DEFAULT COLON.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 24
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      optional_label_decl -> State 134
      sequence_expr -> State 391
      block_expr -> State 42
      call_expr -> State 43
      atom_expr -> State 41
      literal -> State 48
      initialize_expr -> State 47
      access_expr -> State 37
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 38
      cmp_expr -> State 44
      and_expr -> State 39
      or_expr -> State 51
      initializable_type -> State 46
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 40

  State 386:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE case_branches optional_default_branch RBRACE., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 387:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 127
      block_body -> State 392

  State 388:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 76
      STRUCT -> State 34
      ENUM -> State 73
      TRAIT -> State 81
      FUNC -> State 75
      LPAREN -> State 78
      LBRACKET -> State 27
      DOT -> State 72
      QUESTION -> State 79
      EXCLAIM -> State 74
      TILDE_TILDE -> State 80
      BIT_NEG -> State 71
      BIT_AND -> State 70
      LEX_ERROR -> State 77
      initializable_type -> State 87
      atom_type -> State 82
      returnable_type -> State 293
      implicit_struct_def -> State 86
      explicit_struct_def -> State 45
      implicit_enum_def -> State 85
      explicit_enum_def -> State 83
      trait_def -> State 89
      return_type -> State 393
      func_type -> State 84

  State 389:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 390:
    Kernel Items:
      case_branch: CASE match_pattern COLON sequence_expr., *
    Reduce:
      * -> [case_branch]
    Goto:
      (nil)

  State 391:
    Kernel Items:
      default_branch: DEFAULT COLON sequence_expr., RBRACE
    Reduce:
      RBRACE -> [default_branch]
    Goto:
      (nil)

  State 392:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 393:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

Number of states: 393
Number of shift actions: 2561
Number of reduce actions: 318
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
*/
