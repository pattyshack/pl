// Auto-generated from source: grammar.lr

package parser

import (
	fmt "fmt"
	io "io"
	sort "sort"
)

type SymbolId int

const (
	NewlinesToken        = SymbolId(256)
	CommentGroupsToken   = SymbolId(257)
	IntegerLiteralToken  = SymbolId(258)
	FloatLiteralToken    = SymbolId(259)
	RuneLiteralToken     = SymbolId(260)
	StringLiteralToken   = SymbolId(261)
	IdentifierToken      = SymbolId(262)
	TrueToken            = SymbolId(263)
	FalseToken           = SymbolId(264)
	IfToken              = SymbolId(265)
	ElseToken            = SymbolId(266)
	SwitchToken          = SymbolId(267)
	CaseToken            = SymbolId(268)
	DefaultToken         = SymbolId(269)
	ForToken             = SymbolId(270)
	DoToken              = SymbolId(271)
	InToken              = SymbolId(272)
	ReturnToken          = SymbolId(273)
	BreakToken           = SymbolId(274)
	ContinueToken        = SymbolId(275)
	FallthroughToken     = SymbolId(276)
	PackageToken         = SymbolId(277)
	ImportToken          = SymbolId(278)
	AsToken              = SymbolId(279)
	UnsafeToken          = SymbolId(280)
	TypeToken            = SymbolId(281)
	ImplementsToken      = SymbolId(282)
	StructToken          = SymbolId(283)
	EnumToken            = SymbolId(284)
	TraitToken           = SymbolId(285)
	FuncToken            = SymbolId(286)
	AsyncToken           = SymbolId(287)
	DeferToken           = SymbolId(288)
	VarToken             = SymbolId(289)
	LetToken             = SymbolId(290)
	NotToken             = SymbolId(291)
	AndToken             = SymbolId(292)
	OrToken              = SymbolId(293)
	LabelDeclToken       = SymbolId(294)
	JumpLabelToken       = SymbolId(295)
	LbraceToken          = SymbolId(296)
	RbraceToken          = SymbolId(297)
	LparenToken          = SymbolId(298)
	RparenToken          = SymbolId(299)
	LbracketToken        = SymbolId(300)
	RbracketToken        = SymbolId(301)
	DotToken             = SymbolId(302)
	CommaToken           = SymbolId(303)
	QuestionToken        = SymbolId(304)
	SemicolonToken       = SymbolId(305)
	ColonToken           = SymbolId(306)
	ExclaimToken         = SymbolId(307)
	DollarLbracketToken  = SymbolId(308)
	DotDotDotToken       = SymbolId(309)
	TildeTildeToken      = SymbolId(310)
	AssignToken          = SymbolId(311)
	AddAssignToken       = SymbolId(312)
	SubAssignToken       = SymbolId(313)
	MulAssignToken       = SymbolId(314)
	DivAssignToken       = SymbolId(315)
	ModAssignToken       = SymbolId(316)
	AddOneAssignToken    = SymbolId(317)
	SubOneAssignToken    = SymbolId(318)
	BitNegAssignToken    = SymbolId(319)
	BitAndAssignToken    = SymbolId(320)
	BitOrAssignToken     = SymbolId(321)
	BitXorAssignToken    = SymbolId(322)
	BitLshiftAssignToken = SymbolId(323)
	BitRshiftAssignToken = SymbolId(324)
	AddToken             = SymbolId(325)
	SubToken             = SymbolId(326)
	MulToken             = SymbolId(327)
	DivToken             = SymbolId(328)
	ModToken             = SymbolId(329)
	BitNegToken          = SymbolId(330)
	BitAndToken          = SymbolId(331)
	BitXorToken          = SymbolId(332)
	BitOrToken           = SymbolId(333)
	BitLshiftToken       = SymbolId(334)
	BitRshiftToken       = SymbolId(335)
	EqualToken           = SymbolId(336)
	NotEqualToken        = SymbolId(337)
	LessToken            = SymbolId(338)
	LessOrEqualToken     = SymbolId(339)
	GreaterToken         = SymbolId(340)
	GreaterOrEqualToken  = SymbolId(341)
	ParseErrorToken      = SymbolId(342)
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
	End() Location
}

type GenericSymbol struct {
	SymbolId
	StartPos Location
	EndPos   Location
}

func (t GenericSymbol) Id() SymbolId { return t.SymbolId }

func (t GenericSymbol) Loc() Location { return t.StartPos }

func (t GenericSymbol) End() Location { return t.EndPos }

type Lexer interface {
	// Note: Return io.EOF to indicate end of stream
	// Token with unspecified value type should return GenericSymbol
	Next() (Token, error)

	CurrentLocation() Location
}

type SourceReducer interface {
	// 52:29: source -> ...
	ToSource(OptionalDefinitions_ GenericSymbol) ([]SourceDefinition, error)
}

type OptionalDefinitionsReducer interface {
	// 55:2: optional_definitions -> NEWLINES: ...
	NewlinesToOptionalDefinitions(Newlines_ TokenCount) (GenericSymbol, error)

	// 56:2: optional_definitions -> definitions: ...
	DefinitionsToOptionalDefinitions(OptionalNewlines_ GenericSymbol, Definitions_ []SourceDefinition, OptionalNewlines_2 GenericSymbol) (GenericSymbol, error)

	// 57:2: optional_definitions -> nil: ...
	NilToOptionalDefinitions() (GenericSymbol, error)
}

type OptionalNewlinesReducer interface {
	// 60:2: optional_newlines -> NEWLINES: ...
	NewlinesToOptionalNewlines(Newlines_ TokenCount) (GenericSymbol, error)

	// 61:2: optional_newlines -> nil: ...
	NilToOptionalNewlines() (GenericSymbol, error)
}

type DefinitionsReducer interface {
	// 64:2: definitions -> definition: ...
	DefinitionToDefinitions(Definition_ SourceDefinition) ([]SourceDefinition, error)

	// 65:2: definitions -> add: ...
	AddToDefinitions(Definitions_ []SourceDefinition, Newlines_ TokenCount, Definition_ SourceDefinition) ([]SourceDefinition, error)
}

type DefinitionReducer interface {
	// 69:2: definition -> package_def: ...
	PackageDefToDefinition(PackageDef_ SourceDefinition) (SourceDefinition, error)

	// 70:2: definition -> type_def: ...
	TypeDefToDefinition(TypeDef_ SourceDefinition) (SourceDefinition, error)

	// 71:2: definition -> named_func_def: ...
	NamedFuncDefToDefinition(NamedFuncDef_ SourceDefinition) (SourceDefinition, error)

	// 72:2: definition -> global_var_def: ...
	GlobalVarDefToDefinition(VarDeclPattern_ GenericSymbol) (SourceDefinition, error)

	// 73:2: definition -> global_var_assignment: ...
	GlobalVarAssignmentToDefinition(VarDeclPattern_ GenericSymbol, Assign_ TokenValue, Expression_ Expression) (SourceDefinition, error)

	// 76:2: definition -> statement_block: ...
	StatementBlockToDefinition(StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 80:2: definition -> COMMENT_GROUPS: ...
	CommentGroupsToDefinition(CommentGroups_ TokenValue) (SourceDefinition, error)
}

type StatementBlockReducer interface {
	// 100:19: statement_block -> ...
	ToStatementBlock(Lbrace_ TokenValue, Statements_ []Statement, Rbrace_ TokenValue) (GenericSymbol, error)
}

type StatementsReducer interface {
	// 103:2: statements -> empty_list: ...
	EmptyListToStatements() ([]Statement, error)

	// 104:2: statements -> add: ...
	AddToStatements(Statements_ []Statement, Statement_ Statement) ([]Statement, error)
}

type StatementReducer interface {
	// 107:2: statement -> implicit: ...
	ImplicitToStatement(StatementBody_ Statement, Newlines_ TokenCount) (Statement, error)

	// 108:2: statement -> explicit: ...
	ExplicitToStatement(StatementBody_ Statement, Semicolon_ TokenValue) (Statement, error)
}

type StatementBodyReducer interface {

	// 146:2: statement_body -> case_branch_statement: ...
	CaseBranchStatementToStatementBody(Case_ TokenValue, CasePatterns_ GenericSymbol, Colon_ TokenValue, OptionalSimpleStatementBody_ Statement) (Statement, error)

	// 147:2: statement_body -> default_branch_statement: ...
	DefaultBranchStatementToStatementBody(Default_ TokenValue, Colon_ TokenValue, OptionalSimpleStatementBody_ Statement) (Statement, error)
}

type OptionalSimpleStatementBodyReducer interface {

	// 151:2: optional_simple_statement_body -> nil: ...
	NilToOptionalSimpleStatementBody() (Statement, error)
}

type ExpressionOrImproperStructStatementReducer interface {
	// 157:54: expression_or_improper_struct_statement -> ...
	ToExpressionOrImproperStructStatement(Expressions_ GenericSymbol) (Statement, error)
}

type CallbackOpStatementReducer interface {
	// 183:36: callback_op_statement -> ...
	ToCallbackOpStatement(CallbackOp_ TokenValue, CallExpr_ Expression) (Statement, error)
}

type UnsafeStatementReducer interface {
	// 212:31: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ TokenValue, Less_ TokenValue, Identifier_ TokenValue, Greater_ TokenValue, StringLiteral_ TokenValue) (Statement, error)
}

type JumpStatementReducer interface {
	// 219:2: jump_statement -> unlabeled_no_value: ...
	UnlabeledNoValueToJumpStatement(JumpType_ TokenValue) (Statement, error)

	// 220:2: jump_statement -> unlabeled_valued: ...
	UnlabeledValuedToJumpStatement(JumpType_ TokenValue, Expressions_ GenericSymbol) (Statement, error)

	// 221:2: jump_statement -> labeled_no_value: ...
	LabeledNoValueToJumpStatement(JumpType_ TokenValue, JumpLabel_ TokenValue) (Statement, error)

	// 222:2: jump_statement -> labeled_valued: ...
	LabeledValuedToJumpStatement(JumpType_ TokenValue, JumpLabel_ TokenValue, Expressions_ GenericSymbol) (Statement, error)
}

type ExpressionsReducer interface {
	// 230:2: expressions -> expression: ...
	ExpressionToExpressions(Expression_ Expression) (GenericSymbol, error)

	// 231:2: expressions -> add: ...
	AddToExpressions(Expressions_ GenericSymbol, Comma_ TokenValue, Expression_ Expression) (GenericSymbol, error)
}

type FallthroughStatementReducer interface {
	// 233:36: fallthrough_statement -> ...
	ToFallthroughStatement(Fallthrough_ TokenValue) (Statement, error)
}

type AssignStatementReducer interface {
	// 239:31: assign_statement -> ...
	ToAssignStatement(AssignPattern_ GenericSymbol, Assign_ TokenValue, Expression_ Expression) (Statement, error)
}

type UnaryOpAssignStatementReducer interface {
	// 241:40: unary_op_assign_statement -> ...
	ToUnaryOpAssignStatement(AccessibleExpr_ Expression, UnaryOpAssign_ TokenValue) (Statement, error)
}

type BinaryOpAssignStatementReducer interface {
	// 243:41: binary_op_assign_statement -> ...
	ToBinaryOpAssignStatement(AccessibleExpr_ Expression, BinaryOpAssign_ TokenValue, Expression_ Expression) (Statement, error)
}

type ImportStatementReducer interface {
	// 250:2: import_statement -> single: ...
	SingleToImportStatement(Import_ TokenValue, ImportClause_ GenericSymbol) (Statement, error)

	// 251:2: import_statement -> multiple: ...
	MultipleToImportStatement(Import_ TokenValue, Lparen_ TokenValue, ImportClauses_ GenericSymbol, Rparen_ TokenValue) (Statement, error)
}

type ImportClauseReducer interface {
	// 254:2: import_clause -> STRING_LITERAL: ...
	StringLiteralToImportClause(StringLiteral_ TokenValue) (GenericSymbol, error)

	// 255:2: import_clause -> alias: ...
	AliasToImportClause(StringLiteral_ TokenValue, As_ TokenValue, Identifier_ TokenValue) (GenericSymbol, error)
}

type ImportClauseTerminalReducer interface {
	// 258:2: import_clause_terminal -> implicit: ...
	ImplicitToImportClauseTerminal(ImportClause_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)

	// 259:2: import_clause_terminal -> explicit: ...
	ExplicitToImportClauseTerminal(ImportClause_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)
}

type ImportClausesReducer interface {
	// 262:2: import_clauses -> first: ...
	FirstToImportClauses(ImportClauseTerminal_ GenericSymbol) (GenericSymbol, error)

	// 263:2: import_clauses -> add: ...
	AddToImportClauses(ImportClauses_ GenericSymbol, ImportClauseTerminal_ GenericSymbol) (GenericSymbol, error)
}

type CasePatternsReducer interface {
	// 270:2: case_patterns -> case_pattern: ...
	CasePatternToCasePatterns(CasePattern_ GenericSymbol) (GenericSymbol, error)

	// 271:2: case_patterns -> multiple: ...
	MultipleToCasePatterns(CasePatterns_ GenericSymbol, Comma_ TokenValue, CasePattern_ GenericSymbol) (GenericSymbol, error)
}

type VarDeclPatternReducer interface {
	// 280:20: var_decl_pattern -> ...
	ToVarDeclPattern(VarOrLet_ TokenValue, VarPattern_ GenericSymbol, OptionalValueType_ TypeExpression) (GenericSymbol, error)
}

type VarPatternReducer interface {
	// 287:2: var_pattern -> IDENTIFIER: ...
	IdentifierToVarPattern(Identifier_ TokenValue) (GenericSymbol, error)

	// 288:2: var_pattern -> tuple_pattern: ...
	TuplePatternToVarPattern(TuplePattern_ GenericSymbol) (GenericSymbol, error)
}

type TuplePatternReducer interface {
	// 290:17: tuple_pattern -> ...
	ToTuplePattern(Lparen_ TokenValue, FieldVarPatterns_ GenericSymbol, Rparen_ TokenValue) (GenericSymbol, error)
}

type FieldVarPatternsReducer interface {
	// 293:2: field_var_patterns -> field_var_pattern: ...
	FieldVarPatternToFieldVarPatterns(FieldVarPattern_ GenericSymbol) (GenericSymbol, error)

	// 294:2: field_var_patterns -> add: ...
	AddToFieldVarPatterns(FieldVarPatterns_ GenericSymbol, Comma_ TokenValue, FieldVarPattern_ GenericSymbol) (GenericSymbol, error)
}

type FieldVarPatternReducer interface {
	// 297:2: field_var_pattern -> positional: ...
	PositionalToFieldVarPattern(VarPattern_ GenericSymbol) (GenericSymbol, error)

	// 298:2: field_var_pattern -> named: ...
	NamedToFieldVarPattern(Identifier_ TokenValue, Assign_ TokenValue, VarPattern_ GenericSymbol) (GenericSymbol, error)

	// 299:2: field_var_pattern -> DOT_DOT_DOT: ...
	DotDotDotToFieldVarPattern(DotDotDot_ TokenValue) (GenericSymbol, error)
}

type OptionalValueTypeReducer interface {

	// 303:2: optional_value_type -> nil: ...
	NilToOptionalValueType() (TypeExpression, error)
}

type AssignPatternReducer interface {
	// 310:2: assign_pattern -> ...
	ToAssignPattern(SequenceExpr_ Expression) (GenericSymbol, error)
}

type CasePatternReducer interface {

	// 320:2: case_pattern -> enum_match_pattern: ...
	EnumMatchPatternToCasePattern(Dot_ TokenValue, Identifier_ TokenValue, ImplicitStructExpr_ Expression) (GenericSymbol, error)

	// 321:2: case_pattern -> enum_var_decl_pattern: ...
	EnumVarDeclPatternToCasePattern(Var_ TokenValue, Dot_ TokenValue, Identifier_ TokenValue, TuplePattern_ GenericSymbol) (GenericSymbol, error)
}

type ExpressionReducer interface {
	// 328:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ TokenValue, IfExpr_ Expression) (Expression, error)

	// 329:2: expression -> switch_expr: ...
	SwitchExprToExpression(OptionalLabelDecl_ TokenValue, SwitchExpr_ Expression) (Expression, error)

	// 330:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ TokenValue, LoopExpr_ Expression) (Expression, error)
}

type OptionalLabelDeclReducer interface {
	// 334:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ TokenValue) (TokenValue, error)

	// 335:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (TokenValue, error)
}

type SequenceExprReducer interface {

	// 345:2: sequence_expr -> var_decl_pattern: ...
	VarDeclPatternToSequenceExpr(VarDeclPattern_ GenericSymbol) (Expression, error)

	// 349:2: sequence_expr -> assign_var_pattern: ...
	AssignVarPatternToSequenceExpr(Greater_ TokenValue, SequenceExpr_ Expression) (Expression, error)
}

type IfExprReducer interface {
	// 358:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol) (Expression, error)

	// 359:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol, Else_ TokenValue, StatementBlock_2 GenericSymbol) (Expression, error)

	// 360:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol, Else_ TokenValue, IfExpr_ Expression) (Expression, error)
}

type ConditionReducer interface {
	// 363:2: condition -> sequence_expr: ...
	SequenceExprToCondition(SequenceExpr_ Expression) (GenericSymbol, error)

	// 364:2: condition -> case: ...
	CaseToCondition(Case_ TokenValue, CasePatterns_ GenericSymbol, Assign_ TokenValue, SequenceExpr_ Expression) (GenericSymbol, error)
}

type SwitchExprReducer interface {
	// 388:27: switch_expr -> ...
	ToSwitchExpr(Switch_ TokenValue, SequenceExpr_ Expression, StatementBlock_ GenericSymbol) (Expression, error)
}

type LoopExprReducer interface {
	// 402:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 403:2: loop_expr -> do_while: ...
	DoWhileToLoopExpr(Do_ TokenValue, StatementBlock_ GenericSymbol, For_ TokenValue, SequenceExpr_ Expression) (Expression, error)

	// 404:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ TokenValue, SequenceExpr_ Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 405:2: loop_expr -> iterator: ...
	IteratorToLoopExpr(For_ TokenValue, AssignPattern_ GenericSymbol, In_ TokenValue, SequenceExpr_ Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 406:2: loop_expr -> for: ...
	ForToLoopExpr(For_ TokenValue, ForAssignment_ GenericSymbol, Semicolon_ TokenValue, OptionalSequenceExpr_ Expression, Semicolon_2 TokenValue, OptionalSequenceExpr_2 Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)
}

type OptionalSequenceExprReducer interface {
	// 409:2: optional_sequence_expr -> sequence_expr: ...
	SequenceExprToOptionalSequenceExpr(SequenceExpr_ Expression) (Expression, error)

	// 410:2: optional_sequence_expr -> nil: ...
	NilToOptionalSequenceExpr() (Expression, error)
}

type ForAssignmentReducer interface {
	// 413:2: for_assignment -> sequence_expr: ...
	SequenceExprToForAssignment(SequenceExpr_ Expression) (GenericSymbol, error)

	// 414:2: for_assignment -> assign: ...
	AssignToForAssignment(AssignPattern_ GenericSymbol, Assign_ TokenValue, SequenceExpr_ Expression) (GenericSymbol, error)
}

type CallExprReducer interface {
	// 421:2: call_expr -> ...
	ToCallExpr(AccessibleExpr_ Expression, OptionalGenericBinding_ GenericSymbol, Lparen_ TokenValue, OptionalArguments_ GenericSymbol, Rparen_ TokenValue) (Expression, error)
}

type OptionalGenericBindingReducer interface {
	// 424:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ TokenValue, OptionalGenericArguments_ GenericSymbol, Rbracket_ TokenValue) (GenericSymbol, error)

	// 425:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (GenericSymbol, error)
}

type OptionalGenericArgumentsReducer interface {
	// 428:2: optional_generic_arguments -> generic_arguments: ...
	GenericArgumentsToOptionalGenericArguments(GenericArguments_ GenericSymbol) (GenericSymbol, error)

	// 429:2: optional_generic_arguments -> nil: ...
	NilToOptionalGenericArguments() (GenericSymbol, error)
}

type GenericArgumentsReducer interface {
	// 433:2: generic_arguments -> value_type: ...
	ValueTypeToGenericArguments(ValueType_ TypeExpression) (GenericSymbol, error)

	// 434:2: generic_arguments -> add: ...
	AddToGenericArguments(GenericArguments_ GenericSymbol, Comma_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)
}

type OptionalArgumentsReducer interface {
	// 437:2: optional_arguments -> arguments: ...
	ArgumentsToOptionalArguments(Arguments_ GenericSymbol) (GenericSymbol, error)

	// 438:2: optional_arguments -> nil: ...
	NilToOptionalArguments() (GenericSymbol, error)
}

type ArgumentsReducer interface {
	// 441:2: arguments -> argument: ...
	ArgumentToArguments(Argument_ GenericSymbol) (GenericSymbol, error)

	// 442:2: arguments -> add: ...
	AddToArguments(Arguments_ GenericSymbol, Comma_ TokenValue, Argument_ GenericSymbol) (GenericSymbol, error)
}

type ArgumentReducer interface {
	// 445:2: argument -> positional: ...
	PositionalToArgument(Expression_ Expression) (GenericSymbol, error)

	// 446:2: argument -> named: ...
	NamedToArgument(Identifier_ TokenValue, Assign_ TokenValue, Expression_ Expression) (GenericSymbol, error)

	// 447:2: argument -> colon_expressions: ...
	ColonExpressionsToArgument(ColonExpressions_ GenericSymbol) (GenericSymbol, error)

	// 450:2: argument -> DOT_DOT_DOT: ...
	DotDotDotToArgument(DotDotDot_ TokenValue) (GenericSymbol, error)
}

type ColonExpressionsReducer interface {
	// 453:2: colon_expressions -> pair: ...
	PairToColonExpressions(OptionalExpression_ Expression, Colon_ TokenValue, OptionalExpression_2 Expression) (GenericSymbol, error)

	// 454:2: colon_expressions -> add: ...
	AddToColonExpressions(ColonExpressions_ GenericSymbol, Colon_ TokenValue, OptionalExpression_ Expression) (GenericSymbol, error)
}

type OptionalExpressionReducer interface {
	// 457:2: optional_expression -> expression: ...
	ExpressionToOptionalExpression(Expression_ Expression) (Expression, error)

	// 458:2: optional_expression -> nil: ...
	NilToOptionalExpression() (Expression, error)
}

type ParseErrorExprReducer interface {
	// 477:32: parse_error_expr -> ...
	ToParseErrorExpr(ParseError_ ParseErrorSymbol) (Expression, error)
}

type LiteralExprReducer interface {
	// 480:2: literal_expr -> TRUE: ...
	TrueToLiteralExpr(True_ TokenValue) (Expression, error)

	// 481:2: literal_expr -> FALSE: ...
	FalseToLiteralExpr(False_ TokenValue) (Expression, error)

	// 482:2: literal_expr -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteralExpr(IntegerLiteral_ TokenValue) (Expression, error)

	// 483:2: literal_expr -> FLOAT_LITERAL: ...
	FloatLiteralToLiteralExpr(FloatLiteral_ TokenValue) (Expression, error)

	// 484:2: literal_expr -> RUNE_LITERAL: ...
	RuneLiteralToLiteralExpr(RuneLiteral_ TokenValue) (Expression, error)

	// 485:2: literal_expr -> STRING_LITERAL: ...
	StringLiteralToLiteralExpr(StringLiteral_ TokenValue) (Expression, error)
}

type IdentifierExprReducer interface {
	// 487:32: identifier_expr -> ...
	ToIdentifierExpr(Identifier_ TokenValue) (Expression, error)
}

type BlockExprReducer interface {
	// 489:27: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)
}

type InitializeExprReducer interface {
	// 492:2: initialize_expr -> ...
	ToInitializeExpr(InitializableType_ TypeExpression, Lparen_ TokenValue, OptionalArguments_ GenericSymbol, Rparen_ TokenValue) (Expression, error)
}

type ImplicitStructExprReducer interface {
	// 494:36: implicit_struct_expr -> ...
	ToImplicitStructExpr(Lparen_ TokenValue, OptionalArguments_ GenericSymbol, Rparen_ TokenValue) (Expression, error)
}

type AccessExprReducer interface {
	// 502:27: access_expr -> ...
	ToAccessExpr(AccessibleExpr_ Expression, Dot_ TokenValue, Identifier_ TokenValue) (Expression, error)
}

type IndexExprReducer interface {
	// 504:26: index_expr -> ...
	ToIndexExpr(AccessibleExpr_ Expression, Lbracket_ TokenValue, Argument_ GenericSymbol, Rbracket_ TokenValue) (Expression, error)
}

type PostfixUnaryExprReducer interface {
	// 510:35: postfix_unary_expr -> ...
	ToPostfixUnaryExpr(AccessibleExpr_ Expression, Question_ TokenValue) (Expression, error)
}

type PrefixUnaryExprReducer interface {
	// 516:33: prefix_unary_expr -> ...
	ToPrefixUnaryExpr(PrefixUnaryOp_ TokenValue, PrefixableExpr_ Expression) (Expression, error)
}

type BinaryMulExprReducer interface {
	// 533:31: binary_mul_expr -> ...
	ToBinaryMulExpr(MulExpr_ Expression, MulOp_ TokenValue, PrefixableExpr_ Expression) (Expression, error)
}

type BinaryAddExprReducer interface {
	// 547:31: binary_add_expr -> ...
	ToBinaryAddExpr(AddExpr_ Expression, AddOp_ TokenValue, MulExpr_ Expression) (Expression, error)
}

type BinaryCmpExprReducer interface {
	// 559:31: binary_cmp_expr -> ...
	ToBinaryCmpExpr(CmpExpr_ Expression, CmpOp_ TokenValue, AddExpr_ Expression) (Expression, error)
}

type BinaryAndExprReducer interface {
	// 573:31: binary_and_expr -> ...
	ToBinaryAndExpr(AndExpr_ Expression, And_ TokenValue, CmpExpr_ Expression) (Expression, error)
}

type BinaryOrExprReducer interface {
	// 579:30: binary_or_expr -> ...
	ToBinaryOrExpr(OrExpr_ Expression, Or_ TokenValue, AndExpr_ Expression) (Expression, error)
}

type InitializableTypeReducer interface {

	// 589:2: initializable_type -> slice: ...
	SliceToInitializableType(Lbracket_ TokenValue, ValueType_ TypeExpression, Rbracket_ TokenValue) (TypeExpression, error)

	// 590:2: initializable_type -> array: ...
	ArrayToInitializableType(Lbracket_ TokenValue, ValueType_ TypeExpression, Comma_ TokenValue, IntegerLiteral_ TokenValue, Rbracket_ TokenValue) (TypeExpression, error)

	// 591:2: initializable_type -> map: ...
	MapToInitializableType(Lbracket_ TokenValue, ValueType_ TypeExpression, Colon_ TokenValue, ValueType_2 TypeExpression, Rbracket_ TokenValue) (TypeExpression, error)
}

type AtomTypeReducer interface {

	// 595:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ TokenValue, OptionalGenericBinding_ GenericSymbol) (TypeExpression, error)

	// 596:2: atom_type -> extern_named: ...
	ExternNamedToAtomType(Identifier_ TokenValue, Dot_ TokenValue, Identifier_2 TokenValue, OptionalGenericBinding_ GenericSymbol) (TypeExpression, error)

	// 597:2: atom_type -> inferred: ...
	InferredToAtomType(Dot_ TokenValue, OptionalGenericBinding_ GenericSymbol) (TypeExpression, error)
}

type ParseErrorTypeReducer interface {
	// 605:36: parse_error_type -> ...
	ToParseErrorType(ParseError_ ParseErrorSymbol) (TypeExpression, error)
}

type PrefixedTypeReducer interface {
	// 614:33: prefixed_type -> ...
	ToPrefixedType(PrefixTypeOp_ TokenValue, ReturnableType_ TypeExpression) (TypeExpression, error)
}

type TraitOpTypeReducer interface {
	// 629:33: trait_op_type -> ...
	ToTraitOpType(ValueType_ TypeExpression, TraitOp_ TokenValue, ReturnableType_ TypeExpression) (TypeExpression, error)
}

type TypeDefReducer interface {
	// 637:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, ValueType_ TypeExpression) (SourceDefinition, error)

	// 638:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, ValueType_ TypeExpression, Implements_ TokenValue, ValueType_2 TypeExpression) (SourceDefinition, error)

	// 639:2: type_def -> alias: ...
	AliasToTypeDef(Type_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, ValueType_ TypeExpression) (SourceDefinition, error)
}

type GenericParameterDefReducer interface {
	// 647:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ TokenValue) (GenericSymbol, error)

	// 648:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)
}

type GenericParameterDefsReducer interface {
	// 651:2: generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToGenericParameterDefs(GenericParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 652:2: generic_parameter_defs -> add: ...
	AddToGenericParameterDefs(GenericParameterDefs_ GenericSymbol, Comma_ TokenValue, GenericParameterDef_ GenericSymbol) (GenericSymbol, error)
}

type OptionalGenericParameterDefsReducer interface {
	// 655:2: optional_generic_parameter_defs -> generic_parameter_defs: ...
	GenericParameterDefsToOptionalGenericParameterDefs(GenericParameterDefs_ GenericSymbol) (GenericSymbol, error)

	// 656:2: optional_generic_parameter_defs -> nil: ...
	NilToOptionalGenericParameterDefs() (GenericSymbol, error)
}

type OptionalGenericParametersReducer interface {
	// 659:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ TokenValue, OptionalGenericParameterDefs_ GenericSymbol, Rbracket_ TokenValue) (GenericSymbol, error)

	// 660:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (GenericSymbol, error)
}

type FieldDefReducer interface {
	// 667:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 668:2: field_def -> implicit: ...
	ImplicitToFieldDef(ValueType_ TypeExpression) (GenericSymbol, error)

	// 669:2: field_def -> unsafe_statement: ...
	UnsafeStatementToFieldDef(UnsafeStatement_ Statement) (GenericSymbol, error)
}

type ImplicitFieldDefsReducer interface {
	// 672:2: implicit_field_defs -> field_def: ...
	FieldDefToImplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 673:2: implicit_field_defs -> add: ...
	AddToImplicitFieldDefs(ImplicitFieldDefs_ GenericSymbol, Comma_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type OptionalImplicitFieldDefsReducer interface {
	// 676:2: optional_implicit_field_defs -> implicit_field_defs: ...
	ImplicitFieldDefsToOptionalImplicitFieldDefs(ImplicitFieldDefs_ GenericSymbol) (GenericSymbol, error)

	// 677:2: optional_implicit_field_defs -> nil: ...
	NilToOptionalImplicitFieldDefs() (GenericSymbol, error)
}

type ImplicitStructDefReducer interface {
	// 680:2: implicit_struct_def -> ...
	ToImplicitStructDef(Lparen_ TokenValue, OptionalImplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ExplicitFieldDefsReducer interface {
	// 683:2: explicit_field_defs -> field_def: ...
	FieldDefToExplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 684:2: explicit_field_defs -> implicit: ...
	ImplicitToExplicitFieldDefs(ExplicitFieldDefs_ GenericSymbol, Newlines_ TokenCount, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 685:2: explicit_field_defs -> explicit: ...
	ExplicitToExplicitFieldDefs(ExplicitFieldDefs_ GenericSymbol, Comma_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type OptionalExplicitFieldDefsReducer interface {
	// 688:2: optional_explicit_field_defs -> explicit_field_defs: ...
	ExplicitFieldDefsToOptionalExplicitFieldDefs(ExplicitFieldDefs_ GenericSymbol) (GenericSymbol, error)

	// 689:2: optional_explicit_field_defs -> nil: ...
	NilToOptionalExplicitFieldDefs() (GenericSymbol, error)
}

type ExplicitStructDefReducer interface {
	// 692:2: explicit_struct_def -> ...
	ToExplicitStructDef(Struct_ TokenValue, Lparen_ TokenValue, OptionalExplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type EnumValueDefReducer interface {
	// 700:2: enum_value_def -> field_def: ...
	FieldDefToEnumValueDef(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 701:2: enum_value_def -> default: ...
	DefaultToEnumValueDef(FieldDef_ GenericSymbol, Assign_ TokenValue, Default_ TokenValue) (GenericSymbol, error)
}

type ImplicitEnumValueDefsReducer interface {
	// 713:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Or_ TokenValue, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 714:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Or_ TokenValue, EnumValueDef_ GenericSymbol) (GenericSymbol, error)
}

type ImplicitEnumDefReducer interface {
	// 716:37: implicit_enum_def -> ...
	ToImplicitEnumDef(Lparen_ TokenValue, ImplicitEnumValueDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ExplicitEnumValueDefsReducer interface {
	// 719:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Or_ TokenValue, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 720:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Newlines_ TokenCount, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 721:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Or_ TokenValue, EnumValueDef_ GenericSymbol) (GenericSymbol, error)

	// 722:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Newlines_ TokenCount, EnumValueDef_ GenericSymbol) (GenericSymbol, error)
}

type ExplicitEnumDefReducer interface {
	// 724:37: explicit_enum_def -> ...
	ToExplicitEnumDef(Enum_ TokenValue, Lparen_ TokenValue, ExplicitEnumValueDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type TraitPropertyReducer interface {
	// 731:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 732:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ GenericSymbol) (GenericSymbol, error)
}

type TraitPropertiesReducer interface {
	// 735:2: trait_properties -> trait_property: ...
	TraitPropertyToTraitProperties(TraitProperty_ GenericSymbol) (GenericSymbol, error)

	// 736:2: trait_properties -> implicit: ...
	ImplicitToTraitProperties(TraitProperties_ GenericSymbol, Newlines_ TokenCount, TraitProperty_ GenericSymbol) (GenericSymbol, error)

	// 737:2: trait_properties -> explicit: ...
	ExplicitToTraitProperties(TraitProperties_ GenericSymbol, Comma_ TokenValue, TraitProperty_ GenericSymbol) (GenericSymbol, error)
}

type OptionalTraitPropertiesReducer interface {
	// 740:2: optional_trait_properties -> trait_properties: ...
	TraitPropertiesToOptionalTraitProperties(TraitProperties_ GenericSymbol) (GenericSymbol, error)

	// 741:2: optional_trait_properties -> nil: ...
	NilToOptionalTraitProperties() (GenericSymbol, error)
}

type TraitDefReducer interface {
	// 743:29: trait_def -> ...
	ToTraitDef(Trait_ TokenValue, Lparen_ TokenValue, OptionalTraitProperties_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ReturnTypeReducer interface {
	// 751:2: return_type -> returnable_type: ...
	ReturnableTypeToReturnType(ReturnableType_ TypeExpression) (TypeExpression, error)

	// 752:2: return_type -> nil: ...
	NilToReturnType() (TypeExpression, error)
}

type ParameterDeclReducer interface {
	// 755:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 756:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ TokenValue, DotDotDot_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 757:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ TypeExpression) (GenericSymbol, error)

	// 758:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(DotDotDot_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)
}

type ParameterDeclsReducer interface {
	// 761:2: parameter_decls -> parameter_decl: ...
	ParameterDeclToParameterDecls(ParameterDecl_ GenericSymbol) (GenericSymbol, error)

	// 762:2: parameter_decls -> add: ...
	AddToParameterDecls(ParameterDecls_ GenericSymbol, Comma_ TokenValue, ParameterDecl_ GenericSymbol) (GenericSymbol, error)
}

type OptionalParameterDeclsReducer interface {
	// 765:2: optional_parameter_decls -> parameter_decls: ...
	ParameterDeclsToOptionalParameterDecls(ParameterDecls_ GenericSymbol) (GenericSymbol, error)

	// 766:2: optional_parameter_decls -> nil: ...
	NilToOptionalParameterDecls() (GenericSymbol, error)
}

type FuncTypeReducer interface {
	// 769:2: func_type -> ...
	ToFuncType(Func_ TokenValue, Lparen_ TokenValue, OptionalParameterDecls_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression) (TypeExpression, error)
}

type MethodSignatureReducer interface {
	// 780:20: method_signature -> ...
	ToMethodSignature(Func_ TokenValue, Identifier_ TokenValue, Lparen_ TokenValue, OptionalParameterDecls_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression) (GenericSymbol, error)
}

type ParameterDefReducer interface {
	// 786:2: parameter_def -> inferred_ref_arg: ...
	InferredRefArgToParameterDef(Identifier_ TokenValue) (GenericSymbol, error)

	// 787:2: parameter_def -> inferred_ref_vararg: ...
	InferredRefVarargToParameterDef(Identifier_ TokenValue, DotDotDot_ TokenValue) (GenericSymbol, error)

	// 788:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 789:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ TokenValue, DotDotDot_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)
}

type ParameterDefsReducer interface {
	// 792:2: parameter_defs -> parameter_def: ...
	ParameterDefToParameterDefs(ParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 793:2: parameter_defs -> add: ...
	AddToParameterDefs(ParameterDefs_ GenericSymbol, Comma_ TokenValue, ParameterDef_ GenericSymbol) (GenericSymbol, error)
}

type OptionalParameterDefsReducer interface {
	// 796:2: optional_parameter_defs -> parameter_defs: ...
	ParameterDefsToOptionalParameterDefs(ParameterDefs_ GenericSymbol) (GenericSymbol, error)

	// 797:2: optional_parameter_defs -> nil: ...
	NilToOptionalParameterDefs() (GenericSymbol, error)
}

type NamedFuncDefReducer interface {
	// 800:2: named_func_def -> func_def: ...
	FuncDefToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, Lparen_ TokenValue, OptionalParameterDefs_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 801:2: named_func_def -> method_def: ...
	MethodDefToNamedFuncDef(Func_ TokenValue, Lparen_ TokenValue, ParameterDef_ GenericSymbol, Rparen_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, Lparen_2 TokenValue, OptionalParameterDefs_ GenericSymbol, Rparen_2 TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 802:2: named_func_def -> alias: ...
	AliasToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, Expression_ Expression) (SourceDefinition, error)
}

type AnonymousFuncExprReducer interface {
	// 806:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ TokenValue, Lparen_ TokenValue, OptionalParameterDefs_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (Expression, error)
}

type PackageDefReducer interface {
	// 818:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ TokenValue) (SourceDefinition, error)

	// 819:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ TokenValue, StatementBlock_ GenericSymbol) (SourceDefinition, error)
}

type Reducer interface {
	SourceReducer
	OptionalDefinitionsReducer
	OptionalNewlinesReducer
	DefinitionsReducer
	DefinitionReducer
	StatementBlockReducer
	StatementsReducer
	StatementReducer
	StatementBodyReducer
	OptionalSimpleStatementBodyReducer
	ExpressionOrImproperStructStatementReducer
	CallbackOpStatementReducer
	UnsafeStatementReducer
	JumpStatementReducer
	ExpressionsReducer
	FallthroughStatementReducer
	AssignStatementReducer
	UnaryOpAssignStatementReducer
	BinaryOpAssignStatementReducer
	ImportStatementReducer
	ImportClauseReducer
	ImportClauseTerminalReducer
	ImportClausesReducer
	CasePatternsReducer
	VarDeclPatternReducer
	VarPatternReducer
	TuplePatternReducer
	FieldVarPatternsReducer
	FieldVarPatternReducer
	OptionalValueTypeReducer
	AssignPatternReducer
	CasePatternReducer
	ExpressionReducer
	OptionalLabelDeclReducer
	SequenceExprReducer
	IfExprReducer
	ConditionReducer
	SwitchExprReducer
	LoopExprReducer
	OptionalSequenceExprReducer
	ForAssignmentReducer
	CallExprReducer
	OptionalGenericBindingReducer
	OptionalGenericArgumentsReducer
	GenericArgumentsReducer
	OptionalArgumentsReducer
	ArgumentsReducer
	ArgumentReducer
	ColonExpressionsReducer
	OptionalExpressionReducer
	ParseErrorExprReducer
	LiteralExprReducer
	IdentifierExprReducer
	BlockExprReducer
	InitializeExprReducer
	ImplicitStructExprReducer
	AccessExprReducer
	IndexExprReducer
	PostfixUnaryExprReducer
	PrefixUnaryExprReducer
	BinaryMulExprReducer
	BinaryAddExprReducer
	BinaryCmpExprReducer
	BinaryAndExprReducer
	BinaryOrExprReducer
	InitializableTypeReducer
	AtomTypeReducer
	ParseErrorTypeReducer
	PrefixedTypeReducer
	TraitOpTypeReducer
	TypeDefReducer
	GenericParameterDefReducer
	GenericParameterDefsReducer
	OptionalGenericParameterDefsReducer
	OptionalGenericParametersReducer
	FieldDefReducer
	ImplicitFieldDefsReducer
	OptionalImplicitFieldDefsReducer
	ImplicitStructDefReducer
	ExplicitFieldDefsReducer
	OptionalExplicitFieldDefsReducer
	ExplicitStructDefReducer
	EnumValueDefReducer
	ImplicitEnumValueDefsReducer
	ImplicitEnumDefReducer
	ExplicitEnumValueDefsReducer
	ExplicitEnumDefReducer
	TraitPropertyReducer
	TraitPropertiesReducer
	OptionalTraitPropertiesReducer
	TraitDefReducer
	ReturnTypeReducer
	ParameterDeclReducer
	ParameterDeclsReducer
	OptionalParameterDeclsReducer
	FuncTypeReducer
	MethodSignatureReducer
	ParameterDefReducer
	ParameterDefsReducer
	OptionalParameterDefsReducer
	NamedFuncDefReducer
	AnonymousFuncExprReducer
	PackageDefReducer
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

func ParseSource(lexer Lexer, reducer Reducer) ([]SourceDefinition, error) {

	return ParseSourceWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseSourceWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	[]SourceDefinition,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State1)
	if err != nil {
		var errRetVal []SourceDefinition
		return errRetVal, err
	}
	return item.SourceDefinitions, nil
}

func ParsePackageDef(lexer Lexer, reducer Reducer) (SourceDefinition, error) {

	return ParsePackageDefWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParsePackageDefWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	SourceDefinition,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State2)
	if err != nil {
		var errRetVal SourceDefinition
		return errRetVal, err
	}
	return item.SourceDefinition, nil
}

func ParseTypeDef(lexer Lexer, reducer Reducer) (SourceDefinition, error) {

	return ParseTypeDefWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseTypeDefWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	SourceDefinition,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State3)
	if err != nil {
		var errRetVal SourceDefinition
		return errRetVal, err
	}
	return item.SourceDefinition, nil
}

func ParseNamedFuncDef(lexer Lexer, reducer Reducer) (SourceDefinition, error) {

	return ParseNamedFuncDefWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseNamedFuncDefWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	SourceDefinition,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State4)
	if err != nil {
		var errRetVal SourceDefinition
		return errRetVal, err
	}
	return item.SourceDefinition, nil
}

func ParseExpression(lexer Lexer, reducer Reducer) (Expression, error) {

	return ParseExpressionWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseExpressionWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	Expression,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State5)
	if err != nil {
		var errRetVal Expression
		return errRetVal, err
	}
	return item.Expression, nil
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
	case NewlinesToken:
		return "NEWLINES"
	case CommentGroupsToken:
		return "COMMENT_GROUPS"
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
	case FallthroughToken:
		return "FALLTHROUGH"
	case PackageToken:
		return "PACKAGE"
	case ImportToken:
		return "IMPORT"
	case AsToken:
		return "AS"
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
	case LetToken:
		return "LET"
	case NotToken:
		return "NOT"
	case AndToken:
		return "AND"
	case OrToken:
		return "OR"
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
	case DotDotDotToken:
		return "DOT_DOT_DOT"
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
	case ParseErrorToken:
		return "PARSE_ERROR"
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
	case StatementBlockType:
		return "statement_block"
	case StatementsType:
		return "statements"
	case StatementType:
		return "statement"
	case SimpleStatementBodyType:
		return "simple_statement_body"
	case StatementBodyType:
		return "statement_body"
	case OptionalSimpleStatementBodyType:
		return "optional_simple_statement_body"
	case ExpressionOrImproperStructStatementType:
		return "expression_or_improper_struct_statement"
	case CallbackOpType:
		return "callback_op"
	case CallbackOpStatementType:
		return "callback_op_statement"
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
	case ExpressionsType:
		return "expressions"
	case FallthroughStatementType:
		return "fallthrough_statement"
	case AssignStatementType:
		return "assign_statement"
	case UnaryOpAssignStatementType:
		return "unary_op_assign_statement"
	case BinaryOpAssignStatementType:
		return "binary_op_assign_statement"
	case ImportStatementType:
		return "import_statement"
	case ImportClauseType:
		return "import_clause"
	case ImportClauseTerminalType:
		return "import_clause_terminal"
	case ImportClausesType:
		return "import_clauses"
	case CasePatternsType:
		return "case_patterns"
	case VarDeclPatternType:
		return "var_decl_pattern"
	case VarOrLetType:
		return "var_or_let"
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
	case AssignPatternType:
		return "assign_pattern"
	case CasePatternType:
		return "case_pattern"
	case ExpressionType:
		return "expression"
	case OptionalLabelDeclType:
		return "optional_label_decl"
	case SequenceExprType:
		return "sequence_expr"
	case IfExprType:
		return "if_expr"
	case ConditionType:
		return "condition"
	case SwitchExprType:
		return "switch_expr"
	case LoopExprType:
		return "loop_expr"
	case OptionalSequenceExprType:
		return "optional_sequence_expr"
	case ForAssignmentType:
		return "for_assignment"
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
	case ParseErrorExprType:
		return "parse_error_expr"
	case LiteralExprType:
		return "literal_expr"
	case IdentifierExprType:
		return "identifier_expr"
	case BlockExprType:
		return "block_expr"
	case InitializeExprType:
		return "initialize_expr"
	case ImplicitStructExprType:
		return "implicit_struct_expr"
	case AccessibleExprType:
		return "accessible_expr"
	case AccessExprType:
		return "access_expr"
	case IndexExprType:
		return "index_expr"
	case PostfixableExprType:
		return "postfixable_expr"
	case PostfixUnaryExprType:
		return "postfix_unary_expr"
	case PrefixableExprType:
		return "prefixable_expr"
	case PrefixUnaryExprType:
		return "prefix_unary_expr"
	case PrefixUnaryOpType:
		return "prefix_unary_op"
	case MulExprType:
		return "mul_expr"
	case BinaryMulExprType:
		return "binary_mul_expr"
	case MulOpType:
		return "mul_op"
	case AddExprType:
		return "add_expr"
	case BinaryAddExprType:
		return "binary_add_expr"
	case AddOpType:
		return "add_op"
	case CmpExprType:
		return "cmp_expr"
	case BinaryCmpExprType:
		return "binary_cmp_expr"
	case CmpOpType:
		return "cmp_op"
	case AndExprType:
		return "and_expr"
	case BinaryAndExprType:
		return "binary_and_expr"
	case OrExprType:
		return "or_expr"
	case BinaryOrExprType:
		return "binary_or_expr"
	case InitializableTypeType:
		return "initializable_type"
	case AtomTypeType:
		return "atom_type"
	case ParseErrorTypeType:
		return "parse_error_type"
	case ReturnableTypeType:
		return "returnable_type"
	case PrefixedTypeType:
		return "prefixed_type"
	case PrefixTypeOpType:
		return "prefix_type_op"
	case ValueTypeType:
		return "value_type"
	case TraitOpTypeType:
		return "trait_op_type"
	case TraitOpType:
		return "trait_op"
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
	case NamedFuncDefType:
		return "named_func_def"
	case AnonymousFuncExprType:
		return "anonymous_func_expr"
	case PackageDefType:
		return "package_def"
	default:
		return fmt.Sprintf("?unknown symbol %d?", int(i))
	}
}

const (
	_EndMarker      = SymbolId(0)
	_WildcardMarker = SymbolId(-1)

	SourceType                              = SymbolId(343)
	OptionalDefinitionsType                 = SymbolId(344)
	OptionalNewlinesType                    = SymbolId(345)
	DefinitionsType                         = SymbolId(346)
	DefinitionType                          = SymbolId(347)
	StatementBlockType                      = SymbolId(348)
	StatementsType                          = SymbolId(349)
	StatementType                           = SymbolId(350)
	SimpleStatementBodyType                 = SymbolId(351)
	StatementBodyType                       = SymbolId(352)
	OptionalSimpleStatementBodyType         = SymbolId(353)
	ExpressionOrImproperStructStatementType = SymbolId(354)
	CallbackOpType                          = SymbolId(355)
	CallbackOpStatementType                 = SymbolId(356)
	UnaryOpAssignType                       = SymbolId(357)
	BinaryOpAssignType                      = SymbolId(358)
	UnsafeStatementType                     = SymbolId(359)
	JumpStatementType                       = SymbolId(360)
	JumpTypeType                            = SymbolId(361)
	ExpressionsType                         = SymbolId(362)
	FallthroughStatementType                = SymbolId(363)
	AssignStatementType                     = SymbolId(364)
	UnaryOpAssignStatementType              = SymbolId(365)
	BinaryOpAssignStatementType             = SymbolId(366)
	ImportStatementType                     = SymbolId(367)
	ImportClauseType                        = SymbolId(368)
	ImportClauseTerminalType                = SymbolId(369)
	ImportClausesType                       = SymbolId(370)
	CasePatternsType                        = SymbolId(371)
	VarDeclPatternType                      = SymbolId(372)
	VarOrLetType                            = SymbolId(373)
	VarPatternType                          = SymbolId(374)
	TuplePatternType                        = SymbolId(375)
	FieldVarPatternsType                    = SymbolId(376)
	FieldVarPatternType                     = SymbolId(377)
	OptionalValueTypeType                   = SymbolId(378)
	AssignPatternType                       = SymbolId(379)
	CasePatternType                         = SymbolId(380)
	ExpressionType                          = SymbolId(381)
	OptionalLabelDeclType                   = SymbolId(382)
	SequenceExprType                        = SymbolId(383)
	IfExprType                              = SymbolId(384)
	ConditionType                           = SymbolId(385)
	SwitchExprType                          = SymbolId(386)
	LoopExprType                            = SymbolId(387)
	OptionalSequenceExprType                = SymbolId(388)
	ForAssignmentType                       = SymbolId(389)
	CallExprType                            = SymbolId(390)
	OptionalGenericBindingType              = SymbolId(391)
	OptionalGenericArgumentsType            = SymbolId(392)
	GenericArgumentsType                    = SymbolId(393)
	OptionalArgumentsType                   = SymbolId(394)
	ArgumentsType                           = SymbolId(395)
	ArgumentType                            = SymbolId(396)
	ColonExpressionsType                    = SymbolId(397)
	OptionalExpressionType                  = SymbolId(398)
	AtomExprType                            = SymbolId(399)
	ParseErrorExprType                      = SymbolId(400)
	LiteralExprType                         = SymbolId(401)
	IdentifierExprType                      = SymbolId(402)
	BlockExprType                           = SymbolId(403)
	InitializeExprType                      = SymbolId(404)
	ImplicitStructExprType                  = SymbolId(405)
	AccessibleExprType                      = SymbolId(406)
	AccessExprType                          = SymbolId(407)
	IndexExprType                           = SymbolId(408)
	PostfixableExprType                     = SymbolId(409)
	PostfixUnaryExprType                    = SymbolId(410)
	PrefixableExprType                      = SymbolId(411)
	PrefixUnaryExprType                     = SymbolId(412)
	PrefixUnaryOpType                       = SymbolId(413)
	MulExprType                             = SymbolId(414)
	BinaryMulExprType                       = SymbolId(415)
	MulOpType                               = SymbolId(416)
	AddExprType                             = SymbolId(417)
	BinaryAddExprType                       = SymbolId(418)
	AddOpType                               = SymbolId(419)
	CmpExprType                             = SymbolId(420)
	BinaryCmpExprType                       = SymbolId(421)
	CmpOpType                               = SymbolId(422)
	AndExprType                             = SymbolId(423)
	BinaryAndExprType                       = SymbolId(424)
	OrExprType                              = SymbolId(425)
	BinaryOrExprType                        = SymbolId(426)
	InitializableTypeType                   = SymbolId(427)
	AtomTypeType                            = SymbolId(428)
	ParseErrorTypeType                      = SymbolId(429)
	ReturnableTypeType                      = SymbolId(430)
	PrefixedTypeType                        = SymbolId(431)
	PrefixTypeOpType                        = SymbolId(432)
	ValueTypeType                           = SymbolId(433)
	TraitOpTypeType                         = SymbolId(434)
	TraitOpType                             = SymbolId(435)
	TypeDefType                             = SymbolId(436)
	GenericParameterDefType                 = SymbolId(437)
	GenericParameterDefsType                = SymbolId(438)
	OptionalGenericParameterDefsType        = SymbolId(439)
	OptionalGenericParametersType           = SymbolId(440)
	FieldDefType                            = SymbolId(441)
	ImplicitFieldDefsType                   = SymbolId(442)
	OptionalImplicitFieldDefsType           = SymbolId(443)
	ImplicitStructDefType                   = SymbolId(444)
	ExplicitFieldDefsType                   = SymbolId(445)
	OptionalExplicitFieldDefsType           = SymbolId(446)
	ExplicitStructDefType                   = SymbolId(447)
	EnumValueDefType                        = SymbolId(448)
	ImplicitEnumValueDefsType               = SymbolId(449)
	ImplicitEnumDefType                     = SymbolId(450)
	ExplicitEnumValueDefsType               = SymbolId(451)
	ExplicitEnumDefType                     = SymbolId(452)
	TraitPropertyType                       = SymbolId(453)
	TraitPropertiesType                     = SymbolId(454)
	OptionalTraitPropertiesType             = SymbolId(455)
	TraitDefType                            = SymbolId(456)
	ReturnTypeType                          = SymbolId(457)
	ParameterDeclType                       = SymbolId(458)
	ParameterDeclsType                      = SymbolId(459)
	OptionalParameterDeclsType              = SymbolId(460)
	FuncTypeType                            = SymbolId(461)
	MethodSignatureType                     = SymbolId(462)
	ParameterDefType                        = SymbolId(463)
	ParameterDefsType                       = SymbolId(464)
	OptionalParameterDefsType               = SymbolId(465)
	NamedFuncDefType                        = SymbolId(466)
	AnonymousFuncExprType                   = SymbolId(467)
	PackageDefType                          = SymbolId(468)
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
	_ReduceToSource                                                 = _ReduceType(1)
	_ReduceNewlinesToOptionalDefinitions                            = _ReduceType(2)
	_ReduceDefinitionsToOptionalDefinitions                         = _ReduceType(3)
	_ReduceNilToOptionalDefinitions                                 = _ReduceType(4)
	_ReduceNewlinesToOptionalNewlines                               = _ReduceType(5)
	_ReduceNilToOptionalNewlines                                    = _ReduceType(6)
	_ReduceDefinitionToDefinitions                                  = _ReduceType(7)
	_ReduceAddToDefinitions                                         = _ReduceType(8)
	_ReducePackageDefToDefinition                                   = _ReduceType(9)
	_ReduceTypeDefToDefinition                                      = _ReduceType(10)
	_ReduceNamedFuncDefToDefinition                                 = _ReduceType(11)
	_ReduceGlobalVarDefToDefinition                                 = _ReduceType(12)
	_ReduceGlobalVarAssignmentToDefinition                          = _ReduceType(13)
	_ReduceStatementBlockToDefinition                               = _ReduceType(14)
	_ReduceCommentGroupsToDefinition                                = _ReduceType(15)
	_ReduceToStatementBlock                                         = _ReduceType(16)
	_ReduceEmptyListToStatements                                    = _ReduceType(17)
	_ReduceAddToStatements                                          = _ReduceType(18)
	_ReduceImplicitToStatement                                      = _ReduceType(19)
	_ReduceExplicitToStatement                                      = _ReduceType(20)
	_ReduceUnsafeStatementToSimpleStatementBody                     = _ReduceType(21)
	_ReduceExpressionOrImproperStructStatementToSimpleStatementBody = _ReduceType(22)
	_ReduceCallbackOpStatementToSimpleStatementBody                 = _ReduceType(23)
	_ReduceJumpStatementToSimpleStatementBody                       = _ReduceType(24)
	_ReduceAssignStatementToSimpleStatementBody                     = _ReduceType(25)
	_ReduceUnaryOpAssignStatementToSimpleStatementBody              = _ReduceType(26)
	_ReduceBinaryOpAssignStatementToSimpleStatementBody             = _ReduceType(27)
	_ReduceFallthroughStatementToSimpleStatementBody                = _ReduceType(28)
	_ReduceSimpleStatementBodyToStatementBody                       = _ReduceType(29)
	_ReduceImportStatementToStatementBody                           = _ReduceType(30)
	_ReduceCaseBranchStatementToStatementBody                       = _ReduceType(31)
	_ReduceDefaultBranchStatementToStatementBody                    = _ReduceType(32)
	_ReduceSimpleStatementBodyToOptionalSimpleStatementBody         = _ReduceType(33)
	_ReduceNilToOptionalSimpleStatementBody                         = _ReduceType(34)
	_ReduceToExpressionOrImproperStructStatement                    = _ReduceType(35)
	_ReduceAsyncToCallbackOp                                        = _ReduceType(36)
	_ReduceDeferToCallbackOp                                        = _ReduceType(37)
	_ReduceToCallbackOpStatement                                    = _ReduceType(38)
	_ReduceAddOneAssignToUnaryOpAssign                              = _ReduceType(39)
	_ReduceSubOneAssignToUnaryOpAssign                              = _ReduceType(40)
	_ReduceAddAssignToBinaryOpAssign                                = _ReduceType(41)
	_ReduceSubAssignToBinaryOpAssign                                = _ReduceType(42)
	_ReduceMulAssignToBinaryOpAssign                                = _ReduceType(43)
	_ReduceDivAssignToBinaryOpAssign                                = _ReduceType(44)
	_ReduceModAssignToBinaryOpAssign                                = _ReduceType(45)
	_ReduceBitNegAssignToBinaryOpAssign                             = _ReduceType(46)
	_ReduceBitAndAssignToBinaryOpAssign                             = _ReduceType(47)
	_ReduceBitOrAssignToBinaryOpAssign                              = _ReduceType(48)
	_ReduceBitXorAssignToBinaryOpAssign                             = _ReduceType(49)
	_ReduceBitLshiftAssignToBinaryOpAssign                          = _ReduceType(50)
	_ReduceBitRshiftAssignToBinaryOpAssign                          = _ReduceType(51)
	_ReduceToUnsafeStatement                                        = _ReduceType(52)
	_ReduceUnlabeledNoValueToJumpStatement                          = _ReduceType(53)
	_ReduceUnlabeledValuedToJumpStatement                           = _ReduceType(54)
	_ReduceLabeledNoValueToJumpStatement                            = _ReduceType(55)
	_ReduceLabeledValuedToJumpStatement                             = _ReduceType(56)
	_ReduceReturnToJumpType                                         = _ReduceType(57)
	_ReduceBreakToJumpType                                          = _ReduceType(58)
	_ReduceContinueToJumpType                                       = _ReduceType(59)
	_ReduceExpressionToExpressions                                  = _ReduceType(60)
	_ReduceAddToExpressions                                         = _ReduceType(61)
	_ReduceToFallthroughStatement                                   = _ReduceType(62)
	_ReduceToAssignStatement                                        = _ReduceType(63)
	_ReduceToUnaryOpAssignStatement                                 = _ReduceType(64)
	_ReduceToBinaryOpAssignStatement                                = _ReduceType(65)
	_ReduceSingleToImportStatement                                  = _ReduceType(66)
	_ReduceMultipleToImportStatement                                = _ReduceType(67)
	_ReduceStringLiteralToImportClause                              = _ReduceType(68)
	_ReduceAliasToImportClause                                      = _ReduceType(69)
	_ReduceImplicitToImportClauseTerminal                           = _ReduceType(70)
	_ReduceExplicitToImportClauseTerminal                           = _ReduceType(71)
	_ReduceFirstToImportClauses                                     = _ReduceType(72)
	_ReduceAddToImportClauses                                       = _ReduceType(73)
	_ReduceCasePatternToCasePatterns                                = _ReduceType(74)
	_ReduceMultipleToCasePatterns                                   = _ReduceType(75)
	_ReduceToVarDeclPattern                                         = _ReduceType(76)
	_ReduceVarToVarOrLet                                            = _ReduceType(77)
	_ReduceLetToVarOrLet                                            = _ReduceType(78)
	_ReduceIdentifierToVarPattern                                   = _ReduceType(79)
	_ReduceTuplePatternToVarPattern                                 = _ReduceType(80)
	_ReduceToTuplePattern                                           = _ReduceType(81)
	_ReduceFieldVarPatternToFieldVarPatterns                        = _ReduceType(82)
	_ReduceAddToFieldVarPatterns                                    = _ReduceType(83)
	_ReducePositionalToFieldVarPattern                              = _ReduceType(84)
	_ReduceNamedToFieldVarPattern                                   = _ReduceType(85)
	_ReduceDotDotDotToFieldVarPattern                               = _ReduceType(86)
	_ReduceValueTypeToOptionalValueType                             = _ReduceType(87)
	_ReduceNilToOptionalValueType                                   = _ReduceType(88)
	_ReduceToAssignPattern                                          = _ReduceType(89)
	_ReduceAssignPatternToCasePattern                               = _ReduceType(90)
	_ReduceEnumMatchPatternToCasePattern                            = _ReduceType(91)
	_ReduceEnumVarDeclPatternToCasePattern                          = _ReduceType(92)
	_ReduceIfExprToExpression                                       = _ReduceType(93)
	_ReduceSwitchExprToExpression                                   = _ReduceType(94)
	_ReduceLoopExprToExpression                                     = _ReduceType(95)
	_ReduceSequenceExprToExpression                                 = _ReduceType(96)
	_ReduceLabelDeclToOptionalLabelDecl                             = _ReduceType(97)
	_ReduceUnlabelledToOptionalLabelDecl                            = _ReduceType(98)
	_ReduceOrExprToSequenceExpr                                     = _ReduceType(99)
	_ReduceVarDeclPatternToSequenceExpr                             = _ReduceType(100)
	_ReduceAssignVarPatternToSequenceExpr                           = _ReduceType(101)
	_ReduceNoElseToIfExpr                                           = _ReduceType(102)
	_ReduceIfElseToIfExpr                                           = _ReduceType(103)
	_ReduceMultiIfElseToIfExpr                                      = _ReduceType(104)
	_ReduceSequenceExprToCondition                                  = _ReduceType(105)
	_ReduceCaseToCondition                                          = _ReduceType(106)
	_ReduceToSwitchExpr                                             = _ReduceType(107)
	_ReduceInfiniteToLoopExpr                                       = _ReduceType(108)
	_ReduceDoWhileToLoopExpr                                        = _ReduceType(109)
	_ReduceWhileToLoopExpr                                          = _ReduceType(110)
	_ReduceIteratorToLoopExpr                                       = _ReduceType(111)
	_ReduceForToLoopExpr                                            = _ReduceType(112)
	_ReduceSequenceExprToOptionalSequenceExpr                       = _ReduceType(113)
	_ReduceNilToOptionalSequenceExpr                                = _ReduceType(114)
	_ReduceSequenceExprToForAssignment                              = _ReduceType(115)
	_ReduceAssignToForAssignment                                    = _ReduceType(116)
	_ReduceToCallExpr                                               = _ReduceType(117)
	_ReduceBindingToOptionalGenericBinding                          = _ReduceType(118)
	_ReduceNilToOptionalGenericBinding                              = _ReduceType(119)
	_ReduceGenericArgumentsToOptionalGenericArguments               = _ReduceType(120)
	_ReduceNilToOptionalGenericArguments                            = _ReduceType(121)
	_ReduceValueTypeToGenericArguments                              = _ReduceType(122)
	_ReduceAddToGenericArguments                                    = _ReduceType(123)
	_ReduceArgumentsToOptionalArguments                             = _ReduceType(124)
	_ReduceNilToOptionalArguments                                   = _ReduceType(125)
	_ReduceArgumentToArguments                                      = _ReduceType(126)
	_ReduceAddToArguments                                           = _ReduceType(127)
	_ReducePositionalToArgument                                     = _ReduceType(128)
	_ReduceNamedToArgument                                          = _ReduceType(129)
	_ReduceColonExpressionsToArgument                               = _ReduceType(130)
	_ReduceDotDotDotToArgument                                      = _ReduceType(131)
	_ReducePairToColonExpressions                                   = _ReduceType(132)
	_ReduceAddToColonExpressions                                    = _ReduceType(133)
	_ReduceExpressionToOptionalExpression                           = _ReduceType(134)
	_ReduceNilToOptionalExpression                                  = _ReduceType(135)
	_ReduceParseErrorExprToAtomExpr                                 = _ReduceType(136)
	_ReduceLiteralExprToAtomExpr                                    = _ReduceType(137)
	_ReduceIdentifierExprToAtomExpr                                 = _ReduceType(138)
	_ReduceBlockExprToAtomExpr                                      = _ReduceType(139)
	_ReduceAnonymousFuncExprToAtomExpr                              = _ReduceType(140)
	_ReduceInitializeExprToAtomExpr                                 = _ReduceType(141)
	_ReduceImplicitStructExprToAtomExpr                             = _ReduceType(142)
	_ReduceToParseErrorExpr                                         = _ReduceType(143)
	_ReduceTrueToLiteralExpr                                        = _ReduceType(144)
	_ReduceFalseToLiteralExpr                                       = _ReduceType(145)
	_ReduceIntegerLiteralToLiteralExpr                              = _ReduceType(146)
	_ReduceFloatLiteralToLiteralExpr                                = _ReduceType(147)
	_ReduceRuneLiteralToLiteralExpr                                 = _ReduceType(148)
	_ReduceStringLiteralToLiteralExpr                               = _ReduceType(149)
	_ReduceToIdentifierExpr                                         = _ReduceType(150)
	_ReduceToBlockExpr                                              = _ReduceType(151)
	_ReduceToInitializeExpr                                         = _ReduceType(152)
	_ReduceToImplicitStructExpr                                     = _ReduceType(153)
	_ReduceAtomExprToAccessibleExpr                                 = _ReduceType(154)
	_ReduceAccessExprToAccessibleExpr                               = _ReduceType(155)
	_ReduceCallExprToAccessibleExpr                                 = _ReduceType(156)
	_ReduceIndexExprToAccessibleExpr                                = _ReduceType(157)
	_ReduceToAccessExpr                                             = _ReduceType(158)
	_ReduceToIndexExpr                                              = _ReduceType(159)
	_ReduceAccessibleExprToPostfixableExpr                          = _ReduceType(160)
	_ReducePostfixUnaryExprToPostfixableExpr                        = _ReduceType(161)
	_ReduceToPostfixUnaryExpr                                       = _ReduceType(162)
	_ReducePostfixableExprToPrefixableExpr                          = _ReduceType(163)
	_ReducePrefixUnaryExprToPrefixableExpr                          = _ReduceType(164)
	_ReduceToPrefixUnaryExpr                                        = _ReduceType(165)
	_ReduceNotToPrefixUnaryOp                                       = _ReduceType(166)
	_ReduceBitNegToPrefixUnaryOp                                    = _ReduceType(167)
	_ReduceSubToPrefixUnaryOp                                       = _ReduceType(168)
	_ReduceMulToPrefixUnaryOp                                       = _ReduceType(169)
	_ReduceBitAndToPrefixUnaryOp                                    = _ReduceType(170)
	_ReducePrefixableExprToMulExpr                                  = _ReduceType(171)
	_ReduceBinaryMulExprToMulExpr                                   = _ReduceType(172)
	_ReduceToBinaryMulExpr                                          = _ReduceType(173)
	_ReduceMulToMulOp                                               = _ReduceType(174)
	_ReduceDivToMulOp                                               = _ReduceType(175)
	_ReduceModToMulOp                                               = _ReduceType(176)
	_ReduceBitAndToMulOp                                            = _ReduceType(177)
	_ReduceBitLshiftToMulOp                                         = _ReduceType(178)
	_ReduceBitRshiftToMulOp                                         = _ReduceType(179)
	_ReduceMulExprToAddExpr                                         = _ReduceType(180)
	_ReduceBinaryAddExprToAddExpr                                   = _ReduceType(181)
	_ReduceToBinaryAddExpr                                          = _ReduceType(182)
	_ReduceAddToAddOp                                               = _ReduceType(183)
	_ReduceSubToAddOp                                               = _ReduceType(184)
	_ReduceBitOrToAddOp                                             = _ReduceType(185)
	_ReduceBitXorToAddOp                                            = _ReduceType(186)
	_ReduceAddExprToCmpExpr                                         = _ReduceType(187)
	_ReduceBinaryCmpExprToCmpExpr                                   = _ReduceType(188)
	_ReduceToBinaryCmpExpr                                          = _ReduceType(189)
	_ReduceEqualToCmpOp                                             = _ReduceType(190)
	_ReduceNotEqualToCmpOp                                          = _ReduceType(191)
	_ReduceLessToCmpOp                                              = _ReduceType(192)
	_ReduceLessOrEqualToCmpOp                                       = _ReduceType(193)
	_ReduceGreaterToCmpOp                                           = _ReduceType(194)
	_ReduceGreaterOrEqualToCmpOp                                    = _ReduceType(195)
	_ReduceCmpExprToAndExpr                                         = _ReduceType(196)
	_ReduceBinaryAndExprToAndExpr                                   = _ReduceType(197)
	_ReduceToBinaryAndExpr                                          = _ReduceType(198)
	_ReduceAndExprToOrExpr                                          = _ReduceType(199)
	_ReduceBinaryOrExprToOrExpr                                     = _ReduceType(200)
	_ReduceToBinaryOrExpr                                           = _ReduceType(201)
	_ReduceExplicitStructDefToInitializableType                     = _ReduceType(202)
	_ReduceSliceToInitializableType                                 = _ReduceType(203)
	_ReduceArrayToInitializableType                                 = _ReduceType(204)
	_ReduceMapToInitializableType                                   = _ReduceType(205)
	_ReduceInitializableTypeToAtomType                              = _ReduceType(206)
	_ReduceNamedToAtomType                                          = _ReduceType(207)
	_ReduceExternNamedToAtomType                                    = _ReduceType(208)
	_ReduceInferredToAtomType                                       = _ReduceType(209)
	_ReduceImplicitStructDefToAtomType                              = _ReduceType(210)
	_ReduceExplicitEnumDefToAtomType                                = _ReduceType(211)
	_ReduceImplicitEnumDefToAtomType                                = _ReduceType(212)
	_ReduceTraitDefToAtomType                                       = _ReduceType(213)
	_ReduceFuncTypeToAtomType                                       = _ReduceType(214)
	_ReduceParseErrorTypeToAtomType                                 = _ReduceType(215)
	_ReduceToParseErrorType                                         = _ReduceType(216)
	_ReduceAtomTypeToReturnableType                                 = _ReduceType(217)
	_ReducePrefixedTypeToReturnableType                             = _ReduceType(218)
	_ReduceToPrefixedType                                           = _ReduceType(219)
	_ReduceQuestionToPrefixTypeOp                                   = _ReduceType(220)
	_ReduceExclaimToPrefixTypeOp                                    = _ReduceType(221)
	_ReduceBitAndToPrefixTypeOp                                     = _ReduceType(222)
	_ReduceBitNegToPrefixTypeOp                                     = _ReduceType(223)
	_ReduceTildeTildeToPrefixTypeOp                                 = _ReduceType(224)
	_ReduceReturnableTypeToValueType                                = _ReduceType(225)
	_ReduceTraitOpTypeToValueType                                   = _ReduceType(226)
	_ReduceToTraitOpType                                            = _ReduceType(227)
	_ReduceMulToTraitOp                                             = _ReduceType(228)
	_ReduceAddToTraitOp                                             = _ReduceType(229)
	_ReduceSubToTraitOp                                             = _ReduceType(230)
	_ReduceDefinitionToTypeDef                                      = _ReduceType(231)
	_ReduceConstrainedDefToTypeDef                                  = _ReduceType(232)
	_ReduceAliasToTypeDef                                           = _ReduceType(233)
	_ReduceUnconstrainedToGenericParameterDef                       = _ReduceType(234)
	_ReduceConstrainedToGenericParameterDef                         = _ReduceType(235)
	_ReduceGenericParameterDefToGenericParameterDefs                = _ReduceType(236)
	_ReduceAddToGenericParameterDefs                                = _ReduceType(237)
	_ReduceGenericParameterDefsToOptionalGenericParameterDefs       = _ReduceType(238)
	_ReduceNilToOptionalGenericParameterDefs                        = _ReduceType(239)
	_ReduceGenericToOptionalGenericParameters                       = _ReduceType(240)
	_ReduceNilToOptionalGenericParameters                           = _ReduceType(241)
	_ReduceExplicitToFieldDef                                       = _ReduceType(242)
	_ReduceImplicitToFieldDef                                       = _ReduceType(243)
	_ReduceUnsafeStatementToFieldDef                                = _ReduceType(244)
	_ReduceFieldDefToImplicitFieldDefs                              = _ReduceType(245)
	_ReduceAddToImplicitFieldDefs                                   = _ReduceType(246)
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefs             = _ReduceType(247)
	_ReduceNilToOptionalImplicitFieldDefs                           = _ReduceType(248)
	_ReduceToImplicitStructDef                                      = _ReduceType(249)
	_ReduceFieldDefToExplicitFieldDefs                              = _ReduceType(250)
	_ReduceImplicitToExplicitFieldDefs                              = _ReduceType(251)
	_ReduceExplicitToExplicitFieldDefs                              = _ReduceType(252)
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefs             = _ReduceType(253)
	_ReduceNilToOptionalExplicitFieldDefs                           = _ReduceType(254)
	_ReduceToExplicitStructDef                                      = _ReduceType(255)
	_ReduceFieldDefToEnumValueDef                                   = _ReduceType(256)
	_ReduceDefaultToEnumValueDef                                    = _ReduceType(257)
	_ReducePairToImplicitEnumValueDefs                              = _ReduceType(258)
	_ReduceAddToImplicitEnumValueDefs                               = _ReduceType(259)
	_ReduceToImplicitEnumDef                                        = _ReduceType(260)
	_ReduceExplicitPairToExplicitEnumValueDefs                      = _ReduceType(261)
	_ReduceImplicitPairToExplicitEnumValueDefs                      = _ReduceType(262)
	_ReduceExplicitAddToExplicitEnumValueDefs                       = _ReduceType(263)
	_ReduceImplicitAddToExplicitEnumValueDefs                       = _ReduceType(264)
	_ReduceToExplicitEnumDef                                        = _ReduceType(265)
	_ReduceFieldDefToTraitProperty                                  = _ReduceType(266)
	_ReduceMethodSignatureToTraitProperty                           = _ReduceType(267)
	_ReduceTraitPropertyToTraitProperties                           = _ReduceType(268)
	_ReduceImplicitToTraitProperties                                = _ReduceType(269)
	_ReduceExplicitToTraitProperties                                = _ReduceType(270)
	_ReduceTraitPropertiesToOptionalTraitProperties                 = _ReduceType(271)
	_ReduceNilToOptionalTraitProperties                             = _ReduceType(272)
	_ReduceToTraitDef                                               = _ReduceType(273)
	_ReduceReturnableTypeToReturnType                               = _ReduceType(274)
	_ReduceNilToReturnType                                          = _ReduceType(275)
	_ReduceArgToParameterDecl                                       = _ReduceType(276)
	_ReduceVarargToParameterDecl                                    = _ReduceType(277)
	_ReduceUnamedToParameterDecl                                    = _ReduceType(278)
	_ReduceUnnamedVarargToParameterDecl                             = _ReduceType(279)
	_ReduceParameterDeclToParameterDecls                            = _ReduceType(280)
	_ReduceAddToParameterDecls                                      = _ReduceType(281)
	_ReduceParameterDeclsToOptionalParameterDecls                   = _ReduceType(282)
	_ReduceNilToOptionalParameterDecls                              = _ReduceType(283)
	_ReduceToFuncType                                               = _ReduceType(284)
	_ReduceToMethodSignature                                        = _ReduceType(285)
	_ReduceInferredRefArgToParameterDef                             = _ReduceType(286)
	_ReduceInferredRefVarargToParameterDef                          = _ReduceType(287)
	_ReduceArgToParameterDef                                        = _ReduceType(288)
	_ReduceVarargToParameterDef                                     = _ReduceType(289)
	_ReduceParameterDefToParameterDefs                              = _ReduceType(290)
	_ReduceAddToParameterDefs                                       = _ReduceType(291)
	_ReduceParameterDefsToOptionalParameterDefs                     = _ReduceType(292)
	_ReduceNilToOptionalParameterDefs                               = _ReduceType(293)
	_ReduceFuncDefToNamedFuncDef                                    = _ReduceType(294)
	_ReduceMethodDefToNamedFuncDef                                  = _ReduceType(295)
	_ReduceAliasToNamedFuncDef                                      = _ReduceType(296)
	_ReduceToAnonymousFuncExpr                                      = _ReduceType(297)
	_ReduceNoSpecToPackageDef                                       = _ReduceType(298)
	_ReduceWithSpecToPackageDef                                     = _ReduceType(299)
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
	case _ReduceDefinitionToDefinitions:
		return "DefinitionToDefinitions"
	case _ReduceAddToDefinitions:
		return "AddToDefinitions"
	case _ReducePackageDefToDefinition:
		return "PackageDefToDefinition"
	case _ReduceTypeDefToDefinition:
		return "TypeDefToDefinition"
	case _ReduceNamedFuncDefToDefinition:
		return "NamedFuncDefToDefinition"
	case _ReduceGlobalVarDefToDefinition:
		return "GlobalVarDefToDefinition"
	case _ReduceGlobalVarAssignmentToDefinition:
		return "GlobalVarAssignmentToDefinition"
	case _ReduceStatementBlockToDefinition:
		return "StatementBlockToDefinition"
	case _ReduceCommentGroupsToDefinition:
		return "CommentGroupsToDefinition"
	case _ReduceToStatementBlock:
		return "ToStatementBlock"
	case _ReduceEmptyListToStatements:
		return "EmptyListToStatements"
	case _ReduceAddToStatements:
		return "AddToStatements"
	case _ReduceImplicitToStatement:
		return "ImplicitToStatement"
	case _ReduceExplicitToStatement:
		return "ExplicitToStatement"
	case _ReduceUnsafeStatementToSimpleStatementBody:
		return "UnsafeStatementToSimpleStatementBody"
	case _ReduceExpressionOrImproperStructStatementToSimpleStatementBody:
		return "ExpressionOrImproperStructStatementToSimpleStatementBody"
	case _ReduceCallbackOpStatementToSimpleStatementBody:
		return "CallbackOpStatementToSimpleStatementBody"
	case _ReduceJumpStatementToSimpleStatementBody:
		return "JumpStatementToSimpleStatementBody"
	case _ReduceAssignStatementToSimpleStatementBody:
		return "AssignStatementToSimpleStatementBody"
	case _ReduceUnaryOpAssignStatementToSimpleStatementBody:
		return "UnaryOpAssignStatementToSimpleStatementBody"
	case _ReduceBinaryOpAssignStatementToSimpleStatementBody:
		return "BinaryOpAssignStatementToSimpleStatementBody"
	case _ReduceFallthroughStatementToSimpleStatementBody:
		return "FallthroughStatementToSimpleStatementBody"
	case _ReduceSimpleStatementBodyToStatementBody:
		return "SimpleStatementBodyToStatementBody"
	case _ReduceImportStatementToStatementBody:
		return "ImportStatementToStatementBody"
	case _ReduceCaseBranchStatementToStatementBody:
		return "CaseBranchStatementToStatementBody"
	case _ReduceDefaultBranchStatementToStatementBody:
		return "DefaultBranchStatementToStatementBody"
	case _ReduceSimpleStatementBodyToOptionalSimpleStatementBody:
		return "SimpleStatementBodyToOptionalSimpleStatementBody"
	case _ReduceNilToOptionalSimpleStatementBody:
		return "NilToOptionalSimpleStatementBody"
	case _ReduceToExpressionOrImproperStructStatement:
		return "ToExpressionOrImproperStructStatement"
	case _ReduceAsyncToCallbackOp:
		return "AsyncToCallbackOp"
	case _ReduceDeferToCallbackOp:
		return "DeferToCallbackOp"
	case _ReduceToCallbackOpStatement:
		return "ToCallbackOpStatement"
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
	case _ReduceUnlabeledNoValueToJumpStatement:
		return "UnlabeledNoValueToJumpStatement"
	case _ReduceUnlabeledValuedToJumpStatement:
		return "UnlabeledValuedToJumpStatement"
	case _ReduceLabeledNoValueToJumpStatement:
		return "LabeledNoValueToJumpStatement"
	case _ReduceLabeledValuedToJumpStatement:
		return "LabeledValuedToJumpStatement"
	case _ReduceReturnToJumpType:
		return "ReturnToJumpType"
	case _ReduceBreakToJumpType:
		return "BreakToJumpType"
	case _ReduceContinueToJumpType:
		return "ContinueToJumpType"
	case _ReduceExpressionToExpressions:
		return "ExpressionToExpressions"
	case _ReduceAddToExpressions:
		return "AddToExpressions"
	case _ReduceToFallthroughStatement:
		return "ToFallthroughStatement"
	case _ReduceToAssignStatement:
		return "ToAssignStatement"
	case _ReduceToUnaryOpAssignStatement:
		return "ToUnaryOpAssignStatement"
	case _ReduceToBinaryOpAssignStatement:
		return "ToBinaryOpAssignStatement"
	case _ReduceSingleToImportStatement:
		return "SingleToImportStatement"
	case _ReduceMultipleToImportStatement:
		return "MultipleToImportStatement"
	case _ReduceStringLiteralToImportClause:
		return "StringLiteralToImportClause"
	case _ReduceAliasToImportClause:
		return "AliasToImportClause"
	case _ReduceImplicitToImportClauseTerminal:
		return "ImplicitToImportClauseTerminal"
	case _ReduceExplicitToImportClauseTerminal:
		return "ExplicitToImportClauseTerminal"
	case _ReduceFirstToImportClauses:
		return "FirstToImportClauses"
	case _ReduceAddToImportClauses:
		return "AddToImportClauses"
	case _ReduceCasePatternToCasePatterns:
		return "CasePatternToCasePatterns"
	case _ReduceMultipleToCasePatterns:
		return "MultipleToCasePatterns"
	case _ReduceToVarDeclPattern:
		return "ToVarDeclPattern"
	case _ReduceVarToVarOrLet:
		return "VarToVarOrLet"
	case _ReduceLetToVarOrLet:
		return "LetToVarOrLet"
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
	case _ReduceDotDotDotToFieldVarPattern:
		return "DotDotDotToFieldVarPattern"
	case _ReduceValueTypeToOptionalValueType:
		return "ValueTypeToOptionalValueType"
	case _ReduceNilToOptionalValueType:
		return "NilToOptionalValueType"
	case _ReduceToAssignPattern:
		return "ToAssignPattern"
	case _ReduceAssignPatternToCasePattern:
		return "AssignPatternToCasePattern"
	case _ReduceEnumMatchPatternToCasePattern:
		return "EnumMatchPatternToCasePattern"
	case _ReduceEnumVarDeclPatternToCasePattern:
		return "EnumVarDeclPatternToCasePattern"
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
	case _ReduceOrExprToSequenceExpr:
		return "OrExprToSequenceExpr"
	case _ReduceVarDeclPatternToSequenceExpr:
		return "VarDeclPatternToSequenceExpr"
	case _ReduceAssignVarPatternToSequenceExpr:
		return "AssignVarPatternToSequenceExpr"
	case _ReduceNoElseToIfExpr:
		return "NoElseToIfExpr"
	case _ReduceIfElseToIfExpr:
		return "IfElseToIfExpr"
	case _ReduceMultiIfElseToIfExpr:
		return "MultiIfElseToIfExpr"
	case _ReduceSequenceExprToCondition:
		return "SequenceExprToCondition"
	case _ReduceCaseToCondition:
		return "CaseToCondition"
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
	case _ReduceSequenceExprToForAssignment:
		return "SequenceExprToForAssignment"
	case _ReduceAssignToForAssignment:
		return "AssignToForAssignment"
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
	case _ReduceDotDotDotToArgument:
		return "DotDotDotToArgument"
	case _ReducePairToColonExpressions:
		return "PairToColonExpressions"
	case _ReduceAddToColonExpressions:
		return "AddToColonExpressions"
	case _ReduceExpressionToOptionalExpression:
		return "ExpressionToOptionalExpression"
	case _ReduceNilToOptionalExpression:
		return "NilToOptionalExpression"
	case _ReduceParseErrorExprToAtomExpr:
		return "ParseErrorExprToAtomExpr"
	case _ReduceLiteralExprToAtomExpr:
		return "LiteralExprToAtomExpr"
	case _ReduceIdentifierExprToAtomExpr:
		return "IdentifierExprToAtomExpr"
	case _ReduceBlockExprToAtomExpr:
		return "BlockExprToAtomExpr"
	case _ReduceAnonymousFuncExprToAtomExpr:
		return "AnonymousFuncExprToAtomExpr"
	case _ReduceInitializeExprToAtomExpr:
		return "InitializeExprToAtomExpr"
	case _ReduceImplicitStructExprToAtomExpr:
		return "ImplicitStructExprToAtomExpr"
	case _ReduceToParseErrorExpr:
		return "ToParseErrorExpr"
	case _ReduceTrueToLiteralExpr:
		return "TrueToLiteralExpr"
	case _ReduceFalseToLiteralExpr:
		return "FalseToLiteralExpr"
	case _ReduceIntegerLiteralToLiteralExpr:
		return "IntegerLiteralToLiteralExpr"
	case _ReduceFloatLiteralToLiteralExpr:
		return "FloatLiteralToLiteralExpr"
	case _ReduceRuneLiteralToLiteralExpr:
		return "RuneLiteralToLiteralExpr"
	case _ReduceStringLiteralToLiteralExpr:
		return "StringLiteralToLiteralExpr"
	case _ReduceToIdentifierExpr:
		return "ToIdentifierExpr"
	case _ReduceToBlockExpr:
		return "ToBlockExpr"
	case _ReduceToInitializeExpr:
		return "ToInitializeExpr"
	case _ReduceToImplicitStructExpr:
		return "ToImplicitStructExpr"
	case _ReduceAtomExprToAccessibleExpr:
		return "AtomExprToAccessibleExpr"
	case _ReduceAccessExprToAccessibleExpr:
		return "AccessExprToAccessibleExpr"
	case _ReduceCallExprToAccessibleExpr:
		return "CallExprToAccessibleExpr"
	case _ReduceIndexExprToAccessibleExpr:
		return "IndexExprToAccessibleExpr"
	case _ReduceToAccessExpr:
		return "ToAccessExpr"
	case _ReduceToIndexExpr:
		return "ToIndexExpr"
	case _ReduceAccessibleExprToPostfixableExpr:
		return "AccessibleExprToPostfixableExpr"
	case _ReducePostfixUnaryExprToPostfixableExpr:
		return "PostfixUnaryExprToPostfixableExpr"
	case _ReduceToPostfixUnaryExpr:
		return "ToPostfixUnaryExpr"
	case _ReducePostfixableExprToPrefixableExpr:
		return "PostfixableExprToPrefixableExpr"
	case _ReducePrefixUnaryExprToPrefixableExpr:
		return "PrefixUnaryExprToPrefixableExpr"
	case _ReduceToPrefixUnaryExpr:
		return "ToPrefixUnaryExpr"
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
	case _ReducePrefixableExprToMulExpr:
		return "PrefixableExprToMulExpr"
	case _ReduceBinaryMulExprToMulExpr:
		return "BinaryMulExprToMulExpr"
	case _ReduceToBinaryMulExpr:
		return "ToBinaryMulExpr"
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
	case _ReduceMulExprToAddExpr:
		return "MulExprToAddExpr"
	case _ReduceBinaryAddExprToAddExpr:
		return "BinaryAddExprToAddExpr"
	case _ReduceToBinaryAddExpr:
		return "ToBinaryAddExpr"
	case _ReduceAddToAddOp:
		return "AddToAddOp"
	case _ReduceSubToAddOp:
		return "SubToAddOp"
	case _ReduceBitOrToAddOp:
		return "BitOrToAddOp"
	case _ReduceBitXorToAddOp:
		return "BitXorToAddOp"
	case _ReduceAddExprToCmpExpr:
		return "AddExprToCmpExpr"
	case _ReduceBinaryCmpExprToCmpExpr:
		return "BinaryCmpExprToCmpExpr"
	case _ReduceToBinaryCmpExpr:
		return "ToBinaryCmpExpr"
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
	case _ReduceCmpExprToAndExpr:
		return "CmpExprToAndExpr"
	case _ReduceBinaryAndExprToAndExpr:
		return "BinaryAndExprToAndExpr"
	case _ReduceToBinaryAndExpr:
		return "ToBinaryAndExpr"
	case _ReduceAndExprToOrExpr:
		return "AndExprToOrExpr"
	case _ReduceBinaryOrExprToOrExpr:
		return "BinaryOrExprToOrExpr"
	case _ReduceToBinaryOrExpr:
		return "ToBinaryOrExpr"
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
	case _ReduceParseErrorTypeToAtomType:
		return "ParseErrorTypeToAtomType"
	case _ReduceToParseErrorType:
		return "ToParseErrorType"
	case _ReduceAtomTypeToReturnableType:
		return "AtomTypeToReturnableType"
	case _ReducePrefixedTypeToReturnableType:
		return "PrefixedTypeToReturnableType"
	case _ReduceToPrefixedType:
		return "ToPrefixedType"
	case _ReduceQuestionToPrefixTypeOp:
		return "QuestionToPrefixTypeOp"
	case _ReduceExclaimToPrefixTypeOp:
		return "ExclaimToPrefixTypeOp"
	case _ReduceBitAndToPrefixTypeOp:
		return "BitAndToPrefixTypeOp"
	case _ReduceBitNegToPrefixTypeOp:
		return "BitNegToPrefixTypeOp"
	case _ReduceTildeTildeToPrefixTypeOp:
		return "TildeTildeToPrefixTypeOp"
	case _ReduceReturnableTypeToValueType:
		return "ReturnableTypeToValueType"
	case _ReduceTraitOpTypeToValueType:
		return "TraitOpTypeToValueType"
	case _ReduceToTraitOpType:
		return "ToTraitOpType"
	case _ReduceMulToTraitOp:
		return "MulToTraitOp"
	case _ReduceAddToTraitOp:
		return "AddToTraitOp"
	case _ReduceSubToTraitOp:
		return "SubToTraitOp"
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
	case _ReduceInferredRefArgToParameterDef:
		return "InferredRefArgToParameterDef"
	case _ReduceInferredRefVarargToParameterDef:
		return "InferredRefVarargToParameterDef"
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
	case _ReduceFuncDefToNamedFuncDef:
		return "FuncDefToNamedFuncDef"
	case _ReduceMethodDefToNamedFuncDef:
		return "MethodDefToNamedFuncDef"
	case _ReduceAliasToNamedFuncDef:
		return "AliasToNamedFuncDef"
	case _ReduceToAnonymousFuncExpr:
		return "ToAnonymousFuncExpr"
	case _ReduceNoSpecToPackageDef:
		return "NoSpecToPackageDef"
	case _ReduceWithSpecToPackageDef:
		return "WithSpecToPackageDef"
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
	_State428 = _StateId(428)
	_State429 = _StateId(429)
	_State430 = _StateId(430)
	_State431 = _StateId(431)
	_State432 = _StateId(432)
	_State433 = _StateId(433)
	_State434 = _StateId(434)
	_State435 = _StateId(435)
	_State436 = _StateId(436)
	_State437 = _StateId(437)
	_State438 = _StateId(438)
	_State439 = _StateId(439)
	_State440 = _StateId(440)
	_State441 = _StateId(441)
	_State442 = _StateId(442)
	_State443 = _StateId(443)
	_State444 = _StateId(444)
	_State445 = _StateId(445)
	_State446 = _StateId(446)
	_State447 = _StateId(447)
	_State448 = _StateId(448)
	_State449 = _StateId(449)
	_State450 = _StateId(450)
	_State451 = _StateId(451)
	_State452 = _StateId(452)
	_State453 = _StateId(453)
	_State454 = _StateId(454)
	_State455 = _StateId(455)
	_State456 = _StateId(456)
	_State457 = _StateId(457)
	_State458 = _StateId(458)
	_State459 = _StateId(459)
	_State460 = _StateId(460)
	_State461 = _StateId(461)
	_State462 = _StateId(462)
	_State463 = _StateId(463)
	_State464 = _StateId(464)
	_State465 = _StateId(465)
	_State466 = _StateId(466)
	_State467 = _StateId(467)
	_State468 = _StateId(468)
)

type Symbol struct {
	SymbolId_ SymbolId

	Generic_ GenericSymbol

	Count             TokenCount
	Expression        Expression
	ParseError        ParseErrorSymbol
	SourceDefinition  SourceDefinition
	SourceDefinitions []SourceDefinition
	Statement         Statement
	Statements        []Statement
	TypeExpression    TypeExpression
	Value             TokenValue
}

func NewSymbol(token Token) (*Symbol, error) {
	symbol, ok := token.(*Symbol)
	if ok {
		return symbol, nil
	}

	symbol = &Symbol{SymbolId_: token.Id()}
	switch token.Id() {
	case NewlinesToken:
		val, ok := token.(TokenCount)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting TokenCount (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.Count = val
	case _EndMarker:
		val, ok := token.(GenericSymbol)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting GenericSymbol (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.Generic_ = val
	case ParseErrorToken:
		val, ok := token.(ParseErrorSymbol)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting ParseErrorSymbol (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.ParseError = val
	case CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, AsToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, DotDotDotToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken:
		val, ok := token.(TokenValue)
		if !ok {
			return nil, fmt.Errorf(
				"Invalid value type for token %s.  "+
					"Expecting TokenValue (%v)",
				token.Id(),
				token.Loc())
		}
		symbol.Value = val
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
	case NewlinesToken:
		loc, ok := interface{}(s.Count).(locator)
		if ok {
			return loc.Loc()
		}
	case ExpressionType, SequenceExprType, IfExprType, SwitchExprType, LoopExprType, OptionalSequenceExprType, CallExprType, OptionalExpressionType, AtomExprType, ParseErrorExprType, LiteralExprType, IdentifierExprType, BlockExprType, InitializeExprType, ImplicitStructExprType, AccessibleExprType, AccessExprType, IndexExprType, PostfixableExprType, PostfixUnaryExprType, PrefixableExprType, PrefixUnaryExprType, MulExprType, BinaryMulExprType, AddExprType, BinaryAddExprType, CmpExprType, BinaryCmpExprType, AndExprType, BinaryAndExprType, OrExprType, BinaryOrExprType, AnonymousFuncExprType:
		loc, ok := interface{}(s.Expression).(locator)
		if ok {
			return loc.Loc()
		}
	case ParseErrorToken:
		loc, ok := interface{}(s.ParseError).(locator)
		if ok {
			return loc.Loc()
		}
	case DefinitionType, TypeDefType, NamedFuncDefType, PackageDefType:
		loc, ok := interface{}(s.SourceDefinition).(locator)
		if ok {
			return loc.Loc()
		}
	case SourceType, DefinitionsType:
		loc, ok := interface{}(s.SourceDefinitions).(locator)
		if ok {
			return loc.Loc()
		}
	case StatementType, SimpleStatementBodyType, StatementBodyType, OptionalSimpleStatementBodyType, ExpressionOrImproperStructStatementType, CallbackOpStatementType, UnsafeStatementType, JumpStatementType, FallthroughStatementType, AssignStatementType, UnaryOpAssignStatementType, BinaryOpAssignStatementType, ImportStatementType:
		loc, ok := interface{}(s.Statement).(locator)
		if ok {
			return loc.Loc()
		}
	case StatementsType:
		loc, ok := interface{}(s.Statements).(locator)
		if ok {
			return loc.Loc()
		}
	case OptionalValueTypeType, InitializableTypeType, AtomTypeType, ParseErrorTypeType, ReturnableTypeType, PrefixedTypeType, ValueTypeType, TraitOpTypeType, ImplicitStructDefType, ExplicitStructDefType, ImplicitEnumDefType, ExplicitEnumDefType, TraitDefType, ReturnTypeType, FuncTypeType:
		loc, ok := interface{}(s.TypeExpression).(locator)
		if ok {
			return loc.Loc()
		}
	case CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, AsToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, DotDotDotToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, CallbackOpType, UnaryOpAssignType, BinaryOpAssignType, JumpTypeType, VarOrLetType, OptionalLabelDeclType, PrefixUnaryOpType, MulOpType, AddOpType, CmpOpType, PrefixTypeOpType, TraitOpType:
		loc, ok := interface{}(s.Value).(locator)
		if ok {
			return loc.Loc()
		}
	}
	return s.Generic_.Loc()
}

func (s *Symbol) End() Location {
	type locator interface{ End() Location }
	switch s.SymbolId_ {
	case NewlinesToken:
		loc, ok := interface{}(s.Count).(locator)
		if ok {
			return loc.End()
		}
	case ExpressionType, SequenceExprType, IfExprType, SwitchExprType, LoopExprType, OptionalSequenceExprType, CallExprType, OptionalExpressionType, AtomExprType, ParseErrorExprType, LiteralExprType, IdentifierExprType, BlockExprType, InitializeExprType, ImplicitStructExprType, AccessibleExprType, AccessExprType, IndexExprType, PostfixableExprType, PostfixUnaryExprType, PrefixableExprType, PrefixUnaryExprType, MulExprType, BinaryMulExprType, AddExprType, BinaryAddExprType, CmpExprType, BinaryCmpExprType, AndExprType, BinaryAndExprType, OrExprType, BinaryOrExprType, AnonymousFuncExprType:
		loc, ok := interface{}(s.Expression).(locator)
		if ok {
			return loc.End()
		}
	case ParseErrorToken:
		loc, ok := interface{}(s.ParseError).(locator)
		if ok {
			return loc.End()
		}
	case DefinitionType, TypeDefType, NamedFuncDefType, PackageDefType:
		loc, ok := interface{}(s.SourceDefinition).(locator)
		if ok {
			return loc.End()
		}
	case SourceType, DefinitionsType:
		loc, ok := interface{}(s.SourceDefinitions).(locator)
		if ok {
			return loc.End()
		}
	case StatementType, SimpleStatementBodyType, StatementBodyType, OptionalSimpleStatementBodyType, ExpressionOrImproperStructStatementType, CallbackOpStatementType, UnsafeStatementType, JumpStatementType, FallthroughStatementType, AssignStatementType, UnaryOpAssignStatementType, BinaryOpAssignStatementType, ImportStatementType:
		loc, ok := interface{}(s.Statement).(locator)
		if ok {
			return loc.End()
		}
	case StatementsType:
		loc, ok := interface{}(s.Statements).(locator)
		if ok {
			return loc.End()
		}
	case OptionalValueTypeType, InitializableTypeType, AtomTypeType, ParseErrorTypeType, ReturnableTypeType, PrefixedTypeType, ValueTypeType, TraitOpTypeType, ImplicitStructDefType, ExplicitStructDefType, ImplicitEnumDefType, ExplicitEnumDefType, TraitDefType, ReturnTypeType, FuncTypeType:
		loc, ok := interface{}(s.TypeExpression).(locator)
		if ok {
			return loc.End()
		}
	case CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, AsToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, DotDotDotToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, CallbackOpType, UnaryOpAssignType, BinaryOpAssignType, JumpTypeType, VarOrLetType, OptionalLabelDeclType, PrefixUnaryOpType, MulOpType, AddOpType, CmpOpType, PrefixTypeOpType, TraitOpType:
		loc, ok := interface{}(s.Value).(locator)
		if ok {
			return loc.End()
		}
	}
	return s.Generic_.End()
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
			token = GenericSymbol{
				SymbolId: _EndMarker,
				StartPos: stack.lexer.CurrentLocation(),
			}
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
		symbol.SourceDefinitions, err = reducer.ToSource(args[0].Generic_)
	case _ReduceNewlinesToOptionalDefinitions:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalDefinitionsType
		symbol.Generic_, err = reducer.NewlinesToOptionalDefinitions(args[0].Count)
	case _ReduceDefinitionsToOptionalDefinitions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OptionalDefinitionsType
		symbol.Generic_, err = reducer.DefinitionsToOptionalDefinitions(args[0].Generic_, args[1].SourceDefinitions, args[2].Generic_)
	case _ReduceNilToOptionalDefinitions:
		symbol.SymbolId_ = OptionalDefinitionsType
		symbol.Generic_, err = reducer.NilToOptionalDefinitions()
	case _ReduceNewlinesToOptionalNewlines:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalNewlinesType
		symbol.Generic_, err = reducer.NewlinesToOptionalNewlines(args[0].Count)
	case _ReduceNilToOptionalNewlines:
		symbol.SymbolId_ = OptionalNewlinesType
		symbol.Generic_, err = reducer.NilToOptionalNewlines()
	case _ReduceDefinitionToDefinitions:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionsType
		symbol.SourceDefinitions, err = reducer.DefinitionToDefinitions(args[0].SourceDefinition)
	case _ReduceAddToDefinitions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = DefinitionsType
		symbol.SourceDefinitions, err = reducer.AddToDefinitions(args[0].SourceDefinitions, args[1].Count, args[2].SourceDefinition)
	case _ReducePackageDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.SourceDefinition, err = reducer.PackageDefToDefinition(args[0].SourceDefinition)
	case _ReduceTypeDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.SourceDefinition, err = reducer.TypeDefToDefinition(args[0].SourceDefinition)
	case _ReduceNamedFuncDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.SourceDefinition, err = reducer.NamedFuncDefToDefinition(args[0].SourceDefinition)
	case _ReduceGlobalVarDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.SourceDefinition, err = reducer.GlobalVarDefToDefinition(args[0].Generic_)
	case _ReduceGlobalVarAssignmentToDefinition:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = DefinitionType
		symbol.SourceDefinition, err = reducer.GlobalVarAssignmentToDefinition(args[0].Generic_, args[1].Value, args[2].Expression)
	case _ReduceStatementBlockToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.SourceDefinition, err = reducer.StatementBlockToDefinition(args[0].Generic_)
	case _ReduceCommentGroupsToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		symbol.SourceDefinition, err = reducer.CommentGroupsToDefinition(args[0].Value)
	case _ReduceToStatementBlock:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = StatementBlockType
		symbol.Generic_, err = reducer.ToStatementBlock(args[0].Value, args[1].Statements, args[2].Value)
	case _ReduceEmptyListToStatements:
		symbol.SymbolId_ = StatementsType
		symbol.Statements, err = reducer.EmptyListToStatements()
	case _ReduceAddToStatements:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementsType
		symbol.Statements, err = reducer.AddToStatements(args[0].Statements, args[1].Statement)
	case _ReduceImplicitToStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementType
		symbol.Statement, err = reducer.ImplicitToStatement(args[0].Statement, args[1].Count)
	case _ReduceExplicitToStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementType
		symbol.Statement, err = reducer.ExplicitToStatement(args[0].Statement, args[1].Value)
	case _ReduceUnsafeStatementToSimpleStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementBodyType
		//line grammar.lr:111:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceExpressionOrImproperStructStatementToSimpleStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementBodyType
		//line grammar.lr:113:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceCallbackOpStatementToSimpleStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementBodyType
		//line grammar.lr:116:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceJumpStatementToSimpleStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementBodyType
		//line grammar.lr:119:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceAssignStatementToSimpleStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementBodyType
		//line grammar.lr:124:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceUnaryOpAssignStatementToSimpleStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementBodyType
		//line grammar.lr:130:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceBinaryOpAssignStatementToSimpleStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementBodyType
		//line grammar.lr:131:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceFallthroughStatementToSimpleStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementBodyType
		//line grammar.lr:135:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceSimpleStatementBodyToStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementBodyType
		//line grammar.lr:138:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceImportStatementToStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementBodyType
		//line grammar.lr:141:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceCaseBranchStatementToStatementBody:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = StatementBodyType
		symbol.Statement, err = reducer.CaseBranchStatementToStatementBody(args[0].Value, args[1].Generic_, args[2].Value, args[3].Statement)
	case _ReduceDefaultBranchStatementToStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = StatementBodyType
		symbol.Statement, err = reducer.DefaultBranchStatementToStatementBody(args[0].Value, args[1].Value, args[2].Statement)
	case _ReduceSimpleStatementBodyToOptionalSimpleStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalSimpleStatementBodyType
		//line grammar.lr:150:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceNilToOptionalSimpleStatementBody:
		symbol.SymbolId_ = OptionalSimpleStatementBodyType
		symbol.Statement, err = reducer.NilToOptionalSimpleStatementBody()
	case _ReduceToExpressionOrImproperStructStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionOrImproperStructStatementType
		symbol.Statement, err = reducer.ToExpressionOrImproperStructStatement(args[0].Generic_)
	case _ReduceAsyncToCallbackOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CallbackOpType
		//line grammar.lr:164:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDeferToCallbackOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CallbackOpType
		//line grammar.lr:165:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceToCallbackOpStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CallbackOpStatementType
		symbol.Statement, err = reducer.ToCallbackOpStatement(args[0].Value, args[1].Expression)
	case _ReduceAddOneAssignToUnaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpAssignType
		//line grammar.lr:190:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubOneAssignToUnaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpAssignType
		//line grammar.lr:191:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:194:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:195:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:196:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDivAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:197:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceModAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:198:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:199:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:200:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitOrAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:201:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitXorAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:202:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitLshiftAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:203:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitRshiftAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:204:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceToUnsafeStatement:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = UnsafeStatementType
		symbol.Statement, err = reducer.ToUnsafeStatement(args[0].Value, args[1].Value, args[2].Value, args[3].Value, args[4].Value)
	case _ReduceUnlabeledNoValueToJumpStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpStatementType
		symbol.Statement, err = reducer.UnlabeledNoValueToJumpStatement(args[0].Value)
	case _ReduceUnlabeledValuedToJumpStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = JumpStatementType
		symbol.Statement, err = reducer.UnlabeledValuedToJumpStatement(args[0].Value, args[1].Generic_)
	case _ReduceLabeledNoValueToJumpStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = JumpStatementType
		symbol.Statement, err = reducer.LabeledNoValueToJumpStatement(args[0].Value, args[1].Value)
	case _ReduceLabeledValuedToJumpStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = JumpStatementType
		symbol.Statement, err = reducer.LabeledValuedToJumpStatement(args[0].Value, args[1].Value, args[2].Generic_)
	case _ReduceReturnToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		//line grammar.lr:225:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBreakToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		//line grammar.lr:226:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceContinueToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		//line grammar.lr:227:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceExpressionToExpressions:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionsType
		symbol.Generic_, err = reducer.ExpressionToExpressions(args[0].Expression)
	case _ReduceAddToExpressions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExpressionsType
		symbol.Generic_, err = reducer.AddToExpressions(args[0].Generic_, args[1].Value, args[2].Expression)
	case _ReduceToFallthroughStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FallthroughStatementType
		symbol.Statement, err = reducer.ToFallthroughStatement(args[0].Value)
	case _ReduceToAssignStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = AssignStatementType
		symbol.Statement, err = reducer.ToAssignStatement(args[0].Generic_, args[1].Value, args[2].Expression)
	case _ReduceToUnaryOpAssignStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = UnaryOpAssignStatementType
		symbol.Statement, err = reducer.ToUnaryOpAssignStatement(args[0].Expression, args[1].Value)
	case _ReduceToBinaryOpAssignStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BinaryOpAssignStatementType
		symbol.Statement, err = reducer.ToBinaryOpAssignStatement(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceSingleToImportStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportStatementType
		symbol.Statement, err = reducer.SingleToImportStatement(args[0].Value, args[1].Generic_)
	case _ReduceMultipleToImportStatement:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ImportStatementType
		symbol.Statement, err = reducer.MultipleToImportStatement(args[0].Value, args[1].Value, args[2].Generic_, args[3].Value)
	case _ReduceStringLiteralToImportClause:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ImportClauseType
		symbol.Generic_, err = reducer.StringLiteralToImportClause(args[0].Value)
	case _ReduceAliasToImportClause:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImportClauseType
		symbol.Generic_, err = reducer.AliasToImportClause(args[0].Value, args[1].Value, args[2].Value)
	case _ReduceImplicitToImportClauseTerminal:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClauseTerminalType
		symbol.Generic_, err = reducer.ImplicitToImportClauseTerminal(args[0].Generic_, args[1].Count)
	case _ReduceExplicitToImportClauseTerminal:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClauseTerminalType
		symbol.Generic_, err = reducer.ExplicitToImportClauseTerminal(args[0].Generic_, args[1].Value)
	case _ReduceFirstToImportClauses:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ImportClausesType
		symbol.Generic_, err = reducer.FirstToImportClauses(args[0].Generic_)
	case _ReduceAddToImportClauses:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClausesType
		symbol.Generic_, err = reducer.AddToImportClauses(args[0].Generic_, args[1].Generic_)
	case _ReduceCasePatternToCasePatterns:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CasePatternsType
		symbol.Generic_, err = reducer.CasePatternToCasePatterns(args[0].Generic_)
	case _ReduceMultipleToCasePatterns:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CasePatternsType
		symbol.Generic_, err = reducer.MultipleToCasePatterns(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceToVarDeclPattern:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = VarDeclPatternType
		symbol.Generic_, err = reducer.ToVarDeclPattern(args[0].Value, args[1].Generic_, args[2].TypeExpression)
	case _ReduceVarToVarOrLet:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarOrLetType
		//line grammar.lr:283:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLetToVarOrLet:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarOrLetType
		//line grammar.lr:284:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceIdentifierToVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarPatternType
		symbol.Generic_, err = reducer.IdentifierToVarPattern(args[0].Value)
	case _ReduceTuplePatternToVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarPatternType
		symbol.Generic_, err = reducer.TuplePatternToVarPattern(args[0].Generic_)
	case _ReduceToTuplePattern:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TuplePatternType
		symbol.Generic_, err = reducer.ToTuplePattern(args[0].Value, args[1].Generic_, args[2].Value)
	case _ReduceFieldVarPatternToFieldVarPatterns:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldVarPatternsType
		symbol.Generic_, err = reducer.FieldVarPatternToFieldVarPatterns(args[0].Generic_)
	case _ReduceAddToFieldVarPatterns:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = FieldVarPatternsType
		symbol.Generic_, err = reducer.AddToFieldVarPatterns(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReducePositionalToFieldVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldVarPatternType
		symbol.Generic_, err = reducer.PositionalToFieldVarPattern(args[0].Generic_)
	case _ReduceNamedToFieldVarPattern:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = FieldVarPatternType
		symbol.Generic_, err = reducer.NamedToFieldVarPattern(args[0].Value, args[1].Value, args[2].Generic_)
	case _ReduceDotDotDotToFieldVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldVarPatternType
		symbol.Generic_, err = reducer.DotDotDotToFieldVarPattern(args[0].Value)
	case _ReduceValueTypeToOptionalValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalValueTypeType
		//line grammar.lr:302:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceNilToOptionalValueType:
		symbol.SymbolId_ = OptionalValueTypeType
		symbol.TypeExpression, err = reducer.NilToOptionalValueType()
	case _ReduceToAssignPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AssignPatternType
		symbol.Generic_, err = reducer.ToAssignPattern(args[0].Expression)
	case _ReduceAssignPatternToCasePattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CasePatternType
		//line grammar.lr:313:4
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _ReduceEnumMatchPatternToCasePattern:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CasePatternType
		symbol.Generic_, err = reducer.EnumMatchPatternToCasePattern(args[0].Value, args[1].Value, args[2].Expression)
	case _ReduceEnumVarDeclPatternToCasePattern:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CasePatternType
		symbol.Generic_, err = reducer.EnumVarDeclPatternToCasePattern(args[0].Value, args[1].Value, args[2].Value, args[3].Generic_)
	case _ReduceIfExprToExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExpressionType
		symbol.Expression, err = reducer.IfExprToExpression(args[0].Value, args[1].Expression)
	case _ReduceSwitchExprToExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExpressionType
		symbol.Expression, err = reducer.SwitchExprToExpression(args[0].Value, args[1].Expression)
	case _ReduceLoopExprToExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExpressionType
		symbol.Expression, err = reducer.LoopExprToExpression(args[0].Value, args[1].Expression)
	case _ReduceSequenceExprToExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionType
		//line grammar.lr:331:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceLabelDeclToOptionalLabelDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalLabelDeclType
		symbol.Value, err = reducer.LabelDeclToOptionalLabelDecl(args[0].Value)
	case _ReduceUnlabelledToOptionalLabelDecl:
		symbol.SymbolId_ = OptionalLabelDeclType
		symbol.Value, err = reducer.UnlabelledToOptionalLabelDecl()
	case _ReduceOrExprToSequenceExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SequenceExprType
		//line grammar.lr:342:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceVarDeclPatternToSequenceExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SequenceExprType
		symbol.Expression, err = reducer.VarDeclPatternToSequenceExpr(args[0].Generic_)
	case _ReduceAssignVarPatternToSequenceExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = SequenceExprType
		symbol.Expression, err = reducer.AssignVarPatternToSequenceExpr(args[0].Value, args[1].Expression)
	case _ReduceNoElseToIfExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = IfExprType
		symbol.Expression, err = reducer.NoElseToIfExpr(args[0].Value, args[1].Generic_, args[2].Generic_)
	case _ReduceIfElseToIfExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = IfExprType
		symbol.Expression, err = reducer.IfElseToIfExpr(args[0].Value, args[1].Generic_, args[2].Generic_, args[3].Value, args[4].Generic_)
	case _ReduceMultiIfElseToIfExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = IfExprType
		symbol.Expression, err = reducer.MultiIfElseToIfExpr(args[0].Value, args[1].Generic_, args[2].Generic_, args[3].Value, args[4].Expression)
	case _ReduceSequenceExprToCondition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ConditionType
		symbol.Generic_, err = reducer.SequenceExprToCondition(args[0].Expression)
	case _ReduceCaseToCondition:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ConditionType
		symbol.Generic_, err = reducer.CaseToCondition(args[0].Value, args[1].Generic_, args[2].Value, args[3].Expression)
	case _ReduceToSwitchExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = SwitchExprType
		symbol.Expression, err = reducer.ToSwitchExpr(args[0].Value, args[1].Expression, args[2].Generic_)
	case _ReduceInfiniteToLoopExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LoopExprType
		symbol.Expression, err = reducer.InfiniteToLoopExpr(args[0].Value, args[1].Generic_)
	case _ReduceDoWhileToLoopExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = LoopExprType
		symbol.Expression, err = reducer.DoWhileToLoopExpr(args[0].Value, args[1].Generic_, args[2].Value, args[3].Expression)
	case _ReduceWhileToLoopExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = LoopExprType
		symbol.Expression, err = reducer.WhileToLoopExpr(args[0].Value, args[1].Expression, args[2].Value, args[3].Generic_)
	case _ReduceIteratorToLoopExpr:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = LoopExprType
		symbol.Expression, err = reducer.IteratorToLoopExpr(args[0].Value, args[1].Generic_, args[2].Value, args[3].Expression, args[4].Value, args[5].Generic_)
	case _ReduceForToLoopExpr:
		args := stack[len(stack)-8:]
		stack = stack[:len(stack)-8]
		symbol.SymbolId_ = LoopExprType
		symbol.Expression, err = reducer.ForToLoopExpr(args[0].Value, args[1].Generic_, args[2].Value, args[3].Expression, args[4].Value, args[5].Expression, args[6].Value, args[7].Generic_)
	case _ReduceSequenceExprToOptionalSequenceExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalSequenceExprType
		symbol.Expression, err = reducer.SequenceExprToOptionalSequenceExpr(args[0].Expression)
	case _ReduceNilToOptionalSequenceExpr:
		symbol.SymbolId_ = OptionalSequenceExprType
		symbol.Expression, err = reducer.NilToOptionalSequenceExpr()
	case _ReduceSequenceExprToForAssignment:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ForAssignmentType
		symbol.Generic_, err = reducer.SequenceExprToForAssignment(args[0].Expression)
	case _ReduceAssignToForAssignment:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ForAssignmentType
		symbol.Generic_, err = reducer.AssignToForAssignment(args[0].Generic_, args[1].Value, args[2].Expression)
	case _ReduceToCallExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = CallExprType
		symbol.Expression, err = reducer.ToCallExpr(args[0].Expression, args[1].Generic_, args[2].Value, args[3].Generic_, args[4].Value)
	case _ReduceBindingToOptionalGenericBinding:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OptionalGenericBindingType
		symbol.Generic_, err = reducer.BindingToOptionalGenericBinding(args[0].Value, args[1].Generic_, args[2].Value)
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
		symbol.Generic_, err = reducer.ValueTypeToGenericArguments(args[0].TypeExpression)
	case _ReduceAddToGenericArguments:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = GenericArgumentsType
		symbol.Generic_, err = reducer.AddToGenericArguments(args[0].Generic_, args[1].Value, args[2].TypeExpression)
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
		symbol.Generic_, err = reducer.AddToArguments(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReducePositionalToArgument:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentType
		symbol.Generic_, err = reducer.PositionalToArgument(args[0].Expression)
	case _ReduceNamedToArgument:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ArgumentType
		symbol.Generic_, err = reducer.NamedToArgument(args[0].Value, args[1].Value, args[2].Expression)
	case _ReduceColonExpressionsToArgument:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentType
		symbol.Generic_, err = reducer.ColonExpressionsToArgument(args[0].Generic_)
	case _ReduceDotDotDotToArgument:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentType
		symbol.Generic_, err = reducer.DotDotDotToArgument(args[0].Value)
	case _ReducePairToColonExpressions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ColonExpressionsType
		symbol.Generic_, err = reducer.PairToColonExpressions(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceAddToColonExpressions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ColonExpressionsType
		symbol.Generic_, err = reducer.AddToColonExpressions(args[0].Generic_, args[1].Value, args[2].Expression)
	case _ReduceExpressionToOptionalExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalExpressionType
		symbol.Expression, err = reducer.ExpressionToOptionalExpression(args[0].Expression)
	case _ReduceNilToOptionalExpression:
		symbol.SymbolId_ = OptionalExpressionType
		symbol.Expression, err = reducer.NilToOptionalExpression()
	case _ReduceParseErrorExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:469:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceLiteralExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:470:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceIdentifierExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:471:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBlockExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:472:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceAnonymousFuncExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:473:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceInitializeExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:474:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceImplicitStructExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:475:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToParseErrorExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParseErrorExprType
		symbol.Expression, err = reducer.ToParseErrorExpr(args[0].ParseError)
	case _ReduceTrueToLiteralExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralExprType
		symbol.Expression, err = reducer.TrueToLiteralExpr(args[0].Value)
	case _ReduceFalseToLiteralExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralExprType
		symbol.Expression, err = reducer.FalseToLiteralExpr(args[0].Value)
	case _ReduceIntegerLiteralToLiteralExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralExprType
		symbol.Expression, err = reducer.IntegerLiteralToLiteralExpr(args[0].Value)
	case _ReduceFloatLiteralToLiteralExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralExprType
		symbol.Expression, err = reducer.FloatLiteralToLiteralExpr(args[0].Value)
	case _ReduceRuneLiteralToLiteralExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralExprType
		symbol.Expression, err = reducer.RuneLiteralToLiteralExpr(args[0].Value)
	case _ReduceStringLiteralToLiteralExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = LiteralExprType
		symbol.Expression, err = reducer.StringLiteralToLiteralExpr(args[0].Value)
	case _ReduceToIdentifierExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = IdentifierExprType
		symbol.Expression, err = reducer.ToIdentifierExpr(args[0].Value)
	case _ReduceToBlockExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = BlockExprType
		symbol.Expression, err = reducer.ToBlockExpr(args[0].Value, args[1].Generic_)
	case _ReduceToInitializeExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = InitializeExprType
		symbol.Expression, err = reducer.ToInitializeExpr(args[0].TypeExpression, args[1].Value, args[2].Generic_, args[3].Value)
	case _ReduceToImplicitStructExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitStructExprType
		symbol.Expression, err = reducer.ToImplicitStructExpr(args[0].Value, args[1].Generic_, args[2].Value)
	case _ReduceAtomExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:497:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceAccessExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:498:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceCallExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:499:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceIndexExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:500:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToAccessExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = AccessExprType
		symbol.Expression, err = reducer.ToAccessExpr(args[0].Expression, args[1].Value, args[2].Value)
	case _ReduceToIndexExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = IndexExprType
		symbol.Expression, err = reducer.ToIndexExpr(args[0].Expression, args[1].Value, args[2].Generic_, args[3].Value)
	case _ReduceAccessibleExprToPostfixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixableExprType
		//line grammar.lr:507:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReducePostfixUnaryExprToPostfixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixableExprType
		//line grammar.lr:508:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToPostfixUnaryExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PostfixUnaryExprType
		symbol.Expression, err = reducer.ToPostfixUnaryExpr(args[0].Expression, args[1].Value)
	case _ReducePostfixableExprToPrefixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixableExprType
		//line grammar.lr:513:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReducePrefixUnaryExprToPrefixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixableExprType
		//line grammar.lr:514:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToPrefixUnaryExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PrefixUnaryExprType
		symbol.Expression, err = reducer.ToPrefixUnaryExpr(args[0].Value, args[1].Expression)
	case _ReduceNotToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:519:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:520:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:521:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:524:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:527:4
		symbol.Value = args[0].Value
		err = nil
	case _ReducePrefixableExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		//line grammar.lr:530:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryMulExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		//line grammar.lr:531:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToBinaryMulExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BinaryMulExprType
		symbol.Expression, err = reducer.ToBinaryMulExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceMulToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:536:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDivToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:537:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceModToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:538:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:539:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitLshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:540:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitRshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:541:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulExprToAddExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddExprType
		//line grammar.lr:544:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryAddExprToAddExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddExprType
		//line grammar.lr:545:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToBinaryAddExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BinaryAddExprType
		symbol.Expression, err = reducer.ToBinaryAddExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceAddToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:550:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:551:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitOrToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:552:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitXorToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:553:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddExprToCmpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpExprType
		//line grammar.lr:556:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryCmpExprToCmpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpExprType
		//line grammar.lr:557:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToBinaryCmpExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BinaryCmpExprType
		symbol.Expression, err = reducer.ToBinaryCmpExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:562:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceNotEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:563:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLessToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:564:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLessOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:565:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceGreaterToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:566:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceGreaterOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:567:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceCmpExprToAndExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AndExprType
		//line grammar.lr:570:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryAndExprToAndExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AndExprType
		//line grammar.lr:571:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToBinaryAndExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BinaryAndExprType
		symbol.Expression, err = reducer.ToBinaryAndExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceAndExprToOrExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OrExprType
		//line grammar.lr:576:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryOrExprToOrExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OrExprType
		//line grammar.lr:577:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToBinaryOrExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BinaryOrExprType
		symbol.Expression, err = reducer.ToBinaryOrExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceExplicitStructDefToInitializableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeType
		//line grammar.lr:588:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceSliceToInitializableType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = InitializableTypeType
		symbol.TypeExpression, err = reducer.SliceToInitializableType(args[0].Value, args[1].TypeExpression, args[2].Value)
	case _ReduceArrayToInitializableType:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = InitializableTypeType
		symbol.TypeExpression, err = reducer.ArrayToInitializableType(args[0].Value, args[1].TypeExpression, args[2].Value, args[3].Value, args[4].Value)
	case _ReduceMapToInitializableType:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = InitializableTypeType
		symbol.TypeExpression, err = reducer.MapToInitializableType(args[0].Value, args[1].TypeExpression, args[2].Value, args[3].TypeExpression, args[4].Value)
	case _ReduceInitializableTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:594:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceNamedToAtomType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = AtomTypeType
		symbol.TypeExpression, err = reducer.NamedToAtomType(args[0].Value, args[1].Generic_)
	case _ReduceExternNamedToAtomType:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = AtomTypeType
		symbol.TypeExpression, err = reducer.ExternNamedToAtomType(args[0].Value, args[1].Value, args[2].Value, args[3].Generic_)
	case _ReduceInferredToAtomType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = AtomTypeType
		symbol.TypeExpression, err = reducer.InferredToAtomType(args[0].Value, args[1].Generic_)
	case _ReduceImplicitStructDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:598:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceExplicitEnumDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:599:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceImplicitEnumDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:600:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceTraitDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:601:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceFuncTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:602:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceParseErrorTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:603:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceToParseErrorType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParseErrorTypeType
		symbol.TypeExpression, err = reducer.ToParseErrorType(args[0].ParseError)
	case _ReduceAtomTypeToReturnableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableTypeType
		//line grammar.lr:611:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReducePrefixedTypeToReturnableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableTypeType
		//line grammar.lr:612:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceToPrefixedType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PrefixedTypeType
		symbol.TypeExpression, err = reducer.ToPrefixedType(args[0].Value, args[1].TypeExpression)
	case _ReduceQuestionToPrefixTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixTypeOpType
		//line grammar.lr:617:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceExclaimToPrefixTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixTypeOpType
		//line grammar.lr:618:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToPrefixTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixTypeOpType
		//line grammar.lr:619:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegToPrefixTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixTypeOpType
		//line grammar.lr:620:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceTildeTildeToPrefixTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixTypeOpType
		//line grammar.lr:621:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceReturnableTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		//line grammar.lr:626:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceTraitOpTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		//line grammar.lr:627:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceToTraitOpType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TraitOpTypeType
		symbol.TypeExpression, err = reducer.ToTraitOpType(args[0].TypeExpression, args[1].Value, args[2].TypeExpression)
	case _ReduceMulToTraitOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitOpType
		//line grammar.lr:632:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddToTraitOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitOpType
		//line grammar.lr:633:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToTraitOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitOpType
		//line grammar.lr:634:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDefinitionToTypeDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = TypeDefType
		symbol.SourceDefinition, err = reducer.DefinitionToTypeDef(args[0].Value, args[1].Value, args[2].Generic_, args[3].TypeExpression)
	case _ReduceConstrainedDefToTypeDef:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = TypeDefType
		symbol.SourceDefinition, err = reducer.ConstrainedDefToTypeDef(args[0].Value, args[1].Value, args[2].Generic_, args[3].TypeExpression, args[4].Value, args[5].TypeExpression)
	case _ReduceAliasToTypeDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = TypeDefType
		symbol.SourceDefinition, err = reducer.AliasToTypeDef(args[0].Value, args[1].Value, args[2].Value, args[3].TypeExpression)
	case _ReduceUnconstrainedToGenericParameterDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericParameterDefType
		symbol.Generic_, err = reducer.UnconstrainedToGenericParameterDef(args[0].Value)
	case _ReduceConstrainedToGenericParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = GenericParameterDefType
		symbol.Generic_, err = reducer.ConstrainedToGenericParameterDef(args[0].Value, args[1].TypeExpression)
	case _ReduceGenericParameterDefToGenericParameterDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericParameterDefsType
		symbol.Generic_, err = reducer.GenericParameterDefToGenericParameterDefs(args[0].Generic_)
	case _ReduceAddToGenericParameterDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = GenericParameterDefsType
		symbol.Generic_, err = reducer.AddToGenericParameterDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
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
		symbol.Generic_, err = reducer.GenericToOptionalGenericParameters(args[0].Value, args[1].Generic_, args[2].Value)
	case _ReduceNilToOptionalGenericParameters:
		symbol.SymbolId_ = OptionalGenericParametersType
		symbol.Generic_, err = reducer.NilToOptionalGenericParameters()
	case _ReduceExplicitToFieldDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = FieldDefType
		symbol.Generic_, err = reducer.ExplicitToFieldDef(args[0].Value, args[1].TypeExpression)
	case _ReduceImplicitToFieldDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldDefType
		symbol.Generic_, err = reducer.ImplicitToFieldDef(args[0].TypeExpression)
	case _ReduceUnsafeStatementToFieldDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldDefType
		symbol.Generic_, err = reducer.UnsafeStatementToFieldDef(args[0].Statement)
	case _ReduceFieldDefToImplicitFieldDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ImplicitFieldDefsType
		symbol.Generic_, err = reducer.FieldDefToImplicitFieldDefs(args[0].Generic_)
	case _ReduceAddToImplicitFieldDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitFieldDefsType
		symbol.Generic_, err = reducer.AddToImplicitFieldDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
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
		symbol.TypeExpression, err = reducer.ToImplicitStructDef(args[0].Value, args[1].Generic_, args[2].Value)
	case _ReduceFieldDefToExplicitFieldDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExplicitFieldDefsType
		symbol.Generic_, err = reducer.FieldDefToExplicitFieldDefs(args[0].Generic_)
	case _ReduceImplicitToExplicitFieldDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitFieldDefsType
		symbol.Generic_, err = reducer.ImplicitToExplicitFieldDefs(args[0].Generic_, args[1].Count, args[2].Generic_)
	case _ReduceExplicitToExplicitFieldDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitFieldDefsType
		symbol.Generic_, err = reducer.ExplicitToExplicitFieldDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
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
		symbol.TypeExpression, err = reducer.ToExplicitStructDef(args[0].Value, args[1].Value, args[2].Generic_, args[3].Value)
	case _ReduceFieldDefToEnumValueDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = EnumValueDefType
		symbol.Generic_, err = reducer.FieldDefToEnumValueDef(args[0].Generic_)
	case _ReduceDefaultToEnumValueDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = EnumValueDefType
		symbol.Generic_, err = reducer.DefaultToEnumValueDef(args[0].Generic_, args[1].Value, args[2].Value)
	case _ReducePairToImplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitEnumValueDefsType
		symbol.Generic_, err = reducer.PairToImplicitEnumValueDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceAddToImplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitEnumValueDefsType
		symbol.Generic_, err = reducer.AddToImplicitEnumValueDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceToImplicitEnumDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitEnumDefType
		symbol.TypeExpression, err = reducer.ToImplicitEnumDef(args[0].Value, args[1].Generic_, args[2].Value)
	case _ReduceExplicitPairToExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ExplicitPairToExplicitEnumValueDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceImplicitPairToExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ImplicitPairToExplicitEnumValueDefs(args[0].Generic_, args[1].Count, args[2].Generic_)
	case _ReduceExplicitAddToExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ExplicitAddToExplicitEnumValueDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceImplicitAddToExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ImplicitAddToExplicitEnumValueDefs(args[0].Generic_, args[1].Count, args[2].Generic_)
	case _ReduceToExplicitEnumDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ExplicitEnumDefType
		symbol.TypeExpression, err = reducer.ToExplicitEnumDef(args[0].Value, args[1].Value, args[2].Generic_, args[3].Value)
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
		symbol.Generic_, err = reducer.ImplicitToTraitProperties(args[0].Generic_, args[1].Count, args[2].Generic_)
	case _ReduceExplicitToTraitProperties:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = TraitPropertiesType
		symbol.Generic_, err = reducer.ExplicitToTraitProperties(args[0].Generic_, args[1].Value, args[2].Generic_)
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
		symbol.TypeExpression, err = reducer.ToTraitDef(args[0].Value, args[1].Value, args[2].Generic_, args[3].Value)
	case _ReduceReturnableTypeToReturnType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnTypeType
		symbol.TypeExpression, err = reducer.ReturnableTypeToReturnType(args[0].TypeExpression)
	case _ReduceNilToReturnType:
		symbol.SymbolId_ = ReturnTypeType
		symbol.TypeExpression, err = reducer.NilToReturnType()
	case _ReduceArgToParameterDecl:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.ArgToParameterDecl(args[0].Value, args[1].TypeExpression)
	case _ReduceVarargToParameterDecl:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.VarargToParameterDecl(args[0].Value, args[1].Value, args[2].TypeExpression)
	case _ReduceUnamedToParameterDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.UnamedToParameterDecl(args[0].TypeExpression)
	case _ReduceUnnamedVarargToParameterDecl:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.UnnamedVarargToParameterDecl(args[0].Value, args[1].TypeExpression)
	case _ReduceParameterDeclToParameterDecls:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDeclsType
		symbol.Generic_, err = reducer.ParameterDeclToParameterDecls(args[0].Generic_)
	case _ReduceAddToParameterDecls:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDeclsType
		symbol.Generic_, err = reducer.AddToParameterDecls(args[0].Generic_, args[1].Value, args[2].Generic_)
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
		symbol.TypeExpression, err = reducer.ToFuncType(args[0].Value, args[1].Value, args[2].Generic_, args[3].Value, args[4].TypeExpression)
	case _ReduceToMethodSignature:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = MethodSignatureType
		symbol.Generic_, err = reducer.ToMethodSignature(args[0].Value, args[1].Value, args[2].Value, args[3].Generic_, args[4].Value, args[5].TypeExpression)
	case _ReduceInferredRefArgToParameterDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.InferredRefArgToParameterDef(args[0].Value)
	case _ReduceInferredRefVarargToParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.InferredRefVarargToParameterDef(args[0].Value, args[1].Value)
	case _ReduceArgToParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.ArgToParameterDef(args[0].Value, args[1].TypeExpression)
	case _ReduceVarargToParameterDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.VarargToParameterDef(args[0].Value, args[1].Value, args[2].TypeExpression)
	case _ReduceParameterDefToParameterDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDefsType
		symbol.Generic_, err = reducer.ParameterDefToParameterDefs(args[0].Generic_)
	case _ReduceAddToParameterDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDefsType
		symbol.Generic_, err = reducer.AddToParameterDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceParameterDefsToOptionalParameterDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalParameterDefsType
		symbol.Generic_, err = reducer.ParameterDefsToOptionalParameterDefs(args[0].Generic_)
	case _ReduceNilToOptionalParameterDefs:
		symbol.SymbolId_ = OptionalParameterDefsType
		symbol.Generic_, err = reducer.NilToOptionalParameterDefs()
	case _ReduceFuncDefToNamedFuncDef:
		args := stack[len(stack)-8:]
		stack = stack[:len(stack)-8]
		symbol.SymbolId_ = NamedFuncDefType
		symbol.SourceDefinition, err = reducer.FuncDefToNamedFuncDef(args[0].Value, args[1].Value, args[2].Generic_, args[3].Value, args[4].Generic_, args[5].Value, args[6].TypeExpression, args[7].Generic_)
	case _ReduceMethodDefToNamedFuncDef:
		args := stack[len(stack)-11:]
		stack = stack[:len(stack)-11]
		symbol.SymbolId_ = NamedFuncDefType
		symbol.SourceDefinition, err = reducer.MethodDefToNamedFuncDef(args[0].Value, args[1].Value, args[2].Generic_, args[3].Value, args[4].Value, args[5].Generic_, args[6].Value, args[7].Generic_, args[8].Value, args[9].TypeExpression, args[10].Generic_)
	case _ReduceAliasToNamedFuncDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = NamedFuncDefType
		symbol.SourceDefinition, err = reducer.AliasToNamedFuncDef(args[0].Value, args[1].Value, args[2].Value, args[3].Expression)
	case _ReduceToAnonymousFuncExpr:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = AnonymousFuncExprType
		symbol.Expression, err = reducer.ToAnonymousFuncExpr(args[0].Value, args[1].Value, args[2].Generic_, args[3].Value, args[4].TypeExpression, args[5].Generic_)
	case _ReduceNoSpecToPackageDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PackageDefType
		symbol.SourceDefinition, err = reducer.NoSpecToPackageDef(args[0].Value)
	case _ReduceWithSpecToPackageDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PackageDefType
		symbol.SourceDefinition, err = reducer.WithSpecToPackageDef(args[0].Value, args[1].Generic_)
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
	_GotoState1Action                                                     = &_Action{_ShiftAction, _State1, 0}
	_GotoState2Action                                                     = &_Action{_ShiftAction, _State2, 0}
	_GotoState3Action                                                     = &_Action{_ShiftAction, _State3, 0}
	_GotoState4Action                                                     = &_Action{_ShiftAction, _State4, 0}
	_GotoState5Action                                                     = &_Action{_ShiftAction, _State5, 0}
	_GotoState6Action                                                     = &_Action{_ShiftAction, _State6, 0}
	_GotoState7Action                                                     = &_Action{_ShiftAction, _State7, 0}
	_GotoState8Action                                                     = &_Action{_ShiftAction, _State8, 0}
	_GotoState9Action                                                     = &_Action{_ShiftAction, _State9, 0}
	_GotoState10Action                                                    = &_Action{_ShiftAction, _State10, 0}
	_GotoState11Action                                                    = &_Action{_ShiftAction, _State11, 0}
	_GotoState12Action                                                    = &_Action{_ShiftAction, _State12, 0}
	_GotoState13Action                                                    = &_Action{_ShiftAction, _State13, 0}
	_GotoState14Action                                                    = &_Action{_ShiftAction, _State14, 0}
	_GotoState15Action                                                    = &_Action{_ShiftAction, _State15, 0}
	_GotoState16Action                                                    = &_Action{_ShiftAction, _State16, 0}
	_GotoState17Action                                                    = &_Action{_ShiftAction, _State17, 0}
	_GotoState18Action                                                    = &_Action{_ShiftAction, _State18, 0}
	_GotoState19Action                                                    = &_Action{_ShiftAction, _State19, 0}
	_GotoState20Action                                                    = &_Action{_ShiftAction, _State20, 0}
	_GotoState21Action                                                    = &_Action{_ShiftAction, _State21, 0}
	_GotoState22Action                                                    = &_Action{_ShiftAction, _State22, 0}
	_GotoState23Action                                                    = &_Action{_ShiftAction, _State23, 0}
	_GotoState24Action                                                    = &_Action{_ShiftAction, _State24, 0}
	_GotoState25Action                                                    = &_Action{_ShiftAction, _State25, 0}
	_GotoState26Action                                                    = &_Action{_ShiftAction, _State26, 0}
	_GotoState27Action                                                    = &_Action{_ShiftAction, _State27, 0}
	_GotoState28Action                                                    = &_Action{_ShiftAction, _State28, 0}
	_GotoState29Action                                                    = &_Action{_ShiftAction, _State29, 0}
	_GotoState30Action                                                    = &_Action{_ShiftAction, _State30, 0}
	_GotoState31Action                                                    = &_Action{_ShiftAction, _State31, 0}
	_GotoState32Action                                                    = &_Action{_ShiftAction, _State32, 0}
	_GotoState33Action                                                    = &_Action{_ShiftAction, _State33, 0}
	_GotoState34Action                                                    = &_Action{_ShiftAction, _State34, 0}
	_GotoState35Action                                                    = &_Action{_ShiftAction, _State35, 0}
	_GotoState36Action                                                    = &_Action{_ShiftAction, _State36, 0}
	_GotoState37Action                                                    = &_Action{_ShiftAction, _State37, 0}
	_GotoState38Action                                                    = &_Action{_ShiftAction, _State38, 0}
	_GotoState39Action                                                    = &_Action{_ShiftAction, _State39, 0}
	_GotoState40Action                                                    = &_Action{_ShiftAction, _State40, 0}
	_GotoState41Action                                                    = &_Action{_ShiftAction, _State41, 0}
	_GotoState42Action                                                    = &_Action{_ShiftAction, _State42, 0}
	_GotoState43Action                                                    = &_Action{_ShiftAction, _State43, 0}
	_GotoState44Action                                                    = &_Action{_ShiftAction, _State44, 0}
	_GotoState45Action                                                    = &_Action{_ShiftAction, _State45, 0}
	_GotoState46Action                                                    = &_Action{_ShiftAction, _State46, 0}
	_GotoState47Action                                                    = &_Action{_ShiftAction, _State47, 0}
	_GotoState48Action                                                    = &_Action{_ShiftAction, _State48, 0}
	_GotoState49Action                                                    = &_Action{_ShiftAction, _State49, 0}
	_GotoState50Action                                                    = &_Action{_ShiftAction, _State50, 0}
	_GotoState51Action                                                    = &_Action{_ShiftAction, _State51, 0}
	_GotoState52Action                                                    = &_Action{_ShiftAction, _State52, 0}
	_GotoState53Action                                                    = &_Action{_ShiftAction, _State53, 0}
	_GotoState54Action                                                    = &_Action{_ShiftAction, _State54, 0}
	_GotoState55Action                                                    = &_Action{_ShiftAction, _State55, 0}
	_GotoState56Action                                                    = &_Action{_ShiftAction, _State56, 0}
	_GotoState57Action                                                    = &_Action{_ShiftAction, _State57, 0}
	_GotoState58Action                                                    = &_Action{_ShiftAction, _State58, 0}
	_GotoState59Action                                                    = &_Action{_ShiftAction, _State59, 0}
	_GotoState60Action                                                    = &_Action{_ShiftAction, _State60, 0}
	_GotoState61Action                                                    = &_Action{_ShiftAction, _State61, 0}
	_GotoState62Action                                                    = &_Action{_ShiftAction, _State62, 0}
	_GotoState63Action                                                    = &_Action{_ShiftAction, _State63, 0}
	_GotoState64Action                                                    = &_Action{_ShiftAction, _State64, 0}
	_GotoState65Action                                                    = &_Action{_ShiftAction, _State65, 0}
	_GotoState66Action                                                    = &_Action{_ShiftAction, _State66, 0}
	_GotoState67Action                                                    = &_Action{_ShiftAction, _State67, 0}
	_GotoState68Action                                                    = &_Action{_ShiftAction, _State68, 0}
	_GotoState69Action                                                    = &_Action{_ShiftAction, _State69, 0}
	_GotoState70Action                                                    = &_Action{_ShiftAction, _State70, 0}
	_GotoState71Action                                                    = &_Action{_ShiftAction, _State71, 0}
	_GotoState72Action                                                    = &_Action{_ShiftAction, _State72, 0}
	_GotoState73Action                                                    = &_Action{_ShiftAction, _State73, 0}
	_GotoState74Action                                                    = &_Action{_ShiftAction, _State74, 0}
	_GotoState75Action                                                    = &_Action{_ShiftAction, _State75, 0}
	_GotoState76Action                                                    = &_Action{_ShiftAction, _State76, 0}
	_GotoState77Action                                                    = &_Action{_ShiftAction, _State77, 0}
	_GotoState78Action                                                    = &_Action{_ShiftAction, _State78, 0}
	_GotoState79Action                                                    = &_Action{_ShiftAction, _State79, 0}
	_GotoState80Action                                                    = &_Action{_ShiftAction, _State80, 0}
	_GotoState81Action                                                    = &_Action{_ShiftAction, _State81, 0}
	_GotoState82Action                                                    = &_Action{_ShiftAction, _State82, 0}
	_GotoState83Action                                                    = &_Action{_ShiftAction, _State83, 0}
	_GotoState84Action                                                    = &_Action{_ShiftAction, _State84, 0}
	_GotoState85Action                                                    = &_Action{_ShiftAction, _State85, 0}
	_GotoState86Action                                                    = &_Action{_ShiftAction, _State86, 0}
	_GotoState87Action                                                    = &_Action{_ShiftAction, _State87, 0}
	_GotoState88Action                                                    = &_Action{_ShiftAction, _State88, 0}
	_GotoState89Action                                                    = &_Action{_ShiftAction, _State89, 0}
	_GotoState90Action                                                    = &_Action{_ShiftAction, _State90, 0}
	_GotoState91Action                                                    = &_Action{_ShiftAction, _State91, 0}
	_GotoState92Action                                                    = &_Action{_ShiftAction, _State92, 0}
	_GotoState93Action                                                    = &_Action{_ShiftAction, _State93, 0}
	_GotoState94Action                                                    = &_Action{_ShiftAction, _State94, 0}
	_GotoState95Action                                                    = &_Action{_ShiftAction, _State95, 0}
	_GotoState96Action                                                    = &_Action{_ShiftAction, _State96, 0}
	_GotoState97Action                                                    = &_Action{_ShiftAction, _State97, 0}
	_GotoState98Action                                                    = &_Action{_ShiftAction, _State98, 0}
	_GotoState99Action                                                    = &_Action{_ShiftAction, _State99, 0}
	_GotoState100Action                                                   = &_Action{_ShiftAction, _State100, 0}
	_GotoState101Action                                                   = &_Action{_ShiftAction, _State101, 0}
	_GotoState102Action                                                   = &_Action{_ShiftAction, _State102, 0}
	_GotoState103Action                                                   = &_Action{_ShiftAction, _State103, 0}
	_GotoState104Action                                                   = &_Action{_ShiftAction, _State104, 0}
	_GotoState105Action                                                   = &_Action{_ShiftAction, _State105, 0}
	_GotoState106Action                                                   = &_Action{_ShiftAction, _State106, 0}
	_GotoState107Action                                                   = &_Action{_ShiftAction, _State107, 0}
	_GotoState108Action                                                   = &_Action{_ShiftAction, _State108, 0}
	_GotoState109Action                                                   = &_Action{_ShiftAction, _State109, 0}
	_GotoState110Action                                                   = &_Action{_ShiftAction, _State110, 0}
	_GotoState111Action                                                   = &_Action{_ShiftAction, _State111, 0}
	_GotoState112Action                                                   = &_Action{_ShiftAction, _State112, 0}
	_GotoState113Action                                                   = &_Action{_ShiftAction, _State113, 0}
	_GotoState114Action                                                   = &_Action{_ShiftAction, _State114, 0}
	_GotoState115Action                                                   = &_Action{_ShiftAction, _State115, 0}
	_GotoState116Action                                                   = &_Action{_ShiftAction, _State116, 0}
	_GotoState117Action                                                   = &_Action{_ShiftAction, _State117, 0}
	_GotoState118Action                                                   = &_Action{_ShiftAction, _State118, 0}
	_GotoState119Action                                                   = &_Action{_ShiftAction, _State119, 0}
	_GotoState120Action                                                   = &_Action{_ShiftAction, _State120, 0}
	_GotoState121Action                                                   = &_Action{_ShiftAction, _State121, 0}
	_GotoState122Action                                                   = &_Action{_ShiftAction, _State122, 0}
	_GotoState123Action                                                   = &_Action{_ShiftAction, _State123, 0}
	_GotoState124Action                                                   = &_Action{_ShiftAction, _State124, 0}
	_GotoState125Action                                                   = &_Action{_ShiftAction, _State125, 0}
	_GotoState126Action                                                   = &_Action{_ShiftAction, _State126, 0}
	_GotoState127Action                                                   = &_Action{_ShiftAction, _State127, 0}
	_GotoState128Action                                                   = &_Action{_ShiftAction, _State128, 0}
	_GotoState129Action                                                   = &_Action{_ShiftAction, _State129, 0}
	_GotoState130Action                                                   = &_Action{_ShiftAction, _State130, 0}
	_GotoState131Action                                                   = &_Action{_ShiftAction, _State131, 0}
	_GotoState132Action                                                   = &_Action{_ShiftAction, _State132, 0}
	_GotoState133Action                                                   = &_Action{_ShiftAction, _State133, 0}
	_GotoState134Action                                                   = &_Action{_ShiftAction, _State134, 0}
	_GotoState135Action                                                   = &_Action{_ShiftAction, _State135, 0}
	_GotoState136Action                                                   = &_Action{_ShiftAction, _State136, 0}
	_GotoState137Action                                                   = &_Action{_ShiftAction, _State137, 0}
	_GotoState138Action                                                   = &_Action{_ShiftAction, _State138, 0}
	_GotoState139Action                                                   = &_Action{_ShiftAction, _State139, 0}
	_GotoState140Action                                                   = &_Action{_ShiftAction, _State140, 0}
	_GotoState141Action                                                   = &_Action{_ShiftAction, _State141, 0}
	_GotoState142Action                                                   = &_Action{_ShiftAction, _State142, 0}
	_GotoState143Action                                                   = &_Action{_ShiftAction, _State143, 0}
	_GotoState144Action                                                   = &_Action{_ShiftAction, _State144, 0}
	_GotoState145Action                                                   = &_Action{_ShiftAction, _State145, 0}
	_GotoState146Action                                                   = &_Action{_ShiftAction, _State146, 0}
	_GotoState147Action                                                   = &_Action{_ShiftAction, _State147, 0}
	_GotoState148Action                                                   = &_Action{_ShiftAction, _State148, 0}
	_GotoState149Action                                                   = &_Action{_ShiftAction, _State149, 0}
	_GotoState150Action                                                   = &_Action{_ShiftAction, _State150, 0}
	_GotoState151Action                                                   = &_Action{_ShiftAction, _State151, 0}
	_GotoState152Action                                                   = &_Action{_ShiftAction, _State152, 0}
	_GotoState153Action                                                   = &_Action{_ShiftAction, _State153, 0}
	_GotoState154Action                                                   = &_Action{_ShiftAction, _State154, 0}
	_GotoState155Action                                                   = &_Action{_ShiftAction, _State155, 0}
	_GotoState156Action                                                   = &_Action{_ShiftAction, _State156, 0}
	_GotoState157Action                                                   = &_Action{_ShiftAction, _State157, 0}
	_GotoState158Action                                                   = &_Action{_ShiftAction, _State158, 0}
	_GotoState159Action                                                   = &_Action{_ShiftAction, _State159, 0}
	_GotoState160Action                                                   = &_Action{_ShiftAction, _State160, 0}
	_GotoState161Action                                                   = &_Action{_ShiftAction, _State161, 0}
	_GotoState162Action                                                   = &_Action{_ShiftAction, _State162, 0}
	_GotoState163Action                                                   = &_Action{_ShiftAction, _State163, 0}
	_GotoState164Action                                                   = &_Action{_ShiftAction, _State164, 0}
	_GotoState165Action                                                   = &_Action{_ShiftAction, _State165, 0}
	_GotoState166Action                                                   = &_Action{_ShiftAction, _State166, 0}
	_GotoState167Action                                                   = &_Action{_ShiftAction, _State167, 0}
	_GotoState168Action                                                   = &_Action{_ShiftAction, _State168, 0}
	_GotoState169Action                                                   = &_Action{_ShiftAction, _State169, 0}
	_GotoState170Action                                                   = &_Action{_ShiftAction, _State170, 0}
	_GotoState171Action                                                   = &_Action{_ShiftAction, _State171, 0}
	_GotoState172Action                                                   = &_Action{_ShiftAction, _State172, 0}
	_GotoState173Action                                                   = &_Action{_ShiftAction, _State173, 0}
	_GotoState174Action                                                   = &_Action{_ShiftAction, _State174, 0}
	_GotoState175Action                                                   = &_Action{_ShiftAction, _State175, 0}
	_GotoState176Action                                                   = &_Action{_ShiftAction, _State176, 0}
	_GotoState177Action                                                   = &_Action{_ShiftAction, _State177, 0}
	_GotoState178Action                                                   = &_Action{_ShiftAction, _State178, 0}
	_GotoState179Action                                                   = &_Action{_ShiftAction, _State179, 0}
	_GotoState180Action                                                   = &_Action{_ShiftAction, _State180, 0}
	_GotoState181Action                                                   = &_Action{_ShiftAction, _State181, 0}
	_GotoState182Action                                                   = &_Action{_ShiftAction, _State182, 0}
	_GotoState183Action                                                   = &_Action{_ShiftAction, _State183, 0}
	_GotoState184Action                                                   = &_Action{_ShiftAction, _State184, 0}
	_GotoState185Action                                                   = &_Action{_ShiftAction, _State185, 0}
	_GotoState186Action                                                   = &_Action{_ShiftAction, _State186, 0}
	_GotoState187Action                                                   = &_Action{_ShiftAction, _State187, 0}
	_GotoState188Action                                                   = &_Action{_ShiftAction, _State188, 0}
	_GotoState189Action                                                   = &_Action{_ShiftAction, _State189, 0}
	_GotoState190Action                                                   = &_Action{_ShiftAction, _State190, 0}
	_GotoState191Action                                                   = &_Action{_ShiftAction, _State191, 0}
	_GotoState192Action                                                   = &_Action{_ShiftAction, _State192, 0}
	_GotoState193Action                                                   = &_Action{_ShiftAction, _State193, 0}
	_GotoState194Action                                                   = &_Action{_ShiftAction, _State194, 0}
	_GotoState195Action                                                   = &_Action{_ShiftAction, _State195, 0}
	_GotoState196Action                                                   = &_Action{_ShiftAction, _State196, 0}
	_GotoState197Action                                                   = &_Action{_ShiftAction, _State197, 0}
	_GotoState198Action                                                   = &_Action{_ShiftAction, _State198, 0}
	_GotoState199Action                                                   = &_Action{_ShiftAction, _State199, 0}
	_GotoState200Action                                                   = &_Action{_ShiftAction, _State200, 0}
	_GotoState201Action                                                   = &_Action{_ShiftAction, _State201, 0}
	_GotoState202Action                                                   = &_Action{_ShiftAction, _State202, 0}
	_GotoState203Action                                                   = &_Action{_ShiftAction, _State203, 0}
	_GotoState204Action                                                   = &_Action{_ShiftAction, _State204, 0}
	_GotoState205Action                                                   = &_Action{_ShiftAction, _State205, 0}
	_GotoState206Action                                                   = &_Action{_ShiftAction, _State206, 0}
	_GotoState207Action                                                   = &_Action{_ShiftAction, _State207, 0}
	_GotoState208Action                                                   = &_Action{_ShiftAction, _State208, 0}
	_GotoState209Action                                                   = &_Action{_ShiftAction, _State209, 0}
	_GotoState210Action                                                   = &_Action{_ShiftAction, _State210, 0}
	_GotoState211Action                                                   = &_Action{_ShiftAction, _State211, 0}
	_GotoState212Action                                                   = &_Action{_ShiftAction, _State212, 0}
	_GotoState213Action                                                   = &_Action{_ShiftAction, _State213, 0}
	_GotoState214Action                                                   = &_Action{_ShiftAction, _State214, 0}
	_GotoState215Action                                                   = &_Action{_ShiftAction, _State215, 0}
	_GotoState216Action                                                   = &_Action{_ShiftAction, _State216, 0}
	_GotoState217Action                                                   = &_Action{_ShiftAction, _State217, 0}
	_GotoState218Action                                                   = &_Action{_ShiftAction, _State218, 0}
	_GotoState219Action                                                   = &_Action{_ShiftAction, _State219, 0}
	_GotoState220Action                                                   = &_Action{_ShiftAction, _State220, 0}
	_GotoState221Action                                                   = &_Action{_ShiftAction, _State221, 0}
	_GotoState222Action                                                   = &_Action{_ShiftAction, _State222, 0}
	_GotoState223Action                                                   = &_Action{_ShiftAction, _State223, 0}
	_GotoState224Action                                                   = &_Action{_ShiftAction, _State224, 0}
	_GotoState225Action                                                   = &_Action{_ShiftAction, _State225, 0}
	_GotoState226Action                                                   = &_Action{_ShiftAction, _State226, 0}
	_GotoState227Action                                                   = &_Action{_ShiftAction, _State227, 0}
	_GotoState228Action                                                   = &_Action{_ShiftAction, _State228, 0}
	_GotoState229Action                                                   = &_Action{_ShiftAction, _State229, 0}
	_GotoState230Action                                                   = &_Action{_ShiftAction, _State230, 0}
	_GotoState231Action                                                   = &_Action{_ShiftAction, _State231, 0}
	_GotoState232Action                                                   = &_Action{_ShiftAction, _State232, 0}
	_GotoState233Action                                                   = &_Action{_ShiftAction, _State233, 0}
	_GotoState234Action                                                   = &_Action{_ShiftAction, _State234, 0}
	_GotoState235Action                                                   = &_Action{_ShiftAction, _State235, 0}
	_GotoState236Action                                                   = &_Action{_ShiftAction, _State236, 0}
	_GotoState237Action                                                   = &_Action{_ShiftAction, _State237, 0}
	_GotoState238Action                                                   = &_Action{_ShiftAction, _State238, 0}
	_GotoState239Action                                                   = &_Action{_ShiftAction, _State239, 0}
	_GotoState240Action                                                   = &_Action{_ShiftAction, _State240, 0}
	_GotoState241Action                                                   = &_Action{_ShiftAction, _State241, 0}
	_GotoState242Action                                                   = &_Action{_ShiftAction, _State242, 0}
	_GotoState243Action                                                   = &_Action{_ShiftAction, _State243, 0}
	_GotoState244Action                                                   = &_Action{_ShiftAction, _State244, 0}
	_GotoState245Action                                                   = &_Action{_ShiftAction, _State245, 0}
	_GotoState246Action                                                   = &_Action{_ShiftAction, _State246, 0}
	_GotoState247Action                                                   = &_Action{_ShiftAction, _State247, 0}
	_GotoState248Action                                                   = &_Action{_ShiftAction, _State248, 0}
	_GotoState249Action                                                   = &_Action{_ShiftAction, _State249, 0}
	_GotoState250Action                                                   = &_Action{_ShiftAction, _State250, 0}
	_GotoState251Action                                                   = &_Action{_ShiftAction, _State251, 0}
	_GotoState252Action                                                   = &_Action{_ShiftAction, _State252, 0}
	_GotoState253Action                                                   = &_Action{_ShiftAction, _State253, 0}
	_GotoState254Action                                                   = &_Action{_ShiftAction, _State254, 0}
	_GotoState255Action                                                   = &_Action{_ShiftAction, _State255, 0}
	_GotoState256Action                                                   = &_Action{_ShiftAction, _State256, 0}
	_GotoState257Action                                                   = &_Action{_ShiftAction, _State257, 0}
	_GotoState258Action                                                   = &_Action{_ShiftAction, _State258, 0}
	_GotoState259Action                                                   = &_Action{_ShiftAction, _State259, 0}
	_GotoState260Action                                                   = &_Action{_ShiftAction, _State260, 0}
	_GotoState261Action                                                   = &_Action{_ShiftAction, _State261, 0}
	_GotoState262Action                                                   = &_Action{_ShiftAction, _State262, 0}
	_GotoState263Action                                                   = &_Action{_ShiftAction, _State263, 0}
	_GotoState264Action                                                   = &_Action{_ShiftAction, _State264, 0}
	_GotoState265Action                                                   = &_Action{_ShiftAction, _State265, 0}
	_GotoState266Action                                                   = &_Action{_ShiftAction, _State266, 0}
	_GotoState267Action                                                   = &_Action{_ShiftAction, _State267, 0}
	_GotoState268Action                                                   = &_Action{_ShiftAction, _State268, 0}
	_GotoState269Action                                                   = &_Action{_ShiftAction, _State269, 0}
	_GotoState270Action                                                   = &_Action{_ShiftAction, _State270, 0}
	_GotoState271Action                                                   = &_Action{_ShiftAction, _State271, 0}
	_GotoState272Action                                                   = &_Action{_ShiftAction, _State272, 0}
	_GotoState273Action                                                   = &_Action{_ShiftAction, _State273, 0}
	_GotoState274Action                                                   = &_Action{_ShiftAction, _State274, 0}
	_GotoState275Action                                                   = &_Action{_ShiftAction, _State275, 0}
	_GotoState276Action                                                   = &_Action{_ShiftAction, _State276, 0}
	_GotoState277Action                                                   = &_Action{_ShiftAction, _State277, 0}
	_GotoState278Action                                                   = &_Action{_ShiftAction, _State278, 0}
	_GotoState279Action                                                   = &_Action{_ShiftAction, _State279, 0}
	_GotoState280Action                                                   = &_Action{_ShiftAction, _State280, 0}
	_GotoState281Action                                                   = &_Action{_ShiftAction, _State281, 0}
	_GotoState282Action                                                   = &_Action{_ShiftAction, _State282, 0}
	_GotoState283Action                                                   = &_Action{_ShiftAction, _State283, 0}
	_GotoState284Action                                                   = &_Action{_ShiftAction, _State284, 0}
	_GotoState285Action                                                   = &_Action{_ShiftAction, _State285, 0}
	_GotoState286Action                                                   = &_Action{_ShiftAction, _State286, 0}
	_GotoState287Action                                                   = &_Action{_ShiftAction, _State287, 0}
	_GotoState288Action                                                   = &_Action{_ShiftAction, _State288, 0}
	_GotoState289Action                                                   = &_Action{_ShiftAction, _State289, 0}
	_GotoState290Action                                                   = &_Action{_ShiftAction, _State290, 0}
	_GotoState291Action                                                   = &_Action{_ShiftAction, _State291, 0}
	_GotoState292Action                                                   = &_Action{_ShiftAction, _State292, 0}
	_GotoState293Action                                                   = &_Action{_ShiftAction, _State293, 0}
	_GotoState294Action                                                   = &_Action{_ShiftAction, _State294, 0}
	_GotoState295Action                                                   = &_Action{_ShiftAction, _State295, 0}
	_GotoState296Action                                                   = &_Action{_ShiftAction, _State296, 0}
	_GotoState297Action                                                   = &_Action{_ShiftAction, _State297, 0}
	_GotoState298Action                                                   = &_Action{_ShiftAction, _State298, 0}
	_GotoState299Action                                                   = &_Action{_ShiftAction, _State299, 0}
	_GotoState300Action                                                   = &_Action{_ShiftAction, _State300, 0}
	_GotoState301Action                                                   = &_Action{_ShiftAction, _State301, 0}
	_GotoState302Action                                                   = &_Action{_ShiftAction, _State302, 0}
	_GotoState303Action                                                   = &_Action{_ShiftAction, _State303, 0}
	_GotoState304Action                                                   = &_Action{_ShiftAction, _State304, 0}
	_GotoState305Action                                                   = &_Action{_ShiftAction, _State305, 0}
	_GotoState306Action                                                   = &_Action{_ShiftAction, _State306, 0}
	_GotoState307Action                                                   = &_Action{_ShiftAction, _State307, 0}
	_GotoState308Action                                                   = &_Action{_ShiftAction, _State308, 0}
	_GotoState309Action                                                   = &_Action{_ShiftAction, _State309, 0}
	_GotoState310Action                                                   = &_Action{_ShiftAction, _State310, 0}
	_GotoState311Action                                                   = &_Action{_ShiftAction, _State311, 0}
	_GotoState312Action                                                   = &_Action{_ShiftAction, _State312, 0}
	_GotoState313Action                                                   = &_Action{_ShiftAction, _State313, 0}
	_GotoState314Action                                                   = &_Action{_ShiftAction, _State314, 0}
	_GotoState315Action                                                   = &_Action{_ShiftAction, _State315, 0}
	_GotoState316Action                                                   = &_Action{_ShiftAction, _State316, 0}
	_GotoState317Action                                                   = &_Action{_ShiftAction, _State317, 0}
	_GotoState318Action                                                   = &_Action{_ShiftAction, _State318, 0}
	_GotoState319Action                                                   = &_Action{_ShiftAction, _State319, 0}
	_GotoState320Action                                                   = &_Action{_ShiftAction, _State320, 0}
	_GotoState321Action                                                   = &_Action{_ShiftAction, _State321, 0}
	_GotoState322Action                                                   = &_Action{_ShiftAction, _State322, 0}
	_GotoState323Action                                                   = &_Action{_ShiftAction, _State323, 0}
	_GotoState324Action                                                   = &_Action{_ShiftAction, _State324, 0}
	_GotoState325Action                                                   = &_Action{_ShiftAction, _State325, 0}
	_GotoState326Action                                                   = &_Action{_ShiftAction, _State326, 0}
	_GotoState327Action                                                   = &_Action{_ShiftAction, _State327, 0}
	_GotoState328Action                                                   = &_Action{_ShiftAction, _State328, 0}
	_GotoState329Action                                                   = &_Action{_ShiftAction, _State329, 0}
	_GotoState330Action                                                   = &_Action{_ShiftAction, _State330, 0}
	_GotoState331Action                                                   = &_Action{_ShiftAction, _State331, 0}
	_GotoState332Action                                                   = &_Action{_ShiftAction, _State332, 0}
	_GotoState333Action                                                   = &_Action{_ShiftAction, _State333, 0}
	_GotoState334Action                                                   = &_Action{_ShiftAction, _State334, 0}
	_GotoState335Action                                                   = &_Action{_ShiftAction, _State335, 0}
	_GotoState336Action                                                   = &_Action{_ShiftAction, _State336, 0}
	_GotoState337Action                                                   = &_Action{_ShiftAction, _State337, 0}
	_GotoState338Action                                                   = &_Action{_ShiftAction, _State338, 0}
	_GotoState339Action                                                   = &_Action{_ShiftAction, _State339, 0}
	_GotoState340Action                                                   = &_Action{_ShiftAction, _State340, 0}
	_GotoState341Action                                                   = &_Action{_ShiftAction, _State341, 0}
	_GotoState342Action                                                   = &_Action{_ShiftAction, _State342, 0}
	_GotoState343Action                                                   = &_Action{_ShiftAction, _State343, 0}
	_GotoState344Action                                                   = &_Action{_ShiftAction, _State344, 0}
	_GotoState345Action                                                   = &_Action{_ShiftAction, _State345, 0}
	_GotoState346Action                                                   = &_Action{_ShiftAction, _State346, 0}
	_GotoState347Action                                                   = &_Action{_ShiftAction, _State347, 0}
	_GotoState348Action                                                   = &_Action{_ShiftAction, _State348, 0}
	_GotoState349Action                                                   = &_Action{_ShiftAction, _State349, 0}
	_GotoState350Action                                                   = &_Action{_ShiftAction, _State350, 0}
	_GotoState351Action                                                   = &_Action{_ShiftAction, _State351, 0}
	_GotoState352Action                                                   = &_Action{_ShiftAction, _State352, 0}
	_GotoState353Action                                                   = &_Action{_ShiftAction, _State353, 0}
	_GotoState354Action                                                   = &_Action{_ShiftAction, _State354, 0}
	_GotoState355Action                                                   = &_Action{_ShiftAction, _State355, 0}
	_GotoState356Action                                                   = &_Action{_ShiftAction, _State356, 0}
	_GotoState357Action                                                   = &_Action{_ShiftAction, _State357, 0}
	_GotoState358Action                                                   = &_Action{_ShiftAction, _State358, 0}
	_GotoState359Action                                                   = &_Action{_ShiftAction, _State359, 0}
	_GotoState360Action                                                   = &_Action{_ShiftAction, _State360, 0}
	_GotoState361Action                                                   = &_Action{_ShiftAction, _State361, 0}
	_GotoState362Action                                                   = &_Action{_ShiftAction, _State362, 0}
	_GotoState363Action                                                   = &_Action{_ShiftAction, _State363, 0}
	_GotoState364Action                                                   = &_Action{_ShiftAction, _State364, 0}
	_GotoState365Action                                                   = &_Action{_ShiftAction, _State365, 0}
	_GotoState366Action                                                   = &_Action{_ShiftAction, _State366, 0}
	_GotoState367Action                                                   = &_Action{_ShiftAction, _State367, 0}
	_GotoState368Action                                                   = &_Action{_ShiftAction, _State368, 0}
	_GotoState369Action                                                   = &_Action{_ShiftAction, _State369, 0}
	_GotoState370Action                                                   = &_Action{_ShiftAction, _State370, 0}
	_GotoState371Action                                                   = &_Action{_ShiftAction, _State371, 0}
	_GotoState372Action                                                   = &_Action{_ShiftAction, _State372, 0}
	_GotoState373Action                                                   = &_Action{_ShiftAction, _State373, 0}
	_GotoState374Action                                                   = &_Action{_ShiftAction, _State374, 0}
	_GotoState375Action                                                   = &_Action{_ShiftAction, _State375, 0}
	_GotoState376Action                                                   = &_Action{_ShiftAction, _State376, 0}
	_GotoState377Action                                                   = &_Action{_ShiftAction, _State377, 0}
	_GotoState378Action                                                   = &_Action{_ShiftAction, _State378, 0}
	_GotoState379Action                                                   = &_Action{_ShiftAction, _State379, 0}
	_GotoState380Action                                                   = &_Action{_ShiftAction, _State380, 0}
	_GotoState381Action                                                   = &_Action{_ShiftAction, _State381, 0}
	_GotoState382Action                                                   = &_Action{_ShiftAction, _State382, 0}
	_GotoState383Action                                                   = &_Action{_ShiftAction, _State383, 0}
	_GotoState384Action                                                   = &_Action{_ShiftAction, _State384, 0}
	_GotoState385Action                                                   = &_Action{_ShiftAction, _State385, 0}
	_GotoState386Action                                                   = &_Action{_ShiftAction, _State386, 0}
	_GotoState387Action                                                   = &_Action{_ShiftAction, _State387, 0}
	_GotoState388Action                                                   = &_Action{_ShiftAction, _State388, 0}
	_GotoState389Action                                                   = &_Action{_ShiftAction, _State389, 0}
	_GotoState390Action                                                   = &_Action{_ShiftAction, _State390, 0}
	_GotoState391Action                                                   = &_Action{_ShiftAction, _State391, 0}
	_GotoState392Action                                                   = &_Action{_ShiftAction, _State392, 0}
	_GotoState393Action                                                   = &_Action{_ShiftAction, _State393, 0}
	_GotoState394Action                                                   = &_Action{_ShiftAction, _State394, 0}
	_GotoState395Action                                                   = &_Action{_ShiftAction, _State395, 0}
	_GotoState396Action                                                   = &_Action{_ShiftAction, _State396, 0}
	_GotoState397Action                                                   = &_Action{_ShiftAction, _State397, 0}
	_GotoState398Action                                                   = &_Action{_ShiftAction, _State398, 0}
	_GotoState399Action                                                   = &_Action{_ShiftAction, _State399, 0}
	_GotoState400Action                                                   = &_Action{_ShiftAction, _State400, 0}
	_GotoState401Action                                                   = &_Action{_ShiftAction, _State401, 0}
	_GotoState402Action                                                   = &_Action{_ShiftAction, _State402, 0}
	_GotoState403Action                                                   = &_Action{_ShiftAction, _State403, 0}
	_GotoState404Action                                                   = &_Action{_ShiftAction, _State404, 0}
	_GotoState405Action                                                   = &_Action{_ShiftAction, _State405, 0}
	_GotoState406Action                                                   = &_Action{_ShiftAction, _State406, 0}
	_GotoState407Action                                                   = &_Action{_ShiftAction, _State407, 0}
	_GotoState408Action                                                   = &_Action{_ShiftAction, _State408, 0}
	_GotoState409Action                                                   = &_Action{_ShiftAction, _State409, 0}
	_GotoState410Action                                                   = &_Action{_ShiftAction, _State410, 0}
	_GotoState411Action                                                   = &_Action{_ShiftAction, _State411, 0}
	_GotoState412Action                                                   = &_Action{_ShiftAction, _State412, 0}
	_GotoState413Action                                                   = &_Action{_ShiftAction, _State413, 0}
	_GotoState414Action                                                   = &_Action{_ShiftAction, _State414, 0}
	_GotoState415Action                                                   = &_Action{_ShiftAction, _State415, 0}
	_GotoState416Action                                                   = &_Action{_ShiftAction, _State416, 0}
	_GotoState417Action                                                   = &_Action{_ShiftAction, _State417, 0}
	_GotoState418Action                                                   = &_Action{_ShiftAction, _State418, 0}
	_GotoState419Action                                                   = &_Action{_ShiftAction, _State419, 0}
	_GotoState420Action                                                   = &_Action{_ShiftAction, _State420, 0}
	_GotoState421Action                                                   = &_Action{_ShiftAction, _State421, 0}
	_GotoState422Action                                                   = &_Action{_ShiftAction, _State422, 0}
	_GotoState423Action                                                   = &_Action{_ShiftAction, _State423, 0}
	_GotoState424Action                                                   = &_Action{_ShiftAction, _State424, 0}
	_GotoState425Action                                                   = &_Action{_ShiftAction, _State425, 0}
	_GotoState426Action                                                   = &_Action{_ShiftAction, _State426, 0}
	_GotoState427Action                                                   = &_Action{_ShiftAction, _State427, 0}
	_GotoState428Action                                                   = &_Action{_ShiftAction, _State428, 0}
	_GotoState429Action                                                   = &_Action{_ShiftAction, _State429, 0}
	_GotoState430Action                                                   = &_Action{_ShiftAction, _State430, 0}
	_GotoState431Action                                                   = &_Action{_ShiftAction, _State431, 0}
	_GotoState432Action                                                   = &_Action{_ShiftAction, _State432, 0}
	_GotoState433Action                                                   = &_Action{_ShiftAction, _State433, 0}
	_GotoState434Action                                                   = &_Action{_ShiftAction, _State434, 0}
	_GotoState435Action                                                   = &_Action{_ShiftAction, _State435, 0}
	_GotoState436Action                                                   = &_Action{_ShiftAction, _State436, 0}
	_GotoState437Action                                                   = &_Action{_ShiftAction, _State437, 0}
	_GotoState438Action                                                   = &_Action{_ShiftAction, _State438, 0}
	_GotoState439Action                                                   = &_Action{_ShiftAction, _State439, 0}
	_GotoState440Action                                                   = &_Action{_ShiftAction, _State440, 0}
	_GotoState441Action                                                   = &_Action{_ShiftAction, _State441, 0}
	_GotoState442Action                                                   = &_Action{_ShiftAction, _State442, 0}
	_GotoState443Action                                                   = &_Action{_ShiftAction, _State443, 0}
	_GotoState444Action                                                   = &_Action{_ShiftAction, _State444, 0}
	_GotoState445Action                                                   = &_Action{_ShiftAction, _State445, 0}
	_GotoState446Action                                                   = &_Action{_ShiftAction, _State446, 0}
	_GotoState447Action                                                   = &_Action{_ShiftAction, _State447, 0}
	_GotoState448Action                                                   = &_Action{_ShiftAction, _State448, 0}
	_GotoState449Action                                                   = &_Action{_ShiftAction, _State449, 0}
	_GotoState450Action                                                   = &_Action{_ShiftAction, _State450, 0}
	_GotoState451Action                                                   = &_Action{_ShiftAction, _State451, 0}
	_GotoState452Action                                                   = &_Action{_ShiftAction, _State452, 0}
	_GotoState453Action                                                   = &_Action{_ShiftAction, _State453, 0}
	_GotoState454Action                                                   = &_Action{_ShiftAction, _State454, 0}
	_GotoState455Action                                                   = &_Action{_ShiftAction, _State455, 0}
	_GotoState456Action                                                   = &_Action{_ShiftAction, _State456, 0}
	_GotoState457Action                                                   = &_Action{_ShiftAction, _State457, 0}
	_GotoState458Action                                                   = &_Action{_ShiftAction, _State458, 0}
	_GotoState459Action                                                   = &_Action{_ShiftAction, _State459, 0}
	_GotoState460Action                                                   = &_Action{_ShiftAction, _State460, 0}
	_GotoState461Action                                                   = &_Action{_ShiftAction, _State461, 0}
	_GotoState462Action                                                   = &_Action{_ShiftAction, _State462, 0}
	_GotoState463Action                                                   = &_Action{_ShiftAction, _State463, 0}
	_GotoState464Action                                                   = &_Action{_ShiftAction, _State464, 0}
	_GotoState465Action                                                   = &_Action{_ShiftAction, _State465, 0}
	_GotoState466Action                                                   = &_Action{_ShiftAction, _State466, 0}
	_GotoState467Action                                                   = &_Action{_ShiftAction, _State467, 0}
	_GotoState468Action                                                   = &_Action{_ShiftAction, _State468, 0}
	_ReduceToSourceAction                                                 = &_Action{_ReduceAction, 0, _ReduceToSource}
	_ReduceNewlinesToOptionalDefinitionsAction                            = &_Action{_ReduceAction, 0, _ReduceNewlinesToOptionalDefinitions}
	_ReduceDefinitionsToOptionalDefinitionsAction                         = &_Action{_ReduceAction, 0, _ReduceDefinitionsToOptionalDefinitions}
	_ReduceNilToOptionalDefinitionsAction                                 = &_Action{_ReduceAction, 0, _ReduceNilToOptionalDefinitions}
	_ReduceNewlinesToOptionalNewlinesAction                               = &_Action{_ReduceAction, 0, _ReduceNewlinesToOptionalNewlines}
	_ReduceNilToOptionalNewlinesAction                                    = &_Action{_ReduceAction, 0, _ReduceNilToOptionalNewlines}
	_ReduceDefinitionToDefinitionsAction                                  = &_Action{_ReduceAction, 0, _ReduceDefinitionToDefinitions}
	_ReduceAddToDefinitionsAction                                         = &_Action{_ReduceAction, 0, _ReduceAddToDefinitions}
	_ReducePackageDefToDefinitionAction                                   = &_Action{_ReduceAction, 0, _ReducePackageDefToDefinition}
	_ReduceTypeDefToDefinitionAction                                      = &_Action{_ReduceAction, 0, _ReduceTypeDefToDefinition}
	_ReduceNamedFuncDefToDefinitionAction                                 = &_Action{_ReduceAction, 0, _ReduceNamedFuncDefToDefinition}
	_ReduceGlobalVarDefToDefinitionAction                                 = &_Action{_ReduceAction, 0, _ReduceGlobalVarDefToDefinition}
	_ReduceGlobalVarAssignmentToDefinitionAction                          = &_Action{_ReduceAction, 0, _ReduceGlobalVarAssignmentToDefinition}
	_ReduceStatementBlockToDefinitionAction                               = &_Action{_ReduceAction, 0, _ReduceStatementBlockToDefinition}
	_ReduceCommentGroupsToDefinitionAction                                = &_Action{_ReduceAction, 0, _ReduceCommentGroupsToDefinition}
	_ReduceToStatementBlockAction                                         = &_Action{_ReduceAction, 0, _ReduceToStatementBlock}
	_ReduceEmptyListToStatementsAction                                    = &_Action{_ReduceAction, 0, _ReduceEmptyListToStatements}
	_ReduceAddToStatementsAction                                          = &_Action{_ReduceAction, 0, _ReduceAddToStatements}
	_ReduceImplicitToStatementAction                                      = &_Action{_ReduceAction, 0, _ReduceImplicitToStatement}
	_ReduceExplicitToStatementAction                                      = &_Action{_ReduceAction, 0, _ReduceExplicitToStatement}
	_ReduceUnsafeStatementToSimpleStatementBodyAction                     = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToSimpleStatementBody}
	_ReduceExpressionOrImproperStructStatementToSimpleStatementBodyAction = &_Action{_ReduceAction, 0, _ReduceExpressionOrImproperStructStatementToSimpleStatementBody}
	_ReduceCallbackOpStatementToSimpleStatementBodyAction                 = &_Action{_ReduceAction, 0, _ReduceCallbackOpStatementToSimpleStatementBody}
	_ReduceJumpStatementToSimpleStatementBodyAction                       = &_Action{_ReduceAction, 0, _ReduceJumpStatementToSimpleStatementBody}
	_ReduceAssignStatementToSimpleStatementBodyAction                     = &_Action{_ReduceAction, 0, _ReduceAssignStatementToSimpleStatementBody}
	_ReduceUnaryOpAssignStatementToSimpleStatementBodyAction              = &_Action{_ReduceAction, 0, _ReduceUnaryOpAssignStatementToSimpleStatementBody}
	_ReduceBinaryOpAssignStatementToSimpleStatementBodyAction             = &_Action{_ReduceAction, 0, _ReduceBinaryOpAssignStatementToSimpleStatementBody}
	_ReduceFallthroughStatementToSimpleStatementBodyAction                = &_Action{_ReduceAction, 0, _ReduceFallthroughStatementToSimpleStatementBody}
	_ReduceSimpleStatementBodyToStatementBodyAction                       = &_Action{_ReduceAction, 0, _ReduceSimpleStatementBodyToStatementBody}
	_ReduceImportStatementToStatementBodyAction                           = &_Action{_ReduceAction, 0, _ReduceImportStatementToStatementBody}
	_ReduceCaseBranchStatementToStatementBodyAction                       = &_Action{_ReduceAction, 0, _ReduceCaseBranchStatementToStatementBody}
	_ReduceDefaultBranchStatementToStatementBodyAction                    = &_Action{_ReduceAction, 0, _ReduceDefaultBranchStatementToStatementBody}
	_ReduceSimpleStatementBodyToOptionalSimpleStatementBodyAction         = &_Action{_ReduceAction, 0, _ReduceSimpleStatementBodyToOptionalSimpleStatementBody}
	_ReduceNilToOptionalSimpleStatementBodyAction                         = &_Action{_ReduceAction, 0, _ReduceNilToOptionalSimpleStatementBody}
	_ReduceToExpressionOrImproperStructStatementAction                    = &_Action{_ReduceAction, 0, _ReduceToExpressionOrImproperStructStatement}
	_ReduceAsyncToCallbackOpAction                                        = &_Action{_ReduceAction, 0, _ReduceAsyncToCallbackOp}
	_ReduceDeferToCallbackOpAction                                        = &_Action{_ReduceAction, 0, _ReduceDeferToCallbackOp}
	_ReduceToCallbackOpStatementAction                                    = &_Action{_ReduceAction, 0, _ReduceToCallbackOpStatement}
	_ReduceAddOneAssignToUnaryOpAssignAction                              = &_Action{_ReduceAction, 0, _ReduceAddOneAssignToUnaryOpAssign}
	_ReduceSubOneAssignToUnaryOpAssignAction                              = &_Action{_ReduceAction, 0, _ReduceSubOneAssignToUnaryOpAssign}
	_ReduceAddAssignToBinaryOpAssignAction                                = &_Action{_ReduceAction, 0, _ReduceAddAssignToBinaryOpAssign}
	_ReduceSubAssignToBinaryOpAssignAction                                = &_Action{_ReduceAction, 0, _ReduceSubAssignToBinaryOpAssign}
	_ReduceMulAssignToBinaryOpAssignAction                                = &_Action{_ReduceAction, 0, _ReduceMulAssignToBinaryOpAssign}
	_ReduceDivAssignToBinaryOpAssignAction                                = &_Action{_ReduceAction, 0, _ReduceDivAssignToBinaryOpAssign}
	_ReduceModAssignToBinaryOpAssignAction                                = &_Action{_ReduceAction, 0, _ReduceModAssignToBinaryOpAssign}
	_ReduceBitNegAssignToBinaryOpAssignAction                             = &_Action{_ReduceAction, 0, _ReduceBitNegAssignToBinaryOpAssign}
	_ReduceBitAndAssignToBinaryOpAssignAction                             = &_Action{_ReduceAction, 0, _ReduceBitAndAssignToBinaryOpAssign}
	_ReduceBitOrAssignToBinaryOpAssignAction                              = &_Action{_ReduceAction, 0, _ReduceBitOrAssignToBinaryOpAssign}
	_ReduceBitXorAssignToBinaryOpAssignAction                             = &_Action{_ReduceAction, 0, _ReduceBitXorAssignToBinaryOpAssign}
	_ReduceBitLshiftAssignToBinaryOpAssignAction                          = &_Action{_ReduceAction, 0, _ReduceBitLshiftAssignToBinaryOpAssign}
	_ReduceBitRshiftAssignToBinaryOpAssignAction                          = &_Action{_ReduceAction, 0, _ReduceBitRshiftAssignToBinaryOpAssign}
	_ReduceToUnsafeStatementAction                                        = &_Action{_ReduceAction, 0, _ReduceToUnsafeStatement}
	_ReduceUnlabeledNoValueToJumpStatementAction                          = &_Action{_ReduceAction, 0, _ReduceUnlabeledNoValueToJumpStatement}
	_ReduceUnlabeledValuedToJumpStatementAction                           = &_Action{_ReduceAction, 0, _ReduceUnlabeledValuedToJumpStatement}
	_ReduceLabeledNoValueToJumpStatementAction                            = &_Action{_ReduceAction, 0, _ReduceLabeledNoValueToJumpStatement}
	_ReduceLabeledValuedToJumpStatementAction                             = &_Action{_ReduceAction, 0, _ReduceLabeledValuedToJumpStatement}
	_ReduceReturnToJumpTypeAction                                         = &_Action{_ReduceAction, 0, _ReduceReturnToJumpType}
	_ReduceBreakToJumpTypeAction                                          = &_Action{_ReduceAction, 0, _ReduceBreakToJumpType}
	_ReduceContinueToJumpTypeAction                                       = &_Action{_ReduceAction, 0, _ReduceContinueToJumpType}
	_ReduceExpressionToExpressionsAction                                  = &_Action{_ReduceAction, 0, _ReduceExpressionToExpressions}
	_ReduceAddToExpressionsAction                                         = &_Action{_ReduceAction, 0, _ReduceAddToExpressions}
	_ReduceToFallthroughStatementAction                                   = &_Action{_ReduceAction, 0, _ReduceToFallthroughStatement}
	_ReduceToAssignStatementAction                                        = &_Action{_ReduceAction, 0, _ReduceToAssignStatement}
	_ReduceToUnaryOpAssignStatementAction                                 = &_Action{_ReduceAction, 0, _ReduceToUnaryOpAssignStatement}
	_ReduceToBinaryOpAssignStatementAction                                = &_Action{_ReduceAction, 0, _ReduceToBinaryOpAssignStatement}
	_ReduceSingleToImportStatementAction                                  = &_Action{_ReduceAction, 0, _ReduceSingleToImportStatement}
	_ReduceMultipleToImportStatementAction                                = &_Action{_ReduceAction, 0, _ReduceMultipleToImportStatement}
	_ReduceStringLiteralToImportClauseAction                              = &_Action{_ReduceAction, 0, _ReduceStringLiteralToImportClause}
	_ReduceAliasToImportClauseAction                                      = &_Action{_ReduceAction, 0, _ReduceAliasToImportClause}
	_ReduceImplicitToImportClauseTerminalAction                           = &_Action{_ReduceAction, 0, _ReduceImplicitToImportClauseTerminal}
	_ReduceExplicitToImportClauseTerminalAction                           = &_Action{_ReduceAction, 0, _ReduceExplicitToImportClauseTerminal}
	_ReduceFirstToImportClausesAction                                     = &_Action{_ReduceAction, 0, _ReduceFirstToImportClauses}
	_ReduceAddToImportClausesAction                                       = &_Action{_ReduceAction, 0, _ReduceAddToImportClauses}
	_ReduceCasePatternToCasePatternsAction                                = &_Action{_ReduceAction, 0, _ReduceCasePatternToCasePatterns}
	_ReduceMultipleToCasePatternsAction                                   = &_Action{_ReduceAction, 0, _ReduceMultipleToCasePatterns}
	_ReduceToVarDeclPatternAction                                         = &_Action{_ReduceAction, 0, _ReduceToVarDeclPattern}
	_ReduceVarToVarOrLetAction                                            = &_Action{_ReduceAction, 0, _ReduceVarToVarOrLet}
	_ReduceLetToVarOrLetAction                                            = &_Action{_ReduceAction, 0, _ReduceLetToVarOrLet}
	_ReduceIdentifierToVarPatternAction                                   = &_Action{_ReduceAction, 0, _ReduceIdentifierToVarPattern}
	_ReduceTuplePatternToVarPatternAction                                 = &_Action{_ReduceAction, 0, _ReduceTuplePatternToVarPattern}
	_ReduceToTuplePatternAction                                           = &_Action{_ReduceAction, 0, _ReduceToTuplePattern}
	_ReduceFieldVarPatternToFieldVarPatternsAction                        = &_Action{_ReduceAction, 0, _ReduceFieldVarPatternToFieldVarPatterns}
	_ReduceAddToFieldVarPatternsAction                                    = &_Action{_ReduceAction, 0, _ReduceAddToFieldVarPatterns}
	_ReducePositionalToFieldVarPatternAction                              = &_Action{_ReduceAction, 0, _ReducePositionalToFieldVarPattern}
	_ReduceNamedToFieldVarPatternAction                                   = &_Action{_ReduceAction, 0, _ReduceNamedToFieldVarPattern}
	_ReduceDotDotDotToFieldVarPatternAction                               = &_Action{_ReduceAction, 0, _ReduceDotDotDotToFieldVarPattern}
	_ReduceValueTypeToOptionalValueTypeAction                             = &_Action{_ReduceAction, 0, _ReduceValueTypeToOptionalValueType}
	_ReduceNilToOptionalValueTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceNilToOptionalValueType}
	_ReduceToAssignPatternAction                                          = &_Action{_ReduceAction, 0, _ReduceToAssignPattern}
	_ReduceAssignPatternToCasePatternAction                               = &_Action{_ReduceAction, 0, _ReduceAssignPatternToCasePattern}
	_ReduceEnumMatchPatternToCasePatternAction                            = &_Action{_ReduceAction, 0, _ReduceEnumMatchPatternToCasePattern}
	_ReduceEnumVarDeclPatternToCasePatternAction                          = &_Action{_ReduceAction, 0, _ReduceEnumVarDeclPatternToCasePattern}
	_ReduceIfExprToExpressionAction                                       = &_Action{_ReduceAction, 0, _ReduceIfExprToExpression}
	_ReduceSwitchExprToExpressionAction                                   = &_Action{_ReduceAction, 0, _ReduceSwitchExprToExpression}
	_ReduceLoopExprToExpressionAction                                     = &_Action{_ReduceAction, 0, _ReduceLoopExprToExpression}
	_ReduceSequenceExprToExpressionAction                                 = &_Action{_ReduceAction, 0, _ReduceSequenceExprToExpression}
	_ReduceLabelDeclToOptionalLabelDeclAction                             = &_Action{_ReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}
	_ReduceUnlabelledToOptionalLabelDeclAction                            = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}
	_ReduceOrExprToSequenceExprAction                                     = &_Action{_ReduceAction, 0, _ReduceOrExprToSequenceExpr}
	_ReduceVarDeclPatternToSequenceExprAction                             = &_Action{_ReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}
	_ReduceAssignVarPatternToSequenceExprAction                           = &_Action{_ReduceAction, 0, _ReduceAssignVarPatternToSequenceExpr}
	_ReduceNoElseToIfExprAction                                           = &_Action{_ReduceAction, 0, _ReduceNoElseToIfExpr}
	_ReduceIfElseToIfExprAction                                           = &_Action{_ReduceAction, 0, _ReduceIfElseToIfExpr}
	_ReduceMultiIfElseToIfExprAction                                      = &_Action{_ReduceAction, 0, _ReduceMultiIfElseToIfExpr}
	_ReduceSequenceExprToConditionAction                                  = &_Action{_ReduceAction, 0, _ReduceSequenceExprToCondition}
	_ReduceCaseToConditionAction                                          = &_Action{_ReduceAction, 0, _ReduceCaseToCondition}
	_ReduceToSwitchExprAction                                             = &_Action{_ReduceAction, 0, _ReduceToSwitchExpr}
	_ReduceInfiniteToLoopExprAction                                       = &_Action{_ReduceAction, 0, _ReduceInfiniteToLoopExpr}
	_ReduceDoWhileToLoopExprAction                                        = &_Action{_ReduceAction, 0, _ReduceDoWhileToLoopExpr}
	_ReduceWhileToLoopExprAction                                          = &_Action{_ReduceAction, 0, _ReduceWhileToLoopExpr}
	_ReduceIteratorToLoopExprAction                                       = &_Action{_ReduceAction, 0, _ReduceIteratorToLoopExpr}
	_ReduceForToLoopExprAction                                            = &_Action{_ReduceAction, 0, _ReduceForToLoopExpr}
	_ReduceSequenceExprToOptionalSequenceExprAction                       = &_Action{_ReduceAction, 0, _ReduceSequenceExprToOptionalSequenceExpr}
	_ReduceNilToOptionalSequenceExprAction                                = &_Action{_ReduceAction, 0, _ReduceNilToOptionalSequenceExpr}
	_ReduceSequenceExprToForAssignmentAction                              = &_Action{_ReduceAction, 0, _ReduceSequenceExprToForAssignment}
	_ReduceAssignToForAssignmentAction                                    = &_Action{_ReduceAction, 0, _ReduceAssignToForAssignment}
	_ReduceToCallExprAction                                               = &_Action{_ReduceAction, 0, _ReduceToCallExpr}
	_ReduceBindingToOptionalGenericBindingAction                          = &_Action{_ReduceAction, 0, _ReduceBindingToOptionalGenericBinding}
	_ReduceNilToOptionalGenericBindingAction                              = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericBinding}
	_ReduceGenericArgumentsToOptionalGenericArgumentsAction               = &_Action{_ReduceAction, 0, _ReduceGenericArgumentsToOptionalGenericArguments}
	_ReduceNilToOptionalGenericArgumentsAction                            = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericArguments}
	_ReduceValueTypeToGenericArgumentsAction                              = &_Action{_ReduceAction, 0, _ReduceValueTypeToGenericArguments}
	_ReduceAddToGenericArgumentsAction                                    = &_Action{_ReduceAction, 0, _ReduceAddToGenericArguments}
	_ReduceArgumentsToOptionalArgumentsAction                             = &_Action{_ReduceAction, 0, _ReduceArgumentsToOptionalArguments}
	_ReduceNilToOptionalArgumentsAction                                   = &_Action{_ReduceAction, 0, _ReduceNilToOptionalArguments}
	_ReduceArgumentToArgumentsAction                                      = &_Action{_ReduceAction, 0, _ReduceArgumentToArguments}
	_ReduceAddToArgumentsAction                                           = &_Action{_ReduceAction, 0, _ReduceAddToArguments}
	_ReducePositionalToArgumentAction                                     = &_Action{_ReduceAction, 0, _ReducePositionalToArgument}
	_ReduceNamedToArgumentAction                                          = &_Action{_ReduceAction, 0, _ReduceNamedToArgument}
	_ReduceColonExpressionsToArgumentAction                               = &_Action{_ReduceAction, 0, _ReduceColonExpressionsToArgument}
	_ReduceDotDotDotToArgumentAction                                      = &_Action{_ReduceAction, 0, _ReduceDotDotDotToArgument}
	_ReducePairToColonExpressionsAction                                   = &_Action{_ReduceAction, 0, _ReducePairToColonExpressions}
	_ReduceAddToColonExpressionsAction                                    = &_Action{_ReduceAction, 0, _ReduceAddToColonExpressions}
	_ReduceExpressionToOptionalExpressionAction                           = &_Action{_ReduceAction, 0, _ReduceExpressionToOptionalExpression}
	_ReduceNilToOptionalExpressionAction                                  = &_Action{_ReduceAction, 0, _ReduceNilToOptionalExpression}
	_ReduceParseErrorExprToAtomExprAction                                 = &_Action{_ReduceAction, 0, _ReduceParseErrorExprToAtomExpr}
	_ReduceLiteralExprToAtomExprAction                                    = &_Action{_ReduceAction, 0, _ReduceLiteralExprToAtomExpr}
	_ReduceIdentifierExprToAtomExprAction                                 = &_Action{_ReduceAction, 0, _ReduceIdentifierExprToAtomExpr}
	_ReduceBlockExprToAtomExprAction                                      = &_Action{_ReduceAction, 0, _ReduceBlockExprToAtomExpr}
	_ReduceAnonymousFuncExprToAtomExprAction                              = &_Action{_ReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}
	_ReduceInitializeExprToAtomExprAction                                 = &_Action{_ReduceAction, 0, _ReduceInitializeExprToAtomExpr}
	_ReduceImplicitStructExprToAtomExprAction                             = &_Action{_ReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}
	_ReduceToParseErrorExprAction                                         = &_Action{_ReduceAction, 0, _ReduceToParseErrorExpr}
	_ReduceTrueToLiteralExprAction                                        = &_Action{_ReduceAction, 0, _ReduceTrueToLiteralExpr}
	_ReduceFalseToLiteralExprAction                                       = &_Action{_ReduceAction, 0, _ReduceFalseToLiteralExpr}
	_ReduceIntegerLiteralToLiteralExprAction                              = &_Action{_ReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}
	_ReduceFloatLiteralToLiteralExprAction                                = &_Action{_ReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}
	_ReduceRuneLiteralToLiteralExprAction                                 = &_Action{_ReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}
	_ReduceStringLiteralToLiteralExprAction                               = &_Action{_ReduceAction, 0, _ReduceStringLiteralToLiteralExpr}
	_ReduceToIdentifierExprAction                                         = &_Action{_ReduceAction, 0, _ReduceToIdentifierExpr}
	_ReduceToBlockExprAction                                              = &_Action{_ReduceAction, 0, _ReduceToBlockExpr}
	_ReduceToInitializeExprAction                                         = &_Action{_ReduceAction, 0, _ReduceToInitializeExpr}
	_ReduceToImplicitStructExprAction                                     = &_Action{_ReduceAction, 0, _ReduceToImplicitStructExpr}
	_ReduceAtomExprToAccessibleExprAction                                 = &_Action{_ReduceAction, 0, _ReduceAtomExprToAccessibleExpr}
	_ReduceAccessExprToAccessibleExprAction                               = &_Action{_ReduceAction, 0, _ReduceAccessExprToAccessibleExpr}
	_ReduceCallExprToAccessibleExprAction                                 = &_Action{_ReduceAction, 0, _ReduceCallExprToAccessibleExpr}
	_ReduceIndexExprToAccessibleExprAction                                = &_Action{_ReduceAction, 0, _ReduceIndexExprToAccessibleExpr}
	_ReduceToAccessExprAction                                             = &_Action{_ReduceAction, 0, _ReduceToAccessExpr}
	_ReduceToIndexExprAction                                              = &_Action{_ReduceAction, 0, _ReduceToIndexExpr}
	_ReduceAccessibleExprToPostfixableExprAction                          = &_Action{_ReduceAction, 0, _ReduceAccessibleExprToPostfixableExpr}
	_ReducePostfixUnaryExprToPostfixableExprAction                        = &_Action{_ReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}
	_ReduceToPostfixUnaryExprAction                                       = &_Action{_ReduceAction, 0, _ReduceToPostfixUnaryExpr}
	_ReducePostfixableExprToPrefixableExprAction                          = &_Action{_ReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}
	_ReducePrefixUnaryExprToPrefixableExprAction                          = &_Action{_ReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}
	_ReduceToPrefixUnaryExprAction                                        = &_Action{_ReduceAction, 0, _ReduceToPrefixUnaryExpr}
	_ReduceNotToPrefixUnaryOpAction                                       = &_Action{_ReduceAction, 0, _ReduceNotToPrefixUnaryOp}
	_ReduceBitNegToPrefixUnaryOpAction                                    = &_Action{_ReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}
	_ReduceSubToPrefixUnaryOpAction                                       = &_Action{_ReduceAction, 0, _ReduceSubToPrefixUnaryOp}
	_ReduceMulToPrefixUnaryOpAction                                       = &_Action{_ReduceAction, 0, _ReduceMulToPrefixUnaryOp}
	_ReduceBitAndToPrefixUnaryOpAction                                    = &_Action{_ReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}
	_ReducePrefixableExprToMulExprAction                                  = &_Action{_ReduceAction, 0, _ReducePrefixableExprToMulExpr}
	_ReduceBinaryMulExprToMulExprAction                                   = &_Action{_ReduceAction, 0, _ReduceBinaryMulExprToMulExpr}
	_ReduceToBinaryMulExprAction                                          = &_Action{_ReduceAction, 0, _ReduceToBinaryMulExpr}
	_ReduceMulToMulOpAction                                               = &_Action{_ReduceAction, 0, _ReduceMulToMulOp}
	_ReduceDivToMulOpAction                                               = &_Action{_ReduceAction, 0, _ReduceDivToMulOp}
	_ReduceModToMulOpAction                                               = &_Action{_ReduceAction, 0, _ReduceModToMulOp}
	_ReduceBitAndToMulOpAction                                            = &_Action{_ReduceAction, 0, _ReduceBitAndToMulOp}
	_ReduceBitLshiftToMulOpAction                                         = &_Action{_ReduceAction, 0, _ReduceBitLshiftToMulOp}
	_ReduceBitRshiftToMulOpAction                                         = &_Action{_ReduceAction, 0, _ReduceBitRshiftToMulOp}
	_ReduceMulExprToAddExprAction                                         = &_Action{_ReduceAction, 0, _ReduceMulExprToAddExpr}
	_ReduceBinaryAddExprToAddExprAction                                   = &_Action{_ReduceAction, 0, _ReduceBinaryAddExprToAddExpr}
	_ReduceToBinaryAddExprAction                                          = &_Action{_ReduceAction, 0, _ReduceToBinaryAddExpr}
	_ReduceAddToAddOpAction                                               = &_Action{_ReduceAction, 0, _ReduceAddToAddOp}
	_ReduceSubToAddOpAction                                               = &_Action{_ReduceAction, 0, _ReduceSubToAddOp}
	_ReduceBitOrToAddOpAction                                             = &_Action{_ReduceAction, 0, _ReduceBitOrToAddOp}
	_ReduceBitXorToAddOpAction                                            = &_Action{_ReduceAction, 0, _ReduceBitXorToAddOp}
	_ReduceAddExprToCmpExprAction                                         = &_Action{_ReduceAction, 0, _ReduceAddExprToCmpExpr}
	_ReduceBinaryCmpExprToCmpExprAction                                   = &_Action{_ReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}
	_ReduceToBinaryCmpExprAction                                          = &_Action{_ReduceAction, 0, _ReduceToBinaryCmpExpr}
	_ReduceEqualToCmpOpAction                                             = &_Action{_ReduceAction, 0, _ReduceEqualToCmpOp}
	_ReduceNotEqualToCmpOpAction                                          = &_Action{_ReduceAction, 0, _ReduceNotEqualToCmpOp}
	_ReduceLessToCmpOpAction                                              = &_Action{_ReduceAction, 0, _ReduceLessToCmpOp}
	_ReduceLessOrEqualToCmpOpAction                                       = &_Action{_ReduceAction, 0, _ReduceLessOrEqualToCmpOp}
	_ReduceGreaterToCmpOpAction                                           = &_Action{_ReduceAction, 0, _ReduceGreaterToCmpOp}
	_ReduceGreaterOrEqualToCmpOpAction                                    = &_Action{_ReduceAction, 0, _ReduceGreaterOrEqualToCmpOp}
	_ReduceCmpExprToAndExprAction                                         = &_Action{_ReduceAction, 0, _ReduceCmpExprToAndExpr}
	_ReduceBinaryAndExprToAndExprAction                                   = &_Action{_ReduceAction, 0, _ReduceBinaryAndExprToAndExpr}
	_ReduceToBinaryAndExprAction                                          = &_Action{_ReduceAction, 0, _ReduceToBinaryAndExpr}
	_ReduceAndExprToOrExprAction                                          = &_Action{_ReduceAction, 0, _ReduceAndExprToOrExpr}
	_ReduceBinaryOrExprToOrExprAction                                     = &_Action{_ReduceAction, 0, _ReduceBinaryOrExprToOrExpr}
	_ReduceToBinaryOrExprAction                                           = &_Action{_ReduceAction, 0, _ReduceToBinaryOrExpr}
	_ReduceExplicitStructDefToInitializableTypeAction                     = &_Action{_ReduceAction, 0, _ReduceExplicitStructDefToInitializableType}
	_ReduceSliceToInitializableTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceSliceToInitializableType}
	_ReduceArrayToInitializableTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceArrayToInitializableType}
	_ReduceMapToInitializableTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceMapToInitializableType}
	_ReduceInitializableTypeToAtomTypeAction                              = &_Action{_ReduceAction, 0, _ReduceInitializableTypeToAtomType}
	_ReduceNamedToAtomTypeAction                                          = &_Action{_ReduceAction, 0, _ReduceNamedToAtomType}
	_ReduceExternNamedToAtomTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceExternNamedToAtomType}
	_ReduceInferredToAtomTypeAction                                       = &_Action{_ReduceAction, 0, _ReduceInferredToAtomType}
	_ReduceImplicitStructDefToAtomTypeAction                              = &_Action{_ReduceAction, 0, _ReduceImplicitStructDefToAtomType}
	_ReduceExplicitEnumDefToAtomTypeAction                                = &_Action{_ReduceAction, 0, _ReduceExplicitEnumDefToAtomType}
	_ReduceImplicitEnumDefToAtomTypeAction                                = &_Action{_ReduceAction, 0, _ReduceImplicitEnumDefToAtomType}
	_ReduceTraitDefToAtomTypeAction                                       = &_Action{_ReduceAction, 0, _ReduceTraitDefToAtomType}
	_ReduceFuncTypeToAtomTypeAction                                       = &_Action{_ReduceAction, 0, _ReduceFuncTypeToAtomType}
	_ReduceParseErrorTypeToAtomTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceParseErrorTypeToAtomType}
	_ReduceToParseErrorTypeAction                                         = &_Action{_ReduceAction, 0, _ReduceToParseErrorType}
	_ReduceAtomTypeToReturnableTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceAtomTypeToReturnableType}
	_ReducePrefixedTypeToReturnableTypeAction                             = &_Action{_ReduceAction, 0, _ReducePrefixedTypeToReturnableType}
	_ReduceToPrefixedTypeAction                                           = &_Action{_ReduceAction, 0, _ReduceToPrefixedType}
	_ReduceQuestionToPrefixTypeOpAction                                   = &_Action{_ReduceAction, 0, _ReduceQuestionToPrefixTypeOp}
	_ReduceExclaimToPrefixTypeOpAction                                    = &_Action{_ReduceAction, 0, _ReduceExclaimToPrefixTypeOp}
	_ReduceBitAndToPrefixTypeOpAction                                     = &_Action{_ReduceAction, 0, _ReduceBitAndToPrefixTypeOp}
	_ReduceBitNegToPrefixTypeOpAction                                     = &_Action{_ReduceAction, 0, _ReduceBitNegToPrefixTypeOp}
	_ReduceTildeTildeToPrefixTypeOpAction                                 = &_Action{_ReduceAction, 0, _ReduceTildeTildeToPrefixTypeOp}
	_ReduceReturnableTypeToValueTypeAction                                = &_Action{_ReduceAction, 0, _ReduceReturnableTypeToValueType}
	_ReduceTraitOpTypeToValueTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceTraitOpTypeToValueType}
	_ReduceToTraitOpTypeAction                                            = &_Action{_ReduceAction, 0, _ReduceToTraitOpType}
	_ReduceMulToTraitOpAction                                             = &_Action{_ReduceAction, 0, _ReduceMulToTraitOp}
	_ReduceAddToTraitOpAction                                             = &_Action{_ReduceAction, 0, _ReduceAddToTraitOp}
	_ReduceSubToTraitOpAction                                             = &_Action{_ReduceAction, 0, _ReduceSubToTraitOp}
	_ReduceDefinitionToTypeDefAction                                      = &_Action{_ReduceAction, 0, _ReduceDefinitionToTypeDef}
	_ReduceConstrainedDefToTypeDefAction                                  = &_Action{_ReduceAction, 0, _ReduceConstrainedDefToTypeDef}
	_ReduceAliasToTypeDefAction                                           = &_Action{_ReduceAction, 0, _ReduceAliasToTypeDef}
	_ReduceUnconstrainedToGenericParameterDefAction                       = &_Action{_ReduceAction, 0, _ReduceUnconstrainedToGenericParameterDef}
	_ReduceConstrainedToGenericParameterDefAction                         = &_Action{_ReduceAction, 0, _ReduceConstrainedToGenericParameterDef}
	_ReduceGenericParameterDefToGenericParameterDefsAction                = &_Action{_ReduceAction, 0, _ReduceGenericParameterDefToGenericParameterDefs}
	_ReduceAddToGenericParameterDefsAction                                = &_Action{_ReduceAction, 0, _ReduceAddToGenericParameterDefs}
	_ReduceGenericParameterDefsToOptionalGenericParameterDefsAction       = &_Action{_ReduceAction, 0, _ReduceGenericParameterDefsToOptionalGenericParameterDefs}
	_ReduceNilToOptionalGenericParameterDefsAction                        = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericParameterDefs}
	_ReduceGenericToOptionalGenericParametersAction                       = &_Action{_ReduceAction, 0, _ReduceGenericToOptionalGenericParameters}
	_ReduceNilToOptionalGenericParametersAction                           = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericParameters}
	_ReduceExplicitToFieldDefAction                                       = &_Action{_ReduceAction, 0, _ReduceExplicitToFieldDef}
	_ReduceImplicitToFieldDefAction                                       = &_Action{_ReduceAction, 0, _ReduceImplicitToFieldDef}
	_ReduceUnsafeStatementToFieldDefAction                                = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToFieldDef}
	_ReduceFieldDefToImplicitFieldDefsAction                              = &_Action{_ReduceAction, 0, _ReduceFieldDefToImplicitFieldDefs}
	_ReduceAddToImplicitFieldDefsAction                                   = &_Action{_ReduceAction, 0, _ReduceAddToImplicitFieldDefs}
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction             = &_Action{_ReduceAction, 0, _ReduceImplicitFieldDefsToOptionalImplicitFieldDefs}
	_ReduceNilToOptionalImplicitFieldDefsAction                           = &_Action{_ReduceAction, 0, _ReduceNilToOptionalImplicitFieldDefs}
	_ReduceToImplicitStructDefAction                                      = &_Action{_ReduceAction, 0, _ReduceToImplicitStructDef}
	_ReduceFieldDefToExplicitFieldDefsAction                              = &_Action{_ReduceAction, 0, _ReduceFieldDefToExplicitFieldDefs}
	_ReduceImplicitToExplicitFieldDefsAction                              = &_Action{_ReduceAction, 0, _ReduceImplicitToExplicitFieldDefs}
	_ReduceExplicitToExplicitFieldDefsAction                              = &_Action{_ReduceAction, 0, _ReduceExplicitToExplicitFieldDefs}
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction             = &_Action{_ReduceAction, 0, _ReduceExplicitFieldDefsToOptionalExplicitFieldDefs}
	_ReduceNilToOptionalExplicitFieldDefsAction                           = &_Action{_ReduceAction, 0, _ReduceNilToOptionalExplicitFieldDefs}
	_ReduceToExplicitStructDefAction                                      = &_Action{_ReduceAction, 0, _ReduceToExplicitStructDef}
	_ReduceFieldDefToEnumValueDefAction                                   = &_Action{_ReduceAction, 0, _ReduceFieldDefToEnumValueDef}
	_ReduceDefaultToEnumValueDefAction                                    = &_Action{_ReduceAction, 0, _ReduceDefaultToEnumValueDef}
	_ReducePairToImplicitEnumValueDefsAction                              = &_Action{_ReduceAction, 0, _ReducePairToImplicitEnumValueDefs}
	_ReduceAddToImplicitEnumValueDefsAction                               = &_Action{_ReduceAction, 0, _ReduceAddToImplicitEnumValueDefs}
	_ReduceToImplicitEnumDefAction                                        = &_Action{_ReduceAction, 0, _ReduceToImplicitEnumDef}
	_ReduceExplicitPairToExplicitEnumValueDefsAction                      = &_Action{_ReduceAction, 0, _ReduceExplicitPairToExplicitEnumValueDefs}
	_ReduceImplicitPairToExplicitEnumValueDefsAction                      = &_Action{_ReduceAction, 0, _ReduceImplicitPairToExplicitEnumValueDefs}
	_ReduceExplicitAddToExplicitEnumValueDefsAction                       = &_Action{_ReduceAction, 0, _ReduceExplicitAddToExplicitEnumValueDefs}
	_ReduceImplicitAddToExplicitEnumValueDefsAction                       = &_Action{_ReduceAction, 0, _ReduceImplicitAddToExplicitEnumValueDefs}
	_ReduceToExplicitEnumDefAction                                        = &_Action{_ReduceAction, 0, _ReduceToExplicitEnumDef}
	_ReduceFieldDefToTraitPropertyAction                                  = &_Action{_ReduceAction, 0, _ReduceFieldDefToTraitProperty}
	_ReduceMethodSignatureToTraitPropertyAction                           = &_Action{_ReduceAction, 0, _ReduceMethodSignatureToTraitProperty}
	_ReduceTraitPropertyToTraitPropertiesAction                           = &_Action{_ReduceAction, 0, _ReduceTraitPropertyToTraitProperties}
	_ReduceImplicitToTraitPropertiesAction                                = &_Action{_ReduceAction, 0, _ReduceImplicitToTraitProperties}
	_ReduceExplicitToTraitPropertiesAction                                = &_Action{_ReduceAction, 0, _ReduceExplicitToTraitProperties}
	_ReduceTraitPropertiesToOptionalTraitPropertiesAction                 = &_Action{_ReduceAction, 0, _ReduceTraitPropertiesToOptionalTraitProperties}
	_ReduceNilToOptionalTraitPropertiesAction                             = &_Action{_ReduceAction, 0, _ReduceNilToOptionalTraitProperties}
	_ReduceToTraitDefAction                                               = &_Action{_ReduceAction, 0, _ReduceToTraitDef}
	_ReduceReturnableTypeToReturnTypeAction                               = &_Action{_ReduceAction, 0, _ReduceReturnableTypeToReturnType}
	_ReduceNilToReturnTypeAction                                          = &_Action{_ReduceAction, 0, _ReduceNilToReturnType}
	_ReduceArgToParameterDeclAction                                       = &_Action{_ReduceAction, 0, _ReduceArgToParameterDecl}
	_ReduceVarargToParameterDeclAction                                    = &_Action{_ReduceAction, 0, _ReduceVarargToParameterDecl}
	_ReduceUnamedToParameterDeclAction                                    = &_Action{_ReduceAction, 0, _ReduceUnamedToParameterDecl}
	_ReduceUnnamedVarargToParameterDeclAction                             = &_Action{_ReduceAction, 0, _ReduceUnnamedVarargToParameterDecl}
	_ReduceParameterDeclToParameterDeclsAction                            = &_Action{_ReduceAction, 0, _ReduceParameterDeclToParameterDecls}
	_ReduceAddToParameterDeclsAction                                      = &_Action{_ReduceAction, 0, _ReduceAddToParameterDecls}
	_ReduceParameterDeclsToOptionalParameterDeclsAction                   = &_Action{_ReduceAction, 0, _ReduceParameterDeclsToOptionalParameterDecls}
	_ReduceNilToOptionalParameterDeclsAction                              = &_Action{_ReduceAction, 0, _ReduceNilToOptionalParameterDecls}
	_ReduceToFuncTypeAction                                               = &_Action{_ReduceAction, 0, _ReduceToFuncType}
	_ReduceToMethodSignatureAction                                        = &_Action{_ReduceAction, 0, _ReduceToMethodSignature}
	_ReduceInferredRefArgToParameterDefAction                             = &_Action{_ReduceAction, 0, _ReduceInferredRefArgToParameterDef}
	_ReduceInferredRefVarargToParameterDefAction                          = &_Action{_ReduceAction, 0, _ReduceInferredRefVarargToParameterDef}
	_ReduceArgToParameterDefAction                                        = &_Action{_ReduceAction, 0, _ReduceArgToParameterDef}
	_ReduceVarargToParameterDefAction                                     = &_Action{_ReduceAction, 0, _ReduceVarargToParameterDef}
	_ReduceParameterDefToParameterDefsAction                              = &_Action{_ReduceAction, 0, _ReduceParameterDefToParameterDefs}
	_ReduceAddToParameterDefsAction                                       = &_Action{_ReduceAction, 0, _ReduceAddToParameterDefs}
	_ReduceParameterDefsToOptionalParameterDefsAction                     = &_Action{_ReduceAction, 0, _ReduceParameterDefsToOptionalParameterDefs}
	_ReduceNilToOptionalParameterDefsAction                               = &_Action{_ReduceAction, 0, _ReduceNilToOptionalParameterDefs}
	_ReduceFuncDefToNamedFuncDefAction                                    = &_Action{_ReduceAction, 0, _ReduceFuncDefToNamedFuncDef}
	_ReduceMethodDefToNamedFuncDefAction                                  = &_Action{_ReduceAction, 0, _ReduceMethodDefToNamedFuncDef}
	_ReduceAliasToNamedFuncDefAction                                      = &_Action{_ReduceAction, 0, _ReduceAliasToNamedFuncDef}
	_ReduceToAnonymousFuncExprAction                                      = &_Action{_ReduceAction, 0, _ReduceToAnonymousFuncExpr}
	_ReduceNoSpecToPackageDefAction                                       = &_Action{_ReduceAction, 0, _ReduceNoSpecToPackageDef}
	_ReduceWithSpecToPackageDefAction                                     = &_Action{_ReduceAction, 0, _ReduceWithSpecToPackageDef}
)

var _ActionTable = _ActionTableType{
	{_State6, _EndMarker}:                                &_Action{_AcceptAction, 0, 0},
	{_State7, _EndMarker}:                                &_Action{_AcceptAction, 0, 0},
	{_State8, _EndMarker}:                                &_Action{_AcceptAction, 0, 0},
	{_State9, _EndMarker}:                                &_Action{_AcceptAction, 0, 0},
	{_State10, _EndMarker}:                               &_Action{_AcceptAction, 0, 0},
	{_State1, NewlinesToken}:                             _GotoState11Action,
	{_State1, SourceType}:                                _GotoState6Action,
	{_State1, OptionalDefinitionsType}:                   _GotoState12Action,
	{_State1, OptionalNewlinesType}:                      _GotoState13Action,
	{_State2, PackageToken}:                              _GotoState14Action,
	{_State2, PackageDefType}:                            _GotoState7Action,
	{_State3, TypeToken}:                                 _GotoState15Action,
	{_State3, TypeDefType}:                               _GotoState8Action,
	{_State4, FuncToken}:                                 _GotoState16Action,
	{_State4, NamedFuncDefType}:                          _GotoState9Action,
	{_State5, IntegerLiteralToken}:                       _GotoState24Action,
	{_State5, FloatLiteralToken}:                         _GotoState20Action,
	{_State5, RuneLiteralToken}:                          _GotoState32Action,
	{_State5, StringLiteralToken}:                        _GotoState33Action,
	{_State5, IdentifierToken}:                           _GotoState23Action,
	{_State5, TrueToken}:                                 _GotoState36Action,
	{_State5, FalseToken}:                                _GotoState19Action,
	{_State5, StructToken}:                               _GotoState34Action,
	{_State5, FuncToken}:                                 _GotoState21Action,
	{_State5, VarToken}:                                  _GotoState37Action,
	{_State5, LetToken}:                                  _GotoState27Action,
	{_State5, NotToken}:                                  _GotoState30Action,
	{_State5, LabelDeclToken}:                            _GotoState25Action,
	{_State5, LparenToken}:                               _GotoState28Action,
	{_State5, LbracketToken}:                             _GotoState26Action,
	{_State5, SubToken}:                                  _GotoState35Action,
	{_State5, MulToken}:                                  _GotoState29Action,
	{_State5, BitNegToken}:                               _GotoState18Action,
	{_State5, BitAndToken}:                               _GotoState17Action,
	{_State5, GreaterToken}:                              _GotoState22Action,
	{_State5, ParseErrorToken}:                           _GotoState31Action,
	{_State5, VarDeclPatternType}:                        _GotoState69Action,
	{_State5, VarOrLetType}:                              _GotoState70Action,
	{_State5, ExpressionType}:                            _GotoState10Action,
	{_State5, OptionalLabelDeclType}:                     _GotoState60Action,
	{_State5, SequenceExprType}:                          _GotoState68Action,
	{_State5, CallExprType}:                              _GotoState50Action,
	{_State5, AtomExprType}:                              _GotoState43Action,
	{_State5, ParseErrorExprType}:                        _GotoState62Action,
	{_State5, LiteralExprType}:                           _GotoState58Action,
	{_State5, IdentifierExprType}:                        _GotoState53Action,
	{_State5, BlockExprType}:                             _GotoState49Action,
	{_State5, InitializeExprType}:                        _GotoState57Action,
	{_State5, ImplicitStructExprType}:                    _GotoState54Action,
	{_State5, AccessibleExprType}:                        _GotoState39Action,
	{_State5, AccessExprType}:                            _GotoState38Action,
	{_State5, IndexExprType}:                             _GotoState55Action,
	{_State5, PostfixableExprType}:                       _GotoState64Action,
	{_State5, PostfixUnaryExprType}:                      _GotoState63Action,
	{_State5, PrefixableExprType}:                        _GotoState67Action,
	{_State5, PrefixUnaryExprType}:                       _GotoState65Action,
	{_State5, PrefixUnaryOpType}:                         _GotoState66Action,
	{_State5, MulExprType}:                               _GotoState59Action,
	{_State5, BinaryMulExprType}:                         _GotoState47Action,
	{_State5, AddExprType}:                               _GotoState40Action,
	{_State5, BinaryAddExprType}:                         _GotoState44Action,
	{_State5, CmpExprType}:                               _GotoState51Action,
	{_State5, BinaryCmpExprType}:                         _GotoState46Action,
	{_State5, AndExprType}:                               _GotoState41Action,
	{_State5, BinaryAndExprType}:                         _GotoState45Action,
	{_State5, OrExprType}:                                _GotoState61Action,
	{_State5, BinaryOrExprType}:                          _GotoState48Action,
	{_State5, InitializableTypeType}:                     _GotoState56Action,
	{_State5, ExplicitStructDefType}:                     _GotoState52Action,
	{_State5, AnonymousFuncExprType}:                     _GotoState42Action,
	{_State13, CommentGroupsToken}:                       _GotoState71Action,
	{_State13, PackageToken}:                             _GotoState14Action,
	{_State13, TypeToken}:                                _GotoState15Action,
	{_State13, FuncToken}:                                _GotoState16Action,
	{_State13, VarToken}:                                 _GotoState37Action,
	{_State13, LetToken}:                                 _GotoState27Action,
	{_State13, LbraceToken}:                              _GotoState72Action,
	{_State13, DefinitionsType}:                          _GotoState74Action,
	{_State13, DefinitionType}:                           _GotoState73Action,
	{_State13, StatementBlockType}:                       _GotoState77Action,
	{_State13, VarDeclPatternType}:                       _GotoState79Action,
	{_State13, VarOrLetType}:                             _GotoState70Action,
	{_State13, TypeDefType}:                              _GotoState78Action,
	{_State13, NamedFuncDefType}:                         _GotoState75Action,
	{_State13, PackageDefType}:                           _GotoState76Action,
	{_State14, LbraceToken}:                              _GotoState72Action,
	{_State14, StatementBlockType}:                       _GotoState80Action,
	{_State15, IdentifierToken}:                          _GotoState81Action,
	{_State16, IdentifierToken}:                          _GotoState82Action,
	{_State16, LparenToken}:                              _GotoState83Action,
	{_State21, LparenToken}:                              _GotoState84Action,
	{_State22, IntegerLiteralToken}:                      _GotoState24Action,
	{_State22, FloatLiteralToken}:                        _GotoState20Action,
	{_State22, RuneLiteralToken}:                         _GotoState32Action,
	{_State22, StringLiteralToken}:                       _GotoState33Action,
	{_State22, IdentifierToken}:                          _GotoState23Action,
	{_State22, TrueToken}:                                _GotoState36Action,
	{_State22, FalseToken}:                               _GotoState19Action,
	{_State22, StructToken}:                              _GotoState34Action,
	{_State22, FuncToken}:                                _GotoState21Action,
	{_State22, VarToken}:                                 _GotoState37Action,
	{_State22, LetToken}:                                 _GotoState27Action,
	{_State22, NotToken}:                                 _GotoState30Action,
	{_State22, LabelDeclToken}:                           _GotoState25Action,
	{_State22, LparenToken}:                              _GotoState28Action,
	{_State22, LbracketToken}:                            _GotoState26Action,
	{_State22, SubToken}:                                 _GotoState35Action,
	{_State22, MulToken}:                                 _GotoState29Action,
	{_State22, BitNegToken}:                              _GotoState18Action,
	{_State22, BitAndToken}:                              _GotoState17Action,
	{_State22, GreaterToken}:                             _GotoState22Action,
	{_State22, ParseErrorToken}:                          _GotoState31Action,
	{_State22, VarDeclPatternType}:                       _GotoState69Action,
	{_State22, VarOrLetType}:                             _GotoState70Action,
	{_State22, OptionalLabelDeclType}:                    _GotoState85Action,
	{_State22, SequenceExprType}:                         _GotoState86Action,
	{_State22, CallExprType}:                             _GotoState50Action,
	{_State22, AtomExprType}:                             _GotoState43Action,
	{_State22, ParseErrorExprType}:                       _GotoState62Action,
	{_State22, LiteralExprType}:                          _GotoState58Action,
	{_State22, IdentifierExprType}:                       _GotoState53Action,
	{_State22, BlockExprType}:                            _GotoState49Action,
	{_State22, InitializeExprType}:                       _GotoState57Action,
	{_State22, ImplicitStructExprType}:                   _GotoState54Action,
	{_State22, AccessibleExprType}:                       _GotoState39Action,
	{_State22, AccessExprType}:                           _GotoState38Action,
	{_State22, IndexExprType}:                            _GotoState55Action,
	{_State22, PostfixableExprType}:                      _GotoState64Action,
	{_State22, PostfixUnaryExprType}:                     _GotoState63Action,
	{_State22, PrefixableExprType}:                       _GotoState67Action,
	{_State22, PrefixUnaryExprType}:                      _GotoState65Action,
	{_State22, PrefixUnaryOpType}:                        _GotoState66Action,
	{_State22, MulExprType}:                              _GotoState59Action,
	{_State22, BinaryMulExprType}:                        _GotoState47Action,
	{_State22, AddExprType}:                              _GotoState40Action,
	{_State22, BinaryAddExprType}:                        _GotoState44Action,
	{_State22, CmpExprType}:                              _GotoState51Action,
	{_State22, BinaryCmpExprType}:                        _GotoState46Action,
	{_State22, AndExprType}:                              _GotoState41Action,
	{_State22, BinaryAndExprType}:                        _GotoState45Action,
	{_State22, OrExprType}:                               _GotoState61Action,
	{_State22, BinaryOrExprType}:                         _GotoState48Action,
	{_State22, InitializableTypeType}:                    _GotoState56Action,
	{_State22, ExplicitStructDefType}:                    _GotoState52Action,
	{_State22, AnonymousFuncExprType}:                    _GotoState42Action,
	{_State26, IdentifierToken}:                          _GotoState93Action,
	{_State26, StructToken}:                              _GotoState34Action,
	{_State26, EnumToken}:                                _GotoState90Action,
	{_State26, TraitToken}:                               _GotoState98Action,
	{_State26, FuncToken}:                                _GotoState92Action,
	{_State26, LparenToken}:                              _GotoState94Action,
	{_State26, LbracketToken}:                            _GotoState26Action,
	{_State26, DotToken}:                                 _GotoState89Action,
	{_State26, QuestionToken}:                            _GotoState96Action,
	{_State26, ExclaimToken}:                             _GotoState91Action,
	{_State26, TildeTildeToken}:                          _GotoState97Action,
	{_State26, BitNegToken}:                              _GotoState88Action,
	{_State26, BitAndToken}:                              _GotoState87Action,
	{_State26, ParseErrorToken}:                          _GotoState95Action,
	{_State26, InitializableTypeType}:                    _GotoState104Action,
	{_State26, AtomTypeType}:                             _GotoState99Action,
	{_State26, ParseErrorTypeType}:                       _GotoState105Action,
	{_State26, ReturnableTypeType}:                       _GotoState108Action,
	{_State26, PrefixedTypeType}:                         _GotoState107Action,
	{_State26, PrefixTypeOpType}:                         _GotoState106Action,
	{_State26, ValueTypeType}:                            _GotoState111Action,
	{_State26, TraitOpTypeType}:                          _GotoState110Action,
	{_State26, ImplicitStructDefType}:                    _GotoState103Action,
	{_State26, ExplicitStructDefType}:                    _GotoState52Action,
	{_State26, ImplicitEnumDefType}:                      _GotoState102Action,
	{_State26, ExplicitEnumDefType}:                      _GotoState100Action,
	{_State26, TraitDefType}:                             _GotoState109Action,
	{_State26, FuncTypeType}:                             _GotoState101Action,
	{_State28, IntegerLiteralToken}:                      _GotoState24Action,
	{_State28, FloatLiteralToken}:                        _GotoState20Action,
	{_State28, RuneLiteralToken}:                         _GotoState32Action,
	{_State28, StringLiteralToken}:                       _GotoState33Action,
	{_State28, IdentifierToken}:                          _GotoState113Action,
	{_State28, TrueToken}:                                _GotoState36Action,
	{_State28, FalseToken}:                               _GotoState19Action,
	{_State28, StructToken}:                              _GotoState34Action,
	{_State28, FuncToken}:                                _GotoState21Action,
	{_State28, VarToken}:                                 _GotoState37Action,
	{_State28, LetToken}:                                 _GotoState27Action,
	{_State28, NotToken}:                                 _GotoState30Action,
	{_State28, LabelDeclToken}:                           _GotoState25Action,
	{_State28, LparenToken}:                              _GotoState28Action,
	{_State28, LbracketToken}:                            _GotoState26Action,
	{_State28, DotDotDotToken}:                           _GotoState112Action,
	{_State28, SubToken}:                                 _GotoState35Action,
	{_State28, MulToken}:                                 _GotoState29Action,
	{_State28, BitNegToken}:                              _GotoState18Action,
	{_State28, BitAndToken}:                              _GotoState17Action,
	{_State28, GreaterToken}:                             _GotoState22Action,
	{_State28, ParseErrorToken}:                          _GotoState31Action,
	{_State28, VarDeclPatternType}:                       _GotoState69Action,
	{_State28, VarOrLetType}:                             _GotoState70Action,
	{_State28, ExpressionType}:                           _GotoState117Action,
	{_State28, OptionalLabelDeclType}:                    _GotoState60Action,
	{_State28, SequenceExprType}:                         _GotoState68Action,
	{_State28, CallExprType}:                             _GotoState50Action,
	{_State28, OptionalArgumentsType}:                    _GotoState118Action,
	{_State28, ArgumentsType}:                            _GotoState115Action,
	{_State28, ArgumentType}:                             _GotoState114Action,
	{_State28, ColonExpressionsType}:                     _GotoState116Action,
	{_State28, OptionalExpressionType}:                   _GotoState119Action,
	{_State28, AtomExprType}:                             _GotoState43Action,
	{_State28, ParseErrorExprType}:                       _GotoState62Action,
	{_State28, LiteralExprType}:                          _GotoState58Action,
	{_State28, IdentifierExprType}:                       _GotoState53Action,
	{_State28, BlockExprType}:                            _GotoState49Action,
	{_State28, InitializeExprType}:                       _GotoState57Action,
	{_State28, ImplicitStructExprType}:                   _GotoState54Action,
	{_State28, AccessibleExprType}:                       _GotoState39Action,
	{_State28, AccessExprType}:                           _GotoState38Action,
	{_State28, IndexExprType}:                            _GotoState55Action,
	{_State28, PostfixableExprType}:                      _GotoState64Action,
	{_State28, PostfixUnaryExprType}:                     _GotoState63Action,
	{_State28, PrefixableExprType}:                       _GotoState67Action,
	{_State28, PrefixUnaryExprType}:                      _GotoState65Action,
	{_State28, PrefixUnaryOpType}:                        _GotoState66Action,
	{_State28, MulExprType}:                              _GotoState59Action,
	{_State28, BinaryMulExprType}:                        _GotoState47Action,
	{_State28, AddExprType}:                              _GotoState40Action,
	{_State28, BinaryAddExprType}:                        _GotoState44Action,
	{_State28, CmpExprType}:                              _GotoState51Action,
	{_State28, BinaryCmpExprType}:                        _GotoState46Action,
	{_State28, AndExprType}:                              _GotoState41Action,
	{_State28, BinaryAndExprType}:                        _GotoState45Action,
	{_State28, OrExprType}:                               _GotoState61Action,
	{_State28, BinaryOrExprType}:                         _GotoState48Action,
	{_State28, InitializableTypeType}:                    _GotoState56Action,
	{_State28, ExplicitStructDefType}:                    _GotoState52Action,
	{_State28, AnonymousFuncExprType}:                    _GotoState42Action,
	{_State34, LparenToken}:                              _GotoState120Action,
	{_State39, LbracketToken}:                            _GotoState123Action,
	{_State39, DotToken}:                                 _GotoState122Action,
	{_State39, QuestionToken}:                            _GotoState124Action,
	{_State39, DollarLbracketToken}:                      _GotoState121Action,
	{_State39, OptionalGenericBindingType}:               _GotoState125Action,
	{_State40, AddToken}:                                 _GotoState126Action,
	{_State40, SubToken}:                                 _GotoState129Action,
	{_State40, BitXorToken}:                              _GotoState128Action,
	{_State40, BitOrToken}:                               _GotoState127Action,
	{_State40, AddOpType}:                                _GotoState130Action,
	{_State41, AndToken}:                                 _GotoState131Action,
	{_State51, EqualToken}:                               _GotoState132Action,
	{_State51, NotEqualToken}:                            _GotoState137Action,
	{_State51, LessToken}:                                _GotoState135Action,
	{_State51, LessOrEqualToken}:                         _GotoState136Action,
	{_State51, GreaterToken}:                             _GotoState133Action,
	{_State51, GreaterOrEqualToken}:                      _GotoState134Action,
	{_State51, CmpOpType}:                                _GotoState138Action,
	{_State56, LparenToken}:                              _GotoState139Action,
	{_State59, MulToken}:                                 _GotoState145Action,
	{_State59, DivToken}:                                 _GotoState143Action,
	{_State59, ModToken}:                                 _GotoState144Action,
	{_State59, BitAndToken}:                              _GotoState140Action,
	{_State59, BitLshiftToken}:                           _GotoState141Action,
	{_State59, BitRshiftToken}:                           _GotoState142Action,
	{_State59, MulOpType}:                                _GotoState146Action,
	{_State60, IfToken}:                                  _GotoState149Action,
	{_State60, SwitchToken}:                              _GotoState150Action,
	{_State60, ForToken}:                                 _GotoState148Action,
	{_State60, DoToken}:                                  _GotoState147Action,
	{_State60, LbraceToken}:                              _GotoState72Action,
	{_State60, StatementBlockType}:                       _GotoState153Action,
	{_State60, IfExprType}:                               _GotoState151Action,
	{_State60, SwitchExprType}:                           _GotoState154Action,
	{_State60, LoopExprType}:                             _GotoState152Action,
	{_State61, OrToken}:                                  _GotoState155Action,
	{_State66, IntegerLiteralToken}:                      _GotoState24Action,
	{_State66, FloatLiteralToken}:                        _GotoState20Action,
	{_State66, RuneLiteralToken}:                         _GotoState32Action,
	{_State66, StringLiteralToken}:                       _GotoState33Action,
	{_State66, IdentifierToken}:                          _GotoState23Action,
	{_State66, TrueToken}:                                _GotoState36Action,
	{_State66, FalseToken}:                               _GotoState19Action,
	{_State66, StructToken}:                              _GotoState34Action,
	{_State66, FuncToken}:                                _GotoState21Action,
	{_State66, NotToken}:                                 _GotoState30Action,
	{_State66, LabelDeclToken}:                           _GotoState25Action,
	{_State66, LparenToken}:                              _GotoState28Action,
	{_State66, LbracketToken}:                            _GotoState26Action,
	{_State66, SubToken}:                                 _GotoState35Action,
	{_State66, MulToken}:                                 _GotoState29Action,
	{_State66, BitNegToken}:                              _GotoState18Action,
	{_State66, BitAndToken}:                              _GotoState17Action,
	{_State66, ParseErrorToken}:                          _GotoState31Action,
	{_State66, OptionalLabelDeclType}:                    _GotoState85Action,
	{_State66, CallExprType}:                             _GotoState50Action,
	{_State66, AtomExprType}:                             _GotoState43Action,
	{_State66, ParseErrorExprType}:                       _GotoState62Action,
	{_State66, LiteralExprType}:                          _GotoState58Action,
	{_State66, IdentifierExprType}:                       _GotoState53Action,
	{_State66, BlockExprType}:                            _GotoState49Action,
	{_State66, InitializeExprType}:                       _GotoState57Action,
	{_State66, ImplicitStructExprType}:                   _GotoState54Action,
	{_State66, AccessibleExprType}:                       _GotoState39Action,
	{_State66, AccessExprType}:                           _GotoState38Action,
	{_State66, IndexExprType}:                            _GotoState55Action,
	{_State66, PostfixableExprType}:                      _GotoState64Action,
	{_State66, PostfixUnaryExprType}:                     _GotoState63Action,
	{_State66, PrefixableExprType}:                       _GotoState156Action,
	{_State66, PrefixUnaryExprType}:                      _GotoState65Action,
	{_State66, PrefixUnaryOpType}:                        _GotoState66Action,
	{_State66, InitializableTypeType}:                    _GotoState56Action,
	{_State66, ExplicitStructDefType}:                    _GotoState52Action,
	{_State66, AnonymousFuncExprType}:                    _GotoState42Action,
	{_State70, IdentifierToken}:                          _GotoState157Action,
	{_State70, LparenToken}:                              _GotoState158Action,
	{_State70, VarPatternType}:                           _GotoState160Action,
	{_State70, TuplePatternType}:                         _GotoState159Action,
	{_State72, StatementsType}:                           _GotoState161Action,
	{_State74, NewlinesToken}:                            _GotoState162Action,
	{_State74, OptionalNewlinesType}:                     _GotoState163Action,
	{_State79, AssignToken}:                              _GotoState164Action,
	{_State81, DollarLbracketToken}:                      _GotoState166Action,
	{_State81, AssignToken}:                              _GotoState165Action,
	{_State81, OptionalGenericParametersType}:            _GotoState167Action,
	{_State82, DollarLbracketToken}:                      _GotoState166Action,
	{_State82, AssignToken}:                              _GotoState168Action,
	{_State82, OptionalGenericParametersType}:            _GotoState169Action,
	{_State83, IdentifierToken}:                          _GotoState170Action,
	{_State83, ParameterDefType}:                         _GotoState171Action,
	{_State84, IdentifierToken}:                          _GotoState170Action,
	{_State84, ParameterDefType}:                         _GotoState173Action,
	{_State84, ParameterDefsType}:                        _GotoState174Action,
	{_State84, OptionalParameterDefsType}:                _GotoState172Action,
	{_State85, LbraceToken}:                              _GotoState72Action,
	{_State85, StatementBlockType}:                       _GotoState153Action,
	{_State89, DollarLbracketToken}:                      _GotoState121Action,
	{_State89, OptionalGenericBindingType}:               _GotoState175Action,
	{_State90, LparenToken}:                              _GotoState176Action,
	{_State92, LparenToken}:                              _GotoState177Action,
	{_State93, DotToken}:                                 _GotoState178Action,
	{_State93, DollarLbracketToken}:                      _GotoState121Action,
	{_State93, OptionalGenericBindingType}:               _GotoState179Action,
	{_State94, IdentifierToken}:                          _GotoState180Action,
	{_State94, UnsafeToken}:                              _GotoState181Action,
	{_State94, StructToken}:                              _GotoState34Action,
	{_State94, EnumToken}:                                _GotoState90Action,
	{_State94, TraitToken}:                               _GotoState98Action,
	{_State94, FuncToken}:                                _GotoState92Action,
	{_State94, LparenToken}:                              _GotoState94Action,
	{_State94, LbracketToken}:                            _GotoState26Action,
	{_State94, DotToken}:                                 _GotoState89Action,
	{_State94, QuestionToken}:                            _GotoState96Action,
	{_State94, ExclaimToken}:                             _GotoState91Action,
	{_State94, TildeTildeToken}:                          _GotoState97Action,
	{_State94, BitNegToken}:                              _GotoState88Action,
	{_State94, BitAndToken}:                              _GotoState87Action,
	{_State94, ParseErrorToken}:                          _GotoState95Action,
	{_State94, UnsafeStatementType}:                      _GotoState187Action,
	{_State94, InitializableTypeType}:                    _GotoState104Action,
	{_State94, AtomTypeType}:                             _GotoState99Action,
	{_State94, ParseErrorTypeType}:                       _GotoState105Action,
	{_State94, ReturnableTypeType}:                       _GotoState108Action,
	{_State94, PrefixedTypeType}:                         _GotoState107Action,
	{_State94, PrefixTypeOpType}:                         _GotoState106Action,
	{_State94, ValueTypeType}:                            _GotoState188Action,
	{_State94, TraitOpTypeType}:                          _GotoState110Action,
	{_State94, FieldDefType}:                             _GotoState183Action,
	{_State94, ImplicitFieldDefsType}:                    _GotoState185Action,
	{_State94, OptionalImplicitFieldDefsType}:            _GotoState186Action,
	{_State94, ImplicitStructDefType}:                    _GotoState103Action,
	{_State94, ExplicitStructDefType}:                    _GotoState52Action,
	{_State94, EnumValueDefType}:                         _GotoState182Action,
	{_State94, ImplicitEnumValueDefsType}:                _GotoState184Action,
	{_State94, ImplicitEnumDefType}:                      _GotoState102Action,
	{_State94, ExplicitEnumDefType}:                      _GotoState100Action,
	{_State94, TraitDefType}:                             _GotoState109Action,
	{_State94, FuncTypeType}:                             _GotoState101Action,
	{_State98, LparenToken}:                              _GotoState189Action,
	{_State106, IdentifierToken}:                         _GotoState93Action,
	{_State106, StructToken}:                             _GotoState34Action,
	{_State106, EnumToken}:                               _GotoState90Action,
	{_State106, TraitToken}:                              _GotoState98Action,
	{_State106, FuncToken}:                               _GotoState92Action,
	{_State106, LparenToken}:                             _GotoState94Action,
	{_State106, LbracketToken}:                           _GotoState26Action,
	{_State106, DotToken}:                                _GotoState89Action,
	{_State106, QuestionToken}:                           _GotoState96Action,
	{_State106, ExclaimToken}:                            _GotoState91Action,
	{_State106, TildeTildeToken}:                         _GotoState97Action,
	{_State106, BitNegToken}:                             _GotoState88Action,
	{_State106, BitAndToken}:                             _GotoState87Action,
	{_State106, ParseErrorToken}:                         _GotoState95Action,
	{_State106, InitializableTypeType}:                   _GotoState104Action,
	{_State106, AtomTypeType}:                            _GotoState99Action,
	{_State106, ParseErrorTypeType}:                      _GotoState105Action,
	{_State106, ReturnableTypeType}:                      _GotoState190Action,
	{_State106, PrefixedTypeType}:                        _GotoState107Action,
	{_State106, PrefixTypeOpType}:                        _GotoState106Action,
	{_State106, ImplicitStructDefType}:                   _GotoState103Action,
	{_State106, ExplicitStructDefType}:                   _GotoState52Action,
	{_State106, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State106, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State106, TraitDefType}:                            _GotoState109Action,
	{_State106, FuncTypeType}:                            _GotoState101Action,
	{_State111, RbracketToken}:                           _GotoState195Action,
	{_State111, CommaToken}:                              _GotoState193Action,
	{_State111, ColonToken}:                              _GotoState192Action,
	{_State111, AddToken}:                                _GotoState191Action,
	{_State111, SubToken}:                                _GotoState196Action,
	{_State111, MulToken}:                                _GotoState194Action,
	{_State111, TraitOpType}:                             _GotoState197Action,
	{_State113, AssignToken}:                             _GotoState198Action,
	{_State115, CommaToken}:                              _GotoState199Action,
	{_State116, ColonToken}:                              _GotoState200Action,
	{_State118, RparenToken}:                             _GotoState201Action,
	{_State119, ColonToken}:                              _GotoState202Action,
	{_State120, IdentifierToken}:                         _GotoState180Action,
	{_State120, UnsafeToken}:                             _GotoState181Action,
	{_State120, StructToken}:                             _GotoState34Action,
	{_State120, EnumToken}:                               _GotoState90Action,
	{_State120, TraitToken}:                              _GotoState98Action,
	{_State120, FuncToken}:                               _GotoState92Action,
	{_State120, LparenToken}:                             _GotoState94Action,
	{_State120, LbracketToken}:                           _GotoState26Action,
	{_State120, DotToken}:                                _GotoState89Action,
	{_State120, QuestionToken}:                           _GotoState96Action,
	{_State120, ExclaimToken}:                            _GotoState91Action,
	{_State120, TildeTildeToken}:                         _GotoState97Action,
	{_State120, BitNegToken}:                             _GotoState88Action,
	{_State120, BitAndToken}:                             _GotoState87Action,
	{_State120, ParseErrorToken}:                         _GotoState95Action,
	{_State120, UnsafeStatementType}:                     _GotoState187Action,
	{_State120, InitializableTypeType}:                   _GotoState104Action,
	{_State120, AtomTypeType}:                            _GotoState99Action,
	{_State120, ParseErrorTypeType}:                      _GotoState105Action,
	{_State120, ReturnableTypeType}:                      _GotoState108Action,
	{_State120, PrefixedTypeType}:                        _GotoState107Action,
	{_State120, PrefixTypeOpType}:                        _GotoState106Action,
	{_State120, ValueTypeType}:                           _GotoState188Action,
	{_State120, TraitOpTypeType}:                         _GotoState110Action,
	{_State120, FieldDefType}:                            _GotoState204Action,
	{_State120, ImplicitStructDefType}:                   _GotoState103Action,
	{_State120, ExplicitFieldDefsType}:                   _GotoState203Action,
	{_State120, OptionalExplicitFieldDefsType}:           _GotoState205Action,
	{_State120, ExplicitStructDefType}:                   _GotoState52Action,
	{_State120, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State120, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State120, TraitDefType}:                            _GotoState109Action,
	{_State120, FuncTypeType}:                            _GotoState101Action,
	{_State121, IdentifierToken}:                         _GotoState93Action,
	{_State121, StructToken}:                             _GotoState34Action,
	{_State121, EnumToken}:                               _GotoState90Action,
	{_State121, TraitToken}:                              _GotoState98Action,
	{_State121, FuncToken}:                               _GotoState92Action,
	{_State121, LparenToken}:                             _GotoState94Action,
	{_State121, LbracketToken}:                           _GotoState26Action,
	{_State121, DotToken}:                                _GotoState89Action,
	{_State121, QuestionToken}:                           _GotoState96Action,
	{_State121, ExclaimToken}:                            _GotoState91Action,
	{_State121, TildeTildeToken}:                         _GotoState97Action,
	{_State121, BitNegToken}:                             _GotoState88Action,
	{_State121, BitAndToken}:                             _GotoState87Action,
	{_State121, ParseErrorToken}:                         _GotoState95Action,
	{_State121, OptionalGenericArgumentsType}:            _GotoState207Action,
	{_State121, GenericArgumentsType}:                    _GotoState206Action,
	{_State121, InitializableTypeType}:                   _GotoState104Action,
	{_State121, AtomTypeType}:                            _GotoState99Action,
	{_State121, ParseErrorTypeType}:                      _GotoState105Action,
	{_State121, ReturnableTypeType}:                      _GotoState108Action,
	{_State121, PrefixedTypeType}:                        _GotoState107Action,
	{_State121, PrefixTypeOpType}:                        _GotoState106Action,
	{_State121, ValueTypeType}:                           _GotoState208Action,
	{_State121, TraitOpTypeType}:                         _GotoState110Action,
	{_State121, ImplicitStructDefType}:                   _GotoState103Action,
	{_State121, ExplicitStructDefType}:                   _GotoState52Action,
	{_State121, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State121, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State121, TraitDefType}:                            _GotoState109Action,
	{_State121, FuncTypeType}:                            _GotoState101Action,
	{_State122, IdentifierToken}:                         _GotoState209Action,
	{_State123, IntegerLiteralToken}:                     _GotoState24Action,
	{_State123, FloatLiteralToken}:                       _GotoState20Action,
	{_State123, RuneLiteralToken}:                        _GotoState32Action,
	{_State123, StringLiteralToken}:                      _GotoState33Action,
	{_State123, IdentifierToken}:                         _GotoState113Action,
	{_State123, TrueToken}:                               _GotoState36Action,
	{_State123, FalseToken}:                              _GotoState19Action,
	{_State123, StructToken}:                             _GotoState34Action,
	{_State123, FuncToken}:                               _GotoState21Action,
	{_State123, VarToken}:                                _GotoState37Action,
	{_State123, LetToken}:                                _GotoState27Action,
	{_State123, NotToken}:                                _GotoState30Action,
	{_State123, LabelDeclToken}:                          _GotoState25Action,
	{_State123, LparenToken}:                             _GotoState28Action,
	{_State123, LbracketToken}:                           _GotoState26Action,
	{_State123, DotDotDotToken}:                          _GotoState112Action,
	{_State123, SubToken}:                                _GotoState35Action,
	{_State123, MulToken}:                                _GotoState29Action,
	{_State123, BitNegToken}:                             _GotoState18Action,
	{_State123, BitAndToken}:                             _GotoState17Action,
	{_State123, GreaterToken}:                            _GotoState22Action,
	{_State123, ParseErrorToken}:                         _GotoState31Action,
	{_State123, VarDeclPatternType}:                      _GotoState69Action,
	{_State123, VarOrLetType}:                            _GotoState70Action,
	{_State123, ExpressionType}:                          _GotoState117Action,
	{_State123, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State123, SequenceExprType}:                        _GotoState68Action,
	{_State123, CallExprType}:                            _GotoState50Action,
	{_State123, ArgumentType}:                            _GotoState210Action,
	{_State123, ColonExpressionsType}:                    _GotoState116Action,
	{_State123, OptionalExpressionType}:                  _GotoState119Action,
	{_State123, AtomExprType}:                            _GotoState43Action,
	{_State123, ParseErrorExprType}:                      _GotoState62Action,
	{_State123, LiteralExprType}:                         _GotoState58Action,
	{_State123, IdentifierExprType}:                      _GotoState53Action,
	{_State123, BlockExprType}:                           _GotoState49Action,
	{_State123, InitializeExprType}:                      _GotoState57Action,
	{_State123, ImplicitStructExprType}:                  _GotoState54Action,
	{_State123, AccessibleExprType}:                      _GotoState39Action,
	{_State123, AccessExprType}:                          _GotoState38Action,
	{_State123, IndexExprType}:                           _GotoState55Action,
	{_State123, PostfixableExprType}:                     _GotoState64Action,
	{_State123, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State123, PrefixableExprType}:                      _GotoState67Action,
	{_State123, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State123, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State123, MulExprType}:                             _GotoState59Action,
	{_State123, BinaryMulExprType}:                       _GotoState47Action,
	{_State123, AddExprType}:                             _GotoState40Action,
	{_State123, BinaryAddExprType}:                       _GotoState44Action,
	{_State123, CmpExprType}:                             _GotoState51Action,
	{_State123, BinaryCmpExprType}:                       _GotoState46Action,
	{_State123, AndExprType}:                             _GotoState41Action,
	{_State123, BinaryAndExprType}:                       _GotoState45Action,
	{_State123, OrExprType}:                              _GotoState61Action,
	{_State123, BinaryOrExprType}:                        _GotoState48Action,
	{_State123, InitializableTypeType}:                   _GotoState56Action,
	{_State123, ExplicitStructDefType}:                   _GotoState52Action,
	{_State123, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State125, LparenToken}:                             _GotoState211Action,
	{_State130, IntegerLiteralToken}:                     _GotoState24Action,
	{_State130, FloatLiteralToken}:                       _GotoState20Action,
	{_State130, RuneLiteralToken}:                        _GotoState32Action,
	{_State130, StringLiteralToken}:                      _GotoState33Action,
	{_State130, IdentifierToken}:                         _GotoState23Action,
	{_State130, TrueToken}:                               _GotoState36Action,
	{_State130, FalseToken}:                              _GotoState19Action,
	{_State130, StructToken}:                             _GotoState34Action,
	{_State130, FuncToken}:                               _GotoState21Action,
	{_State130, NotToken}:                                _GotoState30Action,
	{_State130, LabelDeclToken}:                          _GotoState25Action,
	{_State130, LparenToken}:                             _GotoState28Action,
	{_State130, LbracketToken}:                           _GotoState26Action,
	{_State130, SubToken}:                                _GotoState35Action,
	{_State130, MulToken}:                                _GotoState29Action,
	{_State130, BitNegToken}:                             _GotoState18Action,
	{_State130, BitAndToken}:                             _GotoState17Action,
	{_State130, ParseErrorToken}:                         _GotoState31Action,
	{_State130, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State130, CallExprType}:                            _GotoState50Action,
	{_State130, AtomExprType}:                            _GotoState43Action,
	{_State130, ParseErrorExprType}:                      _GotoState62Action,
	{_State130, LiteralExprType}:                         _GotoState58Action,
	{_State130, IdentifierExprType}:                      _GotoState53Action,
	{_State130, BlockExprType}:                           _GotoState49Action,
	{_State130, InitializeExprType}:                      _GotoState57Action,
	{_State130, ImplicitStructExprType}:                  _GotoState54Action,
	{_State130, AccessibleExprType}:                      _GotoState39Action,
	{_State130, AccessExprType}:                          _GotoState38Action,
	{_State130, IndexExprType}:                           _GotoState55Action,
	{_State130, PostfixableExprType}:                     _GotoState64Action,
	{_State130, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State130, PrefixableExprType}:                      _GotoState67Action,
	{_State130, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State130, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State130, MulExprType}:                             _GotoState212Action,
	{_State130, BinaryMulExprType}:                       _GotoState47Action,
	{_State130, InitializableTypeType}:                   _GotoState56Action,
	{_State130, ExplicitStructDefType}:                   _GotoState52Action,
	{_State130, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State131, IntegerLiteralToken}:                     _GotoState24Action,
	{_State131, FloatLiteralToken}:                       _GotoState20Action,
	{_State131, RuneLiteralToken}:                        _GotoState32Action,
	{_State131, StringLiteralToken}:                      _GotoState33Action,
	{_State131, IdentifierToken}:                         _GotoState23Action,
	{_State131, TrueToken}:                               _GotoState36Action,
	{_State131, FalseToken}:                              _GotoState19Action,
	{_State131, StructToken}:                             _GotoState34Action,
	{_State131, FuncToken}:                               _GotoState21Action,
	{_State131, NotToken}:                                _GotoState30Action,
	{_State131, LabelDeclToken}:                          _GotoState25Action,
	{_State131, LparenToken}:                             _GotoState28Action,
	{_State131, LbracketToken}:                           _GotoState26Action,
	{_State131, SubToken}:                                _GotoState35Action,
	{_State131, MulToken}:                                _GotoState29Action,
	{_State131, BitNegToken}:                             _GotoState18Action,
	{_State131, BitAndToken}:                             _GotoState17Action,
	{_State131, ParseErrorToken}:                         _GotoState31Action,
	{_State131, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State131, CallExprType}:                            _GotoState50Action,
	{_State131, AtomExprType}:                            _GotoState43Action,
	{_State131, ParseErrorExprType}:                      _GotoState62Action,
	{_State131, LiteralExprType}:                         _GotoState58Action,
	{_State131, IdentifierExprType}:                      _GotoState53Action,
	{_State131, BlockExprType}:                           _GotoState49Action,
	{_State131, InitializeExprType}:                      _GotoState57Action,
	{_State131, ImplicitStructExprType}:                  _GotoState54Action,
	{_State131, AccessibleExprType}:                      _GotoState39Action,
	{_State131, AccessExprType}:                          _GotoState38Action,
	{_State131, IndexExprType}:                           _GotoState55Action,
	{_State131, PostfixableExprType}:                     _GotoState64Action,
	{_State131, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State131, PrefixableExprType}:                      _GotoState67Action,
	{_State131, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State131, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State131, MulExprType}:                             _GotoState59Action,
	{_State131, BinaryMulExprType}:                       _GotoState47Action,
	{_State131, AddExprType}:                             _GotoState40Action,
	{_State131, BinaryAddExprType}:                       _GotoState44Action,
	{_State131, CmpExprType}:                             _GotoState213Action,
	{_State131, BinaryCmpExprType}:                       _GotoState46Action,
	{_State131, InitializableTypeType}:                   _GotoState56Action,
	{_State131, ExplicitStructDefType}:                   _GotoState52Action,
	{_State131, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State138, IntegerLiteralToken}:                     _GotoState24Action,
	{_State138, FloatLiteralToken}:                       _GotoState20Action,
	{_State138, RuneLiteralToken}:                        _GotoState32Action,
	{_State138, StringLiteralToken}:                      _GotoState33Action,
	{_State138, IdentifierToken}:                         _GotoState23Action,
	{_State138, TrueToken}:                               _GotoState36Action,
	{_State138, FalseToken}:                              _GotoState19Action,
	{_State138, StructToken}:                             _GotoState34Action,
	{_State138, FuncToken}:                               _GotoState21Action,
	{_State138, NotToken}:                                _GotoState30Action,
	{_State138, LabelDeclToken}:                          _GotoState25Action,
	{_State138, LparenToken}:                             _GotoState28Action,
	{_State138, LbracketToken}:                           _GotoState26Action,
	{_State138, SubToken}:                                _GotoState35Action,
	{_State138, MulToken}:                                _GotoState29Action,
	{_State138, BitNegToken}:                             _GotoState18Action,
	{_State138, BitAndToken}:                             _GotoState17Action,
	{_State138, ParseErrorToken}:                         _GotoState31Action,
	{_State138, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State138, CallExprType}:                            _GotoState50Action,
	{_State138, AtomExprType}:                            _GotoState43Action,
	{_State138, ParseErrorExprType}:                      _GotoState62Action,
	{_State138, LiteralExprType}:                         _GotoState58Action,
	{_State138, IdentifierExprType}:                      _GotoState53Action,
	{_State138, BlockExprType}:                           _GotoState49Action,
	{_State138, InitializeExprType}:                      _GotoState57Action,
	{_State138, ImplicitStructExprType}:                  _GotoState54Action,
	{_State138, AccessibleExprType}:                      _GotoState39Action,
	{_State138, AccessExprType}:                          _GotoState38Action,
	{_State138, IndexExprType}:                           _GotoState55Action,
	{_State138, PostfixableExprType}:                     _GotoState64Action,
	{_State138, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State138, PrefixableExprType}:                      _GotoState67Action,
	{_State138, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State138, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State138, MulExprType}:                             _GotoState59Action,
	{_State138, BinaryMulExprType}:                       _GotoState47Action,
	{_State138, AddExprType}:                             _GotoState214Action,
	{_State138, BinaryAddExprType}:                       _GotoState44Action,
	{_State138, InitializableTypeType}:                   _GotoState56Action,
	{_State138, ExplicitStructDefType}:                   _GotoState52Action,
	{_State138, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State139, IntegerLiteralToken}:                     _GotoState24Action,
	{_State139, FloatLiteralToken}:                       _GotoState20Action,
	{_State139, RuneLiteralToken}:                        _GotoState32Action,
	{_State139, StringLiteralToken}:                      _GotoState33Action,
	{_State139, IdentifierToken}:                         _GotoState113Action,
	{_State139, TrueToken}:                               _GotoState36Action,
	{_State139, FalseToken}:                              _GotoState19Action,
	{_State139, StructToken}:                             _GotoState34Action,
	{_State139, FuncToken}:                               _GotoState21Action,
	{_State139, VarToken}:                                _GotoState37Action,
	{_State139, LetToken}:                                _GotoState27Action,
	{_State139, NotToken}:                                _GotoState30Action,
	{_State139, LabelDeclToken}:                          _GotoState25Action,
	{_State139, LparenToken}:                             _GotoState28Action,
	{_State139, LbracketToken}:                           _GotoState26Action,
	{_State139, DotDotDotToken}:                          _GotoState112Action,
	{_State139, SubToken}:                                _GotoState35Action,
	{_State139, MulToken}:                                _GotoState29Action,
	{_State139, BitNegToken}:                             _GotoState18Action,
	{_State139, BitAndToken}:                             _GotoState17Action,
	{_State139, GreaterToken}:                            _GotoState22Action,
	{_State139, ParseErrorToken}:                         _GotoState31Action,
	{_State139, VarDeclPatternType}:                      _GotoState69Action,
	{_State139, VarOrLetType}:                            _GotoState70Action,
	{_State139, ExpressionType}:                          _GotoState117Action,
	{_State139, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State139, SequenceExprType}:                        _GotoState68Action,
	{_State139, CallExprType}:                            _GotoState50Action,
	{_State139, OptionalArgumentsType}:                   _GotoState215Action,
	{_State139, ArgumentsType}:                           _GotoState115Action,
	{_State139, ArgumentType}:                            _GotoState114Action,
	{_State139, ColonExpressionsType}:                    _GotoState116Action,
	{_State139, OptionalExpressionType}:                  _GotoState119Action,
	{_State139, AtomExprType}:                            _GotoState43Action,
	{_State139, ParseErrorExprType}:                      _GotoState62Action,
	{_State139, LiteralExprType}:                         _GotoState58Action,
	{_State139, IdentifierExprType}:                      _GotoState53Action,
	{_State139, BlockExprType}:                           _GotoState49Action,
	{_State139, InitializeExprType}:                      _GotoState57Action,
	{_State139, ImplicitStructExprType}:                  _GotoState54Action,
	{_State139, AccessibleExprType}:                      _GotoState39Action,
	{_State139, AccessExprType}:                          _GotoState38Action,
	{_State139, IndexExprType}:                           _GotoState55Action,
	{_State139, PostfixableExprType}:                     _GotoState64Action,
	{_State139, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State139, PrefixableExprType}:                      _GotoState67Action,
	{_State139, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State139, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State139, MulExprType}:                             _GotoState59Action,
	{_State139, BinaryMulExprType}:                       _GotoState47Action,
	{_State139, AddExprType}:                             _GotoState40Action,
	{_State139, BinaryAddExprType}:                       _GotoState44Action,
	{_State139, CmpExprType}:                             _GotoState51Action,
	{_State139, BinaryCmpExprType}:                       _GotoState46Action,
	{_State139, AndExprType}:                             _GotoState41Action,
	{_State139, BinaryAndExprType}:                       _GotoState45Action,
	{_State139, OrExprType}:                              _GotoState61Action,
	{_State139, BinaryOrExprType}:                        _GotoState48Action,
	{_State139, InitializableTypeType}:                   _GotoState56Action,
	{_State139, ExplicitStructDefType}:                   _GotoState52Action,
	{_State139, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State146, IntegerLiteralToken}:                     _GotoState24Action,
	{_State146, FloatLiteralToken}:                       _GotoState20Action,
	{_State146, RuneLiteralToken}:                        _GotoState32Action,
	{_State146, StringLiteralToken}:                      _GotoState33Action,
	{_State146, IdentifierToken}:                         _GotoState23Action,
	{_State146, TrueToken}:                               _GotoState36Action,
	{_State146, FalseToken}:                              _GotoState19Action,
	{_State146, StructToken}:                             _GotoState34Action,
	{_State146, FuncToken}:                               _GotoState21Action,
	{_State146, NotToken}:                                _GotoState30Action,
	{_State146, LabelDeclToken}:                          _GotoState25Action,
	{_State146, LparenToken}:                             _GotoState28Action,
	{_State146, LbracketToken}:                           _GotoState26Action,
	{_State146, SubToken}:                                _GotoState35Action,
	{_State146, MulToken}:                                _GotoState29Action,
	{_State146, BitNegToken}:                             _GotoState18Action,
	{_State146, BitAndToken}:                             _GotoState17Action,
	{_State146, ParseErrorToken}:                         _GotoState31Action,
	{_State146, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State146, CallExprType}:                            _GotoState50Action,
	{_State146, AtomExprType}:                            _GotoState43Action,
	{_State146, ParseErrorExprType}:                      _GotoState62Action,
	{_State146, LiteralExprType}:                         _GotoState58Action,
	{_State146, IdentifierExprType}:                      _GotoState53Action,
	{_State146, BlockExprType}:                           _GotoState49Action,
	{_State146, InitializeExprType}:                      _GotoState57Action,
	{_State146, ImplicitStructExprType}:                  _GotoState54Action,
	{_State146, AccessibleExprType}:                      _GotoState39Action,
	{_State146, AccessExprType}:                          _GotoState38Action,
	{_State146, IndexExprType}:                           _GotoState55Action,
	{_State146, PostfixableExprType}:                     _GotoState64Action,
	{_State146, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State146, PrefixableExprType}:                      _GotoState216Action,
	{_State146, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State146, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State146, InitializableTypeType}:                   _GotoState56Action,
	{_State146, ExplicitStructDefType}:                   _GotoState52Action,
	{_State146, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State147, LbraceToken}:                             _GotoState72Action,
	{_State147, StatementBlockType}:                      _GotoState217Action,
	{_State148, IntegerLiteralToken}:                     _GotoState24Action,
	{_State148, FloatLiteralToken}:                       _GotoState20Action,
	{_State148, RuneLiteralToken}:                        _GotoState32Action,
	{_State148, StringLiteralToken}:                      _GotoState33Action,
	{_State148, IdentifierToken}:                         _GotoState23Action,
	{_State148, TrueToken}:                               _GotoState36Action,
	{_State148, FalseToken}:                              _GotoState19Action,
	{_State148, StructToken}:                             _GotoState34Action,
	{_State148, FuncToken}:                               _GotoState21Action,
	{_State148, VarToken}:                                _GotoState37Action,
	{_State148, LetToken}:                                _GotoState27Action,
	{_State148, NotToken}:                                _GotoState30Action,
	{_State148, LabelDeclToken}:                          _GotoState25Action,
	{_State148, LparenToken}:                             _GotoState28Action,
	{_State148, LbracketToken}:                           _GotoState26Action,
	{_State148, SubToken}:                                _GotoState35Action,
	{_State148, MulToken}:                                _GotoState29Action,
	{_State148, BitNegToken}:                             _GotoState18Action,
	{_State148, BitAndToken}:                             _GotoState17Action,
	{_State148, GreaterToken}:                            _GotoState22Action,
	{_State148, ParseErrorToken}:                         _GotoState31Action,
	{_State148, VarDeclPatternType}:                      _GotoState69Action,
	{_State148, VarOrLetType}:                            _GotoState70Action,
	{_State148, AssignPatternType}:                       _GotoState218Action,
	{_State148, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State148, SequenceExprType}:                        _GotoState220Action,
	{_State148, ForAssignmentType}:                       _GotoState219Action,
	{_State148, CallExprType}:                            _GotoState50Action,
	{_State148, AtomExprType}:                            _GotoState43Action,
	{_State148, ParseErrorExprType}:                      _GotoState62Action,
	{_State148, LiteralExprType}:                         _GotoState58Action,
	{_State148, IdentifierExprType}:                      _GotoState53Action,
	{_State148, BlockExprType}:                           _GotoState49Action,
	{_State148, InitializeExprType}:                      _GotoState57Action,
	{_State148, ImplicitStructExprType}:                  _GotoState54Action,
	{_State148, AccessibleExprType}:                      _GotoState39Action,
	{_State148, AccessExprType}:                          _GotoState38Action,
	{_State148, IndexExprType}:                           _GotoState55Action,
	{_State148, PostfixableExprType}:                     _GotoState64Action,
	{_State148, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State148, PrefixableExprType}:                      _GotoState67Action,
	{_State148, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State148, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State148, MulExprType}:                             _GotoState59Action,
	{_State148, BinaryMulExprType}:                       _GotoState47Action,
	{_State148, AddExprType}:                             _GotoState40Action,
	{_State148, BinaryAddExprType}:                       _GotoState44Action,
	{_State148, CmpExprType}:                             _GotoState51Action,
	{_State148, BinaryCmpExprType}:                       _GotoState46Action,
	{_State148, AndExprType}:                             _GotoState41Action,
	{_State148, BinaryAndExprType}:                       _GotoState45Action,
	{_State148, OrExprType}:                              _GotoState61Action,
	{_State148, BinaryOrExprType}:                        _GotoState48Action,
	{_State148, InitializableTypeType}:                   _GotoState56Action,
	{_State148, ExplicitStructDefType}:                   _GotoState52Action,
	{_State148, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State149, IntegerLiteralToken}:                     _GotoState24Action,
	{_State149, FloatLiteralToken}:                       _GotoState20Action,
	{_State149, RuneLiteralToken}:                        _GotoState32Action,
	{_State149, StringLiteralToken}:                      _GotoState33Action,
	{_State149, IdentifierToken}:                         _GotoState23Action,
	{_State149, TrueToken}:                               _GotoState36Action,
	{_State149, FalseToken}:                              _GotoState19Action,
	{_State149, CaseToken}:                               _GotoState221Action,
	{_State149, StructToken}:                             _GotoState34Action,
	{_State149, FuncToken}:                               _GotoState21Action,
	{_State149, VarToken}:                                _GotoState37Action,
	{_State149, LetToken}:                                _GotoState27Action,
	{_State149, NotToken}:                                _GotoState30Action,
	{_State149, LabelDeclToken}:                          _GotoState25Action,
	{_State149, LparenToken}:                             _GotoState28Action,
	{_State149, LbracketToken}:                           _GotoState26Action,
	{_State149, SubToken}:                                _GotoState35Action,
	{_State149, MulToken}:                                _GotoState29Action,
	{_State149, BitNegToken}:                             _GotoState18Action,
	{_State149, BitAndToken}:                             _GotoState17Action,
	{_State149, GreaterToken}:                            _GotoState22Action,
	{_State149, ParseErrorToken}:                         _GotoState31Action,
	{_State149, VarDeclPatternType}:                      _GotoState69Action,
	{_State149, VarOrLetType}:                            _GotoState70Action,
	{_State149, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State149, SequenceExprType}:                        _GotoState223Action,
	{_State149, ConditionType}:                           _GotoState222Action,
	{_State149, CallExprType}:                            _GotoState50Action,
	{_State149, AtomExprType}:                            _GotoState43Action,
	{_State149, ParseErrorExprType}:                      _GotoState62Action,
	{_State149, LiteralExprType}:                         _GotoState58Action,
	{_State149, IdentifierExprType}:                      _GotoState53Action,
	{_State149, BlockExprType}:                           _GotoState49Action,
	{_State149, InitializeExprType}:                      _GotoState57Action,
	{_State149, ImplicitStructExprType}:                  _GotoState54Action,
	{_State149, AccessibleExprType}:                      _GotoState39Action,
	{_State149, AccessExprType}:                          _GotoState38Action,
	{_State149, IndexExprType}:                           _GotoState55Action,
	{_State149, PostfixableExprType}:                     _GotoState64Action,
	{_State149, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State149, PrefixableExprType}:                      _GotoState67Action,
	{_State149, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State149, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State149, MulExprType}:                             _GotoState59Action,
	{_State149, BinaryMulExprType}:                       _GotoState47Action,
	{_State149, AddExprType}:                             _GotoState40Action,
	{_State149, BinaryAddExprType}:                       _GotoState44Action,
	{_State149, CmpExprType}:                             _GotoState51Action,
	{_State149, BinaryCmpExprType}:                       _GotoState46Action,
	{_State149, AndExprType}:                             _GotoState41Action,
	{_State149, BinaryAndExprType}:                       _GotoState45Action,
	{_State149, OrExprType}:                              _GotoState61Action,
	{_State149, BinaryOrExprType}:                        _GotoState48Action,
	{_State149, InitializableTypeType}:                   _GotoState56Action,
	{_State149, ExplicitStructDefType}:                   _GotoState52Action,
	{_State149, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State150, IntegerLiteralToken}:                     _GotoState24Action,
	{_State150, FloatLiteralToken}:                       _GotoState20Action,
	{_State150, RuneLiteralToken}:                        _GotoState32Action,
	{_State150, StringLiteralToken}:                      _GotoState33Action,
	{_State150, IdentifierToken}:                         _GotoState23Action,
	{_State150, TrueToken}:                               _GotoState36Action,
	{_State150, FalseToken}:                              _GotoState19Action,
	{_State150, StructToken}:                             _GotoState34Action,
	{_State150, FuncToken}:                               _GotoState21Action,
	{_State150, VarToken}:                                _GotoState37Action,
	{_State150, LetToken}:                                _GotoState27Action,
	{_State150, NotToken}:                                _GotoState30Action,
	{_State150, LabelDeclToken}:                          _GotoState25Action,
	{_State150, LparenToken}:                             _GotoState28Action,
	{_State150, LbracketToken}:                           _GotoState26Action,
	{_State150, SubToken}:                                _GotoState35Action,
	{_State150, MulToken}:                                _GotoState29Action,
	{_State150, BitNegToken}:                             _GotoState18Action,
	{_State150, BitAndToken}:                             _GotoState17Action,
	{_State150, GreaterToken}:                            _GotoState22Action,
	{_State150, ParseErrorToken}:                         _GotoState31Action,
	{_State150, VarDeclPatternType}:                      _GotoState69Action,
	{_State150, VarOrLetType}:                            _GotoState70Action,
	{_State150, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State150, SequenceExprType}:                        _GotoState224Action,
	{_State150, CallExprType}:                            _GotoState50Action,
	{_State150, AtomExprType}:                            _GotoState43Action,
	{_State150, ParseErrorExprType}:                      _GotoState62Action,
	{_State150, LiteralExprType}:                         _GotoState58Action,
	{_State150, IdentifierExprType}:                      _GotoState53Action,
	{_State150, BlockExprType}:                           _GotoState49Action,
	{_State150, InitializeExprType}:                      _GotoState57Action,
	{_State150, ImplicitStructExprType}:                  _GotoState54Action,
	{_State150, AccessibleExprType}:                      _GotoState39Action,
	{_State150, AccessExprType}:                          _GotoState38Action,
	{_State150, IndexExprType}:                           _GotoState55Action,
	{_State150, PostfixableExprType}:                     _GotoState64Action,
	{_State150, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State150, PrefixableExprType}:                      _GotoState67Action,
	{_State150, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State150, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State150, MulExprType}:                             _GotoState59Action,
	{_State150, BinaryMulExprType}:                       _GotoState47Action,
	{_State150, AddExprType}:                             _GotoState40Action,
	{_State150, BinaryAddExprType}:                       _GotoState44Action,
	{_State150, CmpExprType}:                             _GotoState51Action,
	{_State150, BinaryCmpExprType}:                       _GotoState46Action,
	{_State150, AndExprType}:                             _GotoState41Action,
	{_State150, BinaryAndExprType}:                       _GotoState45Action,
	{_State150, OrExprType}:                              _GotoState61Action,
	{_State150, BinaryOrExprType}:                        _GotoState48Action,
	{_State150, InitializableTypeType}:                   _GotoState56Action,
	{_State150, ExplicitStructDefType}:                   _GotoState52Action,
	{_State150, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State155, IntegerLiteralToken}:                     _GotoState24Action,
	{_State155, FloatLiteralToken}:                       _GotoState20Action,
	{_State155, RuneLiteralToken}:                        _GotoState32Action,
	{_State155, StringLiteralToken}:                      _GotoState33Action,
	{_State155, IdentifierToken}:                         _GotoState23Action,
	{_State155, TrueToken}:                               _GotoState36Action,
	{_State155, FalseToken}:                              _GotoState19Action,
	{_State155, StructToken}:                             _GotoState34Action,
	{_State155, FuncToken}:                               _GotoState21Action,
	{_State155, NotToken}:                                _GotoState30Action,
	{_State155, LabelDeclToken}:                          _GotoState25Action,
	{_State155, LparenToken}:                             _GotoState28Action,
	{_State155, LbracketToken}:                           _GotoState26Action,
	{_State155, SubToken}:                                _GotoState35Action,
	{_State155, MulToken}:                                _GotoState29Action,
	{_State155, BitNegToken}:                             _GotoState18Action,
	{_State155, BitAndToken}:                             _GotoState17Action,
	{_State155, ParseErrorToken}:                         _GotoState31Action,
	{_State155, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State155, CallExprType}:                            _GotoState50Action,
	{_State155, AtomExprType}:                            _GotoState43Action,
	{_State155, ParseErrorExprType}:                      _GotoState62Action,
	{_State155, LiteralExprType}:                         _GotoState58Action,
	{_State155, IdentifierExprType}:                      _GotoState53Action,
	{_State155, BlockExprType}:                           _GotoState49Action,
	{_State155, InitializeExprType}:                      _GotoState57Action,
	{_State155, ImplicitStructExprType}:                  _GotoState54Action,
	{_State155, AccessibleExprType}:                      _GotoState39Action,
	{_State155, AccessExprType}:                          _GotoState38Action,
	{_State155, IndexExprType}:                           _GotoState55Action,
	{_State155, PostfixableExprType}:                     _GotoState64Action,
	{_State155, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State155, PrefixableExprType}:                      _GotoState67Action,
	{_State155, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State155, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State155, MulExprType}:                             _GotoState59Action,
	{_State155, BinaryMulExprType}:                       _GotoState47Action,
	{_State155, AddExprType}:                             _GotoState40Action,
	{_State155, BinaryAddExprType}:                       _GotoState44Action,
	{_State155, CmpExprType}:                             _GotoState51Action,
	{_State155, BinaryCmpExprType}:                       _GotoState46Action,
	{_State155, AndExprType}:                             _GotoState225Action,
	{_State155, BinaryAndExprType}:                       _GotoState45Action,
	{_State155, InitializableTypeType}:                   _GotoState56Action,
	{_State155, ExplicitStructDefType}:                   _GotoState52Action,
	{_State155, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State158, IdentifierToken}:                         _GotoState227Action,
	{_State158, LparenToken}:                             _GotoState158Action,
	{_State158, DotDotDotToken}:                          _GotoState226Action,
	{_State158, VarPatternType}:                          _GotoState230Action,
	{_State158, TuplePatternType}:                        _GotoState159Action,
	{_State158, FieldVarPatternsType}:                    _GotoState229Action,
	{_State158, FieldVarPatternType}:                     _GotoState228Action,
	{_State160, IdentifierToken}:                         _GotoState93Action,
	{_State160, StructToken}:                             _GotoState34Action,
	{_State160, EnumToken}:                               _GotoState90Action,
	{_State160, TraitToken}:                              _GotoState98Action,
	{_State160, FuncToken}:                               _GotoState92Action,
	{_State160, LparenToken}:                             _GotoState94Action,
	{_State160, LbracketToken}:                           _GotoState26Action,
	{_State160, DotToken}:                                _GotoState89Action,
	{_State160, QuestionToken}:                           _GotoState96Action,
	{_State160, ExclaimToken}:                            _GotoState91Action,
	{_State160, TildeTildeToken}:                         _GotoState97Action,
	{_State160, BitNegToken}:                             _GotoState88Action,
	{_State160, BitAndToken}:                             _GotoState87Action,
	{_State160, ParseErrorToken}:                         _GotoState95Action,
	{_State160, OptionalValueTypeType}:                   _GotoState231Action,
	{_State160, InitializableTypeType}:                   _GotoState104Action,
	{_State160, AtomTypeType}:                            _GotoState99Action,
	{_State160, ParseErrorTypeType}:                      _GotoState105Action,
	{_State160, ReturnableTypeType}:                      _GotoState108Action,
	{_State160, PrefixedTypeType}:                        _GotoState107Action,
	{_State160, PrefixTypeOpType}:                        _GotoState106Action,
	{_State160, ValueTypeType}:                           _GotoState232Action,
	{_State160, TraitOpTypeType}:                         _GotoState110Action,
	{_State160, ImplicitStructDefType}:                   _GotoState103Action,
	{_State160, ExplicitStructDefType}:                   _GotoState52Action,
	{_State160, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State160, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State160, TraitDefType}:                            _GotoState109Action,
	{_State160, FuncTypeType}:                            _GotoState101Action,
	{_State161, IntegerLiteralToken}:                     _GotoState24Action,
	{_State161, FloatLiteralToken}:                       _GotoState20Action,
	{_State161, RuneLiteralToken}:                        _GotoState32Action,
	{_State161, StringLiteralToken}:                      _GotoState33Action,
	{_State161, IdentifierToken}:                         _GotoState23Action,
	{_State161, TrueToken}:                               _GotoState36Action,
	{_State161, FalseToken}:                              _GotoState19Action,
	{_State161, CaseToken}:                               _GotoState235Action,
	{_State161, DefaultToken}:                            _GotoState237Action,
	{_State161, ReturnToken}:                             _GotoState242Action,
	{_State161, BreakToken}:                              _GotoState234Action,
	{_State161, ContinueToken}:                           _GotoState236Action,
	{_State161, FallthroughToken}:                        _GotoState239Action,
	{_State161, ImportToken}:                             _GotoState240Action,
	{_State161, UnsafeToken}:                             _GotoState181Action,
	{_State161, StructToken}:                             _GotoState34Action,
	{_State161, FuncToken}:                               _GotoState21Action,
	{_State161, AsyncToken}:                              _GotoState233Action,
	{_State161, DeferToken}:                              _GotoState238Action,
	{_State161, VarToken}:                                _GotoState37Action,
	{_State161, LetToken}:                                _GotoState27Action,
	{_State161, NotToken}:                                _GotoState30Action,
	{_State161, LabelDeclToken}:                          _GotoState25Action,
	{_State161, RbraceToken}:                             _GotoState241Action,
	{_State161, LparenToken}:                             _GotoState28Action,
	{_State161, LbracketToken}:                           _GotoState26Action,
	{_State161, SubToken}:                                _GotoState35Action,
	{_State161, MulToken}:                                _GotoState29Action,
	{_State161, BitNegToken}:                             _GotoState18Action,
	{_State161, BitAndToken}:                             _GotoState17Action,
	{_State161, GreaterToken}:                            _GotoState22Action,
	{_State161, ParseErrorToken}:                         _GotoState31Action,
	{_State161, StatementType}:                           _GotoState258Action,
	{_State161, SimpleStatementBodyType}:                 _GotoState257Action,
	{_State161, StatementBodyType}:                       _GotoState259Action,
	{_State161, ExpressionOrImproperStructStatementType}: _GotoState250Action,
	{_State161, CallbackOpType}:                          _GotoState247Action,
	{_State161, CallbackOpStatementType}:                 _GotoState248Action,
	{_State161, UnsafeStatementType}:                     _GotoState261Action,
	{_State161, JumpStatementType}:                       _GotoState254Action,
	{_State161, JumpTypeType}:                            _GotoState255Action,
	{_State161, ExpressionsType}:                         _GotoState251Action,
	{_State161, FallthroughStatementType}:                _GotoState252Action,
	{_State161, AssignStatementType}:                     _GotoState245Action,
	{_State161, UnaryOpAssignStatementType}:              _GotoState260Action,
	{_State161, BinaryOpAssignStatementType}:             _GotoState246Action,
	{_State161, ImportStatementType}:                     _GotoState253Action,
	{_State161, VarDeclPatternType}:                      _GotoState69Action,
	{_State161, VarOrLetType}:                            _GotoState70Action,
	{_State161, AssignPatternType}:                       _GotoState244Action,
	{_State161, ExpressionType}:                          _GotoState249Action,
	{_State161, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State161, SequenceExprType}:                        _GotoState256Action,
	{_State161, CallExprType}:                            _GotoState50Action,
	{_State161, AtomExprType}:                            _GotoState43Action,
	{_State161, ParseErrorExprType}:                      _GotoState62Action,
	{_State161, LiteralExprType}:                         _GotoState58Action,
	{_State161, IdentifierExprType}:                      _GotoState53Action,
	{_State161, BlockExprType}:                           _GotoState49Action,
	{_State161, InitializeExprType}:                      _GotoState57Action,
	{_State161, ImplicitStructExprType}:                  _GotoState54Action,
	{_State161, AccessibleExprType}:                      _GotoState243Action,
	{_State161, AccessExprType}:                          _GotoState38Action,
	{_State161, IndexExprType}:                           _GotoState55Action,
	{_State161, PostfixableExprType}:                     _GotoState64Action,
	{_State161, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State161, PrefixableExprType}:                      _GotoState67Action,
	{_State161, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State161, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State161, MulExprType}:                             _GotoState59Action,
	{_State161, BinaryMulExprType}:                       _GotoState47Action,
	{_State161, AddExprType}:                             _GotoState40Action,
	{_State161, BinaryAddExprType}:                       _GotoState44Action,
	{_State161, CmpExprType}:                             _GotoState51Action,
	{_State161, BinaryCmpExprType}:                       _GotoState46Action,
	{_State161, AndExprType}:                             _GotoState41Action,
	{_State161, BinaryAndExprType}:                       _GotoState45Action,
	{_State161, OrExprType}:                              _GotoState61Action,
	{_State161, BinaryOrExprType}:                        _GotoState48Action,
	{_State161, InitializableTypeType}:                   _GotoState56Action,
	{_State161, ExplicitStructDefType}:                   _GotoState52Action,
	{_State161, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State162, CommentGroupsToken}:                      _GotoState71Action,
	{_State162, PackageToken}:                            _GotoState14Action,
	{_State162, TypeToken}:                               _GotoState15Action,
	{_State162, FuncToken}:                               _GotoState16Action,
	{_State162, VarToken}:                                _GotoState37Action,
	{_State162, LetToken}:                                _GotoState27Action,
	{_State162, LbraceToken}:                             _GotoState72Action,
	{_State162, DefinitionType}:                          _GotoState262Action,
	{_State162, StatementBlockType}:                      _GotoState77Action,
	{_State162, VarDeclPatternType}:                      _GotoState79Action,
	{_State162, VarOrLetType}:                            _GotoState70Action,
	{_State162, TypeDefType}:                             _GotoState78Action,
	{_State162, NamedFuncDefType}:                        _GotoState75Action,
	{_State162, PackageDefType}:                          _GotoState76Action,
	{_State164, IntegerLiteralToken}:                     _GotoState24Action,
	{_State164, FloatLiteralToken}:                       _GotoState20Action,
	{_State164, RuneLiteralToken}:                        _GotoState32Action,
	{_State164, StringLiteralToken}:                      _GotoState33Action,
	{_State164, IdentifierToken}:                         _GotoState23Action,
	{_State164, TrueToken}:                               _GotoState36Action,
	{_State164, FalseToken}:                              _GotoState19Action,
	{_State164, StructToken}:                             _GotoState34Action,
	{_State164, FuncToken}:                               _GotoState21Action,
	{_State164, VarToken}:                                _GotoState37Action,
	{_State164, LetToken}:                                _GotoState27Action,
	{_State164, NotToken}:                                _GotoState30Action,
	{_State164, LabelDeclToken}:                          _GotoState25Action,
	{_State164, LparenToken}:                             _GotoState28Action,
	{_State164, LbracketToken}:                           _GotoState26Action,
	{_State164, SubToken}:                                _GotoState35Action,
	{_State164, MulToken}:                                _GotoState29Action,
	{_State164, BitNegToken}:                             _GotoState18Action,
	{_State164, BitAndToken}:                             _GotoState17Action,
	{_State164, GreaterToken}:                            _GotoState22Action,
	{_State164, ParseErrorToken}:                         _GotoState31Action,
	{_State164, VarDeclPatternType}:                      _GotoState69Action,
	{_State164, VarOrLetType}:                            _GotoState70Action,
	{_State164, ExpressionType}:                          _GotoState263Action,
	{_State164, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State164, SequenceExprType}:                        _GotoState68Action,
	{_State164, CallExprType}:                            _GotoState50Action,
	{_State164, AtomExprType}:                            _GotoState43Action,
	{_State164, ParseErrorExprType}:                      _GotoState62Action,
	{_State164, LiteralExprType}:                         _GotoState58Action,
	{_State164, IdentifierExprType}:                      _GotoState53Action,
	{_State164, BlockExprType}:                           _GotoState49Action,
	{_State164, InitializeExprType}:                      _GotoState57Action,
	{_State164, ImplicitStructExprType}:                  _GotoState54Action,
	{_State164, AccessibleExprType}:                      _GotoState39Action,
	{_State164, AccessExprType}:                          _GotoState38Action,
	{_State164, IndexExprType}:                           _GotoState55Action,
	{_State164, PostfixableExprType}:                     _GotoState64Action,
	{_State164, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State164, PrefixableExprType}:                      _GotoState67Action,
	{_State164, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State164, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State164, MulExprType}:                             _GotoState59Action,
	{_State164, BinaryMulExprType}:                       _GotoState47Action,
	{_State164, AddExprType}:                             _GotoState40Action,
	{_State164, BinaryAddExprType}:                       _GotoState44Action,
	{_State164, CmpExprType}:                             _GotoState51Action,
	{_State164, BinaryCmpExprType}:                       _GotoState46Action,
	{_State164, AndExprType}:                             _GotoState41Action,
	{_State164, BinaryAndExprType}:                       _GotoState45Action,
	{_State164, OrExprType}:                              _GotoState61Action,
	{_State164, BinaryOrExprType}:                        _GotoState48Action,
	{_State164, InitializableTypeType}:                   _GotoState56Action,
	{_State164, ExplicitStructDefType}:                   _GotoState52Action,
	{_State164, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State165, IdentifierToken}:                         _GotoState93Action,
	{_State165, StructToken}:                             _GotoState34Action,
	{_State165, EnumToken}:                               _GotoState90Action,
	{_State165, TraitToken}:                              _GotoState98Action,
	{_State165, FuncToken}:                               _GotoState92Action,
	{_State165, LparenToken}:                             _GotoState94Action,
	{_State165, LbracketToken}:                           _GotoState26Action,
	{_State165, DotToken}:                                _GotoState89Action,
	{_State165, QuestionToken}:                           _GotoState96Action,
	{_State165, ExclaimToken}:                            _GotoState91Action,
	{_State165, TildeTildeToken}:                         _GotoState97Action,
	{_State165, BitNegToken}:                             _GotoState88Action,
	{_State165, BitAndToken}:                             _GotoState87Action,
	{_State165, ParseErrorToken}:                         _GotoState95Action,
	{_State165, InitializableTypeType}:                   _GotoState104Action,
	{_State165, AtomTypeType}:                            _GotoState99Action,
	{_State165, ParseErrorTypeType}:                      _GotoState105Action,
	{_State165, ReturnableTypeType}:                      _GotoState108Action,
	{_State165, PrefixedTypeType}:                        _GotoState107Action,
	{_State165, PrefixTypeOpType}:                        _GotoState106Action,
	{_State165, ValueTypeType}:                           _GotoState264Action,
	{_State165, TraitOpTypeType}:                         _GotoState110Action,
	{_State165, ImplicitStructDefType}:                   _GotoState103Action,
	{_State165, ExplicitStructDefType}:                   _GotoState52Action,
	{_State165, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State165, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State165, TraitDefType}:                            _GotoState109Action,
	{_State165, FuncTypeType}:                            _GotoState101Action,
	{_State166, IdentifierToken}:                         _GotoState265Action,
	{_State166, GenericParameterDefType}:                 _GotoState266Action,
	{_State166, GenericParameterDefsType}:                _GotoState267Action,
	{_State166, OptionalGenericParameterDefsType}:        _GotoState268Action,
	{_State167, IdentifierToken}:                         _GotoState93Action,
	{_State167, StructToken}:                             _GotoState34Action,
	{_State167, EnumToken}:                               _GotoState90Action,
	{_State167, TraitToken}:                              _GotoState98Action,
	{_State167, FuncToken}:                               _GotoState92Action,
	{_State167, LparenToken}:                             _GotoState94Action,
	{_State167, LbracketToken}:                           _GotoState26Action,
	{_State167, DotToken}:                                _GotoState89Action,
	{_State167, QuestionToken}:                           _GotoState96Action,
	{_State167, ExclaimToken}:                            _GotoState91Action,
	{_State167, TildeTildeToken}:                         _GotoState97Action,
	{_State167, BitNegToken}:                             _GotoState88Action,
	{_State167, BitAndToken}:                             _GotoState87Action,
	{_State167, ParseErrorToken}:                         _GotoState95Action,
	{_State167, InitializableTypeType}:                   _GotoState104Action,
	{_State167, AtomTypeType}:                            _GotoState99Action,
	{_State167, ParseErrorTypeType}:                      _GotoState105Action,
	{_State167, ReturnableTypeType}:                      _GotoState108Action,
	{_State167, PrefixedTypeType}:                        _GotoState107Action,
	{_State167, PrefixTypeOpType}:                        _GotoState106Action,
	{_State167, ValueTypeType}:                           _GotoState269Action,
	{_State167, TraitOpTypeType}:                         _GotoState110Action,
	{_State167, ImplicitStructDefType}:                   _GotoState103Action,
	{_State167, ExplicitStructDefType}:                   _GotoState52Action,
	{_State167, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State167, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State167, TraitDefType}:                            _GotoState109Action,
	{_State167, FuncTypeType}:                            _GotoState101Action,
	{_State168, IntegerLiteralToken}:                     _GotoState24Action,
	{_State168, FloatLiteralToken}:                       _GotoState20Action,
	{_State168, RuneLiteralToken}:                        _GotoState32Action,
	{_State168, StringLiteralToken}:                      _GotoState33Action,
	{_State168, IdentifierToken}:                         _GotoState23Action,
	{_State168, TrueToken}:                               _GotoState36Action,
	{_State168, FalseToken}:                              _GotoState19Action,
	{_State168, StructToken}:                             _GotoState34Action,
	{_State168, FuncToken}:                               _GotoState21Action,
	{_State168, VarToken}:                                _GotoState37Action,
	{_State168, LetToken}:                                _GotoState27Action,
	{_State168, NotToken}:                                _GotoState30Action,
	{_State168, LabelDeclToken}:                          _GotoState25Action,
	{_State168, LparenToken}:                             _GotoState28Action,
	{_State168, LbracketToken}:                           _GotoState26Action,
	{_State168, SubToken}:                                _GotoState35Action,
	{_State168, MulToken}:                                _GotoState29Action,
	{_State168, BitNegToken}:                             _GotoState18Action,
	{_State168, BitAndToken}:                             _GotoState17Action,
	{_State168, GreaterToken}:                            _GotoState22Action,
	{_State168, ParseErrorToken}:                         _GotoState31Action,
	{_State168, VarDeclPatternType}:                      _GotoState69Action,
	{_State168, VarOrLetType}:                            _GotoState70Action,
	{_State168, ExpressionType}:                          _GotoState270Action,
	{_State168, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State168, SequenceExprType}:                        _GotoState68Action,
	{_State168, CallExprType}:                            _GotoState50Action,
	{_State168, AtomExprType}:                            _GotoState43Action,
	{_State168, ParseErrorExprType}:                      _GotoState62Action,
	{_State168, LiteralExprType}:                         _GotoState58Action,
	{_State168, IdentifierExprType}:                      _GotoState53Action,
	{_State168, BlockExprType}:                           _GotoState49Action,
	{_State168, InitializeExprType}:                      _GotoState57Action,
	{_State168, ImplicitStructExprType}:                  _GotoState54Action,
	{_State168, AccessibleExprType}:                      _GotoState39Action,
	{_State168, AccessExprType}:                          _GotoState38Action,
	{_State168, IndexExprType}:                           _GotoState55Action,
	{_State168, PostfixableExprType}:                     _GotoState64Action,
	{_State168, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State168, PrefixableExprType}:                      _GotoState67Action,
	{_State168, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State168, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State168, MulExprType}:                             _GotoState59Action,
	{_State168, BinaryMulExprType}:                       _GotoState47Action,
	{_State168, AddExprType}:                             _GotoState40Action,
	{_State168, BinaryAddExprType}:                       _GotoState44Action,
	{_State168, CmpExprType}:                             _GotoState51Action,
	{_State168, BinaryCmpExprType}:                       _GotoState46Action,
	{_State168, AndExprType}:                             _GotoState41Action,
	{_State168, BinaryAndExprType}:                       _GotoState45Action,
	{_State168, OrExprType}:                              _GotoState61Action,
	{_State168, BinaryOrExprType}:                        _GotoState48Action,
	{_State168, InitializableTypeType}:                   _GotoState56Action,
	{_State168, ExplicitStructDefType}:                   _GotoState52Action,
	{_State168, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State169, LparenToken}:                             _GotoState271Action,
	{_State170, IdentifierToken}:                         _GotoState93Action,
	{_State170, StructToken}:                             _GotoState34Action,
	{_State170, EnumToken}:                               _GotoState90Action,
	{_State170, TraitToken}:                              _GotoState98Action,
	{_State170, FuncToken}:                               _GotoState92Action,
	{_State170, LparenToken}:                             _GotoState94Action,
	{_State170, LbracketToken}:                           _GotoState26Action,
	{_State170, DotToken}:                                _GotoState89Action,
	{_State170, QuestionToken}:                           _GotoState96Action,
	{_State170, ExclaimToken}:                            _GotoState91Action,
	{_State170, DotDotDotToken}:                          _GotoState272Action,
	{_State170, TildeTildeToken}:                         _GotoState97Action,
	{_State170, BitNegToken}:                             _GotoState88Action,
	{_State170, BitAndToken}:                             _GotoState87Action,
	{_State170, ParseErrorToken}:                         _GotoState95Action,
	{_State170, InitializableTypeType}:                   _GotoState104Action,
	{_State170, AtomTypeType}:                            _GotoState99Action,
	{_State170, ParseErrorTypeType}:                      _GotoState105Action,
	{_State170, ReturnableTypeType}:                      _GotoState108Action,
	{_State170, PrefixedTypeType}:                        _GotoState107Action,
	{_State170, PrefixTypeOpType}:                        _GotoState106Action,
	{_State170, ValueTypeType}:                           _GotoState273Action,
	{_State170, TraitOpTypeType}:                         _GotoState110Action,
	{_State170, ImplicitStructDefType}:                   _GotoState103Action,
	{_State170, ExplicitStructDefType}:                   _GotoState52Action,
	{_State170, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State170, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State170, TraitDefType}:                            _GotoState109Action,
	{_State170, FuncTypeType}:                            _GotoState101Action,
	{_State171, RparenToken}:                             _GotoState274Action,
	{_State172, RparenToken}:                             _GotoState275Action,
	{_State174, CommaToken}:                              _GotoState276Action,
	{_State176, IdentifierToken}:                         _GotoState180Action,
	{_State176, UnsafeToken}:                             _GotoState181Action,
	{_State176, StructToken}:                             _GotoState34Action,
	{_State176, EnumToken}:                               _GotoState90Action,
	{_State176, TraitToken}:                              _GotoState98Action,
	{_State176, FuncToken}:                               _GotoState92Action,
	{_State176, LparenToken}:                             _GotoState94Action,
	{_State176, LbracketToken}:                           _GotoState26Action,
	{_State176, DotToken}:                                _GotoState89Action,
	{_State176, QuestionToken}:                           _GotoState96Action,
	{_State176, ExclaimToken}:                            _GotoState91Action,
	{_State176, TildeTildeToken}:                         _GotoState97Action,
	{_State176, BitNegToken}:                             _GotoState88Action,
	{_State176, BitAndToken}:                             _GotoState87Action,
	{_State176, ParseErrorToken}:                         _GotoState95Action,
	{_State176, UnsafeStatementType}:                     _GotoState187Action,
	{_State176, InitializableTypeType}:                   _GotoState104Action,
	{_State176, AtomTypeType}:                            _GotoState99Action,
	{_State176, ParseErrorTypeType}:                      _GotoState105Action,
	{_State176, ReturnableTypeType}:                      _GotoState108Action,
	{_State176, PrefixedTypeType}:                        _GotoState107Action,
	{_State176, PrefixTypeOpType}:                        _GotoState106Action,
	{_State176, ValueTypeType}:                           _GotoState188Action,
	{_State176, TraitOpTypeType}:                         _GotoState110Action,
	{_State176, FieldDefType}:                            _GotoState279Action,
	{_State176, ImplicitStructDefType}:                   _GotoState103Action,
	{_State176, ExplicitStructDefType}:                   _GotoState52Action,
	{_State176, EnumValueDefType}:                        _GotoState277Action,
	{_State176, ImplicitEnumValueDefsType}:               _GotoState280Action,
	{_State176, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State176, ExplicitEnumValueDefsType}:               _GotoState278Action,
	{_State176, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State176, TraitDefType}:                            _GotoState109Action,
	{_State176, FuncTypeType}:                            _GotoState101Action,
	{_State177, IdentifierToken}:                         _GotoState282Action,
	{_State177, StructToken}:                             _GotoState34Action,
	{_State177, EnumToken}:                               _GotoState90Action,
	{_State177, TraitToken}:                              _GotoState98Action,
	{_State177, FuncToken}:                               _GotoState92Action,
	{_State177, LparenToken}:                             _GotoState94Action,
	{_State177, LbracketToken}:                           _GotoState26Action,
	{_State177, DotToken}:                                _GotoState89Action,
	{_State177, QuestionToken}:                           _GotoState96Action,
	{_State177, ExclaimToken}:                            _GotoState91Action,
	{_State177, DotDotDotToken}:                          _GotoState281Action,
	{_State177, TildeTildeToken}:                         _GotoState97Action,
	{_State177, BitNegToken}:                             _GotoState88Action,
	{_State177, BitAndToken}:                             _GotoState87Action,
	{_State177, ParseErrorToken}:                         _GotoState95Action,
	{_State177, InitializableTypeType}:                   _GotoState104Action,
	{_State177, AtomTypeType}:                            _GotoState99Action,
	{_State177, ParseErrorTypeType}:                      _GotoState105Action,
	{_State177, ReturnableTypeType}:                      _GotoState108Action,
	{_State177, PrefixedTypeType}:                        _GotoState107Action,
	{_State177, PrefixTypeOpType}:                        _GotoState106Action,
	{_State177, ValueTypeType}:                           _GotoState286Action,
	{_State177, TraitOpTypeType}:                         _GotoState110Action,
	{_State177, ImplicitStructDefType}:                   _GotoState103Action,
	{_State177, ExplicitStructDefType}:                   _GotoState52Action,
	{_State177, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State177, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State177, TraitDefType}:                            _GotoState109Action,
	{_State177, ParameterDeclType}:                       _GotoState284Action,
	{_State177, ParameterDeclsType}:                      _GotoState285Action,
	{_State177, OptionalParameterDeclsType}:              _GotoState283Action,
	{_State177, FuncTypeType}:                            _GotoState101Action,
	{_State178, IdentifierToken}:                         _GotoState287Action,
	{_State180, IdentifierToken}:                         _GotoState93Action,
	{_State180, StructToken}:                             _GotoState34Action,
	{_State180, EnumToken}:                               _GotoState90Action,
	{_State180, TraitToken}:                              _GotoState98Action,
	{_State180, FuncToken}:                               _GotoState92Action,
	{_State180, LparenToken}:                             _GotoState94Action,
	{_State180, LbracketToken}:                           _GotoState26Action,
	{_State180, DotToken}:                                _GotoState288Action,
	{_State180, QuestionToken}:                           _GotoState96Action,
	{_State180, ExclaimToken}:                            _GotoState91Action,
	{_State180, DollarLbracketToken}:                     _GotoState121Action,
	{_State180, TildeTildeToken}:                         _GotoState97Action,
	{_State180, BitNegToken}:                             _GotoState88Action,
	{_State180, BitAndToken}:                             _GotoState87Action,
	{_State180, ParseErrorToken}:                         _GotoState95Action,
	{_State180, OptionalGenericBindingType}:              _GotoState179Action,
	{_State180, InitializableTypeType}:                   _GotoState104Action,
	{_State180, AtomTypeType}:                            _GotoState99Action,
	{_State180, ParseErrorTypeType}:                      _GotoState105Action,
	{_State180, ReturnableTypeType}:                      _GotoState108Action,
	{_State180, PrefixedTypeType}:                        _GotoState107Action,
	{_State180, PrefixTypeOpType}:                        _GotoState106Action,
	{_State180, ValueTypeType}:                           _GotoState289Action,
	{_State180, TraitOpTypeType}:                         _GotoState110Action,
	{_State180, ImplicitStructDefType}:                   _GotoState103Action,
	{_State180, ExplicitStructDefType}:                   _GotoState52Action,
	{_State180, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State180, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State180, TraitDefType}:                            _GotoState109Action,
	{_State180, FuncTypeType}:                            _GotoState101Action,
	{_State181, LessToken}:                               _GotoState290Action,
	{_State182, OrToken}:                                 _GotoState291Action,
	{_State183, AssignToken}:                             _GotoState292Action,
	{_State184, OrToken}:                                 _GotoState293Action,
	{_State184, RparenToken}:                             _GotoState294Action,
	{_State185, CommaToken}:                              _GotoState295Action,
	{_State186, RparenToken}:                             _GotoState296Action,
	{_State188, AddToken}:                                _GotoState191Action,
	{_State188, SubToken}:                                _GotoState196Action,
	{_State188, MulToken}:                                _GotoState194Action,
	{_State188, TraitOpType}:                             _GotoState197Action,
	{_State189, IdentifierToken}:                         _GotoState180Action,
	{_State189, UnsafeToken}:                             _GotoState181Action,
	{_State189, StructToken}:                             _GotoState34Action,
	{_State189, EnumToken}:                               _GotoState90Action,
	{_State189, TraitToken}:                              _GotoState98Action,
	{_State189, FuncToken}:                               _GotoState297Action,
	{_State189, LparenToken}:                             _GotoState94Action,
	{_State189, LbracketToken}:                           _GotoState26Action,
	{_State189, DotToken}:                                _GotoState89Action,
	{_State189, QuestionToken}:                           _GotoState96Action,
	{_State189, ExclaimToken}:                            _GotoState91Action,
	{_State189, TildeTildeToken}:                         _GotoState97Action,
	{_State189, BitNegToken}:                             _GotoState88Action,
	{_State189, BitAndToken}:                             _GotoState87Action,
	{_State189, ParseErrorToken}:                         _GotoState95Action,
	{_State189, UnsafeStatementType}:                     _GotoState187Action,
	{_State189, InitializableTypeType}:                   _GotoState104Action,
	{_State189, AtomTypeType}:                            _GotoState99Action,
	{_State189, ParseErrorTypeType}:                      _GotoState105Action,
	{_State189, ReturnableTypeType}:                      _GotoState108Action,
	{_State189, PrefixedTypeType}:                        _GotoState107Action,
	{_State189, PrefixTypeOpType}:                        _GotoState106Action,
	{_State189, ValueTypeType}:                           _GotoState188Action,
	{_State189, TraitOpTypeType}:                         _GotoState110Action,
	{_State189, FieldDefType}:                            _GotoState298Action,
	{_State189, ImplicitStructDefType}:                   _GotoState103Action,
	{_State189, ExplicitStructDefType}:                   _GotoState52Action,
	{_State189, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State189, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State189, TraitPropertyType}:                       _GotoState302Action,
	{_State189, TraitPropertiesType}:                     _GotoState301Action,
	{_State189, OptionalTraitPropertiesType}:             _GotoState300Action,
	{_State189, TraitDefType}:                            _GotoState109Action,
	{_State189, FuncTypeType}:                            _GotoState101Action,
	{_State189, MethodSignatureType}:                     _GotoState299Action,
	{_State192, IdentifierToken}:                         _GotoState93Action,
	{_State192, StructToken}:                             _GotoState34Action,
	{_State192, EnumToken}:                               _GotoState90Action,
	{_State192, TraitToken}:                              _GotoState98Action,
	{_State192, FuncToken}:                               _GotoState92Action,
	{_State192, LparenToken}:                             _GotoState94Action,
	{_State192, LbracketToken}:                           _GotoState26Action,
	{_State192, DotToken}:                                _GotoState89Action,
	{_State192, QuestionToken}:                           _GotoState96Action,
	{_State192, ExclaimToken}:                            _GotoState91Action,
	{_State192, TildeTildeToken}:                         _GotoState97Action,
	{_State192, BitNegToken}:                             _GotoState88Action,
	{_State192, BitAndToken}:                             _GotoState87Action,
	{_State192, ParseErrorToken}:                         _GotoState95Action,
	{_State192, InitializableTypeType}:                   _GotoState104Action,
	{_State192, AtomTypeType}:                            _GotoState99Action,
	{_State192, ParseErrorTypeType}:                      _GotoState105Action,
	{_State192, ReturnableTypeType}:                      _GotoState108Action,
	{_State192, PrefixedTypeType}:                        _GotoState107Action,
	{_State192, PrefixTypeOpType}:                        _GotoState106Action,
	{_State192, ValueTypeType}:                           _GotoState303Action,
	{_State192, TraitOpTypeType}:                         _GotoState110Action,
	{_State192, ImplicitStructDefType}:                   _GotoState103Action,
	{_State192, ExplicitStructDefType}:                   _GotoState52Action,
	{_State192, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State192, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State192, TraitDefType}:                            _GotoState109Action,
	{_State192, FuncTypeType}:                            _GotoState101Action,
	{_State193, IntegerLiteralToken}:                     _GotoState304Action,
	{_State197, IdentifierToken}:                         _GotoState93Action,
	{_State197, StructToken}:                             _GotoState34Action,
	{_State197, EnumToken}:                               _GotoState90Action,
	{_State197, TraitToken}:                              _GotoState98Action,
	{_State197, FuncToken}:                               _GotoState92Action,
	{_State197, LparenToken}:                             _GotoState94Action,
	{_State197, LbracketToken}:                           _GotoState26Action,
	{_State197, DotToken}:                                _GotoState89Action,
	{_State197, QuestionToken}:                           _GotoState96Action,
	{_State197, ExclaimToken}:                            _GotoState91Action,
	{_State197, TildeTildeToken}:                         _GotoState97Action,
	{_State197, BitNegToken}:                             _GotoState88Action,
	{_State197, BitAndToken}:                             _GotoState87Action,
	{_State197, ParseErrorToken}:                         _GotoState95Action,
	{_State197, InitializableTypeType}:                   _GotoState104Action,
	{_State197, AtomTypeType}:                            _GotoState99Action,
	{_State197, ParseErrorTypeType}:                      _GotoState105Action,
	{_State197, ReturnableTypeType}:                      _GotoState305Action,
	{_State197, PrefixedTypeType}:                        _GotoState107Action,
	{_State197, PrefixTypeOpType}:                        _GotoState106Action,
	{_State197, ImplicitStructDefType}:                   _GotoState103Action,
	{_State197, ExplicitStructDefType}:                   _GotoState52Action,
	{_State197, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State197, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State197, TraitDefType}:                            _GotoState109Action,
	{_State197, FuncTypeType}:                            _GotoState101Action,
	{_State198, IntegerLiteralToken}:                     _GotoState24Action,
	{_State198, FloatLiteralToken}:                       _GotoState20Action,
	{_State198, RuneLiteralToken}:                        _GotoState32Action,
	{_State198, StringLiteralToken}:                      _GotoState33Action,
	{_State198, IdentifierToken}:                         _GotoState23Action,
	{_State198, TrueToken}:                               _GotoState36Action,
	{_State198, FalseToken}:                              _GotoState19Action,
	{_State198, StructToken}:                             _GotoState34Action,
	{_State198, FuncToken}:                               _GotoState21Action,
	{_State198, VarToken}:                                _GotoState37Action,
	{_State198, LetToken}:                                _GotoState27Action,
	{_State198, NotToken}:                                _GotoState30Action,
	{_State198, LabelDeclToken}:                          _GotoState25Action,
	{_State198, LparenToken}:                             _GotoState28Action,
	{_State198, LbracketToken}:                           _GotoState26Action,
	{_State198, SubToken}:                                _GotoState35Action,
	{_State198, MulToken}:                                _GotoState29Action,
	{_State198, BitNegToken}:                             _GotoState18Action,
	{_State198, BitAndToken}:                             _GotoState17Action,
	{_State198, GreaterToken}:                            _GotoState22Action,
	{_State198, ParseErrorToken}:                         _GotoState31Action,
	{_State198, VarDeclPatternType}:                      _GotoState69Action,
	{_State198, VarOrLetType}:                            _GotoState70Action,
	{_State198, ExpressionType}:                          _GotoState306Action,
	{_State198, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State198, SequenceExprType}:                        _GotoState68Action,
	{_State198, CallExprType}:                            _GotoState50Action,
	{_State198, AtomExprType}:                            _GotoState43Action,
	{_State198, ParseErrorExprType}:                      _GotoState62Action,
	{_State198, LiteralExprType}:                         _GotoState58Action,
	{_State198, IdentifierExprType}:                      _GotoState53Action,
	{_State198, BlockExprType}:                           _GotoState49Action,
	{_State198, InitializeExprType}:                      _GotoState57Action,
	{_State198, ImplicitStructExprType}:                  _GotoState54Action,
	{_State198, AccessibleExprType}:                      _GotoState39Action,
	{_State198, AccessExprType}:                          _GotoState38Action,
	{_State198, IndexExprType}:                           _GotoState55Action,
	{_State198, PostfixableExprType}:                     _GotoState64Action,
	{_State198, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State198, PrefixableExprType}:                      _GotoState67Action,
	{_State198, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State198, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State198, MulExprType}:                             _GotoState59Action,
	{_State198, BinaryMulExprType}:                       _GotoState47Action,
	{_State198, AddExprType}:                             _GotoState40Action,
	{_State198, BinaryAddExprType}:                       _GotoState44Action,
	{_State198, CmpExprType}:                             _GotoState51Action,
	{_State198, BinaryCmpExprType}:                       _GotoState46Action,
	{_State198, AndExprType}:                             _GotoState41Action,
	{_State198, BinaryAndExprType}:                       _GotoState45Action,
	{_State198, OrExprType}:                              _GotoState61Action,
	{_State198, BinaryOrExprType}:                        _GotoState48Action,
	{_State198, InitializableTypeType}:                   _GotoState56Action,
	{_State198, ExplicitStructDefType}:                   _GotoState52Action,
	{_State198, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State199, IntegerLiteralToken}:                     _GotoState24Action,
	{_State199, FloatLiteralToken}:                       _GotoState20Action,
	{_State199, RuneLiteralToken}:                        _GotoState32Action,
	{_State199, StringLiteralToken}:                      _GotoState33Action,
	{_State199, IdentifierToken}:                         _GotoState113Action,
	{_State199, TrueToken}:                               _GotoState36Action,
	{_State199, FalseToken}:                              _GotoState19Action,
	{_State199, StructToken}:                             _GotoState34Action,
	{_State199, FuncToken}:                               _GotoState21Action,
	{_State199, VarToken}:                                _GotoState37Action,
	{_State199, LetToken}:                                _GotoState27Action,
	{_State199, NotToken}:                                _GotoState30Action,
	{_State199, LabelDeclToken}:                          _GotoState25Action,
	{_State199, LparenToken}:                             _GotoState28Action,
	{_State199, LbracketToken}:                           _GotoState26Action,
	{_State199, DotDotDotToken}:                          _GotoState112Action,
	{_State199, SubToken}:                                _GotoState35Action,
	{_State199, MulToken}:                                _GotoState29Action,
	{_State199, BitNegToken}:                             _GotoState18Action,
	{_State199, BitAndToken}:                             _GotoState17Action,
	{_State199, GreaterToken}:                            _GotoState22Action,
	{_State199, ParseErrorToken}:                         _GotoState31Action,
	{_State199, VarDeclPatternType}:                      _GotoState69Action,
	{_State199, VarOrLetType}:                            _GotoState70Action,
	{_State199, ExpressionType}:                          _GotoState117Action,
	{_State199, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State199, SequenceExprType}:                        _GotoState68Action,
	{_State199, CallExprType}:                            _GotoState50Action,
	{_State199, ArgumentType}:                            _GotoState307Action,
	{_State199, ColonExpressionsType}:                    _GotoState116Action,
	{_State199, OptionalExpressionType}:                  _GotoState119Action,
	{_State199, AtomExprType}:                            _GotoState43Action,
	{_State199, ParseErrorExprType}:                      _GotoState62Action,
	{_State199, LiteralExprType}:                         _GotoState58Action,
	{_State199, IdentifierExprType}:                      _GotoState53Action,
	{_State199, BlockExprType}:                           _GotoState49Action,
	{_State199, InitializeExprType}:                      _GotoState57Action,
	{_State199, ImplicitStructExprType}:                  _GotoState54Action,
	{_State199, AccessibleExprType}:                      _GotoState39Action,
	{_State199, AccessExprType}:                          _GotoState38Action,
	{_State199, IndexExprType}:                           _GotoState55Action,
	{_State199, PostfixableExprType}:                     _GotoState64Action,
	{_State199, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State199, PrefixableExprType}:                      _GotoState67Action,
	{_State199, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State199, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State199, MulExprType}:                             _GotoState59Action,
	{_State199, BinaryMulExprType}:                       _GotoState47Action,
	{_State199, AddExprType}:                             _GotoState40Action,
	{_State199, BinaryAddExprType}:                       _GotoState44Action,
	{_State199, CmpExprType}:                             _GotoState51Action,
	{_State199, BinaryCmpExprType}:                       _GotoState46Action,
	{_State199, AndExprType}:                             _GotoState41Action,
	{_State199, BinaryAndExprType}:                       _GotoState45Action,
	{_State199, OrExprType}:                              _GotoState61Action,
	{_State199, BinaryOrExprType}:                        _GotoState48Action,
	{_State199, InitializableTypeType}:                   _GotoState56Action,
	{_State199, ExplicitStructDefType}:                   _GotoState52Action,
	{_State199, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State200, IntegerLiteralToken}:                     _GotoState24Action,
	{_State200, FloatLiteralToken}:                       _GotoState20Action,
	{_State200, RuneLiteralToken}:                        _GotoState32Action,
	{_State200, StringLiteralToken}:                      _GotoState33Action,
	{_State200, IdentifierToken}:                         _GotoState23Action,
	{_State200, TrueToken}:                               _GotoState36Action,
	{_State200, FalseToken}:                              _GotoState19Action,
	{_State200, StructToken}:                             _GotoState34Action,
	{_State200, FuncToken}:                               _GotoState21Action,
	{_State200, VarToken}:                                _GotoState37Action,
	{_State200, LetToken}:                                _GotoState27Action,
	{_State200, NotToken}:                                _GotoState30Action,
	{_State200, LabelDeclToken}:                          _GotoState25Action,
	{_State200, LparenToken}:                             _GotoState28Action,
	{_State200, LbracketToken}:                           _GotoState26Action,
	{_State200, SubToken}:                                _GotoState35Action,
	{_State200, MulToken}:                                _GotoState29Action,
	{_State200, BitNegToken}:                             _GotoState18Action,
	{_State200, BitAndToken}:                             _GotoState17Action,
	{_State200, GreaterToken}:                            _GotoState22Action,
	{_State200, ParseErrorToken}:                         _GotoState31Action,
	{_State200, VarDeclPatternType}:                      _GotoState69Action,
	{_State200, VarOrLetType}:                            _GotoState70Action,
	{_State200, ExpressionType}:                          _GotoState308Action,
	{_State200, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State200, SequenceExprType}:                        _GotoState68Action,
	{_State200, CallExprType}:                            _GotoState50Action,
	{_State200, OptionalExpressionType}:                  _GotoState309Action,
	{_State200, AtomExprType}:                            _GotoState43Action,
	{_State200, ParseErrorExprType}:                      _GotoState62Action,
	{_State200, LiteralExprType}:                         _GotoState58Action,
	{_State200, IdentifierExprType}:                      _GotoState53Action,
	{_State200, BlockExprType}:                           _GotoState49Action,
	{_State200, InitializeExprType}:                      _GotoState57Action,
	{_State200, ImplicitStructExprType}:                  _GotoState54Action,
	{_State200, AccessibleExprType}:                      _GotoState39Action,
	{_State200, AccessExprType}:                          _GotoState38Action,
	{_State200, IndexExprType}:                           _GotoState55Action,
	{_State200, PostfixableExprType}:                     _GotoState64Action,
	{_State200, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State200, PrefixableExprType}:                      _GotoState67Action,
	{_State200, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State200, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State200, MulExprType}:                             _GotoState59Action,
	{_State200, BinaryMulExprType}:                       _GotoState47Action,
	{_State200, AddExprType}:                             _GotoState40Action,
	{_State200, BinaryAddExprType}:                       _GotoState44Action,
	{_State200, CmpExprType}:                             _GotoState51Action,
	{_State200, BinaryCmpExprType}:                       _GotoState46Action,
	{_State200, AndExprType}:                             _GotoState41Action,
	{_State200, BinaryAndExprType}:                       _GotoState45Action,
	{_State200, OrExprType}:                              _GotoState61Action,
	{_State200, BinaryOrExprType}:                        _GotoState48Action,
	{_State200, InitializableTypeType}:                   _GotoState56Action,
	{_State200, ExplicitStructDefType}:                   _GotoState52Action,
	{_State200, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State202, IntegerLiteralToken}:                     _GotoState24Action,
	{_State202, FloatLiteralToken}:                       _GotoState20Action,
	{_State202, RuneLiteralToken}:                        _GotoState32Action,
	{_State202, StringLiteralToken}:                      _GotoState33Action,
	{_State202, IdentifierToken}:                         _GotoState23Action,
	{_State202, TrueToken}:                               _GotoState36Action,
	{_State202, FalseToken}:                              _GotoState19Action,
	{_State202, StructToken}:                             _GotoState34Action,
	{_State202, FuncToken}:                               _GotoState21Action,
	{_State202, VarToken}:                                _GotoState37Action,
	{_State202, LetToken}:                                _GotoState27Action,
	{_State202, NotToken}:                                _GotoState30Action,
	{_State202, LabelDeclToken}:                          _GotoState25Action,
	{_State202, LparenToken}:                             _GotoState28Action,
	{_State202, LbracketToken}:                           _GotoState26Action,
	{_State202, SubToken}:                                _GotoState35Action,
	{_State202, MulToken}:                                _GotoState29Action,
	{_State202, BitNegToken}:                             _GotoState18Action,
	{_State202, BitAndToken}:                             _GotoState17Action,
	{_State202, GreaterToken}:                            _GotoState22Action,
	{_State202, ParseErrorToken}:                         _GotoState31Action,
	{_State202, VarDeclPatternType}:                      _GotoState69Action,
	{_State202, VarOrLetType}:                            _GotoState70Action,
	{_State202, ExpressionType}:                          _GotoState308Action,
	{_State202, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State202, SequenceExprType}:                        _GotoState68Action,
	{_State202, CallExprType}:                            _GotoState50Action,
	{_State202, OptionalExpressionType}:                  _GotoState310Action,
	{_State202, AtomExprType}:                            _GotoState43Action,
	{_State202, ParseErrorExprType}:                      _GotoState62Action,
	{_State202, LiteralExprType}:                         _GotoState58Action,
	{_State202, IdentifierExprType}:                      _GotoState53Action,
	{_State202, BlockExprType}:                           _GotoState49Action,
	{_State202, InitializeExprType}:                      _GotoState57Action,
	{_State202, ImplicitStructExprType}:                  _GotoState54Action,
	{_State202, AccessibleExprType}:                      _GotoState39Action,
	{_State202, AccessExprType}:                          _GotoState38Action,
	{_State202, IndexExprType}:                           _GotoState55Action,
	{_State202, PostfixableExprType}:                     _GotoState64Action,
	{_State202, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State202, PrefixableExprType}:                      _GotoState67Action,
	{_State202, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State202, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State202, MulExprType}:                             _GotoState59Action,
	{_State202, BinaryMulExprType}:                       _GotoState47Action,
	{_State202, AddExprType}:                             _GotoState40Action,
	{_State202, BinaryAddExprType}:                       _GotoState44Action,
	{_State202, CmpExprType}:                             _GotoState51Action,
	{_State202, BinaryCmpExprType}:                       _GotoState46Action,
	{_State202, AndExprType}:                             _GotoState41Action,
	{_State202, BinaryAndExprType}:                       _GotoState45Action,
	{_State202, OrExprType}:                              _GotoState61Action,
	{_State202, BinaryOrExprType}:                        _GotoState48Action,
	{_State202, InitializableTypeType}:                   _GotoState56Action,
	{_State202, ExplicitStructDefType}:                   _GotoState52Action,
	{_State202, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State203, NewlinesToken}:                           _GotoState312Action,
	{_State203, CommaToken}:                              _GotoState311Action,
	{_State205, RparenToken}:                             _GotoState313Action,
	{_State206, CommaToken}:                              _GotoState314Action,
	{_State207, RbracketToken}:                           _GotoState315Action,
	{_State208, AddToken}:                                _GotoState191Action,
	{_State208, SubToken}:                                _GotoState196Action,
	{_State208, MulToken}:                                _GotoState194Action,
	{_State208, TraitOpType}:                             _GotoState197Action,
	{_State210, RbracketToken}:                           _GotoState316Action,
	{_State211, IntegerLiteralToken}:                     _GotoState24Action,
	{_State211, FloatLiteralToken}:                       _GotoState20Action,
	{_State211, RuneLiteralToken}:                        _GotoState32Action,
	{_State211, StringLiteralToken}:                      _GotoState33Action,
	{_State211, IdentifierToken}:                         _GotoState113Action,
	{_State211, TrueToken}:                               _GotoState36Action,
	{_State211, FalseToken}:                              _GotoState19Action,
	{_State211, StructToken}:                             _GotoState34Action,
	{_State211, FuncToken}:                               _GotoState21Action,
	{_State211, VarToken}:                                _GotoState37Action,
	{_State211, LetToken}:                                _GotoState27Action,
	{_State211, NotToken}:                                _GotoState30Action,
	{_State211, LabelDeclToken}:                          _GotoState25Action,
	{_State211, LparenToken}:                             _GotoState28Action,
	{_State211, LbracketToken}:                           _GotoState26Action,
	{_State211, DotDotDotToken}:                          _GotoState112Action,
	{_State211, SubToken}:                                _GotoState35Action,
	{_State211, MulToken}:                                _GotoState29Action,
	{_State211, BitNegToken}:                             _GotoState18Action,
	{_State211, BitAndToken}:                             _GotoState17Action,
	{_State211, GreaterToken}:                            _GotoState22Action,
	{_State211, ParseErrorToken}:                         _GotoState31Action,
	{_State211, VarDeclPatternType}:                      _GotoState69Action,
	{_State211, VarOrLetType}:                            _GotoState70Action,
	{_State211, ExpressionType}:                          _GotoState117Action,
	{_State211, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State211, SequenceExprType}:                        _GotoState68Action,
	{_State211, CallExprType}:                            _GotoState50Action,
	{_State211, OptionalArgumentsType}:                   _GotoState317Action,
	{_State211, ArgumentsType}:                           _GotoState115Action,
	{_State211, ArgumentType}:                            _GotoState114Action,
	{_State211, ColonExpressionsType}:                    _GotoState116Action,
	{_State211, OptionalExpressionType}:                  _GotoState119Action,
	{_State211, AtomExprType}:                            _GotoState43Action,
	{_State211, ParseErrorExprType}:                      _GotoState62Action,
	{_State211, LiteralExprType}:                         _GotoState58Action,
	{_State211, IdentifierExprType}:                      _GotoState53Action,
	{_State211, BlockExprType}:                           _GotoState49Action,
	{_State211, InitializeExprType}:                      _GotoState57Action,
	{_State211, ImplicitStructExprType}:                  _GotoState54Action,
	{_State211, AccessibleExprType}:                      _GotoState39Action,
	{_State211, AccessExprType}:                          _GotoState38Action,
	{_State211, IndexExprType}:                           _GotoState55Action,
	{_State211, PostfixableExprType}:                     _GotoState64Action,
	{_State211, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State211, PrefixableExprType}:                      _GotoState67Action,
	{_State211, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State211, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State211, MulExprType}:                             _GotoState59Action,
	{_State211, BinaryMulExprType}:                       _GotoState47Action,
	{_State211, AddExprType}:                             _GotoState40Action,
	{_State211, BinaryAddExprType}:                       _GotoState44Action,
	{_State211, CmpExprType}:                             _GotoState51Action,
	{_State211, BinaryCmpExprType}:                       _GotoState46Action,
	{_State211, AndExprType}:                             _GotoState41Action,
	{_State211, BinaryAndExprType}:                       _GotoState45Action,
	{_State211, OrExprType}:                              _GotoState61Action,
	{_State211, BinaryOrExprType}:                        _GotoState48Action,
	{_State211, InitializableTypeType}:                   _GotoState56Action,
	{_State211, ExplicitStructDefType}:                   _GotoState52Action,
	{_State211, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State212, MulToken}:                                _GotoState145Action,
	{_State212, DivToken}:                                _GotoState143Action,
	{_State212, ModToken}:                                _GotoState144Action,
	{_State212, BitAndToken}:                             _GotoState140Action,
	{_State212, BitLshiftToken}:                          _GotoState141Action,
	{_State212, BitRshiftToken}:                          _GotoState142Action,
	{_State212, MulOpType}:                               _GotoState146Action,
	{_State213, EqualToken}:                              _GotoState132Action,
	{_State213, NotEqualToken}:                           _GotoState137Action,
	{_State213, LessToken}:                               _GotoState135Action,
	{_State213, LessOrEqualToken}:                        _GotoState136Action,
	{_State213, GreaterToken}:                            _GotoState133Action,
	{_State213, GreaterOrEqualToken}:                     _GotoState134Action,
	{_State213, CmpOpType}:                               _GotoState138Action,
	{_State214, AddToken}:                                _GotoState126Action,
	{_State214, SubToken}:                                _GotoState129Action,
	{_State214, BitXorToken}:                             _GotoState128Action,
	{_State214, BitOrToken}:                              _GotoState127Action,
	{_State214, AddOpType}:                               _GotoState130Action,
	{_State215, RparenToken}:                             _GotoState318Action,
	{_State217, ForToken}:                                _GotoState319Action,
	{_State218, InToken}:                                 _GotoState321Action,
	{_State218, AssignToken}:                             _GotoState320Action,
	{_State219, SemicolonToken}:                          _GotoState322Action,
	{_State220, DoToken}:                                 _GotoState323Action,
	{_State221, IntegerLiteralToken}:                     _GotoState24Action,
	{_State221, FloatLiteralToken}:                       _GotoState20Action,
	{_State221, RuneLiteralToken}:                        _GotoState32Action,
	{_State221, StringLiteralToken}:                      _GotoState33Action,
	{_State221, IdentifierToken}:                         _GotoState23Action,
	{_State221, TrueToken}:                               _GotoState36Action,
	{_State221, FalseToken}:                              _GotoState19Action,
	{_State221, StructToken}:                             _GotoState34Action,
	{_State221, FuncToken}:                               _GotoState21Action,
	{_State221, VarToken}:                                _GotoState325Action,
	{_State221, LetToken}:                                _GotoState27Action,
	{_State221, NotToken}:                                _GotoState30Action,
	{_State221, LabelDeclToken}:                          _GotoState25Action,
	{_State221, LparenToken}:                             _GotoState28Action,
	{_State221, LbracketToken}:                           _GotoState26Action,
	{_State221, DotToken}:                                _GotoState324Action,
	{_State221, SubToken}:                                _GotoState35Action,
	{_State221, MulToken}:                                _GotoState29Action,
	{_State221, BitNegToken}:                             _GotoState18Action,
	{_State221, BitAndToken}:                             _GotoState17Action,
	{_State221, GreaterToken}:                            _GotoState22Action,
	{_State221, ParseErrorToken}:                         _GotoState31Action,
	{_State221, CasePatternsType}:                        _GotoState328Action,
	{_State221, VarDeclPatternType}:                      _GotoState69Action,
	{_State221, VarOrLetType}:                            _GotoState70Action,
	{_State221, AssignPatternType}:                       _GotoState326Action,
	{_State221, CasePatternType}:                         _GotoState327Action,
	{_State221, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State221, SequenceExprType}:                        _GotoState329Action,
	{_State221, CallExprType}:                            _GotoState50Action,
	{_State221, AtomExprType}:                            _GotoState43Action,
	{_State221, ParseErrorExprType}:                      _GotoState62Action,
	{_State221, LiteralExprType}:                         _GotoState58Action,
	{_State221, IdentifierExprType}:                      _GotoState53Action,
	{_State221, BlockExprType}:                           _GotoState49Action,
	{_State221, InitializeExprType}:                      _GotoState57Action,
	{_State221, ImplicitStructExprType}:                  _GotoState54Action,
	{_State221, AccessibleExprType}:                      _GotoState39Action,
	{_State221, AccessExprType}:                          _GotoState38Action,
	{_State221, IndexExprType}:                           _GotoState55Action,
	{_State221, PostfixableExprType}:                     _GotoState64Action,
	{_State221, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State221, PrefixableExprType}:                      _GotoState67Action,
	{_State221, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State221, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State221, MulExprType}:                             _GotoState59Action,
	{_State221, BinaryMulExprType}:                       _GotoState47Action,
	{_State221, AddExprType}:                             _GotoState40Action,
	{_State221, BinaryAddExprType}:                       _GotoState44Action,
	{_State221, CmpExprType}:                             _GotoState51Action,
	{_State221, BinaryCmpExprType}:                       _GotoState46Action,
	{_State221, AndExprType}:                             _GotoState41Action,
	{_State221, BinaryAndExprType}:                       _GotoState45Action,
	{_State221, OrExprType}:                              _GotoState61Action,
	{_State221, BinaryOrExprType}:                        _GotoState48Action,
	{_State221, InitializableTypeType}:                   _GotoState56Action,
	{_State221, ExplicitStructDefType}:                   _GotoState52Action,
	{_State221, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State222, LbraceToken}:                             _GotoState72Action,
	{_State222, StatementBlockType}:                      _GotoState330Action,
	{_State224, LbraceToken}:                             _GotoState72Action,
	{_State224, StatementBlockType}:                      _GotoState331Action,
	{_State225, AndToken}:                                _GotoState131Action,
	{_State227, AssignToken}:                             _GotoState332Action,
	{_State229, RparenToken}:                             _GotoState334Action,
	{_State229, CommaToken}:                              _GotoState333Action,
	{_State232, AddToken}:                                _GotoState191Action,
	{_State232, SubToken}:                                _GotoState196Action,
	{_State232, MulToken}:                                _GotoState194Action,
	{_State232, TraitOpType}:                             _GotoState197Action,
	{_State235, IntegerLiteralToken}:                     _GotoState24Action,
	{_State235, FloatLiteralToken}:                       _GotoState20Action,
	{_State235, RuneLiteralToken}:                        _GotoState32Action,
	{_State235, StringLiteralToken}:                      _GotoState33Action,
	{_State235, IdentifierToken}:                         _GotoState23Action,
	{_State235, TrueToken}:                               _GotoState36Action,
	{_State235, FalseToken}:                              _GotoState19Action,
	{_State235, StructToken}:                             _GotoState34Action,
	{_State235, FuncToken}:                               _GotoState21Action,
	{_State235, VarToken}:                                _GotoState325Action,
	{_State235, LetToken}:                                _GotoState27Action,
	{_State235, NotToken}:                                _GotoState30Action,
	{_State235, LabelDeclToken}:                          _GotoState25Action,
	{_State235, LparenToken}:                             _GotoState28Action,
	{_State235, LbracketToken}:                           _GotoState26Action,
	{_State235, DotToken}:                                _GotoState324Action,
	{_State235, SubToken}:                                _GotoState35Action,
	{_State235, MulToken}:                                _GotoState29Action,
	{_State235, BitNegToken}:                             _GotoState18Action,
	{_State235, BitAndToken}:                             _GotoState17Action,
	{_State235, GreaterToken}:                            _GotoState22Action,
	{_State235, ParseErrorToken}:                         _GotoState31Action,
	{_State235, CasePatternsType}:                        _GotoState335Action,
	{_State235, VarDeclPatternType}:                      _GotoState69Action,
	{_State235, VarOrLetType}:                            _GotoState70Action,
	{_State235, AssignPatternType}:                       _GotoState326Action,
	{_State235, CasePatternType}:                         _GotoState327Action,
	{_State235, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State235, SequenceExprType}:                        _GotoState329Action,
	{_State235, CallExprType}:                            _GotoState50Action,
	{_State235, AtomExprType}:                            _GotoState43Action,
	{_State235, ParseErrorExprType}:                      _GotoState62Action,
	{_State235, LiteralExprType}:                         _GotoState58Action,
	{_State235, IdentifierExprType}:                      _GotoState53Action,
	{_State235, BlockExprType}:                           _GotoState49Action,
	{_State235, InitializeExprType}:                      _GotoState57Action,
	{_State235, ImplicitStructExprType}:                  _GotoState54Action,
	{_State235, AccessibleExprType}:                      _GotoState39Action,
	{_State235, AccessExprType}:                          _GotoState38Action,
	{_State235, IndexExprType}:                           _GotoState55Action,
	{_State235, PostfixableExprType}:                     _GotoState64Action,
	{_State235, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State235, PrefixableExprType}:                      _GotoState67Action,
	{_State235, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State235, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State235, MulExprType}:                             _GotoState59Action,
	{_State235, BinaryMulExprType}:                       _GotoState47Action,
	{_State235, AddExprType}:                             _GotoState40Action,
	{_State235, BinaryAddExprType}:                       _GotoState44Action,
	{_State235, CmpExprType}:                             _GotoState51Action,
	{_State235, BinaryCmpExprType}:                       _GotoState46Action,
	{_State235, AndExprType}:                             _GotoState41Action,
	{_State235, BinaryAndExprType}:                       _GotoState45Action,
	{_State235, OrExprType}:                              _GotoState61Action,
	{_State235, BinaryOrExprType}:                        _GotoState48Action,
	{_State235, InitializableTypeType}:                   _GotoState56Action,
	{_State235, ExplicitStructDefType}:                   _GotoState52Action,
	{_State235, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State237, ColonToken}:                              _GotoState336Action,
	{_State240, StringLiteralToken}:                      _GotoState338Action,
	{_State240, LparenToken}:                             _GotoState337Action,
	{_State240, ImportClauseType}:                        _GotoState339Action,
	{_State243, LbracketToken}:                           _GotoState123Action,
	{_State243, DotToken}:                                _GotoState122Action,
	{_State243, QuestionToken}:                           _GotoState124Action,
	{_State243, DollarLbracketToken}:                     _GotoState121Action,
	{_State243, AddAssignToken}:                          _GotoState340Action,
	{_State243, SubAssignToken}:                          _GotoState351Action,
	{_State243, MulAssignToken}:                          _GotoState350Action,
	{_State243, DivAssignToken}:                          _GotoState348Action,
	{_State243, ModAssignToken}:                          _GotoState349Action,
	{_State243, AddOneAssignToken}:                       _GotoState341Action,
	{_State243, SubOneAssignToken}:                       _GotoState352Action,
	{_State243, BitNegAssignToken}:                       _GotoState344Action,
	{_State243, BitAndAssignToken}:                       _GotoState342Action,
	{_State243, BitOrAssignToken}:                        _GotoState345Action,
	{_State243, BitXorAssignToken}:                       _GotoState347Action,
	{_State243, BitLshiftAssignToken}:                    _GotoState343Action,
	{_State243, BitRshiftAssignToken}:                    _GotoState346Action,
	{_State243, UnaryOpAssignType}:                       _GotoState354Action,
	{_State243, BinaryOpAssignType}:                      _GotoState353Action,
	{_State243, OptionalGenericBindingType}:              _GotoState125Action,
	{_State244, AssignToken}:                             _GotoState355Action,
	{_State247, IntegerLiteralToken}:                     _GotoState24Action,
	{_State247, FloatLiteralToken}:                       _GotoState20Action,
	{_State247, RuneLiteralToken}:                        _GotoState32Action,
	{_State247, StringLiteralToken}:                      _GotoState33Action,
	{_State247, IdentifierToken}:                         _GotoState23Action,
	{_State247, TrueToken}:                               _GotoState36Action,
	{_State247, FalseToken}:                              _GotoState19Action,
	{_State247, StructToken}:                             _GotoState34Action,
	{_State247, FuncToken}:                               _GotoState21Action,
	{_State247, LabelDeclToken}:                          _GotoState25Action,
	{_State247, LparenToken}:                             _GotoState28Action,
	{_State247, LbracketToken}:                           _GotoState26Action,
	{_State247, ParseErrorToken}:                         _GotoState31Action,
	{_State247, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State247, CallExprType}:                            _GotoState357Action,
	{_State247, AtomExprType}:                            _GotoState43Action,
	{_State247, ParseErrorExprType}:                      _GotoState62Action,
	{_State247, LiteralExprType}:                         _GotoState58Action,
	{_State247, IdentifierExprType}:                      _GotoState53Action,
	{_State247, BlockExprType}:                           _GotoState49Action,
	{_State247, InitializeExprType}:                      _GotoState57Action,
	{_State247, ImplicitStructExprType}:                  _GotoState54Action,
	{_State247, AccessibleExprType}:                      _GotoState356Action,
	{_State247, AccessExprType}:                          _GotoState38Action,
	{_State247, IndexExprType}:                           _GotoState55Action,
	{_State247, InitializableTypeType}:                   _GotoState56Action,
	{_State247, ExplicitStructDefType}:                   _GotoState52Action,
	{_State247, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State251, CommaToken}:                              _GotoState358Action,
	{_State255, IntegerLiteralToken}:                     _GotoState24Action,
	{_State255, FloatLiteralToken}:                       _GotoState20Action,
	{_State255, RuneLiteralToken}:                        _GotoState32Action,
	{_State255, StringLiteralToken}:                      _GotoState33Action,
	{_State255, IdentifierToken}:                         _GotoState23Action,
	{_State255, TrueToken}:                               _GotoState36Action,
	{_State255, FalseToken}:                              _GotoState19Action,
	{_State255, StructToken}:                             _GotoState34Action,
	{_State255, FuncToken}:                               _GotoState21Action,
	{_State255, VarToken}:                                _GotoState37Action,
	{_State255, LetToken}:                                _GotoState27Action,
	{_State255, NotToken}:                                _GotoState30Action,
	{_State255, LabelDeclToken}:                          _GotoState25Action,
	{_State255, JumpLabelToken}:                          _GotoState359Action,
	{_State255, LparenToken}:                             _GotoState28Action,
	{_State255, LbracketToken}:                           _GotoState26Action,
	{_State255, SubToken}:                                _GotoState35Action,
	{_State255, MulToken}:                                _GotoState29Action,
	{_State255, BitNegToken}:                             _GotoState18Action,
	{_State255, BitAndToken}:                             _GotoState17Action,
	{_State255, GreaterToken}:                            _GotoState22Action,
	{_State255, ParseErrorToken}:                         _GotoState31Action,
	{_State255, ExpressionsType}:                         _GotoState360Action,
	{_State255, VarDeclPatternType}:                      _GotoState69Action,
	{_State255, VarOrLetType}:                            _GotoState70Action,
	{_State255, ExpressionType}:                          _GotoState249Action,
	{_State255, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State255, SequenceExprType}:                        _GotoState68Action,
	{_State255, CallExprType}:                            _GotoState50Action,
	{_State255, AtomExprType}:                            _GotoState43Action,
	{_State255, ParseErrorExprType}:                      _GotoState62Action,
	{_State255, LiteralExprType}:                         _GotoState58Action,
	{_State255, IdentifierExprType}:                      _GotoState53Action,
	{_State255, BlockExprType}:                           _GotoState49Action,
	{_State255, InitializeExprType}:                      _GotoState57Action,
	{_State255, ImplicitStructExprType}:                  _GotoState54Action,
	{_State255, AccessibleExprType}:                      _GotoState39Action,
	{_State255, AccessExprType}:                          _GotoState38Action,
	{_State255, IndexExprType}:                           _GotoState55Action,
	{_State255, PostfixableExprType}:                     _GotoState64Action,
	{_State255, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State255, PrefixableExprType}:                      _GotoState67Action,
	{_State255, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State255, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State255, MulExprType}:                             _GotoState59Action,
	{_State255, BinaryMulExprType}:                       _GotoState47Action,
	{_State255, AddExprType}:                             _GotoState40Action,
	{_State255, BinaryAddExprType}:                       _GotoState44Action,
	{_State255, CmpExprType}:                             _GotoState51Action,
	{_State255, BinaryCmpExprType}:                       _GotoState46Action,
	{_State255, AndExprType}:                             _GotoState41Action,
	{_State255, BinaryAndExprType}:                       _GotoState45Action,
	{_State255, OrExprType}:                              _GotoState61Action,
	{_State255, BinaryOrExprType}:                        _GotoState48Action,
	{_State255, InitializableTypeType}:                   _GotoState56Action,
	{_State255, ExplicitStructDefType}:                   _GotoState52Action,
	{_State255, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State259, NewlinesToken}:                           _GotoState361Action,
	{_State259, SemicolonToken}:                          _GotoState362Action,
	{_State264, AddToken}:                                _GotoState191Action,
	{_State264, SubToken}:                                _GotoState196Action,
	{_State264, MulToken}:                                _GotoState194Action,
	{_State264, TraitOpType}:                             _GotoState197Action,
	{_State265, IdentifierToken}:                         _GotoState93Action,
	{_State265, StructToken}:                             _GotoState34Action,
	{_State265, EnumToken}:                               _GotoState90Action,
	{_State265, TraitToken}:                              _GotoState98Action,
	{_State265, FuncToken}:                               _GotoState92Action,
	{_State265, LparenToken}:                             _GotoState94Action,
	{_State265, LbracketToken}:                           _GotoState26Action,
	{_State265, DotToken}:                                _GotoState89Action,
	{_State265, QuestionToken}:                           _GotoState96Action,
	{_State265, ExclaimToken}:                            _GotoState91Action,
	{_State265, TildeTildeToken}:                         _GotoState97Action,
	{_State265, BitNegToken}:                             _GotoState88Action,
	{_State265, BitAndToken}:                             _GotoState87Action,
	{_State265, ParseErrorToken}:                         _GotoState95Action,
	{_State265, InitializableTypeType}:                   _GotoState104Action,
	{_State265, AtomTypeType}:                            _GotoState99Action,
	{_State265, ParseErrorTypeType}:                      _GotoState105Action,
	{_State265, ReturnableTypeType}:                      _GotoState108Action,
	{_State265, PrefixedTypeType}:                        _GotoState107Action,
	{_State265, PrefixTypeOpType}:                        _GotoState106Action,
	{_State265, ValueTypeType}:                           _GotoState363Action,
	{_State265, TraitOpTypeType}:                         _GotoState110Action,
	{_State265, ImplicitStructDefType}:                   _GotoState103Action,
	{_State265, ExplicitStructDefType}:                   _GotoState52Action,
	{_State265, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State265, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State265, TraitDefType}:                            _GotoState109Action,
	{_State265, FuncTypeType}:                            _GotoState101Action,
	{_State267, CommaToken}:                              _GotoState364Action,
	{_State268, RbracketToken}:                           _GotoState365Action,
	{_State269, ImplementsToken}:                         _GotoState366Action,
	{_State269, AddToken}:                                _GotoState191Action,
	{_State269, SubToken}:                                _GotoState196Action,
	{_State269, MulToken}:                                _GotoState194Action,
	{_State269, TraitOpType}:                             _GotoState197Action,
	{_State271, IdentifierToken}:                         _GotoState170Action,
	{_State271, ParameterDefType}:                        _GotoState173Action,
	{_State271, ParameterDefsType}:                       _GotoState174Action,
	{_State271, OptionalParameterDefsType}:               _GotoState367Action,
	{_State272, IdentifierToken}:                         _GotoState93Action,
	{_State272, StructToken}:                             _GotoState34Action,
	{_State272, EnumToken}:                               _GotoState90Action,
	{_State272, TraitToken}:                              _GotoState98Action,
	{_State272, FuncToken}:                               _GotoState92Action,
	{_State272, LparenToken}:                             _GotoState94Action,
	{_State272, LbracketToken}:                           _GotoState26Action,
	{_State272, DotToken}:                                _GotoState89Action,
	{_State272, QuestionToken}:                           _GotoState96Action,
	{_State272, ExclaimToken}:                            _GotoState91Action,
	{_State272, TildeTildeToken}:                         _GotoState97Action,
	{_State272, BitNegToken}:                             _GotoState88Action,
	{_State272, BitAndToken}:                             _GotoState87Action,
	{_State272, ParseErrorToken}:                         _GotoState95Action,
	{_State272, InitializableTypeType}:                   _GotoState104Action,
	{_State272, AtomTypeType}:                            _GotoState99Action,
	{_State272, ParseErrorTypeType}:                      _GotoState105Action,
	{_State272, ReturnableTypeType}:                      _GotoState108Action,
	{_State272, PrefixedTypeType}:                        _GotoState107Action,
	{_State272, PrefixTypeOpType}:                        _GotoState106Action,
	{_State272, ValueTypeType}:                           _GotoState368Action,
	{_State272, TraitOpTypeType}:                         _GotoState110Action,
	{_State272, ImplicitStructDefType}:                   _GotoState103Action,
	{_State272, ExplicitStructDefType}:                   _GotoState52Action,
	{_State272, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State272, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State272, TraitDefType}:                            _GotoState109Action,
	{_State272, FuncTypeType}:                            _GotoState101Action,
	{_State273, AddToken}:                                _GotoState191Action,
	{_State273, SubToken}:                                _GotoState196Action,
	{_State273, MulToken}:                                _GotoState194Action,
	{_State273, TraitOpType}:                             _GotoState197Action,
	{_State274, IdentifierToken}:                         _GotoState369Action,
	{_State275, IdentifierToken}:                         _GotoState93Action,
	{_State275, StructToken}:                             _GotoState34Action,
	{_State275, EnumToken}:                               _GotoState90Action,
	{_State275, TraitToken}:                              _GotoState98Action,
	{_State275, FuncToken}:                               _GotoState92Action,
	{_State275, LparenToken}:                             _GotoState94Action,
	{_State275, LbracketToken}:                           _GotoState26Action,
	{_State275, DotToken}:                                _GotoState89Action,
	{_State275, QuestionToken}:                           _GotoState96Action,
	{_State275, ExclaimToken}:                            _GotoState91Action,
	{_State275, TildeTildeToken}:                         _GotoState97Action,
	{_State275, BitNegToken}:                             _GotoState88Action,
	{_State275, BitAndToken}:                             _GotoState87Action,
	{_State275, ParseErrorToken}:                         _GotoState95Action,
	{_State275, InitializableTypeType}:                   _GotoState104Action,
	{_State275, AtomTypeType}:                            _GotoState99Action,
	{_State275, ParseErrorTypeType}:                      _GotoState105Action,
	{_State275, ReturnableTypeType}:                      _GotoState371Action,
	{_State275, PrefixedTypeType}:                        _GotoState107Action,
	{_State275, PrefixTypeOpType}:                        _GotoState106Action,
	{_State275, ImplicitStructDefType}:                   _GotoState103Action,
	{_State275, ExplicitStructDefType}:                   _GotoState52Action,
	{_State275, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State275, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State275, TraitDefType}:                            _GotoState109Action,
	{_State275, ReturnTypeType}:                          _GotoState370Action,
	{_State275, FuncTypeType}:                            _GotoState101Action,
	{_State276, IdentifierToken}:                         _GotoState170Action,
	{_State276, ParameterDefType}:                        _GotoState372Action,
	{_State277, NewlinesToken}:                           _GotoState373Action,
	{_State277, OrToken}:                                 _GotoState374Action,
	{_State278, RparenToken}:                             _GotoState375Action,
	{_State279, AssignToken}:                             _GotoState292Action,
	{_State280, NewlinesToken}:                           _GotoState376Action,
	{_State280, OrToken}:                                 _GotoState377Action,
	{_State281, IdentifierToken}:                         _GotoState93Action,
	{_State281, StructToken}:                             _GotoState34Action,
	{_State281, EnumToken}:                               _GotoState90Action,
	{_State281, TraitToken}:                              _GotoState98Action,
	{_State281, FuncToken}:                               _GotoState92Action,
	{_State281, LparenToken}:                             _GotoState94Action,
	{_State281, LbracketToken}:                           _GotoState26Action,
	{_State281, DotToken}:                                _GotoState89Action,
	{_State281, QuestionToken}:                           _GotoState96Action,
	{_State281, ExclaimToken}:                            _GotoState91Action,
	{_State281, TildeTildeToken}:                         _GotoState97Action,
	{_State281, BitNegToken}:                             _GotoState88Action,
	{_State281, BitAndToken}:                             _GotoState87Action,
	{_State281, ParseErrorToken}:                         _GotoState95Action,
	{_State281, InitializableTypeType}:                   _GotoState104Action,
	{_State281, AtomTypeType}:                            _GotoState99Action,
	{_State281, ParseErrorTypeType}:                      _GotoState105Action,
	{_State281, ReturnableTypeType}:                      _GotoState108Action,
	{_State281, PrefixedTypeType}:                        _GotoState107Action,
	{_State281, PrefixTypeOpType}:                        _GotoState106Action,
	{_State281, ValueTypeType}:                           _GotoState378Action,
	{_State281, TraitOpTypeType}:                         _GotoState110Action,
	{_State281, ImplicitStructDefType}:                   _GotoState103Action,
	{_State281, ExplicitStructDefType}:                   _GotoState52Action,
	{_State281, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State281, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State281, TraitDefType}:                            _GotoState109Action,
	{_State281, FuncTypeType}:                            _GotoState101Action,
	{_State282, IdentifierToken}:                         _GotoState93Action,
	{_State282, StructToken}:                             _GotoState34Action,
	{_State282, EnumToken}:                               _GotoState90Action,
	{_State282, TraitToken}:                              _GotoState98Action,
	{_State282, FuncToken}:                               _GotoState92Action,
	{_State282, LparenToken}:                             _GotoState94Action,
	{_State282, LbracketToken}:                           _GotoState26Action,
	{_State282, DotToken}:                                _GotoState288Action,
	{_State282, QuestionToken}:                           _GotoState96Action,
	{_State282, ExclaimToken}:                            _GotoState91Action,
	{_State282, DollarLbracketToken}:                     _GotoState121Action,
	{_State282, DotDotDotToken}:                          _GotoState379Action,
	{_State282, TildeTildeToken}:                         _GotoState97Action,
	{_State282, BitNegToken}:                             _GotoState88Action,
	{_State282, BitAndToken}:                             _GotoState87Action,
	{_State282, ParseErrorToken}:                         _GotoState95Action,
	{_State282, OptionalGenericBindingType}:              _GotoState179Action,
	{_State282, InitializableTypeType}:                   _GotoState104Action,
	{_State282, AtomTypeType}:                            _GotoState99Action,
	{_State282, ParseErrorTypeType}:                      _GotoState105Action,
	{_State282, ReturnableTypeType}:                      _GotoState108Action,
	{_State282, PrefixedTypeType}:                        _GotoState107Action,
	{_State282, PrefixTypeOpType}:                        _GotoState106Action,
	{_State282, ValueTypeType}:                           _GotoState380Action,
	{_State282, TraitOpTypeType}:                         _GotoState110Action,
	{_State282, ImplicitStructDefType}:                   _GotoState103Action,
	{_State282, ExplicitStructDefType}:                   _GotoState52Action,
	{_State282, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State282, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State282, TraitDefType}:                            _GotoState109Action,
	{_State282, FuncTypeType}:                            _GotoState101Action,
	{_State283, RparenToken}:                             _GotoState381Action,
	{_State285, CommaToken}:                              _GotoState382Action,
	{_State286, AddToken}:                                _GotoState191Action,
	{_State286, SubToken}:                                _GotoState196Action,
	{_State286, MulToken}:                                _GotoState194Action,
	{_State286, TraitOpType}:                             _GotoState197Action,
	{_State287, DollarLbracketToken}:                     _GotoState121Action,
	{_State287, OptionalGenericBindingType}:              _GotoState383Action,
	{_State288, IdentifierToken}:                         _GotoState287Action,
	{_State288, DollarLbracketToken}:                     _GotoState121Action,
	{_State288, OptionalGenericBindingType}:              _GotoState175Action,
	{_State289, AddToken}:                                _GotoState191Action,
	{_State289, SubToken}:                                _GotoState196Action,
	{_State289, MulToken}:                                _GotoState194Action,
	{_State289, TraitOpType}:                             _GotoState197Action,
	{_State290, IdentifierToken}:                         _GotoState384Action,
	{_State291, IdentifierToken}:                         _GotoState180Action,
	{_State291, UnsafeToken}:                             _GotoState181Action,
	{_State291, StructToken}:                             _GotoState34Action,
	{_State291, EnumToken}:                               _GotoState90Action,
	{_State291, TraitToken}:                              _GotoState98Action,
	{_State291, FuncToken}:                               _GotoState92Action,
	{_State291, LparenToken}:                             _GotoState94Action,
	{_State291, LbracketToken}:                           _GotoState26Action,
	{_State291, DotToken}:                                _GotoState89Action,
	{_State291, QuestionToken}:                           _GotoState96Action,
	{_State291, ExclaimToken}:                            _GotoState91Action,
	{_State291, TildeTildeToken}:                         _GotoState97Action,
	{_State291, BitNegToken}:                             _GotoState88Action,
	{_State291, BitAndToken}:                             _GotoState87Action,
	{_State291, ParseErrorToken}:                         _GotoState95Action,
	{_State291, UnsafeStatementType}:                     _GotoState187Action,
	{_State291, InitializableTypeType}:                   _GotoState104Action,
	{_State291, AtomTypeType}:                            _GotoState99Action,
	{_State291, ParseErrorTypeType}:                      _GotoState105Action,
	{_State291, ReturnableTypeType}:                      _GotoState108Action,
	{_State291, PrefixedTypeType}:                        _GotoState107Action,
	{_State291, PrefixTypeOpType}:                        _GotoState106Action,
	{_State291, ValueTypeType}:                           _GotoState188Action,
	{_State291, TraitOpTypeType}:                         _GotoState110Action,
	{_State291, FieldDefType}:                            _GotoState279Action,
	{_State291, ImplicitStructDefType}:                   _GotoState103Action,
	{_State291, ExplicitStructDefType}:                   _GotoState52Action,
	{_State291, EnumValueDefType}:                        _GotoState385Action,
	{_State291, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State291, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State291, TraitDefType}:                            _GotoState109Action,
	{_State291, FuncTypeType}:                            _GotoState101Action,
	{_State292, DefaultToken}:                            _GotoState386Action,
	{_State293, IdentifierToken}:                         _GotoState180Action,
	{_State293, UnsafeToken}:                             _GotoState181Action,
	{_State293, StructToken}:                             _GotoState34Action,
	{_State293, EnumToken}:                               _GotoState90Action,
	{_State293, TraitToken}:                              _GotoState98Action,
	{_State293, FuncToken}:                               _GotoState92Action,
	{_State293, LparenToken}:                             _GotoState94Action,
	{_State293, LbracketToken}:                           _GotoState26Action,
	{_State293, DotToken}:                                _GotoState89Action,
	{_State293, QuestionToken}:                           _GotoState96Action,
	{_State293, ExclaimToken}:                            _GotoState91Action,
	{_State293, TildeTildeToken}:                         _GotoState97Action,
	{_State293, BitNegToken}:                             _GotoState88Action,
	{_State293, BitAndToken}:                             _GotoState87Action,
	{_State293, ParseErrorToken}:                         _GotoState95Action,
	{_State293, UnsafeStatementType}:                     _GotoState187Action,
	{_State293, InitializableTypeType}:                   _GotoState104Action,
	{_State293, AtomTypeType}:                            _GotoState99Action,
	{_State293, ParseErrorTypeType}:                      _GotoState105Action,
	{_State293, ReturnableTypeType}:                      _GotoState108Action,
	{_State293, PrefixedTypeType}:                        _GotoState107Action,
	{_State293, PrefixTypeOpType}:                        _GotoState106Action,
	{_State293, ValueTypeType}:                           _GotoState188Action,
	{_State293, TraitOpTypeType}:                         _GotoState110Action,
	{_State293, FieldDefType}:                            _GotoState279Action,
	{_State293, ImplicitStructDefType}:                   _GotoState103Action,
	{_State293, ExplicitStructDefType}:                   _GotoState52Action,
	{_State293, EnumValueDefType}:                        _GotoState387Action,
	{_State293, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State293, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State293, TraitDefType}:                            _GotoState109Action,
	{_State293, FuncTypeType}:                            _GotoState101Action,
	{_State295, IdentifierToken}:                         _GotoState180Action,
	{_State295, UnsafeToken}:                             _GotoState181Action,
	{_State295, StructToken}:                             _GotoState34Action,
	{_State295, EnumToken}:                               _GotoState90Action,
	{_State295, TraitToken}:                              _GotoState98Action,
	{_State295, FuncToken}:                               _GotoState92Action,
	{_State295, LparenToken}:                             _GotoState94Action,
	{_State295, LbracketToken}:                           _GotoState26Action,
	{_State295, DotToken}:                                _GotoState89Action,
	{_State295, QuestionToken}:                           _GotoState96Action,
	{_State295, ExclaimToken}:                            _GotoState91Action,
	{_State295, TildeTildeToken}:                         _GotoState97Action,
	{_State295, BitNegToken}:                             _GotoState88Action,
	{_State295, BitAndToken}:                             _GotoState87Action,
	{_State295, ParseErrorToken}:                         _GotoState95Action,
	{_State295, UnsafeStatementType}:                     _GotoState187Action,
	{_State295, InitializableTypeType}:                   _GotoState104Action,
	{_State295, AtomTypeType}:                            _GotoState99Action,
	{_State295, ParseErrorTypeType}:                      _GotoState105Action,
	{_State295, ReturnableTypeType}:                      _GotoState108Action,
	{_State295, PrefixedTypeType}:                        _GotoState107Action,
	{_State295, PrefixTypeOpType}:                        _GotoState106Action,
	{_State295, ValueTypeType}:                           _GotoState188Action,
	{_State295, TraitOpTypeType}:                         _GotoState110Action,
	{_State295, FieldDefType}:                            _GotoState388Action,
	{_State295, ImplicitStructDefType}:                   _GotoState103Action,
	{_State295, ExplicitStructDefType}:                   _GotoState52Action,
	{_State295, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State295, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State295, TraitDefType}:                            _GotoState109Action,
	{_State295, FuncTypeType}:                            _GotoState101Action,
	{_State297, IdentifierToken}:                         _GotoState389Action,
	{_State297, LparenToken}:                             _GotoState177Action,
	{_State300, RparenToken}:                             _GotoState390Action,
	{_State301, NewlinesToken}:                           _GotoState392Action,
	{_State301, CommaToken}:                              _GotoState391Action,
	{_State303, RbracketToken}:                           _GotoState393Action,
	{_State303, AddToken}:                                _GotoState191Action,
	{_State303, SubToken}:                                _GotoState196Action,
	{_State303, MulToken}:                                _GotoState194Action,
	{_State303, TraitOpType}:                             _GotoState197Action,
	{_State304, RbracketToken}:                           _GotoState394Action,
	{_State311, IdentifierToken}:                         _GotoState180Action,
	{_State311, UnsafeToken}:                             _GotoState181Action,
	{_State311, StructToken}:                             _GotoState34Action,
	{_State311, EnumToken}:                               _GotoState90Action,
	{_State311, TraitToken}:                              _GotoState98Action,
	{_State311, FuncToken}:                               _GotoState92Action,
	{_State311, LparenToken}:                             _GotoState94Action,
	{_State311, LbracketToken}:                           _GotoState26Action,
	{_State311, DotToken}:                                _GotoState89Action,
	{_State311, QuestionToken}:                           _GotoState96Action,
	{_State311, ExclaimToken}:                            _GotoState91Action,
	{_State311, TildeTildeToken}:                         _GotoState97Action,
	{_State311, BitNegToken}:                             _GotoState88Action,
	{_State311, BitAndToken}:                             _GotoState87Action,
	{_State311, ParseErrorToken}:                         _GotoState95Action,
	{_State311, UnsafeStatementType}:                     _GotoState187Action,
	{_State311, InitializableTypeType}:                   _GotoState104Action,
	{_State311, AtomTypeType}:                            _GotoState99Action,
	{_State311, ParseErrorTypeType}:                      _GotoState105Action,
	{_State311, ReturnableTypeType}:                      _GotoState108Action,
	{_State311, PrefixedTypeType}:                        _GotoState107Action,
	{_State311, PrefixTypeOpType}:                        _GotoState106Action,
	{_State311, ValueTypeType}:                           _GotoState188Action,
	{_State311, TraitOpTypeType}:                         _GotoState110Action,
	{_State311, FieldDefType}:                            _GotoState395Action,
	{_State311, ImplicitStructDefType}:                   _GotoState103Action,
	{_State311, ExplicitStructDefType}:                   _GotoState52Action,
	{_State311, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State311, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State311, TraitDefType}:                            _GotoState109Action,
	{_State311, FuncTypeType}:                            _GotoState101Action,
	{_State312, IdentifierToken}:                         _GotoState180Action,
	{_State312, UnsafeToken}:                             _GotoState181Action,
	{_State312, StructToken}:                             _GotoState34Action,
	{_State312, EnumToken}:                               _GotoState90Action,
	{_State312, TraitToken}:                              _GotoState98Action,
	{_State312, FuncToken}:                               _GotoState92Action,
	{_State312, LparenToken}:                             _GotoState94Action,
	{_State312, LbracketToken}:                           _GotoState26Action,
	{_State312, DotToken}:                                _GotoState89Action,
	{_State312, QuestionToken}:                           _GotoState96Action,
	{_State312, ExclaimToken}:                            _GotoState91Action,
	{_State312, TildeTildeToken}:                         _GotoState97Action,
	{_State312, BitNegToken}:                             _GotoState88Action,
	{_State312, BitAndToken}:                             _GotoState87Action,
	{_State312, ParseErrorToken}:                         _GotoState95Action,
	{_State312, UnsafeStatementType}:                     _GotoState187Action,
	{_State312, InitializableTypeType}:                   _GotoState104Action,
	{_State312, AtomTypeType}:                            _GotoState99Action,
	{_State312, ParseErrorTypeType}:                      _GotoState105Action,
	{_State312, ReturnableTypeType}:                      _GotoState108Action,
	{_State312, PrefixedTypeType}:                        _GotoState107Action,
	{_State312, PrefixTypeOpType}:                        _GotoState106Action,
	{_State312, ValueTypeType}:                           _GotoState188Action,
	{_State312, TraitOpTypeType}:                         _GotoState110Action,
	{_State312, FieldDefType}:                            _GotoState396Action,
	{_State312, ImplicitStructDefType}:                   _GotoState103Action,
	{_State312, ExplicitStructDefType}:                   _GotoState52Action,
	{_State312, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State312, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State312, TraitDefType}:                            _GotoState109Action,
	{_State312, FuncTypeType}:                            _GotoState101Action,
	{_State314, IdentifierToken}:                         _GotoState93Action,
	{_State314, StructToken}:                             _GotoState34Action,
	{_State314, EnumToken}:                               _GotoState90Action,
	{_State314, TraitToken}:                              _GotoState98Action,
	{_State314, FuncToken}:                               _GotoState92Action,
	{_State314, LparenToken}:                             _GotoState94Action,
	{_State314, LbracketToken}:                           _GotoState26Action,
	{_State314, DotToken}:                                _GotoState89Action,
	{_State314, QuestionToken}:                           _GotoState96Action,
	{_State314, ExclaimToken}:                            _GotoState91Action,
	{_State314, TildeTildeToken}:                         _GotoState97Action,
	{_State314, BitNegToken}:                             _GotoState88Action,
	{_State314, BitAndToken}:                             _GotoState87Action,
	{_State314, ParseErrorToken}:                         _GotoState95Action,
	{_State314, InitializableTypeType}:                   _GotoState104Action,
	{_State314, AtomTypeType}:                            _GotoState99Action,
	{_State314, ParseErrorTypeType}:                      _GotoState105Action,
	{_State314, ReturnableTypeType}:                      _GotoState108Action,
	{_State314, PrefixedTypeType}:                        _GotoState107Action,
	{_State314, PrefixTypeOpType}:                        _GotoState106Action,
	{_State314, ValueTypeType}:                           _GotoState397Action,
	{_State314, TraitOpTypeType}:                         _GotoState110Action,
	{_State314, ImplicitStructDefType}:                   _GotoState103Action,
	{_State314, ExplicitStructDefType}:                   _GotoState52Action,
	{_State314, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State314, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State314, TraitDefType}:                            _GotoState109Action,
	{_State314, FuncTypeType}:                            _GotoState101Action,
	{_State317, RparenToken}:                             _GotoState398Action,
	{_State319, IntegerLiteralToken}:                     _GotoState24Action,
	{_State319, FloatLiteralToken}:                       _GotoState20Action,
	{_State319, RuneLiteralToken}:                        _GotoState32Action,
	{_State319, StringLiteralToken}:                      _GotoState33Action,
	{_State319, IdentifierToken}:                         _GotoState23Action,
	{_State319, TrueToken}:                               _GotoState36Action,
	{_State319, FalseToken}:                              _GotoState19Action,
	{_State319, StructToken}:                             _GotoState34Action,
	{_State319, FuncToken}:                               _GotoState21Action,
	{_State319, VarToken}:                                _GotoState37Action,
	{_State319, LetToken}:                                _GotoState27Action,
	{_State319, NotToken}:                                _GotoState30Action,
	{_State319, LabelDeclToken}:                          _GotoState25Action,
	{_State319, LparenToken}:                             _GotoState28Action,
	{_State319, LbracketToken}:                           _GotoState26Action,
	{_State319, SubToken}:                                _GotoState35Action,
	{_State319, MulToken}:                                _GotoState29Action,
	{_State319, BitNegToken}:                             _GotoState18Action,
	{_State319, BitAndToken}:                             _GotoState17Action,
	{_State319, GreaterToken}:                            _GotoState22Action,
	{_State319, ParseErrorToken}:                         _GotoState31Action,
	{_State319, VarDeclPatternType}:                      _GotoState69Action,
	{_State319, VarOrLetType}:                            _GotoState70Action,
	{_State319, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State319, SequenceExprType}:                        _GotoState399Action,
	{_State319, CallExprType}:                            _GotoState50Action,
	{_State319, AtomExprType}:                            _GotoState43Action,
	{_State319, ParseErrorExprType}:                      _GotoState62Action,
	{_State319, LiteralExprType}:                         _GotoState58Action,
	{_State319, IdentifierExprType}:                      _GotoState53Action,
	{_State319, BlockExprType}:                           _GotoState49Action,
	{_State319, InitializeExprType}:                      _GotoState57Action,
	{_State319, ImplicitStructExprType}:                  _GotoState54Action,
	{_State319, AccessibleExprType}:                      _GotoState39Action,
	{_State319, AccessExprType}:                          _GotoState38Action,
	{_State319, IndexExprType}:                           _GotoState55Action,
	{_State319, PostfixableExprType}:                     _GotoState64Action,
	{_State319, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State319, PrefixableExprType}:                      _GotoState67Action,
	{_State319, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State319, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State319, MulExprType}:                             _GotoState59Action,
	{_State319, BinaryMulExprType}:                       _GotoState47Action,
	{_State319, AddExprType}:                             _GotoState40Action,
	{_State319, BinaryAddExprType}:                       _GotoState44Action,
	{_State319, CmpExprType}:                             _GotoState51Action,
	{_State319, BinaryCmpExprType}:                       _GotoState46Action,
	{_State319, AndExprType}:                             _GotoState41Action,
	{_State319, BinaryAndExprType}:                       _GotoState45Action,
	{_State319, OrExprType}:                              _GotoState61Action,
	{_State319, BinaryOrExprType}:                        _GotoState48Action,
	{_State319, InitializableTypeType}:                   _GotoState56Action,
	{_State319, ExplicitStructDefType}:                   _GotoState52Action,
	{_State319, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State320, IntegerLiteralToken}:                     _GotoState24Action,
	{_State320, FloatLiteralToken}:                       _GotoState20Action,
	{_State320, RuneLiteralToken}:                        _GotoState32Action,
	{_State320, StringLiteralToken}:                      _GotoState33Action,
	{_State320, IdentifierToken}:                         _GotoState23Action,
	{_State320, TrueToken}:                               _GotoState36Action,
	{_State320, FalseToken}:                              _GotoState19Action,
	{_State320, StructToken}:                             _GotoState34Action,
	{_State320, FuncToken}:                               _GotoState21Action,
	{_State320, VarToken}:                                _GotoState37Action,
	{_State320, LetToken}:                                _GotoState27Action,
	{_State320, NotToken}:                                _GotoState30Action,
	{_State320, LabelDeclToken}:                          _GotoState25Action,
	{_State320, LparenToken}:                             _GotoState28Action,
	{_State320, LbracketToken}:                           _GotoState26Action,
	{_State320, SubToken}:                                _GotoState35Action,
	{_State320, MulToken}:                                _GotoState29Action,
	{_State320, BitNegToken}:                             _GotoState18Action,
	{_State320, BitAndToken}:                             _GotoState17Action,
	{_State320, GreaterToken}:                            _GotoState22Action,
	{_State320, ParseErrorToken}:                         _GotoState31Action,
	{_State320, VarDeclPatternType}:                      _GotoState69Action,
	{_State320, VarOrLetType}:                            _GotoState70Action,
	{_State320, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State320, SequenceExprType}:                        _GotoState400Action,
	{_State320, CallExprType}:                            _GotoState50Action,
	{_State320, AtomExprType}:                            _GotoState43Action,
	{_State320, ParseErrorExprType}:                      _GotoState62Action,
	{_State320, LiteralExprType}:                         _GotoState58Action,
	{_State320, IdentifierExprType}:                      _GotoState53Action,
	{_State320, BlockExprType}:                           _GotoState49Action,
	{_State320, InitializeExprType}:                      _GotoState57Action,
	{_State320, ImplicitStructExprType}:                  _GotoState54Action,
	{_State320, AccessibleExprType}:                      _GotoState39Action,
	{_State320, AccessExprType}:                          _GotoState38Action,
	{_State320, IndexExprType}:                           _GotoState55Action,
	{_State320, PostfixableExprType}:                     _GotoState64Action,
	{_State320, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State320, PrefixableExprType}:                      _GotoState67Action,
	{_State320, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State320, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State320, MulExprType}:                             _GotoState59Action,
	{_State320, BinaryMulExprType}:                       _GotoState47Action,
	{_State320, AddExprType}:                             _GotoState40Action,
	{_State320, BinaryAddExprType}:                       _GotoState44Action,
	{_State320, CmpExprType}:                             _GotoState51Action,
	{_State320, BinaryCmpExprType}:                       _GotoState46Action,
	{_State320, AndExprType}:                             _GotoState41Action,
	{_State320, BinaryAndExprType}:                       _GotoState45Action,
	{_State320, OrExprType}:                              _GotoState61Action,
	{_State320, BinaryOrExprType}:                        _GotoState48Action,
	{_State320, InitializableTypeType}:                   _GotoState56Action,
	{_State320, ExplicitStructDefType}:                   _GotoState52Action,
	{_State320, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State321, IntegerLiteralToken}:                     _GotoState24Action,
	{_State321, FloatLiteralToken}:                       _GotoState20Action,
	{_State321, RuneLiteralToken}:                        _GotoState32Action,
	{_State321, StringLiteralToken}:                      _GotoState33Action,
	{_State321, IdentifierToken}:                         _GotoState23Action,
	{_State321, TrueToken}:                               _GotoState36Action,
	{_State321, FalseToken}:                              _GotoState19Action,
	{_State321, StructToken}:                             _GotoState34Action,
	{_State321, FuncToken}:                               _GotoState21Action,
	{_State321, VarToken}:                                _GotoState37Action,
	{_State321, LetToken}:                                _GotoState27Action,
	{_State321, NotToken}:                                _GotoState30Action,
	{_State321, LabelDeclToken}:                          _GotoState25Action,
	{_State321, LparenToken}:                             _GotoState28Action,
	{_State321, LbracketToken}:                           _GotoState26Action,
	{_State321, SubToken}:                                _GotoState35Action,
	{_State321, MulToken}:                                _GotoState29Action,
	{_State321, BitNegToken}:                             _GotoState18Action,
	{_State321, BitAndToken}:                             _GotoState17Action,
	{_State321, GreaterToken}:                            _GotoState22Action,
	{_State321, ParseErrorToken}:                         _GotoState31Action,
	{_State321, VarDeclPatternType}:                      _GotoState69Action,
	{_State321, VarOrLetType}:                            _GotoState70Action,
	{_State321, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State321, SequenceExprType}:                        _GotoState401Action,
	{_State321, CallExprType}:                            _GotoState50Action,
	{_State321, AtomExprType}:                            _GotoState43Action,
	{_State321, ParseErrorExprType}:                      _GotoState62Action,
	{_State321, LiteralExprType}:                         _GotoState58Action,
	{_State321, IdentifierExprType}:                      _GotoState53Action,
	{_State321, BlockExprType}:                           _GotoState49Action,
	{_State321, InitializeExprType}:                      _GotoState57Action,
	{_State321, ImplicitStructExprType}:                  _GotoState54Action,
	{_State321, AccessibleExprType}:                      _GotoState39Action,
	{_State321, AccessExprType}:                          _GotoState38Action,
	{_State321, IndexExprType}:                           _GotoState55Action,
	{_State321, PostfixableExprType}:                     _GotoState64Action,
	{_State321, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State321, PrefixableExprType}:                      _GotoState67Action,
	{_State321, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State321, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State321, MulExprType}:                             _GotoState59Action,
	{_State321, BinaryMulExprType}:                       _GotoState47Action,
	{_State321, AddExprType}:                             _GotoState40Action,
	{_State321, BinaryAddExprType}:                       _GotoState44Action,
	{_State321, CmpExprType}:                             _GotoState51Action,
	{_State321, BinaryCmpExprType}:                       _GotoState46Action,
	{_State321, AndExprType}:                             _GotoState41Action,
	{_State321, BinaryAndExprType}:                       _GotoState45Action,
	{_State321, OrExprType}:                              _GotoState61Action,
	{_State321, BinaryOrExprType}:                        _GotoState48Action,
	{_State321, InitializableTypeType}:                   _GotoState56Action,
	{_State321, ExplicitStructDefType}:                   _GotoState52Action,
	{_State321, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State322, IntegerLiteralToken}:                     _GotoState24Action,
	{_State322, FloatLiteralToken}:                       _GotoState20Action,
	{_State322, RuneLiteralToken}:                        _GotoState32Action,
	{_State322, StringLiteralToken}:                      _GotoState33Action,
	{_State322, IdentifierToken}:                         _GotoState23Action,
	{_State322, TrueToken}:                               _GotoState36Action,
	{_State322, FalseToken}:                              _GotoState19Action,
	{_State322, StructToken}:                             _GotoState34Action,
	{_State322, FuncToken}:                               _GotoState21Action,
	{_State322, VarToken}:                                _GotoState37Action,
	{_State322, LetToken}:                                _GotoState27Action,
	{_State322, NotToken}:                                _GotoState30Action,
	{_State322, LabelDeclToken}:                          _GotoState25Action,
	{_State322, LparenToken}:                             _GotoState28Action,
	{_State322, LbracketToken}:                           _GotoState26Action,
	{_State322, SubToken}:                                _GotoState35Action,
	{_State322, MulToken}:                                _GotoState29Action,
	{_State322, BitNegToken}:                             _GotoState18Action,
	{_State322, BitAndToken}:                             _GotoState17Action,
	{_State322, GreaterToken}:                            _GotoState22Action,
	{_State322, ParseErrorToken}:                         _GotoState31Action,
	{_State322, VarDeclPatternType}:                      _GotoState69Action,
	{_State322, VarOrLetType}:                            _GotoState70Action,
	{_State322, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State322, SequenceExprType}:                        _GotoState403Action,
	{_State322, OptionalSequenceExprType}:                _GotoState402Action,
	{_State322, CallExprType}:                            _GotoState50Action,
	{_State322, AtomExprType}:                            _GotoState43Action,
	{_State322, ParseErrorExprType}:                      _GotoState62Action,
	{_State322, LiteralExprType}:                         _GotoState58Action,
	{_State322, IdentifierExprType}:                      _GotoState53Action,
	{_State322, BlockExprType}:                           _GotoState49Action,
	{_State322, InitializeExprType}:                      _GotoState57Action,
	{_State322, ImplicitStructExprType}:                  _GotoState54Action,
	{_State322, AccessibleExprType}:                      _GotoState39Action,
	{_State322, AccessExprType}:                          _GotoState38Action,
	{_State322, IndexExprType}:                           _GotoState55Action,
	{_State322, PostfixableExprType}:                     _GotoState64Action,
	{_State322, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State322, PrefixableExprType}:                      _GotoState67Action,
	{_State322, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State322, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State322, MulExprType}:                             _GotoState59Action,
	{_State322, BinaryMulExprType}:                       _GotoState47Action,
	{_State322, AddExprType}:                             _GotoState40Action,
	{_State322, BinaryAddExprType}:                       _GotoState44Action,
	{_State322, CmpExprType}:                             _GotoState51Action,
	{_State322, BinaryCmpExprType}:                       _GotoState46Action,
	{_State322, AndExprType}:                             _GotoState41Action,
	{_State322, BinaryAndExprType}:                       _GotoState45Action,
	{_State322, OrExprType}:                              _GotoState61Action,
	{_State322, BinaryOrExprType}:                        _GotoState48Action,
	{_State322, InitializableTypeType}:                   _GotoState56Action,
	{_State322, ExplicitStructDefType}:                   _GotoState52Action,
	{_State322, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State323, LbraceToken}:                             _GotoState72Action,
	{_State323, StatementBlockType}:                      _GotoState404Action,
	{_State324, IdentifierToken}:                         _GotoState405Action,
	{_State325, DotToken}:                                _GotoState406Action,
	{_State328, CommaToken}:                              _GotoState408Action,
	{_State328, AssignToken}:                             _GotoState407Action,
	{_State330, ElseToken}:                               _GotoState409Action,
	{_State332, IdentifierToken}:                         _GotoState157Action,
	{_State332, LparenToken}:                             _GotoState158Action,
	{_State332, VarPatternType}:                          _GotoState410Action,
	{_State332, TuplePatternType}:                        _GotoState159Action,
	{_State333, IdentifierToken}:                         _GotoState227Action,
	{_State333, LparenToken}:                             _GotoState158Action,
	{_State333, DotDotDotToken}:                          _GotoState226Action,
	{_State333, VarPatternType}:                          _GotoState230Action,
	{_State333, TuplePatternType}:                        _GotoState159Action,
	{_State333, FieldVarPatternType}:                     _GotoState411Action,
	{_State335, CommaToken}:                              _GotoState408Action,
	{_State335, ColonToken}:                              _GotoState412Action,
	{_State336, IntegerLiteralToken}:                     _GotoState24Action,
	{_State336, FloatLiteralToken}:                       _GotoState20Action,
	{_State336, RuneLiteralToken}:                        _GotoState32Action,
	{_State336, StringLiteralToken}:                      _GotoState33Action,
	{_State336, IdentifierToken}:                         _GotoState23Action,
	{_State336, TrueToken}:                               _GotoState36Action,
	{_State336, FalseToken}:                              _GotoState19Action,
	{_State336, ReturnToken}:                             _GotoState242Action,
	{_State336, BreakToken}:                              _GotoState234Action,
	{_State336, ContinueToken}:                           _GotoState236Action,
	{_State336, FallthroughToken}:                        _GotoState239Action,
	{_State336, UnsafeToken}:                             _GotoState181Action,
	{_State336, StructToken}:                             _GotoState34Action,
	{_State336, FuncToken}:                               _GotoState21Action,
	{_State336, AsyncToken}:                              _GotoState233Action,
	{_State336, DeferToken}:                              _GotoState238Action,
	{_State336, VarToken}:                                _GotoState37Action,
	{_State336, LetToken}:                                _GotoState27Action,
	{_State336, NotToken}:                                _GotoState30Action,
	{_State336, LabelDeclToken}:                          _GotoState25Action,
	{_State336, LparenToken}:                             _GotoState28Action,
	{_State336, LbracketToken}:                           _GotoState26Action,
	{_State336, SubToken}:                                _GotoState35Action,
	{_State336, MulToken}:                                _GotoState29Action,
	{_State336, BitNegToken}:                             _GotoState18Action,
	{_State336, BitAndToken}:                             _GotoState17Action,
	{_State336, GreaterToken}:                            _GotoState22Action,
	{_State336, ParseErrorToken}:                         _GotoState31Action,
	{_State336, SimpleStatementBodyType}:                 _GotoState414Action,
	{_State336, OptionalSimpleStatementBodyType}:         _GotoState413Action,
	{_State336, ExpressionOrImproperStructStatementType}: _GotoState250Action,
	{_State336, CallbackOpType}:                          _GotoState247Action,
	{_State336, CallbackOpStatementType}:                 _GotoState248Action,
	{_State336, UnsafeStatementType}:                     _GotoState261Action,
	{_State336, JumpStatementType}:                       _GotoState254Action,
	{_State336, JumpTypeType}:                            _GotoState255Action,
	{_State336, ExpressionsType}:                         _GotoState251Action,
	{_State336, FallthroughStatementType}:                _GotoState252Action,
	{_State336, AssignStatementType}:                     _GotoState245Action,
	{_State336, UnaryOpAssignStatementType}:              _GotoState260Action,
	{_State336, BinaryOpAssignStatementType}:             _GotoState246Action,
	{_State336, VarDeclPatternType}:                      _GotoState69Action,
	{_State336, VarOrLetType}:                            _GotoState70Action,
	{_State336, AssignPatternType}:                       _GotoState244Action,
	{_State336, ExpressionType}:                          _GotoState249Action,
	{_State336, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State336, SequenceExprType}:                        _GotoState256Action,
	{_State336, CallExprType}:                            _GotoState50Action,
	{_State336, AtomExprType}:                            _GotoState43Action,
	{_State336, ParseErrorExprType}:                      _GotoState62Action,
	{_State336, LiteralExprType}:                         _GotoState58Action,
	{_State336, IdentifierExprType}:                      _GotoState53Action,
	{_State336, BlockExprType}:                           _GotoState49Action,
	{_State336, InitializeExprType}:                      _GotoState57Action,
	{_State336, ImplicitStructExprType}:                  _GotoState54Action,
	{_State336, AccessibleExprType}:                      _GotoState243Action,
	{_State336, AccessExprType}:                          _GotoState38Action,
	{_State336, IndexExprType}:                           _GotoState55Action,
	{_State336, PostfixableExprType}:                     _GotoState64Action,
	{_State336, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State336, PrefixableExprType}:                      _GotoState67Action,
	{_State336, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State336, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State336, MulExprType}:                             _GotoState59Action,
	{_State336, BinaryMulExprType}:                       _GotoState47Action,
	{_State336, AddExprType}:                             _GotoState40Action,
	{_State336, BinaryAddExprType}:                       _GotoState44Action,
	{_State336, CmpExprType}:                             _GotoState51Action,
	{_State336, BinaryCmpExprType}:                       _GotoState46Action,
	{_State336, AndExprType}:                             _GotoState41Action,
	{_State336, BinaryAndExprType}:                       _GotoState45Action,
	{_State336, OrExprType}:                              _GotoState61Action,
	{_State336, BinaryOrExprType}:                        _GotoState48Action,
	{_State336, InitializableTypeType}:                   _GotoState56Action,
	{_State336, ExplicitStructDefType}:                   _GotoState52Action,
	{_State336, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State337, StringLiteralToken}:                      _GotoState338Action,
	{_State337, ImportClauseType}:                        _GotoState415Action,
	{_State337, ImportClauseTerminalType}:                _GotoState416Action,
	{_State337, ImportClausesType}:                       _GotoState417Action,
	{_State338, AsToken}:                                 _GotoState418Action,
	{_State353, IntegerLiteralToken}:                     _GotoState24Action,
	{_State353, FloatLiteralToken}:                       _GotoState20Action,
	{_State353, RuneLiteralToken}:                        _GotoState32Action,
	{_State353, StringLiteralToken}:                      _GotoState33Action,
	{_State353, IdentifierToken}:                         _GotoState23Action,
	{_State353, TrueToken}:                               _GotoState36Action,
	{_State353, FalseToken}:                              _GotoState19Action,
	{_State353, StructToken}:                             _GotoState34Action,
	{_State353, FuncToken}:                               _GotoState21Action,
	{_State353, VarToken}:                                _GotoState37Action,
	{_State353, LetToken}:                                _GotoState27Action,
	{_State353, NotToken}:                                _GotoState30Action,
	{_State353, LabelDeclToken}:                          _GotoState25Action,
	{_State353, LparenToken}:                             _GotoState28Action,
	{_State353, LbracketToken}:                           _GotoState26Action,
	{_State353, SubToken}:                                _GotoState35Action,
	{_State353, MulToken}:                                _GotoState29Action,
	{_State353, BitNegToken}:                             _GotoState18Action,
	{_State353, BitAndToken}:                             _GotoState17Action,
	{_State353, GreaterToken}:                            _GotoState22Action,
	{_State353, ParseErrorToken}:                         _GotoState31Action,
	{_State353, VarDeclPatternType}:                      _GotoState69Action,
	{_State353, VarOrLetType}:                            _GotoState70Action,
	{_State353, ExpressionType}:                          _GotoState419Action,
	{_State353, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State353, SequenceExprType}:                        _GotoState68Action,
	{_State353, CallExprType}:                            _GotoState50Action,
	{_State353, AtomExprType}:                            _GotoState43Action,
	{_State353, ParseErrorExprType}:                      _GotoState62Action,
	{_State353, LiteralExprType}:                         _GotoState58Action,
	{_State353, IdentifierExprType}:                      _GotoState53Action,
	{_State353, BlockExprType}:                           _GotoState49Action,
	{_State353, InitializeExprType}:                      _GotoState57Action,
	{_State353, ImplicitStructExprType}:                  _GotoState54Action,
	{_State353, AccessibleExprType}:                      _GotoState39Action,
	{_State353, AccessExprType}:                          _GotoState38Action,
	{_State353, IndexExprType}:                           _GotoState55Action,
	{_State353, PostfixableExprType}:                     _GotoState64Action,
	{_State353, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State353, PrefixableExprType}:                      _GotoState67Action,
	{_State353, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State353, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State353, MulExprType}:                             _GotoState59Action,
	{_State353, BinaryMulExprType}:                       _GotoState47Action,
	{_State353, AddExprType}:                             _GotoState40Action,
	{_State353, BinaryAddExprType}:                       _GotoState44Action,
	{_State353, CmpExprType}:                             _GotoState51Action,
	{_State353, BinaryCmpExprType}:                       _GotoState46Action,
	{_State353, AndExprType}:                             _GotoState41Action,
	{_State353, BinaryAndExprType}:                       _GotoState45Action,
	{_State353, OrExprType}:                              _GotoState61Action,
	{_State353, BinaryOrExprType}:                        _GotoState48Action,
	{_State353, InitializableTypeType}:                   _GotoState56Action,
	{_State353, ExplicitStructDefType}:                   _GotoState52Action,
	{_State353, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State355, IntegerLiteralToken}:                     _GotoState24Action,
	{_State355, FloatLiteralToken}:                       _GotoState20Action,
	{_State355, RuneLiteralToken}:                        _GotoState32Action,
	{_State355, StringLiteralToken}:                      _GotoState33Action,
	{_State355, IdentifierToken}:                         _GotoState23Action,
	{_State355, TrueToken}:                               _GotoState36Action,
	{_State355, FalseToken}:                              _GotoState19Action,
	{_State355, StructToken}:                             _GotoState34Action,
	{_State355, FuncToken}:                               _GotoState21Action,
	{_State355, VarToken}:                                _GotoState37Action,
	{_State355, LetToken}:                                _GotoState27Action,
	{_State355, NotToken}:                                _GotoState30Action,
	{_State355, LabelDeclToken}:                          _GotoState25Action,
	{_State355, LparenToken}:                             _GotoState28Action,
	{_State355, LbracketToken}:                           _GotoState26Action,
	{_State355, SubToken}:                                _GotoState35Action,
	{_State355, MulToken}:                                _GotoState29Action,
	{_State355, BitNegToken}:                             _GotoState18Action,
	{_State355, BitAndToken}:                             _GotoState17Action,
	{_State355, GreaterToken}:                            _GotoState22Action,
	{_State355, ParseErrorToken}:                         _GotoState31Action,
	{_State355, VarDeclPatternType}:                      _GotoState69Action,
	{_State355, VarOrLetType}:                            _GotoState70Action,
	{_State355, ExpressionType}:                          _GotoState420Action,
	{_State355, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State355, SequenceExprType}:                        _GotoState68Action,
	{_State355, CallExprType}:                            _GotoState50Action,
	{_State355, AtomExprType}:                            _GotoState43Action,
	{_State355, ParseErrorExprType}:                      _GotoState62Action,
	{_State355, LiteralExprType}:                         _GotoState58Action,
	{_State355, IdentifierExprType}:                      _GotoState53Action,
	{_State355, BlockExprType}:                           _GotoState49Action,
	{_State355, InitializeExprType}:                      _GotoState57Action,
	{_State355, ImplicitStructExprType}:                  _GotoState54Action,
	{_State355, AccessibleExprType}:                      _GotoState39Action,
	{_State355, AccessExprType}:                          _GotoState38Action,
	{_State355, IndexExprType}:                           _GotoState55Action,
	{_State355, PostfixableExprType}:                     _GotoState64Action,
	{_State355, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State355, PrefixableExprType}:                      _GotoState67Action,
	{_State355, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State355, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State355, MulExprType}:                             _GotoState59Action,
	{_State355, BinaryMulExprType}:                       _GotoState47Action,
	{_State355, AddExprType}:                             _GotoState40Action,
	{_State355, BinaryAddExprType}:                       _GotoState44Action,
	{_State355, CmpExprType}:                             _GotoState51Action,
	{_State355, BinaryCmpExprType}:                       _GotoState46Action,
	{_State355, AndExprType}:                             _GotoState41Action,
	{_State355, BinaryAndExprType}:                       _GotoState45Action,
	{_State355, OrExprType}:                              _GotoState61Action,
	{_State355, BinaryOrExprType}:                        _GotoState48Action,
	{_State355, InitializableTypeType}:                   _GotoState56Action,
	{_State355, ExplicitStructDefType}:                   _GotoState52Action,
	{_State355, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State356, LbracketToken}:                           _GotoState123Action,
	{_State356, DotToken}:                                _GotoState122Action,
	{_State356, DollarLbracketToken}:                     _GotoState121Action,
	{_State356, OptionalGenericBindingType}:              _GotoState125Action,
	{_State358, IntegerLiteralToken}:                     _GotoState24Action,
	{_State358, FloatLiteralToken}:                       _GotoState20Action,
	{_State358, RuneLiteralToken}:                        _GotoState32Action,
	{_State358, StringLiteralToken}:                      _GotoState33Action,
	{_State358, IdentifierToken}:                         _GotoState23Action,
	{_State358, TrueToken}:                               _GotoState36Action,
	{_State358, FalseToken}:                              _GotoState19Action,
	{_State358, StructToken}:                             _GotoState34Action,
	{_State358, FuncToken}:                               _GotoState21Action,
	{_State358, VarToken}:                                _GotoState37Action,
	{_State358, LetToken}:                                _GotoState27Action,
	{_State358, NotToken}:                                _GotoState30Action,
	{_State358, LabelDeclToken}:                          _GotoState25Action,
	{_State358, LparenToken}:                             _GotoState28Action,
	{_State358, LbracketToken}:                           _GotoState26Action,
	{_State358, SubToken}:                                _GotoState35Action,
	{_State358, MulToken}:                                _GotoState29Action,
	{_State358, BitNegToken}:                             _GotoState18Action,
	{_State358, BitAndToken}:                             _GotoState17Action,
	{_State358, GreaterToken}:                            _GotoState22Action,
	{_State358, ParseErrorToken}:                         _GotoState31Action,
	{_State358, VarDeclPatternType}:                      _GotoState69Action,
	{_State358, VarOrLetType}:                            _GotoState70Action,
	{_State358, ExpressionType}:                          _GotoState421Action,
	{_State358, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State358, SequenceExprType}:                        _GotoState68Action,
	{_State358, CallExprType}:                            _GotoState50Action,
	{_State358, AtomExprType}:                            _GotoState43Action,
	{_State358, ParseErrorExprType}:                      _GotoState62Action,
	{_State358, LiteralExprType}:                         _GotoState58Action,
	{_State358, IdentifierExprType}:                      _GotoState53Action,
	{_State358, BlockExprType}:                           _GotoState49Action,
	{_State358, InitializeExprType}:                      _GotoState57Action,
	{_State358, ImplicitStructExprType}:                  _GotoState54Action,
	{_State358, AccessibleExprType}:                      _GotoState39Action,
	{_State358, AccessExprType}:                          _GotoState38Action,
	{_State358, IndexExprType}:                           _GotoState55Action,
	{_State358, PostfixableExprType}:                     _GotoState64Action,
	{_State358, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State358, PrefixableExprType}:                      _GotoState67Action,
	{_State358, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State358, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State358, MulExprType}:                             _GotoState59Action,
	{_State358, BinaryMulExprType}:                       _GotoState47Action,
	{_State358, AddExprType}:                             _GotoState40Action,
	{_State358, BinaryAddExprType}:                       _GotoState44Action,
	{_State358, CmpExprType}:                             _GotoState51Action,
	{_State358, BinaryCmpExprType}:                       _GotoState46Action,
	{_State358, AndExprType}:                             _GotoState41Action,
	{_State358, BinaryAndExprType}:                       _GotoState45Action,
	{_State358, OrExprType}:                              _GotoState61Action,
	{_State358, BinaryOrExprType}:                        _GotoState48Action,
	{_State358, InitializableTypeType}:                   _GotoState56Action,
	{_State358, ExplicitStructDefType}:                   _GotoState52Action,
	{_State358, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State359, IntegerLiteralToken}:                     _GotoState24Action,
	{_State359, FloatLiteralToken}:                       _GotoState20Action,
	{_State359, RuneLiteralToken}:                        _GotoState32Action,
	{_State359, StringLiteralToken}:                      _GotoState33Action,
	{_State359, IdentifierToken}:                         _GotoState23Action,
	{_State359, TrueToken}:                               _GotoState36Action,
	{_State359, FalseToken}:                              _GotoState19Action,
	{_State359, StructToken}:                             _GotoState34Action,
	{_State359, FuncToken}:                               _GotoState21Action,
	{_State359, VarToken}:                                _GotoState37Action,
	{_State359, LetToken}:                                _GotoState27Action,
	{_State359, NotToken}:                                _GotoState30Action,
	{_State359, LabelDeclToken}:                          _GotoState25Action,
	{_State359, LparenToken}:                             _GotoState28Action,
	{_State359, LbracketToken}:                           _GotoState26Action,
	{_State359, SubToken}:                                _GotoState35Action,
	{_State359, MulToken}:                                _GotoState29Action,
	{_State359, BitNegToken}:                             _GotoState18Action,
	{_State359, BitAndToken}:                             _GotoState17Action,
	{_State359, GreaterToken}:                            _GotoState22Action,
	{_State359, ParseErrorToken}:                         _GotoState31Action,
	{_State359, ExpressionsType}:                         _GotoState422Action,
	{_State359, VarDeclPatternType}:                      _GotoState69Action,
	{_State359, VarOrLetType}:                            _GotoState70Action,
	{_State359, ExpressionType}:                          _GotoState249Action,
	{_State359, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State359, SequenceExprType}:                        _GotoState68Action,
	{_State359, CallExprType}:                            _GotoState50Action,
	{_State359, AtomExprType}:                            _GotoState43Action,
	{_State359, ParseErrorExprType}:                      _GotoState62Action,
	{_State359, LiteralExprType}:                         _GotoState58Action,
	{_State359, IdentifierExprType}:                      _GotoState53Action,
	{_State359, BlockExprType}:                           _GotoState49Action,
	{_State359, InitializeExprType}:                      _GotoState57Action,
	{_State359, ImplicitStructExprType}:                  _GotoState54Action,
	{_State359, AccessibleExprType}:                      _GotoState39Action,
	{_State359, AccessExprType}:                          _GotoState38Action,
	{_State359, IndexExprType}:                           _GotoState55Action,
	{_State359, PostfixableExprType}:                     _GotoState64Action,
	{_State359, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State359, PrefixableExprType}:                      _GotoState67Action,
	{_State359, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State359, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State359, MulExprType}:                             _GotoState59Action,
	{_State359, BinaryMulExprType}:                       _GotoState47Action,
	{_State359, AddExprType}:                             _GotoState40Action,
	{_State359, BinaryAddExprType}:                       _GotoState44Action,
	{_State359, CmpExprType}:                             _GotoState51Action,
	{_State359, BinaryCmpExprType}:                       _GotoState46Action,
	{_State359, AndExprType}:                             _GotoState41Action,
	{_State359, BinaryAndExprType}:                       _GotoState45Action,
	{_State359, OrExprType}:                              _GotoState61Action,
	{_State359, BinaryOrExprType}:                        _GotoState48Action,
	{_State359, InitializableTypeType}:                   _GotoState56Action,
	{_State359, ExplicitStructDefType}:                   _GotoState52Action,
	{_State359, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State360, CommaToken}:                              _GotoState358Action,
	{_State363, AddToken}:                                _GotoState191Action,
	{_State363, SubToken}:                                _GotoState196Action,
	{_State363, MulToken}:                                _GotoState194Action,
	{_State363, TraitOpType}:                             _GotoState197Action,
	{_State364, IdentifierToken}:                         _GotoState265Action,
	{_State364, GenericParameterDefType}:                 _GotoState423Action,
	{_State366, IdentifierToken}:                         _GotoState93Action,
	{_State366, StructToken}:                             _GotoState34Action,
	{_State366, EnumToken}:                               _GotoState90Action,
	{_State366, TraitToken}:                              _GotoState98Action,
	{_State366, FuncToken}:                               _GotoState92Action,
	{_State366, LparenToken}:                             _GotoState94Action,
	{_State366, LbracketToken}:                           _GotoState26Action,
	{_State366, DotToken}:                                _GotoState89Action,
	{_State366, QuestionToken}:                           _GotoState96Action,
	{_State366, ExclaimToken}:                            _GotoState91Action,
	{_State366, TildeTildeToken}:                         _GotoState97Action,
	{_State366, BitNegToken}:                             _GotoState88Action,
	{_State366, BitAndToken}:                             _GotoState87Action,
	{_State366, ParseErrorToken}:                         _GotoState95Action,
	{_State366, InitializableTypeType}:                   _GotoState104Action,
	{_State366, AtomTypeType}:                            _GotoState99Action,
	{_State366, ParseErrorTypeType}:                      _GotoState105Action,
	{_State366, ReturnableTypeType}:                      _GotoState108Action,
	{_State366, PrefixedTypeType}:                        _GotoState107Action,
	{_State366, PrefixTypeOpType}:                        _GotoState106Action,
	{_State366, ValueTypeType}:                           _GotoState424Action,
	{_State366, TraitOpTypeType}:                         _GotoState110Action,
	{_State366, ImplicitStructDefType}:                   _GotoState103Action,
	{_State366, ExplicitStructDefType}:                   _GotoState52Action,
	{_State366, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State366, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State366, TraitDefType}:                            _GotoState109Action,
	{_State366, FuncTypeType}:                            _GotoState101Action,
	{_State367, RparenToken}:                             _GotoState425Action,
	{_State368, AddToken}:                                _GotoState191Action,
	{_State368, SubToken}:                                _GotoState196Action,
	{_State368, MulToken}:                                _GotoState194Action,
	{_State368, TraitOpType}:                             _GotoState197Action,
	{_State369, DollarLbracketToken}:                     _GotoState166Action,
	{_State369, OptionalGenericParametersType}:           _GotoState426Action,
	{_State370, LbraceToken}:                             _GotoState72Action,
	{_State370, StatementBlockType}:                      _GotoState427Action,
	{_State373, IdentifierToken}:                         _GotoState180Action,
	{_State373, UnsafeToken}:                             _GotoState181Action,
	{_State373, StructToken}:                             _GotoState34Action,
	{_State373, EnumToken}:                               _GotoState90Action,
	{_State373, TraitToken}:                              _GotoState98Action,
	{_State373, FuncToken}:                               _GotoState92Action,
	{_State373, LparenToken}:                             _GotoState94Action,
	{_State373, LbracketToken}:                           _GotoState26Action,
	{_State373, DotToken}:                                _GotoState89Action,
	{_State373, QuestionToken}:                           _GotoState96Action,
	{_State373, ExclaimToken}:                            _GotoState91Action,
	{_State373, TildeTildeToken}:                         _GotoState97Action,
	{_State373, BitNegToken}:                             _GotoState88Action,
	{_State373, BitAndToken}:                             _GotoState87Action,
	{_State373, ParseErrorToken}:                         _GotoState95Action,
	{_State373, UnsafeStatementType}:                     _GotoState187Action,
	{_State373, InitializableTypeType}:                   _GotoState104Action,
	{_State373, AtomTypeType}:                            _GotoState99Action,
	{_State373, ParseErrorTypeType}:                      _GotoState105Action,
	{_State373, ReturnableTypeType}:                      _GotoState108Action,
	{_State373, PrefixedTypeType}:                        _GotoState107Action,
	{_State373, PrefixTypeOpType}:                        _GotoState106Action,
	{_State373, ValueTypeType}:                           _GotoState188Action,
	{_State373, TraitOpTypeType}:                         _GotoState110Action,
	{_State373, FieldDefType}:                            _GotoState279Action,
	{_State373, ImplicitStructDefType}:                   _GotoState103Action,
	{_State373, ExplicitStructDefType}:                   _GotoState52Action,
	{_State373, EnumValueDefType}:                        _GotoState428Action,
	{_State373, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State373, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State373, TraitDefType}:                            _GotoState109Action,
	{_State373, FuncTypeType}:                            _GotoState101Action,
	{_State374, IdentifierToken}:                         _GotoState180Action,
	{_State374, UnsafeToken}:                             _GotoState181Action,
	{_State374, StructToken}:                             _GotoState34Action,
	{_State374, EnumToken}:                               _GotoState90Action,
	{_State374, TraitToken}:                              _GotoState98Action,
	{_State374, FuncToken}:                               _GotoState92Action,
	{_State374, LparenToken}:                             _GotoState94Action,
	{_State374, LbracketToken}:                           _GotoState26Action,
	{_State374, DotToken}:                                _GotoState89Action,
	{_State374, QuestionToken}:                           _GotoState96Action,
	{_State374, ExclaimToken}:                            _GotoState91Action,
	{_State374, TildeTildeToken}:                         _GotoState97Action,
	{_State374, BitNegToken}:                             _GotoState88Action,
	{_State374, BitAndToken}:                             _GotoState87Action,
	{_State374, ParseErrorToken}:                         _GotoState95Action,
	{_State374, UnsafeStatementType}:                     _GotoState187Action,
	{_State374, InitializableTypeType}:                   _GotoState104Action,
	{_State374, AtomTypeType}:                            _GotoState99Action,
	{_State374, ParseErrorTypeType}:                      _GotoState105Action,
	{_State374, ReturnableTypeType}:                      _GotoState108Action,
	{_State374, PrefixedTypeType}:                        _GotoState107Action,
	{_State374, PrefixTypeOpType}:                        _GotoState106Action,
	{_State374, ValueTypeType}:                           _GotoState188Action,
	{_State374, TraitOpTypeType}:                         _GotoState110Action,
	{_State374, FieldDefType}:                            _GotoState279Action,
	{_State374, ImplicitStructDefType}:                   _GotoState103Action,
	{_State374, ExplicitStructDefType}:                   _GotoState52Action,
	{_State374, EnumValueDefType}:                        _GotoState429Action,
	{_State374, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State374, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State374, TraitDefType}:                            _GotoState109Action,
	{_State374, FuncTypeType}:                            _GotoState101Action,
	{_State376, IdentifierToken}:                         _GotoState180Action,
	{_State376, UnsafeToken}:                             _GotoState181Action,
	{_State376, StructToken}:                             _GotoState34Action,
	{_State376, EnumToken}:                               _GotoState90Action,
	{_State376, TraitToken}:                              _GotoState98Action,
	{_State376, FuncToken}:                               _GotoState92Action,
	{_State376, LparenToken}:                             _GotoState94Action,
	{_State376, LbracketToken}:                           _GotoState26Action,
	{_State376, DotToken}:                                _GotoState89Action,
	{_State376, QuestionToken}:                           _GotoState96Action,
	{_State376, ExclaimToken}:                            _GotoState91Action,
	{_State376, TildeTildeToken}:                         _GotoState97Action,
	{_State376, BitNegToken}:                             _GotoState88Action,
	{_State376, BitAndToken}:                             _GotoState87Action,
	{_State376, ParseErrorToken}:                         _GotoState95Action,
	{_State376, UnsafeStatementType}:                     _GotoState187Action,
	{_State376, InitializableTypeType}:                   _GotoState104Action,
	{_State376, AtomTypeType}:                            _GotoState99Action,
	{_State376, ParseErrorTypeType}:                      _GotoState105Action,
	{_State376, ReturnableTypeType}:                      _GotoState108Action,
	{_State376, PrefixedTypeType}:                        _GotoState107Action,
	{_State376, PrefixTypeOpType}:                        _GotoState106Action,
	{_State376, ValueTypeType}:                           _GotoState188Action,
	{_State376, TraitOpTypeType}:                         _GotoState110Action,
	{_State376, FieldDefType}:                            _GotoState279Action,
	{_State376, ImplicitStructDefType}:                   _GotoState103Action,
	{_State376, ExplicitStructDefType}:                   _GotoState52Action,
	{_State376, EnumValueDefType}:                        _GotoState430Action,
	{_State376, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State376, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State376, TraitDefType}:                            _GotoState109Action,
	{_State376, FuncTypeType}:                            _GotoState101Action,
	{_State377, IdentifierToken}:                         _GotoState180Action,
	{_State377, UnsafeToken}:                             _GotoState181Action,
	{_State377, StructToken}:                             _GotoState34Action,
	{_State377, EnumToken}:                               _GotoState90Action,
	{_State377, TraitToken}:                              _GotoState98Action,
	{_State377, FuncToken}:                               _GotoState92Action,
	{_State377, LparenToken}:                             _GotoState94Action,
	{_State377, LbracketToken}:                           _GotoState26Action,
	{_State377, DotToken}:                                _GotoState89Action,
	{_State377, QuestionToken}:                           _GotoState96Action,
	{_State377, ExclaimToken}:                            _GotoState91Action,
	{_State377, TildeTildeToken}:                         _GotoState97Action,
	{_State377, BitNegToken}:                             _GotoState88Action,
	{_State377, BitAndToken}:                             _GotoState87Action,
	{_State377, ParseErrorToken}:                         _GotoState95Action,
	{_State377, UnsafeStatementType}:                     _GotoState187Action,
	{_State377, InitializableTypeType}:                   _GotoState104Action,
	{_State377, AtomTypeType}:                            _GotoState99Action,
	{_State377, ParseErrorTypeType}:                      _GotoState105Action,
	{_State377, ReturnableTypeType}:                      _GotoState108Action,
	{_State377, PrefixedTypeType}:                        _GotoState107Action,
	{_State377, PrefixTypeOpType}:                        _GotoState106Action,
	{_State377, ValueTypeType}:                           _GotoState188Action,
	{_State377, TraitOpTypeType}:                         _GotoState110Action,
	{_State377, FieldDefType}:                            _GotoState279Action,
	{_State377, ImplicitStructDefType}:                   _GotoState103Action,
	{_State377, ExplicitStructDefType}:                   _GotoState52Action,
	{_State377, EnumValueDefType}:                        _GotoState431Action,
	{_State377, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State377, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State377, TraitDefType}:                            _GotoState109Action,
	{_State377, FuncTypeType}:                            _GotoState101Action,
	{_State378, AddToken}:                                _GotoState191Action,
	{_State378, SubToken}:                                _GotoState196Action,
	{_State378, MulToken}:                                _GotoState194Action,
	{_State378, TraitOpType}:                             _GotoState197Action,
	{_State379, IdentifierToken}:                         _GotoState93Action,
	{_State379, StructToken}:                             _GotoState34Action,
	{_State379, EnumToken}:                               _GotoState90Action,
	{_State379, TraitToken}:                              _GotoState98Action,
	{_State379, FuncToken}:                               _GotoState92Action,
	{_State379, LparenToken}:                             _GotoState94Action,
	{_State379, LbracketToken}:                           _GotoState26Action,
	{_State379, DotToken}:                                _GotoState89Action,
	{_State379, QuestionToken}:                           _GotoState96Action,
	{_State379, ExclaimToken}:                            _GotoState91Action,
	{_State379, TildeTildeToken}:                         _GotoState97Action,
	{_State379, BitNegToken}:                             _GotoState88Action,
	{_State379, BitAndToken}:                             _GotoState87Action,
	{_State379, ParseErrorToken}:                         _GotoState95Action,
	{_State379, InitializableTypeType}:                   _GotoState104Action,
	{_State379, AtomTypeType}:                            _GotoState99Action,
	{_State379, ParseErrorTypeType}:                      _GotoState105Action,
	{_State379, ReturnableTypeType}:                      _GotoState108Action,
	{_State379, PrefixedTypeType}:                        _GotoState107Action,
	{_State379, PrefixTypeOpType}:                        _GotoState106Action,
	{_State379, ValueTypeType}:                           _GotoState432Action,
	{_State379, TraitOpTypeType}:                         _GotoState110Action,
	{_State379, ImplicitStructDefType}:                   _GotoState103Action,
	{_State379, ExplicitStructDefType}:                   _GotoState52Action,
	{_State379, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State379, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State379, TraitDefType}:                            _GotoState109Action,
	{_State379, FuncTypeType}:                            _GotoState101Action,
	{_State380, AddToken}:                                _GotoState191Action,
	{_State380, SubToken}:                                _GotoState196Action,
	{_State380, MulToken}:                                _GotoState194Action,
	{_State380, TraitOpType}:                             _GotoState197Action,
	{_State381, IdentifierToken}:                         _GotoState93Action,
	{_State381, StructToken}:                             _GotoState34Action,
	{_State381, EnumToken}:                               _GotoState90Action,
	{_State381, TraitToken}:                              _GotoState98Action,
	{_State381, FuncToken}:                               _GotoState92Action,
	{_State381, LparenToken}:                             _GotoState94Action,
	{_State381, LbracketToken}:                           _GotoState26Action,
	{_State381, DotToken}:                                _GotoState89Action,
	{_State381, QuestionToken}:                           _GotoState96Action,
	{_State381, ExclaimToken}:                            _GotoState91Action,
	{_State381, TildeTildeToken}:                         _GotoState97Action,
	{_State381, BitNegToken}:                             _GotoState88Action,
	{_State381, BitAndToken}:                             _GotoState87Action,
	{_State381, ParseErrorToken}:                         _GotoState95Action,
	{_State381, InitializableTypeType}:                   _GotoState104Action,
	{_State381, AtomTypeType}:                            _GotoState99Action,
	{_State381, ParseErrorTypeType}:                      _GotoState105Action,
	{_State381, ReturnableTypeType}:                      _GotoState371Action,
	{_State381, PrefixedTypeType}:                        _GotoState107Action,
	{_State381, PrefixTypeOpType}:                        _GotoState106Action,
	{_State381, ImplicitStructDefType}:                   _GotoState103Action,
	{_State381, ExplicitStructDefType}:                   _GotoState52Action,
	{_State381, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State381, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State381, TraitDefType}:                            _GotoState109Action,
	{_State381, ReturnTypeType}:                          _GotoState433Action,
	{_State381, FuncTypeType}:                            _GotoState101Action,
	{_State382, IdentifierToken}:                         _GotoState282Action,
	{_State382, StructToken}:                             _GotoState34Action,
	{_State382, EnumToken}:                               _GotoState90Action,
	{_State382, TraitToken}:                              _GotoState98Action,
	{_State382, FuncToken}:                               _GotoState92Action,
	{_State382, LparenToken}:                             _GotoState94Action,
	{_State382, LbracketToken}:                           _GotoState26Action,
	{_State382, DotToken}:                                _GotoState89Action,
	{_State382, QuestionToken}:                           _GotoState96Action,
	{_State382, ExclaimToken}:                            _GotoState91Action,
	{_State382, DotDotDotToken}:                          _GotoState281Action,
	{_State382, TildeTildeToken}:                         _GotoState97Action,
	{_State382, BitNegToken}:                             _GotoState88Action,
	{_State382, BitAndToken}:                             _GotoState87Action,
	{_State382, ParseErrorToken}:                         _GotoState95Action,
	{_State382, InitializableTypeType}:                   _GotoState104Action,
	{_State382, AtomTypeType}:                            _GotoState99Action,
	{_State382, ParseErrorTypeType}:                      _GotoState105Action,
	{_State382, ReturnableTypeType}:                      _GotoState108Action,
	{_State382, PrefixedTypeType}:                        _GotoState107Action,
	{_State382, PrefixTypeOpType}:                        _GotoState106Action,
	{_State382, ValueTypeType}:                           _GotoState286Action,
	{_State382, TraitOpTypeType}:                         _GotoState110Action,
	{_State382, ImplicitStructDefType}:                   _GotoState103Action,
	{_State382, ExplicitStructDefType}:                   _GotoState52Action,
	{_State382, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State382, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State382, TraitDefType}:                            _GotoState109Action,
	{_State382, ParameterDeclType}:                       _GotoState434Action,
	{_State382, FuncTypeType}:                            _GotoState101Action,
	{_State384, GreaterToken}:                            _GotoState435Action,
	{_State389, LparenToken}:                             _GotoState436Action,
	{_State391, IdentifierToken}:                         _GotoState180Action,
	{_State391, UnsafeToken}:                             _GotoState181Action,
	{_State391, StructToken}:                             _GotoState34Action,
	{_State391, EnumToken}:                               _GotoState90Action,
	{_State391, TraitToken}:                              _GotoState98Action,
	{_State391, FuncToken}:                               _GotoState297Action,
	{_State391, LparenToken}:                             _GotoState94Action,
	{_State391, LbracketToken}:                           _GotoState26Action,
	{_State391, DotToken}:                                _GotoState89Action,
	{_State391, QuestionToken}:                           _GotoState96Action,
	{_State391, ExclaimToken}:                            _GotoState91Action,
	{_State391, TildeTildeToken}:                         _GotoState97Action,
	{_State391, BitNegToken}:                             _GotoState88Action,
	{_State391, BitAndToken}:                             _GotoState87Action,
	{_State391, ParseErrorToken}:                         _GotoState95Action,
	{_State391, UnsafeStatementType}:                     _GotoState187Action,
	{_State391, InitializableTypeType}:                   _GotoState104Action,
	{_State391, AtomTypeType}:                            _GotoState99Action,
	{_State391, ParseErrorTypeType}:                      _GotoState105Action,
	{_State391, ReturnableTypeType}:                      _GotoState108Action,
	{_State391, PrefixedTypeType}:                        _GotoState107Action,
	{_State391, PrefixTypeOpType}:                        _GotoState106Action,
	{_State391, ValueTypeType}:                           _GotoState188Action,
	{_State391, TraitOpTypeType}:                         _GotoState110Action,
	{_State391, FieldDefType}:                            _GotoState298Action,
	{_State391, ImplicitStructDefType}:                   _GotoState103Action,
	{_State391, ExplicitStructDefType}:                   _GotoState52Action,
	{_State391, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State391, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State391, TraitPropertyType}:                       _GotoState437Action,
	{_State391, TraitDefType}:                            _GotoState109Action,
	{_State391, FuncTypeType}:                            _GotoState101Action,
	{_State391, MethodSignatureType}:                     _GotoState299Action,
	{_State392, IdentifierToken}:                         _GotoState180Action,
	{_State392, UnsafeToken}:                             _GotoState181Action,
	{_State392, StructToken}:                             _GotoState34Action,
	{_State392, EnumToken}:                               _GotoState90Action,
	{_State392, TraitToken}:                              _GotoState98Action,
	{_State392, FuncToken}:                               _GotoState297Action,
	{_State392, LparenToken}:                             _GotoState94Action,
	{_State392, LbracketToken}:                           _GotoState26Action,
	{_State392, DotToken}:                                _GotoState89Action,
	{_State392, QuestionToken}:                           _GotoState96Action,
	{_State392, ExclaimToken}:                            _GotoState91Action,
	{_State392, TildeTildeToken}:                         _GotoState97Action,
	{_State392, BitNegToken}:                             _GotoState88Action,
	{_State392, BitAndToken}:                             _GotoState87Action,
	{_State392, ParseErrorToken}:                         _GotoState95Action,
	{_State392, UnsafeStatementType}:                     _GotoState187Action,
	{_State392, InitializableTypeType}:                   _GotoState104Action,
	{_State392, AtomTypeType}:                            _GotoState99Action,
	{_State392, ParseErrorTypeType}:                      _GotoState105Action,
	{_State392, ReturnableTypeType}:                      _GotoState108Action,
	{_State392, PrefixedTypeType}:                        _GotoState107Action,
	{_State392, PrefixTypeOpType}:                        _GotoState106Action,
	{_State392, ValueTypeType}:                           _GotoState188Action,
	{_State392, TraitOpTypeType}:                         _GotoState110Action,
	{_State392, FieldDefType}:                            _GotoState298Action,
	{_State392, ImplicitStructDefType}:                   _GotoState103Action,
	{_State392, ExplicitStructDefType}:                   _GotoState52Action,
	{_State392, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State392, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State392, TraitPropertyType}:                       _GotoState438Action,
	{_State392, TraitDefType}:                            _GotoState109Action,
	{_State392, FuncTypeType}:                            _GotoState101Action,
	{_State392, MethodSignatureType}:                     _GotoState299Action,
	{_State397, AddToken}:                                _GotoState191Action,
	{_State397, SubToken}:                                _GotoState196Action,
	{_State397, MulToken}:                                _GotoState194Action,
	{_State397, TraitOpType}:                             _GotoState197Action,
	{_State401, DoToken}:                                 _GotoState439Action,
	{_State402, SemicolonToken}:                          _GotoState440Action,
	{_State405, LparenToken}:                             _GotoState28Action,
	{_State405, ImplicitStructExprType}:                  _GotoState441Action,
	{_State406, IdentifierToken}:                         _GotoState442Action,
	{_State407, IntegerLiteralToken}:                     _GotoState24Action,
	{_State407, FloatLiteralToken}:                       _GotoState20Action,
	{_State407, RuneLiteralToken}:                        _GotoState32Action,
	{_State407, StringLiteralToken}:                      _GotoState33Action,
	{_State407, IdentifierToken}:                         _GotoState23Action,
	{_State407, TrueToken}:                               _GotoState36Action,
	{_State407, FalseToken}:                              _GotoState19Action,
	{_State407, StructToken}:                             _GotoState34Action,
	{_State407, FuncToken}:                               _GotoState21Action,
	{_State407, VarToken}:                                _GotoState37Action,
	{_State407, LetToken}:                                _GotoState27Action,
	{_State407, NotToken}:                                _GotoState30Action,
	{_State407, LabelDeclToken}:                          _GotoState25Action,
	{_State407, LparenToken}:                             _GotoState28Action,
	{_State407, LbracketToken}:                           _GotoState26Action,
	{_State407, SubToken}:                                _GotoState35Action,
	{_State407, MulToken}:                                _GotoState29Action,
	{_State407, BitNegToken}:                             _GotoState18Action,
	{_State407, BitAndToken}:                             _GotoState17Action,
	{_State407, GreaterToken}:                            _GotoState22Action,
	{_State407, ParseErrorToken}:                         _GotoState31Action,
	{_State407, VarDeclPatternType}:                      _GotoState69Action,
	{_State407, VarOrLetType}:                            _GotoState70Action,
	{_State407, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State407, SequenceExprType}:                        _GotoState443Action,
	{_State407, CallExprType}:                            _GotoState50Action,
	{_State407, AtomExprType}:                            _GotoState43Action,
	{_State407, ParseErrorExprType}:                      _GotoState62Action,
	{_State407, LiteralExprType}:                         _GotoState58Action,
	{_State407, IdentifierExprType}:                      _GotoState53Action,
	{_State407, BlockExprType}:                           _GotoState49Action,
	{_State407, InitializeExprType}:                      _GotoState57Action,
	{_State407, ImplicitStructExprType}:                  _GotoState54Action,
	{_State407, AccessibleExprType}:                      _GotoState39Action,
	{_State407, AccessExprType}:                          _GotoState38Action,
	{_State407, IndexExprType}:                           _GotoState55Action,
	{_State407, PostfixableExprType}:                     _GotoState64Action,
	{_State407, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State407, PrefixableExprType}:                      _GotoState67Action,
	{_State407, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State407, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State407, MulExprType}:                             _GotoState59Action,
	{_State407, BinaryMulExprType}:                       _GotoState47Action,
	{_State407, AddExprType}:                             _GotoState40Action,
	{_State407, BinaryAddExprType}:                       _GotoState44Action,
	{_State407, CmpExprType}:                             _GotoState51Action,
	{_State407, BinaryCmpExprType}:                       _GotoState46Action,
	{_State407, AndExprType}:                             _GotoState41Action,
	{_State407, BinaryAndExprType}:                       _GotoState45Action,
	{_State407, OrExprType}:                              _GotoState61Action,
	{_State407, BinaryOrExprType}:                        _GotoState48Action,
	{_State407, InitializableTypeType}:                   _GotoState56Action,
	{_State407, ExplicitStructDefType}:                   _GotoState52Action,
	{_State407, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State408, IntegerLiteralToken}:                     _GotoState24Action,
	{_State408, FloatLiteralToken}:                       _GotoState20Action,
	{_State408, RuneLiteralToken}:                        _GotoState32Action,
	{_State408, StringLiteralToken}:                      _GotoState33Action,
	{_State408, IdentifierToken}:                         _GotoState23Action,
	{_State408, TrueToken}:                               _GotoState36Action,
	{_State408, FalseToken}:                              _GotoState19Action,
	{_State408, StructToken}:                             _GotoState34Action,
	{_State408, FuncToken}:                               _GotoState21Action,
	{_State408, VarToken}:                                _GotoState325Action,
	{_State408, LetToken}:                                _GotoState27Action,
	{_State408, NotToken}:                                _GotoState30Action,
	{_State408, LabelDeclToken}:                          _GotoState25Action,
	{_State408, LparenToken}:                             _GotoState28Action,
	{_State408, LbracketToken}:                           _GotoState26Action,
	{_State408, DotToken}:                                _GotoState324Action,
	{_State408, SubToken}:                                _GotoState35Action,
	{_State408, MulToken}:                                _GotoState29Action,
	{_State408, BitNegToken}:                             _GotoState18Action,
	{_State408, BitAndToken}:                             _GotoState17Action,
	{_State408, GreaterToken}:                            _GotoState22Action,
	{_State408, ParseErrorToken}:                         _GotoState31Action,
	{_State408, VarDeclPatternType}:                      _GotoState69Action,
	{_State408, VarOrLetType}:                            _GotoState70Action,
	{_State408, AssignPatternType}:                       _GotoState326Action,
	{_State408, CasePatternType}:                         _GotoState444Action,
	{_State408, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State408, SequenceExprType}:                        _GotoState329Action,
	{_State408, CallExprType}:                            _GotoState50Action,
	{_State408, AtomExprType}:                            _GotoState43Action,
	{_State408, ParseErrorExprType}:                      _GotoState62Action,
	{_State408, LiteralExprType}:                         _GotoState58Action,
	{_State408, IdentifierExprType}:                      _GotoState53Action,
	{_State408, BlockExprType}:                           _GotoState49Action,
	{_State408, InitializeExprType}:                      _GotoState57Action,
	{_State408, ImplicitStructExprType}:                  _GotoState54Action,
	{_State408, AccessibleExprType}:                      _GotoState39Action,
	{_State408, AccessExprType}:                          _GotoState38Action,
	{_State408, IndexExprType}:                           _GotoState55Action,
	{_State408, PostfixableExprType}:                     _GotoState64Action,
	{_State408, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State408, PrefixableExprType}:                      _GotoState67Action,
	{_State408, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State408, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State408, MulExprType}:                             _GotoState59Action,
	{_State408, BinaryMulExprType}:                       _GotoState47Action,
	{_State408, AddExprType}:                             _GotoState40Action,
	{_State408, BinaryAddExprType}:                       _GotoState44Action,
	{_State408, CmpExprType}:                             _GotoState51Action,
	{_State408, BinaryCmpExprType}:                       _GotoState46Action,
	{_State408, AndExprType}:                             _GotoState41Action,
	{_State408, BinaryAndExprType}:                       _GotoState45Action,
	{_State408, OrExprType}:                              _GotoState61Action,
	{_State408, BinaryOrExprType}:                        _GotoState48Action,
	{_State408, InitializableTypeType}:                   _GotoState56Action,
	{_State408, ExplicitStructDefType}:                   _GotoState52Action,
	{_State408, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State409, IfToken}:                                 _GotoState149Action,
	{_State409, LbraceToken}:                             _GotoState72Action,
	{_State409, StatementBlockType}:                      _GotoState446Action,
	{_State409, IfExprType}:                              _GotoState445Action,
	{_State412, IntegerLiteralToken}:                     _GotoState24Action,
	{_State412, FloatLiteralToken}:                       _GotoState20Action,
	{_State412, RuneLiteralToken}:                        _GotoState32Action,
	{_State412, StringLiteralToken}:                      _GotoState33Action,
	{_State412, IdentifierToken}:                         _GotoState23Action,
	{_State412, TrueToken}:                               _GotoState36Action,
	{_State412, FalseToken}:                              _GotoState19Action,
	{_State412, ReturnToken}:                             _GotoState242Action,
	{_State412, BreakToken}:                              _GotoState234Action,
	{_State412, ContinueToken}:                           _GotoState236Action,
	{_State412, FallthroughToken}:                        _GotoState239Action,
	{_State412, UnsafeToken}:                             _GotoState181Action,
	{_State412, StructToken}:                             _GotoState34Action,
	{_State412, FuncToken}:                               _GotoState21Action,
	{_State412, AsyncToken}:                              _GotoState233Action,
	{_State412, DeferToken}:                              _GotoState238Action,
	{_State412, VarToken}:                                _GotoState37Action,
	{_State412, LetToken}:                                _GotoState27Action,
	{_State412, NotToken}:                                _GotoState30Action,
	{_State412, LabelDeclToken}:                          _GotoState25Action,
	{_State412, LparenToken}:                             _GotoState28Action,
	{_State412, LbracketToken}:                           _GotoState26Action,
	{_State412, SubToken}:                                _GotoState35Action,
	{_State412, MulToken}:                                _GotoState29Action,
	{_State412, BitNegToken}:                             _GotoState18Action,
	{_State412, BitAndToken}:                             _GotoState17Action,
	{_State412, GreaterToken}:                            _GotoState22Action,
	{_State412, ParseErrorToken}:                         _GotoState31Action,
	{_State412, SimpleStatementBodyType}:                 _GotoState414Action,
	{_State412, OptionalSimpleStatementBodyType}:         _GotoState447Action,
	{_State412, ExpressionOrImproperStructStatementType}: _GotoState250Action,
	{_State412, CallbackOpType}:                          _GotoState247Action,
	{_State412, CallbackOpStatementType}:                 _GotoState248Action,
	{_State412, UnsafeStatementType}:                     _GotoState261Action,
	{_State412, JumpStatementType}:                       _GotoState254Action,
	{_State412, JumpTypeType}:                            _GotoState255Action,
	{_State412, ExpressionsType}:                         _GotoState251Action,
	{_State412, FallthroughStatementType}:                _GotoState252Action,
	{_State412, AssignStatementType}:                     _GotoState245Action,
	{_State412, UnaryOpAssignStatementType}:              _GotoState260Action,
	{_State412, BinaryOpAssignStatementType}:             _GotoState246Action,
	{_State412, VarDeclPatternType}:                      _GotoState69Action,
	{_State412, VarOrLetType}:                            _GotoState70Action,
	{_State412, AssignPatternType}:                       _GotoState244Action,
	{_State412, ExpressionType}:                          _GotoState249Action,
	{_State412, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State412, SequenceExprType}:                        _GotoState256Action,
	{_State412, CallExprType}:                            _GotoState50Action,
	{_State412, AtomExprType}:                            _GotoState43Action,
	{_State412, ParseErrorExprType}:                      _GotoState62Action,
	{_State412, LiteralExprType}:                         _GotoState58Action,
	{_State412, IdentifierExprType}:                      _GotoState53Action,
	{_State412, BlockExprType}:                           _GotoState49Action,
	{_State412, InitializeExprType}:                      _GotoState57Action,
	{_State412, ImplicitStructExprType}:                  _GotoState54Action,
	{_State412, AccessibleExprType}:                      _GotoState243Action,
	{_State412, AccessExprType}:                          _GotoState38Action,
	{_State412, IndexExprType}:                           _GotoState55Action,
	{_State412, PostfixableExprType}:                     _GotoState64Action,
	{_State412, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State412, PrefixableExprType}:                      _GotoState67Action,
	{_State412, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State412, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State412, MulExprType}:                             _GotoState59Action,
	{_State412, BinaryMulExprType}:                       _GotoState47Action,
	{_State412, AddExprType}:                             _GotoState40Action,
	{_State412, BinaryAddExprType}:                       _GotoState44Action,
	{_State412, CmpExprType}:                             _GotoState51Action,
	{_State412, BinaryCmpExprType}:                       _GotoState46Action,
	{_State412, AndExprType}:                             _GotoState41Action,
	{_State412, BinaryAndExprType}:                       _GotoState45Action,
	{_State412, OrExprType}:                              _GotoState61Action,
	{_State412, BinaryOrExprType}:                        _GotoState48Action,
	{_State412, InitializableTypeType}:                   _GotoState56Action,
	{_State412, ExplicitStructDefType}:                   _GotoState52Action,
	{_State412, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State415, NewlinesToken}:                           _GotoState449Action,
	{_State415, CommaToken}:                              _GotoState448Action,
	{_State417, StringLiteralToken}:                      _GotoState338Action,
	{_State417, RparenToken}:                             _GotoState450Action,
	{_State417, ImportClauseType}:                        _GotoState415Action,
	{_State417, ImportClauseTerminalType}:                _GotoState451Action,
	{_State418, IdentifierToken}:                         _GotoState452Action,
	{_State422, CommaToken}:                              _GotoState358Action,
	{_State424, AddToken}:                                _GotoState191Action,
	{_State424, SubToken}:                                _GotoState196Action,
	{_State424, MulToken}:                                _GotoState194Action,
	{_State424, TraitOpType}:                             _GotoState197Action,
	{_State425, IdentifierToken}:                         _GotoState93Action,
	{_State425, StructToken}:                             _GotoState34Action,
	{_State425, EnumToken}:                               _GotoState90Action,
	{_State425, TraitToken}:                              _GotoState98Action,
	{_State425, FuncToken}:                               _GotoState92Action,
	{_State425, LparenToken}:                             _GotoState94Action,
	{_State425, LbracketToken}:                           _GotoState26Action,
	{_State425, DotToken}:                                _GotoState89Action,
	{_State425, QuestionToken}:                           _GotoState96Action,
	{_State425, ExclaimToken}:                            _GotoState91Action,
	{_State425, TildeTildeToken}:                         _GotoState97Action,
	{_State425, BitNegToken}:                             _GotoState88Action,
	{_State425, BitAndToken}:                             _GotoState87Action,
	{_State425, ParseErrorToken}:                         _GotoState95Action,
	{_State425, InitializableTypeType}:                   _GotoState104Action,
	{_State425, AtomTypeType}:                            _GotoState99Action,
	{_State425, ParseErrorTypeType}:                      _GotoState105Action,
	{_State425, ReturnableTypeType}:                      _GotoState371Action,
	{_State425, PrefixedTypeType}:                        _GotoState107Action,
	{_State425, PrefixTypeOpType}:                        _GotoState106Action,
	{_State425, ImplicitStructDefType}:                   _GotoState103Action,
	{_State425, ExplicitStructDefType}:                   _GotoState52Action,
	{_State425, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State425, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State425, TraitDefType}:                            _GotoState109Action,
	{_State425, ReturnTypeType}:                          _GotoState453Action,
	{_State425, FuncTypeType}:                            _GotoState101Action,
	{_State426, LparenToken}:                             _GotoState454Action,
	{_State432, AddToken}:                                _GotoState191Action,
	{_State432, SubToken}:                                _GotoState196Action,
	{_State432, MulToken}:                                _GotoState194Action,
	{_State432, TraitOpType}:                             _GotoState197Action,
	{_State435, StringLiteralToken}:                      _GotoState455Action,
	{_State436, IdentifierToken}:                         _GotoState282Action,
	{_State436, StructToken}:                             _GotoState34Action,
	{_State436, EnumToken}:                               _GotoState90Action,
	{_State436, TraitToken}:                              _GotoState98Action,
	{_State436, FuncToken}:                               _GotoState92Action,
	{_State436, LparenToken}:                             _GotoState94Action,
	{_State436, LbracketToken}:                           _GotoState26Action,
	{_State436, DotToken}:                                _GotoState89Action,
	{_State436, QuestionToken}:                           _GotoState96Action,
	{_State436, ExclaimToken}:                            _GotoState91Action,
	{_State436, DotDotDotToken}:                          _GotoState281Action,
	{_State436, TildeTildeToken}:                         _GotoState97Action,
	{_State436, BitNegToken}:                             _GotoState88Action,
	{_State436, BitAndToken}:                             _GotoState87Action,
	{_State436, ParseErrorToken}:                         _GotoState95Action,
	{_State436, InitializableTypeType}:                   _GotoState104Action,
	{_State436, AtomTypeType}:                            _GotoState99Action,
	{_State436, ParseErrorTypeType}:                      _GotoState105Action,
	{_State436, ReturnableTypeType}:                      _GotoState108Action,
	{_State436, PrefixedTypeType}:                        _GotoState107Action,
	{_State436, PrefixTypeOpType}:                        _GotoState106Action,
	{_State436, ValueTypeType}:                           _GotoState286Action,
	{_State436, TraitOpTypeType}:                         _GotoState110Action,
	{_State436, ImplicitStructDefType}:                   _GotoState103Action,
	{_State436, ExplicitStructDefType}:                   _GotoState52Action,
	{_State436, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State436, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State436, TraitDefType}:                            _GotoState109Action,
	{_State436, ParameterDeclType}:                       _GotoState284Action,
	{_State436, ParameterDeclsType}:                      _GotoState285Action,
	{_State436, OptionalParameterDeclsType}:              _GotoState456Action,
	{_State436, FuncTypeType}:                            _GotoState101Action,
	{_State439, LbraceToken}:                             _GotoState72Action,
	{_State439, StatementBlockType}:                      _GotoState457Action,
	{_State440, IntegerLiteralToken}:                     _GotoState24Action,
	{_State440, FloatLiteralToken}:                       _GotoState20Action,
	{_State440, RuneLiteralToken}:                        _GotoState32Action,
	{_State440, StringLiteralToken}:                      _GotoState33Action,
	{_State440, IdentifierToken}:                         _GotoState23Action,
	{_State440, TrueToken}:                               _GotoState36Action,
	{_State440, FalseToken}:                              _GotoState19Action,
	{_State440, StructToken}:                             _GotoState34Action,
	{_State440, FuncToken}:                               _GotoState21Action,
	{_State440, VarToken}:                                _GotoState37Action,
	{_State440, LetToken}:                                _GotoState27Action,
	{_State440, NotToken}:                                _GotoState30Action,
	{_State440, LabelDeclToken}:                          _GotoState25Action,
	{_State440, LparenToken}:                             _GotoState28Action,
	{_State440, LbracketToken}:                           _GotoState26Action,
	{_State440, SubToken}:                                _GotoState35Action,
	{_State440, MulToken}:                                _GotoState29Action,
	{_State440, BitNegToken}:                             _GotoState18Action,
	{_State440, BitAndToken}:                             _GotoState17Action,
	{_State440, GreaterToken}:                            _GotoState22Action,
	{_State440, ParseErrorToken}:                         _GotoState31Action,
	{_State440, VarDeclPatternType}:                      _GotoState69Action,
	{_State440, VarOrLetType}:                            _GotoState70Action,
	{_State440, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State440, SequenceExprType}:                        _GotoState403Action,
	{_State440, OptionalSequenceExprType}:                _GotoState458Action,
	{_State440, CallExprType}:                            _GotoState50Action,
	{_State440, AtomExprType}:                            _GotoState43Action,
	{_State440, ParseErrorExprType}:                      _GotoState62Action,
	{_State440, LiteralExprType}:                         _GotoState58Action,
	{_State440, IdentifierExprType}:                      _GotoState53Action,
	{_State440, BlockExprType}:                           _GotoState49Action,
	{_State440, InitializeExprType}:                      _GotoState57Action,
	{_State440, ImplicitStructExprType}:                  _GotoState54Action,
	{_State440, AccessibleExprType}:                      _GotoState39Action,
	{_State440, AccessExprType}:                          _GotoState38Action,
	{_State440, IndexExprType}:                           _GotoState55Action,
	{_State440, PostfixableExprType}:                     _GotoState64Action,
	{_State440, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State440, PrefixableExprType}:                      _GotoState67Action,
	{_State440, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State440, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State440, MulExprType}:                             _GotoState59Action,
	{_State440, BinaryMulExprType}:                       _GotoState47Action,
	{_State440, AddExprType}:                             _GotoState40Action,
	{_State440, BinaryAddExprType}:                       _GotoState44Action,
	{_State440, CmpExprType}:                             _GotoState51Action,
	{_State440, BinaryCmpExprType}:                       _GotoState46Action,
	{_State440, AndExprType}:                             _GotoState41Action,
	{_State440, BinaryAndExprType}:                       _GotoState45Action,
	{_State440, OrExprType}:                              _GotoState61Action,
	{_State440, BinaryOrExprType}:                        _GotoState48Action,
	{_State440, InitializableTypeType}:                   _GotoState56Action,
	{_State440, ExplicitStructDefType}:                   _GotoState52Action,
	{_State440, AnonymousFuncExprType}:                   _GotoState42Action,
	{_State442, LparenToken}:                             _GotoState158Action,
	{_State442, TuplePatternType}:                        _GotoState459Action,
	{_State453, LbraceToken}:                             _GotoState72Action,
	{_State453, StatementBlockType}:                      _GotoState460Action,
	{_State454, IdentifierToken}:                         _GotoState170Action,
	{_State454, ParameterDefType}:                        _GotoState173Action,
	{_State454, ParameterDefsType}:                       _GotoState174Action,
	{_State454, OptionalParameterDefsType}:               _GotoState461Action,
	{_State456, RparenToken}:                             _GotoState462Action,
	{_State458, DoToken}:                                 _GotoState463Action,
	{_State461, RparenToken}:                             _GotoState464Action,
	{_State462, IdentifierToken}:                         _GotoState93Action,
	{_State462, StructToken}:                             _GotoState34Action,
	{_State462, EnumToken}:                               _GotoState90Action,
	{_State462, TraitToken}:                              _GotoState98Action,
	{_State462, FuncToken}:                               _GotoState92Action,
	{_State462, LparenToken}:                             _GotoState94Action,
	{_State462, LbracketToken}:                           _GotoState26Action,
	{_State462, DotToken}:                                _GotoState89Action,
	{_State462, QuestionToken}:                           _GotoState96Action,
	{_State462, ExclaimToken}:                            _GotoState91Action,
	{_State462, TildeTildeToken}:                         _GotoState97Action,
	{_State462, BitNegToken}:                             _GotoState88Action,
	{_State462, BitAndToken}:                             _GotoState87Action,
	{_State462, ParseErrorToken}:                         _GotoState95Action,
	{_State462, InitializableTypeType}:                   _GotoState104Action,
	{_State462, AtomTypeType}:                            _GotoState99Action,
	{_State462, ParseErrorTypeType}:                      _GotoState105Action,
	{_State462, ReturnableTypeType}:                      _GotoState371Action,
	{_State462, PrefixedTypeType}:                        _GotoState107Action,
	{_State462, PrefixTypeOpType}:                        _GotoState106Action,
	{_State462, ImplicitStructDefType}:                   _GotoState103Action,
	{_State462, ExplicitStructDefType}:                   _GotoState52Action,
	{_State462, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State462, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State462, TraitDefType}:                            _GotoState109Action,
	{_State462, ReturnTypeType}:                          _GotoState465Action,
	{_State462, FuncTypeType}:                            _GotoState101Action,
	{_State463, LbraceToken}:                             _GotoState72Action,
	{_State463, StatementBlockType}:                      _GotoState466Action,
	{_State464, IdentifierToken}:                         _GotoState93Action,
	{_State464, StructToken}:                             _GotoState34Action,
	{_State464, EnumToken}:                               _GotoState90Action,
	{_State464, TraitToken}:                              _GotoState98Action,
	{_State464, FuncToken}:                               _GotoState92Action,
	{_State464, LparenToken}:                             _GotoState94Action,
	{_State464, LbracketToken}:                           _GotoState26Action,
	{_State464, DotToken}:                                _GotoState89Action,
	{_State464, QuestionToken}:                           _GotoState96Action,
	{_State464, ExclaimToken}:                            _GotoState91Action,
	{_State464, TildeTildeToken}:                         _GotoState97Action,
	{_State464, BitNegToken}:                             _GotoState88Action,
	{_State464, BitAndToken}:                             _GotoState87Action,
	{_State464, ParseErrorToken}:                         _GotoState95Action,
	{_State464, InitializableTypeType}:                   _GotoState104Action,
	{_State464, AtomTypeType}:                            _GotoState99Action,
	{_State464, ParseErrorTypeType}:                      _GotoState105Action,
	{_State464, ReturnableTypeType}:                      _GotoState371Action,
	{_State464, PrefixedTypeType}:                        _GotoState107Action,
	{_State464, PrefixTypeOpType}:                        _GotoState106Action,
	{_State464, ImplicitStructDefType}:                   _GotoState103Action,
	{_State464, ExplicitStructDefType}:                   _GotoState52Action,
	{_State464, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State464, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State464, TraitDefType}:                            _GotoState109Action,
	{_State464, ReturnTypeType}:                          _GotoState467Action,
	{_State464, FuncTypeType}:                            _GotoState101Action,
	{_State467, LbraceToken}:                             _GotoState72Action,
	{_State467, StatementBlockType}:                      _GotoState468Action,
	{_State1, _EndMarker}:                                _ReduceNilToOptionalDefinitionsAction,
	{_State1, _WildcardMarker}:                           _ReduceNilToOptionalNewlinesAction,
	{_State5, _WildcardMarker}:                           _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State11, _EndMarker}:                               _ReduceNewlinesToOptionalDefinitionsAction,
	{_State11, _WildcardMarker}:                          _ReduceNewlinesToOptionalNewlinesAction,
	{_State12, _EndMarker}:                               _ReduceToSourceAction,
	{_State14, _WildcardMarker}:                          _ReduceNoSpecToPackageDefAction,
	{_State17, _WildcardMarker}:                          _ReduceBitAndToPrefixUnaryOpAction,
	{_State18, _WildcardMarker}:                          _ReduceBitNegToPrefixUnaryOpAction,
	{_State19, _WildcardMarker}:                          _ReduceFalseToLiteralExprAction,
	{_State20, _WildcardMarker}:                          _ReduceFloatLiteralToLiteralExprAction,
	{_State22, LbraceToken}:                              _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State23, _WildcardMarker}:                          _ReduceToIdentifierExprAction,
	{_State24, _WildcardMarker}:                          _ReduceIntegerLiteralToLiteralExprAction,
	{_State25, _WildcardMarker}:                          _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State27, _WildcardMarker}:                          _ReduceLetToVarOrLetAction,
	{_State28, _WildcardMarker}:                          _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State28, RparenToken}:                              _ReduceNilToOptionalArgumentsAction,
	{_State28, ColonToken}:                               _ReduceNilToOptionalExpressionAction,
	{_State29, _WildcardMarker}:                          _ReduceMulToPrefixUnaryOpAction,
	{_State30, _WildcardMarker}:                          _ReduceNotToPrefixUnaryOpAction,
	{_State31, _WildcardMarker}:                          _ReduceToParseErrorExprAction,
	{_State32, _WildcardMarker}:                          _ReduceRuneLiteralToLiteralExprAction,
	{_State33, _WildcardMarker}:                          _ReduceStringLiteralToLiteralExprAction,
	{_State35, _WildcardMarker}:                          _ReduceSubToPrefixUnaryOpAction,
	{_State36, _WildcardMarker}:                          _ReduceTrueToLiteralExprAction,
	{_State37, _WildcardMarker}:                          _ReduceVarToVarOrLetAction,
	{_State38, _WildcardMarker}:                          _ReduceAccessExprToAccessibleExprAction,
	{_State39, _WildcardMarker}:                          _ReduceAccessibleExprToPostfixableExprAction,
	{_State39, LparenToken}:                              _ReduceNilToOptionalGenericBindingAction,
	{_State40, _WildcardMarker}:                          _ReduceAddExprToCmpExprAction,
	{_State41, _WildcardMarker}:                          _ReduceAndExprToOrExprAction,
	{_State42, _WildcardMarker}:                          _ReduceAnonymousFuncExprToAtomExprAction,
	{_State43, _WildcardMarker}:                          _ReduceAtomExprToAccessibleExprAction,
	{_State44, _WildcardMarker}:                          _ReduceBinaryAddExprToAddExprAction,
	{_State45, _WildcardMarker}:                          _ReduceBinaryAndExprToAndExprAction,
	{_State46, _WildcardMarker}:                          _ReduceBinaryCmpExprToCmpExprAction,
	{_State47, _WildcardMarker}:                          _ReduceBinaryMulExprToMulExprAction,
	{_State48, _WildcardMarker}:                          _ReduceBinaryOrExprToOrExprAction,
	{_State49, _WildcardMarker}:                          _ReduceBlockExprToAtomExprAction,
	{_State50, _WildcardMarker}:                          _ReduceCallExprToAccessibleExprAction,
	{_State51, _WildcardMarker}:                          _ReduceCmpExprToAndExprAction,
	{_State52, _WildcardMarker}:                          _ReduceExplicitStructDefToInitializableTypeAction,
	{_State53, _WildcardMarker}:                          _ReduceIdentifierExprToAtomExprAction,
	{_State54, _WildcardMarker}:                          _ReduceImplicitStructExprToAtomExprAction,
	{_State55, _WildcardMarker}:                          _ReduceIndexExprToAccessibleExprAction,
	{_State57, _WildcardMarker}:                          _ReduceInitializeExprToAtomExprAction,
	{_State58, _WildcardMarker}:                          _ReduceLiteralExprToAtomExprAction,
	{_State59, _WildcardMarker}:                          _ReduceMulExprToAddExprAction,
	{_State61, _WildcardMarker}:                          _ReduceOrExprToSequenceExprAction,
	{_State62, _WildcardMarker}:                          _ReduceParseErrorExprToAtomExprAction,
	{_State63, _WildcardMarker}:                          _ReducePostfixUnaryExprToPostfixableExprAction,
	{_State64, _WildcardMarker}:                          _ReducePostfixableExprToPrefixableExprAction,
	{_State65, _WildcardMarker}:                          _ReducePrefixUnaryExprToPrefixableExprAction,
	{_State66, LbraceToken}:                              _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State67, _WildcardMarker}:                          _ReducePrefixableExprToMulExprAction,
	{_State68, _EndMarker}:                               _ReduceSequenceExprToExpressionAction,
	{_State69, _EndMarker}:                               _ReduceVarDeclPatternToSequenceExprAction,
	{_State71, _WildcardMarker}:                          _ReduceCommentGroupsToDefinitionAction,
	{_State72, _WildcardMarker}:                          _ReduceEmptyListToStatementsAction,
	{_State73, _WildcardMarker}:                          _ReduceDefinitionToDefinitionsAction,
	{_State74, _EndMarker}:                               _ReduceNilToOptionalNewlinesAction,
	{_State75, _WildcardMarker}:                          _ReduceNamedFuncDefToDefinitionAction,
	{_State76, _WildcardMarker}:                          _ReducePackageDefToDefinitionAction,
	{_State77, _WildcardMarker}:                          _ReduceStatementBlockToDefinitionAction,
	{_State78, _WildcardMarker}:                          _ReduceTypeDefToDefinitionAction,
	{_State79, _WildcardMarker}:                          _ReduceGlobalVarDefToDefinitionAction,
	{_State80, _EndMarker}:                               _ReduceWithSpecToPackageDefAction,
	{_State81, _WildcardMarker}:                          _ReduceNilToOptionalGenericParametersAction,
	{_State82, LparenToken}:                              _ReduceNilToOptionalGenericParametersAction,
	{_State84, RparenToken}:                              _ReduceNilToOptionalParameterDefsAction,
	{_State86, _EndMarker}:                               _ReduceAssignVarPatternToSequenceExprAction,
	{_State87, _WildcardMarker}:                          _ReduceBitAndToPrefixTypeOpAction,
	{_State88, _WildcardMarker}:                          _ReduceBitNegToPrefixTypeOpAction,
	{_State89, _WildcardMarker}:                          _ReduceNilToOptionalGenericBindingAction,
	{_State91, _WildcardMarker}:                          _ReduceExclaimToPrefixTypeOpAction,
	{_State93, _WildcardMarker}:                          _ReduceNilToOptionalGenericBindingAction,
	{_State94, RparenToken}:                              _ReduceNilToOptionalImplicitFieldDefsAction,
	{_State95, _WildcardMarker}:                          _ReduceToParseErrorTypeAction,
	{_State96, _WildcardMarker}:                          _ReduceQuestionToPrefixTypeOpAction,
	{_State97, _WildcardMarker}:                          _ReduceTildeTildeToPrefixTypeOpAction,
	{_State99, _WildcardMarker}:                          _ReduceAtomTypeToReturnableTypeAction,
	{_State100, _WildcardMarker}:                         _ReduceExplicitEnumDefToAtomTypeAction,
	{_State101, _WildcardMarker}:                         _ReduceFuncTypeToAtomTypeAction,
	{_State102, _WildcardMarker}:                         _ReduceImplicitEnumDefToAtomTypeAction,
	{_State103, _WildcardMarker}:                         _ReduceImplicitStructDefToAtomTypeAction,
	{_State104, _WildcardMarker}:                         _ReduceInitializableTypeToAtomTypeAction,
	{_State105, _WildcardMarker}:                         _ReduceParseErrorTypeToAtomTypeAction,
	{_State107, _WildcardMarker}:                         _ReducePrefixedTypeToReturnableTypeAction,
	{_State108, _WildcardMarker}:                         _ReduceReturnableTypeToValueTypeAction,
	{_State109, _WildcardMarker}:                         _ReduceTraitDefToAtomTypeAction,
	{_State110, _WildcardMarker}:                         _ReduceTraitOpTypeToValueTypeAction,
	{_State112, _WildcardMarker}:                         _ReduceDotDotDotToArgumentAction,
	{_State113, _WildcardMarker}:                         _ReduceToIdentifierExprAction,
	{_State114, _WildcardMarker}:                         _ReduceArgumentToArgumentsAction,
	{_State115, RparenToken}:                             _ReduceArgumentsToOptionalArgumentsAction,
	{_State116, _WildcardMarker}:                         _ReduceColonExpressionsToArgumentAction,
	{_State117, _WildcardMarker}:                         _ReducePositionalToArgumentAction,
	{_State117, ColonToken}:                              _ReduceExpressionToOptionalExpressionAction,
	{_State120, RparenToken}:                             _ReduceNilToOptionalExplicitFieldDefsAction,
	{_State121, RbracketToken}:                           _ReduceNilToOptionalGenericArgumentsAction,
	{_State123, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State123, ColonToken}:                              _ReduceNilToOptionalExpressionAction,
	{_State124, _WildcardMarker}:                         _ReduceToPostfixUnaryExprAction,
	{_State126, _WildcardMarker}:                         _ReduceAddToAddOpAction,
	{_State127, _WildcardMarker}:                         _ReduceBitOrToAddOpAction,
	{_State128, _WildcardMarker}:                         _ReduceBitXorToAddOpAction,
	{_State129, _WildcardMarker}:                         _ReduceSubToAddOpAction,
	{_State130, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State131, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State132, _WildcardMarker}:                         _ReduceEqualToCmpOpAction,
	{_State133, _WildcardMarker}:                         _ReduceGreaterToCmpOpAction,
	{_State134, _WildcardMarker}:                         _ReduceGreaterOrEqualToCmpOpAction,
	{_State135, _WildcardMarker}:                         _ReduceLessToCmpOpAction,
	{_State136, _WildcardMarker}:                         _ReduceLessOrEqualToCmpOpAction,
	{_State137, _WildcardMarker}:                         _ReduceNotEqualToCmpOpAction,
	{_State138, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State139, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State139, RparenToken}:                             _ReduceNilToOptionalArgumentsAction,
	{_State139, ColonToken}:                              _ReduceNilToOptionalExpressionAction,
	{_State140, _WildcardMarker}:                         _ReduceBitAndToMulOpAction,
	{_State141, _WildcardMarker}:                         _ReduceBitLshiftToMulOpAction,
	{_State142, _WildcardMarker}:                         _ReduceBitRshiftToMulOpAction,
	{_State143, _WildcardMarker}:                         _ReduceDivToMulOpAction,
	{_State144, _WildcardMarker}:                         _ReduceModToMulOpAction,
	{_State145, _WildcardMarker}:                         _ReduceMulToMulOpAction,
	{_State146, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State148, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State149, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State150, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State151, _EndMarker}:                              _ReduceIfExprToExpressionAction,
	{_State152, _EndMarker}:                              _ReduceLoopExprToExpressionAction,
	{_State153, _WildcardMarker}:                         _ReduceToBlockExprAction,
	{_State154, _EndMarker}:                              _ReduceSwitchExprToExpressionAction,
	{_State155, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State156, _WildcardMarker}:                         _ReduceToPrefixUnaryExprAction,
	{_State157, _WildcardMarker}:                         _ReduceIdentifierToVarPatternAction,
	{_State159, _WildcardMarker}:                         _ReduceTuplePatternToVarPatternAction,
	{_State160, _WildcardMarker}:                         _ReduceNilToOptionalValueTypeAction,
	{_State161, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State162, _EndMarker}:                              _ReduceNewlinesToOptionalNewlinesAction,
	{_State163, _EndMarker}:                              _ReduceDefinitionsToOptionalDefinitionsAction,
	{_State164, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State166, RbracketToken}:                           _ReduceNilToOptionalGenericParameterDefsAction,
	{_State168, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State170, _WildcardMarker}:                         _ReduceInferredRefArgToParameterDefAction,
	{_State173, _WildcardMarker}:                         _ReduceParameterDefToParameterDefsAction,
	{_State174, RparenToken}:                             _ReduceParameterDefsToOptionalParameterDefsAction,
	{_State175, _WildcardMarker}:                         _ReduceInferredToAtomTypeAction,
	{_State177, RparenToken}:                             _ReduceNilToOptionalParameterDeclsAction,
	{_State179, _WildcardMarker}:                         _ReduceNamedToAtomTypeAction,
	{_State180, _WildcardMarker}:                         _ReduceNilToOptionalGenericBindingAction,
	{_State183, _WildcardMarker}:                         _ReduceFieldDefToImplicitFieldDefsAction,
	{_State183, OrToken}:                                 _ReduceFieldDefToEnumValueDefAction,
	{_State185, RparenToken}:                             _ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction,
	{_State187, _WildcardMarker}:                         _ReduceUnsafeStatementToFieldDefAction,
	{_State188, _WildcardMarker}:                         _ReduceImplicitToFieldDefAction,
	{_State189, RparenToken}:                             _ReduceNilToOptionalTraitPropertiesAction,
	{_State190, _WildcardMarker}:                         _ReduceToPrefixedTypeAction,
	{_State191, _WildcardMarker}:                         _ReduceAddToTraitOpAction,
	{_State194, _WildcardMarker}:                         _ReduceMulToTraitOpAction,
	{_State195, _WildcardMarker}:                         _ReduceSliceToInitializableTypeAction,
	{_State196, _WildcardMarker}:                         _ReduceSubToTraitOpAction,
	{_State198, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State199, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State199, ColonToken}:                              _ReduceNilToOptionalExpressionAction,
	{_State200, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State200, ColonToken}:                              _ReduceNilToOptionalExpressionAction,
	{_State200, CommaToken}:                              _ReduceNilToOptionalExpressionAction,
	{_State200, RbracketToken}:                           _ReduceNilToOptionalExpressionAction,
	{_State200, RparenToken}:                             _ReduceNilToOptionalExpressionAction,
	{_State201, _WildcardMarker}:                         _ReduceToImplicitStructExprAction,
	{_State202, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State202, ColonToken}:                              _ReduceNilToOptionalExpressionAction,
	{_State202, CommaToken}:                              _ReduceNilToOptionalExpressionAction,
	{_State202, RbracketToken}:                           _ReduceNilToOptionalExpressionAction,
	{_State202, RparenToken}:                             _ReduceNilToOptionalExpressionAction,
	{_State203, RparenToken}:                             _ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction,
	{_State204, _WildcardMarker}:                         _ReduceFieldDefToExplicitFieldDefsAction,
	{_State206, RbracketToken}:                           _ReduceGenericArgumentsToOptionalGenericArgumentsAction,
	{_State208, _WildcardMarker}:                         _ReduceValueTypeToGenericArgumentsAction,
	{_State209, _WildcardMarker}:                         _ReduceToAccessExprAction,
	{_State211, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State211, RparenToken}:                             _ReduceNilToOptionalArgumentsAction,
	{_State211, ColonToken}:                              _ReduceNilToOptionalExpressionAction,
	{_State212, _WildcardMarker}:                         _ReduceToBinaryAddExprAction,
	{_State213, _WildcardMarker}:                         _ReduceToBinaryAndExprAction,
	{_State214, _WildcardMarker}:                         _ReduceToBinaryCmpExprAction,
	{_State216, _WildcardMarker}:                         _ReduceToBinaryMulExprAction,
	{_State217, _WildcardMarker}:                         _ReduceInfiniteToLoopExprAction,
	{_State220, _WildcardMarker}:                         _ReduceToAssignPatternAction,
	{_State220, SemicolonToken}:                          _ReduceSequenceExprToForAssignmentAction,
	{_State221, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State223, LbraceToken}:                             _ReduceSequenceExprToConditionAction,
	{_State225, _WildcardMarker}:                         _ReduceToBinaryOrExprAction,
	{_State226, _WildcardMarker}:                         _ReduceDotDotDotToFieldVarPatternAction,
	{_State227, _WildcardMarker}:                         _ReduceIdentifierToVarPatternAction,
	{_State228, _WildcardMarker}:                         _ReduceFieldVarPatternToFieldVarPatternsAction,
	{_State230, _WildcardMarker}:                         _ReducePositionalToFieldVarPatternAction,
	{_State231, _EndMarker}:                              _ReduceToVarDeclPatternAction,
	{_State232, _WildcardMarker}:                         _ReduceValueTypeToOptionalValueTypeAction,
	{_State233, _WildcardMarker}:                         _ReduceAsyncToCallbackOpAction,
	{_State234, _WildcardMarker}:                         _ReduceBreakToJumpTypeAction,
	{_State235, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State236, _WildcardMarker}:                         _ReduceContinueToJumpTypeAction,
	{_State238, _WildcardMarker}:                         _ReduceDeferToCallbackOpAction,
	{_State239, _WildcardMarker}:                         _ReduceToFallthroughStatementAction,
	{_State241, _EndMarker}:                              _ReduceToStatementBlockAction,
	{_State242, _WildcardMarker}:                         _ReduceReturnToJumpTypeAction,
	{_State243, _WildcardMarker}:                         _ReduceAccessibleExprToPostfixableExprAction,
	{_State243, LparenToken}:                             _ReduceNilToOptionalGenericBindingAction,
	{_State245, _WildcardMarker}:                         _ReduceAssignStatementToSimpleStatementBodyAction,
	{_State246, _WildcardMarker}:                         _ReduceBinaryOpAssignStatementToSimpleStatementBodyAction,
	{_State247, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State248, _WildcardMarker}:                         _ReduceCallbackOpStatementToSimpleStatementBodyAction,
	{_State249, _WildcardMarker}:                         _ReduceExpressionToExpressionsAction,
	{_State250, _WildcardMarker}:                         _ReduceExpressionOrImproperStructStatementToSimpleStatementBodyAction,
	{_State251, _WildcardMarker}:                         _ReduceToExpressionOrImproperStructStatementAction,
	{_State252, _WildcardMarker}:                         _ReduceFallthroughStatementToSimpleStatementBodyAction,
	{_State253, _WildcardMarker}:                         _ReduceImportStatementToStatementBodyAction,
	{_State254, _WildcardMarker}:                         _ReduceJumpStatementToSimpleStatementBodyAction,
	{_State255, NewlinesToken}:                           _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State255, SemicolonToken}:                          _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State255, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State256, AssignToken}:                             _ReduceToAssignPatternAction,
	{_State256, _WildcardMarker}:                         _ReduceSequenceExprToExpressionAction,
	{_State257, _WildcardMarker}:                         _ReduceSimpleStatementBodyToStatementBodyAction,
	{_State258, _WildcardMarker}:                         _ReduceAddToStatementsAction,
	{_State260, _WildcardMarker}:                         _ReduceUnaryOpAssignStatementToSimpleStatementBodyAction,
	{_State261, _WildcardMarker}:                         _ReduceUnsafeStatementToSimpleStatementBodyAction,
	{_State262, _WildcardMarker}:                         _ReduceAddToDefinitionsAction,
	{_State263, _WildcardMarker}:                         _ReduceGlobalVarAssignmentToDefinitionAction,
	{_State264, _WildcardMarker}:                         _ReduceAliasToTypeDefAction,
	{_State265, _WildcardMarker}:                         _ReduceUnconstrainedToGenericParameterDefAction,
	{_State266, _WildcardMarker}:                         _ReduceGenericParameterDefToGenericParameterDefsAction,
	{_State267, RbracketToken}:                           _ReduceGenericParameterDefsToOptionalGenericParameterDefsAction,
	{_State269, _WildcardMarker}:                         _ReduceDefinitionToTypeDefAction,
	{_State270, _EndMarker}:                              _ReduceAliasToNamedFuncDefAction,
	{_State271, RparenToken}:                             _ReduceNilToOptionalParameterDefsAction,
	{_State272, _WildcardMarker}:                         _ReduceInferredRefVarargToParameterDefAction,
	{_State273, _WildcardMarker}:                         _ReduceArgToParameterDefAction,
	{_State275, LbraceToken}:                             _ReduceNilToReturnTypeAction,
	{_State279, _WildcardMarker}:                         _ReduceFieldDefToEnumValueDefAction,
	{_State282, _WildcardMarker}:                         _ReduceNilToOptionalGenericBindingAction,
	{_State284, _WildcardMarker}:                         _ReduceParameterDeclToParameterDeclsAction,
	{_State285, RparenToken}:                             _ReduceParameterDeclsToOptionalParameterDeclsAction,
	{_State286, _WildcardMarker}:                         _ReduceUnamedToParameterDeclAction,
	{_State287, _WildcardMarker}:                         _ReduceNilToOptionalGenericBindingAction,
	{_State288, _WildcardMarker}:                         _ReduceNilToOptionalGenericBindingAction,
	{_State289, _WildcardMarker}:                         _ReduceExplicitToFieldDefAction,
	{_State294, _WildcardMarker}:                         _ReduceToImplicitEnumDefAction,
	{_State296, _WildcardMarker}:                         _ReduceToImplicitStructDefAction,
	{_State298, _WildcardMarker}:                         _ReduceFieldDefToTraitPropertyAction,
	{_State299, _WildcardMarker}:                         _ReduceMethodSignatureToTraitPropertyAction,
	{_State301, RparenToken}:                             _ReduceTraitPropertiesToOptionalTraitPropertiesAction,
	{_State302, _WildcardMarker}:                         _ReduceTraitPropertyToTraitPropertiesAction,
	{_State305, _WildcardMarker}:                         _ReduceToTraitOpTypeAction,
	{_State306, _WildcardMarker}:                         _ReduceNamedToArgumentAction,
	{_State307, _WildcardMarker}:                         _ReduceAddToArgumentsAction,
	{_State308, _WildcardMarker}:                         _ReduceExpressionToOptionalExpressionAction,
	{_State309, _WildcardMarker}:                         _ReduceAddToColonExpressionsAction,
	{_State310, _WildcardMarker}:                         _ReducePairToColonExpressionsAction,
	{_State313, _WildcardMarker}:                         _ReduceToExplicitStructDefAction,
	{_State315, _WildcardMarker}:                         _ReduceBindingToOptionalGenericBindingAction,
	{_State316, _WildcardMarker}:                         _ReduceToIndexExprAction,
	{_State318, _WildcardMarker}:                         _ReduceToInitializeExprAction,
	{_State319, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State320, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State321, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State322, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State322, SemicolonToken}:                          _ReduceNilToOptionalSequenceExprAction,
	{_State325, _WildcardMarker}:                         _ReduceVarToVarOrLetAction,
	{_State326, _WildcardMarker}:                         _ReduceAssignPatternToCasePatternAction,
	{_State327, _WildcardMarker}:                         _ReduceCasePatternToCasePatternsAction,
	{_State329, _WildcardMarker}:                         _ReduceToAssignPatternAction,
	{_State330, _WildcardMarker}:                         _ReduceNoElseToIfExprAction,
	{_State331, _EndMarker}:                              _ReduceToSwitchExprAction,
	{_State334, _WildcardMarker}:                         _ReduceToTuplePatternAction,
	{_State336, NewlinesToken}:                           _ReduceNilToOptionalSimpleStatementBodyAction,
	{_State336, SemicolonToken}:                          _ReduceNilToOptionalSimpleStatementBodyAction,
	{_State336, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State338, _WildcardMarker}:                         _ReduceStringLiteralToImportClauseAction,
	{_State339, _WildcardMarker}:                         _ReduceSingleToImportStatementAction,
	{_State340, _WildcardMarker}:                         _ReduceAddAssignToBinaryOpAssignAction,
	{_State341, _WildcardMarker}:                         _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State342, _WildcardMarker}:                         _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State343, _WildcardMarker}:                         _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State344, _WildcardMarker}:                         _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State345, _WildcardMarker}:                         _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State346, _WildcardMarker}:                         _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State347, _WildcardMarker}:                         _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State348, _WildcardMarker}:                         _ReduceDivAssignToBinaryOpAssignAction,
	{_State349, _WildcardMarker}:                         _ReduceModAssignToBinaryOpAssignAction,
	{_State350, _WildcardMarker}:                         _ReduceMulAssignToBinaryOpAssignAction,
	{_State351, _WildcardMarker}:                         _ReduceSubAssignToBinaryOpAssignAction,
	{_State352, _WildcardMarker}:                         _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State353, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State354, _WildcardMarker}:                         _ReduceToUnaryOpAssignStatementAction,
	{_State355, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State356, LparenToken}:                             _ReduceNilToOptionalGenericBindingAction,
	{_State357, NewlinesToken}:                           _ReduceToCallbackOpStatementAction,
	{_State357, SemicolonToken}:                          _ReduceToCallbackOpStatementAction,
	{_State357, _WildcardMarker}:                         _ReduceCallExprToAccessibleExprAction,
	{_State358, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State359, NewlinesToken}:                           _ReduceLabeledNoValueToJumpStatementAction,
	{_State359, SemicolonToken}:                          _ReduceLabeledNoValueToJumpStatementAction,
	{_State359, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State360, _WildcardMarker}:                         _ReduceUnlabeledValuedToJumpStatementAction,
	{_State361, _WildcardMarker}:                         _ReduceImplicitToStatementAction,
	{_State362, _WildcardMarker}:                         _ReduceExplicitToStatementAction,
	{_State363, _WildcardMarker}:                         _ReduceConstrainedToGenericParameterDefAction,
	{_State365, _WildcardMarker}:                         _ReduceGenericToOptionalGenericParametersAction,
	{_State368, _WildcardMarker}:                         _ReduceVarargToParameterDefAction,
	{_State369, LparenToken}:                             _ReduceNilToOptionalGenericParametersAction,
	{_State371, _WildcardMarker}:                         _ReduceReturnableTypeToReturnTypeAction,
	{_State372, _WildcardMarker}:                         _ReduceAddToParameterDefsAction,
	{_State375, _WildcardMarker}:                         _ReduceToExplicitEnumDefAction,
	{_State378, _WildcardMarker}:                         _ReduceUnnamedVarargToParameterDeclAction,
	{_State380, _WildcardMarker}:                         _ReduceArgToParameterDeclAction,
	{_State381, _WildcardMarker}:                         _ReduceNilToReturnTypeAction,
	{_State383, _WildcardMarker}:                         _ReduceExternNamedToAtomTypeAction,
	{_State385, _WildcardMarker}:                         _ReducePairToImplicitEnumValueDefsAction,
	{_State386, _WildcardMarker}:                         _ReduceDefaultToEnumValueDefAction,
	{_State387, _WildcardMarker}:                         _ReduceAddToImplicitEnumValueDefsAction,
	{_State388, _WildcardMarker}:                         _ReduceAddToImplicitFieldDefsAction,
	{_State390, _WildcardMarker}:                         _ReduceToTraitDefAction,
	{_State393, _WildcardMarker}:                         _ReduceMapToInitializableTypeAction,
	{_State394, _WildcardMarker}:                         _ReduceArrayToInitializableTypeAction,
	{_State395, _WildcardMarker}:                         _ReduceExplicitToExplicitFieldDefsAction,
	{_State396, _WildcardMarker}:                         _ReduceImplicitToExplicitFieldDefsAction,
	{_State397, _WildcardMarker}:                         _ReduceAddToGenericArgumentsAction,
	{_State398, _WildcardMarker}:                         _ReduceToCallExprAction,
	{_State399, _EndMarker}:                              _ReduceDoWhileToLoopExprAction,
	{_State400, SemicolonToken}:                          _ReduceAssignToForAssignmentAction,
	{_State403, DoToken}:                                 _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State404, _EndMarker}:                              _ReduceWhileToLoopExprAction,
	{_State407, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State408, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State410, _WildcardMarker}:                         _ReduceNamedToFieldVarPatternAction,
	{_State411, _WildcardMarker}:                         _ReduceAddToFieldVarPatternsAction,
	{_State412, NewlinesToken}:                           _ReduceNilToOptionalSimpleStatementBodyAction,
	{_State412, SemicolonToken}:                          _ReduceNilToOptionalSimpleStatementBodyAction,
	{_State412, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State413, _WildcardMarker}:                         _ReduceDefaultBranchStatementToStatementBodyAction,
	{_State414, _WildcardMarker}:                         _ReduceSimpleStatementBodyToOptionalSimpleStatementBodyAction,
	{_State416, _WildcardMarker}:                         _ReduceFirstToImportClausesAction,
	{_State419, _WildcardMarker}:                         _ReduceToBinaryOpAssignStatementAction,
	{_State420, _WildcardMarker}:                         _ReduceToAssignStatementAction,
	{_State421, _WildcardMarker}:                         _ReduceAddToExpressionsAction,
	{_State422, _WildcardMarker}:                         _ReduceLabeledValuedToJumpStatementAction,
	{_State423, _WildcardMarker}:                         _ReduceAddToGenericParameterDefsAction,
	{_State424, _WildcardMarker}:                         _ReduceConstrainedDefToTypeDefAction,
	{_State425, LbraceToken}:                             _ReduceNilToReturnTypeAction,
	{_State427, _WildcardMarker}:                         _ReduceToAnonymousFuncExprAction,
	{_State428, RparenToken}:                             _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State429, _WildcardMarker}:                         _ReducePairToImplicitEnumValueDefsAction,
	{_State429, RparenToken}:                             _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State430, RparenToken}:                             _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State431, _WildcardMarker}:                         _ReduceAddToImplicitEnumValueDefsAction,
	{_State431, RparenToken}:                             _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State432, _WildcardMarker}:                         _ReduceVarargToParameterDeclAction,
	{_State433, _WildcardMarker}:                         _ReduceToFuncTypeAction,
	{_State434, _WildcardMarker}:                         _ReduceAddToParameterDeclsAction,
	{_State436, RparenToken}:                             _ReduceNilToOptionalParameterDeclsAction,
	{_State437, _WildcardMarker}:                         _ReduceExplicitToTraitPropertiesAction,
	{_State438, _WildcardMarker}:                         _ReduceImplicitToTraitPropertiesAction,
	{_State440, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State440, DoToken}:                                 _ReduceNilToOptionalSequenceExprAction,
	{_State441, _WildcardMarker}:                         _ReduceEnumMatchPatternToCasePatternAction,
	{_State443, LbraceToken}:                             _ReduceCaseToConditionAction,
	{_State444, _WildcardMarker}:                         _ReduceMultipleToCasePatternsAction,
	{_State445, _EndMarker}:                              _ReduceMultiIfElseToIfExprAction,
	{_State446, _EndMarker}:                              _ReduceIfElseToIfExprAction,
	{_State447, _WildcardMarker}:                         _ReduceCaseBranchStatementToStatementBodyAction,
	{_State448, _WildcardMarker}:                         _ReduceExplicitToImportClauseTerminalAction,
	{_State449, _WildcardMarker}:                         _ReduceImplicitToImportClauseTerminalAction,
	{_State450, _WildcardMarker}:                         _ReduceMultipleToImportStatementAction,
	{_State451, _WildcardMarker}:                         _ReduceAddToImportClausesAction,
	{_State452, _WildcardMarker}:                         _ReduceAliasToImportClauseAction,
	{_State454, RparenToken}:                             _ReduceNilToOptionalParameterDefsAction,
	{_State455, _WildcardMarker}:                         _ReduceToUnsafeStatementAction,
	{_State457, _EndMarker}:                              _ReduceIteratorToLoopExprAction,
	{_State459, _WildcardMarker}:                         _ReduceEnumVarDeclPatternToCasePatternAction,
	{_State460, _EndMarker}:                              _ReduceFuncDefToNamedFuncDefAction,
	{_State462, _WildcardMarker}:                         _ReduceNilToReturnTypeAction,
	{_State464, LbraceToken}:                             _ReduceNilToReturnTypeAction,
	{_State465, _WildcardMarker}:                         _ReduceToMethodSignatureAction,
	{_State466, _EndMarker}:                              _ReduceForToLoopExprAction,
	{_State468, _EndMarker}:                              _ReduceMethodDefToNamedFuncDefAction,
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
      NEWLINES -> State 11
      source -> State 6
      optional_definitions -> State 12
      optional_newlines -> State 13

  State 2:
    Kernel Items:
      #accept: ^.package_def
    Reduce:
      (nil)
    Goto:
      PACKAGE -> State 14
      package_def -> State 7

  State 3:
    Kernel Items:
      #accept: ^.type_def
    Reduce:
      (nil)
    Goto:
      TYPE -> State 15
      type_def -> State 8

  State 4:
    Kernel Items:
      #accept: ^.named_func_def
    Reduce:
      (nil)
    Goto:
      FUNC -> State 16
      named_func_def -> State 9

  State 5:
    Kernel Items:
      #accept: ^.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 10
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 6:
    Kernel Items:
      #accept: ^ source., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

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
      #accept: ^ named_func_def., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      #accept: ^ expression., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 11:
    Kernel Items:
      optional_definitions: NEWLINES., $
      optional_newlines: NEWLINES., *
    Reduce:
      * -> [optional_newlines]
      $ -> [optional_definitions]
    Goto:
      (nil)

  State 12:
    Kernel Items:
      source: optional_definitions., $
    Reduce:
      $ -> [source]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      optional_definitions: optional_newlines.definitions optional_newlines
    Reduce:
      (nil)
    Goto:
      COMMENT_GROUPS -> State 71
      PACKAGE -> State 14
      TYPE -> State 15
      FUNC -> State 16
      VAR -> State 37
      LET -> State 27
      LBRACE -> State 72
      definitions -> State 74
      definition -> State 73
      statement_block -> State 77
      var_decl_pattern -> State 79
      var_or_let -> State 70
      type_def -> State 78
      named_func_def -> State 75
      package_def -> State 76

  State 14:
    Kernel Items:
      package_def: PACKAGE., *
      package_def: PACKAGE.statement_block
    Reduce:
      * -> [package_def]
    Goto:
      LBRACE -> State 72
      statement_block -> State 80

  State 15:
    Kernel Items:
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE.IDENTIFIER ASSIGN value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81

  State 16:
    Kernel Items:
      named_func_def: FUNC.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.IDENTIFIER ASSIGN expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 82
      LPAREN -> State 83

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
      literal_expr: FALSE., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 20:
    Kernel Items:
      literal_expr: FLOAT_LITERAL., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 21:
    Kernel Items:
      anonymous_func_expr: FUNC.LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 84

  State 22:
    Kernel Items:
      sequence_expr: GREATER.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      optional_label_decl -> State 85
      sequence_expr -> State 86
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 23:
    Kernel Items:
      identifier_expr: IDENTIFIER., *
    Reduce:
      * -> [identifier_expr]
    Goto:
      (nil)

  State 24:
    Kernel Items:
      literal_expr: INTEGER_LITERAL., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 25:
    Kernel Items:
      optional_label_decl: LABEL_DECL., *
    Reduce:
      * -> [optional_label_decl]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      initializable_type: LBRACKET.value_type RBRACKET
      initializable_type: LBRACKET.value_type COMMA INTEGER_LITERAL RBRACKET
      initializable_type: LBRACKET.value_type COLON value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 111
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 27:
    Kernel Items:
      var_or_let: LET., *
    Reduce:
      * -> [var_or_let]
    Goto:
      (nil)

  State 28:
    Kernel Items:
      implicit_struct_expr: LPAREN.optional_arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_arguments]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 113
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT_DOT_DOT -> State 112
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 117
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      optional_arguments -> State 118
      arguments -> State 115
      argument -> State 114
      colon_expressions -> State 116
      optional_expression -> State 119
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

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
      parse_error_expr: PARSE_ERROR., *
    Reduce:
      * -> [parse_error_expr]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      literal_expr: RUNE_LITERAL., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 33:
    Kernel Items:
      literal_expr: STRING_LITERAL., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 34:
    Kernel Items:
      explicit_struct_def: STRUCT.LPAREN optional_explicit_field_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 120

  State 35:
    Kernel Items:
      prefix_unary_op: SUB., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 36:
    Kernel Items:
      literal_expr: TRUE., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 37:
    Kernel Items:
      var_or_let: VAR., *
    Reduce:
      * -> [var_or_let]
    Goto:
      (nil)

  State 38:
    Kernel Items:
      accessible_expr: access_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 39:
    Kernel Items:
      call_expr: accessible_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
      postfixable_expr: accessible_expr., *
      postfix_unary_expr: accessible_expr.QUESTION
    Reduce:
      * -> [postfixable_expr]
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 123
      DOT -> State 122
      QUESTION -> State 124
      DOLLAR_LBRACKET -> State 121
      optional_generic_binding -> State 125

  State 40:
    Kernel Items:
      binary_add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 126
      SUB -> State 129
      BIT_XOR -> State 128
      BIT_OR -> State 127
      add_op -> State 130

  State 41:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 131

  State 42:
    Kernel Items:
      atom_expr: anonymous_func_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 43:
    Kernel Items:
      accessible_expr: atom_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 44:
    Kernel Items:
      add_expr: binary_add_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      (nil)

  State 45:
    Kernel Items:
      and_expr: binary_and_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      (nil)

  State 46:
    Kernel Items:
      cmp_expr: binary_cmp_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      (nil)

  State 47:
    Kernel Items:
      mul_expr: binary_mul_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      or_expr: binary_or_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      (nil)

  State 49:
    Kernel Items:
      atom_expr: block_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 50:
    Kernel Items:
      accessible_expr: call_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 51:
    Kernel Items:
      binary_cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 132
      NOT_EQUAL -> State 137
      LESS -> State 135
      LESS_OR_EQUAL -> State 136
      GREATER -> State 133
      GREATER_OR_EQUAL -> State 134
      cmp_op -> State 138

  State 52:
    Kernel Items:
      initializable_type: explicit_struct_def., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 53:
    Kernel Items:
      atom_expr: identifier_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 54:
    Kernel Items:
      atom_expr: implicit_struct_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 55:
    Kernel Items:
      accessible_expr: index_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 56:
    Kernel Items:
      initialize_expr: initializable_type.LPAREN optional_arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 139

  State 57:
    Kernel Items:
      atom_expr: initialize_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      atom_expr: literal_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 59:
    Kernel Items:
      binary_mul_expr: mul_expr.mul_op prefixable_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 145
      DIV -> State 143
      MOD -> State 144
      BIT_AND -> State 140
      BIT_LSHIFT -> State 141
      BIT_RSHIFT -> State 142
      mul_op -> State 146

  State 60:
    Kernel Items:
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
      block_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    Goto:
      IF -> State 149
      SWITCH -> State 150
      FOR -> State 148
      DO -> State 147
      LBRACE -> State 72
      statement_block -> State 153
      if_expr -> State 151
      switch_expr -> State 154
      loop_expr -> State 152

  State 61:
    Kernel Items:
      sequence_expr: or_expr., *
      binary_or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    Goto:
      OR -> State 155

  State 62:
    Kernel Items:
      atom_expr: parse_error_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 63:
    Kernel Items:
      postfixable_expr: postfix_unary_expr., *
    Reduce:
      * -> [postfixable_expr]
    Goto:
      (nil)

  State 64:
    Kernel Items:
      prefixable_expr: postfixable_expr., *
    Reduce:
      * -> [prefixable_expr]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      prefixable_expr: prefix_unary_expr., *
    Reduce:
      * -> [prefixable_expr]
    Goto:
      (nil)

  State 66:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op.prefixable_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      PARSE_ERROR -> State 31
      optional_label_decl -> State 85
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 156
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 67:
    Kernel Items:
      mul_expr: prefixable_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 68:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 69:
    Kernel Items:
      sequence_expr: var_decl_pattern., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 70:
    Kernel Items:
      var_decl_pattern: var_or_let.var_pattern optional_value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      LPAREN -> State 158
      var_pattern -> State 160
      tuple_pattern -> State 159

  State 71:
    Kernel Items:
      definition: COMMENT_GROUPS., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 72:
    Kernel Items:
      statement_block: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 161

  State 73:
    Kernel Items:
      definitions: definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 74:
    Kernel Items:
      optional_definitions: optional_newlines definitions.optional_newlines
      definitions: definitions.NEWLINES definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      NEWLINES -> State 162
      optional_newlines -> State 163

  State 75:
    Kernel Items:
      definition: named_func_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 76:
    Kernel Items:
      definition: package_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 77:
    Kernel Items:
      definition: statement_block., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 78:
    Kernel Items:
      definition: type_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 79:
    Kernel Items:
      definition: var_decl_pattern., *
      definition: var_decl_pattern.ASSIGN expression
    Reduce:
      * -> [definition]
    Goto:
      ASSIGN -> State 164

  State 80:
    Kernel Items:
      package_def: PACKAGE statement_block., $
    Reduce:
      $ -> [package_def]
    Goto:
      (nil)

  State 81:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE IDENTIFIER.ASSIGN value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 166
      ASSIGN -> State 165
      optional_generic_parameters -> State 167

  State 82:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC IDENTIFIER.ASSIGN expression
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 166
      ASSIGN -> State 168
      optional_generic_parameters -> State 169

  State 83:
    Kernel Items:
      named_func_def: FUNC LPAREN.parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 170
      parameter_def -> State 171

  State 84:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 170
      parameter_def -> State 173
      parameter_defs -> State 174
      optional_parameter_defs -> State 172

  State 85:
    Kernel Items:
      block_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 72
      statement_block -> State 153

  State 86:
    Kernel Items:
      sequence_expr: GREATER sequence_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 87:
    Kernel Items:
      prefix_type_op: BIT_AND., *
    Reduce:
      * -> [prefix_type_op]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      prefix_type_op: BIT_NEG., *
    Reduce:
      * -> [prefix_type_op]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 121
      optional_generic_binding -> State 175

  State 90:
    Kernel Items:
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 176

  State 91:
    Kernel Items:
      prefix_type_op: EXCLAIM., *
    Reduce:
      * -> [prefix_type_op]
    Goto:
      (nil)

  State 92:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 177

  State 93:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOT -> State 178
      DOLLAR_LBRACKET -> State 121
      optional_generic_binding -> State 179

  State 94:
    Kernel Items:
      implicit_struct_def: LPAREN.optional_implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      IDENTIFIER -> State 180
      UNSAFE -> State 181
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 187
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 188
      trait_op_type -> State 110
      field_def -> State 183
      implicit_field_defs -> State 185
      optional_implicit_field_defs -> State 186
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      enum_value_def -> State 182
      implicit_enum_value_defs -> State 184
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 95:
    Kernel Items:
      parse_error_type: PARSE_ERROR., *
    Reduce:
      * -> [parse_error_type]
    Goto:
      (nil)

  State 96:
    Kernel Items:
      prefix_type_op: QUESTION., *
    Reduce:
      * -> [prefix_type_op]
    Goto:
      (nil)

  State 97:
    Kernel Items:
      prefix_type_op: TILDE_TILDE., *
    Reduce:
      * -> [prefix_type_op]
    Goto:
      (nil)

  State 98:
    Kernel Items:
      trait_def: TRAIT.LPAREN optional_trait_properties RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 189

  State 99:
    Kernel Items:
      returnable_type: atom_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 100:
    Kernel Items:
      atom_type: explicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 101:
    Kernel Items:
      atom_type: func_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 102:
    Kernel Items:
      atom_type: implicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 103:
    Kernel Items:
      atom_type: implicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 104:
    Kernel Items:
      atom_type: initializable_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 105:
    Kernel Items:
      atom_type: parse_error_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 106:
    Kernel Items:
      prefixed_type: prefix_type_op.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 190
      prefixed_type -> State 107
      prefix_type_op -> State 106
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 107:
    Kernel Items:
      returnable_type: prefixed_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 108:
    Kernel Items:
      value_type: returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 109:
    Kernel Items:
      atom_type: trait_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 110:
    Kernel Items:
      value_type: trait_op_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 111:
    Kernel Items:
      initializable_type: LBRACKET value_type.RBRACKET
      initializable_type: LBRACKET value_type.COMMA INTEGER_LITERAL RBRACKET
      initializable_type: LBRACKET value_type.COLON value_type RBRACKET
      trait_op_type: value_type.trait_op returnable_type
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 195
      COMMA -> State 193
      COLON -> State 192
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 112:
    Kernel Items:
      argument: DOT_DOT_DOT., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 113:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expression
      identifier_expr: IDENTIFIER., *
    Reduce:
      * -> [identifier_expr]
    Goto:
      ASSIGN -> State 198

  State 114:
    Kernel Items:
      arguments: argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      optional_arguments: arguments., RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 199

  State 116:
    Kernel Items:
      argument: colon_expressions., *
      colon_expressions: colon_expressions.COLON optional_expression
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 200

  State 117:
    Kernel Items:
      argument: expression., *
      optional_expression: expression., COLON
    Reduce:
      * -> [argument]
      COLON -> [optional_expression]
    Goto:
      (nil)

  State 118:
    Kernel Items:
      implicit_struct_expr: LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 201

  State 119:
    Kernel Items:
      colon_expressions: optional_expression.COLON optional_expression
    Reduce:
      (nil)
    Goto:
      COLON -> State 202

  State 120:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN.optional_explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      IDENTIFIER -> State 180
      UNSAFE -> State 181
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 187
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 188
      trait_op_type -> State 110
      field_def -> State 204
      implicit_struct_def -> State 103
      explicit_field_defs -> State 203
      optional_explicit_field_defs -> State 205
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 121:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      optional_generic_arguments -> State 207
      generic_arguments -> State 206
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 208
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 122:
    Kernel Items:
      access_expr: accessible_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 209

  State 123:
    Kernel Items:
      index_expr: accessible_expr LBRACKET.argument RBRACKET
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 113
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT_DOT_DOT -> State 112
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 117
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      argument -> State 210
      colon_expressions -> State 116
      optional_expression -> State 119
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 124:
    Kernel Items:
      postfix_unary_expr: accessible_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      call_expr: accessible_expr optional_generic_binding.LPAREN optional_arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 211

  State 126:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 127:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 128:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 129:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 130:
    Kernel Items:
      binary_add_expr: add_expr add_op.mul_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      PARSE_ERROR -> State 31
      optional_label_decl -> State 85
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 212
      binary_mul_expr -> State 47
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 131:
    Kernel Items:
      binary_and_expr: and_expr AND.cmp_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      PARSE_ERROR -> State 31
      optional_label_decl -> State 85
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 213
      binary_cmp_expr -> State 46
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 132:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 133:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 134:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 135:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 136:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 137:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 138:
    Kernel Items:
      binary_cmp_expr: cmp_expr cmp_op.add_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      PARSE_ERROR -> State 31
      optional_label_decl -> State 85
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 214
      binary_add_expr -> State 44
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 139:
    Kernel Items:
      initialize_expr: initializable_type LPAREN.optional_arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_arguments]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 113
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT_DOT_DOT -> State 112
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 117
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      optional_arguments -> State 215
      arguments -> State 115
      argument -> State 114
      colon_expressions -> State 116
      optional_expression -> State 119
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 140:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 141:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 142:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 143:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 144:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 145:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 146:
    Kernel Items:
      binary_mul_expr: mul_expr mul_op.prefixable_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      PARSE_ERROR -> State 31
      optional_label_decl -> State 85
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 216
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 147:
    Kernel Items:
      loop_expr: DO.statement_block
      loop_expr: DO.statement_block FOR sequence_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 72
      statement_block -> State 217

  State 148:
    Kernel Items:
      loop_expr: FOR.sequence_expr DO statement_block
      loop_expr: FOR.assign_pattern IN sequence_expr DO statement_block
      loop_expr: FOR.for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      assign_pattern -> State 218
      optional_label_decl -> State 85
      sequence_expr -> State 220
      for_assignment -> State 219
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 149:
    Kernel Items:
      if_expr: IF.condition statement_block
      if_expr: IF.condition statement_block ELSE statement_block
      if_expr: IF.condition statement_block ELSE if_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      CASE -> State 221
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      optional_label_decl -> State 85
      sequence_expr -> State 223
      condition -> State 222
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 150:
    Kernel Items:
      switch_expr: SWITCH.sequence_expr statement_block
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      optional_label_decl -> State 85
      sequence_expr -> State 224
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 151:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 152:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 153:
    Kernel Items:
      block_expr: optional_label_decl statement_block., *
    Reduce:
      * -> [block_expr]
    Goto:
      (nil)

  State 154:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 155:
    Kernel Items:
      binary_or_expr: or_expr OR.and_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      PARSE_ERROR -> State 31
      optional_label_decl -> State 85
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 225
      binary_and_expr -> State 45
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 156:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefixable_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 157:
    Kernel Items:
      var_pattern: IDENTIFIER., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 158:
    Kernel Items:
      tuple_pattern: LPAREN.field_var_patterns RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 227
      LPAREN -> State 158
      DOT_DOT_DOT -> State 226
      var_pattern -> State 230
      tuple_pattern -> State 159
      field_var_patterns -> State 229
      field_var_pattern -> State 228

  State 159:
    Kernel Items:
      var_pattern: tuple_pattern., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 160:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern.optional_value_type
    Reduce:
      * -> [optional_value_type]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      optional_value_type -> State 231
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 232
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 161:
    Kernel Items:
      statement_block: LBRACE statements.RBRACE
      statements: statements.statement
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      CASE -> State 235
      DEFAULT -> State 237
      RETURN -> State 242
      BREAK -> State 234
      CONTINUE -> State 236
      FALLTHROUGH -> State 239
      IMPORT -> State 240
      UNSAFE -> State 181
      STRUCT -> State 34
      FUNC -> State 21
      ASYNC -> State 233
      DEFER -> State 238
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      RBRACE -> State 241
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      statement -> State 258
      simple_statement_body -> State 257
      statement_body -> State 259
      expression_or_improper_struct_statement -> State 250
      callback_op -> State 247
      callback_op_statement -> State 248
      unsafe_statement -> State 261
      jump_statement -> State 254
      jump_type -> State 255
      expressions -> State 251
      fallthrough_statement -> State 252
      assign_statement -> State 245
      unary_op_assign_statement -> State 260
      binary_op_assign_statement -> State 246
      import_statement -> State 253
      var_decl_pattern -> State 69
      var_or_let -> State 70
      assign_pattern -> State 244
      expression -> State 249
      optional_label_decl -> State 60
      sequence_expr -> State 256
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 243
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 162:
    Kernel Items:
      optional_newlines: NEWLINES., $
      definitions: definitions NEWLINES.definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      COMMENT_GROUPS -> State 71
      PACKAGE -> State 14
      TYPE -> State 15
      FUNC -> State 16
      VAR -> State 37
      LET -> State 27
      LBRACE -> State 72
      definition -> State 262
      statement_block -> State 77
      var_decl_pattern -> State 79
      var_or_let -> State 70
      type_def -> State 78
      named_func_def -> State 75
      package_def -> State 76

  State 163:
    Kernel Items:
      optional_definitions: optional_newlines definitions optional_newlines., $
    Reduce:
      $ -> [optional_definitions]
    Goto:
      (nil)

  State 164:
    Kernel Items:
      definition: var_decl_pattern ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 263
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 165:
    Kernel Items:
      type_def: TYPE IDENTIFIER ASSIGN.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 264
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 166:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.optional_generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 265
      generic_parameter_def -> State 266
      generic_parameter_defs -> State 267
      optional_generic_parameter_defs -> State 268

  State 167:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 269
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 168:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 270
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 169:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 271

  State 170:
    Kernel Items:
      parameter_def: IDENTIFIER., *
      parameter_def: IDENTIFIER.DOT_DOT_DOT
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.DOT_DOT_DOT value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      DOT_DOT_DOT -> State 272
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 273
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 171:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def.RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 274

  State 172:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 275

  State 173:
    Kernel Items:
      parameter_defs: parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 174:
    Kernel Items:
      parameter_defs: parameter_defs.COMMA parameter_def
      optional_parameter_defs: parameter_defs., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      COMMA -> State 276

  State 175:
    Kernel Items:
      atom_type: DOT optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 176:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 180
      UNSAFE -> State 181
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 187
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 188
      trait_op_type -> State 110
      field_def -> State 279
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      enum_value_def -> State 277
      implicit_enum_value_defs -> State 280
      implicit_enum_def -> State 102
      explicit_enum_value_defs -> State 278
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 177:
    Kernel Items:
      func_type: FUNC LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 282
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      DOT_DOT_DOT -> State 281
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 286
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      parameter_decl -> State 284
      parameter_decls -> State 285
      optional_parameter_decls -> State 283
      func_type -> State 101

  State 178:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 287

  State 179:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 180:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 288
      QUESTION -> State 96
      EXCLAIM -> State 91
      DOLLAR_LBRACKET -> State 121
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      optional_generic_binding -> State 179
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 289
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 181:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 290

  State 182:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 291

  State 183:
    Kernel Items:
      implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 292

  State 184:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      OR -> State 293
      RPAREN -> State 294

  State 185:
    Kernel Items:
      implicit_field_defs: implicit_field_defs.COMMA field_def
      optional_implicit_field_defs: implicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      COMMA -> State 295

  State 186:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 296

  State 187:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 188:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 189:
    Kernel Items:
      trait_def: TRAIT LPAREN.optional_trait_properties RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      IDENTIFIER -> State 180
      UNSAFE -> State 181
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 297
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 187
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 188
      trait_op_type -> State 110
      field_def -> State 298
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_property -> State 302
      trait_properties -> State 301
      optional_trait_properties -> State 300
      trait_def -> State 109
      func_type -> State 101
      method_signature -> State 299

  State 190:
    Kernel Items:
      prefixed_type: prefix_type_op returnable_type., *
    Reduce:
      * -> [prefixed_type]
    Goto:
      (nil)

  State 191:
    Kernel Items:
      trait_op: ADD., *
    Reduce:
      * -> [trait_op]
    Goto:
      (nil)

  State 192:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON.value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 303
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 193:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 304

  State 194:
    Kernel Items:
      trait_op: MUL., *
    Reduce:
      * -> [trait_op]
    Goto:
      (nil)

  State 195:
    Kernel Items:
      initializable_type: LBRACKET value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 196:
    Kernel Items:
      trait_op: SUB., *
    Reduce:
      * -> [trait_op]
    Goto:
      (nil)

  State 197:
    Kernel Items:
      trait_op_type: value_type trait_op.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 305
      prefixed_type -> State 107
      prefix_type_op -> State 106
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 198:
    Kernel Items:
      argument: IDENTIFIER ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 306
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 199:
    Kernel Items:
      arguments: arguments COMMA.argument
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 113
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT_DOT_DOT -> State 112
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 117
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      argument -> State 307
      colon_expressions -> State 116
      optional_expression -> State 119
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 200:
    Kernel Items:
      colon_expressions: colon_expressions COLON.optional_expression
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_expression]
      RBRACKET -> [optional_expression]
      COMMA -> [optional_expression]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 308
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      optional_expression -> State 309
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 201:
    Kernel Items:
      implicit_struct_expr: LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [implicit_struct_expr]
    Goto:
      (nil)

  State 202:
    Kernel Items:
      colon_expressions: optional_expression COLON.optional_expression
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_expression]
      RBRACKET -> [optional_expression]
      COMMA -> [optional_expression]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 308
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      optional_expression -> State 310
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 203:
    Kernel Items:
      explicit_field_defs: explicit_field_defs.NEWLINES field_def
      explicit_field_defs: explicit_field_defs.COMMA field_def
      optional_explicit_field_defs: explicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      NEWLINES -> State 312
      COMMA -> State 311

  State 204:
    Kernel Items:
      explicit_field_defs: field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 205:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 313

  State 206:
    Kernel Items:
      optional_generic_arguments: generic_arguments., RBRACKET
      generic_arguments: generic_arguments.COMMA value_type
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 314

  State 207:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 315

  State 208:
    Kernel Items:
      generic_arguments: value_type., *
      trait_op_type: value_type.trait_op returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 209:
    Kernel Items:
      access_expr: accessible_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 316

  State 211:
    Kernel Items:
      call_expr: accessible_expr optional_generic_binding LPAREN.optional_arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_arguments]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 113
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT_DOT_DOT -> State 112
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 117
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      optional_arguments -> State 317
      arguments -> State 115
      argument -> State 114
      colon_expressions -> State 116
      optional_expression -> State 119
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 212:
    Kernel Items:
      binary_mul_expr: mul_expr.mul_op prefixable_expr
      binary_add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [binary_add_expr]
    Goto:
      MUL -> State 145
      DIV -> State 143
      MOD -> State 144
      BIT_AND -> State 140
      BIT_LSHIFT -> State 141
      BIT_RSHIFT -> State 142
      mul_op -> State 146

  State 213:
    Kernel Items:
      binary_cmp_expr: cmp_expr.cmp_op add_expr
      binary_and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [binary_and_expr]
    Goto:
      EQUAL -> State 132
      NOT_EQUAL -> State 137
      LESS -> State 135
      LESS_OR_EQUAL -> State 136
      GREATER -> State 133
      GREATER_OR_EQUAL -> State 134
      cmp_op -> State 138

  State 214:
    Kernel Items:
      binary_add_expr: add_expr.add_op mul_expr
      binary_cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [binary_cmp_expr]
    Goto:
      ADD -> State 126
      SUB -> State 129
      BIT_XOR -> State 128
      BIT_OR -> State 127
      add_op -> State 130

  State 215:
    Kernel Items:
      initialize_expr: initializable_type LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 318

  State 216:
    Kernel Items:
      binary_mul_expr: mul_expr mul_op prefixable_expr., *
    Reduce:
      * -> [binary_mul_expr]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      loop_expr: DO statement_block., *
      loop_expr: DO statement_block.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 319

  State 218:
    Kernel Items:
      loop_expr: FOR assign_pattern.IN sequence_expr DO statement_block
      for_assignment: assign_pattern.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      IN -> State 321
      ASSIGN -> State 320

  State 219:
    Kernel Items:
      loop_expr: FOR for_assignment.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 322

  State 220:
    Kernel Items:
      assign_pattern: sequence_expr., *
      loop_expr: FOR sequence_expr.DO statement_block
      for_assignment: sequence_expr., SEMICOLON
    Reduce:
      * -> [assign_pattern]
      SEMICOLON -> [for_assignment]
    Goto:
      DO -> State 323

  State 221:
    Kernel Items:
      condition: CASE.case_patterns ASSIGN sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 325
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT -> State 324
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      case_patterns -> State 328
      var_decl_pattern -> State 69
      var_or_let -> State 70
      assign_pattern -> State 326
      case_pattern -> State 327
      optional_label_decl -> State 85
      sequence_expr -> State 329
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 222:
    Kernel Items:
      if_expr: IF condition.statement_block
      if_expr: IF condition.statement_block ELSE statement_block
      if_expr: IF condition.statement_block ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 72
      statement_block -> State 330

  State 223:
    Kernel Items:
      condition: sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 224:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 72
      statement_block -> State 331

  State 225:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      binary_or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [binary_or_expr]
    Goto:
      AND -> State 131

  State 226:
    Kernel Items:
      field_var_pattern: DOT_DOT_DOT., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 227:
    Kernel Items:
      var_pattern: IDENTIFIER., *
      field_var_pattern: IDENTIFIER.ASSIGN var_pattern
    Reduce:
      * -> [var_pattern]
    Goto:
      ASSIGN -> State 332

  State 228:
    Kernel Items:
      field_var_patterns: field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 229:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns.RPAREN
      field_var_patterns: field_var_patterns.COMMA field_var_pattern
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 334
      COMMA -> State 333

  State 230:
    Kernel Items:
      field_var_pattern: var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 231:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern optional_value_type., $
    Reduce:
      $ -> [var_decl_pattern]
    Goto:
      (nil)

  State 232:
    Kernel Items:
      optional_value_type: value_type., *
      trait_op_type: value_type.trait_op returnable_type
    Reduce:
      * -> [optional_value_type]
    Goto:
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 233:
    Kernel Items:
      callback_op: ASYNC., *
    Reduce:
      * -> [callback_op]
    Goto:
      (nil)

  State 234:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 235:
    Kernel Items:
      statement_body: CASE.case_patterns COLON optional_simple_statement_body
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 325
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT -> State 324
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      case_patterns -> State 335
      var_decl_pattern -> State 69
      var_or_let -> State 70
      assign_pattern -> State 326
      case_pattern -> State 327
      optional_label_decl -> State 85
      sequence_expr -> State 329
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 236:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 237:
    Kernel Items:
      statement_body: DEFAULT.COLON optional_simple_statement_body
    Reduce:
      (nil)
    Goto:
      COLON -> State 336

  State 238:
    Kernel Items:
      callback_op: DEFER., *
    Reduce:
      * -> [callback_op]
    Goto:
      (nil)

  State 239:
    Kernel Items:
      fallthrough_statement: FALLTHROUGH., *
    Reduce:
      * -> [fallthrough_statement]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      import_statement: IMPORT.import_clause
      import_statement: IMPORT.LPAREN import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 338
      LPAREN -> State 337
      import_clause -> State 339

  State 241:
    Kernel Items:
      statement_block: LBRACE statements RBRACE., $
    Reduce:
      $ -> [statement_block]
    Goto:
      (nil)

  State 242:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 243:
    Kernel Items:
      unary_op_assign_statement: accessible_expr.unary_op_assign
      binary_op_assign_statement: accessible_expr.binary_op_assign expression
      call_expr: accessible_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
      postfixable_expr: accessible_expr., *
      postfix_unary_expr: accessible_expr.QUESTION
    Reduce:
      * -> [postfixable_expr]
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 123
      DOT -> State 122
      QUESTION -> State 124
      DOLLAR_LBRACKET -> State 121
      ADD_ASSIGN -> State 340
      SUB_ASSIGN -> State 351
      MUL_ASSIGN -> State 350
      DIV_ASSIGN -> State 348
      MOD_ASSIGN -> State 349
      ADD_ONE_ASSIGN -> State 341
      SUB_ONE_ASSIGN -> State 352
      BIT_NEG_ASSIGN -> State 344
      BIT_AND_ASSIGN -> State 342
      BIT_OR_ASSIGN -> State 345
      BIT_XOR_ASSIGN -> State 347
      BIT_LSHIFT_ASSIGN -> State 343
      BIT_RSHIFT_ASSIGN -> State 346
      unary_op_assign -> State 354
      binary_op_assign -> State 353
      optional_generic_binding -> State 125

  State 244:
    Kernel Items:
      assign_statement: assign_pattern.ASSIGN expression
    Reduce:
      (nil)
    Goto:
      ASSIGN -> State 355

  State 245:
    Kernel Items:
      simple_statement_body: assign_statement., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 246:
    Kernel Items:
      simple_statement_body: binary_op_assign_statement., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 247:
    Kernel Items:
      callback_op_statement: callback_op.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      PARSE_ERROR -> State 31
      optional_label_decl -> State 85
      call_expr -> State 357
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 356
      access_expr -> State 38
      index_expr -> State 55
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 248:
    Kernel Items:
      simple_statement_body: callback_op_statement., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 249:
    Kernel Items:
      expressions: expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 250:
    Kernel Items:
      simple_statement_body: expression_or_improper_struct_statement., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 251:
    Kernel Items:
      expression_or_improper_struct_statement: expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [expression_or_improper_struct_statement]
    Goto:
      COMMA -> State 358

  State 252:
    Kernel Items:
      simple_statement_body: fallthrough_statement., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 253:
    Kernel Items:
      statement_body: import_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 254:
    Kernel Items:
      simple_statement_body: jump_statement., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 255:
    Kernel Items:
      jump_statement: jump_type., NEWLINES
      jump_statement: jump_type., SEMICOLON
      jump_statement: jump_type.expressions
      jump_statement: jump_type.JUMP_LABEL
      jump_statement: jump_type.JUMP_LABEL expressions
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [jump_statement]
      SEMICOLON -> [jump_statement]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      JUMP_LABEL -> State 359
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      expressions -> State 360
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 249
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 256:
    Kernel Items:
      assign_pattern: sequence_expr., ASSIGN
      expression: sequence_expr., *
    Reduce:
      * -> [expression]
      ASSIGN -> [assign_pattern]
    Goto:
      (nil)

  State 257:
    Kernel Items:
      statement_body: simple_statement_body., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 258:
    Kernel Items:
      statements: statements statement., *
    Reduce:
      * -> [statements]
    Goto:
      (nil)

  State 259:
    Kernel Items:
      statement: statement_body.NEWLINES
      statement: statement_body.SEMICOLON
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 361
      SEMICOLON -> State 362

  State 260:
    Kernel Items:
      simple_statement_body: unary_op_assign_statement., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 261:
    Kernel Items:
      simple_statement_body: unsafe_statement., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 262:
    Kernel Items:
      definitions: definitions NEWLINES definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 263:
    Kernel Items:
      definition: var_decl_pattern ASSIGN expression., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 264:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      type_def: TYPE IDENTIFIER ASSIGN value_type., *
    Reduce:
      * -> [type_def]
    Goto:
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 265:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 363
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 266:
    Kernel Items:
      generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 267:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs.COMMA generic_parameter_def
      optional_generic_parameter_defs: generic_parameter_defs., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      COMMA -> State 364

  State 268:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 365

  State 269:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., *
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 366
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 270:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN expression., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 271:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 170
      parameter_def -> State 173
      parameter_defs -> State 174
      optional_parameter_defs -> State 367

  State 272:
    Kernel Items:
      parameter_def: IDENTIFIER DOT_DOT_DOT., *
      parameter_def: IDENTIFIER DOT_DOT_DOT.value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 368
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 273:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 274:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 369

  State 275:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 371
      prefixed_type -> State 107
      prefix_type_op -> State 106
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      return_type -> State 370
      func_type -> State 101

  State 276:
    Kernel Items:
      parameter_defs: parameter_defs COMMA.parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 170
      parameter_def -> State 372

  State 277:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 373
      OR -> State 374

  State 278:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 375

  State 279:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 292

  State 280:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 376
      OR -> State 377

  State 281:
    Kernel Items:
      parameter_decl: DOT_DOT_DOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 378
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 282:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.DOT_DOT_DOT value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 288
      QUESTION -> State 96
      EXCLAIM -> State 91
      DOLLAR_LBRACKET -> State 121
      DOT_DOT_DOT -> State 379
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      optional_generic_binding -> State 179
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 380
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 283:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 381

  State 284:
    Kernel Items:
      parameter_decls: parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 285:
    Kernel Items:
      parameter_decls: parameter_decls.COMMA parameter_decl
      optional_parameter_decls: parameter_decls., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      COMMA -> State 382

  State 286:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 287:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 121
      optional_generic_binding -> State 383

  State 288:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 287
      DOLLAR_LBRACKET -> State 121
      optional_generic_binding -> State 175

  State 289:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 290:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 384

  State 291:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 180
      UNSAFE -> State 181
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 187
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 188
      trait_op_type -> State 110
      field_def -> State 279
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      enum_value_def -> State 385
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 292:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 386

  State 293:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 180
      UNSAFE -> State 181
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 187
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 188
      trait_op_type -> State 110
      field_def -> State 279
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      enum_value_def -> State 387
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 294:
    Kernel Items:
      implicit_enum_def: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_def]
    Goto:
      (nil)

  State 295:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 180
      UNSAFE -> State 181
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 187
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 188
      trait_op_type -> State 110
      field_def -> State 388
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 296:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_def]
    Goto:
      (nil)

  State 297:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 389
      LPAREN -> State 177

  State 298:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 299:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 300:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 390

  State 301:
    Kernel Items:
      trait_properties: trait_properties.NEWLINES trait_property
      trait_properties: trait_properties.COMMA trait_property
      optional_trait_properties: trait_properties., RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      NEWLINES -> State 392
      COMMA -> State 391

  State 302:
    Kernel Items:
      trait_properties: trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 303:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type.RBRACKET
      trait_op_type: value_type.trait_op returnable_type
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 393
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 304:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 394

  State 305:
    Kernel Items:
      trait_op_type: value_type trait_op returnable_type., *
    Reduce:
      * -> [trait_op_type]
    Goto:
      (nil)

  State 306:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 307:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 308:
    Kernel Items:
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      (nil)

  State 309:
    Kernel Items:
      colon_expressions: colon_expressions COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 310:
    Kernel Items:
      colon_expressions: optional_expression COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 311:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 180
      UNSAFE -> State 181
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 187
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 188
      trait_op_type -> State 110
      field_def -> State 395
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 312:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 180
      UNSAFE -> State 181
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 187
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 188
      trait_op_type -> State 110
      field_def -> State 396
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 313:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_def]
    Goto:
      (nil)

  State 314:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 397
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 315:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 316:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [index_expr]
    Goto:
      (nil)

  State 317:
    Kernel Items:
      call_expr: accessible_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 398

  State 318:
    Kernel Items:
      initialize_expr: initializable_type LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [initialize_expr]
    Goto:
      (nil)

  State 319:
    Kernel Items:
      loop_expr: DO statement_block FOR.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      optional_label_decl -> State 85
      sequence_expr -> State 399
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 320:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      optional_label_decl -> State 85
      sequence_expr -> State 400
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 321:
    Kernel Items:
      loop_expr: FOR assign_pattern IN.sequence_expr DO statement_block
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      optional_label_decl -> State 85
      sequence_expr -> State 401
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 322:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON.optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      LBRACE -> [optional_label_decl]
      SEMICOLON -> [optional_sequence_expr]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      optional_label_decl -> State 85
      sequence_expr -> State 403
      optional_sequence_expr -> State 402
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 323:
    Kernel Items:
      loop_expr: FOR sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 72
      statement_block -> State 404

  State 324:
    Kernel Items:
      case_pattern: DOT.IDENTIFIER implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 405

  State 325:
    Kernel Items:
      var_or_let: VAR., *
      case_pattern: VAR.DOT IDENTIFIER tuple_pattern
    Reduce:
      * -> [var_or_let]
    Goto:
      DOT -> State 406

  State 326:
    Kernel Items:
      case_pattern: assign_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 327:
    Kernel Items:
      case_patterns: case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 328:
    Kernel Items:
      case_patterns: case_patterns.COMMA case_pattern
      condition: CASE case_patterns.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      COMMA -> State 408
      ASSIGN -> State 407

  State 329:
    Kernel Items:
      assign_pattern: sequence_expr., *
    Reduce:
      * -> [assign_pattern]
    Goto:
      (nil)

  State 330:
    Kernel Items:
      if_expr: IF condition statement_block., *
      if_expr: IF condition statement_block.ELSE statement_block
      if_expr: IF condition statement_block.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 409

  State 331:
    Kernel Items:
      switch_expr: SWITCH sequence_expr statement_block., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN.var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      LPAREN -> State 158
      var_pattern -> State 410
      tuple_pattern -> State 159

  State 333:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA.field_var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 227
      LPAREN -> State 158
      DOT_DOT_DOT -> State 226
      var_pattern -> State 230
      tuple_pattern -> State 159
      field_var_pattern -> State 411

  State 334:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns RPAREN., *
    Reduce:
      * -> [tuple_pattern]
    Goto:
      (nil)

  State 335:
    Kernel Items:
      statement_body: CASE case_patterns.COLON optional_simple_statement_body
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    Goto:
      COMMA -> State 408
      COLON -> State 412

  State 336:
    Kernel Items:
      statement_body: DEFAULT COLON.optional_simple_statement_body
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_simple_statement_body]
      SEMICOLON -> [optional_simple_statement_body]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      RETURN -> State 242
      BREAK -> State 234
      CONTINUE -> State 236
      FALLTHROUGH -> State 239
      UNSAFE -> State 181
      STRUCT -> State 34
      FUNC -> State 21
      ASYNC -> State 233
      DEFER -> State 238
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      simple_statement_body -> State 414
      optional_simple_statement_body -> State 413
      expression_or_improper_struct_statement -> State 250
      callback_op -> State 247
      callback_op_statement -> State 248
      unsafe_statement -> State 261
      jump_statement -> State 254
      jump_type -> State 255
      expressions -> State 251
      fallthrough_statement -> State 252
      assign_statement -> State 245
      unary_op_assign_statement -> State 260
      binary_op_assign_statement -> State 246
      var_decl_pattern -> State 69
      var_or_let -> State 70
      assign_pattern -> State 244
      expression -> State 249
      optional_label_decl -> State 60
      sequence_expr -> State 256
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 243
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 337:
    Kernel Items:
      import_statement: IMPORT LPAREN.import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 338
      import_clause -> State 415
      import_clause_terminal -> State 416
      import_clauses -> State 417

  State 338:
    Kernel Items:
      import_clause: STRING_LITERAL., *
      import_clause: STRING_LITERAL.AS IDENTIFIER
    Reduce:
      * -> [import_clause]
    Goto:
      AS -> State 418

  State 339:
    Kernel Items:
      import_statement: IMPORT import_clause., *
    Reduce:
      * -> [import_statement]
    Goto:
      (nil)

  State 340:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 342:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 343:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 344:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 345:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 346:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 347:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 348:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 349:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 350:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 351:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 352:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 353:
    Kernel Items:
      binary_op_assign_statement: accessible_expr binary_op_assign.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 419
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 354:
    Kernel Items:
      unary_op_assign_statement: accessible_expr unary_op_assign., *
    Reduce:
      * -> [unary_op_assign_statement]
    Goto:
      (nil)

  State 355:
    Kernel Items:
      assign_statement: assign_pattern ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 420
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 356:
    Kernel Items:
      call_expr: accessible_expr.optional_generic_binding LPAREN optional_arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
    Reduce:
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 123
      DOT -> State 122
      DOLLAR_LBRACKET -> State 121
      optional_generic_binding -> State 125

  State 357:
    Kernel Items:
      callback_op_statement: callback_op call_expr., NEWLINES
      callback_op_statement: callback_op call_expr., SEMICOLON
      accessible_expr: call_expr., *
    Reduce:
      * -> [accessible_expr]
      NEWLINES -> [callback_op_statement]
      SEMICOLON -> [callback_op_statement]
    Goto:
      (nil)

  State 358:
    Kernel Items:
      expressions: expressions COMMA.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 421
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 359:
    Kernel Items:
      jump_statement: jump_type JUMP_LABEL., NEWLINES
      jump_statement: jump_type JUMP_LABEL., SEMICOLON
      jump_statement: jump_type JUMP_LABEL.expressions
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [jump_statement]
      SEMICOLON -> [jump_statement]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      expressions -> State 422
      var_decl_pattern -> State 69
      var_or_let -> State 70
      expression -> State 249
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 360:
    Kernel Items:
      jump_statement: jump_type expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [jump_statement]
    Goto:
      COMMA -> State 358

  State 361:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 362:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 363:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 364:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA.generic_parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 265
      generic_parameter_def -> State 423

  State 365:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 366:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 424
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 367:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 425

  State 368:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_def: IDENTIFIER DOT_DOT_DOT value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 369:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 166
      optional_generic_parameters -> State 426

  State 370:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 72
      statement_block -> State 427

  State 371:
    Kernel Items:
      return_type: returnable_type., *
    Reduce:
      * -> [return_type]
    Goto:
      (nil)

  State 372:
    Kernel Items:
      parameter_defs: parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 373:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 180
      UNSAFE -> State 181
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 187
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 188
      trait_op_type -> State 110
      field_def -> State 279
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      enum_value_def -> State 428
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 374:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 180
      UNSAFE -> State 181
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 187
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 188
      trait_op_type -> State 110
      field_def -> State 279
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      enum_value_def -> State 429
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 375:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_def]
    Goto:
      (nil)

  State 376:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 180
      UNSAFE -> State 181
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 187
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 188
      trait_op_type -> State 110
      field_def -> State 279
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      enum_value_def -> State 430
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 377:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 180
      UNSAFE -> State 181
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 187
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 188
      trait_op_type -> State 110
      field_def -> State 279
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      enum_value_def -> State 431
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 378:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_decl: DOT_DOT_DOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 379:
    Kernel Items:
      parameter_decl: IDENTIFIER DOT_DOT_DOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 432
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 380:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 381:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 371
      prefixed_type -> State 107
      prefix_type_op -> State 106
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      return_type -> State 433
      func_type -> State 101

  State 382:
    Kernel Items:
      parameter_decls: parameter_decls COMMA.parameter_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 282
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      DOT_DOT_DOT -> State 281
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 286
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      parameter_decl -> State 434
      func_type -> State 101

  State 383:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 384:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 435

  State 385:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 386:
    Kernel Items:
      enum_value_def: field_def ASSIGN DEFAULT., *
    Reduce:
      * -> [enum_value_def]
    Goto:
      (nil)

  State 387:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 388:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [implicit_field_defs]
    Goto:
      (nil)

  State 389:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 436

  State 390:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties RPAREN., *
    Reduce:
      * -> [trait_def]
    Goto:
      (nil)

  State 391:
    Kernel Items:
      trait_properties: trait_properties COMMA.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 180
      UNSAFE -> State 181
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 297
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 187
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 188
      trait_op_type -> State 110
      field_def -> State 298
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_property -> State 437
      trait_def -> State 109
      func_type -> State 101
      method_signature -> State 299

  State 392:
    Kernel Items:
      trait_properties: trait_properties NEWLINES.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 180
      UNSAFE -> State 181
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 297
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 187
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 188
      trait_op_type -> State 110
      field_def -> State 298
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_property -> State 438
      trait_def -> State 109
      func_type -> State 101
      method_signature -> State 299

  State 393:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 394:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 395:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 396:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 397:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
      trait_op_type: value_type.trait_op returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 398:
    Kernel Items:
      call_expr: accessible_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 399:
    Kernel Items:
      loop_expr: DO statement_block FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 400:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN sequence_expr., SEMICOLON
    Reduce:
      SEMICOLON -> [for_assignment]
    Goto:
      (nil)

  State 401:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 439

  State 402:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 440

  State 403:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 404:
    Kernel Items:
      loop_expr: FOR sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 405:
    Kernel Items:
      case_pattern: DOT IDENTIFIER.implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 28
      implicit_struct_expr -> State 441

  State 406:
    Kernel Items:
      case_pattern: VAR DOT.IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 442

  State 407:
    Kernel Items:
      condition: CASE case_patterns ASSIGN.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      optional_label_decl -> State 85
      sequence_expr -> State 443
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 408:
    Kernel Items:
      case_patterns: case_patterns COMMA.case_pattern
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 325
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT -> State 324
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      assign_pattern -> State 326
      case_pattern -> State 444
      optional_label_decl -> State 85
      sequence_expr -> State 329
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 409:
    Kernel Items:
      if_expr: IF condition statement_block ELSE.statement_block
      if_expr: IF condition statement_block ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 149
      LBRACE -> State 72
      statement_block -> State 446
      if_expr -> State 445

  State 410:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 411:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 412:
    Kernel Items:
      statement_body: CASE case_patterns COLON.optional_simple_statement_body
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_simple_statement_body]
      SEMICOLON -> [optional_simple_statement_body]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      RETURN -> State 242
      BREAK -> State 234
      CONTINUE -> State 236
      FALLTHROUGH -> State 239
      UNSAFE -> State 181
      STRUCT -> State 34
      FUNC -> State 21
      ASYNC -> State 233
      DEFER -> State 238
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      simple_statement_body -> State 414
      optional_simple_statement_body -> State 447
      expression_or_improper_struct_statement -> State 250
      callback_op -> State 247
      callback_op_statement -> State 248
      unsafe_statement -> State 261
      jump_statement -> State 254
      jump_type -> State 255
      expressions -> State 251
      fallthrough_statement -> State 252
      assign_statement -> State 245
      unary_op_assign_statement -> State 260
      binary_op_assign_statement -> State 246
      var_decl_pattern -> State 69
      var_or_let -> State 70
      assign_pattern -> State 244
      expression -> State 249
      optional_label_decl -> State 60
      sequence_expr -> State 256
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 243
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 413:
    Kernel Items:
      statement_body: DEFAULT COLON optional_simple_statement_body., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 414:
    Kernel Items:
      optional_simple_statement_body: simple_statement_body., *
    Reduce:
      * -> [optional_simple_statement_body]
    Goto:
      (nil)

  State 415:
    Kernel Items:
      import_clause_terminal: import_clause.NEWLINES
      import_clause_terminal: import_clause.COMMA
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 449
      COMMA -> State 448

  State 416:
    Kernel Items:
      import_clauses: import_clause_terminal., *
    Reduce:
      * -> [import_clauses]
    Goto:
      (nil)

  State 417:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses.RPAREN
      import_clauses: import_clauses.import_clause_terminal
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 338
      RPAREN -> State 450
      import_clause -> State 415
      import_clause_terminal -> State 451

  State 418:
    Kernel Items:
      import_clause: STRING_LITERAL AS.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 452

  State 419:
    Kernel Items:
      binary_op_assign_statement: accessible_expr binary_op_assign expression., *
    Reduce:
      * -> [binary_op_assign_statement]
    Goto:
      (nil)

  State 420:
    Kernel Items:
      assign_statement: assign_pattern ASSIGN expression., *
    Reduce:
      * -> [assign_statement]
    Goto:
      (nil)

  State 421:
    Kernel Items:
      expressions: expressions COMMA expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 422:
    Kernel Items:
      jump_statement: jump_type JUMP_LABEL expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [jump_statement]
    Goto:
      COMMA -> State 358

  State 423:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 424:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., *
    Reduce:
      * -> [type_def]
    Goto:
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 425:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 371
      prefixed_type -> State 107
      prefix_type_op -> State 106
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      return_type -> State 453
      func_type -> State 101

  State 426:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 454

  State 427:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 428:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 429:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
      explicit_enum_value_defs: enum_value_def OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 430:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 431:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 432:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_decl: IDENTIFIER DOT_DOT_DOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 191
      SUB -> State 196
      MUL -> State 194
      trait_op -> State 197

  State 433:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [func_type]
    Goto:
      (nil)

  State 434:
    Kernel Items:
      parameter_decls: parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 435:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 455

  State 436:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 282
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      DOT_DOT_DOT -> State 281
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 286
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      parameter_decl -> State 284
      parameter_decls -> State 285
      optional_parameter_decls -> State 456
      func_type -> State 101

  State 437:
    Kernel Items:
      trait_properties: trait_properties COMMA trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 438:
    Kernel Items:
      trait_properties: trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 439:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 72
      statement_block -> State 457

  State 440:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON.optional_sequence_expr DO statement_block
    Reduce:
      DO -> [optional_sequence_expr]
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 23
      TRUE -> State 36
      FALSE -> State 19
      STRUCT -> State 34
      FUNC -> State 21
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 69
      var_or_let -> State 70
      optional_label_decl -> State 85
      sequence_expr -> State 403
      optional_sequence_expr -> State 458
      call_expr -> State 50
      atom_expr -> State 43
      parse_error_expr -> State 62
      literal_expr -> State 58
      identifier_expr -> State 53
      block_expr -> State 49
      initialize_expr -> State 57
      implicit_struct_expr -> State 54
      accessible_expr -> State 39
      access_expr -> State 38
      index_expr -> State 55
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 47
      add_expr -> State 40
      binary_add_expr -> State 44
      cmp_expr -> State 51
      binary_cmp_expr -> State 46
      and_expr -> State 41
      binary_and_expr -> State 45
      or_expr -> State 61
      binary_or_expr -> State 48
      initializable_type -> State 56
      explicit_struct_def -> State 52
      anonymous_func_expr -> State 42

  State 441:
    Kernel Items:
      case_pattern: DOT IDENTIFIER implicit_struct_expr., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 442:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER.tuple_pattern
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 158
      tuple_pattern -> State 459

  State 443:
    Kernel Items:
      condition: CASE case_patterns ASSIGN sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 444:
    Kernel Items:
      case_patterns: case_patterns COMMA case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 445:
    Kernel Items:
      if_expr: IF condition statement_block ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 446:
    Kernel Items:
      if_expr: IF condition statement_block ELSE statement_block., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 447:
    Kernel Items:
      statement_body: CASE case_patterns COLON optional_simple_statement_body., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 448:
    Kernel Items:
      import_clause_terminal: import_clause COMMA., *
    Reduce:
      * -> [import_clause_terminal]
    Goto:
      (nil)

  State 449:
    Kernel Items:
      import_clause_terminal: import_clause NEWLINES., *
    Reduce:
      * -> [import_clause_terminal]
    Goto:
      (nil)

  State 450:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses RPAREN., *
    Reduce:
      * -> [import_statement]
    Goto:
      (nil)

  State 451:
    Kernel Items:
      import_clauses: import_clauses import_clause_terminal., *
    Reduce:
      * -> [import_clauses]
    Goto:
      (nil)

  State 452:
    Kernel Items:
      import_clause: STRING_LITERAL AS IDENTIFIER., *
    Reduce:
      * -> [import_clause]
    Goto:
      (nil)

  State 453:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 72
      statement_block -> State 460

  State 454:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 170
      parameter_def -> State 173
      parameter_defs -> State 174
      optional_parameter_defs -> State 461

  State 455:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 456:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 462

  State 457:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 458:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 463

  State 459:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER tuple_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 460:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 461:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 464

  State 462:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 371
      prefixed_type -> State 107
      prefix_type_op -> State 106
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      return_type -> State 465
      func_type -> State 101

  State 463:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 72
      statement_block -> State 466

  State 464:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 34
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 26
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 371
      prefixed_type -> State 107
      prefix_type_op -> State 106
      implicit_struct_def -> State 103
      explicit_struct_def -> State 52
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      return_type -> State 467
      func_type -> State 101

  State 465:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 466:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 467:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 72
      statement_block -> State 468

  State 468:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

Number of states: 468
Number of shift actions: 3728
Number of reduce actions: 394
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 3902
Number of unoptimized shift actions: 31484
Number of unoptimized reduce actions: 26006
*/
