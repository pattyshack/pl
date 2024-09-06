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
	EllipsisToken        = SymbolId(309)
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
	ToSource(Definitions_ []SourceDefinition) ([]SourceDefinition, error)
}

type DefinitionsReducer interface {
	// 55:2: definitions -> add: ...
	AddToDefinitions(Definitions_ []SourceDefinition, Definition_ SourceDefinition, Newlines_ TokenCount) ([]SourceDefinition, error)

	// 56:2: definitions -> nil: ...
	NilToDefinitions() ([]SourceDefinition, error)
}

type DefinitionReducer interface {

	// 63:2: definition -> global_var_def: ...
	GlobalVarDefToDefinition(VarDeclPattern_ GenericSymbol) (SourceDefinition, error)

	// 64:2: definition -> global_var_assignment: ...
	GlobalVarAssignmentToDefinition(VarDeclPattern_ GenericSymbol, Assign_ TokenValue, Expression_ Expression) (SourceDefinition, error)

	// 67:2: definition -> statement_block: ...
	StatementBlockToDefinition(StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 71:2: definition -> COMMENT_GROUPS: ...
	CommentGroupsToDefinition(CommentGroups_ TokenValue) (SourceDefinition, error)
}

type StatementBlockReducer interface {
	// 91:19: statement_block -> ...
	ToStatementBlock(Lbrace_ TokenValue, Statements_ []Statement, Rbrace_ TokenValue) (GenericSymbol, error)
}

type ProperStatementsReducer interface {
	// 94:2: proper_statements -> add_implicit: ...
	AddImplicitToProperStatements(ProperStatements_ []Statement, Newlines_ TokenCount, Statement_ Statement) ([]Statement, error)

	// 95:2: proper_statements -> add_explicit: ...
	AddExplicitToProperStatements(ProperStatements_ []Statement, Semicolon_ TokenValue, Statement_ Statement) ([]Statement, error)

	// 96:2: proper_statements -> statement: ...
	StatementToProperStatements(Statement_ Statement) ([]Statement, error)
}

type StatementsReducer interface {

	// 100:2: statements -> improper_implicit: ...
	ImproperImplicitToStatements(ProperStatements_ []Statement, Newlines_ TokenCount) ([]Statement, error)

	// 101:2: statements -> improper_explicit: ...
	ImproperExplicitToStatements(ProperStatements_ []Statement, Semicolon_ TokenValue) ([]Statement, error)

	// 102:2: statements -> nil: ...
	NilToStatements() ([]Statement, error)
}

type StatementReducer interface {

	// 140:2: statement -> case_branch_statement: ...
	CaseBranchStatementToStatement(Case_ TokenValue, CasePatterns_ GenericSymbol, Colon_ TokenValue, OptionalSimpleStatement_ Statement) (Statement, error)

	// 141:2: statement -> default_branch_statement: ...
	DefaultBranchStatementToStatement(Default_ TokenValue, Colon_ TokenValue, OptionalSimpleStatement_ Statement) (Statement, error)
}

type OptionalSimpleStatementReducer interface {

	// 145:2: optional_simple_statement -> nil: ...
	NilToOptionalSimpleStatement() (Statement, error)
}

type ExpressionOrImproperStructStatementReducer interface {
	// 151:54: expression_or_improper_struct_statement -> ...
	ToExpressionOrImproperStructStatement(Expressions_ GenericSymbol) (Statement, error)
}

type ExpressionsReducer interface {
	// 154:2: expressions -> expression: ...
	ExpressionToExpressions(Expression_ Expression) (GenericSymbol, error)

	// 155:2: expressions -> add: ...
	AddToExpressions(Expressions_ GenericSymbol, Comma_ TokenValue, Expression_ Expression) (GenericSymbol, error)
}

type CallbackOpStatementReducer interface {
	// 181:36: callback_op_statement -> ...
	ToCallbackOpStatement(CallbackOp_ TokenValue, CallExpr_ Expression) (Statement, error)
}

type UnsafeStatementReducer interface {
	// 189:31: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ TokenValue, Less_ TokenValue, Identifier_ TokenValue, Greater_ TokenValue, StringLiteral_ TokenValue) (Statement, error)
}

type JumpStatementReducer interface {
	// 196:2: jump_statement -> unlabeled_no_value: ...
	UnlabeledNoValueToJumpStatement(JumpType_ TokenValue) (Statement, error)

	// 197:2: jump_statement -> unlabeled_valued: ...
	UnlabeledValuedToJumpStatement(JumpType_ TokenValue, Expressions_ GenericSymbol) (Statement, error)

	// 198:2: jump_statement -> labeled_no_value: ...
	LabeledNoValueToJumpStatement(JumpType_ TokenValue, JumpLabel_ TokenValue) (Statement, error)

	// 199:2: jump_statement -> labeled_valued: ...
	LabeledValuedToJumpStatement(JumpType_ TokenValue, JumpLabel_ TokenValue, Expressions_ GenericSymbol) (Statement, error)
}

type FallthroughStatementReducer interface {
	// 206:36: fallthrough_statement -> ...
	ToFallthroughStatement(Fallthrough_ TokenValue) (Statement, error)
}

type AssignStatementReducer interface {
	// 212:31: assign_statement -> ...
	ToAssignStatement(AssignPattern_ GenericSymbol, Assign_ TokenValue, Expression_ Expression) (Statement, error)
}

type UnaryOpAssignStatementReducer interface {
	// 214:40: unary_op_assign_statement -> ...
	ToUnaryOpAssignStatement(AccessibleExpr_ Expression, UnaryOpAssign_ TokenValue) (Statement, error)
}

type BinaryOpAssignStatementReducer interface {
	// 220:41: binary_op_assign_statement -> ...
	ToBinaryOpAssignStatement(AccessibleExpr_ Expression, BinaryOpAssign_ TokenValue, Expression_ Expression) (Statement, error)
}

type ImportStatementReducer interface {
	// 240:2: import_statement -> single: ...
	SingleToImportStatement(Import_ TokenValue, ImportClause_ GenericSymbol) (Statement, error)

	// 241:2: import_statement -> multiple: ...
	MultipleToImportStatement(Import_ TokenValue, Lparen_ TokenValue, ImportClauses_ GenericSymbol, Rparen_ TokenValue) (Statement, error)
}

type ImportClauseReducer interface {
	// 244:2: import_clause -> STRING_LITERAL: ...
	StringLiteralToImportClause(StringLiteral_ TokenValue) (GenericSymbol, error)

	// 245:2: import_clause -> alias: ...
	AliasToImportClause(StringLiteral_ TokenValue, As_ TokenValue, Identifier_ TokenValue) (GenericSymbol, error)
}

type ProperImportClausesReducer interface {
	// 248:2: proper_import_clauses -> add_implicit: ...
	AddImplicitToProperImportClauses(ProperImportClauses_ GenericSymbol, Newlines_ TokenCount, ImportClause_ GenericSymbol) (GenericSymbol, error)

	// 249:2: proper_import_clauses -> add_explicit: ...
	AddExplicitToProperImportClauses(ProperImportClauses_ GenericSymbol, Comma_ TokenValue, ImportClause_ GenericSymbol) (GenericSymbol, error)

	// 250:2: proper_import_clauses -> import_clause: ...
	ImportClauseToProperImportClauses(ImportClause_ GenericSymbol) (GenericSymbol, error)
}

type ImportClausesReducer interface {

	// 254:2: import_clauses -> implicit: ...
	ImplicitToImportClauses(ProperImportClauses_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)

	// 255:2: import_clauses -> explicit: ...
	ExplicitToImportClauses(ProperImportClauses_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)
}

type CasePatternsReducer interface {
	// 262:2: case_patterns -> case_pattern: ...
	CasePatternToCasePatterns(CasePattern_ GenericSymbol) (GenericSymbol, error)

	// 263:2: case_patterns -> multiple: ...
	MultipleToCasePatterns(CasePatterns_ GenericSymbol, Comma_ TokenValue, CasePattern_ GenericSymbol) (GenericSymbol, error)
}

type VarDeclPatternReducer interface {
	// 272:20: var_decl_pattern -> ...
	ToVarDeclPattern(VarOrLet_ TokenValue, VarPattern_ GenericSymbol, OptionalValueType_ TypeExpression) (GenericSymbol, error)
}

type VarPatternReducer interface {
	// 279:2: var_pattern -> IDENTIFIER: ...
	IdentifierToVarPattern(Identifier_ TokenValue) (GenericSymbol, error)

	// 280:2: var_pattern -> tuple_pattern: ...
	TuplePatternToVarPattern(TuplePattern_ GenericSymbol) (GenericSymbol, error)
}

type TuplePatternReducer interface {
	// 282:17: tuple_pattern -> ...
	ToTuplePattern(Lparen_ TokenValue, FieldVarPatterns_ GenericSymbol, Rparen_ TokenValue) (GenericSymbol, error)
}

type FieldVarPatternsReducer interface {
	// 285:2: field_var_patterns -> field_var_pattern: ...
	FieldVarPatternToFieldVarPatterns(FieldVarPattern_ GenericSymbol) (GenericSymbol, error)

	// 286:2: field_var_patterns -> add: ...
	AddToFieldVarPatterns(FieldVarPatterns_ GenericSymbol, Comma_ TokenValue, FieldVarPattern_ GenericSymbol) (GenericSymbol, error)
}

type FieldVarPatternReducer interface {
	// 289:2: field_var_pattern -> positional: ...
	PositionalToFieldVarPattern(VarPattern_ GenericSymbol) (GenericSymbol, error)

	// 290:2: field_var_pattern -> named: ...
	NamedToFieldVarPattern(Identifier_ TokenValue, Assign_ TokenValue, VarPattern_ GenericSymbol) (GenericSymbol, error)

	// 291:2: field_var_pattern -> ELLIPSIS: ...
	EllipsisToFieldVarPattern(Ellipsis_ TokenValue) (GenericSymbol, error)
}

type OptionalValueTypeReducer interface {

	// 295:2: optional_value_type -> nil: ...
	NilToOptionalValueType() (TypeExpression, error)
}

type AssignPatternReducer interface {
	// 302:2: assign_pattern -> ...
	ToAssignPattern(SequenceExpr_ Expression) (GenericSymbol, error)
}

type CasePatternReducer interface {

	// 312:2: case_pattern -> enum_match_pattern: ...
	EnumMatchPatternToCasePattern(Dot_ TokenValue, Identifier_ TokenValue, ImplicitStructExpr_ Expression) (GenericSymbol, error)

	// 313:2: case_pattern -> enum_var_decl_pattern: ...
	EnumVarDeclPatternToCasePattern(Var_ TokenValue, Dot_ TokenValue, Identifier_ TokenValue, TuplePattern_ GenericSymbol) (GenericSymbol, error)
}

type ExpressionReducer interface {
	// 320:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ TokenValue, IfExpr_ Expression) (Expression, error)

	// 321:2: expression -> switch_expr: ...
	SwitchExprToExpression(OptionalLabelDecl_ TokenValue, SwitchExpr_ Expression) (Expression, error)

	// 322:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ TokenValue, LoopExpr_ Expression) (Expression, error)
}

type OptionalLabelDeclReducer interface {
	// 326:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ TokenValue) (TokenValue, error)

	// 327:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (TokenValue, error)
}

type SequenceExprReducer interface {

	// 337:2: sequence_expr -> var_decl_pattern: ...
	VarDeclPatternToSequenceExpr(VarDeclPattern_ GenericSymbol) (Expression, error)

	// 341:2: sequence_expr -> assign_var_pattern: ...
	AssignVarPatternToSequenceExpr(Greater_ TokenValue, SequenceExpr_ Expression) (Expression, error)
}

type IfExprReducer interface {
	// 350:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol) (Expression, error)

	// 351:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol, Else_ TokenValue, StatementBlock_2 GenericSymbol) (Expression, error)

	// 352:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol, Else_ TokenValue, IfExpr_ Expression) (Expression, error)
}

type ConditionReducer interface {
	// 355:2: condition -> sequence_expr: ...
	SequenceExprToCondition(SequenceExpr_ Expression) (GenericSymbol, error)

	// 356:2: condition -> case: ...
	CaseToCondition(Case_ TokenValue, CasePatterns_ GenericSymbol, Assign_ TokenValue, SequenceExpr_ Expression) (GenericSymbol, error)
}

type SwitchExprReducer interface {
	// 380:27: switch_expr -> ...
	ToSwitchExpr(Switch_ TokenValue, SequenceExpr_ Expression, StatementBlock_ GenericSymbol) (Expression, error)
}

type LoopExprReducer interface {
	// 394:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 395:2: loop_expr -> do_while: ...
	DoWhileToLoopExpr(Do_ TokenValue, StatementBlock_ GenericSymbol, For_ TokenValue, SequenceExpr_ Expression) (Expression, error)

	// 396:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ TokenValue, SequenceExpr_ Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 397:2: loop_expr -> iterator: ...
	IteratorToLoopExpr(For_ TokenValue, AssignPattern_ GenericSymbol, In_ TokenValue, SequenceExpr_ Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 398:2: loop_expr -> for: ...
	ForToLoopExpr(For_ TokenValue, ForAssignment_ GenericSymbol, Semicolon_ TokenValue, OptionalSequenceExpr_ Expression, Semicolon_2 TokenValue, OptionalSequenceExpr_2 Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)
}

type OptionalSequenceExprReducer interface {
	// 401:2: optional_sequence_expr -> sequence_expr: ...
	SequenceExprToOptionalSequenceExpr(SequenceExpr_ Expression) (Expression, error)

	// 402:2: optional_sequence_expr -> nil: ...
	NilToOptionalSequenceExpr() (Expression, error)
}

type ForAssignmentReducer interface {
	// 405:2: for_assignment -> sequence_expr: ...
	SequenceExprToForAssignment(SequenceExpr_ Expression) (GenericSymbol, error)

	// 406:2: for_assignment -> assign: ...
	AssignToForAssignment(AssignPattern_ GenericSymbol, Assign_ TokenValue, SequenceExpr_ Expression) (GenericSymbol, error)
}

type CallExprReducer interface {
	// 413:2: call_expr -> ...
	ToCallExpr(AccessibleExpr_ Expression, OptionalGenericBinding_ GenericSymbol, Lparen_ TokenValue, Arguments_ *ArgumentList, Rparen_ TokenValue) (Expression, error)
}

type OptionalGenericBindingReducer interface {
	// 416:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ TokenValue, GenericArguments_ GenericSymbol, Rbracket_ TokenValue) (GenericSymbol, error)

	// 417:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (GenericSymbol, error)
}

type ProperGenericArgumentsReducer interface {
	// 420:2: proper_generic_arguments -> add: ...
	AddToProperGenericArguments(ProperGenericArguments_ GenericSymbol, Comma_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 421:2: proper_generic_arguments -> value_type: ...
	ValueTypeToProperGenericArguments(ValueType_ TypeExpression) (GenericSymbol, error)
}

type GenericArgumentsReducer interface {

	// 425:2: generic_arguments -> improper: ...
	ImproperToGenericArguments(ProperGenericArguments_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 426:2: generic_arguments -> nil: ...
	NilToGenericArguments() (GenericSymbol, error)
}

type ProperArgumentsReducer interface {
	// 429:2: proper_arguments -> add: ...
	AddToProperArguments(ProperArguments_ *ArgumentList, Comma_ TokenValue, Argument_ *Argument) (*ArgumentList, error)

	// 430:2: proper_arguments -> argument: ...
	ArgumentToProperArguments(Argument_ *Argument) (*ArgumentList, error)
}

type ArgumentsReducer interface {

	// 434:2: arguments -> improper: ...
	ImproperToArguments(ProperArguments_ *ArgumentList, Comma_ TokenValue) (*ArgumentList, error)

	// 435:2: arguments -> nil: ...
	NilToArguments() (*ArgumentList, error)
}

type ArgumentReducer interface {
	// 438:2: argument -> positional: ...
	PositionalToArgument(Expression_ Expression) (*Argument, error)

	// 439:2: argument -> colon_expr: ...
	ColonExprToArgument(ColonExpr_ Expression) (*Argument, error)

	// 440:2: argument -> named_assignment: ...
	NamedAssignmentToArgument(Identifier_ TokenValue, Assign_ TokenValue, Expression_ Expression) (*Argument, error)

	// 444:2: argument -> vararg_assignment: ...
	VarargAssignmentToArgument(Expression_ Expression, Ellipsis_ TokenValue) (*Argument, error)

	// 447:2: argument -> skip_pattern: ...
	SkipPatternToArgument(Ellipsis_ TokenValue) (*Argument, error)
}

type ColonExprReducer interface {
	// 451:2: colon_expr -> unit_unit_pair: ...
	UnitUnitPairToColonExpr(Colon_ TokenValue) (Expression, error)

	// 452:2: colon_expr -> expr_unit_pair: ...
	ExprUnitPairToColonExpr(Expression_ Expression, Colon_ TokenValue) (Expression, error)

	// 453:2: colon_expr -> unit_expr_pair: ...
	UnitExprPairToColonExpr(Colon_ TokenValue, Expression_ Expression) (Expression, error)

	// 454:2: colon_expr -> expr_expr_pair: ...
	ExprExprPairToColonExpr(Expression_ Expression, Colon_ TokenValue, Expression_2 Expression) (Expression, error)

	// 455:2: colon_expr -> colon_expr_unit_tuple: ...
	ColonExprUnitTupleToColonExpr(ColonExpr_ Expression, Colon_ TokenValue) (Expression, error)

	// 456:2: colon_expr -> colon_expr_expr_tuple: ...
	ColonExprExprTupleToColonExpr(ColonExpr_ Expression, Colon_ TokenValue, Expression_ Expression) (Expression, error)
}

type ParseErrorExprReducer interface {
	// 475:32: parse_error_expr -> ...
	ToParseErrorExpr(ParseError_ ParseErrorSymbol) (Expression, error)
}

type LiteralExprReducer interface {
	// 478:2: literal_expr -> TRUE: ...
	TrueToLiteralExpr(True_ TokenValue) (Expression, error)

	// 479:2: literal_expr -> FALSE: ...
	FalseToLiteralExpr(False_ TokenValue) (Expression, error)

	// 480:2: literal_expr -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteralExpr(IntegerLiteral_ TokenValue) (Expression, error)

	// 481:2: literal_expr -> FLOAT_LITERAL: ...
	FloatLiteralToLiteralExpr(FloatLiteral_ TokenValue) (Expression, error)

	// 482:2: literal_expr -> RUNE_LITERAL: ...
	RuneLiteralToLiteralExpr(RuneLiteral_ TokenValue) (Expression, error)

	// 483:2: literal_expr -> STRING_LITERAL: ...
	StringLiteralToLiteralExpr(StringLiteral_ TokenValue) (Expression, error)
}

type IdentifierExprReducer interface {
	// 485:32: identifier_expr -> ...
	ToIdentifierExpr(Identifier_ TokenValue) (Expression, error)
}

type BlockExprReducer interface {
	// 487:27: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)
}

type InitializeExprReducer interface {
	// 489:31: initialize_expr -> ...
	ToInitializeExpr(InitializableType_ TypeExpression, Lparen_ TokenValue, Arguments_ *ArgumentList, Rparen_ TokenValue) (Expression, error)
}

type ImplicitStructExprReducer interface {
	// 491:36: implicit_struct_expr -> ...
	ToImplicitStructExpr(Lparen_ TokenValue, Arguments_ *ArgumentList, Rparen_ TokenValue) (Expression, error)
}

type AccessExprReducer interface {
	// 499:27: access_expr -> ...
	ToAccessExpr(AccessibleExpr_ Expression, Dot_ TokenValue, Identifier_ TokenValue) (Expression, error)
}

type IndexExprReducer interface {
	// 503:26: index_expr -> ...
	ToIndexExpr(AccessibleExpr_ Expression, Lbracket_ TokenValue, Argument_ *Argument, Rbracket_ TokenValue) (Expression, error)
}

type PostfixUnaryExprReducer interface {
	// 509:35: postfix_unary_expr -> ...
	ToPostfixUnaryExpr(AccessibleExpr_ Expression, Question_ TokenValue) (Expression, error)
}

type PrefixUnaryExprReducer interface {
	// 515:33: prefix_unary_expr -> ...
	ToPrefixUnaryExpr(PrefixUnaryOp_ TokenValue, PrefixableExpr_ Expression) (Expression, error)
}

type BinaryMulExprReducer interface {
	// 532:31: binary_mul_expr -> ...
	ToBinaryMulExpr(MulExpr_ Expression, MulOp_ TokenValue, PrefixableExpr_ Expression) (Expression, error)
}

type BinaryAddExprReducer interface {
	// 546:31: binary_add_expr -> ...
	ToBinaryAddExpr(AddExpr_ Expression, AddOp_ TokenValue, MulExpr_ Expression) (Expression, error)
}

type BinaryCmpExprReducer interface {
	// 558:31: binary_cmp_expr -> ...
	ToBinaryCmpExpr(CmpExpr_ Expression, CmpOp_ TokenValue, AddExpr_ Expression) (Expression, error)
}

type BinaryAndExprReducer interface {
	// 572:31: binary_and_expr -> ...
	ToBinaryAndExpr(AndExpr_ Expression, And_ TokenValue, CmpExpr_ Expression) (Expression, error)
}

type BinaryOrExprReducer interface {
	// 578:30: binary_or_expr -> ...
	ToBinaryOrExpr(OrExpr_ Expression, Or_ TokenValue, AndExpr_ Expression) (Expression, error)
}

type SliceTypeReducer interface {
	// 593:30: slice_type -> ...
	ToSliceType(Lbracket_ TokenValue, ValueType_ TypeExpression, Rbracket_ TokenValue) (TypeExpression, error)
}

type ArrayTypeReducer interface {
	// 595:30: array_type -> ...
	ToArrayType(Lbracket_ TokenValue, ValueType_ TypeExpression, Comma_ TokenValue, IntegerLiteral_ TokenValue, Rbracket_ TokenValue) (TypeExpression, error)
}

type MapTypeReducer interface {
	// 598:28: map_type -> ...
	ToMapType(Lbracket_ TokenValue, ValueType_ TypeExpression, Colon_ TokenValue, ValueType_2 TypeExpression, Rbracket_ TokenValue) (TypeExpression, error)
}

type AtomTypeReducer interface {

	// 602:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ TokenValue, OptionalGenericBinding_ GenericSymbol) (TypeExpression, error)

	// 603:2: atom_type -> extern_named: ...
	ExternNamedToAtomType(Identifier_ TokenValue, Dot_ TokenValue, Identifier_2 TokenValue, OptionalGenericBinding_ GenericSymbol) (TypeExpression, error)

	// 604:2: atom_type -> inferred: ...
	InferredToAtomType(Dot_ TokenValue, OptionalGenericBinding_ GenericSymbol) (TypeExpression, error)
}

type ParseErrorTypeReducer interface {
	// 612:36: parse_error_type -> ...
	ToParseErrorType(ParseError_ ParseErrorSymbol) (TypeExpression, error)
}

type PrefixedTypeReducer interface {
	// 621:33: prefixed_type -> ...
	ToPrefixedType(PrefixTypeOp_ TokenValue, ReturnableType_ TypeExpression) (TypeExpression, error)
}

type TraitOpTypeReducer interface {
	// 636:33: trait_op_type -> ...
	ToTraitOpType(ValueType_ TypeExpression, TraitOp_ TokenValue, ReturnableType_ TypeExpression) (TypeExpression, error)
}

type TypeDefReducer interface {
	// 644:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, ValueType_ TypeExpression) (SourceDefinition, error)

	// 645:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, ValueType_ TypeExpression, Implements_ TokenValue, ValueType_2 TypeExpression) (SourceDefinition, error)

	// 646:2: type_def -> alias: ...
	AliasToTypeDef(Type_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, ValueType_ TypeExpression) (SourceDefinition, error)
}

type GenericParameterDefReducer interface {
	// 654:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ TokenValue) (GenericSymbol, error)

	// 655:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)
}

type ProperGenericParameterDefsReducer interface {
	// 658:2: proper_generic_parameter_defs -> add: ...
	AddToProperGenericParameterDefs(ProperGenericParameterDefs_ GenericSymbol, Comma_ TokenValue, GenericParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 659:2: proper_generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToProperGenericParameterDefs(GenericParameterDef_ GenericSymbol) (GenericSymbol, error)
}

type GenericParameterDefsReducer interface {

	// 663:2: generic_parameter_defs -> improper: ...
	ImproperToGenericParameterDefs(ProperGenericParameterDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 664:2: generic_parameter_defs -> nil: ...
	NilToGenericParameterDefs() (GenericSymbol, error)
}

type OptionalGenericParametersReducer interface {
	// 667:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ TokenValue, GenericParameterDefs_ GenericSymbol, Rbracket_ TokenValue) (GenericSymbol, error)

	// 668:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (GenericSymbol, error)
}

type FieldDefReducer interface {
	// 675:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 676:2: field_def -> implicit: ...
	ImplicitToFieldDef(ValueType_ TypeExpression) (GenericSymbol, error)

	// 677:2: field_def -> unsafe_statement: ...
	UnsafeStatementToFieldDef(UnsafeStatement_ Statement) (GenericSymbol, error)
}

type ProperImplicitFieldDefsReducer interface {
	// 680:2: proper_implicit_field_defs -> add: ...
	AddToProperImplicitFieldDefs(ProperImplicitFieldDefs_ GenericSymbol, Comma_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 681:2: proper_implicit_field_defs -> field_def: ...
	FieldDefToProperImplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type ImplicitFieldDefsReducer interface {

	// 685:2: implicit_field_defs -> improper: ...
	ImproperToImplicitFieldDefs(ProperImplicitFieldDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 686:2: implicit_field_defs -> nil: ...
	NilToImplicitFieldDefs() (GenericSymbol, error)
}

type ImplicitStructDefReducer interface {
	// 689:2: implicit_struct_def -> ...
	ToImplicitStructDef(Lparen_ TokenValue, ImplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ProperExplicitFieldDefsReducer interface {
	// 692:2: proper_explicit_field_defs -> add_implicit: ...
	AddImplicitToProperExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Newlines_ TokenCount, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 693:2: proper_explicit_field_defs -> add_explicit: ...
	AddExplicitToProperExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Comma_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 694:2: proper_explicit_field_defs -> field_def: ...
	FieldDefToProperExplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type ExplicitFieldDefsReducer interface {

	// 698:2: explicit_field_defs -> improper_implicit: ...
	ImproperImplicitToExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)

	// 699:2: explicit_field_defs -> improper_explicit: ...
	ImproperExplicitToExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 700:2: explicit_field_defs -> nil: ...
	NilToExplicitFieldDefs() (GenericSymbol, error)
}

type ExplicitStructDefReducer interface {
	// 703:2: explicit_struct_def -> ...
	ToExplicitStructDef(Struct_ TokenValue, Lparen_ TokenValue, ExplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type EnumValueDefReducer interface {
	// 711:2: enum_value_def -> field_def: ...
	FieldDefToEnumValueDef(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 712:2: enum_value_def -> default: ...
	DefaultToEnumValueDef(FieldDef_ GenericSymbol, Assign_ TokenValue, Default_ TokenValue) (GenericSymbol, error)
}

type ImplicitEnumValueDefsReducer interface {
	// 724:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Or_ TokenValue, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 725:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Or_ TokenValue, EnumValueDef_ GenericSymbol) (GenericSymbol, error)
}

type ImplicitEnumDefReducer interface {
	// 727:37: implicit_enum_def -> ...
	ToImplicitEnumDef(Lparen_ TokenValue, ImplicitEnumValueDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ExplicitEnumValueDefsReducer interface {
	// 730:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Or_ TokenValue, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 731:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Newlines_ TokenCount, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 732:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Or_ TokenValue, EnumValueDef_ GenericSymbol) (GenericSymbol, error)

	// 733:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Newlines_ TokenCount, EnumValueDef_ GenericSymbol) (GenericSymbol, error)
}

type ExplicitEnumDefReducer interface {
	// 735:37: explicit_enum_def -> ...
	ToExplicitEnumDef(Enum_ TokenValue, Lparen_ TokenValue, ExplicitEnumValueDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type TraitPropertyReducer interface {
	// 742:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 743:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ GenericSymbol) (GenericSymbol, error)
}

type ProperTraitPropertiesReducer interface {
	// 746:2: proper_trait_properties -> implicit: ...
	ImplicitToProperTraitProperties(ProperTraitProperties_ GenericSymbol, Newlines_ TokenCount, TraitProperty_ GenericSymbol) (GenericSymbol, error)

	// 747:2: proper_trait_properties -> explicit: ...
	ExplicitToProperTraitProperties(ProperTraitProperties_ GenericSymbol, Comma_ TokenValue, TraitProperty_ GenericSymbol) (GenericSymbol, error)

	// 748:2: proper_trait_properties -> trait_property: ...
	TraitPropertyToProperTraitProperties(TraitProperty_ GenericSymbol) (GenericSymbol, error)
}

type TraitPropertiesReducer interface {

	// 752:2: trait_properties -> improper_implicit: ...
	ImproperImplicitToTraitProperties(ProperTraitProperties_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)

	// 753:2: trait_properties -> improper_explicit: ...
	ImproperExplicitToTraitProperties(ProperTraitProperties_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 754:2: trait_properties -> nil: ...
	NilToTraitProperties() (GenericSymbol, error)
}

type TraitDefReducer interface {
	// 756:29: trait_def -> ...
	ToTraitDef(Trait_ TokenValue, Lparen_ TokenValue, TraitProperties_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ReturnTypeReducer interface {
	// 764:2: return_type -> returnable_type: ...
	ReturnableTypeToReturnType(ReturnableType_ TypeExpression) (TypeExpression, error)

	// 765:2: return_type -> nil: ...
	NilToReturnType() (TypeExpression, error)
}

type ParameterDeclReducer interface {
	// 768:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 769:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ TokenValue, Ellipsis_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 770:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ TypeExpression) (GenericSymbol, error)

	// 771:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(Ellipsis_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)
}

type ProperParameterDeclsReducer interface {
	// 774:2: proper_parameter_decls -> add: ...
	AddToProperParameterDecls(ProperParameterDecls_ GenericSymbol, Comma_ TokenValue, ParameterDecl_ GenericSymbol) (GenericSymbol, error)

	// 775:2: proper_parameter_decls -> parameter_decl: ...
	ParameterDeclToProperParameterDecls(ParameterDecl_ GenericSymbol) (GenericSymbol, error)
}

type ParameterDeclsReducer interface {

	// 779:2: parameter_decls -> improper: ...
	ImproperToParameterDecls(ProperParameterDecls_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 780:2: parameter_decls -> nil: ...
	NilToParameterDecls() (GenericSymbol, error)
}

type FuncTypeReducer interface {
	// 783:2: func_type -> ...
	ToFuncType(Func_ TokenValue, Lparen_ TokenValue, ParameterDecls_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression) (TypeExpression, error)
}

type MethodSignatureReducer interface {
	// 794:20: method_signature -> ...
	ToMethodSignature(Func_ TokenValue, Identifier_ TokenValue, Lparen_ TokenValue, ParameterDecls_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression) (GenericSymbol, error)
}

type ParameterDefReducer interface {
	// 800:2: parameter_def -> inferred_ref_arg: ...
	InferredRefArgToParameterDef(Identifier_ TokenValue) (GenericSymbol, error)

	// 801:2: parameter_def -> inferred_ref_vararg: ...
	InferredRefVarargToParameterDef(Identifier_ TokenValue, Ellipsis_ TokenValue) (GenericSymbol, error)

	// 802:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 803:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ TokenValue, Ellipsis_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)
}

type ProperParameterDefsReducer interface {
	// 806:2: proper_parameter_defs -> add: ...
	AddToProperParameterDefs(ProperParameterDefs_ GenericSymbol, Comma_ TokenValue, ParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 807:2: proper_parameter_defs -> parameter_def: ...
	ParameterDefToProperParameterDefs(ParameterDef_ GenericSymbol) (GenericSymbol, error)
}

type ParameterDefsReducer interface {

	// 811:2: parameter_defs -> improper: ...
	ImproperToParameterDefs(ProperParameterDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 812:2: parameter_defs -> nil: ...
	NilToParameterDefs() (GenericSymbol, error)
}

type NamedFuncDefReducer interface {
	// 815:2: named_func_def -> func_def: ...
	FuncDefToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, Lparen_ TokenValue, ParameterDefs_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 816:2: named_func_def -> method_def: ...
	MethodDefToNamedFuncDef(Func_ TokenValue, Lparen_ TokenValue, ParameterDef_ GenericSymbol, Rparen_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, Lparen_2 TokenValue, ParameterDefs_ GenericSymbol, Rparen_2 TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 817:2: named_func_def -> alias: ...
	AliasToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, Expression_ Expression) (SourceDefinition, error)
}

type AnonymousFuncExprReducer interface {
	// 821:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ TokenValue, Lparen_ TokenValue, ParameterDefs_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (Expression, error)
}

type PackageDefReducer interface {
	// 833:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ TokenValue) (SourceDefinition, error)

	// 834:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ TokenValue, StatementBlock_ GenericSymbol) (SourceDefinition, error)
}

type Reducer interface {
	SourceReducer
	DefinitionsReducer
	DefinitionReducer
	StatementBlockReducer
	ProperStatementsReducer
	StatementsReducer
	StatementReducer
	OptionalSimpleStatementReducer
	ExpressionOrImproperStructStatementReducer
	ExpressionsReducer
	CallbackOpStatementReducer
	UnsafeStatementReducer
	JumpStatementReducer
	FallthroughStatementReducer
	AssignStatementReducer
	UnaryOpAssignStatementReducer
	BinaryOpAssignStatementReducer
	ImportStatementReducer
	ImportClauseReducer
	ProperImportClausesReducer
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
	ProperGenericArgumentsReducer
	GenericArgumentsReducer
	ProperArgumentsReducer
	ArgumentsReducer
	ArgumentReducer
	ColonExprReducer
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
	SliceTypeReducer
	ArrayTypeReducer
	MapTypeReducer
	AtomTypeReducer
	ParseErrorTypeReducer
	PrefixedTypeReducer
	TraitOpTypeReducer
	TypeDefReducer
	GenericParameterDefReducer
	ProperGenericParameterDefsReducer
	GenericParameterDefsReducer
	OptionalGenericParametersReducer
	FieldDefReducer
	ProperImplicitFieldDefsReducer
	ImplicitFieldDefsReducer
	ImplicitStructDefReducer
	ProperExplicitFieldDefsReducer
	ExplicitFieldDefsReducer
	ExplicitStructDefReducer
	EnumValueDefReducer
	ImplicitEnumValueDefsReducer
	ImplicitEnumDefReducer
	ExplicitEnumValueDefsReducer
	ExplicitEnumDefReducer
	TraitPropertyReducer
	ProperTraitPropertiesReducer
	TraitPropertiesReducer
	TraitDefReducer
	ReturnTypeReducer
	ParameterDeclReducer
	ProperParameterDeclsReducer
	ParameterDeclsReducer
	FuncTypeReducer
	MethodSignatureReducer
	ParameterDefReducer
	ProperParameterDefsReducer
	ParameterDefsReducer
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
	case EllipsisToken:
		return "ELLIPSIS"
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
	case DefinitionsType:
		return "definitions"
	case DefinitionType:
		return "definition"
	case StatementBlockType:
		return "statement_block"
	case ProperStatementsType:
		return "proper_statements"
	case StatementsType:
		return "statements"
	case SimpleStatementType:
		return "simple_statement"
	case StatementType:
		return "statement"
	case OptionalSimpleStatementType:
		return "optional_simple_statement"
	case ExpressionOrImproperStructStatementType:
		return "expression_or_improper_struct_statement"
	case ExpressionsType:
		return "expressions"
	case CallbackOpType:
		return "callback_op"
	case CallbackOpStatementType:
		return "callback_op_statement"
	case UnsafeStatementType:
		return "unsafe_statement"
	case JumpStatementType:
		return "jump_statement"
	case JumpTypeType:
		return "jump_type"
	case FallthroughStatementType:
		return "fallthrough_statement"
	case AssignStatementType:
		return "assign_statement"
	case UnaryOpAssignStatementType:
		return "unary_op_assign_statement"
	case UnaryOpAssignType:
		return "unary_op_assign"
	case BinaryOpAssignStatementType:
		return "binary_op_assign_statement"
	case BinaryOpAssignType:
		return "binary_op_assign"
	case ImportStatementType:
		return "import_statement"
	case ImportClauseType:
		return "import_clause"
	case ProperImportClausesType:
		return "proper_import_clauses"
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
	case ProperGenericArgumentsType:
		return "proper_generic_arguments"
	case GenericArgumentsType:
		return "generic_arguments"
	case ProperArgumentsType:
		return "proper_arguments"
	case ArgumentsType:
		return "arguments"
	case ArgumentType:
		return "argument"
	case ColonExprType:
		return "colon_expr"
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
	case SliceTypeType:
		return "slice_type"
	case ArrayTypeType:
		return "array_type"
	case MapTypeType:
		return "map_type"
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
	case ProperGenericParameterDefsType:
		return "proper_generic_parameter_defs"
	case GenericParameterDefsType:
		return "generic_parameter_defs"
	case OptionalGenericParametersType:
		return "optional_generic_parameters"
	case FieldDefType:
		return "field_def"
	case ProperImplicitFieldDefsType:
		return "proper_implicit_field_defs"
	case ImplicitFieldDefsType:
		return "implicit_field_defs"
	case ImplicitStructDefType:
		return "implicit_struct_def"
	case ProperExplicitFieldDefsType:
		return "proper_explicit_field_defs"
	case ExplicitFieldDefsType:
		return "explicit_field_defs"
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
	case ProperTraitPropertiesType:
		return "proper_trait_properties"
	case TraitPropertiesType:
		return "trait_properties"
	case TraitDefType:
		return "trait_def"
	case ReturnTypeType:
		return "return_type"
	case ParameterDeclType:
		return "parameter_decl"
	case ProperParameterDeclsType:
		return "proper_parameter_decls"
	case ParameterDeclsType:
		return "parameter_decls"
	case FuncTypeType:
		return "func_type"
	case MethodSignatureType:
		return "method_signature"
	case ParameterDefType:
		return "parameter_def"
	case ProperParameterDefsType:
		return "proper_parameter_defs"
	case ParameterDefsType:
		return "parameter_defs"
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
	DefinitionsType                         = SymbolId(344)
	DefinitionType                          = SymbolId(345)
	StatementBlockType                      = SymbolId(346)
	ProperStatementsType                    = SymbolId(347)
	StatementsType                          = SymbolId(348)
	SimpleStatementType                     = SymbolId(349)
	StatementType                           = SymbolId(350)
	OptionalSimpleStatementType             = SymbolId(351)
	ExpressionOrImproperStructStatementType = SymbolId(352)
	ExpressionsType                         = SymbolId(353)
	CallbackOpType                          = SymbolId(354)
	CallbackOpStatementType                 = SymbolId(355)
	UnsafeStatementType                     = SymbolId(356)
	JumpStatementType                       = SymbolId(357)
	JumpTypeType                            = SymbolId(358)
	FallthroughStatementType                = SymbolId(359)
	AssignStatementType                     = SymbolId(360)
	UnaryOpAssignStatementType              = SymbolId(361)
	UnaryOpAssignType                       = SymbolId(362)
	BinaryOpAssignStatementType             = SymbolId(363)
	BinaryOpAssignType                      = SymbolId(364)
	ImportStatementType                     = SymbolId(365)
	ImportClauseType                        = SymbolId(366)
	ProperImportClausesType                 = SymbolId(367)
	ImportClausesType                       = SymbolId(368)
	CasePatternsType                        = SymbolId(369)
	VarDeclPatternType                      = SymbolId(370)
	VarOrLetType                            = SymbolId(371)
	VarPatternType                          = SymbolId(372)
	TuplePatternType                        = SymbolId(373)
	FieldVarPatternsType                    = SymbolId(374)
	FieldVarPatternType                     = SymbolId(375)
	OptionalValueTypeType                   = SymbolId(376)
	AssignPatternType                       = SymbolId(377)
	CasePatternType                         = SymbolId(378)
	ExpressionType                          = SymbolId(379)
	OptionalLabelDeclType                   = SymbolId(380)
	SequenceExprType                        = SymbolId(381)
	IfExprType                              = SymbolId(382)
	ConditionType                           = SymbolId(383)
	SwitchExprType                          = SymbolId(384)
	LoopExprType                            = SymbolId(385)
	OptionalSequenceExprType                = SymbolId(386)
	ForAssignmentType                       = SymbolId(387)
	CallExprType                            = SymbolId(388)
	OptionalGenericBindingType              = SymbolId(389)
	ProperGenericArgumentsType              = SymbolId(390)
	GenericArgumentsType                    = SymbolId(391)
	ProperArgumentsType                     = SymbolId(392)
	ArgumentsType                           = SymbolId(393)
	ArgumentType                            = SymbolId(394)
	ColonExprType                           = SymbolId(395)
	AtomExprType                            = SymbolId(396)
	ParseErrorExprType                      = SymbolId(397)
	LiteralExprType                         = SymbolId(398)
	IdentifierExprType                      = SymbolId(399)
	BlockExprType                           = SymbolId(400)
	InitializeExprType                      = SymbolId(401)
	ImplicitStructExprType                  = SymbolId(402)
	AccessibleExprType                      = SymbolId(403)
	AccessExprType                          = SymbolId(404)
	IndexExprType                           = SymbolId(405)
	PostfixableExprType                     = SymbolId(406)
	PostfixUnaryExprType                    = SymbolId(407)
	PrefixableExprType                      = SymbolId(408)
	PrefixUnaryExprType                     = SymbolId(409)
	PrefixUnaryOpType                       = SymbolId(410)
	MulExprType                             = SymbolId(411)
	BinaryMulExprType                       = SymbolId(412)
	MulOpType                               = SymbolId(413)
	AddExprType                             = SymbolId(414)
	BinaryAddExprType                       = SymbolId(415)
	AddOpType                               = SymbolId(416)
	CmpExprType                             = SymbolId(417)
	BinaryCmpExprType                       = SymbolId(418)
	CmpOpType                               = SymbolId(419)
	AndExprType                             = SymbolId(420)
	BinaryAndExprType                       = SymbolId(421)
	OrExprType                              = SymbolId(422)
	BinaryOrExprType                        = SymbolId(423)
	InitializableTypeType                   = SymbolId(424)
	SliceTypeType                           = SymbolId(425)
	ArrayTypeType                           = SymbolId(426)
	MapTypeType                             = SymbolId(427)
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
	ProperGenericParameterDefsType          = SymbolId(438)
	GenericParameterDefsType                = SymbolId(439)
	OptionalGenericParametersType           = SymbolId(440)
	FieldDefType                            = SymbolId(441)
	ProperImplicitFieldDefsType             = SymbolId(442)
	ImplicitFieldDefsType                   = SymbolId(443)
	ImplicitStructDefType                   = SymbolId(444)
	ProperExplicitFieldDefsType             = SymbolId(445)
	ExplicitFieldDefsType                   = SymbolId(446)
	ExplicitStructDefType                   = SymbolId(447)
	EnumValueDefType                        = SymbolId(448)
	ImplicitEnumValueDefsType               = SymbolId(449)
	ImplicitEnumDefType                     = SymbolId(450)
	ExplicitEnumValueDefsType               = SymbolId(451)
	ExplicitEnumDefType                     = SymbolId(452)
	TraitPropertyType                       = SymbolId(453)
	ProperTraitPropertiesType               = SymbolId(454)
	TraitPropertiesType                     = SymbolId(455)
	TraitDefType                            = SymbolId(456)
	ReturnTypeType                          = SymbolId(457)
	ParameterDeclType                       = SymbolId(458)
	ProperParameterDeclsType                = SymbolId(459)
	ParameterDeclsType                      = SymbolId(460)
	FuncTypeType                            = SymbolId(461)
	MethodSignatureType                     = SymbolId(462)
	ParameterDefType                        = SymbolId(463)
	ProperParameterDefsType                 = SymbolId(464)
	ParameterDefsType                       = SymbolId(465)
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
	_ReduceToSource                                             = _ReduceType(1)
	_ReduceAddToDefinitions                                     = _ReduceType(2)
	_ReduceNilToDefinitions                                     = _ReduceType(3)
	_ReducePackageDefToDefinition                               = _ReduceType(4)
	_ReduceTypeDefToDefinition                                  = _ReduceType(5)
	_ReduceNamedFuncDefToDefinition                             = _ReduceType(6)
	_ReduceGlobalVarDefToDefinition                             = _ReduceType(7)
	_ReduceGlobalVarAssignmentToDefinition                      = _ReduceType(8)
	_ReduceStatementBlockToDefinition                           = _ReduceType(9)
	_ReduceCommentGroupsToDefinition                            = _ReduceType(10)
	_ReduceToStatementBlock                                     = _ReduceType(11)
	_ReduceAddImplicitToProperStatements                        = _ReduceType(12)
	_ReduceAddExplicitToProperStatements                        = _ReduceType(13)
	_ReduceStatementToProperStatements                          = _ReduceType(14)
	_ReduceProperStatementsToStatements                         = _ReduceType(15)
	_ReduceImproperImplicitToStatements                         = _ReduceType(16)
	_ReduceImproperExplicitToStatements                         = _ReduceType(17)
	_ReduceNilToStatements                                      = _ReduceType(18)
	_ReduceUnsafeStatementToSimpleStatement                     = _ReduceType(19)
	_ReduceExpressionOrImproperStructStatementToSimpleStatement = _ReduceType(20)
	_ReduceCallbackOpStatementToSimpleStatement                 = _ReduceType(21)
	_ReduceJumpStatementToSimpleStatement                       = _ReduceType(22)
	_ReduceAssignStatementToSimpleStatement                     = _ReduceType(23)
	_ReduceUnaryOpAssignStatementToSimpleStatement              = _ReduceType(24)
	_ReduceBinaryOpAssignStatementToSimpleStatement             = _ReduceType(25)
	_ReduceFallthroughStatementToSimpleStatement                = _ReduceType(26)
	_ReduceSimpleStatementToStatement                           = _ReduceType(27)
	_ReduceImportStatementToStatement                           = _ReduceType(28)
	_ReduceCaseBranchStatementToStatement                       = _ReduceType(29)
	_ReduceDefaultBranchStatementToStatement                    = _ReduceType(30)
	_ReduceSimpleStatementToOptionalSimpleStatement             = _ReduceType(31)
	_ReduceNilToOptionalSimpleStatement                         = _ReduceType(32)
	_ReduceToExpressionOrImproperStructStatement                = _ReduceType(33)
	_ReduceExpressionToExpressions                              = _ReduceType(34)
	_ReduceAddToExpressions                                     = _ReduceType(35)
	_ReduceAsyncToCallbackOp                                    = _ReduceType(36)
	_ReduceDeferToCallbackOp                                    = _ReduceType(37)
	_ReduceToCallbackOpStatement                                = _ReduceType(38)
	_ReduceToUnsafeStatement                                    = _ReduceType(39)
	_ReduceUnlabeledNoValueToJumpStatement                      = _ReduceType(40)
	_ReduceUnlabeledValuedToJumpStatement                       = _ReduceType(41)
	_ReduceLabeledNoValueToJumpStatement                        = _ReduceType(42)
	_ReduceLabeledValuedToJumpStatement                         = _ReduceType(43)
	_ReduceReturnToJumpType                                     = _ReduceType(44)
	_ReduceBreakToJumpType                                      = _ReduceType(45)
	_ReduceContinueToJumpType                                   = _ReduceType(46)
	_ReduceToFallthroughStatement                               = _ReduceType(47)
	_ReduceToAssignStatement                                    = _ReduceType(48)
	_ReduceToUnaryOpAssignStatement                             = _ReduceType(49)
	_ReduceAddOneAssignToUnaryOpAssign                          = _ReduceType(50)
	_ReduceSubOneAssignToUnaryOpAssign                          = _ReduceType(51)
	_ReduceToBinaryOpAssignStatement                            = _ReduceType(52)
	_ReduceAddAssignToBinaryOpAssign                            = _ReduceType(53)
	_ReduceSubAssignToBinaryOpAssign                            = _ReduceType(54)
	_ReduceMulAssignToBinaryOpAssign                            = _ReduceType(55)
	_ReduceDivAssignToBinaryOpAssign                            = _ReduceType(56)
	_ReduceModAssignToBinaryOpAssign                            = _ReduceType(57)
	_ReduceBitNegAssignToBinaryOpAssign                         = _ReduceType(58)
	_ReduceBitAndAssignToBinaryOpAssign                         = _ReduceType(59)
	_ReduceBitOrAssignToBinaryOpAssign                          = _ReduceType(60)
	_ReduceBitXorAssignToBinaryOpAssign                         = _ReduceType(61)
	_ReduceBitLshiftAssignToBinaryOpAssign                      = _ReduceType(62)
	_ReduceBitRshiftAssignToBinaryOpAssign                      = _ReduceType(63)
	_ReduceSingleToImportStatement                              = _ReduceType(64)
	_ReduceMultipleToImportStatement                            = _ReduceType(65)
	_ReduceStringLiteralToImportClause                          = _ReduceType(66)
	_ReduceAliasToImportClause                                  = _ReduceType(67)
	_ReduceAddImplicitToProperImportClauses                     = _ReduceType(68)
	_ReduceAddExplicitToProperImportClauses                     = _ReduceType(69)
	_ReduceImportClauseToProperImportClauses                    = _ReduceType(70)
	_ReduceProperImportClausesToImportClauses                   = _ReduceType(71)
	_ReduceImplicitToImportClauses                              = _ReduceType(72)
	_ReduceExplicitToImportClauses                              = _ReduceType(73)
	_ReduceCasePatternToCasePatterns                            = _ReduceType(74)
	_ReduceMultipleToCasePatterns                               = _ReduceType(75)
	_ReduceToVarDeclPattern                                     = _ReduceType(76)
	_ReduceVarToVarOrLet                                        = _ReduceType(77)
	_ReduceLetToVarOrLet                                        = _ReduceType(78)
	_ReduceIdentifierToVarPattern                               = _ReduceType(79)
	_ReduceTuplePatternToVarPattern                             = _ReduceType(80)
	_ReduceToTuplePattern                                       = _ReduceType(81)
	_ReduceFieldVarPatternToFieldVarPatterns                    = _ReduceType(82)
	_ReduceAddToFieldVarPatterns                                = _ReduceType(83)
	_ReducePositionalToFieldVarPattern                          = _ReduceType(84)
	_ReduceNamedToFieldVarPattern                               = _ReduceType(85)
	_ReduceEllipsisToFieldVarPattern                            = _ReduceType(86)
	_ReduceValueTypeToOptionalValueType                         = _ReduceType(87)
	_ReduceNilToOptionalValueType                               = _ReduceType(88)
	_ReduceToAssignPattern                                      = _ReduceType(89)
	_ReduceAssignPatternToCasePattern                           = _ReduceType(90)
	_ReduceEnumMatchPatternToCasePattern                        = _ReduceType(91)
	_ReduceEnumVarDeclPatternToCasePattern                      = _ReduceType(92)
	_ReduceIfExprToExpression                                   = _ReduceType(93)
	_ReduceSwitchExprToExpression                               = _ReduceType(94)
	_ReduceLoopExprToExpression                                 = _ReduceType(95)
	_ReduceSequenceExprToExpression                             = _ReduceType(96)
	_ReduceLabelDeclToOptionalLabelDecl                         = _ReduceType(97)
	_ReduceUnlabelledToOptionalLabelDecl                        = _ReduceType(98)
	_ReduceOrExprToSequenceExpr                                 = _ReduceType(99)
	_ReduceVarDeclPatternToSequenceExpr                         = _ReduceType(100)
	_ReduceAssignVarPatternToSequenceExpr                       = _ReduceType(101)
	_ReduceNoElseToIfExpr                                       = _ReduceType(102)
	_ReduceIfElseToIfExpr                                       = _ReduceType(103)
	_ReduceMultiIfElseToIfExpr                                  = _ReduceType(104)
	_ReduceSequenceExprToCondition                              = _ReduceType(105)
	_ReduceCaseToCondition                                      = _ReduceType(106)
	_ReduceToSwitchExpr                                         = _ReduceType(107)
	_ReduceInfiniteToLoopExpr                                   = _ReduceType(108)
	_ReduceDoWhileToLoopExpr                                    = _ReduceType(109)
	_ReduceWhileToLoopExpr                                      = _ReduceType(110)
	_ReduceIteratorToLoopExpr                                   = _ReduceType(111)
	_ReduceForToLoopExpr                                        = _ReduceType(112)
	_ReduceSequenceExprToOptionalSequenceExpr                   = _ReduceType(113)
	_ReduceNilToOptionalSequenceExpr                            = _ReduceType(114)
	_ReduceSequenceExprToForAssignment                          = _ReduceType(115)
	_ReduceAssignToForAssignment                                = _ReduceType(116)
	_ReduceToCallExpr                                           = _ReduceType(117)
	_ReduceBindingToOptionalGenericBinding                      = _ReduceType(118)
	_ReduceNilToOptionalGenericBinding                          = _ReduceType(119)
	_ReduceAddToProperGenericArguments                          = _ReduceType(120)
	_ReduceValueTypeToProperGenericArguments                    = _ReduceType(121)
	_ReduceProperGenericArgumentsToGenericArguments             = _ReduceType(122)
	_ReduceImproperToGenericArguments                           = _ReduceType(123)
	_ReduceNilToGenericArguments                                = _ReduceType(124)
	_ReduceAddToProperArguments                                 = _ReduceType(125)
	_ReduceArgumentToProperArguments                            = _ReduceType(126)
	_ReduceProperArgumentsToArguments                           = _ReduceType(127)
	_ReduceImproperToArguments                                  = _ReduceType(128)
	_ReduceNilToArguments                                       = _ReduceType(129)
	_ReducePositionalToArgument                                 = _ReduceType(130)
	_ReduceColonExprToArgument                                  = _ReduceType(131)
	_ReduceNamedAssignmentToArgument                            = _ReduceType(132)
	_ReduceVarargAssignmentToArgument                           = _ReduceType(133)
	_ReduceSkipPatternToArgument                                = _ReduceType(134)
	_ReduceUnitUnitPairToColonExpr                              = _ReduceType(135)
	_ReduceExprUnitPairToColonExpr                              = _ReduceType(136)
	_ReduceUnitExprPairToColonExpr                              = _ReduceType(137)
	_ReduceExprExprPairToColonExpr                              = _ReduceType(138)
	_ReduceColonExprUnitTupleToColonExpr                        = _ReduceType(139)
	_ReduceColonExprExprTupleToColonExpr                        = _ReduceType(140)
	_ReduceParseErrorExprToAtomExpr                             = _ReduceType(141)
	_ReduceLiteralExprToAtomExpr                                = _ReduceType(142)
	_ReduceIdentifierExprToAtomExpr                             = _ReduceType(143)
	_ReduceBlockExprToAtomExpr                                  = _ReduceType(144)
	_ReduceAnonymousFuncExprToAtomExpr                          = _ReduceType(145)
	_ReduceInitializeExprToAtomExpr                             = _ReduceType(146)
	_ReduceImplicitStructExprToAtomExpr                         = _ReduceType(147)
	_ReduceToParseErrorExpr                                     = _ReduceType(148)
	_ReduceTrueToLiteralExpr                                    = _ReduceType(149)
	_ReduceFalseToLiteralExpr                                   = _ReduceType(150)
	_ReduceIntegerLiteralToLiteralExpr                          = _ReduceType(151)
	_ReduceFloatLiteralToLiteralExpr                            = _ReduceType(152)
	_ReduceRuneLiteralToLiteralExpr                             = _ReduceType(153)
	_ReduceStringLiteralToLiteralExpr                           = _ReduceType(154)
	_ReduceToIdentifierExpr                                     = _ReduceType(155)
	_ReduceToBlockExpr                                          = _ReduceType(156)
	_ReduceToInitializeExpr                                     = _ReduceType(157)
	_ReduceToImplicitStructExpr                                 = _ReduceType(158)
	_ReduceAtomExprToAccessibleExpr                             = _ReduceType(159)
	_ReduceAccessExprToAccessibleExpr                           = _ReduceType(160)
	_ReduceCallExprToAccessibleExpr                             = _ReduceType(161)
	_ReduceIndexExprToAccessibleExpr                            = _ReduceType(162)
	_ReduceToAccessExpr                                         = _ReduceType(163)
	_ReduceToIndexExpr                                          = _ReduceType(164)
	_ReduceAccessibleExprToPostfixableExpr                      = _ReduceType(165)
	_ReducePostfixUnaryExprToPostfixableExpr                    = _ReduceType(166)
	_ReduceToPostfixUnaryExpr                                   = _ReduceType(167)
	_ReducePostfixableExprToPrefixableExpr                      = _ReduceType(168)
	_ReducePrefixUnaryExprToPrefixableExpr                      = _ReduceType(169)
	_ReduceToPrefixUnaryExpr                                    = _ReduceType(170)
	_ReduceNotToPrefixUnaryOp                                   = _ReduceType(171)
	_ReduceBitNegToPrefixUnaryOp                                = _ReduceType(172)
	_ReduceSubToPrefixUnaryOp                                   = _ReduceType(173)
	_ReduceMulToPrefixUnaryOp                                   = _ReduceType(174)
	_ReduceBitAndToPrefixUnaryOp                                = _ReduceType(175)
	_ReducePrefixableExprToMulExpr                              = _ReduceType(176)
	_ReduceBinaryMulExprToMulExpr                               = _ReduceType(177)
	_ReduceToBinaryMulExpr                                      = _ReduceType(178)
	_ReduceMulToMulOp                                           = _ReduceType(179)
	_ReduceDivToMulOp                                           = _ReduceType(180)
	_ReduceModToMulOp                                           = _ReduceType(181)
	_ReduceBitAndToMulOp                                        = _ReduceType(182)
	_ReduceBitLshiftToMulOp                                     = _ReduceType(183)
	_ReduceBitRshiftToMulOp                                     = _ReduceType(184)
	_ReduceMulExprToAddExpr                                     = _ReduceType(185)
	_ReduceBinaryAddExprToAddExpr                               = _ReduceType(186)
	_ReduceToBinaryAddExpr                                      = _ReduceType(187)
	_ReduceAddToAddOp                                           = _ReduceType(188)
	_ReduceSubToAddOp                                           = _ReduceType(189)
	_ReduceBitOrToAddOp                                         = _ReduceType(190)
	_ReduceBitXorToAddOp                                        = _ReduceType(191)
	_ReduceAddExprToCmpExpr                                     = _ReduceType(192)
	_ReduceBinaryCmpExprToCmpExpr                               = _ReduceType(193)
	_ReduceToBinaryCmpExpr                                      = _ReduceType(194)
	_ReduceEqualToCmpOp                                         = _ReduceType(195)
	_ReduceNotEqualToCmpOp                                      = _ReduceType(196)
	_ReduceLessToCmpOp                                          = _ReduceType(197)
	_ReduceLessOrEqualToCmpOp                                   = _ReduceType(198)
	_ReduceGreaterToCmpOp                                       = _ReduceType(199)
	_ReduceGreaterOrEqualToCmpOp                                = _ReduceType(200)
	_ReduceCmpExprToAndExpr                                     = _ReduceType(201)
	_ReduceBinaryAndExprToAndExpr                               = _ReduceType(202)
	_ReduceToBinaryAndExpr                                      = _ReduceType(203)
	_ReduceAndExprToOrExpr                                      = _ReduceType(204)
	_ReduceBinaryOrExprToOrExpr                                 = _ReduceType(205)
	_ReduceToBinaryOrExpr                                       = _ReduceType(206)
	_ReduceExplicitStructDefToInitializableType                 = _ReduceType(207)
	_ReduceSliceTypeToInitializableType                         = _ReduceType(208)
	_ReduceArrayTypeToInitializableType                         = _ReduceType(209)
	_ReduceMapTypeToInitializableType                           = _ReduceType(210)
	_ReduceToSliceType                                          = _ReduceType(211)
	_ReduceToArrayType                                          = _ReduceType(212)
	_ReduceToMapType                                            = _ReduceType(213)
	_ReduceInitializableTypeToAtomType                          = _ReduceType(214)
	_ReduceNamedToAtomType                                      = _ReduceType(215)
	_ReduceExternNamedToAtomType                                = _ReduceType(216)
	_ReduceInferredToAtomType                                   = _ReduceType(217)
	_ReduceImplicitStructDefToAtomType                          = _ReduceType(218)
	_ReduceExplicitEnumDefToAtomType                            = _ReduceType(219)
	_ReduceImplicitEnumDefToAtomType                            = _ReduceType(220)
	_ReduceTraitDefToAtomType                                   = _ReduceType(221)
	_ReduceFuncTypeToAtomType                                   = _ReduceType(222)
	_ReduceParseErrorTypeToAtomType                             = _ReduceType(223)
	_ReduceToParseErrorType                                     = _ReduceType(224)
	_ReduceAtomTypeToReturnableType                             = _ReduceType(225)
	_ReducePrefixedTypeToReturnableType                         = _ReduceType(226)
	_ReduceToPrefixedType                                       = _ReduceType(227)
	_ReduceQuestionToPrefixTypeOp                               = _ReduceType(228)
	_ReduceExclaimToPrefixTypeOp                                = _ReduceType(229)
	_ReduceBitAndToPrefixTypeOp                                 = _ReduceType(230)
	_ReduceBitNegToPrefixTypeOp                                 = _ReduceType(231)
	_ReduceTildeTildeToPrefixTypeOp                             = _ReduceType(232)
	_ReduceReturnableTypeToValueType                            = _ReduceType(233)
	_ReduceTraitOpTypeToValueType                               = _ReduceType(234)
	_ReduceToTraitOpType                                        = _ReduceType(235)
	_ReduceMulToTraitOp                                         = _ReduceType(236)
	_ReduceAddToTraitOp                                         = _ReduceType(237)
	_ReduceSubToTraitOp                                         = _ReduceType(238)
	_ReduceDefinitionToTypeDef                                  = _ReduceType(239)
	_ReduceConstrainedDefToTypeDef                              = _ReduceType(240)
	_ReduceAliasToTypeDef                                       = _ReduceType(241)
	_ReduceUnconstrainedToGenericParameterDef                   = _ReduceType(242)
	_ReduceConstrainedToGenericParameterDef                     = _ReduceType(243)
	_ReduceAddToProperGenericParameterDefs                      = _ReduceType(244)
	_ReduceGenericParameterDefToProperGenericParameterDefs      = _ReduceType(245)
	_ReduceProperGenericParameterDefsToGenericParameterDefs     = _ReduceType(246)
	_ReduceImproperToGenericParameterDefs                       = _ReduceType(247)
	_ReduceNilToGenericParameterDefs                            = _ReduceType(248)
	_ReduceGenericToOptionalGenericParameters                   = _ReduceType(249)
	_ReduceNilToOptionalGenericParameters                       = _ReduceType(250)
	_ReduceExplicitToFieldDef                                   = _ReduceType(251)
	_ReduceImplicitToFieldDef                                   = _ReduceType(252)
	_ReduceUnsafeStatementToFieldDef                            = _ReduceType(253)
	_ReduceAddToProperImplicitFieldDefs                         = _ReduceType(254)
	_ReduceFieldDefToProperImplicitFieldDefs                    = _ReduceType(255)
	_ReduceProperImplicitFieldDefsToImplicitFieldDefs           = _ReduceType(256)
	_ReduceImproperToImplicitFieldDefs                          = _ReduceType(257)
	_ReduceNilToImplicitFieldDefs                               = _ReduceType(258)
	_ReduceToImplicitStructDef                                  = _ReduceType(259)
	_ReduceAddImplicitToProperExplicitFieldDefs                 = _ReduceType(260)
	_ReduceAddExplicitToProperExplicitFieldDefs                 = _ReduceType(261)
	_ReduceFieldDefToProperExplicitFieldDefs                    = _ReduceType(262)
	_ReduceProperExplicitFieldDefsToExplicitFieldDefs           = _ReduceType(263)
	_ReduceImproperImplicitToExplicitFieldDefs                  = _ReduceType(264)
	_ReduceImproperExplicitToExplicitFieldDefs                  = _ReduceType(265)
	_ReduceNilToExplicitFieldDefs                               = _ReduceType(266)
	_ReduceToExplicitStructDef                                  = _ReduceType(267)
	_ReduceFieldDefToEnumValueDef                               = _ReduceType(268)
	_ReduceDefaultToEnumValueDef                                = _ReduceType(269)
	_ReducePairToImplicitEnumValueDefs                          = _ReduceType(270)
	_ReduceAddToImplicitEnumValueDefs                           = _ReduceType(271)
	_ReduceToImplicitEnumDef                                    = _ReduceType(272)
	_ReduceExplicitPairToExplicitEnumValueDefs                  = _ReduceType(273)
	_ReduceImplicitPairToExplicitEnumValueDefs                  = _ReduceType(274)
	_ReduceExplicitAddToExplicitEnumValueDefs                   = _ReduceType(275)
	_ReduceImplicitAddToExplicitEnumValueDefs                   = _ReduceType(276)
	_ReduceToExplicitEnumDef                                    = _ReduceType(277)
	_ReduceFieldDefToTraitProperty                              = _ReduceType(278)
	_ReduceMethodSignatureToTraitProperty                       = _ReduceType(279)
	_ReduceImplicitToProperTraitProperties                      = _ReduceType(280)
	_ReduceExplicitToProperTraitProperties                      = _ReduceType(281)
	_ReduceTraitPropertyToProperTraitProperties                 = _ReduceType(282)
	_ReduceProperTraitPropertiesToTraitProperties               = _ReduceType(283)
	_ReduceImproperImplicitToTraitProperties                    = _ReduceType(284)
	_ReduceImproperExplicitToTraitProperties                    = _ReduceType(285)
	_ReduceNilToTraitProperties                                 = _ReduceType(286)
	_ReduceToTraitDef                                           = _ReduceType(287)
	_ReduceReturnableTypeToReturnType                           = _ReduceType(288)
	_ReduceNilToReturnType                                      = _ReduceType(289)
	_ReduceArgToParameterDecl                                   = _ReduceType(290)
	_ReduceVarargToParameterDecl                                = _ReduceType(291)
	_ReduceUnamedToParameterDecl                                = _ReduceType(292)
	_ReduceUnnamedVarargToParameterDecl                         = _ReduceType(293)
	_ReduceAddToProperParameterDecls                            = _ReduceType(294)
	_ReduceParameterDeclToProperParameterDecls                  = _ReduceType(295)
	_ReduceProperParameterDeclsToParameterDecls                 = _ReduceType(296)
	_ReduceImproperToParameterDecls                             = _ReduceType(297)
	_ReduceNilToParameterDecls                                  = _ReduceType(298)
	_ReduceToFuncType                                           = _ReduceType(299)
	_ReduceToMethodSignature                                    = _ReduceType(300)
	_ReduceInferredRefArgToParameterDef                         = _ReduceType(301)
	_ReduceInferredRefVarargToParameterDef                      = _ReduceType(302)
	_ReduceArgToParameterDef                                    = _ReduceType(303)
	_ReduceVarargToParameterDef                                 = _ReduceType(304)
	_ReduceAddToProperParameterDefs                             = _ReduceType(305)
	_ReduceParameterDefToProperParameterDefs                    = _ReduceType(306)
	_ReduceProperParameterDefsToParameterDefs                   = _ReduceType(307)
	_ReduceImproperToParameterDefs                              = _ReduceType(308)
	_ReduceNilToParameterDefs                                   = _ReduceType(309)
	_ReduceFuncDefToNamedFuncDef                                = _ReduceType(310)
	_ReduceMethodDefToNamedFuncDef                              = _ReduceType(311)
	_ReduceAliasToNamedFuncDef                                  = _ReduceType(312)
	_ReduceToAnonymousFuncExpr                                  = _ReduceType(313)
	_ReduceNoSpecToPackageDef                                   = _ReduceType(314)
	_ReduceWithSpecToPackageDef                                 = _ReduceType(315)
)

func (i _ReduceType) String() string {
	switch i {
	case _ReduceToSource:
		return "ToSource"
	case _ReduceAddToDefinitions:
		return "AddToDefinitions"
	case _ReduceNilToDefinitions:
		return "NilToDefinitions"
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
	case _ReduceAddImplicitToProperStatements:
		return "AddImplicitToProperStatements"
	case _ReduceAddExplicitToProperStatements:
		return "AddExplicitToProperStatements"
	case _ReduceStatementToProperStatements:
		return "StatementToProperStatements"
	case _ReduceProperStatementsToStatements:
		return "ProperStatementsToStatements"
	case _ReduceImproperImplicitToStatements:
		return "ImproperImplicitToStatements"
	case _ReduceImproperExplicitToStatements:
		return "ImproperExplicitToStatements"
	case _ReduceNilToStatements:
		return "NilToStatements"
	case _ReduceUnsafeStatementToSimpleStatement:
		return "UnsafeStatementToSimpleStatement"
	case _ReduceExpressionOrImproperStructStatementToSimpleStatement:
		return "ExpressionOrImproperStructStatementToSimpleStatement"
	case _ReduceCallbackOpStatementToSimpleStatement:
		return "CallbackOpStatementToSimpleStatement"
	case _ReduceJumpStatementToSimpleStatement:
		return "JumpStatementToSimpleStatement"
	case _ReduceAssignStatementToSimpleStatement:
		return "AssignStatementToSimpleStatement"
	case _ReduceUnaryOpAssignStatementToSimpleStatement:
		return "UnaryOpAssignStatementToSimpleStatement"
	case _ReduceBinaryOpAssignStatementToSimpleStatement:
		return "BinaryOpAssignStatementToSimpleStatement"
	case _ReduceFallthroughStatementToSimpleStatement:
		return "FallthroughStatementToSimpleStatement"
	case _ReduceSimpleStatementToStatement:
		return "SimpleStatementToStatement"
	case _ReduceImportStatementToStatement:
		return "ImportStatementToStatement"
	case _ReduceCaseBranchStatementToStatement:
		return "CaseBranchStatementToStatement"
	case _ReduceDefaultBranchStatementToStatement:
		return "DefaultBranchStatementToStatement"
	case _ReduceSimpleStatementToOptionalSimpleStatement:
		return "SimpleStatementToOptionalSimpleStatement"
	case _ReduceNilToOptionalSimpleStatement:
		return "NilToOptionalSimpleStatement"
	case _ReduceToExpressionOrImproperStructStatement:
		return "ToExpressionOrImproperStructStatement"
	case _ReduceExpressionToExpressions:
		return "ExpressionToExpressions"
	case _ReduceAddToExpressions:
		return "AddToExpressions"
	case _ReduceAsyncToCallbackOp:
		return "AsyncToCallbackOp"
	case _ReduceDeferToCallbackOp:
		return "DeferToCallbackOp"
	case _ReduceToCallbackOpStatement:
		return "ToCallbackOpStatement"
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
	case _ReduceToFallthroughStatement:
		return "ToFallthroughStatement"
	case _ReduceToAssignStatement:
		return "ToAssignStatement"
	case _ReduceToUnaryOpAssignStatement:
		return "ToUnaryOpAssignStatement"
	case _ReduceAddOneAssignToUnaryOpAssign:
		return "AddOneAssignToUnaryOpAssign"
	case _ReduceSubOneAssignToUnaryOpAssign:
		return "SubOneAssignToUnaryOpAssign"
	case _ReduceToBinaryOpAssignStatement:
		return "ToBinaryOpAssignStatement"
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
	case _ReduceSingleToImportStatement:
		return "SingleToImportStatement"
	case _ReduceMultipleToImportStatement:
		return "MultipleToImportStatement"
	case _ReduceStringLiteralToImportClause:
		return "StringLiteralToImportClause"
	case _ReduceAliasToImportClause:
		return "AliasToImportClause"
	case _ReduceAddImplicitToProperImportClauses:
		return "AddImplicitToProperImportClauses"
	case _ReduceAddExplicitToProperImportClauses:
		return "AddExplicitToProperImportClauses"
	case _ReduceImportClauseToProperImportClauses:
		return "ImportClauseToProperImportClauses"
	case _ReduceProperImportClausesToImportClauses:
		return "ProperImportClausesToImportClauses"
	case _ReduceImplicitToImportClauses:
		return "ImplicitToImportClauses"
	case _ReduceExplicitToImportClauses:
		return "ExplicitToImportClauses"
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
	case _ReduceEllipsisToFieldVarPattern:
		return "EllipsisToFieldVarPattern"
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
	case _ReduceAddToProperGenericArguments:
		return "AddToProperGenericArguments"
	case _ReduceValueTypeToProperGenericArguments:
		return "ValueTypeToProperGenericArguments"
	case _ReduceProperGenericArgumentsToGenericArguments:
		return "ProperGenericArgumentsToGenericArguments"
	case _ReduceImproperToGenericArguments:
		return "ImproperToGenericArguments"
	case _ReduceNilToGenericArguments:
		return "NilToGenericArguments"
	case _ReduceAddToProperArguments:
		return "AddToProperArguments"
	case _ReduceArgumentToProperArguments:
		return "ArgumentToProperArguments"
	case _ReduceProperArgumentsToArguments:
		return "ProperArgumentsToArguments"
	case _ReduceImproperToArguments:
		return "ImproperToArguments"
	case _ReduceNilToArguments:
		return "NilToArguments"
	case _ReducePositionalToArgument:
		return "PositionalToArgument"
	case _ReduceColonExprToArgument:
		return "ColonExprToArgument"
	case _ReduceNamedAssignmentToArgument:
		return "NamedAssignmentToArgument"
	case _ReduceVarargAssignmentToArgument:
		return "VarargAssignmentToArgument"
	case _ReduceSkipPatternToArgument:
		return "SkipPatternToArgument"
	case _ReduceUnitUnitPairToColonExpr:
		return "UnitUnitPairToColonExpr"
	case _ReduceExprUnitPairToColonExpr:
		return "ExprUnitPairToColonExpr"
	case _ReduceUnitExprPairToColonExpr:
		return "UnitExprPairToColonExpr"
	case _ReduceExprExprPairToColonExpr:
		return "ExprExprPairToColonExpr"
	case _ReduceColonExprUnitTupleToColonExpr:
		return "ColonExprUnitTupleToColonExpr"
	case _ReduceColonExprExprTupleToColonExpr:
		return "ColonExprExprTupleToColonExpr"
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
	case _ReduceSliceTypeToInitializableType:
		return "SliceTypeToInitializableType"
	case _ReduceArrayTypeToInitializableType:
		return "ArrayTypeToInitializableType"
	case _ReduceMapTypeToInitializableType:
		return "MapTypeToInitializableType"
	case _ReduceToSliceType:
		return "ToSliceType"
	case _ReduceToArrayType:
		return "ToArrayType"
	case _ReduceToMapType:
		return "ToMapType"
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
	case _ReduceAddToProperGenericParameterDefs:
		return "AddToProperGenericParameterDefs"
	case _ReduceGenericParameterDefToProperGenericParameterDefs:
		return "GenericParameterDefToProperGenericParameterDefs"
	case _ReduceProperGenericParameterDefsToGenericParameterDefs:
		return "ProperGenericParameterDefsToGenericParameterDefs"
	case _ReduceImproperToGenericParameterDefs:
		return "ImproperToGenericParameterDefs"
	case _ReduceNilToGenericParameterDefs:
		return "NilToGenericParameterDefs"
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
	case _ReduceAddToProperImplicitFieldDefs:
		return "AddToProperImplicitFieldDefs"
	case _ReduceFieldDefToProperImplicitFieldDefs:
		return "FieldDefToProperImplicitFieldDefs"
	case _ReduceProperImplicitFieldDefsToImplicitFieldDefs:
		return "ProperImplicitFieldDefsToImplicitFieldDefs"
	case _ReduceImproperToImplicitFieldDefs:
		return "ImproperToImplicitFieldDefs"
	case _ReduceNilToImplicitFieldDefs:
		return "NilToImplicitFieldDefs"
	case _ReduceToImplicitStructDef:
		return "ToImplicitStructDef"
	case _ReduceAddImplicitToProperExplicitFieldDefs:
		return "AddImplicitToProperExplicitFieldDefs"
	case _ReduceAddExplicitToProperExplicitFieldDefs:
		return "AddExplicitToProperExplicitFieldDefs"
	case _ReduceFieldDefToProperExplicitFieldDefs:
		return "FieldDefToProperExplicitFieldDefs"
	case _ReduceProperExplicitFieldDefsToExplicitFieldDefs:
		return "ProperExplicitFieldDefsToExplicitFieldDefs"
	case _ReduceImproperImplicitToExplicitFieldDefs:
		return "ImproperImplicitToExplicitFieldDefs"
	case _ReduceImproperExplicitToExplicitFieldDefs:
		return "ImproperExplicitToExplicitFieldDefs"
	case _ReduceNilToExplicitFieldDefs:
		return "NilToExplicitFieldDefs"
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
	case _ReduceImplicitToProperTraitProperties:
		return "ImplicitToProperTraitProperties"
	case _ReduceExplicitToProperTraitProperties:
		return "ExplicitToProperTraitProperties"
	case _ReduceTraitPropertyToProperTraitProperties:
		return "TraitPropertyToProperTraitProperties"
	case _ReduceProperTraitPropertiesToTraitProperties:
		return "ProperTraitPropertiesToTraitProperties"
	case _ReduceImproperImplicitToTraitProperties:
		return "ImproperImplicitToTraitProperties"
	case _ReduceImproperExplicitToTraitProperties:
		return "ImproperExplicitToTraitProperties"
	case _ReduceNilToTraitProperties:
		return "NilToTraitProperties"
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
	case _ReduceAddToProperParameterDecls:
		return "AddToProperParameterDecls"
	case _ReduceParameterDeclToProperParameterDecls:
		return "ParameterDeclToProperParameterDecls"
	case _ReduceProperParameterDeclsToParameterDecls:
		return "ProperParameterDeclsToParameterDecls"
	case _ReduceImproperToParameterDecls:
		return "ImproperToParameterDecls"
	case _ReduceNilToParameterDecls:
		return "NilToParameterDecls"
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
	case _ReduceAddToProperParameterDefs:
		return "AddToProperParameterDefs"
	case _ReduceParameterDefToProperParameterDefs:
		return "ParameterDefToProperParameterDefs"
	case _ReduceProperParameterDefsToParameterDefs:
		return "ProperParameterDefsToParameterDefs"
	case _ReduceImproperToParameterDefs:
		return "ImproperToParameterDefs"
	case _ReduceNilToParameterDefs:
		return "NilToParameterDefs"
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
	_State469 = _StateId(469)
	_State470 = _StateId(470)
)

type Symbol struct {
	SymbolId_ SymbolId

	Generic_ GenericSymbol

	Argument          *Argument
	ArgumentList      *ArgumentList
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
	case CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, AsToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, EllipsisToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken:
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
	case ArgumentType:
		loc, ok := interface{}(s.Argument).(locator)
		if ok {
			return loc.Loc()
		}
	case ProperArgumentsType, ArgumentsType:
		loc, ok := interface{}(s.ArgumentList).(locator)
		if ok {
			return loc.Loc()
		}
	case NewlinesToken:
		loc, ok := interface{}(s.Count).(locator)
		if ok {
			return loc.Loc()
		}
	case ExpressionType, SequenceExprType, IfExprType, SwitchExprType, LoopExprType, OptionalSequenceExprType, CallExprType, ColonExprType, AtomExprType, ParseErrorExprType, LiteralExprType, IdentifierExprType, BlockExprType, InitializeExprType, ImplicitStructExprType, AccessibleExprType, AccessExprType, IndexExprType, PostfixableExprType, PostfixUnaryExprType, PrefixableExprType, PrefixUnaryExprType, MulExprType, BinaryMulExprType, AddExprType, BinaryAddExprType, CmpExprType, BinaryCmpExprType, AndExprType, BinaryAndExprType, OrExprType, BinaryOrExprType, AnonymousFuncExprType:
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
	case SimpleStatementType, StatementType, OptionalSimpleStatementType, ExpressionOrImproperStructStatementType, CallbackOpStatementType, UnsafeStatementType, JumpStatementType, FallthroughStatementType, AssignStatementType, UnaryOpAssignStatementType, BinaryOpAssignStatementType, ImportStatementType:
		loc, ok := interface{}(s.Statement).(locator)
		if ok {
			return loc.Loc()
		}
	case ProperStatementsType, StatementsType:
		loc, ok := interface{}(s.Statements).(locator)
		if ok {
			return loc.Loc()
		}
	case OptionalValueTypeType, InitializableTypeType, SliceTypeType, ArrayTypeType, MapTypeType, AtomTypeType, ParseErrorTypeType, ReturnableTypeType, PrefixedTypeType, ValueTypeType, TraitOpTypeType, ImplicitStructDefType, ExplicitStructDefType, ImplicitEnumDefType, ExplicitEnumDefType, TraitDefType, ReturnTypeType, FuncTypeType:
		loc, ok := interface{}(s.TypeExpression).(locator)
		if ok {
			return loc.Loc()
		}
	case CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, AsToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, EllipsisToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, CallbackOpType, JumpTypeType, UnaryOpAssignType, BinaryOpAssignType, VarOrLetType, OptionalLabelDeclType, PrefixUnaryOpType, MulOpType, AddOpType, CmpOpType, PrefixTypeOpType, TraitOpType:
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
	case ArgumentType:
		loc, ok := interface{}(s.Argument).(locator)
		if ok {
			return loc.End()
		}
	case ProperArgumentsType, ArgumentsType:
		loc, ok := interface{}(s.ArgumentList).(locator)
		if ok {
			return loc.End()
		}
	case NewlinesToken:
		loc, ok := interface{}(s.Count).(locator)
		if ok {
			return loc.End()
		}
	case ExpressionType, SequenceExprType, IfExprType, SwitchExprType, LoopExprType, OptionalSequenceExprType, CallExprType, ColonExprType, AtomExprType, ParseErrorExprType, LiteralExprType, IdentifierExprType, BlockExprType, InitializeExprType, ImplicitStructExprType, AccessibleExprType, AccessExprType, IndexExprType, PostfixableExprType, PostfixUnaryExprType, PrefixableExprType, PrefixUnaryExprType, MulExprType, BinaryMulExprType, AddExprType, BinaryAddExprType, CmpExprType, BinaryCmpExprType, AndExprType, BinaryAndExprType, OrExprType, BinaryOrExprType, AnonymousFuncExprType:
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
	case SimpleStatementType, StatementType, OptionalSimpleStatementType, ExpressionOrImproperStructStatementType, CallbackOpStatementType, UnsafeStatementType, JumpStatementType, FallthroughStatementType, AssignStatementType, UnaryOpAssignStatementType, BinaryOpAssignStatementType, ImportStatementType:
		loc, ok := interface{}(s.Statement).(locator)
		if ok {
			return loc.End()
		}
	case ProperStatementsType, StatementsType:
		loc, ok := interface{}(s.Statements).(locator)
		if ok {
			return loc.End()
		}
	case OptionalValueTypeType, InitializableTypeType, SliceTypeType, ArrayTypeType, MapTypeType, AtomTypeType, ParseErrorTypeType, ReturnableTypeType, PrefixedTypeType, ValueTypeType, TraitOpTypeType, ImplicitStructDefType, ExplicitStructDefType, ImplicitEnumDefType, ExplicitEnumDefType, TraitDefType, ReturnTypeType, FuncTypeType:
		loc, ok := interface{}(s.TypeExpression).(locator)
		if ok {
			return loc.End()
		}
	case CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, AsToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, EllipsisToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, CallbackOpType, JumpTypeType, UnaryOpAssignType, BinaryOpAssignType, VarOrLetType, OptionalLabelDeclType, PrefixUnaryOpType, MulOpType, AddOpType, CmpOpType, PrefixTypeOpType, TraitOpType:
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
		symbol.SourceDefinitions, err = reducer.ToSource(args[0].SourceDefinitions)
	case _ReduceAddToDefinitions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = DefinitionsType
		symbol.SourceDefinitions, err = reducer.AddToDefinitions(args[0].SourceDefinitions, args[1].SourceDefinition, args[2].Count)
	case _ReduceNilToDefinitions:
		symbol.SymbolId_ = DefinitionsType
		symbol.SourceDefinitions, err = reducer.NilToDefinitions()
	case _ReducePackageDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		//line grammar.lr:60:4
		symbol.SourceDefinition = args[0].SourceDefinition
		err = nil
	case _ReduceTypeDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		//line grammar.lr:61:4
		symbol.SourceDefinition = args[0].SourceDefinition
		err = nil
	case _ReduceNamedFuncDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		//line grammar.lr:62:4
		symbol.SourceDefinition = args[0].SourceDefinition
		err = nil
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
	case _ReduceAddImplicitToProperStatements:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperStatementsType
		symbol.Statements, err = reducer.AddImplicitToProperStatements(args[0].Statements, args[1].Count, args[2].Statement)
	case _ReduceAddExplicitToProperStatements:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperStatementsType
		symbol.Statements, err = reducer.AddExplicitToProperStatements(args[0].Statements, args[1].Value, args[2].Statement)
	case _ReduceStatementToProperStatements:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperStatementsType
		symbol.Statements, err = reducer.StatementToProperStatements(args[0].Statement)
	case _ReduceProperStatementsToStatements:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementsType
		//line grammar.lr:99:4
		symbol.Statements = args[0].Statements
		err = nil
	case _ReduceImproperImplicitToStatements:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementsType
		symbol.Statements, err = reducer.ImproperImplicitToStatements(args[0].Statements, args[1].Count)
	case _ReduceImproperExplicitToStatements:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = StatementsType
		symbol.Statements, err = reducer.ImproperExplicitToStatements(args[0].Statements, args[1].Value)
	case _ReduceNilToStatements:
		symbol.SymbolId_ = StatementsType
		symbol.Statements, err = reducer.NilToStatements()
	case _ReduceUnsafeStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:105:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceExpressionOrImproperStructStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:107:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceCallbackOpStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:110:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceJumpStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:113:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceAssignStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:118:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceUnaryOpAssignStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:124:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceBinaryOpAssignStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:125:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceFallthroughStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:129:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceSimpleStatementToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:132:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceImportStatementToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:135:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceCaseBranchStatementToStatement:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = StatementType
		symbol.Statement, err = reducer.CaseBranchStatementToStatement(args[0].Value, args[1].Generic_, args[2].Value, args[3].Statement)
	case _ReduceDefaultBranchStatementToStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = StatementType
		symbol.Statement, err = reducer.DefaultBranchStatementToStatement(args[0].Value, args[1].Value, args[2].Statement)
	case _ReduceSimpleStatementToOptionalSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalSimpleStatementType
		//line grammar.lr:144:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceNilToOptionalSimpleStatement:
		symbol.SymbolId_ = OptionalSimpleStatementType
		symbol.Statement, err = reducer.NilToOptionalSimpleStatement()
	case _ReduceToExpressionOrImproperStructStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExpressionOrImproperStructStatementType
		symbol.Statement, err = reducer.ToExpressionOrImproperStructStatement(args[0].Generic_)
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
	case _ReduceAsyncToCallbackOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CallbackOpType
		//line grammar.lr:162:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDeferToCallbackOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CallbackOpType
		//line grammar.lr:163:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceToCallbackOpStatement:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CallbackOpStatementType
		symbol.Statement, err = reducer.ToCallbackOpStatement(args[0].Value, args[1].Expression)
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
		//line grammar.lr:202:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBreakToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		//line grammar.lr:203:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceContinueToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		//line grammar.lr:204:4
		symbol.Value = args[0].Value
		err = nil
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
	case _ReduceAddOneAssignToUnaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpAssignType
		//line grammar.lr:217:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubOneAssignToUnaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpAssignType
		//line grammar.lr:218:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceToBinaryOpAssignStatement:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BinaryOpAssignStatementType
		symbol.Statement, err = reducer.ToBinaryOpAssignStatement(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceAddAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:223:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:224:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:225:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDivAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:226:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceModAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:227:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:228:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:229:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitOrAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:230:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitXorAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:231:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitLshiftAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:232:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitRshiftAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:233:4
		symbol.Value = args[0].Value
		err = nil
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
	case _ReduceAddImplicitToProperImportClauses:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperImportClausesType
		symbol.Generic_, err = reducer.AddImplicitToProperImportClauses(args[0].Generic_, args[1].Count, args[2].Generic_)
	case _ReduceAddExplicitToProperImportClauses:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperImportClausesType
		symbol.Generic_, err = reducer.AddExplicitToProperImportClauses(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceImportClauseToProperImportClauses:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperImportClausesType
		symbol.Generic_, err = reducer.ImportClauseToProperImportClauses(args[0].Generic_)
	case _ReduceProperImportClausesToImportClauses:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ImportClausesType
		//line grammar.lr:253:4
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _ReduceImplicitToImportClauses:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClausesType
		symbol.Generic_, err = reducer.ImplicitToImportClauses(args[0].Generic_, args[1].Count)
	case _ReduceExplicitToImportClauses:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClausesType
		symbol.Generic_, err = reducer.ExplicitToImportClauses(args[0].Generic_, args[1].Value)
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
		//line grammar.lr:275:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLetToVarOrLet:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarOrLetType
		//line grammar.lr:276:4
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
	case _ReduceEllipsisToFieldVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldVarPatternType
		symbol.Generic_, err = reducer.EllipsisToFieldVarPattern(args[0].Value)
	case _ReduceValueTypeToOptionalValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalValueTypeType
		//line grammar.lr:294:4
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
		//line grammar.lr:305:4
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
		//line grammar.lr:323:4
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
		//line grammar.lr:334:4
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
		symbol.Expression, err = reducer.ToCallExpr(args[0].Expression, args[1].Generic_, args[2].Value, args[3].ArgumentList, args[4].Value)
	case _ReduceBindingToOptionalGenericBinding:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = OptionalGenericBindingType
		symbol.Generic_, err = reducer.BindingToOptionalGenericBinding(args[0].Value, args[1].Generic_, args[2].Value)
	case _ReduceNilToOptionalGenericBinding:
		symbol.SymbolId_ = OptionalGenericBindingType
		symbol.Generic_, err = reducer.NilToOptionalGenericBinding()
	case _ReduceAddToProperGenericArguments:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperGenericArgumentsType
		symbol.Generic_, err = reducer.AddToProperGenericArguments(args[0].Generic_, args[1].Value, args[2].TypeExpression)
	case _ReduceValueTypeToProperGenericArguments:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperGenericArgumentsType
		symbol.Generic_, err = reducer.ValueTypeToProperGenericArguments(args[0].TypeExpression)
	case _ReduceProperGenericArgumentsToGenericArguments:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericArgumentsType
		//line grammar.lr:424:4
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _ReduceImproperToGenericArguments:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = GenericArgumentsType
		symbol.Generic_, err = reducer.ImproperToGenericArguments(args[0].Generic_, args[1].Value)
	case _ReduceNilToGenericArguments:
		symbol.SymbolId_ = GenericArgumentsType
		symbol.Generic_, err = reducer.NilToGenericArguments()
	case _ReduceAddToProperArguments:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperArgumentsType
		symbol.ArgumentList, err = reducer.AddToProperArguments(args[0].ArgumentList, args[1].Value, args[2].Argument)
	case _ReduceArgumentToProperArguments:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperArgumentsType
		symbol.ArgumentList, err = reducer.ArgumentToProperArguments(args[0].Argument)
	case _ReduceProperArgumentsToArguments:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentsType
		//line grammar.lr:433:4
		symbol.ArgumentList = args[0].ArgumentList
		err = nil
	case _ReduceImproperToArguments:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ArgumentsType
		symbol.ArgumentList, err = reducer.ImproperToArguments(args[0].ArgumentList, args[1].Value)
	case _ReduceNilToArguments:
		symbol.SymbolId_ = ArgumentsType
		symbol.ArgumentList, err = reducer.NilToArguments()
	case _ReducePositionalToArgument:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentType
		symbol.Argument, err = reducer.PositionalToArgument(args[0].Expression)
	case _ReduceColonExprToArgument:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentType
		symbol.Argument, err = reducer.ColonExprToArgument(args[0].Expression)
	case _ReduceNamedAssignmentToArgument:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ArgumentType
		symbol.Argument, err = reducer.NamedAssignmentToArgument(args[0].Value, args[1].Value, args[2].Expression)
	case _ReduceVarargAssignmentToArgument:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ArgumentType
		symbol.Argument, err = reducer.VarargAssignmentToArgument(args[0].Expression, args[1].Value)
	case _ReduceSkipPatternToArgument:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ArgumentType
		symbol.Argument, err = reducer.SkipPatternToArgument(args[0].Value)
	case _ReduceUnitUnitPairToColonExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ColonExprType
		symbol.Expression, err = reducer.UnitUnitPairToColonExpr(args[0].Value)
	case _ReduceExprUnitPairToColonExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ColonExprType
		symbol.Expression, err = reducer.ExprUnitPairToColonExpr(args[0].Expression, args[1].Value)
	case _ReduceUnitExprPairToColonExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ColonExprType
		symbol.Expression, err = reducer.UnitExprPairToColonExpr(args[0].Value, args[1].Expression)
	case _ReduceExprExprPairToColonExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ColonExprType
		symbol.Expression, err = reducer.ExprExprPairToColonExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceColonExprUnitTupleToColonExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ColonExprType
		symbol.Expression, err = reducer.ColonExprUnitTupleToColonExpr(args[0].Expression, args[1].Value)
	case _ReduceColonExprExprTupleToColonExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ColonExprType
		symbol.Expression, err = reducer.ColonExprExprTupleToColonExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceParseErrorExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:467:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceLiteralExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:468:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceIdentifierExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:469:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBlockExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:470:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceAnonymousFuncExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:471:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceInitializeExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:472:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceImplicitStructExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:473:4
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
		symbol.Expression, err = reducer.ToInitializeExpr(args[0].TypeExpression, args[1].Value, args[2].ArgumentList, args[3].Value)
	case _ReduceToImplicitStructExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitStructExprType
		symbol.Expression, err = reducer.ToImplicitStructExpr(args[0].Value, args[1].ArgumentList, args[2].Value)
	case _ReduceAtomExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:494:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceAccessExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:495:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceCallExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:496:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceIndexExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:497:4
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
		symbol.Expression, err = reducer.ToIndexExpr(args[0].Expression, args[1].Value, args[2].Argument, args[3].Value)
	case _ReduceAccessibleExprToPostfixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixableExprType
		//line grammar.lr:506:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReducePostfixUnaryExprToPostfixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixableExprType
		//line grammar.lr:507:4
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
		//line grammar.lr:512:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReducePrefixUnaryExprToPrefixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixableExprType
		//line grammar.lr:513:4
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
		//line grammar.lr:518:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:519:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:520:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:523:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:526:4
		symbol.Value = args[0].Value
		err = nil
	case _ReducePrefixableExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		//line grammar.lr:529:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryMulExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		//line grammar.lr:530:4
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
		//line grammar.lr:535:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDivToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:536:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceModToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:537:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:538:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitLshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:539:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitRshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:540:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulExprToAddExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddExprType
		//line grammar.lr:543:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryAddExprToAddExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddExprType
		//line grammar.lr:544:4
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
		//line grammar.lr:549:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:550:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitOrToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:551:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitXorToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:552:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddExprToCmpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpExprType
		//line grammar.lr:555:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryCmpExprToCmpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpExprType
		//line grammar.lr:556:4
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
		//line grammar.lr:561:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceNotEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:562:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLessToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:563:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLessOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:564:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceGreaterToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:565:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceGreaterOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:566:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceCmpExprToAndExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AndExprType
		//line grammar.lr:569:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryAndExprToAndExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AndExprType
		//line grammar.lr:570:4
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
		//line grammar.lr:575:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryOrExprToOrExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OrExprType
		//line grammar.lr:576:4
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
		//line grammar.lr:587:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceSliceTypeToInitializableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeType
		//line grammar.lr:588:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceArrayTypeToInitializableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeType
		//line grammar.lr:589:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceMapTypeToInitializableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeType
		//line grammar.lr:590:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceToSliceType:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = SliceTypeType
		symbol.TypeExpression, err = reducer.ToSliceType(args[0].Value, args[1].TypeExpression, args[2].Value)
	case _ReduceToArrayType:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = ArrayTypeType
		symbol.TypeExpression, err = reducer.ToArrayType(args[0].Value, args[1].TypeExpression, args[2].Value, args[3].Value, args[4].Value)
	case _ReduceToMapType:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = MapTypeType
		symbol.TypeExpression, err = reducer.ToMapType(args[0].Value, args[1].TypeExpression, args[2].Value, args[3].TypeExpression, args[4].Value)
	case _ReduceInitializableTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:601:4
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
		//line grammar.lr:605:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceExplicitEnumDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:606:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceImplicitEnumDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:607:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceTraitDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:608:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceFuncTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:609:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceParseErrorTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:610:4
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
		//line grammar.lr:618:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReducePrefixedTypeToReturnableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableTypeType
		//line grammar.lr:619:4
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
		//line grammar.lr:624:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceExclaimToPrefixTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixTypeOpType
		//line grammar.lr:625:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToPrefixTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixTypeOpType
		//line grammar.lr:626:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegToPrefixTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixTypeOpType
		//line grammar.lr:627:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceTildeTildeToPrefixTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixTypeOpType
		//line grammar.lr:628:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceReturnableTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		//line grammar.lr:633:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceTraitOpTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		//line grammar.lr:634:4
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
		//line grammar.lr:639:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddToTraitOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitOpType
		//line grammar.lr:640:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToTraitOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitOpType
		//line grammar.lr:641:4
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
	case _ReduceAddToProperGenericParameterDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperGenericParameterDefsType
		symbol.Generic_, err = reducer.AddToProperGenericParameterDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceGenericParameterDefToProperGenericParameterDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperGenericParameterDefsType
		symbol.Generic_, err = reducer.GenericParameterDefToProperGenericParameterDefs(args[0].Generic_)
	case _ReduceProperGenericParameterDefsToGenericParameterDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericParameterDefsType
		//line grammar.lr:662:4
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _ReduceImproperToGenericParameterDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = GenericParameterDefsType
		symbol.Generic_, err = reducer.ImproperToGenericParameterDefs(args[0].Generic_, args[1].Value)
	case _ReduceNilToGenericParameterDefs:
		symbol.SymbolId_ = GenericParameterDefsType
		symbol.Generic_, err = reducer.NilToGenericParameterDefs()
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
	case _ReduceAddToProperImplicitFieldDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperImplicitFieldDefsType
		symbol.Generic_, err = reducer.AddToProperImplicitFieldDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceFieldDefToProperImplicitFieldDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperImplicitFieldDefsType
		symbol.Generic_, err = reducer.FieldDefToProperImplicitFieldDefs(args[0].Generic_)
	case _ReduceProperImplicitFieldDefsToImplicitFieldDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ImplicitFieldDefsType
		//line grammar.lr:684:4
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _ReduceImproperToImplicitFieldDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImplicitFieldDefsType
		symbol.Generic_, err = reducer.ImproperToImplicitFieldDefs(args[0].Generic_, args[1].Value)
	case _ReduceNilToImplicitFieldDefs:
		symbol.SymbolId_ = ImplicitFieldDefsType
		symbol.Generic_, err = reducer.NilToImplicitFieldDefs()
	case _ReduceToImplicitStructDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitStructDefType
		symbol.TypeExpression, err = reducer.ToImplicitStructDef(args[0].Value, args[1].Generic_, args[2].Value)
	case _ReduceAddImplicitToProperExplicitFieldDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperExplicitFieldDefsType
		symbol.Generic_, err = reducer.AddImplicitToProperExplicitFieldDefs(args[0].Generic_, args[1].Count, args[2].Generic_)
	case _ReduceAddExplicitToProperExplicitFieldDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperExplicitFieldDefsType
		symbol.Generic_, err = reducer.AddExplicitToProperExplicitFieldDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceFieldDefToProperExplicitFieldDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperExplicitFieldDefsType
		symbol.Generic_, err = reducer.FieldDefToProperExplicitFieldDefs(args[0].Generic_)
	case _ReduceProperExplicitFieldDefsToExplicitFieldDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExplicitFieldDefsType
		//line grammar.lr:697:4
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _ReduceImproperImplicitToExplicitFieldDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExplicitFieldDefsType
		symbol.Generic_, err = reducer.ImproperImplicitToExplicitFieldDefs(args[0].Generic_, args[1].Count)
	case _ReduceImproperExplicitToExplicitFieldDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExplicitFieldDefsType
		symbol.Generic_, err = reducer.ImproperExplicitToExplicitFieldDefs(args[0].Generic_, args[1].Value)
	case _ReduceNilToExplicitFieldDefs:
		symbol.SymbolId_ = ExplicitFieldDefsType
		symbol.Generic_, err = reducer.NilToExplicitFieldDefs()
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
	case _ReduceImplicitToProperTraitProperties:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperTraitPropertiesType
		symbol.Generic_, err = reducer.ImplicitToProperTraitProperties(args[0].Generic_, args[1].Count, args[2].Generic_)
	case _ReduceExplicitToProperTraitProperties:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperTraitPropertiesType
		symbol.Generic_, err = reducer.ExplicitToProperTraitProperties(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceTraitPropertyToProperTraitProperties:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperTraitPropertiesType
		symbol.Generic_, err = reducer.TraitPropertyToProperTraitProperties(args[0].Generic_)
	case _ReduceProperTraitPropertiesToTraitProperties:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitPropertiesType
		//line grammar.lr:751:4
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _ReduceImproperImplicitToTraitProperties:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = TraitPropertiesType
		symbol.Generic_, err = reducer.ImproperImplicitToTraitProperties(args[0].Generic_, args[1].Count)
	case _ReduceImproperExplicitToTraitProperties:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = TraitPropertiesType
		symbol.Generic_, err = reducer.ImproperExplicitToTraitProperties(args[0].Generic_, args[1].Value)
	case _ReduceNilToTraitProperties:
		symbol.SymbolId_ = TraitPropertiesType
		symbol.Generic_, err = reducer.NilToTraitProperties()
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
	case _ReduceAddToProperParameterDecls:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperParameterDeclsType
		symbol.Generic_, err = reducer.AddToProperParameterDecls(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceParameterDeclToProperParameterDecls:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperParameterDeclsType
		symbol.Generic_, err = reducer.ParameterDeclToProperParameterDecls(args[0].Generic_)
	case _ReduceProperParameterDeclsToParameterDecls:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDeclsType
		//line grammar.lr:778:4
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _ReduceImproperToParameterDecls:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDeclsType
		symbol.Generic_, err = reducer.ImproperToParameterDecls(args[0].Generic_, args[1].Value)
	case _ReduceNilToParameterDecls:
		symbol.SymbolId_ = ParameterDeclsType
		symbol.Generic_, err = reducer.NilToParameterDecls()
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
	case _ReduceAddToProperParameterDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperParameterDefsType
		symbol.Generic_, err = reducer.AddToProperParameterDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceParameterDefToProperParameterDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperParameterDefsType
		symbol.Generic_, err = reducer.ParameterDefToProperParameterDefs(args[0].Generic_)
	case _ReduceProperParameterDefsToParameterDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDefsType
		//line grammar.lr:810:4
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _ReduceImproperToParameterDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDefsType
		symbol.Generic_, err = reducer.ImproperToParameterDefs(args[0].Generic_, args[1].Value)
	case _ReduceNilToParameterDefs:
		symbol.SymbolId_ = ParameterDefsType
		symbol.Generic_, err = reducer.NilToParameterDefs()
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
	_GotoState1Action                                                 = &_Action{_ShiftAction, _State1, 0}
	_GotoState2Action                                                 = &_Action{_ShiftAction, _State2, 0}
	_GotoState3Action                                                 = &_Action{_ShiftAction, _State3, 0}
	_GotoState4Action                                                 = &_Action{_ShiftAction, _State4, 0}
	_GotoState5Action                                                 = &_Action{_ShiftAction, _State5, 0}
	_GotoState6Action                                                 = &_Action{_ShiftAction, _State6, 0}
	_GotoState7Action                                                 = &_Action{_ShiftAction, _State7, 0}
	_GotoState8Action                                                 = &_Action{_ShiftAction, _State8, 0}
	_GotoState9Action                                                 = &_Action{_ShiftAction, _State9, 0}
	_GotoState10Action                                                = &_Action{_ShiftAction, _State10, 0}
	_GotoState11Action                                                = &_Action{_ShiftAction, _State11, 0}
	_GotoState12Action                                                = &_Action{_ShiftAction, _State12, 0}
	_GotoState13Action                                                = &_Action{_ShiftAction, _State13, 0}
	_GotoState14Action                                                = &_Action{_ShiftAction, _State14, 0}
	_GotoState15Action                                                = &_Action{_ShiftAction, _State15, 0}
	_GotoState16Action                                                = &_Action{_ShiftAction, _State16, 0}
	_GotoState17Action                                                = &_Action{_ShiftAction, _State17, 0}
	_GotoState18Action                                                = &_Action{_ShiftAction, _State18, 0}
	_GotoState19Action                                                = &_Action{_ShiftAction, _State19, 0}
	_GotoState20Action                                                = &_Action{_ShiftAction, _State20, 0}
	_GotoState21Action                                                = &_Action{_ShiftAction, _State21, 0}
	_GotoState22Action                                                = &_Action{_ShiftAction, _State22, 0}
	_GotoState23Action                                                = &_Action{_ShiftAction, _State23, 0}
	_GotoState24Action                                                = &_Action{_ShiftAction, _State24, 0}
	_GotoState25Action                                                = &_Action{_ShiftAction, _State25, 0}
	_GotoState26Action                                                = &_Action{_ShiftAction, _State26, 0}
	_GotoState27Action                                                = &_Action{_ShiftAction, _State27, 0}
	_GotoState28Action                                                = &_Action{_ShiftAction, _State28, 0}
	_GotoState29Action                                                = &_Action{_ShiftAction, _State29, 0}
	_GotoState30Action                                                = &_Action{_ShiftAction, _State30, 0}
	_GotoState31Action                                                = &_Action{_ShiftAction, _State31, 0}
	_GotoState32Action                                                = &_Action{_ShiftAction, _State32, 0}
	_GotoState33Action                                                = &_Action{_ShiftAction, _State33, 0}
	_GotoState34Action                                                = &_Action{_ShiftAction, _State34, 0}
	_GotoState35Action                                                = &_Action{_ShiftAction, _State35, 0}
	_GotoState36Action                                                = &_Action{_ShiftAction, _State36, 0}
	_GotoState37Action                                                = &_Action{_ShiftAction, _State37, 0}
	_GotoState38Action                                                = &_Action{_ShiftAction, _State38, 0}
	_GotoState39Action                                                = &_Action{_ShiftAction, _State39, 0}
	_GotoState40Action                                                = &_Action{_ShiftAction, _State40, 0}
	_GotoState41Action                                                = &_Action{_ShiftAction, _State41, 0}
	_GotoState42Action                                                = &_Action{_ShiftAction, _State42, 0}
	_GotoState43Action                                                = &_Action{_ShiftAction, _State43, 0}
	_GotoState44Action                                                = &_Action{_ShiftAction, _State44, 0}
	_GotoState45Action                                                = &_Action{_ShiftAction, _State45, 0}
	_GotoState46Action                                                = &_Action{_ShiftAction, _State46, 0}
	_GotoState47Action                                                = &_Action{_ShiftAction, _State47, 0}
	_GotoState48Action                                                = &_Action{_ShiftAction, _State48, 0}
	_GotoState49Action                                                = &_Action{_ShiftAction, _State49, 0}
	_GotoState50Action                                                = &_Action{_ShiftAction, _State50, 0}
	_GotoState51Action                                                = &_Action{_ShiftAction, _State51, 0}
	_GotoState52Action                                                = &_Action{_ShiftAction, _State52, 0}
	_GotoState53Action                                                = &_Action{_ShiftAction, _State53, 0}
	_GotoState54Action                                                = &_Action{_ShiftAction, _State54, 0}
	_GotoState55Action                                                = &_Action{_ShiftAction, _State55, 0}
	_GotoState56Action                                                = &_Action{_ShiftAction, _State56, 0}
	_GotoState57Action                                                = &_Action{_ShiftAction, _State57, 0}
	_GotoState58Action                                                = &_Action{_ShiftAction, _State58, 0}
	_GotoState59Action                                                = &_Action{_ShiftAction, _State59, 0}
	_GotoState60Action                                                = &_Action{_ShiftAction, _State60, 0}
	_GotoState61Action                                                = &_Action{_ShiftAction, _State61, 0}
	_GotoState62Action                                                = &_Action{_ShiftAction, _State62, 0}
	_GotoState63Action                                                = &_Action{_ShiftAction, _State63, 0}
	_GotoState64Action                                                = &_Action{_ShiftAction, _State64, 0}
	_GotoState65Action                                                = &_Action{_ShiftAction, _State65, 0}
	_GotoState66Action                                                = &_Action{_ShiftAction, _State66, 0}
	_GotoState67Action                                                = &_Action{_ShiftAction, _State67, 0}
	_GotoState68Action                                                = &_Action{_ShiftAction, _State68, 0}
	_GotoState69Action                                                = &_Action{_ShiftAction, _State69, 0}
	_GotoState70Action                                                = &_Action{_ShiftAction, _State70, 0}
	_GotoState71Action                                                = &_Action{_ShiftAction, _State71, 0}
	_GotoState72Action                                                = &_Action{_ShiftAction, _State72, 0}
	_GotoState73Action                                                = &_Action{_ShiftAction, _State73, 0}
	_GotoState74Action                                                = &_Action{_ShiftAction, _State74, 0}
	_GotoState75Action                                                = &_Action{_ShiftAction, _State75, 0}
	_GotoState76Action                                                = &_Action{_ShiftAction, _State76, 0}
	_GotoState77Action                                                = &_Action{_ShiftAction, _State77, 0}
	_GotoState78Action                                                = &_Action{_ShiftAction, _State78, 0}
	_GotoState79Action                                                = &_Action{_ShiftAction, _State79, 0}
	_GotoState80Action                                                = &_Action{_ShiftAction, _State80, 0}
	_GotoState81Action                                                = &_Action{_ShiftAction, _State81, 0}
	_GotoState82Action                                                = &_Action{_ShiftAction, _State82, 0}
	_GotoState83Action                                                = &_Action{_ShiftAction, _State83, 0}
	_GotoState84Action                                                = &_Action{_ShiftAction, _State84, 0}
	_GotoState85Action                                                = &_Action{_ShiftAction, _State85, 0}
	_GotoState86Action                                                = &_Action{_ShiftAction, _State86, 0}
	_GotoState87Action                                                = &_Action{_ShiftAction, _State87, 0}
	_GotoState88Action                                                = &_Action{_ShiftAction, _State88, 0}
	_GotoState89Action                                                = &_Action{_ShiftAction, _State89, 0}
	_GotoState90Action                                                = &_Action{_ShiftAction, _State90, 0}
	_GotoState91Action                                                = &_Action{_ShiftAction, _State91, 0}
	_GotoState92Action                                                = &_Action{_ShiftAction, _State92, 0}
	_GotoState93Action                                                = &_Action{_ShiftAction, _State93, 0}
	_GotoState94Action                                                = &_Action{_ShiftAction, _State94, 0}
	_GotoState95Action                                                = &_Action{_ShiftAction, _State95, 0}
	_GotoState96Action                                                = &_Action{_ShiftAction, _State96, 0}
	_GotoState97Action                                                = &_Action{_ShiftAction, _State97, 0}
	_GotoState98Action                                                = &_Action{_ShiftAction, _State98, 0}
	_GotoState99Action                                                = &_Action{_ShiftAction, _State99, 0}
	_GotoState100Action                                               = &_Action{_ShiftAction, _State100, 0}
	_GotoState101Action                                               = &_Action{_ShiftAction, _State101, 0}
	_GotoState102Action                                               = &_Action{_ShiftAction, _State102, 0}
	_GotoState103Action                                               = &_Action{_ShiftAction, _State103, 0}
	_GotoState104Action                                               = &_Action{_ShiftAction, _State104, 0}
	_GotoState105Action                                               = &_Action{_ShiftAction, _State105, 0}
	_GotoState106Action                                               = &_Action{_ShiftAction, _State106, 0}
	_GotoState107Action                                               = &_Action{_ShiftAction, _State107, 0}
	_GotoState108Action                                               = &_Action{_ShiftAction, _State108, 0}
	_GotoState109Action                                               = &_Action{_ShiftAction, _State109, 0}
	_GotoState110Action                                               = &_Action{_ShiftAction, _State110, 0}
	_GotoState111Action                                               = &_Action{_ShiftAction, _State111, 0}
	_GotoState112Action                                               = &_Action{_ShiftAction, _State112, 0}
	_GotoState113Action                                               = &_Action{_ShiftAction, _State113, 0}
	_GotoState114Action                                               = &_Action{_ShiftAction, _State114, 0}
	_GotoState115Action                                               = &_Action{_ShiftAction, _State115, 0}
	_GotoState116Action                                               = &_Action{_ShiftAction, _State116, 0}
	_GotoState117Action                                               = &_Action{_ShiftAction, _State117, 0}
	_GotoState118Action                                               = &_Action{_ShiftAction, _State118, 0}
	_GotoState119Action                                               = &_Action{_ShiftAction, _State119, 0}
	_GotoState120Action                                               = &_Action{_ShiftAction, _State120, 0}
	_GotoState121Action                                               = &_Action{_ShiftAction, _State121, 0}
	_GotoState122Action                                               = &_Action{_ShiftAction, _State122, 0}
	_GotoState123Action                                               = &_Action{_ShiftAction, _State123, 0}
	_GotoState124Action                                               = &_Action{_ShiftAction, _State124, 0}
	_GotoState125Action                                               = &_Action{_ShiftAction, _State125, 0}
	_GotoState126Action                                               = &_Action{_ShiftAction, _State126, 0}
	_GotoState127Action                                               = &_Action{_ShiftAction, _State127, 0}
	_GotoState128Action                                               = &_Action{_ShiftAction, _State128, 0}
	_GotoState129Action                                               = &_Action{_ShiftAction, _State129, 0}
	_GotoState130Action                                               = &_Action{_ShiftAction, _State130, 0}
	_GotoState131Action                                               = &_Action{_ShiftAction, _State131, 0}
	_GotoState132Action                                               = &_Action{_ShiftAction, _State132, 0}
	_GotoState133Action                                               = &_Action{_ShiftAction, _State133, 0}
	_GotoState134Action                                               = &_Action{_ShiftAction, _State134, 0}
	_GotoState135Action                                               = &_Action{_ShiftAction, _State135, 0}
	_GotoState136Action                                               = &_Action{_ShiftAction, _State136, 0}
	_GotoState137Action                                               = &_Action{_ShiftAction, _State137, 0}
	_GotoState138Action                                               = &_Action{_ShiftAction, _State138, 0}
	_GotoState139Action                                               = &_Action{_ShiftAction, _State139, 0}
	_GotoState140Action                                               = &_Action{_ShiftAction, _State140, 0}
	_GotoState141Action                                               = &_Action{_ShiftAction, _State141, 0}
	_GotoState142Action                                               = &_Action{_ShiftAction, _State142, 0}
	_GotoState143Action                                               = &_Action{_ShiftAction, _State143, 0}
	_GotoState144Action                                               = &_Action{_ShiftAction, _State144, 0}
	_GotoState145Action                                               = &_Action{_ShiftAction, _State145, 0}
	_GotoState146Action                                               = &_Action{_ShiftAction, _State146, 0}
	_GotoState147Action                                               = &_Action{_ShiftAction, _State147, 0}
	_GotoState148Action                                               = &_Action{_ShiftAction, _State148, 0}
	_GotoState149Action                                               = &_Action{_ShiftAction, _State149, 0}
	_GotoState150Action                                               = &_Action{_ShiftAction, _State150, 0}
	_GotoState151Action                                               = &_Action{_ShiftAction, _State151, 0}
	_GotoState152Action                                               = &_Action{_ShiftAction, _State152, 0}
	_GotoState153Action                                               = &_Action{_ShiftAction, _State153, 0}
	_GotoState154Action                                               = &_Action{_ShiftAction, _State154, 0}
	_GotoState155Action                                               = &_Action{_ShiftAction, _State155, 0}
	_GotoState156Action                                               = &_Action{_ShiftAction, _State156, 0}
	_GotoState157Action                                               = &_Action{_ShiftAction, _State157, 0}
	_GotoState158Action                                               = &_Action{_ShiftAction, _State158, 0}
	_GotoState159Action                                               = &_Action{_ShiftAction, _State159, 0}
	_GotoState160Action                                               = &_Action{_ShiftAction, _State160, 0}
	_GotoState161Action                                               = &_Action{_ShiftAction, _State161, 0}
	_GotoState162Action                                               = &_Action{_ShiftAction, _State162, 0}
	_GotoState163Action                                               = &_Action{_ShiftAction, _State163, 0}
	_GotoState164Action                                               = &_Action{_ShiftAction, _State164, 0}
	_GotoState165Action                                               = &_Action{_ShiftAction, _State165, 0}
	_GotoState166Action                                               = &_Action{_ShiftAction, _State166, 0}
	_GotoState167Action                                               = &_Action{_ShiftAction, _State167, 0}
	_GotoState168Action                                               = &_Action{_ShiftAction, _State168, 0}
	_GotoState169Action                                               = &_Action{_ShiftAction, _State169, 0}
	_GotoState170Action                                               = &_Action{_ShiftAction, _State170, 0}
	_GotoState171Action                                               = &_Action{_ShiftAction, _State171, 0}
	_GotoState172Action                                               = &_Action{_ShiftAction, _State172, 0}
	_GotoState173Action                                               = &_Action{_ShiftAction, _State173, 0}
	_GotoState174Action                                               = &_Action{_ShiftAction, _State174, 0}
	_GotoState175Action                                               = &_Action{_ShiftAction, _State175, 0}
	_GotoState176Action                                               = &_Action{_ShiftAction, _State176, 0}
	_GotoState177Action                                               = &_Action{_ShiftAction, _State177, 0}
	_GotoState178Action                                               = &_Action{_ShiftAction, _State178, 0}
	_GotoState179Action                                               = &_Action{_ShiftAction, _State179, 0}
	_GotoState180Action                                               = &_Action{_ShiftAction, _State180, 0}
	_GotoState181Action                                               = &_Action{_ShiftAction, _State181, 0}
	_GotoState182Action                                               = &_Action{_ShiftAction, _State182, 0}
	_GotoState183Action                                               = &_Action{_ShiftAction, _State183, 0}
	_GotoState184Action                                               = &_Action{_ShiftAction, _State184, 0}
	_GotoState185Action                                               = &_Action{_ShiftAction, _State185, 0}
	_GotoState186Action                                               = &_Action{_ShiftAction, _State186, 0}
	_GotoState187Action                                               = &_Action{_ShiftAction, _State187, 0}
	_GotoState188Action                                               = &_Action{_ShiftAction, _State188, 0}
	_GotoState189Action                                               = &_Action{_ShiftAction, _State189, 0}
	_GotoState190Action                                               = &_Action{_ShiftAction, _State190, 0}
	_GotoState191Action                                               = &_Action{_ShiftAction, _State191, 0}
	_GotoState192Action                                               = &_Action{_ShiftAction, _State192, 0}
	_GotoState193Action                                               = &_Action{_ShiftAction, _State193, 0}
	_GotoState194Action                                               = &_Action{_ShiftAction, _State194, 0}
	_GotoState195Action                                               = &_Action{_ShiftAction, _State195, 0}
	_GotoState196Action                                               = &_Action{_ShiftAction, _State196, 0}
	_GotoState197Action                                               = &_Action{_ShiftAction, _State197, 0}
	_GotoState198Action                                               = &_Action{_ShiftAction, _State198, 0}
	_GotoState199Action                                               = &_Action{_ShiftAction, _State199, 0}
	_GotoState200Action                                               = &_Action{_ShiftAction, _State200, 0}
	_GotoState201Action                                               = &_Action{_ShiftAction, _State201, 0}
	_GotoState202Action                                               = &_Action{_ShiftAction, _State202, 0}
	_GotoState203Action                                               = &_Action{_ShiftAction, _State203, 0}
	_GotoState204Action                                               = &_Action{_ShiftAction, _State204, 0}
	_GotoState205Action                                               = &_Action{_ShiftAction, _State205, 0}
	_GotoState206Action                                               = &_Action{_ShiftAction, _State206, 0}
	_GotoState207Action                                               = &_Action{_ShiftAction, _State207, 0}
	_GotoState208Action                                               = &_Action{_ShiftAction, _State208, 0}
	_GotoState209Action                                               = &_Action{_ShiftAction, _State209, 0}
	_GotoState210Action                                               = &_Action{_ShiftAction, _State210, 0}
	_GotoState211Action                                               = &_Action{_ShiftAction, _State211, 0}
	_GotoState212Action                                               = &_Action{_ShiftAction, _State212, 0}
	_GotoState213Action                                               = &_Action{_ShiftAction, _State213, 0}
	_GotoState214Action                                               = &_Action{_ShiftAction, _State214, 0}
	_GotoState215Action                                               = &_Action{_ShiftAction, _State215, 0}
	_GotoState216Action                                               = &_Action{_ShiftAction, _State216, 0}
	_GotoState217Action                                               = &_Action{_ShiftAction, _State217, 0}
	_GotoState218Action                                               = &_Action{_ShiftAction, _State218, 0}
	_GotoState219Action                                               = &_Action{_ShiftAction, _State219, 0}
	_GotoState220Action                                               = &_Action{_ShiftAction, _State220, 0}
	_GotoState221Action                                               = &_Action{_ShiftAction, _State221, 0}
	_GotoState222Action                                               = &_Action{_ShiftAction, _State222, 0}
	_GotoState223Action                                               = &_Action{_ShiftAction, _State223, 0}
	_GotoState224Action                                               = &_Action{_ShiftAction, _State224, 0}
	_GotoState225Action                                               = &_Action{_ShiftAction, _State225, 0}
	_GotoState226Action                                               = &_Action{_ShiftAction, _State226, 0}
	_GotoState227Action                                               = &_Action{_ShiftAction, _State227, 0}
	_GotoState228Action                                               = &_Action{_ShiftAction, _State228, 0}
	_GotoState229Action                                               = &_Action{_ShiftAction, _State229, 0}
	_GotoState230Action                                               = &_Action{_ShiftAction, _State230, 0}
	_GotoState231Action                                               = &_Action{_ShiftAction, _State231, 0}
	_GotoState232Action                                               = &_Action{_ShiftAction, _State232, 0}
	_GotoState233Action                                               = &_Action{_ShiftAction, _State233, 0}
	_GotoState234Action                                               = &_Action{_ShiftAction, _State234, 0}
	_GotoState235Action                                               = &_Action{_ShiftAction, _State235, 0}
	_GotoState236Action                                               = &_Action{_ShiftAction, _State236, 0}
	_GotoState237Action                                               = &_Action{_ShiftAction, _State237, 0}
	_GotoState238Action                                               = &_Action{_ShiftAction, _State238, 0}
	_GotoState239Action                                               = &_Action{_ShiftAction, _State239, 0}
	_GotoState240Action                                               = &_Action{_ShiftAction, _State240, 0}
	_GotoState241Action                                               = &_Action{_ShiftAction, _State241, 0}
	_GotoState242Action                                               = &_Action{_ShiftAction, _State242, 0}
	_GotoState243Action                                               = &_Action{_ShiftAction, _State243, 0}
	_GotoState244Action                                               = &_Action{_ShiftAction, _State244, 0}
	_GotoState245Action                                               = &_Action{_ShiftAction, _State245, 0}
	_GotoState246Action                                               = &_Action{_ShiftAction, _State246, 0}
	_GotoState247Action                                               = &_Action{_ShiftAction, _State247, 0}
	_GotoState248Action                                               = &_Action{_ShiftAction, _State248, 0}
	_GotoState249Action                                               = &_Action{_ShiftAction, _State249, 0}
	_GotoState250Action                                               = &_Action{_ShiftAction, _State250, 0}
	_GotoState251Action                                               = &_Action{_ShiftAction, _State251, 0}
	_GotoState252Action                                               = &_Action{_ShiftAction, _State252, 0}
	_GotoState253Action                                               = &_Action{_ShiftAction, _State253, 0}
	_GotoState254Action                                               = &_Action{_ShiftAction, _State254, 0}
	_GotoState255Action                                               = &_Action{_ShiftAction, _State255, 0}
	_GotoState256Action                                               = &_Action{_ShiftAction, _State256, 0}
	_GotoState257Action                                               = &_Action{_ShiftAction, _State257, 0}
	_GotoState258Action                                               = &_Action{_ShiftAction, _State258, 0}
	_GotoState259Action                                               = &_Action{_ShiftAction, _State259, 0}
	_GotoState260Action                                               = &_Action{_ShiftAction, _State260, 0}
	_GotoState261Action                                               = &_Action{_ShiftAction, _State261, 0}
	_GotoState262Action                                               = &_Action{_ShiftAction, _State262, 0}
	_GotoState263Action                                               = &_Action{_ShiftAction, _State263, 0}
	_GotoState264Action                                               = &_Action{_ShiftAction, _State264, 0}
	_GotoState265Action                                               = &_Action{_ShiftAction, _State265, 0}
	_GotoState266Action                                               = &_Action{_ShiftAction, _State266, 0}
	_GotoState267Action                                               = &_Action{_ShiftAction, _State267, 0}
	_GotoState268Action                                               = &_Action{_ShiftAction, _State268, 0}
	_GotoState269Action                                               = &_Action{_ShiftAction, _State269, 0}
	_GotoState270Action                                               = &_Action{_ShiftAction, _State270, 0}
	_GotoState271Action                                               = &_Action{_ShiftAction, _State271, 0}
	_GotoState272Action                                               = &_Action{_ShiftAction, _State272, 0}
	_GotoState273Action                                               = &_Action{_ShiftAction, _State273, 0}
	_GotoState274Action                                               = &_Action{_ShiftAction, _State274, 0}
	_GotoState275Action                                               = &_Action{_ShiftAction, _State275, 0}
	_GotoState276Action                                               = &_Action{_ShiftAction, _State276, 0}
	_GotoState277Action                                               = &_Action{_ShiftAction, _State277, 0}
	_GotoState278Action                                               = &_Action{_ShiftAction, _State278, 0}
	_GotoState279Action                                               = &_Action{_ShiftAction, _State279, 0}
	_GotoState280Action                                               = &_Action{_ShiftAction, _State280, 0}
	_GotoState281Action                                               = &_Action{_ShiftAction, _State281, 0}
	_GotoState282Action                                               = &_Action{_ShiftAction, _State282, 0}
	_GotoState283Action                                               = &_Action{_ShiftAction, _State283, 0}
	_GotoState284Action                                               = &_Action{_ShiftAction, _State284, 0}
	_GotoState285Action                                               = &_Action{_ShiftAction, _State285, 0}
	_GotoState286Action                                               = &_Action{_ShiftAction, _State286, 0}
	_GotoState287Action                                               = &_Action{_ShiftAction, _State287, 0}
	_GotoState288Action                                               = &_Action{_ShiftAction, _State288, 0}
	_GotoState289Action                                               = &_Action{_ShiftAction, _State289, 0}
	_GotoState290Action                                               = &_Action{_ShiftAction, _State290, 0}
	_GotoState291Action                                               = &_Action{_ShiftAction, _State291, 0}
	_GotoState292Action                                               = &_Action{_ShiftAction, _State292, 0}
	_GotoState293Action                                               = &_Action{_ShiftAction, _State293, 0}
	_GotoState294Action                                               = &_Action{_ShiftAction, _State294, 0}
	_GotoState295Action                                               = &_Action{_ShiftAction, _State295, 0}
	_GotoState296Action                                               = &_Action{_ShiftAction, _State296, 0}
	_GotoState297Action                                               = &_Action{_ShiftAction, _State297, 0}
	_GotoState298Action                                               = &_Action{_ShiftAction, _State298, 0}
	_GotoState299Action                                               = &_Action{_ShiftAction, _State299, 0}
	_GotoState300Action                                               = &_Action{_ShiftAction, _State300, 0}
	_GotoState301Action                                               = &_Action{_ShiftAction, _State301, 0}
	_GotoState302Action                                               = &_Action{_ShiftAction, _State302, 0}
	_GotoState303Action                                               = &_Action{_ShiftAction, _State303, 0}
	_GotoState304Action                                               = &_Action{_ShiftAction, _State304, 0}
	_GotoState305Action                                               = &_Action{_ShiftAction, _State305, 0}
	_GotoState306Action                                               = &_Action{_ShiftAction, _State306, 0}
	_GotoState307Action                                               = &_Action{_ShiftAction, _State307, 0}
	_GotoState308Action                                               = &_Action{_ShiftAction, _State308, 0}
	_GotoState309Action                                               = &_Action{_ShiftAction, _State309, 0}
	_GotoState310Action                                               = &_Action{_ShiftAction, _State310, 0}
	_GotoState311Action                                               = &_Action{_ShiftAction, _State311, 0}
	_GotoState312Action                                               = &_Action{_ShiftAction, _State312, 0}
	_GotoState313Action                                               = &_Action{_ShiftAction, _State313, 0}
	_GotoState314Action                                               = &_Action{_ShiftAction, _State314, 0}
	_GotoState315Action                                               = &_Action{_ShiftAction, _State315, 0}
	_GotoState316Action                                               = &_Action{_ShiftAction, _State316, 0}
	_GotoState317Action                                               = &_Action{_ShiftAction, _State317, 0}
	_GotoState318Action                                               = &_Action{_ShiftAction, _State318, 0}
	_GotoState319Action                                               = &_Action{_ShiftAction, _State319, 0}
	_GotoState320Action                                               = &_Action{_ShiftAction, _State320, 0}
	_GotoState321Action                                               = &_Action{_ShiftAction, _State321, 0}
	_GotoState322Action                                               = &_Action{_ShiftAction, _State322, 0}
	_GotoState323Action                                               = &_Action{_ShiftAction, _State323, 0}
	_GotoState324Action                                               = &_Action{_ShiftAction, _State324, 0}
	_GotoState325Action                                               = &_Action{_ShiftAction, _State325, 0}
	_GotoState326Action                                               = &_Action{_ShiftAction, _State326, 0}
	_GotoState327Action                                               = &_Action{_ShiftAction, _State327, 0}
	_GotoState328Action                                               = &_Action{_ShiftAction, _State328, 0}
	_GotoState329Action                                               = &_Action{_ShiftAction, _State329, 0}
	_GotoState330Action                                               = &_Action{_ShiftAction, _State330, 0}
	_GotoState331Action                                               = &_Action{_ShiftAction, _State331, 0}
	_GotoState332Action                                               = &_Action{_ShiftAction, _State332, 0}
	_GotoState333Action                                               = &_Action{_ShiftAction, _State333, 0}
	_GotoState334Action                                               = &_Action{_ShiftAction, _State334, 0}
	_GotoState335Action                                               = &_Action{_ShiftAction, _State335, 0}
	_GotoState336Action                                               = &_Action{_ShiftAction, _State336, 0}
	_GotoState337Action                                               = &_Action{_ShiftAction, _State337, 0}
	_GotoState338Action                                               = &_Action{_ShiftAction, _State338, 0}
	_GotoState339Action                                               = &_Action{_ShiftAction, _State339, 0}
	_GotoState340Action                                               = &_Action{_ShiftAction, _State340, 0}
	_GotoState341Action                                               = &_Action{_ShiftAction, _State341, 0}
	_GotoState342Action                                               = &_Action{_ShiftAction, _State342, 0}
	_GotoState343Action                                               = &_Action{_ShiftAction, _State343, 0}
	_GotoState344Action                                               = &_Action{_ShiftAction, _State344, 0}
	_GotoState345Action                                               = &_Action{_ShiftAction, _State345, 0}
	_GotoState346Action                                               = &_Action{_ShiftAction, _State346, 0}
	_GotoState347Action                                               = &_Action{_ShiftAction, _State347, 0}
	_GotoState348Action                                               = &_Action{_ShiftAction, _State348, 0}
	_GotoState349Action                                               = &_Action{_ShiftAction, _State349, 0}
	_GotoState350Action                                               = &_Action{_ShiftAction, _State350, 0}
	_GotoState351Action                                               = &_Action{_ShiftAction, _State351, 0}
	_GotoState352Action                                               = &_Action{_ShiftAction, _State352, 0}
	_GotoState353Action                                               = &_Action{_ShiftAction, _State353, 0}
	_GotoState354Action                                               = &_Action{_ShiftAction, _State354, 0}
	_GotoState355Action                                               = &_Action{_ShiftAction, _State355, 0}
	_GotoState356Action                                               = &_Action{_ShiftAction, _State356, 0}
	_GotoState357Action                                               = &_Action{_ShiftAction, _State357, 0}
	_GotoState358Action                                               = &_Action{_ShiftAction, _State358, 0}
	_GotoState359Action                                               = &_Action{_ShiftAction, _State359, 0}
	_GotoState360Action                                               = &_Action{_ShiftAction, _State360, 0}
	_GotoState361Action                                               = &_Action{_ShiftAction, _State361, 0}
	_GotoState362Action                                               = &_Action{_ShiftAction, _State362, 0}
	_GotoState363Action                                               = &_Action{_ShiftAction, _State363, 0}
	_GotoState364Action                                               = &_Action{_ShiftAction, _State364, 0}
	_GotoState365Action                                               = &_Action{_ShiftAction, _State365, 0}
	_GotoState366Action                                               = &_Action{_ShiftAction, _State366, 0}
	_GotoState367Action                                               = &_Action{_ShiftAction, _State367, 0}
	_GotoState368Action                                               = &_Action{_ShiftAction, _State368, 0}
	_GotoState369Action                                               = &_Action{_ShiftAction, _State369, 0}
	_GotoState370Action                                               = &_Action{_ShiftAction, _State370, 0}
	_GotoState371Action                                               = &_Action{_ShiftAction, _State371, 0}
	_GotoState372Action                                               = &_Action{_ShiftAction, _State372, 0}
	_GotoState373Action                                               = &_Action{_ShiftAction, _State373, 0}
	_GotoState374Action                                               = &_Action{_ShiftAction, _State374, 0}
	_GotoState375Action                                               = &_Action{_ShiftAction, _State375, 0}
	_GotoState376Action                                               = &_Action{_ShiftAction, _State376, 0}
	_GotoState377Action                                               = &_Action{_ShiftAction, _State377, 0}
	_GotoState378Action                                               = &_Action{_ShiftAction, _State378, 0}
	_GotoState379Action                                               = &_Action{_ShiftAction, _State379, 0}
	_GotoState380Action                                               = &_Action{_ShiftAction, _State380, 0}
	_GotoState381Action                                               = &_Action{_ShiftAction, _State381, 0}
	_GotoState382Action                                               = &_Action{_ShiftAction, _State382, 0}
	_GotoState383Action                                               = &_Action{_ShiftAction, _State383, 0}
	_GotoState384Action                                               = &_Action{_ShiftAction, _State384, 0}
	_GotoState385Action                                               = &_Action{_ShiftAction, _State385, 0}
	_GotoState386Action                                               = &_Action{_ShiftAction, _State386, 0}
	_GotoState387Action                                               = &_Action{_ShiftAction, _State387, 0}
	_GotoState388Action                                               = &_Action{_ShiftAction, _State388, 0}
	_GotoState389Action                                               = &_Action{_ShiftAction, _State389, 0}
	_GotoState390Action                                               = &_Action{_ShiftAction, _State390, 0}
	_GotoState391Action                                               = &_Action{_ShiftAction, _State391, 0}
	_GotoState392Action                                               = &_Action{_ShiftAction, _State392, 0}
	_GotoState393Action                                               = &_Action{_ShiftAction, _State393, 0}
	_GotoState394Action                                               = &_Action{_ShiftAction, _State394, 0}
	_GotoState395Action                                               = &_Action{_ShiftAction, _State395, 0}
	_GotoState396Action                                               = &_Action{_ShiftAction, _State396, 0}
	_GotoState397Action                                               = &_Action{_ShiftAction, _State397, 0}
	_GotoState398Action                                               = &_Action{_ShiftAction, _State398, 0}
	_GotoState399Action                                               = &_Action{_ShiftAction, _State399, 0}
	_GotoState400Action                                               = &_Action{_ShiftAction, _State400, 0}
	_GotoState401Action                                               = &_Action{_ShiftAction, _State401, 0}
	_GotoState402Action                                               = &_Action{_ShiftAction, _State402, 0}
	_GotoState403Action                                               = &_Action{_ShiftAction, _State403, 0}
	_GotoState404Action                                               = &_Action{_ShiftAction, _State404, 0}
	_GotoState405Action                                               = &_Action{_ShiftAction, _State405, 0}
	_GotoState406Action                                               = &_Action{_ShiftAction, _State406, 0}
	_GotoState407Action                                               = &_Action{_ShiftAction, _State407, 0}
	_GotoState408Action                                               = &_Action{_ShiftAction, _State408, 0}
	_GotoState409Action                                               = &_Action{_ShiftAction, _State409, 0}
	_GotoState410Action                                               = &_Action{_ShiftAction, _State410, 0}
	_GotoState411Action                                               = &_Action{_ShiftAction, _State411, 0}
	_GotoState412Action                                               = &_Action{_ShiftAction, _State412, 0}
	_GotoState413Action                                               = &_Action{_ShiftAction, _State413, 0}
	_GotoState414Action                                               = &_Action{_ShiftAction, _State414, 0}
	_GotoState415Action                                               = &_Action{_ShiftAction, _State415, 0}
	_GotoState416Action                                               = &_Action{_ShiftAction, _State416, 0}
	_GotoState417Action                                               = &_Action{_ShiftAction, _State417, 0}
	_GotoState418Action                                               = &_Action{_ShiftAction, _State418, 0}
	_GotoState419Action                                               = &_Action{_ShiftAction, _State419, 0}
	_GotoState420Action                                               = &_Action{_ShiftAction, _State420, 0}
	_GotoState421Action                                               = &_Action{_ShiftAction, _State421, 0}
	_GotoState422Action                                               = &_Action{_ShiftAction, _State422, 0}
	_GotoState423Action                                               = &_Action{_ShiftAction, _State423, 0}
	_GotoState424Action                                               = &_Action{_ShiftAction, _State424, 0}
	_GotoState425Action                                               = &_Action{_ShiftAction, _State425, 0}
	_GotoState426Action                                               = &_Action{_ShiftAction, _State426, 0}
	_GotoState427Action                                               = &_Action{_ShiftAction, _State427, 0}
	_GotoState428Action                                               = &_Action{_ShiftAction, _State428, 0}
	_GotoState429Action                                               = &_Action{_ShiftAction, _State429, 0}
	_GotoState430Action                                               = &_Action{_ShiftAction, _State430, 0}
	_GotoState431Action                                               = &_Action{_ShiftAction, _State431, 0}
	_GotoState432Action                                               = &_Action{_ShiftAction, _State432, 0}
	_GotoState433Action                                               = &_Action{_ShiftAction, _State433, 0}
	_GotoState434Action                                               = &_Action{_ShiftAction, _State434, 0}
	_GotoState435Action                                               = &_Action{_ShiftAction, _State435, 0}
	_GotoState436Action                                               = &_Action{_ShiftAction, _State436, 0}
	_GotoState437Action                                               = &_Action{_ShiftAction, _State437, 0}
	_GotoState438Action                                               = &_Action{_ShiftAction, _State438, 0}
	_GotoState439Action                                               = &_Action{_ShiftAction, _State439, 0}
	_GotoState440Action                                               = &_Action{_ShiftAction, _State440, 0}
	_GotoState441Action                                               = &_Action{_ShiftAction, _State441, 0}
	_GotoState442Action                                               = &_Action{_ShiftAction, _State442, 0}
	_GotoState443Action                                               = &_Action{_ShiftAction, _State443, 0}
	_GotoState444Action                                               = &_Action{_ShiftAction, _State444, 0}
	_GotoState445Action                                               = &_Action{_ShiftAction, _State445, 0}
	_GotoState446Action                                               = &_Action{_ShiftAction, _State446, 0}
	_GotoState447Action                                               = &_Action{_ShiftAction, _State447, 0}
	_GotoState448Action                                               = &_Action{_ShiftAction, _State448, 0}
	_GotoState449Action                                               = &_Action{_ShiftAction, _State449, 0}
	_GotoState450Action                                               = &_Action{_ShiftAction, _State450, 0}
	_GotoState451Action                                               = &_Action{_ShiftAction, _State451, 0}
	_GotoState452Action                                               = &_Action{_ShiftAction, _State452, 0}
	_GotoState453Action                                               = &_Action{_ShiftAction, _State453, 0}
	_GotoState454Action                                               = &_Action{_ShiftAction, _State454, 0}
	_GotoState455Action                                               = &_Action{_ShiftAction, _State455, 0}
	_GotoState456Action                                               = &_Action{_ShiftAction, _State456, 0}
	_GotoState457Action                                               = &_Action{_ShiftAction, _State457, 0}
	_GotoState458Action                                               = &_Action{_ShiftAction, _State458, 0}
	_GotoState459Action                                               = &_Action{_ShiftAction, _State459, 0}
	_GotoState460Action                                               = &_Action{_ShiftAction, _State460, 0}
	_GotoState461Action                                               = &_Action{_ShiftAction, _State461, 0}
	_GotoState462Action                                               = &_Action{_ShiftAction, _State462, 0}
	_GotoState463Action                                               = &_Action{_ShiftAction, _State463, 0}
	_GotoState464Action                                               = &_Action{_ShiftAction, _State464, 0}
	_GotoState465Action                                               = &_Action{_ShiftAction, _State465, 0}
	_GotoState466Action                                               = &_Action{_ShiftAction, _State466, 0}
	_GotoState467Action                                               = &_Action{_ShiftAction, _State467, 0}
	_GotoState468Action                                               = &_Action{_ShiftAction, _State468, 0}
	_GotoState469Action                                               = &_Action{_ShiftAction, _State469, 0}
	_GotoState470Action                                               = &_Action{_ShiftAction, _State470, 0}
	_ReduceToSourceAction                                             = &_Action{_ReduceAction, 0, _ReduceToSource}
	_ReduceAddToDefinitionsAction                                     = &_Action{_ReduceAction, 0, _ReduceAddToDefinitions}
	_ReduceNilToDefinitionsAction                                     = &_Action{_ReduceAction, 0, _ReduceNilToDefinitions}
	_ReducePackageDefToDefinitionAction                               = &_Action{_ReduceAction, 0, _ReducePackageDefToDefinition}
	_ReduceTypeDefToDefinitionAction                                  = &_Action{_ReduceAction, 0, _ReduceTypeDefToDefinition}
	_ReduceNamedFuncDefToDefinitionAction                             = &_Action{_ReduceAction, 0, _ReduceNamedFuncDefToDefinition}
	_ReduceGlobalVarDefToDefinitionAction                             = &_Action{_ReduceAction, 0, _ReduceGlobalVarDefToDefinition}
	_ReduceGlobalVarAssignmentToDefinitionAction                      = &_Action{_ReduceAction, 0, _ReduceGlobalVarAssignmentToDefinition}
	_ReduceStatementBlockToDefinitionAction                           = &_Action{_ReduceAction, 0, _ReduceStatementBlockToDefinition}
	_ReduceCommentGroupsToDefinitionAction                            = &_Action{_ReduceAction, 0, _ReduceCommentGroupsToDefinition}
	_ReduceToStatementBlockAction                                     = &_Action{_ReduceAction, 0, _ReduceToStatementBlock}
	_ReduceAddImplicitToProperStatementsAction                        = &_Action{_ReduceAction, 0, _ReduceAddImplicitToProperStatements}
	_ReduceAddExplicitToProperStatementsAction                        = &_Action{_ReduceAction, 0, _ReduceAddExplicitToProperStatements}
	_ReduceStatementToProperStatementsAction                          = &_Action{_ReduceAction, 0, _ReduceStatementToProperStatements}
	_ReduceProperStatementsToStatementsAction                         = &_Action{_ReduceAction, 0, _ReduceProperStatementsToStatements}
	_ReduceImproperImplicitToStatementsAction                         = &_Action{_ReduceAction, 0, _ReduceImproperImplicitToStatements}
	_ReduceImproperExplicitToStatementsAction                         = &_Action{_ReduceAction, 0, _ReduceImproperExplicitToStatements}
	_ReduceNilToStatementsAction                                      = &_Action{_ReduceAction, 0, _ReduceNilToStatements}
	_ReduceUnsafeStatementToSimpleStatementAction                     = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToSimpleStatement}
	_ReduceExpressionOrImproperStructStatementToSimpleStatementAction = &_Action{_ReduceAction, 0, _ReduceExpressionOrImproperStructStatementToSimpleStatement}
	_ReduceCallbackOpStatementToSimpleStatementAction                 = &_Action{_ReduceAction, 0, _ReduceCallbackOpStatementToSimpleStatement}
	_ReduceJumpStatementToSimpleStatementAction                       = &_Action{_ReduceAction, 0, _ReduceJumpStatementToSimpleStatement}
	_ReduceAssignStatementToSimpleStatementAction                     = &_Action{_ReduceAction, 0, _ReduceAssignStatementToSimpleStatement}
	_ReduceUnaryOpAssignStatementToSimpleStatementAction              = &_Action{_ReduceAction, 0, _ReduceUnaryOpAssignStatementToSimpleStatement}
	_ReduceBinaryOpAssignStatementToSimpleStatementAction             = &_Action{_ReduceAction, 0, _ReduceBinaryOpAssignStatementToSimpleStatement}
	_ReduceFallthroughStatementToSimpleStatementAction                = &_Action{_ReduceAction, 0, _ReduceFallthroughStatementToSimpleStatement}
	_ReduceSimpleStatementToStatementAction                           = &_Action{_ReduceAction, 0, _ReduceSimpleStatementToStatement}
	_ReduceImportStatementToStatementAction                           = &_Action{_ReduceAction, 0, _ReduceImportStatementToStatement}
	_ReduceCaseBranchStatementToStatementAction                       = &_Action{_ReduceAction, 0, _ReduceCaseBranchStatementToStatement}
	_ReduceDefaultBranchStatementToStatementAction                    = &_Action{_ReduceAction, 0, _ReduceDefaultBranchStatementToStatement}
	_ReduceSimpleStatementToOptionalSimpleStatementAction             = &_Action{_ReduceAction, 0, _ReduceSimpleStatementToOptionalSimpleStatement}
	_ReduceNilToOptionalSimpleStatementAction                         = &_Action{_ReduceAction, 0, _ReduceNilToOptionalSimpleStatement}
	_ReduceToExpressionOrImproperStructStatementAction                = &_Action{_ReduceAction, 0, _ReduceToExpressionOrImproperStructStatement}
	_ReduceExpressionToExpressionsAction                              = &_Action{_ReduceAction, 0, _ReduceExpressionToExpressions}
	_ReduceAddToExpressionsAction                                     = &_Action{_ReduceAction, 0, _ReduceAddToExpressions}
	_ReduceAsyncToCallbackOpAction                                    = &_Action{_ReduceAction, 0, _ReduceAsyncToCallbackOp}
	_ReduceDeferToCallbackOpAction                                    = &_Action{_ReduceAction, 0, _ReduceDeferToCallbackOp}
	_ReduceToCallbackOpStatementAction                                = &_Action{_ReduceAction, 0, _ReduceToCallbackOpStatement}
	_ReduceToUnsafeStatementAction                                    = &_Action{_ReduceAction, 0, _ReduceToUnsafeStatement}
	_ReduceUnlabeledNoValueToJumpStatementAction                      = &_Action{_ReduceAction, 0, _ReduceUnlabeledNoValueToJumpStatement}
	_ReduceUnlabeledValuedToJumpStatementAction                       = &_Action{_ReduceAction, 0, _ReduceUnlabeledValuedToJumpStatement}
	_ReduceLabeledNoValueToJumpStatementAction                        = &_Action{_ReduceAction, 0, _ReduceLabeledNoValueToJumpStatement}
	_ReduceLabeledValuedToJumpStatementAction                         = &_Action{_ReduceAction, 0, _ReduceLabeledValuedToJumpStatement}
	_ReduceReturnToJumpTypeAction                                     = &_Action{_ReduceAction, 0, _ReduceReturnToJumpType}
	_ReduceBreakToJumpTypeAction                                      = &_Action{_ReduceAction, 0, _ReduceBreakToJumpType}
	_ReduceContinueToJumpTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceContinueToJumpType}
	_ReduceToFallthroughStatementAction                               = &_Action{_ReduceAction, 0, _ReduceToFallthroughStatement}
	_ReduceToAssignStatementAction                                    = &_Action{_ReduceAction, 0, _ReduceToAssignStatement}
	_ReduceToUnaryOpAssignStatementAction                             = &_Action{_ReduceAction, 0, _ReduceToUnaryOpAssignStatement}
	_ReduceAddOneAssignToUnaryOpAssignAction                          = &_Action{_ReduceAction, 0, _ReduceAddOneAssignToUnaryOpAssign}
	_ReduceSubOneAssignToUnaryOpAssignAction                          = &_Action{_ReduceAction, 0, _ReduceSubOneAssignToUnaryOpAssign}
	_ReduceToBinaryOpAssignStatementAction                            = &_Action{_ReduceAction, 0, _ReduceToBinaryOpAssignStatement}
	_ReduceAddAssignToBinaryOpAssignAction                            = &_Action{_ReduceAction, 0, _ReduceAddAssignToBinaryOpAssign}
	_ReduceSubAssignToBinaryOpAssignAction                            = &_Action{_ReduceAction, 0, _ReduceSubAssignToBinaryOpAssign}
	_ReduceMulAssignToBinaryOpAssignAction                            = &_Action{_ReduceAction, 0, _ReduceMulAssignToBinaryOpAssign}
	_ReduceDivAssignToBinaryOpAssignAction                            = &_Action{_ReduceAction, 0, _ReduceDivAssignToBinaryOpAssign}
	_ReduceModAssignToBinaryOpAssignAction                            = &_Action{_ReduceAction, 0, _ReduceModAssignToBinaryOpAssign}
	_ReduceBitNegAssignToBinaryOpAssignAction                         = &_Action{_ReduceAction, 0, _ReduceBitNegAssignToBinaryOpAssign}
	_ReduceBitAndAssignToBinaryOpAssignAction                         = &_Action{_ReduceAction, 0, _ReduceBitAndAssignToBinaryOpAssign}
	_ReduceBitOrAssignToBinaryOpAssignAction                          = &_Action{_ReduceAction, 0, _ReduceBitOrAssignToBinaryOpAssign}
	_ReduceBitXorAssignToBinaryOpAssignAction                         = &_Action{_ReduceAction, 0, _ReduceBitXorAssignToBinaryOpAssign}
	_ReduceBitLshiftAssignToBinaryOpAssignAction                      = &_Action{_ReduceAction, 0, _ReduceBitLshiftAssignToBinaryOpAssign}
	_ReduceBitRshiftAssignToBinaryOpAssignAction                      = &_Action{_ReduceAction, 0, _ReduceBitRshiftAssignToBinaryOpAssign}
	_ReduceSingleToImportStatementAction                              = &_Action{_ReduceAction, 0, _ReduceSingleToImportStatement}
	_ReduceMultipleToImportStatementAction                            = &_Action{_ReduceAction, 0, _ReduceMultipleToImportStatement}
	_ReduceStringLiteralToImportClauseAction                          = &_Action{_ReduceAction, 0, _ReduceStringLiteralToImportClause}
	_ReduceAliasToImportClauseAction                                  = &_Action{_ReduceAction, 0, _ReduceAliasToImportClause}
	_ReduceAddImplicitToProperImportClausesAction                     = &_Action{_ReduceAction, 0, _ReduceAddImplicitToProperImportClauses}
	_ReduceAddExplicitToProperImportClausesAction                     = &_Action{_ReduceAction, 0, _ReduceAddExplicitToProperImportClauses}
	_ReduceImportClauseToProperImportClausesAction                    = &_Action{_ReduceAction, 0, _ReduceImportClauseToProperImportClauses}
	_ReduceProperImportClausesToImportClausesAction                   = &_Action{_ReduceAction, 0, _ReduceProperImportClausesToImportClauses}
	_ReduceImplicitToImportClausesAction                              = &_Action{_ReduceAction, 0, _ReduceImplicitToImportClauses}
	_ReduceExplicitToImportClausesAction                              = &_Action{_ReduceAction, 0, _ReduceExplicitToImportClauses}
	_ReduceCasePatternToCasePatternsAction                            = &_Action{_ReduceAction, 0, _ReduceCasePatternToCasePatterns}
	_ReduceMultipleToCasePatternsAction                               = &_Action{_ReduceAction, 0, _ReduceMultipleToCasePatterns}
	_ReduceToVarDeclPatternAction                                     = &_Action{_ReduceAction, 0, _ReduceToVarDeclPattern}
	_ReduceVarToVarOrLetAction                                        = &_Action{_ReduceAction, 0, _ReduceVarToVarOrLet}
	_ReduceLetToVarOrLetAction                                        = &_Action{_ReduceAction, 0, _ReduceLetToVarOrLet}
	_ReduceIdentifierToVarPatternAction                               = &_Action{_ReduceAction, 0, _ReduceIdentifierToVarPattern}
	_ReduceTuplePatternToVarPatternAction                             = &_Action{_ReduceAction, 0, _ReduceTuplePatternToVarPattern}
	_ReduceToTuplePatternAction                                       = &_Action{_ReduceAction, 0, _ReduceToTuplePattern}
	_ReduceFieldVarPatternToFieldVarPatternsAction                    = &_Action{_ReduceAction, 0, _ReduceFieldVarPatternToFieldVarPatterns}
	_ReduceAddToFieldVarPatternsAction                                = &_Action{_ReduceAction, 0, _ReduceAddToFieldVarPatterns}
	_ReducePositionalToFieldVarPatternAction                          = &_Action{_ReduceAction, 0, _ReducePositionalToFieldVarPattern}
	_ReduceNamedToFieldVarPatternAction                               = &_Action{_ReduceAction, 0, _ReduceNamedToFieldVarPattern}
	_ReduceEllipsisToFieldVarPatternAction                            = &_Action{_ReduceAction, 0, _ReduceEllipsisToFieldVarPattern}
	_ReduceValueTypeToOptionalValueTypeAction                         = &_Action{_ReduceAction, 0, _ReduceValueTypeToOptionalValueType}
	_ReduceNilToOptionalValueTypeAction                               = &_Action{_ReduceAction, 0, _ReduceNilToOptionalValueType}
	_ReduceToAssignPatternAction                                      = &_Action{_ReduceAction, 0, _ReduceToAssignPattern}
	_ReduceAssignPatternToCasePatternAction                           = &_Action{_ReduceAction, 0, _ReduceAssignPatternToCasePattern}
	_ReduceEnumMatchPatternToCasePatternAction                        = &_Action{_ReduceAction, 0, _ReduceEnumMatchPatternToCasePattern}
	_ReduceEnumVarDeclPatternToCasePatternAction                      = &_Action{_ReduceAction, 0, _ReduceEnumVarDeclPatternToCasePattern}
	_ReduceIfExprToExpressionAction                                   = &_Action{_ReduceAction, 0, _ReduceIfExprToExpression}
	_ReduceSwitchExprToExpressionAction                               = &_Action{_ReduceAction, 0, _ReduceSwitchExprToExpression}
	_ReduceLoopExprToExpressionAction                                 = &_Action{_ReduceAction, 0, _ReduceLoopExprToExpression}
	_ReduceSequenceExprToExpressionAction                             = &_Action{_ReduceAction, 0, _ReduceSequenceExprToExpression}
	_ReduceLabelDeclToOptionalLabelDeclAction                         = &_Action{_ReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}
	_ReduceUnlabelledToOptionalLabelDeclAction                        = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}
	_ReduceOrExprToSequenceExprAction                                 = &_Action{_ReduceAction, 0, _ReduceOrExprToSequenceExpr}
	_ReduceVarDeclPatternToSequenceExprAction                         = &_Action{_ReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}
	_ReduceAssignVarPatternToSequenceExprAction                       = &_Action{_ReduceAction, 0, _ReduceAssignVarPatternToSequenceExpr}
	_ReduceNoElseToIfExprAction                                       = &_Action{_ReduceAction, 0, _ReduceNoElseToIfExpr}
	_ReduceIfElseToIfExprAction                                       = &_Action{_ReduceAction, 0, _ReduceIfElseToIfExpr}
	_ReduceMultiIfElseToIfExprAction                                  = &_Action{_ReduceAction, 0, _ReduceMultiIfElseToIfExpr}
	_ReduceSequenceExprToConditionAction                              = &_Action{_ReduceAction, 0, _ReduceSequenceExprToCondition}
	_ReduceCaseToConditionAction                                      = &_Action{_ReduceAction, 0, _ReduceCaseToCondition}
	_ReduceToSwitchExprAction                                         = &_Action{_ReduceAction, 0, _ReduceToSwitchExpr}
	_ReduceInfiniteToLoopExprAction                                   = &_Action{_ReduceAction, 0, _ReduceInfiniteToLoopExpr}
	_ReduceDoWhileToLoopExprAction                                    = &_Action{_ReduceAction, 0, _ReduceDoWhileToLoopExpr}
	_ReduceWhileToLoopExprAction                                      = &_Action{_ReduceAction, 0, _ReduceWhileToLoopExpr}
	_ReduceIteratorToLoopExprAction                                   = &_Action{_ReduceAction, 0, _ReduceIteratorToLoopExpr}
	_ReduceForToLoopExprAction                                        = &_Action{_ReduceAction, 0, _ReduceForToLoopExpr}
	_ReduceSequenceExprToOptionalSequenceExprAction                   = &_Action{_ReduceAction, 0, _ReduceSequenceExprToOptionalSequenceExpr}
	_ReduceNilToOptionalSequenceExprAction                            = &_Action{_ReduceAction, 0, _ReduceNilToOptionalSequenceExpr}
	_ReduceSequenceExprToForAssignmentAction                          = &_Action{_ReduceAction, 0, _ReduceSequenceExprToForAssignment}
	_ReduceAssignToForAssignmentAction                                = &_Action{_ReduceAction, 0, _ReduceAssignToForAssignment}
	_ReduceToCallExprAction                                           = &_Action{_ReduceAction, 0, _ReduceToCallExpr}
	_ReduceBindingToOptionalGenericBindingAction                      = &_Action{_ReduceAction, 0, _ReduceBindingToOptionalGenericBinding}
	_ReduceNilToOptionalGenericBindingAction                          = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericBinding}
	_ReduceAddToProperGenericArgumentsAction                          = &_Action{_ReduceAction, 0, _ReduceAddToProperGenericArguments}
	_ReduceValueTypeToProperGenericArgumentsAction                    = &_Action{_ReduceAction, 0, _ReduceValueTypeToProperGenericArguments}
	_ReduceProperGenericArgumentsToGenericArgumentsAction             = &_Action{_ReduceAction, 0, _ReduceProperGenericArgumentsToGenericArguments}
	_ReduceImproperToGenericArgumentsAction                           = &_Action{_ReduceAction, 0, _ReduceImproperToGenericArguments}
	_ReduceNilToGenericArgumentsAction                                = &_Action{_ReduceAction, 0, _ReduceNilToGenericArguments}
	_ReduceAddToProperArgumentsAction                                 = &_Action{_ReduceAction, 0, _ReduceAddToProperArguments}
	_ReduceArgumentToProperArgumentsAction                            = &_Action{_ReduceAction, 0, _ReduceArgumentToProperArguments}
	_ReduceProperArgumentsToArgumentsAction                           = &_Action{_ReduceAction, 0, _ReduceProperArgumentsToArguments}
	_ReduceImproperToArgumentsAction                                  = &_Action{_ReduceAction, 0, _ReduceImproperToArguments}
	_ReduceNilToArgumentsAction                                       = &_Action{_ReduceAction, 0, _ReduceNilToArguments}
	_ReducePositionalToArgumentAction                                 = &_Action{_ReduceAction, 0, _ReducePositionalToArgument}
	_ReduceColonExprToArgumentAction                                  = &_Action{_ReduceAction, 0, _ReduceColonExprToArgument}
	_ReduceNamedAssignmentToArgumentAction                            = &_Action{_ReduceAction, 0, _ReduceNamedAssignmentToArgument}
	_ReduceVarargAssignmentToArgumentAction                           = &_Action{_ReduceAction, 0, _ReduceVarargAssignmentToArgument}
	_ReduceSkipPatternToArgumentAction                                = &_Action{_ReduceAction, 0, _ReduceSkipPatternToArgument}
	_ReduceUnitUnitPairToColonExprAction                              = &_Action{_ReduceAction, 0, _ReduceUnitUnitPairToColonExpr}
	_ReduceExprUnitPairToColonExprAction                              = &_Action{_ReduceAction, 0, _ReduceExprUnitPairToColonExpr}
	_ReduceUnitExprPairToColonExprAction                              = &_Action{_ReduceAction, 0, _ReduceUnitExprPairToColonExpr}
	_ReduceExprExprPairToColonExprAction                              = &_Action{_ReduceAction, 0, _ReduceExprExprPairToColonExpr}
	_ReduceColonExprUnitTupleToColonExprAction                        = &_Action{_ReduceAction, 0, _ReduceColonExprUnitTupleToColonExpr}
	_ReduceColonExprExprTupleToColonExprAction                        = &_Action{_ReduceAction, 0, _ReduceColonExprExprTupleToColonExpr}
	_ReduceParseErrorExprToAtomExprAction                             = &_Action{_ReduceAction, 0, _ReduceParseErrorExprToAtomExpr}
	_ReduceLiteralExprToAtomExprAction                                = &_Action{_ReduceAction, 0, _ReduceLiteralExprToAtomExpr}
	_ReduceIdentifierExprToAtomExprAction                             = &_Action{_ReduceAction, 0, _ReduceIdentifierExprToAtomExpr}
	_ReduceBlockExprToAtomExprAction                                  = &_Action{_ReduceAction, 0, _ReduceBlockExprToAtomExpr}
	_ReduceAnonymousFuncExprToAtomExprAction                          = &_Action{_ReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}
	_ReduceInitializeExprToAtomExprAction                             = &_Action{_ReduceAction, 0, _ReduceInitializeExprToAtomExpr}
	_ReduceImplicitStructExprToAtomExprAction                         = &_Action{_ReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}
	_ReduceToParseErrorExprAction                                     = &_Action{_ReduceAction, 0, _ReduceToParseErrorExpr}
	_ReduceTrueToLiteralExprAction                                    = &_Action{_ReduceAction, 0, _ReduceTrueToLiteralExpr}
	_ReduceFalseToLiteralExprAction                                   = &_Action{_ReduceAction, 0, _ReduceFalseToLiteralExpr}
	_ReduceIntegerLiteralToLiteralExprAction                          = &_Action{_ReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}
	_ReduceFloatLiteralToLiteralExprAction                            = &_Action{_ReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}
	_ReduceRuneLiteralToLiteralExprAction                             = &_Action{_ReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}
	_ReduceStringLiteralToLiteralExprAction                           = &_Action{_ReduceAction, 0, _ReduceStringLiteralToLiteralExpr}
	_ReduceToIdentifierExprAction                                     = &_Action{_ReduceAction, 0, _ReduceToIdentifierExpr}
	_ReduceToBlockExprAction                                          = &_Action{_ReduceAction, 0, _ReduceToBlockExpr}
	_ReduceToInitializeExprAction                                     = &_Action{_ReduceAction, 0, _ReduceToInitializeExpr}
	_ReduceToImplicitStructExprAction                                 = &_Action{_ReduceAction, 0, _ReduceToImplicitStructExpr}
	_ReduceAtomExprToAccessibleExprAction                             = &_Action{_ReduceAction, 0, _ReduceAtomExprToAccessibleExpr}
	_ReduceAccessExprToAccessibleExprAction                           = &_Action{_ReduceAction, 0, _ReduceAccessExprToAccessibleExpr}
	_ReduceCallExprToAccessibleExprAction                             = &_Action{_ReduceAction, 0, _ReduceCallExprToAccessibleExpr}
	_ReduceIndexExprToAccessibleExprAction                            = &_Action{_ReduceAction, 0, _ReduceIndexExprToAccessibleExpr}
	_ReduceToAccessExprAction                                         = &_Action{_ReduceAction, 0, _ReduceToAccessExpr}
	_ReduceToIndexExprAction                                          = &_Action{_ReduceAction, 0, _ReduceToIndexExpr}
	_ReduceAccessibleExprToPostfixableExprAction                      = &_Action{_ReduceAction, 0, _ReduceAccessibleExprToPostfixableExpr}
	_ReducePostfixUnaryExprToPostfixableExprAction                    = &_Action{_ReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}
	_ReduceToPostfixUnaryExprAction                                   = &_Action{_ReduceAction, 0, _ReduceToPostfixUnaryExpr}
	_ReducePostfixableExprToPrefixableExprAction                      = &_Action{_ReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}
	_ReducePrefixUnaryExprToPrefixableExprAction                      = &_Action{_ReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}
	_ReduceToPrefixUnaryExprAction                                    = &_Action{_ReduceAction, 0, _ReduceToPrefixUnaryExpr}
	_ReduceNotToPrefixUnaryOpAction                                   = &_Action{_ReduceAction, 0, _ReduceNotToPrefixUnaryOp}
	_ReduceBitNegToPrefixUnaryOpAction                                = &_Action{_ReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}
	_ReduceSubToPrefixUnaryOpAction                                   = &_Action{_ReduceAction, 0, _ReduceSubToPrefixUnaryOp}
	_ReduceMulToPrefixUnaryOpAction                                   = &_Action{_ReduceAction, 0, _ReduceMulToPrefixUnaryOp}
	_ReduceBitAndToPrefixUnaryOpAction                                = &_Action{_ReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}
	_ReducePrefixableExprToMulExprAction                              = &_Action{_ReduceAction, 0, _ReducePrefixableExprToMulExpr}
	_ReduceBinaryMulExprToMulExprAction                               = &_Action{_ReduceAction, 0, _ReduceBinaryMulExprToMulExpr}
	_ReduceToBinaryMulExprAction                                      = &_Action{_ReduceAction, 0, _ReduceToBinaryMulExpr}
	_ReduceMulToMulOpAction                                           = &_Action{_ReduceAction, 0, _ReduceMulToMulOp}
	_ReduceDivToMulOpAction                                           = &_Action{_ReduceAction, 0, _ReduceDivToMulOp}
	_ReduceModToMulOpAction                                           = &_Action{_ReduceAction, 0, _ReduceModToMulOp}
	_ReduceBitAndToMulOpAction                                        = &_Action{_ReduceAction, 0, _ReduceBitAndToMulOp}
	_ReduceBitLshiftToMulOpAction                                     = &_Action{_ReduceAction, 0, _ReduceBitLshiftToMulOp}
	_ReduceBitRshiftToMulOpAction                                     = &_Action{_ReduceAction, 0, _ReduceBitRshiftToMulOp}
	_ReduceMulExprToAddExprAction                                     = &_Action{_ReduceAction, 0, _ReduceMulExprToAddExpr}
	_ReduceBinaryAddExprToAddExprAction                               = &_Action{_ReduceAction, 0, _ReduceBinaryAddExprToAddExpr}
	_ReduceToBinaryAddExprAction                                      = &_Action{_ReduceAction, 0, _ReduceToBinaryAddExpr}
	_ReduceAddToAddOpAction                                           = &_Action{_ReduceAction, 0, _ReduceAddToAddOp}
	_ReduceSubToAddOpAction                                           = &_Action{_ReduceAction, 0, _ReduceSubToAddOp}
	_ReduceBitOrToAddOpAction                                         = &_Action{_ReduceAction, 0, _ReduceBitOrToAddOp}
	_ReduceBitXorToAddOpAction                                        = &_Action{_ReduceAction, 0, _ReduceBitXorToAddOp}
	_ReduceAddExprToCmpExprAction                                     = &_Action{_ReduceAction, 0, _ReduceAddExprToCmpExpr}
	_ReduceBinaryCmpExprToCmpExprAction                               = &_Action{_ReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}
	_ReduceToBinaryCmpExprAction                                      = &_Action{_ReduceAction, 0, _ReduceToBinaryCmpExpr}
	_ReduceEqualToCmpOpAction                                         = &_Action{_ReduceAction, 0, _ReduceEqualToCmpOp}
	_ReduceNotEqualToCmpOpAction                                      = &_Action{_ReduceAction, 0, _ReduceNotEqualToCmpOp}
	_ReduceLessToCmpOpAction                                          = &_Action{_ReduceAction, 0, _ReduceLessToCmpOp}
	_ReduceLessOrEqualToCmpOpAction                                   = &_Action{_ReduceAction, 0, _ReduceLessOrEqualToCmpOp}
	_ReduceGreaterToCmpOpAction                                       = &_Action{_ReduceAction, 0, _ReduceGreaterToCmpOp}
	_ReduceGreaterOrEqualToCmpOpAction                                = &_Action{_ReduceAction, 0, _ReduceGreaterOrEqualToCmpOp}
	_ReduceCmpExprToAndExprAction                                     = &_Action{_ReduceAction, 0, _ReduceCmpExprToAndExpr}
	_ReduceBinaryAndExprToAndExprAction                               = &_Action{_ReduceAction, 0, _ReduceBinaryAndExprToAndExpr}
	_ReduceToBinaryAndExprAction                                      = &_Action{_ReduceAction, 0, _ReduceToBinaryAndExpr}
	_ReduceAndExprToOrExprAction                                      = &_Action{_ReduceAction, 0, _ReduceAndExprToOrExpr}
	_ReduceBinaryOrExprToOrExprAction                                 = &_Action{_ReduceAction, 0, _ReduceBinaryOrExprToOrExpr}
	_ReduceToBinaryOrExprAction                                       = &_Action{_ReduceAction, 0, _ReduceToBinaryOrExpr}
	_ReduceExplicitStructDefToInitializableTypeAction                 = &_Action{_ReduceAction, 0, _ReduceExplicitStructDefToInitializableType}
	_ReduceSliceTypeToInitializableTypeAction                         = &_Action{_ReduceAction, 0, _ReduceSliceTypeToInitializableType}
	_ReduceArrayTypeToInitializableTypeAction                         = &_Action{_ReduceAction, 0, _ReduceArrayTypeToInitializableType}
	_ReduceMapTypeToInitializableTypeAction                           = &_Action{_ReduceAction, 0, _ReduceMapTypeToInitializableType}
	_ReduceToSliceTypeAction                                          = &_Action{_ReduceAction, 0, _ReduceToSliceType}
	_ReduceToArrayTypeAction                                          = &_Action{_ReduceAction, 0, _ReduceToArrayType}
	_ReduceToMapTypeAction                                            = &_Action{_ReduceAction, 0, _ReduceToMapType}
	_ReduceInitializableTypeToAtomTypeAction                          = &_Action{_ReduceAction, 0, _ReduceInitializableTypeToAtomType}
	_ReduceNamedToAtomTypeAction                                      = &_Action{_ReduceAction, 0, _ReduceNamedToAtomType}
	_ReduceExternNamedToAtomTypeAction                                = &_Action{_ReduceAction, 0, _ReduceExternNamedToAtomType}
	_ReduceInferredToAtomTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceInferredToAtomType}
	_ReduceImplicitStructDefToAtomTypeAction                          = &_Action{_ReduceAction, 0, _ReduceImplicitStructDefToAtomType}
	_ReduceExplicitEnumDefToAtomTypeAction                            = &_Action{_ReduceAction, 0, _ReduceExplicitEnumDefToAtomType}
	_ReduceImplicitEnumDefToAtomTypeAction                            = &_Action{_ReduceAction, 0, _ReduceImplicitEnumDefToAtomType}
	_ReduceTraitDefToAtomTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceTraitDefToAtomType}
	_ReduceFuncTypeToAtomTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceFuncTypeToAtomType}
	_ReduceParseErrorTypeToAtomTypeAction                             = &_Action{_ReduceAction, 0, _ReduceParseErrorTypeToAtomType}
	_ReduceToParseErrorTypeAction                                     = &_Action{_ReduceAction, 0, _ReduceToParseErrorType}
	_ReduceAtomTypeToReturnableTypeAction                             = &_Action{_ReduceAction, 0, _ReduceAtomTypeToReturnableType}
	_ReducePrefixedTypeToReturnableTypeAction                         = &_Action{_ReduceAction, 0, _ReducePrefixedTypeToReturnableType}
	_ReduceToPrefixedTypeAction                                       = &_Action{_ReduceAction, 0, _ReduceToPrefixedType}
	_ReduceQuestionToPrefixTypeOpAction                               = &_Action{_ReduceAction, 0, _ReduceQuestionToPrefixTypeOp}
	_ReduceExclaimToPrefixTypeOpAction                                = &_Action{_ReduceAction, 0, _ReduceExclaimToPrefixTypeOp}
	_ReduceBitAndToPrefixTypeOpAction                                 = &_Action{_ReduceAction, 0, _ReduceBitAndToPrefixTypeOp}
	_ReduceBitNegToPrefixTypeOpAction                                 = &_Action{_ReduceAction, 0, _ReduceBitNegToPrefixTypeOp}
	_ReduceTildeTildeToPrefixTypeOpAction                             = &_Action{_ReduceAction, 0, _ReduceTildeTildeToPrefixTypeOp}
	_ReduceReturnableTypeToValueTypeAction                            = &_Action{_ReduceAction, 0, _ReduceReturnableTypeToValueType}
	_ReduceTraitOpTypeToValueTypeAction                               = &_Action{_ReduceAction, 0, _ReduceTraitOpTypeToValueType}
	_ReduceToTraitOpTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceToTraitOpType}
	_ReduceMulToTraitOpAction                                         = &_Action{_ReduceAction, 0, _ReduceMulToTraitOp}
	_ReduceAddToTraitOpAction                                         = &_Action{_ReduceAction, 0, _ReduceAddToTraitOp}
	_ReduceSubToTraitOpAction                                         = &_Action{_ReduceAction, 0, _ReduceSubToTraitOp}
	_ReduceDefinitionToTypeDefAction                                  = &_Action{_ReduceAction, 0, _ReduceDefinitionToTypeDef}
	_ReduceConstrainedDefToTypeDefAction                              = &_Action{_ReduceAction, 0, _ReduceConstrainedDefToTypeDef}
	_ReduceAliasToTypeDefAction                                       = &_Action{_ReduceAction, 0, _ReduceAliasToTypeDef}
	_ReduceUnconstrainedToGenericParameterDefAction                   = &_Action{_ReduceAction, 0, _ReduceUnconstrainedToGenericParameterDef}
	_ReduceConstrainedToGenericParameterDefAction                     = &_Action{_ReduceAction, 0, _ReduceConstrainedToGenericParameterDef}
	_ReduceAddToProperGenericParameterDefsAction                      = &_Action{_ReduceAction, 0, _ReduceAddToProperGenericParameterDefs}
	_ReduceGenericParameterDefToProperGenericParameterDefsAction      = &_Action{_ReduceAction, 0, _ReduceGenericParameterDefToProperGenericParameterDefs}
	_ReduceProperGenericParameterDefsToGenericParameterDefsAction     = &_Action{_ReduceAction, 0, _ReduceProperGenericParameterDefsToGenericParameterDefs}
	_ReduceImproperToGenericParameterDefsAction                       = &_Action{_ReduceAction, 0, _ReduceImproperToGenericParameterDefs}
	_ReduceNilToGenericParameterDefsAction                            = &_Action{_ReduceAction, 0, _ReduceNilToGenericParameterDefs}
	_ReduceGenericToOptionalGenericParametersAction                   = &_Action{_ReduceAction, 0, _ReduceGenericToOptionalGenericParameters}
	_ReduceNilToOptionalGenericParametersAction                       = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericParameters}
	_ReduceExplicitToFieldDefAction                                   = &_Action{_ReduceAction, 0, _ReduceExplicitToFieldDef}
	_ReduceImplicitToFieldDefAction                                   = &_Action{_ReduceAction, 0, _ReduceImplicitToFieldDef}
	_ReduceUnsafeStatementToFieldDefAction                            = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToFieldDef}
	_ReduceAddToProperImplicitFieldDefsAction                         = &_Action{_ReduceAction, 0, _ReduceAddToProperImplicitFieldDefs}
	_ReduceFieldDefToProperImplicitFieldDefsAction                    = &_Action{_ReduceAction, 0, _ReduceFieldDefToProperImplicitFieldDefs}
	_ReduceProperImplicitFieldDefsToImplicitFieldDefsAction           = &_Action{_ReduceAction, 0, _ReduceProperImplicitFieldDefsToImplicitFieldDefs}
	_ReduceImproperToImplicitFieldDefsAction                          = &_Action{_ReduceAction, 0, _ReduceImproperToImplicitFieldDefs}
	_ReduceNilToImplicitFieldDefsAction                               = &_Action{_ReduceAction, 0, _ReduceNilToImplicitFieldDefs}
	_ReduceToImplicitStructDefAction                                  = &_Action{_ReduceAction, 0, _ReduceToImplicitStructDef}
	_ReduceAddImplicitToProperExplicitFieldDefsAction                 = &_Action{_ReduceAction, 0, _ReduceAddImplicitToProperExplicitFieldDefs}
	_ReduceAddExplicitToProperExplicitFieldDefsAction                 = &_Action{_ReduceAction, 0, _ReduceAddExplicitToProperExplicitFieldDefs}
	_ReduceFieldDefToProperExplicitFieldDefsAction                    = &_Action{_ReduceAction, 0, _ReduceFieldDefToProperExplicitFieldDefs}
	_ReduceProperExplicitFieldDefsToExplicitFieldDefsAction           = &_Action{_ReduceAction, 0, _ReduceProperExplicitFieldDefsToExplicitFieldDefs}
	_ReduceImproperImplicitToExplicitFieldDefsAction                  = &_Action{_ReduceAction, 0, _ReduceImproperImplicitToExplicitFieldDefs}
	_ReduceImproperExplicitToExplicitFieldDefsAction                  = &_Action{_ReduceAction, 0, _ReduceImproperExplicitToExplicitFieldDefs}
	_ReduceNilToExplicitFieldDefsAction                               = &_Action{_ReduceAction, 0, _ReduceNilToExplicitFieldDefs}
	_ReduceToExplicitStructDefAction                                  = &_Action{_ReduceAction, 0, _ReduceToExplicitStructDef}
	_ReduceFieldDefToEnumValueDefAction                               = &_Action{_ReduceAction, 0, _ReduceFieldDefToEnumValueDef}
	_ReduceDefaultToEnumValueDefAction                                = &_Action{_ReduceAction, 0, _ReduceDefaultToEnumValueDef}
	_ReducePairToImplicitEnumValueDefsAction                          = &_Action{_ReduceAction, 0, _ReducePairToImplicitEnumValueDefs}
	_ReduceAddToImplicitEnumValueDefsAction                           = &_Action{_ReduceAction, 0, _ReduceAddToImplicitEnumValueDefs}
	_ReduceToImplicitEnumDefAction                                    = &_Action{_ReduceAction, 0, _ReduceToImplicitEnumDef}
	_ReduceExplicitPairToExplicitEnumValueDefsAction                  = &_Action{_ReduceAction, 0, _ReduceExplicitPairToExplicitEnumValueDefs}
	_ReduceImplicitPairToExplicitEnumValueDefsAction                  = &_Action{_ReduceAction, 0, _ReduceImplicitPairToExplicitEnumValueDefs}
	_ReduceExplicitAddToExplicitEnumValueDefsAction                   = &_Action{_ReduceAction, 0, _ReduceExplicitAddToExplicitEnumValueDefs}
	_ReduceImplicitAddToExplicitEnumValueDefsAction                   = &_Action{_ReduceAction, 0, _ReduceImplicitAddToExplicitEnumValueDefs}
	_ReduceToExplicitEnumDefAction                                    = &_Action{_ReduceAction, 0, _ReduceToExplicitEnumDef}
	_ReduceFieldDefToTraitPropertyAction                              = &_Action{_ReduceAction, 0, _ReduceFieldDefToTraitProperty}
	_ReduceMethodSignatureToTraitPropertyAction                       = &_Action{_ReduceAction, 0, _ReduceMethodSignatureToTraitProperty}
	_ReduceImplicitToProperTraitPropertiesAction                      = &_Action{_ReduceAction, 0, _ReduceImplicitToProperTraitProperties}
	_ReduceExplicitToProperTraitPropertiesAction                      = &_Action{_ReduceAction, 0, _ReduceExplicitToProperTraitProperties}
	_ReduceTraitPropertyToProperTraitPropertiesAction                 = &_Action{_ReduceAction, 0, _ReduceTraitPropertyToProperTraitProperties}
	_ReduceProperTraitPropertiesToTraitPropertiesAction               = &_Action{_ReduceAction, 0, _ReduceProperTraitPropertiesToTraitProperties}
	_ReduceImproperImplicitToTraitPropertiesAction                    = &_Action{_ReduceAction, 0, _ReduceImproperImplicitToTraitProperties}
	_ReduceImproperExplicitToTraitPropertiesAction                    = &_Action{_ReduceAction, 0, _ReduceImproperExplicitToTraitProperties}
	_ReduceNilToTraitPropertiesAction                                 = &_Action{_ReduceAction, 0, _ReduceNilToTraitProperties}
	_ReduceToTraitDefAction                                           = &_Action{_ReduceAction, 0, _ReduceToTraitDef}
	_ReduceReturnableTypeToReturnTypeAction                           = &_Action{_ReduceAction, 0, _ReduceReturnableTypeToReturnType}
	_ReduceNilToReturnTypeAction                                      = &_Action{_ReduceAction, 0, _ReduceNilToReturnType}
	_ReduceArgToParameterDeclAction                                   = &_Action{_ReduceAction, 0, _ReduceArgToParameterDecl}
	_ReduceVarargToParameterDeclAction                                = &_Action{_ReduceAction, 0, _ReduceVarargToParameterDecl}
	_ReduceUnamedToParameterDeclAction                                = &_Action{_ReduceAction, 0, _ReduceUnamedToParameterDecl}
	_ReduceUnnamedVarargToParameterDeclAction                         = &_Action{_ReduceAction, 0, _ReduceUnnamedVarargToParameterDecl}
	_ReduceAddToProperParameterDeclsAction                            = &_Action{_ReduceAction, 0, _ReduceAddToProperParameterDecls}
	_ReduceParameterDeclToProperParameterDeclsAction                  = &_Action{_ReduceAction, 0, _ReduceParameterDeclToProperParameterDecls}
	_ReduceProperParameterDeclsToParameterDeclsAction                 = &_Action{_ReduceAction, 0, _ReduceProperParameterDeclsToParameterDecls}
	_ReduceImproperToParameterDeclsAction                             = &_Action{_ReduceAction, 0, _ReduceImproperToParameterDecls}
	_ReduceNilToParameterDeclsAction                                  = &_Action{_ReduceAction, 0, _ReduceNilToParameterDecls}
	_ReduceToFuncTypeAction                                           = &_Action{_ReduceAction, 0, _ReduceToFuncType}
	_ReduceToMethodSignatureAction                                    = &_Action{_ReduceAction, 0, _ReduceToMethodSignature}
	_ReduceInferredRefArgToParameterDefAction                         = &_Action{_ReduceAction, 0, _ReduceInferredRefArgToParameterDef}
	_ReduceInferredRefVarargToParameterDefAction                      = &_Action{_ReduceAction, 0, _ReduceInferredRefVarargToParameterDef}
	_ReduceArgToParameterDefAction                                    = &_Action{_ReduceAction, 0, _ReduceArgToParameterDef}
	_ReduceVarargToParameterDefAction                                 = &_Action{_ReduceAction, 0, _ReduceVarargToParameterDef}
	_ReduceAddToProperParameterDefsAction                             = &_Action{_ReduceAction, 0, _ReduceAddToProperParameterDefs}
	_ReduceParameterDefToProperParameterDefsAction                    = &_Action{_ReduceAction, 0, _ReduceParameterDefToProperParameterDefs}
	_ReduceProperParameterDefsToParameterDefsAction                   = &_Action{_ReduceAction, 0, _ReduceProperParameterDefsToParameterDefs}
	_ReduceImproperToParameterDefsAction                              = &_Action{_ReduceAction, 0, _ReduceImproperToParameterDefs}
	_ReduceNilToParameterDefsAction                                   = &_Action{_ReduceAction, 0, _ReduceNilToParameterDefs}
	_ReduceFuncDefToNamedFuncDefAction                                = &_Action{_ReduceAction, 0, _ReduceFuncDefToNamedFuncDef}
	_ReduceMethodDefToNamedFuncDefAction                              = &_Action{_ReduceAction, 0, _ReduceMethodDefToNamedFuncDef}
	_ReduceAliasToNamedFuncDefAction                                  = &_Action{_ReduceAction, 0, _ReduceAliasToNamedFuncDef}
	_ReduceToAnonymousFuncExprAction                                  = &_Action{_ReduceAction, 0, _ReduceToAnonymousFuncExpr}
	_ReduceNoSpecToPackageDefAction                                   = &_Action{_ReduceAction, 0, _ReduceNoSpecToPackageDef}
	_ReduceWithSpecToPackageDefAction                                 = &_Action{_ReduceAction, 0, _ReduceWithSpecToPackageDef}
)

var _ActionTable = _ActionTableType{
	{_State6, _EndMarker}:                                &_Action{_AcceptAction, 0, 0},
	{_State7, _EndMarker}:                                &_Action{_AcceptAction, 0, 0},
	{_State8, _EndMarker}:                                &_Action{_AcceptAction, 0, 0},
	{_State9, _EndMarker}:                                &_Action{_AcceptAction, 0, 0},
	{_State10, _EndMarker}:                               &_Action{_AcceptAction, 0, 0},
	{_State1, SourceType}:                                _GotoState6Action,
	{_State1, DefinitionsType}:                           _GotoState11Action,
	{_State2, PackageToken}:                              _GotoState12Action,
	{_State2, PackageDefType}:                            _GotoState7Action,
	{_State3, TypeToken}:                                 _GotoState13Action,
	{_State3, TypeDefType}:                               _GotoState8Action,
	{_State4, FuncToken}:                                 _GotoState14Action,
	{_State4, NamedFuncDefType}:                          _GotoState9Action,
	{_State5, IntegerLiteralToken}:                       _GotoState22Action,
	{_State5, FloatLiteralToken}:                         _GotoState18Action,
	{_State5, RuneLiteralToken}:                          _GotoState30Action,
	{_State5, StringLiteralToken}:                        _GotoState31Action,
	{_State5, IdentifierToken}:                           _GotoState21Action,
	{_State5, TrueToken}:                                 _GotoState34Action,
	{_State5, FalseToken}:                                _GotoState17Action,
	{_State5, StructToken}:                               _GotoState32Action,
	{_State5, FuncToken}:                                 _GotoState19Action,
	{_State5, VarToken}:                                  _GotoState35Action,
	{_State5, LetToken}:                                  _GotoState25Action,
	{_State5, NotToken}:                                  _GotoState28Action,
	{_State5, LabelDeclToken}:                            _GotoState23Action,
	{_State5, LparenToken}:                               _GotoState26Action,
	{_State5, LbracketToken}:                             _GotoState24Action,
	{_State5, SubToken}:                                  _GotoState33Action,
	{_State5, MulToken}:                                  _GotoState27Action,
	{_State5, BitNegToken}:                               _GotoState16Action,
	{_State5, BitAndToken}:                               _GotoState15Action,
	{_State5, GreaterToken}:                              _GotoState20Action,
	{_State5, ParseErrorToken}:                           _GotoState29Action,
	{_State5, VarDeclPatternType}:                        _GotoState70Action,
	{_State5, VarOrLetType}:                              _GotoState71Action,
	{_State5, ExpressionType}:                            _GotoState10Action,
	{_State5, OptionalLabelDeclType}:                     _GotoState60Action,
	{_State5, SequenceExprType}:                          _GotoState68Action,
	{_State5, CallExprType}:                              _GotoState49Action,
	{_State5, AtomExprType}:                              _GotoState42Action,
	{_State5, ParseErrorExprType}:                        _GotoState62Action,
	{_State5, LiteralExprType}:                           _GotoState57Action,
	{_State5, IdentifierExprType}:                        _GotoState52Action,
	{_State5, BlockExprType}:                             _GotoState48Action,
	{_State5, InitializeExprType}:                        _GotoState56Action,
	{_State5, ImplicitStructExprType}:                    _GotoState53Action,
	{_State5, AccessibleExprType}:                        _GotoState37Action,
	{_State5, AccessExprType}:                            _GotoState36Action,
	{_State5, IndexExprType}:                             _GotoState54Action,
	{_State5, PostfixableExprType}:                       _GotoState64Action,
	{_State5, PostfixUnaryExprType}:                      _GotoState63Action,
	{_State5, PrefixableExprType}:                        _GotoState67Action,
	{_State5, PrefixUnaryExprType}:                       _GotoState65Action,
	{_State5, PrefixUnaryOpType}:                         _GotoState66Action,
	{_State5, MulExprType}:                               _GotoState59Action,
	{_State5, BinaryMulExprType}:                         _GotoState46Action,
	{_State5, AddExprType}:                               _GotoState38Action,
	{_State5, BinaryAddExprType}:                         _GotoState43Action,
	{_State5, CmpExprType}:                               _GotoState50Action,
	{_State5, BinaryCmpExprType}:                         _GotoState45Action,
	{_State5, AndExprType}:                               _GotoState39Action,
	{_State5, BinaryAndExprType}:                         _GotoState44Action,
	{_State5, OrExprType}:                                _GotoState61Action,
	{_State5, BinaryOrExprType}:                          _GotoState47Action,
	{_State5, InitializableTypeType}:                     _GotoState55Action,
	{_State5, SliceTypeType}:                             _GotoState69Action,
	{_State5, ArrayTypeType}:                             _GotoState41Action,
	{_State5, MapTypeType}:                               _GotoState58Action,
	{_State5, ExplicitStructDefType}:                     _GotoState51Action,
	{_State5, AnonymousFuncExprType}:                     _GotoState40Action,
	{_State11, CommentGroupsToken}:                       _GotoState72Action,
	{_State11, PackageToken}:                             _GotoState12Action,
	{_State11, TypeToken}:                                _GotoState13Action,
	{_State11, FuncToken}:                                _GotoState14Action,
	{_State11, VarToken}:                                 _GotoState35Action,
	{_State11, LetToken}:                                 _GotoState25Action,
	{_State11, LbraceToken}:                              _GotoState73Action,
	{_State11, DefinitionType}:                           _GotoState74Action,
	{_State11, StatementBlockType}:                       _GotoState77Action,
	{_State11, VarDeclPatternType}:                       _GotoState79Action,
	{_State11, VarOrLetType}:                             _GotoState71Action,
	{_State11, TypeDefType}:                              _GotoState78Action,
	{_State11, NamedFuncDefType}:                         _GotoState75Action,
	{_State11, PackageDefType}:                           _GotoState76Action,
	{_State12, LbraceToken}:                              _GotoState73Action,
	{_State12, StatementBlockType}:                       _GotoState80Action,
	{_State13, IdentifierToken}:                          _GotoState81Action,
	{_State14, IdentifierToken}:                          _GotoState82Action,
	{_State14, LparenToken}:                              _GotoState83Action,
	{_State19, LparenToken}:                              _GotoState84Action,
	{_State20, IntegerLiteralToken}:                      _GotoState22Action,
	{_State20, FloatLiteralToken}:                        _GotoState18Action,
	{_State20, RuneLiteralToken}:                         _GotoState30Action,
	{_State20, StringLiteralToken}:                       _GotoState31Action,
	{_State20, IdentifierToken}:                          _GotoState21Action,
	{_State20, TrueToken}:                                _GotoState34Action,
	{_State20, FalseToken}:                               _GotoState17Action,
	{_State20, StructToken}:                              _GotoState32Action,
	{_State20, FuncToken}:                                _GotoState19Action,
	{_State20, VarToken}:                                 _GotoState35Action,
	{_State20, LetToken}:                                 _GotoState25Action,
	{_State20, NotToken}:                                 _GotoState28Action,
	{_State20, LabelDeclToken}:                           _GotoState23Action,
	{_State20, LparenToken}:                              _GotoState26Action,
	{_State20, LbracketToken}:                            _GotoState24Action,
	{_State20, SubToken}:                                 _GotoState33Action,
	{_State20, MulToken}:                                 _GotoState27Action,
	{_State20, BitNegToken}:                              _GotoState16Action,
	{_State20, BitAndToken}:                              _GotoState15Action,
	{_State20, GreaterToken}:                             _GotoState20Action,
	{_State20, ParseErrorToken}:                          _GotoState29Action,
	{_State20, VarDeclPatternType}:                       _GotoState70Action,
	{_State20, VarOrLetType}:                             _GotoState71Action,
	{_State20, OptionalLabelDeclType}:                    _GotoState85Action,
	{_State20, SequenceExprType}:                         _GotoState86Action,
	{_State20, CallExprType}:                             _GotoState49Action,
	{_State20, AtomExprType}:                             _GotoState42Action,
	{_State20, ParseErrorExprType}:                       _GotoState62Action,
	{_State20, LiteralExprType}:                          _GotoState57Action,
	{_State20, IdentifierExprType}:                       _GotoState52Action,
	{_State20, BlockExprType}:                            _GotoState48Action,
	{_State20, InitializeExprType}:                       _GotoState56Action,
	{_State20, ImplicitStructExprType}:                   _GotoState53Action,
	{_State20, AccessibleExprType}:                       _GotoState37Action,
	{_State20, AccessExprType}:                           _GotoState36Action,
	{_State20, IndexExprType}:                            _GotoState54Action,
	{_State20, PostfixableExprType}:                      _GotoState64Action,
	{_State20, PostfixUnaryExprType}:                     _GotoState63Action,
	{_State20, PrefixableExprType}:                       _GotoState67Action,
	{_State20, PrefixUnaryExprType}:                      _GotoState65Action,
	{_State20, PrefixUnaryOpType}:                        _GotoState66Action,
	{_State20, MulExprType}:                              _GotoState59Action,
	{_State20, BinaryMulExprType}:                        _GotoState46Action,
	{_State20, AddExprType}:                              _GotoState38Action,
	{_State20, BinaryAddExprType}:                        _GotoState43Action,
	{_State20, CmpExprType}:                              _GotoState50Action,
	{_State20, BinaryCmpExprType}:                        _GotoState45Action,
	{_State20, AndExprType}:                              _GotoState39Action,
	{_State20, BinaryAndExprType}:                        _GotoState44Action,
	{_State20, OrExprType}:                               _GotoState61Action,
	{_State20, BinaryOrExprType}:                         _GotoState47Action,
	{_State20, InitializableTypeType}:                    _GotoState55Action,
	{_State20, SliceTypeType}:                            _GotoState69Action,
	{_State20, ArrayTypeType}:                            _GotoState41Action,
	{_State20, MapTypeType}:                              _GotoState58Action,
	{_State20, ExplicitStructDefType}:                    _GotoState51Action,
	{_State20, AnonymousFuncExprType}:                    _GotoState40Action,
	{_State24, IdentifierToken}:                          _GotoState93Action,
	{_State24, StructToken}:                              _GotoState32Action,
	{_State24, EnumToken}:                                _GotoState90Action,
	{_State24, TraitToken}:                               _GotoState98Action,
	{_State24, FuncToken}:                                _GotoState92Action,
	{_State24, LparenToken}:                              _GotoState94Action,
	{_State24, LbracketToken}:                            _GotoState24Action,
	{_State24, DotToken}:                                 _GotoState89Action,
	{_State24, QuestionToken}:                            _GotoState96Action,
	{_State24, ExclaimToken}:                             _GotoState91Action,
	{_State24, TildeTildeToken}:                          _GotoState97Action,
	{_State24, BitNegToken}:                              _GotoState88Action,
	{_State24, BitAndToken}:                              _GotoState87Action,
	{_State24, ParseErrorToken}:                          _GotoState95Action,
	{_State24, InitializableTypeType}:                    _GotoState104Action,
	{_State24, SliceTypeType}:                            _GotoState69Action,
	{_State24, ArrayTypeType}:                            _GotoState41Action,
	{_State24, MapTypeType}:                              _GotoState58Action,
	{_State24, AtomTypeType}:                             _GotoState99Action,
	{_State24, ParseErrorTypeType}:                       _GotoState105Action,
	{_State24, ReturnableTypeType}:                       _GotoState108Action,
	{_State24, PrefixedTypeType}:                         _GotoState107Action,
	{_State24, PrefixTypeOpType}:                         _GotoState106Action,
	{_State24, ValueTypeType}:                            _GotoState111Action,
	{_State24, TraitOpTypeType}:                          _GotoState110Action,
	{_State24, ImplicitStructDefType}:                    _GotoState103Action,
	{_State24, ExplicitStructDefType}:                    _GotoState51Action,
	{_State24, ImplicitEnumDefType}:                      _GotoState102Action,
	{_State24, ExplicitEnumDefType}:                      _GotoState100Action,
	{_State24, TraitDefType}:                             _GotoState109Action,
	{_State24, FuncTypeType}:                             _GotoState101Action,
	{_State26, IntegerLiteralToken}:                      _GotoState22Action,
	{_State26, FloatLiteralToken}:                        _GotoState18Action,
	{_State26, RuneLiteralToken}:                         _GotoState30Action,
	{_State26, StringLiteralToken}:                       _GotoState31Action,
	{_State26, IdentifierToken}:                          _GotoState114Action,
	{_State26, TrueToken}:                                _GotoState34Action,
	{_State26, FalseToken}:                               _GotoState17Action,
	{_State26, StructToken}:                              _GotoState32Action,
	{_State26, FuncToken}:                                _GotoState19Action,
	{_State26, VarToken}:                                 _GotoState35Action,
	{_State26, LetToken}:                                 _GotoState25Action,
	{_State26, NotToken}:                                 _GotoState28Action,
	{_State26, LabelDeclToken}:                           _GotoState23Action,
	{_State26, LparenToken}:                              _GotoState26Action,
	{_State26, LbracketToken}:                            _GotoState24Action,
	{_State26, ColonToken}:                               _GotoState112Action,
	{_State26, EllipsisToken}:                            _GotoState113Action,
	{_State26, SubToken}:                                 _GotoState33Action,
	{_State26, MulToken}:                                 _GotoState27Action,
	{_State26, BitNegToken}:                              _GotoState16Action,
	{_State26, BitAndToken}:                              _GotoState15Action,
	{_State26, GreaterToken}:                             _GotoState20Action,
	{_State26, ParseErrorToken}:                          _GotoState29Action,
	{_State26, VarDeclPatternType}:                       _GotoState70Action,
	{_State26, VarOrLetType}:                             _GotoState71Action,
	{_State26, ExpressionType}:                           _GotoState118Action,
	{_State26, OptionalLabelDeclType}:                    _GotoState60Action,
	{_State26, SequenceExprType}:                         _GotoState68Action,
	{_State26, CallExprType}:                             _GotoState49Action,
	{_State26, ProperArgumentsType}:                      _GotoState119Action,
	{_State26, ArgumentsType}:                            _GotoState116Action,
	{_State26, ArgumentType}:                             _GotoState115Action,
	{_State26, ColonExprType}:                            _GotoState117Action,
	{_State26, AtomExprType}:                             _GotoState42Action,
	{_State26, ParseErrorExprType}:                       _GotoState62Action,
	{_State26, LiteralExprType}:                          _GotoState57Action,
	{_State26, IdentifierExprType}:                       _GotoState52Action,
	{_State26, BlockExprType}:                            _GotoState48Action,
	{_State26, InitializeExprType}:                       _GotoState56Action,
	{_State26, ImplicitStructExprType}:                   _GotoState53Action,
	{_State26, AccessibleExprType}:                       _GotoState37Action,
	{_State26, AccessExprType}:                           _GotoState36Action,
	{_State26, IndexExprType}:                            _GotoState54Action,
	{_State26, PostfixableExprType}:                      _GotoState64Action,
	{_State26, PostfixUnaryExprType}:                     _GotoState63Action,
	{_State26, PrefixableExprType}:                       _GotoState67Action,
	{_State26, PrefixUnaryExprType}:                      _GotoState65Action,
	{_State26, PrefixUnaryOpType}:                        _GotoState66Action,
	{_State26, MulExprType}:                              _GotoState59Action,
	{_State26, BinaryMulExprType}:                        _GotoState46Action,
	{_State26, AddExprType}:                              _GotoState38Action,
	{_State26, BinaryAddExprType}:                        _GotoState43Action,
	{_State26, CmpExprType}:                              _GotoState50Action,
	{_State26, BinaryCmpExprType}:                        _GotoState45Action,
	{_State26, AndExprType}:                              _GotoState39Action,
	{_State26, BinaryAndExprType}:                        _GotoState44Action,
	{_State26, OrExprType}:                               _GotoState61Action,
	{_State26, BinaryOrExprType}:                         _GotoState47Action,
	{_State26, InitializableTypeType}:                    _GotoState55Action,
	{_State26, SliceTypeType}:                            _GotoState69Action,
	{_State26, ArrayTypeType}:                            _GotoState41Action,
	{_State26, MapTypeType}:                              _GotoState58Action,
	{_State26, ExplicitStructDefType}:                    _GotoState51Action,
	{_State26, AnonymousFuncExprType}:                    _GotoState40Action,
	{_State32, LparenToken}:                              _GotoState120Action,
	{_State37, LbracketToken}:                            _GotoState123Action,
	{_State37, DotToken}:                                 _GotoState122Action,
	{_State37, QuestionToken}:                            _GotoState124Action,
	{_State37, DollarLbracketToken}:                      _GotoState121Action,
	{_State37, OptionalGenericBindingType}:               _GotoState125Action,
	{_State38, AddToken}:                                 _GotoState126Action,
	{_State38, SubToken}:                                 _GotoState129Action,
	{_State38, BitXorToken}:                              _GotoState128Action,
	{_State38, BitOrToken}:                               _GotoState127Action,
	{_State38, AddOpType}:                                _GotoState130Action,
	{_State39, AndToken}:                                 _GotoState131Action,
	{_State50, EqualToken}:                               _GotoState132Action,
	{_State50, NotEqualToken}:                            _GotoState137Action,
	{_State50, LessToken}:                                _GotoState135Action,
	{_State50, LessOrEqualToken}:                         _GotoState136Action,
	{_State50, GreaterToken}:                             _GotoState133Action,
	{_State50, GreaterOrEqualToken}:                      _GotoState134Action,
	{_State50, CmpOpType}:                                _GotoState138Action,
	{_State55, LparenToken}:                              _GotoState139Action,
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
	{_State60, LbraceToken}:                              _GotoState73Action,
	{_State60, StatementBlockType}:                       _GotoState153Action,
	{_State60, IfExprType}:                               _GotoState151Action,
	{_State60, SwitchExprType}:                           _GotoState154Action,
	{_State60, LoopExprType}:                             _GotoState152Action,
	{_State61, OrToken}:                                  _GotoState155Action,
	{_State66, IntegerLiteralToken}:                      _GotoState22Action,
	{_State66, FloatLiteralToken}:                        _GotoState18Action,
	{_State66, RuneLiteralToken}:                         _GotoState30Action,
	{_State66, StringLiteralToken}:                       _GotoState31Action,
	{_State66, IdentifierToken}:                          _GotoState21Action,
	{_State66, TrueToken}:                                _GotoState34Action,
	{_State66, FalseToken}:                               _GotoState17Action,
	{_State66, StructToken}:                              _GotoState32Action,
	{_State66, FuncToken}:                                _GotoState19Action,
	{_State66, NotToken}:                                 _GotoState28Action,
	{_State66, LabelDeclToken}:                           _GotoState23Action,
	{_State66, LparenToken}:                              _GotoState26Action,
	{_State66, LbracketToken}:                            _GotoState24Action,
	{_State66, SubToken}:                                 _GotoState33Action,
	{_State66, MulToken}:                                 _GotoState27Action,
	{_State66, BitNegToken}:                              _GotoState16Action,
	{_State66, BitAndToken}:                              _GotoState15Action,
	{_State66, ParseErrorToken}:                          _GotoState29Action,
	{_State66, OptionalLabelDeclType}:                    _GotoState85Action,
	{_State66, CallExprType}:                             _GotoState49Action,
	{_State66, AtomExprType}:                             _GotoState42Action,
	{_State66, ParseErrorExprType}:                       _GotoState62Action,
	{_State66, LiteralExprType}:                          _GotoState57Action,
	{_State66, IdentifierExprType}:                       _GotoState52Action,
	{_State66, BlockExprType}:                            _GotoState48Action,
	{_State66, InitializeExprType}:                       _GotoState56Action,
	{_State66, ImplicitStructExprType}:                   _GotoState53Action,
	{_State66, AccessibleExprType}:                       _GotoState37Action,
	{_State66, AccessExprType}:                           _GotoState36Action,
	{_State66, IndexExprType}:                            _GotoState54Action,
	{_State66, PostfixableExprType}:                      _GotoState64Action,
	{_State66, PostfixUnaryExprType}:                     _GotoState63Action,
	{_State66, PrefixableExprType}:                       _GotoState156Action,
	{_State66, PrefixUnaryExprType}:                      _GotoState65Action,
	{_State66, PrefixUnaryOpType}:                        _GotoState66Action,
	{_State66, InitializableTypeType}:                    _GotoState55Action,
	{_State66, SliceTypeType}:                            _GotoState69Action,
	{_State66, ArrayTypeType}:                            _GotoState41Action,
	{_State66, MapTypeType}:                              _GotoState58Action,
	{_State66, ExplicitStructDefType}:                    _GotoState51Action,
	{_State66, AnonymousFuncExprType}:                    _GotoState40Action,
	{_State71, IdentifierToken}:                          _GotoState157Action,
	{_State71, LparenToken}:                              _GotoState158Action,
	{_State71, VarPatternType}:                           _GotoState160Action,
	{_State71, TuplePatternType}:                         _GotoState159Action,
	{_State73, IntegerLiteralToken}:                      _GotoState22Action,
	{_State73, FloatLiteralToken}:                        _GotoState18Action,
	{_State73, RuneLiteralToken}:                         _GotoState30Action,
	{_State73, StringLiteralToken}:                       _GotoState31Action,
	{_State73, IdentifierToken}:                          _GotoState21Action,
	{_State73, TrueToken}:                                _GotoState34Action,
	{_State73, FalseToken}:                               _GotoState17Action,
	{_State73, CaseToken}:                                _GotoState163Action,
	{_State73, DefaultToken}:                             _GotoState165Action,
	{_State73, ReturnToken}:                              _GotoState169Action,
	{_State73, BreakToken}:                               _GotoState162Action,
	{_State73, ContinueToken}:                            _GotoState164Action,
	{_State73, FallthroughToken}:                         _GotoState167Action,
	{_State73, ImportToken}:                              _GotoState168Action,
	{_State73, UnsafeToken}:                              _GotoState170Action,
	{_State73, StructToken}:                              _GotoState32Action,
	{_State73, FuncToken}:                                _GotoState19Action,
	{_State73, AsyncToken}:                               _GotoState161Action,
	{_State73, DeferToken}:                               _GotoState166Action,
	{_State73, VarToken}:                                 _GotoState35Action,
	{_State73, LetToken}:                                 _GotoState25Action,
	{_State73, NotToken}:                                 _GotoState28Action,
	{_State73, LabelDeclToken}:                           _GotoState23Action,
	{_State73, LparenToken}:                              _GotoState26Action,
	{_State73, LbracketToken}:                            _GotoState24Action,
	{_State73, SubToken}:                                 _GotoState33Action,
	{_State73, MulToken}:                                 _GotoState27Action,
	{_State73, BitNegToken}:                              _GotoState16Action,
	{_State73, BitAndToken}:                              _GotoState15Action,
	{_State73, GreaterToken}:                             _GotoState20Action,
	{_State73, ParseErrorToken}:                          _GotoState29Action,
	{_State73, ProperStatementsType}:                     _GotoState184Action,
	{_State73, StatementsType}:                           _GotoState188Action,
	{_State73, SimpleStatementType}:                      _GotoState186Action,
	{_State73, StatementType}:                            _GotoState187Action,
	{_State73, ExpressionOrImproperStructStatementType}:  _GotoState178Action,
	{_State73, ExpressionsType}:                          _GotoState179Action,
	{_State73, CallbackOpType}:                           _GotoState175Action,
	{_State73, CallbackOpStatementType}:                  _GotoState176Action,
	{_State73, UnsafeStatementType}:                      _GotoState190Action,
	{_State73, JumpStatementType}:                        _GotoState182Action,
	{_State73, JumpTypeType}:                             _GotoState183Action,
	{_State73, FallthroughStatementType}:                 _GotoState180Action,
	{_State73, AssignStatementType}:                      _GotoState173Action,
	{_State73, UnaryOpAssignStatementType}:               _GotoState189Action,
	{_State73, BinaryOpAssignStatementType}:              _GotoState174Action,
	{_State73, ImportStatementType}:                      _GotoState181Action,
	{_State73, VarDeclPatternType}:                       _GotoState70Action,
	{_State73, VarOrLetType}:                             _GotoState71Action,
	{_State73, AssignPatternType}:                        _GotoState172Action,
	{_State73, ExpressionType}:                           _GotoState177Action,
	{_State73, OptionalLabelDeclType}:                    _GotoState60Action,
	{_State73, SequenceExprType}:                         _GotoState185Action,
	{_State73, CallExprType}:                             _GotoState49Action,
	{_State73, AtomExprType}:                             _GotoState42Action,
	{_State73, ParseErrorExprType}:                       _GotoState62Action,
	{_State73, LiteralExprType}:                          _GotoState57Action,
	{_State73, IdentifierExprType}:                       _GotoState52Action,
	{_State73, BlockExprType}:                            _GotoState48Action,
	{_State73, InitializeExprType}:                       _GotoState56Action,
	{_State73, ImplicitStructExprType}:                   _GotoState53Action,
	{_State73, AccessibleExprType}:                       _GotoState171Action,
	{_State73, AccessExprType}:                           _GotoState36Action,
	{_State73, IndexExprType}:                            _GotoState54Action,
	{_State73, PostfixableExprType}:                      _GotoState64Action,
	{_State73, PostfixUnaryExprType}:                     _GotoState63Action,
	{_State73, PrefixableExprType}:                       _GotoState67Action,
	{_State73, PrefixUnaryExprType}:                      _GotoState65Action,
	{_State73, PrefixUnaryOpType}:                        _GotoState66Action,
	{_State73, MulExprType}:                              _GotoState59Action,
	{_State73, BinaryMulExprType}:                        _GotoState46Action,
	{_State73, AddExprType}:                              _GotoState38Action,
	{_State73, BinaryAddExprType}:                        _GotoState43Action,
	{_State73, CmpExprType}:                              _GotoState50Action,
	{_State73, BinaryCmpExprType}:                        _GotoState45Action,
	{_State73, AndExprType}:                              _GotoState39Action,
	{_State73, BinaryAndExprType}:                        _GotoState44Action,
	{_State73, OrExprType}:                               _GotoState61Action,
	{_State73, BinaryOrExprType}:                         _GotoState47Action,
	{_State73, InitializableTypeType}:                    _GotoState55Action,
	{_State73, SliceTypeType}:                            _GotoState69Action,
	{_State73, ArrayTypeType}:                            _GotoState41Action,
	{_State73, MapTypeType}:                              _GotoState58Action,
	{_State73, ExplicitStructDefType}:                    _GotoState51Action,
	{_State73, AnonymousFuncExprType}:                    _GotoState40Action,
	{_State74, NewlinesToken}:                            _GotoState191Action,
	{_State79, AssignToken}:                              _GotoState192Action,
	{_State81, DollarLbracketToken}:                      _GotoState194Action,
	{_State81, AssignToken}:                              _GotoState193Action,
	{_State81, OptionalGenericParametersType}:            _GotoState195Action,
	{_State82, DollarLbracketToken}:                      _GotoState194Action,
	{_State82, AssignToken}:                              _GotoState196Action,
	{_State82, OptionalGenericParametersType}:            _GotoState197Action,
	{_State83, IdentifierToken}:                          _GotoState198Action,
	{_State83, ParameterDefType}:                         _GotoState199Action,
	{_State84, IdentifierToken}:                          _GotoState198Action,
	{_State84, ParameterDefType}:                         _GotoState200Action,
	{_State84, ProperParameterDefsType}:                  _GotoState202Action,
	{_State84, ParameterDefsType}:                        _GotoState201Action,
	{_State85, LbraceToken}:                              _GotoState73Action,
	{_State85, StatementBlockType}:                       _GotoState153Action,
	{_State89, DollarLbracketToken}:                      _GotoState121Action,
	{_State89, OptionalGenericBindingType}:               _GotoState203Action,
	{_State90, LparenToken}:                              _GotoState204Action,
	{_State92, LparenToken}:                              _GotoState205Action,
	{_State93, DotToken}:                                 _GotoState206Action,
	{_State93, DollarLbracketToken}:                      _GotoState121Action,
	{_State93, OptionalGenericBindingType}:               _GotoState207Action,
	{_State94, IdentifierToken}:                          _GotoState208Action,
	{_State94, UnsafeToken}:                              _GotoState170Action,
	{_State94, StructToken}:                              _GotoState32Action,
	{_State94, EnumToken}:                                _GotoState90Action,
	{_State94, TraitToken}:                               _GotoState98Action,
	{_State94, FuncToken}:                                _GotoState92Action,
	{_State94, LparenToken}:                              _GotoState94Action,
	{_State94, LbracketToken}:                            _GotoState24Action,
	{_State94, DotToken}:                                 _GotoState89Action,
	{_State94, QuestionToken}:                            _GotoState96Action,
	{_State94, ExclaimToken}:                             _GotoState91Action,
	{_State94, TildeTildeToken}:                          _GotoState97Action,
	{_State94, BitNegToken}:                              _GotoState88Action,
	{_State94, BitAndToken}:                              _GotoState87Action,
	{_State94, ParseErrorToken}:                          _GotoState95Action,
	{_State94, UnsafeStatementType}:                      _GotoState214Action,
	{_State94, InitializableTypeType}:                    _GotoState104Action,
	{_State94, SliceTypeType}:                            _GotoState69Action,
	{_State94, ArrayTypeType}:                            _GotoState41Action,
	{_State94, MapTypeType}:                              _GotoState58Action,
	{_State94, AtomTypeType}:                             _GotoState99Action,
	{_State94, ParseErrorTypeType}:                       _GotoState105Action,
	{_State94, ReturnableTypeType}:                       _GotoState108Action,
	{_State94, PrefixedTypeType}:                         _GotoState107Action,
	{_State94, PrefixTypeOpType}:                         _GotoState106Action,
	{_State94, ValueTypeType}:                            _GotoState215Action,
	{_State94, TraitOpTypeType}:                          _GotoState110Action,
	{_State94, FieldDefType}:                             _GotoState210Action,
	{_State94, ProperImplicitFieldDefsType}:              _GotoState213Action,
	{_State94, ImplicitFieldDefsType}:                    _GotoState212Action,
	{_State94, ImplicitStructDefType}:                    _GotoState103Action,
	{_State94, ExplicitStructDefType}:                    _GotoState51Action,
	{_State94, EnumValueDefType}:                         _GotoState209Action,
	{_State94, ImplicitEnumValueDefsType}:                _GotoState211Action,
	{_State94, ImplicitEnumDefType}:                      _GotoState102Action,
	{_State94, ExplicitEnumDefType}:                      _GotoState100Action,
	{_State94, TraitDefType}:                             _GotoState109Action,
	{_State94, FuncTypeType}:                             _GotoState101Action,
	{_State98, LparenToken}:                              _GotoState216Action,
	{_State106, IdentifierToken}:                         _GotoState93Action,
	{_State106, StructToken}:                             _GotoState32Action,
	{_State106, EnumToken}:                               _GotoState90Action,
	{_State106, TraitToken}:                              _GotoState98Action,
	{_State106, FuncToken}:                               _GotoState92Action,
	{_State106, LparenToken}:                             _GotoState94Action,
	{_State106, LbracketToken}:                           _GotoState24Action,
	{_State106, DotToken}:                                _GotoState89Action,
	{_State106, QuestionToken}:                           _GotoState96Action,
	{_State106, ExclaimToken}:                            _GotoState91Action,
	{_State106, TildeTildeToken}:                         _GotoState97Action,
	{_State106, BitNegToken}:                             _GotoState88Action,
	{_State106, BitAndToken}:                             _GotoState87Action,
	{_State106, ParseErrorToken}:                         _GotoState95Action,
	{_State106, InitializableTypeType}:                   _GotoState104Action,
	{_State106, SliceTypeType}:                           _GotoState69Action,
	{_State106, ArrayTypeType}:                           _GotoState41Action,
	{_State106, MapTypeType}:                             _GotoState58Action,
	{_State106, AtomTypeType}:                            _GotoState99Action,
	{_State106, ParseErrorTypeType}:                      _GotoState105Action,
	{_State106, ReturnableTypeType}:                      _GotoState217Action,
	{_State106, PrefixedTypeType}:                        _GotoState107Action,
	{_State106, PrefixTypeOpType}:                        _GotoState106Action,
	{_State106, ImplicitStructDefType}:                   _GotoState103Action,
	{_State106, ExplicitStructDefType}:                   _GotoState51Action,
	{_State106, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State106, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State106, TraitDefType}:                            _GotoState109Action,
	{_State106, FuncTypeType}:                            _GotoState101Action,
	{_State111, RbracketToken}:                           _GotoState222Action,
	{_State111, CommaToken}:                              _GotoState220Action,
	{_State111, ColonToken}:                              _GotoState219Action,
	{_State111, AddToken}:                                _GotoState218Action,
	{_State111, SubToken}:                                _GotoState223Action,
	{_State111, MulToken}:                                _GotoState221Action,
	{_State111, TraitOpType}:                             _GotoState224Action,
	{_State112, IntegerLiteralToken}:                     _GotoState22Action,
	{_State112, FloatLiteralToken}:                       _GotoState18Action,
	{_State112, RuneLiteralToken}:                        _GotoState30Action,
	{_State112, StringLiteralToken}:                      _GotoState31Action,
	{_State112, IdentifierToken}:                         _GotoState21Action,
	{_State112, TrueToken}:                               _GotoState34Action,
	{_State112, FalseToken}:                              _GotoState17Action,
	{_State112, StructToken}:                             _GotoState32Action,
	{_State112, FuncToken}:                               _GotoState19Action,
	{_State112, VarToken}:                                _GotoState35Action,
	{_State112, LetToken}:                                _GotoState25Action,
	{_State112, NotToken}:                                _GotoState28Action,
	{_State112, LabelDeclToken}:                          _GotoState23Action,
	{_State112, LparenToken}:                             _GotoState26Action,
	{_State112, LbracketToken}:                           _GotoState24Action,
	{_State112, SubToken}:                                _GotoState33Action,
	{_State112, MulToken}:                                _GotoState27Action,
	{_State112, BitNegToken}:                             _GotoState16Action,
	{_State112, BitAndToken}:                             _GotoState15Action,
	{_State112, GreaterToken}:                            _GotoState20Action,
	{_State112, ParseErrorToken}:                         _GotoState29Action,
	{_State112, VarDeclPatternType}:                      _GotoState70Action,
	{_State112, VarOrLetType}:                            _GotoState71Action,
	{_State112, ExpressionType}:                          _GotoState225Action,
	{_State112, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State112, SequenceExprType}:                        _GotoState68Action,
	{_State112, CallExprType}:                            _GotoState49Action,
	{_State112, AtomExprType}:                            _GotoState42Action,
	{_State112, ParseErrorExprType}:                      _GotoState62Action,
	{_State112, LiteralExprType}:                         _GotoState57Action,
	{_State112, IdentifierExprType}:                      _GotoState52Action,
	{_State112, BlockExprType}:                           _GotoState48Action,
	{_State112, InitializeExprType}:                      _GotoState56Action,
	{_State112, ImplicitStructExprType}:                  _GotoState53Action,
	{_State112, AccessibleExprType}:                      _GotoState37Action,
	{_State112, AccessExprType}:                          _GotoState36Action,
	{_State112, IndexExprType}:                           _GotoState54Action,
	{_State112, PostfixableExprType}:                     _GotoState64Action,
	{_State112, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State112, PrefixableExprType}:                      _GotoState67Action,
	{_State112, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State112, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State112, MulExprType}:                             _GotoState59Action,
	{_State112, BinaryMulExprType}:                       _GotoState46Action,
	{_State112, AddExprType}:                             _GotoState38Action,
	{_State112, BinaryAddExprType}:                       _GotoState43Action,
	{_State112, CmpExprType}:                             _GotoState50Action,
	{_State112, BinaryCmpExprType}:                       _GotoState45Action,
	{_State112, AndExprType}:                             _GotoState39Action,
	{_State112, BinaryAndExprType}:                       _GotoState44Action,
	{_State112, OrExprType}:                              _GotoState61Action,
	{_State112, BinaryOrExprType}:                        _GotoState47Action,
	{_State112, InitializableTypeType}:                   _GotoState55Action,
	{_State112, SliceTypeType}:                           _GotoState69Action,
	{_State112, ArrayTypeType}:                           _GotoState41Action,
	{_State112, MapTypeType}:                             _GotoState58Action,
	{_State112, ExplicitStructDefType}:                   _GotoState51Action,
	{_State112, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State114, AssignToken}:                             _GotoState226Action,
	{_State116, RparenToken}:                             _GotoState227Action,
	{_State117, ColonToken}:                              _GotoState228Action,
	{_State118, ColonToken}:                              _GotoState229Action,
	{_State118, EllipsisToken}:                           _GotoState230Action,
	{_State119, CommaToken}:                              _GotoState231Action,
	{_State120, IdentifierToken}:                         _GotoState208Action,
	{_State120, UnsafeToken}:                             _GotoState170Action,
	{_State120, StructToken}:                             _GotoState32Action,
	{_State120, EnumToken}:                               _GotoState90Action,
	{_State120, TraitToken}:                              _GotoState98Action,
	{_State120, FuncToken}:                               _GotoState92Action,
	{_State120, LparenToken}:                             _GotoState94Action,
	{_State120, LbracketToken}:                           _GotoState24Action,
	{_State120, DotToken}:                                _GotoState89Action,
	{_State120, QuestionToken}:                           _GotoState96Action,
	{_State120, ExclaimToken}:                            _GotoState91Action,
	{_State120, TildeTildeToken}:                         _GotoState97Action,
	{_State120, BitNegToken}:                             _GotoState88Action,
	{_State120, BitAndToken}:                             _GotoState87Action,
	{_State120, ParseErrorToken}:                         _GotoState95Action,
	{_State120, UnsafeStatementType}:                     _GotoState214Action,
	{_State120, InitializableTypeType}:                   _GotoState104Action,
	{_State120, SliceTypeType}:                           _GotoState69Action,
	{_State120, ArrayTypeType}:                           _GotoState41Action,
	{_State120, MapTypeType}:                             _GotoState58Action,
	{_State120, AtomTypeType}:                            _GotoState99Action,
	{_State120, ParseErrorTypeType}:                      _GotoState105Action,
	{_State120, ReturnableTypeType}:                      _GotoState108Action,
	{_State120, PrefixedTypeType}:                        _GotoState107Action,
	{_State120, PrefixTypeOpType}:                        _GotoState106Action,
	{_State120, ValueTypeType}:                           _GotoState215Action,
	{_State120, TraitOpTypeType}:                         _GotoState110Action,
	{_State120, FieldDefType}:                            _GotoState233Action,
	{_State120, ImplicitStructDefType}:                   _GotoState103Action,
	{_State120, ProperExplicitFieldDefsType}:             _GotoState234Action,
	{_State120, ExplicitFieldDefsType}:                   _GotoState232Action,
	{_State120, ExplicitStructDefType}:                   _GotoState51Action,
	{_State120, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State120, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State120, TraitDefType}:                            _GotoState109Action,
	{_State120, FuncTypeType}:                            _GotoState101Action,
	{_State121, IdentifierToken}:                         _GotoState93Action,
	{_State121, StructToken}:                             _GotoState32Action,
	{_State121, EnumToken}:                               _GotoState90Action,
	{_State121, TraitToken}:                              _GotoState98Action,
	{_State121, FuncToken}:                               _GotoState92Action,
	{_State121, LparenToken}:                             _GotoState94Action,
	{_State121, LbracketToken}:                           _GotoState24Action,
	{_State121, DotToken}:                                _GotoState89Action,
	{_State121, QuestionToken}:                           _GotoState96Action,
	{_State121, ExclaimToken}:                            _GotoState91Action,
	{_State121, TildeTildeToken}:                         _GotoState97Action,
	{_State121, BitNegToken}:                             _GotoState88Action,
	{_State121, BitAndToken}:                             _GotoState87Action,
	{_State121, ParseErrorToken}:                         _GotoState95Action,
	{_State121, ProperGenericArgumentsType}:              _GotoState236Action,
	{_State121, GenericArgumentsType}:                    _GotoState235Action,
	{_State121, InitializableTypeType}:                   _GotoState104Action,
	{_State121, SliceTypeType}:                           _GotoState69Action,
	{_State121, ArrayTypeType}:                           _GotoState41Action,
	{_State121, MapTypeType}:                             _GotoState58Action,
	{_State121, AtomTypeType}:                            _GotoState99Action,
	{_State121, ParseErrorTypeType}:                      _GotoState105Action,
	{_State121, ReturnableTypeType}:                      _GotoState108Action,
	{_State121, PrefixedTypeType}:                        _GotoState107Action,
	{_State121, PrefixTypeOpType}:                        _GotoState106Action,
	{_State121, ValueTypeType}:                           _GotoState237Action,
	{_State121, TraitOpTypeType}:                         _GotoState110Action,
	{_State121, ImplicitStructDefType}:                   _GotoState103Action,
	{_State121, ExplicitStructDefType}:                   _GotoState51Action,
	{_State121, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State121, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State121, TraitDefType}:                            _GotoState109Action,
	{_State121, FuncTypeType}:                            _GotoState101Action,
	{_State122, IdentifierToken}:                         _GotoState238Action,
	{_State123, IntegerLiteralToken}:                     _GotoState22Action,
	{_State123, FloatLiteralToken}:                       _GotoState18Action,
	{_State123, RuneLiteralToken}:                        _GotoState30Action,
	{_State123, StringLiteralToken}:                      _GotoState31Action,
	{_State123, IdentifierToken}:                         _GotoState114Action,
	{_State123, TrueToken}:                               _GotoState34Action,
	{_State123, FalseToken}:                              _GotoState17Action,
	{_State123, StructToken}:                             _GotoState32Action,
	{_State123, FuncToken}:                               _GotoState19Action,
	{_State123, VarToken}:                                _GotoState35Action,
	{_State123, LetToken}:                                _GotoState25Action,
	{_State123, NotToken}:                                _GotoState28Action,
	{_State123, LabelDeclToken}:                          _GotoState23Action,
	{_State123, LparenToken}:                             _GotoState26Action,
	{_State123, LbracketToken}:                           _GotoState24Action,
	{_State123, ColonToken}:                              _GotoState112Action,
	{_State123, EllipsisToken}:                           _GotoState113Action,
	{_State123, SubToken}:                                _GotoState33Action,
	{_State123, MulToken}:                                _GotoState27Action,
	{_State123, BitNegToken}:                             _GotoState16Action,
	{_State123, BitAndToken}:                             _GotoState15Action,
	{_State123, GreaterToken}:                            _GotoState20Action,
	{_State123, ParseErrorToken}:                         _GotoState29Action,
	{_State123, VarDeclPatternType}:                      _GotoState70Action,
	{_State123, VarOrLetType}:                            _GotoState71Action,
	{_State123, ExpressionType}:                          _GotoState118Action,
	{_State123, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State123, SequenceExprType}:                        _GotoState68Action,
	{_State123, CallExprType}:                            _GotoState49Action,
	{_State123, ArgumentType}:                            _GotoState239Action,
	{_State123, ColonExprType}:                           _GotoState117Action,
	{_State123, AtomExprType}:                            _GotoState42Action,
	{_State123, ParseErrorExprType}:                      _GotoState62Action,
	{_State123, LiteralExprType}:                         _GotoState57Action,
	{_State123, IdentifierExprType}:                      _GotoState52Action,
	{_State123, BlockExprType}:                           _GotoState48Action,
	{_State123, InitializeExprType}:                      _GotoState56Action,
	{_State123, ImplicitStructExprType}:                  _GotoState53Action,
	{_State123, AccessibleExprType}:                      _GotoState37Action,
	{_State123, AccessExprType}:                          _GotoState36Action,
	{_State123, IndexExprType}:                           _GotoState54Action,
	{_State123, PostfixableExprType}:                     _GotoState64Action,
	{_State123, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State123, PrefixableExprType}:                      _GotoState67Action,
	{_State123, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State123, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State123, MulExprType}:                             _GotoState59Action,
	{_State123, BinaryMulExprType}:                       _GotoState46Action,
	{_State123, AddExprType}:                             _GotoState38Action,
	{_State123, BinaryAddExprType}:                       _GotoState43Action,
	{_State123, CmpExprType}:                             _GotoState50Action,
	{_State123, BinaryCmpExprType}:                       _GotoState45Action,
	{_State123, AndExprType}:                             _GotoState39Action,
	{_State123, BinaryAndExprType}:                       _GotoState44Action,
	{_State123, OrExprType}:                              _GotoState61Action,
	{_State123, BinaryOrExprType}:                        _GotoState47Action,
	{_State123, InitializableTypeType}:                   _GotoState55Action,
	{_State123, SliceTypeType}:                           _GotoState69Action,
	{_State123, ArrayTypeType}:                           _GotoState41Action,
	{_State123, MapTypeType}:                             _GotoState58Action,
	{_State123, ExplicitStructDefType}:                   _GotoState51Action,
	{_State123, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State125, LparenToken}:                             _GotoState240Action,
	{_State130, IntegerLiteralToken}:                     _GotoState22Action,
	{_State130, FloatLiteralToken}:                       _GotoState18Action,
	{_State130, RuneLiteralToken}:                        _GotoState30Action,
	{_State130, StringLiteralToken}:                      _GotoState31Action,
	{_State130, IdentifierToken}:                         _GotoState21Action,
	{_State130, TrueToken}:                               _GotoState34Action,
	{_State130, FalseToken}:                              _GotoState17Action,
	{_State130, StructToken}:                             _GotoState32Action,
	{_State130, FuncToken}:                               _GotoState19Action,
	{_State130, NotToken}:                                _GotoState28Action,
	{_State130, LabelDeclToken}:                          _GotoState23Action,
	{_State130, LparenToken}:                             _GotoState26Action,
	{_State130, LbracketToken}:                           _GotoState24Action,
	{_State130, SubToken}:                                _GotoState33Action,
	{_State130, MulToken}:                                _GotoState27Action,
	{_State130, BitNegToken}:                             _GotoState16Action,
	{_State130, BitAndToken}:                             _GotoState15Action,
	{_State130, ParseErrorToken}:                         _GotoState29Action,
	{_State130, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State130, CallExprType}:                            _GotoState49Action,
	{_State130, AtomExprType}:                            _GotoState42Action,
	{_State130, ParseErrorExprType}:                      _GotoState62Action,
	{_State130, LiteralExprType}:                         _GotoState57Action,
	{_State130, IdentifierExprType}:                      _GotoState52Action,
	{_State130, BlockExprType}:                           _GotoState48Action,
	{_State130, InitializeExprType}:                      _GotoState56Action,
	{_State130, ImplicitStructExprType}:                  _GotoState53Action,
	{_State130, AccessibleExprType}:                      _GotoState37Action,
	{_State130, AccessExprType}:                          _GotoState36Action,
	{_State130, IndexExprType}:                           _GotoState54Action,
	{_State130, PostfixableExprType}:                     _GotoState64Action,
	{_State130, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State130, PrefixableExprType}:                      _GotoState67Action,
	{_State130, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State130, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State130, MulExprType}:                             _GotoState241Action,
	{_State130, BinaryMulExprType}:                       _GotoState46Action,
	{_State130, InitializableTypeType}:                   _GotoState55Action,
	{_State130, SliceTypeType}:                           _GotoState69Action,
	{_State130, ArrayTypeType}:                           _GotoState41Action,
	{_State130, MapTypeType}:                             _GotoState58Action,
	{_State130, ExplicitStructDefType}:                   _GotoState51Action,
	{_State130, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State131, IntegerLiteralToken}:                     _GotoState22Action,
	{_State131, FloatLiteralToken}:                       _GotoState18Action,
	{_State131, RuneLiteralToken}:                        _GotoState30Action,
	{_State131, StringLiteralToken}:                      _GotoState31Action,
	{_State131, IdentifierToken}:                         _GotoState21Action,
	{_State131, TrueToken}:                               _GotoState34Action,
	{_State131, FalseToken}:                              _GotoState17Action,
	{_State131, StructToken}:                             _GotoState32Action,
	{_State131, FuncToken}:                               _GotoState19Action,
	{_State131, NotToken}:                                _GotoState28Action,
	{_State131, LabelDeclToken}:                          _GotoState23Action,
	{_State131, LparenToken}:                             _GotoState26Action,
	{_State131, LbracketToken}:                           _GotoState24Action,
	{_State131, SubToken}:                                _GotoState33Action,
	{_State131, MulToken}:                                _GotoState27Action,
	{_State131, BitNegToken}:                             _GotoState16Action,
	{_State131, BitAndToken}:                             _GotoState15Action,
	{_State131, ParseErrorToken}:                         _GotoState29Action,
	{_State131, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State131, CallExprType}:                            _GotoState49Action,
	{_State131, AtomExprType}:                            _GotoState42Action,
	{_State131, ParseErrorExprType}:                      _GotoState62Action,
	{_State131, LiteralExprType}:                         _GotoState57Action,
	{_State131, IdentifierExprType}:                      _GotoState52Action,
	{_State131, BlockExprType}:                           _GotoState48Action,
	{_State131, InitializeExprType}:                      _GotoState56Action,
	{_State131, ImplicitStructExprType}:                  _GotoState53Action,
	{_State131, AccessibleExprType}:                      _GotoState37Action,
	{_State131, AccessExprType}:                          _GotoState36Action,
	{_State131, IndexExprType}:                           _GotoState54Action,
	{_State131, PostfixableExprType}:                     _GotoState64Action,
	{_State131, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State131, PrefixableExprType}:                      _GotoState67Action,
	{_State131, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State131, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State131, MulExprType}:                             _GotoState59Action,
	{_State131, BinaryMulExprType}:                       _GotoState46Action,
	{_State131, AddExprType}:                             _GotoState38Action,
	{_State131, BinaryAddExprType}:                       _GotoState43Action,
	{_State131, CmpExprType}:                             _GotoState242Action,
	{_State131, BinaryCmpExprType}:                       _GotoState45Action,
	{_State131, InitializableTypeType}:                   _GotoState55Action,
	{_State131, SliceTypeType}:                           _GotoState69Action,
	{_State131, ArrayTypeType}:                           _GotoState41Action,
	{_State131, MapTypeType}:                             _GotoState58Action,
	{_State131, ExplicitStructDefType}:                   _GotoState51Action,
	{_State131, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State138, IntegerLiteralToken}:                     _GotoState22Action,
	{_State138, FloatLiteralToken}:                       _GotoState18Action,
	{_State138, RuneLiteralToken}:                        _GotoState30Action,
	{_State138, StringLiteralToken}:                      _GotoState31Action,
	{_State138, IdentifierToken}:                         _GotoState21Action,
	{_State138, TrueToken}:                               _GotoState34Action,
	{_State138, FalseToken}:                              _GotoState17Action,
	{_State138, StructToken}:                             _GotoState32Action,
	{_State138, FuncToken}:                               _GotoState19Action,
	{_State138, NotToken}:                                _GotoState28Action,
	{_State138, LabelDeclToken}:                          _GotoState23Action,
	{_State138, LparenToken}:                             _GotoState26Action,
	{_State138, LbracketToken}:                           _GotoState24Action,
	{_State138, SubToken}:                                _GotoState33Action,
	{_State138, MulToken}:                                _GotoState27Action,
	{_State138, BitNegToken}:                             _GotoState16Action,
	{_State138, BitAndToken}:                             _GotoState15Action,
	{_State138, ParseErrorToken}:                         _GotoState29Action,
	{_State138, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State138, CallExprType}:                            _GotoState49Action,
	{_State138, AtomExprType}:                            _GotoState42Action,
	{_State138, ParseErrorExprType}:                      _GotoState62Action,
	{_State138, LiteralExprType}:                         _GotoState57Action,
	{_State138, IdentifierExprType}:                      _GotoState52Action,
	{_State138, BlockExprType}:                           _GotoState48Action,
	{_State138, InitializeExprType}:                      _GotoState56Action,
	{_State138, ImplicitStructExprType}:                  _GotoState53Action,
	{_State138, AccessibleExprType}:                      _GotoState37Action,
	{_State138, AccessExprType}:                          _GotoState36Action,
	{_State138, IndexExprType}:                           _GotoState54Action,
	{_State138, PostfixableExprType}:                     _GotoState64Action,
	{_State138, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State138, PrefixableExprType}:                      _GotoState67Action,
	{_State138, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State138, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State138, MulExprType}:                             _GotoState59Action,
	{_State138, BinaryMulExprType}:                       _GotoState46Action,
	{_State138, AddExprType}:                             _GotoState243Action,
	{_State138, BinaryAddExprType}:                       _GotoState43Action,
	{_State138, InitializableTypeType}:                   _GotoState55Action,
	{_State138, SliceTypeType}:                           _GotoState69Action,
	{_State138, ArrayTypeType}:                           _GotoState41Action,
	{_State138, MapTypeType}:                             _GotoState58Action,
	{_State138, ExplicitStructDefType}:                   _GotoState51Action,
	{_State138, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State139, IntegerLiteralToken}:                     _GotoState22Action,
	{_State139, FloatLiteralToken}:                       _GotoState18Action,
	{_State139, RuneLiteralToken}:                        _GotoState30Action,
	{_State139, StringLiteralToken}:                      _GotoState31Action,
	{_State139, IdentifierToken}:                         _GotoState114Action,
	{_State139, TrueToken}:                               _GotoState34Action,
	{_State139, FalseToken}:                              _GotoState17Action,
	{_State139, StructToken}:                             _GotoState32Action,
	{_State139, FuncToken}:                               _GotoState19Action,
	{_State139, VarToken}:                                _GotoState35Action,
	{_State139, LetToken}:                                _GotoState25Action,
	{_State139, NotToken}:                                _GotoState28Action,
	{_State139, LabelDeclToken}:                          _GotoState23Action,
	{_State139, LparenToken}:                             _GotoState26Action,
	{_State139, LbracketToken}:                           _GotoState24Action,
	{_State139, ColonToken}:                              _GotoState112Action,
	{_State139, EllipsisToken}:                           _GotoState113Action,
	{_State139, SubToken}:                                _GotoState33Action,
	{_State139, MulToken}:                                _GotoState27Action,
	{_State139, BitNegToken}:                             _GotoState16Action,
	{_State139, BitAndToken}:                             _GotoState15Action,
	{_State139, GreaterToken}:                            _GotoState20Action,
	{_State139, ParseErrorToken}:                         _GotoState29Action,
	{_State139, VarDeclPatternType}:                      _GotoState70Action,
	{_State139, VarOrLetType}:                            _GotoState71Action,
	{_State139, ExpressionType}:                          _GotoState118Action,
	{_State139, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State139, SequenceExprType}:                        _GotoState68Action,
	{_State139, CallExprType}:                            _GotoState49Action,
	{_State139, ProperArgumentsType}:                     _GotoState119Action,
	{_State139, ArgumentsType}:                           _GotoState244Action,
	{_State139, ArgumentType}:                            _GotoState115Action,
	{_State139, ColonExprType}:                           _GotoState117Action,
	{_State139, AtomExprType}:                            _GotoState42Action,
	{_State139, ParseErrorExprType}:                      _GotoState62Action,
	{_State139, LiteralExprType}:                         _GotoState57Action,
	{_State139, IdentifierExprType}:                      _GotoState52Action,
	{_State139, BlockExprType}:                           _GotoState48Action,
	{_State139, InitializeExprType}:                      _GotoState56Action,
	{_State139, ImplicitStructExprType}:                  _GotoState53Action,
	{_State139, AccessibleExprType}:                      _GotoState37Action,
	{_State139, AccessExprType}:                          _GotoState36Action,
	{_State139, IndexExprType}:                           _GotoState54Action,
	{_State139, PostfixableExprType}:                     _GotoState64Action,
	{_State139, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State139, PrefixableExprType}:                      _GotoState67Action,
	{_State139, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State139, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State139, MulExprType}:                             _GotoState59Action,
	{_State139, BinaryMulExprType}:                       _GotoState46Action,
	{_State139, AddExprType}:                             _GotoState38Action,
	{_State139, BinaryAddExprType}:                       _GotoState43Action,
	{_State139, CmpExprType}:                             _GotoState50Action,
	{_State139, BinaryCmpExprType}:                       _GotoState45Action,
	{_State139, AndExprType}:                             _GotoState39Action,
	{_State139, BinaryAndExprType}:                       _GotoState44Action,
	{_State139, OrExprType}:                              _GotoState61Action,
	{_State139, BinaryOrExprType}:                        _GotoState47Action,
	{_State139, InitializableTypeType}:                   _GotoState55Action,
	{_State139, SliceTypeType}:                           _GotoState69Action,
	{_State139, ArrayTypeType}:                           _GotoState41Action,
	{_State139, MapTypeType}:                             _GotoState58Action,
	{_State139, ExplicitStructDefType}:                   _GotoState51Action,
	{_State139, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State146, IntegerLiteralToken}:                     _GotoState22Action,
	{_State146, FloatLiteralToken}:                       _GotoState18Action,
	{_State146, RuneLiteralToken}:                        _GotoState30Action,
	{_State146, StringLiteralToken}:                      _GotoState31Action,
	{_State146, IdentifierToken}:                         _GotoState21Action,
	{_State146, TrueToken}:                               _GotoState34Action,
	{_State146, FalseToken}:                              _GotoState17Action,
	{_State146, StructToken}:                             _GotoState32Action,
	{_State146, FuncToken}:                               _GotoState19Action,
	{_State146, NotToken}:                                _GotoState28Action,
	{_State146, LabelDeclToken}:                          _GotoState23Action,
	{_State146, LparenToken}:                             _GotoState26Action,
	{_State146, LbracketToken}:                           _GotoState24Action,
	{_State146, SubToken}:                                _GotoState33Action,
	{_State146, MulToken}:                                _GotoState27Action,
	{_State146, BitNegToken}:                             _GotoState16Action,
	{_State146, BitAndToken}:                             _GotoState15Action,
	{_State146, ParseErrorToken}:                         _GotoState29Action,
	{_State146, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State146, CallExprType}:                            _GotoState49Action,
	{_State146, AtomExprType}:                            _GotoState42Action,
	{_State146, ParseErrorExprType}:                      _GotoState62Action,
	{_State146, LiteralExprType}:                         _GotoState57Action,
	{_State146, IdentifierExprType}:                      _GotoState52Action,
	{_State146, BlockExprType}:                           _GotoState48Action,
	{_State146, InitializeExprType}:                      _GotoState56Action,
	{_State146, ImplicitStructExprType}:                  _GotoState53Action,
	{_State146, AccessibleExprType}:                      _GotoState37Action,
	{_State146, AccessExprType}:                          _GotoState36Action,
	{_State146, IndexExprType}:                           _GotoState54Action,
	{_State146, PostfixableExprType}:                     _GotoState64Action,
	{_State146, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State146, PrefixableExprType}:                      _GotoState245Action,
	{_State146, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State146, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State146, InitializableTypeType}:                   _GotoState55Action,
	{_State146, SliceTypeType}:                           _GotoState69Action,
	{_State146, ArrayTypeType}:                           _GotoState41Action,
	{_State146, MapTypeType}:                             _GotoState58Action,
	{_State146, ExplicitStructDefType}:                   _GotoState51Action,
	{_State146, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State147, LbraceToken}:                             _GotoState73Action,
	{_State147, StatementBlockType}:                      _GotoState246Action,
	{_State148, IntegerLiteralToken}:                     _GotoState22Action,
	{_State148, FloatLiteralToken}:                       _GotoState18Action,
	{_State148, RuneLiteralToken}:                        _GotoState30Action,
	{_State148, StringLiteralToken}:                      _GotoState31Action,
	{_State148, IdentifierToken}:                         _GotoState21Action,
	{_State148, TrueToken}:                               _GotoState34Action,
	{_State148, FalseToken}:                              _GotoState17Action,
	{_State148, StructToken}:                             _GotoState32Action,
	{_State148, FuncToken}:                               _GotoState19Action,
	{_State148, VarToken}:                                _GotoState35Action,
	{_State148, LetToken}:                                _GotoState25Action,
	{_State148, NotToken}:                                _GotoState28Action,
	{_State148, LabelDeclToken}:                          _GotoState23Action,
	{_State148, LparenToken}:                             _GotoState26Action,
	{_State148, LbracketToken}:                           _GotoState24Action,
	{_State148, SubToken}:                                _GotoState33Action,
	{_State148, MulToken}:                                _GotoState27Action,
	{_State148, BitNegToken}:                             _GotoState16Action,
	{_State148, BitAndToken}:                             _GotoState15Action,
	{_State148, GreaterToken}:                            _GotoState20Action,
	{_State148, ParseErrorToken}:                         _GotoState29Action,
	{_State148, VarDeclPatternType}:                      _GotoState70Action,
	{_State148, VarOrLetType}:                            _GotoState71Action,
	{_State148, AssignPatternType}:                       _GotoState247Action,
	{_State148, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State148, SequenceExprType}:                        _GotoState249Action,
	{_State148, ForAssignmentType}:                       _GotoState248Action,
	{_State148, CallExprType}:                            _GotoState49Action,
	{_State148, AtomExprType}:                            _GotoState42Action,
	{_State148, ParseErrorExprType}:                      _GotoState62Action,
	{_State148, LiteralExprType}:                         _GotoState57Action,
	{_State148, IdentifierExprType}:                      _GotoState52Action,
	{_State148, BlockExprType}:                           _GotoState48Action,
	{_State148, InitializeExprType}:                      _GotoState56Action,
	{_State148, ImplicitStructExprType}:                  _GotoState53Action,
	{_State148, AccessibleExprType}:                      _GotoState37Action,
	{_State148, AccessExprType}:                          _GotoState36Action,
	{_State148, IndexExprType}:                           _GotoState54Action,
	{_State148, PostfixableExprType}:                     _GotoState64Action,
	{_State148, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State148, PrefixableExprType}:                      _GotoState67Action,
	{_State148, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State148, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State148, MulExprType}:                             _GotoState59Action,
	{_State148, BinaryMulExprType}:                       _GotoState46Action,
	{_State148, AddExprType}:                             _GotoState38Action,
	{_State148, BinaryAddExprType}:                       _GotoState43Action,
	{_State148, CmpExprType}:                             _GotoState50Action,
	{_State148, BinaryCmpExprType}:                       _GotoState45Action,
	{_State148, AndExprType}:                             _GotoState39Action,
	{_State148, BinaryAndExprType}:                       _GotoState44Action,
	{_State148, OrExprType}:                              _GotoState61Action,
	{_State148, BinaryOrExprType}:                        _GotoState47Action,
	{_State148, InitializableTypeType}:                   _GotoState55Action,
	{_State148, SliceTypeType}:                           _GotoState69Action,
	{_State148, ArrayTypeType}:                           _GotoState41Action,
	{_State148, MapTypeType}:                             _GotoState58Action,
	{_State148, ExplicitStructDefType}:                   _GotoState51Action,
	{_State148, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State149, IntegerLiteralToken}:                     _GotoState22Action,
	{_State149, FloatLiteralToken}:                       _GotoState18Action,
	{_State149, RuneLiteralToken}:                        _GotoState30Action,
	{_State149, StringLiteralToken}:                      _GotoState31Action,
	{_State149, IdentifierToken}:                         _GotoState21Action,
	{_State149, TrueToken}:                               _GotoState34Action,
	{_State149, FalseToken}:                              _GotoState17Action,
	{_State149, CaseToken}:                               _GotoState250Action,
	{_State149, StructToken}:                             _GotoState32Action,
	{_State149, FuncToken}:                               _GotoState19Action,
	{_State149, VarToken}:                                _GotoState35Action,
	{_State149, LetToken}:                                _GotoState25Action,
	{_State149, NotToken}:                                _GotoState28Action,
	{_State149, LabelDeclToken}:                          _GotoState23Action,
	{_State149, LparenToken}:                             _GotoState26Action,
	{_State149, LbracketToken}:                           _GotoState24Action,
	{_State149, SubToken}:                                _GotoState33Action,
	{_State149, MulToken}:                                _GotoState27Action,
	{_State149, BitNegToken}:                             _GotoState16Action,
	{_State149, BitAndToken}:                             _GotoState15Action,
	{_State149, GreaterToken}:                            _GotoState20Action,
	{_State149, ParseErrorToken}:                         _GotoState29Action,
	{_State149, VarDeclPatternType}:                      _GotoState70Action,
	{_State149, VarOrLetType}:                            _GotoState71Action,
	{_State149, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State149, SequenceExprType}:                        _GotoState252Action,
	{_State149, ConditionType}:                           _GotoState251Action,
	{_State149, CallExprType}:                            _GotoState49Action,
	{_State149, AtomExprType}:                            _GotoState42Action,
	{_State149, ParseErrorExprType}:                      _GotoState62Action,
	{_State149, LiteralExprType}:                         _GotoState57Action,
	{_State149, IdentifierExprType}:                      _GotoState52Action,
	{_State149, BlockExprType}:                           _GotoState48Action,
	{_State149, InitializeExprType}:                      _GotoState56Action,
	{_State149, ImplicitStructExprType}:                  _GotoState53Action,
	{_State149, AccessibleExprType}:                      _GotoState37Action,
	{_State149, AccessExprType}:                          _GotoState36Action,
	{_State149, IndexExprType}:                           _GotoState54Action,
	{_State149, PostfixableExprType}:                     _GotoState64Action,
	{_State149, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State149, PrefixableExprType}:                      _GotoState67Action,
	{_State149, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State149, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State149, MulExprType}:                             _GotoState59Action,
	{_State149, BinaryMulExprType}:                       _GotoState46Action,
	{_State149, AddExprType}:                             _GotoState38Action,
	{_State149, BinaryAddExprType}:                       _GotoState43Action,
	{_State149, CmpExprType}:                             _GotoState50Action,
	{_State149, BinaryCmpExprType}:                       _GotoState45Action,
	{_State149, AndExprType}:                             _GotoState39Action,
	{_State149, BinaryAndExprType}:                       _GotoState44Action,
	{_State149, OrExprType}:                              _GotoState61Action,
	{_State149, BinaryOrExprType}:                        _GotoState47Action,
	{_State149, InitializableTypeType}:                   _GotoState55Action,
	{_State149, SliceTypeType}:                           _GotoState69Action,
	{_State149, ArrayTypeType}:                           _GotoState41Action,
	{_State149, MapTypeType}:                             _GotoState58Action,
	{_State149, ExplicitStructDefType}:                   _GotoState51Action,
	{_State149, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State150, IntegerLiteralToken}:                     _GotoState22Action,
	{_State150, FloatLiteralToken}:                       _GotoState18Action,
	{_State150, RuneLiteralToken}:                        _GotoState30Action,
	{_State150, StringLiteralToken}:                      _GotoState31Action,
	{_State150, IdentifierToken}:                         _GotoState21Action,
	{_State150, TrueToken}:                               _GotoState34Action,
	{_State150, FalseToken}:                              _GotoState17Action,
	{_State150, StructToken}:                             _GotoState32Action,
	{_State150, FuncToken}:                               _GotoState19Action,
	{_State150, VarToken}:                                _GotoState35Action,
	{_State150, LetToken}:                                _GotoState25Action,
	{_State150, NotToken}:                                _GotoState28Action,
	{_State150, LabelDeclToken}:                          _GotoState23Action,
	{_State150, LparenToken}:                             _GotoState26Action,
	{_State150, LbracketToken}:                           _GotoState24Action,
	{_State150, SubToken}:                                _GotoState33Action,
	{_State150, MulToken}:                                _GotoState27Action,
	{_State150, BitNegToken}:                             _GotoState16Action,
	{_State150, BitAndToken}:                             _GotoState15Action,
	{_State150, GreaterToken}:                            _GotoState20Action,
	{_State150, ParseErrorToken}:                         _GotoState29Action,
	{_State150, VarDeclPatternType}:                      _GotoState70Action,
	{_State150, VarOrLetType}:                            _GotoState71Action,
	{_State150, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State150, SequenceExprType}:                        _GotoState253Action,
	{_State150, CallExprType}:                            _GotoState49Action,
	{_State150, AtomExprType}:                            _GotoState42Action,
	{_State150, ParseErrorExprType}:                      _GotoState62Action,
	{_State150, LiteralExprType}:                         _GotoState57Action,
	{_State150, IdentifierExprType}:                      _GotoState52Action,
	{_State150, BlockExprType}:                           _GotoState48Action,
	{_State150, InitializeExprType}:                      _GotoState56Action,
	{_State150, ImplicitStructExprType}:                  _GotoState53Action,
	{_State150, AccessibleExprType}:                      _GotoState37Action,
	{_State150, AccessExprType}:                          _GotoState36Action,
	{_State150, IndexExprType}:                           _GotoState54Action,
	{_State150, PostfixableExprType}:                     _GotoState64Action,
	{_State150, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State150, PrefixableExprType}:                      _GotoState67Action,
	{_State150, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State150, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State150, MulExprType}:                             _GotoState59Action,
	{_State150, BinaryMulExprType}:                       _GotoState46Action,
	{_State150, AddExprType}:                             _GotoState38Action,
	{_State150, BinaryAddExprType}:                       _GotoState43Action,
	{_State150, CmpExprType}:                             _GotoState50Action,
	{_State150, BinaryCmpExprType}:                       _GotoState45Action,
	{_State150, AndExprType}:                             _GotoState39Action,
	{_State150, BinaryAndExprType}:                       _GotoState44Action,
	{_State150, OrExprType}:                              _GotoState61Action,
	{_State150, BinaryOrExprType}:                        _GotoState47Action,
	{_State150, InitializableTypeType}:                   _GotoState55Action,
	{_State150, SliceTypeType}:                           _GotoState69Action,
	{_State150, ArrayTypeType}:                           _GotoState41Action,
	{_State150, MapTypeType}:                             _GotoState58Action,
	{_State150, ExplicitStructDefType}:                   _GotoState51Action,
	{_State150, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State155, IntegerLiteralToken}:                     _GotoState22Action,
	{_State155, FloatLiteralToken}:                       _GotoState18Action,
	{_State155, RuneLiteralToken}:                        _GotoState30Action,
	{_State155, StringLiteralToken}:                      _GotoState31Action,
	{_State155, IdentifierToken}:                         _GotoState21Action,
	{_State155, TrueToken}:                               _GotoState34Action,
	{_State155, FalseToken}:                              _GotoState17Action,
	{_State155, StructToken}:                             _GotoState32Action,
	{_State155, FuncToken}:                               _GotoState19Action,
	{_State155, NotToken}:                                _GotoState28Action,
	{_State155, LabelDeclToken}:                          _GotoState23Action,
	{_State155, LparenToken}:                             _GotoState26Action,
	{_State155, LbracketToken}:                           _GotoState24Action,
	{_State155, SubToken}:                                _GotoState33Action,
	{_State155, MulToken}:                                _GotoState27Action,
	{_State155, BitNegToken}:                             _GotoState16Action,
	{_State155, BitAndToken}:                             _GotoState15Action,
	{_State155, ParseErrorToken}:                         _GotoState29Action,
	{_State155, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State155, CallExprType}:                            _GotoState49Action,
	{_State155, AtomExprType}:                            _GotoState42Action,
	{_State155, ParseErrorExprType}:                      _GotoState62Action,
	{_State155, LiteralExprType}:                         _GotoState57Action,
	{_State155, IdentifierExprType}:                      _GotoState52Action,
	{_State155, BlockExprType}:                           _GotoState48Action,
	{_State155, InitializeExprType}:                      _GotoState56Action,
	{_State155, ImplicitStructExprType}:                  _GotoState53Action,
	{_State155, AccessibleExprType}:                      _GotoState37Action,
	{_State155, AccessExprType}:                          _GotoState36Action,
	{_State155, IndexExprType}:                           _GotoState54Action,
	{_State155, PostfixableExprType}:                     _GotoState64Action,
	{_State155, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State155, PrefixableExprType}:                      _GotoState67Action,
	{_State155, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State155, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State155, MulExprType}:                             _GotoState59Action,
	{_State155, BinaryMulExprType}:                       _GotoState46Action,
	{_State155, AddExprType}:                             _GotoState38Action,
	{_State155, BinaryAddExprType}:                       _GotoState43Action,
	{_State155, CmpExprType}:                             _GotoState50Action,
	{_State155, BinaryCmpExprType}:                       _GotoState45Action,
	{_State155, AndExprType}:                             _GotoState254Action,
	{_State155, BinaryAndExprType}:                       _GotoState44Action,
	{_State155, InitializableTypeType}:                   _GotoState55Action,
	{_State155, SliceTypeType}:                           _GotoState69Action,
	{_State155, ArrayTypeType}:                           _GotoState41Action,
	{_State155, MapTypeType}:                             _GotoState58Action,
	{_State155, ExplicitStructDefType}:                   _GotoState51Action,
	{_State155, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State158, IdentifierToken}:                         _GotoState256Action,
	{_State158, LparenToken}:                             _GotoState158Action,
	{_State158, EllipsisToken}:                           _GotoState255Action,
	{_State158, VarPatternType}:                          _GotoState259Action,
	{_State158, TuplePatternType}:                        _GotoState159Action,
	{_State158, FieldVarPatternsType}:                    _GotoState258Action,
	{_State158, FieldVarPatternType}:                     _GotoState257Action,
	{_State160, IdentifierToken}:                         _GotoState93Action,
	{_State160, StructToken}:                             _GotoState32Action,
	{_State160, EnumToken}:                               _GotoState90Action,
	{_State160, TraitToken}:                              _GotoState98Action,
	{_State160, FuncToken}:                               _GotoState92Action,
	{_State160, LparenToken}:                             _GotoState94Action,
	{_State160, LbracketToken}:                           _GotoState24Action,
	{_State160, DotToken}:                                _GotoState89Action,
	{_State160, QuestionToken}:                           _GotoState96Action,
	{_State160, ExclaimToken}:                            _GotoState91Action,
	{_State160, TildeTildeToken}:                         _GotoState97Action,
	{_State160, BitNegToken}:                             _GotoState88Action,
	{_State160, BitAndToken}:                             _GotoState87Action,
	{_State160, ParseErrorToken}:                         _GotoState95Action,
	{_State160, OptionalValueTypeType}:                   _GotoState260Action,
	{_State160, InitializableTypeType}:                   _GotoState104Action,
	{_State160, SliceTypeType}:                           _GotoState69Action,
	{_State160, ArrayTypeType}:                           _GotoState41Action,
	{_State160, MapTypeType}:                             _GotoState58Action,
	{_State160, AtomTypeType}:                            _GotoState99Action,
	{_State160, ParseErrorTypeType}:                      _GotoState105Action,
	{_State160, ReturnableTypeType}:                      _GotoState108Action,
	{_State160, PrefixedTypeType}:                        _GotoState107Action,
	{_State160, PrefixTypeOpType}:                        _GotoState106Action,
	{_State160, ValueTypeType}:                           _GotoState261Action,
	{_State160, TraitOpTypeType}:                         _GotoState110Action,
	{_State160, ImplicitStructDefType}:                   _GotoState103Action,
	{_State160, ExplicitStructDefType}:                   _GotoState51Action,
	{_State160, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State160, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State160, TraitDefType}:                            _GotoState109Action,
	{_State160, FuncTypeType}:                            _GotoState101Action,
	{_State163, IntegerLiteralToken}:                     _GotoState22Action,
	{_State163, FloatLiteralToken}:                       _GotoState18Action,
	{_State163, RuneLiteralToken}:                        _GotoState30Action,
	{_State163, StringLiteralToken}:                      _GotoState31Action,
	{_State163, IdentifierToken}:                         _GotoState21Action,
	{_State163, TrueToken}:                               _GotoState34Action,
	{_State163, FalseToken}:                              _GotoState17Action,
	{_State163, StructToken}:                             _GotoState32Action,
	{_State163, FuncToken}:                               _GotoState19Action,
	{_State163, VarToken}:                                _GotoState263Action,
	{_State163, LetToken}:                                _GotoState25Action,
	{_State163, NotToken}:                                _GotoState28Action,
	{_State163, LabelDeclToken}:                          _GotoState23Action,
	{_State163, LparenToken}:                             _GotoState26Action,
	{_State163, LbracketToken}:                           _GotoState24Action,
	{_State163, DotToken}:                                _GotoState262Action,
	{_State163, SubToken}:                                _GotoState33Action,
	{_State163, MulToken}:                                _GotoState27Action,
	{_State163, BitNegToken}:                             _GotoState16Action,
	{_State163, BitAndToken}:                             _GotoState15Action,
	{_State163, GreaterToken}:                            _GotoState20Action,
	{_State163, ParseErrorToken}:                         _GotoState29Action,
	{_State163, CasePatternsType}:                        _GotoState266Action,
	{_State163, VarDeclPatternType}:                      _GotoState70Action,
	{_State163, VarOrLetType}:                            _GotoState71Action,
	{_State163, AssignPatternType}:                       _GotoState264Action,
	{_State163, CasePatternType}:                         _GotoState265Action,
	{_State163, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State163, SequenceExprType}:                        _GotoState267Action,
	{_State163, CallExprType}:                            _GotoState49Action,
	{_State163, AtomExprType}:                            _GotoState42Action,
	{_State163, ParseErrorExprType}:                      _GotoState62Action,
	{_State163, LiteralExprType}:                         _GotoState57Action,
	{_State163, IdentifierExprType}:                      _GotoState52Action,
	{_State163, BlockExprType}:                           _GotoState48Action,
	{_State163, InitializeExprType}:                      _GotoState56Action,
	{_State163, ImplicitStructExprType}:                  _GotoState53Action,
	{_State163, AccessibleExprType}:                      _GotoState37Action,
	{_State163, AccessExprType}:                          _GotoState36Action,
	{_State163, IndexExprType}:                           _GotoState54Action,
	{_State163, PostfixableExprType}:                     _GotoState64Action,
	{_State163, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State163, PrefixableExprType}:                      _GotoState67Action,
	{_State163, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State163, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State163, MulExprType}:                             _GotoState59Action,
	{_State163, BinaryMulExprType}:                       _GotoState46Action,
	{_State163, AddExprType}:                             _GotoState38Action,
	{_State163, BinaryAddExprType}:                       _GotoState43Action,
	{_State163, CmpExprType}:                             _GotoState50Action,
	{_State163, BinaryCmpExprType}:                       _GotoState45Action,
	{_State163, AndExprType}:                             _GotoState39Action,
	{_State163, BinaryAndExprType}:                       _GotoState44Action,
	{_State163, OrExprType}:                              _GotoState61Action,
	{_State163, BinaryOrExprType}:                        _GotoState47Action,
	{_State163, InitializableTypeType}:                   _GotoState55Action,
	{_State163, SliceTypeType}:                           _GotoState69Action,
	{_State163, ArrayTypeType}:                           _GotoState41Action,
	{_State163, MapTypeType}:                             _GotoState58Action,
	{_State163, ExplicitStructDefType}:                   _GotoState51Action,
	{_State163, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State165, ColonToken}:                              _GotoState268Action,
	{_State168, StringLiteralToken}:                      _GotoState270Action,
	{_State168, LparenToken}:                             _GotoState269Action,
	{_State168, ImportClauseType}:                        _GotoState271Action,
	{_State170, LessToken}:                               _GotoState272Action,
	{_State171, LbracketToken}:                           _GotoState123Action,
	{_State171, DotToken}:                                _GotoState122Action,
	{_State171, QuestionToken}:                           _GotoState124Action,
	{_State171, DollarLbracketToken}:                     _GotoState121Action,
	{_State171, AddAssignToken}:                          _GotoState273Action,
	{_State171, SubAssignToken}:                          _GotoState284Action,
	{_State171, MulAssignToken}:                          _GotoState283Action,
	{_State171, DivAssignToken}:                          _GotoState281Action,
	{_State171, ModAssignToken}:                          _GotoState282Action,
	{_State171, AddOneAssignToken}:                       _GotoState274Action,
	{_State171, SubOneAssignToken}:                       _GotoState285Action,
	{_State171, BitNegAssignToken}:                       _GotoState277Action,
	{_State171, BitAndAssignToken}:                       _GotoState275Action,
	{_State171, BitOrAssignToken}:                        _GotoState278Action,
	{_State171, BitXorAssignToken}:                       _GotoState280Action,
	{_State171, BitLshiftAssignToken}:                    _GotoState276Action,
	{_State171, BitRshiftAssignToken}:                    _GotoState279Action,
	{_State171, UnaryOpAssignType}:                       _GotoState287Action,
	{_State171, BinaryOpAssignType}:                      _GotoState286Action,
	{_State171, OptionalGenericBindingType}:              _GotoState125Action,
	{_State172, AssignToken}:                             _GotoState288Action,
	{_State175, IntegerLiteralToken}:                     _GotoState22Action,
	{_State175, FloatLiteralToken}:                       _GotoState18Action,
	{_State175, RuneLiteralToken}:                        _GotoState30Action,
	{_State175, StringLiteralToken}:                      _GotoState31Action,
	{_State175, IdentifierToken}:                         _GotoState21Action,
	{_State175, TrueToken}:                               _GotoState34Action,
	{_State175, FalseToken}:                              _GotoState17Action,
	{_State175, StructToken}:                             _GotoState32Action,
	{_State175, FuncToken}:                               _GotoState19Action,
	{_State175, LabelDeclToken}:                          _GotoState23Action,
	{_State175, LparenToken}:                             _GotoState26Action,
	{_State175, LbracketToken}:                           _GotoState24Action,
	{_State175, ParseErrorToken}:                         _GotoState29Action,
	{_State175, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State175, CallExprType}:                            _GotoState290Action,
	{_State175, AtomExprType}:                            _GotoState42Action,
	{_State175, ParseErrorExprType}:                      _GotoState62Action,
	{_State175, LiteralExprType}:                         _GotoState57Action,
	{_State175, IdentifierExprType}:                      _GotoState52Action,
	{_State175, BlockExprType}:                           _GotoState48Action,
	{_State175, InitializeExprType}:                      _GotoState56Action,
	{_State175, ImplicitStructExprType}:                  _GotoState53Action,
	{_State175, AccessibleExprType}:                      _GotoState289Action,
	{_State175, AccessExprType}:                          _GotoState36Action,
	{_State175, IndexExprType}:                           _GotoState54Action,
	{_State175, InitializableTypeType}:                   _GotoState55Action,
	{_State175, SliceTypeType}:                           _GotoState69Action,
	{_State175, ArrayTypeType}:                           _GotoState41Action,
	{_State175, MapTypeType}:                             _GotoState58Action,
	{_State175, ExplicitStructDefType}:                   _GotoState51Action,
	{_State175, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State179, CommaToken}:                              _GotoState291Action,
	{_State183, IntegerLiteralToken}:                     _GotoState22Action,
	{_State183, FloatLiteralToken}:                       _GotoState18Action,
	{_State183, RuneLiteralToken}:                        _GotoState30Action,
	{_State183, StringLiteralToken}:                      _GotoState31Action,
	{_State183, IdentifierToken}:                         _GotoState21Action,
	{_State183, TrueToken}:                               _GotoState34Action,
	{_State183, FalseToken}:                              _GotoState17Action,
	{_State183, StructToken}:                             _GotoState32Action,
	{_State183, FuncToken}:                               _GotoState19Action,
	{_State183, VarToken}:                                _GotoState35Action,
	{_State183, LetToken}:                                _GotoState25Action,
	{_State183, NotToken}:                                _GotoState28Action,
	{_State183, LabelDeclToken}:                          _GotoState23Action,
	{_State183, JumpLabelToken}:                          _GotoState292Action,
	{_State183, LparenToken}:                             _GotoState26Action,
	{_State183, LbracketToken}:                           _GotoState24Action,
	{_State183, SubToken}:                                _GotoState33Action,
	{_State183, MulToken}:                                _GotoState27Action,
	{_State183, BitNegToken}:                             _GotoState16Action,
	{_State183, BitAndToken}:                             _GotoState15Action,
	{_State183, GreaterToken}:                            _GotoState20Action,
	{_State183, ParseErrorToken}:                         _GotoState29Action,
	{_State183, ExpressionsType}:                         _GotoState293Action,
	{_State183, VarDeclPatternType}:                      _GotoState70Action,
	{_State183, VarOrLetType}:                            _GotoState71Action,
	{_State183, ExpressionType}:                          _GotoState177Action,
	{_State183, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State183, SequenceExprType}:                        _GotoState68Action,
	{_State183, CallExprType}:                            _GotoState49Action,
	{_State183, AtomExprType}:                            _GotoState42Action,
	{_State183, ParseErrorExprType}:                      _GotoState62Action,
	{_State183, LiteralExprType}:                         _GotoState57Action,
	{_State183, IdentifierExprType}:                      _GotoState52Action,
	{_State183, BlockExprType}:                           _GotoState48Action,
	{_State183, InitializeExprType}:                      _GotoState56Action,
	{_State183, ImplicitStructExprType}:                  _GotoState53Action,
	{_State183, AccessibleExprType}:                      _GotoState37Action,
	{_State183, AccessExprType}:                          _GotoState36Action,
	{_State183, IndexExprType}:                           _GotoState54Action,
	{_State183, PostfixableExprType}:                     _GotoState64Action,
	{_State183, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State183, PrefixableExprType}:                      _GotoState67Action,
	{_State183, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State183, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State183, MulExprType}:                             _GotoState59Action,
	{_State183, BinaryMulExprType}:                       _GotoState46Action,
	{_State183, AddExprType}:                             _GotoState38Action,
	{_State183, BinaryAddExprType}:                       _GotoState43Action,
	{_State183, CmpExprType}:                             _GotoState50Action,
	{_State183, BinaryCmpExprType}:                       _GotoState45Action,
	{_State183, AndExprType}:                             _GotoState39Action,
	{_State183, BinaryAndExprType}:                       _GotoState44Action,
	{_State183, OrExprType}:                              _GotoState61Action,
	{_State183, BinaryOrExprType}:                        _GotoState47Action,
	{_State183, InitializableTypeType}:                   _GotoState55Action,
	{_State183, SliceTypeType}:                           _GotoState69Action,
	{_State183, ArrayTypeType}:                           _GotoState41Action,
	{_State183, MapTypeType}:                             _GotoState58Action,
	{_State183, ExplicitStructDefType}:                   _GotoState51Action,
	{_State183, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State184, NewlinesToken}:                           _GotoState294Action,
	{_State184, SemicolonToken}:                          _GotoState295Action,
	{_State188, RbraceToken}:                             _GotoState296Action,
	{_State192, IntegerLiteralToken}:                     _GotoState22Action,
	{_State192, FloatLiteralToken}:                       _GotoState18Action,
	{_State192, RuneLiteralToken}:                        _GotoState30Action,
	{_State192, StringLiteralToken}:                      _GotoState31Action,
	{_State192, IdentifierToken}:                         _GotoState21Action,
	{_State192, TrueToken}:                               _GotoState34Action,
	{_State192, FalseToken}:                              _GotoState17Action,
	{_State192, StructToken}:                             _GotoState32Action,
	{_State192, FuncToken}:                               _GotoState19Action,
	{_State192, VarToken}:                                _GotoState35Action,
	{_State192, LetToken}:                                _GotoState25Action,
	{_State192, NotToken}:                                _GotoState28Action,
	{_State192, LabelDeclToken}:                          _GotoState23Action,
	{_State192, LparenToken}:                             _GotoState26Action,
	{_State192, LbracketToken}:                           _GotoState24Action,
	{_State192, SubToken}:                                _GotoState33Action,
	{_State192, MulToken}:                                _GotoState27Action,
	{_State192, BitNegToken}:                             _GotoState16Action,
	{_State192, BitAndToken}:                             _GotoState15Action,
	{_State192, GreaterToken}:                            _GotoState20Action,
	{_State192, ParseErrorToken}:                         _GotoState29Action,
	{_State192, VarDeclPatternType}:                      _GotoState70Action,
	{_State192, VarOrLetType}:                            _GotoState71Action,
	{_State192, ExpressionType}:                          _GotoState297Action,
	{_State192, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State192, SequenceExprType}:                        _GotoState68Action,
	{_State192, CallExprType}:                            _GotoState49Action,
	{_State192, AtomExprType}:                            _GotoState42Action,
	{_State192, ParseErrorExprType}:                      _GotoState62Action,
	{_State192, LiteralExprType}:                         _GotoState57Action,
	{_State192, IdentifierExprType}:                      _GotoState52Action,
	{_State192, BlockExprType}:                           _GotoState48Action,
	{_State192, InitializeExprType}:                      _GotoState56Action,
	{_State192, ImplicitStructExprType}:                  _GotoState53Action,
	{_State192, AccessibleExprType}:                      _GotoState37Action,
	{_State192, AccessExprType}:                          _GotoState36Action,
	{_State192, IndexExprType}:                           _GotoState54Action,
	{_State192, PostfixableExprType}:                     _GotoState64Action,
	{_State192, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State192, PrefixableExprType}:                      _GotoState67Action,
	{_State192, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State192, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State192, MulExprType}:                             _GotoState59Action,
	{_State192, BinaryMulExprType}:                       _GotoState46Action,
	{_State192, AddExprType}:                             _GotoState38Action,
	{_State192, BinaryAddExprType}:                       _GotoState43Action,
	{_State192, CmpExprType}:                             _GotoState50Action,
	{_State192, BinaryCmpExprType}:                       _GotoState45Action,
	{_State192, AndExprType}:                             _GotoState39Action,
	{_State192, BinaryAndExprType}:                       _GotoState44Action,
	{_State192, OrExprType}:                              _GotoState61Action,
	{_State192, BinaryOrExprType}:                        _GotoState47Action,
	{_State192, InitializableTypeType}:                   _GotoState55Action,
	{_State192, SliceTypeType}:                           _GotoState69Action,
	{_State192, ArrayTypeType}:                           _GotoState41Action,
	{_State192, MapTypeType}:                             _GotoState58Action,
	{_State192, ExplicitStructDefType}:                   _GotoState51Action,
	{_State192, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State193, IdentifierToken}:                         _GotoState93Action,
	{_State193, StructToken}:                             _GotoState32Action,
	{_State193, EnumToken}:                               _GotoState90Action,
	{_State193, TraitToken}:                              _GotoState98Action,
	{_State193, FuncToken}:                               _GotoState92Action,
	{_State193, LparenToken}:                             _GotoState94Action,
	{_State193, LbracketToken}:                           _GotoState24Action,
	{_State193, DotToken}:                                _GotoState89Action,
	{_State193, QuestionToken}:                           _GotoState96Action,
	{_State193, ExclaimToken}:                            _GotoState91Action,
	{_State193, TildeTildeToken}:                         _GotoState97Action,
	{_State193, BitNegToken}:                             _GotoState88Action,
	{_State193, BitAndToken}:                             _GotoState87Action,
	{_State193, ParseErrorToken}:                         _GotoState95Action,
	{_State193, InitializableTypeType}:                   _GotoState104Action,
	{_State193, SliceTypeType}:                           _GotoState69Action,
	{_State193, ArrayTypeType}:                           _GotoState41Action,
	{_State193, MapTypeType}:                             _GotoState58Action,
	{_State193, AtomTypeType}:                            _GotoState99Action,
	{_State193, ParseErrorTypeType}:                      _GotoState105Action,
	{_State193, ReturnableTypeType}:                      _GotoState108Action,
	{_State193, PrefixedTypeType}:                        _GotoState107Action,
	{_State193, PrefixTypeOpType}:                        _GotoState106Action,
	{_State193, ValueTypeType}:                           _GotoState298Action,
	{_State193, TraitOpTypeType}:                         _GotoState110Action,
	{_State193, ImplicitStructDefType}:                   _GotoState103Action,
	{_State193, ExplicitStructDefType}:                   _GotoState51Action,
	{_State193, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State193, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State193, TraitDefType}:                            _GotoState109Action,
	{_State193, FuncTypeType}:                            _GotoState101Action,
	{_State194, IdentifierToken}:                         _GotoState299Action,
	{_State194, GenericParameterDefType}:                 _GotoState300Action,
	{_State194, ProperGenericParameterDefsType}:          _GotoState302Action,
	{_State194, GenericParameterDefsType}:                _GotoState301Action,
	{_State195, IdentifierToken}:                         _GotoState93Action,
	{_State195, StructToken}:                             _GotoState32Action,
	{_State195, EnumToken}:                               _GotoState90Action,
	{_State195, TraitToken}:                              _GotoState98Action,
	{_State195, FuncToken}:                               _GotoState92Action,
	{_State195, LparenToken}:                             _GotoState94Action,
	{_State195, LbracketToken}:                           _GotoState24Action,
	{_State195, DotToken}:                                _GotoState89Action,
	{_State195, QuestionToken}:                           _GotoState96Action,
	{_State195, ExclaimToken}:                            _GotoState91Action,
	{_State195, TildeTildeToken}:                         _GotoState97Action,
	{_State195, BitNegToken}:                             _GotoState88Action,
	{_State195, BitAndToken}:                             _GotoState87Action,
	{_State195, ParseErrorToken}:                         _GotoState95Action,
	{_State195, InitializableTypeType}:                   _GotoState104Action,
	{_State195, SliceTypeType}:                           _GotoState69Action,
	{_State195, ArrayTypeType}:                           _GotoState41Action,
	{_State195, MapTypeType}:                             _GotoState58Action,
	{_State195, AtomTypeType}:                            _GotoState99Action,
	{_State195, ParseErrorTypeType}:                      _GotoState105Action,
	{_State195, ReturnableTypeType}:                      _GotoState108Action,
	{_State195, PrefixedTypeType}:                        _GotoState107Action,
	{_State195, PrefixTypeOpType}:                        _GotoState106Action,
	{_State195, ValueTypeType}:                           _GotoState303Action,
	{_State195, TraitOpTypeType}:                         _GotoState110Action,
	{_State195, ImplicitStructDefType}:                   _GotoState103Action,
	{_State195, ExplicitStructDefType}:                   _GotoState51Action,
	{_State195, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State195, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State195, TraitDefType}:                            _GotoState109Action,
	{_State195, FuncTypeType}:                            _GotoState101Action,
	{_State196, IntegerLiteralToken}:                     _GotoState22Action,
	{_State196, FloatLiteralToken}:                       _GotoState18Action,
	{_State196, RuneLiteralToken}:                        _GotoState30Action,
	{_State196, StringLiteralToken}:                      _GotoState31Action,
	{_State196, IdentifierToken}:                         _GotoState21Action,
	{_State196, TrueToken}:                               _GotoState34Action,
	{_State196, FalseToken}:                              _GotoState17Action,
	{_State196, StructToken}:                             _GotoState32Action,
	{_State196, FuncToken}:                               _GotoState19Action,
	{_State196, VarToken}:                                _GotoState35Action,
	{_State196, LetToken}:                                _GotoState25Action,
	{_State196, NotToken}:                                _GotoState28Action,
	{_State196, LabelDeclToken}:                          _GotoState23Action,
	{_State196, LparenToken}:                             _GotoState26Action,
	{_State196, LbracketToken}:                           _GotoState24Action,
	{_State196, SubToken}:                                _GotoState33Action,
	{_State196, MulToken}:                                _GotoState27Action,
	{_State196, BitNegToken}:                             _GotoState16Action,
	{_State196, BitAndToken}:                             _GotoState15Action,
	{_State196, GreaterToken}:                            _GotoState20Action,
	{_State196, ParseErrorToken}:                         _GotoState29Action,
	{_State196, VarDeclPatternType}:                      _GotoState70Action,
	{_State196, VarOrLetType}:                            _GotoState71Action,
	{_State196, ExpressionType}:                          _GotoState304Action,
	{_State196, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State196, SequenceExprType}:                        _GotoState68Action,
	{_State196, CallExprType}:                            _GotoState49Action,
	{_State196, AtomExprType}:                            _GotoState42Action,
	{_State196, ParseErrorExprType}:                      _GotoState62Action,
	{_State196, LiteralExprType}:                         _GotoState57Action,
	{_State196, IdentifierExprType}:                      _GotoState52Action,
	{_State196, BlockExprType}:                           _GotoState48Action,
	{_State196, InitializeExprType}:                      _GotoState56Action,
	{_State196, ImplicitStructExprType}:                  _GotoState53Action,
	{_State196, AccessibleExprType}:                      _GotoState37Action,
	{_State196, AccessExprType}:                          _GotoState36Action,
	{_State196, IndexExprType}:                           _GotoState54Action,
	{_State196, PostfixableExprType}:                     _GotoState64Action,
	{_State196, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State196, PrefixableExprType}:                      _GotoState67Action,
	{_State196, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State196, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State196, MulExprType}:                             _GotoState59Action,
	{_State196, BinaryMulExprType}:                       _GotoState46Action,
	{_State196, AddExprType}:                             _GotoState38Action,
	{_State196, BinaryAddExprType}:                       _GotoState43Action,
	{_State196, CmpExprType}:                             _GotoState50Action,
	{_State196, BinaryCmpExprType}:                       _GotoState45Action,
	{_State196, AndExprType}:                             _GotoState39Action,
	{_State196, BinaryAndExprType}:                       _GotoState44Action,
	{_State196, OrExprType}:                              _GotoState61Action,
	{_State196, BinaryOrExprType}:                        _GotoState47Action,
	{_State196, InitializableTypeType}:                   _GotoState55Action,
	{_State196, SliceTypeType}:                           _GotoState69Action,
	{_State196, ArrayTypeType}:                           _GotoState41Action,
	{_State196, MapTypeType}:                             _GotoState58Action,
	{_State196, ExplicitStructDefType}:                   _GotoState51Action,
	{_State196, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State197, LparenToken}:                             _GotoState305Action,
	{_State198, IdentifierToken}:                         _GotoState93Action,
	{_State198, StructToken}:                             _GotoState32Action,
	{_State198, EnumToken}:                               _GotoState90Action,
	{_State198, TraitToken}:                              _GotoState98Action,
	{_State198, FuncToken}:                               _GotoState92Action,
	{_State198, LparenToken}:                             _GotoState94Action,
	{_State198, LbracketToken}:                           _GotoState24Action,
	{_State198, DotToken}:                                _GotoState89Action,
	{_State198, QuestionToken}:                           _GotoState96Action,
	{_State198, ExclaimToken}:                            _GotoState91Action,
	{_State198, EllipsisToken}:                           _GotoState306Action,
	{_State198, TildeTildeToken}:                         _GotoState97Action,
	{_State198, BitNegToken}:                             _GotoState88Action,
	{_State198, BitAndToken}:                             _GotoState87Action,
	{_State198, ParseErrorToken}:                         _GotoState95Action,
	{_State198, InitializableTypeType}:                   _GotoState104Action,
	{_State198, SliceTypeType}:                           _GotoState69Action,
	{_State198, ArrayTypeType}:                           _GotoState41Action,
	{_State198, MapTypeType}:                             _GotoState58Action,
	{_State198, AtomTypeType}:                            _GotoState99Action,
	{_State198, ParseErrorTypeType}:                      _GotoState105Action,
	{_State198, ReturnableTypeType}:                      _GotoState108Action,
	{_State198, PrefixedTypeType}:                        _GotoState107Action,
	{_State198, PrefixTypeOpType}:                        _GotoState106Action,
	{_State198, ValueTypeType}:                           _GotoState307Action,
	{_State198, TraitOpTypeType}:                         _GotoState110Action,
	{_State198, ImplicitStructDefType}:                   _GotoState103Action,
	{_State198, ExplicitStructDefType}:                   _GotoState51Action,
	{_State198, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State198, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State198, TraitDefType}:                            _GotoState109Action,
	{_State198, FuncTypeType}:                            _GotoState101Action,
	{_State199, RparenToken}:                             _GotoState308Action,
	{_State201, RparenToken}:                             _GotoState309Action,
	{_State202, CommaToken}:                              _GotoState310Action,
	{_State204, IdentifierToken}:                         _GotoState208Action,
	{_State204, UnsafeToken}:                             _GotoState170Action,
	{_State204, StructToken}:                             _GotoState32Action,
	{_State204, EnumToken}:                               _GotoState90Action,
	{_State204, TraitToken}:                              _GotoState98Action,
	{_State204, FuncToken}:                               _GotoState92Action,
	{_State204, LparenToken}:                             _GotoState94Action,
	{_State204, LbracketToken}:                           _GotoState24Action,
	{_State204, DotToken}:                                _GotoState89Action,
	{_State204, QuestionToken}:                           _GotoState96Action,
	{_State204, ExclaimToken}:                            _GotoState91Action,
	{_State204, TildeTildeToken}:                         _GotoState97Action,
	{_State204, BitNegToken}:                             _GotoState88Action,
	{_State204, BitAndToken}:                             _GotoState87Action,
	{_State204, ParseErrorToken}:                         _GotoState95Action,
	{_State204, UnsafeStatementType}:                     _GotoState214Action,
	{_State204, InitializableTypeType}:                   _GotoState104Action,
	{_State204, SliceTypeType}:                           _GotoState69Action,
	{_State204, ArrayTypeType}:                           _GotoState41Action,
	{_State204, MapTypeType}:                             _GotoState58Action,
	{_State204, AtomTypeType}:                            _GotoState99Action,
	{_State204, ParseErrorTypeType}:                      _GotoState105Action,
	{_State204, ReturnableTypeType}:                      _GotoState108Action,
	{_State204, PrefixedTypeType}:                        _GotoState107Action,
	{_State204, PrefixTypeOpType}:                        _GotoState106Action,
	{_State204, ValueTypeType}:                           _GotoState215Action,
	{_State204, TraitOpTypeType}:                         _GotoState110Action,
	{_State204, FieldDefType}:                            _GotoState313Action,
	{_State204, ImplicitStructDefType}:                   _GotoState103Action,
	{_State204, ExplicitStructDefType}:                   _GotoState51Action,
	{_State204, EnumValueDefType}:                        _GotoState311Action,
	{_State204, ImplicitEnumValueDefsType}:               _GotoState314Action,
	{_State204, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State204, ExplicitEnumValueDefsType}:               _GotoState312Action,
	{_State204, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State204, TraitDefType}:                            _GotoState109Action,
	{_State204, FuncTypeType}:                            _GotoState101Action,
	{_State205, IdentifierToken}:                         _GotoState316Action,
	{_State205, StructToken}:                             _GotoState32Action,
	{_State205, EnumToken}:                               _GotoState90Action,
	{_State205, TraitToken}:                              _GotoState98Action,
	{_State205, FuncToken}:                               _GotoState92Action,
	{_State205, LparenToken}:                             _GotoState94Action,
	{_State205, LbracketToken}:                           _GotoState24Action,
	{_State205, DotToken}:                                _GotoState89Action,
	{_State205, QuestionToken}:                           _GotoState96Action,
	{_State205, ExclaimToken}:                            _GotoState91Action,
	{_State205, EllipsisToken}:                           _GotoState315Action,
	{_State205, TildeTildeToken}:                         _GotoState97Action,
	{_State205, BitNegToken}:                             _GotoState88Action,
	{_State205, BitAndToken}:                             _GotoState87Action,
	{_State205, ParseErrorToken}:                         _GotoState95Action,
	{_State205, InitializableTypeType}:                   _GotoState104Action,
	{_State205, SliceTypeType}:                           _GotoState69Action,
	{_State205, ArrayTypeType}:                           _GotoState41Action,
	{_State205, MapTypeType}:                             _GotoState58Action,
	{_State205, AtomTypeType}:                            _GotoState99Action,
	{_State205, ParseErrorTypeType}:                      _GotoState105Action,
	{_State205, ReturnableTypeType}:                      _GotoState108Action,
	{_State205, PrefixedTypeType}:                        _GotoState107Action,
	{_State205, PrefixTypeOpType}:                        _GotoState106Action,
	{_State205, ValueTypeType}:                           _GotoState320Action,
	{_State205, TraitOpTypeType}:                         _GotoState110Action,
	{_State205, ImplicitStructDefType}:                   _GotoState103Action,
	{_State205, ExplicitStructDefType}:                   _GotoState51Action,
	{_State205, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State205, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State205, TraitDefType}:                            _GotoState109Action,
	{_State205, ParameterDeclType}:                       _GotoState317Action,
	{_State205, ProperParameterDeclsType}:                _GotoState319Action,
	{_State205, ParameterDeclsType}:                      _GotoState318Action,
	{_State205, FuncTypeType}:                            _GotoState101Action,
	{_State206, IdentifierToken}:                         _GotoState321Action,
	{_State208, IdentifierToken}:                         _GotoState93Action,
	{_State208, StructToken}:                             _GotoState32Action,
	{_State208, EnumToken}:                               _GotoState90Action,
	{_State208, TraitToken}:                              _GotoState98Action,
	{_State208, FuncToken}:                               _GotoState92Action,
	{_State208, LparenToken}:                             _GotoState94Action,
	{_State208, LbracketToken}:                           _GotoState24Action,
	{_State208, DotToken}:                                _GotoState322Action,
	{_State208, QuestionToken}:                           _GotoState96Action,
	{_State208, ExclaimToken}:                            _GotoState91Action,
	{_State208, DollarLbracketToken}:                     _GotoState121Action,
	{_State208, TildeTildeToken}:                         _GotoState97Action,
	{_State208, BitNegToken}:                             _GotoState88Action,
	{_State208, BitAndToken}:                             _GotoState87Action,
	{_State208, ParseErrorToken}:                         _GotoState95Action,
	{_State208, OptionalGenericBindingType}:              _GotoState207Action,
	{_State208, InitializableTypeType}:                   _GotoState104Action,
	{_State208, SliceTypeType}:                           _GotoState69Action,
	{_State208, ArrayTypeType}:                           _GotoState41Action,
	{_State208, MapTypeType}:                             _GotoState58Action,
	{_State208, AtomTypeType}:                            _GotoState99Action,
	{_State208, ParseErrorTypeType}:                      _GotoState105Action,
	{_State208, ReturnableTypeType}:                      _GotoState108Action,
	{_State208, PrefixedTypeType}:                        _GotoState107Action,
	{_State208, PrefixTypeOpType}:                        _GotoState106Action,
	{_State208, ValueTypeType}:                           _GotoState323Action,
	{_State208, TraitOpTypeType}:                         _GotoState110Action,
	{_State208, ImplicitStructDefType}:                   _GotoState103Action,
	{_State208, ExplicitStructDefType}:                   _GotoState51Action,
	{_State208, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State208, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State208, TraitDefType}:                            _GotoState109Action,
	{_State208, FuncTypeType}:                            _GotoState101Action,
	{_State209, OrToken}:                                 _GotoState324Action,
	{_State210, AssignToken}:                             _GotoState325Action,
	{_State211, OrToken}:                                 _GotoState326Action,
	{_State211, RparenToken}:                             _GotoState327Action,
	{_State212, RparenToken}:                             _GotoState328Action,
	{_State213, CommaToken}:                              _GotoState329Action,
	{_State215, AddToken}:                                _GotoState218Action,
	{_State215, SubToken}:                                _GotoState223Action,
	{_State215, MulToken}:                                _GotoState221Action,
	{_State215, TraitOpType}:                             _GotoState224Action,
	{_State216, IdentifierToken}:                         _GotoState208Action,
	{_State216, UnsafeToken}:                             _GotoState170Action,
	{_State216, StructToken}:                             _GotoState32Action,
	{_State216, EnumToken}:                               _GotoState90Action,
	{_State216, TraitToken}:                              _GotoState98Action,
	{_State216, FuncToken}:                               _GotoState330Action,
	{_State216, LparenToken}:                             _GotoState94Action,
	{_State216, LbracketToken}:                           _GotoState24Action,
	{_State216, DotToken}:                                _GotoState89Action,
	{_State216, QuestionToken}:                           _GotoState96Action,
	{_State216, ExclaimToken}:                            _GotoState91Action,
	{_State216, TildeTildeToken}:                         _GotoState97Action,
	{_State216, BitNegToken}:                             _GotoState88Action,
	{_State216, BitAndToken}:                             _GotoState87Action,
	{_State216, ParseErrorToken}:                         _GotoState95Action,
	{_State216, UnsafeStatementType}:                     _GotoState214Action,
	{_State216, InitializableTypeType}:                   _GotoState104Action,
	{_State216, SliceTypeType}:                           _GotoState69Action,
	{_State216, ArrayTypeType}:                           _GotoState41Action,
	{_State216, MapTypeType}:                             _GotoState58Action,
	{_State216, AtomTypeType}:                            _GotoState99Action,
	{_State216, ParseErrorTypeType}:                      _GotoState105Action,
	{_State216, ReturnableTypeType}:                      _GotoState108Action,
	{_State216, PrefixedTypeType}:                        _GotoState107Action,
	{_State216, PrefixTypeOpType}:                        _GotoState106Action,
	{_State216, ValueTypeType}:                           _GotoState215Action,
	{_State216, TraitOpTypeType}:                         _GotoState110Action,
	{_State216, FieldDefType}:                            _GotoState331Action,
	{_State216, ImplicitStructDefType}:                   _GotoState103Action,
	{_State216, ExplicitStructDefType}:                   _GotoState51Action,
	{_State216, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State216, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State216, TraitPropertyType}:                       _GotoState335Action,
	{_State216, ProperTraitPropertiesType}:               _GotoState333Action,
	{_State216, TraitPropertiesType}:                     _GotoState334Action,
	{_State216, TraitDefType}:                            _GotoState109Action,
	{_State216, FuncTypeType}:                            _GotoState101Action,
	{_State216, MethodSignatureType}:                     _GotoState332Action,
	{_State219, IdentifierToken}:                         _GotoState93Action,
	{_State219, StructToken}:                             _GotoState32Action,
	{_State219, EnumToken}:                               _GotoState90Action,
	{_State219, TraitToken}:                              _GotoState98Action,
	{_State219, FuncToken}:                               _GotoState92Action,
	{_State219, LparenToken}:                             _GotoState94Action,
	{_State219, LbracketToken}:                           _GotoState24Action,
	{_State219, DotToken}:                                _GotoState89Action,
	{_State219, QuestionToken}:                           _GotoState96Action,
	{_State219, ExclaimToken}:                            _GotoState91Action,
	{_State219, TildeTildeToken}:                         _GotoState97Action,
	{_State219, BitNegToken}:                             _GotoState88Action,
	{_State219, BitAndToken}:                             _GotoState87Action,
	{_State219, ParseErrorToken}:                         _GotoState95Action,
	{_State219, InitializableTypeType}:                   _GotoState104Action,
	{_State219, SliceTypeType}:                           _GotoState69Action,
	{_State219, ArrayTypeType}:                           _GotoState41Action,
	{_State219, MapTypeType}:                             _GotoState58Action,
	{_State219, AtomTypeType}:                            _GotoState99Action,
	{_State219, ParseErrorTypeType}:                      _GotoState105Action,
	{_State219, ReturnableTypeType}:                      _GotoState108Action,
	{_State219, PrefixedTypeType}:                        _GotoState107Action,
	{_State219, PrefixTypeOpType}:                        _GotoState106Action,
	{_State219, ValueTypeType}:                           _GotoState336Action,
	{_State219, TraitOpTypeType}:                         _GotoState110Action,
	{_State219, ImplicitStructDefType}:                   _GotoState103Action,
	{_State219, ExplicitStructDefType}:                   _GotoState51Action,
	{_State219, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State219, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State219, TraitDefType}:                            _GotoState109Action,
	{_State219, FuncTypeType}:                            _GotoState101Action,
	{_State220, IntegerLiteralToken}:                     _GotoState337Action,
	{_State224, IdentifierToken}:                         _GotoState93Action,
	{_State224, StructToken}:                             _GotoState32Action,
	{_State224, EnumToken}:                               _GotoState90Action,
	{_State224, TraitToken}:                              _GotoState98Action,
	{_State224, FuncToken}:                               _GotoState92Action,
	{_State224, LparenToken}:                             _GotoState94Action,
	{_State224, LbracketToken}:                           _GotoState24Action,
	{_State224, DotToken}:                                _GotoState89Action,
	{_State224, QuestionToken}:                           _GotoState96Action,
	{_State224, ExclaimToken}:                            _GotoState91Action,
	{_State224, TildeTildeToken}:                         _GotoState97Action,
	{_State224, BitNegToken}:                             _GotoState88Action,
	{_State224, BitAndToken}:                             _GotoState87Action,
	{_State224, ParseErrorToken}:                         _GotoState95Action,
	{_State224, InitializableTypeType}:                   _GotoState104Action,
	{_State224, SliceTypeType}:                           _GotoState69Action,
	{_State224, ArrayTypeType}:                           _GotoState41Action,
	{_State224, MapTypeType}:                             _GotoState58Action,
	{_State224, AtomTypeType}:                            _GotoState99Action,
	{_State224, ParseErrorTypeType}:                      _GotoState105Action,
	{_State224, ReturnableTypeType}:                      _GotoState338Action,
	{_State224, PrefixedTypeType}:                        _GotoState107Action,
	{_State224, PrefixTypeOpType}:                        _GotoState106Action,
	{_State224, ImplicitStructDefType}:                   _GotoState103Action,
	{_State224, ExplicitStructDefType}:                   _GotoState51Action,
	{_State224, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State224, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State224, TraitDefType}:                            _GotoState109Action,
	{_State224, FuncTypeType}:                            _GotoState101Action,
	{_State226, IntegerLiteralToken}:                     _GotoState22Action,
	{_State226, FloatLiteralToken}:                       _GotoState18Action,
	{_State226, RuneLiteralToken}:                        _GotoState30Action,
	{_State226, StringLiteralToken}:                      _GotoState31Action,
	{_State226, IdentifierToken}:                         _GotoState21Action,
	{_State226, TrueToken}:                               _GotoState34Action,
	{_State226, FalseToken}:                              _GotoState17Action,
	{_State226, StructToken}:                             _GotoState32Action,
	{_State226, FuncToken}:                               _GotoState19Action,
	{_State226, VarToken}:                                _GotoState35Action,
	{_State226, LetToken}:                                _GotoState25Action,
	{_State226, NotToken}:                                _GotoState28Action,
	{_State226, LabelDeclToken}:                          _GotoState23Action,
	{_State226, LparenToken}:                             _GotoState26Action,
	{_State226, LbracketToken}:                           _GotoState24Action,
	{_State226, SubToken}:                                _GotoState33Action,
	{_State226, MulToken}:                                _GotoState27Action,
	{_State226, BitNegToken}:                             _GotoState16Action,
	{_State226, BitAndToken}:                             _GotoState15Action,
	{_State226, GreaterToken}:                            _GotoState20Action,
	{_State226, ParseErrorToken}:                         _GotoState29Action,
	{_State226, VarDeclPatternType}:                      _GotoState70Action,
	{_State226, VarOrLetType}:                            _GotoState71Action,
	{_State226, ExpressionType}:                          _GotoState339Action,
	{_State226, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State226, SequenceExprType}:                        _GotoState68Action,
	{_State226, CallExprType}:                            _GotoState49Action,
	{_State226, AtomExprType}:                            _GotoState42Action,
	{_State226, ParseErrorExprType}:                      _GotoState62Action,
	{_State226, LiteralExprType}:                         _GotoState57Action,
	{_State226, IdentifierExprType}:                      _GotoState52Action,
	{_State226, BlockExprType}:                           _GotoState48Action,
	{_State226, InitializeExprType}:                      _GotoState56Action,
	{_State226, ImplicitStructExprType}:                  _GotoState53Action,
	{_State226, AccessibleExprType}:                      _GotoState37Action,
	{_State226, AccessExprType}:                          _GotoState36Action,
	{_State226, IndexExprType}:                           _GotoState54Action,
	{_State226, PostfixableExprType}:                     _GotoState64Action,
	{_State226, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State226, PrefixableExprType}:                      _GotoState67Action,
	{_State226, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State226, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State226, MulExprType}:                             _GotoState59Action,
	{_State226, BinaryMulExprType}:                       _GotoState46Action,
	{_State226, AddExprType}:                             _GotoState38Action,
	{_State226, BinaryAddExprType}:                       _GotoState43Action,
	{_State226, CmpExprType}:                             _GotoState50Action,
	{_State226, BinaryCmpExprType}:                       _GotoState45Action,
	{_State226, AndExprType}:                             _GotoState39Action,
	{_State226, BinaryAndExprType}:                       _GotoState44Action,
	{_State226, OrExprType}:                              _GotoState61Action,
	{_State226, BinaryOrExprType}:                        _GotoState47Action,
	{_State226, InitializableTypeType}:                   _GotoState55Action,
	{_State226, SliceTypeType}:                           _GotoState69Action,
	{_State226, ArrayTypeType}:                           _GotoState41Action,
	{_State226, MapTypeType}:                             _GotoState58Action,
	{_State226, ExplicitStructDefType}:                   _GotoState51Action,
	{_State226, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State228, IntegerLiteralToken}:                     _GotoState22Action,
	{_State228, FloatLiteralToken}:                       _GotoState18Action,
	{_State228, RuneLiteralToken}:                        _GotoState30Action,
	{_State228, StringLiteralToken}:                      _GotoState31Action,
	{_State228, IdentifierToken}:                         _GotoState21Action,
	{_State228, TrueToken}:                               _GotoState34Action,
	{_State228, FalseToken}:                              _GotoState17Action,
	{_State228, StructToken}:                             _GotoState32Action,
	{_State228, FuncToken}:                               _GotoState19Action,
	{_State228, VarToken}:                                _GotoState35Action,
	{_State228, LetToken}:                                _GotoState25Action,
	{_State228, NotToken}:                                _GotoState28Action,
	{_State228, LabelDeclToken}:                          _GotoState23Action,
	{_State228, LparenToken}:                             _GotoState26Action,
	{_State228, LbracketToken}:                           _GotoState24Action,
	{_State228, SubToken}:                                _GotoState33Action,
	{_State228, MulToken}:                                _GotoState27Action,
	{_State228, BitNegToken}:                             _GotoState16Action,
	{_State228, BitAndToken}:                             _GotoState15Action,
	{_State228, GreaterToken}:                            _GotoState20Action,
	{_State228, ParseErrorToken}:                         _GotoState29Action,
	{_State228, VarDeclPatternType}:                      _GotoState70Action,
	{_State228, VarOrLetType}:                            _GotoState71Action,
	{_State228, ExpressionType}:                          _GotoState340Action,
	{_State228, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State228, SequenceExprType}:                        _GotoState68Action,
	{_State228, CallExprType}:                            _GotoState49Action,
	{_State228, AtomExprType}:                            _GotoState42Action,
	{_State228, ParseErrorExprType}:                      _GotoState62Action,
	{_State228, LiteralExprType}:                         _GotoState57Action,
	{_State228, IdentifierExprType}:                      _GotoState52Action,
	{_State228, BlockExprType}:                           _GotoState48Action,
	{_State228, InitializeExprType}:                      _GotoState56Action,
	{_State228, ImplicitStructExprType}:                  _GotoState53Action,
	{_State228, AccessibleExprType}:                      _GotoState37Action,
	{_State228, AccessExprType}:                          _GotoState36Action,
	{_State228, IndexExprType}:                           _GotoState54Action,
	{_State228, PostfixableExprType}:                     _GotoState64Action,
	{_State228, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State228, PrefixableExprType}:                      _GotoState67Action,
	{_State228, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State228, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State228, MulExprType}:                             _GotoState59Action,
	{_State228, BinaryMulExprType}:                       _GotoState46Action,
	{_State228, AddExprType}:                             _GotoState38Action,
	{_State228, BinaryAddExprType}:                       _GotoState43Action,
	{_State228, CmpExprType}:                             _GotoState50Action,
	{_State228, BinaryCmpExprType}:                       _GotoState45Action,
	{_State228, AndExprType}:                             _GotoState39Action,
	{_State228, BinaryAndExprType}:                       _GotoState44Action,
	{_State228, OrExprType}:                              _GotoState61Action,
	{_State228, BinaryOrExprType}:                        _GotoState47Action,
	{_State228, InitializableTypeType}:                   _GotoState55Action,
	{_State228, SliceTypeType}:                           _GotoState69Action,
	{_State228, ArrayTypeType}:                           _GotoState41Action,
	{_State228, MapTypeType}:                             _GotoState58Action,
	{_State228, ExplicitStructDefType}:                   _GotoState51Action,
	{_State228, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State229, IntegerLiteralToken}:                     _GotoState22Action,
	{_State229, FloatLiteralToken}:                       _GotoState18Action,
	{_State229, RuneLiteralToken}:                        _GotoState30Action,
	{_State229, StringLiteralToken}:                      _GotoState31Action,
	{_State229, IdentifierToken}:                         _GotoState21Action,
	{_State229, TrueToken}:                               _GotoState34Action,
	{_State229, FalseToken}:                              _GotoState17Action,
	{_State229, StructToken}:                             _GotoState32Action,
	{_State229, FuncToken}:                               _GotoState19Action,
	{_State229, VarToken}:                                _GotoState35Action,
	{_State229, LetToken}:                                _GotoState25Action,
	{_State229, NotToken}:                                _GotoState28Action,
	{_State229, LabelDeclToken}:                          _GotoState23Action,
	{_State229, LparenToken}:                             _GotoState26Action,
	{_State229, LbracketToken}:                           _GotoState24Action,
	{_State229, SubToken}:                                _GotoState33Action,
	{_State229, MulToken}:                                _GotoState27Action,
	{_State229, BitNegToken}:                             _GotoState16Action,
	{_State229, BitAndToken}:                             _GotoState15Action,
	{_State229, GreaterToken}:                            _GotoState20Action,
	{_State229, ParseErrorToken}:                         _GotoState29Action,
	{_State229, VarDeclPatternType}:                      _GotoState70Action,
	{_State229, VarOrLetType}:                            _GotoState71Action,
	{_State229, ExpressionType}:                          _GotoState341Action,
	{_State229, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State229, SequenceExprType}:                        _GotoState68Action,
	{_State229, CallExprType}:                            _GotoState49Action,
	{_State229, AtomExprType}:                            _GotoState42Action,
	{_State229, ParseErrorExprType}:                      _GotoState62Action,
	{_State229, LiteralExprType}:                         _GotoState57Action,
	{_State229, IdentifierExprType}:                      _GotoState52Action,
	{_State229, BlockExprType}:                           _GotoState48Action,
	{_State229, InitializeExprType}:                      _GotoState56Action,
	{_State229, ImplicitStructExprType}:                  _GotoState53Action,
	{_State229, AccessibleExprType}:                      _GotoState37Action,
	{_State229, AccessExprType}:                          _GotoState36Action,
	{_State229, IndexExprType}:                           _GotoState54Action,
	{_State229, PostfixableExprType}:                     _GotoState64Action,
	{_State229, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State229, PrefixableExprType}:                      _GotoState67Action,
	{_State229, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State229, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State229, MulExprType}:                             _GotoState59Action,
	{_State229, BinaryMulExprType}:                       _GotoState46Action,
	{_State229, AddExprType}:                             _GotoState38Action,
	{_State229, BinaryAddExprType}:                       _GotoState43Action,
	{_State229, CmpExprType}:                             _GotoState50Action,
	{_State229, BinaryCmpExprType}:                       _GotoState45Action,
	{_State229, AndExprType}:                             _GotoState39Action,
	{_State229, BinaryAndExprType}:                       _GotoState44Action,
	{_State229, OrExprType}:                              _GotoState61Action,
	{_State229, BinaryOrExprType}:                        _GotoState47Action,
	{_State229, InitializableTypeType}:                   _GotoState55Action,
	{_State229, SliceTypeType}:                           _GotoState69Action,
	{_State229, ArrayTypeType}:                           _GotoState41Action,
	{_State229, MapTypeType}:                             _GotoState58Action,
	{_State229, ExplicitStructDefType}:                   _GotoState51Action,
	{_State229, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State231, IntegerLiteralToken}:                     _GotoState22Action,
	{_State231, FloatLiteralToken}:                       _GotoState18Action,
	{_State231, RuneLiteralToken}:                        _GotoState30Action,
	{_State231, StringLiteralToken}:                      _GotoState31Action,
	{_State231, IdentifierToken}:                         _GotoState114Action,
	{_State231, TrueToken}:                               _GotoState34Action,
	{_State231, FalseToken}:                              _GotoState17Action,
	{_State231, StructToken}:                             _GotoState32Action,
	{_State231, FuncToken}:                               _GotoState19Action,
	{_State231, VarToken}:                                _GotoState35Action,
	{_State231, LetToken}:                                _GotoState25Action,
	{_State231, NotToken}:                                _GotoState28Action,
	{_State231, LabelDeclToken}:                          _GotoState23Action,
	{_State231, LparenToken}:                             _GotoState26Action,
	{_State231, LbracketToken}:                           _GotoState24Action,
	{_State231, ColonToken}:                              _GotoState112Action,
	{_State231, EllipsisToken}:                           _GotoState113Action,
	{_State231, SubToken}:                                _GotoState33Action,
	{_State231, MulToken}:                                _GotoState27Action,
	{_State231, BitNegToken}:                             _GotoState16Action,
	{_State231, BitAndToken}:                             _GotoState15Action,
	{_State231, GreaterToken}:                            _GotoState20Action,
	{_State231, ParseErrorToken}:                         _GotoState29Action,
	{_State231, VarDeclPatternType}:                      _GotoState70Action,
	{_State231, VarOrLetType}:                            _GotoState71Action,
	{_State231, ExpressionType}:                          _GotoState118Action,
	{_State231, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State231, SequenceExprType}:                        _GotoState68Action,
	{_State231, CallExprType}:                            _GotoState49Action,
	{_State231, ArgumentType}:                            _GotoState342Action,
	{_State231, ColonExprType}:                           _GotoState117Action,
	{_State231, AtomExprType}:                            _GotoState42Action,
	{_State231, ParseErrorExprType}:                      _GotoState62Action,
	{_State231, LiteralExprType}:                         _GotoState57Action,
	{_State231, IdentifierExprType}:                      _GotoState52Action,
	{_State231, BlockExprType}:                           _GotoState48Action,
	{_State231, InitializeExprType}:                      _GotoState56Action,
	{_State231, ImplicitStructExprType}:                  _GotoState53Action,
	{_State231, AccessibleExprType}:                      _GotoState37Action,
	{_State231, AccessExprType}:                          _GotoState36Action,
	{_State231, IndexExprType}:                           _GotoState54Action,
	{_State231, PostfixableExprType}:                     _GotoState64Action,
	{_State231, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State231, PrefixableExprType}:                      _GotoState67Action,
	{_State231, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State231, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State231, MulExprType}:                             _GotoState59Action,
	{_State231, BinaryMulExprType}:                       _GotoState46Action,
	{_State231, AddExprType}:                             _GotoState38Action,
	{_State231, BinaryAddExprType}:                       _GotoState43Action,
	{_State231, CmpExprType}:                             _GotoState50Action,
	{_State231, BinaryCmpExprType}:                       _GotoState45Action,
	{_State231, AndExprType}:                             _GotoState39Action,
	{_State231, BinaryAndExprType}:                       _GotoState44Action,
	{_State231, OrExprType}:                              _GotoState61Action,
	{_State231, BinaryOrExprType}:                        _GotoState47Action,
	{_State231, InitializableTypeType}:                   _GotoState55Action,
	{_State231, SliceTypeType}:                           _GotoState69Action,
	{_State231, ArrayTypeType}:                           _GotoState41Action,
	{_State231, MapTypeType}:                             _GotoState58Action,
	{_State231, ExplicitStructDefType}:                   _GotoState51Action,
	{_State231, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State232, RparenToken}:                             _GotoState343Action,
	{_State234, NewlinesToken}:                           _GotoState345Action,
	{_State234, CommaToken}:                              _GotoState344Action,
	{_State235, RbracketToken}:                           _GotoState346Action,
	{_State236, CommaToken}:                              _GotoState347Action,
	{_State237, AddToken}:                                _GotoState218Action,
	{_State237, SubToken}:                                _GotoState223Action,
	{_State237, MulToken}:                                _GotoState221Action,
	{_State237, TraitOpType}:                             _GotoState224Action,
	{_State239, RbracketToken}:                           _GotoState348Action,
	{_State240, IntegerLiteralToken}:                     _GotoState22Action,
	{_State240, FloatLiteralToken}:                       _GotoState18Action,
	{_State240, RuneLiteralToken}:                        _GotoState30Action,
	{_State240, StringLiteralToken}:                      _GotoState31Action,
	{_State240, IdentifierToken}:                         _GotoState114Action,
	{_State240, TrueToken}:                               _GotoState34Action,
	{_State240, FalseToken}:                              _GotoState17Action,
	{_State240, StructToken}:                             _GotoState32Action,
	{_State240, FuncToken}:                               _GotoState19Action,
	{_State240, VarToken}:                                _GotoState35Action,
	{_State240, LetToken}:                                _GotoState25Action,
	{_State240, NotToken}:                                _GotoState28Action,
	{_State240, LabelDeclToken}:                          _GotoState23Action,
	{_State240, LparenToken}:                             _GotoState26Action,
	{_State240, LbracketToken}:                           _GotoState24Action,
	{_State240, ColonToken}:                              _GotoState112Action,
	{_State240, EllipsisToken}:                           _GotoState113Action,
	{_State240, SubToken}:                                _GotoState33Action,
	{_State240, MulToken}:                                _GotoState27Action,
	{_State240, BitNegToken}:                             _GotoState16Action,
	{_State240, BitAndToken}:                             _GotoState15Action,
	{_State240, GreaterToken}:                            _GotoState20Action,
	{_State240, ParseErrorToken}:                         _GotoState29Action,
	{_State240, VarDeclPatternType}:                      _GotoState70Action,
	{_State240, VarOrLetType}:                            _GotoState71Action,
	{_State240, ExpressionType}:                          _GotoState118Action,
	{_State240, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State240, SequenceExprType}:                        _GotoState68Action,
	{_State240, CallExprType}:                            _GotoState49Action,
	{_State240, ProperArgumentsType}:                     _GotoState119Action,
	{_State240, ArgumentsType}:                           _GotoState349Action,
	{_State240, ArgumentType}:                            _GotoState115Action,
	{_State240, ColonExprType}:                           _GotoState117Action,
	{_State240, AtomExprType}:                            _GotoState42Action,
	{_State240, ParseErrorExprType}:                      _GotoState62Action,
	{_State240, LiteralExprType}:                         _GotoState57Action,
	{_State240, IdentifierExprType}:                      _GotoState52Action,
	{_State240, BlockExprType}:                           _GotoState48Action,
	{_State240, InitializeExprType}:                      _GotoState56Action,
	{_State240, ImplicitStructExprType}:                  _GotoState53Action,
	{_State240, AccessibleExprType}:                      _GotoState37Action,
	{_State240, AccessExprType}:                          _GotoState36Action,
	{_State240, IndexExprType}:                           _GotoState54Action,
	{_State240, PostfixableExprType}:                     _GotoState64Action,
	{_State240, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State240, PrefixableExprType}:                      _GotoState67Action,
	{_State240, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State240, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State240, MulExprType}:                             _GotoState59Action,
	{_State240, BinaryMulExprType}:                       _GotoState46Action,
	{_State240, AddExprType}:                             _GotoState38Action,
	{_State240, BinaryAddExprType}:                       _GotoState43Action,
	{_State240, CmpExprType}:                             _GotoState50Action,
	{_State240, BinaryCmpExprType}:                       _GotoState45Action,
	{_State240, AndExprType}:                             _GotoState39Action,
	{_State240, BinaryAndExprType}:                       _GotoState44Action,
	{_State240, OrExprType}:                              _GotoState61Action,
	{_State240, BinaryOrExprType}:                        _GotoState47Action,
	{_State240, InitializableTypeType}:                   _GotoState55Action,
	{_State240, SliceTypeType}:                           _GotoState69Action,
	{_State240, ArrayTypeType}:                           _GotoState41Action,
	{_State240, MapTypeType}:                             _GotoState58Action,
	{_State240, ExplicitStructDefType}:                   _GotoState51Action,
	{_State240, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State241, MulToken}:                                _GotoState145Action,
	{_State241, DivToken}:                                _GotoState143Action,
	{_State241, ModToken}:                                _GotoState144Action,
	{_State241, BitAndToken}:                             _GotoState140Action,
	{_State241, BitLshiftToken}:                          _GotoState141Action,
	{_State241, BitRshiftToken}:                          _GotoState142Action,
	{_State241, MulOpType}:                               _GotoState146Action,
	{_State242, EqualToken}:                              _GotoState132Action,
	{_State242, NotEqualToken}:                           _GotoState137Action,
	{_State242, LessToken}:                               _GotoState135Action,
	{_State242, LessOrEqualToken}:                        _GotoState136Action,
	{_State242, GreaterToken}:                            _GotoState133Action,
	{_State242, GreaterOrEqualToken}:                     _GotoState134Action,
	{_State242, CmpOpType}:                               _GotoState138Action,
	{_State243, AddToken}:                                _GotoState126Action,
	{_State243, SubToken}:                                _GotoState129Action,
	{_State243, BitXorToken}:                             _GotoState128Action,
	{_State243, BitOrToken}:                              _GotoState127Action,
	{_State243, AddOpType}:                               _GotoState130Action,
	{_State244, RparenToken}:                             _GotoState350Action,
	{_State246, ForToken}:                                _GotoState351Action,
	{_State247, InToken}:                                 _GotoState353Action,
	{_State247, AssignToken}:                             _GotoState352Action,
	{_State248, SemicolonToken}:                          _GotoState354Action,
	{_State249, DoToken}:                                 _GotoState355Action,
	{_State250, IntegerLiteralToken}:                     _GotoState22Action,
	{_State250, FloatLiteralToken}:                       _GotoState18Action,
	{_State250, RuneLiteralToken}:                        _GotoState30Action,
	{_State250, StringLiteralToken}:                      _GotoState31Action,
	{_State250, IdentifierToken}:                         _GotoState21Action,
	{_State250, TrueToken}:                               _GotoState34Action,
	{_State250, FalseToken}:                              _GotoState17Action,
	{_State250, StructToken}:                             _GotoState32Action,
	{_State250, FuncToken}:                               _GotoState19Action,
	{_State250, VarToken}:                                _GotoState263Action,
	{_State250, LetToken}:                                _GotoState25Action,
	{_State250, NotToken}:                                _GotoState28Action,
	{_State250, LabelDeclToken}:                          _GotoState23Action,
	{_State250, LparenToken}:                             _GotoState26Action,
	{_State250, LbracketToken}:                           _GotoState24Action,
	{_State250, DotToken}:                                _GotoState262Action,
	{_State250, SubToken}:                                _GotoState33Action,
	{_State250, MulToken}:                                _GotoState27Action,
	{_State250, BitNegToken}:                             _GotoState16Action,
	{_State250, BitAndToken}:                             _GotoState15Action,
	{_State250, GreaterToken}:                            _GotoState20Action,
	{_State250, ParseErrorToken}:                         _GotoState29Action,
	{_State250, CasePatternsType}:                        _GotoState356Action,
	{_State250, VarDeclPatternType}:                      _GotoState70Action,
	{_State250, VarOrLetType}:                            _GotoState71Action,
	{_State250, AssignPatternType}:                       _GotoState264Action,
	{_State250, CasePatternType}:                         _GotoState265Action,
	{_State250, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State250, SequenceExprType}:                        _GotoState267Action,
	{_State250, CallExprType}:                            _GotoState49Action,
	{_State250, AtomExprType}:                            _GotoState42Action,
	{_State250, ParseErrorExprType}:                      _GotoState62Action,
	{_State250, LiteralExprType}:                         _GotoState57Action,
	{_State250, IdentifierExprType}:                      _GotoState52Action,
	{_State250, BlockExprType}:                           _GotoState48Action,
	{_State250, InitializeExprType}:                      _GotoState56Action,
	{_State250, ImplicitStructExprType}:                  _GotoState53Action,
	{_State250, AccessibleExprType}:                      _GotoState37Action,
	{_State250, AccessExprType}:                          _GotoState36Action,
	{_State250, IndexExprType}:                           _GotoState54Action,
	{_State250, PostfixableExprType}:                     _GotoState64Action,
	{_State250, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State250, PrefixableExprType}:                      _GotoState67Action,
	{_State250, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State250, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State250, MulExprType}:                             _GotoState59Action,
	{_State250, BinaryMulExprType}:                       _GotoState46Action,
	{_State250, AddExprType}:                             _GotoState38Action,
	{_State250, BinaryAddExprType}:                       _GotoState43Action,
	{_State250, CmpExprType}:                             _GotoState50Action,
	{_State250, BinaryCmpExprType}:                       _GotoState45Action,
	{_State250, AndExprType}:                             _GotoState39Action,
	{_State250, BinaryAndExprType}:                       _GotoState44Action,
	{_State250, OrExprType}:                              _GotoState61Action,
	{_State250, BinaryOrExprType}:                        _GotoState47Action,
	{_State250, InitializableTypeType}:                   _GotoState55Action,
	{_State250, SliceTypeType}:                           _GotoState69Action,
	{_State250, ArrayTypeType}:                           _GotoState41Action,
	{_State250, MapTypeType}:                             _GotoState58Action,
	{_State250, ExplicitStructDefType}:                   _GotoState51Action,
	{_State250, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State251, LbraceToken}:                             _GotoState73Action,
	{_State251, StatementBlockType}:                      _GotoState357Action,
	{_State253, LbraceToken}:                             _GotoState73Action,
	{_State253, StatementBlockType}:                      _GotoState358Action,
	{_State254, AndToken}:                                _GotoState131Action,
	{_State256, AssignToken}:                             _GotoState359Action,
	{_State258, RparenToken}:                             _GotoState361Action,
	{_State258, CommaToken}:                              _GotoState360Action,
	{_State261, AddToken}:                                _GotoState218Action,
	{_State261, SubToken}:                                _GotoState223Action,
	{_State261, MulToken}:                                _GotoState221Action,
	{_State261, TraitOpType}:                             _GotoState224Action,
	{_State262, IdentifierToken}:                         _GotoState362Action,
	{_State263, DotToken}:                                _GotoState363Action,
	{_State266, CommaToken}:                              _GotoState365Action,
	{_State266, ColonToken}:                              _GotoState364Action,
	{_State268, IntegerLiteralToken}:                     _GotoState22Action,
	{_State268, FloatLiteralToken}:                       _GotoState18Action,
	{_State268, RuneLiteralToken}:                        _GotoState30Action,
	{_State268, StringLiteralToken}:                      _GotoState31Action,
	{_State268, IdentifierToken}:                         _GotoState21Action,
	{_State268, TrueToken}:                               _GotoState34Action,
	{_State268, FalseToken}:                              _GotoState17Action,
	{_State268, ReturnToken}:                             _GotoState169Action,
	{_State268, BreakToken}:                              _GotoState162Action,
	{_State268, ContinueToken}:                           _GotoState164Action,
	{_State268, FallthroughToken}:                        _GotoState167Action,
	{_State268, UnsafeToken}:                             _GotoState170Action,
	{_State268, StructToken}:                             _GotoState32Action,
	{_State268, FuncToken}:                               _GotoState19Action,
	{_State268, AsyncToken}:                              _GotoState161Action,
	{_State268, DeferToken}:                              _GotoState166Action,
	{_State268, VarToken}:                                _GotoState35Action,
	{_State268, LetToken}:                                _GotoState25Action,
	{_State268, NotToken}:                                _GotoState28Action,
	{_State268, LabelDeclToken}:                          _GotoState23Action,
	{_State268, LparenToken}:                             _GotoState26Action,
	{_State268, LbracketToken}:                           _GotoState24Action,
	{_State268, SubToken}:                                _GotoState33Action,
	{_State268, MulToken}:                                _GotoState27Action,
	{_State268, BitNegToken}:                             _GotoState16Action,
	{_State268, BitAndToken}:                             _GotoState15Action,
	{_State268, GreaterToken}:                            _GotoState20Action,
	{_State268, ParseErrorToken}:                         _GotoState29Action,
	{_State268, SimpleStatementType}:                     _GotoState367Action,
	{_State268, OptionalSimpleStatementType}:             _GotoState366Action,
	{_State268, ExpressionOrImproperStructStatementType}: _GotoState178Action,
	{_State268, ExpressionsType}:                         _GotoState179Action,
	{_State268, CallbackOpType}:                          _GotoState175Action,
	{_State268, CallbackOpStatementType}:                 _GotoState176Action,
	{_State268, UnsafeStatementType}:                     _GotoState190Action,
	{_State268, JumpStatementType}:                       _GotoState182Action,
	{_State268, JumpTypeType}:                            _GotoState183Action,
	{_State268, FallthroughStatementType}:                _GotoState180Action,
	{_State268, AssignStatementType}:                     _GotoState173Action,
	{_State268, UnaryOpAssignStatementType}:              _GotoState189Action,
	{_State268, BinaryOpAssignStatementType}:             _GotoState174Action,
	{_State268, VarDeclPatternType}:                      _GotoState70Action,
	{_State268, VarOrLetType}:                            _GotoState71Action,
	{_State268, AssignPatternType}:                       _GotoState172Action,
	{_State268, ExpressionType}:                          _GotoState177Action,
	{_State268, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State268, SequenceExprType}:                        _GotoState185Action,
	{_State268, CallExprType}:                            _GotoState49Action,
	{_State268, AtomExprType}:                            _GotoState42Action,
	{_State268, ParseErrorExprType}:                      _GotoState62Action,
	{_State268, LiteralExprType}:                         _GotoState57Action,
	{_State268, IdentifierExprType}:                      _GotoState52Action,
	{_State268, BlockExprType}:                           _GotoState48Action,
	{_State268, InitializeExprType}:                      _GotoState56Action,
	{_State268, ImplicitStructExprType}:                  _GotoState53Action,
	{_State268, AccessibleExprType}:                      _GotoState171Action,
	{_State268, AccessExprType}:                          _GotoState36Action,
	{_State268, IndexExprType}:                           _GotoState54Action,
	{_State268, PostfixableExprType}:                     _GotoState64Action,
	{_State268, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State268, PrefixableExprType}:                      _GotoState67Action,
	{_State268, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State268, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State268, MulExprType}:                             _GotoState59Action,
	{_State268, BinaryMulExprType}:                       _GotoState46Action,
	{_State268, AddExprType}:                             _GotoState38Action,
	{_State268, BinaryAddExprType}:                       _GotoState43Action,
	{_State268, CmpExprType}:                             _GotoState50Action,
	{_State268, BinaryCmpExprType}:                       _GotoState45Action,
	{_State268, AndExprType}:                             _GotoState39Action,
	{_State268, BinaryAndExprType}:                       _GotoState44Action,
	{_State268, OrExprType}:                              _GotoState61Action,
	{_State268, BinaryOrExprType}:                        _GotoState47Action,
	{_State268, InitializableTypeType}:                   _GotoState55Action,
	{_State268, SliceTypeType}:                           _GotoState69Action,
	{_State268, ArrayTypeType}:                           _GotoState41Action,
	{_State268, MapTypeType}:                             _GotoState58Action,
	{_State268, ExplicitStructDefType}:                   _GotoState51Action,
	{_State268, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State269, StringLiteralToken}:                      _GotoState270Action,
	{_State269, ImportClauseType}:                        _GotoState368Action,
	{_State269, ProperImportClausesType}:                 _GotoState370Action,
	{_State269, ImportClausesType}:                       _GotoState369Action,
	{_State270, AsToken}:                                 _GotoState371Action,
	{_State272, IdentifierToken}:                         _GotoState372Action,
	{_State286, IntegerLiteralToken}:                     _GotoState22Action,
	{_State286, FloatLiteralToken}:                       _GotoState18Action,
	{_State286, RuneLiteralToken}:                        _GotoState30Action,
	{_State286, StringLiteralToken}:                      _GotoState31Action,
	{_State286, IdentifierToken}:                         _GotoState21Action,
	{_State286, TrueToken}:                               _GotoState34Action,
	{_State286, FalseToken}:                              _GotoState17Action,
	{_State286, StructToken}:                             _GotoState32Action,
	{_State286, FuncToken}:                               _GotoState19Action,
	{_State286, VarToken}:                                _GotoState35Action,
	{_State286, LetToken}:                                _GotoState25Action,
	{_State286, NotToken}:                                _GotoState28Action,
	{_State286, LabelDeclToken}:                          _GotoState23Action,
	{_State286, LparenToken}:                             _GotoState26Action,
	{_State286, LbracketToken}:                           _GotoState24Action,
	{_State286, SubToken}:                                _GotoState33Action,
	{_State286, MulToken}:                                _GotoState27Action,
	{_State286, BitNegToken}:                             _GotoState16Action,
	{_State286, BitAndToken}:                             _GotoState15Action,
	{_State286, GreaterToken}:                            _GotoState20Action,
	{_State286, ParseErrorToken}:                         _GotoState29Action,
	{_State286, VarDeclPatternType}:                      _GotoState70Action,
	{_State286, VarOrLetType}:                            _GotoState71Action,
	{_State286, ExpressionType}:                          _GotoState373Action,
	{_State286, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State286, SequenceExprType}:                        _GotoState68Action,
	{_State286, CallExprType}:                            _GotoState49Action,
	{_State286, AtomExprType}:                            _GotoState42Action,
	{_State286, ParseErrorExprType}:                      _GotoState62Action,
	{_State286, LiteralExprType}:                         _GotoState57Action,
	{_State286, IdentifierExprType}:                      _GotoState52Action,
	{_State286, BlockExprType}:                           _GotoState48Action,
	{_State286, InitializeExprType}:                      _GotoState56Action,
	{_State286, ImplicitStructExprType}:                  _GotoState53Action,
	{_State286, AccessibleExprType}:                      _GotoState37Action,
	{_State286, AccessExprType}:                          _GotoState36Action,
	{_State286, IndexExprType}:                           _GotoState54Action,
	{_State286, PostfixableExprType}:                     _GotoState64Action,
	{_State286, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State286, PrefixableExprType}:                      _GotoState67Action,
	{_State286, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State286, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State286, MulExprType}:                             _GotoState59Action,
	{_State286, BinaryMulExprType}:                       _GotoState46Action,
	{_State286, AddExprType}:                             _GotoState38Action,
	{_State286, BinaryAddExprType}:                       _GotoState43Action,
	{_State286, CmpExprType}:                             _GotoState50Action,
	{_State286, BinaryCmpExprType}:                       _GotoState45Action,
	{_State286, AndExprType}:                             _GotoState39Action,
	{_State286, BinaryAndExprType}:                       _GotoState44Action,
	{_State286, OrExprType}:                              _GotoState61Action,
	{_State286, BinaryOrExprType}:                        _GotoState47Action,
	{_State286, InitializableTypeType}:                   _GotoState55Action,
	{_State286, SliceTypeType}:                           _GotoState69Action,
	{_State286, ArrayTypeType}:                           _GotoState41Action,
	{_State286, MapTypeType}:                             _GotoState58Action,
	{_State286, ExplicitStructDefType}:                   _GotoState51Action,
	{_State286, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State288, IntegerLiteralToken}:                     _GotoState22Action,
	{_State288, FloatLiteralToken}:                       _GotoState18Action,
	{_State288, RuneLiteralToken}:                        _GotoState30Action,
	{_State288, StringLiteralToken}:                      _GotoState31Action,
	{_State288, IdentifierToken}:                         _GotoState21Action,
	{_State288, TrueToken}:                               _GotoState34Action,
	{_State288, FalseToken}:                              _GotoState17Action,
	{_State288, StructToken}:                             _GotoState32Action,
	{_State288, FuncToken}:                               _GotoState19Action,
	{_State288, VarToken}:                                _GotoState35Action,
	{_State288, LetToken}:                                _GotoState25Action,
	{_State288, NotToken}:                                _GotoState28Action,
	{_State288, LabelDeclToken}:                          _GotoState23Action,
	{_State288, LparenToken}:                             _GotoState26Action,
	{_State288, LbracketToken}:                           _GotoState24Action,
	{_State288, SubToken}:                                _GotoState33Action,
	{_State288, MulToken}:                                _GotoState27Action,
	{_State288, BitNegToken}:                             _GotoState16Action,
	{_State288, BitAndToken}:                             _GotoState15Action,
	{_State288, GreaterToken}:                            _GotoState20Action,
	{_State288, ParseErrorToken}:                         _GotoState29Action,
	{_State288, VarDeclPatternType}:                      _GotoState70Action,
	{_State288, VarOrLetType}:                            _GotoState71Action,
	{_State288, ExpressionType}:                          _GotoState374Action,
	{_State288, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State288, SequenceExprType}:                        _GotoState68Action,
	{_State288, CallExprType}:                            _GotoState49Action,
	{_State288, AtomExprType}:                            _GotoState42Action,
	{_State288, ParseErrorExprType}:                      _GotoState62Action,
	{_State288, LiteralExprType}:                         _GotoState57Action,
	{_State288, IdentifierExprType}:                      _GotoState52Action,
	{_State288, BlockExprType}:                           _GotoState48Action,
	{_State288, InitializeExprType}:                      _GotoState56Action,
	{_State288, ImplicitStructExprType}:                  _GotoState53Action,
	{_State288, AccessibleExprType}:                      _GotoState37Action,
	{_State288, AccessExprType}:                          _GotoState36Action,
	{_State288, IndexExprType}:                           _GotoState54Action,
	{_State288, PostfixableExprType}:                     _GotoState64Action,
	{_State288, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State288, PrefixableExprType}:                      _GotoState67Action,
	{_State288, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State288, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State288, MulExprType}:                             _GotoState59Action,
	{_State288, BinaryMulExprType}:                       _GotoState46Action,
	{_State288, AddExprType}:                             _GotoState38Action,
	{_State288, BinaryAddExprType}:                       _GotoState43Action,
	{_State288, CmpExprType}:                             _GotoState50Action,
	{_State288, BinaryCmpExprType}:                       _GotoState45Action,
	{_State288, AndExprType}:                             _GotoState39Action,
	{_State288, BinaryAndExprType}:                       _GotoState44Action,
	{_State288, OrExprType}:                              _GotoState61Action,
	{_State288, BinaryOrExprType}:                        _GotoState47Action,
	{_State288, InitializableTypeType}:                   _GotoState55Action,
	{_State288, SliceTypeType}:                           _GotoState69Action,
	{_State288, ArrayTypeType}:                           _GotoState41Action,
	{_State288, MapTypeType}:                             _GotoState58Action,
	{_State288, ExplicitStructDefType}:                   _GotoState51Action,
	{_State288, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State289, LbracketToken}:                           _GotoState123Action,
	{_State289, DotToken}:                                _GotoState122Action,
	{_State289, DollarLbracketToken}:                     _GotoState121Action,
	{_State289, OptionalGenericBindingType}:              _GotoState125Action,
	{_State291, IntegerLiteralToken}:                     _GotoState22Action,
	{_State291, FloatLiteralToken}:                       _GotoState18Action,
	{_State291, RuneLiteralToken}:                        _GotoState30Action,
	{_State291, StringLiteralToken}:                      _GotoState31Action,
	{_State291, IdentifierToken}:                         _GotoState21Action,
	{_State291, TrueToken}:                               _GotoState34Action,
	{_State291, FalseToken}:                              _GotoState17Action,
	{_State291, StructToken}:                             _GotoState32Action,
	{_State291, FuncToken}:                               _GotoState19Action,
	{_State291, VarToken}:                                _GotoState35Action,
	{_State291, LetToken}:                                _GotoState25Action,
	{_State291, NotToken}:                                _GotoState28Action,
	{_State291, LabelDeclToken}:                          _GotoState23Action,
	{_State291, LparenToken}:                             _GotoState26Action,
	{_State291, LbracketToken}:                           _GotoState24Action,
	{_State291, SubToken}:                                _GotoState33Action,
	{_State291, MulToken}:                                _GotoState27Action,
	{_State291, BitNegToken}:                             _GotoState16Action,
	{_State291, BitAndToken}:                             _GotoState15Action,
	{_State291, GreaterToken}:                            _GotoState20Action,
	{_State291, ParseErrorToken}:                         _GotoState29Action,
	{_State291, VarDeclPatternType}:                      _GotoState70Action,
	{_State291, VarOrLetType}:                            _GotoState71Action,
	{_State291, ExpressionType}:                          _GotoState375Action,
	{_State291, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State291, SequenceExprType}:                        _GotoState68Action,
	{_State291, CallExprType}:                            _GotoState49Action,
	{_State291, AtomExprType}:                            _GotoState42Action,
	{_State291, ParseErrorExprType}:                      _GotoState62Action,
	{_State291, LiteralExprType}:                         _GotoState57Action,
	{_State291, IdentifierExprType}:                      _GotoState52Action,
	{_State291, BlockExprType}:                           _GotoState48Action,
	{_State291, InitializeExprType}:                      _GotoState56Action,
	{_State291, ImplicitStructExprType}:                  _GotoState53Action,
	{_State291, AccessibleExprType}:                      _GotoState37Action,
	{_State291, AccessExprType}:                          _GotoState36Action,
	{_State291, IndexExprType}:                           _GotoState54Action,
	{_State291, PostfixableExprType}:                     _GotoState64Action,
	{_State291, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State291, PrefixableExprType}:                      _GotoState67Action,
	{_State291, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State291, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State291, MulExprType}:                             _GotoState59Action,
	{_State291, BinaryMulExprType}:                       _GotoState46Action,
	{_State291, AddExprType}:                             _GotoState38Action,
	{_State291, BinaryAddExprType}:                       _GotoState43Action,
	{_State291, CmpExprType}:                             _GotoState50Action,
	{_State291, BinaryCmpExprType}:                       _GotoState45Action,
	{_State291, AndExprType}:                             _GotoState39Action,
	{_State291, BinaryAndExprType}:                       _GotoState44Action,
	{_State291, OrExprType}:                              _GotoState61Action,
	{_State291, BinaryOrExprType}:                        _GotoState47Action,
	{_State291, InitializableTypeType}:                   _GotoState55Action,
	{_State291, SliceTypeType}:                           _GotoState69Action,
	{_State291, ArrayTypeType}:                           _GotoState41Action,
	{_State291, MapTypeType}:                             _GotoState58Action,
	{_State291, ExplicitStructDefType}:                   _GotoState51Action,
	{_State291, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State292, IntegerLiteralToken}:                     _GotoState22Action,
	{_State292, FloatLiteralToken}:                       _GotoState18Action,
	{_State292, RuneLiteralToken}:                        _GotoState30Action,
	{_State292, StringLiteralToken}:                      _GotoState31Action,
	{_State292, IdentifierToken}:                         _GotoState21Action,
	{_State292, TrueToken}:                               _GotoState34Action,
	{_State292, FalseToken}:                              _GotoState17Action,
	{_State292, StructToken}:                             _GotoState32Action,
	{_State292, FuncToken}:                               _GotoState19Action,
	{_State292, VarToken}:                                _GotoState35Action,
	{_State292, LetToken}:                                _GotoState25Action,
	{_State292, NotToken}:                                _GotoState28Action,
	{_State292, LabelDeclToken}:                          _GotoState23Action,
	{_State292, LparenToken}:                             _GotoState26Action,
	{_State292, LbracketToken}:                           _GotoState24Action,
	{_State292, SubToken}:                                _GotoState33Action,
	{_State292, MulToken}:                                _GotoState27Action,
	{_State292, BitNegToken}:                             _GotoState16Action,
	{_State292, BitAndToken}:                             _GotoState15Action,
	{_State292, GreaterToken}:                            _GotoState20Action,
	{_State292, ParseErrorToken}:                         _GotoState29Action,
	{_State292, ExpressionsType}:                         _GotoState376Action,
	{_State292, VarDeclPatternType}:                      _GotoState70Action,
	{_State292, VarOrLetType}:                            _GotoState71Action,
	{_State292, ExpressionType}:                          _GotoState177Action,
	{_State292, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State292, SequenceExprType}:                        _GotoState68Action,
	{_State292, CallExprType}:                            _GotoState49Action,
	{_State292, AtomExprType}:                            _GotoState42Action,
	{_State292, ParseErrorExprType}:                      _GotoState62Action,
	{_State292, LiteralExprType}:                         _GotoState57Action,
	{_State292, IdentifierExprType}:                      _GotoState52Action,
	{_State292, BlockExprType}:                           _GotoState48Action,
	{_State292, InitializeExprType}:                      _GotoState56Action,
	{_State292, ImplicitStructExprType}:                  _GotoState53Action,
	{_State292, AccessibleExprType}:                      _GotoState37Action,
	{_State292, AccessExprType}:                          _GotoState36Action,
	{_State292, IndexExprType}:                           _GotoState54Action,
	{_State292, PostfixableExprType}:                     _GotoState64Action,
	{_State292, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State292, PrefixableExprType}:                      _GotoState67Action,
	{_State292, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State292, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State292, MulExprType}:                             _GotoState59Action,
	{_State292, BinaryMulExprType}:                       _GotoState46Action,
	{_State292, AddExprType}:                             _GotoState38Action,
	{_State292, BinaryAddExprType}:                       _GotoState43Action,
	{_State292, CmpExprType}:                             _GotoState50Action,
	{_State292, BinaryCmpExprType}:                       _GotoState45Action,
	{_State292, AndExprType}:                             _GotoState39Action,
	{_State292, BinaryAndExprType}:                       _GotoState44Action,
	{_State292, OrExprType}:                              _GotoState61Action,
	{_State292, BinaryOrExprType}:                        _GotoState47Action,
	{_State292, InitializableTypeType}:                   _GotoState55Action,
	{_State292, SliceTypeType}:                           _GotoState69Action,
	{_State292, ArrayTypeType}:                           _GotoState41Action,
	{_State292, MapTypeType}:                             _GotoState58Action,
	{_State292, ExplicitStructDefType}:                   _GotoState51Action,
	{_State292, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State293, CommaToken}:                              _GotoState291Action,
	{_State294, IntegerLiteralToken}:                     _GotoState22Action,
	{_State294, FloatLiteralToken}:                       _GotoState18Action,
	{_State294, RuneLiteralToken}:                        _GotoState30Action,
	{_State294, StringLiteralToken}:                      _GotoState31Action,
	{_State294, IdentifierToken}:                         _GotoState21Action,
	{_State294, TrueToken}:                               _GotoState34Action,
	{_State294, FalseToken}:                              _GotoState17Action,
	{_State294, CaseToken}:                               _GotoState163Action,
	{_State294, DefaultToken}:                            _GotoState165Action,
	{_State294, ReturnToken}:                             _GotoState169Action,
	{_State294, BreakToken}:                              _GotoState162Action,
	{_State294, ContinueToken}:                           _GotoState164Action,
	{_State294, FallthroughToken}:                        _GotoState167Action,
	{_State294, ImportToken}:                             _GotoState168Action,
	{_State294, UnsafeToken}:                             _GotoState170Action,
	{_State294, StructToken}:                             _GotoState32Action,
	{_State294, FuncToken}:                               _GotoState19Action,
	{_State294, AsyncToken}:                              _GotoState161Action,
	{_State294, DeferToken}:                              _GotoState166Action,
	{_State294, VarToken}:                                _GotoState35Action,
	{_State294, LetToken}:                                _GotoState25Action,
	{_State294, NotToken}:                                _GotoState28Action,
	{_State294, LabelDeclToken}:                          _GotoState23Action,
	{_State294, LparenToken}:                             _GotoState26Action,
	{_State294, LbracketToken}:                           _GotoState24Action,
	{_State294, SubToken}:                                _GotoState33Action,
	{_State294, MulToken}:                                _GotoState27Action,
	{_State294, BitNegToken}:                             _GotoState16Action,
	{_State294, BitAndToken}:                             _GotoState15Action,
	{_State294, GreaterToken}:                            _GotoState20Action,
	{_State294, ParseErrorToken}:                         _GotoState29Action,
	{_State294, SimpleStatementType}:                     _GotoState186Action,
	{_State294, StatementType}:                           _GotoState377Action,
	{_State294, ExpressionOrImproperStructStatementType}: _GotoState178Action,
	{_State294, ExpressionsType}:                         _GotoState179Action,
	{_State294, CallbackOpType}:                          _GotoState175Action,
	{_State294, CallbackOpStatementType}:                 _GotoState176Action,
	{_State294, UnsafeStatementType}:                     _GotoState190Action,
	{_State294, JumpStatementType}:                       _GotoState182Action,
	{_State294, JumpTypeType}:                            _GotoState183Action,
	{_State294, FallthroughStatementType}:                _GotoState180Action,
	{_State294, AssignStatementType}:                     _GotoState173Action,
	{_State294, UnaryOpAssignStatementType}:              _GotoState189Action,
	{_State294, BinaryOpAssignStatementType}:             _GotoState174Action,
	{_State294, ImportStatementType}:                     _GotoState181Action,
	{_State294, VarDeclPatternType}:                      _GotoState70Action,
	{_State294, VarOrLetType}:                            _GotoState71Action,
	{_State294, AssignPatternType}:                       _GotoState172Action,
	{_State294, ExpressionType}:                          _GotoState177Action,
	{_State294, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State294, SequenceExprType}:                        _GotoState185Action,
	{_State294, CallExprType}:                            _GotoState49Action,
	{_State294, AtomExprType}:                            _GotoState42Action,
	{_State294, ParseErrorExprType}:                      _GotoState62Action,
	{_State294, LiteralExprType}:                         _GotoState57Action,
	{_State294, IdentifierExprType}:                      _GotoState52Action,
	{_State294, BlockExprType}:                           _GotoState48Action,
	{_State294, InitializeExprType}:                      _GotoState56Action,
	{_State294, ImplicitStructExprType}:                  _GotoState53Action,
	{_State294, AccessibleExprType}:                      _GotoState171Action,
	{_State294, AccessExprType}:                          _GotoState36Action,
	{_State294, IndexExprType}:                           _GotoState54Action,
	{_State294, PostfixableExprType}:                     _GotoState64Action,
	{_State294, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State294, PrefixableExprType}:                      _GotoState67Action,
	{_State294, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State294, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State294, MulExprType}:                             _GotoState59Action,
	{_State294, BinaryMulExprType}:                       _GotoState46Action,
	{_State294, AddExprType}:                             _GotoState38Action,
	{_State294, BinaryAddExprType}:                       _GotoState43Action,
	{_State294, CmpExprType}:                             _GotoState50Action,
	{_State294, BinaryCmpExprType}:                       _GotoState45Action,
	{_State294, AndExprType}:                             _GotoState39Action,
	{_State294, BinaryAndExprType}:                       _GotoState44Action,
	{_State294, OrExprType}:                              _GotoState61Action,
	{_State294, BinaryOrExprType}:                        _GotoState47Action,
	{_State294, InitializableTypeType}:                   _GotoState55Action,
	{_State294, SliceTypeType}:                           _GotoState69Action,
	{_State294, ArrayTypeType}:                           _GotoState41Action,
	{_State294, MapTypeType}:                             _GotoState58Action,
	{_State294, ExplicitStructDefType}:                   _GotoState51Action,
	{_State294, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State295, IntegerLiteralToken}:                     _GotoState22Action,
	{_State295, FloatLiteralToken}:                       _GotoState18Action,
	{_State295, RuneLiteralToken}:                        _GotoState30Action,
	{_State295, StringLiteralToken}:                      _GotoState31Action,
	{_State295, IdentifierToken}:                         _GotoState21Action,
	{_State295, TrueToken}:                               _GotoState34Action,
	{_State295, FalseToken}:                              _GotoState17Action,
	{_State295, CaseToken}:                               _GotoState163Action,
	{_State295, DefaultToken}:                            _GotoState165Action,
	{_State295, ReturnToken}:                             _GotoState169Action,
	{_State295, BreakToken}:                              _GotoState162Action,
	{_State295, ContinueToken}:                           _GotoState164Action,
	{_State295, FallthroughToken}:                        _GotoState167Action,
	{_State295, ImportToken}:                             _GotoState168Action,
	{_State295, UnsafeToken}:                             _GotoState170Action,
	{_State295, StructToken}:                             _GotoState32Action,
	{_State295, FuncToken}:                               _GotoState19Action,
	{_State295, AsyncToken}:                              _GotoState161Action,
	{_State295, DeferToken}:                              _GotoState166Action,
	{_State295, VarToken}:                                _GotoState35Action,
	{_State295, LetToken}:                                _GotoState25Action,
	{_State295, NotToken}:                                _GotoState28Action,
	{_State295, LabelDeclToken}:                          _GotoState23Action,
	{_State295, LparenToken}:                             _GotoState26Action,
	{_State295, LbracketToken}:                           _GotoState24Action,
	{_State295, SubToken}:                                _GotoState33Action,
	{_State295, MulToken}:                                _GotoState27Action,
	{_State295, BitNegToken}:                             _GotoState16Action,
	{_State295, BitAndToken}:                             _GotoState15Action,
	{_State295, GreaterToken}:                            _GotoState20Action,
	{_State295, ParseErrorToken}:                         _GotoState29Action,
	{_State295, SimpleStatementType}:                     _GotoState186Action,
	{_State295, StatementType}:                           _GotoState378Action,
	{_State295, ExpressionOrImproperStructStatementType}: _GotoState178Action,
	{_State295, ExpressionsType}:                         _GotoState179Action,
	{_State295, CallbackOpType}:                          _GotoState175Action,
	{_State295, CallbackOpStatementType}:                 _GotoState176Action,
	{_State295, UnsafeStatementType}:                     _GotoState190Action,
	{_State295, JumpStatementType}:                       _GotoState182Action,
	{_State295, JumpTypeType}:                            _GotoState183Action,
	{_State295, FallthroughStatementType}:                _GotoState180Action,
	{_State295, AssignStatementType}:                     _GotoState173Action,
	{_State295, UnaryOpAssignStatementType}:              _GotoState189Action,
	{_State295, BinaryOpAssignStatementType}:             _GotoState174Action,
	{_State295, ImportStatementType}:                     _GotoState181Action,
	{_State295, VarDeclPatternType}:                      _GotoState70Action,
	{_State295, VarOrLetType}:                            _GotoState71Action,
	{_State295, AssignPatternType}:                       _GotoState172Action,
	{_State295, ExpressionType}:                          _GotoState177Action,
	{_State295, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State295, SequenceExprType}:                        _GotoState185Action,
	{_State295, CallExprType}:                            _GotoState49Action,
	{_State295, AtomExprType}:                            _GotoState42Action,
	{_State295, ParseErrorExprType}:                      _GotoState62Action,
	{_State295, LiteralExprType}:                         _GotoState57Action,
	{_State295, IdentifierExprType}:                      _GotoState52Action,
	{_State295, BlockExprType}:                           _GotoState48Action,
	{_State295, InitializeExprType}:                      _GotoState56Action,
	{_State295, ImplicitStructExprType}:                  _GotoState53Action,
	{_State295, AccessibleExprType}:                      _GotoState171Action,
	{_State295, AccessExprType}:                          _GotoState36Action,
	{_State295, IndexExprType}:                           _GotoState54Action,
	{_State295, PostfixableExprType}:                     _GotoState64Action,
	{_State295, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State295, PrefixableExprType}:                      _GotoState67Action,
	{_State295, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State295, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State295, MulExprType}:                             _GotoState59Action,
	{_State295, BinaryMulExprType}:                       _GotoState46Action,
	{_State295, AddExprType}:                             _GotoState38Action,
	{_State295, BinaryAddExprType}:                       _GotoState43Action,
	{_State295, CmpExprType}:                             _GotoState50Action,
	{_State295, BinaryCmpExprType}:                       _GotoState45Action,
	{_State295, AndExprType}:                             _GotoState39Action,
	{_State295, BinaryAndExprType}:                       _GotoState44Action,
	{_State295, OrExprType}:                              _GotoState61Action,
	{_State295, BinaryOrExprType}:                        _GotoState47Action,
	{_State295, InitializableTypeType}:                   _GotoState55Action,
	{_State295, SliceTypeType}:                           _GotoState69Action,
	{_State295, ArrayTypeType}:                           _GotoState41Action,
	{_State295, MapTypeType}:                             _GotoState58Action,
	{_State295, ExplicitStructDefType}:                   _GotoState51Action,
	{_State295, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State298, AddToken}:                                _GotoState218Action,
	{_State298, SubToken}:                                _GotoState223Action,
	{_State298, MulToken}:                                _GotoState221Action,
	{_State298, TraitOpType}:                             _GotoState224Action,
	{_State299, IdentifierToken}:                         _GotoState93Action,
	{_State299, StructToken}:                             _GotoState32Action,
	{_State299, EnumToken}:                               _GotoState90Action,
	{_State299, TraitToken}:                              _GotoState98Action,
	{_State299, FuncToken}:                               _GotoState92Action,
	{_State299, LparenToken}:                             _GotoState94Action,
	{_State299, LbracketToken}:                           _GotoState24Action,
	{_State299, DotToken}:                                _GotoState89Action,
	{_State299, QuestionToken}:                           _GotoState96Action,
	{_State299, ExclaimToken}:                            _GotoState91Action,
	{_State299, TildeTildeToken}:                         _GotoState97Action,
	{_State299, BitNegToken}:                             _GotoState88Action,
	{_State299, BitAndToken}:                             _GotoState87Action,
	{_State299, ParseErrorToken}:                         _GotoState95Action,
	{_State299, InitializableTypeType}:                   _GotoState104Action,
	{_State299, SliceTypeType}:                           _GotoState69Action,
	{_State299, ArrayTypeType}:                           _GotoState41Action,
	{_State299, MapTypeType}:                             _GotoState58Action,
	{_State299, AtomTypeType}:                            _GotoState99Action,
	{_State299, ParseErrorTypeType}:                      _GotoState105Action,
	{_State299, ReturnableTypeType}:                      _GotoState108Action,
	{_State299, PrefixedTypeType}:                        _GotoState107Action,
	{_State299, PrefixTypeOpType}:                        _GotoState106Action,
	{_State299, ValueTypeType}:                           _GotoState379Action,
	{_State299, TraitOpTypeType}:                         _GotoState110Action,
	{_State299, ImplicitStructDefType}:                   _GotoState103Action,
	{_State299, ExplicitStructDefType}:                   _GotoState51Action,
	{_State299, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State299, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State299, TraitDefType}:                            _GotoState109Action,
	{_State299, FuncTypeType}:                            _GotoState101Action,
	{_State301, RbracketToken}:                           _GotoState380Action,
	{_State302, CommaToken}:                              _GotoState381Action,
	{_State303, ImplementsToken}:                         _GotoState382Action,
	{_State303, AddToken}:                                _GotoState218Action,
	{_State303, SubToken}:                                _GotoState223Action,
	{_State303, MulToken}:                                _GotoState221Action,
	{_State303, TraitOpType}:                             _GotoState224Action,
	{_State305, IdentifierToken}:                         _GotoState198Action,
	{_State305, ParameterDefType}:                        _GotoState200Action,
	{_State305, ProperParameterDefsType}:                 _GotoState202Action,
	{_State305, ParameterDefsType}:                       _GotoState383Action,
	{_State306, IdentifierToken}:                         _GotoState93Action,
	{_State306, StructToken}:                             _GotoState32Action,
	{_State306, EnumToken}:                               _GotoState90Action,
	{_State306, TraitToken}:                              _GotoState98Action,
	{_State306, FuncToken}:                               _GotoState92Action,
	{_State306, LparenToken}:                             _GotoState94Action,
	{_State306, LbracketToken}:                           _GotoState24Action,
	{_State306, DotToken}:                                _GotoState89Action,
	{_State306, QuestionToken}:                           _GotoState96Action,
	{_State306, ExclaimToken}:                            _GotoState91Action,
	{_State306, TildeTildeToken}:                         _GotoState97Action,
	{_State306, BitNegToken}:                             _GotoState88Action,
	{_State306, BitAndToken}:                             _GotoState87Action,
	{_State306, ParseErrorToken}:                         _GotoState95Action,
	{_State306, InitializableTypeType}:                   _GotoState104Action,
	{_State306, SliceTypeType}:                           _GotoState69Action,
	{_State306, ArrayTypeType}:                           _GotoState41Action,
	{_State306, MapTypeType}:                             _GotoState58Action,
	{_State306, AtomTypeType}:                            _GotoState99Action,
	{_State306, ParseErrorTypeType}:                      _GotoState105Action,
	{_State306, ReturnableTypeType}:                      _GotoState108Action,
	{_State306, PrefixedTypeType}:                        _GotoState107Action,
	{_State306, PrefixTypeOpType}:                        _GotoState106Action,
	{_State306, ValueTypeType}:                           _GotoState384Action,
	{_State306, TraitOpTypeType}:                         _GotoState110Action,
	{_State306, ImplicitStructDefType}:                   _GotoState103Action,
	{_State306, ExplicitStructDefType}:                   _GotoState51Action,
	{_State306, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State306, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State306, TraitDefType}:                            _GotoState109Action,
	{_State306, FuncTypeType}:                            _GotoState101Action,
	{_State307, AddToken}:                                _GotoState218Action,
	{_State307, SubToken}:                                _GotoState223Action,
	{_State307, MulToken}:                                _GotoState221Action,
	{_State307, TraitOpType}:                             _GotoState224Action,
	{_State308, IdentifierToken}:                         _GotoState385Action,
	{_State309, IdentifierToken}:                         _GotoState93Action,
	{_State309, StructToken}:                             _GotoState32Action,
	{_State309, EnumToken}:                               _GotoState90Action,
	{_State309, TraitToken}:                              _GotoState98Action,
	{_State309, FuncToken}:                               _GotoState92Action,
	{_State309, LparenToken}:                             _GotoState94Action,
	{_State309, LbracketToken}:                           _GotoState24Action,
	{_State309, DotToken}:                                _GotoState89Action,
	{_State309, QuestionToken}:                           _GotoState96Action,
	{_State309, ExclaimToken}:                            _GotoState91Action,
	{_State309, TildeTildeToken}:                         _GotoState97Action,
	{_State309, BitNegToken}:                             _GotoState88Action,
	{_State309, BitAndToken}:                             _GotoState87Action,
	{_State309, ParseErrorToken}:                         _GotoState95Action,
	{_State309, InitializableTypeType}:                   _GotoState104Action,
	{_State309, SliceTypeType}:                           _GotoState69Action,
	{_State309, ArrayTypeType}:                           _GotoState41Action,
	{_State309, MapTypeType}:                             _GotoState58Action,
	{_State309, AtomTypeType}:                            _GotoState99Action,
	{_State309, ParseErrorTypeType}:                      _GotoState105Action,
	{_State309, ReturnableTypeType}:                      _GotoState387Action,
	{_State309, PrefixedTypeType}:                        _GotoState107Action,
	{_State309, PrefixTypeOpType}:                        _GotoState106Action,
	{_State309, ImplicitStructDefType}:                   _GotoState103Action,
	{_State309, ExplicitStructDefType}:                   _GotoState51Action,
	{_State309, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State309, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State309, TraitDefType}:                            _GotoState109Action,
	{_State309, ReturnTypeType}:                          _GotoState386Action,
	{_State309, FuncTypeType}:                            _GotoState101Action,
	{_State310, IdentifierToken}:                         _GotoState198Action,
	{_State310, ParameterDefType}:                        _GotoState388Action,
	{_State311, NewlinesToken}:                           _GotoState389Action,
	{_State311, OrToken}:                                 _GotoState390Action,
	{_State312, RparenToken}:                             _GotoState391Action,
	{_State313, AssignToken}:                             _GotoState325Action,
	{_State314, NewlinesToken}:                           _GotoState392Action,
	{_State314, OrToken}:                                 _GotoState393Action,
	{_State315, IdentifierToken}:                         _GotoState93Action,
	{_State315, StructToken}:                             _GotoState32Action,
	{_State315, EnumToken}:                               _GotoState90Action,
	{_State315, TraitToken}:                              _GotoState98Action,
	{_State315, FuncToken}:                               _GotoState92Action,
	{_State315, LparenToken}:                             _GotoState94Action,
	{_State315, LbracketToken}:                           _GotoState24Action,
	{_State315, DotToken}:                                _GotoState89Action,
	{_State315, QuestionToken}:                           _GotoState96Action,
	{_State315, ExclaimToken}:                            _GotoState91Action,
	{_State315, TildeTildeToken}:                         _GotoState97Action,
	{_State315, BitNegToken}:                             _GotoState88Action,
	{_State315, BitAndToken}:                             _GotoState87Action,
	{_State315, ParseErrorToken}:                         _GotoState95Action,
	{_State315, InitializableTypeType}:                   _GotoState104Action,
	{_State315, SliceTypeType}:                           _GotoState69Action,
	{_State315, ArrayTypeType}:                           _GotoState41Action,
	{_State315, MapTypeType}:                             _GotoState58Action,
	{_State315, AtomTypeType}:                            _GotoState99Action,
	{_State315, ParseErrorTypeType}:                      _GotoState105Action,
	{_State315, ReturnableTypeType}:                      _GotoState108Action,
	{_State315, PrefixedTypeType}:                        _GotoState107Action,
	{_State315, PrefixTypeOpType}:                        _GotoState106Action,
	{_State315, ValueTypeType}:                           _GotoState394Action,
	{_State315, TraitOpTypeType}:                         _GotoState110Action,
	{_State315, ImplicitStructDefType}:                   _GotoState103Action,
	{_State315, ExplicitStructDefType}:                   _GotoState51Action,
	{_State315, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State315, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State315, TraitDefType}:                            _GotoState109Action,
	{_State315, FuncTypeType}:                            _GotoState101Action,
	{_State316, IdentifierToken}:                         _GotoState93Action,
	{_State316, StructToken}:                             _GotoState32Action,
	{_State316, EnumToken}:                               _GotoState90Action,
	{_State316, TraitToken}:                              _GotoState98Action,
	{_State316, FuncToken}:                               _GotoState92Action,
	{_State316, LparenToken}:                             _GotoState94Action,
	{_State316, LbracketToken}:                           _GotoState24Action,
	{_State316, DotToken}:                                _GotoState322Action,
	{_State316, QuestionToken}:                           _GotoState96Action,
	{_State316, ExclaimToken}:                            _GotoState91Action,
	{_State316, DollarLbracketToken}:                     _GotoState121Action,
	{_State316, EllipsisToken}:                           _GotoState395Action,
	{_State316, TildeTildeToken}:                         _GotoState97Action,
	{_State316, BitNegToken}:                             _GotoState88Action,
	{_State316, BitAndToken}:                             _GotoState87Action,
	{_State316, ParseErrorToken}:                         _GotoState95Action,
	{_State316, OptionalGenericBindingType}:              _GotoState207Action,
	{_State316, InitializableTypeType}:                   _GotoState104Action,
	{_State316, SliceTypeType}:                           _GotoState69Action,
	{_State316, ArrayTypeType}:                           _GotoState41Action,
	{_State316, MapTypeType}:                             _GotoState58Action,
	{_State316, AtomTypeType}:                            _GotoState99Action,
	{_State316, ParseErrorTypeType}:                      _GotoState105Action,
	{_State316, ReturnableTypeType}:                      _GotoState108Action,
	{_State316, PrefixedTypeType}:                        _GotoState107Action,
	{_State316, PrefixTypeOpType}:                        _GotoState106Action,
	{_State316, ValueTypeType}:                           _GotoState396Action,
	{_State316, TraitOpTypeType}:                         _GotoState110Action,
	{_State316, ImplicitStructDefType}:                   _GotoState103Action,
	{_State316, ExplicitStructDefType}:                   _GotoState51Action,
	{_State316, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State316, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State316, TraitDefType}:                            _GotoState109Action,
	{_State316, FuncTypeType}:                            _GotoState101Action,
	{_State318, RparenToken}:                             _GotoState397Action,
	{_State319, CommaToken}:                              _GotoState398Action,
	{_State320, AddToken}:                                _GotoState218Action,
	{_State320, SubToken}:                                _GotoState223Action,
	{_State320, MulToken}:                                _GotoState221Action,
	{_State320, TraitOpType}:                             _GotoState224Action,
	{_State321, DollarLbracketToken}:                     _GotoState121Action,
	{_State321, OptionalGenericBindingType}:              _GotoState399Action,
	{_State322, IdentifierToken}:                         _GotoState321Action,
	{_State322, DollarLbracketToken}:                     _GotoState121Action,
	{_State322, OptionalGenericBindingType}:              _GotoState203Action,
	{_State323, AddToken}:                                _GotoState218Action,
	{_State323, SubToken}:                                _GotoState223Action,
	{_State323, MulToken}:                                _GotoState221Action,
	{_State323, TraitOpType}:                             _GotoState224Action,
	{_State324, IdentifierToken}:                         _GotoState208Action,
	{_State324, UnsafeToken}:                             _GotoState170Action,
	{_State324, StructToken}:                             _GotoState32Action,
	{_State324, EnumToken}:                               _GotoState90Action,
	{_State324, TraitToken}:                              _GotoState98Action,
	{_State324, FuncToken}:                               _GotoState92Action,
	{_State324, LparenToken}:                             _GotoState94Action,
	{_State324, LbracketToken}:                           _GotoState24Action,
	{_State324, DotToken}:                                _GotoState89Action,
	{_State324, QuestionToken}:                           _GotoState96Action,
	{_State324, ExclaimToken}:                            _GotoState91Action,
	{_State324, TildeTildeToken}:                         _GotoState97Action,
	{_State324, BitNegToken}:                             _GotoState88Action,
	{_State324, BitAndToken}:                             _GotoState87Action,
	{_State324, ParseErrorToken}:                         _GotoState95Action,
	{_State324, UnsafeStatementType}:                     _GotoState214Action,
	{_State324, InitializableTypeType}:                   _GotoState104Action,
	{_State324, SliceTypeType}:                           _GotoState69Action,
	{_State324, ArrayTypeType}:                           _GotoState41Action,
	{_State324, MapTypeType}:                             _GotoState58Action,
	{_State324, AtomTypeType}:                            _GotoState99Action,
	{_State324, ParseErrorTypeType}:                      _GotoState105Action,
	{_State324, ReturnableTypeType}:                      _GotoState108Action,
	{_State324, PrefixedTypeType}:                        _GotoState107Action,
	{_State324, PrefixTypeOpType}:                        _GotoState106Action,
	{_State324, ValueTypeType}:                           _GotoState215Action,
	{_State324, TraitOpTypeType}:                         _GotoState110Action,
	{_State324, FieldDefType}:                            _GotoState313Action,
	{_State324, ImplicitStructDefType}:                   _GotoState103Action,
	{_State324, ExplicitStructDefType}:                   _GotoState51Action,
	{_State324, EnumValueDefType}:                        _GotoState400Action,
	{_State324, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State324, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State324, TraitDefType}:                            _GotoState109Action,
	{_State324, FuncTypeType}:                            _GotoState101Action,
	{_State325, DefaultToken}:                            _GotoState401Action,
	{_State326, IdentifierToken}:                         _GotoState208Action,
	{_State326, UnsafeToken}:                             _GotoState170Action,
	{_State326, StructToken}:                             _GotoState32Action,
	{_State326, EnumToken}:                               _GotoState90Action,
	{_State326, TraitToken}:                              _GotoState98Action,
	{_State326, FuncToken}:                               _GotoState92Action,
	{_State326, LparenToken}:                             _GotoState94Action,
	{_State326, LbracketToken}:                           _GotoState24Action,
	{_State326, DotToken}:                                _GotoState89Action,
	{_State326, QuestionToken}:                           _GotoState96Action,
	{_State326, ExclaimToken}:                            _GotoState91Action,
	{_State326, TildeTildeToken}:                         _GotoState97Action,
	{_State326, BitNegToken}:                             _GotoState88Action,
	{_State326, BitAndToken}:                             _GotoState87Action,
	{_State326, ParseErrorToken}:                         _GotoState95Action,
	{_State326, UnsafeStatementType}:                     _GotoState214Action,
	{_State326, InitializableTypeType}:                   _GotoState104Action,
	{_State326, SliceTypeType}:                           _GotoState69Action,
	{_State326, ArrayTypeType}:                           _GotoState41Action,
	{_State326, MapTypeType}:                             _GotoState58Action,
	{_State326, AtomTypeType}:                            _GotoState99Action,
	{_State326, ParseErrorTypeType}:                      _GotoState105Action,
	{_State326, ReturnableTypeType}:                      _GotoState108Action,
	{_State326, PrefixedTypeType}:                        _GotoState107Action,
	{_State326, PrefixTypeOpType}:                        _GotoState106Action,
	{_State326, ValueTypeType}:                           _GotoState215Action,
	{_State326, TraitOpTypeType}:                         _GotoState110Action,
	{_State326, FieldDefType}:                            _GotoState313Action,
	{_State326, ImplicitStructDefType}:                   _GotoState103Action,
	{_State326, ExplicitStructDefType}:                   _GotoState51Action,
	{_State326, EnumValueDefType}:                        _GotoState402Action,
	{_State326, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State326, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State326, TraitDefType}:                            _GotoState109Action,
	{_State326, FuncTypeType}:                            _GotoState101Action,
	{_State329, IdentifierToken}:                         _GotoState208Action,
	{_State329, UnsafeToken}:                             _GotoState170Action,
	{_State329, StructToken}:                             _GotoState32Action,
	{_State329, EnumToken}:                               _GotoState90Action,
	{_State329, TraitToken}:                              _GotoState98Action,
	{_State329, FuncToken}:                               _GotoState92Action,
	{_State329, LparenToken}:                             _GotoState94Action,
	{_State329, LbracketToken}:                           _GotoState24Action,
	{_State329, DotToken}:                                _GotoState89Action,
	{_State329, QuestionToken}:                           _GotoState96Action,
	{_State329, ExclaimToken}:                            _GotoState91Action,
	{_State329, TildeTildeToken}:                         _GotoState97Action,
	{_State329, BitNegToken}:                             _GotoState88Action,
	{_State329, BitAndToken}:                             _GotoState87Action,
	{_State329, ParseErrorToken}:                         _GotoState95Action,
	{_State329, UnsafeStatementType}:                     _GotoState214Action,
	{_State329, InitializableTypeType}:                   _GotoState104Action,
	{_State329, SliceTypeType}:                           _GotoState69Action,
	{_State329, ArrayTypeType}:                           _GotoState41Action,
	{_State329, MapTypeType}:                             _GotoState58Action,
	{_State329, AtomTypeType}:                            _GotoState99Action,
	{_State329, ParseErrorTypeType}:                      _GotoState105Action,
	{_State329, ReturnableTypeType}:                      _GotoState108Action,
	{_State329, PrefixedTypeType}:                        _GotoState107Action,
	{_State329, PrefixTypeOpType}:                        _GotoState106Action,
	{_State329, ValueTypeType}:                           _GotoState215Action,
	{_State329, TraitOpTypeType}:                         _GotoState110Action,
	{_State329, FieldDefType}:                            _GotoState403Action,
	{_State329, ImplicitStructDefType}:                   _GotoState103Action,
	{_State329, ExplicitStructDefType}:                   _GotoState51Action,
	{_State329, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State329, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State329, TraitDefType}:                            _GotoState109Action,
	{_State329, FuncTypeType}:                            _GotoState101Action,
	{_State330, IdentifierToken}:                         _GotoState404Action,
	{_State330, LparenToken}:                             _GotoState205Action,
	{_State333, NewlinesToken}:                           _GotoState406Action,
	{_State333, CommaToken}:                              _GotoState405Action,
	{_State334, RparenToken}:                             _GotoState407Action,
	{_State336, RbracketToken}:                           _GotoState408Action,
	{_State336, AddToken}:                                _GotoState218Action,
	{_State336, SubToken}:                                _GotoState223Action,
	{_State336, MulToken}:                                _GotoState221Action,
	{_State336, TraitOpType}:                             _GotoState224Action,
	{_State337, RbracketToken}:                           _GotoState409Action,
	{_State344, IdentifierToken}:                         _GotoState208Action,
	{_State344, UnsafeToken}:                             _GotoState170Action,
	{_State344, StructToken}:                             _GotoState32Action,
	{_State344, EnumToken}:                               _GotoState90Action,
	{_State344, TraitToken}:                              _GotoState98Action,
	{_State344, FuncToken}:                               _GotoState92Action,
	{_State344, LparenToken}:                             _GotoState94Action,
	{_State344, LbracketToken}:                           _GotoState24Action,
	{_State344, DotToken}:                                _GotoState89Action,
	{_State344, QuestionToken}:                           _GotoState96Action,
	{_State344, ExclaimToken}:                            _GotoState91Action,
	{_State344, TildeTildeToken}:                         _GotoState97Action,
	{_State344, BitNegToken}:                             _GotoState88Action,
	{_State344, BitAndToken}:                             _GotoState87Action,
	{_State344, ParseErrorToken}:                         _GotoState95Action,
	{_State344, UnsafeStatementType}:                     _GotoState214Action,
	{_State344, InitializableTypeType}:                   _GotoState104Action,
	{_State344, SliceTypeType}:                           _GotoState69Action,
	{_State344, ArrayTypeType}:                           _GotoState41Action,
	{_State344, MapTypeType}:                             _GotoState58Action,
	{_State344, AtomTypeType}:                            _GotoState99Action,
	{_State344, ParseErrorTypeType}:                      _GotoState105Action,
	{_State344, ReturnableTypeType}:                      _GotoState108Action,
	{_State344, PrefixedTypeType}:                        _GotoState107Action,
	{_State344, PrefixTypeOpType}:                        _GotoState106Action,
	{_State344, ValueTypeType}:                           _GotoState215Action,
	{_State344, TraitOpTypeType}:                         _GotoState110Action,
	{_State344, FieldDefType}:                            _GotoState410Action,
	{_State344, ImplicitStructDefType}:                   _GotoState103Action,
	{_State344, ExplicitStructDefType}:                   _GotoState51Action,
	{_State344, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State344, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State344, TraitDefType}:                            _GotoState109Action,
	{_State344, FuncTypeType}:                            _GotoState101Action,
	{_State345, IdentifierToken}:                         _GotoState208Action,
	{_State345, UnsafeToken}:                             _GotoState170Action,
	{_State345, StructToken}:                             _GotoState32Action,
	{_State345, EnumToken}:                               _GotoState90Action,
	{_State345, TraitToken}:                              _GotoState98Action,
	{_State345, FuncToken}:                               _GotoState92Action,
	{_State345, LparenToken}:                             _GotoState94Action,
	{_State345, LbracketToken}:                           _GotoState24Action,
	{_State345, DotToken}:                                _GotoState89Action,
	{_State345, QuestionToken}:                           _GotoState96Action,
	{_State345, ExclaimToken}:                            _GotoState91Action,
	{_State345, TildeTildeToken}:                         _GotoState97Action,
	{_State345, BitNegToken}:                             _GotoState88Action,
	{_State345, BitAndToken}:                             _GotoState87Action,
	{_State345, ParseErrorToken}:                         _GotoState95Action,
	{_State345, UnsafeStatementType}:                     _GotoState214Action,
	{_State345, InitializableTypeType}:                   _GotoState104Action,
	{_State345, SliceTypeType}:                           _GotoState69Action,
	{_State345, ArrayTypeType}:                           _GotoState41Action,
	{_State345, MapTypeType}:                             _GotoState58Action,
	{_State345, AtomTypeType}:                            _GotoState99Action,
	{_State345, ParseErrorTypeType}:                      _GotoState105Action,
	{_State345, ReturnableTypeType}:                      _GotoState108Action,
	{_State345, PrefixedTypeType}:                        _GotoState107Action,
	{_State345, PrefixTypeOpType}:                        _GotoState106Action,
	{_State345, ValueTypeType}:                           _GotoState215Action,
	{_State345, TraitOpTypeType}:                         _GotoState110Action,
	{_State345, FieldDefType}:                            _GotoState411Action,
	{_State345, ImplicitStructDefType}:                   _GotoState103Action,
	{_State345, ExplicitStructDefType}:                   _GotoState51Action,
	{_State345, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State345, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State345, TraitDefType}:                            _GotoState109Action,
	{_State345, FuncTypeType}:                            _GotoState101Action,
	{_State347, IdentifierToken}:                         _GotoState93Action,
	{_State347, StructToken}:                             _GotoState32Action,
	{_State347, EnumToken}:                               _GotoState90Action,
	{_State347, TraitToken}:                              _GotoState98Action,
	{_State347, FuncToken}:                               _GotoState92Action,
	{_State347, LparenToken}:                             _GotoState94Action,
	{_State347, LbracketToken}:                           _GotoState24Action,
	{_State347, DotToken}:                                _GotoState89Action,
	{_State347, QuestionToken}:                           _GotoState96Action,
	{_State347, ExclaimToken}:                            _GotoState91Action,
	{_State347, TildeTildeToken}:                         _GotoState97Action,
	{_State347, BitNegToken}:                             _GotoState88Action,
	{_State347, BitAndToken}:                             _GotoState87Action,
	{_State347, ParseErrorToken}:                         _GotoState95Action,
	{_State347, InitializableTypeType}:                   _GotoState104Action,
	{_State347, SliceTypeType}:                           _GotoState69Action,
	{_State347, ArrayTypeType}:                           _GotoState41Action,
	{_State347, MapTypeType}:                             _GotoState58Action,
	{_State347, AtomTypeType}:                            _GotoState99Action,
	{_State347, ParseErrorTypeType}:                      _GotoState105Action,
	{_State347, ReturnableTypeType}:                      _GotoState108Action,
	{_State347, PrefixedTypeType}:                        _GotoState107Action,
	{_State347, PrefixTypeOpType}:                        _GotoState106Action,
	{_State347, ValueTypeType}:                           _GotoState412Action,
	{_State347, TraitOpTypeType}:                         _GotoState110Action,
	{_State347, ImplicitStructDefType}:                   _GotoState103Action,
	{_State347, ExplicitStructDefType}:                   _GotoState51Action,
	{_State347, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State347, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State347, TraitDefType}:                            _GotoState109Action,
	{_State347, FuncTypeType}:                            _GotoState101Action,
	{_State349, RparenToken}:                             _GotoState413Action,
	{_State351, IntegerLiteralToken}:                     _GotoState22Action,
	{_State351, FloatLiteralToken}:                       _GotoState18Action,
	{_State351, RuneLiteralToken}:                        _GotoState30Action,
	{_State351, StringLiteralToken}:                      _GotoState31Action,
	{_State351, IdentifierToken}:                         _GotoState21Action,
	{_State351, TrueToken}:                               _GotoState34Action,
	{_State351, FalseToken}:                              _GotoState17Action,
	{_State351, StructToken}:                             _GotoState32Action,
	{_State351, FuncToken}:                               _GotoState19Action,
	{_State351, VarToken}:                                _GotoState35Action,
	{_State351, LetToken}:                                _GotoState25Action,
	{_State351, NotToken}:                                _GotoState28Action,
	{_State351, LabelDeclToken}:                          _GotoState23Action,
	{_State351, LparenToken}:                             _GotoState26Action,
	{_State351, LbracketToken}:                           _GotoState24Action,
	{_State351, SubToken}:                                _GotoState33Action,
	{_State351, MulToken}:                                _GotoState27Action,
	{_State351, BitNegToken}:                             _GotoState16Action,
	{_State351, BitAndToken}:                             _GotoState15Action,
	{_State351, GreaterToken}:                            _GotoState20Action,
	{_State351, ParseErrorToken}:                         _GotoState29Action,
	{_State351, VarDeclPatternType}:                      _GotoState70Action,
	{_State351, VarOrLetType}:                            _GotoState71Action,
	{_State351, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State351, SequenceExprType}:                        _GotoState414Action,
	{_State351, CallExprType}:                            _GotoState49Action,
	{_State351, AtomExprType}:                            _GotoState42Action,
	{_State351, ParseErrorExprType}:                      _GotoState62Action,
	{_State351, LiteralExprType}:                         _GotoState57Action,
	{_State351, IdentifierExprType}:                      _GotoState52Action,
	{_State351, BlockExprType}:                           _GotoState48Action,
	{_State351, InitializeExprType}:                      _GotoState56Action,
	{_State351, ImplicitStructExprType}:                  _GotoState53Action,
	{_State351, AccessibleExprType}:                      _GotoState37Action,
	{_State351, AccessExprType}:                          _GotoState36Action,
	{_State351, IndexExprType}:                           _GotoState54Action,
	{_State351, PostfixableExprType}:                     _GotoState64Action,
	{_State351, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State351, PrefixableExprType}:                      _GotoState67Action,
	{_State351, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State351, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State351, MulExprType}:                             _GotoState59Action,
	{_State351, BinaryMulExprType}:                       _GotoState46Action,
	{_State351, AddExprType}:                             _GotoState38Action,
	{_State351, BinaryAddExprType}:                       _GotoState43Action,
	{_State351, CmpExprType}:                             _GotoState50Action,
	{_State351, BinaryCmpExprType}:                       _GotoState45Action,
	{_State351, AndExprType}:                             _GotoState39Action,
	{_State351, BinaryAndExprType}:                       _GotoState44Action,
	{_State351, OrExprType}:                              _GotoState61Action,
	{_State351, BinaryOrExprType}:                        _GotoState47Action,
	{_State351, InitializableTypeType}:                   _GotoState55Action,
	{_State351, SliceTypeType}:                           _GotoState69Action,
	{_State351, ArrayTypeType}:                           _GotoState41Action,
	{_State351, MapTypeType}:                             _GotoState58Action,
	{_State351, ExplicitStructDefType}:                   _GotoState51Action,
	{_State351, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State352, IntegerLiteralToken}:                     _GotoState22Action,
	{_State352, FloatLiteralToken}:                       _GotoState18Action,
	{_State352, RuneLiteralToken}:                        _GotoState30Action,
	{_State352, StringLiteralToken}:                      _GotoState31Action,
	{_State352, IdentifierToken}:                         _GotoState21Action,
	{_State352, TrueToken}:                               _GotoState34Action,
	{_State352, FalseToken}:                              _GotoState17Action,
	{_State352, StructToken}:                             _GotoState32Action,
	{_State352, FuncToken}:                               _GotoState19Action,
	{_State352, VarToken}:                                _GotoState35Action,
	{_State352, LetToken}:                                _GotoState25Action,
	{_State352, NotToken}:                                _GotoState28Action,
	{_State352, LabelDeclToken}:                          _GotoState23Action,
	{_State352, LparenToken}:                             _GotoState26Action,
	{_State352, LbracketToken}:                           _GotoState24Action,
	{_State352, SubToken}:                                _GotoState33Action,
	{_State352, MulToken}:                                _GotoState27Action,
	{_State352, BitNegToken}:                             _GotoState16Action,
	{_State352, BitAndToken}:                             _GotoState15Action,
	{_State352, GreaterToken}:                            _GotoState20Action,
	{_State352, ParseErrorToken}:                         _GotoState29Action,
	{_State352, VarDeclPatternType}:                      _GotoState70Action,
	{_State352, VarOrLetType}:                            _GotoState71Action,
	{_State352, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State352, SequenceExprType}:                        _GotoState415Action,
	{_State352, CallExprType}:                            _GotoState49Action,
	{_State352, AtomExprType}:                            _GotoState42Action,
	{_State352, ParseErrorExprType}:                      _GotoState62Action,
	{_State352, LiteralExprType}:                         _GotoState57Action,
	{_State352, IdentifierExprType}:                      _GotoState52Action,
	{_State352, BlockExprType}:                           _GotoState48Action,
	{_State352, InitializeExprType}:                      _GotoState56Action,
	{_State352, ImplicitStructExprType}:                  _GotoState53Action,
	{_State352, AccessibleExprType}:                      _GotoState37Action,
	{_State352, AccessExprType}:                          _GotoState36Action,
	{_State352, IndexExprType}:                           _GotoState54Action,
	{_State352, PostfixableExprType}:                     _GotoState64Action,
	{_State352, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State352, PrefixableExprType}:                      _GotoState67Action,
	{_State352, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State352, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State352, MulExprType}:                             _GotoState59Action,
	{_State352, BinaryMulExprType}:                       _GotoState46Action,
	{_State352, AddExprType}:                             _GotoState38Action,
	{_State352, BinaryAddExprType}:                       _GotoState43Action,
	{_State352, CmpExprType}:                             _GotoState50Action,
	{_State352, BinaryCmpExprType}:                       _GotoState45Action,
	{_State352, AndExprType}:                             _GotoState39Action,
	{_State352, BinaryAndExprType}:                       _GotoState44Action,
	{_State352, OrExprType}:                              _GotoState61Action,
	{_State352, BinaryOrExprType}:                        _GotoState47Action,
	{_State352, InitializableTypeType}:                   _GotoState55Action,
	{_State352, SliceTypeType}:                           _GotoState69Action,
	{_State352, ArrayTypeType}:                           _GotoState41Action,
	{_State352, MapTypeType}:                             _GotoState58Action,
	{_State352, ExplicitStructDefType}:                   _GotoState51Action,
	{_State352, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State353, IntegerLiteralToken}:                     _GotoState22Action,
	{_State353, FloatLiteralToken}:                       _GotoState18Action,
	{_State353, RuneLiteralToken}:                        _GotoState30Action,
	{_State353, StringLiteralToken}:                      _GotoState31Action,
	{_State353, IdentifierToken}:                         _GotoState21Action,
	{_State353, TrueToken}:                               _GotoState34Action,
	{_State353, FalseToken}:                              _GotoState17Action,
	{_State353, StructToken}:                             _GotoState32Action,
	{_State353, FuncToken}:                               _GotoState19Action,
	{_State353, VarToken}:                                _GotoState35Action,
	{_State353, LetToken}:                                _GotoState25Action,
	{_State353, NotToken}:                                _GotoState28Action,
	{_State353, LabelDeclToken}:                          _GotoState23Action,
	{_State353, LparenToken}:                             _GotoState26Action,
	{_State353, LbracketToken}:                           _GotoState24Action,
	{_State353, SubToken}:                                _GotoState33Action,
	{_State353, MulToken}:                                _GotoState27Action,
	{_State353, BitNegToken}:                             _GotoState16Action,
	{_State353, BitAndToken}:                             _GotoState15Action,
	{_State353, GreaterToken}:                            _GotoState20Action,
	{_State353, ParseErrorToken}:                         _GotoState29Action,
	{_State353, VarDeclPatternType}:                      _GotoState70Action,
	{_State353, VarOrLetType}:                            _GotoState71Action,
	{_State353, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State353, SequenceExprType}:                        _GotoState416Action,
	{_State353, CallExprType}:                            _GotoState49Action,
	{_State353, AtomExprType}:                            _GotoState42Action,
	{_State353, ParseErrorExprType}:                      _GotoState62Action,
	{_State353, LiteralExprType}:                         _GotoState57Action,
	{_State353, IdentifierExprType}:                      _GotoState52Action,
	{_State353, BlockExprType}:                           _GotoState48Action,
	{_State353, InitializeExprType}:                      _GotoState56Action,
	{_State353, ImplicitStructExprType}:                  _GotoState53Action,
	{_State353, AccessibleExprType}:                      _GotoState37Action,
	{_State353, AccessExprType}:                          _GotoState36Action,
	{_State353, IndexExprType}:                           _GotoState54Action,
	{_State353, PostfixableExprType}:                     _GotoState64Action,
	{_State353, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State353, PrefixableExprType}:                      _GotoState67Action,
	{_State353, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State353, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State353, MulExprType}:                             _GotoState59Action,
	{_State353, BinaryMulExprType}:                       _GotoState46Action,
	{_State353, AddExprType}:                             _GotoState38Action,
	{_State353, BinaryAddExprType}:                       _GotoState43Action,
	{_State353, CmpExprType}:                             _GotoState50Action,
	{_State353, BinaryCmpExprType}:                       _GotoState45Action,
	{_State353, AndExprType}:                             _GotoState39Action,
	{_State353, BinaryAndExprType}:                       _GotoState44Action,
	{_State353, OrExprType}:                              _GotoState61Action,
	{_State353, BinaryOrExprType}:                        _GotoState47Action,
	{_State353, InitializableTypeType}:                   _GotoState55Action,
	{_State353, SliceTypeType}:                           _GotoState69Action,
	{_State353, ArrayTypeType}:                           _GotoState41Action,
	{_State353, MapTypeType}:                             _GotoState58Action,
	{_State353, ExplicitStructDefType}:                   _GotoState51Action,
	{_State353, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State354, IntegerLiteralToken}:                     _GotoState22Action,
	{_State354, FloatLiteralToken}:                       _GotoState18Action,
	{_State354, RuneLiteralToken}:                        _GotoState30Action,
	{_State354, StringLiteralToken}:                      _GotoState31Action,
	{_State354, IdentifierToken}:                         _GotoState21Action,
	{_State354, TrueToken}:                               _GotoState34Action,
	{_State354, FalseToken}:                              _GotoState17Action,
	{_State354, StructToken}:                             _GotoState32Action,
	{_State354, FuncToken}:                               _GotoState19Action,
	{_State354, VarToken}:                                _GotoState35Action,
	{_State354, LetToken}:                                _GotoState25Action,
	{_State354, NotToken}:                                _GotoState28Action,
	{_State354, LabelDeclToken}:                          _GotoState23Action,
	{_State354, LparenToken}:                             _GotoState26Action,
	{_State354, LbracketToken}:                           _GotoState24Action,
	{_State354, SubToken}:                                _GotoState33Action,
	{_State354, MulToken}:                                _GotoState27Action,
	{_State354, BitNegToken}:                             _GotoState16Action,
	{_State354, BitAndToken}:                             _GotoState15Action,
	{_State354, GreaterToken}:                            _GotoState20Action,
	{_State354, ParseErrorToken}:                         _GotoState29Action,
	{_State354, VarDeclPatternType}:                      _GotoState70Action,
	{_State354, VarOrLetType}:                            _GotoState71Action,
	{_State354, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State354, SequenceExprType}:                        _GotoState418Action,
	{_State354, OptionalSequenceExprType}:                _GotoState417Action,
	{_State354, CallExprType}:                            _GotoState49Action,
	{_State354, AtomExprType}:                            _GotoState42Action,
	{_State354, ParseErrorExprType}:                      _GotoState62Action,
	{_State354, LiteralExprType}:                         _GotoState57Action,
	{_State354, IdentifierExprType}:                      _GotoState52Action,
	{_State354, BlockExprType}:                           _GotoState48Action,
	{_State354, InitializeExprType}:                      _GotoState56Action,
	{_State354, ImplicitStructExprType}:                  _GotoState53Action,
	{_State354, AccessibleExprType}:                      _GotoState37Action,
	{_State354, AccessExprType}:                          _GotoState36Action,
	{_State354, IndexExprType}:                           _GotoState54Action,
	{_State354, PostfixableExprType}:                     _GotoState64Action,
	{_State354, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State354, PrefixableExprType}:                      _GotoState67Action,
	{_State354, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State354, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State354, MulExprType}:                             _GotoState59Action,
	{_State354, BinaryMulExprType}:                       _GotoState46Action,
	{_State354, AddExprType}:                             _GotoState38Action,
	{_State354, BinaryAddExprType}:                       _GotoState43Action,
	{_State354, CmpExprType}:                             _GotoState50Action,
	{_State354, BinaryCmpExprType}:                       _GotoState45Action,
	{_State354, AndExprType}:                             _GotoState39Action,
	{_State354, BinaryAndExprType}:                       _GotoState44Action,
	{_State354, OrExprType}:                              _GotoState61Action,
	{_State354, BinaryOrExprType}:                        _GotoState47Action,
	{_State354, InitializableTypeType}:                   _GotoState55Action,
	{_State354, SliceTypeType}:                           _GotoState69Action,
	{_State354, ArrayTypeType}:                           _GotoState41Action,
	{_State354, MapTypeType}:                             _GotoState58Action,
	{_State354, ExplicitStructDefType}:                   _GotoState51Action,
	{_State354, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State355, LbraceToken}:                             _GotoState73Action,
	{_State355, StatementBlockType}:                      _GotoState419Action,
	{_State356, CommaToken}:                              _GotoState365Action,
	{_State356, AssignToken}:                             _GotoState420Action,
	{_State357, ElseToken}:                               _GotoState421Action,
	{_State359, IdentifierToken}:                         _GotoState157Action,
	{_State359, LparenToken}:                             _GotoState158Action,
	{_State359, VarPatternType}:                          _GotoState422Action,
	{_State359, TuplePatternType}:                        _GotoState159Action,
	{_State360, IdentifierToken}:                         _GotoState256Action,
	{_State360, LparenToken}:                             _GotoState158Action,
	{_State360, EllipsisToken}:                           _GotoState255Action,
	{_State360, VarPatternType}:                          _GotoState259Action,
	{_State360, TuplePatternType}:                        _GotoState159Action,
	{_State360, FieldVarPatternType}:                     _GotoState423Action,
	{_State362, LparenToken}:                             _GotoState26Action,
	{_State362, ImplicitStructExprType}:                  _GotoState424Action,
	{_State363, IdentifierToken}:                         _GotoState425Action,
	{_State364, IntegerLiteralToken}:                     _GotoState22Action,
	{_State364, FloatLiteralToken}:                       _GotoState18Action,
	{_State364, RuneLiteralToken}:                        _GotoState30Action,
	{_State364, StringLiteralToken}:                      _GotoState31Action,
	{_State364, IdentifierToken}:                         _GotoState21Action,
	{_State364, TrueToken}:                               _GotoState34Action,
	{_State364, FalseToken}:                              _GotoState17Action,
	{_State364, ReturnToken}:                             _GotoState169Action,
	{_State364, BreakToken}:                              _GotoState162Action,
	{_State364, ContinueToken}:                           _GotoState164Action,
	{_State364, FallthroughToken}:                        _GotoState167Action,
	{_State364, UnsafeToken}:                             _GotoState170Action,
	{_State364, StructToken}:                             _GotoState32Action,
	{_State364, FuncToken}:                               _GotoState19Action,
	{_State364, AsyncToken}:                              _GotoState161Action,
	{_State364, DeferToken}:                              _GotoState166Action,
	{_State364, VarToken}:                                _GotoState35Action,
	{_State364, LetToken}:                                _GotoState25Action,
	{_State364, NotToken}:                                _GotoState28Action,
	{_State364, LabelDeclToken}:                          _GotoState23Action,
	{_State364, LparenToken}:                             _GotoState26Action,
	{_State364, LbracketToken}:                           _GotoState24Action,
	{_State364, SubToken}:                                _GotoState33Action,
	{_State364, MulToken}:                                _GotoState27Action,
	{_State364, BitNegToken}:                             _GotoState16Action,
	{_State364, BitAndToken}:                             _GotoState15Action,
	{_State364, GreaterToken}:                            _GotoState20Action,
	{_State364, ParseErrorToken}:                         _GotoState29Action,
	{_State364, SimpleStatementType}:                     _GotoState367Action,
	{_State364, OptionalSimpleStatementType}:             _GotoState426Action,
	{_State364, ExpressionOrImproperStructStatementType}: _GotoState178Action,
	{_State364, ExpressionsType}:                         _GotoState179Action,
	{_State364, CallbackOpType}:                          _GotoState175Action,
	{_State364, CallbackOpStatementType}:                 _GotoState176Action,
	{_State364, UnsafeStatementType}:                     _GotoState190Action,
	{_State364, JumpStatementType}:                       _GotoState182Action,
	{_State364, JumpTypeType}:                            _GotoState183Action,
	{_State364, FallthroughStatementType}:                _GotoState180Action,
	{_State364, AssignStatementType}:                     _GotoState173Action,
	{_State364, UnaryOpAssignStatementType}:              _GotoState189Action,
	{_State364, BinaryOpAssignStatementType}:             _GotoState174Action,
	{_State364, VarDeclPatternType}:                      _GotoState70Action,
	{_State364, VarOrLetType}:                            _GotoState71Action,
	{_State364, AssignPatternType}:                       _GotoState172Action,
	{_State364, ExpressionType}:                          _GotoState177Action,
	{_State364, OptionalLabelDeclType}:                   _GotoState60Action,
	{_State364, SequenceExprType}:                        _GotoState185Action,
	{_State364, CallExprType}:                            _GotoState49Action,
	{_State364, AtomExprType}:                            _GotoState42Action,
	{_State364, ParseErrorExprType}:                      _GotoState62Action,
	{_State364, LiteralExprType}:                         _GotoState57Action,
	{_State364, IdentifierExprType}:                      _GotoState52Action,
	{_State364, BlockExprType}:                           _GotoState48Action,
	{_State364, InitializeExprType}:                      _GotoState56Action,
	{_State364, ImplicitStructExprType}:                  _GotoState53Action,
	{_State364, AccessibleExprType}:                      _GotoState171Action,
	{_State364, AccessExprType}:                          _GotoState36Action,
	{_State364, IndexExprType}:                           _GotoState54Action,
	{_State364, PostfixableExprType}:                     _GotoState64Action,
	{_State364, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State364, PrefixableExprType}:                      _GotoState67Action,
	{_State364, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State364, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State364, MulExprType}:                             _GotoState59Action,
	{_State364, BinaryMulExprType}:                       _GotoState46Action,
	{_State364, AddExprType}:                             _GotoState38Action,
	{_State364, BinaryAddExprType}:                       _GotoState43Action,
	{_State364, CmpExprType}:                             _GotoState50Action,
	{_State364, BinaryCmpExprType}:                       _GotoState45Action,
	{_State364, AndExprType}:                             _GotoState39Action,
	{_State364, BinaryAndExprType}:                       _GotoState44Action,
	{_State364, OrExprType}:                              _GotoState61Action,
	{_State364, BinaryOrExprType}:                        _GotoState47Action,
	{_State364, InitializableTypeType}:                   _GotoState55Action,
	{_State364, SliceTypeType}:                           _GotoState69Action,
	{_State364, ArrayTypeType}:                           _GotoState41Action,
	{_State364, MapTypeType}:                             _GotoState58Action,
	{_State364, ExplicitStructDefType}:                   _GotoState51Action,
	{_State364, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State365, IntegerLiteralToken}:                     _GotoState22Action,
	{_State365, FloatLiteralToken}:                       _GotoState18Action,
	{_State365, RuneLiteralToken}:                        _GotoState30Action,
	{_State365, StringLiteralToken}:                      _GotoState31Action,
	{_State365, IdentifierToken}:                         _GotoState21Action,
	{_State365, TrueToken}:                               _GotoState34Action,
	{_State365, FalseToken}:                              _GotoState17Action,
	{_State365, StructToken}:                             _GotoState32Action,
	{_State365, FuncToken}:                               _GotoState19Action,
	{_State365, VarToken}:                                _GotoState263Action,
	{_State365, LetToken}:                                _GotoState25Action,
	{_State365, NotToken}:                                _GotoState28Action,
	{_State365, LabelDeclToken}:                          _GotoState23Action,
	{_State365, LparenToken}:                             _GotoState26Action,
	{_State365, LbracketToken}:                           _GotoState24Action,
	{_State365, DotToken}:                                _GotoState262Action,
	{_State365, SubToken}:                                _GotoState33Action,
	{_State365, MulToken}:                                _GotoState27Action,
	{_State365, BitNegToken}:                             _GotoState16Action,
	{_State365, BitAndToken}:                             _GotoState15Action,
	{_State365, GreaterToken}:                            _GotoState20Action,
	{_State365, ParseErrorToken}:                         _GotoState29Action,
	{_State365, VarDeclPatternType}:                      _GotoState70Action,
	{_State365, VarOrLetType}:                            _GotoState71Action,
	{_State365, AssignPatternType}:                       _GotoState264Action,
	{_State365, CasePatternType}:                         _GotoState427Action,
	{_State365, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State365, SequenceExprType}:                        _GotoState267Action,
	{_State365, CallExprType}:                            _GotoState49Action,
	{_State365, AtomExprType}:                            _GotoState42Action,
	{_State365, ParseErrorExprType}:                      _GotoState62Action,
	{_State365, LiteralExprType}:                         _GotoState57Action,
	{_State365, IdentifierExprType}:                      _GotoState52Action,
	{_State365, BlockExprType}:                           _GotoState48Action,
	{_State365, InitializeExprType}:                      _GotoState56Action,
	{_State365, ImplicitStructExprType}:                  _GotoState53Action,
	{_State365, AccessibleExprType}:                      _GotoState37Action,
	{_State365, AccessExprType}:                          _GotoState36Action,
	{_State365, IndexExprType}:                           _GotoState54Action,
	{_State365, PostfixableExprType}:                     _GotoState64Action,
	{_State365, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State365, PrefixableExprType}:                      _GotoState67Action,
	{_State365, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State365, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State365, MulExprType}:                             _GotoState59Action,
	{_State365, BinaryMulExprType}:                       _GotoState46Action,
	{_State365, AddExprType}:                             _GotoState38Action,
	{_State365, BinaryAddExprType}:                       _GotoState43Action,
	{_State365, CmpExprType}:                             _GotoState50Action,
	{_State365, BinaryCmpExprType}:                       _GotoState45Action,
	{_State365, AndExprType}:                             _GotoState39Action,
	{_State365, BinaryAndExprType}:                       _GotoState44Action,
	{_State365, OrExprType}:                              _GotoState61Action,
	{_State365, BinaryOrExprType}:                        _GotoState47Action,
	{_State365, InitializableTypeType}:                   _GotoState55Action,
	{_State365, SliceTypeType}:                           _GotoState69Action,
	{_State365, ArrayTypeType}:                           _GotoState41Action,
	{_State365, MapTypeType}:                             _GotoState58Action,
	{_State365, ExplicitStructDefType}:                   _GotoState51Action,
	{_State365, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State369, RparenToken}:                             _GotoState428Action,
	{_State370, NewlinesToken}:                           _GotoState430Action,
	{_State370, CommaToken}:                              _GotoState429Action,
	{_State371, IdentifierToken}:                         _GotoState431Action,
	{_State372, GreaterToken}:                            _GotoState432Action,
	{_State376, CommaToken}:                              _GotoState291Action,
	{_State379, AddToken}:                                _GotoState218Action,
	{_State379, SubToken}:                                _GotoState223Action,
	{_State379, MulToken}:                                _GotoState221Action,
	{_State379, TraitOpType}:                             _GotoState224Action,
	{_State381, IdentifierToken}:                         _GotoState299Action,
	{_State381, GenericParameterDefType}:                 _GotoState433Action,
	{_State382, IdentifierToken}:                         _GotoState93Action,
	{_State382, StructToken}:                             _GotoState32Action,
	{_State382, EnumToken}:                               _GotoState90Action,
	{_State382, TraitToken}:                              _GotoState98Action,
	{_State382, FuncToken}:                               _GotoState92Action,
	{_State382, LparenToken}:                             _GotoState94Action,
	{_State382, LbracketToken}:                           _GotoState24Action,
	{_State382, DotToken}:                                _GotoState89Action,
	{_State382, QuestionToken}:                           _GotoState96Action,
	{_State382, ExclaimToken}:                            _GotoState91Action,
	{_State382, TildeTildeToken}:                         _GotoState97Action,
	{_State382, BitNegToken}:                             _GotoState88Action,
	{_State382, BitAndToken}:                             _GotoState87Action,
	{_State382, ParseErrorToken}:                         _GotoState95Action,
	{_State382, InitializableTypeType}:                   _GotoState104Action,
	{_State382, SliceTypeType}:                           _GotoState69Action,
	{_State382, ArrayTypeType}:                           _GotoState41Action,
	{_State382, MapTypeType}:                             _GotoState58Action,
	{_State382, AtomTypeType}:                            _GotoState99Action,
	{_State382, ParseErrorTypeType}:                      _GotoState105Action,
	{_State382, ReturnableTypeType}:                      _GotoState108Action,
	{_State382, PrefixedTypeType}:                        _GotoState107Action,
	{_State382, PrefixTypeOpType}:                        _GotoState106Action,
	{_State382, ValueTypeType}:                           _GotoState434Action,
	{_State382, TraitOpTypeType}:                         _GotoState110Action,
	{_State382, ImplicitStructDefType}:                   _GotoState103Action,
	{_State382, ExplicitStructDefType}:                   _GotoState51Action,
	{_State382, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State382, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State382, TraitDefType}:                            _GotoState109Action,
	{_State382, FuncTypeType}:                            _GotoState101Action,
	{_State383, RparenToken}:                             _GotoState435Action,
	{_State384, AddToken}:                                _GotoState218Action,
	{_State384, SubToken}:                                _GotoState223Action,
	{_State384, MulToken}:                                _GotoState221Action,
	{_State384, TraitOpType}:                             _GotoState224Action,
	{_State385, DollarLbracketToken}:                     _GotoState194Action,
	{_State385, OptionalGenericParametersType}:           _GotoState436Action,
	{_State386, LbraceToken}:                             _GotoState73Action,
	{_State386, StatementBlockType}:                      _GotoState437Action,
	{_State389, IdentifierToken}:                         _GotoState208Action,
	{_State389, UnsafeToken}:                             _GotoState170Action,
	{_State389, StructToken}:                             _GotoState32Action,
	{_State389, EnumToken}:                               _GotoState90Action,
	{_State389, TraitToken}:                              _GotoState98Action,
	{_State389, FuncToken}:                               _GotoState92Action,
	{_State389, LparenToken}:                             _GotoState94Action,
	{_State389, LbracketToken}:                           _GotoState24Action,
	{_State389, DotToken}:                                _GotoState89Action,
	{_State389, QuestionToken}:                           _GotoState96Action,
	{_State389, ExclaimToken}:                            _GotoState91Action,
	{_State389, TildeTildeToken}:                         _GotoState97Action,
	{_State389, BitNegToken}:                             _GotoState88Action,
	{_State389, BitAndToken}:                             _GotoState87Action,
	{_State389, ParseErrorToken}:                         _GotoState95Action,
	{_State389, UnsafeStatementType}:                     _GotoState214Action,
	{_State389, InitializableTypeType}:                   _GotoState104Action,
	{_State389, SliceTypeType}:                           _GotoState69Action,
	{_State389, ArrayTypeType}:                           _GotoState41Action,
	{_State389, MapTypeType}:                             _GotoState58Action,
	{_State389, AtomTypeType}:                            _GotoState99Action,
	{_State389, ParseErrorTypeType}:                      _GotoState105Action,
	{_State389, ReturnableTypeType}:                      _GotoState108Action,
	{_State389, PrefixedTypeType}:                        _GotoState107Action,
	{_State389, PrefixTypeOpType}:                        _GotoState106Action,
	{_State389, ValueTypeType}:                           _GotoState215Action,
	{_State389, TraitOpTypeType}:                         _GotoState110Action,
	{_State389, FieldDefType}:                            _GotoState313Action,
	{_State389, ImplicitStructDefType}:                   _GotoState103Action,
	{_State389, ExplicitStructDefType}:                   _GotoState51Action,
	{_State389, EnumValueDefType}:                        _GotoState438Action,
	{_State389, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State389, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State389, TraitDefType}:                            _GotoState109Action,
	{_State389, FuncTypeType}:                            _GotoState101Action,
	{_State390, IdentifierToken}:                         _GotoState208Action,
	{_State390, UnsafeToken}:                             _GotoState170Action,
	{_State390, StructToken}:                             _GotoState32Action,
	{_State390, EnumToken}:                               _GotoState90Action,
	{_State390, TraitToken}:                              _GotoState98Action,
	{_State390, FuncToken}:                               _GotoState92Action,
	{_State390, LparenToken}:                             _GotoState94Action,
	{_State390, LbracketToken}:                           _GotoState24Action,
	{_State390, DotToken}:                                _GotoState89Action,
	{_State390, QuestionToken}:                           _GotoState96Action,
	{_State390, ExclaimToken}:                            _GotoState91Action,
	{_State390, TildeTildeToken}:                         _GotoState97Action,
	{_State390, BitNegToken}:                             _GotoState88Action,
	{_State390, BitAndToken}:                             _GotoState87Action,
	{_State390, ParseErrorToken}:                         _GotoState95Action,
	{_State390, UnsafeStatementType}:                     _GotoState214Action,
	{_State390, InitializableTypeType}:                   _GotoState104Action,
	{_State390, SliceTypeType}:                           _GotoState69Action,
	{_State390, ArrayTypeType}:                           _GotoState41Action,
	{_State390, MapTypeType}:                             _GotoState58Action,
	{_State390, AtomTypeType}:                            _GotoState99Action,
	{_State390, ParseErrorTypeType}:                      _GotoState105Action,
	{_State390, ReturnableTypeType}:                      _GotoState108Action,
	{_State390, PrefixedTypeType}:                        _GotoState107Action,
	{_State390, PrefixTypeOpType}:                        _GotoState106Action,
	{_State390, ValueTypeType}:                           _GotoState215Action,
	{_State390, TraitOpTypeType}:                         _GotoState110Action,
	{_State390, FieldDefType}:                            _GotoState313Action,
	{_State390, ImplicitStructDefType}:                   _GotoState103Action,
	{_State390, ExplicitStructDefType}:                   _GotoState51Action,
	{_State390, EnumValueDefType}:                        _GotoState439Action,
	{_State390, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State390, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State390, TraitDefType}:                            _GotoState109Action,
	{_State390, FuncTypeType}:                            _GotoState101Action,
	{_State392, IdentifierToken}:                         _GotoState208Action,
	{_State392, UnsafeToken}:                             _GotoState170Action,
	{_State392, StructToken}:                             _GotoState32Action,
	{_State392, EnumToken}:                               _GotoState90Action,
	{_State392, TraitToken}:                              _GotoState98Action,
	{_State392, FuncToken}:                               _GotoState92Action,
	{_State392, LparenToken}:                             _GotoState94Action,
	{_State392, LbracketToken}:                           _GotoState24Action,
	{_State392, DotToken}:                                _GotoState89Action,
	{_State392, QuestionToken}:                           _GotoState96Action,
	{_State392, ExclaimToken}:                            _GotoState91Action,
	{_State392, TildeTildeToken}:                         _GotoState97Action,
	{_State392, BitNegToken}:                             _GotoState88Action,
	{_State392, BitAndToken}:                             _GotoState87Action,
	{_State392, ParseErrorToken}:                         _GotoState95Action,
	{_State392, UnsafeStatementType}:                     _GotoState214Action,
	{_State392, InitializableTypeType}:                   _GotoState104Action,
	{_State392, SliceTypeType}:                           _GotoState69Action,
	{_State392, ArrayTypeType}:                           _GotoState41Action,
	{_State392, MapTypeType}:                             _GotoState58Action,
	{_State392, AtomTypeType}:                            _GotoState99Action,
	{_State392, ParseErrorTypeType}:                      _GotoState105Action,
	{_State392, ReturnableTypeType}:                      _GotoState108Action,
	{_State392, PrefixedTypeType}:                        _GotoState107Action,
	{_State392, PrefixTypeOpType}:                        _GotoState106Action,
	{_State392, ValueTypeType}:                           _GotoState215Action,
	{_State392, TraitOpTypeType}:                         _GotoState110Action,
	{_State392, FieldDefType}:                            _GotoState313Action,
	{_State392, ImplicitStructDefType}:                   _GotoState103Action,
	{_State392, ExplicitStructDefType}:                   _GotoState51Action,
	{_State392, EnumValueDefType}:                        _GotoState440Action,
	{_State392, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State392, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State392, TraitDefType}:                            _GotoState109Action,
	{_State392, FuncTypeType}:                            _GotoState101Action,
	{_State393, IdentifierToken}:                         _GotoState208Action,
	{_State393, UnsafeToken}:                             _GotoState170Action,
	{_State393, StructToken}:                             _GotoState32Action,
	{_State393, EnumToken}:                               _GotoState90Action,
	{_State393, TraitToken}:                              _GotoState98Action,
	{_State393, FuncToken}:                               _GotoState92Action,
	{_State393, LparenToken}:                             _GotoState94Action,
	{_State393, LbracketToken}:                           _GotoState24Action,
	{_State393, DotToken}:                                _GotoState89Action,
	{_State393, QuestionToken}:                           _GotoState96Action,
	{_State393, ExclaimToken}:                            _GotoState91Action,
	{_State393, TildeTildeToken}:                         _GotoState97Action,
	{_State393, BitNegToken}:                             _GotoState88Action,
	{_State393, BitAndToken}:                             _GotoState87Action,
	{_State393, ParseErrorToken}:                         _GotoState95Action,
	{_State393, UnsafeStatementType}:                     _GotoState214Action,
	{_State393, InitializableTypeType}:                   _GotoState104Action,
	{_State393, SliceTypeType}:                           _GotoState69Action,
	{_State393, ArrayTypeType}:                           _GotoState41Action,
	{_State393, MapTypeType}:                             _GotoState58Action,
	{_State393, AtomTypeType}:                            _GotoState99Action,
	{_State393, ParseErrorTypeType}:                      _GotoState105Action,
	{_State393, ReturnableTypeType}:                      _GotoState108Action,
	{_State393, PrefixedTypeType}:                        _GotoState107Action,
	{_State393, PrefixTypeOpType}:                        _GotoState106Action,
	{_State393, ValueTypeType}:                           _GotoState215Action,
	{_State393, TraitOpTypeType}:                         _GotoState110Action,
	{_State393, FieldDefType}:                            _GotoState313Action,
	{_State393, ImplicitStructDefType}:                   _GotoState103Action,
	{_State393, ExplicitStructDefType}:                   _GotoState51Action,
	{_State393, EnumValueDefType}:                        _GotoState441Action,
	{_State393, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State393, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State393, TraitDefType}:                            _GotoState109Action,
	{_State393, FuncTypeType}:                            _GotoState101Action,
	{_State394, AddToken}:                                _GotoState218Action,
	{_State394, SubToken}:                                _GotoState223Action,
	{_State394, MulToken}:                                _GotoState221Action,
	{_State394, TraitOpType}:                             _GotoState224Action,
	{_State395, IdentifierToken}:                         _GotoState93Action,
	{_State395, StructToken}:                             _GotoState32Action,
	{_State395, EnumToken}:                               _GotoState90Action,
	{_State395, TraitToken}:                              _GotoState98Action,
	{_State395, FuncToken}:                               _GotoState92Action,
	{_State395, LparenToken}:                             _GotoState94Action,
	{_State395, LbracketToken}:                           _GotoState24Action,
	{_State395, DotToken}:                                _GotoState89Action,
	{_State395, QuestionToken}:                           _GotoState96Action,
	{_State395, ExclaimToken}:                            _GotoState91Action,
	{_State395, TildeTildeToken}:                         _GotoState97Action,
	{_State395, BitNegToken}:                             _GotoState88Action,
	{_State395, BitAndToken}:                             _GotoState87Action,
	{_State395, ParseErrorToken}:                         _GotoState95Action,
	{_State395, InitializableTypeType}:                   _GotoState104Action,
	{_State395, SliceTypeType}:                           _GotoState69Action,
	{_State395, ArrayTypeType}:                           _GotoState41Action,
	{_State395, MapTypeType}:                             _GotoState58Action,
	{_State395, AtomTypeType}:                            _GotoState99Action,
	{_State395, ParseErrorTypeType}:                      _GotoState105Action,
	{_State395, ReturnableTypeType}:                      _GotoState108Action,
	{_State395, PrefixedTypeType}:                        _GotoState107Action,
	{_State395, PrefixTypeOpType}:                        _GotoState106Action,
	{_State395, ValueTypeType}:                           _GotoState442Action,
	{_State395, TraitOpTypeType}:                         _GotoState110Action,
	{_State395, ImplicitStructDefType}:                   _GotoState103Action,
	{_State395, ExplicitStructDefType}:                   _GotoState51Action,
	{_State395, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State395, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State395, TraitDefType}:                            _GotoState109Action,
	{_State395, FuncTypeType}:                            _GotoState101Action,
	{_State396, AddToken}:                                _GotoState218Action,
	{_State396, SubToken}:                                _GotoState223Action,
	{_State396, MulToken}:                                _GotoState221Action,
	{_State396, TraitOpType}:                             _GotoState224Action,
	{_State397, IdentifierToken}:                         _GotoState93Action,
	{_State397, StructToken}:                             _GotoState32Action,
	{_State397, EnumToken}:                               _GotoState90Action,
	{_State397, TraitToken}:                              _GotoState98Action,
	{_State397, FuncToken}:                               _GotoState92Action,
	{_State397, LparenToken}:                             _GotoState94Action,
	{_State397, LbracketToken}:                           _GotoState24Action,
	{_State397, DotToken}:                                _GotoState89Action,
	{_State397, QuestionToken}:                           _GotoState96Action,
	{_State397, ExclaimToken}:                            _GotoState91Action,
	{_State397, TildeTildeToken}:                         _GotoState97Action,
	{_State397, BitNegToken}:                             _GotoState88Action,
	{_State397, BitAndToken}:                             _GotoState87Action,
	{_State397, ParseErrorToken}:                         _GotoState95Action,
	{_State397, InitializableTypeType}:                   _GotoState104Action,
	{_State397, SliceTypeType}:                           _GotoState69Action,
	{_State397, ArrayTypeType}:                           _GotoState41Action,
	{_State397, MapTypeType}:                             _GotoState58Action,
	{_State397, AtomTypeType}:                            _GotoState99Action,
	{_State397, ParseErrorTypeType}:                      _GotoState105Action,
	{_State397, ReturnableTypeType}:                      _GotoState387Action,
	{_State397, PrefixedTypeType}:                        _GotoState107Action,
	{_State397, PrefixTypeOpType}:                        _GotoState106Action,
	{_State397, ImplicitStructDefType}:                   _GotoState103Action,
	{_State397, ExplicitStructDefType}:                   _GotoState51Action,
	{_State397, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State397, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State397, TraitDefType}:                            _GotoState109Action,
	{_State397, ReturnTypeType}:                          _GotoState443Action,
	{_State397, FuncTypeType}:                            _GotoState101Action,
	{_State398, IdentifierToken}:                         _GotoState316Action,
	{_State398, StructToken}:                             _GotoState32Action,
	{_State398, EnumToken}:                               _GotoState90Action,
	{_State398, TraitToken}:                              _GotoState98Action,
	{_State398, FuncToken}:                               _GotoState92Action,
	{_State398, LparenToken}:                             _GotoState94Action,
	{_State398, LbracketToken}:                           _GotoState24Action,
	{_State398, DotToken}:                                _GotoState89Action,
	{_State398, QuestionToken}:                           _GotoState96Action,
	{_State398, ExclaimToken}:                            _GotoState91Action,
	{_State398, EllipsisToken}:                           _GotoState315Action,
	{_State398, TildeTildeToken}:                         _GotoState97Action,
	{_State398, BitNegToken}:                             _GotoState88Action,
	{_State398, BitAndToken}:                             _GotoState87Action,
	{_State398, ParseErrorToken}:                         _GotoState95Action,
	{_State398, InitializableTypeType}:                   _GotoState104Action,
	{_State398, SliceTypeType}:                           _GotoState69Action,
	{_State398, ArrayTypeType}:                           _GotoState41Action,
	{_State398, MapTypeType}:                             _GotoState58Action,
	{_State398, AtomTypeType}:                            _GotoState99Action,
	{_State398, ParseErrorTypeType}:                      _GotoState105Action,
	{_State398, ReturnableTypeType}:                      _GotoState108Action,
	{_State398, PrefixedTypeType}:                        _GotoState107Action,
	{_State398, PrefixTypeOpType}:                        _GotoState106Action,
	{_State398, ValueTypeType}:                           _GotoState320Action,
	{_State398, TraitOpTypeType}:                         _GotoState110Action,
	{_State398, ImplicitStructDefType}:                   _GotoState103Action,
	{_State398, ExplicitStructDefType}:                   _GotoState51Action,
	{_State398, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State398, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State398, TraitDefType}:                            _GotoState109Action,
	{_State398, ParameterDeclType}:                       _GotoState444Action,
	{_State398, FuncTypeType}:                            _GotoState101Action,
	{_State404, LparenToken}:                             _GotoState445Action,
	{_State405, IdentifierToken}:                         _GotoState208Action,
	{_State405, UnsafeToken}:                             _GotoState170Action,
	{_State405, StructToken}:                             _GotoState32Action,
	{_State405, EnumToken}:                               _GotoState90Action,
	{_State405, TraitToken}:                              _GotoState98Action,
	{_State405, FuncToken}:                               _GotoState330Action,
	{_State405, LparenToken}:                             _GotoState94Action,
	{_State405, LbracketToken}:                           _GotoState24Action,
	{_State405, DotToken}:                                _GotoState89Action,
	{_State405, QuestionToken}:                           _GotoState96Action,
	{_State405, ExclaimToken}:                            _GotoState91Action,
	{_State405, TildeTildeToken}:                         _GotoState97Action,
	{_State405, BitNegToken}:                             _GotoState88Action,
	{_State405, BitAndToken}:                             _GotoState87Action,
	{_State405, ParseErrorToken}:                         _GotoState95Action,
	{_State405, UnsafeStatementType}:                     _GotoState214Action,
	{_State405, InitializableTypeType}:                   _GotoState104Action,
	{_State405, SliceTypeType}:                           _GotoState69Action,
	{_State405, ArrayTypeType}:                           _GotoState41Action,
	{_State405, MapTypeType}:                             _GotoState58Action,
	{_State405, AtomTypeType}:                            _GotoState99Action,
	{_State405, ParseErrorTypeType}:                      _GotoState105Action,
	{_State405, ReturnableTypeType}:                      _GotoState108Action,
	{_State405, PrefixedTypeType}:                        _GotoState107Action,
	{_State405, PrefixTypeOpType}:                        _GotoState106Action,
	{_State405, ValueTypeType}:                           _GotoState215Action,
	{_State405, TraitOpTypeType}:                         _GotoState110Action,
	{_State405, FieldDefType}:                            _GotoState331Action,
	{_State405, ImplicitStructDefType}:                   _GotoState103Action,
	{_State405, ExplicitStructDefType}:                   _GotoState51Action,
	{_State405, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State405, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State405, TraitPropertyType}:                       _GotoState446Action,
	{_State405, TraitDefType}:                            _GotoState109Action,
	{_State405, FuncTypeType}:                            _GotoState101Action,
	{_State405, MethodSignatureType}:                     _GotoState332Action,
	{_State406, IdentifierToken}:                         _GotoState208Action,
	{_State406, UnsafeToken}:                             _GotoState170Action,
	{_State406, StructToken}:                             _GotoState32Action,
	{_State406, EnumToken}:                               _GotoState90Action,
	{_State406, TraitToken}:                              _GotoState98Action,
	{_State406, FuncToken}:                               _GotoState330Action,
	{_State406, LparenToken}:                             _GotoState94Action,
	{_State406, LbracketToken}:                           _GotoState24Action,
	{_State406, DotToken}:                                _GotoState89Action,
	{_State406, QuestionToken}:                           _GotoState96Action,
	{_State406, ExclaimToken}:                            _GotoState91Action,
	{_State406, TildeTildeToken}:                         _GotoState97Action,
	{_State406, BitNegToken}:                             _GotoState88Action,
	{_State406, BitAndToken}:                             _GotoState87Action,
	{_State406, ParseErrorToken}:                         _GotoState95Action,
	{_State406, UnsafeStatementType}:                     _GotoState214Action,
	{_State406, InitializableTypeType}:                   _GotoState104Action,
	{_State406, SliceTypeType}:                           _GotoState69Action,
	{_State406, ArrayTypeType}:                           _GotoState41Action,
	{_State406, MapTypeType}:                             _GotoState58Action,
	{_State406, AtomTypeType}:                            _GotoState99Action,
	{_State406, ParseErrorTypeType}:                      _GotoState105Action,
	{_State406, ReturnableTypeType}:                      _GotoState108Action,
	{_State406, PrefixedTypeType}:                        _GotoState107Action,
	{_State406, PrefixTypeOpType}:                        _GotoState106Action,
	{_State406, ValueTypeType}:                           _GotoState215Action,
	{_State406, TraitOpTypeType}:                         _GotoState110Action,
	{_State406, FieldDefType}:                            _GotoState331Action,
	{_State406, ImplicitStructDefType}:                   _GotoState103Action,
	{_State406, ExplicitStructDefType}:                   _GotoState51Action,
	{_State406, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State406, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State406, TraitPropertyType}:                       _GotoState447Action,
	{_State406, TraitDefType}:                            _GotoState109Action,
	{_State406, FuncTypeType}:                            _GotoState101Action,
	{_State406, MethodSignatureType}:                     _GotoState332Action,
	{_State412, AddToken}:                                _GotoState218Action,
	{_State412, SubToken}:                                _GotoState223Action,
	{_State412, MulToken}:                                _GotoState221Action,
	{_State412, TraitOpType}:                             _GotoState224Action,
	{_State416, DoToken}:                                 _GotoState448Action,
	{_State417, SemicolonToken}:                          _GotoState449Action,
	{_State420, IntegerLiteralToken}:                     _GotoState22Action,
	{_State420, FloatLiteralToken}:                       _GotoState18Action,
	{_State420, RuneLiteralToken}:                        _GotoState30Action,
	{_State420, StringLiteralToken}:                      _GotoState31Action,
	{_State420, IdentifierToken}:                         _GotoState21Action,
	{_State420, TrueToken}:                               _GotoState34Action,
	{_State420, FalseToken}:                              _GotoState17Action,
	{_State420, StructToken}:                             _GotoState32Action,
	{_State420, FuncToken}:                               _GotoState19Action,
	{_State420, VarToken}:                                _GotoState35Action,
	{_State420, LetToken}:                                _GotoState25Action,
	{_State420, NotToken}:                                _GotoState28Action,
	{_State420, LabelDeclToken}:                          _GotoState23Action,
	{_State420, LparenToken}:                             _GotoState26Action,
	{_State420, LbracketToken}:                           _GotoState24Action,
	{_State420, SubToken}:                                _GotoState33Action,
	{_State420, MulToken}:                                _GotoState27Action,
	{_State420, BitNegToken}:                             _GotoState16Action,
	{_State420, BitAndToken}:                             _GotoState15Action,
	{_State420, GreaterToken}:                            _GotoState20Action,
	{_State420, ParseErrorToken}:                         _GotoState29Action,
	{_State420, VarDeclPatternType}:                      _GotoState70Action,
	{_State420, VarOrLetType}:                            _GotoState71Action,
	{_State420, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State420, SequenceExprType}:                        _GotoState450Action,
	{_State420, CallExprType}:                            _GotoState49Action,
	{_State420, AtomExprType}:                            _GotoState42Action,
	{_State420, ParseErrorExprType}:                      _GotoState62Action,
	{_State420, LiteralExprType}:                         _GotoState57Action,
	{_State420, IdentifierExprType}:                      _GotoState52Action,
	{_State420, BlockExprType}:                           _GotoState48Action,
	{_State420, InitializeExprType}:                      _GotoState56Action,
	{_State420, ImplicitStructExprType}:                  _GotoState53Action,
	{_State420, AccessibleExprType}:                      _GotoState37Action,
	{_State420, AccessExprType}:                          _GotoState36Action,
	{_State420, IndexExprType}:                           _GotoState54Action,
	{_State420, PostfixableExprType}:                     _GotoState64Action,
	{_State420, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State420, PrefixableExprType}:                      _GotoState67Action,
	{_State420, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State420, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State420, MulExprType}:                             _GotoState59Action,
	{_State420, BinaryMulExprType}:                       _GotoState46Action,
	{_State420, AddExprType}:                             _GotoState38Action,
	{_State420, BinaryAddExprType}:                       _GotoState43Action,
	{_State420, CmpExprType}:                             _GotoState50Action,
	{_State420, BinaryCmpExprType}:                       _GotoState45Action,
	{_State420, AndExprType}:                             _GotoState39Action,
	{_State420, BinaryAndExprType}:                       _GotoState44Action,
	{_State420, OrExprType}:                              _GotoState61Action,
	{_State420, BinaryOrExprType}:                        _GotoState47Action,
	{_State420, InitializableTypeType}:                   _GotoState55Action,
	{_State420, SliceTypeType}:                           _GotoState69Action,
	{_State420, ArrayTypeType}:                           _GotoState41Action,
	{_State420, MapTypeType}:                             _GotoState58Action,
	{_State420, ExplicitStructDefType}:                   _GotoState51Action,
	{_State420, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State421, IfToken}:                                 _GotoState149Action,
	{_State421, LbraceToken}:                             _GotoState73Action,
	{_State421, StatementBlockType}:                      _GotoState452Action,
	{_State421, IfExprType}:                              _GotoState451Action,
	{_State425, LparenToken}:                             _GotoState158Action,
	{_State425, TuplePatternType}:                        _GotoState453Action,
	{_State429, StringLiteralToken}:                      _GotoState270Action,
	{_State429, ImportClauseType}:                        _GotoState454Action,
	{_State430, StringLiteralToken}:                      _GotoState270Action,
	{_State430, ImportClauseType}:                        _GotoState455Action,
	{_State432, StringLiteralToken}:                      _GotoState456Action,
	{_State434, AddToken}:                                _GotoState218Action,
	{_State434, SubToken}:                                _GotoState223Action,
	{_State434, MulToken}:                                _GotoState221Action,
	{_State434, TraitOpType}:                             _GotoState224Action,
	{_State435, IdentifierToken}:                         _GotoState93Action,
	{_State435, StructToken}:                             _GotoState32Action,
	{_State435, EnumToken}:                               _GotoState90Action,
	{_State435, TraitToken}:                              _GotoState98Action,
	{_State435, FuncToken}:                               _GotoState92Action,
	{_State435, LparenToken}:                             _GotoState94Action,
	{_State435, LbracketToken}:                           _GotoState24Action,
	{_State435, DotToken}:                                _GotoState89Action,
	{_State435, QuestionToken}:                           _GotoState96Action,
	{_State435, ExclaimToken}:                            _GotoState91Action,
	{_State435, TildeTildeToken}:                         _GotoState97Action,
	{_State435, BitNegToken}:                             _GotoState88Action,
	{_State435, BitAndToken}:                             _GotoState87Action,
	{_State435, ParseErrorToken}:                         _GotoState95Action,
	{_State435, InitializableTypeType}:                   _GotoState104Action,
	{_State435, SliceTypeType}:                           _GotoState69Action,
	{_State435, ArrayTypeType}:                           _GotoState41Action,
	{_State435, MapTypeType}:                             _GotoState58Action,
	{_State435, AtomTypeType}:                            _GotoState99Action,
	{_State435, ParseErrorTypeType}:                      _GotoState105Action,
	{_State435, ReturnableTypeType}:                      _GotoState387Action,
	{_State435, PrefixedTypeType}:                        _GotoState107Action,
	{_State435, PrefixTypeOpType}:                        _GotoState106Action,
	{_State435, ImplicitStructDefType}:                   _GotoState103Action,
	{_State435, ExplicitStructDefType}:                   _GotoState51Action,
	{_State435, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State435, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State435, TraitDefType}:                            _GotoState109Action,
	{_State435, ReturnTypeType}:                          _GotoState457Action,
	{_State435, FuncTypeType}:                            _GotoState101Action,
	{_State436, LparenToken}:                             _GotoState458Action,
	{_State442, AddToken}:                                _GotoState218Action,
	{_State442, SubToken}:                                _GotoState223Action,
	{_State442, MulToken}:                                _GotoState221Action,
	{_State442, TraitOpType}:                             _GotoState224Action,
	{_State445, IdentifierToken}:                         _GotoState316Action,
	{_State445, StructToken}:                             _GotoState32Action,
	{_State445, EnumToken}:                               _GotoState90Action,
	{_State445, TraitToken}:                              _GotoState98Action,
	{_State445, FuncToken}:                               _GotoState92Action,
	{_State445, LparenToken}:                             _GotoState94Action,
	{_State445, LbracketToken}:                           _GotoState24Action,
	{_State445, DotToken}:                                _GotoState89Action,
	{_State445, QuestionToken}:                           _GotoState96Action,
	{_State445, ExclaimToken}:                            _GotoState91Action,
	{_State445, EllipsisToken}:                           _GotoState315Action,
	{_State445, TildeTildeToken}:                         _GotoState97Action,
	{_State445, BitNegToken}:                             _GotoState88Action,
	{_State445, BitAndToken}:                             _GotoState87Action,
	{_State445, ParseErrorToken}:                         _GotoState95Action,
	{_State445, InitializableTypeType}:                   _GotoState104Action,
	{_State445, SliceTypeType}:                           _GotoState69Action,
	{_State445, ArrayTypeType}:                           _GotoState41Action,
	{_State445, MapTypeType}:                             _GotoState58Action,
	{_State445, AtomTypeType}:                            _GotoState99Action,
	{_State445, ParseErrorTypeType}:                      _GotoState105Action,
	{_State445, ReturnableTypeType}:                      _GotoState108Action,
	{_State445, PrefixedTypeType}:                        _GotoState107Action,
	{_State445, PrefixTypeOpType}:                        _GotoState106Action,
	{_State445, ValueTypeType}:                           _GotoState320Action,
	{_State445, TraitOpTypeType}:                         _GotoState110Action,
	{_State445, ImplicitStructDefType}:                   _GotoState103Action,
	{_State445, ExplicitStructDefType}:                   _GotoState51Action,
	{_State445, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State445, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State445, TraitDefType}:                            _GotoState109Action,
	{_State445, ParameterDeclType}:                       _GotoState317Action,
	{_State445, ProperParameterDeclsType}:                _GotoState319Action,
	{_State445, ParameterDeclsType}:                      _GotoState459Action,
	{_State445, FuncTypeType}:                            _GotoState101Action,
	{_State448, LbraceToken}:                             _GotoState73Action,
	{_State448, StatementBlockType}:                      _GotoState460Action,
	{_State449, IntegerLiteralToken}:                     _GotoState22Action,
	{_State449, FloatLiteralToken}:                       _GotoState18Action,
	{_State449, RuneLiteralToken}:                        _GotoState30Action,
	{_State449, StringLiteralToken}:                      _GotoState31Action,
	{_State449, IdentifierToken}:                         _GotoState21Action,
	{_State449, TrueToken}:                               _GotoState34Action,
	{_State449, FalseToken}:                              _GotoState17Action,
	{_State449, StructToken}:                             _GotoState32Action,
	{_State449, FuncToken}:                               _GotoState19Action,
	{_State449, VarToken}:                                _GotoState35Action,
	{_State449, LetToken}:                                _GotoState25Action,
	{_State449, NotToken}:                                _GotoState28Action,
	{_State449, LabelDeclToken}:                          _GotoState23Action,
	{_State449, LparenToken}:                             _GotoState26Action,
	{_State449, LbracketToken}:                           _GotoState24Action,
	{_State449, SubToken}:                                _GotoState33Action,
	{_State449, MulToken}:                                _GotoState27Action,
	{_State449, BitNegToken}:                             _GotoState16Action,
	{_State449, BitAndToken}:                             _GotoState15Action,
	{_State449, GreaterToken}:                            _GotoState20Action,
	{_State449, ParseErrorToken}:                         _GotoState29Action,
	{_State449, VarDeclPatternType}:                      _GotoState70Action,
	{_State449, VarOrLetType}:                            _GotoState71Action,
	{_State449, OptionalLabelDeclType}:                   _GotoState85Action,
	{_State449, SequenceExprType}:                        _GotoState418Action,
	{_State449, OptionalSequenceExprType}:                _GotoState461Action,
	{_State449, CallExprType}:                            _GotoState49Action,
	{_State449, AtomExprType}:                            _GotoState42Action,
	{_State449, ParseErrorExprType}:                      _GotoState62Action,
	{_State449, LiteralExprType}:                         _GotoState57Action,
	{_State449, IdentifierExprType}:                      _GotoState52Action,
	{_State449, BlockExprType}:                           _GotoState48Action,
	{_State449, InitializeExprType}:                      _GotoState56Action,
	{_State449, ImplicitStructExprType}:                  _GotoState53Action,
	{_State449, AccessibleExprType}:                      _GotoState37Action,
	{_State449, AccessExprType}:                          _GotoState36Action,
	{_State449, IndexExprType}:                           _GotoState54Action,
	{_State449, PostfixableExprType}:                     _GotoState64Action,
	{_State449, PostfixUnaryExprType}:                    _GotoState63Action,
	{_State449, PrefixableExprType}:                      _GotoState67Action,
	{_State449, PrefixUnaryExprType}:                     _GotoState65Action,
	{_State449, PrefixUnaryOpType}:                       _GotoState66Action,
	{_State449, MulExprType}:                             _GotoState59Action,
	{_State449, BinaryMulExprType}:                       _GotoState46Action,
	{_State449, AddExprType}:                             _GotoState38Action,
	{_State449, BinaryAddExprType}:                       _GotoState43Action,
	{_State449, CmpExprType}:                             _GotoState50Action,
	{_State449, BinaryCmpExprType}:                       _GotoState45Action,
	{_State449, AndExprType}:                             _GotoState39Action,
	{_State449, BinaryAndExprType}:                       _GotoState44Action,
	{_State449, OrExprType}:                              _GotoState61Action,
	{_State449, BinaryOrExprType}:                        _GotoState47Action,
	{_State449, InitializableTypeType}:                   _GotoState55Action,
	{_State449, SliceTypeType}:                           _GotoState69Action,
	{_State449, ArrayTypeType}:                           _GotoState41Action,
	{_State449, MapTypeType}:                             _GotoState58Action,
	{_State449, ExplicitStructDefType}:                   _GotoState51Action,
	{_State449, AnonymousFuncExprType}:                   _GotoState40Action,
	{_State457, LbraceToken}:                             _GotoState73Action,
	{_State457, StatementBlockType}:                      _GotoState462Action,
	{_State458, IdentifierToken}:                         _GotoState198Action,
	{_State458, ParameterDefType}:                        _GotoState200Action,
	{_State458, ProperParameterDefsType}:                 _GotoState202Action,
	{_State458, ParameterDefsType}:                       _GotoState463Action,
	{_State459, RparenToken}:                             _GotoState464Action,
	{_State461, DoToken}:                                 _GotoState465Action,
	{_State463, RparenToken}:                             _GotoState466Action,
	{_State464, IdentifierToken}:                         _GotoState93Action,
	{_State464, StructToken}:                             _GotoState32Action,
	{_State464, EnumToken}:                               _GotoState90Action,
	{_State464, TraitToken}:                              _GotoState98Action,
	{_State464, FuncToken}:                               _GotoState92Action,
	{_State464, LparenToken}:                             _GotoState94Action,
	{_State464, LbracketToken}:                           _GotoState24Action,
	{_State464, DotToken}:                                _GotoState89Action,
	{_State464, QuestionToken}:                           _GotoState96Action,
	{_State464, ExclaimToken}:                            _GotoState91Action,
	{_State464, TildeTildeToken}:                         _GotoState97Action,
	{_State464, BitNegToken}:                             _GotoState88Action,
	{_State464, BitAndToken}:                             _GotoState87Action,
	{_State464, ParseErrorToken}:                         _GotoState95Action,
	{_State464, InitializableTypeType}:                   _GotoState104Action,
	{_State464, SliceTypeType}:                           _GotoState69Action,
	{_State464, ArrayTypeType}:                           _GotoState41Action,
	{_State464, MapTypeType}:                             _GotoState58Action,
	{_State464, AtomTypeType}:                            _GotoState99Action,
	{_State464, ParseErrorTypeType}:                      _GotoState105Action,
	{_State464, ReturnableTypeType}:                      _GotoState387Action,
	{_State464, PrefixedTypeType}:                        _GotoState107Action,
	{_State464, PrefixTypeOpType}:                        _GotoState106Action,
	{_State464, ImplicitStructDefType}:                   _GotoState103Action,
	{_State464, ExplicitStructDefType}:                   _GotoState51Action,
	{_State464, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State464, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State464, TraitDefType}:                            _GotoState109Action,
	{_State464, ReturnTypeType}:                          _GotoState467Action,
	{_State464, FuncTypeType}:                            _GotoState101Action,
	{_State465, LbraceToken}:                             _GotoState73Action,
	{_State465, StatementBlockType}:                      _GotoState468Action,
	{_State466, IdentifierToken}:                         _GotoState93Action,
	{_State466, StructToken}:                             _GotoState32Action,
	{_State466, EnumToken}:                               _GotoState90Action,
	{_State466, TraitToken}:                              _GotoState98Action,
	{_State466, FuncToken}:                               _GotoState92Action,
	{_State466, LparenToken}:                             _GotoState94Action,
	{_State466, LbracketToken}:                           _GotoState24Action,
	{_State466, DotToken}:                                _GotoState89Action,
	{_State466, QuestionToken}:                           _GotoState96Action,
	{_State466, ExclaimToken}:                            _GotoState91Action,
	{_State466, TildeTildeToken}:                         _GotoState97Action,
	{_State466, BitNegToken}:                             _GotoState88Action,
	{_State466, BitAndToken}:                             _GotoState87Action,
	{_State466, ParseErrorToken}:                         _GotoState95Action,
	{_State466, InitializableTypeType}:                   _GotoState104Action,
	{_State466, SliceTypeType}:                           _GotoState69Action,
	{_State466, ArrayTypeType}:                           _GotoState41Action,
	{_State466, MapTypeType}:                             _GotoState58Action,
	{_State466, AtomTypeType}:                            _GotoState99Action,
	{_State466, ParseErrorTypeType}:                      _GotoState105Action,
	{_State466, ReturnableTypeType}:                      _GotoState387Action,
	{_State466, PrefixedTypeType}:                        _GotoState107Action,
	{_State466, PrefixTypeOpType}:                        _GotoState106Action,
	{_State466, ImplicitStructDefType}:                   _GotoState103Action,
	{_State466, ExplicitStructDefType}:                   _GotoState51Action,
	{_State466, ImplicitEnumDefType}:                     _GotoState102Action,
	{_State466, ExplicitEnumDefType}:                     _GotoState100Action,
	{_State466, TraitDefType}:                            _GotoState109Action,
	{_State466, ReturnTypeType}:                          _GotoState469Action,
	{_State466, FuncTypeType}:                            _GotoState101Action,
	{_State469, LbraceToken}:                             _GotoState73Action,
	{_State469, StatementBlockType}:                      _GotoState470Action,
	{_State1, _WildcardMarker}:                           _ReduceNilToDefinitionsAction,
	{_State5, _WildcardMarker}:                           _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State11, _EndMarker}:                               _ReduceToSourceAction,
	{_State12, _WildcardMarker}:                          _ReduceNoSpecToPackageDefAction,
	{_State15, _WildcardMarker}:                          _ReduceBitAndToPrefixUnaryOpAction,
	{_State16, _WildcardMarker}:                          _ReduceBitNegToPrefixUnaryOpAction,
	{_State17, _WildcardMarker}:                          _ReduceFalseToLiteralExprAction,
	{_State18, _WildcardMarker}:                          _ReduceFloatLiteralToLiteralExprAction,
	{_State20, LbraceToken}:                              _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State21, _WildcardMarker}:                          _ReduceToIdentifierExprAction,
	{_State22, _WildcardMarker}:                          _ReduceIntegerLiteralToLiteralExprAction,
	{_State23, _WildcardMarker}:                          _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State25, _WildcardMarker}:                          _ReduceLetToVarOrLetAction,
	{_State26, _WildcardMarker}:                          _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State26, RparenToken}:                              _ReduceNilToArgumentsAction,
	{_State27, _WildcardMarker}:                          _ReduceMulToPrefixUnaryOpAction,
	{_State28, _WildcardMarker}:                          _ReduceNotToPrefixUnaryOpAction,
	{_State29, _WildcardMarker}:                          _ReduceToParseErrorExprAction,
	{_State30, _WildcardMarker}:                          _ReduceRuneLiteralToLiteralExprAction,
	{_State31, _WildcardMarker}:                          _ReduceStringLiteralToLiteralExprAction,
	{_State33, _WildcardMarker}:                          _ReduceSubToPrefixUnaryOpAction,
	{_State34, _WildcardMarker}:                          _ReduceTrueToLiteralExprAction,
	{_State35, _WildcardMarker}:                          _ReduceVarToVarOrLetAction,
	{_State36, _WildcardMarker}:                          _ReduceAccessExprToAccessibleExprAction,
	{_State37, _WildcardMarker}:                          _ReduceAccessibleExprToPostfixableExprAction,
	{_State37, LparenToken}:                              _ReduceNilToOptionalGenericBindingAction,
	{_State38, _WildcardMarker}:                          _ReduceAddExprToCmpExprAction,
	{_State39, _WildcardMarker}:                          _ReduceAndExprToOrExprAction,
	{_State40, _WildcardMarker}:                          _ReduceAnonymousFuncExprToAtomExprAction,
	{_State41, _WildcardMarker}:                          _ReduceArrayTypeToInitializableTypeAction,
	{_State42, _WildcardMarker}:                          _ReduceAtomExprToAccessibleExprAction,
	{_State43, _WildcardMarker}:                          _ReduceBinaryAddExprToAddExprAction,
	{_State44, _WildcardMarker}:                          _ReduceBinaryAndExprToAndExprAction,
	{_State45, _WildcardMarker}:                          _ReduceBinaryCmpExprToCmpExprAction,
	{_State46, _WildcardMarker}:                          _ReduceBinaryMulExprToMulExprAction,
	{_State47, _WildcardMarker}:                          _ReduceBinaryOrExprToOrExprAction,
	{_State48, _WildcardMarker}:                          _ReduceBlockExprToAtomExprAction,
	{_State49, _WildcardMarker}:                          _ReduceCallExprToAccessibleExprAction,
	{_State50, _WildcardMarker}:                          _ReduceCmpExprToAndExprAction,
	{_State51, _WildcardMarker}:                          _ReduceExplicitStructDefToInitializableTypeAction,
	{_State52, _WildcardMarker}:                          _ReduceIdentifierExprToAtomExprAction,
	{_State53, _WildcardMarker}:                          _ReduceImplicitStructExprToAtomExprAction,
	{_State54, _WildcardMarker}:                          _ReduceIndexExprToAccessibleExprAction,
	{_State56, _WildcardMarker}:                          _ReduceInitializeExprToAtomExprAction,
	{_State57, _WildcardMarker}:                          _ReduceLiteralExprToAtomExprAction,
	{_State58, _WildcardMarker}:                          _ReduceMapTypeToInitializableTypeAction,
	{_State59, _WildcardMarker}:                          _ReduceMulExprToAddExprAction,
	{_State61, _WildcardMarker}:                          _ReduceOrExprToSequenceExprAction,
	{_State62, _WildcardMarker}:                          _ReduceParseErrorExprToAtomExprAction,
	{_State63, _WildcardMarker}:                          _ReducePostfixUnaryExprToPostfixableExprAction,
	{_State64, _WildcardMarker}:                          _ReducePostfixableExprToPrefixableExprAction,
	{_State65, _WildcardMarker}:                          _ReducePrefixUnaryExprToPrefixableExprAction,
	{_State66, LbraceToken}:                              _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State67, _WildcardMarker}:                          _ReducePrefixableExprToMulExprAction,
	{_State68, _EndMarker}:                               _ReduceSequenceExprToExpressionAction,
	{_State69, _WildcardMarker}:                          _ReduceSliceTypeToInitializableTypeAction,
	{_State70, _EndMarker}:                               _ReduceVarDeclPatternToSequenceExprAction,
	{_State72, NewlinesToken}:                            _ReduceCommentGroupsToDefinitionAction,
	{_State73, RbraceToken}:                              _ReduceNilToStatementsAction,
	{_State73, _WildcardMarker}:                          _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State75, NewlinesToken}:                            _ReduceNamedFuncDefToDefinitionAction,
	{_State76, NewlinesToken}:                            _ReducePackageDefToDefinitionAction,
	{_State77, NewlinesToken}:                            _ReduceStatementBlockToDefinitionAction,
	{_State78, NewlinesToken}:                            _ReduceTypeDefToDefinitionAction,
	{_State79, NewlinesToken}:                            _ReduceGlobalVarDefToDefinitionAction,
	{_State80, _EndMarker}:                               _ReduceWithSpecToPackageDefAction,
	{_State81, _WildcardMarker}:                          _ReduceNilToOptionalGenericParametersAction,
	{_State82, LparenToken}:                              _ReduceNilToOptionalGenericParametersAction,
	{_State84, RparenToken}:                              _ReduceNilToParameterDefsAction,
	{_State86, _EndMarker}:                               _ReduceAssignVarPatternToSequenceExprAction,
	{_State87, _WildcardMarker}:                          _ReduceBitAndToPrefixTypeOpAction,
	{_State88, _WildcardMarker}:                          _ReduceBitNegToPrefixTypeOpAction,
	{_State89, _WildcardMarker}:                          _ReduceNilToOptionalGenericBindingAction,
	{_State91, _WildcardMarker}:                          _ReduceExclaimToPrefixTypeOpAction,
	{_State93, _WildcardMarker}:                          _ReduceNilToOptionalGenericBindingAction,
	{_State94, RparenToken}:                              _ReduceNilToImplicitFieldDefsAction,
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
	{_State112, ColonToken}:                              _ReduceUnitUnitPairToColonExprAction,
	{_State112, CommaToken}:                              _ReduceUnitUnitPairToColonExprAction,
	{_State112, RbracketToken}:                           _ReduceUnitUnitPairToColonExprAction,
	{_State112, RparenToken}:                             _ReduceUnitUnitPairToColonExprAction,
	{_State112, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State113, _WildcardMarker}:                         _ReduceSkipPatternToArgumentAction,
	{_State114, _WildcardMarker}:                         _ReduceToIdentifierExprAction,
	{_State115, _WildcardMarker}:                         _ReduceArgumentToProperArgumentsAction,
	{_State117, _WildcardMarker}:                         _ReduceColonExprToArgumentAction,
	{_State118, _WildcardMarker}:                         _ReducePositionalToArgumentAction,
	{_State119, RparenToken}:                             _ReduceProperArgumentsToArgumentsAction,
	{_State120, RparenToken}:                             _ReduceNilToExplicitFieldDefsAction,
	{_State121, RbracketToken}:                           _ReduceNilToGenericArgumentsAction,
	{_State123, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
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
	{_State139, RparenToken}:                             _ReduceNilToArgumentsAction,
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
	{_State161, _WildcardMarker}:                         _ReduceAsyncToCallbackOpAction,
	{_State162, _WildcardMarker}:                         _ReduceBreakToJumpTypeAction,
	{_State163, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State164, _WildcardMarker}:                         _ReduceContinueToJumpTypeAction,
	{_State166, _WildcardMarker}:                         _ReduceDeferToCallbackOpAction,
	{_State167, _WildcardMarker}:                         _ReduceToFallthroughStatementAction,
	{_State169, _WildcardMarker}:                         _ReduceReturnToJumpTypeAction,
	{_State171, _WildcardMarker}:                         _ReduceAccessibleExprToPostfixableExprAction,
	{_State171, LparenToken}:                             _ReduceNilToOptionalGenericBindingAction,
	{_State173, _WildcardMarker}:                         _ReduceAssignStatementToSimpleStatementAction,
	{_State174, _WildcardMarker}:                         _ReduceBinaryOpAssignStatementToSimpleStatementAction,
	{_State175, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State176, _WildcardMarker}:                         _ReduceCallbackOpStatementToSimpleStatementAction,
	{_State177, _WildcardMarker}:                         _ReduceExpressionToExpressionsAction,
	{_State178, _WildcardMarker}:                         _ReduceExpressionOrImproperStructStatementToSimpleStatementAction,
	{_State179, _WildcardMarker}:                         _ReduceToExpressionOrImproperStructStatementAction,
	{_State180, _WildcardMarker}:                         _ReduceFallthroughStatementToSimpleStatementAction,
	{_State181, _WildcardMarker}:                         _ReduceImportStatementToStatementAction,
	{_State182, _WildcardMarker}:                         _ReduceJumpStatementToSimpleStatementAction,
	{_State183, NewlinesToken}:                           _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State183, RbraceToken}:                             _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State183, SemicolonToken}:                          _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State183, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State184, RbraceToken}:                             _ReduceProperStatementsToStatementsAction,
	{_State185, AssignToken}:                             _ReduceToAssignPatternAction,
	{_State185, _WildcardMarker}:                         _ReduceSequenceExprToExpressionAction,
	{_State186, _WildcardMarker}:                         _ReduceSimpleStatementToStatementAction,
	{_State187, _WildcardMarker}:                         _ReduceStatementToProperStatementsAction,
	{_State189, _WildcardMarker}:                         _ReduceUnaryOpAssignStatementToSimpleStatementAction,
	{_State190, _WildcardMarker}:                         _ReduceUnsafeStatementToSimpleStatementAction,
	{_State191, _WildcardMarker}:                         _ReduceAddToDefinitionsAction,
	{_State192, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State194, RbracketToken}:                           _ReduceNilToGenericParameterDefsAction,
	{_State196, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State198, _WildcardMarker}:                         _ReduceInferredRefArgToParameterDefAction,
	{_State200, _WildcardMarker}:                         _ReduceParameterDefToProperParameterDefsAction,
	{_State202, RparenToken}:                             _ReduceProperParameterDefsToParameterDefsAction,
	{_State203, _WildcardMarker}:                         _ReduceInferredToAtomTypeAction,
	{_State205, RparenToken}:                             _ReduceNilToParameterDeclsAction,
	{_State207, _WildcardMarker}:                         _ReduceNamedToAtomTypeAction,
	{_State208, _WildcardMarker}:                         _ReduceNilToOptionalGenericBindingAction,
	{_State210, _WildcardMarker}:                         _ReduceFieldDefToProperImplicitFieldDefsAction,
	{_State210, OrToken}:                                 _ReduceFieldDefToEnumValueDefAction,
	{_State213, RparenToken}:                             _ReduceProperImplicitFieldDefsToImplicitFieldDefsAction,
	{_State214, _WildcardMarker}:                         _ReduceUnsafeStatementToFieldDefAction,
	{_State215, _WildcardMarker}:                         _ReduceImplicitToFieldDefAction,
	{_State216, RparenToken}:                             _ReduceNilToTraitPropertiesAction,
	{_State217, _WildcardMarker}:                         _ReduceToPrefixedTypeAction,
	{_State218, _WildcardMarker}:                         _ReduceAddToTraitOpAction,
	{_State221, _WildcardMarker}:                         _ReduceMulToTraitOpAction,
	{_State222, _WildcardMarker}:                         _ReduceToSliceTypeAction,
	{_State223, _WildcardMarker}:                         _ReduceSubToTraitOpAction,
	{_State225, _WildcardMarker}:                         _ReduceUnitExprPairToColonExprAction,
	{_State226, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State227, _WildcardMarker}:                         _ReduceToImplicitStructExprAction,
	{_State228, ColonToken}:                              _ReduceColonExprUnitTupleToColonExprAction,
	{_State228, CommaToken}:                              _ReduceColonExprUnitTupleToColonExprAction,
	{_State228, RbracketToken}:                           _ReduceColonExprUnitTupleToColonExprAction,
	{_State228, RparenToken}:                             _ReduceColonExprUnitTupleToColonExprAction,
	{_State228, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State229, ColonToken}:                              _ReduceExprUnitPairToColonExprAction,
	{_State229, CommaToken}:                              _ReduceExprUnitPairToColonExprAction,
	{_State229, RbracketToken}:                           _ReduceExprUnitPairToColonExprAction,
	{_State229, RparenToken}:                             _ReduceExprUnitPairToColonExprAction,
	{_State229, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State230, _WildcardMarker}:                         _ReduceVarargAssignmentToArgumentAction,
	{_State231, RparenToken}:                             _ReduceImproperToArgumentsAction,
	{_State231, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State233, _WildcardMarker}:                         _ReduceFieldDefToProperExplicitFieldDefsAction,
	{_State234, RparenToken}:                             _ReduceProperExplicitFieldDefsToExplicitFieldDefsAction,
	{_State236, RbracketToken}:                           _ReduceProperGenericArgumentsToGenericArgumentsAction,
	{_State237, _WildcardMarker}:                         _ReduceValueTypeToProperGenericArgumentsAction,
	{_State238, _WildcardMarker}:                         _ReduceToAccessExprAction,
	{_State240, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State240, RparenToken}:                             _ReduceNilToArgumentsAction,
	{_State241, _WildcardMarker}:                         _ReduceToBinaryAddExprAction,
	{_State242, _WildcardMarker}:                         _ReduceToBinaryAndExprAction,
	{_State243, _WildcardMarker}:                         _ReduceToBinaryCmpExprAction,
	{_State245, _WildcardMarker}:                         _ReduceToBinaryMulExprAction,
	{_State246, _WildcardMarker}:                         _ReduceInfiniteToLoopExprAction,
	{_State249, _WildcardMarker}:                         _ReduceToAssignPatternAction,
	{_State249, SemicolonToken}:                          _ReduceSequenceExprToForAssignmentAction,
	{_State250, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State252, LbraceToken}:                             _ReduceSequenceExprToConditionAction,
	{_State254, _WildcardMarker}:                         _ReduceToBinaryOrExprAction,
	{_State255, _WildcardMarker}:                         _ReduceEllipsisToFieldVarPatternAction,
	{_State256, _WildcardMarker}:                         _ReduceIdentifierToVarPatternAction,
	{_State257, _WildcardMarker}:                         _ReduceFieldVarPatternToFieldVarPatternsAction,
	{_State259, _WildcardMarker}:                         _ReducePositionalToFieldVarPatternAction,
	{_State260, _EndMarker}:                              _ReduceToVarDeclPatternAction,
	{_State261, _WildcardMarker}:                         _ReduceValueTypeToOptionalValueTypeAction,
	{_State263, _WildcardMarker}:                         _ReduceVarToVarOrLetAction,
	{_State264, _WildcardMarker}:                         _ReduceAssignPatternToCasePatternAction,
	{_State265, _WildcardMarker}:                         _ReduceCasePatternToCasePatternsAction,
	{_State267, _WildcardMarker}:                         _ReduceToAssignPatternAction,
	{_State268, NewlinesToken}:                           _ReduceNilToOptionalSimpleStatementAction,
	{_State268, RbraceToken}:                             _ReduceNilToOptionalSimpleStatementAction,
	{_State268, SemicolonToken}:                          _ReduceNilToOptionalSimpleStatementAction,
	{_State268, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State270, _WildcardMarker}:                         _ReduceStringLiteralToImportClauseAction,
	{_State271, _WildcardMarker}:                         _ReduceSingleToImportStatementAction,
	{_State273, _WildcardMarker}:                         _ReduceAddAssignToBinaryOpAssignAction,
	{_State274, _WildcardMarker}:                         _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State275, _WildcardMarker}:                         _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State276, _WildcardMarker}:                         _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State277, _WildcardMarker}:                         _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State278, _WildcardMarker}:                         _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State279, _WildcardMarker}:                         _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State280, _WildcardMarker}:                         _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State281, _WildcardMarker}:                         _ReduceDivAssignToBinaryOpAssignAction,
	{_State282, _WildcardMarker}:                         _ReduceModAssignToBinaryOpAssignAction,
	{_State283, _WildcardMarker}:                         _ReduceMulAssignToBinaryOpAssignAction,
	{_State284, _WildcardMarker}:                         _ReduceSubAssignToBinaryOpAssignAction,
	{_State285, _WildcardMarker}:                         _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State286, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State287, _WildcardMarker}:                         _ReduceToUnaryOpAssignStatementAction,
	{_State288, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State289, LparenToken}:                             _ReduceNilToOptionalGenericBindingAction,
	{_State290, NewlinesToken}:                           _ReduceToCallbackOpStatementAction,
	{_State290, RbraceToken}:                             _ReduceToCallbackOpStatementAction,
	{_State290, SemicolonToken}:                          _ReduceToCallbackOpStatementAction,
	{_State290, _WildcardMarker}:                         _ReduceCallExprToAccessibleExprAction,
	{_State291, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State292, NewlinesToken}:                           _ReduceLabeledNoValueToJumpStatementAction,
	{_State292, RbraceToken}:                             _ReduceLabeledNoValueToJumpStatementAction,
	{_State292, SemicolonToken}:                          _ReduceLabeledNoValueToJumpStatementAction,
	{_State292, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State293, _WildcardMarker}:                         _ReduceUnlabeledValuedToJumpStatementAction,
	{_State294, RbraceToken}:                             _ReduceImproperImplicitToStatementsAction,
	{_State294, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State295, RbraceToken}:                             _ReduceImproperExplicitToStatementsAction,
	{_State295, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State296, _EndMarker}:                              _ReduceToStatementBlockAction,
	{_State297, NewlinesToken}:                           _ReduceGlobalVarAssignmentToDefinitionAction,
	{_State298, _WildcardMarker}:                         _ReduceAliasToTypeDefAction,
	{_State299, _WildcardMarker}:                         _ReduceUnconstrainedToGenericParameterDefAction,
	{_State300, _WildcardMarker}:                         _ReduceGenericParameterDefToProperGenericParameterDefsAction,
	{_State302, RbracketToken}:                           _ReduceProperGenericParameterDefsToGenericParameterDefsAction,
	{_State303, _WildcardMarker}:                         _ReduceDefinitionToTypeDefAction,
	{_State304, _EndMarker}:                              _ReduceAliasToNamedFuncDefAction,
	{_State305, RparenToken}:                             _ReduceNilToParameterDefsAction,
	{_State306, _WildcardMarker}:                         _ReduceInferredRefVarargToParameterDefAction,
	{_State307, _WildcardMarker}:                         _ReduceArgToParameterDefAction,
	{_State309, LbraceToken}:                             _ReduceNilToReturnTypeAction,
	{_State310, RparenToken}:                             _ReduceImproperToParameterDefsAction,
	{_State313, _WildcardMarker}:                         _ReduceFieldDefToEnumValueDefAction,
	{_State316, _WildcardMarker}:                         _ReduceNilToOptionalGenericBindingAction,
	{_State317, _WildcardMarker}:                         _ReduceParameterDeclToProperParameterDeclsAction,
	{_State319, RparenToken}:                             _ReduceProperParameterDeclsToParameterDeclsAction,
	{_State320, _WildcardMarker}:                         _ReduceUnamedToParameterDeclAction,
	{_State321, _WildcardMarker}:                         _ReduceNilToOptionalGenericBindingAction,
	{_State322, _WildcardMarker}:                         _ReduceNilToOptionalGenericBindingAction,
	{_State323, _WildcardMarker}:                         _ReduceExplicitToFieldDefAction,
	{_State327, _WildcardMarker}:                         _ReduceToImplicitEnumDefAction,
	{_State328, _WildcardMarker}:                         _ReduceToImplicitStructDefAction,
	{_State329, RparenToken}:                             _ReduceImproperToImplicitFieldDefsAction,
	{_State331, _WildcardMarker}:                         _ReduceFieldDefToTraitPropertyAction,
	{_State332, _WildcardMarker}:                         _ReduceMethodSignatureToTraitPropertyAction,
	{_State333, RparenToken}:                             _ReduceProperTraitPropertiesToTraitPropertiesAction,
	{_State335, _WildcardMarker}:                         _ReduceTraitPropertyToProperTraitPropertiesAction,
	{_State338, _WildcardMarker}:                         _ReduceToTraitOpTypeAction,
	{_State339, _WildcardMarker}:                         _ReduceNamedAssignmentToArgumentAction,
	{_State340, _WildcardMarker}:                         _ReduceColonExprExprTupleToColonExprAction,
	{_State341, _WildcardMarker}:                         _ReduceExprExprPairToColonExprAction,
	{_State342, _WildcardMarker}:                         _ReduceAddToProperArgumentsAction,
	{_State343, _WildcardMarker}:                         _ReduceToExplicitStructDefAction,
	{_State344, RparenToken}:                             _ReduceImproperExplicitToExplicitFieldDefsAction,
	{_State345, RparenToken}:                             _ReduceImproperImplicitToExplicitFieldDefsAction,
	{_State346, _WildcardMarker}:                         _ReduceBindingToOptionalGenericBindingAction,
	{_State347, RbracketToken}:                           _ReduceImproperToGenericArgumentsAction,
	{_State348, _WildcardMarker}:                         _ReduceToIndexExprAction,
	{_State350, _WildcardMarker}:                         _ReduceToInitializeExprAction,
	{_State351, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State352, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State353, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State354, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State354, SemicolonToken}:                          _ReduceNilToOptionalSequenceExprAction,
	{_State357, _WildcardMarker}:                         _ReduceNoElseToIfExprAction,
	{_State358, _EndMarker}:                              _ReduceToSwitchExprAction,
	{_State361, _WildcardMarker}:                         _ReduceToTuplePatternAction,
	{_State364, NewlinesToken}:                           _ReduceNilToOptionalSimpleStatementAction,
	{_State364, RbraceToken}:                             _ReduceNilToOptionalSimpleStatementAction,
	{_State364, SemicolonToken}:                          _ReduceNilToOptionalSimpleStatementAction,
	{_State364, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State365, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State366, _WildcardMarker}:                         _ReduceDefaultBranchStatementToStatementAction,
	{_State367, _WildcardMarker}:                         _ReduceSimpleStatementToOptionalSimpleStatementAction,
	{_State368, _WildcardMarker}:                         _ReduceImportClauseToProperImportClausesAction,
	{_State370, RparenToken}:                             _ReduceProperImportClausesToImportClausesAction,
	{_State373, _WildcardMarker}:                         _ReduceToBinaryOpAssignStatementAction,
	{_State374, _WildcardMarker}:                         _ReduceToAssignStatementAction,
	{_State375, _WildcardMarker}:                         _ReduceAddToExpressionsAction,
	{_State376, _WildcardMarker}:                         _ReduceLabeledValuedToJumpStatementAction,
	{_State377, _WildcardMarker}:                         _ReduceAddImplicitToProperStatementsAction,
	{_State378, _WildcardMarker}:                         _ReduceAddExplicitToProperStatementsAction,
	{_State379, _WildcardMarker}:                         _ReduceConstrainedToGenericParameterDefAction,
	{_State380, _WildcardMarker}:                         _ReduceGenericToOptionalGenericParametersAction,
	{_State381, RbracketToken}:                           _ReduceImproperToGenericParameterDefsAction,
	{_State384, _WildcardMarker}:                         _ReduceVarargToParameterDefAction,
	{_State385, LparenToken}:                             _ReduceNilToOptionalGenericParametersAction,
	{_State387, _WildcardMarker}:                         _ReduceReturnableTypeToReturnTypeAction,
	{_State388, _WildcardMarker}:                         _ReduceAddToProperParameterDefsAction,
	{_State391, _WildcardMarker}:                         _ReduceToExplicitEnumDefAction,
	{_State394, _WildcardMarker}:                         _ReduceUnnamedVarargToParameterDeclAction,
	{_State396, _WildcardMarker}:                         _ReduceArgToParameterDeclAction,
	{_State397, _WildcardMarker}:                         _ReduceNilToReturnTypeAction,
	{_State398, RparenToken}:                             _ReduceImproperToParameterDeclsAction,
	{_State399, _WildcardMarker}:                         _ReduceExternNamedToAtomTypeAction,
	{_State400, _WildcardMarker}:                         _ReducePairToImplicitEnumValueDefsAction,
	{_State401, _WildcardMarker}:                         _ReduceDefaultToEnumValueDefAction,
	{_State402, _WildcardMarker}:                         _ReduceAddToImplicitEnumValueDefsAction,
	{_State403, _WildcardMarker}:                         _ReduceAddToProperImplicitFieldDefsAction,
	{_State405, RparenToken}:                             _ReduceImproperExplicitToTraitPropertiesAction,
	{_State406, RparenToken}:                             _ReduceImproperImplicitToTraitPropertiesAction,
	{_State407, _WildcardMarker}:                         _ReduceToTraitDefAction,
	{_State408, _WildcardMarker}:                         _ReduceToMapTypeAction,
	{_State409, _WildcardMarker}:                         _ReduceToArrayTypeAction,
	{_State410, _WildcardMarker}:                         _ReduceAddExplicitToProperExplicitFieldDefsAction,
	{_State411, _WildcardMarker}:                         _ReduceAddImplicitToProperExplicitFieldDefsAction,
	{_State412, _WildcardMarker}:                         _ReduceAddToProperGenericArgumentsAction,
	{_State413, _WildcardMarker}:                         _ReduceToCallExprAction,
	{_State414, _EndMarker}:                              _ReduceDoWhileToLoopExprAction,
	{_State415, SemicolonToken}:                          _ReduceAssignToForAssignmentAction,
	{_State418, DoToken}:                                 _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State419, _EndMarker}:                              _ReduceWhileToLoopExprAction,
	{_State420, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State422, _WildcardMarker}:                         _ReduceNamedToFieldVarPatternAction,
	{_State423, _WildcardMarker}:                         _ReduceAddToFieldVarPatternsAction,
	{_State424, _WildcardMarker}:                         _ReduceEnumMatchPatternToCasePatternAction,
	{_State426, _WildcardMarker}:                         _ReduceCaseBranchStatementToStatementAction,
	{_State427, _WildcardMarker}:                         _ReduceMultipleToCasePatternsAction,
	{_State428, _WildcardMarker}:                         _ReduceMultipleToImportStatementAction,
	{_State429, RparenToken}:                             _ReduceExplicitToImportClausesAction,
	{_State430, RparenToken}:                             _ReduceImplicitToImportClausesAction,
	{_State431, _WildcardMarker}:                         _ReduceAliasToImportClauseAction,
	{_State433, _WildcardMarker}:                         _ReduceAddToProperGenericParameterDefsAction,
	{_State434, _WildcardMarker}:                         _ReduceConstrainedDefToTypeDefAction,
	{_State435, LbraceToken}:                             _ReduceNilToReturnTypeAction,
	{_State437, _WildcardMarker}:                         _ReduceToAnonymousFuncExprAction,
	{_State438, RparenToken}:                             _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State439, _WildcardMarker}:                         _ReducePairToImplicitEnumValueDefsAction,
	{_State439, RparenToken}:                             _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State440, RparenToken}:                             _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State441, _WildcardMarker}:                         _ReduceAddToImplicitEnumValueDefsAction,
	{_State441, RparenToken}:                             _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State442, _WildcardMarker}:                         _ReduceVarargToParameterDeclAction,
	{_State443, _WildcardMarker}:                         _ReduceToFuncTypeAction,
	{_State444, _WildcardMarker}:                         _ReduceAddToProperParameterDeclsAction,
	{_State445, RparenToken}:                             _ReduceNilToParameterDeclsAction,
	{_State446, _WildcardMarker}:                         _ReduceExplicitToProperTraitPropertiesAction,
	{_State447, _WildcardMarker}:                         _ReduceImplicitToProperTraitPropertiesAction,
	{_State449, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State449, DoToken}:                                 _ReduceNilToOptionalSequenceExprAction,
	{_State450, LbraceToken}:                             _ReduceCaseToConditionAction,
	{_State451, _EndMarker}:                              _ReduceMultiIfElseToIfExprAction,
	{_State452, _EndMarker}:                              _ReduceIfElseToIfExprAction,
	{_State453, _WildcardMarker}:                         _ReduceEnumVarDeclPatternToCasePatternAction,
	{_State454, _WildcardMarker}:                         _ReduceAddExplicitToProperImportClausesAction,
	{_State455, _WildcardMarker}:                         _ReduceAddImplicitToProperImportClausesAction,
	{_State456, _WildcardMarker}:                         _ReduceToUnsafeStatementAction,
	{_State458, RparenToken}:                             _ReduceNilToParameterDefsAction,
	{_State460, _EndMarker}:                              _ReduceIteratorToLoopExprAction,
	{_State462, _EndMarker}:                              _ReduceFuncDefToNamedFuncDefAction,
	{_State464, _WildcardMarker}:                         _ReduceNilToReturnTypeAction,
	{_State466, LbraceToken}:                             _ReduceNilToReturnTypeAction,
	{_State467, _WildcardMarker}:                         _ReduceToMethodSignatureAction,
	{_State468, _EndMarker}:                              _ReduceForToLoopExprAction,
	{_State470, _EndMarker}:                              _ReduceMethodDefToNamedFuncDefAction,
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.source
    Reduce:
      * -> [definitions]
    Goto:
      source -> State 6
      definitions -> State 11

  State 2:
    Kernel Items:
      #accept: ^.package_def
    Reduce:
      (nil)
    Goto:
      PACKAGE -> State 12
      package_def -> State 7

  State 3:
    Kernel Items:
      #accept: ^.type_def
    Reduce:
      (nil)
    Goto:
      TYPE -> State 13
      type_def -> State 8

  State 4:
    Kernel Items:
      #accept: ^.named_func_def
    Reduce:
      (nil)
    Goto:
      FUNC -> State 14
      named_func_def -> State 9

  State 5:
    Kernel Items:
      #accept: ^.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 10
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

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
      source: definitions., $
      definitions: definitions.definition NEWLINES
    Reduce:
      $ -> [source]
    Goto:
      COMMENT_GROUPS -> State 72
      PACKAGE -> State 12
      TYPE -> State 13
      FUNC -> State 14
      VAR -> State 35
      LET -> State 25
      LBRACE -> State 73
      definition -> State 74
      statement_block -> State 77
      var_decl_pattern -> State 79
      var_or_let -> State 71
      type_def -> State 78
      named_func_def -> State 75
      package_def -> State 76

  State 12:
    Kernel Items:
      package_def: PACKAGE., *
      package_def: PACKAGE.statement_block
    Reduce:
      * -> [package_def]
    Goto:
      LBRACE -> State 73
      statement_block -> State 80

  State 13:
    Kernel Items:
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE.IDENTIFIER ASSIGN value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 81

  State 14:
    Kernel Items:
      named_func_def: FUNC.IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.IDENTIFIER ASSIGN expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 82
      LPAREN -> State 83

  State 15:
    Kernel Items:
      prefix_unary_op: BIT_AND., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 16:
    Kernel Items:
      prefix_unary_op: BIT_NEG., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 17:
    Kernel Items:
      literal_expr: FALSE., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 18:
    Kernel Items:
      literal_expr: FLOAT_LITERAL., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 19:
    Kernel Items:
      anonymous_func_expr: FUNC.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 84

  State 20:
    Kernel Items:
      sequence_expr: GREATER.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      optional_label_decl -> State 85
      sequence_expr -> State 86
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 21:
    Kernel Items:
      identifier_expr: IDENTIFIER., *
    Reduce:
      * -> [identifier_expr]
    Goto:
      (nil)

  State 22:
    Kernel Items:
      literal_expr: INTEGER_LITERAL., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      optional_label_decl: LABEL_DECL., *
    Reduce:
      * -> [optional_label_decl]
    Goto:
      (nil)

  State 24:
    Kernel Items:
      slice_type: LBRACKET.value_type RBRACKET
      array_type: LBRACKET.value_type COMMA INTEGER_LITERAL RBRACKET
      map_type: LBRACKET.value_type COLON value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 111
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 25:
    Kernel Items:
      var_or_let: LET., *
    Reduce:
      * -> [var_or_let]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      implicit_struct_expr: LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [arguments]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 114
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      COLON -> State 112
      ELLIPSIS -> State 113
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 118
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      proper_arguments -> State 119
      arguments -> State 116
      argument -> State 115
      colon_expr -> State 117
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

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
      parse_error_expr: PARSE_ERROR., *
    Reduce:
      * -> [parse_error_expr]
    Goto:
      (nil)

  State 30:
    Kernel Items:
      literal_expr: RUNE_LITERAL., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 31:
    Kernel Items:
      literal_expr: STRING_LITERAL., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 32:
    Kernel Items:
      explicit_struct_def: STRUCT.LPAREN explicit_field_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 120

  State 33:
    Kernel Items:
      prefix_unary_op: SUB., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 34:
    Kernel Items:
      literal_expr: TRUE., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 35:
    Kernel Items:
      var_or_let: VAR., *
    Reduce:
      * -> [var_or_let]
    Goto:
      (nil)

  State 36:
    Kernel Items:
      accessible_expr: access_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 37:
    Kernel Items:
      call_expr: accessible_expr.optional_generic_binding LPAREN arguments RPAREN
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

  State 38:
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

  State 39:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 131

  State 40:
    Kernel Items:
      atom_expr: anonymous_func_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 41:
    Kernel Items:
      initializable_type: array_type., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 42:
    Kernel Items:
      accessible_expr: atom_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 43:
    Kernel Items:
      add_expr: binary_add_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      (nil)

  State 44:
    Kernel Items:
      and_expr: binary_and_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      (nil)

  State 45:
    Kernel Items:
      cmp_expr: binary_cmp_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      (nil)

  State 46:
    Kernel Items:
      mul_expr: binary_mul_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 47:
    Kernel Items:
      or_expr: binary_or_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      atom_expr: block_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 49:
    Kernel Items:
      accessible_expr: call_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 50:
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

  State 51:
    Kernel Items:
      initializable_type: explicit_struct_def., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 52:
    Kernel Items:
      atom_expr: identifier_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 53:
    Kernel Items:
      atom_expr: implicit_struct_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 54:
    Kernel Items:
      accessible_expr: index_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 55:
    Kernel Items:
      initialize_expr: initializable_type.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 139

  State 56:
    Kernel Items:
      atom_expr: initialize_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 57:
    Kernel Items:
      atom_expr: literal_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 58:
    Kernel Items:
      initializable_type: map_type., *
    Reduce:
      * -> [initializable_type]
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
      LBRACE -> State 73
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
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      PARSE_ERROR -> State 29
      optional_label_decl -> State 85
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 156
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

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
      initializable_type: slice_type., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 70:
    Kernel Items:
      sequence_expr: var_decl_pattern., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 71:
    Kernel Items:
      var_decl_pattern: var_or_let.var_pattern optional_value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      LPAREN -> State 158
      var_pattern -> State 160
      tuple_pattern -> State 159

  State 72:
    Kernel Items:
      definition: COMMENT_GROUPS., NEWLINES
    Reduce:
      NEWLINES -> [definition]
    Goto:
      (nil)

  State 73:
    Kernel Items:
      statement_block: LBRACE.statements RBRACE
    Reduce:
      * -> [optional_label_decl]
      RBRACE -> [statements]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      CASE -> State 163
      DEFAULT -> State 165
      RETURN -> State 169
      BREAK -> State 162
      CONTINUE -> State 164
      FALLTHROUGH -> State 167
      IMPORT -> State 168
      UNSAFE -> State 170
      STRUCT -> State 32
      FUNC -> State 19
      ASYNC -> State 161
      DEFER -> State 166
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      proper_statements -> State 184
      statements -> State 188
      simple_statement -> State 186
      statement -> State 187
      expression_or_improper_struct_statement -> State 178
      expressions -> State 179
      callback_op -> State 175
      callback_op_statement -> State 176
      unsafe_statement -> State 190
      jump_statement -> State 182
      jump_type -> State 183
      fallthrough_statement -> State 180
      assign_statement -> State 173
      unary_op_assign_statement -> State 189
      binary_op_assign_statement -> State 174
      import_statement -> State 181
      var_decl_pattern -> State 70
      var_or_let -> State 71
      assign_pattern -> State 172
      expression -> State 177
      optional_label_decl -> State 60
      sequence_expr -> State 185
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 171
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 74:
    Kernel Items:
      definitions: definitions definition.NEWLINES
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 191

  State 75:
    Kernel Items:
      definition: named_func_def., NEWLINES
    Reduce:
      NEWLINES -> [definition]
    Goto:
      (nil)

  State 76:
    Kernel Items:
      definition: package_def., NEWLINES
    Reduce:
      NEWLINES -> [definition]
    Goto:
      (nil)

  State 77:
    Kernel Items:
      definition: statement_block., NEWLINES
    Reduce:
      NEWLINES -> [definition]
    Goto:
      (nil)

  State 78:
    Kernel Items:
      definition: type_def., NEWLINES
    Reduce:
      NEWLINES -> [definition]
    Goto:
      (nil)

  State 79:
    Kernel Items:
      definition: var_decl_pattern., NEWLINES
      definition: var_decl_pattern.ASSIGN expression
    Reduce:
      NEWLINES -> [definition]
    Goto:
      ASSIGN -> State 192

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
      DOLLAR_LBRACKET -> State 194
      ASSIGN -> State 193
      optional_generic_parameters -> State 195

  State 82:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER.optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC IDENTIFIER.ASSIGN expression
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 194
      ASSIGN -> State 196
      optional_generic_parameters -> State 197

  State 83:
    Kernel Items:
      named_func_def: FUNC LPAREN.parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 198
      parameter_def -> State 199

  State 84:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 198
      parameter_def -> State 200
      proper_parameter_defs -> State 202
      parameter_defs -> State 201

  State 85:
    Kernel Items:
      block_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 73
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
      optional_generic_binding -> State 203

  State 90:
    Kernel Items:
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 204

  State 91:
    Kernel Items:
      prefix_type_op: EXCLAIM., *
    Reduce:
      * -> [prefix_type_op]
    Goto:
      (nil)

  State 92:
    Kernel Items:
      func_type: FUNC.LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 205

  State 93:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOT -> State 206
      DOLLAR_LBRACKET -> State 121
      optional_generic_binding -> State 207

  State 94:
    Kernel Items:
      implicit_struct_def: LPAREN.implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      IDENTIFIER -> State 208
      UNSAFE -> State 170
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 214
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 215
      trait_op_type -> State 110
      field_def -> State 210
      proper_implicit_field_defs -> State 213
      implicit_field_defs -> State 212
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      enum_value_def -> State 209
      implicit_enum_value_defs -> State 211
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
      trait_def: TRAIT.LPAREN trait_properties RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 216

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
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 217
      prefixed_type -> State 107
      prefix_type_op -> State 106
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
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
      slice_type: LBRACKET value_type.RBRACKET
      array_type: LBRACKET value_type.COMMA INTEGER_LITERAL RBRACKET
      map_type: LBRACKET value_type.COLON value_type RBRACKET
      trait_op_type: value_type.trait_op returnable_type
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 222
      COMMA -> State 220
      COLON -> State 219
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 112:
    Kernel Items:
      colon_expr: COLON., COLON
      colon_expr: COLON., COMMA
      colon_expr: COLON., RBRACKET
      colon_expr: COLON., RPAREN
      colon_expr: COLON.expression
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [colon_expr]
      RBRACKET -> [colon_expr]
      COMMA -> [colon_expr]
      COLON -> [colon_expr]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 225
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 113:
    Kernel Items:
      argument: ELLIPSIS., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 114:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expression
      identifier_expr: IDENTIFIER., *
    Reduce:
      * -> [identifier_expr]
    Goto:
      ASSIGN -> State 226

  State 115:
    Kernel Items:
      proper_arguments: argument., *
    Reduce:
      * -> [proper_arguments]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 227

  State 117:
    Kernel Items:
      argument: colon_expr., *
      colon_expr: colon_expr.COLON
      colon_expr: colon_expr.COLON expression
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 228

  State 118:
    Kernel Items:
      argument: expression., *
      argument: expression.ELLIPSIS
      colon_expr: expression.COLON
      colon_expr: expression.COLON expression
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 229
      ELLIPSIS -> State 230

  State 119:
    Kernel Items:
      proper_arguments: proper_arguments.COMMA argument
      arguments: proper_arguments., RPAREN
      arguments: proper_arguments.COMMA
    Reduce:
      RPAREN -> [arguments]
    Goto:
      COMMA -> State 231

  State 120:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN.explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 208
      UNSAFE -> State 170
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 214
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 215
      trait_op_type -> State 110
      field_def -> State 233
      implicit_struct_def -> State 103
      proper_explicit_field_defs -> State 234
      explicit_field_defs -> State 232
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 121:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [generic_arguments]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      proper_generic_arguments -> State 236
      generic_arguments -> State 235
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 237
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
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
      IDENTIFIER -> State 238

  State 123:
    Kernel Items:
      index_expr: accessible_expr LBRACKET.argument RBRACKET
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 114
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      COLON -> State 112
      ELLIPSIS -> State 113
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 118
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      argument -> State 239
      colon_expr -> State 117
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 124:
    Kernel Items:
      postfix_unary_expr: accessible_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      call_expr: accessible_expr optional_generic_binding.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 240

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
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      PARSE_ERROR -> State 29
      optional_label_decl -> State 85
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 241
      binary_mul_expr -> State 46
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 131:
    Kernel Items:
      binary_and_expr: and_expr AND.cmp_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      PARSE_ERROR -> State 29
      optional_label_decl -> State 85
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 242
      binary_cmp_expr -> State 45
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

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
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      PARSE_ERROR -> State 29
      optional_label_decl -> State 85
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 243
      binary_add_expr -> State 43
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 139:
    Kernel Items:
      initialize_expr: initializable_type LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [arguments]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 114
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      COLON -> State 112
      ELLIPSIS -> State 113
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 118
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      proper_arguments -> State 119
      arguments -> State 244
      argument -> State 115
      colon_expr -> State 117
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

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
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      PARSE_ERROR -> State 29
      optional_label_decl -> State 85
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 245
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 147:
    Kernel Items:
      loop_expr: DO.statement_block
      loop_expr: DO.statement_block FOR sequence_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 73
      statement_block -> State 246

  State 148:
    Kernel Items:
      loop_expr: FOR.sequence_expr DO statement_block
      loop_expr: FOR.assign_pattern IN sequence_expr DO statement_block
      loop_expr: FOR.for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      assign_pattern -> State 247
      optional_label_decl -> State 85
      sequence_expr -> State 249
      for_assignment -> State 248
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 149:
    Kernel Items:
      if_expr: IF.condition statement_block
      if_expr: IF.condition statement_block ELSE statement_block
      if_expr: IF.condition statement_block ELSE if_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      CASE -> State 250
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      optional_label_decl -> State 85
      sequence_expr -> State 252
      condition -> State 251
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 150:
    Kernel Items:
      switch_expr: SWITCH.sequence_expr statement_block
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      optional_label_decl -> State 85
      sequence_expr -> State 253
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

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
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      PARSE_ERROR -> State 29
      optional_label_decl -> State 85
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 254
      binary_and_expr -> State 44
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

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
      IDENTIFIER -> State 256
      LPAREN -> State 158
      ELLIPSIS -> State 255
      var_pattern -> State 259
      tuple_pattern -> State 159
      field_var_patterns -> State 258
      field_var_pattern -> State 257

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
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      optional_value_type -> State 260
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 261
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 161:
    Kernel Items:
      callback_op: ASYNC., *
    Reduce:
      * -> [callback_op]
    Goto:
      (nil)

  State 162:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 163:
    Kernel Items:
      statement: CASE.case_patterns COLON optional_simple_statement
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 263
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      DOT -> State 262
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      case_patterns -> State 266
      var_decl_pattern -> State 70
      var_or_let -> State 71
      assign_pattern -> State 264
      case_pattern -> State 265
      optional_label_decl -> State 85
      sequence_expr -> State 267
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 164:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 165:
    Kernel Items:
      statement: DEFAULT.COLON optional_simple_statement
    Reduce:
      (nil)
    Goto:
      COLON -> State 268

  State 166:
    Kernel Items:
      callback_op: DEFER., *
    Reduce:
      * -> [callback_op]
    Goto:
      (nil)

  State 167:
    Kernel Items:
      fallthrough_statement: FALLTHROUGH., *
    Reduce:
      * -> [fallthrough_statement]
    Goto:
      (nil)

  State 168:
    Kernel Items:
      import_statement: IMPORT.import_clause
      import_statement: IMPORT.LPAREN import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 270
      LPAREN -> State 269
      import_clause -> State 271

  State 169:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 170:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 272

  State 171:
    Kernel Items:
      unary_op_assign_statement: accessible_expr.unary_op_assign
      binary_op_assign_statement: accessible_expr.binary_op_assign expression
      call_expr: accessible_expr.optional_generic_binding LPAREN arguments RPAREN
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
      ADD_ASSIGN -> State 273
      SUB_ASSIGN -> State 284
      MUL_ASSIGN -> State 283
      DIV_ASSIGN -> State 281
      MOD_ASSIGN -> State 282
      ADD_ONE_ASSIGN -> State 274
      SUB_ONE_ASSIGN -> State 285
      BIT_NEG_ASSIGN -> State 277
      BIT_AND_ASSIGN -> State 275
      BIT_OR_ASSIGN -> State 278
      BIT_XOR_ASSIGN -> State 280
      BIT_LSHIFT_ASSIGN -> State 276
      BIT_RSHIFT_ASSIGN -> State 279
      unary_op_assign -> State 287
      binary_op_assign -> State 286
      optional_generic_binding -> State 125

  State 172:
    Kernel Items:
      assign_statement: assign_pattern.ASSIGN expression
    Reduce:
      (nil)
    Goto:
      ASSIGN -> State 288

  State 173:
    Kernel Items:
      simple_statement: assign_statement., *
    Reduce:
      * -> [simple_statement]
    Goto:
      (nil)

  State 174:
    Kernel Items:
      simple_statement: binary_op_assign_statement., *
    Reduce:
      * -> [simple_statement]
    Goto:
      (nil)

  State 175:
    Kernel Items:
      callback_op_statement: callback_op.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      PARSE_ERROR -> State 29
      optional_label_decl -> State 85
      call_expr -> State 290
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 289
      access_expr -> State 36
      index_expr -> State 54
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 176:
    Kernel Items:
      simple_statement: callback_op_statement., *
    Reduce:
      * -> [simple_statement]
    Goto:
      (nil)

  State 177:
    Kernel Items:
      expressions: expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 178:
    Kernel Items:
      simple_statement: expression_or_improper_struct_statement., *
    Reduce:
      * -> [simple_statement]
    Goto:
      (nil)

  State 179:
    Kernel Items:
      expression_or_improper_struct_statement: expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [expression_or_improper_struct_statement]
    Goto:
      COMMA -> State 291

  State 180:
    Kernel Items:
      simple_statement: fallthrough_statement., *
    Reduce:
      * -> [simple_statement]
    Goto:
      (nil)

  State 181:
    Kernel Items:
      statement: import_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 182:
    Kernel Items:
      simple_statement: jump_statement., *
    Reduce:
      * -> [simple_statement]
    Goto:
      (nil)

  State 183:
    Kernel Items:
      jump_statement: jump_type., NEWLINES
      jump_statement: jump_type., RBRACE
      jump_statement: jump_type., SEMICOLON
      jump_statement: jump_type.expressions
      jump_statement: jump_type.JUMP_LABEL
      jump_statement: jump_type.JUMP_LABEL expressions
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [jump_statement]
      RBRACE -> [jump_statement]
      SEMICOLON -> [jump_statement]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      JUMP_LABEL -> State 292
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      expressions -> State 293
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 177
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 184:
    Kernel Items:
      proper_statements: proper_statements.NEWLINES statement
      proper_statements: proper_statements.SEMICOLON statement
      statements: proper_statements., RBRACE
      statements: proper_statements.NEWLINES
      statements: proper_statements.SEMICOLON
    Reduce:
      RBRACE -> [statements]
    Goto:
      NEWLINES -> State 294
      SEMICOLON -> State 295

  State 185:
    Kernel Items:
      assign_pattern: sequence_expr., ASSIGN
      expression: sequence_expr., *
    Reduce:
      * -> [expression]
      ASSIGN -> [assign_pattern]
    Goto:
      (nil)

  State 186:
    Kernel Items:
      statement: simple_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 187:
    Kernel Items:
      proper_statements: statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 188:
    Kernel Items:
      statement_block: LBRACE statements.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 296

  State 189:
    Kernel Items:
      simple_statement: unary_op_assign_statement., *
    Reduce:
      * -> [simple_statement]
    Goto:
      (nil)

  State 190:
    Kernel Items:
      simple_statement: unsafe_statement., *
    Reduce:
      * -> [simple_statement]
    Goto:
      (nil)

  State 191:
    Kernel Items:
      definitions: definitions definition NEWLINES., *
    Reduce:
      * -> [definitions]
    Goto:
      (nil)

  State 192:
    Kernel Items:
      definition: var_decl_pattern ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 297
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 193:
    Kernel Items:
      type_def: TYPE IDENTIFIER ASSIGN.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 298
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 194:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 299
      generic_parameter_def -> State 300
      proper_generic_parameter_defs -> State 302
      generic_parameter_defs -> State 301

  State 195:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 303
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 196:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 304
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 197:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 305

  State 198:
    Kernel Items:
      parameter_def: IDENTIFIER., *
      parameter_def: IDENTIFIER.ELLIPSIS
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.ELLIPSIS value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      ELLIPSIS -> State 306
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 307
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 199:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def.RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 308

  State 200:
    Kernel Items:
      proper_parameter_defs: parameter_def., *
    Reduce:
      * -> [proper_parameter_defs]
    Goto:
      (nil)

  State 201:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 309

  State 202:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs.COMMA parameter_def
      parameter_defs: proper_parameter_defs., RPAREN
      parameter_defs: proper_parameter_defs.COMMA
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      COMMA -> State 310

  State 203:
    Kernel Items:
      atom_type: DOT optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 204:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 208
      UNSAFE -> State 170
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 214
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 215
      trait_op_type -> State 110
      field_def -> State 313
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      enum_value_def -> State 311
      implicit_enum_value_defs -> State 314
      implicit_enum_def -> State 102
      explicit_enum_value_defs -> State 312
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 205:
    Kernel Items:
      func_type: FUNC LPAREN.parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 316
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      ELLIPSIS -> State 315
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 320
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      parameter_decl -> State 317
      proper_parameter_decls -> State 319
      parameter_decls -> State 318
      func_type -> State 101

  State 206:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 321

  State 207:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 208:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 322
      QUESTION -> State 96
      EXCLAIM -> State 91
      DOLLAR_LBRACKET -> State 121
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      optional_generic_binding -> State 207
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 323
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 209:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 324

  State 210:
    Kernel Items:
      proper_implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [proper_implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 325

  State 211:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      OR -> State 326
      RPAREN -> State 327

  State 212:
    Kernel Items:
      implicit_struct_def: LPAREN implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 328

  State 213:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs.COMMA field_def
      implicit_field_defs: proper_implicit_field_defs., RPAREN
      implicit_field_defs: proper_implicit_field_defs.COMMA
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      COMMA -> State 329

  State 214:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 215:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 216:
    Kernel Items:
      trait_def: TRAIT LPAREN.trait_properties RPAREN
    Reduce:
      RPAREN -> [trait_properties]
    Goto:
      IDENTIFIER -> State 208
      UNSAFE -> State 170
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 330
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 214
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 215
      trait_op_type -> State 110
      field_def -> State 331
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_property -> State 335
      proper_trait_properties -> State 333
      trait_properties -> State 334
      trait_def -> State 109
      func_type -> State 101
      method_signature -> State 332

  State 217:
    Kernel Items:
      prefixed_type: prefix_type_op returnable_type., *
    Reduce:
      * -> [prefixed_type]
    Goto:
      (nil)

  State 218:
    Kernel Items:
      trait_op: ADD., *
    Reduce:
      * -> [trait_op]
    Goto:
      (nil)

  State 219:
    Kernel Items:
      map_type: LBRACKET value_type COLON.value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 336
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 220:
    Kernel Items:
      array_type: LBRACKET value_type COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 337

  State 221:
    Kernel Items:
      trait_op: MUL., *
    Reduce:
      * -> [trait_op]
    Goto:
      (nil)

  State 222:
    Kernel Items:
      slice_type: LBRACKET value_type RBRACKET., *
    Reduce:
      * -> [slice_type]
    Goto:
      (nil)

  State 223:
    Kernel Items:
      trait_op: SUB., *
    Reduce:
      * -> [trait_op]
    Goto:
      (nil)

  State 224:
    Kernel Items:
      trait_op_type: value_type trait_op.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 338
      prefixed_type -> State 107
      prefix_type_op -> State 106
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 225:
    Kernel Items:
      colon_expr: COLON expression., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 226:
    Kernel Items:
      argument: IDENTIFIER ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 339
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 227:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [implicit_struct_expr]
    Goto:
      (nil)

  State 228:
    Kernel Items:
      colon_expr: colon_expr COLON., COLON
      colon_expr: colon_expr COLON., COMMA
      colon_expr: colon_expr COLON., RBRACKET
      colon_expr: colon_expr COLON., RPAREN
      colon_expr: colon_expr COLON.expression
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [colon_expr]
      RBRACKET -> [colon_expr]
      COMMA -> [colon_expr]
      COLON -> [colon_expr]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 340
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 229:
    Kernel Items:
      colon_expr: expression COLON., COLON
      colon_expr: expression COLON., COMMA
      colon_expr: expression COLON., RBRACKET
      colon_expr: expression COLON., RPAREN
      colon_expr: expression COLON.expression
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [colon_expr]
      RBRACKET -> [colon_expr]
      COMMA -> [colon_expr]
      COLON -> [colon_expr]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 341
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 230:
    Kernel Items:
      argument: expression ELLIPSIS., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 231:
    Kernel Items:
      proper_arguments: proper_arguments COMMA.argument
      arguments: proper_arguments COMMA., RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [arguments]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 114
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      COLON -> State 112
      ELLIPSIS -> State 113
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 118
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      argument -> State 342
      colon_expr -> State 117
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 232:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 343

  State 233:
    Kernel Items:
      proper_explicit_field_defs: field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 234:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs.NEWLINES field_def
      proper_explicit_field_defs: proper_explicit_field_defs.COMMA field_def
      explicit_field_defs: proper_explicit_field_defs., RPAREN
      explicit_field_defs: proper_explicit_field_defs.NEWLINES
      explicit_field_defs: proper_explicit_field_defs.COMMA
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      NEWLINES -> State 345
      COMMA -> State 344

  State 235:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 346

  State 236:
    Kernel Items:
      proper_generic_arguments: proper_generic_arguments.COMMA value_type
      generic_arguments: proper_generic_arguments., RBRACKET
      generic_arguments: proper_generic_arguments.COMMA
    Reduce:
      RBRACKET -> [generic_arguments]
    Goto:
      COMMA -> State 347

  State 237:
    Kernel Items:
      proper_generic_arguments: value_type., *
      trait_op_type: value_type.trait_op returnable_type
    Reduce:
      * -> [proper_generic_arguments]
    Goto:
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 238:
    Kernel Items:
      access_expr: accessible_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 239:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 348

  State 240:
    Kernel Items:
      call_expr: accessible_expr optional_generic_binding LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [arguments]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 114
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      COLON -> State 112
      ELLIPSIS -> State 113
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 118
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      proper_arguments -> State 119
      arguments -> State 349
      argument -> State 115
      colon_expr -> State 117
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 241:
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

  State 242:
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

  State 243:
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

  State 244:
    Kernel Items:
      initialize_expr: initializable_type LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 350

  State 245:
    Kernel Items:
      binary_mul_expr: mul_expr mul_op prefixable_expr., *
    Reduce:
      * -> [binary_mul_expr]
    Goto:
      (nil)

  State 246:
    Kernel Items:
      loop_expr: DO statement_block., *
      loop_expr: DO statement_block.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 351

  State 247:
    Kernel Items:
      loop_expr: FOR assign_pattern.IN sequence_expr DO statement_block
      for_assignment: assign_pattern.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      IN -> State 353
      ASSIGN -> State 352

  State 248:
    Kernel Items:
      loop_expr: FOR for_assignment.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 354

  State 249:
    Kernel Items:
      assign_pattern: sequence_expr., *
      loop_expr: FOR sequence_expr.DO statement_block
      for_assignment: sequence_expr., SEMICOLON
    Reduce:
      * -> [assign_pattern]
      SEMICOLON -> [for_assignment]
    Goto:
      DO -> State 355

  State 250:
    Kernel Items:
      condition: CASE.case_patterns ASSIGN sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 263
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      DOT -> State 262
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      case_patterns -> State 356
      var_decl_pattern -> State 70
      var_or_let -> State 71
      assign_pattern -> State 264
      case_pattern -> State 265
      optional_label_decl -> State 85
      sequence_expr -> State 267
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 251:
    Kernel Items:
      if_expr: IF condition.statement_block
      if_expr: IF condition.statement_block ELSE statement_block
      if_expr: IF condition.statement_block ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 73
      statement_block -> State 357

  State 252:
    Kernel Items:
      condition: sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 253:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 73
      statement_block -> State 358

  State 254:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      binary_or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [binary_or_expr]
    Goto:
      AND -> State 131

  State 255:
    Kernel Items:
      field_var_pattern: ELLIPSIS., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 256:
    Kernel Items:
      var_pattern: IDENTIFIER., *
      field_var_pattern: IDENTIFIER.ASSIGN var_pattern
    Reduce:
      * -> [var_pattern]
    Goto:
      ASSIGN -> State 359

  State 257:
    Kernel Items:
      field_var_patterns: field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 258:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns.RPAREN
      field_var_patterns: field_var_patterns.COMMA field_var_pattern
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 361
      COMMA -> State 360

  State 259:
    Kernel Items:
      field_var_pattern: var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 260:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern optional_value_type., $
    Reduce:
      $ -> [var_decl_pattern]
    Goto:
      (nil)

  State 261:
    Kernel Items:
      optional_value_type: value_type., *
      trait_op_type: value_type.trait_op returnable_type
    Reduce:
      * -> [optional_value_type]
    Goto:
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 262:
    Kernel Items:
      case_pattern: DOT.IDENTIFIER implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 362

  State 263:
    Kernel Items:
      var_or_let: VAR., *
      case_pattern: VAR.DOT IDENTIFIER tuple_pattern
    Reduce:
      * -> [var_or_let]
    Goto:
      DOT -> State 363

  State 264:
    Kernel Items:
      case_pattern: assign_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 265:
    Kernel Items:
      case_patterns: case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 266:
    Kernel Items:
      statement: CASE case_patterns.COLON optional_simple_statement
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    Goto:
      COMMA -> State 365
      COLON -> State 364

  State 267:
    Kernel Items:
      assign_pattern: sequence_expr., *
    Reduce:
      * -> [assign_pattern]
    Goto:
      (nil)

  State 268:
    Kernel Items:
      statement: DEFAULT COLON.optional_simple_statement
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_simple_statement]
      RBRACE -> [optional_simple_statement]
      SEMICOLON -> [optional_simple_statement]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      RETURN -> State 169
      BREAK -> State 162
      CONTINUE -> State 164
      FALLTHROUGH -> State 167
      UNSAFE -> State 170
      STRUCT -> State 32
      FUNC -> State 19
      ASYNC -> State 161
      DEFER -> State 166
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      simple_statement -> State 367
      optional_simple_statement -> State 366
      expression_or_improper_struct_statement -> State 178
      expressions -> State 179
      callback_op -> State 175
      callback_op_statement -> State 176
      unsafe_statement -> State 190
      jump_statement -> State 182
      jump_type -> State 183
      fallthrough_statement -> State 180
      assign_statement -> State 173
      unary_op_assign_statement -> State 189
      binary_op_assign_statement -> State 174
      var_decl_pattern -> State 70
      var_or_let -> State 71
      assign_pattern -> State 172
      expression -> State 177
      optional_label_decl -> State 60
      sequence_expr -> State 185
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 171
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 269:
    Kernel Items:
      import_statement: IMPORT LPAREN.import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 270
      import_clause -> State 368
      proper_import_clauses -> State 370
      import_clauses -> State 369

  State 270:
    Kernel Items:
      import_clause: STRING_LITERAL., *
      import_clause: STRING_LITERAL.AS IDENTIFIER
    Reduce:
      * -> [import_clause]
    Goto:
      AS -> State 371

  State 271:
    Kernel Items:
      import_statement: IMPORT import_clause., *
    Reduce:
      * -> [import_statement]
    Goto:
      (nil)

  State 272:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 372

  State 273:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 274:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 275:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 276:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 277:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 280:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 281:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 282:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 283:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 284:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 285:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., *
    Reduce:
      * -> [unary_op_assign]
    Goto:
      (nil)

  State 286:
    Kernel Items:
      binary_op_assign_statement: accessible_expr binary_op_assign.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 373
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 287:
    Kernel Items:
      unary_op_assign_statement: accessible_expr unary_op_assign., *
    Reduce:
      * -> [unary_op_assign_statement]
    Goto:
      (nil)

  State 288:
    Kernel Items:
      assign_statement: assign_pattern ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 374
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 289:
    Kernel Items:
      call_expr: accessible_expr.optional_generic_binding LPAREN arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
    Reduce:
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 123
      DOT -> State 122
      DOLLAR_LBRACKET -> State 121
      optional_generic_binding -> State 125

  State 290:
    Kernel Items:
      callback_op_statement: callback_op call_expr., NEWLINES
      callback_op_statement: callback_op call_expr., RBRACE
      callback_op_statement: callback_op call_expr., SEMICOLON
      accessible_expr: call_expr., *
    Reduce:
      * -> [accessible_expr]
      NEWLINES -> [callback_op_statement]
      RBRACE -> [callback_op_statement]
      SEMICOLON -> [callback_op_statement]
    Goto:
      (nil)

  State 291:
    Kernel Items:
      expressions: expressions COMMA.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 375
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 292:
    Kernel Items:
      jump_statement: jump_type JUMP_LABEL., NEWLINES
      jump_statement: jump_type JUMP_LABEL., RBRACE
      jump_statement: jump_type JUMP_LABEL., SEMICOLON
      jump_statement: jump_type JUMP_LABEL.expressions
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [jump_statement]
      RBRACE -> [jump_statement]
      SEMICOLON -> [jump_statement]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      expressions -> State 376
      var_decl_pattern -> State 70
      var_or_let -> State 71
      expression -> State 177
      optional_label_decl -> State 60
      sequence_expr -> State 68
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 293:
    Kernel Items:
      expressions: expressions.COMMA expression
      jump_statement: jump_type expressions., *
    Reduce:
      * -> [jump_statement]
    Goto:
      COMMA -> State 291

  State 294:
    Kernel Items:
      proper_statements: proper_statements NEWLINES.statement
      statements: proper_statements NEWLINES., RBRACE
    Reduce:
      * -> [optional_label_decl]
      RBRACE -> [statements]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      CASE -> State 163
      DEFAULT -> State 165
      RETURN -> State 169
      BREAK -> State 162
      CONTINUE -> State 164
      FALLTHROUGH -> State 167
      IMPORT -> State 168
      UNSAFE -> State 170
      STRUCT -> State 32
      FUNC -> State 19
      ASYNC -> State 161
      DEFER -> State 166
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      simple_statement -> State 186
      statement -> State 377
      expression_or_improper_struct_statement -> State 178
      expressions -> State 179
      callback_op -> State 175
      callback_op_statement -> State 176
      unsafe_statement -> State 190
      jump_statement -> State 182
      jump_type -> State 183
      fallthrough_statement -> State 180
      assign_statement -> State 173
      unary_op_assign_statement -> State 189
      binary_op_assign_statement -> State 174
      import_statement -> State 181
      var_decl_pattern -> State 70
      var_or_let -> State 71
      assign_pattern -> State 172
      expression -> State 177
      optional_label_decl -> State 60
      sequence_expr -> State 185
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 171
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 295:
    Kernel Items:
      proper_statements: proper_statements SEMICOLON.statement
      statements: proper_statements SEMICOLON., RBRACE
    Reduce:
      * -> [optional_label_decl]
      RBRACE -> [statements]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      CASE -> State 163
      DEFAULT -> State 165
      RETURN -> State 169
      BREAK -> State 162
      CONTINUE -> State 164
      FALLTHROUGH -> State 167
      IMPORT -> State 168
      UNSAFE -> State 170
      STRUCT -> State 32
      FUNC -> State 19
      ASYNC -> State 161
      DEFER -> State 166
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      simple_statement -> State 186
      statement -> State 378
      expression_or_improper_struct_statement -> State 178
      expressions -> State 179
      callback_op -> State 175
      callback_op_statement -> State 176
      unsafe_statement -> State 190
      jump_statement -> State 182
      jump_type -> State 183
      fallthrough_statement -> State 180
      assign_statement -> State 173
      unary_op_assign_statement -> State 189
      binary_op_assign_statement -> State 174
      import_statement -> State 181
      var_decl_pattern -> State 70
      var_or_let -> State 71
      assign_pattern -> State 172
      expression -> State 177
      optional_label_decl -> State 60
      sequence_expr -> State 185
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 171
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 296:
    Kernel Items:
      statement_block: LBRACE statements RBRACE., $
    Reduce:
      $ -> [statement_block]
    Goto:
      (nil)

  State 297:
    Kernel Items:
      definition: var_decl_pattern ASSIGN expression., NEWLINES
    Reduce:
      NEWLINES -> [definition]
    Goto:
      (nil)

  State 298:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      type_def: TYPE IDENTIFIER ASSIGN value_type., *
    Reduce:
      * -> [type_def]
    Goto:
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 299:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 379
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 300:
    Kernel Items:
      proper_generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [proper_generic_parameter_defs]
    Goto:
      (nil)

  State 301:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 380

  State 302:
    Kernel Items:
      proper_generic_parameter_defs: proper_generic_parameter_defs.COMMA generic_parameter_def
      generic_parameter_defs: proper_generic_parameter_defs., RBRACKET
      generic_parameter_defs: proper_generic_parameter_defs.COMMA
    Reduce:
      RBRACKET -> [generic_parameter_defs]
    Goto:
      COMMA -> State 381

  State 303:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., *
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 382
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 304:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN expression., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 305:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 198
      parameter_def -> State 200
      proper_parameter_defs -> State 202
      parameter_defs -> State 383

  State 306:
    Kernel Items:
      parameter_def: IDENTIFIER ELLIPSIS., *
      parameter_def: IDENTIFIER ELLIPSIS.value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 384
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 307:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 308:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN.IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 385

  State 309:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 387
      prefixed_type -> State 107
      prefix_type_op -> State 106
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      return_type -> State 386
      func_type -> State 101

  State 310:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs COMMA.parameter_def
      parameter_defs: proper_parameter_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 198
      parameter_def -> State 388

  State 311:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 389
      OR -> State 390

  State 312:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 391

  State 313:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 325

  State 314:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 392
      OR -> State 393

  State 315:
    Kernel Items:
      parameter_decl: ELLIPSIS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 394
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 316:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.ELLIPSIS value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 322
      QUESTION -> State 96
      EXCLAIM -> State 91
      DOLLAR_LBRACKET -> State 121
      ELLIPSIS -> State 395
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      optional_generic_binding -> State 207
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 396
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 317:
    Kernel Items:
      proper_parameter_decls: parameter_decl., *
    Reduce:
      * -> [proper_parameter_decls]
    Goto:
      (nil)

  State 318:
    Kernel Items:
      func_type: FUNC LPAREN parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 397

  State 319:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls.COMMA parameter_decl
      parameter_decls: proper_parameter_decls., RPAREN
      parameter_decls: proper_parameter_decls.COMMA
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      COMMA -> State 398

  State 320:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 321:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 121
      optional_generic_binding -> State 399

  State 322:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 321
      DOLLAR_LBRACKET -> State 121
      optional_generic_binding -> State 203

  State 323:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 324:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 208
      UNSAFE -> State 170
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 214
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 215
      trait_op_type -> State 110
      field_def -> State 313
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      enum_value_def -> State 400
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 325:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 401

  State 326:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 208
      UNSAFE -> State 170
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 214
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 215
      trait_op_type -> State 110
      field_def -> State 313
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      enum_value_def -> State 402
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 327:
    Kernel Items:
      implicit_enum_def: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_def]
    Goto:
      (nil)

  State 328:
    Kernel Items:
      implicit_struct_def: LPAREN implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_def]
    Goto:
      (nil)

  State 329:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs COMMA.field_def
      implicit_field_defs: proper_implicit_field_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      IDENTIFIER -> State 208
      UNSAFE -> State 170
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 214
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 215
      trait_op_type -> State 110
      field_def -> State 403
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 330:
    Kernel Items:
      func_type: FUNC.LPAREN parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 404
      LPAREN -> State 205

  State 331:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 332:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 333:
    Kernel Items:
      proper_trait_properties: proper_trait_properties.NEWLINES trait_property
      proper_trait_properties: proper_trait_properties.COMMA trait_property
      trait_properties: proper_trait_properties., RPAREN
      trait_properties: proper_trait_properties.NEWLINES
      trait_properties: proper_trait_properties.COMMA
    Reduce:
      RPAREN -> [trait_properties]
    Goto:
      NEWLINES -> State 406
      COMMA -> State 405

  State 334:
    Kernel Items:
      trait_def: TRAIT LPAREN trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 407

  State 335:
    Kernel Items:
      proper_trait_properties: trait_property., *
    Reduce:
      * -> [proper_trait_properties]
    Goto:
      (nil)

  State 336:
    Kernel Items:
      map_type: LBRACKET value_type COLON value_type.RBRACKET
      trait_op_type: value_type.trait_op returnable_type
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 408
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 337:
    Kernel Items:
      array_type: LBRACKET value_type COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 409

  State 338:
    Kernel Items:
      trait_op_type: value_type trait_op returnable_type., *
    Reduce:
      * -> [trait_op_type]
    Goto:
      (nil)

  State 339:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 340:
    Kernel Items:
      colon_expr: colon_expr COLON expression., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      colon_expr: expression COLON expression., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 342:
    Kernel Items:
      proper_arguments: proper_arguments COMMA argument., *
    Reduce:
      * -> [proper_arguments]
    Goto:
      (nil)

  State 343:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_def]
    Goto:
      (nil)

  State 344:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs COMMA.field_def
      explicit_field_defs: proper_explicit_field_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 208
      UNSAFE -> State 170
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 214
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 215
      trait_op_type -> State 110
      field_def -> State 410
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 345:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs NEWLINES.field_def
      explicit_field_defs: proper_explicit_field_defs NEWLINES., RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 208
      UNSAFE -> State 170
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 214
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 215
      trait_op_type -> State 110
      field_def -> State 411
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 346:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 347:
    Kernel Items:
      proper_generic_arguments: proper_generic_arguments COMMA.value_type
      generic_arguments: proper_generic_arguments COMMA., RBRACKET
    Reduce:
      RBRACKET -> [generic_arguments]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 412
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 348:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [index_expr]
    Goto:
      (nil)

  State 349:
    Kernel Items:
      call_expr: accessible_expr optional_generic_binding LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 413

  State 350:
    Kernel Items:
      initialize_expr: initializable_type LPAREN arguments RPAREN., *
    Reduce:
      * -> [initialize_expr]
    Goto:
      (nil)

  State 351:
    Kernel Items:
      loop_expr: DO statement_block FOR.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      optional_label_decl -> State 85
      sequence_expr -> State 414
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 352:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      optional_label_decl -> State 85
      sequence_expr -> State 415
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 353:
    Kernel Items:
      loop_expr: FOR assign_pattern IN.sequence_expr DO statement_block
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      optional_label_decl -> State 85
      sequence_expr -> State 416
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 354:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON.optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      LBRACE -> [optional_label_decl]
      SEMICOLON -> [optional_sequence_expr]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      optional_label_decl -> State 85
      sequence_expr -> State 418
      optional_sequence_expr -> State 417
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 355:
    Kernel Items:
      loop_expr: FOR sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 73
      statement_block -> State 419

  State 356:
    Kernel Items:
      case_patterns: case_patterns.COMMA case_pattern
      condition: CASE case_patterns.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      COMMA -> State 365
      ASSIGN -> State 420

  State 357:
    Kernel Items:
      if_expr: IF condition statement_block., *
      if_expr: IF condition statement_block.ELSE statement_block
      if_expr: IF condition statement_block.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 421

  State 358:
    Kernel Items:
      switch_expr: SWITCH sequence_expr statement_block., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 359:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN.var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 157
      LPAREN -> State 158
      var_pattern -> State 422
      tuple_pattern -> State 159

  State 360:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA.field_var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 256
      LPAREN -> State 158
      ELLIPSIS -> State 255
      var_pattern -> State 259
      tuple_pattern -> State 159
      field_var_pattern -> State 423

  State 361:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns RPAREN., *
    Reduce:
      * -> [tuple_pattern]
    Goto:
      (nil)

  State 362:
    Kernel Items:
      case_pattern: DOT IDENTIFIER.implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 26
      implicit_struct_expr -> State 424

  State 363:
    Kernel Items:
      case_pattern: VAR DOT.IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 425

  State 364:
    Kernel Items:
      statement: CASE case_patterns COLON.optional_simple_statement
    Reduce:
      * -> [optional_label_decl]
      NEWLINES -> [optional_simple_statement]
      RBRACE -> [optional_simple_statement]
      SEMICOLON -> [optional_simple_statement]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      RETURN -> State 169
      BREAK -> State 162
      CONTINUE -> State 164
      FALLTHROUGH -> State 167
      UNSAFE -> State 170
      STRUCT -> State 32
      FUNC -> State 19
      ASYNC -> State 161
      DEFER -> State 166
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      simple_statement -> State 367
      optional_simple_statement -> State 426
      expression_or_improper_struct_statement -> State 178
      expressions -> State 179
      callback_op -> State 175
      callback_op_statement -> State 176
      unsafe_statement -> State 190
      jump_statement -> State 182
      jump_type -> State 183
      fallthrough_statement -> State 180
      assign_statement -> State 173
      unary_op_assign_statement -> State 189
      binary_op_assign_statement -> State 174
      var_decl_pattern -> State 70
      var_or_let -> State 71
      assign_pattern -> State 172
      expression -> State 177
      optional_label_decl -> State 60
      sequence_expr -> State 185
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 171
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 365:
    Kernel Items:
      case_patterns: case_patterns COMMA.case_pattern
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 263
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      DOT -> State 262
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      assign_pattern -> State 264
      case_pattern -> State 427
      optional_label_decl -> State 85
      sequence_expr -> State 267
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 366:
    Kernel Items:
      statement: DEFAULT COLON optional_simple_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 367:
    Kernel Items:
      optional_simple_statement: simple_statement., *
    Reduce:
      * -> [optional_simple_statement]
    Goto:
      (nil)

  State 368:
    Kernel Items:
      proper_import_clauses: import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 369:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 428

  State 370:
    Kernel Items:
      proper_import_clauses: proper_import_clauses.NEWLINES import_clause
      proper_import_clauses: proper_import_clauses.COMMA import_clause
      import_clauses: proper_import_clauses., RPAREN
      import_clauses: proper_import_clauses.NEWLINES
      import_clauses: proper_import_clauses.COMMA
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      NEWLINES -> State 430
      COMMA -> State 429

  State 371:
    Kernel Items:
      import_clause: STRING_LITERAL AS.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 431

  State 372:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 432

  State 373:
    Kernel Items:
      binary_op_assign_statement: accessible_expr binary_op_assign expression., *
    Reduce:
      * -> [binary_op_assign_statement]
    Goto:
      (nil)

  State 374:
    Kernel Items:
      assign_statement: assign_pattern ASSIGN expression., *
    Reduce:
      * -> [assign_statement]
    Goto:
      (nil)

  State 375:
    Kernel Items:
      expressions: expressions COMMA expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 376:
    Kernel Items:
      expressions: expressions.COMMA expression
      jump_statement: jump_type JUMP_LABEL expressions., *
    Reduce:
      * -> [jump_statement]
    Goto:
      COMMA -> State 291

  State 377:
    Kernel Items:
      proper_statements: proper_statements NEWLINES statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 378:
    Kernel Items:
      proper_statements: proper_statements SEMICOLON statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 379:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 380:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 381:
    Kernel Items:
      proper_generic_parameter_defs: proper_generic_parameter_defs COMMA.generic_parameter_def
      generic_parameter_defs: proper_generic_parameter_defs COMMA., RBRACKET
    Reduce:
      RBRACKET -> [generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 299
      generic_parameter_def -> State 433

  State 382:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 434
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 383:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 435

  State 384:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_def: IDENTIFIER ELLIPSIS value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 385:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER.optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 194
      optional_generic_parameters -> State 436

  State 386:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 73
      statement_block -> State 437

  State 387:
    Kernel Items:
      return_type: returnable_type., *
    Reduce:
      * -> [return_type]
    Goto:
      (nil)

  State 388:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [proper_parameter_defs]
    Goto:
      (nil)

  State 389:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 208
      UNSAFE -> State 170
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 214
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 215
      trait_op_type -> State 110
      field_def -> State 313
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      enum_value_def -> State 438
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 390:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 208
      UNSAFE -> State 170
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 214
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 215
      trait_op_type -> State 110
      field_def -> State 313
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      enum_value_def -> State 439
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 391:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_def]
    Goto:
      (nil)

  State 392:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 208
      UNSAFE -> State 170
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 214
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 215
      trait_op_type -> State 110
      field_def -> State 313
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      enum_value_def -> State 440
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 393:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 208
      UNSAFE -> State 170
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 214
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 215
      trait_op_type -> State 110
      field_def -> State 313
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      enum_value_def -> State 441
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 394:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_decl: ELLIPSIS value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 395:
    Kernel Items:
      parameter_decl: IDENTIFIER ELLIPSIS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 442
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      func_type -> State 101

  State 396:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 397:
    Kernel Items:
      func_type: FUNC LPAREN parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 387
      prefixed_type -> State 107
      prefix_type_op -> State 106
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      return_type -> State 443
      func_type -> State 101

  State 398:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls COMMA.parameter_decl
      parameter_decls: proper_parameter_decls COMMA., RPAREN
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 316
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      ELLIPSIS -> State 315
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 320
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      parameter_decl -> State 444
      func_type -> State 101

  State 399:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 400:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 401:
    Kernel Items:
      enum_value_def: field_def ASSIGN DEFAULT., *
    Reduce:
      * -> [enum_value_def]
    Goto:
      (nil)

  State 402:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 403:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [proper_implicit_field_defs]
    Goto:
      (nil)

  State 404:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 445

  State 405:
    Kernel Items:
      proper_trait_properties: proper_trait_properties COMMA.trait_property
      trait_properties: proper_trait_properties COMMA., RPAREN
    Reduce:
      RPAREN -> [trait_properties]
    Goto:
      IDENTIFIER -> State 208
      UNSAFE -> State 170
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 330
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 214
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 215
      trait_op_type -> State 110
      field_def -> State 331
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_property -> State 446
      trait_def -> State 109
      func_type -> State 101
      method_signature -> State 332

  State 406:
    Kernel Items:
      proper_trait_properties: proper_trait_properties NEWLINES.trait_property
      trait_properties: proper_trait_properties NEWLINES., RPAREN
    Reduce:
      RPAREN -> [trait_properties]
    Goto:
      IDENTIFIER -> State 208
      UNSAFE -> State 170
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 330
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      unsafe_statement -> State 214
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 215
      trait_op_type -> State 110
      field_def -> State 331
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_property -> State 447
      trait_def -> State 109
      func_type -> State 101
      method_signature -> State 332

  State 407:
    Kernel Items:
      trait_def: TRAIT LPAREN trait_properties RPAREN., *
    Reduce:
      * -> [trait_def]
    Goto:
      (nil)

  State 408:
    Kernel Items:
      map_type: LBRACKET value_type COLON value_type RBRACKET., *
    Reduce:
      * -> [map_type]
    Goto:
      (nil)

  State 409:
    Kernel Items:
      array_type: LBRACKET value_type COMMA INTEGER_LITERAL RBRACKET., *
    Reduce:
      * -> [array_type]
    Goto:
      (nil)

  State 410:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 411:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 412:
    Kernel Items:
      proper_generic_arguments: proper_generic_arguments COMMA value_type., *
      trait_op_type: value_type.trait_op returnable_type
    Reduce:
      * -> [proper_generic_arguments]
    Goto:
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 413:
    Kernel Items:
      call_expr: accessible_expr optional_generic_binding LPAREN arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 414:
    Kernel Items:
      loop_expr: DO statement_block FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 415:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN sequence_expr., SEMICOLON
    Reduce:
      SEMICOLON -> [for_assignment]
    Goto:
      (nil)

  State 416:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 448

  State 417:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 449

  State 418:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 419:
    Kernel Items:
      loop_expr: FOR sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 420:
    Kernel Items:
      condition: CASE case_patterns ASSIGN.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      optional_label_decl -> State 85
      sequence_expr -> State 450
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 421:
    Kernel Items:
      if_expr: IF condition statement_block ELSE.statement_block
      if_expr: IF condition statement_block ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 149
      LBRACE -> State 73
      statement_block -> State 452
      if_expr -> State 451

  State 422:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 423:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 424:
    Kernel Items:
      case_pattern: DOT IDENTIFIER implicit_struct_expr., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 425:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER.tuple_pattern
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 158
      tuple_pattern -> State 453

  State 426:
    Kernel Items:
      statement: CASE case_patterns COLON optional_simple_statement., *
    Reduce:
      * -> [statement]
    Goto:
      (nil)

  State 427:
    Kernel Items:
      case_patterns: case_patterns COMMA case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 428:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses RPAREN., *
    Reduce:
      * -> [import_statement]
    Goto:
      (nil)

  State 429:
    Kernel Items:
      proper_import_clauses: proper_import_clauses COMMA.import_clause
      import_clauses: proper_import_clauses COMMA., RPAREN
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      STRING_LITERAL -> State 270
      import_clause -> State 454

  State 430:
    Kernel Items:
      proper_import_clauses: proper_import_clauses NEWLINES.import_clause
      import_clauses: proper_import_clauses NEWLINES., RPAREN
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      STRING_LITERAL -> State 270
      import_clause -> State 455

  State 431:
    Kernel Items:
      import_clause: STRING_LITERAL AS IDENTIFIER., *
    Reduce:
      * -> [import_clause]
    Goto:
      (nil)

  State 432:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 456

  State 433:
    Kernel Items:
      proper_generic_parameter_defs: proper_generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [proper_generic_parameter_defs]
    Goto:
      (nil)

  State 434:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., *
    Reduce:
      * -> [type_def]
    Goto:
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 435:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 387
      prefixed_type -> State 107
      prefix_type_op -> State 106
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      return_type -> State 457
      func_type -> State 101

  State 436:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 458

  State 437:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 438:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 439:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
      explicit_enum_value_defs: enum_value_def OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 440:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 441:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 442:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_decl: IDENTIFIER ELLIPSIS value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 218
      SUB -> State 223
      MUL -> State 221
      trait_op -> State 224

  State 443:
    Kernel Items:
      func_type: FUNC LPAREN parameter_decls RPAREN return_type., *
    Reduce:
      * -> [func_type]
    Goto:
      (nil)

  State 444:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [proper_parameter_decls]
    Goto:
      (nil)

  State 445:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 316
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      ELLIPSIS -> State 315
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 108
      prefixed_type -> State 107
      prefix_type_op -> State 106
      value_type -> State 320
      trait_op_type -> State 110
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      parameter_decl -> State 317
      proper_parameter_decls -> State 319
      parameter_decls -> State 459
      func_type -> State 101

  State 446:
    Kernel Items:
      proper_trait_properties: proper_trait_properties COMMA trait_property., *
    Reduce:
      * -> [proper_trait_properties]
    Goto:
      (nil)

  State 447:
    Kernel Items:
      proper_trait_properties: proper_trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [proper_trait_properties]
    Goto:
      (nil)

  State 448:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 73
      statement_block -> State 460

  State 449:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON.optional_sequence_expr DO statement_block
    Reduce:
      DO -> [optional_sequence_expr]
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 22
      FLOAT_LITERAL -> State 18
      RUNE_LITERAL -> State 30
      STRING_LITERAL -> State 31
      IDENTIFIER -> State 21
      TRUE -> State 34
      FALSE -> State 17
      STRUCT -> State 32
      FUNC -> State 19
      VAR -> State 35
      LET -> State 25
      NOT -> State 28
      LABEL_DECL -> State 23
      LPAREN -> State 26
      LBRACKET -> State 24
      SUB -> State 33
      MUL -> State 27
      BIT_NEG -> State 16
      BIT_AND -> State 15
      GREATER -> State 20
      PARSE_ERROR -> State 29
      var_decl_pattern -> State 70
      var_or_let -> State 71
      optional_label_decl -> State 85
      sequence_expr -> State 418
      optional_sequence_expr -> State 461
      call_expr -> State 49
      atom_expr -> State 42
      parse_error_expr -> State 62
      literal_expr -> State 57
      identifier_expr -> State 52
      block_expr -> State 48
      initialize_expr -> State 56
      implicit_struct_expr -> State 53
      accessible_expr -> State 37
      access_expr -> State 36
      index_expr -> State 54
      postfixable_expr -> State 64
      postfix_unary_expr -> State 63
      prefixable_expr -> State 67
      prefix_unary_expr -> State 65
      prefix_unary_op -> State 66
      mul_expr -> State 59
      binary_mul_expr -> State 46
      add_expr -> State 38
      binary_add_expr -> State 43
      cmp_expr -> State 50
      binary_cmp_expr -> State 45
      and_expr -> State 39
      binary_and_expr -> State 44
      or_expr -> State 61
      binary_or_expr -> State 47
      initializable_type -> State 55
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      explicit_struct_def -> State 51
      anonymous_func_expr -> State 40

  State 450:
    Kernel Items:
      condition: CASE case_patterns ASSIGN sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 451:
    Kernel Items:
      if_expr: IF condition statement_block ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 452:
    Kernel Items:
      if_expr: IF condition statement_block ELSE statement_block., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 453:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER tuple_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 454:
    Kernel Items:
      proper_import_clauses: proper_import_clauses COMMA import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 455:
    Kernel Items:
      proper_import_clauses: proper_import_clauses NEWLINES import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 456:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., *
    Reduce:
      * -> [unsafe_statement]
    Goto:
      (nil)

  State 457:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 73
      statement_block -> State 462

  State 458:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 198
      parameter_def -> State 200
      proper_parameter_defs -> State 202
      parameter_defs -> State 463

  State 459:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 464

  State 460:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 461:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 465

  State 462:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

  State 463:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 466

  State 464:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 387
      prefixed_type -> State 107
      prefix_type_op -> State 106
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      return_type -> State 467
      func_type -> State 101

  State 465:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 73
      statement_block -> State 468

  State 466:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 93
      STRUCT -> State 32
      ENUM -> State 90
      TRAIT -> State 98
      FUNC -> State 92
      LPAREN -> State 94
      LBRACKET -> State 24
      DOT -> State 89
      QUESTION -> State 96
      EXCLAIM -> State 91
      TILDE_TILDE -> State 97
      BIT_NEG -> State 88
      BIT_AND -> State 87
      PARSE_ERROR -> State 95
      initializable_type -> State 104
      slice_type -> State 69
      array_type -> State 41
      map_type -> State 58
      atom_type -> State 99
      parse_error_type -> State 105
      returnable_type -> State 387
      prefixed_type -> State 107
      prefix_type_op -> State 106
      implicit_struct_def -> State 103
      explicit_struct_def -> State 51
      implicit_enum_def -> State 102
      explicit_enum_def -> State 100
      trait_def -> State 109
      return_type -> State 469
      func_type -> State 101

  State 467:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 468:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 469:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 73
      statement_block -> State 470

  State 470:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block., $
    Reduce:
      $ -> [named_func_def]
    Goto:
      (nil)

Number of states: 470
Number of shift actions: 4171
Number of reduce actions: 412
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 4374
Number of unoptimized shift actions: 37626
Number of unoptimized reduce actions: 27942
*/
