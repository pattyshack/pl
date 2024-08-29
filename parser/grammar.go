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
	VarToken             = SymbolId(287)
	LabelDeclToken       = SymbolId(288)
	JumpLabelToken       = SymbolId(289)
	LbraceToken          = SymbolId(290)
	RbraceToken          = SymbolId(291)
	LparenToken          = SymbolId(292)
	RparenToken          = SymbolId(293)
	LbracketToken        = SymbolId(294)
	RbracketToken        = SymbolId(295)
	DotToken             = SymbolId(296)
	CommaToken           = SymbolId(297)
	QuestionToken        = SymbolId(298)
	SemicolonToken       = SymbolId(299)
	ColonToken           = SymbolId(300)
	ExclaimToken         = SymbolId(301)
	DollarLbracketToken  = SymbolId(302)
	DotdotdotToken       = SymbolId(303)
	TildeTildeToken      = SymbolId(304)
	AssignToken          = SymbolId(305)
	AddAssignToken       = SymbolId(306)
	SubAssignToken       = SymbolId(307)
	MulAssignToken       = SymbolId(308)
	DivAssignToken       = SymbolId(309)
	ModAssignToken       = SymbolId(310)
	AddOneAssignToken    = SymbolId(311)
	SubOneAssignToken    = SymbolId(312)
	BitNegAssignToken    = SymbolId(313)
	BitAndAssignToken    = SymbolId(314)
	BitOrAssignToken     = SymbolId(315)
	BitXorAssignToken    = SymbolId(316)
	BitLshiftAssignToken = SymbolId(317)
	BitRshiftAssignToken = SymbolId(318)
	NotToken             = SymbolId(319)
	AndToken             = SymbolId(320)
	OrToken              = SymbolId(321)
	AddToken             = SymbolId(322)
	SubToken             = SymbolId(323)
	MulToken             = SymbolId(324)
	DivToken             = SymbolId(325)
	ModToken             = SymbolId(326)
	BitNegToken          = SymbolId(327)
	BitAndToken          = SymbolId(328)
	BitXorToken          = SymbolId(329)
	BitOrToken           = SymbolId(330)
	BitLshiftToken       = SymbolId(331)
	BitRshiftToken       = SymbolId(332)
	EqualToken           = SymbolId(333)
	NotEqualToken        = SymbolId(334)
	LessToken            = SymbolId(335)
	LessOrEqualToken     = SymbolId(336)
	GreaterToken         = SymbolId(337)
	GreaterOrEqualToken  = SymbolId(338)
	LexErrorToken        = SymbolId(339)
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
	// 47:10: source -> ...
	ToSource(OptionalDefinitions_ *GenericSymbol) (*GenericSymbol, error)

	// 50:2: optional_definitions -> NEWLINES: ...
	NewlinesToOptionalDefinitions(Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 51:2: optional_definitions -> definitions: ...
	DefinitionsToOptionalDefinitions(OptionalNewlines_ *GenericSymbol, Definitions_ *GenericSymbol, OptionalNewlines_2 *GenericSymbol) (*GenericSymbol, error)

	// 52:2: optional_definitions -> nil: ...
	NilToOptionalDefinitions() (*GenericSymbol, error)

	// 55:2: optional_newlines -> NEWLINES: ...
	NewlinesToOptionalNewlines(Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 56:2: optional_newlines -> nil: ...
	NilToOptionalNewlines() (*GenericSymbol, error)

	// 59:2: definitions -> nil: ...
	NilToDefinitions(Definition_ *GenericSymbol) (*GenericSymbol, error)

	// 60:2: definitions -> add: ...
	AddToDefinitions(Definitions_ *GenericSymbol, Newlines_ *GenericSymbol, Definition_ *GenericSymbol) (*GenericSymbol, error)

	// 63:2: definition -> package_def: ...
	PackageDefToDefinition(PackageDef_ *GenericSymbol) (*GenericSymbol, error)

	// 64:2: definition -> type_def: ...
	TypeDefToDefinition(TypeDef_ *GenericSymbol) (*GenericSymbol, error)

	// 65:2: definition -> named_func_def: ...
	NamedFuncDefToDefinition(NamedFuncDef_ *GenericSymbol) (*GenericSymbol, error)

	// 66:2: definition -> unsafe_statement: ...
	UnsafeStatementToDefinition(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 73:2: var_pattern -> IDENTIFIER: ...
	IdentifierToVarPattern(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 74:2: var_pattern -> tuple_pattern: ...
	TuplePatternToVarPattern(TuplePattern_ *GenericSymbol) (*GenericSymbol, error)

	// 76:17: tuple_pattern -> ...
	ToTuplePattern(Lparen_ *GenericSymbol, FieldVarPatterns_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 79:2: field_var_patterns -> field_var_pattern: ...
	FieldVarPatternToFieldVarPatterns(FieldVarPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 80:2: field_var_patterns -> add: ...
	AddToFieldVarPatterns(FieldVarPatterns_ *GenericSymbol, Comma_ *GenericSymbol, FieldVarPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 83:2: field_var_pattern -> positional: ...
	PositionalToFieldVarPattern(VarPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 84:2: field_var_pattern -> named: ...
	NamedToFieldVarPattern(Identifier_ *GenericSymbol, Assign_ *GenericSymbol, VarPattern_ *GenericSymbol) (*GenericSymbol, error)

	// 85:2: field_var_pattern -> DOTDOTDOT: ...
	DotdotdotToFieldVarPattern(Dotdotdot_ *GenericSymbol) (*GenericSymbol, error)

	// 88:2: optional_value_type -> value_type: ...
	ValueTypeToOptionalValueType(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 89:2: optional_value_type -> nil: ...
	NilToOptionalValueType() (*GenericSymbol, error)

	// 92:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 93:2: expression -> switch_expr: ...
	SwitchExprToExpression(OptionalLabelDecl_ *GenericSymbol, SwitchExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 94:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ *GenericSymbol, LoopExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 95:2: expression -> sequence_expr: ...
	SequenceExprToExpression(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 98:2: expression -> var_pattern: ...
	VarPatternToExpression(Var_ *GenericSymbol, VarPattern_ *GenericSymbol, OptionalValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 101:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 102:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (*GenericSymbol, error)

	// 111:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 112:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, BlockBody_2 *GenericSymbol) (*GenericSymbol, error)

	// 113:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ *GenericSymbol, SequenceExpr_ *GenericSymbol, BlockBody_ *GenericSymbol, Else_ *GenericSymbol, IfExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 137:15: switch_expr -> ...
	ToSwitchExpr(Switch_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Lbrace_ *GenericSymbol, CaseBranches_ *GenericSymbol, OptionalDefaultBranch_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 140:2: case_branches -> case_branch: ...
	CaseBranchToCaseBranches(CaseBranch_ *GenericSymbol) (*GenericSymbol, error)

	// 141:2: case_branches -> add: ...
	AddToCaseBranches(CaseBranches_ *GenericSymbol, CaseBranch_ *GenericSymbol) (*GenericSymbol, error)

	// 143:15: case_branch -> ...
	ToCaseBranch(Case_ *GenericSymbol, Patterns_ *GenericSymbol, Colon_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 146:2: patterns -> pattern: ...
	PatternToPatterns(Pattern_ *GenericSymbol) (*GenericSymbol, error)

	// 147:2: patterns -> multiple: ...
	MultipleToPatterns(Patterns_ *GenericSymbol, Comma_ *GenericSymbol, Pattern_ *GenericSymbol) (*GenericSymbol, error)

	// 150:2: optional_default_branch -> default_branch: ...
	DefaultBranchToOptionalDefaultBranch(DefaultBranch_ *GenericSymbol) (*GenericSymbol, error)

	// 151:2: optional_default_branch -> nil: ...
	NilToOptionalDefaultBranch() (*GenericSymbol, error)

	// 153:18: default_branch -> ...
	ToDefaultBranch(Default_ *GenericSymbol, Colon_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 167:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 168:2: loop_expr -> do_while: ...
	DoWhileToLoopExpr(Do_ *GenericSymbol, BlockBody_ *GenericSymbol, For_ *GenericSymbol, SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 169:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 170:2: loop_expr -> iterator: ...
	IteratorToLoopExpr(For_ *GenericSymbol, Pattern_ *GenericSymbol, In_ *GenericSymbol, SequenceExpr_ *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 171:2: loop_expr -> for: ...
	ForToLoopExpr(For_ *GenericSymbol, Semicolon_ *GenericSymbol, OptionalSequenceExpr_ *GenericSymbol, Semicolon_2 *GenericSymbol, OptionalSequenceExpr_2 *GenericSymbol, Do_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 174:2: optional_sequence_expr -> sequence_expr: ...
	SequenceExprToOptionalSequenceExpr(SequenceExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 175:2: optional_sequence_expr -> nil: ...
	NilToOptionalSequenceExpr() (*GenericSymbol, error)

	// 181:17: sequence_expr -> ...
	ToSequenceExpr(OrExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 183:14: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 185:14: block_body -> ...
	ToBlockBody(Lbrace_ *GenericSymbol, Statements_ *GenericSymbol, Rbrace_ *GenericSymbol) (*GenericSymbol, error)

	// 188:2: statements -> empty_list: ...
	EmptyListToStatements() (*GenericSymbol, error)

	// 189:2: statements -> add: ...
	AddToStatements(Statements_ *GenericSymbol, Statement_ *GenericSymbol) (*GenericSymbol, error)

	// 192:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 193:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 196:2: statement_body -> unsafe_statement: ...
	UnsafeStatementToStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 198:2: statement_body -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToStatementBody(Expressions_ *GenericSymbol) (*GenericSymbol, error)

	// 204:2: statement_body -> async: ...
	AsyncToStatementBody(Async_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 218:2: statement_body -> defer: ...
	DeferToStatementBody(Defer_ *GenericSymbol, CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 222:2: statement_body -> jump_statement: ...
	JumpStatementToStatementBody(JumpStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 229:2: statement_body -> assign_statement: ...
	AssignStatementToStatementBody(Pattern_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 233:2: statement_body -> unary_op_assign_statement: ...
	UnaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, UnaryOpAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 234:2: statement_body -> binary_op_assign_statement: ...
	BinaryOpAssignStatementToStatementBody(AccessExpr_ *GenericSymbol, BinaryOpAssign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 245:2: pattern -> expression: ...
	ExpressionToPattern(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 251:2: pattern -> enum_match_pattern: ...
	EnumMatchPatternToPattern(Dot_ *GenericSymbol, Identifier_ *GenericSymbol, ImplicitStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 252:2: pattern -> enum_var_pattern: ...
	EnumVarPatternToPattern(Var_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol, TuplePattern_ *GenericSymbol) (*GenericSymbol, error)

	// 255:2: unary_op_assign -> ADD_ONE_ASSIGN: ...
	AddOneAssignToUnaryOpAssign(AddOneAssign_ *GenericSymbol) (*GenericSymbol, error)

	// 256:2: unary_op_assign -> SUB_ONE_ASSIGN: ...
	SubOneAssignToUnaryOpAssign(SubOneAssign_ *GenericSymbol) (*GenericSymbol, error)

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

	// 277:20: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ *GenericSymbol, Less_ *GenericSymbol, Identifier_ *GenericSymbol, Greater_ *GenericSymbol, StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 284:2: jump_statement -> ...
	ToJumpStatement(JumpType_ *GenericSymbol, OptionalJumpLabel_ *GenericSymbol, OptionalExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 287:2: jump_type -> RETURN: ...
	ReturnToJumpType(Return_ *GenericSymbol) (*GenericSymbol, error)

	// 288:2: jump_type -> BREAK: ...
	BreakToJumpType(Break_ *GenericSymbol) (*GenericSymbol, error)

	// 289:2: jump_type -> CONTINUE: ...
	ContinueToJumpType(Continue_ *GenericSymbol) (*GenericSymbol, error)

	// 292:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ *GenericSymbol) (*GenericSymbol, error)

	// 293:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (*GenericSymbol, error)

	// 296:2: expressions -> expression: ...
	ExpressionToExpressions(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 297:2: expressions -> add: ...
	AddToExpressions(Expressions_ *GenericSymbol, Comma_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 300:2: optional_expressions -> expressions: ...
	ExpressionsToOptionalExpressions(Expressions_ *GenericSymbol) (*GenericSymbol, error)

	// 301:2: optional_expressions -> nil: ...
	NilToOptionalExpressions() (*GenericSymbol, error)

	// 307:13: call_expr -> ...
	ToCallExpr(AccessExpr_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalArguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 310:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ *GenericSymbol, OptionalGenericArguments_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 311:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (*GenericSymbol, error)

	// 314:2: optional_generic_arguments -> generic_arguments: ...
	GenericArgumentsToOptionalGenericArguments(GenericArguments_ *GenericSymbol) (*GenericSymbol, error)

	// 315:2: optional_generic_arguments -> nil: ...
	NilToOptionalGenericArguments() (*GenericSymbol, error)

	// 319:2: generic_arguments -> value_type: ...
	ValueTypeToGenericArguments(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 320:2: generic_arguments -> add: ...
	AddToGenericArguments(GenericArguments_ *GenericSymbol, Comma_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 323:2: optional_arguments -> arguments: ...
	ArgumentsToOptionalArguments(Arguments_ *GenericSymbol) (*GenericSymbol, error)

	// 324:2: optional_arguments -> nil: ...
	NilToOptionalArguments() (*GenericSymbol, error)

	// 327:2: arguments -> argument: ...
	ArgumentToArguments(Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 328:2: arguments -> add: ...
	AddToArguments(Arguments_ *GenericSymbol, Comma_ *GenericSymbol, Argument_ *GenericSymbol) (*GenericSymbol, error)

	// 331:2: argument -> positional: ...
	PositionalToArgument(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 332:2: argument -> named: ...
	NamedToArgument(Identifier_ *GenericSymbol, Assign_ *GenericSymbol, Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 333:2: argument -> colon_expressions: ...
	ColonExpressionsToArgument(ColonExpressions_ *GenericSymbol) (*GenericSymbol, error)

	// 336:2: argument -> DOTDOTDOT: ...
	DotdotdotToArgument(Dotdotdot_ *GenericSymbol) (*GenericSymbol, error)

	// 339:2: colon_expressions -> pair: ...
	PairToColonExpressions(OptionalExpression_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_2 *GenericSymbol) (*GenericSymbol, error)

	// 340:2: colon_expressions -> add: ...
	AddToColonExpressions(ColonExpressions_ *GenericSymbol, Colon_ *GenericSymbol, OptionalExpression_ *GenericSymbol) (*GenericSymbol, error)

	// 343:2: optional_expression -> expression: ...
	ExpressionToOptionalExpression(Expression_ *GenericSymbol) (*GenericSymbol, error)

	// 344:2: optional_expression -> nil: ...
	NilToOptionalExpression() (*GenericSymbol, error)

	// 354:2: atom_expr -> literal: ...
	LiteralToAtomExpr(Literal_ *GenericSymbol) (*GenericSymbol, error)

	// 355:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 356:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(BlockExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 357:2: atom_expr -> anonymous_func_expr: ...
	AnonymousFuncExprToAtomExpr(AnonymousFuncExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 358:2: atom_expr -> initialize_expr: ...
	InitializeExprToAtomExpr(InitializableType_ *GenericSymbol, Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 359:2: atom_expr -> implicit_struct_expr: ...
	ImplicitStructExprToAtomExpr(ImplicitStructExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 360:2: atom_expr -> LEX_ERROR: ...
	LexErrorToAtomExpr(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 363:2: literal -> TRUE: ...
	TrueToLiteral(True_ *GenericSymbol) (*GenericSymbol, error)

	// 364:2: literal -> FALSE: ...
	FalseToLiteral(False_ *GenericSymbol) (*GenericSymbol, error)

	// 365:2: literal -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteral(IntegerLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 366:2: literal -> FLOAT_LITERAL: ...
	FloatLiteralToLiteral(FloatLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 367:2: literal -> RUNE_LITERAL: ...
	RuneLiteralToLiteral(RuneLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 368:2: literal -> STRING_LITERAL: ...
	StringLiteralToLiteral(StringLiteral_ *GenericSymbol) (*GenericSymbol, error)

	// 370:24: implicit_struct_expr -> ...
	ToImplicitStructExpr(Lparen_ *GenericSymbol, Arguments_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 373:2: access_expr -> atom_expr: ...
	AtomExprToAccessExpr(AtomExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 374:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 375:2: access_expr -> call_expr: ...
	CallExprToAccessExpr(CallExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 376:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ *GenericSymbol, Lbracket_ *GenericSymbol, Argument_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 379:2: postfix_unary_expr -> access_expr: ...
	AccessExprToPostfixUnaryExpr(AccessExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 380:2: postfix_unary_expr -> question: ...
	QuestionToPostfixUnaryExpr(AccessExpr_ *GenericSymbol, Question_ *GenericSymbol) (*GenericSymbol, error)

	// 383:2: prefix_unary_op -> NOT: ...
	NotToPrefixUnaryOp(Not_ *GenericSymbol) (*GenericSymbol, error)

	// 384:2: prefix_unary_op -> BIT_NEG: ...
	BitNegToPrefixUnaryOp(BitNeg_ *GenericSymbol) (*GenericSymbol, error)

	// 385:2: prefix_unary_op -> SUB: ...
	SubToPrefixUnaryOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 388:2: prefix_unary_op -> MUL: ...
	MulToPrefixUnaryOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 391:2: prefix_unary_op -> BIT_AND: ...
	BitAndToPrefixUnaryOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 394:2: prefix_unary_expr -> postfix_unary_expr: ...
	PostfixUnaryExprToPrefixUnaryExpr(PostfixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 395:2: prefix_unary_expr -> prefix_op: ...
	PrefixOpToPrefixUnaryExpr(PrefixUnaryOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 398:2: mul_op -> MUL: ...
	MulToMulOp(Mul_ *GenericSymbol) (*GenericSymbol, error)

	// 399:2: mul_op -> DIV: ...
	DivToMulOp(Div_ *GenericSymbol) (*GenericSymbol, error)

	// 400:2: mul_op -> MOD: ...
	ModToMulOp(Mod_ *GenericSymbol) (*GenericSymbol, error)

	// 401:2: mul_op -> BIT_AND: ...
	BitAndToMulOp(BitAnd_ *GenericSymbol) (*GenericSymbol, error)

	// 402:2: mul_op -> BIT_LSHIFT: ...
	BitLshiftToMulOp(BitLshift_ *GenericSymbol) (*GenericSymbol, error)

	// 403:2: mul_op -> BIT_RSHIFT: ...
	BitRshiftToMulOp(BitRshift_ *GenericSymbol) (*GenericSymbol, error)

	// 406:2: mul_expr -> prefix_unary_expr: ...
	PrefixUnaryExprToMulExpr(PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 407:2: mul_expr -> op: ...
	OpToMulExpr(MulExpr_ *GenericSymbol, MulOp_ *GenericSymbol, PrefixUnaryExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 410:2: add_op -> ADD: ...
	AddToAddOp(Add_ *GenericSymbol) (*GenericSymbol, error)

	// 411:2: add_op -> SUB: ...
	SubToAddOp(Sub_ *GenericSymbol) (*GenericSymbol, error)

	// 412:2: add_op -> BIT_OR: ...
	BitOrToAddOp(BitOr_ *GenericSymbol) (*GenericSymbol, error)

	// 413:2: add_op -> BIT_XOR: ...
	BitXorToAddOp(BitXor_ *GenericSymbol) (*GenericSymbol, error)

	// 416:2: add_expr -> mul_expr: ...
	MulExprToAddExpr(MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 417:2: add_expr -> op: ...
	OpToAddExpr(AddExpr_ *GenericSymbol, AddOp_ *GenericSymbol, MulExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 420:2: cmp_op -> EQUAL: ...
	EqualToCmpOp(Equal_ *GenericSymbol) (*GenericSymbol, error)

	// 421:2: cmp_op -> NOT_EQUAL: ...
	NotEqualToCmpOp(NotEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 422:2: cmp_op -> LESS: ...
	LessToCmpOp(Less_ *GenericSymbol) (*GenericSymbol, error)

	// 423:2: cmp_op -> LESS_OR_EQUAL: ...
	LessOrEqualToCmpOp(LessOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 424:2: cmp_op -> GREATER: ...
	GreaterToCmpOp(Greater_ *GenericSymbol) (*GenericSymbol, error)

	// 425:2: cmp_op -> GREATER_OR_EQUAL: ...
	GreaterOrEqualToCmpOp(GreaterOrEqual_ *GenericSymbol) (*GenericSymbol, error)

	// 428:2: cmp_expr -> add_expr: ...
	AddExprToCmpExpr(AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 429:2: cmp_expr -> op: ...
	OpToCmpExpr(CmpExpr_ *GenericSymbol, CmpOp_ *GenericSymbol, AddExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 432:2: and_expr -> cmp_expr: ...
	CmpExprToAndExpr(CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 433:2: and_expr -> op: ...
	OpToAndExpr(AndExpr_ *GenericSymbol, And_ *GenericSymbol, CmpExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 436:2: or_expr -> and_expr: ...
	AndExprToOrExpr(AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 437:2: or_expr -> op: ...
	OpToOrExpr(OrExpr_ *GenericSymbol, Or_ *GenericSymbol, AndExpr_ *GenericSymbol) (*GenericSymbol, error)

	// 446:2: initializable_type -> explicit_struct_def: ...
	ExplicitStructDefToInitializableType(ExplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 447:2: initializable_type -> slice: ...
	SliceToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 448:2: initializable_type -> array: ...
	ArrayToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Comma_ *GenericSymbol, IntegerLiteral_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 449:2: initializable_type -> map: ...
	MapToInitializableType(Lbracket_ *GenericSymbol, ValueType_ *GenericSymbol, Colon_ *GenericSymbol, ValueType_2 *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 452:2: atom_type -> initializable_type: ...
	InitializableTypeToAtomType(InitializableType_ *GenericSymbol) (*GenericSymbol, error)

	// 453:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 454:2: atom_type -> extern_named: ...
	ExternNamedToAtomType(Identifier_ *GenericSymbol, Dot_ *GenericSymbol, Identifier_2 *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 455:2: atom_type -> inferred: ...
	InferredToAtomType(Dot_ *GenericSymbol, OptionalGenericBinding_ *GenericSymbol) (*GenericSymbol, error)

	// 456:2: atom_type -> implicit_struct_def: ...
	ImplicitStructDefToAtomType(ImplicitStructDef_ *GenericSymbol) (*GenericSymbol, error)

	// 457:2: atom_type -> explicit_enum_def: ...
	ExplicitEnumDefToAtomType(ExplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 458:2: atom_type -> implicit_enum_def: ...
	ImplicitEnumDefToAtomType(ImplicitEnumDef_ *GenericSymbol) (*GenericSymbol, error)

	// 459:2: atom_type -> trait_def: ...
	TraitDefToAtomType(TraitDef_ *GenericSymbol) (*GenericSymbol, error)

	// 460:2: atom_type -> func_type: ...
	FuncTypeToAtomType(FuncType_ *GenericSymbol) (*GenericSymbol, error)

	// 461:2: atom_type -> LEX_ERROR: ...
	LexErrorToAtomType(LexError_ *GenericSymbol) (*GenericSymbol, error)

	// 467:2: returnable_type -> atom_type: ...
	AtomTypeToReturnableType(AtomType_ *GenericSymbol) (*GenericSymbol, error)

	// 468:2: returnable_type -> optional: ...
	OptionalToReturnableType(Question_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 469:2: returnable_type -> result: ...
	ResultToReturnableType(Exclaim_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 470:2: returnable_type -> reference: ...
	ReferenceToReturnableType(BitAnd_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 471:2: returnable_type -> public_methods_trait: ...
	PublicMethodsTraitToReturnableType(BitNeg_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 472:2: returnable_type -> public_trait: ...
	PublicTraitToReturnableType(TildeTilde_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 477:2: value_type -> returnable_type: ...
	ReturnableTypeToValueType(ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 478:2: value_type -> trait_intersect: ...
	TraitIntersectToValueType(ValueType_ *GenericSymbol, Mul_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 479:2: value_type -> trait_union: ...
	TraitUnionToValueType(ValueType_ *GenericSymbol, Add_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 480:2: value_type -> trait_difference: ...
	TraitDifferenceToValueType(ValueType_ *GenericSymbol, Sub_ *GenericSymbol, ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 483:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 484:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, ValueType_ *GenericSymbol, Implements_ *GenericSymbol, ValueType_2 *GenericSymbol) (*GenericSymbol, error)

	// 485:2: type_def -> alias: ...
	AliasToTypeDef(Type_ *GenericSymbol, Identifier_ *GenericSymbol, Equal_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 493:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 494:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 497:2: generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToGenericParameterDefs(GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 498:2: generic_parameter_defs -> add: ...
	AddToGenericParameterDefs(GenericParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, GenericParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 501:2: optional_generic_parameter_defs -> generic_parameter_defs: ...
	GenericParameterDefsToOptionalGenericParameterDefs(GenericParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 502:2: optional_generic_parameter_defs -> nil: ...
	NilToOptionalGenericParameterDefs() (*GenericSymbol, error)

	// 505:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ *GenericSymbol, OptionalGenericParameterDefs_ *GenericSymbol, Rbracket_ *GenericSymbol) (*GenericSymbol, error)

	// 506:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (*GenericSymbol, error)

	// 513:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 514:2: field_def -> implicit: ...
	ImplicitToFieldDef(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 515:2: field_def -> unsafe_statement: ...
	UnsafeStatementToFieldDef(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 518:2: implicit_field_defs -> field_def: ...
	FieldDefToImplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 519:2: implicit_field_defs -> add: ...
	AddToImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 522:2: optional_implicit_field_defs -> implicit_field_defs: ...
	ImplicitFieldDefsToOptionalImplicitFieldDefs(ImplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 523:2: optional_implicit_field_defs -> nil: ...
	NilToOptionalImplicitFieldDefs() (*GenericSymbol, error)

	// 525:23: implicit_struct_def -> ...
	ToImplicitStructDef(Lparen_ *GenericSymbol, OptionalImplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 528:2: explicit_field_defs -> field_def: ...
	FieldDefToExplicitFieldDefs(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 529:2: explicit_field_defs -> implicit: ...
	ImplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Newlines_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 530:2: explicit_field_defs -> explicit: ...
	ExplicitToExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol, Comma_ *GenericSymbol, FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 533:2: optional_explicit_field_defs -> explicit_field_defs: ...
	ExplicitFieldDefsToOptionalExplicitFieldDefs(ExplicitFieldDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 534:2: optional_explicit_field_defs -> nil: ...
	NilToOptionalExplicitFieldDefs() (*GenericSymbol, error)

	// 536:23: explicit_struct_def -> ...
	ToExplicitStructDef(Struct_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalExplicitFieldDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 546:2: enum_value_def -> field_def: ...
	FieldDefToEnumValueDef(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 547:2: enum_value_def -> default: ...
	DefaultToEnumValueDef(FieldDef_ *GenericSymbol, Assign_ *GenericSymbol, Default_ *GenericSymbol) (*GenericSymbol, error)

	// 551:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 552:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 554:21: implicit_enum_def -> ...
	ToImplicitEnumDef(Lparen_ *GenericSymbol, ImplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 557:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 558:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(EnumValueDef_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_2 *GenericSymbol) (*GenericSymbol, error)

	// 559:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Or_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 560:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ *GenericSymbol, Newlines_ *GenericSymbol, EnumValueDef_ *GenericSymbol) (*GenericSymbol, error)

	// 562:21: explicit_enum_def -> ...
	ToExplicitEnumDef(Enum_ *GenericSymbol, Lparen_ *GenericSymbol, ExplicitEnumValueDefs_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 569:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ *GenericSymbol) (*GenericSymbol, error)

	// 570:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ *GenericSymbol) (*GenericSymbol, error)

	// 573:2: trait_properties -> trait_property: ...
	TraitPropertyToTraitProperties(TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 574:2: trait_properties -> implicit: ...
	ImplicitToTraitProperties(TraitProperties_ *GenericSymbol, Newlines_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 575:2: trait_properties -> explicit: ...
	ExplicitToTraitProperties(TraitProperties_ *GenericSymbol, Comma_ *GenericSymbol, TraitProperty_ *GenericSymbol) (*GenericSymbol, error)

	// 578:2: optional_trait_properties -> trait_properties: ...
	TraitPropertiesToOptionalTraitProperties(TraitProperties_ *GenericSymbol) (*GenericSymbol, error)

	// 579:2: optional_trait_properties -> nil: ...
	NilToOptionalTraitProperties() (*GenericSymbol, error)

	// 581:13: trait_def -> ...
	ToTraitDef(Trait_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalTraitProperties_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 589:2: return_type -> returnable_type: ...
	ReturnableTypeToReturnType(ReturnableType_ *GenericSymbol) (*GenericSymbol, error)

	// 590:2: return_type -> nil: ...
	NilToReturnType() (*GenericSymbol, error)

	// 593:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 594:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 595:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 596:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 599:2: parameter_decls -> parameter_decl: ...
	ParameterDeclToParameterDecls(ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 600:2: parameter_decls -> add: ...
	AddToParameterDecls(ParameterDecls_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDecl_ *GenericSymbol) (*GenericSymbol, error)

	// 603:2: optional_parameter_decls -> parameter_decls: ...
	ParameterDeclsToOptionalParameterDecls(ParameterDecls_ *GenericSymbol) (*GenericSymbol, error)

	// 604:2: optional_parameter_decls -> nil: ...
	NilToOptionalParameterDecls() (*GenericSymbol, error)

	// 606:13: func_type -> ...
	ToFuncType(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 617:20: method_signature -> ...
	ToMethodSignature(Func_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDecls_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol) (*GenericSymbol, error)

	// 620:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 621:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ *GenericSymbol, Dotdotdot_ *GenericSymbol, ValueType_ *GenericSymbol) (*GenericSymbol, error)

	// 624:2: parameter_defs -> parameter_def: ...
	ParameterDefToParameterDefs(ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 625:2: parameter_defs -> add: ...
	AddToParameterDefs(ParameterDefs_ *GenericSymbol, Comma_ *GenericSymbol, ParameterDef_ *GenericSymbol) (*GenericSymbol, error)

	// 628:2: optional_parameter_defs -> parameter_defs: ...
	ParameterDefsToOptionalParameterDefs(ParameterDefs_ *GenericSymbol) (*GenericSymbol, error)

	// 629:2: optional_parameter_defs -> nil: ...
	NilToOptionalParameterDefs() (*GenericSymbol, error)

	// 632:2: optional_receiver -> receiver: ...
	ReceiverToOptionalReceiver(Lparen_ *GenericSymbol, ParameterDef_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 633:2: optional_receiver -> nil: ...
	NilToOptionalReceiver() (*GenericSymbol, error)

	// 636:2: named_func_def -> ...
	ToNamedFuncDef(Func_ *GenericSymbol, OptionalReceiver_ *GenericSymbol, Identifier_ *GenericSymbol, OptionalGenericParameters_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 639:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ *GenericSymbol, Lparen_ *GenericSymbol, OptionalParameterDefs_ *GenericSymbol, Rparen_ *GenericSymbol, ReturnType_ *GenericSymbol, BlockBody_ *GenericSymbol) (*GenericSymbol, error)

	// 646:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol) (*GenericSymbol, error)

	// 647:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ *GenericSymbol, Identifier_ *GenericSymbol, Lparen_ *GenericSymbol, PackageStatements_ *GenericSymbol, Rparen_ *GenericSymbol) (*GenericSymbol, error)

	// 649:26: package_statement_body -> ...
	ToPackageStatementBody(UnsafeStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 652:2: package_statement -> implicit: ...
	ImplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Newlines_ *GenericSymbol) (*GenericSymbol, error)

	// 653:2: package_statement -> explicit: ...
	ExplicitToPackageStatement(PackageStatementBody_ *GenericSymbol, Semicolon_ *GenericSymbol) (*GenericSymbol, error)

	// 656:2: package_statements -> empty_list: ...
	EmptyListToPackageStatements() (*GenericSymbol, error)

	// 657:2: package_statements -> add: ...
	AddToPackageStatements(PackageStatements_ *GenericSymbol, PackageStatement_ *GenericSymbol) (*GenericSymbol, error)

	// 661:2: lex_internal_tokens -> SPACES: ...
	SpacesToLexInternalTokens(Spaces_ *GenericSymbol) (*GenericSymbol, error)

	// 662:2: lex_internal_tokens -> COMMENT: ...
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
	case VarToken:
		return "VAR"
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
	case VarPatternType:
		return "var_pattern"
	case TuplePatternType:
		return "tuple_pattern"
	case FieldVarPatternsType:
		return "field_var_patterns"
	case FieldVarPatternType:
		return "field_var_pattern"
	case OptionalValueTypeType:
		return "optional_value_type"
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
	case PatternsType:
		return "patterns"
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
	case PatternType:
		return "pattern"
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
	case ImplicitStructExprType:
		return "implicit_struct_expr"
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

	SourceType                       = SymbolId(340)
	OptionalDefinitionsType          = SymbolId(341)
	OptionalNewlinesType             = SymbolId(342)
	DefinitionsType                  = SymbolId(343)
	DefinitionType                   = SymbolId(344)
	VarPatternType                   = SymbolId(345)
	TuplePatternType                 = SymbolId(346)
	FieldVarPatternsType             = SymbolId(347)
	FieldVarPatternType              = SymbolId(348)
	OptionalValueTypeType            = SymbolId(349)
	ExpressionType                   = SymbolId(350)
	OptionalLabelDeclType            = SymbolId(351)
	IfExprType                       = SymbolId(352)
	SwitchExprType                   = SymbolId(353)
	CaseBranchesType                 = SymbolId(354)
	CaseBranchType                   = SymbolId(355)
	PatternsType                     = SymbolId(356)
	OptionalDefaultBranchType        = SymbolId(357)
	DefaultBranchType                = SymbolId(358)
	LoopExprType                     = SymbolId(359)
	OptionalSequenceExprType         = SymbolId(360)
	SequenceExprType                 = SymbolId(361)
	BlockExprType                    = SymbolId(362)
	BlockBodyType                    = SymbolId(363)
	StatementsType                   = SymbolId(364)
	StatementType                    = SymbolId(365)
	StatementBodyType                = SymbolId(366)
	PatternType                      = SymbolId(367)
	UnaryOpAssignType                = SymbolId(368)
	BinaryOpAssignType               = SymbolId(369)
	UnsafeStatementType              = SymbolId(370)
	JumpStatementType                = SymbolId(371)
	JumpTypeType                     = SymbolId(372)
	OptionalJumpLabelType            = SymbolId(373)
	ExpressionsType                  = SymbolId(374)
	OptionalExpressionsType          = SymbolId(375)
	CallExprType                     = SymbolId(376)
	OptionalGenericBindingType       = SymbolId(377)
	OptionalGenericArgumentsType     = SymbolId(378)
	GenericArgumentsType             = SymbolId(379)
	OptionalArgumentsType            = SymbolId(380)
	ArgumentsType                    = SymbolId(381)
	ArgumentType                     = SymbolId(382)
	ColonExpressionsType             = SymbolId(383)
	OptionalExpressionType           = SymbolId(384)
	AtomExprType                     = SymbolId(385)
	LiteralType                      = SymbolId(386)
	ImplicitStructExprType           = SymbolId(387)
	AccessExprType                   = SymbolId(388)
	PostfixUnaryExprType             = SymbolId(389)
	PrefixUnaryOpType                = SymbolId(390)
	PrefixUnaryExprType              = SymbolId(391)
	MulOpType                        = SymbolId(392)
	MulExprType                      = SymbolId(393)
	AddOpType                        = SymbolId(394)
	AddExprType                      = SymbolId(395)
	CmpOpType                        = SymbolId(396)
	CmpExprType                      = SymbolId(397)
	AndExprType                      = SymbolId(398)
	OrExprType                       = SymbolId(399)
	InitializableTypeType            = SymbolId(400)
	AtomTypeType                     = SymbolId(401)
	ReturnableTypeType               = SymbolId(402)
	ValueTypeType                    = SymbolId(403)
	TypeDefType                      = SymbolId(404)
	GenericParameterDefType          = SymbolId(405)
	GenericParameterDefsType         = SymbolId(406)
	OptionalGenericParameterDefsType = SymbolId(407)
	OptionalGenericParametersType    = SymbolId(408)
	FieldDefType                     = SymbolId(409)
	ImplicitFieldDefsType            = SymbolId(410)
	OptionalImplicitFieldDefsType    = SymbolId(411)
	ImplicitStructDefType            = SymbolId(412)
	ExplicitFieldDefsType            = SymbolId(413)
	OptionalExplicitFieldDefsType    = SymbolId(414)
	ExplicitStructDefType            = SymbolId(415)
	EnumValueDefType                 = SymbolId(416)
	ImplicitEnumValueDefsType        = SymbolId(417)
	ImplicitEnumDefType              = SymbolId(418)
	ExplicitEnumValueDefsType        = SymbolId(419)
	ExplicitEnumDefType              = SymbolId(420)
	TraitPropertyType                = SymbolId(421)
	TraitPropertiesType              = SymbolId(422)
	OptionalTraitPropertiesType      = SymbolId(423)
	TraitDefType                     = SymbolId(424)
	ReturnTypeType                   = SymbolId(425)
	ParameterDeclType                = SymbolId(426)
	ParameterDeclsType               = SymbolId(427)
	OptionalParameterDeclsType       = SymbolId(428)
	FuncTypeType                     = SymbolId(429)
	MethodSignatureType              = SymbolId(430)
	ParameterDefType                 = SymbolId(431)
	ParameterDefsType                = SymbolId(432)
	OptionalParameterDefsType        = SymbolId(433)
	OptionalReceiverType             = SymbolId(434)
	NamedFuncDefType                 = SymbolId(435)
	AnonymousFuncExprType            = SymbolId(436)
	PackageDefType                   = SymbolId(437)
	PackageStatementBodyType         = SymbolId(438)
	PackageStatementType             = SymbolId(439)
	PackageStatementsType            = SymbolId(440)
	LexInternalTokensType            = SymbolId(441)
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
	_ReduceIdentifierToVarPattern                             = _ReduceType(13)
	_ReduceTuplePatternToVarPattern                           = _ReduceType(14)
	_ReduceToTuplePattern                                     = _ReduceType(15)
	_ReduceFieldVarPatternToFieldVarPatterns                  = _ReduceType(16)
	_ReduceAddToFieldVarPatterns                              = _ReduceType(17)
	_ReducePositionalToFieldVarPattern                        = _ReduceType(18)
	_ReduceNamedToFieldVarPattern                             = _ReduceType(19)
	_ReduceDotdotdotToFieldVarPattern                         = _ReduceType(20)
	_ReduceValueTypeToOptionalValueType                       = _ReduceType(21)
	_ReduceNilToOptionalValueType                             = _ReduceType(22)
	_ReduceIfExprToExpression                                 = _ReduceType(23)
	_ReduceSwitchExprToExpression                             = _ReduceType(24)
	_ReduceLoopExprToExpression                               = _ReduceType(25)
	_ReduceSequenceExprToExpression                           = _ReduceType(26)
	_ReduceVarPatternToExpression                             = _ReduceType(27)
	_ReduceLabelDeclToOptionalLabelDecl                       = _ReduceType(28)
	_ReduceUnlabelledToOptionalLabelDecl                      = _ReduceType(29)
	_ReduceNoElseToIfExpr                                     = _ReduceType(30)
	_ReduceIfElseToIfExpr                                     = _ReduceType(31)
	_ReduceMultiIfElseToIfExpr                                = _ReduceType(32)
	_ReduceToSwitchExpr                                       = _ReduceType(33)
	_ReduceCaseBranchToCaseBranches                           = _ReduceType(34)
	_ReduceAddToCaseBranches                                  = _ReduceType(35)
	_ReduceToCaseBranch                                       = _ReduceType(36)
	_ReducePatternToPatterns                                  = _ReduceType(37)
	_ReduceMultipleToPatterns                                 = _ReduceType(38)
	_ReduceDefaultBranchToOptionalDefaultBranch               = _ReduceType(39)
	_ReduceNilToOptionalDefaultBranch                         = _ReduceType(40)
	_ReduceToDefaultBranch                                    = _ReduceType(41)
	_ReduceInfiniteToLoopExpr                                 = _ReduceType(42)
	_ReduceDoWhileToLoopExpr                                  = _ReduceType(43)
	_ReduceWhileToLoopExpr                                    = _ReduceType(44)
	_ReduceIteratorToLoopExpr                                 = _ReduceType(45)
	_ReduceForToLoopExpr                                      = _ReduceType(46)
	_ReduceSequenceExprToOptionalSequenceExpr                 = _ReduceType(47)
	_ReduceNilToOptionalSequenceExpr                          = _ReduceType(48)
	_ReduceToSequenceExpr                                     = _ReduceType(49)
	_ReduceToBlockExpr                                        = _ReduceType(50)
	_ReduceToBlockBody                                        = _ReduceType(51)
	_ReduceEmptyListToStatements                              = _ReduceType(52)
	_ReduceAddToStatements                                    = _ReduceType(53)
	_ReduceImplicitToStatement                                = _ReduceType(54)
	_ReduceExplicitToStatement                                = _ReduceType(55)
	_ReduceUnsafeStatementToStatementBody                     = _ReduceType(56)
	_ReduceExpressionOrImplicitStructToStatementBody          = _ReduceType(57)
	_ReduceAsyncToStatementBody                               = _ReduceType(58)
	_ReduceDeferToStatementBody                               = _ReduceType(59)
	_ReduceJumpStatementToStatementBody                       = _ReduceType(60)
	_ReduceAssignStatementToStatementBody                     = _ReduceType(61)
	_ReduceUnaryOpAssignStatementToStatementBody              = _ReduceType(62)
	_ReduceBinaryOpAssignStatementToStatementBody             = _ReduceType(63)
	_ReduceExpressionToPattern                                = _ReduceType(64)
	_ReduceEnumMatchPatternToPattern                          = _ReduceType(65)
	_ReduceEnumVarPatternToPattern                            = _ReduceType(66)
	_ReduceAddOneAssignToUnaryOpAssign                        = _ReduceType(67)
	_ReduceSubOneAssignToUnaryOpAssign                        = _ReduceType(68)
	_ReduceAddAssignToBinaryOpAssign                          = _ReduceType(69)
	_ReduceSubAssignToBinaryOpAssign                          = _ReduceType(70)
	_ReduceMulAssignToBinaryOpAssign                          = _ReduceType(71)
	_ReduceDivAssignToBinaryOpAssign                          = _ReduceType(72)
	_ReduceModAssignToBinaryOpAssign                          = _ReduceType(73)
	_ReduceBitNegAssignToBinaryOpAssign                       = _ReduceType(74)
	_ReduceBitAndAssignToBinaryOpAssign                       = _ReduceType(75)
	_ReduceBitOrAssignToBinaryOpAssign                        = _ReduceType(76)
	_ReduceBitXorAssignToBinaryOpAssign                       = _ReduceType(77)
	_ReduceBitLshiftAssignToBinaryOpAssign                    = _ReduceType(78)
	_ReduceBitRshiftAssignToBinaryOpAssign                    = _ReduceType(79)
	_ReduceToUnsafeStatement                                  = _ReduceType(80)
	_ReduceToJumpStatement                                    = _ReduceType(81)
	_ReduceReturnToJumpType                                   = _ReduceType(82)
	_ReduceBreakToJumpType                                    = _ReduceType(83)
	_ReduceContinueToJumpType                                 = _ReduceType(84)
	_ReduceJumpLabelToOptionalJumpLabel                       = _ReduceType(85)
	_ReduceUnlabelledToOptionalJumpLabel                      = _ReduceType(86)
	_ReduceExpressionToExpressions                            = _ReduceType(87)
	_ReduceAddToExpressions                                   = _ReduceType(88)
	_ReduceExpressionsToOptionalExpressions                   = _ReduceType(89)
	_ReduceNilToOptionalExpressions                           = _ReduceType(90)
	_ReduceToCallExpr                                         = _ReduceType(91)
	_ReduceBindingToOptionalGenericBinding                    = _ReduceType(92)
	_ReduceNilToOptionalGenericBinding                        = _ReduceType(93)
	_ReduceGenericArgumentsToOptionalGenericArguments         = _ReduceType(94)
	_ReduceNilToOptionalGenericArguments                      = _ReduceType(95)
	_ReduceValueTypeToGenericArguments                        = _ReduceType(96)
	_ReduceAddToGenericArguments                              = _ReduceType(97)
	_ReduceArgumentsToOptionalArguments                       = _ReduceType(98)
	_ReduceNilToOptionalArguments                             = _ReduceType(99)
	_ReduceArgumentToArguments                                = _ReduceType(100)
	_ReduceAddToArguments                                     = _ReduceType(101)
	_ReducePositionalToArgument                               = _ReduceType(102)
	_ReduceNamedToArgument                                    = _ReduceType(103)
	_ReduceColonExpressionsToArgument                         = _ReduceType(104)
	_ReduceDotdotdotToArgument                                = _ReduceType(105)
	_ReducePairToColonExpressions                             = _ReduceType(106)
	_ReduceAddToColonExpressions                              = _ReduceType(107)
	_ReduceExpressionToOptionalExpression                     = _ReduceType(108)
	_ReduceNilToOptionalExpression                            = _ReduceType(109)
	_ReduceLiteralToAtomExpr                                  = _ReduceType(110)
	_ReduceIdentifierToAtomExpr                               = _ReduceType(111)
	_ReduceBlockExprToAtomExpr                                = _ReduceType(112)
	_ReduceAnonymousFuncExprToAtomExpr                        = _ReduceType(113)
	_ReduceInitializeExprToAtomExpr                           = _ReduceType(114)
	_ReduceImplicitStructExprToAtomExpr                       = _ReduceType(115)
	_ReduceLexErrorToAtomExpr                                 = _ReduceType(116)
	_ReduceTrueToLiteral                                      = _ReduceType(117)
	_ReduceFalseToLiteral                                     = _ReduceType(118)
	_ReduceIntegerLiteralToLiteral                            = _ReduceType(119)
	_ReduceFloatLiteralToLiteral                              = _ReduceType(120)
	_ReduceRuneLiteralToLiteral                               = _ReduceType(121)
	_ReduceStringLiteralToLiteral                             = _ReduceType(122)
	_ReduceToImplicitStructExpr                               = _ReduceType(123)
	_ReduceAtomExprToAccessExpr                               = _ReduceType(124)
	_ReduceAccessToAccessExpr                                 = _ReduceType(125)
	_ReduceCallExprToAccessExpr                               = _ReduceType(126)
	_ReduceIndexToAccessExpr                                  = _ReduceType(127)
	_ReduceAccessExprToPostfixUnaryExpr                       = _ReduceType(128)
	_ReduceQuestionToPostfixUnaryExpr                         = _ReduceType(129)
	_ReduceNotToPrefixUnaryOp                                 = _ReduceType(130)
	_ReduceBitNegToPrefixUnaryOp                              = _ReduceType(131)
	_ReduceSubToPrefixUnaryOp                                 = _ReduceType(132)
	_ReduceMulToPrefixUnaryOp                                 = _ReduceType(133)
	_ReduceBitAndToPrefixUnaryOp                              = _ReduceType(134)
	_ReducePostfixUnaryExprToPrefixUnaryExpr                  = _ReduceType(135)
	_ReducePrefixOpToPrefixUnaryExpr                          = _ReduceType(136)
	_ReduceMulToMulOp                                         = _ReduceType(137)
	_ReduceDivToMulOp                                         = _ReduceType(138)
	_ReduceModToMulOp                                         = _ReduceType(139)
	_ReduceBitAndToMulOp                                      = _ReduceType(140)
	_ReduceBitLshiftToMulOp                                   = _ReduceType(141)
	_ReduceBitRshiftToMulOp                                   = _ReduceType(142)
	_ReducePrefixUnaryExprToMulExpr                           = _ReduceType(143)
	_ReduceOpToMulExpr                                        = _ReduceType(144)
	_ReduceAddToAddOp                                         = _ReduceType(145)
	_ReduceSubToAddOp                                         = _ReduceType(146)
	_ReduceBitOrToAddOp                                       = _ReduceType(147)
	_ReduceBitXorToAddOp                                      = _ReduceType(148)
	_ReduceMulExprToAddExpr                                   = _ReduceType(149)
	_ReduceOpToAddExpr                                        = _ReduceType(150)
	_ReduceEqualToCmpOp                                       = _ReduceType(151)
	_ReduceNotEqualToCmpOp                                    = _ReduceType(152)
	_ReduceLessToCmpOp                                        = _ReduceType(153)
	_ReduceLessOrEqualToCmpOp                                 = _ReduceType(154)
	_ReduceGreaterToCmpOp                                     = _ReduceType(155)
	_ReduceGreaterOrEqualToCmpOp                              = _ReduceType(156)
	_ReduceAddExprToCmpExpr                                   = _ReduceType(157)
	_ReduceOpToCmpExpr                                        = _ReduceType(158)
	_ReduceCmpExprToAndExpr                                   = _ReduceType(159)
	_ReduceOpToAndExpr                                        = _ReduceType(160)
	_ReduceAndExprToOrExpr                                    = _ReduceType(161)
	_ReduceOpToOrExpr                                         = _ReduceType(162)
	_ReduceExplicitStructDefToInitializableType               = _ReduceType(163)
	_ReduceSliceToInitializableType                           = _ReduceType(164)
	_ReduceArrayToInitializableType                           = _ReduceType(165)
	_ReduceMapToInitializableType                             = _ReduceType(166)
	_ReduceInitializableTypeToAtomType                        = _ReduceType(167)
	_ReduceNamedToAtomType                                    = _ReduceType(168)
	_ReduceExternNamedToAtomType                              = _ReduceType(169)
	_ReduceInferredToAtomType                                 = _ReduceType(170)
	_ReduceImplicitStructDefToAtomType                        = _ReduceType(171)
	_ReduceExplicitEnumDefToAtomType                          = _ReduceType(172)
	_ReduceImplicitEnumDefToAtomType                          = _ReduceType(173)
	_ReduceTraitDefToAtomType                                 = _ReduceType(174)
	_ReduceFuncTypeToAtomType                                 = _ReduceType(175)
	_ReduceLexErrorToAtomType                                 = _ReduceType(176)
	_ReduceAtomTypeToReturnableType                           = _ReduceType(177)
	_ReduceOptionalToReturnableType                           = _ReduceType(178)
	_ReduceResultToReturnableType                             = _ReduceType(179)
	_ReduceReferenceToReturnableType                          = _ReduceType(180)
	_ReducePublicMethodsTraitToReturnableType                 = _ReduceType(181)
	_ReducePublicTraitToReturnableType                        = _ReduceType(182)
	_ReduceReturnableTypeToValueType                          = _ReduceType(183)
	_ReduceTraitIntersectToValueType                          = _ReduceType(184)
	_ReduceTraitUnionToValueType                              = _ReduceType(185)
	_ReduceTraitDifferenceToValueType                         = _ReduceType(186)
	_ReduceDefinitionToTypeDef                                = _ReduceType(187)
	_ReduceConstrainedDefToTypeDef                            = _ReduceType(188)
	_ReduceAliasToTypeDef                                     = _ReduceType(189)
	_ReduceUnconstrainedToGenericParameterDef                 = _ReduceType(190)
	_ReduceConstrainedToGenericParameterDef                   = _ReduceType(191)
	_ReduceGenericParameterDefToGenericParameterDefs          = _ReduceType(192)
	_ReduceAddToGenericParameterDefs                          = _ReduceType(193)
	_ReduceGenericParameterDefsToOptionalGenericParameterDefs = _ReduceType(194)
	_ReduceNilToOptionalGenericParameterDefs                  = _ReduceType(195)
	_ReduceGenericToOptionalGenericParameters                 = _ReduceType(196)
	_ReduceNilToOptionalGenericParameters                     = _ReduceType(197)
	_ReduceExplicitToFieldDef                                 = _ReduceType(198)
	_ReduceImplicitToFieldDef                                 = _ReduceType(199)
	_ReduceUnsafeStatementToFieldDef                          = _ReduceType(200)
	_ReduceFieldDefToImplicitFieldDefs                        = _ReduceType(201)
	_ReduceAddToImplicitFieldDefs                             = _ReduceType(202)
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefs       = _ReduceType(203)
	_ReduceNilToOptionalImplicitFieldDefs                     = _ReduceType(204)
	_ReduceToImplicitStructDef                                = _ReduceType(205)
	_ReduceFieldDefToExplicitFieldDefs                        = _ReduceType(206)
	_ReduceImplicitToExplicitFieldDefs                        = _ReduceType(207)
	_ReduceExplicitToExplicitFieldDefs                        = _ReduceType(208)
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefs       = _ReduceType(209)
	_ReduceNilToOptionalExplicitFieldDefs                     = _ReduceType(210)
	_ReduceToExplicitStructDef                                = _ReduceType(211)
	_ReduceFieldDefToEnumValueDef                             = _ReduceType(212)
	_ReduceDefaultToEnumValueDef                              = _ReduceType(213)
	_ReducePairToImplicitEnumValueDefs                        = _ReduceType(214)
	_ReduceAddToImplicitEnumValueDefs                         = _ReduceType(215)
	_ReduceToImplicitEnumDef                                  = _ReduceType(216)
	_ReduceExplicitPairToExplicitEnumValueDefs                = _ReduceType(217)
	_ReduceImplicitPairToExplicitEnumValueDefs                = _ReduceType(218)
	_ReduceExplicitAddToExplicitEnumValueDefs                 = _ReduceType(219)
	_ReduceImplicitAddToExplicitEnumValueDefs                 = _ReduceType(220)
	_ReduceToExplicitEnumDef                                  = _ReduceType(221)
	_ReduceFieldDefToTraitProperty                            = _ReduceType(222)
	_ReduceMethodSignatureToTraitProperty                     = _ReduceType(223)
	_ReduceTraitPropertyToTraitProperties                     = _ReduceType(224)
	_ReduceImplicitToTraitProperties                          = _ReduceType(225)
	_ReduceExplicitToTraitProperties                          = _ReduceType(226)
	_ReduceTraitPropertiesToOptionalTraitProperties           = _ReduceType(227)
	_ReduceNilToOptionalTraitProperties                       = _ReduceType(228)
	_ReduceToTraitDef                                         = _ReduceType(229)
	_ReduceReturnableTypeToReturnType                         = _ReduceType(230)
	_ReduceNilToReturnType                                    = _ReduceType(231)
	_ReduceArgToParameterDecl                                 = _ReduceType(232)
	_ReduceVarargToParameterDecl                              = _ReduceType(233)
	_ReduceUnamedToParameterDecl                              = _ReduceType(234)
	_ReduceUnnamedVarargToParameterDecl                       = _ReduceType(235)
	_ReduceParameterDeclToParameterDecls                      = _ReduceType(236)
	_ReduceAddToParameterDecls                                = _ReduceType(237)
	_ReduceParameterDeclsToOptionalParameterDecls             = _ReduceType(238)
	_ReduceNilToOptionalParameterDecls                        = _ReduceType(239)
	_ReduceToFuncType                                         = _ReduceType(240)
	_ReduceToMethodSignature                                  = _ReduceType(241)
	_ReduceArgToParameterDef                                  = _ReduceType(242)
	_ReduceVarargToParameterDef                               = _ReduceType(243)
	_ReduceParameterDefToParameterDefs                        = _ReduceType(244)
	_ReduceAddToParameterDefs                                 = _ReduceType(245)
	_ReduceParameterDefsToOptionalParameterDefs               = _ReduceType(246)
	_ReduceNilToOptionalParameterDefs                         = _ReduceType(247)
	_ReduceReceiverToOptionalReceiver                         = _ReduceType(248)
	_ReduceNilToOptionalReceiver                              = _ReduceType(249)
	_ReduceToNamedFuncDef                                     = _ReduceType(250)
	_ReduceToAnonymousFuncExpr                                = _ReduceType(251)
	_ReduceNoSpecToPackageDef                                 = _ReduceType(252)
	_ReduceWithSpecToPackageDef                               = _ReduceType(253)
	_ReduceToPackageStatementBody                             = _ReduceType(254)
	_ReduceImplicitToPackageStatement                         = _ReduceType(255)
	_ReduceExplicitToPackageStatement                         = _ReduceType(256)
	_ReduceEmptyListToPackageStatements                       = _ReduceType(257)
	_ReduceAddToPackageStatements                             = _ReduceType(258)
	_ReduceSpacesToLexInternalTokens                          = _ReduceType(259)
	_ReduceCommentToLexInternalTokens                         = _ReduceType(260)
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
	case _ReduceIdentifierToVarPattern:
		return "IdentifierToVarPattern"
	case _ReduceTuplePatternToVarPattern:
		return "TuplePatternToVarPattern"
	case _ReduceToTuplePattern:
		return "ToTuplePattern"
	case _ReduceFieldVarPatternToFieldVarPatterns:
		return "FieldVarPatternToFieldVarPatterns"
	case _ReduceAddToFieldVarPatterns:
		return "AddToFieldVarPatterns"
	case _ReducePositionalToFieldVarPattern:
		return "PositionalToFieldVarPattern"
	case _ReduceNamedToFieldVarPattern:
		return "NamedToFieldVarPattern"
	case _ReduceDotdotdotToFieldVarPattern:
		return "DotdotdotToFieldVarPattern"
	case _ReduceValueTypeToOptionalValueType:
		return "ValueTypeToOptionalValueType"
	case _ReduceNilToOptionalValueType:
		return "NilToOptionalValueType"
	case _ReduceIfExprToExpression:
		return "IfExprToExpression"
	case _ReduceSwitchExprToExpression:
		return "SwitchExprToExpression"
	case _ReduceLoopExprToExpression:
		return "LoopExprToExpression"
	case _ReduceSequenceExprToExpression:
		return "SequenceExprToExpression"
	case _ReduceVarPatternToExpression:
		return "VarPatternToExpression"
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
	case _ReducePatternToPatterns:
		return "PatternToPatterns"
	case _ReduceMultipleToPatterns:
		return "MultipleToPatterns"
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
	case _ReduceAssignStatementToStatementBody:
		return "AssignStatementToStatementBody"
	case _ReduceUnaryOpAssignStatementToStatementBody:
		return "UnaryOpAssignStatementToStatementBody"
	case _ReduceBinaryOpAssignStatementToStatementBody:
		return "BinaryOpAssignStatementToStatementBody"
	case _ReduceExpressionToPattern:
		return "ExpressionToPattern"
	case _ReduceEnumMatchPatternToPattern:
		return "EnumMatchPatternToPattern"
	case _ReduceEnumVarPatternToPattern:
		return "EnumVarPatternToPattern"
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
	case _ReduceDotdotdotToArgument:
		return "DotdotdotToArgument"
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
	case _ReduceImplicitStructExprToAtomExpr:
		return "ImplicitStructExprToAtomExpr"
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
	case _ReduceToImplicitStructExpr:
		return "ToImplicitStructExpr"
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
	_State394 = _StateId(394)
	_State395 = _StateId(395)
	_State396 = _StateId(396)
	_State397 = _StateId(397)
	_State398 = _StateId(398)
	_State399 = _StateId(399)
	_State400 = _StateId(400)
	_State401 = _StateId(401)
	_State402 = _StateId(402)
	_State403 = _StateId(403)
	_State404 = _StateId(404)
	_State405 = _StateId(405)
	_State406 = _StateId(406)
	_State407 = _StateId(407)
	_State408 = _StateId(408)
	_State409 = _StateId(409)
	_State410 = _StateId(410)
	_State411 = _StateId(411)
	_State412 = _StateId(412)
	_State413 = _StateId(413)
	_State414 = _StateId(414)
	_State415 = _StateId(415)
	_State416 = _StateId(416)
	_State417 = _StateId(417)
	_State418 = _StateId(418)
	_State419 = _StateId(419)
	_State420 = _StateId(420)
	_State421 = _StateId(421)
	_State422 = _StateId(422)
	_State423 = _StateId(423)
	_State424 = _StateId(424)
	_State425 = _StateId(425)
	_State426 = _StateId(426)
	_State427 = _StateId(427)
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
	case _EndMarker, SpacesToken, NewlinesToken, CommentToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, PackageToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, DotdotdotToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, NotToken, AndToken, OrToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, LexErrorToken:
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
	case _ReduceIdentifierToVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarPatternType
		symbol.Generic_, err = reducer.IdentifierToVarPattern(args[0].Generic_)
	case _ReduceTuplePatternToVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarPatternType
		symbol.Generic_, err = reducer.TuplePatternToVarPattern(args[0].Generic_)
	case _ReduceToTuplePattern:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TuplePatternType
		symbol.Generic_, err = reducer.ToTuplePattern(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceFieldVarPatternToFieldVarPatterns:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldVarPatternsType
		symbol.Generic_, err = reducer.FieldVarPatternToFieldVarPatterns(args[0].Generic_)
	case _ReduceAddToFieldVarPatterns:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = FieldVarPatternsType
		symbol.Generic_, err = reducer.AddToFieldVarPatterns(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReducePositionalToFieldVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldVarPatternType
		symbol.Generic_, err = reducer.PositionalToFieldVarPattern(args[0].Generic_)
	case _ReduceNamedToFieldVarPattern:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = FieldVarPatternType
		symbol.Generic_, err = reducer.NamedToFieldVarPattern(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceDotdotdotToFieldVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldVarPatternType
		symbol.Generic_, err = reducer.DotdotdotToFieldVarPattern(args[0].Generic_)
	case _ReduceValueTypeToOptionalValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalValueTypeType
		symbol.Generic_, err = reducer.ValueTypeToOptionalValueType(args[0].Generic_)
	case _ReduceNilToOptionalValueType:
		symbol.SymbolId_ = OptionalValueTypeType
		symbol.Generic_, err = reducer.NilToOptionalValueType()
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
	case _ReduceVarPatternToExpression:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExpressionType
		symbol.Generic_, err = reducer.VarPatternToExpression(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	case _ReducePatternToPatterns:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PatternsType
		symbol.Generic_, err = reducer.PatternToPatterns(args[0].Generic_)
	case _ReduceMultipleToPatterns:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = PatternsType
		symbol.Generic_, err = reducer.MultipleToPatterns(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = LoopExprType
		symbol.Generic_, err = reducer.IteratorToLoopExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_, args[4].Generic_, args[5].Generic_)
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
	case _ReduceAssignStatementToStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = StatementBodyType
		symbol.Generic_, err = reducer.AssignStatementToStatementBody(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	case _ReduceExpressionToPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PatternType
		symbol.Generic_, err = reducer.ExpressionToPattern(args[0].Generic_)
	case _ReduceEnumMatchPatternToPattern:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = PatternType
		symbol.Generic_, err = reducer.EnumMatchPatternToPattern(args[0].Generic_, args[1].Generic_, args[2].Generic_)
	case _ReduceEnumVarPatternToPattern:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = PatternType
		symbol.Generic_, err = reducer.EnumVarPatternToPattern(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
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
	case _ReduceDotdotdotToArgument:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentType
		symbol.Generic_, err = reducer.DotdotdotToArgument(args[0].Generic_)
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
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.InitializeExprToAtomExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_, args[3].Generic_)
	case _ReduceImplicitStructExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Generic_, err = reducer.ImplicitStructExprToAtomExpr(args[0].Generic_)
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
	case _ReduceToImplicitStructExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitStructExprType
		symbol.Generic_, err = reducer.ToImplicitStructExpr(args[0].Generic_, args[1].Generic_, args[2].Generic_)
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
	_GotoState394Action                                             = &_Action{_ShiftAction, _State394, 0}
	_GotoState395Action                                             = &_Action{_ShiftAction, _State395, 0}
	_GotoState396Action                                             = &_Action{_ShiftAction, _State396, 0}
	_GotoState397Action                                             = &_Action{_ShiftAction, _State397, 0}
	_GotoState398Action                                             = &_Action{_ShiftAction, _State398, 0}
	_GotoState399Action                                             = &_Action{_ShiftAction, _State399, 0}
	_GotoState400Action                                             = &_Action{_ShiftAction, _State400, 0}
	_GotoState401Action                                             = &_Action{_ShiftAction, _State401, 0}
	_GotoState402Action                                             = &_Action{_ShiftAction, _State402, 0}
	_GotoState403Action                                             = &_Action{_ShiftAction, _State403, 0}
	_GotoState404Action                                             = &_Action{_ShiftAction, _State404, 0}
	_GotoState405Action                                             = &_Action{_ShiftAction, _State405, 0}
	_GotoState406Action                                             = &_Action{_ShiftAction, _State406, 0}
	_GotoState407Action                                             = &_Action{_ShiftAction, _State407, 0}
	_GotoState408Action                                             = &_Action{_ShiftAction, _State408, 0}
	_GotoState409Action                                             = &_Action{_ShiftAction, _State409, 0}
	_GotoState410Action                                             = &_Action{_ShiftAction, _State410, 0}
	_GotoState411Action                                             = &_Action{_ShiftAction, _State411, 0}
	_GotoState412Action                                             = &_Action{_ShiftAction, _State412, 0}
	_GotoState413Action                                             = &_Action{_ShiftAction, _State413, 0}
	_GotoState414Action                                             = &_Action{_ShiftAction, _State414, 0}
	_GotoState415Action                                             = &_Action{_ShiftAction, _State415, 0}
	_GotoState416Action                                             = &_Action{_ShiftAction, _State416, 0}
	_GotoState417Action                                             = &_Action{_ShiftAction, _State417, 0}
	_GotoState418Action                                             = &_Action{_ShiftAction, _State418, 0}
	_GotoState419Action                                             = &_Action{_ShiftAction, _State419, 0}
	_GotoState420Action                                             = &_Action{_ShiftAction, _State420, 0}
	_GotoState421Action                                             = &_Action{_ShiftAction, _State421, 0}
	_GotoState422Action                                             = &_Action{_ShiftAction, _State422, 0}
	_GotoState423Action                                             = &_Action{_ShiftAction, _State423, 0}
	_GotoState424Action                                             = &_Action{_ShiftAction, _State424, 0}
	_GotoState425Action                                             = &_Action{_ShiftAction, _State425, 0}
	_GotoState426Action                                             = &_Action{_ShiftAction, _State426, 0}
	_GotoState427Action                                             = &_Action{_ShiftAction, _State427, 0}
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
	_ReduceIdentifierToVarPatternAction                             = &_Action{_ReduceAction, 0, _ReduceIdentifierToVarPattern}
	_ReduceTuplePatternToVarPatternAction                           = &_Action{_ReduceAction, 0, _ReduceTuplePatternToVarPattern}
	_ReduceToTuplePatternAction                                     = &_Action{_ReduceAction, 0, _ReduceToTuplePattern}
	_ReduceFieldVarPatternToFieldVarPatternsAction                  = &_Action{_ReduceAction, 0, _ReduceFieldVarPatternToFieldVarPatterns}
	_ReduceAddToFieldVarPatternsAction                              = &_Action{_ReduceAction, 0, _ReduceAddToFieldVarPatterns}
	_ReducePositionalToFieldVarPatternAction                        = &_Action{_ReduceAction, 0, _ReducePositionalToFieldVarPattern}
	_ReduceNamedToFieldVarPatternAction                             = &_Action{_ReduceAction, 0, _ReduceNamedToFieldVarPattern}
	_ReduceDotdotdotToFieldVarPatternAction                         = &_Action{_ReduceAction, 0, _ReduceDotdotdotToFieldVarPattern}
	_ReduceValueTypeToOptionalValueTypeAction                       = &_Action{_ReduceAction, 0, _ReduceValueTypeToOptionalValueType}
	_ReduceNilToOptionalValueTypeAction                             = &_Action{_ReduceAction, 0, _ReduceNilToOptionalValueType}
	_ReduceIfExprToExpressionAction                                 = &_Action{_ReduceAction, 0, _ReduceIfExprToExpression}
	_ReduceSwitchExprToExpressionAction                             = &_Action{_ReduceAction, 0, _ReduceSwitchExprToExpression}
	_ReduceLoopExprToExpressionAction                               = &_Action{_ReduceAction, 0, _ReduceLoopExprToExpression}
	_ReduceSequenceExprToExpressionAction                           = &_Action{_ReduceAction, 0, _ReduceSequenceExprToExpression}
	_ReduceVarPatternToExpressionAction                             = &_Action{_ReduceAction, 0, _ReduceVarPatternToExpression}
	_ReduceLabelDeclToOptionalLabelDeclAction                       = &_Action{_ReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}
	_ReduceUnlabelledToOptionalLabelDeclAction                      = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}
	_ReduceNoElseToIfExprAction                                     = &_Action{_ReduceAction, 0, _ReduceNoElseToIfExpr}
	_ReduceIfElseToIfExprAction                                     = &_Action{_ReduceAction, 0, _ReduceIfElseToIfExpr}
	_ReduceMultiIfElseToIfExprAction                                = &_Action{_ReduceAction, 0, _ReduceMultiIfElseToIfExpr}
	_ReduceToSwitchExprAction                                       = &_Action{_ReduceAction, 0, _ReduceToSwitchExpr}
	_ReduceCaseBranchToCaseBranchesAction                           = &_Action{_ReduceAction, 0, _ReduceCaseBranchToCaseBranches}
	_ReduceAddToCaseBranchesAction                                  = &_Action{_ReduceAction, 0, _ReduceAddToCaseBranches}
	_ReduceToCaseBranchAction                                       = &_Action{_ReduceAction, 0, _ReduceToCaseBranch}
	_ReducePatternToPatternsAction                                  = &_Action{_ReduceAction, 0, _ReducePatternToPatterns}
	_ReduceMultipleToPatternsAction                                 = &_Action{_ReduceAction, 0, _ReduceMultipleToPatterns}
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
	_ReduceAssignStatementToStatementBodyAction                     = &_Action{_ReduceAction, 0, _ReduceAssignStatementToStatementBody}
	_ReduceUnaryOpAssignStatementToStatementBodyAction              = &_Action{_ReduceAction, 0, _ReduceUnaryOpAssignStatementToStatementBody}
	_ReduceBinaryOpAssignStatementToStatementBodyAction             = &_Action{_ReduceAction, 0, _ReduceBinaryOpAssignStatementToStatementBody}
	_ReduceExpressionToPatternAction                                = &_Action{_ReduceAction, 0, _ReduceExpressionToPattern}
	_ReduceEnumMatchPatternToPatternAction                          = &_Action{_ReduceAction, 0, _ReduceEnumMatchPatternToPattern}
	_ReduceEnumVarPatternToPatternAction                            = &_Action{_ReduceAction, 0, _ReduceEnumVarPatternToPattern}
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
	_ReduceDotdotdotToArgumentAction                                = &_Action{_ReduceAction, 0, _ReduceDotdotdotToArgument}
	_ReducePairToColonExpressionsAction                             = &_Action{_ReduceAction, 0, _ReducePairToColonExpressions}
	_ReduceAddToColonExpressionsAction                              = &_Action{_ReduceAction, 0, _ReduceAddToColonExpressions}
	_ReduceExpressionToOptionalExpressionAction                     = &_Action{_ReduceAction, 0, _ReduceExpressionToOptionalExpression}
	_ReduceNilToOptionalExpressionAction                            = &_Action{_ReduceAction, 0, _ReduceNilToOptionalExpression}
	_ReduceLiteralToAtomExprAction                                  = &_Action{_ReduceAction, 0, _ReduceLiteralToAtomExpr}
	_ReduceIdentifierToAtomExprAction                               = &_Action{_ReduceAction, 0, _ReduceIdentifierToAtomExpr}
	_ReduceBlockExprToAtomExprAction                                = &_Action{_ReduceAction, 0, _ReduceBlockExprToAtomExpr}
	_ReduceAnonymousFuncExprToAtomExprAction                        = &_Action{_ReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}
	_ReduceInitializeExprToAtomExprAction                           = &_Action{_ReduceAction, 0, _ReduceInitializeExprToAtomExpr}
	_ReduceImplicitStructExprToAtomExprAction                       = &_Action{_ReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}
	_ReduceLexErrorToAtomExprAction                                 = &_Action{_ReduceAction, 0, _ReduceLexErrorToAtomExpr}
	_ReduceTrueToLiteralAction                                      = &_Action{_ReduceAction, 0, _ReduceTrueToLiteral}
	_ReduceFalseToLiteralAction                                     = &_Action{_ReduceAction, 0, _ReduceFalseToLiteral}
	_ReduceIntegerLiteralToLiteralAction                            = &_Action{_ReduceAction, 0, _ReduceIntegerLiteralToLiteral}
	_ReduceFloatLiteralToLiteralAction                              = &_Action{_ReduceAction, 0, _ReduceFloatLiteralToLiteral}
	_ReduceRuneLiteralToLiteralAction                               = &_Action{_ReduceAction, 0, _ReduceRuneLiteralToLiteral}
	_ReduceStringLiteralToLiteralAction                             = &_Action{_ReduceAction, 0, _ReduceStringLiteralToLiteral}
	_ReduceToImplicitStructExprAction                               = &_Action{_ReduceAction, 0, _ReduceToImplicitStructExpr}
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
	{_State5, VarToken}:                           _GotoState37Action,
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
	{_State5, OptionalLabelDeclType}:              _GotoState51Action,
	{_State5, SequenceExprType}:                   _GotoState56Action,
	{_State5, BlockExprType}:                      _GotoState43Action,
	{_State5, CallExprType}:                       _GotoState44Action,
	{_State5, AtomExprType}:                       _GotoState42Action,
	{_State5, LiteralType}:                        _GotoState49Action,
	{_State5, ImplicitStructExprType}:             _GotoState47Action,
	{_State5, AccessExprType}:                     _GotoState38Action,
	{_State5, PostfixUnaryExprType}:               _GotoState53Action,
	{_State5, PrefixUnaryOpType}:                  _GotoState55Action,
	{_State5, PrefixUnaryExprType}:                _GotoState54Action,
	{_State5, MulExprType}:                        _GotoState50Action,
	{_State5, AddExprType}:                        _GotoState39Action,
	{_State5, CmpExprType}:                        _GotoState45Action,
	{_State5, AndExprType}:                        _GotoState40Action,
	{_State5, OrExprType}:                         _GotoState52Action,
	{_State5, InitializableTypeType}:              _GotoState48Action,
	{_State5, ExplicitStructDefType}:              _GotoState46Action,
	{_State5, AnonymousFuncExprType}:              _GotoState41Action,
	{_State6, SpacesToken}:                        _GotoState58Action,
	{_State6, CommentToken}:                       _GotoState57Action,
	{_State6, LexInternalTokensType}:              _GotoState12Action,
	{_State15, PackageToken}:                      _GotoState16Action,
	{_State15, UnsafeToken}:                       _GotoState59Action,
	{_State15, TypeToken}:                         _GotoState17Action,
	{_State15, FuncToken}:                         _GotoState18Action,
	{_State15, DefinitionsType}:                   _GotoState61Action,
	{_State15, DefinitionType}:                    _GotoState60Action,
	{_State15, UnsafeStatementType}:               _GotoState65Action,
	{_State15, TypeDefType}:                       _GotoState64Action,
	{_State15, NamedFuncDefType}:                  _GotoState62Action,
	{_State15, PackageDefType}:                    _GotoState63Action,
	{_State16, IdentifierToken}:                   _GotoState66Action,
	{_State17, IdentifierToken}:                   _GotoState67Action,
	{_State18, LparenToken}:                       _GotoState68Action,
	{_State18, OptionalReceiverType}:              _GotoState69Action,
	{_State23, LparenToken}:                       _GotoState70Action,
	{_State27, IdentifierToken}:                   _GotoState77Action,
	{_State27, StructToken}:                       _GotoState34Action,
	{_State27, EnumToken}:                         _GotoState74Action,
	{_State27, TraitToken}:                        _GotoState82Action,
	{_State27, FuncToken}:                         _GotoState76Action,
	{_State27, LparenToken}:                       _GotoState79Action,
	{_State27, LbracketToken}:                     _GotoState27Action,
	{_State27, DotToken}:                          _GotoState73Action,
	{_State27, QuestionToken}:                     _GotoState80Action,
	{_State27, ExclaimToken}:                      _GotoState75Action,
	{_State27, TildeTildeToken}:                   _GotoState81Action,
	{_State27, BitNegToken}:                       _GotoState72Action,
	{_State27, BitAndToken}:                       _GotoState71Action,
	{_State27, LexErrorToken}:                     _GotoState78Action,
	{_State27, InitializableTypeType}:             _GotoState88Action,
	{_State27, AtomTypeType}:                      _GotoState83Action,
	{_State27, ReturnableTypeType}:                _GotoState89Action,
	{_State27, ValueTypeType}:                     _GotoState91Action,
	{_State27, ImplicitStructDefType}:             _GotoState87Action,
	{_State27, ExplicitStructDefType}:             _GotoState46Action,
	{_State27, ImplicitEnumDefType}:               _GotoState86Action,
	{_State27, ExplicitEnumDefType}:               _GotoState84Action,
	{_State27, TraitDefType}:                      _GotoState90Action,
	{_State27, FuncTypeType}:                      _GotoState85Action,
	{_State29, IntegerLiteralToken}:               _GotoState25Action,
	{_State29, FloatLiteralToken}:                 _GotoState22Action,
	{_State29, RuneLiteralToken}:                  _GotoState32Action,
	{_State29, StringLiteralToken}:                _GotoState33Action,
	{_State29, IdentifierToken}:                   _GotoState93Action,
	{_State29, TrueToken}:                         _GotoState36Action,
	{_State29, FalseToken}:                        _GotoState21Action,
	{_State29, StructToken}:                       _GotoState34Action,
	{_State29, FuncToken}:                         _GotoState23Action,
	{_State29, VarToken}:                          _GotoState37Action,
	{_State29, LabelDeclToken}:                    _GotoState26Action,
	{_State29, LparenToken}:                       _GotoState29Action,
	{_State29, LbracketToken}:                     _GotoState27Action,
	{_State29, DotdotdotToken}:                    _GotoState92Action,
	{_State29, NotToken}:                          _GotoState31Action,
	{_State29, SubToken}:                          _GotoState35Action,
	{_State29, MulToken}:                          _GotoState30Action,
	{_State29, BitNegToken}:                       _GotoState20Action,
	{_State29, BitAndToken}:                       _GotoState19Action,
	{_State29, LexErrorToken}:                     _GotoState28Action,
	{_State29, ExpressionType}:                    _GotoState97Action,
	{_State29, OptionalLabelDeclType}:             _GotoState51Action,
	{_State29, SequenceExprType}:                  _GotoState56Action,
	{_State29, BlockExprType}:                     _GotoState43Action,
	{_State29, CallExprType}:                      _GotoState44Action,
	{_State29, ArgumentsType}:                     _GotoState95Action,
	{_State29, ArgumentType}:                      _GotoState94Action,
	{_State29, ColonExpressionsType}:              _GotoState96Action,
	{_State29, OptionalExpressionType}:            _GotoState98Action,
	{_State29, AtomExprType}:                      _GotoState42Action,
	{_State29, LiteralType}:                       _GotoState49Action,
	{_State29, ImplicitStructExprType}:            _GotoState47Action,
	{_State29, AccessExprType}:                    _GotoState38Action,
	{_State29, PostfixUnaryExprType}:              _GotoState53Action,
	{_State29, PrefixUnaryOpType}:                 _GotoState55Action,
	{_State29, PrefixUnaryExprType}:               _GotoState54Action,
	{_State29, MulExprType}:                       _GotoState50Action,
	{_State29, AddExprType}:                       _GotoState39Action,
	{_State29, CmpExprType}:                       _GotoState45Action,
	{_State29, AndExprType}:                       _GotoState40Action,
	{_State29, OrExprType}:                        _GotoState52Action,
	{_State29, InitializableTypeType}:             _GotoState48Action,
	{_State29, ExplicitStructDefType}:             _GotoState46Action,
	{_State29, AnonymousFuncExprType}:             _GotoState41Action,
	{_State34, LparenToken}:                       _GotoState99Action,
	{_State37, IdentifierToken}:                   _GotoState100Action,
	{_State37, LparenToken}:                       _GotoState101Action,
	{_State37, VarPatternType}:                    _GotoState103Action,
	{_State37, TuplePatternType}:                  _GotoState102Action,
	{_State38, LbracketToken}:                     _GotoState106Action,
	{_State38, DotToken}:                          _GotoState105Action,
	{_State38, QuestionToken}:                     _GotoState107Action,
	{_State38, DollarLbracketToken}:               _GotoState104Action,
	{_State38, OptionalGenericBindingType}:        _GotoState108Action,
	{_State39, AddToken}:                          _GotoState109Action,
	{_State39, SubToken}:                          _GotoState112Action,
	{_State39, BitXorToken}:                       _GotoState111Action,
	{_State39, BitOrToken}:                        _GotoState110Action,
	{_State39, AddOpType}:                         _GotoState113Action,
	{_State40, AndToken}:                          _GotoState114Action,
	{_State45, EqualToken}:                        _GotoState115Action,
	{_State45, NotEqualToken}:                     _GotoState120Action,
	{_State45, LessToken}:                         _GotoState118Action,
	{_State45, LessOrEqualToken}:                  _GotoState119Action,
	{_State45, GreaterToken}:                      _GotoState116Action,
	{_State45, GreaterOrEqualToken}:               _GotoState117Action,
	{_State45, CmpOpType}:                         _GotoState121Action,
	{_State48, LparenToken}:                       _GotoState122Action,
	{_State50, MulToken}:                          _GotoState128Action,
	{_State50, DivToken}:                          _GotoState126Action,
	{_State50, ModToken}:                          _GotoState127Action,
	{_State50, BitAndToken}:                       _GotoState123Action,
	{_State50, BitLshiftToken}:                    _GotoState124Action,
	{_State50, BitRshiftToken}:                    _GotoState125Action,
	{_State50, MulOpType}:                         _GotoState129Action,
	{_State51, IfToken}:                           _GotoState132Action,
	{_State51, SwitchToken}:                       _GotoState134Action,
	{_State51, ForToken}:                          _GotoState131Action,
	{_State51, DoToken}:                           _GotoState130Action,
	{_State51, LbraceToken}:                       _GotoState133Action,
	{_State51, IfExprType}:                        _GotoState136Action,
	{_State51, SwitchExprType}:                    _GotoState138Action,
	{_State51, LoopExprType}:                      _GotoState137Action,
	{_State51, BlockBodyType}:                     _GotoState135Action,
	{_State52, OrToken}:                           _GotoState139Action,
	{_State55, IntegerLiteralToken}:               _GotoState25Action,
	{_State55, FloatLiteralToken}:                 _GotoState22Action,
	{_State55, RuneLiteralToken}:                  _GotoState32Action,
	{_State55, StringLiteralToken}:                _GotoState33Action,
	{_State55, IdentifierToken}:                   _GotoState24Action,
	{_State55, TrueToken}:                         _GotoState36Action,
	{_State55, FalseToken}:                        _GotoState21Action,
	{_State55, StructToken}:                       _GotoState34Action,
	{_State55, FuncToken}:                         _GotoState23Action,
	{_State55, LabelDeclToken}:                    _GotoState26Action,
	{_State55, LparenToken}:                       _GotoState29Action,
	{_State55, LbracketToken}:                     _GotoState27Action,
	{_State55, NotToken}:                          _GotoState31Action,
	{_State55, SubToken}:                          _GotoState35Action,
	{_State55, MulToken}:                          _GotoState30Action,
	{_State55, BitNegToken}:                       _GotoState20Action,
	{_State55, BitAndToken}:                       _GotoState19Action,
	{_State55, LexErrorToken}:                     _GotoState28Action,
	{_State55, OptionalLabelDeclType}:             _GotoState140Action,
	{_State55, BlockExprType}:                     _GotoState43Action,
	{_State55, CallExprType}:                      _GotoState44Action,
	{_State55, AtomExprType}:                      _GotoState42Action,
	{_State55, LiteralType}:                       _GotoState49Action,
	{_State55, ImplicitStructExprType}:            _GotoState47Action,
	{_State55, AccessExprType}:                    _GotoState38Action,
	{_State55, PostfixUnaryExprType}:              _GotoState53Action,
	{_State55, PrefixUnaryOpType}:                 _GotoState55Action,
	{_State55, PrefixUnaryExprType}:               _GotoState141Action,
	{_State55, InitializableTypeType}:             _GotoState48Action,
	{_State55, ExplicitStructDefType}:             _GotoState46Action,
	{_State55, AnonymousFuncExprType}:             _GotoState41Action,
	{_State59, LessToken}:                         _GotoState142Action,
	{_State61, NewlinesToken}:                     _GotoState143Action,
	{_State61, OptionalNewlinesType}:              _GotoState144Action,
	{_State66, LparenToken}:                       _GotoState145Action,
	{_State67, DollarLbracketToken}:               _GotoState146Action,
	{_State67, EqualToken}:                        _GotoState147Action,
	{_State67, OptionalGenericParametersType}:     _GotoState148Action,
	{_State68, IdentifierToken}:                   _GotoState149Action,
	{_State68, ParameterDefType}:                  _GotoState150Action,
	{_State69, IdentifierToken}:                   _GotoState151Action,
	{_State70, IdentifierToken}:                   _GotoState149Action,
	{_State70, ParameterDefType}:                  _GotoState153Action,
	{_State70, ParameterDefsType}:                 _GotoState154Action,
	{_State70, OptionalParameterDefsType}:         _GotoState152Action,
	{_State71, IdentifierToken}:                   _GotoState77Action,
	{_State71, StructToken}:                       _GotoState34Action,
	{_State71, EnumToken}:                         _GotoState74Action,
	{_State71, TraitToken}:                        _GotoState82Action,
	{_State71, FuncToken}:                         _GotoState76Action,
	{_State71, LparenToken}:                       _GotoState79Action,
	{_State71, LbracketToken}:                     _GotoState27Action,
	{_State71, DotToken}:                          _GotoState73Action,
	{_State71, QuestionToken}:                     _GotoState80Action,
	{_State71, ExclaimToken}:                      _GotoState75Action,
	{_State71, TildeTildeToken}:                   _GotoState81Action,
	{_State71, BitNegToken}:                       _GotoState72Action,
	{_State71, BitAndToken}:                       _GotoState71Action,
	{_State71, LexErrorToken}:                     _GotoState78Action,
	{_State71, InitializableTypeType}:             _GotoState88Action,
	{_State71, AtomTypeType}:                      _GotoState83Action,
	{_State71, ReturnableTypeType}:                _GotoState155Action,
	{_State71, ImplicitStructDefType}:             _GotoState87Action,
	{_State71, ExplicitStructDefType}:             _GotoState46Action,
	{_State71, ImplicitEnumDefType}:               _GotoState86Action,
	{_State71, ExplicitEnumDefType}:               _GotoState84Action,
	{_State71, TraitDefType}:                      _GotoState90Action,
	{_State71, FuncTypeType}:                      _GotoState85Action,
	{_State72, IdentifierToken}:                   _GotoState77Action,
	{_State72, StructToken}:                       _GotoState34Action,
	{_State72, EnumToken}:                         _GotoState74Action,
	{_State72, TraitToken}:                        _GotoState82Action,
	{_State72, FuncToken}:                         _GotoState76Action,
	{_State72, LparenToken}:                       _GotoState79Action,
	{_State72, LbracketToken}:                     _GotoState27Action,
	{_State72, DotToken}:                          _GotoState73Action,
	{_State72, QuestionToken}:                     _GotoState80Action,
	{_State72, ExclaimToken}:                      _GotoState75Action,
	{_State72, TildeTildeToken}:                   _GotoState81Action,
	{_State72, BitNegToken}:                       _GotoState72Action,
	{_State72, BitAndToken}:                       _GotoState71Action,
	{_State72, LexErrorToken}:                     _GotoState78Action,
	{_State72, InitializableTypeType}:             _GotoState88Action,
	{_State72, AtomTypeType}:                      _GotoState83Action,
	{_State72, ReturnableTypeType}:                _GotoState156Action,
	{_State72, ImplicitStructDefType}:             _GotoState87Action,
	{_State72, ExplicitStructDefType}:             _GotoState46Action,
	{_State72, ImplicitEnumDefType}:               _GotoState86Action,
	{_State72, ExplicitEnumDefType}:               _GotoState84Action,
	{_State72, TraitDefType}:                      _GotoState90Action,
	{_State72, FuncTypeType}:                      _GotoState85Action,
	{_State73, DollarLbracketToken}:               _GotoState104Action,
	{_State73, OptionalGenericBindingType}:        _GotoState157Action,
	{_State74, LparenToken}:                       _GotoState158Action,
	{_State75, IdentifierToken}:                   _GotoState77Action,
	{_State75, StructToken}:                       _GotoState34Action,
	{_State75, EnumToken}:                         _GotoState74Action,
	{_State75, TraitToken}:                        _GotoState82Action,
	{_State75, FuncToken}:                         _GotoState76Action,
	{_State75, LparenToken}:                       _GotoState79Action,
	{_State75, LbracketToken}:                     _GotoState27Action,
	{_State75, DotToken}:                          _GotoState73Action,
	{_State75, QuestionToken}:                     _GotoState80Action,
	{_State75, ExclaimToken}:                      _GotoState75Action,
	{_State75, TildeTildeToken}:                   _GotoState81Action,
	{_State75, BitNegToken}:                       _GotoState72Action,
	{_State75, BitAndToken}:                       _GotoState71Action,
	{_State75, LexErrorToken}:                     _GotoState78Action,
	{_State75, InitializableTypeType}:             _GotoState88Action,
	{_State75, AtomTypeType}:                      _GotoState83Action,
	{_State75, ReturnableTypeType}:                _GotoState159Action,
	{_State75, ImplicitStructDefType}:             _GotoState87Action,
	{_State75, ExplicitStructDefType}:             _GotoState46Action,
	{_State75, ImplicitEnumDefType}:               _GotoState86Action,
	{_State75, ExplicitEnumDefType}:               _GotoState84Action,
	{_State75, TraitDefType}:                      _GotoState90Action,
	{_State75, FuncTypeType}:                      _GotoState85Action,
	{_State76, LparenToken}:                       _GotoState160Action,
	{_State77, DotToken}:                          _GotoState161Action,
	{_State77, DollarLbracketToken}:               _GotoState104Action,
	{_State77, OptionalGenericBindingType}:        _GotoState162Action,
	{_State79, IdentifierToken}:                   _GotoState163Action,
	{_State79, UnsafeToken}:                       _GotoState59Action,
	{_State79, StructToken}:                       _GotoState34Action,
	{_State79, EnumToken}:                         _GotoState74Action,
	{_State79, TraitToken}:                        _GotoState82Action,
	{_State79, FuncToken}:                         _GotoState76Action,
	{_State79, LparenToken}:                       _GotoState79Action,
	{_State79, LbracketToken}:                     _GotoState27Action,
	{_State79, DotToken}:                          _GotoState73Action,
	{_State79, QuestionToken}:                     _GotoState80Action,
	{_State79, ExclaimToken}:                      _GotoState75Action,
	{_State79, TildeTildeToken}:                   _GotoState81Action,
	{_State79, BitNegToken}:                       _GotoState72Action,
	{_State79, BitAndToken}:                       _GotoState71Action,
	{_State79, LexErrorToken}:                     _GotoState78Action,
	{_State79, UnsafeStatementType}:               _GotoState169Action,
	{_State79, InitializableTypeType}:             _GotoState88Action,
	{_State79, AtomTypeType}:                      _GotoState83Action,
	{_State79, ReturnableTypeType}:                _GotoState89Action,
	{_State79, ValueTypeType}:                     _GotoState170Action,
	{_State79, FieldDefType}:                      _GotoState165Action,
	{_State79, ImplicitFieldDefsType}:             _GotoState167Action,
	{_State79, OptionalImplicitFieldDefsType}:     _GotoState168Action,
	{_State79, ImplicitStructDefType}:             _GotoState87Action,
	{_State79, ExplicitStructDefType}:             _GotoState46Action,
	{_State79, EnumValueDefType}:                  _GotoState164Action,
	{_State79, ImplicitEnumValueDefsType}:         _GotoState166Action,
	{_State79, ImplicitEnumDefType}:               _GotoState86Action,
	{_State79, ExplicitEnumDefType}:               _GotoState84Action,
	{_State79, TraitDefType}:                      _GotoState90Action,
	{_State79, FuncTypeType}:                      _GotoState85Action,
	{_State80, IdentifierToken}:                   _GotoState77Action,
	{_State80, StructToken}:                       _GotoState34Action,
	{_State80, EnumToken}:                         _GotoState74Action,
	{_State80, TraitToken}:                        _GotoState82Action,
	{_State80, FuncToken}:                         _GotoState76Action,
	{_State80, LparenToken}:                       _GotoState79Action,
	{_State80, LbracketToken}:                     _GotoState27Action,
	{_State80, DotToken}:                          _GotoState73Action,
	{_State80, QuestionToken}:                     _GotoState80Action,
	{_State80, ExclaimToken}:                      _GotoState75Action,
	{_State80, TildeTildeToken}:                   _GotoState81Action,
	{_State80, BitNegToken}:                       _GotoState72Action,
	{_State80, BitAndToken}:                       _GotoState71Action,
	{_State80, LexErrorToken}:                     _GotoState78Action,
	{_State80, InitializableTypeType}:             _GotoState88Action,
	{_State80, AtomTypeType}:                      _GotoState83Action,
	{_State80, ReturnableTypeType}:                _GotoState171Action,
	{_State80, ImplicitStructDefType}:             _GotoState87Action,
	{_State80, ExplicitStructDefType}:             _GotoState46Action,
	{_State80, ImplicitEnumDefType}:               _GotoState86Action,
	{_State80, ExplicitEnumDefType}:               _GotoState84Action,
	{_State80, TraitDefType}:                      _GotoState90Action,
	{_State80, FuncTypeType}:                      _GotoState85Action,
	{_State81, IdentifierToken}:                   _GotoState77Action,
	{_State81, StructToken}:                       _GotoState34Action,
	{_State81, EnumToken}:                         _GotoState74Action,
	{_State81, TraitToken}:                        _GotoState82Action,
	{_State81, FuncToken}:                         _GotoState76Action,
	{_State81, LparenToken}:                       _GotoState79Action,
	{_State81, LbracketToken}:                     _GotoState27Action,
	{_State81, DotToken}:                          _GotoState73Action,
	{_State81, QuestionToken}:                     _GotoState80Action,
	{_State81, ExclaimToken}:                      _GotoState75Action,
	{_State81, TildeTildeToken}:                   _GotoState81Action,
	{_State81, BitNegToken}:                       _GotoState72Action,
	{_State81, BitAndToken}:                       _GotoState71Action,
	{_State81, LexErrorToken}:                     _GotoState78Action,
	{_State81, InitializableTypeType}:             _GotoState88Action,
	{_State81, AtomTypeType}:                      _GotoState83Action,
	{_State81, ReturnableTypeType}:                _GotoState172Action,
	{_State81, ImplicitStructDefType}:             _GotoState87Action,
	{_State81, ExplicitStructDefType}:             _GotoState46Action,
	{_State81, ImplicitEnumDefType}:               _GotoState86Action,
	{_State81, ExplicitEnumDefType}:               _GotoState84Action,
	{_State81, TraitDefType}:                      _GotoState90Action,
	{_State81, FuncTypeType}:                      _GotoState85Action,
	{_State82, LparenToken}:                       _GotoState173Action,
	{_State91, RbracketToken}:                     _GotoState178Action,
	{_State91, CommaToken}:                        _GotoState176Action,
	{_State91, ColonToken}:                        _GotoState175Action,
	{_State91, AddToken}:                          _GotoState174Action,
	{_State91, SubToken}:                          _GotoState179Action,
	{_State91, MulToken}:                          _GotoState177Action,
	{_State93, AssignToken}:                       _GotoState180Action,
	{_State95, RparenToken}:                       _GotoState182Action,
	{_State95, CommaToken}:                        _GotoState181Action,
	{_State96, ColonToken}:                        _GotoState183Action,
	{_State98, ColonToken}:                        _GotoState184Action,
	{_State99, IdentifierToken}:                   _GotoState163Action,
	{_State99, UnsafeToken}:                       _GotoState59Action,
	{_State99, StructToken}:                       _GotoState34Action,
	{_State99, EnumToken}:                         _GotoState74Action,
	{_State99, TraitToken}:                        _GotoState82Action,
	{_State99, FuncToken}:                         _GotoState76Action,
	{_State99, LparenToken}:                       _GotoState79Action,
	{_State99, LbracketToken}:                     _GotoState27Action,
	{_State99, DotToken}:                          _GotoState73Action,
	{_State99, QuestionToken}:                     _GotoState80Action,
	{_State99, ExclaimToken}:                      _GotoState75Action,
	{_State99, TildeTildeToken}:                   _GotoState81Action,
	{_State99, BitNegToken}:                       _GotoState72Action,
	{_State99, BitAndToken}:                       _GotoState71Action,
	{_State99, LexErrorToken}:                     _GotoState78Action,
	{_State99, UnsafeStatementType}:               _GotoState169Action,
	{_State99, InitializableTypeType}:             _GotoState88Action,
	{_State99, AtomTypeType}:                      _GotoState83Action,
	{_State99, ReturnableTypeType}:                _GotoState89Action,
	{_State99, ValueTypeType}:                     _GotoState170Action,
	{_State99, FieldDefType}:                      _GotoState186Action,
	{_State99, ImplicitStructDefType}:             _GotoState87Action,
	{_State99, ExplicitFieldDefsType}:             _GotoState185Action,
	{_State99, OptionalExplicitFieldDefsType}:     _GotoState187Action,
	{_State99, ExplicitStructDefType}:             _GotoState46Action,
	{_State99, ImplicitEnumDefType}:               _GotoState86Action,
	{_State99, ExplicitEnumDefType}:               _GotoState84Action,
	{_State99, TraitDefType}:                      _GotoState90Action,
	{_State99, FuncTypeType}:                      _GotoState85Action,
	{_State101, IdentifierToken}:                  _GotoState189Action,
	{_State101, LparenToken}:                      _GotoState101Action,
	{_State101, DotdotdotToken}:                   _GotoState188Action,
	{_State101, VarPatternType}:                   _GotoState192Action,
	{_State101, TuplePatternType}:                 _GotoState102Action,
	{_State101, FieldVarPatternsType}:             _GotoState191Action,
	{_State101, FieldVarPatternType}:              _GotoState190Action,
	{_State103, IdentifierToken}:                  _GotoState77Action,
	{_State103, StructToken}:                      _GotoState34Action,
	{_State103, EnumToken}:                        _GotoState74Action,
	{_State103, TraitToken}:                       _GotoState82Action,
	{_State103, FuncToken}:                        _GotoState76Action,
	{_State103, LparenToken}:                      _GotoState79Action,
	{_State103, LbracketToken}:                    _GotoState27Action,
	{_State103, DotToken}:                         _GotoState73Action,
	{_State103, QuestionToken}:                    _GotoState80Action,
	{_State103, ExclaimToken}:                     _GotoState75Action,
	{_State103, TildeTildeToken}:                  _GotoState81Action,
	{_State103, BitNegToken}:                      _GotoState72Action,
	{_State103, BitAndToken}:                      _GotoState71Action,
	{_State103, LexErrorToken}:                    _GotoState78Action,
	{_State103, OptionalValueTypeType}:            _GotoState193Action,
	{_State103, InitializableTypeType}:            _GotoState88Action,
	{_State103, AtomTypeType}:                     _GotoState83Action,
	{_State103, ReturnableTypeType}:               _GotoState89Action,
	{_State103, ValueTypeType}:                    _GotoState194Action,
	{_State103, ImplicitStructDefType}:            _GotoState87Action,
	{_State103, ExplicitStructDefType}:            _GotoState46Action,
	{_State103, ImplicitEnumDefType}:              _GotoState86Action,
	{_State103, ExplicitEnumDefType}:              _GotoState84Action,
	{_State103, TraitDefType}:                     _GotoState90Action,
	{_State103, FuncTypeType}:                     _GotoState85Action,
	{_State104, IdentifierToken}:                  _GotoState77Action,
	{_State104, StructToken}:                      _GotoState34Action,
	{_State104, EnumToken}:                        _GotoState74Action,
	{_State104, TraitToken}:                       _GotoState82Action,
	{_State104, FuncToken}:                        _GotoState76Action,
	{_State104, LparenToken}:                      _GotoState79Action,
	{_State104, LbracketToken}:                    _GotoState27Action,
	{_State104, DotToken}:                         _GotoState73Action,
	{_State104, QuestionToken}:                    _GotoState80Action,
	{_State104, ExclaimToken}:                     _GotoState75Action,
	{_State104, TildeTildeToken}:                  _GotoState81Action,
	{_State104, BitNegToken}:                      _GotoState72Action,
	{_State104, BitAndToken}:                      _GotoState71Action,
	{_State104, LexErrorToken}:                    _GotoState78Action,
	{_State104, OptionalGenericArgumentsType}:     _GotoState196Action,
	{_State104, GenericArgumentsType}:             _GotoState195Action,
	{_State104, InitializableTypeType}:            _GotoState88Action,
	{_State104, AtomTypeType}:                     _GotoState83Action,
	{_State104, ReturnableTypeType}:               _GotoState89Action,
	{_State104, ValueTypeType}:                    _GotoState197Action,
	{_State104, ImplicitStructDefType}:            _GotoState87Action,
	{_State104, ExplicitStructDefType}:            _GotoState46Action,
	{_State104, ImplicitEnumDefType}:              _GotoState86Action,
	{_State104, ExplicitEnumDefType}:              _GotoState84Action,
	{_State104, TraitDefType}:                     _GotoState90Action,
	{_State104, FuncTypeType}:                     _GotoState85Action,
	{_State105, IdentifierToken}:                  _GotoState198Action,
	{_State106, IntegerLiteralToken}:              _GotoState25Action,
	{_State106, FloatLiteralToken}:                _GotoState22Action,
	{_State106, RuneLiteralToken}:                 _GotoState32Action,
	{_State106, StringLiteralToken}:               _GotoState33Action,
	{_State106, IdentifierToken}:                  _GotoState93Action,
	{_State106, TrueToken}:                        _GotoState36Action,
	{_State106, FalseToken}:                       _GotoState21Action,
	{_State106, StructToken}:                      _GotoState34Action,
	{_State106, FuncToken}:                        _GotoState23Action,
	{_State106, VarToken}:                         _GotoState37Action,
	{_State106, LabelDeclToken}:                   _GotoState26Action,
	{_State106, LparenToken}:                      _GotoState29Action,
	{_State106, LbracketToken}:                    _GotoState27Action,
	{_State106, DotdotdotToken}:                   _GotoState92Action,
	{_State106, NotToken}:                         _GotoState31Action,
	{_State106, SubToken}:                         _GotoState35Action,
	{_State106, MulToken}:                         _GotoState30Action,
	{_State106, BitNegToken}:                      _GotoState20Action,
	{_State106, BitAndToken}:                      _GotoState19Action,
	{_State106, LexErrorToken}:                    _GotoState28Action,
	{_State106, ExpressionType}:                   _GotoState97Action,
	{_State106, OptionalLabelDeclType}:            _GotoState51Action,
	{_State106, SequenceExprType}:                 _GotoState56Action,
	{_State106, BlockExprType}:                    _GotoState43Action,
	{_State106, CallExprType}:                     _GotoState44Action,
	{_State106, ArgumentType}:                     _GotoState199Action,
	{_State106, ColonExpressionsType}:             _GotoState96Action,
	{_State106, OptionalExpressionType}:           _GotoState98Action,
	{_State106, AtomExprType}:                     _GotoState42Action,
	{_State106, LiteralType}:                      _GotoState49Action,
	{_State106, ImplicitStructExprType}:           _GotoState47Action,
	{_State106, AccessExprType}:                   _GotoState38Action,
	{_State106, PostfixUnaryExprType}:             _GotoState53Action,
	{_State106, PrefixUnaryOpType}:                _GotoState55Action,
	{_State106, PrefixUnaryExprType}:              _GotoState54Action,
	{_State106, MulExprType}:                      _GotoState50Action,
	{_State106, AddExprType}:                      _GotoState39Action,
	{_State106, CmpExprType}:                      _GotoState45Action,
	{_State106, AndExprType}:                      _GotoState40Action,
	{_State106, OrExprType}:                       _GotoState52Action,
	{_State106, InitializableTypeType}:            _GotoState48Action,
	{_State106, ExplicitStructDefType}:            _GotoState46Action,
	{_State106, AnonymousFuncExprType}:            _GotoState41Action,
	{_State108, LparenToken}:                      _GotoState200Action,
	{_State113, IntegerLiteralToken}:              _GotoState25Action,
	{_State113, FloatLiteralToken}:                _GotoState22Action,
	{_State113, RuneLiteralToken}:                 _GotoState32Action,
	{_State113, StringLiteralToken}:               _GotoState33Action,
	{_State113, IdentifierToken}:                  _GotoState24Action,
	{_State113, TrueToken}:                        _GotoState36Action,
	{_State113, FalseToken}:                       _GotoState21Action,
	{_State113, StructToken}:                      _GotoState34Action,
	{_State113, FuncToken}:                        _GotoState23Action,
	{_State113, LabelDeclToken}:                   _GotoState26Action,
	{_State113, LparenToken}:                      _GotoState29Action,
	{_State113, LbracketToken}:                    _GotoState27Action,
	{_State113, NotToken}:                         _GotoState31Action,
	{_State113, SubToken}:                         _GotoState35Action,
	{_State113, MulToken}:                         _GotoState30Action,
	{_State113, BitNegToken}:                      _GotoState20Action,
	{_State113, BitAndToken}:                      _GotoState19Action,
	{_State113, LexErrorToken}:                    _GotoState28Action,
	{_State113, OptionalLabelDeclType}:            _GotoState140Action,
	{_State113, BlockExprType}:                    _GotoState43Action,
	{_State113, CallExprType}:                     _GotoState44Action,
	{_State113, AtomExprType}:                     _GotoState42Action,
	{_State113, LiteralType}:                      _GotoState49Action,
	{_State113, ImplicitStructExprType}:           _GotoState47Action,
	{_State113, AccessExprType}:                   _GotoState38Action,
	{_State113, PostfixUnaryExprType}:             _GotoState53Action,
	{_State113, PrefixUnaryOpType}:                _GotoState55Action,
	{_State113, PrefixUnaryExprType}:              _GotoState54Action,
	{_State113, MulExprType}:                      _GotoState201Action,
	{_State113, InitializableTypeType}:            _GotoState48Action,
	{_State113, ExplicitStructDefType}:            _GotoState46Action,
	{_State113, AnonymousFuncExprType}:            _GotoState41Action,
	{_State114, IntegerLiteralToken}:              _GotoState25Action,
	{_State114, FloatLiteralToken}:                _GotoState22Action,
	{_State114, RuneLiteralToken}:                 _GotoState32Action,
	{_State114, StringLiteralToken}:               _GotoState33Action,
	{_State114, IdentifierToken}:                  _GotoState24Action,
	{_State114, TrueToken}:                        _GotoState36Action,
	{_State114, FalseToken}:                       _GotoState21Action,
	{_State114, StructToken}:                      _GotoState34Action,
	{_State114, FuncToken}:                        _GotoState23Action,
	{_State114, LabelDeclToken}:                   _GotoState26Action,
	{_State114, LparenToken}:                      _GotoState29Action,
	{_State114, LbracketToken}:                    _GotoState27Action,
	{_State114, NotToken}:                         _GotoState31Action,
	{_State114, SubToken}:                         _GotoState35Action,
	{_State114, MulToken}:                         _GotoState30Action,
	{_State114, BitNegToken}:                      _GotoState20Action,
	{_State114, BitAndToken}:                      _GotoState19Action,
	{_State114, LexErrorToken}:                    _GotoState28Action,
	{_State114, OptionalLabelDeclType}:            _GotoState140Action,
	{_State114, BlockExprType}:                    _GotoState43Action,
	{_State114, CallExprType}:                     _GotoState44Action,
	{_State114, AtomExprType}:                     _GotoState42Action,
	{_State114, LiteralType}:                      _GotoState49Action,
	{_State114, ImplicitStructExprType}:           _GotoState47Action,
	{_State114, AccessExprType}:                   _GotoState38Action,
	{_State114, PostfixUnaryExprType}:             _GotoState53Action,
	{_State114, PrefixUnaryOpType}:                _GotoState55Action,
	{_State114, PrefixUnaryExprType}:              _GotoState54Action,
	{_State114, MulExprType}:                      _GotoState50Action,
	{_State114, AddExprType}:                      _GotoState39Action,
	{_State114, CmpExprType}:                      _GotoState202Action,
	{_State114, InitializableTypeType}:            _GotoState48Action,
	{_State114, ExplicitStructDefType}:            _GotoState46Action,
	{_State114, AnonymousFuncExprType}:            _GotoState41Action,
	{_State121, IntegerLiteralToken}:              _GotoState25Action,
	{_State121, FloatLiteralToken}:                _GotoState22Action,
	{_State121, RuneLiteralToken}:                 _GotoState32Action,
	{_State121, StringLiteralToken}:               _GotoState33Action,
	{_State121, IdentifierToken}:                  _GotoState24Action,
	{_State121, TrueToken}:                        _GotoState36Action,
	{_State121, FalseToken}:                       _GotoState21Action,
	{_State121, StructToken}:                      _GotoState34Action,
	{_State121, FuncToken}:                        _GotoState23Action,
	{_State121, LabelDeclToken}:                   _GotoState26Action,
	{_State121, LparenToken}:                      _GotoState29Action,
	{_State121, LbracketToken}:                    _GotoState27Action,
	{_State121, NotToken}:                         _GotoState31Action,
	{_State121, SubToken}:                         _GotoState35Action,
	{_State121, MulToken}:                         _GotoState30Action,
	{_State121, BitNegToken}:                      _GotoState20Action,
	{_State121, BitAndToken}:                      _GotoState19Action,
	{_State121, LexErrorToken}:                    _GotoState28Action,
	{_State121, OptionalLabelDeclType}:            _GotoState140Action,
	{_State121, BlockExprType}:                    _GotoState43Action,
	{_State121, CallExprType}:                     _GotoState44Action,
	{_State121, AtomExprType}:                     _GotoState42Action,
	{_State121, LiteralType}:                      _GotoState49Action,
	{_State121, ImplicitStructExprType}:           _GotoState47Action,
	{_State121, AccessExprType}:                   _GotoState38Action,
	{_State121, PostfixUnaryExprType}:             _GotoState53Action,
	{_State121, PrefixUnaryOpType}:                _GotoState55Action,
	{_State121, PrefixUnaryExprType}:              _GotoState54Action,
	{_State121, MulExprType}:                      _GotoState50Action,
	{_State121, AddExprType}:                      _GotoState203Action,
	{_State121, InitializableTypeType}:            _GotoState48Action,
	{_State121, ExplicitStructDefType}:            _GotoState46Action,
	{_State121, AnonymousFuncExprType}:            _GotoState41Action,
	{_State122, IntegerLiteralToken}:              _GotoState25Action,
	{_State122, FloatLiteralToken}:                _GotoState22Action,
	{_State122, RuneLiteralToken}:                 _GotoState32Action,
	{_State122, StringLiteralToken}:               _GotoState33Action,
	{_State122, IdentifierToken}:                  _GotoState93Action,
	{_State122, TrueToken}:                        _GotoState36Action,
	{_State122, FalseToken}:                       _GotoState21Action,
	{_State122, StructToken}:                      _GotoState34Action,
	{_State122, FuncToken}:                        _GotoState23Action,
	{_State122, VarToken}:                         _GotoState37Action,
	{_State122, LabelDeclToken}:                   _GotoState26Action,
	{_State122, LparenToken}:                      _GotoState29Action,
	{_State122, LbracketToken}:                    _GotoState27Action,
	{_State122, DotdotdotToken}:                   _GotoState92Action,
	{_State122, NotToken}:                         _GotoState31Action,
	{_State122, SubToken}:                         _GotoState35Action,
	{_State122, MulToken}:                         _GotoState30Action,
	{_State122, BitNegToken}:                      _GotoState20Action,
	{_State122, BitAndToken}:                      _GotoState19Action,
	{_State122, LexErrorToken}:                    _GotoState28Action,
	{_State122, ExpressionType}:                   _GotoState97Action,
	{_State122, OptionalLabelDeclType}:            _GotoState51Action,
	{_State122, SequenceExprType}:                 _GotoState56Action,
	{_State122, BlockExprType}:                    _GotoState43Action,
	{_State122, CallExprType}:                     _GotoState44Action,
	{_State122, ArgumentsType}:                    _GotoState204Action,
	{_State122, ArgumentType}:                     _GotoState94Action,
	{_State122, ColonExpressionsType}:             _GotoState96Action,
	{_State122, OptionalExpressionType}:           _GotoState98Action,
	{_State122, AtomExprType}:                     _GotoState42Action,
	{_State122, LiteralType}:                      _GotoState49Action,
	{_State122, ImplicitStructExprType}:           _GotoState47Action,
	{_State122, AccessExprType}:                   _GotoState38Action,
	{_State122, PostfixUnaryExprType}:             _GotoState53Action,
	{_State122, PrefixUnaryOpType}:                _GotoState55Action,
	{_State122, PrefixUnaryExprType}:              _GotoState54Action,
	{_State122, MulExprType}:                      _GotoState50Action,
	{_State122, AddExprType}:                      _GotoState39Action,
	{_State122, CmpExprType}:                      _GotoState45Action,
	{_State122, AndExprType}:                      _GotoState40Action,
	{_State122, OrExprType}:                       _GotoState52Action,
	{_State122, InitializableTypeType}:            _GotoState48Action,
	{_State122, ExplicitStructDefType}:            _GotoState46Action,
	{_State122, AnonymousFuncExprType}:            _GotoState41Action,
	{_State129, IntegerLiteralToken}:              _GotoState25Action,
	{_State129, FloatLiteralToken}:                _GotoState22Action,
	{_State129, RuneLiteralToken}:                 _GotoState32Action,
	{_State129, StringLiteralToken}:               _GotoState33Action,
	{_State129, IdentifierToken}:                  _GotoState24Action,
	{_State129, TrueToken}:                        _GotoState36Action,
	{_State129, FalseToken}:                       _GotoState21Action,
	{_State129, StructToken}:                      _GotoState34Action,
	{_State129, FuncToken}:                        _GotoState23Action,
	{_State129, LabelDeclToken}:                   _GotoState26Action,
	{_State129, LparenToken}:                      _GotoState29Action,
	{_State129, LbracketToken}:                    _GotoState27Action,
	{_State129, NotToken}:                         _GotoState31Action,
	{_State129, SubToken}:                         _GotoState35Action,
	{_State129, MulToken}:                         _GotoState30Action,
	{_State129, BitNegToken}:                      _GotoState20Action,
	{_State129, BitAndToken}:                      _GotoState19Action,
	{_State129, LexErrorToken}:                    _GotoState28Action,
	{_State129, OptionalLabelDeclType}:            _GotoState140Action,
	{_State129, BlockExprType}:                    _GotoState43Action,
	{_State129, CallExprType}:                     _GotoState44Action,
	{_State129, AtomExprType}:                     _GotoState42Action,
	{_State129, LiteralType}:                      _GotoState49Action,
	{_State129, ImplicitStructExprType}:           _GotoState47Action,
	{_State129, AccessExprType}:                   _GotoState38Action,
	{_State129, PostfixUnaryExprType}:             _GotoState53Action,
	{_State129, PrefixUnaryOpType}:                _GotoState55Action,
	{_State129, PrefixUnaryExprType}:              _GotoState205Action,
	{_State129, InitializableTypeType}:            _GotoState48Action,
	{_State129, ExplicitStructDefType}:            _GotoState46Action,
	{_State129, AnonymousFuncExprType}:            _GotoState41Action,
	{_State130, LbraceToken}:                      _GotoState133Action,
	{_State130, BlockBodyType}:                    _GotoState206Action,
	{_State131, IntegerLiteralToken}:              _GotoState25Action,
	{_State131, FloatLiteralToken}:                _GotoState22Action,
	{_State131, RuneLiteralToken}:                 _GotoState32Action,
	{_State131, StringLiteralToken}:               _GotoState33Action,
	{_State131, IdentifierToken}:                  _GotoState24Action,
	{_State131, TrueToken}:                        _GotoState36Action,
	{_State131, FalseToken}:                       _GotoState21Action,
	{_State131, StructToken}:                      _GotoState34Action,
	{_State131, FuncToken}:                        _GotoState23Action,
	{_State131, VarToken}:                         _GotoState209Action,
	{_State131, LabelDeclToken}:                   _GotoState26Action,
	{_State131, LparenToken}:                      _GotoState29Action,
	{_State131, LbracketToken}:                    _GotoState27Action,
	{_State131, DotToken}:                         _GotoState207Action,
	{_State131, SemicolonToken}:                   _GotoState208Action,
	{_State131, NotToken}:                         _GotoState31Action,
	{_State131, SubToken}:                         _GotoState35Action,
	{_State131, MulToken}:                         _GotoState30Action,
	{_State131, BitNegToken}:                      _GotoState20Action,
	{_State131, BitAndToken}:                      _GotoState19Action,
	{_State131, LexErrorToken}:                    _GotoState28Action,
	{_State131, ExpressionType}:                   _GotoState210Action,
	{_State131, OptionalLabelDeclType}:            _GotoState51Action,
	{_State131, SequenceExprType}:                 _GotoState212Action,
	{_State131, BlockExprType}:                    _GotoState43Action,
	{_State131, PatternType}:                      _GotoState211Action,
	{_State131, CallExprType}:                     _GotoState44Action,
	{_State131, AtomExprType}:                     _GotoState42Action,
	{_State131, LiteralType}:                      _GotoState49Action,
	{_State131, ImplicitStructExprType}:           _GotoState47Action,
	{_State131, AccessExprType}:                   _GotoState38Action,
	{_State131, PostfixUnaryExprType}:             _GotoState53Action,
	{_State131, PrefixUnaryOpType}:                _GotoState55Action,
	{_State131, PrefixUnaryExprType}:              _GotoState54Action,
	{_State131, MulExprType}:                      _GotoState50Action,
	{_State131, AddExprType}:                      _GotoState39Action,
	{_State131, CmpExprType}:                      _GotoState45Action,
	{_State131, AndExprType}:                      _GotoState40Action,
	{_State131, OrExprType}:                       _GotoState52Action,
	{_State131, InitializableTypeType}:            _GotoState48Action,
	{_State131, ExplicitStructDefType}:            _GotoState46Action,
	{_State131, AnonymousFuncExprType}:            _GotoState41Action,
	{_State132, IntegerLiteralToken}:              _GotoState25Action,
	{_State132, FloatLiteralToken}:                _GotoState22Action,
	{_State132, RuneLiteralToken}:                 _GotoState32Action,
	{_State132, StringLiteralToken}:               _GotoState33Action,
	{_State132, IdentifierToken}:                  _GotoState24Action,
	{_State132, TrueToken}:                        _GotoState36Action,
	{_State132, FalseToken}:                       _GotoState21Action,
	{_State132, StructToken}:                      _GotoState34Action,
	{_State132, FuncToken}:                        _GotoState23Action,
	{_State132, LabelDeclToken}:                   _GotoState26Action,
	{_State132, LparenToken}:                      _GotoState29Action,
	{_State132, LbracketToken}:                    _GotoState27Action,
	{_State132, NotToken}:                         _GotoState31Action,
	{_State132, SubToken}:                         _GotoState35Action,
	{_State132, MulToken}:                         _GotoState30Action,
	{_State132, BitNegToken}:                      _GotoState20Action,
	{_State132, BitAndToken}:                      _GotoState19Action,
	{_State132, LexErrorToken}:                    _GotoState28Action,
	{_State132, OptionalLabelDeclType}:            _GotoState140Action,
	{_State132, SequenceExprType}:                 _GotoState213Action,
	{_State132, BlockExprType}:                    _GotoState43Action,
	{_State132, CallExprType}:                     _GotoState44Action,
	{_State132, AtomExprType}:                     _GotoState42Action,
	{_State132, LiteralType}:                      _GotoState49Action,
	{_State132, ImplicitStructExprType}:           _GotoState47Action,
	{_State132, AccessExprType}:                   _GotoState38Action,
	{_State132, PostfixUnaryExprType}:             _GotoState53Action,
	{_State132, PrefixUnaryOpType}:                _GotoState55Action,
	{_State132, PrefixUnaryExprType}:              _GotoState54Action,
	{_State132, MulExprType}:                      _GotoState50Action,
	{_State132, AddExprType}:                      _GotoState39Action,
	{_State132, CmpExprType}:                      _GotoState45Action,
	{_State132, AndExprType}:                      _GotoState40Action,
	{_State132, OrExprType}:                       _GotoState52Action,
	{_State132, InitializableTypeType}:            _GotoState48Action,
	{_State132, ExplicitStructDefType}:            _GotoState46Action,
	{_State132, AnonymousFuncExprType}:            _GotoState41Action,
	{_State133, StatementsType}:                   _GotoState214Action,
	{_State134, IntegerLiteralToken}:              _GotoState25Action,
	{_State134, FloatLiteralToken}:                _GotoState22Action,
	{_State134, RuneLiteralToken}:                 _GotoState32Action,
	{_State134, StringLiteralToken}:               _GotoState33Action,
	{_State134, IdentifierToken}:                  _GotoState24Action,
	{_State134, TrueToken}:                        _GotoState36Action,
	{_State134, FalseToken}:                       _GotoState21Action,
	{_State134, StructToken}:                      _GotoState34Action,
	{_State134, FuncToken}:                        _GotoState23Action,
	{_State134, LabelDeclToken}:                   _GotoState26Action,
	{_State134, LparenToken}:                      _GotoState29Action,
	{_State134, LbracketToken}:                    _GotoState27Action,
	{_State134, NotToken}:                         _GotoState31Action,
	{_State134, SubToken}:                         _GotoState35Action,
	{_State134, MulToken}:                         _GotoState30Action,
	{_State134, BitNegToken}:                      _GotoState20Action,
	{_State134, BitAndToken}:                      _GotoState19Action,
	{_State134, LexErrorToken}:                    _GotoState28Action,
	{_State134, OptionalLabelDeclType}:            _GotoState140Action,
	{_State134, SequenceExprType}:                 _GotoState215Action,
	{_State134, BlockExprType}:                    _GotoState43Action,
	{_State134, CallExprType}:                     _GotoState44Action,
	{_State134, AtomExprType}:                     _GotoState42Action,
	{_State134, LiteralType}:                      _GotoState49Action,
	{_State134, ImplicitStructExprType}:           _GotoState47Action,
	{_State134, AccessExprType}:                   _GotoState38Action,
	{_State134, PostfixUnaryExprType}:             _GotoState53Action,
	{_State134, PrefixUnaryOpType}:                _GotoState55Action,
	{_State134, PrefixUnaryExprType}:              _GotoState54Action,
	{_State134, MulExprType}:                      _GotoState50Action,
	{_State134, AddExprType}:                      _GotoState39Action,
	{_State134, CmpExprType}:                      _GotoState45Action,
	{_State134, AndExprType}:                      _GotoState40Action,
	{_State134, OrExprType}:                       _GotoState52Action,
	{_State134, InitializableTypeType}:            _GotoState48Action,
	{_State134, ExplicitStructDefType}:            _GotoState46Action,
	{_State134, AnonymousFuncExprType}:            _GotoState41Action,
	{_State139, IntegerLiteralToken}:              _GotoState25Action,
	{_State139, FloatLiteralToken}:                _GotoState22Action,
	{_State139, RuneLiteralToken}:                 _GotoState32Action,
	{_State139, StringLiteralToken}:               _GotoState33Action,
	{_State139, IdentifierToken}:                  _GotoState24Action,
	{_State139, TrueToken}:                        _GotoState36Action,
	{_State139, FalseToken}:                       _GotoState21Action,
	{_State139, StructToken}:                      _GotoState34Action,
	{_State139, FuncToken}:                        _GotoState23Action,
	{_State139, LabelDeclToken}:                   _GotoState26Action,
	{_State139, LparenToken}:                      _GotoState29Action,
	{_State139, LbracketToken}:                    _GotoState27Action,
	{_State139, NotToken}:                         _GotoState31Action,
	{_State139, SubToken}:                         _GotoState35Action,
	{_State139, MulToken}:                         _GotoState30Action,
	{_State139, BitNegToken}:                      _GotoState20Action,
	{_State139, BitAndToken}:                      _GotoState19Action,
	{_State139, LexErrorToken}:                    _GotoState28Action,
	{_State139, OptionalLabelDeclType}:            _GotoState140Action,
	{_State139, BlockExprType}:                    _GotoState43Action,
	{_State139, CallExprType}:                     _GotoState44Action,
	{_State139, AtomExprType}:                     _GotoState42Action,
	{_State139, LiteralType}:                      _GotoState49Action,
	{_State139, ImplicitStructExprType}:           _GotoState47Action,
	{_State139, AccessExprType}:                   _GotoState38Action,
	{_State139, PostfixUnaryExprType}:             _GotoState53Action,
	{_State139, PrefixUnaryOpType}:                _GotoState55Action,
	{_State139, PrefixUnaryExprType}:              _GotoState54Action,
	{_State139, MulExprType}:                      _GotoState50Action,
	{_State139, AddExprType}:                      _GotoState39Action,
	{_State139, CmpExprType}:                      _GotoState45Action,
	{_State139, AndExprType}:                      _GotoState216Action,
	{_State139, InitializableTypeType}:            _GotoState48Action,
	{_State139, ExplicitStructDefType}:            _GotoState46Action,
	{_State139, AnonymousFuncExprType}:            _GotoState41Action,
	{_State140, LbraceToken}:                      _GotoState133Action,
	{_State140, BlockBodyType}:                    _GotoState135Action,
	{_State142, IdentifierToken}:                  _GotoState217Action,
	{_State143, PackageToken}:                     _GotoState16Action,
	{_State143, UnsafeToken}:                      _GotoState59Action,
	{_State143, TypeToken}:                        _GotoState17Action,
	{_State143, FuncToken}:                        _GotoState18Action,
	{_State143, DefinitionType}:                   _GotoState218Action,
	{_State143, UnsafeStatementType}:              _GotoState65Action,
	{_State143, TypeDefType}:                      _GotoState64Action,
	{_State143, NamedFuncDefType}:                 _GotoState62Action,
	{_State143, PackageDefType}:                   _GotoState63Action,
	{_State145, PackageStatementsType}:            _GotoState219Action,
	{_State146, IdentifierToken}:                  _GotoState220Action,
	{_State146, GenericParameterDefType}:          _GotoState221Action,
	{_State146, GenericParameterDefsType}:         _GotoState222Action,
	{_State146, OptionalGenericParameterDefsType}: _GotoState223Action,
	{_State147, IdentifierToken}:                  _GotoState77Action,
	{_State147, StructToken}:                      _GotoState34Action,
	{_State147, EnumToken}:                        _GotoState74Action,
	{_State147, TraitToken}:                       _GotoState82Action,
	{_State147, FuncToken}:                        _GotoState76Action,
	{_State147, LparenToken}:                      _GotoState79Action,
	{_State147, LbracketToken}:                    _GotoState27Action,
	{_State147, DotToken}:                         _GotoState73Action,
	{_State147, QuestionToken}:                    _GotoState80Action,
	{_State147, ExclaimToken}:                     _GotoState75Action,
	{_State147, TildeTildeToken}:                  _GotoState81Action,
	{_State147, BitNegToken}:                      _GotoState72Action,
	{_State147, BitAndToken}:                      _GotoState71Action,
	{_State147, LexErrorToken}:                    _GotoState78Action,
	{_State147, InitializableTypeType}:            _GotoState88Action,
	{_State147, AtomTypeType}:                     _GotoState83Action,
	{_State147, ReturnableTypeType}:               _GotoState89Action,
	{_State147, ValueTypeType}:                    _GotoState224Action,
	{_State147, ImplicitStructDefType}:            _GotoState87Action,
	{_State147, ExplicitStructDefType}:            _GotoState46Action,
	{_State147, ImplicitEnumDefType}:              _GotoState86Action,
	{_State147, ExplicitEnumDefType}:              _GotoState84Action,
	{_State147, TraitDefType}:                     _GotoState90Action,
	{_State147, FuncTypeType}:                     _GotoState85Action,
	{_State148, IdentifierToken}:                  _GotoState77Action,
	{_State148, StructToken}:                      _GotoState34Action,
	{_State148, EnumToken}:                        _GotoState74Action,
	{_State148, TraitToken}:                       _GotoState82Action,
	{_State148, FuncToken}:                        _GotoState76Action,
	{_State148, LparenToken}:                      _GotoState79Action,
	{_State148, LbracketToken}:                    _GotoState27Action,
	{_State148, DotToken}:                         _GotoState73Action,
	{_State148, QuestionToken}:                    _GotoState80Action,
	{_State148, ExclaimToken}:                     _GotoState75Action,
	{_State148, TildeTildeToken}:                  _GotoState81Action,
	{_State148, BitNegToken}:                      _GotoState72Action,
	{_State148, BitAndToken}:                      _GotoState71Action,
	{_State148, LexErrorToken}:                    _GotoState78Action,
	{_State148, InitializableTypeType}:            _GotoState88Action,
	{_State148, AtomTypeType}:                     _GotoState83Action,
	{_State148, ReturnableTypeType}:               _GotoState89Action,
	{_State148, ValueTypeType}:                    _GotoState225Action,
	{_State148, ImplicitStructDefType}:            _GotoState87Action,
	{_State148, ExplicitStructDefType}:            _GotoState46Action,
	{_State148, ImplicitEnumDefType}:              _GotoState86Action,
	{_State148, ExplicitEnumDefType}:              _GotoState84Action,
	{_State148, TraitDefType}:                     _GotoState90Action,
	{_State148, FuncTypeType}:                     _GotoState85Action,
	{_State149, IdentifierToken}:                  _GotoState77Action,
	{_State149, StructToken}:                      _GotoState34Action,
	{_State149, EnumToken}:                        _GotoState74Action,
	{_State149, TraitToken}:                       _GotoState82Action,
	{_State149, FuncToken}:                        _GotoState76Action,
	{_State149, LparenToken}:                      _GotoState79Action,
	{_State149, LbracketToken}:                    _GotoState27Action,
	{_State149, DotToken}:                         _GotoState73Action,
	{_State149, QuestionToken}:                    _GotoState80Action,
	{_State149, ExclaimToken}:                     _GotoState75Action,
	{_State149, DotdotdotToken}:                   _GotoState226Action,
	{_State149, TildeTildeToken}:                  _GotoState81Action,
	{_State149, BitNegToken}:                      _GotoState72Action,
	{_State149, BitAndToken}:                      _GotoState71Action,
	{_State149, LexErrorToken}:                    _GotoState78Action,
	{_State149, InitializableTypeType}:            _GotoState88Action,
	{_State149, AtomTypeType}:                     _GotoState83Action,
	{_State149, ReturnableTypeType}:               _GotoState89Action,
	{_State149, ValueTypeType}:                    _GotoState227Action,
	{_State149, ImplicitStructDefType}:            _GotoState87Action,
	{_State149, ExplicitStructDefType}:            _GotoState46Action,
	{_State149, ImplicitEnumDefType}:              _GotoState86Action,
	{_State149, ExplicitEnumDefType}:              _GotoState84Action,
	{_State149, TraitDefType}:                     _GotoState90Action,
	{_State149, FuncTypeType}:                     _GotoState85Action,
	{_State150, RparenToken}:                      _GotoState228Action,
	{_State151, DollarLbracketToken}:              _GotoState146Action,
	{_State151, OptionalGenericParametersType}:    _GotoState229Action,
	{_State152, RparenToken}:                      _GotoState230Action,
	{_State154, CommaToken}:                       _GotoState231Action,
	{_State158, IdentifierToken}:                  _GotoState163Action,
	{_State158, UnsafeToken}:                      _GotoState59Action,
	{_State158, StructToken}:                      _GotoState34Action,
	{_State158, EnumToken}:                        _GotoState74Action,
	{_State158, TraitToken}:                       _GotoState82Action,
	{_State158, FuncToken}:                        _GotoState76Action,
	{_State158, LparenToken}:                      _GotoState79Action,
	{_State158, LbracketToken}:                    _GotoState27Action,
	{_State158, DotToken}:                         _GotoState73Action,
	{_State158, QuestionToken}:                    _GotoState80Action,
	{_State158, ExclaimToken}:                     _GotoState75Action,
	{_State158, TildeTildeToken}:                  _GotoState81Action,
	{_State158, BitNegToken}:                      _GotoState72Action,
	{_State158, BitAndToken}:                      _GotoState71Action,
	{_State158, LexErrorToken}:                    _GotoState78Action,
	{_State158, UnsafeStatementType}:              _GotoState169Action,
	{_State158, InitializableTypeType}:            _GotoState88Action,
	{_State158, AtomTypeType}:                     _GotoState83Action,
	{_State158, ReturnableTypeType}:               _GotoState89Action,
	{_State158, ValueTypeType}:                    _GotoState170Action,
	{_State158, FieldDefType}:                     _GotoState234Action,
	{_State158, ImplicitStructDefType}:            _GotoState87Action,
	{_State158, ExplicitStructDefType}:            _GotoState46Action,
	{_State158, EnumValueDefType}:                 _GotoState232Action,
	{_State158, ImplicitEnumValueDefsType}:        _GotoState235Action,
	{_State158, ImplicitEnumDefType}:              _GotoState86Action,
	{_State158, ExplicitEnumValueDefsType}:        _GotoState233Action,
	{_State158, ExplicitEnumDefType}:              _GotoState84Action,
	{_State158, TraitDefType}:                     _GotoState90Action,
	{_State158, FuncTypeType}:                     _GotoState85Action,
	{_State160, IdentifierToken}:                  _GotoState237Action,
	{_State160, StructToken}:                      _GotoState34Action,
	{_State160, EnumToken}:                        _GotoState74Action,
	{_State160, TraitToken}:                       _GotoState82Action,
	{_State160, FuncToken}:                        _GotoState76Action,
	{_State160, LparenToken}:                      _GotoState79Action,
	{_State160, LbracketToken}:                    _GotoState27Action,
	{_State160, DotToken}:                         _GotoState73Action,
	{_State160, QuestionToken}:                    _GotoState80Action,
	{_State160, ExclaimToken}:                     _GotoState75Action,
	{_State160, DotdotdotToken}:                   _GotoState236Action,
	{_State160, TildeTildeToken}:                  _GotoState81Action,
	{_State160, BitNegToken}:                      _GotoState72Action,
	{_State160, BitAndToken}:                      _GotoState71Action,
	{_State160, LexErrorToken}:                    _GotoState78Action,
	{_State160, InitializableTypeType}:            _GotoState88Action,
	{_State160, AtomTypeType}:                     _GotoState83Action,
	{_State160, ReturnableTypeType}:               _GotoState89Action,
	{_State160, ValueTypeType}:                    _GotoState241Action,
	{_State160, ImplicitStructDefType}:            _GotoState87Action,
	{_State160, ExplicitStructDefType}:            _GotoState46Action,
	{_State160, ImplicitEnumDefType}:              _GotoState86Action,
	{_State160, ExplicitEnumDefType}:              _GotoState84Action,
	{_State160, TraitDefType}:                     _GotoState90Action,
	{_State160, ParameterDeclType}:                _GotoState239Action,
	{_State160, ParameterDeclsType}:               _GotoState240Action,
	{_State160, OptionalParameterDeclsType}:       _GotoState238Action,
	{_State160, FuncTypeType}:                     _GotoState85Action,
	{_State161, IdentifierToken}:                  _GotoState242Action,
	{_State163, IdentifierToken}:                  _GotoState77Action,
	{_State163, StructToken}:                      _GotoState34Action,
	{_State163, EnumToken}:                        _GotoState74Action,
	{_State163, TraitToken}:                       _GotoState82Action,
	{_State163, FuncToken}:                        _GotoState76Action,
	{_State163, LparenToken}:                      _GotoState79Action,
	{_State163, LbracketToken}:                    _GotoState27Action,
	{_State163, DotToken}:                         _GotoState243Action,
	{_State163, QuestionToken}:                    _GotoState80Action,
	{_State163, ExclaimToken}:                     _GotoState75Action,
	{_State163, DollarLbracketToken}:              _GotoState104Action,
	{_State163, TildeTildeToken}:                  _GotoState81Action,
	{_State163, BitNegToken}:                      _GotoState72Action,
	{_State163, BitAndToken}:                      _GotoState71Action,
	{_State163, LexErrorToken}:                    _GotoState78Action,
	{_State163, OptionalGenericBindingType}:       _GotoState162Action,
	{_State163, InitializableTypeType}:            _GotoState88Action,
	{_State163, AtomTypeType}:                     _GotoState83Action,
	{_State163, ReturnableTypeType}:               _GotoState89Action,
	{_State163, ValueTypeType}:                    _GotoState244Action,
	{_State163, ImplicitStructDefType}:            _GotoState87Action,
	{_State163, ExplicitStructDefType}:            _GotoState46Action,
	{_State163, ImplicitEnumDefType}:              _GotoState86Action,
	{_State163, ExplicitEnumDefType}:              _GotoState84Action,
	{_State163, TraitDefType}:                     _GotoState90Action,
	{_State163, FuncTypeType}:                     _GotoState85Action,
	{_State164, OrToken}:                          _GotoState245Action,
	{_State165, AssignToken}:                      _GotoState246Action,
	{_State166, RparenToken}:                      _GotoState248Action,
	{_State166, OrToken}:                          _GotoState247Action,
	{_State167, CommaToken}:                       _GotoState249Action,
	{_State168, RparenToken}:                      _GotoState250Action,
	{_State170, AddToken}:                         _GotoState174Action,
	{_State170, SubToken}:                         _GotoState179Action,
	{_State170, MulToken}:                         _GotoState177Action,
	{_State173, IdentifierToken}:                  _GotoState163Action,
	{_State173, UnsafeToken}:                      _GotoState59Action,
	{_State173, StructToken}:                      _GotoState34Action,
	{_State173, EnumToken}:                        _GotoState74Action,
	{_State173, TraitToken}:                       _GotoState82Action,
	{_State173, FuncToken}:                        _GotoState251Action,
	{_State173, LparenToken}:                      _GotoState79Action,
	{_State173, LbracketToken}:                    _GotoState27Action,
	{_State173, DotToken}:                         _GotoState73Action,
	{_State173, QuestionToken}:                    _GotoState80Action,
	{_State173, ExclaimToken}:                     _GotoState75Action,
	{_State173, TildeTildeToken}:                  _GotoState81Action,
	{_State173, BitNegToken}:                      _GotoState72Action,
	{_State173, BitAndToken}:                      _GotoState71Action,
	{_State173, LexErrorToken}:                    _GotoState78Action,
	{_State173, UnsafeStatementType}:              _GotoState169Action,
	{_State173, InitializableTypeType}:            _GotoState88Action,
	{_State173, AtomTypeType}:                     _GotoState83Action,
	{_State173, ReturnableTypeType}:               _GotoState89Action,
	{_State173, ValueTypeType}:                    _GotoState170Action,
	{_State173, FieldDefType}:                     _GotoState252Action,
	{_State173, ImplicitStructDefType}:            _GotoState87Action,
	{_State173, ExplicitStructDefType}:            _GotoState46Action,
	{_State173, ImplicitEnumDefType}:              _GotoState86Action,
	{_State173, ExplicitEnumDefType}:              _GotoState84Action,
	{_State173, TraitPropertyType}:                _GotoState256Action,
	{_State173, TraitPropertiesType}:              _GotoState255Action,
	{_State173, OptionalTraitPropertiesType}:      _GotoState254Action,
	{_State173, TraitDefType}:                     _GotoState90Action,
	{_State173, FuncTypeType}:                     _GotoState85Action,
	{_State173, MethodSignatureType}:              _GotoState253Action,
	{_State174, IdentifierToken}:                  _GotoState77Action,
	{_State174, StructToken}:                      _GotoState34Action,
	{_State174, EnumToken}:                        _GotoState74Action,
	{_State174, TraitToken}:                       _GotoState82Action,
	{_State174, FuncToken}:                        _GotoState76Action,
	{_State174, LparenToken}:                      _GotoState79Action,
	{_State174, LbracketToken}:                    _GotoState27Action,
	{_State174, DotToken}:                         _GotoState73Action,
	{_State174, QuestionToken}:                    _GotoState80Action,
	{_State174, ExclaimToken}:                     _GotoState75Action,
	{_State174, TildeTildeToken}:                  _GotoState81Action,
	{_State174, BitNegToken}:                      _GotoState72Action,
	{_State174, BitAndToken}:                      _GotoState71Action,
	{_State174, LexErrorToken}:                    _GotoState78Action,
	{_State174, InitializableTypeType}:            _GotoState88Action,
	{_State174, AtomTypeType}:                     _GotoState83Action,
	{_State174, ReturnableTypeType}:               _GotoState257Action,
	{_State174, ImplicitStructDefType}:            _GotoState87Action,
	{_State174, ExplicitStructDefType}:            _GotoState46Action,
	{_State174, ImplicitEnumDefType}:              _GotoState86Action,
	{_State174, ExplicitEnumDefType}:              _GotoState84Action,
	{_State174, TraitDefType}:                     _GotoState90Action,
	{_State174, FuncTypeType}:                     _GotoState85Action,
	{_State175, IdentifierToken}:                  _GotoState77Action,
	{_State175, StructToken}:                      _GotoState34Action,
	{_State175, EnumToken}:                        _GotoState74Action,
	{_State175, TraitToken}:                       _GotoState82Action,
	{_State175, FuncToken}:                        _GotoState76Action,
	{_State175, LparenToken}:                      _GotoState79Action,
	{_State175, LbracketToken}:                    _GotoState27Action,
	{_State175, DotToken}:                         _GotoState73Action,
	{_State175, QuestionToken}:                    _GotoState80Action,
	{_State175, ExclaimToken}:                     _GotoState75Action,
	{_State175, TildeTildeToken}:                  _GotoState81Action,
	{_State175, BitNegToken}:                      _GotoState72Action,
	{_State175, BitAndToken}:                      _GotoState71Action,
	{_State175, LexErrorToken}:                    _GotoState78Action,
	{_State175, InitializableTypeType}:            _GotoState88Action,
	{_State175, AtomTypeType}:                     _GotoState83Action,
	{_State175, ReturnableTypeType}:               _GotoState89Action,
	{_State175, ValueTypeType}:                    _GotoState258Action,
	{_State175, ImplicitStructDefType}:            _GotoState87Action,
	{_State175, ExplicitStructDefType}:            _GotoState46Action,
	{_State175, ImplicitEnumDefType}:              _GotoState86Action,
	{_State175, ExplicitEnumDefType}:              _GotoState84Action,
	{_State175, TraitDefType}:                     _GotoState90Action,
	{_State175, FuncTypeType}:                     _GotoState85Action,
	{_State176, IntegerLiteralToken}:              _GotoState259Action,
	{_State177, IdentifierToken}:                  _GotoState77Action,
	{_State177, StructToken}:                      _GotoState34Action,
	{_State177, EnumToken}:                        _GotoState74Action,
	{_State177, TraitToken}:                       _GotoState82Action,
	{_State177, FuncToken}:                        _GotoState76Action,
	{_State177, LparenToken}:                      _GotoState79Action,
	{_State177, LbracketToken}:                    _GotoState27Action,
	{_State177, DotToken}:                         _GotoState73Action,
	{_State177, QuestionToken}:                    _GotoState80Action,
	{_State177, ExclaimToken}:                     _GotoState75Action,
	{_State177, TildeTildeToken}:                  _GotoState81Action,
	{_State177, BitNegToken}:                      _GotoState72Action,
	{_State177, BitAndToken}:                      _GotoState71Action,
	{_State177, LexErrorToken}:                    _GotoState78Action,
	{_State177, InitializableTypeType}:            _GotoState88Action,
	{_State177, AtomTypeType}:                     _GotoState83Action,
	{_State177, ReturnableTypeType}:               _GotoState260Action,
	{_State177, ImplicitStructDefType}:            _GotoState87Action,
	{_State177, ExplicitStructDefType}:            _GotoState46Action,
	{_State177, ImplicitEnumDefType}:              _GotoState86Action,
	{_State177, ExplicitEnumDefType}:              _GotoState84Action,
	{_State177, TraitDefType}:                     _GotoState90Action,
	{_State177, FuncTypeType}:                     _GotoState85Action,
	{_State179, IdentifierToken}:                  _GotoState77Action,
	{_State179, StructToken}:                      _GotoState34Action,
	{_State179, EnumToken}:                        _GotoState74Action,
	{_State179, TraitToken}:                       _GotoState82Action,
	{_State179, FuncToken}:                        _GotoState76Action,
	{_State179, LparenToken}:                      _GotoState79Action,
	{_State179, LbracketToken}:                    _GotoState27Action,
	{_State179, DotToken}:                         _GotoState73Action,
	{_State179, QuestionToken}:                    _GotoState80Action,
	{_State179, ExclaimToken}:                     _GotoState75Action,
	{_State179, TildeTildeToken}:                  _GotoState81Action,
	{_State179, BitNegToken}:                      _GotoState72Action,
	{_State179, BitAndToken}:                      _GotoState71Action,
	{_State179, LexErrorToken}:                    _GotoState78Action,
	{_State179, InitializableTypeType}:            _GotoState88Action,
	{_State179, AtomTypeType}:                     _GotoState83Action,
	{_State179, ReturnableTypeType}:               _GotoState261Action,
	{_State179, ImplicitStructDefType}:            _GotoState87Action,
	{_State179, ExplicitStructDefType}:            _GotoState46Action,
	{_State179, ImplicitEnumDefType}:              _GotoState86Action,
	{_State179, ExplicitEnumDefType}:              _GotoState84Action,
	{_State179, TraitDefType}:                     _GotoState90Action,
	{_State179, FuncTypeType}:                     _GotoState85Action,
	{_State180, IntegerLiteralToken}:              _GotoState25Action,
	{_State180, FloatLiteralToken}:                _GotoState22Action,
	{_State180, RuneLiteralToken}:                 _GotoState32Action,
	{_State180, StringLiteralToken}:               _GotoState33Action,
	{_State180, IdentifierToken}:                  _GotoState24Action,
	{_State180, TrueToken}:                        _GotoState36Action,
	{_State180, FalseToken}:                       _GotoState21Action,
	{_State180, StructToken}:                      _GotoState34Action,
	{_State180, FuncToken}:                        _GotoState23Action,
	{_State180, VarToken}:                         _GotoState37Action,
	{_State180, LabelDeclToken}:                   _GotoState26Action,
	{_State180, LparenToken}:                      _GotoState29Action,
	{_State180, LbracketToken}:                    _GotoState27Action,
	{_State180, NotToken}:                         _GotoState31Action,
	{_State180, SubToken}:                         _GotoState35Action,
	{_State180, MulToken}:                         _GotoState30Action,
	{_State180, BitNegToken}:                      _GotoState20Action,
	{_State180, BitAndToken}:                      _GotoState19Action,
	{_State180, LexErrorToken}:                    _GotoState28Action,
	{_State180, ExpressionType}:                   _GotoState262Action,
	{_State180, OptionalLabelDeclType}:            _GotoState51Action,
	{_State180, SequenceExprType}:                 _GotoState56Action,
	{_State180, BlockExprType}:                    _GotoState43Action,
	{_State180, CallExprType}:                     _GotoState44Action,
	{_State180, AtomExprType}:                     _GotoState42Action,
	{_State180, LiteralType}:                      _GotoState49Action,
	{_State180, ImplicitStructExprType}:           _GotoState47Action,
	{_State180, AccessExprType}:                   _GotoState38Action,
	{_State180, PostfixUnaryExprType}:             _GotoState53Action,
	{_State180, PrefixUnaryOpType}:                _GotoState55Action,
	{_State180, PrefixUnaryExprType}:              _GotoState54Action,
	{_State180, MulExprType}:                      _GotoState50Action,
	{_State180, AddExprType}:                      _GotoState39Action,
	{_State180, CmpExprType}:                      _GotoState45Action,
	{_State180, AndExprType}:                      _GotoState40Action,
	{_State180, OrExprType}:                       _GotoState52Action,
	{_State180, InitializableTypeType}:            _GotoState48Action,
	{_State180, ExplicitStructDefType}:            _GotoState46Action,
	{_State180, AnonymousFuncExprType}:            _GotoState41Action,
	{_State181, IntegerLiteralToken}:              _GotoState25Action,
	{_State181, FloatLiteralToken}:                _GotoState22Action,
	{_State181, RuneLiteralToken}:                 _GotoState32Action,
	{_State181, StringLiteralToken}:               _GotoState33Action,
	{_State181, IdentifierToken}:                  _GotoState93Action,
	{_State181, TrueToken}:                        _GotoState36Action,
	{_State181, FalseToken}:                       _GotoState21Action,
	{_State181, StructToken}:                      _GotoState34Action,
	{_State181, FuncToken}:                        _GotoState23Action,
	{_State181, VarToken}:                         _GotoState37Action,
	{_State181, LabelDeclToken}:                   _GotoState26Action,
	{_State181, LparenToken}:                      _GotoState29Action,
	{_State181, LbracketToken}:                    _GotoState27Action,
	{_State181, DotdotdotToken}:                   _GotoState92Action,
	{_State181, NotToken}:                         _GotoState31Action,
	{_State181, SubToken}:                         _GotoState35Action,
	{_State181, MulToken}:                         _GotoState30Action,
	{_State181, BitNegToken}:                      _GotoState20Action,
	{_State181, BitAndToken}:                      _GotoState19Action,
	{_State181, LexErrorToken}:                    _GotoState28Action,
	{_State181, ExpressionType}:                   _GotoState97Action,
	{_State181, OptionalLabelDeclType}:            _GotoState51Action,
	{_State181, SequenceExprType}:                 _GotoState56Action,
	{_State181, BlockExprType}:                    _GotoState43Action,
	{_State181, CallExprType}:                     _GotoState44Action,
	{_State181, ArgumentType}:                     _GotoState263Action,
	{_State181, ColonExpressionsType}:             _GotoState96Action,
	{_State181, OptionalExpressionType}:           _GotoState98Action,
	{_State181, AtomExprType}:                     _GotoState42Action,
	{_State181, LiteralType}:                      _GotoState49Action,
	{_State181, ImplicitStructExprType}:           _GotoState47Action,
	{_State181, AccessExprType}:                   _GotoState38Action,
	{_State181, PostfixUnaryExprType}:             _GotoState53Action,
	{_State181, PrefixUnaryOpType}:                _GotoState55Action,
	{_State181, PrefixUnaryExprType}:              _GotoState54Action,
	{_State181, MulExprType}:                      _GotoState50Action,
	{_State181, AddExprType}:                      _GotoState39Action,
	{_State181, CmpExprType}:                      _GotoState45Action,
	{_State181, AndExprType}:                      _GotoState40Action,
	{_State181, OrExprType}:                       _GotoState52Action,
	{_State181, InitializableTypeType}:            _GotoState48Action,
	{_State181, ExplicitStructDefType}:            _GotoState46Action,
	{_State181, AnonymousFuncExprType}:            _GotoState41Action,
	{_State183, IntegerLiteralToken}:              _GotoState25Action,
	{_State183, FloatLiteralToken}:                _GotoState22Action,
	{_State183, RuneLiteralToken}:                 _GotoState32Action,
	{_State183, StringLiteralToken}:               _GotoState33Action,
	{_State183, IdentifierToken}:                  _GotoState24Action,
	{_State183, TrueToken}:                        _GotoState36Action,
	{_State183, FalseToken}:                       _GotoState21Action,
	{_State183, StructToken}:                      _GotoState34Action,
	{_State183, FuncToken}:                        _GotoState23Action,
	{_State183, VarToken}:                         _GotoState37Action,
	{_State183, LabelDeclToken}:                   _GotoState26Action,
	{_State183, LparenToken}:                      _GotoState29Action,
	{_State183, LbracketToken}:                    _GotoState27Action,
	{_State183, NotToken}:                         _GotoState31Action,
	{_State183, SubToken}:                         _GotoState35Action,
	{_State183, MulToken}:                         _GotoState30Action,
	{_State183, BitNegToken}:                      _GotoState20Action,
	{_State183, BitAndToken}:                      _GotoState19Action,
	{_State183, LexErrorToken}:                    _GotoState28Action,
	{_State183, ExpressionType}:                   _GotoState264Action,
	{_State183, OptionalLabelDeclType}:            _GotoState51Action,
	{_State183, SequenceExprType}:                 _GotoState56Action,
	{_State183, BlockExprType}:                    _GotoState43Action,
	{_State183, CallExprType}:                     _GotoState44Action,
	{_State183, OptionalExpressionType}:           _GotoState265Action,
	{_State183, AtomExprType}:                     _GotoState42Action,
	{_State183, LiteralType}:                      _GotoState49Action,
	{_State183, ImplicitStructExprType}:           _GotoState47Action,
	{_State183, AccessExprType}:                   _GotoState38Action,
	{_State183, PostfixUnaryExprType}:             _GotoState53Action,
	{_State183, PrefixUnaryOpType}:                _GotoState55Action,
	{_State183, PrefixUnaryExprType}:              _GotoState54Action,
	{_State183, MulExprType}:                      _GotoState50Action,
	{_State183, AddExprType}:                      _GotoState39Action,
	{_State183, CmpExprType}:                      _GotoState45Action,
	{_State183, AndExprType}:                      _GotoState40Action,
	{_State183, OrExprType}:                       _GotoState52Action,
	{_State183, InitializableTypeType}:            _GotoState48Action,
	{_State183, ExplicitStructDefType}:            _GotoState46Action,
	{_State183, AnonymousFuncExprType}:            _GotoState41Action,
	{_State184, IntegerLiteralToken}:              _GotoState25Action,
	{_State184, FloatLiteralToken}:                _GotoState22Action,
	{_State184, RuneLiteralToken}:                 _GotoState32Action,
	{_State184, StringLiteralToken}:               _GotoState33Action,
	{_State184, IdentifierToken}:                  _GotoState24Action,
	{_State184, TrueToken}:                        _GotoState36Action,
	{_State184, FalseToken}:                       _GotoState21Action,
	{_State184, StructToken}:                      _GotoState34Action,
	{_State184, FuncToken}:                        _GotoState23Action,
	{_State184, VarToken}:                         _GotoState37Action,
	{_State184, LabelDeclToken}:                   _GotoState26Action,
	{_State184, LparenToken}:                      _GotoState29Action,
	{_State184, LbracketToken}:                    _GotoState27Action,
	{_State184, NotToken}:                         _GotoState31Action,
	{_State184, SubToken}:                         _GotoState35Action,
	{_State184, MulToken}:                         _GotoState30Action,
	{_State184, BitNegToken}:                      _GotoState20Action,
	{_State184, BitAndToken}:                      _GotoState19Action,
	{_State184, LexErrorToken}:                    _GotoState28Action,
	{_State184, ExpressionType}:                   _GotoState264Action,
	{_State184, OptionalLabelDeclType}:            _GotoState51Action,
	{_State184, SequenceExprType}:                 _GotoState56Action,
	{_State184, BlockExprType}:                    _GotoState43Action,
	{_State184, CallExprType}:                     _GotoState44Action,
	{_State184, OptionalExpressionType}:           _GotoState266Action,
	{_State184, AtomExprType}:                     _GotoState42Action,
	{_State184, LiteralType}:                      _GotoState49Action,
	{_State184, ImplicitStructExprType}:           _GotoState47Action,
	{_State184, AccessExprType}:                   _GotoState38Action,
	{_State184, PostfixUnaryExprType}:             _GotoState53Action,
	{_State184, PrefixUnaryOpType}:                _GotoState55Action,
	{_State184, PrefixUnaryExprType}:              _GotoState54Action,
	{_State184, MulExprType}:                      _GotoState50Action,
	{_State184, AddExprType}:                      _GotoState39Action,
	{_State184, CmpExprType}:                      _GotoState45Action,
	{_State184, AndExprType}:                      _GotoState40Action,
	{_State184, OrExprType}:                       _GotoState52Action,
	{_State184, InitializableTypeType}:            _GotoState48Action,
	{_State184, ExplicitStructDefType}:            _GotoState46Action,
	{_State184, AnonymousFuncExprType}:            _GotoState41Action,
	{_State185, NewlinesToken}:                    _GotoState268Action,
	{_State185, CommaToken}:                       _GotoState267Action,
	{_State187, RparenToken}:                      _GotoState269Action,
	{_State189, AssignToken}:                      _GotoState270Action,
	{_State191, RparenToken}:                      _GotoState272Action,
	{_State191, CommaToken}:                       _GotoState271Action,
	{_State194, AddToken}:                         _GotoState174Action,
	{_State194, SubToken}:                         _GotoState179Action,
	{_State194, MulToken}:                         _GotoState177Action,
	{_State195, CommaToken}:                       _GotoState273Action,
	{_State196, RbracketToken}:                    _GotoState274Action,
	{_State197, AddToken}:                         _GotoState174Action,
	{_State197, SubToken}:                         _GotoState179Action,
	{_State197, MulToken}:                         _GotoState177Action,
	{_State199, RbracketToken}:                    _GotoState275Action,
	{_State200, IntegerLiteralToken}:              _GotoState25Action,
	{_State200, FloatLiteralToken}:                _GotoState22Action,
	{_State200, RuneLiteralToken}:                 _GotoState32Action,
	{_State200, StringLiteralToken}:               _GotoState33Action,
	{_State200, IdentifierToken}:                  _GotoState93Action,
	{_State200, TrueToken}:                        _GotoState36Action,
	{_State200, FalseToken}:                       _GotoState21Action,
	{_State200, StructToken}:                      _GotoState34Action,
	{_State200, FuncToken}:                        _GotoState23Action,
	{_State200, VarToken}:                         _GotoState37Action,
	{_State200, LabelDeclToken}:                   _GotoState26Action,
	{_State200, LparenToken}:                      _GotoState29Action,
	{_State200, LbracketToken}:                    _GotoState27Action,
	{_State200, DotdotdotToken}:                   _GotoState92Action,
	{_State200, NotToken}:                         _GotoState31Action,
	{_State200, SubToken}:                         _GotoState35Action,
	{_State200, MulToken}:                         _GotoState30Action,
	{_State200, BitNegToken}:                      _GotoState20Action,
	{_State200, BitAndToken}:                      _GotoState19Action,
	{_State200, LexErrorToken}:                    _GotoState28Action,
	{_State200, ExpressionType}:                   _GotoState97Action,
	{_State200, OptionalLabelDeclType}:            _GotoState51Action,
	{_State200, SequenceExprType}:                 _GotoState56Action,
	{_State200, BlockExprType}:                    _GotoState43Action,
	{_State200, CallExprType}:                     _GotoState44Action,
	{_State200, OptionalArgumentsType}:            _GotoState277Action,
	{_State200, ArgumentsType}:                    _GotoState276Action,
	{_State200, ArgumentType}:                     _GotoState94Action,
	{_State200, ColonExpressionsType}:             _GotoState96Action,
	{_State200, OptionalExpressionType}:           _GotoState98Action,
	{_State200, AtomExprType}:                     _GotoState42Action,
	{_State200, LiteralType}:                      _GotoState49Action,
	{_State200, ImplicitStructExprType}:           _GotoState47Action,
	{_State200, AccessExprType}:                   _GotoState38Action,
	{_State200, PostfixUnaryExprType}:             _GotoState53Action,
	{_State200, PrefixUnaryOpType}:                _GotoState55Action,
	{_State200, PrefixUnaryExprType}:              _GotoState54Action,
	{_State200, MulExprType}:                      _GotoState50Action,
	{_State200, AddExprType}:                      _GotoState39Action,
	{_State200, CmpExprType}:                      _GotoState45Action,
	{_State200, AndExprType}:                      _GotoState40Action,
	{_State200, OrExprType}:                       _GotoState52Action,
	{_State200, InitializableTypeType}:            _GotoState48Action,
	{_State200, ExplicitStructDefType}:            _GotoState46Action,
	{_State200, AnonymousFuncExprType}:            _GotoState41Action,
	{_State201, MulToken}:                         _GotoState128Action,
	{_State201, DivToken}:                         _GotoState126Action,
	{_State201, ModToken}:                         _GotoState127Action,
	{_State201, BitAndToken}:                      _GotoState123Action,
	{_State201, BitLshiftToken}:                   _GotoState124Action,
	{_State201, BitRshiftToken}:                   _GotoState125Action,
	{_State201, MulOpType}:                        _GotoState129Action,
	{_State202, EqualToken}:                       _GotoState115Action,
	{_State202, NotEqualToken}:                    _GotoState120Action,
	{_State202, LessToken}:                        _GotoState118Action,
	{_State202, LessOrEqualToken}:                 _GotoState119Action,
	{_State202, GreaterToken}:                     _GotoState116Action,
	{_State202, GreaterOrEqualToken}:              _GotoState117Action,
	{_State202, CmpOpType}:                        _GotoState121Action,
	{_State203, AddToken}:                         _GotoState109Action,
	{_State203, SubToken}:                         _GotoState112Action,
	{_State203, BitXorToken}:                      _GotoState111Action,
	{_State203, BitOrToken}:                       _GotoState110Action,
	{_State203, AddOpType}:                        _GotoState113Action,
	{_State204, RparenToken}:                      _GotoState278Action,
	{_State204, CommaToken}:                       _GotoState181Action,
	{_State206, ForToken}:                         _GotoState279Action,
	{_State207, IdentifierToken}:                  _GotoState280Action,
	{_State208, IntegerLiteralToken}:              _GotoState25Action,
	{_State208, FloatLiteralToken}:                _GotoState22Action,
	{_State208, RuneLiteralToken}:                 _GotoState32Action,
	{_State208, StringLiteralToken}:               _GotoState33Action,
	{_State208, IdentifierToken}:                  _GotoState24Action,
	{_State208, TrueToken}:                        _GotoState36Action,
	{_State208, FalseToken}:                       _GotoState21Action,
	{_State208, StructToken}:                      _GotoState34Action,
	{_State208, FuncToken}:                        _GotoState23Action,
	{_State208, LabelDeclToken}:                   _GotoState26Action,
	{_State208, LparenToken}:                      _GotoState29Action,
	{_State208, LbracketToken}:                    _GotoState27Action,
	{_State208, NotToken}:                         _GotoState31Action,
	{_State208, SubToken}:                         _GotoState35Action,
	{_State208, MulToken}:                         _GotoState30Action,
	{_State208, BitNegToken}:                      _GotoState20Action,
	{_State208, BitAndToken}:                      _GotoState19Action,
	{_State208, LexErrorToken}:                    _GotoState28Action,
	{_State208, OptionalLabelDeclType}:            _GotoState140Action,
	{_State208, OptionalSequenceExprType}:         _GotoState281Action,
	{_State208, SequenceExprType}:                 _GotoState282Action,
	{_State208, BlockExprType}:                    _GotoState43Action,
	{_State208, CallExprType}:                     _GotoState44Action,
	{_State208, AtomExprType}:                     _GotoState42Action,
	{_State208, LiteralType}:                      _GotoState49Action,
	{_State208, ImplicitStructExprType}:           _GotoState47Action,
	{_State208, AccessExprType}:                   _GotoState38Action,
	{_State208, PostfixUnaryExprType}:             _GotoState53Action,
	{_State208, PrefixUnaryOpType}:                _GotoState55Action,
	{_State208, PrefixUnaryExprType}:              _GotoState54Action,
	{_State208, MulExprType}:                      _GotoState50Action,
	{_State208, AddExprType}:                      _GotoState39Action,
	{_State208, CmpExprType}:                      _GotoState45Action,
	{_State208, AndExprType}:                      _GotoState40Action,
	{_State208, OrExprType}:                       _GotoState52Action,
	{_State208, InitializableTypeType}:            _GotoState48Action,
	{_State208, ExplicitStructDefType}:            _GotoState46Action,
	{_State208, AnonymousFuncExprType}:            _GotoState41Action,
	{_State209, IdentifierToken}:                  _GotoState100Action,
	{_State209, LparenToken}:                      _GotoState101Action,
	{_State209, DotToken}:                         _GotoState283Action,
	{_State209, VarPatternType}:                   _GotoState103Action,
	{_State209, TuplePatternType}:                 _GotoState102Action,
	{_State211, InToken}:                          _GotoState284Action,
	{_State212, DoToken}:                          _GotoState285Action,
	{_State213, LbraceToken}:                      _GotoState133Action,
	{_State213, BlockBodyType}:                    _GotoState286Action,
	{_State214, IntegerLiteralToken}:              _GotoState25Action,
	{_State214, FloatLiteralToken}:                _GotoState22Action,
	{_State214, RuneLiteralToken}:                 _GotoState32Action,
	{_State214, StringLiteralToken}:               _GotoState33Action,
	{_State214, IdentifierToken}:                  _GotoState24Action,
	{_State214, TrueToken}:                        _GotoState36Action,
	{_State214, FalseToken}:                       _GotoState21Action,
	{_State214, ReturnToken}:                      _GotoState292Action,
	{_State214, BreakToken}:                       _GotoState288Action,
	{_State214, ContinueToken}:                    _GotoState289Action,
	{_State214, UnsafeToken}:                      _GotoState59Action,
	{_State214, StructToken}:                      _GotoState34Action,
	{_State214, FuncToken}:                        _GotoState23Action,
	{_State214, AsyncToken}:                       _GotoState287Action,
	{_State214, DeferToken}:                       _GotoState290Action,
	{_State214, VarToken}:                         _GotoState209Action,
	{_State214, LabelDeclToken}:                   _GotoState26Action,
	{_State214, RbraceToken}:                      _GotoState291Action,
	{_State214, LparenToken}:                      _GotoState29Action,
	{_State214, LbracketToken}:                    _GotoState27Action,
	{_State214, DotToken}:                         _GotoState207Action,
	{_State214, NotToken}:                         _GotoState31Action,
	{_State214, SubToken}:                         _GotoState35Action,
	{_State214, MulToken}:                         _GotoState30Action,
	{_State214, BitNegToken}:                      _GotoState20Action,
	{_State214, BitAndToken}:                      _GotoState19Action,
	{_State214, LexErrorToken}:                    _GotoState28Action,
	{_State214, ExpressionType}:                   _GotoState294Action,
	{_State214, OptionalLabelDeclType}:            _GotoState51Action,
	{_State214, SequenceExprType}:                 _GotoState56Action,
	{_State214, BlockExprType}:                    _GotoState43Action,
	{_State214, StatementType}:                    _GotoState299Action,
	{_State214, StatementBodyType}:                _GotoState300Action,
	{_State214, PatternType}:                      _GotoState298Action,
	{_State214, UnsafeStatementType}:              _GotoState301Action,
	{_State214, JumpStatementType}:                _GotoState296Action,
	{_State214, JumpTypeType}:                     _GotoState297Action,
	{_State214, ExpressionsType}:                  _GotoState295Action,
	{_State214, CallExprType}:                     _GotoState44Action,
	{_State214, AtomExprType}:                     _GotoState42Action,
	{_State214, LiteralType}:                      _GotoState49Action,
	{_State214, ImplicitStructExprType}:           _GotoState47Action,
	{_State214, AccessExprType}:                   _GotoState293Action,
	{_State214, PostfixUnaryExprType}:             _GotoState53Action,
	{_State214, PrefixUnaryOpType}:                _GotoState55Action,
	{_State214, PrefixUnaryExprType}:              _GotoState54Action,
	{_State214, MulExprType}:                      _GotoState50Action,
	{_State214, AddExprType}:                      _GotoState39Action,
	{_State214, CmpExprType}:                      _GotoState45Action,
	{_State214, AndExprType}:                      _GotoState40Action,
	{_State214, OrExprType}:                       _GotoState52Action,
	{_State214, InitializableTypeType}:            _GotoState48Action,
	{_State214, ExplicitStructDefType}:            _GotoState46Action,
	{_State214, AnonymousFuncExprType}:            _GotoState41Action,
	{_State215, LbraceToken}:                      _GotoState302Action,
	{_State216, AndToken}:                         _GotoState114Action,
	{_State217, GreaterToken}:                     _GotoState303Action,
	{_State219, UnsafeToken}:                      _GotoState59Action,
	{_State219, RparenToken}:                      _GotoState304Action,
	{_State219, UnsafeStatementType}:              _GotoState307Action,
	{_State219, PackageStatementBodyType}:         _GotoState306Action,
	{_State219, PackageStatementType}:             _GotoState305Action,
	{_State220, IdentifierToken}:                  _GotoState77Action,
	{_State220, StructToken}:                      _GotoState34Action,
	{_State220, EnumToken}:                        _GotoState74Action,
	{_State220, TraitToken}:                       _GotoState82Action,
	{_State220, FuncToken}:                        _GotoState76Action,
	{_State220, LparenToken}:                      _GotoState79Action,
	{_State220, LbracketToken}:                    _GotoState27Action,
	{_State220, DotToken}:                         _GotoState73Action,
	{_State220, QuestionToken}:                    _GotoState80Action,
	{_State220, ExclaimToken}:                     _GotoState75Action,
	{_State220, TildeTildeToken}:                  _GotoState81Action,
	{_State220, BitNegToken}:                      _GotoState72Action,
	{_State220, BitAndToken}:                      _GotoState71Action,
	{_State220, LexErrorToken}:                    _GotoState78Action,
	{_State220, InitializableTypeType}:            _GotoState88Action,
	{_State220, AtomTypeType}:                     _GotoState83Action,
	{_State220, ReturnableTypeType}:               _GotoState89Action,
	{_State220, ValueTypeType}:                    _GotoState308Action,
	{_State220, ImplicitStructDefType}:            _GotoState87Action,
	{_State220, ExplicitStructDefType}:            _GotoState46Action,
	{_State220, ImplicitEnumDefType}:              _GotoState86Action,
	{_State220, ExplicitEnumDefType}:              _GotoState84Action,
	{_State220, TraitDefType}:                     _GotoState90Action,
	{_State220, FuncTypeType}:                     _GotoState85Action,
	{_State222, CommaToken}:                       _GotoState309Action,
	{_State223, RbracketToken}:                    _GotoState310Action,
	{_State224, AddToken}:                         _GotoState174Action,
	{_State224, SubToken}:                         _GotoState179Action,
	{_State224, MulToken}:                         _GotoState177Action,
	{_State225, ImplementsToken}:                  _GotoState311Action,
	{_State225, AddToken}:                         _GotoState174Action,
	{_State225, SubToken}:                         _GotoState179Action,
	{_State225, MulToken}:                         _GotoState177Action,
	{_State226, IdentifierToken}:                  _GotoState77Action,
	{_State226, StructToken}:                      _GotoState34Action,
	{_State226, EnumToken}:                        _GotoState74Action,
	{_State226, TraitToken}:                       _GotoState82Action,
	{_State226, FuncToken}:                        _GotoState76Action,
	{_State226, LparenToken}:                      _GotoState79Action,
	{_State226, LbracketToken}:                    _GotoState27Action,
	{_State226, DotToken}:                         _GotoState73Action,
	{_State226, QuestionToken}:                    _GotoState80Action,
	{_State226, ExclaimToken}:                     _GotoState75Action,
	{_State226, TildeTildeToken}:                  _GotoState81Action,
	{_State226, BitNegToken}:                      _GotoState72Action,
	{_State226, BitAndToken}:                      _GotoState71Action,
	{_State226, LexErrorToken}:                    _GotoState78Action,
	{_State226, InitializableTypeType}:            _GotoState88Action,
	{_State226, AtomTypeType}:                     _GotoState83Action,
	{_State226, ReturnableTypeType}:               _GotoState89Action,
	{_State226, ValueTypeType}:                    _GotoState312Action,
	{_State226, ImplicitStructDefType}:            _GotoState87Action,
	{_State226, ExplicitStructDefType}:            _GotoState46Action,
	{_State226, ImplicitEnumDefType}:              _GotoState86Action,
	{_State226, ExplicitEnumDefType}:              _GotoState84Action,
	{_State226, TraitDefType}:                     _GotoState90Action,
	{_State226, FuncTypeType}:                     _GotoState85Action,
	{_State227, AddToken}:                         _GotoState174Action,
	{_State227, SubToken}:                         _GotoState179Action,
	{_State227, MulToken}:                         _GotoState177Action,
	{_State229, LparenToken}:                      _GotoState313Action,
	{_State230, IdentifierToken}:                  _GotoState77Action,
	{_State230, StructToken}:                      _GotoState34Action,
	{_State230, EnumToken}:                        _GotoState74Action,
	{_State230, TraitToken}:                       _GotoState82Action,
	{_State230, FuncToken}:                        _GotoState76Action,
	{_State230, LparenToken}:                      _GotoState79Action,
	{_State230, LbracketToken}:                    _GotoState27Action,
	{_State230, DotToken}:                         _GotoState73Action,
	{_State230, QuestionToken}:                    _GotoState80Action,
	{_State230, ExclaimToken}:                     _GotoState75Action,
	{_State230, TildeTildeToken}:                  _GotoState81Action,
	{_State230, BitNegToken}:                      _GotoState72Action,
	{_State230, BitAndToken}:                      _GotoState71Action,
	{_State230, LexErrorToken}:                    _GotoState78Action,
	{_State230, InitializableTypeType}:            _GotoState88Action,
	{_State230, AtomTypeType}:                     _GotoState83Action,
	{_State230, ReturnableTypeType}:               _GotoState315Action,
	{_State230, ImplicitStructDefType}:            _GotoState87Action,
	{_State230, ExplicitStructDefType}:            _GotoState46Action,
	{_State230, ImplicitEnumDefType}:              _GotoState86Action,
	{_State230, ExplicitEnumDefType}:              _GotoState84Action,
	{_State230, TraitDefType}:                     _GotoState90Action,
	{_State230, ReturnTypeType}:                   _GotoState314Action,
	{_State230, FuncTypeType}:                     _GotoState85Action,
	{_State231, IdentifierToken}:                  _GotoState149Action,
	{_State231, ParameterDefType}:                 _GotoState316Action,
	{_State232, NewlinesToken}:                    _GotoState317Action,
	{_State232, OrToken}:                          _GotoState318Action,
	{_State233, RparenToken}:                      _GotoState319Action,
	{_State234, AssignToken}:                      _GotoState246Action,
	{_State235, NewlinesToken}:                    _GotoState320Action,
	{_State235, OrToken}:                          _GotoState321Action,
	{_State236, IdentifierToken}:                  _GotoState77Action,
	{_State236, StructToken}:                      _GotoState34Action,
	{_State236, EnumToken}:                        _GotoState74Action,
	{_State236, TraitToken}:                       _GotoState82Action,
	{_State236, FuncToken}:                        _GotoState76Action,
	{_State236, LparenToken}:                      _GotoState79Action,
	{_State236, LbracketToken}:                    _GotoState27Action,
	{_State236, DotToken}:                         _GotoState73Action,
	{_State236, QuestionToken}:                    _GotoState80Action,
	{_State236, ExclaimToken}:                     _GotoState75Action,
	{_State236, TildeTildeToken}:                  _GotoState81Action,
	{_State236, BitNegToken}:                      _GotoState72Action,
	{_State236, BitAndToken}:                      _GotoState71Action,
	{_State236, LexErrorToken}:                    _GotoState78Action,
	{_State236, InitializableTypeType}:            _GotoState88Action,
	{_State236, AtomTypeType}:                     _GotoState83Action,
	{_State236, ReturnableTypeType}:               _GotoState89Action,
	{_State236, ValueTypeType}:                    _GotoState322Action,
	{_State236, ImplicitStructDefType}:            _GotoState87Action,
	{_State236, ExplicitStructDefType}:            _GotoState46Action,
	{_State236, ImplicitEnumDefType}:              _GotoState86Action,
	{_State236, ExplicitEnumDefType}:              _GotoState84Action,
	{_State236, TraitDefType}:                     _GotoState90Action,
	{_State236, FuncTypeType}:                     _GotoState85Action,
	{_State237, IdentifierToken}:                  _GotoState77Action,
	{_State237, StructToken}:                      _GotoState34Action,
	{_State237, EnumToken}:                        _GotoState74Action,
	{_State237, TraitToken}:                       _GotoState82Action,
	{_State237, FuncToken}:                        _GotoState76Action,
	{_State237, LparenToken}:                      _GotoState79Action,
	{_State237, LbracketToken}:                    _GotoState27Action,
	{_State237, DotToken}:                         _GotoState243Action,
	{_State237, QuestionToken}:                    _GotoState80Action,
	{_State237, ExclaimToken}:                     _GotoState75Action,
	{_State237, DollarLbracketToken}:              _GotoState104Action,
	{_State237, DotdotdotToken}:                   _GotoState323Action,
	{_State237, TildeTildeToken}:                  _GotoState81Action,
	{_State237, BitNegToken}:                      _GotoState72Action,
	{_State237, BitAndToken}:                      _GotoState71Action,
	{_State237, LexErrorToken}:                    _GotoState78Action,
	{_State237, OptionalGenericBindingType}:       _GotoState162Action,
	{_State237, InitializableTypeType}:            _GotoState88Action,
	{_State237, AtomTypeType}:                     _GotoState83Action,
	{_State237, ReturnableTypeType}:               _GotoState89Action,
	{_State237, ValueTypeType}:                    _GotoState324Action,
	{_State237, ImplicitStructDefType}:            _GotoState87Action,
	{_State237, ExplicitStructDefType}:            _GotoState46Action,
	{_State237, ImplicitEnumDefType}:              _GotoState86Action,
	{_State237, ExplicitEnumDefType}:              _GotoState84Action,
	{_State237, TraitDefType}:                     _GotoState90Action,
	{_State237, FuncTypeType}:                     _GotoState85Action,
	{_State238, RparenToken}:                      _GotoState325Action,
	{_State240, CommaToken}:                       _GotoState326Action,
	{_State241, AddToken}:                         _GotoState174Action,
	{_State241, SubToken}:                         _GotoState179Action,
	{_State241, MulToken}:                         _GotoState177Action,
	{_State242, DollarLbracketToken}:              _GotoState104Action,
	{_State242, OptionalGenericBindingType}:       _GotoState327Action,
	{_State243, IdentifierToken}:                  _GotoState242Action,
	{_State243, DollarLbracketToken}:              _GotoState104Action,
	{_State243, OptionalGenericBindingType}:       _GotoState157Action,
	{_State244, AddToken}:                         _GotoState174Action,
	{_State244, SubToken}:                         _GotoState179Action,
	{_State244, MulToken}:                         _GotoState177Action,
	{_State245, IdentifierToken}:                  _GotoState163Action,
	{_State245, UnsafeToken}:                      _GotoState59Action,
	{_State245, StructToken}:                      _GotoState34Action,
	{_State245, EnumToken}:                        _GotoState74Action,
	{_State245, TraitToken}:                       _GotoState82Action,
	{_State245, FuncToken}:                        _GotoState76Action,
	{_State245, LparenToken}:                      _GotoState79Action,
	{_State245, LbracketToken}:                    _GotoState27Action,
	{_State245, DotToken}:                         _GotoState73Action,
	{_State245, QuestionToken}:                    _GotoState80Action,
	{_State245, ExclaimToken}:                     _GotoState75Action,
	{_State245, TildeTildeToken}:                  _GotoState81Action,
	{_State245, BitNegToken}:                      _GotoState72Action,
	{_State245, BitAndToken}:                      _GotoState71Action,
	{_State245, LexErrorToken}:                    _GotoState78Action,
	{_State245, UnsafeStatementType}:              _GotoState169Action,
	{_State245, InitializableTypeType}:            _GotoState88Action,
	{_State245, AtomTypeType}:                     _GotoState83Action,
	{_State245, ReturnableTypeType}:               _GotoState89Action,
	{_State245, ValueTypeType}:                    _GotoState170Action,
	{_State245, FieldDefType}:                     _GotoState234Action,
	{_State245, ImplicitStructDefType}:            _GotoState87Action,
	{_State245, ExplicitStructDefType}:            _GotoState46Action,
	{_State245, EnumValueDefType}:                 _GotoState328Action,
	{_State245, ImplicitEnumDefType}:              _GotoState86Action,
	{_State245, ExplicitEnumDefType}:              _GotoState84Action,
	{_State245, TraitDefType}:                     _GotoState90Action,
	{_State245, FuncTypeType}:                     _GotoState85Action,
	{_State246, DefaultToken}:                     _GotoState329Action,
	{_State247, IdentifierToken}:                  _GotoState163Action,
	{_State247, UnsafeToken}:                      _GotoState59Action,
	{_State247, StructToken}:                      _GotoState34Action,
	{_State247, EnumToken}:                        _GotoState74Action,
	{_State247, TraitToken}:                       _GotoState82Action,
	{_State247, FuncToken}:                        _GotoState76Action,
	{_State247, LparenToken}:                      _GotoState79Action,
	{_State247, LbracketToken}:                    _GotoState27Action,
	{_State247, DotToken}:                         _GotoState73Action,
	{_State247, QuestionToken}:                    _GotoState80Action,
	{_State247, ExclaimToken}:                     _GotoState75Action,
	{_State247, TildeTildeToken}:                  _GotoState81Action,
	{_State247, BitNegToken}:                      _GotoState72Action,
	{_State247, BitAndToken}:                      _GotoState71Action,
	{_State247, LexErrorToken}:                    _GotoState78Action,
	{_State247, UnsafeStatementType}:              _GotoState169Action,
	{_State247, InitializableTypeType}:            _GotoState88Action,
	{_State247, AtomTypeType}:                     _GotoState83Action,
	{_State247, ReturnableTypeType}:               _GotoState89Action,
	{_State247, ValueTypeType}:                    _GotoState170Action,
	{_State247, FieldDefType}:                     _GotoState234Action,
	{_State247, ImplicitStructDefType}:            _GotoState87Action,
	{_State247, ExplicitStructDefType}:            _GotoState46Action,
	{_State247, EnumValueDefType}:                 _GotoState330Action,
	{_State247, ImplicitEnumDefType}:              _GotoState86Action,
	{_State247, ExplicitEnumDefType}:              _GotoState84Action,
	{_State247, TraitDefType}:                     _GotoState90Action,
	{_State247, FuncTypeType}:                     _GotoState85Action,
	{_State249, IdentifierToken}:                  _GotoState163Action,
	{_State249, UnsafeToken}:                      _GotoState59Action,
	{_State249, StructToken}:                      _GotoState34Action,
	{_State249, EnumToken}:                        _GotoState74Action,
	{_State249, TraitToken}:                       _GotoState82Action,
	{_State249, FuncToken}:                        _GotoState76Action,
	{_State249, LparenToken}:                      _GotoState79Action,
	{_State249, LbracketToken}:                    _GotoState27Action,
	{_State249, DotToken}:                         _GotoState73Action,
	{_State249, QuestionToken}:                    _GotoState80Action,
	{_State249, ExclaimToken}:                     _GotoState75Action,
	{_State249, TildeTildeToken}:                  _GotoState81Action,
	{_State249, BitNegToken}:                      _GotoState72Action,
	{_State249, BitAndToken}:                      _GotoState71Action,
	{_State249, LexErrorToken}:                    _GotoState78Action,
	{_State249, UnsafeStatementType}:              _GotoState169Action,
	{_State249, InitializableTypeType}:            _GotoState88Action,
	{_State249, AtomTypeType}:                     _GotoState83Action,
	{_State249, ReturnableTypeType}:               _GotoState89Action,
	{_State249, ValueTypeType}:                    _GotoState170Action,
	{_State249, FieldDefType}:                     _GotoState331Action,
	{_State249, ImplicitStructDefType}:            _GotoState87Action,
	{_State249, ExplicitStructDefType}:            _GotoState46Action,
	{_State249, ImplicitEnumDefType}:              _GotoState86Action,
	{_State249, ExplicitEnumDefType}:              _GotoState84Action,
	{_State249, TraitDefType}:                     _GotoState90Action,
	{_State249, FuncTypeType}:                     _GotoState85Action,
	{_State251, IdentifierToken}:                  _GotoState332Action,
	{_State251, LparenToken}:                      _GotoState160Action,
	{_State254, RparenToken}:                      _GotoState333Action,
	{_State255, NewlinesToken}:                    _GotoState335Action,
	{_State255, CommaToken}:                       _GotoState334Action,
	{_State258, RbracketToken}:                    _GotoState336Action,
	{_State258, AddToken}:                         _GotoState174Action,
	{_State258, SubToken}:                         _GotoState179Action,
	{_State258, MulToken}:                         _GotoState177Action,
	{_State259, RbracketToken}:                    _GotoState337Action,
	{_State267, IdentifierToken}:                  _GotoState163Action,
	{_State267, UnsafeToken}:                      _GotoState59Action,
	{_State267, StructToken}:                      _GotoState34Action,
	{_State267, EnumToken}:                        _GotoState74Action,
	{_State267, TraitToken}:                       _GotoState82Action,
	{_State267, FuncToken}:                        _GotoState76Action,
	{_State267, LparenToken}:                      _GotoState79Action,
	{_State267, LbracketToken}:                    _GotoState27Action,
	{_State267, DotToken}:                         _GotoState73Action,
	{_State267, QuestionToken}:                    _GotoState80Action,
	{_State267, ExclaimToken}:                     _GotoState75Action,
	{_State267, TildeTildeToken}:                  _GotoState81Action,
	{_State267, BitNegToken}:                      _GotoState72Action,
	{_State267, BitAndToken}:                      _GotoState71Action,
	{_State267, LexErrorToken}:                    _GotoState78Action,
	{_State267, UnsafeStatementType}:              _GotoState169Action,
	{_State267, InitializableTypeType}:            _GotoState88Action,
	{_State267, AtomTypeType}:                     _GotoState83Action,
	{_State267, ReturnableTypeType}:               _GotoState89Action,
	{_State267, ValueTypeType}:                    _GotoState170Action,
	{_State267, FieldDefType}:                     _GotoState338Action,
	{_State267, ImplicitStructDefType}:            _GotoState87Action,
	{_State267, ExplicitStructDefType}:            _GotoState46Action,
	{_State267, ImplicitEnumDefType}:              _GotoState86Action,
	{_State267, ExplicitEnumDefType}:              _GotoState84Action,
	{_State267, TraitDefType}:                     _GotoState90Action,
	{_State267, FuncTypeType}:                     _GotoState85Action,
	{_State268, IdentifierToken}:                  _GotoState163Action,
	{_State268, UnsafeToken}:                      _GotoState59Action,
	{_State268, StructToken}:                      _GotoState34Action,
	{_State268, EnumToken}:                        _GotoState74Action,
	{_State268, TraitToken}:                       _GotoState82Action,
	{_State268, FuncToken}:                        _GotoState76Action,
	{_State268, LparenToken}:                      _GotoState79Action,
	{_State268, LbracketToken}:                    _GotoState27Action,
	{_State268, DotToken}:                         _GotoState73Action,
	{_State268, QuestionToken}:                    _GotoState80Action,
	{_State268, ExclaimToken}:                     _GotoState75Action,
	{_State268, TildeTildeToken}:                  _GotoState81Action,
	{_State268, BitNegToken}:                      _GotoState72Action,
	{_State268, BitAndToken}:                      _GotoState71Action,
	{_State268, LexErrorToken}:                    _GotoState78Action,
	{_State268, UnsafeStatementType}:              _GotoState169Action,
	{_State268, InitializableTypeType}:            _GotoState88Action,
	{_State268, AtomTypeType}:                     _GotoState83Action,
	{_State268, ReturnableTypeType}:               _GotoState89Action,
	{_State268, ValueTypeType}:                    _GotoState170Action,
	{_State268, FieldDefType}:                     _GotoState339Action,
	{_State268, ImplicitStructDefType}:            _GotoState87Action,
	{_State268, ExplicitStructDefType}:            _GotoState46Action,
	{_State268, ImplicitEnumDefType}:              _GotoState86Action,
	{_State268, ExplicitEnumDefType}:              _GotoState84Action,
	{_State268, TraitDefType}:                     _GotoState90Action,
	{_State268, FuncTypeType}:                     _GotoState85Action,
	{_State270, IdentifierToken}:                  _GotoState100Action,
	{_State270, LparenToken}:                      _GotoState101Action,
	{_State270, VarPatternType}:                   _GotoState340Action,
	{_State270, TuplePatternType}:                 _GotoState102Action,
	{_State271, IdentifierToken}:                  _GotoState189Action,
	{_State271, LparenToken}:                      _GotoState101Action,
	{_State271, DotdotdotToken}:                   _GotoState188Action,
	{_State271, VarPatternType}:                   _GotoState192Action,
	{_State271, TuplePatternType}:                 _GotoState102Action,
	{_State271, FieldVarPatternType}:              _GotoState341Action,
	{_State273, IdentifierToken}:                  _GotoState77Action,
	{_State273, StructToken}:                      _GotoState34Action,
	{_State273, EnumToken}:                        _GotoState74Action,
	{_State273, TraitToken}:                       _GotoState82Action,
	{_State273, FuncToken}:                        _GotoState76Action,
	{_State273, LparenToken}:                      _GotoState79Action,
	{_State273, LbracketToken}:                    _GotoState27Action,
	{_State273, DotToken}:                         _GotoState73Action,
	{_State273, QuestionToken}:                    _GotoState80Action,
	{_State273, ExclaimToken}:                     _GotoState75Action,
	{_State273, TildeTildeToken}:                  _GotoState81Action,
	{_State273, BitNegToken}:                      _GotoState72Action,
	{_State273, BitAndToken}:                      _GotoState71Action,
	{_State273, LexErrorToken}:                    _GotoState78Action,
	{_State273, InitializableTypeType}:            _GotoState88Action,
	{_State273, AtomTypeType}:                     _GotoState83Action,
	{_State273, ReturnableTypeType}:               _GotoState89Action,
	{_State273, ValueTypeType}:                    _GotoState342Action,
	{_State273, ImplicitStructDefType}:            _GotoState87Action,
	{_State273, ExplicitStructDefType}:            _GotoState46Action,
	{_State273, ImplicitEnumDefType}:              _GotoState86Action,
	{_State273, ExplicitEnumDefType}:              _GotoState84Action,
	{_State273, TraitDefType}:                     _GotoState90Action,
	{_State273, FuncTypeType}:                     _GotoState85Action,
	{_State276, CommaToken}:                       _GotoState181Action,
	{_State277, RparenToken}:                      _GotoState343Action,
	{_State279, IntegerLiteralToken}:              _GotoState25Action,
	{_State279, FloatLiteralToken}:                _GotoState22Action,
	{_State279, RuneLiteralToken}:                 _GotoState32Action,
	{_State279, StringLiteralToken}:               _GotoState33Action,
	{_State279, IdentifierToken}:                  _GotoState24Action,
	{_State279, TrueToken}:                        _GotoState36Action,
	{_State279, FalseToken}:                       _GotoState21Action,
	{_State279, StructToken}:                      _GotoState34Action,
	{_State279, FuncToken}:                        _GotoState23Action,
	{_State279, LabelDeclToken}:                   _GotoState26Action,
	{_State279, LparenToken}:                      _GotoState29Action,
	{_State279, LbracketToken}:                    _GotoState27Action,
	{_State279, NotToken}:                         _GotoState31Action,
	{_State279, SubToken}:                         _GotoState35Action,
	{_State279, MulToken}:                         _GotoState30Action,
	{_State279, BitNegToken}:                      _GotoState20Action,
	{_State279, BitAndToken}:                      _GotoState19Action,
	{_State279, LexErrorToken}:                    _GotoState28Action,
	{_State279, OptionalLabelDeclType}:            _GotoState140Action,
	{_State279, SequenceExprType}:                 _GotoState344Action,
	{_State279, BlockExprType}:                    _GotoState43Action,
	{_State279, CallExprType}:                     _GotoState44Action,
	{_State279, AtomExprType}:                     _GotoState42Action,
	{_State279, LiteralType}:                      _GotoState49Action,
	{_State279, ImplicitStructExprType}:           _GotoState47Action,
	{_State279, AccessExprType}:                   _GotoState38Action,
	{_State279, PostfixUnaryExprType}:             _GotoState53Action,
	{_State279, PrefixUnaryOpType}:                _GotoState55Action,
	{_State279, PrefixUnaryExprType}:              _GotoState54Action,
	{_State279, MulExprType}:                      _GotoState50Action,
	{_State279, AddExprType}:                      _GotoState39Action,
	{_State279, CmpExprType}:                      _GotoState45Action,
	{_State279, AndExprType}:                      _GotoState40Action,
	{_State279, OrExprType}:                       _GotoState52Action,
	{_State279, InitializableTypeType}:            _GotoState48Action,
	{_State279, ExplicitStructDefType}:            _GotoState46Action,
	{_State279, AnonymousFuncExprType}:            _GotoState41Action,
	{_State280, LparenToken}:                      _GotoState29Action,
	{_State280, ImplicitStructExprType}:           _GotoState345Action,
	{_State281, SemicolonToken}:                   _GotoState346Action,
	{_State283, IdentifierToken}:                  _GotoState347Action,
	{_State284, IntegerLiteralToken}:              _GotoState25Action,
	{_State284, FloatLiteralToken}:                _GotoState22Action,
	{_State284, RuneLiteralToken}:                 _GotoState32Action,
	{_State284, StringLiteralToken}:               _GotoState33Action,
	{_State284, IdentifierToken}:                  _GotoState24Action,
	{_State284, TrueToken}:                        _GotoState36Action,
	{_State284, FalseToken}:                       _GotoState21Action,
	{_State284, StructToken}:                      _GotoState34Action,
	{_State284, FuncToken}:                        _GotoState23Action,
	{_State284, LabelDeclToken}:                   _GotoState26Action,
	{_State284, LparenToken}:                      _GotoState29Action,
	{_State284, LbracketToken}:                    _GotoState27Action,
	{_State284, NotToken}:                         _GotoState31Action,
	{_State284, SubToken}:                         _GotoState35Action,
	{_State284, MulToken}:                         _GotoState30Action,
	{_State284, BitNegToken}:                      _GotoState20Action,
	{_State284, BitAndToken}:                      _GotoState19Action,
	{_State284, LexErrorToken}:                    _GotoState28Action,
	{_State284, OptionalLabelDeclType}:            _GotoState140Action,
	{_State284, SequenceExprType}:                 _GotoState348Action,
	{_State284, BlockExprType}:                    _GotoState43Action,
	{_State284, CallExprType}:                     _GotoState44Action,
	{_State284, AtomExprType}:                     _GotoState42Action,
	{_State284, LiteralType}:                      _GotoState49Action,
	{_State284, ImplicitStructExprType}:           _GotoState47Action,
	{_State284, AccessExprType}:                   _GotoState38Action,
	{_State284, PostfixUnaryExprType}:             _GotoState53Action,
	{_State284, PrefixUnaryOpType}:                _GotoState55Action,
	{_State284, PrefixUnaryExprType}:              _GotoState54Action,
	{_State284, MulExprType}:                      _GotoState50Action,
	{_State284, AddExprType}:                      _GotoState39Action,
	{_State284, CmpExprType}:                      _GotoState45Action,
	{_State284, AndExprType}:                      _GotoState40Action,
	{_State284, OrExprType}:                       _GotoState52Action,
	{_State284, InitializableTypeType}:            _GotoState48Action,
	{_State284, ExplicitStructDefType}:            _GotoState46Action,
	{_State284, AnonymousFuncExprType}:            _GotoState41Action,
	{_State285, LbraceToken}:                      _GotoState133Action,
	{_State285, BlockBodyType}:                    _GotoState349Action,
	{_State286, ElseToken}:                        _GotoState350Action,
	{_State287, IntegerLiteralToken}:              _GotoState25Action,
	{_State287, FloatLiteralToken}:                _GotoState22Action,
	{_State287, RuneLiteralToken}:                 _GotoState32Action,
	{_State287, StringLiteralToken}:               _GotoState33Action,
	{_State287, IdentifierToken}:                  _GotoState24Action,
	{_State287, TrueToken}:                        _GotoState36Action,
	{_State287, FalseToken}:                       _GotoState21Action,
	{_State287, StructToken}:                      _GotoState34Action,
	{_State287, FuncToken}:                        _GotoState23Action,
	{_State287, LabelDeclToken}:                   _GotoState26Action,
	{_State287, LparenToken}:                      _GotoState29Action,
	{_State287, LbracketToken}:                    _GotoState27Action,
	{_State287, LexErrorToken}:                    _GotoState28Action,
	{_State287, OptionalLabelDeclType}:            _GotoState140Action,
	{_State287, BlockExprType}:                    _GotoState43Action,
	{_State287, CallExprType}:                     _GotoState352Action,
	{_State287, AtomExprType}:                     _GotoState42Action,
	{_State287, LiteralType}:                      _GotoState49Action,
	{_State287, ImplicitStructExprType}:           _GotoState47Action,
	{_State287, AccessExprType}:                   _GotoState351Action,
	{_State287, InitializableTypeType}:            _GotoState48Action,
	{_State287, ExplicitStructDefType}:            _GotoState46Action,
	{_State287, AnonymousFuncExprType}:            _GotoState41Action,
	{_State290, IntegerLiteralToken}:              _GotoState25Action,
	{_State290, FloatLiteralToken}:                _GotoState22Action,
	{_State290, RuneLiteralToken}:                 _GotoState32Action,
	{_State290, StringLiteralToken}:               _GotoState33Action,
	{_State290, IdentifierToken}:                  _GotoState24Action,
	{_State290, TrueToken}:                        _GotoState36Action,
	{_State290, FalseToken}:                       _GotoState21Action,
	{_State290, StructToken}:                      _GotoState34Action,
	{_State290, FuncToken}:                        _GotoState23Action,
	{_State290, LabelDeclToken}:                   _GotoState26Action,
	{_State290, LparenToken}:                      _GotoState29Action,
	{_State290, LbracketToken}:                    _GotoState27Action,
	{_State290, LexErrorToken}:                    _GotoState28Action,
	{_State290, OptionalLabelDeclType}:            _GotoState140Action,
	{_State290, BlockExprType}:                    _GotoState43Action,
	{_State290, CallExprType}:                     _GotoState353Action,
	{_State290, AtomExprType}:                     _GotoState42Action,
	{_State290, LiteralType}:                      _GotoState49Action,
	{_State290, ImplicitStructExprType}:           _GotoState47Action,
	{_State290, AccessExprType}:                   _GotoState351Action,
	{_State290, InitializableTypeType}:            _GotoState48Action,
	{_State290, ExplicitStructDefType}:            _GotoState46Action,
	{_State290, AnonymousFuncExprType}:            _GotoState41Action,
	{_State293, LbracketToken}:                    _GotoState106Action,
	{_State293, DotToken}:                         _GotoState105Action,
	{_State293, QuestionToken}:                    _GotoState107Action,
	{_State293, DollarLbracketToken}:              _GotoState104Action,
	{_State293, AddAssignToken}:                   _GotoState354Action,
	{_State293, SubAssignToken}:                   _GotoState365Action,
	{_State293, MulAssignToken}:                   _GotoState364Action,
	{_State293, DivAssignToken}:                   _GotoState362Action,
	{_State293, ModAssignToken}:                   _GotoState363Action,
	{_State293, AddOneAssignToken}:                _GotoState355Action,
	{_State293, SubOneAssignToken}:                _GotoState366Action,
	{_State293, BitNegAssignToken}:                _GotoState358Action,
	{_State293, BitAndAssignToken}:                _GotoState356Action,
	{_State293, BitOrAssignToken}:                 _GotoState359Action,
	{_State293, BitXorAssignToken}:                _GotoState361Action,
	{_State293, BitLshiftAssignToken}:             _GotoState357Action,
	{_State293, BitRshiftAssignToken}:             _GotoState360Action,
	{_State293, UnaryOpAssignType}:                _GotoState368Action,
	{_State293, BinaryOpAssignType}:               _GotoState367Action,
	{_State293, OptionalGenericBindingType}:       _GotoState108Action,
	{_State295, CommaToken}:                       _GotoState369Action,
	{_State297, JumpLabelToken}:                   _GotoState370Action,
	{_State297, OptionalJumpLabelType}:            _GotoState371Action,
	{_State298, AssignToken}:                      _GotoState372Action,
	{_State300, NewlinesToken}:                    _GotoState373Action,
	{_State300, SemicolonToken}:                   _GotoState374Action,
	{_State302, CaseToken}:                        _GotoState375Action,
	{_State302, CaseBranchesType}:                 _GotoState377Action,
	{_State302, CaseBranchType}:                   _GotoState376Action,
	{_State303, StringLiteralToken}:               _GotoState378Action,
	{_State306, NewlinesToken}:                    _GotoState379Action,
	{_State306, SemicolonToken}:                   _GotoState380Action,
	{_State308, AddToken}:                         _GotoState174Action,
	{_State308, SubToken}:                         _GotoState179Action,
	{_State308, MulToken}:                         _GotoState177Action,
	{_State309, IdentifierToken}:                  _GotoState220Action,
	{_State309, GenericParameterDefType}:          _GotoState381Action,
	{_State311, IdentifierToken}:                  _GotoState77Action,
	{_State311, StructToken}:                      _GotoState34Action,
	{_State311, EnumToken}:                        _GotoState74Action,
	{_State311, TraitToken}:                       _GotoState82Action,
	{_State311, FuncToken}:                        _GotoState76Action,
	{_State311, LparenToken}:                      _GotoState79Action,
	{_State311, LbracketToken}:                    _GotoState27Action,
	{_State311, DotToken}:                         _GotoState73Action,
	{_State311, QuestionToken}:                    _GotoState80Action,
	{_State311, ExclaimToken}:                     _GotoState75Action,
	{_State311, TildeTildeToken}:                  _GotoState81Action,
	{_State311, BitNegToken}:                      _GotoState72Action,
	{_State311, BitAndToken}:                      _GotoState71Action,
	{_State311, LexErrorToken}:                    _GotoState78Action,
	{_State311, InitializableTypeType}:            _GotoState88Action,
	{_State311, AtomTypeType}:                     _GotoState83Action,
	{_State311, ReturnableTypeType}:               _GotoState89Action,
	{_State311, ValueTypeType}:                    _GotoState382Action,
	{_State311, ImplicitStructDefType}:            _GotoState87Action,
	{_State311, ExplicitStructDefType}:            _GotoState46Action,
	{_State311, ImplicitEnumDefType}:              _GotoState86Action,
	{_State311, ExplicitEnumDefType}:              _GotoState84Action,
	{_State311, TraitDefType}:                     _GotoState90Action,
	{_State311, FuncTypeType}:                     _GotoState85Action,
	{_State312, AddToken}:                         _GotoState174Action,
	{_State312, SubToken}:                         _GotoState179Action,
	{_State312, MulToken}:                         _GotoState177Action,
	{_State313, IdentifierToken}:                  _GotoState149Action,
	{_State313, ParameterDefType}:                 _GotoState153Action,
	{_State313, ParameterDefsType}:                _GotoState154Action,
	{_State313, OptionalParameterDefsType}:        _GotoState383Action,
	{_State314, LbraceToken}:                      _GotoState133Action,
	{_State314, BlockBodyType}:                    _GotoState384Action,
	{_State317, IdentifierToken}:                  _GotoState163Action,
	{_State317, UnsafeToken}:                      _GotoState59Action,
	{_State317, StructToken}:                      _GotoState34Action,
	{_State317, EnumToken}:                        _GotoState74Action,
	{_State317, TraitToken}:                       _GotoState82Action,
	{_State317, FuncToken}:                        _GotoState76Action,
	{_State317, LparenToken}:                      _GotoState79Action,
	{_State317, LbracketToken}:                    _GotoState27Action,
	{_State317, DotToken}:                         _GotoState73Action,
	{_State317, QuestionToken}:                    _GotoState80Action,
	{_State317, ExclaimToken}:                     _GotoState75Action,
	{_State317, TildeTildeToken}:                  _GotoState81Action,
	{_State317, BitNegToken}:                      _GotoState72Action,
	{_State317, BitAndToken}:                      _GotoState71Action,
	{_State317, LexErrorToken}:                    _GotoState78Action,
	{_State317, UnsafeStatementType}:              _GotoState169Action,
	{_State317, InitializableTypeType}:            _GotoState88Action,
	{_State317, AtomTypeType}:                     _GotoState83Action,
	{_State317, ReturnableTypeType}:               _GotoState89Action,
	{_State317, ValueTypeType}:                    _GotoState170Action,
	{_State317, FieldDefType}:                     _GotoState234Action,
	{_State317, ImplicitStructDefType}:            _GotoState87Action,
	{_State317, ExplicitStructDefType}:            _GotoState46Action,
	{_State317, EnumValueDefType}:                 _GotoState385Action,
	{_State317, ImplicitEnumDefType}:              _GotoState86Action,
	{_State317, ExplicitEnumDefType}:              _GotoState84Action,
	{_State317, TraitDefType}:                     _GotoState90Action,
	{_State317, FuncTypeType}:                     _GotoState85Action,
	{_State318, IdentifierToken}:                  _GotoState163Action,
	{_State318, UnsafeToken}:                      _GotoState59Action,
	{_State318, StructToken}:                      _GotoState34Action,
	{_State318, EnumToken}:                        _GotoState74Action,
	{_State318, TraitToken}:                       _GotoState82Action,
	{_State318, FuncToken}:                        _GotoState76Action,
	{_State318, LparenToken}:                      _GotoState79Action,
	{_State318, LbracketToken}:                    _GotoState27Action,
	{_State318, DotToken}:                         _GotoState73Action,
	{_State318, QuestionToken}:                    _GotoState80Action,
	{_State318, ExclaimToken}:                     _GotoState75Action,
	{_State318, TildeTildeToken}:                  _GotoState81Action,
	{_State318, BitNegToken}:                      _GotoState72Action,
	{_State318, BitAndToken}:                      _GotoState71Action,
	{_State318, LexErrorToken}:                    _GotoState78Action,
	{_State318, UnsafeStatementType}:              _GotoState169Action,
	{_State318, InitializableTypeType}:            _GotoState88Action,
	{_State318, AtomTypeType}:                     _GotoState83Action,
	{_State318, ReturnableTypeType}:               _GotoState89Action,
	{_State318, ValueTypeType}:                    _GotoState170Action,
	{_State318, FieldDefType}:                     _GotoState234Action,
	{_State318, ImplicitStructDefType}:            _GotoState87Action,
	{_State318, ExplicitStructDefType}:            _GotoState46Action,
	{_State318, EnumValueDefType}:                 _GotoState386Action,
	{_State318, ImplicitEnumDefType}:              _GotoState86Action,
	{_State318, ExplicitEnumDefType}:              _GotoState84Action,
	{_State318, TraitDefType}:                     _GotoState90Action,
	{_State318, FuncTypeType}:                     _GotoState85Action,
	{_State320, IdentifierToken}:                  _GotoState163Action,
	{_State320, UnsafeToken}:                      _GotoState59Action,
	{_State320, StructToken}:                      _GotoState34Action,
	{_State320, EnumToken}:                        _GotoState74Action,
	{_State320, TraitToken}:                       _GotoState82Action,
	{_State320, FuncToken}:                        _GotoState76Action,
	{_State320, LparenToken}:                      _GotoState79Action,
	{_State320, LbracketToken}:                    _GotoState27Action,
	{_State320, DotToken}:                         _GotoState73Action,
	{_State320, QuestionToken}:                    _GotoState80Action,
	{_State320, ExclaimToken}:                     _GotoState75Action,
	{_State320, TildeTildeToken}:                  _GotoState81Action,
	{_State320, BitNegToken}:                      _GotoState72Action,
	{_State320, BitAndToken}:                      _GotoState71Action,
	{_State320, LexErrorToken}:                    _GotoState78Action,
	{_State320, UnsafeStatementType}:              _GotoState169Action,
	{_State320, InitializableTypeType}:            _GotoState88Action,
	{_State320, AtomTypeType}:                     _GotoState83Action,
	{_State320, ReturnableTypeType}:               _GotoState89Action,
	{_State320, ValueTypeType}:                    _GotoState170Action,
	{_State320, FieldDefType}:                     _GotoState234Action,
	{_State320, ImplicitStructDefType}:            _GotoState87Action,
	{_State320, ExplicitStructDefType}:            _GotoState46Action,
	{_State320, EnumValueDefType}:                 _GotoState387Action,
	{_State320, ImplicitEnumDefType}:              _GotoState86Action,
	{_State320, ExplicitEnumDefType}:              _GotoState84Action,
	{_State320, TraitDefType}:                     _GotoState90Action,
	{_State320, FuncTypeType}:                     _GotoState85Action,
	{_State321, IdentifierToken}:                  _GotoState163Action,
	{_State321, UnsafeToken}:                      _GotoState59Action,
	{_State321, StructToken}:                      _GotoState34Action,
	{_State321, EnumToken}:                        _GotoState74Action,
	{_State321, TraitToken}:                       _GotoState82Action,
	{_State321, FuncToken}:                        _GotoState76Action,
	{_State321, LparenToken}:                      _GotoState79Action,
	{_State321, LbracketToken}:                    _GotoState27Action,
	{_State321, DotToken}:                         _GotoState73Action,
	{_State321, QuestionToken}:                    _GotoState80Action,
	{_State321, ExclaimToken}:                     _GotoState75Action,
	{_State321, TildeTildeToken}:                  _GotoState81Action,
	{_State321, BitNegToken}:                      _GotoState72Action,
	{_State321, BitAndToken}:                      _GotoState71Action,
	{_State321, LexErrorToken}:                    _GotoState78Action,
	{_State321, UnsafeStatementType}:              _GotoState169Action,
	{_State321, InitializableTypeType}:            _GotoState88Action,
	{_State321, AtomTypeType}:                     _GotoState83Action,
	{_State321, ReturnableTypeType}:               _GotoState89Action,
	{_State321, ValueTypeType}:                    _GotoState170Action,
	{_State321, FieldDefType}:                     _GotoState234Action,
	{_State321, ImplicitStructDefType}:            _GotoState87Action,
	{_State321, ExplicitStructDefType}:            _GotoState46Action,
	{_State321, EnumValueDefType}:                 _GotoState388Action,
	{_State321, ImplicitEnumDefType}:              _GotoState86Action,
	{_State321, ExplicitEnumDefType}:              _GotoState84Action,
	{_State321, TraitDefType}:                     _GotoState90Action,
	{_State321, FuncTypeType}:                     _GotoState85Action,
	{_State322, AddToken}:                         _GotoState174Action,
	{_State322, SubToken}:                         _GotoState179Action,
	{_State322, MulToken}:                         _GotoState177Action,
	{_State323, IdentifierToken}:                  _GotoState77Action,
	{_State323, StructToken}:                      _GotoState34Action,
	{_State323, EnumToken}:                        _GotoState74Action,
	{_State323, TraitToken}:                       _GotoState82Action,
	{_State323, FuncToken}:                        _GotoState76Action,
	{_State323, LparenToken}:                      _GotoState79Action,
	{_State323, LbracketToken}:                    _GotoState27Action,
	{_State323, DotToken}:                         _GotoState73Action,
	{_State323, QuestionToken}:                    _GotoState80Action,
	{_State323, ExclaimToken}:                     _GotoState75Action,
	{_State323, TildeTildeToken}:                  _GotoState81Action,
	{_State323, BitNegToken}:                      _GotoState72Action,
	{_State323, BitAndToken}:                      _GotoState71Action,
	{_State323, LexErrorToken}:                    _GotoState78Action,
	{_State323, InitializableTypeType}:            _GotoState88Action,
	{_State323, AtomTypeType}:                     _GotoState83Action,
	{_State323, ReturnableTypeType}:               _GotoState89Action,
	{_State323, ValueTypeType}:                    _GotoState389Action,
	{_State323, ImplicitStructDefType}:            _GotoState87Action,
	{_State323, ExplicitStructDefType}:            _GotoState46Action,
	{_State323, ImplicitEnumDefType}:              _GotoState86Action,
	{_State323, ExplicitEnumDefType}:              _GotoState84Action,
	{_State323, TraitDefType}:                     _GotoState90Action,
	{_State323, FuncTypeType}:                     _GotoState85Action,
	{_State324, AddToken}:                         _GotoState174Action,
	{_State324, SubToken}:                         _GotoState179Action,
	{_State324, MulToken}:                         _GotoState177Action,
	{_State325, IdentifierToken}:                  _GotoState77Action,
	{_State325, StructToken}:                      _GotoState34Action,
	{_State325, EnumToken}:                        _GotoState74Action,
	{_State325, TraitToken}:                       _GotoState82Action,
	{_State325, FuncToken}:                        _GotoState76Action,
	{_State325, LparenToken}:                      _GotoState79Action,
	{_State325, LbracketToken}:                    _GotoState27Action,
	{_State325, DotToken}:                         _GotoState73Action,
	{_State325, QuestionToken}:                    _GotoState80Action,
	{_State325, ExclaimToken}:                     _GotoState75Action,
	{_State325, TildeTildeToken}:                  _GotoState81Action,
	{_State325, BitNegToken}:                      _GotoState72Action,
	{_State325, BitAndToken}:                      _GotoState71Action,
	{_State325, LexErrorToken}:                    _GotoState78Action,
	{_State325, InitializableTypeType}:            _GotoState88Action,
	{_State325, AtomTypeType}:                     _GotoState83Action,
	{_State325, ReturnableTypeType}:               _GotoState315Action,
	{_State325, ImplicitStructDefType}:            _GotoState87Action,
	{_State325, ExplicitStructDefType}:            _GotoState46Action,
	{_State325, ImplicitEnumDefType}:              _GotoState86Action,
	{_State325, ExplicitEnumDefType}:              _GotoState84Action,
	{_State325, TraitDefType}:                     _GotoState90Action,
	{_State325, ReturnTypeType}:                   _GotoState390Action,
	{_State325, FuncTypeType}:                     _GotoState85Action,
	{_State326, IdentifierToken}:                  _GotoState237Action,
	{_State326, StructToken}:                      _GotoState34Action,
	{_State326, EnumToken}:                        _GotoState74Action,
	{_State326, TraitToken}:                       _GotoState82Action,
	{_State326, FuncToken}:                        _GotoState76Action,
	{_State326, LparenToken}:                      _GotoState79Action,
	{_State326, LbracketToken}:                    _GotoState27Action,
	{_State326, DotToken}:                         _GotoState73Action,
	{_State326, QuestionToken}:                    _GotoState80Action,
	{_State326, ExclaimToken}:                     _GotoState75Action,
	{_State326, DotdotdotToken}:                   _GotoState236Action,
	{_State326, TildeTildeToken}:                  _GotoState81Action,
	{_State326, BitNegToken}:                      _GotoState72Action,
	{_State326, BitAndToken}:                      _GotoState71Action,
	{_State326, LexErrorToken}:                    _GotoState78Action,
	{_State326, InitializableTypeType}:            _GotoState88Action,
	{_State326, AtomTypeType}:                     _GotoState83Action,
	{_State326, ReturnableTypeType}:               _GotoState89Action,
	{_State326, ValueTypeType}:                    _GotoState241Action,
	{_State326, ImplicitStructDefType}:            _GotoState87Action,
	{_State326, ExplicitStructDefType}:            _GotoState46Action,
	{_State326, ImplicitEnumDefType}:              _GotoState86Action,
	{_State326, ExplicitEnumDefType}:              _GotoState84Action,
	{_State326, TraitDefType}:                     _GotoState90Action,
	{_State326, ParameterDeclType}:                _GotoState391Action,
	{_State326, FuncTypeType}:                     _GotoState85Action,
	{_State332, LparenToken}:                      _GotoState392Action,
	{_State334, IdentifierToken}:                  _GotoState163Action,
	{_State334, UnsafeToken}:                      _GotoState59Action,
	{_State334, StructToken}:                      _GotoState34Action,
	{_State334, EnumToken}:                        _GotoState74Action,
	{_State334, TraitToken}:                       _GotoState82Action,
	{_State334, FuncToken}:                        _GotoState251Action,
	{_State334, LparenToken}:                      _GotoState79Action,
	{_State334, LbracketToken}:                    _GotoState27Action,
	{_State334, DotToken}:                         _GotoState73Action,
	{_State334, QuestionToken}:                    _GotoState80Action,
	{_State334, ExclaimToken}:                     _GotoState75Action,
	{_State334, TildeTildeToken}:                  _GotoState81Action,
	{_State334, BitNegToken}:                      _GotoState72Action,
	{_State334, BitAndToken}:                      _GotoState71Action,
	{_State334, LexErrorToken}:                    _GotoState78Action,
	{_State334, UnsafeStatementType}:              _GotoState169Action,
	{_State334, InitializableTypeType}:            _GotoState88Action,
	{_State334, AtomTypeType}:                     _GotoState83Action,
	{_State334, ReturnableTypeType}:               _GotoState89Action,
	{_State334, ValueTypeType}:                    _GotoState170Action,
	{_State334, FieldDefType}:                     _GotoState252Action,
	{_State334, ImplicitStructDefType}:            _GotoState87Action,
	{_State334, ExplicitStructDefType}:            _GotoState46Action,
	{_State334, ImplicitEnumDefType}:              _GotoState86Action,
	{_State334, ExplicitEnumDefType}:              _GotoState84Action,
	{_State334, TraitPropertyType}:                _GotoState393Action,
	{_State334, TraitDefType}:                     _GotoState90Action,
	{_State334, FuncTypeType}:                     _GotoState85Action,
	{_State334, MethodSignatureType}:              _GotoState253Action,
	{_State335, IdentifierToken}:                  _GotoState163Action,
	{_State335, UnsafeToken}:                      _GotoState59Action,
	{_State335, StructToken}:                      _GotoState34Action,
	{_State335, EnumToken}:                        _GotoState74Action,
	{_State335, TraitToken}:                       _GotoState82Action,
	{_State335, FuncToken}:                        _GotoState251Action,
	{_State335, LparenToken}:                      _GotoState79Action,
	{_State335, LbracketToken}:                    _GotoState27Action,
	{_State335, DotToken}:                         _GotoState73Action,
	{_State335, QuestionToken}:                    _GotoState80Action,
	{_State335, ExclaimToken}:                     _GotoState75Action,
	{_State335, TildeTildeToken}:                  _GotoState81Action,
	{_State335, BitNegToken}:                      _GotoState72Action,
	{_State335, BitAndToken}:                      _GotoState71Action,
	{_State335, LexErrorToken}:                    _GotoState78Action,
	{_State335, UnsafeStatementType}:              _GotoState169Action,
	{_State335, InitializableTypeType}:            _GotoState88Action,
	{_State335, AtomTypeType}:                     _GotoState83Action,
	{_State335, ReturnableTypeType}:               _GotoState89Action,
	{_State335, ValueTypeType}:                    _GotoState170Action,
	{_State335, FieldDefType}:                     _GotoState252Action,
	{_State335, ImplicitStructDefType}:            _GotoState87Action,
	{_State335, ExplicitStructDefType}:            _GotoState46Action,
	{_State335, ImplicitEnumDefType}:              _GotoState86Action,
	{_State335, ExplicitEnumDefType}:              _GotoState84Action,
	{_State335, TraitPropertyType}:                _GotoState394Action,
	{_State335, TraitDefType}:                     _GotoState90Action,
	{_State335, FuncTypeType}:                     _GotoState85Action,
	{_State335, MethodSignatureType}:              _GotoState253Action,
	{_State342, AddToken}:                         _GotoState174Action,
	{_State342, SubToken}:                         _GotoState179Action,
	{_State342, MulToken}:                         _GotoState177Action,
	{_State346, IntegerLiteralToken}:              _GotoState25Action,
	{_State346, FloatLiteralToken}:                _GotoState22Action,
	{_State346, RuneLiteralToken}:                 _GotoState32Action,
	{_State346, StringLiteralToken}:               _GotoState33Action,
	{_State346, IdentifierToken}:                  _GotoState24Action,
	{_State346, TrueToken}:                        _GotoState36Action,
	{_State346, FalseToken}:                       _GotoState21Action,
	{_State346, StructToken}:                      _GotoState34Action,
	{_State346, FuncToken}:                        _GotoState23Action,
	{_State346, LabelDeclToken}:                   _GotoState26Action,
	{_State346, LparenToken}:                      _GotoState29Action,
	{_State346, LbracketToken}:                    _GotoState27Action,
	{_State346, NotToken}:                         _GotoState31Action,
	{_State346, SubToken}:                         _GotoState35Action,
	{_State346, MulToken}:                         _GotoState30Action,
	{_State346, BitNegToken}:                      _GotoState20Action,
	{_State346, BitAndToken}:                      _GotoState19Action,
	{_State346, LexErrorToken}:                    _GotoState28Action,
	{_State346, OptionalLabelDeclType}:            _GotoState140Action,
	{_State346, OptionalSequenceExprType}:         _GotoState395Action,
	{_State346, SequenceExprType}:                 _GotoState282Action,
	{_State346, BlockExprType}:                    _GotoState43Action,
	{_State346, CallExprType}:                     _GotoState44Action,
	{_State346, AtomExprType}:                     _GotoState42Action,
	{_State346, LiteralType}:                      _GotoState49Action,
	{_State346, ImplicitStructExprType}:           _GotoState47Action,
	{_State346, AccessExprType}:                   _GotoState38Action,
	{_State346, PostfixUnaryExprType}:             _GotoState53Action,
	{_State346, PrefixUnaryOpType}:                _GotoState55Action,
	{_State346, PrefixUnaryExprType}:              _GotoState54Action,
	{_State346, MulExprType}:                      _GotoState50Action,
	{_State346, AddExprType}:                      _GotoState39Action,
	{_State346, CmpExprType}:                      _GotoState45Action,
	{_State346, AndExprType}:                      _GotoState40Action,
	{_State346, OrExprType}:                       _GotoState52Action,
	{_State346, InitializableTypeType}:            _GotoState48Action,
	{_State346, ExplicitStructDefType}:            _GotoState46Action,
	{_State346, AnonymousFuncExprType}:            _GotoState41Action,
	{_State347, LparenToken}:                      _GotoState101Action,
	{_State347, TuplePatternType}:                 _GotoState396Action,
	{_State348, DoToken}:                          _GotoState397Action,
	{_State350, IfToken}:                          _GotoState132Action,
	{_State350, LbraceToken}:                      _GotoState133Action,
	{_State350, IfExprType}:                       _GotoState399Action,
	{_State350, BlockBodyType}:                    _GotoState398Action,
	{_State351, LbracketToken}:                    _GotoState106Action,
	{_State351, DotToken}:                         _GotoState105Action,
	{_State351, DollarLbracketToken}:              _GotoState104Action,
	{_State351, OptionalGenericBindingType}:       _GotoState108Action,
	{_State367, IntegerLiteralToken}:              _GotoState25Action,
	{_State367, FloatLiteralToken}:                _GotoState22Action,
	{_State367, RuneLiteralToken}:                 _GotoState32Action,
	{_State367, StringLiteralToken}:               _GotoState33Action,
	{_State367, IdentifierToken}:                  _GotoState24Action,
	{_State367, TrueToken}:                        _GotoState36Action,
	{_State367, FalseToken}:                       _GotoState21Action,
	{_State367, StructToken}:                      _GotoState34Action,
	{_State367, FuncToken}:                        _GotoState23Action,
	{_State367, VarToken}:                         _GotoState37Action,
	{_State367, LabelDeclToken}:                   _GotoState26Action,
	{_State367, LparenToken}:                      _GotoState29Action,
	{_State367, LbracketToken}:                    _GotoState27Action,
	{_State367, NotToken}:                         _GotoState31Action,
	{_State367, SubToken}:                         _GotoState35Action,
	{_State367, MulToken}:                         _GotoState30Action,
	{_State367, BitNegToken}:                      _GotoState20Action,
	{_State367, BitAndToken}:                      _GotoState19Action,
	{_State367, LexErrorToken}:                    _GotoState28Action,
	{_State367, ExpressionType}:                   _GotoState400Action,
	{_State367, OptionalLabelDeclType}:            _GotoState51Action,
	{_State367, SequenceExprType}:                 _GotoState56Action,
	{_State367, BlockExprType}:                    _GotoState43Action,
	{_State367, CallExprType}:                     _GotoState44Action,
	{_State367, AtomExprType}:                     _GotoState42Action,
	{_State367, LiteralType}:                      _GotoState49Action,
	{_State367, ImplicitStructExprType}:           _GotoState47Action,
	{_State367, AccessExprType}:                   _GotoState38Action,
	{_State367, PostfixUnaryExprType}:             _GotoState53Action,
	{_State367, PrefixUnaryOpType}:                _GotoState55Action,
	{_State367, PrefixUnaryExprType}:              _GotoState54Action,
	{_State367, MulExprType}:                      _GotoState50Action,
	{_State367, AddExprType}:                      _GotoState39Action,
	{_State367, CmpExprType}:                      _GotoState45Action,
	{_State367, AndExprType}:                      _GotoState40Action,
	{_State367, OrExprType}:                       _GotoState52Action,
	{_State367, InitializableTypeType}:            _GotoState48Action,
	{_State367, ExplicitStructDefType}:            _GotoState46Action,
	{_State367, AnonymousFuncExprType}:            _GotoState41Action,
	{_State369, IntegerLiteralToken}:              _GotoState25Action,
	{_State369, FloatLiteralToken}:                _GotoState22Action,
	{_State369, RuneLiteralToken}:                 _GotoState32Action,
	{_State369, StringLiteralToken}:               _GotoState33Action,
	{_State369, IdentifierToken}:                  _GotoState24Action,
	{_State369, TrueToken}:                        _GotoState36Action,
	{_State369, FalseToken}:                       _GotoState21Action,
	{_State369, StructToken}:                      _GotoState34Action,
	{_State369, FuncToken}:                        _GotoState23Action,
	{_State369, VarToken}:                         _GotoState37Action,
	{_State369, LabelDeclToken}:                   _GotoState26Action,
	{_State369, LparenToken}:                      _GotoState29Action,
	{_State369, LbracketToken}:                    _GotoState27Action,
	{_State369, NotToken}:                         _GotoState31Action,
	{_State369, SubToken}:                         _GotoState35Action,
	{_State369, MulToken}:                         _GotoState30Action,
	{_State369, BitNegToken}:                      _GotoState20Action,
	{_State369, BitAndToken}:                      _GotoState19Action,
	{_State369, LexErrorToken}:                    _GotoState28Action,
	{_State369, ExpressionType}:                   _GotoState401Action,
	{_State369, OptionalLabelDeclType}:            _GotoState51Action,
	{_State369, SequenceExprType}:                 _GotoState56Action,
	{_State369, BlockExprType}:                    _GotoState43Action,
	{_State369, CallExprType}:                     _GotoState44Action,
	{_State369, AtomExprType}:                     _GotoState42Action,
	{_State369, LiteralType}:                      _GotoState49Action,
	{_State369, ImplicitStructExprType}:           _GotoState47Action,
	{_State369, AccessExprType}:                   _GotoState38Action,
	{_State369, PostfixUnaryExprType}:             _GotoState53Action,
	{_State369, PrefixUnaryOpType}:                _GotoState55Action,
	{_State369, PrefixUnaryExprType}:              _GotoState54Action,
	{_State369, MulExprType}:                      _GotoState50Action,
	{_State369, AddExprType}:                      _GotoState39Action,
	{_State369, CmpExprType}:                      _GotoState45Action,
	{_State369, AndExprType}:                      _GotoState40Action,
	{_State369, OrExprType}:                       _GotoState52Action,
	{_State369, InitializableTypeType}:            _GotoState48Action,
	{_State369, ExplicitStructDefType}:            _GotoState46Action,
	{_State369, AnonymousFuncExprType}:            _GotoState41Action,
	{_State371, IntegerLiteralToken}:              _GotoState25Action,
	{_State371, FloatLiteralToken}:                _GotoState22Action,
	{_State371, RuneLiteralToken}:                 _GotoState32Action,
	{_State371, StringLiteralToken}:               _GotoState33Action,
	{_State371, IdentifierToken}:                  _GotoState24Action,
	{_State371, TrueToken}:                        _GotoState36Action,
	{_State371, FalseToken}:                       _GotoState21Action,
	{_State371, StructToken}:                      _GotoState34Action,
	{_State371, FuncToken}:                        _GotoState23Action,
	{_State371, VarToken}:                         _GotoState37Action,
	{_State371, LabelDeclToken}:                   _GotoState26Action,
	{_State371, LparenToken}:                      _GotoState29Action,
	{_State371, LbracketToken}:                    _GotoState27Action,
	{_State371, NotToken}:                         _GotoState31Action,
	{_State371, SubToken}:                         _GotoState35Action,
	{_State371, MulToken}:                         _GotoState30Action,
	{_State371, BitNegToken}:                      _GotoState20Action,
	{_State371, BitAndToken}:                      _GotoState19Action,
	{_State371, LexErrorToken}:                    _GotoState28Action,
	{_State371, ExpressionType}:                   _GotoState402Action,
	{_State371, OptionalLabelDeclType}:            _GotoState51Action,
	{_State371, SequenceExprType}:                 _GotoState56Action,
	{_State371, BlockExprType}:                    _GotoState43Action,
	{_State371, ExpressionsType}:                  _GotoState403Action,
	{_State371, OptionalExpressionsType}:          _GotoState404Action,
	{_State371, CallExprType}:                     _GotoState44Action,
	{_State371, AtomExprType}:                     _GotoState42Action,
	{_State371, LiteralType}:                      _GotoState49Action,
	{_State371, ImplicitStructExprType}:           _GotoState47Action,
	{_State371, AccessExprType}:                   _GotoState38Action,
	{_State371, PostfixUnaryExprType}:             _GotoState53Action,
	{_State371, PrefixUnaryOpType}:                _GotoState55Action,
	{_State371, PrefixUnaryExprType}:              _GotoState54Action,
	{_State371, MulExprType}:                      _GotoState50Action,
	{_State371, AddExprType}:                      _GotoState39Action,
	{_State371, CmpExprType}:                      _GotoState45Action,
	{_State371, AndExprType}:                      _GotoState40Action,
	{_State371, OrExprType}:                       _GotoState52Action,
	{_State371, InitializableTypeType}:            _GotoState48Action,
	{_State371, ExplicitStructDefType}:            _GotoState46Action,
	{_State371, AnonymousFuncExprType}:            _GotoState41Action,
	{_State372, IntegerLiteralToken}:              _GotoState25Action,
	{_State372, FloatLiteralToken}:                _GotoState22Action,
	{_State372, RuneLiteralToken}:                 _GotoState32Action,
	{_State372, StringLiteralToken}:               _GotoState33Action,
	{_State372, IdentifierToken}:                  _GotoState24Action,
	{_State372, TrueToken}:                        _GotoState36Action,
	{_State372, FalseToken}:                       _GotoState21Action,
	{_State372, StructToken}:                      _GotoState34Action,
	{_State372, FuncToken}:                        _GotoState23Action,
	{_State372, VarToken}:                         _GotoState37Action,
	{_State372, LabelDeclToken}:                   _GotoState26Action,
	{_State372, LparenToken}:                      _GotoState29Action,
	{_State372, LbracketToken}:                    _GotoState27Action,
	{_State372, NotToken}:                         _GotoState31Action,
	{_State372, SubToken}:                         _GotoState35Action,
	{_State372, MulToken}:                         _GotoState30Action,
	{_State372, BitNegToken}:                      _GotoState20Action,
	{_State372, BitAndToken}:                      _GotoState19Action,
	{_State372, LexErrorToken}:                    _GotoState28Action,
	{_State372, ExpressionType}:                   _GotoState405Action,
	{_State372, OptionalLabelDeclType}:            _GotoState51Action,
	{_State372, SequenceExprType}:                 _GotoState56Action,
	{_State372, BlockExprType}:                    _GotoState43Action,
	{_State372, CallExprType}:                     _GotoState44Action,
	{_State372, AtomExprType}:                     _GotoState42Action,
	{_State372, LiteralType}:                      _GotoState49Action,
	{_State372, ImplicitStructExprType}:           _GotoState47Action,
	{_State372, AccessExprType}:                   _GotoState38Action,
	{_State372, PostfixUnaryExprType}:             _GotoState53Action,
	{_State372, PrefixUnaryOpType}:                _GotoState55Action,
	{_State372, PrefixUnaryExprType}:              _GotoState54Action,
	{_State372, MulExprType}:                      _GotoState50Action,
	{_State372, AddExprType}:                      _GotoState39Action,
	{_State372, CmpExprType}:                      _GotoState45Action,
	{_State372, AndExprType}:                      _GotoState40Action,
	{_State372, OrExprType}:                       _GotoState52Action,
	{_State372, InitializableTypeType}:            _GotoState48Action,
	{_State372, ExplicitStructDefType}:            _GotoState46Action,
	{_State372, AnonymousFuncExprType}:            _GotoState41Action,
	{_State375, IntegerLiteralToken}:              _GotoState25Action,
	{_State375, FloatLiteralToken}:                _GotoState22Action,
	{_State375, RuneLiteralToken}:                 _GotoState32Action,
	{_State375, StringLiteralToken}:               _GotoState33Action,
	{_State375, IdentifierToken}:                  _GotoState24Action,
	{_State375, TrueToken}:                        _GotoState36Action,
	{_State375, FalseToken}:                       _GotoState21Action,
	{_State375, StructToken}:                      _GotoState34Action,
	{_State375, FuncToken}:                        _GotoState23Action,
	{_State375, VarToken}:                         _GotoState209Action,
	{_State375, LabelDeclToken}:                   _GotoState26Action,
	{_State375, LparenToken}:                      _GotoState29Action,
	{_State375, LbracketToken}:                    _GotoState27Action,
	{_State375, DotToken}:                         _GotoState207Action,
	{_State375, NotToken}:                         _GotoState31Action,
	{_State375, SubToken}:                         _GotoState35Action,
	{_State375, MulToken}:                         _GotoState30Action,
	{_State375, BitNegToken}:                      _GotoState20Action,
	{_State375, BitAndToken}:                      _GotoState19Action,
	{_State375, LexErrorToken}:                    _GotoState28Action,
	{_State375, ExpressionType}:                   _GotoState210Action,
	{_State375, OptionalLabelDeclType}:            _GotoState51Action,
	{_State375, PatternsType}:                     _GotoState407Action,
	{_State375, SequenceExprType}:                 _GotoState56Action,
	{_State375, BlockExprType}:                    _GotoState43Action,
	{_State375, PatternType}:                      _GotoState406Action,
	{_State375, CallExprType}:                     _GotoState44Action,
	{_State375, AtomExprType}:                     _GotoState42Action,
	{_State375, LiteralType}:                      _GotoState49Action,
	{_State375, ImplicitStructExprType}:           _GotoState47Action,
	{_State375, AccessExprType}:                   _GotoState38Action,
	{_State375, PostfixUnaryExprType}:             _GotoState53Action,
	{_State375, PrefixUnaryOpType}:                _GotoState55Action,
	{_State375, PrefixUnaryExprType}:              _GotoState54Action,
	{_State375, MulExprType}:                      _GotoState50Action,
	{_State375, AddExprType}:                      _GotoState39Action,
	{_State375, CmpExprType}:                      _GotoState45Action,
	{_State375, AndExprType}:                      _GotoState40Action,
	{_State375, OrExprType}:                       _GotoState52Action,
	{_State375, InitializableTypeType}:            _GotoState48Action,
	{_State375, ExplicitStructDefType}:            _GotoState46Action,
	{_State375, AnonymousFuncExprType}:            _GotoState41Action,
	{_State377, CaseToken}:                        _GotoState375Action,
	{_State377, DefaultToken}:                     _GotoState408Action,
	{_State377, CaseBranchType}:                   _GotoState409Action,
	{_State377, OptionalDefaultBranchType}:        _GotoState411Action,
	{_State377, DefaultBranchType}:                _GotoState410Action,
	{_State382, AddToken}:                         _GotoState174Action,
	{_State382, SubToken}:                         _GotoState179Action,
	{_State382, MulToken}:                         _GotoState177Action,
	{_State383, RparenToken}:                      _GotoState412Action,
	{_State389, AddToken}:                         _GotoState174Action,
	{_State389, SubToken}:                         _GotoState179Action,
	{_State389, MulToken}:                         _GotoState177Action,
	{_State392, IdentifierToken}:                  _GotoState237Action,
	{_State392, StructToken}:                      _GotoState34Action,
	{_State392, EnumToken}:                        _GotoState74Action,
	{_State392, TraitToken}:                       _GotoState82Action,
	{_State392, FuncToken}:                        _GotoState76Action,
	{_State392, LparenToken}:                      _GotoState79Action,
	{_State392, LbracketToken}:                    _GotoState27Action,
	{_State392, DotToken}:                         _GotoState73Action,
	{_State392, QuestionToken}:                    _GotoState80Action,
	{_State392, ExclaimToken}:                     _GotoState75Action,
	{_State392, DotdotdotToken}:                   _GotoState236Action,
	{_State392, TildeTildeToken}:                  _GotoState81Action,
	{_State392, BitNegToken}:                      _GotoState72Action,
	{_State392, BitAndToken}:                      _GotoState71Action,
	{_State392, LexErrorToken}:                    _GotoState78Action,
	{_State392, InitializableTypeType}:            _GotoState88Action,
	{_State392, AtomTypeType}:                     _GotoState83Action,
	{_State392, ReturnableTypeType}:               _GotoState89Action,
	{_State392, ValueTypeType}:                    _GotoState241Action,
	{_State392, ImplicitStructDefType}:            _GotoState87Action,
	{_State392, ExplicitStructDefType}:            _GotoState46Action,
	{_State392, ImplicitEnumDefType}:              _GotoState86Action,
	{_State392, ExplicitEnumDefType}:              _GotoState84Action,
	{_State392, TraitDefType}:                     _GotoState90Action,
	{_State392, ParameterDeclType}:                _GotoState239Action,
	{_State392, ParameterDeclsType}:               _GotoState240Action,
	{_State392, OptionalParameterDeclsType}:       _GotoState413Action,
	{_State392, FuncTypeType}:                     _GotoState85Action,
	{_State395, DoToken}:                          _GotoState414Action,
	{_State397, LbraceToken}:                      _GotoState133Action,
	{_State397, BlockBodyType}:                    _GotoState415Action,
	{_State403, CommaToken}:                       _GotoState369Action,
	{_State407, CommaToken}:                       _GotoState417Action,
	{_State407, ColonToken}:                       _GotoState416Action,
	{_State408, ColonToken}:                       _GotoState418Action,
	{_State411, RbraceToken}:                      _GotoState419Action,
	{_State412, IdentifierToken}:                  _GotoState77Action,
	{_State412, StructToken}:                      _GotoState34Action,
	{_State412, EnumToken}:                        _GotoState74Action,
	{_State412, TraitToken}:                       _GotoState82Action,
	{_State412, FuncToken}:                        _GotoState76Action,
	{_State412, LparenToken}:                      _GotoState79Action,
	{_State412, LbracketToken}:                    _GotoState27Action,
	{_State412, DotToken}:                         _GotoState73Action,
	{_State412, QuestionToken}:                    _GotoState80Action,
	{_State412, ExclaimToken}:                     _GotoState75Action,
	{_State412, TildeTildeToken}:                  _GotoState81Action,
	{_State412, BitNegToken}:                      _GotoState72Action,
	{_State412, BitAndToken}:                      _GotoState71Action,
	{_State412, LexErrorToken}:                    _GotoState78Action,
	{_State412, InitializableTypeType}:            _GotoState88Action,
	{_State412, AtomTypeType}:                     _GotoState83Action,
	{_State412, ReturnableTypeType}:               _GotoState315Action,
	{_State412, ImplicitStructDefType}:            _GotoState87Action,
	{_State412, ExplicitStructDefType}:            _GotoState46Action,
	{_State412, ImplicitEnumDefType}:              _GotoState86Action,
	{_State412, ExplicitEnumDefType}:              _GotoState84Action,
	{_State412, TraitDefType}:                     _GotoState90Action,
	{_State412, ReturnTypeType}:                   _GotoState420Action,
	{_State412, FuncTypeType}:                     _GotoState85Action,
	{_State413, RparenToken}:                      _GotoState421Action,
	{_State414, LbraceToken}:                      _GotoState133Action,
	{_State414, BlockBodyType}:                    _GotoState422Action,
	{_State416, IntegerLiteralToken}:              _GotoState25Action,
	{_State416, FloatLiteralToken}:                _GotoState22Action,
	{_State416, RuneLiteralToken}:                 _GotoState32Action,
	{_State416, StringLiteralToken}:               _GotoState33Action,
	{_State416, IdentifierToken}:                  _GotoState24Action,
	{_State416, TrueToken}:                        _GotoState36Action,
	{_State416, FalseToken}:                       _GotoState21Action,
	{_State416, StructToken}:                      _GotoState34Action,
	{_State416, FuncToken}:                        _GotoState23Action,
	{_State416, LabelDeclToken}:                   _GotoState26Action,
	{_State416, LparenToken}:                      _GotoState29Action,
	{_State416, LbracketToken}:                    _GotoState27Action,
	{_State416, NotToken}:                         _GotoState31Action,
	{_State416, SubToken}:                         _GotoState35Action,
	{_State416, MulToken}:                         _GotoState30Action,
	{_State416, BitNegToken}:                      _GotoState20Action,
	{_State416, BitAndToken}:                      _GotoState19Action,
	{_State416, LexErrorToken}:                    _GotoState28Action,
	{_State416, OptionalLabelDeclType}:            _GotoState140Action,
	{_State416, SequenceExprType}:                 _GotoState423Action,
	{_State416, BlockExprType}:                    _GotoState43Action,
	{_State416, CallExprType}:                     _GotoState44Action,
	{_State416, AtomExprType}:                     _GotoState42Action,
	{_State416, LiteralType}:                      _GotoState49Action,
	{_State416, ImplicitStructExprType}:           _GotoState47Action,
	{_State416, AccessExprType}:                   _GotoState38Action,
	{_State416, PostfixUnaryExprType}:             _GotoState53Action,
	{_State416, PrefixUnaryOpType}:                _GotoState55Action,
	{_State416, PrefixUnaryExprType}:              _GotoState54Action,
	{_State416, MulExprType}:                      _GotoState50Action,
	{_State416, AddExprType}:                      _GotoState39Action,
	{_State416, CmpExprType}:                      _GotoState45Action,
	{_State416, AndExprType}:                      _GotoState40Action,
	{_State416, OrExprType}:                       _GotoState52Action,
	{_State416, InitializableTypeType}:            _GotoState48Action,
	{_State416, ExplicitStructDefType}:            _GotoState46Action,
	{_State416, AnonymousFuncExprType}:            _GotoState41Action,
	{_State417, IntegerLiteralToken}:              _GotoState25Action,
	{_State417, FloatLiteralToken}:                _GotoState22Action,
	{_State417, RuneLiteralToken}:                 _GotoState32Action,
	{_State417, StringLiteralToken}:               _GotoState33Action,
	{_State417, IdentifierToken}:                  _GotoState24Action,
	{_State417, TrueToken}:                        _GotoState36Action,
	{_State417, FalseToken}:                       _GotoState21Action,
	{_State417, StructToken}:                      _GotoState34Action,
	{_State417, FuncToken}:                        _GotoState23Action,
	{_State417, VarToken}:                         _GotoState209Action,
	{_State417, LabelDeclToken}:                   _GotoState26Action,
	{_State417, LparenToken}:                      _GotoState29Action,
	{_State417, LbracketToken}:                    _GotoState27Action,
	{_State417, DotToken}:                         _GotoState207Action,
	{_State417, NotToken}:                         _GotoState31Action,
	{_State417, SubToken}:                         _GotoState35Action,
	{_State417, MulToken}:                         _GotoState30Action,
	{_State417, BitNegToken}:                      _GotoState20Action,
	{_State417, BitAndToken}:                      _GotoState19Action,
	{_State417, LexErrorToken}:                    _GotoState28Action,
	{_State417, ExpressionType}:                   _GotoState210Action,
	{_State417, OptionalLabelDeclType}:            _GotoState51Action,
	{_State417, SequenceExprType}:                 _GotoState56Action,
	{_State417, BlockExprType}:                    _GotoState43Action,
	{_State417, PatternType}:                      _GotoState424Action,
	{_State417, CallExprType}:                     _GotoState44Action,
	{_State417, AtomExprType}:                     _GotoState42Action,
	{_State417, LiteralType}:                      _GotoState49Action,
	{_State417, ImplicitStructExprType}:           _GotoState47Action,
	{_State417, AccessExprType}:                   _GotoState38Action,
	{_State417, PostfixUnaryExprType}:             _GotoState53Action,
	{_State417, PrefixUnaryOpType}:                _GotoState55Action,
	{_State417, PrefixUnaryExprType}:              _GotoState54Action,
	{_State417, MulExprType}:                      _GotoState50Action,
	{_State417, AddExprType}:                      _GotoState39Action,
	{_State417, CmpExprType}:                      _GotoState45Action,
	{_State417, AndExprType}:                      _GotoState40Action,
	{_State417, OrExprType}:                       _GotoState52Action,
	{_State417, InitializableTypeType}:            _GotoState48Action,
	{_State417, ExplicitStructDefType}:            _GotoState46Action,
	{_State417, AnonymousFuncExprType}:            _GotoState41Action,
	{_State418, IntegerLiteralToken}:              _GotoState25Action,
	{_State418, FloatLiteralToken}:                _GotoState22Action,
	{_State418, RuneLiteralToken}:                 _GotoState32Action,
	{_State418, StringLiteralToken}:               _GotoState33Action,
	{_State418, IdentifierToken}:                  _GotoState24Action,
	{_State418, TrueToken}:                        _GotoState36Action,
	{_State418, FalseToken}:                       _GotoState21Action,
	{_State418, StructToken}:                      _GotoState34Action,
	{_State418, FuncToken}:                        _GotoState23Action,
	{_State418, LabelDeclToken}:                   _GotoState26Action,
	{_State418, LparenToken}:                      _GotoState29Action,
	{_State418, LbracketToken}:                    _GotoState27Action,
	{_State418, NotToken}:                         _GotoState31Action,
	{_State418, SubToken}:                         _GotoState35Action,
	{_State418, MulToken}:                         _GotoState30Action,
	{_State418, BitNegToken}:                      _GotoState20Action,
	{_State418, BitAndToken}:                      _GotoState19Action,
	{_State418, LexErrorToken}:                    _GotoState28Action,
	{_State418, OptionalLabelDeclType}:            _GotoState140Action,
	{_State418, SequenceExprType}:                 _GotoState425Action,
	{_State418, BlockExprType}:                    _GotoState43Action,
	{_State418, CallExprType}:                     _GotoState44Action,
	{_State418, AtomExprType}:                     _GotoState42Action,
	{_State418, LiteralType}:                      _GotoState49Action,
	{_State418, ImplicitStructExprType}:           _GotoState47Action,
	{_State418, AccessExprType}:                   _GotoState38Action,
	{_State418, PostfixUnaryExprType}:             _GotoState53Action,
	{_State418, PrefixUnaryOpType}:                _GotoState55Action,
	{_State418, PrefixUnaryExprType}:              _GotoState54Action,
	{_State418, MulExprType}:                      _GotoState50Action,
	{_State418, AddExprType}:                      _GotoState39Action,
	{_State418, CmpExprType}:                      _GotoState45Action,
	{_State418, AndExprType}:                      _GotoState40Action,
	{_State418, OrExprType}:                       _GotoState52Action,
	{_State418, InitializableTypeType}:            _GotoState48Action,
	{_State418, ExplicitStructDefType}:            _GotoState46Action,
	{_State418, AnonymousFuncExprType}:            _GotoState41Action,
	{_State420, LbraceToken}:                      _GotoState133Action,
	{_State420, BlockBodyType}:                    _GotoState426Action,
	{_State421, IdentifierToken}:                  _GotoState77Action,
	{_State421, StructToken}:                      _GotoState34Action,
	{_State421, EnumToken}:                        _GotoState74Action,
	{_State421, TraitToken}:                       _GotoState82Action,
	{_State421, FuncToken}:                        _GotoState76Action,
	{_State421, LparenToken}:                      _GotoState79Action,
	{_State421, LbracketToken}:                    _GotoState27Action,
	{_State421, DotToken}:                         _GotoState73Action,
	{_State421, QuestionToken}:                    _GotoState80Action,
	{_State421, ExclaimToken}:                     _GotoState75Action,
	{_State421, TildeTildeToken}:                  _GotoState81Action,
	{_State421, BitNegToken}:                      _GotoState72Action,
	{_State421, BitAndToken}:                      _GotoState71Action,
	{_State421, LexErrorToken}:                    _GotoState78Action,
	{_State421, InitializableTypeType}:            _GotoState88Action,
	{_State421, AtomTypeType}:                     _GotoState83Action,
	{_State421, ReturnableTypeType}:               _GotoState315Action,
	{_State421, ImplicitStructDefType}:            _GotoState87Action,
	{_State421, ExplicitStructDefType}:            _GotoState46Action,
	{_State421, ImplicitEnumDefType}:              _GotoState86Action,
	{_State421, ExplicitEnumDefType}:              _GotoState84Action,
	{_State421, TraitDefType}:                     _GotoState90Action,
	{_State421, ReturnTypeType}:                   _GotoState427Action,
	{_State421, FuncTypeType}:                     _GotoState85Action,
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
	{_State38, _WildcardMarker}:                   _ReduceAccessExprToPostfixUnaryExprAction,
	{_State38, LparenToken}:                       _ReduceNilToOptionalGenericBindingAction,
	{_State39, _WildcardMarker}:                   _ReduceAddExprToCmpExprAction,
	{_State40, _WildcardMarker}:                   _ReduceAndExprToOrExprAction,
	{_State41, _WildcardMarker}:                   _ReduceAnonymousFuncExprToAtomExprAction,
	{_State42, _WildcardMarker}:                   _ReduceAtomExprToAccessExprAction,
	{_State43, _WildcardMarker}:                   _ReduceBlockExprToAtomExprAction,
	{_State44, _WildcardMarker}:                   _ReduceCallExprToAccessExprAction,
	{_State45, _WildcardMarker}:                   _ReduceCmpExprToAndExprAction,
	{_State46, _WildcardMarker}:                   _ReduceExplicitStructDefToInitializableTypeAction,
	{_State47, _WildcardMarker}:                   _ReduceImplicitStructExprToAtomExprAction,
	{_State49, _WildcardMarker}:                   _ReduceLiteralToAtomExprAction,
	{_State50, _WildcardMarker}:                   _ReduceMulExprToAddExprAction,
	{_State52, _WildcardMarker}:                   _ReduceToSequenceExprAction,
	{_State53, _WildcardMarker}:                   _ReducePostfixUnaryExprToPrefixUnaryExprAction,
	{_State54, _WildcardMarker}:                   _ReducePrefixUnaryExprToMulExprAction,
	{_State55, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State56, _EndMarker}:                        _ReduceSequenceExprToExpressionAction,
	{_State57, _EndMarker}:                        _ReduceCommentToLexInternalTokensAction,
	{_State58, _EndMarker}:                        _ReduceSpacesToLexInternalTokensAction,
	{_State60, _WildcardMarker}:                   _ReduceNilToDefinitionsAction,
	{_State61, _EndMarker}:                        _ReduceNilToOptionalNewlinesAction,
	{_State62, _WildcardMarker}:                   _ReduceNamedFuncDefToDefinitionAction,
	{_State63, _WildcardMarker}:                   _ReducePackageDefToDefinitionAction,
	{_State64, _WildcardMarker}:                   _ReduceTypeDefToDefinitionAction,
	{_State65, _WildcardMarker}:                   _ReduceUnsafeStatementToDefinitionAction,
	{_State66, _WildcardMarker}:                   _ReduceNoSpecToPackageDefAction,
	{_State67, _WildcardMarker}:                   _ReduceNilToOptionalGenericParametersAction,
	{_State70, RparenToken}:                       _ReduceNilToOptionalParameterDefsAction,
	{_State73, _WildcardMarker}:                   _ReduceNilToOptionalGenericBindingAction,
	{_State77, _WildcardMarker}:                   _ReduceNilToOptionalGenericBindingAction,
	{_State78, _WildcardMarker}:                   _ReduceLexErrorToAtomTypeAction,
	{_State79, RparenToken}:                       _ReduceNilToOptionalImplicitFieldDefsAction,
	{_State83, _WildcardMarker}:                   _ReduceAtomTypeToReturnableTypeAction,
	{_State84, _WildcardMarker}:                   _ReduceExplicitEnumDefToAtomTypeAction,
	{_State85, _WildcardMarker}:                   _ReduceFuncTypeToAtomTypeAction,
	{_State86, _WildcardMarker}:                   _ReduceImplicitEnumDefToAtomTypeAction,
	{_State87, _WildcardMarker}:                   _ReduceImplicitStructDefToAtomTypeAction,
	{_State88, _WildcardMarker}:                   _ReduceInitializableTypeToAtomTypeAction,
	{_State89, _WildcardMarker}:                   _ReduceReturnableTypeToValueTypeAction,
	{_State90, _WildcardMarker}:                   _ReduceTraitDefToAtomTypeAction,
	{_State92, _WildcardMarker}:                   _ReduceDotdotdotToArgumentAction,
	{_State93, _WildcardMarker}:                   _ReduceIdentifierToAtomExprAction,
	{_State94, _WildcardMarker}:                   _ReduceArgumentToArgumentsAction,
	{_State96, _WildcardMarker}:                   _ReduceColonExpressionsToArgumentAction,
	{_State97, _WildcardMarker}:                   _ReducePositionalToArgumentAction,
	{_State97, ColonToken}:                        _ReduceExpressionToOptionalExpressionAction,
	{_State99, RparenToken}:                       _ReduceNilToOptionalExplicitFieldDefsAction,
	{_State100, _WildcardMarker}:                  _ReduceIdentifierToVarPatternAction,
	{_State102, _WildcardMarker}:                  _ReduceTuplePatternToVarPatternAction,
	{_State103, _WildcardMarker}:                  _ReduceNilToOptionalValueTypeAction,
	{_State104, RbracketToken}:                    _ReduceNilToOptionalGenericArgumentsAction,
	{_State106, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State106, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State107, _WildcardMarker}:                  _ReduceQuestionToPostfixUnaryExprAction,
	{_State109, _WildcardMarker}:                  _ReduceAddToAddOpAction,
	{_State110, _WildcardMarker}:                  _ReduceBitOrToAddOpAction,
	{_State111, _WildcardMarker}:                  _ReduceBitXorToAddOpAction,
	{_State112, _WildcardMarker}:                  _ReduceSubToAddOpAction,
	{_State113, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State114, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State115, _WildcardMarker}:                  _ReduceEqualToCmpOpAction,
	{_State116, _WildcardMarker}:                  _ReduceGreaterToCmpOpAction,
	{_State117, _WildcardMarker}:                  _ReduceGreaterOrEqualToCmpOpAction,
	{_State118, _WildcardMarker}:                  _ReduceLessToCmpOpAction,
	{_State119, _WildcardMarker}:                  _ReduceLessOrEqualToCmpOpAction,
	{_State120, _WildcardMarker}:                  _ReduceNotEqualToCmpOpAction,
	{_State121, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State122, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State122, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State123, _WildcardMarker}:                  _ReduceBitAndToMulOpAction,
	{_State124, _WildcardMarker}:                  _ReduceBitLshiftToMulOpAction,
	{_State125, _WildcardMarker}:                  _ReduceBitRshiftToMulOpAction,
	{_State126, _WildcardMarker}:                  _ReduceDivToMulOpAction,
	{_State127, _WildcardMarker}:                  _ReduceModToMulOpAction,
	{_State128, _WildcardMarker}:                  _ReduceMulToMulOpAction,
	{_State129, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State131, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State132, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State133, _WildcardMarker}:                  _ReduceEmptyListToStatementsAction,
	{_State134, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State135, _WildcardMarker}:                  _ReduceToBlockExprAction,
	{_State136, _EndMarker}:                       _ReduceIfExprToExpressionAction,
	{_State137, _EndMarker}:                       _ReduceLoopExprToExpressionAction,
	{_State138, _EndMarker}:                       _ReduceSwitchExprToExpressionAction,
	{_State139, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State141, _WildcardMarker}:                  _ReducePrefixOpToPrefixUnaryExprAction,
	{_State143, _EndMarker}:                       _ReduceNewlinesToOptionalNewlinesAction,
	{_State144, _EndMarker}:                       _ReduceDefinitionsToOptionalDefinitionsAction,
	{_State145, _WildcardMarker}:                  _ReduceEmptyListToPackageStatementsAction,
	{_State146, RbracketToken}:                    _ReduceNilToOptionalGenericParameterDefsAction,
	{_State151, LparenToken}:                      _ReduceNilToOptionalGenericParametersAction,
	{_State153, _WildcardMarker}:                  _ReduceParameterDefToParameterDefsAction,
	{_State154, RparenToken}:                      _ReduceParameterDefsToOptionalParameterDefsAction,
	{_State155, _WildcardMarker}:                  _ReduceReferenceToReturnableTypeAction,
	{_State156, _WildcardMarker}:                  _ReducePublicMethodsTraitToReturnableTypeAction,
	{_State157, _WildcardMarker}:                  _ReduceInferredToAtomTypeAction,
	{_State159, _WildcardMarker}:                  _ReduceResultToReturnableTypeAction,
	{_State160, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State162, _WildcardMarker}:                  _ReduceNamedToAtomTypeAction,
	{_State163, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State165, _WildcardMarker}:                  _ReduceFieldDefToImplicitFieldDefsAction,
	{_State165, OrToken}:                          _ReduceFieldDefToEnumValueDefAction,
	{_State167, RparenToken}:                      _ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction,
	{_State169, _WildcardMarker}:                  _ReduceUnsafeStatementToFieldDefAction,
	{_State170, _WildcardMarker}:                  _ReduceImplicitToFieldDefAction,
	{_State171, _WildcardMarker}:                  _ReduceOptionalToReturnableTypeAction,
	{_State172, _WildcardMarker}:                  _ReducePublicTraitToReturnableTypeAction,
	{_State173, RparenToken}:                      _ReduceNilToOptionalTraitPropertiesAction,
	{_State178, _WildcardMarker}:                  _ReduceSliceToInitializableTypeAction,
	{_State180, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State181, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State181, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State182, _WildcardMarker}:                  _ReduceToImplicitStructExprAction,
	{_State183, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State183, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State183, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State183, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State183, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State184, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State184, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State184, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State184, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State184, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State185, RparenToken}:                      _ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction,
	{_State186, _WildcardMarker}:                  _ReduceFieldDefToExplicitFieldDefsAction,
	{_State188, _WildcardMarker}:                  _ReduceDotdotdotToFieldVarPatternAction,
	{_State189, _WildcardMarker}:                  _ReduceIdentifierToVarPatternAction,
	{_State190, _WildcardMarker}:                  _ReduceFieldVarPatternToFieldVarPatternsAction,
	{_State192, _WildcardMarker}:                  _ReducePositionalToFieldVarPatternAction,
	{_State193, _EndMarker}:                       _ReduceVarPatternToExpressionAction,
	{_State194, _WildcardMarker}:                  _ReduceValueTypeToOptionalValueTypeAction,
	{_State195, RbracketToken}:                    _ReduceGenericArgumentsToOptionalGenericArgumentsAction,
	{_State197, _WildcardMarker}:                  _ReduceValueTypeToGenericArgumentsAction,
	{_State198, _WildcardMarker}:                  _ReduceAccessToAccessExprAction,
	{_State200, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State200, RparenToken}:                      _ReduceNilToOptionalArgumentsAction,
	{_State200, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State201, _WildcardMarker}:                  _ReduceOpToAddExprAction,
	{_State202, _WildcardMarker}:                  _ReduceOpToAndExprAction,
	{_State203, _WildcardMarker}:                  _ReduceOpToCmpExprAction,
	{_State205, _WildcardMarker}:                  _ReduceOpToMulExprAction,
	{_State206, _WildcardMarker}:                  _ReduceInfiniteToLoopExprAction,
	{_State208, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State208, SemicolonToken}:                   _ReduceNilToOptionalSequenceExprAction,
	{_State210, _WildcardMarker}:                  _ReduceExpressionToPatternAction,
	{_State212, InToken}:                          _ReduceSequenceExprToExpressionAction,
	{_State214, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State216, _WildcardMarker}:                  _ReduceOpToOrExprAction,
	{_State218, _WildcardMarker}:                  _ReduceAddToDefinitionsAction,
	{_State220, _WildcardMarker}:                  _ReduceUnconstrainedToGenericParameterDefAction,
	{_State221, _WildcardMarker}:                  _ReduceGenericParameterDefToGenericParameterDefsAction,
	{_State222, RbracketToken}:                    _ReduceGenericParameterDefsToOptionalGenericParameterDefsAction,
	{_State224, _EndMarker}:                       _ReduceAliasToTypeDefAction,
	{_State225, _WildcardMarker}:                  _ReduceDefinitionToTypeDefAction,
	{_State227, _WildcardMarker}:                  _ReduceArgToParameterDefAction,
	{_State228, IdentifierToken}:                  _ReduceReceiverToOptionalReceiverAction,
	{_State230, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State234, _WildcardMarker}:                  _ReduceFieldDefToEnumValueDefAction,
	{_State237, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State239, _WildcardMarker}:                  _ReduceParameterDeclToParameterDeclsAction,
	{_State240, RparenToken}:                      _ReduceParameterDeclsToOptionalParameterDeclsAction,
	{_State241, _WildcardMarker}:                  _ReduceUnamedToParameterDeclAction,
	{_State242, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State243, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State244, _WildcardMarker}:                  _ReduceExplicitToFieldDefAction,
	{_State248, _WildcardMarker}:                  _ReduceToImplicitEnumDefAction,
	{_State250, _WildcardMarker}:                  _ReduceToImplicitStructDefAction,
	{_State252, _WildcardMarker}:                  _ReduceFieldDefToTraitPropertyAction,
	{_State253, _WildcardMarker}:                  _ReduceMethodSignatureToTraitPropertyAction,
	{_State255, RparenToken}:                      _ReduceTraitPropertiesToOptionalTraitPropertiesAction,
	{_State256, _WildcardMarker}:                  _ReduceTraitPropertyToTraitPropertiesAction,
	{_State257, _WildcardMarker}:                  _ReduceTraitUnionToValueTypeAction,
	{_State260, _WildcardMarker}:                  _ReduceTraitIntersectToValueTypeAction,
	{_State261, _WildcardMarker}:                  _ReduceTraitDifferenceToValueTypeAction,
	{_State262, _WildcardMarker}:                  _ReduceNamedToArgumentAction,
	{_State263, _WildcardMarker}:                  _ReduceAddToArgumentsAction,
	{_State264, _WildcardMarker}:                  _ReduceExpressionToOptionalExpressionAction,
	{_State265, _WildcardMarker}:                  _ReduceAddToColonExpressionsAction,
	{_State266, _WildcardMarker}:                  _ReducePairToColonExpressionsAction,
	{_State269, _WildcardMarker}:                  _ReduceToExplicitStructDefAction,
	{_State272, _WildcardMarker}:                  _ReduceToTuplePatternAction,
	{_State274, _WildcardMarker}:                  _ReduceBindingToOptionalGenericBindingAction,
	{_State275, _WildcardMarker}:                  _ReduceIndexToAccessExprAction,
	{_State276, RparenToken}:                      _ReduceArgumentsToOptionalArgumentsAction,
	{_State278, _WildcardMarker}:                  _ReduceInitializeExprToAtomExprAction,
	{_State279, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State282, DoToken}:                          _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State284, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State286, _WildcardMarker}:                  _ReduceNoElseToIfExprAction,
	{_State287, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State288, _WildcardMarker}:                  _ReduceBreakToJumpTypeAction,
	{_State289, _WildcardMarker}:                  _ReduceContinueToJumpTypeAction,
	{_State290, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State291, _EndMarker}:                       _ReduceToBlockBodyAction,
	{_State292, _WildcardMarker}:                  _ReduceReturnToJumpTypeAction,
	{_State293, _WildcardMarker}:                  _ReduceAccessExprToPostfixUnaryExprAction,
	{_State293, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State294, AssignToken}:                      _ReduceExpressionToPatternAction,
	{_State294, _WildcardMarker}:                  _ReduceExpressionToExpressionsAction,
	{_State295, _WildcardMarker}:                  _ReduceExpressionOrImplicitStructToStatementBodyAction,
	{_State296, _WildcardMarker}:                  _ReduceJumpStatementToStatementBodyAction,
	{_State297, _WildcardMarker}:                  _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State299, _WildcardMarker}:                  _ReduceAddToStatementsAction,
	{_State301, _WildcardMarker}:                  _ReduceUnsafeStatementToStatementBodyAction,
	{_State304, _EndMarker}:                       _ReduceWithSpecToPackageDefAction,
	{_State305, _WildcardMarker}:                  _ReduceAddToPackageStatementsAction,
	{_State307, _WildcardMarker}:                  _ReduceToPackageStatementBodyAction,
	{_State308, _WildcardMarker}:                  _ReduceConstrainedToGenericParameterDefAction,
	{_State310, _WildcardMarker}:                  _ReduceGenericToOptionalGenericParametersAction,
	{_State312, _WildcardMarker}:                  _ReduceVarargToParameterDefAction,
	{_State313, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State315, _WildcardMarker}:                  _ReduceReturnableTypeToReturnTypeAction,
	{_State316, _WildcardMarker}:                  _ReduceAddToParameterDefsAction,
	{_State319, _WildcardMarker}:                  _ReduceToExplicitEnumDefAction,
	{_State322, _WildcardMarker}:                  _ReduceUnnamedVarargToParameterDeclAction,
	{_State324, _WildcardMarker}:                  _ReduceArgToParameterDeclAction,
	{_State325, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State327, _WildcardMarker}:                  _ReduceExternNamedToAtomTypeAction,
	{_State328, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State329, _WildcardMarker}:                  _ReduceDefaultToEnumValueDefAction,
	{_State330, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State331, _WildcardMarker}:                  _ReduceAddToImplicitFieldDefsAction,
	{_State333, _WildcardMarker}:                  _ReduceToTraitDefAction,
	{_State336, _WildcardMarker}:                  _ReduceMapToInitializableTypeAction,
	{_State337, _WildcardMarker}:                  _ReduceArrayToInitializableTypeAction,
	{_State338, _WildcardMarker}:                  _ReduceExplicitToExplicitFieldDefsAction,
	{_State339, _WildcardMarker}:                  _ReduceImplicitToExplicitFieldDefsAction,
	{_State340, _WildcardMarker}:                  _ReduceNamedToFieldVarPatternAction,
	{_State341, _WildcardMarker}:                  _ReduceAddToFieldVarPatternsAction,
	{_State342, _WildcardMarker}:                  _ReduceAddToGenericArgumentsAction,
	{_State343, _WildcardMarker}:                  _ReduceToCallExprAction,
	{_State344, _EndMarker}:                       _ReduceDoWhileToLoopExprAction,
	{_State345, AssignToken}:                      _ReduceEnumMatchPatternToPatternAction,
	{_State346, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State346, DoToken}:                          _ReduceNilToOptionalSequenceExprAction,
	{_State349, _EndMarker}:                       _ReduceWhileToLoopExprAction,
	{_State351, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State352, NewlinesToken}:                    _ReduceAsyncToStatementBodyAction,
	{_State352, SemicolonToken}:                   _ReduceAsyncToStatementBodyAction,
	{_State352, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State353, NewlinesToken}:                    _ReduceDeferToStatementBodyAction,
	{_State353, SemicolonToken}:                   _ReduceDeferToStatementBodyAction,
	{_State353, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State354, _WildcardMarker}:                  _ReduceAddAssignToBinaryOpAssignAction,
	{_State355, _WildcardMarker}:                  _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State356, _WildcardMarker}:                  _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State357, _WildcardMarker}:                  _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State358, _WildcardMarker}:                  _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State359, _WildcardMarker}:                  _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State360, _WildcardMarker}:                  _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State361, _WildcardMarker}:                  _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State362, _WildcardMarker}:                  _ReduceDivAssignToBinaryOpAssignAction,
	{_State363, _WildcardMarker}:                  _ReduceModAssignToBinaryOpAssignAction,
	{_State364, _WildcardMarker}:                  _ReduceMulAssignToBinaryOpAssignAction,
	{_State365, _WildcardMarker}:                  _ReduceSubAssignToBinaryOpAssignAction,
	{_State366, _WildcardMarker}:                  _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State367, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State368, _WildcardMarker}:                  _ReduceUnaryOpAssignStatementToStatementBodyAction,
	{_State369, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State370, _WildcardMarker}:                  _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State371, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State371, NewlinesToken}:                    _ReduceNilToOptionalExpressionsAction,
	{_State371, SemicolonToken}:                   _ReduceNilToOptionalExpressionsAction,
	{_State372, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State373, _WildcardMarker}:                  _ReduceImplicitToStatementAction,
	{_State374, _WildcardMarker}:                  _ReduceExplicitToStatementAction,
	{_State375, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State376, _WildcardMarker}:                  _ReduceCaseBranchToCaseBranchesAction,
	{_State377, RbraceToken}:                      _ReduceNilToOptionalDefaultBranchAction,
	{_State378, _WildcardMarker}:                  _ReduceToUnsafeStatementAction,
	{_State379, _WildcardMarker}:                  _ReduceImplicitToPackageStatementAction,
	{_State380, _WildcardMarker}:                  _ReduceExplicitToPackageStatementAction,
	{_State381, _WildcardMarker}:                  _ReduceAddToGenericParameterDefsAction,
	{_State382, _EndMarker}:                       _ReduceConstrainedDefToTypeDefAction,
	{_State384, _WildcardMarker}:                  _ReduceToAnonymousFuncExprAction,
	{_State385, RparenToken}:                      _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State386, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State386, RparenToken}:                      _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State387, RparenToken}:                      _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State388, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State388, RparenToken}:                      _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State389, _WildcardMarker}:                  _ReduceVarargToParameterDeclAction,
	{_State390, _WildcardMarker}:                  _ReduceToFuncTypeAction,
	{_State391, _WildcardMarker}:                  _ReduceAddToParameterDeclsAction,
	{_State392, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State393, _WildcardMarker}:                  _ReduceExplicitToTraitPropertiesAction,
	{_State394, _WildcardMarker}:                  _ReduceImplicitToTraitPropertiesAction,
	{_State396, AssignToken}:                      _ReduceEnumVarPatternToPatternAction,
	{_State398, _EndMarker}:                       _ReduceIfElseToIfExprAction,
	{_State399, _EndMarker}:                       _ReduceMultiIfElseToIfExprAction,
	{_State400, _WildcardMarker}:                  _ReduceBinaryOpAssignStatementToStatementBodyAction,
	{_State401, _WildcardMarker}:                  _ReduceAddToExpressionsAction,
	{_State402, _WildcardMarker}:                  _ReduceExpressionToExpressionsAction,
	{_State403, _WildcardMarker}:                  _ReduceExpressionsToOptionalExpressionsAction,
	{_State404, _WildcardMarker}:                  _ReduceToJumpStatementAction,
	{_State405, _WildcardMarker}:                  _ReduceAssignStatementToStatementBodyAction,
	{_State406, _WildcardMarker}:                  _ReducePatternToPatternsAction,
	{_State409, _WildcardMarker}:                  _ReduceAddToCaseBranchesAction,
	{_State410, RbraceToken}:                      _ReduceDefaultBranchToOptionalDefaultBranchAction,
	{_State412, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State415, _EndMarker}:                       _ReduceIteratorToLoopExprAction,
	{_State416, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State417, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State418, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State419, _EndMarker}:                       _ReduceToSwitchExprAction,
	{_State421, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State422, _EndMarker}:                       _ReduceForToLoopExprAction,
	{_State423, _WildcardMarker}:                  _ReduceToCaseBranchAction,
	{_State424, _WildcardMarker}:                  _ReduceMultipleToPatternsAction,
	{_State425, RbraceToken}:                      _ReduceToDefaultBranchAction,
	{_State426, _EndMarker}:                       _ReduceToNamedFuncDefAction,
	{_State427, _WildcardMarker}:                  _ReduceToMethodSignatureAction,
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
      VAR -> State 37
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
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 6:
    Kernel Items:
      #accept: ^.lex_internal_tokens
    Reduce:
      (nil)
    Goto:
      SPACES -> State 58
      COMMENT -> State 57
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
      UNSAFE -> State 59
      TYPE -> State 17
      FUNC -> State 18
      definitions -> State 61
      definition -> State 60
      unsafe_statement -> State 65
      type_def -> State 64
      named_func_def -> State 62
      package_def -> State 63

  State 16:
    Kernel Items:
      package_def: PACKAGE.IDENTIFIER
      package_def: PACKAGE.IDENTIFIER LPAREN package_statements RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 66

  State 17:
    Kernel Items:
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE.IDENTIFIER EQUAL value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 67

  State 18:
    Kernel Items:
      named_func_def: FUNC.optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      LPAREN -> State 68
      optional_receiver -> State 69

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
      LPAREN -> State 70

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
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 91
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 28:
    Kernel Items:
      atom_expr: LEX_ERROR., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 29:
    Kernel Items:
      implicit_struct_expr: LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 93
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      VAR -> State 37
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      DOTDOTDOT -> State 92
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 97
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      arguments -> State 95
      argument -> State 94
      colon_expressions -> State 96
      optional_expression -> State 98
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

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
      LPAREN -> State 99

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
      expression: VAR.var_pattern optional_value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 100
      LPAREN -> State 101
      var_pattern -> State 103
      tuple_pattern -> State 102

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
      LBRACKET -> State 106
      DOT -> State 105
      QUESTION -> State 107
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 108

  State 39:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 109
      SUB -> State 112
      BIT_XOR -> State 111
      BIT_OR -> State 110
      add_op -> State 113

  State 40:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 114

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
      EQUAL -> State 115
      NOT_EQUAL -> State 120
      LESS -> State 118
      LESS_OR_EQUAL -> State 119
      GREATER -> State 116
      GREATER_OR_EQUAL -> State 117
      cmp_op -> State 121

  State 46:
    Kernel Items:
      initializable_type: explicit_struct_def., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 47:
    Kernel Items:
      atom_expr: implicit_struct_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      atom_expr: initializable_type.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 122

  State 49:
    Kernel Items:
      atom_expr: literal., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 50:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 128
      DIV -> State 126
      MOD -> State 127
      BIT_AND -> State 123
      BIT_LSHIFT -> State 124
      BIT_RSHIFT -> State 125
      mul_op -> State 129

  State 51:
    Kernel Items:
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      IF -> State 132
      SWITCH -> State 134
      FOR -> State 131
      DO -> State 130
      LBRACE -> State 133
      if_expr -> State 136
      switch_expr -> State 138
      loop_expr -> State 137
      block_body -> State 135

  State 52:
    Kernel Items:
      sequence_expr: or_expr., *
      or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    Goto:
      OR -> State 139

  State 53:
    Kernel Items:
      prefix_unary_expr: postfix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 54:
    Kernel Items:
      mul_expr: prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 55:
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
      optional_label_decl -> State 140
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 141
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 56:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 57:
    Kernel Items:
      lex_internal_tokens: COMMENT., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      lex_internal_tokens: SPACES., $
    Reduce:
      $ -> [lex_internal_tokens]
    Goto:
      (nil)

  State 59:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 142

  State 60:
    Kernel Items:
      definitions: definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 61:
    Kernel Items:
      optional_definitions: optional_newlines definitions.optional_newlines
      definitions: definitions.NEWLINES definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      NEWLINES -> State 143
      optional_newlines -> State 144

  State 62:
    Kernel Items:
      definition: named_func_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 63:
    Kernel Items:
      definition: package_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 64:
    Kernel Items:
      definition: type_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      definition: unsafe_statement., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 66:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER., *
      package_def: PACKAGE IDENTIFIER.LPAREN package_statements RPAREN
    Reduce:
      * -> [package_def]
    Goto:
      LPAREN -> State 145

  State 67:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE IDENTIFIER.EQUAL value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 146
      EQUAL -> State 147
      optional_generic_parameters -> State 148

  State 68:
    Kernel Items:
      optional_receiver: LPAREN.parameter_def RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      parameter_def -> State 150

  State 69:
    Kernel Items:
      named_func_def: FUNC optional_receiver.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 151

  State 70:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 149
      parameter_def -> State 153
      parameter_defs -> State 154
      optional_parameter_defs -> State 152

  State 71:
    Kernel Items:
      returnable_type: BIT_AND.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 155
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 72:
    Kernel Items:
      returnable_type: BIT_NEG.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 156
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 73:
    Kernel Items:
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 157

  State 74:
    Kernel Items:
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 158

  State 75:
    Kernel Items:
      returnable_type: EXCLAIM.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 159
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 76:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 160

  State 77:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOT -> State 161
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 162

  State 78:
    Kernel Items:
      atom_type: LEX_ERROR., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 79:
    Kernel Items:
      implicit_struct_def: LPAREN.optional_implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      IDENTIFIER -> State 163
      UNSAFE -> State 59
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      unsafe_statement -> State 169
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 170
      field_def -> State 165
      implicit_field_defs -> State 167
      optional_implicit_field_defs -> State 168
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      enum_value_def -> State 164
      implicit_enum_value_defs -> State 166
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 80:
    Kernel Items:
      returnable_type: QUESTION.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 171
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 81:
    Kernel Items:
      returnable_type: TILDE_TILDE.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 172
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 82:
    Kernel Items:
      trait_def: TRAIT.LPAREN optional_trait_properties RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 173

  State 83:
    Kernel Items:
      returnable_type: atom_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 84:
    Kernel Items:
      atom_type: explicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 85:
    Kernel Items:
      atom_type: func_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 86:
    Kernel Items:
      atom_type: implicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 87:
    Kernel Items:
      atom_type: implicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      atom_type: initializable_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      value_type: returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 90:
    Kernel Items:
      atom_type: trait_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 91:
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
      RBRACKET -> State 178
      COMMA -> State 176
      COLON -> State 175
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 92:
    Kernel Items:
      argument: DOTDOTDOT., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 93:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expression
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      ASSIGN -> State 180

  State 94:
    Kernel Items:
      arguments: argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 95:
    Kernel Items:
      arguments: arguments.COMMA argument
      implicit_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 182
      COMMA -> State 181

  State 96:
    Kernel Items:
      argument: colon_expressions., *
      colon_expressions: colon_expressions.COLON optional_expression
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 183

  State 97:
    Kernel Items:
      argument: expression., *
      optional_expression: expression., COLON
    Reduce:
      * -> [argument]
      COLON -> [optional_expression]
    Goto:
      (nil)

  State 98:
    Kernel Items:
      colon_expressions: optional_expression.COLON optional_expression
    Reduce:
      (nil)
    Goto:
      COLON -> State 184

  State 99:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN.optional_explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      IDENTIFIER -> State 163
      UNSAFE -> State 59
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      unsafe_statement -> State 169
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 170
      field_def -> State 186
      implicit_struct_def -> State 87
      explicit_field_defs -> State 185
      optional_explicit_field_defs -> State 187
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 100:
    Kernel Items:
      var_pattern: IDENTIFIER., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 101:
    Kernel Items:
      tuple_pattern: LPAREN.field_var_patterns RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 189
      LPAREN -> State 101
      DOTDOTDOT -> State 188
      var_pattern -> State 192
      tuple_pattern -> State 102
      field_var_patterns -> State 191
      field_var_pattern -> State 190

  State 102:
    Kernel Items:
      var_pattern: tuple_pattern., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 103:
    Kernel Items:
      expression: VAR var_pattern.optional_value_type
    Reduce:
      * -> [optional_value_type]
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      optional_value_type -> State 193
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 194
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 104:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      optional_generic_arguments -> State 196
      generic_arguments -> State 195
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 197
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 105:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 198

  State 106:
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
      IDENTIFIER -> State 93
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      VAR -> State 37
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      DOTDOTDOT -> State 92
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 97
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      argument -> State 199
      colon_expressions -> State 96
      optional_expression -> State 98
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 107:
    Kernel Items:
      postfix_unary_expr: access_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 108:
    Kernel Items:
      call_expr: access_expr optional_generic_binding.LPAREN optional_arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 200

  State 109:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 110:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 111:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 112:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 113:
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
      optional_label_decl -> State 140
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 201
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 114:
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
      optional_label_decl -> State 140
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 202
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 115:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 117:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 118:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 119:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 121:
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
      optional_label_decl -> State 140
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 203
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 122:
    Kernel Items:
      atom_expr: initializable_type LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 25
      FLOAT_LITERAL -> State 22
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 93
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      VAR -> State 37
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      DOTDOTDOT -> State 92
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 97
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      arguments -> State 204
      argument -> State 94
      colon_expressions -> State 96
      optional_expression -> State 98
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 123:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 126:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 127:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 128:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 129:
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
      optional_label_decl -> State 140
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 205
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 130:
    Kernel Items:
      loop_expr: DO.block_body
      loop_expr: DO.block_body FOR sequence_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 133
      block_body -> State 206

  State 131:
    Kernel Items:
      loop_expr: FOR.sequence_expr DO block_body
      loop_expr: FOR.pattern IN sequence_expr DO block_body
      loop_expr: FOR.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body
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
      VAR -> State 209
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      DOT -> State 207
      SEMICOLON -> State 208
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 210
      optional_label_decl -> State 51
      sequence_expr -> State 212
      block_expr -> State 43
      pattern -> State 211
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 132:
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
      optional_label_decl -> State 140
      sequence_expr -> State 213
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 133:
    Kernel Items:
      block_body: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 214

  State 134:
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
      optional_label_decl -> State 140
      sequence_expr -> State 215
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 135:
    Kernel Items:
      block_expr: optional_label_decl block_body., *
    Reduce:
      * -> [block_expr]
    Goto:
      (nil)

  State 136:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 137:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 138:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 139:
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
      optional_label_decl -> State 140
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 216
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 140:
    Kernel Items:
      block_expr: optional_label_decl.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 133
      block_body -> State 135

  State 141:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 142:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 217

  State 143:
    Kernel Items:
      optional_newlines: NEWLINES., $
      definitions: definitions NEWLINES.definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      PACKAGE -> State 16
      UNSAFE -> State 59
      TYPE -> State 17
      FUNC -> State 18
      definition -> State 218
      unsafe_statement -> State 65
      type_def -> State 64
      named_func_def -> State 62
      package_def -> State 63

  State 144:
    Kernel Items:
      optional_definitions: optional_newlines definitions optional_newlines., $
    Reduce:
      $ -> [optional_definitions]
    Goto:
      (nil)

  State 145:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN.package_statements RPAREN
    Reduce:
      * -> [package_statements]
    Goto:
      package_statements -> State 219

  State 146:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.optional_generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 220
      generic_parameter_def -> State 221
      generic_parameter_defs -> State 222
      optional_generic_parameter_defs -> State 223

  State 147:
    Kernel Items:
      type_def: TYPE IDENTIFIER EQUAL.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 224
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 148:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 225
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 149:
    Kernel Items:
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      DOTDOTDOT -> State 226
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 227
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 150:
    Kernel Items:
      optional_receiver: LPAREN parameter_def.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 228

  State 151:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 146
      optional_generic_parameters -> State 229

  State 152:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 230

  State 153:
    Kernel Items:
      parameter_defs: parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 154:
    Kernel Items:
      parameter_defs: parameter_defs.COMMA parameter_def
      optional_parameter_defs: parameter_defs., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      COMMA -> State 231

  State 155:
    Kernel Items:
      returnable_type: BIT_AND returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 156:
    Kernel Items:
      returnable_type: BIT_NEG returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 157:
    Kernel Items:
      atom_type: DOT optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 158:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 163
      UNSAFE -> State 59
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      unsafe_statement -> State 169
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 170
      field_def -> State 234
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      enum_value_def -> State 232
      implicit_enum_value_defs -> State 235
      implicit_enum_def -> State 86
      explicit_enum_value_defs -> State 233
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 159:
    Kernel Items:
      returnable_type: EXCLAIM returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 160:
    Kernel Items:
      func_type: FUNC LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 237
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      DOTDOTDOT -> State 236
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 241
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      parameter_decl -> State 239
      parameter_decls -> State 240
      optional_parameter_decls -> State 238
      func_type -> State 85

  State 161:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 242

  State 162:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 163:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 243
      QUESTION -> State 80
      EXCLAIM -> State 75
      DOLLAR_LBRACKET -> State 104
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      optional_generic_binding -> State 162
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 244
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 164:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 245

  State 165:
    Kernel Items:
      implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 246

  State 166:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 248
      OR -> State 247

  State 167:
    Kernel Items:
      implicit_field_defs: implicit_field_defs.COMMA field_def
      optional_implicit_field_defs: implicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      COMMA -> State 249

  State 168:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 250

  State 169:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 170:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 171:
    Kernel Items:
      returnable_type: QUESTION returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 172:
    Kernel Items:
      returnable_type: TILDE_TILDE returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 173:
    Kernel Items:
      trait_def: TRAIT LPAREN.optional_trait_properties RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      IDENTIFIER -> State 163
      UNSAFE -> State 59
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 251
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      unsafe_statement -> State 169
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 170
      field_def -> State 252
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_property -> State 256
      trait_properties -> State 255
      optional_trait_properties -> State 254
      trait_def -> State 90
      func_type -> State 85
      method_signature -> State 253

  State 174:
    Kernel Items:
      value_type: value_type ADD.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 257
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 175:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON.value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 258
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 176:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 259

  State 177:
    Kernel Items:
      value_type: value_type MUL.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 260
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 178:
    Kernel Items:
      initializable_type: LBRACKET value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 179:
    Kernel Items:
      value_type: value_type SUB.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 261
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 180:
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
      VAR -> State 37
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 262
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 181:
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
      IDENTIFIER -> State 93
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      VAR -> State 37
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      DOTDOTDOT -> State 92
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 97
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      argument -> State 263
      colon_expressions -> State 96
      optional_expression -> State 98
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 182:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [implicit_struct_expr]
    Goto:
      (nil)

  State 183:
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
      VAR -> State 37
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 264
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      optional_expression -> State 265
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 184:
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
      VAR -> State 37
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 264
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      optional_expression -> State 266
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 185:
    Kernel Items:
      explicit_field_defs: explicit_field_defs.NEWLINES field_def
      explicit_field_defs: explicit_field_defs.COMMA field_def
      optional_explicit_field_defs: explicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      NEWLINES -> State 268
      COMMA -> State 267

  State 186:
    Kernel Items:
      explicit_field_defs: field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 187:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 269

  State 188:
    Kernel Items:
      field_var_pattern: DOTDOTDOT., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 189:
    Kernel Items:
      var_pattern: IDENTIFIER., *
      field_var_pattern: IDENTIFIER.ASSIGN var_pattern
    Reduce:
      * -> [var_pattern]
    Goto:
      ASSIGN -> State 270

  State 190:
    Kernel Items:
      field_var_patterns: field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 191:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns.RPAREN
      field_var_patterns: field_var_patterns.COMMA field_var_pattern
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 272
      COMMA -> State 271

  State 192:
    Kernel Items:
      field_var_pattern: var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 193:
    Kernel Items:
      expression: VAR var_pattern optional_value_type., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 194:
    Kernel Items:
      optional_value_type: value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [optional_value_type]
    Goto:
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 195:
    Kernel Items:
      optional_generic_arguments: generic_arguments., RBRACKET
      generic_arguments: generic_arguments.COMMA value_type
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 273

  State 196:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 274

  State 197:
    Kernel Items:
      generic_arguments: value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 198:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 199:
    Kernel Items:
      access_expr: access_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 275

  State 200:
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
      IDENTIFIER -> State 93
      TRUE -> State 36
      FALSE -> State 21
      STRUCT -> State 34
      FUNC -> State 23
      VAR -> State 37
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      DOTDOTDOT -> State 92
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 97
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      optional_arguments -> State 277
      arguments -> State 276
      argument -> State 94
      colon_expressions -> State 96
      optional_expression -> State 98
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 201:
    Kernel Items:
      mul_expr: mul_expr.mul_op prefix_unary_expr
      add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 128
      DIV -> State 126
      MOD -> State 127
      BIT_AND -> State 123
      BIT_LSHIFT -> State 124
      BIT_RSHIFT -> State 125
      mul_op -> State 129

  State 202:
    Kernel Items:
      cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 115
      NOT_EQUAL -> State 120
      LESS -> State 118
      LESS_OR_EQUAL -> State 119
      GREATER -> State 116
      GREATER_OR_EQUAL -> State 117
      cmp_op -> State 121

  State 203:
    Kernel Items:
      add_expr: add_expr.add_op mul_expr
      cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 109
      SUB -> State 112
      BIT_XOR -> State 111
      BIT_OR -> State 110
      add_op -> State 113

  State 204:
    Kernel Items:
      arguments: arguments.COMMA argument
      atom_expr: initializable_type LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 278
      COMMA -> State 181

  State 205:
    Kernel Items:
      mul_expr: mul_expr mul_op prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 206:
    Kernel Items:
      loop_expr: DO block_body., *
      loop_expr: DO block_body.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 279

  State 207:
    Kernel Items:
      pattern: DOT.IDENTIFIER implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 280

  State 208:
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
      optional_label_decl -> State 140
      optional_sequence_expr -> State 281
      sequence_expr -> State 282
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 209:
    Kernel Items:
      expression: VAR.var_pattern optional_value_type
      pattern: VAR.DOT IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 100
      LPAREN -> State 101
      DOT -> State 283
      var_pattern -> State 103
      tuple_pattern -> State 102

  State 210:
    Kernel Items:
      pattern: expression., *
    Reduce:
      * -> [pattern]
    Goto:
      (nil)

  State 211:
    Kernel Items:
      loop_expr: FOR pattern.IN sequence_expr DO block_body
    Reduce:
      (nil)
    Goto:
      IN -> State 284

  State 212:
    Kernel Items:
      expression: sequence_expr., IN
      loop_expr: FOR sequence_expr.DO block_body
    Reduce:
      IN -> [expression]
    Goto:
      DO -> State 285

  State 213:
    Kernel Items:
      if_expr: IF sequence_expr.block_body
      if_expr: IF sequence_expr.block_body ELSE block_body
      if_expr: IF sequence_expr.block_body ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 133
      block_body -> State 286

  State 214:
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
      RETURN -> State 292
      BREAK -> State 288
      CONTINUE -> State 289
      UNSAFE -> State 59
      STRUCT -> State 34
      FUNC -> State 23
      ASYNC -> State 287
      DEFER -> State 290
      VAR -> State 209
      LABEL_DECL -> State 26
      RBRACE -> State 291
      LPAREN -> State 29
      LBRACKET -> State 27
      DOT -> State 207
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 294
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      statement -> State 299
      statement_body -> State 300
      pattern -> State 298
      unsafe_statement -> State 301
      jump_statement -> State 296
      jump_type -> State 297
      expressions -> State 295
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 293
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 215:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.LBRACE case_branches optional_default_branch RBRACE
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 302

  State 216:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 114

  State 217:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 303

  State 218:
    Kernel Items:
      definitions: definitions NEWLINES definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 219:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements.RPAREN
      package_statements: package_statements.package_statement
    Reduce:
      (nil)
    Goto:
      UNSAFE -> State 59
      RPAREN -> State 304
      unsafe_statement -> State 307
      package_statement_body -> State 306
      package_statement -> State 305

  State 220:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 308
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 221:
    Kernel Items:
      generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 222:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs.COMMA generic_parameter_def
      optional_generic_parameter_defs: generic_parameter_defs., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      COMMA -> State 309

  State 223:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 310

  State 224:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER EQUAL value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 225:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., *
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 311
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 226:
    Kernel Items:
      parameter_def: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 312
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 227:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 228:
    Kernel Items:
      optional_receiver: LPAREN parameter_def RPAREN., IDENTIFIER
    Reduce:
      IDENTIFIER -> [optional_receiver]
    Goto:
      (nil)

  State 229:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 313

  State 230:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 315
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      return_type -> State 314
      func_type -> State 85

  State 231:
    Kernel Items:
      parameter_defs: parameter_defs COMMA.parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 149
      parameter_def -> State 316

  State 232:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 317
      OR -> State 318

  State 233:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 319

  State 234:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 246

  State 235:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 320
      OR -> State 321

  State 236:
    Kernel Items:
      parameter_decl: DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 322
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 237:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.DOTDOTDOT value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 243
      QUESTION -> State 80
      EXCLAIM -> State 75
      DOLLAR_LBRACKET -> State 104
      DOTDOTDOT -> State 323
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      optional_generic_binding -> State 162
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 324
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 238:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 325

  State 239:
    Kernel Items:
      parameter_decls: parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      parameter_decls: parameter_decls.COMMA parameter_decl
      optional_parameter_decls: parameter_decls., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      COMMA -> State 326

  State 241:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 242:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 327

  State 243:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 242
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 157

  State 244:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 245:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 163
      UNSAFE -> State 59
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      unsafe_statement -> State 169
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 170
      field_def -> State 234
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      enum_value_def -> State 328
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 246:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 329

  State 247:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 163
      UNSAFE -> State 59
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      unsafe_statement -> State 169
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 170
      field_def -> State 234
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      enum_value_def -> State 330
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

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
      IDENTIFIER -> State 163
      UNSAFE -> State 59
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      unsafe_statement -> State 169
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 170
      field_def -> State 331
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 250:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_def]
    Goto:
      (nil)

  State 251:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 332
      LPAREN -> State 160

  State 252:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 253:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 254:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 333

  State 255:
    Kernel Items:
      trait_properties: trait_properties.NEWLINES trait_property
      trait_properties: trait_properties.COMMA trait_property
      optional_trait_properties: trait_properties., RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      NEWLINES -> State 335
      COMMA -> State 334

  State 256:
    Kernel Items:
      trait_properties: trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 257:
    Kernel Items:
      value_type: value_type ADD returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 258:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type.RBRACKET
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 336
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 259:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 337

  State 260:
    Kernel Items:
      value_type: value_type MUL returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 261:
    Kernel Items:
      value_type: value_type SUB returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 262:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 263:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 264:
    Kernel Items:
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      (nil)

  State 265:
    Kernel Items:
      colon_expressions: colon_expressions COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 266:
    Kernel Items:
      colon_expressions: optional_expression COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 267:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 163
      UNSAFE -> State 59
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      unsafe_statement -> State 169
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 170
      field_def -> State 338
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 268:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 163
      UNSAFE -> State 59
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      unsafe_statement -> State 169
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 170
      field_def -> State 339
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 269:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_def]
    Goto:
      (nil)

  State 270:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN.var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 100
      LPAREN -> State 101
      var_pattern -> State 340
      tuple_pattern -> State 102

  State 271:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA.field_var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 189
      LPAREN -> State 101
      DOTDOTDOT -> State 188
      var_pattern -> State 192
      tuple_pattern -> State 102
      field_var_pattern -> State 341

  State 272:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns RPAREN., *
    Reduce:
      * -> [tuple_pattern]
    Goto:
      (nil)

  State 273:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 342
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 274:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 275:
    Kernel Items:
      access_expr: access_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 276:
    Kernel Items:
      optional_arguments: arguments., RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 181

  State 277:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 343

  State 278:
    Kernel Items:
      atom_expr: initializable_type LPAREN arguments RPAREN., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 279:
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
      optional_label_decl -> State 140
      sequence_expr -> State 344
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 280:
    Kernel Items:
      pattern: DOT IDENTIFIER.implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 29
      implicit_struct_expr -> State 345

  State 281:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO block_body
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 346

  State 282:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 283:
    Kernel Items:
      pattern: VAR DOT.IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 347

  State 284:
    Kernel Items:
      loop_expr: FOR pattern IN.sequence_expr DO block_body
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
      optional_label_decl -> State 140
      sequence_expr -> State 348
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 285:
    Kernel Items:
      loop_expr: FOR sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 133
      block_body -> State 349

  State 286:
    Kernel Items:
      if_expr: IF sequence_expr block_body., *
      if_expr: IF sequence_expr block_body.ELSE block_body
      if_expr: IF sequence_expr block_body.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 350

  State 287:
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
      optional_label_decl -> State 140
      block_expr -> State 43
      call_expr -> State 352
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 351
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 288:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 289:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 290:
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
      optional_label_decl -> State 140
      block_expr -> State 43
      call_expr -> State 353
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 351
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 291:
    Kernel Items:
      block_body: LBRACE statements RBRACE., $
    Reduce:
      $ -> [block_body]
    Goto:
      (nil)

  State 292:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 293:
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
      LBRACKET -> State 106
      DOT -> State 105
      QUESTION -> State 107
      DOLLAR_LBRACKET -> State 104
      ADD_ASSIGN -> State 354
      SUB_ASSIGN -> State 365
      MUL_ASSIGN -> State 364
      DIV_ASSIGN -> State 362
      MOD_ASSIGN -> State 363
      ADD_ONE_ASSIGN -> State 355
      SUB_ONE_ASSIGN -> State 366
      BIT_NEG_ASSIGN -> State 358
      BIT_AND_ASSIGN -> State 356
      BIT_OR_ASSIGN -> State 359
      BIT_XOR_ASSIGN -> State 361
      BIT_LSHIFT_ASSIGN -> State 357
      BIT_RSHIFT_ASSIGN -> State 360
      unary_op_assign -> State 368
      binary_op_assign -> State 367
      optional_generic_binding -> State 108

  State 294:
    Kernel Items:
      pattern: expression., ASSIGN
      expressions: expression., *
    Reduce:
      * -> [expressions]
      ASSIGN -> [pattern]
    Goto:
      (nil)

  State 295:
    Kernel Items:
      statement_body: expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [statement_body]
    Goto:
      COMMA -> State 369

  State 296:
    Kernel Items:
      statement_body: jump_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 297:
    Kernel Items:
      jump_statement: jump_type.optional_jump_label optional_expressions
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 370
      optional_jump_label -> State 371

  State 298:
    Kernel Items:
      statement_body: pattern.ASSIGN expression
    Reduce:
      (nil)
    Goto:
      ASSIGN -> State 372

  State 299:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 300:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 373
      SEMICOLON -> State 374

  State 301:
    Kernel Items:
      statement_body: unsafe_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 302:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE.case_branches optional_default_branch RBRACE
    Reduce:
      (nil)
    Goto:
      CASE -> State 375
      case_branches -> State 377
      case_branch -> State 376

  State 303:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 378

  State 304:
    Kernel Items:
      package_def: PACKAGE IDENTIFIER LPAREN package_statements RPAREN., $
    Reduce:
      $ -> [package_def]
    Goto:
      (nil)

  State 305:
    Kernel Items:
      package_statements: package_statements package_statement., *
    Reduce:
      * -> [package_statements]
    Goto:
      (nil)

  State 306:
    Kernel Items:
      package_statement: package_statement_body.NEWLINES
      package_statement: package_statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 379
      SEMICOLON -> State 380

  State 307:
    Kernel Items:
      package_statement_body: unsafe_statement., *
    Reduce:
      * -> [package_statement_body]
    Goto:
      (nil)

  State 308:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 309:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA.generic_parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 220
      generic_parameter_def -> State 381

  State 310:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 311:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 382
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 312:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_def: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 313:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type block_body
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 149
      parameter_def -> State 153
      parameter_defs -> State 154
      optional_parameter_defs -> State 383

  State 314:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 133
      block_body -> State 384

  State 315:
    Kernel Items:
      return_type: returnable_type., *
    Reduce:
      * -> [return_type]
    Goto:
      (nil)

  State 316:
    Kernel Items:
      parameter_defs: parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 317:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 163
      UNSAFE -> State 59
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      unsafe_statement -> State 169
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 170
      field_def -> State 234
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      enum_value_def -> State 385
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 318:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 163
      UNSAFE -> State 59
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      unsafe_statement -> State 169
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 170
      field_def -> State 234
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      enum_value_def -> State 386
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 319:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_def]
    Goto:
      (nil)

  State 320:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 163
      UNSAFE -> State 59
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      unsafe_statement -> State 169
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 170
      field_def -> State 234
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      enum_value_def -> State 387
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 321:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 163
      UNSAFE -> State 59
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      unsafe_statement -> State 169
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 170
      field_def -> State 234
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      enum_value_def -> State 388
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 322:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 323:
    Kernel Items:
      parameter_decl: IDENTIFIER DOTDOTDOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 389
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      func_type -> State 85

  State 324:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 325:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 315
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      return_type -> State 390
      func_type -> State 85

  State 326:
    Kernel Items:
      parameter_decls: parameter_decls COMMA.parameter_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 237
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      DOTDOTDOT -> State 236
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 241
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      parameter_decl -> State 391
      func_type -> State 85

  State 327:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 328:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 329:
    Kernel Items:
      enum_value_def: field_def ASSIGN DEFAULT., *
    Reduce:
      * -> [enum_value_def]
    Goto:
      (nil)

  State 330:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 331:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [implicit_field_defs]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 392

  State 333:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties RPAREN., *
    Reduce:
      * -> [trait_def]
    Goto:
      (nil)

  State 334:
    Kernel Items:
      trait_properties: trait_properties COMMA.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 163
      UNSAFE -> State 59
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 251
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      unsafe_statement -> State 169
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 170
      field_def -> State 252
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_property -> State 393
      trait_def -> State 90
      func_type -> State 85
      method_signature -> State 253

  State 335:
    Kernel Items:
      trait_properties: trait_properties NEWLINES.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 163
      UNSAFE -> State 59
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 251
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      unsafe_statement -> State 169
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 170
      field_def -> State 252
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_property -> State 394
      trait_def -> State 90
      func_type -> State 85
      method_signature -> State 253

  State 336:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 337:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 338:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 339:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 340:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 342:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 343:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 344:
    Kernel Items:
      loop_expr: DO block_body FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 345:
    Kernel Items:
      pattern: DOT IDENTIFIER implicit_struct_expr., ASSIGN
    Reduce:
      ASSIGN -> [pattern]
    Goto:
      (nil)

  State 346:
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
      optional_label_decl -> State 140
      optional_sequence_expr -> State 395
      sequence_expr -> State 282
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 347:
    Kernel Items:
      pattern: VAR DOT IDENTIFIER.tuple_pattern
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 101
      tuple_pattern -> State 396

  State 348:
    Kernel Items:
      loop_expr: FOR pattern IN sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 397

  State 349:
    Kernel Items:
      loop_expr: FOR sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 350:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE.block_body
      if_expr: IF sequence_expr block_body ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 132
      LBRACE -> State 133
      if_expr -> State 399
      block_body -> State 398

  State 351:
    Kernel Items:
      call_expr: access_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: access_expr.DOT IDENTIFIER
      access_expr: access_expr.LBRACKET argument RBRACKET
    Reduce:
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 106
      DOT -> State 105
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 108

  State 352:
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

  State 353:
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

  State 354:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 355:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 356:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 357:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 358:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 359:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 360:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 361:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 362:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 363:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 364:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 365:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 366:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 367:
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
      VAR -> State 37
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 400
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 368:
    Kernel Items:
      statement_body: access_expr unary_op_assign., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 369:
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
      VAR -> State 37
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 401
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 370:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 371:
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
      VAR -> State 37
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 402
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      expressions -> State 403
      optional_expressions -> State 404
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 372:
    Kernel Items:
      statement_body: pattern ASSIGN.expression
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
      VAR -> State 37
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 405
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 373:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 374:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 375:
    Kernel Items:
      case_branch: CASE.patterns COLON sequence_expr
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
      VAR -> State 209
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      DOT -> State 207
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 210
      optional_label_decl -> State 51
      patterns -> State 407
      sequence_expr -> State 56
      block_expr -> State 43
      pattern -> State 406
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 376:
    Kernel Items:
      case_branches: case_branch., *
    Reduce:
      * -> [case_branches]
    Goto:
      (nil)

  State 377:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE case_branches.optional_default_branch RBRACE
      case_branches: case_branches.case_branch
    Reduce:
      RBRACE -> [optional_default_branch]
    Goto:
      CASE -> State 375
      DEFAULT -> State 408
      case_branch -> State 409
      optional_default_branch -> State 411
      default_branch -> State 410

  State 378:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 379:
    Kernel Items:
      package_statement: package_statement_body NEWLINES., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 380:
    Kernel Items:
      package_statement: package_statement_body SEMICOLON., *
    Reduce:
      * -> [package_statement]
    Goto:
      (nil)

  State 381:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 382:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 383:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type block_body
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 412

  State 384:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type block_body., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 385:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 386:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
      explicit_enum_value_defs: enum_value_def OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 387:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 388:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 389:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: IDENTIFIER DOTDOTDOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 174
      SUB -> State 179
      MUL -> State 177

  State 390:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [func_type]
    Goto:
      (nil)

  State 391:
    Kernel Items:
      parameter_decls: parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 392:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 237
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      DOTDOTDOT -> State 236
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 89
      value_type -> State 241
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      parameter_decl -> State 239
      parameter_decls -> State 240
      optional_parameter_decls -> State 413
      func_type -> State 85

  State 393:
    Kernel Items:
      trait_properties: trait_properties COMMA trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 394:
    Kernel Items:
      trait_properties: trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 395:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO block_body
    Reduce:
      (nil)
    Goto:
      DO -> State 414

  State 396:
    Kernel Items:
      pattern: VAR DOT IDENTIFIER tuple_pattern., ASSIGN
    Reduce:
      ASSIGN -> [pattern]
    Goto:
      (nil)

  State 397:
    Kernel Items:
      loop_expr: FOR pattern IN sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 133
      block_body -> State 415

  State 398:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE block_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 399:
    Kernel Items:
      if_expr: IF sequence_expr block_body ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 400:
    Kernel Items:
      statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 401:
    Kernel Items:
      expressions: expressions COMMA expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 402:
    Kernel Items:
      expressions: expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 403:
    Kernel Items:
      expressions: expressions.COMMA expression
      optional_expressions: expressions., *
    Reduce:
      * -> [optional_expressions]
    Goto:
      COMMA -> State 369

  State 404:
    Kernel Items:
      jump_statement: jump_type optional_jump_label optional_expressions., *
    Reduce:
      * -> [jump_statement]
    Goto:
      (nil)

  State 405:
    Kernel Items:
      statement_body: pattern ASSIGN expression., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 406:
    Kernel Items:
      patterns: pattern., *
    Reduce:
      * -> [patterns]
    Goto:
      (nil)

  State 407:
    Kernel Items:
      case_branch: CASE patterns.COLON sequence_expr
      patterns: patterns.COMMA pattern
    Reduce:
      (nil)
    Goto:
      COMMA -> State 417
      COLON -> State 416

  State 408:
    Kernel Items:
      default_branch: DEFAULT.COLON sequence_expr
    Reduce:
      (nil)
    Goto:
      COLON -> State 418

  State 409:
    Kernel Items:
      case_branches: case_branches case_branch., *
    Reduce:
      * -> [case_branches]
    Goto:
      (nil)

  State 410:
    Kernel Items:
      optional_default_branch: default_branch., RBRACE
    Reduce:
      RBRACE -> [optional_default_branch]
    Goto:
      (nil)

  State 411:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE case_branches optional_default_branch.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 419

  State 412:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type block_body
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 315
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      return_type -> State 420
      func_type -> State 85

  State 413:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 421

  State 414:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 133
      block_body -> State 422

  State 415:
    Kernel Items:
      loop_expr: FOR pattern IN sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 416:
    Kernel Items:
      case_branch: CASE patterns COLON.sequence_expr
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
      optional_label_decl -> State 140
      sequence_expr -> State 423
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 417:
    Kernel Items:
      patterns: patterns COMMA.pattern
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
      VAR -> State 209
      LABEL_DECL -> State 26
      LPAREN -> State 29
      LBRACKET -> State 27
      DOT -> State 207
      NOT -> State 31
      SUB -> State 35
      MUL -> State 30
      BIT_NEG -> State 20
      BIT_AND -> State 19
      LEX_ERROR -> State 28
      expression -> State 210
      optional_label_decl -> State 51
      sequence_expr -> State 56
      block_expr -> State 43
      pattern -> State 424
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 418:
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
      optional_label_decl -> State 140
      sequence_expr -> State 425
      block_expr -> State 43
      call_expr -> State 44
      atom_expr -> State 42
      literal -> State 49
      implicit_struct_expr -> State 47
      access_expr -> State 38
      postfix_unary_expr -> State 53
      prefix_unary_op -> State 55
      prefix_unary_expr -> State 54
      mul_expr -> State 50
      add_expr -> State 39
      cmp_expr -> State 45
      and_expr -> State 40
      or_expr -> State 52
      initializable_type -> State 48
      explicit_struct_def -> State 46
      anonymous_func_expr -> State 41

  State 419:
    Kernel Items:
      switch_expr: SWITCH sequence_expr LBRACE case_branches optional_default_branch RBRACE., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 420:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.block_body
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 133
      block_body -> State 426

  State 421:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 77
      STRUCT -> State 34
      ENUM -> State 74
      TRAIT -> State 82
      FUNC -> State 76
      LPAREN -> State 79
      LBRACKET -> State 27
      DOT -> State 73
      QUESTION -> State 80
      EXCLAIM -> State 75
      TILDE_TILDE -> State 81
      BIT_NEG -> State 72
      BIT_AND -> State 71
      LEX_ERROR -> State 78
      initializable_type -> State 88
      atom_type -> State 83
      returnable_type -> State 315
      implicit_struct_def -> State 87
      explicit_struct_def -> State 46
      implicit_enum_def -> State 86
      explicit_enum_def -> State 84
      trait_def -> State 90
      return_type -> State 427
      func_type -> State 85

  State 422:
    Kernel Items:
      loop_expr: FOR SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO block_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 423:
    Kernel Items:
      case_branch: CASE patterns COLON sequence_expr., *
    Reduce:
      * -> [case_branch]
    Goto:
      (nil)

  State 424:
    Kernel Items:
      patterns: patterns COMMA pattern., *
    Reduce:
      * -> [patterns]
    Goto:
      (nil)

  State 425:
    Kernel Items:
      default_branch: DEFAULT COLON sequence_expr., RBRACE
    Reduce:
      RBRACE -> [default_branch]
    Goto:
      (nil)

  State 426:
    Kernel Items:
      named_func_def: FUNC optional_receiver IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type block_body., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 427:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

Number of states: 427
Number of shift actions: 2771
Number of reduce actions: 342
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 3660
Number of unoptimized shift actions: 27260
Number of unoptimized reduce actions: 22249
*/
