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

type ProperDefinitionsReducer interface {
	// 55:2: proper_definitions -> add: ...
	AddToProperDefinitions(ProperDefinitions_ []SourceDefinition, Newlines_ TokenCount, Definition_ SourceDefinition) ([]SourceDefinition, error)

	// 56:2: proper_definitions -> definition: ...
	DefinitionToProperDefinitions(Definition_ SourceDefinition) ([]SourceDefinition, error)
}

type DefinitionsReducer interface {

	// 60:2: definitions -> improper: ...
	ImproperToDefinitions(ProperDefinitions_ []SourceDefinition, Newlines_ TokenCount) ([]SourceDefinition, error)

	// 61:2: definitions -> nil: ...
	NilToDefinitions() ([]SourceDefinition, error)
}

type DefinitionReducer interface {

	// 68:2: definition -> global_var_def: ...
	GlobalVarDefToDefinition(VarDeclPattern_ GenericSymbol) (SourceDefinition, error)

	// 69:2: definition -> global_var_assignment: ...
	GlobalVarAssignmentToDefinition(VarDeclPattern_ GenericSymbol, Assign_ TokenValue, Expression_ Expression) (SourceDefinition, error)

	// 72:2: definition -> statement_block: ...
	StatementBlockToDefinition(StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 76:2: definition -> COMMENT_GROUPS: ...
	CommentGroupsToDefinition(CommentGroups_ TokenValue) (SourceDefinition, error)
}

type StatementBlockReducer interface {
	// 96:19: statement_block -> ...
	ToStatementBlock(Lbrace_ TokenValue, Statements_ []Statement, Rbrace_ TokenValue) (GenericSymbol, error)
}

type ProperStatementsReducer interface {
	// 99:2: proper_statements -> add_implicit: ...
	AddImplicitToProperStatements(ProperStatements_ []Statement, Newlines_ TokenCount, Statement_ Statement) ([]Statement, error)

	// 100:2: proper_statements -> add_explicit: ...
	AddExplicitToProperStatements(ProperStatements_ []Statement, Semicolon_ TokenValue, Statement_ Statement) ([]Statement, error)

	// 101:2: proper_statements -> statement: ...
	StatementToProperStatements(Statement_ Statement) ([]Statement, error)
}

type StatementsReducer interface {

	// 105:2: statements -> improper_implicit: ...
	ImproperImplicitToStatements(ProperStatements_ []Statement, Newlines_ TokenCount) ([]Statement, error)

	// 106:2: statements -> improper_explicit: ...
	ImproperExplicitToStatements(ProperStatements_ []Statement, Semicolon_ TokenValue) ([]Statement, error)

	// 107:2: statements -> nil: ...
	NilToStatements() ([]Statement, error)
}

type StatementReducer interface {

	// 145:2: statement -> case_branch_statement: ...
	CaseBranchStatementToStatement(Case_ TokenValue, CasePatterns_ GenericSymbol, Colon_ TokenValue, OptionalSimpleStatement_ Statement) (Statement, error)

	// 146:2: statement -> default_branch_statement: ...
	DefaultBranchStatementToStatement(Default_ TokenValue, Colon_ TokenValue, OptionalSimpleStatement_ Statement) (Statement, error)
}

type OptionalSimpleStatementReducer interface {

	// 150:2: optional_simple_statement -> nil: ...
	NilToOptionalSimpleStatement() (Statement, error)
}

type ExpressionOrImproperStructStatementReducer interface {
	// 156:54: expression_or_improper_struct_statement -> ...
	ToExpressionOrImproperStructStatement(Expressions_ GenericSymbol) (Statement, error)
}

type ExpressionsReducer interface {
	// 159:2: expressions -> expression: ...
	ExpressionToExpressions(Expression_ Expression) (GenericSymbol, error)

	// 160:2: expressions -> add: ...
	AddToExpressions(Expressions_ GenericSymbol, Comma_ TokenValue, Expression_ Expression) (GenericSymbol, error)
}

type CallbackOpStatementReducer interface {
	// 186:36: callback_op_statement -> ...
	ToCallbackOpStatement(CallbackOp_ TokenValue, CallExpr_ Expression) (Statement, error)
}

type UnsafeStatementReducer interface {
	// 194:31: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ TokenValue, Less_ TokenValue, Identifier_ TokenValue, Greater_ TokenValue, StringLiteral_ TokenValue) (Statement, error)
}

type JumpStatementReducer interface {
	// 201:2: jump_statement -> unlabeled_no_value: ...
	UnlabeledNoValueToJumpStatement(JumpType_ TokenValue) (Statement, error)

	// 202:2: jump_statement -> unlabeled_valued: ...
	UnlabeledValuedToJumpStatement(JumpType_ TokenValue, Expressions_ GenericSymbol) (Statement, error)

	// 203:2: jump_statement -> labeled_no_value: ...
	LabeledNoValueToJumpStatement(JumpType_ TokenValue, JumpLabel_ TokenValue) (Statement, error)

	// 204:2: jump_statement -> labeled_valued: ...
	LabeledValuedToJumpStatement(JumpType_ TokenValue, JumpLabel_ TokenValue, Expressions_ GenericSymbol) (Statement, error)
}

type FallthroughStatementReducer interface {
	// 211:36: fallthrough_statement -> ...
	ToFallthroughStatement(Fallthrough_ TokenValue) (Statement, error)
}

type AssignStatementReducer interface {
	// 217:31: assign_statement -> ...
	ToAssignStatement(AssignPattern_ GenericSymbol, Assign_ TokenValue, Expression_ Expression) (Statement, error)
}

type UnaryOpAssignStatementReducer interface {
	// 219:40: unary_op_assign_statement -> ...
	ToUnaryOpAssignStatement(AccessibleExpr_ Expression, UnaryOpAssign_ TokenValue) (Statement, error)
}

type BinaryOpAssignStatementReducer interface {
	// 225:41: binary_op_assign_statement -> ...
	ToBinaryOpAssignStatement(AccessibleExpr_ Expression, BinaryOpAssign_ TokenValue, Expression_ Expression) (Statement, error)
}

type ImportStatementReducer interface {
	// 245:2: import_statement -> single: ...
	SingleToImportStatement(Import_ TokenValue, ImportClause_ GenericSymbol) (Statement, error)

	// 246:2: import_statement -> multiple: ...
	MultipleToImportStatement(Import_ TokenValue, Lparen_ TokenValue, ImportClauses_ GenericSymbol, Rparen_ TokenValue) (Statement, error)
}

type ImportClauseReducer interface {
	// 249:2: import_clause -> STRING_LITERAL: ...
	StringLiteralToImportClause(StringLiteral_ TokenValue) (GenericSymbol, error)

	// 250:2: import_clause -> alias: ...
	AliasToImportClause(StringLiteral_ TokenValue, As_ TokenValue, Identifier_ TokenValue) (GenericSymbol, error)
}

type ProperImportClausesReducer interface {
	// 253:2: proper_import_clauses -> add_implicit: ...
	AddImplicitToProperImportClauses(ProperImportClauses_ GenericSymbol, Newlines_ TokenCount, ImportClause_ GenericSymbol) (GenericSymbol, error)

	// 254:2: proper_import_clauses -> add_explicit: ...
	AddExplicitToProperImportClauses(ProperImportClauses_ GenericSymbol, Comma_ TokenValue, ImportClause_ GenericSymbol) (GenericSymbol, error)

	// 255:2: proper_import_clauses -> import_clause: ...
	ImportClauseToProperImportClauses(ImportClause_ GenericSymbol) (GenericSymbol, error)
}

type ImportClausesReducer interface {

	// 259:2: import_clauses -> implicit: ...
	ImplicitToImportClauses(ProperImportClauses_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)

	// 260:2: import_clauses -> explicit: ...
	ExplicitToImportClauses(ProperImportClauses_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)
}

type CasePatternsReducer interface {
	// 267:2: case_patterns -> case_pattern: ...
	CasePatternToCasePatterns(CasePattern_ GenericSymbol) (GenericSymbol, error)

	// 268:2: case_patterns -> multiple: ...
	MultipleToCasePatterns(CasePatterns_ GenericSymbol, Comma_ TokenValue, CasePattern_ GenericSymbol) (GenericSymbol, error)
}

type VarDeclPatternReducer interface {
	// 277:20: var_decl_pattern -> ...
	ToVarDeclPattern(VarOrLet_ TokenValue, VarPattern_ GenericSymbol, OptionalValueType_ TypeExpression) (GenericSymbol, error)
}

type VarPatternReducer interface {
	// 284:2: var_pattern -> IDENTIFIER: ...
	IdentifierToVarPattern(Identifier_ TokenValue) (GenericSymbol, error)

	// 285:2: var_pattern -> tuple_pattern: ...
	TuplePatternToVarPattern(TuplePattern_ GenericSymbol) (GenericSymbol, error)
}

type TuplePatternReducer interface {
	// 287:17: tuple_pattern -> ...
	ToTuplePattern(Lparen_ TokenValue, FieldVarPatterns_ GenericSymbol, Rparen_ TokenValue) (GenericSymbol, error)
}

type FieldVarPatternsReducer interface {
	// 290:2: field_var_patterns -> field_var_pattern: ...
	FieldVarPatternToFieldVarPatterns(FieldVarPattern_ GenericSymbol) (GenericSymbol, error)

	// 291:2: field_var_patterns -> add: ...
	AddToFieldVarPatterns(FieldVarPatterns_ GenericSymbol, Comma_ TokenValue, FieldVarPattern_ GenericSymbol) (GenericSymbol, error)
}

type FieldVarPatternReducer interface {
	// 294:2: field_var_pattern -> positional: ...
	PositionalToFieldVarPattern(VarPattern_ GenericSymbol) (GenericSymbol, error)

	// 295:2: field_var_pattern -> named: ...
	NamedToFieldVarPattern(Identifier_ TokenValue, Assign_ TokenValue, VarPattern_ GenericSymbol) (GenericSymbol, error)

	// 296:2: field_var_pattern -> ELLIPSIS: ...
	EllipsisToFieldVarPattern(Ellipsis_ TokenValue) (GenericSymbol, error)
}

type OptionalValueTypeReducer interface {

	// 300:2: optional_value_type -> nil: ...
	NilToOptionalValueType() (TypeExpression, error)
}

type AssignPatternReducer interface {
	// 307:2: assign_pattern -> ...
	ToAssignPattern(SequenceExpr_ Expression) (GenericSymbol, error)
}

type CasePatternReducer interface {

	// 317:2: case_pattern -> enum_match_pattern: ...
	EnumMatchPatternToCasePattern(Dot_ TokenValue, Identifier_ TokenValue, ImplicitStructExpr_ Expression) (GenericSymbol, error)

	// 318:2: case_pattern -> enum_var_decl_pattern: ...
	EnumVarDeclPatternToCasePattern(Var_ TokenValue, Dot_ TokenValue, Identifier_ TokenValue, TuplePattern_ GenericSymbol) (GenericSymbol, error)
}

type ExpressionReducer interface {
	// 325:2: expression -> if_expr: ...
	IfExprToExpression(OptionalLabelDecl_ TokenValue, IfExpr_ Expression) (Expression, error)

	// 326:2: expression -> switch_expr: ...
	SwitchExprToExpression(OptionalLabelDecl_ TokenValue, SwitchExpr_ Expression) (Expression, error)

	// 327:2: expression -> loop_expr: ...
	LoopExprToExpression(OptionalLabelDecl_ TokenValue, LoopExpr_ Expression) (Expression, error)
}

type OptionalLabelDeclReducer interface {
	// 331:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ TokenValue) (TokenValue, error)

	// 332:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (TokenValue, error)
}

type SequenceExprReducer interface {

	// 342:2: sequence_expr -> var_decl_pattern: ...
	VarDeclPatternToSequenceExpr(VarDeclPattern_ GenericSymbol) (Expression, error)

	// 346:2: sequence_expr -> assign_var_pattern: ...
	AssignVarPatternToSequenceExpr(Greater_ TokenValue, SequenceExpr_ Expression) (Expression, error)
}

type IfExprReducer interface {
	// 355:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol) (Expression, error)

	// 356:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol, Else_ TokenValue, StatementBlock_2 GenericSymbol) (Expression, error)

	// 357:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol, Else_ TokenValue, IfExpr_ Expression) (Expression, error)
}

type ConditionReducer interface {
	// 360:2: condition -> sequence_expr: ...
	SequenceExprToCondition(SequenceExpr_ Expression) (GenericSymbol, error)

	// 361:2: condition -> case: ...
	CaseToCondition(Case_ TokenValue, CasePatterns_ GenericSymbol, Assign_ TokenValue, SequenceExpr_ Expression) (GenericSymbol, error)
}

type SwitchExprReducer interface {
	// 385:27: switch_expr -> ...
	ToSwitchExpr(Switch_ TokenValue, SequenceExpr_ Expression, StatementBlock_ GenericSymbol) (Expression, error)
}

type LoopExprReducer interface {
	// 399:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 400:2: loop_expr -> do_while: ...
	DoWhileToLoopExpr(Do_ TokenValue, StatementBlock_ GenericSymbol, For_ TokenValue, SequenceExpr_ Expression) (Expression, error)

	// 401:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ TokenValue, SequenceExpr_ Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 402:2: loop_expr -> iterator: ...
	IteratorToLoopExpr(For_ TokenValue, AssignPattern_ GenericSymbol, In_ TokenValue, SequenceExpr_ Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 403:2: loop_expr -> for: ...
	ForToLoopExpr(For_ TokenValue, ForAssignment_ GenericSymbol, Semicolon_ TokenValue, OptionalSequenceExpr_ Expression, Semicolon_2 TokenValue, OptionalSequenceExpr_2 Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)
}

type OptionalSequenceExprReducer interface {
	// 406:2: optional_sequence_expr -> sequence_expr: ...
	SequenceExprToOptionalSequenceExpr(SequenceExpr_ Expression) (Expression, error)

	// 407:2: optional_sequence_expr -> nil: ...
	NilToOptionalSequenceExpr() (Expression, error)
}

type ForAssignmentReducer interface {
	// 410:2: for_assignment -> sequence_expr: ...
	SequenceExprToForAssignment(SequenceExpr_ Expression) (GenericSymbol, error)

	// 411:2: for_assignment -> assign: ...
	AssignToForAssignment(AssignPattern_ GenericSymbol, Assign_ TokenValue, SequenceExpr_ Expression) (GenericSymbol, error)
}

type CallExprReducer interface {
	// 418:2: call_expr -> ...
	ToCallExpr(AccessibleExpr_ Expression, OptionalGenericBinding_ GenericSymbol, Lparen_ TokenValue, Arguments_ *ArgumentList, Rparen_ TokenValue) (Expression, error)
}

type OptionalGenericBindingReducer interface {
	// 421:2: optional_generic_binding -> binding: ...
	BindingToOptionalGenericBinding(DollarLbracket_ TokenValue, GenericArguments_ GenericSymbol, Rbracket_ TokenValue) (GenericSymbol, error)

	// 422:2: optional_generic_binding -> nil: ...
	NilToOptionalGenericBinding() (GenericSymbol, error)
}

type ProperGenericArgumentsReducer interface {
	// 425:2: proper_generic_arguments -> add: ...
	AddToProperGenericArguments(ProperGenericArguments_ GenericSymbol, Comma_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 426:2: proper_generic_arguments -> value_type: ...
	ValueTypeToProperGenericArguments(ValueType_ TypeExpression) (GenericSymbol, error)
}

type GenericArgumentsReducer interface {

	// 430:2: generic_arguments -> improper: ...
	ImproperToGenericArguments(ProperGenericArguments_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 431:2: generic_arguments -> nil: ...
	NilToGenericArguments() (GenericSymbol, error)
}

type ProperArgumentsReducer interface {
	// 434:2: proper_arguments -> add: ...
	AddToProperArguments(ProperArguments_ *ArgumentList, Comma_ TokenValue, Argument_ *Argument) (*ArgumentList, error)

	// 435:2: proper_arguments -> argument: ...
	ArgumentToProperArguments(Argument_ *Argument) (*ArgumentList, error)
}

type ArgumentsReducer interface {

	// 439:2: arguments -> improper: ...
	ImproperToArguments(ProperArguments_ *ArgumentList, Comma_ TokenValue) (*ArgumentList, error)

	// 440:2: arguments -> nil: ...
	NilToArguments() (*ArgumentList, error)
}

type ArgumentReducer interface {
	// 443:2: argument -> positional: ...
	PositionalToArgument(Expression_ Expression) (*Argument, error)

	// 444:2: argument -> colon_expr: ...
	ColonExprToArgument(ColonExpr_ Expression) (*Argument, error)

	// 445:2: argument -> named_assignment: ...
	NamedAssignmentToArgument(Identifier_ TokenValue, Assign_ TokenValue, Expression_ Expression) (*Argument, error)

	// 449:2: argument -> vararg_assignment: ...
	VarargAssignmentToArgument(Expression_ Expression, Ellipsis_ TokenValue) (*Argument, error)

	// 452:2: argument -> skip_pattern: ...
	SkipPatternToArgument(Ellipsis_ TokenValue) (*Argument, error)
}

type ColonExprReducer interface {
	// 456:2: colon_expr -> unit_unit_pair: ...
	UnitUnitPairToColonExpr(Colon_ TokenValue) (Expression, error)

	// 457:2: colon_expr -> expr_unit_pair: ...
	ExprUnitPairToColonExpr(Expression_ Expression, Colon_ TokenValue) (Expression, error)

	// 458:2: colon_expr -> unit_expr_pair: ...
	UnitExprPairToColonExpr(Colon_ TokenValue, Expression_ Expression) (Expression, error)

	// 459:2: colon_expr -> expr_expr_pair: ...
	ExprExprPairToColonExpr(Expression_ Expression, Colon_ TokenValue, Expression_2 Expression) (Expression, error)

	// 460:2: colon_expr -> colon_expr_unit_tuple: ...
	ColonExprUnitTupleToColonExpr(ColonExpr_ Expression, Colon_ TokenValue) (Expression, error)

	// 461:2: colon_expr -> colon_expr_expr_tuple: ...
	ColonExprExprTupleToColonExpr(ColonExpr_ Expression, Colon_ TokenValue, Expression_ Expression) (Expression, error)
}

type ParseErrorExprReducer interface {
	// 480:32: parse_error_expr -> ...
	ToParseErrorExpr(ParseError_ ParseErrorSymbol) (Expression, error)
}

type LiteralExprReducer interface {
	// 483:2: literal_expr -> TRUE: ...
	TrueToLiteralExpr(True_ TokenValue) (Expression, error)

	// 484:2: literal_expr -> FALSE: ...
	FalseToLiteralExpr(False_ TokenValue) (Expression, error)

	// 485:2: literal_expr -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteralExpr(IntegerLiteral_ TokenValue) (Expression, error)

	// 486:2: literal_expr -> FLOAT_LITERAL: ...
	FloatLiteralToLiteralExpr(FloatLiteral_ TokenValue) (Expression, error)

	// 487:2: literal_expr -> RUNE_LITERAL: ...
	RuneLiteralToLiteralExpr(RuneLiteral_ TokenValue) (Expression, error)

	// 488:2: literal_expr -> STRING_LITERAL: ...
	StringLiteralToLiteralExpr(StringLiteral_ TokenValue) (Expression, error)
}

type IdentifierExprReducer interface {
	// 490:32: identifier_expr -> ...
	ToIdentifierExpr(Identifier_ TokenValue) (Expression, error)
}

type BlockExprReducer interface {
	// 492:27: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)
}

type InitializeExprReducer interface {
	// 494:31: initialize_expr -> ...
	ToInitializeExpr(InitializableType_ TypeExpression, Lparen_ TokenValue, Arguments_ *ArgumentList, Rparen_ TokenValue) (Expression, error)
}

type ImplicitStructExprReducer interface {
	// 496:36: implicit_struct_expr -> ...
	ToImplicitStructExpr(Lparen_ TokenValue, Arguments_ *ArgumentList, Rparen_ TokenValue) (Expression, error)
}

type AccessExprReducer interface {
	// 504:27: access_expr -> ...
	ToAccessExpr(AccessibleExpr_ Expression, Dot_ TokenValue, Identifier_ TokenValue) (Expression, error)
}

type IndexExprReducer interface {
	// 508:26: index_expr -> ...
	ToIndexExpr(AccessibleExpr_ Expression, Lbracket_ TokenValue, Argument_ *Argument, Rbracket_ TokenValue) (Expression, error)
}

type PostfixUnaryExprReducer interface {
	// 514:35: postfix_unary_expr -> ...
	ToPostfixUnaryExpr(AccessibleExpr_ Expression, Question_ TokenValue) (Expression, error)
}

type PrefixUnaryExprReducer interface {
	// 520:33: prefix_unary_expr -> ...
	ToPrefixUnaryExpr(PrefixUnaryOp_ TokenValue, PrefixableExpr_ Expression) (Expression, error)
}

type BinaryMulExprReducer interface {
	// 537:31: binary_mul_expr -> ...
	ToBinaryMulExpr(MulExpr_ Expression, MulOp_ TokenValue, PrefixableExpr_ Expression) (Expression, error)
}

type BinaryAddExprReducer interface {
	// 551:31: binary_add_expr -> ...
	ToBinaryAddExpr(AddExpr_ Expression, AddOp_ TokenValue, MulExpr_ Expression) (Expression, error)
}

type BinaryCmpExprReducer interface {
	// 563:31: binary_cmp_expr -> ...
	ToBinaryCmpExpr(CmpExpr_ Expression, CmpOp_ TokenValue, AddExpr_ Expression) (Expression, error)
}

type BinaryAndExprReducer interface {
	// 577:31: binary_and_expr -> ...
	ToBinaryAndExpr(AndExpr_ Expression, And_ TokenValue, CmpExpr_ Expression) (Expression, error)
}

type BinaryOrExprReducer interface {
	// 583:30: binary_or_expr -> ...
	ToBinaryOrExpr(OrExpr_ Expression, Or_ TokenValue, AndExpr_ Expression) (Expression, error)
}

type SliceTypeReducer interface {
	// 598:30: slice_type -> ...
	ToSliceType(Lbracket_ TokenValue, ValueType_ TypeExpression, Rbracket_ TokenValue) (TypeExpression, error)
}

type ArrayTypeReducer interface {
	// 600:30: array_type -> ...
	ToArrayType(Lbracket_ TokenValue, ValueType_ TypeExpression, Comma_ TokenValue, IntegerLiteral_ TokenValue, Rbracket_ TokenValue) (TypeExpression, error)
}

type MapTypeReducer interface {
	// 603:28: map_type -> ...
	ToMapType(Lbracket_ TokenValue, ValueType_ TypeExpression, Colon_ TokenValue, ValueType_2 TypeExpression, Rbracket_ TokenValue) (TypeExpression, error)
}

type AtomTypeReducer interface {

	// 607:2: atom_type -> named: ...
	NamedToAtomType(Identifier_ TokenValue, OptionalGenericBinding_ GenericSymbol) (TypeExpression, error)

	// 608:2: atom_type -> extern_named: ...
	ExternNamedToAtomType(Identifier_ TokenValue, Dot_ TokenValue, Identifier_2 TokenValue, OptionalGenericBinding_ GenericSymbol) (TypeExpression, error)

	// 609:2: atom_type -> inferred: ...
	InferredToAtomType(Dot_ TokenValue, OptionalGenericBinding_ GenericSymbol) (TypeExpression, error)
}

type ParseErrorTypeReducer interface {
	// 617:36: parse_error_type -> ...
	ToParseErrorType(ParseError_ ParseErrorSymbol) (TypeExpression, error)
}

type PrefixedTypeReducer interface {
	// 626:33: prefixed_type -> ...
	ToPrefixedType(PrefixTypeOp_ TokenValue, ReturnableType_ TypeExpression) (TypeExpression, error)
}

type TraitOpTypeReducer interface {
	// 641:33: trait_op_type -> ...
	ToTraitOpType(ValueType_ TypeExpression, TraitOp_ TokenValue, ReturnableType_ TypeExpression) (TypeExpression, error)
}

type TypeDefReducer interface {
	// 649:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, ValueType_ TypeExpression) (SourceDefinition, error)

	// 650:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, ValueType_ TypeExpression, Implements_ TokenValue, ValueType_2 TypeExpression) (SourceDefinition, error)

	// 651:2: type_def -> alias: ...
	AliasToTypeDef(Type_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, ValueType_ TypeExpression) (SourceDefinition, error)
}

type GenericParameterDefReducer interface {
	// 659:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ TokenValue) (GenericSymbol, error)

	// 660:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)
}

type ProperGenericParameterDefsReducer interface {
	// 663:2: proper_generic_parameter_defs -> add: ...
	AddToProperGenericParameterDefs(ProperGenericParameterDefs_ GenericSymbol, Comma_ TokenValue, GenericParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 664:2: proper_generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToProperGenericParameterDefs(GenericParameterDef_ GenericSymbol) (GenericSymbol, error)
}

type GenericParameterDefsReducer interface {

	// 668:2: generic_parameter_defs -> improper: ...
	ImproperToGenericParameterDefs(ProperGenericParameterDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 669:2: generic_parameter_defs -> nil: ...
	NilToGenericParameterDefs() (GenericSymbol, error)
}

type OptionalGenericParametersReducer interface {
	// 672:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ TokenValue, GenericParameterDefs_ GenericSymbol, Rbracket_ TokenValue) (GenericSymbol, error)

	// 673:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (GenericSymbol, error)
}

type FieldDefReducer interface {
	// 680:2: field_def -> explicit: ...
	ExplicitToFieldDef(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 681:2: field_def -> implicit: ...
	ImplicitToFieldDef(ValueType_ TypeExpression) (GenericSymbol, error)

	// 682:2: field_def -> unsafe_statement: ...
	UnsafeStatementToFieldDef(UnsafeStatement_ Statement) (GenericSymbol, error)
}

type ProperImplicitFieldDefsReducer interface {
	// 685:2: proper_implicit_field_defs -> add: ...
	AddToProperImplicitFieldDefs(ProperImplicitFieldDefs_ GenericSymbol, Comma_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 686:2: proper_implicit_field_defs -> field_def: ...
	FieldDefToProperImplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type ImplicitFieldDefsReducer interface {

	// 690:2: implicit_field_defs -> improper: ...
	ImproperToImplicitFieldDefs(ProperImplicitFieldDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 691:2: implicit_field_defs -> nil: ...
	NilToImplicitFieldDefs() (GenericSymbol, error)
}

type ImplicitStructDefReducer interface {
	// 694:2: implicit_struct_def -> ...
	ToImplicitStructDef(Lparen_ TokenValue, ImplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ProperExplicitFieldDefsReducer interface {
	// 697:2: proper_explicit_field_defs -> add_implicit: ...
	AddImplicitToProperExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Newlines_ TokenCount, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 698:2: proper_explicit_field_defs -> add_explicit: ...
	AddExplicitToProperExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Comma_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 699:2: proper_explicit_field_defs -> field_def: ...
	FieldDefToProperExplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type ExplicitFieldDefsReducer interface {

	// 703:2: explicit_field_defs -> improper_implicit: ...
	ImproperImplicitToExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)

	// 704:2: explicit_field_defs -> improper_explicit: ...
	ImproperExplicitToExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 705:2: explicit_field_defs -> nil: ...
	NilToExplicitFieldDefs() (GenericSymbol, error)
}

type ExplicitStructDefReducer interface {
	// 708:2: explicit_struct_def -> ...
	ToExplicitStructDef(Struct_ TokenValue, Lparen_ TokenValue, ExplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type EnumValueDefReducer interface {
	// 716:2: enum_value_def -> field_def: ...
	FieldDefToEnumValueDef(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 717:2: enum_value_def -> default: ...
	DefaultToEnumValueDef(FieldDef_ GenericSymbol, Assign_ TokenValue, Default_ TokenValue) (GenericSymbol, error)
}

type ImplicitEnumValueDefsReducer interface {
	// 729:2: implicit_enum_value_defs -> pair: ...
	PairToImplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Or_ TokenValue, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 730:2: implicit_enum_value_defs -> add: ...
	AddToImplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Or_ TokenValue, EnumValueDef_ GenericSymbol) (GenericSymbol, error)
}

type ImplicitEnumDefReducer interface {
	// 732:37: implicit_enum_def -> ...
	ToImplicitEnumDef(Lparen_ TokenValue, ImplicitEnumValueDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ExplicitEnumValueDefsReducer interface {
	// 735:2: explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToExplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Or_ TokenValue, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 736:2: explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToExplicitEnumValueDefs(EnumValueDef_ GenericSymbol, Newlines_ TokenCount, EnumValueDef_2 GenericSymbol) (GenericSymbol, error)

	// 737:2: explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Or_ TokenValue, EnumValueDef_ GenericSymbol) (GenericSymbol, error)

	// 738:2: explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToExplicitEnumValueDefs(ImplicitEnumValueDefs_ GenericSymbol, Newlines_ TokenCount, EnumValueDef_ GenericSymbol) (GenericSymbol, error)
}

type ExplicitEnumDefReducer interface {
	// 740:37: explicit_enum_def -> ...
	ToExplicitEnumDef(Enum_ TokenValue, Lparen_ TokenValue, ExplicitEnumValueDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type TraitPropertyReducer interface {
	// 747:2: trait_property -> field_def: ...
	FieldDefToTraitProperty(FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 748:2: trait_property -> method_signature: ...
	MethodSignatureToTraitProperty(MethodSignature_ GenericSymbol) (GenericSymbol, error)
}

type ProperTraitPropertiesReducer interface {
	// 751:2: proper_trait_properties -> implicit: ...
	ImplicitToProperTraitProperties(ProperTraitProperties_ GenericSymbol, Newlines_ TokenCount, TraitProperty_ GenericSymbol) (GenericSymbol, error)

	// 752:2: proper_trait_properties -> explicit: ...
	ExplicitToProperTraitProperties(ProperTraitProperties_ GenericSymbol, Comma_ TokenValue, TraitProperty_ GenericSymbol) (GenericSymbol, error)

	// 753:2: proper_trait_properties -> trait_property: ...
	TraitPropertyToProperTraitProperties(TraitProperty_ GenericSymbol) (GenericSymbol, error)
}

type TraitPropertiesReducer interface {

	// 757:2: trait_properties -> improper_implicit: ...
	ImproperImplicitToTraitProperties(ProperTraitProperties_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)

	// 758:2: trait_properties -> improper_explicit: ...
	ImproperExplicitToTraitProperties(ProperTraitProperties_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 759:2: trait_properties -> nil: ...
	NilToTraitProperties() (GenericSymbol, error)
}

type TraitDefReducer interface {
	// 761:29: trait_def -> ...
	ToTraitDef(Trait_ TokenValue, Lparen_ TokenValue, TraitProperties_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ReturnTypeReducer interface {
	// 769:2: return_type -> returnable_type: ...
	ReturnableTypeToReturnType(ReturnableType_ TypeExpression) (TypeExpression, error)

	// 770:2: return_type -> nil: ...
	NilToReturnType() (TypeExpression, error)
}

type ParameterDeclReducer interface {
	// 773:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 774:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ TokenValue, Ellipsis_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 775:2: parameter_decl -> unamed: ...
	UnamedToParameterDecl(ValueType_ TypeExpression) (GenericSymbol, error)

	// 776:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(Ellipsis_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)
}

type ProperParameterDeclsReducer interface {
	// 779:2: proper_parameter_decls -> add: ...
	AddToProperParameterDecls(ProperParameterDecls_ GenericSymbol, Comma_ TokenValue, ParameterDecl_ GenericSymbol) (GenericSymbol, error)

	// 780:2: proper_parameter_decls -> parameter_decl: ...
	ParameterDeclToProperParameterDecls(ParameterDecl_ GenericSymbol) (GenericSymbol, error)
}

type ParameterDeclsReducer interface {

	// 784:2: parameter_decls -> improper: ...
	ImproperToParameterDecls(ProperParameterDecls_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 785:2: parameter_decls -> nil: ...
	NilToParameterDecls() (GenericSymbol, error)
}

type FuncTypeReducer interface {
	// 788:2: func_type -> ...
	ToFuncType(Func_ TokenValue, Lparen_ TokenValue, ParameterDecls_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression) (TypeExpression, error)
}

type MethodSignatureReducer interface {
	// 799:20: method_signature -> ...
	ToMethodSignature(Func_ TokenValue, Identifier_ TokenValue, Lparen_ TokenValue, ParameterDecls_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression) (GenericSymbol, error)
}

type ParameterDefReducer interface {
	// 805:2: parameter_def -> inferred_ref_arg: ...
	InferredRefArgToParameterDef(Identifier_ TokenValue) (GenericSymbol, error)

	// 806:2: parameter_def -> inferred_ref_vararg: ...
	InferredRefVarargToParameterDef(Identifier_ TokenValue, Ellipsis_ TokenValue) (GenericSymbol, error)

	// 807:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)

	// 808:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ TokenValue, Ellipsis_ TokenValue, ValueType_ TypeExpression) (GenericSymbol, error)
}

type ProperParameterDefsReducer interface {
	// 811:2: proper_parameter_defs -> add: ...
	AddToProperParameterDefs(ProperParameterDefs_ GenericSymbol, Comma_ TokenValue, ParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 812:2: proper_parameter_defs -> parameter_def: ...
	ParameterDefToProperParameterDefs(ParameterDef_ GenericSymbol) (GenericSymbol, error)
}

type ParameterDefsReducer interface {

	// 816:2: parameter_defs -> improper: ...
	ImproperToParameterDefs(ProperParameterDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 817:2: parameter_defs -> nil: ...
	NilToParameterDefs() (GenericSymbol, error)
}

type NamedFuncDefReducer interface {
	// 820:2: named_func_def -> func_def: ...
	FuncDefToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, Lparen_ TokenValue, ParameterDefs_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 821:2: named_func_def -> method_def: ...
	MethodDefToNamedFuncDef(Func_ TokenValue, Lparen_ TokenValue, ParameterDef_ GenericSymbol, Rparen_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, Lparen_2 TokenValue, ParameterDefs_ GenericSymbol, Rparen_2 TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 822:2: named_func_def -> alias: ...
	AliasToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, Expression_ Expression) (SourceDefinition, error)
}

type AnonymousFuncExprReducer interface {
	// 826:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ TokenValue, Lparen_ TokenValue, ParameterDefs_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (Expression, error)
}

type PackageDefReducer interface {
	// 838:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ TokenValue) (SourceDefinition, error)

	// 839:2: package_def -> with_spec: ...
	WithSpecToPackageDef(Package_ TokenValue, StatementBlock_ GenericSymbol) (SourceDefinition, error)
}

type Reducer interface {
	SourceReducer
	ProperDefinitionsReducer
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

func ParseStatement(lexer Lexer, reducer Reducer) (Statement, error) {

	return ParseStatementWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseStatementWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	Statement,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State2)
	if err != nil {
		var errRetVal Statement
		return errRetVal, err
	}
	return item.Statement, nil
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

	item, err := _Parse(lexer, reducer, errHandler, _State3)
	if err != nil {
		var errRetVal Expression
		return errRetVal, err
	}
	return item.Expression, nil
}

func ParseValueType(lexer Lexer, reducer Reducer) (TypeExpression, error) {

	return ParseValueTypeWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseValueTypeWithCustomErrorHandler(
	lexer Lexer,
	reducer Reducer,
	errHandler ParseErrorHandler) (
	TypeExpression,
	error) {

	item, err := _Parse(lexer, reducer, errHandler, _State4)
	if err != nil {
		var errRetVal TypeExpression
		return errRetVal, err
	}
	return item.TypeExpression, nil
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
	case ProperDefinitionsType:
		return "proper_definitions"
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
	ProperDefinitionsType                   = SymbolId(344)
	DefinitionsType                         = SymbolId(345)
	DefinitionType                          = SymbolId(346)
	StatementBlockType                      = SymbolId(347)
	ProperStatementsType                    = SymbolId(348)
	StatementsType                          = SymbolId(349)
	SimpleStatementType                     = SymbolId(350)
	StatementType                           = SymbolId(351)
	OptionalSimpleStatementType             = SymbolId(352)
	ExpressionOrImproperStructStatementType = SymbolId(353)
	ExpressionsType                         = SymbolId(354)
	CallbackOpType                          = SymbolId(355)
	CallbackOpStatementType                 = SymbolId(356)
	UnsafeStatementType                     = SymbolId(357)
	JumpStatementType                       = SymbolId(358)
	JumpTypeType                            = SymbolId(359)
	FallthroughStatementType                = SymbolId(360)
	AssignStatementType                     = SymbolId(361)
	UnaryOpAssignStatementType              = SymbolId(362)
	UnaryOpAssignType                       = SymbolId(363)
	BinaryOpAssignStatementType             = SymbolId(364)
	BinaryOpAssignType                      = SymbolId(365)
	ImportStatementType                     = SymbolId(366)
	ImportClauseType                        = SymbolId(367)
	ProperImportClausesType                 = SymbolId(368)
	ImportClausesType                       = SymbolId(369)
	CasePatternsType                        = SymbolId(370)
	VarDeclPatternType                      = SymbolId(371)
	VarOrLetType                            = SymbolId(372)
	VarPatternType                          = SymbolId(373)
	TuplePatternType                        = SymbolId(374)
	FieldVarPatternsType                    = SymbolId(375)
	FieldVarPatternType                     = SymbolId(376)
	OptionalValueTypeType                   = SymbolId(377)
	AssignPatternType                       = SymbolId(378)
	CasePatternType                         = SymbolId(379)
	ExpressionType                          = SymbolId(380)
	OptionalLabelDeclType                   = SymbolId(381)
	SequenceExprType                        = SymbolId(382)
	IfExprType                              = SymbolId(383)
	ConditionType                           = SymbolId(384)
	SwitchExprType                          = SymbolId(385)
	LoopExprType                            = SymbolId(386)
	OptionalSequenceExprType                = SymbolId(387)
	ForAssignmentType                       = SymbolId(388)
	CallExprType                            = SymbolId(389)
	OptionalGenericBindingType              = SymbolId(390)
	ProperGenericArgumentsType              = SymbolId(391)
	GenericArgumentsType                    = SymbolId(392)
	ProperArgumentsType                     = SymbolId(393)
	ArgumentsType                           = SymbolId(394)
	ArgumentType                            = SymbolId(395)
	ColonExprType                           = SymbolId(396)
	AtomExprType                            = SymbolId(397)
	ParseErrorExprType                      = SymbolId(398)
	LiteralExprType                         = SymbolId(399)
	IdentifierExprType                      = SymbolId(400)
	BlockExprType                           = SymbolId(401)
	InitializeExprType                      = SymbolId(402)
	ImplicitStructExprType                  = SymbolId(403)
	AccessibleExprType                      = SymbolId(404)
	AccessExprType                          = SymbolId(405)
	IndexExprType                           = SymbolId(406)
	PostfixableExprType                     = SymbolId(407)
	PostfixUnaryExprType                    = SymbolId(408)
	PrefixableExprType                      = SymbolId(409)
	PrefixUnaryExprType                     = SymbolId(410)
	PrefixUnaryOpType                       = SymbolId(411)
	MulExprType                             = SymbolId(412)
	BinaryMulExprType                       = SymbolId(413)
	MulOpType                               = SymbolId(414)
	AddExprType                             = SymbolId(415)
	BinaryAddExprType                       = SymbolId(416)
	AddOpType                               = SymbolId(417)
	CmpExprType                             = SymbolId(418)
	BinaryCmpExprType                       = SymbolId(419)
	CmpOpType                               = SymbolId(420)
	AndExprType                             = SymbolId(421)
	BinaryAndExprType                       = SymbolId(422)
	OrExprType                              = SymbolId(423)
	BinaryOrExprType                        = SymbolId(424)
	InitializableTypeType                   = SymbolId(425)
	SliceTypeType                           = SymbolId(426)
	ArrayTypeType                           = SymbolId(427)
	MapTypeType                             = SymbolId(428)
	AtomTypeType                            = SymbolId(429)
	ParseErrorTypeType                      = SymbolId(430)
	ReturnableTypeType                      = SymbolId(431)
	PrefixedTypeType                        = SymbolId(432)
	PrefixTypeOpType                        = SymbolId(433)
	ValueTypeType                           = SymbolId(434)
	TraitOpTypeType                         = SymbolId(435)
	TraitOpType                             = SymbolId(436)
	TypeDefType                             = SymbolId(437)
	GenericParameterDefType                 = SymbolId(438)
	ProperGenericParameterDefsType          = SymbolId(439)
	GenericParameterDefsType                = SymbolId(440)
	OptionalGenericParametersType           = SymbolId(441)
	FieldDefType                            = SymbolId(442)
	ProperImplicitFieldDefsType             = SymbolId(443)
	ImplicitFieldDefsType                   = SymbolId(444)
	ImplicitStructDefType                   = SymbolId(445)
	ProperExplicitFieldDefsType             = SymbolId(446)
	ExplicitFieldDefsType                   = SymbolId(447)
	ExplicitStructDefType                   = SymbolId(448)
	EnumValueDefType                        = SymbolId(449)
	ImplicitEnumValueDefsType               = SymbolId(450)
	ImplicitEnumDefType                     = SymbolId(451)
	ExplicitEnumValueDefsType               = SymbolId(452)
	ExplicitEnumDefType                     = SymbolId(453)
	TraitPropertyType                       = SymbolId(454)
	ProperTraitPropertiesType               = SymbolId(455)
	TraitPropertiesType                     = SymbolId(456)
	TraitDefType                            = SymbolId(457)
	ReturnTypeType                          = SymbolId(458)
	ParameterDeclType                       = SymbolId(459)
	ProperParameterDeclsType                = SymbolId(460)
	ParameterDeclsType                      = SymbolId(461)
	FuncTypeType                            = SymbolId(462)
	MethodSignatureType                     = SymbolId(463)
	ParameterDefType                        = SymbolId(464)
	ProperParameterDefsType                 = SymbolId(465)
	ParameterDefsType                       = SymbolId(466)
	NamedFuncDefType                        = SymbolId(467)
	AnonymousFuncExprType                   = SymbolId(468)
	PackageDefType                          = SymbolId(469)
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
	_ReduceAddToProperDefinitions                               = _ReduceType(2)
	_ReduceDefinitionToProperDefinitions                        = _ReduceType(3)
	_ReduceProperDefinitionsToDefinitions                       = _ReduceType(4)
	_ReduceImproperToDefinitions                                = _ReduceType(5)
	_ReduceNilToDefinitions                                     = _ReduceType(6)
	_ReducePackageDefToDefinition                               = _ReduceType(7)
	_ReduceTypeDefToDefinition                                  = _ReduceType(8)
	_ReduceNamedFuncDefToDefinition                             = _ReduceType(9)
	_ReduceGlobalVarDefToDefinition                             = _ReduceType(10)
	_ReduceGlobalVarAssignmentToDefinition                      = _ReduceType(11)
	_ReduceStatementBlockToDefinition                           = _ReduceType(12)
	_ReduceCommentGroupsToDefinition                            = _ReduceType(13)
	_ReduceToStatementBlock                                     = _ReduceType(14)
	_ReduceAddImplicitToProperStatements                        = _ReduceType(15)
	_ReduceAddExplicitToProperStatements                        = _ReduceType(16)
	_ReduceStatementToProperStatements                          = _ReduceType(17)
	_ReduceProperStatementsToStatements                         = _ReduceType(18)
	_ReduceImproperImplicitToStatements                         = _ReduceType(19)
	_ReduceImproperExplicitToStatements                         = _ReduceType(20)
	_ReduceNilToStatements                                      = _ReduceType(21)
	_ReduceUnsafeStatementToSimpleStatement                     = _ReduceType(22)
	_ReduceExpressionOrImproperStructStatementToSimpleStatement = _ReduceType(23)
	_ReduceCallbackOpStatementToSimpleStatement                 = _ReduceType(24)
	_ReduceJumpStatementToSimpleStatement                       = _ReduceType(25)
	_ReduceAssignStatementToSimpleStatement                     = _ReduceType(26)
	_ReduceUnaryOpAssignStatementToSimpleStatement              = _ReduceType(27)
	_ReduceBinaryOpAssignStatementToSimpleStatement             = _ReduceType(28)
	_ReduceFallthroughStatementToSimpleStatement                = _ReduceType(29)
	_ReduceSimpleStatementToStatement                           = _ReduceType(30)
	_ReduceImportStatementToStatement                           = _ReduceType(31)
	_ReduceCaseBranchStatementToStatement                       = _ReduceType(32)
	_ReduceDefaultBranchStatementToStatement                    = _ReduceType(33)
	_ReduceSimpleStatementToOptionalSimpleStatement             = _ReduceType(34)
	_ReduceNilToOptionalSimpleStatement                         = _ReduceType(35)
	_ReduceToExpressionOrImproperStructStatement                = _ReduceType(36)
	_ReduceExpressionToExpressions                              = _ReduceType(37)
	_ReduceAddToExpressions                                     = _ReduceType(38)
	_ReduceAsyncToCallbackOp                                    = _ReduceType(39)
	_ReduceDeferToCallbackOp                                    = _ReduceType(40)
	_ReduceToCallbackOpStatement                                = _ReduceType(41)
	_ReduceToUnsafeStatement                                    = _ReduceType(42)
	_ReduceUnlabeledNoValueToJumpStatement                      = _ReduceType(43)
	_ReduceUnlabeledValuedToJumpStatement                       = _ReduceType(44)
	_ReduceLabeledNoValueToJumpStatement                        = _ReduceType(45)
	_ReduceLabeledValuedToJumpStatement                         = _ReduceType(46)
	_ReduceReturnToJumpType                                     = _ReduceType(47)
	_ReduceBreakToJumpType                                      = _ReduceType(48)
	_ReduceContinueToJumpType                                   = _ReduceType(49)
	_ReduceToFallthroughStatement                               = _ReduceType(50)
	_ReduceToAssignStatement                                    = _ReduceType(51)
	_ReduceToUnaryOpAssignStatement                             = _ReduceType(52)
	_ReduceAddOneAssignToUnaryOpAssign                          = _ReduceType(53)
	_ReduceSubOneAssignToUnaryOpAssign                          = _ReduceType(54)
	_ReduceToBinaryOpAssignStatement                            = _ReduceType(55)
	_ReduceAddAssignToBinaryOpAssign                            = _ReduceType(56)
	_ReduceSubAssignToBinaryOpAssign                            = _ReduceType(57)
	_ReduceMulAssignToBinaryOpAssign                            = _ReduceType(58)
	_ReduceDivAssignToBinaryOpAssign                            = _ReduceType(59)
	_ReduceModAssignToBinaryOpAssign                            = _ReduceType(60)
	_ReduceBitNegAssignToBinaryOpAssign                         = _ReduceType(61)
	_ReduceBitAndAssignToBinaryOpAssign                         = _ReduceType(62)
	_ReduceBitOrAssignToBinaryOpAssign                          = _ReduceType(63)
	_ReduceBitXorAssignToBinaryOpAssign                         = _ReduceType(64)
	_ReduceBitLshiftAssignToBinaryOpAssign                      = _ReduceType(65)
	_ReduceBitRshiftAssignToBinaryOpAssign                      = _ReduceType(66)
	_ReduceSingleToImportStatement                              = _ReduceType(67)
	_ReduceMultipleToImportStatement                            = _ReduceType(68)
	_ReduceStringLiteralToImportClause                          = _ReduceType(69)
	_ReduceAliasToImportClause                                  = _ReduceType(70)
	_ReduceAddImplicitToProperImportClauses                     = _ReduceType(71)
	_ReduceAddExplicitToProperImportClauses                     = _ReduceType(72)
	_ReduceImportClauseToProperImportClauses                    = _ReduceType(73)
	_ReduceProperImportClausesToImportClauses                   = _ReduceType(74)
	_ReduceImplicitToImportClauses                              = _ReduceType(75)
	_ReduceExplicitToImportClauses                              = _ReduceType(76)
	_ReduceCasePatternToCasePatterns                            = _ReduceType(77)
	_ReduceMultipleToCasePatterns                               = _ReduceType(78)
	_ReduceToVarDeclPattern                                     = _ReduceType(79)
	_ReduceVarToVarOrLet                                        = _ReduceType(80)
	_ReduceLetToVarOrLet                                        = _ReduceType(81)
	_ReduceIdentifierToVarPattern                               = _ReduceType(82)
	_ReduceTuplePatternToVarPattern                             = _ReduceType(83)
	_ReduceToTuplePattern                                       = _ReduceType(84)
	_ReduceFieldVarPatternToFieldVarPatterns                    = _ReduceType(85)
	_ReduceAddToFieldVarPatterns                                = _ReduceType(86)
	_ReducePositionalToFieldVarPattern                          = _ReduceType(87)
	_ReduceNamedToFieldVarPattern                               = _ReduceType(88)
	_ReduceEllipsisToFieldVarPattern                            = _ReduceType(89)
	_ReduceValueTypeToOptionalValueType                         = _ReduceType(90)
	_ReduceNilToOptionalValueType                               = _ReduceType(91)
	_ReduceToAssignPattern                                      = _ReduceType(92)
	_ReduceAssignPatternToCasePattern                           = _ReduceType(93)
	_ReduceEnumMatchPatternToCasePattern                        = _ReduceType(94)
	_ReduceEnumVarDeclPatternToCasePattern                      = _ReduceType(95)
	_ReduceIfExprToExpression                                   = _ReduceType(96)
	_ReduceSwitchExprToExpression                               = _ReduceType(97)
	_ReduceLoopExprToExpression                                 = _ReduceType(98)
	_ReduceSequenceExprToExpression                             = _ReduceType(99)
	_ReduceLabelDeclToOptionalLabelDecl                         = _ReduceType(100)
	_ReduceUnlabelledToOptionalLabelDecl                        = _ReduceType(101)
	_ReduceOrExprToSequenceExpr                                 = _ReduceType(102)
	_ReduceVarDeclPatternToSequenceExpr                         = _ReduceType(103)
	_ReduceAssignVarPatternToSequenceExpr                       = _ReduceType(104)
	_ReduceNoElseToIfExpr                                       = _ReduceType(105)
	_ReduceIfElseToIfExpr                                       = _ReduceType(106)
	_ReduceMultiIfElseToIfExpr                                  = _ReduceType(107)
	_ReduceSequenceExprToCondition                              = _ReduceType(108)
	_ReduceCaseToCondition                                      = _ReduceType(109)
	_ReduceToSwitchExpr                                         = _ReduceType(110)
	_ReduceInfiniteToLoopExpr                                   = _ReduceType(111)
	_ReduceDoWhileToLoopExpr                                    = _ReduceType(112)
	_ReduceWhileToLoopExpr                                      = _ReduceType(113)
	_ReduceIteratorToLoopExpr                                   = _ReduceType(114)
	_ReduceForToLoopExpr                                        = _ReduceType(115)
	_ReduceSequenceExprToOptionalSequenceExpr                   = _ReduceType(116)
	_ReduceNilToOptionalSequenceExpr                            = _ReduceType(117)
	_ReduceSequenceExprToForAssignment                          = _ReduceType(118)
	_ReduceAssignToForAssignment                                = _ReduceType(119)
	_ReduceToCallExpr                                           = _ReduceType(120)
	_ReduceBindingToOptionalGenericBinding                      = _ReduceType(121)
	_ReduceNilToOptionalGenericBinding                          = _ReduceType(122)
	_ReduceAddToProperGenericArguments                          = _ReduceType(123)
	_ReduceValueTypeToProperGenericArguments                    = _ReduceType(124)
	_ReduceProperGenericArgumentsToGenericArguments             = _ReduceType(125)
	_ReduceImproperToGenericArguments                           = _ReduceType(126)
	_ReduceNilToGenericArguments                                = _ReduceType(127)
	_ReduceAddToProperArguments                                 = _ReduceType(128)
	_ReduceArgumentToProperArguments                            = _ReduceType(129)
	_ReduceProperArgumentsToArguments                           = _ReduceType(130)
	_ReduceImproperToArguments                                  = _ReduceType(131)
	_ReduceNilToArguments                                       = _ReduceType(132)
	_ReducePositionalToArgument                                 = _ReduceType(133)
	_ReduceColonExprToArgument                                  = _ReduceType(134)
	_ReduceNamedAssignmentToArgument                            = _ReduceType(135)
	_ReduceVarargAssignmentToArgument                           = _ReduceType(136)
	_ReduceSkipPatternToArgument                                = _ReduceType(137)
	_ReduceUnitUnitPairToColonExpr                              = _ReduceType(138)
	_ReduceExprUnitPairToColonExpr                              = _ReduceType(139)
	_ReduceUnitExprPairToColonExpr                              = _ReduceType(140)
	_ReduceExprExprPairToColonExpr                              = _ReduceType(141)
	_ReduceColonExprUnitTupleToColonExpr                        = _ReduceType(142)
	_ReduceColonExprExprTupleToColonExpr                        = _ReduceType(143)
	_ReduceParseErrorExprToAtomExpr                             = _ReduceType(144)
	_ReduceLiteralExprToAtomExpr                                = _ReduceType(145)
	_ReduceIdentifierExprToAtomExpr                             = _ReduceType(146)
	_ReduceBlockExprToAtomExpr                                  = _ReduceType(147)
	_ReduceAnonymousFuncExprToAtomExpr                          = _ReduceType(148)
	_ReduceInitializeExprToAtomExpr                             = _ReduceType(149)
	_ReduceImplicitStructExprToAtomExpr                         = _ReduceType(150)
	_ReduceToParseErrorExpr                                     = _ReduceType(151)
	_ReduceTrueToLiteralExpr                                    = _ReduceType(152)
	_ReduceFalseToLiteralExpr                                   = _ReduceType(153)
	_ReduceIntegerLiteralToLiteralExpr                          = _ReduceType(154)
	_ReduceFloatLiteralToLiteralExpr                            = _ReduceType(155)
	_ReduceRuneLiteralToLiteralExpr                             = _ReduceType(156)
	_ReduceStringLiteralToLiteralExpr                           = _ReduceType(157)
	_ReduceToIdentifierExpr                                     = _ReduceType(158)
	_ReduceToBlockExpr                                          = _ReduceType(159)
	_ReduceToInitializeExpr                                     = _ReduceType(160)
	_ReduceToImplicitStructExpr                                 = _ReduceType(161)
	_ReduceAtomExprToAccessibleExpr                             = _ReduceType(162)
	_ReduceAccessExprToAccessibleExpr                           = _ReduceType(163)
	_ReduceCallExprToAccessibleExpr                             = _ReduceType(164)
	_ReduceIndexExprToAccessibleExpr                            = _ReduceType(165)
	_ReduceToAccessExpr                                         = _ReduceType(166)
	_ReduceToIndexExpr                                          = _ReduceType(167)
	_ReduceAccessibleExprToPostfixableExpr                      = _ReduceType(168)
	_ReducePostfixUnaryExprToPostfixableExpr                    = _ReduceType(169)
	_ReduceToPostfixUnaryExpr                                   = _ReduceType(170)
	_ReducePostfixableExprToPrefixableExpr                      = _ReduceType(171)
	_ReducePrefixUnaryExprToPrefixableExpr                      = _ReduceType(172)
	_ReduceToPrefixUnaryExpr                                    = _ReduceType(173)
	_ReduceNotToPrefixUnaryOp                                   = _ReduceType(174)
	_ReduceBitNegToPrefixUnaryOp                                = _ReduceType(175)
	_ReduceSubToPrefixUnaryOp                                   = _ReduceType(176)
	_ReduceMulToPrefixUnaryOp                                   = _ReduceType(177)
	_ReduceBitAndToPrefixUnaryOp                                = _ReduceType(178)
	_ReducePrefixableExprToMulExpr                              = _ReduceType(179)
	_ReduceBinaryMulExprToMulExpr                               = _ReduceType(180)
	_ReduceToBinaryMulExpr                                      = _ReduceType(181)
	_ReduceMulToMulOp                                           = _ReduceType(182)
	_ReduceDivToMulOp                                           = _ReduceType(183)
	_ReduceModToMulOp                                           = _ReduceType(184)
	_ReduceBitAndToMulOp                                        = _ReduceType(185)
	_ReduceBitLshiftToMulOp                                     = _ReduceType(186)
	_ReduceBitRshiftToMulOp                                     = _ReduceType(187)
	_ReduceMulExprToAddExpr                                     = _ReduceType(188)
	_ReduceBinaryAddExprToAddExpr                               = _ReduceType(189)
	_ReduceToBinaryAddExpr                                      = _ReduceType(190)
	_ReduceAddToAddOp                                           = _ReduceType(191)
	_ReduceSubToAddOp                                           = _ReduceType(192)
	_ReduceBitOrToAddOp                                         = _ReduceType(193)
	_ReduceBitXorToAddOp                                        = _ReduceType(194)
	_ReduceAddExprToCmpExpr                                     = _ReduceType(195)
	_ReduceBinaryCmpExprToCmpExpr                               = _ReduceType(196)
	_ReduceToBinaryCmpExpr                                      = _ReduceType(197)
	_ReduceEqualToCmpOp                                         = _ReduceType(198)
	_ReduceNotEqualToCmpOp                                      = _ReduceType(199)
	_ReduceLessToCmpOp                                          = _ReduceType(200)
	_ReduceLessOrEqualToCmpOp                                   = _ReduceType(201)
	_ReduceGreaterToCmpOp                                       = _ReduceType(202)
	_ReduceGreaterOrEqualToCmpOp                                = _ReduceType(203)
	_ReduceCmpExprToAndExpr                                     = _ReduceType(204)
	_ReduceBinaryAndExprToAndExpr                               = _ReduceType(205)
	_ReduceToBinaryAndExpr                                      = _ReduceType(206)
	_ReduceAndExprToOrExpr                                      = _ReduceType(207)
	_ReduceBinaryOrExprToOrExpr                                 = _ReduceType(208)
	_ReduceToBinaryOrExpr                                       = _ReduceType(209)
	_ReduceExplicitStructDefToInitializableType                 = _ReduceType(210)
	_ReduceSliceTypeToInitializableType                         = _ReduceType(211)
	_ReduceArrayTypeToInitializableType                         = _ReduceType(212)
	_ReduceMapTypeToInitializableType                           = _ReduceType(213)
	_ReduceToSliceType                                          = _ReduceType(214)
	_ReduceToArrayType                                          = _ReduceType(215)
	_ReduceToMapType                                            = _ReduceType(216)
	_ReduceInitializableTypeToAtomType                          = _ReduceType(217)
	_ReduceNamedToAtomType                                      = _ReduceType(218)
	_ReduceExternNamedToAtomType                                = _ReduceType(219)
	_ReduceInferredToAtomType                                   = _ReduceType(220)
	_ReduceImplicitStructDefToAtomType                          = _ReduceType(221)
	_ReduceExplicitEnumDefToAtomType                            = _ReduceType(222)
	_ReduceImplicitEnumDefToAtomType                            = _ReduceType(223)
	_ReduceTraitDefToAtomType                                   = _ReduceType(224)
	_ReduceFuncTypeToAtomType                                   = _ReduceType(225)
	_ReduceParseErrorTypeToAtomType                             = _ReduceType(226)
	_ReduceToParseErrorType                                     = _ReduceType(227)
	_ReduceAtomTypeToReturnableType                             = _ReduceType(228)
	_ReducePrefixedTypeToReturnableType                         = _ReduceType(229)
	_ReduceToPrefixedType                                       = _ReduceType(230)
	_ReduceQuestionToPrefixTypeOp                               = _ReduceType(231)
	_ReduceExclaimToPrefixTypeOp                                = _ReduceType(232)
	_ReduceBitAndToPrefixTypeOp                                 = _ReduceType(233)
	_ReduceBitNegToPrefixTypeOp                                 = _ReduceType(234)
	_ReduceTildeTildeToPrefixTypeOp                             = _ReduceType(235)
	_ReduceReturnableTypeToValueType                            = _ReduceType(236)
	_ReduceTraitOpTypeToValueType                               = _ReduceType(237)
	_ReduceToTraitOpType                                        = _ReduceType(238)
	_ReduceMulToTraitOp                                         = _ReduceType(239)
	_ReduceAddToTraitOp                                         = _ReduceType(240)
	_ReduceSubToTraitOp                                         = _ReduceType(241)
	_ReduceDefinitionToTypeDef                                  = _ReduceType(242)
	_ReduceConstrainedDefToTypeDef                              = _ReduceType(243)
	_ReduceAliasToTypeDef                                       = _ReduceType(244)
	_ReduceUnconstrainedToGenericParameterDef                   = _ReduceType(245)
	_ReduceConstrainedToGenericParameterDef                     = _ReduceType(246)
	_ReduceAddToProperGenericParameterDefs                      = _ReduceType(247)
	_ReduceGenericParameterDefToProperGenericParameterDefs      = _ReduceType(248)
	_ReduceProperGenericParameterDefsToGenericParameterDefs     = _ReduceType(249)
	_ReduceImproperToGenericParameterDefs                       = _ReduceType(250)
	_ReduceNilToGenericParameterDefs                            = _ReduceType(251)
	_ReduceGenericToOptionalGenericParameters                   = _ReduceType(252)
	_ReduceNilToOptionalGenericParameters                       = _ReduceType(253)
	_ReduceExplicitToFieldDef                                   = _ReduceType(254)
	_ReduceImplicitToFieldDef                                   = _ReduceType(255)
	_ReduceUnsafeStatementToFieldDef                            = _ReduceType(256)
	_ReduceAddToProperImplicitFieldDefs                         = _ReduceType(257)
	_ReduceFieldDefToProperImplicitFieldDefs                    = _ReduceType(258)
	_ReduceProperImplicitFieldDefsToImplicitFieldDefs           = _ReduceType(259)
	_ReduceImproperToImplicitFieldDefs                          = _ReduceType(260)
	_ReduceNilToImplicitFieldDefs                               = _ReduceType(261)
	_ReduceToImplicitStructDef                                  = _ReduceType(262)
	_ReduceAddImplicitToProperExplicitFieldDefs                 = _ReduceType(263)
	_ReduceAddExplicitToProperExplicitFieldDefs                 = _ReduceType(264)
	_ReduceFieldDefToProperExplicitFieldDefs                    = _ReduceType(265)
	_ReduceProperExplicitFieldDefsToExplicitFieldDefs           = _ReduceType(266)
	_ReduceImproperImplicitToExplicitFieldDefs                  = _ReduceType(267)
	_ReduceImproperExplicitToExplicitFieldDefs                  = _ReduceType(268)
	_ReduceNilToExplicitFieldDefs                               = _ReduceType(269)
	_ReduceToExplicitStructDef                                  = _ReduceType(270)
	_ReduceFieldDefToEnumValueDef                               = _ReduceType(271)
	_ReduceDefaultToEnumValueDef                                = _ReduceType(272)
	_ReducePairToImplicitEnumValueDefs                          = _ReduceType(273)
	_ReduceAddToImplicitEnumValueDefs                           = _ReduceType(274)
	_ReduceToImplicitEnumDef                                    = _ReduceType(275)
	_ReduceExplicitPairToExplicitEnumValueDefs                  = _ReduceType(276)
	_ReduceImplicitPairToExplicitEnumValueDefs                  = _ReduceType(277)
	_ReduceExplicitAddToExplicitEnumValueDefs                   = _ReduceType(278)
	_ReduceImplicitAddToExplicitEnumValueDefs                   = _ReduceType(279)
	_ReduceToExplicitEnumDef                                    = _ReduceType(280)
	_ReduceFieldDefToTraitProperty                              = _ReduceType(281)
	_ReduceMethodSignatureToTraitProperty                       = _ReduceType(282)
	_ReduceImplicitToProperTraitProperties                      = _ReduceType(283)
	_ReduceExplicitToProperTraitProperties                      = _ReduceType(284)
	_ReduceTraitPropertyToProperTraitProperties                 = _ReduceType(285)
	_ReduceProperTraitPropertiesToTraitProperties               = _ReduceType(286)
	_ReduceImproperImplicitToTraitProperties                    = _ReduceType(287)
	_ReduceImproperExplicitToTraitProperties                    = _ReduceType(288)
	_ReduceNilToTraitProperties                                 = _ReduceType(289)
	_ReduceToTraitDef                                           = _ReduceType(290)
	_ReduceReturnableTypeToReturnType                           = _ReduceType(291)
	_ReduceNilToReturnType                                      = _ReduceType(292)
	_ReduceArgToParameterDecl                                   = _ReduceType(293)
	_ReduceVarargToParameterDecl                                = _ReduceType(294)
	_ReduceUnamedToParameterDecl                                = _ReduceType(295)
	_ReduceUnnamedVarargToParameterDecl                         = _ReduceType(296)
	_ReduceAddToProperParameterDecls                            = _ReduceType(297)
	_ReduceParameterDeclToProperParameterDecls                  = _ReduceType(298)
	_ReduceProperParameterDeclsToParameterDecls                 = _ReduceType(299)
	_ReduceImproperToParameterDecls                             = _ReduceType(300)
	_ReduceNilToParameterDecls                                  = _ReduceType(301)
	_ReduceToFuncType                                           = _ReduceType(302)
	_ReduceToMethodSignature                                    = _ReduceType(303)
	_ReduceInferredRefArgToParameterDef                         = _ReduceType(304)
	_ReduceInferredRefVarargToParameterDef                      = _ReduceType(305)
	_ReduceArgToParameterDef                                    = _ReduceType(306)
	_ReduceVarargToParameterDef                                 = _ReduceType(307)
	_ReduceAddToProperParameterDefs                             = _ReduceType(308)
	_ReduceParameterDefToProperParameterDefs                    = _ReduceType(309)
	_ReduceProperParameterDefsToParameterDefs                   = _ReduceType(310)
	_ReduceImproperToParameterDefs                              = _ReduceType(311)
	_ReduceNilToParameterDefs                                   = _ReduceType(312)
	_ReduceFuncDefToNamedFuncDef                                = _ReduceType(313)
	_ReduceMethodDefToNamedFuncDef                              = _ReduceType(314)
	_ReduceAliasToNamedFuncDef                                  = _ReduceType(315)
	_ReduceToAnonymousFuncExpr                                  = _ReduceType(316)
	_ReduceNoSpecToPackageDef                                   = _ReduceType(317)
	_ReduceWithSpecToPackageDef                                 = _ReduceType(318)
)

func (i _ReduceType) String() string {
	switch i {
	case _ReduceToSource:
		return "ToSource"
	case _ReduceAddToProperDefinitions:
		return "AddToProperDefinitions"
	case _ReduceDefinitionToProperDefinitions:
		return "DefinitionToProperDefinitions"
	case _ReduceProperDefinitionsToDefinitions:
		return "ProperDefinitionsToDefinitions"
	case _ReduceImproperToDefinitions:
		return "ImproperToDefinitions"
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
	case SourceType, ProperDefinitionsType, DefinitionsType:
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
	case SourceType, ProperDefinitionsType, DefinitionsType:
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
	case _ReduceAddToProperDefinitions:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperDefinitionsType
		symbol.SourceDefinitions, err = reducer.AddToProperDefinitions(args[0].SourceDefinitions, args[1].Count, args[2].SourceDefinition)
	case _ReduceDefinitionToProperDefinitions:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperDefinitionsType
		symbol.SourceDefinitions, err = reducer.DefinitionToProperDefinitions(args[0].SourceDefinition)
	case _ReduceProperDefinitionsToDefinitions:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionsType
		//line grammar.lr:59:4
		symbol.SourceDefinitions = args[0].SourceDefinitions
		err = nil
	case _ReduceImproperToDefinitions:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = DefinitionsType
		symbol.SourceDefinitions, err = reducer.ImproperToDefinitions(args[0].SourceDefinitions, args[1].Count)
	case _ReduceNilToDefinitions:
		symbol.SymbolId_ = DefinitionsType
		symbol.SourceDefinitions, err = reducer.NilToDefinitions()
	case _ReducePackageDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		//line grammar.lr:65:4
		symbol.SourceDefinition = args[0].SourceDefinition
		err = nil
	case _ReduceTypeDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		//line grammar.lr:66:4
		symbol.SourceDefinition = args[0].SourceDefinition
		err = nil
	case _ReduceNamedFuncDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		//line grammar.lr:67:4
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
		//line grammar.lr:104:4
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
		//line grammar.lr:110:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceExpressionOrImproperStructStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:112:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceCallbackOpStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:115:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceJumpStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:118:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceAssignStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:123:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceUnaryOpAssignStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:129:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceBinaryOpAssignStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:130:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceFallthroughStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:134:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceSimpleStatementToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:137:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceImportStatementToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:140:4
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
		//line grammar.lr:149:4
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
		//line grammar.lr:167:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDeferToCallbackOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CallbackOpType
		//line grammar.lr:168:4
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
		//line grammar.lr:207:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBreakToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		//line grammar.lr:208:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceContinueToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		//line grammar.lr:209:4
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
		//line grammar.lr:222:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubOneAssignToUnaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpAssignType
		//line grammar.lr:223:4
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
		//line grammar.lr:228:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:229:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:230:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDivAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:231:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceModAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:232:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:233:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:234:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitOrAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:235:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitXorAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:236:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitLshiftAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:237:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitRshiftAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:238:4
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
		//line grammar.lr:258:4
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
		//line grammar.lr:280:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLetToVarOrLet:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarOrLetType
		//line grammar.lr:281:4
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
		//line grammar.lr:299:4
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
		//line grammar.lr:310:4
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
		//line grammar.lr:328:4
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
		//line grammar.lr:339:4
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
		//line grammar.lr:429:4
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
		//line grammar.lr:438:4
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
		//line grammar.lr:472:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceLiteralExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:473:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceIdentifierExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:474:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBlockExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:475:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceAnonymousFuncExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:476:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceInitializeExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:477:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceImplicitStructExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:478:4
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
		//line grammar.lr:499:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceAccessExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:500:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceCallExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:501:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceIndexExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:502:4
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
		//line grammar.lr:511:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReducePostfixUnaryExprToPostfixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixableExprType
		//line grammar.lr:512:4
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
		//line grammar.lr:517:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReducePrefixUnaryExprToPrefixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixableExprType
		//line grammar.lr:518:4
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
		//line grammar.lr:523:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:524:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:525:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:528:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:531:4
		symbol.Value = args[0].Value
		err = nil
	case _ReducePrefixableExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		//line grammar.lr:534:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryMulExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		//line grammar.lr:535:4
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
		//line grammar.lr:540:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDivToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:541:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceModToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:542:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:543:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitLshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:544:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitRshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:545:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulExprToAddExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddExprType
		//line grammar.lr:548:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryAddExprToAddExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddExprType
		//line grammar.lr:549:4
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
		//line grammar.lr:554:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:555:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitOrToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:556:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitXorToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:557:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddExprToCmpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpExprType
		//line grammar.lr:560:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryCmpExprToCmpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpExprType
		//line grammar.lr:561:4
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
		//line grammar.lr:566:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceNotEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:567:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLessToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:568:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLessOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:569:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceGreaterToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:570:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceGreaterOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:571:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceCmpExprToAndExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AndExprType
		//line grammar.lr:574:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryAndExprToAndExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AndExprType
		//line grammar.lr:575:4
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
		//line grammar.lr:580:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryOrExprToOrExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OrExprType
		//line grammar.lr:581:4
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
		//line grammar.lr:592:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceSliceTypeToInitializableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeType
		//line grammar.lr:593:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceArrayTypeToInitializableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeType
		//line grammar.lr:594:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceMapTypeToInitializableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeType
		//line grammar.lr:595:4
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
		//line grammar.lr:606:4
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
		//line grammar.lr:610:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceExplicitEnumDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:611:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceImplicitEnumDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:612:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceTraitDefToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:613:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceFuncTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:614:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceParseErrorTypeToAtomType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeType
		//line grammar.lr:615:4
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
		//line grammar.lr:623:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReducePrefixedTypeToReturnableType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableTypeType
		//line grammar.lr:624:4
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
		//line grammar.lr:629:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceExclaimToPrefixTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixTypeOpType
		//line grammar.lr:630:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToPrefixTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixTypeOpType
		//line grammar.lr:631:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegToPrefixTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixTypeOpType
		//line grammar.lr:632:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceTildeTildeToPrefixTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixTypeOpType
		//line grammar.lr:633:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceReturnableTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		//line grammar.lr:638:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceTraitOpTypeToValueType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ValueTypeType
		//line grammar.lr:639:4
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
		//line grammar.lr:644:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddToTraitOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitOpType
		//line grammar.lr:645:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToTraitOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TraitOpType
		//line grammar.lr:646:4
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
		//line grammar.lr:667:4
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
		//line grammar.lr:689:4
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
		//line grammar.lr:702:4
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
		//line grammar.lr:756:4
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
		//line grammar.lr:783:4
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
		//line grammar.lr:815:4
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
	_ReduceAddToProperDefinitionsAction                               = &_Action{_ReduceAction, 0, _ReduceAddToProperDefinitions}
	_ReduceDefinitionToProperDefinitionsAction                        = &_Action{_ReduceAction, 0, _ReduceDefinitionToProperDefinitions}
	_ReduceProperDefinitionsToDefinitionsAction                       = &_Action{_ReduceAction, 0, _ReduceProperDefinitionsToDefinitions}
	_ReduceImproperToDefinitionsAction                                = &_Action{_ReduceAction, 0, _ReduceImproperToDefinitions}
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
	{_State5, _EndMarker}:                                &_Action{_AcceptAction, 0, 0},
	{_State6, _EndMarker}:                                &_Action{_AcceptAction, 0, 0},
	{_State7, _EndMarker}:                                &_Action{_AcceptAction, 0, 0},
	{_State8, _EndMarker}:                                &_Action{_AcceptAction, 0, 0},
	{_State1, CommentGroupsToken}:                        _GotoState9Action,
	{_State1, PackageToken}:                              _GotoState13Action,
	{_State1, TypeToken}:                                 _GotoState14Action,
	{_State1, FuncToken}:                                 _GotoState10Action,
	{_State1, VarToken}:                                  _GotoState15Action,
	{_State1, LetToken}:                                  _GotoState12Action,
	{_State1, LbraceToken}:                               _GotoState11Action,
	{_State1, SourceType}:                                _GotoState5Action,
	{_State1, ProperDefinitionsType}:                     _GotoState20Action,
	{_State1, DefinitionsType}:                           _GotoState17Action,
	{_State1, DefinitionType}:                            _GotoState16Action,
	{_State1, StatementBlockType}:                        _GotoState21Action,
	{_State1, VarDeclPatternType}:                        _GotoState23Action,
	{_State1, VarOrLetType}:                              _GotoState24Action,
	{_State1, TypeDefType}:                               _GotoState22Action,
	{_State1, NamedFuncDefType}:                          _GotoState18Action,
	{_State1, PackageDefType}:                            _GotoState19Action,
	{_State2, IntegerLiteralToken}:                       _GotoState40Action,
	{_State2, FloatLiteralToken}:                         _GotoState35Action,
	{_State2, RuneLiteralToken}:                          _GotoState48Action,
	{_State2, StringLiteralToken}:                        _GotoState49Action,
	{_State2, IdentifierToken}:                           _GotoState38Action,
	{_State2, TrueToken}:                                 _GotoState52Action,
	{_State2, FalseToken}:                                _GotoState34Action,
	{_State2, CaseToken}:                                 _GotoState29Action,
	{_State2, DefaultToken}:                              _GotoState31Action,
	{_State2, ReturnToken}:                               _GotoState47Action,
	{_State2, BreakToken}:                                _GotoState28Action,
	{_State2, ContinueToken}:                             _GotoState30Action,
	{_State2, FallthroughToken}:                          _GotoState33Action,
	{_State2, ImportToken}:                               _GotoState39Action,
	{_State2, UnsafeToken}:                               _GotoState53Action,
	{_State2, StructToken}:                               _GotoState50Action,
	{_State2, FuncToken}:                                 _GotoState36Action,
	{_State2, AsyncToken}:                                _GotoState25Action,
	{_State2, DeferToken}:                                _GotoState32Action,
	{_State2, VarToken}:                                  _GotoState15Action,
	{_State2, LetToken}:                                  _GotoState12Action,
	{_State2, NotToken}:                                  _GotoState45Action,
	{_State2, LabelDeclToken}:                            _GotoState41Action,
	{_State2, LparenToken}:                               _GotoState43Action,
	{_State2, LbracketToken}:                             _GotoState42Action,
	{_State2, SubToken}:                                  _GotoState51Action,
	{_State2, MulToken}:                                  _GotoState44Action,
	{_State2, BitNegToken}:                               _GotoState27Action,
	{_State2, BitAndToken}:                               _GotoState26Action,
	{_State2, GreaterToken}:                              _GotoState37Action,
	{_State2, ParseErrorToken}:                           _GotoState46Action,
	{_State2, SimpleStatementType}:                       _GotoState99Action,
	{_State2, StatementType}:                             _GotoState6Action,
	{_State2, ExpressionOrImproperStructStatementType}:   _GotoState76Action,
	{_State2, ExpressionsType}:                           _GotoState77Action,
	{_State2, CallbackOpType}:                            _GotoState71Action,
	{_State2, CallbackOpStatementType}:                   _GotoState72Action,
	{_State2, UnsafeStatementType}:                       _GotoState102Action,
	{_State2, JumpStatementType}:                         _GotoState85Action,
	{_State2, JumpTypeType}:                              _GotoState86Action,
	{_State2, FallthroughStatementType}:                  _GotoState78Action,
	{_State2, AssignStatementType}:                       _GotoState61Action,
	{_State2, UnaryOpAssignStatementType}:                _GotoState101Action,
	{_State2, BinaryOpAssignStatementType}:               _GotoState67Action,
	{_State2, ImportStatementType}:                       _GotoState81Action,
	{_State2, VarDeclPatternType}:                        _GotoState103Action,
	{_State2, VarOrLetType}:                              _GotoState24Action,
	{_State2, AssignPatternType}:                         _GotoState60Action,
	{_State2, ExpressionType}:                            _GotoState75Action,
	{_State2, OptionalLabelDeclType}:                     _GotoState90Action,
	{_State2, SequenceExprType}:                          _GotoState98Action,
	{_State2, CallExprType}:                              _GotoState70Action,
	{_State2, AtomExprType}:                              _GotoState62Action,
	{_State2, ParseErrorExprType}:                        _GotoState92Action,
	{_State2, LiteralExprType}:                           _GotoState87Action,
	{_State2, IdentifierExprType}:                        _GotoState79Action,
	{_State2, BlockExprType}:                             _GotoState69Action,
	{_State2, InitializeExprType}:                        _GotoState84Action,
	{_State2, ImplicitStructExprType}:                    _GotoState80Action,
	{_State2, AccessibleExprType}:                        _GotoState55Action,
	{_State2, AccessExprType}:                            _GotoState54Action,
	{_State2, IndexExprType}:                             _GotoState82Action,
	{_State2, PostfixableExprType}:                       _GotoState94Action,
	{_State2, PostfixUnaryExprType}:                      _GotoState93Action,
	{_State2, PrefixableExprType}:                        _GotoState97Action,
	{_State2, PrefixUnaryExprType}:                       _GotoState95Action,
	{_State2, PrefixUnaryOpType}:                         _GotoState96Action,
	{_State2, MulExprType}:                               _GotoState89Action,
	{_State2, BinaryMulExprType}:                         _GotoState66Action,
	{_State2, AddExprType}:                               _GotoState56Action,
	{_State2, BinaryAddExprType}:                         _GotoState63Action,
	{_State2, CmpExprType}:                               _GotoState73Action,
	{_State2, BinaryCmpExprType}:                         _GotoState65Action,
	{_State2, AndExprType}:                               _GotoState57Action,
	{_State2, BinaryAndExprType}:                         _GotoState64Action,
	{_State2, OrExprType}:                                _GotoState91Action,
	{_State2, BinaryOrExprType}:                          _GotoState68Action,
	{_State2, InitializableTypeType}:                     _GotoState83Action,
	{_State2, SliceTypeType}:                             _GotoState100Action,
	{_State2, ArrayTypeType}:                             _GotoState59Action,
	{_State2, MapTypeType}:                               _GotoState88Action,
	{_State2, ExplicitStructDefType}:                     _GotoState74Action,
	{_State2, AnonymousFuncExprType}:                     _GotoState58Action,
	{_State3, IntegerLiteralToken}:                       _GotoState40Action,
	{_State3, FloatLiteralToken}:                         _GotoState35Action,
	{_State3, RuneLiteralToken}:                          _GotoState48Action,
	{_State3, StringLiteralToken}:                        _GotoState49Action,
	{_State3, IdentifierToken}:                           _GotoState38Action,
	{_State3, TrueToken}:                                 _GotoState52Action,
	{_State3, FalseToken}:                                _GotoState34Action,
	{_State3, StructToken}:                               _GotoState50Action,
	{_State3, FuncToken}:                                 _GotoState36Action,
	{_State3, VarToken}:                                  _GotoState15Action,
	{_State3, LetToken}:                                  _GotoState12Action,
	{_State3, NotToken}:                                  _GotoState45Action,
	{_State3, LabelDeclToken}:                            _GotoState41Action,
	{_State3, LparenToken}:                               _GotoState43Action,
	{_State3, LbracketToken}:                             _GotoState42Action,
	{_State3, SubToken}:                                  _GotoState51Action,
	{_State3, MulToken}:                                  _GotoState44Action,
	{_State3, BitNegToken}:                               _GotoState27Action,
	{_State3, BitAndToken}:                               _GotoState26Action,
	{_State3, GreaterToken}:                              _GotoState37Action,
	{_State3, ParseErrorToken}:                           _GotoState46Action,
	{_State3, VarDeclPatternType}:                        _GotoState103Action,
	{_State3, VarOrLetType}:                              _GotoState24Action,
	{_State3, ExpressionType}:                            _GotoState7Action,
	{_State3, OptionalLabelDeclType}:                     _GotoState90Action,
	{_State3, SequenceExprType}:                          _GotoState105Action,
	{_State3, CallExprType}:                              _GotoState70Action,
	{_State3, AtomExprType}:                              _GotoState62Action,
	{_State3, ParseErrorExprType}:                        _GotoState92Action,
	{_State3, LiteralExprType}:                           _GotoState87Action,
	{_State3, IdentifierExprType}:                        _GotoState79Action,
	{_State3, BlockExprType}:                             _GotoState69Action,
	{_State3, InitializeExprType}:                        _GotoState84Action,
	{_State3, ImplicitStructExprType}:                    _GotoState80Action,
	{_State3, AccessibleExprType}:                        _GotoState104Action,
	{_State3, AccessExprType}:                            _GotoState54Action,
	{_State3, IndexExprType}:                             _GotoState82Action,
	{_State3, PostfixableExprType}:                       _GotoState94Action,
	{_State3, PostfixUnaryExprType}:                      _GotoState93Action,
	{_State3, PrefixableExprType}:                        _GotoState97Action,
	{_State3, PrefixUnaryExprType}:                       _GotoState95Action,
	{_State3, PrefixUnaryOpType}:                         _GotoState96Action,
	{_State3, MulExprType}:                               _GotoState89Action,
	{_State3, BinaryMulExprType}:                         _GotoState66Action,
	{_State3, AddExprType}:                               _GotoState56Action,
	{_State3, BinaryAddExprType}:                         _GotoState63Action,
	{_State3, CmpExprType}:                               _GotoState73Action,
	{_State3, BinaryCmpExprType}:                         _GotoState65Action,
	{_State3, AndExprType}:                               _GotoState57Action,
	{_State3, BinaryAndExprType}:                         _GotoState64Action,
	{_State3, OrExprType}:                                _GotoState91Action,
	{_State3, BinaryOrExprType}:                          _GotoState68Action,
	{_State3, InitializableTypeType}:                     _GotoState83Action,
	{_State3, SliceTypeType}:                             _GotoState100Action,
	{_State3, ArrayTypeType}:                             _GotoState59Action,
	{_State3, MapTypeType}:                               _GotoState88Action,
	{_State3, ExplicitStructDefType}:                     _GotoState74Action,
	{_State3, AnonymousFuncExprType}:                     _GotoState58Action,
	{_State4, IdentifierToken}:                           _GotoState112Action,
	{_State4, StructToken}:                               _GotoState50Action,
	{_State4, EnumToken}:                                 _GotoState109Action,
	{_State4, TraitToken}:                                _GotoState117Action,
	{_State4, FuncToken}:                                 _GotoState111Action,
	{_State4, LparenToken}:                               _GotoState113Action,
	{_State4, LbracketToken}:                             _GotoState42Action,
	{_State4, DotToken}:                                  _GotoState108Action,
	{_State4, QuestionToken}:                             _GotoState115Action,
	{_State4, ExclaimToken}:                              _GotoState110Action,
	{_State4, TildeTildeToken}:                           _GotoState116Action,
	{_State4, BitNegToken}:                               _GotoState107Action,
	{_State4, BitAndToken}:                               _GotoState106Action,
	{_State4, ParseErrorToken}:                           _GotoState114Action,
	{_State4, InitializableTypeType}:                     _GotoState123Action,
	{_State4, SliceTypeType}:                             _GotoState100Action,
	{_State4, ArrayTypeType}:                             _GotoState59Action,
	{_State4, MapTypeType}:                               _GotoState88Action,
	{_State4, AtomTypeType}:                              _GotoState118Action,
	{_State4, ParseErrorTypeType}:                        _GotoState124Action,
	{_State4, ReturnableTypeType}:                        _GotoState127Action,
	{_State4, PrefixedTypeType}:                          _GotoState126Action,
	{_State4, PrefixTypeOpType}:                          _GotoState125Action,
	{_State4, ValueTypeType}:                             _GotoState8Action,
	{_State4, TraitOpTypeType}:                           _GotoState129Action,
	{_State4, ImplicitStructDefType}:                     _GotoState122Action,
	{_State4, ExplicitStructDefType}:                     _GotoState74Action,
	{_State4, ImplicitEnumDefType}:                       _GotoState121Action,
	{_State4, ExplicitEnumDefType}:                       _GotoState119Action,
	{_State4, TraitDefType}:                              _GotoState128Action,
	{_State4, FuncTypeType}:                              _GotoState120Action,
	{_State8, AddToken}:                                  _GotoState239Action,
	{_State8, SubToken}:                                  _GotoState241Action,
	{_State8, MulToken}:                                  _GotoState240Action,
	{_State8, TraitOpType}:                               _GotoState242Action,
	{_State10, IdentifierToken}:                          _GotoState130Action,
	{_State10, LparenToken}:                              _GotoState131Action,
	{_State11, IntegerLiteralToken}:                      _GotoState40Action,
	{_State11, FloatLiteralToken}:                        _GotoState35Action,
	{_State11, RuneLiteralToken}:                         _GotoState48Action,
	{_State11, StringLiteralToken}:                       _GotoState49Action,
	{_State11, IdentifierToken}:                          _GotoState38Action,
	{_State11, TrueToken}:                                _GotoState52Action,
	{_State11, FalseToken}:                               _GotoState34Action,
	{_State11, CaseToken}:                                _GotoState29Action,
	{_State11, DefaultToken}:                             _GotoState31Action,
	{_State11, ReturnToken}:                              _GotoState47Action,
	{_State11, BreakToken}:                               _GotoState28Action,
	{_State11, ContinueToken}:                            _GotoState30Action,
	{_State11, FallthroughToken}:                         _GotoState33Action,
	{_State11, ImportToken}:                              _GotoState39Action,
	{_State11, UnsafeToken}:                              _GotoState53Action,
	{_State11, StructToken}:                              _GotoState50Action,
	{_State11, FuncToken}:                                _GotoState36Action,
	{_State11, AsyncToken}:                               _GotoState25Action,
	{_State11, DeferToken}:                               _GotoState32Action,
	{_State11, VarToken}:                                 _GotoState15Action,
	{_State11, LetToken}:                                 _GotoState12Action,
	{_State11, NotToken}:                                 _GotoState45Action,
	{_State11, LabelDeclToken}:                           _GotoState41Action,
	{_State11, LparenToken}:                              _GotoState43Action,
	{_State11, LbracketToken}:                            _GotoState42Action,
	{_State11, SubToken}:                                 _GotoState51Action,
	{_State11, MulToken}:                                 _GotoState44Action,
	{_State11, BitNegToken}:                              _GotoState27Action,
	{_State11, BitAndToken}:                              _GotoState26Action,
	{_State11, GreaterToken}:                             _GotoState37Action,
	{_State11, ParseErrorToken}:                          _GotoState46Action,
	{_State11, ProperStatementsType}:                     _GotoState132Action,
	{_State11, StatementsType}:                           _GotoState134Action,
	{_State11, SimpleStatementType}:                      _GotoState99Action,
	{_State11, StatementType}:                            _GotoState133Action,
	{_State11, ExpressionOrImproperStructStatementType}:  _GotoState76Action,
	{_State11, ExpressionsType}:                          _GotoState77Action,
	{_State11, CallbackOpType}:                           _GotoState71Action,
	{_State11, CallbackOpStatementType}:                  _GotoState72Action,
	{_State11, UnsafeStatementType}:                      _GotoState102Action,
	{_State11, JumpStatementType}:                        _GotoState85Action,
	{_State11, JumpTypeType}:                             _GotoState86Action,
	{_State11, FallthroughStatementType}:                 _GotoState78Action,
	{_State11, AssignStatementType}:                      _GotoState61Action,
	{_State11, UnaryOpAssignStatementType}:               _GotoState101Action,
	{_State11, BinaryOpAssignStatementType}:              _GotoState67Action,
	{_State11, ImportStatementType}:                      _GotoState81Action,
	{_State11, VarDeclPatternType}:                       _GotoState103Action,
	{_State11, VarOrLetType}:                             _GotoState24Action,
	{_State11, AssignPatternType}:                        _GotoState60Action,
	{_State11, ExpressionType}:                           _GotoState75Action,
	{_State11, OptionalLabelDeclType}:                    _GotoState90Action,
	{_State11, SequenceExprType}:                         _GotoState98Action,
	{_State11, CallExprType}:                             _GotoState70Action,
	{_State11, AtomExprType}:                             _GotoState62Action,
	{_State11, ParseErrorExprType}:                       _GotoState92Action,
	{_State11, LiteralExprType}:                          _GotoState87Action,
	{_State11, IdentifierExprType}:                       _GotoState79Action,
	{_State11, BlockExprType}:                            _GotoState69Action,
	{_State11, InitializeExprType}:                       _GotoState84Action,
	{_State11, ImplicitStructExprType}:                   _GotoState80Action,
	{_State11, AccessibleExprType}:                       _GotoState55Action,
	{_State11, AccessExprType}:                           _GotoState54Action,
	{_State11, IndexExprType}:                            _GotoState82Action,
	{_State11, PostfixableExprType}:                      _GotoState94Action,
	{_State11, PostfixUnaryExprType}:                     _GotoState93Action,
	{_State11, PrefixableExprType}:                       _GotoState97Action,
	{_State11, PrefixUnaryExprType}:                      _GotoState95Action,
	{_State11, PrefixUnaryOpType}:                        _GotoState96Action,
	{_State11, MulExprType}:                              _GotoState89Action,
	{_State11, BinaryMulExprType}:                        _GotoState66Action,
	{_State11, AddExprType}:                              _GotoState56Action,
	{_State11, BinaryAddExprType}:                        _GotoState63Action,
	{_State11, CmpExprType}:                              _GotoState73Action,
	{_State11, BinaryCmpExprType}:                        _GotoState65Action,
	{_State11, AndExprType}:                              _GotoState57Action,
	{_State11, BinaryAndExprType}:                        _GotoState64Action,
	{_State11, OrExprType}:                               _GotoState91Action,
	{_State11, BinaryOrExprType}:                         _GotoState68Action,
	{_State11, InitializableTypeType}:                    _GotoState83Action,
	{_State11, SliceTypeType}:                            _GotoState100Action,
	{_State11, ArrayTypeType}:                            _GotoState59Action,
	{_State11, MapTypeType}:                              _GotoState88Action,
	{_State11, ExplicitStructDefType}:                    _GotoState74Action,
	{_State11, AnonymousFuncExprType}:                    _GotoState58Action,
	{_State13, LbraceToken}:                              _GotoState11Action,
	{_State13, StatementBlockType}:                       _GotoState135Action,
	{_State14, IdentifierToken}:                          _GotoState136Action,
	{_State20, NewlinesToken}:                            _GotoState137Action,
	{_State23, AssignToken}:                              _GotoState138Action,
	{_State24, IdentifierToken}:                          _GotoState139Action,
	{_State24, LparenToken}:                              _GotoState140Action,
	{_State24, VarPatternType}:                           _GotoState142Action,
	{_State24, TuplePatternType}:                         _GotoState141Action,
	{_State29, IntegerLiteralToken}:                      _GotoState40Action,
	{_State29, FloatLiteralToken}:                        _GotoState35Action,
	{_State29, RuneLiteralToken}:                         _GotoState48Action,
	{_State29, StringLiteralToken}:                       _GotoState49Action,
	{_State29, IdentifierToken}:                          _GotoState38Action,
	{_State29, TrueToken}:                                _GotoState52Action,
	{_State29, FalseToken}:                               _GotoState34Action,
	{_State29, StructToken}:                              _GotoState50Action,
	{_State29, FuncToken}:                                _GotoState36Action,
	{_State29, VarToken}:                                 _GotoState144Action,
	{_State29, LetToken}:                                 _GotoState12Action,
	{_State29, NotToken}:                                 _GotoState45Action,
	{_State29, LabelDeclToken}:                           _GotoState41Action,
	{_State29, LparenToken}:                              _GotoState43Action,
	{_State29, LbracketToken}:                            _GotoState42Action,
	{_State29, DotToken}:                                 _GotoState143Action,
	{_State29, SubToken}:                                 _GotoState51Action,
	{_State29, MulToken}:                                 _GotoState44Action,
	{_State29, BitNegToken}:                              _GotoState27Action,
	{_State29, BitAndToken}:                              _GotoState26Action,
	{_State29, GreaterToken}:                             _GotoState37Action,
	{_State29, ParseErrorToken}:                          _GotoState46Action,
	{_State29, CasePatternsType}:                         _GotoState147Action,
	{_State29, VarDeclPatternType}:                       _GotoState103Action,
	{_State29, VarOrLetType}:                             _GotoState24Action,
	{_State29, AssignPatternType}:                        _GotoState145Action,
	{_State29, CasePatternType}:                          _GotoState146Action,
	{_State29, OptionalLabelDeclType}:                    _GotoState148Action,
	{_State29, SequenceExprType}:                         _GotoState149Action,
	{_State29, CallExprType}:                             _GotoState70Action,
	{_State29, AtomExprType}:                             _GotoState62Action,
	{_State29, ParseErrorExprType}:                       _GotoState92Action,
	{_State29, LiteralExprType}:                          _GotoState87Action,
	{_State29, IdentifierExprType}:                       _GotoState79Action,
	{_State29, BlockExprType}:                            _GotoState69Action,
	{_State29, InitializeExprType}:                       _GotoState84Action,
	{_State29, ImplicitStructExprType}:                   _GotoState80Action,
	{_State29, AccessibleExprType}:                       _GotoState104Action,
	{_State29, AccessExprType}:                           _GotoState54Action,
	{_State29, IndexExprType}:                            _GotoState82Action,
	{_State29, PostfixableExprType}:                      _GotoState94Action,
	{_State29, PostfixUnaryExprType}:                     _GotoState93Action,
	{_State29, PrefixableExprType}:                       _GotoState97Action,
	{_State29, PrefixUnaryExprType}:                      _GotoState95Action,
	{_State29, PrefixUnaryOpType}:                        _GotoState96Action,
	{_State29, MulExprType}:                              _GotoState89Action,
	{_State29, BinaryMulExprType}:                        _GotoState66Action,
	{_State29, AddExprType}:                              _GotoState56Action,
	{_State29, BinaryAddExprType}:                        _GotoState63Action,
	{_State29, CmpExprType}:                              _GotoState73Action,
	{_State29, BinaryCmpExprType}:                        _GotoState65Action,
	{_State29, AndExprType}:                              _GotoState57Action,
	{_State29, BinaryAndExprType}:                        _GotoState64Action,
	{_State29, OrExprType}:                               _GotoState91Action,
	{_State29, BinaryOrExprType}:                         _GotoState68Action,
	{_State29, InitializableTypeType}:                    _GotoState83Action,
	{_State29, SliceTypeType}:                            _GotoState100Action,
	{_State29, ArrayTypeType}:                            _GotoState59Action,
	{_State29, MapTypeType}:                              _GotoState88Action,
	{_State29, ExplicitStructDefType}:                    _GotoState74Action,
	{_State29, AnonymousFuncExprType}:                    _GotoState58Action,
	{_State31, ColonToken}:                               _GotoState150Action,
	{_State36, LparenToken}:                              _GotoState151Action,
	{_State37, IntegerLiteralToken}:                      _GotoState40Action,
	{_State37, FloatLiteralToken}:                        _GotoState35Action,
	{_State37, RuneLiteralToken}:                         _GotoState48Action,
	{_State37, StringLiteralToken}:                       _GotoState49Action,
	{_State37, IdentifierToken}:                          _GotoState38Action,
	{_State37, TrueToken}:                                _GotoState52Action,
	{_State37, FalseToken}:                               _GotoState34Action,
	{_State37, StructToken}:                              _GotoState50Action,
	{_State37, FuncToken}:                                _GotoState36Action,
	{_State37, VarToken}:                                 _GotoState15Action,
	{_State37, LetToken}:                                 _GotoState12Action,
	{_State37, NotToken}:                                 _GotoState45Action,
	{_State37, LabelDeclToken}:                           _GotoState41Action,
	{_State37, LparenToken}:                              _GotoState43Action,
	{_State37, LbracketToken}:                            _GotoState42Action,
	{_State37, SubToken}:                                 _GotoState51Action,
	{_State37, MulToken}:                                 _GotoState44Action,
	{_State37, BitNegToken}:                              _GotoState27Action,
	{_State37, BitAndToken}:                              _GotoState26Action,
	{_State37, GreaterToken}:                             _GotoState37Action,
	{_State37, ParseErrorToken}:                          _GotoState46Action,
	{_State37, VarDeclPatternType}:                       _GotoState103Action,
	{_State37, VarOrLetType}:                             _GotoState24Action,
	{_State37, OptionalLabelDeclType}:                    _GotoState148Action,
	{_State37, SequenceExprType}:                         _GotoState152Action,
	{_State37, CallExprType}:                             _GotoState70Action,
	{_State37, AtomExprType}:                             _GotoState62Action,
	{_State37, ParseErrorExprType}:                       _GotoState92Action,
	{_State37, LiteralExprType}:                          _GotoState87Action,
	{_State37, IdentifierExprType}:                       _GotoState79Action,
	{_State37, BlockExprType}:                            _GotoState69Action,
	{_State37, InitializeExprType}:                       _GotoState84Action,
	{_State37, ImplicitStructExprType}:                   _GotoState80Action,
	{_State37, AccessibleExprType}:                       _GotoState104Action,
	{_State37, AccessExprType}:                           _GotoState54Action,
	{_State37, IndexExprType}:                            _GotoState82Action,
	{_State37, PostfixableExprType}:                      _GotoState94Action,
	{_State37, PostfixUnaryExprType}:                     _GotoState93Action,
	{_State37, PrefixableExprType}:                       _GotoState97Action,
	{_State37, PrefixUnaryExprType}:                      _GotoState95Action,
	{_State37, PrefixUnaryOpType}:                        _GotoState96Action,
	{_State37, MulExprType}:                              _GotoState89Action,
	{_State37, BinaryMulExprType}:                        _GotoState66Action,
	{_State37, AddExprType}:                              _GotoState56Action,
	{_State37, BinaryAddExprType}:                        _GotoState63Action,
	{_State37, CmpExprType}:                              _GotoState73Action,
	{_State37, BinaryCmpExprType}:                        _GotoState65Action,
	{_State37, AndExprType}:                              _GotoState57Action,
	{_State37, BinaryAndExprType}:                        _GotoState64Action,
	{_State37, OrExprType}:                               _GotoState91Action,
	{_State37, BinaryOrExprType}:                         _GotoState68Action,
	{_State37, InitializableTypeType}:                    _GotoState83Action,
	{_State37, SliceTypeType}:                            _GotoState100Action,
	{_State37, ArrayTypeType}:                            _GotoState59Action,
	{_State37, MapTypeType}:                              _GotoState88Action,
	{_State37, ExplicitStructDefType}:                    _GotoState74Action,
	{_State37, AnonymousFuncExprType}:                    _GotoState58Action,
	{_State39, StringLiteralToken}:                       _GotoState154Action,
	{_State39, LparenToken}:                              _GotoState153Action,
	{_State39, ImportClauseType}:                         _GotoState155Action,
	{_State42, IdentifierToken}:                          _GotoState112Action,
	{_State42, StructToken}:                              _GotoState50Action,
	{_State42, EnumToken}:                                _GotoState109Action,
	{_State42, TraitToken}:                               _GotoState117Action,
	{_State42, FuncToken}:                                _GotoState111Action,
	{_State42, LparenToken}:                              _GotoState113Action,
	{_State42, LbracketToken}:                            _GotoState42Action,
	{_State42, DotToken}:                                 _GotoState108Action,
	{_State42, QuestionToken}:                            _GotoState115Action,
	{_State42, ExclaimToken}:                             _GotoState110Action,
	{_State42, TildeTildeToken}:                          _GotoState116Action,
	{_State42, BitNegToken}:                              _GotoState107Action,
	{_State42, BitAndToken}:                              _GotoState106Action,
	{_State42, ParseErrorToken}:                          _GotoState114Action,
	{_State42, InitializableTypeType}:                    _GotoState123Action,
	{_State42, SliceTypeType}:                            _GotoState100Action,
	{_State42, ArrayTypeType}:                            _GotoState59Action,
	{_State42, MapTypeType}:                              _GotoState88Action,
	{_State42, AtomTypeType}:                             _GotoState118Action,
	{_State42, ParseErrorTypeType}:                       _GotoState124Action,
	{_State42, ReturnableTypeType}:                       _GotoState127Action,
	{_State42, PrefixedTypeType}:                         _GotoState126Action,
	{_State42, PrefixTypeOpType}:                         _GotoState125Action,
	{_State42, ValueTypeType}:                            _GotoState156Action,
	{_State42, TraitOpTypeType}:                          _GotoState129Action,
	{_State42, ImplicitStructDefType}:                    _GotoState122Action,
	{_State42, ExplicitStructDefType}:                    _GotoState74Action,
	{_State42, ImplicitEnumDefType}:                      _GotoState121Action,
	{_State42, ExplicitEnumDefType}:                      _GotoState119Action,
	{_State42, TraitDefType}:                             _GotoState128Action,
	{_State42, FuncTypeType}:                             _GotoState120Action,
	{_State43, IntegerLiteralToken}:                      _GotoState40Action,
	{_State43, FloatLiteralToken}:                        _GotoState35Action,
	{_State43, RuneLiteralToken}:                         _GotoState48Action,
	{_State43, StringLiteralToken}:                       _GotoState49Action,
	{_State43, IdentifierToken}:                          _GotoState159Action,
	{_State43, TrueToken}:                                _GotoState52Action,
	{_State43, FalseToken}:                               _GotoState34Action,
	{_State43, StructToken}:                              _GotoState50Action,
	{_State43, FuncToken}:                                _GotoState36Action,
	{_State43, VarToken}:                                 _GotoState15Action,
	{_State43, LetToken}:                                 _GotoState12Action,
	{_State43, NotToken}:                                 _GotoState45Action,
	{_State43, LabelDeclToken}:                           _GotoState41Action,
	{_State43, LparenToken}:                              _GotoState43Action,
	{_State43, LbracketToken}:                            _GotoState42Action,
	{_State43, ColonToken}:                               _GotoState157Action,
	{_State43, EllipsisToken}:                            _GotoState158Action,
	{_State43, SubToken}:                                 _GotoState51Action,
	{_State43, MulToken}:                                 _GotoState44Action,
	{_State43, BitNegToken}:                              _GotoState27Action,
	{_State43, BitAndToken}:                              _GotoState26Action,
	{_State43, GreaterToken}:                             _GotoState37Action,
	{_State43, ParseErrorToken}:                          _GotoState46Action,
	{_State43, VarDeclPatternType}:                       _GotoState103Action,
	{_State43, VarOrLetType}:                             _GotoState24Action,
	{_State43, ExpressionType}:                           _GotoState163Action,
	{_State43, OptionalLabelDeclType}:                    _GotoState90Action,
	{_State43, SequenceExprType}:                         _GotoState105Action,
	{_State43, CallExprType}:                             _GotoState70Action,
	{_State43, ProperArgumentsType}:                      _GotoState164Action,
	{_State43, ArgumentsType}:                            _GotoState161Action,
	{_State43, ArgumentType}:                             _GotoState160Action,
	{_State43, ColonExprType}:                            _GotoState162Action,
	{_State43, AtomExprType}:                             _GotoState62Action,
	{_State43, ParseErrorExprType}:                       _GotoState92Action,
	{_State43, LiteralExprType}:                          _GotoState87Action,
	{_State43, IdentifierExprType}:                       _GotoState79Action,
	{_State43, BlockExprType}:                            _GotoState69Action,
	{_State43, InitializeExprType}:                       _GotoState84Action,
	{_State43, ImplicitStructExprType}:                   _GotoState80Action,
	{_State43, AccessibleExprType}:                       _GotoState104Action,
	{_State43, AccessExprType}:                           _GotoState54Action,
	{_State43, IndexExprType}:                            _GotoState82Action,
	{_State43, PostfixableExprType}:                      _GotoState94Action,
	{_State43, PostfixUnaryExprType}:                     _GotoState93Action,
	{_State43, PrefixableExprType}:                       _GotoState97Action,
	{_State43, PrefixUnaryExprType}:                      _GotoState95Action,
	{_State43, PrefixUnaryOpType}:                        _GotoState96Action,
	{_State43, MulExprType}:                              _GotoState89Action,
	{_State43, BinaryMulExprType}:                        _GotoState66Action,
	{_State43, AddExprType}:                              _GotoState56Action,
	{_State43, BinaryAddExprType}:                        _GotoState63Action,
	{_State43, CmpExprType}:                              _GotoState73Action,
	{_State43, BinaryCmpExprType}:                        _GotoState65Action,
	{_State43, AndExprType}:                              _GotoState57Action,
	{_State43, BinaryAndExprType}:                        _GotoState64Action,
	{_State43, OrExprType}:                               _GotoState91Action,
	{_State43, BinaryOrExprType}:                         _GotoState68Action,
	{_State43, InitializableTypeType}:                    _GotoState83Action,
	{_State43, SliceTypeType}:                            _GotoState100Action,
	{_State43, ArrayTypeType}:                            _GotoState59Action,
	{_State43, MapTypeType}:                              _GotoState88Action,
	{_State43, ExplicitStructDefType}:                    _GotoState74Action,
	{_State43, AnonymousFuncExprType}:                    _GotoState58Action,
	{_State50, LparenToken}:                              _GotoState165Action,
	{_State53, LessToken}:                                _GotoState166Action,
	{_State55, LbracketToken}:                            _GotoState178Action,
	{_State55, DotToken}:                                 _GotoState177Action,
	{_State55, QuestionToken}:                            _GotoState181Action,
	{_State55, DollarLbracketToken}:                      _GotoState176Action,
	{_State55, AddAssignToken}:                           _GotoState167Action,
	{_State55, SubAssignToken}:                           _GotoState182Action,
	{_State55, MulAssignToken}:                           _GotoState180Action,
	{_State55, DivAssignToken}:                           _GotoState175Action,
	{_State55, ModAssignToken}:                           _GotoState179Action,
	{_State55, AddOneAssignToken}:                        _GotoState168Action,
	{_State55, SubOneAssignToken}:                        _GotoState183Action,
	{_State55, BitNegAssignToken}:                        _GotoState171Action,
	{_State55, BitAndAssignToken}:                        _GotoState169Action,
	{_State55, BitOrAssignToken}:                         _GotoState172Action,
	{_State55, BitXorAssignToken}:                        _GotoState174Action,
	{_State55, BitLshiftAssignToken}:                     _GotoState170Action,
	{_State55, BitRshiftAssignToken}:                     _GotoState173Action,
	{_State55, UnaryOpAssignType}:                        _GotoState186Action,
	{_State55, BinaryOpAssignType}:                       _GotoState184Action,
	{_State55, OptionalGenericBindingType}:               _GotoState185Action,
	{_State56, AddToken}:                                 _GotoState187Action,
	{_State56, SubToken}:                                 _GotoState190Action,
	{_State56, BitXorToken}:                              _GotoState189Action,
	{_State56, BitOrToken}:                               _GotoState188Action,
	{_State56, AddOpType}:                                _GotoState191Action,
	{_State57, AndToken}:                                 _GotoState192Action,
	{_State60, AssignToken}:                              _GotoState193Action,
	{_State71, IntegerLiteralToken}:                      _GotoState40Action,
	{_State71, FloatLiteralToken}:                        _GotoState35Action,
	{_State71, RuneLiteralToken}:                         _GotoState48Action,
	{_State71, StringLiteralToken}:                       _GotoState49Action,
	{_State71, IdentifierToken}:                          _GotoState38Action,
	{_State71, TrueToken}:                                _GotoState52Action,
	{_State71, FalseToken}:                               _GotoState34Action,
	{_State71, StructToken}:                              _GotoState50Action,
	{_State71, FuncToken}:                                _GotoState36Action,
	{_State71, LabelDeclToken}:                           _GotoState41Action,
	{_State71, LparenToken}:                              _GotoState43Action,
	{_State71, LbracketToken}:                            _GotoState42Action,
	{_State71, ParseErrorToken}:                          _GotoState46Action,
	{_State71, OptionalLabelDeclType}:                    _GotoState148Action,
	{_State71, CallExprType}:                             _GotoState195Action,
	{_State71, AtomExprType}:                             _GotoState62Action,
	{_State71, ParseErrorExprType}:                       _GotoState92Action,
	{_State71, LiteralExprType}:                          _GotoState87Action,
	{_State71, IdentifierExprType}:                       _GotoState79Action,
	{_State71, BlockExprType}:                            _GotoState69Action,
	{_State71, InitializeExprType}:                       _GotoState84Action,
	{_State71, ImplicitStructExprType}:                   _GotoState80Action,
	{_State71, AccessibleExprType}:                       _GotoState194Action,
	{_State71, AccessExprType}:                           _GotoState54Action,
	{_State71, IndexExprType}:                            _GotoState82Action,
	{_State71, InitializableTypeType}:                    _GotoState83Action,
	{_State71, SliceTypeType}:                            _GotoState100Action,
	{_State71, ArrayTypeType}:                            _GotoState59Action,
	{_State71, MapTypeType}:                              _GotoState88Action,
	{_State71, ExplicitStructDefType}:                    _GotoState74Action,
	{_State71, AnonymousFuncExprType}:                    _GotoState58Action,
	{_State73, EqualToken}:                               _GotoState196Action,
	{_State73, NotEqualToken}:                            _GotoState201Action,
	{_State73, LessToken}:                                _GotoState199Action,
	{_State73, LessOrEqualToken}:                         _GotoState200Action,
	{_State73, GreaterToken}:                             _GotoState197Action,
	{_State73, GreaterOrEqualToken}:                      _GotoState198Action,
	{_State73, CmpOpType}:                                _GotoState202Action,
	{_State77, CommaToken}:                               _GotoState203Action,
	{_State83, LparenToken}:                              _GotoState204Action,
	{_State86, IntegerLiteralToken}:                      _GotoState40Action,
	{_State86, FloatLiteralToken}:                        _GotoState35Action,
	{_State86, RuneLiteralToken}:                         _GotoState48Action,
	{_State86, StringLiteralToken}:                       _GotoState49Action,
	{_State86, IdentifierToken}:                          _GotoState38Action,
	{_State86, TrueToken}:                                _GotoState52Action,
	{_State86, FalseToken}:                               _GotoState34Action,
	{_State86, StructToken}:                              _GotoState50Action,
	{_State86, FuncToken}:                                _GotoState36Action,
	{_State86, VarToken}:                                 _GotoState15Action,
	{_State86, LetToken}:                                 _GotoState12Action,
	{_State86, NotToken}:                                 _GotoState45Action,
	{_State86, LabelDeclToken}:                           _GotoState41Action,
	{_State86, JumpLabelToken}:                           _GotoState205Action,
	{_State86, LparenToken}:                              _GotoState43Action,
	{_State86, LbracketToken}:                            _GotoState42Action,
	{_State86, SubToken}:                                 _GotoState51Action,
	{_State86, MulToken}:                                 _GotoState44Action,
	{_State86, BitNegToken}:                              _GotoState27Action,
	{_State86, BitAndToken}:                              _GotoState26Action,
	{_State86, GreaterToken}:                             _GotoState37Action,
	{_State86, ParseErrorToken}:                          _GotoState46Action,
	{_State86, ExpressionsType}:                          _GotoState206Action,
	{_State86, VarDeclPatternType}:                       _GotoState103Action,
	{_State86, VarOrLetType}:                             _GotoState24Action,
	{_State86, ExpressionType}:                           _GotoState75Action,
	{_State86, OptionalLabelDeclType}:                    _GotoState90Action,
	{_State86, SequenceExprType}:                         _GotoState105Action,
	{_State86, CallExprType}:                             _GotoState70Action,
	{_State86, AtomExprType}:                             _GotoState62Action,
	{_State86, ParseErrorExprType}:                       _GotoState92Action,
	{_State86, LiteralExprType}:                          _GotoState87Action,
	{_State86, IdentifierExprType}:                       _GotoState79Action,
	{_State86, BlockExprType}:                            _GotoState69Action,
	{_State86, InitializeExprType}:                       _GotoState84Action,
	{_State86, ImplicitStructExprType}:                   _GotoState80Action,
	{_State86, AccessibleExprType}:                       _GotoState104Action,
	{_State86, AccessExprType}:                           _GotoState54Action,
	{_State86, IndexExprType}:                            _GotoState82Action,
	{_State86, PostfixableExprType}:                      _GotoState94Action,
	{_State86, PostfixUnaryExprType}:                     _GotoState93Action,
	{_State86, PrefixableExprType}:                       _GotoState97Action,
	{_State86, PrefixUnaryExprType}:                      _GotoState95Action,
	{_State86, PrefixUnaryOpType}:                        _GotoState96Action,
	{_State86, MulExprType}:                              _GotoState89Action,
	{_State86, BinaryMulExprType}:                        _GotoState66Action,
	{_State86, AddExprType}:                              _GotoState56Action,
	{_State86, BinaryAddExprType}:                        _GotoState63Action,
	{_State86, CmpExprType}:                              _GotoState73Action,
	{_State86, BinaryCmpExprType}:                        _GotoState65Action,
	{_State86, AndExprType}:                              _GotoState57Action,
	{_State86, BinaryAndExprType}:                        _GotoState64Action,
	{_State86, OrExprType}:                               _GotoState91Action,
	{_State86, BinaryOrExprType}:                         _GotoState68Action,
	{_State86, InitializableTypeType}:                    _GotoState83Action,
	{_State86, SliceTypeType}:                            _GotoState100Action,
	{_State86, ArrayTypeType}:                            _GotoState59Action,
	{_State86, MapTypeType}:                              _GotoState88Action,
	{_State86, ExplicitStructDefType}:                    _GotoState74Action,
	{_State86, AnonymousFuncExprType}:                    _GotoState58Action,
	{_State89, MulToken}:                                 _GotoState212Action,
	{_State89, DivToken}:                                 _GotoState210Action,
	{_State89, ModToken}:                                 _GotoState211Action,
	{_State89, BitAndToken}:                              _GotoState207Action,
	{_State89, BitLshiftToken}:                           _GotoState208Action,
	{_State89, BitRshiftToken}:                           _GotoState209Action,
	{_State89, MulOpType}:                                _GotoState213Action,
	{_State90, IfToken}:                                  _GotoState216Action,
	{_State90, SwitchToken}:                              _GotoState217Action,
	{_State90, ForToken}:                                 _GotoState215Action,
	{_State90, DoToken}:                                  _GotoState214Action,
	{_State90, LbraceToken}:                              _GotoState11Action,
	{_State90, StatementBlockType}:                       _GotoState220Action,
	{_State90, IfExprType}:                               _GotoState218Action,
	{_State90, SwitchExprType}:                           _GotoState221Action,
	{_State90, LoopExprType}:                             _GotoState219Action,
	{_State91, OrToken}:                                  _GotoState222Action,
	{_State96, IntegerLiteralToken}:                      _GotoState40Action,
	{_State96, FloatLiteralToken}:                        _GotoState35Action,
	{_State96, RuneLiteralToken}:                         _GotoState48Action,
	{_State96, StringLiteralToken}:                       _GotoState49Action,
	{_State96, IdentifierToken}:                          _GotoState38Action,
	{_State96, TrueToken}:                                _GotoState52Action,
	{_State96, FalseToken}:                               _GotoState34Action,
	{_State96, StructToken}:                              _GotoState50Action,
	{_State96, FuncToken}:                                _GotoState36Action,
	{_State96, NotToken}:                                 _GotoState45Action,
	{_State96, LabelDeclToken}:                           _GotoState41Action,
	{_State96, LparenToken}:                              _GotoState43Action,
	{_State96, LbracketToken}:                            _GotoState42Action,
	{_State96, SubToken}:                                 _GotoState51Action,
	{_State96, MulToken}:                                 _GotoState44Action,
	{_State96, BitNegToken}:                              _GotoState27Action,
	{_State96, BitAndToken}:                              _GotoState26Action,
	{_State96, ParseErrorToken}:                          _GotoState46Action,
	{_State96, OptionalLabelDeclType}:                    _GotoState148Action,
	{_State96, CallExprType}:                             _GotoState70Action,
	{_State96, AtomExprType}:                             _GotoState62Action,
	{_State96, ParseErrorExprType}:                       _GotoState92Action,
	{_State96, LiteralExprType}:                          _GotoState87Action,
	{_State96, IdentifierExprType}:                       _GotoState79Action,
	{_State96, BlockExprType}:                            _GotoState69Action,
	{_State96, InitializeExprType}:                       _GotoState84Action,
	{_State96, ImplicitStructExprType}:                   _GotoState80Action,
	{_State96, AccessibleExprType}:                       _GotoState104Action,
	{_State96, AccessExprType}:                           _GotoState54Action,
	{_State96, IndexExprType}:                            _GotoState82Action,
	{_State96, PostfixableExprType}:                      _GotoState94Action,
	{_State96, PostfixUnaryExprType}:                     _GotoState93Action,
	{_State96, PrefixableExprType}:                       _GotoState223Action,
	{_State96, PrefixUnaryExprType}:                      _GotoState95Action,
	{_State96, PrefixUnaryOpType}:                        _GotoState96Action,
	{_State96, InitializableTypeType}:                    _GotoState83Action,
	{_State96, SliceTypeType}:                            _GotoState100Action,
	{_State96, ArrayTypeType}:                            _GotoState59Action,
	{_State96, MapTypeType}:                              _GotoState88Action,
	{_State96, ExplicitStructDefType}:                    _GotoState74Action,
	{_State96, AnonymousFuncExprType}:                    _GotoState58Action,
	{_State104, LbracketToken}:                           _GotoState178Action,
	{_State104, DotToken}:                                _GotoState177Action,
	{_State104, QuestionToken}:                           _GotoState181Action,
	{_State104, DollarLbracketToken}:                     _GotoState176Action,
	{_State104, OptionalGenericBindingType}:              _GotoState185Action,
	{_State108, DollarLbracketToken}:                     _GotoState176Action,
	{_State108, OptionalGenericBindingType}:              _GotoState224Action,
	{_State109, LparenToken}:                             _GotoState225Action,
	{_State111, LparenToken}:                             _GotoState226Action,
	{_State112, DotToken}:                                _GotoState227Action,
	{_State112, DollarLbracketToken}:                     _GotoState176Action,
	{_State112, OptionalGenericBindingType}:              _GotoState228Action,
	{_State113, IdentifierToken}:                         _GotoState229Action,
	{_State113, UnsafeToken}:                             _GotoState53Action,
	{_State113, StructToken}:                             _GotoState50Action,
	{_State113, EnumToken}:                               _GotoState109Action,
	{_State113, TraitToken}:                              _GotoState117Action,
	{_State113, FuncToken}:                               _GotoState111Action,
	{_State113, LparenToken}:                             _GotoState113Action,
	{_State113, LbracketToken}:                           _GotoState42Action,
	{_State113, DotToken}:                                _GotoState108Action,
	{_State113, QuestionToken}:                           _GotoState115Action,
	{_State113, ExclaimToken}:                            _GotoState110Action,
	{_State113, TildeTildeToken}:                         _GotoState116Action,
	{_State113, BitNegToken}:                             _GotoState107Action,
	{_State113, BitAndToken}:                             _GotoState106Action,
	{_State113, ParseErrorToken}:                         _GotoState114Action,
	{_State113, UnsafeStatementType}:                     _GotoState235Action,
	{_State113, InitializableTypeType}:                   _GotoState123Action,
	{_State113, SliceTypeType}:                           _GotoState100Action,
	{_State113, ArrayTypeType}:                           _GotoState59Action,
	{_State113, MapTypeType}:                             _GotoState88Action,
	{_State113, AtomTypeType}:                            _GotoState118Action,
	{_State113, ParseErrorTypeType}:                      _GotoState124Action,
	{_State113, ReturnableTypeType}:                      _GotoState127Action,
	{_State113, PrefixedTypeType}:                        _GotoState126Action,
	{_State113, PrefixTypeOpType}:                        _GotoState125Action,
	{_State113, ValueTypeType}:                           _GotoState236Action,
	{_State113, TraitOpTypeType}:                         _GotoState129Action,
	{_State113, FieldDefType}:                            _GotoState231Action,
	{_State113, ProperImplicitFieldDefsType}:             _GotoState234Action,
	{_State113, ImplicitFieldDefsType}:                   _GotoState233Action,
	{_State113, ImplicitStructDefType}:                   _GotoState122Action,
	{_State113, ExplicitStructDefType}:                   _GotoState74Action,
	{_State113, EnumValueDefType}:                        _GotoState230Action,
	{_State113, ImplicitEnumValueDefsType}:               _GotoState232Action,
	{_State113, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State113, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State113, TraitDefType}:                            _GotoState128Action,
	{_State113, FuncTypeType}:                            _GotoState120Action,
	{_State117, LparenToken}:                             _GotoState237Action,
	{_State125, IdentifierToken}:                         _GotoState112Action,
	{_State125, StructToken}:                             _GotoState50Action,
	{_State125, EnumToken}:                               _GotoState109Action,
	{_State125, TraitToken}:                              _GotoState117Action,
	{_State125, FuncToken}:                               _GotoState111Action,
	{_State125, LparenToken}:                             _GotoState113Action,
	{_State125, LbracketToken}:                           _GotoState42Action,
	{_State125, DotToken}:                                _GotoState108Action,
	{_State125, QuestionToken}:                           _GotoState115Action,
	{_State125, ExclaimToken}:                            _GotoState110Action,
	{_State125, TildeTildeToken}:                         _GotoState116Action,
	{_State125, BitNegToken}:                             _GotoState107Action,
	{_State125, BitAndToken}:                             _GotoState106Action,
	{_State125, ParseErrorToken}:                         _GotoState114Action,
	{_State125, InitializableTypeType}:                   _GotoState123Action,
	{_State125, SliceTypeType}:                           _GotoState100Action,
	{_State125, ArrayTypeType}:                           _GotoState59Action,
	{_State125, MapTypeType}:                             _GotoState88Action,
	{_State125, AtomTypeType}:                            _GotoState118Action,
	{_State125, ParseErrorTypeType}:                      _GotoState124Action,
	{_State125, ReturnableTypeType}:                      _GotoState238Action,
	{_State125, PrefixedTypeType}:                        _GotoState126Action,
	{_State125, PrefixTypeOpType}:                        _GotoState125Action,
	{_State125, ImplicitStructDefType}:                   _GotoState122Action,
	{_State125, ExplicitStructDefType}:                   _GotoState74Action,
	{_State125, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State125, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State125, TraitDefType}:                            _GotoState128Action,
	{_State125, FuncTypeType}:                            _GotoState120Action,
	{_State130, DollarLbracketToken}:                     _GotoState244Action,
	{_State130, AssignToken}:                             _GotoState243Action,
	{_State130, OptionalGenericParametersType}:           _GotoState245Action,
	{_State131, IdentifierToken}:                         _GotoState246Action,
	{_State131, ParameterDefType}:                        _GotoState247Action,
	{_State132, NewlinesToken}:                           _GotoState248Action,
	{_State132, SemicolonToken}:                          _GotoState249Action,
	{_State134, RbraceToken}:                             _GotoState250Action,
	{_State136, DollarLbracketToken}:                     _GotoState244Action,
	{_State136, AssignToken}:                             _GotoState251Action,
	{_State136, OptionalGenericParametersType}:           _GotoState252Action,
	{_State137, CommentGroupsToken}:                      _GotoState9Action,
	{_State137, PackageToken}:                            _GotoState13Action,
	{_State137, TypeToken}:                               _GotoState14Action,
	{_State137, FuncToken}:                               _GotoState10Action,
	{_State137, VarToken}:                                _GotoState15Action,
	{_State137, LetToken}:                                _GotoState12Action,
	{_State137, LbraceToken}:                             _GotoState11Action,
	{_State137, DefinitionType}:                          _GotoState253Action,
	{_State137, StatementBlockType}:                      _GotoState21Action,
	{_State137, VarDeclPatternType}:                      _GotoState23Action,
	{_State137, VarOrLetType}:                            _GotoState24Action,
	{_State137, TypeDefType}:                             _GotoState22Action,
	{_State137, NamedFuncDefType}:                        _GotoState18Action,
	{_State137, PackageDefType}:                          _GotoState19Action,
	{_State138, IntegerLiteralToken}:                     _GotoState40Action,
	{_State138, FloatLiteralToken}:                       _GotoState35Action,
	{_State138, RuneLiteralToken}:                        _GotoState48Action,
	{_State138, StringLiteralToken}:                      _GotoState49Action,
	{_State138, IdentifierToken}:                         _GotoState38Action,
	{_State138, TrueToken}:                               _GotoState52Action,
	{_State138, FalseToken}:                              _GotoState34Action,
	{_State138, StructToken}:                             _GotoState50Action,
	{_State138, FuncToken}:                               _GotoState36Action,
	{_State138, VarToken}:                                _GotoState15Action,
	{_State138, LetToken}:                                _GotoState12Action,
	{_State138, NotToken}:                                _GotoState45Action,
	{_State138, LabelDeclToken}:                          _GotoState41Action,
	{_State138, LparenToken}:                             _GotoState43Action,
	{_State138, LbracketToken}:                           _GotoState42Action,
	{_State138, SubToken}:                                _GotoState51Action,
	{_State138, MulToken}:                                _GotoState44Action,
	{_State138, BitNegToken}:                             _GotoState27Action,
	{_State138, BitAndToken}:                             _GotoState26Action,
	{_State138, GreaterToken}:                            _GotoState37Action,
	{_State138, ParseErrorToken}:                         _GotoState46Action,
	{_State138, VarDeclPatternType}:                      _GotoState103Action,
	{_State138, VarOrLetType}:                            _GotoState24Action,
	{_State138, ExpressionType}:                          _GotoState254Action,
	{_State138, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State138, SequenceExprType}:                        _GotoState105Action,
	{_State138, CallExprType}:                            _GotoState70Action,
	{_State138, AtomExprType}:                            _GotoState62Action,
	{_State138, ParseErrorExprType}:                      _GotoState92Action,
	{_State138, LiteralExprType}:                         _GotoState87Action,
	{_State138, IdentifierExprType}:                      _GotoState79Action,
	{_State138, BlockExprType}:                           _GotoState69Action,
	{_State138, InitializeExprType}:                      _GotoState84Action,
	{_State138, ImplicitStructExprType}:                  _GotoState80Action,
	{_State138, AccessibleExprType}:                      _GotoState104Action,
	{_State138, AccessExprType}:                          _GotoState54Action,
	{_State138, IndexExprType}:                           _GotoState82Action,
	{_State138, PostfixableExprType}:                     _GotoState94Action,
	{_State138, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State138, PrefixableExprType}:                      _GotoState97Action,
	{_State138, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State138, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State138, MulExprType}:                             _GotoState89Action,
	{_State138, BinaryMulExprType}:                       _GotoState66Action,
	{_State138, AddExprType}:                             _GotoState56Action,
	{_State138, BinaryAddExprType}:                       _GotoState63Action,
	{_State138, CmpExprType}:                             _GotoState73Action,
	{_State138, BinaryCmpExprType}:                       _GotoState65Action,
	{_State138, AndExprType}:                             _GotoState57Action,
	{_State138, BinaryAndExprType}:                       _GotoState64Action,
	{_State138, OrExprType}:                              _GotoState91Action,
	{_State138, BinaryOrExprType}:                        _GotoState68Action,
	{_State138, InitializableTypeType}:                   _GotoState83Action,
	{_State138, SliceTypeType}:                           _GotoState100Action,
	{_State138, ArrayTypeType}:                           _GotoState59Action,
	{_State138, MapTypeType}:                             _GotoState88Action,
	{_State138, ExplicitStructDefType}:                   _GotoState74Action,
	{_State138, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State140, IdentifierToken}:                         _GotoState256Action,
	{_State140, LparenToken}:                             _GotoState140Action,
	{_State140, EllipsisToken}:                           _GotoState255Action,
	{_State140, VarPatternType}:                          _GotoState259Action,
	{_State140, TuplePatternType}:                        _GotoState141Action,
	{_State140, FieldVarPatternsType}:                    _GotoState258Action,
	{_State140, FieldVarPatternType}:                     _GotoState257Action,
	{_State142, IdentifierToken}:                         _GotoState112Action,
	{_State142, StructToken}:                             _GotoState50Action,
	{_State142, EnumToken}:                               _GotoState109Action,
	{_State142, TraitToken}:                              _GotoState117Action,
	{_State142, FuncToken}:                               _GotoState111Action,
	{_State142, LparenToken}:                             _GotoState113Action,
	{_State142, LbracketToken}:                           _GotoState42Action,
	{_State142, DotToken}:                                _GotoState108Action,
	{_State142, QuestionToken}:                           _GotoState115Action,
	{_State142, ExclaimToken}:                            _GotoState110Action,
	{_State142, TildeTildeToken}:                         _GotoState116Action,
	{_State142, BitNegToken}:                             _GotoState107Action,
	{_State142, BitAndToken}:                             _GotoState106Action,
	{_State142, ParseErrorToken}:                         _GotoState114Action,
	{_State142, OptionalValueTypeType}:                   _GotoState260Action,
	{_State142, InitializableTypeType}:                   _GotoState123Action,
	{_State142, SliceTypeType}:                           _GotoState100Action,
	{_State142, ArrayTypeType}:                           _GotoState59Action,
	{_State142, MapTypeType}:                             _GotoState88Action,
	{_State142, AtomTypeType}:                            _GotoState118Action,
	{_State142, ParseErrorTypeType}:                      _GotoState124Action,
	{_State142, ReturnableTypeType}:                      _GotoState127Action,
	{_State142, PrefixedTypeType}:                        _GotoState126Action,
	{_State142, PrefixTypeOpType}:                        _GotoState125Action,
	{_State142, ValueTypeType}:                           _GotoState261Action,
	{_State142, TraitOpTypeType}:                         _GotoState129Action,
	{_State142, ImplicitStructDefType}:                   _GotoState122Action,
	{_State142, ExplicitStructDefType}:                   _GotoState74Action,
	{_State142, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State142, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State142, TraitDefType}:                            _GotoState128Action,
	{_State142, FuncTypeType}:                            _GotoState120Action,
	{_State143, IdentifierToken}:                         _GotoState262Action,
	{_State144, DotToken}:                                _GotoState263Action,
	{_State147, CommaToken}:                              _GotoState265Action,
	{_State147, ColonToken}:                              _GotoState264Action,
	{_State148, LbraceToken}:                             _GotoState11Action,
	{_State148, StatementBlockType}:                      _GotoState220Action,
	{_State150, IntegerLiteralToken}:                     _GotoState40Action,
	{_State150, FloatLiteralToken}:                       _GotoState35Action,
	{_State150, RuneLiteralToken}:                        _GotoState48Action,
	{_State150, StringLiteralToken}:                      _GotoState49Action,
	{_State150, IdentifierToken}:                         _GotoState38Action,
	{_State150, TrueToken}:                               _GotoState52Action,
	{_State150, FalseToken}:                              _GotoState34Action,
	{_State150, ReturnToken}:                             _GotoState47Action,
	{_State150, BreakToken}:                              _GotoState28Action,
	{_State150, ContinueToken}:                           _GotoState30Action,
	{_State150, FallthroughToken}:                        _GotoState33Action,
	{_State150, UnsafeToken}:                             _GotoState53Action,
	{_State150, StructToken}:                             _GotoState50Action,
	{_State150, FuncToken}:                               _GotoState36Action,
	{_State150, AsyncToken}:                              _GotoState25Action,
	{_State150, DeferToken}:                              _GotoState32Action,
	{_State150, VarToken}:                                _GotoState15Action,
	{_State150, LetToken}:                                _GotoState12Action,
	{_State150, NotToken}:                                _GotoState45Action,
	{_State150, LabelDeclToken}:                          _GotoState41Action,
	{_State150, LparenToken}:                             _GotoState43Action,
	{_State150, LbracketToken}:                           _GotoState42Action,
	{_State150, SubToken}:                                _GotoState51Action,
	{_State150, MulToken}:                                _GotoState44Action,
	{_State150, BitNegToken}:                             _GotoState27Action,
	{_State150, BitAndToken}:                             _GotoState26Action,
	{_State150, GreaterToken}:                            _GotoState37Action,
	{_State150, ParseErrorToken}:                         _GotoState46Action,
	{_State150, SimpleStatementType}:                     _GotoState267Action,
	{_State150, OptionalSimpleStatementType}:             _GotoState266Action,
	{_State150, ExpressionOrImproperStructStatementType}: _GotoState76Action,
	{_State150, ExpressionsType}:                         _GotoState77Action,
	{_State150, CallbackOpType}:                          _GotoState71Action,
	{_State150, CallbackOpStatementType}:                 _GotoState72Action,
	{_State150, UnsafeStatementType}:                     _GotoState102Action,
	{_State150, JumpStatementType}:                       _GotoState85Action,
	{_State150, JumpTypeType}:                            _GotoState86Action,
	{_State150, FallthroughStatementType}:                _GotoState78Action,
	{_State150, AssignStatementType}:                     _GotoState61Action,
	{_State150, UnaryOpAssignStatementType}:              _GotoState101Action,
	{_State150, BinaryOpAssignStatementType}:             _GotoState67Action,
	{_State150, VarDeclPatternType}:                      _GotoState103Action,
	{_State150, VarOrLetType}:                            _GotoState24Action,
	{_State150, AssignPatternType}:                       _GotoState60Action,
	{_State150, ExpressionType}:                          _GotoState75Action,
	{_State150, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State150, SequenceExprType}:                        _GotoState98Action,
	{_State150, CallExprType}:                            _GotoState70Action,
	{_State150, AtomExprType}:                            _GotoState62Action,
	{_State150, ParseErrorExprType}:                      _GotoState92Action,
	{_State150, LiteralExprType}:                         _GotoState87Action,
	{_State150, IdentifierExprType}:                      _GotoState79Action,
	{_State150, BlockExprType}:                           _GotoState69Action,
	{_State150, InitializeExprType}:                      _GotoState84Action,
	{_State150, ImplicitStructExprType}:                  _GotoState80Action,
	{_State150, AccessibleExprType}:                      _GotoState55Action,
	{_State150, AccessExprType}:                          _GotoState54Action,
	{_State150, IndexExprType}:                           _GotoState82Action,
	{_State150, PostfixableExprType}:                     _GotoState94Action,
	{_State150, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State150, PrefixableExprType}:                      _GotoState97Action,
	{_State150, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State150, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State150, MulExprType}:                             _GotoState89Action,
	{_State150, BinaryMulExprType}:                       _GotoState66Action,
	{_State150, AddExprType}:                             _GotoState56Action,
	{_State150, BinaryAddExprType}:                       _GotoState63Action,
	{_State150, CmpExprType}:                             _GotoState73Action,
	{_State150, BinaryCmpExprType}:                       _GotoState65Action,
	{_State150, AndExprType}:                             _GotoState57Action,
	{_State150, BinaryAndExprType}:                       _GotoState64Action,
	{_State150, OrExprType}:                              _GotoState91Action,
	{_State150, BinaryOrExprType}:                        _GotoState68Action,
	{_State150, InitializableTypeType}:                   _GotoState83Action,
	{_State150, SliceTypeType}:                           _GotoState100Action,
	{_State150, ArrayTypeType}:                           _GotoState59Action,
	{_State150, MapTypeType}:                             _GotoState88Action,
	{_State150, ExplicitStructDefType}:                   _GotoState74Action,
	{_State150, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State151, IdentifierToken}:                         _GotoState246Action,
	{_State151, ParameterDefType}:                        _GotoState268Action,
	{_State151, ProperParameterDefsType}:                 _GotoState270Action,
	{_State151, ParameterDefsType}:                       _GotoState269Action,
	{_State153, StringLiteralToken}:                      _GotoState154Action,
	{_State153, ImportClauseType}:                        _GotoState271Action,
	{_State153, ProperImportClausesType}:                 _GotoState273Action,
	{_State153, ImportClausesType}:                       _GotoState272Action,
	{_State154, AsToken}:                                 _GotoState274Action,
	{_State156, RbracketToken}:                           _GotoState277Action,
	{_State156, CommaToken}:                              _GotoState276Action,
	{_State156, ColonToken}:                              _GotoState275Action,
	{_State156, AddToken}:                                _GotoState239Action,
	{_State156, SubToken}:                                _GotoState241Action,
	{_State156, MulToken}:                                _GotoState240Action,
	{_State156, TraitOpType}:                             _GotoState242Action,
	{_State157, IntegerLiteralToken}:                     _GotoState40Action,
	{_State157, FloatLiteralToken}:                       _GotoState35Action,
	{_State157, RuneLiteralToken}:                        _GotoState48Action,
	{_State157, StringLiteralToken}:                      _GotoState49Action,
	{_State157, IdentifierToken}:                         _GotoState38Action,
	{_State157, TrueToken}:                               _GotoState52Action,
	{_State157, FalseToken}:                              _GotoState34Action,
	{_State157, StructToken}:                             _GotoState50Action,
	{_State157, FuncToken}:                               _GotoState36Action,
	{_State157, VarToken}:                                _GotoState15Action,
	{_State157, LetToken}:                                _GotoState12Action,
	{_State157, NotToken}:                                _GotoState45Action,
	{_State157, LabelDeclToken}:                          _GotoState41Action,
	{_State157, LparenToken}:                             _GotoState43Action,
	{_State157, LbracketToken}:                           _GotoState42Action,
	{_State157, SubToken}:                                _GotoState51Action,
	{_State157, MulToken}:                                _GotoState44Action,
	{_State157, BitNegToken}:                             _GotoState27Action,
	{_State157, BitAndToken}:                             _GotoState26Action,
	{_State157, GreaterToken}:                            _GotoState37Action,
	{_State157, ParseErrorToken}:                         _GotoState46Action,
	{_State157, VarDeclPatternType}:                      _GotoState103Action,
	{_State157, VarOrLetType}:                            _GotoState24Action,
	{_State157, ExpressionType}:                          _GotoState278Action,
	{_State157, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State157, SequenceExprType}:                        _GotoState105Action,
	{_State157, CallExprType}:                            _GotoState70Action,
	{_State157, AtomExprType}:                            _GotoState62Action,
	{_State157, ParseErrorExprType}:                      _GotoState92Action,
	{_State157, LiteralExprType}:                         _GotoState87Action,
	{_State157, IdentifierExprType}:                      _GotoState79Action,
	{_State157, BlockExprType}:                           _GotoState69Action,
	{_State157, InitializeExprType}:                      _GotoState84Action,
	{_State157, ImplicitStructExprType}:                  _GotoState80Action,
	{_State157, AccessibleExprType}:                      _GotoState104Action,
	{_State157, AccessExprType}:                          _GotoState54Action,
	{_State157, IndexExprType}:                           _GotoState82Action,
	{_State157, PostfixableExprType}:                     _GotoState94Action,
	{_State157, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State157, PrefixableExprType}:                      _GotoState97Action,
	{_State157, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State157, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State157, MulExprType}:                             _GotoState89Action,
	{_State157, BinaryMulExprType}:                       _GotoState66Action,
	{_State157, AddExprType}:                             _GotoState56Action,
	{_State157, BinaryAddExprType}:                       _GotoState63Action,
	{_State157, CmpExprType}:                             _GotoState73Action,
	{_State157, BinaryCmpExprType}:                       _GotoState65Action,
	{_State157, AndExprType}:                             _GotoState57Action,
	{_State157, BinaryAndExprType}:                       _GotoState64Action,
	{_State157, OrExprType}:                              _GotoState91Action,
	{_State157, BinaryOrExprType}:                        _GotoState68Action,
	{_State157, InitializableTypeType}:                   _GotoState83Action,
	{_State157, SliceTypeType}:                           _GotoState100Action,
	{_State157, ArrayTypeType}:                           _GotoState59Action,
	{_State157, MapTypeType}:                             _GotoState88Action,
	{_State157, ExplicitStructDefType}:                   _GotoState74Action,
	{_State157, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State159, AssignToken}:                             _GotoState279Action,
	{_State161, RparenToken}:                             _GotoState280Action,
	{_State162, ColonToken}:                              _GotoState281Action,
	{_State163, ColonToken}:                              _GotoState282Action,
	{_State163, EllipsisToken}:                           _GotoState283Action,
	{_State164, CommaToken}:                              _GotoState284Action,
	{_State165, IdentifierToken}:                         _GotoState229Action,
	{_State165, UnsafeToken}:                             _GotoState53Action,
	{_State165, StructToken}:                             _GotoState50Action,
	{_State165, EnumToken}:                               _GotoState109Action,
	{_State165, TraitToken}:                              _GotoState117Action,
	{_State165, FuncToken}:                               _GotoState111Action,
	{_State165, LparenToken}:                             _GotoState113Action,
	{_State165, LbracketToken}:                           _GotoState42Action,
	{_State165, DotToken}:                                _GotoState108Action,
	{_State165, QuestionToken}:                           _GotoState115Action,
	{_State165, ExclaimToken}:                            _GotoState110Action,
	{_State165, TildeTildeToken}:                         _GotoState116Action,
	{_State165, BitNegToken}:                             _GotoState107Action,
	{_State165, BitAndToken}:                             _GotoState106Action,
	{_State165, ParseErrorToken}:                         _GotoState114Action,
	{_State165, UnsafeStatementType}:                     _GotoState235Action,
	{_State165, InitializableTypeType}:                   _GotoState123Action,
	{_State165, SliceTypeType}:                           _GotoState100Action,
	{_State165, ArrayTypeType}:                           _GotoState59Action,
	{_State165, MapTypeType}:                             _GotoState88Action,
	{_State165, AtomTypeType}:                            _GotoState118Action,
	{_State165, ParseErrorTypeType}:                      _GotoState124Action,
	{_State165, ReturnableTypeType}:                      _GotoState127Action,
	{_State165, PrefixedTypeType}:                        _GotoState126Action,
	{_State165, PrefixTypeOpType}:                        _GotoState125Action,
	{_State165, ValueTypeType}:                           _GotoState236Action,
	{_State165, TraitOpTypeType}:                         _GotoState129Action,
	{_State165, FieldDefType}:                            _GotoState286Action,
	{_State165, ImplicitStructDefType}:                   _GotoState122Action,
	{_State165, ProperExplicitFieldDefsType}:             _GotoState287Action,
	{_State165, ExplicitFieldDefsType}:                   _GotoState285Action,
	{_State165, ExplicitStructDefType}:                   _GotoState74Action,
	{_State165, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State165, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State165, TraitDefType}:                            _GotoState128Action,
	{_State165, FuncTypeType}:                            _GotoState120Action,
	{_State166, IdentifierToken}:                         _GotoState288Action,
	{_State176, IdentifierToken}:                         _GotoState112Action,
	{_State176, StructToken}:                             _GotoState50Action,
	{_State176, EnumToken}:                               _GotoState109Action,
	{_State176, TraitToken}:                              _GotoState117Action,
	{_State176, FuncToken}:                               _GotoState111Action,
	{_State176, LparenToken}:                             _GotoState113Action,
	{_State176, LbracketToken}:                           _GotoState42Action,
	{_State176, DotToken}:                                _GotoState108Action,
	{_State176, QuestionToken}:                           _GotoState115Action,
	{_State176, ExclaimToken}:                            _GotoState110Action,
	{_State176, TildeTildeToken}:                         _GotoState116Action,
	{_State176, BitNegToken}:                             _GotoState107Action,
	{_State176, BitAndToken}:                             _GotoState106Action,
	{_State176, ParseErrorToken}:                         _GotoState114Action,
	{_State176, ProperGenericArgumentsType}:              _GotoState290Action,
	{_State176, GenericArgumentsType}:                    _GotoState289Action,
	{_State176, InitializableTypeType}:                   _GotoState123Action,
	{_State176, SliceTypeType}:                           _GotoState100Action,
	{_State176, ArrayTypeType}:                           _GotoState59Action,
	{_State176, MapTypeType}:                             _GotoState88Action,
	{_State176, AtomTypeType}:                            _GotoState118Action,
	{_State176, ParseErrorTypeType}:                      _GotoState124Action,
	{_State176, ReturnableTypeType}:                      _GotoState127Action,
	{_State176, PrefixedTypeType}:                        _GotoState126Action,
	{_State176, PrefixTypeOpType}:                        _GotoState125Action,
	{_State176, ValueTypeType}:                           _GotoState291Action,
	{_State176, TraitOpTypeType}:                         _GotoState129Action,
	{_State176, ImplicitStructDefType}:                   _GotoState122Action,
	{_State176, ExplicitStructDefType}:                   _GotoState74Action,
	{_State176, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State176, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State176, TraitDefType}:                            _GotoState128Action,
	{_State176, FuncTypeType}:                            _GotoState120Action,
	{_State177, IdentifierToken}:                         _GotoState292Action,
	{_State178, IntegerLiteralToken}:                     _GotoState40Action,
	{_State178, FloatLiteralToken}:                       _GotoState35Action,
	{_State178, RuneLiteralToken}:                        _GotoState48Action,
	{_State178, StringLiteralToken}:                      _GotoState49Action,
	{_State178, IdentifierToken}:                         _GotoState159Action,
	{_State178, TrueToken}:                               _GotoState52Action,
	{_State178, FalseToken}:                              _GotoState34Action,
	{_State178, StructToken}:                             _GotoState50Action,
	{_State178, FuncToken}:                               _GotoState36Action,
	{_State178, VarToken}:                                _GotoState15Action,
	{_State178, LetToken}:                                _GotoState12Action,
	{_State178, NotToken}:                                _GotoState45Action,
	{_State178, LabelDeclToken}:                          _GotoState41Action,
	{_State178, LparenToken}:                             _GotoState43Action,
	{_State178, LbracketToken}:                           _GotoState42Action,
	{_State178, ColonToken}:                              _GotoState157Action,
	{_State178, EllipsisToken}:                           _GotoState158Action,
	{_State178, SubToken}:                                _GotoState51Action,
	{_State178, MulToken}:                                _GotoState44Action,
	{_State178, BitNegToken}:                             _GotoState27Action,
	{_State178, BitAndToken}:                             _GotoState26Action,
	{_State178, GreaterToken}:                            _GotoState37Action,
	{_State178, ParseErrorToken}:                         _GotoState46Action,
	{_State178, VarDeclPatternType}:                      _GotoState103Action,
	{_State178, VarOrLetType}:                            _GotoState24Action,
	{_State178, ExpressionType}:                          _GotoState163Action,
	{_State178, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State178, SequenceExprType}:                        _GotoState105Action,
	{_State178, CallExprType}:                            _GotoState70Action,
	{_State178, ArgumentType}:                            _GotoState293Action,
	{_State178, ColonExprType}:                           _GotoState162Action,
	{_State178, AtomExprType}:                            _GotoState62Action,
	{_State178, ParseErrorExprType}:                      _GotoState92Action,
	{_State178, LiteralExprType}:                         _GotoState87Action,
	{_State178, IdentifierExprType}:                      _GotoState79Action,
	{_State178, BlockExprType}:                           _GotoState69Action,
	{_State178, InitializeExprType}:                      _GotoState84Action,
	{_State178, ImplicitStructExprType}:                  _GotoState80Action,
	{_State178, AccessibleExprType}:                      _GotoState104Action,
	{_State178, AccessExprType}:                          _GotoState54Action,
	{_State178, IndexExprType}:                           _GotoState82Action,
	{_State178, PostfixableExprType}:                     _GotoState94Action,
	{_State178, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State178, PrefixableExprType}:                      _GotoState97Action,
	{_State178, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State178, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State178, MulExprType}:                             _GotoState89Action,
	{_State178, BinaryMulExprType}:                       _GotoState66Action,
	{_State178, AddExprType}:                             _GotoState56Action,
	{_State178, BinaryAddExprType}:                       _GotoState63Action,
	{_State178, CmpExprType}:                             _GotoState73Action,
	{_State178, BinaryCmpExprType}:                       _GotoState65Action,
	{_State178, AndExprType}:                             _GotoState57Action,
	{_State178, BinaryAndExprType}:                       _GotoState64Action,
	{_State178, OrExprType}:                              _GotoState91Action,
	{_State178, BinaryOrExprType}:                        _GotoState68Action,
	{_State178, InitializableTypeType}:                   _GotoState83Action,
	{_State178, SliceTypeType}:                           _GotoState100Action,
	{_State178, ArrayTypeType}:                           _GotoState59Action,
	{_State178, MapTypeType}:                             _GotoState88Action,
	{_State178, ExplicitStructDefType}:                   _GotoState74Action,
	{_State178, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State184, IntegerLiteralToken}:                     _GotoState40Action,
	{_State184, FloatLiteralToken}:                       _GotoState35Action,
	{_State184, RuneLiteralToken}:                        _GotoState48Action,
	{_State184, StringLiteralToken}:                      _GotoState49Action,
	{_State184, IdentifierToken}:                         _GotoState38Action,
	{_State184, TrueToken}:                               _GotoState52Action,
	{_State184, FalseToken}:                              _GotoState34Action,
	{_State184, StructToken}:                             _GotoState50Action,
	{_State184, FuncToken}:                               _GotoState36Action,
	{_State184, VarToken}:                                _GotoState15Action,
	{_State184, LetToken}:                                _GotoState12Action,
	{_State184, NotToken}:                                _GotoState45Action,
	{_State184, LabelDeclToken}:                          _GotoState41Action,
	{_State184, LparenToken}:                             _GotoState43Action,
	{_State184, LbracketToken}:                           _GotoState42Action,
	{_State184, SubToken}:                                _GotoState51Action,
	{_State184, MulToken}:                                _GotoState44Action,
	{_State184, BitNegToken}:                             _GotoState27Action,
	{_State184, BitAndToken}:                             _GotoState26Action,
	{_State184, GreaterToken}:                            _GotoState37Action,
	{_State184, ParseErrorToken}:                         _GotoState46Action,
	{_State184, VarDeclPatternType}:                      _GotoState103Action,
	{_State184, VarOrLetType}:                            _GotoState24Action,
	{_State184, ExpressionType}:                          _GotoState294Action,
	{_State184, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State184, SequenceExprType}:                        _GotoState105Action,
	{_State184, CallExprType}:                            _GotoState70Action,
	{_State184, AtomExprType}:                            _GotoState62Action,
	{_State184, ParseErrorExprType}:                      _GotoState92Action,
	{_State184, LiteralExprType}:                         _GotoState87Action,
	{_State184, IdentifierExprType}:                      _GotoState79Action,
	{_State184, BlockExprType}:                           _GotoState69Action,
	{_State184, InitializeExprType}:                      _GotoState84Action,
	{_State184, ImplicitStructExprType}:                  _GotoState80Action,
	{_State184, AccessibleExprType}:                      _GotoState104Action,
	{_State184, AccessExprType}:                          _GotoState54Action,
	{_State184, IndexExprType}:                           _GotoState82Action,
	{_State184, PostfixableExprType}:                     _GotoState94Action,
	{_State184, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State184, PrefixableExprType}:                      _GotoState97Action,
	{_State184, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State184, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State184, MulExprType}:                             _GotoState89Action,
	{_State184, BinaryMulExprType}:                       _GotoState66Action,
	{_State184, AddExprType}:                             _GotoState56Action,
	{_State184, BinaryAddExprType}:                       _GotoState63Action,
	{_State184, CmpExprType}:                             _GotoState73Action,
	{_State184, BinaryCmpExprType}:                       _GotoState65Action,
	{_State184, AndExprType}:                             _GotoState57Action,
	{_State184, BinaryAndExprType}:                       _GotoState64Action,
	{_State184, OrExprType}:                              _GotoState91Action,
	{_State184, BinaryOrExprType}:                        _GotoState68Action,
	{_State184, InitializableTypeType}:                   _GotoState83Action,
	{_State184, SliceTypeType}:                           _GotoState100Action,
	{_State184, ArrayTypeType}:                           _GotoState59Action,
	{_State184, MapTypeType}:                             _GotoState88Action,
	{_State184, ExplicitStructDefType}:                   _GotoState74Action,
	{_State184, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State185, LparenToken}:                             _GotoState295Action,
	{_State191, IntegerLiteralToken}:                     _GotoState40Action,
	{_State191, FloatLiteralToken}:                       _GotoState35Action,
	{_State191, RuneLiteralToken}:                        _GotoState48Action,
	{_State191, StringLiteralToken}:                      _GotoState49Action,
	{_State191, IdentifierToken}:                         _GotoState38Action,
	{_State191, TrueToken}:                               _GotoState52Action,
	{_State191, FalseToken}:                              _GotoState34Action,
	{_State191, StructToken}:                             _GotoState50Action,
	{_State191, FuncToken}:                               _GotoState36Action,
	{_State191, NotToken}:                                _GotoState45Action,
	{_State191, LabelDeclToken}:                          _GotoState41Action,
	{_State191, LparenToken}:                             _GotoState43Action,
	{_State191, LbracketToken}:                           _GotoState42Action,
	{_State191, SubToken}:                                _GotoState51Action,
	{_State191, MulToken}:                                _GotoState44Action,
	{_State191, BitNegToken}:                             _GotoState27Action,
	{_State191, BitAndToken}:                             _GotoState26Action,
	{_State191, ParseErrorToken}:                         _GotoState46Action,
	{_State191, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State191, CallExprType}:                            _GotoState70Action,
	{_State191, AtomExprType}:                            _GotoState62Action,
	{_State191, ParseErrorExprType}:                      _GotoState92Action,
	{_State191, LiteralExprType}:                         _GotoState87Action,
	{_State191, IdentifierExprType}:                      _GotoState79Action,
	{_State191, BlockExprType}:                           _GotoState69Action,
	{_State191, InitializeExprType}:                      _GotoState84Action,
	{_State191, ImplicitStructExprType}:                  _GotoState80Action,
	{_State191, AccessibleExprType}:                      _GotoState104Action,
	{_State191, AccessExprType}:                          _GotoState54Action,
	{_State191, IndexExprType}:                           _GotoState82Action,
	{_State191, PostfixableExprType}:                     _GotoState94Action,
	{_State191, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State191, PrefixableExprType}:                      _GotoState97Action,
	{_State191, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State191, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State191, MulExprType}:                             _GotoState296Action,
	{_State191, BinaryMulExprType}:                       _GotoState66Action,
	{_State191, InitializableTypeType}:                   _GotoState83Action,
	{_State191, SliceTypeType}:                           _GotoState100Action,
	{_State191, ArrayTypeType}:                           _GotoState59Action,
	{_State191, MapTypeType}:                             _GotoState88Action,
	{_State191, ExplicitStructDefType}:                   _GotoState74Action,
	{_State191, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State192, IntegerLiteralToken}:                     _GotoState40Action,
	{_State192, FloatLiteralToken}:                       _GotoState35Action,
	{_State192, RuneLiteralToken}:                        _GotoState48Action,
	{_State192, StringLiteralToken}:                      _GotoState49Action,
	{_State192, IdentifierToken}:                         _GotoState38Action,
	{_State192, TrueToken}:                               _GotoState52Action,
	{_State192, FalseToken}:                              _GotoState34Action,
	{_State192, StructToken}:                             _GotoState50Action,
	{_State192, FuncToken}:                               _GotoState36Action,
	{_State192, NotToken}:                                _GotoState45Action,
	{_State192, LabelDeclToken}:                          _GotoState41Action,
	{_State192, LparenToken}:                             _GotoState43Action,
	{_State192, LbracketToken}:                           _GotoState42Action,
	{_State192, SubToken}:                                _GotoState51Action,
	{_State192, MulToken}:                                _GotoState44Action,
	{_State192, BitNegToken}:                             _GotoState27Action,
	{_State192, BitAndToken}:                             _GotoState26Action,
	{_State192, ParseErrorToken}:                         _GotoState46Action,
	{_State192, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State192, CallExprType}:                            _GotoState70Action,
	{_State192, AtomExprType}:                            _GotoState62Action,
	{_State192, ParseErrorExprType}:                      _GotoState92Action,
	{_State192, LiteralExprType}:                         _GotoState87Action,
	{_State192, IdentifierExprType}:                      _GotoState79Action,
	{_State192, BlockExprType}:                           _GotoState69Action,
	{_State192, InitializeExprType}:                      _GotoState84Action,
	{_State192, ImplicitStructExprType}:                  _GotoState80Action,
	{_State192, AccessibleExprType}:                      _GotoState104Action,
	{_State192, AccessExprType}:                          _GotoState54Action,
	{_State192, IndexExprType}:                           _GotoState82Action,
	{_State192, PostfixableExprType}:                     _GotoState94Action,
	{_State192, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State192, PrefixableExprType}:                      _GotoState97Action,
	{_State192, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State192, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State192, MulExprType}:                             _GotoState89Action,
	{_State192, BinaryMulExprType}:                       _GotoState66Action,
	{_State192, AddExprType}:                             _GotoState56Action,
	{_State192, BinaryAddExprType}:                       _GotoState63Action,
	{_State192, CmpExprType}:                             _GotoState297Action,
	{_State192, BinaryCmpExprType}:                       _GotoState65Action,
	{_State192, InitializableTypeType}:                   _GotoState83Action,
	{_State192, SliceTypeType}:                           _GotoState100Action,
	{_State192, ArrayTypeType}:                           _GotoState59Action,
	{_State192, MapTypeType}:                             _GotoState88Action,
	{_State192, ExplicitStructDefType}:                   _GotoState74Action,
	{_State192, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State193, IntegerLiteralToken}:                     _GotoState40Action,
	{_State193, FloatLiteralToken}:                       _GotoState35Action,
	{_State193, RuneLiteralToken}:                        _GotoState48Action,
	{_State193, StringLiteralToken}:                      _GotoState49Action,
	{_State193, IdentifierToken}:                         _GotoState38Action,
	{_State193, TrueToken}:                               _GotoState52Action,
	{_State193, FalseToken}:                              _GotoState34Action,
	{_State193, StructToken}:                             _GotoState50Action,
	{_State193, FuncToken}:                               _GotoState36Action,
	{_State193, VarToken}:                                _GotoState15Action,
	{_State193, LetToken}:                                _GotoState12Action,
	{_State193, NotToken}:                                _GotoState45Action,
	{_State193, LabelDeclToken}:                          _GotoState41Action,
	{_State193, LparenToken}:                             _GotoState43Action,
	{_State193, LbracketToken}:                           _GotoState42Action,
	{_State193, SubToken}:                                _GotoState51Action,
	{_State193, MulToken}:                                _GotoState44Action,
	{_State193, BitNegToken}:                             _GotoState27Action,
	{_State193, BitAndToken}:                             _GotoState26Action,
	{_State193, GreaterToken}:                            _GotoState37Action,
	{_State193, ParseErrorToken}:                         _GotoState46Action,
	{_State193, VarDeclPatternType}:                      _GotoState103Action,
	{_State193, VarOrLetType}:                            _GotoState24Action,
	{_State193, ExpressionType}:                          _GotoState298Action,
	{_State193, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State193, SequenceExprType}:                        _GotoState105Action,
	{_State193, CallExprType}:                            _GotoState70Action,
	{_State193, AtomExprType}:                            _GotoState62Action,
	{_State193, ParseErrorExprType}:                      _GotoState92Action,
	{_State193, LiteralExprType}:                         _GotoState87Action,
	{_State193, IdentifierExprType}:                      _GotoState79Action,
	{_State193, BlockExprType}:                           _GotoState69Action,
	{_State193, InitializeExprType}:                      _GotoState84Action,
	{_State193, ImplicitStructExprType}:                  _GotoState80Action,
	{_State193, AccessibleExprType}:                      _GotoState104Action,
	{_State193, AccessExprType}:                          _GotoState54Action,
	{_State193, IndexExprType}:                           _GotoState82Action,
	{_State193, PostfixableExprType}:                     _GotoState94Action,
	{_State193, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State193, PrefixableExprType}:                      _GotoState97Action,
	{_State193, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State193, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State193, MulExprType}:                             _GotoState89Action,
	{_State193, BinaryMulExprType}:                       _GotoState66Action,
	{_State193, AddExprType}:                             _GotoState56Action,
	{_State193, BinaryAddExprType}:                       _GotoState63Action,
	{_State193, CmpExprType}:                             _GotoState73Action,
	{_State193, BinaryCmpExprType}:                       _GotoState65Action,
	{_State193, AndExprType}:                             _GotoState57Action,
	{_State193, BinaryAndExprType}:                       _GotoState64Action,
	{_State193, OrExprType}:                              _GotoState91Action,
	{_State193, BinaryOrExprType}:                        _GotoState68Action,
	{_State193, InitializableTypeType}:                   _GotoState83Action,
	{_State193, SliceTypeType}:                           _GotoState100Action,
	{_State193, ArrayTypeType}:                           _GotoState59Action,
	{_State193, MapTypeType}:                             _GotoState88Action,
	{_State193, ExplicitStructDefType}:                   _GotoState74Action,
	{_State193, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State194, LbracketToken}:                           _GotoState178Action,
	{_State194, DotToken}:                                _GotoState177Action,
	{_State194, DollarLbracketToken}:                     _GotoState176Action,
	{_State194, OptionalGenericBindingType}:              _GotoState185Action,
	{_State202, IntegerLiteralToken}:                     _GotoState40Action,
	{_State202, FloatLiteralToken}:                       _GotoState35Action,
	{_State202, RuneLiteralToken}:                        _GotoState48Action,
	{_State202, StringLiteralToken}:                      _GotoState49Action,
	{_State202, IdentifierToken}:                         _GotoState38Action,
	{_State202, TrueToken}:                               _GotoState52Action,
	{_State202, FalseToken}:                              _GotoState34Action,
	{_State202, StructToken}:                             _GotoState50Action,
	{_State202, FuncToken}:                               _GotoState36Action,
	{_State202, NotToken}:                                _GotoState45Action,
	{_State202, LabelDeclToken}:                          _GotoState41Action,
	{_State202, LparenToken}:                             _GotoState43Action,
	{_State202, LbracketToken}:                           _GotoState42Action,
	{_State202, SubToken}:                                _GotoState51Action,
	{_State202, MulToken}:                                _GotoState44Action,
	{_State202, BitNegToken}:                             _GotoState27Action,
	{_State202, BitAndToken}:                             _GotoState26Action,
	{_State202, ParseErrorToken}:                         _GotoState46Action,
	{_State202, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State202, CallExprType}:                            _GotoState70Action,
	{_State202, AtomExprType}:                            _GotoState62Action,
	{_State202, ParseErrorExprType}:                      _GotoState92Action,
	{_State202, LiteralExprType}:                         _GotoState87Action,
	{_State202, IdentifierExprType}:                      _GotoState79Action,
	{_State202, BlockExprType}:                           _GotoState69Action,
	{_State202, InitializeExprType}:                      _GotoState84Action,
	{_State202, ImplicitStructExprType}:                  _GotoState80Action,
	{_State202, AccessibleExprType}:                      _GotoState104Action,
	{_State202, AccessExprType}:                          _GotoState54Action,
	{_State202, IndexExprType}:                           _GotoState82Action,
	{_State202, PostfixableExprType}:                     _GotoState94Action,
	{_State202, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State202, PrefixableExprType}:                      _GotoState97Action,
	{_State202, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State202, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State202, MulExprType}:                             _GotoState89Action,
	{_State202, BinaryMulExprType}:                       _GotoState66Action,
	{_State202, AddExprType}:                             _GotoState299Action,
	{_State202, BinaryAddExprType}:                       _GotoState63Action,
	{_State202, InitializableTypeType}:                   _GotoState83Action,
	{_State202, SliceTypeType}:                           _GotoState100Action,
	{_State202, ArrayTypeType}:                           _GotoState59Action,
	{_State202, MapTypeType}:                             _GotoState88Action,
	{_State202, ExplicitStructDefType}:                   _GotoState74Action,
	{_State202, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State203, IntegerLiteralToken}:                     _GotoState40Action,
	{_State203, FloatLiteralToken}:                       _GotoState35Action,
	{_State203, RuneLiteralToken}:                        _GotoState48Action,
	{_State203, StringLiteralToken}:                      _GotoState49Action,
	{_State203, IdentifierToken}:                         _GotoState38Action,
	{_State203, TrueToken}:                               _GotoState52Action,
	{_State203, FalseToken}:                              _GotoState34Action,
	{_State203, StructToken}:                             _GotoState50Action,
	{_State203, FuncToken}:                               _GotoState36Action,
	{_State203, VarToken}:                                _GotoState15Action,
	{_State203, LetToken}:                                _GotoState12Action,
	{_State203, NotToken}:                                _GotoState45Action,
	{_State203, LabelDeclToken}:                          _GotoState41Action,
	{_State203, LparenToken}:                             _GotoState43Action,
	{_State203, LbracketToken}:                           _GotoState42Action,
	{_State203, SubToken}:                                _GotoState51Action,
	{_State203, MulToken}:                                _GotoState44Action,
	{_State203, BitNegToken}:                             _GotoState27Action,
	{_State203, BitAndToken}:                             _GotoState26Action,
	{_State203, GreaterToken}:                            _GotoState37Action,
	{_State203, ParseErrorToken}:                         _GotoState46Action,
	{_State203, VarDeclPatternType}:                      _GotoState103Action,
	{_State203, VarOrLetType}:                            _GotoState24Action,
	{_State203, ExpressionType}:                          _GotoState300Action,
	{_State203, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State203, SequenceExprType}:                        _GotoState105Action,
	{_State203, CallExprType}:                            _GotoState70Action,
	{_State203, AtomExprType}:                            _GotoState62Action,
	{_State203, ParseErrorExprType}:                      _GotoState92Action,
	{_State203, LiteralExprType}:                         _GotoState87Action,
	{_State203, IdentifierExprType}:                      _GotoState79Action,
	{_State203, BlockExprType}:                           _GotoState69Action,
	{_State203, InitializeExprType}:                      _GotoState84Action,
	{_State203, ImplicitStructExprType}:                  _GotoState80Action,
	{_State203, AccessibleExprType}:                      _GotoState104Action,
	{_State203, AccessExprType}:                          _GotoState54Action,
	{_State203, IndexExprType}:                           _GotoState82Action,
	{_State203, PostfixableExprType}:                     _GotoState94Action,
	{_State203, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State203, PrefixableExprType}:                      _GotoState97Action,
	{_State203, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State203, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State203, MulExprType}:                             _GotoState89Action,
	{_State203, BinaryMulExprType}:                       _GotoState66Action,
	{_State203, AddExprType}:                             _GotoState56Action,
	{_State203, BinaryAddExprType}:                       _GotoState63Action,
	{_State203, CmpExprType}:                             _GotoState73Action,
	{_State203, BinaryCmpExprType}:                       _GotoState65Action,
	{_State203, AndExprType}:                             _GotoState57Action,
	{_State203, BinaryAndExprType}:                       _GotoState64Action,
	{_State203, OrExprType}:                              _GotoState91Action,
	{_State203, BinaryOrExprType}:                        _GotoState68Action,
	{_State203, InitializableTypeType}:                   _GotoState83Action,
	{_State203, SliceTypeType}:                           _GotoState100Action,
	{_State203, ArrayTypeType}:                           _GotoState59Action,
	{_State203, MapTypeType}:                             _GotoState88Action,
	{_State203, ExplicitStructDefType}:                   _GotoState74Action,
	{_State203, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State204, IntegerLiteralToken}:                     _GotoState40Action,
	{_State204, FloatLiteralToken}:                       _GotoState35Action,
	{_State204, RuneLiteralToken}:                        _GotoState48Action,
	{_State204, StringLiteralToken}:                      _GotoState49Action,
	{_State204, IdentifierToken}:                         _GotoState159Action,
	{_State204, TrueToken}:                               _GotoState52Action,
	{_State204, FalseToken}:                              _GotoState34Action,
	{_State204, StructToken}:                             _GotoState50Action,
	{_State204, FuncToken}:                               _GotoState36Action,
	{_State204, VarToken}:                                _GotoState15Action,
	{_State204, LetToken}:                                _GotoState12Action,
	{_State204, NotToken}:                                _GotoState45Action,
	{_State204, LabelDeclToken}:                          _GotoState41Action,
	{_State204, LparenToken}:                             _GotoState43Action,
	{_State204, LbracketToken}:                           _GotoState42Action,
	{_State204, ColonToken}:                              _GotoState157Action,
	{_State204, EllipsisToken}:                           _GotoState158Action,
	{_State204, SubToken}:                                _GotoState51Action,
	{_State204, MulToken}:                                _GotoState44Action,
	{_State204, BitNegToken}:                             _GotoState27Action,
	{_State204, BitAndToken}:                             _GotoState26Action,
	{_State204, GreaterToken}:                            _GotoState37Action,
	{_State204, ParseErrorToken}:                         _GotoState46Action,
	{_State204, VarDeclPatternType}:                      _GotoState103Action,
	{_State204, VarOrLetType}:                            _GotoState24Action,
	{_State204, ExpressionType}:                          _GotoState163Action,
	{_State204, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State204, SequenceExprType}:                        _GotoState105Action,
	{_State204, CallExprType}:                            _GotoState70Action,
	{_State204, ProperArgumentsType}:                     _GotoState164Action,
	{_State204, ArgumentsType}:                           _GotoState301Action,
	{_State204, ArgumentType}:                            _GotoState160Action,
	{_State204, ColonExprType}:                           _GotoState162Action,
	{_State204, AtomExprType}:                            _GotoState62Action,
	{_State204, ParseErrorExprType}:                      _GotoState92Action,
	{_State204, LiteralExprType}:                         _GotoState87Action,
	{_State204, IdentifierExprType}:                      _GotoState79Action,
	{_State204, BlockExprType}:                           _GotoState69Action,
	{_State204, InitializeExprType}:                      _GotoState84Action,
	{_State204, ImplicitStructExprType}:                  _GotoState80Action,
	{_State204, AccessibleExprType}:                      _GotoState104Action,
	{_State204, AccessExprType}:                          _GotoState54Action,
	{_State204, IndexExprType}:                           _GotoState82Action,
	{_State204, PostfixableExprType}:                     _GotoState94Action,
	{_State204, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State204, PrefixableExprType}:                      _GotoState97Action,
	{_State204, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State204, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State204, MulExprType}:                             _GotoState89Action,
	{_State204, BinaryMulExprType}:                       _GotoState66Action,
	{_State204, AddExprType}:                             _GotoState56Action,
	{_State204, BinaryAddExprType}:                       _GotoState63Action,
	{_State204, CmpExprType}:                             _GotoState73Action,
	{_State204, BinaryCmpExprType}:                       _GotoState65Action,
	{_State204, AndExprType}:                             _GotoState57Action,
	{_State204, BinaryAndExprType}:                       _GotoState64Action,
	{_State204, OrExprType}:                              _GotoState91Action,
	{_State204, BinaryOrExprType}:                        _GotoState68Action,
	{_State204, InitializableTypeType}:                   _GotoState83Action,
	{_State204, SliceTypeType}:                           _GotoState100Action,
	{_State204, ArrayTypeType}:                           _GotoState59Action,
	{_State204, MapTypeType}:                             _GotoState88Action,
	{_State204, ExplicitStructDefType}:                   _GotoState74Action,
	{_State204, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State205, IntegerLiteralToken}:                     _GotoState40Action,
	{_State205, FloatLiteralToken}:                       _GotoState35Action,
	{_State205, RuneLiteralToken}:                        _GotoState48Action,
	{_State205, StringLiteralToken}:                      _GotoState49Action,
	{_State205, IdentifierToken}:                         _GotoState38Action,
	{_State205, TrueToken}:                               _GotoState52Action,
	{_State205, FalseToken}:                              _GotoState34Action,
	{_State205, StructToken}:                             _GotoState50Action,
	{_State205, FuncToken}:                               _GotoState36Action,
	{_State205, VarToken}:                                _GotoState15Action,
	{_State205, LetToken}:                                _GotoState12Action,
	{_State205, NotToken}:                                _GotoState45Action,
	{_State205, LabelDeclToken}:                          _GotoState41Action,
	{_State205, LparenToken}:                             _GotoState43Action,
	{_State205, LbracketToken}:                           _GotoState42Action,
	{_State205, SubToken}:                                _GotoState51Action,
	{_State205, MulToken}:                                _GotoState44Action,
	{_State205, BitNegToken}:                             _GotoState27Action,
	{_State205, BitAndToken}:                             _GotoState26Action,
	{_State205, GreaterToken}:                            _GotoState37Action,
	{_State205, ParseErrorToken}:                         _GotoState46Action,
	{_State205, ExpressionsType}:                         _GotoState302Action,
	{_State205, VarDeclPatternType}:                      _GotoState103Action,
	{_State205, VarOrLetType}:                            _GotoState24Action,
	{_State205, ExpressionType}:                          _GotoState75Action,
	{_State205, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State205, SequenceExprType}:                        _GotoState105Action,
	{_State205, CallExprType}:                            _GotoState70Action,
	{_State205, AtomExprType}:                            _GotoState62Action,
	{_State205, ParseErrorExprType}:                      _GotoState92Action,
	{_State205, LiteralExprType}:                         _GotoState87Action,
	{_State205, IdentifierExprType}:                      _GotoState79Action,
	{_State205, BlockExprType}:                           _GotoState69Action,
	{_State205, InitializeExprType}:                      _GotoState84Action,
	{_State205, ImplicitStructExprType}:                  _GotoState80Action,
	{_State205, AccessibleExprType}:                      _GotoState104Action,
	{_State205, AccessExprType}:                          _GotoState54Action,
	{_State205, IndexExprType}:                           _GotoState82Action,
	{_State205, PostfixableExprType}:                     _GotoState94Action,
	{_State205, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State205, PrefixableExprType}:                      _GotoState97Action,
	{_State205, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State205, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State205, MulExprType}:                             _GotoState89Action,
	{_State205, BinaryMulExprType}:                       _GotoState66Action,
	{_State205, AddExprType}:                             _GotoState56Action,
	{_State205, BinaryAddExprType}:                       _GotoState63Action,
	{_State205, CmpExprType}:                             _GotoState73Action,
	{_State205, BinaryCmpExprType}:                       _GotoState65Action,
	{_State205, AndExprType}:                             _GotoState57Action,
	{_State205, BinaryAndExprType}:                       _GotoState64Action,
	{_State205, OrExprType}:                              _GotoState91Action,
	{_State205, BinaryOrExprType}:                        _GotoState68Action,
	{_State205, InitializableTypeType}:                   _GotoState83Action,
	{_State205, SliceTypeType}:                           _GotoState100Action,
	{_State205, ArrayTypeType}:                           _GotoState59Action,
	{_State205, MapTypeType}:                             _GotoState88Action,
	{_State205, ExplicitStructDefType}:                   _GotoState74Action,
	{_State205, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State206, CommaToken}:                              _GotoState203Action,
	{_State213, IntegerLiteralToken}:                     _GotoState40Action,
	{_State213, FloatLiteralToken}:                       _GotoState35Action,
	{_State213, RuneLiteralToken}:                        _GotoState48Action,
	{_State213, StringLiteralToken}:                      _GotoState49Action,
	{_State213, IdentifierToken}:                         _GotoState38Action,
	{_State213, TrueToken}:                               _GotoState52Action,
	{_State213, FalseToken}:                              _GotoState34Action,
	{_State213, StructToken}:                             _GotoState50Action,
	{_State213, FuncToken}:                               _GotoState36Action,
	{_State213, NotToken}:                                _GotoState45Action,
	{_State213, LabelDeclToken}:                          _GotoState41Action,
	{_State213, LparenToken}:                             _GotoState43Action,
	{_State213, LbracketToken}:                           _GotoState42Action,
	{_State213, SubToken}:                                _GotoState51Action,
	{_State213, MulToken}:                                _GotoState44Action,
	{_State213, BitNegToken}:                             _GotoState27Action,
	{_State213, BitAndToken}:                             _GotoState26Action,
	{_State213, ParseErrorToken}:                         _GotoState46Action,
	{_State213, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State213, CallExprType}:                            _GotoState70Action,
	{_State213, AtomExprType}:                            _GotoState62Action,
	{_State213, ParseErrorExprType}:                      _GotoState92Action,
	{_State213, LiteralExprType}:                         _GotoState87Action,
	{_State213, IdentifierExprType}:                      _GotoState79Action,
	{_State213, BlockExprType}:                           _GotoState69Action,
	{_State213, InitializeExprType}:                      _GotoState84Action,
	{_State213, ImplicitStructExprType}:                  _GotoState80Action,
	{_State213, AccessibleExprType}:                      _GotoState104Action,
	{_State213, AccessExprType}:                          _GotoState54Action,
	{_State213, IndexExprType}:                           _GotoState82Action,
	{_State213, PostfixableExprType}:                     _GotoState94Action,
	{_State213, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State213, PrefixableExprType}:                      _GotoState303Action,
	{_State213, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State213, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State213, InitializableTypeType}:                   _GotoState83Action,
	{_State213, SliceTypeType}:                           _GotoState100Action,
	{_State213, ArrayTypeType}:                           _GotoState59Action,
	{_State213, MapTypeType}:                             _GotoState88Action,
	{_State213, ExplicitStructDefType}:                   _GotoState74Action,
	{_State213, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State214, LbraceToken}:                             _GotoState11Action,
	{_State214, StatementBlockType}:                      _GotoState304Action,
	{_State215, IntegerLiteralToken}:                     _GotoState40Action,
	{_State215, FloatLiteralToken}:                       _GotoState35Action,
	{_State215, RuneLiteralToken}:                        _GotoState48Action,
	{_State215, StringLiteralToken}:                      _GotoState49Action,
	{_State215, IdentifierToken}:                         _GotoState38Action,
	{_State215, TrueToken}:                               _GotoState52Action,
	{_State215, FalseToken}:                              _GotoState34Action,
	{_State215, StructToken}:                             _GotoState50Action,
	{_State215, FuncToken}:                               _GotoState36Action,
	{_State215, VarToken}:                                _GotoState15Action,
	{_State215, LetToken}:                                _GotoState12Action,
	{_State215, NotToken}:                                _GotoState45Action,
	{_State215, LabelDeclToken}:                          _GotoState41Action,
	{_State215, LparenToken}:                             _GotoState43Action,
	{_State215, LbracketToken}:                           _GotoState42Action,
	{_State215, SubToken}:                                _GotoState51Action,
	{_State215, MulToken}:                                _GotoState44Action,
	{_State215, BitNegToken}:                             _GotoState27Action,
	{_State215, BitAndToken}:                             _GotoState26Action,
	{_State215, GreaterToken}:                            _GotoState37Action,
	{_State215, ParseErrorToken}:                         _GotoState46Action,
	{_State215, VarDeclPatternType}:                      _GotoState103Action,
	{_State215, VarOrLetType}:                            _GotoState24Action,
	{_State215, AssignPatternType}:                       _GotoState305Action,
	{_State215, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State215, SequenceExprType}:                        _GotoState307Action,
	{_State215, ForAssignmentType}:                       _GotoState306Action,
	{_State215, CallExprType}:                            _GotoState70Action,
	{_State215, AtomExprType}:                            _GotoState62Action,
	{_State215, ParseErrorExprType}:                      _GotoState92Action,
	{_State215, LiteralExprType}:                         _GotoState87Action,
	{_State215, IdentifierExprType}:                      _GotoState79Action,
	{_State215, BlockExprType}:                           _GotoState69Action,
	{_State215, InitializeExprType}:                      _GotoState84Action,
	{_State215, ImplicitStructExprType}:                  _GotoState80Action,
	{_State215, AccessibleExprType}:                      _GotoState104Action,
	{_State215, AccessExprType}:                          _GotoState54Action,
	{_State215, IndexExprType}:                           _GotoState82Action,
	{_State215, PostfixableExprType}:                     _GotoState94Action,
	{_State215, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State215, PrefixableExprType}:                      _GotoState97Action,
	{_State215, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State215, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State215, MulExprType}:                             _GotoState89Action,
	{_State215, BinaryMulExprType}:                       _GotoState66Action,
	{_State215, AddExprType}:                             _GotoState56Action,
	{_State215, BinaryAddExprType}:                       _GotoState63Action,
	{_State215, CmpExprType}:                             _GotoState73Action,
	{_State215, BinaryCmpExprType}:                       _GotoState65Action,
	{_State215, AndExprType}:                             _GotoState57Action,
	{_State215, BinaryAndExprType}:                       _GotoState64Action,
	{_State215, OrExprType}:                              _GotoState91Action,
	{_State215, BinaryOrExprType}:                        _GotoState68Action,
	{_State215, InitializableTypeType}:                   _GotoState83Action,
	{_State215, SliceTypeType}:                           _GotoState100Action,
	{_State215, ArrayTypeType}:                           _GotoState59Action,
	{_State215, MapTypeType}:                             _GotoState88Action,
	{_State215, ExplicitStructDefType}:                   _GotoState74Action,
	{_State215, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State216, IntegerLiteralToken}:                     _GotoState40Action,
	{_State216, FloatLiteralToken}:                       _GotoState35Action,
	{_State216, RuneLiteralToken}:                        _GotoState48Action,
	{_State216, StringLiteralToken}:                      _GotoState49Action,
	{_State216, IdentifierToken}:                         _GotoState38Action,
	{_State216, TrueToken}:                               _GotoState52Action,
	{_State216, FalseToken}:                              _GotoState34Action,
	{_State216, CaseToken}:                               _GotoState308Action,
	{_State216, StructToken}:                             _GotoState50Action,
	{_State216, FuncToken}:                               _GotoState36Action,
	{_State216, VarToken}:                                _GotoState15Action,
	{_State216, LetToken}:                                _GotoState12Action,
	{_State216, NotToken}:                                _GotoState45Action,
	{_State216, LabelDeclToken}:                          _GotoState41Action,
	{_State216, LparenToken}:                             _GotoState43Action,
	{_State216, LbracketToken}:                           _GotoState42Action,
	{_State216, SubToken}:                                _GotoState51Action,
	{_State216, MulToken}:                                _GotoState44Action,
	{_State216, BitNegToken}:                             _GotoState27Action,
	{_State216, BitAndToken}:                             _GotoState26Action,
	{_State216, GreaterToken}:                            _GotoState37Action,
	{_State216, ParseErrorToken}:                         _GotoState46Action,
	{_State216, VarDeclPatternType}:                      _GotoState103Action,
	{_State216, VarOrLetType}:                            _GotoState24Action,
	{_State216, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State216, SequenceExprType}:                        _GotoState310Action,
	{_State216, ConditionType}:                           _GotoState309Action,
	{_State216, CallExprType}:                            _GotoState70Action,
	{_State216, AtomExprType}:                            _GotoState62Action,
	{_State216, ParseErrorExprType}:                      _GotoState92Action,
	{_State216, LiteralExprType}:                         _GotoState87Action,
	{_State216, IdentifierExprType}:                      _GotoState79Action,
	{_State216, BlockExprType}:                           _GotoState69Action,
	{_State216, InitializeExprType}:                      _GotoState84Action,
	{_State216, ImplicitStructExprType}:                  _GotoState80Action,
	{_State216, AccessibleExprType}:                      _GotoState104Action,
	{_State216, AccessExprType}:                          _GotoState54Action,
	{_State216, IndexExprType}:                           _GotoState82Action,
	{_State216, PostfixableExprType}:                     _GotoState94Action,
	{_State216, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State216, PrefixableExprType}:                      _GotoState97Action,
	{_State216, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State216, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State216, MulExprType}:                             _GotoState89Action,
	{_State216, BinaryMulExprType}:                       _GotoState66Action,
	{_State216, AddExprType}:                             _GotoState56Action,
	{_State216, BinaryAddExprType}:                       _GotoState63Action,
	{_State216, CmpExprType}:                             _GotoState73Action,
	{_State216, BinaryCmpExprType}:                       _GotoState65Action,
	{_State216, AndExprType}:                             _GotoState57Action,
	{_State216, BinaryAndExprType}:                       _GotoState64Action,
	{_State216, OrExprType}:                              _GotoState91Action,
	{_State216, BinaryOrExprType}:                        _GotoState68Action,
	{_State216, InitializableTypeType}:                   _GotoState83Action,
	{_State216, SliceTypeType}:                           _GotoState100Action,
	{_State216, ArrayTypeType}:                           _GotoState59Action,
	{_State216, MapTypeType}:                             _GotoState88Action,
	{_State216, ExplicitStructDefType}:                   _GotoState74Action,
	{_State216, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State217, IntegerLiteralToken}:                     _GotoState40Action,
	{_State217, FloatLiteralToken}:                       _GotoState35Action,
	{_State217, RuneLiteralToken}:                        _GotoState48Action,
	{_State217, StringLiteralToken}:                      _GotoState49Action,
	{_State217, IdentifierToken}:                         _GotoState38Action,
	{_State217, TrueToken}:                               _GotoState52Action,
	{_State217, FalseToken}:                              _GotoState34Action,
	{_State217, StructToken}:                             _GotoState50Action,
	{_State217, FuncToken}:                               _GotoState36Action,
	{_State217, VarToken}:                                _GotoState15Action,
	{_State217, LetToken}:                                _GotoState12Action,
	{_State217, NotToken}:                                _GotoState45Action,
	{_State217, LabelDeclToken}:                          _GotoState41Action,
	{_State217, LparenToken}:                             _GotoState43Action,
	{_State217, LbracketToken}:                           _GotoState42Action,
	{_State217, SubToken}:                                _GotoState51Action,
	{_State217, MulToken}:                                _GotoState44Action,
	{_State217, BitNegToken}:                             _GotoState27Action,
	{_State217, BitAndToken}:                             _GotoState26Action,
	{_State217, GreaterToken}:                            _GotoState37Action,
	{_State217, ParseErrorToken}:                         _GotoState46Action,
	{_State217, VarDeclPatternType}:                      _GotoState103Action,
	{_State217, VarOrLetType}:                            _GotoState24Action,
	{_State217, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State217, SequenceExprType}:                        _GotoState311Action,
	{_State217, CallExprType}:                            _GotoState70Action,
	{_State217, AtomExprType}:                            _GotoState62Action,
	{_State217, ParseErrorExprType}:                      _GotoState92Action,
	{_State217, LiteralExprType}:                         _GotoState87Action,
	{_State217, IdentifierExprType}:                      _GotoState79Action,
	{_State217, BlockExprType}:                           _GotoState69Action,
	{_State217, InitializeExprType}:                      _GotoState84Action,
	{_State217, ImplicitStructExprType}:                  _GotoState80Action,
	{_State217, AccessibleExprType}:                      _GotoState104Action,
	{_State217, AccessExprType}:                          _GotoState54Action,
	{_State217, IndexExprType}:                           _GotoState82Action,
	{_State217, PostfixableExprType}:                     _GotoState94Action,
	{_State217, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State217, PrefixableExprType}:                      _GotoState97Action,
	{_State217, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State217, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State217, MulExprType}:                             _GotoState89Action,
	{_State217, BinaryMulExprType}:                       _GotoState66Action,
	{_State217, AddExprType}:                             _GotoState56Action,
	{_State217, BinaryAddExprType}:                       _GotoState63Action,
	{_State217, CmpExprType}:                             _GotoState73Action,
	{_State217, BinaryCmpExprType}:                       _GotoState65Action,
	{_State217, AndExprType}:                             _GotoState57Action,
	{_State217, BinaryAndExprType}:                       _GotoState64Action,
	{_State217, OrExprType}:                              _GotoState91Action,
	{_State217, BinaryOrExprType}:                        _GotoState68Action,
	{_State217, InitializableTypeType}:                   _GotoState83Action,
	{_State217, SliceTypeType}:                           _GotoState100Action,
	{_State217, ArrayTypeType}:                           _GotoState59Action,
	{_State217, MapTypeType}:                             _GotoState88Action,
	{_State217, ExplicitStructDefType}:                   _GotoState74Action,
	{_State217, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State222, IntegerLiteralToken}:                     _GotoState40Action,
	{_State222, FloatLiteralToken}:                       _GotoState35Action,
	{_State222, RuneLiteralToken}:                        _GotoState48Action,
	{_State222, StringLiteralToken}:                      _GotoState49Action,
	{_State222, IdentifierToken}:                         _GotoState38Action,
	{_State222, TrueToken}:                               _GotoState52Action,
	{_State222, FalseToken}:                              _GotoState34Action,
	{_State222, StructToken}:                             _GotoState50Action,
	{_State222, FuncToken}:                               _GotoState36Action,
	{_State222, NotToken}:                                _GotoState45Action,
	{_State222, LabelDeclToken}:                          _GotoState41Action,
	{_State222, LparenToken}:                             _GotoState43Action,
	{_State222, LbracketToken}:                           _GotoState42Action,
	{_State222, SubToken}:                                _GotoState51Action,
	{_State222, MulToken}:                                _GotoState44Action,
	{_State222, BitNegToken}:                             _GotoState27Action,
	{_State222, BitAndToken}:                             _GotoState26Action,
	{_State222, ParseErrorToken}:                         _GotoState46Action,
	{_State222, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State222, CallExprType}:                            _GotoState70Action,
	{_State222, AtomExprType}:                            _GotoState62Action,
	{_State222, ParseErrorExprType}:                      _GotoState92Action,
	{_State222, LiteralExprType}:                         _GotoState87Action,
	{_State222, IdentifierExprType}:                      _GotoState79Action,
	{_State222, BlockExprType}:                           _GotoState69Action,
	{_State222, InitializeExprType}:                      _GotoState84Action,
	{_State222, ImplicitStructExprType}:                  _GotoState80Action,
	{_State222, AccessibleExprType}:                      _GotoState104Action,
	{_State222, AccessExprType}:                          _GotoState54Action,
	{_State222, IndexExprType}:                           _GotoState82Action,
	{_State222, PostfixableExprType}:                     _GotoState94Action,
	{_State222, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State222, PrefixableExprType}:                      _GotoState97Action,
	{_State222, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State222, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State222, MulExprType}:                             _GotoState89Action,
	{_State222, BinaryMulExprType}:                       _GotoState66Action,
	{_State222, AddExprType}:                             _GotoState56Action,
	{_State222, BinaryAddExprType}:                       _GotoState63Action,
	{_State222, CmpExprType}:                             _GotoState73Action,
	{_State222, BinaryCmpExprType}:                       _GotoState65Action,
	{_State222, AndExprType}:                             _GotoState312Action,
	{_State222, BinaryAndExprType}:                       _GotoState64Action,
	{_State222, InitializableTypeType}:                   _GotoState83Action,
	{_State222, SliceTypeType}:                           _GotoState100Action,
	{_State222, ArrayTypeType}:                           _GotoState59Action,
	{_State222, MapTypeType}:                             _GotoState88Action,
	{_State222, ExplicitStructDefType}:                   _GotoState74Action,
	{_State222, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State225, IdentifierToken}:                         _GotoState229Action,
	{_State225, UnsafeToken}:                             _GotoState53Action,
	{_State225, StructToken}:                             _GotoState50Action,
	{_State225, EnumToken}:                               _GotoState109Action,
	{_State225, TraitToken}:                              _GotoState117Action,
	{_State225, FuncToken}:                               _GotoState111Action,
	{_State225, LparenToken}:                             _GotoState113Action,
	{_State225, LbracketToken}:                           _GotoState42Action,
	{_State225, DotToken}:                                _GotoState108Action,
	{_State225, QuestionToken}:                           _GotoState115Action,
	{_State225, ExclaimToken}:                            _GotoState110Action,
	{_State225, TildeTildeToken}:                         _GotoState116Action,
	{_State225, BitNegToken}:                             _GotoState107Action,
	{_State225, BitAndToken}:                             _GotoState106Action,
	{_State225, ParseErrorToken}:                         _GotoState114Action,
	{_State225, UnsafeStatementType}:                     _GotoState235Action,
	{_State225, InitializableTypeType}:                   _GotoState123Action,
	{_State225, SliceTypeType}:                           _GotoState100Action,
	{_State225, ArrayTypeType}:                           _GotoState59Action,
	{_State225, MapTypeType}:                             _GotoState88Action,
	{_State225, AtomTypeType}:                            _GotoState118Action,
	{_State225, ParseErrorTypeType}:                      _GotoState124Action,
	{_State225, ReturnableTypeType}:                      _GotoState127Action,
	{_State225, PrefixedTypeType}:                        _GotoState126Action,
	{_State225, PrefixTypeOpType}:                        _GotoState125Action,
	{_State225, ValueTypeType}:                           _GotoState236Action,
	{_State225, TraitOpTypeType}:                         _GotoState129Action,
	{_State225, FieldDefType}:                            _GotoState315Action,
	{_State225, ImplicitStructDefType}:                   _GotoState122Action,
	{_State225, ExplicitStructDefType}:                   _GotoState74Action,
	{_State225, EnumValueDefType}:                        _GotoState313Action,
	{_State225, ImplicitEnumValueDefsType}:               _GotoState316Action,
	{_State225, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State225, ExplicitEnumValueDefsType}:               _GotoState314Action,
	{_State225, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State225, TraitDefType}:                            _GotoState128Action,
	{_State225, FuncTypeType}:                            _GotoState120Action,
	{_State226, IdentifierToken}:                         _GotoState318Action,
	{_State226, StructToken}:                             _GotoState50Action,
	{_State226, EnumToken}:                               _GotoState109Action,
	{_State226, TraitToken}:                              _GotoState117Action,
	{_State226, FuncToken}:                               _GotoState111Action,
	{_State226, LparenToken}:                             _GotoState113Action,
	{_State226, LbracketToken}:                           _GotoState42Action,
	{_State226, DotToken}:                                _GotoState108Action,
	{_State226, QuestionToken}:                           _GotoState115Action,
	{_State226, ExclaimToken}:                            _GotoState110Action,
	{_State226, EllipsisToken}:                           _GotoState317Action,
	{_State226, TildeTildeToken}:                         _GotoState116Action,
	{_State226, BitNegToken}:                             _GotoState107Action,
	{_State226, BitAndToken}:                             _GotoState106Action,
	{_State226, ParseErrorToken}:                         _GotoState114Action,
	{_State226, InitializableTypeType}:                   _GotoState123Action,
	{_State226, SliceTypeType}:                           _GotoState100Action,
	{_State226, ArrayTypeType}:                           _GotoState59Action,
	{_State226, MapTypeType}:                             _GotoState88Action,
	{_State226, AtomTypeType}:                            _GotoState118Action,
	{_State226, ParseErrorTypeType}:                      _GotoState124Action,
	{_State226, ReturnableTypeType}:                      _GotoState127Action,
	{_State226, PrefixedTypeType}:                        _GotoState126Action,
	{_State226, PrefixTypeOpType}:                        _GotoState125Action,
	{_State226, ValueTypeType}:                           _GotoState322Action,
	{_State226, TraitOpTypeType}:                         _GotoState129Action,
	{_State226, ImplicitStructDefType}:                   _GotoState122Action,
	{_State226, ExplicitStructDefType}:                   _GotoState74Action,
	{_State226, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State226, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State226, TraitDefType}:                            _GotoState128Action,
	{_State226, ParameterDeclType}:                       _GotoState319Action,
	{_State226, ProperParameterDeclsType}:                _GotoState321Action,
	{_State226, ParameterDeclsType}:                      _GotoState320Action,
	{_State226, FuncTypeType}:                            _GotoState120Action,
	{_State227, IdentifierToken}:                         _GotoState323Action,
	{_State229, IdentifierToken}:                         _GotoState112Action,
	{_State229, StructToken}:                             _GotoState50Action,
	{_State229, EnumToken}:                               _GotoState109Action,
	{_State229, TraitToken}:                              _GotoState117Action,
	{_State229, FuncToken}:                               _GotoState111Action,
	{_State229, LparenToken}:                             _GotoState113Action,
	{_State229, LbracketToken}:                           _GotoState42Action,
	{_State229, DotToken}:                                _GotoState324Action,
	{_State229, QuestionToken}:                           _GotoState115Action,
	{_State229, ExclaimToken}:                            _GotoState110Action,
	{_State229, DollarLbracketToken}:                     _GotoState176Action,
	{_State229, TildeTildeToken}:                         _GotoState116Action,
	{_State229, BitNegToken}:                             _GotoState107Action,
	{_State229, BitAndToken}:                             _GotoState106Action,
	{_State229, ParseErrorToken}:                         _GotoState114Action,
	{_State229, OptionalGenericBindingType}:              _GotoState228Action,
	{_State229, InitializableTypeType}:                   _GotoState123Action,
	{_State229, SliceTypeType}:                           _GotoState100Action,
	{_State229, ArrayTypeType}:                           _GotoState59Action,
	{_State229, MapTypeType}:                             _GotoState88Action,
	{_State229, AtomTypeType}:                            _GotoState118Action,
	{_State229, ParseErrorTypeType}:                      _GotoState124Action,
	{_State229, ReturnableTypeType}:                      _GotoState127Action,
	{_State229, PrefixedTypeType}:                        _GotoState126Action,
	{_State229, PrefixTypeOpType}:                        _GotoState125Action,
	{_State229, ValueTypeType}:                           _GotoState325Action,
	{_State229, TraitOpTypeType}:                         _GotoState129Action,
	{_State229, ImplicitStructDefType}:                   _GotoState122Action,
	{_State229, ExplicitStructDefType}:                   _GotoState74Action,
	{_State229, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State229, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State229, TraitDefType}:                            _GotoState128Action,
	{_State229, FuncTypeType}:                            _GotoState120Action,
	{_State230, OrToken}:                                 _GotoState326Action,
	{_State231, AssignToken}:                             _GotoState327Action,
	{_State232, OrToken}:                                 _GotoState328Action,
	{_State232, RparenToken}:                             _GotoState329Action,
	{_State233, RparenToken}:                             _GotoState330Action,
	{_State234, CommaToken}:                              _GotoState331Action,
	{_State236, AddToken}:                                _GotoState239Action,
	{_State236, SubToken}:                                _GotoState241Action,
	{_State236, MulToken}:                                _GotoState240Action,
	{_State236, TraitOpType}:                             _GotoState242Action,
	{_State237, IdentifierToken}:                         _GotoState229Action,
	{_State237, UnsafeToken}:                             _GotoState53Action,
	{_State237, StructToken}:                             _GotoState50Action,
	{_State237, EnumToken}:                               _GotoState109Action,
	{_State237, TraitToken}:                              _GotoState117Action,
	{_State237, FuncToken}:                               _GotoState332Action,
	{_State237, LparenToken}:                             _GotoState113Action,
	{_State237, LbracketToken}:                           _GotoState42Action,
	{_State237, DotToken}:                                _GotoState108Action,
	{_State237, QuestionToken}:                           _GotoState115Action,
	{_State237, ExclaimToken}:                            _GotoState110Action,
	{_State237, TildeTildeToken}:                         _GotoState116Action,
	{_State237, BitNegToken}:                             _GotoState107Action,
	{_State237, BitAndToken}:                             _GotoState106Action,
	{_State237, ParseErrorToken}:                         _GotoState114Action,
	{_State237, UnsafeStatementType}:                     _GotoState235Action,
	{_State237, InitializableTypeType}:                   _GotoState123Action,
	{_State237, SliceTypeType}:                           _GotoState100Action,
	{_State237, ArrayTypeType}:                           _GotoState59Action,
	{_State237, MapTypeType}:                             _GotoState88Action,
	{_State237, AtomTypeType}:                            _GotoState118Action,
	{_State237, ParseErrorTypeType}:                      _GotoState124Action,
	{_State237, ReturnableTypeType}:                      _GotoState127Action,
	{_State237, PrefixedTypeType}:                        _GotoState126Action,
	{_State237, PrefixTypeOpType}:                        _GotoState125Action,
	{_State237, ValueTypeType}:                           _GotoState236Action,
	{_State237, TraitOpTypeType}:                         _GotoState129Action,
	{_State237, FieldDefType}:                            _GotoState333Action,
	{_State237, ImplicitStructDefType}:                   _GotoState122Action,
	{_State237, ExplicitStructDefType}:                   _GotoState74Action,
	{_State237, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State237, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State237, TraitPropertyType}:                       _GotoState337Action,
	{_State237, ProperTraitPropertiesType}:               _GotoState335Action,
	{_State237, TraitPropertiesType}:                     _GotoState336Action,
	{_State237, TraitDefType}:                            _GotoState128Action,
	{_State237, FuncTypeType}:                            _GotoState120Action,
	{_State237, MethodSignatureType}:                     _GotoState334Action,
	{_State242, IdentifierToken}:                         _GotoState112Action,
	{_State242, StructToken}:                             _GotoState50Action,
	{_State242, EnumToken}:                               _GotoState109Action,
	{_State242, TraitToken}:                              _GotoState117Action,
	{_State242, FuncToken}:                               _GotoState111Action,
	{_State242, LparenToken}:                             _GotoState113Action,
	{_State242, LbracketToken}:                           _GotoState42Action,
	{_State242, DotToken}:                                _GotoState108Action,
	{_State242, QuestionToken}:                           _GotoState115Action,
	{_State242, ExclaimToken}:                            _GotoState110Action,
	{_State242, TildeTildeToken}:                         _GotoState116Action,
	{_State242, BitNegToken}:                             _GotoState107Action,
	{_State242, BitAndToken}:                             _GotoState106Action,
	{_State242, ParseErrorToken}:                         _GotoState114Action,
	{_State242, InitializableTypeType}:                   _GotoState123Action,
	{_State242, SliceTypeType}:                           _GotoState100Action,
	{_State242, ArrayTypeType}:                           _GotoState59Action,
	{_State242, MapTypeType}:                             _GotoState88Action,
	{_State242, AtomTypeType}:                            _GotoState118Action,
	{_State242, ParseErrorTypeType}:                      _GotoState124Action,
	{_State242, ReturnableTypeType}:                      _GotoState338Action,
	{_State242, PrefixedTypeType}:                        _GotoState126Action,
	{_State242, PrefixTypeOpType}:                        _GotoState125Action,
	{_State242, ImplicitStructDefType}:                   _GotoState122Action,
	{_State242, ExplicitStructDefType}:                   _GotoState74Action,
	{_State242, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State242, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State242, TraitDefType}:                            _GotoState128Action,
	{_State242, FuncTypeType}:                            _GotoState120Action,
	{_State243, IntegerLiteralToken}:                     _GotoState40Action,
	{_State243, FloatLiteralToken}:                       _GotoState35Action,
	{_State243, RuneLiteralToken}:                        _GotoState48Action,
	{_State243, StringLiteralToken}:                      _GotoState49Action,
	{_State243, IdentifierToken}:                         _GotoState38Action,
	{_State243, TrueToken}:                               _GotoState52Action,
	{_State243, FalseToken}:                              _GotoState34Action,
	{_State243, StructToken}:                             _GotoState50Action,
	{_State243, FuncToken}:                               _GotoState36Action,
	{_State243, VarToken}:                                _GotoState15Action,
	{_State243, LetToken}:                                _GotoState12Action,
	{_State243, NotToken}:                                _GotoState45Action,
	{_State243, LabelDeclToken}:                          _GotoState41Action,
	{_State243, LparenToken}:                             _GotoState43Action,
	{_State243, LbracketToken}:                           _GotoState42Action,
	{_State243, SubToken}:                                _GotoState51Action,
	{_State243, MulToken}:                                _GotoState44Action,
	{_State243, BitNegToken}:                             _GotoState27Action,
	{_State243, BitAndToken}:                             _GotoState26Action,
	{_State243, GreaterToken}:                            _GotoState37Action,
	{_State243, ParseErrorToken}:                         _GotoState46Action,
	{_State243, VarDeclPatternType}:                      _GotoState103Action,
	{_State243, VarOrLetType}:                            _GotoState24Action,
	{_State243, ExpressionType}:                          _GotoState339Action,
	{_State243, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State243, SequenceExprType}:                        _GotoState105Action,
	{_State243, CallExprType}:                            _GotoState70Action,
	{_State243, AtomExprType}:                            _GotoState62Action,
	{_State243, ParseErrorExprType}:                      _GotoState92Action,
	{_State243, LiteralExprType}:                         _GotoState87Action,
	{_State243, IdentifierExprType}:                      _GotoState79Action,
	{_State243, BlockExprType}:                           _GotoState69Action,
	{_State243, InitializeExprType}:                      _GotoState84Action,
	{_State243, ImplicitStructExprType}:                  _GotoState80Action,
	{_State243, AccessibleExprType}:                      _GotoState104Action,
	{_State243, AccessExprType}:                          _GotoState54Action,
	{_State243, IndexExprType}:                           _GotoState82Action,
	{_State243, PostfixableExprType}:                     _GotoState94Action,
	{_State243, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State243, PrefixableExprType}:                      _GotoState97Action,
	{_State243, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State243, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State243, MulExprType}:                             _GotoState89Action,
	{_State243, BinaryMulExprType}:                       _GotoState66Action,
	{_State243, AddExprType}:                             _GotoState56Action,
	{_State243, BinaryAddExprType}:                       _GotoState63Action,
	{_State243, CmpExprType}:                             _GotoState73Action,
	{_State243, BinaryCmpExprType}:                       _GotoState65Action,
	{_State243, AndExprType}:                             _GotoState57Action,
	{_State243, BinaryAndExprType}:                       _GotoState64Action,
	{_State243, OrExprType}:                              _GotoState91Action,
	{_State243, BinaryOrExprType}:                        _GotoState68Action,
	{_State243, InitializableTypeType}:                   _GotoState83Action,
	{_State243, SliceTypeType}:                           _GotoState100Action,
	{_State243, ArrayTypeType}:                           _GotoState59Action,
	{_State243, MapTypeType}:                             _GotoState88Action,
	{_State243, ExplicitStructDefType}:                   _GotoState74Action,
	{_State243, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State244, IdentifierToken}:                         _GotoState340Action,
	{_State244, GenericParameterDefType}:                 _GotoState341Action,
	{_State244, ProperGenericParameterDefsType}:          _GotoState343Action,
	{_State244, GenericParameterDefsType}:                _GotoState342Action,
	{_State245, LparenToken}:                             _GotoState344Action,
	{_State246, IdentifierToken}:                         _GotoState112Action,
	{_State246, StructToken}:                             _GotoState50Action,
	{_State246, EnumToken}:                               _GotoState109Action,
	{_State246, TraitToken}:                              _GotoState117Action,
	{_State246, FuncToken}:                               _GotoState111Action,
	{_State246, LparenToken}:                             _GotoState113Action,
	{_State246, LbracketToken}:                           _GotoState42Action,
	{_State246, DotToken}:                                _GotoState108Action,
	{_State246, QuestionToken}:                           _GotoState115Action,
	{_State246, ExclaimToken}:                            _GotoState110Action,
	{_State246, EllipsisToken}:                           _GotoState345Action,
	{_State246, TildeTildeToken}:                         _GotoState116Action,
	{_State246, BitNegToken}:                             _GotoState107Action,
	{_State246, BitAndToken}:                             _GotoState106Action,
	{_State246, ParseErrorToken}:                         _GotoState114Action,
	{_State246, InitializableTypeType}:                   _GotoState123Action,
	{_State246, SliceTypeType}:                           _GotoState100Action,
	{_State246, ArrayTypeType}:                           _GotoState59Action,
	{_State246, MapTypeType}:                             _GotoState88Action,
	{_State246, AtomTypeType}:                            _GotoState118Action,
	{_State246, ParseErrorTypeType}:                      _GotoState124Action,
	{_State246, ReturnableTypeType}:                      _GotoState127Action,
	{_State246, PrefixedTypeType}:                        _GotoState126Action,
	{_State246, PrefixTypeOpType}:                        _GotoState125Action,
	{_State246, ValueTypeType}:                           _GotoState346Action,
	{_State246, TraitOpTypeType}:                         _GotoState129Action,
	{_State246, ImplicitStructDefType}:                   _GotoState122Action,
	{_State246, ExplicitStructDefType}:                   _GotoState74Action,
	{_State246, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State246, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State246, TraitDefType}:                            _GotoState128Action,
	{_State246, FuncTypeType}:                            _GotoState120Action,
	{_State247, RparenToken}:                             _GotoState347Action,
	{_State248, IntegerLiteralToken}:                     _GotoState40Action,
	{_State248, FloatLiteralToken}:                       _GotoState35Action,
	{_State248, RuneLiteralToken}:                        _GotoState48Action,
	{_State248, StringLiteralToken}:                      _GotoState49Action,
	{_State248, IdentifierToken}:                         _GotoState38Action,
	{_State248, TrueToken}:                               _GotoState52Action,
	{_State248, FalseToken}:                              _GotoState34Action,
	{_State248, CaseToken}:                               _GotoState29Action,
	{_State248, DefaultToken}:                            _GotoState31Action,
	{_State248, ReturnToken}:                             _GotoState47Action,
	{_State248, BreakToken}:                              _GotoState28Action,
	{_State248, ContinueToken}:                           _GotoState30Action,
	{_State248, FallthroughToken}:                        _GotoState33Action,
	{_State248, ImportToken}:                             _GotoState39Action,
	{_State248, UnsafeToken}:                             _GotoState53Action,
	{_State248, StructToken}:                             _GotoState50Action,
	{_State248, FuncToken}:                               _GotoState36Action,
	{_State248, AsyncToken}:                              _GotoState25Action,
	{_State248, DeferToken}:                              _GotoState32Action,
	{_State248, VarToken}:                                _GotoState15Action,
	{_State248, LetToken}:                                _GotoState12Action,
	{_State248, NotToken}:                                _GotoState45Action,
	{_State248, LabelDeclToken}:                          _GotoState41Action,
	{_State248, LparenToken}:                             _GotoState43Action,
	{_State248, LbracketToken}:                           _GotoState42Action,
	{_State248, SubToken}:                                _GotoState51Action,
	{_State248, MulToken}:                                _GotoState44Action,
	{_State248, BitNegToken}:                             _GotoState27Action,
	{_State248, BitAndToken}:                             _GotoState26Action,
	{_State248, GreaterToken}:                            _GotoState37Action,
	{_State248, ParseErrorToken}:                         _GotoState46Action,
	{_State248, SimpleStatementType}:                     _GotoState99Action,
	{_State248, StatementType}:                           _GotoState348Action,
	{_State248, ExpressionOrImproperStructStatementType}: _GotoState76Action,
	{_State248, ExpressionsType}:                         _GotoState77Action,
	{_State248, CallbackOpType}:                          _GotoState71Action,
	{_State248, CallbackOpStatementType}:                 _GotoState72Action,
	{_State248, UnsafeStatementType}:                     _GotoState102Action,
	{_State248, JumpStatementType}:                       _GotoState85Action,
	{_State248, JumpTypeType}:                            _GotoState86Action,
	{_State248, FallthroughStatementType}:                _GotoState78Action,
	{_State248, AssignStatementType}:                     _GotoState61Action,
	{_State248, UnaryOpAssignStatementType}:              _GotoState101Action,
	{_State248, BinaryOpAssignStatementType}:             _GotoState67Action,
	{_State248, ImportStatementType}:                     _GotoState81Action,
	{_State248, VarDeclPatternType}:                      _GotoState103Action,
	{_State248, VarOrLetType}:                            _GotoState24Action,
	{_State248, AssignPatternType}:                       _GotoState60Action,
	{_State248, ExpressionType}:                          _GotoState75Action,
	{_State248, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State248, SequenceExprType}:                        _GotoState98Action,
	{_State248, CallExprType}:                            _GotoState70Action,
	{_State248, AtomExprType}:                            _GotoState62Action,
	{_State248, ParseErrorExprType}:                      _GotoState92Action,
	{_State248, LiteralExprType}:                         _GotoState87Action,
	{_State248, IdentifierExprType}:                      _GotoState79Action,
	{_State248, BlockExprType}:                           _GotoState69Action,
	{_State248, InitializeExprType}:                      _GotoState84Action,
	{_State248, ImplicitStructExprType}:                  _GotoState80Action,
	{_State248, AccessibleExprType}:                      _GotoState55Action,
	{_State248, AccessExprType}:                          _GotoState54Action,
	{_State248, IndexExprType}:                           _GotoState82Action,
	{_State248, PostfixableExprType}:                     _GotoState94Action,
	{_State248, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State248, PrefixableExprType}:                      _GotoState97Action,
	{_State248, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State248, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State248, MulExprType}:                             _GotoState89Action,
	{_State248, BinaryMulExprType}:                       _GotoState66Action,
	{_State248, AddExprType}:                             _GotoState56Action,
	{_State248, BinaryAddExprType}:                       _GotoState63Action,
	{_State248, CmpExprType}:                             _GotoState73Action,
	{_State248, BinaryCmpExprType}:                       _GotoState65Action,
	{_State248, AndExprType}:                             _GotoState57Action,
	{_State248, BinaryAndExprType}:                       _GotoState64Action,
	{_State248, OrExprType}:                              _GotoState91Action,
	{_State248, BinaryOrExprType}:                        _GotoState68Action,
	{_State248, InitializableTypeType}:                   _GotoState83Action,
	{_State248, SliceTypeType}:                           _GotoState100Action,
	{_State248, ArrayTypeType}:                           _GotoState59Action,
	{_State248, MapTypeType}:                             _GotoState88Action,
	{_State248, ExplicitStructDefType}:                   _GotoState74Action,
	{_State248, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State249, IntegerLiteralToken}:                     _GotoState40Action,
	{_State249, FloatLiteralToken}:                       _GotoState35Action,
	{_State249, RuneLiteralToken}:                        _GotoState48Action,
	{_State249, StringLiteralToken}:                      _GotoState49Action,
	{_State249, IdentifierToken}:                         _GotoState38Action,
	{_State249, TrueToken}:                               _GotoState52Action,
	{_State249, FalseToken}:                              _GotoState34Action,
	{_State249, CaseToken}:                               _GotoState29Action,
	{_State249, DefaultToken}:                            _GotoState31Action,
	{_State249, ReturnToken}:                             _GotoState47Action,
	{_State249, BreakToken}:                              _GotoState28Action,
	{_State249, ContinueToken}:                           _GotoState30Action,
	{_State249, FallthroughToken}:                        _GotoState33Action,
	{_State249, ImportToken}:                             _GotoState39Action,
	{_State249, UnsafeToken}:                             _GotoState53Action,
	{_State249, StructToken}:                             _GotoState50Action,
	{_State249, FuncToken}:                               _GotoState36Action,
	{_State249, AsyncToken}:                              _GotoState25Action,
	{_State249, DeferToken}:                              _GotoState32Action,
	{_State249, VarToken}:                                _GotoState15Action,
	{_State249, LetToken}:                                _GotoState12Action,
	{_State249, NotToken}:                                _GotoState45Action,
	{_State249, LabelDeclToken}:                          _GotoState41Action,
	{_State249, LparenToken}:                             _GotoState43Action,
	{_State249, LbracketToken}:                           _GotoState42Action,
	{_State249, SubToken}:                                _GotoState51Action,
	{_State249, MulToken}:                                _GotoState44Action,
	{_State249, BitNegToken}:                             _GotoState27Action,
	{_State249, BitAndToken}:                             _GotoState26Action,
	{_State249, GreaterToken}:                            _GotoState37Action,
	{_State249, ParseErrorToken}:                         _GotoState46Action,
	{_State249, SimpleStatementType}:                     _GotoState99Action,
	{_State249, StatementType}:                           _GotoState349Action,
	{_State249, ExpressionOrImproperStructStatementType}: _GotoState76Action,
	{_State249, ExpressionsType}:                         _GotoState77Action,
	{_State249, CallbackOpType}:                          _GotoState71Action,
	{_State249, CallbackOpStatementType}:                 _GotoState72Action,
	{_State249, UnsafeStatementType}:                     _GotoState102Action,
	{_State249, JumpStatementType}:                       _GotoState85Action,
	{_State249, JumpTypeType}:                            _GotoState86Action,
	{_State249, FallthroughStatementType}:                _GotoState78Action,
	{_State249, AssignStatementType}:                     _GotoState61Action,
	{_State249, UnaryOpAssignStatementType}:              _GotoState101Action,
	{_State249, BinaryOpAssignStatementType}:             _GotoState67Action,
	{_State249, ImportStatementType}:                     _GotoState81Action,
	{_State249, VarDeclPatternType}:                      _GotoState103Action,
	{_State249, VarOrLetType}:                            _GotoState24Action,
	{_State249, AssignPatternType}:                       _GotoState60Action,
	{_State249, ExpressionType}:                          _GotoState75Action,
	{_State249, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State249, SequenceExprType}:                        _GotoState98Action,
	{_State249, CallExprType}:                            _GotoState70Action,
	{_State249, AtomExprType}:                            _GotoState62Action,
	{_State249, ParseErrorExprType}:                      _GotoState92Action,
	{_State249, LiteralExprType}:                         _GotoState87Action,
	{_State249, IdentifierExprType}:                      _GotoState79Action,
	{_State249, BlockExprType}:                           _GotoState69Action,
	{_State249, InitializeExprType}:                      _GotoState84Action,
	{_State249, ImplicitStructExprType}:                  _GotoState80Action,
	{_State249, AccessibleExprType}:                      _GotoState55Action,
	{_State249, AccessExprType}:                          _GotoState54Action,
	{_State249, IndexExprType}:                           _GotoState82Action,
	{_State249, PostfixableExprType}:                     _GotoState94Action,
	{_State249, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State249, PrefixableExprType}:                      _GotoState97Action,
	{_State249, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State249, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State249, MulExprType}:                             _GotoState89Action,
	{_State249, BinaryMulExprType}:                       _GotoState66Action,
	{_State249, AddExprType}:                             _GotoState56Action,
	{_State249, BinaryAddExprType}:                       _GotoState63Action,
	{_State249, CmpExprType}:                             _GotoState73Action,
	{_State249, BinaryCmpExprType}:                       _GotoState65Action,
	{_State249, AndExprType}:                             _GotoState57Action,
	{_State249, BinaryAndExprType}:                       _GotoState64Action,
	{_State249, OrExprType}:                              _GotoState91Action,
	{_State249, BinaryOrExprType}:                        _GotoState68Action,
	{_State249, InitializableTypeType}:                   _GotoState83Action,
	{_State249, SliceTypeType}:                           _GotoState100Action,
	{_State249, ArrayTypeType}:                           _GotoState59Action,
	{_State249, MapTypeType}:                             _GotoState88Action,
	{_State249, ExplicitStructDefType}:                   _GotoState74Action,
	{_State249, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State251, IdentifierToken}:                         _GotoState112Action,
	{_State251, StructToken}:                             _GotoState50Action,
	{_State251, EnumToken}:                               _GotoState109Action,
	{_State251, TraitToken}:                              _GotoState117Action,
	{_State251, FuncToken}:                               _GotoState111Action,
	{_State251, LparenToken}:                             _GotoState113Action,
	{_State251, LbracketToken}:                           _GotoState42Action,
	{_State251, DotToken}:                                _GotoState108Action,
	{_State251, QuestionToken}:                           _GotoState115Action,
	{_State251, ExclaimToken}:                            _GotoState110Action,
	{_State251, TildeTildeToken}:                         _GotoState116Action,
	{_State251, BitNegToken}:                             _GotoState107Action,
	{_State251, BitAndToken}:                             _GotoState106Action,
	{_State251, ParseErrorToken}:                         _GotoState114Action,
	{_State251, InitializableTypeType}:                   _GotoState123Action,
	{_State251, SliceTypeType}:                           _GotoState100Action,
	{_State251, ArrayTypeType}:                           _GotoState59Action,
	{_State251, MapTypeType}:                             _GotoState88Action,
	{_State251, AtomTypeType}:                            _GotoState118Action,
	{_State251, ParseErrorTypeType}:                      _GotoState124Action,
	{_State251, ReturnableTypeType}:                      _GotoState127Action,
	{_State251, PrefixedTypeType}:                        _GotoState126Action,
	{_State251, PrefixTypeOpType}:                        _GotoState125Action,
	{_State251, ValueTypeType}:                           _GotoState350Action,
	{_State251, TraitOpTypeType}:                         _GotoState129Action,
	{_State251, ImplicitStructDefType}:                   _GotoState122Action,
	{_State251, ExplicitStructDefType}:                   _GotoState74Action,
	{_State251, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State251, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State251, TraitDefType}:                            _GotoState128Action,
	{_State251, FuncTypeType}:                            _GotoState120Action,
	{_State252, IdentifierToken}:                         _GotoState112Action,
	{_State252, StructToken}:                             _GotoState50Action,
	{_State252, EnumToken}:                               _GotoState109Action,
	{_State252, TraitToken}:                              _GotoState117Action,
	{_State252, FuncToken}:                               _GotoState111Action,
	{_State252, LparenToken}:                             _GotoState113Action,
	{_State252, LbracketToken}:                           _GotoState42Action,
	{_State252, DotToken}:                                _GotoState108Action,
	{_State252, QuestionToken}:                           _GotoState115Action,
	{_State252, ExclaimToken}:                            _GotoState110Action,
	{_State252, TildeTildeToken}:                         _GotoState116Action,
	{_State252, BitNegToken}:                             _GotoState107Action,
	{_State252, BitAndToken}:                             _GotoState106Action,
	{_State252, ParseErrorToken}:                         _GotoState114Action,
	{_State252, InitializableTypeType}:                   _GotoState123Action,
	{_State252, SliceTypeType}:                           _GotoState100Action,
	{_State252, ArrayTypeType}:                           _GotoState59Action,
	{_State252, MapTypeType}:                             _GotoState88Action,
	{_State252, AtomTypeType}:                            _GotoState118Action,
	{_State252, ParseErrorTypeType}:                      _GotoState124Action,
	{_State252, ReturnableTypeType}:                      _GotoState127Action,
	{_State252, PrefixedTypeType}:                        _GotoState126Action,
	{_State252, PrefixTypeOpType}:                        _GotoState125Action,
	{_State252, ValueTypeType}:                           _GotoState351Action,
	{_State252, TraitOpTypeType}:                         _GotoState129Action,
	{_State252, ImplicitStructDefType}:                   _GotoState122Action,
	{_State252, ExplicitStructDefType}:                   _GotoState74Action,
	{_State252, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State252, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State252, TraitDefType}:                            _GotoState128Action,
	{_State252, FuncTypeType}:                            _GotoState120Action,
	{_State256, AssignToken}:                             _GotoState352Action,
	{_State258, RparenToken}:                             _GotoState354Action,
	{_State258, CommaToken}:                              _GotoState353Action,
	{_State261, AddToken}:                                _GotoState239Action,
	{_State261, SubToken}:                                _GotoState241Action,
	{_State261, MulToken}:                                _GotoState240Action,
	{_State261, TraitOpType}:                             _GotoState242Action,
	{_State262, LparenToken}:                             _GotoState43Action,
	{_State262, ImplicitStructExprType}:                  _GotoState355Action,
	{_State263, IdentifierToken}:                         _GotoState356Action,
	{_State264, IntegerLiteralToken}:                     _GotoState40Action,
	{_State264, FloatLiteralToken}:                       _GotoState35Action,
	{_State264, RuneLiteralToken}:                        _GotoState48Action,
	{_State264, StringLiteralToken}:                      _GotoState49Action,
	{_State264, IdentifierToken}:                         _GotoState38Action,
	{_State264, TrueToken}:                               _GotoState52Action,
	{_State264, FalseToken}:                              _GotoState34Action,
	{_State264, ReturnToken}:                             _GotoState47Action,
	{_State264, BreakToken}:                              _GotoState28Action,
	{_State264, ContinueToken}:                           _GotoState30Action,
	{_State264, FallthroughToken}:                        _GotoState33Action,
	{_State264, UnsafeToken}:                             _GotoState53Action,
	{_State264, StructToken}:                             _GotoState50Action,
	{_State264, FuncToken}:                               _GotoState36Action,
	{_State264, AsyncToken}:                              _GotoState25Action,
	{_State264, DeferToken}:                              _GotoState32Action,
	{_State264, VarToken}:                                _GotoState15Action,
	{_State264, LetToken}:                                _GotoState12Action,
	{_State264, NotToken}:                                _GotoState45Action,
	{_State264, LabelDeclToken}:                          _GotoState41Action,
	{_State264, LparenToken}:                             _GotoState43Action,
	{_State264, LbracketToken}:                           _GotoState42Action,
	{_State264, SubToken}:                                _GotoState51Action,
	{_State264, MulToken}:                                _GotoState44Action,
	{_State264, BitNegToken}:                             _GotoState27Action,
	{_State264, BitAndToken}:                             _GotoState26Action,
	{_State264, GreaterToken}:                            _GotoState37Action,
	{_State264, ParseErrorToken}:                         _GotoState46Action,
	{_State264, SimpleStatementType}:                     _GotoState267Action,
	{_State264, OptionalSimpleStatementType}:             _GotoState357Action,
	{_State264, ExpressionOrImproperStructStatementType}: _GotoState76Action,
	{_State264, ExpressionsType}:                         _GotoState77Action,
	{_State264, CallbackOpType}:                          _GotoState71Action,
	{_State264, CallbackOpStatementType}:                 _GotoState72Action,
	{_State264, UnsafeStatementType}:                     _GotoState102Action,
	{_State264, JumpStatementType}:                       _GotoState85Action,
	{_State264, JumpTypeType}:                            _GotoState86Action,
	{_State264, FallthroughStatementType}:                _GotoState78Action,
	{_State264, AssignStatementType}:                     _GotoState61Action,
	{_State264, UnaryOpAssignStatementType}:              _GotoState101Action,
	{_State264, BinaryOpAssignStatementType}:             _GotoState67Action,
	{_State264, VarDeclPatternType}:                      _GotoState103Action,
	{_State264, VarOrLetType}:                            _GotoState24Action,
	{_State264, AssignPatternType}:                       _GotoState60Action,
	{_State264, ExpressionType}:                          _GotoState75Action,
	{_State264, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State264, SequenceExprType}:                        _GotoState98Action,
	{_State264, CallExprType}:                            _GotoState70Action,
	{_State264, AtomExprType}:                            _GotoState62Action,
	{_State264, ParseErrorExprType}:                      _GotoState92Action,
	{_State264, LiteralExprType}:                         _GotoState87Action,
	{_State264, IdentifierExprType}:                      _GotoState79Action,
	{_State264, BlockExprType}:                           _GotoState69Action,
	{_State264, InitializeExprType}:                      _GotoState84Action,
	{_State264, ImplicitStructExprType}:                  _GotoState80Action,
	{_State264, AccessibleExprType}:                      _GotoState55Action,
	{_State264, AccessExprType}:                          _GotoState54Action,
	{_State264, IndexExprType}:                           _GotoState82Action,
	{_State264, PostfixableExprType}:                     _GotoState94Action,
	{_State264, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State264, PrefixableExprType}:                      _GotoState97Action,
	{_State264, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State264, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State264, MulExprType}:                             _GotoState89Action,
	{_State264, BinaryMulExprType}:                       _GotoState66Action,
	{_State264, AddExprType}:                             _GotoState56Action,
	{_State264, BinaryAddExprType}:                       _GotoState63Action,
	{_State264, CmpExprType}:                             _GotoState73Action,
	{_State264, BinaryCmpExprType}:                       _GotoState65Action,
	{_State264, AndExprType}:                             _GotoState57Action,
	{_State264, BinaryAndExprType}:                       _GotoState64Action,
	{_State264, OrExprType}:                              _GotoState91Action,
	{_State264, BinaryOrExprType}:                        _GotoState68Action,
	{_State264, InitializableTypeType}:                   _GotoState83Action,
	{_State264, SliceTypeType}:                           _GotoState100Action,
	{_State264, ArrayTypeType}:                           _GotoState59Action,
	{_State264, MapTypeType}:                             _GotoState88Action,
	{_State264, ExplicitStructDefType}:                   _GotoState74Action,
	{_State264, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State265, IntegerLiteralToken}:                     _GotoState40Action,
	{_State265, FloatLiteralToken}:                       _GotoState35Action,
	{_State265, RuneLiteralToken}:                        _GotoState48Action,
	{_State265, StringLiteralToken}:                      _GotoState49Action,
	{_State265, IdentifierToken}:                         _GotoState38Action,
	{_State265, TrueToken}:                               _GotoState52Action,
	{_State265, FalseToken}:                              _GotoState34Action,
	{_State265, StructToken}:                             _GotoState50Action,
	{_State265, FuncToken}:                               _GotoState36Action,
	{_State265, VarToken}:                                _GotoState144Action,
	{_State265, LetToken}:                                _GotoState12Action,
	{_State265, NotToken}:                                _GotoState45Action,
	{_State265, LabelDeclToken}:                          _GotoState41Action,
	{_State265, LparenToken}:                             _GotoState43Action,
	{_State265, LbracketToken}:                           _GotoState42Action,
	{_State265, DotToken}:                                _GotoState143Action,
	{_State265, SubToken}:                                _GotoState51Action,
	{_State265, MulToken}:                                _GotoState44Action,
	{_State265, BitNegToken}:                             _GotoState27Action,
	{_State265, BitAndToken}:                             _GotoState26Action,
	{_State265, GreaterToken}:                            _GotoState37Action,
	{_State265, ParseErrorToken}:                         _GotoState46Action,
	{_State265, VarDeclPatternType}:                      _GotoState103Action,
	{_State265, VarOrLetType}:                            _GotoState24Action,
	{_State265, AssignPatternType}:                       _GotoState145Action,
	{_State265, CasePatternType}:                         _GotoState358Action,
	{_State265, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State265, SequenceExprType}:                        _GotoState149Action,
	{_State265, CallExprType}:                            _GotoState70Action,
	{_State265, AtomExprType}:                            _GotoState62Action,
	{_State265, ParseErrorExprType}:                      _GotoState92Action,
	{_State265, LiteralExprType}:                         _GotoState87Action,
	{_State265, IdentifierExprType}:                      _GotoState79Action,
	{_State265, BlockExprType}:                           _GotoState69Action,
	{_State265, InitializeExprType}:                      _GotoState84Action,
	{_State265, ImplicitStructExprType}:                  _GotoState80Action,
	{_State265, AccessibleExprType}:                      _GotoState104Action,
	{_State265, AccessExprType}:                          _GotoState54Action,
	{_State265, IndexExprType}:                           _GotoState82Action,
	{_State265, PostfixableExprType}:                     _GotoState94Action,
	{_State265, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State265, PrefixableExprType}:                      _GotoState97Action,
	{_State265, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State265, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State265, MulExprType}:                             _GotoState89Action,
	{_State265, BinaryMulExprType}:                       _GotoState66Action,
	{_State265, AddExprType}:                             _GotoState56Action,
	{_State265, BinaryAddExprType}:                       _GotoState63Action,
	{_State265, CmpExprType}:                             _GotoState73Action,
	{_State265, BinaryCmpExprType}:                       _GotoState65Action,
	{_State265, AndExprType}:                             _GotoState57Action,
	{_State265, BinaryAndExprType}:                       _GotoState64Action,
	{_State265, OrExprType}:                              _GotoState91Action,
	{_State265, BinaryOrExprType}:                        _GotoState68Action,
	{_State265, InitializableTypeType}:                   _GotoState83Action,
	{_State265, SliceTypeType}:                           _GotoState100Action,
	{_State265, ArrayTypeType}:                           _GotoState59Action,
	{_State265, MapTypeType}:                             _GotoState88Action,
	{_State265, ExplicitStructDefType}:                   _GotoState74Action,
	{_State265, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State269, RparenToken}:                             _GotoState359Action,
	{_State270, CommaToken}:                              _GotoState360Action,
	{_State272, RparenToken}:                             _GotoState361Action,
	{_State273, NewlinesToken}:                           _GotoState363Action,
	{_State273, CommaToken}:                              _GotoState362Action,
	{_State274, IdentifierToken}:                         _GotoState364Action,
	{_State275, IdentifierToken}:                         _GotoState112Action,
	{_State275, StructToken}:                             _GotoState50Action,
	{_State275, EnumToken}:                               _GotoState109Action,
	{_State275, TraitToken}:                              _GotoState117Action,
	{_State275, FuncToken}:                               _GotoState111Action,
	{_State275, LparenToken}:                             _GotoState113Action,
	{_State275, LbracketToken}:                           _GotoState42Action,
	{_State275, DotToken}:                                _GotoState108Action,
	{_State275, QuestionToken}:                           _GotoState115Action,
	{_State275, ExclaimToken}:                            _GotoState110Action,
	{_State275, TildeTildeToken}:                         _GotoState116Action,
	{_State275, BitNegToken}:                             _GotoState107Action,
	{_State275, BitAndToken}:                             _GotoState106Action,
	{_State275, ParseErrorToken}:                         _GotoState114Action,
	{_State275, InitializableTypeType}:                   _GotoState123Action,
	{_State275, SliceTypeType}:                           _GotoState100Action,
	{_State275, ArrayTypeType}:                           _GotoState59Action,
	{_State275, MapTypeType}:                             _GotoState88Action,
	{_State275, AtomTypeType}:                            _GotoState118Action,
	{_State275, ParseErrorTypeType}:                      _GotoState124Action,
	{_State275, ReturnableTypeType}:                      _GotoState127Action,
	{_State275, PrefixedTypeType}:                        _GotoState126Action,
	{_State275, PrefixTypeOpType}:                        _GotoState125Action,
	{_State275, ValueTypeType}:                           _GotoState365Action,
	{_State275, TraitOpTypeType}:                         _GotoState129Action,
	{_State275, ImplicitStructDefType}:                   _GotoState122Action,
	{_State275, ExplicitStructDefType}:                   _GotoState74Action,
	{_State275, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State275, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State275, TraitDefType}:                            _GotoState128Action,
	{_State275, FuncTypeType}:                            _GotoState120Action,
	{_State276, IntegerLiteralToken}:                     _GotoState366Action,
	{_State279, IntegerLiteralToken}:                     _GotoState40Action,
	{_State279, FloatLiteralToken}:                       _GotoState35Action,
	{_State279, RuneLiteralToken}:                        _GotoState48Action,
	{_State279, StringLiteralToken}:                      _GotoState49Action,
	{_State279, IdentifierToken}:                         _GotoState38Action,
	{_State279, TrueToken}:                               _GotoState52Action,
	{_State279, FalseToken}:                              _GotoState34Action,
	{_State279, StructToken}:                             _GotoState50Action,
	{_State279, FuncToken}:                               _GotoState36Action,
	{_State279, VarToken}:                                _GotoState15Action,
	{_State279, LetToken}:                                _GotoState12Action,
	{_State279, NotToken}:                                _GotoState45Action,
	{_State279, LabelDeclToken}:                          _GotoState41Action,
	{_State279, LparenToken}:                             _GotoState43Action,
	{_State279, LbracketToken}:                           _GotoState42Action,
	{_State279, SubToken}:                                _GotoState51Action,
	{_State279, MulToken}:                                _GotoState44Action,
	{_State279, BitNegToken}:                             _GotoState27Action,
	{_State279, BitAndToken}:                             _GotoState26Action,
	{_State279, GreaterToken}:                            _GotoState37Action,
	{_State279, ParseErrorToken}:                         _GotoState46Action,
	{_State279, VarDeclPatternType}:                      _GotoState103Action,
	{_State279, VarOrLetType}:                            _GotoState24Action,
	{_State279, ExpressionType}:                          _GotoState367Action,
	{_State279, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State279, SequenceExprType}:                        _GotoState105Action,
	{_State279, CallExprType}:                            _GotoState70Action,
	{_State279, AtomExprType}:                            _GotoState62Action,
	{_State279, ParseErrorExprType}:                      _GotoState92Action,
	{_State279, LiteralExprType}:                         _GotoState87Action,
	{_State279, IdentifierExprType}:                      _GotoState79Action,
	{_State279, BlockExprType}:                           _GotoState69Action,
	{_State279, InitializeExprType}:                      _GotoState84Action,
	{_State279, ImplicitStructExprType}:                  _GotoState80Action,
	{_State279, AccessibleExprType}:                      _GotoState104Action,
	{_State279, AccessExprType}:                          _GotoState54Action,
	{_State279, IndexExprType}:                           _GotoState82Action,
	{_State279, PostfixableExprType}:                     _GotoState94Action,
	{_State279, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State279, PrefixableExprType}:                      _GotoState97Action,
	{_State279, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State279, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State279, MulExprType}:                             _GotoState89Action,
	{_State279, BinaryMulExprType}:                       _GotoState66Action,
	{_State279, AddExprType}:                             _GotoState56Action,
	{_State279, BinaryAddExprType}:                       _GotoState63Action,
	{_State279, CmpExprType}:                             _GotoState73Action,
	{_State279, BinaryCmpExprType}:                       _GotoState65Action,
	{_State279, AndExprType}:                             _GotoState57Action,
	{_State279, BinaryAndExprType}:                       _GotoState64Action,
	{_State279, OrExprType}:                              _GotoState91Action,
	{_State279, BinaryOrExprType}:                        _GotoState68Action,
	{_State279, InitializableTypeType}:                   _GotoState83Action,
	{_State279, SliceTypeType}:                           _GotoState100Action,
	{_State279, ArrayTypeType}:                           _GotoState59Action,
	{_State279, MapTypeType}:                             _GotoState88Action,
	{_State279, ExplicitStructDefType}:                   _GotoState74Action,
	{_State279, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State281, IntegerLiteralToken}:                     _GotoState40Action,
	{_State281, FloatLiteralToken}:                       _GotoState35Action,
	{_State281, RuneLiteralToken}:                        _GotoState48Action,
	{_State281, StringLiteralToken}:                      _GotoState49Action,
	{_State281, IdentifierToken}:                         _GotoState38Action,
	{_State281, TrueToken}:                               _GotoState52Action,
	{_State281, FalseToken}:                              _GotoState34Action,
	{_State281, StructToken}:                             _GotoState50Action,
	{_State281, FuncToken}:                               _GotoState36Action,
	{_State281, VarToken}:                                _GotoState15Action,
	{_State281, LetToken}:                                _GotoState12Action,
	{_State281, NotToken}:                                _GotoState45Action,
	{_State281, LabelDeclToken}:                          _GotoState41Action,
	{_State281, LparenToken}:                             _GotoState43Action,
	{_State281, LbracketToken}:                           _GotoState42Action,
	{_State281, SubToken}:                                _GotoState51Action,
	{_State281, MulToken}:                                _GotoState44Action,
	{_State281, BitNegToken}:                             _GotoState27Action,
	{_State281, BitAndToken}:                             _GotoState26Action,
	{_State281, GreaterToken}:                            _GotoState37Action,
	{_State281, ParseErrorToken}:                         _GotoState46Action,
	{_State281, VarDeclPatternType}:                      _GotoState103Action,
	{_State281, VarOrLetType}:                            _GotoState24Action,
	{_State281, ExpressionType}:                          _GotoState368Action,
	{_State281, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State281, SequenceExprType}:                        _GotoState105Action,
	{_State281, CallExprType}:                            _GotoState70Action,
	{_State281, AtomExprType}:                            _GotoState62Action,
	{_State281, ParseErrorExprType}:                      _GotoState92Action,
	{_State281, LiteralExprType}:                         _GotoState87Action,
	{_State281, IdentifierExprType}:                      _GotoState79Action,
	{_State281, BlockExprType}:                           _GotoState69Action,
	{_State281, InitializeExprType}:                      _GotoState84Action,
	{_State281, ImplicitStructExprType}:                  _GotoState80Action,
	{_State281, AccessibleExprType}:                      _GotoState104Action,
	{_State281, AccessExprType}:                          _GotoState54Action,
	{_State281, IndexExprType}:                           _GotoState82Action,
	{_State281, PostfixableExprType}:                     _GotoState94Action,
	{_State281, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State281, PrefixableExprType}:                      _GotoState97Action,
	{_State281, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State281, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State281, MulExprType}:                             _GotoState89Action,
	{_State281, BinaryMulExprType}:                       _GotoState66Action,
	{_State281, AddExprType}:                             _GotoState56Action,
	{_State281, BinaryAddExprType}:                       _GotoState63Action,
	{_State281, CmpExprType}:                             _GotoState73Action,
	{_State281, BinaryCmpExprType}:                       _GotoState65Action,
	{_State281, AndExprType}:                             _GotoState57Action,
	{_State281, BinaryAndExprType}:                       _GotoState64Action,
	{_State281, OrExprType}:                              _GotoState91Action,
	{_State281, BinaryOrExprType}:                        _GotoState68Action,
	{_State281, InitializableTypeType}:                   _GotoState83Action,
	{_State281, SliceTypeType}:                           _GotoState100Action,
	{_State281, ArrayTypeType}:                           _GotoState59Action,
	{_State281, MapTypeType}:                             _GotoState88Action,
	{_State281, ExplicitStructDefType}:                   _GotoState74Action,
	{_State281, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State282, IntegerLiteralToken}:                     _GotoState40Action,
	{_State282, FloatLiteralToken}:                       _GotoState35Action,
	{_State282, RuneLiteralToken}:                        _GotoState48Action,
	{_State282, StringLiteralToken}:                      _GotoState49Action,
	{_State282, IdentifierToken}:                         _GotoState38Action,
	{_State282, TrueToken}:                               _GotoState52Action,
	{_State282, FalseToken}:                              _GotoState34Action,
	{_State282, StructToken}:                             _GotoState50Action,
	{_State282, FuncToken}:                               _GotoState36Action,
	{_State282, VarToken}:                                _GotoState15Action,
	{_State282, LetToken}:                                _GotoState12Action,
	{_State282, NotToken}:                                _GotoState45Action,
	{_State282, LabelDeclToken}:                          _GotoState41Action,
	{_State282, LparenToken}:                             _GotoState43Action,
	{_State282, LbracketToken}:                           _GotoState42Action,
	{_State282, SubToken}:                                _GotoState51Action,
	{_State282, MulToken}:                                _GotoState44Action,
	{_State282, BitNegToken}:                             _GotoState27Action,
	{_State282, BitAndToken}:                             _GotoState26Action,
	{_State282, GreaterToken}:                            _GotoState37Action,
	{_State282, ParseErrorToken}:                         _GotoState46Action,
	{_State282, VarDeclPatternType}:                      _GotoState103Action,
	{_State282, VarOrLetType}:                            _GotoState24Action,
	{_State282, ExpressionType}:                          _GotoState369Action,
	{_State282, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State282, SequenceExprType}:                        _GotoState105Action,
	{_State282, CallExprType}:                            _GotoState70Action,
	{_State282, AtomExprType}:                            _GotoState62Action,
	{_State282, ParseErrorExprType}:                      _GotoState92Action,
	{_State282, LiteralExprType}:                         _GotoState87Action,
	{_State282, IdentifierExprType}:                      _GotoState79Action,
	{_State282, BlockExprType}:                           _GotoState69Action,
	{_State282, InitializeExprType}:                      _GotoState84Action,
	{_State282, ImplicitStructExprType}:                  _GotoState80Action,
	{_State282, AccessibleExprType}:                      _GotoState104Action,
	{_State282, AccessExprType}:                          _GotoState54Action,
	{_State282, IndexExprType}:                           _GotoState82Action,
	{_State282, PostfixableExprType}:                     _GotoState94Action,
	{_State282, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State282, PrefixableExprType}:                      _GotoState97Action,
	{_State282, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State282, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State282, MulExprType}:                             _GotoState89Action,
	{_State282, BinaryMulExprType}:                       _GotoState66Action,
	{_State282, AddExprType}:                             _GotoState56Action,
	{_State282, BinaryAddExprType}:                       _GotoState63Action,
	{_State282, CmpExprType}:                             _GotoState73Action,
	{_State282, BinaryCmpExprType}:                       _GotoState65Action,
	{_State282, AndExprType}:                             _GotoState57Action,
	{_State282, BinaryAndExprType}:                       _GotoState64Action,
	{_State282, OrExprType}:                              _GotoState91Action,
	{_State282, BinaryOrExprType}:                        _GotoState68Action,
	{_State282, InitializableTypeType}:                   _GotoState83Action,
	{_State282, SliceTypeType}:                           _GotoState100Action,
	{_State282, ArrayTypeType}:                           _GotoState59Action,
	{_State282, MapTypeType}:                             _GotoState88Action,
	{_State282, ExplicitStructDefType}:                   _GotoState74Action,
	{_State282, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State284, IntegerLiteralToken}:                     _GotoState40Action,
	{_State284, FloatLiteralToken}:                       _GotoState35Action,
	{_State284, RuneLiteralToken}:                        _GotoState48Action,
	{_State284, StringLiteralToken}:                      _GotoState49Action,
	{_State284, IdentifierToken}:                         _GotoState159Action,
	{_State284, TrueToken}:                               _GotoState52Action,
	{_State284, FalseToken}:                              _GotoState34Action,
	{_State284, StructToken}:                             _GotoState50Action,
	{_State284, FuncToken}:                               _GotoState36Action,
	{_State284, VarToken}:                                _GotoState15Action,
	{_State284, LetToken}:                                _GotoState12Action,
	{_State284, NotToken}:                                _GotoState45Action,
	{_State284, LabelDeclToken}:                          _GotoState41Action,
	{_State284, LparenToken}:                             _GotoState43Action,
	{_State284, LbracketToken}:                           _GotoState42Action,
	{_State284, ColonToken}:                              _GotoState157Action,
	{_State284, EllipsisToken}:                           _GotoState158Action,
	{_State284, SubToken}:                                _GotoState51Action,
	{_State284, MulToken}:                                _GotoState44Action,
	{_State284, BitNegToken}:                             _GotoState27Action,
	{_State284, BitAndToken}:                             _GotoState26Action,
	{_State284, GreaterToken}:                            _GotoState37Action,
	{_State284, ParseErrorToken}:                         _GotoState46Action,
	{_State284, VarDeclPatternType}:                      _GotoState103Action,
	{_State284, VarOrLetType}:                            _GotoState24Action,
	{_State284, ExpressionType}:                          _GotoState163Action,
	{_State284, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State284, SequenceExprType}:                        _GotoState105Action,
	{_State284, CallExprType}:                            _GotoState70Action,
	{_State284, ArgumentType}:                            _GotoState370Action,
	{_State284, ColonExprType}:                           _GotoState162Action,
	{_State284, AtomExprType}:                            _GotoState62Action,
	{_State284, ParseErrorExprType}:                      _GotoState92Action,
	{_State284, LiteralExprType}:                         _GotoState87Action,
	{_State284, IdentifierExprType}:                      _GotoState79Action,
	{_State284, BlockExprType}:                           _GotoState69Action,
	{_State284, InitializeExprType}:                      _GotoState84Action,
	{_State284, ImplicitStructExprType}:                  _GotoState80Action,
	{_State284, AccessibleExprType}:                      _GotoState104Action,
	{_State284, AccessExprType}:                          _GotoState54Action,
	{_State284, IndexExprType}:                           _GotoState82Action,
	{_State284, PostfixableExprType}:                     _GotoState94Action,
	{_State284, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State284, PrefixableExprType}:                      _GotoState97Action,
	{_State284, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State284, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State284, MulExprType}:                             _GotoState89Action,
	{_State284, BinaryMulExprType}:                       _GotoState66Action,
	{_State284, AddExprType}:                             _GotoState56Action,
	{_State284, BinaryAddExprType}:                       _GotoState63Action,
	{_State284, CmpExprType}:                             _GotoState73Action,
	{_State284, BinaryCmpExprType}:                       _GotoState65Action,
	{_State284, AndExprType}:                             _GotoState57Action,
	{_State284, BinaryAndExprType}:                       _GotoState64Action,
	{_State284, OrExprType}:                              _GotoState91Action,
	{_State284, BinaryOrExprType}:                        _GotoState68Action,
	{_State284, InitializableTypeType}:                   _GotoState83Action,
	{_State284, SliceTypeType}:                           _GotoState100Action,
	{_State284, ArrayTypeType}:                           _GotoState59Action,
	{_State284, MapTypeType}:                             _GotoState88Action,
	{_State284, ExplicitStructDefType}:                   _GotoState74Action,
	{_State284, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State285, RparenToken}:                             _GotoState371Action,
	{_State287, NewlinesToken}:                           _GotoState373Action,
	{_State287, CommaToken}:                              _GotoState372Action,
	{_State288, GreaterToken}:                            _GotoState374Action,
	{_State289, RbracketToken}:                           _GotoState375Action,
	{_State290, CommaToken}:                              _GotoState376Action,
	{_State291, AddToken}:                                _GotoState239Action,
	{_State291, SubToken}:                                _GotoState241Action,
	{_State291, MulToken}:                                _GotoState240Action,
	{_State291, TraitOpType}:                             _GotoState242Action,
	{_State293, RbracketToken}:                           _GotoState377Action,
	{_State295, IntegerLiteralToken}:                     _GotoState40Action,
	{_State295, FloatLiteralToken}:                       _GotoState35Action,
	{_State295, RuneLiteralToken}:                        _GotoState48Action,
	{_State295, StringLiteralToken}:                      _GotoState49Action,
	{_State295, IdentifierToken}:                         _GotoState159Action,
	{_State295, TrueToken}:                               _GotoState52Action,
	{_State295, FalseToken}:                              _GotoState34Action,
	{_State295, StructToken}:                             _GotoState50Action,
	{_State295, FuncToken}:                               _GotoState36Action,
	{_State295, VarToken}:                                _GotoState15Action,
	{_State295, LetToken}:                                _GotoState12Action,
	{_State295, NotToken}:                                _GotoState45Action,
	{_State295, LabelDeclToken}:                          _GotoState41Action,
	{_State295, LparenToken}:                             _GotoState43Action,
	{_State295, LbracketToken}:                           _GotoState42Action,
	{_State295, ColonToken}:                              _GotoState157Action,
	{_State295, EllipsisToken}:                           _GotoState158Action,
	{_State295, SubToken}:                                _GotoState51Action,
	{_State295, MulToken}:                                _GotoState44Action,
	{_State295, BitNegToken}:                             _GotoState27Action,
	{_State295, BitAndToken}:                             _GotoState26Action,
	{_State295, GreaterToken}:                            _GotoState37Action,
	{_State295, ParseErrorToken}:                         _GotoState46Action,
	{_State295, VarDeclPatternType}:                      _GotoState103Action,
	{_State295, VarOrLetType}:                            _GotoState24Action,
	{_State295, ExpressionType}:                          _GotoState163Action,
	{_State295, OptionalLabelDeclType}:                   _GotoState90Action,
	{_State295, SequenceExprType}:                        _GotoState105Action,
	{_State295, CallExprType}:                            _GotoState70Action,
	{_State295, ProperArgumentsType}:                     _GotoState164Action,
	{_State295, ArgumentsType}:                           _GotoState378Action,
	{_State295, ArgumentType}:                            _GotoState160Action,
	{_State295, ColonExprType}:                           _GotoState162Action,
	{_State295, AtomExprType}:                            _GotoState62Action,
	{_State295, ParseErrorExprType}:                      _GotoState92Action,
	{_State295, LiteralExprType}:                         _GotoState87Action,
	{_State295, IdentifierExprType}:                      _GotoState79Action,
	{_State295, BlockExprType}:                           _GotoState69Action,
	{_State295, InitializeExprType}:                      _GotoState84Action,
	{_State295, ImplicitStructExprType}:                  _GotoState80Action,
	{_State295, AccessibleExprType}:                      _GotoState104Action,
	{_State295, AccessExprType}:                          _GotoState54Action,
	{_State295, IndexExprType}:                           _GotoState82Action,
	{_State295, PostfixableExprType}:                     _GotoState94Action,
	{_State295, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State295, PrefixableExprType}:                      _GotoState97Action,
	{_State295, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State295, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State295, MulExprType}:                             _GotoState89Action,
	{_State295, BinaryMulExprType}:                       _GotoState66Action,
	{_State295, AddExprType}:                             _GotoState56Action,
	{_State295, BinaryAddExprType}:                       _GotoState63Action,
	{_State295, CmpExprType}:                             _GotoState73Action,
	{_State295, BinaryCmpExprType}:                       _GotoState65Action,
	{_State295, AndExprType}:                             _GotoState57Action,
	{_State295, BinaryAndExprType}:                       _GotoState64Action,
	{_State295, OrExprType}:                              _GotoState91Action,
	{_State295, BinaryOrExprType}:                        _GotoState68Action,
	{_State295, InitializableTypeType}:                   _GotoState83Action,
	{_State295, SliceTypeType}:                           _GotoState100Action,
	{_State295, ArrayTypeType}:                           _GotoState59Action,
	{_State295, MapTypeType}:                             _GotoState88Action,
	{_State295, ExplicitStructDefType}:                   _GotoState74Action,
	{_State295, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State296, MulToken}:                                _GotoState212Action,
	{_State296, DivToken}:                                _GotoState210Action,
	{_State296, ModToken}:                                _GotoState211Action,
	{_State296, BitAndToken}:                             _GotoState207Action,
	{_State296, BitLshiftToken}:                          _GotoState208Action,
	{_State296, BitRshiftToken}:                          _GotoState209Action,
	{_State296, MulOpType}:                               _GotoState213Action,
	{_State297, EqualToken}:                              _GotoState196Action,
	{_State297, NotEqualToken}:                           _GotoState201Action,
	{_State297, LessToken}:                               _GotoState199Action,
	{_State297, LessOrEqualToken}:                        _GotoState200Action,
	{_State297, GreaterToken}:                            _GotoState197Action,
	{_State297, GreaterOrEqualToken}:                     _GotoState198Action,
	{_State297, CmpOpType}:                               _GotoState202Action,
	{_State299, AddToken}:                                _GotoState187Action,
	{_State299, SubToken}:                                _GotoState190Action,
	{_State299, BitXorToken}:                             _GotoState189Action,
	{_State299, BitOrToken}:                              _GotoState188Action,
	{_State299, AddOpType}:                               _GotoState191Action,
	{_State301, RparenToken}:                             _GotoState379Action,
	{_State302, CommaToken}:                              _GotoState203Action,
	{_State304, ForToken}:                                _GotoState380Action,
	{_State305, InToken}:                                 _GotoState382Action,
	{_State305, AssignToken}:                             _GotoState381Action,
	{_State306, SemicolonToken}:                          _GotoState383Action,
	{_State307, DoToken}:                                 _GotoState384Action,
	{_State308, IntegerLiteralToken}:                     _GotoState40Action,
	{_State308, FloatLiteralToken}:                       _GotoState35Action,
	{_State308, RuneLiteralToken}:                        _GotoState48Action,
	{_State308, StringLiteralToken}:                      _GotoState49Action,
	{_State308, IdentifierToken}:                         _GotoState38Action,
	{_State308, TrueToken}:                               _GotoState52Action,
	{_State308, FalseToken}:                              _GotoState34Action,
	{_State308, StructToken}:                             _GotoState50Action,
	{_State308, FuncToken}:                               _GotoState36Action,
	{_State308, VarToken}:                                _GotoState144Action,
	{_State308, LetToken}:                                _GotoState12Action,
	{_State308, NotToken}:                                _GotoState45Action,
	{_State308, LabelDeclToken}:                          _GotoState41Action,
	{_State308, LparenToken}:                             _GotoState43Action,
	{_State308, LbracketToken}:                           _GotoState42Action,
	{_State308, DotToken}:                                _GotoState143Action,
	{_State308, SubToken}:                                _GotoState51Action,
	{_State308, MulToken}:                                _GotoState44Action,
	{_State308, BitNegToken}:                             _GotoState27Action,
	{_State308, BitAndToken}:                             _GotoState26Action,
	{_State308, GreaterToken}:                            _GotoState37Action,
	{_State308, ParseErrorToken}:                         _GotoState46Action,
	{_State308, CasePatternsType}:                        _GotoState385Action,
	{_State308, VarDeclPatternType}:                      _GotoState103Action,
	{_State308, VarOrLetType}:                            _GotoState24Action,
	{_State308, AssignPatternType}:                       _GotoState145Action,
	{_State308, CasePatternType}:                         _GotoState146Action,
	{_State308, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State308, SequenceExprType}:                        _GotoState149Action,
	{_State308, CallExprType}:                            _GotoState70Action,
	{_State308, AtomExprType}:                            _GotoState62Action,
	{_State308, ParseErrorExprType}:                      _GotoState92Action,
	{_State308, LiteralExprType}:                         _GotoState87Action,
	{_State308, IdentifierExprType}:                      _GotoState79Action,
	{_State308, BlockExprType}:                           _GotoState69Action,
	{_State308, InitializeExprType}:                      _GotoState84Action,
	{_State308, ImplicitStructExprType}:                  _GotoState80Action,
	{_State308, AccessibleExprType}:                      _GotoState104Action,
	{_State308, AccessExprType}:                          _GotoState54Action,
	{_State308, IndexExprType}:                           _GotoState82Action,
	{_State308, PostfixableExprType}:                     _GotoState94Action,
	{_State308, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State308, PrefixableExprType}:                      _GotoState97Action,
	{_State308, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State308, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State308, MulExprType}:                             _GotoState89Action,
	{_State308, BinaryMulExprType}:                       _GotoState66Action,
	{_State308, AddExprType}:                             _GotoState56Action,
	{_State308, BinaryAddExprType}:                       _GotoState63Action,
	{_State308, CmpExprType}:                             _GotoState73Action,
	{_State308, BinaryCmpExprType}:                       _GotoState65Action,
	{_State308, AndExprType}:                             _GotoState57Action,
	{_State308, BinaryAndExprType}:                       _GotoState64Action,
	{_State308, OrExprType}:                              _GotoState91Action,
	{_State308, BinaryOrExprType}:                        _GotoState68Action,
	{_State308, InitializableTypeType}:                   _GotoState83Action,
	{_State308, SliceTypeType}:                           _GotoState100Action,
	{_State308, ArrayTypeType}:                           _GotoState59Action,
	{_State308, MapTypeType}:                             _GotoState88Action,
	{_State308, ExplicitStructDefType}:                   _GotoState74Action,
	{_State308, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State309, LbraceToken}:                             _GotoState11Action,
	{_State309, StatementBlockType}:                      _GotoState386Action,
	{_State311, LbraceToken}:                             _GotoState11Action,
	{_State311, StatementBlockType}:                      _GotoState387Action,
	{_State312, AndToken}:                                _GotoState192Action,
	{_State313, NewlinesToken}:                           _GotoState388Action,
	{_State313, OrToken}:                                 _GotoState389Action,
	{_State314, RparenToken}:                             _GotoState390Action,
	{_State315, AssignToken}:                             _GotoState327Action,
	{_State316, NewlinesToken}:                           _GotoState391Action,
	{_State316, OrToken}:                                 _GotoState392Action,
	{_State317, IdentifierToken}:                         _GotoState112Action,
	{_State317, StructToken}:                             _GotoState50Action,
	{_State317, EnumToken}:                               _GotoState109Action,
	{_State317, TraitToken}:                              _GotoState117Action,
	{_State317, FuncToken}:                               _GotoState111Action,
	{_State317, LparenToken}:                             _GotoState113Action,
	{_State317, LbracketToken}:                           _GotoState42Action,
	{_State317, DotToken}:                                _GotoState108Action,
	{_State317, QuestionToken}:                           _GotoState115Action,
	{_State317, ExclaimToken}:                            _GotoState110Action,
	{_State317, TildeTildeToken}:                         _GotoState116Action,
	{_State317, BitNegToken}:                             _GotoState107Action,
	{_State317, BitAndToken}:                             _GotoState106Action,
	{_State317, ParseErrorToken}:                         _GotoState114Action,
	{_State317, InitializableTypeType}:                   _GotoState123Action,
	{_State317, SliceTypeType}:                           _GotoState100Action,
	{_State317, ArrayTypeType}:                           _GotoState59Action,
	{_State317, MapTypeType}:                             _GotoState88Action,
	{_State317, AtomTypeType}:                            _GotoState118Action,
	{_State317, ParseErrorTypeType}:                      _GotoState124Action,
	{_State317, ReturnableTypeType}:                      _GotoState127Action,
	{_State317, PrefixedTypeType}:                        _GotoState126Action,
	{_State317, PrefixTypeOpType}:                        _GotoState125Action,
	{_State317, ValueTypeType}:                           _GotoState393Action,
	{_State317, TraitOpTypeType}:                         _GotoState129Action,
	{_State317, ImplicitStructDefType}:                   _GotoState122Action,
	{_State317, ExplicitStructDefType}:                   _GotoState74Action,
	{_State317, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State317, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State317, TraitDefType}:                            _GotoState128Action,
	{_State317, FuncTypeType}:                            _GotoState120Action,
	{_State318, IdentifierToken}:                         _GotoState112Action,
	{_State318, StructToken}:                             _GotoState50Action,
	{_State318, EnumToken}:                               _GotoState109Action,
	{_State318, TraitToken}:                              _GotoState117Action,
	{_State318, FuncToken}:                               _GotoState111Action,
	{_State318, LparenToken}:                             _GotoState113Action,
	{_State318, LbracketToken}:                           _GotoState42Action,
	{_State318, DotToken}:                                _GotoState324Action,
	{_State318, QuestionToken}:                           _GotoState115Action,
	{_State318, ExclaimToken}:                            _GotoState110Action,
	{_State318, DollarLbracketToken}:                     _GotoState176Action,
	{_State318, EllipsisToken}:                           _GotoState394Action,
	{_State318, TildeTildeToken}:                         _GotoState116Action,
	{_State318, BitNegToken}:                             _GotoState107Action,
	{_State318, BitAndToken}:                             _GotoState106Action,
	{_State318, ParseErrorToken}:                         _GotoState114Action,
	{_State318, OptionalGenericBindingType}:              _GotoState228Action,
	{_State318, InitializableTypeType}:                   _GotoState123Action,
	{_State318, SliceTypeType}:                           _GotoState100Action,
	{_State318, ArrayTypeType}:                           _GotoState59Action,
	{_State318, MapTypeType}:                             _GotoState88Action,
	{_State318, AtomTypeType}:                            _GotoState118Action,
	{_State318, ParseErrorTypeType}:                      _GotoState124Action,
	{_State318, ReturnableTypeType}:                      _GotoState127Action,
	{_State318, PrefixedTypeType}:                        _GotoState126Action,
	{_State318, PrefixTypeOpType}:                        _GotoState125Action,
	{_State318, ValueTypeType}:                           _GotoState395Action,
	{_State318, TraitOpTypeType}:                         _GotoState129Action,
	{_State318, ImplicitStructDefType}:                   _GotoState122Action,
	{_State318, ExplicitStructDefType}:                   _GotoState74Action,
	{_State318, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State318, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State318, TraitDefType}:                            _GotoState128Action,
	{_State318, FuncTypeType}:                            _GotoState120Action,
	{_State320, RparenToken}:                             _GotoState396Action,
	{_State321, CommaToken}:                              _GotoState397Action,
	{_State322, AddToken}:                                _GotoState239Action,
	{_State322, SubToken}:                                _GotoState241Action,
	{_State322, MulToken}:                                _GotoState240Action,
	{_State322, TraitOpType}:                             _GotoState242Action,
	{_State323, DollarLbracketToken}:                     _GotoState176Action,
	{_State323, OptionalGenericBindingType}:              _GotoState398Action,
	{_State324, IdentifierToken}:                         _GotoState323Action,
	{_State324, DollarLbracketToken}:                     _GotoState176Action,
	{_State324, OptionalGenericBindingType}:              _GotoState224Action,
	{_State325, AddToken}:                                _GotoState239Action,
	{_State325, SubToken}:                                _GotoState241Action,
	{_State325, MulToken}:                                _GotoState240Action,
	{_State325, TraitOpType}:                             _GotoState242Action,
	{_State326, IdentifierToken}:                         _GotoState229Action,
	{_State326, UnsafeToken}:                             _GotoState53Action,
	{_State326, StructToken}:                             _GotoState50Action,
	{_State326, EnumToken}:                               _GotoState109Action,
	{_State326, TraitToken}:                              _GotoState117Action,
	{_State326, FuncToken}:                               _GotoState111Action,
	{_State326, LparenToken}:                             _GotoState113Action,
	{_State326, LbracketToken}:                           _GotoState42Action,
	{_State326, DotToken}:                                _GotoState108Action,
	{_State326, QuestionToken}:                           _GotoState115Action,
	{_State326, ExclaimToken}:                            _GotoState110Action,
	{_State326, TildeTildeToken}:                         _GotoState116Action,
	{_State326, BitNegToken}:                             _GotoState107Action,
	{_State326, BitAndToken}:                             _GotoState106Action,
	{_State326, ParseErrorToken}:                         _GotoState114Action,
	{_State326, UnsafeStatementType}:                     _GotoState235Action,
	{_State326, InitializableTypeType}:                   _GotoState123Action,
	{_State326, SliceTypeType}:                           _GotoState100Action,
	{_State326, ArrayTypeType}:                           _GotoState59Action,
	{_State326, MapTypeType}:                             _GotoState88Action,
	{_State326, AtomTypeType}:                            _GotoState118Action,
	{_State326, ParseErrorTypeType}:                      _GotoState124Action,
	{_State326, ReturnableTypeType}:                      _GotoState127Action,
	{_State326, PrefixedTypeType}:                        _GotoState126Action,
	{_State326, PrefixTypeOpType}:                        _GotoState125Action,
	{_State326, ValueTypeType}:                           _GotoState236Action,
	{_State326, TraitOpTypeType}:                         _GotoState129Action,
	{_State326, FieldDefType}:                            _GotoState315Action,
	{_State326, ImplicitStructDefType}:                   _GotoState122Action,
	{_State326, ExplicitStructDefType}:                   _GotoState74Action,
	{_State326, EnumValueDefType}:                        _GotoState399Action,
	{_State326, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State326, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State326, TraitDefType}:                            _GotoState128Action,
	{_State326, FuncTypeType}:                            _GotoState120Action,
	{_State327, DefaultToken}:                            _GotoState400Action,
	{_State328, IdentifierToken}:                         _GotoState229Action,
	{_State328, UnsafeToken}:                             _GotoState53Action,
	{_State328, StructToken}:                             _GotoState50Action,
	{_State328, EnumToken}:                               _GotoState109Action,
	{_State328, TraitToken}:                              _GotoState117Action,
	{_State328, FuncToken}:                               _GotoState111Action,
	{_State328, LparenToken}:                             _GotoState113Action,
	{_State328, LbracketToken}:                           _GotoState42Action,
	{_State328, DotToken}:                                _GotoState108Action,
	{_State328, QuestionToken}:                           _GotoState115Action,
	{_State328, ExclaimToken}:                            _GotoState110Action,
	{_State328, TildeTildeToken}:                         _GotoState116Action,
	{_State328, BitNegToken}:                             _GotoState107Action,
	{_State328, BitAndToken}:                             _GotoState106Action,
	{_State328, ParseErrorToken}:                         _GotoState114Action,
	{_State328, UnsafeStatementType}:                     _GotoState235Action,
	{_State328, InitializableTypeType}:                   _GotoState123Action,
	{_State328, SliceTypeType}:                           _GotoState100Action,
	{_State328, ArrayTypeType}:                           _GotoState59Action,
	{_State328, MapTypeType}:                             _GotoState88Action,
	{_State328, AtomTypeType}:                            _GotoState118Action,
	{_State328, ParseErrorTypeType}:                      _GotoState124Action,
	{_State328, ReturnableTypeType}:                      _GotoState127Action,
	{_State328, PrefixedTypeType}:                        _GotoState126Action,
	{_State328, PrefixTypeOpType}:                        _GotoState125Action,
	{_State328, ValueTypeType}:                           _GotoState236Action,
	{_State328, TraitOpTypeType}:                         _GotoState129Action,
	{_State328, FieldDefType}:                            _GotoState315Action,
	{_State328, ImplicitStructDefType}:                   _GotoState122Action,
	{_State328, ExplicitStructDefType}:                   _GotoState74Action,
	{_State328, EnumValueDefType}:                        _GotoState401Action,
	{_State328, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State328, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State328, TraitDefType}:                            _GotoState128Action,
	{_State328, FuncTypeType}:                            _GotoState120Action,
	{_State331, IdentifierToken}:                         _GotoState229Action,
	{_State331, UnsafeToken}:                             _GotoState53Action,
	{_State331, StructToken}:                             _GotoState50Action,
	{_State331, EnumToken}:                               _GotoState109Action,
	{_State331, TraitToken}:                              _GotoState117Action,
	{_State331, FuncToken}:                               _GotoState111Action,
	{_State331, LparenToken}:                             _GotoState113Action,
	{_State331, LbracketToken}:                           _GotoState42Action,
	{_State331, DotToken}:                                _GotoState108Action,
	{_State331, QuestionToken}:                           _GotoState115Action,
	{_State331, ExclaimToken}:                            _GotoState110Action,
	{_State331, TildeTildeToken}:                         _GotoState116Action,
	{_State331, BitNegToken}:                             _GotoState107Action,
	{_State331, BitAndToken}:                             _GotoState106Action,
	{_State331, ParseErrorToken}:                         _GotoState114Action,
	{_State331, UnsafeStatementType}:                     _GotoState235Action,
	{_State331, InitializableTypeType}:                   _GotoState123Action,
	{_State331, SliceTypeType}:                           _GotoState100Action,
	{_State331, ArrayTypeType}:                           _GotoState59Action,
	{_State331, MapTypeType}:                             _GotoState88Action,
	{_State331, AtomTypeType}:                            _GotoState118Action,
	{_State331, ParseErrorTypeType}:                      _GotoState124Action,
	{_State331, ReturnableTypeType}:                      _GotoState127Action,
	{_State331, PrefixedTypeType}:                        _GotoState126Action,
	{_State331, PrefixTypeOpType}:                        _GotoState125Action,
	{_State331, ValueTypeType}:                           _GotoState236Action,
	{_State331, TraitOpTypeType}:                         _GotoState129Action,
	{_State331, FieldDefType}:                            _GotoState402Action,
	{_State331, ImplicitStructDefType}:                   _GotoState122Action,
	{_State331, ExplicitStructDefType}:                   _GotoState74Action,
	{_State331, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State331, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State331, TraitDefType}:                            _GotoState128Action,
	{_State331, FuncTypeType}:                            _GotoState120Action,
	{_State332, IdentifierToken}:                         _GotoState403Action,
	{_State332, LparenToken}:                             _GotoState226Action,
	{_State335, NewlinesToken}:                           _GotoState405Action,
	{_State335, CommaToken}:                              _GotoState404Action,
	{_State336, RparenToken}:                             _GotoState406Action,
	{_State340, IdentifierToken}:                         _GotoState112Action,
	{_State340, StructToken}:                             _GotoState50Action,
	{_State340, EnumToken}:                               _GotoState109Action,
	{_State340, TraitToken}:                              _GotoState117Action,
	{_State340, FuncToken}:                               _GotoState111Action,
	{_State340, LparenToken}:                             _GotoState113Action,
	{_State340, LbracketToken}:                           _GotoState42Action,
	{_State340, DotToken}:                                _GotoState108Action,
	{_State340, QuestionToken}:                           _GotoState115Action,
	{_State340, ExclaimToken}:                            _GotoState110Action,
	{_State340, TildeTildeToken}:                         _GotoState116Action,
	{_State340, BitNegToken}:                             _GotoState107Action,
	{_State340, BitAndToken}:                             _GotoState106Action,
	{_State340, ParseErrorToken}:                         _GotoState114Action,
	{_State340, InitializableTypeType}:                   _GotoState123Action,
	{_State340, SliceTypeType}:                           _GotoState100Action,
	{_State340, ArrayTypeType}:                           _GotoState59Action,
	{_State340, MapTypeType}:                             _GotoState88Action,
	{_State340, AtomTypeType}:                            _GotoState118Action,
	{_State340, ParseErrorTypeType}:                      _GotoState124Action,
	{_State340, ReturnableTypeType}:                      _GotoState127Action,
	{_State340, PrefixedTypeType}:                        _GotoState126Action,
	{_State340, PrefixTypeOpType}:                        _GotoState125Action,
	{_State340, ValueTypeType}:                           _GotoState407Action,
	{_State340, TraitOpTypeType}:                         _GotoState129Action,
	{_State340, ImplicitStructDefType}:                   _GotoState122Action,
	{_State340, ExplicitStructDefType}:                   _GotoState74Action,
	{_State340, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State340, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State340, TraitDefType}:                            _GotoState128Action,
	{_State340, FuncTypeType}:                            _GotoState120Action,
	{_State342, RbracketToken}:                           _GotoState408Action,
	{_State343, CommaToken}:                              _GotoState409Action,
	{_State344, IdentifierToken}:                         _GotoState246Action,
	{_State344, ParameterDefType}:                        _GotoState268Action,
	{_State344, ProperParameterDefsType}:                 _GotoState270Action,
	{_State344, ParameterDefsType}:                       _GotoState410Action,
	{_State345, IdentifierToken}:                         _GotoState112Action,
	{_State345, StructToken}:                             _GotoState50Action,
	{_State345, EnumToken}:                               _GotoState109Action,
	{_State345, TraitToken}:                              _GotoState117Action,
	{_State345, FuncToken}:                               _GotoState111Action,
	{_State345, LparenToken}:                             _GotoState113Action,
	{_State345, LbracketToken}:                           _GotoState42Action,
	{_State345, DotToken}:                                _GotoState108Action,
	{_State345, QuestionToken}:                           _GotoState115Action,
	{_State345, ExclaimToken}:                            _GotoState110Action,
	{_State345, TildeTildeToken}:                         _GotoState116Action,
	{_State345, BitNegToken}:                             _GotoState107Action,
	{_State345, BitAndToken}:                             _GotoState106Action,
	{_State345, ParseErrorToken}:                         _GotoState114Action,
	{_State345, InitializableTypeType}:                   _GotoState123Action,
	{_State345, SliceTypeType}:                           _GotoState100Action,
	{_State345, ArrayTypeType}:                           _GotoState59Action,
	{_State345, MapTypeType}:                             _GotoState88Action,
	{_State345, AtomTypeType}:                            _GotoState118Action,
	{_State345, ParseErrorTypeType}:                      _GotoState124Action,
	{_State345, ReturnableTypeType}:                      _GotoState127Action,
	{_State345, PrefixedTypeType}:                        _GotoState126Action,
	{_State345, PrefixTypeOpType}:                        _GotoState125Action,
	{_State345, ValueTypeType}:                           _GotoState411Action,
	{_State345, TraitOpTypeType}:                         _GotoState129Action,
	{_State345, ImplicitStructDefType}:                   _GotoState122Action,
	{_State345, ExplicitStructDefType}:                   _GotoState74Action,
	{_State345, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State345, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State345, TraitDefType}:                            _GotoState128Action,
	{_State345, FuncTypeType}:                            _GotoState120Action,
	{_State346, AddToken}:                                _GotoState239Action,
	{_State346, SubToken}:                                _GotoState241Action,
	{_State346, MulToken}:                                _GotoState240Action,
	{_State346, TraitOpType}:                             _GotoState242Action,
	{_State347, IdentifierToken}:                         _GotoState412Action,
	{_State350, AddToken}:                                _GotoState239Action,
	{_State350, SubToken}:                                _GotoState241Action,
	{_State350, MulToken}:                                _GotoState240Action,
	{_State350, TraitOpType}:                             _GotoState242Action,
	{_State351, ImplementsToken}:                         _GotoState413Action,
	{_State351, AddToken}:                                _GotoState239Action,
	{_State351, SubToken}:                                _GotoState241Action,
	{_State351, MulToken}:                                _GotoState240Action,
	{_State351, TraitOpType}:                             _GotoState242Action,
	{_State352, IdentifierToken}:                         _GotoState139Action,
	{_State352, LparenToken}:                             _GotoState140Action,
	{_State352, VarPatternType}:                          _GotoState414Action,
	{_State352, TuplePatternType}:                        _GotoState141Action,
	{_State353, IdentifierToken}:                         _GotoState256Action,
	{_State353, LparenToken}:                             _GotoState140Action,
	{_State353, EllipsisToken}:                           _GotoState255Action,
	{_State353, VarPatternType}:                          _GotoState259Action,
	{_State353, TuplePatternType}:                        _GotoState141Action,
	{_State353, FieldVarPatternType}:                     _GotoState415Action,
	{_State356, LparenToken}:                             _GotoState140Action,
	{_State356, TuplePatternType}:                        _GotoState416Action,
	{_State359, IdentifierToken}:                         _GotoState112Action,
	{_State359, StructToken}:                             _GotoState50Action,
	{_State359, EnumToken}:                               _GotoState109Action,
	{_State359, TraitToken}:                              _GotoState117Action,
	{_State359, FuncToken}:                               _GotoState111Action,
	{_State359, LparenToken}:                             _GotoState113Action,
	{_State359, LbracketToken}:                           _GotoState42Action,
	{_State359, DotToken}:                                _GotoState108Action,
	{_State359, QuestionToken}:                           _GotoState115Action,
	{_State359, ExclaimToken}:                            _GotoState110Action,
	{_State359, TildeTildeToken}:                         _GotoState116Action,
	{_State359, BitNegToken}:                             _GotoState107Action,
	{_State359, BitAndToken}:                             _GotoState106Action,
	{_State359, ParseErrorToken}:                         _GotoState114Action,
	{_State359, InitializableTypeType}:                   _GotoState123Action,
	{_State359, SliceTypeType}:                           _GotoState100Action,
	{_State359, ArrayTypeType}:                           _GotoState59Action,
	{_State359, MapTypeType}:                             _GotoState88Action,
	{_State359, AtomTypeType}:                            _GotoState118Action,
	{_State359, ParseErrorTypeType}:                      _GotoState124Action,
	{_State359, ReturnableTypeType}:                      _GotoState418Action,
	{_State359, PrefixedTypeType}:                        _GotoState126Action,
	{_State359, PrefixTypeOpType}:                        _GotoState125Action,
	{_State359, ImplicitStructDefType}:                   _GotoState122Action,
	{_State359, ExplicitStructDefType}:                   _GotoState74Action,
	{_State359, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State359, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State359, TraitDefType}:                            _GotoState128Action,
	{_State359, ReturnTypeType}:                          _GotoState417Action,
	{_State359, FuncTypeType}:                            _GotoState120Action,
	{_State360, IdentifierToken}:                         _GotoState246Action,
	{_State360, ParameterDefType}:                        _GotoState419Action,
	{_State362, StringLiteralToken}:                      _GotoState154Action,
	{_State362, ImportClauseType}:                        _GotoState420Action,
	{_State363, StringLiteralToken}:                      _GotoState154Action,
	{_State363, ImportClauseType}:                        _GotoState421Action,
	{_State365, RbracketToken}:                           _GotoState422Action,
	{_State365, AddToken}:                                _GotoState239Action,
	{_State365, SubToken}:                                _GotoState241Action,
	{_State365, MulToken}:                                _GotoState240Action,
	{_State365, TraitOpType}:                             _GotoState242Action,
	{_State366, RbracketToken}:                           _GotoState423Action,
	{_State372, IdentifierToken}:                         _GotoState229Action,
	{_State372, UnsafeToken}:                             _GotoState53Action,
	{_State372, StructToken}:                             _GotoState50Action,
	{_State372, EnumToken}:                               _GotoState109Action,
	{_State372, TraitToken}:                              _GotoState117Action,
	{_State372, FuncToken}:                               _GotoState111Action,
	{_State372, LparenToken}:                             _GotoState113Action,
	{_State372, LbracketToken}:                           _GotoState42Action,
	{_State372, DotToken}:                                _GotoState108Action,
	{_State372, QuestionToken}:                           _GotoState115Action,
	{_State372, ExclaimToken}:                            _GotoState110Action,
	{_State372, TildeTildeToken}:                         _GotoState116Action,
	{_State372, BitNegToken}:                             _GotoState107Action,
	{_State372, BitAndToken}:                             _GotoState106Action,
	{_State372, ParseErrorToken}:                         _GotoState114Action,
	{_State372, UnsafeStatementType}:                     _GotoState235Action,
	{_State372, InitializableTypeType}:                   _GotoState123Action,
	{_State372, SliceTypeType}:                           _GotoState100Action,
	{_State372, ArrayTypeType}:                           _GotoState59Action,
	{_State372, MapTypeType}:                             _GotoState88Action,
	{_State372, AtomTypeType}:                            _GotoState118Action,
	{_State372, ParseErrorTypeType}:                      _GotoState124Action,
	{_State372, ReturnableTypeType}:                      _GotoState127Action,
	{_State372, PrefixedTypeType}:                        _GotoState126Action,
	{_State372, PrefixTypeOpType}:                        _GotoState125Action,
	{_State372, ValueTypeType}:                           _GotoState236Action,
	{_State372, TraitOpTypeType}:                         _GotoState129Action,
	{_State372, FieldDefType}:                            _GotoState424Action,
	{_State372, ImplicitStructDefType}:                   _GotoState122Action,
	{_State372, ExplicitStructDefType}:                   _GotoState74Action,
	{_State372, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State372, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State372, TraitDefType}:                            _GotoState128Action,
	{_State372, FuncTypeType}:                            _GotoState120Action,
	{_State373, IdentifierToken}:                         _GotoState229Action,
	{_State373, UnsafeToken}:                             _GotoState53Action,
	{_State373, StructToken}:                             _GotoState50Action,
	{_State373, EnumToken}:                               _GotoState109Action,
	{_State373, TraitToken}:                              _GotoState117Action,
	{_State373, FuncToken}:                               _GotoState111Action,
	{_State373, LparenToken}:                             _GotoState113Action,
	{_State373, LbracketToken}:                           _GotoState42Action,
	{_State373, DotToken}:                                _GotoState108Action,
	{_State373, QuestionToken}:                           _GotoState115Action,
	{_State373, ExclaimToken}:                            _GotoState110Action,
	{_State373, TildeTildeToken}:                         _GotoState116Action,
	{_State373, BitNegToken}:                             _GotoState107Action,
	{_State373, BitAndToken}:                             _GotoState106Action,
	{_State373, ParseErrorToken}:                         _GotoState114Action,
	{_State373, UnsafeStatementType}:                     _GotoState235Action,
	{_State373, InitializableTypeType}:                   _GotoState123Action,
	{_State373, SliceTypeType}:                           _GotoState100Action,
	{_State373, ArrayTypeType}:                           _GotoState59Action,
	{_State373, MapTypeType}:                             _GotoState88Action,
	{_State373, AtomTypeType}:                            _GotoState118Action,
	{_State373, ParseErrorTypeType}:                      _GotoState124Action,
	{_State373, ReturnableTypeType}:                      _GotoState127Action,
	{_State373, PrefixedTypeType}:                        _GotoState126Action,
	{_State373, PrefixTypeOpType}:                        _GotoState125Action,
	{_State373, ValueTypeType}:                           _GotoState236Action,
	{_State373, TraitOpTypeType}:                         _GotoState129Action,
	{_State373, FieldDefType}:                            _GotoState425Action,
	{_State373, ImplicitStructDefType}:                   _GotoState122Action,
	{_State373, ExplicitStructDefType}:                   _GotoState74Action,
	{_State373, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State373, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State373, TraitDefType}:                            _GotoState128Action,
	{_State373, FuncTypeType}:                            _GotoState120Action,
	{_State374, StringLiteralToken}:                      _GotoState426Action,
	{_State376, IdentifierToken}:                         _GotoState112Action,
	{_State376, StructToken}:                             _GotoState50Action,
	{_State376, EnumToken}:                               _GotoState109Action,
	{_State376, TraitToken}:                              _GotoState117Action,
	{_State376, FuncToken}:                               _GotoState111Action,
	{_State376, LparenToken}:                             _GotoState113Action,
	{_State376, LbracketToken}:                           _GotoState42Action,
	{_State376, DotToken}:                                _GotoState108Action,
	{_State376, QuestionToken}:                           _GotoState115Action,
	{_State376, ExclaimToken}:                            _GotoState110Action,
	{_State376, TildeTildeToken}:                         _GotoState116Action,
	{_State376, BitNegToken}:                             _GotoState107Action,
	{_State376, BitAndToken}:                             _GotoState106Action,
	{_State376, ParseErrorToken}:                         _GotoState114Action,
	{_State376, InitializableTypeType}:                   _GotoState123Action,
	{_State376, SliceTypeType}:                           _GotoState100Action,
	{_State376, ArrayTypeType}:                           _GotoState59Action,
	{_State376, MapTypeType}:                             _GotoState88Action,
	{_State376, AtomTypeType}:                            _GotoState118Action,
	{_State376, ParseErrorTypeType}:                      _GotoState124Action,
	{_State376, ReturnableTypeType}:                      _GotoState127Action,
	{_State376, PrefixedTypeType}:                        _GotoState126Action,
	{_State376, PrefixTypeOpType}:                        _GotoState125Action,
	{_State376, ValueTypeType}:                           _GotoState427Action,
	{_State376, TraitOpTypeType}:                         _GotoState129Action,
	{_State376, ImplicitStructDefType}:                   _GotoState122Action,
	{_State376, ExplicitStructDefType}:                   _GotoState74Action,
	{_State376, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State376, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State376, TraitDefType}:                            _GotoState128Action,
	{_State376, FuncTypeType}:                            _GotoState120Action,
	{_State378, RparenToken}:                             _GotoState428Action,
	{_State380, IntegerLiteralToken}:                     _GotoState40Action,
	{_State380, FloatLiteralToken}:                       _GotoState35Action,
	{_State380, RuneLiteralToken}:                        _GotoState48Action,
	{_State380, StringLiteralToken}:                      _GotoState49Action,
	{_State380, IdentifierToken}:                         _GotoState38Action,
	{_State380, TrueToken}:                               _GotoState52Action,
	{_State380, FalseToken}:                              _GotoState34Action,
	{_State380, StructToken}:                             _GotoState50Action,
	{_State380, FuncToken}:                               _GotoState36Action,
	{_State380, VarToken}:                                _GotoState15Action,
	{_State380, LetToken}:                                _GotoState12Action,
	{_State380, NotToken}:                                _GotoState45Action,
	{_State380, LabelDeclToken}:                          _GotoState41Action,
	{_State380, LparenToken}:                             _GotoState43Action,
	{_State380, LbracketToken}:                           _GotoState42Action,
	{_State380, SubToken}:                                _GotoState51Action,
	{_State380, MulToken}:                                _GotoState44Action,
	{_State380, BitNegToken}:                             _GotoState27Action,
	{_State380, BitAndToken}:                             _GotoState26Action,
	{_State380, GreaterToken}:                            _GotoState37Action,
	{_State380, ParseErrorToken}:                         _GotoState46Action,
	{_State380, VarDeclPatternType}:                      _GotoState103Action,
	{_State380, VarOrLetType}:                            _GotoState24Action,
	{_State380, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State380, SequenceExprType}:                        _GotoState429Action,
	{_State380, CallExprType}:                            _GotoState70Action,
	{_State380, AtomExprType}:                            _GotoState62Action,
	{_State380, ParseErrorExprType}:                      _GotoState92Action,
	{_State380, LiteralExprType}:                         _GotoState87Action,
	{_State380, IdentifierExprType}:                      _GotoState79Action,
	{_State380, BlockExprType}:                           _GotoState69Action,
	{_State380, InitializeExprType}:                      _GotoState84Action,
	{_State380, ImplicitStructExprType}:                  _GotoState80Action,
	{_State380, AccessibleExprType}:                      _GotoState104Action,
	{_State380, AccessExprType}:                          _GotoState54Action,
	{_State380, IndexExprType}:                           _GotoState82Action,
	{_State380, PostfixableExprType}:                     _GotoState94Action,
	{_State380, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State380, PrefixableExprType}:                      _GotoState97Action,
	{_State380, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State380, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State380, MulExprType}:                             _GotoState89Action,
	{_State380, BinaryMulExprType}:                       _GotoState66Action,
	{_State380, AddExprType}:                             _GotoState56Action,
	{_State380, BinaryAddExprType}:                       _GotoState63Action,
	{_State380, CmpExprType}:                             _GotoState73Action,
	{_State380, BinaryCmpExprType}:                       _GotoState65Action,
	{_State380, AndExprType}:                             _GotoState57Action,
	{_State380, BinaryAndExprType}:                       _GotoState64Action,
	{_State380, OrExprType}:                              _GotoState91Action,
	{_State380, BinaryOrExprType}:                        _GotoState68Action,
	{_State380, InitializableTypeType}:                   _GotoState83Action,
	{_State380, SliceTypeType}:                           _GotoState100Action,
	{_State380, ArrayTypeType}:                           _GotoState59Action,
	{_State380, MapTypeType}:                             _GotoState88Action,
	{_State380, ExplicitStructDefType}:                   _GotoState74Action,
	{_State380, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State381, IntegerLiteralToken}:                     _GotoState40Action,
	{_State381, FloatLiteralToken}:                       _GotoState35Action,
	{_State381, RuneLiteralToken}:                        _GotoState48Action,
	{_State381, StringLiteralToken}:                      _GotoState49Action,
	{_State381, IdentifierToken}:                         _GotoState38Action,
	{_State381, TrueToken}:                               _GotoState52Action,
	{_State381, FalseToken}:                              _GotoState34Action,
	{_State381, StructToken}:                             _GotoState50Action,
	{_State381, FuncToken}:                               _GotoState36Action,
	{_State381, VarToken}:                                _GotoState15Action,
	{_State381, LetToken}:                                _GotoState12Action,
	{_State381, NotToken}:                                _GotoState45Action,
	{_State381, LabelDeclToken}:                          _GotoState41Action,
	{_State381, LparenToken}:                             _GotoState43Action,
	{_State381, LbracketToken}:                           _GotoState42Action,
	{_State381, SubToken}:                                _GotoState51Action,
	{_State381, MulToken}:                                _GotoState44Action,
	{_State381, BitNegToken}:                             _GotoState27Action,
	{_State381, BitAndToken}:                             _GotoState26Action,
	{_State381, GreaterToken}:                            _GotoState37Action,
	{_State381, ParseErrorToken}:                         _GotoState46Action,
	{_State381, VarDeclPatternType}:                      _GotoState103Action,
	{_State381, VarOrLetType}:                            _GotoState24Action,
	{_State381, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State381, SequenceExprType}:                        _GotoState430Action,
	{_State381, CallExprType}:                            _GotoState70Action,
	{_State381, AtomExprType}:                            _GotoState62Action,
	{_State381, ParseErrorExprType}:                      _GotoState92Action,
	{_State381, LiteralExprType}:                         _GotoState87Action,
	{_State381, IdentifierExprType}:                      _GotoState79Action,
	{_State381, BlockExprType}:                           _GotoState69Action,
	{_State381, InitializeExprType}:                      _GotoState84Action,
	{_State381, ImplicitStructExprType}:                  _GotoState80Action,
	{_State381, AccessibleExprType}:                      _GotoState104Action,
	{_State381, AccessExprType}:                          _GotoState54Action,
	{_State381, IndexExprType}:                           _GotoState82Action,
	{_State381, PostfixableExprType}:                     _GotoState94Action,
	{_State381, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State381, PrefixableExprType}:                      _GotoState97Action,
	{_State381, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State381, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State381, MulExprType}:                             _GotoState89Action,
	{_State381, BinaryMulExprType}:                       _GotoState66Action,
	{_State381, AddExprType}:                             _GotoState56Action,
	{_State381, BinaryAddExprType}:                       _GotoState63Action,
	{_State381, CmpExprType}:                             _GotoState73Action,
	{_State381, BinaryCmpExprType}:                       _GotoState65Action,
	{_State381, AndExprType}:                             _GotoState57Action,
	{_State381, BinaryAndExprType}:                       _GotoState64Action,
	{_State381, OrExprType}:                              _GotoState91Action,
	{_State381, BinaryOrExprType}:                        _GotoState68Action,
	{_State381, InitializableTypeType}:                   _GotoState83Action,
	{_State381, SliceTypeType}:                           _GotoState100Action,
	{_State381, ArrayTypeType}:                           _GotoState59Action,
	{_State381, MapTypeType}:                             _GotoState88Action,
	{_State381, ExplicitStructDefType}:                   _GotoState74Action,
	{_State381, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State382, IntegerLiteralToken}:                     _GotoState40Action,
	{_State382, FloatLiteralToken}:                       _GotoState35Action,
	{_State382, RuneLiteralToken}:                        _GotoState48Action,
	{_State382, StringLiteralToken}:                      _GotoState49Action,
	{_State382, IdentifierToken}:                         _GotoState38Action,
	{_State382, TrueToken}:                               _GotoState52Action,
	{_State382, FalseToken}:                              _GotoState34Action,
	{_State382, StructToken}:                             _GotoState50Action,
	{_State382, FuncToken}:                               _GotoState36Action,
	{_State382, VarToken}:                                _GotoState15Action,
	{_State382, LetToken}:                                _GotoState12Action,
	{_State382, NotToken}:                                _GotoState45Action,
	{_State382, LabelDeclToken}:                          _GotoState41Action,
	{_State382, LparenToken}:                             _GotoState43Action,
	{_State382, LbracketToken}:                           _GotoState42Action,
	{_State382, SubToken}:                                _GotoState51Action,
	{_State382, MulToken}:                                _GotoState44Action,
	{_State382, BitNegToken}:                             _GotoState27Action,
	{_State382, BitAndToken}:                             _GotoState26Action,
	{_State382, GreaterToken}:                            _GotoState37Action,
	{_State382, ParseErrorToken}:                         _GotoState46Action,
	{_State382, VarDeclPatternType}:                      _GotoState103Action,
	{_State382, VarOrLetType}:                            _GotoState24Action,
	{_State382, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State382, SequenceExprType}:                        _GotoState431Action,
	{_State382, CallExprType}:                            _GotoState70Action,
	{_State382, AtomExprType}:                            _GotoState62Action,
	{_State382, ParseErrorExprType}:                      _GotoState92Action,
	{_State382, LiteralExprType}:                         _GotoState87Action,
	{_State382, IdentifierExprType}:                      _GotoState79Action,
	{_State382, BlockExprType}:                           _GotoState69Action,
	{_State382, InitializeExprType}:                      _GotoState84Action,
	{_State382, ImplicitStructExprType}:                  _GotoState80Action,
	{_State382, AccessibleExprType}:                      _GotoState104Action,
	{_State382, AccessExprType}:                          _GotoState54Action,
	{_State382, IndexExprType}:                           _GotoState82Action,
	{_State382, PostfixableExprType}:                     _GotoState94Action,
	{_State382, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State382, PrefixableExprType}:                      _GotoState97Action,
	{_State382, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State382, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State382, MulExprType}:                             _GotoState89Action,
	{_State382, BinaryMulExprType}:                       _GotoState66Action,
	{_State382, AddExprType}:                             _GotoState56Action,
	{_State382, BinaryAddExprType}:                       _GotoState63Action,
	{_State382, CmpExprType}:                             _GotoState73Action,
	{_State382, BinaryCmpExprType}:                       _GotoState65Action,
	{_State382, AndExprType}:                             _GotoState57Action,
	{_State382, BinaryAndExprType}:                       _GotoState64Action,
	{_State382, OrExprType}:                              _GotoState91Action,
	{_State382, BinaryOrExprType}:                        _GotoState68Action,
	{_State382, InitializableTypeType}:                   _GotoState83Action,
	{_State382, SliceTypeType}:                           _GotoState100Action,
	{_State382, ArrayTypeType}:                           _GotoState59Action,
	{_State382, MapTypeType}:                             _GotoState88Action,
	{_State382, ExplicitStructDefType}:                   _GotoState74Action,
	{_State382, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State383, IntegerLiteralToken}:                     _GotoState40Action,
	{_State383, FloatLiteralToken}:                       _GotoState35Action,
	{_State383, RuneLiteralToken}:                        _GotoState48Action,
	{_State383, StringLiteralToken}:                      _GotoState49Action,
	{_State383, IdentifierToken}:                         _GotoState38Action,
	{_State383, TrueToken}:                               _GotoState52Action,
	{_State383, FalseToken}:                              _GotoState34Action,
	{_State383, StructToken}:                             _GotoState50Action,
	{_State383, FuncToken}:                               _GotoState36Action,
	{_State383, VarToken}:                                _GotoState15Action,
	{_State383, LetToken}:                                _GotoState12Action,
	{_State383, NotToken}:                                _GotoState45Action,
	{_State383, LabelDeclToken}:                          _GotoState41Action,
	{_State383, LparenToken}:                             _GotoState43Action,
	{_State383, LbracketToken}:                           _GotoState42Action,
	{_State383, SubToken}:                                _GotoState51Action,
	{_State383, MulToken}:                                _GotoState44Action,
	{_State383, BitNegToken}:                             _GotoState27Action,
	{_State383, BitAndToken}:                             _GotoState26Action,
	{_State383, GreaterToken}:                            _GotoState37Action,
	{_State383, ParseErrorToken}:                         _GotoState46Action,
	{_State383, VarDeclPatternType}:                      _GotoState103Action,
	{_State383, VarOrLetType}:                            _GotoState24Action,
	{_State383, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State383, SequenceExprType}:                        _GotoState433Action,
	{_State383, OptionalSequenceExprType}:                _GotoState432Action,
	{_State383, CallExprType}:                            _GotoState70Action,
	{_State383, AtomExprType}:                            _GotoState62Action,
	{_State383, ParseErrorExprType}:                      _GotoState92Action,
	{_State383, LiteralExprType}:                         _GotoState87Action,
	{_State383, IdentifierExprType}:                      _GotoState79Action,
	{_State383, BlockExprType}:                           _GotoState69Action,
	{_State383, InitializeExprType}:                      _GotoState84Action,
	{_State383, ImplicitStructExprType}:                  _GotoState80Action,
	{_State383, AccessibleExprType}:                      _GotoState104Action,
	{_State383, AccessExprType}:                          _GotoState54Action,
	{_State383, IndexExprType}:                           _GotoState82Action,
	{_State383, PostfixableExprType}:                     _GotoState94Action,
	{_State383, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State383, PrefixableExprType}:                      _GotoState97Action,
	{_State383, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State383, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State383, MulExprType}:                             _GotoState89Action,
	{_State383, BinaryMulExprType}:                       _GotoState66Action,
	{_State383, AddExprType}:                             _GotoState56Action,
	{_State383, BinaryAddExprType}:                       _GotoState63Action,
	{_State383, CmpExprType}:                             _GotoState73Action,
	{_State383, BinaryCmpExprType}:                       _GotoState65Action,
	{_State383, AndExprType}:                             _GotoState57Action,
	{_State383, BinaryAndExprType}:                       _GotoState64Action,
	{_State383, OrExprType}:                              _GotoState91Action,
	{_State383, BinaryOrExprType}:                        _GotoState68Action,
	{_State383, InitializableTypeType}:                   _GotoState83Action,
	{_State383, SliceTypeType}:                           _GotoState100Action,
	{_State383, ArrayTypeType}:                           _GotoState59Action,
	{_State383, MapTypeType}:                             _GotoState88Action,
	{_State383, ExplicitStructDefType}:                   _GotoState74Action,
	{_State383, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State384, LbraceToken}:                             _GotoState11Action,
	{_State384, StatementBlockType}:                      _GotoState434Action,
	{_State385, CommaToken}:                              _GotoState265Action,
	{_State385, AssignToken}:                             _GotoState435Action,
	{_State386, ElseToken}:                               _GotoState436Action,
	{_State388, IdentifierToken}:                         _GotoState229Action,
	{_State388, UnsafeToken}:                             _GotoState53Action,
	{_State388, StructToken}:                             _GotoState50Action,
	{_State388, EnumToken}:                               _GotoState109Action,
	{_State388, TraitToken}:                              _GotoState117Action,
	{_State388, FuncToken}:                               _GotoState111Action,
	{_State388, LparenToken}:                             _GotoState113Action,
	{_State388, LbracketToken}:                           _GotoState42Action,
	{_State388, DotToken}:                                _GotoState108Action,
	{_State388, QuestionToken}:                           _GotoState115Action,
	{_State388, ExclaimToken}:                            _GotoState110Action,
	{_State388, TildeTildeToken}:                         _GotoState116Action,
	{_State388, BitNegToken}:                             _GotoState107Action,
	{_State388, BitAndToken}:                             _GotoState106Action,
	{_State388, ParseErrorToken}:                         _GotoState114Action,
	{_State388, UnsafeStatementType}:                     _GotoState235Action,
	{_State388, InitializableTypeType}:                   _GotoState123Action,
	{_State388, SliceTypeType}:                           _GotoState100Action,
	{_State388, ArrayTypeType}:                           _GotoState59Action,
	{_State388, MapTypeType}:                             _GotoState88Action,
	{_State388, AtomTypeType}:                            _GotoState118Action,
	{_State388, ParseErrorTypeType}:                      _GotoState124Action,
	{_State388, ReturnableTypeType}:                      _GotoState127Action,
	{_State388, PrefixedTypeType}:                        _GotoState126Action,
	{_State388, PrefixTypeOpType}:                        _GotoState125Action,
	{_State388, ValueTypeType}:                           _GotoState236Action,
	{_State388, TraitOpTypeType}:                         _GotoState129Action,
	{_State388, FieldDefType}:                            _GotoState315Action,
	{_State388, ImplicitStructDefType}:                   _GotoState122Action,
	{_State388, ExplicitStructDefType}:                   _GotoState74Action,
	{_State388, EnumValueDefType}:                        _GotoState437Action,
	{_State388, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State388, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State388, TraitDefType}:                            _GotoState128Action,
	{_State388, FuncTypeType}:                            _GotoState120Action,
	{_State389, IdentifierToken}:                         _GotoState229Action,
	{_State389, UnsafeToken}:                             _GotoState53Action,
	{_State389, StructToken}:                             _GotoState50Action,
	{_State389, EnumToken}:                               _GotoState109Action,
	{_State389, TraitToken}:                              _GotoState117Action,
	{_State389, FuncToken}:                               _GotoState111Action,
	{_State389, LparenToken}:                             _GotoState113Action,
	{_State389, LbracketToken}:                           _GotoState42Action,
	{_State389, DotToken}:                                _GotoState108Action,
	{_State389, QuestionToken}:                           _GotoState115Action,
	{_State389, ExclaimToken}:                            _GotoState110Action,
	{_State389, TildeTildeToken}:                         _GotoState116Action,
	{_State389, BitNegToken}:                             _GotoState107Action,
	{_State389, BitAndToken}:                             _GotoState106Action,
	{_State389, ParseErrorToken}:                         _GotoState114Action,
	{_State389, UnsafeStatementType}:                     _GotoState235Action,
	{_State389, InitializableTypeType}:                   _GotoState123Action,
	{_State389, SliceTypeType}:                           _GotoState100Action,
	{_State389, ArrayTypeType}:                           _GotoState59Action,
	{_State389, MapTypeType}:                             _GotoState88Action,
	{_State389, AtomTypeType}:                            _GotoState118Action,
	{_State389, ParseErrorTypeType}:                      _GotoState124Action,
	{_State389, ReturnableTypeType}:                      _GotoState127Action,
	{_State389, PrefixedTypeType}:                        _GotoState126Action,
	{_State389, PrefixTypeOpType}:                        _GotoState125Action,
	{_State389, ValueTypeType}:                           _GotoState236Action,
	{_State389, TraitOpTypeType}:                         _GotoState129Action,
	{_State389, FieldDefType}:                            _GotoState315Action,
	{_State389, ImplicitStructDefType}:                   _GotoState122Action,
	{_State389, ExplicitStructDefType}:                   _GotoState74Action,
	{_State389, EnumValueDefType}:                        _GotoState438Action,
	{_State389, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State389, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State389, TraitDefType}:                            _GotoState128Action,
	{_State389, FuncTypeType}:                            _GotoState120Action,
	{_State391, IdentifierToken}:                         _GotoState229Action,
	{_State391, UnsafeToken}:                             _GotoState53Action,
	{_State391, StructToken}:                             _GotoState50Action,
	{_State391, EnumToken}:                               _GotoState109Action,
	{_State391, TraitToken}:                              _GotoState117Action,
	{_State391, FuncToken}:                               _GotoState111Action,
	{_State391, LparenToken}:                             _GotoState113Action,
	{_State391, LbracketToken}:                           _GotoState42Action,
	{_State391, DotToken}:                                _GotoState108Action,
	{_State391, QuestionToken}:                           _GotoState115Action,
	{_State391, ExclaimToken}:                            _GotoState110Action,
	{_State391, TildeTildeToken}:                         _GotoState116Action,
	{_State391, BitNegToken}:                             _GotoState107Action,
	{_State391, BitAndToken}:                             _GotoState106Action,
	{_State391, ParseErrorToken}:                         _GotoState114Action,
	{_State391, UnsafeStatementType}:                     _GotoState235Action,
	{_State391, InitializableTypeType}:                   _GotoState123Action,
	{_State391, SliceTypeType}:                           _GotoState100Action,
	{_State391, ArrayTypeType}:                           _GotoState59Action,
	{_State391, MapTypeType}:                             _GotoState88Action,
	{_State391, AtomTypeType}:                            _GotoState118Action,
	{_State391, ParseErrorTypeType}:                      _GotoState124Action,
	{_State391, ReturnableTypeType}:                      _GotoState127Action,
	{_State391, PrefixedTypeType}:                        _GotoState126Action,
	{_State391, PrefixTypeOpType}:                        _GotoState125Action,
	{_State391, ValueTypeType}:                           _GotoState236Action,
	{_State391, TraitOpTypeType}:                         _GotoState129Action,
	{_State391, FieldDefType}:                            _GotoState315Action,
	{_State391, ImplicitStructDefType}:                   _GotoState122Action,
	{_State391, ExplicitStructDefType}:                   _GotoState74Action,
	{_State391, EnumValueDefType}:                        _GotoState439Action,
	{_State391, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State391, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State391, TraitDefType}:                            _GotoState128Action,
	{_State391, FuncTypeType}:                            _GotoState120Action,
	{_State392, IdentifierToken}:                         _GotoState229Action,
	{_State392, UnsafeToken}:                             _GotoState53Action,
	{_State392, StructToken}:                             _GotoState50Action,
	{_State392, EnumToken}:                               _GotoState109Action,
	{_State392, TraitToken}:                              _GotoState117Action,
	{_State392, FuncToken}:                               _GotoState111Action,
	{_State392, LparenToken}:                             _GotoState113Action,
	{_State392, LbracketToken}:                           _GotoState42Action,
	{_State392, DotToken}:                                _GotoState108Action,
	{_State392, QuestionToken}:                           _GotoState115Action,
	{_State392, ExclaimToken}:                            _GotoState110Action,
	{_State392, TildeTildeToken}:                         _GotoState116Action,
	{_State392, BitNegToken}:                             _GotoState107Action,
	{_State392, BitAndToken}:                             _GotoState106Action,
	{_State392, ParseErrorToken}:                         _GotoState114Action,
	{_State392, UnsafeStatementType}:                     _GotoState235Action,
	{_State392, InitializableTypeType}:                   _GotoState123Action,
	{_State392, SliceTypeType}:                           _GotoState100Action,
	{_State392, ArrayTypeType}:                           _GotoState59Action,
	{_State392, MapTypeType}:                             _GotoState88Action,
	{_State392, AtomTypeType}:                            _GotoState118Action,
	{_State392, ParseErrorTypeType}:                      _GotoState124Action,
	{_State392, ReturnableTypeType}:                      _GotoState127Action,
	{_State392, PrefixedTypeType}:                        _GotoState126Action,
	{_State392, PrefixTypeOpType}:                        _GotoState125Action,
	{_State392, ValueTypeType}:                           _GotoState236Action,
	{_State392, TraitOpTypeType}:                         _GotoState129Action,
	{_State392, FieldDefType}:                            _GotoState315Action,
	{_State392, ImplicitStructDefType}:                   _GotoState122Action,
	{_State392, ExplicitStructDefType}:                   _GotoState74Action,
	{_State392, EnumValueDefType}:                        _GotoState440Action,
	{_State392, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State392, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State392, TraitDefType}:                            _GotoState128Action,
	{_State392, FuncTypeType}:                            _GotoState120Action,
	{_State393, AddToken}:                                _GotoState239Action,
	{_State393, SubToken}:                                _GotoState241Action,
	{_State393, MulToken}:                                _GotoState240Action,
	{_State393, TraitOpType}:                             _GotoState242Action,
	{_State394, IdentifierToken}:                         _GotoState112Action,
	{_State394, StructToken}:                             _GotoState50Action,
	{_State394, EnumToken}:                               _GotoState109Action,
	{_State394, TraitToken}:                              _GotoState117Action,
	{_State394, FuncToken}:                               _GotoState111Action,
	{_State394, LparenToken}:                             _GotoState113Action,
	{_State394, LbracketToken}:                           _GotoState42Action,
	{_State394, DotToken}:                                _GotoState108Action,
	{_State394, QuestionToken}:                           _GotoState115Action,
	{_State394, ExclaimToken}:                            _GotoState110Action,
	{_State394, TildeTildeToken}:                         _GotoState116Action,
	{_State394, BitNegToken}:                             _GotoState107Action,
	{_State394, BitAndToken}:                             _GotoState106Action,
	{_State394, ParseErrorToken}:                         _GotoState114Action,
	{_State394, InitializableTypeType}:                   _GotoState123Action,
	{_State394, SliceTypeType}:                           _GotoState100Action,
	{_State394, ArrayTypeType}:                           _GotoState59Action,
	{_State394, MapTypeType}:                             _GotoState88Action,
	{_State394, AtomTypeType}:                            _GotoState118Action,
	{_State394, ParseErrorTypeType}:                      _GotoState124Action,
	{_State394, ReturnableTypeType}:                      _GotoState127Action,
	{_State394, PrefixedTypeType}:                        _GotoState126Action,
	{_State394, PrefixTypeOpType}:                        _GotoState125Action,
	{_State394, ValueTypeType}:                           _GotoState441Action,
	{_State394, TraitOpTypeType}:                         _GotoState129Action,
	{_State394, ImplicitStructDefType}:                   _GotoState122Action,
	{_State394, ExplicitStructDefType}:                   _GotoState74Action,
	{_State394, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State394, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State394, TraitDefType}:                            _GotoState128Action,
	{_State394, FuncTypeType}:                            _GotoState120Action,
	{_State395, AddToken}:                                _GotoState239Action,
	{_State395, SubToken}:                                _GotoState241Action,
	{_State395, MulToken}:                                _GotoState240Action,
	{_State395, TraitOpType}:                             _GotoState242Action,
	{_State396, IdentifierToken}:                         _GotoState112Action,
	{_State396, StructToken}:                             _GotoState50Action,
	{_State396, EnumToken}:                               _GotoState109Action,
	{_State396, TraitToken}:                              _GotoState117Action,
	{_State396, FuncToken}:                               _GotoState111Action,
	{_State396, LparenToken}:                             _GotoState113Action,
	{_State396, LbracketToken}:                           _GotoState42Action,
	{_State396, DotToken}:                                _GotoState108Action,
	{_State396, QuestionToken}:                           _GotoState115Action,
	{_State396, ExclaimToken}:                            _GotoState110Action,
	{_State396, TildeTildeToken}:                         _GotoState116Action,
	{_State396, BitNegToken}:                             _GotoState107Action,
	{_State396, BitAndToken}:                             _GotoState106Action,
	{_State396, ParseErrorToken}:                         _GotoState114Action,
	{_State396, InitializableTypeType}:                   _GotoState123Action,
	{_State396, SliceTypeType}:                           _GotoState100Action,
	{_State396, ArrayTypeType}:                           _GotoState59Action,
	{_State396, MapTypeType}:                             _GotoState88Action,
	{_State396, AtomTypeType}:                            _GotoState118Action,
	{_State396, ParseErrorTypeType}:                      _GotoState124Action,
	{_State396, ReturnableTypeType}:                      _GotoState418Action,
	{_State396, PrefixedTypeType}:                        _GotoState126Action,
	{_State396, PrefixTypeOpType}:                        _GotoState125Action,
	{_State396, ImplicitStructDefType}:                   _GotoState122Action,
	{_State396, ExplicitStructDefType}:                   _GotoState74Action,
	{_State396, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State396, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State396, TraitDefType}:                            _GotoState128Action,
	{_State396, ReturnTypeType}:                          _GotoState442Action,
	{_State396, FuncTypeType}:                            _GotoState120Action,
	{_State397, IdentifierToken}:                         _GotoState318Action,
	{_State397, StructToken}:                             _GotoState50Action,
	{_State397, EnumToken}:                               _GotoState109Action,
	{_State397, TraitToken}:                              _GotoState117Action,
	{_State397, FuncToken}:                               _GotoState111Action,
	{_State397, LparenToken}:                             _GotoState113Action,
	{_State397, LbracketToken}:                           _GotoState42Action,
	{_State397, DotToken}:                                _GotoState108Action,
	{_State397, QuestionToken}:                           _GotoState115Action,
	{_State397, ExclaimToken}:                            _GotoState110Action,
	{_State397, EllipsisToken}:                           _GotoState317Action,
	{_State397, TildeTildeToken}:                         _GotoState116Action,
	{_State397, BitNegToken}:                             _GotoState107Action,
	{_State397, BitAndToken}:                             _GotoState106Action,
	{_State397, ParseErrorToken}:                         _GotoState114Action,
	{_State397, InitializableTypeType}:                   _GotoState123Action,
	{_State397, SliceTypeType}:                           _GotoState100Action,
	{_State397, ArrayTypeType}:                           _GotoState59Action,
	{_State397, MapTypeType}:                             _GotoState88Action,
	{_State397, AtomTypeType}:                            _GotoState118Action,
	{_State397, ParseErrorTypeType}:                      _GotoState124Action,
	{_State397, ReturnableTypeType}:                      _GotoState127Action,
	{_State397, PrefixedTypeType}:                        _GotoState126Action,
	{_State397, PrefixTypeOpType}:                        _GotoState125Action,
	{_State397, ValueTypeType}:                           _GotoState322Action,
	{_State397, TraitOpTypeType}:                         _GotoState129Action,
	{_State397, ImplicitStructDefType}:                   _GotoState122Action,
	{_State397, ExplicitStructDefType}:                   _GotoState74Action,
	{_State397, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State397, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State397, TraitDefType}:                            _GotoState128Action,
	{_State397, ParameterDeclType}:                       _GotoState443Action,
	{_State397, FuncTypeType}:                            _GotoState120Action,
	{_State403, LparenToken}:                             _GotoState444Action,
	{_State404, IdentifierToken}:                         _GotoState229Action,
	{_State404, UnsafeToken}:                             _GotoState53Action,
	{_State404, StructToken}:                             _GotoState50Action,
	{_State404, EnumToken}:                               _GotoState109Action,
	{_State404, TraitToken}:                              _GotoState117Action,
	{_State404, FuncToken}:                               _GotoState332Action,
	{_State404, LparenToken}:                             _GotoState113Action,
	{_State404, LbracketToken}:                           _GotoState42Action,
	{_State404, DotToken}:                                _GotoState108Action,
	{_State404, QuestionToken}:                           _GotoState115Action,
	{_State404, ExclaimToken}:                            _GotoState110Action,
	{_State404, TildeTildeToken}:                         _GotoState116Action,
	{_State404, BitNegToken}:                             _GotoState107Action,
	{_State404, BitAndToken}:                             _GotoState106Action,
	{_State404, ParseErrorToken}:                         _GotoState114Action,
	{_State404, UnsafeStatementType}:                     _GotoState235Action,
	{_State404, InitializableTypeType}:                   _GotoState123Action,
	{_State404, SliceTypeType}:                           _GotoState100Action,
	{_State404, ArrayTypeType}:                           _GotoState59Action,
	{_State404, MapTypeType}:                             _GotoState88Action,
	{_State404, AtomTypeType}:                            _GotoState118Action,
	{_State404, ParseErrorTypeType}:                      _GotoState124Action,
	{_State404, ReturnableTypeType}:                      _GotoState127Action,
	{_State404, PrefixedTypeType}:                        _GotoState126Action,
	{_State404, PrefixTypeOpType}:                        _GotoState125Action,
	{_State404, ValueTypeType}:                           _GotoState236Action,
	{_State404, TraitOpTypeType}:                         _GotoState129Action,
	{_State404, FieldDefType}:                            _GotoState333Action,
	{_State404, ImplicitStructDefType}:                   _GotoState122Action,
	{_State404, ExplicitStructDefType}:                   _GotoState74Action,
	{_State404, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State404, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State404, TraitPropertyType}:                       _GotoState445Action,
	{_State404, TraitDefType}:                            _GotoState128Action,
	{_State404, FuncTypeType}:                            _GotoState120Action,
	{_State404, MethodSignatureType}:                     _GotoState334Action,
	{_State405, IdentifierToken}:                         _GotoState229Action,
	{_State405, UnsafeToken}:                             _GotoState53Action,
	{_State405, StructToken}:                             _GotoState50Action,
	{_State405, EnumToken}:                               _GotoState109Action,
	{_State405, TraitToken}:                              _GotoState117Action,
	{_State405, FuncToken}:                               _GotoState332Action,
	{_State405, LparenToken}:                             _GotoState113Action,
	{_State405, LbracketToken}:                           _GotoState42Action,
	{_State405, DotToken}:                                _GotoState108Action,
	{_State405, QuestionToken}:                           _GotoState115Action,
	{_State405, ExclaimToken}:                            _GotoState110Action,
	{_State405, TildeTildeToken}:                         _GotoState116Action,
	{_State405, BitNegToken}:                             _GotoState107Action,
	{_State405, BitAndToken}:                             _GotoState106Action,
	{_State405, ParseErrorToken}:                         _GotoState114Action,
	{_State405, UnsafeStatementType}:                     _GotoState235Action,
	{_State405, InitializableTypeType}:                   _GotoState123Action,
	{_State405, SliceTypeType}:                           _GotoState100Action,
	{_State405, ArrayTypeType}:                           _GotoState59Action,
	{_State405, MapTypeType}:                             _GotoState88Action,
	{_State405, AtomTypeType}:                            _GotoState118Action,
	{_State405, ParseErrorTypeType}:                      _GotoState124Action,
	{_State405, ReturnableTypeType}:                      _GotoState127Action,
	{_State405, PrefixedTypeType}:                        _GotoState126Action,
	{_State405, PrefixTypeOpType}:                        _GotoState125Action,
	{_State405, ValueTypeType}:                           _GotoState236Action,
	{_State405, TraitOpTypeType}:                         _GotoState129Action,
	{_State405, FieldDefType}:                            _GotoState333Action,
	{_State405, ImplicitStructDefType}:                   _GotoState122Action,
	{_State405, ExplicitStructDefType}:                   _GotoState74Action,
	{_State405, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State405, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State405, TraitPropertyType}:                       _GotoState446Action,
	{_State405, TraitDefType}:                            _GotoState128Action,
	{_State405, FuncTypeType}:                            _GotoState120Action,
	{_State405, MethodSignatureType}:                     _GotoState334Action,
	{_State407, AddToken}:                                _GotoState239Action,
	{_State407, SubToken}:                                _GotoState241Action,
	{_State407, MulToken}:                                _GotoState240Action,
	{_State407, TraitOpType}:                             _GotoState242Action,
	{_State409, IdentifierToken}:                         _GotoState340Action,
	{_State409, GenericParameterDefType}:                 _GotoState447Action,
	{_State410, RparenToken}:                             _GotoState448Action,
	{_State411, AddToken}:                                _GotoState239Action,
	{_State411, SubToken}:                                _GotoState241Action,
	{_State411, MulToken}:                                _GotoState240Action,
	{_State411, TraitOpType}:                             _GotoState242Action,
	{_State412, DollarLbracketToken}:                     _GotoState244Action,
	{_State412, OptionalGenericParametersType}:           _GotoState449Action,
	{_State413, IdentifierToken}:                         _GotoState112Action,
	{_State413, StructToken}:                             _GotoState50Action,
	{_State413, EnumToken}:                               _GotoState109Action,
	{_State413, TraitToken}:                              _GotoState117Action,
	{_State413, FuncToken}:                               _GotoState111Action,
	{_State413, LparenToken}:                             _GotoState113Action,
	{_State413, LbracketToken}:                           _GotoState42Action,
	{_State413, DotToken}:                                _GotoState108Action,
	{_State413, QuestionToken}:                           _GotoState115Action,
	{_State413, ExclaimToken}:                            _GotoState110Action,
	{_State413, TildeTildeToken}:                         _GotoState116Action,
	{_State413, BitNegToken}:                             _GotoState107Action,
	{_State413, BitAndToken}:                             _GotoState106Action,
	{_State413, ParseErrorToken}:                         _GotoState114Action,
	{_State413, InitializableTypeType}:                   _GotoState123Action,
	{_State413, SliceTypeType}:                           _GotoState100Action,
	{_State413, ArrayTypeType}:                           _GotoState59Action,
	{_State413, MapTypeType}:                             _GotoState88Action,
	{_State413, AtomTypeType}:                            _GotoState118Action,
	{_State413, ParseErrorTypeType}:                      _GotoState124Action,
	{_State413, ReturnableTypeType}:                      _GotoState127Action,
	{_State413, PrefixedTypeType}:                        _GotoState126Action,
	{_State413, PrefixTypeOpType}:                        _GotoState125Action,
	{_State413, ValueTypeType}:                           _GotoState450Action,
	{_State413, TraitOpTypeType}:                         _GotoState129Action,
	{_State413, ImplicitStructDefType}:                   _GotoState122Action,
	{_State413, ExplicitStructDefType}:                   _GotoState74Action,
	{_State413, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State413, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State413, TraitDefType}:                            _GotoState128Action,
	{_State413, FuncTypeType}:                            _GotoState120Action,
	{_State417, LbraceToken}:                             _GotoState11Action,
	{_State417, StatementBlockType}:                      _GotoState451Action,
	{_State427, AddToken}:                                _GotoState239Action,
	{_State427, SubToken}:                                _GotoState241Action,
	{_State427, MulToken}:                                _GotoState240Action,
	{_State427, TraitOpType}:                             _GotoState242Action,
	{_State431, DoToken}:                                 _GotoState452Action,
	{_State432, SemicolonToken}:                          _GotoState453Action,
	{_State435, IntegerLiteralToken}:                     _GotoState40Action,
	{_State435, FloatLiteralToken}:                       _GotoState35Action,
	{_State435, RuneLiteralToken}:                        _GotoState48Action,
	{_State435, StringLiteralToken}:                      _GotoState49Action,
	{_State435, IdentifierToken}:                         _GotoState38Action,
	{_State435, TrueToken}:                               _GotoState52Action,
	{_State435, FalseToken}:                              _GotoState34Action,
	{_State435, StructToken}:                             _GotoState50Action,
	{_State435, FuncToken}:                               _GotoState36Action,
	{_State435, VarToken}:                                _GotoState15Action,
	{_State435, LetToken}:                                _GotoState12Action,
	{_State435, NotToken}:                                _GotoState45Action,
	{_State435, LabelDeclToken}:                          _GotoState41Action,
	{_State435, LparenToken}:                             _GotoState43Action,
	{_State435, LbracketToken}:                           _GotoState42Action,
	{_State435, SubToken}:                                _GotoState51Action,
	{_State435, MulToken}:                                _GotoState44Action,
	{_State435, BitNegToken}:                             _GotoState27Action,
	{_State435, BitAndToken}:                             _GotoState26Action,
	{_State435, GreaterToken}:                            _GotoState37Action,
	{_State435, ParseErrorToken}:                         _GotoState46Action,
	{_State435, VarDeclPatternType}:                      _GotoState103Action,
	{_State435, VarOrLetType}:                            _GotoState24Action,
	{_State435, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State435, SequenceExprType}:                        _GotoState454Action,
	{_State435, CallExprType}:                            _GotoState70Action,
	{_State435, AtomExprType}:                            _GotoState62Action,
	{_State435, ParseErrorExprType}:                      _GotoState92Action,
	{_State435, LiteralExprType}:                         _GotoState87Action,
	{_State435, IdentifierExprType}:                      _GotoState79Action,
	{_State435, BlockExprType}:                           _GotoState69Action,
	{_State435, InitializeExprType}:                      _GotoState84Action,
	{_State435, ImplicitStructExprType}:                  _GotoState80Action,
	{_State435, AccessibleExprType}:                      _GotoState104Action,
	{_State435, AccessExprType}:                          _GotoState54Action,
	{_State435, IndexExprType}:                           _GotoState82Action,
	{_State435, PostfixableExprType}:                     _GotoState94Action,
	{_State435, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State435, PrefixableExprType}:                      _GotoState97Action,
	{_State435, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State435, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State435, MulExprType}:                             _GotoState89Action,
	{_State435, BinaryMulExprType}:                       _GotoState66Action,
	{_State435, AddExprType}:                             _GotoState56Action,
	{_State435, BinaryAddExprType}:                       _GotoState63Action,
	{_State435, CmpExprType}:                             _GotoState73Action,
	{_State435, BinaryCmpExprType}:                       _GotoState65Action,
	{_State435, AndExprType}:                             _GotoState57Action,
	{_State435, BinaryAndExprType}:                       _GotoState64Action,
	{_State435, OrExprType}:                              _GotoState91Action,
	{_State435, BinaryOrExprType}:                        _GotoState68Action,
	{_State435, InitializableTypeType}:                   _GotoState83Action,
	{_State435, SliceTypeType}:                           _GotoState100Action,
	{_State435, ArrayTypeType}:                           _GotoState59Action,
	{_State435, MapTypeType}:                             _GotoState88Action,
	{_State435, ExplicitStructDefType}:                   _GotoState74Action,
	{_State435, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State436, IfToken}:                                 _GotoState216Action,
	{_State436, LbraceToken}:                             _GotoState11Action,
	{_State436, StatementBlockType}:                      _GotoState456Action,
	{_State436, IfExprType}:                              _GotoState455Action,
	{_State441, AddToken}:                                _GotoState239Action,
	{_State441, SubToken}:                                _GotoState241Action,
	{_State441, MulToken}:                                _GotoState240Action,
	{_State441, TraitOpType}:                             _GotoState242Action,
	{_State444, IdentifierToken}:                         _GotoState318Action,
	{_State444, StructToken}:                             _GotoState50Action,
	{_State444, EnumToken}:                               _GotoState109Action,
	{_State444, TraitToken}:                              _GotoState117Action,
	{_State444, FuncToken}:                               _GotoState111Action,
	{_State444, LparenToken}:                             _GotoState113Action,
	{_State444, LbracketToken}:                           _GotoState42Action,
	{_State444, DotToken}:                                _GotoState108Action,
	{_State444, QuestionToken}:                           _GotoState115Action,
	{_State444, ExclaimToken}:                            _GotoState110Action,
	{_State444, EllipsisToken}:                           _GotoState317Action,
	{_State444, TildeTildeToken}:                         _GotoState116Action,
	{_State444, BitNegToken}:                             _GotoState107Action,
	{_State444, BitAndToken}:                             _GotoState106Action,
	{_State444, ParseErrorToken}:                         _GotoState114Action,
	{_State444, InitializableTypeType}:                   _GotoState123Action,
	{_State444, SliceTypeType}:                           _GotoState100Action,
	{_State444, ArrayTypeType}:                           _GotoState59Action,
	{_State444, MapTypeType}:                             _GotoState88Action,
	{_State444, AtomTypeType}:                            _GotoState118Action,
	{_State444, ParseErrorTypeType}:                      _GotoState124Action,
	{_State444, ReturnableTypeType}:                      _GotoState127Action,
	{_State444, PrefixedTypeType}:                        _GotoState126Action,
	{_State444, PrefixTypeOpType}:                        _GotoState125Action,
	{_State444, ValueTypeType}:                           _GotoState322Action,
	{_State444, TraitOpTypeType}:                         _GotoState129Action,
	{_State444, ImplicitStructDefType}:                   _GotoState122Action,
	{_State444, ExplicitStructDefType}:                   _GotoState74Action,
	{_State444, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State444, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State444, TraitDefType}:                            _GotoState128Action,
	{_State444, ParameterDeclType}:                       _GotoState319Action,
	{_State444, ProperParameterDeclsType}:                _GotoState321Action,
	{_State444, ParameterDeclsType}:                      _GotoState457Action,
	{_State444, FuncTypeType}:                            _GotoState120Action,
	{_State448, IdentifierToken}:                         _GotoState112Action,
	{_State448, StructToken}:                             _GotoState50Action,
	{_State448, EnumToken}:                               _GotoState109Action,
	{_State448, TraitToken}:                              _GotoState117Action,
	{_State448, FuncToken}:                               _GotoState111Action,
	{_State448, LparenToken}:                             _GotoState113Action,
	{_State448, LbracketToken}:                           _GotoState42Action,
	{_State448, DotToken}:                                _GotoState108Action,
	{_State448, QuestionToken}:                           _GotoState115Action,
	{_State448, ExclaimToken}:                            _GotoState110Action,
	{_State448, TildeTildeToken}:                         _GotoState116Action,
	{_State448, BitNegToken}:                             _GotoState107Action,
	{_State448, BitAndToken}:                             _GotoState106Action,
	{_State448, ParseErrorToken}:                         _GotoState114Action,
	{_State448, InitializableTypeType}:                   _GotoState123Action,
	{_State448, SliceTypeType}:                           _GotoState100Action,
	{_State448, ArrayTypeType}:                           _GotoState59Action,
	{_State448, MapTypeType}:                             _GotoState88Action,
	{_State448, AtomTypeType}:                            _GotoState118Action,
	{_State448, ParseErrorTypeType}:                      _GotoState124Action,
	{_State448, ReturnableTypeType}:                      _GotoState418Action,
	{_State448, PrefixedTypeType}:                        _GotoState126Action,
	{_State448, PrefixTypeOpType}:                        _GotoState125Action,
	{_State448, ImplicitStructDefType}:                   _GotoState122Action,
	{_State448, ExplicitStructDefType}:                   _GotoState74Action,
	{_State448, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State448, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State448, TraitDefType}:                            _GotoState128Action,
	{_State448, ReturnTypeType}:                          _GotoState458Action,
	{_State448, FuncTypeType}:                            _GotoState120Action,
	{_State449, LparenToken}:                             _GotoState459Action,
	{_State450, AddToken}:                                _GotoState239Action,
	{_State450, SubToken}:                                _GotoState241Action,
	{_State450, MulToken}:                                _GotoState240Action,
	{_State450, TraitOpType}:                             _GotoState242Action,
	{_State452, LbraceToken}:                             _GotoState11Action,
	{_State452, StatementBlockType}:                      _GotoState460Action,
	{_State453, IntegerLiteralToken}:                     _GotoState40Action,
	{_State453, FloatLiteralToken}:                       _GotoState35Action,
	{_State453, RuneLiteralToken}:                        _GotoState48Action,
	{_State453, StringLiteralToken}:                      _GotoState49Action,
	{_State453, IdentifierToken}:                         _GotoState38Action,
	{_State453, TrueToken}:                               _GotoState52Action,
	{_State453, FalseToken}:                              _GotoState34Action,
	{_State453, StructToken}:                             _GotoState50Action,
	{_State453, FuncToken}:                               _GotoState36Action,
	{_State453, VarToken}:                                _GotoState15Action,
	{_State453, LetToken}:                                _GotoState12Action,
	{_State453, NotToken}:                                _GotoState45Action,
	{_State453, LabelDeclToken}:                          _GotoState41Action,
	{_State453, LparenToken}:                             _GotoState43Action,
	{_State453, LbracketToken}:                           _GotoState42Action,
	{_State453, SubToken}:                                _GotoState51Action,
	{_State453, MulToken}:                                _GotoState44Action,
	{_State453, BitNegToken}:                             _GotoState27Action,
	{_State453, BitAndToken}:                             _GotoState26Action,
	{_State453, GreaterToken}:                            _GotoState37Action,
	{_State453, ParseErrorToken}:                         _GotoState46Action,
	{_State453, VarDeclPatternType}:                      _GotoState103Action,
	{_State453, VarOrLetType}:                            _GotoState24Action,
	{_State453, OptionalLabelDeclType}:                   _GotoState148Action,
	{_State453, SequenceExprType}:                        _GotoState433Action,
	{_State453, OptionalSequenceExprType}:                _GotoState461Action,
	{_State453, CallExprType}:                            _GotoState70Action,
	{_State453, AtomExprType}:                            _GotoState62Action,
	{_State453, ParseErrorExprType}:                      _GotoState92Action,
	{_State453, LiteralExprType}:                         _GotoState87Action,
	{_State453, IdentifierExprType}:                      _GotoState79Action,
	{_State453, BlockExprType}:                           _GotoState69Action,
	{_State453, InitializeExprType}:                      _GotoState84Action,
	{_State453, ImplicitStructExprType}:                  _GotoState80Action,
	{_State453, AccessibleExprType}:                      _GotoState104Action,
	{_State453, AccessExprType}:                          _GotoState54Action,
	{_State453, IndexExprType}:                           _GotoState82Action,
	{_State453, PostfixableExprType}:                     _GotoState94Action,
	{_State453, PostfixUnaryExprType}:                    _GotoState93Action,
	{_State453, PrefixableExprType}:                      _GotoState97Action,
	{_State453, PrefixUnaryExprType}:                     _GotoState95Action,
	{_State453, PrefixUnaryOpType}:                       _GotoState96Action,
	{_State453, MulExprType}:                             _GotoState89Action,
	{_State453, BinaryMulExprType}:                       _GotoState66Action,
	{_State453, AddExprType}:                             _GotoState56Action,
	{_State453, BinaryAddExprType}:                       _GotoState63Action,
	{_State453, CmpExprType}:                             _GotoState73Action,
	{_State453, BinaryCmpExprType}:                       _GotoState65Action,
	{_State453, AndExprType}:                             _GotoState57Action,
	{_State453, BinaryAndExprType}:                       _GotoState64Action,
	{_State453, OrExprType}:                              _GotoState91Action,
	{_State453, BinaryOrExprType}:                        _GotoState68Action,
	{_State453, InitializableTypeType}:                   _GotoState83Action,
	{_State453, SliceTypeType}:                           _GotoState100Action,
	{_State453, ArrayTypeType}:                           _GotoState59Action,
	{_State453, MapTypeType}:                             _GotoState88Action,
	{_State453, ExplicitStructDefType}:                   _GotoState74Action,
	{_State453, AnonymousFuncExprType}:                   _GotoState58Action,
	{_State457, RparenToken}:                             _GotoState462Action,
	{_State458, LbraceToken}:                             _GotoState11Action,
	{_State458, StatementBlockType}:                      _GotoState463Action,
	{_State459, IdentifierToken}:                         _GotoState246Action,
	{_State459, ParameterDefType}:                        _GotoState268Action,
	{_State459, ProperParameterDefsType}:                 _GotoState270Action,
	{_State459, ParameterDefsType}:                       _GotoState464Action,
	{_State461, DoToken}:                                 _GotoState465Action,
	{_State462, IdentifierToken}:                         _GotoState112Action,
	{_State462, StructToken}:                             _GotoState50Action,
	{_State462, EnumToken}:                               _GotoState109Action,
	{_State462, TraitToken}:                              _GotoState117Action,
	{_State462, FuncToken}:                               _GotoState111Action,
	{_State462, LparenToken}:                             _GotoState113Action,
	{_State462, LbracketToken}:                           _GotoState42Action,
	{_State462, DotToken}:                                _GotoState108Action,
	{_State462, QuestionToken}:                           _GotoState115Action,
	{_State462, ExclaimToken}:                            _GotoState110Action,
	{_State462, TildeTildeToken}:                         _GotoState116Action,
	{_State462, BitNegToken}:                             _GotoState107Action,
	{_State462, BitAndToken}:                             _GotoState106Action,
	{_State462, ParseErrorToken}:                         _GotoState114Action,
	{_State462, InitializableTypeType}:                   _GotoState123Action,
	{_State462, SliceTypeType}:                           _GotoState100Action,
	{_State462, ArrayTypeType}:                           _GotoState59Action,
	{_State462, MapTypeType}:                             _GotoState88Action,
	{_State462, AtomTypeType}:                            _GotoState118Action,
	{_State462, ParseErrorTypeType}:                      _GotoState124Action,
	{_State462, ReturnableTypeType}:                      _GotoState418Action,
	{_State462, PrefixedTypeType}:                        _GotoState126Action,
	{_State462, PrefixTypeOpType}:                        _GotoState125Action,
	{_State462, ImplicitStructDefType}:                   _GotoState122Action,
	{_State462, ExplicitStructDefType}:                   _GotoState74Action,
	{_State462, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State462, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State462, TraitDefType}:                            _GotoState128Action,
	{_State462, ReturnTypeType}:                          _GotoState466Action,
	{_State462, FuncTypeType}:                            _GotoState120Action,
	{_State464, RparenToken}:                             _GotoState467Action,
	{_State465, LbraceToken}:                             _GotoState11Action,
	{_State465, StatementBlockType}:                      _GotoState468Action,
	{_State467, IdentifierToken}:                         _GotoState112Action,
	{_State467, StructToken}:                             _GotoState50Action,
	{_State467, EnumToken}:                               _GotoState109Action,
	{_State467, TraitToken}:                              _GotoState117Action,
	{_State467, FuncToken}:                               _GotoState111Action,
	{_State467, LparenToken}:                             _GotoState113Action,
	{_State467, LbracketToken}:                           _GotoState42Action,
	{_State467, DotToken}:                                _GotoState108Action,
	{_State467, QuestionToken}:                           _GotoState115Action,
	{_State467, ExclaimToken}:                            _GotoState110Action,
	{_State467, TildeTildeToken}:                         _GotoState116Action,
	{_State467, BitNegToken}:                             _GotoState107Action,
	{_State467, BitAndToken}:                             _GotoState106Action,
	{_State467, ParseErrorToken}:                         _GotoState114Action,
	{_State467, InitializableTypeType}:                   _GotoState123Action,
	{_State467, SliceTypeType}:                           _GotoState100Action,
	{_State467, ArrayTypeType}:                           _GotoState59Action,
	{_State467, MapTypeType}:                             _GotoState88Action,
	{_State467, AtomTypeType}:                            _GotoState118Action,
	{_State467, ParseErrorTypeType}:                      _GotoState124Action,
	{_State467, ReturnableTypeType}:                      _GotoState418Action,
	{_State467, PrefixedTypeType}:                        _GotoState126Action,
	{_State467, PrefixTypeOpType}:                        _GotoState125Action,
	{_State467, ImplicitStructDefType}:                   _GotoState122Action,
	{_State467, ExplicitStructDefType}:                   _GotoState74Action,
	{_State467, ImplicitEnumDefType}:                     _GotoState121Action,
	{_State467, ExplicitEnumDefType}:                     _GotoState119Action,
	{_State467, TraitDefType}:                            _GotoState128Action,
	{_State467, ReturnTypeType}:                          _GotoState469Action,
	{_State467, FuncTypeType}:                            _GotoState120Action,
	{_State469, LbraceToken}:                             _GotoState11Action,
	{_State469, StatementBlockType}:                      _GotoState470Action,
	{_State1, _EndMarker}:                                _ReduceNilToDefinitionsAction,
	{_State2, _WildcardMarker}:                           _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State3, _WildcardMarker}:                           _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State9, _WildcardMarker}:                           _ReduceCommentGroupsToDefinitionAction,
	{_State11, RbraceToken}:                              _ReduceNilToStatementsAction,
	{_State11, _WildcardMarker}:                          _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State12, _WildcardMarker}:                          _ReduceLetToVarOrLetAction,
	{_State13, _WildcardMarker}:                          _ReduceNoSpecToPackageDefAction,
	{_State15, _WildcardMarker}:                          _ReduceVarToVarOrLetAction,
	{_State16, _WildcardMarker}:                          _ReduceDefinitionToProperDefinitionsAction,
	{_State17, _EndMarker}:                               _ReduceToSourceAction,
	{_State18, _WildcardMarker}:                          _ReduceNamedFuncDefToDefinitionAction,
	{_State19, _WildcardMarker}:                          _ReducePackageDefToDefinitionAction,
	{_State20, _EndMarker}:                               _ReduceProperDefinitionsToDefinitionsAction,
	{_State21, _WildcardMarker}:                          _ReduceStatementBlockToDefinitionAction,
	{_State22, _WildcardMarker}:                          _ReduceTypeDefToDefinitionAction,
	{_State23, _WildcardMarker}:                          _ReduceGlobalVarDefToDefinitionAction,
	{_State25, _WildcardMarker}:                          _ReduceAsyncToCallbackOpAction,
	{_State26, _WildcardMarker}:                          _ReduceBitAndToPrefixUnaryOpAction,
	{_State27, _WildcardMarker}:                          _ReduceBitNegToPrefixUnaryOpAction,
	{_State28, _WildcardMarker}:                          _ReduceBreakToJumpTypeAction,
	{_State29, LbraceToken}:                              _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State30, _WildcardMarker}:                          _ReduceContinueToJumpTypeAction,
	{_State32, _WildcardMarker}:                          _ReduceDeferToCallbackOpAction,
	{_State33, _EndMarker}:                               _ReduceToFallthroughStatementAction,
	{_State34, _WildcardMarker}:                          _ReduceFalseToLiteralExprAction,
	{_State35, _WildcardMarker}:                          _ReduceFloatLiteralToLiteralExprAction,
	{_State37, LbraceToken}:                              _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State38, _WildcardMarker}:                          _ReduceToIdentifierExprAction,
	{_State40, _WildcardMarker}:                          _ReduceIntegerLiteralToLiteralExprAction,
	{_State41, _WildcardMarker}:                          _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State43, _WildcardMarker}:                          _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State43, RparenToken}:                              _ReduceNilToArgumentsAction,
	{_State44, _WildcardMarker}:                          _ReduceMulToPrefixUnaryOpAction,
	{_State45, _WildcardMarker}:                          _ReduceNotToPrefixUnaryOpAction,
	{_State46, _WildcardMarker}:                          _ReduceToParseErrorExprAction,
	{_State47, _WildcardMarker}:                          _ReduceReturnToJumpTypeAction,
	{_State48, _WildcardMarker}:                          _ReduceRuneLiteralToLiteralExprAction,
	{_State49, _WildcardMarker}:                          _ReduceStringLiteralToLiteralExprAction,
	{_State51, _WildcardMarker}:                          _ReduceSubToPrefixUnaryOpAction,
	{_State52, _WildcardMarker}:                          _ReduceTrueToLiteralExprAction,
	{_State54, _WildcardMarker}:                          _ReduceAccessExprToAccessibleExprAction,
	{_State55, _WildcardMarker}:                          _ReduceAccessibleExprToPostfixableExprAction,
	{_State55, LparenToken}:                              _ReduceNilToOptionalGenericBindingAction,
	{_State56, _WildcardMarker}:                          _ReduceAddExprToCmpExprAction,
	{_State57, _WildcardMarker}:                          _ReduceAndExprToOrExprAction,
	{_State58, _WildcardMarker}:                          _ReduceAnonymousFuncExprToAtomExprAction,
	{_State59, _WildcardMarker}:                          _ReduceArrayTypeToInitializableTypeAction,
	{_State61, _EndMarker}:                               _ReduceAssignStatementToSimpleStatementAction,
	{_State62, _WildcardMarker}:                          _ReduceAtomExprToAccessibleExprAction,
	{_State63, _WildcardMarker}:                          _ReduceBinaryAddExprToAddExprAction,
	{_State64, _WildcardMarker}:                          _ReduceBinaryAndExprToAndExprAction,
	{_State65, _WildcardMarker}:                          _ReduceBinaryCmpExprToCmpExprAction,
	{_State66, _WildcardMarker}:                          _ReduceBinaryMulExprToMulExprAction,
	{_State67, _EndMarker}:                               _ReduceBinaryOpAssignStatementToSimpleStatementAction,
	{_State68, _WildcardMarker}:                          _ReduceBinaryOrExprToOrExprAction,
	{_State69, _WildcardMarker}:                          _ReduceBlockExprToAtomExprAction,
	{_State70, _WildcardMarker}:                          _ReduceCallExprToAccessibleExprAction,
	{_State71, LbraceToken}:                              _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State72, _EndMarker}:                               _ReduceCallbackOpStatementToSimpleStatementAction,
	{_State73, _WildcardMarker}:                          _ReduceCmpExprToAndExprAction,
	{_State74, _WildcardMarker}:                          _ReduceExplicitStructDefToInitializableTypeAction,
	{_State75, _WildcardMarker}:                          _ReduceExpressionToExpressionsAction,
	{_State76, _EndMarker}:                               _ReduceExpressionOrImproperStructStatementToSimpleStatementAction,
	{_State77, _WildcardMarker}:                          _ReduceToExpressionOrImproperStructStatementAction,
	{_State78, _EndMarker}:                               _ReduceFallthroughStatementToSimpleStatementAction,
	{_State79, _WildcardMarker}:                          _ReduceIdentifierExprToAtomExprAction,
	{_State80, _WildcardMarker}:                          _ReduceImplicitStructExprToAtomExprAction,
	{_State81, _EndMarker}:                               _ReduceImportStatementToStatementAction,
	{_State82, _WildcardMarker}:                          _ReduceIndexExprToAccessibleExprAction,
	{_State84, _WildcardMarker}:                          _ReduceInitializeExprToAtomExprAction,
	{_State85, _EndMarker}:                               _ReduceJumpStatementToSimpleStatementAction,
	{_State86, _EndMarker}:                               _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State86, NewlinesToken}:                            _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State86, RbraceToken}:                              _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State86, SemicolonToken}:                           _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State86, _WildcardMarker}:                          _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State87, _WildcardMarker}:                          _ReduceLiteralExprToAtomExprAction,
	{_State88, _WildcardMarker}:                          _ReduceMapTypeToInitializableTypeAction,
	{_State89, _WildcardMarker}:                          _ReduceMulExprToAddExprAction,
	{_State91, _WildcardMarker}:                          _ReduceOrExprToSequenceExprAction,
	{_State92, _WildcardMarker}:                          _ReduceParseErrorExprToAtomExprAction,
	{_State93, _WildcardMarker}:                          _ReducePostfixUnaryExprToPostfixableExprAction,
	{_State94, _WildcardMarker}:                          _ReducePostfixableExprToPrefixableExprAction,
	{_State95, _WildcardMarker}:                          _ReducePrefixUnaryExprToPrefixableExprAction,
	{_State96, LbraceToken}:                              _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State97, _WildcardMarker}:                          _ReducePrefixableExprToMulExprAction,
	{_State98, AssignToken}:                              _ReduceToAssignPatternAction,
	{_State98, _WildcardMarker}:                          _ReduceSequenceExprToExpressionAction,
	{_State99, _EndMarker}:                               _ReduceSimpleStatementToStatementAction,
	{_State100, _WildcardMarker}:                         _ReduceSliceTypeToInitializableTypeAction,
	{_State101, _EndMarker}:                              _ReduceUnaryOpAssignStatementToSimpleStatementAction,
	{_State102, _EndMarker}:                              _ReduceUnsafeStatementToSimpleStatementAction,
	{_State103, _EndMarker}:                              _ReduceVarDeclPatternToSequenceExprAction,
	{_State104, _WildcardMarker}:                         _ReduceAccessibleExprToPostfixableExprAction,
	{_State104, LparenToken}:                             _ReduceNilToOptionalGenericBindingAction,
	{_State105, _EndMarker}:                              _ReduceSequenceExprToExpressionAction,
	{_State106, _WildcardMarker}:                         _ReduceBitAndToPrefixTypeOpAction,
	{_State107, _WildcardMarker}:                         _ReduceBitNegToPrefixTypeOpAction,
	{_State108, _WildcardMarker}:                         _ReduceNilToOptionalGenericBindingAction,
	{_State110, _WildcardMarker}:                         _ReduceExclaimToPrefixTypeOpAction,
	{_State112, _WildcardMarker}:                         _ReduceNilToOptionalGenericBindingAction,
	{_State113, RparenToken}:                             _ReduceNilToImplicitFieldDefsAction,
	{_State114, _WildcardMarker}:                         _ReduceToParseErrorTypeAction,
	{_State115, _WildcardMarker}:                         _ReduceQuestionToPrefixTypeOpAction,
	{_State116, _WildcardMarker}:                         _ReduceTildeTildeToPrefixTypeOpAction,
	{_State118, _WildcardMarker}:                         _ReduceAtomTypeToReturnableTypeAction,
	{_State119, _WildcardMarker}:                         _ReduceExplicitEnumDefToAtomTypeAction,
	{_State120, _WildcardMarker}:                         _ReduceFuncTypeToAtomTypeAction,
	{_State121, _WildcardMarker}:                         _ReduceImplicitEnumDefToAtomTypeAction,
	{_State122, _WildcardMarker}:                         _ReduceImplicitStructDefToAtomTypeAction,
	{_State123, _WildcardMarker}:                         _ReduceInitializableTypeToAtomTypeAction,
	{_State124, _WildcardMarker}:                         _ReduceParseErrorTypeToAtomTypeAction,
	{_State126, _WildcardMarker}:                         _ReducePrefixedTypeToReturnableTypeAction,
	{_State127, _WildcardMarker}:                         _ReduceReturnableTypeToValueTypeAction,
	{_State128, _WildcardMarker}:                         _ReduceTraitDefToAtomTypeAction,
	{_State129, _WildcardMarker}:                         _ReduceTraitOpTypeToValueTypeAction,
	{_State130, LparenToken}:                             _ReduceNilToOptionalGenericParametersAction,
	{_State132, RbraceToken}:                             _ReduceProperStatementsToStatementsAction,
	{_State133, _WildcardMarker}:                         _ReduceStatementToProperStatementsAction,
	{_State135, _WildcardMarker}:                         _ReduceWithSpecToPackageDefAction,
	{_State136, _WildcardMarker}:                         _ReduceNilToOptionalGenericParametersAction,
	{_State137, _EndMarker}:                              _ReduceImproperToDefinitionsAction,
	{_State138, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State139, _WildcardMarker}:                         _ReduceIdentifierToVarPatternAction,
	{_State141, _WildcardMarker}:                         _ReduceTuplePatternToVarPatternAction,
	{_State142, _WildcardMarker}:                         _ReduceNilToOptionalValueTypeAction,
	{_State144, _WildcardMarker}:                         _ReduceVarToVarOrLetAction,
	{_State145, _WildcardMarker}:                         _ReduceAssignPatternToCasePatternAction,
	{_State146, _WildcardMarker}:                         _ReduceCasePatternToCasePatternsAction,
	{_State149, _WildcardMarker}:                         _ReduceToAssignPatternAction,
	{_State150, _EndMarker}:                              _ReduceNilToOptionalSimpleStatementAction,
	{_State150, NewlinesToken}:                           _ReduceNilToOptionalSimpleStatementAction,
	{_State150, RbraceToken}:                             _ReduceNilToOptionalSimpleStatementAction,
	{_State150, SemicolonToken}:                          _ReduceNilToOptionalSimpleStatementAction,
	{_State150, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State151, RparenToken}:                             _ReduceNilToParameterDefsAction,
	{_State152, _EndMarker}:                              _ReduceAssignVarPatternToSequenceExprAction,
	{_State154, _WildcardMarker}:                         _ReduceStringLiteralToImportClauseAction,
	{_State155, _EndMarker}:                              _ReduceSingleToImportStatementAction,
	{_State157, ColonToken}:                              _ReduceUnitUnitPairToColonExprAction,
	{_State157, CommaToken}:                              _ReduceUnitUnitPairToColonExprAction,
	{_State157, RbracketToken}:                           _ReduceUnitUnitPairToColonExprAction,
	{_State157, RparenToken}:                             _ReduceUnitUnitPairToColonExprAction,
	{_State157, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State158, _WildcardMarker}:                         _ReduceSkipPatternToArgumentAction,
	{_State159, _WildcardMarker}:                         _ReduceToIdentifierExprAction,
	{_State160, _WildcardMarker}:                         _ReduceArgumentToProperArgumentsAction,
	{_State162, _WildcardMarker}:                         _ReduceColonExprToArgumentAction,
	{_State163, _WildcardMarker}:                         _ReducePositionalToArgumentAction,
	{_State164, RparenToken}:                             _ReduceProperArgumentsToArgumentsAction,
	{_State165, RparenToken}:                             _ReduceNilToExplicitFieldDefsAction,
	{_State167, _WildcardMarker}:                         _ReduceAddAssignToBinaryOpAssignAction,
	{_State168, _EndMarker}:                              _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State169, _WildcardMarker}:                         _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State170, _WildcardMarker}:                         _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State171, _WildcardMarker}:                         _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State172, _WildcardMarker}:                         _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State173, _WildcardMarker}:                         _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State174, _WildcardMarker}:                         _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State175, _WildcardMarker}:                         _ReduceDivAssignToBinaryOpAssignAction,
	{_State176, RbracketToken}:                           _ReduceNilToGenericArgumentsAction,
	{_State178, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State179, _WildcardMarker}:                         _ReduceModAssignToBinaryOpAssignAction,
	{_State180, _WildcardMarker}:                         _ReduceMulAssignToBinaryOpAssignAction,
	{_State181, _WildcardMarker}:                         _ReduceToPostfixUnaryExprAction,
	{_State182, _WildcardMarker}:                         _ReduceSubAssignToBinaryOpAssignAction,
	{_State183, _EndMarker}:                              _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State184, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State186, _EndMarker}:                              _ReduceToUnaryOpAssignStatementAction,
	{_State187, _WildcardMarker}:                         _ReduceAddToAddOpAction,
	{_State188, _WildcardMarker}:                         _ReduceBitOrToAddOpAction,
	{_State189, _WildcardMarker}:                         _ReduceBitXorToAddOpAction,
	{_State190, _WildcardMarker}:                         _ReduceSubToAddOpAction,
	{_State191, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State192, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State193, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State194, LparenToken}:                             _ReduceNilToOptionalGenericBindingAction,
	{_State195, _EndMarker}:                              _ReduceToCallbackOpStatementAction,
	{_State195, NewlinesToken}:                           _ReduceToCallbackOpStatementAction,
	{_State195, RbraceToken}:                             _ReduceToCallbackOpStatementAction,
	{_State195, SemicolonToken}:                          _ReduceToCallbackOpStatementAction,
	{_State195, _WildcardMarker}:                         _ReduceCallExprToAccessibleExprAction,
	{_State196, _WildcardMarker}:                         _ReduceEqualToCmpOpAction,
	{_State197, _WildcardMarker}:                         _ReduceGreaterToCmpOpAction,
	{_State198, _WildcardMarker}:                         _ReduceGreaterOrEqualToCmpOpAction,
	{_State199, _WildcardMarker}:                         _ReduceLessToCmpOpAction,
	{_State200, _WildcardMarker}:                         _ReduceLessOrEqualToCmpOpAction,
	{_State201, _WildcardMarker}:                         _ReduceNotEqualToCmpOpAction,
	{_State202, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State203, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State204, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State204, RparenToken}:                             _ReduceNilToArgumentsAction,
	{_State205, _EndMarker}:                              _ReduceLabeledNoValueToJumpStatementAction,
	{_State205, NewlinesToken}:                           _ReduceLabeledNoValueToJumpStatementAction,
	{_State205, RbraceToken}:                             _ReduceLabeledNoValueToJumpStatementAction,
	{_State205, SemicolonToken}:                          _ReduceLabeledNoValueToJumpStatementAction,
	{_State205, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State206, _EndMarker}:                              _ReduceUnlabeledValuedToJumpStatementAction,
	{_State207, _WildcardMarker}:                         _ReduceBitAndToMulOpAction,
	{_State208, _WildcardMarker}:                         _ReduceBitLshiftToMulOpAction,
	{_State209, _WildcardMarker}:                         _ReduceBitRshiftToMulOpAction,
	{_State210, _WildcardMarker}:                         _ReduceDivToMulOpAction,
	{_State211, _WildcardMarker}:                         _ReduceModToMulOpAction,
	{_State212, _WildcardMarker}:                         _ReduceMulToMulOpAction,
	{_State213, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State215, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State216, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State217, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State218, _EndMarker}:                              _ReduceIfExprToExpressionAction,
	{_State219, _EndMarker}:                              _ReduceLoopExprToExpressionAction,
	{_State220, _WildcardMarker}:                         _ReduceToBlockExprAction,
	{_State221, _EndMarker}:                              _ReduceSwitchExprToExpressionAction,
	{_State222, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State223, _WildcardMarker}:                         _ReduceToPrefixUnaryExprAction,
	{_State224, _WildcardMarker}:                         _ReduceInferredToAtomTypeAction,
	{_State226, RparenToken}:                             _ReduceNilToParameterDeclsAction,
	{_State228, _WildcardMarker}:                         _ReduceNamedToAtomTypeAction,
	{_State229, _WildcardMarker}:                         _ReduceNilToOptionalGenericBindingAction,
	{_State231, _WildcardMarker}:                         _ReduceFieldDefToProperImplicitFieldDefsAction,
	{_State231, OrToken}:                                 _ReduceFieldDefToEnumValueDefAction,
	{_State234, RparenToken}:                             _ReduceProperImplicitFieldDefsToImplicitFieldDefsAction,
	{_State235, _WildcardMarker}:                         _ReduceUnsafeStatementToFieldDefAction,
	{_State236, _WildcardMarker}:                         _ReduceImplicitToFieldDefAction,
	{_State237, RparenToken}:                             _ReduceNilToTraitPropertiesAction,
	{_State238, _WildcardMarker}:                         _ReduceToPrefixedTypeAction,
	{_State239, _WildcardMarker}:                         _ReduceAddToTraitOpAction,
	{_State240, _WildcardMarker}:                         _ReduceMulToTraitOpAction,
	{_State241, _WildcardMarker}:                         _ReduceSubToTraitOpAction,
	{_State243, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State244, RbracketToken}:                           _ReduceNilToGenericParameterDefsAction,
	{_State246, _WildcardMarker}:                         _ReduceInferredRefArgToParameterDefAction,
	{_State248, RbraceToken}:                             _ReduceImproperImplicitToStatementsAction,
	{_State248, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State249, RbraceToken}:                             _ReduceImproperExplicitToStatementsAction,
	{_State249, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State250, _EndMarker}:                              _ReduceToStatementBlockAction,
	{_State253, _WildcardMarker}:                         _ReduceAddToProperDefinitionsAction,
	{_State254, _WildcardMarker}:                         _ReduceGlobalVarAssignmentToDefinitionAction,
	{_State255, _WildcardMarker}:                         _ReduceEllipsisToFieldVarPatternAction,
	{_State256, _WildcardMarker}:                         _ReduceIdentifierToVarPatternAction,
	{_State257, _WildcardMarker}:                         _ReduceFieldVarPatternToFieldVarPatternsAction,
	{_State259, _WildcardMarker}:                         _ReducePositionalToFieldVarPatternAction,
	{_State260, _EndMarker}:                              _ReduceToVarDeclPatternAction,
	{_State261, _WildcardMarker}:                         _ReduceValueTypeToOptionalValueTypeAction,
	{_State264, _EndMarker}:                              _ReduceNilToOptionalSimpleStatementAction,
	{_State264, NewlinesToken}:                           _ReduceNilToOptionalSimpleStatementAction,
	{_State264, RbraceToken}:                             _ReduceNilToOptionalSimpleStatementAction,
	{_State264, SemicolonToken}:                          _ReduceNilToOptionalSimpleStatementAction,
	{_State264, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State265, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State266, _EndMarker}:                              _ReduceDefaultBranchStatementToStatementAction,
	{_State267, _EndMarker}:                              _ReduceSimpleStatementToOptionalSimpleStatementAction,
	{_State268, _WildcardMarker}:                         _ReduceParameterDefToProperParameterDefsAction,
	{_State270, RparenToken}:                             _ReduceProperParameterDefsToParameterDefsAction,
	{_State271, _WildcardMarker}:                         _ReduceImportClauseToProperImportClausesAction,
	{_State273, RparenToken}:                             _ReduceProperImportClausesToImportClausesAction,
	{_State277, _WildcardMarker}:                         _ReduceToSliceTypeAction,
	{_State278, _WildcardMarker}:                         _ReduceUnitExprPairToColonExprAction,
	{_State279, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State280, _WildcardMarker}:                         _ReduceToImplicitStructExprAction,
	{_State281, ColonToken}:                              _ReduceColonExprUnitTupleToColonExprAction,
	{_State281, CommaToken}:                              _ReduceColonExprUnitTupleToColonExprAction,
	{_State281, RbracketToken}:                           _ReduceColonExprUnitTupleToColonExprAction,
	{_State281, RparenToken}:                             _ReduceColonExprUnitTupleToColonExprAction,
	{_State281, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State282, ColonToken}:                              _ReduceExprUnitPairToColonExprAction,
	{_State282, CommaToken}:                              _ReduceExprUnitPairToColonExprAction,
	{_State282, RbracketToken}:                           _ReduceExprUnitPairToColonExprAction,
	{_State282, RparenToken}:                             _ReduceExprUnitPairToColonExprAction,
	{_State282, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State283, _WildcardMarker}:                         _ReduceVarargAssignmentToArgumentAction,
	{_State284, RparenToken}:                             _ReduceImproperToArgumentsAction,
	{_State284, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State286, _WildcardMarker}:                         _ReduceFieldDefToProperExplicitFieldDefsAction,
	{_State287, RparenToken}:                             _ReduceProperExplicitFieldDefsToExplicitFieldDefsAction,
	{_State290, RbracketToken}:                           _ReduceProperGenericArgumentsToGenericArgumentsAction,
	{_State291, _WildcardMarker}:                         _ReduceValueTypeToProperGenericArgumentsAction,
	{_State292, _WildcardMarker}:                         _ReduceToAccessExprAction,
	{_State294, _EndMarker}:                              _ReduceToBinaryOpAssignStatementAction,
	{_State295, _WildcardMarker}:                         _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State295, RparenToken}:                             _ReduceNilToArgumentsAction,
	{_State296, _WildcardMarker}:                         _ReduceToBinaryAddExprAction,
	{_State297, _WildcardMarker}:                         _ReduceToBinaryAndExprAction,
	{_State298, _EndMarker}:                              _ReduceToAssignStatementAction,
	{_State299, _WildcardMarker}:                         _ReduceToBinaryCmpExprAction,
	{_State300, _WildcardMarker}:                         _ReduceAddToExpressionsAction,
	{_State302, _EndMarker}:                              _ReduceLabeledValuedToJumpStatementAction,
	{_State303, _WildcardMarker}:                         _ReduceToBinaryMulExprAction,
	{_State304, _WildcardMarker}:                         _ReduceInfiniteToLoopExprAction,
	{_State307, _WildcardMarker}:                         _ReduceToAssignPatternAction,
	{_State307, SemicolonToken}:                          _ReduceSequenceExprToForAssignmentAction,
	{_State308, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State310, LbraceToken}:                             _ReduceSequenceExprToConditionAction,
	{_State312, _WildcardMarker}:                         _ReduceToBinaryOrExprAction,
	{_State315, _WildcardMarker}:                         _ReduceFieldDefToEnumValueDefAction,
	{_State318, _WildcardMarker}:                         _ReduceNilToOptionalGenericBindingAction,
	{_State319, _WildcardMarker}:                         _ReduceParameterDeclToProperParameterDeclsAction,
	{_State321, RparenToken}:                             _ReduceProperParameterDeclsToParameterDeclsAction,
	{_State322, _WildcardMarker}:                         _ReduceUnamedToParameterDeclAction,
	{_State323, _WildcardMarker}:                         _ReduceNilToOptionalGenericBindingAction,
	{_State324, _WildcardMarker}:                         _ReduceNilToOptionalGenericBindingAction,
	{_State325, _WildcardMarker}:                         _ReduceExplicitToFieldDefAction,
	{_State329, _WildcardMarker}:                         _ReduceToImplicitEnumDefAction,
	{_State330, _WildcardMarker}:                         _ReduceToImplicitStructDefAction,
	{_State331, RparenToken}:                             _ReduceImproperToImplicitFieldDefsAction,
	{_State333, _WildcardMarker}:                         _ReduceFieldDefToTraitPropertyAction,
	{_State334, _WildcardMarker}:                         _ReduceMethodSignatureToTraitPropertyAction,
	{_State335, RparenToken}:                             _ReduceProperTraitPropertiesToTraitPropertiesAction,
	{_State337, _WildcardMarker}:                         _ReduceTraitPropertyToProperTraitPropertiesAction,
	{_State338, _WildcardMarker}:                         _ReduceToTraitOpTypeAction,
	{_State339, _WildcardMarker}:                         _ReduceAliasToNamedFuncDefAction,
	{_State340, _WildcardMarker}:                         _ReduceUnconstrainedToGenericParameterDefAction,
	{_State341, _WildcardMarker}:                         _ReduceGenericParameterDefToProperGenericParameterDefsAction,
	{_State343, RbracketToken}:                           _ReduceProperGenericParameterDefsToGenericParameterDefsAction,
	{_State344, RparenToken}:                             _ReduceNilToParameterDefsAction,
	{_State345, _WildcardMarker}:                         _ReduceInferredRefVarargToParameterDefAction,
	{_State346, _WildcardMarker}:                         _ReduceArgToParameterDefAction,
	{_State348, _WildcardMarker}:                         _ReduceAddImplicitToProperStatementsAction,
	{_State349, _WildcardMarker}:                         _ReduceAddExplicitToProperStatementsAction,
	{_State350, _WildcardMarker}:                         _ReduceAliasToTypeDefAction,
	{_State351, _WildcardMarker}:                         _ReduceDefinitionToTypeDefAction,
	{_State354, _WildcardMarker}:                         _ReduceToTuplePatternAction,
	{_State355, _WildcardMarker}:                         _ReduceEnumMatchPatternToCasePatternAction,
	{_State357, _EndMarker}:                              _ReduceCaseBranchStatementToStatementAction,
	{_State358, _WildcardMarker}:                         _ReduceMultipleToCasePatternsAction,
	{_State359, LbraceToken}:                             _ReduceNilToReturnTypeAction,
	{_State360, RparenToken}:                             _ReduceImproperToParameterDefsAction,
	{_State361, _EndMarker}:                              _ReduceMultipleToImportStatementAction,
	{_State362, RparenToken}:                             _ReduceExplicitToImportClausesAction,
	{_State363, RparenToken}:                             _ReduceImplicitToImportClausesAction,
	{_State364, _EndMarker}:                              _ReduceAliasToImportClauseAction,
	{_State367, _WildcardMarker}:                         _ReduceNamedAssignmentToArgumentAction,
	{_State368, _WildcardMarker}:                         _ReduceColonExprExprTupleToColonExprAction,
	{_State369, _WildcardMarker}:                         _ReduceExprExprPairToColonExprAction,
	{_State370, _WildcardMarker}:                         _ReduceAddToProperArgumentsAction,
	{_State371, _WildcardMarker}:                         _ReduceToExplicitStructDefAction,
	{_State372, RparenToken}:                             _ReduceImproperExplicitToExplicitFieldDefsAction,
	{_State373, RparenToken}:                             _ReduceImproperImplicitToExplicitFieldDefsAction,
	{_State375, _WildcardMarker}:                         _ReduceBindingToOptionalGenericBindingAction,
	{_State376, RbracketToken}:                           _ReduceImproperToGenericArgumentsAction,
	{_State377, _WildcardMarker}:                         _ReduceToIndexExprAction,
	{_State379, _WildcardMarker}:                         _ReduceToInitializeExprAction,
	{_State380, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State381, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State382, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State383, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State383, SemicolonToken}:                          _ReduceNilToOptionalSequenceExprAction,
	{_State386, _WildcardMarker}:                         _ReduceNoElseToIfExprAction,
	{_State387, _EndMarker}:                              _ReduceToSwitchExprAction,
	{_State390, _WildcardMarker}:                         _ReduceToExplicitEnumDefAction,
	{_State393, _WildcardMarker}:                         _ReduceUnnamedVarargToParameterDeclAction,
	{_State395, _WildcardMarker}:                         _ReduceArgToParameterDeclAction,
	{_State396, _WildcardMarker}:                         _ReduceNilToReturnTypeAction,
	{_State397, RparenToken}:                             _ReduceImproperToParameterDeclsAction,
	{_State398, _WildcardMarker}:                         _ReduceExternNamedToAtomTypeAction,
	{_State399, _WildcardMarker}:                         _ReducePairToImplicitEnumValueDefsAction,
	{_State400, _WildcardMarker}:                         _ReduceDefaultToEnumValueDefAction,
	{_State401, _WildcardMarker}:                         _ReduceAddToImplicitEnumValueDefsAction,
	{_State402, _WildcardMarker}:                         _ReduceAddToProperImplicitFieldDefsAction,
	{_State404, RparenToken}:                             _ReduceImproperExplicitToTraitPropertiesAction,
	{_State405, RparenToken}:                             _ReduceImproperImplicitToTraitPropertiesAction,
	{_State406, _WildcardMarker}:                         _ReduceToTraitDefAction,
	{_State407, _WildcardMarker}:                         _ReduceConstrainedToGenericParameterDefAction,
	{_State408, _WildcardMarker}:                         _ReduceGenericToOptionalGenericParametersAction,
	{_State409, RbracketToken}:                           _ReduceImproperToGenericParameterDefsAction,
	{_State411, _WildcardMarker}:                         _ReduceVarargToParameterDefAction,
	{_State412, LparenToken}:                             _ReduceNilToOptionalGenericParametersAction,
	{_State414, _WildcardMarker}:                         _ReduceNamedToFieldVarPatternAction,
	{_State415, _WildcardMarker}:                         _ReduceAddToFieldVarPatternsAction,
	{_State416, _WildcardMarker}:                         _ReduceEnumVarDeclPatternToCasePatternAction,
	{_State418, _WildcardMarker}:                         _ReduceReturnableTypeToReturnTypeAction,
	{_State419, _WildcardMarker}:                         _ReduceAddToProperParameterDefsAction,
	{_State420, _WildcardMarker}:                         _ReduceAddExplicitToProperImportClausesAction,
	{_State421, _WildcardMarker}:                         _ReduceAddImplicitToProperImportClausesAction,
	{_State422, _WildcardMarker}:                         _ReduceToMapTypeAction,
	{_State423, _WildcardMarker}:                         _ReduceToArrayTypeAction,
	{_State424, _WildcardMarker}:                         _ReduceAddExplicitToProperExplicitFieldDefsAction,
	{_State425, _WildcardMarker}:                         _ReduceAddImplicitToProperExplicitFieldDefsAction,
	{_State426, _EndMarker}:                              _ReduceToUnsafeStatementAction,
	{_State427, _WildcardMarker}:                         _ReduceAddToProperGenericArgumentsAction,
	{_State428, _WildcardMarker}:                         _ReduceToCallExprAction,
	{_State429, _EndMarker}:                              _ReduceDoWhileToLoopExprAction,
	{_State430, SemicolonToken}:                          _ReduceAssignToForAssignmentAction,
	{_State433, DoToken}:                                 _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State434, _EndMarker}:                              _ReduceWhileToLoopExprAction,
	{_State435, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State437, RparenToken}:                             _ReduceImplicitPairToExplicitEnumValueDefsAction,
	{_State438, _WildcardMarker}:                         _ReducePairToImplicitEnumValueDefsAction,
	{_State438, RparenToken}:                             _ReduceExplicitPairToExplicitEnumValueDefsAction,
	{_State439, RparenToken}:                             _ReduceImplicitAddToExplicitEnumValueDefsAction,
	{_State440, _WildcardMarker}:                         _ReduceAddToImplicitEnumValueDefsAction,
	{_State440, RparenToken}:                             _ReduceExplicitAddToExplicitEnumValueDefsAction,
	{_State441, _WildcardMarker}:                         _ReduceVarargToParameterDeclAction,
	{_State442, _WildcardMarker}:                         _ReduceToFuncTypeAction,
	{_State443, _WildcardMarker}:                         _ReduceAddToProperParameterDeclsAction,
	{_State444, RparenToken}:                             _ReduceNilToParameterDeclsAction,
	{_State445, _WildcardMarker}:                         _ReduceExplicitToProperTraitPropertiesAction,
	{_State446, _WildcardMarker}:                         _ReduceImplicitToProperTraitPropertiesAction,
	{_State447, _WildcardMarker}:                         _ReduceAddToProperGenericParameterDefsAction,
	{_State448, LbraceToken}:                             _ReduceNilToReturnTypeAction,
	{_State450, _WildcardMarker}:                         _ReduceConstrainedDefToTypeDefAction,
	{_State451, _WildcardMarker}:                         _ReduceToAnonymousFuncExprAction,
	{_State453, LbraceToken}:                             _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State453, DoToken}:                                 _ReduceNilToOptionalSequenceExprAction,
	{_State454, LbraceToken}:                             _ReduceCaseToConditionAction,
	{_State455, _EndMarker}:                              _ReduceMultiIfElseToIfExprAction,
	{_State456, _EndMarker}:                              _ReduceIfElseToIfExprAction,
	{_State459, RparenToken}:                             _ReduceNilToParameterDefsAction,
	{_State460, _EndMarker}:                              _ReduceIteratorToLoopExprAction,
	{_State462, _WildcardMarker}:                         _ReduceNilToReturnTypeAction,
	{_State463, _WildcardMarker}:                         _ReduceFuncDefToNamedFuncDefAction,
	{_State466, _WildcardMarker}:                         _ReduceToMethodSignatureAction,
	{_State467, LbraceToken}:                             _ReduceNilToReturnTypeAction,
	{_State468, _EndMarker}:                              _ReduceForToLoopExprAction,
	{_State470, _WildcardMarker}:                         _ReduceMethodDefToNamedFuncDefAction,
}

/*
Parser Debug States:
  State 1:
    Kernel Items:
      #accept: ^.source
    Reduce:
      $ -> [definitions]
    Goto:
      COMMENT_GROUPS -> State 9
      PACKAGE -> State 13
      TYPE -> State 14
      FUNC -> State 10
      VAR -> State 15
      LET -> State 12
      LBRACE -> State 11
      source -> State 5
      proper_definitions -> State 20
      definitions -> State 17
      definition -> State 16
      statement_block -> State 21
      var_decl_pattern -> State 23
      var_or_let -> State 24
      type_def -> State 22
      named_func_def -> State 18
      package_def -> State 19

  State 2:
    Kernel Items:
      #accept: ^.statement
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      CASE -> State 29
      DEFAULT -> State 31
      RETURN -> State 47
      BREAK -> State 28
      CONTINUE -> State 30
      FALLTHROUGH -> State 33
      IMPORT -> State 39
      UNSAFE -> State 53
      STRUCT -> State 50
      FUNC -> State 36
      ASYNC -> State 25
      DEFER -> State 32
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      simple_statement -> State 99
      statement -> State 6
      expression_or_improper_struct_statement -> State 76
      expressions -> State 77
      callback_op -> State 71
      callback_op_statement -> State 72
      unsafe_statement -> State 102
      jump_statement -> State 85
      jump_type -> State 86
      fallthrough_statement -> State 78
      assign_statement -> State 61
      unary_op_assign_statement -> State 101
      binary_op_assign_statement -> State 67
      import_statement -> State 81
      var_decl_pattern -> State 103
      var_or_let -> State 24
      assign_pattern -> State 60
      expression -> State 75
      optional_label_decl -> State 90
      sequence_expr -> State 98
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 55
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 3:
    Kernel Items:
      #accept: ^.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 7
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 4:
    Kernel Items:
      #accept: ^.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 8
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 5:
    Kernel Items:
      #accept: ^ source., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 6:
    Kernel Items:
      #accept: ^ statement., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 7:
    Kernel Items:
      #accept: ^ expression., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      #accept: ^ value_type., $
      trait_op_type: value_type.trait_op returnable_type
    Reduce:
      $ -> [#accept]
    Goto:
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 9:
    Kernel Items:
      definition: COMMENT_GROUPS., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      named_func_def: FUNC.IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.IDENTIFIER ASSIGN expression
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 130
      LPAREN -> State 131

  State 11:
    Kernel Items:
      statement_block: LBRACE.statements RBRACE
    Reduce:
      * -> [optional_label_decl]
      RBRACE -> [statements]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      CASE -> State 29
      DEFAULT -> State 31
      RETURN -> State 47
      BREAK -> State 28
      CONTINUE -> State 30
      FALLTHROUGH -> State 33
      IMPORT -> State 39
      UNSAFE -> State 53
      STRUCT -> State 50
      FUNC -> State 36
      ASYNC -> State 25
      DEFER -> State 32
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      proper_statements -> State 132
      statements -> State 134
      simple_statement -> State 99
      statement -> State 133
      expression_or_improper_struct_statement -> State 76
      expressions -> State 77
      callback_op -> State 71
      callback_op_statement -> State 72
      unsafe_statement -> State 102
      jump_statement -> State 85
      jump_type -> State 86
      fallthrough_statement -> State 78
      assign_statement -> State 61
      unary_op_assign_statement -> State 101
      binary_op_assign_statement -> State 67
      import_statement -> State 81
      var_decl_pattern -> State 103
      var_or_let -> State 24
      assign_pattern -> State 60
      expression -> State 75
      optional_label_decl -> State 90
      sequence_expr -> State 98
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 55
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 12:
    Kernel Items:
      var_or_let: LET., *
    Reduce:
      * -> [var_or_let]
    Goto:
      (nil)

  State 13:
    Kernel Items:
      package_def: PACKAGE., *
      package_def: PACKAGE.statement_block
    Reduce:
      * -> [package_def]
    Goto:
      LBRACE -> State 11
      statement_block -> State 135

  State 14:
    Kernel Items:
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type
      type_def: TYPE.IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE.IDENTIFIER ASSIGN value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 136

  State 15:
    Kernel Items:
      var_or_let: VAR., *
    Reduce:
      * -> [var_or_let]
    Goto:
      (nil)

  State 16:
    Kernel Items:
      proper_definitions: definition., *
    Reduce:
      * -> [proper_definitions]
    Goto:
      (nil)

  State 17:
    Kernel Items:
      source: definitions., $
    Reduce:
      $ -> [source]
    Goto:
      (nil)

  State 18:
    Kernel Items:
      definition: named_func_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 19:
    Kernel Items:
      definition: package_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 20:
    Kernel Items:
      proper_definitions: proper_definitions.NEWLINES definition
      definitions: proper_definitions., $
      definitions: proper_definitions.NEWLINES
    Reduce:
      $ -> [definitions]
    Goto:
      NEWLINES -> State 137

  State 21:
    Kernel Items:
      definition: statement_block., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 22:
    Kernel Items:
      definition: type_def., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 23:
    Kernel Items:
      definition: var_decl_pattern., *
      definition: var_decl_pattern.ASSIGN expression
    Reduce:
      * -> [definition]
    Goto:
      ASSIGN -> State 138

  State 24:
    Kernel Items:
      var_decl_pattern: var_or_let.var_pattern optional_value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      LPAREN -> State 140
      var_pattern -> State 142
      tuple_pattern -> State 141

  State 25:
    Kernel Items:
      callback_op: ASYNC., *
    Reduce:
      * -> [callback_op]
    Goto:
      (nil)

  State 26:
    Kernel Items:
      prefix_unary_op: BIT_AND., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 27:
    Kernel Items:
      prefix_unary_op: BIT_NEG., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 28:
    Kernel Items:
      jump_type: BREAK., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 29:
    Kernel Items:
      statement: CASE.case_patterns COLON optional_simple_statement
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 144
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      DOT -> State 143
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      case_patterns -> State 147
      var_decl_pattern -> State 103
      var_or_let -> State 24
      assign_pattern -> State 145
      case_pattern -> State 146
      optional_label_decl -> State 148
      sequence_expr -> State 149
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 30:
    Kernel Items:
      jump_type: CONTINUE., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 31:
    Kernel Items:
      statement: DEFAULT.COLON optional_simple_statement
    Reduce:
      (nil)
    Goto:
      COLON -> State 150

  State 32:
    Kernel Items:
      callback_op: DEFER., *
    Reduce:
      * -> [callback_op]
    Goto:
      (nil)

  State 33:
    Kernel Items:
      fallthrough_statement: FALLTHROUGH., $
    Reduce:
      $ -> [fallthrough_statement]
    Goto:
      (nil)

  State 34:
    Kernel Items:
      literal_expr: FALSE., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 35:
    Kernel Items:
      literal_expr: FLOAT_LITERAL., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 36:
    Kernel Items:
      anonymous_func_expr: FUNC.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 151

  State 37:
    Kernel Items:
      sequence_expr: GREATER.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      optional_label_decl -> State 148
      sequence_expr -> State 152
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 38:
    Kernel Items:
      identifier_expr: IDENTIFIER., *
    Reduce:
      * -> [identifier_expr]
    Goto:
      (nil)

  State 39:
    Kernel Items:
      import_statement: IMPORT.import_clause
      import_statement: IMPORT.LPAREN import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 154
      LPAREN -> State 153
      import_clause -> State 155

  State 40:
    Kernel Items:
      literal_expr: INTEGER_LITERAL., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 41:
    Kernel Items:
      optional_label_decl: LABEL_DECL., *
    Reduce:
      * -> [optional_label_decl]
    Goto:
      (nil)

  State 42:
    Kernel Items:
      slice_type: LBRACKET.value_type RBRACKET
      array_type: LBRACKET.value_type COMMA INTEGER_LITERAL RBRACKET
      map_type: LBRACKET.value_type COLON value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 156
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 43:
    Kernel Items:
      implicit_struct_expr: LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [arguments]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 159
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      COLON -> State 157
      ELLIPSIS -> State 158
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 163
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      proper_arguments -> State 164
      arguments -> State 161
      argument -> State 160
      colon_expr -> State 162
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 44:
    Kernel Items:
      prefix_unary_op: MUL., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 45:
    Kernel Items:
      prefix_unary_op: NOT., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 46:
    Kernel Items:
      parse_error_expr: PARSE_ERROR., *
    Reduce:
      * -> [parse_error_expr]
    Goto:
      (nil)

  State 47:
    Kernel Items:
      jump_type: RETURN., *
    Reduce:
      * -> [jump_type]
    Goto:
      (nil)

  State 48:
    Kernel Items:
      literal_expr: RUNE_LITERAL., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 49:
    Kernel Items:
      literal_expr: STRING_LITERAL., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 50:
    Kernel Items:
      explicit_struct_def: STRUCT.LPAREN explicit_field_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 165

  State 51:
    Kernel Items:
      prefix_unary_op: SUB., *
    Reduce:
      * -> [prefix_unary_op]
    Goto:
      (nil)

  State 52:
    Kernel Items:
      literal_expr: TRUE., *
    Reduce:
      * -> [literal_expr]
    Goto:
      (nil)

  State 53:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 166

  State 54:
    Kernel Items:
      accessible_expr: access_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 55:
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
      LBRACKET -> State 178
      DOT -> State 177
      QUESTION -> State 181
      DOLLAR_LBRACKET -> State 176
      ADD_ASSIGN -> State 167
      SUB_ASSIGN -> State 182
      MUL_ASSIGN -> State 180
      DIV_ASSIGN -> State 175
      MOD_ASSIGN -> State 179
      ADD_ONE_ASSIGN -> State 168
      SUB_ONE_ASSIGN -> State 183
      BIT_NEG_ASSIGN -> State 171
      BIT_AND_ASSIGN -> State 169
      BIT_OR_ASSIGN -> State 172
      BIT_XOR_ASSIGN -> State 174
      BIT_LSHIFT_ASSIGN -> State 170
      BIT_RSHIFT_ASSIGN -> State 173
      unary_op_assign -> State 186
      binary_op_assign -> State 184
      optional_generic_binding -> State 185

  State 56:
    Kernel Items:
      binary_add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 187
      SUB -> State 190
      BIT_XOR -> State 189
      BIT_OR -> State 188
      add_op -> State 191

  State 57:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 192

  State 58:
    Kernel Items:
      atom_expr: anonymous_func_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 59:
    Kernel Items:
      initializable_type: array_type., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 60:
    Kernel Items:
      assign_statement: assign_pattern.ASSIGN expression
    Reduce:
      (nil)
    Goto:
      ASSIGN -> State 193

  State 61:
    Kernel Items:
      simple_statement: assign_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 62:
    Kernel Items:
      accessible_expr: atom_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 63:
    Kernel Items:
      add_expr: binary_add_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      (nil)

  State 64:
    Kernel Items:
      and_expr: binary_and_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      cmp_expr: binary_cmp_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      (nil)

  State 66:
    Kernel Items:
      mul_expr: binary_mul_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 67:
    Kernel Items:
      simple_statement: binary_op_assign_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 68:
    Kernel Items:
      or_expr: binary_or_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      (nil)

  State 69:
    Kernel Items:
      atom_expr: block_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 70:
    Kernel Items:
      accessible_expr: call_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 71:
    Kernel Items:
      callback_op_statement: callback_op.call_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      PARSE_ERROR -> State 46
      optional_label_decl -> State 148
      call_expr -> State 195
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 194
      access_expr -> State 54
      index_expr -> State 82
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 72:
    Kernel Items:
      simple_statement: callback_op_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 73:
    Kernel Items:
      binary_cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 196
      NOT_EQUAL -> State 201
      LESS -> State 199
      LESS_OR_EQUAL -> State 200
      GREATER -> State 197
      GREATER_OR_EQUAL -> State 198
      cmp_op -> State 202

  State 74:
    Kernel Items:
      initializable_type: explicit_struct_def., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 75:
    Kernel Items:
      expressions: expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 76:
    Kernel Items:
      simple_statement: expression_or_improper_struct_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 77:
    Kernel Items:
      expression_or_improper_struct_statement: expressions., *
      expressions: expressions.COMMA expression
    Reduce:
      * -> [expression_or_improper_struct_statement]
    Goto:
      COMMA -> State 203

  State 78:
    Kernel Items:
      simple_statement: fallthrough_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 79:
    Kernel Items:
      atom_expr: identifier_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 80:
    Kernel Items:
      atom_expr: implicit_struct_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 81:
    Kernel Items:
      statement: import_statement., $
    Reduce:
      $ -> [statement]
    Goto:
      (nil)

  State 82:
    Kernel Items:
      accessible_expr: index_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      initialize_expr: initializable_type.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 204

  State 84:
    Kernel Items:
      atom_expr: initialize_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 85:
    Kernel Items:
      simple_statement: jump_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 86:
    Kernel Items:
      jump_statement: jump_type., $
      jump_statement: jump_type., NEWLINES
      jump_statement: jump_type., RBRACE
      jump_statement: jump_type., SEMICOLON
      jump_statement: jump_type.expressions
      jump_statement: jump_type.JUMP_LABEL
      jump_statement: jump_type.JUMP_LABEL expressions
    Reduce:
      * -> [optional_label_decl]
      $ -> [jump_statement]
      NEWLINES -> [jump_statement]
      RBRACE -> [jump_statement]
      SEMICOLON -> [jump_statement]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      JUMP_LABEL -> State 205
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      expressions -> State 206
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 75
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 87:
    Kernel Items:
      atom_expr: literal_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      initializable_type: map_type., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      binary_mul_expr: mul_expr.mul_op prefixable_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 212
      DIV -> State 210
      MOD -> State 211
      BIT_AND -> State 207
      BIT_LSHIFT -> State 208
      BIT_RSHIFT -> State 209
      mul_op -> State 213

  State 90:
    Kernel Items:
      expression: optional_label_decl.if_expr
      expression: optional_label_decl.switch_expr
      expression: optional_label_decl.loop_expr
      block_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    Goto:
      IF -> State 216
      SWITCH -> State 217
      FOR -> State 215
      DO -> State 214
      LBRACE -> State 11
      statement_block -> State 220
      if_expr -> State 218
      switch_expr -> State 221
      loop_expr -> State 219

  State 91:
    Kernel Items:
      sequence_expr: or_expr., *
      binary_or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    Goto:
      OR -> State 222

  State 92:
    Kernel Items:
      atom_expr: parse_error_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 93:
    Kernel Items:
      postfixable_expr: postfix_unary_expr., *
    Reduce:
      * -> [postfixable_expr]
    Goto:
      (nil)

  State 94:
    Kernel Items:
      prefixable_expr: postfixable_expr., *
    Reduce:
      * -> [prefixable_expr]
    Goto:
      (nil)

  State 95:
    Kernel Items:
      prefixable_expr: prefix_unary_expr., *
    Reduce:
      * -> [prefixable_expr]
    Goto:
      (nil)

  State 96:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op.prefixable_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      PARSE_ERROR -> State 46
      optional_label_decl -> State 148
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 223
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 97:
    Kernel Items:
      mul_expr: prefixable_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 98:
    Kernel Items:
      assign_pattern: sequence_expr., ASSIGN
      expression: sequence_expr., *
    Reduce:
      * -> [expression]
      ASSIGN -> [assign_pattern]
    Goto:
      (nil)

  State 99:
    Kernel Items:
      statement: simple_statement., $
    Reduce:
      $ -> [statement]
    Goto:
      (nil)

  State 100:
    Kernel Items:
      initializable_type: slice_type., *
    Reduce:
      * -> [initializable_type]
    Goto:
      (nil)

  State 101:
    Kernel Items:
      simple_statement: unary_op_assign_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 102:
    Kernel Items:
      simple_statement: unsafe_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 103:
    Kernel Items:
      sequence_expr: var_decl_pattern., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 104:
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
      LBRACKET -> State 178
      DOT -> State 177
      QUESTION -> State 181
      DOLLAR_LBRACKET -> State 176
      optional_generic_binding -> State 185

  State 105:
    Kernel Items:
      expression: sequence_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 106:
    Kernel Items:
      prefix_type_op: BIT_AND., *
    Reduce:
      * -> [prefix_type_op]
    Goto:
      (nil)

  State 107:
    Kernel Items:
      prefix_type_op: BIT_NEG., *
    Reduce:
      * -> [prefix_type_op]
    Goto:
      (nil)

  State 108:
    Kernel Items:
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 176
      optional_generic_binding -> State 224

  State 109:
    Kernel Items:
      explicit_enum_def: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 225

  State 110:
    Kernel Items:
      prefix_type_op: EXCLAIM., *
    Reduce:
      * -> [prefix_type_op]
    Goto:
      (nil)

  State 111:
    Kernel Items:
      func_type: FUNC.LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 226

  State 112:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOT -> State 227
      DOLLAR_LBRACKET -> State 176
      optional_generic_binding -> State 228

  State 113:
    Kernel Items:
      implicit_struct_def: LPAREN.implicit_field_defs RPAREN
      implicit_enum_def: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      IDENTIFIER -> State 229
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 235
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 236
      trait_op_type -> State 129
      field_def -> State 231
      proper_implicit_field_defs -> State 234
      implicit_field_defs -> State 233
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      enum_value_def -> State 230
      implicit_enum_value_defs -> State 232
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 114:
    Kernel Items:
      parse_error_type: PARSE_ERROR., *
    Reduce:
      * -> [parse_error_type]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      prefix_type_op: QUESTION., *
    Reduce:
      * -> [prefix_type_op]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      prefix_type_op: TILDE_TILDE., *
    Reduce:
      * -> [prefix_type_op]
    Goto:
      (nil)

  State 117:
    Kernel Items:
      trait_def: TRAIT.LPAREN trait_properties RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 237

  State 118:
    Kernel Items:
      returnable_type: atom_type., *
    Reduce:
      * -> [returnable_type]
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
      atom_type: func_type., *
    Reduce:
      * -> [atom_type]
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
      atom_type: initializable_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      atom_type: parse_error_type., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      prefixed_type: prefix_type_op.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 238
      prefixed_type -> State 126
      prefix_type_op -> State 125
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 126:
    Kernel Items:
      returnable_type: prefixed_type., *
    Reduce:
      * -> [returnable_type]
    Goto:
      (nil)

  State 127:
    Kernel Items:
      value_type: returnable_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 128:
    Kernel Items:
      atom_type: trait_def., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 129:
    Kernel Items:
      value_type: trait_op_type., *
    Reduce:
      * -> [value_type]
    Goto:
      (nil)

  State 130:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER.optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC IDENTIFIER.ASSIGN expression
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 244
      ASSIGN -> State 243
      optional_generic_parameters -> State 245

  State 131:
    Kernel Items:
      named_func_def: FUNC LPAREN.parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 246
      parameter_def -> State 247

  State 132:
    Kernel Items:
      proper_statements: proper_statements.NEWLINES statement
      proper_statements: proper_statements.SEMICOLON statement
      statements: proper_statements., RBRACE
      statements: proper_statements.NEWLINES
      statements: proper_statements.SEMICOLON
    Reduce:
      RBRACE -> [statements]
    Goto:
      NEWLINES -> State 248
      SEMICOLON -> State 249

  State 133:
    Kernel Items:
      proper_statements: statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 134:
    Kernel Items:
      statement_block: LBRACE statements.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 250

  State 135:
    Kernel Items:
      package_def: PACKAGE statement_block., *
    Reduce:
      * -> [package_def]
    Goto:
      (nil)

  State 136:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type
      type_def: TYPE IDENTIFIER.optional_generic_parameters value_type IMPLEMENTS value_type
      type_def: TYPE IDENTIFIER.ASSIGN value_type
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 244
      ASSIGN -> State 251
      optional_generic_parameters -> State 252

  State 137:
    Kernel Items:
      proper_definitions: proper_definitions NEWLINES.definition
      definitions: proper_definitions NEWLINES., $
    Reduce:
      $ -> [definitions]
    Goto:
      COMMENT_GROUPS -> State 9
      PACKAGE -> State 13
      TYPE -> State 14
      FUNC -> State 10
      VAR -> State 15
      LET -> State 12
      LBRACE -> State 11
      definition -> State 253
      statement_block -> State 21
      var_decl_pattern -> State 23
      var_or_let -> State 24
      type_def -> State 22
      named_func_def -> State 18
      package_def -> State 19

  State 138:
    Kernel Items:
      definition: var_decl_pattern ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 254
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 139:
    Kernel Items:
      var_pattern: IDENTIFIER., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 140:
    Kernel Items:
      tuple_pattern: LPAREN.field_var_patterns RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 256
      LPAREN -> State 140
      ELLIPSIS -> State 255
      var_pattern -> State 259
      tuple_pattern -> State 141
      field_var_patterns -> State 258
      field_var_pattern -> State 257

  State 141:
    Kernel Items:
      var_pattern: tuple_pattern., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 142:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern.optional_value_type
    Reduce:
      * -> [optional_value_type]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      optional_value_type -> State 260
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 261
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 143:
    Kernel Items:
      case_pattern: DOT.IDENTIFIER implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 262

  State 144:
    Kernel Items:
      var_or_let: VAR., *
      case_pattern: VAR.DOT IDENTIFIER tuple_pattern
    Reduce:
      * -> [var_or_let]
    Goto:
      DOT -> State 263

  State 145:
    Kernel Items:
      case_pattern: assign_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 146:
    Kernel Items:
      case_patterns: case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 147:
    Kernel Items:
      statement: CASE case_patterns.COLON optional_simple_statement
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    Goto:
      COMMA -> State 265
      COLON -> State 264

  State 148:
    Kernel Items:
      block_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 220

  State 149:
    Kernel Items:
      assign_pattern: sequence_expr., *
    Reduce:
      * -> [assign_pattern]
    Goto:
      (nil)

  State 150:
    Kernel Items:
      statement: DEFAULT COLON.optional_simple_statement
    Reduce:
      * -> [optional_label_decl]
      $ -> [optional_simple_statement]
      NEWLINES -> [optional_simple_statement]
      RBRACE -> [optional_simple_statement]
      SEMICOLON -> [optional_simple_statement]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      RETURN -> State 47
      BREAK -> State 28
      CONTINUE -> State 30
      FALLTHROUGH -> State 33
      UNSAFE -> State 53
      STRUCT -> State 50
      FUNC -> State 36
      ASYNC -> State 25
      DEFER -> State 32
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      simple_statement -> State 267
      optional_simple_statement -> State 266
      expression_or_improper_struct_statement -> State 76
      expressions -> State 77
      callback_op -> State 71
      callback_op_statement -> State 72
      unsafe_statement -> State 102
      jump_statement -> State 85
      jump_type -> State 86
      fallthrough_statement -> State 78
      assign_statement -> State 61
      unary_op_assign_statement -> State 101
      binary_op_assign_statement -> State 67
      var_decl_pattern -> State 103
      var_or_let -> State 24
      assign_pattern -> State 60
      expression -> State 75
      optional_label_decl -> State 90
      sequence_expr -> State 98
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 55
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 151:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 246
      parameter_def -> State 268
      proper_parameter_defs -> State 270
      parameter_defs -> State 269

  State 152:
    Kernel Items:
      sequence_expr: GREATER sequence_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 153:
    Kernel Items:
      import_statement: IMPORT LPAREN.import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 154
      import_clause -> State 271
      proper_import_clauses -> State 273
      import_clauses -> State 272

  State 154:
    Kernel Items:
      import_clause: STRING_LITERAL., *
      import_clause: STRING_LITERAL.AS IDENTIFIER
    Reduce:
      * -> [import_clause]
    Goto:
      AS -> State 274

  State 155:
    Kernel Items:
      import_statement: IMPORT import_clause., $
    Reduce:
      $ -> [import_statement]
    Goto:
      (nil)

  State 156:
    Kernel Items:
      slice_type: LBRACKET value_type.RBRACKET
      array_type: LBRACKET value_type.COMMA INTEGER_LITERAL RBRACKET
      map_type: LBRACKET value_type.COLON value_type RBRACKET
      trait_op_type: value_type.trait_op returnable_type
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 277
      COMMA -> State 276
      COLON -> State 275
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 157:
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
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 278
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 158:
    Kernel Items:
      argument: ELLIPSIS., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 159:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expression
      identifier_expr: IDENTIFIER., *
    Reduce:
      * -> [identifier_expr]
    Goto:
      ASSIGN -> State 279

  State 160:
    Kernel Items:
      proper_arguments: argument., *
    Reduce:
      * -> [proper_arguments]
    Goto:
      (nil)

  State 161:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 280

  State 162:
    Kernel Items:
      argument: colon_expr., *
      colon_expr: colon_expr.COLON
      colon_expr: colon_expr.COLON expression
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 281

  State 163:
    Kernel Items:
      argument: expression., *
      argument: expression.ELLIPSIS
      colon_expr: expression.COLON
      colon_expr: expression.COLON expression
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 282
      ELLIPSIS -> State 283

  State 164:
    Kernel Items:
      proper_arguments: proper_arguments.COMMA argument
      arguments: proper_arguments., RPAREN
      arguments: proper_arguments.COMMA
    Reduce:
      RPAREN -> [arguments]
    Goto:
      COMMA -> State 284

  State 165:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN.explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 229
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 235
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 236
      trait_op_type -> State 129
      field_def -> State 286
      implicit_struct_def -> State 122
      proper_explicit_field_defs -> State 287
      explicit_field_defs -> State 285
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 166:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 288

  State 167:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 168:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., $
    Reduce:
      $ -> [unary_op_assign]
    Goto:
      (nil)

  State 169:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 170:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 171:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 172:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 173:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 174:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 175:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 176:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET.generic_arguments RBRACKET
    Reduce:
      RBRACKET -> [generic_arguments]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      proper_generic_arguments -> State 290
      generic_arguments -> State 289
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 291
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 177:
    Kernel Items:
      access_expr: accessible_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 292

  State 178:
    Kernel Items:
      index_expr: accessible_expr LBRACKET.argument RBRACKET
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 159
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      COLON -> State 157
      ELLIPSIS -> State 158
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 163
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      argument -> State 293
      colon_expr -> State 162
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 179:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 180:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 181:
    Kernel Items:
      postfix_unary_expr: accessible_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 182:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 183:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., $
    Reduce:
      $ -> [unary_op_assign]
    Goto:
      (nil)

  State 184:
    Kernel Items:
      binary_op_assign_statement: accessible_expr binary_op_assign.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 294
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 185:
    Kernel Items:
      call_expr: accessible_expr optional_generic_binding.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 295

  State 186:
    Kernel Items:
      unary_op_assign_statement: accessible_expr unary_op_assign., $
    Reduce:
      $ -> [unary_op_assign_statement]
    Goto:
      (nil)

  State 187:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 188:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 189:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 190:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 191:
    Kernel Items:
      binary_add_expr: add_expr add_op.mul_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      PARSE_ERROR -> State 46
      optional_label_decl -> State 148
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 296
      binary_mul_expr -> State 66
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 192:
    Kernel Items:
      binary_and_expr: and_expr AND.cmp_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      PARSE_ERROR -> State 46
      optional_label_decl -> State 148
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 297
      binary_cmp_expr -> State 65
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 193:
    Kernel Items:
      assign_statement: assign_pattern ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 298
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 194:
    Kernel Items:
      call_expr: accessible_expr.optional_generic_binding LPAREN arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
    Reduce:
      LPAREN -> [optional_generic_binding]
    Goto:
      LBRACKET -> State 178
      DOT -> State 177
      DOLLAR_LBRACKET -> State 176
      optional_generic_binding -> State 185

  State 195:
    Kernel Items:
      callback_op_statement: callback_op call_expr., $
      callback_op_statement: callback_op call_expr., NEWLINES
      callback_op_statement: callback_op call_expr., RBRACE
      callback_op_statement: callback_op call_expr., SEMICOLON
      accessible_expr: call_expr., *
    Reduce:
      * -> [accessible_expr]
      $ -> [callback_op_statement]
      NEWLINES -> [callback_op_statement]
      RBRACE -> [callback_op_statement]
      SEMICOLON -> [callback_op_statement]
    Goto:
      (nil)

  State 196:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 197:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 198:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 199:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 200:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 201:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 202:
    Kernel Items:
      binary_cmp_expr: cmp_expr cmp_op.add_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      PARSE_ERROR -> State 46
      optional_label_decl -> State 148
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 299
      binary_add_expr -> State 63
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 203:
    Kernel Items:
      expressions: expressions COMMA.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 300
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 204:
    Kernel Items:
      initialize_expr: initializable_type LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [arguments]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 159
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      COLON -> State 157
      ELLIPSIS -> State 158
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 163
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      proper_arguments -> State 164
      arguments -> State 301
      argument -> State 160
      colon_expr -> State 162
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 205:
    Kernel Items:
      jump_statement: jump_type JUMP_LABEL., $
      jump_statement: jump_type JUMP_LABEL., NEWLINES
      jump_statement: jump_type JUMP_LABEL., RBRACE
      jump_statement: jump_type JUMP_LABEL., SEMICOLON
      jump_statement: jump_type JUMP_LABEL.expressions
    Reduce:
      * -> [optional_label_decl]
      $ -> [jump_statement]
      NEWLINES -> [jump_statement]
      RBRACE -> [jump_statement]
      SEMICOLON -> [jump_statement]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      expressions -> State 302
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 75
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 206:
    Kernel Items:
      expressions: expressions.COMMA expression
      jump_statement: jump_type expressions., $
    Reduce:
      $ -> [jump_statement]
    Goto:
      COMMA -> State 203

  State 207:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 208:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 209:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 211:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 212:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      binary_mul_expr: mul_expr mul_op.prefixable_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      PARSE_ERROR -> State 46
      optional_label_decl -> State 148
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 303
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 214:
    Kernel Items:
      loop_expr: DO.statement_block
      loop_expr: DO.statement_block FOR sequence_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 304

  State 215:
    Kernel Items:
      loop_expr: FOR.sequence_expr DO statement_block
      loop_expr: FOR.assign_pattern IN sequence_expr DO statement_block
      loop_expr: FOR.for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      assign_pattern -> State 305
      optional_label_decl -> State 148
      sequence_expr -> State 307
      for_assignment -> State 306
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 216:
    Kernel Items:
      if_expr: IF.condition statement_block
      if_expr: IF.condition statement_block ELSE statement_block
      if_expr: IF.condition statement_block ELSE if_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      CASE -> State 308
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      optional_label_decl -> State 148
      sequence_expr -> State 310
      condition -> State 309
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 217:
    Kernel Items:
      switch_expr: SWITCH.sequence_expr statement_block
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      optional_label_decl -> State 148
      sequence_expr -> State 311
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 218:
    Kernel Items:
      expression: optional_label_decl if_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 219:
    Kernel Items:
      expression: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 220:
    Kernel Items:
      block_expr: optional_label_decl statement_block., *
    Reduce:
      * -> [block_expr]
    Goto:
      (nil)

  State 221:
    Kernel Items:
      expression: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expression]
    Goto:
      (nil)

  State 222:
    Kernel Items:
      binary_or_expr: or_expr OR.and_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      PARSE_ERROR -> State 46
      optional_label_decl -> State 148
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 312
      binary_and_expr -> State 64
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 223:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefixable_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 224:
    Kernel Items:
      atom_type: DOT optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 225:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 229
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 235
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 236
      trait_op_type -> State 129
      field_def -> State 315
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      enum_value_def -> State 313
      implicit_enum_value_defs -> State 316
      implicit_enum_def -> State 121
      explicit_enum_value_defs -> State 314
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 226:
    Kernel Items:
      func_type: FUNC LPAREN.parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 318
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      ELLIPSIS -> State 317
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 322
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      parameter_decl -> State 319
      proper_parameter_decls -> State 321
      parameter_decls -> State 320
      func_type -> State 120

  State 227:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 323

  State 228:
    Kernel Items:
      atom_type: IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 229:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      field_def: IDENTIFIER.value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 324
      QUESTION -> State 115
      EXCLAIM -> State 110
      DOLLAR_LBRACKET -> State 176
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      optional_generic_binding -> State 228
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 325
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 230:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
    Reduce:
      (nil)
    Goto:
      OR -> State 326

  State 231:
    Kernel Items:
      proper_implicit_field_defs: field_def., *
      enum_value_def: field_def., OR
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [proper_implicit_field_defs]
      OR -> [enum_value_def]
    Goto:
      ASSIGN -> State 327

  State 232:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      implicit_enum_def: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      OR -> State 328
      RPAREN -> State 329

  State 233:
    Kernel Items:
      implicit_struct_def: LPAREN implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 330

  State 234:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs.COMMA field_def
      implicit_field_defs: proper_implicit_field_defs., RPAREN
      implicit_field_defs: proper_implicit_field_defs.COMMA
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      COMMA -> State 331

  State 235:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 236:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      field_def: value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 237:
    Kernel Items:
      trait_def: TRAIT LPAREN.trait_properties RPAREN
    Reduce:
      RPAREN -> [trait_properties]
    Goto:
      IDENTIFIER -> State 229
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 332
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 235
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 236
      trait_op_type -> State 129
      field_def -> State 333
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_property -> State 337
      proper_trait_properties -> State 335
      trait_properties -> State 336
      trait_def -> State 128
      func_type -> State 120
      method_signature -> State 334

  State 238:
    Kernel Items:
      prefixed_type: prefix_type_op returnable_type., *
    Reduce:
      * -> [prefixed_type]
    Goto:
      (nil)

  State 239:
    Kernel Items:
      trait_op: ADD., *
    Reduce:
      * -> [trait_op]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      trait_op: MUL., *
    Reduce:
      * -> [trait_op]
    Goto:
      (nil)

  State 241:
    Kernel Items:
      trait_op: SUB., *
    Reduce:
      * -> [trait_op]
    Goto:
      (nil)

  State 242:
    Kernel Items:
      trait_op_type: value_type trait_op.returnable_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 338
      prefixed_type -> State 126
      prefix_type_op -> State 125
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 243:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 339
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 244:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET.generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 340
      generic_parameter_def -> State 341
      proper_generic_parameter_defs -> State 343
      generic_parameter_defs -> State 342

  State 245:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 344

  State 246:
    Kernel Items:
      parameter_def: IDENTIFIER., *
      parameter_def: IDENTIFIER.ELLIPSIS
      parameter_def: IDENTIFIER.value_type
      parameter_def: IDENTIFIER.ELLIPSIS value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      ELLIPSIS -> State 345
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 346
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 247:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def.RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 347

  State 248:
    Kernel Items:
      proper_statements: proper_statements NEWLINES.statement
      statements: proper_statements NEWLINES., RBRACE
    Reduce:
      * -> [optional_label_decl]
      RBRACE -> [statements]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      CASE -> State 29
      DEFAULT -> State 31
      RETURN -> State 47
      BREAK -> State 28
      CONTINUE -> State 30
      FALLTHROUGH -> State 33
      IMPORT -> State 39
      UNSAFE -> State 53
      STRUCT -> State 50
      FUNC -> State 36
      ASYNC -> State 25
      DEFER -> State 32
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      simple_statement -> State 99
      statement -> State 348
      expression_or_improper_struct_statement -> State 76
      expressions -> State 77
      callback_op -> State 71
      callback_op_statement -> State 72
      unsafe_statement -> State 102
      jump_statement -> State 85
      jump_type -> State 86
      fallthrough_statement -> State 78
      assign_statement -> State 61
      unary_op_assign_statement -> State 101
      binary_op_assign_statement -> State 67
      import_statement -> State 81
      var_decl_pattern -> State 103
      var_or_let -> State 24
      assign_pattern -> State 60
      expression -> State 75
      optional_label_decl -> State 90
      sequence_expr -> State 98
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 55
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 249:
    Kernel Items:
      proper_statements: proper_statements SEMICOLON.statement
      statements: proper_statements SEMICOLON., RBRACE
    Reduce:
      * -> [optional_label_decl]
      RBRACE -> [statements]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      CASE -> State 29
      DEFAULT -> State 31
      RETURN -> State 47
      BREAK -> State 28
      CONTINUE -> State 30
      FALLTHROUGH -> State 33
      IMPORT -> State 39
      UNSAFE -> State 53
      STRUCT -> State 50
      FUNC -> State 36
      ASYNC -> State 25
      DEFER -> State 32
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      simple_statement -> State 99
      statement -> State 349
      expression_or_improper_struct_statement -> State 76
      expressions -> State 77
      callback_op -> State 71
      callback_op_statement -> State 72
      unsafe_statement -> State 102
      jump_statement -> State 85
      jump_type -> State 86
      fallthrough_statement -> State 78
      assign_statement -> State 61
      unary_op_assign_statement -> State 101
      binary_op_assign_statement -> State 67
      import_statement -> State 81
      var_decl_pattern -> State 103
      var_or_let -> State 24
      assign_pattern -> State 60
      expression -> State 75
      optional_label_decl -> State 90
      sequence_expr -> State 98
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 55
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 250:
    Kernel Items:
      statement_block: LBRACE statements RBRACE., $
    Reduce:
      $ -> [statement_block]
    Goto:
      (nil)

  State 251:
    Kernel Items:
      type_def: TYPE IDENTIFIER ASSIGN.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 350
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 252:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type
      type_def: TYPE IDENTIFIER optional_generic_parameters.value_type IMPLEMENTS value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 351
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 253:
    Kernel Items:
      proper_definitions: proper_definitions NEWLINES definition., *
    Reduce:
      * -> [proper_definitions]
    Goto:
      (nil)

  State 254:
    Kernel Items:
      definition: var_decl_pattern ASSIGN expression., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

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
      ASSIGN -> State 352

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
      RPAREN -> State 354
      COMMA -> State 353

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
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 262:
    Kernel Items:
      case_pattern: DOT IDENTIFIER.implicit_struct_expr
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 43
      implicit_struct_expr -> State 355

  State 263:
    Kernel Items:
      case_pattern: VAR DOT.IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 356

  State 264:
    Kernel Items:
      statement: CASE case_patterns COLON.optional_simple_statement
    Reduce:
      * -> [optional_label_decl]
      $ -> [optional_simple_statement]
      NEWLINES -> [optional_simple_statement]
      RBRACE -> [optional_simple_statement]
      SEMICOLON -> [optional_simple_statement]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      RETURN -> State 47
      BREAK -> State 28
      CONTINUE -> State 30
      FALLTHROUGH -> State 33
      UNSAFE -> State 53
      STRUCT -> State 50
      FUNC -> State 36
      ASYNC -> State 25
      DEFER -> State 32
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      simple_statement -> State 267
      optional_simple_statement -> State 357
      expression_or_improper_struct_statement -> State 76
      expressions -> State 77
      callback_op -> State 71
      callback_op_statement -> State 72
      unsafe_statement -> State 102
      jump_statement -> State 85
      jump_type -> State 86
      fallthrough_statement -> State 78
      assign_statement -> State 61
      unary_op_assign_statement -> State 101
      binary_op_assign_statement -> State 67
      var_decl_pattern -> State 103
      var_or_let -> State 24
      assign_pattern -> State 60
      expression -> State 75
      optional_label_decl -> State 90
      sequence_expr -> State 98
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 55
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 265:
    Kernel Items:
      case_patterns: case_patterns COMMA.case_pattern
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 144
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      DOT -> State 143
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      assign_pattern -> State 145
      case_pattern -> State 358
      optional_label_decl -> State 148
      sequence_expr -> State 149
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 266:
    Kernel Items:
      statement: DEFAULT COLON optional_simple_statement., $
    Reduce:
      $ -> [statement]
    Goto:
      (nil)

  State 267:
    Kernel Items:
      optional_simple_statement: simple_statement., $
    Reduce:
      $ -> [optional_simple_statement]
    Goto:
      (nil)

  State 268:
    Kernel Items:
      proper_parameter_defs: parameter_def., *
    Reduce:
      * -> [proper_parameter_defs]
    Goto:
      (nil)

  State 269:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 359

  State 270:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs.COMMA parameter_def
      parameter_defs: proper_parameter_defs., RPAREN
      parameter_defs: proper_parameter_defs.COMMA
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      COMMA -> State 360

  State 271:
    Kernel Items:
      proper_import_clauses: import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 272:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 361

  State 273:
    Kernel Items:
      proper_import_clauses: proper_import_clauses.NEWLINES import_clause
      proper_import_clauses: proper_import_clauses.COMMA import_clause
      import_clauses: proper_import_clauses., RPAREN
      import_clauses: proper_import_clauses.NEWLINES
      import_clauses: proper_import_clauses.COMMA
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      NEWLINES -> State 363
      COMMA -> State 362

  State 274:
    Kernel Items:
      import_clause: STRING_LITERAL AS.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 364

  State 275:
    Kernel Items:
      map_type: LBRACKET value_type COLON.value_type RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 365
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 276:
    Kernel Items:
      array_type: LBRACKET value_type COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 366

  State 277:
    Kernel Items:
      slice_type: LBRACKET value_type RBRACKET., *
    Reduce:
      * -> [slice_type]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      colon_expr: COLON expression., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      argument: IDENTIFIER ASSIGN.expression
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 367
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 280:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [implicit_struct_expr]
    Goto:
      (nil)

  State 281:
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
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 368
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 282:
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
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 369
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 283:
    Kernel Items:
      argument: expression ELLIPSIS., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 284:
    Kernel Items:
      proper_arguments: proper_arguments COMMA.argument
      arguments: proper_arguments COMMA., RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [arguments]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 159
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      COLON -> State 157
      ELLIPSIS -> State 158
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 163
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      argument -> State 370
      colon_expr -> State 162
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 285:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 371

  State 286:
    Kernel Items:
      proper_explicit_field_defs: field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 287:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs.NEWLINES field_def
      proper_explicit_field_defs: proper_explicit_field_defs.COMMA field_def
      explicit_field_defs: proper_explicit_field_defs., RPAREN
      explicit_field_defs: proper_explicit_field_defs.NEWLINES
      explicit_field_defs: proper_explicit_field_defs.COMMA
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      NEWLINES -> State 373
      COMMA -> State 372

  State 288:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 374

  State 289:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET generic_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 375

  State 290:
    Kernel Items:
      proper_generic_arguments: proper_generic_arguments.COMMA value_type
      generic_arguments: proper_generic_arguments., RBRACKET
      generic_arguments: proper_generic_arguments.COMMA
    Reduce:
      RBRACKET -> [generic_arguments]
    Goto:
      COMMA -> State 376

  State 291:
    Kernel Items:
      proper_generic_arguments: value_type., *
      trait_op_type: value_type.trait_op returnable_type
    Reduce:
      * -> [proper_generic_arguments]
    Goto:
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 292:
    Kernel Items:
      access_expr: accessible_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 293:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 377

  State 294:
    Kernel Items:
      binary_op_assign_statement: accessible_expr binary_op_assign expression., $
    Reduce:
      $ -> [binary_op_assign_statement]
    Goto:
      (nil)

  State 295:
    Kernel Items:
      call_expr: accessible_expr optional_generic_binding LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [arguments]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 159
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      COLON -> State 157
      ELLIPSIS -> State 158
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      expression -> State 163
      optional_label_decl -> State 90
      sequence_expr -> State 105
      call_expr -> State 70
      proper_arguments -> State 164
      arguments -> State 378
      argument -> State 160
      colon_expr -> State 162
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 296:
    Kernel Items:
      binary_mul_expr: mul_expr.mul_op prefixable_expr
      binary_add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [binary_add_expr]
    Goto:
      MUL -> State 212
      DIV -> State 210
      MOD -> State 211
      BIT_AND -> State 207
      BIT_LSHIFT -> State 208
      BIT_RSHIFT -> State 209
      mul_op -> State 213

  State 297:
    Kernel Items:
      binary_cmp_expr: cmp_expr.cmp_op add_expr
      binary_and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [binary_and_expr]
    Goto:
      EQUAL -> State 196
      NOT_EQUAL -> State 201
      LESS -> State 199
      LESS_OR_EQUAL -> State 200
      GREATER -> State 197
      GREATER_OR_EQUAL -> State 198
      cmp_op -> State 202

  State 298:
    Kernel Items:
      assign_statement: assign_pattern ASSIGN expression., $
    Reduce:
      $ -> [assign_statement]
    Goto:
      (nil)

  State 299:
    Kernel Items:
      binary_add_expr: add_expr.add_op mul_expr
      binary_cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [binary_cmp_expr]
    Goto:
      ADD -> State 187
      SUB -> State 190
      BIT_XOR -> State 189
      BIT_OR -> State 188
      add_op -> State 191

  State 300:
    Kernel Items:
      expressions: expressions COMMA expression., *
    Reduce:
      * -> [expressions]
    Goto:
      (nil)

  State 301:
    Kernel Items:
      initialize_expr: initializable_type LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 379

  State 302:
    Kernel Items:
      expressions: expressions.COMMA expression
      jump_statement: jump_type JUMP_LABEL expressions., $
    Reduce:
      $ -> [jump_statement]
    Goto:
      COMMA -> State 203

  State 303:
    Kernel Items:
      binary_mul_expr: mul_expr mul_op prefixable_expr., *
    Reduce:
      * -> [binary_mul_expr]
    Goto:
      (nil)

  State 304:
    Kernel Items:
      loop_expr: DO statement_block., *
      loop_expr: DO statement_block.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 380

  State 305:
    Kernel Items:
      loop_expr: FOR assign_pattern.IN sequence_expr DO statement_block
      for_assignment: assign_pattern.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      IN -> State 382
      ASSIGN -> State 381

  State 306:
    Kernel Items:
      loop_expr: FOR for_assignment.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 383

  State 307:
    Kernel Items:
      assign_pattern: sequence_expr., *
      loop_expr: FOR sequence_expr.DO statement_block
      for_assignment: sequence_expr., SEMICOLON
    Reduce:
      * -> [assign_pattern]
      SEMICOLON -> [for_assignment]
    Goto:
      DO -> State 384

  State 308:
    Kernel Items:
      condition: CASE.case_patterns ASSIGN sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 144
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      DOT -> State 143
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      case_patterns -> State 385
      var_decl_pattern -> State 103
      var_or_let -> State 24
      assign_pattern -> State 145
      case_pattern -> State 146
      optional_label_decl -> State 148
      sequence_expr -> State 149
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 309:
    Kernel Items:
      if_expr: IF condition.statement_block
      if_expr: IF condition.statement_block ELSE statement_block
      if_expr: IF condition.statement_block ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 386

  State 310:
    Kernel Items:
      condition: sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 311:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 387

  State 312:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      binary_or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [binary_or_expr]
    Goto:
      AND -> State 192

  State 313:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.OR enum_value_def
      explicit_enum_value_defs: enum_value_def.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 388
      OR -> State 389

  State 314:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 390

  State 315:
    Kernel Items:
      enum_value_def: field_def., *
      enum_value_def: field_def.ASSIGN DEFAULT
    Reduce:
      * -> [enum_value_def]
    Goto:
      ASSIGN -> State 327

  State 316:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.OR enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs.NEWLINES enum_value_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 391
      OR -> State 392

  State 317:
    Kernel Items:
      parameter_decl: ELLIPSIS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 393
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 318:
    Kernel Items:
      atom_type: IDENTIFIER.optional_generic_binding
      atom_type: IDENTIFIER.DOT IDENTIFIER optional_generic_binding
      parameter_decl: IDENTIFIER.value_type
      parameter_decl: IDENTIFIER.ELLIPSIS value_type
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 324
      QUESTION -> State 115
      EXCLAIM -> State 110
      DOLLAR_LBRACKET -> State 176
      ELLIPSIS -> State 394
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      optional_generic_binding -> State 228
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 395
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 319:
    Kernel Items:
      proper_parameter_decls: parameter_decl., *
    Reduce:
      * -> [proper_parameter_decls]
    Goto:
      (nil)

  State 320:
    Kernel Items:
      func_type: FUNC LPAREN parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 396

  State 321:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls.COMMA parameter_decl
      parameter_decls: proper_parameter_decls., RPAREN
      parameter_decls: proper_parameter_decls.COMMA
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      COMMA -> State 397

  State 322:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_decl: value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 323:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      DOLLAR_LBRACKET -> State 176
      optional_generic_binding -> State 398

  State 324:
    Kernel Items:
      atom_type: IDENTIFIER DOT.IDENTIFIER optional_generic_binding
      atom_type: DOT.optional_generic_binding
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      IDENTIFIER -> State 323
      DOLLAR_LBRACKET -> State 176
      optional_generic_binding -> State 224

  State 325:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      field_def: IDENTIFIER value_type., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 326:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 229
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 235
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 236
      trait_op_type -> State 129
      field_def -> State 315
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      enum_value_def -> State 399
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 327:
    Kernel Items:
      enum_value_def: field_def ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 400

  State 328:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 229
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 235
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 236
      trait_op_type -> State 129
      field_def -> State 315
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      enum_value_def -> State 401
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 329:
    Kernel Items:
      implicit_enum_def: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_def]
    Goto:
      (nil)

  State 330:
    Kernel Items:
      implicit_struct_def: LPAREN implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_def]
    Goto:
      (nil)

  State 331:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs COMMA.field_def
      implicit_field_defs: proper_implicit_field_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      IDENTIFIER -> State 229
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 235
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 236
      trait_op_type -> State 129
      field_def -> State 402
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 332:
    Kernel Items:
      func_type: FUNC.LPAREN parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 403
      LPAREN -> State 226

  State 333:
    Kernel Items:
      trait_property: field_def., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 334:
    Kernel Items:
      trait_property: method_signature., *
    Reduce:
      * -> [trait_property]
    Goto:
      (nil)

  State 335:
    Kernel Items:
      proper_trait_properties: proper_trait_properties.NEWLINES trait_property
      proper_trait_properties: proper_trait_properties.COMMA trait_property
      trait_properties: proper_trait_properties., RPAREN
      trait_properties: proper_trait_properties.NEWLINES
      trait_properties: proper_trait_properties.COMMA
    Reduce:
      RPAREN -> [trait_properties]
    Goto:
      NEWLINES -> State 405
      COMMA -> State 404

  State 336:
    Kernel Items:
      trait_def: TRAIT LPAREN trait_properties.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 406

  State 337:
    Kernel Items:
      proper_trait_properties: trait_property., *
    Reduce:
      * -> [proper_trait_properties]
    Goto:
      (nil)

  State 338:
    Kernel Items:
      trait_op_type: value_type trait_op returnable_type., *
    Reduce:
      * -> [trait_op_type]
    Goto:
      (nil)

  State 339:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [named_func_def]
    Goto:
      (nil)

  State 340:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.value_type
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 407
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 341:
    Kernel Items:
      proper_generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [proper_generic_parameter_defs]
    Goto:
      (nil)

  State 342:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 408

  State 343:
    Kernel Items:
      proper_generic_parameter_defs: proper_generic_parameter_defs.COMMA generic_parameter_def
      generic_parameter_defs: proper_generic_parameter_defs., RBRACKET
      generic_parameter_defs: proper_generic_parameter_defs.COMMA
    Reduce:
      RBRACKET -> [generic_parameter_defs]
    Goto:
      COMMA -> State 409

  State 344:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 246
      parameter_def -> State 268
      proper_parameter_defs -> State 270
      parameter_defs -> State 410

  State 345:
    Kernel Items:
      parameter_def: IDENTIFIER ELLIPSIS., *
      parameter_def: IDENTIFIER ELLIPSIS.value_type
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 411
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 346:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 347:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN.IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 412

  State 348:
    Kernel Items:
      proper_statements: proper_statements NEWLINES statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 349:
    Kernel Items:
      proper_statements: proper_statements SEMICOLON statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 350:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      type_def: TYPE IDENTIFIER ASSIGN value_type., *
    Reduce:
      * -> [type_def]
    Goto:
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 351:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type., *
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type.IMPLEMENTS value_type
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 413
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 352:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN.var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 139
      LPAREN -> State 140
      var_pattern -> State 414
      tuple_pattern -> State 141

  State 353:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA.field_var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 256
      LPAREN -> State 140
      ELLIPSIS -> State 255
      var_pattern -> State 259
      tuple_pattern -> State 141
      field_var_pattern -> State 415

  State 354:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns RPAREN., *
    Reduce:
      * -> [tuple_pattern]
    Goto:
      (nil)

  State 355:
    Kernel Items:
      case_pattern: DOT IDENTIFIER implicit_struct_expr., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 356:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER.tuple_pattern
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 140
      tuple_pattern -> State 416

  State 357:
    Kernel Items:
      statement: CASE case_patterns COLON optional_simple_statement., $
    Reduce:
      $ -> [statement]
    Goto:
      (nil)

  State 358:
    Kernel Items:
      case_patterns: case_patterns COMMA case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 359:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 418
      prefixed_type -> State 126
      prefix_type_op -> State 125
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      return_type -> State 417
      func_type -> State 120

  State 360:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs COMMA.parameter_def
      parameter_defs: proper_parameter_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 246
      parameter_def -> State 419

  State 361:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses RPAREN., $
    Reduce:
      $ -> [import_statement]
    Goto:
      (nil)

  State 362:
    Kernel Items:
      proper_import_clauses: proper_import_clauses COMMA.import_clause
      import_clauses: proper_import_clauses COMMA., RPAREN
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      STRING_LITERAL -> State 154
      import_clause -> State 420

  State 363:
    Kernel Items:
      proper_import_clauses: proper_import_clauses NEWLINES.import_clause
      import_clauses: proper_import_clauses NEWLINES., RPAREN
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      STRING_LITERAL -> State 154
      import_clause -> State 421

  State 364:
    Kernel Items:
      import_clause: STRING_LITERAL AS IDENTIFIER., $
    Reduce:
      $ -> [import_clause]
    Goto:
      (nil)

  State 365:
    Kernel Items:
      map_type: LBRACKET value_type COLON value_type.RBRACKET
      trait_op_type: value_type.trait_op returnable_type
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 422
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 366:
    Kernel Items:
      array_type: LBRACKET value_type COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 423

  State 367:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expression., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 368:
    Kernel Items:
      colon_expr: colon_expr COLON expression., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 369:
    Kernel Items:
      colon_expr: expression COLON expression., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 370:
    Kernel Items:
      proper_arguments: proper_arguments COMMA argument., *
    Reduce:
      * -> [proper_arguments]
    Goto:
      (nil)

  State 371:
    Kernel Items:
      explicit_struct_def: STRUCT LPAREN explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_def]
    Goto:
      (nil)

  State 372:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs COMMA.field_def
      explicit_field_defs: proper_explicit_field_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 229
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 235
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 236
      trait_op_type -> State 129
      field_def -> State 424
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 373:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs NEWLINES.field_def
      explicit_field_defs: proper_explicit_field_defs NEWLINES., RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 229
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 235
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 236
      trait_op_type -> State 129
      field_def -> State 425
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 374:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 426

  State 375:
    Kernel Items:
      optional_generic_binding: DOLLAR_LBRACKET generic_arguments RBRACKET., *
    Reduce:
      * -> [optional_generic_binding]
    Goto:
      (nil)

  State 376:
    Kernel Items:
      proper_generic_arguments: proper_generic_arguments COMMA.value_type
      generic_arguments: proper_generic_arguments COMMA., RBRACKET
    Reduce:
      RBRACKET -> [generic_arguments]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 427
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 377:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [index_expr]
    Goto:
      (nil)

  State 378:
    Kernel Items:
      call_expr: accessible_expr optional_generic_binding LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 428

  State 379:
    Kernel Items:
      initialize_expr: initializable_type LPAREN arguments RPAREN., *
    Reduce:
      * -> [initialize_expr]
    Goto:
      (nil)

  State 380:
    Kernel Items:
      loop_expr: DO statement_block FOR.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      optional_label_decl -> State 148
      sequence_expr -> State 429
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 381:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      optional_label_decl -> State 148
      sequence_expr -> State 430
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 382:
    Kernel Items:
      loop_expr: FOR assign_pattern IN.sequence_expr DO statement_block
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      optional_label_decl -> State 148
      sequence_expr -> State 431
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 383:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON.optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      LBRACE -> [optional_label_decl]
      SEMICOLON -> [optional_sequence_expr]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      optional_label_decl -> State 148
      sequence_expr -> State 433
      optional_sequence_expr -> State 432
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 384:
    Kernel Items:
      loop_expr: FOR sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 434

  State 385:
    Kernel Items:
      case_patterns: case_patterns.COMMA case_pattern
      condition: CASE case_patterns.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      COMMA -> State 265
      ASSIGN -> State 435

  State 386:
    Kernel Items:
      if_expr: IF condition statement_block., *
      if_expr: IF condition statement_block.ELSE statement_block
      if_expr: IF condition statement_block.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 436

  State 387:
    Kernel Items:
      switch_expr: SWITCH sequence_expr statement_block., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 388:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 229
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 235
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 236
      trait_op_type -> State 129
      field_def -> State 315
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      enum_value_def -> State 437
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 389:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR.enum_value_def
      explicit_enum_value_defs: enum_value_def OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 229
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 235
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 236
      trait_op_type -> State 129
      field_def -> State 315
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      enum_value_def -> State 438
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 390:
    Kernel Items:
      explicit_enum_def: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_def]
    Goto:
      (nil)

  State 391:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 229
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 235
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 236
      trait_op_type -> State 129
      field_def -> State 315
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      enum_value_def -> State 439
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 392:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
      explicit_enum_value_defs: implicit_enum_value_defs OR.enum_value_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 229
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 235
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 236
      trait_op_type -> State 129
      field_def -> State 315
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      enum_value_def -> State 440
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 393:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_decl: ELLIPSIS value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 394:
    Kernel Items:
      parameter_decl: IDENTIFIER ELLIPSIS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 441
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 395:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_decl: IDENTIFIER value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 396:
    Kernel Items:
      func_type: FUNC LPAREN parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 418
      prefixed_type -> State 126
      prefix_type_op -> State 125
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      return_type -> State 442
      func_type -> State 120

  State 397:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls COMMA.parameter_decl
      parameter_decls: proper_parameter_decls COMMA., RPAREN
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 318
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      ELLIPSIS -> State 317
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 322
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      parameter_decl -> State 443
      func_type -> State 120

  State 398:
    Kernel Items:
      atom_type: IDENTIFIER DOT IDENTIFIER optional_generic_binding., *
    Reduce:
      * -> [atom_type]
    Goto:
      (nil)

  State 399:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 400:
    Kernel Items:
      enum_value_def: field_def ASSIGN DEFAULT., *
    Reduce:
      * -> [enum_value_def]
    Goto:
      (nil)

  State 401:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
    Reduce:
      * -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 402:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [proper_implicit_field_defs]
    Goto:
      (nil)

  State 403:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 444

  State 404:
    Kernel Items:
      proper_trait_properties: proper_trait_properties COMMA.trait_property
      trait_properties: proper_trait_properties COMMA., RPAREN
    Reduce:
      RPAREN -> [trait_properties]
    Goto:
      IDENTIFIER -> State 229
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 332
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 235
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 236
      trait_op_type -> State 129
      field_def -> State 333
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_property -> State 445
      trait_def -> State 128
      func_type -> State 120
      method_signature -> State 334

  State 405:
    Kernel Items:
      proper_trait_properties: proper_trait_properties NEWLINES.trait_property
      trait_properties: proper_trait_properties NEWLINES., RPAREN
    Reduce:
      RPAREN -> [trait_properties]
    Goto:
      IDENTIFIER -> State 229
      UNSAFE -> State 53
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 332
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      unsafe_statement -> State 235
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 236
      trait_op_type -> State 129
      field_def -> State 333
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_property -> State 446
      trait_def -> State 128
      func_type -> State 120
      method_signature -> State 334

  State 406:
    Kernel Items:
      trait_def: TRAIT LPAREN trait_properties RPAREN., *
    Reduce:
      * -> [trait_def]
    Goto:
      (nil)

  State 407:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      generic_parameter_def: IDENTIFIER value_type., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 408:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 409:
    Kernel Items:
      proper_generic_parameter_defs: proper_generic_parameter_defs COMMA.generic_parameter_def
      generic_parameter_defs: proper_generic_parameter_defs COMMA., RBRACKET
    Reduce:
      RBRACKET -> [generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 340
      generic_parameter_def -> State 447

  State 410:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 448

  State 411:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_def: IDENTIFIER ELLIPSIS value_type., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 412:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER.optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 244
      optional_generic_parameters -> State 449

  State 413:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS.value_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 450
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      func_type -> State 120

  State 414:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 415:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 416:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER tuple_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 417:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 451

  State 418:
    Kernel Items:
      return_type: returnable_type., *
    Reduce:
      * -> [return_type]
    Goto:
      (nil)

  State 419:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [proper_parameter_defs]
    Goto:
      (nil)

  State 420:
    Kernel Items:
      proper_import_clauses: proper_import_clauses COMMA import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 421:
    Kernel Items:
      proper_import_clauses: proper_import_clauses NEWLINES import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 422:
    Kernel Items:
      map_type: LBRACKET value_type COLON value_type RBRACKET., *
    Reduce:
      * -> [map_type]
    Goto:
      (nil)

  State 423:
    Kernel Items:
      array_type: LBRACKET value_type COMMA INTEGER_LITERAL RBRACKET., *
    Reduce:
      * -> [array_type]
    Goto:
      (nil)

  State 424:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 425:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 426:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., $
    Reduce:
      $ -> [unsafe_statement]
    Goto:
      (nil)

  State 427:
    Kernel Items:
      proper_generic_arguments: proper_generic_arguments COMMA value_type., *
      trait_op_type: value_type.trait_op returnable_type
    Reduce:
      * -> [proper_generic_arguments]
    Goto:
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 428:
    Kernel Items:
      call_expr: accessible_expr optional_generic_binding LPAREN arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 429:
    Kernel Items:
      loop_expr: DO statement_block FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 430:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN sequence_expr., SEMICOLON
    Reduce:
      SEMICOLON -> [for_assignment]
    Goto:
      (nil)

  State 431:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 452

  State 432:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 453

  State 433:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 434:
    Kernel Items:
      loop_expr: FOR sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 435:
    Kernel Items:
      condition: CASE case_patterns ASSIGN.sequence_expr
    Reduce:
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      optional_label_decl -> State 148
      sequence_expr -> State 454
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 436:
    Kernel Items:
      if_expr: IF condition statement_block ELSE.statement_block
      if_expr: IF condition statement_block ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 216
      LBRACE -> State 11
      statement_block -> State 456
      if_expr -> State 455

  State 437:
    Kernel Items:
      explicit_enum_value_defs: enum_value_def NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 438:
    Kernel Items:
      implicit_enum_value_defs: enum_value_def OR enum_value_def., *
      explicit_enum_value_defs: enum_value_def OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 439:
    Kernel Items:
      explicit_enum_value_defs: implicit_enum_value_defs NEWLINES enum_value_def., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 440:
    Kernel Items:
      implicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., *
      explicit_enum_value_defs: implicit_enum_value_defs OR enum_value_def., RPAREN
    Reduce:
      * -> [implicit_enum_value_defs]
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      (nil)

  State 441:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      parameter_decl: IDENTIFIER ELLIPSIS value_type., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 442:
    Kernel Items:
      func_type: FUNC LPAREN parameter_decls RPAREN return_type., *
    Reduce:
      * -> [func_type]
    Goto:
      (nil)

  State 443:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [proper_parameter_decls]
    Goto:
      (nil)

  State 444:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 318
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      ELLIPSIS -> State 317
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 127
      prefixed_type -> State 126
      prefix_type_op -> State 125
      value_type -> State 322
      trait_op_type -> State 129
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      parameter_decl -> State 319
      proper_parameter_decls -> State 321
      parameter_decls -> State 457
      func_type -> State 120

  State 445:
    Kernel Items:
      proper_trait_properties: proper_trait_properties COMMA trait_property., *
    Reduce:
      * -> [proper_trait_properties]
    Goto:
      (nil)

  State 446:
    Kernel Items:
      proper_trait_properties: proper_trait_properties NEWLINES trait_property., *
    Reduce:
      * -> [proper_trait_properties]
    Goto:
      (nil)

  State 447:
    Kernel Items:
      proper_generic_parameter_defs: proper_generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [proper_generic_parameter_defs]
    Goto:
      (nil)

  State 448:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 418
      prefixed_type -> State 126
      prefix_type_op -> State 125
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      return_type -> State 458
      func_type -> State 120

  State 449:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 459

  State 450:
    Kernel Items:
      trait_op_type: value_type.trait_op returnable_type
      type_def: TYPE IDENTIFIER optional_generic_parameters value_type IMPLEMENTS value_type., *
    Reduce:
      * -> [type_def]
    Goto:
      ADD -> State 239
      SUB -> State 241
      MUL -> State 240
      trait_op -> State 242

  State 451:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 452:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 460

  State 453:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON.optional_sequence_expr DO statement_block
    Reduce:
      DO -> [optional_sequence_expr]
      LBRACE -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 103
      var_or_let -> State 24
      optional_label_decl -> State 148
      sequence_expr -> State 433
      optional_sequence_expr -> State 461
      call_expr -> State 70
      atom_expr -> State 62
      parse_error_expr -> State 92
      literal_expr -> State 87
      identifier_expr -> State 79
      block_expr -> State 69
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 104
      access_expr -> State 54
      index_expr -> State 82
      postfixable_expr -> State 94
      postfix_unary_expr -> State 93
      prefixable_expr -> State 97
      prefix_unary_expr -> State 95
      prefix_unary_op -> State 96
      mul_expr -> State 89
      binary_mul_expr -> State 66
      add_expr -> State 56
      binary_add_expr -> State 63
      cmp_expr -> State 73
      binary_cmp_expr -> State 65
      and_expr -> State 57
      binary_and_expr -> State 64
      or_expr -> State 91
      binary_or_expr -> State 68
      initializable_type -> State 83
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      explicit_struct_def -> State 74
      anonymous_func_expr -> State 58

  State 454:
    Kernel Items:
      condition: CASE case_patterns ASSIGN sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 455:
    Kernel Items:
      if_expr: IF condition statement_block ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 456:
    Kernel Items:
      if_expr: IF condition statement_block ELSE statement_block., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 457:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 462

  State 458:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 463

  State 459:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 246
      parameter_def -> State 268
      proper_parameter_defs -> State 270
      parameter_defs -> State 464

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
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 418
      prefixed_type -> State 126
      prefix_type_op -> State 125
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      return_type -> State 466
      func_type -> State 120

  State 463:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [named_func_def]
    Goto:
      (nil)

  State 464:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 467

  State 465:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 468

  State 466:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 467:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 112
      STRUCT -> State 50
      ENUM -> State 109
      TRAIT -> State 117
      FUNC -> State 111
      LPAREN -> State 113
      LBRACKET -> State 42
      DOT -> State 108
      QUESTION -> State 115
      EXCLAIM -> State 110
      TILDE_TILDE -> State 116
      BIT_NEG -> State 107
      BIT_AND -> State 106
      PARSE_ERROR -> State 114
      initializable_type -> State 123
      slice_type -> State 100
      array_type -> State 59
      map_type -> State 88
      atom_type -> State 118
      parse_error_type -> State 124
      returnable_type -> State 418
      prefixed_type -> State 126
      prefix_type_op -> State 125
      implicit_struct_def -> State 122
      explicit_struct_def -> State 74
      implicit_enum_def -> State 121
      explicit_enum_def -> State 119
      trait_def -> State 128
      return_type -> State 469
      func_type -> State 120

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
      LBRACE -> State 11
      statement_block -> State 470

  State 470:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [named_func_def]
    Goto:
      (nil)

Number of states: 470
Number of shift actions: 4298
Number of reduce actions: 420
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 4727
Number of unoptimized shift actions: 41252
Number of unoptimized reduce actions: 31824
*/
