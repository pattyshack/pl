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

type SimpleStatementBodyReducer interface {

	// 113:2: simple_statement_body -> expression_or_implicit_struct: ...
	ExpressionOrImplicitStructToSimpleStatementBody(Expressions_ GenericSymbol) (Statement, error)

	// 121:2: simple_statement_body -> async: ...
	AsyncToSimpleStatementBody(Async_ TokenValue, CallExpr_ Expression) (Statement, error)

	// 139:2: simple_statement_body -> defer: ...
	DeferToSimpleStatementBody(Defer_ TokenValue, CallExpr_ Expression) (Statement, error)

	// 143:2: simple_statement_body -> jump_statement: ...
	JumpStatementToSimpleStatementBody(JumpType_ TokenValue, OptionalJumpLabel_ TokenValue, OptionalExpressions_ GenericSymbol) (Statement, error)

	// 148:2: simple_statement_body -> assign_statement: ...
	AssignStatementToSimpleStatementBody(AssignPattern_ GenericSymbol, Assign_ TokenValue, Expression_ Expression) (Statement, error)

	// 154:2: simple_statement_body -> unary_op_assign_statement: ...
	UnaryOpAssignStatementToSimpleStatementBody(AccessExpr_ Expression, UnaryOpAssign_ TokenValue) (Statement, error)

	// 155:2: simple_statement_body -> binary_op_assign_statement: ...
	BinaryOpAssignStatementToSimpleStatementBody(AccessExpr_ Expression, BinaryOpAssign_ TokenValue, Expression_ Expression) (Statement, error)

	// 159:2: simple_statement_body -> FALLTHROUGH: ...
	FallthroughToSimpleStatementBody(Fallthrough_ TokenValue) (Statement, error)
}

type StatementBodyReducer interface {

	// 170:2: statement_body -> case_branch_statement: ...
	CaseBranchStatementToStatementBody(Case_ TokenValue, CasePatterns_ GenericSymbol, Colon_ TokenValue, OptionalSimpleStatementBody_ Statement) (Statement, error)

	// 171:2: statement_body -> default_branch_statement: ...
	DefaultBranchStatementToStatementBody(Default_ TokenValue, Colon_ TokenValue, OptionalSimpleStatementBody_ Statement) (Statement, error)
}

type OptionalSimpleStatementBodyReducer interface {

	// 175:2: optional_simple_statement_body -> nil: ...
	NilToOptionalSimpleStatementBody() (Statement, error)
}

type UnsafeStatementReducer interface {
	// 204:31: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ TokenValue, Less_ TokenValue, Identifier_ TokenValue, Greater_ TokenValue, StringLiteral_ TokenValue) (Statement, error)
}

type OptionalJumpLabelReducer interface {
	// 216:2: optional_jump_label -> JUMP_LABEL: ...
	JumpLabelToOptionalJumpLabel(JumpLabel_ TokenValue) (TokenValue, error)

	// 217:2: optional_jump_label -> unlabelled: ...
	UnlabelledToOptionalJumpLabel() (TokenValue, error)
}

type ExpressionsReducer interface {
	// 220:2: expressions -> expression: ...
	ExpressionToExpressions(Expression_ Expression) (GenericSymbol, error)

	// 221:2: expressions -> add: ...
	AddToExpressions(Expressions_ GenericSymbol, Comma_ TokenValue, Expression_ Expression) (GenericSymbol, error)
}

type OptionalExpressionsReducer interface {
	// 224:2: optional_expressions -> expressions: ...
	ExpressionsToOptionalExpressions(Expressions_ GenericSymbol) (GenericSymbol, error)

	// 225:2: optional_expressions -> nil: ...
	NilToOptionalExpressions() (GenericSymbol, error)
}

type ImportStatementReducer interface {
	// 232:2: import_statement -> single: ...
	SingleToImportStatement(Import_ TokenValue, ImportClause_ GenericSymbol) (Statement, error)

	// 233:2: import_statement -> multiple: ...
	MultipleToImportStatement(Import_ TokenValue, Lparen_ TokenValue, ImportClauses_ GenericSymbol, Rparen_ TokenValue) (Statement, error)
}

type ImportClauseReducer interface {
	// 236:2: import_clause -> STRING_LITERAL: ...
	StringLiteralToImportClause(StringLiteral_ TokenValue) (GenericSymbol, error)

	// 237:2: import_clause -> alias: ...
	AliasToImportClause(StringLiteral_ TokenValue, As_ TokenValue, Identifier_ TokenValue) (GenericSymbol, error)
}

type ImportClauseTerminalReducer interface {
	// 240:2: import_clause_terminal -> implicit: ...
	ImplicitToImportClauseTerminal(ImportClause_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)

	// 241:2: import_clause_terminal -> explicit: ...
	ExplicitToImportClauseTerminal(ImportClause_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)
}

type ImportClausesReducer interface {
	// 244:2: import_clauses -> first: ...
	FirstToImportClauses(ImportClauseTerminal_ GenericSymbol) (GenericSymbol, error)

	// 245:2: import_clauses -> add: ...
	AddToImportClauses(ImportClauses_ GenericSymbol, ImportClauseTerminal_ GenericSymbol) (GenericSymbol, error)
}

type CasePatternsReducer interface {
	// 252:2: case_patterns -> case_pattern: ...
	CasePatternToCasePatterns(CasePattern_ GenericSymbol) (GenericSymbol, error)

	// 253:2: case_patterns -> multiple: ...
	MultipleToCasePatterns(CasePatterns_ GenericSymbol, Comma_ TokenValue, CasePattern_ GenericSymbol) (GenericSymbol, error)
}

type VarDeclPatternReducer interface {
	// 262:20: var_decl_pattern -> ...
	ToVarDeclPattern(VarOrLet_ TokenValue, VarPattern_ GenericSymbol, OptionalValueType_ TypeExpression) (GenericSymbol, error)
}

type VarPatternReducer interface {
	// 269:2: var_pattern -> IDENTIFIER: ...
	IdentifierToVarPattern(Identifier_ TokenValue) (GenericSymbol, error)

	// 270:2: var_pattern -> tuple_pattern: ...
	TuplePatternToVarPattern(TuplePattern_ GenericSymbol) (GenericSymbol, error)
}

type TuplePatternReducer interface {
	// 272:17: tuple_pattern -> ...
	ToTuplePattern(Lparen_ TokenValue, FieldVarPatterns_ GenericSymbol, Rparen_ TokenValue) (GenericSymbol, error)
}

type FieldVarPatternsReducer interface {
	// 275:2: field_var_patterns -> field_var_pattern: ...
	FieldVarPatternToFieldVarPatterns(FieldVarPattern_ GenericSymbol) (GenericSymbol, error)

	// 276:2: field_var_patterns -> add: ...
	AddToFieldVarPatterns(FieldVarPatterns_ GenericSymbol, Comma_ TokenValue, FieldVarPattern_ GenericSymbol) (GenericSymbol, error)
}

type FieldVarPatternReducer interface {
	// 279:2: field_var_pattern -> positional: ...
	PositionalToFieldVarPattern(VarPattern_ GenericSymbol) (GenericSymbol, error)

	// 280:2: field_var_pattern -> named: ...
	NamedToFieldVarPattern(Identifier_ TokenValue, Assign_ TokenValue, VarPattern_ GenericSymbol) (GenericSymbol, error)

	// 281:2: field_var_pattern -> DOT_DOT_DOT: ...
	DotDotDotToFieldVarPattern(DotDotDot_ TokenValue) (GenericSymbol, error)
}

type OptionalValueTypeReducer interface {

	// 285:2: optional_value_type -> nil: ...
	NilToOptionalValueType() (TypeExpression, error)
}

type AssignPatternReducer interface {
	// 292:2: assign_pattern -> ...
	ToAssignPattern(SequenceExpr_ Expression) (GenericSymbol, error)
}

type CasePatternReducer interface {

	// 302:2: case_pattern -> enum_match_pattern: ...
	EnumMatchPatternToCasePattern(Dot_ TokenValue, Identifier_ TokenValue, ImplicitStructExpr_ Expression) (GenericSymbol, error)

	// 303:2: case_pattern -> enum_var_decl_pattern: ...
	EnumVarDeclPatternToCasePattern(Var_ TokenValue, Dot_ TokenValue, Identifier_ TokenValue, TuplePattern_ GenericSymbol) (GenericSymbol, error)
}

type ExpressionReducer interface {
	// 310:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ TokenValue, IfExpr_ Expression) (Expression, error)

	// 311:2: expression -> switch_expr: ...
	SwitchExprToExpression(OptionalLabelDecl_ TokenValue, SwitchExpr_ GenericSymbol) (Expression, error)

	// 312:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ TokenValue, LoopExpr_ Expression) (Expression, error)
}

type OptionalLabelDeclReducer interface {
	// 316:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ TokenValue) (TokenValue, error)

	// 317:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (TokenValue, error)
}

type SequenceExprReducer interface {

	// 327:2: sequence_expr -> var_decl_pattern: ...
	VarDeclPatternToSequenceExpr(VarDeclPattern_ GenericSymbol) (Expression, error)

	// 331:2: sequence_expr -> assign_var_pattern: ...
	AssignVarPatternToSequenceExpr(Greater_ TokenValue, SequenceExpr_ Expression) (Expression, error)
}

type IfExprReducer interface {
	// 340:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol) (Expression, error)

	// 341:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol, Else_ TokenValue, StatementBlock_2 GenericSymbol) (Expression, error)

	// 342:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol, Else_ TokenValue, IfExpr_ Expression) (Expression, error)
}

type ConditionReducer interface {
	// 345:2: condition -> sequence_expr: ...
	SequenceExprToCondition(SequenceExpr_ Expression) (GenericSymbol, error)

	// 346:2: condition -> case: ...
	CaseToCondition(Case_ TokenValue, CasePatterns_ GenericSymbol, Assign_ TokenValue, SequenceExpr_ Expression) (GenericSymbol, error)
}

type SwitchExprReducer interface {
	// 370:15: switch_expr -> ...
	ToSwitchExpr(Switch_ TokenValue, SequenceExpr_ Expression, StatementBlock_ GenericSymbol) (GenericSymbol, error)
}

type LoopExprReducer interface {
	// 384:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 385:2: loop_expr -> do_while: ...
	DoWhileToLoopExpr(Do_ TokenValue, StatementBlock_ GenericSymbol, For_ TokenValue, SequenceExpr_ Expression) (Expression, error)

	// 386:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ TokenValue, SequenceExpr_ Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 387:2: loop_expr -> iterator: ...
	IteratorToLoopExpr(For_ TokenValue, AssignPattern_ GenericSymbol, In_ TokenValue, SequenceExpr_ Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 388:2: loop_expr -> for: ...
	ForToLoopExpr(For_ TokenValue, ForAssignment_ GenericSymbol, Semicolon_ TokenValue, OptionalSequenceExpr_ Expression, Semicolon_2 TokenValue, OptionalSequenceExpr_2 Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)
}

type OptionalSequenceExprReducer interface {
	// 391:2: optional_sequence_expr -> sequence_expr: ...
	SequenceExprToOptionalSequenceExpr(SequenceExpr_ Expression) (Expression, error)

	// 392:2: optional_sequence_expr -> nil: ...
	NilToOptionalSequenceExpr() (Expression, error)
}

type ForAssignmentReducer interface {
	// 395:2: for_assignment -> sequence_expr: ...
	SequenceExprToForAssignment(SequenceExpr_ Expression) (GenericSymbol, error)

	// 396:2: for_assignment -> assign: ...
	AssignToForAssignment(AssignPattern_ GenericSymbol, Assign_ TokenValue, SequenceExpr_ Expression) (GenericSymbol, error)
}

type CallExprReducer interface {
	// 403:2: call_expr -> ...
	ToCallExpr(AccessExpr_ Expression, OptionalGenericBinding_ GenericSymbol, Lparen_ TokenValue, OptionalArguments_ GenericSymbol, Rparen_ TokenValue) (Expression, error)
}

type OptionalGenericBindingReducer interface {
	// 406:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ TokenValue, OptionalGenericArguments_ GenericSymbol, Rbracket_ TokenValue) (GenericSymbol, error)

	// 407:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (GenericSymbol, error)
}

type OptionalGenericArgumentsReducer interface {
	// 410:2: optional_generic_arguments -> generic_arguments: ...
	GenericArgumentsToOptionalGenericArguments(GenericArguments_ GenericSymbol) (GenericSymbol, error)

	// 411:2: optional_generic_arguments -> nil: ...
	NilToOptionalGenericArguments() (GenericSymbol, error)
}

type GenericArgumentsReducer interface {
	// 415:2: generic_arguments -> value_type: ...
	ValueTypeToGenericArguments(ValueType_ TypeExpression) (GenericSymbol, error)

	// 416:2: generic_arguments -> add: ...
	AddToGenericArguments(GenericArguments_ GenericSymbol, Comma_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)
}

type OptionalArgumentsReducer interface {
	// 419:2: optional_arguments -> arguments: ...
	ArgumentsToOptionalArguments(Arguments_ GenericSymbol) (GenericSymbol, error)

	// 420:2: optional_arguments -> nil: ...
	NilToOptionalArguments() (GenericSymbol, error)
}

type ArgumentsReducer interface {
	// 423:2: arguments -> argument: ...
	ArgumentToArguments(Argument_ GenericSymbol) (GenericSymbol, error)

	// 424:2: arguments -> add: ...
	AddToArguments(Arguments_ GenericSymbol, Comma_ TokenValue, Argument_ GenericSymbol) (GenericSymbol, error)
}

type ArgumentReducer interface {
	// 427:2: argument -> positional: ...
	PositionalToArgument(Expression_ Expression) (GenericSymbol, error)

	// 428:2: argument -> named: ...
	NamedToArgument(Identifier_ TokenValue, Assign_ TokenValue, Expression_ Expression) (GenericSymbol, error)

	// 429:2: argument -> colon_expressions: ...
	ColonExpressionsToArgument(ColonExpressions_ GenericSymbol) (GenericSymbol, error)

	// 432:2: argument -> DOT_DOT_DOT: ...
	DotDotDotToArgument(DotDotDot_ TokenValue) (GenericSymbol, error)
}

type ColonExpressionsReducer interface {
	// 435:2: colon_expressions -> pair: ...
	PairToColonExpressions(OptionalExpression_ Expression, Colon_ TokenValue, OptionalExpression_2 Expression) (GenericSymbol, error)

	// 436:2: colon_expressions -> add: ...
	AddToColonExpressions(ColonExpressions_ GenericSymbol, Colon_ TokenValue, OptionalExpression_ Expression) (GenericSymbol, error)
}

type OptionalExpressionReducer interface {
	// 439:2: optional_expression -> expression: ...
	ExpressionToOptionalExpression(Expression_ Expression) (Expression, error)

	// 440:2: optional_expression -> nil: ...
	NilToOptionalExpression() (Expression, error)
}

type AtomExprReducer interface {

	// 452:2: atom_expr -> IDENTIFIER: ...
	IdentifierToAtomExpr(Identifier_ TokenValue) (Expression, error)

	// 453:2: atom_expr -> block_expr: ...
	BlockExprToAtomExpr(OptionalLabelDecl_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 455:2: atom_expr -> initialize_expr: ...
	InitializeExprToAtomExpr(InitializableType_ TypeExpression, Lparen_ TokenValue, OptionalArguments_ GenericSymbol, Rparen_ TokenValue) (Expression, error)

	// 457:2: atom_expr -> PARSE_ERROR: ...
	ParseErrorToAtomExpr(ParseError_ ParseErrorSymbol) (Expression, error)
}

type LiteralExprReducer interface {
	// 460:2: literal_expr -> TRUE: ...
	TrueToLiteralExpr(True_ TokenValue) (Expression, error)

	// 461:2: literal_expr -> FALSE: ...
	FalseToLiteralExpr(False_ TokenValue) (Expression, error)

	// 462:2: literal_expr -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteralExpr(IntegerLiteral_ TokenValue) (Expression, error)

	// 463:2: literal_expr -> FLOAT_LITERAL: ...
	FloatLiteralToLiteralExpr(FloatLiteral_ TokenValue) (Expression, error)

	// 464:2: literal_expr -> RUNE_LITERAL: ...
	RuneLiteralToLiteralExpr(RuneLiteral_ TokenValue) (Expression, error)

	// 465:2: literal_expr -> STRING_LITERAL: ...
	StringLiteralToLiteralExpr(StringLiteral_ TokenValue) (Expression, error)
}

type ImplicitStructExprReducer interface {
	// 467:36: implicit_struct_expr -> ...
	ToImplicitStructExpr(Lparen_ TokenValue, OptionalArguments_ GenericSymbol, Rparen_ TokenValue) (Expression, error)
}

type AccessExprReducer interface {

	// 471:2: access_expr -> access: ...
	AccessToAccessExpr(AccessExpr_ Expression, Dot_ TokenValue, Identifier_ TokenValue) (Expression, error)

	// 473:2: access_expr -> index: ...
	IndexToAccessExpr(AccessExpr_ Expression, Lbracket_ TokenValue, Argument_ GenericSymbol, Rbracket_ TokenValue) (Expression, error)
}

type PostfixUnaryExprReducer interface {

	// 477:2: postfix_unary_expr -> ...
	ToPostfixUnaryExpr(AccessExpr_ Expression, Question_ TokenValue) (Expression, error)
}

type PrefixUnaryExprReducer interface {

	// 492:2: prefix_unary_expr -> ...
	ToPrefixUnaryExpr(PrefixUnaryOp_ TokenValue, PrefixUnaryExpr_ Expression) (Expression, error)
}

type MulExprReducer interface {

	// 504:2: mul_expr -> ...
	ToMulExpr(MulExpr_ Expression, MulOp_ TokenValue, PrefixUnaryExpr_ Expression) (Expression, error)
}

type AddExprReducer interface {

	// 514:2: add_expr -> ...
	ToAddExpr(AddExpr_ Expression, AddOp_ TokenValue, MulExpr_ Expression) (Expression, error)
}

type CmpExprReducer interface {

	// 526:2: cmp_expr -> ...
	ToCmpExpr(CmpExpr_ Expression, CmpOp_ TokenValue, AddExpr_ Expression) (Expression, error)
}

type AndExprReducer interface {

	// 530:2: and_expr -> ...
	ToAndExpr(AndExpr_ Expression, And_ TokenValue, CmpExpr_ Expression) (Expression, error)
}

type OrExprReducer interface {

	// 534:2: or_expr -> ...
	ToOrExpr(OrExpr_ Expression, Or_ TokenValue, AndExpr_ Expression) (Expression, error)
}

type InitializableTypeReducer interface {

	// 544:2: initializable_type -> slice: ...
	SliceToInitializableType(Lbracket_ TokenValue, ValueType_ TypeExpression, Rbracket_ TokenValue) (TypeExpression, error)

	// 545:2: initializable_type -> array: ...
	ArrayToInitializableType(Lbracket_ TokenValue, ValueType_ TypeExpression, Comma_ TokenValue, IntegerLiteral_ TokenValue, Rbracket_ TokenValue) (TypeExpression, error)

	// 546:2: initializable_type -> map: ...
	MapToInitializableType(Lbracket_ TokenValue, ValueType_ TypeExpression, Colon_ TokenValue, ValueType_2 TypeExpression, Rbracket_ TokenValue) (TypeExpression, error)
}

type AtomTypeReducer interface {

	// 550:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ TokenValue, OptionalGenericBinding_ GenericSymbol) (TypeExpression, error)

	// 551:2: atom_type -> extern_named: ...
	ExternNamedToAtomType(Identifier_ TokenValue, Dot_ TokenValue, Identifier_2 TokenValue, OptionalGenericBinding_ GenericSymbol) (TypeExpression, error)

	// 552:2: atom_type -> inferred: ...
	InferredToAtomType(Dot_ TokenValue, OptionalGenericBinding_ GenericSymbol) (TypeExpression, error)

	// 558:2: atom_type -> PARSE_ERROR: ...
	ParseErrorToAtomType(ParseError_ ParseErrorSymbol) (TypeExpression, error)
}

type ReturnableTypeReducer interface {

	// 565:2: returnable_type -> optional: ...
	OptionalToReturnableType(Question_ TokenValue, ReturnableType_ TypeExpression) (TypeExpression, error)

	// 566:2: returnable_type -> result: ...
	ResultToReturnableType(Exclaim_ TokenValue, ReturnableType_ TypeExpression) (TypeExpression, error)

	// 567:2: returnable_type -> reference: ...
	ReferenceToReturnableType(BitAnd_ TokenValue, ReturnableType_ TypeExpression) (TypeExpression, error)

	// 568:2: returnable_type -> public_methods_trait: ...
	PublicMethodsTraitToReturnableType(BitNeg_ TokenValue, ReturnableType_ TypeExpression) (TypeExpression, error)

	// 569:2: returnable_type -> public_trait: ...
	PublicTraitToReturnableType(TildeTilde_ TokenValue, ReturnableType_ TypeExpression) (TypeExpression, error)
}

type ValueTypeReducer interface {

	// 575:2: value_type -> trait_intersect: ...
	TraitIntersectToValueType(ValueType_ TypeExpression, Mul_ TokenValue, ReturnableType_ TypeExpression) (TypeExpression, error)

	// 576:2: value_type -> trait_union: ...
	TraitUnionToValueType(ValueType_ TypeExpression, Add_ TokenValue, ReturnableType_ TypeExpression) (TypeExpression, error)

	// 577:2: value_type -> trait_difference: ...
	TraitDifferenceToValueType(ValueType_ TypeExpression, Sub_ TokenValue, ReturnableType_ TypeExpression) (TypeExpression, error)
}

type TypeDefReducer interface {
	// 580:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, ValueType_ TypeExpression) (SourceDefinition, error)

	// 581:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, ValueType_ TypeExpression, Implements_ TokenValue, ValueType_2 TypeExpression) (SourceDefinition, error)

	// 582:2: type_def -> alias: ...
	AliasToTypeDef(Type_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, ValueType_ TypeExpression) (SourceDefinition, error)
}

type GenericParameterDefReducer interface {
	// 590:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ TokenValue) (GenericSymbol, error)

	// 591:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)
}

type GenericParameterDefsReducer interface {
	// 594:2: generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToGenericParameterDefs(GenericParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 595:2: generic_parameter_defs -> add: ...
	AddToGenericParameterDefs(GenericParameterDefs_ GenericSymbol, Comma_ TokenValue, GenericParameterDef_ GenericSymbol) (GenericSymbol, error)
}

type OptionalGenericParameterDefsReducer interface {
	// 598:2: optional_generic_parameter_defs -> generic_parameter_defs: ...
	GenericParameterDefsToOptionalGenericParameterDefs(GenericParameterDefs_ GenericSymbol) (GenericSymbol, error)

	// 599:2: optional_generic_parameter_defs -> nil: ...
	NilToOptionalGenericParameterDefs() (GenericSymbol, error)
}

type OptionalGenericParametersReducer interface {
	// 602:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ TokenValue, OptionalGenericParameterDefs_ GenericSymbol, Rbracket_ TokenValue) (GenericSymbol, error)

	// 603:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (GenericSymbol, error)
}

type FieldDefReducer interface {
	// 610:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 611:2: field_def -> implicit: ...
	ImplicitToFieldDef(ValueType_ TypeExpression) (GenericSymbol, error)

	// 612:2: field_def -> unsafe_statement: ...
	UnsafeStatementToFieldDef(UnsafeStatement_ Statement) (GenericSymbol, error)
}

type ImplicitFieldDefsReducer interface {
	// 615:2: implicit_field_defs -> field_def: ...
	FieldDefToImplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 616:2: implicit_field_defs -> add: ...
	AddToImplicitFieldDefs(ImplicitFieldDefs_ GenericSymbol, Comma_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type OptionalImplicitFieldDefsReducer interface {
	// 619:2: optional_implicit_field_defs -> implicit_field_defs: ...
	ImplicitFieldDefsToOptionalImplicitFieldDefs(ImplicitFieldDefs_ GenericSymbol) (GenericSymbol, error)

	// 620:2: optional_implicit_field_defs -> nil: ...
	NilToOptionalImplicitFieldDefs() (GenericSymbol, error)
}

type ImplicitStructDefReducer interface {
	// 623:2: implicit_struct_def -> ...
	ToImplicitStructDef(Lparen_ TokenValue, OptionalImplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ExplicitFieldDefsReducer interface {
	// 626:2: explicit_field_defs -> field_def: ...
	FieldDefToExplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 627:2: explicit_field_defs -> implicit: ...
	ImplicitToExplicitFieldDefs(ExplicitFieldDefs_ GenericSymbol, Newlines_ TokenCount, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 628:2: explicit_field_defs -> explicit: ...
	ExplicitToExplicitFieldDefs(ExplicitFieldDefs_ GenericSymbol, Comma_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type OptionalExplicitFieldDefsReducer interface {
	// 631:2: optional_explicit_field_defs -> explicit_field_defs: ...
	ExplicitFieldDefsToOptionalExplicitFieldDefs(ExplicitFieldDefs_ GenericSymbol) (GenericSymbol, error)

	// 632:2: optional_explicit_field_defs -> nil: ...
	NilToOptionalExplicitFieldDefs() (GenericSymbol, error)
}

type ExplicitStructDefReducer interface {
	// 635:2: explicit_struct_def -> ...
	ToExplicitStructDef(Struct_ TokenValue, Lparen_ TokenValue, OptionalExplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type EnumValueDefReducer interface {
	// 643:2: enum_value_def -> field_def: ...
	FieldDefToEnumValueDef(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 644:2: enum_value_def -> default: ...
	DefaultToEnumValueDef(FieldDef_ GenericSymbol, Assign_ TokenValue, Default_ TokenValue) (GenericSymbol, error)
}

type ImplicitEnumValueDefsReducer interface {
	// 656:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Or_ TokenValue, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 657:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Or_ TokenValue, EnumValueDef_ GenericSymbol) (GenericSymbol, error)
}

type ImplicitEnumDefReducer interface {
	// 659:37: implicit_enum_def -> ...
	ToImplicitEnumDef(Lparen_ TokenValue, ImplicitEnumValueDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ExplicitEnumValueDefsReducer interface {
	// 662:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Or_ TokenValue, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 663:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Newlines_ TokenCount, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 664:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Or_ TokenValue, EnumValueDef_ GenericSymbol) (GenericSymbol, error)

	// 665:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Newlines_ TokenCount, EnumValueDef_ GenericSymbol) (GenericSymbol, error)
}

type ExplicitEnumDefReducer interface {
	// 667:37: explicit_enum_def -> ...
	ToExplicitEnumDef(Enum_ TokenValue, Lparen_ TokenValue, ExplicitEnumValueDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type TraitPropertyReducer interface {
	// 674:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 675:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ GenericSymbol) (GenericSymbol, error)
}

type TraitPropertiesReducer interface {
	// 678:2: trait_properties -> trait_property: ...
	TraitPropertyToTraitProperties(TraitProperty_ GenericSymbol) (GenericSymbol, error)

	// 679:2: trait_properties -> implicit: ...
	ImplicitToTraitProperties(TraitProperties_ GenericSymbol, Newlines_ TokenCount, TraitProperty_ GenericSymbol) (GenericSymbol, error)

	// 680:2: trait_properties -> explicit: ...
	ExplicitToTraitProperties(TraitProperties_ GenericSymbol, Comma_ TokenValue, TraitProperty_ GenericSymbol) (GenericSymbol, error)
}

type OptionalTraitPropertiesReducer interface {
	// 683:2: optional_trait_properties -> trait_properties: ...
	TraitPropertiesToOptionalTraitProperties(TraitProperties_ GenericSymbol) (GenericSymbol, error)

	// 684:2: optional_trait_properties -> nil: ...
	NilToOptionalTraitProperties() (GenericSymbol, error)
}

type TraitDefReducer interface {
	// 686:29: trait_def -> ...
	ToTraitDef(Trait_ TokenValue, Lparen_ TokenValue, OptionalTraitProperties_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ReturnTypeReducer interface {
	// 694:2: return_type -> returnable_type: ...
	ReturnableTypeToReturnType(ReturnableType_ TypeExpression) (TypeExpression, error)

	// 695:2: return_type -> nil: ...
	NilToReturnType() (TypeExpression, error)
}

type ParameterDeclReducer interface {
	// 698:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 699:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ TokenValue, DotDotDot_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 700:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ TypeExpression) (GenericSymbol, error)

	// 701:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(DotDotDot_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)
}

type ParameterDeclsReducer interface {
	// 704:2: parameter_decls -> parameter_decl: ...
	ParameterDeclToParameterDecls(ParameterDecl_ GenericSymbol) (GenericSymbol, error)

	// 705:2: parameter_decls -> add: ...
	AddToParameterDecls(ParameterDecls_ GenericSymbol, Comma_ TokenValue, ParameterDecl_ GenericSymbol) (GenericSymbol, error)
}

type OptionalParameterDeclsReducer interface {
	// 708:2: optional_parameter_decls -> parameter_decls: ...
	ParameterDeclsToOptionalParameterDecls(ParameterDecls_ GenericSymbol) (GenericSymbol, error)

	// 709:2: optional_parameter_decls -> nil: ...
	NilToOptionalParameterDecls() (GenericSymbol, error)
}

type FuncTypeReducer interface {
	// 712:2: func_type -> ...
	ToFuncType(Func_ TokenValue, Lparen_ TokenValue, OptionalParameterDecls_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression) (TypeExpression, error)
}

type MethodSignatureReducer interface {
	// 723:20: method_signature -> ...
	ToMethodSignature(Func_ TokenValue, Identifier_ TokenValue, Lparen_ TokenValue, OptionalParameterDecls_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression) (GenericSymbol, error)
}

type ParameterDefReducer interface {
	// 729:2: parameter_def -> inferred_ref_arg: ...
	InferredRefArgToParameterDef(Identifier_ TokenValue) (GenericSymbol, error)

	// 730:2: parameter_def -> inferred_ref_vararg: ...
	InferredRefVarargToParameterDef(Identifier_ TokenValue, DotDotDot_ TokenValue) (GenericSymbol, error)

	// 731:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 732:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ TokenValue, DotDotDot_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)
}

type ParameterDefsReducer interface {
	// 735:2: parameter_defs -> parameter_def: ...
	ParameterDefToParameterDefs(ParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 736:2: parameter_defs -> add: ...
	AddToParameterDefs(ParameterDefs_ GenericSymbol, Comma_ TokenValue, ParameterDef_ GenericSymbol) (GenericSymbol, error)
}

type OptionalParameterDefsReducer interface {
	// 739:2: optional_parameter_defs -> parameter_defs: ...
	ParameterDefsToOptionalParameterDefs(ParameterDefs_ GenericSymbol) (GenericSymbol, error)

	// 740:2: optional_parameter_defs -> nil: ...
	NilToOptionalParameterDefs() (GenericSymbol, error)
}

type NamedFuncDefReducer interface {
	// 743:2: named_func_def -> func_def: ...
	FuncDefToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, Lparen_ TokenValue, OptionalParameterDefs_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 744:2: named_func_def -> method_def: ...
	MethodDefToNamedFuncDef(Func_ TokenValue, Lparen_ TokenValue, ParameterDef_ GenericSymbol, Rparen_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, Lparen_2 TokenValue, OptionalParameterDefs_ GenericSymbol, Rparen_2 TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 745:2: named_func_def -> alias: ...
	AliasToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, Expression_ Expression) (SourceDefinition, error)
}

type AnonymousFuncExprReducer interface {
	// 749:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ TokenValue, Lparen_ TokenValue, OptionalParameterDefs_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (Expression, error)
}

type PackageDefReducer interface {
	// 761:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ TokenValue) (SourceDefinition, error)

	// 762:2: package_def -> with_spec: ...
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
	SimpleStatementBodyReducer
	StatementBodyReducer
	OptionalSimpleStatementBodyReducer
	UnsafeStatementReducer
	OptionalJumpLabelReducer
	ExpressionsReducer
	OptionalExpressionsReducer
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
	AtomExprReducer
	LiteralExprReducer
	ImplicitStructExprReducer
	AccessExprReducer
	PostfixUnaryExprReducer
	PrefixUnaryExprReducer
	MulExprReducer
	AddExprReducer
	CmpExprReducer
	AndExprReducer
	OrExprReducer
	InitializableTypeReducer
	AtomTypeReducer
	ReturnableTypeReducer
	ValueTypeReducer
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
	case UnaryOpAssignType:
		return "unary_op_assign"
	case BinaryOpAssignType:
		return "binary_op_assign"
	case UnsafeStatementType:
		return "unsafe_statement"
	case JumpTypeType:
		return "jump_type"
	case OptionalJumpLabelType:
		return "optional_jump_label"
	case ExpressionsType:
		return "expressions"
	case OptionalExpressionsType:
		return "optional_expressions"
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
	case LiteralExprType:
		return "literal_expr"
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

	SourceType                       = SymbolId(343)
	OptionalDefinitionsType          = SymbolId(344)
	OptionalNewlinesType             = SymbolId(345)
	DefinitionsType                  = SymbolId(346)
	DefinitionType                   = SymbolId(347)
	StatementBlockType               = SymbolId(348)
	StatementsType                   = SymbolId(349)
	StatementType                    = SymbolId(350)
	SimpleStatementBodyType          = SymbolId(351)
	StatementBodyType                = SymbolId(352)
	OptionalSimpleStatementBodyType  = SymbolId(353)
	UnaryOpAssignType                = SymbolId(354)
	BinaryOpAssignType               = SymbolId(355)
	UnsafeStatementType              = SymbolId(356)
	JumpTypeType                     = SymbolId(357)
	OptionalJumpLabelType            = SymbolId(358)
	ExpressionsType                  = SymbolId(359)
	OptionalExpressionsType          = SymbolId(360)
	ImportStatementType              = SymbolId(361)
	ImportClauseType                 = SymbolId(362)
	ImportClauseTerminalType         = SymbolId(363)
	ImportClausesType                = SymbolId(364)
	CasePatternsType                 = SymbolId(365)
	VarDeclPatternType               = SymbolId(366)
	VarOrLetType                     = SymbolId(367)
	VarPatternType                   = SymbolId(368)
	TuplePatternType                 = SymbolId(369)
	FieldVarPatternsType             = SymbolId(370)
	FieldVarPatternType              = SymbolId(371)
	OptionalValueTypeType            = SymbolId(372)
	AssignPatternType                = SymbolId(373)
	CasePatternType                  = SymbolId(374)
	ExpressionType                   = SymbolId(375)
	OptionalLabelDeclType            = SymbolId(376)
	SequenceExprType                 = SymbolId(377)
	IfExprType                       = SymbolId(378)
	ConditionType                    = SymbolId(379)
	SwitchExprType                   = SymbolId(380)
	LoopExprType                     = SymbolId(381)
	OptionalSequenceExprType         = SymbolId(382)
	ForAssignmentType                = SymbolId(383)
	CallExprType                     = SymbolId(384)
	OptionalGenericBindingType       = SymbolId(385)
	OptionalGenericArgumentsType     = SymbolId(386)
	GenericArgumentsType             = SymbolId(387)
	OptionalArgumentsType            = SymbolId(388)
	ArgumentsType                    = SymbolId(389)
	ArgumentType                     = SymbolId(390)
	ColonExpressionsType             = SymbolId(391)
	OptionalExpressionType           = SymbolId(392)
	AtomExprType                     = SymbolId(393)
	LiteralExprType                  = SymbolId(394)
	ImplicitStructExprType           = SymbolId(395)
	AccessExprType                   = SymbolId(396)
	PostfixUnaryExprType             = SymbolId(397)
	PrefixUnaryOpType                = SymbolId(398)
	PrefixUnaryExprType              = SymbolId(399)
	MulOpType                        = SymbolId(400)
	MulExprType                      = SymbolId(401)
	AddOpType                        = SymbolId(402)
	AddExprType                      = SymbolId(403)
	CmpOpType                        = SymbolId(404)
	CmpExprType                      = SymbolId(405)
	AndExprType                      = SymbolId(406)
	OrExprType                       = SymbolId(407)
	InitializableTypeType            = SymbolId(408)
	AtomTypeType                     = SymbolId(409)
	ReturnableTypeType               = SymbolId(410)
	ValueTypeType                    = SymbolId(411)
	TypeDefType                      = SymbolId(412)
	GenericParameterDefType          = SymbolId(413)
	GenericParameterDefsType         = SymbolId(414)
	OptionalGenericParameterDefsType = SymbolId(415)
	OptionalGenericParametersType    = SymbolId(416)
	FieldDefType                     = SymbolId(417)
	ImplicitFieldDefsType            = SymbolId(418)
	OptionalImplicitFieldDefsType    = SymbolId(419)
	ImplicitStructDefType            = SymbolId(420)
	ExplicitFieldDefsType            = SymbolId(421)
	OptionalExplicitFieldDefsType    = SymbolId(422)
	ExplicitStructDefType            = SymbolId(423)
	EnumValueDefType                 = SymbolId(424)
	ImplicitEnumValueDefsType        = SymbolId(425)
	ImplicitEnumDefType              = SymbolId(426)
	ExplicitEnumValueDefsType        = SymbolId(427)
	ExplicitEnumDefType              = SymbolId(428)
	TraitPropertyType                = SymbolId(429)
	TraitPropertiesType              = SymbolId(430)
	OptionalTraitPropertiesType      = SymbolId(431)
	TraitDefType                     = SymbolId(432)
	ReturnTypeType                   = SymbolId(433)
	ParameterDeclType                = SymbolId(434)
	ParameterDeclsType               = SymbolId(435)
	OptionalParameterDeclsType       = SymbolId(436)
	FuncTypeType                     = SymbolId(437)
	MethodSignatureType              = SymbolId(438)
	ParameterDefType                 = SymbolId(439)
	ParameterDefsType                = SymbolId(440)
	OptionalParameterDefsType        = SymbolId(441)
	NamedFuncDefType                 = SymbolId(442)
	AnonymousFuncExprType            = SymbolId(443)
	PackageDefType                   = SymbolId(444)
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
	_ReduceDefinitionToDefinitions                            = _ReduceType(7)
	_ReduceAddToDefinitions                                   = _ReduceType(8)
	_ReducePackageDefToDefinition                             = _ReduceType(9)
	_ReduceTypeDefToDefinition                                = _ReduceType(10)
	_ReduceNamedFuncDefToDefinition                           = _ReduceType(11)
	_ReduceGlobalVarDefToDefinition                           = _ReduceType(12)
	_ReduceGlobalVarAssignmentToDefinition                    = _ReduceType(13)
	_ReduceStatementBlockToDefinition                         = _ReduceType(14)
	_ReduceCommentGroupsToDefinition                          = _ReduceType(15)
	_ReduceToStatementBlock                                   = _ReduceType(16)
	_ReduceEmptyListToStatements                              = _ReduceType(17)
	_ReduceAddToStatements                                    = _ReduceType(18)
	_ReduceImplicitToStatement                                = _ReduceType(19)
	_ReduceExplicitToStatement                                = _ReduceType(20)
	_ReduceUnsafeStatementToSimpleStatementBody               = _ReduceType(21)
	_ReduceExpressionOrImplicitStructToSimpleStatementBody    = _ReduceType(22)
	_ReduceAsyncToSimpleStatementBody                         = _ReduceType(23)
	_ReduceDeferToSimpleStatementBody                         = _ReduceType(24)
	_ReduceJumpStatementToSimpleStatementBody                 = _ReduceType(25)
	_ReduceAssignStatementToSimpleStatementBody               = _ReduceType(26)
	_ReduceUnaryOpAssignStatementToSimpleStatementBody        = _ReduceType(27)
	_ReduceBinaryOpAssignStatementToSimpleStatementBody       = _ReduceType(28)
	_ReduceFallthroughToSimpleStatementBody                   = _ReduceType(29)
	_ReduceSimpleStatementBodyToStatementBody                 = _ReduceType(30)
	_ReduceImportStatementToStatementBody                     = _ReduceType(31)
	_ReduceCaseBranchStatementToStatementBody                 = _ReduceType(32)
	_ReduceDefaultBranchStatementToStatementBody              = _ReduceType(33)
	_ReduceSimpleStatementBodyToOptionalSimpleStatementBody   = _ReduceType(34)
	_ReduceNilToOptionalSimpleStatementBody                   = _ReduceType(35)
	_ReduceAddOneAssignToUnaryOpAssign                        = _ReduceType(36)
	_ReduceSubOneAssignToUnaryOpAssign                        = _ReduceType(37)
	_ReduceAddAssignToBinaryOpAssign                          = _ReduceType(38)
	_ReduceSubAssignToBinaryOpAssign                          = _ReduceType(39)
	_ReduceMulAssignToBinaryOpAssign                          = _ReduceType(40)
	_ReduceDivAssignToBinaryOpAssign                          = _ReduceType(41)
	_ReduceModAssignToBinaryOpAssign                          = _ReduceType(42)
	_ReduceBitNegAssignToBinaryOpAssign                       = _ReduceType(43)
	_ReduceBitAndAssignToBinaryOpAssign                       = _ReduceType(44)
	_ReduceBitOrAssignToBinaryOpAssign                        = _ReduceType(45)
	_ReduceBitXorAssignToBinaryOpAssign                       = _ReduceType(46)
	_ReduceBitLshiftAssignToBinaryOpAssign                    = _ReduceType(47)
	_ReduceBitRshiftAssignToBinaryOpAssign                    = _ReduceType(48)
	_ReduceToUnsafeStatement                                  = _ReduceType(49)
	_ReduceReturnToJumpType                                   = _ReduceType(50)
	_ReduceBreakToJumpType                                    = _ReduceType(51)
	_ReduceContinueToJumpType                                 = _ReduceType(52)
	_ReduceJumpLabelToOptionalJumpLabel                       = _ReduceType(53)
	_ReduceUnlabelledToOptionalJumpLabel                      = _ReduceType(54)
	_ReduceExpressionToExpressions                            = _ReduceType(55)
	_ReduceAddToExpressions                                   = _ReduceType(56)
	_ReduceExpressionsToOptionalExpressions                   = _ReduceType(57)
	_ReduceNilToOptionalExpressions                           = _ReduceType(58)
	_ReduceSingleToImportStatement                            = _ReduceType(59)
	_ReduceMultipleToImportStatement                          = _ReduceType(60)
	_ReduceStringLiteralToImportClause                        = _ReduceType(61)
	_ReduceAliasToImportClause                                = _ReduceType(62)
	_ReduceImplicitToImportClauseTerminal                     = _ReduceType(63)
	_ReduceExplicitToImportClauseTerminal                     = _ReduceType(64)
	_ReduceFirstToImportClauses                               = _ReduceType(65)
	_ReduceAddToImportClauses                                 = _ReduceType(66)
	_ReduceCasePatternToCasePatterns                          = _ReduceType(67)
	_ReduceMultipleToCasePatterns                             = _ReduceType(68)
	_ReduceToVarDeclPattern                                   = _ReduceType(69)
	_ReduceVarToVarOrLet                                      = _ReduceType(70)
	_ReduceLetToVarOrLet                                      = _ReduceType(71)
	_ReduceIdentifierToVarPattern                             = _ReduceType(72)
	_ReduceTuplePatternToVarPattern                           = _ReduceType(73)
	_ReduceToTuplePattern                                     = _ReduceType(74)
	_ReduceFieldVarPatternToFieldVarPatterns                  = _ReduceType(75)
	_ReduceAddToFieldVarPatterns                              = _ReduceType(76)
	_ReducePositionalToFieldVarPattern                        = _ReduceType(77)
	_ReduceNamedToFieldVarPattern                             = _ReduceType(78)
	_ReduceDotDotDotToFieldVarPattern                         = _ReduceType(79)
	_ReduceValueTypeToOptionalValueType                       = _ReduceType(80)
	_ReduceNilToOptionalValueType                             = _ReduceType(81)
	_ReduceToAssignPattern                                    = _ReduceType(82)
	_ReduceAssignPatternToCasePattern                         = _ReduceType(83)
	_ReduceEnumMatchPatternToCasePattern                      = _ReduceType(84)
	_ReduceEnumVarDeclPatternToCasePattern                    = _ReduceType(85)
	_ReduceIfExprToExpression                                 = _ReduceType(86)
	_ReduceSwitchExprToExpression                             = _ReduceType(87)
	_ReduceLoopExprToExpression                               = _ReduceType(88)
	_ReduceSequenceExprToExpression                           = _ReduceType(89)
	_ReduceLabelDeclToOptionalLabelDecl                       = _ReduceType(90)
	_ReduceUnlabelledToOptionalLabelDecl                      = _ReduceType(91)
	_ReduceOrExprToSequenceExpr                               = _ReduceType(92)
	_ReduceVarDeclPatternToSequenceExpr                       = _ReduceType(93)
	_ReduceAssignVarPatternToSequenceExpr                     = _ReduceType(94)
	_ReduceNoElseToIfExpr                                     = _ReduceType(95)
	_ReduceIfElseToIfExpr                                     = _ReduceType(96)
	_ReduceMultiIfElseToIfExpr                                = _ReduceType(97)
	_ReduceSequenceExprToCondition                            = _ReduceType(98)
	_ReduceCaseToCondition                                    = _ReduceType(99)
	_ReduceToSwitchExpr                                       = _ReduceType(100)
	_ReduceInfiniteToLoopExpr                                 = _ReduceType(101)
	_ReduceDoWhileToLoopExpr                                  = _ReduceType(102)
	_ReduceWhileToLoopExpr                                    = _ReduceType(103)
	_ReduceIteratorToLoopExpr                                 = _ReduceType(104)
	_ReduceForToLoopExpr                                      = _ReduceType(105)
	_ReduceSequenceExprToOptionalSequenceExpr                 = _ReduceType(106)
	_ReduceNilToOptionalSequenceExpr                          = _ReduceType(107)
	_ReduceSequenceExprToForAssignment                        = _ReduceType(108)
	_ReduceAssignToForAssignment                              = _ReduceType(109)
	_ReduceToCallExpr                                         = _ReduceType(110)
	_ReduceBindingToOptionalGenericBinding                    = _ReduceType(111)
	_ReduceNilToOptionalGenericBinding                        = _ReduceType(112)
	_ReduceGenericArgumentsToOptionalGenericArguments         = _ReduceType(113)
	_ReduceNilToOptionalGenericArguments                      = _ReduceType(114)
	_ReduceValueTypeToGenericArguments                        = _ReduceType(115)
	_ReduceAddToGenericArguments                              = _ReduceType(116)
	_ReduceArgumentsToOptionalArguments                       = _ReduceType(117)
	_ReduceNilToOptionalArguments                             = _ReduceType(118)
	_ReduceArgumentToArguments                                = _ReduceType(119)
	_ReduceAddToArguments                                     = _ReduceType(120)
	_ReducePositionalToArgument                               = _ReduceType(121)
	_ReduceNamedToArgument                                    = _ReduceType(122)
	_ReduceColonExpressionsToArgument                         = _ReduceType(123)
	_ReduceDotDotDotToArgument                                = _ReduceType(124)
	_ReducePairToColonExpressions                             = _ReduceType(125)
	_ReduceAddToColonExpressions                              = _ReduceType(126)
	_ReduceExpressionToOptionalExpression                     = _ReduceType(127)
	_ReduceNilToOptionalExpression                            = _ReduceType(128)
	_ReduceLiteralExprToAtomExpr                              = _ReduceType(129)
	_ReduceIdentifierToAtomExpr                               = _ReduceType(130)
	_ReduceBlockExprToAtomExpr                                = _ReduceType(131)
	_ReduceAnonymousFuncExprToAtomExpr                        = _ReduceType(132)
	_ReduceInitializeExprToAtomExpr                           = _ReduceType(133)
	_ReduceImplicitStructExprToAtomExpr                       = _ReduceType(134)
	_ReduceParseErrorToAtomExpr                               = _ReduceType(135)
	_ReduceTrueToLiteralExpr                                  = _ReduceType(136)
	_ReduceFalseToLiteralExpr                                 = _ReduceType(137)
	_ReduceIntegerLiteralToLiteralExpr                        = _ReduceType(138)
	_ReduceFloatLiteralToLiteralExpr                          = _ReduceType(139)
	_ReduceRuneLiteralToLiteralExpr                           = _ReduceType(140)
	_ReduceStringLiteralToLiteralExpr                         = _ReduceType(141)
	_ReduceToImplicitStructExpr                               = _ReduceType(142)
	_ReduceAtomExprToAccessExpr                               = _ReduceType(143)
	_ReduceAccessToAccessExpr                                 = _ReduceType(144)
	_ReduceCallExprToAccessExpr                               = _ReduceType(145)
	_ReduceIndexToAccessExpr                                  = _ReduceType(146)
	_ReduceAccessExprToPostfixUnaryExpr                       = _ReduceType(147)
	_ReduceToPostfixUnaryExpr                                 = _ReduceType(148)
	_ReduceNotToPrefixUnaryOp                                 = _ReduceType(149)
	_ReduceBitNegToPrefixUnaryOp                              = _ReduceType(150)
	_ReduceSubToPrefixUnaryOp                                 = _ReduceType(151)
	_ReduceMulToPrefixUnaryOp                                 = _ReduceType(152)
	_ReduceBitAndToPrefixUnaryOp                              = _ReduceType(153)
	_ReducePostfixUnaryExprToPrefixUnaryExpr                  = _ReduceType(154)
	_ReduceToPrefixUnaryExpr                                  = _ReduceType(155)
	_ReduceMulToMulOp                                         = _ReduceType(156)
	_ReduceDivToMulOp                                         = _ReduceType(157)
	_ReduceModToMulOp                                         = _ReduceType(158)
	_ReduceBitAndToMulOp                                      = _ReduceType(159)
	_ReduceBitLshiftToMulOp                                   = _ReduceType(160)
	_ReduceBitRshiftToMulOp                                   = _ReduceType(161)
	_ReducePrefixUnaryExprToMulExpr                           = _ReduceType(162)
	_ReduceToMulExpr                                          = _ReduceType(163)
	_ReduceAddToAddOp                                         = _ReduceType(164)
	_ReduceSubToAddOp                                         = _ReduceType(165)
	_ReduceBitOrToAddOp                                       = _ReduceType(166)
	_ReduceBitXorToAddOp                                      = _ReduceType(167)
	_ReduceMulExprToAddExpr                                   = _ReduceType(168)
	_ReduceToAddExpr                                          = _ReduceType(169)
	_ReduceEqualToCmpOp                                       = _ReduceType(170)
	_ReduceNotEqualToCmpOp                                    = _ReduceType(171)
	_ReduceLessToCmpOp                                        = _ReduceType(172)
	_ReduceLessOrEqualToCmpOp                                 = _ReduceType(173)
	_ReduceGreaterToCmpOp                                     = _ReduceType(174)
	_ReduceGreaterOrEqualToCmpOp                              = _ReduceType(175)
	_ReduceAddExprToCmpExpr                                   = _ReduceType(176)
	_ReduceToCmpExpr                                          = _ReduceType(177)
	_ReduceCmpExprToAndExpr                                   = _ReduceType(178)
	_ReduceToAndExpr                                          = _ReduceType(179)
	_ReduceAndExprToOrExpr                                    = _ReduceType(180)
	_ReduceToOrExpr                                           = _ReduceType(181)
	_ReduceExplicitStructDefToInitializableType               = _ReduceType(182)
	_ReduceSliceToInitializableType                           = _ReduceType(183)
	_ReduceArrayToInitializableType                           = _ReduceType(184)
	_ReduceMapToInitializableType                             = _ReduceType(185)
	_ReduceInitializableTypeToAtomType                        = _ReduceType(186)
	_ReduceNamedToAtomType                                    = _ReduceType(187)
	_ReduceExternNamedToAtomType                              = _ReduceType(188)
	_ReduceInferredToAtomType                                 = _ReduceType(189)
	_ReduceImplicitStructDefToAtomType                        = _ReduceType(190)
	_ReduceExplicitEnumDefToAtomType                          = _ReduceType(191)
	_ReduceImplicitEnumDefToAtomType                          = _ReduceType(192)
	_ReduceTraitDefToAtomType                                 = _ReduceType(193)
	_ReduceFuncTypeToAtomType                                 = _ReduceType(194)
	_ReduceParseErrorToAtomType                               = _ReduceType(195)
	_ReduceAtomTypeToReturnableType                           = _ReduceType(196)
	_ReduceOptionalToReturnableType                           = _ReduceType(197)
	_ReduceResultToReturnableType                             = _ReduceType(198)
	_ReduceReferenceToReturnableType                          = _ReduceType(199)
	_ReducePublicMethodsTraitToReturnableType                 = _ReduceType(200)
	_ReducePublicTraitToReturnableType                        = _ReduceType(201)
	_ReduceReturnableTypeToValueType                          = _ReduceType(202)
	_ReduceTraitIntersectToValueType                          = _ReduceType(203)
	_ReduceTraitUnionToValueType                              = _ReduceType(204)
	_ReduceTraitDifferenceToValueType                         = _ReduceType(205)
	_ReduceDefinitionToTypeDef                                = _ReduceType(206)
	_ReduceConstrainedDefToTypeDef                            = _ReduceType(207)
	_ReduceAliasToTypeDef                                     = _ReduceType(208)
	_ReduceUnconstrainedToGenericParameterDef                 = _ReduceType(209)
	_ReduceConstrainedToGenericParameterDef                   = _ReduceType(210)
	_ReduceGenericParameterDefToGenericParameterDefs          = _ReduceType(211)
	_ReduceAddToGenericParameterDefs                          = _ReduceType(212)
	_ReduceGenericParameterDefsToOptionalGenericParameterDefs = _ReduceType(213)
	_ReduceNilToOptionalGenericParameterDefs                  = _ReduceType(214)
	_ReduceGenericToOptionalGenericParameters                 = _ReduceType(215)
	_ReduceNilToOptionalGenericParameters                     = _ReduceType(216)
	_ReduceExplicitToFieldDef                                 = _ReduceType(217)
	_ReduceImplicitToFieldDef                                 = _ReduceType(218)
	_ReduceUnsafeStatementToFieldDef                          = _ReduceType(219)
	_ReduceFieldDefToImplicitFieldDefs                        = _ReduceType(220)
	_ReduceAddToImplicitFieldDefs                             = _ReduceType(221)
	_ReduceImplicitFieldDefsToOptionalImplicitFieldDefs       = _ReduceType(222)
	_ReduceNilToOptionalImplicitFieldDefs                     = _ReduceType(223)
	_ReduceToImplicitStructDef                                = _ReduceType(224)
	_ReduceFieldDefToExplicitFieldDefs                        = _ReduceType(225)
	_ReduceImplicitToExplicitFieldDefs                        = _ReduceType(226)
	_ReduceExplicitToExplicitFieldDefs                        = _ReduceType(227)
	_ReduceExplicitFieldDefsToOptionalExplicitFieldDefs       = _ReduceType(228)
	_ReduceNilToOptionalExplicitFieldDefs                     = _ReduceType(229)
	_ReduceToExplicitStructDef                                = _ReduceType(230)
	_ReduceFieldDefToEnumValueDef                             = _ReduceType(231)
	_ReduceDefaultToEnumValueDef                              = _ReduceType(232)
	_ReducePairToImplicitEnumValueDefs                        = _ReduceType(233)
	_ReduceAddToImplicitEnumValueDefs                         = _ReduceType(234)
	_ReduceToImplicitEnumDef                                  = _ReduceType(235)
	_ReduceExplicitPairToExplicitEnumValueDefs                = _ReduceType(236)
	_ReduceImplicitPairToExplicitEnumValueDefs                = _ReduceType(237)
	_ReduceExplicitAddToExplicitEnumValueDefs                 = _ReduceType(238)
	_ReduceImplicitAddToExplicitEnumValueDefs                 = _ReduceType(239)
	_ReduceToExplicitEnumDef                                  = _ReduceType(240)
	_ReduceFieldDefToTraitProperty                            = _ReduceType(241)
	_ReduceMethodSignatureToTraitProperty                     = _ReduceType(242)
	_ReduceTraitPropertyToTraitProperties                     = _ReduceType(243)
	_ReduceImplicitToTraitProperties                          = _ReduceType(244)
	_ReduceExplicitToTraitProperties                          = _ReduceType(245)
	_ReduceTraitPropertiesToOptionalTraitProperties           = _ReduceType(246)
	_ReduceNilToOptionalTraitProperties                       = _ReduceType(247)
	_ReduceToTraitDef                                         = _ReduceType(248)
	_ReduceReturnableTypeToReturnType                         = _ReduceType(249)
	_ReduceNilToReturnType                                    = _ReduceType(250)
	_ReduceArgToParameterDecl                                 = _ReduceType(251)
	_ReduceVarargToParameterDecl                              = _ReduceType(252)
	_ReduceUnamedToParameterDecl                              = _ReduceType(253)
	_ReduceUnnamedVarargToParameterDecl                       = _ReduceType(254)
	_ReduceParameterDeclToParameterDecls                      = _ReduceType(255)
	_ReduceAddToParameterDecls                                = _ReduceType(256)
	_ReduceParameterDeclsToOptionalParameterDecls             = _ReduceType(257)
	_ReduceNilToOptionalParameterDecls                        = _ReduceType(258)
	_ReduceToFuncType                                         = _ReduceType(259)
	_ReduceToMethodSignature                                  = _ReduceType(260)
	_ReduceInferredRefArgToParameterDef                       = _ReduceType(261)
	_ReduceInferredRefVarargToParameterDef                    = _ReduceType(262)
	_ReduceArgToParameterDef                                  = _ReduceType(263)
	_ReduceVarargToParameterDef                               = _ReduceType(264)
	_ReduceParameterDefToParameterDefs                        = _ReduceType(265)
	_ReduceAddToParameterDefs                                 = _ReduceType(266)
	_ReduceParameterDefsToOptionalParameterDefs               = _ReduceType(267)
	_ReduceNilToOptionalParameterDefs                         = _ReduceType(268)
	_ReduceFuncDefToNamedFuncDef                              = _ReduceType(269)
	_ReduceMethodDefToNamedFuncDef                            = _ReduceType(270)
	_ReduceAliasToNamedFuncDef                                = _ReduceType(271)
	_ReduceToAnonymousFuncExpr                                = _ReduceType(272)
	_ReduceNoSpecToPackageDef                                 = _ReduceType(273)
	_ReduceWithSpecToPackageDef                               = _ReduceType(274)
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
	case _ReduceExpressionOrImplicitStructToSimpleStatementBody:
		return "ExpressionOrImplicitStructToSimpleStatementBody"
	case _ReduceAsyncToSimpleStatementBody:
		return "AsyncToSimpleStatementBody"
	case _ReduceDeferToSimpleStatementBody:
		return "DeferToSimpleStatementBody"
	case _ReduceJumpStatementToSimpleStatementBody:
		return "JumpStatementToSimpleStatementBody"
	case _ReduceAssignStatementToSimpleStatementBody:
		return "AssignStatementToSimpleStatementBody"
	case _ReduceUnaryOpAssignStatementToSimpleStatementBody:
		return "UnaryOpAssignStatementToSimpleStatementBody"
	case _ReduceBinaryOpAssignStatementToSimpleStatementBody:
		return "BinaryOpAssignStatementToSimpleStatementBody"
	case _ReduceFallthroughToSimpleStatementBody:
		return "FallthroughToSimpleStatementBody"
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
	case _ReduceLiteralExprToAtomExpr:
		return "LiteralExprToAtomExpr"
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
	case _ReduceParseErrorToAtomExpr:
		return "ParseErrorToAtomExpr"
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
	case _ReduceToPostfixUnaryExpr:
		return "ToPostfixUnaryExpr"
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
	case _ReduceToPrefixUnaryExpr:
		return "ToPrefixUnaryExpr"
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
	case _ReduceToMulExpr:
		return "ToMulExpr"
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
	case _ReduceToAddExpr:
		return "ToAddExpr"
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
	case _ReduceToCmpExpr:
		return "ToCmpExpr"
	case _ReduceCmpExprToAndExpr:
		return "CmpExprToAndExpr"
	case _ReduceToAndExpr:
		return "ToAndExpr"
	case _ReduceAndExprToOrExpr:
		return "AndExprToOrExpr"
	case _ReduceToOrExpr:
		return "ToOrExpr"
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
	case _ReduceParseErrorToAtomType:
		return "ParseErrorToAtomType"
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
	case ExpressionType, SequenceExprType, IfExprType, LoopExprType, OptionalSequenceExprType, CallExprType, OptionalExpressionType, AtomExprType, LiteralExprType, ImplicitStructExprType, AccessExprType, PostfixUnaryExprType, PrefixUnaryExprType, MulExprType, AddExprType, CmpExprType, AndExprType, OrExprType, AnonymousFuncExprType:
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
	case StatementType, SimpleStatementBodyType, StatementBodyType, OptionalSimpleStatementBodyType, UnsafeStatementType, ImportStatementType:
		loc, ok := interface{}(s.Statement).(locator)
		if ok {
			return loc.Loc()
		}
	case StatementsType:
		loc, ok := interface{}(s.Statements).(locator)
		if ok {
			return loc.Loc()
		}
	case OptionalValueTypeType, InitializableTypeType, AtomTypeType, ReturnableTypeType, ValueTypeType, ImplicitStructDefType, ExplicitStructDefType, ImplicitEnumDefType, ExplicitEnumDefType, TraitDefType, ReturnTypeType, FuncTypeType:
		loc, ok := interface{}(s.TypeExpression).(locator)
		if ok {
			return loc.Loc()
		}
	case CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, AsToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, DotDotDotToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, UnaryOpAssignType, BinaryOpAssignType, JumpTypeType, OptionalJumpLabelType, VarOrLetType, OptionalLabelDeclType, PrefixUnaryOpType, MulOpType, AddOpType, CmpOpType:
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
	case ExpressionType, SequenceExprType, IfExprType, LoopExprType, OptionalSequenceExprType, CallExprType, OptionalExpressionType, AtomExprType, LiteralExprType, ImplicitStructExprType, AccessExprType, PostfixUnaryExprType, PrefixUnaryExprType, MulExprType, AddExprType, CmpExprType, AndExprType, OrExprType, AnonymousFuncExprType:
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
	case StatementType, SimpleStatementBodyType, StatementBodyType, OptionalSimpleStatementBodyType, UnsafeStatementType, ImportStatementType:
		loc, ok := interface{}(s.Statement).(locator)
		if ok {
			return loc.End()
		}
	case StatementsType:
		loc, ok := interface{}(s.Statements).(locator)
		if ok {
			return loc.End()
		}
	case OptionalValueTypeType, InitializableTypeType, AtomTypeType, ReturnableTypeType, ValueTypeType, ImplicitStructDefType, ExplicitStructDefType, ImplicitEnumDefType, ExplicitEnumDefType, TraitDefType, ReturnTypeType, FuncTypeType:
		loc, ok := interface{}(s.TypeExpression).(locator)
		if ok {
			return loc.End()
		}
	case CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, AsToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, DotDotDotToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, UnaryOpAssignType, BinaryOpAssignType, JumpTypeType, OptionalJumpLabelType, VarOrLetType, OptionalLabelDeclType, PrefixUnaryOpType, MulOpType, AddOpType, CmpOpType:
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
	case _ReduceExpressionOrImplicitStructToSimpleStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Statement, err = reducer.ExpressionOrImplicitStructToSimpleStatementBody(args[0].Generic_)
	case _ReduceAsyncToSimpleStatementBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Statement, err = reducer.AsyncToSimpleStatementBody(args[0].Value, args[1].Expression)
	case _ReduceDeferToSimpleStatementBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Statement, err = reducer.DeferToSimpleStatementBody(args[0].Value, args[1].Expression)
	case _ReduceJumpStatementToSimpleStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Statement, err = reducer.JumpStatementToSimpleStatementBody(args[0].Value, args[1].Value, args[2].Generic_)
	case _ReduceAssignStatementToSimpleStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Statement, err = reducer.AssignStatementToSimpleStatementBody(args[0].Generic_, args[1].Value, args[2].Expression)
	case _ReduceUnaryOpAssignStatementToSimpleStatementBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Statement, err = reducer.UnaryOpAssignStatementToSimpleStatementBody(args[0].Expression, args[1].Value)
	case _ReduceBinaryOpAssignStatementToSimpleStatementBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Statement, err = reducer.BinaryOpAssignStatementToSimpleStatementBody(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceFallthroughToSimpleStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementBodyType
		symbol.Statement, err = reducer.FallthroughToSimpleStatementBody(args[0].Value)
	case _ReduceSimpleStatementBodyToStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementBodyType
		//line grammar.lr:162:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceImportStatementToStatementBody:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementBodyType
		//line grammar.lr:165:4
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
		//line grammar.lr:174:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceNilToOptionalSimpleStatementBody:
		symbol.SymbolId_ = OptionalSimpleStatementBodyType
		symbol.Statement, err = reducer.NilToOptionalSimpleStatementBody()
	case _ReduceAddOneAssignToUnaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpAssignType
		//line grammar.lr:182:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubOneAssignToUnaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpAssignType
		//line grammar.lr:183:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:186:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:187:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:188:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDivAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:189:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceModAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:190:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:191:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:192:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitOrAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:193:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitXorAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:194:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitLshiftAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:195:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitRshiftAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:196:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceToUnsafeStatement:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = UnsafeStatementType
		symbol.Statement, err = reducer.ToUnsafeStatement(args[0].Value, args[1].Value, args[2].Value, args[3].Value, args[4].Value)
	case _ReduceReturnToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		//line grammar.lr:211:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBreakToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		//line grammar.lr:212:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceContinueToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		//line grammar.lr:213:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceJumpLabelToOptionalJumpLabel:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalJumpLabelType
		symbol.Value, err = reducer.JumpLabelToOptionalJumpLabel(args[0].Value)
	case _ReduceUnlabelledToOptionalJumpLabel:
		symbol.SymbolId_ = OptionalJumpLabelType
		symbol.Value, err = reducer.UnlabelledToOptionalJumpLabel()
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
	case _ReduceExpressionsToOptionalExpressions:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalExpressionsType
		symbol.Generic_, err = reducer.ExpressionsToOptionalExpressions(args[0].Generic_)
	case _ReduceNilToOptionalExpressions:
		symbol.SymbolId_ = OptionalExpressionsType
		symbol.Generic_, err = reducer.NilToOptionalExpressions()
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
		//line grammar.lr:265:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLetToVarOrLet:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarOrLetType
		//line grammar.lr:266:4
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
		//line grammar.lr:284:4
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
		//line grammar.lr:295:4
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
		symbol.Expression, err = reducer.SwitchExprToExpression(args[0].Value, args[1].Generic_)
	case _ReduceLoopExprToExpression:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExpressionType
		symbol.Expression, err = reducer.LoopExprToExpression(args[0].Value, args[1].Expression)
	case _ReduceSequenceExprToExpression:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionType
		//line grammar.lr:313:4
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
		//line grammar.lr:324:4
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
		symbol.Generic_, err = reducer.ToSwitchExpr(args[0].Value, args[1].Expression, args[2].Generic_)
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
	case _ReduceLiteralExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:451:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceIdentifierToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Expression, err = reducer.IdentifierToAtomExpr(args[0].Value)
	case _ReduceBlockExprToAtomExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = AtomExprType
		symbol.Expression, err = reducer.BlockExprToAtomExpr(args[0].Value, args[1].Generic_)
	case _ReduceAnonymousFuncExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:454:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceInitializeExprToAtomExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = AtomExprType
		symbol.Expression, err = reducer.InitializeExprToAtomExpr(args[0].TypeExpression, args[1].Value, args[2].Generic_, args[3].Value)
	case _ReduceImplicitStructExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:456:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceParseErrorToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		symbol.Expression, err = reducer.ParseErrorToAtomExpr(args[0].ParseError)
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
	case _ReduceToImplicitStructExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitStructExprType
		symbol.Expression, err = reducer.ToImplicitStructExpr(args[0].Value, args[1].Generic_, args[2].Value)
	case _ReduceAtomExprToAccessExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessExprType
		//line grammar.lr:470:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceAccessToAccessExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = AccessExprType
		symbol.Expression, err = reducer.AccessToAccessExpr(args[0].Expression, args[1].Value, args[2].Value)
	case _ReduceCallExprToAccessExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessExprType
		//line grammar.lr:472:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceIndexToAccessExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = AccessExprType
		symbol.Expression, err = reducer.IndexToAccessExpr(args[0].Expression, args[1].Value, args[2].Generic_, args[3].Value)
	case _ReduceAccessExprToPostfixUnaryExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixUnaryExprType
		//line grammar.lr:476:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToPostfixUnaryExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PostfixUnaryExprType
		symbol.Expression, err = reducer.ToPostfixUnaryExpr(args[0].Expression, args[1].Value)
	case _ReduceNotToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:480:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:481:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:482:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:485:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:488:4
		symbol.Value = args[0].Value
		err = nil
	case _ReducePostfixUnaryExprToPrefixUnaryExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryExprType
		//line grammar.lr:491:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToPrefixUnaryExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PrefixUnaryExprType
		symbol.Expression, err = reducer.ToPrefixUnaryExpr(args[0].Value, args[1].Expression)
	case _ReduceMulToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:495:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDivToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:496:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceModToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:497:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:498:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitLshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:499:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitRshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:500:4
		symbol.Value = args[0].Value
		err = nil
	case _ReducePrefixUnaryExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		//line grammar.lr:503:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToMulExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = MulExprType
		symbol.Expression, err = reducer.ToMulExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceAddToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:507:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:508:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitOrToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:509:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitXorToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:510:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulExprToAddExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddExprType
		//line grammar.lr:513:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToAddExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = AddExprType
		symbol.Expression, err = reducer.ToAddExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:517:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceNotEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:518:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLessToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:519:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLessOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:520:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceGreaterToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:521:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceGreaterOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:522:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddExprToCmpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpExprType
		//line grammar.lr:525:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToCmpExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CmpExprType
		symbol.Expression, err = reducer.ToCmpExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceCmpExprToAndExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AndExprType
		//line grammar.lr:529:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToAndExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = AndExprType
		symbol.Expression, err = reducer.ToAndExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceAndExprToOrExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OrExprType
		//line grammar.lr:533:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceToOrExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OrExprType
		symbol.Expression, err = reducer.ToOrExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceExplicitStructDefToInitializableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeType
		//line grammar.lr:543:4
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
		//line grammar.lr:549:4
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
		//line grammar.lr:553:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceExplicitEnumDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:554:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceImplicitEnumDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:555:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceTraitDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:556:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceFuncTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:557:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceParseErrorToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		symbol.TypeExpression, err = reducer.ParseErrorToAtomType(args[0].ParseError)
	case _ReduceAtomTypeToReturnableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableTypeType
		//line grammar.lr:564:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceOptionalToReturnableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.TypeExpression, err = reducer.OptionalToReturnableType(args[0].Value, args[1].TypeExpression)
	case _ReduceResultToReturnableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.TypeExpression, err = reducer.ResultToReturnableType(args[0].Value, args[1].TypeExpression)
	case _ReduceReferenceToReturnableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.TypeExpression, err = reducer.ReferenceToReturnableType(args[0].Value, args[1].TypeExpression)
	case _ReducePublicMethodsTraitToReturnableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.TypeExpression, err = reducer.PublicMethodsTraitToReturnableType(args[0].Value, args[1].TypeExpression)
	case _ReducePublicTraitToReturnableType:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ReturnableTypeType
		symbol.TypeExpression, err = reducer.PublicTraitToReturnableType(args[0].Value, args[1].TypeExpression)
	case _ReduceReturnableTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		//line grammar.lr:574:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceTraitIntersectToValueType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ValueTypeType
		symbol.TypeExpression, err = reducer.TraitIntersectToValueType(args[0].TypeExpression, args[1].Value, args[2].TypeExpression)
	case _ReduceTraitUnionToValueType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ValueTypeType
		symbol.TypeExpression, err = reducer.TraitUnionToValueType(args[0].TypeExpression, args[1].Value, args[2].TypeExpression)
	case _ReduceTraitDifferenceToValueType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ValueTypeType
		symbol.TypeExpression, err = reducer.TraitDifferenceToValueType(args[0].TypeExpression, args[1].Value, args[2].TypeExpression)
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
	_GotoState428Action                                             = &_Action{_ShiftAction, _State428, 0}
	_GotoState429Action                                             = &_Action{_ShiftAction, _State429, 0}
	_GotoState430Action                                             = &_Action{_ShiftAction, _State430, 0}
	_GotoState431Action                                             = &_Action{_ShiftAction, _State431, 0}
	_GotoState432Action                                             = &_Action{_ShiftAction, _State432, 0}
	_GotoState433Action                                             = &_Action{_ShiftAction, _State433, 0}
	_GotoState434Action                                             = &_Action{_ShiftAction, _State434, 0}
	_GotoState435Action                                             = &_Action{_ShiftAction, _State435, 0}
	_GotoState436Action                                             = &_Action{_ShiftAction, _State436, 0}
	_GotoState437Action                                             = &_Action{_ShiftAction, _State437, 0}
	_GotoState438Action                                             = &_Action{_ShiftAction, _State438, 0}
	_GotoState439Action                                             = &_Action{_ShiftAction, _State439, 0}
	_GotoState440Action                                             = &_Action{_ShiftAction, _State440, 0}
	_GotoState441Action                                             = &_Action{_ShiftAction, _State441, 0}
	_GotoState442Action                                             = &_Action{_ShiftAction, _State442, 0}
	_GotoState443Action                                             = &_Action{_ShiftAction, _State443, 0}
	_GotoState444Action                                             = &_Action{_ShiftAction, _State444, 0}
	_GotoState445Action                                             = &_Action{_ShiftAction, _State445, 0}
	_GotoState446Action                                             = &_Action{_ShiftAction, _State446, 0}
	_GotoState447Action                                             = &_Action{_ShiftAction, _State447, 0}
	_GotoState448Action                                             = &_Action{_ShiftAction, _State448, 0}
	_GotoState449Action                                             = &_Action{_ShiftAction, _State449, 0}
	_GotoState450Action                                             = &_Action{_ShiftAction, _State450, 0}
	_ReduceToSourceAction                                           = &_Action{_ReduceAction, 0, _ReduceToSource}
	_ReduceNewlinesToOptionalDefinitionsAction                      = &_Action{_ReduceAction, 0, _ReduceNewlinesToOptionalDefinitions}
	_ReduceDefinitionsToOptionalDefinitionsAction                   = &_Action{_ReduceAction, 0, _ReduceDefinitionsToOptionalDefinitions}
	_ReduceNilToOptionalDefinitionsAction                           = &_Action{_ReduceAction, 0, _ReduceNilToOptionalDefinitions}
	_ReduceNewlinesToOptionalNewlinesAction                         = &_Action{_ReduceAction, 0, _ReduceNewlinesToOptionalNewlines}
	_ReduceNilToOptionalNewlinesAction                              = &_Action{_ReduceAction, 0, _ReduceNilToOptionalNewlines}
	_ReduceDefinitionToDefinitionsAction                            = &_Action{_ReduceAction, 0, _ReduceDefinitionToDefinitions}
	_ReduceAddToDefinitionsAction                                   = &_Action{_ReduceAction, 0, _ReduceAddToDefinitions}
	_ReducePackageDefToDefinitionAction                             = &_Action{_ReduceAction, 0, _ReducePackageDefToDefinition}
	_ReduceTypeDefToDefinitionAction                                = &_Action{_ReduceAction, 0, _ReduceTypeDefToDefinition}
	_ReduceNamedFuncDefToDefinitionAction                           = &_Action{_ReduceAction, 0, _ReduceNamedFuncDefToDefinition}
	_ReduceGlobalVarDefToDefinitionAction                           = &_Action{_ReduceAction, 0, _ReduceGlobalVarDefToDefinition}
	_ReduceGlobalVarAssignmentToDefinitionAction                    = &_Action{_ReduceAction, 0, _ReduceGlobalVarAssignmentToDefinition}
	_ReduceStatementBlockToDefinitionAction                         = &_Action{_ReduceAction, 0, _ReduceStatementBlockToDefinition}
	_ReduceCommentGroupsToDefinitionAction                          = &_Action{_ReduceAction, 0, _ReduceCommentGroupsToDefinition}
	_ReduceToStatementBlockAction                                   = &_Action{_ReduceAction, 0, _ReduceToStatementBlock}
	_ReduceEmptyListToStatementsAction                              = &_Action{_ReduceAction, 0, _ReduceEmptyListToStatements}
	_ReduceAddToStatementsAction                                    = &_Action{_ReduceAction, 0, _ReduceAddToStatements}
	_ReduceImplicitToStatementAction                                = &_Action{_ReduceAction, 0, _ReduceImplicitToStatement}
	_ReduceExplicitToStatementAction                                = &_Action{_ReduceAction, 0, _ReduceExplicitToStatement}
	_ReduceUnsafeStatementToSimpleStatementBodyAction               = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToSimpleStatementBody}
	_ReduceExpressionOrImplicitStructToSimpleStatementBodyAction    = &_Action{_ReduceAction, 0, _ReduceExpressionOrImplicitStructToSimpleStatementBody}
	_ReduceAsyncToSimpleStatementBodyAction                         = &_Action{_ReduceAction, 0, _ReduceAsyncToSimpleStatementBody}
	_ReduceDeferToSimpleStatementBodyAction                         = &_Action{_ReduceAction, 0, _ReduceDeferToSimpleStatementBody}
	_ReduceJumpStatementToSimpleStatementBodyAction                 = &_Action{_ReduceAction, 0, _ReduceJumpStatementToSimpleStatementBody}
	_ReduceAssignStatementToSimpleStatementBodyAction               = &_Action{_ReduceAction, 0, _ReduceAssignStatementToSimpleStatementBody}
	_ReduceUnaryOpAssignStatementToSimpleStatementBodyAction        = &_Action{_ReduceAction, 0, _ReduceUnaryOpAssignStatementToSimpleStatementBody}
	_ReduceBinaryOpAssignStatementToSimpleStatementBodyAction       = &_Action{_ReduceAction, 0, _ReduceBinaryOpAssignStatementToSimpleStatementBody}
	_ReduceFallthroughToSimpleStatementBodyAction                   = &_Action{_ReduceAction, 0, _ReduceFallthroughToSimpleStatementBody}
	_ReduceSimpleStatementBodyToStatementBodyAction                 = &_Action{_ReduceAction, 0, _ReduceSimpleStatementBodyToStatementBody}
	_ReduceImportStatementToStatementBodyAction                     = &_Action{_ReduceAction, 0, _ReduceImportStatementToStatementBody}
	_ReduceCaseBranchStatementToStatementBodyAction                 = &_Action{_ReduceAction, 0, _ReduceCaseBranchStatementToStatementBody}
	_ReduceDefaultBranchStatementToStatementBodyAction              = &_Action{_ReduceAction, 0, _ReduceDefaultBranchStatementToStatementBody}
	_ReduceSimpleStatementBodyToOptionalSimpleStatementBodyAction   = &_Action{_ReduceAction, 0, _ReduceSimpleStatementBodyToOptionalSimpleStatementBody}
	_ReduceNilToOptionalSimpleStatementBodyAction                   = &_Action{_ReduceAction, 0, _ReduceNilToOptionalSimpleStatementBody}
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
	_ReduceReturnToJumpTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceReturnToJumpType}
	_ReduceBreakToJumpTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceBreakToJumpType}
	_ReduceContinueToJumpTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceContinueToJumpType}
	_ReduceJumpLabelToOptionalJumpLabelAction                       = &_Action{_ReduceAction, 0, _ReduceJumpLabelToOptionalJumpLabel}
	_ReduceUnlabelledToOptionalJumpLabelAction                      = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalJumpLabel}
	_ReduceExpressionToExpressionsAction                            = &_Action{_ReduceAction, 0, _ReduceExpressionToExpressions}
	_ReduceAddToExpressionsAction                                   = &_Action{_ReduceAction, 0, _ReduceAddToExpressions}
	_ReduceExpressionsToOptionalExpressionsAction                   = &_Action{_ReduceAction, 0, _ReduceExpressionsToOptionalExpressions}
	_ReduceNilToOptionalExpressionsAction                           = &_Action{_ReduceAction, 0, _ReduceNilToOptionalExpressions}
	_ReduceSingleToImportStatementAction                            = &_Action{_ReduceAction, 0, _ReduceSingleToImportStatement}
	_ReduceMultipleToImportStatementAction                          = &_Action{_ReduceAction, 0, _ReduceMultipleToImportStatement}
	_ReduceStringLiteralToImportClauseAction                        = &_Action{_ReduceAction, 0, _ReduceStringLiteralToImportClause}
	_ReduceAliasToImportClauseAction                                = &_Action{_ReduceAction, 0, _ReduceAliasToImportClause}
	_ReduceImplicitToImportClauseTerminalAction                     = &_Action{_ReduceAction, 0, _ReduceImplicitToImportClauseTerminal}
	_ReduceExplicitToImportClauseTerminalAction                     = &_Action{_ReduceAction, 0, _ReduceExplicitToImportClauseTerminal}
	_ReduceFirstToImportClausesAction                               = &_Action{_ReduceAction, 0, _ReduceFirstToImportClauses}
	_ReduceAddToImportClausesAction                                 = &_Action{_ReduceAction, 0, _ReduceAddToImportClauses}
	_ReduceCasePatternToCasePatternsAction                          = &_Action{_ReduceAction, 0, _ReduceCasePatternToCasePatterns}
	_ReduceMultipleToCasePatternsAction                             = &_Action{_ReduceAction, 0, _ReduceMultipleToCasePatterns}
	_ReduceToVarDeclPatternAction                                   = &_Action{_ReduceAction, 0, _ReduceToVarDeclPattern}
	_ReduceVarToVarOrLetAction                                      = &_Action{_ReduceAction, 0, _ReduceVarToVarOrLet}
	_ReduceLetToVarOrLetAction                                      = &_Action{_ReduceAction, 0, _ReduceLetToVarOrLet}
	_ReduceIdentifierToVarPatternAction                             = &_Action{_ReduceAction, 0, _ReduceIdentifierToVarPattern}
	_ReduceTuplePatternToVarPatternAction                           = &_Action{_ReduceAction, 0, _ReduceTuplePatternToVarPattern}
	_ReduceToTuplePatternAction                                     = &_Action{_ReduceAction, 0, _ReduceToTuplePattern}
	_ReduceFieldVarPatternToFieldVarPatternsAction                  = &_Action{_ReduceAction, 0, _ReduceFieldVarPatternToFieldVarPatterns}
	_ReduceAddToFieldVarPatternsAction                              = &_Action{_ReduceAction, 0, _ReduceAddToFieldVarPatterns}
	_ReducePositionalToFieldVarPatternAction                        = &_Action{_ReduceAction, 0, _ReducePositionalToFieldVarPattern}
	_ReduceNamedToFieldVarPatternAction                             = &_Action{_ReduceAction, 0, _ReduceNamedToFieldVarPattern}
	_ReduceDotDotDotToFieldVarPatternAction                         = &_Action{_ReduceAction, 0, _ReduceDotDotDotToFieldVarPattern}
	_ReduceValueTypeToOptionalValueTypeAction                       = &_Action{_ReduceAction, 0, _ReduceValueTypeToOptionalValueType}
	_ReduceNilToOptionalValueTypeAction                             = &_Action{_ReduceAction, 0, _ReduceNilToOptionalValueType}
	_ReduceToAssignPatternAction                                    = &_Action{_ReduceAction, 0, _ReduceToAssignPattern}
	_ReduceAssignPatternToCasePatternAction                         = &_Action{_ReduceAction, 0, _ReduceAssignPatternToCasePattern}
	_ReduceEnumMatchPatternToCasePatternAction                      = &_Action{_ReduceAction, 0, _ReduceEnumMatchPatternToCasePattern}
	_ReduceEnumVarDeclPatternToCasePatternAction                    = &_Action{_ReduceAction, 0, _ReduceEnumVarDeclPatternToCasePattern}
	_ReduceIfExprToExpressionAction                                 = &_Action{_ReduceAction, 0, _ReduceIfExprToExpression}
	_ReduceSwitchExprToExpressionAction                             = &_Action{_ReduceAction, 0, _ReduceSwitchExprToExpression}
	_ReduceLoopExprToExpressionAction                               = &_Action{_ReduceAction, 0, _ReduceLoopExprToExpression}
	_ReduceSequenceExprToExpressionAction                           = &_Action{_ReduceAction, 0, _ReduceSequenceExprToExpression}
	_ReduceLabelDeclToOptionalLabelDeclAction                       = &_Action{_ReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}
	_ReduceUnlabelledToOptionalLabelDeclAction                      = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}
	_ReduceOrExprToSequenceExprAction                               = &_Action{_ReduceAction, 0, _ReduceOrExprToSequenceExpr}
	_ReduceVarDeclPatternToSequenceExprAction                       = &_Action{_ReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}
	_ReduceAssignVarPatternToSequenceExprAction                     = &_Action{_ReduceAction, 0, _ReduceAssignVarPatternToSequenceExpr}
	_ReduceNoElseToIfExprAction                                     = &_Action{_ReduceAction, 0, _ReduceNoElseToIfExpr}
	_ReduceIfElseToIfExprAction                                     = &_Action{_ReduceAction, 0, _ReduceIfElseToIfExpr}
	_ReduceMultiIfElseToIfExprAction                                = &_Action{_ReduceAction, 0, _ReduceMultiIfElseToIfExpr}
	_ReduceSequenceExprToConditionAction                            = &_Action{_ReduceAction, 0, _ReduceSequenceExprToCondition}
	_ReduceCaseToConditionAction                                    = &_Action{_ReduceAction, 0, _ReduceCaseToCondition}
	_ReduceToSwitchExprAction                                       = &_Action{_ReduceAction, 0, _ReduceToSwitchExpr}
	_ReduceInfiniteToLoopExprAction                                 = &_Action{_ReduceAction, 0, _ReduceInfiniteToLoopExpr}
	_ReduceDoWhileToLoopExprAction                                  = &_Action{_ReduceAction, 0, _ReduceDoWhileToLoopExpr}
	_ReduceWhileToLoopExprAction                                    = &_Action{_ReduceAction, 0, _ReduceWhileToLoopExpr}
	_ReduceIteratorToLoopExprAction                                 = &_Action{_ReduceAction, 0, _ReduceIteratorToLoopExpr}
	_ReduceForToLoopExprAction                                      = &_Action{_ReduceAction, 0, _ReduceForToLoopExpr}
	_ReduceSequenceExprToOptionalSequenceExprAction                 = &_Action{_ReduceAction, 0, _ReduceSequenceExprToOptionalSequenceExpr}
	_ReduceNilToOptionalSequenceExprAction                          = &_Action{_ReduceAction, 0, _ReduceNilToOptionalSequenceExpr}
	_ReduceSequenceExprToForAssignmentAction                        = &_Action{_ReduceAction, 0, _ReduceSequenceExprToForAssignment}
	_ReduceAssignToForAssignmentAction                              = &_Action{_ReduceAction, 0, _ReduceAssignToForAssignment}
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
	_ReduceDotDotDotToArgumentAction                                = &_Action{_ReduceAction, 0, _ReduceDotDotDotToArgument}
	_ReducePairToColonExpressionsAction                             = &_Action{_ReduceAction, 0, _ReducePairToColonExpressions}
	_ReduceAddToColonExpressionsAction                              = &_Action{_ReduceAction, 0, _ReduceAddToColonExpressions}
	_ReduceExpressionToOptionalExpressionAction                     = &_Action{_ReduceAction, 0, _ReduceExpressionToOptionalExpression}
	_ReduceNilToOptionalExpressionAction                            = &_Action{_ReduceAction, 0, _ReduceNilToOptionalExpression}
	_ReduceLiteralExprToAtomExprAction                              = &_Action{_ReduceAction, 0, _ReduceLiteralExprToAtomExpr}
	_ReduceIdentifierToAtomExprAction                               = &_Action{_ReduceAction, 0, _ReduceIdentifierToAtomExpr}
	_ReduceBlockExprToAtomExprAction                                = &_Action{_ReduceAction, 0, _ReduceBlockExprToAtomExpr}
	_ReduceAnonymousFuncExprToAtomExprAction                        = &_Action{_ReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}
	_ReduceInitializeExprToAtomExprAction                           = &_Action{_ReduceAction, 0, _ReduceInitializeExprToAtomExpr}
	_ReduceImplicitStructExprToAtomExprAction                       = &_Action{_ReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}
	_ReduceParseErrorToAtomExprAction                               = &_Action{_ReduceAction, 0, _ReduceParseErrorToAtomExpr}
	_ReduceTrueToLiteralExprAction                                  = &_Action{_ReduceAction, 0, _ReduceTrueToLiteralExpr}
	_ReduceFalseToLiteralExprAction                                 = &_Action{_ReduceAction, 0, _ReduceFalseToLiteralExpr}
	_ReduceIntegerLiteralToLiteralExprAction                        = &_Action{_ReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}
	_ReduceFloatLiteralToLiteralExprAction                          = &_Action{_ReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}
	_ReduceRuneLiteralToLiteralExprAction                           = &_Action{_ReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}
	_ReduceStringLiteralToLiteralExprAction                         = &_Action{_ReduceAction, 0, _ReduceStringLiteralToLiteralExpr}
	_ReduceToImplicitStructExprAction                               = &_Action{_ReduceAction, 0, _ReduceToImplicitStructExpr}
	_ReduceAtomExprToAccessExprAction                               = &_Action{_ReduceAction, 0, _ReduceAtomExprToAccessExpr}
	_ReduceAccessToAccessExprAction                                 = &_Action{_ReduceAction, 0, _ReduceAccessToAccessExpr}
	_ReduceCallExprToAccessExprAction                               = &_Action{_ReduceAction, 0, _ReduceCallExprToAccessExpr}
	_ReduceIndexToAccessExprAction                                  = &_Action{_ReduceAction, 0, _ReduceIndexToAccessExpr}
	_ReduceAccessExprToPostfixUnaryExprAction                       = &_Action{_ReduceAction, 0, _ReduceAccessExprToPostfixUnaryExpr}
	_ReduceToPostfixUnaryExprAction                                 = &_Action{_ReduceAction, 0, _ReduceToPostfixUnaryExpr}
	_ReduceNotToPrefixUnaryOpAction                                 = &_Action{_ReduceAction, 0, _ReduceNotToPrefixUnaryOp}
	_ReduceBitNegToPrefixUnaryOpAction                              = &_Action{_ReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}
	_ReduceSubToPrefixUnaryOpAction                                 = &_Action{_ReduceAction, 0, _ReduceSubToPrefixUnaryOp}
	_ReduceMulToPrefixUnaryOpAction                                 = &_Action{_ReduceAction, 0, _ReduceMulToPrefixUnaryOp}
	_ReduceBitAndToPrefixUnaryOpAction                              = &_Action{_ReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}
	_ReducePostfixUnaryExprToPrefixUnaryExprAction                  = &_Action{_ReduceAction, 0, _ReducePostfixUnaryExprToPrefixUnaryExpr}
	_ReduceToPrefixUnaryExprAction                                  = &_Action{_ReduceAction, 0, _ReduceToPrefixUnaryExpr}
	_ReduceMulToMulOpAction                                         = &_Action{_ReduceAction, 0, _ReduceMulToMulOp}
	_ReduceDivToMulOpAction                                         = &_Action{_ReduceAction, 0, _ReduceDivToMulOp}
	_ReduceModToMulOpAction                                         = &_Action{_ReduceAction, 0, _ReduceModToMulOp}
	_ReduceBitAndToMulOpAction                                      = &_Action{_ReduceAction, 0, _ReduceBitAndToMulOp}
	_ReduceBitLshiftToMulOpAction                                   = &_Action{_ReduceAction, 0, _ReduceBitLshiftToMulOp}
	_ReduceBitRshiftToMulOpAction                                   = &_Action{_ReduceAction, 0, _ReduceBitRshiftToMulOp}
	_ReducePrefixUnaryExprToMulExprAction                           = &_Action{_ReduceAction, 0, _ReducePrefixUnaryExprToMulExpr}
	_ReduceToMulExprAction                                          = &_Action{_ReduceAction, 0, _ReduceToMulExpr}
	_ReduceAddToAddOpAction                                         = &_Action{_ReduceAction, 0, _ReduceAddToAddOp}
	_ReduceSubToAddOpAction                                         = &_Action{_ReduceAction, 0, _ReduceSubToAddOp}
	_ReduceBitOrToAddOpAction                                       = &_Action{_ReduceAction, 0, _ReduceBitOrToAddOp}
	_ReduceBitXorToAddOpAction                                      = &_Action{_ReduceAction, 0, _ReduceBitXorToAddOp}
	_ReduceMulExprToAddExprAction                                   = &_Action{_ReduceAction, 0, _ReduceMulExprToAddExpr}
	_ReduceToAddExprAction                                          = &_Action{_ReduceAction, 0, _ReduceToAddExpr}
	_ReduceEqualToCmpOpAction                                       = &_Action{_ReduceAction, 0, _ReduceEqualToCmpOp}
	_ReduceNotEqualToCmpOpAction                                    = &_Action{_ReduceAction, 0, _ReduceNotEqualToCmpOp}
	_ReduceLessToCmpOpAction                                        = &_Action{_ReduceAction, 0, _ReduceLessToCmpOp}
	_ReduceLessOrEqualToCmpOpAction                                 = &_Action{_ReduceAction, 0, _ReduceLessOrEqualToCmpOp}
	_ReduceGreaterToCmpOpAction                                     = &_Action{_ReduceAction, 0, _ReduceGreaterToCmpOp}
	_ReduceGreaterOrEqualToCmpOpAction                              = &_Action{_ReduceAction, 0, _ReduceGreaterOrEqualToCmpOp}
	_ReduceAddExprToCmpExprAction                                   = &_Action{_ReduceAction, 0, _ReduceAddExprToCmpExpr}
	_ReduceToCmpExprAction                                          = &_Action{_ReduceAction, 0, _ReduceToCmpExpr}
	_ReduceCmpExprToAndExprAction                                   = &_Action{_ReduceAction, 0, _ReduceCmpExprToAndExpr}
	_ReduceToAndExprAction                                          = &_Action{_ReduceAction, 0, _ReduceToAndExpr}
	_ReduceAndExprToOrExprAction                                    = &_Action{_ReduceAction, 0, _ReduceAndExprToOrExpr}
	_ReduceToOrExprAction                                           = &_Action{_ReduceAction, 0, _ReduceToOrExpr}
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
	_ReduceParseErrorToAtomTypeAction                               = &_Action{_ReduceAction, 0, _ReduceParseErrorToAtomType}
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
	_ReduceInferredRefArgToParameterDefAction                       = &_Action{_ReduceAction, 0, _ReduceInferredRefArgToParameterDef}
	_ReduceInferredRefVarargToParameterDefAction                    = &_Action{_ReduceAction, 0, _ReduceInferredRefVarargToParameterDef}
	_ReduceArgToParameterDefAction                                  = &_Action{_ReduceAction, 0, _ReduceArgToParameterDef}
	_ReduceVarargToParameterDefAction                               = &_Action{_ReduceAction, 0, _ReduceVarargToParameterDef}
	_ReduceParameterDefToParameterDefsAction                        = &_Action{_ReduceAction, 0, _ReduceParameterDefToParameterDefs}
	_ReduceAddToParameterDefsAction                                 = &_Action{_ReduceAction, 0, _ReduceAddToParameterDefs}
	_ReduceParameterDefsToOptionalParameterDefsAction               = &_Action{_ReduceAction, 0, _ReduceParameterDefsToOptionalParameterDefs}
	_ReduceNilToOptionalParameterDefsAction                         = &_Action{_ReduceAction, 0, _ReduceNilToOptionalParameterDefs}
	_ReduceFuncDefToNamedFuncDefAction                              = &_Action{_ReduceAction, 0, _ReduceFuncDefToNamedFuncDef}
	_ReduceMethodDefToNamedFuncDefAction                            = &_Action{_ReduceAction, 0, _ReduceMethodDefToNamedFuncDef}
	_ReduceAliasToNamedFuncDefAction                                = &_Action{_ReduceAction, 0, _ReduceAliasToNamedFuncDef}
	_ReduceToAnonymousFuncExprAction                                = &_Action{_ReduceAction, 0, _ReduceToAnonymousFuncExpr}
	_ReduceNoSpecToPackageDefAction                                 = &_Action{_ReduceAction, 0, _ReduceNoSpecToPackageDef}
	_ReduceWithSpecToPackageDefAction                               = &_Action{_ReduceAction, 0, _ReduceWithSpecToPackageDef}
)

var _ActionTable = _ActionTableType{
	{_State6, _EndMarker}:                         &_Action{_AcceptAction, 0, 0},
	{_State7, _EndMarker}:                         &_Action{_AcceptAction, 0, 0},
	{_State8, _EndMarker}:                         &_Action{_AcceptAction, 0, 0},
	{_State9, _EndMarker}:                         &_Action{_AcceptAction, 0, 0},
	{_State10, _EndMarker}:                        &_Action{_AcceptAction, 0, 0},
	{_State1, NewlinesToken}:                      _GotoState11Action,
	{_State1, SourceType}:                         _GotoState6Action,
	{_State1, OptionalDefinitionsType}:            _GotoState12Action,
	{_State1, OptionalNewlinesType}:               _GotoState13Action,
	{_State2, PackageToken}:                       _GotoState14Action,
	{_State2, PackageDefType}:                     _GotoState7Action,
	{_State3, TypeToken}:                          _GotoState15Action,
	{_State3, TypeDefType}:                        _GotoState8Action,
	{_State4, FuncToken}:                          _GotoState16Action,
	{_State4, NamedFuncDefType}:                   _GotoState9Action,
	{_State5, IntegerLiteralToken}:                _GotoState24Action,
	{_State5, FloatLiteralToken}:                  _GotoState20Action,
	{_State5, RuneLiteralToken}:                   _GotoState32Action,
	{_State5, StringLiteralToken}:                 _GotoState33Action,
	{_State5, IdentifierToken}:                    _GotoState23Action,
	{_State5, TrueToken}:                          _GotoState36Action,
	{_State5, FalseToken}:                         _GotoState19Action,
	{_State5, StructToken}:                        _GotoState34Action,
	{_State5, FuncToken}:                          _GotoState21Action,
	{_State5, VarToken}:                           _GotoState37Action,
	{_State5, LetToken}:                           _GotoState27Action,
	{_State5, NotToken}:                           _GotoState30Action,
	{_State5, LabelDeclToken}:                     _GotoState25Action,
	{_State5, LparenToken}:                        _GotoState28Action,
	{_State5, LbracketToken}:                      _GotoState26Action,
	{_State5, SubToken}:                           _GotoState35Action,
	{_State5, MulToken}:                           _GotoState29Action,
	{_State5, BitNegToken}:                        _GotoState18Action,
	{_State5, BitAndToken}:                        _GotoState17Action,
	{_State5, GreaterToken}:                       _GotoState22Action,
	{_State5, ParseErrorToken}:                    _GotoState31Action,
	{_State5, VarDeclPatternType}:                 _GotoState56Action,
	{_State5, VarOrLetType}:                       _GotoState57Action,
	{_State5, ExpressionType}:                     _GotoState10Action,
	{_State5, OptionalLabelDeclType}:              _GotoState50Action,
	{_State5, SequenceExprType}:                   _GotoState55Action,
	{_State5, CallExprType}:                       _GotoState43Action,
	{_State5, AtomExprType}:                       _GotoState42Action,
	{_State5, LiteralExprType}:                    _GotoState48Action,
	{_State5, ImplicitStructExprType}:             _GotoState46Action,
	{_State5, AccessExprType}:                     _GotoState38Action,
	{_State5, PostfixUnaryExprType}:               _GotoState52Action,
	{_State5, PrefixUnaryOpType}:                  _GotoState54Action,
	{_State5, PrefixUnaryExprType}:                _GotoState53Action,
	{_State5, MulExprType}:                        _GotoState49Action,
	{_State5, AddExprType}:                        _GotoState39Action,
	{_State5, CmpExprType}:                        _GotoState44Action,
	{_State5, AndExprType}:                        _GotoState40Action,
	{_State5, OrExprType}:                         _GotoState51Action,
	{_State5, InitializableTypeType}:              _GotoState47Action,
	{_State5, ExplicitStructDefType}:              _GotoState45Action,
	{_State5, AnonymousFuncExprType}:              _GotoState41Action,
	{_State13, CommentGroupsToken}:                _GotoState58Action,
	{_State13, PackageToken}:                      _GotoState14Action,
	{_State13, TypeToken}:                         _GotoState15Action,
	{_State13, FuncToken}:                         _GotoState16Action,
	{_State13, VarToken}:                          _GotoState37Action,
	{_State13, LetToken}:                          _GotoState27Action,
	{_State13, LbraceToken}:                       _GotoState59Action,
	{_State13, DefinitionsType}:                   _GotoState61Action,
	{_State13, DefinitionType}:                    _GotoState60Action,
	{_State13, StatementBlockType}:                _GotoState64Action,
	{_State13, VarDeclPatternType}:                _GotoState66Action,
	{_State13, VarOrLetType}:                      _GotoState57Action,
	{_State13, TypeDefType}:                       _GotoState65Action,
	{_State13, NamedFuncDefType}:                  _GotoState62Action,
	{_State13, PackageDefType}:                    _GotoState63Action,
	{_State14, LbraceToken}:                       _GotoState59Action,
	{_State14, StatementBlockType}:                _GotoState67Action,
	{_State15, IdentifierToken}:                   _GotoState68Action,
	{_State16, IdentifierToken}:                   _GotoState69Action,
	{_State16, LparenToken}:                       _GotoState70Action,
	{_State21, LparenToken}:                       _GotoState71Action,
	{_State22, IntegerLiteralToken}:               _GotoState24Action,
	{_State22, FloatLiteralToken}:                 _GotoState20Action,
	{_State22, RuneLiteralToken}:                  _GotoState32Action,
	{_State22, StringLiteralToken}:                _GotoState33Action,
	{_State22, IdentifierToken}:                   _GotoState23Action,
	{_State22, TrueToken}:                         _GotoState36Action,
	{_State22, FalseToken}:                        _GotoState19Action,
	{_State22, StructToken}:                       _GotoState34Action,
	{_State22, FuncToken}:                         _GotoState21Action,
	{_State22, VarToken}:                          _GotoState37Action,
	{_State22, LetToken}:                          _GotoState27Action,
	{_State22, NotToken}:                          _GotoState30Action,
	{_State22, LabelDeclToken}:                    _GotoState25Action,
	{_State22, LparenToken}:                       _GotoState28Action,
	{_State22, LbracketToken}:                     _GotoState26Action,
	{_State22, SubToken}:                          _GotoState35Action,
	{_State22, MulToken}:                          _GotoState29Action,
	{_State22, BitNegToken}:                       _GotoState18Action,
	{_State22, BitAndToken}:                       _GotoState17Action,
	{_State22, GreaterToken}:                      _GotoState22Action,
	{_State22, ParseErrorToken}:                   _GotoState31Action,
	{_State22, VarDeclPatternType}:                _GotoState56Action,
	{_State22, VarOrLetType}:                      _GotoState57Action,
	{_State22, OptionalLabelDeclType}:             _GotoState72Action,
	{_State22, SequenceExprType}:                  _GotoState73Action,
	{_State22, CallExprType}:                      _GotoState43Action,
	{_State22, AtomExprType}:                      _GotoState42Action,
	{_State22, LiteralExprType}:                   _GotoState48Action,
	{_State22, ImplicitStructExprType}:            _GotoState46Action,
	{_State22, AccessExprType}:                    _GotoState38Action,
	{_State22, PostfixUnaryExprType}:              _GotoState52Action,
	{_State22, PrefixUnaryOpType}:                 _GotoState54Action,
	{_State22, PrefixUnaryExprType}:               _GotoState53Action,
	{_State22, MulExprType}:                       _GotoState49Action,
	{_State22, AddExprType}:                       _GotoState39Action,
	{_State22, CmpExprType}:                       _GotoState44Action,
	{_State22, AndExprType}:                       _GotoState40Action,
	{_State22, OrExprType}:                        _GotoState51Action,
	{_State22, InitializableTypeType}:             _GotoState47Action,
	{_State22, ExplicitStructDefType}:             _GotoState45Action,
	{_State22, AnonymousFuncExprType}:             _GotoState41Action,
	{_State26, IdentifierToken}:                   _GotoState80Action,
	{_State26, StructToken}:                       _GotoState34Action,
	{_State26, EnumToken}:                         _GotoState77Action,
	{_State26, TraitToken}:                        _GotoState85Action,
	{_State26, FuncToken}:                         _GotoState79Action,
	{_State26, LparenToken}:                       _GotoState81Action,
	{_State26, LbracketToken}:                     _GotoState26Action,
	{_State26, DotToken}:                          _GotoState76Action,
	{_State26, QuestionToken}:                     _GotoState83Action,
	{_State26, ExclaimToken}:                      _GotoState78Action,
	{_State26, TildeTildeToken}:                   _GotoState84Action,
	{_State26, BitNegToken}:                       _GotoState75Action,
	{_State26, BitAndToken}:                       _GotoState74Action,
	{_State26, ParseErrorToken}:                   _GotoState82Action,
	{_State26, InitializableTypeType}:             _GotoState91Action,
	{_State26, AtomTypeType}:                      _GotoState86Action,
	{_State26, ReturnableTypeType}:                _GotoState92Action,
	{_State26, ValueTypeType}:                     _GotoState94Action,
	{_State26, ImplicitStructDefType}:             _GotoState90Action,
	{_State26, ExplicitStructDefType}:             _GotoState45Action,
	{_State26, ImplicitEnumDefType}:               _GotoState89Action,
	{_State26, ExplicitEnumDefType}:               _GotoState87Action,
	{_State26, TraitDefType}:                      _GotoState93Action,
	{_State26, FuncTypeType}:                      _GotoState88Action,
	{_State28, IntegerLiteralToken}:               _GotoState24Action,
	{_State28, FloatLiteralToken}:                 _GotoState20Action,
	{_State28, RuneLiteralToken}:                  _GotoState32Action,
	{_State28, StringLiteralToken}:                _GotoState33Action,
	{_State28, IdentifierToken}:                   _GotoState96Action,
	{_State28, TrueToken}:                         _GotoState36Action,
	{_State28, FalseToken}:                        _GotoState19Action,
	{_State28, StructToken}:                       _GotoState34Action,
	{_State28, FuncToken}:                         _GotoState21Action,
	{_State28, VarToken}:                          _GotoState37Action,
	{_State28, LetToken}:                          _GotoState27Action,
	{_State28, NotToken}:                          _GotoState30Action,
	{_State28, LabelDeclToken}:                    _GotoState25Action,
	{_State28, LparenToken}:                       _GotoState28Action,
	{_State28, LbracketToken}:                     _GotoState26Action,
	{_State28, DotDotDotToken}:                    _GotoState95Action,
	{_State28, SubToken}:                          _GotoState35Action,
	{_State28, MulToken}:                          _GotoState29Action,
	{_State28, BitNegToken}:                       _GotoState18Action,
	{_State28, BitAndToken}:                       _GotoState17Action,
	{_State28, GreaterToken}:                      _GotoState22Action,
	{_State28, ParseErrorToken}:                   _GotoState31Action,
	{_State28, VarDeclPatternType}:                _GotoState56Action,
	{_State28, VarOrLetType}:                      _GotoState57Action,
	{_State28, ExpressionType}:                    _GotoState100Action,
	{_State28, OptionalLabelDeclType}:             _GotoState50Action,
	{_State28, SequenceExprType}:                  _GotoState55Action,
	{_State28, CallExprType}:                      _GotoState43Action,
	{_State28, OptionalArgumentsType}:             _GotoState101Action,
	{_State28, ArgumentsType}:                     _GotoState98Action,
	{_State28, ArgumentType}:                      _GotoState97Action,
	{_State28, ColonExpressionsType}:              _GotoState99Action,
	{_State28, OptionalExpressionType}:            _GotoState102Action,
	{_State28, AtomExprType}:                      _GotoState42Action,
	{_State28, LiteralExprType}:                   _GotoState48Action,
	{_State28, ImplicitStructExprType}:            _GotoState46Action,
	{_State28, AccessExprType}:                    _GotoState38Action,
	{_State28, PostfixUnaryExprType}:              _GotoState52Action,
	{_State28, PrefixUnaryOpType}:                 _GotoState54Action,
	{_State28, PrefixUnaryExprType}:               _GotoState53Action,
	{_State28, MulExprType}:                       _GotoState49Action,
	{_State28, AddExprType}:                       _GotoState39Action,
	{_State28, CmpExprType}:                       _GotoState44Action,
	{_State28, AndExprType}:                       _GotoState40Action,
	{_State28, OrExprType}:                        _GotoState51Action,
	{_State28, InitializableTypeType}:             _GotoState47Action,
	{_State28, ExplicitStructDefType}:             _GotoState45Action,
	{_State28, AnonymousFuncExprType}:             _GotoState41Action,
	{_State34, LparenToken}:                       _GotoState103Action,
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
	{_State44, EqualToken}:                        _GotoState115Action,
	{_State44, NotEqualToken}:                     _GotoState120Action,
	{_State44, LessToken}:                         _GotoState118Action,
	{_State44, LessOrEqualToken}:                  _GotoState119Action,
	{_State44, GreaterToken}:                      _GotoState116Action,
	{_State44, GreaterOrEqualToken}:               _GotoState117Action,
	{_State44, CmpOpType}:                         _GotoState121Action,
	{_State47, LparenToken}:                       _GotoState122Action,
	{_State49, MulToken}:                          _GotoState128Action,
	{_State49, DivToken}:                          _GotoState126Action,
	{_State49, ModToken}:                          _GotoState127Action,
	{_State49, BitAndToken}:                       _GotoState123Action,
	{_State49, BitLshiftToken}:                    _GotoState124Action,
	{_State49, BitRshiftToken}:                    _GotoState125Action,
	{_State49, MulOpType}:                         _GotoState129Action,
	{_State50, IfToken}:                           _GotoState132Action,
	{_State50, SwitchToken}:                       _GotoState133Action,
	{_State50, ForToken}:                          _GotoState131Action,
	{_State50, DoToken}:                           _GotoState130Action,
	{_State50, LbraceToken}:                       _GotoState59Action,
	{_State50, StatementBlockType}:                _GotoState136Action,
	{_State50, IfExprType}:                        _GotoState134Action,
	{_State50, SwitchExprType}:                    _GotoState137Action,
	{_State50, LoopExprType}:                      _GotoState135Action,
	{_State51, OrToken}:                           _GotoState138Action,
	{_State54, IntegerLiteralToken}:               _GotoState24Action,
	{_State54, FloatLiteralToken}:                 _GotoState20Action,
	{_State54, RuneLiteralToken}:                  _GotoState32Action,
	{_State54, StringLiteralToken}:                _GotoState33Action,
	{_State54, IdentifierToken}:                   _GotoState23Action,
	{_State54, TrueToken}:                         _GotoState36Action,
	{_State54, FalseToken}:                        _GotoState19Action,
	{_State54, StructToken}:                       _GotoState34Action,
	{_State54, FuncToken}:                         _GotoState21Action,
	{_State54, NotToken}:                          _GotoState30Action,
	{_State54, LabelDeclToken}:                    _GotoState25Action,
	{_State54, LparenToken}:                       _GotoState28Action,
	{_State54, LbracketToken}:                     _GotoState26Action,
	{_State54, SubToken}:                          _GotoState35Action,
	{_State54, MulToken}:                          _GotoState29Action,
	{_State54, BitNegToken}:                       _GotoState18Action,
	{_State54, BitAndToken}:                       _GotoState17Action,
	{_State54, ParseErrorToken}:                   _GotoState31Action,
	{_State54, OptionalLabelDeclType}:             _GotoState72Action,
	{_State54, CallExprType}:                      _GotoState43Action,
	{_State54, AtomExprType}:                      _GotoState42Action,
	{_State54, LiteralExprType}:                   _GotoState48Action,
	{_State54, ImplicitStructExprType}:            _GotoState46Action,
	{_State54, AccessExprType}:                    _GotoState38Action,
	{_State54, PostfixUnaryExprType}:              _GotoState52Action,
	{_State54, PrefixUnaryOpType}:                 _GotoState54Action,
	{_State54, PrefixUnaryExprType}:               _GotoState139Action,
	{_State54, InitializableTypeType}:             _GotoState47Action,
	{_State54, ExplicitStructDefType}:             _GotoState45Action,
	{_State54, AnonymousFuncExprType}:             _GotoState41Action,
	{_State57, IdentifierToken}:                   _GotoState140Action,
	{_State57, LparenToken}:                       _GotoState141Action,
	{_State57, VarPatternType}:                    _GotoState143Action,
	{_State57, TuplePatternType}:                  _GotoState142Action,
	{_State59, StatementsType}:                    _GotoState144Action,
	{_State61, NewlinesToken}:                     _GotoState145Action,
	{_State61, OptionalNewlinesType}:              _GotoState146Action,
	{_State66, AssignToken}:                       _GotoState147Action,
	{_State68, DollarLbracketToken}:               _GotoState149Action,
	{_State68, AssignToken}:                       _GotoState148Action,
	{_State68, OptionalGenericParametersType}:     _GotoState150Action,
	{_State69, DollarLbracketToken}:               _GotoState149Action,
	{_State69, AssignToken}:                       _GotoState151Action,
	{_State69, OptionalGenericParametersType}:     _GotoState152Action,
	{_State70, IdentifierToken}:                   _GotoState153Action,
	{_State70, ParameterDefType}:                  _GotoState154Action,
	{_State71, IdentifierToken}:                   _GotoState153Action,
	{_State71, ParameterDefType}:                  _GotoState156Action,
	{_State71, ParameterDefsType}:                 _GotoState157Action,
	{_State71, OptionalParameterDefsType}:         _GotoState155Action,
	{_State72, LbraceToken}:                       _GotoState59Action,
	{_State72, StatementBlockType}:                _GotoState136Action,
	{_State74, IdentifierToken}:                   _GotoState80Action,
	{_State74, StructToken}:                       _GotoState34Action,
	{_State74, EnumToken}:                         _GotoState77Action,
	{_State74, TraitToken}:                        _GotoState85Action,
	{_State74, FuncToken}:                         _GotoState79Action,
	{_State74, LparenToken}:                       _GotoState81Action,
	{_State74, LbracketToken}:                     _GotoState26Action,
	{_State74, DotToken}:                          _GotoState76Action,
	{_State74, QuestionToken}:                     _GotoState83Action,
	{_State74, ExclaimToken}:                      _GotoState78Action,
	{_State74, TildeTildeToken}:                   _GotoState84Action,
	{_State74, BitNegToken}:                       _GotoState75Action,
	{_State74, BitAndToken}:                       _GotoState74Action,
	{_State74, ParseErrorToken}:                   _GotoState82Action,
	{_State74, InitializableTypeType}:             _GotoState91Action,
	{_State74, AtomTypeType}:                      _GotoState86Action,
	{_State74, ReturnableTypeType}:                _GotoState158Action,
	{_State74, ImplicitStructDefType}:             _GotoState90Action,
	{_State74, ExplicitStructDefType}:             _GotoState45Action,
	{_State74, ImplicitEnumDefType}:               _GotoState89Action,
	{_State74, ExplicitEnumDefType}:               _GotoState87Action,
	{_State74, TraitDefType}:                      _GotoState93Action,
	{_State74, FuncTypeType}:                      _GotoState88Action,
	{_State75, IdentifierToken}:                   _GotoState80Action,
	{_State75, StructToken}:                       _GotoState34Action,
	{_State75, EnumToken}:                         _GotoState77Action,
	{_State75, TraitToken}:                        _GotoState85Action,
	{_State75, FuncToken}:                         _GotoState79Action,
	{_State75, LparenToken}:                       _GotoState81Action,
	{_State75, LbracketToken}:                     _GotoState26Action,
	{_State75, DotToken}:                          _GotoState76Action,
	{_State75, QuestionToken}:                     _GotoState83Action,
	{_State75, ExclaimToken}:                      _GotoState78Action,
	{_State75, TildeTildeToken}:                   _GotoState84Action,
	{_State75, BitNegToken}:                       _GotoState75Action,
	{_State75, BitAndToken}:                       _GotoState74Action,
	{_State75, ParseErrorToken}:                   _GotoState82Action,
	{_State75, InitializableTypeType}:             _GotoState91Action,
	{_State75, AtomTypeType}:                      _GotoState86Action,
	{_State75, ReturnableTypeType}:                _GotoState159Action,
	{_State75, ImplicitStructDefType}:             _GotoState90Action,
	{_State75, ExplicitStructDefType}:             _GotoState45Action,
	{_State75, ImplicitEnumDefType}:               _GotoState89Action,
	{_State75, ExplicitEnumDefType}:               _GotoState87Action,
	{_State75, TraitDefType}:                      _GotoState93Action,
	{_State75, FuncTypeType}:                      _GotoState88Action,
	{_State76, DollarLbracketToken}:               _GotoState104Action,
	{_State76, OptionalGenericBindingType}:        _GotoState160Action,
	{_State77, LparenToken}:                       _GotoState161Action,
	{_State78, IdentifierToken}:                   _GotoState80Action,
	{_State78, StructToken}:                       _GotoState34Action,
	{_State78, EnumToken}:                         _GotoState77Action,
	{_State78, TraitToken}:                        _GotoState85Action,
	{_State78, FuncToken}:                         _GotoState79Action,
	{_State78, LparenToken}:                       _GotoState81Action,
	{_State78, LbracketToken}:                     _GotoState26Action,
	{_State78, DotToken}:                          _GotoState76Action,
	{_State78, QuestionToken}:                     _GotoState83Action,
	{_State78, ExclaimToken}:                      _GotoState78Action,
	{_State78, TildeTildeToken}:                   _GotoState84Action,
	{_State78, BitNegToken}:                       _GotoState75Action,
	{_State78, BitAndToken}:                       _GotoState74Action,
	{_State78, ParseErrorToken}:                   _GotoState82Action,
	{_State78, InitializableTypeType}:             _GotoState91Action,
	{_State78, AtomTypeType}:                      _GotoState86Action,
	{_State78, ReturnableTypeType}:                _GotoState162Action,
	{_State78, ImplicitStructDefType}:             _GotoState90Action,
	{_State78, ExplicitStructDefType}:             _GotoState45Action,
	{_State78, ImplicitEnumDefType}:               _GotoState89Action,
	{_State78, ExplicitEnumDefType}:               _GotoState87Action,
	{_State78, TraitDefType}:                      _GotoState93Action,
	{_State78, FuncTypeType}:                      _GotoState88Action,
	{_State79, LparenToken}:                       _GotoState163Action,
	{_State80, DotToken}:                          _GotoState164Action,
	{_State80, DollarLbracketToken}:               _GotoState104Action,
	{_State80, OptionalGenericBindingType}:        _GotoState165Action,
	{_State81, IdentifierToken}:                   _GotoState166Action,
	{_State81, UnsafeToken}:                       _GotoState167Action,
	{_State81, StructToken}:                       _GotoState34Action,
	{_State81, EnumToken}:                         _GotoState77Action,
	{_State81, TraitToken}:                        _GotoState85Action,
	{_State81, FuncToken}:                         _GotoState79Action,
	{_State81, LparenToken}:                       _GotoState81Action,
	{_State81, LbracketToken}:                     _GotoState26Action,
	{_State81, DotToken}:                          _GotoState76Action,
	{_State81, QuestionToken}:                     _GotoState83Action,
	{_State81, ExclaimToken}:                      _GotoState78Action,
	{_State81, TildeTildeToken}:                   _GotoState84Action,
	{_State81, BitNegToken}:                       _GotoState75Action,
	{_State81, BitAndToken}:                       _GotoState74Action,
	{_State81, ParseErrorToken}:                   _GotoState82Action,
	{_State81, UnsafeStatementType}:               _GotoState173Action,
	{_State81, InitializableTypeType}:             _GotoState91Action,
	{_State81, AtomTypeType}:                      _GotoState86Action,
	{_State81, ReturnableTypeType}:                _GotoState92Action,
	{_State81, ValueTypeType}:                     _GotoState174Action,
	{_State81, FieldDefType}:                      _GotoState169Action,
	{_State81, ImplicitFieldDefsType}:             _GotoState171Action,
	{_State81, OptionalImplicitFieldDefsType}:     _GotoState172Action,
	{_State81, ImplicitStructDefType}:             _GotoState90Action,
	{_State81, ExplicitStructDefType}:             _GotoState45Action,
	{_State81, EnumValueDefType}:                  _GotoState168Action,
	{_State81, ImplicitEnumValueDefsType}:         _GotoState170Action,
	{_State81, ImplicitEnumDefType}:               _GotoState89Action,
	{_State81, ExplicitEnumDefType}:               _GotoState87Action,
	{_State81, TraitDefType}:                      _GotoState93Action,
	{_State81, FuncTypeType}:                      _GotoState88Action,
	{_State83, IdentifierToken}:                   _GotoState80Action,
	{_State83, StructToken}:                       _GotoState34Action,
	{_State83, EnumToken}:                         _GotoState77Action,
	{_State83, TraitToken}:                        _GotoState85Action,
	{_State83, FuncToken}:                         _GotoState79Action,
	{_State83, LparenToken}:                       _GotoState81Action,
	{_State83, LbracketToken}:                     _GotoState26Action,
	{_State83, DotToken}:                          _GotoState76Action,
	{_State83, QuestionToken}:                     _GotoState83Action,
	{_State83, ExclaimToken}:                      _GotoState78Action,
	{_State83, TildeTildeToken}:                   _GotoState84Action,
	{_State83, BitNegToken}:                       _GotoState75Action,
	{_State83, BitAndToken}:                       _GotoState74Action,
	{_State83, ParseErrorToken}:                   _GotoState82Action,
	{_State83, InitializableTypeType}:             _GotoState91Action,
	{_State83, AtomTypeType}:                      _GotoState86Action,
	{_State83, ReturnableTypeType}:                _GotoState175Action,
	{_State83, ImplicitStructDefType}:             _GotoState90Action,
	{_State83, ExplicitStructDefType}:             _GotoState45Action,
	{_State83, ImplicitEnumDefType}:               _GotoState89Action,
	{_State83, ExplicitEnumDefType}:               _GotoState87Action,
	{_State83, TraitDefType}:                      _GotoState93Action,
	{_State83, FuncTypeType}:                      _GotoState88Action,
	{_State84, IdentifierToken}:                   _GotoState80Action,
	{_State84, StructToken}:                       _GotoState34Action,
	{_State84, EnumToken}:                         _GotoState77Action,
	{_State84, TraitToken}:                        _GotoState85Action,
	{_State84, FuncToken}:                         _GotoState79Action,
	{_State84, LparenToken}:                       _GotoState81Action,
	{_State84, LbracketToken}:                     _GotoState26Action,
	{_State84, DotToken}:                          _GotoState76Action,
	{_State84, QuestionToken}:                     _GotoState83Action,
	{_State84, ExclaimToken}:                      _GotoState78Action,
	{_State84, TildeTildeToken}:                   _GotoState84Action,
	{_State84, BitNegToken}:                       _GotoState75Action,
	{_State84, BitAndToken}:                       _GotoState74Action,
	{_State84, ParseErrorToken}:                   _GotoState82Action,
	{_State84, InitializableTypeType}:             _GotoState91Action,
	{_State84, AtomTypeType}:                      _GotoState86Action,
	{_State84, ReturnableTypeType}:                _GotoState176Action,
	{_State84, ImplicitStructDefType}:             _GotoState90Action,
	{_State84, ExplicitStructDefType}:             _GotoState45Action,
	{_State84, ImplicitEnumDefType}:               _GotoState89Action,
	{_State84, ExplicitEnumDefType}:               _GotoState87Action,
	{_State84, TraitDefType}:                      _GotoState93Action,
	{_State84, FuncTypeType}:                      _GotoState88Action,
	{_State85, LparenToken}:                       _GotoState177Action,
	{_State94, RbracketToken}:                     _GotoState182Action,
	{_State94, CommaToken}:                        _GotoState180Action,
	{_State94, ColonToken}:                        _GotoState179Action,
	{_State94, AddToken}:                          _GotoState178Action,
	{_State94, SubToken}:                          _GotoState183Action,
	{_State94, MulToken}:                          _GotoState181Action,
	{_State96, AssignToken}:                       _GotoState184Action,
	{_State98, CommaToken}:                        _GotoState185Action,
	{_State99, ColonToken}:                        _GotoState186Action,
	{_State101, RparenToken}:                      _GotoState187Action,
	{_State102, ColonToken}:                       _GotoState188Action,
	{_State103, IdentifierToken}:                  _GotoState166Action,
	{_State103, UnsafeToken}:                      _GotoState167Action,
	{_State103, StructToken}:                      _GotoState34Action,
	{_State103, EnumToken}:                        _GotoState77Action,
	{_State103, TraitToken}:                       _GotoState85Action,
	{_State103, FuncToken}:                        _GotoState79Action,
	{_State103, LparenToken}:                      _GotoState81Action,
	{_State103, LbracketToken}:                    _GotoState26Action,
	{_State103, DotToken}:                         _GotoState76Action,
	{_State103, QuestionToken}:                    _GotoState83Action,
	{_State103, ExclaimToken}:                     _GotoState78Action,
	{_State103, TildeTildeToken}:                  _GotoState84Action,
	{_State103, BitNegToken}:                      _GotoState75Action,
	{_State103, BitAndToken}:                      _GotoState74Action,
	{_State103, ParseErrorToken}:                  _GotoState82Action,
	{_State103, UnsafeStatementType}:              _GotoState173Action,
	{_State103, InitializableTypeType}:            _GotoState91Action,
	{_State103, AtomTypeType}:                     _GotoState86Action,
	{_State103, ReturnableTypeType}:               _GotoState92Action,
	{_State103, ValueTypeType}:                    _GotoState174Action,
	{_State103, FieldDefType}:                     _GotoState190Action,
	{_State103, ImplicitStructDefType}:            _GotoState90Action,
	{_State103, ExplicitFieldDefsType}:            _GotoState189Action,
	{_State103, OptionalExplicitFieldDefsType}:    _GotoState191Action,
	{_State103, ExplicitStructDefType}:            _GotoState45Action,
	{_State103, ImplicitEnumDefType}:              _GotoState89Action,
	{_State103, ExplicitEnumDefType}:              _GotoState87Action,
	{_State103, TraitDefType}:                     _GotoState93Action,
	{_State103, FuncTypeType}:                     _GotoState88Action,
	{_State104, IdentifierToken}:                  _GotoState80Action,
	{_State104, StructToken}:                      _GotoState34Action,
	{_State104, EnumToken}:                        _GotoState77Action,
	{_State104, TraitToken}:                       _GotoState85Action,
	{_State104, FuncToken}:                        _GotoState79Action,
	{_State104, LparenToken}:                      _GotoState81Action,
	{_State104, LbracketToken}:                    _GotoState26Action,
	{_State104, DotToken}:                         _GotoState76Action,
	{_State104, QuestionToken}:                    _GotoState83Action,
	{_State104, ExclaimToken}:                     _GotoState78Action,
	{_State104, TildeTildeToken}:                  _GotoState84Action,
	{_State104, BitNegToken}:                      _GotoState75Action,
	{_State104, BitAndToken}:                      _GotoState74Action,
	{_State104, ParseErrorToken}:                  _GotoState82Action,
	{_State104, OptionalGenericArgumentsType}:     _GotoState193Action,
	{_State104, GenericArgumentsType}:             _GotoState192Action,
	{_State104, InitializableTypeType}:            _GotoState91Action,
	{_State104, AtomTypeType}:                     _GotoState86Action,
	{_State104, ReturnableTypeType}:               _GotoState92Action,
	{_State104, ValueTypeType}:                    _GotoState194Action,
	{_State104, ImplicitStructDefType}:            _GotoState90Action,
	{_State104, ExplicitStructDefType}:            _GotoState45Action,
	{_State104, ImplicitEnumDefType}:              _GotoState89Action,
	{_State104, ExplicitEnumDefType}:              _GotoState87Action,
	{_State104, TraitDefType}:                     _GotoState93Action,
	{_State104, FuncTypeType}:                     _GotoState88Action,
	{_State105, IdentifierToken}:                  _GotoState195Action,
	{_State106, IntegerLiteralToken}:              _GotoState24Action,
	{_State106, FloatLiteralToken}:                _GotoState20Action,
	{_State106, RuneLiteralToken}:                 _GotoState32Action,
	{_State106, StringLiteralToken}:               _GotoState33Action,
	{_State106, IdentifierToken}:                  _GotoState96Action,
	{_State106, TrueToken}:                        _GotoState36Action,
	{_State106, FalseToken}:                       _GotoState19Action,
	{_State106, StructToken}:                      _GotoState34Action,
	{_State106, FuncToken}:                        _GotoState21Action,
	{_State106, VarToken}:                         _GotoState37Action,
	{_State106, LetToken}:                         _GotoState27Action,
	{_State106, NotToken}:                         _GotoState30Action,
	{_State106, LabelDeclToken}:                   _GotoState25Action,
	{_State106, LparenToken}:                      _GotoState28Action,
	{_State106, LbracketToken}:                    _GotoState26Action,
	{_State106, DotDotDotToken}:                   _GotoState95Action,
	{_State106, SubToken}:                         _GotoState35Action,
	{_State106, MulToken}:                         _GotoState29Action,
	{_State106, BitNegToken}:                      _GotoState18Action,
	{_State106, BitAndToken}:                      _GotoState17Action,
	{_State106, GreaterToken}:                     _GotoState22Action,
	{_State106, ParseErrorToken}:                  _GotoState31Action,
	{_State106, VarDeclPatternType}:               _GotoState56Action,
	{_State106, VarOrLetType}:                     _GotoState57Action,
	{_State106, ExpressionType}:                   _GotoState100Action,
	{_State106, OptionalLabelDeclType}:            _GotoState50Action,
	{_State106, SequenceExprType}:                 _GotoState55Action,
	{_State106, CallExprType}:                     _GotoState43Action,
	{_State106, ArgumentType}:                     _GotoState196Action,
	{_State106, ColonExpressionsType}:             _GotoState99Action,
	{_State106, OptionalExpressionType}:           _GotoState102Action,
	{_State106, AtomExprType}:                     _GotoState42Action,
	{_State106, LiteralExprType}:                  _GotoState48Action,
	{_State106, ImplicitStructExprType}:           _GotoState46Action,
	{_State106, AccessExprType}:                   _GotoState38Action,
	{_State106, PostfixUnaryExprType}:             _GotoState52Action,
	{_State106, PrefixUnaryOpType}:                _GotoState54Action,
	{_State106, PrefixUnaryExprType}:              _GotoState53Action,
	{_State106, MulExprType}:                      _GotoState49Action,
	{_State106, AddExprType}:                      _GotoState39Action,
	{_State106, CmpExprType}:                      _GotoState44Action,
	{_State106, AndExprType}:                      _GotoState40Action,
	{_State106, OrExprType}:                       _GotoState51Action,
	{_State106, InitializableTypeType}:            _GotoState47Action,
	{_State106, ExplicitStructDefType}:            _GotoState45Action,
	{_State106, AnonymousFuncExprType}:            _GotoState41Action,
	{_State108, LparenToken}:                      _GotoState197Action,
	{_State113, IntegerLiteralToken}:              _GotoState24Action,
	{_State113, FloatLiteralToken}:                _GotoState20Action,
	{_State113, RuneLiteralToken}:                 _GotoState32Action,
	{_State113, StringLiteralToken}:               _GotoState33Action,
	{_State113, IdentifierToken}:                  _GotoState23Action,
	{_State113, TrueToken}:                        _GotoState36Action,
	{_State113, FalseToken}:                       _GotoState19Action,
	{_State113, StructToken}:                      _GotoState34Action,
	{_State113, FuncToken}:                        _GotoState21Action,
	{_State113, NotToken}:                         _GotoState30Action,
	{_State113, LabelDeclToken}:                   _GotoState25Action,
	{_State113, LparenToken}:                      _GotoState28Action,
	{_State113, LbracketToken}:                    _GotoState26Action,
	{_State113, SubToken}:                         _GotoState35Action,
	{_State113, MulToken}:                         _GotoState29Action,
	{_State113, BitNegToken}:                      _GotoState18Action,
	{_State113, BitAndToken}:                      _GotoState17Action,
	{_State113, ParseErrorToken}:                  _GotoState31Action,
	{_State113, OptionalLabelDeclType}:            _GotoState72Action,
	{_State113, CallExprType}:                     _GotoState43Action,
	{_State113, AtomExprType}:                     _GotoState42Action,
	{_State113, LiteralExprType}:                  _GotoState48Action,
	{_State113, ImplicitStructExprType}:           _GotoState46Action,
	{_State113, AccessExprType}:                   _GotoState38Action,
	{_State113, PostfixUnaryExprType}:             _GotoState52Action,
	{_State113, PrefixUnaryOpType}:                _GotoState54Action,
	{_State113, PrefixUnaryExprType}:              _GotoState53Action,
	{_State113, MulExprType}:                      _GotoState198Action,
	{_State113, InitializableTypeType}:            _GotoState47Action,
	{_State113, ExplicitStructDefType}:            _GotoState45Action,
	{_State113, AnonymousFuncExprType}:            _GotoState41Action,
	{_State114, IntegerLiteralToken}:              _GotoState24Action,
	{_State114, FloatLiteralToken}:                _GotoState20Action,
	{_State114, RuneLiteralToken}:                 _GotoState32Action,
	{_State114, StringLiteralToken}:               _GotoState33Action,
	{_State114, IdentifierToken}:                  _GotoState23Action,
	{_State114, TrueToken}:                        _GotoState36Action,
	{_State114, FalseToken}:                       _GotoState19Action,
	{_State114, StructToken}:                      _GotoState34Action,
	{_State114, FuncToken}:                        _GotoState21Action,
	{_State114, NotToken}:                         _GotoState30Action,
	{_State114, LabelDeclToken}:                   _GotoState25Action,
	{_State114, LparenToken}:                      _GotoState28Action,
	{_State114, LbracketToken}:                    _GotoState26Action,
	{_State114, SubToken}:                         _GotoState35Action,
	{_State114, MulToken}:                         _GotoState29Action,
	{_State114, BitNegToken}:                      _GotoState18Action,
	{_State114, BitAndToken}:                      _GotoState17Action,
	{_State114, ParseErrorToken}:                  _GotoState31Action,
	{_State114, OptionalLabelDeclType}:            _GotoState72Action,
	{_State114, CallExprType}:                     _GotoState43Action,
	{_State114, AtomExprType}:                     _GotoState42Action,
	{_State114, LiteralExprType}:                  _GotoState48Action,
	{_State114, ImplicitStructExprType}:           _GotoState46Action,
	{_State114, AccessExprType}:                   _GotoState38Action,
	{_State114, PostfixUnaryExprType}:             _GotoState52Action,
	{_State114, PrefixUnaryOpType}:                _GotoState54Action,
	{_State114, PrefixUnaryExprType}:              _GotoState53Action,
	{_State114, MulExprType}:                      _GotoState49Action,
	{_State114, AddExprType}:                      _GotoState39Action,
	{_State114, CmpExprType}:                      _GotoState199Action,
	{_State114, InitializableTypeType}:            _GotoState47Action,
	{_State114, ExplicitStructDefType}:            _GotoState45Action,
	{_State114, AnonymousFuncExprType}:            _GotoState41Action,
	{_State121, IntegerLiteralToken}:              _GotoState24Action,
	{_State121, FloatLiteralToken}:                _GotoState20Action,
	{_State121, RuneLiteralToken}:                 _GotoState32Action,
	{_State121, StringLiteralToken}:               _GotoState33Action,
	{_State121, IdentifierToken}:                  _GotoState23Action,
	{_State121, TrueToken}:                        _GotoState36Action,
	{_State121, FalseToken}:                       _GotoState19Action,
	{_State121, StructToken}:                      _GotoState34Action,
	{_State121, FuncToken}:                        _GotoState21Action,
	{_State121, NotToken}:                         _GotoState30Action,
	{_State121, LabelDeclToken}:                   _GotoState25Action,
	{_State121, LparenToken}:                      _GotoState28Action,
	{_State121, LbracketToken}:                    _GotoState26Action,
	{_State121, SubToken}:                         _GotoState35Action,
	{_State121, MulToken}:                         _GotoState29Action,
	{_State121, BitNegToken}:                      _GotoState18Action,
	{_State121, BitAndToken}:                      _GotoState17Action,
	{_State121, ParseErrorToken}:                  _GotoState31Action,
	{_State121, OptionalLabelDeclType}:            _GotoState72Action,
	{_State121, CallExprType}:                     _GotoState43Action,
	{_State121, AtomExprType}:                     _GotoState42Action,
	{_State121, LiteralExprType}:                  _GotoState48Action,
	{_State121, ImplicitStructExprType}:           _GotoState46Action,
	{_State121, AccessExprType}:                   _GotoState38Action,
	{_State121, PostfixUnaryExprType}:             _GotoState52Action,
	{_State121, PrefixUnaryOpType}:                _GotoState54Action,
	{_State121, PrefixUnaryExprType}:              _GotoState53Action,
	{_State121, MulExprType}:                      _GotoState49Action,
	{_State121, AddExprType}:                      _GotoState200Action,
	{_State121, InitializableTypeType}:            _GotoState47Action,
	{_State121, ExplicitStructDefType}:            _GotoState45Action,
	{_State121, AnonymousFuncExprType}:            _GotoState41Action,
	{_State122, IntegerLiteralToken}:              _GotoState24Action,
	{_State122, FloatLiteralToken}:                _GotoState20Action,
	{_State122, RuneLiteralToken}:                 _GotoState32Action,
	{_State122, StringLiteralToken}:               _GotoState33Action,
	{_State122, IdentifierToken}:                  _GotoState96Action,
	{_State122, TrueToken}:                        _GotoState36Action,
	{_State122, FalseToken}:                       _GotoState19Action,
	{_State122, StructToken}:                      _GotoState34Action,
	{_State122, FuncToken}:                        _GotoState21Action,
	{_State122, VarToken}:                         _GotoState37Action,
	{_State122, LetToken}:                         _GotoState27Action,
	{_State122, NotToken}:                         _GotoState30Action,
	{_State122, LabelDeclToken}:                   _GotoState25Action,
	{_State122, LparenToken}:                      _GotoState28Action,
	{_State122, LbracketToken}:                    _GotoState26Action,
	{_State122, DotDotDotToken}:                   _GotoState95Action,
	{_State122, SubToken}:                         _GotoState35Action,
	{_State122, MulToken}:                         _GotoState29Action,
	{_State122, BitNegToken}:                      _GotoState18Action,
	{_State122, BitAndToken}:                      _GotoState17Action,
	{_State122, GreaterToken}:                     _GotoState22Action,
	{_State122, ParseErrorToken}:                  _GotoState31Action,
	{_State122, VarDeclPatternType}:               _GotoState56Action,
	{_State122, VarOrLetType}:                     _GotoState57Action,
	{_State122, ExpressionType}:                   _GotoState100Action,
	{_State122, OptionalLabelDeclType}:            _GotoState50Action,
	{_State122, SequenceExprType}:                 _GotoState55Action,
	{_State122, CallExprType}:                     _GotoState43Action,
	{_State122, OptionalArgumentsType}:            _GotoState201Action,
	{_State122, ArgumentsType}:                    _GotoState98Action,
	{_State122, ArgumentType}:                     _GotoState97Action,
	{_State122, ColonExpressionsType}:             _GotoState99Action,
	{_State122, OptionalExpressionType}:           _GotoState102Action,
	{_State122, AtomExprType}:                     _GotoState42Action,
	{_State122, LiteralExprType}:                  _GotoState48Action,
	{_State122, ImplicitStructExprType}:           _GotoState46Action,
	{_State122, AccessExprType}:                   _GotoState38Action,
	{_State122, PostfixUnaryExprType}:             _GotoState52Action,
	{_State122, PrefixUnaryOpType}:                _GotoState54Action,
	{_State122, PrefixUnaryExprType}:              _GotoState53Action,
	{_State122, MulExprType}:                      _GotoState49Action,
	{_State122, AddExprType}:                      _GotoState39Action,
	{_State122, CmpExprType}:                      _GotoState44Action,
	{_State122, AndExprType}:                      _GotoState40Action,
	{_State122, OrExprType}:                       _GotoState51Action,
	{_State122, InitializableTypeType}:            _GotoState47Action,
	{_State122, ExplicitStructDefType}:            _GotoState45Action,
	{_State122, AnonymousFuncExprType}:            _GotoState41Action,
	{_State129, IntegerLiteralToken}:              _GotoState24Action,
	{_State129, FloatLiteralToken}:                _GotoState20Action,
	{_State129, RuneLiteralToken}:                 _GotoState32Action,
	{_State129, StringLiteralToken}:               _GotoState33Action,
	{_State129, IdentifierToken}:                  _GotoState23Action,
	{_State129, TrueToken}:                        _GotoState36Action,
	{_State129, FalseToken}:                       _GotoState19Action,
	{_State129, StructToken}:                      _GotoState34Action,
	{_State129, FuncToken}:                        _GotoState21Action,
	{_State129, NotToken}:                         _GotoState30Action,
	{_State129, LabelDeclToken}:                   _GotoState25Action,
	{_State129, LparenToken}:                      _GotoState28Action,
	{_State129, LbracketToken}:                    _GotoState26Action,
	{_State129, SubToken}:                         _GotoState35Action,
	{_State129, MulToken}:                         _GotoState29Action,
	{_State129, BitNegToken}:                      _GotoState18Action,
	{_State129, BitAndToken}:                      _GotoState17Action,
	{_State129, ParseErrorToken}:                  _GotoState31Action,
	{_State129, OptionalLabelDeclType}:            _GotoState72Action,
	{_State129, CallExprType}:                     _GotoState43Action,
	{_State129, AtomExprType}:                     _GotoState42Action,
	{_State129, LiteralExprType}:                  _GotoState48Action,
	{_State129, ImplicitStructExprType}:           _GotoState46Action,
	{_State129, AccessExprType}:                   _GotoState38Action,
	{_State129, PostfixUnaryExprType}:             _GotoState52Action,
	{_State129, PrefixUnaryOpType}:                _GotoState54Action,
	{_State129, PrefixUnaryExprType}:              _GotoState202Action,
	{_State129, InitializableTypeType}:            _GotoState47Action,
	{_State129, ExplicitStructDefType}:            _GotoState45Action,
	{_State129, AnonymousFuncExprType}:            _GotoState41Action,
	{_State130, LbraceToken}:                      _GotoState59Action,
	{_State130, StatementBlockType}:               _GotoState203Action,
	{_State131, IntegerLiteralToken}:              _GotoState24Action,
	{_State131, FloatLiteralToken}:                _GotoState20Action,
	{_State131, RuneLiteralToken}:                 _GotoState32Action,
	{_State131, StringLiteralToken}:               _GotoState33Action,
	{_State131, IdentifierToken}:                  _GotoState23Action,
	{_State131, TrueToken}:                        _GotoState36Action,
	{_State131, FalseToken}:                       _GotoState19Action,
	{_State131, StructToken}:                      _GotoState34Action,
	{_State131, FuncToken}:                        _GotoState21Action,
	{_State131, VarToken}:                         _GotoState37Action,
	{_State131, LetToken}:                         _GotoState27Action,
	{_State131, NotToken}:                         _GotoState30Action,
	{_State131, LabelDeclToken}:                   _GotoState25Action,
	{_State131, LparenToken}:                      _GotoState28Action,
	{_State131, LbracketToken}:                    _GotoState26Action,
	{_State131, SubToken}:                         _GotoState35Action,
	{_State131, MulToken}:                         _GotoState29Action,
	{_State131, BitNegToken}:                      _GotoState18Action,
	{_State131, BitAndToken}:                      _GotoState17Action,
	{_State131, GreaterToken}:                     _GotoState22Action,
	{_State131, ParseErrorToken}:                  _GotoState31Action,
	{_State131, VarDeclPatternType}:               _GotoState56Action,
	{_State131, VarOrLetType}:                     _GotoState57Action,
	{_State131, AssignPatternType}:                _GotoState204Action,
	{_State131, OptionalLabelDeclType}:            _GotoState72Action,
	{_State131, SequenceExprType}:                 _GotoState206Action,
	{_State131, ForAssignmentType}:                _GotoState205Action,
	{_State131, CallExprType}:                     _GotoState43Action,
	{_State131, AtomExprType}:                     _GotoState42Action,
	{_State131, LiteralExprType}:                  _GotoState48Action,
	{_State131, ImplicitStructExprType}:           _GotoState46Action,
	{_State131, AccessExprType}:                   _GotoState38Action,
	{_State131, PostfixUnaryExprType}:             _GotoState52Action,
	{_State131, PrefixUnaryOpType}:                _GotoState54Action,
	{_State131, PrefixUnaryExprType}:              _GotoState53Action,
	{_State131, MulExprType}:                      _GotoState49Action,
	{_State131, AddExprType}:                      _GotoState39Action,
	{_State131, CmpExprType}:                      _GotoState44Action,
	{_State131, AndExprType}:                      _GotoState40Action,
	{_State131, OrExprType}:                       _GotoState51Action,
	{_State131, InitializableTypeType}:            _GotoState47Action,
	{_State131, ExplicitStructDefType}:            _GotoState45Action,
	{_State131, AnonymousFuncExprType}:            _GotoState41Action,
	{_State132, IntegerLiteralToken}:              _GotoState24Action,
	{_State132, FloatLiteralToken}:                _GotoState20Action,
	{_State132, RuneLiteralToken}:                 _GotoState32Action,
	{_State132, StringLiteralToken}:               _GotoState33Action,
	{_State132, IdentifierToken}:                  _GotoState23Action,
	{_State132, TrueToken}:                        _GotoState36Action,
	{_State132, FalseToken}:                       _GotoState19Action,
	{_State132, CaseToken}:                        _GotoState207Action,
	{_State132, StructToken}:                      _GotoState34Action,
	{_State132, FuncToken}:                        _GotoState21Action,
	{_State132, VarToken}:                         _GotoState37Action,
	{_State132, LetToken}:                         _GotoState27Action,
	{_State132, NotToken}:                         _GotoState30Action,
	{_State132, LabelDeclToken}:                   _GotoState25Action,
	{_State132, LparenToken}:                      _GotoState28Action,
	{_State132, LbracketToken}:                    _GotoState26Action,
	{_State132, SubToken}:                         _GotoState35Action,
	{_State132, MulToken}:                         _GotoState29Action,
	{_State132, BitNegToken}:                      _GotoState18Action,
	{_State132, BitAndToken}:                      _GotoState17Action,
	{_State132, GreaterToken}:                     _GotoState22Action,
	{_State132, ParseErrorToken}:                  _GotoState31Action,
	{_State132, VarDeclPatternType}:               _GotoState56Action,
	{_State132, VarOrLetType}:                     _GotoState57Action,
	{_State132, OptionalLabelDeclType}:            _GotoState72Action,
	{_State132, SequenceExprType}:                 _GotoState209Action,
	{_State132, ConditionType}:                    _GotoState208Action,
	{_State132, CallExprType}:                     _GotoState43Action,
	{_State132, AtomExprType}:                     _GotoState42Action,
	{_State132, LiteralExprType}:                  _GotoState48Action,
	{_State132, ImplicitStructExprType}:           _GotoState46Action,
	{_State132, AccessExprType}:                   _GotoState38Action,
	{_State132, PostfixUnaryExprType}:             _GotoState52Action,
	{_State132, PrefixUnaryOpType}:                _GotoState54Action,
	{_State132, PrefixUnaryExprType}:              _GotoState53Action,
	{_State132, MulExprType}:                      _GotoState49Action,
	{_State132, AddExprType}:                      _GotoState39Action,
	{_State132, CmpExprType}:                      _GotoState44Action,
	{_State132, AndExprType}:                      _GotoState40Action,
	{_State132, OrExprType}:                       _GotoState51Action,
	{_State132, InitializableTypeType}:            _GotoState47Action,
	{_State132, ExplicitStructDefType}:            _GotoState45Action,
	{_State132, AnonymousFuncExprType}:            _GotoState41Action,
	{_State133, IntegerLiteralToken}:              _GotoState24Action,
	{_State133, FloatLiteralToken}:                _GotoState20Action,
	{_State133, RuneLiteralToken}:                 _GotoState32Action,
	{_State133, StringLiteralToken}:               _GotoState33Action,
	{_State133, IdentifierToken}:                  _GotoState23Action,
	{_State133, TrueToken}:                        _GotoState36Action,
	{_State133, FalseToken}:                       _GotoState19Action,
	{_State133, StructToken}:                      _GotoState34Action,
	{_State133, FuncToken}:                        _GotoState21Action,
	{_State133, VarToken}:                         _GotoState37Action,
	{_State133, LetToken}:                         _GotoState27Action,
	{_State133, NotToken}:                         _GotoState30Action,
	{_State133, LabelDeclToken}:                   _GotoState25Action,
	{_State133, LparenToken}:                      _GotoState28Action,
	{_State133, LbracketToken}:                    _GotoState26Action,
	{_State133, SubToken}:                         _GotoState35Action,
	{_State133, MulToken}:                         _GotoState29Action,
	{_State133, BitNegToken}:                      _GotoState18Action,
	{_State133, BitAndToken}:                      _GotoState17Action,
	{_State133, GreaterToken}:                     _GotoState22Action,
	{_State133, ParseErrorToken}:                  _GotoState31Action,
	{_State133, VarDeclPatternType}:               _GotoState56Action,
	{_State133, VarOrLetType}:                     _GotoState57Action,
	{_State133, OptionalLabelDeclType}:            _GotoState72Action,
	{_State133, SequenceExprType}:                 _GotoState210Action,
	{_State133, CallExprType}:                     _GotoState43Action,
	{_State133, AtomExprType}:                     _GotoState42Action,
	{_State133, LiteralExprType}:                  _GotoState48Action,
	{_State133, ImplicitStructExprType}:           _GotoState46Action,
	{_State133, AccessExprType}:                   _GotoState38Action,
	{_State133, PostfixUnaryExprType}:             _GotoState52Action,
	{_State133, PrefixUnaryOpType}:                _GotoState54Action,
	{_State133, PrefixUnaryExprType}:              _GotoState53Action,
	{_State133, MulExprType}:                      _GotoState49Action,
	{_State133, AddExprType}:                      _GotoState39Action,
	{_State133, CmpExprType}:                      _GotoState44Action,
	{_State133, AndExprType}:                      _GotoState40Action,
	{_State133, OrExprType}:                       _GotoState51Action,
	{_State133, InitializableTypeType}:            _GotoState47Action,
	{_State133, ExplicitStructDefType}:            _GotoState45Action,
	{_State133, AnonymousFuncExprType}:            _GotoState41Action,
	{_State138, IntegerLiteralToken}:              _GotoState24Action,
	{_State138, FloatLiteralToken}:                _GotoState20Action,
	{_State138, RuneLiteralToken}:                 _GotoState32Action,
	{_State138, StringLiteralToken}:               _GotoState33Action,
	{_State138, IdentifierToken}:                  _GotoState23Action,
	{_State138, TrueToken}:                        _GotoState36Action,
	{_State138, FalseToken}:                       _GotoState19Action,
	{_State138, StructToken}:                      _GotoState34Action,
	{_State138, FuncToken}:                        _GotoState21Action,
	{_State138, NotToken}:                         _GotoState30Action,
	{_State138, LabelDeclToken}:                   _GotoState25Action,
	{_State138, LparenToken}:                      _GotoState28Action,
	{_State138, LbracketToken}:                    _GotoState26Action,
	{_State138, SubToken}:                         _GotoState35Action,
	{_State138, MulToken}:                         _GotoState29Action,
	{_State138, BitNegToken}:                      _GotoState18Action,
	{_State138, BitAndToken}:                      _GotoState17Action,
	{_State138, ParseErrorToken}:                  _GotoState31Action,
	{_State138, OptionalLabelDeclType}:            _GotoState72Action,
	{_State138, CallExprType}:                     _GotoState43Action,
	{_State138, AtomExprType}:                     _GotoState42Action,
	{_State138, LiteralExprType}:                  _GotoState48Action,
	{_State138, ImplicitStructExprType}:           _GotoState46Action,
	{_State138, AccessExprType}:                   _GotoState38Action,
	{_State138, PostfixUnaryExprType}:             _GotoState52Action,
	{_State138, PrefixUnaryOpType}:                _GotoState54Action,
	{_State138, PrefixUnaryExprType}:              _GotoState53Action,
	{_State138, MulExprType}:                      _GotoState49Action,
	{_State138, AddExprType}:                      _GotoState39Action,
	{_State138, CmpExprType}:                      _GotoState44Action,
	{_State138, AndExprType}:                      _GotoState211Action,
	{_State138, InitializableTypeType}:            _GotoState47Action,
	{_State138, ExplicitStructDefType}:            _GotoState45Action,
	{_State138, AnonymousFuncExprType}:            _GotoState41Action,
	{_State141, IdentifierToken}:                  _GotoState213Action,
	{_State141, LparenToken}:                      _GotoState141Action,
	{_State141, DotDotDotToken}:                   _GotoState212Action,
	{_State141, VarPatternType}:                   _GotoState216Action,
	{_State141, TuplePatternType}:                 _GotoState142Action,
	{_State141, FieldVarPatternsType}:             _GotoState215Action,
	{_State141, FieldVarPatternType}:              _GotoState214Action,
	{_State143, IdentifierToken}:                  _GotoState80Action,
	{_State143, StructToken}:                      _GotoState34Action,
	{_State143, EnumToken}:                        _GotoState77Action,
	{_State143, TraitToken}:                       _GotoState85Action,
	{_State143, FuncToken}:                        _GotoState79Action,
	{_State143, LparenToken}:                      _GotoState81Action,
	{_State143, LbracketToken}:                    _GotoState26Action,
	{_State143, DotToken}:                         _GotoState76Action,
	{_State143, QuestionToken}:                    _GotoState83Action,
	{_State143, ExclaimToken}:                     _GotoState78Action,
	{_State143, TildeTildeToken}:                  _GotoState84Action,
	{_State143, BitNegToken}:                      _GotoState75Action,
	{_State143, BitAndToken}:                      _GotoState74Action,
	{_State143, ParseErrorToken}:                  _GotoState82Action,
	{_State143, OptionalValueTypeType}:            _GotoState217Action,
	{_State143, InitializableTypeType}:            _GotoState91Action,
	{_State143, AtomTypeType}:                     _GotoState86Action,
	{_State143, ReturnableTypeType}:               _GotoState92Action,
	{_State143, ValueTypeType}:                    _GotoState218Action,
	{_State143, ImplicitStructDefType}:            _GotoState90Action,
	{_State143, ExplicitStructDefType}:            _GotoState45Action,
	{_State143, ImplicitEnumDefType}:              _GotoState89Action,
	{_State143, ExplicitEnumDefType}:              _GotoState87Action,
	{_State143, TraitDefType}:                     _GotoState93Action,
	{_State143, FuncTypeType}:                     _GotoState88Action,
	{_State144, IntegerLiteralToken}:              _GotoState24Action,
	{_State144, FloatLiteralToken}:                _GotoState20Action,
	{_State144, RuneLiteralToken}:                 _GotoState32Action,
	{_State144, StringLiteralToken}:               _GotoState33Action,
	{_State144, IdentifierToken}:                  _GotoState23Action,
	{_State144, TrueToken}:                        _GotoState36Action,
	{_State144, FalseToken}:                       _GotoState19Action,
	{_State144, CaseToken}:                        _GotoState221Action,
	{_State144, DefaultToken}:                     _GotoState223Action,
	{_State144, ReturnToken}:                      _GotoState228Action,
	{_State144, BreakToken}:                       _GotoState220Action,
	{_State144, ContinueToken}:                    _GotoState222Action,
	{_State144, FallthroughToken}:                 _GotoState225Action,
	{_State144, ImportToken}:                      _GotoState226Action,
	{_State144, UnsafeToken}:                      _GotoState167Action,
	{_State144, StructToken}:                      _GotoState34Action,
	{_State144, FuncToken}:                        _GotoState21Action,
	{_State144, AsyncToken}:                       _GotoState219Action,
	{_State144, DeferToken}:                       _GotoState224Action,
	{_State144, VarToken}:                         _GotoState37Action,
	{_State144, LetToken}:                         _GotoState27Action,
	{_State144, NotToken}:                         _GotoState30Action,
	{_State144, LabelDeclToken}:                   _GotoState25Action,
	{_State144, RbraceToken}:                      _GotoState227Action,
	{_State144, LparenToken}:                      _GotoState28Action,
	{_State144, LbracketToken}:                    _GotoState26Action,
	{_State144, SubToken}:                         _GotoState35Action,
	{_State144, MulToken}:                         _GotoState29Action,
	{_State144, BitNegToken}:                      _GotoState18Action,
	{_State144, BitAndToken}:                      _GotoState17Action,
	{_State144, GreaterToken}:                     _GotoState22Action,
	{_State144, ParseErrorToken}:                  _GotoState31Action,
	{_State144, StatementType}:                    _GotoState237Action,
	{_State144, SimpleStatementBodyType}:          _GotoState236Action,
	{_State144, StatementBodyType}:                _GotoState238Action,
	{_State144, UnsafeStatementType}:              _GotoState239Action,
	{_State144, JumpTypeType}:                     _GotoState234Action,
	{_State144, ExpressionsType}:                  _GotoState232Action,
	{_State144, ImportStatementType}:              _GotoState233Action,
	{_State144, VarDeclPatternType}:               _GotoState56Action,
	{_State144, VarOrLetType}:                     _GotoState57Action,
	{_State144, AssignPatternType}:                _GotoState230Action,
	{_State144, ExpressionType}:                   _GotoState231Action,
	{_State144, OptionalLabelDeclType}:            _GotoState50Action,
	{_State144, SequenceExprType}:                 _GotoState235Action,
	{_State144, CallExprType}:                     _GotoState43Action,
	{_State144, AtomExprType}:                     _GotoState42Action,
	{_State144, LiteralExprType}:                  _GotoState48Action,
	{_State144, ImplicitStructExprType}:           _GotoState46Action,
	{_State144, AccessExprType}:                   _GotoState229Action,
	{_State144, PostfixUnaryExprType}:             _GotoState52Action,
	{_State144, PrefixUnaryOpType}:                _GotoState54Action,
	{_State144, PrefixUnaryExprType}:              _GotoState53Action,
	{_State144, MulExprType}:                      _GotoState49Action,
	{_State144, AddExprType}:                      _GotoState39Action,
	{_State144, CmpExprType}:                      _GotoState44Action,
	{_State144, AndExprType}:                      _GotoState40Action,
	{_State144, OrExprType}:                       _GotoState51Action,
	{_State144, InitializableTypeType}:            _GotoState47Action,
	{_State144, ExplicitStructDefType}:            _GotoState45Action,
	{_State144, AnonymousFuncExprType}:            _GotoState41Action,
	{_State145, CommentGroupsToken}:               _GotoState58Action,
	{_State145, PackageToken}:                     _GotoState14Action,
	{_State145, TypeToken}:                        _GotoState15Action,
	{_State145, FuncToken}:                        _GotoState16Action,
	{_State145, VarToken}:                         _GotoState37Action,
	{_State145, LetToken}:                         _GotoState27Action,
	{_State145, LbraceToken}:                      _GotoState59Action,
	{_State145, DefinitionType}:                   _GotoState240Action,
	{_State145, StatementBlockType}:               _GotoState64Action,
	{_State145, VarDeclPatternType}:               _GotoState66Action,
	{_State145, VarOrLetType}:                     _GotoState57Action,
	{_State145, TypeDefType}:                      _GotoState65Action,
	{_State145, NamedFuncDefType}:                 _GotoState62Action,
	{_State145, PackageDefType}:                   _GotoState63Action,
	{_State147, IntegerLiteralToken}:              _GotoState24Action,
	{_State147, FloatLiteralToken}:                _GotoState20Action,
	{_State147, RuneLiteralToken}:                 _GotoState32Action,
	{_State147, StringLiteralToken}:               _GotoState33Action,
	{_State147, IdentifierToken}:                  _GotoState23Action,
	{_State147, TrueToken}:                        _GotoState36Action,
	{_State147, FalseToken}:                       _GotoState19Action,
	{_State147, StructToken}:                      _GotoState34Action,
	{_State147, FuncToken}:                        _GotoState21Action,
	{_State147, VarToken}:                         _GotoState37Action,
	{_State147, LetToken}:                         _GotoState27Action,
	{_State147, NotToken}:                         _GotoState30Action,
	{_State147, LabelDeclToken}:                   _GotoState25Action,
	{_State147, LparenToken}:                      _GotoState28Action,
	{_State147, LbracketToken}:                    _GotoState26Action,
	{_State147, SubToken}:                         _GotoState35Action,
	{_State147, MulToken}:                         _GotoState29Action,
	{_State147, BitNegToken}:                      _GotoState18Action,
	{_State147, BitAndToken}:                      _GotoState17Action,
	{_State147, GreaterToken}:                     _GotoState22Action,
	{_State147, ParseErrorToken}:                  _GotoState31Action,
	{_State147, VarDeclPatternType}:               _GotoState56Action,
	{_State147, VarOrLetType}:                     _GotoState57Action,
	{_State147, ExpressionType}:                   _GotoState241Action,
	{_State147, OptionalLabelDeclType}:            _GotoState50Action,
	{_State147, SequenceExprType}:                 _GotoState55Action,
	{_State147, CallExprType}:                     _GotoState43Action,
	{_State147, AtomExprType}:                     _GotoState42Action,
	{_State147, LiteralExprType}:                  _GotoState48Action,
	{_State147, ImplicitStructExprType}:           _GotoState46Action,
	{_State147, AccessExprType}:                   _GotoState38Action,
	{_State147, PostfixUnaryExprType}:             _GotoState52Action,
	{_State147, PrefixUnaryOpType}:                _GotoState54Action,
	{_State147, PrefixUnaryExprType}:              _GotoState53Action,
	{_State147, MulExprType}:                      _GotoState49Action,
	{_State147, AddExprType}:                      _GotoState39Action,
	{_State147, CmpExprType}:                      _GotoState44Action,
	{_State147, AndExprType}:                      _GotoState40Action,
	{_State147, OrExprType}:                       _GotoState51Action,
	{_State147, InitializableTypeType}:            _GotoState47Action,
	{_State147, ExplicitStructDefType}:            _GotoState45Action,
	{_State147, AnonymousFuncExprType}:            _GotoState41Action,
	{_State148, IdentifierToken}:                  _GotoState80Action,
	{_State148, StructToken}:                      _GotoState34Action,
	{_State148, EnumToken}:                        _GotoState77Action,
	{_State148, TraitToken}:                       _GotoState85Action,
	{_State148, FuncToken}:                        _GotoState79Action,
	{_State148, LparenToken}:                      _GotoState81Action,
	{_State148, LbracketToken}:                    _GotoState26Action,
	{_State148, DotToken}:                         _GotoState76Action,
	{_State148, QuestionToken}:                    _GotoState83Action,
	{_State148, ExclaimToken}:                     _GotoState78Action,
	{_State148, TildeTildeToken}:                  _GotoState84Action,
	{_State148, BitNegToken}:                      _GotoState75Action,
	{_State148, BitAndToken}:                      _GotoState74Action,
	{_State148, ParseErrorToken}:                  _GotoState82Action,
	{_State148, InitializableTypeType}:            _GotoState91Action,
	{_State148, AtomTypeType}:                     _GotoState86Action,
	{_State148, ReturnableTypeType}:               _GotoState92Action,
	{_State148, ValueTypeType}:                    _GotoState242Action,
	{_State148, ImplicitStructDefType}:            _GotoState90Action,
	{_State148, ExplicitStructDefType}:            _GotoState45Action,
	{_State148, ImplicitEnumDefType}:              _GotoState89Action,
	{_State148, ExplicitEnumDefType}:              _GotoState87Action,
	{_State148, TraitDefType}:                     _GotoState93Action,
	{_State148, FuncTypeType}:                     _GotoState88Action,
	{_State149, IdentifierToken}:                  _GotoState243Action,
	{_State149, GenericParameterDefType}:          _GotoState244Action,
	{_State149, GenericParameterDefsType}:         _GotoState245Action,
	{_State149, OptionalGenericParameterDefsType}: _GotoState246Action,
	{_State150, IdentifierToken}:                  _GotoState80Action,
	{_State150, StructToken}:                      _GotoState34Action,
	{_State150, EnumToken}:                        _GotoState77Action,
	{_State150, TraitToken}:                       _GotoState85Action,
	{_State150, FuncToken}:                        _GotoState79Action,
	{_State150, LparenToken}:                      _GotoState81Action,
	{_State150, LbracketToken}:                    _GotoState26Action,
	{_State150, DotToken}:                         _GotoState76Action,
	{_State150, QuestionToken}:                    _GotoState83Action,
	{_State150, ExclaimToken}:                     _GotoState78Action,
	{_State150, TildeTildeToken}:                  _GotoState84Action,
	{_State150, BitNegToken}:                      _GotoState75Action,
	{_State150, BitAndToken}:                      _GotoState74Action,
	{_State150, ParseErrorToken}:                  _GotoState82Action,
	{_State150, InitializableTypeType}:            _GotoState91Action,
	{_State150, AtomTypeType}:                     _GotoState86Action,
	{_State150, ReturnableTypeType}:               _GotoState92Action,
	{_State150, ValueTypeType}:                    _GotoState247Action,
	{_State150, ImplicitStructDefType}:            _GotoState90Action,
	{_State150, ExplicitStructDefType}:            _GotoState45Action,
	{_State150, ImplicitEnumDefType}:              _GotoState89Action,
	{_State150, ExplicitEnumDefType}:              _GotoState87Action,
	{_State150, TraitDefType}:                     _GotoState93Action,
	{_State150, FuncTypeType}:                     _GotoState88Action,
	{_State151, IntegerLiteralToken}:              _GotoState24Action,
	{_State151, FloatLiteralToken}:                _GotoState20Action,
	{_State151, RuneLiteralToken}:                 _GotoState32Action,
	{_State151, StringLiteralToken}:               _GotoState33Action,
	{_State151, IdentifierToken}:                  _GotoState23Action,
	{_State151, TrueToken}:                        _GotoState36Action,
	{_State151, FalseToken}:                       _GotoState19Action,
	{_State151, StructToken}:                      _GotoState34Action,
	{_State151, FuncToken}:                        _GotoState21Action,
	{_State151, VarToken}:                         _GotoState37Action,
	{_State151, LetToken}:                         _GotoState27Action,
	{_State151, NotToken}:                         _GotoState30Action,
	{_State151, LabelDeclToken}:                   _GotoState25Action,
	{_State151, LparenToken}:                      _GotoState28Action,
	{_State151, LbracketToken}:                    _GotoState26Action,
	{_State151, SubToken}:                         _GotoState35Action,
	{_State151, MulToken}:                         _GotoState29Action,
	{_State151, BitNegToken}:                      _GotoState18Action,
	{_State151, BitAndToken}:                      _GotoState17Action,
	{_State151, GreaterToken}:                     _GotoState22Action,
	{_State151, ParseErrorToken}:                  _GotoState31Action,
	{_State151, VarDeclPatternType}:               _GotoState56Action,
	{_State151, VarOrLetType}:                     _GotoState57Action,
	{_State151, ExpressionType}:                   _GotoState248Action,
	{_State151, OptionalLabelDeclType}:            _GotoState50Action,
	{_State151, SequenceExprType}:                 _GotoState55Action,
	{_State151, CallExprType}:                     _GotoState43Action,
	{_State151, AtomExprType}:                     _GotoState42Action,
	{_State151, LiteralExprType}:                  _GotoState48Action,
	{_State151, ImplicitStructExprType}:           _GotoState46Action,
	{_State151, AccessExprType}:                   _GotoState38Action,
	{_State151, PostfixUnaryExprType}:             _GotoState52Action,
	{_State151, PrefixUnaryOpType}:                _GotoState54Action,
	{_State151, PrefixUnaryExprType}:              _GotoState53Action,
	{_State151, MulExprType}:                      _GotoState49Action,
	{_State151, AddExprType}:                      _GotoState39Action,
	{_State151, CmpExprType}:                      _GotoState44Action,
	{_State151, AndExprType}:                      _GotoState40Action,
	{_State151, OrExprType}:                       _GotoState51Action,
	{_State151, InitializableTypeType}:            _GotoState47Action,
	{_State151, ExplicitStructDefType}:            _GotoState45Action,
	{_State151, AnonymousFuncExprType}:            _GotoState41Action,
	{_State152, LparenToken}:                      _GotoState249Action,
	{_State153, IdentifierToken}:                  _GotoState80Action,
	{_State153, StructToken}:                      _GotoState34Action,
	{_State153, EnumToken}:                        _GotoState77Action,
	{_State153, TraitToken}:                       _GotoState85Action,
	{_State153, FuncToken}:                        _GotoState79Action,
	{_State153, LparenToken}:                      _GotoState81Action,
	{_State153, LbracketToken}:                    _GotoState26Action,
	{_State153, DotToken}:                         _GotoState76Action,
	{_State153, QuestionToken}:                    _GotoState83Action,
	{_State153, ExclaimToken}:                     _GotoState78Action,
	{_State153, DotDotDotToken}:                   _GotoState250Action,
	{_State153, TildeTildeToken}:                  _GotoState84Action,
	{_State153, BitNegToken}:                      _GotoState75Action,
	{_State153, BitAndToken}:                      _GotoState74Action,
	{_State153, ParseErrorToken}:                  _GotoState82Action,
	{_State153, InitializableTypeType}:            _GotoState91Action,
	{_State153, AtomTypeType}:                     _GotoState86Action,
	{_State153, ReturnableTypeType}:               _GotoState92Action,
	{_State153, ValueTypeType}:                    _GotoState251Action,
	{_State153, ImplicitStructDefType}:            _GotoState90Action,
	{_State153, ExplicitStructDefType}:            _GotoState45Action,
	{_State153, ImplicitEnumDefType}:              _GotoState89Action,
	{_State153, ExplicitEnumDefType}:              _GotoState87Action,
	{_State153, TraitDefType}:                     _GotoState93Action,
	{_State153, FuncTypeType}:                     _GotoState88Action,
	{_State154, RparenToken}:                      _GotoState252Action,
	{_State155, RparenToken}:                      _GotoState253Action,
	{_State157, CommaToken}:                       _GotoState254Action,
	{_State161, IdentifierToken}:                  _GotoState166Action,
	{_State161, UnsafeToken}:                      _GotoState167Action,
	{_State161, StructToken}:                      _GotoState34Action,
	{_State161, EnumToken}:                        _GotoState77Action,
	{_State161, TraitToken}:                       _GotoState85Action,
	{_State161, FuncToken}:                        _GotoState79Action,
	{_State161, LparenToken}:                      _GotoState81Action,
	{_State161, LbracketToken}:                    _GotoState26Action,
	{_State161, DotToken}:                         _GotoState76Action,
	{_State161, QuestionToken}:                    _GotoState83Action,
	{_State161, ExclaimToken}:                     _GotoState78Action,
	{_State161, TildeTildeToken}:                  _GotoState84Action,
	{_State161, BitNegToken}:                      _GotoState75Action,
	{_State161, BitAndToken}:                      _GotoState74Action,
	{_State161, ParseErrorToken}:                  _GotoState82Action,
	{_State161, UnsafeStatementType}:              _GotoState173Action,
	{_State161, InitializableTypeType}:            _GotoState91Action,
	{_State161, AtomTypeType}:                     _GotoState86Action,
	{_State161, ReturnableTypeType}:               _GotoState92Action,
	{_State161, ValueTypeType}:                    _GotoState174Action,
	{_State161, FieldDefType}:                     _GotoState257Action,
	{_State161, ImplicitStructDefType}:            _GotoState90Action,
	{_State161, ExplicitStructDefType}:            _GotoState45Action,
	{_State161, EnumValueDefType}:                 _GotoState255Action,
	{_State161, ImplicitEnumValueDefsType}:        _GotoState258Action,
	{_State161, ImplicitEnumDefType}:              _GotoState89Action,
	{_State161, ExplicitEnumValueDefsType}:        _GotoState256Action,
	{_State161, ExplicitEnumDefType}:              _GotoState87Action,
	{_State161, TraitDefType}:                     _GotoState93Action,
	{_State161, FuncTypeType}:                     _GotoState88Action,
	{_State163, IdentifierToken}:                  _GotoState260Action,
	{_State163, StructToken}:                      _GotoState34Action,
	{_State163, EnumToken}:                        _GotoState77Action,
	{_State163, TraitToken}:                       _GotoState85Action,
	{_State163, FuncToken}:                        _GotoState79Action,
	{_State163, LparenToken}:                      _GotoState81Action,
	{_State163, LbracketToken}:                    _GotoState26Action,
	{_State163, DotToken}:                         _GotoState76Action,
	{_State163, QuestionToken}:                    _GotoState83Action,
	{_State163, ExclaimToken}:                     _GotoState78Action,
	{_State163, DotDotDotToken}:                   _GotoState259Action,
	{_State163, TildeTildeToken}:                  _GotoState84Action,
	{_State163, BitNegToken}:                      _GotoState75Action,
	{_State163, BitAndToken}:                      _GotoState74Action,
	{_State163, ParseErrorToken}:                  _GotoState82Action,
	{_State163, InitializableTypeType}:            _GotoState91Action,
	{_State163, AtomTypeType}:                     _GotoState86Action,
	{_State163, ReturnableTypeType}:               _GotoState92Action,
	{_State163, ValueTypeType}:                    _GotoState264Action,
	{_State163, ImplicitStructDefType}:            _GotoState90Action,
	{_State163, ExplicitStructDefType}:            _GotoState45Action,
	{_State163, ImplicitEnumDefType}:              _GotoState89Action,
	{_State163, ExplicitEnumDefType}:              _GotoState87Action,
	{_State163, TraitDefType}:                     _GotoState93Action,
	{_State163, ParameterDeclType}:                _GotoState262Action,
	{_State163, ParameterDeclsType}:               _GotoState263Action,
	{_State163, OptionalParameterDeclsType}:       _GotoState261Action,
	{_State163, FuncTypeType}:                     _GotoState88Action,
	{_State164, IdentifierToken}:                  _GotoState265Action,
	{_State166, IdentifierToken}:                  _GotoState80Action,
	{_State166, StructToken}:                      _GotoState34Action,
	{_State166, EnumToken}:                        _GotoState77Action,
	{_State166, TraitToken}:                       _GotoState85Action,
	{_State166, FuncToken}:                        _GotoState79Action,
	{_State166, LparenToken}:                      _GotoState81Action,
	{_State166, LbracketToken}:                    _GotoState26Action,
	{_State166, DotToken}:                         _GotoState266Action,
	{_State166, QuestionToken}:                    _GotoState83Action,
	{_State166, ExclaimToken}:                     _GotoState78Action,
	{_State166, DollarLbracketToken}:              _GotoState104Action,
	{_State166, TildeTildeToken}:                  _GotoState84Action,
	{_State166, BitNegToken}:                      _GotoState75Action,
	{_State166, BitAndToken}:                      _GotoState74Action,
	{_State166, ParseErrorToken}:                  _GotoState82Action,
	{_State166, OptionalGenericBindingType}:       _GotoState165Action,
	{_State166, InitializableTypeType}:            _GotoState91Action,
	{_State166, AtomTypeType}:                     _GotoState86Action,
	{_State166, ReturnableTypeType}:               _GotoState92Action,
	{_State166, ValueTypeType}:                    _GotoState267Action,
	{_State166, ImplicitStructDefType}:            _GotoState90Action,
	{_State166, ExplicitStructDefType}:            _GotoState45Action,
	{_State166, ImplicitEnumDefType}:              _GotoState89Action,
	{_State166, ExplicitEnumDefType}:              _GotoState87Action,
	{_State166, TraitDefType}:                     _GotoState93Action,
	{_State166, FuncTypeType}:                     _GotoState88Action,
	{_State167, LessToken}:                        _GotoState268Action,
	{_State168, OrToken}:                          _GotoState269Action,
	{_State169, AssignToken}:                      _GotoState270Action,
	{_State170, OrToken}:                          _GotoState271Action,
	{_State170, RparenToken}:                      _GotoState272Action,
	{_State171, CommaToken}:                       _GotoState273Action,
	{_State172, RparenToken}:                      _GotoState274Action,
	{_State174, AddToken}:                         _GotoState178Action,
	{_State174, SubToken}:                         _GotoState183Action,
	{_State174, MulToken}:                         _GotoState181Action,
	{_State177, IdentifierToken}:                  _GotoState166Action,
	{_State177, UnsafeToken}:                      _GotoState167Action,
	{_State177, StructToken}:                      _GotoState34Action,
	{_State177, EnumToken}:                        _GotoState77Action,
	{_State177, TraitToken}:                       _GotoState85Action,
	{_State177, FuncToken}:                        _GotoState275Action,
	{_State177, LparenToken}:                      _GotoState81Action,
	{_State177, LbracketToken}:                    _GotoState26Action,
	{_State177, DotToken}:                         _GotoState76Action,
	{_State177, QuestionToken}:                    _GotoState83Action,
	{_State177, ExclaimToken}:                     _GotoState78Action,
	{_State177, TildeTildeToken}:                  _GotoState84Action,
	{_State177, BitNegToken}:                      _GotoState75Action,
	{_State177, BitAndToken}:                      _GotoState74Action,
	{_State177, ParseErrorToken}:                  _GotoState82Action,
	{_State177, UnsafeStatementType}:              _GotoState173Action,
	{_State177, InitializableTypeType}:            _GotoState91Action,
	{_State177, AtomTypeType}:                     _GotoState86Action,
	{_State177, ReturnableTypeType}:               _GotoState92Action,
	{_State177, ValueTypeType}:                    _GotoState174Action,
	{_State177, FieldDefType}:                     _GotoState276Action,
	{_State177, ImplicitStructDefType}:            _GotoState90Action,
	{_State177, ExplicitStructDefType}:            _GotoState45Action,
	{_State177, ImplicitEnumDefType}:              _GotoState89Action,
	{_State177, ExplicitEnumDefType}:              _GotoState87Action,
	{_State177, TraitPropertyType}:                _GotoState280Action,
	{_State177, TraitPropertiesType}:              _GotoState279Action,
	{_State177, OptionalTraitPropertiesType}:      _GotoState278Action,
	{_State177, TraitDefType}:                     _GotoState93Action,
	{_State177, FuncTypeType}:                     _GotoState88Action,
	{_State177, MethodSignatureType}:              _GotoState277Action,
	{_State178, IdentifierToken}:                  _GotoState80Action,
	{_State178, StructToken}:                      _GotoState34Action,
	{_State178, EnumToken}:                        _GotoState77Action,
	{_State178, TraitToken}:                       _GotoState85Action,
	{_State178, FuncToken}:                        _GotoState79Action,
	{_State178, LparenToken}:                      _GotoState81Action,
	{_State178, LbracketToken}:                    _GotoState26Action,
	{_State178, DotToken}:                         _GotoState76Action,
	{_State178, QuestionToken}:                    _GotoState83Action,
	{_State178, ExclaimToken}:                     _GotoState78Action,
	{_State178, TildeTildeToken}:                  _GotoState84Action,
	{_State178, BitNegToken}:                      _GotoState75Action,
	{_State178, BitAndToken}:                      _GotoState74Action,
	{_State178, ParseErrorToken}:                  _GotoState82Action,
	{_State178, InitializableTypeType}:            _GotoState91Action,
	{_State178, AtomTypeType}:                     _GotoState86Action,
	{_State178, ReturnableTypeType}:               _GotoState281Action,
	{_State178, ImplicitStructDefType}:            _GotoState90Action,
	{_State178, ExplicitStructDefType}:            _GotoState45Action,
	{_State178, ImplicitEnumDefType}:              _GotoState89Action,
	{_State178, ExplicitEnumDefType}:              _GotoState87Action,
	{_State178, TraitDefType}:                     _GotoState93Action,
	{_State178, FuncTypeType}:                     _GotoState88Action,
	{_State179, IdentifierToken}:                  _GotoState80Action,
	{_State179, StructToken}:                      _GotoState34Action,
	{_State179, EnumToken}:                        _GotoState77Action,
	{_State179, TraitToken}:                       _GotoState85Action,
	{_State179, FuncToken}:                        _GotoState79Action,
	{_State179, LparenToken}:                      _GotoState81Action,
	{_State179, LbracketToken}:                    _GotoState26Action,
	{_State179, DotToken}:                         _GotoState76Action,
	{_State179, QuestionToken}:                    _GotoState83Action,
	{_State179, ExclaimToken}:                     _GotoState78Action,
	{_State179, TildeTildeToken}:                  _GotoState84Action,
	{_State179, BitNegToken}:                      _GotoState75Action,
	{_State179, BitAndToken}:                      _GotoState74Action,
	{_State179, ParseErrorToken}:                  _GotoState82Action,
	{_State179, InitializableTypeType}:            _GotoState91Action,
	{_State179, AtomTypeType}:                     _GotoState86Action,
	{_State179, ReturnableTypeType}:               _GotoState92Action,
	{_State179, ValueTypeType}:                    _GotoState282Action,
	{_State179, ImplicitStructDefType}:            _GotoState90Action,
	{_State179, ExplicitStructDefType}:            _GotoState45Action,
	{_State179, ImplicitEnumDefType}:              _GotoState89Action,
	{_State179, ExplicitEnumDefType}:              _GotoState87Action,
	{_State179, TraitDefType}:                     _GotoState93Action,
	{_State179, FuncTypeType}:                     _GotoState88Action,
	{_State180, IntegerLiteralToken}:              _GotoState283Action,
	{_State181, IdentifierToken}:                  _GotoState80Action,
	{_State181, StructToken}:                      _GotoState34Action,
	{_State181, EnumToken}:                        _GotoState77Action,
	{_State181, TraitToken}:                       _GotoState85Action,
	{_State181, FuncToken}:                        _GotoState79Action,
	{_State181, LparenToken}:                      _GotoState81Action,
	{_State181, LbracketToken}:                    _GotoState26Action,
	{_State181, DotToken}:                         _GotoState76Action,
	{_State181, QuestionToken}:                    _GotoState83Action,
	{_State181, ExclaimToken}:                     _GotoState78Action,
	{_State181, TildeTildeToken}:                  _GotoState84Action,
	{_State181, BitNegToken}:                      _GotoState75Action,
	{_State181, BitAndToken}:                      _GotoState74Action,
	{_State181, ParseErrorToken}:                  _GotoState82Action,
	{_State181, InitializableTypeType}:            _GotoState91Action,
	{_State181, AtomTypeType}:                     _GotoState86Action,
	{_State181, ReturnableTypeType}:               _GotoState284Action,
	{_State181, ImplicitStructDefType}:            _GotoState90Action,
	{_State181, ExplicitStructDefType}:            _GotoState45Action,
	{_State181, ImplicitEnumDefType}:              _GotoState89Action,
	{_State181, ExplicitEnumDefType}:              _GotoState87Action,
	{_State181, TraitDefType}:                     _GotoState93Action,
	{_State181, FuncTypeType}:                     _GotoState88Action,
	{_State183, IdentifierToken}:                  _GotoState80Action,
	{_State183, StructToken}:                      _GotoState34Action,
	{_State183, EnumToken}:                        _GotoState77Action,
	{_State183, TraitToken}:                       _GotoState85Action,
	{_State183, FuncToken}:                        _GotoState79Action,
	{_State183, LparenToken}:                      _GotoState81Action,
	{_State183, LbracketToken}:                    _GotoState26Action,
	{_State183, DotToken}:                         _GotoState76Action,
	{_State183, QuestionToken}:                    _GotoState83Action,
	{_State183, ExclaimToken}:                     _GotoState78Action,
	{_State183, TildeTildeToken}:                  _GotoState84Action,
	{_State183, BitNegToken}:                      _GotoState75Action,
	{_State183, BitAndToken}:                      _GotoState74Action,
	{_State183, ParseErrorToken}:                  _GotoState82Action,
	{_State183, InitializableTypeType}:            _GotoState91Action,
	{_State183, AtomTypeType}:                     _GotoState86Action,
	{_State183, ReturnableTypeType}:               _GotoState285Action,
	{_State183, ImplicitStructDefType}:            _GotoState90Action,
	{_State183, ExplicitStructDefType}:            _GotoState45Action,
	{_State183, ImplicitEnumDefType}:              _GotoState89Action,
	{_State183, ExplicitEnumDefType}:              _GotoState87Action,
	{_State183, TraitDefType}:                     _GotoState93Action,
	{_State183, FuncTypeType}:                     _GotoState88Action,
	{_State184, IntegerLiteralToken}:              _GotoState24Action,
	{_State184, FloatLiteralToken}:                _GotoState20Action,
	{_State184, RuneLiteralToken}:                 _GotoState32Action,
	{_State184, StringLiteralToken}:               _GotoState33Action,
	{_State184, IdentifierToken}:                  _GotoState23Action,
	{_State184, TrueToken}:                        _GotoState36Action,
	{_State184, FalseToken}:                       _GotoState19Action,
	{_State184, StructToken}:                      _GotoState34Action,
	{_State184, FuncToken}:                        _GotoState21Action,
	{_State184, VarToken}:                         _GotoState37Action,
	{_State184, LetToken}:                         _GotoState27Action,
	{_State184, NotToken}:                         _GotoState30Action,
	{_State184, LabelDeclToken}:                   _GotoState25Action,
	{_State184, LparenToken}:                      _GotoState28Action,
	{_State184, LbracketToken}:                    _GotoState26Action,
	{_State184, SubToken}:                         _GotoState35Action,
	{_State184, MulToken}:                         _GotoState29Action,
	{_State184, BitNegToken}:                      _GotoState18Action,
	{_State184, BitAndToken}:                      _GotoState17Action,
	{_State184, GreaterToken}:                     _GotoState22Action,
	{_State184, ParseErrorToken}:                  _GotoState31Action,
	{_State184, VarDeclPatternType}:               _GotoState56Action,
	{_State184, VarOrLetType}:                     _GotoState57Action,
	{_State184, ExpressionType}:                   _GotoState286Action,
	{_State184, OptionalLabelDeclType}:            _GotoState50Action,
	{_State184, SequenceExprType}:                 _GotoState55Action,
	{_State184, CallExprType}:                     _GotoState43Action,
	{_State184, AtomExprType}:                     _GotoState42Action,
	{_State184, LiteralExprType}:                  _GotoState48Action,
	{_State184, ImplicitStructExprType}:           _GotoState46Action,
	{_State184, AccessExprType}:                   _GotoState38Action,
	{_State184, PostfixUnaryExprType}:             _GotoState52Action,
	{_State184, PrefixUnaryOpType}:                _GotoState54Action,
	{_State184, PrefixUnaryExprType}:              _GotoState53Action,
	{_State184, MulExprType}:                      _GotoState49Action,
	{_State184, AddExprType}:                      _GotoState39Action,
	{_State184, CmpExprType}:                      _GotoState44Action,
	{_State184, AndExprType}:                      _GotoState40Action,
	{_State184, OrExprType}:                       _GotoState51Action,
	{_State184, InitializableTypeType}:            _GotoState47Action,
	{_State184, ExplicitStructDefType}:            _GotoState45Action,
	{_State184, AnonymousFuncExprType}:            _GotoState41Action,
	{_State185, IntegerLiteralToken}:              _GotoState24Action,
	{_State185, FloatLiteralToken}:                _GotoState20Action,
	{_State185, RuneLiteralToken}:                 _GotoState32Action,
	{_State185, StringLiteralToken}:               _GotoState33Action,
	{_State185, IdentifierToken}:                  _GotoState96Action,
	{_State185, TrueToken}:                        _GotoState36Action,
	{_State185, FalseToken}:                       _GotoState19Action,
	{_State185, StructToken}:                      _GotoState34Action,
	{_State185, FuncToken}:                        _GotoState21Action,
	{_State185, VarToken}:                         _GotoState37Action,
	{_State185, LetToken}:                         _GotoState27Action,
	{_State185, NotToken}:                         _GotoState30Action,
	{_State185, LabelDeclToken}:                   _GotoState25Action,
	{_State185, LparenToken}:                      _GotoState28Action,
	{_State185, LbracketToken}:                    _GotoState26Action,
	{_State185, DotDotDotToken}:                   _GotoState95Action,
	{_State185, SubToken}:                         _GotoState35Action,
	{_State185, MulToken}:                         _GotoState29Action,
	{_State185, BitNegToken}:                      _GotoState18Action,
	{_State185, BitAndToken}:                      _GotoState17Action,
	{_State185, GreaterToken}:                     _GotoState22Action,
	{_State185, ParseErrorToken}:                  _GotoState31Action,
	{_State185, VarDeclPatternType}:               _GotoState56Action,
	{_State185, VarOrLetType}:                     _GotoState57Action,
	{_State185, ExpressionType}:                   _GotoState100Action,
	{_State185, OptionalLabelDeclType}:            _GotoState50Action,
	{_State185, SequenceExprType}:                 _GotoState55Action,
	{_State185, CallExprType}:                     _GotoState43Action,
	{_State185, ArgumentType}:                     _GotoState287Action,
	{_State185, ColonExpressionsType}:             _GotoState99Action,
	{_State185, OptionalExpressionType}:           _GotoState102Action,
	{_State185, AtomExprType}:                     _GotoState42Action,
	{_State185, LiteralExprType}:                  _GotoState48Action,
	{_State185, ImplicitStructExprType}:           _GotoState46Action,
	{_State185, AccessExprType}:                   _GotoState38Action,
	{_State185, PostfixUnaryExprType}:             _GotoState52Action,
	{_State185, PrefixUnaryOpType}:                _GotoState54Action,
	{_State185, PrefixUnaryExprType}:              _GotoState53Action,
	{_State185, MulExprType}:                      _GotoState49Action,
	{_State185, AddExprType}:                      _GotoState39Action,
	{_State185, CmpExprType}:                      _GotoState44Action,
	{_State185, AndExprType}:                      _GotoState40Action,
	{_State185, OrExprType}:                       _GotoState51Action,
	{_State185, InitializableTypeType}:            _GotoState47Action,
	{_State185, ExplicitStructDefType}:            _GotoState45Action,
	{_State185, AnonymousFuncExprType}:            _GotoState41Action,
	{_State186, IntegerLiteralToken}:              _GotoState24Action,
	{_State186, FloatLiteralToken}:                _GotoState20Action,
	{_State186, RuneLiteralToken}:                 _GotoState32Action,
	{_State186, StringLiteralToken}:               _GotoState33Action,
	{_State186, IdentifierToken}:                  _GotoState23Action,
	{_State186, TrueToken}:                        _GotoState36Action,
	{_State186, FalseToken}:                       _GotoState19Action,
	{_State186, StructToken}:                      _GotoState34Action,
	{_State186, FuncToken}:                        _GotoState21Action,
	{_State186, VarToken}:                         _GotoState37Action,
	{_State186, LetToken}:                         _GotoState27Action,
	{_State186, NotToken}:                         _GotoState30Action,
	{_State186, LabelDeclToken}:                   _GotoState25Action,
	{_State186, LparenToken}:                      _GotoState28Action,
	{_State186, LbracketToken}:                    _GotoState26Action,
	{_State186, SubToken}:                         _GotoState35Action,
	{_State186, MulToken}:                         _GotoState29Action,
	{_State186, BitNegToken}:                      _GotoState18Action,
	{_State186, BitAndToken}:                      _GotoState17Action,
	{_State186, GreaterToken}:                     _GotoState22Action,
	{_State186, ParseErrorToken}:                  _GotoState31Action,
	{_State186, VarDeclPatternType}:               _GotoState56Action,
	{_State186, VarOrLetType}:                     _GotoState57Action,
	{_State186, ExpressionType}:                   _GotoState288Action,
	{_State186, OptionalLabelDeclType}:            _GotoState50Action,
	{_State186, SequenceExprType}:                 _GotoState55Action,
	{_State186, CallExprType}:                     _GotoState43Action,
	{_State186, OptionalExpressionType}:           _GotoState289Action,
	{_State186, AtomExprType}:                     _GotoState42Action,
	{_State186, LiteralExprType}:                  _GotoState48Action,
	{_State186, ImplicitStructExprType}:           _GotoState46Action,
	{_State186, AccessExprType}:                   _GotoState38Action,
	{_State186, PostfixUnaryExprType}:             _GotoState52Action,
	{_State186, PrefixUnaryOpType}:                _GotoState54Action,
	{_State186, PrefixUnaryExprType}:              _GotoState53Action,
	{_State186, MulExprType}:                      _GotoState49Action,
	{_State186, AddExprType}:                      _GotoState39Action,
	{_State186, CmpExprType}:                      _GotoState44Action,
	{_State186, AndExprType}:                      _GotoState40Action,
	{_State186, OrExprType}:                       _GotoState51Action,
	{_State186, InitializableTypeType}:            _GotoState47Action,
	{_State186, ExplicitStructDefType}:            _GotoState45Action,
	{_State186, AnonymousFuncExprType}:            _GotoState41Action,
	{_State188, IntegerLiteralToken}:              _GotoState24Action,
	{_State188, FloatLiteralToken}:                _GotoState20Action,
	{_State188, RuneLiteralToken}:                 _GotoState32Action,
	{_State188, StringLiteralToken}:               _GotoState33Action,
	{_State188, IdentifierToken}:                  _GotoState23Action,
	{_State188, TrueToken}:                        _GotoState36Action,
	{_State188, FalseToken}:                       _GotoState19Action,
	{_State188, StructToken}:                      _GotoState34Action,
	{_State188, FuncToken}:                        _GotoState21Action,
	{_State188, VarToken}:                         _GotoState37Action,
	{_State188, LetToken}:                         _GotoState27Action,
	{_State188, NotToken}:                         _GotoState30Action,
	{_State188, LabelDeclToken}:                   _GotoState25Action,
	{_State188, LparenToken}:                      _GotoState28Action,
	{_State188, LbracketToken}:                    _GotoState26Action,
	{_State188, SubToken}:                         _GotoState35Action,
	{_State188, MulToken}:                         _GotoState29Action,
	{_State188, BitNegToken}:                      _GotoState18Action,
	{_State188, BitAndToken}:                      _GotoState17Action,
	{_State188, GreaterToken}:                     _GotoState22Action,
	{_State188, ParseErrorToken}:                  _GotoState31Action,
	{_State188, VarDeclPatternType}:               _GotoState56Action,
	{_State188, VarOrLetType}:                     _GotoState57Action,
	{_State188, ExpressionType}:                   _GotoState288Action,
	{_State188, OptionalLabelDeclType}:            _GotoState50Action,
	{_State188, SequenceExprType}:                 _GotoState55Action,
	{_State188, CallExprType}:                     _GotoState43Action,
	{_State188, OptionalExpressionType}:           _GotoState290Action,
	{_State188, AtomExprType}:                     _GotoState42Action,
	{_State188, LiteralExprType}:                  _GotoState48Action,
	{_State188, ImplicitStructExprType}:           _GotoState46Action,
	{_State188, AccessExprType}:                   _GotoState38Action,
	{_State188, PostfixUnaryExprType}:             _GotoState52Action,
	{_State188, PrefixUnaryOpType}:                _GotoState54Action,
	{_State188, PrefixUnaryExprType}:              _GotoState53Action,
	{_State188, MulExprType}:                      _GotoState49Action,
	{_State188, AddExprType}:                      _GotoState39Action,
	{_State188, CmpExprType}:                      _GotoState44Action,
	{_State188, AndExprType}:                      _GotoState40Action,
	{_State188, OrExprType}:                       _GotoState51Action,
	{_State188, InitializableTypeType}:            _GotoState47Action,
	{_State188, ExplicitStructDefType}:            _GotoState45Action,
	{_State188, AnonymousFuncExprType}:            _GotoState41Action,
	{_State189, NewlinesToken}:                    _GotoState292Action,
	{_State189, CommaToken}:                       _GotoState291Action,
	{_State191, RparenToken}:                      _GotoState293Action,
	{_State192, CommaToken}:                       _GotoState294Action,
	{_State193, RbracketToken}:                    _GotoState295Action,
	{_State194, AddToken}:                         _GotoState178Action,
	{_State194, SubToken}:                         _GotoState183Action,
	{_State194, MulToken}:                         _GotoState181Action,
	{_State196, RbracketToken}:                    _GotoState296Action,
	{_State197, IntegerLiteralToken}:              _GotoState24Action,
	{_State197, FloatLiteralToken}:                _GotoState20Action,
	{_State197, RuneLiteralToken}:                 _GotoState32Action,
	{_State197, StringLiteralToken}:               _GotoState33Action,
	{_State197, IdentifierToken}:                  _GotoState96Action,
	{_State197, TrueToken}:                        _GotoState36Action,
	{_State197, FalseToken}:                       _GotoState19Action,
	{_State197, StructToken}:                      _GotoState34Action,
	{_State197, FuncToken}:                        _GotoState21Action,
	{_State197, VarToken}:                         _GotoState37Action,
	{_State197, LetToken}:                         _GotoState27Action,
	{_State197, NotToken}:                         _GotoState30Action,
	{_State197, LabelDeclToken}:                   _GotoState25Action,
	{_State197, LparenToken}:                      _GotoState28Action,
	{_State197, LbracketToken}:                    _GotoState26Action,
	{_State197, DotDotDotToken}:                   _GotoState95Action,
	{_State197, SubToken}:                         _GotoState35Action,
	{_State197, MulToken}:                         _GotoState29Action,
	{_State197, BitNegToken}:                      _GotoState18Action,
	{_State197, BitAndToken}:                      _GotoState17Action,
	{_State197, GreaterToken}:                     _GotoState22Action,
	{_State197, ParseErrorToken}:                  _GotoState31Action,
	{_State197, VarDeclPatternType}:               _GotoState56Action,
	{_State197, VarOrLetType}:                     _GotoState57Action,
	{_State197, ExpressionType}:                   _GotoState100Action,
	{_State197, OptionalLabelDeclType}:            _GotoState50Action,
	{_State197, SequenceExprType}:                 _GotoState55Action,
	{_State197, CallExprType}:                     _GotoState43Action,
	{_State197, OptionalArgumentsType}:            _GotoState297Action,
	{_State197, ArgumentsType}:                    _GotoState98Action,
	{_State197, ArgumentType}:                     _GotoState97Action,
	{_State197, ColonExpressionsType}:             _GotoState99Action,
	{_State197, OptionalExpressionType}:           _GotoState102Action,
	{_State197, AtomExprType}:                     _GotoState42Action,
	{_State197, LiteralExprType}:                  _GotoState48Action,
	{_State197, ImplicitStructExprType}:           _GotoState46Action,
	{_State197, AccessExprType}:                   _GotoState38Action,
	{_State197, PostfixUnaryExprType}:             _GotoState52Action,
	{_State197, PrefixUnaryOpType}:                _GotoState54Action,
	{_State197, PrefixUnaryExprType}:              _GotoState53Action,
	{_State197, MulExprType}:                      _GotoState49Action,
	{_State197, AddExprType}:                      _GotoState39Action,
	{_State197, CmpExprType}:                      _GotoState44Action,
	{_State197, AndExprType}:                      _GotoState40Action,
	{_State197, OrExprType}:                       _GotoState51Action,
	{_State197, InitializableTypeType}:            _GotoState47Action,
	{_State197, ExplicitStructDefType}:            _GotoState45Action,
	{_State197, AnonymousFuncExprType}:            _GotoState41Action,
	{_State198, MulToken}:                         _GotoState128Action,
	{_State198, DivToken}:                         _GotoState126Action,
	{_State198, ModToken}:                         _GotoState127Action,
	{_State198, BitAndToken}:                      _GotoState123Action,
	{_State198, BitLshiftToken}:                   _GotoState124Action,
	{_State198, BitRshiftToken}:                   _GotoState125Action,
	{_State198, MulOpType}:                        _GotoState129Action,
	{_State199, EqualToken}:                       _GotoState115Action,
	{_State199, NotEqualToken}:                    _GotoState120Action,
	{_State199, LessToken}:                        _GotoState118Action,
	{_State199, LessOrEqualToken}:                 _GotoState119Action,
	{_State199, GreaterToken}:                     _GotoState116Action,
	{_State199, GreaterOrEqualToken}:              _GotoState117Action,
	{_State199, CmpOpType}:                        _GotoState121Action,
	{_State200, AddToken}:                         _GotoState109Action,
	{_State200, SubToken}:                         _GotoState112Action,
	{_State200, BitXorToken}:                      _GotoState111Action,
	{_State200, BitOrToken}:                       _GotoState110Action,
	{_State200, AddOpType}:                        _GotoState113Action,
	{_State201, RparenToken}:                      _GotoState298Action,
	{_State203, ForToken}:                         _GotoState299Action,
	{_State204, InToken}:                          _GotoState301Action,
	{_State204, AssignToken}:                      _GotoState300Action,
	{_State205, SemicolonToken}:                   _GotoState302Action,
	{_State206, DoToken}:                          _GotoState303Action,
	{_State207, IntegerLiteralToken}:              _GotoState24Action,
	{_State207, FloatLiteralToken}:                _GotoState20Action,
	{_State207, RuneLiteralToken}:                 _GotoState32Action,
	{_State207, StringLiteralToken}:               _GotoState33Action,
	{_State207, IdentifierToken}:                  _GotoState23Action,
	{_State207, TrueToken}:                        _GotoState36Action,
	{_State207, FalseToken}:                       _GotoState19Action,
	{_State207, StructToken}:                      _GotoState34Action,
	{_State207, FuncToken}:                        _GotoState21Action,
	{_State207, VarToken}:                         _GotoState305Action,
	{_State207, LetToken}:                         _GotoState27Action,
	{_State207, NotToken}:                         _GotoState30Action,
	{_State207, LabelDeclToken}:                   _GotoState25Action,
	{_State207, LparenToken}:                      _GotoState28Action,
	{_State207, LbracketToken}:                    _GotoState26Action,
	{_State207, DotToken}:                         _GotoState304Action,
	{_State207, SubToken}:                         _GotoState35Action,
	{_State207, MulToken}:                         _GotoState29Action,
	{_State207, BitNegToken}:                      _GotoState18Action,
	{_State207, BitAndToken}:                      _GotoState17Action,
	{_State207, GreaterToken}:                     _GotoState22Action,
	{_State207, ParseErrorToken}:                  _GotoState31Action,
	{_State207, CasePatternsType}:                 _GotoState308Action,
	{_State207, VarDeclPatternType}:               _GotoState56Action,
	{_State207, VarOrLetType}:                     _GotoState57Action,
	{_State207, AssignPatternType}:                _GotoState306Action,
	{_State207, CasePatternType}:                  _GotoState307Action,
	{_State207, OptionalLabelDeclType}:            _GotoState72Action,
	{_State207, SequenceExprType}:                 _GotoState309Action,
	{_State207, CallExprType}:                     _GotoState43Action,
	{_State207, AtomExprType}:                     _GotoState42Action,
	{_State207, LiteralExprType}:                  _GotoState48Action,
	{_State207, ImplicitStructExprType}:           _GotoState46Action,
	{_State207, AccessExprType}:                   _GotoState38Action,
	{_State207, PostfixUnaryExprType}:             _GotoState52Action,
	{_State207, PrefixUnaryOpType}:                _GotoState54Action,
	{_State207, PrefixUnaryExprType}:              _GotoState53Action,
	{_State207, MulExprType}:                      _GotoState49Action,
	{_State207, AddExprType}:                      _GotoState39Action,
	{_State207, CmpExprType}:                      _GotoState44Action,
	{_State207, AndExprType}:                      _GotoState40Action,
	{_State207, OrExprType}:                       _GotoState51Action,
	{_State207, InitializableTypeType}:            _GotoState47Action,
	{_State207, ExplicitStructDefType}:            _GotoState45Action,
	{_State207, AnonymousFuncExprType}:            _GotoState41Action,
	{_State208, LbraceToken}:                      _GotoState59Action,
	{_State208, StatementBlockType}:               _GotoState310Action,
	{_State210, LbraceToken}:                      _GotoState59Action,
	{_State210, StatementBlockType}:               _GotoState311Action,
	{_State211, AndToken}:                         _GotoState114Action,
	{_State213, AssignToken}:                      _GotoState312Action,
	{_State215, RparenToken}:                      _GotoState314Action,
	{_State215, CommaToken}:                       _GotoState313Action,
	{_State218, AddToken}:                         _GotoState178Action,
	{_State218, SubToken}:                         _GotoState183Action,
	{_State218, MulToken}:                         _GotoState181Action,
	{_State219, IntegerLiteralToken}:              _GotoState24Action,
	{_State219, FloatLiteralToken}:                _GotoState20Action,
	{_State219, RuneLiteralToken}:                 _GotoState32Action,
	{_State219, StringLiteralToken}:               _GotoState33Action,
	{_State219, IdentifierToken}:                  _GotoState23Action,
	{_State219, TrueToken}:                        _GotoState36Action,
	{_State219, FalseToken}:                       _GotoState19Action,
	{_State219, StructToken}:                      _GotoState34Action,
	{_State219, FuncToken}:                        _GotoState21Action,
	{_State219, LabelDeclToken}:                   _GotoState25Action,
	{_State219, LparenToken}:                      _GotoState28Action,
	{_State219, LbracketToken}:                    _GotoState26Action,
	{_State219, ParseErrorToken}:                  _GotoState31Action,
	{_State219, OptionalLabelDeclType}:            _GotoState72Action,
	{_State219, CallExprType}:                     _GotoState316Action,
	{_State219, AtomExprType}:                     _GotoState42Action,
	{_State219, LiteralExprType}:                  _GotoState48Action,
	{_State219, ImplicitStructExprType}:           _GotoState46Action,
	{_State219, AccessExprType}:                   _GotoState315Action,
	{_State219, InitializableTypeType}:            _GotoState47Action,
	{_State219, ExplicitStructDefType}:            _GotoState45Action,
	{_State219, AnonymousFuncExprType}:            _GotoState41Action,
	{_State221, IntegerLiteralToken}:              _GotoState24Action,
	{_State221, FloatLiteralToken}:                _GotoState20Action,
	{_State221, RuneLiteralToken}:                 _GotoState32Action,
	{_State221, StringLiteralToken}:               _GotoState33Action,
	{_State221, IdentifierToken}:                  _GotoState23Action,
	{_State221, TrueToken}:                        _GotoState36Action,
	{_State221, FalseToken}:                       _GotoState19Action,
	{_State221, StructToken}:                      _GotoState34Action,
	{_State221, FuncToken}:                        _GotoState21Action,
	{_State221, VarToken}:                         _GotoState305Action,
	{_State221, LetToken}:                         _GotoState27Action,
	{_State221, NotToken}:                         _GotoState30Action,
	{_State221, LabelDeclToken}:                   _GotoState25Action,
	{_State221, LparenToken}:                      _GotoState28Action,
	{_State221, LbracketToken}:                    _GotoState26Action,
	{_State221, DotToken}:                         _GotoState304Action,
	{_State221, SubToken}:                         _GotoState35Action,
	{_State221, MulToken}:                         _GotoState29Action,
	{_State221, BitNegToken}:                      _GotoState18Action,
	{_State221, BitAndToken}:                      _GotoState17Action,
	{_State221, GreaterToken}:                     _GotoState22Action,
	{_State221, ParseErrorToken}:                  _GotoState31Action,
	{_State221, CasePatternsType}:                 _GotoState317Action,
	{_State221, VarDeclPatternType}:               _GotoState56Action,
	{_State221, VarOrLetType}:                     _GotoState57Action,
	{_State221, AssignPatternType}:                _GotoState306Action,
	{_State221, CasePatternType}:                  _GotoState307Action,
	{_State221, OptionalLabelDeclType}:            _GotoState72Action,
	{_State221, SequenceExprType}:                 _GotoState309Action,
	{_State221, CallExprType}:                     _GotoState43Action,
	{_State221, AtomExprType}:                     _GotoState42Action,
	{_State221, LiteralExprType}:                  _GotoState48Action,
	{_State221, ImplicitStructExprType}:           _GotoState46Action,
	{_State221, AccessExprType}:                   _GotoState38Action,
	{_State221, PostfixUnaryExprType}:             _GotoState52Action,
	{_State221, PrefixUnaryOpType}:                _GotoState54Action,
	{_State221, PrefixUnaryExprType}:              _GotoState53Action,
	{_State221, MulExprType}:                      _GotoState49Action,
	{_State221, AddExprType}:                      _GotoState39Action,
	{_State221, CmpExprType}:                      _GotoState44Action,
	{_State221, AndExprType}:                      _GotoState40Action,
	{_State221, OrExprType}:                       _GotoState51Action,
	{_State221, InitializableTypeType}:            _GotoState47Action,
	{_State221, ExplicitStructDefType}:            _GotoState45Action,
	{_State221, AnonymousFuncExprType}:            _GotoState41Action,
	{_State223, ColonToken}:                       _GotoState318Action,
	{_State224, IntegerLiteralToken}:              _GotoState24Action,
	{_State224, FloatLiteralToken}:                _GotoState20Action,
	{_State224, RuneLiteralToken}:                 _GotoState32Action,
	{_State224, StringLiteralToken}:               _GotoState33Action,
	{_State224, IdentifierToken}:                  _GotoState23Action,
	{_State224, TrueToken}:                        _GotoState36Action,
	{_State224, FalseToken}:                       _GotoState19Action,
	{_State224, StructToken}:                      _GotoState34Action,
	{_State224, FuncToken}:                        _GotoState21Action,
	{_State224, LabelDeclToken}:                   _GotoState25Action,
	{_State224, LparenToken}:                      _GotoState28Action,
	{_State224, LbracketToken}:                    _GotoState26Action,
	{_State224, ParseErrorToken}:                  _GotoState31Action,
	{_State224, OptionalLabelDeclType}:            _GotoState72Action,
	{_State224, CallExprType}:                     _GotoState319Action,
	{_State224, AtomExprType}:                     _GotoState42Action,
	{_State224, LiteralExprType}:                  _GotoState48Action,
	{_State224, ImplicitStructExprType}:           _GotoState46Action,
	{_State224, AccessExprType}:                   _GotoState315Action,
	{_State224, InitializableTypeType}:            _GotoState47Action,
	{_State224, ExplicitStructDefType}:            _GotoState45Action,
	{_State224, AnonymousFuncExprType}:            _GotoState41Action,
	{_State226, StringLiteralToken}:               _GotoState321Action,
	{_State226, LparenToken}:                      _GotoState320Action,
	{_State226, ImportClauseType}:                 _GotoState322Action,
	{_State229, LbracketToken}:                    _GotoState106Action,
	{_State229, DotToken}:                         _GotoState105Action,
	{_State229, QuestionToken}:                    _GotoState107Action,
	{_State229, DollarLbracketToken}:              _GotoState104Action,
	{_State229, AddAssignToken}:                   _GotoState323Action,
	{_State229, SubAssignToken}:                   _GotoState334Action,
	{_State229, MulAssignToken}:                   _GotoState333Action,
	{_State229, DivAssignToken}:                   _GotoState331Action,
	{_State229, ModAssignToken}:                   _GotoState332Action,
	{_State229, AddOneAssignToken}:                _GotoState324Action,
	{_State229, SubOneAssignToken}:                _GotoState335Action,
	{_State229, BitNegAssignToken}:                _GotoState327Action,
	{_State229, BitAndAssignToken}:                _GotoState325Action,
	{_State229, BitOrAssignToken}:                 _GotoState328Action,
	{_State229, BitXorAssignToken}:                _GotoState330Action,
	{_State229, BitLshiftAssignToken}:             _GotoState326Action,
	{_State229, BitRshiftAssignToken}:             _GotoState329Action,
	{_State229, UnaryOpAssignType}:                _GotoState337Action,
	{_State229, BinaryOpAssignType}:               _GotoState336Action,
	{_State229, OptionalGenericBindingType}:       _GotoState108Action,
	{_State230, AssignToken}:                      _GotoState338Action,
	{_State232, CommaToken}:                       _GotoState339Action,
	{_State234, JumpLabelToken}:                   _GotoState340Action,
	{_State234, OptionalJumpLabelType}:            _GotoState341Action,
	{_State238, NewlinesToken}:                    _GotoState342Action,
	{_State238, SemicolonToken}:                   _GotoState343Action,
	{_State242, AddToken}:                         _GotoState178Action,
	{_State242, SubToken}:                         _GotoState183Action,
	{_State242, MulToken}:                         _GotoState181Action,
	{_State243, IdentifierToken}:                  _GotoState80Action,
	{_State243, StructToken}:                      _GotoState34Action,
	{_State243, EnumToken}:                        _GotoState77Action,
	{_State243, TraitToken}:                       _GotoState85Action,
	{_State243, FuncToken}:                        _GotoState79Action,
	{_State243, LparenToken}:                      _GotoState81Action,
	{_State243, LbracketToken}:                    _GotoState26Action,
	{_State243, DotToken}:                         _GotoState76Action,
	{_State243, QuestionToken}:                    _GotoState83Action,
	{_State243, ExclaimToken}:                     _GotoState78Action,
	{_State243, TildeTildeToken}:                  _GotoState84Action,
	{_State243, BitNegToken}:                      _GotoState75Action,
	{_State243, BitAndToken}:                      _GotoState74Action,
	{_State243, ParseErrorToken}:                  _GotoState82Action,
	{_State243, InitializableTypeType}:            _GotoState91Action,
	{_State243, AtomTypeType}:                     _GotoState86Action,
	{_State243, ReturnableTypeType}:               _GotoState92Action,
	{_State243, ValueTypeType}:                    _GotoState344Action,
	{_State243, ImplicitStructDefType}:            _GotoState90Action,
	{_State243, ExplicitStructDefType}:            _GotoState45Action,
	{_State243, ImplicitEnumDefType}:              _GotoState89Action,
	{_State243, ExplicitEnumDefType}:              _GotoState87Action,
	{_State243, TraitDefType}:                     _GotoState93Action,
	{_State243, FuncTypeType}:                     _GotoState88Action,
	{_State245, CommaToken}:                       _GotoState345Action,
	{_State246, RbracketToken}:                    _GotoState346Action,
	{_State247, ImplementsToken}:                  _GotoState347Action,
	{_State247, AddToken}:                         _GotoState178Action,
	{_State247, SubToken}:                         _GotoState183Action,
	{_State247, MulToken}:                         _GotoState181Action,
	{_State249, IdentifierToken}:                  _GotoState153Action,
	{_State249, ParameterDefType}:                 _GotoState156Action,
	{_State249, ParameterDefsType}:                _GotoState157Action,
	{_State249, OptionalParameterDefsType}:        _GotoState348Action,
	{_State250, IdentifierToken}:                  _GotoState80Action,
	{_State250, StructToken}:                      _GotoState34Action,
	{_State250, EnumToken}:                        _GotoState77Action,
	{_State250, TraitToken}:                       _GotoState85Action,
	{_State250, FuncToken}:                        _GotoState79Action,
	{_State250, LparenToken}:                      _GotoState81Action,
	{_State250, LbracketToken}:                    _GotoState26Action,
	{_State250, DotToken}:                         _GotoState76Action,
	{_State250, QuestionToken}:                    _GotoState83Action,
	{_State250, ExclaimToken}:                     _GotoState78Action,
	{_State250, TildeTildeToken}:                  _GotoState84Action,
	{_State250, BitNegToken}:                      _GotoState75Action,
	{_State250, BitAndToken}:                      _GotoState74Action,
	{_State250, ParseErrorToken}:                  _GotoState82Action,
	{_State250, InitializableTypeType}:            _GotoState91Action,
	{_State250, AtomTypeType}:                     _GotoState86Action,
	{_State250, ReturnableTypeType}:               _GotoState92Action,
	{_State250, ValueTypeType}:                    _GotoState349Action,
	{_State250, ImplicitStructDefType}:            _GotoState90Action,
	{_State250, ExplicitStructDefType}:            _GotoState45Action,
	{_State250, ImplicitEnumDefType}:              _GotoState89Action,
	{_State250, ExplicitEnumDefType}:              _GotoState87Action,
	{_State250, TraitDefType}:                     _GotoState93Action,
	{_State250, FuncTypeType}:                     _GotoState88Action,
	{_State251, AddToken}:                         _GotoState178Action,
	{_State251, SubToken}:                         _GotoState183Action,
	{_State251, MulToken}:                         _GotoState181Action,
	{_State252, IdentifierToken}:                  _GotoState350Action,
	{_State253, IdentifierToken}:                  _GotoState80Action,
	{_State253, StructToken}:                      _GotoState34Action,
	{_State253, EnumToken}:                        _GotoState77Action,
	{_State253, TraitToken}:                       _GotoState85Action,
	{_State253, FuncToken}:                        _GotoState79Action,
	{_State253, LparenToken}:                      _GotoState81Action,
	{_State253, LbracketToken}:                    _GotoState26Action,
	{_State253, DotToken}:                         _GotoState76Action,
	{_State253, QuestionToken}:                    _GotoState83Action,
	{_State253, ExclaimToken}:                     _GotoState78Action,
	{_State253, TildeTildeToken}:                  _GotoState84Action,
	{_State253, BitNegToken}:                      _GotoState75Action,
	{_State253, BitAndToken}:                      _GotoState74Action,
	{_State253, ParseErrorToken}:                  _GotoState82Action,
	{_State253, InitializableTypeType}:            _GotoState91Action,
	{_State253, AtomTypeType}:                     _GotoState86Action,
	{_State253, ReturnableTypeType}:               _GotoState352Action,
	{_State253, ImplicitStructDefType}:            _GotoState90Action,
	{_State253, ExplicitStructDefType}:            _GotoState45Action,
	{_State253, ImplicitEnumDefType}:              _GotoState89Action,
	{_State253, ExplicitEnumDefType}:              _GotoState87Action,
	{_State253, TraitDefType}:                     _GotoState93Action,
	{_State253, ReturnTypeType}:                   _GotoState351Action,
	{_State253, FuncTypeType}:                     _GotoState88Action,
	{_State254, IdentifierToken}:                  _GotoState153Action,
	{_State254, ParameterDefType}:                 _GotoState353Action,
	{_State255, NewlinesToken}:                    _GotoState354Action,
	{_State255, OrToken}:                          _GotoState355Action,
	{_State256, RparenToken}:                      _GotoState356Action,
	{_State257, AssignToken}:                      _GotoState270Action,
	{_State258, NewlinesToken}:                    _GotoState357Action,
	{_State258, OrToken}:                          _GotoState358Action,
	{_State259, IdentifierToken}:                  _GotoState80Action,
	{_State259, StructToken}:                      _GotoState34Action,
	{_State259, EnumToken}:                        _GotoState77Action,
	{_State259, TraitToken}:                       _GotoState85Action,
	{_State259, FuncToken}:                        _GotoState79Action,
	{_State259, LparenToken}:                      _GotoState81Action,
	{_State259, LbracketToken}:                    _GotoState26Action,
	{_State259, DotToken}:                         _GotoState76Action,
	{_State259, QuestionToken}:                    _GotoState83Action,
	{_State259, ExclaimToken}:                     _GotoState78Action,
	{_State259, TildeTildeToken}:                  _GotoState84Action,
	{_State259, BitNegToken}:                      _GotoState75Action,
	{_State259, BitAndToken}:                      _GotoState74Action,
	{_State259, ParseErrorToken}:                  _GotoState82Action,
	{_State259, InitializableTypeType}:            _GotoState91Action,
	{_State259, AtomTypeType}:                     _GotoState86Action,
	{_State259, ReturnableTypeType}:               _GotoState92Action,
	{_State259, ValueTypeType}:                    _GotoState359Action,
	{_State259, ImplicitStructDefType}:            _GotoState90Action,
	{_State259, ExplicitStructDefType}:            _GotoState45Action,
	{_State259, ImplicitEnumDefType}:              _GotoState89Action,
	{_State259, ExplicitEnumDefType}:              _GotoState87Action,
	{_State259, TraitDefType}:                     _GotoState93Action,
	{_State259, FuncTypeType}:                     _GotoState88Action,
	{_State260, IdentifierToken}:                  _GotoState80Action,
	{_State260, StructToken}:                      _GotoState34Action,
	{_State260, EnumToken}:                        _GotoState77Action,
	{_State260, TraitToken}:                       _GotoState85Action,
	{_State260, FuncToken}:                        _GotoState79Action,
	{_State260, LparenToken}:                      _GotoState81Action,
	{_State260, LbracketToken}:                    _GotoState26Action,
	{_State260, DotToken}:                         _GotoState266Action,
	{_State260, QuestionToken}:                    _GotoState83Action,
	{_State260, ExclaimToken}:                     _GotoState78Action,
	{_State260, DollarLbracketToken}:              _GotoState104Action,
	{_State260, DotDotDotToken}:                   _GotoState360Action,
	{_State260, TildeTildeToken}:                  _GotoState84Action,
	{_State260, BitNegToken}:                      _GotoState75Action,
	{_State260, BitAndToken}:                      _GotoState74Action,
	{_State260, ParseErrorToken}:                  _GotoState82Action,
	{_State260, OptionalGenericBindingType}:       _GotoState165Action,
	{_State260, InitializableTypeType}:            _GotoState91Action,
	{_State260, AtomTypeType}:                     _GotoState86Action,
	{_State260, ReturnableTypeType}:               _GotoState92Action,
	{_State260, ValueTypeType}:                    _GotoState361Action,
	{_State260, ImplicitStructDefType}:            _GotoState90Action,
	{_State260, ExplicitStructDefType}:            _GotoState45Action,
	{_State260, ImplicitEnumDefType}:              _GotoState89Action,
	{_State260, ExplicitEnumDefType}:              _GotoState87Action,
	{_State260, TraitDefType}:                     _GotoState93Action,
	{_State260, FuncTypeType}:                     _GotoState88Action,
	{_State261, RparenToken}:                      _GotoState362Action,
	{_State263, CommaToken}:                       _GotoState363Action,
	{_State264, AddToken}:                         _GotoState178Action,
	{_State264, SubToken}:                         _GotoState183Action,
	{_State264, MulToken}:                         _GotoState181Action,
	{_State265, DollarLbracketToken}:              _GotoState104Action,
	{_State265, OptionalGenericBindingType}:       _GotoState364Action,
	{_State266, IdentifierToken}:                  _GotoState265Action,
	{_State266, DollarLbracketToken}:              _GotoState104Action,
	{_State266, OptionalGenericBindingType}:       _GotoState160Action,
	{_State267, AddToken}:                         _GotoState178Action,
	{_State267, SubToken}:                         _GotoState183Action,
	{_State267, MulToken}:                         _GotoState181Action,
	{_State268, IdentifierToken}:                  _GotoState365Action,
	{_State269, IdentifierToken}:                  _GotoState166Action,
	{_State269, UnsafeToken}:                      _GotoState167Action,
	{_State269, StructToken}:                      _GotoState34Action,
	{_State269, EnumToken}:                        _GotoState77Action,
	{_State269, TraitToken}:                       _GotoState85Action,
	{_State269, FuncToken}:                        _GotoState79Action,
	{_State269, LparenToken}:                      _GotoState81Action,
	{_State269, LbracketToken}:                    _GotoState26Action,
	{_State269, DotToken}:                         _GotoState76Action,
	{_State269, QuestionToken}:                    _GotoState83Action,
	{_State269, ExclaimToken}:                     _GotoState78Action,
	{_State269, TildeTildeToken}:                  _GotoState84Action,
	{_State269, BitNegToken}:                      _GotoState75Action,
	{_State269, BitAndToken}:                      _GotoState74Action,
	{_State269, ParseErrorToken}:                  _GotoState82Action,
	{_State269, UnsafeStatementType}:              _GotoState173Action,
	{_State269, InitializableTypeType}:            _GotoState91Action,
	{_State269, AtomTypeType}:                     _GotoState86Action,
	{_State269, ReturnableTypeType}:               _GotoState92Action,
	{_State269, ValueTypeType}:                    _GotoState174Action,
	{_State269, FieldDefType}:                     _GotoState257Action,
	{_State269, ImplicitStructDefType}:            _GotoState90Action,
	{_State269, ExplicitStructDefType}:            _GotoState45Action,
	{_State269, EnumValueDefType}:                 _GotoState366Action,
	{_State269, ImplicitEnumDefType}:              _GotoState89Action,
	{_State269, ExplicitEnumDefType}:              _GotoState87Action,
	{_State269, TraitDefType}:                     _GotoState93Action,
	{_State269, FuncTypeType}:                     _GotoState88Action,
	{_State270, DefaultToken}:                     _GotoState367Action,
	{_State271, IdentifierToken}:                  _GotoState166Action,
	{_State271, UnsafeToken}:                      _GotoState167Action,
	{_State271, StructToken}:                      _GotoState34Action,
	{_State271, EnumToken}:                        _GotoState77Action,
	{_State271, TraitToken}:                       _GotoState85Action,
	{_State271, FuncToken}:                        _GotoState79Action,
	{_State271, LparenToken}:                      _GotoState81Action,
	{_State271, LbracketToken}:                    _GotoState26Action,
	{_State271, DotToken}:                         _GotoState76Action,
	{_State271, QuestionToken}:                    _GotoState83Action,
	{_State271, ExclaimToken}:                     _GotoState78Action,
	{_State271, TildeTildeToken}:                  _GotoState84Action,
	{_State271, BitNegToken}:                      _GotoState75Action,
	{_State271, BitAndToken}:                      _GotoState74Action,
	{_State271, ParseErrorToken}:                  _GotoState82Action,
	{_State271, UnsafeStatementType}:              _GotoState173Action,
	{_State271, InitializableTypeType}:            _GotoState91Action,
	{_State271, AtomTypeType}:                     _GotoState86Action,
	{_State271, ReturnableTypeType}:               _GotoState92Action,
	{_State271, ValueTypeType}:                    _GotoState174Action,
	{_State271, FieldDefType}:                     _GotoState257Action,
	{_State271, ImplicitStructDefType}:            _GotoState90Action,
	{_State271, ExplicitStructDefType}:            _GotoState45Action,
	{_State271, EnumValueDefType}:                 _GotoState368Action,
	{_State271, ImplicitEnumDefType}:              _GotoState89Action,
	{_State271, ExplicitEnumDefType}:              _GotoState87Action,
	{_State271, TraitDefType}:                     _GotoState93Action,
	{_State271, FuncTypeType}:                     _GotoState88Action,
	{_State273, IdentifierToken}:                  _GotoState166Action,
	{_State273, UnsafeToken}:                      _GotoState167Action,
	{_State273, StructToken}:                      _GotoState34Action,
	{_State273, EnumToken}:                        _GotoState77Action,
	{_State273, TraitToken}:                       _GotoState85Action,
	{_State273, FuncToken}:                        _GotoState79Action,
	{_State273, LparenToken}:                      _GotoState81Action,
	{_State273, LbracketToken}:                    _GotoState26Action,
	{_State273, DotToken}:                         _GotoState76Action,
	{_State273, QuestionToken}:                    _GotoState83Action,
	{_State273, ExclaimToken}:                     _GotoState78Action,
	{_State273, TildeTildeToken}:                  _GotoState84Action,
	{_State273, BitNegToken}:                      _GotoState75Action,
	{_State273, BitAndToken}:                      _GotoState74Action,
	{_State273, ParseErrorToken}:                  _GotoState82Action,
	{_State273, UnsafeStatementType}:              _GotoState173Action,
	{_State273, InitializableTypeType}:            _GotoState91Action,
	{_State273, AtomTypeType}:                     _GotoState86Action,
	{_State273, ReturnableTypeType}:               _GotoState92Action,
	{_State273, ValueTypeType}:                    _GotoState174Action,
	{_State273, FieldDefType}:                     _GotoState369Action,
	{_State273, ImplicitStructDefType}:            _GotoState90Action,
	{_State273, ExplicitStructDefType}:            _GotoState45Action,
	{_State273, ImplicitEnumDefType}:              _GotoState89Action,
	{_State273, ExplicitEnumDefType}:              _GotoState87Action,
	{_State273, TraitDefType}:                     _GotoState93Action,
	{_State273, FuncTypeType}:                     _GotoState88Action,
	{_State275, IdentifierToken}:                  _GotoState370Action,
	{_State275, LparenToken}:                      _GotoState163Action,
	{_State278, RparenToken}:                      _GotoState371Action,
	{_State279, NewlinesToken}:                    _GotoState373Action,
	{_State279, CommaToken}:                       _GotoState372Action,
	{_State282, RbracketToken}:                    _GotoState374Action,
	{_State282, AddToken}:                         _GotoState178Action,
	{_State282, SubToken}:                         _GotoState183Action,
	{_State282, MulToken}:                         _GotoState181Action,
	{_State283, RbracketToken}:                    _GotoState375Action,
	{_State291, IdentifierToken}:                  _GotoState166Action,
	{_State291, UnsafeToken}:                      _GotoState167Action,
	{_State291, StructToken}:                      _GotoState34Action,
	{_State291, EnumToken}:                        _GotoState77Action,
	{_State291, TraitToken}:                       _GotoState85Action,
	{_State291, FuncToken}:                        _GotoState79Action,
	{_State291, LparenToken}:                      _GotoState81Action,
	{_State291, LbracketToken}:                    _GotoState26Action,
	{_State291, DotToken}:                         _GotoState76Action,
	{_State291, QuestionToken}:                    _GotoState83Action,
	{_State291, ExclaimToken}:                     _GotoState78Action,
	{_State291, TildeTildeToken}:                  _GotoState84Action,
	{_State291, BitNegToken}:                      _GotoState75Action,
	{_State291, BitAndToken}:                      _GotoState74Action,
	{_State291, ParseErrorToken}:                  _GotoState82Action,
	{_State291, UnsafeStatementType}:              _GotoState173Action,
	{_State291, InitializableTypeType}:            _GotoState91Action,
	{_State291, AtomTypeType}:                     _GotoState86Action,
	{_State291, ReturnableTypeType}:               _GotoState92Action,
	{_State291, ValueTypeType}:                    _GotoState174Action,
	{_State291, FieldDefType}:                     _GotoState376Action,
	{_State291, ImplicitStructDefType}:            _GotoState90Action,
	{_State291, ExplicitStructDefType}:            _GotoState45Action,
	{_State291, ImplicitEnumDefType}:              _GotoState89Action,
	{_State291, ExplicitEnumDefType}:              _GotoState87Action,
	{_State291, TraitDefType}:                     _GotoState93Action,
	{_State291, FuncTypeType}:                     _GotoState88Action,
	{_State292, IdentifierToken}:                  _GotoState166Action,
	{_State292, UnsafeToken}:                      _GotoState167Action,
	{_State292, StructToken}:                      _GotoState34Action,
	{_State292, EnumToken}:                        _GotoState77Action,
	{_State292, TraitToken}:                       _GotoState85Action,
	{_State292, FuncToken}:                        _GotoState79Action,
	{_State292, LparenToken}:                      _GotoState81Action,
	{_State292, LbracketToken}:                    _GotoState26Action,
	{_State292, DotToken}:                         _GotoState76Action,
	{_State292, QuestionToken}:                    _GotoState83Action,
	{_State292, ExclaimToken}:                     _GotoState78Action,
	{_State292, TildeTildeToken}:                  _GotoState84Action,
	{_State292, BitNegToken}:                      _GotoState75Action,
	{_State292, BitAndToken}:                      _GotoState74Action,
	{_State292, ParseErrorToken}:                  _GotoState82Action,
	{_State292, UnsafeStatementType}:              _GotoState173Action,
	{_State292, InitializableTypeType}:            _GotoState91Action,
	{_State292, AtomTypeType}:                     _GotoState86Action,
	{_State292, ReturnableTypeType}:               _GotoState92Action,
	{_State292, ValueTypeType}:                    _GotoState174Action,
	{_State292, FieldDefType}:                     _GotoState377Action,
	{_State292, ImplicitStructDefType}:            _GotoState90Action,
	{_State292, ExplicitStructDefType}:            _GotoState45Action,
	{_State292, ImplicitEnumDefType}:              _GotoState89Action,
	{_State292, ExplicitEnumDefType}:              _GotoState87Action,
	{_State292, TraitDefType}:                     _GotoState93Action,
	{_State292, FuncTypeType}:                     _GotoState88Action,
	{_State294, IdentifierToken}:                  _GotoState80Action,
	{_State294, StructToken}:                      _GotoState34Action,
	{_State294, EnumToken}:                        _GotoState77Action,
	{_State294, TraitToken}:                       _GotoState85Action,
	{_State294, FuncToken}:                        _GotoState79Action,
	{_State294, LparenToken}:                      _GotoState81Action,
	{_State294, LbracketToken}:                    _GotoState26Action,
	{_State294, DotToken}:                         _GotoState76Action,
	{_State294, QuestionToken}:                    _GotoState83Action,
	{_State294, ExclaimToken}:                     _GotoState78Action,
	{_State294, TildeTildeToken}:                  _GotoState84Action,
	{_State294, BitNegToken}:                      _GotoState75Action,
	{_State294, BitAndToken}:                      _GotoState74Action,
	{_State294, ParseErrorToken}:                  _GotoState82Action,
	{_State294, InitializableTypeType}:            _GotoState91Action,
	{_State294, AtomTypeType}:                     _GotoState86Action,
	{_State294, ReturnableTypeType}:               _GotoState92Action,
	{_State294, ValueTypeType}:                    _GotoState378Action,
	{_State294, ImplicitStructDefType}:            _GotoState90Action,
	{_State294, ExplicitStructDefType}:            _GotoState45Action,
	{_State294, ImplicitEnumDefType}:              _GotoState89Action,
	{_State294, ExplicitEnumDefType}:              _GotoState87Action,
	{_State294, TraitDefType}:                     _GotoState93Action,
	{_State294, FuncTypeType}:                     _GotoState88Action,
	{_State297, RparenToken}:                      _GotoState379Action,
	{_State299, IntegerLiteralToken}:              _GotoState24Action,
	{_State299, FloatLiteralToken}:                _GotoState20Action,
	{_State299, RuneLiteralToken}:                 _GotoState32Action,
	{_State299, StringLiteralToken}:               _GotoState33Action,
	{_State299, IdentifierToken}:                  _GotoState23Action,
	{_State299, TrueToken}:                        _GotoState36Action,
	{_State299, FalseToken}:                       _GotoState19Action,
	{_State299, StructToken}:                      _GotoState34Action,
	{_State299, FuncToken}:                        _GotoState21Action,
	{_State299, VarToken}:                         _GotoState37Action,
	{_State299, LetToken}:                         _GotoState27Action,
	{_State299, NotToken}:                         _GotoState30Action,
	{_State299, LabelDeclToken}:                   _GotoState25Action,
	{_State299, LparenToken}:                      _GotoState28Action,
	{_State299, LbracketToken}:                    _GotoState26Action,
	{_State299, SubToken}:                         _GotoState35Action,
	{_State299, MulToken}:                         _GotoState29Action,
	{_State299, BitNegToken}:                      _GotoState18Action,
	{_State299, BitAndToken}:                      _GotoState17Action,
	{_State299, GreaterToken}:                     _GotoState22Action,
	{_State299, ParseErrorToken}:                  _GotoState31Action,
	{_State299, VarDeclPatternType}:               _GotoState56Action,
	{_State299, VarOrLetType}:                     _GotoState57Action,
	{_State299, OptionalLabelDeclType}:            _GotoState72Action,
	{_State299, SequenceExprType}:                 _GotoState380Action,
	{_State299, CallExprType}:                     _GotoState43Action,
	{_State299, AtomExprType}:                     _GotoState42Action,
	{_State299, LiteralExprType}:                  _GotoState48Action,
	{_State299, ImplicitStructExprType}:           _GotoState46Action,
	{_State299, AccessExprType}:                   _GotoState38Action,
	{_State299, PostfixUnaryExprType}:             _GotoState52Action,
	{_State299, PrefixUnaryOpType}:                _GotoState54Action,
	{_State299, PrefixUnaryExprType}:              _GotoState53Action,
	{_State299, MulExprType}:                      _GotoState49Action,
	{_State299, AddExprType}:                      _GotoState39Action,
	{_State299, CmpExprType}:                      _GotoState44Action,
	{_State299, AndExprType}:                      _GotoState40Action,
	{_State299, OrExprType}:                       _GotoState51Action,
	{_State299, InitializableTypeType}:            _GotoState47Action,
	{_State299, ExplicitStructDefType}:            _GotoState45Action,
	{_State299, AnonymousFuncExprType}:            _GotoState41Action,
	{_State300, IntegerLiteralToken}:              _GotoState24Action,
	{_State300, FloatLiteralToken}:                _GotoState20Action,
	{_State300, RuneLiteralToken}:                 _GotoState32Action,
	{_State300, StringLiteralToken}:               _GotoState33Action,
	{_State300, IdentifierToken}:                  _GotoState23Action,
	{_State300, TrueToken}:                        _GotoState36Action,
	{_State300, FalseToken}:                       _GotoState19Action,
	{_State300, StructToken}:                      _GotoState34Action,
	{_State300, FuncToken}:                        _GotoState21Action,
	{_State300, VarToken}:                         _GotoState37Action,
	{_State300, LetToken}:                         _GotoState27Action,
	{_State300, NotToken}:                         _GotoState30Action,
	{_State300, LabelDeclToken}:                   _GotoState25Action,
	{_State300, LparenToken}:                      _GotoState28Action,
	{_State300, LbracketToken}:                    _GotoState26Action,
	{_State300, SubToken}:                         _GotoState35Action,
	{_State300, MulToken}:                         _GotoState29Action,
	{_State300, BitNegToken}:                      _GotoState18Action,
	{_State300, BitAndToken}:                      _GotoState17Action,
	{_State300, GreaterToken}:                     _GotoState22Action,
	{_State300, ParseErrorToken}:                  _GotoState31Action,
	{_State300, VarDeclPatternType}:               _GotoState56Action,
	{_State300, VarOrLetType}:                     _GotoState57Action,
	{_State300, OptionalLabelDeclType}:            _GotoState72Action,
	{_State300, SequenceExprType}:                 _GotoState381Action,
	{_State300, CallExprType}:                     _GotoState43Action,
	{_State300, AtomExprType}:                     _GotoState42Action,
	{_State300, LiteralExprType}:                  _GotoState48Action,
	{_State300, ImplicitStructExprType}:           _GotoState46Action,
	{_State300, AccessExprType}:                   _GotoState38Action,
	{_State300, PostfixUnaryExprType}:             _GotoState52Action,
	{_State300, PrefixUnaryOpType}:                _GotoState54Action,
	{_State300, PrefixUnaryExprType}:              _GotoState53Action,
	{_State300, MulExprType}:                      _GotoState49Action,
	{_State300, AddExprType}:                      _GotoState39Action,
	{_State300, CmpExprType}:                      _GotoState44Action,
	{_State300, AndExprType}:                      _GotoState40Action,
	{_State300, OrExprType}:                       _GotoState51Action,
	{_State300, InitializableTypeType}:            _GotoState47Action,
	{_State300, ExplicitStructDefType}:            _GotoState45Action,
	{_State300, AnonymousFuncExprType}:            _GotoState41Action,
	{_State301, IntegerLiteralToken}:              _GotoState24Action,
	{_State301, FloatLiteralToken}:                _GotoState20Action,
	{_State301, RuneLiteralToken}:                 _GotoState32Action,
	{_State301, StringLiteralToken}:               _GotoState33Action,
	{_State301, IdentifierToken}:                  _GotoState23Action,
	{_State301, TrueToken}:                        _GotoState36Action,
	{_State301, FalseToken}:                       _GotoState19Action,
	{_State301, StructToken}:                      _GotoState34Action,
	{_State301, FuncToken}:                        _GotoState21Action,
	{_State301, VarToken}:                         _GotoState37Action,
	{_State301, LetToken}:                         _GotoState27Action,
	{_State301, NotToken}:                         _GotoState30Action,
	{_State301, LabelDeclToken}:                   _GotoState25Action,
	{_State301, LparenToken}:                      _GotoState28Action,
	{_State301, LbracketToken}:                    _GotoState26Action,
	{_State301, SubToken}:                         _GotoState35Action,
	{_State301, MulToken}:                         _GotoState29Action,
	{_State301, BitNegToken}:                      _GotoState18Action,
	{_State301, BitAndToken}:                      _GotoState17Action,
	{_State301, GreaterToken}:                     _GotoState22Action,
	{_State301, ParseErrorToken}:                  _GotoState31Action,
	{_State301, VarDeclPatternType}:               _GotoState56Action,
	{_State301, VarOrLetType}:                     _GotoState57Action,
	{_State301, OptionalLabelDeclType}:            _GotoState72Action,
	{_State301, SequenceExprType}:                 _GotoState382Action,
	{_State301, CallExprType}:                     _GotoState43Action,
	{_State301, AtomExprType}:                     _GotoState42Action,
	{_State301, LiteralExprType}:                  _GotoState48Action,
	{_State301, ImplicitStructExprType}:           _GotoState46Action,
	{_State301, AccessExprType}:                   _GotoState38Action,
	{_State301, PostfixUnaryExprType}:             _GotoState52Action,
	{_State301, PrefixUnaryOpType}:                _GotoState54Action,
	{_State301, PrefixUnaryExprType}:              _GotoState53Action,
	{_State301, MulExprType}:                      _GotoState49Action,
	{_State301, AddExprType}:                      _GotoState39Action,
	{_State301, CmpExprType}:                      _GotoState44Action,
	{_State301, AndExprType}:                      _GotoState40Action,
	{_State301, OrExprType}:                       _GotoState51Action,
	{_State301, InitializableTypeType}:            _GotoState47Action,
	{_State301, ExplicitStructDefType}:            _GotoState45Action,
	{_State301, AnonymousFuncExprType}:            _GotoState41Action,
	{_State302, IntegerLiteralToken}:              _GotoState24Action,
	{_State302, FloatLiteralToken}:                _GotoState20Action,
	{_State302, RuneLiteralToken}:                 _GotoState32Action,
	{_State302, StringLiteralToken}:               _GotoState33Action,
	{_State302, IdentifierToken}:                  _GotoState23Action,
	{_State302, TrueToken}:                        _GotoState36Action,
	{_State302, FalseToken}:                       _GotoState19Action,
	{_State302, StructToken}:                      _GotoState34Action,
	{_State302, FuncToken}:                        _GotoState21Action,
	{_State302, VarToken}:                         _GotoState37Action,
	{_State302, LetToken}:                         _GotoState27Action,
	{_State302, NotToken}:                         _GotoState30Action,
	{_State302, LabelDeclToken}:                   _GotoState25Action,
	{_State302, LparenToken}:                      _GotoState28Action,
	{_State302, LbracketToken}:                    _GotoState26Action,
	{_State302, SubToken}:                         _GotoState35Action,
	{_State302, MulToken}:                         _GotoState29Action,
	{_State302, BitNegToken}:                      _GotoState18Action,
	{_State302, BitAndToken}:                      _GotoState17Action,
	{_State302, GreaterToken}:                     _GotoState22Action,
	{_State302, ParseErrorToken}:                  _GotoState31Action,
	{_State302, VarDeclPatternType}:               _GotoState56Action,
	{_State302, VarOrLetType}:                     _GotoState57Action,
	{_State302, OptionalLabelDeclType}:            _GotoState72Action,
	{_State302, SequenceExprType}:                 _GotoState384Action,
	{_State302, OptionalSequenceExprType}:         _GotoState383Action,
	{_State302, CallExprType}:                     _GotoState43Action,
	{_State302, AtomExprType}:                     _GotoState42Action,
	{_State302, LiteralExprType}:                  _GotoState48Action,
	{_State302, ImplicitStructExprType}:           _GotoState46Action,
	{_State302, AccessExprType}:                   _GotoState38Action,
	{_State302, PostfixUnaryExprType}:             _GotoState52Action,
	{_State302, PrefixUnaryOpType}:                _GotoState54Action,
	{_State302, PrefixUnaryExprType}:              _GotoState53Action,
	{_State302, MulExprType}:                      _GotoState49Action,
	{_State302, AddExprType}:                      _GotoState39Action,
	{_State302, CmpExprType}:                      _GotoState44Action,
	{_State302, AndExprType}:                      _GotoState40Action,
	{_State302, OrExprType}:                       _GotoState51Action,
	{_State302, InitializableTypeType}:            _GotoState47Action,
	{_State302, ExplicitStructDefType}:            _GotoState45Action,
	{_State302, AnonymousFuncExprType}:            _GotoState41Action,
	{_State303, LbraceToken}:                      _GotoState59Action,
	{_State303, StatementBlockType}:               _GotoState385Action,
	{_State304, IdentifierToken}:                  _GotoState386Action,
	{_State305, DotToken}:                         _GotoState387Action,
	{_State308, CommaToken}:                       _GotoState389Action,
	{_State308, AssignToken}:                      _GotoState388Action,
	{_State310, ElseToken}:                        _GotoState390Action,
	{_State312, IdentifierToken}:                  _GotoState140Action,
	{_State312, LparenToken}:                      _GotoState141Action,
	{_State312, VarPatternType}:                   _GotoState391Action,
	{_State312, TuplePatternType}:                 _GotoState142Action,
	{_State313, IdentifierToken}:                  _GotoState213Action,
	{_State313, LparenToken}:                      _GotoState141Action,
	{_State313, DotDotDotToken}:                   _GotoState212Action,
	{_State313, VarPatternType}:                   _GotoState216Action,
	{_State313, TuplePatternType}:                 _GotoState142Action,
	{_State313, FieldVarPatternType}:              _GotoState392Action,
	{_State315, LbracketToken}:                    _GotoState106Action,
	{_State315, DotToken}:                         _GotoState105Action,
	{_State315, DollarLbracketToken}:              _GotoState104Action,
	{_State315, OptionalGenericBindingType}:       _GotoState108Action,
	{_State317, CommaToken}:                       _GotoState389Action,
	{_State317, ColonToken}:                       _GotoState393Action,
	{_State318, IntegerLiteralToken}:              _GotoState24Action,
	{_State318, FloatLiteralToken}:                _GotoState20Action,
	{_State318, RuneLiteralToken}:                 _GotoState32Action,
	{_State318, StringLiteralToken}:               _GotoState33Action,
	{_State318, IdentifierToken}:                  _GotoState23Action,
	{_State318, TrueToken}:                        _GotoState36Action,
	{_State318, FalseToken}:                       _GotoState19Action,
	{_State318, ReturnToken}:                      _GotoState228Action,
	{_State318, BreakToken}:                       _GotoState220Action,
	{_State318, ContinueToken}:                    _GotoState222Action,
	{_State318, FallthroughToken}:                 _GotoState225Action,
	{_State318, UnsafeToken}:                      _GotoState167Action,
	{_State318, StructToken}:                      _GotoState34Action,
	{_State318, FuncToken}:                        _GotoState21Action,
	{_State318, AsyncToken}:                       _GotoState219Action,
	{_State318, DeferToken}:                       _GotoState224Action,
	{_State318, VarToken}:                         _GotoState37Action,
	{_State318, LetToken}:                         _GotoState27Action,
	{_State318, NotToken}:                         _GotoState30Action,
	{_State318, LabelDeclToken}:                   _GotoState25Action,
	{_State318, LparenToken}:                      _GotoState28Action,
	{_State318, LbracketToken}:                    _GotoState26Action,
	{_State318, SubToken}:                         _GotoState35Action,
	{_State318, MulToken}:                         _GotoState29Action,
	{_State318, BitNegToken}:                      _GotoState18Action,
	{_State318, BitAndToken}:                      _GotoState17Action,
	{_State318, GreaterToken}:                     _GotoState22Action,
	{_State318, ParseErrorToken}:                  _GotoState31Action,
	{_State318, SimpleStatementBodyType}:          _GotoState395Action,
	{_State318, OptionalSimpleStatementBodyType}:  _GotoState394Action,
	{_State318, UnsafeStatementType}:              _GotoState239Action,
	{_State318, JumpTypeType}:                     _GotoState234Action,
	{_State318, ExpressionsType}:                  _GotoState232Action,
	{_State318, VarDeclPatternType}:               _GotoState56Action,
	{_State318, VarOrLetType}:                     _GotoState57Action,
	{_State318, AssignPatternType}:                _GotoState230Action,
	{_State318, ExpressionType}:                   _GotoState231Action,
	{_State318, OptionalLabelDeclType}:            _GotoState50Action,
	{_State318, SequenceExprType}:                 _GotoState235Action,
	{_State318, CallExprType}:                     _GotoState43Action,
	{_State318, AtomExprType}:                     _GotoState42Action,
	{_State318, LiteralExprType}:                  _GotoState48Action,
	{_State318, ImplicitStructExprType}:           _GotoState46Action,
	{_State318, AccessExprType}:                   _GotoState229Action,
	{_State318, PostfixUnaryExprType}:             _GotoState52Action,
	{_State318, PrefixUnaryOpType}:                _GotoState54Action,
	{_State318, PrefixUnaryExprType}:              _GotoState53Action,
	{_State318, MulExprType}:                      _GotoState49Action,
	{_State318, AddExprType}:                      _GotoState39Action,
	{_State318, CmpExprType}:                      _GotoState44Action,
	{_State318, AndExprType}:                      _GotoState40Action,
	{_State318, OrExprType}:                       _GotoState51Action,
	{_State318, InitializableTypeType}:            _GotoState47Action,
	{_State318, ExplicitStructDefType}:            _GotoState45Action,
	{_State318, AnonymousFuncExprType}:            _GotoState41Action,
	{_State320, StringLiteralToken}:               _GotoState321Action,
	{_State320, ImportClauseType}:                 _GotoState396Action,
	{_State320, ImportClauseTerminalType}:         _GotoState397Action,
	{_State320, ImportClausesType}:                _GotoState398Action,
	{_State321, AsToken}:                          _GotoState399Action,
	{_State336, IntegerLiteralToken}:              _GotoState24Action,
	{_State336, FloatLiteralToken}:                _GotoState20Action,
	{_State336, RuneLiteralToken}:                 _GotoState32Action,
	{_State336, StringLiteralToken}:               _GotoState33Action,
	{_State336, IdentifierToken}:                  _GotoState23Action,
	{_State336, TrueToken}:                        _GotoState36Action,
	{_State336, FalseToken}:                       _GotoState19Action,
	{_State336, StructToken}:                      _GotoState34Action,
	{_State336, FuncToken}:                        _GotoState21Action,
	{_State336, VarToken}:                         _GotoState37Action,
	{_State336, LetToken}:                         _GotoState27Action,
	{_State336, NotToken}:                         _GotoState30Action,
	{_State336, LabelDeclToken}:                   _GotoState25Action,
	{_State336, LparenToken}:                      _GotoState28Action,
	{_State336, LbracketToken}:                    _GotoState26Action,
	{_State336, SubToken}:                         _GotoState35Action,
	{_State336, MulToken}:                         _GotoState29Action,
	{_State336, BitNegToken}:                      _GotoState18Action,
	{_State336, BitAndToken}:                      _GotoState17Action,
	{_State336, GreaterToken}:                     _GotoState22Action,
	{_State336, ParseErrorToken}:                  _GotoState31Action,
	{_State336, VarDeclPatternType}:               _GotoState56Action,
	{_State336, VarOrLetType}:                     _GotoState57Action,
	{_State336, ExpressionType}:                   _GotoState400Action,
	{_State336, OptionalLabelDeclType}:            _GotoState50Action,
	{_State336, SequenceExprType}:                 _GotoState55Action,
	{_State336, CallExprType}:                     _GotoState43Action,
	{_State336, AtomExprType}:                     _GotoState42Action,
	{_State336, LiteralExprType}:                  _GotoState48Action,
	{_State336, ImplicitStructExprType}:           _GotoState46Action,
	{_State336, AccessExprType}:                   _GotoState38Action,
	{_State336, PostfixUnaryExprType}:             _GotoState52Action,
	{_State336, PrefixUnaryOpType}:                _GotoState54Action,
	{_State336, PrefixUnaryExprType}:              _GotoState53Action,
	{_State336, MulExprType}:                      _GotoState49Action,
	{_State336, AddExprType}:                      _GotoState39Action,
	{_State336, CmpExprType}:                      _GotoState44Action,
	{_State336, AndExprType}:                      _GotoState40Action,
	{_State336, OrExprType}:                       _GotoState51Action,
	{_State336, InitializableTypeType}:            _GotoState47Action,
	{_State336, ExplicitStructDefType}:            _GotoState45Action,
	{_State336, AnonymousFuncExprType}:            _GotoState41Action,
	{_State338, IntegerLiteralToken}:              _GotoState24Action,
	{_State338, FloatLiteralToken}:                _GotoState20Action,
	{_State338, RuneLiteralToken}:                 _GotoState32Action,
	{_State338, StringLiteralToken}:               _GotoState33Action,
	{_State338, IdentifierToken}:                  _GotoState23Action,
	{_State338, TrueToken}:                        _GotoState36Action,
	{_State338, FalseToken}:                       _GotoState19Action,
	{_State338, StructToken}:                      _GotoState34Action,
	{_State338, FuncToken}:                        _GotoState21Action,
	{_State338, VarToken}:                         _GotoState37Action,
	{_State338, LetToken}:                         _GotoState27Action,
	{_State338, NotToken}:                         _GotoState30Action,
	{_State338, LabelDeclToken}:                   _GotoState25Action,
	{_State338, LparenToken}:                      _GotoState28Action,
	{_State338, LbracketToken}:                    _GotoState26Action,
	{_State338, SubToken}:                         _GotoState35Action,
	{_State338, MulToken}:                         _GotoState29Action,
	{_State338, BitNegToken}:                      _GotoState18Action,
	{_State338, BitAndToken}:                      _GotoState17Action,
	{_State338, GreaterToken}:                     _GotoState22Action,
	{_State338, ParseErrorToken}:                  _GotoState31Action,
	{_State338, VarDeclPatternType}:               _GotoState56Action,
	{_State338, VarOrLetType}:                     _GotoState57Action,
	{_State338, ExpressionType}:                   _GotoState401Action,
	{_State338, OptionalLabelDeclType}:            _GotoState50Action,
	{_State338, SequenceExprType}:                 _GotoState55Action,
	{_State338, CallExprType}:                     _GotoState43Action,
	{_State338, AtomExprType}:                     _GotoState42Action,
	{_State338, LiteralExprType}:                  _GotoState48Action,
	{_State338, ImplicitStructExprType}:           _GotoState46Action,
	{_State338, AccessExprType}:                   _GotoState38Action,
	{_State338, PostfixUnaryExprType}:             _GotoState52Action,
	{_State338, PrefixUnaryOpType}:                _GotoState54Action,
	{_State338, PrefixUnaryExprType}:              _GotoState53Action,
	{_State338, MulExprType}:                      _GotoState49Action,
	{_State338, AddExprType}:                      _GotoState39Action,
	{_State338, CmpExprType}:                      _GotoState44Action,
	{_State338, AndExprType}:                      _GotoState40Action,
	{_State338, OrExprType}:                       _GotoState51Action,
	{_State338, InitializableTypeType}:            _GotoState47Action,
	{_State338, ExplicitStructDefType}:            _GotoState45Action,
	{_State338, AnonymousFuncExprType}:            _GotoState41Action,
	{_State339, IntegerLiteralToken}:              _GotoState24Action,
	{_State339, FloatLiteralToken}:                _GotoState20Action,
	{_State339, RuneLiteralToken}:                 _GotoState32Action,
	{_State339, StringLiteralToken}:               _GotoState33Action,
	{_State339, IdentifierToken}:                  _GotoState23Action,
	{_State339, TrueToken}:                        _GotoState36Action,
	{_State339, FalseToken}:                       _GotoState19Action,
	{_State339, StructToken}:                      _GotoState34Action,
	{_State339, FuncToken}:                        _GotoState21Action,
	{_State339, VarToken}:                         _GotoState37Action,
	{_State339, LetToken}:                         _GotoState27Action,
	{_State339, NotToken}:                         _GotoState30Action,
	{_State339, LabelDeclToken}:                   _GotoState25Action,
	{_State339, LparenToken}:                      _GotoState28Action,
	{_State339, LbracketToken}:                    _GotoState26Action,
	{_State339, SubToken}:                         _GotoState35Action,
	{_State339, MulToken}:                         _GotoState29Action,
	{_State339, BitNegToken}:                      _GotoState18Action,
	{_State339, BitAndToken}:                      _GotoState17Action,
	{_State339, GreaterToken}:                     _GotoState22Action,
	{_State339, ParseErrorToken}:                  _GotoState31Action,
	{_State339, VarDeclPatternType}:               _GotoState56Action,
	{_State339, VarOrLetType}:                     _GotoState57Action,
	{_State339, ExpressionType}:                   _GotoState402Action,
	{_State339, OptionalLabelDeclType}:            _GotoState50Action,
	{_State339, SequenceExprType}:                 _GotoState55Action,
	{_State339, CallExprType}:                     _GotoState43Action,
	{_State339, AtomExprType}:                     _GotoState42Action,
	{_State339, LiteralExprType}:                  _GotoState48Action,
	{_State339, ImplicitStructExprType}:           _GotoState46Action,
	{_State339, AccessExprType}:                   _GotoState38Action,
	{_State339, PostfixUnaryExprType}:             _GotoState52Action,
	{_State339, PrefixUnaryOpType}:                _GotoState54Action,
	{_State339, PrefixUnaryExprType}:              _GotoState53Action,
	{_State339, MulExprType}:                      _GotoState49Action,
	{_State339, AddExprType}:                      _GotoState39Action,
	{_State339, CmpExprType}:                      _GotoState44Action,
	{_State339, AndExprType}:                      _GotoState40Action,
	{_State339, OrExprType}:                       _GotoState51Action,
	{_State339, InitializableTypeType}:            _GotoState47Action,
	{_State339, ExplicitStructDefType}:            _GotoState45Action,
	{_State339, AnonymousFuncExprType}:            _GotoState41Action,
	{_State341, IntegerLiteralToken}:              _GotoState24Action,
	{_State341, FloatLiteralToken}:                _GotoState20Action,
	{_State341, RuneLiteralToken}:                 _GotoState32Action,
	{_State341, StringLiteralToken}:               _GotoState33Action,
	{_State341, IdentifierToken}:                  _GotoState23Action,
	{_State341, TrueToken}:                        _GotoState36Action,
	{_State341, FalseToken}:                       _GotoState19Action,
	{_State341, StructToken}:                      _GotoState34Action,
	{_State341, FuncToken}:                        _GotoState21Action,
	{_State341, VarToken}:                         _GotoState37Action,
	{_State341, LetToken}:                         _GotoState27Action,
	{_State341, NotToken}:                         _GotoState30Action,
	{_State341, LabelDeclToken}:                   _GotoState25Action,
	{_State341, LparenToken}:                      _GotoState28Action,
	{_State341, LbracketToken}:                    _GotoState26Action,
	{_State341, SubToken}:                         _GotoState35Action,
	{_State341, MulToken}:                         _GotoState29Action,
	{_State341, BitNegToken}:                      _GotoState18Action,
	{_State341, BitAndToken}:                      _GotoState17Action,
	{_State341, GreaterToken}:                     _GotoState22Action,
	{_State341, ParseErrorToken}:                  _GotoState31Action,
	{_State341, ExpressionsType}:                  _GotoState403Action,
	{_State341, OptionalExpressionsType}:          _GotoState404Action,
	{_State341, VarDeclPatternType}:               _GotoState56Action,
	{_State341, VarOrLetType}:                     _GotoState57Action,
	{_State341, ExpressionType}:                   _GotoState231Action,
	{_State341, OptionalLabelDeclType}:            _GotoState50Action,
	{_State341, SequenceExprType}:                 _GotoState55Action,
	{_State341, CallExprType}:                     _GotoState43Action,
	{_State341, AtomExprType}:                     _GotoState42Action,
	{_State341, LiteralExprType}:                  _GotoState48Action,
	{_State341, ImplicitStructExprType}:           _GotoState46Action,
	{_State341, AccessExprType}:                   _GotoState38Action,
	{_State341, PostfixUnaryExprType}:             _GotoState52Action,
	{_State341, PrefixUnaryOpType}:                _GotoState54Action,
	{_State341, PrefixUnaryExprType}:              _GotoState53Action,
	{_State341, MulExprType}:                      _GotoState49Action,
	{_State341, AddExprType}:                      _GotoState39Action,
	{_State341, CmpExprType}:                      _GotoState44Action,
	{_State341, AndExprType}:                      _GotoState40Action,
	{_State341, OrExprType}:                       _GotoState51Action,
	{_State341, InitializableTypeType}:            _GotoState47Action,
	{_State341, ExplicitStructDefType}:            _GotoState45Action,
	{_State341, AnonymousFuncExprType}:            _GotoState41Action,
	{_State344, AddToken}:                         _GotoState178Action,
	{_State344, SubToken}:                         _GotoState183Action,
	{_State344, MulToken}:                         _GotoState181Action,
	{_State345, IdentifierToken}:                  _GotoState243Action,
	{_State345, GenericParameterDefType}:          _GotoState405Action,
	{_State347, IdentifierToken}:                  _GotoState80Action,
	{_State347, StructToken}:                      _GotoState34Action,
	{_State347, EnumToken}:                        _GotoState77Action,
	{_State347, TraitToken}:                       _GotoState85Action,
	{_State347, FuncToken}:                        _GotoState79Action,
	{_State347, LparenToken}:                      _GotoState81Action,
	{_State347, LbracketToken}:                    _GotoState26Action,
	{_State347, DotToken}:                         _GotoState76Action,
	{_State347, QuestionToken}:                    _GotoState83Action,
	{_State347, ExclaimToken}:                     _GotoState78Action,
	{_State347, TildeTildeToken}:                  _GotoState84Action,
	{_State347, BitNegToken}:                      _GotoState75Action,
	{_State347, BitAndToken}:                      _GotoState74Action,
	{_State347, ParseErrorToken}:                  _GotoState82Action,
	{_State347, InitializableTypeType}:            _GotoState91Action,
	{_State347, AtomTypeType}:                     _GotoState86Action,
	{_State347, ReturnableTypeType}:               _GotoState92Action,
	{_State347, ValueTypeType}:                    _GotoState406Action,
	{_State347, ImplicitStructDefType}:            _GotoState90Action,
	{_State347, ExplicitStructDefType}:            _GotoState45Action,
	{_State347, ImplicitEnumDefType}:              _GotoState89Action,
	{_State347, ExplicitEnumDefType}:              _GotoState87Action,
	{_State347, TraitDefType}:                     _GotoState93Action,
	{_State347, FuncTypeType}:                     _GotoState88Action,
	{_State348, RparenToken}:                      _GotoState407Action,
	{_State349, AddToken}:                         _GotoState178Action,
	{_State349, SubToken}:                         _GotoState183Action,
	{_State349, MulToken}:                         _GotoState181Action,
	{_State350, DollarLbracketToken}:              _GotoState149Action,
	{_State350, OptionalGenericParametersType}:    _GotoState408Action,
	{_State351, LbraceToken}:                      _GotoState59Action,
	{_State351, StatementBlockType}:               _GotoState409Action,
	{_State354, IdentifierToken}:                  _GotoState166Action,
	{_State354, UnsafeToken}:                      _GotoState167Action,
	{_State354, StructToken}:                      _GotoState34Action,
	{_State354, EnumToken}:                        _GotoState77Action,
	{_State354, TraitToken}:                       _GotoState85Action,
	{_State354, FuncToken}:                        _GotoState79Action,
	{_State354, LparenToken}:                      _GotoState81Action,
	{_State354, LbracketToken}:                    _GotoState26Action,
	{_State354, DotToken}:                         _GotoState76Action,
	{_State354, QuestionToken}:                    _GotoState83Action,
	{_State354, ExclaimToken}:                     _GotoState78Action,
	{_State354, TildeTildeToken}:                  _GotoState84Action,
	{_State354, BitNegToken}:                      _GotoState75Action,
	{_State354, BitAndToken}:                      _GotoState74Action,
	{_State354, ParseErrorToken}:                  _GotoState82Action,
	{_State354, UnsafeStatementType}:              _GotoState173Action,
	{_State354, InitializableTypeType}:            _GotoState91Action,
	{_State354, AtomTypeType}:                     _GotoState86Action,
	{_State354, ReturnableTypeType}:               _GotoState92Action,
	{_State354, ValueTypeType}:                    _GotoState174Action,
	{_State354, FieldDefType}:                     _GotoState257Action,
	{_State354, ImplicitStructDefType}:            _GotoState90Action,
	{_State354, ExplicitStructDefType}:            _GotoState45Action,
	{_State354, EnumValueDefType}:                 _GotoState410Action,
	{_State354, ImplicitEnumDefType}:              _GotoState89Action,
	{_State354, ExplicitEnumDefType}:              _GotoState87Action,
	{_State354, TraitDefType}:                     _GotoState93Action,
	{_State354, FuncTypeType}:                     _GotoState88Action,
	{_State355, IdentifierToken}:                  _GotoState166Action,
	{_State355, UnsafeToken}:                      _GotoState167Action,
	{_State355, StructToken}:                      _GotoState34Action,
	{_State355, EnumToken}:                        _GotoState77Action,
	{_State355, TraitToken}:                       _GotoState85Action,
	{_State355, FuncToken}:                        _GotoState79Action,
	{_State355, LparenToken}:                      _GotoState81Action,
	{_State355, LbracketToken}:                    _GotoState26Action,
	{_State355, DotToken}:                         _GotoState76Action,
	{_State355, QuestionToken}:                    _GotoState83Action,
	{_State355, ExclaimToken}:                     _GotoState78Action,
	{_State355, TildeTildeToken}:                  _GotoState84Action,
	{_State355, BitNegToken}:                      _GotoState75Action,
	{_State355, BitAndToken}:                      _GotoState74Action,
	{_State355, ParseErrorToken}:                  _GotoState82Action,
	{_State355, UnsafeStatementType}:              _GotoState173Action,
	{_State355, InitializableTypeType}:            _GotoState91Action,
	{_State355, AtomTypeType}:                     _GotoState86Action,
	{_State355, ReturnableTypeType}:               _GotoState92Action,
	{_State355, ValueTypeType}:                    _GotoState174Action,
	{_State355, FieldDefType}:                     _GotoState257Action,
	{_State355, ImplicitStructDefType}:            _GotoState90Action,
	{_State355, ExplicitStructDefType}:            _GotoState45Action,
	{_State355, EnumValueDefType}:                 _GotoState411Action,
	{_State355, ImplicitEnumDefType}:              _GotoState89Action,
	{_State355, ExplicitEnumDefType}:              _GotoState87Action,
	{_State355, TraitDefType}:                     _GotoState93Action,
	{_State355, FuncTypeType}:                     _GotoState88Action,
	{_State357, IdentifierToken}:                  _GotoState166Action,
	{_State357, UnsafeToken}:                      _GotoState167Action,
	{_State357, StructToken}:                      _GotoState34Action,
	{_State357, EnumToken}:                        _GotoState77Action,
	{_State357, TraitToken}:                       _GotoState85Action,
	{_State357, FuncToken}:                        _GotoState79Action,
	{_State357, LparenToken}:                      _GotoState81Action,
	{_State357, LbracketToken}:                    _GotoState26Action,
	{_State357, DotToken}:                         _GotoState76Action,
	{_State357, QuestionToken}:                    _GotoState83Action,
	{_State357, ExclaimToken}:                     _GotoState78Action,
	{_State357, TildeTildeToken}:                  _GotoState84Action,
	{_State357, BitNegToken}:                      _GotoState75Action,
	{_State357, BitAndToken}:                      _GotoState74Action,
	{_State357, ParseErrorToken}:                  _GotoState82Action,
	{_State357, UnsafeStatementType}:              _GotoState173Action,
	{_State357, InitializableTypeType}:            _GotoState91Action,
	{_State357, AtomTypeType}:                     _GotoState86Action,
	{_State357, ReturnableTypeType}:               _GotoState92Action,
	{_State357, ValueTypeType}:                    _GotoState174Action,
	{_State357, FieldDefType}:                     _GotoState257Action,
	{_State357, ImplicitStructDefType}:            _GotoState90Action,
	{_State357, ExplicitStructDefType}:            _GotoState45Action,
	{_State357, EnumValueDefType}:                 _GotoState412Action,
	{_State357, ImplicitEnumDefType}:              _GotoState89Action,
	{_State357, ExplicitEnumDefType}:              _GotoState87Action,
	{_State357, TraitDefType}:                     _GotoState93Action,
	{_State357, FuncTypeType}:                     _GotoState88Action,
	{_State358, IdentifierToken}:                  _GotoState166Action,
	{_State358, UnsafeToken}:                      _GotoState167Action,
	{_State358, StructToken}:                      _GotoState34Action,
	{_State358, EnumToken}:                        _GotoState77Action,
	{_State358, TraitToken}:                       _GotoState85Action,
	{_State358, FuncToken}:                        _GotoState79Action,
	{_State358, LparenToken}:                      _GotoState81Action,
	{_State358, LbracketToken}:                    _GotoState26Action,
	{_State358, DotToken}:                         _GotoState76Action,
	{_State358, QuestionToken}:                    _GotoState83Action,
	{_State358, ExclaimToken}:                     _GotoState78Action,
	{_State358, TildeTildeToken}:                  _GotoState84Action,
	{_State358, BitNegToken}:                      _GotoState75Action,
	{_State358, BitAndToken}:                      _GotoState74Action,
	{_State358, ParseErrorToken}:                  _GotoState82Action,
	{_State358, UnsafeStatementType}:              _GotoState173Action,
	{_State358, InitializableTypeType}:            _GotoState91Action,
	{_State358, AtomTypeType}:                     _GotoState86Action,
	{_State358, ReturnableTypeType}:               _GotoState92Action,
	{_State358, ValueTypeType}:                    _GotoState174Action,
	{_State358, FieldDefType}:                     _GotoState257Action,
	{_State358, ImplicitStructDefType}:            _GotoState90Action,
	{_State358, ExplicitStructDefType}:            _GotoState45Action,
	{_State358, EnumValueDefType}:                 _GotoState413Action,
	{_State358, ImplicitEnumDefType}:              _GotoState89Action,
	{_State358, ExplicitEnumDefType}:              _GotoState87Action,
	{_State358, TraitDefType}:                     _GotoState93Action,
	{_State358, FuncTypeType}:                     _GotoState88Action,
	{_State359, AddToken}:                         _GotoState178Action,
	{_State359, SubToken}:                         _GotoState183Action,
	{_State359, MulToken}:                         _GotoState181Action,
	{_State360, IdentifierToken}:                  _GotoState80Action,
	{_State360, StructToken}:                      _GotoState34Action,
	{_State360, EnumToken}:                        _GotoState77Action,
	{_State360, TraitToken}:                       _GotoState85Action,
	{_State360, FuncToken}:                        _GotoState79Action,
	{_State360, LparenToken}:                      _GotoState81Action,
	{_State360, LbracketToken}:                    _GotoState26Action,
	{_State360, DotToken}:                         _GotoState76Action,
	{_State360, QuestionToken}:                    _GotoState83Action,
	{_State360, ExclaimToken}:                     _GotoState78Action,
	{_State360, TildeTildeToken}:                  _GotoState84Action,
	{_State360, BitNegToken}:                      _GotoState75Action,
	{_State360, BitAndToken}:                      _GotoState74Action,
	{_State360, ParseErrorToken}:                  _GotoState82Action,
	{_State360, InitializableTypeType}:            _GotoState91Action,
	{_State360, AtomTypeType}:                     _GotoState86Action,
	{_State360, ReturnableTypeType}:               _GotoState92Action,
	{_State360, ValueTypeType}:                    _GotoState414Action,
	{_State360, ImplicitStructDefType}:            _GotoState90Action,
	{_State360, ExplicitStructDefType}:            _GotoState45Action,
	{_State360, ImplicitEnumDefType}:              _GotoState89Action,
	{_State360, ExplicitEnumDefType}:              _GotoState87Action,
	{_State360, TraitDefType}:                     _GotoState93Action,
	{_State360, FuncTypeType}:                     _GotoState88Action,
	{_State361, AddToken}:                         _GotoState178Action,
	{_State361, SubToken}:                         _GotoState183Action,
	{_State361, MulToken}:                         _GotoState181Action,
	{_State362, IdentifierToken}:                  _GotoState80Action,
	{_State362, StructToken}:                      _GotoState34Action,
	{_State362, EnumToken}:                        _GotoState77Action,
	{_State362, TraitToken}:                       _GotoState85Action,
	{_State362, FuncToken}:                        _GotoState79Action,
	{_State362, LparenToken}:                      _GotoState81Action,
	{_State362, LbracketToken}:                    _GotoState26Action,
	{_State362, DotToken}:                         _GotoState76Action,
	{_State362, QuestionToken}:                    _GotoState83Action,
	{_State362, ExclaimToken}:                     _GotoState78Action,
	{_State362, TildeTildeToken}:                  _GotoState84Action,
	{_State362, BitNegToken}:                      _GotoState75Action,
	{_State362, BitAndToken}:                      _GotoState74Action,
	{_State362, ParseErrorToken}:                  _GotoState82Action,
	{_State362, InitializableTypeType}:            _GotoState91Action,
	{_State362, AtomTypeType}:                     _GotoState86Action,
	{_State362, ReturnableTypeType}:               _GotoState352Action,
	{_State362, ImplicitStructDefType}:            _GotoState90Action,
	{_State362, ExplicitStructDefType}:            _GotoState45Action,
	{_State362, ImplicitEnumDefType}:              _GotoState89Action,
	{_State362, ExplicitEnumDefType}:              _GotoState87Action,
	{_State362, TraitDefType}:                     _GotoState93Action,
	{_State362, ReturnTypeType}:                   _GotoState415Action,
	{_State362, FuncTypeType}:                     _GotoState88Action,
	{_State363, IdentifierToken}:                  _GotoState260Action,
	{_State363, StructToken}:                      _GotoState34Action,
	{_State363, EnumToken}:                        _GotoState77Action,
	{_State363, TraitToken}:                       _GotoState85Action,
	{_State363, FuncToken}:                        _GotoState79Action,
	{_State363, LparenToken}:                      _GotoState81Action,
	{_State363, LbracketToken}:                    _GotoState26Action,
	{_State363, DotToken}:                         _GotoState76Action,
	{_State363, QuestionToken}:                    _GotoState83Action,
	{_State363, ExclaimToken}:                     _GotoState78Action,
	{_State363, DotDotDotToken}:                   _GotoState259Action,
	{_State363, TildeTildeToken}:                  _GotoState84Action,
	{_State363, BitNegToken}:                      _GotoState75Action,
	{_State363, BitAndToken}:                      _GotoState74Action,
	{_State363, ParseErrorToken}:                  _GotoState82Action,
	{_State363, InitializableTypeType}:            _GotoState91Action,
	{_State363, AtomTypeType}:                     _GotoState86Action,
	{_State363, ReturnableTypeType}:               _GotoState92Action,
	{_State363, ValueTypeType}:                    _GotoState264Action,
	{_State363, ImplicitStructDefType}:            _GotoState90Action,
	{_State363, ExplicitStructDefType}:            _GotoState45Action,
	{_State363, ImplicitEnumDefType}:              _GotoState89Action,
	{_State363, ExplicitEnumDefType}:              _GotoState87Action,
	{_State363, TraitDefType}:                     _GotoState93Action,
	{_State363, ParameterDeclType}:                _GotoState416Action,
	{_State363, FuncTypeType}:                     _GotoState88Action,
	{_State365, GreaterToken}:                     _GotoState417Action,
	{_State370, LparenToken}:                      _GotoState418Action,
	{_State372, IdentifierToken}:                  _GotoState166Action,
	{_State372, UnsafeToken}:                      _GotoState167Action,
	{_State372, StructToken}:                      _GotoState34Action,
	{_State372, EnumToken}:                        _GotoState77Action,
	{_State372, TraitToken}:                       _GotoState85Action,
	{_State372, FuncToken}:                        _GotoState275Action,
	{_State372, LparenToken}:                      _GotoState81Action,
	{_State372, LbracketToken}:                    _GotoState26Action,
	{_State372, DotToken}:                         _GotoState76Action,
	{_State372, QuestionToken}:                    _GotoState83Action,
	{_State372, ExclaimToken}:                     _GotoState78Action,
	{_State372, TildeTildeToken}:                  _GotoState84Action,
	{_State372, BitNegToken}:                      _GotoState75Action,
	{_State372, BitAndToken}:                      _GotoState74Action,
	{_State372, ParseErrorToken}:                  _GotoState82Action,
	{_State372, UnsafeStatementType}:              _GotoState173Action,
	{_State372, InitializableTypeType}:            _GotoState91Action,
	{_State372, AtomTypeType}:                     _GotoState86Action,
	{_State372, ReturnableTypeType}:               _GotoState92Action,
	{_State372, ValueTypeType}:                    _GotoState174Action,
	{_State372, FieldDefType}:                     _GotoState276Action,
	{_State372, ImplicitStructDefType}:            _GotoState90Action,
	{_State372, ExplicitStructDefType}:            _GotoState45Action,
	{_State372, ImplicitEnumDefType}:              _GotoState89Action,
	{_State372, ExplicitEnumDefType}:              _GotoState87Action,
	{_State372, TraitPropertyType}:                _GotoState419Action,
	{_State372, TraitDefType}:                     _GotoState93Action,
	{_State372, FuncTypeType}:                     _GotoState88Action,
	{_State372, MethodSignatureType}:              _GotoState277Action,
	{_State373, IdentifierToken}:                  _GotoState166Action,
	{_State373, UnsafeToken}:                      _GotoState167Action,
	{_State373, StructToken}:                      _GotoState34Action,
	{_State373, EnumToken}:                        _GotoState77Action,
	{_State373, TraitToken}:                       _GotoState85Action,
	{_State373, FuncToken}:                        _GotoState275Action,
	{_State373, LparenToken}:                      _GotoState81Action,
	{_State373, LbracketToken}:                    _GotoState26Action,
	{_State373, DotToken}:                         _GotoState76Action,
	{_State373, QuestionToken}:                    _GotoState83Action,
	{_State373, ExclaimToken}:                     _GotoState78Action,
	{_State373, TildeTildeToken}:                  _GotoState84Action,
	{_State373, BitNegToken}:                      _GotoState75Action,
	{_State373, BitAndToken}:                      _GotoState74Action,
	{_State373, ParseErrorToken}:                  _GotoState82Action,
	{_State373, UnsafeStatementType}:              _GotoState173Action,
	{_State373, InitializableTypeType}:            _GotoState91Action,
	{_State373, AtomTypeType}:                     _GotoState86Action,
	{_State373, ReturnableTypeType}:               _GotoState92Action,
	{_State373, ValueTypeType}:                    _GotoState174Action,
	{_State373, FieldDefType}:                     _GotoState276Action,
	{_State373, ImplicitStructDefType}:            _GotoState90Action,
	{_State373, ExplicitStructDefType}:            _GotoState45Action,
	{_State373, ImplicitEnumDefType}:              _GotoState89Action,
	{_State373, ExplicitEnumDefType}:              _GotoState87Action,
	{_State373, TraitPropertyType}:                _GotoState420Action,
	{_State373, TraitDefType}:                     _GotoState93Action,
	{_State373, FuncTypeType}:                     _GotoState88Action,
	{_State373, MethodSignatureType}:              _GotoState277Action,
	{_State378, AddToken}:                         _GotoState178Action,
	{_State378, SubToken}:                         _GotoState183Action,
	{_State378, MulToken}:                         _GotoState181Action,
	{_State382, DoToken}:                          _GotoState421Action,
	{_State383, SemicolonToken}:                   _GotoState422Action,
	{_State386, LparenToken}:                      _GotoState28Action,
	{_State386, ImplicitStructExprType}:           _GotoState423Action,
	{_State387, IdentifierToken}:                  _GotoState424Action,
	{_State388, IntegerLiteralToken}:              _GotoState24Action,
	{_State388, FloatLiteralToken}:                _GotoState20Action,
	{_State388, RuneLiteralToken}:                 _GotoState32Action,
	{_State388, StringLiteralToken}:               _GotoState33Action,
	{_State388, IdentifierToken}:                  _GotoState23Action,
	{_State388, TrueToken}:                        _GotoState36Action,
	{_State388, FalseToken}:                       _GotoState19Action,
	{_State388, StructToken}:                      _GotoState34Action,
	{_State388, FuncToken}:                        _GotoState21Action,
	{_State388, VarToken}:                         _GotoState37Action,
	{_State388, LetToken}:                         _GotoState27Action,
	{_State388, NotToken}:                         _GotoState30Action,
	{_State388, LabelDeclToken}:                   _GotoState25Action,
	{_State388, LparenToken}:                      _GotoState28Action,
	{_State388, LbracketToken}:                    _GotoState26Action,
	{_State388, SubToken}:                         _GotoState35Action,
	{_State388, MulToken}:                         _GotoState29Action,
	{_State388, BitNegToken}:                      _GotoState18Action,
	{_State388, BitAndToken}:                      _GotoState17Action,
	{_State388, GreaterToken}:                     _GotoState22Action,
	{_State388, ParseErrorToken}:                  _GotoState31Action,
	{_State388, VarDeclPatternType}:               _GotoState56Action,
	{_State388, VarOrLetType}:                     _GotoState57Action,
	{_State388, OptionalLabelDeclType}:            _GotoState72Action,
	{_State388, SequenceExprType}:                 _GotoState425Action,
	{_State388, CallExprType}:                     _GotoState43Action,
	{_State388, AtomExprType}:                     _GotoState42Action,
	{_State388, LiteralExprType}:                  _GotoState48Action,
	{_State388, ImplicitStructExprType}:           _GotoState46Action,
	{_State388, AccessExprType}:                   _GotoState38Action,
	{_State388, PostfixUnaryExprType}:             _GotoState52Action,
	{_State388, PrefixUnaryOpType}:                _GotoState54Action,
	{_State388, PrefixUnaryExprType}:              _GotoState53Action,
	{_State388, MulExprType}:                      _GotoState49Action,
	{_State388, AddExprType}:                      _GotoState39Action,
	{_State388, CmpExprType}:                      _GotoState44Action,
	{_State388, AndExprType}:                      _GotoState40Action,
	{_State388, OrExprType}:                       _GotoState51Action,
	{_State388, InitializableTypeType}:            _GotoState47Action,
	{_State388, ExplicitStructDefType}:            _GotoState45Action,
	{_State388, AnonymousFuncExprType}:            _GotoState41Action,
	{_State389, IntegerLiteralToken}:              _GotoState24Action,
	{_State389, FloatLiteralToken}:                _GotoState20Action,
	{_State389, RuneLiteralToken}:                 _GotoState32Action,
	{_State389, StringLiteralToken}:               _GotoState33Action,
	{_State389, IdentifierToken}:                  _GotoState23Action,
	{_State389, TrueToken}:                        _GotoState36Action,
	{_State389, FalseToken}:                       _GotoState19Action,
	{_State389, StructToken}:                      _GotoState34Action,
	{_State389, FuncToken}:                        _GotoState21Action,
	{_State389, VarToken}:                         _GotoState305Action,
	{_State389, LetToken}:                         _GotoState27Action,
	{_State389, NotToken}:                         _GotoState30Action,
	{_State389, LabelDeclToken}:                   _GotoState25Action,
	{_State389, LparenToken}:                      _GotoState28Action,
	{_State389, LbracketToken}:                    _GotoState26Action,
	{_State389, DotToken}:                         _GotoState304Action,
	{_State389, SubToken}:                         _GotoState35Action,
	{_State389, MulToken}:                         _GotoState29Action,
	{_State389, BitNegToken}:                      _GotoState18Action,
	{_State389, BitAndToken}:                      _GotoState17Action,
	{_State389, GreaterToken}:                     _GotoState22Action,
	{_State389, ParseErrorToken}:                  _GotoState31Action,
	{_State389, VarDeclPatternType}:               _GotoState56Action,
	{_State389, VarOrLetType}:                     _GotoState57Action,
	{_State389, AssignPatternType}:                _GotoState306Action,
	{_State389, CasePatternType}:                  _GotoState426Action,
	{_State389, OptionalLabelDeclType}:            _GotoState72Action,
	{_State389, SequenceExprType}:                 _GotoState309Action,
	{_State389, CallExprType}:                     _GotoState43Action,
	{_State389, AtomExprType}:                     _GotoState42Action,
	{_State389, LiteralExprType}:                  _GotoState48Action,
	{_State389, ImplicitStructExprType}:           _GotoState46Action,
	{_State389, AccessExprType}:                   _GotoState38Action,
	{_State389, PostfixUnaryExprType}:             _GotoState52Action,
	{_State389, PrefixUnaryOpType}:                _GotoState54Action,
	{_State389, PrefixUnaryExprType}:              _GotoState53Action,
	{_State389, MulExprType}:                      _GotoState49Action,
	{_State389, AddExprType}:                      _GotoState39Action,
	{_State389, CmpExprType}:                      _GotoState44Action,
	{_State389, AndExprType}:                      _GotoState40Action,
	{_State389, OrExprType}:                       _GotoState51Action,
	{_State389, InitializableTypeType}:            _GotoState47Action,
	{_State389, ExplicitStructDefType}:            _GotoState45Action,
	{_State389, AnonymousFuncExprType}:            _GotoState41Action,
	{_State390, IfToken}:                          _GotoState132Action,
	{_State390, LbraceToken}:                      _GotoState59Action,
	{_State390, StatementBlockType}:               _GotoState428Action,
	{_State390, IfExprType}:                       _GotoState427Action,
	{_State393, IntegerLiteralToken}:              _GotoState24Action,
	{_State393, FloatLiteralToken}:                _GotoState20Action,
	{_State393, RuneLiteralToken}:                 _GotoState32Action,
	{_State393, StringLiteralToken}:               _GotoState33Action,
	{_State393, IdentifierToken}:                  _GotoState23Action,
	{_State393, TrueToken}:                        _GotoState36Action,
	{_State393, FalseToken}:                       _GotoState19Action,
	{_State393, ReturnToken}:                      _GotoState228Action,
	{_State393, BreakToken}:                       _GotoState220Action,
	{_State393, ContinueToken}:                    _GotoState222Action,
	{_State393, FallthroughToken}:                 _GotoState225Action,
	{_State393, UnsafeToken}:                      _GotoState167Action,
	{_State393, StructToken}:                      _GotoState34Action,
	{_State393, FuncToken}:                        _GotoState21Action,
	{_State393, AsyncToken}:                       _GotoState219Action,
	{_State393, DeferToken}:                       _GotoState224Action,
	{_State393, VarToken}:                         _GotoState37Action,
	{_State393, LetToken}:                         _GotoState27Action,
	{_State393, NotToken}:                         _GotoState30Action,
	{_State393, LabelDeclToken}:                   _GotoState25Action,
	{_State393, LparenToken}:                      _GotoState28Action,
	{_State393, LbracketToken}:                    _GotoState26Action,
	{_State393, SubToken}:                         _GotoState35Action,
	{_State393, MulToken}:                         _GotoState29Action,
	{_State393, BitNegToken}:                      _GotoState18Action,
	{_State393, BitAndToken}:                      _GotoState17Action,
	{_State393, GreaterToken}:                     _GotoState22Action,
	{_State393, ParseErrorToken}:                  _GotoState31Action,
	{_State393, SimpleStatementBodyType}:          _GotoState395Action,
	{_State393, OptionalSimpleStatementBodyType}:  _GotoState429Action,
	{_State393, UnsafeStatementType}:              _GotoState239Action,
	{_State393, JumpTypeType}:                     _GotoState234Action,
	{_State393, ExpressionsType}:                  _GotoState232Action,
	{_State393, VarDeclPatternType}:               _GotoState56Action,
	{_State393, VarOrLetType}:                     _GotoState57Action,
	{_State393, AssignPatternType}:                _GotoState230Action,
	{_State393, ExpressionType}:                   _GotoState231Action,
	{_State393, OptionalLabelDeclType}:            _GotoState50Action,
	{_State393, SequenceExprType}:                 _GotoState235Action,
	{_State393, CallExprType}:                     _GotoState43Action,
	{_State393, AtomExprType}:                     _GotoState42Action,
	{_State393, LiteralExprType}:                  _GotoState48Action,
	{_State393, ImplicitStructExprType}:           _GotoState46Action,
	{_State393, AccessExprType}:                   _GotoState229Action,
	{_State393, PostfixUnaryExprType}:             _GotoState52Action,
	{_State393, PrefixUnaryOpType}:                _GotoState54Action,
	{_State393, PrefixUnaryExprType}:              _GotoState53Action,
	{_State393, MulExprType}:                      _GotoState49Action,
	{_State393, AddExprType}:                      _GotoState39Action,
	{_State393, CmpExprType}:                      _GotoState44Action,
	{_State393, AndExprType}:                      _GotoState40Action,
	{_State393, OrExprType}:                       _GotoState51Action,
	{_State393, InitializableTypeType}:            _GotoState47Action,
	{_State393, ExplicitStructDefType}:            _GotoState45Action,
	{_State393, AnonymousFuncExprType}:            _GotoState41Action,
	{_State396, NewlinesToken}:                    _GotoState431Action,
	{_State396, CommaToken}:                       _GotoState430Action,
	{_State398, StringLiteralToken}:               _GotoState321Action,
	{_State398, RparenToken}:                      _GotoState432Action,
	{_State398, ImportClauseType}:                 _GotoState396Action,
	{_State398, ImportClauseTerminalType}:         _GotoState433Action,
	{_State399, IdentifierToken}:                  _GotoState434Action,
	{_State403, CommaToken}:                       _GotoState339Action,
	{_State406, AddToken}:                         _GotoState178Action,
	{_State406, SubToken}:                         _GotoState183Action,
	{_State406, MulToken}:                         _GotoState181Action,
	{_State407, IdentifierToken}:                  _GotoState80Action,
	{_State407, StructToken}:                      _GotoState34Action,
	{_State407, EnumToken}:                        _GotoState77Action,
	{_State407, TraitToken}:                       _GotoState85Action,
	{_State407, FuncToken}:                        _GotoState79Action,
	{_State407, LparenToken}:                      _GotoState81Action,
	{_State407, LbracketToken}:                    _GotoState26Action,
	{_State407, DotToken}:                         _GotoState76Action,
	{_State407, QuestionToken}:                    _GotoState83Action,
	{_State407, ExclaimToken}:                     _GotoState78Action,
	{_State407, TildeTildeToken}:                  _GotoState84Action,
	{_State407, BitNegToken}:                      _GotoState75Action,
	{_State407, BitAndToken}:                      _GotoState74Action,
	{_State407, ParseErrorToken}:                  _GotoState82Action,
	{_State407, InitializableTypeType}:            _GotoState91Action,
	{_State407, AtomTypeType}:                     _GotoState86Action,
	{_State407, ReturnableTypeType}:               _GotoState352Action,
	{_State407, ImplicitStructDefType}:            _GotoState90Action,
	{_State407, ExplicitStructDefType}:            _GotoState45Action,
	{_State407, ImplicitEnumDefType}:              _GotoState89Action,
	{_State407, ExplicitEnumDefType}:              _GotoState87Action,
	{_State407, TraitDefType}:                     _GotoState93Action,
	{_State407, ReturnTypeType}:                   _GotoState435Action,
	{_State407, FuncTypeType}:                     _GotoState88Action,
	{_State408, LparenToken}:                      _GotoState436Action,
	{_State414, AddToken}:                         _GotoState178Action,
	{_State414, SubToken}:                         _GotoState183Action,
	{_State414, MulToken}:                         _GotoState181Action,
	{_State417, StringLiteralToken}:               _GotoState437Action,
	{_State418, IdentifierToken}:                  _GotoState260Action,
	{_State418, StructToken}:                      _GotoState34Action,
	{_State418, EnumToken}:                        _GotoState77Action,
	{_State418, TraitToken}:                       _GotoState85Action,
	{_State418, FuncToken}:                        _GotoState79Action,
	{_State418, LparenToken}:                      _GotoState81Action,
	{_State418, LbracketToken}:                    _GotoState26Action,
	{_State418, DotToken}:                         _GotoState76Action,
	{_State418, QuestionToken}:                    _GotoState83Action,
	{_State418, ExclaimToken}:                     _GotoState78Action,
	{_State418, DotDotDotToken}:                   _GotoState259Action,
	{_State418, TildeTildeToken}:                  _GotoState84Action,
	{_State418, BitNegToken}:                      _GotoState75Action,
	{_State418, BitAndToken}:                      _GotoState74Action,
	{_State418, ParseErrorToken}:                  _GotoState82Action,
	{_State418, InitializableTypeType}:            _GotoState91Action,
	{_State418, AtomTypeType}:                     _GotoState86Action,
	{_State418, ReturnableTypeType}:               _GotoState92Action,
	{_State418, ValueTypeType}:                    _GotoState264Action,
	{_State418, ImplicitStructDefType}:            _GotoState90Action,
	{_State418, ExplicitStructDefType}:            _GotoState45Action,
	{_State418, ImplicitEnumDefType}:              _GotoState89Action,
	{_State418, ExplicitEnumDefType}:              _GotoState87Action,
	{_State418, TraitDefType}:                     _GotoState93Action,
	{_State418, ParameterDeclType}:                _GotoState262Action,
	{_State418, ParameterDeclsType}:               _GotoState263Action,
	{_State418, OptionalParameterDeclsType}:       _GotoState438Action,
	{_State418, FuncTypeType}:                     _GotoState88Action,
	{_State421, LbraceToken}:                      _GotoState59Action,
	{_State421, StatementBlockType}:               _GotoState439Action,
	{_State422, IntegerLiteralToken}:              _GotoState24Action,
	{_State422, FloatLiteralToken}:                _GotoState20Action,
	{_State422, RuneLiteralToken}:                 _GotoState32Action,
	{_State422, StringLiteralToken}:               _GotoState33Action,
	{_State422, IdentifierToken}:                  _GotoState23Action,
	{_State422, TrueToken}:                        _GotoState36Action,
	{_State422, FalseToken}:                       _GotoState19Action,
	{_State422, StructToken}:                      _GotoState34Action,
	{_State422, FuncToken}:                        _GotoState21Action,
	{_State422, VarToken}:                         _GotoState37Action,
	{_State422, LetToken}:                         _GotoState27Action,
	{_State422, NotToken}:                         _GotoState30Action,
	{_State422, LabelDeclToken}:                   _GotoState25Action,
	{_State422, LparenToken}:                      _GotoState28Action,
	{_State422, LbracketToken}:                    _GotoState26Action,
	{_State422, SubToken}:                         _GotoState35Action,
	{_State422, MulToken}:                         _GotoState29Action,
	{_State422, BitNegToken}:                      _GotoState18Action,
	{_State422, BitAndToken}:                      _GotoState17Action,
	{_State422, GreaterToken}:                     _GotoState22Action,
	{_State422, ParseErrorToken}:                  _GotoState31Action,
	{_State422, VarDeclPatternType}:               _GotoState56Action,
	{_State422, VarOrLetType}:                     _GotoState57Action,
	{_State422, OptionalLabelDeclType}:            _GotoState72Action,
	{_State422, SequenceExprType}:                 _GotoState384Action,
	{_State422, OptionalSequenceExprType}:         _GotoState440Action,
	{_State422, CallExprType}:                     _GotoState43Action,
	{_State422, AtomExprType}:                     _GotoState42Action,
	{_State422, LiteralExprType}:                  _GotoState48Action,
	{_State422, ImplicitStructExprType}:           _GotoState46Action,
	{_State422, AccessExprType}:                   _GotoState38Action,
	{_State422, PostfixUnaryExprType}:             _GotoState52Action,
	{_State422, PrefixUnaryOpType}:                _GotoState54Action,
	{_State422, PrefixUnaryExprType}:              _GotoState53Action,
	{_State422, MulExprType}:                      _GotoState49Action,
	{_State422, AddExprType}:                      _GotoState39Action,
	{_State422, CmpExprType}:                      _GotoState44Action,
	{_State422, AndExprType}:                      _GotoState40Action,
	{_State422, OrExprType}:                       _GotoState51Action,
	{_State422, InitializableTypeType}:            _GotoState47Action,
	{_State422, ExplicitStructDefType}:            _GotoState45Action,
	{_State422, AnonymousFuncExprType}:            _GotoState41Action,
	{_State424, LparenToken}:                      _GotoState141Action,
	{_State424, TuplePatternType}:                 _GotoState441Action,
	{_State435, LbraceToken}:                      _GotoState59Action,
	{_State435, StatementBlockType}:               _GotoState442Action,
	{_State436, IdentifierToken}:                  _GotoState153Action,
	{_State436, ParameterDefType}:                 _GotoState156Action,
	{_State436, ParameterDefsType}:                _GotoState157Action,
	{_State436, OptionalParameterDefsType}:        _GotoState443Action,
	{_State438, RparenToken}:                      _GotoState444Action,
	{_State440, DoToken}:                          _GotoState445Action,
	{_State443, RparenToken}:                      _GotoState446Action,
	{_State444, IdentifierToken}:                  _GotoState80Action,
	{_State444, StructToken}:                      _GotoState34Action,
	{_State444, EnumToken}:                        _GotoState77Action,
	{_State444, TraitToken}:                       _GotoState85Action,
	{_State444, FuncToken}:                        _GotoState79Action,
	{_State444, LparenToken}:                      _GotoState81Action,
	{_State444, LbracketToken}:                    _GotoState26Action,
	{_State444, DotToken}:                         _GotoState76Action,
	{_State444, QuestionToken}:                    _GotoState83Action,
	{_State444, ExclaimToken}:                     _GotoState78Action,
	{_State444, TildeTildeToken}:                  _GotoState84Action,
	{_State444, BitNegToken}:                      _GotoState75Action,
	{_State444, BitAndToken}:                      _GotoState74Action,
	{_State444, ParseErrorToken}:                  _GotoState82Action,
	{_State444, InitializableTypeType}:            _GotoState91Action,
	{_State444, AtomTypeType}:                     _GotoState86Action,
	{_State444, ReturnableTypeType}:               _GotoState352Action,
	{_State444, ImplicitStructDefType}:            _GotoState90Action,
	{_State444, ExplicitStructDefType}:            _GotoState45Action,
	{_State444, ImplicitEnumDefType}:              _GotoState89Action,
	{_State444, ExplicitEnumDefType}:              _GotoState87Action,
	{_State444, TraitDefType}:                     _GotoState93Action,
	{_State444, ReturnTypeType}:                   _GotoState447Action,
	{_State444, FuncTypeType}:                     _GotoState88Action,
	{_State445, LbraceToken}:                      _GotoState59Action,
	{_State445, StatementBlockType}:               _GotoState448Action,
	{_State446, IdentifierToken}:                  _GotoState80Action,
	{_State446, StructToken}:                      _GotoState34Action,
	{_State446, EnumToken}:                        _GotoState77Action,
	{_State446, TraitToken}:                       _GotoState85Action,
	{_State446, FuncToken}:                        _GotoState79Action,
	{_State446, LparenToken}:                      _GotoState81Action,
	{_State446, LbracketToken}:                    _GotoState26Action,
	{_State446, DotToken}:                         _GotoState76Action,
	{_State446, QuestionToken}:                    _GotoState83Action,
	{_State446, ExclaimToken}:                     _GotoState78Action,
	{_State446, TildeTildeToken}:                  _GotoState84Action,
	{_State446, BitNegToken}:                      _GotoState75Action,
	{_State446, BitAndToken}:                      _GotoState74Action,
	{_State446, ParseErrorToken}:                  _GotoState82Action,
	{_State446, InitializableTypeType}:            _GotoState91Action,
	{_State446, AtomTypeType}:                     _GotoState86Action,
	{_State446, ReturnableTypeType}:               _GotoState352Action,
	{_State446, ImplicitStructDefType}:            _GotoState90Action,
	{_State446, ExplicitStructDefType}:            _GotoState45Action,
	{_State446, ImplicitEnumDefType}:              _GotoState89Action,
	{_State446, ExplicitEnumDefType}:              _GotoState87Action,
	{_State446, TraitDefType}:                     _GotoState93Action,
	{_State446, ReturnTypeType}:                   _GotoState449Action,
	{_State446, FuncTypeType}:                     _GotoState88Action,
	{_State449, LbraceToken}:                      _GotoState59Action,
	{_State449, StatementBlockType}:               _GotoState450Action,
	{_State1, _EndMarker}:                         _ReduceNilToOptionalDefinitionsAction,
	{_State1, _WildcardMarker}:                    _ReduceNilToOptionalNewlinesAction,
	{_State5, _WildcardMarker}:                    _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State11, _EndMarker}:                        _ReduceNewlinesToOptionalDefinitionsAction,
	{_State11, _WildcardMarker}:                   _ReduceNewlinesToOptionalNewlinesAction,
	{_State12, _EndMarker}:                        _ReduceToSourceAction,
	{_State14, _WildcardMarker}:                   _ReduceNoSpecToPackageDefAction,
	{_State17, _WildcardMarker}:                   _ReduceBitAndToPrefixUnaryOpAction,
	{_State18, _WildcardMarker}:                   _ReduceBitNegToPrefixUnaryOpAction,
	{_State19, _WildcardMarker}:                   _ReduceFalseToLiteralExprAction,
	{_State20, _WildcardMarker}:                   _ReduceFloatLiteralToLiteralExprAction,
	{_State22, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State23, _WildcardMarker}:                   _ReduceIdentifierToAtomExprAction,
	{_State24, _WildcardMarker}:                   _ReduceIntegerLiteralToLiteralExprAction,
	{_State25, _WildcardMarker}:                   _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State27, _WildcardMarker}:                   _ReduceLetToVarOrLetAction,
	{_State28, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State28, RparenToken}:                       _ReduceNilToOptionalArgumentsAction,
	{_State28, ColonToken}:                        _ReduceNilToOptionalExpressionAction,
	{_State29, _WildcardMarker}:                   _ReduceMulToPrefixUnaryOpAction,
	{_State30, _WildcardMarker}:                   _ReduceNotToPrefixUnaryOpAction,
	{_State31, _WildcardMarker}:                   _ReduceParseErrorToAtomExprAction,
	{_State32, _WildcardMarker}:                   _ReduceRuneLiteralToLiteralExprAction,
	{_State33, _WildcardMarker}:                   _ReduceStringLiteralToLiteralExprAction,
	{_State35, _WildcardMarker}:                   _ReduceSubToPrefixUnaryOpAction,
	{_State36, _WildcardMarker}:                   _ReduceTrueToLiteralExprAction,
	{_State37, _WildcardMarker}:                   _ReduceVarToVarOrLetAction,
	{_State38, _WildcardMarker}:                   _ReduceAccessExprToPostfixUnaryExprAction,
	{_State38, LparenToken}:                       _ReduceNilToOptionalGenericBindingAction,
	{_State39, _WildcardMarker}:                   _ReduceAddExprToCmpExprAction,
	{_State40, _WildcardMarker}:                   _ReduceAndExprToOrExprAction,
	{_State41, _WildcardMarker}:                   _ReduceAnonymousFuncExprToAtomExprAction,
	{_State42, _WildcardMarker}:                   _ReduceAtomExprToAccessExprAction,
	{_State43, _WildcardMarker}:                   _ReduceCallExprToAccessExprAction,
	{_State44, _WildcardMarker}:                   _ReduceCmpExprToAndExprAction,
	{_State45, _WildcardMarker}:                   _ReduceExplicitStructDefToInitializableTypeAction,
	{_State46, _WildcardMarker}:                   _ReduceImplicitStructExprToAtomExprAction,
	{_State48, _WildcardMarker}:                   _ReduceLiteralExprToAtomExprAction,
	{_State49, _WildcardMarker}:                   _ReduceMulExprToAddExprAction,
	{_State51, _WildcardMarker}:                   _ReduceOrExprToSequenceExprAction,
	{_State52, _WildcardMarker}:                   _ReducePostfixUnaryExprToPrefixUnaryExprAction,
	{_State53, _WildcardMarker}:                   _ReducePrefixUnaryExprToMulExprAction,
	{_State54, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State55, _EndMarker}:                        _ReduceSequenceExprToExpressionAction,
	{_State56, _EndMarker}:                        _ReduceVarDeclPatternToSequenceExprAction,
	{_State58, _WildcardMarker}:                   _ReduceCommentGroupsToDefinitionAction,
	{_State59, _WildcardMarker}:                   _ReduceEmptyListToStatementsAction,
	{_State60, _WildcardMarker}:                   _ReduceDefinitionToDefinitionsAction,
	{_State61, _EndMarker}:                        _ReduceNilToOptionalNewlinesAction,
	{_State62, _WildcardMarker}:                   _ReduceNamedFuncDefToDefinitionAction,
	{_State63, _WildcardMarker}:                   _ReducePackageDefToDefinitionAction,
	{_State64, _WildcardMarker}:                   _ReduceStatementBlockToDefinitionAction,
	{_State65, _WildcardMarker}:                   _ReduceTypeDefToDefinitionAction,
	{_State66, _WildcardMarker}:                   _ReduceGlobalVarDefToDefinitionAction,
	{_State67, _EndMarker}:                        _ReduceWithSpecToPackageDefAction,
	{_State68, _WildcardMarker}:                   _ReduceNilToOptionalGenericParametersAction,
	{_State69, LparenToken}:                       _ReduceNilToOptionalGenericParametersAction,
	{_State71, RparenToken}:                       _ReduceNilToOptionalParameterDefsAction,
	{_State73, _EndMarker}:                        _ReduceAssignVarPatternToSequenceExprAction,
	{_State76, _WildcardMarker}:                   _ReduceNilToOptionalGenericBindingAction,
	{_State80, _WildcardMarker}:                   _ReduceNilToOptionalGenericBindingAction,
	{_State81, RparenToken}:                       _ReduceNilToOptionalImplicitFieldDefsAction,
	{_State82, _WildcardMarker}:                   _ReduceParseErrorToAtomTypeAction,
	{_State86, _WildcardMarker}:                   _ReduceAtomTypeToReturnableTypeAction,
	{_State87, _WildcardMarker}:                   _ReduceExplicitEnumDefToAtomTypeAction,
	{_State88, _WildcardMarker}:                   _ReduceFuncTypeToAtomTypeAction,
	{_State89, _WildcardMarker}:                   _ReduceImplicitEnumDefToAtomTypeAction,
	{_State90, _WildcardMarker}:                   _ReduceImplicitStructDefToAtomTypeAction,
	{_State91, _WildcardMarker}:                   _ReduceInitializableTypeToAtomTypeAction,
	{_State92, _WildcardMarker}:                   _ReduceReturnableTypeToValueTypeAction,
	{_State93, _WildcardMarker}:                   _ReduceTraitDefToAtomTypeAction,
	{_State95, _WildcardMarker}:                   _ReduceDotDotDotToArgumentAction,
	{_State96, _WildcardMarker}:                   _ReduceIdentifierToAtomExprAction,
	{_State97, _WildcardMarker}:                   _ReduceArgumentToArgumentsAction,
	{_State98, RparenToken}:                       _ReduceArgumentsToOptionalArgumentsAction,
	{_State99, _WildcardMarker}:                   _ReduceColonExpressionsToArgumentAction,
	{_State100, _WildcardMarker}:                  _ReducePositionalToArgumentAction,
	{_State100, ColonToken}:                       _ReduceExpressionToOptionalExpressionAction,
	{_State103, RparenToken}:                      _ReduceNilToOptionalExplicitFieldDefsAction,
	{_State104, RbracketToken}:                    _ReduceNilToOptionalGenericArgumentsAction,
	{_State106, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State106, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State107, _WildcardMarker}:                  _ReduceToPostfixUnaryExprAction,
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
	{_State122, RparenToken}:                      _ReduceNilToOptionalArgumentsAction,
	{_State122, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State123, _WildcardMarker}:                  _ReduceBitAndToMulOpAction,
	{_State124, _WildcardMarker}:                  _ReduceBitLshiftToMulOpAction,
	{_State125, _WildcardMarker}:                  _ReduceBitRshiftToMulOpAction,
	{_State126, _WildcardMarker}:                  _ReduceDivToMulOpAction,
	{_State127, _WildcardMarker}:                  _ReduceModToMulOpAction,
	{_State128, _WildcardMarker}:                  _ReduceMulToMulOpAction,
	{_State129, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State131, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State132, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State133, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State134, _EndMarker}:                       _ReduceIfExprToExpressionAction,
	{_State135, _EndMarker}:                       _ReduceLoopExprToExpressionAction,
	{_State136, _WildcardMarker}:                  _ReduceBlockExprToAtomExprAction,
	{_State137, _EndMarker}:                       _ReduceSwitchExprToExpressionAction,
	{_State138, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State139, _WildcardMarker}:                  _ReduceToPrefixUnaryExprAction,
	{_State140, _WildcardMarker}:                  _ReduceIdentifierToVarPatternAction,
	{_State142, _WildcardMarker}:                  _ReduceTuplePatternToVarPatternAction,
	{_State143, _WildcardMarker}:                  _ReduceNilToOptionalValueTypeAction,
	{_State144, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State145, _EndMarker}:                       _ReduceNewlinesToOptionalNewlinesAction,
	{_State146, _EndMarker}:                       _ReduceDefinitionsToOptionalDefinitionsAction,
	{_State147, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State149, RbracketToken}:                    _ReduceNilToOptionalGenericParameterDefsAction,
	{_State151, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State153, _WildcardMarker}:                  _ReduceInferredRefArgToParameterDefAction,
	{_State156, _WildcardMarker}:                  _ReduceParameterDefToParameterDefsAction,
	{_State157, RparenToken}:                      _ReduceParameterDefsToOptionalParameterDefsAction,
	{_State158, _WildcardMarker}:                  _ReduceReferenceToReturnableTypeAction,
	{_State159, _WildcardMarker}:                  _ReducePublicMethodsTraitToReturnableTypeAction,
	{_State160, _WildcardMarker}:                  _ReduceInferredToAtomTypeAction,
	{_State162, _WildcardMarker}:                  _ReduceResultToReturnableTypeAction,
	{_State163, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State165, _WildcardMarker}:                  _ReduceNamedToAtomTypeAction,
	{_State166, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State169, _WildcardMarker}:                  _ReduceFieldDefToImplicitFieldDefsAction,
	{_State169, OrToken}:                          _ReduceFieldDefToEnumValueDefAction,
	{_State171, RparenToken}:                      _ReduceImplicitFieldDefsToOptionalImplicitFieldDefsAction,
	{_State173, _WildcardMarker}:                  _ReduceUnsafeStatementToFieldDefAction,
	{_State174, _WildcardMarker}:                  _ReduceImplicitToFieldDefAction,
	{_State175, _WildcardMarker}:                  _ReduceOptionalToReturnableTypeAction,
	{_State176, _WildcardMarker}:                  _ReducePublicTraitToReturnableTypeAction,
	{_State177, RparenToken}:                      _ReduceNilToOptionalTraitPropertiesAction,
	{_State182, _WildcardMarker}:                  _ReduceSliceToInitializableTypeAction,
	{_State184, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State185, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State185, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State186, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State186, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State186, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State186, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State186, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State187, _WildcardMarker}:                  _ReduceToImplicitStructExprAction,
	{_State188, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State188, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State188, CommaToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State188, RbracketToken}:                    _ReduceNilToOptionalExpressionAction,
	{_State188, RparenToken}:                      _ReduceNilToOptionalExpressionAction,
	{_State189, RparenToken}:                      _ReduceExplicitFieldDefsToOptionalExplicitFieldDefsAction,
	{_State190, _WildcardMarker}:                  _ReduceFieldDefToExplicitFieldDefsAction,
	{_State192, RbracketToken}:                    _ReduceGenericArgumentsToOptionalGenericArgumentsAction,
	{_State194, _WildcardMarker}:                  _ReduceValueTypeToGenericArgumentsAction,
	{_State195, _WildcardMarker}:                  _ReduceAccessToAccessExprAction,
	{_State197, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State197, RparenToken}:                      _ReduceNilToOptionalArgumentsAction,
	{_State197, ColonToken}:                       _ReduceNilToOptionalExpressionAction,
	{_State198, _WildcardMarker}:                  _ReduceToAddExprAction,
	{_State199, _WildcardMarker}:                  _ReduceToAndExprAction,
	{_State200, _WildcardMarker}:                  _ReduceToCmpExprAction,
	{_State202, _WildcardMarker}:                  _ReduceToMulExprAction,
	{_State203, _WildcardMarker}:                  _ReduceInfiniteToLoopExprAction,
	{_State206, _WildcardMarker}:                  _ReduceToAssignPatternAction,
	{_State206, SemicolonToken}:                   _ReduceSequenceExprToForAssignmentAction,
	{_State207, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State209, LbraceToken}:                      _ReduceSequenceExprToConditionAction,
	{_State211, _WildcardMarker}:                  _ReduceToOrExprAction,
	{_State212, _WildcardMarker}:                  _ReduceDotDotDotToFieldVarPatternAction,
	{_State213, _WildcardMarker}:                  _ReduceIdentifierToVarPatternAction,
	{_State214, _WildcardMarker}:                  _ReduceFieldVarPatternToFieldVarPatternsAction,
	{_State216, _WildcardMarker}:                  _ReducePositionalToFieldVarPatternAction,
	{_State217, _EndMarker}:                       _ReduceToVarDeclPatternAction,
	{_State218, _WildcardMarker}:                  _ReduceValueTypeToOptionalValueTypeAction,
	{_State219, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State220, _WildcardMarker}:                  _ReduceBreakToJumpTypeAction,
	{_State221, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State222, _WildcardMarker}:                  _ReduceContinueToJumpTypeAction,
	{_State224, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State225, _WildcardMarker}:                  _ReduceFallthroughToSimpleStatementBodyAction,
	{_State227, _EndMarker}:                       _ReduceToStatementBlockAction,
	{_State228, _WildcardMarker}:                  _ReduceReturnToJumpTypeAction,
	{_State229, _WildcardMarker}:                  _ReduceAccessExprToPostfixUnaryExprAction,
	{_State229, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State231, _WildcardMarker}:                  _ReduceExpressionToExpressionsAction,
	{_State232, _WildcardMarker}:                  _ReduceExpressionOrImplicitStructToSimpleStatementBodyAction,
	{_State233, _WildcardMarker}:                  _ReduceImportStatementToStatementBodyAction,
	{_State234, _WildcardMarker}:                  _ReduceUnlabelledToOptionalJumpLabelAction,
	{_State235, AssignToken}:                      _ReduceToAssignPatternAction,
	{_State235, _WildcardMarker}:                  _ReduceSequenceExprToExpressionAction,
	{_State236, _WildcardMarker}:                  _ReduceSimpleStatementBodyToStatementBodyAction,
	{_State237, _WildcardMarker}:                  _ReduceAddToStatementsAction,
	{_State239, _WildcardMarker}:                  _ReduceUnsafeStatementToSimpleStatementBodyAction,
	{_State240, _WildcardMarker}:                  _ReduceAddToDefinitionsAction,
	{_State241, _WildcardMarker}:                  _ReduceGlobalVarAssignmentToDefinitionAction,
	{_State242, _EndMarker}:                       _ReduceAliasToTypeDefAction,
	{_State243, _WildcardMarker}:                  _ReduceUnconstrainedToGenericParameterDefAction,
	{_State244, _WildcardMarker}:                  _ReduceGenericParameterDefToGenericParameterDefsAction,
	{_State245, RbracketToken}:                    _ReduceGenericParameterDefsToOptionalGenericParameterDefsAction,
	{_State247, _WildcardMarker}:                  _ReduceDefinitionToTypeDefAction,
	{_State248, _EndMarker}:                       _ReduceAliasToNamedFuncDefAction,
	{_State249, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State250, _WildcardMarker}:                  _ReduceInferredRefVarargToParameterDefAction,
	{_State251, _WildcardMarker}:                  _ReduceArgToParameterDefAction,
	{_State253, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State257, _WildcardMarker}:                  _ReduceFieldDefToEnumValueDefAction,
	{_State260, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State262, _WildcardMarker}:                  _ReduceParameterDeclToParameterDeclsAction,
	{_State263, RparenToken}:                      _ReduceParameterDeclsToOptionalParameterDeclsAction,
	{_State264, _WildcardMarker}:                  _ReduceUnamedToParameterDeclAction,
	{_State265, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State266, _WildcardMarker}:                  _ReduceNilToOptionalGenericBindingAction,
	{_State267, _WildcardMarker}:                  _ReduceExplicitToFieldDefAction,
	{_State272, _WildcardMarker}:                  _ReduceToImplicitEnumDefAction,
	{_State274, _WildcardMarker}:                  _ReduceToImplicitStructDefAction,
	{_State276, _WildcardMarker}:                  _ReduceFieldDefToTraitPropertyAction,
	{_State277, _WildcardMarker}:                  _ReduceMethodSignatureToTraitPropertyAction,
	{_State279, RparenToken}:                      _ReduceTraitPropertiesToOptionalTraitPropertiesAction,
	{_State280, _WildcardMarker}:                  _ReduceTraitPropertyToTraitPropertiesAction,
	{_State281, _WildcardMarker}:                  _ReduceTraitUnionToValueTypeAction,
	{_State284, _WildcardMarker}:                  _ReduceTraitIntersectToValueTypeAction,
	{_State285, _WildcardMarker}:                  _ReduceTraitDifferenceToValueTypeAction,
	{_State286, _WildcardMarker}:                  _ReduceNamedToArgumentAction,
	{_State287, _WildcardMarker}:                  _ReduceAddToArgumentsAction,
	{_State288, _WildcardMarker}:                  _ReduceExpressionToOptionalExpressionAction,
	{_State289, _WildcardMarker}:                  _ReduceAddToColonExpressionsAction,
	{_State290, _WildcardMarker}:                  _ReducePairToColonExpressionsAction,
	{_State293, _WildcardMarker}:                  _ReduceToExplicitStructDefAction,
	{_State295, _WildcardMarker}:                  _ReduceBindingToOptionalGenericBindingAction,
	{_State296, _WildcardMarker}:                  _ReduceIndexToAccessExprAction,
	{_State298, _WildcardMarker}:                  _ReduceInitializeExprToAtomExprAction,
	{_State299, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State300, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State301, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State302, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State302, SemicolonToken}:                   _ReduceNilToOptionalSequenceExprAction,
	{_State305, _WildcardMarker}:                  _ReduceVarToVarOrLetAction,
	{_State306, _WildcardMarker}:                  _ReduceAssignPatternToCasePatternAction,
	{_State307, _WildcardMarker}:                  _ReduceCasePatternToCasePatternsAction,
	{_State309, _WildcardMarker}:                  _ReduceToAssignPatternAction,
	{_State310, _WildcardMarker}:                  _ReduceNoElseToIfExprAction,
	{_State311, _EndMarker}:                       _ReduceToSwitchExprAction,
	{_State314, _WildcardMarker}:                  _ReduceToTuplePatternAction,
	{_State315, LparenToken}:                      _ReduceNilToOptionalGenericBindingAction,
	{_State316, NewlinesToken}:                    _ReduceAsyncToSimpleStatementBodyAction,
	{_State316, SemicolonToken}:                   _ReduceAsyncToSimpleStatementBodyAction,
	{_State316, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State318, NewlinesToken}:                    _ReduceNilToOptionalSimpleStatementBodyAction,
	{_State318, SemicolonToken}:                   _ReduceNilToOptionalSimpleStatementBodyAction,
	{_State318, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State319, NewlinesToken}:                    _ReduceDeferToSimpleStatementBodyAction,
	{_State319, SemicolonToken}:                   _ReduceDeferToSimpleStatementBodyAction,
	{_State319, _WildcardMarker}:                  _ReduceCallExprToAccessExprAction,
	{_State321, _WildcardMarker}:                  _ReduceStringLiteralToImportClauseAction,
	{_State322, _WildcardMarker}:                  _ReduceSingleToImportStatementAction,
	{_State323, _WildcardMarker}:                  _ReduceAddAssignToBinaryOpAssignAction,
	{_State324, _WildcardMarker}:                  _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State325, _WildcardMarker}:                  _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State326, _WildcardMarker}:                  _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State327, _WildcardMarker}:                  _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State328, _WildcardMarker}:                  _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State329, _WildcardMarker}:                  _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State330, _WildcardMarker}:                  _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State331, _WildcardMarker}:                  _ReduceDivAssignToBinaryOpAssignAction,
	{_State332, _WildcardMarker}:                  _ReduceModAssignToBinaryOpAssignAction,
	{_State333, _WildcardMarker}:                  _ReduceMulAssignToBinaryOpAssignAction,
	{_State334, _WildcardMarker}:                  _ReduceSubAssignToBinaryOpAssignAction,
	{_State335, _WildcardMarker}:                  _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State336, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State337, _WildcardMarker}:                  _ReduceUnaryOpAssignStatementToSimpleStatementBodyAction,
	{_State338, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State339, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State340, _WildcardMarker}:                  _ReduceJumpLabelToOptionalJumpLabelAction,
	{_State341, NewlinesToken}:                    _ReduceNilToOptionalExpressionsAction,
	{_State341, SemicolonToken}:                   _ReduceNilToOptionalExpressionsAction,
	{_State341, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State342, _WildcardMarker}:                  _ReduceImplicitToStatementAction,
	{_State343, _WildcardMarker}:                  _ReduceExplicitToStatementAction,
	{_State344, _WildcardMarker}:                  _ReduceConstrainedToGenericParameterDefAction,
	{_State346, _WildcardMarker}:                  _ReduceGenericToOptionalGenericParametersAction,
	{_State349, _WildcardMarker}:                  _ReduceVarargToParameterDefAction,
	{_State350, LparenToken}:                      _ReduceNilToOptionalGenericParametersAction,
	{_State352, _WildcardMarker}:                  _ReduceReturnableTypeToReturnTypeAction,
	{_State353, _WildcardMarker}:                  _ReduceAddToParameterDefsAction,
	{_State356, _WildcardMarker}:                  _ReduceToExplicitEnumDefAction,
	{_State359, _WildcardMarker}:                  _ReduceUnnamedVarargToParameterDeclAction,
	{_State361, _WildcardMarker}:                  _ReduceArgToParameterDeclAction,
	{_State362, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State364, _WildcardMarker}:                  _ReduceExternNamedToAtomTypeAction,
	{_State366, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State367, _WildcardMarker}:                  _ReduceDefaultToEnumValueDefAction,
	{_State368, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State369, _WildcardMarker}:                  _ReduceAddToImplicitFieldDefsAction,
	{_State371, _WildcardMarker}:                  _ReduceToTraitDefAction,
	{_State374, _WildcardMarker}:                  _ReduceMapToInitializableTypeAction,
	{_State375, _WildcardMarker}:                  _ReduceArrayToInitializableTypeAction,
	{_State376, _WildcardMarker}:                  _ReduceExplicitToExplicitFieldDefsAction,
	{_State377, _WildcardMarker}:                  _ReduceImplicitToExplicitFieldDefsAction,
	{_State378, _WildcardMarker}:                  _ReduceAddToGenericArgumentsAction,
	{_State379, _WildcardMarker}:                  _ReduceToCallExprAction,
	{_State380, _EndMarker}:                       _ReduceDoWhileToLoopExprAction,
	{_State381, SemicolonToken}:                   _ReduceAssignToForAssignmentAction,
	{_State384, DoToken}:                          _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State385, _EndMarker}:                       _ReduceWhileToLoopExprAction,
	{_State388, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State389, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State391, _WildcardMarker}:                  _ReduceNamedToFieldVarPatternAction,
	{_State392, _WildcardMarker}:                  _ReduceAddToFieldVarPatternsAction,
	{_State393, NewlinesToken}:                    _ReduceNilToOptionalSimpleStatementBodyAction,
	{_State393, SemicolonToken}:                   _ReduceNilToOptionalSimpleStatementBodyAction,
	{_State393, _WildcardMarker}:                  _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State394, _WildcardMarker}:                  _ReduceDefaultBranchStatementToStatementBodyAction,
	{_State395, _WildcardMarker}:                  _ReduceSimpleStatementBodyToOptionalSimpleStatementBodyAction,
	{_State397, _WildcardMarker}:                  _ReduceFirstToImportClausesAction,
	{_State400, _WildcardMarker}:                  _ReduceBinaryOpAssignStatementToSimpleStatementBodyAction,
	{_State401, _WildcardMarker}:                  _ReduceAssignStatementToSimpleStatementBodyAction,
	{_State402, _WildcardMarker}:                  _ReduceAddToExpressionsAction,
	{_State403, _WildcardMarker}:                  _ReduceExpressionsToOptionalExpressionsAction,
	{_State404, _WildcardMarker}:                  _ReduceJumpStatementToSimpleStatementBodyAction,
	{_State405, _WildcardMarker}:                  _ReduceAddToGenericParameterDefsAction,
	{_State406, _EndMarker}:                       _ReduceConstrainedDefToTypeDefAction,
	{_State407, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State409, _WildcardMarker}:                  _ReduceToAnonymousFuncExprAction,
	{_State410, RparenToken}:                      _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State411, _WildcardMarker}:                  _ReducePairToImplicitEnumValueDefsAction,
	{_State411, RparenToken}:                      _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State412, RparenToken}:                      _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State413, _WildcardMarker}:                  _ReduceAddToImplicitEnumValueDefsAction,
	{_State413, RparenToken}:                      _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State414, _WildcardMarker}:                  _ReduceVarargToParameterDeclAction,
	{_State415, _WildcardMarker}:                  _ReduceToFuncTypeAction,
	{_State416, _WildcardMarker}:                  _ReduceAddToParameterDeclsAction,
	{_State418, RparenToken}:                      _ReduceNilToOptionalParameterDeclsAction,
	{_State419, _WildcardMarker}:                  _ReduceExplicitToTraitPropertiesAction,
	{_State420, _WildcardMarker}:                  _ReduceImplicitToTraitPropertiesAction,
	{_State422, LbraceToken}:                      _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State422, DoToken}:                          _ReduceNilToOptionalSequenceExprAction,
	{_State423, _WildcardMarker}:                  _ReduceEnumMatchPatternToCasePatternAction,
	{_State425, LbraceToken}:                      _ReduceCaseToConditionAction,
	{_State426, _WildcardMarker}:                  _ReduceMultipleToCasePatternsAction,
	{_State427, _EndMarker}:                       _ReduceMultiIfElseToIfExprAction,
	{_State428, _EndMarker}:                       _ReduceIfElseToIfExprAction,
	{_State429, _WildcardMarker}:                  _ReduceCaseBranchStatementToStatementBodyAction,
	{_State430, _WildcardMarker}:                  _ReduceExplicitToImportClauseTerminalAction,
	{_State431, _WildcardMarker}:                  _ReduceImplicitToImportClauseTerminalAction,
	{_State432, _WildcardMarker}:                  _ReduceMultipleToImportStatementAction,
	{_State433, _WildcardMarker}:                  _ReduceAddToImportClausesAction,
	{_State434, _WildcardMarker}:                  _ReduceAliasToImportClauseAction,
	{_State436, RparenToken}:                      _ReduceNilToOptionalParameterDefsAction,
	{_State437, _WildcardMarker}:                  _ReduceToUnsafeStatementAction,
	{_State439, _EndMarker}:                       _ReduceIteratorToLoopExprAction,
	{_State441, _WildcardMarker}:                  _ReduceEnumVarDeclPatternToCasePatternAction,
	{_State442, _EndMarker}:                       _ReduceFuncDefToNamedFuncDefAction,
	{_State444, _WildcardMarker}:                  _ReduceNilToReturnTypeAction,
	{_State446, LbraceToken}:                      _ReduceNilToReturnTypeAction,
	{_State447, _WildcardMarker}:                  _ReduceToMethodSignatureAction,
	{_State448, _EndMarker}:                       _ReduceForToLoopExprAction,
	{_State450, _EndMarker}:                       _ReduceMethodDefToNamedFuncDefAction,
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 10
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

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
      COMMENT_GROUPS -> State 58
      PACKAGE -> State 14
      TYPE -> State 15
      FUNC -> State 16
      VAR -> State 37
      LET -> State 27
      LBRACE -> State 59
      definitions -> State 61
      definition -> State 60
      statement_block -> State 64
      var_decl_pattern -> State 66
      var_or_let -> State 57
      type_def -> State 65
      named_func_def -> State 62
      package_def -> State 63

  State 14:
    Kernel Items:
      package_def: PACKAGE., *
      package_def: PACKAGE.statement_block
    Reduce:
      * -> [package_def]
    Goto:
      LBRACE -> State 59
      statement_block -> State 67

  State 15:
    Kernel Items:
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE.IDENTIFIER ASSIGN value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 68

  State 16:
    Kernel Items:
      named_func_def: FUNC.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.IDENTIFIER ASSIGN expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 69
      LPAREN -> State 70

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
      LPAREN -> State 71

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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      optional_label_decl -> State 72
      sequence_expr -> State 73
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 23:
    Kernel Items:
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
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
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 94
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

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
      IDENTIFIER -> State 96
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
      DOT_DOT_DOT -> State 95
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 100
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      optional_arguments -> State 101
      arguments -> State 98
      argument -> State 97
      colon_expressions -> State 99
      optional_expression -> State 102
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

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
      atom_expr: PARSE_ERROR., *
    Reduce:
      * -> [atom_expr]
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
      LPAREN -> State 103

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
      EQUAL -> State 115
      NOT_EQUAL -> State 120
      LESS -> State 118
      LESS_OR_EQUAL -> State 119
      GREATER -> State 116
      GREATER_OR_EQUAL -> State 117
      cmp_op -> State 121

  State 45:
    Kernel Items:
      initializable_type: explicit_struct_def., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 46:
    Kernel Items:
      atom_expr: implicit_struct_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 47:
    Kernel Items:
      atom_expr: initializable_type.LPAREN optional_arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 122

  State 48:
    Kernel Items:
      atom_expr: literal_expr., *
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
      MUL -> State 128
      DIV -> State 126
      MOD -> State 127
      BIT_AND -> State 123
      BIT_LSHIFT -> State 124
      BIT_RSHIFT -> State 125
      mul_op -> State 129

  State 50:
    Kernel Items:
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
      atom_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    Goto:
      IF -> State 132
      SWITCH -> State 133
      FOR -> State 131
      DO -> State 130
      LBRACE -> State 59
      statement_block -> State 136
      if_expr -> State 134
      switch_expr -> State 137
      loop_expr -> State 135

  State 51:
    Kernel Items:
      sequence_expr: or_expr., *
      or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    Goto:
      OR -> State 138

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
      optional_label_decl -> State 72
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 139
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 55:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 56:
    Kernel Items:
      sequence_expr: var_decl_pattern., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 57:
    Kernel Items:
      var_decl_pattern: var_or_let.var_pattern optional_value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 140
      LPAREN -> State 141
      var_pattern -> State 143
      tuple_pattern -> State 142

  State 58:
    Kernel Items:
      definition: COMMENT_GROUPS., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 59:
    Kernel Items:
      statement_block: LBRACE.statements RBRACE
    Reduce:
      * -> [statements]
    Goto:
      statements -> State 144

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
      NEWLINES -> State 145
      optional_newlines -> State 146

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
      definition: statement_block., *
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
      definition: var_decl_pattern., *
      definition: var_decl_pattern.ASSIGN expression
    Reduce:
      * -> [definition]
    Goto:
      ASSIGN -> State 147

  State 67:
    Kernel Items:
      package_def: PACKAGE statement_block., $
    Reduce:
      $ -> [package_def]
    Goto:
      (nil)

  State 68:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE IDENTIFIER.ASSIGN value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 149
      ASSIGN -> State 148
      optional_generic_parameters -> State 150

  State 69:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC IDENTIFIER.ASSIGN expression
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 149
      ASSIGN -> State 151
      optional_generic_parameters -> State 152

  State 70:
    Kernel Items:
      named_func_def: FUNC LPAREN.parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 153
      parameter_def -> State 154

  State 71:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 153
      parameter_def -> State 156
      parameter_defs -> State 157
      optional_parameter_defs -> State 155

  State 72:
    Kernel Items:
      atom_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      statement_block -> State 136

  State 73:
    Kernel Items:
      sequence_expr: GREATER sequence_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 74:
    Kernel Items:
      returnable_type: BIT_AND.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 158
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 75:
    Kernel Items:
      returnable_type: BIT_NEG.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 159
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 76:
    Kernel Items:
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 160

  State 77:
    Kernel Items:
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 161

  State 78:
    Kernel Items:
      returnable_type: EXCLAIM.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 162
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 79:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 163

  State 80:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOT -> State 164
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 165

  State 81:
    Kernel Items:
      implicit_struct_def: LPAREN.optional_implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 169
      implicit_field_defs -> State 171
      optional_implicit_field_defs -> State 172
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      enum_value_def -> State 168
      implicit_enum_value_defs -> State 170
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 82:
    Kernel Items:
      atom_type: PARSE_ERROR., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      returnable_type: QUESTION.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 175
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 84:
    Kernel Items:
      returnable_type: TILDE_TILDE.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 176
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 85:
    Kernel Items:
      trait_def: TRAIT.LPAREN optional_trait_properties RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 177

  State 86:
    Kernel Items:
      returnable_type: atom_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 87:
    Kernel Items:
      atom_type: explicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      atom_type: func_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      atom_type: implicit_enum_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 90:
    Kernel Items:
      atom_type: implicit_struct_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 91:
    Kernel Items:
      atom_type: initializable_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 92:
    Kernel Items:
      value_type: returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 93:
    Kernel Items:
      atom_type: trait_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 94:
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
      RBRACKET -> State 182
      COMMA -> State 180
      COLON -> State 179
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 95:
    Kernel Items:
      argument: DOT_DOT_DOT., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 96:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expression
      atom_expr: IDENTIFIER., *
    Reduce:
      * -> [atom_expr]
    Goto:
      ASSIGN -> State 184

  State 97:
    Kernel Items:
      arguments: argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 98:
    Kernel Items:
      optional_arguments: arguments., RPAREN
      arguments: arguments.COMMA argument
    Reduce:
      RPAREN -> [optional_arguments]
    Goto:
      COMMA -> State 185

  State 99:
    Kernel Items:
      argument: colon_expressions., *
      colon_expressions: colon_expressions.COLON optional_expression
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 186

  State 100:
    Kernel Items:
      argument: expression., *
      optional_expression: expression., COLON
    Reduce:
      * -> [argument]
      COLON -> [optional_expression]
    Goto:
      (nil)

  State 101:
    Kernel Items:
      implicit_struct_expr: LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 187

  State 102:
    Kernel Items:
      colon_expressions: optional_expression.COLON optional_expression
    Reduce:
      (nil)
    Goto:
      COLON -> State 188

  State 103:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN.optional_explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 190
      implicit_struct_def -> State 90
      explicit_field_defs -> State 189
      optional_explicit_field_defs -> State 191
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 104:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.optional_generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      optional_generic_arguments -> State 193
      generic_arguments -> State 192
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 194
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 105:
    Kernel Items:
      access_expr: access_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 195

  State 106:
    Kernel Items:
      access_expr: access_expr LBRACKET.argument RBRACKET
    Reduce:
      * -> [optional_label_decl]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 96
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
      DOT_DOT_DOT -> State 95
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 100
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      argument -> State 196
      colon_expressions -> State 99
      optional_expression -> State 102
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
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
      LPAREN -> State 197

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
      optional_label_decl -> State 72
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 198
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 114:
    Kernel Items:
      and_expr: and_expr AND.cmp_expr
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
      optional_label_decl -> State 72
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 199
      initializable_type -> State 47
      explicit_struct_def -> State 45
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
      optional_label_decl -> State 72
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 200
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 122:
    Kernel Items:
      atom_expr: initializable_type LPAREN.optional_arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_arguments]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 96
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
      DOT_DOT_DOT -> State 95
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 100
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      optional_arguments -> State 201
      arguments -> State 98
      argument -> State 97
      colon_expressions -> State 99
      optional_expression -> State 102
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
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
      optional_label_decl -> State 72
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 202
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 130:
    Kernel Items:
      loop_expr: DO.statement_block
      loop_expr: DO.statement_block FOR sequence_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      statement_block -> State 203

  State 131:
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      assign_pattern -> State 204
      optional_label_decl -> State 72
      sequence_expr -> State 206
      for_assignment -> State 205
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 132:
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
      CASE -> State 207
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      optional_label_decl -> State 72
      sequence_expr -> State 209
      condition -> State 208
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 133:
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      optional_label_decl -> State 72
      sequence_expr -> State 210
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 134:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 135:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 136:
    Kernel Items:
      atom_expr: optional_label_decl statement_block., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 137:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 138:
    Kernel Items:
      or_expr: or_expr OR.and_expr
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
      optional_label_decl -> State 72
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 211
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 139:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefix_unary_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 140:
    Kernel Items:
      var_pattern: IDENTIFIER., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 141:
    Kernel Items:
      tuple_pattern: LPAREN.field_var_patterns RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 213
      LPAREN -> State 141
      DOT_DOT_DOT -> State 212
      var_pattern -> State 216
      tuple_pattern -> State 142
      field_var_patterns -> State 215
      field_var_pattern -> State 214

  State 142:
    Kernel Items:
      var_pattern: tuple_pattern., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 143:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern.optional_value_type
    Reduce:
      * -> [optional_value_type]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      optional_value_type -> State 217
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 218
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 144:
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
      CASE -> State 221
      DEFAULT -> State 223
      RETURN -> State 228
      BREAK -> State 220
      CONTINUE -> State 222
      FALLTHROUGH -> State 225
      IMPORT -> State 226
      UNSAFE -> State 167
      STRUCT -> State 34
      FUNC -> State 21
      ASYNC -> State 219
      DEFER -> State 224
      VAR -> State 37
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      RBRACE -> State 227
      LPAREN -> State 28
      LBRACKET -> State 26
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      statement -> State 237
      simple_statement_body -> State 236
      statement_body -> State 238
      unsafe_statement -> State 239
      jump_type -> State 234
      expressions -> State 232
      import_statement -> State 233
      var_decl_pattern -> State 56
      var_or_let -> State 57
      assign_pattern -> State 230
      expression -> State 231
      optional_label_decl -> State 50
      sequence_expr -> State 235
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 229
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 145:
    Kernel Items:
      optional_newlines: NEWLINES., $
      definitions: definitions NEWLINES.definition
    Reduce:
      $ -> [optional_newlines]
    Goto:
      COMMENT_GROUPS -> State 58
      PACKAGE -> State 14
      TYPE -> State 15
      FUNC -> State 16
      VAR -> State 37
      LET -> State 27
      LBRACE -> State 59
      definition -> State 240
      statement_block -> State 64
      var_decl_pattern -> State 66
      var_or_let -> State 57
      type_def -> State 65
      named_func_def -> State 62
      package_def -> State 63

  State 146:
    Kernel Items:
      optional_definitions: optional_newlines definitions optional_newlines., $
    Reduce:
      $ -> [optional_definitions]
    Goto:
      (nil)

  State 147:
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 241
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 148:
    Kernel Items:
      type_def: TYPE IDENTIFIER ASSIGN.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 242
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 149:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.optional_generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 243
      generic_parameter_def -> State 244
      generic_parameter_defs -> State 245
      optional_generic_parameter_defs -> State 246

  State 150:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 247
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 151:
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 248
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 152:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 249

  State 153:
    Kernel Items:
      parameter_def: IDENTIFIER., *
      parameter_def: IDENTIFIER.DOT_DOT_DOT
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.DOT_DOT_DOT value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      DOT_DOT_DOT -> State 250
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 251
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 154:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def.RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 252

  State 155:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 253

  State 156:
    Kernel Items:
      parameter_defs: parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 157:
    Kernel Items:
      parameter_defs: parameter_defs.COMMA parameter_def
      optional_parameter_defs: parameter_defs., RPAREN
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      COMMA -> State 254

  State 158:
    Kernel Items:
      returnable_type: BIT_AND returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 159:
    Kernel Items:
      returnable_type: BIT_NEG returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 160:
    Kernel Items:
      atom_type: DOT optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 161:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 257
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      enum_value_def -> State 255
      implicit_enum_value_defs -> State 258
      implicit_enum_def -> State 89
      explicit_enum_value_defs -> State 256
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 162:
    Kernel Items:
      returnable_type: EXCLAIM returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 163:
    Kernel Items:
      func_type: FUNC LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 260
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      DOT_DOT_DOT -> State 259
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 264
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      parameter_decl -> State 262
      parameter_decls -> State 263
      optional_parameter_decls -> State 261
      func_type -> State 88

  State 164:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 265

  State 165:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 166:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 266
      QUESTION -> State 83
      EXCLAIM -> State 78
      DOLLAR_LBRACKET -> State 104
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      optional_generic_binding -> State 165
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 267
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 167:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 268

  State 168:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 269

  State 169:
    Kernel Items:
      implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 270

  State 170:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      OR -> State 271
      RPAREN -> State 272

  State 171:
    Kernel Items:
      implicit_field_defs: implicit_field_defs.COMMA field_def
      optional_implicit_field_defs: implicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_implicit_field_defs]
    Goto:
      COMMA -> State 273

  State 172:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 274

  State 173:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 174:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 175:
    Kernel Items:
      returnable_type: QUESTION returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 176:
    Kernel Items:
      returnable_type: TILDE_TILDE returnable_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 177:
    Kernel Items:
      trait_def: TRAIT LPAREN.optional_trait_properties RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 275
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 276
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_property -> State 280
      trait_properties -> State 279
      optional_trait_properties -> State 278
      trait_def -> State 93
      func_type -> State 88
      method_signature -> State 277

  State 178:
    Kernel Items:
      value_type: value_type ADD.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 281
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 179:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON.value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 282
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 180:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 283

  State 181:
    Kernel Items:
      value_type: value_type MUL.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 284
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 182:
    Kernel Items:
      initializable_type: LBRACKET value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 183:
    Kernel Items:
      value_type: value_type SUB.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 285
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 184:
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 286
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 185:
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
      IDENTIFIER -> State 96
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
      DOT_DOT_DOT -> State 95
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 100
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      argument -> State 287
      colon_expressions -> State 99
      optional_expression -> State 102
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 186:
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 288
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      optional_expression -> State 289
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 187:
    Kernel Items:
      implicit_struct_expr: LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [implicit_struct_expr]
    Goto:
      (nil)

  State 188:
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 288
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      optional_expression -> State 290
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 189:
    Kernel Items:
      explicit_field_defs: explicit_field_defs.NEWLINES field_def
      explicit_field_defs: explicit_field_defs.COMMA field_def
      optional_explicit_field_defs: explicit_field_defs., RPAREN
    Reduce:
      RPAREN -> [optional_explicit_field_defs]
    Goto:
      NEWLINES -> State 292
      COMMA -> State 291

  State 190:
    Kernel Items:
      explicit_field_defs: field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 191:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 293

  State 192:
    Kernel Items:
      optional_generic_arguments: generic_arguments., RBRACKET
      generic_arguments: generic_arguments.COMMA value_type
    Reduce:
      RBRACKET -> [optional_generic_arguments]
    Goto:
      COMMA -> State 294

  State 193:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 295

  State 194:
    Kernel Items:
      generic_arguments: value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 195:
    Kernel Items:
      access_expr: access_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 196:
    Kernel Items:
      access_expr: access_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 296

  State 197:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN.optional_arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [optional_arguments]
      COLON -> [optional_expression]
    Goto:
      INTEGER_LITERAL -> State 24
      FLOAT_LITERAL -> State 20
      RUNE_LITERAL -> State 32
      STRING_LITERAL -> State 33
      IDENTIFIER -> State 96
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
      DOT_DOT_DOT -> State 95
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 100
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      optional_arguments -> State 297
      arguments -> State 98
      argument -> State 97
      colon_expressions -> State 99
      optional_expression -> State 102
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 198:
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

  State 199:
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

  State 200:
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

  State 201:
    Kernel Items:
      atom_expr: initializable_type LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 298

  State 202:
    Kernel Items:
      mul_expr: mul_expr mul_op prefix_unary_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 203:
    Kernel Items:
      loop_expr: DO statement_block., *
      loop_expr: DO statement_block.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 299

  State 204:
    Kernel Items:
      loop_expr: FOR assign_pattern.IN sequence_expr DO statement_block
      for_assignment: assign_pattern.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      IN -> State 301
      ASSIGN -> State 300

  State 205:
    Kernel Items:
      loop_expr: FOR for_assignment.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 302

  State 206:
    Kernel Items:
      assign_pattern: sequence_expr., *
      loop_expr: FOR sequence_expr.DO statement_block
      for_assignment: sequence_expr., SEMICOLON
    Reduce:
      * -> [assign_pattern]
      SEMICOLON -> [for_assignment]
    Goto:
      DO -> State 303

  State 207:
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
      VAR -> State 305
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT -> State 304
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      case_patterns -> State 308
      var_decl_pattern -> State 56
      var_or_let -> State 57
      assign_pattern -> State 306
      case_pattern -> State 307
      optional_label_decl -> State 72
      sequence_expr -> State 309
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 208:
    Kernel Items:
      if_expr: IF condition.statement_block
      if_expr: IF condition.statement_block ELSE statement_block
      if_expr: IF condition.statement_block ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      statement_block -> State 310

  State 209:
    Kernel Items:
      condition: sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      statement_block -> State 311

  State 211:
    Kernel Items:
      and_expr: and_expr.AND cmp_expr
      or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 114

  State 212:
    Kernel Items:
      field_var_pattern: DOT_DOT_DOT., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      var_pattern: IDENTIFIER., *
      field_var_pattern: IDENTIFIER.ASSIGN var_pattern
    Reduce:
      * -> [var_pattern]
    Goto:
      ASSIGN -> State 312

  State 214:
    Kernel Items:
      field_var_patterns: field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 215:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns.RPAREN
      field_var_patterns: field_var_patterns.COMMA field_var_pattern
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 314
      COMMA -> State 313

  State 216:
    Kernel Items:
      field_var_pattern: var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern optional_value_type., $
    Reduce:
      $ -> [var_decl_pattern]
    Goto:
      (nil)

  State 218:
    Kernel Items:
      optional_value_type: value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [optional_value_type]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 219:
    Kernel Items:
      simple_statement_body: ASYNC.call_expr
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
      optional_label_decl -> State 72
      call_expr -> State 316
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 315
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 220:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 221:
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
      VAR -> State 305
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT -> State 304
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      case_patterns -> State 317
      var_decl_pattern -> State 56
      var_or_let -> State 57
      assign_pattern -> State 306
      case_pattern -> State 307
      optional_label_decl -> State 72
      sequence_expr -> State 309
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 222:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 223:
    Kernel Items:
      statement_body: DEFAULT.COLON optional_simple_statement_body
    Reduce:
      (nil)
    Goto:
      COLON -> State 318

  State 224:
    Kernel Items:
      simple_statement_body: DEFER.call_expr
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
      optional_label_decl -> State 72
      call_expr -> State 319
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 315
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 225:
    Kernel Items:
      simple_statement_body: FALLTHROUGH., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 226:
    Kernel Items:
      import_statement: IMPORT.import_clause
      import_statement: IMPORT.LPAREN import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 321
      LPAREN -> State 320
      import_clause -> State 322

  State 227:
    Kernel Items:
      statement_block: LBRACE statements RBRACE., $
    Reduce:
      $ -> [statement_block]
    Goto:
      (nil)

  State 228:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 229:
    Kernel Items:
      simple_statement_body: access_expr.unary_op_assign
      simple_statement_body: access_expr.binary_op_assign expression
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
      ADD_ASSIGN -> State 323
      SUB_ASSIGN -> State 334
      MUL_ASSIGN -> State 333
      DIV_ASSIGN -> State 331
      MOD_ASSIGN -> State 332
      ADD_ONE_ASSIGN -> State 324
      SUB_ONE_ASSIGN -> State 335
      BIT_NEG_ASSIGN -> State 327
      BIT_AND_ASSIGN -> State 325
      BIT_OR_ASSIGN -> State 328
      BIT_XOR_ASSIGN -> State 330
      BIT_LSHIFT_ASSIGN -> State 326
      BIT_RSHIFT_ASSIGN -> State 329
      unary_op_assign -> State 337
      binary_op_assign -> State 336
      optional_generic_binding -> State 108

  State 230:
    Kernel Items:
      simple_statement_body: assign_pattern.ASSIGN expression
    Reduce:
      (nil)
    Goto:
      ASSIGN -> State 338

  State 231:
    Kernel Items:
      expressions: expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 232:
    Kernel Items:
      simple_statement_body: expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [simple_statement_body]
    Goto:
      COMMA -> State 339

  State 233:
    Kernel Items:
      statement_body: import_statement., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 234:
    Kernel Items:
      simple_statement_body: jump_type.optional_jump_label optional_expressions
    Reduce:
      * -> [optional_jump_label]
    Goto:
      JUMP_LABEL -> State 340
      optional_jump_label -> State 341

  State 235:
    Kernel Items:
      assign_pattern: sequence_expr., ASSIGN
      expression: sequence_expr., *
    Reduce:
      * -> [expression]
      ASSIGN -> [assign_pattern]
    Goto:
      (nil)

  State 236:
    Kernel Items:
      statement_body: simple_statement_body., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

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
      NEWLINES -> State 342
      SEMICOLON -> State 343

  State 239:
    Kernel Items:
      simple_statement_body: unsafe_statement., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      definitions: definitions NEWLINES definition., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      definition: var_decl_pattern ASSIGN expression., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 242:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER ASSIGN value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 243:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 344
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 244:
    Kernel Items:
      generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 245:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs.COMMA generic_parameter_def
      optional_generic_parameter_defs: generic_parameter_defs., RBRACKET
    Reduce:
      RBRACKET -> [optional_generic_parameter_defs]
    Goto:
      COMMA -> State 345

  State 246:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 346

  State 247:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., *
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 347
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 248:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN expression., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 249:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 153
      parameter_def -> State 156
      parameter_defs -> State 157
      optional_parameter_defs -> State 348

  State 250:
    Kernel Items:
      parameter_def: IDENTIFIER DOT_DOT_DOT., *
      parameter_def: IDENTIFIER DOT_DOT_DOT.value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 349
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 251:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 252:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN.IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 350

  State 253:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 352
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      return_type -> State 351
      func_type -> State 88

  State 254:
    Kernel Items:
      parameter_defs: parameter_defs COMMA.parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 153
      parameter_def -> State 353

  State 255:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 354
      OR -> State 355

  State 256:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 356

  State 257:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 270

  State 258:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 357
      OR -> State 358

  State 259:
    Kernel Items:
      parameter_decl: DOT_DOT_DOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 359
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 260:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.DOT_DOT_DOT value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 266
      QUESTION -> State 83
      EXCLAIM -> State 78
      DOLLAR_LBRACKET -> State 104
      DOT_DOT_DOT -> State 360
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      optional_generic_binding -> State 165
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 361
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 261:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 362

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
      COMMA -> State 363

  State 264:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 265:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 364

  State 266:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 265
      DOLLAR_LBRACKET -> State 104
      optional_generic_binding -> State 160

  State 267:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 268:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 365

  State 269:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 257
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      enum_value_def -> State 366
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 270:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 367

  State 271:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 257
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      enum_value_def -> State 368
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 272:
    Kernel Items:
      implicit_enum_def: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_def]
    Goto:
      (nil)

  State 273:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 369
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 274:
    Kernel Items:
      implicit_struct_def: LPAREN optional_implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_def]
    Goto:
      (nil)

  State 275:
    Kernel Items:
      func_type: FUNC.LPAREN optional_parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 370
      LPAREN -> State 163

  State 276:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 277:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 371

  State 279:
    Kernel Items:
      trait_properties: trait_properties.NEWLINES trait_property
      trait_properties: trait_properties.COMMA trait_property
      optional_trait_properties: trait_properties., RPAREN
    Reduce:
      RPAREN -> [optional_trait_properties]
    Goto:
      NEWLINES -> State 373
      COMMA -> State 372

  State 280:
    Kernel Items:
      trait_properties: trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 281:
    Kernel Items:
      value_type: value_type ADD returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 282:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type.RBRACKET
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 374
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 283:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 375

  State 284:
    Kernel Items:
      value_type: value_type MUL returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 285:
    Kernel Items:
      value_type: value_type SUB returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 286:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 287:
    Kernel Items:
      arguments: arguments COMMA argument., *
    Reduce:
      * -> [arguments]
    Goto:
      (nil)

  State 288:
    Kernel Items:
      optional_expression: expression., *
    Reduce:
      * -> [optional_expression]
    Goto:
      (nil)

  State 289:
    Kernel Items:
      colon_expressions: colon_expressions COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 290:
    Kernel Items:
      colon_expressions: optional_expression COLON optional_expression., *
    Reduce:
      * -> [colon_expressions]
    Goto:
      (nil)

  State 291:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 376
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 292:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 377
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 293:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN optional_explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_def]
    Goto:
      (nil)

  State 294:
    Kernel Items:
      generic_arguments: generic_arguments COMMA.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 378
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 295:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET optional_generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 296:
    Kernel Items:
      access_expr: access_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 297:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 379

  State 298:
    Kernel Items:
      atom_expr: initializable_type LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 299:
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      optional_label_decl -> State 72
      sequence_expr -> State 380
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 300:
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      optional_label_decl -> State 72
      sequence_expr -> State 381
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 301:
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      optional_label_decl -> State 72
      sequence_expr -> State 382
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 302:
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      optional_label_decl -> State 72
      sequence_expr -> State 384
      optional_sequence_expr -> State 383
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 303:
    Kernel Items:
      loop_expr: FOR sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      statement_block -> State 385

  State 304:
    Kernel Items:
      case_pattern: DOT.IDENTIFIER implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 386

  State 305:
    Kernel Items:
      var_or_let: VAR., *
      case_pattern: VAR.DOT IDENTIFIER tuple_pattern
    Reduce:
      * -> [var_or_let]
    Goto:
      DOT -> State 387

  State 306:
    Kernel Items:
      case_pattern: assign_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 307:
    Kernel Items:
      case_patterns: case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 308:
    Kernel Items:
      case_patterns: case_patterns.COMMA case_pattern
      condition: CASE case_patterns.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      COMMA -> State 389
      ASSIGN -> State 388

  State 309:
    Kernel Items:
      assign_pattern: sequence_expr., *
    Reduce:
      * -> [assign_pattern]
    Goto:
      (nil)

  State 310:
    Kernel Items:
      if_expr: IF condition statement_block., *
      if_expr: IF condition statement_block.ELSE statement_block
      if_expr: IF condition statement_block.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 390

  State 311:
    Kernel Items:
      switch_expr: SWITCH sequence_expr statement_block., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 312:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN.var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 140
      LPAREN -> State 141
      var_pattern -> State 391
      tuple_pattern -> State 142

  State 313:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA.field_var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 213
      LPAREN -> State 141
      DOT_DOT_DOT -> State 212
      var_pattern -> State 216
      tuple_pattern -> State 142
      field_var_pattern -> State 392

  State 314:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns RPAREN., *
    Reduce:
      * -> [tuple_pattern]
    Goto:
      (nil)

  State 315:
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

  State 316:
    Kernel Items:
      simple_statement_body: ASYNC call_expr., NEWLINES
      simple_statement_body: ASYNC call_expr., SEMICOLON
      access_expr: call_expr., *
    Reduce:
      * -> [access_expr]
      NEWLINES -> [simple_statement_body]
      SEMICOLON -> [simple_statement_body]
    Goto:
      (nil)

  State 317:
    Kernel Items:
      statement_body: CASE case_patterns.COLON optional_simple_statement_body
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    Goto:
      COMMA -> State 389
      COLON -> State 393

  State 318:
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
      RETURN -> State 228
      BREAK -> State 220
      CONTINUE -> State 222
      FALLTHROUGH -> State 225
      UNSAFE -> State 167
      STRUCT -> State 34
      FUNC -> State 21
      ASYNC -> State 219
      DEFER -> State 224
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
      simple_statement_body -> State 395
      optional_simple_statement_body -> State 394
      unsafe_statement -> State 239
      jump_type -> State 234
      expressions -> State 232
      var_decl_pattern -> State 56
      var_or_let -> State 57
      assign_pattern -> State 230
      expression -> State 231
      optional_label_decl -> State 50
      sequence_expr -> State 235
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 229
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 319:
    Kernel Items:
      simple_statement_body: DEFER call_expr., NEWLINES
      simple_statement_body: DEFER call_expr., SEMICOLON
      access_expr: call_expr., *
    Reduce:
      * -> [access_expr]
      NEWLINES -> [simple_statement_body]
      SEMICOLON -> [simple_statement_body]
    Goto:
      (nil)

  State 320:
    Kernel Items:
      import_statement: IMPORT LPAREN.import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 321
      import_clause -> State 396
      import_clause_terminal -> State 397
      import_clauses -> State 398

  State 321:
    Kernel Items:
      import_clause: STRING_LITERAL., *
      import_clause: STRING_LITERAL.AS IDENTIFIER
    Reduce:
      * -> [import_clause]
    Goto:
      AS -> State 399

  State 322:
    Kernel Items:
      import_statement: IMPORT import_clause., *
    Reduce:
      * -> [import_statement]
    Goto:
      (nil)

  State 323:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 324:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 325:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 326:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 327:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 328:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 329:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 330:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 331:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 333:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 334:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 335:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 336:
    Kernel Items:
      simple_statement_body: access_expr binary_op_assign.expression
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 400
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 337:
    Kernel Items:
      simple_statement_body: access_expr unary_op_assign., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 338:
    Kernel Items:
      simple_statement_body: assign_pattern ASSIGN.expression
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 401
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 339:
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 402
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 340:
    Kernel Items:
      optional_jump_label: JUMP_LABEL., *
    Reduce:
      * -> [optional_jump_label]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      simple_statement_body: jump_type optional_jump_label.optional_expressions
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_expressions]
      SEMICOLON -> [optional_expressions]
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
      expressions -> State 403
      optional_expressions -> State 404
      var_decl_pattern -> State 56
      var_or_let -> State 57
      expression -> State 231
      optional_label_decl -> State 50
      sequence_expr -> State 55
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 342:
    Kernel Items:
      statement: statement_body NEWLINES., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 343:
    Kernel Items:
      statement: statement_body SEMICOLON., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 344:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 345:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA.generic_parameter_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 243
      generic_parameter_def -> State 405

  State 346:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET optional_generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 347:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 406
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 348:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 407

  State 349:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_def: IDENTIFIER DOT_DOT_DOT value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 350:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER.optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 149
      optional_generic_parameters -> State 408

  State 351:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      statement_block -> State 409

  State 352:
    Kernel Items:
      return_type: returnable_type., *
    Reduce:
      * -> [return_type]
    Goto:
      (nil)

  State 353:
    Kernel Items:
      parameter_defs: parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [parameter_defs]
    Goto:
      (nil)

  State 354:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 257
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      enum_value_def -> State 410
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 355:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 257
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      enum_value_def -> State 411
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 356:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_def]
    Goto:
      (nil)

  State 357:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 257
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      enum_value_def -> State 412
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 358:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 257
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      enum_value_def -> State 413
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 359:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: DOT_DOT_DOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 360:
    Kernel Items:
      parameter_decl: IDENTIFIER DOT_DOT_DOT.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 414
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      func_type -> State 88

  State 361:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 362:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 352
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      return_type -> State 415
      func_type -> State 88

  State 363:
    Kernel Items:
      parameter_decls: parameter_decls COMMA.parameter_decl
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 260
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      DOT_DOT_DOT -> State 259
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 264
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      parameter_decl -> State 416
      func_type -> State 88

  State 364:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 365:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 417

  State 366:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 367:
    Kernel Items:
      enum_value_def: field_def ASSIGN DEFAULT., *
    Reduce:
      * -> [enum_value_def]
    Goto:
      (nil)

  State 368:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 369:
    Kernel Items:
      implicit_field_defs: implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [implicit_field_defs]
    Goto:
      (nil)

  State 370:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN optional_parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 418

  State 371:
    Kernel Items:
      trait_def: TRAIT LPAREN optional_trait_properties RPAREN., *
    Reduce:
      * -> [trait_def]
    Goto:
      (nil)

  State 372:
    Kernel Items:
      trait_properties: trait_properties COMMA.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 275
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 276
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_property -> State 419
      trait_def -> State 93
      func_type -> State 88
      method_signature -> State 277

  State 373:
    Kernel Items:
      trait_properties: trait_properties NEWLINES.trait_property
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 166
      UNSAFE -> State 167
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 275
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      unsafe_statement -> State 173
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 174
      field_def -> State 276
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_property -> State 420
      trait_def -> State 93
      func_type -> State 88
      method_signature -> State 277

  State 374:
    Kernel Items:
      initializable_type: LBRACKET value_type COLON value_type RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 375:
    Kernel Items:
      initializable_type: LBRACKET value_type COMMA INTEGER_LITERAL RBRACKET., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 376:
    Kernel Items:
      explicit_field_defs: explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 377:
    Kernel Items:
      explicit_field_defs: explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [explicit_field_defs]
    Goto:
      (nil)

  State 378:
    Kernel Items:
      generic_arguments: generic_arguments COMMA value_type., *
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
    Reduce:
      * -> [generic_arguments]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 379:
    Kernel Items:
      call_expr: access_expr optional_generic_binding LPAREN optional_arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 380:
    Kernel Items:
      loop_expr: DO statement_block FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 381:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN sequence_expr., SEMICOLON
    Reduce:
      SEMICOLON -> [for_assignment]
    Goto:
      (nil)

  State 382:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 421

  State 383:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 422

  State 384:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 385:
    Kernel Items:
      loop_expr: FOR sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 386:
    Kernel Items:
      case_pattern: DOT IDENTIFIER.implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 28
      implicit_struct_expr -> State 423

  State 387:
    Kernel Items:
      case_pattern: VAR DOT.IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 424

  State 388:
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      optional_label_decl -> State 72
      sequence_expr -> State 425
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 389:
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
      VAR -> State 305
      LET -> State 27
      NOT -> State 30
      LABEL_DECL -> State 25
      LPAREN -> State 28
      LBRACKET -> State 26
      DOT -> State 304
      SUB -> State 35
      MUL -> State 29
      BIT_NEG -> State 18
      BIT_AND -> State 17
      GREATER -> State 22
      PARSE_ERROR -> State 31
      var_decl_pattern -> State 56
      var_or_let -> State 57
      assign_pattern -> State 306
      case_pattern -> State 426
      optional_label_decl -> State 72
      sequence_expr -> State 309
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 390:
    Kernel Items:
      if_expr: IF condition statement_block ELSE.statement_block
      if_expr: IF condition statement_block ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 132
      LBRACE -> State 59
      statement_block -> State 428
      if_expr -> State 427

  State 391:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 392:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 393:
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
      RETURN -> State 228
      BREAK -> State 220
      CONTINUE -> State 222
      FALLTHROUGH -> State 225
      UNSAFE -> State 167
      STRUCT -> State 34
      FUNC -> State 21
      ASYNC -> State 219
      DEFER -> State 224
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
      simple_statement_body -> State 395
      optional_simple_statement_body -> State 429
      unsafe_statement -> State 239
      jump_type -> State 234
      expressions -> State 232
      var_decl_pattern -> State 56
      var_or_let -> State 57
      assign_pattern -> State 230
      expression -> State 231
      optional_label_decl -> State 50
      sequence_expr -> State 235
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 229
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 394:
    Kernel Items:
      statement_body: DEFAULT COLON optional_simple_statement_body., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 395:
    Kernel Items:
      optional_simple_statement_body: simple_statement_body., *
    Reduce:
      * -> [optional_simple_statement_body]
    Goto:
      (nil)

  State 396:
    Kernel Items:
      import_clause_terminal: import_clause.NEWLINES
      import_clause_terminal: import_clause.COMMA
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 431
      COMMA -> State 430

  State 397:
    Kernel Items:
      import_clauses: import_clause_terminal., *
    Reduce:
      * -> [import_clauses]
    Goto:
      (nil)

  State 398:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses.RPAREN
      import_clauses: import_clauses.import_clause_terminal
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 321
      RPAREN -> State 432
      import_clause -> State 396
      import_clause_terminal -> State 433

  State 399:
    Kernel Items:
      import_clause: STRING_LITERAL AS.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 434

  State 400:
    Kernel Items:
      simple_statement_body: access_expr binary_op_assign expression., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 401:
    Kernel Items:
      simple_statement_body: assign_pattern ASSIGN expression., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 402:
    Kernel Items:
      expressions: expressions COMMA expression., *
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
      COMMA -> State 339

  State 404:
    Kernel Items:
      simple_statement_body: jump_type optional_jump_label optional_expressions., *
    Reduce:
      * -> [simple_statement_body]
    Goto:
      (nil)

  State 405:
    Kernel Items:
      generic_parameter_defs: generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 406:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., $
    Reduce:
      $ -> [type_def]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 407:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 352
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      return_type -> State 435
      func_type -> State 88

  State 408:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters.LPAREN optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 436

  State 409:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN optional_parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 410:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 411:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
      explicit_enum_value_defs: enum_value_def OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 412:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 413:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 414:
    Kernel Items:
      value_type: value_type.MUL returnable_type
      value_type: value_type.ADD returnable_type
      value_type: value_type.SUB returnable_type
      parameter_decl: IDENTIFIER DOT_DOT_DOT value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 178
      SUB -> State 183
      MUL -> State 181

  State 415:
    Kernel Items:
      func_type: FUNC LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [func_type]
    Goto:
      (nil)

  State 416:
    Kernel Items:
      parameter_decls: parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [parameter_decls]
    Goto:
      (nil)

  State 417:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 437

  State 418:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.optional_parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [optional_parameter_decls]
    Goto:
      IDENTIFIER -> State 260
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      DOT_DOT_DOT -> State 259
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 92
      value_type -> State 264
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      parameter_decl -> State 262
      parameter_decls -> State 263
      optional_parameter_decls -> State 438
      func_type -> State 88

  State 419:
    Kernel Items:
      trait_properties: trait_properties COMMA trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 420:
    Kernel Items:
      trait_properties: trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [trait_properties]
    Goto:
      (nil)

  State 421:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      statement_block -> State 439

  State 422:
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
      var_decl_pattern -> State 56
      var_or_let -> State 57
      optional_label_decl -> State 72
      sequence_expr -> State 384
      optional_sequence_expr -> State 440
      call_expr -> State 43
      atom_expr -> State 42
      literal_expr -> State 48
      implicit_struct_expr -> State 46
      access_expr -> State 38
      postfix_unary_expr -> State 52
      prefix_unary_op -> State 54
      prefix_unary_expr -> State 53
      mul_expr -> State 49
      add_expr -> State 39
      cmp_expr -> State 44
      and_expr -> State 40
      or_expr -> State 51
      initializable_type -> State 47
      explicit_struct_def -> State 45
      anonymous_func_expr -> State 41

  State 423:
    Kernel Items:
      case_pattern: DOT IDENTIFIER implicit_struct_expr., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 424:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER.tuple_pattern
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 141
      tuple_pattern -> State 441

  State 425:
    Kernel Items:
      condition: CASE case_patterns ASSIGN sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 426:
    Kernel Items:
      case_patterns: case_patterns COMMA case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 427:
    Kernel Items:
      if_expr: IF condition statement_block ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 428:
    Kernel Items:
      if_expr: IF condition statement_block ELSE statement_block., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 429:
    Kernel Items:
      statement_body: CASE case_patterns COLON optional_simple_statement_body., *
    Reduce:
      * -> [statement_body]
    Goto:
      (nil)

  State 430:
    Kernel Items:
      import_clause_terminal: import_clause COMMA., *
    Reduce:
      * -> [import_clause_terminal]
    Goto:
      (nil)

  State 431:
    Kernel Items:
      import_clause_terminal: import_clause NEWLINES., *
    Reduce:
      * -> [import_clause_terminal]
    Goto:
      (nil)

  State 432:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses RPAREN., *
    Reduce:
      * -> [import_statement]
    Goto:
      (nil)

  State 433:
    Kernel Items:
      import_clauses: import_clauses import_clause_terminal., *
    Reduce:
      * -> [import_clauses]
    Goto:
      (nil)

  State 434:
    Kernel Items:
      import_clause: STRING_LITERAL AS IDENTIFIER., *
    Reduce:
      * -> [import_clause]
    Goto:
      (nil)

  State 435:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      statement_block -> State 442

  State 436:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN.optional_parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [optional_parameter_defs]
    Goto:
      IDENTIFIER -> State 153
      parameter_def -> State 156
      parameter_defs -> State 157
      optional_parameter_defs -> State 443

  State 437:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 438:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 444

  State 439:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 440:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 445

  State 441:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER tuple_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 442:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 443:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 446

  State 444:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 352
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      return_type -> State 447
      func_type -> State 88

  State 445:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      statement_block -> State 448

  State 446:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 80
      STRUCT -> State 34
      ENUM -> State 77
      TRAIT -> State 85
      FUNC -> State 79
      LPAREN -> State 81
      LBRACKET -> State 26
      DOT -> State 76
      QUESTION -> State 83
      EXCLAIM -> State 78
      TILDE_TILDE -> State 84
      BIT_NEG -> State 75
      BIT_AND -> State 74
      PARSE_ERROR -> State 82
      initializable_type -> State 91
      atom_type -> State 86
      returnable_type -> State 352
      implicit_struct_def -> State 90
      explicit_struct_def -> State 45
      implicit_enum_def -> State 89
      explicit_enum_def -> State 87
      trait_def -> State 93
      return_type -> State 449
      func_type -> State 88

  State 447:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN optional_parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 448:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 449:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 59
      statement_block -> State 450

  State 450:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN optional_parameter_defs RPAREN return_type statement_block., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

Number of states: 450
Number of shift actions: 3172
Number of reduce actions: 370
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 3936
Number of unoptimized shift actions: 29900
Number of unoptimized reduce actions: 22459
*/
