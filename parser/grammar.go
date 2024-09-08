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
	UnderscoreToken      = SymbolId(263)
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
	FallthroughToken     = SymbolId(277)
	PackageToken         = SymbolId(278)
	ImportToken          = SymbolId(279)
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
	// 59:29: source -> ...
	ToSource(Definitions_ []SourceDefinition) ([]SourceDefinition, error)
}

type ProperDefinitionsReducer interface {
	// 62:2: proper_definitions -> add: ...
	AddToProperDefinitions(ProperDefinitions_ []SourceDefinition, Newlines_ TokenCount, Definition_ SourceDefinition) ([]SourceDefinition, error)

	// 63:2: proper_definitions -> definition: ...
	DefinitionToProperDefinitions(Definition_ SourceDefinition) ([]SourceDefinition, error)
}

type DefinitionsReducer interface {

	// 67:2: definitions -> improper: ...
	ImproperToDefinitions(ProperDefinitions_ []SourceDefinition, Newlines_ TokenCount) ([]SourceDefinition, error)

	// 68:2: definitions -> nil: ...
	NilToDefinitions() ([]SourceDefinition, error)
}

type DefinitionReducer interface {

	// 75:2: definition -> global_var_def: ...
	GlobalVarDefToDefinition(VarDeclPattern_ GenericSymbol) (SourceDefinition, error)

	// 76:2: definition -> global_var_assignment: ...
	GlobalVarAssignmentToDefinition(VarDeclPattern_ GenericSymbol, Assign_ TokenValue, Expr_ Expression) (SourceDefinition, error)

	// 79:2: definition -> statement_block: ...
	StatementBlockToDefinition(StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 83:2: definition -> COMMENT_GROUPS: ...
	CommentGroupsToDefinition(CommentGroups_ TokenValue) (SourceDefinition, error)
}

type StatementBlockReducer interface {
	// 103:19: statement_block -> ...
	ToStatementBlock(Lbrace_ TokenValue, Statements_ *StatementList, Rbrace_ TokenValue) (GenericSymbol, error)
}

type ProperStatementsReducer interface {
	// 106:2: proper_statements -> add_implicit: ...
	AddImplicitToProperStatements(ProperStatements_ *StatementList, Newlines_ TokenCount, Statement_ Statement) (*StatementList, error)

	// 107:2: proper_statements -> add_explicit: ...
	AddExplicitToProperStatements(ProperStatements_ *StatementList, Semicolon_ TokenValue, Statement_ Statement) (*StatementList, error)

	// 108:2: proper_statements -> statement: ...
	StatementToProperStatements(Statement_ Statement) (*StatementList, error)
}

type StatementsReducer interface {

	// 112:2: statements -> improper_implicit: ...
	ImproperImplicitToStatements(ProperStatements_ *StatementList, Newlines_ TokenCount) (*StatementList, error)

	// 113:2: statements -> improper_explicit: ...
	ImproperExplicitToStatements(ProperStatements_ *StatementList, Semicolon_ TokenValue) (*StatementList, error)

	// 114:2: statements -> nil: ...
	NilToStatements() (*StatementList, error)
}

type StatementReducer interface {

	// 152:2: statement -> case_branch_statement: ...
	CaseBranchStatementToStatement(Case_ TokenValue, CasePatterns_ GenericSymbol, Colon_ TokenValue, OptionalSimpleStatement_ Statement) (Statement, error)

	// 153:2: statement -> default_branch_statement: ...
	DefaultBranchStatementToStatement(Default_ TokenValue, Colon_ TokenValue, OptionalSimpleStatement_ Statement) (Statement, error)
}

type OptionalSimpleStatementReducer interface {

	// 157:2: optional_simple_statement -> nil: ...
	NilToOptionalSimpleStatement() (Statement, error)
}

type ExprOrImproperStructStatementReducer interface {
	// 163:48: expr_or_improper_struct_statement -> ...
	ToExprOrImproperStructStatement(Exprs_ GenericSymbol) (Statement, error)
}

type ExprsReducer interface {
	// 166:2: exprs -> expr: ...
	ExprToExprs(Expr_ Expression) (GenericSymbol, error)

	// 167:2: exprs -> add: ...
	AddToExprs(Exprs_ GenericSymbol, Comma_ TokenValue, Expr_ Expression) (GenericSymbol, error)
}

type CallbackOpStatementReducer interface {
	// 193:36: callback_op_statement -> ...
	ToCallbackOpStatement(CallbackOp_ TokenValue, CallExpr_ Expression) (Statement, error)
}

type UnsafeStatementReducer interface {
	// 201:31: unsafe_statement -> ...
	ToUnsafeStatement(Unsafe_ TokenValue, Less_ TokenValue, Identifier_ TokenValue, Greater_ TokenValue, StringLiteral_ TokenValue) (Statement, error)
}

type JumpStatementReducer interface {
	// 208:2: jump_statement -> unlabeled_no_value: ...
	UnlabeledNoValueToJumpStatement(JumpType_ TokenValue) (Statement, error)

	// 209:2: jump_statement -> unlabeled_valued: ...
	UnlabeledValuedToJumpStatement(JumpType_ TokenValue, Exprs_ GenericSymbol) (Statement, error)

	// 210:2: jump_statement -> labeled_no_value: ...
	LabeledNoValueToJumpStatement(JumpType_ TokenValue, JumpLabel_ TokenValue) (Statement, error)

	// 211:2: jump_statement -> labeled_valued: ...
	LabeledValuedToJumpStatement(JumpType_ TokenValue, JumpLabel_ TokenValue, Exprs_ GenericSymbol) (Statement, error)
}

type FallthroughStatementReducer interface {
	// 218:36: fallthrough_statement -> ...
	ToFallthroughStatement(Fallthrough_ TokenValue) (Statement, error)
}

type AssignStatementReducer interface {
	// 224:31: assign_statement -> ...
	ToAssignStatement(AssignPattern_ GenericSymbol, Assign_ TokenValue, Expr_ Expression) (Statement, error)
}

type UnaryOpAssignStatementReducer interface {
	// 226:40: unary_op_assign_statement -> ...
	ToUnaryOpAssignStatement(AccessibleExpr_ Expression, UnaryOpAssign_ TokenValue) (Statement, error)
}

type BinaryOpAssignStatementReducer interface {
	// 232:41: binary_op_assign_statement -> ...
	ToBinaryOpAssignStatement(AccessibleExpr_ Expression, BinaryOpAssign_ TokenValue, Expr_ Expression) (Statement, error)
}

type ImportStatementReducer interface {
	// 252:2: import_statement -> single: ...
	SingleToImportStatement(Import_ TokenValue, ImportClause_ GenericSymbol) (Statement, error)

	// 253:2: import_statement -> multiple: ...
	MultipleToImportStatement(Import_ TokenValue, Lparen_ TokenValue, ImportClauses_ GenericSymbol, Rparen_ TokenValue) (Statement, error)
}

type ImportClauseReducer interface {
	// 256:2: import_clause -> STRING_LITERAL: ...
	StringLiteralToImportClause(StringLiteral_ TokenValue) (GenericSymbol, error)

	// 257:2: import_clause -> alias: ...
	AliasToImportClause(Identifier_ TokenValue, StringLiteral_ TokenValue) (GenericSymbol, error)

	// 258:2: import_clause -> unusable_import: ...
	UnusableImportToImportClause(Underscore_ TokenValue, StringLiteral_ TokenValue) (GenericSymbol, error)

	// 259:2: import_clause -> import_to_local: ...
	ImportToLocalToImportClause(Dot_ TokenValue, StringLiteral_ TokenValue) (GenericSymbol, error)
}

type ProperImportClausesReducer interface {
	// 262:2: proper_import_clauses -> add_implicit: ...
	AddImplicitToProperImportClauses(ProperImportClauses_ GenericSymbol, Newlines_ TokenCount, ImportClause_ GenericSymbol) (GenericSymbol, error)

	// 263:2: proper_import_clauses -> add_explicit: ...
	AddExplicitToProperImportClauses(ProperImportClauses_ GenericSymbol, Comma_ TokenValue, ImportClause_ GenericSymbol) (GenericSymbol, error)

	// 264:2: proper_import_clauses -> import_clause: ...
	ImportClauseToProperImportClauses(ImportClause_ GenericSymbol) (GenericSymbol, error)
}

type ImportClausesReducer interface {

	// 268:2: import_clauses -> implicit: ...
	ImplicitToImportClauses(ProperImportClauses_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)

	// 269:2: import_clauses -> explicit: ...
	ExplicitToImportClauses(ProperImportClauses_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)
}

type CasePatternsReducer interface {
	// 276:2: case_patterns -> case_pattern: ...
	CasePatternToCasePatterns(CasePattern_ GenericSymbol) (GenericSymbol, error)

	// 277:2: case_patterns -> multiple: ...
	MultipleToCasePatterns(CasePatterns_ GenericSymbol, Comma_ TokenValue, CasePattern_ GenericSymbol) (GenericSymbol, error)
}

type VarDeclPatternReducer interface {
	// 286:20: var_decl_pattern -> ...
	ToVarDeclPattern(VarOrLet_ TokenValue, VarPattern_ GenericSymbol, OptionalTypeExpr_ TypeExpression) (GenericSymbol, error)
}

type VarPatternReducer interface {
	// 293:2: var_pattern -> IDENTIFIER: ...
	IdentifierToVarPattern(Identifier_ TokenValue) (GenericSymbol, error)

	// 294:2: var_pattern -> UNDERSCORE: ...
	UnderscoreToVarPattern(Underscore_ TokenValue) (GenericSymbol, error)

	// 295:2: var_pattern -> tuple_pattern: ...
	TuplePatternToVarPattern(TuplePattern_ GenericSymbol) (GenericSymbol, error)
}

type TuplePatternReducer interface {
	// 297:17: tuple_pattern -> ...
	ToTuplePattern(Lparen_ TokenValue, FieldVarPatterns_ GenericSymbol, Rparen_ TokenValue) (GenericSymbol, error)
}

type FieldVarPatternsReducer interface {
	// 300:2: field_var_patterns -> field_var_pattern: ...
	FieldVarPatternToFieldVarPatterns(FieldVarPattern_ GenericSymbol) (GenericSymbol, error)

	// 301:2: field_var_patterns -> add: ...
	AddToFieldVarPatterns(FieldVarPatterns_ GenericSymbol, Comma_ TokenValue, FieldVarPattern_ GenericSymbol) (GenericSymbol, error)
}

type FieldVarPatternReducer interface {
	// 304:2: field_var_pattern -> positional: ...
	PositionalToFieldVarPattern(VarPattern_ GenericSymbol) (GenericSymbol, error)

	// 305:2: field_var_pattern -> named: ...
	NamedToFieldVarPattern(Identifier_ TokenValue, Assign_ TokenValue, VarPattern_ GenericSymbol) (GenericSymbol, error)

	// 306:2: field_var_pattern -> ELLIPSIS: ...
	EllipsisToFieldVarPattern(Ellipsis_ TokenValue) (GenericSymbol, error)
}

type OptionalTypeExprReducer interface {

	// 310:2: optional_type_expr -> nil: ...
	NilToOptionalTypeExpr() (TypeExpression, error)
}

type AssignPatternReducer interface {
	// 317:2: assign_pattern -> ...
	ToAssignPattern(SequenceExpr_ Expression) (GenericSymbol, error)
}

type CasePatternReducer interface {

	// 327:2: case_pattern -> enum_match_pattern: ...
	EnumMatchPatternToCasePattern(Dot_ TokenValue, Identifier_ TokenValue, ImplicitStructExpr_ Expression) (GenericSymbol, error)

	// 328:2: case_pattern -> enum_nondata_match_patten: ...
	EnumNondataMatchPattenToCasePattern(Dot_ TokenValue, Identifier_ TokenValue) (GenericSymbol, error)

	// 329:2: case_pattern -> enum_var_decl_pattern: ...
	EnumVarDeclPatternToCasePattern(Var_ TokenValue, Dot_ TokenValue, Identifier_ TokenValue, TuplePattern_ GenericSymbol) (GenericSymbol, error)
}

type ExprReducer interface {
	// 336:2: expr -> if_expr: ...
	IfExprToExpr(OptionalLabelDecl_ TokenValue, IfExpr_ Expression) (Expression, error)

	// 337:2: expr -> switch_expr: ...
	SwitchExprToExpr(OptionalLabelDecl_ TokenValue, SwitchExpr_ Expression) (Expression, error)

	// 338:2: expr -> loop_expr: ...
	LoopExprToExpr(OptionalLabelDecl_ TokenValue, LoopExpr_ Expression) (Expression, error)
}

type OptionalLabelDeclReducer interface {
	// 342:2: optional_label_decl -> LABEL_DECL: ...
	LabelDeclToOptionalLabelDecl(LabelDecl_ TokenValue) (TokenValue, error)

	// 343:2: optional_label_decl -> unlabelled: ...
	UnlabelledToOptionalLabelDecl() (TokenValue, error)
}

type SequenceExprReducer interface {

	// 353:2: sequence_expr -> var_decl_pattern: ...
	VarDeclPatternToSequenceExpr(VarDeclPattern_ GenericSymbol) (Expression, error)

	// 357:2: sequence_expr -> assign_var_pattern: ...
	AssignVarPatternToSequenceExpr(Greater_ TokenValue, SequenceExpr_ Expression) (Expression, error)
}

type IfExprReducer interface {
	// 366:2: if_expr -> no_else: ...
	NoElseToIfExpr(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol) (Expression, error)

	// 367:2: if_expr -> if_else: ...
	IfElseToIfExpr(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol, Else_ TokenValue, StatementBlock_2 GenericSymbol) (Expression, error)

	// 368:2: if_expr -> multi_if_else: ...
	MultiIfElseToIfExpr(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol, Else_ TokenValue, IfExpr_ Expression) (Expression, error)
}

type ConditionReducer interface {
	// 371:2: condition -> sequence_expr: ...
	SequenceExprToCondition(SequenceExpr_ Expression) (GenericSymbol, error)

	// 372:2: condition -> case: ...
	CaseToCondition(Case_ TokenValue, CasePatterns_ GenericSymbol, Assign_ TokenValue, SequenceExpr_ Expression) (GenericSymbol, error)
}

type SwitchExprReducer interface {
	// 396:27: switch_expr -> ...
	ToSwitchExpr(Switch_ TokenValue, SequenceExpr_ Expression, StatementBlock_ GenericSymbol) (Expression, error)
}

type LoopExprReducer interface {
	// 410:2: loop_expr -> infinite: ...
	InfiniteToLoopExpr(Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 411:2: loop_expr -> do_while: ...
	DoWhileToLoopExpr(Do_ TokenValue, StatementBlock_ GenericSymbol, For_ TokenValue, SequenceExpr_ Expression) (Expression, error)

	// 412:2: loop_expr -> while: ...
	WhileToLoopExpr(For_ TokenValue, SequenceExpr_ Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 413:2: loop_expr -> iterator: ...
	IteratorToLoopExpr(For_ TokenValue, AssignPattern_ GenericSymbol, In_ TokenValue, SequenceExpr_ Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 414:2: loop_expr -> for: ...
	ForToLoopExpr(For_ TokenValue, ForAssignment_ GenericSymbol, Semicolon_ TokenValue, OptionalSequenceExpr_ Expression, Semicolon_2 TokenValue, OptionalSequenceExpr_2 Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)
}

type OptionalSequenceExprReducer interface {
	// 417:2: optional_sequence_expr -> sequence_expr: ...
	SequenceExprToOptionalSequenceExpr(SequenceExpr_ Expression) (Expression, error)

	// 418:2: optional_sequence_expr -> nil: ...
	NilToOptionalSequenceExpr() (Expression, error)
}

type ForAssignmentReducer interface {
	// 421:2: for_assignment -> sequence_expr: ...
	SequenceExprToForAssignment(SequenceExpr_ Expression) (GenericSymbol, error)

	// 422:2: for_assignment -> assign: ...
	AssignToForAssignment(AssignPattern_ GenericSymbol, Assign_ TokenValue, SequenceExpr_ Expression) (GenericSymbol, error)
}

type CallExprReducer interface {
	// 429:2: call_expr -> ...
	ToCallExpr(AccessibleExpr_ Expression, GenericTypeArguments_ *TypeArgumentList, Lparen_ TokenValue, Arguments_ *ArgumentList, Rparen_ TokenValue) (Expression, error)
}

type ProperArgumentsReducer interface {
	// 432:2: proper_arguments -> add: ...
	AddToProperArguments(ProperArguments_ *ArgumentList, Comma_ TokenValue, Argument_ *Argument) (*ArgumentList, error)

	// 433:2: proper_arguments -> argument: ...
	ArgumentToProperArguments(Argument_ *Argument) (*ArgumentList, error)
}

type ArgumentsReducer interface {

	// 437:2: arguments -> improper: ...
	ImproperToArguments(ProperArguments_ *ArgumentList, Comma_ TokenValue) (*ArgumentList, error)

	// 438:2: arguments -> nil: ...
	NilToArguments() (*ArgumentList, error)
}

type ArgumentReducer interface {
	// 441:2: argument -> positional: ...
	PositionalToArgument(Expr_ Expression) (*Argument, error)

	// 442:2: argument -> colon_expr: ...
	ColonExprToArgument(ColonExpr_ *ColonExpr) (*Argument, error)

	// 443:2: argument -> named_assignment: ...
	NamedAssignmentToArgument(Identifier_ TokenValue, Assign_ TokenValue, Expr_ Expression) (*Argument, error)

	// 447:2: argument -> vararg_assignment: ...
	VarargAssignmentToArgument(Expr_ Expression, Ellipsis_ TokenValue) (*Argument, error)

	// 450:2: argument -> skip_pattern: ...
	SkipPatternToArgument(Ellipsis_ TokenValue) (*Argument, error)
}

type ColonExprReducer interface {
	// 454:2: colon_expr -> unit_unit_pair: ...
	UnitUnitPairToColonExpr(Colon_ TokenValue) (*ColonExpr, error)

	// 455:2: colon_expr -> expr_unit_pair: ...
	ExprUnitPairToColonExpr(Expr_ Expression, Colon_ TokenValue) (*ColonExpr, error)

	// 456:2: colon_expr -> unit_expr_pair: ...
	UnitExprPairToColonExpr(Colon_ TokenValue, Expr_ Expression) (*ColonExpr, error)

	// 457:2: colon_expr -> expr_expr_pair: ...
	ExprExprPairToColonExpr(Expr_ Expression, Colon_ TokenValue, Expr_2 Expression) (*ColonExpr, error)

	// 458:2: colon_expr -> colon_expr_unit_tuple: ...
	ColonExprUnitTupleToColonExpr(ColonExpr_ *ColonExpr, Colon_ TokenValue) (*ColonExpr, error)

	// 459:2: colon_expr -> colon_expr_expr_tuple: ...
	ColonExprExprTupleToColonExpr(ColonExpr_ *ColonExpr, Colon_ TokenValue, Expr_ Expression) (*ColonExpr, error)
}

type ParseErrorExprReducer interface {
	// 478:32: parse_error_expr -> ...
	ToParseErrorExpr(ParseError_ ParseErrorSymbol) (Expression, error)
}

type LiteralExprReducer interface {
	// 481:2: literal_expr -> TRUE: ...
	TrueToLiteralExpr(True_ TokenValue) (Expression, error)

	// 482:2: literal_expr -> FALSE: ...
	FalseToLiteralExpr(False_ TokenValue) (Expression, error)

	// 483:2: literal_expr -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteralExpr(IntegerLiteral_ TokenValue) (Expression, error)

	// 484:2: literal_expr -> FLOAT_LITERAL: ...
	FloatLiteralToLiteralExpr(FloatLiteral_ TokenValue) (Expression, error)

	// 485:2: literal_expr -> RUNE_LITERAL: ...
	RuneLiteralToLiteralExpr(RuneLiteral_ TokenValue) (Expression, error)

	// 486:2: literal_expr -> STRING_LITERAL: ...
	StringLiteralToLiteralExpr(StringLiteral_ TokenValue) (Expression, error)
}

type NamedExprReducer interface {
	// 489:2: named_expr -> IDENTIFIER: ...
	IdentifierToNamedExpr(Identifier_ TokenValue) (Expression, error)

	// 490:2: named_expr -> UNDERSCORE: ...
	UnderscoreToNamedExpr(Underscore_ TokenValue) (Expression, error)
}

type BlockExprReducer interface {
	// 492:27: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)
}

type InitializeExprReducer interface {
	// 494:31: initialize_expr -> ...
	ToInitializeExpr(InitializableTypeExpr_ TypeExpression, Lparen_ TokenValue, Arguments_ *ArgumentList, Rparen_ TokenValue) (Expression, error)
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

type SliceTypeExprReducer interface {
	// 598:35: slice_type_expr -> ...
	ToSliceTypeExpr(Lbracket_ TokenValue, TypeExpr_ TypeExpression, Rbracket_ TokenValue) (TypeExpression, error)
}

type ArrayTypeExprReducer interface {
	// 601:2: array_type_expr -> ...
	ToArrayTypeExpr(Lbracket_ TokenValue, TypeExpr_ TypeExpression, Comma_ TokenValue, IntegerLiteral_ TokenValue, Rbracket_ TokenValue) (TypeExpression, error)
}

type MapTypeExprReducer interface {
	// 604:33: map_type_expr -> ...
	ToMapTypeExpr(Lbracket_ TokenValue, TypeExpr_ TypeExpression, Colon_ TokenValue, TypeExpr_2 TypeExpression, Rbracket_ TokenValue) (TypeExpression, error)
}

type NamedTypeExprReducer interface {
	// 618:2: named_type_expr -> local: ...
	LocalToNamedTypeExpr(Identifier_ TokenValue, GenericTypeArguments_ *TypeArgumentList) (TypeExpression, error)

	// 619:2: named_type_expr -> external: ...
	ExternalToNamedTypeExpr(Identifier_ TokenValue, Dot_ TokenValue, Identifier_2 TokenValue, GenericTypeArguments_ *TypeArgumentList) (TypeExpression, error)
}

type InferredTypeExprReducer interface {
	// 627:2: inferred_type_expr -> DOT: ...
	DotToInferredTypeExpr(Dot_ TokenValue) (TypeExpression, error)

	// 628:2: inferred_type_expr -> UNDERSCORE: ...
	UnderscoreToInferredTypeExpr(Underscore_ TokenValue) (TypeExpression, error)
}

type ParseErrorTypeExprReducer interface {
	// 630:41: parse_error_type_expr -> ...
	ToParseErrorTypeExpr(ParseError_ ParseErrorSymbol) (TypeExpression, error)
}

type PrefixUnaryTypeExprReducer interface {
	// 640:2: prefix_unary_type_expr -> ...
	ToPrefixUnaryTypeExpr(PrefixUnaryTypeOp_ TokenValue, ReturnableTypeExpr_ TypeExpression) (TypeExpression, error)
}

type BinaryTypeExprReducer interface {
	// 656:2: binary_type_expr -> ...
	ToBinaryTypeExpr(TypeExpr_ TypeExpression, BinaryTypeOp_ TokenValue, ReturnableTypeExpr_ TypeExpression) (TypeExpression, error)
}

type TypeDefReducer interface {
	// 664:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ TokenValue, Identifier_ TokenValue, GenericParameterDefs_ GenericSymbol, TypeExpr_ TypeExpression) (SourceDefinition, error)

	// 665:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ TokenValue, Identifier_ TokenValue, GenericParameterDefs_ GenericSymbol, TypeExpr_ TypeExpression, Implements_ TokenValue, TypeExpr_2 TypeExpression) (SourceDefinition, error)

	// 666:2: type_def -> alias: ...
	AliasToTypeDef(Type_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, TypeExpr_ TypeExpression) (SourceDefinition, error)
}

type GenericParameterDefReducer interface {
	// 674:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ TokenValue) (GenericSymbol, error)

	// 675:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)
}

type GenericParameterDefsReducer interface {
	// 678:2: generic_parameter_defs -> generic: ...
	GenericToGenericParameterDefs(DollarLbracket_ TokenValue, GenericParameterDefList_ GenericSymbol, Rbracket_ TokenValue) (GenericSymbol, error)

	// 679:2: generic_parameter_defs -> nil: ...
	NilToGenericParameterDefs() (GenericSymbol, error)
}

type ProperGenericParameterDefListReducer interface {
	// 682:2: proper_generic_parameter_def_list -> add: ...
	AddToProperGenericParameterDefList(ProperGenericParameterDefList_ GenericSymbol, Comma_ TokenValue, GenericParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 683:2: proper_generic_parameter_def_list -> generic_parameter_def: ...
	GenericParameterDefToProperGenericParameterDefList(GenericParameterDef_ GenericSymbol) (GenericSymbol, error)
}

type GenericParameterDefListReducer interface {

	// 687:2: generic_parameter_def_list -> improper: ...
	ImproperToGenericParameterDefList(ProperGenericParameterDefList_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 688:2: generic_parameter_def_list -> nil: ...
	NilToGenericParameterDefList() (GenericSymbol, error)
}

type GenericTypeArgumentsReducer interface {
	// 691:2: generic_type_arguments -> binding: ...
	BindingToGenericTypeArguments(DollarLbracket_ TokenValue, TypeArguments_ *TypeArgumentList, Rbracket_ TokenValue) (*TypeArgumentList, error)

	// 692:2: generic_type_arguments -> nil: ...
	NilToGenericTypeArguments() (*TypeArgumentList, error)
}

type ProperTypeArgumentsReducer interface {
	// 695:2: proper_type_arguments -> add: ...
	AddToProperTypeArguments(ProperTypeArguments_ *TypeArgumentList, Comma_ TokenValue, TypeExpr_ TypeExpression) (*TypeArgumentList, error)

	// 696:2: proper_type_arguments -> type_expr: ...
	TypeExprToProperTypeArguments(TypeExpr_ TypeExpression) (*TypeArgumentList, error)
}

type TypeArgumentsReducer interface {

	// 700:2: type_arguments -> improper: ...
	ImproperToTypeArguments(ProperTypeArguments_ *TypeArgumentList, Comma_ TokenValue) (*TypeArgumentList, error)

	// 701:2: type_arguments -> nil: ...
	NilToTypeArguments() (*TypeArgumentList, error)
}

type FieldDefReducer interface {
	// 715:2: field_def -> named: ...
	NamedToFieldDef(Identifier_ TokenValue, TypeExpr_ TypeExpression, OptionalDefault_ GenericSymbol) (GenericSymbol, error)

	// 716:2: field_def -> unnamed: ...
	UnnamedToFieldDef(TypeExpr_ TypeExpression, OptionalDefault_ GenericSymbol) (GenericSymbol, error)

	// 717:2: field_def -> struct_padding: ...
	StructPaddingToFieldDef(Underscore_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 719:2: field_def -> unsafe_statement: ...
	UnsafeStatementToFieldDef(UnsafeStatement_ Statement) (GenericSymbol, error)
}

type OptionalDefaultReducer interface {
	// 722:2: optional_default -> default: ...
	DefaultToOptionalDefault(Assign_ TokenValue, Default_ TokenValue) (GenericSymbol, error)

	// 723:2: optional_default -> nil: ...
	NilToOptionalDefault() (GenericSymbol, error)
}

type ProperImplicitFieldDefsReducer interface {
	// 726:2: proper_implicit_field_defs -> add: ...
	AddToProperImplicitFieldDefs(ProperImplicitFieldDefs_ GenericSymbol, Comma_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 727:2: proper_implicit_field_defs -> field_def: ...
	FieldDefToProperImplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type ImplicitFieldDefsReducer interface {

	// 731:2: implicit_field_defs -> improper: ...
	ImproperToImplicitFieldDefs(ProperImplicitFieldDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 732:2: implicit_field_defs -> nil: ...
	NilToImplicitFieldDefs() (GenericSymbol, error)
}

type ImplicitStructTypeExprReducer interface {
	// 735:2: implicit_struct_type_expr -> ...
	ToImplicitStructTypeExpr(Lparen_ TokenValue, ImplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ProperExplicitFieldDefsReducer interface {
	// 738:2: proper_explicit_field_defs -> add_implicit: ...
	AddImplicitToProperExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Newlines_ TokenCount, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 739:2: proper_explicit_field_defs -> add_explicit: ...
	AddExplicitToProperExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Comma_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 740:2: proper_explicit_field_defs -> field_def: ...
	FieldDefToProperExplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type ExplicitFieldDefsReducer interface {

	// 744:2: explicit_field_defs -> improper_implicit: ...
	ImproperImplicitToExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)

	// 745:2: explicit_field_defs -> improper_explicit: ...
	ImproperExplicitToExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 746:2: explicit_field_defs -> nil: ...
	NilToExplicitFieldDefs() (GenericSymbol, error)
}

type ExplicitStructTypeExprReducer interface {
	// 749:2: explicit_struct_type_expr -> ...
	ToExplicitStructTypeExpr(Struct_ TokenValue, Lparen_ TokenValue, ExplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type TraitTypeExprReducer interface {
	// 752:2: trait_type_expr -> ...
	ToTraitTypeExpr(Trait_ TokenValue, Lparen_ TokenValue, ExplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ProperImplicitEnumValueDefsReducer interface {
	// 763:2: proper_implicit_enum_value_defs -> pair: ...
	PairToProperImplicitEnumValueDefs(FieldDef_ GenericSymbol, Or_ TokenValue, FieldDef_2 GenericSymbol) (GenericSymbol, error)

	// 764:2: proper_implicit_enum_value_defs -> add: ...
	AddToProperImplicitEnumValueDefs(ProperImplicitEnumValueDefs_ GenericSymbol, Or_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type ImplicitEnumValueDefsReducer interface {

	// 769:2: implicit_enum_value_defs -> improper: ...
	ImproperToImplicitEnumValueDefs(ProperImplicitEnumValueDefs_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)
}

type ImplicitEnumTypeExprReducer interface {
	// 772:2: implicit_enum_type_expr -> ...
	ToImplicitEnumTypeExpr(Lparen_ TokenValue, ImplicitEnumValueDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ProperExplicitEnumValueDefsReducer interface {
	// 775:2: proper_explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToProperExplicitEnumValueDefs(FieldDef_ GenericSymbol, Or_ TokenValue, FieldDef_2 GenericSymbol) (GenericSymbol, error)

	// 776:2: proper_explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToProperExplicitEnumValueDefs(FieldDef_ GenericSymbol, Newlines_ TokenCount, FieldDef_2 GenericSymbol) (GenericSymbol, error)

	// 777:2: proper_explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToProperExplicitEnumValueDefs(ProperExplicitEnumValueDefs_ GenericSymbol, Or_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 778:2: proper_explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToProperExplicitEnumValueDefs(ProperExplicitEnumValueDefs_ GenericSymbol, Newlines_ TokenCount, FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type ExplicitEnumValueDefsReducer interface {

	// 783:2: explicit_enum_value_defs -> improper: ...
	ImproperToExplicitEnumValueDefs(ProperExplicitEnumValueDefs_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)
}

type ExplicitEnumTypeExprReducer interface {
	// 786:2: explicit_enum_type_expr -> ...
	ToExplicitEnumTypeExpr(Enum_ TokenValue, Lparen_ TokenValue, ExplicitEnumValueDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ReturnTypeReducer interface {
	// 794:2: return_type -> returnable_type_expr: ...
	ReturnableTypeExprToReturnType(ReturnableTypeExpr_ TypeExpression) (TypeExpression, error)

	// 795:2: return_type -> nil: ...
	NilToReturnType() (TypeExpression, error)
}

type ProperParameterDefReducer interface {
	// 798:2: proper_parameter_def -> named_typed_arg: ...
	NamedTypedArgToProperParameterDef(Identifier_ TokenValue, TypeExpr_ TypeExpression) (*Parameter, error)

	// 799:2: proper_parameter_def -> named_typed_vararg: ...
	NamedTypedVarargToProperParameterDef(Identifier_ TokenValue, Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (*Parameter, error)

	// 800:2: proper_parameter_def -> named_inferred_vararg: ...
	NamedInferredVarargToProperParameterDef(Identifier_ TokenValue, Ellipsis_ TokenValue) (*Parameter, error)

	// 801:2: proper_parameter_def -> ignore_typed_arg: ...
	IgnoreTypedArgToProperParameterDef(Underscore_ TokenValue, TypeExpr_ TypeExpression) (*Parameter, error)

	// 802:2: proper_parameter_def -> ignore_inferred_vararg: ...
	IgnoreInferredVarargToProperParameterDef(Underscore_ TokenValue, Ellipsis_ TokenValue) (*Parameter, error)

	// 803:2: proper_parameter_def -> ignore_typed_vararg: ...
	IgnoreTypedVarargToProperParameterDef(Underscore_ TokenValue, Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (*Parameter, error)
}

type ParameterDeclReducer interface {

	// 809:2: parameter_decl -> unnamed_typed_arg: ...
	UnnamedTypedArgToParameterDecl(TypeExpr_ TypeExpression) (*Parameter, error)

	// 810:2: parameter_decl -> unnamed_inferred_vararg: ...
	UnnamedInferredVarargToParameterDecl(Ellipsis_ TokenValue) (*Parameter, error)

	// 811:2: parameter_decl -> unnamed_typed_vararg: ...
	UnnamedTypedVarargToParameterDecl(Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (*Parameter, error)
}

type ParameterDefReducer interface {

	// 819:2: parameter_def -> named_inferred_arg: ...
	NamedInferredArgToParameterDef(Identifier_ TokenValue) (*Parameter, error)

	// 820:2: parameter_def -> ignore_inferred_arg: ...
	IgnoreInferredArgToParameterDef(Underscore_ TokenValue) (*Parameter, error)
}

type ProperParameterDeclsReducer interface {
	// 823:2: proper_parameter_decls -> add: ...
	AddToProperParameterDecls(ProperParameterDecls_ *ParameterList, Comma_ TokenValue, ParameterDecl_ *Parameter) (*ParameterList, error)

	// 824:2: proper_parameter_decls -> parameter_decl: ...
	ParameterDeclToProperParameterDecls(ParameterDecl_ *Parameter) (*ParameterList, error)
}

type ParameterDeclsReducer interface {

	// 828:2: parameter_decls -> improper: ...
	ImproperToParameterDecls(ProperParameterDecls_ *ParameterList, Comma_ TokenValue) (*ParameterList, error)

	// 829:2: parameter_decls -> nil: ...
	NilToParameterDecls() (*ParameterList, error)
}

type ProperParameterDefsReducer interface {
	// 832:2: proper_parameter_defs -> add: ...
	AddToProperParameterDefs(ProperParameterDefs_ *ParameterList, Comma_ TokenValue, ParameterDef_ *Parameter) (*ParameterList, error)

	// 833:2: proper_parameter_defs -> parameter_def: ...
	ParameterDefToProperParameterDefs(ParameterDef_ *Parameter) (*ParameterList, error)
}

type ParameterDefsReducer interface {

	// 837:2: parameter_defs -> improper: ...
	ImproperToParameterDefs(ProperParameterDefs_ *ParameterList, Comma_ TokenValue) (*ParameterList, error)

	// 838:2: parameter_defs -> nil: ...
	NilToParameterDefs() (*ParameterList, error)
}

type FuncTypeExprReducer interface {
	// 841:2: func_type_expr -> ...
	ToFuncTypeExpr(Func_ TokenValue, Lparen_ TokenValue, ParameterDecls_ *ParameterList, Rparen_ TokenValue, ReturnType_ TypeExpression) (TypeExpression, error)
}

type MethodSignatureReducer interface {
	// 852:20: method_signature -> ...
	ToMethodSignature(Func_ TokenValue, Identifier_ TokenValue, Lparen_ TokenValue, ParameterDecls_ *ParameterList, Rparen_ TokenValue, ReturnType_ TypeExpression) (GenericSymbol, error)
}

type NamedFuncDefReducer interface {
	// 856:2: named_func_def -> func_def: ...
	FuncDefToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, GenericParameterDefs_ GenericSymbol, Lparen_ TokenValue, ParameterDefs_ *ParameterList, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 857:2: named_func_def -> method_def: ...
	MethodDefToNamedFuncDef(Func_ TokenValue, Lparen_ TokenValue, ParameterDef_ *Parameter, Rparen_ TokenValue, Identifier_ TokenValue, Lparen_2 TokenValue, ParameterDefs_ *ParameterList, Rparen_2 TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 858:2: named_func_def -> alias: ...
	AliasToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, Expr_ Expression) (SourceDefinition, error)
}

type AnonymousFuncExprReducer interface {
	// 861:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ TokenValue, Lparen_ TokenValue, ParameterDefs_ *ParameterList, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (Expression, error)
}

type PackageDefReducer interface {
	// 873:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ TokenValue) (SourceDefinition, error)

	// 874:2: package_def -> with_spec: ...
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
	ExprOrImproperStructStatementReducer
	ExprsReducer
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
	OptionalTypeExprReducer
	AssignPatternReducer
	CasePatternReducer
	ExprReducer
	OptionalLabelDeclReducer
	SequenceExprReducer
	IfExprReducer
	ConditionReducer
	SwitchExprReducer
	LoopExprReducer
	OptionalSequenceExprReducer
	ForAssignmentReducer
	CallExprReducer
	ProperArgumentsReducer
	ArgumentsReducer
	ArgumentReducer
	ColonExprReducer
	ParseErrorExprReducer
	LiteralExprReducer
	NamedExprReducer
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
	SliceTypeExprReducer
	ArrayTypeExprReducer
	MapTypeExprReducer
	NamedTypeExprReducer
	InferredTypeExprReducer
	ParseErrorTypeExprReducer
	PrefixUnaryTypeExprReducer
	BinaryTypeExprReducer
	TypeDefReducer
	GenericParameterDefReducer
	GenericParameterDefsReducer
	ProperGenericParameterDefListReducer
	GenericParameterDefListReducer
	GenericTypeArgumentsReducer
	ProperTypeArgumentsReducer
	TypeArgumentsReducer
	FieldDefReducer
	OptionalDefaultReducer
	ProperImplicitFieldDefsReducer
	ImplicitFieldDefsReducer
	ImplicitStructTypeExprReducer
	ProperExplicitFieldDefsReducer
	ExplicitFieldDefsReducer
	ExplicitStructTypeExprReducer
	TraitTypeExprReducer
	ProperImplicitEnumValueDefsReducer
	ImplicitEnumValueDefsReducer
	ImplicitEnumTypeExprReducer
	ProperExplicitEnumValueDefsReducer
	ExplicitEnumValueDefsReducer
	ExplicitEnumTypeExprReducer
	ReturnTypeReducer
	ProperParameterDefReducer
	ParameterDeclReducer
	ParameterDefReducer
	ProperParameterDeclsReducer
	ParameterDeclsReducer
	ProperParameterDefsReducer
	ParameterDefsReducer
	FuncTypeExprReducer
	MethodSignatureReducer
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

func ParseExpr(lexer Lexer, reducer Reducer) (Expression, error) {

	return ParseExprWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseExprWithCustomErrorHandler(
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

func ParseTypeExpr(lexer Lexer, reducer Reducer) (TypeExpression, error) {

	return ParseTypeExprWithCustomErrorHandler(
		lexer,
		reducer,
		DefaultParseErrorHandler{})
}

func ParseTypeExprWithCustomErrorHandler(
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
	case UnderscoreToken:
		return "UNDERSCORE"
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
	case ExprOrImproperStructStatementType:
		return "expr_or_improper_struct_statement"
	case ExprsType:
		return "exprs"
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
	case OptionalTypeExprType:
		return "optional_type_expr"
	case AssignPatternType:
		return "assign_pattern"
	case CasePatternType:
		return "case_pattern"
	case ExprType:
		return "expr"
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
	case NamedExprType:
		return "named_expr"
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
	case InitializableTypeExprType:
		return "initializable_type_expr"
	case SliceTypeExprType:
		return "slice_type_expr"
	case ArrayTypeExprType:
		return "array_type_expr"
	case MapTypeExprType:
		return "map_type_expr"
	case AtomTypeExprType:
		return "atom_type_expr"
	case NamedTypeExprType:
		return "named_type_expr"
	case InferredTypeExprType:
		return "inferred_type_expr"
	case ParseErrorTypeExprType:
		return "parse_error_type_expr"
	case ReturnableTypeExprType:
		return "returnable_type_expr"
	case PrefixUnaryTypeExprType:
		return "prefix_unary_type_expr"
	case PrefixUnaryTypeOpType:
		return "prefix_unary_type_op"
	case TypeExprType:
		return "type_expr"
	case BinaryTypeExprType:
		return "binary_type_expr"
	case BinaryTypeOpType:
		return "binary_type_op"
	case TypeDefType:
		return "type_def"
	case GenericParameterDefType:
		return "generic_parameter_def"
	case GenericParameterDefsType:
		return "generic_parameter_defs"
	case ProperGenericParameterDefListType:
		return "proper_generic_parameter_def_list"
	case GenericParameterDefListType:
		return "generic_parameter_def_list"
	case GenericTypeArgumentsType:
		return "generic_type_arguments"
	case ProperTypeArgumentsType:
		return "proper_type_arguments"
	case TypeArgumentsType:
		return "type_arguments"
	case FieldDefType:
		return "field_def"
	case OptionalDefaultType:
		return "optional_default"
	case ProperImplicitFieldDefsType:
		return "proper_implicit_field_defs"
	case ImplicitFieldDefsType:
		return "implicit_field_defs"
	case ImplicitStructTypeExprType:
		return "implicit_struct_type_expr"
	case ProperExplicitFieldDefsType:
		return "proper_explicit_field_defs"
	case ExplicitFieldDefsType:
		return "explicit_field_defs"
	case ExplicitStructTypeExprType:
		return "explicit_struct_type_expr"
	case TraitTypeExprType:
		return "trait_type_expr"
	case ProperImplicitEnumValueDefsType:
		return "proper_implicit_enum_value_defs"
	case ImplicitEnumValueDefsType:
		return "implicit_enum_value_defs"
	case ImplicitEnumTypeExprType:
		return "implicit_enum_type_expr"
	case ProperExplicitEnumValueDefsType:
		return "proper_explicit_enum_value_defs"
	case ExplicitEnumValueDefsType:
		return "explicit_enum_value_defs"
	case ExplicitEnumTypeExprType:
		return "explicit_enum_type_expr"
	case ReturnTypeType:
		return "return_type"
	case ProperParameterDefType:
		return "proper_parameter_def"
	case ParameterDeclType:
		return "parameter_decl"
	case ParameterDefType:
		return "parameter_def"
	case ProperParameterDeclsType:
		return "proper_parameter_decls"
	case ParameterDeclsType:
		return "parameter_decls"
	case ProperParameterDefsType:
		return "proper_parameter_defs"
	case ParameterDefsType:
		return "parameter_defs"
	case FuncTypeExprType:
		return "func_type_expr"
	case MethodSignatureType:
		return "method_signature"
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

	SourceType                        = SymbolId(343)
	ProperDefinitionsType             = SymbolId(344)
	DefinitionsType                   = SymbolId(345)
	DefinitionType                    = SymbolId(346)
	StatementBlockType                = SymbolId(347)
	ProperStatementsType              = SymbolId(348)
	StatementsType                    = SymbolId(349)
	SimpleStatementType               = SymbolId(350)
	StatementType                     = SymbolId(351)
	OptionalSimpleStatementType       = SymbolId(352)
	ExprOrImproperStructStatementType = SymbolId(353)
	ExprsType                         = SymbolId(354)
	CallbackOpType                    = SymbolId(355)
	CallbackOpStatementType           = SymbolId(356)
	UnsafeStatementType               = SymbolId(357)
	JumpStatementType                 = SymbolId(358)
	JumpTypeType                      = SymbolId(359)
	FallthroughStatementType          = SymbolId(360)
	AssignStatementType               = SymbolId(361)
	UnaryOpAssignStatementType        = SymbolId(362)
	UnaryOpAssignType                 = SymbolId(363)
	BinaryOpAssignStatementType       = SymbolId(364)
	BinaryOpAssignType                = SymbolId(365)
	ImportStatementType               = SymbolId(366)
	ImportClauseType                  = SymbolId(367)
	ProperImportClausesType           = SymbolId(368)
	ImportClausesType                 = SymbolId(369)
	CasePatternsType                  = SymbolId(370)
	VarDeclPatternType                = SymbolId(371)
	VarOrLetType                      = SymbolId(372)
	VarPatternType                    = SymbolId(373)
	TuplePatternType                  = SymbolId(374)
	FieldVarPatternsType              = SymbolId(375)
	FieldVarPatternType               = SymbolId(376)
	OptionalTypeExprType              = SymbolId(377)
	AssignPatternType                 = SymbolId(378)
	CasePatternType                   = SymbolId(379)
	ExprType                          = SymbolId(380)
	OptionalLabelDeclType             = SymbolId(381)
	SequenceExprType                  = SymbolId(382)
	IfExprType                        = SymbolId(383)
	ConditionType                     = SymbolId(384)
	SwitchExprType                    = SymbolId(385)
	LoopExprType                      = SymbolId(386)
	OptionalSequenceExprType          = SymbolId(387)
	ForAssignmentType                 = SymbolId(388)
	CallExprType                      = SymbolId(389)
	ProperArgumentsType               = SymbolId(390)
	ArgumentsType                     = SymbolId(391)
	ArgumentType                      = SymbolId(392)
	ColonExprType                     = SymbolId(393)
	AtomExprType                      = SymbolId(394)
	ParseErrorExprType                = SymbolId(395)
	LiteralExprType                   = SymbolId(396)
	NamedExprType                     = SymbolId(397)
	BlockExprType                     = SymbolId(398)
	InitializeExprType                = SymbolId(399)
	ImplicitStructExprType            = SymbolId(400)
	AccessibleExprType                = SymbolId(401)
	AccessExprType                    = SymbolId(402)
	IndexExprType                     = SymbolId(403)
	PostfixableExprType               = SymbolId(404)
	PostfixUnaryExprType              = SymbolId(405)
	PrefixableExprType                = SymbolId(406)
	PrefixUnaryExprType               = SymbolId(407)
	PrefixUnaryOpType                 = SymbolId(408)
	MulExprType                       = SymbolId(409)
	BinaryMulExprType                 = SymbolId(410)
	MulOpType                         = SymbolId(411)
	AddExprType                       = SymbolId(412)
	BinaryAddExprType                 = SymbolId(413)
	AddOpType                         = SymbolId(414)
	CmpExprType                       = SymbolId(415)
	BinaryCmpExprType                 = SymbolId(416)
	CmpOpType                         = SymbolId(417)
	AndExprType                       = SymbolId(418)
	BinaryAndExprType                 = SymbolId(419)
	OrExprType                        = SymbolId(420)
	BinaryOrExprType                  = SymbolId(421)
	InitializableTypeExprType         = SymbolId(422)
	SliceTypeExprType                 = SymbolId(423)
	ArrayTypeExprType                 = SymbolId(424)
	MapTypeExprType                   = SymbolId(425)
	AtomTypeExprType                  = SymbolId(426)
	NamedTypeExprType                 = SymbolId(427)
	InferredTypeExprType              = SymbolId(428)
	ParseErrorTypeExprType            = SymbolId(429)
	ReturnableTypeExprType            = SymbolId(430)
	PrefixUnaryTypeExprType           = SymbolId(431)
	PrefixUnaryTypeOpType             = SymbolId(432)
	TypeExprType                      = SymbolId(433)
	BinaryTypeExprType                = SymbolId(434)
	BinaryTypeOpType                  = SymbolId(435)
	TypeDefType                       = SymbolId(436)
	GenericParameterDefType           = SymbolId(437)
	GenericParameterDefsType          = SymbolId(438)
	ProperGenericParameterDefListType = SymbolId(439)
	GenericParameterDefListType       = SymbolId(440)
	GenericTypeArgumentsType          = SymbolId(441)
	ProperTypeArgumentsType           = SymbolId(442)
	TypeArgumentsType                 = SymbolId(443)
	FieldDefType                      = SymbolId(444)
	OptionalDefaultType               = SymbolId(445)
	ProperImplicitFieldDefsType       = SymbolId(446)
	ImplicitFieldDefsType             = SymbolId(447)
	ImplicitStructTypeExprType        = SymbolId(448)
	ProperExplicitFieldDefsType       = SymbolId(449)
	ExplicitFieldDefsType             = SymbolId(450)
	ExplicitStructTypeExprType        = SymbolId(451)
	TraitTypeExprType                 = SymbolId(452)
	ProperImplicitEnumValueDefsType   = SymbolId(453)
	ImplicitEnumValueDefsType         = SymbolId(454)
	ImplicitEnumTypeExprType          = SymbolId(455)
	ProperExplicitEnumValueDefsType   = SymbolId(456)
	ExplicitEnumValueDefsType         = SymbolId(457)
	ExplicitEnumTypeExprType          = SymbolId(458)
	ReturnTypeType                    = SymbolId(459)
	ProperParameterDefType            = SymbolId(460)
	ParameterDeclType                 = SymbolId(461)
	ParameterDefType                  = SymbolId(462)
	ProperParameterDeclsType          = SymbolId(463)
	ParameterDeclsType                = SymbolId(464)
	ProperParameterDefsType           = SymbolId(465)
	ParameterDefsType                 = SymbolId(466)
	FuncTypeExprType                  = SymbolId(467)
	MethodSignatureType               = SymbolId(468)
	NamedFuncDefType                  = SymbolId(469)
	AnonymousFuncExprType             = SymbolId(470)
	PackageDefType                    = SymbolId(471)
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
	_ReduceToSource                                               = _ReduceType(1)
	_ReduceAddToProperDefinitions                                 = _ReduceType(2)
	_ReduceDefinitionToProperDefinitions                          = _ReduceType(3)
	_ReduceProperDefinitionsToDefinitions                         = _ReduceType(4)
	_ReduceImproperToDefinitions                                  = _ReduceType(5)
	_ReduceNilToDefinitions                                       = _ReduceType(6)
	_ReducePackageDefToDefinition                                 = _ReduceType(7)
	_ReduceTypeDefToDefinition                                    = _ReduceType(8)
	_ReduceNamedFuncDefToDefinition                               = _ReduceType(9)
	_ReduceGlobalVarDefToDefinition                               = _ReduceType(10)
	_ReduceGlobalVarAssignmentToDefinition                        = _ReduceType(11)
	_ReduceStatementBlockToDefinition                             = _ReduceType(12)
	_ReduceCommentGroupsToDefinition                              = _ReduceType(13)
	_ReduceToStatementBlock                                       = _ReduceType(14)
	_ReduceAddImplicitToProperStatements                          = _ReduceType(15)
	_ReduceAddExplicitToProperStatements                          = _ReduceType(16)
	_ReduceStatementToProperStatements                            = _ReduceType(17)
	_ReduceProperStatementsToStatements                           = _ReduceType(18)
	_ReduceImproperImplicitToStatements                           = _ReduceType(19)
	_ReduceImproperExplicitToStatements                           = _ReduceType(20)
	_ReduceNilToStatements                                        = _ReduceType(21)
	_ReduceUnsafeStatementToSimpleStatement                       = _ReduceType(22)
	_ReduceExprOrImproperStructStatementToSimpleStatement         = _ReduceType(23)
	_ReduceCallbackOpStatementToSimpleStatement                   = _ReduceType(24)
	_ReduceJumpStatementToSimpleStatement                         = _ReduceType(25)
	_ReduceAssignStatementToSimpleStatement                       = _ReduceType(26)
	_ReduceUnaryOpAssignStatementToSimpleStatement                = _ReduceType(27)
	_ReduceBinaryOpAssignStatementToSimpleStatement               = _ReduceType(28)
	_ReduceFallthroughStatementToSimpleStatement                  = _ReduceType(29)
	_ReduceSimpleStatementToStatement                             = _ReduceType(30)
	_ReduceImportStatementToStatement                             = _ReduceType(31)
	_ReduceCaseBranchStatementToStatement                         = _ReduceType(32)
	_ReduceDefaultBranchStatementToStatement                      = _ReduceType(33)
	_ReduceSimpleStatementToOptionalSimpleStatement               = _ReduceType(34)
	_ReduceNilToOptionalSimpleStatement                           = _ReduceType(35)
	_ReduceToExprOrImproperStructStatement                        = _ReduceType(36)
	_ReduceExprToExprs                                            = _ReduceType(37)
	_ReduceAddToExprs                                             = _ReduceType(38)
	_ReduceAsyncToCallbackOp                                      = _ReduceType(39)
	_ReduceDeferToCallbackOp                                      = _ReduceType(40)
	_ReduceToCallbackOpStatement                                  = _ReduceType(41)
	_ReduceToUnsafeStatement                                      = _ReduceType(42)
	_ReduceUnlabeledNoValueToJumpStatement                        = _ReduceType(43)
	_ReduceUnlabeledValuedToJumpStatement                         = _ReduceType(44)
	_ReduceLabeledNoValueToJumpStatement                          = _ReduceType(45)
	_ReduceLabeledValuedToJumpStatement                           = _ReduceType(46)
	_ReduceReturnToJumpType                                       = _ReduceType(47)
	_ReduceBreakToJumpType                                        = _ReduceType(48)
	_ReduceContinueToJumpType                                     = _ReduceType(49)
	_ReduceToFallthroughStatement                                 = _ReduceType(50)
	_ReduceToAssignStatement                                      = _ReduceType(51)
	_ReduceToUnaryOpAssignStatement                               = _ReduceType(52)
	_ReduceAddOneAssignToUnaryOpAssign                            = _ReduceType(53)
	_ReduceSubOneAssignToUnaryOpAssign                            = _ReduceType(54)
	_ReduceToBinaryOpAssignStatement                              = _ReduceType(55)
	_ReduceAddAssignToBinaryOpAssign                              = _ReduceType(56)
	_ReduceSubAssignToBinaryOpAssign                              = _ReduceType(57)
	_ReduceMulAssignToBinaryOpAssign                              = _ReduceType(58)
	_ReduceDivAssignToBinaryOpAssign                              = _ReduceType(59)
	_ReduceModAssignToBinaryOpAssign                              = _ReduceType(60)
	_ReduceBitNegAssignToBinaryOpAssign                           = _ReduceType(61)
	_ReduceBitAndAssignToBinaryOpAssign                           = _ReduceType(62)
	_ReduceBitOrAssignToBinaryOpAssign                            = _ReduceType(63)
	_ReduceBitXorAssignToBinaryOpAssign                           = _ReduceType(64)
	_ReduceBitLshiftAssignToBinaryOpAssign                        = _ReduceType(65)
	_ReduceBitRshiftAssignToBinaryOpAssign                        = _ReduceType(66)
	_ReduceSingleToImportStatement                                = _ReduceType(67)
	_ReduceMultipleToImportStatement                              = _ReduceType(68)
	_ReduceStringLiteralToImportClause                            = _ReduceType(69)
	_ReduceAliasToImportClause                                    = _ReduceType(70)
	_ReduceUnusableImportToImportClause                           = _ReduceType(71)
	_ReduceImportToLocalToImportClause                            = _ReduceType(72)
	_ReduceAddImplicitToProperImportClauses                       = _ReduceType(73)
	_ReduceAddExplicitToProperImportClauses                       = _ReduceType(74)
	_ReduceImportClauseToProperImportClauses                      = _ReduceType(75)
	_ReduceProperImportClausesToImportClauses                     = _ReduceType(76)
	_ReduceImplicitToImportClauses                                = _ReduceType(77)
	_ReduceExplicitToImportClauses                                = _ReduceType(78)
	_ReduceCasePatternToCasePatterns                              = _ReduceType(79)
	_ReduceMultipleToCasePatterns                                 = _ReduceType(80)
	_ReduceToVarDeclPattern                                       = _ReduceType(81)
	_ReduceVarToVarOrLet                                          = _ReduceType(82)
	_ReduceLetToVarOrLet                                          = _ReduceType(83)
	_ReduceIdentifierToVarPattern                                 = _ReduceType(84)
	_ReduceUnderscoreToVarPattern                                 = _ReduceType(85)
	_ReduceTuplePatternToVarPattern                               = _ReduceType(86)
	_ReduceToTuplePattern                                         = _ReduceType(87)
	_ReduceFieldVarPatternToFieldVarPatterns                      = _ReduceType(88)
	_ReduceAddToFieldVarPatterns                                  = _ReduceType(89)
	_ReducePositionalToFieldVarPattern                            = _ReduceType(90)
	_ReduceNamedToFieldVarPattern                                 = _ReduceType(91)
	_ReduceEllipsisToFieldVarPattern                              = _ReduceType(92)
	_ReduceTypeExprToOptionalTypeExpr                             = _ReduceType(93)
	_ReduceNilToOptionalTypeExpr                                  = _ReduceType(94)
	_ReduceToAssignPattern                                        = _ReduceType(95)
	_ReduceAssignPatternToCasePattern                             = _ReduceType(96)
	_ReduceEnumMatchPatternToCasePattern                          = _ReduceType(97)
	_ReduceEnumNondataMatchPattenToCasePattern                    = _ReduceType(98)
	_ReduceEnumVarDeclPatternToCasePattern                        = _ReduceType(99)
	_ReduceIfExprToExpr                                           = _ReduceType(100)
	_ReduceSwitchExprToExpr                                       = _ReduceType(101)
	_ReduceLoopExprToExpr                                         = _ReduceType(102)
	_ReduceSequenceExprToExpr                                     = _ReduceType(103)
	_ReduceLabelDeclToOptionalLabelDecl                           = _ReduceType(104)
	_ReduceUnlabelledToOptionalLabelDecl                          = _ReduceType(105)
	_ReduceOrExprToSequenceExpr                                   = _ReduceType(106)
	_ReduceVarDeclPatternToSequenceExpr                           = _ReduceType(107)
	_ReduceAssignVarPatternToSequenceExpr                         = _ReduceType(108)
	_ReduceNoElseToIfExpr                                         = _ReduceType(109)
	_ReduceIfElseToIfExpr                                         = _ReduceType(110)
	_ReduceMultiIfElseToIfExpr                                    = _ReduceType(111)
	_ReduceSequenceExprToCondition                                = _ReduceType(112)
	_ReduceCaseToCondition                                        = _ReduceType(113)
	_ReduceToSwitchExpr                                           = _ReduceType(114)
	_ReduceInfiniteToLoopExpr                                     = _ReduceType(115)
	_ReduceDoWhileToLoopExpr                                      = _ReduceType(116)
	_ReduceWhileToLoopExpr                                        = _ReduceType(117)
	_ReduceIteratorToLoopExpr                                     = _ReduceType(118)
	_ReduceForToLoopExpr                                          = _ReduceType(119)
	_ReduceSequenceExprToOptionalSequenceExpr                     = _ReduceType(120)
	_ReduceNilToOptionalSequenceExpr                              = _ReduceType(121)
	_ReduceSequenceExprToForAssignment                            = _ReduceType(122)
	_ReduceAssignToForAssignment                                  = _ReduceType(123)
	_ReduceToCallExpr                                             = _ReduceType(124)
	_ReduceAddToProperArguments                                   = _ReduceType(125)
	_ReduceArgumentToProperArguments                              = _ReduceType(126)
	_ReduceProperArgumentsToArguments                             = _ReduceType(127)
	_ReduceImproperToArguments                                    = _ReduceType(128)
	_ReduceNilToArguments                                         = _ReduceType(129)
	_ReducePositionalToArgument                                   = _ReduceType(130)
	_ReduceColonExprToArgument                                    = _ReduceType(131)
	_ReduceNamedAssignmentToArgument                              = _ReduceType(132)
	_ReduceVarargAssignmentToArgument                             = _ReduceType(133)
	_ReduceSkipPatternToArgument                                  = _ReduceType(134)
	_ReduceUnitUnitPairToColonExpr                                = _ReduceType(135)
	_ReduceExprUnitPairToColonExpr                                = _ReduceType(136)
	_ReduceUnitExprPairToColonExpr                                = _ReduceType(137)
	_ReduceExprExprPairToColonExpr                                = _ReduceType(138)
	_ReduceColonExprUnitTupleToColonExpr                          = _ReduceType(139)
	_ReduceColonExprExprTupleToColonExpr                          = _ReduceType(140)
	_ReduceParseErrorExprToAtomExpr                               = _ReduceType(141)
	_ReduceLiteralExprToAtomExpr                                  = _ReduceType(142)
	_ReduceNamedExprToAtomExpr                                    = _ReduceType(143)
	_ReduceBlockExprToAtomExpr                                    = _ReduceType(144)
	_ReduceAnonymousFuncExprToAtomExpr                            = _ReduceType(145)
	_ReduceInitializeExprToAtomExpr                               = _ReduceType(146)
	_ReduceImplicitStructExprToAtomExpr                           = _ReduceType(147)
	_ReduceToParseErrorExpr                                       = _ReduceType(148)
	_ReduceTrueToLiteralExpr                                      = _ReduceType(149)
	_ReduceFalseToLiteralExpr                                     = _ReduceType(150)
	_ReduceIntegerLiteralToLiteralExpr                            = _ReduceType(151)
	_ReduceFloatLiteralToLiteralExpr                              = _ReduceType(152)
	_ReduceRuneLiteralToLiteralExpr                               = _ReduceType(153)
	_ReduceStringLiteralToLiteralExpr                             = _ReduceType(154)
	_ReduceIdentifierToNamedExpr                                  = _ReduceType(155)
	_ReduceUnderscoreToNamedExpr                                  = _ReduceType(156)
	_ReduceToBlockExpr                                            = _ReduceType(157)
	_ReduceToInitializeExpr                                       = _ReduceType(158)
	_ReduceToImplicitStructExpr                                   = _ReduceType(159)
	_ReduceAtomExprToAccessibleExpr                               = _ReduceType(160)
	_ReduceAccessExprToAccessibleExpr                             = _ReduceType(161)
	_ReduceCallExprToAccessibleExpr                               = _ReduceType(162)
	_ReduceIndexExprToAccessibleExpr                              = _ReduceType(163)
	_ReduceToAccessExpr                                           = _ReduceType(164)
	_ReduceToIndexExpr                                            = _ReduceType(165)
	_ReduceAccessibleExprToPostfixableExpr                        = _ReduceType(166)
	_ReducePostfixUnaryExprToPostfixableExpr                      = _ReduceType(167)
	_ReduceToPostfixUnaryExpr                                     = _ReduceType(168)
	_ReducePostfixableExprToPrefixableExpr                        = _ReduceType(169)
	_ReducePrefixUnaryExprToPrefixableExpr                        = _ReduceType(170)
	_ReduceToPrefixUnaryExpr                                      = _ReduceType(171)
	_ReduceNotToPrefixUnaryOp                                     = _ReduceType(172)
	_ReduceBitNegToPrefixUnaryOp                                  = _ReduceType(173)
	_ReduceSubToPrefixUnaryOp                                     = _ReduceType(174)
	_ReduceMulToPrefixUnaryOp                                     = _ReduceType(175)
	_ReduceBitAndToPrefixUnaryOp                                  = _ReduceType(176)
	_ReducePrefixableExprToMulExpr                                = _ReduceType(177)
	_ReduceBinaryMulExprToMulExpr                                 = _ReduceType(178)
	_ReduceToBinaryMulExpr                                        = _ReduceType(179)
	_ReduceMulToMulOp                                             = _ReduceType(180)
	_ReduceDivToMulOp                                             = _ReduceType(181)
	_ReduceModToMulOp                                             = _ReduceType(182)
	_ReduceBitAndToMulOp                                          = _ReduceType(183)
	_ReduceBitLshiftToMulOp                                       = _ReduceType(184)
	_ReduceBitRshiftToMulOp                                       = _ReduceType(185)
	_ReduceMulExprToAddExpr                                       = _ReduceType(186)
	_ReduceBinaryAddExprToAddExpr                                 = _ReduceType(187)
	_ReduceToBinaryAddExpr                                        = _ReduceType(188)
	_ReduceAddToAddOp                                             = _ReduceType(189)
	_ReduceSubToAddOp                                             = _ReduceType(190)
	_ReduceBitOrToAddOp                                           = _ReduceType(191)
	_ReduceBitXorToAddOp                                          = _ReduceType(192)
	_ReduceAddExprToCmpExpr                                       = _ReduceType(193)
	_ReduceBinaryCmpExprToCmpExpr                                 = _ReduceType(194)
	_ReduceToBinaryCmpExpr                                        = _ReduceType(195)
	_ReduceEqualToCmpOp                                           = _ReduceType(196)
	_ReduceNotEqualToCmpOp                                        = _ReduceType(197)
	_ReduceLessToCmpOp                                            = _ReduceType(198)
	_ReduceLessOrEqualToCmpOp                                     = _ReduceType(199)
	_ReduceGreaterToCmpOp                                         = _ReduceType(200)
	_ReduceGreaterOrEqualToCmpOp                                  = _ReduceType(201)
	_ReduceCmpExprToAndExpr                                       = _ReduceType(202)
	_ReduceBinaryAndExprToAndExpr                                 = _ReduceType(203)
	_ReduceToBinaryAndExpr                                        = _ReduceType(204)
	_ReduceAndExprToOrExpr                                        = _ReduceType(205)
	_ReduceBinaryOrExprToOrExpr                                   = _ReduceType(206)
	_ReduceToBinaryOrExpr                                         = _ReduceType(207)
	_ReduceExplicitStructTypeExprToInitializableTypeExpr          = _ReduceType(208)
	_ReduceSliceTypeExprToInitializableTypeExpr                   = _ReduceType(209)
	_ReduceArrayTypeExprToInitializableTypeExpr                   = _ReduceType(210)
	_ReduceMapTypeExprToInitializableTypeExpr                     = _ReduceType(211)
	_ReduceToSliceTypeExpr                                        = _ReduceType(212)
	_ReduceToArrayTypeExpr                                        = _ReduceType(213)
	_ReduceToMapTypeExpr                                          = _ReduceType(214)
	_ReduceInitializableTypeExprToAtomTypeExpr                    = _ReduceType(215)
	_ReduceNamedTypeExprToAtomTypeExpr                            = _ReduceType(216)
	_ReduceInferredTypeExprToAtomTypeExpr                         = _ReduceType(217)
	_ReduceImplicitStructTypeExprToAtomTypeExpr                   = _ReduceType(218)
	_ReduceExplicitEnumTypeExprToAtomTypeExpr                     = _ReduceType(219)
	_ReduceImplicitEnumTypeExprToAtomTypeExpr                     = _ReduceType(220)
	_ReduceTraitTypeExprToAtomTypeExpr                            = _ReduceType(221)
	_ReduceFuncTypeExprToAtomTypeExpr                             = _ReduceType(222)
	_ReduceParseErrorTypeExprToAtomTypeExpr                       = _ReduceType(223)
	_ReduceLocalToNamedTypeExpr                                   = _ReduceType(224)
	_ReduceExternalToNamedTypeExpr                                = _ReduceType(225)
	_ReduceDotToInferredTypeExpr                                  = _ReduceType(226)
	_ReduceUnderscoreToInferredTypeExpr                           = _ReduceType(227)
	_ReduceToParseErrorTypeExpr                                   = _ReduceType(228)
	_ReduceAtomTypeExprToReturnableTypeExpr                       = _ReduceType(229)
	_ReducePrefixUnaryTypeExprToReturnableTypeExpr                = _ReduceType(230)
	_ReduceToPrefixUnaryTypeExpr                                  = _ReduceType(231)
	_ReduceQuestionToPrefixUnaryTypeOp                            = _ReduceType(232)
	_ReduceExclaimToPrefixUnaryTypeOp                             = _ReduceType(233)
	_ReduceBitAndToPrefixUnaryTypeOp                              = _ReduceType(234)
	_ReduceBitNegToPrefixUnaryTypeOp                              = _ReduceType(235)
	_ReduceTildeTildeToPrefixUnaryTypeOp                          = _ReduceType(236)
	_ReduceReturnableTypeExprToTypeExpr                           = _ReduceType(237)
	_ReduceBinaryTypeExprToTypeExpr                               = _ReduceType(238)
	_ReduceToBinaryTypeExpr                                       = _ReduceType(239)
	_ReduceMulToBinaryTypeOp                                      = _ReduceType(240)
	_ReduceAddToBinaryTypeOp                                      = _ReduceType(241)
	_ReduceSubToBinaryTypeOp                                      = _ReduceType(242)
	_ReduceDefinitionToTypeDef                                    = _ReduceType(243)
	_ReduceConstrainedDefToTypeDef                                = _ReduceType(244)
	_ReduceAliasToTypeDef                                         = _ReduceType(245)
	_ReduceUnconstrainedToGenericParameterDef                     = _ReduceType(246)
	_ReduceConstrainedToGenericParameterDef                       = _ReduceType(247)
	_ReduceGenericToGenericParameterDefs                          = _ReduceType(248)
	_ReduceNilToGenericParameterDefs                              = _ReduceType(249)
	_ReduceAddToProperGenericParameterDefList                     = _ReduceType(250)
	_ReduceGenericParameterDefToProperGenericParameterDefList     = _ReduceType(251)
	_ReduceProperGenericParameterDefListToGenericParameterDefList = _ReduceType(252)
	_ReduceImproperToGenericParameterDefList                      = _ReduceType(253)
	_ReduceNilToGenericParameterDefList                           = _ReduceType(254)
	_ReduceBindingToGenericTypeArguments                          = _ReduceType(255)
	_ReduceNilToGenericTypeArguments                              = _ReduceType(256)
	_ReduceAddToProperTypeArguments                               = _ReduceType(257)
	_ReduceTypeExprToProperTypeArguments                          = _ReduceType(258)
	_ReduceProperTypeArgumentsToTypeArguments                     = _ReduceType(259)
	_ReduceImproperToTypeArguments                                = _ReduceType(260)
	_ReduceNilToTypeArguments                                     = _ReduceType(261)
	_ReduceNamedToFieldDef                                        = _ReduceType(262)
	_ReduceUnnamedToFieldDef                                      = _ReduceType(263)
	_ReduceStructPaddingToFieldDef                                = _ReduceType(264)
	_ReduceMethodSignatureToFieldDef                              = _ReduceType(265)
	_ReduceUnsafeStatementToFieldDef                              = _ReduceType(266)
	_ReduceDefaultToOptionalDefault                               = _ReduceType(267)
	_ReduceNilToOptionalDefault                                   = _ReduceType(268)
	_ReduceAddToProperImplicitFieldDefs                           = _ReduceType(269)
	_ReduceFieldDefToProperImplicitFieldDefs                      = _ReduceType(270)
	_ReduceProperImplicitFieldDefsToImplicitFieldDefs             = _ReduceType(271)
	_ReduceImproperToImplicitFieldDefs                            = _ReduceType(272)
	_ReduceNilToImplicitFieldDefs                                 = _ReduceType(273)
	_ReduceToImplicitStructTypeExpr                               = _ReduceType(274)
	_ReduceAddImplicitToProperExplicitFieldDefs                   = _ReduceType(275)
	_ReduceAddExplicitToProperExplicitFieldDefs                   = _ReduceType(276)
	_ReduceFieldDefToProperExplicitFieldDefs                      = _ReduceType(277)
	_ReduceProperExplicitFieldDefsToExplicitFieldDefs             = _ReduceType(278)
	_ReduceImproperImplicitToExplicitFieldDefs                    = _ReduceType(279)
	_ReduceImproperExplicitToExplicitFieldDefs                    = _ReduceType(280)
	_ReduceNilToExplicitFieldDefs                                 = _ReduceType(281)
	_ReduceToExplicitStructTypeExpr                               = _ReduceType(282)
	_ReduceToTraitTypeExpr                                        = _ReduceType(283)
	_ReducePairToProperImplicitEnumValueDefs                      = _ReduceType(284)
	_ReduceAddToProperImplicitEnumValueDefs                       = _ReduceType(285)
	_ReduceProperImplicitEnumValueDefsToImplicitEnumValueDefs     = _ReduceType(286)
	_ReduceImproperToImplicitEnumValueDefs                        = _ReduceType(287)
	_ReduceToImplicitEnumTypeExpr                                 = _ReduceType(288)
	_ReduceExplicitPairToProperExplicitEnumValueDefs              = _ReduceType(289)
	_ReduceImplicitPairToProperExplicitEnumValueDefs              = _ReduceType(290)
	_ReduceExplicitAddToProperExplicitEnumValueDefs               = _ReduceType(291)
	_ReduceImplicitAddToProperExplicitEnumValueDefs               = _ReduceType(292)
	_ReduceProperExplicitEnumValueDefsToExplicitEnumValueDefs     = _ReduceType(293)
	_ReduceImproperToExplicitEnumValueDefs                        = _ReduceType(294)
	_ReduceToExplicitEnumTypeExpr                                 = _ReduceType(295)
	_ReduceReturnableTypeExprToReturnType                         = _ReduceType(296)
	_ReduceNilToReturnType                                        = _ReduceType(297)
	_ReduceNamedTypedArgToProperParameterDef                      = _ReduceType(298)
	_ReduceNamedTypedVarargToProperParameterDef                   = _ReduceType(299)
	_ReduceNamedInferredVarargToProperParameterDef                = _ReduceType(300)
	_ReduceIgnoreTypedArgToProperParameterDef                     = _ReduceType(301)
	_ReduceIgnoreInferredVarargToProperParameterDef               = _ReduceType(302)
	_ReduceIgnoreTypedVarargToProperParameterDef                  = _ReduceType(303)
	_ReduceProperParameterDefToParameterDecl                      = _ReduceType(304)
	_ReduceUnnamedTypedArgToParameterDecl                         = _ReduceType(305)
	_ReduceUnnamedInferredVarargToParameterDecl                   = _ReduceType(306)
	_ReduceUnnamedTypedVarargToParameterDecl                      = _ReduceType(307)
	_ReduceProperParameterDefToParameterDef                       = _ReduceType(308)
	_ReduceNamedInferredArgToParameterDef                         = _ReduceType(309)
	_ReduceIgnoreInferredArgToParameterDef                        = _ReduceType(310)
	_ReduceAddToProperParameterDecls                              = _ReduceType(311)
	_ReduceParameterDeclToProperParameterDecls                    = _ReduceType(312)
	_ReduceProperParameterDeclsToParameterDecls                   = _ReduceType(313)
	_ReduceImproperToParameterDecls                               = _ReduceType(314)
	_ReduceNilToParameterDecls                                    = _ReduceType(315)
	_ReduceAddToProperParameterDefs                               = _ReduceType(316)
	_ReduceParameterDefToProperParameterDefs                      = _ReduceType(317)
	_ReduceProperParameterDefsToParameterDefs                     = _ReduceType(318)
	_ReduceImproperToParameterDefs                                = _ReduceType(319)
	_ReduceNilToParameterDefs                                     = _ReduceType(320)
	_ReduceToFuncTypeExpr                                         = _ReduceType(321)
	_ReduceToMethodSignature                                      = _ReduceType(322)
	_ReduceFuncDefToNamedFuncDef                                  = _ReduceType(323)
	_ReduceMethodDefToNamedFuncDef                                = _ReduceType(324)
	_ReduceAliasToNamedFuncDef                                    = _ReduceType(325)
	_ReduceToAnonymousFuncExpr                                    = _ReduceType(326)
	_ReduceNoSpecToPackageDef                                     = _ReduceType(327)
	_ReduceWithSpecToPackageDef                                   = _ReduceType(328)
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
	case _ReduceExprOrImproperStructStatementToSimpleStatement:
		return "ExprOrImproperStructStatementToSimpleStatement"
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
	case _ReduceToExprOrImproperStructStatement:
		return "ToExprOrImproperStructStatement"
	case _ReduceExprToExprs:
		return "ExprToExprs"
	case _ReduceAddToExprs:
		return "AddToExprs"
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
	case _ReduceUnusableImportToImportClause:
		return "UnusableImportToImportClause"
	case _ReduceImportToLocalToImportClause:
		return "ImportToLocalToImportClause"
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
	case _ReduceUnderscoreToVarPattern:
		return "UnderscoreToVarPattern"
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
	case _ReduceTypeExprToOptionalTypeExpr:
		return "TypeExprToOptionalTypeExpr"
	case _ReduceNilToOptionalTypeExpr:
		return "NilToOptionalTypeExpr"
	case _ReduceToAssignPattern:
		return "ToAssignPattern"
	case _ReduceAssignPatternToCasePattern:
		return "AssignPatternToCasePattern"
	case _ReduceEnumMatchPatternToCasePattern:
		return "EnumMatchPatternToCasePattern"
	case _ReduceEnumNondataMatchPattenToCasePattern:
		return "EnumNondataMatchPattenToCasePattern"
	case _ReduceEnumVarDeclPatternToCasePattern:
		return "EnumVarDeclPatternToCasePattern"
	case _ReduceIfExprToExpr:
		return "IfExprToExpr"
	case _ReduceSwitchExprToExpr:
		return "SwitchExprToExpr"
	case _ReduceLoopExprToExpr:
		return "LoopExprToExpr"
	case _ReduceSequenceExprToExpr:
		return "SequenceExprToExpr"
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
	case _ReduceNamedExprToAtomExpr:
		return "NamedExprToAtomExpr"
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
	case _ReduceIdentifierToNamedExpr:
		return "IdentifierToNamedExpr"
	case _ReduceUnderscoreToNamedExpr:
		return "UnderscoreToNamedExpr"
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
	case _ReduceExplicitStructTypeExprToInitializableTypeExpr:
		return "ExplicitStructTypeExprToInitializableTypeExpr"
	case _ReduceSliceTypeExprToInitializableTypeExpr:
		return "SliceTypeExprToInitializableTypeExpr"
	case _ReduceArrayTypeExprToInitializableTypeExpr:
		return "ArrayTypeExprToInitializableTypeExpr"
	case _ReduceMapTypeExprToInitializableTypeExpr:
		return "MapTypeExprToInitializableTypeExpr"
	case _ReduceToSliceTypeExpr:
		return "ToSliceTypeExpr"
	case _ReduceToArrayTypeExpr:
		return "ToArrayTypeExpr"
	case _ReduceToMapTypeExpr:
		return "ToMapTypeExpr"
	case _ReduceInitializableTypeExprToAtomTypeExpr:
		return "InitializableTypeExprToAtomTypeExpr"
	case _ReduceNamedTypeExprToAtomTypeExpr:
		return "NamedTypeExprToAtomTypeExpr"
	case _ReduceInferredTypeExprToAtomTypeExpr:
		return "InferredTypeExprToAtomTypeExpr"
	case _ReduceImplicitStructTypeExprToAtomTypeExpr:
		return "ImplicitStructTypeExprToAtomTypeExpr"
	case _ReduceExplicitEnumTypeExprToAtomTypeExpr:
		return "ExplicitEnumTypeExprToAtomTypeExpr"
	case _ReduceImplicitEnumTypeExprToAtomTypeExpr:
		return "ImplicitEnumTypeExprToAtomTypeExpr"
	case _ReduceTraitTypeExprToAtomTypeExpr:
		return "TraitTypeExprToAtomTypeExpr"
	case _ReduceFuncTypeExprToAtomTypeExpr:
		return "FuncTypeExprToAtomTypeExpr"
	case _ReduceParseErrorTypeExprToAtomTypeExpr:
		return "ParseErrorTypeExprToAtomTypeExpr"
	case _ReduceLocalToNamedTypeExpr:
		return "LocalToNamedTypeExpr"
	case _ReduceExternalToNamedTypeExpr:
		return "ExternalToNamedTypeExpr"
	case _ReduceDotToInferredTypeExpr:
		return "DotToInferredTypeExpr"
	case _ReduceUnderscoreToInferredTypeExpr:
		return "UnderscoreToInferredTypeExpr"
	case _ReduceToParseErrorTypeExpr:
		return "ToParseErrorTypeExpr"
	case _ReduceAtomTypeExprToReturnableTypeExpr:
		return "AtomTypeExprToReturnableTypeExpr"
	case _ReducePrefixUnaryTypeExprToReturnableTypeExpr:
		return "PrefixUnaryTypeExprToReturnableTypeExpr"
	case _ReduceToPrefixUnaryTypeExpr:
		return "ToPrefixUnaryTypeExpr"
	case _ReduceQuestionToPrefixUnaryTypeOp:
		return "QuestionToPrefixUnaryTypeOp"
	case _ReduceExclaimToPrefixUnaryTypeOp:
		return "ExclaimToPrefixUnaryTypeOp"
	case _ReduceBitAndToPrefixUnaryTypeOp:
		return "BitAndToPrefixUnaryTypeOp"
	case _ReduceBitNegToPrefixUnaryTypeOp:
		return "BitNegToPrefixUnaryTypeOp"
	case _ReduceTildeTildeToPrefixUnaryTypeOp:
		return "TildeTildeToPrefixUnaryTypeOp"
	case _ReduceReturnableTypeExprToTypeExpr:
		return "ReturnableTypeExprToTypeExpr"
	case _ReduceBinaryTypeExprToTypeExpr:
		return "BinaryTypeExprToTypeExpr"
	case _ReduceToBinaryTypeExpr:
		return "ToBinaryTypeExpr"
	case _ReduceMulToBinaryTypeOp:
		return "MulToBinaryTypeOp"
	case _ReduceAddToBinaryTypeOp:
		return "AddToBinaryTypeOp"
	case _ReduceSubToBinaryTypeOp:
		return "SubToBinaryTypeOp"
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
	case _ReduceGenericToGenericParameterDefs:
		return "GenericToGenericParameterDefs"
	case _ReduceNilToGenericParameterDefs:
		return "NilToGenericParameterDefs"
	case _ReduceAddToProperGenericParameterDefList:
		return "AddToProperGenericParameterDefList"
	case _ReduceGenericParameterDefToProperGenericParameterDefList:
		return "GenericParameterDefToProperGenericParameterDefList"
	case _ReduceProperGenericParameterDefListToGenericParameterDefList:
		return "ProperGenericParameterDefListToGenericParameterDefList"
	case _ReduceImproperToGenericParameterDefList:
		return "ImproperToGenericParameterDefList"
	case _ReduceNilToGenericParameterDefList:
		return "NilToGenericParameterDefList"
	case _ReduceBindingToGenericTypeArguments:
		return "BindingToGenericTypeArguments"
	case _ReduceNilToGenericTypeArguments:
		return "NilToGenericTypeArguments"
	case _ReduceAddToProperTypeArguments:
		return "AddToProperTypeArguments"
	case _ReduceTypeExprToProperTypeArguments:
		return "TypeExprToProperTypeArguments"
	case _ReduceProperTypeArgumentsToTypeArguments:
		return "ProperTypeArgumentsToTypeArguments"
	case _ReduceImproperToTypeArguments:
		return "ImproperToTypeArguments"
	case _ReduceNilToTypeArguments:
		return "NilToTypeArguments"
	case _ReduceNamedToFieldDef:
		return "NamedToFieldDef"
	case _ReduceUnnamedToFieldDef:
		return "UnnamedToFieldDef"
	case _ReduceStructPaddingToFieldDef:
		return "StructPaddingToFieldDef"
	case _ReduceMethodSignatureToFieldDef:
		return "MethodSignatureToFieldDef"
	case _ReduceUnsafeStatementToFieldDef:
		return "UnsafeStatementToFieldDef"
	case _ReduceDefaultToOptionalDefault:
		return "DefaultToOptionalDefault"
	case _ReduceNilToOptionalDefault:
		return "NilToOptionalDefault"
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
	case _ReduceToImplicitStructTypeExpr:
		return "ToImplicitStructTypeExpr"
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
	case _ReduceToExplicitStructTypeExpr:
		return "ToExplicitStructTypeExpr"
	case _ReduceToTraitTypeExpr:
		return "ToTraitTypeExpr"
	case _ReducePairToProperImplicitEnumValueDefs:
		return "PairToProperImplicitEnumValueDefs"
	case _ReduceAddToProperImplicitEnumValueDefs:
		return "AddToProperImplicitEnumValueDefs"
	case _ReduceProperImplicitEnumValueDefsToImplicitEnumValueDefs:
		return "ProperImplicitEnumValueDefsToImplicitEnumValueDefs"
	case _ReduceImproperToImplicitEnumValueDefs:
		return "ImproperToImplicitEnumValueDefs"
	case _ReduceToImplicitEnumTypeExpr:
		return "ToImplicitEnumTypeExpr"
	case _ReduceExplicitPairToProperExplicitEnumValueDefs:
		return "ExplicitPairToProperExplicitEnumValueDefs"
	case _ReduceImplicitPairToProperExplicitEnumValueDefs:
		return "ImplicitPairToProperExplicitEnumValueDefs"
	case _ReduceExplicitAddToProperExplicitEnumValueDefs:
		return "ExplicitAddToProperExplicitEnumValueDefs"
	case _ReduceImplicitAddToProperExplicitEnumValueDefs:
		return "ImplicitAddToProperExplicitEnumValueDefs"
	case _ReduceProperExplicitEnumValueDefsToExplicitEnumValueDefs:
		return "ProperExplicitEnumValueDefsToExplicitEnumValueDefs"
	case _ReduceImproperToExplicitEnumValueDefs:
		return "ImproperToExplicitEnumValueDefs"
	case _ReduceToExplicitEnumTypeExpr:
		return "ToExplicitEnumTypeExpr"
	case _ReduceReturnableTypeExprToReturnType:
		return "ReturnableTypeExprToReturnType"
	case _ReduceNilToReturnType:
		return "NilToReturnType"
	case _ReduceNamedTypedArgToProperParameterDef:
		return "NamedTypedArgToProperParameterDef"
	case _ReduceNamedTypedVarargToProperParameterDef:
		return "NamedTypedVarargToProperParameterDef"
	case _ReduceNamedInferredVarargToProperParameterDef:
		return "NamedInferredVarargToProperParameterDef"
	case _ReduceIgnoreTypedArgToProperParameterDef:
		return "IgnoreTypedArgToProperParameterDef"
	case _ReduceIgnoreInferredVarargToProperParameterDef:
		return "IgnoreInferredVarargToProperParameterDef"
	case _ReduceIgnoreTypedVarargToProperParameterDef:
		return "IgnoreTypedVarargToProperParameterDef"
	case _ReduceProperParameterDefToParameterDecl:
		return "ProperParameterDefToParameterDecl"
	case _ReduceUnnamedTypedArgToParameterDecl:
		return "UnnamedTypedArgToParameterDecl"
	case _ReduceUnnamedInferredVarargToParameterDecl:
		return "UnnamedInferredVarargToParameterDecl"
	case _ReduceUnnamedTypedVarargToParameterDecl:
		return "UnnamedTypedVarargToParameterDecl"
	case _ReduceProperParameterDefToParameterDef:
		return "ProperParameterDefToParameterDef"
	case _ReduceNamedInferredArgToParameterDef:
		return "NamedInferredArgToParameterDef"
	case _ReduceIgnoreInferredArgToParameterDef:
		return "IgnoreInferredArgToParameterDef"
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
	case _ReduceToFuncTypeExpr:
		return "ToFuncTypeExpr"
	case _ReduceToMethodSignature:
		return "ToMethodSignature"
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
	_State471 = _StateId(471)
	_State472 = _StateId(472)
	_State473 = _StateId(473)
	_State474 = _StateId(474)
	_State475 = _StateId(475)
	_State476 = _StateId(476)
	_State477 = _StateId(477)
	_State478 = _StateId(478)
)

type Symbol struct {
	SymbolId_ SymbolId

	Generic_ GenericSymbol

	Argument          *Argument
	ArgumentList      *ArgumentList
	ColonExpr         *ColonExpr
	Count             TokenCount
	Expression        Expression
	Parameter         *Parameter
	Parameters        *ParameterList
	ParseError        ParseErrorSymbol
	SourceDefinition  SourceDefinition
	SourceDefinitions []SourceDefinition
	Statement         Statement
	Statements        *StatementList
	TypeArgumentList  *TypeArgumentList
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
	case CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, EllipsisToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken:
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
	case ColonExprType:
		loc, ok := interface{}(s.ColonExpr).(locator)
		if ok {
			return loc.Loc()
		}
	case NewlinesToken:
		loc, ok := interface{}(s.Count).(locator)
		if ok {
			return loc.Loc()
		}
	case ExprType, SequenceExprType, IfExprType, SwitchExprType, LoopExprType, OptionalSequenceExprType, CallExprType, AtomExprType, ParseErrorExprType, LiteralExprType, NamedExprType, BlockExprType, InitializeExprType, ImplicitStructExprType, AccessibleExprType, AccessExprType, IndexExprType, PostfixableExprType, PostfixUnaryExprType, PrefixableExprType, PrefixUnaryExprType, MulExprType, BinaryMulExprType, AddExprType, BinaryAddExprType, CmpExprType, BinaryCmpExprType, AndExprType, BinaryAndExprType, OrExprType, BinaryOrExprType, AnonymousFuncExprType:
		loc, ok := interface{}(s.Expression).(locator)
		if ok {
			return loc.Loc()
		}
	case ProperParameterDefType, ParameterDeclType, ParameterDefType:
		loc, ok := interface{}(s.Parameter).(locator)
		if ok {
			return loc.Loc()
		}
	case ProperParameterDeclsType, ParameterDeclsType, ProperParameterDefsType, ParameterDefsType:
		loc, ok := interface{}(s.Parameters).(locator)
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
	case SimpleStatementType, StatementType, OptionalSimpleStatementType, ExprOrImproperStructStatementType, CallbackOpStatementType, UnsafeStatementType, JumpStatementType, FallthroughStatementType, AssignStatementType, UnaryOpAssignStatementType, BinaryOpAssignStatementType, ImportStatementType:
		loc, ok := interface{}(s.Statement).(locator)
		if ok {
			return loc.Loc()
		}
	case ProperStatementsType, StatementsType:
		loc, ok := interface{}(s.Statements).(locator)
		if ok {
			return loc.Loc()
		}
	case GenericTypeArgumentsType, ProperTypeArgumentsType, TypeArgumentsType:
		loc, ok := interface{}(s.TypeArgumentList).(locator)
		if ok {
			return loc.Loc()
		}
	case OptionalTypeExprType, InitializableTypeExprType, SliceTypeExprType, ArrayTypeExprType, MapTypeExprType, AtomTypeExprType, NamedTypeExprType, InferredTypeExprType, ParseErrorTypeExprType, ReturnableTypeExprType, PrefixUnaryTypeExprType, TypeExprType, BinaryTypeExprType, ImplicitStructTypeExprType, ExplicitStructTypeExprType, TraitTypeExprType, ImplicitEnumTypeExprType, ExplicitEnumTypeExprType, ReturnTypeType, FuncTypeExprType:
		loc, ok := interface{}(s.TypeExpression).(locator)
		if ok {
			return loc.Loc()
		}
	case CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, EllipsisToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, CallbackOpType, JumpTypeType, UnaryOpAssignType, BinaryOpAssignType, VarOrLetType, OptionalLabelDeclType, PrefixUnaryOpType, MulOpType, AddOpType, CmpOpType, PrefixUnaryTypeOpType, BinaryTypeOpType:
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
	case ColonExprType:
		loc, ok := interface{}(s.ColonExpr).(locator)
		if ok {
			return loc.End()
		}
	case NewlinesToken:
		loc, ok := interface{}(s.Count).(locator)
		if ok {
			return loc.End()
		}
	case ExprType, SequenceExprType, IfExprType, SwitchExprType, LoopExprType, OptionalSequenceExprType, CallExprType, AtomExprType, ParseErrorExprType, LiteralExprType, NamedExprType, BlockExprType, InitializeExprType, ImplicitStructExprType, AccessibleExprType, AccessExprType, IndexExprType, PostfixableExprType, PostfixUnaryExprType, PrefixableExprType, PrefixUnaryExprType, MulExprType, BinaryMulExprType, AddExprType, BinaryAddExprType, CmpExprType, BinaryCmpExprType, AndExprType, BinaryAndExprType, OrExprType, BinaryOrExprType, AnonymousFuncExprType:
		loc, ok := interface{}(s.Expression).(locator)
		if ok {
			return loc.End()
		}
	case ProperParameterDefType, ParameterDeclType, ParameterDefType:
		loc, ok := interface{}(s.Parameter).(locator)
		if ok {
			return loc.End()
		}
	case ProperParameterDeclsType, ParameterDeclsType, ProperParameterDefsType, ParameterDefsType:
		loc, ok := interface{}(s.Parameters).(locator)
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
	case SimpleStatementType, StatementType, OptionalSimpleStatementType, ExprOrImproperStructStatementType, CallbackOpStatementType, UnsafeStatementType, JumpStatementType, FallthroughStatementType, AssignStatementType, UnaryOpAssignStatementType, BinaryOpAssignStatementType, ImportStatementType:
		loc, ok := interface{}(s.Statement).(locator)
		if ok {
			return loc.End()
		}
	case ProperStatementsType, StatementsType:
		loc, ok := interface{}(s.Statements).(locator)
		if ok {
			return loc.End()
		}
	case GenericTypeArgumentsType, ProperTypeArgumentsType, TypeArgumentsType:
		loc, ok := interface{}(s.TypeArgumentList).(locator)
		if ok {
			return loc.End()
		}
	case OptionalTypeExprType, InitializableTypeExprType, SliceTypeExprType, ArrayTypeExprType, MapTypeExprType, AtomTypeExprType, NamedTypeExprType, InferredTypeExprType, ParseErrorTypeExprType, ReturnableTypeExprType, PrefixUnaryTypeExprType, TypeExprType, BinaryTypeExprType, ImplicitStructTypeExprType, ExplicitStructTypeExprType, TraitTypeExprType, ImplicitEnumTypeExprType, ExplicitEnumTypeExprType, ReturnTypeType, FuncTypeExprType:
		loc, ok := interface{}(s.TypeExpression).(locator)
		if ok {
			return loc.End()
		}
	case CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, EllipsisToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, CallbackOpType, JumpTypeType, UnaryOpAssignType, BinaryOpAssignType, VarOrLetType, OptionalLabelDeclType, PrefixUnaryOpType, MulOpType, AddOpType, CmpOpType, PrefixUnaryTypeOpType, BinaryTypeOpType:
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
		//line grammar.lr:66:4
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
		//line grammar.lr:72:4
		symbol.SourceDefinition = args[0].SourceDefinition
		err = nil
	case _ReduceTypeDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		//line grammar.lr:73:4
		symbol.SourceDefinition = args[0].SourceDefinition
		err = nil
	case _ReduceNamedFuncDefToDefinition:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = DefinitionType
		//line grammar.lr:74:4
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
		//line grammar.lr:111:4
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
		//line grammar.lr:117:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceExprOrImproperStructStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:119:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceCallbackOpStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:122:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceJumpStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:125:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceAssignStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:130:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceUnaryOpAssignStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:136:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceBinaryOpAssignStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:137:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceFallthroughStatementToSimpleStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = SimpleStatementType
		//line grammar.lr:141:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceSimpleStatementToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:144:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceImportStatementToStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = StatementType
		//line grammar.lr:147:4
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
		//line grammar.lr:156:4
		symbol.Statement = args[0].Statement
		err = nil
	case _ReduceNilToOptionalSimpleStatement:
		symbol.SymbolId_ = OptionalSimpleStatementType
		symbol.Statement, err = reducer.NilToOptionalSimpleStatement()
	case _ReduceToExprOrImproperStructStatement:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExprOrImproperStructStatementType
		symbol.Statement, err = reducer.ToExprOrImproperStructStatement(args[0].Generic_)
	case _ReduceExprToExprs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExprsType
		symbol.Generic_, err = reducer.ExprToExprs(args[0].Expression)
	case _ReduceAddToExprs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ExprsType
		symbol.Generic_, err = reducer.AddToExprs(args[0].Generic_, args[1].Value, args[2].Expression)
	case _ReduceAsyncToCallbackOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CallbackOpType
		//line grammar.lr:174:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDeferToCallbackOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CallbackOpType
		//line grammar.lr:175:4
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
		//line grammar.lr:214:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBreakToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		//line grammar.lr:215:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceContinueToJumpType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = JumpTypeType
		//line grammar.lr:216:4
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
		//line grammar.lr:229:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubOneAssignToUnaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = UnaryOpAssignType
		//line grammar.lr:230:4
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
		//line grammar.lr:235:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:236:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:237:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDivAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:238:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceModAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:239:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:240:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:241:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitOrAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:242:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitXorAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:243:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitLshiftAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:244:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitRshiftAssignToBinaryOpAssign:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryOpAssignType
		//line grammar.lr:245:4
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
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClauseType
		symbol.Generic_, err = reducer.AliasToImportClause(args[0].Value, args[1].Value)
	case _ReduceUnusableImportToImportClause:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClauseType
		symbol.Generic_, err = reducer.UnusableImportToImportClause(args[0].Value, args[1].Value)
	case _ReduceImportToLocalToImportClause:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImportClauseType
		symbol.Generic_, err = reducer.ImportToLocalToImportClause(args[0].Value, args[1].Value)
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
		//line grammar.lr:267:4
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
		//line grammar.lr:289:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLetToVarOrLet:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarOrLetType
		//line grammar.lr:290:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceIdentifierToVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarPatternType
		symbol.Generic_, err = reducer.IdentifierToVarPattern(args[0].Value)
	case _ReduceUnderscoreToVarPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = VarPatternType
		symbol.Generic_, err = reducer.UnderscoreToVarPattern(args[0].Value)
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
	case _ReduceTypeExprToOptionalTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OptionalTypeExprType
		//line grammar.lr:309:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceNilToOptionalTypeExpr:
		symbol.SymbolId_ = OptionalTypeExprType
		symbol.TypeExpression, err = reducer.NilToOptionalTypeExpr()
	case _ReduceToAssignPattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AssignPatternType
		symbol.Generic_, err = reducer.ToAssignPattern(args[0].Expression)
	case _ReduceAssignPatternToCasePattern:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CasePatternType
		//line grammar.lr:320:4
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _ReduceEnumMatchPatternToCasePattern:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = CasePatternType
		symbol.Generic_, err = reducer.EnumMatchPatternToCasePattern(args[0].Value, args[1].Value, args[2].Expression)
	case _ReduceEnumNondataMatchPattenToCasePattern:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = CasePatternType
		symbol.Generic_, err = reducer.EnumNondataMatchPattenToCasePattern(args[0].Value, args[1].Value)
	case _ReduceEnumVarDeclPatternToCasePattern:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = CasePatternType
		symbol.Generic_, err = reducer.EnumVarDeclPatternToCasePattern(args[0].Value, args[1].Value, args[2].Value, args[3].Generic_)
	case _ReduceIfExprToExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExprType
		symbol.Expression, err = reducer.IfExprToExpr(args[0].Value, args[1].Expression)
	case _ReduceSwitchExprToExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExprType
		symbol.Expression, err = reducer.SwitchExprToExpr(args[0].Value, args[1].Expression)
	case _ReduceLoopExprToExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExprType
		symbol.Expression, err = reducer.LoopExprToExpr(args[0].Value, args[1].Expression)
	case _ReduceSequenceExprToExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExprType
		//line grammar.lr:339:4
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
		//line grammar.lr:350:4
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
		symbol.Expression, err = reducer.ToCallExpr(args[0].Expression, args[1].TypeArgumentList, args[2].Value, args[3].ArgumentList, args[4].Value)
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
		//line grammar.lr:436:4
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
		symbol.Argument, err = reducer.ColonExprToArgument(args[0].ColonExpr)
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
		symbol.ColonExpr, err = reducer.UnitUnitPairToColonExpr(args[0].Value)
	case _ReduceExprUnitPairToColonExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ColonExprType
		symbol.ColonExpr, err = reducer.ExprUnitPairToColonExpr(args[0].Expression, args[1].Value)
	case _ReduceUnitExprPairToColonExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ColonExprType
		symbol.ColonExpr, err = reducer.UnitExprPairToColonExpr(args[0].Value, args[1].Expression)
	case _ReduceExprExprPairToColonExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ColonExprType
		symbol.ColonExpr, err = reducer.ExprExprPairToColonExpr(args[0].Expression, args[1].Value, args[2].Expression)
	case _ReduceColonExprUnitTupleToColonExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ColonExprType
		symbol.ColonExpr, err = reducer.ColonExprUnitTupleToColonExpr(args[0].ColonExpr, args[1].Value)
	case _ReduceColonExprExprTupleToColonExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ColonExprType
		symbol.ColonExpr, err = reducer.ColonExprExprTupleToColonExpr(args[0].ColonExpr, args[1].Value, args[2].Expression)
	case _ReduceParseErrorExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:470:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceLiteralExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:471:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceNamedExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:472:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBlockExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:473:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceAnonymousFuncExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:474:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceInitializeExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:475:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceImplicitStructExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:476:4
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
	case _ReduceIdentifierToNamedExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = NamedExprType
		symbol.Expression, err = reducer.IdentifierToNamedExpr(args[0].Value)
	case _ReduceUnderscoreToNamedExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = NamedExprType
		symbol.Expression, err = reducer.UnderscoreToNamedExpr(args[0].Value)
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
	case _ReduceExplicitStructTypeExprToInitializableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeExprType
		//line grammar.lr:592:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceSliceTypeExprToInitializableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeExprType
		//line grammar.lr:593:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceArrayTypeExprToInitializableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeExprType
		//line grammar.lr:594:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceMapTypeExprToInitializableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeExprType
		//line grammar.lr:595:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceToSliceTypeExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = SliceTypeExprType
		symbol.TypeExpression, err = reducer.ToSliceTypeExpr(args[0].Value, args[1].TypeExpression, args[2].Value)
	case _ReduceToArrayTypeExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = ArrayTypeExprType
		symbol.TypeExpression, err = reducer.ToArrayTypeExpr(args[0].Value, args[1].TypeExpression, args[2].Value, args[3].Value, args[4].Value)
	case _ReduceToMapTypeExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = MapTypeExprType
		symbol.TypeExpression, err = reducer.ToMapTypeExpr(args[0].Value, args[1].TypeExpression, args[2].Value, args[3].TypeExpression, args[4].Value)
	case _ReduceInitializableTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:607:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceNamedTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:608:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceInferredTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:609:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceImplicitStructTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:610:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceExplicitEnumTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:611:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceImplicitEnumTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:612:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceTraitTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:613:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceFuncTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:614:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceParseErrorTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:615:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceLocalToNamedTypeExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = NamedTypeExprType
		symbol.TypeExpression, err = reducer.LocalToNamedTypeExpr(args[0].Value, args[1].TypeArgumentList)
	case _ReduceExternalToNamedTypeExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = NamedTypeExprType
		symbol.TypeExpression, err = reducer.ExternalToNamedTypeExpr(args[0].Value, args[1].Value, args[2].Value, args[3].TypeArgumentList)
	case _ReduceDotToInferredTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InferredTypeExprType
		symbol.TypeExpression, err = reducer.DotToInferredTypeExpr(args[0].Value)
	case _ReduceUnderscoreToInferredTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InferredTypeExprType
		symbol.TypeExpression, err = reducer.UnderscoreToInferredTypeExpr(args[0].Value)
	case _ReduceToParseErrorTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParseErrorTypeExprType
		symbol.TypeExpression, err = reducer.ToParseErrorTypeExpr(args[0].ParseError)
	case _ReduceAtomTypeExprToReturnableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableTypeExprType
		//line grammar.lr:636:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReducePrefixUnaryTypeExprToReturnableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableTypeExprType
		//line grammar.lr:637:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceToPrefixUnaryTypeExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PrefixUnaryTypeExprType
		symbol.TypeExpression, err = reducer.ToPrefixUnaryTypeExpr(args[0].Value, args[1].TypeExpression)
	case _ReduceQuestionToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:643:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceExclaimToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:644:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:645:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:646:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceTildeTildeToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:647:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceReturnableTypeExprToTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TypeExprType
		//line grammar.lr:652:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceBinaryTypeExprToTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TypeExprType
		//line grammar.lr:653:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceToBinaryTypeExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = BinaryTypeExprType
		symbol.TypeExpression, err = reducer.ToBinaryTypeExpr(args[0].TypeExpression, args[1].Value, args[2].TypeExpression)
	case _ReduceMulToBinaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryTypeOpType
		//line grammar.lr:659:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddToBinaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryTypeOpType
		//line grammar.lr:660:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToBinaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryTypeOpType
		//line grammar.lr:661:4
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
	case _ReduceGenericToGenericParameterDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = GenericParameterDefsType
		symbol.Generic_, err = reducer.GenericToGenericParameterDefs(args[0].Value, args[1].Generic_, args[2].Value)
	case _ReduceNilToGenericParameterDefs:
		symbol.SymbolId_ = GenericParameterDefsType
		symbol.Generic_, err = reducer.NilToGenericParameterDefs()
	case _ReduceAddToProperGenericParameterDefList:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperGenericParameterDefListType
		symbol.Generic_, err = reducer.AddToProperGenericParameterDefList(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceGenericParameterDefToProperGenericParameterDefList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperGenericParameterDefListType
		symbol.Generic_, err = reducer.GenericParameterDefToProperGenericParameterDefList(args[0].Generic_)
	case _ReduceProperGenericParameterDefListToGenericParameterDefList:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = GenericParameterDefListType
		//line grammar.lr:686:4
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _ReduceImproperToGenericParameterDefList:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = GenericParameterDefListType
		symbol.Generic_, err = reducer.ImproperToGenericParameterDefList(args[0].Generic_, args[1].Value)
	case _ReduceNilToGenericParameterDefList:
		symbol.SymbolId_ = GenericParameterDefListType
		symbol.Generic_, err = reducer.NilToGenericParameterDefList()
	case _ReduceBindingToGenericTypeArguments:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = GenericTypeArgumentsType
		symbol.TypeArgumentList, err = reducer.BindingToGenericTypeArguments(args[0].Value, args[1].TypeArgumentList, args[2].Value)
	case _ReduceNilToGenericTypeArguments:
		symbol.SymbolId_ = GenericTypeArgumentsType
		symbol.TypeArgumentList, err = reducer.NilToGenericTypeArguments()
	case _ReduceAddToProperTypeArguments:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperTypeArgumentsType
		symbol.TypeArgumentList, err = reducer.AddToProperTypeArguments(args[0].TypeArgumentList, args[1].Value, args[2].TypeExpression)
	case _ReduceTypeExprToProperTypeArguments:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperTypeArgumentsType
		symbol.TypeArgumentList, err = reducer.TypeExprToProperTypeArguments(args[0].TypeExpression)
	case _ReduceProperTypeArgumentsToTypeArguments:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TypeArgumentsType
		//line grammar.lr:699:4
		symbol.TypeArgumentList = args[0].TypeArgumentList
		err = nil
	case _ReduceImproperToTypeArguments:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = TypeArgumentsType
		symbol.TypeArgumentList, err = reducer.ImproperToTypeArguments(args[0].TypeArgumentList, args[1].Value)
	case _ReduceNilToTypeArguments:
		symbol.SymbolId_ = TypeArgumentsType
		symbol.TypeArgumentList, err = reducer.NilToTypeArguments()
	case _ReduceNamedToFieldDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = FieldDefType
		symbol.Generic_, err = reducer.NamedToFieldDef(args[0].Value, args[1].TypeExpression, args[2].Generic_)
	case _ReduceUnnamedToFieldDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = FieldDefType
		symbol.Generic_, err = reducer.UnnamedToFieldDef(args[0].TypeExpression, args[1].Generic_)
	case _ReduceStructPaddingToFieldDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = FieldDefType
		symbol.Generic_, err = reducer.StructPaddingToFieldDef(args[0].Value, args[1].TypeExpression)
	case _ReduceMethodSignatureToFieldDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldDefType
		//line grammar.lr:718:4
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _ReduceUnsafeStatementToFieldDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = FieldDefType
		symbol.Generic_, err = reducer.UnsafeStatementToFieldDef(args[0].Statement)
	case _ReduceDefaultToOptionalDefault:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = OptionalDefaultType
		symbol.Generic_, err = reducer.DefaultToOptionalDefault(args[0].Value, args[1].Value)
	case _ReduceNilToOptionalDefault:
		symbol.SymbolId_ = OptionalDefaultType
		symbol.Generic_, err = reducer.NilToOptionalDefault()
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
		//line grammar.lr:730:4
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
	case _ReduceToImplicitStructTypeExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitStructTypeExprType
		symbol.TypeExpression, err = reducer.ToImplicitStructTypeExpr(args[0].Value, args[1].Generic_, args[2].Value)
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
		//line grammar.lr:743:4
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
	case _ReduceToExplicitStructTypeExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ExplicitStructTypeExprType
		symbol.TypeExpression, err = reducer.ToExplicitStructTypeExpr(args[0].Value, args[1].Value, args[2].Generic_, args[3].Value)
	case _ReduceToTraitTypeExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = TraitTypeExprType
		symbol.TypeExpression, err = reducer.ToTraitTypeExpr(args[0].Value, args[1].Value, args[2].Generic_, args[3].Value)
	case _ReducePairToProperImplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperImplicitEnumValueDefsType
		symbol.Generic_, err = reducer.PairToProperImplicitEnumValueDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceAddToProperImplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperImplicitEnumValueDefsType
		symbol.Generic_, err = reducer.AddToProperImplicitEnumValueDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceProperImplicitEnumValueDefsToImplicitEnumValueDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ImplicitEnumValueDefsType
		//line grammar.lr:767:4
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _ReduceImproperToImplicitEnumValueDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ImplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ImproperToImplicitEnumValueDefs(args[0].Generic_, args[1].Count)
	case _ReduceToImplicitEnumTypeExpr:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ImplicitEnumTypeExprType
		symbol.TypeExpression, err = reducer.ToImplicitEnumTypeExpr(args[0].Value, args[1].Generic_, args[2].Value)
	case _ReduceExplicitPairToProperExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ExplicitPairToProperExplicitEnumValueDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceImplicitPairToProperExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ImplicitPairToProperExplicitEnumValueDefs(args[0].Generic_, args[1].Count, args[2].Generic_)
	case _ReduceExplicitAddToProperExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ExplicitAddToProperExplicitEnumValueDefs(args[0].Generic_, args[1].Value, args[2].Generic_)
	case _ReduceImplicitAddToProperExplicitEnumValueDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ImplicitAddToProperExplicitEnumValueDefs(args[0].Generic_, args[1].Count, args[2].Generic_)
	case _ReduceProperExplicitEnumValueDefsToExplicitEnumValueDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		//line grammar.lr:781:4
		symbol.Generic_ = args[0].Generic_
		err = nil
	case _ReduceImproperToExplicitEnumValueDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ExplicitEnumValueDefsType
		symbol.Generic_, err = reducer.ImproperToExplicitEnumValueDefs(args[0].Generic_, args[1].Count)
	case _ReduceToExplicitEnumTypeExpr:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = ExplicitEnumTypeExprType
		symbol.TypeExpression, err = reducer.ToExplicitEnumTypeExpr(args[0].Value, args[1].Value, args[2].Generic_, args[3].Value)
	case _ReduceReturnableTypeExprToReturnType:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnTypeType
		symbol.TypeExpression, err = reducer.ReturnableTypeExprToReturnType(args[0].TypeExpression)
	case _ReduceNilToReturnType:
		symbol.SymbolId_ = ReturnTypeType
		symbol.TypeExpression, err = reducer.NilToReturnType()
	case _ReduceNamedTypedArgToProperParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ProperParameterDefType
		symbol.Parameter, err = reducer.NamedTypedArgToProperParameterDef(args[0].Value, args[1].TypeExpression)
	case _ReduceNamedTypedVarargToProperParameterDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperParameterDefType
		symbol.Parameter, err = reducer.NamedTypedVarargToProperParameterDef(args[0].Value, args[1].Value, args[2].TypeExpression)
	case _ReduceNamedInferredVarargToProperParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ProperParameterDefType
		symbol.Parameter, err = reducer.NamedInferredVarargToProperParameterDef(args[0].Value, args[1].Value)
	case _ReduceIgnoreTypedArgToProperParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ProperParameterDefType
		symbol.Parameter, err = reducer.IgnoreTypedArgToProperParameterDef(args[0].Value, args[1].TypeExpression)
	case _ReduceIgnoreInferredVarargToProperParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ProperParameterDefType
		symbol.Parameter, err = reducer.IgnoreInferredVarargToProperParameterDef(args[0].Value, args[1].Value)
	case _ReduceIgnoreTypedVarargToProperParameterDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperParameterDefType
		symbol.Parameter, err = reducer.IgnoreTypedVarargToProperParameterDef(args[0].Value, args[1].Value, args[2].TypeExpression)
	case _ReduceProperParameterDefToParameterDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDeclType
		//line grammar.lr:808:4
		symbol.Parameter = args[0].Parameter
		err = nil
	case _ReduceUnnamedTypedArgToParameterDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Parameter, err = reducer.UnnamedTypedArgToParameterDecl(args[0].TypeExpression)
	case _ReduceUnnamedInferredVarargToParameterDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Parameter, err = reducer.UnnamedInferredVarargToParameterDecl(args[0].Value)
	case _ReduceUnnamedTypedVarargToParameterDecl:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Parameter, err = reducer.UnnamedTypedVarargToParameterDecl(args[0].Value, args[1].TypeExpression)
	case _ReduceProperParameterDefToParameterDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDefType
		//line grammar.lr:818:4
		symbol.Parameter = args[0].Parameter
		err = nil
	case _ReduceNamedInferredArgToParameterDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDefType
		symbol.Parameter, err = reducer.NamedInferredArgToParameterDef(args[0].Value)
	case _ReduceIgnoreInferredArgToParameterDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDefType
		symbol.Parameter, err = reducer.IgnoreInferredArgToParameterDef(args[0].Value)
	case _ReduceAddToProperParameterDecls:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperParameterDeclsType
		symbol.Parameters, err = reducer.AddToProperParameterDecls(args[0].Parameters, args[1].Value, args[2].Parameter)
	case _ReduceParameterDeclToProperParameterDecls:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperParameterDeclsType
		symbol.Parameters, err = reducer.ParameterDeclToProperParameterDecls(args[0].Parameter)
	case _ReduceProperParameterDeclsToParameterDecls:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDeclsType
		//line grammar.lr:827:4
		symbol.Parameters = args[0].Parameters
		err = nil
	case _ReduceImproperToParameterDecls:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDeclsType
		symbol.Parameters, err = reducer.ImproperToParameterDecls(args[0].Parameters, args[1].Value)
	case _ReduceNilToParameterDecls:
		symbol.SymbolId_ = ParameterDeclsType
		symbol.Parameters, err = reducer.NilToParameterDecls()
	case _ReduceAddToProperParameterDefs:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ProperParameterDefsType
		symbol.Parameters, err = reducer.AddToProperParameterDefs(args[0].Parameters, args[1].Value, args[2].Parameter)
	case _ReduceParameterDefToProperParameterDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ProperParameterDefsType
		symbol.Parameters, err = reducer.ParameterDefToProperParameterDefs(args[0].Parameter)
	case _ReduceProperParameterDefsToParameterDefs:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDefsType
		//line grammar.lr:836:4
		symbol.Parameters = args[0].Parameters
		err = nil
	case _ReduceImproperToParameterDefs:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDefsType
		symbol.Parameters, err = reducer.ImproperToParameterDefs(args[0].Parameters, args[1].Value)
	case _ReduceNilToParameterDefs:
		symbol.SymbolId_ = ParameterDefsType
		symbol.Parameters, err = reducer.NilToParameterDefs()
	case _ReduceToFuncTypeExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = FuncTypeExprType
		symbol.TypeExpression, err = reducer.ToFuncTypeExpr(args[0].Value, args[1].Value, args[2].Parameters, args[3].Value, args[4].TypeExpression)
	case _ReduceToMethodSignature:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = MethodSignatureType
		symbol.Generic_, err = reducer.ToMethodSignature(args[0].Value, args[1].Value, args[2].Value, args[3].Parameters, args[4].Value, args[5].TypeExpression)
	case _ReduceFuncDefToNamedFuncDef:
		args := stack[len(stack)-8:]
		stack = stack[:len(stack)-8]
		symbol.SymbolId_ = NamedFuncDefType
		symbol.SourceDefinition, err = reducer.FuncDefToNamedFuncDef(args[0].Value, args[1].Value, args[2].Generic_, args[3].Value, args[4].Parameters, args[5].Value, args[6].TypeExpression, args[7].Generic_)
	case _ReduceMethodDefToNamedFuncDef:
		args := stack[len(stack)-10:]
		stack = stack[:len(stack)-10]
		symbol.SymbolId_ = NamedFuncDefType
		symbol.SourceDefinition, err = reducer.MethodDefToNamedFuncDef(args[0].Value, args[1].Value, args[2].Parameter, args[3].Value, args[4].Value, args[5].Value, args[6].Parameters, args[7].Value, args[8].TypeExpression, args[9].Generic_)
	case _ReduceAliasToNamedFuncDef:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = NamedFuncDefType
		symbol.SourceDefinition, err = reducer.AliasToNamedFuncDef(args[0].Value, args[1].Value, args[2].Value, args[3].Expression)
	case _ReduceToAnonymousFuncExpr:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = AnonymousFuncExprType
		symbol.Expression, err = reducer.ToAnonymousFuncExpr(args[0].Value, args[1].Value, args[2].Parameters, args[3].Value, args[4].TypeExpression, args[5].Generic_)
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
	_GotoState1Action                                                   = &_Action{_ShiftAction, _State1, 0}
	_GotoState2Action                                                   = &_Action{_ShiftAction, _State2, 0}
	_GotoState3Action                                                   = &_Action{_ShiftAction, _State3, 0}
	_GotoState4Action                                                   = &_Action{_ShiftAction, _State4, 0}
	_GotoState5Action                                                   = &_Action{_ShiftAction, _State5, 0}
	_GotoState6Action                                                   = &_Action{_ShiftAction, _State6, 0}
	_GotoState7Action                                                   = &_Action{_ShiftAction, _State7, 0}
	_GotoState8Action                                                   = &_Action{_ShiftAction, _State8, 0}
	_GotoState9Action                                                   = &_Action{_ShiftAction, _State9, 0}
	_GotoState10Action                                                  = &_Action{_ShiftAction, _State10, 0}
	_GotoState11Action                                                  = &_Action{_ShiftAction, _State11, 0}
	_GotoState12Action                                                  = &_Action{_ShiftAction, _State12, 0}
	_GotoState13Action                                                  = &_Action{_ShiftAction, _State13, 0}
	_GotoState14Action                                                  = &_Action{_ShiftAction, _State14, 0}
	_GotoState15Action                                                  = &_Action{_ShiftAction, _State15, 0}
	_GotoState16Action                                                  = &_Action{_ShiftAction, _State16, 0}
	_GotoState17Action                                                  = &_Action{_ShiftAction, _State17, 0}
	_GotoState18Action                                                  = &_Action{_ShiftAction, _State18, 0}
	_GotoState19Action                                                  = &_Action{_ShiftAction, _State19, 0}
	_GotoState20Action                                                  = &_Action{_ShiftAction, _State20, 0}
	_GotoState21Action                                                  = &_Action{_ShiftAction, _State21, 0}
	_GotoState22Action                                                  = &_Action{_ShiftAction, _State22, 0}
	_GotoState23Action                                                  = &_Action{_ShiftAction, _State23, 0}
	_GotoState24Action                                                  = &_Action{_ShiftAction, _State24, 0}
	_GotoState25Action                                                  = &_Action{_ShiftAction, _State25, 0}
	_GotoState26Action                                                  = &_Action{_ShiftAction, _State26, 0}
	_GotoState27Action                                                  = &_Action{_ShiftAction, _State27, 0}
	_GotoState28Action                                                  = &_Action{_ShiftAction, _State28, 0}
	_GotoState29Action                                                  = &_Action{_ShiftAction, _State29, 0}
	_GotoState30Action                                                  = &_Action{_ShiftAction, _State30, 0}
	_GotoState31Action                                                  = &_Action{_ShiftAction, _State31, 0}
	_GotoState32Action                                                  = &_Action{_ShiftAction, _State32, 0}
	_GotoState33Action                                                  = &_Action{_ShiftAction, _State33, 0}
	_GotoState34Action                                                  = &_Action{_ShiftAction, _State34, 0}
	_GotoState35Action                                                  = &_Action{_ShiftAction, _State35, 0}
	_GotoState36Action                                                  = &_Action{_ShiftAction, _State36, 0}
	_GotoState37Action                                                  = &_Action{_ShiftAction, _State37, 0}
	_GotoState38Action                                                  = &_Action{_ShiftAction, _State38, 0}
	_GotoState39Action                                                  = &_Action{_ShiftAction, _State39, 0}
	_GotoState40Action                                                  = &_Action{_ShiftAction, _State40, 0}
	_GotoState41Action                                                  = &_Action{_ShiftAction, _State41, 0}
	_GotoState42Action                                                  = &_Action{_ShiftAction, _State42, 0}
	_GotoState43Action                                                  = &_Action{_ShiftAction, _State43, 0}
	_GotoState44Action                                                  = &_Action{_ShiftAction, _State44, 0}
	_GotoState45Action                                                  = &_Action{_ShiftAction, _State45, 0}
	_GotoState46Action                                                  = &_Action{_ShiftAction, _State46, 0}
	_GotoState47Action                                                  = &_Action{_ShiftAction, _State47, 0}
	_GotoState48Action                                                  = &_Action{_ShiftAction, _State48, 0}
	_GotoState49Action                                                  = &_Action{_ShiftAction, _State49, 0}
	_GotoState50Action                                                  = &_Action{_ShiftAction, _State50, 0}
	_GotoState51Action                                                  = &_Action{_ShiftAction, _State51, 0}
	_GotoState52Action                                                  = &_Action{_ShiftAction, _State52, 0}
	_GotoState53Action                                                  = &_Action{_ShiftAction, _State53, 0}
	_GotoState54Action                                                  = &_Action{_ShiftAction, _State54, 0}
	_GotoState55Action                                                  = &_Action{_ShiftAction, _State55, 0}
	_GotoState56Action                                                  = &_Action{_ShiftAction, _State56, 0}
	_GotoState57Action                                                  = &_Action{_ShiftAction, _State57, 0}
	_GotoState58Action                                                  = &_Action{_ShiftAction, _State58, 0}
	_GotoState59Action                                                  = &_Action{_ShiftAction, _State59, 0}
	_GotoState60Action                                                  = &_Action{_ShiftAction, _State60, 0}
	_GotoState61Action                                                  = &_Action{_ShiftAction, _State61, 0}
	_GotoState62Action                                                  = &_Action{_ShiftAction, _State62, 0}
	_GotoState63Action                                                  = &_Action{_ShiftAction, _State63, 0}
	_GotoState64Action                                                  = &_Action{_ShiftAction, _State64, 0}
	_GotoState65Action                                                  = &_Action{_ShiftAction, _State65, 0}
	_GotoState66Action                                                  = &_Action{_ShiftAction, _State66, 0}
	_GotoState67Action                                                  = &_Action{_ShiftAction, _State67, 0}
	_GotoState68Action                                                  = &_Action{_ShiftAction, _State68, 0}
	_GotoState69Action                                                  = &_Action{_ShiftAction, _State69, 0}
	_GotoState70Action                                                  = &_Action{_ShiftAction, _State70, 0}
	_GotoState71Action                                                  = &_Action{_ShiftAction, _State71, 0}
	_GotoState72Action                                                  = &_Action{_ShiftAction, _State72, 0}
	_GotoState73Action                                                  = &_Action{_ShiftAction, _State73, 0}
	_GotoState74Action                                                  = &_Action{_ShiftAction, _State74, 0}
	_GotoState75Action                                                  = &_Action{_ShiftAction, _State75, 0}
	_GotoState76Action                                                  = &_Action{_ShiftAction, _State76, 0}
	_GotoState77Action                                                  = &_Action{_ShiftAction, _State77, 0}
	_GotoState78Action                                                  = &_Action{_ShiftAction, _State78, 0}
	_GotoState79Action                                                  = &_Action{_ShiftAction, _State79, 0}
	_GotoState80Action                                                  = &_Action{_ShiftAction, _State80, 0}
	_GotoState81Action                                                  = &_Action{_ShiftAction, _State81, 0}
	_GotoState82Action                                                  = &_Action{_ShiftAction, _State82, 0}
	_GotoState83Action                                                  = &_Action{_ShiftAction, _State83, 0}
	_GotoState84Action                                                  = &_Action{_ShiftAction, _State84, 0}
	_GotoState85Action                                                  = &_Action{_ShiftAction, _State85, 0}
	_GotoState86Action                                                  = &_Action{_ShiftAction, _State86, 0}
	_GotoState87Action                                                  = &_Action{_ShiftAction, _State87, 0}
	_GotoState88Action                                                  = &_Action{_ShiftAction, _State88, 0}
	_GotoState89Action                                                  = &_Action{_ShiftAction, _State89, 0}
	_GotoState90Action                                                  = &_Action{_ShiftAction, _State90, 0}
	_GotoState91Action                                                  = &_Action{_ShiftAction, _State91, 0}
	_GotoState92Action                                                  = &_Action{_ShiftAction, _State92, 0}
	_GotoState93Action                                                  = &_Action{_ShiftAction, _State93, 0}
	_GotoState94Action                                                  = &_Action{_ShiftAction, _State94, 0}
	_GotoState95Action                                                  = &_Action{_ShiftAction, _State95, 0}
	_GotoState96Action                                                  = &_Action{_ShiftAction, _State96, 0}
	_GotoState97Action                                                  = &_Action{_ShiftAction, _State97, 0}
	_GotoState98Action                                                  = &_Action{_ShiftAction, _State98, 0}
	_GotoState99Action                                                  = &_Action{_ShiftAction, _State99, 0}
	_GotoState100Action                                                 = &_Action{_ShiftAction, _State100, 0}
	_GotoState101Action                                                 = &_Action{_ShiftAction, _State101, 0}
	_GotoState102Action                                                 = &_Action{_ShiftAction, _State102, 0}
	_GotoState103Action                                                 = &_Action{_ShiftAction, _State103, 0}
	_GotoState104Action                                                 = &_Action{_ShiftAction, _State104, 0}
	_GotoState105Action                                                 = &_Action{_ShiftAction, _State105, 0}
	_GotoState106Action                                                 = &_Action{_ShiftAction, _State106, 0}
	_GotoState107Action                                                 = &_Action{_ShiftAction, _State107, 0}
	_GotoState108Action                                                 = &_Action{_ShiftAction, _State108, 0}
	_GotoState109Action                                                 = &_Action{_ShiftAction, _State109, 0}
	_GotoState110Action                                                 = &_Action{_ShiftAction, _State110, 0}
	_GotoState111Action                                                 = &_Action{_ShiftAction, _State111, 0}
	_GotoState112Action                                                 = &_Action{_ShiftAction, _State112, 0}
	_GotoState113Action                                                 = &_Action{_ShiftAction, _State113, 0}
	_GotoState114Action                                                 = &_Action{_ShiftAction, _State114, 0}
	_GotoState115Action                                                 = &_Action{_ShiftAction, _State115, 0}
	_GotoState116Action                                                 = &_Action{_ShiftAction, _State116, 0}
	_GotoState117Action                                                 = &_Action{_ShiftAction, _State117, 0}
	_GotoState118Action                                                 = &_Action{_ShiftAction, _State118, 0}
	_GotoState119Action                                                 = &_Action{_ShiftAction, _State119, 0}
	_GotoState120Action                                                 = &_Action{_ShiftAction, _State120, 0}
	_GotoState121Action                                                 = &_Action{_ShiftAction, _State121, 0}
	_GotoState122Action                                                 = &_Action{_ShiftAction, _State122, 0}
	_GotoState123Action                                                 = &_Action{_ShiftAction, _State123, 0}
	_GotoState124Action                                                 = &_Action{_ShiftAction, _State124, 0}
	_GotoState125Action                                                 = &_Action{_ShiftAction, _State125, 0}
	_GotoState126Action                                                 = &_Action{_ShiftAction, _State126, 0}
	_GotoState127Action                                                 = &_Action{_ShiftAction, _State127, 0}
	_GotoState128Action                                                 = &_Action{_ShiftAction, _State128, 0}
	_GotoState129Action                                                 = &_Action{_ShiftAction, _State129, 0}
	_GotoState130Action                                                 = &_Action{_ShiftAction, _State130, 0}
	_GotoState131Action                                                 = &_Action{_ShiftAction, _State131, 0}
	_GotoState132Action                                                 = &_Action{_ShiftAction, _State132, 0}
	_GotoState133Action                                                 = &_Action{_ShiftAction, _State133, 0}
	_GotoState134Action                                                 = &_Action{_ShiftAction, _State134, 0}
	_GotoState135Action                                                 = &_Action{_ShiftAction, _State135, 0}
	_GotoState136Action                                                 = &_Action{_ShiftAction, _State136, 0}
	_GotoState137Action                                                 = &_Action{_ShiftAction, _State137, 0}
	_GotoState138Action                                                 = &_Action{_ShiftAction, _State138, 0}
	_GotoState139Action                                                 = &_Action{_ShiftAction, _State139, 0}
	_GotoState140Action                                                 = &_Action{_ShiftAction, _State140, 0}
	_GotoState141Action                                                 = &_Action{_ShiftAction, _State141, 0}
	_GotoState142Action                                                 = &_Action{_ShiftAction, _State142, 0}
	_GotoState143Action                                                 = &_Action{_ShiftAction, _State143, 0}
	_GotoState144Action                                                 = &_Action{_ShiftAction, _State144, 0}
	_GotoState145Action                                                 = &_Action{_ShiftAction, _State145, 0}
	_GotoState146Action                                                 = &_Action{_ShiftAction, _State146, 0}
	_GotoState147Action                                                 = &_Action{_ShiftAction, _State147, 0}
	_GotoState148Action                                                 = &_Action{_ShiftAction, _State148, 0}
	_GotoState149Action                                                 = &_Action{_ShiftAction, _State149, 0}
	_GotoState150Action                                                 = &_Action{_ShiftAction, _State150, 0}
	_GotoState151Action                                                 = &_Action{_ShiftAction, _State151, 0}
	_GotoState152Action                                                 = &_Action{_ShiftAction, _State152, 0}
	_GotoState153Action                                                 = &_Action{_ShiftAction, _State153, 0}
	_GotoState154Action                                                 = &_Action{_ShiftAction, _State154, 0}
	_GotoState155Action                                                 = &_Action{_ShiftAction, _State155, 0}
	_GotoState156Action                                                 = &_Action{_ShiftAction, _State156, 0}
	_GotoState157Action                                                 = &_Action{_ShiftAction, _State157, 0}
	_GotoState158Action                                                 = &_Action{_ShiftAction, _State158, 0}
	_GotoState159Action                                                 = &_Action{_ShiftAction, _State159, 0}
	_GotoState160Action                                                 = &_Action{_ShiftAction, _State160, 0}
	_GotoState161Action                                                 = &_Action{_ShiftAction, _State161, 0}
	_GotoState162Action                                                 = &_Action{_ShiftAction, _State162, 0}
	_GotoState163Action                                                 = &_Action{_ShiftAction, _State163, 0}
	_GotoState164Action                                                 = &_Action{_ShiftAction, _State164, 0}
	_GotoState165Action                                                 = &_Action{_ShiftAction, _State165, 0}
	_GotoState166Action                                                 = &_Action{_ShiftAction, _State166, 0}
	_GotoState167Action                                                 = &_Action{_ShiftAction, _State167, 0}
	_GotoState168Action                                                 = &_Action{_ShiftAction, _State168, 0}
	_GotoState169Action                                                 = &_Action{_ShiftAction, _State169, 0}
	_GotoState170Action                                                 = &_Action{_ShiftAction, _State170, 0}
	_GotoState171Action                                                 = &_Action{_ShiftAction, _State171, 0}
	_GotoState172Action                                                 = &_Action{_ShiftAction, _State172, 0}
	_GotoState173Action                                                 = &_Action{_ShiftAction, _State173, 0}
	_GotoState174Action                                                 = &_Action{_ShiftAction, _State174, 0}
	_GotoState175Action                                                 = &_Action{_ShiftAction, _State175, 0}
	_GotoState176Action                                                 = &_Action{_ShiftAction, _State176, 0}
	_GotoState177Action                                                 = &_Action{_ShiftAction, _State177, 0}
	_GotoState178Action                                                 = &_Action{_ShiftAction, _State178, 0}
	_GotoState179Action                                                 = &_Action{_ShiftAction, _State179, 0}
	_GotoState180Action                                                 = &_Action{_ShiftAction, _State180, 0}
	_GotoState181Action                                                 = &_Action{_ShiftAction, _State181, 0}
	_GotoState182Action                                                 = &_Action{_ShiftAction, _State182, 0}
	_GotoState183Action                                                 = &_Action{_ShiftAction, _State183, 0}
	_GotoState184Action                                                 = &_Action{_ShiftAction, _State184, 0}
	_GotoState185Action                                                 = &_Action{_ShiftAction, _State185, 0}
	_GotoState186Action                                                 = &_Action{_ShiftAction, _State186, 0}
	_GotoState187Action                                                 = &_Action{_ShiftAction, _State187, 0}
	_GotoState188Action                                                 = &_Action{_ShiftAction, _State188, 0}
	_GotoState189Action                                                 = &_Action{_ShiftAction, _State189, 0}
	_GotoState190Action                                                 = &_Action{_ShiftAction, _State190, 0}
	_GotoState191Action                                                 = &_Action{_ShiftAction, _State191, 0}
	_GotoState192Action                                                 = &_Action{_ShiftAction, _State192, 0}
	_GotoState193Action                                                 = &_Action{_ShiftAction, _State193, 0}
	_GotoState194Action                                                 = &_Action{_ShiftAction, _State194, 0}
	_GotoState195Action                                                 = &_Action{_ShiftAction, _State195, 0}
	_GotoState196Action                                                 = &_Action{_ShiftAction, _State196, 0}
	_GotoState197Action                                                 = &_Action{_ShiftAction, _State197, 0}
	_GotoState198Action                                                 = &_Action{_ShiftAction, _State198, 0}
	_GotoState199Action                                                 = &_Action{_ShiftAction, _State199, 0}
	_GotoState200Action                                                 = &_Action{_ShiftAction, _State200, 0}
	_GotoState201Action                                                 = &_Action{_ShiftAction, _State201, 0}
	_GotoState202Action                                                 = &_Action{_ShiftAction, _State202, 0}
	_GotoState203Action                                                 = &_Action{_ShiftAction, _State203, 0}
	_GotoState204Action                                                 = &_Action{_ShiftAction, _State204, 0}
	_GotoState205Action                                                 = &_Action{_ShiftAction, _State205, 0}
	_GotoState206Action                                                 = &_Action{_ShiftAction, _State206, 0}
	_GotoState207Action                                                 = &_Action{_ShiftAction, _State207, 0}
	_GotoState208Action                                                 = &_Action{_ShiftAction, _State208, 0}
	_GotoState209Action                                                 = &_Action{_ShiftAction, _State209, 0}
	_GotoState210Action                                                 = &_Action{_ShiftAction, _State210, 0}
	_GotoState211Action                                                 = &_Action{_ShiftAction, _State211, 0}
	_GotoState212Action                                                 = &_Action{_ShiftAction, _State212, 0}
	_GotoState213Action                                                 = &_Action{_ShiftAction, _State213, 0}
	_GotoState214Action                                                 = &_Action{_ShiftAction, _State214, 0}
	_GotoState215Action                                                 = &_Action{_ShiftAction, _State215, 0}
	_GotoState216Action                                                 = &_Action{_ShiftAction, _State216, 0}
	_GotoState217Action                                                 = &_Action{_ShiftAction, _State217, 0}
	_GotoState218Action                                                 = &_Action{_ShiftAction, _State218, 0}
	_GotoState219Action                                                 = &_Action{_ShiftAction, _State219, 0}
	_GotoState220Action                                                 = &_Action{_ShiftAction, _State220, 0}
	_GotoState221Action                                                 = &_Action{_ShiftAction, _State221, 0}
	_GotoState222Action                                                 = &_Action{_ShiftAction, _State222, 0}
	_GotoState223Action                                                 = &_Action{_ShiftAction, _State223, 0}
	_GotoState224Action                                                 = &_Action{_ShiftAction, _State224, 0}
	_GotoState225Action                                                 = &_Action{_ShiftAction, _State225, 0}
	_GotoState226Action                                                 = &_Action{_ShiftAction, _State226, 0}
	_GotoState227Action                                                 = &_Action{_ShiftAction, _State227, 0}
	_GotoState228Action                                                 = &_Action{_ShiftAction, _State228, 0}
	_GotoState229Action                                                 = &_Action{_ShiftAction, _State229, 0}
	_GotoState230Action                                                 = &_Action{_ShiftAction, _State230, 0}
	_GotoState231Action                                                 = &_Action{_ShiftAction, _State231, 0}
	_GotoState232Action                                                 = &_Action{_ShiftAction, _State232, 0}
	_GotoState233Action                                                 = &_Action{_ShiftAction, _State233, 0}
	_GotoState234Action                                                 = &_Action{_ShiftAction, _State234, 0}
	_GotoState235Action                                                 = &_Action{_ShiftAction, _State235, 0}
	_GotoState236Action                                                 = &_Action{_ShiftAction, _State236, 0}
	_GotoState237Action                                                 = &_Action{_ShiftAction, _State237, 0}
	_GotoState238Action                                                 = &_Action{_ShiftAction, _State238, 0}
	_GotoState239Action                                                 = &_Action{_ShiftAction, _State239, 0}
	_GotoState240Action                                                 = &_Action{_ShiftAction, _State240, 0}
	_GotoState241Action                                                 = &_Action{_ShiftAction, _State241, 0}
	_GotoState242Action                                                 = &_Action{_ShiftAction, _State242, 0}
	_GotoState243Action                                                 = &_Action{_ShiftAction, _State243, 0}
	_GotoState244Action                                                 = &_Action{_ShiftAction, _State244, 0}
	_GotoState245Action                                                 = &_Action{_ShiftAction, _State245, 0}
	_GotoState246Action                                                 = &_Action{_ShiftAction, _State246, 0}
	_GotoState247Action                                                 = &_Action{_ShiftAction, _State247, 0}
	_GotoState248Action                                                 = &_Action{_ShiftAction, _State248, 0}
	_GotoState249Action                                                 = &_Action{_ShiftAction, _State249, 0}
	_GotoState250Action                                                 = &_Action{_ShiftAction, _State250, 0}
	_GotoState251Action                                                 = &_Action{_ShiftAction, _State251, 0}
	_GotoState252Action                                                 = &_Action{_ShiftAction, _State252, 0}
	_GotoState253Action                                                 = &_Action{_ShiftAction, _State253, 0}
	_GotoState254Action                                                 = &_Action{_ShiftAction, _State254, 0}
	_GotoState255Action                                                 = &_Action{_ShiftAction, _State255, 0}
	_GotoState256Action                                                 = &_Action{_ShiftAction, _State256, 0}
	_GotoState257Action                                                 = &_Action{_ShiftAction, _State257, 0}
	_GotoState258Action                                                 = &_Action{_ShiftAction, _State258, 0}
	_GotoState259Action                                                 = &_Action{_ShiftAction, _State259, 0}
	_GotoState260Action                                                 = &_Action{_ShiftAction, _State260, 0}
	_GotoState261Action                                                 = &_Action{_ShiftAction, _State261, 0}
	_GotoState262Action                                                 = &_Action{_ShiftAction, _State262, 0}
	_GotoState263Action                                                 = &_Action{_ShiftAction, _State263, 0}
	_GotoState264Action                                                 = &_Action{_ShiftAction, _State264, 0}
	_GotoState265Action                                                 = &_Action{_ShiftAction, _State265, 0}
	_GotoState266Action                                                 = &_Action{_ShiftAction, _State266, 0}
	_GotoState267Action                                                 = &_Action{_ShiftAction, _State267, 0}
	_GotoState268Action                                                 = &_Action{_ShiftAction, _State268, 0}
	_GotoState269Action                                                 = &_Action{_ShiftAction, _State269, 0}
	_GotoState270Action                                                 = &_Action{_ShiftAction, _State270, 0}
	_GotoState271Action                                                 = &_Action{_ShiftAction, _State271, 0}
	_GotoState272Action                                                 = &_Action{_ShiftAction, _State272, 0}
	_GotoState273Action                                                 = &_Action{_ShiftAction, _State273, 0}
	_GotoState274Action                                                 = &_Action{_ShiftAction, _State274, 0}
	_GotoState275Action                                                 = &_Action{_ShiftAction, _State275, 0}
	_GotoState276Action                                                 = &_Action{_ShiftAction, _State276, 0}
	_GotoState277Action                                                 = &_Action{_ShiftAction, _State277, 0}
	_GotoState278Action                                                 = &_Action{_ShiftAction, _State278, 0}
	_GotoState279Action                                                 = &_Action{_ShiftAction, _State279, 0}
	_GotoState280Action                                                 = &_Action{_ShiftAction, _State280, 0}
	_GotoState281Action                                                 = &_Action{_ShiftAction, _State281, 0}
	_GotoState282Action                                                 = &_Action{_ShiftAction, _State282, 0}
	_GotoState283Action                                                 = &_Action{_ShiftAction, _State283, 0}
	_GotoState284Action                                                 = &_Action{_ShiftAction, _State284, 0}
	_GotoState285Action                                                 = &_Action{_ShiftAction, _State285, 0}
	_GotoState286Action                                                 = &_Action{_ShiftAction, _State286, 0}
	_GotoState287Action                                                 = &_Action{_ShiftAction, _State287, 0}
	_GotoState288Action                                                 = &_Action{_ShiftAction, _State288, 0}
	_GotoState289Action                                                 = &_Action{_ShiftAction, _State289, 0}
	_GotoState290Action                                                 = &_Action{_ShiftAction, _State290, 0}
	_GotoState291Action                                                 = &_Action{_ShiftAction, _State291, 0}
	_GotoState292Action                                                 = &_Action{_ShiftAction, _State292, 0}
	_GotoState293Action                                                 = &_Action{_ShiftAction, _State293, 0}
	_GotoState294Action                                                 = &_Action{_ShiftAction, _State294, 0}
	_GotoState295Action                                                 = &_Action{_ShiftAction, _State295, 0}
	_GotoState296Action                                                 = &_Action{_ShiftAction, _State296, 0}
	_GotoState297Action                                                 = &_Action{_ShiftAction, _State297, 0}
	_GotoState298Action                                                 = &_Action{_ShiftAction, _State298, 0}
	_GotoState299Action                                                 = &_Action{_ShiftAction, _State299, 0}
	_GotoState300Action                                                 = &_Action{_ShiftAction, _State300, 0}
	_GotoState301Action                                                 = &_Action{_ShiftAction, _State301, 0}
	_GotoState302Action                                                 = &_Action{_ShiftAction, _State302, 0}
	_GotoState303Action                                                 = &_Action{_ShiftAction, _State303, 0}
	_GotoState304Action                                                 = &_Action{_ShiftAction, _State304, 0}
	_GotoState305Action                                                 = &_Action{_ShiftAction, _State305, 0}
	_GotoState306Action                                                 = &_Action{_ShiftAction, _State306, 0}
	_GotoState307Action                                                 = &_Action{_ShiftAction, _State307, 0}
	_GotoState308Action                                                 = &_Action{_ShiftAction, _State308, 0}
	_GotoState309Action                                                 = &_Action{_ShiftAction, _State309, 0}
	_GotoState310Action                                                 = &_Action{_ShiftAction, _State310, 0}
	_GotoState311Action                                                 = &_Action{_ShiftAction, _State311, 0}
	_GotoState312Action                                                 = &_Action{_ShiftAction, _State312, 0}
	_GotoState313Action                                                 = &_Action{_ShiftAction, _State313, 0}
	_GotoState314Action                                                 = &_Action{_ShiftAction, _State314, 0}
	_GotoState315Action                                                 = &_Action{_ShiftAction, _State315, 0}
	_GotoState316Action                                                 = &_Action{_ShiftAction, _State316, 0}
	_GotoState317Action                                                 = &_Action{_ShiftAction, _State317, 0}
	_GotoState318Action                                                 = &_Action{_ShiftAction, _State318, 0}
	_GotoState319Action                                                 = &_Action{_ShiftAction, _State319, 0}
	_GotoState320Action                                                 = &_Action{_ShiftAction, _State320, 0}
	_GotoState321Action                                                 = &_Action{_ShiftAction, _State321, 0}
	_GotoState322Action                                                 = &_Action{_ShiftAction, _State322, 0}
	_GotoState323Action                                                 = &_Action{_ShiftAction, _State323, 0}
	_GotoState324Action                                                 = &_Action{_ShiftAction, _State324, 0}
	_GotoState325Action                                                 = &_Action{_ShiftAction, _State325, 0}
	_GotoState326Action                                                 = &_Action{_ShiftAction, _State326, 0}
	_GotoState327Action                                                 = &_Action{_ShiftAction, _State327, 0}
	_GotoState328Action                                                 = &_Action{_ShiftAction, _State328, 0}
	_GotoState329Action                                                 = &_Action{_ShiftAction, _State329, 0}
	_GotoState330Action                                                 = &_Action{_ShiftAction, _State330, 0}
	_GotoState331Action                                                 = &_Action{_ShiftAction, _State331, 0}
	_GotoState332Action                                                 = &_Action{_ShiftAction, _State332, 0}
	_GotoState333Action                                                 = &_Action{_ShiftAction, _State333, 0}
	_GotoState334Action                                                 = &_Action{_ShiftAction, _State334, 0}
	_GotoState335Action                                                 = &_Action{_ShiftAction, _State335, 0}
	_GotoState336Action                                                 = &_Action{_ShiftAction, _State336, 0}
	_GotoState337Action                                                 = &_Action{_ShiftAction, _State337, 0}
	_GotoState338Action                                                 = &_Action{_ShiftAction, _State338, 0}
	_GotoState339Action                                                 = &_Action{_ShiftAction, _State339, 0}
	_GotoState340Action                                                 = &_Action{_ShiftAction, _State340, 0}
	_GotoState341Action                                                 = &_Action{_ShiftAction, _State341, 0}
	_GotoState342Action                                                 = &_Action{_ShiftAction, _State342, 0}
	_GotoState343Action                                                 = &_Action{_ShiftAction, _State343, 0}
	_GotoState344Action                                                 = &_Action{_ShiftAction, _State344, 0}
	_GotoState345Action                                                 = &_Action{_ShiftAction, _State345, 0}
	_GotoState346Action                                                 = &_Action{_ShiftAction, _State346, 0}
	_GotoState347Action                                                 = &_Action{_ShiftAction, _State347, 0}
	_GotoState348Action                                                 = &_Action{_ShiftAction, _State348, 0}
	_GotoState349Action                                                 = &_Action{_ShiftAction, _State349, 0}
	_GotoState350Action                                                 = &_Action{_ShiftAction, _State350, 0}
	_GotoState351Action                                                 = &_Action{_ShiftAction, _State351, 0}
	_GotoState352Action                                                 = &_Action{_ShiftAction, _State352, 0}
	_GotoState353Action                                                 = &_Action{_ShiftAction, _State353, 0}
	_GotoState354Action                                                 = &_Action{_ShiftAction, _State354, 0}
	_GotoState355Action                                                 = &_Action{_ShiftAction, _State355, 0}
	_GotoState356Action                                                 = &_Action{_ShiftAction, _State356, 0}
	_GotoState357Action                                                 = &_Action{_ShiftAction, _State357, 0}
	_GotoState358Action                                                 = &_Action{_ShiftAction, _State358, 0}
	_GotoState359Action                                                 = &_Action{_ShiftAction, _State359, 0}
	_GotoState360Action                                                 = &_Action{_ShiftAction, _State360, 0}
	_GotoState361Action                                                 = &_Action{_ShiftAction, _State361, 0}
	_GotoState362Action                                                 = &_Action{_ShiftAction, _State362, 0}
	_GotoState363Action                                                 = &_Action{_ShiftAction, _State363, 0}
	_GotoState364Action                                                 = &_Action{_ShiftAction, _State364, 0}
	_GotoState365Action                                                 = &_Action{_ShiftAction, _State365, 0}
	_GotoState366Action                                                 = &_Action{_ShiftAction, _State366, 0}
	_GotoState367Action                                                 = &_Action{_ShiftAction, _State367, 0}
	_GotoState368Action                                                 = &_Action{_ShiftAction, _State368, 0}
	_GotoState369Action                                                 = &_Action{_ShiftAction, _State369, 0}
	_GotoState370Action                                                 = &_Action{_ShiftAction, _State370, 0}
	_GotoState371Action                                                 = &_Action{_ShiftAction, _State371, 0}
	_GotoState372Action                                                 = &_Action{_ShiftAction, _State372, 0}
	_GotoState373Action                                                 = &_Action{_ShiftAction, _State373, 0}
	_GotoState374Action                                                 = &_Action{_ShiftAction, _State374, 0}
	_GotoState375Action                                                 = &_Action{_ShiftAction, _State375, 0}
	_GotoState376Action                                                 = &_Action{_ShiftAction, _State376, 0}
	_GotoState377Action                                                 = &_Action{_ShiftAction, _State377, 0}
	_GotoState378Action                                                 = &_Action{_ShiftAction, _State378, 0}
	_GotoState379Action                                                 = &_Action{_ShiftAction, _State379, 0}
	_GotoState380Action                                                 = &_Action{_ShiftAction, _State380, 0}
	_GotoState381Action                                                 = &_Action{_ShiftAction, _State381, 0}
	_GotoState382Action                                                 = &_Action{_ShiftAction, _State382, 0}
	_GotoState383Action                                                 = &_Action{_ShiftAction, _State383, 0}
	_GotoState384Action                                                 = &_Action{_ShiftAction, _State384, 0}
	_GotoState385Action                                                 = &_Action{_ShiftAction, _State385, 0}
	_GotoState386Action                                                 = &_Action{_ShiftAction, _State386, 0}
	_GotoState387Action                                                 = &_Action{_ShiftAction, _State387, 0}
	_GotoState388Action                                                 = &_Action{_ShiftAction, _State388, 0}
	_GotoState389Action                                                 = &_Action{_ShiftAction, _State389, 0}
	_GotoState390Action                                                 = &_Action{_ShiftAction, _State390, 0}
	_GotoState391Action                                                 = &_Action{_ShiftAction, _State391, 0}
	_GotoState392Action                                                 = &_Action{_ShiftAction, _State392, 0}
	_GotoState393Action                                                 = &_Action{_ShiftAction, _State393, 0}
	_GotoState394Action                                                 = &_Action{_ShiftAction, _State394, 0}
	_GotoState395Action                                                 = &_Action{_ShiftAction, _State395, 0}
	_GotoState396Action                                                 = &_Action{_ShiftAction, _State396, 0}
	_GotoState397Action                                                 = &_Action{_ShiftAction, _State397, 0}
	_GotoState398Action                                                 = &_Action{_ShiftAction, _State398, 0}
	_GotoState399Action                                                 = &_Action{_ShiftAction, _State399, 0}
	_GotoState400Action                                                 = &_Action{_ShiftAction, _State400, 0}
	_GotoState401Action                                                 = &_Action{_ShiftAction, _State401, 0}
	_GotoState402Action                                                 = &_Action{_ShiftAction, _State402, 0}
	_GotoState403Action                                                 = &_Action{_ShiftAction, _State403, 0}
	_GotoState404Action                                                 = &_Action{_ShiftAction, _State404, 0}
	_GotoState405Action                                                 = &_Action{_ShiftAction, _State405, 0}
	_GotoState406Action                                                 = &_Action{_ShiftAction, _State406, 0}
	_GotoState407Action                                                 = &_Action{_ShiftAction, _State407, 0}
	_GotoState408Action                                                 = &_Action{_ShiftAction, _State408, 0}
	_GotoState409Action                                                 = &_Action{_ShiftAction, _State409, 0}
	_GotoState410Action                                                 = &_Action{_ShiftAction, _State410, 0}
	_GotoState411Action                                                 = &_Action{_ShiftAction, _State411, 0}
	_GotoState412Action                                                 = &_Action{_ShiftAction, _State412, 0}
	_GotoState413Action                                                 = &_Action{_ShiftAction, _State413, 0}
	_GotoState414Action                                                 = &_Action{_ShiftAction, _State414, 0}
	_GotoState415Action                                                 = &_Action{_ShiftAction, _State415, 0}
	_GotoState416Action                                                 = &_Action{_ShiftAction, _State416, 0}
	_GotoState417Action                                                 = &_Action{_ShiftAction, _State417, 0}
	_GotoState418Action                                                 = &_Action{_ShiftAction, _State418, 0}
	_GotoState419Action                                                 = &_Action{_ShiftAction, _State419, 0}
	_GotoState420Action                                                 = &_Action{_ShiftAction, _State420, 0}
	_GotoState421Action                                                 = &_Action{_ShiftAction, _State421, 0}
	_GotoState422Action                                                 = &_Action{_ShiftAction, _State422, 0}
	_GotoState423Action                                                 = &_Action{_ShiftAction, _State423, 0}
	_GotoState424Action                                                 = &_Action{_ShiftAction, _State424, 0}
	_GotoState425Action                                                 = &_Action{_ShiftAction, _State425, 0}
	_GotoState426Action                                                 = &_Action{_ShiftAction, _State426, 0}
	_GotoState427Action                                                 = &_Action{_ShiftAction, _State427, 0}
	_GotoState428Action                                                 = &_Action{_ShiftAction, _State428, 0}
	_GotoState429Action                                                 = &_Action{_ShiftAction, _State429, 0}
	_GotoState430Action                                                 = &_Action{_ShiftAction, _State430, 0}
	_GotoState431Action                                                 = &_Action{_ShiftAction, _State431, 0}
	_GotoState432Action                                                 = &_Action{_ShiftAction, _State432, 0}
	_GotoState433Action                                                 = &_Action{_ShiftAction, _State433, 0}
	_GotoState434Action                                                 = &_Action{_ShiftAction, _State434, 0}
	_GotoState435Action                                                 = &_Action{_ShiftAction, _State435, 0}
	_GotoState436Action                                                 = &_Action{_ShiftAction, _State436, 0}
	_GotoState437Action                                                 = &_Action{_ShiftAction, _State437, 0}
	_GotoState438Action                                                 = &_Action{_ShiftAction, _State438, 0}
	_GotoState439Action                                                 = &_Action{_ShiftAction, _State439, 0}
	_GotoState440Action                                                 = &_Action{_ShiftAction, _State440, 0}
	_GotoState441Action                                                 = &_Action{_ShiftAction, _State441, 0}
	_GotoState442Action                                                 = &_Action{_ShiftAction, _State442, 0}
	_GotoState443Action                                                 = &_Action{_ShiftAction, _State443, 0}
	_GotoState444Action                                                 = &_Action{_ShiftAction, _State444, 0}
	_GotoState445Action                                                 = &_Action{_ShiftAction, _State445, 0}
	_GotoState446Action                                                 = &_Action{_ShiftAction, _State446, 0}
	_GotoState447Action                                                 = &_Action{_ShiftAction, _State447, 0}
	_GotoState448Action                                                 = &_Action{_ShiftAction, _State448, 0}
	_GotoState449Action                                                 = &_Action{_ShiftAction, _State449, 0}
	_GotoState450Action                                                 = &_Action{_ShiftAction, _State450, 0}
	_GotoState451Action                                                 = &_Action{_ShiftAction, _State451, 0}
	_GotoState452Action                                                 = &_Action{_ShiftAction, _State452, 0}
	_GotoState453Action                                                 = &_Action{_ShiftAction, _State453, 0}
	_GotoState454Action                                                 = &_Action{_ShiftAction, _State454, 0}
	_GotoState455Action                                                 = &_Action{_ShiftAction, _State455, 0}
	_GotoState456Action                                                 = &_Action{_ShiftAction, _State456, 0}
	_GotoState457Action                                                 = &_Action{_ShiftAction, _State457, 0}
	_GotoState458Action                                                 = &_Action{_ShiftAction, _State458, 0}
	_GotoState459Action                                                 = &_Action{_ShiftAction, _State459, 0}
	_GotoState460Action                                                 = &_Action{_ShiftAction, _State460, 0}
	_GotoState461Action                                                 = &_Action{_ShiftAction, _State461, 0}
	_GotoState462Action                                                 = &_Action{_ShiftAction, _State462, 0}
	_GotoState463Action                                                 = &_Action{_ShiftAction, _State463, 0}
	_GotoState464Action                                                 = &_Action{_ShiftAction, _State464, 0}
	_GotoState465Action                                                 = &_Action{_ShiftAction, _State465, 0}
	_GotoState466Action                                                 = &_Action{_ShiftAction, _State466, 0}
	_GotoState467Action                                                 = &_Action{_ShiftAction, _State467, 0}
	_GotoState468Action                                                 = &_Action{_ShiftAction, _State468, 0}
	_GotoState469Action                                                 = &_Action{_ShiftAction, _State469, 0}
	_GotoState470Action                                                 = &_Action{_ShiftAction, _State470, 0}
	_GotoState471Action                                                 = &_Action{_ShiftAction, _State471, 0}
	_GotoState472Action                                                 = &_Action{_ShiftAction, _State472, 0}
	_GotoState473Action                                                 = &_Action{_ShiftAction, _State473, 0}
	_GotoState474Action                                                 = &_Action{_ShiftAction, _State474, 0}
	_GotoState475Action                                                 = &_Action{_ShiftAction, _State475, 0}
	_GotoState476Action                                                 = &_Action{_ShiftAction, _State476, 0}
	_GotoState477Action                                                 = &_Action{_ShiftAction, _State477, 0}
	_GotoState478Action                                                 = &_Action{_ShiftAction, _State478, 0}
	_ReduceToSourceAction                                               = &_Action{_ReduceAction, 0, _ReduceToSource}
	_ReduceAddToProperDefinitionsAction                                 = &_Action{_ReduceAction, 0, _ReduceAddToProperDefinitions}
	_ReduceDefinitionToProperDefinitionsAction                          = &_Action{_ReduceAction, 0, _ReduceDefinitionToProperDefinitions}
	_ReduceProperDefinitionsToDefinitionsAction                         = &_Action{_ReduceAction, 0, _ReduceProperDefinitionsToDefinitions}
	_ReduceImproperToDefinitionsAction                                  = &_Action{_ReduceAction, 0, _ReduceImproperToDefinitions}
	_ReduceNilToDefinitionsAction                                       = &_Action{_ReduceAction, 0, _ReduceNilToDefinitions}
	_ReducePackageDefToDefinitionAction                                 = &_Action{_ReduceAction, 0, _ReducePackageDefToDefinition}
	_ReduceTypeDefToDefinitionAction                                    = &_Action{_ReduceAction, 0, _ReduceTypeDefToDefinition}
	_ReduceNamedFuncDefToDefinitionAction                               = &_Action{_ReduceAction, 0, _ReduceNamedFuncDefToDefinition}
	_ReduceGlobalVarDefToDefinitionAction                               = &_Action{_ReduceAction, 0, _ReduceGlobalVarDefToDefinition}
	_ReduceGlobalVarAssignmentToDefinitionAction                        = &_Action{_ReduceAction, 0, _ReduceGlobalVarAssignmentToDefinition}
	_ReduceStatementBlockToDefinitionAction                             = &_Action{_ReduceAction, 0, _ReduceStatementBlockToDefinition}
	_ReduceCommentGroupsToDefinitionAction                              = &_Action{_ReduceAction, 0, _ReduceCommentGroupsToDefinition}
	_ReduceToStatementBlockAction                                       = &_Action{_ReduceAction, 0, _ReduceToStatementBlock}
	_ReduceAddImplicitToProperStatementsAction                          = &_Action{_ReduceAction, 0, _ReduceAddImplicitToProperStatements}
	_ReduceAddExplicitToProperStatementsAction                          = &_Action{_ReduceAction, 0, _ReduceAddExplicitToProperStatements}
	_ReduceStatementToProperStatementsAction                            = &_Action{_ReduceAction, 0, _ReduceStatementToProperStatements}
	_ReduceProperStatementsToStatementsAction                           = &_Action{_ReduceAction, 0, _ReduceProperStatementsToStatements}
	_ReduceImproperImplicitToStatementsAction                           = &_Action{_ReduceAction, 0, _ReduceImproperImplicitToStatements}
	_ReduceImproperExplicitToStatementsAction                           = &_Action{_ReduceAction, 0, _ReduceImproperExplicitToStatements}
	_ReduceNilToStatementsAction                                        = &_Action{_ReduceAction, 0, _ReduceNilToStatements}
	_ReduceUnsafeStatementToSimpleStatementAction                       = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToSimpleStatement}
	_ReduceExprOrImproperStructStatementToSimpleStatementAction         = &_Action{_ReduceAction, 0, _ReduceExprOrImproperStructStatementToSimpleStatement}
	_ReduceCallbackOpStatementToSimpleStatementAction                   = &_Action{_ReduceAction, 0, _ReduceCallbackOpStatementToSimpleStatement}
	_ReduceJumpStatementToSimpleStatementAction                         = &_Action{_ReduceAction, 0, _ReduceJumpStatementToSimpleStatement}
	_ReduceAssignStatementToSimpleStatementAction                       = &_Action{_ReduceAction, 0, _ReduceAssignStatementToSimpleStatement}
	_ReduceUnaryOpAssignStatementToSimpleStatementAction                = &_Action{_ReduceAction, 0, _ReduceUnaryOpAssignStatementToSimpleStatement}
	_ReduceBinaryOpAssignStatementToSimpleStatementAction               = &_Action{_ReduceAction, 0, _ReduceBinaryOpAssignStatementToSimpleStatement}
	_ReduceFallthroughStatementToSimpleStatementAction                  = &_Action{_ReduceAction, 0, _ReduceFallthroughStatementToSimpleStatement}
	_ReduceSimpleStatementToStatementAction                             = &_Action{_ReduceAction, 0, _ReduceSimpleStatementToStatement}
	_ReduceImportStatementToStatementAction                             = &_Action{_ReduceAction, 0, _ReduceImportStatementToStatement}
	_ReduceCaseBranchStatementToStatementAction                         = &_Action{_ReduceAction, 0, _ReduceCaseBranchStatementToStatement}
	_ReduceDefaultBranchStatementToStatementAction                      = &_Action{_ReduceAction, 0, _ReduceDefaultBranchStatementToStatement}
	_ReduceSimpleStatementToOptionalSimpleStatementAction               = &_Action{_ReduceAction, 0, _ReduceSimpleStatementToOptionalSimpleStatement}
	_ReduceNilToOptionalSimpleStatementAction                           = &_Action{_ReduceAction, 0, _ReduceNilToOptionalSimpleStatement}
	_ReduceToExprOrImproperStructStatementAction                        = &_Action{_ReduceAction, 0, _ReduceToExprOrImproperStructStatement}
	_ReduceExprToExprsAction                                            = &_Action{_ReduceAction, 0, _ReduceExprToExprs}
	_ReduceAddToExprsAction                                             = &_Action{_ReduceAction, 0, _ReduceAddToExprs}
	_ReduceAsyncToCallbackOpAction                                      = &_Action{_ReduceAction, 0, _ReduceAsyncToCallbackOp}
	_ReduceDeferToCallbackOpAction                                      = &_Action{_ReduceAction, 0, _ReduceDeferToCallbackOp}
	_ReduceToCallbackOpStatementAction                                  = &_Action{_ReduceAction, 0, _ReduceToCallbackOpStatement}
	_ReduceToUnsafeStatementAction                                      = &_Action{_ReduceAction, 0, _ReduceToUnsafeStatement}
	_ReduceUnlabeledNoValueToJumpStatementAction                        = &_Action{_ReduceAction, 0, _ReduceUnlabeledNoValueToJumpStatement}
	_ReduceUnlabeledValuedToJumpStatementAction                         = &_Action{_ReduceAction, 0, _ReduceUnlabeledValuedToJumpStatement}
	_ReduceLabeledNoValueToJumpStatementAction                          = &_Action{_ReduceAction, 0, _ReduceLabeledNoValueToJumpStatement}
	_ReduceLabeledValuedToJumpStatementAction                           = &_Action{_ReduceAction, 0, _ReduceLabeledValuedToJumpStatement}
	_ReduceReturnToJumpTypeAction                                       = &_Action{_ReduceAction, 0, _ReduceReturnToJumpType}
	_ReduceBreakToJumpTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceBreakToJumpType}
	_ReduceContinueToJumpTypeAction                                     = &_Action{_ReduceAction, 0, _ReduceContinueToJumpType}
	_ReduceToFallthroughStatementAction                                 = &_Action{_ReduceAction, 0, _ReduceToFallthroughStatement}
	_ReduceToAssignStatementAction                                      = &_Action{_ReduceAction, 0, _ReduceToAssignStatement}
	_ReduceToUnaryOpAssignStatementAction                               = &_Action{_ReduceAction, 0, _ReduceToUnaryOpAssignStatement}
	_ReduceAddOneAssignToUnaryOpAssignAction                            = &_Action{_ReduceAction, 0, _ReduceAddOneAssignToUnaryOpAssign}
	_ReduceSubOneAssignToUnaryOpAssignAction                            = &_Action{_ReduceAction, 0, _ReduceSubOneAssignToUnaryOpAssign}
	_ReduceToBinaryOpAssignStatementAction                              = &_Action{_ReduceAction, 0, _ReduceToBinaryOpAssignStatement}
	_ReduceAddAssignToBinaryOpAssignAction                              = &_Action{_ReduceAction, 0, _ReduceAddAssignToBinaryOpAssign}
	_ReduceSubAssignToBinaryOpAssignAction                              = &_Action{_ReduceAction, 0, _ReduceSubAssignToBinaryOpAssign}
	_ReduceMulAssignToBinaryOpAssignAction                              = &_Action{_ReduceAction, 0, _ReduceMulAssignToBinaryOpAssign}
	_ReduceDivAssignToBinaryOpAssignAction                              = &_Action{_ReduceAction, 0, _ReduceDivAssignToBinaryOpAssign}
	_ReduceModAssignToBinaryOpAssignAction                              = &_Action{_ReduceAction, 0, _ReduceModAssignToBinaryOpAssign}
	_ReduceBitNegAssignToBinaryOpAssignAction                           = &_Action{_ReduceAction, 0, _ReduceBitNegAssignToBinaryOpAssign}
	_ReduceBitAndAssignToBinaryOpAssignAction                           = &_Action{_ReduceAction, 0, _ReduceBitAndAssignToBinaryOpAssign}
	_ReduceBitOrAssignToBinaryOpAssignAction                            = &_Action{_ReduceAction, 0, _ReduceBitOrAssignToBinaryOpAssign}
	_ReduceBitXorAssignToBinaryOpAssignAction                           = &_Action{_ReduceAction, 0, _ReduceBitXorAssignToBinaryOpAssign}
	_ReduceBitLshiftAssignToBinaryOpAssignAction                        = &_Action{_ReduceAction, 0, _ReduceBitLshiftAssignToBinaryOpAssign}
	_ReduceBitRshiftAssignToBinaryOpAssignAction                        = &_Action{_ReduceAction, 0, _ReduceBitRshiftAssignToBinaryOpAssign}
	_ReduceSingleToImportStatementAction                                = &_Action{_ReduceAction, 0, _ReduceSingleToImportStatement}
	_ReduceMultipleToImportStatementAction                              = &_Action{_ReduceAction, 0, _ReduceMultipleToImportStatement}
	_ReduceStringLiteralToImportClauseAction                            = &_Action{_ReduceAction, 0, _ReduceStringLiteralToImportClause}
	_ReduceAliasToImportClauseAction                                    = &_Action{_ReduceAction, 0, _ReduceAliasToImportClause}
	_ReduceUnusableImportToImportClauseAction                           = &_Action{_ReduceAction, 0, _ReduceUnusableImportToImportClause}
	_ReduceImportToLocalToImportClauseAction                            = &_Action{_ReduceAction, 0, _ReduceImportToLocalToImportClause}
	_ReduceAddImplicitToProperImportClausesAction                       = &_Action{_ReduceAction, 0, _ReduceAddImplicitToProperImportClauses}
	_ReduceAddExplicitToProperImportClausesAction                       = &_Action{_ReduceAction, 0, _ReduceAddExplicitToProperImportClauses}
	_ReduceImportClauseToProperImportClausesAction                      = &_Action{_ReduceAction, 0, _ReduceImportClauseToProperImportClauses}
	_ReduceProperImportClausesToImportClausesAction                     = &_Action{_ReduceAction, 0, _ReduceProperImportClausesToImportClauses}
	_ReduceImplicitToImportClausesAction                                = &_Action{_ReduceAction, 0, _ReduceImplicitToImportClauses}
	_ReduceExplicitToImportClausesAction                                = &_Action{_ReduceAction, 0, _ReduceExplicitToImportClauses}
	_ReduceCasePatternToCasePatternsAction                              = &_Action{_ReduceAction, 0, _ReduceCasePatternToCasePatterns}
	_ReduceMultipleToCasePatternsAction                                 = &_Action{_ReduceAction, 0, _ReduceMultipleToCasePatterns}
	_ReduceToVarDeclPatternAction                                       = &_Action{_ReduceAction, 0, _ReduceToVarDeclPattern}
	_ReduceVarToVarOrLetAction                                          = &_Action{_ReduceAction, 0, _ReduceVarToVarOrLet}
	_ReduceLetToVarOrLetAction                                          = &_Action{_ReduceAction, 0, _ReduceLetToVarOrLet}
	_ReduceIdentifierToVarPatternAction                                 = &_Action{_ReduceAction, 0, _ReduceIdentifierToVarPattern}
	_ReduceUnderscoreToVarPatternAction                                 = &_Action{_ReduceAction, 0, _ReduceUnderscoreToVarPattern}
	_ReduceTuplePatternToVarPatternAction                               = &_Action{_ReduceAction, 0, _ReduceTuplePatternToVarPattern}
	_ReduceToTuplePatternAction                                         = &_Action{_ReduceAction, 0, _ReduceToTuplePattern}
	_ReduceFieldVarPatternToFieldVarPatternsAction                      = &_Action{_ReduceAction, 0, _ReduceFieldVarPatternToFieldVarPatterns}
	_ReduceAddToFieldVarPatternsAction                                  = &_Action{_ReduceAction, 0, _ReduceAddToFieldVarPatterns}
	_ReducePositionalToFieldVarPatternAction                            = &_Action{_ReduceAction, 0, _ReducePositionalToFieldVarPattern}
	_ReduceNamedToFieldVarPatternAction                                 = &_Action{_ReduceAction, 0, _ReduceNamedToFieldVarPattern}
	_ReduceEllipsisToFieldVarPatternAction                              = &_Action{_ReduceAction, 0, _ReduceEllipsisToFieldVarPattern}
	_ReduceTypeExprToOptionalTypeExprAction                             = &_Action{_ReduceAction, 0, _ReduceTypeExprToOptionalTypeExpr}
	_ReduceNilToOptionalTypeExprAction                                  = &_Action{_ReduceAction, 0, _ReduceNilToOptionalTypeExpr}
	_ReduceToAssignPatternAction                                        = &_Action{_ReduceAction, 0, _ReduceToAssignPattern}
	_ReduceAssignPatternToCasePatternAction                             = &_Action{_ReduceAction, 0, _ReduceAssignPatternToCasePattern}
	_ReduceEnumMatchPatternToCasePatternAction                          = &_Action{_ReduceAction, 0, _ReduceEnumMatchPatternToCasePattern}
	_ReduceEnumNondataMatchPattenToCasePatternAction                    = &_Action{_ReduceAction, 0, _ReduceEnumNondataMatchPattenToCasePattern}
	_ReduceEnumVarDeclPatternToCasePatternAction                        = &_Action{_ReduceAction, 0, _ReduceEnumVarDeclPatternToCasePattern}
	_ReduceIfExprToExprAction                                           = &_Action{_ReduceAction, 0, _ReduceIfExprToExpr}
	_ReduceSwitchExprToExprAction                                       = &_Action{_ReduceAction, 0, _ReduceSwitchExprToExpr}
	_ReduceLoopExprToExprAction                                         = &_Action{_ReduceAction, 0, _ReduceLoopExprToExpr}
	_ReduceSequenceExprToExprAction                                     = &_Action{_ReduceAction, 0, _ReduceSequenceExprToExpr}
	_ReduceLabelDeclToOptionalLabelDeclAction                           = &_Action{_ReduceAction, 0, _ReduceLabelDeclToOptionalLabelDecl}
	_ReduceUnlabelledToOptionalLabelDeclAction                          = &_Action{_ReduceAction, 0, _ReduceUnlabelledToOptionalLabelDecl}
	_ReduceOrExprToSequenceExprAction                                   = &_Action{_ReduceAction, 0, _ReduceOrExprToSequenceExpr}
	_ReduceVarDeclPatternToSequenceExprAction                           = &_Action{_ReduceAction, 0, _ReduceVarDeclPatternToSequenceExpr}
	_ReduceAssignVarPatternToSequenceExprAction                         = &_Action{_ReduceAction, 0, _ReduceAssignVarPatternToSequenceExpr}
	_ReduceNoElseToIfExprAction                                         = &_Action{_ReduceAction, 0, _ReduceNoElseToIfExpr}
	_ReduceIfElseToIfExprAction                                         = &_Action{_ReduceAction, 0, _ReduceIfElseToIfExpr}
	_ReduceMultiIfElseToIfExprAction                                    = &_Action{_ReduceAction, 0, _ReduceMultiIfElseToIfExpr}
	_ReduceSequenceExprToConditionAction                                = &_Action{_ReduceAction, 0, _ReduceSequenceExprToCondition}
	_ReduceCaseToConditionAction                                        = &_Action{_ReduceAction, 0, _ReduceCaseToCondition}
	_ReduceToSwitchExprAction                                           = &_Action{_ReduceAction, 0, _ReduceToSwitchExpr}
	_ReduceInfiniteToLoopExprAction                                     = &_Action{_ReduceAction, 0, _ReduceInfiniteToLoopExpr}
	_ReduceDoWhileToLoopExprAction                                      = &_Action{_ReduceAction, 0, _ReduceDoWhileToLoopExpr}
	_ReduceWhileToLoopExprAction                                        = &_Action{_ReduceAction, 0, _ReduceWhileToLoopExpr}
	_ReduceIteratorToLoopExprAction                                     = &_Action{_ReduceAction, 0, _ReduceIteratorToLoopExpr}
	_ReduceForToLoopExprAction                                          = &_Action{_ReduceAction, 0, _ReduceForToLoopExpr}
	_ReduceSequenceExprToOptionalSequenceExprAction                     = &_Action{_ReduceAction, 0, _ReduceSequenceExprToOptionalSequenceExpr}
	_ReduceNilToOptionalSequenceExprAction                              = &_Action{_ReduceAction, 0, _ReduceNilToOptionalSequenceExpr}
	_ReduceSequenceExprToForAssignmentAction                            = &_Action{_ReduceAction, 0, _ReduceSequenceExprToForAssignment}
	_ReduceAssignToForAssignmentAction                                  = &_Action{_ReduceAction, 0, _ReduceAssignToForAssignment}
	_ReduceToCallExprAction                                             = &_Action{_ReduceAction, 0, _ReduceToCallExpr}
	_ReduceAddToProperArgumentsAction                                   = &_Action{_ReduceAction, 0, _ReduceAddToProperArguments}
	_ReduceArgumentToProperArgumentsAction                              = &_Action{_ReduceAction, 0, _ReduceArgumentToProperArguments}
	_ReduceProperArgumentsToArgumentsAction                             = &_Action{_ReduceAction, 0, _ReduceProperArgumentsToArguments}
	_ReduceImproperToArgumentsAction                                    = &_Action{_ReduceAction, 0, _ReduceImproperToArguments}
	_ReduceNilToArgumentsAction                                         = &_Action{_ReduceAction, 0, _ReduceNilToArguments}
	_ReducePositionalToArgumentAction                                   = &_Action{_ReduceAction, 0, _ReducePositionalToArgument}
	_ReduceColonExprToArgumentAction                                    = &_Action{_ReduceAction, 0, _ReduceColonExprToArgument}
	_ReduceNamedAssignmentToArgumentAction                              = &_Action{_ReduceAction, 0, _ReduceNamedAssignmentToArgument}
	_ReduceVarargAssignmentToArgumentAction                             = &_Action{_ReduceAction, 0, _ReduceVarargAssignmentToArgument}
	_ReduceSkipPatternToArgumentAction                                  = &_Action{_ReduceAction, 0, _ReduceSkipPatternToArgument}
	_ReduceUnitUnitPairToColonExprAction                                = &_Action{_ReduceAction, 0, _ReduceUnitUnitPairToColonExpr}
	_ReduceExprUnitPairToColonExprAction                                = &_Action{_ReduceAction, 0, _ReduceExprUnitPairToColonExpr}
	_ReduceUnitExprPairToColonExprAction                                = &_Action{_ReduceAction, 0, _ReduceUnitExprPairToColonExpr}
	_ReduceExprExprPairToColonExprAction                                = &_Action{_ReduceAction, 0, _ReduceExprExprPairToColonExpr}
	_ReduceColonExprUnitTupleToColonExprAction                          = &_Action{_ReduceAction, 0, _ReduceColonExprUnitTupleToColonExpr}
	_ReduceColonExprExprTupleToColonExprAction                          = &_Action{_ReduceAction, 0, _ReduceColonExprExprTupleToColonExpr}
	_ReduceParseErrorExprToAtomExprAction                               = &_Action{_ReduceAction, 0, _ReduceParseErrorExprToAtomExpr}
	_ReduceLiteralExprToAtomExprAction                                  = &_Action{_ReduceAction, 0, _ReduceLiteralExprToAtomExpr}
	_ReduceNamedExprToAtomExprAction                                    = &_Action{_ReduceAction, 0, _ReduceNamedExprToAtomExpr}
	_ReduceBlockExprToAtomExprAction                                    = &_Action{_ReduceAction, 0, _ReduceBlockExprToAtomExpr}
	_ReduceAnonymousFuncExprToAtomExprAction                            = &_Action{_ReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}
	_ReduceInitializeExprToAtomExprAction                               = &_Action{_ReduceAction, 0, _ReduceInitializeExprToAtomExpr}
	_ReduceImplicitStructExprToAtomExprAction                           = &_Action{_ReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}
	_ReduceToParseErrorExprAction                                       = &_Action{_ReduceAction, 0, _ReduceToParseErrorExpr}
	_ReduceTrueToLiteralExprAction                                      = &_Action{_ReduceAction, 0, _ReduceTrueToLiteralExpr}
	_ReduceFalseToLiteralExprAction                                     = &_Action{_ReduceAction, 0, _ReduceFalseToLiteralExpr}
	_ReduceIntegerLiteralToLiteralExprAction                            = &_Action{_ReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}
	_ReduceFloatLiteralToLiteralExprAction                              = &_Action{_ReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}
	_ReduceRuneLiteralToLiteralExprAction                               = &_Action{_ReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}
	_ReduceStringLiteralToLiteralExprAction                             = &_Action{_ReduceAction, 0, _ReduceStringLiteralToLiteralExpr}
	_ReduceIdentifierToNamedExprAction                                  = &_Action{_ReduceAction, 0, _ReduceIdentifierToNamedExpr}
	_ReduceUnderscoreToNamedExprAction                                  = &_Action{_ReduceAction, 0, _ReduceUnderscoreToNamedExpr}
	_ReduceToBlockExprAction                                            = &_Action{_ReduceAction, 0, _ReduceToBlockExpr}
	_ReduceToInitializeExprAction                                       = &_Action{_ReduceAction, 0, _ReduceToInitializeExpr}
	_ReduceToImplicitStructExprAction                                   = &_Action{_ReduceAction, 0, _ReduceToImplicitStructExpr}
	_ReduceAtomExprToAccessibleExprAction                               = &_Action{_ReduceAction, 0, _ReduceAtomExprToAccessibleExpr}
	_ReduceAccessExprToAccessibleExprAction                             = &_Action{_ReduceAction, 0, _ReduceAccessExprToAccessibleExpr}
	_ReduceCallExprToAccessibleExprAction                               = &_Action{_ReduceAction, 0, _ReduceCallExprToAccessibleExpr}
	_ReduceIndexExprToAccessibleExprAction                              = &_Action{_ReduceAction, 0, _ReduceIndexExprToAccessibleExpr}
	_ReduceToAccessExprAction                                           = &_Action{_ReduceAction, 0, _ReduceToAccessExpr}
	_ReduceToIndexExprAction                                            = &_Action{_ReduceAction, 0, _ReduceToIndexExpr}
	_ReduceAccessibleExprToPostfixableExprAction                        = &_Action{_ReduceAction, 0, _ReduceAccessibleExprToPostfixableExpr}
	_ReducePostfixUnaryExprToPostfixableExprAction                      = &_Action{_ReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}
	_ReduceToPostfixUnaryExprAction                                     = &_Action{_ReduceAction, 0, _ReduceToPostfixUnaryExpr}
	_ReducePostfixableExprToPrefixableExprAction                        = &_Action{_ReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}
	_ReducePrefixUnaryExprToPrefixableExprAction                        = &_Action{_ReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}
	_ReduceToPrefixUnaryExprAction                                      = &_Action{_ReduceAction, 0, _ReduceToPrefixUnaryExpr}
	_ReduceNotToPrefixUnaryOpAction                                     = &_Action{_ReduceAction, 0, _ReduceNotToPrefixUnaryOp}
	_ReduceBitNegToPrefixUnaryOpAction                                  = &_Action{_ReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}
	_ReduceSubToPrefixUnaryOpAction                                     = &_Action{_ReduceAction, 0, _ReduceSubToPrefixUnaryOp}
	_ReduceMulToPrefixUnaryOpAction                                     = &_Action{_ReduceAction, 0, _ReduceMulToPrefixUnaryOp}
	_ReduceBitAndToPrefixUnaryOpAction                                  = &_Action{_ReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}
	_ReducePrefixableExprToMulExprAction                                = &_Action{_ReduceAction, 0, _ReducePrefixableExprToMulExpr}
	_ReduceBinaryMulExprToMulExprAction                                 = &_Action{_ReduceAction, 0, _ReduceBinaryMulExprToMulExpr}
	_ReduceToBinaryMulExprAction                                        = &_Action{_ReduceAction, 0, _ReduceToBinaryMulExpr}
	_ReduceMulToMulOpAction                                             = &_Action{_ReduceAction, 0, _ReduceMulToMulOp}
	_ReduceDivToMulOpAction                                             = &_Action{_ReduceAction, 0, _ReduceDivToMulOp}
	_ReduceModToMulOpAction                                             = &_Action{_ReduceAction, 0, _ReduceModToMulOp}
	_ReduceBitAndToMulOpAction                                          = &_Action{_ReduceAction, 0, _ReduceBitAndToMulOp}
	_ReduceBitLshiftToMulOpAction                                       = &_Action{_ReduceAction, 0, _ReduceBitLshiftToMulOp}
	_ReduceBitRshiftToMulOpAction                                       = &_Action{_ReduceAction, 0, _ReduceBitRshiftToMulOp}
	_ReduceMulExprToAddExprAction                                       = &_Action{_ReduceAction, 0, _ReduceMulExprToAddExpr}
	_ReduceBinaryAddExprToAddExprAction                                 = &_Action{_ReduceAction, 0, _ReduceBinaryAddExprToAddExpr}
	_ReduceToBinaryAddExprAction                                        = &_Action{_ReduceAction, 0, _ReduceToBinaryAddExpr}
	_ReduceAddToAddOpAction                                             = &_Action{_ReduceAction, 0, _ReduceAddToAddOp}
	_ReduceSubToAddOpAction                                             = &_Action{_ReduceAction, 0, _ReduceSubToAddOp}
	_ReduceBitOrToAddOpAction                                           = &_Action{_ReduceAction, 0, _ReduceBitOrToAddOp}
	_ReduceBitXorToAddOpAction                                          = &_Action{_ReduceAction, 0, _ReduceBitXorToAddOp}
	_ReduceAddExprToCmpExprAction                                       = &_Action{_ReduceAction, 0, _ReduceAddExprToCmpExpr}
	_ReduceBinaryCmpExprToCmpExprAction                                 = &_Action{_ReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}
	_ReduceToBinaryCmpExprAction                                        = &_Action{_ReduceAction, 0, _ReduceToBinaryCmpExpr}
	_ReduceEqualToCmpOpAction                                           = &_Action{_ReduceAction, 0, _ReduceEqualToCmpOp}
	_ReduceNotEqualToCmpOpAction                                        = &_Action{_ReduceAction, 0, _ReduceNotEqualToCmpOp}
	_ReduceLessToCmpOpAction                                            = &_Action{_ReduceAction, 0, _ReduceLessToCmpOp}
	_ReduceLessOrEqualToCmpOpAction                                     = &_Action{_ReduceAction, 0, _ReduceLessOrEqualToCmpOp}
	_ReduceGreaterToCmpOpAction                                         = &_Action{_ReduceAction, 0, _ReduceGreaterToCmpOp}
	_ReduceGreaterOrEqualToCmpOpAction                                  = &_Action{_ReduceAction, 0, _ReduceGreaterOrEqualToCmpOp}
	_ReduceCmpExprToAndExprAction                                       = &_Action{_ReduceAction, 0, _ReduceCmpExprToAndExpr}
	_ReduceBinaryAndExprToAndExprAction                                 = &_Action{_ReduceAction, 0, _ReduceBinaryAndExprToAndExpr}
	_ReduceToBinaryAndExprAction                                        = &_Action{_ReduceAction, 0, _ReduceToBinaryAndExpr}
	_ReduceAndExprToOrExprAction                                        = &_Action{_ReduceAction, 0, _ReduceAndExprToOrExpr}
	_ReduceBinaryOrExprToOrExprAction                                   = &_Action{_ReduceAction, 0, _ReduceBinaryOrExprToOrExpr}
	_ReduceToBinaryOrExprAction                                         = &_Action{_ReduceAction, 0, _ReduceToBinaryOrExpr}
	_ReduceExplicitStructTypeExprToInitializableTypeExprAction          = &_Action{_ReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}
	_ReduceSliceTypeExprToInitializableTypeExprAction                   = &_Action{_ReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}
	_ReduceArrayTypeExprToInitializableTypeExprAction                   = &_Action{_ReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}
	_ReduceMapTypeExprToInitializableTypeExprAction                     = &_Action{_ReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}
	_ReduceToSliceTypeExprAction                                        = &_Action{_ReduceAction, 0, _ReduceToSliceTypeExpr}
	_ReduceToArrayTypeExprAction                                        = &_Action{_ReduceAction, 0, _ReduceToArrayTypeExpr}
	_ReduceToMapTypeExprAction                                          = &_Action{_ReduceAction, 0, _ReduceToMapTypeExpr}
	_ReduceInitializableTypeExprToAtomTypeExprAction                    = &_Action{_ReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}
	_ReduceNamedTypeExprToAtomTypeExprAction                            = &_Action{_ReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}
	_ReduceInferredTypeExprToAtomTypeExprAction                         = &_Action{_ReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}
	_ReduceImplicitStructTypeExprToAtomTypeExprAction                   = &_Action{_ReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}
	_ReduceExplicitEnumTypeExprToAtomTypeExprAction                     = &_Action{_ReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}
	_ReduceImplicitEnumTypeExprToAtomTypeExprAction                     = &_Action{_ReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}
	_ReduceTraitTypeExprToAtomTypeExprAction                            = &_Action{_ReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}
	_ReduceFuncTypeExprToAtomTypeExprAction                             = &_Action{_ReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}
	_ReduceParseErrorTypeExprToAtomTypeExprAction                       = &_Action{_ReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}
	_ReduceLocalToNamedTypeExprAction                                   = &_Action{_ReduceAction, 0, _ReduceLocalToNamedTypeExpr}
	_ReduceExternalToNamedTypeExprAction                                = &_Action{_ReduceAction, 0, _ReduceExternalToNamedTypeExpr}
	_ReduceDotToInferredTypeExprAction                                  = &_Action{_ReduceAction, 0, _ReduceDotToInferredTypeExpr}
	_ReduceUnderscoreToInferredTypeExprAction                           = &_Action{_ReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}
	_ReduceToParseErrorTypeExprAction                                   = &_Action{_ReduceAction, 0, _ReduceToParseErrorTypeExpr}
	_ReduceAtomTypeExprToReturnableTypeExprAction                       = &_Action{_ReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}
	_ReducePrefixUnaryTypeExprToReturnableTypeExprAction                = &_Action{_ReduceAction, 0, _ReducePrefixUnaryTypeExprToReturnableTypeExpr}
	_ReduceToPrefixUnaryTypeExprAction                                  = &_Action{_ReduceAction, 0, _ReduceToPrefixUnaryTypeExpr}
	_ReduceQuestionToPrefixUnaryTypeOpAction                            = &_Action{_ReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}
	_ReduceExclaimToPrefixUnaryTypeOpAction                             = &_Action{_ReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}
	_ReduceBitAndToPrefixUnaryTypeOpAction                              = &_Action{_ReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}
	_ReduceBitNegToPrefixUnaryTypeOpAction                              = &_Action{_ReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}
	_ReduceTildeTildeToPrefixUnaryTypeOpAction                          = &_Action{_ReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}
	_ReduceReturnableTypeExprToTypeExprAction                           = &_Action{_ReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}
	_ReduceBinaryTypeExprToTypeExprAction                               = &_Action{_ReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}
	_ReduceToBinaryTypeExprAction                                       = &_Action{_ReduceAction, 0, _ReduceToBinaryTypeExpr}
	_ReduceMulToBinaryTypeOpAction                                      = &_Action{_ReduceAction, 0, _ReduceMulToBinaryTypeOp}
	_ReduceAddToBinaryTypeOpAction                                      = &_Action{_ReduceAction, 0, _ReduceAddToBinaryTypeOp}
	_ReduceSubToBinaryTypeOpAction                                      = &_Action{_ReduceAction, 0, _ReduceSubToBinaryTypeOp}
	_ReduceDefinitionToTypeDefAction                                    = &_Action{_ReduceAction, 0, _ReduceDefinitionToTypeDef}
	_ReduceConstrainedDefToTypeDefAction                                = &_Action{_ReduceAction, 0, _ReduceConstrainedDefToTypeDef}
	_ReduceAliasToTypeDefAction                                         = &_Action{_ReduceAction, 0, _ReduceAliasToTypeDef}
	_ReduceUnconstrainedToGenericParameterDefAction                     = &_Action{_ReduceAction, 0, _ReduceUnconstrainedToGenericParameterDef}
	_ReduceConstrainedToGenericParameterDefAction                       = &_Action{_ReduceAction, 0, _ReduceConstrainedToGenericParameterDef}
	_ReduceGenericToGenericParameterDefsAction                          = &_Action{_ReduceAction, 0, _ReduceGenericToGenericParameterDefs}
	_ReduceNilToGenericParameterDefsAction                              = &_Action{_ReduceAction, 0, _ReduceNilToGenericParameterDefs}
	_ReduceAddToProperGenericParameterDefListAction                     = &_Action{_ReduceAction, 0, _ReduceAddToProperGenericParameterDefList}
	_ReduceGenericParameterDefToProperGenericParameterDefListAction     = &_Action{_ReduceAction, 0, _ReduceGenericParameterDefToProperGenericParameterDefList}
	_ReduceProperGenericParameterDefListToGenericParameterDefListAction = &_Action{_ReduceAction, 0, _ReduceProperGenericParameterDefListToGenericParameterDefList}
	_ReduceImproperToGenericParameterDefListAction                      = &_Action{_ReduceAction, 0, _ReduceImproperToGenericParameterDefList}
	_ReduceNilToGenericParameterDefListAction                           = &_Action{_ReduceAction, 0, _ReduceNilToGenericParameterDefList}
	_ReduceBindingToGenericTypeArgumentsAction                          = &_Action{_ReduceAction, 0, _ReduceBindingToGenericTypeArguments}
	_ReduceNilToGenericTypeArgumentsAction                              = &_Action{_ReduceAction, 0, _ReduceNilToGenericTypeArguments}
	_ReduceAddToProperTypeArgumentsAction                               = &_Action{_ReduceAction, 0, _ReduceAddToProperTypeArguments}
	_ReduceTypeExprToProperTypeArgumentsAction                          = &_Action{_ReduceAction, 0, _ReduceTypeExprToProperTypeArguments}
	_ReduceProperTypeArgumentsToTypeArgumentsAction                     = &_Action{_ReduceAction, 0, _ReduceProperTypeArgumentsToTypeArguments}
	_ReduceImproperToTypeArgumentsAction                                = &_Action{_ReduceAction, 0, _ReduceImproperToTypeArguments}
	_ReduceNilToTypeArgumentsAction                                     = &_Action{_ReduceAction, 0, _ReduceNilToTypeArguments}
	_ReduceNamedToFieldDefAction                                        = &_Action{_ReduceAction, 0, _ReduceNamedToFieldDef}
	_ReduceUnnamedToFieldDefAction                                      = &_Action{_ReduceAction, 0, _ReduceUnnamedToFieldDef}
	_ReduceStructPaddingToFieldDefAction                                = &_Action{_ReduceAction, 0, _ReduceStructPaddingToFieldDef}
	_ReduceMethodSignatureToFieldDefAction                              = &_Action{_ReduceAction, 0, _ReduceMethodSignatureToFieldDef}
	_ReduceUnsafeStatementToFieldDefAction                              = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToFieldDef}
	_ReduceDefaultToOptionalDefaultAction                               = &_Action{_ReduceAction, 0, _ReduceDefaultToOptionalDefault}
	_ReduceNilToOptionalDefaultAction                                   = &_Action{_ReduceAction, 0, _ReduceNilToOptionalDefault}
	_ReduceAddToProperImplicitFieldDefsAction                           = &_Action{_ReduceAction, 0, _ReduceAddToProperImplicitFieldDefs}
	_ReduceFieldDefToProperImplicitFieldDefsAction                      = &_Action{_ReduceAction, 0, _ReduceFieldDefToProperImplicitFieldDefs}
	_ReduceProperImplicitFieldDefsToImplicitFieldDefsAction             = &_Action{_ReduceAction, 0, _ReduceProperImplicitFieldDefsToImplicitFieldDefs}
	_ReduceImproperToImplicitFieldDefsAction                            = &_Action{_ReduceAction, 0, _ReduceImproperToImplicitFieldDefs}
	_ReduceNilToImplicitFieldDefsAction                                 = &_Action{_ReduceAction, 0, _ReduceNilToImplicitFieldDefs}
	_ReduceToImplicitStructTypeExprAction                               = &_Action{_ReduceAction, 0, _ReduceToImplicitStructTypeExpr}
	_ReduceAddImplicitToProperExplicitFieldDefsAction                   = &_Action{_ReduceAction, 0, _ReduceAddImplicitToProperExplicitFieldDefs}
	_ReduceAddExplicitToProperExplicitFieldDefsAction                   = &_Action{_ReduceAction, 0, _ReduceAddExplicitToProperExplicitFieldDefs}
	_ReduceFieldDefToProperExplicitFieldDefsAction                      = &_Action{_ReduceAction, 0, _ReduceFieldDefToProperExplicitFieldDefs}
	_ReduceProperExplicitFieldDefsToExplicitFieldDefsAction             = &_Action{_ReduceAction, 0, _ReduceProperExplicitFieldDefsToExplicitFieldDefs}
	_ReduceImproperImplicitToExplicitFieldDefsAction                    = &_Action{_ReduceAction, 0, _ReduceImproperImplicitToExplicitFieldDefs}
	_ReduceImproperExplicitToExplicitFieldDefsAction                    = &_Action{_ReduceAction, 0, _ReduceImproperExplicitToExplicitFieldDefs}
	_ReduceNilToExplicitFieldDefsAction                                 = &_Action{_ReduceAction, 0, _ReduceNilToExplicitFieldDefs}
	_ReduceToExplicitStructTypeExprAction                               = &_Action{_ReduceAction, 0, _ReduceToExplicitStructTypeExpr}
	_ReduceToTraitTypeExprAction                                        = &_Action{_ReduceAction, 0, _ReduceToTraitTypeExpr}
	_ReducePairToProperImplicitEnumValueDefsAction                      = &_Action{_ReduceAction, 0, _ReducePairToProperImplicitEnumValueDefs}
	_ReduceAddToProperImplicitEnumValueDefsAction                       = &_Action{_ReduceAction, 0, _ReduceAddToProperImplicitEnumValueDefs}
	_ReduceProperImplicitEnumValueDefsToImplicitEnumValueDefsAction     = &_Action{_ReduceAction, 0, _ReduceProperImplicitEnumValueDefsToImplicitEnumValueDefs}
	_ReduceImproperToImplicitEnumValueDefsAction                        = &_Action{_ReduceAction, 0, _ReduceImproperToImplicitEnumValueDefs}
	_ReduceToImplicitEnumTypeExprAction                                 = &_Action{_ReduceAction, 0, _ReduceToImplicitEnumTypeExpr}
	_ReduceExplicitPairToProperExplicitEnumValueDefsAction              = &_Action{_ReduceAction, 0, _ReduceExplicitPairToProperExplicitEnumValueDefs}
	_ReduceImplicitPairToProperExplicitEnumValueDefsAction              = &_Action{_ReduceAction, 0, _ReduceImplicitPairToProperExplicitEnumValueDefs}
	_ReduceExplicitAddToProperExplicitEnumValueDefsAction               = &_Action{_ReduceAction, 0, _ReduceExplicitAddToProperExplicitEnumValueDefs}
	_ReduceImplicitAddToProperExplicitEnumValueDefsAction               = &_Action{_ReduceAction, 0, _ReduceImplicitAddToProperExplicitEnumValueDefs}
	_ReduceProperExplicitEnumValueDefsToExplicitEnumValueDefsAction     = &_Action{_ReduceAction, 0, _ReduceProperExplicitEnumValueDefsToExplicitEnumValueDefs}
	_ReduceImproperToExplicitEnumValueDefsAction                        = &_Action{_ReduceAction, 0, _ReduceImproperToExplicitEnumValueDefs}
	_ReduceToExplicitEnumTypeExprAction                                 = &_Action{_ReduceAction, 0, _ReduceToExplicitEnumTypeExpr}
	_ReduceReturnableTypeExprToReturnTypeAction                         = &_Action{_ReduceAction, 0, _ReduceReturnableTypeExprToReturnType}
	_ReduceNilToReturnTypeAction                                        = &_Action{_ReduceAction, 0, _ReduceNilToReturnType}
	_ReduceNamedTypedArgToProperParameterDefAction                      = &_Action{_ReduceAction, 0, _ReduceNamedTypedArgToProperParameterDef}
	_ReduceNamedTypedVarargToProperParameterDefAction                   = &_Action{_ReduceAction, 0, _ReduceNamedTypedVarargToProperParameterDef}
	_ReduceNamedInferredVarargToProperParameterDefAction                = &_Action{_ReduceAction, 0, _ReduceNamedInferredVarargToProperParameterDef}
	_ReduceIgnoreTypedArgToProperParameterDefAction                     = &_Action{_ReduceAction, 0, _ReduceIgnoreTypedArgToProperParameterDef}
	_ReduceIgnoreInferredVarargToProperParameterDefAction               = &_Action{_ReduceAction, 0, _ReduceIgnoreInferredVarargToProperParameterDef}
	_ReduceIgnoreTypedVarargToProperParameterDefAction                  = &_Action{_ReduceAction, 0, _ReduceIgnoreTypedVarargToProperParameterDef}
	_ReduceProperParameterDefToParameterDeclAction                      = &_Action{_ReduceAction, 0, _ReduceProperParameterDefToParameterDecl}
	_ReduceUnnamedTypedArgToParameterDeclAction                         = &_Action{_ReduceAction, 0, _ReduceUnnamedTypedArgToParameterDecl}
	_ReduceUnnamedInferredVarargToParameterDeclAction                   = &_Action{_ReduceAction, 0, _ReduceUnnamedInferredVarargToParameterDecl}
	_ReduceUnnamedTypedVarargToParameterDeclAction                      = &_Action{_ReduceAction, 0, _ReduceUnnamedTypedVarargToParameterDecl}
	_ReduceProperParameterDefToParameterDefAction                       = &_Action{_ReduceAction, 0, _ReduceProperParameterDefToParameterDef}
	_ReduceNamedInferredArgToParameterDefAction                         = &_Action{_ReduceAction, 0, _ReduceNamedInferredArgToParameterDef}
	_ReduceIgnoreInferredArgToParameterDefAction                        = &_Action{_ReduceAction, 0, _ReduceIgnoreInferredArgToParameterDef}
	_ReduceAddToProperParameterDeclsAction                              = &_Action{_ReduceAction, 0, _ReduceAddToProperParameterDecls}
	_ReduceParameterDeclToProperParameterDeclsAction                    = &_Action{_ReduceAction, 0, _ReduceParameterDeclToProperParameterDecls}
	_ReduceProperParameterDeclsToParameterDeclsAction                   = &_Action{_ReduceAction, 0, _ReduceProperParameterDeclsToParameterDecls}
	_ReduceImproperToParameterDeclsAction                               = &_Action{_ReduceAction, 0, _ReduceImproperToParameterDecls}
	_ReduceNilToParameterDeclsAction                                    = &_Action{_ReduceAction, 0, _ReduceNilToParameterDecls}
	_ReduceAddToProperParameterDefsAction                               = &_Action{_ReduceAction, 0, _ReduceAddToProperParameterDefs}
	_ReduceParameterDefToProperParameterDefsAction                      = &_Action{_ReduceAction, 0, _ReduceParameterDefToProperParameterDefs}
	_ReduceProperParameterDefsToParameterDefsAction                     = &_Action{_ReduceAction, 0, _ReduceProperParameterDefsToParameterDefs}
	_ReduceImproperToParameterDefsAction                                = &_Action{_ReduceAction, 0, _ReduceImproperToParameterDefs}
	_ReduceNilToParameterDefsAction                                     = &_Action{_ReduceAction, 0, _ReduceNilToParameterDefs}
	_ReduceToFuncTypeExprAction                                         = &_Action{_ReduceAction, 0, _ReduceToFuncTypeExpr}
	_ReduceToMethodSignatureAction                                      = &_Action{_ReduceAction, 0, _ReduceToMethodSignature}
	_ReduceFuncDefToNamedFuncDefAction                                  = &_Action{_ReduceAction, 0, _ReduceFuncDefToNamedFuncDef}
	_ReduceMethodDefToNamedFuncDefAction                                = &_Action{_ReduceAction, 0, _ReduceMethodDefToNamedFuncDef}
	_ReduceAliasToNamedFuncDefAction                                    = &_Action{_ReduceAction, 0, _ReduceAliasToNamedFuncDef}
	_ReduceToAnonymousFuncExprAction                                    = &_Action{_ReduceAction, 0, _ReduceToAnonymousFuncExpr}
	_ReduceNoSpecToPackageDefAction                                     = &_Action{_ReduceAction, 0, _ReduceNoSpecToPackageDef}
	_ReduceWithSpecToPackageDefAction                                   = &_Action{_ReduceAction, 0, _ReduceWithSpecToPackageDef}
)

var _ActionTable = _ActionTableType{
	{_State5, _EndMarker}:                          &_Action{_AcceptAction, 0, 0},
	{_State6, _EndMarker}:                          &_Action{_AcceptAction, 0, 0},
	{_State7, _EndMarker}:                          &_Action{_AcceptAction, 0, 0},
	{_State8, _EndMarker}:                          &_Action{_AcceptAction, 0, 0},
	{_State1, CommentGroupsToken}:                  _GotoState9Action,
	{_State1, PackageToken}:                        _GotoState13Action,
	{_State1, TypeToken}:                           _GotoState14Action,
	{_State1, FuncToken}:                           _GotoState10Action,
	{_State1, VarToken}:                            _GotoState15Action,
	{_State1, LetToken}:                            _GotoState12Action,
	{_State1, LbraceToken}:                         _GotoState11Action,
	{_State1, SourceType}:                          _GotoState5Action,
	{_State1, ProperDefinitionsType}:               _GotoState20Action,
	{_State1, DefinitionsType}:                     _GotoState17Action,
	{_State1, DefinitionType}:                      _GotoState16Action,
	{_State1, StatementBlockType}:                  _GotoState21Action,
	{_State1, VarDeclPatternType}:                  _GotoState23Action,
	{_State1, VarOrLetType}:                        _GotoState24Action,
	{_State1, TypeDefType}:                         _GotoState22Action,
	{_State1, NamedFuncDefType}:                    _GotoState18Action,
	{_State1, PackageDefType}:                      _GotoState19Action,
	{_State2, IntegerLiteralToken}:                 _GotoState40Action,
	{_State2, FloatLiteralToken}:                   _GotoState35Action,
	{_State2, RuneLiteralToken}:                    _GotoState48Action,
	{_State2, StringLiteralToken}:                  _GotoState49Action,
	{_State2, IdentifierToken}:                     _GotoState38Action,
	{_State2, UnderscoreToken}:                     _GotoState53Action,
	{_State2, TrueToken}:                           _GotoState52Action,
	{_State2, FalseToken}:                          _GotoState34Action,
	{_State2, CaseToken}:                           _GotoState29Action,
	{_State2, DefaultToken}:                        _GotoState31Action,
	{_State2, ReturnToken}:                         _GotoState47Action,
	{_State2, BreakToken}:                          _GotoState28Action,
	{_State2, ContinueToken}:                       _GotoState30Action,
	{_State2, FallthroughToken}:                    _GotoState33Action,
	{_State2, ImportToken}:                         _GotoState39Action,
	{_State2, UnsafeToken}:                         _GotoState54Action,
	{_State2, StructToken}:                         _GotoState50Action,
	{_State2, FuncToken}:                           _GotoState36Action,
	{_State2, AsyncToken}:                          _GotoState25Action,
	{_State2, DeferToken}:                          _GotoState32Action,
	{_State2, VarToken}:                            _GotoState15Action,
	{_State2, LetToken}:                            _GotoState12Action,
	{_State2, NotToken}:                            _GotoState45Action,
	{_State2, LabelDeclToken}:                      _GotoState41Action,
	{_State2, LparenToken}:                         _GotoState43Action,
	{_State2, LbracketToken}:                       _GotoState42Action,
	{_State2, SubToken}:                            _GotoState51Action,
	{_State2, MulToken}:                            _GotoState44Action,
	{_State2, BitNegToken}:                         _GotoState27Action,
	{_State2, BitAndToken}:                         _GotoState26Action,
	{_State2, GreaterToken}:                        _GotoState37Action,
	{_State2, ParseErrorToken}:                     _GotoState46Action,
	{_State2, SimpleStatementType}:                 _GotoState100Action,
	{_State2, StatementType}:                       _GotoState6Action,
	{_State2, ExprOrImproperStructStatementType}:   _GotoState77Action,
	{_State2, ExprsType}:                           _GotoState78Action,
	{_State2, CallbackOpType}:                      _GotoState72Action,
	{_State2, CallbackOpStatementType}:             _GotoState73Action,
	{_State2, UnsafeStatementType}:                 _GotoState103Action,
	{_State2, JumpStatementType}:                   _GotoState85Action,
	{_State2, JumpTypeType}:                        _GotoState86Action,
	{_State2, FallthroughStatementType}:            _GotoState79Action,
	{_State2, AssignStatementType}:                 _GotoState62Action,
	{_State2, UnaryOpAssignStatementType}:          _GotoState102Action,
	{_State2, BinaryOpAssignStatementType}:         _GotoState68Action,
	{_State2, ImportStatementType}:                 _GotoState81Action,
	{_State2, VarDeclPatternType}:                  _GotoState104Action,
	{_State2, VarOrLetType}:                        _GotoState24Action,
	{_State2, AssignPatternType}:                   _GotoState61Action,
	{_State2, ExprType}:                            _GotoState76Action,
	{_State2, OptionalLabelDeclType}:               _GotoState91Action,
	{_State2, SequenceExprType}:                    _GotoState99Action,
	{_State2, CallExprType}:                        _GotoState71Action,
	{_State2, AtomExprType}:                        _GotoState63Action,
	{_State2, ParseErrorExprType}:                  _GotoState93Action,
	{_State2, LiteralExprType}:                     _GotoState87Action,
	{_State2, NamedExprType}:                       _GotoState90Action,
	{_State2, BlockExprType}:                       _GotoState70Action,
	{_State2, InitializeExprType}:                  _GotoState84Action,
	{_State2, ImplicitStructExprType}:              _GotoState80Action,
	{_State2, AccessibleExprType}:                  _GotoState56Action,
	{_State2, AccessExprType}:                      _GotoState55Action,
	{_State2, IndexExprType}:                       _GotoState82Action,
	{_State2, PostfixableExprType}:                 _GotoState95Action,
	{_State2, PostfixUnaryExprType}:                _GotoState94Action,
	{_State2, PrefixableExprType}:                  _GotoState98Action,
	{_State2, PrefixUnaryExprType}:                 _GotoState96Action,
	{_State2, PrefixUnaryOpType}:                   _GotoState97Action,
	{_State2, MulExprType}:                         _GotoState89Action,
	{_State2, BinaryMulExprType}:                   _GotoState67Action,
	{_State2, AddExprType}:                         _GotoState57Action,
	{_State2, BinaryAddExprType}:                   _GotoState64Action,
	{_State2, CmpExprType}:                         _GotoState74Action,
	{_State2, BinaryCmpExprType}:                   _GotoState66Action,
	{_State2, AndExprType}:                         _GotoState58Action,
	{_State2, BinaryAndExprType}:                   _GotoState65Action,
	{_State2, OrExprType}:                          _GotoState92Action,
	{_State2, BinaryOrExprType}:                    _GotoState69Action,
	{_State2, InitializableTypeExprType}:           _GotoState83Action,
	{_State2, SliceTypeExprType}:                   _GotoState101Action,
	{_State2, ArrayTypeExprType}:                   _GotoState60Action,
	{_State2, MapTypeExprType}:                     _GotoState88Action,
	{_State2, ExplicitStructTypeExprType}:          _GotoState75Action,
	{_State2, AnonymousFuncExprType}:               _GotoState59Action,
	{_State3, IntegerLiteralToken}:                 _GotoState40Action,
	{_State3, FloatLiteralToken}:                   _GotoState35Action,
	{_State3, RuneLiteralToken}:                    _GotoState48Action,
	{_State3, StringLiteralToken}:                  _GotoState49Action,
	{_State3, IdentifierToken}:                     _GotoState38Action,
	{_State3, UnderscoreToken}:                     _GotoState53Action,
	{_State3, TrueToken}:                           _GotoState52Action,
	{_State3, FalseToken}:                          _GotoState34Action,
	{_State3, StructToken}:                         _GotoState50Action,
	{_State3, FuncToken}:                           _GotoState36Action,
	{_State3, VarToken}:                            _GotoState15Action,
	{_State3, LetToken}:                            _GotoState12Action,
	{_State3, NotToken}:                            _GotoState45Action,
	{_State3, LabelDeclToken}:                      _GotoState41Action,
	{_State3, LparenToken}:                         _GotoState43Action,
	{_State3, LbracketToken}:                       _GotoState42Action,
	{_State3, SubToken}:                            _GotoState51Action,
	{_State3, MulToken}:                            _GotoState44Action,
	{_State3, BitNegToken}:                         _GotoState27Action,
	{_State3, BitAndToken}:                         _GotoState26Action,
	{_State3, GreaterToken}:                        _GotoState37Action,
	{_State3, ParseErrorToken}:                     _GotoState46Action,
	{_State3, VarDeclPatternType}:                  _GotoState104Action,
	{_State3, VarOrLetType}:                        _GotoState24Action,
	{_State3, ExprType}:                            _GotoState7Action,
	{_State3, OptionalLabelDeclType}:               _GotoState91Action,
	{_State3, SequenceExprType}:                    _GotoState106Action,
	{_State3, CallExprType}:                        _GotoState71Action,
	{_State3, AtomExprType}:                        _GotoState63Action,
	{_State3, ParseErrorExprType}:                  _GotoState93Action,
	{_State3, LiteralExprType}:                     _GotoState87Action,
	{_State3, NamedExprType}:                       _GotoState90Action,
	{_State3, BlockExprType}:                       _GotoState70Action,
	{_State3, InitializeExprType}:                  _GotoState84Action,
	{_State3, ImplicitStructExprType}:              _GotoState80Action,
	{_State3, AccessibleExprType}:                  _GotoState105Action,
	{_State3, AccessExprType}:                      _GotoState55Action,
	{_State3, IndexExprType}:                       _GotoState82Action,
	{_State3, PostfixableExprType}:                 _GotoState95Action,
	{_State3, PostfixUnaryExprType}:                _GotoState94Action,
	{_State3, PrefixableExprType}:                  _GotoState98Action,
	{_State3, PrefixUnaryExprType}:                 _GotoState96Action,
	{_State3, PrefixUnaryOpType}:                   _GotoState97Action,
	{_State3, MulExprType}:                         _GotoState89Action,
	{_State3, BinaryMulExprType}:                   _GotoState67Action,
	{_State3, AddExprType}:                         _GotoState57Action,
	{_State3, BinaryAddExprType}:                   _GotoState64Action,
	{_State3, CmpExprType}:                         _GotoState74Action,
	{_State3, BinaryCmpExprType}:                   _GotoState66Action,
	{_State3, AndExprType}:                         _GotoState58Action,
	{_State3, BinaryAndExprType}:                   _GotoState65Action,
	{_State3, OrExprType}:                          _GotoState92Action,
	{_State3, BinaryOrExprType}:                    _GotoState69Action,
	{_State3, InitializableTypeExprType}:           _GotoState83Action,
	{_State3, SliceTypeExprType}:                   _GotoState101Action,
	{_State3, ArrayTypeExprType}:                   _GotoState60Action,
	{_State3, MapTypeExprType}:                     _GotoState88Action,
	{_State3, ExplicitStructTypeExprType}:          _GotoState75Action,
	{_State3, AnonymousFuncExprType}:               _GotoState59Action,
	{_State4, IdentifierToken}:                     _GotoState113Action,
	{_State4, UnderscoreToken}:                     _GotoState119Action,
	{_State4, StructToken}:                         _GotoState50Action,
	{_State4, EnumToken}:                           _GotoState110Action,
	{_State4, TraitToken}:                          _GotoState118Action,
	{_State4, FuncToken}:                           _GotoState112Action,
	{_State4, LparenToken}:                         _GotoState114Action,
	{_State4, LbracketToken}:                       _GotoState42Action,
	{_State4, DotToken}:                            _GotoState109Action,
	{_State4, QuestionToken}:                       _GotoState116Action,
	{_State4, ExclaimToken}:                        _GotoState111Action,
	{_State4, TildeTildeToken}:                     _GotoState117Action,
	{_State4, BitNegToken}:                         _GotoState108Action,
	{_State4, BitAndToken}:                         _GotoState107Action,
	{_State4, ParseErrorToken}:                     _GotoState115Action,
	{_State4, InitializableTypeExprType}:           _GotoState127Action,
	{_State4, SliceTypeExprType}:                   _GotoState101Action,
	{_State4, ArrayTypeExprType}:                   _GotoState60Action,
	{_State4, MapTypeExprType}:                     _GotoState88Action,
	{_State4, AtomTypeExprType}:                    _GotoState120Action,
	{_State4, NamedTypeExprType}:                   _GotoState128Action,
	{_State4, InferredTypeExprType}:                _GotoState126Action,
	{_State4, ParseErrorTypeExprType}:              _GotoState129Action,
	{_State4, ReturnableTypeExprType}:              _GotoState132Action,
	{_State4, PrefixUnaryTypeExprType}:             _GotoState130Action,
	{_State4, PrefixUnaryTypeOpType}:               _GotoState131Action,
	{_State4, TypeExprType}:                        _GotoState8Action,
	{_State4, BinaryTypeExprType}:                  _GotoState121Action,
	{_State4, ImplicitStructTypeExprType}:          _GotoState125Action,
	{_State4, ExplicitStructTypeExprType}:          _GotoState75Action,
	{_State4, TraitTypeExprType}:                   _GotoState133Action,
	{_State4, ImplicitEnumTypeExprType}:            _GotoState124Action,
	{_State4, ExplicitEnumTypeExprType}:            _GotoState122Action,
	{_State4, FuncTypeExprType}:                    _GotoState123Action,
	{_State8, AddToken}:                            _GotoState249Action,
	{_State8, SubToken}:                            _GotoState251Action,
	{_State8, MulToken}:                            _GotoState250Action,
	{_State8, BinaryTypeOpType}:                    _GotoState252Action,
	{_State10, IdentifierToken}:                    _GotoState134Action,
	{_State10, LparenToken}:                        _GotoState135Action,
	{_State11, IntegerLiteralToken}:                _GotoState40Action,
	{_State11, FloatLiteralToken}:                  _GotoState35Action,
	{_State11, RuneLiteralToken}:                   _GotoState48Action,
	{_State11, StringLiteralToken}:                 _GotoState49Action,
	{_State11, IdentifierToken}:                    _GotoState38Action,
	{_State11, UnderscoreToken}:                    _GotoState53Action,
	{_State11, TrueToken}:                          _GotoState52Action,
	{_State11, FalseToken}:                         _GotoState34Action,
	{_State11, CaseToken}:                          _GotoState29Action,
	{_State11, DefaultToken}:                       _GotoState31Action,
	{_State11, ReturnToken}:                        _GotoState47Action,
	{_State11, BreakToken}:                         _GotoState28Action,
	{_State11, ContinueToken}:                      _GotoState30Action,
	{_State11, FallthroughToken}:                   _GotoState33Action,
	{_State11, ImportToken}:                        _GotoState39Action,
	{_State11, UnsafeToken}:                        _GotoState54Action,
	{_State11, StructToken}:                        _GotoState50Action,
	{_State11, FuncToken}:                          _GotoState36Action,
	{_State11, AsyncToken}:                         _GotoState25Action,
	{_State11, DeferToken}:                         _GotoState32Action,
	{_State11, VarToken}:                           _GotoState15Action,
	{_State11, LetToken}:                           _GotoState12Action,
	{_State11, NotToken}:                           _GotoState45Action,
	{_State11, LabelDeclToken}:                     _GotoState41Action,
	{_State11, LparenToken}:                        _GotoState43Action,
	{_State11, LbracketToken}:                      _GotoState42Action,
	{_State11, SubToken}:                           _GotoState51Action,
	{_State11, MulToken}:                           _GotoState44Action,
	{_State11, BitNegToken}:                        _GotoState27Action,
	{_State11, BitAndToken}:                        _GotoState26Action,
	{_State11, GreaterToken}:                       _GotoState37Action,
	{_State11, ParseErrorToken}:                    _GotoState46Action,
	{_State11, ProperStatementsType}:               _GotoState136Action,
	{_State11, StatementsType}:                     _GotoState138Action,
	{_State11, SimpleStatementType}:                _GotoState100Action,
	{_State11, StatementType}:                      _GotoState137Action,
	{_State11, ExprOrImproperStructStatementType}:  _GotoState77Action,
	{_State11, ExprsType}:                          _GotoState78Action,
	{_State11, CallbackOpType}:                     _GotoState72Action,
	{_State11, CallbackOpStatementType}:            _GotoState73Action,
	{_State11, UnsafeStatementType}:                _GotoState103Action,
	{_State11, JumpStatementType}:                  _GotoState85Action,
	{_State11, JumpTypeType}:                       _GotoState86Action,
	{_State11, FallthroughStatementType}:           _GotoState79Action,
	{_State11, AssignStatementType}:                _GotoState62Action,
	{_State11, UnaryOpAssignStatementType}:         _GotoState102Action,
	{_State11, BinaryOpAssignStatementType}:        _GotoState68Action,
	{_State11, ImportStatementType}:                _GotoState81Action,
	{_State11, VarDeclPatternType}:                 _GotoState104Action,
	{_State11, VarOrLetType}:                       _GotoState24Action,
	{_State11, AssignPatternType}:                  _GotoState61Action,
	{_State11, ExprType}:                           _GotoState76Action,
	{_State11, OptionalLabelDeclType}:              _GotoState91Action,
	{_State11, SequenceExprType}:                   _GotoState99Action,
	{_State11, CallExprType}:                       _GotoState71Action,
	{_State11, AtomExprType}:                       _GotoState63Action,
	{_State11, ParseErrorExprType}:                 _GotoState93Action,
	{_State11, LiteralExprType}:                    _GotoState87Action,
	{_State11, NamedExprType}:                      _GotoState90Action,
	{_State11, BlockExprType}:                      _GotoState70Action,
	{_State11, InitializeExprType}:                 _GotoState84Action,
	{_State11, ImplicitStructExprType}:             _GotoState80Action,
	{_State11, AccessibleExprType}:                 _GotoState56Action,
	{_State11, AccessExprType}:                     _GotoState55Action,
	{_State11, IndexExprType}:                      _GotoState82Action,
	{_State11, PostfixableExprType}:                _GotoState95Action,
	{_State11, PostfixUnaryExprType}:               _GotoState94Action,
	{_State11, PrefixableExprType}:                 _GotoState98Action,
	{_State11, PrefixUnaryExprType}:                _GotoState96Action,
	{_State11, PrefixUnaryOpType}:                  _GotoState97Action,
	{_State11, MulExprType}:                        _GotoState89Action,
	{_State11, BinaryMulExprType}:                  _GotoState67Action,
	{_State11, AddExprType}:                        _GotoState57Action,
	{_State11, BinaryAddExprType}:                  _GotoState64Action,
	{_State11, CmpExprType}:                        _GotoState74Action,
	{_State11, BinaryCmpExprType}:                  _GotoState66Action,
	{_State11, AndExprType}:                        _GotoState58Action,
	{_State11, BinaryAndExprType}:                  _GotoState65Action,
	{_State11, OrExprType}:                         _GotoState92Action,
	{_State11, BinaryOrExprType}:                   _GotoState69Action,
	{_State11, InitializableTypeExprType}:          _GotoState83Action,
	{_State11, SliceTypeExprType}:                  _GotoState101Action,
	{_State11, ArrayTypeExprType}:                  _GotoState60Action,
	{_State11, MapTypeExprType}:                    _GotoState88Action,
	{_State11, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State11, AnonymousFuncExprType}:              _GotoState59Action,
	{_State13, LbraceToken}:                        _GotoState11Action,
	{_State13, StatementBlockType}:                 _GotoState139Action,
	{_State14, IdentifierToken}:                    _GotoState140Action,
	{_State20, NewlinesToken}:                      _GotoState141Action,
	{_State23, AssignToken}:                        _GotoState142Action,
	{_State24, IdentifierToken}:                    _GotoState143Action,
	{_State24, UnderscoreToken}:                    _GotoState145Action,
	{_State24, LparenToken}:                        _GotoState144Action,
	{_State24, VarPatternType}:                     _GotoState147Action,
	{_State24, TuplePatternType}:                   _GotoState146Action,
	{_State29, IntegerLiteralToken}:                _GotoState40Action,
	{_State29, FloatLiteralToken}:                  _GotoState35Action,
	{_State29, RuneLiteralToken}:                   _GotoState48Action,
	{_State29, StringLiteralToken}:                 _GotoState49Action,
	{_State29, IdentifierToken}:                    _GotoState38Action,
	{_State29, UnderscoreToken}:                    _GotoState53Action,
	{_State29, TrueToken}:                          _GotoState52Action,
	{_State29, FalseToken}:                         _GotoState34Action,
	{_State29, StructToken}:                        _GotoState50Action,
	{_State29, FuncToken}:                          _GotoState36Action,
	{_State29, VarToken}:                           _GotoState149Action,
	{_State29, LetToken}:                           _GotoState12Action,
	{_State29, NotToken}:                           _GotoState45Action,
	{_State29, LabelDeclToken}:                     _GotoState41Action,
	{_State29, LparenToken}:                        _GotoState43Action,
	{_State29, LbracketToken}:                      _GotoState42Action,
	{_State29, DotToken}:                           _GotoState148Action,
	{_State29, SubToken}:                           _GotoState51Action,
	{_State29, MulToken}:                           _GotoState44Action,
	{_State29, BitNegToken}:                        _GotoState27Action,
	{_State29, BitAndToken}:                        _GotoState26Action,
	{_State29, GreaterToken}:                       _GotoState37Action,
	{_State29, ParseErrorToken}:                    _GotoState46Action,
	{_State29, CasePatternsType}:                   _GotoState152Action,
	{_State29, VarDeclPatternType}:                 _GotoState104Action,
	{_State29, VarOrLetType}:                       _GotoState24Action,
	{_State29, AssignPatternType}:                  _GotoState150Action,
	{_State29, CasePatternType}:                    _GotoState151Action,
	{_State29, OptionalLabelDeclType}:              _GotoState153Action,
	{_State29, SequenceExprType}:                   _GotoState154Action,
	{_State29, CallExprType}:                       _GotoState71Action,
	{_State29, AtomExprType}:                       _GotoState63Action,
	{_State29, ParseErrorExprType}:                 _GotoState93Action,
	{_State29, LiteralExprType}:                    _GotoState87Action,
	{_State29, NamedExprType}:                      _GotoState90Action,
	{_State29, BlockExprType}:                      _GotoState70Action,
	{_State29, InitializeExprType}:                 _GotoState84Action,
	{_State29, ImplicitStructExprType}:             _GotoState80Action,
	{_State29, AccessibleExprType}:                 _GotoState105Action,
	{_State29, AccessExprType}:                     _GotoState55Action,
	{_State29, IndexExprType}:                      _GotoState82Action,
	{_State29, PostfixableExprType}:                _GotoState95Action,
	{_State29, PostfixUnaryExprType}:               _GotoState94Action,
	{_State29, PrefixableExprType}:                 _GotoState98Action,
	{_State29, PrefixUnaryExprType}:                _GotoState96Action,
	{_State29, PrefixUnaryOpType}:                  _GotoState97Action,
	{_State29, MulExprType}:                        _GotoState89Action,
	{_State29, BinaryMulExprType}:                  _GotoState67Action,
	{_State29, AddExprType}:                        _GotoState57Action,
	{_State29, BinaryAddExprType}:                  _GotoState64Action,
	{_State29, CmpExprType}:                        _GotoState74Action,
	{_State29, BinaryCmpExprType}:                  _GotoState66Action,
	{_State29, AndExprType}:                        _GotoState58Action,
	{_State29, BinaryAndExprType}:                  _GotoState65Action,
	{_State29, OrExprType}:                         _GotoState92Action,
	{_State29, BinaryOrExprType}:                   _GotoState69Action,
	{_State29, InitializableTypeExprType}:          _GotoState83Action,
	{_State29, SliceTypeExprType}:                  _GotoState101Action,
	{_State29, ArrayTypeExprType}:                  _GotoState60Action,
	{_State29, MapTypeExprType}:                    _GotoState88Action,
	{_State29, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State29, AnonymousFuncExprType}:              _GotoState59Action,
	{_State31, ColonToken}:                         _GotoState155Action,
	{_State36, LparenToken}:                        _GotoState156Action,
	{_State37, IntegerLiteralToken}:                _GotoState40Action,
	{_State37, FloatLiteralToken}:                  _GotoState35Action,
	{_State37, RuneLiteralToken}:                   _GotoState48Action,
	{_State37, StringLiteralToken}:                 _GotoState49Action,
	{_State37, IdentifierToken}:                    _GotoState38Action,
	{_State37, UnderscoreToken}:                    _GotoState53Action,
	{_State37, TrueToken}:                          _GotoState52Action,
	{_State37, FalseToken}:                         _GotoState34Action,
	{_State37, StructToken}:                        _GotoState50Action,
	{_State37, FuncToken}:                          _GotoState36Action,
	{_State37, VarToken}:                           _GotoState15Action,
	{_State37, LetToken}:                           _GotoState12Action,
	{_State37, NotToken}:                           _GotoState45Action,
	{_State37, LabelDeclToken}:                     _GotoState41Action,
	{_State37, LparenToken}:                        _GotoState43Action,
	{_State37, LbracketToken}:                      _GotoState42Action,
	{_State37, SubToken}:                           _GotoState51Action,
	{_State37, MulToken}:                           _GotoState44Action,
	{_State37, BitNegToken}:                        _GotoState27Action,
	{_State37, BitAndToken}:                        _GotoState26Action,
	{_State37, GreaterToken}:                       _GotoState37Action,
	{_State37, ParseErrorToken}:                    _GotoState46Action,
	{_State37, VarDeclPatternType}:                 _GotoState104Action,
	{_State37, VarOrLetType}:                       _GotoState24Action,
	{_State37, OptionalLabelDeclType}:              _GotoState153Action,
	{_State37, SequenceExprType}:                   _GotoState157Action,
	{_State37, CallExprType}:                       _GotoState71Action,
	{_State37, AtomExprType}:                       _GotoState63Action,
	{_State37, ParseErrorExprType}:                 _GotoState93Action,
	{_State37, LiteralExprType}:                    _GotoState87Action,
	{_State37, NamedExprType}:                      _GotoState90Action,
	{_State37, BlockExprType}:                      _GotoState70Action,
	{_State37, InitializeExprType}:                 _GotoState84Action,
	{_State37, ImplicitStructExprType}:             _GotoState80Action,
	{_State37, AccessibleExprType}:                 _GotoState105Action,
	{_State37, AccessExprType}:                     _GotoState55Action,
	{_State37, IndexExprType}:                      _GotoState82Action,
	{_State37, PostfixableExprType}:                _GotoState95Action,
	{_State37, PostfixUnaryExprType}:               _GotoState94Action,
	{_State37, PrefixableExprType}:                 _GotoState98Action,
	{_State37, PrefixUnaryExprType}:                _GotoState96Action,
	{_State37, PrefixUnaryOpType}:                  _GotoState97Action,
	{_State37, MulExprType}:                        _GotoState89Action,
	{_State37, BinaryMulExprType}:                  _GotoState67Action,
	{_State37, AddExprType}:                        _GotoState57Action,
	{_State37, BinaryAddExprType}:                  _GotoState64Action,
	{_State37, CmpExprType}:                        _GotoState74Action,
	{_State37, BinaryCmpExprType}:                  _GotoState66Action,
	{_State37, AndExprType}:                        _GotoState58Action,
	{_State37, BinaryAndExprType}:                  _GotoState65Action,
	{_State37, OrExprType}:                         _GotoState92Action,
	{_State37, BinaryOrExprType}:                   _GotoState69Action,
	{_State37, InitializableTypeExprType}:          _GotoState83Action,
	{_State37, SliceTypeExprType}:                  _GotoState101Action,
	{_State37, ArrayTypeExprType}:                  _GotoState60Action,
	{_State37, MapTypeExprType}:                    _GotoState88Action,
	{_State37, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State37, AnonymousFuncExprType}:              _GotoState59Action,
	{_State39, StringLiteralToken}:                 _GotoState161Action,
	{_State39, IdentifierToken}:                    _GotoState159Action,
	{_State39, UnderscoreToken}:                    _GotoState162Action,
	{_State39, LparenToken}:                        _GotoState160Action,
	{_State39, DotToken}:                           _GotoState158Action,
	{_State39, ImportClauseType}:                   _GotoState163Action,
	{_State42, IdentifierToken}:                    _GotoState113Action,
	{_State42, UnderscoreToken}:                    _GotoState119Action,
	{_State42, StructToken}:                        _GotoState50Action,
	{_State42, EnumToken}:                          _GotoState110Action,
	{_State42, TraitToken}:                         _GotoState118Action,
	{_State42, FuncToken}:                          _GotoState112Action,
	{_State42, LparenToken}:                        _GotoState114Action,
	{_State42, LbracketToken}:                      _GotoState42Action,
	{_State42, DotToken}:                           _GotoState109Action,
	{_State42, QuestionToken}:                      _GotoState116Action,
	{_State42, ExclaimToken}:                       _GotoState111Action,
	{_State42, TildeTildeToken}:                    _GotoState117Action,
	{_State42, BitNegToken}:                        _GotoState108Action,
	{_State42, BitAndToken}:                        _GotoState107Action,
	{_State42, ParseErrorToken}:                    _GotoState115Action,
	{_State42, InitializableTypeExprType}:          _GotoState127Action,
	{_State42, SliceTypeExprType}:                  _GotoState101Action,
	{_State42, ArrayTypeExprType}:                  _GotoState60Action,
	{_State42, MapTypeExprType}:                    _GotoState88Action,
	{_State42, AtomTypeExprType}:                   _GotoState120Action,
	{_State42, NamedTypeExprType}:                  _GotoState128Action,
	{_State42, InferredTypeExprType}:               _GotoState126Action,
	{_State42, ParseErrorTypeExprType}:             _GotoState129Action,
	{_State42, ReturnableTypeExprType}:             _GotoState132Action,
	{_State42, PrefixUnaryTypeExprType}:            _GotoState130Action,
	{_State42, PrefixUnaryTypeOpType}:              _GotoState131Action,
	{_State42, TypeExprType}:                       _GotoState164Action,
	{_State42, BinaryTypeExprType}:                 _GotoState121Action,
	{_State42, ImplicitStructTypeExprType}:         _GotoState125Action,
	{_State42, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State42, TraitTypeExprType}:                  _GotoState133Action,
	{_State42, ImplicitEnumTypeExprType}:           _GotoState124Action,
	{_State42, ExplicitEnumTypeExprType}:           _GotoState122Action,
	{_State42, FuncTypeExprType}:                   _GotoState123Action,
	{_State43, IntegerLiteralToken}:                _GotoState40Action,
	{_State43, FloatLiteralToken}:                  _GotoState35Action,
	{_State43, RuneLiteralToken}:                   _GotoState48Action,
	{_State43, StringLiteralToken}:                 _GotoState49Action,
	{_State43, IdentifierToken}:                    _GotoState167Action,
	{_State43, UnderscoreToken}:                    _GotoState53Action,
	{_State43, TrueToken}:                          _GotoState52Action,
	{_State43, FalseToken}:                         _GotoState34Action,
	{_State43, StructToken}:                        _GotoState50Action,
	{_State43, FuncToken}:                          _GotoState36Action,
	{_State43, VarToken}:                           _GotoState15Action,
	{_State43, LetToken}:                           _GotoState12Action,
	{_State43, NotToken}:                           _GotoState45Action,
	{_State43, LabelDeclToken}:                     _GotoState41Action,
	{_State43, LparenToken}:                        _GotoState43Action,
	{_State43, LbracketToken}:                      _GotoState42Action,
	{_State43, ColonToken}:                         _GotoState165Action,
	{_State43, EllipsisToken}:                      _GotoState166Action,
	{_State43, SubToken}:                           _GotoState51Action,
	{_State43, MulToken}:                           _GotoState44Action,
	{_State43, BitNegToken}:                        _GotoState27Action,
	{_State43, BitAndToken}:                        _GotoState26Action,
	{_State43, GreaterToken}:                       _GotoState37Action,
	{_State43, ParseErrorToken}:                    _GotoState46Action,
	{_State43, VarDeclPatternType}:                 _GotoState104Action,
	{_State43, VarOrLetType}:                       _GotoState24Action,
	{_State43, ExprType}:                           _GotoState171Action,
	{_State43, OptionalLabelDeclType}:              _GotoState91Action,
	{_State43, SequenceExprType}:                   _GotoState106Action,
	{_State43, CallExprType}:                       _GotoState71Action,
	{_State43, ProperArgumentsType}:                _GotoState172Action,
	{_State43, ArgumentsType}:                      _GotoState169Action,
	{_State43, ArgumentType}:                       _GotoState168Action,
	{_State43, ColonExprType}:                      _GotoState170Action,
	{_State43, AtomExprType}:                       _GotoState63Action,
	{_State43, ParseErrorExprType}:                 _GotoState93Action,
	{_State43, LiteralExprType}:                    _GotoState87Action,
	{_State43, NamedExprType}:                      _GotoState90Action,
	{_State43, BlockExprType}:                      _GotoState70Action,
	{_State43, InitializeExprType}:                 _GotoState84Action,
	{_State43, ImplicitStructExprType}:             _GotoState80Action,
	{_State43, AccessibleExprType}:                 _GotoState105Action,
	{_State43, AccessExprType}:                     _GotoState55Action,
	{_State43, IndexExprType}:                      _GotoState82Action,
	{_State43, PostfixableExprType}:                _GotoState95Action,
	{_State43, PostfixUnaryExprType}:               _GotoState94Action,
	{_State43, PrefixableExprType}:                 _GotoState98Action,
	{_State43, PrefixUnaryExprType}:                _GotoState96Action,
	{_State43, PrefixUnaryOpType}:                  _GotoState97Action,
	{_State43, MulExprType}:                        _GotoState89Action,
	{_State43, BinaryMulExprType}:                  _GotoState67Action,
	{_State43, AddExprType}:                        _GotoState57Action,
	{_State43, BinaryAddExprType}:                  _GotoState64Action,
	{_State43, CmpExprType}:                        _GotoState74Action,
	{_State43, BinaryCmpExprType}:                  _GotoState66Action,
	{_State43, AndExprType}:                        _GotoState58Action,
	{_State43, BinaryAndExprType}:                  _GotoState65Action,
	{_State43, OrExprType}:                         _GotoState92Action,
	{_State43, BinaryOrExprType}:                   _GotoState69Action,
	{_State43, InitializableTypeExprType}:          _GotoState83Action,
	{_State43, SliceTypeExprType}:                  _GotoState101Action,
	{_State43, ArrayTypeExprType}:                  _GotoState60Action,
	{_State43, MapTypeExprType}:                    _GotoState88Action,
	{_State43, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State43, AnonymousFuncExprType}:              _GotoState59Action,
	{_State50, LparenToken}:                        _GotoState173Action,
	{_State54, LessToken}:                          _GotoState174Action,
	{_State56, LbracketToken}:                      _GotoState186Action,
	{_State56, DotToken}:                           _GotoState185Action,
	{_State56, QuestionToken}:                      _GotoState189Action,
	{_State56, DollarLbracketToken}:                _GotoState184Action,
	{_State56, AddAssignToken}:                     _GotoState175Action,
	{_State56, SubAssignToken}:                     _GotoState190Action,
	{_State56, MulAssignToken}:                     _GotoState188Action,
	{_State56, DivAssignToken}:                     _GotoState183Action,
	{_State56, ModAssignToken}:                     _GotoState187Action,
	{_State56, AddOneAssignToken}:                  _GotoState176Action,
	{_State56, SubOneAssignToken}:                  _GotoState191Action,
	{_State56, BitNegAssignToken}:                  _GotoState179Action,
	{_State56, BitAndAssignToken}:                  _GotoState177Action,
	{_State56, BitOrAssignToken}:                   _GotoState180Action,
	{_State56, BitXorAssignToken}:                  _GotoState182Action,
	{_State56, BitLshiftAssignToken}:               _GotoState178Action,
	{_State56, BitRshiftAssignToken}:               _GotoState181Action,
	{_State56, UnaryOpAssignType}:                  _GotoState194Action,
	{_State56, BinaryOpAssignType}:                 _GotoState192Action,
	{_State56, GenericTypeArgumentsType}:           _GotoState193Action,
	{_State57, AddToken}:                           _GotoState195Action,
	{_State57, SubToken}:                           _GotoState198Action,
	{_State57, BitXorToken}:                        _GotoState197Action,
	{_State57, BitOrToken}:                         _GotoState196Action,
	{_State57, AddOpType}:                          _GotoState199Action,
	{_State58, AndToken}:                           _GotoState200Action,
	{_State61, AssignToken}:                        _GotoState201Action,
	{_State72, IntegerLiteralToken}:                _GotoState40Action,
	{_State72, FloatLiteralToken}:                  _GotoState35Action,
	{_State72, RuneLiteralToken}:                   _GotoState48Action,
	{_State72, StringLiteralToken}:                 _GotoState49Action,
	{_State72, IdentifierToken}:                    _GotoState38Action,
	{_State72, UnderscoreToken}:                    _GotoState53Action,
	{_State72, TrueToken}:                          _GotoState52Action,
	{_State72, FalseToken}:                         _GotoState34Action,
	{_State72, StructToken}:                        _GotoState50Action,
	{_State72, FuncToken}:                          _GotoState36Action,
	{_State72, LabelDeclToken}:                     _GotoState41Action,
	{_State72, LparenToken}:                        _GotoState43Action,
	{_State72, LbracketToken}:                      _GotoState42Action,
	{_State72, ParseErrorToken}:                    _GotoState46Action,
	{_State72, OptionalLabelDeclType}:              _GotoState153Action,
	{_State72, CallExprType}:                       _GotoState203Action,
	{_State72, AtomExprType}:                       _GotoState63Action,
	{_State72, ParseErrorExprType}:                 _GotoState93Action,
	{_State72, LiteralExprType}:                    _GotoState87Action,
	{_State72, NamedExprType}:                      _GotoState90Action,
	{_State72, BlockExprType}:                      _GotoState70Action,
	{_State72, InitializeExprType}:                 _GotoState84Action,
	{_State72, ImplicitStructExprType}:             _GotoState80Action,
	{_State72, AccessibleExprType}:                 _GotoState202Action,
	{_State72, AccessExprType}:                     _GotoState55Action,
	{_State72, IndexExprType}:                      _GotoState82Action,
	{_State72, InitializableTypeExprType}:          _GotoState83Action,
	{_State72, SliceTypeExprType}:                  _GotoState101Action,
	{_State72, ArrayTypeExprType}:                  _GotoState60Action,
	{_State72, MapTypeExprType}:                    _GotoState88Action,
	{_State72, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State72, AnonymousFuncExprType}:              _GotoState59Action,
	{_State74, EqualToken}:                         _GotoState204Action,
	{_State74, NotEqualToken}:                      _GotoState209Action,
	{_State74, LessToken}:                          _GotoState207Action,
	{_State74, LessOrEqualToken}:                   _GotoState208Action,
	{_State74, GreaterToken}:                       _GotoState205Action,
	{_State74, GreaterOrEqualToken}:                _GotoState206Action,
	{_State74, CmpOpType}:                          _GotoState210Action,
	{_State78, CommaToken}:                         _GotoState211Action,
	{_State83, LparenToken}:                        _GotoState212Action,
	{_State86, IntegerLiteralToken}:                _GotoState40Action,
	{_State86, FloatLiteralToken}:                  _GotoState35Action,
	{_State86, RuneLiteralToken}:                   _GotoState48Action,
	{_State86, StringLiteralToken}:                 _GotoState49Action,
	{_State86, IdentifierToken}:                    _GotoState38Action,
	{_State86, UnderscoreToken}:                    _GotoState53Action,
	{_State86, TrueToken}:                          _GotoState52Action,
	{_State86, FalseToken}:                         _GotoState34Action,
	{_State86, StructToken}:                        _GotoState50Action,
	{_State86, FuncToken}:                          _GotoState36Action,
	{_State86, VarToken}:                           _GotoState15Action,
	{_State86, LetToken}:                           _GotoState12Action,
	{_State86, NotToken}:                           _GotoState45Action,
	{_State86, LabelDeclToken}:                     _GotoState41Action,
	{_State86, JumpLabelToken}:                     _GotoState213Action,
	{_State86, LparenToken}:                        _GotoState43Action,
	{_State86, LbracketToken}:                      _GotoState42Action,
	{_State86, SubToken}:                           _GotoState51Action,
	{_State86, MulToken}:                           _GotoState44Action,
	{_State86, BitNegToken}:                        _GotoState27Action,
	{_State86, BitAndToken}:                        _GotoState26Action,
	{_State86, GreaterToken}:                       _GotoState37Action,
	{_State86, ParseErrorToken}:                    _GotoState46Action,
	{_State86, ExprsType}:                          _GotoState214Action,
	{_State86, VarDeclPatternType}:                 _GotoState104Action,
	{_State86, VarOrLetType}:                       _GotoState24Action,
	{_State86, ExprType}:                           _GotoState76Action,
	{_State86, OptionalLabelDeclType}:              _GotoState91Action,
	{_State86, SequenceExprType}:                   _GotoState106Action,
	{_State86, CallExprType}:                       _GotoState71Action,
	{_State86, AtomExprType}:                       _GotoState63Action,
	{_State86, ParseErrorExprType}:                 _GotoState93Action,
	{_State86, LiteralExprType}:                    _GotoState87Action,
	{_State86, NamedExprType}:                      _GotoState90Action,
	{_State86, BlockExprType}:                      _GotoState70Action,
	{_State86, InitializeExprType}:                 _GotoState84Action,
	{_State86, ImplicitStructExprType}:             _GotoState80Action,
	{_State86, AccessibleExprType}:                 _GotoState105Action,
	{_State86, AccessExprType}:                     _GotoState55Action,
	{_State86, IndexExprType}:                      _GotoState82Action,
	{_State86, PostfixableExprType}:                _GotoState95Action,
	{_State86, PostfixUnaryExprType}:               _GotoState94Action,
	{_State86, PrefixableExprType}:                 _GotoState98Action,
	{_State86, PrefixUnaryExprType}:                _GotoState96Action,
	{_State86, PrefixUnaryOpType}:                  _GotoState97Action,
	{_State86, MulExprType}:                        _GotoState89Action,
	{_State86, BinaryMulExprType}:                  _GotoState67Action,
	{_State86, AddExprType}:                        _GotoState57Action,
	{_State86, BinaryAddExprType}:                  _GotoState64Action,
	{_State86, CmpExprType}:                        _GotoState74Action,
	{_State86, BinaryCmpExprType}:                  _GotoState66Action,
	{_State86, AndExprType}:                        _GotoState58Action,
	{_State86, BinaryAndExprType}:                  _GotoState65Action,
	{_State86, OrExprType}:                         _GotoState92Action,
	{_State86, BinaryOrExprType}:                   _GotoState69Action,
	{_State86, InitializableTypeExprType}:          _GotoState83Action,
	{_State86, SliceTypeExprType}:                  _GotoState101Action,
	{_State86, ArrayTypeExprType}:                  _GotoState60Action,
	{_State86, MapTypeExprType}:                    _GotoState88Action,
	{_State86, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State86, AnonymousFuncExprType}:              _GotoState59Action,
	{_State89, MulToken}:                           _GotoState220Action,
	{_State89, DivToken}:                           _GotoState218Action,
	{_State89, ModToken}:                           _GotoState219Action,
	{_State89, BitAndToken}:                        _GotoState215Action,
	{_State89, BitLshiftToken}:                     _GotoState216Action,
	{_State89, BitRshiftToken}:                     _GotoState217Action,
	{_State89, MulOpType}:                          _GotoState221Action,
	{_State91, IfToken}:                            _GotoState224Action,
	{_State91, SwitchToken}:                        _GotoState225Action,
	{_State91, ForToken}:                           _GotoState223Action,
	{_State91, DoToken}:                            _GotoState222Action,
	{_State91, LbraceToken}:                        _GotoState11Action,
	{_State91, StatementBlockType}:                 _GotoState228Action,
	{_State91, IfExprType}:                         _GotoState226Action,
	{_State91, SwitchExprType}:                     _GotoState229Action,
	{_State91, LoopExprType}:                       _GotoState227Action,
	{_State92, OrToken}:                            _GotoState230Action,
	{_State97, IntegerLiteralToken}:                _GotoState40Action,
	{_State97, FloatLiteralToken}:                  _GotoState35Action,
	{_State97, RuneLiteralToken}:                   _GotoState48Action,
	{_State97, StringLiteralToken}:                 _GotoState49Action,
	{_State97, IdentifierToken}:                    _GotoState38Action,
	{_State97, UnderscoreToken}:                    _GotoState53Action,
	{_State97, TrueToken}:                          _GotoState52Action,
	{_State97, FalseToken}:                         _GotoState34Action,
	{_State97, StructToken}:                        _GotoState50Action,
	{_State97, FuncToken}:                          _GotoState36Action,
	{_State97, NotToken}:                           _GotoState45Action,
	{_State97, LabelDeclToken}:                     _GotoState41Action,
	{_State97, LparenToken}:                        _GotoState43Action,
	{_State97, LbracketToken}:                      _GotoState42Action,
	{_State97, SubToken}:                           _GotoState51Action,
	{_State97, MulToken}:                           _GotoState44Action,
	{_State97, BitNegToken}:                        _GotoState27Action,
	{_State97, BitAndToken}:                        _GotoState26Action,
	{_State97, ParseErrorToken}:                    _GotoState46Action,
	{_State97, OptionalLabelDeclType}:              _GotoState153Action,
	{_State97, CallExprType}:                       _GotoState71Action,
	{_State97, AtomExprType}:                       _GotoState63Action,
	{_State97, ParseErrorExprType}:                 _GotoState93Action,
	{_State97, LiteralExprType}:                    _GotoState87Action,
	{_State97, NamedExprType}:                      _GotoState90Action,
	{_State97, BlockExprType}:                      _GotoState70Action,
	{_State97, InitializeExprType}:                 _GotoState84Action,
	{_State97, ImplicitStructExprType}:             _GotoState80Action,
	{_State97, AccessibleExprType}:                 _GotoState105Action,
	{_State97, AccessExprType}:                     _GotoState55Action,
	{_State97, IndexExprType}:                      _GotoState82Action,
	{_State97, PostfixableExprType}:                _GotoState95Action,
	{_State97, PostfixUnaryExprType}:               _GotoState94Action,
	{_State97, PrefixableExprType}:                 _GotoState231Action,
	{_State97, PrefixUnaryExprType}:                _GotoState96Action,
	{_State97, PrefixUnaryOpType}:                  _GotoState97Action,
	{_State97, InitializableTypeExprType}:          _GotoState83Action,
	{_State97, SliceTypeExprType}:                  _GotoState101Action,
	{_State97, ArrayTypeExprType}:                  _GotoState60Action,
	{_State97, MapTypeExprType}:                    _GotoState88Action,
	{_State97, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State97, AnonymousFuncExprType}:              _GotoState59Action,
	{_State105, LbracketToken}:                     _GotoState186Action,
	{_State105, DotToken}:                          _GotoState185Action,
	{_State105, QuestionToken}:                     _GotoState189Action,
	{_State105, DollarLbracketToken}:               _GotoState184Action,
	{_State105, GenericTypeArgumentsType}:          _GotoState193Action,
	{_State110, LparenToken}:                       _GotoState232Action,
	{_State112, LparenToken}:                       _GotoState233Action,
	{_State113, DotToken}:                          _GotoState234Action,
	{_State113, DollarLbracketToken}:               _GotoState184Action,
	{_State113, GenericTypeArgumentsType}:          _GotoState235Action,
	{_State114, IdentifierToken}:                   _GotoState237Action,
	{_State114, UnderscoreToken}:                   _GotoState238Action,
	{_State114, UnsafeToken}:                       _GotoState54Action,
	{_State114, StructToken}:                       _GotoState50Action,
	{_State114, EnumToken}:                         _GotoState110Action,
	{_State114, TraitToken}:                        _GotoState118Action,
	{_State114, FuncToken}:                         _GotoState236Action,
	{_State114, LparenToken}:                       _GotoState114Action,
	{_State114, LbracketToken}:                     _GotoState42Action,
	{_State114, DotToken}:                          _GotoState109Action,
	{_State114, QuestionToken}:                     _GotoState116Action,
	{_State114, ExclaimToken}:                      _GotoState111Action,
	{_State114, TildeTildeToken}:                   _GotoState117Action,
	{_State114, BitNegToken}:                       _GotoState108Action,
	{_State114, BitAndToken}:                       _GotoState107Action,
	{_State114, ParseErrorToken}:                   _GotoState115Action,
	{_State114, UnsafeStatementType}:               _GotoState246Action,
	{_State114, InitializableTypeExprType}:         _GotoState127Action,
	{_State114, SliceTypeExprType}:                 _GotoState101Action,
	{_State114, ArrayTypeExprType}:                 _GotoState60Action,
	{_State114, MapTypeExprType}:                   _GotoState88Action,
	{_State114, AtomTypeExprType}:                  _GotoState120Action,
	{_State114, NamedTypeExprType}:                 _GotoState128Action,
	{_State114, InferredTypeExprType}:              _GotoState126Action,
	{_State114, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State114, ReturnableTypeExprType}:            _GotoState132Action,
	{_State114, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State114, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State114, TypeExprType}:                      _GotoState245Action,
	{_State114, BinaryTypeExprType}:                _GotoState121Action,
	{_State114, FieldDefType}:                      _GotoState239Action,
	{_State114, ProperImplicitFieldDefsType}:       _GotoState244Action,
	{_State114, ImplicitFieldDefsType}:             _GotoState241Action,
	{_State114, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State114, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State114, TraitTypeExprType}:                 _GotoState133Action,
	{_State114, ProperImplicitEnumValueDefsType}:   _GotoState243Action,
	{_State114, ImplicitEnumValueDefsType}:         _GotoState240Action,
	{_State114, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State114, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State114, FuncTypeExprType}:                  _GotoState123Action,
	{_State114, MethodSignatureType}:               _GotoState242Action,
	{_State118, LparenToken}:                       _GotoState247Action,
	{_State131, IdentifierToken}:                   _GotoState113Action,
	{_State131, UnderscoreToken}:                   _GotoState119Action,
	{_State131, StructToken}:                       _GotoState50Action,
	{_State131, EnumToken}:                         _GotoState110Action,
	{_State131, TraitToken}:                        _GotoState118Action,
	{_State131, FuncToken}:                         _GotoState112Action,
	{_State131, LparenToken}:                       _GotoState114Action,
	{_State131, LbracketToken}:                     _GotoState42Action,
	{_State131, DotToken}:                          _GotoState109Action,
	{_State131, QuestionToken}:                     _GotoState116Action,
	{_State131, ExclaimToken}:                      _GotoState111Action,
	{_State131, TildeTildeToken}:                   _GotoState117Action,
	{_State131, BitNegToken}:                       _GotoState108Action,
	{_State131, BitAndToken}:                       _GotoState107Action,
	{_State131, ParseErrorToken}:                   _GotoState115Action,
	{_State131, InitializableTypeExprType}:         _GotoState127Action,
	{_State131, SliceTypeExprType}:                 _GotoState101Action,
	{_State131, ArrayTypeExprType}:                 _GotoState60Action,
	{_State131, MapTypeExprType}:                   _GotoState88Action,
	{_State131, AtomTypeExprType}:                  _GotoState120Action,
	{_State131, NamedTypeExprType}:                 _GotoState128Action,
	{_State131, InferredTypeExprType}:              _GotoState126Action,
	{_State131, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State131, ReturnableTypeExprType}:            _GotoState248Action,
	{_State131, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State131, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State131, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State131, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State131, TraitTypeExprType}:                 _GotoState133Action,
	{_State131, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State131, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State131, FuncTypeExprType}:                  _GotoState123Action,
	{_State134, DollarLbracketToken}:               _GotoState254Action,
	{_State134, AssignToken}:                       _GotoState253Action,
	{_State134, GenericParameterDefsType}:          _GotoState255Action,
	{_State135, IdentifierToken}:                   _GotoState256Action,
	{_State135, UnderscoreToken}:                   _GotoState257Action,
	{_State135, ProperParameterDefType}:            _GotoState259Action,
	{_State135, ParameterDefType}:                  _GotoState258Action,
	{_State136, NewlinesToken}:                     _GotoState260Action,
	{_State136, SemicolonToken}:                    _GotoState261Action,
	{_State138, RbraceToken}:                       _GotoState262Action,
	{_State140, DollarLbracketToken}:               _GotoState254Action,
	{_State140, AssignToken}:                       _GotoState263Action,
	{_State140, GenericParameterDefsType}:          _GotoState264Action,
	{_State141, CommentGroupsToken}:                _GotoState9Action,
	{_State141, PackageToken}:                      _GotoState13Action,
	{_State141, TypeToken}:                         _GotoState14Action,
	{_State141, FuncToken}:                         _GotoState10Action,
	{_State141, VarToken}:                          _GotoState15Action,
	{_State141, LetToken}:                          _GotoState12Action,
	{_State141, LbraceToken}:                       _GotoState11Action,
	{_State141, DefinitionType}:                    _GotoState265Action,
	{_State141, StatementBlockType}:                _GotoState21Action,
	{_State141, VarDeclPatternType}:                _GotoState23Action,
	{_State141, VarOrLetType}:                      _GotoState24Action,
	{_State141, TypeDefType}:                       _GotoState22Action,
	{_State141, NamedFuncDefType}:                  _GotoState18Action,
	{_State141, PackageDefType}:                    _GotoState19Action,
	{_State142, IntegerLiteralToken}:               _GotoState40Action,
	{_State142, FloatLiteralToken}:                 _GotoState35Action,
	{_State142, RuneLiteralToken}:                  _GotoState48Action,
	{_State142, StringLiteralToken}:                _GotoState49Action,
	{_State142, IdentifierToken}:                   _GotoState38Action,
	{_State142, UnderscoreToken}:                   _GotoState53Action,
	{_State142, TrueToken}:                         _GotoState52Action,
	{_State142, FalseToken}:                        _GotoState34Action,
	{_State142, StructToken}:                       _GotoState50Action,
	{_State142, FuncToken}:                         _GotoState36Action,
	{_State142, VarToken}:                          _GotoState15Action,
	{_State142, LetToken}:                          _GotoState12Action,
	{_State142, NotToken}:                          _GotoState45Action,
	{_State142, LabelDeclToken}:                    _GotoState41Action,
	{_State142, LparenToken}:                       _GotoState43Action,
	{_State142, LbracketToken}:                     _GotoState42Action,
	{_State142, SubToken}:                          _GotoState51Action,
	{_State142, MulToken}:                          _GotoState44Action,
	{_State142, BitNegToken}:                       _GotoState27Action,
	{_State142, BitAndToken}:                       _GotoState26Action,
	{_State142, GreaterToken}:                      _GotoState37Action,
	{_State142, ParseErrorToken}:                   _GotoState46Action,
	{_State142, VarDeclPatternType}:                _GotoState104Action,
	{_State142, VarOrLetType}:                      _GotoState24Action,
	{_State142, ExprType}:                          _GotoState266Action,
	{_State142, OptionalLabelDeclType}:             _GotoState91Action,
	{_State142, SequenceExprType}:                  _GotoState106Action,
	{_State142, CallExprType}:                      _GotoState71Action,
	{_State142, AtomExprType}:                      _GotoState63Action,
	{_State142, ParseErrorExprType}:                _GotoState93Action,
	{_State142, LiteralExprType}:                   _GotoState87Action,
	{_State142, NamedExprType}:                     _GotoState90Action,
	{_State142, BlockExprType}:                     _GotoState70Action,
	{_State142, InitializeExprType}:                _GotoState84Action,
	{_State142, ImplicitStructExprType}:            _GotoState80Action,
	{_State142, AccessibleExprType}:                _GotoState105Action,
	{_State142, AccessExprType}:                    _GotoState55Action,
	{_State142, IndexExprType}:                     _GotoState82Action,
	{_State142, PostfixableExprType}:               _GotoState95Action,
	{_State142, PostfixUnaryExprType}:              _GotoState94Action,
	{_State142, PrefixableExprType}:                _GotoState98Action,
	{_State142, PrefixUnaryExprType}:               _GotoState96Action,
	{_State142, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State142, MulExprType}:                       _GotoState89Action,
	{_State142, BinaryMulExprType}:                 _GotoState67Action,
	{_State142, AddExprType}:                       _GotoState57Action,
	{_State142, BinaryAddExprType}:                 _GotoState64Action,
	{_State142, CmpExprType}:                       _GotoState74Action,
	{_State142, BinaryCmpExprType}:                 _GotoState66Action,
	{_State142, AndExprType}:                       _GotoState58Action,
	{_State142, BinaryAndExprType}:                 _GotoState65Action,
	{_State142, OrExprType}:                        _GotoState92Action,
	{_State142, BinaryOrExprType}:                  _GotoState69Action,
	{_State142, InitializableTypeExprType}:         _GotoState83Action,
	{_State142, SliceTypeExprType}:                 _GotoState101Action,
	{_State142, ArrayTypeExprType}:                 _GotoState60Action,
	{_State142, MapTypeExprType}:                   _GotoState88Action,
	{_State142, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State142, AnonymousFuncExprType}:             _GotoState59Action,
	{_State144, IdentifierToken}:                   _GotoState268Action,
	{_State144, UnderscoreToken}:                   _GotoState145Action,
	{_State144, LparenToken}:                       _GotoState144Action,
	{_State144, EllipsisToken}:                     _GotoState267Action,
	{_State144, VarPatternType}:                    _GotoState271Action,
	{_State144, TuplePatternType}:                  _GotoState146Action,
	{_State144, FieldVarPatternsType}:              _GotoState270Action,
	{_State144, FieldVarPatternType}:               _GotoState269Action,
	{_State147, IdentifierToken}:                   _GotoState113Action,
	{_State147, UnderscoreToken}:                   _GotoState119Action,
	{_State147, StructToken}:                       _GotoState50Action,
	{_State147, EnumToken}:                         _GotoState110Action,
	{_State147, TraitToken}:                        _GotoState118Action,
	{_State147, FuncToken}:                         _GotoState112Action,
	{_State147, LparenToken}:                       _GotoState114Action,
	{_State147, LbracketToken}:                     _GotoState42Action,
	{_State147, DotToken}:                          _GotoState109Action,
	{_State147, QuestionToken}:                     _GotoState116Action,
	{_State147, ExclaimToken}:                      _GotoState111Action,
	{_State147, TildeTildeToken}:                   _GotoState117Action,
	{_State147, BitNegToken}:                       _GotoState108Action,
	{_State147, BitAndToken}:                       _GotoState107Action,
	{_State147, ParseErrorToken}:                   _GotoState115Action,
	{_State147, OptionalTypeExprType}:              _GotoState272Action,
	{_State147, InitializableTypeExprType}:         _GotoState127Action,
	{_State147, SliceTypeExprType}:                 _GotoState101Action,
	{_State147, ArrayTypeExprType}:                 _GotoState60Action,
	{_State147, MapTypeExprType}:                   _GotoState88Action,
	{_State147, AtomTypeExprType}:                  _GotoState120Action,
	{_State147, NamedTypeExprType}:                 _GotoState128Action,
	{_State147, InferredTypeExprType}:              _GotoState126Action,
	{_State147, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State147, ReturnableTypeExprType}:            _GotoState132Action,
	{_State147, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State147, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State147, TypeExprType}:                      _GotoState273Action,
	{_State147, BinaryTypeExprType}:                _GotoState121Action,
	{_State147, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State147, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State147, TraitTypeExprType}:                 _GotoState133Action,
	{_State147, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State147, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State147, FuncTypeExprType}:                  _GotoState123Action,
	{_State148, IdentifierToken}:                   _GotoState274Action,
	{_State149, DotToken}:                          _GotoState275Action,
	{_State152, CommaToken}:                        _GotoState277Action,
	{_State152, ColonToken}:                        _GotoState276Action,
	{_State153, LbraceToken}:                       _GotoState11Action,
	{_State153, StatementBlockType}:                _GotoState228Action,
	{_State155, IntegerLiteralToken}:               _GotoState40Action,
	{_State155, FloatLiteralToken}:                 _GotoState35Action,
	{_State155, RuneLiteralToken}:                  _GotoState48Action,
	{_State155, StringLiteralToken}:                _GotoState49Action,
	{_State155, IdentifierToken}:                   _GotoState38Action,
	{_State155, UnderscoreToken}:                   _GotoState53Action,
	{_State155, TrueToken}:                         _GotoState52Action,
	{_State155, FalseToken}:                        _GotoState34Action,
	{_State155, ReturnToken}:                       _GotoState47Action,
	{_State155, BreakToken}:                        _GotoState28Action,
	{_State155, ContinueToken}:                     _GotoState30Action,
	{_State155, FallthroughToken}:                  _GotoState33Action,
	{_State155, UnsafeToken}:                       _GotoState54Action,
	{_State155, StructToken}:                       _GotoState50Action,
	{_State155, FuncToken}:                         _GotoState36Action,
	{_State155, AsyncToken}:                        _GotoState25Action,
	{_State155, DeferToken}:                        _GotoState32Action,
	{_State155, VarToken}:                          _GotoState15Action,
	{_State155, LetToken}:                          _GotoState12Action,
	{_State155, NotToken}:                          _GotoState45Action,
	{_State155, LabelDeclToken}:                    _GotoState41Action,
	{_State155, LparenToken}:                       _GotoState43Action,
	{_State155, LbracketToken}:                     _GotoState42Action,
	{_State155, SubToken}:                          _GotoState51Action,
	{_State155, MulToken}:                          _GotoState44Action,
	{_State155, BitNegToken}:                       _GotoState27Action,
	{_State155, BitAndToken}:                       _GotoState26Action,
	{_State155, GreaterToken}:                      _GotoState37Action,
	{_State155, ParseErrorToken}:                   _GotoState46Action,
	{_State155, SimpleStatementType}:               _GotoState279Action,
	{_State155, OptionalSimpleStatementType}:       _GotoState278Action,
	{_State155, ExprOrImproperStructStatementType}: _GotoState77Action,
	{_State155, ExprsType}:                         _GotoState78Action,
	{_State155, CallbackOpType}:                    _GotoState72Action,
	{_State155, CallbackOpStatementType}:           _GotoState73Action,
	{_State155, UnsafeStatementType}:               _GotoState103Action,
	{_State155, JumpStatementType}:                 _GotoState85Action,
	{_State155, JumpTypeType}:                      _GotoState86Action,
	{_State155, FallthroughStatementType}:          _GotoState79Action,
	{_State155, AssignStatementType}:               _GotoState62Action,
	{_State155, UnaryOpAssignStatementType}:        _GotoState102Action,
	{_State155, BinaryOpAssignStatementType}:       _GotoState68Action,
	{_State155, VarDeclPatternType}:                _GotoState104Action,
	{_State155, VarOrLetType}:                      _GotoState24Action,
	{_State155, AssignPatternType}:                 _GotoState61Action,
	{_State155, ExprType}:                          _GotoState76Action,
	{_State155, OptionalLabelDeclType}:             _GotoState91Action,
	{_State155, SequenceExprType}:                  _GotoState99Action,
	{_State155, CallExprType}:                      _GotoState71Action,
	{_State155, AtomExprType}:                      _GotoState63Action,
	{_State155, ParseErrorExprType}:                _GotoState93Action,
	{_State155, LiteralExprType}:                   _GotoState87Action,
	{_State155, NamedExprType}:                     _GotoState90Action,
	{_State155, BlockExprType}:                     _GotoState70Action,
	{_State155, InitializeExprType}:                _GotoState84Action,
	{_State155, ImplicitStructExprType}:            _GotoState80Action,
	{_State155, AccessibleExprType}:                _GotoState56Action,
	{_State155, AccessExprType}:                    _GotoState55Action,
	{_State155, IndexExprType}:                     _GotoState82Action,
	{_State155, PostfixableExprType}:               _GotoState95Action,
	{_State155, PostfixUnaryExprType}:              _GotoState94Action,
	{_State155, PrefixableExprType}:                _GotoState98Action,
	{_State155, PrefixUnaryExprType}:               _GotoState96Action,
	{_State155, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State155, MulExprType}:                       _GotoState89Action,
	{_State155, BinaryMulExprType}:                 _GotoState67Action,
	{_State155, AddExprType}:                       _GotoState57Action,
	{_State155, BinaryAddExprType}:                 _GotoState64Action,
	{_State155, CmpExprType}:                       _GotoState74Action,
	{_State155, BinaryCmpExprType}:                 _GotoState66Action,
	{_State155, AndExprType}:                       _GotoState58Action,
	{_State155, BinaryAndExprType}:                 _GotoState65Action,
	{_State155, OrExprType}:                        _GotoState92Action,
	{_State155, BinaryOrExprType}:                  _GotoState69Action,
	{_State155, InitializableTypeExprType}:         _GotoState83Action,
	{_State155, SliceTypeExprType}:                 _GotoState101Action,
	{_State155, ArrayTypeExprType}:                 _GotoState60Action,
	{_State155, MapTypeExprType}:                   _GotoState88Action,
	{_State155, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State155, AnonymousFuncExprType}:             _GotoState59Action,
	{_State156, IdentifierToken}:                   _GotoState256Action,
	{_State156, UnderscoreToken}:                   _GotoState257Action,
	{_State156, ProperParameterDefType}:            _GotoState259Action,
	{_State156, ParameterDefType}:                  _GotoState280Action,
	{_State156, ProperParameterDefsType}:           _GotoState282Action,
	{_State156, ParameterDefsType}:                 _GotoState281Action,
	{_State158, StringLiteralToken}:                _GotoState283Action,
	{_State159, StringLiteralToken}:                _GotoState284Action,
	{_State160, StringLiteralToken}:                _GotoState161Action,
	{_State160, IdentifierToken}:                   _GotoState159Action,
	{_State160, UnderscoreToken}:                   _GotoState162Action,
	{_State160, DotToken}:                          _GotoState158Action,
	{_State160, ImportClauseType}:                  _GotoState285Action,
	{_State160, ProperImportClausesType}:           _GotoState287Action,
	{_State160, ImportClausesType}:                 _GotoState286Action,
	{_State162, StringLiteralToken}:                _GotoState288Action,
	{_State164, RbracketToken}:                     _GotoState291Action,
	{_State164, CommaToken}:                        _GotoState290Action,
	{_State164, ColonToken}:                        _GotoState289Action,
	{_State164, AddToken}:                          _GotoState249Action,
	{_State164, SubToken}:                          _GotoState251Action,
	{_State164, MulToken}:                          _GotoState250Action,
	{_State164, BinaryTypeOpType}:                  _GotoState252Action,
	{_State165, IntegerLiteralToken}:               _GotoState40Action,
	{_State165, FloatLiteralToken}:                 _GotoState35Action,
	{_State165, RuneLiteralToken}:                  _GotoState48Action,
	{_State165, StringLiteralToken}:                _GotoState49Action,
	{_State165, IdentifierToken}:                   _GotoState38Action,
	{_State165, UnderscoreToken}:                   _GotoState53Action,
	{_State165, TrueToken}:                         _GotoState52Action,
	{_State165, FalseToken}:                        _GotoState34Action,
	{_State165, StructToken}:                       _GotoState50Action,
	{_State165, FuncToken}:                         _GotoState36Action,
	{_State165, VarToken}:                          _GotoState15Action,
	{_State165, LetToken}:                          _GotoState12Action,
	{_State165, NotToken}:                          _GotoState45Action,
	{_State165, LabelDeclToken}:                    _GotoState41Action,
	{_State165, LparenToken}:                       _GotoState43Action,
	{_State165, LbracketToken}:                     _GotoState42Action,
	{_State165, SubToken}:                          _GotoState51Action,
	{_State165, MulToken}:                          _GotoState44Action,
	{_State165, BitNegToken}:                       _GotoState27Action,
	{_State165, BitAndToken}:                       _GotoState26Action,
	{_State165, GreaterToken}:                      _GotoState37Action,
	{_State165, ParseErrorToken}:                   _GotoState46Action,
	{_State165, VarDeclPatternType}:                _GotoState104Action,
	{_State165, VarOrLetType}:                      _GotoState24Action,
	{_State165, ExprType}:                          _GotoState292Action,
	{_State165, OptionalLabelDeclType}:             _GotoState91Action,
	{_State165, SequenceExprType}:                  _GotoState106Action,
	{_State165, CallExprType}:                      _GotoState71Action,
	{_State165, AtomExprType}:                      _GotoState63Action,
	{_State165, ParseErrorExprType}:                _GotoState93Action,
	{_State165, LiteralExprType}:                   _GotoState87Action,
	{_State165, NamedExprType}:                     _GotoState90Action,
	{_State165, BlockExprType}:                     _GotoState70Action,
	{_State165, InitializeExprType}:                _GotoState84Action,
	{_State165, ImplicitStructExprType}:            _GotoState80Action,
	{_State165, AccessibleExprType}:                _GotoState105Action,
	{_State165, AccessExprType}:                    _GotoState55Action,
	{_State165, IndexExprType}:                     _GotoState82Action,
	{_State165, PostfixableExprType}:               _GotoState95Action,
	{_State165, PostfixUnaryExprType}:              _GotoState94Action,
	{_State165, PrefixableExprType}:                _GotoState98Action,
	{_State165, PrefixUnaryExprType}:               _GotoState96Action,
	{_State165, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State165, MulExprType}:                       _GotoState89Action,
	{_State165, BinaryMulExprType}:                 _GotoState67Action,
	{_State165, AddExprType}:                       _GotoState57Action,
	{_State165, BinaryAddExprType}:                 _GotoState64Action,
	{_State165, CmpExprType}:                       _GotoState74Action,
	{_State165, BinaryCmpExprType}:                 _GotoState66Action,
	{_State165, AndExprType}:                       _GotoState58Action,
	{_State165, BinaryAndExprType}:                 _GotoState65Action,
	{_State165, OrExprType}:                        _GotoState92Action,
	{_State165, BinaryOrExprType}:                  _GotoState69Action,
	{_State165, InitializableTypeExprType}:         _GotoState83Action,
	{_State165, SliceTypeExprType}:                 _GotoState101Action,
	{_State165, ArrayTypeExprType}:                 _GotoState60Action,
	{_State165, MapTypeExprType}:                   _GotoState88Action,
	{_State165, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State165, AnonymousFuncExprType}:             _GotoState59Action,
	{_State167, AssignToken}:                       _GotoState293Action,
	{_State169, RparenToken}:                       _GotoState294Action,
	{_State170, ColonToken}:                        _GotoState295Action,
	{_State171, ColonToken}:                        _GotoState296Action,
	{_State171, EllipsisToken}:                     _GotoState297Action,
	{_State172, CommaToken}:                        _GotoState298Action,
	{_State173, IdentifierToken}:                   _GotoState237Action,
	{_State173, UnderscoreToken}:                   _GotoState238Action,
	{_State173, UnsafeToken}:                       _GotoState54Action,
	{_State173, StructToken}:                       _GotoState50Action,
	{_State173, EnumToken}:                         _GotoState110Action,
	{_State173, TraitToken}:                        _GotoState118Action,
	{_State173, FuncToken}:                         _GotoState236Action,
	{_State173, LparenToken}:                       _GotoState114Action,
	{_State173, LbracketToken}:                     _GotoState42Action,
	{_State173, DotToken}:                          _GotoState109Action,
	{_State173, QuestionToken}:                     _GotoState116Action,
	{_State173, ExclaimToken}:                      _GotoState111Action,
	{_State173, TildeTildeToken}:                   _GotoState117Action,
	{_State173, BitNegToken}:                       _GotoState108Action,
	{_State173, BitAndToken}:                       _GotoState107Action,
	{_State173, ParseErrorToken}:                   _GotoState115Action,
	{_State173, UnsafeStatementType}:               _GotoState246Action,
	{_State173, InitializableTypeExprType}:         _GotoState127Action,
	{_State173, SliceTypeExprType}:                 _GotoState101Action,
	{_State173, ArrayTypeExprType}:                 _GotoState60Action,
	{_State173, MapTypeExprType}:                   _GotoState88Action,
	{_State173, AtomTypeExprType}:                  _GotoState120Action,
	{_State173, NamedTypeExprType}:                 _GotoState128Action,
	{_State173, InferredTypeExprType}:              _GotoState126Action,
	{_State173, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State173, ReturnableTypeExprType}:            _GotoState132Action,
	{_State173, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State173, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State173, TypeExprType}:                      _GotoState245Action,
	{_State173, BinaryTypeExprType}:                _GotoState121Action,
	{_State173, FieldDefType}:                      _GotoState300Action,
	{_State173, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State173, ProperExplicitFieldDefsType}:       _GotoState301Action,
	{_State173, ExplicitFieldDefsType}:             _GotoState299Action,
	{_State173, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State173, TraitTypeExprType}:                 _GotoState133Action,
	{_State173, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State173, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State173, FuncTypeExprType}:                  _GotoState123Action,
	{_State173, MethodSignatureType}:               _GotoState242Action,
	{_State174, IdentifierToken}:                   _GotoState302Action,
	{_State184, IdentifierToken}:                   _GotoState113Action,
	{_State184, UnderscoreToken}:                   _GotoState119Action,
	{_State184, StructToken}:                       _GotoState50Action,
	{_State184, EnumToken}:                         _GotoState110Action,
	{_State184, TraitToken}:                        _GotoState118Action,
	{_State184, FuncToken}:                         _GotoState112Action,
	{_State184, LparenToken}:                       _GotoState114Action,
	{_State184, LbracketToken}:                     _GotoState42Action,
	{_State184, DotToken}:                          _GotoState109Action,
	{_State184, QuestionToken}:                     _GotoState116Action,
	{_State184, ExclaimToken}:                      _GotoState111Action,
	{_State184, TildeTildeToken}:                   _GotoState117Action,
	{_State184, BitNegToken}:                       _GotoState108Action,
	{_State184, BitAndToken}:                       _GotoState107Action,
	{_State184, ParseErrorToken}:                   _GotoState115Action,
	{_State184, InitializableTypeExprType}:         _GotoState127Action,
	{_State184, SliceTypeExprType}:                 _GotoState101Action,
	{_State184, ArrayTypeExprType}:                 _GotoState60Action,
	{_State184, MapTypeExprType}:                   _GotoState88Action,
	{_State184, AtomTypeExprType}:                  _GotoState120Action,
	{_State184, NamedTypeExprType}:                 _GotoState128Action,
	{_State184, InferredTypeExprType}:              _GotoState126Action,
	{_State184, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State184, ReturnableTypeExprType}:            _GotoState132Action,
	{_State184, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State184, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State184, TypeExprType}:                      _GotoState305Action,
	{_State184, BinaryTypeExprType}:                _GotoState121Action,
	{_State184, ProperTypeArgumentsType}:           _GotoState303Action,
	{_State184, TypeArgumentsType}:                 _GotoState304Action,
	{_State184, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State184, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State184, TraitTypeExprType}:                 _GotoState133Action,
	{_State184, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State184, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State184, FuncTypeExprType}:                  _GotoState123Action,
	{_State185, IdentifierToken}:                   _GotoState306Action,
	{_State186, IntegerLiteralToken}:               _GotoState40Action,
	{_State186, FloatLiteralToken}:                 _GotoState35Action,
	{_State186, RuneLiteralToken}:                  _GotoState48Action,
	{_State186, StringLiteralToken}:                _GotoState49Action,
	{_State186, IdentifierToken}:                   _GotoState167Action,
	{_State186, UnderscoreToken}:                   _GotoState53Action,
	{_State186, TrueToken}:                         _GotoState52Action,
	{_State186, FalseToken}:                        _GotoState34Action,
	{_State186, StructToken}:                       _GotoState50Action,
	{_State186, FuncToken}:                         _GotoState36Action,
	{_State186, VarToken}:                          _GotoState15Action,
	{_State186, LetToken}:                          _GotoState12Action,
	{_State186, NotToken}:                          _GotoState45Action,
	{_State186, LabelDeclToken}:                    _GotoState41Action,
	{_State186, LparenToken}:                       _GotoState43Action,
	{_State186, LbracketToken}:                     _GotoState42Action,
	{_State186, ColonToken}:                        _GotoState165Action,
	{_State186, EllipsisToken}:                     _GotoState166Action,
	{_State186, SubToken}:                          _GotoState51Action,
	{_State186, MulToken}:                          _GotoState44Action,
	{_State186, BitNegToken}:                       _GotoState27Action,
	{_State186, BitAndToken}:                       _GotoState26Action,
	{_State186, GreaterToken}:                      _GotoState37Action,
	{_State186, ParseErrorToken}:                   _GotoState46Action,
	{_State186, VarDeclPatternType}:                _GotoState104Action,
	{_State186, VarOrLetType}:                      _GotoState24Action,
	{_State186, ExprType}:                          _GotoState171Action,
	{_State186, OptionalLabelDeclType}:             _GotoState91Action,
	{_State186, SequenceExprType}:                  _GotoState106Action,
	{_State186, CallExprType}:                      _GotoState71Action,
	{_State186, ArgumentType}:                      _GotoState307Action,
	{_State186, ColonExprType}:                     _GotoState170Action,
	{_State186, AtomExprType}:                      _GotoState63Action,
	{_State186, ParseErrorExprType}:                _GotoState93Action,
	{_State186, LiteralExprType}:                   _GotoState87Action,
	{_State186, NamedExprType}:                     _GotoState90Action,
	{_State186, BlockExprType}:                     _GotoState70Action,
	{_State186, InitializeExprType}:                _GotoState84Action,
	{_State186, ImplicitStructExprType}:            _GotoState80Action,
	{_State186, AccessibleExprType}:                _GotoState105Action,
	{_State186, AccessExprType}:                    _GotoState55Action,
	{_State186, IndexExprType}:                     _GotoState82Action,
	{_State186, PostfixableExprType}:               _GotoState95Action,
	{_State186, PostfixUnaryExprType}:              _GotoState94Action,
	{_State186, PrefixableExprType}:                _GotoState98Action,
	{_State186, PrefixUnaryExprType}:               _GotoState96Action,
	{_State186, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State186, MulExprType}:                       _GotoState89Action,
	{_State186, BinaryMulExprType}:                 _GotoState67Action,
	{_State186, AddExprType}:                       _GotoState57Action,
	{_State186, BinaryAddExprType}:                 _GotoState64Action,
	{_State186, CmpExprType}:                       _GotoState74Action,
	{_State186, BinaryCmpExprType}:                 _GotoState66Action,
	{_State186, AndExprType}:                       _GotoState58Action,
	{_State186, BinaryAndExprType}:                 _GotoState65Action,
	{_State186, OrExprType}:                        _GotoState92Action,
	{_State186, BinaryOrExprType}:                  _GotoState69Action,
	{_State186, InitializableTypeExprType}:         _GotoState83Action,
	{_State186, SliceTypeExprType}:                 _GotoState101Action,
	{_State186, ArrayTypeExprType}:                 _GotoState60Action,
	{_State186, MapTypeExprType}:                   _GotoState88Action,
	{_State186, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State186, AnonymousFuncExprType}:             _GotoState59Action,
	{_State192, IntegerLiteralToken}:               _GotoState40Action,
	{_State192, FloatLiteralToken}:                 _GotoState35Action,
	{_State192, RuneLiteralToken}:                  _GotoState48Action,
	{_State192, StringLiteralToken}:                _GotoState49Action,
	{_State192, IdentifierToken}:                   _GotoState38Action,
	{_State192, UnderscoreToken}:                   _GotoState53Action,
	{_State192, TrueToken}:                         _GotoState52Action,
	{_State192, FalseToken}:                        _GotoState34Action,
	{_State192, StructToken}:                       _GotoState50Action,
	{_State192, FuncToken}:                         _GotoState36Action,
	{_State192, VarToken}:                          _GotoState15Action,
	{_State192, LetToken}:                          _GotoState12Action,
	{_State192, NotToken}:                          _GotoState45Action,
	{_State192, LabelDeclToken}:                    _GotoState41Action,
	{_State192, LparenToken}:                       _GotoState43Action,
	{_State192, LbracketToken}:                     _GotoState42Action,
	{_State192, SubToken}:                          _GotoState51Action,
	{_State192, MulToken}:                          _GotoState44Action,
	{_State192, BitNegToken}:                       _GotoState27Action,
	{_State192, BitAndToken}:                       _GotoState26Action,
	{_State192, GreaterToken}:                      _GotoState37Action,
	{_State192, ParseErrorToken}:                   _GotoState46Action,
	{_State192, VarDeclPatternType}:                _GotoState104Action,
	{_State192, VarOrLetType}:                      _GotoState24Action,
	{_State192, ExprType}:                          _GotoState308Action,
	{_State192, OptionalLabelDeclType}:             _GotoState91Action,
	{_State192, SequenceExprType}:                  _GotoState106Action,
	{_State192, CallExprType}:                      _GotoState71Action,
	{_State192, AtomExprType}:                      _GotoState63Action,
	{_State192, ParseErrorExprType}:                _GotoState93Action,
	{_State192, LiteralExprType}:                   _GotoState87Action,
	{_State192, NamedExprType}:                     _GotoState90Action,
	{_State192, BlockExprType}:                     _GotoState70Action,
	{_State192, InitializeExprType}:                _GotoState84Action,
	{_State192, ImplicitStructExprType}:            _GotoState80Action,
	{_State192, AccessibleExprType}:                _GotoState105Action,
	{_State192, AccessExprType}:                    _GotoState55Action,
	{_State192, IndexExprType}:                     _GotoState82Action,
	{_State192, PostfixableExprType}:               _GotoState95Action,
	{_State192, PostfixUnaryExprType}:              _GotoState94Action,
	{_State192, PrefixableExprType}:                _GotoState98Action,
	{_State192, PrefixUnaryExprType}:               _GotoState96Action,
	{_State192, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State192, MulExprType}:                       _GotoState89Action,
	{_State192, BinaryMulExprType}:                 _GotoState67Action,
	{_State192, AddExprType}:                       _GotoState57Action,
	{_State192, BinaryAddExprType}:                 _GotoState64Action,
	{_State192, CmpExprType}:                       _GotoState74Action,
	{_State192, BinaryCmpExprType}:                 _GotoState66Action,
	{_State192, AndExprType}:                       _GotoState58Action,
	{_State192, BinaryAndExprType}:                 _GotoState65Action,
	{_State192, OrExprType}:                        _GotoState92Action,
	{_State192, BinaryOrExprType}:                  _GotoState69Action,
	{_State192, InitializableTypeExprType}:         _GotoState83Action,
	{_State192, SliceTypeExprType}:                 _GotoState101Action,
	{_State192, ArrayTypeExprType}:                 _GotoState60Action,
	{_State192, MapTypeExprType}:                   _GotoState88Action,
	{_State192, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State192, AnonymousFuncExprType}:             _GotoState59Action,
	{_State193, LparenToken}:                       _GotoState309Action,
	{_State199, IntegerLiteralToken}:               _GotoState40Action,
	{_State199, FloatLiteralToken}:                 _GotoState35Action,
	{_State199, RuneLiteralToken}:                  _GotoState48Action,
	{_State199, StringLiteralToken}:                _GotoState49Action,
	{_State199, IdentifierToken}:                   _GotoState38Action,
	{_State199, UnderscoreToken}:                   _GotoState53Action,
	{_State199, TrueToken}:                         _GotoState52Action,
	{_State199, FalseToken}:                        _GotoState34Action,
	{_State199, StructToken}:                       _GotoState50Action,
	{_State199, FuncToken}:                         _GotoState36Action,
	{_State199, NotToken}:                          _GotoState45Action,
	{_State199, LabelDeclToken}:                    _GotoState41Action,
	{_State199, LparenToken}:                       _GotoState43Action,
	{_State199, LbracketToken}:                     _GotoState42Action,
	{_State199, SubToken}:                          _GotoState51Action,
	{_State199, MulToken}:                          _GotoState44Action,
	{_State199, BitNegToken}:                       _GotoState27Action,
	{_State199, BitAndToken}:                       _GotoState26Action,
	{_State199, ParseErrorToken}:                   _GotoState46Action,
	{_State199, OptionalLabelDeclType}:             _GotoState153Action,
	{_State199, CallExprType}:                      _GotoState71Action,
	{_State199, AtomExprType}:                      _GotoState63Action,
	{_State199, ParseErrorExprType}:                _GotoState93Action,
	{_State199, LiteralExprType}:                   _GotoState87Action,
	{_State199, NamedExprType}:                     _GotoState90Action,
	{_State199, BlockExprType}:                     _GotoState70Action,
	{_State199, InitializeExprType}:                _GotoState84Action,
	{_State199, ImplicitStructExprType}:            _GotoState80Action,
	{_State199, AccessibleExprType}:                _GotoState105Action,
	{_State199, AccessExprType}:                    _GotoState55Action,
	{_State199, IndexExprType}:                     _GotoState82Action,
	{_State199, PostfixableExprType}:               _GotoState95Action,
	{_State199, PostfixUnaryExprType}:              _GotoState94Action,
	{_State199, PrefixableExprType}:                _GotoState98Action,
	{_State199, PrefixUnaryExprType}:               _GotoState96Action,
	{_State199, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State199, MulExprType}:                       _GotoState310Action,
	{_State199, BinaryMulExprType}:                 _GotoState67Action,
	{_State199, InitializableTypeExprType}:         _GotoState83Action,
	{_State199, SliceTypeExprType}:                 _GotoState101Action,
	{_State199, ArrayTypeExprType}:                 _GotoState60Action,
	{_State199, MapTypeExprType}:                   _GotoState88Action,
	{_State199, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State199, AnonymousFuncExprType}:             _GotoState59Action,
	{_State200, IntegerLiteralToken}:               _GotoState40Action,
	{_State200, FloatLiteralToken}:                 _GotoState35Action,
	{_State200, RuneLiteralToken}:                  _GotoState48Action,
	{_State200, StringLiteralToken}:                _GotoState49Action,
	{_State200, IdentifierToken}:                   _GotoState38Action,
	{_State200, UnderscoreToken}:                   _GotoState53Action,
	{_State200, TrueToken}:                         _GotoState52Action,
	{_State200, FalseToken}:                        _GotoState34Action,
	{_State200, StructToken}:                       _GotoState50Action,
	{_State200, FuncToken}:                         _GotoState36Action,
	{_State200, NotToken}:                          _GotoState45Action,
	{_State200, LabelDeclToken}:                    _GotoState41Action,
	{_State200, LparenToken}:                       _GotoState43Action,
	{_State200, LbracketToken}:                     _GotoState42Action,
	{_State200, SubToken}:                          _GotoState51Action,
	{_State200, MulToken}:                          _GotoState44Action,
	{_State200, BitNegToken}:                       _GotoState27Action,
	{_State200, BitAndToken}:                       _GotoState26Action,
	{_State200, ParseErrorToken}:                   _GotoState46Action,
	{_State200, OptionalLabelDeclType}:             _GotoState153Action,
	{_State200, CallExprType}:                      _GotoState71Action,
	{_State200, AtomExprType}:                      _GotoState63Action,
	{_State200, ParseErrorExprType}:                _GotoState93Action,
	{_State200, LiteralExprType}:                   _GotoState87Action,
	{_State200, NamedExprType}:                     _GotoState90Action,
	{_State200, BlockExprType}:                     _GotoState70Action,
	{_State200, InitializeExprType}:                _GotoState84Action,
	{_State200, ImplicitStructExprType}:            _GotoState80Action,
	{_State200, AccessibleExprType}:                _GotoState105Action,
	{_State200, AccessExprType}:                    _GotoState55Action,
	{_State200, IndexExprType}:                     _GotoState82Action,
	{_State200, PostfixableExprType}:               _GotoState95Action,
	{_State200, PostfixUnaryExprType}:              _GotoState94Action,
	{_State200, PrefixableExprType}:                _GotoState98Action,
	{_State200, PrefixUnaryExprType}:               _GotoState96Action,
	{_State200, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State200, MulExprType}:                       _GotoState89Action,
	{_State200, BinaryMulExprType}:                 _GotoState67Action,
	{_State200, AddExprType}:                       _GotoState57Action,
	{_State200, BinaryAddExprType}:                 _GotoState64Action,
	{_State200, CmpExprType}:                       _GotoState311Action,
	{_State200, BinaryCmpExprType}:                 _GotoState66Action,
	{_State200, InitializableTypeExprType}:         _GotoState83Action,
	{_State200, SliceTypeExprType}:                 _GotoState101Action,
	{_State200, ArrayTypeExprType}:                 _GotoState60Action,
	{_State200, MapTypeExprType}:                   _GotoState88Action,
	{_State200, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State200, AnonymousFuncExprType}:             _GotoState59Action,
	{_State201, IntegerLiteralToken}:               _GotoState40Action,
	{_State201, FloatLiteralToken}:                 _GotoState35Action,
	{_State201, RuneLiteralToken}:                  _GotoState48Action,
	{_State201, StringLiteralToken}:                _GotoState49Action,
	{_State201, IdentifierToken}:                   _GotoState38Action,
	{_State201, UnderscoreToken}:                   _GotoState53Action,
	{_State201, TrueToken}:                         _GotoState52Action,
	{_State201, FalseToken}:                        _GotoState34Action,
	{_State201, StructToken}:                       _GotoState50Action,
	{_State201, FuncToken}:                         _GotoState36Action,
	{_State201, VarToken}:                          _GotoState15Action,
	{_State201, LetToken}:                          _GotoState12Action,
	{_State201, NotToken}:                          _GotoState45Action,
	{_State201, LabelDeclToken}:                    _GotoState41Action,
	{_State201, LparenToken}:                       _GotoState43Action,
	{_State201, LbracketToken}:                     _GotoState42Action,
	{_State201, SubToken}:                          _GotoState51Action,
	{_State201, MulToken}:                          _GotoState44Action,
	{_State201, BitNegToken}:                       _GotoState27Action,
	{_State201, BitAndToken}:                       _GotoState26Action,
	{_State201, GreaterToken}:                      _GotoState37Action,
	{_State201, ParseErrorToken}:                   _GotoState46Action,
	{_State201, VarDeclPatternType}:                _GotoState104Action,
	{_State201, VarOrLetType}:                      _GotoState24Action,
	{_State201, ExprType}:                          _GotoState312Action,
	{_State201, OptionalLabelDeclType}:             _GotoState91Action,
	{_State201, SequenceExprType}:                  _GotoState106Action,
	{_State201, CallExprType}:                      _GotoState71Action,
	{_State201, AtomExprType}:                      _GotoState63Action,
	{_State201, ParseErrorExprType}:                _GotoState93Action,
	{_State201, LiteralExprType}:                   _GotoState87Action,
	{_State201, NamedExprType}:                     _GotoState90Action,
	{_State201, BlockExprType}:                     _GotoState70Action,
	{_State201, InitializeExprType}:                _GotoState84Action,
	{_State201, ImplicitStructExprType}:            _GotoState80Action,
	{_State201, AccessibleExprType}:                _GotoState105Action,
	{_State201, AccessExprType}:                    _GotoState55Action,
	{_State201, IndexExprType}:                     _GotoState82Action,
	{_State201, PostfixableExprType}:               _GotoState95Action,
	{_State201, PostfixUnaryExprType}:              _GotoState94Action,
	{_State201, PrefixableExprType}:                _GotoState98Action,
	{_State201, PrefixUnaryExprType}:               _GotoState96Action,
	{_State201, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State201, MulExprType}:                       _GotoState89Action,
	{_State201, BinaryMulExprType}:                 _GotoState67Action,
	{_State201, AddExprType}:                       _GotoState57Action,
	{_State201, BinaryAddExprType}:                 _GotoState64Action,
	{_State201, CmpExprType}:                       _GotoState74Action,
	{_State201, BinaryCmpExprType}:                 _GotoState66Action,
	{_State201, AndExprType}:                       _GotoState58Action,
	{_State201, BinaryAndExprType}:                 _GotoState65Action,
	{_State201, OrExprType}:                        _GotoState92Action,
	{_State201, BinaryOrExprType}:                  _GotoState69Action,
	{_State201, InitializableTypeExprType}:         _GotoState83Action,
	{_State201, SliceTypeExprType}:                 _GotoState101Action,
	{_State201, ArrayTypeExprType}:                 _GotoState60Action,
	{_State201, MapTypeExprType}:                   _GotoState88Action,
	{_State201, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State201, AnonymousFuncExprType}:             _GotoState59Action,
	{_State202, LbracketToken}:                     _GotoState186Action,
	{_State202, DotToken}:                          _GotoState185Action,
	{_State202, DollarLbracketToken}:               _GotoState184Action,
	{_State202, GenericTypeArgumentsType}:          _GotoState193Action,
	{_State210, IntegerLiteralToken}:               _GotoState40Action,
	{_State210, FloatLiteralToken}:                 _GotoState35Action,
	{_State210, RuneLiteralToken}:                  _GotoState48Action,
	{_State210, StringLiteralToken}:                _GotoState49Action,
	{_State210, IdentifierToken}:                   _GotoState38Action,
	{_State210, UnderscoreToken}:                   _GotoState53Action,
	{_State210, TrueToken}:                         _GotoState52Action,
	{_State210, FalseToken}:                        _GotoState34Action,
	{_State210, StructToken}:                       _GotoState50Action,
	{_State210, FuncToken}:                         _GotoState36Action,
	{_State210, NotToken}:                          _GotoState45Action,
	{_State210, LabelDeclToken}:                    _GotoState41Action,
	{_State210, LparenToken}:                       _GotoState43Action,
	{_State210, LbracketToken}:                     _GotoState42Action,
	{_State210, SubToken}:                          _GotoState51Action,
	{_State210, MulToken}:                          _GotoState44Action,
	{_State210, BitNegToken}:                       _GotoState27Action,
	{_State210, BitAndToken}:                       _GotoState26Action,
	{_State210, ParseErrorToken}:                   _GotoState46Action,
	{_State210, OptionalLabelDeclType}:             _GotoState153Action,
	{_State210, CallExprType}:                      _GotoState71Action,
	{_State210, AtomExprType}:                      _GotoState63Action,
	{_State210, ParseErrorExprType}:                _GotoState93Action,
	{_State210, LiteralExprType}:                   _GotoState87Action,
	{_State210, NamedExprType}:                     _GotoState90Action,
	{_State210, BlockExprType}:                     _GotoState70Action,
	{_State210, InitializeExprType}:                _GotoState84Action,
	{_State210, ImplicitStructExprType}:            _GotoState80Action,
	{_State210, AccessibleExprType}:                _GotoState105Action,
	{_State210, AccessExprType}:                    _GotoState55Action,
	{_State210, IndexExprType}:                     _GotoState82Action,
	{_State210, PostfixableExprType}:               _GotoState95Action,
	{_State210, PostfixUnaryExprType}:              _GotoState94Action,
	{_State210, PrefixableExprType}:                _GotoState98Action,
	{_State210, PrefixUnaryExprType}:               _GotoState96Action,
	{_State210, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State210, MulExprType}:                       _GotoState89Action,
	{_State210, BinaryMulExprType}:                 _GotoState67Action,
	{_State210, AddExprType}:                       _GotoState313Action,
	{_State210, BinaryAddExprType}:                 _GotoState64Action,
	{_State210, InitializableTypeExprType}:         _GotoState83Action,
	{_State210, SliceTypeExprType}:                 _GotoState101Action,
	{_State210, ArrayTypeExprType}:                 _GotoState60Action,
	{_State210, MapTypeExprType}:                   _GotoState88Action,
	{_State210, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State210, AnonymousFuncExprType}:             _GotoState59Action,
	{_State211, IntegerLiteralToken}:               _GotoState40Action,
	{_State211, FloatLiteralToken}:                 _GotoState35Action,
	{_State211, RuneLiteralToken}:                  _GotoState48Action,
	{_State211, StringLiteralToken}:                _GotoState49Action,
	{_State211, IdentifierToken}:                   _GotoState38Action,
	{_State211, UnderscoreToken}:                   _GotoState53Action,
	{_State211, TrueToken}:                         _GotoState52Action,
	{_State211, FalseToken}:                        _GotoState34Action,
	{_State211, StructToken}:                       _GotoState50Action,
	{_State211, FuncToken}:                         _GotoState36Action,
	{_State211, VarToken}:                          _GotoState15Action,
	{_State211, LetToken}:                          _GotoState12Action,
	{_State211, NotToken}:                          _GotoState45Action,
	{_State211, LabelDeclToken}:                    _GotoState41Action,
	{_State211, LparenToken}:                       _GotoState43Action,
	{_State211, LbracketToken}:                     _GotoState42Action,
	{_State211, SubToken}:                          _GotoState51Action,
	{_State211, MulToken}:                          _GotoState44Action,
	{_State211, BitNegToken}:                       _GotoState27Action,
	{_State211, BitAndToken}:                       _GotoState26Action,
	{_State211, GreaterToken}:                      _GotoState37Action,
	{_State211, ParseErrorToken}:                   _GotoState46Action,
	{_State211, VarDeclPatternType}:                _GotoState104Action,
	{_State211, VarOrLetType}:                      _GotoState24Action,
	{_State211, ExprType}:                          _GotoState314Action,
	{_State211, OptionalLabelDeclType}:             _GotoState91Action,
	{_State211, SequenceExprType}:                  _GotoState106Action,
	{_State211, CallExprType}:                      _GotoState71Action,
	{_State211, AtomExprType}:                      _GotoState63Action,
	{_State211, ParseErrorExprType}:                _GotoState93Action,
	{_State211, LiteralExprType}:                   _GotoState87Action,
	{_State211, NamedExprType}:                     _GotoState90Action,
	{_State211, BlockExprType}:                     _GotoState70Action,
	{_State211, InitializeExprType}:                _GotoState84Action,
	{_State211, ImplicitStructExprType}:            _GotoState80Action,
	{_State211, AccessibleExprType}:                _GotoState105Action,
	{_State211, AccessExprType}:                    _GotoState55Action,
	{_State211, IndexExprType}:                     _GotoState82Action,
	{_State211, PostfixableExprType}:               _GotoState95Action,
	{_State211, PostfixUnaryExprType}:              _GotoState94Action,
	{_State211, PrefixableExprType}:                _GotoState98Action,
	{_State211, PrefixUnaryExprType}:               _GotoState96Action,
	{_State211, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State211, MulExprType}:                       _GotoState89Action,
	{_State211, BinaryMulExprType}:                 _GotoState67Action,
	{_State211, AddExprType}:                       _GotoState57Action,
	{_State211, BinaryAddExprType}:                 _GotoState64Action,
	{_State211, CmpExprType}:                       _GotoState74Action,
	{_State211, BinaryCmpExprType}:                 _GotoState66Action,
	{_State211, AndExprType}:                       _GotoState58Action,
	{_State211, BinaryAndExprType}:                 _GotoState65Action,
	{_State211, OrExprType}:                        _GotoState92Action,
	{_State211, BinaryOrExprType}:                  _GotoState69Action,
	{_State211, InitializableTypeExprType}:         _GotoState83Action,
	{_State211, SliceTypeExprType}:                 _GotoState101Action,
	{_State211, ArrayTypeExprType}:                 _GotoState60Action,
	{_State211, MapTypeExprType}:                   _GotoState88Action,
	{_State211, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State211, AnonymousFuncExprType}:             _GotoState59Action,
	{_State212, IntegerLiteralToken}:               _GotoState40Action,
	{_State212, FloatLiteralToken}:                 _GotoState35Action,
	{_State212, RuneLiteralToken}:                  _GotoState48Action,
	{_State212, StringLiteralToken}:                _GotoState49Action,
	{_State212, IdentifierToken}:                   _GotoState167Action,
	{_State212, UnderscoreToken}:                   _GotoState53Action,
	{_State212, TrueToken}:                         _GotoState52Action,
	{_State212, FalseToken}:                        _GotoState34Action,
	{_State212, StructToken}:                       _GotoState50Action,
	{_State212, FuncToken}:                         _GotoState36Action,
	{_State212, VarToken}:                          _GotoState15Action,
	{_State212, LetToken}:                          _GotoState12Action,
	{_State212, NotToken}:                          _GotoState45Action,
	{_State212, LabelDeclToken}:                    _GotoState41Action,
	{_State212, LparenToken}:                       _GotoState43Action,
	{_State212, LbracketToken}:                     _GotoState42Action,
	{_State212, ColonToken}:                        _GotoState165Action,
	{_State212, EllipsisToken}:                     _GotoState166Action,
	{_State212, SubToken}:                          _GotoState51Action,
	{_State212, MulToken}:                          _GotoState44Action,
	{_State212, BitNegToken}:                       _GotoState27Action,
	{_State212, BitAndToken}:                       _GotoState26Action,
	{_State212, GreaterToken}:                      _GotoState37Action,
	{_State212, ParseErrorToken}:                   _GotoState46Action,
	{_State212, VarDeclPatternType}:                _GotoState104Action,
	{_State212, VarOrLetType}:                      _GotoState24Action,
	{_State212, ExprType}:                          _GotoState171Action,
	{_State212, OptionalLabelDeclType}:             _GotoState91Action,
	{_State212, SequenceExprType}:                  _GotoState106Action,
	{_State212, CallExprType}:                      _GotoState71Action,
	{_State212, ProperArgumentsType}:               _GotoState172Action,
	{_State212, ArgumentsType}:                     _GotoState315Action,
	{_State212, ArgumentType}:                      _GotoState168Action,
	{_State212, ColonExprType}:                     _GotoState170Action,
	{_State212, AtomExprType}:                      _GotoState63Action,
	{_State212, ParseErrorExprType}:                _GotoState93Action,
	{_State212, LiteralExprType}:                   _GotoState87Action,
	{_State212, NamedExprType}:                     _GotoState90Action,
	{_State212, BlockExprType}:                     _GotoState70Action,
	{_State212, InitializeExprType}:                _GotoState84Action,
	{_State212, ImplicitStructExprType}:            _GotoState80Action,
	{_State212, AccessibleExprType}:                _GotoState105Action,
	{_State212, AccessExprType}:                    _GotoState55Action,
	{_State212, IndexExprType}:                     _GotoState82Action,
	{_State212, PostfixableExprType}:               _GotoState95Action,
	{_State212, PostfixUnaryExprType}:              _GotoState94Action,
	{_State212, PrefixableExprType}:                _GotoState98Action,
	{_State212, PrefixUnaryExprType}:               _GotoState96Action,
	{_State212, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State212, MulExprType}:                       _GotoState89Action,
	{_State212, BinaryMulExprType}:                 _GotoState67Action,
	{_State212, AddExprType}:                       _GotoState57Action,
	{_State212, BinaryAddExprType}:                 _GotoState64Action,
	{_State212, CmpExprType}:                       _GotoState74Action,
	{_State212, BinaryCmpExprType}:                 _GotoState66Action,
	{_State212, AndExprType}:                       _GotoState58Action,
	{_State212, BinaryAndExprType}:                 _GotoState65Action,
	{_State212, OrExprType}:                        _GotoState92Action,
	{_State212, BinaryOrExprType}:                  _GotoState69Action,
	{_State212, InitializableTypeExprType}:         _GotoState83Action,
	{_State212, SliceTypeExprType}:                 _GotoState101Action,
	{_State212, ArrayTypeExprType}:                 _GotoState60Action,
	{_State212, MapTypeExprType}:                   _GotoState88Action,
	{_State212, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State212, AnonymousFuncExprType}:             _GotoState59Action,
	{_State213, IntegerLiteralToken}:               _GotoState40Action,
	{_State213, FloatLiteralToken}:                 _GotoState35Action,
	{_State213, RuneLiteralToken}:                  _GotoState48Action,
	{_State213, StringLiteralToken}:                _GotoState49Action,
	{_State213, IdentifierToken}:                   _GotoState38Action,
	{_State213, UnderscoreToken}:                   _GotoState53Action,
	{_State213, TrueToken}:                         _GotoState52Action,
	{_State213, FalseToken}:                        _GotoState34Action,
	{_State213, StructToken}:                       _GotoState50Action,
	{_State213, FuncToken}:                         _GotoState36Action,
	{_State213, VarToken}:                          _GotoState15Action,
	{_State213, LetToken}:                          _GotoState12Action,
	{_State213, NotToken}:                          _GotoState45Action,
	{_State213, LabelDeclToken}:                    _GotoState41Action,
	{_State213, LparenToken}:                       _GotoState43Action,
	{_State213, LbracketToken}:                     _GotoState42Action,
	{_State213, SubToken}:                          _GotoState51Action,
	{_State213, MulToken}:                          _GotoState44Action,
	{_State213, BitNegToken}:                       _GotoState27Action,
	{_State213, BitAndToken}:                       _GotoState26Action,
	{_State213, GreaterToken}:                      _GotoState37Action,
	{_State213, ParseErrorToken}:                   _GotoState46Action,
	{_State213, ExprsType}:                         _GotoState316Action,
	{_State213, VarDeclPatternType}:                _GotoState104Action,
	{_State213, VarOrLetType}:                      _GotoState24Action,
	{_State213, ExprType}:                          _GotoState76Action,
	{_State213, OptionalLabelDeclType}:             _GotoState91Action,
	{_State213, SequenceExprType}:                  _GotoState106Action,
	{_State213, CallExprType}:                      _GotoState71Action,
	{_State213, AtomExprType}:                      _GotoState63Action,
	{_State213, ParseErrorExprType}:                _GotoState93Action,
	{_State213, LiteralExprType}:                   _GotoState87Action,
	{_State213, NamedExprType}:                     _GotoState90Action,
	{_State213, BlockExprType}:                     _GotoState70Action,
	{_State213, InitializeExprType}:                _GotoState84Action,
	{_State213, ImplicitStructExprType}:            _GotoState80Action,
	{_State213, AccessibleExprType}:                _GotoState105Action,
	{_State213, AccessExprType}:                    _GotoState55Action,
	{_State213, IndexExprType}:                     _GotoState82Action,
	{_State213, PostfixableExprType}:               _GotoState95Action,
	{_State213, PostfixUnaryExprType}:              _GotoState94Action,
	{_State213, PrefixableExprType}:                _GotoState98Action,
	{_State213, PrefixUnaryExprType}:               _GotoState96Action,
	{_State213, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State213, MulExprType}:                       _GotoState89Action,
	{_State213, BinaryMulExprType}:                 _GotoState67Action,
	{_State213, AddExprType}:                       _GotoState57Action,
	{_State213, BinaryAddExprType}:                 _GotoState64Action,
	{_State213, CmpExprType}:                       _GotoState74Action,
	{_State213, BinaryCmpExprType}:                 _GotoState66Action,
	{_State213, AndExprType}:                       _GotoState58Action,
	{_State213, BinaryAndExprType}:                 _GotoState65Action,
	{_State213, OrExprType}:                        _GotoState92Action,
	{_State213, BinaryOrExprType}:                  _GotoState69Action,
	{_State213, InitializableTypeExprType}:         _GotoState83Action,
	{_State213, SliceTypeExprType}:                 _GotoState101Action,
	{_State213, ArrayTypeExprType}:                 _GotoState60Action,
	{_State213, MapTypeExprType}:                   _GotoState88Action,
	{_State213, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State213, AnonymousFuncExprType}:             _GotoState59Action,
	{_State214, CommaToken}:                        _GotoState211Action,
	{_State221, IntegerLiteralToken}:               _GotoState40Action,
	{_State221, FloatLiteralToken}:                 _GotoState35Action,
	{_State221, RuneLiteralToken}:                  _GotoState48Action,
	{_State221, StringLiteralToken}:                _GotoState49Action,
	{_State221, IdentifierToken}:                   _GotoState38Action,
	{_State221, UnderscoreToken}:                   _GotoState53Action,
	{_State221, TrueToken}:                         _GotoState52Action,
	{_State221, FalseToken}:                        _GotoState34Action,
	{_State221, StructToken}:                       _GotoState50Action,
	{_State221, FuncToken}:                         _GotoState36Action,
	{_State221, NotToken}:                          _GotoState45Action,
	{_State221, LabelDeclToken}:                    _GotoState41Action,
	{_State221, LparenToken}:                       _GotoState43Action,
	{_State221, LbracketToken}:                     _GotoState42Action,
	{_State221, SubToken}:                          _GotoState51Action,
	{_State221, MulToken}:                          _GotoState44Action,
	{_State221, BitNegToken}:                       _GotoState27Action,
	{_State221, BitAndToken}:                       _GotoState26Action,
	{_State221, ParseErrorToken}:                   _GotoState46Action,
	{_State221, OptionalLabelDeclType}:             _GotoState153Action,
	{_State221, CallExprType}:                      _GotoState71Action,
	{_State221, AtomExprType}:                      _GotoState63Action,
	{_State221, ParseErrorExprType}:                _GotoState93Action,
	{_State221, LiteralExprType}:                   _GotoState87Action,
	{_State221, NamedExprType}:                     _GotoState90Action,
	{_State221, BlockExprType}:                     _GotoState70Action,
	{_State221, InitializeExprType}:                _GotoState84Action,
	{_State221, ImplicitStructExprType}:            _GotoState80Action,
	{_State221, AccessibleExprType}:                _GotoState105Action,
	{_State221, AccessExprType}:                    _GotoState55Action,
	{_State221, IndexExprType}:                     _GotoState82Action,
	{_State221, PostfixableExprType}:               _GotoState95Action,
	{_State221, PostfixUnaryExprType}:              _GotoState94Action,
	{_State221, PrefixableExprType}:                _GotoState317Action,
	{_State221, PrefixUnaryExprType}:               _GotoState96Action,
	{_State221, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State221, InitializableTypeExprType}:         _GotoState83Action,
	{_State221, SliceTypeExprType}:                 _GotoState101Action,
	{_State221, ArrayTypeExprType}:                 _GotoState60Action,
	{_State221, MapTypeExprType}:                   _GotoState88Action,
	{_State221, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State221, AnonymousFuncExprType}:             _GotoState59Action,
	{_State222, LbraceToken}:                       _GotoState11Action,
	{_State222, StatementBlockType}:                _GotoState318Action,
	{_State223, IntegerLiteralToken}:               _GotoState40Action,
	{_State223, FloatLiteralToken}:                 _GotoState35Action,
	{_State223, RuneLiteralToken}:                  _GotoState48Action,
	{_State223, StringLiteralToken}:                _GotoState49Action,
	{_State223, IdentifierToken}:                   _GotoState38Action,
	{_State223, UnderscoreToken}:                   _GotoState53Action,
	{_State223, TrueToken}:                         _GotoState52Action,
	{_State223, FalseToken}:                        _GotoState34Action,
	{_State223, StructToken}:                       _GotoState50Action,
	{_State223, FuncToken}:                         _GotoState36Action,
	{_State223, VarToken}:                          _GotoState15Action,
	{_State223, LetToken}:                          _GotoState12Action,
	{_State223, NotToken}:                          _GotoState45Action,
	{_State223, LabelDeclToken}:                    _GotoState41Action,
	{_State223, LparenToken}:                       _GotoState43Action,
	{_State223, LbracketToken}:                     _GotoState42Action,
	{_State223, SubToken}:                          _GotoState51Action,
	{_State223, MulToken}:                          _GotoState44Action,
	{_State223, BitNegToken}:                       _GotoState27Action,
	{_State223, BitAndToken}:                       _GotoState26Action,
	{_State223, GreaterToken}:                      _GotoState37Action,
	{_State223, ParseErrorToken}:                   _GotoState46Action,
	{_State223, VarDeclPatternType}:                _GotoState104Action,
	{_State223, VarOrLetType}:                      _GotoState24Action,
	{_State223, AssignPatternType}:                 _GotoState319Action,
	{_State223, OptionalLabelDeclType}:             _GotoState153Action,
	{_State223, SequenceExprType}:                  _GotoState321Action,
	{_State223, ForAssignmentType}:                 _GotoState320Action,
	{_State223, CallExprType}:                      _GotoState71Action,
	{_State223, AtomExprType}:                      _GotoState63Action,
	{_State223, ParseErrorExprType}:                _GotoState93Action,
	{_State223, LiteralExprType}:                   _GotoState87Action,
	{_State223, NamedExprType}:                     _GotoState90Action,
	{_State223, BlockExprType}:                     _GotoState70Action,
	{_State223, InitializeExprType}:                _GotoState84Action,
	{_State223, ImplicitStructExprType}:            _GotoState80Action,
	{_State223, AccessibleExprType}:                _GotoState105Action,
	{_State223, AccessExprType}:                    _GotoState55Action,
	{_State223, IndexExprType}:                     _GotoState82Action,
	{_State223, PostfixableExprType}:               _GotoState95Action,
	{_State223, PostfixUnaryExprType}:              _GotoState94Action,
	{_State223, PrefixableExprType}:                _GotoState98Action,
	{_State223, PrefixUnaryExprType}:               _GotoState96Action,
	{_State223, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State223, MulExprType}:                       _GotoState89Action,
	{_State223, BinaryMulExprType}:                 _GotoState67Action,
	{_State223, AddExprType}:                       _GotoState57Action,
	{_State223, BinaryAddExprType}:                 _GotoState64Action,
	{_State223, CmpExprType}:                       _GotoState74Action,
	{_State223, BinaryCmpExprType}:                 _GotoState66Action,
	{_State223, AndExprType}:                       _GotoState58Action,
	{_State223, BinaryAndExprType}:                 _GotoState65Action,
	{_State223, OrExprType}:                        _GotoState92Action,
	{_State223, BinaryOrExprType}:                  _GotoState69Action,
	{_State223, InitializableTypeExprType}:         _GotoState83Action,
	{_State223, SliceTypeExprType}:                 _GotoState101Action,
	{_State223, ArrayTypeExprType}:                 _GotoState60Action,
	{_State223, MapTypeExprType}:                   _GotoState88Action,
	{_State223, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State223, AnonymousFuncExprType}:             _GotoState59Action,
	{_State224, IntegerLiteralToken}:               _GotoState40Action,
	{_State224, FloatLiteralToken}:                 _GotoState35Action,
	{_State224, RuneLiteralToken}:                  _GotoState48Action,
	{_State224, StringLiteralToken}:                _GotoState49Action,
	{_State224, IdentifierToken}:                   _GotoState38Action,
	{_State224, UnderscoreToken}:                   _GotoState53Action,
	{_State224, TrueToken}:                         _GotoState52Action,
	{_State224, FalseToken}:                        _GotoState34Action,
	{_State224, CaseToken}:                         _GotoState322Action,
	{_State224, StructToken}:                       _GotoState50Action,
	{_State224, FuncToken}:                         _GotoState36Action,
	{_State224, VarToken}:                          _GotoState15Action,
	{_State224, LetToken}:                          _GotoState12Action,
	{_State224, NotToken}:                          _GotoState45Action,
	{_State224, LabelDeclToken}:                    _GotoState41Action,
	{_State224, LparenToken}:                       _GotoState43Action,
	{_State224, LbracketToken}:                     _GotoState42Action,
	{_State224, SubToken}:                          _GotoState51Action,
	{_State224, MulToken}:                          _GotoState44Action,
	{_State224, BitNegToken}:                       _GotoState27Action,
	{_State224, BitAndToken}:                       _GotoState26Action,
	{_State224, GreaterToken}:                      _GotoState37Action,
	{_State224, ParseErrorToken}:                   _GotoState46Action,
	{_State224, VarDeclPatternType}:                _GotoState104Action,
	{_State224, VarOrLetType}:                      _GotoState24Action,
	{_State224, OptionalLabelDeclType}:             _GotoState153Action,
	{_State224, SequenceExprType}:                  _GotoState324Action,
	{_State224, ConditionType}:                     _GotoState323Action,
	{_State224, CallExprType}:                      _GotoState71Action,
	{_State224, AtomExprType}:                      _GotoState63Action,
	{_State224, ParseErrorExprType}:                _GotoState93Action,
	{_State224, LiteralExprType}:                   _GotoState87Action,
	{_State224, NamedExprType}:                     _GotoState90Action,
	{_State224, BlockExprType}:                     _GotoState70Action,
	{_State224, InitializeExprType}:                _GotoState84Action,
	{_State224, ImplicitStructExprType}:            _GotoState80Action,
	{_State224, AccessibleExprType}:                _GotoState105Action,
	{_State224, AccessExprType}:                    _GotoState55Action,
	{_State224, IndexExprType}:                     _GotoState82Action,
	{_State224, PostfixableExprType}:               _GotoState95Action,
	{_State224, PostfixUnaryExprType}:              _GotoState94Action,
	{_State224, PrefixableExprType}:                _GotoState98Action,
	{_State224, PrefixUnaryExprType}:               _GotoState96Action,
	{_State224, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State224, MulExprType}:                       _GotoState89Action,
	{_State224, BinaryMulExprType}:                 _GotoState67Action,
	{_State224, AddExprType}:                       _GotoState57Action,
	{_State224, BinaryAddExprType}:                 _GotoState64Action,
	{_State224, CmpExprType}:                       _GotoState74Action,
	{_State224, BinaryCmpExprType}:                 _GotoState66Action,
	{_State224, AndExprType}:                       _GotoState58Action,
	{_State224, BinaryAndExprType}:                 _GotoState65Action,
	{_State224, OrExprType}:                        _GotoState92Action,
	{_State224, BinaryOrExprType}:                  _GotoState69Action,
	{_State224, InitializableTypeExprType}:         _GotoState83Action,
	{_State224, SliceTypeExprType}:                 _GotoState101Action,
	{_State224, ArrayTypeExprType}:                 _GotoState60Action,
	{_State224, MapTypeExprType}:                   _GotoState88Action,
	{_State224, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State224, AnonymousFuncExprType}:             _GotoState59Action,
	{_State225, IntegerLiteralToken}:               _GotoState40Action,
	{_State225, FloatLiteralToken}:                 _GotoState35Action,
	{_State225, RuneLiteralToken}:                  _GotoState48Action,
	{_State225, StringLiteralToken}:                _GotoState49Action,
	{_State225, IdentifierToken}:                   _GotoState38Action,
	{_State225, UnderscoreToken}:                   _GotoState53Action,
	{_State225, TrueToken}:                         _GotoState52Action,
	{_State225, FalseToken}:                        _GotoState34Action,
	{_State225, StructToken}:                       _GotoState50Action,
	{_State225, FuncToken}:                         _GotoState36Action,
	{_State225, VarToken}:                          _GotoState15Action,
	{_State225, LetToken}:                          _GotoState12Action,
	{_State225, NotToken}:                          _GotoState45Action,
	{_State225, LabelDeclToken}:                    _GotoState41Action,
	{_State225, LparenToken}:                       _GotoState43Action,
	{_State225, LbracketToken}:                     _GotoState42Action,
	{_State225, SubToken}:                          _GotoState51Action,
	{_State225, MulToken}:                          _GotoState44Action,
	{_State225, BitNegToken}:                       _GotoState27Action,
	{_State225, BitAndToken}:                       _GotoState26Action,
	{_State225, GreaterToken}:                      _GotoState37Action,
	{_State225, ParseErrorToken}:                   _GotoState46Action,
	{_State225, VarDeclPatternType}:                _GotoState104Action,
	{_State225, VarOrLetType}:                      _GotoState24Action,
	{_State225, OptionalLabelDeclType}:             _GotoState153Action,
	{_State225, SequenceExprType}:                  _GotoState325Action,
	{_State225, CallExprType}:                      _GotoState71Action,
	{_State225, AtomExprType}:                      _GotoState63Action,
	{_State225, ParseErrorExprType}:                _GotoState93Action,
	{_State225, LiteralExprType}:                   _GotoState87Action,
	{_State225, NamedExprType}:                     _GotoState90Action,
	{_State225, BlockExprType}:                     _GotoState70Action,
	{_State225, InitializeExprType}:                _GotoState84Action,
	{_State225, ImplicitStructExprType}:            _GotoState80Action,
	{_State225, AccessibleExprType}:                _GotoState105Action,
	{_State225, AccessExprType}:                    _GotoState55Action,
	{_State225, IndexExprType}:                     _GotoState82Action,
	{_State225, PostfixableExprType}:               _GotoState95Action,
	{_State225, PostfixUnaryExprType}:              _GotoState94Action,
	{_State225, PrefixableExprType}:                _GotoState98Action,
	{_State225, PrefixUnaryExprType}:               _GotoState96Action,
	{_State225, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State225, MulExprType}:                       _GotoState89Action,
	{_State225, BinaryMulExprType}:                 _GotoState67Action,
	{_State225, AddExprType}:                       _GotoState57Action,
	{_State225, BinaryAddExprType}:                 _GotoState64Action,
	{_State225, CmpExprType}:                       _GotoState74Action,
	{_State225, BinaryCmpExprType}:                 _GotoState66Action,
	{_State225, AndExprType}:                       _GotoState58Action,
	{_State225, BinaryAndExprType}:                 _GotoState65Action,
	{_State225, OrExprType}:                        _GotoState92Action,
	{_State225, BinaryOrExprType}:                  _GotoState69Action,
	{_State225, InitializableTypeExprType}:         _GotoState83Action,
	{_State225, SliceTypeExprType}:                 _GotoState101Action,
	{_State225, ArrayTypeExprType}:                 _GotoState60Action,
	{_State225, MapTypeExprType}:                   _GotoState88Action,
	{_State225, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State225, AnonymousFuncExprType}:             _GotoState59Action,
	{_State230, IntegerLiteralToken}:               _GotoState40Action,
	{_State230, FloatLiteralToken}:                 _GotoState35Action,
	{_State230, RuneLiteralToken}:                  _GotoState48Action,
	{_State230, StringLiteralToken}:                _GotoState49Action,
	{_State230, IdentifierToken}:                   _GotoState38Action,
	{_State230, UnderscoreToken}:                   _GotoState53Action,
	{_State230, TrueToken}:                         _GotoState52Action,
	{_State230, FalseToken}:                        _GotoState34Action,
	{_State230, StructToken}:                       _GotoState50Action,
	{_State230, FuncToken}:                         _GotoState36Action,
	{_State230, NotToken}:                          _GotoState45Action,
	{_State230, LabelDeclToken}:                    _GotoState41Action,
	{_State230, LparenToken}:                       _GotoState43Action,
	{_State230, LbracketToken}:                     _GotoState42Action,
	{_State230, SubToken}:                          _GotoState51Action,
	{_State230, MulToken}:                          _GotoState44Action,
	{_State230, BitNegToken}:                       _GotoState27Action,
	{_State230, BitAndToken}:                       _GotoState26Action,
	{_State230, ParseErrorToken}:                   _GotoState46Action,
	{_State230, OptionalLabelDeclType}:             _GotoState153Action,
	{_State230, CallExprType}:                      _GotoState71Action,
	{_State230, AtomExprType}:                      _GotoState63Action,
	{_State230, ParseErrorExprType}:                _GotoState93Action,
	{_State230, LiteralExprType}:                   _GotoState87Action,
	{_State230, NamedExprType}:                     _GotoState90Action,
	{_State230, BlockExprType}:                     _GotoState70Action,
	{_State230, InitializeExprType}:                _GotoState84Action,
	{_State230, ImplicitStructExprType}:            _GotoState80Action,
	{_State230, AccessibleExprType}:                _GotoState105Action,
	{_State230, AccessExprType}:                    _GotoState55Action,
	{_State230, IndexExprType}:                     _GotoState82Action,
	{_State230, PostfixableExprType}:               _GotoState95Action,
	{_State230, PostfixUnaryExprType}:              _GotoState94Action,
	{_State230, PrefixableExprType}:                _GotoState98Action,
	{_State230, PrefixUnaryExprType}:               _GotoState96Action,
	{_State230, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State230, MulExprType}:                       _GotoState89Action,
	{_State230, BinaryMulExprType}:                 _GotoState67Action,
	{_State230, AddExprType}:                       _GotoState57Action,
	{_State230, BinaryAddExprType}:                 _GotoState64Action,
	{_State230, CmpExprType}:                       _GotoState74Action,
	{_State230, BinaryCmpExprType}:                 _GotoState66Action,
	{_State230, AndExprType}:                       _GotoState326Action,
	{_State230, BinaryAndExprType}:                 _GotoState65Action,
	{_State230, InitializableTypeExprType}:         _GotoState83Action,
	{_State230, SliceTypeExprType}:                 _GotoState101Action,
	{_State230, ArrayTypeExprType}:                 _GotoState60Action,
	{_State230, MapTypeExprType}:                   _GotoState88Action,
	{_State230, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State230, AnonymousFuncExprType}:             _GotoState59Action,
	{_State232, IdentifierToken}:                   _GotoState237Action,
	{_State232, UnderscoreToken}:                   _GotoState238Action,
	{_State232, UnsafeToken}:                       _GotoState54Action,
	{_State232, StructToken}:                       _GotoState50Action,
	{_State232, EnumToken}:                         _GotoState110Action,
	{_State232, TraitToken}:                        _GotoState118Action,
	{_State232, FuncToken}:                         _GotoState236Action,
	{_State232, LparenToken}:                       _GotoState114Action,
	{_State232, LbracketToken}:                     _GotoState42Action,
	{_State232, DotToken}:                          _GotoState109Action,
	{_State232, QuestionToken}:                     _GotoState116Action,
	{_State232, ExclaimToken}:                      _GotoState111Action,
	{_State232, TildeTildeToken}:                   _GotoState117Action,
	{_State232, BitNegToken}:                       _GotoState108Action,
	{_State232, BitAndToken}:                       _GotoState107Action,
	{_State232, ParseErrorToken}:                   _GotoState115Action,
	{_State232, UnsafeStatementType}:               _GotoState246Action,
	{_State232, InitializableTypeExprType}:         _GotoState127Action,
	{_State232, SliceTypeExprType}:                 _GotoState101Action,
	{_State232, ArrayTypeExprType}:                 _GotoState60Action,
	{_State232, MapTypeExprType}:                   _GotoState88Action,
	{_State232, AtomTypeExprType}:                  _GotoState120Action,
	{_State232, NamedTypeExprType}:                 _GotoState128Action,
	{_State232, InferredTypeExprType}:              _GotoState126Action,
	{_State232, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State232, ReturnableTypeExprType}:            _GotoState132Action,
	{_State232, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State232, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State232, TypeExprType}:                      _GotoState245Action,
	{_State232, BinaryTypeExprType}:                _GotoState121Action,
	{_State232, FieldDefType}:                      _GotoState328Action,
	{_State232, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State232, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State232, TraitTypeExprType}:                 _GotoState133Action,
	{_State232, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State232, ProperExplicitEnumValueDefsType}:   _GotoState329Action,
	{_State232, ExplicitEnumValueDefsType}:         _GotoState327Action,
	{_State232, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State232, FuncTypeExprType}:                  _GotoState123Action,
	{_State232, MethodSignatureType}:               _GotoState242Action,
	{_State233, IdentifierToken}:                   _GotoState331Action,
	{_State233, UnderscoreToken}:                   _GotoState332Action,
	{_State233, StructToken}:                       _GotoState50Action,
	{_State233, EnumToken}:                         _GotoState110Action,
	{_State233, TraitToken}:                        _GotoState118Action,
	{_State233, FuncToken}:                         _GotoState112Action,
	{_State233, LparenToken}:                       _GotoState114Action,
	{_State233, LbracketToken}:                     _GotoState42Action,
	{_State233, DotToken}:                          _GotoState109Action,
	{_State233, QuestionToken}:                     _GotoState116Action,
	{_State233, ExclaimToken}:                      _GotoState111Action,
	{_State233, EllipsisToken}:                     _GotoState330Action,
	{_State233, TildeTildeToken}:                   _GotoState117Action,
	{_State233, BitNegToken}:                       _GotoState108Action,
	{_State233, BitAndToken}:                       _GotoState107Action,
	{_State233, ParseErrorToken}:                   _GotoState115Action,
	{_State233, InitializableTypeExprType}:         _GotoState127Action,
	{_State233, SliceTypeExprType}:                 _GotoState101Action,
	{_State233, ArrayTypeExprType}:                 _GotoState60Action,
	{_State233, MapTypeExprType}:                   _GotoState88Action,
	{_State233, AtomTypeExprType}:                  _GotoState120Action,
	{_State233, NamedTypeExprType}:                 _GotoState128Action,
	{_State233, InferredTypeExprType}:              _GotoState126Action,
	{_State233, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State233, ReturnableTypeExprType}:            _GotoState132Action,
	{_State233, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State233, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State233, TypeExprType}:                      _GotoState337Action,
	{_State233, BinaryTypeExprType}:                _GotoState121Action,
	{_State233, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State233, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State233, TraitTypeExprType}:                 _GotoState133Action,
	{_State233, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State233, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State233, ProperParameterDefType}:            _GotoState336Action,
	{_State233, ParameterDeclType}:                 _GotoState333Action,
	{_State233, ProperParameterDeclsType}:          _GotoState335Action,
	{_State233, ParameterDeclsType}:                _GotoState334Action,
	{_State233, FuncTypeExprType}:                  _GotoState123Action,
	{_State234, IdentifierToken}:                   _GotoState338Action,
	{_State236, IdentifierToken}:                   _GotoState339Action,
	{_State236, LparenToken}:                       _GotoState233Action,
	{_State237, IdentifierToken}:                   _GotoState113Action,
	{_State237, UnderscoreToken}:                   _GotoState119Action,
	{_State237, StructToken}:                       _GotoState50Action,
	{_State237, EnumToken}:                         _GotoState110Action,
	{_State237, TraitToken}:                        _GotoState118Action,
	{_State237, FuncToken}:                         _GotoState112Action,
	{_State237, LparenToken}:                       _GotoState114Action,
	{_State237, LbracketToken}:                     _GotoState42Action,
	{_State237, DotToken}:                          _GotoState340Action,
	{_State237, QuestionToken}:                     _GotoState116Action,
	{_State237, ExclaimToken}:                      _GotoState111Action,
	{_State237, DollarLbracketToken}:               _GotoState184Action,
	{_State237, TildeTildeToken}:                   _GotoState117Action,
	{_State237, BitNegToken}:                       _GotoState108Action,
	{_State237, BitAndToken}:                       _GotoState107Action,
	{_State237, ParseErrorToken}:                   _GotoState115Action,
	{_State237, InitializableTypeExprType}:         _GotoState127Action,
	{_State237, SliceTypeExprType}:                 _GotoState101Action,
	{_State237, ArrayTypeExprType}:                 _GotoState60Action,
	{_State237, MapTypeExprType}:                   _GotoState88Action,
	{_State237, AtomTypeExprType}:                  _GotoState120Action,
	{_State237, NamedTypeExprType}:                 _GotoState128Action,
	{_State237, InferredTypeExprType}:              _GotoState126Action,
	{_State237, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State237, ReturnableTypeExprType}:            _GotoState132Action,
	{_State237, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State237, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State237, TypeExprType}:                      _GotoState341Action,
	{_State237, BinaryTypeExprType}:                _GotoState121Action,
	{_State237, GenericTypeArgumentsType}:          _GotoState235Action,
	{_State237, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State237, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State237, TraitTypeExprType}:                 _GotoState133Action,
	{_State237, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State237, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State237, FuncTypeExprType}:                  _GotoState123Action,
	{_State238, IdentifierToken}:                   _GotoState113Action,
	{_State238, UnderscoreToken}:                   _GotoState119Action,
	{_State238, StructToken}:                       _GotoState50Action,
	{_State238, EnumToken}:                         _GotoState110Action,
	{_State238, TraitToken}:                        _GotoState118Action,
	{_State238, FuncToken}:                         _GotoState112Action,
	{_State238, LparenToken}:                       _GotoState114Action,
	{_State238, LbracketToken}:                     _GotoState42Action,
	{_State238, DotToken}:                          _GotoState109Action,
	{_State238, QuestionToken}:                     _GotoState116Action,
	{_State238, ExclaimToken}:                      _GotoState111Action,
	{_State238, TildeTildeToken}:                   _GotoState117Action,
	{_State238, BitNegToken}:                       _GotoState108Action,
	{_State238, BitAndToken}:                       _GotoState107Action,
	{_State238, ParseErrorToken}:                   _GotoState115Action,
	{_State238, InitializableTypeExprType}:         _GotoState127Action,
	{_State238, SliceTypeExprType}:                 _GotoState101Action,
	{_State238, ArrayTypeExprType}:                 _GotoState60Action,
	{_State238, MapTypeExprType}:                   _GotoState88Action,
	{_State238, AtomTypeExprType}:                  _GotoState120Action,
	{_State238, NamedTypeExprType}:                 _GotoState128Action,
	{_State238, InferredTypeExprType}:              _GotoState126Action,
	{_State238, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State238, ReturnableTypeExprType}:            _GotoState132Action,
	{_State238, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State238, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State238, TypeExprType}:                      _GotoState342Action,
	{_State238, BinaryTypeExprType}:                _GotoState121Action,
	{_State238, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State238, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State238, TraitTypeExprType}:                 _GotoState133Action,
	{_State238, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State238, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State238, FuncTypeExprType}:                  _GotoState123Action,
	{_State239, OrToken}:                           _GotoState343Action,
	{_State240, RparenToken}:                       _GotoState344Action,
	{_State241, RparenToken}:                       _GotoState345Action,
	{_State243, NewlinesToken}:                     _GotoState346Action,
	{_State243, OrToken}:                           _GotoState347Action,
	{_State244, CommaToken}:                        _GotoState348Action,
	{_State245, AssignToken}:                       _GotoState349Action,
	{_State245, AddToken}:                          _GotoState249Action,
	{_State245, SubToken}:                          _GotoState251Action,
	{_State245, MulToken}:                          _GotoState250Action,
	{_State245, BinaryTypeOpType}:                  _GotoState252Action,
	{_State245, OptionalDefaultType}:               _GotoState350Action,
	{_State247, IdentifierToken}:                   _GotoState237Action,
	{_State247, UnderscoreToken}:                   _GotoState238Action,
	{_State247, UnsafeToken}:                       _GotoState54Action,
	{_State247, StructToken}:                       _GotoState50Action,
	{_State247, EnumToken}:                         _GotoState110Action,
	{_State247, TraitToken}:                        _GotoState118Action,
	{_State247, FuncToken}:                         _GotoState236Action,
	{_State247, LparenToken}:                       _GotoState114Action,
	{_State247, LbracketToken}:                     _GotoState42Action,
	{_State247, DotToken}:                          _GotoState109Action,
	{_State247, QuestionToken}:                     _GotoState116Action,
	{_State247, ExclaimToken}:                      _GotoState111Action,
	{_State247, TildeTildeToken}:                   _GotoState117Action,
	{_State247, BitNegToken}:                       _GotoState108Action,
	{_State247, BitAndToken}:                       _GotoState107Action,
	{_State247, ParseErrorToken}:                   _GotoState115Action,
	{_State247, UnsafeStatementType}:               _GotoState246Action,
	{_State247, InitializableTypeExprType}:         _GotoState127Action,
	{_State247, SliceTypeExprType}:                 _GotoState101Action,
	{_State247, ArrayTypeExprType}:                 _GotoState60Action,
	{_State247, MapTypeExprType}:                   _GotoState88Action,
	{_State247, AtomTypeExprType}:                  _GotoState120Action,
	{_State247, NamedTypeExprType}:                 _GotoState128Action,
	{_State247, InferredTypeExprType}:              _GotoState126Action,
	{_State247, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State247, ReturnableTypeExprType}:            _GotoState132Action,
	{_State247, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State247, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State247, TypeExprType}:                      _GotoState245Action,
	{_State247, BinaryTypeExprType}:                _GotoState121Action,
	{_State247, FieldDefType}:                      _GotoState300Action,
	{_State247, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State247, ProperExplicitFieldDefsType}:       _GotoState301Action,
	{_State247, ExplicitFieldDefsType}:             _GotoState351Action,
	{_State247, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State247, TraitTypeExprType}:                 _GotoState133Action,
	{_State247, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State247, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State247, FuncTypeExprType}:                  _GotoState123Action,
	{_State247, MethodSignatureType}:               _GotoState242Action,
	{_State252, IdentifierToken}:                   _GotoState113Action,
	{_State252, UnderscoreToken}:                   _GotoState119Action,
	{_State252, StructToken}:                       _GotoState50Action,
	{_State252, EnumToken}:                         _GotoState110Action,
	{_State252, TraitToken}:                        _GotoState118Action,
	{_State252, FuncToken}:                         _GotoState112Action,
	{_State252, LparenToken}:                       _GotoState114Action,
	{_State252, LbracketToken}:                     _GotoState42Action,
	{_State252, DotToken}:                          _GotoState109Action,
	{_State252, QuestionToken}:                     _GotoState116Action,
	{_State252, ExclaimToken}:                      _GotoState111Action,
	{_State252, TildeTildeToken}:                   _GotoState117Action,
	{_State252, BitNegToken}:                       _GotoState108Action,
	{_State252, BitAndToken}:                       _GotoState107Action,
	{_State252, ParseErrorToken}:                   _GotoState115Action,
	{_State252, InitializableTypeExprType}:         _GotoState127Action,
	{_State252, SliceTypeExprType}:                 _GotoState101Action,
	{_State252, ArrayTypeExprType}:                 _GotoState60Action,
	{_State252, MapTypeExprType}:                   _GotoState88Action,
	{_State252, AtomTypeExprType}:                  _GotoState120Action,
	{_State252, NamedTypeExprType}:                 _GotoState128Action,
	{_State252, InferredTypeExprType}:              _GotoState126Action,
	{_State252, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State252, ReturnableTypeExprType}:            _GotoState352Action,
	{_State252, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State252, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State252, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State252, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State252, TraitTypeExprType}:                 _GotoState133Action,
	{_State252, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State252, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State252, FuncTypeExprType}:                  _GotoState123Action,
	{_State253, IntegerLiteralToken}:               _GotoState40Action,
	{_State253, FloatLiteralToken}:                 _GotoState35Action,
	{_State253, RuneLiteralToken}:                  _GotoState48Action,
	{_State253, StringLiteralToken}:                _GotoState49Action,
	{_State253, IdentifierToken}:                   _GotoState38Action,
	{_State253, UnderscoreToken}:                   _GotoState53Action,
	{_State253, TrueToken}:                         _GotoState52Action,
	{_State253, FalseToken}:                        _GotoState34Action,
	{_State253, StructToken}:                       _GotoState50Action,
	{_State253, FuncToken}:                         _GotoState36Action,
	{_State253, VarToken}:                          _GotoState15Action,
	{_State253, LetToken}:                          _GotoState12Action,
	{_State253, NotToken}:                          _GotoState45Action,
	{_State253, LabelDeclToken}:                    _GotoState41Action,
	{_State253, LparenToken}:                       _GotoState43Action,
	{_State253, LbracketToken}:                     _GotoState42Action,
	{_State253, SubToken}:                          _GotoState51Action,
	{_State253, MulToken}:                          _GotoState44Action,
	{_State253, BitNegToken}:                       _GotoState27Action,
	{_State253, BitAndToken}:                       _GotoState26Action,
	{_State253, GreaterToken}:                      _GotoState37Action,
	{_State253, ParseErrorToken}:                   _GotoState46Action,
	{_State253, VarDeclPatternType}:                _GotoState104Action,
	{_State253, VarOrLetType}:                      _GotoState24Action,
	{_State253, ExprType}:                          _GotoState353Action,
	{_State253, OptionalLabelDeclType}:             _GotoState91Action,
	{_State253, SequenceExprType}:                  _GotoState106Action,
	{_State253, CallExprType}:                      _GotoState71Action,
	{_State253, AtomExprType}:                      _GotoState63Action,
	{_State253, ParseErrorExprType}:                _GotoState93Action,
	{_State253, LiteralExprType}:                   _GotoState87Action,
	{_State253, NamedExprType}:                     _GotoState90Action,
	{_State253, BlockExprType}:                     _GotoState70Action,
	{_State253, InitializeExprType}:                _GotoState84Action,
	{_State253, ImplicitStructExprType}:            _GotoState80Action,
	{_State253, AccessibleExprType}:                _GotoState105Action,
	{_State253, AccessExprType}:                    _GotoState55Action,
	{_State253, IndexExprType}:                     _GotoState82Action,
	{_State253, PostfixableExprType}:               _GotoState95Action,
	{_State253, PostfixUnaryExprType}:              _GotoState94Action,
	{_State253, PrefixableExprType}:                _GotoState98Action,
	{_State253, PrefixUnaryExprType}:               _GotoState96Action,
	{_State253, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State253, MulExprType}:                       _GotoState89Action,
	{_State253, BinaryMulExprType}:                 _GotoState67Action,
	{_State253, AddExprType}:                       _GotoState57Action,
	{_State253, BinaryAddExprType}:                 _GotoState64Action,
	{_State253, CmpExprType}:                       _GotoState74Action,
	{_State253, BinaryCmpExprType}:                 _GotoState66Action,
	{_State253, AndExprType}:                       _GotoState58Action,
	{_State253, BinaryAndExprType}:                 _GotoState65Action,
	{_State253, OrExprType}:                        _GotoState92Action,
	{_State253, BinaryOrExprType}:                  _GotoState69Action,
	{_State253, InitializableTypeExprType}:         _GotoState83Action,
	{_State253, SliceTypeExprType}:                 _GotoState101Action,
	{_State253, ArrayTypeExprType}:                 _GotoState60Action,
	{_State253, MapTypeExprType}:                   _GotoState88Action,
	{_State253, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State253, AnonymousFuncExprType}:             _GotoState59Action,
	{_State254, IdentifierToken}:                   _GotoState354Action,
	{_State254, GenericParameterDefType}:           _GotoState355Action,
	{_State254, ProperGenericParameterDefListType}: _GotoState357Action,
	{_State254, GenericParameterDefListType}:       _GotoState356Action,
	{_State255, LparenToken}:                       _GotoState358Action,
	{_State256, IdentifierToken}:                   _GotoState113Action,
	{_State256, UnderscoreToken}:                   _GotoState119Action,
	{_State256, StructToken}:                       _GotoState50Action,
	{_State256, EnumToken}:                         _GotoState110Action,
	{_State256, TraitToken}:                        _GotoState118Action,
	{_State256, FuncToken}:                         _GotoState112Action,
	{_State256, LparenToken}:                       _GotoState114Action,
	{_State256, LbracketToken}:                     _GotoState42Action,
	{_State256, DotToken}:                          _GotoState109Action,
	{_State256, QuestionToken}:                     _GotoState116Action,
	{_State256, ExclaimToken}:                      _GotoState111Action,
	{_State256, EllipsisToken}:                     _GotoState359Action,
	{_State256, TildeTildeToken}:                   _GotoState117Action,
	{_State256, BitNegToken}:                       _GotoState108Action,
	{_State256, BitAndToken}:                       _GotoState107Action,
	{_State256, ParseErrorToken}:                   _GotoState115Action,
	{_State256, InitializableTypeExprType}:         _GotoState127Action,
	{_State256, SliceTypeExprType}:                 _GotoState101Action,
	{_State256, ArrayTypeExprType}:                 _GotoState60Action,
	{_State256, MapTypeExprType}:                   _GotoState88Action,
	{_State256, AtomTypeExprType}:                  _GotoState120Action,
	{_State256, NamedTypeExprType}:                 _GotoState128Action,
	{_State256, InferredTypeExprType}:              _GotoState126Action,
	{_State256, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State256, ReturnableTypeExprType}:            _GotoState132Action,
	{_State256, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State256, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State256, TypeExprType}:                      _GotoState360Action,
	{_State256, BinaryTypeExprType}:                _GotoState121Action,
	{_State256, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State256, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State256, TraitTypeExprType}:                 _GotoState133Action,
	{_State256, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State256, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State256, FuncTypeExprType}:                  _GotoState123Action,
	{_State257, IdentifierToken}:                   _GotoState113Action,
	{_State257, UnderscoreToken}:                   _GotoState119Action,
	{_State257, StructToken}:                       _GotoState50Action,
	{_State257, EnumToken}:                         _GotoState110Action,
	{_State257, TraitToken}:                        _GotoState118Action,
	{_State257, FuncToken}:                         _GotoState112Action,
	{_State257, LparenToken}:                       _GotoState114Action,
	{_State257, LbracketToken}:                     _GotoState42Action,
	{_State257, DotToken}:                          _GotoState109Action,
	{_State257, QuestionToken}:                     _GotoState116Action,
	{_State257, ExclaimToken}:                      _GotoState111Action,
	{_State257, EllipsisToken}:                     _GotoState361Action,
	{_State257, TildeTildeToken}:                   _GotoState117Action,
	{_State257, BitNegToken}:                       _GotoState108Action,
	{_State257, BitAndToken}:                       _GotoState107Action,
	{_State257, ParseErrorToken}:                   _GotoState115Action,
	{_State257, InitializableTypeExprType}:         _GotoState127Action,
	{_State257, SliceTypeExprType}:                 _GotoState101Action,
	{_State257, ArrayTypeExprType}:                 _GotoState60Action,
	{_State257, MapTypeExprType}:                   _GotoState88Action,
	{_State257, AtomTypeExprType}:                  _GotoState120Action,
	{_State257, NamedTypeExprType}:                 _GotoState128Action,
	{_State257, InferredTypeExprType}:              _GotoState126Action,
	{_State257, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State257, ReturnableTypeExprType}:            _GotoState132Action,
	{_State257, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State257, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State257, TypeExprType}:                      _GotoState362Action,
	{_State257, BinaryTypeExprType}:                _GotoState121Action,
	{_State257, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State257, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State257, TraitTypeExprType}:                 _GotoState133Action,
	{_State257, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State257, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State257, FuncTypeExprType}:                  _GotoState123Action,
	{_State258, RparenToken}:                       _GotoState363Action,
	{_State260, IntegerLiteralToken}:               _GotoState40Action,
	{_State260, FloatLiteralToken}:                 _GotoState35Action,
	{_State260, RuneLiteralToken}:                  _GotoState48Action,
	{_State260, StringLiteralToken}:                _GotoState49Action,
	{_State260, IdentifierToken}:                   _GotoState38Action,
	{_State260, UnderscoreToken}:                   _GotoState53Action,
	{_State260, TrueToken}:                         _GotoState52Action,
	{_State260, FalseToken}:                        _GotoState34Action,
	{_State260, CaseToken}:                         _GotoState29Action,
	{_State260, DefaultToken}:                      _GotoState31Action,
	{_State260, ReturnToken}:                       _GotoState47Action,
	{_State260, BreakToken}:                        _GotoState28Action,
	{_State260, ContinueToken}:                     _GotoState30Action,
	{_State260, FallthroughToken}:                  _GotoState33Action,
	{_State260, ImportToken}:                       _GotoState39Action,
	{_State260, UnsafeToken}:                       _GotoState54Action,
	{_State260, StructToken}:                       _GotoState50Action,
	{_State260, FuncToken}:                         _GotoState36Action,
	{_State260, AsyncToken}:                        _GotoState25Action,
	{_State260, DeferToken}:                        _GotoState32Action,
	{_State260, VarToken}:                          _GotoState15Action,
	{_State260, LetToken}:                          _GotoState12Action,
	{_State260, NotToken}:                          _GotoState45Action,
	{_State260, LabelDeclToken}:                    _GotoState41Action,
	{_State260, LparenToken}:                       _GotoState43Action,
	{_State260, LbracketToken}:                     _GotoState42Action,
	{_State260, SubToken}:                          _GotoState51Action,
	{_State260, MulToken}:                          _GotoState44Action,
	{_State260, BitNegToken}:                       _GotoState27Action,
	{_State260, BitAndToken}:                       _GotoState26Action,
	{_State260, GreaterToken}:                      _GotoState37Action,
	{_State260, ParseErrorToken}:                   _GotoState46Action,
	{_State260, SimpleStatementType}:               _GotoState100Action,
	{_State260, StatementType}:                     _GotoState364Action,
	{_State260, ExprOrImproperStructStatementType}: _GotoState77Action,
	{_State260, ExprsType}:                         _GotoState78Action,
	{_State260, CallbackOpType}:                    _GotoState72Action,
	{_State260, CallbackOpStatementType}:           _GotoState73Action,
	{_State260, UnsafeStatementType}:               _GotoState103Action,
	{_State260, JumpStatementType}:                 _GotoState85Action,
	{_State260, JumpTypeType}:                      _GotoState86Action,
	{_State260, FallthroughStatementType}:          _GotoState79Action,
	{_State260, AssignStatementType}:               _GotoState62Action,
	{_State260, UnaryOpAssignStatementType}:        _GotoState102Action,
	{_State260, BinaryOpAssignStatementType}:       _GotoState68Action,
	{_State260, ImportStatementType}:               _GotoState81Action,
	{_State260, VarDeclPatternType}:                _GotoState104Action,
	{_State260, VarOrLetType}:                      _GotoState24Action,
	{_State260, AssignPatternType}:                 _GotoState61Action,
	{_State260, ExprType}:                          _GotoState76Action,
	{_State260, OptionalLabelDeclType}:             _GotoState91Action,
	{_State260, SequenceExprType}:                  _GotoState99Action,
	{_State260, CallExprType}:                      _GotoState71Action,
	{_State260, AtomExprType}:                      _GotoState63Action,
	{_State260, ParseErrorExprType}:                _GotoState93Action,
	{_State260, LiteralExprType}:                   _GotoState87Action,
	{_State260, NamedExprType}:                     _GotoState90Action,
	{_State260, BlockExprType}:                     _GotoState70Action,
	{_State260, InitializeExprType}:                _GotoState84Action,
	{_State260, ImplicitStructExprType}:            _GotoState80Action,
	{_State260, AccessibleExprType}:                _GotoState56Action,
	{_State260, AccessExprType}:                    _GotoState55Action,
	{_State260, IndexExprType}:                     _GotoState82Action,
	{_State260, PostfixableExprType}:               _GotoState95Action,
	{_State260, PostfixUnaryExprType}:              _GotoState94Action,
	{_State260, PrefixableExprType}:                _GotoState98Action,
	{_State260, PrefixUnaryExprType}:               _GotoState96Action,
	{_State260, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State260, MulExprType}:                       _GotoState89Action,
	{_State260, BinaryMulExprType}:                 _GotoState67Action,
	{_State260, AddExprType}:                       _GotoState57Action,
	{_State260, BinaryAddExprType}:                 _GotoState64Action,
	{_State260, CmpExprType}:                       _GotoState74Action,
	{_State260, BinaryCmpExprType}:                 _GotoState66Action,
	{_State260, AndExprType}:                       _GotoState58Action,
	{_State260, BinaryAndExprType}:                 _GotoState65Action,
	{_State260, OrExprType}:                        _GotoState92Action,
	{_State260, BinaryOrExprType}:                  _GotoState69Action,
	{_State260, InitializableTypeExprType}:         _GotoState83Action,
	{_State260, SliceTypeExprType}:                 _GotoState101Action,
	{_State260, ArrayTypeExprType}:                 _GotoState60Action,
	{_State260, MapTypeExprType}:                   _GotoState88Action,
	{_State260, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State260, AnonymousFuncExprType}:             _GotoState59Action,
	{_State261, IntegerLiteralToken}:               _GotoState40Action,
	{_State261, FloatLiteralToken}:                 _GotoState35Action,
	{_State261, RuneLiteralToken}:                  _GotoState48Action,
	{_State261, StringLiteralToken}:                _GotoState49Action,
	{_State261, IdentifierToken}:                   _GotoState38Action,
	{_State261, UnderscoreToken}:                   _GotoState53Action,
	{_State261, TrueToken}:                         _GotoState52Action,
	{_State261, FalseToken}:                        _GotoState34Action,
	{_State261, CaseToken}:                         _GotoState29Action,
	{_State261, DefaultToken}:                      _GotoState31Action,
	{_State261, ReturnToken}:                       _GotoState47Action,
	{_State261, BreakToken}:                        _GotoState28Action,
	{_State261, ContinueToken}:                     _GotoState30Action,
	{_State261, FallthroughToken}:                  _GotoState33Action,
	{_State261, ImportToken}:                       _GotoState39Action,
	{_State261, UnsafeToken}:                       _GotoState54Action,
	{_State261, StructToken}:                       _GotoState50Action,
	{_State261, FuncToken}:                         _GotoState36Action,
	{_State261, AsyncToken}:                        _GotoState25Action,
	{_State261, DeferToken}:                        _GotoState32Action,
	{_State261, VarToken}:                          _GotoState15Action,
	{_State261, LetToken}:                          _GotoState12Action,
	{_State261, NotToken}:                          _GotoState45Action,
	{_State261, LabelDeclToken}:                    _GotoState41Action,
	{_State261, LparenToken}:                       _GotoState43Action,
	{_State261, LbracketToken}:                     _GotoState42Action,
	{_State261, SubToken}:                          _GotoState51Action,
	{_State261, MulToken}:                          _GotoState44Action,
	{_State261, BitNegToken}:                       _GotoState27Action,
	{_State261, BitAndToken}:                       _GotoState26Action,
	{_State261, GreaterToken}:                      _GotoState37Action,
	{_State261, ParseErrorToken}:                   _GotoState46Action,
	{_State261, SimpleStatementType}:               _GotoState100Action,
	{_State261, StatementType}:                     _GotoState365Action,
	{_State261, ExprOrImproperStructStatementType}: _GotoState77Action,
	{_State261, ExprsType}:                         _GotoState78Action,
	{_State261, CallbackOpType}:                    _GotoState72Action,
	{_State261, CallbackOpStatementType}:           _GotoState73Action,
	{_State261, UnsafeStatementType}:               _GotoState103Action,
	{_State261, JumpStatementType}:                 _GotoState85Action,
	{_State261, JumpTypeType}:                      _GotoState86Action,
	{_State261, FallthroughStatementType}:          _GotoState79Action,
	{_State261, AssignStatementType}:               _GotoState62Action,
	{_State261, UnaryOpAssignStatementType}:        _GotoState102Action,
	{_State261, BinaryOpAssignStatementType}:       _GotoState68Action,
	{_State261, ImportStatementType}:               _GotoState81Action,
	{_State261, VarDeclPatternType}:                _GotoState104Action,
	{_State261, VarOrLetType}:                      _GotoState24Action,
	{_State261, AssignPatternType}:                 _GotoState61Action,
	{_State261, ExprType}:                          _GotoState76Action,
	{_State261, OptionalLabelDeclType}:             _GotoState91Action,
	{_State261, SequenceExprType}:                  _GotoState99Action,
	{_State261, CallExprType}:                      _GotoState71Action,
	{_State261, AtomExprType}:                      _GotoState63Action,
	{_State261, ParseErrorExprType}:                _GotoState93Action,
	{_State261, LiteralExprType}:                   _GotoState87Action,
	{_State261, NamedExprType}:                     _GotoState90Action,
	{_State261, BlockExprType}:                     _GotoState70Action,
	{_State261, InitializeExprType}:                _GotoState84Action,
	{_State261, ImplicitStructExprType}:            _GotoState80Action,
	{_State261, AccessibleExprType}:                _GotoState56Action,
	{_State261, AccessExprType}:                    _GotoState55Action,
	{_State261, IndexExprType}:                     _GotoState82Action,
	{_State261, PostfixableExprType}:               _GotoState95Action,
	{_State261, PostfixUnaryExprType}:              _GotoState94Action,
	{_State261, PrefixableExprType}:                _GotoState98Action,
	{_State261, PrefixUnaryExprType}:               _GotoState96Action,
	{_State261, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State261, MulExprType}:                       _GotoState89Action,
	{_State261, BinaryMulExprType}:                 _GotoState67Action,
	{_State261, AddExprType}:                       _GotoState57Action,
	{_State261, BinaryAddExprType}:                 _GotoState64Action,
	{_State261, CmpExprType}:                       _GotoState74Action,
	{_State261, BinaryCmpExprType}:                 _GotoState66Action,
	{_State261, AndExprType}:                       _GotoState58Action,
	{_State261, BinaryAndExprType}:                 _GotoState65Action,
	{_State261, OrExprType}:                        _GotoState92Action,
	{_State261, BinaryOrExprType}:                  _GotoState69Action,
	{_State261, InitializableTypeExprType}:         _GotoState83Action,
	{_State261, SliceTypeExprType}:                 _GotoState101Action,
	{_State261, ArrayTypeExprType}:                 _GotoState60Action,
	{_State261, MapTypeExprType}:                   _GotoState88Action,
	{_State261, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State261, AnonymousFuncExprType}:             _GotoState59Action,
	{_State263, IdentifierToken}:                   _GotoState113Action,
	{_State263, UnderscoreToken}:                   _GotoState119Action,
	{_State263, StructToken}:                       _GotoState50Action,
	{_State263, EnumToken}:                         _GotoState110Action,
	{_State263, TraitToken}:                        _GotoState118Action,
	{_State263, FuncToken}:                         _GotoState112Action,
	{_State263, LparenToken}:                       _GotoState114Action,
	{_State263, LbracketToken}:                     _GotoState42Action,
	{_State263, DotToken}:                          _GotoState109Action,
	{_State263, QuestionToken}:                     _GotoState116Action,
	{_State263, ExclaimToken}:                      _GotoState111Action,
	{_State263, TildeTildeToken}:                   _GotoState117Action,
	{_State263, BitNegToken}:                       _GotoState108Action,
	{_State263, BitAndToken}:                       _GotoState107Action,
	{_State263, ParseErrorToken}:                   _GotoState115Action,
	{_State263, InitializableTypeExprType}:         _GotoState127Action,
	{_State263, SliceTypeExprType}:                 _GotoState101Action,
	{_State263, ArrayTypeExprType}:                 _GotoState60Action,
	{_State263, MapTypeExprType}:                   _GotoState88Action,
	{_State263, AtomTypeExprType}:                  _GotoState120Action,
	{_State263, NamedTypeExprType}:                 _GotoState128Action,
	{_State263, InferredTypeExprType}:              _GotoState126Action,
	{_State263, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State263, ReturnableTypeExprType}:            _GotoState132Action,
	{_State263, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State263, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State263, TypeExprType}:                      _GotoState366Action,
	{_State263, BinaryTypeExprType}:                _GotoState121Action,
	{_State263, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State263, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State263, TraitTypeExprType}:                 _GotoState133Action,
	{_State263, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State263, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State263, FuncTypeExprType}:                  _GotoState123Action,
	{_State264, IdentifierToken}:                   _GotoState113Action,
	{_State264, UnderscoreToken}:                   _GotoState119Action,
	{_State264, StructToken}:                       _GotoState50Action,
	{_State264, EnumToken}:                         _GotoState110Action,
	{_State264, TraitToken}:                        _GotoState118Action,
	{_State264, FuncToken}:                         _GotoState112Action,
	{_State264, LparenToken}:                       _GotoState114Action,
	{_State264, LbracketToken}:                     _GotoState42Action,
	{_State264, DotToken}:                          _GotoState109Action,
	{_State264, QuestionToken}:                     _GotoState116Action,
	{_State264, ExclaimToken}:                      _GotoState111Action,
	{_State264, TildeTildeToken}:                   _GotoState117Action,
	{_State264, BitNegToken}:                       _GotoState108Action,
	{_State264, BitAndToken}:                       _GotoState107Action,
	{_State264, ParseErrorToken}:                   _GotoState115Action,
	{_State264, InitializableTypeExprType}:         _GotoState127Action,
	{_State264, SliceTypeExprType}:                 _GotoState101Action,
	{_State264, ArrayTypeExprType}:                 _GotoState60Action,
	{_State264, MapTypeExprType}:                   _GotoState88Action,
	{_State264, AtomTypeExprType}:                  _GotoState120Action,
	{_State264, NamedTypeExprType}:                 _GotoState128Action,
	{_State264, InferredTypeExprType}:              _GotoState126Action,
	{_State264, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State264, ReturnableTypeExprType}:            _GotoState132Action,
	{_State264, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State264, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State264, TypeExprType}:                      _GotoState367Action,
	{_State264, BinaryTypeExprType}:                _GotoState121Action,
	{_State264, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State264, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State264, TraitTypeExprType}:                 _GotoState133Action,
	{_State264, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State264, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State264, FuncTypeExprType}:                  _GotoState123Action,
	{_State268, AssignToken}:                       _GotoState368Action,
	{_State270, RparenToken}:                       _GotoState370Action,
	{_State270, CommaToken}:                        _GotoState369Action,
	{_State273, AddToken}:                          _GotoState249Action,
	{_State273, SubToken}:                          _GotoState251Action,
	{_State273, MulToken}:                          _GotoState250Action,
	{_State273, BinaryTypeOpType}:                  _GotoState252Action,
	{_State274, LparenToken}:                       _GotoState43Action,
	{_State274, ImplicitStructExprType}:            _GotoState371Action,
	{_State275, IdentifierToken}:                   _GotoState372Action,
	{_State276, IntegerLiteralToken}:               _GotoState40Action,
	{_State276, FloatLiteralToken}:                 _GotoState35Action,
	{_State276, RuneLiteralToken}:                  _GotoState48Action,
	{_State276, StringLiteralToken}:                _GotoState49Action,
	{_State276, IdentifierToken}:                   _GotoState38Action,
	{_State276, UnderscoreToken}:                   _GotoState53Action,
	{_State276, TrueToken}:                         _GotoState52Action,
	{_State276, FalseToken}:                        _GotoState34Action,
	{_State276, ReturnToken}:                       _GotoState47Action,
	{_State276, BreakToken}:                        _GotoState28Action,
	{_State276, ContinueToken}:                     _GotoState30Action,
	{_State276, FallthroughToken}:                  _GotoState33Action,
	{_State276, UnsafeToken}:                       _GotoState54Action,
	{_State276, StructToken}:                       _GotoState50Action,
	{_State276, FuncToken}:                         _GotoState36Action,
	{_State276, AsyncToken}:                        _GotoState25Action,
	{_State276, DeferToken}:                        _GotoState32Action,
	{_State276, VarToken}:                          _GotoState15Action,
	{_State276, LetToken}:                          _GotoState12Action,
	{_State276, NotToken}:                          _GotoState45Action,
	{_State276, LabelDeclToken}:                    _GotoState41Action,
	{_State276, LparenToken}:                       _GotoState43Action,
	{_State276, LbracketToken}:                     _GotoState42Action,
	{_State276, SubToken}:                          _GotoState51Action,
	{_State276, MulToken}:                          _GotoState44Action,
	{_State276, BitNegToken}:                       _GotoState27Action,
	{_State276, BitAndToken}:                       _GotoState26Action,
	{_State276, GreaterToken}:                      _GotoState37Action,
	{_State276, ParseErrorToken}:                   _GotoState46Action,
	{_State276, SimpleStatementType}:               _GotoState279Action,
	{_State276, OptionalSimpleStatementType}:       _GotoState373Action,
	{_State276, ExprOrImproperStructStatementType}: _GotoState77Action,
	{_State276, ExprsType}:                         _GotoState78Action,
	{_State276, CallbackOpType}:                    _GotoState72Action,
	{_State276, CallbackOpStatementType}:           _GotoState73Action,
	{_State276, UnsafeStatementType}:               _GotoState103Action,
	{_State276, JumpStatementType}:                 _GotoState85Action,
	{_State276, JumpTypeType}:                      _GotoState86Action,
	{_State276, FallthroughStatementType}:          _GotoState79Action,
	{_State276, AssignStatementType}:               _GotoState62Action,
	{_State276, UnaryOpAssignStatementType}:        _GotoState102Action,
	{_State276, BinaryOpAssignStatementType}:       _GotoState68Action,
	{_State276, VarDeclPatternType}:                _GotoState104Action,
	{_State276, VarOrLetType}:                      _GotoState24Action,
	{_State276, AssignPatternType}:                 _GotoState61Action,
	{_State276, ExprType}:                          _GotoState76Action,
	{_State276, OptionalLabelDeclType}:             _GotoState91Action,
	{_State276, SequenceExprType}:                  _GotoState99Action,
	{_State276, CallExprType}:                      _GotoState71Action,
	{_State276, AtomExprType}:                      _GotoState63Action,
	{_State276, ParseErrorExprType}:                _GotoState93Action,
	{_State276, LiteralExprType}:                   _GotoState87Action,
	{_State276, NamedExprType}:                     _GotoState90Action,
	{_State276, BlockExprType}:                     _GotoState70Action,
	{_State276, InitializeExprType}:                _GotoState84Action,
	{_State276, ImplicitStructExprType}:            _GotoState80Action,
	{_State276, AccessibleExprType}:                _GotoState56Action,
	{_State276, AccessExprType}:                    _GotoState55Action,
	{_State276, IndexExprType}:                     _GotoState82Action,
	{_State276, PostfixableExprType}:               _GotoState95Action,
	{_State276, PostfixUnaryExprType}:              _GotoState94Action,
	{_State276, PrefixableExprType}:                _GotoState98Action,
	{_State276, PrefixUnaryExprType}:               _GotoState96Action,
	{_State276, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State276, MulExprType}:                       _GotoState89Action,
	{_State276, BinaryMulExprType}:                 _GotoState67Action,
	{_State276, AddExprType}:                       _GotoState57Action,
	{_State276, BinaryAddExprType}:                 _GotoState64Action,
	{_State276, CmpExprType}:                       _GotoState74Action,
	{_State276, BinaryCmpExprType}:                 _GotoState66Action,
	{_State276, AndExprType}:                       _GotoState58Action,
	{_State276, BinaryAndExprType}:                 _GotoState65Action,
	{_State276, OrExprType}:                        _GotoState92Action,
	{_State276, BinaryOrExprType}:                  _GotoState69Action,
	{_State276, InitializableTypeExprType}:         _GotoState83Action,
	{_State276, SliceTypeExprType}:                 _GotoState101Action,
	{_State276, ArrayTypeExprType}:                 _GotoState60Action,
	{_State276, MapTypeExprType}:                   _GotoState88Action,
	{_State276, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State276, AnonymousFuncExprType}:             _GotoState59Action,
	{_State277, IntegerLiteralToken}:               _GotoState40Action,
	{_State277, FloatLiteralToken}:                 _GotoState35Action,
	{_State277, RuneLiteralToken}:                  _GotoState48Action,
	{_State277, StringLiteralToken}:                _GotoState49Action,
	{_State277, IdentifierToken}:                   _GotoState38Action,
	{_State277, UnderscoreToken}:                   _GotoState53Action,
	{_State277, TrueToken}:                         _GotoState52Action,
	{_State277, FalseToken}:                        _GotoState34Action,
	{_State277, StructToken}:                       _GotoState50Action,
	{_State277, FuncToken}:                         _GotoState36Action,
	{_State277, VarToken}:                          _GotoState149Action,
	{_State277, LetToken}:                          _GotoState12Action,
	{_State277, NotToken}:                          _GotoState45Action,
	{_State277, LabelDeclToken}:                    _GotoState41Action,
	{_State277, LparenToken}:                       _GotoState43Action,
	{_State277, LbracketToken}:                     _GotoState42Action,
	{_State277, DotToken}:                          _GotoState148Action,
	{_State277, SubToken}:                          _GotoState51Action,
	{_State277, MulToken}:                          _GotoState44Action,
	{_State277, BitNegToken}:                       _GotoState27Action,
	{_State277, BitAndToken}:                       _GotoState26Action,
	{_State277, GreaterToken}:                      _GotoState37Action,
	{_State277, ParseErrorToken}:                   _GotoState46Action,
	{_State277, VarDeclPatternType}:                _GotoState104Action,
	{_State277, VarOrLetType}:                      _GotoState24Action,
	{_State277, AssignPatternType}:                 _GotoState150Action,
	{_State277, CasePatternType}:                   _GotoState374Action,
	{_State277, OptionalLabelDeclType}:             _GotoState153Action,
	{_State277, SequenceExprType}:                  _GotoState154Action,
	{_State277, CallExprType}:                      _GotoState71Action,
	{_State277, AtomExprType}:                      _GotoState63Action,
	{_State277, ParseErrorExprType}:                _GotoState93Action,
	{_State277, LiteralExprType}:                   _GotoState87Action,
	{_State277, NamedExprType}:                     _GotoState90Action,
	{_State277, BlockExprType}:                     _GotoState70Action,
	{_State277, InitializeExprType}:                _GotoState84Action,
	{_State277, ImplicitStructExprType}:            _GotoState80Action,
	{_State277, AccessibleExprType}:                _GotoState105Action,
	{_State277, AccessExprType}:                    _GotoState55Action,
	{_State277, IndexExprType}:                     _GotoState82Action,
	{_State277, PostfixableExprType}:               _GotoState95Action,
	{_State277, PostfixUnaryExprType}:              _GotoState94Action,
	{_State277, PrefixableExprType}:                _GotoState98Action,
	{_State277, PrefixUnaryExprType}:               _GotoState96Action,
	{_State277, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State277, MulExprType}:                       _GotoState89Action,
	{_State277, BinaryMulExprType}:                 _GotoState67Action,
	{_State277, AddExprType}:                       _GotoState57Action,
	{_State277, BinaryAddExprType}:                 _GotoState64Action,
	{_State277, CmpExprType}:                       _GotoState74Action,
	{_State277, BinaryCmpExprType}:                 _GotoState66Action,
	{_State277, AndExprType}:                       _GotoState58Action,
	{_State277, BinaryAndExprType}:                 _GotoState65Action,
	{_State277, OrExprType}:                        _GotoState92Action,
	{_State277, BinaryOrExprType}:                  _GotoState69Action,
	{_State277, InitializableTypeExprType}:         _GotoState83Action,
	{_State277, SliceTypeExprType}:                 _GotoState101Action,
	{_State277, ArrayTypeExprType}:                 _GotoState60Action,
	{_State277, MapTypeExprType}:                   _GotoState88Action,
	{_State277, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State277, AnonymousFuncExprType}:             _GotoState59Action,
	{_State281, RparenToken}:                       _GotoState375Action,
	{_State282, CommaToken}:                        _GotoState376Action,
	{_State286, RparenToken}:                       _GotoState377Action,
	{_State287, NewlinesToken}:                     _GotoState379Action,
	{_State287, CommaToken}:                        _GotoState378Action,
	{_State289, IdentifierToken}:                   _GotoState113Action,
	{_State289, UnderscoreToken}:                   _GotoState119Action,
	{_State289, StructToken}:                       _GotoState50Action,
	{_State289, EnumToken}:                         _GotoState110Action,
	{_State289, TraitToken}:                        _GotoState118Action,
	{_State289, FuncToken}:                         _GotoState112Action,
	{_State289, LparenToken}:                       _GotoState114Action,
	{_State289, LbracketToken}:                     _GotoState42Action,
	{_State289, DotToken}:                          _GotoState109Action,
	{_State289, QuestionToken}:                     _GotoState116Action,
	{_State289, ExclaimToken}:                      _GotoState111Action,
	{_State289, TildeTildeToken}:                   _GotoState117Action,
	{_State289, BitNegToken}:                       _GotoState108Action,
	{_State289, BitAndToken}:                       _GotoState107Action,
	{_State289, ParseErrorToken}:                   _GotoState115Action,
	{_State289, InitializableTypeExprType}:         _GotoState127Action,
	{_State289, SliceTypeExprType}:                 _GotoState101Action,
	{_State289, ArrayTypeExprType}:                 _GotoState60Action,
	{_State289, MapTypeExprType}:                   _GotoState88Action,
	{_State289, AtomTypeExprType}:                  _GotoState120Action,
	{_State289, NamedTypeExprType}:                 _GotoState128Action,
	{_State289, InferredTypeExprType}:              _GotoState126Action,
	{_State289, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State289, ReturnableTypeExprType}:            _GotoState132Action,
	{_State289, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State289, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State289, TypeExprType}:                      _GotoState380Action,
	{_State289, BinaryTypeExprType}:                _GotoState121Action,
	{_State289, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State289, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State289, TraitTypeExprType}:                 _GotoState133Action,
	{_State289, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State289, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State289, FuncTypeExprType}:                  _GotoState123Action,
	{_State290, IntegerLiteralToken}:               _GotoState381Action,
	{_State293, IntegerLiteralToken}:               _GotoState40Action,
	{_State293, FloatLiteralToken}:                 _GotoState35Action,
	{_State293, RuneLiteralToken}:                  _GotoState48Action,
	{_State293, StringLiteralToken}:                _GotoState49Action,
	{_State293, IdentifierToken}:                   _GotoState38Action,
	{_State293, UnderscoreToken}:                   _GotoState53Action,
	{_State293, TrueToken}:                         _GotoState52Action,
	{_State293, FalseToken}:                        _GotoState34Action,
	{_State293, StructToken}:                       _GotoState50Action,
	{_State293, FuncToken}:                         _GotoState36Action,
	{_State293, VarToken}:                          _GotoState15Action,
	{_State293, LetToken}:                          _GotoState12Action,
	{_State293, NotToken}:                          _GotoState45Action,
	{_State293, LabelDeclToken}:                    _GotoState41Action,
	{_State293, LparenToken}:                       _GotoState43Action,
	{_State293, LbracketToken}:                     _GotoState42Action,
	{_State293, SubToken}:                          _GotoState51Action,
	{_State293, MulToken}:                          _GotoState44Action,
	{_State293, BitNegToken}:                       _GotoState27Action,
	{_State293, BitAndToken}:                       _GotoState26Action,
	{_State293, GreaterToken}:                      _GotoState37Action,
	{_State293, ParseErrorToken}:                   _GotoState46Action,
	{_State293, VarDeclPatternType}:                _GotoState104Action,
	{_State293, VarOrLetType}:                      _GotoState24Action,
	{_State293, ExprType}:                          _GotoState382Action,
	{_State293, OptionalLabelDeclType}:             _GotoState91Action,
	{_State293, SequenceExprType}:                  _GotoState106Action,
	{_State293, CallExprType}:                      _GotoState71Action,
	{_State293, AtomExprType}:                      _GotoState63Action,
	{_State293, ParseErrorExprType}:                _GotoState93Action,
	{_State293, LiteralExprType}:                   _GotoState87Action,
	{_State293, NamedExprType}:                     _GotoState90Action,
	{_State293, BlockExprType}:                     _GotoState70Action,
	{_State293, InitializeExprType}:                _GotoState84Action,
	{_State293, ImplicitStructExprType}:            _GotoState80Action,
	{_State293, AccessibleExprType}:                _GotoState105Action,
	{_State293, AccessExprType}:                    _GotoState55Action,
	{_State293, IndexExprType}:                     _GotoState82Action,
	{_State293, PostfixableExprType}:               _GotoState95Action,
	{_State293, PostfixUnaryExprType}:              _GotoState94Action,
	{_State293, PrefixableExprType}:                _GotoState98Action,
	{_State293, PrefixUnaryExprType}:               _GotoState96Action,
	{_State293, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State293, MulExprType}:                       _GotoState89Action,
	{_State293, BinaryMulExprType}:                 _GotoState67Action,
	{_State293, AddExprType}:                       _GotoState57Action,
	{_State293, BinaryAddExprType}:                 _GotoState64Action,
	{_State293, CmpExprType}:                       _GotoState74Action,
	{_State293, BinaryCmpExprType}:                 _GotoState66Action,
	{_State293, AndExprType}:                       _GotoState58Action,
	{_State293, BinaryAndExprType}:                 _GotoState65Action,
	{_State293, OrExprType}:                        _GotoState92Action,
	{_State293, BinaryOrExprType}:                  _GotoState69Action,
	{_State293, InitializableTypeExprType}:         _GotoState83Action,
	{_State293, SliceTypeExprType}:                 _GotoState101Action,
	{_State293, ArrayTypeExprType}:                 _GotoState60Action,
	{_State293, MapTypeExprType}:                   _GotoState88Action,
	{_State293, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State293, AnonymousFuncExprType}:             _GotoState59Action,
	{_State295, IntegerLiteralToken}:               _GotoState40Action,
	{_State295, FloatLiteralToken}:                 _GotoState35Action,
	{_State295, RuneLiteralToken}:                  _GotoState48Action,
	{_State295, StringLiteralToken}:                _GotoState49Action,
	{_State295, IdentifierToken}:                   _GotoState38Action,
	{_State295, UnderscoreToken}:                   _GotoState53Action,
	{_State295, TrueToken}:                         _GotoState52Action,
	{_State295, FalseToken}:                        _GotoState34Action,
	{_State295, StructToken}:                       _GotoState50Action,
	{_State295, FuncToken}:                         _GotoState36Action,
	{_State295, VarToken}:                          _GotoState15Action,
	{_State295, LetToken}:                          _GotoState12Action,
	{_State295, NotToken}:                          _GotoState45Action,
	{_State295, LabelDeclToken}:                    _GotoState41Action,
	{_State295, LparenToken}:                       _GotoState43Action,
	{_State295, LbracketToken}:                     _GotoState42Action,
	{_State295, SubToken}:                          _GotoState51Action,
	{_State295, MulToken}:                          _GotoState44Action,
	{_State295, BitNegToken}:                       _GotoState27Action,
	{_State295, BitAndToken}:                       _GotoState26Action,
	{_State295, GreaterToken}:                      _GotoState37Action,
	{_State295, ParseErrorToken}:                   _GotoState46Action,
	{_State295, VarDeclPatternType}:                _GotoState104Action,
	{_State295, VarOrLetType}:                      _GotoState24Action,
	{_State295, ExprType}:                          _GotoState383Action,
	{_State295, OptionalLabelDeclType}:             _GotoState91Action,
	{_State295, SequenceExprType}:                  _GotoState106Action,
	{_State295, CallExprType}:                      _GotoState71Action,
	{_State295, AtomExprType}:                      _GotoState63Action,
	{_State295, ParseErrorExprType}:                _GotoState93Action,
	{_State295, LiteralExprType}:                   _GotoState87Action,
	{_State295, NamedExprType}:                     _GotoState90Action,
	{_State295, BlockExprType}:                     _GotoState70Action,
	{_State295, InitializeExprType}:                _GotoState84Action,
	{_State295, ImplicitStructExprType}:            _GotoState80Action,
	{_State295, AccessibleExprType}:                _GotoState105Action,
	{_State295, AccessExprType}:                    _GotoState55Action,
	{_State295, IndexExprType}:                     _GotoState82Action,
	{_State295, PostfixableExprType}:               _GotoState95Action,
	{_State295, PostfixUnaryExprType}:              _GotoState94Action,
	{_State295, PrefixableExprType}:                _GotoState98Action,
	{_State295, PrefixUnaryExprType}:               _GotoState96Action,
	{_State295, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State295, MulExprType}:                       _GotoState89Action,
	{_State295, BinaryMulExprType}:                 _GotoState67Action,
	{_State295, AddExprType}:                       _GotoState57Action,
	{_State295, BinaryAddExprType}:                 _GotoState64Action,
	{_State295, CmpExprType}:                       _GotoState74Action,
	{_State295, BinaryCmpExprType}:                 _GotoState66Action,
	{_State295, AndExprType}:                       _GotoState58Action,
	{_State295, BinaryAndExprType}:                 _GotoState65Action,
	{_State295, OrExprType}:                        _GotoState92Action,
	{_State295, BinaryOrExprType}:                  _GotoState69Action,
	{_State295, InitializableTypeExprType}:         _GotoState83Action,
	{_State295, SliceTypeExprType}:                 _GotoState101Action,
	{_State295, ArrayTypeExprType}:                 _GotoState60Action,
	{_State295, MapTypeExprType}:                   _GotoState88Action,
	{_State295, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State295, AnonymousFuncExprType}:             _GotoState59Action,
	{_State296, IntegerLiteralToken}:               _GotoState40Action,
	{_State296, FloatLiteralToken}:                 _GotoState35Action,
	{_State296, RuneLiteralToken}:                  _GotoState48Action,
	{_State296, StringLiteralToken}:                _GotoState49Action,
	{_State296, IdentifierToken}:                   _GotoState38Action,
	{_State296, UnderscoreToken}:                   _GotoState53Action,
	{_State296, TrueToken}:                         _GotoState52Action,
	{_State296, FalseToken}:                        _GotoState34Action,
	{_State296, StructToken}:                       _GotoState50Action,
	{_State296, FuncToken}:                         _GotoState36Action,
	{_State296, VarToken}:                          _GotoState15Action,
	{_State296, LetToken}:                          _GotoState12Action,
	{_State296, NotToken}:                          _GotoState45Action,
	{_State296, LabelDeclToken}:                    _GotoState41Action,
	{_State296, LparenToken}:                       _GotoState43Action,
	{_State296, LbracketToken}:                     _GotoState42Action,
	{_State296, SubToken}:                          _GotoState51Action,
	{_State296, MulToken}:                          _GotoState44Action,
	{_State296, BitNegToken}:                       _GotoState27Action,
	{_State296, BitAndToken}:                       _GotoState26Action,
	{_State296, GreaterToken}:                      _GotoState37Action,
	{_State296, ParseErrorToken}:                   _GotoState46Action,
	{_State296, VarDeclPatternType}:                _GotoState104Action,
	{_State296, VarOrLetType}:                      _GotoState24Action,
	{_State296, ExprType}:                          _GotoState384Action,
	{_State296, OptionalLabelDeclType}:             _GotoState91Action,
	{_State296, SequenceExprType}:                  _GotoState106Action,
	{_State296, CallExprType}:                      _GotoState71Action,
	{_State296, AtomExprType}:                      _GotoState63Action,
	{_State296, ParseErrorExprType}:                _GotoState93Action,
	{_State296, LiteralExprType}:                   _GotoState87Action,
	{_State296, NamedExprType}:                     _GotoState90Action,
	{_State296, BlockExprType}:                     _GotoState70Action,
	{_State296, InitializeExprType}:                _GotoState84Action,
	{_State296, ImplicitStructExprType}:            _GotoState80Action,
	{_State296, AccessibleExprType}:                _GotoState105Action,
	{_State296, AccessExprType}:                    _GotoState55Action,
	{_State296, IndexExprType}:                     _GotoState82Action,
	{_State296, PostfixableExprType}:               _GotoState95Action,
	{_State296, PostfixUnaryExprType}:              _GotoState94Action,
	{_State296, PrefixableExprType}:                _GotoState98Action,
	{_State296, PrefixUnaryExprType}:               _GotoState96Action,
	{_State296, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State296, MulExprType}:                       _GotoState89Action,
	{_State296, BinaryMulExprType}:                 _GotoState67Action,
	{_State296, AddExprType}:                       _GotoState57Action,
	{_State296, BinaryAddExprType}:                 _GotoState64Action,
	{_State296, CmpExprType}:                       _GotoState74Action,
	{_State296, BinaryCmpExprType}:                 _GotoState66Action,
	{_State296, AndExprType}:                       _GotoState58Action,
	{_State296, BinaryAndExprType}:                 _GotoState65Action,
	{_State296, OrExprType}:                        _GotoState92Action,
	{_State296, BinaryOrExprType}:                  _GotoState69Action,
	{_State296, InitializableTypeExprType}:         _GotoState83Action,
	{_State296, SliceTypeExprType}:                 _GotoState101Action,
	{_State296, ArrayTypeExprType}:                 _GotoState60Action,
	{_State296, MapTypeExprType}:                   _GotoState88Action,
	{_State296, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State296, AnonymousFuncExprType}:             _GotoState59Action,
	{_State298, IntegerLiteralToken}:               _GotoState40Action,
	{_State298, FloatLiteralToken}:                 _GotoState35Action,
	{_State298, RuneLiteralToken}:                  _GotoState48Action,
	{_State298, StringLiteralToken}:                _GotoState49Action,
	{_State298, IdentifierToken}:                   _GotoState167Action,
	{_State298, UnderscoreToken}:                   _GotoState53Action,
	{_State298, TrueToken}:                         _GotoState52Action,
	{_State298, FalseToken}:                        _GotoState34Action,
	{_State298, StructToken}:                       _GotoState50Action,
	{_State298, FuncToken}:                         _GotoState36Action,
	{_State298, VarToken}:                          _GotoState15Action,
	{_State298, LetToken}:                          _GotoState12Action,
	{_State298, NotToken}:                          _GotoState45Action,
	{_State298, LabelDeclToken}:                    _GotoState41Action,
	{_State298, LparenToken}:                       _GotoState43Action,
	{_State298, LbracketToken}:                     _GotoState42Action,
	{_State298, ColonToken}:                        _GotoState165Action,
	{_State298, EllipsisToken}:                     _GotoState166Action,
	{_State298, SubToken}:                          _GotoState51Action,
	{_State298, MulToken}:                          _GotoState44Action,
	{_State298, BitNegToken}:                       _GotoState27Action,
	{_State298, BitAndToken}:                       _GotoState26Action,
	{_State298, GreaterToken}:                      _GotoState37Action,
	{_State298, ParseErrorToken}:                   _GotoState46Action,
	{_State298, VarDeclPatternType}:                _GotoState104Action,
	{_State298, VarOrLetType}:                      _GotoState24Action,
	{_State298, ExprType}:                          _GotoState171Action,
	{_State298, OptionalLabelDeclType}:             _GotoState91Action,
	{_State298, SequenceExprType}:                  _GotoState106Action,
	{_State298, CallExprType}:                      _GotoState71Action,
	{_State298, ArgumentType}:                      _GotoState385Action,
	{_State298, ColonExprType}:                     _GotoState170Action,
	{_State298, AtomExprType}:                      _GotoState63Action,
	{_State298, ParseErrorExprType}:                _GotoState93Action,
	{_State298, LiteralExprType}:                   _GotoState87Action,
	{_State298, NamedExprType}:                     _GotoState90Action,
	{_State298, BlockExprType}:                     _GotoState70Action,
	{_State298, InitializeExprType}:                _GotoState84Action,
	{_State298, ImplicitStructExprType}:            _GotoState80Action,
	{_State298, AccessibleExprType}:                _GotoState105Action,
	{_State298, AccessExprType}:                    _GotoState55Action,
	{_State298, IndexExprType}:                     _GotoState82Action,
	{_State298, PostfixableExprType}:               _GotoState95Action,
	{_State298, PostfixUnaryExprType}:              _GotoState94Action,
	{_State298, PrefixableExprType}:                _GotoState98Action,
	{_State298, PrefixUnaryExprType}:               _GotoState96Action,
	{_State298, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State298, MulExprType}:                       _GotoState89Action,
	{_State298, BinaryMulExprType}:                 _GotoState67Action,
	{_State298, AddExprType}:                       _GotoState57Action,
	{_State298, BinaryAddExprType}:                 _GotoState64Action,
	{_State298, CmpExprType}:                       _GotoState74Action,
	{_State298, BinaryCmpExprType}:                 _GotoState66Action,
	{_State298, AndExprType}:                       _GotoState58Action,
	{_State298, BinaryAndExprType}:                 _GotoState65Action,
	{_State298, OrExprType}:                        _GotoState92Action,
	{_State298, BinaryOrExprType}:                  _GotoState69Action,
	{_State298, InitializableTypeExprType}:         _GotoState83Action,
	{_State298, SliceTypeExprType}:                 _GotoState101Action,
	{_State298, ArrayTypeExprType}:                 _GotoState60Action,
	{_State298, MapTypeExprType}:                   _GotoState88Action,
	{_State298, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State298, AnonymousFuncExprType}:             _GotoState59Action,
	{_State299, RparenToken}:                       _GotoState386Action,
	{_State301, NewlinesToken}:                     _GotoState388Action,
	{_State301, CommaToken}:                        _GotoState387Action,
	{_State302, GreaterToken}:                      _GotoState389Action,
	{_State303, CommaToken}:                        _GotoState390Action,
	{_State304, RbracketToken}:                     _GotoState391Action,
	{_State305, AddToken}:                          _GotoState249Action,
	{_State305, SubToken}:                          _GotoState251Action,
	{_State305, MulToken}:                          _GotoState250Action,
	{_State305, BinaryTypeOpType}:                  _GotoState252Action,
	{_State307, RbracketToken}:                     _GotoState392Action,
	{_State309, IntegerLiteralToken}:               _GotoState40Action,
	{_State309, FloatLiteralToken}:                 _GotoState35Action,
	{_State309, RuneLiteralToken}:                  _GotoState48Action,
	{_State309, StringLiteralToken}:                _GotoState49Action,
	{_State309, IdentifierToken}:                   _GotoState167Action,
	{_State309, UnderscoreToken}:                   _GotoState53Action,
	{_State309, TrueToken}:                         _GotoState52Action,
	{_State309, FalseToken}:                        _GotoState34Action,
	{_State309, StructToken}:                       _GotoState50Action,
	{_State309, FuncToken}:                         _GotoState36Action,
	{_State309, VarToken}:                          _GotoState15Action,
	{_State309, LetToken}:                          _GotoState12Action,
	{_State309, NotToken}:                          _GotoState45Action,
	{_State309, LabelDeclToken}:                    _GotoState41Action,
	{_State309, LparenToken}:                       _GotoState43Action,
	{_State309, LbracketToken}:                     _GotoState42Action,
	{_State309, ColonToken}:                        _GotoState165Action,
	{_State309, EllipsisToken}:                     _GotoState166Action,
	{_State309, SubToken}:                          _GotoState51Action,
	{_State309, MulToken}:                          _GotoState44Action,
	{_State309, BitNegToken}:                       _GotoState27Action,
	{_State309, BitAndToken}:                       _GotoState26Action,
	{_State309, GreaterToken}:                      _GotoState37Action,
	{_State309, ParseErrorToken}:                   _GotoState46Action,
	{_State309, VarDeclPatternType}:                _GotoState104Action,
	{_State309, VarOrLetType}:                      _GotoState24Action,
	{_State309, ExprType}:                          _GotoState171Action,
	{_State309, OptionalLabelDeclType}:             _GotoState91Action,
	{_State309, SequenceExprType}:                  _GotoState106Action,
	{_State309, CallExprType}:                      _GotoState71Action,
	{_State309, ProperArgumentsType}:               _GotoState172Action,
	{_State309, ArgumentsType}:                     _GotoState393Action,
	{_State309, ArgumentType}:                      _GotoState168Action,
	{_State309, ColonExprType}:                     _GotoState170Action,
	{_State309, AtomExprType}:                      _GotoState63Action,
	{_State309, ParseErrorExprType}:                _GotoState93Action,
	{_State309, LiteralExprType}:                   _GotoState87Action,
	{_State309, NamedExprType}:                     _GotoState90Action,
	{_State309, BlockExprType}:                     _GotoState70Action,
	{_State309, InitializeExprType}:                _GotoState84Action,
	{_State309, ImplicitStructExprType}:            _GotoState80Action,
	{_State309, AccessibleExprType}:                _GotoState105Action,
	{_State309, AccessExprType}:                    _GotoState55Action,
	{_State309, IndexExprType}:                     _GotoState82Action,
	{_State309, PostfixableExprType}:               _GotoState95Action,
	{_State309, PostfixUnaryExprType}:              _GotoState94Action,
	{_State309, PrefixableExprType}:                _GotoState98Action,
	{_State309, PrefixUnaryExprType}:               _GotoState96Action,
	{_State309, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State309, MulExprType}:                       _GotoState89Action,
	{_State309, BinaryMulExprType}:                 _GotoState67Action,
	{_State309, AddExprType}:                       _GotoState57Action,
	{_State309, BinaryAddExprType}:                 _GotoState64Action,
	{_State309, CmpExprType}:                       _GotoState74Action,
	{_State309, BinaryCmpExprType}:                 _GotoState66Action,
	{_State309, AndExprType}:                       _GotoState58Action,
	{_State309, BinaryAndExprType}:                 _GotoState65Action,
	{_State309, OrExprType}:                        _GotoState92Action,
	{_State309, BinaryOrExprType}:                  _GotoState69Action,
	{_State309, InitializableTypeExprType}:         _GotoState83Action,
	{_State309, SliceTypeExprType}:                 _GotoState101Action,
	{_State309, ArrayTypeExprType}:                 _GotoState60Action,
	{_State309, MapTypeExprType}:                   _GotoState88Action,
	{_State309, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State309, AnonymousFuncExprType}:             _GotoState59Action,
	{_State310, MulToken}:                          _GotoState220Action,
	{_State310, DivToken}:                          _GotoState218Action,
	{_State310, ModToken}:                          _GotoState219Action,
	{_State310, BitAndToken}:                       _GotoState215Action,
	{_State310, BitLshiftToken}:                    _GotoState216Action,
	{_State310, BitRshiftToken}:                    _GotoState217Action,
	{_State310, MulOpType}:                         _GotoState221Action,
	{_State311, EqualToken}:                        _GotoState204Action,
	{_State311, NotEqualToken}:                     _GotoState209Action,
	{_State311, LessToken}:                         _GotoState207Action,
	{_State311, LessOrEqualToken}:                  _GotoState208Action,
	{_State311, GreaterToken}:                      _GotoState205Action,
	{_State311, GreaterOrEqualToken}:               _GotoState206Action,
	{_State311, CmpOpType}:                         _GotoState210Action,
	{_State313, AddToken}:                          _GotoState195Action,
	{_State313, SubToken}:                          _GotoState198Action,
	{_State313, BitXorToken}:                       _GotoState197Action,
	{_State313, BitOrToken}:                        _GotoState196Action,
	{_State313, AddOpType}:                         _GotoState199Action,
	{_State315, RparenToken}:                       _GotoState394Action,
	{_State316, CommaToken}:                        _GotoState211Action,
	{_State318, ForToken}:                          _GotoState395Action,
	{_State319, InToken}:                           _GotoState397Action,
	{_State319, AssignToken}:                       _GotoState396Action,
	{_State320, SemicolonToken}:                    _GotoState398Action,
	{_State321, DoToken}:                           _GotoState399Action,
	{_State322, IntegerLiteralToken}:               _GotoState40Action,
	{_State322, FloatLiteralToken}:                 _GotoState35Action,
	{_State322, RuneLiteralToken}:                  _GotoState48Action,
	{_State322, StringLiteralToken}:                _GotoState49Action,
	{_State322, IdentifierToken}:                   _GotoState38Action,
	{_State322, UnderscoreToken}:                   _GotoState53Action,
	{_State322, TrueToken}:                         _GotoState52Action,
	{_State322, FalseToken}:                        _GotoState34Action,
	{_State322, StructToken}:                       _GotoState50Action,
	{_State322, FuncToken}:                         _GotoState36Action,
	{_State322, VarToken}:                          _GotoState149Action,
	{_State322, LetToken}:                          _GotoState12Action,
	{_State322, NotToken}:                          _GotoState45Action,
	{_State322, LabelDeclToken}:                    _GotoState41Action,
	{_State322, LparenToken}:                       _GotoState43Action,
	{_State322, LbracketToken}:                     _GotoState42Action,
	{_State322, DotToken}:                          _GotoState148Action,
	{_State322, SubToken}:                          _GotoState51Action,
	{_State322, MulToken}:                          _GotoState44Action,
	{_State322, BitNegToken}:                       _GotoState27Action,
	{_State322, BitAndToken}:                       _GotoState26Action,
	{_State322, GreaterToken}:                      _GotoState37Action,
	{_State322, ParseErrorToken}:                   _GotoState46Action,
	{_State322, CasePatternsType}:                  _GotoState400Action,
	{_State322, VarDeclPatternType}:                _GotoState104Action,
	{_State322, VarOrLetType}:                      _GotoState24Action,
	{_State322, AssignPatternType}:                 _GotoState150Action,
	{_State322, CasePatternType}:                   _GotoState151Action,
	{_State322, OptionalLabelDeclType}:             _GotoState153Action,
	{_State322, SequenceExprType}:                  _GotoState154Action,
	{_State322, CallExprType}:                      _GotoState71Action,
	{_State322, AtomExprType}:                      _GotoState63Action,
	{_State322, ParseErrorExprType}:                _GotoState93Action,
	{_State322, LiteralExprType}:                   _GotoState87Action,
	{_State322, NamedExprType}:                     _GotoState90Action,
	{_State322, BlockExprType}:                     _GotoState70Action,
	{_State322, InitializeExprType}:                _GotoState84Action,
	{_State322, ImplicitStructExprType}:            _GotoState80Action,
	{_State322, AccessibleExprType}:                _GotoState105Action,
	{_State322, AccessExprType}:                    _GotoState55Action,
	{_State322, IndexExprType}:                     _GotoState82Action,
	{_State322, PostfixableExprType}:               _GotoState95Action,
	{_State322, PostfixUnaryExprType}:              _GotoState94Action,
	{_State322, PrefixableExprType}:                _GotoState98Action,
	{_State322, PrefixUnaryExprType}:               _GotoState96Action,
	{_State322, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State322, MulExprType}:                       _GotoState89Action,
	{_State322, BinaryMulExprType}:                 _GotoState67Action,
	{_State322, AddExprType}:                       _GotoState57Action,
	{_State322, BinaryAddExprType}:                 _GotoState64Action,
	{_State322, CmpExprType}:                       _GotoState74Action,
	{_State322, BinaryCmpExprType}:                 _GotoState66Action,
	{_State322, AndExprType}:                       _GotoState58Action,
	{_State322, BinaryAndExprType}:                 _GotoState65Action,
	{_State322, OrExprType}:                        _GotoState92Action,
	{_State322, BinaryOrExprType}:                  _GotoState69Action,
	{_State322, InitializableTypeExprType}:         _GotoState83Action,
	{_State322, SliceTypeExprType}:                 _GotoState101Action,
	{_State322, ArrayTypeExprType}:                 _GotoState60Action,
	{_State322, MapTypeExprType}:                   _GotoState88Action,
	{_State322, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State322, AnonymousFuncExprType}:             _GotoState59Action,
	{_State323, LbraceToken}:                       _GotoState11Action,
	{_State323, StatementBlockType}:                _GotoState401Action,
	{_State325, LbraceToken}:                       _GotoState11Action,
	{_State325, StatementBlockType}:                _GotoState402Action,
	{_State326, AndToken}:                          _GotoState200Action,
	{_State327, RparenToken}:                       _GotoState403Action,
	{_State328, NewlinesToken}:                     _GotoState404Action,
	{_State328, OrToken}:                           _GotoState405Action,
	{_State329, NewlinesToken}:                     _GotoState406Action,
	{_State329, OrToken}:                           _GotoState407Action,
	{_State330, IdentifierToken}:                   _GotoState113Action,
	{_State330, UnderscoreToken}:                   _GotoState119Action,
	{_State330, StructToken}:                       _GotoState50Action,
	{_State330, EnumToken}:                         _GotoState110Action,
	{_State330, TraitToken}:                        _GotoState118Action,
	{_State330, FuncToken}:                         _GotoState112Action,
	{_State330, LparenToken}:                       _GotoState114Action,
	{_State330, LbracketToken}:                     _GotoState42Action,
	{_State330, DotToken}:                          _GotoState109Action,
	{_State330, QuestionToken}:                     _GotoState116Action,
	{_State330, ExclaimToken}:                      _GotoState111Action,
	{_State330, TildeTildeToken}:                   _GotoState117Action,
	{_State330, BitNegToken}:                       _GotoState108Action,
	{_State330, BitAndToken}:                       _GotoState107Action,
	{_State330, ParseErrorToken}:                   _GotoState115Action,
	{_State330, InitializableTypeExprType}:         _GotoState127Action,
	{_State330, SliceTypeExprType}:                 _GotoState101Action,
	{_State330, ArrayTypeExprType}:                 _GotoState60Action,
	{_State330, MapTypeExprType}:                   _GotoState88Action,
	{_State330, AtomTypeExprType}:                  _GotoState120Action,
	{_State330, NamedTypeExprType}:                 _GotoState128Action,
	{_State330, InferredTypeExprType}:              _GotoState126Action,
	{_State330, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State330, ReturnableTypeExprType}:            _GotoState132Action,
	{_State330, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State330, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State330, TypeExprType}:                      _GotoState408Action,
	{_State330, BinaryTypeExprType}:                _GotoState121Action,
	{_State330, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State330, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State330, TraitTypeExprType}:                 _GotoState133Action,
	{_State330, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State330, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State330, FuncTypeExprType}:                  _GotoState123Action,
	{_State331, IdentifierToken}:                   _GotoState113Action,
	{_State331, UnderscoreToken}:                   _GotoState119Action,
	{_State331, StructToken}:                       _GotoState50Action,
	{_State331, EnumToken}:                         _GotoState110Action,
	{_State331, TraitToken}:                        _GotoState118Action,
	{_State331, FuncToken}:                         _GotoState112Action,
	{_State331, LparenToken}:                       _GotoState114Action,
	{_State331, LbracketToken}:                     _GotoState42Action,
	{_State331, DotToken}:                          _GotoState340Action,
	{_State331, QuestionToken}:                     _GotoState116Action,
	{_State331, ExclaimToken}:                      _GotoState111Action,
	{_State331, DollarLbracketToken}:               _GotoState184Action,
	{_State331, EllipsisToken}:                     _GotoState359Action,
	{_State331, TildeTildeToken}:                   _GotoState117Action,
	{_State331, BitNegToken}:                       _GotoState108Action,
	{_State331, BitAndToken}:                       _GotoState107Action,
	{_State331, ParseErrorToken}:                   _GotoState115Action,
	{_State331, InitializableTypeExprType}:         _GotoState127Action,
	{_State331, SliceTypeExprType}:                 _GotoState101Action,
	{_State331, ArrayTypeExprType}:                 _GotoState60Action,
	{_State331, MapTypeExprType}:                   _GotoState88Action,
	{_State331, AtomTypeExprType}:                  _GotoState120Action,
	{_State331, NamedTypeExprType}:                 _GotoState128Action,
	{_State331, InferredTypeExprType}:              _GotoState126Action,
	{_State331, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State331, ReturnableTypeExprType}:            _GotoState132Action,
	{_State331, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State331, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State331, TypeExprType}:                      _GotoState360Action,
	{_State331, BinaryTypeExprType}:                _GotoState121Action,
	{_State331, GenericTypeArgumentsType}:          _GotoState235Action,
	{_State331, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State331, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State331, TraitTypeExprType}:                 _GotoState133Action,
	{_State331, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State331, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State331, FuncTypeExprType}:                  _GotoState123Action,
	{_State332, IdentifierToken}:                   _GotoState113Action,
	{_State332, UnderscoreToken}:                   _GotoState119Action,
	{_State332, StructToken}:                       _GotoState50Action,
	{_State332, EnumToken}:                         _GotoState110Action,
	{_State332, TraitToken}:                        _GotoState118Action,
	{_State332, FuncToken}:                         _GotoState112Action,
	{_State332, LparenToken}:                       _GotoState114Action,
	{_State332, LbracketToken}:                     _GotoState42Action,
	{_State332, DotToken}:                          _GotoState109Action,
	{_State332, QuestionToken}:                     _GotoState116Action,
	{_State332, ExclaimToken}:                      _GotoState111Action,
	{_State332, EllipsisToken}:                     _GotoState361Action,
	{_State332, TildeTildeToken}:                   _GotoState117Action,
	{_State332, BitNegToken}:                       _GotoState108Action,
	{_State332, BitAndToken}:                       _GotoState107Action,
	{_State332, ParseErrorToken}:                   _GotoState115Action,
	{_State332, InitializableTypeExprType}:         _GotoState127Action,
	{_State332, SliceTypeExprType}:                 _GotoState101Action,
	{_State332, ArrayTypeExprType}:                 _GotoState60Action,
	{_State332, MapTypeExprType}:                   _GotoState88Action,
	{_State332, AtomTypeExprType}:                  _GotoState120Action,
	{_State332, NamedTypeExprType}:                 _GotoState128Action,
	{_State332, InferredTypeExprType}:              _GotoState126Action,
	{_State332, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State332, ReturnableTypeExprType}:            _GotoState132Action,
	{_State332, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State332, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State332, TypeExprType}:                      _GotoState362Action,
	{_State332, BinaryTypeExprType}:                _GotoState121Action,
	{_State332, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State332, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State332, TraitTypeExprType}:                 _GotoState133Action,
	{_State332, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State332, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State332, FuncTypeExprType}:                  _GotoState123Action,
	{_State334, RparenToken}:                       _GotoState409Action,
	{_State335, CommaToken}:                        _GotoState410Action,
	{_State337, AddToken}:                          _GotoState249Action,
	{_State337, SubToken}:                          _GotoState251Action,
	{_State337, MulToken}:                          _GotoState250Action,
	{_State337, BinaryTypeOpType}:                  _GotoState252Action,
	{_State338, DollarLbracketToken}:               _GotoState184Action,
	{_State338, GenericTypeArgumentsType}:          _GotoState411Action,
	{_State339, LparenToken}:                       _GotoState412Action,
	{_State340, IdentifierToken}:                   _GotoState338Action,
	{_State341, AssignToken}:                       _GotoState349Action,
	{_State341, AddToken}:                          _GotoState249Action,
	{_State341, SubToken}:                          _GotoState251Action,
	{_State341, MulToken}:                          _GotoState250Action,
	{_State341, BinaryTypeOpType}:                  _GotoState252Action,
	{_State341, OptionalDefaultType}:               _GotoState413Action,
	{_State342, AddToken}:                          _GotoState249Action,
	{_State342, SubToken}:                          _GotoState251Action,
	{_State342, MulToken}:                          _GotoState250Action,
	{_State342, BinaryTypeOpType}:                  _GotoState252Action,
	{_State343, IdentifierToken}:                   _GotoState237Action,
	{_State343, UnderscoreToken}:                   _GotoState238Action,
	{_State343, UnsafeToken}:                       _GotoState54Action,
	{_State343, StructToken}:                       _GotoState50Action,
	{_State343, EnumToken}:                         _GotoState110Action,
	{_State343, TraitToken}:                        _GotoState118Action,
	{_State343, FuncToken}:                         _GotoState236Action,
	{_State343, LparenToken}:                       _GotoState114Action,
	{_State343, LbracketToken}:                     _GotoState42Action,
	{_State343, DotToken}:                          _GotoState109Action,
	{_State343, QuestionToken}:                     _GotoState116Action,
	{_State343, ExclaimToken}:                      _GotoState111Action,
	{_State343, TildeTildeToken}:                   _GotoState117Action,
	{_State343, BitNegToken}:                       _GotoState108Action,
	{_State343, BitAndToken}:                       _GotoState107Action,
	{_State343, ParseErrorToken}:                   _GotoState115Action,
	{_State343, UnsafeStatementType}:               _GotoState246Action,
	{_State343, InitializableTypeExprType}:         _GotoState127Action,
	{_State343, SliceTypeExprType}:                 _GotoState101Action,
	{_State343, ArrayTypeExprType}:                 _GotoState60Action,
	{_State343, MapTypeExprType}:                   _GotoState88Action,
	{_State343, AtomTypeExprType}:                  _GotoState120Action,
	{_State343, NamedTypeExprType}:                 _GotoState128Action,
	{_State343, InferredTypeExprType}:              _GotoState126Action,
	{_State343, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State343, ReturnableTypeExprType}:            _GotoState132Action,
	{_State343, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State343, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State343, TypeExprType}:                      _GotoState245Action,
	{_State343, BinaryTypeExprType}:                _GotoState121Action,
	{_State343, FieldDefType}:                      _GotoState414Action,
	{_State343, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State343, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State343, TraitTypeExprType}:                 _GotoState133Action,
	{_State343, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State343, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State343, FuncTypeExprType}:                  _GotoState123Action,
	{_State343, MethodSignatureType}:               _GotoState242Action,
	{_State347, IdentifierToken}:                   _GotoState237Action,
	{_State347, UnderscoreToken}:                   _GotoState238Action,
	{_State347, UnsafeToken}:                       _GotoState54Action,
	{_State347, StructToken}:                       _GotoState50Action,
	{_State347, EnumToken}:                         _GotoState110Action,
	{_State347, TraitToken}:                        _GotoState118Action,
	{_State347, FuncToken}:                         _GotoState236Action,
	{_State347, LparenToken}:                       _GotoState114Action,
	{_State347, LbracketToken}:                     _GotoState42Action,
	{_State347, DotToken}:                          _GotoState109Action,
	{_State347, QuestionToken}:                     _GotoState116Action,
	{_State347, ExclaimToken}:                      _GotoState111Action,
	{_State347, TildeTildeToken}:                   _GotoState117Action,
	{_State347, BitNegToken}:                       _GotoState108Action,
	{_State347, BitAndToken}:                       _GotoState107Action,
	{_State347, ParseErrorToken}:                   _GotoState115Action,
	{_State347, UnsafeStatementType}:               _GotoState246Action,
	{_State347, InitializableTypeExprType}:         _GotoState127Action,
	{_State347, SliceTypeExprType}:                 _GotoState101Action,
	{_State347, ArrayTypeExprType}:                 _GotoState60Action,
	{_State347, MapTypeExprType}:                   _GotoState88Action,
	{_State347, AtomTypeExprType}:                  _GotoState120Action,
	{_State347, NamedTypeExprType}:                 _GotoState128Action,
	{_State347, InferredTypeExprType}:              _GotoState126Action,
	{_State347, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State347, ReturnableTypeExprType}:            _GotoState132Action,
	{_State347, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State347, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State347, TypeExprType}:                      _GotoState245Action,
	{_State347, BinaryTypeExprType}:                _GotoState121Action,
	{_State347, FieldDefType}:                      _GotoState415Action,
	{_State347, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State347, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State347, TraitTypeExprType}:                 _GotoState133Action,
	{_State347, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State347, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State347, FuncTypeExprType}:                  _GotoState123Action,
	{_State347, MethodSignatureType}:               _GotoState242Action,
	{_State348, IdentifierToken}:                   _GotoState237Action,
	{_State348, UnderscoreToken}:                   _GotoState238Action,
	{_State348, UnsafeToken}:                       _GotoState54Action,
	{_State348, StructToken}:                       _GotoState50Action,
	{_State348, EnumToken}:                         _GotoState110Action,
	{_State348, TraitToken}:                        _GotoState118Action,
	{_State348, FuncToken}:                         _GotoState236Action,
	{_State348, LparenToken}:                       _GotoState114Action,
	{_State348, LbracketToken}:                     _GotoState42Action,
	{_State348, DotToken}:                          _GotoState109Action,
	{_State348, QuestionToken}:                     _GotoState116Action,
	{_State348, ExclaimToken}:                      _GotoState111Action,
	{_State348, TildeTildeToken}:                   _GotoState117Action,
	{_State348, BitNegToken}:                       _GotoState108Action,
	{_State348, BitAndToken}:                       _GotoState107Action,
	{_State348, ParseErrorToken}:                   _GotoState115Action,
	{_State348, UnsafeStatementType}:               _GotoState246Action,
	{_State348, InitializableTypeExprType}:         _GotoState127Action,
	{_State348, SliceTypeExprType}:                 _GotoState101Action,
	{_State348, ArrayTypeExprType}:                 _GotoState60Action,
	{_State348, MapTypeExprType}:                   _GotoState88Action,
	{_State348, AtomTypeExprType}:                  _GotoState120Action,
	{_State348, NamedTypeExprType}:                 _GotoState128Action,
	{_State348, InferredTypeExprType}:              _GotoState126Action,
	{_State348, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State348, ReturnableTypeExprType}:            _GotoState132Action,
	{_State348, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State348, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State348, TypeExprType}:                      _GotoState245Action,
	{_State348, BinaryTypeExprType}:                _GotoState121Action,
	{_State348, FieldDefType}:                      _GotoState416Action,
	{_State348, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State348, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State348, TraitTypeExprType}:                 _GotoState133Action,
	{_State348, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State348, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State348, FuncTypeExprType}:                  _GotoState123Action,
	{_State348, MethodSignatureType}:               _GotoState242Action,
	{_State349, DefaultToken}:                      _GotoState417Action,
	{_State351, RparenToken}:                       _GotoState418Action,
	{_State354, IdentifierToken}:                   _GotoState113Action,
	{_State354, UnderscoreToken}:                   _GotoState119Action,
	{_State354, StructToken}:                       _GotoState50Action,
	{_State354, EnumToken}:                         _GotoState110Action,
	{_State354, TraitToken}:                        _GotoState118Action,
	{_State354, FuncToken}:                         _GotoState112Action,
	{_State354, LparenToken}:                       _GotoState114Action,
	{_State354, LbracketToken}:                     _GotoState42Action,
	{_State354, DotToken}:                          _GotoState109Action,
	{_State354, QuestionToken}:                     _GotoState116Action,
	{_State354, ExclaimToken}:                      _GotoState111Action,
	{_State354, TildeTildeToken}:                   _GotoState117Action,
	{_State354, BitNegToken}:                       _GotoState108Action,
	{_State354, BitAndToken}:                       _GotoState107Action,
	{_State354, ParseErrorToken}:                   _GotoState115Action,
	{_State354, InitializableTypeExprType}:         _GotoState127Action,
	{_State354, SliceTypeExprType}:                 _GotoState101Action,
	{_State354, ArrayTypeExprType}:                 _GotoState60Action,
	{_State354, MapTypeExprType}:                   _GotoState88Action,
	{_State354, AtomTypeExprType}:                  _GotoState120Action,
	{_State354, NamedTypeExprType}:                 _GotoState128Action,
	{_State354, InferredTypeExprType}:              _GotoState126Action,
	{_State354, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State354, ReturnableTypeExprType}:            _GotoState132Action,
	{_State354, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State354, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State354, TypeExprType}:                      _GotoState419Action,
	{_State354, BinaryTypeExprType}:                _GotoState121Action,
	{_State354, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State354, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State354, TraitTypeExprType}:                 _GotoState133Action,
	{_State354, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State354, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State354, FuncTypeExprType}:                  _GotoState123Action,
	{_State356, RbracketToken}:                     _GotoState420Action,
	{_State357, CommaToken}:                        _GotoState421Action,
	{_State358, IdentifierToken}:                   _GotoState256Action,
	{_State358, UnderscoreToken}:                   _GotoState257Action,
	{_State358, ProperParameterDefType}:            _GotoState259Action,
	{_State358, ParameterDefType}:                  _GotoState280Action,
	{_State358, ProperParameterDefsType}:           _GotoState282Action,
	{_State358, ParameterDefsType}:                 _GotoState422Action,
	{_State359, IdentifierToken}:                   _GotoState113Action,
	{_State359, UnderscoreToken}:                   _GotoState119Action,
	{_State359, StructToken}:                       _GotoState50Action,
	{_State359, EnumToken}:                         _GotoState110Action,
	{_State359, TraitToken}:                        _GotoState118Action,
	{_State359, FuncToken}:                         _GotoState112Action,
	{_State359, LparenToken}:                       _GotoState114Action,
	{_State359, LbracketToken}:                     _GotoState42Action,
	{_State359, DotToken}:                          _GotoState109Action,
	{_State359, QuestionToken}:                     _GotoState116Action,
	{_State359, ExclaimToken}:                      _GotoState111Action,
	{_State359, TildeTildeToken}:                   _GotoState117Action,
	{_State359, BitNegToken}:                       _GotoState108Action,
	{_State359, BitAndToken}:                       _GotoState107Action,
	{_State359, ParseErrorToken}:                   _GotoState115Action,
	{_State359, InitializableTypeExprType}:         _GotoState127Action,
	{_State359, SliceTypeExprType}:                 _GotoState101Action,
	{_State359, ArrayTypeExprType}:                 _GotoState60Action,
	{_State359, MapTypeExprType}:                   _GotoState88Action,
	{_State359, AtomTypeExprType}:                  _GotoState120Action,
	{_State359, NamedTypeExprType}:                 _GotoState128Action,
	{_State359, InferredTypeExprType}:              _GotoState126Action,
	{_State359, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State359, ReturnableTypeExprType}:            _GotoState132Action,
	{_State359, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State359, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State359, TypeExprType}:                      _GotoState423Action,
	{_State359, BinaryTypeExprType}:                _GotoState121Action,
	{_State359, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State359, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State359, TraitTypeExprType}:                 _GotoState133Action,
	{_State359, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State359, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State359, FuncTypeExprType}:                  _GotoState123Action,
	{_State360, AddToken}:                          _GotoState249Action,
	{_State360, SubToken}:                          _GotoState251Action,
	{_State360, MulToken}:                          _GotoState250Action,
	{_State360, BinaryTypeOpType}:                  _GotoState252Action,
	{_State361, IdentifierToken}:                   _GotoState113Action,
	{_State361, UnderscoreToken}:                   _GotoState119Action,
	{_State361, StructToken}:                       _GotoState50Action,
	{_State361, EnumToken}:                         _GotoState110Action,
	{_State361, TraitToken}:                        _GotoState118Action,
	{_State361, FuncToken}:                         _GotoState112Action,
	{_State361, LparenToken}:                       _GotoState114Action,
	{_State361, LbracketToken}:                     _GotoState42Action,
	{_State361, DotToken}:                          _GotoState109Action,
	{_State361, QuestionToken}:                     _GotoState116Action,
	{_State361, ExclaimToken}:                      _GotoState111Action,
	{_State361, TildeTildeToken}:                   _GotoState117Action,
	{_State361, BitNegToken}:                       _GotoState108Action,
	{_State361, BitAndToken}:                       _GotoState107Action,
	{_State361, ParseErrorToken}:                   _GotoState115Action,
	{_State361, InitializableTypeExprType}:         _GotoState127Action,
	{_State361, SliceTypeExprType}:                 _GotoState101Action,
	{_State361, ArrayTypeExprType}:                 _GotoState60Action,
	{_State361, MapTypeExprType}:                   _GotoState88Action,
	{_State361, AtomTypeExprType}:                  _GotoState120Action,
	{_State361, NamedTypeExprType}:                 _GotoState128Action,
	{_State361, InferredTypeExprType}:              _GotoState126Action,
	{_State361, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State361, ReturnableTypeExprType}:            _GotoState132Action,
	{_State361, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State361, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State361, TypeExprType}:                      _GotoState424Action,
	{_State361, BinaryTypeExprType}:                _GotoState121Action,
	{_State361, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State361, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State361, TraitTypeExprType}:                 _GotoState133Action,
	{_State361, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State361, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State361, FuncTypeExprType}:                  _GotoState123Action,
	{_State362, AddToken}:                          _GotoState249Action,
	{_State362, SubToken}:                          _GotoState251Action,
	{_State362, MulToken}:                          _GotoState250Action,
	{_State362, BinaryTypeOpType}:                  _GotoState252Action,
	{_State363, IdentifierToken}:                   _GotoState425Action,
	{_State366, AddToken}:                          _GotoState249Action,
	{_State366, SubToken}:                          _GotoState251Action,
	{_State366, MulToken}:                          _GotoState250Action,
	{_State366, BinaryTypeOpType}:                  _GotoState252Action,
	{_State367, ImplementsToken}:                   _GotoState426Action,
	{_State367, AddToken}:                          _GotoState249Action,
	{_State367, SubToken}:                          _GotoState251Action,
	{_State367, MulToken}:                          _GotoState250Action,
	{_State367, BinaryTypeOpType}:                  _GotoState252Action,
	{_State368, IdentifierToken}:                   _GotoState143Action,
	{_State368, UnderscoreToken}:                   _GotoState145Action,
	{_State368, LparenToken}:                       _GotoState144Action,
	{_State368, VarPatternType}:                    _GotoState427Action,
	{_State368, TuplePatternType}:                  _GotoState146Action,
	{_State369, IdentifierToken}:                   _GotoState268Action,
	{_State369, UnderscoreToken}:                   _GotoState145Action,
	{_State369, LparenToken}:                       _GotoState144Action,
	{_State369, EllipsisToken}:                     _GotoState267Action,
	{_State369, VarPatternType}:                    _GotoState271Action,
	{_State369, TuplePatternType}:                  _GotoState146Action,
	{_State369, FieldVarPatternType}:               _GotoState428Action,
	{_State372, LparenToken}:                       _GotoState144Action,
	{_State372, TuplePatternType}:                  _GotoState429Action,
	{_State375, IdentifierToken}:                   _GotoState113Action,
	{_State375, UnderscoreToken}:                   _GotoState119Action,
	{_State375, StructToken}:                       _GotoState50Action,
	{_State375, EnumToken}:                         _GotoState110Action,
	{_State375, TraitToken}:                        _GotoState118Action,
	{_State375, FuncToken}:                         _GotoState112Action,
	{_State375, LparenToken}:                       _GotoState114Action,
	{_State375, LbracketToken}:                     _GotoState42Action,
	{_State375, DotToken}:                          _GotoState109Action,
	{_State375, QuestionToken}:                     _GotoState116Action,
	{_State375, ExclaimToken}:                      _GotoState111Action,
	{_State375, TildeTildeToken}:                   _GotoState117Action,
	{_State375, BitNegToken}:                       _GotoState108Action,
	{_State375, BitAndToken}:                       _GotoState107Action,
	{_State375, ParseErrorToken}:                   _GotoState115Action,
	{_State375, InitializableTypeExprType}:         _GotoState127Action,
	{_State375, SliceTypeExprType}:                 _GotoState101Action,
	{_State375, ArrayTypeExprType}:                 _GotoState60Action,
	{_State375, MapTypeExprType}:                   _GotoState88Action,
	{_State375, AtomTypeExprType}:                  _GotoState120Action,
	{_State375, NamedTypeExprType}:                 _GotoState128Action,
	{_State375, InferredTypeExprType}:              _GotoState126Action,
	{_State375, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State375, ReturnableTypeExprType}:            _GotoState431Action,
	{_State375, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State375, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State375, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State375, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State375, TraitTypeExprType}:                 _GotoState133Action,
	{_State375, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State375, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State375, ReturnTypeType}:                    _GotoState430Action,
	{_State375, FuncTypeExprType}:                  _GotoState123Action,
	{_State376, IdentifierToken}:                   _GotoState256Action,
	{_State376, UnderscoreToken}:                   _GotoState257Action,
	{_State376, ProperParameterDefType}:            _GotoState259Action,
	{_State376, ParameterDefType}:                  _GotoState432Action,
	{_State378, StringLiteralToken}:                _GotoState161Action,
	{_State378, IdentifierToken}:                   _GotoState159Action,
	{_State378, UnderscoreToken}:                   _GotoState162Action,
	{_State378, DotToken}:                          _GotoState158Action,
	{_State378, ImportClauseType}:                  _GotoState433Action,
	{_State379, StringLiteralToken}:                _GotoState161Action,
	{_State379, IdentifierToken}:                   _GotoState159Action,
	{_State379, UnderscoreToken}:                   _GotoState162Action,
	{_State379, DotToken}:                          _GotoState158Action,
	{_State379, ImportClauseType}:                  _GotoState434Action,
	{_State380, RbracketToken}:                     _GotoState435Action,
	{_State380, AddToken}:                          _GotoState249Action,
	{_State380, SubToken}:                          _GotoState251Action,
	{_State380, MulToken}:                          _GotoState250Action,
	{_State380, BinaryTypeOpType}:                  _GotoState252Action,
	{_State381, RbracketToken}:                     _GotoState436Action,
	{_State387, IdentifierToken}:                   _GotoState237Action,
	{_State387, UnderscoreToken}:                   _GotoState238Action,
	{_State387, UnsafeToken}:                       _GotoState54Action,
	{_State387, StructToken}:                       _GotoState50Action,
	{_State387, EnumToken}:                         _GotoState110Action,
	{_State387, TraitToken}:                        _GotoState118Action,
	{_State387, FuncToken}:                         _GotoState236Action,
	{_State387, LparenToken}:                       _GotoState114Action,
	{_State387, LbracketToken}:                     _GotoState42Action,
	{_State387, DotToken}:                          _GotoState109Action,
	{_State387, QuestionToken}:                     _GotoState116Action,
	{_State387, ExclaimToken}:                      _GotoState111Action,
	{_State387, TildeTildeToken}:                   _GotoState117Action,
	{_State387, BitNegToken}:                       _GotoState108Action,
	{_State387, BitAndToken}:                       _GotoState107Action,
	{_State387, ParseErrorToken}:                   _GotoState115Action,
	{_State387, UnsafeStatementType}:               _GotoState246Action,
	{_State387, InitializableTypeExprType}:         _GotoState127Action,
	{_State387, SliceTypeExprType}:                 _GotoState101Action,
	{_State387, ArrayTypeExprType}:                 _GotoState60Action,
	{_State387, MapTypeExprType}:                   _GotoState88Action,
	{_State387, AtomTypeExprType}:                  _GotoState120Action,
	{_State387, NamedTypeExprType}:                 _GotoState128Action,
	{_State387, InferredTypeExprType}:              _GotoState126Action,
	{_State387, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State387, ReturnableTypeExprType}:            _GotoState132Action,
	{_State387, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State387, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State387, TypeExprType}:                      _GotoState245Action,
	{_State387, BinaryTypeExprType}:                _GotoState121Action,
	{_State387, FieldDefType}:                      _GotoState437Action,
	{_State387, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State387, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State387, TraitTypeExprType}:                 _GotoState133Action,
	{_State387, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State387, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State387, FuncTypeExprType}:                  _GotoState123Action,
	{_State387, MethodSignatureType}:               _GotoState242Action,
	{_State388, IdentifierToken}:                   _GotoState237Action,
	{_State388, UnderscoreToken}:                   _GotoState238Action,
	{_State388, UnsafeToken}:                       _GotoState54Action,
	{_State388, StructToken}:                       _GotoState50Action,
	{_State388, EnumToken}:                         _GotoState110Action,
	{_State388, TraitToken}:                        _GotoState118Action,
	{_State388, FuncToken}:                         _GotoState236Action,
	{_State388, LparenToken}:                       _GotoState114Action,
	{_State388, LbracketToken}:                     _GotoState42Action,
	{_State388, DotToken}:                          _GotoState109Action,
	{_State388, QuestionToken}:                     _GotoState116Action,
	{_State388, ExclaimToken}:                      _GotoState111Action,
	{_State388, TildeTildeToken}:                   _GotoState117Action,
	{_State388, BitNegToken}:                       _GotoState108Action,
	{_State388, BitAndToken}:                       _GotoState107Action,
	{_State388, ParseErrorToken}:                   _GotoState115Action,
	{_State388, UnsafeStatementType}:               _GotoState246Action,
	{_State388, InitializableTypeExprType}:         _GotoState127Action,
	{_State388, SliceTypeExprType}:                 _GotoState101Action,
	{_State388, ArrayTypeExprType}:                 _GotoState60Action,
	{_State388, MapTypeExprType}:                   _GotoState88Action,
	{_State388, AtomTypeExprType}:                  _GotoState120Action,
	{_State388, NamedTypeExprType}:                 _GotoState128Action,
	{_State388, InferredTypeExprType}:              _GotoState126Action,
	{_State388, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State388, ReturnableTypeExprType}:            _GotoState132Action,
	{_State388, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State388, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State388, TypeExprType}:                      _GotoState245Action,
	{_State388, BinaryTypeExprType}:                _GotoState121Action,
	{_State388, FieldDefType}:                      _GotoState438Action,
	{_State388, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State388, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State388, TraitTypeExprType}:                 _GotoState133Action,
	{_State388, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State388, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State388, FuncTypeExprType}:                  _GotoState123Action,
	{_State388, MethodSignatureType}:               _GotoState242Action,
	{_State389, StringLiteralToken}:                _GotoState439Action,
	{_State390, IdentifierToken}:                   _GotoState113Action,
	{_State390, UnderscoreToken}:                   _GotoState119Action,
	{_State390, StructToken}:                       _GotoState50Action,
	{_State390, EnumToken}:                         _GotoState110Action,
	{_State390, TraitToken}:                        _GotoState118Action,
	{_State390, FuncToken}:                         _GotoState112Action,
	{_State390, LparenToken}:                       _GotoState114Action,
	{_State390, LbracketToken}:                     _GotoState42Action,
	{_State390, DotToken}:                          _GotoState109Action,
	{_State390, QuestionToken}:                     _GotoState116Action,
	{_State390, ExclaimToken}:                      _GotoState111Action,
	{_State390, TildeTildeToken}:                   _GotoState117Action,
	{_State390, BitNegToken}:                       _GotoState108Action,
	{_State390, BitAndToken}:                       _GotoState107Action,
	{_State390, ParseErrorToken}:                   _GotoState115Action,
	{_State390, InitializableTypeExprType}:         _GotoState127Action,
	{_State390, SliceTypeExprType}:                 _GotoState101Action,
	{_State390, ArrayTypeExprType}:                 _GotoState60Action,
	{_State390, MapTypeExprType}:                   _GotoState88Action,
	{_State390, AtomTypeExprType}:                  _GotoState120Action,
	{_State390, NamedTypeExprType}:                 _GotoState128Action,
	{_State390, InferredTypeExprType}:              _GotoState126Action,
	{_State390, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State390, ReturnableTypeExprType}:            _GotoState132Action,
	{_State390, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State390, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State390, TypeExprType}:                      _GotoState440Action,
	{_State390, BinaryTypeExprType}:                _GotoState121Action,
	{_State390, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State390, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State390, TraitTypeExprType}:                 _GotoState133Action,
	{_State390, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State390, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State390, FuncTypeExprType}:                  _GotoState123Action,
	{_State393, RparenToken}:                       _GotoState441Action,
	{_State395, IntegerLiteralToken}:               _GotoState40Action,
	{_State395, FloatLiteralToken}:                 _GotoState35Action,
	{_State395, RuneLiteralToken}:                  _GotoState48Action,
	{_State395, StringLiteralToken}:                _GotoState49Action,
	{_State395, IdentifierToken}:                   _GotoState38Action,
	{_State395, UnderscoreToken}:                   _GotoState53Action,
	{_State395, TrueToken}:                         _GotoState52Action,
	{_State395, FalseToken}:                        _GotoState34Action,
	{_State395, StructToken}:                       _GotoState50Action,
	{_State395, FuncToken}:                         _GotoState36Action,
	{_State395, VarToken}:                          _GotoState15Action,
	{_State395, LetToken}:                          _GotoState12Action,
	{_State395, NotToken}:                          _GotoState45Action,
	{_State395, LabelDeclToken}:                    _GotoState41Action,
	{_State395, LparenToken}:                       _GotoState43Action,
	{_State395, LbracketToken}:                     _GotoState42Action,
	{_State395, SubToken}:                          _GotoState51Action,
	{_State395, MulToken}:                          _GotoState44Action,
	{_State395, BitNegToken}:                       _GotoState27Action,
	{_State395, BitAndToken}:                       _GotoState26Action,
	{_State395, GreaterToken}:                      _GotoState37Action,
	{_State395, ParseErrorToken}:                   _GotoState46Action,
	{_State395, VarDeclPatternType}:                _GotoState104Action,
	{_State395, VarOrLetType}:                      _GotoState24Action,
	{_State395, OptionalLabelDeclType}:             _GotoState153Action,
	{_State395, SequenceExprType}:                  _GotoState442Action,
	{_State395, CallExprType}:                      _GotoState71Action,
	{_State395, AtomExprType}:                      _GotoState63Action,
	{_State395, ParseErrorExprType}:                _GotoState93Action,
	{_State395, LiteralExprType}:                   _GotoState87Action,
	{_State395, NamedExprType}:                     _GotoState90Action,
	{_State395, BlockExprType}:                     _GotoState70Action,
	{_State395, InitializeExprType}:                _GotoState84Action,
	{_State395, ImplicitStructExprType}:            _GotoState80Action,
	{_State395, AccessibleExprType}:                _GotoState105Action,
	{_State395, AccessExprType}:                    _GotoState55Action,
	{_State395, IndexExprType}:                     _GotoState82Action,
	{_State395, PostfixableExprType}:               _GotoState95Action,
	{_State395, PostfixUnaryExprType}:              _GotoState94Action,
	{_State395, PrefixableExprType}:                _GotoState98Action,
	{_State395, PrefixUnaryExprType}:               _GotoState96Action,
	{_State395, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State395, MulExprType}:                       _GotoState89Action,
	{_State395, BinaryMulExprType}:                 _GotoState67Action,
	{_State395, AddExprType}:                       _GotoState57Action,
	{_State395, BinaryAddExprType}:                 _GotoState64Action,
	{_State395, CmpExprType}:                       _GotoState74Action,
	{_State395, BinaryCmpExprType}:                 _GotoState66Action,
	{_State395, AndExprType}:                       _GotoState58Action,
	{_State395, BinaryAndExprType}:                 _GotoState65Action,
	{_State395, OrExprType}:                        _GotoState92Action,
	{_State395, BinaryOrExprType}:                  _GotoState69Action,
	{_State395, InitializableTypeExprType}:         _GotoState83Action,
	{_State395, SliceTypeExprType}:                 _GotoState101Action,
	{_State395, ArrayTypeExprType}:                 _GotoState60Action,
	{_State395, MapTypeExprType}:                   _GotoState88Action,
	{_State395, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State395, AnonymousFuncExprType}:             _GotoState59Action,
	{_State396, IntegerLiteralToken}:               _GotoState40Action,
	{_State396, FloatLiteralToken}:                 _GotoState35Action,
	{_State396, RuneLiteralToken}:                  _GotoState48Action,
	{_State396, StringLiteralToken}:                _GotoState49Action,
	{_State396, IdentifierToken}:                   _GotoState38Action,
	{_State396, UnderscoreToken}:                   _GotoState53Action,
	{_State396, TrueToken}:                         _GotoState52Action,
	{_State396, FalseToken}:                        _GotoState34Action,
	{_State396, StructToken}:                       _GotoState50Action,
	{_State396, FuncToken}:                         _GotoState36Action,
	{_State396, VarToken}:                          _GotoState15Action,
	{_State396, LetToken}:                          _GotoState12Action,
	{_State396, NotToken}:                          _GotoState45Action,
	{_State396, LabelDeclToken}:                    _GotoState41Action,
	{_State396, LparenToken}:                       _GotoState43Action,
	{_State396, LbracketToken}:                     _GotoState42Action,
	{_State396, SubToken}:                          _GotoState51Action,
	{_State396, MulToken}:                          _GotoState44Action,
	{_State396, BitNegToken}:                       _GotoState27Action,
	{_State396, BitAndToken}:                       _GotoState26Action,
	{_State396, GreaterToken}:                      _GotoState37Action,
	{_State396, ParseErrorToken}:                   _GotoState46Action,
	{_State396, VarDeclPatternType}:                _GotoState104Action,
	{_State396, VarOrLetType}:                      _GotoState24Action,
	{_State396, OptionalLabelDeclType}:             _GotoState153Action,
	{_State396, SequenceExprType}:                  _GotoState443Action,
	{_State396, CallExprType}:                      _GotoState71Action,
	{_State396, AtomExprType}:                      _GotoState63Action,
	{_State396, ParseErrorExprType}:                _GotoState93Action,
	{_State396, LiteralExprType}:                   _GotoState87Action,
	{_State396, NamedExprType}:                     _GotoState90Action,
	{_State396, BlockExprType}:                     _GotoState70Action,
	{_State396, InitializeExprType}:                _GotoState84Action,
	{_State396, ImplicitStructExprType}:            _GotoState80Action,
	{_State396, AccessibleExprType}:                _GotoState105Action,
	{_State396, AccessExprType}:                    _GotoState55Action,
	{_State396, IndexExprType}:                     _GotoState82Action,
	{_State396, PostfixableExprType}:               _GotoState95Action,
	{_State396, PostfixUnaryExprType}:              _GotoState94Action,
	{_State396, PrefixableExprType}:                _GotoState98Action,
	{_State396, PrefixUnaryExprType}:               _GotoState96Action,
	{_State396, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State396, MulExprType}:                       _GotoState89Action,
	{_State396, BinaryMulExprType}:                 _GotoState67Action,
	{_State396, AddExprType}:                       _GotoState57Action,
	{_State396, BinaryAddExprType}:                 _GotoState64Action,
	{_State396, CmpExprType}:                       _GotoState74Action,
	{_State396, BinaryCmpExprType}:                 _GotoState66Action,
	{_State396, AndExprType}:                       _GotoState58Action,
	{_State396, BinaryAndExprType}:                 _GotoState65Action,
	{_State396, OrExprType}:                        _GotoState92Action,
	{_State396, BinaryOrExprType}:                  _GotoState69Action,
	{_State396, InitializableTypeExprType}:         _GotoState83Action,
	{_State396, SliceTypeExprType}:                 _GotoState101Action,
	{_State396, ArrayTypeExprType}:                 _GotoState60Action,
	{_State396, MapTypeExprType}:                   _GotoState88Action,
	{_State396, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State396, AnonymousFuncExprType}:             _GotoState59Action,
	{_State397, IntegerLiteralToken}:               _GotoState40Action,
	{_State397, FloatLiteralToken}:                 _GotoState35Action,
	{_State397, RuneLiteralToken}:                  _GotoState48Action,
	{_State397, StringLiteralToken}:                _GotoState49Action,
	{_State397, IdentifierToken}:                   _GotoState38Action,
	{_State397, UnderscoreToken}:                   _GotoState53Action,
	{_State397, TrueToken}:                         _GotoState52Action,
	{_State397, FalseToken}:                        _GotoState34Action,
	{_State397, StructToken}:                       _GotoState50Action,
	{_State397, FuncToken}:                         _GotoState36Action,
	{_State397, VarToken}:                          _GotoState15Action,
	{_State397, LetToken}:                          _GotoState12Action,
	{_State397, NotToken}:                          _GotoState45Action,
	{_State397, LabelDeclToken}:                    _GotoState41Action,
	{_State397, LparenToken}:                       _GotoState43Action,
	{_State397, LbracketToken}:                     _GotoState42Action,
	{_State397, SubToken}:                          _GotoState51Action,
	{_State397, MulToken}:                          _GotoState44Action,
	{_State397, BitNegToken}:                       _GotoState27Action,
	{_State397, BitAndToken}:                       _GotoState26Action,
	{_State397, GreaterToken}:                      _GotoState37Action,
	{_State397, ParseErrorToken}:                   _GotoState46Action,
	{_State397, VarDeclPatternType}:                _GotoState104Action,
	{_State397, VarOrLetType}:                      _GotoState24Action,
	{_State397, OptionalLabelDeclType}:             _GotoState153Action,
	{_State397, SequenceExprType}:                  _GotoState444Action,
	{_State397, CallExprType}:                      _GotoState71Action,
	{_State397, AtomExprType}:                      _GotoState63Action,
	{_State397, ParseErrorExprType}:                _GotoState93Action,
	{_State397, LiteralExprType}:                   _GotoState87Action,
	{_State397, NamedExprType}:                     _GotoState90Action,
	{_State397, BlockExprType}:                     _GotoState70Action,
	{_State397, InitializeExprType}:                _GotoState84Action,
	{_State397, ImplicitStructExprType}:            _GotoState80Action,
	{_State397, AccessibleExprType}:                _GotoState105Action,
	{_State397, AccessExprType}:                    _GotoState55Action,
	{_State397, IndexExprType}:                     _GotoState82Action,
	{_State397, PostfixableExprType}:               _GotoState95Action,
	{_State397, PostfixUnaryExprType}:              _GotoState94Action,
	{_State397, PrefixableExprType}:                _GotoState98Action,
	{_State397, PrefixUnaryExprType}:               _GotoState96Action,
	{_State397, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State397, MulExprType}:                       _GotoState89Action,
	{_State397, BinaryMulExprType}:                 _GotoState67Action,
	{_State397, AddExprType}:                       _GotoState57Action,
	{_State397, BinaryAddExprType}:                 _GotoState64Action,
	{_State397, CmpExprType}:                       _GotoState74Action,
	{_State397, BinaryCmpExprType}:                 _GotoState66Action,
	{_State397, AndExprType}:                       _GotoState58Action,
	{_State397, BinaryAndExprType}:                 _GotoState65Action,
	{_State397, OrExprType}:                        _GotoState92Action,
	{_State397, BinaryOrExprType}:                  _GotoState69Action,
	{_State397, InitializableTypeExprType}:         _GotoState83Action,
	{_State397, SliceTypeExprType}:                 _GotoState101Action,
	{_State397, ArrayTypeExprType}:                 _GotoState60Action,
	{_State397, MapTypeExprType}:                   _GotoState88Action,
	{_State397, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State397, AnonymousFuncExprType}:             _GotoState59Action,
	{_State398, IntegerLiteralToken}:               _GotoState40Action,
	{_State398, FloatLiteralToken}:                 _GotoState35Action,
	{_State398, RuneLiteralToken}:                  _GotoState48Action,
	{_State398, StringLiteralToken}:                _GotoState49Action,
	{_State398, IdentifierToken}:                   _GotoState38Action,
	{_State398, UnderscoreToken}:                   _GotoState53Action,
	{_State398, TrueToken}:                         _GotoState52Action,
	{_State398, FalseToken}:                        _GotoState34Action,
	{_State398, StructToken}:                       _GotoState50Action,
	{_State398, FuncToken}:                         _GotoState36Action,
	{_State398, VarToken}:                          _GotoState15Action,
	{_State398, LetToken}:                          _GotoState12Action,
	{_State398, NotToken}:                          _GotoState45Action,
	{_State398, LabelDeclToken}:                    _GotoState41Action,
	{_State398, LparenToken}:                       _GotoState43Action,
	{_State398, LbracketToken}:                     _GotoState42Action,
	{_State398, SubToken}:                          _GotoState51Action,
	{_State398, MulToken}:                          _GotoState44Action,
	{_State398, BitNegToken}:                       _GotoState27Action,
	{_State398, BitAndToken}:                       _GotoState26Action,
	{_State398, GreaterToken}:                      _GotoState37Action,
	{_State398, ParseErrorToken}:                   _GotoState46Action,
	{_State398, VarDeclPatternType}:                _GotoState104Action,
	{_State398, VarOrLetType}:                      _GotoState24Action,
	{_State398, OptionalLabelDeclType}:             _GotoState153Action,
	{_State398, SequenceExprType}:                  _GotoState446Action,
	{_State398, OptionalSequenceExprType}:          _GotoState445Action,
	{_State398, CallExprType}:                      _GotoState71Action,
	{_State398, AtomExprType}:                      _GotoState63Action,
	{_State398, ParseErrorExprType}:                _GotoState93Action,
	{_State398, LiteralExprType}:                   _GotoState87Action,
	{_State398, NamedExprType}:                     _GotoState90Action,
	{_State398, BlockExprType}:                     _GotoState70Action,
	{_State398, InitializeExprType}:                _GotoState84Action,
	{_State398, ImplicitStructExprType}:            _GotoState80Action,
	{_State398, AccessibleExprType}:                _GotoState105Action,
	{_State398, AccessExprType}:                    _GotoState55Action,
	{_State398, IndexExprType}:                     _GotoState82Action,
	{_State398, PostfixableExprType}:               _GotoState95Action,
	{_State398, PostfixUnaryExprType}:              _GotoState94Action,
	{_State398, PrefixableExprType}:                _GotoState98Action,
	{_State398, PrefixUnaryExprType}:               _GotoState96Action,
	{_State398, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State398, MulExprType}:                       _GotoState89Action,
	{_State398, BinaryMulExprType}:                 _GotoState67Action,
	{_State398, AddExprType}:                       _GotoState57Action,
	{_State398, BinaryAddExprType}:                 _GotoState64Action,
	{_State398, CmpExprType}:                       _GotoState74Action,
	{_State398, BinaryCmpExprType}:                 _GotoState66Action,
	{_State398, AndExprType}:                       _GotoState58Action,
	{_State398, BinaryAndExprType}:                 _GotoState65Action,
	{_State398, OrExprType}:                        _GotoState92Action,
	{_State398, BinaryOrExprType}:                  _GotoState69Action,
	{_State398, InitializableTypeExprType}:         _GotoState83Action,
	{_State398, SliceTypeExprType}:                 _GotoState101Action,
	{_State398, ArrayTypeExprType}:                 _GotoState60Action,
	{_State398, MapTypeExprType}:                   _GotoState88Action,
	{_State398, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State398, AnonymousFuncExprType}:             _GotoState59Action,
	{_State399, LbraceToken}:                       _GotoState11Action,
	{_State399, StatementBlockType}:                _GotoState447Action,
	{_State400, CommaToken}:                        _GotoState277Action,
	{_State400, AssignToken}:                       _GotoState448Action,
	{_State401, ElseToken}:                         _GotoState449Action,
	{_State404, IdentifierToken}:                   _GotoState237Action,
	{_State404, UnderscoreToken}:                   _GotoState238Action,
	{_State404, UnsafeToken}:                       _GotoState54Action,
	{_State404, StructToken}:                       _GotoState50Action,
	{_State404, EnumToken}:                         _GotoState110Action,
	{_State404, TraitToken}:                        _GotoState118Action,
	{_State404, FuncToken}:                         _GotoState236Action,
	{_State404, LparenToken}:                       _GotoState114Action,
	{_State404, LbracketToken}:                     _GotoState42Action,
	{_State404, DotToken}:                          _GotoState109Action,
	{_State404, QuestionToken}:                     _GotoState116Action,
	{_State404, ExclaimToken}:                      _GotoState111Action,
	{_State404, TildeTildeToken}:                   _GotoState117Action,
	{_State404, BitNegToken}:                       _GotoState108Action,
	{_State404, BitAndToken}:                       _GotoState107Action,
	{_State404, ParseErrorToken}:                   _GotoState115Action,
	{_State404, UnsafeStatementType}:               _GotoState246Action,
	{_State404, InitializableTypeExprType}:         _GotoState127Action,
	{_State404, SliceTypeExprType}:                 _GotoState101Action,
	{_State404, ArrayTypeExprType}:                 _GotoState60Action,
	{_State404, MapTypeExprType}:                   _GotoState88Action,
	{_State404, AtomTypeExprType}:                  _GotoState120Action,
	{_State404, NamedTypeExprType}:                 _GotoState128Action,
	{_State404, InferredTypeExprType}:              _GotoState126Action,
	{_State404, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State404, ReturnableTypeExprType}:            _GotoState132Action,
	{_State404, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State404, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State404, TypeExprType}:                      _GotoState245Action,
	{_State404, BinaryTypeExprType}:                _GotoState121Action,
	{_State404, FieldDefType}:                      _GotoState450Action,
	{_State404, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State404, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State404, TraitTypeExprType}:                 _GotoState133Action,
	{_State404, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State404, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State404, FuncTypeExprType}:                  _GotoState123Action,
	{_State404, MethodSignatureType}:               _GotoState242Action,
	{_State405, IdentifierToken}:                   _GotoState237Action,
	{_State405, UnderscoreToken}:                   _GotoState238Action,
	{_State405, UnsafeToken}:                       _GotoState54Action,
	{_State405, StructToken}:                       _GotoState50Action,
	{_State405, EnumToken}:                         _GotoState110Action,
	{_State405, TraitToken}:                        _GotoState118Action,
	{_State405, FuncToken}:                         _GotoState236Action,
	{_State405, LparenToken}:                       _GotoState114Action,
	{_State405, LbracketToken}:                     _GotoState42Action,
	{_State405, DotToken}:                          _GotoState109Action,
	{_State405, QuestionToken}:                     _GotoState116Action,
	{_State405, ExclaimToken}:                      _GotoState111Action,
	{_State405, TildeTildeToken}:                   _GotoState117Action,
	{_State405, BitNegToken}:                       _GotoState108Action,
	{_State405, BitAndToken}:                       _GotoState107Action,
	{_State405, ParseErrorToken}:                   _GotoState115Action,
	{_State405, UnsafeStatementType}:               _GotoState246Action,
	{_State405, InitializableTypeExprType}:         _GotoState127Action,
	{_State405, SliceTypeExprType}:                 _GotoState101Action,
	{_State405, ArrayTypeExprType}:                 _GotoState60Action,
	{_State405, MapTypeExprType}:                   _GotoState88Action,
	{_State405, AtomTypeExprType}:                  _GotoState120Action,
	{_State405, NamedTypeExprType}:                 _GotoState128Action,
	{_State405, InferredTypeExprType}:              _GotoState126Action,
	{_State405, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State405, ReturnableTypeExprType}:            _GotoState132Action,
	{_State405, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State405, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State405, TypeExprType}:                      _GotoState245Action,
	{_State405, BinaryTypeExprType}:                _GotoState121Action,
	{_State405, FieldDefType}:                      _GotoState451Action,
	{_State405, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State405, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State405, TraitTypeExprType}:                 _GotoState133Action,
	{_State405, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State405, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State405, FuncTypeExprType}:                  _GotoState123Action,
	{_State405, MethodSignatureType}:               _GotoState242Action,
	{_State406, IdentifierToken}:                   _GotoState237Action,
	{_State406, UnderscoreToken}:                   _GotoState238Action,
	{_State406, UnsafeToken}:                       _GotoState54Action,
	{_State406, StructToken}:                       _GotoState50Action,
	{_State406, EnumToken}:                         _GotoState110Action,
	{_State406, TraitToken}:                        _GotoState118Action,
	{_State406, FuncToken}:                         _GotoState236Action,
	{_State406, LparenToken}:                       _GotoState114Action,
	{_State406, LbracketToken}:                     _GotoState42Action,
	{_State406, DotToken}:                          _GotoState109Action,
	{_State406, QuestionToken}:                     _GotoState116Action,
	{_State406, ExclaimToken}:                      _GotoState111Action,
	{_State406, TildeTildeToken}:                   _GotoState117Action,
	{_State406, BitNegToken}:                       _GotoState108Action,
	{_State406, BitAndToken}:                       _GotoState107Action,
	{_State406, ParseErrorToken}:                   _GotoState115Action,
	{_State406, UnsafeStatementType}:               _GotoState246Action,
	{_State406, InitializableTypeExprType}:         _GotoState127Action,
	{_State406, SliceTypeExprType}:                 _GotoState101Action,
	{_State406, ArrayTypeExprType}:                 _GotoState60Action,
	{_State406, MapTypeExprType}:                   _GotoState88Action,
	{_State406, AtomTypeExprType}:                  _GotoState120Action,
	{_State406, NamedTypeExprType}:                 _GotoState128Action,
	{_State406, InferredTypeExprType}:              _GotoState126Action,
	{_State406, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State406, ReturnableTypeExprType}:            _GotoState132Action,
	{_State406, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State406, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State406, TypeExprType}:                      _GotoState245Action,
	{_State406, BinaryTypeExprType}:                _GotoState121Action,
	{_State406, FieldDefType}:                      _GotoState452Action,
	{_State406, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State406, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State406, TraitTypeExprType}:                 _GotoState133Action,
	{_State406, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State406, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State406, FuncTypeExprType}:                  _GotoState123Action,
	{_State406, MethodSignatureType}:               _GotoState242Action,
	{_State407, IdentifierToken}:                   _GotoState237Action,
	{_State407, UnderscoreToken}:                   _GotoState238Action,
	{_State407, UnsafeToken}:                       _GotoState54Action,
	{_State407, StructToken}:                       _GotoState50Action,
	{_State407, EnumToken}:                         _GotoState110Action,
	{_State407, TraitToken}:                        _GotoState118Action,
	{_State407, FuncToken}:                         _GotoState236Action,
	{_State407, LparenToken}:                       _GotoState114Action,
	{_State407, LbracketToken}:                     _GotoState42Action,
	{_State407, DotToken}:                          _GotoState109Action,
	{_State407, QuestionToken}:                     _GotoState116Action,
	{_State407, ExclaimToken}:                      _GotoState111Action,
	{_State407, TildeTildeToken}:                   _GotoState117Action,
	{_State407, BitNegToken}:                       _GotoState108Action,
	{_State407, BitAndToken}:                       _GotoState107Action,
	{_State407, ParseErrorToken}:                   _GotoState115Action,
	{_State407, UnsafeStatementType}:               _GotoState246Action,
	{_State407, InitializableTypeExprType}:         _GotoState127Action,
	{_State407, SliceTypeExprType}:                 _GotoState101Action,
	{_State407, ArrayTypeExprType}:                 _GotoState60Action,
	{_State407, MapTypeExprType}:                   _GotoState88Action,
	{_State407, AtomTypeExprType}:                  _GotoState120Action,
	{_State407, NamedTypeExprType}:                 _GotoState128Action,
	{_State407, InferredTypeExprType}:              _GotoState126Action,
	{_State407, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State407, ReturnableTypeExprType}:            _GotoState132Action,
	{_State407, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State407, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State407, TypeExprType}:                      _GotoState245Action,
	{_State407, BinaryTypeExprType}:                _GotoState121Action,
	{_State407, FieldDefType}:                      _GotoState453Action,
	{_State407, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State407, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State407, TraitTypeExprType}:                 _GotoState133Action,
	{_State407, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State407, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State407, FuncTypeExprType}:                  _GotoState123Action,
	{_State407, MethodSignatureType}:               _GotoState242Action,
	{_State408, AddToken}:                          _GotoState249Action,
	{_State408, SubToken}:                          _GotoState251Action,
	{_State408, MulToken}:                          _GotoState250Action,
	{_State408, BinaryTypeOpType}:                  _GotoState252Action,
	{_State409, IdentifierToken}:                   _GotoState113Action,
	{_State409, UnderscoreToken}:                   _GotoState119Action,
	{_State409, StructToken}:                       _GotoState50Action,
	{_State409, EnumToken}:                         _GotoState110Action,
	{_State409, TraitToken}:                        _GotoState118Action,
	{_State409, FuncToken}:                         _GotoState112Action,
	{_State409, LparenToken}:                       _GotoState114Action,
	{_State409, LbracketToken}:                     _GotoState42Action,
	{_State409, DotToken}:                          _GotoState109Action,
	{_State409, QuestionToken}:                     _GotoState116Action,
	{_State409, ExclaimToken}:                      _GotoState111Action,
	{_State409, TildeTildeToken}:                   _GotoState117Action,
	{_State409, BitNegToken}:                       _GotoState108Action,
	{_State409, BitAndToken}:                       _GotoState107Action,
	{_State409, ParseErrorToken}:                   _GotoState115Action,
	{_State409, InitializableTypeExprType}:         _GotoState127Action,
	{_State409, SliceTypeExprType}:                 _GotoState101Action,
	{_State409, ArrayTypeExprType}:                 _GotoState60Action,
	{_State409, MapTypeExprType}:                   _GotoState88Action,
	{_State409, AtomTypeExprType}:                  _GotoState120Action,
	{_State409, NamedTypeExprType}:                 _GotoState128Action,
	{_State409, InferredTypeExprType}:              _GotoState126Action,
	{_State409, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State409, ReturnableTypeExprType}:            _GotoState431Action,
	{_State409, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State409, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State409, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State409, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State409, TraitTypeExprType}:                 _GotoState133Action,
	{_State409, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State409, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State409, ReturnTypeType}:                    _GotoState454Action,
	{_State409, FuncTypeExprType}:                  _GotoState123Action,
	{_State410, IdentifierToken}:                   _GotoState331Action,
	{_State410, UnderscoreToken}:                   _GotoState332Action,
	{_State410, StructToken}:                       _GotoState50Action,
	{_State410, EnumToken}:                         _GotoState110Action,
	{_State410, TraitToken}:                        _GotoState118Action,
	{_State410, FuncToken}:                         _GotoState112Action,
	{_State410, LparenToken}:                       _GotoState114Action,
	{_State410, LbracketToken}:                     _GotoState42Action,
	{_State410, DotToken}:                          _GotoState109Action,
	{_State410, QuestionToken}:                     _GotoState116Action,
	{_State410, ExclaimToken}:                      _GotoState111Action,
	{_State410, EllipsisToken}:                     _GotoState330Action,
	{_State410, TildeTildeToken}:                   _GotoState117Action,
	{_State410, BitNegToken}:                       _GotoState108Action,
	{_State410, BitAndToken}:                       _GotoState107Action,
	{_State410, ParseErrorToken}:                   _GotoState115Action,
	{_State410, InitializableTypeExprType}:         _GotoState127Action,
	{_State410, SliceTypeExprType}:                 _GotoState101Action,
	{_State410, ArrayTypeExprType}:                 _GotoState60Action,
	{_State410, MapTypeExprType}:                   _GotoState88Action,
	{_State410, AtomTypeExprType}:                  _GotoState120Action,
	{_State410, NamedTypeExprType}:                 _GotoState128Action,
	{_State410, InferredTypeExprType}:              _GotoState126Action,
	{_State410, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State410, ReturnableTypeExprType}:            _GotoState132Action,
	{_State410, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State410, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State410, TypeExprType}:                      _GotoState337Action,
	{_State410, BinaryTypeExprType}:                _GotoState121Action,
	{_State410, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State410, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State410, TraitTypeExprType}:                 _GotoState133Action,
	{_State410, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State410, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State410, ProperParameterDefType}:            _GotoState336Action,
	{_State410, ParameterDeclType}:                 _GotoState455Action,
	{_State410, FuncTypeExprType}:                  _GotoState123Action,
	{_State412, IdentifierToken}:                   _GotoState331Action,
	{_State412, UnderscoreToken}:                   _GotoState332Action,
	{_State412, StructToken}:                       _GotoState50Action,
	{_State412, EnumToken}:                         _GotoState110Action,
	{_State412, TraitToken}:                        _GotoState118Action,
	{_State412, FuncToken}:                         _GotoState112Action,
	{_State412, LparenToken}:                       _GotoState114Action,
	{_State412, LbracketToken}:                     _GotoState42Action,
	{_State412, DotToken}:                          _GotoState109Action,
	{_State412, QuestionToken}:                     _GotoState116Action,
	{_State412, ExclaimToken}:                      _GotoState111Action,
	{_State412, EllipsisToken}:                     _GotoState330Action,
	{_State412, TildeTildeToken}:                   _GotoState117Action,
	{_State412, BitNegToken}:                       _GotoState108Action,
	{_State412, BitAndToken}:                       _GotoState107Action,
	{_State412, ParseErrorToken}:                   _GotoState115Action,
	{_State412, InitializableTypeExprType}:         _GotoState127Action,
	{_State412, SliceTypeExprType}:                 _GotoState101Action,
	{_State412, ArrayTypeExprType}:                 _GotoState60Action,
	{_State412, MapTypeExprType}:                   _GotoState88Action,
	{_State412, AtomTypeExprType}:                  _GotoState120Action,
	{_State412, NamedTypeExprType}:                 _GotoState128Action,
	{_State412, InferredTypeExprType}:              _GotoState126Action,
	{_State412, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State412, ReturnableTypeExprType}:            _GotoState132Action,
	{_State412, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State412, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State412, TypeExprType}:                      _GotoState337Action,
	{_State412, BinaryTypeExprType}:                _GotoState121Action,
	{_State412, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State412, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State412, TraitTypeExprType}:                 _GotoState133Action,
	{_State412, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State412, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State412, ProperParameterDefType}:            _GotoState336Action,
	{_State412, ParameterDeclType}:                 _GotoState333Action,
	{_State412, ProperParameterDeclsType}:          _GotoState335Action,
	{_State412, ParameterDeclsType}:                _GotoState456Action,
	{_State412, FuncTypeExprType}:                  _GotoState123Action,
	{_State419, AddToken}:                          _GotoState249Action,
	{_State419, SubToken}:                          _GotoState251Action,
	{_State419, MulToken}:                          _GotoState250Action,
	{_State419, BinaryTypeOpType}:                  _GotoState252Action,
	{_State421, IdentifierToken}:                   _GotoState354Action,
	{_State421, GenericParameterDefType}:           _GotoState457Action,
	{_State422, RparenToken}:                       _GotoState458Action,
	{_State423, AddToken}:                          _GotoState249Action,
	{_State423, SubToken}:                          _GotoState251Action,
	{_State423, MulToken}:                          _GotoState250Action,
	{_State423, BinaryTypeOpType}:                  _GotoState252Action,
	{_State424, AddToken}:                          _GotoState249Action,
	{_State424, SubToken}:                          _GotoState251Action,
	{_State424, MulToken}:                          _GotoState250Action,
	{_State424, BinaryTypeOpType}:                  _GotoState252Action,
	{_State425, LparenToken}:                       _GotoState459Action,
	{_State426, IdentifierToken}:                   _GotoState113Action,
	{_State426, UnderscoreToken}:                   _GotoState119Action,
	{_State426, StructToken}:                       _GotoState50Action,
	{_State426, EnumToken}:                         _GotoState110Action,
	{_State426, TraitToken}:                        _GotoState118Action,
	{_State426, FuncToken}:                         _GotoState112Action,
	{_State426, LparenToken}:                       _GotoState114Action,
	{_State426, LbracketToken}:                     _GotoState42Action,
	{_State426, DotToken}:                          _GotoState109Action,
	{_State426, QuestionToken}:                     _GotoState116Action,
	{_State426, ExclaimToken}:                      _GotoState111Action,
	{_State426, TildeTildeToken}:                   _GotoState117Action,
	{_State426, BitNegToken}:                       _GotoState108Action,
	{_State426, BitAndToken}:                       _GotoState107Action,
	{_State426, ParseErrorToken}:                   _GotoState115Action,
	{_State426, InitializableTypeExprType}:         _GotoState127Action,
	{_State426, SliceTypeExprType}:                 _GotoState101Action,
	{_State426, ArrayTypeExprType}:                 _GotoState60Action,
	{_State426, MapTypeExprType}:                   _GotoState88Action,
	{_State426, AtomTypeExprType}:                  _GotoState120Action,
	{_State426, NamedTypeExprType}:                 _GotoState128Action,
	{_State426, InferredTypeExprType}:              _GotoState126Action,
	{_State426, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State426, ReturnableTypeExprType}:            _GotoState132Action,
	{_State426, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State426, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State426, TypeExprType}:                      _GotoState460Action,
	{_State426, BinaryTypeExprType}:                _GotoState121Action,
	{_State426, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State426, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State426, TraitTypeExprType}:                 _GotoState133Action,
	{_State426, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State426, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State426, FuncTypeExprType}:                  _GotoState123Action,
	{_State430, LbraceToken}:                       _GotoState11Action,
	{_State430, StatementBlockType}:                _GotoState461Action,
	{_State440, AddToken}:                          _GotoState249Action,
	{_State440, SubToken}:                          _GotoState251Action,
	{_State440, MulToken}:                          _GotoState250Action,
	{_State440, BinaryTypeOpType}:                  _GotoState252Action,
	{_State444, DoToken}:                           _GotoState462Action,
	{_State445, SemicolonToken}:                    _GotoState463Action,
	{_State448, IntegerLiteralToken}:               _GotoState40Action,
	{_State448, FloatLiteralToken}:                 _GotoState35Action,
	{_State448, RuneLiteralToken}:                  _GotoState48Action,
	{_State448, StringLiteralToken}:                _GotoState49Action,
	{_State448, IdentifierToken}:                   _GotoState38Action,
	{_State448, UnderscoreToken}:                   _GotoState53Action,
	{_State448, TrueToken}:                         _GotoState52Action,
	{_State448, FalseToken}:                        _GotoState34Action,
	{_State448, StructToken}:                       _GotoState50Action,
	{_State448, FuncToken}:                         _GotoState36Action,
	{_State448, VarToken}:                          _GotoState15Action,
	{_State448, LetToken}:                          _GotoState12Action,
	{_State448, NotToken}:                          _GotoState45Action,
	{_State448, LabelDeclToken}:                    _GotoState41Action,
	{_State448, LparenToken}:                       _GotoState43Action,
	{_State448, LbracketToken}:                     _GotoState42Action,
	{_State448, SubToken}:                          _GotoState51Action,
	{_State448, MulToken}:                          _GotoState44Action,
	{_State448, BitNegToken}:                       _GotoState27Action,
	{_State448, BitAndToken}:                       _GotoState26Action,
	{_State448, GreaterToken}:                      _GotoState37Action,
	{_State448, ParseErrorToken}:                   _GotoState46Action,
	{_State448, VarDeclPatternType}:                _GotoState104Action,
	{_State448, VarOrLetType}:                      _GotoState24Action,
	{_State448, OptionalLabelDeclType}:             _GotoState153Action,
	{_State448, SequenceExprType}:                  _GotoState464Action,
	{_State448, CallExprType}:                      _GotoState71Action,
	{_State448, AtomExprType}:                      _GotoState63Action,
	{_State448, ParseErrorExprType}:                _GotoState93Action,
	{_State448, LiteralExprType}:                   _GotoState87Action,
	{_State448, NamedExprType}:                     _GotoState90Action,
	{_State448, BlockExprType}:                     _GotoState70Action,
	{_State448, InitializeExprType}:                _GotoState84Action,
	{_State448, ImplicitStructExprType}:            _GotoState80Action,
	{_State448, AccessibleExprType}:                _GotoState105Action,
	{_State448, AccessExprType}:                    _GotoState55Action,
	{_State448, IndexExprType}:                     _GotoState82Action,
	{_State448, PostfixableExprType}:               _GotoState95Action,
	{_State448, PostfixUnaryExprType}:              _GotoState94Action,
	{_State448, PrefixableExprType}:                _GotoState98Action,
	{_State448, PrefixUnaryExprType}:               _GotoState96Action,
	{_State448, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State448, MulExprType}:                       _GotoState89Action,
	{_State448, BinaryMulExprType}:                 _GotoState67Action,
	{_State448, AddExprType}:                       _GotoState57Action,
	{_State448, BinaryAddExprType}:                 _GotoState64Action,
	{_State448, CmpExprType}:                       _GotoState74Action,
	{_State448, BinaryCmpExprType}:                 _GotoState66Action,
	{_State448, AndExprType}:                       _GotoState58Action,
	{_State448, BinaryAndExprType}:                 _GotoState65Action,
	{_State448, OrExprType}:                        _GotoState92Action,
	{_State448, BinaryOrExprType}:                  _GotoState69Action,
	{_State448, InitializableTypeExprType}:         _GotoState83Action,
	{_State448, SliceTypeExprType}:                 _GotoState101Action,
	{_State448, ArrayTypeExprType}:                 _GotoState60Action,
	{_State448, MapTypeExprType}:                   _GotoState88Action,
	{_State448, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State448, AnonymousFuncExprType}:             _GotoState59Action,
	{_State449, IfToken}:                           _GotoState224Action,
	{_State449, LbraceToken}:                       _GotoState11Action,
	{_State449, StatementBlockType}:                _GotoState466Action,
	{_State449, IfExprType}:                        _GotoState465Action,
	{_State456, RparenToken}:                       _GotoState467Action,
	{_State458, IdentifierToken}:                   _GotoState113Action,
	{_State458, UnderscoreToken}:                   _GotoState119Action,
	{_State458, StructToken}:                       _GotoState50Action,
	{_State458, EnumToken}:                         _GotoState110Action,
	{_State458, TraitToken}:                        _GotoState118Action,
	{_State458, FuncToken}:                         _GotoState112Action,
	{_State458, LparenToken}:                       _GotoState114Action,
	{_State458, LbracketToken}:                     _GotoState42Action,
	{_State458, DotToken}:                          _GotoState109Action,
	{_State458, QuestionToken}:                     _GotoState116Action,
	{_State458, ExclaimToken}:                      _GotoState111Action,
	{_State458, TildeTildeToken}:                   _GotoState117Action,
	{_State458, BitNegToken}:                       _GotoState108Action,
	{_State458, BitAndToken}:                       _GotoState107Action,
	{_State458, ParseErrorToken}:                   _GotoState115Action,
	{_State458, InitializableTypeExprType}:         _GotoState127Action,
	{_State458, SliceTypeExprType}:                 _GotoState101Action,
	{_State458, ArrayTypeExprType}:                 _GotoState60Action,
	{_State458, MapTypeExprType}:                   _GotoState88Action,
	{_State458, AtomTypeExprType}:                  _GotoState120Action,
	{_State458, NamedTypeExprType}:                 _GotoState128Action,
	{_State458, InferredTypeExprType}:              _GotoState126Action,
	{_State458, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State458, ReturnableTypeExprType}:            _GotoState431Action,
	{_State458, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State458, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State458, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State458, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State458, TraitTypeExprType}:                 _GotoState133Action,
	{_State458, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State458, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State458, ReturnTypeType}:                    _GotoState468Action,
	{_State458, FuncTypeExprType}:                  _GotoState123Action,
	{_State459, IdentifierToken}:                   _GotoState256Action,
	{_State459, UnderscoreToken}:                   _GotoState257Action,
	{_State459, ProperParameterDefType}:            _GotoState259Action,
	{_State459, ParameterDefType}:                  _GotoState280Action,
	{_State459, ProperParameterDefsType}:           _GotoState282Action,
	{_State459, ParameterDefsType}:                 _GotoState469Action,
	{_State460, AddToken}:                          _GotoState249Action,
	{_State460, SubToken}:                          _GotoState251Action,
	{_State460, MulToken}:                          _GotoState250Action,
	{_State460, BinaryTypeOpType}:                  _GotoState252Action,
	{_State462, LbraceToken}:                       _GotoState11Action,
	{_State462, StatementBlockType}:                _GotoState470Action,
	{_State463, IntegerLiteralToken}:               _GotoState40Action,
	{_State463, FloatLiteralToken}:                 _GotoState35Action,
	{_State463, RuneLiteralToken}:                  _GotoState48Action,
	{_State463, StringLiteralToken}:                _GotoState49Action,
	{_State463, IdentifierToken}:                   _GotoState38Action,
	{_State463, UnderscoreToken}:                   _GotoState53Action,
	{_State463, TrueToken}:                         _GotoState52Action,
	{_State463, FalseToken}:                        _GotoState34Action,
	{_State463, StructToken}:                       _GotoState50Action,
	{_State463, FuncToken}:                         _GotoState36Action,
	{_State463, VarToken}:                          _GotoState15Action,
	{_State463, LetToken}:                          _GotoState12Action,
	{_State463, NotToken}:                          _GotoState45Action,
	{_State463, LabelDeclToken}:                    _GotoState41Action,
	{_State463, LparenToken}:                       _GotoState43Action,
	{_State463, LbracketToken}:                     _GotoState42Action,
	{_State463, SubToken}:                          _GotoState51Action,
	{_State463, MulToken}:                          _GotoState44Action,
	{_State463, BitNegToken}:                       _GotoState27Action,
	{_State463, BitAndToken}:                       _GotoState26Action,
	{_State463, GreaterToken}:                      _GotoState37Action,
	{_State463, ParseErrorToken}:                   _GotoState46Action,
	{_State463, VarDeclPatternType}:                _GotoState104Action,
	{_State463, VarOrLetType}:                      _GotoState24Action,
	{_State463, OptionalLabelDeclType}:             _GotoState153Action,
	{_State463, SequenceExprType}:                  _GotoState446Action,
	{_State463, OptionalSequenceExprType}:          _GotoState471Action,
	{_State463, CallExprType}:                      _GotoState71Action,
	{_State463, AtomExprType}:                      _GotoState63Action,
	{_State463, ParseErrorExprType}:                _GotoState93Action,
	{_State463, LiteralExprType}:                   _GotoState87Action,
	{_State463, NamedExprType}:                     _GotoState90Action,
	{_State463, BlockExprType}:                     _GotoState70Action,
	{_State463, InitializeExprType}:                _GotoState84Action,
	{_State463, ImplicitStructExprType}:            _GotoState80Action,
	{_State463, AccessibleExprType}:                _GotoState105Action,
	{_State463, AccessExprType}:                    _GotoState55Action,
	{_State463, IndexExprType}:                     _GotoState82Action,
	{_State463, PostfixableExprType}:               _GotoState95Action,
	{_State463, PostfixUnaryExprType}:              _GotoState94Action,
	{_State463, PrefixableExprType}:                _GotoState98Action,
	{_State463, PrefixUnaryExprType}:               _GotoState96Action,
	{_State463, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State463, MulExprType}:                       _GotoState89Action,
	{_State463, BinaryMulExprType}:                 _GotoState67Action,
	{_State463, AddExprType}:                       _GotoState57Action,
	{_State463, BinaryAddExprType}:                 _GotoState64Action,
	{_State463, CmpExprType}:                       _GotoState74Action,
	{_State463, BinaryCmpExprType}:                 _GotoState66Action,
	{_State463, AndExprType}:                       _GotoState58Action,
	{_State463, BinaryAndExprType}:                 _GotoState65Action,
	{_State463, OrExprType}:                        _GotoState92Action,
	{_State463, BinaryOrExprType}:                  _GotoState69Action,
	{_State463, InitializableTypeExprType}:         _GotoState83Action,
	{_State463, SliceTypeExprType}:                 _GotoState101Action,
	{_State463, ArrayTypeExprType}:                 _GotoState60Action,
	{_State463, MapTypeExprType}:                   _GotoState88Action,
	{_State463, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State463, AnonymousFuncExprType}:             _GotoState59Action,
	{_State467, IdentifierToken}:                   _GotoState113Action,
	{_State467, UnderscoreToken}:                   _GotoState119Action,
	{_State467, StructToken}:                       _GotoState50Action,
	{_State467, EnumToken}:                         _GotoState110Action,
	{_State467, TraitToken}:                        _GotoState118Action,
	{_State467, FuncToken}:                         _GotoState112Action,
	{_State467, LparenToken}:                       _GotoState114Action,
	{_State467, LbracketToken}:                     _GotoState42Action,
	{_State467, DotToken}:                          _GotoState109Action,
	{_State467, QuestionToken}:                     _GotoState116Action,
	{_State467, ExclaimToken}:                      _GotoState111Action,
	{_State467, TildeTildeToken}:                   _GotoState117Action,
	{_State467, BitNegToken}:                       _GotoState108Action,
	{_State467, BitAndToken}:                       _GotoState107Action,
	{_State467, ParseErrorToken}:                   _GotoState115Action,
	{_State467, InitializableTypeExprType}:         _GotoState127Action,
	{_State467, SliceTypeExprType}:                 _GotoState101Action,
	{_State467, ArrayTypeExprType}:                 _GotoState60Action,
	{_State467, MapTypeExprType}:                   _GotoState88Action,
	{_State467, AtomTypeExprType}:                  _GotoState120Action,
	{_State467, NamedTypeExprType}:                 _GotoState128Action,
	{_State467, InferredTypeExprType}:              _GotoState126Action,
	{_State467, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State467, ReturnableTypeExprType}:            _GotoState431Action,
	{_State467, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State467, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State467, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State467, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State467, TraitTypeExprType}:                 _GotoState133Action,
	{_State467, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State467, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State467, ReturnTypeType}:                    _GotoState472Action,
	{_State467, FuncTypeExprType}:                  _GotoState123Action,
	{_State468, LbraceToken}:                       _GotoState11Action,
	{_State468, StatementBlockType}:                _GotoState473Action,
	{_State469, RparenToken}:                       _GotoState474Action,
	{_State471, DoToken}:                           _GotoState475Action,
	{_State474, IdentifierToken}:                   _GotoState113Action,
	{_State474, UnderscoreToken}:                   _GotoState119Action,
	{_State474, StructToken}:                       _GotoState50Action,
	{_State474, EnumToken}:                         _GotoState110Action,
	{_State474, TraitToken}:                        _GotoState118Action,
	{_State474, FuncToken}:                         _GotoState112Action,
	{_State474, LparenToken}:                       _GotoState114Action,
	{_State474, LbracketToken}:                     _GotoState42Action,
	{_State474, DotToken}:                          _GotoState109Action,
	{_State474, QuestionToken}:                     _GotoState116Action,
	{_State474, ExclaimToken}:                      _GotoState111Action,
	{_State474, TildeTildeToken}:                   _GotoState117Action,
	{_State474, BitNegToken}:                       _GotoState108Action,
	{_State474, BitAndToken}:                       _GotoState107Action,
	{_State474, ParseErrorToken}:                   _GotoState115Action,
	{_State474, InitializableTypeExprType}:         _GotoState127Action,
	{_State474, SliceTypeExprType}:                 _GotoState101Action,
	{_State474, ArrayTypeExprType}:                 _GotoState60Action,
	{_State474, MapTypeExprType}:                   _GotoState88Action,
	{_State474, AtomTypeExprType}:                  _GotoState120Action,
	{_State474, NamedTypeExprType}:                 _GotoState128Action,
	{_State474, InferredTypeExprType}:              _GotoState126Action,
	{_State474, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State474, ReturnableTypeExprType}:            _GotoState431Action,
	{_State474, PrefixUnaryTypeExprType}:           _GotoState130Action,
	{_State474, PrefixUnaryTypeOpType}:             _GotoState131Action,
	{_State474, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State474, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State474, TraitTypeExprType}:                 _GotoState133Action,
	{_State474, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State474, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State474, ReturnTypeType}:                    _GotoState476Action,
	{_State474, FuncTypeExprType}:                  _GotoState123Action,
	{_State475, LbraceToken}:                       _GotoState11Action,
	{_State475, StatementBlockType}:                _GotoState477Action,
	{_State476, LbraceToken}:                       _GotoState11Action,
	{_State476, StatementBlockType}:                _GotoState478Action,
	{_State1, _EndMarker}:                          _ReduceNilToDefinitionsAction,
	{_State2, _WildcardMarker}:                     _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State3, _WildcardMarker}:                     _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State9, _WildcardMarker}:                     _ReduceCommentGroupsToDefinitionAction,
	{_State11, RbraceToken}:                        _ReduceNilToStatementsAction,
	{_State11, _WildcardMarker}:                    _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State12, _WildcardMarker}:                    _ReduceLetToVarOrLetAction,
	{_State13, _WildcardMarker}:                    _ReduceNoSpecToPackageDefAction,
	{_State15, _WildcardMarker}:                    _ReduceVarToVarOrLetAction,
	{_State16, _WildcardMarker}:                    _ReduceDefinitionToProperDefinitionsAction,
	{_State17, _EndMarker}:                         _ReduceToSourceAction,
	{_State18, _WildcardMarker}:                    _ReduceNamedFuncDefToDefinitionAction,
	{_State19, _WildcardMarker}:                    _ReducePackageDefToDefinitionAction,
	{_State20, _EndMarker}:                         _ReduceProperDefinitionsToDefinitionsAction,
	{_State21, _WildcardMarker}:                    _ReduceStatementBlockToDefinitionAction,
	{_State22, _WildcardMarker}:                    _ReduceTypeDefToDefinitionAction,
	{_State23, _WildcardMarker}:                    _ReduceGlobalVarDefToDefinitionAction,
	{_State25, _WildcardMarker}:                    _ReduceAsyncToCallbackOpAction,
	{_State26, _WildcardMarker}:                    _ReduceBitAndToPrefixUnaryOpAction,
	{_State27, _WildcardMarker}:                    _ReduceBitNegToPrefixUnaryOpAction,
	{_State28, _WildcardMarker}:                    _ReduceBreakToJumpTypeAction,
	{_State29, LbraceToken}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State30, _WildcardMarker}:                    _ReduceContinueToJumpTypeAction,
	{_State32, _WildcardMarker}:                    _ReduceDeferToCallbackOpAction,
	{_State33, _EndMarker}:                         _ReduceToFallthroughStatementAction,
	{_State34, _WildcardMarker}:                    _ReduceFalseToLiteralExprAction,
	{_State35, _WildcardMarker}:                    _ReduceFloatLiteralToLiteralExprAction,
	{_State37, LbraceToken}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State38, _WildcardMarker}:                    _ReduceIdentifierToNamedExprAction,
	{_State40, _WildcardMarker}:                    _ReduceIntegerLiteralToLiteralExprAction,
	{_State41, _WildcardMarker}:                    _ReduceLabelDeclToOptionalLabelDeclAction,
	{_State43, _WildcardMarker}:                    _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State43, RparenToken}:                        _ReduceNilToArgumentsAction,
	{_State44, _WildcardMarker}:                    _ReduceMulToPrefixUnaryOpAction,
	{_State45, _WildcardMarker}:                    _ReduceNotToPrefixUnaryOpAction,
	{_State46, _WildcardMarker}:                    _ReduceToParseErrorExprAction,
	{_State47, _WildcardMarker}:                    _ReduceReturnToJumpTypeAction,
	{_State48, _WildcardMarker}:                    _ReduceRuneLiteralToLiteralExprAction,
	{_State49, _WildcardMarker}:                    _ReduceStringLiteralToLiteralExprAction,
	{_State51, _WildcardMarker}:                    _ReduceSubToPrefixUnaryOpAction,
	{_State52, _WildcardMarker}:                    _ReduceTrueToLiteralExprAction,
	{_State53, _WildcardMarker}:                    _ReduceUnderscoreToNamedExprAction,
	{_State55, _WildcardMarker}:                    _ReduceAccessExprToAccessibleExprAction,
	{_State56, _WildcardMarker}:                    _ReduceAccessibleExprToPostfixableExprAction,
	{_State56, LparenToken}:                        _ReduceNilToGenericTypeArgumentsAction,
	{_State57, _WildcardMarker}:                    _ReduceAddExprToCmpExprAction,
	{_State58, _WildcardMarker}:                    _ReduceAndExprToOrExprAction,
	{_State59, _WildcardMarker}:                    _ReduceAnonymousFuncExprToAtomExprAction,
	{_State60, _WildcardMarker}:                    _ReduceArrayTypeExprToInitializableTypeExprAction,
	{_State62, _EndMarker}:                         _ReduceAssignStatementToSimpleStatementAction,
	{_State63, _WildcardMarker}:                    _ReduceAtomExprToAccessibleExprAction,
	{_State64, _WildcardMarker}:                    _ReduceBinaryAddExprToAddExprAction,
	{_State65, _WildcardMarker}:                    _ReduceBinaryAndExprToAndExprAction,
	{_State66, _WildcardMarker}:                    _ReduceBinaryCmpExprToCmpExprAction,
	{_State67, _WildcardMarker}:                    _ReduceBinaryMulExprToMulExprAction,
	{_State68, _EndMarker}:                         _ReduceBinaryOpAssignStatementToSimpleStatementAction,
	{_State69, _WildcardMarker}:                    _ReduceBinaryOrExprToOrExprAction,
	{_State70, _WildcardMarker}:                    _ReduceBlockExprToAtomExprAction,
	{_State71, _WildcardMarker}:                    _ReduceCallExprToAccessibleExprAction,
	{_State72, LbraceToken}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State73, _EndMarker}:                         _ReduceCallbackOpStatementToSimpleStatementAction,
	{_State74, _WildcardMarker}:                    _ReduceCmpExprToAndExprAction,
	{_State75, _WildcardMarker}:                    _ReduceExplicitStructTypeExprToInitializableTypeExprAction,
	{_State76, _WildcardMarker}:                    _ReduceExprToExprsAction,
	{_State77, _EndMarker}:                         _ReduceExprOrImproperStructStatementToSimpleStatementAction,
	{_State78, _WildcardMarker}:                    _ReduceToExprOrImproperStructStatementAction,
	{_State79, _EndMarker}:                         _ReduceFallthroughStatementToSimpleStatementAction,
	{_State80, _WildcardMarker}:                    _ReduceImplicitStructExprToAtomExprAction,
	{_State81, _EndMarker}:                         _ReduceImportStatementToStatementAction,
	{_State82, _WildcardMarker}:                    _ReduceIndexExprToAccessibleExprAction,
	{_State84, _WildcardMarker}:                    _ReduceInitializeExprToAtomExprAction,
	{_State85, _EndMarker}:                         _ReduceJumpStatementToSimpleStatementAction,
	{_State86, _EndMarker}:                         _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State86, NewlinesToken}:                      _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State86, RbraceToken}:                        _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State86, SemicolonToken}:                     _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State86, _WildcardMarker}:                    _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State87, _WildcardMarker}:                    _ReduceLiteralExprToAtomExprAction,
	{_State88, _WildcardMarker}:                    _ReduceMapTypeExprToInitializableTypeExprAction,
	{_State89, _WildcardMarker}:                    _ReduceMulExprToAddExprAction,
	{_State90, _WildcardMarker}:                    _ReduceNamedExprToAtomExprAction,
	{_State92, _WildcardMarker}:                    _ReduceOrExprToSequenceExprAction,
	{_State93, _WildcardMarker}:                    _ReduceParseErrorExprToAtomExprAction,
	{_State94, _WildcardMarker}:                    _ReducePostfixUnaryExprToPostfixableExprAction,
	{_State95, _WildcardMarker}:                    _ReducePostfixableExprToPrefixableExprAction,
	{_State96, _WildcardMarker}:                    _ReducePrefixUnaryExprToPrefixableExprAction,
	{_State97, LbraceToken}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State98, _WildcardMarker}:                    _ReducePrefixableExprToMulExprAction,
	{_State99, AssignToken}:                        _ReduceToAssignPatternAction,
	{_State99, _WildcardMarker}:                    _ReduceSequenceExprToExprAction,
	{_State100, _EndMarker}:                        _ReduceSimpleStatementToStatementAction,
	{_State101, _WildcardMarker}:                   _ReduceSliceTypeExprToInitializableTypeExprAction,
	{_State102, _EndMarker}:                        _ReduceUnaryOpAssignStatementToSimpleStatementAction,
	{_State103, _EndMarker}:                        _ReduceUnsafeStatementToSimpleStatementAction,
	{_State104, _EndMarker}:                        _ReduceVarDeclPatternToSequenceExprAction,
	{_State105, _WildcardMarker}:                   _ReduceAccessibleExprToPostfixableExprAction,
	{_State105, LparenToken}:                       _ReduceNilToGenericTypeArgumentsAction,
	{_State106, _EndMarker}:                        _ReduceSequenceExprToExprAction,
	{_State107, _WildcardMarker}:                   _ReduceBitAndToPrefixUnaryTypeOpAction,
	{_State108, _WildcardMarker}:                   _ReduceBitNegToPrefixUnaryTypeOpAction,
	{_State109, _WildcardMarker}:                   _ReduceDotToInferredTypeExprAction,
	{_State111, _WildcardMarker}:                   _ReduceExclaimToPrefixUnaryTypeOpAction,
	{_State113, _WildcardMarker}:                   _ReduceNilToGenericTypeArgumentsAction,
	{_State114, RparenToken}:                       _ReduceNilToImplicitFieldDefsAction,
	{_State115, _WildcardMarker}:                   _ReduceToParseErrorTypeExprAction,
	{_State116, _WildcardMarker}:                   _ReduceQuestionToPrefixUnaryTypeOpAction,
	{_State117, _WildcardMarker}:                   _ReduceTildeTildeToPrefixUnaryTypeOpAction,
	{_State119, _WildcardMarker}:                   _ReduceUnderscoreToInferredTypeExprAction,
	{_State120, _WildcardMarker}:                   _ReduceAtomTypeExprToReturnableTypeExprAction,
	{_State121, _WildcardMarker}:                   _ReduceBinaryTypeExprToTypeExprAction,
	{_State122, _WildcardMarker}:                   _ReduceExplicitEnumTypeExprToAtomTypeExprAction,
	{_State123, _WildcardMarker}:                   _ReduceFuncTypeExprToAtomTypeExprAction,
	{_State124, _WildcardMarker}:                   _ReduceImplicitEnumTypeExprToAtomTypeExprAction,
	{_State125, _WildcardMarker}:                   _ReduceImplicitStructTypeExprToAtomTypeExprAction,
	{_State126, _WildcardMarker}:                   _ReduceInferredTypeExprToAtomTypeExprAction,
	{_State127, _WildcardMarker}:                   _ReduceInitializableTypeExprToAtomTypeExprAction,
	{_State128, _WildcardMarker}:                   _ReduceNamedTypeExprToAtomTypeExprAction,
	{_State129, _WildcardMarker}:                   _ReduceParseErrorTypeExprToAtomTypeExprAction,
	{_State130, _WildcardMarker}:                   _ReducePrefixUnaryTypeExprToReturnableTypeExprAction,
	{_State132, _WildcardMarker}:                   _ReduceReturnableTypeExprToTypeExprAction,
	{_State133, _WildcardMarker}:                   _ReduceTraitTypeExprToAtomTypeExprAction,
	{_State134, LparenToken}:                       _ReduceNilToGenericParameterDefsAction,
	{_State136, RbraceToken}:                       _ReduceProperStatementsToStatementsAction,
	{_State137, _WildcardMarker}:                   _ReduceStatementToProperStatementsAction,
	{_State139, _WildcardMarker}:                   _ReduceWithSpecToPackageDefAction,
	{_State140, _WildcardMarker}:                   _ReduceNilToGenericParameterDefsAction,
	{_State141, _EndMarker}:                        _ReduceImproperToDefinitionsAction,
	{_State142, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State143, _WildcardMarker}:                   _ReduceIdentifierToVarPatternAction,
	{_State145, _WildcardMarker}:                   _ReduceUnderscoreToVarPatternAction,
	{_State146, _WildcardMarker}:                   _ReduceTuplePatternToVarPatternAction,
	{_State147, _WildcardMarker}:                   _ReduceNilToOptionalTypeExprAction,
	{_State149, _WildcardMarker}:                   _ReduceVarToVarOrLetAction,
	{_State150, _WildcardMarker}:                   _ReduceAssignPatternToCasePatternAction,
	{_State151, _WildcardMarker}:                   _ReduceCasePatternToCasePatternsAction,
	{_State154, _WildcardMarker}:                   _ReduceToAssignPatternAction,
	{_State155, _EndMarker}:                        _ReduceNilToOptionalSimpleStatementAction,
	{_State155, NewlinesToken}:                     _ReduceNilToOptionalSimpleStatementAction,
	{_State155, RbraceToken}:                       _ReduceNilToOptionalSimpleStatementAction,
	{_State155, SemicolonToken}:                    _ReduceNilToOptionalSimpleStatementAction,
	{_State155, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State156, RparenToken}:                       _ReduceNilToParameterDefsAction,
	{_State157, _EndMarker}:                        _ReduceAssignVarPatternToSequenceExprAction,
	{_State161, _EndMarker}:                        _ReduceStringLiteralToImportClauseAction,
	{_State163, _EndMarker}:                        _ReduceSingleToImportStatementAction,
	{_State165, ColonToken}:                        _ReduceUnitUnitPairToColonExprAction,
	{_State165, CommaToken}:                        _ReduceUnitUnitPairToColonExprAction,
	{_State165, RbracketToken}:                     _ReduceUnitUnitPairToColonExprAction,
	{_State165, RparenToken}:                       _ReduceUnitUnitPairToColonExprAction,
	{_State165, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State166, _WildcardMarker}:                   _ReduceSkipPatternToArgumentAction,
	{_State167, _WildcardMarker}:                   _ReduceIdentifierToNamedExprAction,
	{_State168, _WildcardMarker}:                   _ReduceArgumentToProperArgumentsAction,
	{_State170, _WildcardMarker}:                   _ReduceColonExprToArgumentAction,
	{_State171, _WildcardMarker}:                   _ReducePositionalToArgumentAction,
	{_State172, RparenToken}:                       _ReduceProperArgumentsToArgumentsAction,
	{_State173, RparenToken}:                       _ReduceNilToExplicitFieldDefsAction,
	{_State175, _WildcardMarker}:                   _ReduceAddAssignToBinaryOpAssignAction,
	{_State176, _EndMarker}:                        _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State177, _WildcardMarker}:                   _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State178, _WildcardMarker}:                   _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State179, _WildcardMarker}:                   _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State180, _WildcardMarker}:                   _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State181, _WildcardMarker}:                   _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State182, _WildcardMarker}:                   _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State183, _WildcardMarker}:                   _ReduceDivAssignToBinaryOpAssignAction,
	{_State184, RbracketToken}:                     _ReduceNilToTypeArgumentsAction,
	{_State186, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State187, _WildcardMarker}:                   _ReduceModAssignToBinaryOpAssignAction,
	{_State188, _WildcardMarker}:                   _ReduceMulAssignToBinaryOpAssignAction,
	{_State189, _WildcardMarker}:                   _ReduceToPostfixUnaryExprAction,
	{_State190, _WildcardMarker}:                   _ReduceSubAssignToBinaryOpAssignAction,
	{_State191, _EndMarker}:                        _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State192, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State194, _EndMarker}:                        _ReduceToUnaryOpAssignStatementAction,
	{_State195, _WildcardMarker}:                   _ReduceAddToAddOpAction,
	{_State196, _WildcardMarker}:                   _ReduceBitOrToAddOpAction,
	{_State197, _WildcardMarker}:                   _ReduceBitXorToAddOpAction,
	{_State198, _WildcardMarker}:                   _ReduceSubToAddOpAction,
	{_State199, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State200, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State201, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State202, LparenToken}:                       _ReduceNilToGenericTypeArgumentsAction,
	{_State203, _EndMarker}:                        _ReduceToCallbackOpStatementAction,
	{_State203, NewlinesToken}:                     _ReduceToCallbackOpStatementAction,
	{_State203, RbraceToken}:                       _ReduceToCallbackOpStatementAction,
	{_State203, SemicolonToken}:                    _ReduceToCallbackOpStatementAction,
	{_State203, _WildcardMarker}:                   _ReduceCallExprToAccessibleExprAction,
	{_State204, _WildcardMarker}:                   _ReduceEqualToCmpOpAction,
	{_State205, _WildcardMarker}:                   _ReduceGreaterToCmpOpAction,
	{_State206, _WildcardMarker}:                   _ReduceGreaterOrEqualToCmpOpAction,
	{_State207, _WildcardMarker}:                   _ReduceLessToCmpOpAction,
	{_State208, _WildcardMarker}:                   _ReduceLessOrEqualToCmpOpAction,
	{_State209, _WildcardMarker}:                   _ReduceNotEqualToCmpOpAction,
	{_State210, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State211, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State212, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State212, RparenToken}:                       _ReduceNilToArgumentsAction,
	{_State213, _EndMarker}:                        _ReduceLabeledNoValueToJumpStatementAction,
	{_State213, NewlinesToken}:                     _ReduceLabeledNoValueToJumpStatementAction,
	{_State213, RbraceToken}:                       _ReduceLabeledNoValueToJumpStatementAction,
	{_State213, SemicolonToken}:                    _ReduceLabeledNoValueToJumpStatementAction,
	{_State213, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State214, _EndMarker}:                        _ReduceUnlabeledValuedToJumpStatementAction,
	{_State215, _WildcardMarker}:                   _ReduceBitAndToMulOpAction,
	{_State216, _WildcardMarker}:                   _ReduceBitLshiftToMulOpAction,
	{_State217, _WildcardMarker}:                   _ReduceBitRshiftToMulOpAction,
	{_State218, _WildcardMarker}:                   _ReduceDivToMulOpAction,
	{_State219, _WildcardMarker}:                   _ReduceModToMulOpAction,
	{_State220, _WildcardMarker}:                   _ReduceMulToMulOpAction,
	{_State221, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State223, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State224, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State225, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State226, _EndMarker}:                        _ReduceIfExprToExprAction,
	{_State227, _EndMarker}:                        _ReduceLoopExprToExprAction,
	{_State228, _WildcardMarker}:                   _ReduceToBlockExprAction,
	{_State229, _EndMarker}:                        _ReduceSwitchExprToExprAction,
	{_State230, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State231, _WildcardMarker}:                   _ReduceToPrefixUnaryExprAction,
	{_State233, RparenToken}:                       _ReduceNilToParameterDeclsAction,
	{_State235, _WildcardMarker}:                   _ReduceLocalToNamedTypeExprAction,
	{_State237, _WildcardMarker}:                   _ReduceNilToGenericTypeArgumentsAction,
	{_State238, _WildcardMarker}:                   _ReduceUnderscoreToInferredTypeExprAction,
	{_State239, _WildcardMarker}:                   _ReduceFieldDefToProperImplicitFieldDefsAction,
	{_State242, _WildcardMarker}:                   _ReduceMethodSignatureToFieldDefAction,
	{_State243, RparenToken}:                       _ReduceProperImplicitEnumValueDefsToImplicitEnumValueDefsAction,
	{_State244, RparenToken}:                       _ReduceProperImplicitFieldDefsToImplicitFieldDefsAction,
	{_State245, _WildcardMarker}:                   _ReduceNilToOptionalDefaultAction,
	{_State246, _WildcardMarker}:                   _ReduceUnsafeStatementToFieldDefAction,
	{_State247, RparenToken}:                       _ReduceNilToExplicitFieldDefsAction,
	{_State248, _WildcardMarker}:                   _ReduceToPrefixUnaryTypeExprAction,
	{_State249, _WildcardMarker}:                   _ReduceAddToBinaryTypeOpAction,
	{_State250, _WildcardMarker}:                   _ReduceMulToBinaryTypeOpAction,
	{_State251, _WildcardMarker}:                   _ReduceSubToBinaryTypeOpAction,
	{_State253, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State254, RbracketToken}:                     _ReduceNilToGenericParameterDefListAction,
	{_State256, _WildcardMarker}:                   _ReduceNamedInferredArgToParameterDefAction,
	{_State257, _WildcardMarker}:                   _ReduceIgnoreInferredArgToParameterDefAction,
	{_State259, _WildcardMarker}:                   _ReduceProperParameterDefToParameterDefAction,
	{_State260, RbraceToken}:                       _ReduceImproperImplicitToStatementsAction,
	{_State260, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State261, RbraceToken}:                       _ReduceImproperExplicitToStatementsAction,
	{_State261, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State262, _EndMarker}:                        _ReduceToStatementBlockAction,
	{_State265, _WildcardMarker}:                   _ReduceAddToProperDefinitionsAction,
	{_State266, _WildcardMarker}:                   _ReduceGlobalVarAssignmentToDefinitionAction,
	{_State267, _WildcardMarker}:                   _ReduceEllipsisToFieldVarPatternAction,
	{_State268, _WildcardMarker}:                   _ReduceIdentifierToVarPatternAction,
	{_State269, _WildcardMarker}:                   _ReduceFieldVarPatternToFieldVarPatternsAction,
	{_State271, _WildcardMarker}:                   _ReducePositionalToFieldVarPatternAction,
	{_State272, _EndMarker}:                        _ReduceToVarDeclPatternAction,
	{_State273, _WildcardMarker}:                   _ReduceTypeExprToOptionalTypeExprAction,
	{_State274, _WildcardMarker}:                   _ReduceEnumNondataMatchPattenToCasePatternAction,
	{_State276, _EndMarker}:                        _ReduceNilToOptionalSimpleStatementAction,
	{_State276, NewlinesToken}:                     _ReduceNilToOptionalSimpleStatementAction,
	{_State276, RbraceToken}:                       _ReduceNilToOptionalSimpleStatementAction,
	{_State276, SemicolonToken}:                    _ReduceNilToOptionalSimpleStatementAction,
	{_State276, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State277, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State278, _EndMarker}:                        _ReduceDefaultBranchStatementToStatementAction,
	{_State279, _EndMarker}:                        _ReduceSimpleStatementToOptionalSimpleStatementAction,
	{_State280, _WildcardMarker}:                   _ReduceParameterDefToProperParameterDefsAction,
	{_State282, RparenToken}:                       _ReduceProperParameterDefsToParameterDefsAction,
	{_State283, _EndMarker}:                        _ReduceImportToLocalToImportClauseAction,
	{_State284, _EndMarker}:                        _ReduceAliasToImportClauseAction,
	{_State285, _WildcardMarker}:                   _ReduceImportClauseToProperImportClausesAction,
	{_State287, RparenToken}:                       _ReduceProperImportClausesToImportClausesAction,
	{_State288, _EndMarker}:                        _ReduceUnusableImportToImportClauseAction,
	{_State291, _WildcardMarker}:                   _ReduceToSliceTypeExprAction,
	{_State292, _WildcardMarker}:                   _ReduceUnitExprPairToColonExprAction,
	{_State293, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State294, _WildcardMarker}:                   _ReduceToImplicitStructExprAction,
	{_State295, ColonToken}:                        _ReduceColonExprUnitTupleToColonExprAction,
	{_State295, CommaToken}:                        _ReduceColonExprUnitTupleToColonExprAction,
	{_State295, RbracketToken}:                     _ReduceColonExprUnitTupleToColonExprAction,
	{_State295, RparenToken}:                       _ReduceColonExprUnitTupleToColonExprAction,
	{_State295, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State296, ColonToken}:                        _ReduceExprUnitPairToColonExprAction,
	{_State296, CommaToken}:                        _ReduceExprUnitPairToColonExprAction,
	{_State296, RbracketToken}:                     _ReduceExprUnitPairToColonExprAction,
	{_State296, RparenToken}:                       _ReduceExprUnitPairToColonExprAction,
	{_State296, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State297, _WildcardMarker}:                   _ReduceVarargAssignmentToArgumentAction,
	{_State298, RparenToken}:                       _ReduceImproperToArgumentsAction,
	{_State298, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State300, _WildcardMarker}:                   _ReduceFieldDefToProperExplicitFieldDefsAction,
	{_State301, RparenToken}:                       _ReduceProperExplicitFieldDefsToExplicitFieldDefsAction,
	{_State303, RbracketToken}:                     _ReduceProperTypeArgumentsToTypeArgumentsAction,
	{_State305, _WildcardMarker}:                   _ReduceTypeExprToProperTypeArgumentsAction,
	{_State306, _WildcardMarker}:                   _ReduceToAccessExprAction,
	{_State308, _EndMarker}:                        _ReduceToBinaryOpAssignStatementAction,
	{_State309, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State309, RparenToken}:                       _ReduceNilToArgumentsAction,
	{_State310, _WildcardMarker}:                   _ReduceToBinaryAddExprAction,
	{_State311, _WildcardMarker}:                   _ReduceToBinaryAndExprAction,
	{_State312, _EndMarker}:                        _ReduceToAssignStatementAction,
	{_State313, _WildcardMarker}:                   _ReduceToBinaryCmpExprAction,
	{_State314, _WildcardMarker}:                   _ReduceAddToExprsAction,
	{_State316, _EndMarker}:                        _ReduceLabeledValuedToJumpStatementAction,
	{_State317, _WildcardMarker}:                   _ReduceToBinaryMulExprAction,
	{_State318, _WildcardMarker}:                   _ReduceInfiniteToLoopExprAction,
	{_State321, _WildcardMarker}:                   _ReduceToAssignPatternAction,
	{_State321, SemicolonToken}:                    _ReduceSequenceExprToForAssignmentAction,
	{_State322, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State324, LbraceToken}:                       _ReduceSequenceExprToConditionAction,
	{_State326, _WildcardMarker}:                   _ReduceToBinaryOrExprAction,
	{_State329, RparenToken}:                       _ReduceProperExplicitEnumValueDefsToExplicitEnumValueDefsAction,
	{_State330, _WildcardMarker}:                   _ReduceUnnamedInferredVarargToParameterDeclAction,
	{_State331, _WildcardMarker}:                   _ReduceNilToGenericTypeArgumentsAction,
	{_State332, _WildcardMarker}:                   _ReduceUnderscoreToInferredTypeExprAction,
	{_State333, _WildcardMarker}:                   _ReduceParameterDeclToProperParameterDeclsAction,
	{_State335, RparenToken}:                       _ReduceProperParameterDeclsToParameterDeclsAction,
	{_State336, _WildcardMarker}:                   _ReduceProperParameterDefToParameterDeclAction,
	{_State337, _WildcardMarker}:                   _ReduceUnnamedTypedArgToParameterDeclAction,
	{_State338, _WildcardMarker}:                   _ReduceNilToGenericTypeArgumentsAction,
	{_State340, _WildcardMarker}:                   _ReduceDotToInferredTypeExprAction,
	{_State341, _WildcardMarker}:                   _ReduceNilToOptionalDefaultAction,
	{_State342, _WildcardMarker}:                   _ReduceStructPaddingToFieldDefAction,
	{_State344, _WildcardMarker}:                   _ReduceToImplicitEnumTypeExprAction,
	{_State345, _WildcardMarker}:                   _ReduceToImplicitStructTypeExprAction,
	{_State346, RparenToken}:                       _ReduceImproperToImplicitEnumValueDefsAction,
	{_State348, RparenToken}:                       _ReduceImproperToImplicitFieldDefsAction,
	{_State350, _WildcardMarker}:                   _ReduceUnnamedToFieldDefAction,
	{_State352, _WildcardMarker}:                   _ReduceToBinaryTypeExprAction,
	{_State353, _WildcardMarker}:                   _ReduceAliasToNamedFuncDefAction,
	{_State354, _WildcardMarker}:                   _ReduceUnconstrainedToGenericParameterDefAction,
	{_State355, _WildcardMarker}:                   _ReduceGenericParameterDefToProperGenericParameterDefListAction,
	{_State357, RbracketToken}:                     _ReduceProperGenericParameterDefListToGenericParameterDefListAction,
	{_State358, RparenToken}:                       _ReduceNilToParameterDefsAction,
	{_State359, _WildcardMarker}:                   _ReduceNamedInferredVarargToProperParameterDefAction,
	{_State360, _WildcardMarker}:                   _ReduceNamedTypedArgToProperParameterDefAction,
	{_State361, _WildcardMarker}:                   _ReduceIgnoreInferredVarargToProperParameterDefAction,
	{_State362, _WildcardMarker}:                   _ReduceIgnoreTypedArgToProperParameterDefAction,
	{_State364, _WildcardMarker}:                   _ReduceAddImplicitToProperStatementsAction,
	{_State365, _WildcardMarker}:                   _ReduceAddExplicitToProperStatementsAction,
	{_State366, _WildcardMarker}:                   _ReduceAliasToTypeDefAction,
	{_State367, _WildcardMarker}:                   _ReduceDefinitionToTypeDefAction,
	{_State370, _WildcardMarker}:                   _ReduceToTuplePatternAction,
	{_State371, _WildcardMarker}:                   _ReduceEnumMatchPatternToCasePatternAction,
	{_State373, _EndMarker}:                        _ReduceCaseBranchStatementToStatementAction,
	{_State374, _WildcardMarker}:                   _ReduceMultipleToCasePatternsAction,
	{_State375, LbraceToken}:                       _ReduceNilToReturnTypeAction,
	{_State376, RparenToken}:                       _ReduceImproperToParameterDefsAction,
	{_State377, _EndMarker}:                        _ReduceMultipleToImportStatementAction,
	{_State378, RparenToken}:                       _ReduceExplicitToImportClausesAction,
	{_State379, RparenToken}:                       _ReduceImplicitToImportClausesAction,
	{_State382, _WildcardMarker}:                   _ReduceNamedAssignmentToArgumentAction,
	{_State383, _WildcardMarker}:                   _ReduceColonExprExprTupleToColonExprAction,
	{_State384, _WildcardMarker}:                   _ReduceExprExprPairToColonExprAction,
	{_State385, _WildcardMarker}:                   _ReduceAddToProperArgumentsAction,
	{_State386, _WildcardMarker}:                   _ReduceToExplicitStructTypeExprAction,
	{_State387, RparenToken}:                       _ReduceImproperExplicitToExplicitFieldDefsAction,
	{_State388, RparenToken}:                       _ReduceImproperImplicitToExplicitFieldDefsAction,
	{_State390, RbracketToken}:                     _ReduceImproperToTypeArgumentsAction,
	{_State391, _WildcardMarker}:                   _ReduceBindingToGenericTypeArgumentsAction,
	{_State392, _WildcardMarker}:                   _ReduceToIndexExprAction,
	{_State394, _WildcardMarker}:                   _ReduceToInitializeExprAction,
	{_State395, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State396, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State397, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State398, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State398, SemicolonToken}:                    _ReduceNilToOptionalSequenceExprAction,
	{_State401, _WildcardMarker}:                   _ReduceNoElseToIfExprAction,
	{_State402, _EndMarker}:                        _ReduceToSwitchExprAction,
	{_State403, _WildcardMarker}:                   _ReduceToExplicitEnumTypeExprAction,
	{_State406, RparenToken}:                       _ReduceImproperToExplicitEnumValueDefsAction,
	{_State408, _WildcardMarker}:                   _ReduceUnnamedTypedVarargToParameterDeclAction,
	{_State409, _WildcardMarker}:                   _ReduceNilToReturnTypeAction,
	{_State410, RparenToken}:                       _ReduceImproperToParameterDeclsAction,
	{_State411, _WildcardMarker}:                   _ReduceExternalToNamedTypeExprAction,
	{_State412, RparenToken}:                       _ReduceNilToParameterDeclsAction,
	{_State413, _WildcardMarker}:                   _ReduceNamedToFieldDefAction,
	{_State414, _WildcardMarker}:                   _ReducePairToProperImplicitEnumValueDefsAction,
	{_State415, _WildcardMarker}:                   _ReduceAddToProperImplicitEnumValueDefsAction,
	{_State416, _WildcardMarker}:                   _ReduceAddToProperImplicitFieldDefsAction,
	{_State417, _WildcardMarker}:                   _ReduceDefaultToOptionalDefaultAction,
	{_State418, _WildcardMarker}:                   _ReduceToTraitTypeExprAction,
	{_State419, _WildcardMarker}:                   _ReduceConstrainedToGenericParameterDefAction,
	{_State420, _WildcardMarker}:                   _ReduceGenericToGenericParameterDefsAction,
	{_State421, RbracketToken}:                     _ReduceImproperToGenericParameterDefListAction,
	{_State423, _WildcardMarker}:                   _ReduceNamedTypedVarargToProperParameterDefAction,
	{_State424, _WildcardMarker}:                   _ReduceIgnoreTypedVarargToProperParameterDefAction,
	{_State427, _WildcardMarker}:                   _ReduceNamedToFieldVarPatternAction,
	{_State428, _WildcardMarker}:                   _ReduceAddToFieldVarPatternsAction,
	{_State429, _WildcardMarker}:                   _ReduceEnumVarDeclPatternToCasePatternAction,
	{_State431, _WildcardMarker}:                   _ReduceReturnableTypeExprToReturnTypeAction,
	{_State432, _WildcardMarker}:                   _ReduceAddToProperParameterDefsAction,
	{_State433, _WildcardMarker}:                   _ReduceAddExplicitToProperImportClausesAction,
	{_State434, _WildcardMarker}:                   _ReduceAddImplicitToProperImportClausesAction,
	{_State435, _WildcardMarker}:                   _ReduceToMapTypeExprAction,
	{_State436, _WildcardMarker}:                   _ReduceToArrayTypeExprAction,
	{_State437, _WildcardMarker}:                   _ReduceAddExplicitToProperExplicitFieldDefsAction,
	{_State438, _WildcardMarker}:                   _ReduceAddImplicitToProperExplicitFieldDefsAction,
	{_State439, _EndMarker}:                        _ReduceToUnsafeStatementAction,
	{_State440, _WildcardMarker}:                   _ReduceAddToProperTypeArgumentsAction,
	{_State441, _WildcardMarker}:                   _ReduceToCallExprAction,
	{_State442, _EndMarker}:                        _ReduceDoWhileToLoopExprAction,
	{_State443, SemicolonToken}:                    _ReduceAssignToForAssignmentAction,
	{_State446, DoToken}:                           _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State447, _EndMarker}:                        _ReduceWhileToLoopExprAction,
	{_State448, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State450, _WildcardMarker}:                   _ReduceImplicitPairToProperExplicitEnumValueDefsAction,
	{_State451, _WildcardMarker}:                   _ReduceExplicitPairToProperExplicitEnumValueDefsAction,
	{_State452, _WildcardMarker}:                   _ReduceImplicitAddToProperExplicitEnumValueDefsAction,
	{_State453, _WildcardMarker}:                   _ReduceExplicitAddToProperExplicitEnumValueDefsAction,
	{_State454, _WildcardMarker}:                   _ReduceToFuncTypeExprAction,
	{_State455, _WildcardMarker}:                   _ReduceAddToProperParameterDeclsAction,
	{_State457, _WildcardMarker}:                   _ReduceAddToProperGenericParameterDefListAction,
	{_State458, LbraceToken}:                       _ReduceNilToReturnTypeAction,
	{_State459, RparenToken}:                       _ReduceNilToParameterDefsAction,
	{_State460, _WildcardMarker}:                   _ReduceConstrainedDefToTypeDefAction,
	{_State461, _WildcardMarker}:                   _ReduceToAnonymousFuncExprAction,
	{_State463, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State463, DoToken}:                           _ReduceNilToOptionalSequenceExprAction,
	{_State464, LbraceToken}:                       _ReduceCaseToConditionAction,
	{_State465, _EndMarker}:                        _ReduceMultiIfElseToIfExprAction,
	{_State466, _EndMarker}:                        _ReduceIfElseToIfExprAction,
	{_State467, _WildcardMarker}:                   _ReduceNilToReturnTypeAction,
	{_State470, _EndMarker}:                        _ReduceIteratorToLoopExprAction,
	{_State472, _WildcardMarker}:                   _ReduceToMethodSignatureAction,
	{_State473, _WildcardMarker}:                   _ReduceFuncDefToNamedFuncDefAction,
	{_State474, LbraceToken}:                       _ReduceNilToReturnTypeAction,
	{_State477, _EndMarker}:                        _ReduceForToLoopExprAction,
	{_State478, _WildcardMarker}:                   _ReduceMethodDefToNamedFuncDefAction,
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
      UNDERSCORE -> State 53
      TRUE -> State 52
      FALSE -> State 34
      CASE -> State 29
      DEFAULT -> State 31
      RETURN -> State 47
      BREAK -> State 28
      CONTINUE -> State 30
      FALLTHROUGH -> State 33
      IMPORT -> State 39
      UNSAFE -> State 54
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
      simple_statement -> State 100
      statement -> State 6
      expr_or_improper_struct_statement -> State 77
      exprs -> State 78
      callback_op -> State 72
      callback_op_statement -> State 73
      unsafe_statement -> State 103
      jump_statement -> State 85
      jump_type -> State 86
      fallthrough_statement -> State 79
      assign_statement -> State 62
      unary_op_assign_statement -> State 102
      binary_op_assign_statement -> State 68
      import_statement -> State 81
      var_decl_pattern -> State 104
      var_or_let -> State 24
      assign_pattern -> State 61
      expr -> State 76
      optional_label_decl -> State 91
      sequence_expr -> State 99
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 56
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 3:
    Kernel Items:
      #accept: ^.expr
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 7
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 4:
    Kernel Items:
      #accept: ^.type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 8
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

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
      #accept: ^ expr., $
    Reduce:
      $ -> [#accept]
    Goto:
      (nil)

  State 8:
    Kernel Items:
      #accept: ^ type_expr., $
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      $ -> [#accept]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 9:
    Kernel Items:
      definition: COMMENT_GROUPS., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 10:
    Kernel Items:
      named_func_def: FUNC.IDENTIFIER generic_parameter_defs LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.IDENTIFIER ASSIGN expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 134
      LPAREN -> State 135

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
      UNDERSCORE -> State 53
      TRUE -> State 52
      FALSE -> State 34
      CASE -> State 29
      DEFAULT -> State 31
      RETURN -> State 47
      BREAK -> State 28
      CONTINUE -> State 30
      FALLTHROUGH -> State 33
      IMPORT -> State 39
      UNSAFE -> State 54
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
      proper_statements -> State 136
      statements -> State 138
      simple_statement -> State 100
      statement -> State 137
      expr_or_improper_struct_statement -> State 77
      exprs -> State 78
      callback_op -> State 72
      callback_op_statement -> State 73
      unsafe_statement -> State 103
      jump_statement -> State 85
      jump_type -> State 86
      fallthrough_statement -> State 79
      assign_statement -> State 62
      unary_op_assign_statement -> State 102
      binary_op_assign_statement -> State 68
      import_statement -> State 81
      var_decl_pattern -> State 104
      var_or_let -> State 24
      assign_pattern -> State 61
      expr -> State 76
      optional_label_decl -> State 91
      sequence_expr -> State 99
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 56
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

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
      statement_block -> State 139

  State 14:
    Kernel Items:
      type_def: TYPE.IDENTIFIER generic_parameter_defs type_expr
      type_def: TYPE.IDENTIFIER generic_parameter_defs type_expr IMPLEMENTS type_expr
      type_def: TYPE.IDENTIFIER ASSIGN type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 140

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
      NEWLINES -> State 141

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
      definition: var_decl_pattern.ASSIGN expr
    Reduce:
      * -> [definition]
    Goto:
      ASSIGN -> State 142

  State 24:
    Kernel Items:
      var_decl_pattern: var_or_let.var_pattern optional_type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 143
      UNDERSCORE -> State 145
      LPAREN -> State 144
      var_pattern -> State 147
      tuple_pattern -> State 146

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
      UNDERSCORE -> State 53
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 149
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      DOT -> State 148
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      case_patterns -> State 152
      var_decl_pattern -> State 104
      var_or_let -> State 24
      assign_pattern -> State 150
      case_pattern -> State 151
      optional_label_decl -> State 153
      sequence_expr -> State 154
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

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
      COLON -> State 155

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
      LPAREN -> State 156

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
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      optional_label_decl -> State 153
      sequence_expr -> State 157
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 38:
    Kernel Items:
      named_expr: IDENTIFIER., *
    Reduce:
      * -> [named_expr]
    Goto:
      (nil)

  State 39:
    Kernel Items:
      import_statement: IMPORT.import_clause
      import_statement: IMPORT.LPAREN import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 161
      IDENTIFIER -> State 159
      UNDERSCORE -> State 162
      LPAREN -> State 160
      DOT -> State 158
      import_clause -> State 163

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
      slice_type_expr: LBRACKET.type_expr RBRACKET
      array_type_expr: LBRACKET.type_expr COMMA INTEGER_LITERAL RBRACKET
      map_type_expr: LBRACKET.type_expr COLON type_expr RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 164
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

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
      IDENTIFIER -> State 167
      UNDERSCORE -> State 53
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
      COLON -> State 165
      ELLIPSIS -> State 166
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 171
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      proper_arguments -> State 172
      arguments -> State 169
      argument -> State 168
      colon_expr -> State 170
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

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
      explicit_struct_type_expr: STRUCT.LPAREN explicit_field_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 173

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
      named_expr: UNDERSCORE., *
    Reduce:
      * -> [named_expr]
    Goto:
      (nil)

  State 54:
    Kernel Items:
      unsafe_statement: UNSAFE.LESS IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      LESS -> State 174

  State 55:
    Kernel Items:
      accessible_expr: access_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 56:
    Kernel Items:
      unary_op_assign_statement: accessible_expr.unary_op_assign
      binary_op_assign_statement: accessible_expr.binary_op_assign expr
      call_expr: accessible_expr.generic_type_arguments LPAREN arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
      postfixable_expr: accessible_expr., *
      postfix_unary_expr: accessible_expr.QUESTION
    Reduce:
      * -> [postfixable_expr]
      LPAREN -> [generic_type_arguments]
    Goto:
      LBRACKET -> State 186
      DOT -> State 185
      QUESTION -> State 189
      DOLLAR_LBRACKET -> State 184
      ADD_ASSIGN -> State 175
      SUB_ASSIGN -> State 190
      MUL_ASSIGN -> State 188
      DIV_ASSIGN -> State 183
      MOD_ASSIGN -> State 187
      ADD_ONE_ASSIGN -> State 176
      SUB_ONE_ASSIGN -> State 191
      BIT_NEG_ASSIGN -> State 179
      BIT_AND_ASSIGN -> State 177
      BIT_OR_ASSIGN -> State 180
      BIT_XOR_ASSIGN -> State 182
      BIT_LSHIFT_ASSIGN -> State 178
      BIT_RSHIFT_ASSIGN -> State 181
      unary_op_assign -> State 194
      binary_op_assign -> State 192
      generic_type_arguments -> State 193

  State 57:
    Kernel Items:
      binary_add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 195
      SUB -> State 198
      BIT_XOR -> State 197
      BIT_OR -> State 196
      add_op -> State 199

  State 58:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 200

  State 59:
    Kernel Items:
      atom_expr: anonymous_func_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 60:
    Kernel Items:
      initializable_type_expr: array_type_expr., *
    Reduce:
      * -> [initializable_type_expr]
    Goto:
      (nil)

  State 61:
    Kernel Items:
      assign_statement: assign_pattern.ASSIGN expr
    Reduce:
      (nil)
    Goto:
      ASSIGN -> State 201

  State 62:
    Kernel Items:
      simple_statement: assign_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 63:
    Kernel Items:
      accessible_expr: atom_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 64:
    Kernel Items:
      add_expr: binary_add_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      (nil)

  State 65:
    Kernel Items:
      and_expr: binary_and_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      (nil)

  State 66:
    Kernel Items:
      cmp_expr: binary_cmp_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      (nil)

  State 67:
    Kernel Items:
      mul_expr: binary_mul_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 68:
    Kernel Items:
      simple_statement: binary_op_assign_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 69:
    Kernel Items:
      or_expr: binary_or_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      (nil)

  State 70:
    Kernel Items:
      atom_expr: block_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 71:
    Kernel Items:
      accessible_expr: call_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 72:
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
      UNDERSCORE -> State 53
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      PARSE_ERROR -> State 46
      optional_label_decl -> State 153
      call_expr -> State 203
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 202
      access_expr -> State 55
      index_expr -> State 82
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 73:
    Kernel Items:
      simple_statement: callback_op_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 74:
    Kernel Items:
      binary_cmp_expr: cmp_expr.cmp_op add_expr
      and_expr: cmp_expr., *
    Reduce:
      * -> [and_expr]
    Goto:
      EQUAL -> State 204
      NOT_EQUAL -> State 209
      LESS -> State 207
      LESS_OR_EQUAL -> State 208
      GREATER -> State 205
      GREATER_OR_EQUAL -> State 206
      cmp_op -> State 210

  State 75:
    Kernel Items:
      initializable_type_expr: explicit_struct_type_expr., *
    Reduce:
      * -> [initializable_type_expr]
    Goto:
      (nil)

  State 76:
    Kernel Items:
      exprs: expr., *
    Reduce:
      * -> [exprs]
    Goto:
      (nil)

  State 77:
    Kernel Items:
      simple_statement: expr_or_improper_struct_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 78:
    Kernel Items:
      expr_or_improper_struct_statement: exprs., *
      exprs: exprs.COMMA expr
    Reduce:
      * -> [expr_or_improper_struct_statement]
    Goto:
      COMMA -> State 211

  State 79:
    Kernel Items:
      simple_statement: fallthrough_statement., $
    Reduce:
      $ -> [simple_statement]
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
      initialize_expr: initializable_type_expr.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 212

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
      jump_statement: jump_type.exprs
      jump_statement: jump_type.JUMP_LABEL
      jump_statement: jump_type.JUMP_LABEL exprs
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
      UNDERSCORE -> State 53
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 15
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      JUMP_LABEL -> State 213
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      exprs -> State 214
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 76
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 87:
    Kernel Items:
      atom_expr: literal_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 88:
    Kernel Items:
      initializable_type_expr: map_type_expr., *
    Reduce:
      * -> [initializable_type_expr]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      binary_mul_expr: mul_expr.mul_op prefixable_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 220
      DIV -> State 218
      MOD -> State 219
      BIT_AND -> State 215
      BIT_LSHIFT -> State 216
      BIT_RSHIFT -> State 217
      mul_op -> State 221

  State 90:
    Kernel Items:
      atom_expr: named_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 91:
    Kernel Items:
      expr: optional_label_decl.if_expr
      expr: optional_label_decl.switch_expr
      expr: optional_label_decl.loop_expr
      block_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    Goto:
      IF -> State 224
      SWITCH -> State 225
      FOR -> State 223
      DO -> State 222
      LBRACE -> State 11
      statement_block -> State 228
      if_expr -> State 226
      switch_expr -> State 229
      loop_expr -> State 227

  State 92:
    Kernel Items:
      sequence_expr: or_expr., *
      binary_or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    Goto:
      OR -> State 230

  State 93:
    Kernel Items:
      atom_expr: parse_error_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 94:
    Kernel Items:
      postfixable_expr: postfix_unary_expr., *
    Reduce:
      * -> [postfixable_expr]
    Goto:
      (nil)

  State 95:
    Kernel Items:
      prefixable_expr: postfixable_expr., *
    Reduce:
      * -> [prefixable_expr]
    Goto:
      (nil)

  State 96:
    Kernel Items:
      prefixable_expr: prefix_unary_expr., *
    Reduce:
      * -> [prefixable_expr]
    Goto:
      (nil)

  State 97:
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
      UNDERSCORE -> State 53
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
      optional_label_decl -> State 153
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 231
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 98:
    Kernel Items:
      mul_expr: prefixable_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 99:
    Kernel Items:
      assign_pattern: sequence_expr., ASSIGN
      expr: sequence_expr., *
    Reduce:
      * -> [expr]
      ASSIGN -> [assign_pattern]
    Goto:
      (nil)

  State 100:
    Kernel Items:
      statement: simple_statement., $
    Reduce:
      $ -> [statement]
    Goto:
      (nil)

  State 101:
    Kernel Items:
      initializable_type_expr: slice_type_expr., *
    Reduce:
      * -> [initializable_type_expr]
    Goto:
      (nil)

  State 102:
    Kernel Items:
      simple_statement: unary_op_assign_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 103:
    Kernel Items:
      simple_statement: unsafe_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 104:
    Kernel Items:
      sequence_expr: var_decl_pattern., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 105:
    Kernel Items:
      call_expr: accessible_expr.generic_type_arguments LPAREN arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
      postfixable_expr: accessible_expr., *
      postfix_unary_expr: accessible_expr.QUESTION
    Reduce:
      * -> [postfixable_expr]
      LPAREN -> [generic_type_arguments]
    Goto:
      LBRACKET -> State 186
      DOT -> State 185
      QUESTION -> State 189
      DOLLAR_LBRACKET -> State 184
      generic_type_arguments -> State 193

  State 106:
    Kernel Items:
      expr: sequence_expr., $
    Reduce:
      $ -> [expr]
    Goto:
      (nil)

  State 107:
    Kernel Items:
      prefix_unary_type_op: BIT_AND., *
    Reduce:
      * -> [prefix_unary_type_op]
    Goto:
      (nil)

  State 108:
    Kernel Items:
      prefix_unary_type_op: BIT_NEG., *
    Reduce:
      * -> [prefix_unary_type_op]
    Goto:
      (nil)

  State 109:
    Kernel Items:
      inferred_type_expr: DOT., *
    Reduce:
      * -> [inferred_type_expr]
    Goto:
      (nil)

  State 110:
    Kernel Items:
      explicit_enum_type_expr: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 232

  State 111:
    Kernel Items:
      prefix_unary_type_op: EXCLAIM., *
    Reduce:
      * -> [prefix_unary_type_op]
    Goto:
      (nil)

  State 112:
    Kernel Items:
      func_type_expr: FUNC.LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 233

  State 113:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_type_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_type_arguments
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      DOT -> State 234
      DOLLAR_LBRACKET -> State 184
      generic_type_arguments -> State 235

  State 114:
    Kernel Items:
      implicit_struct_type_expr: LPAREN.implicit_field_defs RPAREN
      implicit_enum_type_expr: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      IDENTIFIER -> State 237
      UNDERSCORE -> State 238
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 236
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      unsafe_statement -> State 246
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 239
      proper_implicit_field_defs -> State 244
      implicit_field_defs -> State 241
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      proper_implicit_enum_value_defs -> State 243
      implicit_enum_value_defs -> State 240
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 115:
    Kernel Items:
      parse_error_type_expr: PARSE_ERROR., *
    Reduce:
      * -> [parse_error_type_expr]
    Goto:
      (nil)

  State 116:
    Kernel Items:
      prefix_unary_type_op: QUESTION., *
    Reduce:
      * -> [prefix_unary_type_op]
    Goto:
      (nil)

  State 117:
    Kernel Items:
      prefix_unary_type_op: TILDE_TILDE., *
    Reduce:
      * -> [prefix_unary_type_op]
    Goto:
      (nil)

  State 118:
    Kernel Items:
      trait_type_expr: TRAIT.LPAREN explicit_field_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 247

  State 119:
    Kernel Items:
      inferred_type_expr: UNDERSCORE., *
    Reduce:
      * -> [inferred_type_expr]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      returnable_type_expr: atom_type_expr., *
    Reduce:
      * -> [returnable_type_expr]
    Goto:
      (nil)

  State 121:
    Kernel Items:
      type_expr: binary_type_expr., *
    Reduce:
      * -> [type_expr]
    Goto:
      (nil)

  State 122:
    Kernel Items:
      atom_type_expr: explicit_enum_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 123:
    Kernel Items:
      atom_type_expr: func_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      atom_type_expr: implicit_enum_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      atom_type_expr: implicit_struct_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 126:
    Kernel Items:
      atom_type_expr: inferred_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 127:
    Kernel Items:
      atom_type_expr: initializable_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 128:
    Kernel Items:
      atom_type_expr: named_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 129:
    Kernel Items:
      atom_type_expr: parse_error_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 130:
    Kernel Items:
      returnable_type_expr: prefix_unary_type_expr., *
    Reduce:
      * -> [returnable_type_expr]
    Goto:
      (nil)

  State 131:
    Kernel Items:
      prefix_unary_type_expr: prefix_unary_type_op.returnable_type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 248
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 132:
    Kernel Items:
      type_expr: returnable_type_expr., *
    Reduce:
      * -> [type_expr]
    Goto:
      (nil)

  State 133:
    Kernel Items:
      atom_type_expr: trait_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 134:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER.generic_parameter_defs LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC IDENTIFIER.ASSIGN expr
    Reduce:
      LPAREN -> [generic_parameter_defs]
    Goto:
      DOLLAR_LBRACKET -> State 254
      ASSIGN -> State 253
      generic_parameter_defs -> State 255

  State 135:
    Kernel Items:
      named_func_def: FUNC LPAREN.parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 256
      UNDERSCORE -> State 257
      proper_parameter_def -> State 259
      parameter_def -> State 258

  State 136:
    Kernel Items:
      proper_statements: proper_statements.NEWLINES statement
      proper_statements: proper_statements.SEMICOLON statement
      statements: proper_statements., RBRACE
      statements: proper_statements.NEWLINES
      statements: proper_statements.SEMICOLON
    Reduce:
      RBRACE -> [statements]
    Goto:
      NEWLINES -> State 260
      SEMICOLON -> State 261

  State 137:
    Kernel Items:
      proper_statements: statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 138:
    Kernel Items:
      statement_block: LBRACE statements.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 262

  State 139:
    Kernel Items:
      package_def: PACKAGE statement_block., *
    Reduce:
      * -> [package_def]
    Goto:
      (nil)

  State 140:
    Kernel Items:
      type_def: TYPE IDENTIFIER.generic_parameter_defs type_expr
      type_def: TYPE IDENTIFIER.generic_parameter_defs type_expr IMPLEMENTS type_expr
      type_def: TYPE IDENTIFIER.ASSIGN type_expr
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      DOLLAR_LBRACKET -> State 254
      ASSIGN -> State 263
      generic_parameter_defs -> State 264

  State 141:
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
      definition -> State 265
      statement_block -> State 21
      var_decl_pattern -> State 23
      var_or_let -> State 24
      type_def -> State 22
      named_func_def -> State 18
      package_def -> State 19

  State 142:
    Kernel Items:
      definition: var_decl_pattern ASSIGN.expr
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 266
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 143:
    Kernel Items:
      var_pattern: IDENTIFIER., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 144:
    Kernel Items:
      tuple_pattern: LPAREN.field_var_patterns RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 268
      UNDERSCORE -> State 145
      LPAREN -> State 144
      ELLIPSIS -> State 267
      var_pattern -> State 271
      tuple_pattern -> State 146
      field_var_patterns -> State 270
      field_var_pattern -> State 269

  State 145:
    Kernel Items:
      var_pattern: UNDERSCORE., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 146:
    Kernel Items:
      var_pattern: tuple_pattern., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 147:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern.optional_type_expr
    Reduce:
      * -> [optional_type_expr]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      optional_type_expr -> State 272
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 273
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 148:
    Kernel Items:
      case_pattern: DOT.IDENTIFIER implicit_struct_expr
      case_pattern: DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 274

  State 149:
    Kernel Items:
      var_or_let: VAR., *
      case_pattern: VAR.DOT IDENTIFIER tuple_pattern
    Reduce:
      * -> [var_or_let]
    Goto:
      DOT -> State 275

  State 150:
    Kernel Items:
      case_pattern: assign_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 151:
    Kernel Items:
      case_patterns: case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 152:
    Kernel Items:
      statement: CASE case_patterns.COLON optional_simple_statement
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    Goto:
      COMMA -> State 277
      COLON -> State 276

  State 153:
    Kernel Items:
      block_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 228

  State 154:
    Kernel Items:
      assign_pattern: sequence_expr., *
    Reduce:
      * -> [assign_pattern]
    Goto:
      (nil)

  State 155:
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
      UNDERSCORE -> State 53
      TRUE -> State 52
      FALSE -> State 34
      RETURN -> State 47
      BREAK -> State 28
      CONTINUE -> State 30
      FALLTHROUGH -> State 33
      UNSAFE -> State 54
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
      simple_statement -> State 279
      optional_simple_statement -> State 278
      expr_or_improper_struct_statement -> State 77
      exprs -> State 78
      callback_op -> State 72
      callback_op_statement -> State 73
      unsafe_statement -> State 103
      jump_statement -> State 85
      jump_type -> State 86
      fallthrough_statement -> State 79
      assign_statement -> State 62
      unary_op_assign_statement -> State 102
      binary_op_assign_statement -> State 68
      var_decl_pattern -> State 104
      var_or_let -> State 24
      assign_pattern -> State 61
      expr -> State 76
      optional_label_decl -> State 91
      sequence_expr -> State 99
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 56
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 156:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 256
      UNDERSCORE -> State 257
      proper_parameter_def -> State 259
      parameter_def -> State 280
      proper_parameter_defs -> State 282
      parameter_defs -> State 281

  State 157:
    Kernel Items:
      sequence_expr: GREATER sequence_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 158:
    Kernel Items:
      import_clause: DOT.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 283

  State 159:
    Kernel Items:
      import_clause: IDENTIFIER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 284

  State 160:
    Kernel Items:
      import_statement: IMPORT LPAREN.import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 161
      IDENTIFIER -> State 159
      UNDERSCORE -> State 162
      DOT -> State 158
      import_clause -> State 285
      proper_import_clauses -> State 287
      import_clauses -> State 286

  State 161:
    Kernel Items:
      import_clause: STRING_LITERAL., $
    Reduce:
      $ -> [import_clause]
    Goto:
      (nil)

  State 162:
    Kernel Items:
      import_clause: UNDERSCORE.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 288

  State 163:
    Kernel Items:
      import_statement: IMPORT import_clause., $
    Reduce:
      $ -> [import_statement]
    Goto:
      (nil)

  State 164:
    Kernel Items:
      slice_type_expr: LBRACKET type_expr.RBRACKET
      array_type_expr: LBRACKET type_expr.COMMA INTEGER_LITERAL RBRACKET
      map_type_expr: LBRACKET type_expr.COLON type_expr RBRACKET
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 291
      COMMA -> State 290
      COLON -> State 289
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 165:
    Kernel Items:
      colon_expr: COLON., COLON
      colon_expr: COLON., COMMA
      colon_expr: COLON., RBRACKET
      colon_expr: COLON., RPAREN
      colon_expr: COLON.expr
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
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 292
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 166:
    Kernel Items:
      argument: ELLIPSIS., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 167:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expr
      named_expr: IDENTIFIER., *
    Reduce:
      * -> [named_expr]
    Goto:
      ASSIGN -> State 293

  State 168:
    Kernel Items:
      proper_arguments: argument., *
    Reduce:
      * -> [proper_arguments]
    Goto:
      (nil)

  State 169:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 294

  State 170:
    Kernel Items:
      argument: colon_expr., *
      colon_expr: colon_expr.COLON
      colon_expr: colon_expr.COLON expr
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 295

  State 171:
    Kernel Items:
      argument: expr., *
      argument: expr.ELLIPSIS
      colon_expr: expr.COLON
      colon_expr: expr.COLON expr
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 296
      ELLIPSIS -> State 297

  State 172:
    Kernel Items:
      proper_arguments: proper_arguments.COMMA argument
      arguments: proper_arguments., RPAREN
      arguments: proper_arguments.COMMA
    Reduce:
      RPAREN -> [arguments]
    Goto:
      COMMA -> State 298

  State 173:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN.explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 237
      UNDERSCORE -> State 238
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 236
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      unsafe_statement -> State 246
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 300
      implicit_struct_type_expr -> State 125
      proper_explicit_field_defs -> State 301
      explicit_field_defs -> State 299
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 174:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 302

  State 175:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 176:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., $
    Reduce:
      $ -> [unary_op_assign]
    Goto:
      (nil)

  State 177:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 178:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 179:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 180:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 181:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 182:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 183:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 184:
    Kernel Items:
      generic_type_arguments: DOLLAR_LBRACKET.type_arguments RBRACKET
    Reduce:
      RBRACKET -> [type_arguments]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 305
      binary_type_expr -> State 121
      proper_type_arguments -> State 303
      type_arguments -> State 304
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 185:
    Kernel Items:
      access_expr: accessible_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 306

  State 186:
    Kernel Items:
      index_expr: accessible_expr LBRACKET.argument RBRACKET
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 167
      UNDERSCORE -> State 53
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
      COLON -> State 165
      ELLIPSIS -> State 166
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 171
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      argument -> State 307
      colon_expr -> State 170
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 187:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 188:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 189:
    Kernel Items:
      postfix_unary_expr: accessible_expr QUESTION., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 190:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 191:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., $
    Reduce:
      $ -> [unary_op_assign]
    Goto:
      (nil)

  State 192:
    Kernel Items:
      binary_op_assign_statement: accessible_expr binary_op_assign.expr
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 308
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 193:
    Kernel Items:
      call_expr: accessible_expr generic_type_arguments.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 309

  State 194:
    Kernel Items:
      unary_op_assign_statement: accessible_expr unary_op_assign., $
    Reduce:
      $ -> [unary_op_assign_statement]
    Goto:
      (nil)

  State 195:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 196:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 197:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 198:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 199:
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
      UNDERSCORE -> State 53
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
      optional_label_decl -> State 153
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 310
      binary_mul_expr -> State 67
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 200:
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
      UNDERSCORE -> State 53
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
      optional_label_decl -> State 153
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 311
      binary_cmp_expr -> State 66
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 201:
    Kernel Items:
      assign_statement: assign_pattern ASSIGN.expr
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 312
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 202:
    Kernel Items:
      call_expr: accessible_expr.generic_type_arguments LPAREN arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
    Reduce:
      LPAREN -> [generic_type_arguments]
    Goto:
      LBRACKET -> State 186
      DOT -> State 185
      DOLLAR_LBRACKET -> State 184
      generic_type_arguments -> State 193

  State 203:
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

  State 204:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 205:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 206:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 207:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 208:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 209:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 210:
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
      UNDERSCORE -> State 53
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
      optional_label_decl -> State 153
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 313
      binary_add_expr -> State 64
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 211:
    Kernel Items:
      exprs: exprs COMMA.expr
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 314
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 212:
    Kernel Items:
      initialize_expr: initializable_type_expr LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [arguments]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 167
      UNDERSCORE -> State 53
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
      COLON -> State 165
      ELLIPSIS -> State 166
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 171
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      proper_arguments -> State 172
      arguments -> State 315
      argument -> State 168
      colon_expr -> State 170
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 213:
    Kernel Items:
      jump_statement: jump_type JUMP_LABEL., $
      jump_statement: jump_type JUMP_LABEL., NEWLINES
      jump_statement: jump_type JUMP_LABEL., RBRACE
      jump_statement: jump_type JUMP_LABEL., SEMICOLON
      jump_statement: jump_type JUMP_LABEL.exprs
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
      UNDERSCORE -> State 53
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
      exprs -> State 316
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 76
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 214:
    Kernel Items:
      exprs: exprs.COMMA expr
      jump_statement: jump_type exprs., $
    Reduce:
      $ -> [jump_statement]
    Goto:
      COMMA -> State 211

  State 215:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 216:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 217:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 218:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 219:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 220:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 221:
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
      UNDERSCORE -> State 53
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
      optional_label_decl -> State 153
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 317
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 222:
    Kernel Items:
      loop_expr: DO.statement_block
      loop_expr: DO.statement_block FOR sequence_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 318

  State 223:
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
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      assign_pattern -> State 319
      optional_label_decl -> State 153
      sequence_expr -> State 321
      for_assignment -> State 320
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 224:
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
      UNDERSCORE -> State 53
      TRUE -> State 52
      FALSE -> State 34
      CASE -> State 322
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      optional_label_decl -> State 153
      sequence_expr -> State 324
      condition -> State 323
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 225:
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
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      optional_label_decl -> State 153
      sequence_expr -> State 325
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 226:
    Kernel Items:
      expr: optional_label_decl if_expr., $
    Reduce:
      $ -> [expr]
    Goto:
      (nil)

  State 227:
    Kernel Items:
      expr: optional_label_decl loop_expr., $
    Reduce:
      $ -> [expr]
    Goto:
      (nil)

  State 228:
    Kernel Items:
      block_expr: optional_label_decl statement_block., *
    Reduce:
      * -> [block_expr]
    Goto:
      (nil)

  State 229:
    Kernel Items:
      expr: optional_label_decl switch_expr., $
    Reduce:
      $ -> [expr]
    Goto:
      (nil)

  State 230:
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
      UNDERSCORE -> State 53
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
      optional_label_decl -> State 153
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 326
      binary_and_expr -> State 65
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 231:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefixable_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 232:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 237
      UNDERSCORE -> State 238
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 236
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      unsafe_statement -> State 246
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 328
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      proper_explicit_enum_value_defs -> State 329
      explicit_enum_value_defs -> State 327
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 233:
    Kernel Items:
      func_type_expr: FUNC LPAREN.parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 331
      UNDERSCORE -> State 332
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      ELLIPSIS -> State 330
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 337
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      proper_parameter_def -> State 336
      parameter_decl -> State 333
      proper_parameter_decls -> State 335
      parameter_decls -> State 334
      func_type_expr -> State 123

  State 234:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT.IDENTIFIER generic_type_arguments
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 338

  State 235:
    Kernel Items:
      named_type_expr: IDENTIFIER generic_type_arguments., *
    Reduce:
      * -> [named_type_expr]
    Goto:
      (nil)

  State 236:
    Kernel Items:
      func_type_expr: FUNC.LPAREN parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 339
      LPAREN -> State 233

  State 237:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_type_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_type_arguments
      field_def: IDENTIFIER.type_expr optional_default
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 340
      QUESTION -> State 116
      EXCLAIM -> State 111
      DOLLAR_LBRACKET -> State 184
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 341
      binary_type_expr -> State 121
      generic_type_arguments -> State 235
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 238:
    Kernel Items:
      inferred_type_expr: UNDERSCORE., *
      field_def: UNDERSCORE.type_expr
    Reduce:
      * -> [inferred_type_expr]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 342
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 239:
    Kernel Items:
      proper_implicit_field_defs: field_def., *
      proper_implicit_enum_value_defs: field_def.OR field_def
    Reduce:
      * -> [proper_implicit_field_defs]
    Goto:
      OR -> State 343

  State 240:
    Kernel Items:
      implicit_enum_type_expr: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 344

  State 241:
    Kernel Items:
      implicit_struct_type_expr: LPAREN implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 345

  State 242:
    Kernel Items:
      field_def: method_signature., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 243:
    Kernel Items:
      proper_implicit_enum_value_defs: proper_implicit_enum_value_defs.OR field_def
      implicit_enum_value_defs: proper_implicit_enum_value_defs., RPAREN
      implicit_enum_value_defs: proper_implicit_enum_value_defs.NEWLINES
    Reduce:
      RPAREN -> [implicit_enum_value_defs]
    Goto:
      NEWLINES -> State 346
      OR -> State 347

  State 244:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs.COMMA field_def
      implicit_field_defs: proper_implicit_field_defs., RPAREN
      implicit_field_defs: proper_implicit_field_defs.COMMA
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      COMMA -> State 348

  State 245:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      field_def: type_expr.optional_default
    Reduce:
      * -> [optional_default]
    Goto:
      ASSIGN -> State 349
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252
      optional_default -> State 350

  State 246:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 247:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN.explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 237
      UNDERSCORE -> State 238
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 236
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      unsafe_statement -> State 246
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 300
      implicit_struct_type_expr -> State 125
      proper_explicit_field_defs -> State 301
      explicit_field_defs -> State 351
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 248:
    Kernel Items:
      prefix_unary_type_expr: prefix_unary_type_op returnable_type_expr., *
    Reduce:
      * -> [prefix_unary_type_expr]
    Goto:
      (nil)

  State 249:
    Kernel Items:
      binary_type_op: ADD., *
    Reduce:
      * -> [binary_type_op]
    Goto:
      (nil)

  State 250:
    Kernel Items:
      binary_type_op: MUL., *
    Reduce:
      * -> [binary_type_op]
    Goto:
      (nil)

  State 251:
    Kernel Items:
      binary_type_op: SUB., *
    Reduce:
      * -> [binary_type_op]
    Goto:
      (nil)

  State 252:
    Kernel Items:
      binary_type_expr: type_expr binary_type_op.returnable_type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 352
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 253:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN.expr
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 353
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 254:
    Kernel Items:
      generic_parameter_defs: DOLLAR_LBRACKET.generic_parameter_def_list RBRACKET
    Reduce:
      RBRACKET -> [generic_parameter_def_list]
    Goto:
      IDENTIFIER -> State 354
      generic_parameter_def -> State 355
      proper_generic_parameter_def_list -> State 357
      generic_parameter_def_list -> State 356

  State 255:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 358

  State 256:
    Kernel Items:
      proper_parameter_def: IDENTIFIER.type_expr
      proper_parameter_def: IDENTIFIER.ELLIPSIS type_expr
      proper_parameter_def: IDENTIFIER.ELLIPSIS
      parameter_def: IDENTIFIER., *
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      ELLIPSIS -> State 359
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 360
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 257:
    Kernel Items:
      proper_parameter_def: UNDERSCORE.type_expr
      proper_parameter_def: UNDERSCORE.ELLIPSIS
      proper_parameter_def: UNDERSCORE.ELLIPSIS type_expr
      parameter_def: UNDERSCORE., *
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      ELLIPSIS -> State 361
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 362
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 258:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def.RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 363

  State 259:
    Kernel Items:
      parameter_def: proper_parameter_def., *
    Reduce:
      * -> [parameter_def]
    Goto:
      (nil)

  State 260:
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
      UNDERSCORE -> State 53
      TRUE -> State 52
      FALSE -> State 34
      CASE -> State 29
      DEFAULT -> State 31
      RETURN -> State 47
      BREAK -> State 28
      CONTINUE -> State 30
      FALLTHROUGH -> State 33
      IMPORT -> State 39
      UNSAFE -> State 54
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
      simple_statement -> State 100
      statement -> State 364
      expr_or_improper_struct_statement -> State 77
      exprs -> State 78
      callback_op -> State 72
      callback_op_statement -> State 73
      unsafe_statement -> State 103
      jump_statement -> State 85
      jump_type -> State 86
      fallthrough_statement -> State 79
      assign_statement -> State 62
      unary_op_assign_statement -> State 102
      binary_op_assign_statement -> State 68
      import_statement -> State 81
      var_decl_pattern -> State 104
      var_or_let -> State 24
      assign_pattern -> State 61
      expr -> State 76
      optional_label_decl -> State 91
      sequence_expr -> State 99
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 56
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 261:
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
      UNDERSCORE -> State 53
      TRUE -> State 52
      FALSE -> State 34
      CASE -> State 29
      DEFAULT -> State 31
      RETURN -> State 47
      BREAK -> State 28
      CONTINUE -> State 30
      FALLTHROUGH -> State 33
      IMPORT -> State 39
      UNSAFE -> State 54
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
      simple_statement -> State 100
      statement -> State 365
      expr_or_improper_struct_statement -> State 77
      exprs -> State 78
      callback_op -> State 72
      callback_op_statement -> State 73
      unsafe_statement -> State 103
      jump_statement -> State 85
      jump_type -> State 86
      fallthrough_statement -> State 79
      assign_statement -> State 62
      unary_op_assign_statement -> State 102
      binary_op_assign_statement -> State 68
      import_statement -> State 81
      var_decl_pattern -> State 104
      var_or_let -> State 24
      assign_pattern -> State 61
      expr -> State 76
      optional_label_decl -> State 91
      sequence_expr -> State 99
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 56
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 262:
    Kernel Items:
      statement_block: LBRACE statements RBRACE., $
    Reduce:
      $ -> [statement_block]
    Goto:
      (nil)

  State 263:
    Kernel Items:
      type_def: TYPE IDENTIFIER ASSIGN.type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 366
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 264:
    Kernel Items:
      type_def: TYPE IDENTIFIER generic_parameter_defs.type_expr
      type_def: TYPE IDENTIFIER generic_parameter_defs.type_expr IMPLEMENTS type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 367
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 265:
    Kernel Items:
      proper_definitions: proper_definitions NEWLINES definition., *
    Reduce:
      * -> [proper_definitions]
    Goto:
      (nil)

  State 266:
    Kernel Items:
      definition: var_decl_pattern ASSIGN expr., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 267:
    Kernel Items:
      field_var_pattern: ELLIPSIS., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 268:
    Kernel Items:
      var_pattern: IDENTIFIER., *
      field_var_pattern: IDENTIFIER.ASSIGN var_pattern
    Reduce:
      * -> [var_pattern]
    Goto:
      ASSIGN -> State 368

  State 269:
    Kernel Items:
      field_var_patterns: field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 270:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns.RPAREN
      field_var_patterns: field_var_patterns.COMMA field_var_pattern
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 370
      COMMA -> State 369

  State 271:
    Kernel Items:
      field_var_pattern: var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 272:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern optional_type_expr., $
    Reduce:
      $ -> [var_decl_pattern]
    Goto:
      (nil)

  State 273:
    Kernel Items:
      optional_type_expr: type_expr., *
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      * -> [optional_type_expr]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 274:
    Kernel Items:
      case_pattern: DOT IDENTIFIER.implicit_struct_expr
      case_pattern: DOT IDENTIFIER., *
    Reduce:
      * -> [case_pattern]
    Goto:
      LPAREN -> State 43
      implicit_struct_expr -> State 371

  State 275:
    Kernel Items:
      case_pattern: VAR DOT.IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 372

  State 276:
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
      UNDERSCORE -> State 53
      TRUE -> State 52
      FALSE -> State 34
      RETURN -> State 47
      BREAK -> State 28
      CONTINUE -> State 30
      FALLTHROUGH -> State 33
      UNSAFE -> State 54
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
      simple_statement -> State 279
      optional_simple_statement -> State 373
      expr_or_improper_struct_statement -> State 77
      exprs -> State 78
      callback_op -> State 72
      callback_op_statement -> State 73
      unsafe_statement -> State 103
      jump_statement -> State 85
      jump_type -> State 86
      fallthrough_statement -> State 79
      assign_statement -> State 62
      unary_op_assign_statement -> State 102
      binary_op_assign_statement -> State 68
      var_decl_pattern -> State 104
      var_or_let -> State 24
      assign_pattern -> State 61
      expr -> State 76
      optional_label_decl -> State 91
      sequence_expr -> State 99
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 56
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 277:
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
      UNDERSCORE -> State 53
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 149
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      DOT -> State 148
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 104
      var_or_let -> State 24
      assign_pattern -> State 150
      case_pattern -> State 374
      optional_label_decl -> State 153
      sequence_expr -> State 154
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 278:
    Kernel Items:
      statement: DEFAULT COLON optional_simple_statement., $
    Reduce:
      $ -> [statement]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      optional_simple_statement: simple_statement., $
    Reduce:
      $ -> [optional_simple_statement]
    Goto:
      (nil)

  State 280:
    Kernel Items:
      proper_parameter_defs: parameter_def., *
    Reduce:
      * -> [proper_parameter_defs]
    Goto:
      (nil)

  State 281:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 375

  State 282:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs.COMMA parameter_def
      parameter_defs: proper_parameter_defs., RPAREN
      parameter_defs: proper_parameter_defs.COMMA
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      COMMA -> State 376

  State 283:
    Kernel Items:
      import_clause: DOT STRING_LITERAL., $
    Reduce:
      $ -> [import_clause]
    Goto:
      (nil)

  State 284:
    Kernel Items:
      import_clause: IDENTIFIER STRING_LITERAL., $
    Reduce:
      $ -> [import_clause]
    Goto:
      (nil)

  State 285:
    Kernel Items:
      proper_import_clauses: import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 286:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 377

  State 287:
    Kernel Items:
      proper_import_clauses: proper_import_clauses.NEWLINES import_clause
      proper_import_clauses: proper_import_clauses.COMMA import_clause
      import_clauses: proper_import_clauses., RPAREN
      import_clauses: proper_import_clauses.NEWLINES
      import_clauses: proper_import_clauses.COMMA
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      NEWLINES -> State 379
      COMMA -> State 378

  State 288:
    Kernel Items:
      import_clause: UNDERSCORE STRING_LITERAL., $
    Reduce:
      $ -> [import_clause]
    Goto:
      (nil)

  State 289:
    Kernel Items:
      map_type_expr: LBRACKET type_expr COLON.type_expr RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 380
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 290:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 381

  State 291:
    Kernel Items:
      slice_type_expr: LBRACKET type_expr RBRACKET., *
    Reduce:
      * -> [slice_type_expr]
    Goto:
      (nil)

  State 292:
    Kernel Items:
      colon_expr: COLON expr., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 293:
    Kernel Items:
      argument: IDENTIFIER ASSIGN.expr
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 38
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 382
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 294:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [implicit_struct_expr]
    Goto:
      (nil)

  State 295:
    Kernel Items:
      colon_expr: colon_expr COLON., COLON
      colon_expr: colon_expr COLON., COMMA
      colon_expr: colon_expr COLON., RBRACKET
      colon_expr: colon_expr COLON., RPAREN
      colon_expr: colon_expr COLON.expr
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
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 383
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 296:
    Kernel Items:
      colon_expr: expr COLON., COLON
      colon_expr: expr COLON., COMMA
      colon_expr: expr COLON., RBRACKET
      colon_expr: expr COLON., RPAREN
      colon_expr: expr COLON.expr
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
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 384
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 297:
    Kernel Items:
      argument: expr ELLIPSIS., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 298:
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
      IDENTIFIER -> State 167
      UNDERSCORE -> State 53
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
      COLON -> State 165
      ELLIPSIS -> State 166
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 171
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      argument -> State 385
      colon_expr -> State 170
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 299:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 386

  State 300:
    Kernel Items:
      proper_explicit_field_defs: field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 301:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs.NEWLINES field_def
      proper_explicit_field_defs: proper_explicit_field_defs.COMMA field_def
      explicit_field_defs: proper_explicit_field_defs., RPAREN
      explicit_field_defs: proper_explicit_field_defs.NEWLINES
      explicit_field_defs: proper_explicit_field_defs.COMMA
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      NEWLINES -> State 388
      COMMA -> State 387

  State 302:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 389

  State 303:
    Kernel Items:
      proper_type_arguments: proper_type_arguments.COMMA type_expr
      type_arguments: proper_type_arguments., RBRACKET
      type_arguments: proper_type_arguments.COMMA
    Reduce:
      RBRACKET -> [type_arguments]
    Goto:
      COMMA -> State 390

  State 304:
    Kernel Items:
      generic_type_arguments: DOLLAR_LBRACKET type_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 391

  State 305:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_type_arguments: type_expr., *
    Reduce:
      * -> [proper_type_arguments]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 306:
    Kernel Items:
      access_expr: accessible_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 307:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 392

  State 308:
    Kernel Items:
      binary_op_assign_statement: accessible_expr binary_op_assign expr., $
    Reduce:
      $ -> [binary_op_assign_statement]
    Goto:
      (nil)

  State 309:
    Kernel Items:
      call_expr: accessible_expr generic_type_arguments LPAREN.arguments RPAREN
    Reduce:
      * -> [optional_label_decl]
      RPAREN -> [arguments]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 167
      UNDERSCORE -> State 53
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
      COLON -> State 165
      ELLIPSIS -> State 166
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 104
      var_or_let -> State 24
      expr -> State 171
      optional_label_decl -> State 91
      sequence_expr -> State 106
      call_expr -> State 71
      proper_arguments -> State 172
      arguments -> State 393
      argument -> State 168
      colon_expr -> State 170
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 310:
    Kernel Items:
      binary_mul_expr: mul_expr.mul_op prefixable_expr
      binary_add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [binary_add_expr]
    Goto:
      MUL -> State 220
      DIV -> State 218
      MOD -> State 219
      BIT_AND -> State 215
      BIT_LSHIFT -> State 216
      BIT_RSHIFT -> State 217
      mul_op -> State 221

  State 311:
    Kernel Items:
      binary_cmp_expr: cmp_expr.cmp_op add_expr
      binary_and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [binary_and_expr]
    Goto:
      EQUAL -> State 204
      NOT_EQUAL -> State 209
      LESS -> State 207
      LESS_OR_EQUAL -> State 208
      GREATER -> State 205
      GREATER_OR_EQUAL -> State 206
      cmp_op -> State 210

  State 312:
    Kernel Items:
      assign_statement: assign_pattern ASSIGN expr., $
    Reduce:
      $ -> [assign_statement]
    Goto:
      (nil)

  State 313:
    Kernel Items:
      binary_add_expr: add_expr.add_op mul_expr
      binary_cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [binary_cmp_expr]
    Goto:
      ADD -> State 195
      SUB -> State 198
      BIT_XOR -> State 197
      BIT_OR -> State 196
      add_op -> State 199

  State 314:
    Kernel Items:
      exprs: exprs COMMA expr., *
    Reduce:
      * -> [exprs]
    Goto:
      (nil)

  State 315:
    Kernel Items:
      initialize_expr: initializable_type_expr LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 394

  State 316:
    Kernel Items:
      exprs: exprs.COMMA expr
      jump_statement: jump_type JUMP_LABEL exprs., $
    Reduce:
      $ -> [jump_statement]
    Goto:
      COMMA -> State 211

  State 317:
    Kernel Items:
      binary_mul_expr: mul_expr mul_op prefixable_expr., *
    Reduce:
      * -> [binary_mul_expr]
    Goto:
      (nil)

  State 318:
    Kernel Items:
      loop_expr: DO statement_block., *
      loop_expr: DO statement_block.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 395

  State 319:
    Kernel Items:
      loop_expr: FOR assign_pattern.IN sequence_expr DO statement_block
      for_assignment: assign_pattern.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      IN -> State 397
      ASSIGN -> State 396

  State 320:
    Kernel Items:
      loop_expr: FOR for_assignment.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 398

  State 321:
    Kernel Items:
      assign_pattern: sequence_expr., *
      loop_expr: FOR sequence_expr.DO statement_block
      for_assignment: sequence_expr., SEMICOLON
    Reduce:
      * -> [assign_pattern]
      SEMICOLON -> [for_assignment]
    Goto:
      DO -> State 399

  State 322:
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
      UNDERSCORE -> State 53
      TRUE -> State 52
      FALSE -> State 34
      STRUCT -> State 50
      FUNC -> State 36
      VAR -> State 149
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      DOT -> State 148
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      case_patterns -> State 400
      var_decl_pattern -> State 104
      var_or_let -> State 24
      assign_pattern -> State 150
      case_pattern -> State 151
      optional_label_decl -> State 153
      sequence_expr -> State 154
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 323:
    Kernel Items:
      if_expr: IF condition.statement_block
      if_expr: IF condition.statement_block ELSE statement_block
      if_expr: IF condition.statement_block ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 401

  State 324:
    Kernel Items:
      condition: sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 325:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 402

  State 326:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      binary_or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [binary_or_expr]
    Goto:
      AND -> State 200

  State 327:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 403

  State 328:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def.OR field_def
      proper_explicit_enum_value_defs: field_def.NEWLINES field_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 404
      OR -> State 405

  State 329:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs.OR field_def
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs.NEWLINES field_def
      explicit_enum_value_defs: proper_explicit_enum_value_defs., RPAREN
      explicit_enum_value_defs: proper_explicit_enum_value_defs.NEWLINES
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      NEWLINES -> State 406
      OR -> State 407

  State 330:
    Kernel Items:
      parameter_decl: ELLIPSIS., *
      parameter_decl: ELLIPSIS.type_expr
    Reduce:
      * -> [parameter_decl]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 408
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 331:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_type_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_type_arguments
      proper_parameter_def: IDENTIFIER.type_expr
      proper_parameter_def: IDENTIFIER.ELLIPSIS type_expr
      proper_parameter_def: IDENTIFIER.ELLIPSIS
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 340
      QUESTION -> State 116
      EXCLAIM -> State 111
      DOLLAR_LBRACKET -> State 184
      ELLIPSIS -> State 359
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 360
      binary_type_expr -> State 121
      generic_type_arguments -> State 235
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 332:
    Kernel Items:
      inferred_type_expr: UNDERSCORE., *
      proper_parameter_def: UNDERSCORE.type_expr
      proper_parameter_def: UNDERSCORE.ELLIPSIS
      proper_parameter_def: UNDERSCORE.ELLIPSIS type_expr
    Reduce:
      * -> [inferred_type_expr]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      ELLIPSIS -> State 361
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 362
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 333:
    Kernel Items:
      proper_parameter_decls: parameter_decl., *
    Reduce:
      * -> [proper_parameter_decls]
    Goto:
      (nil)

  State 334:
    Kernel Items:
      func_type_expr: FUNC LPAREN parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 409

  State 335:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls.COMMA parameter_decl
      parameter_decls: proper_parameter_decls., RPAREN
      parameter_decls: proper_parameter_decls.COMMA
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      COMMA -> State 410

  State 336:
    Kernel Items:
      parameter_decl: proper_parameter_def., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 337:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_decl: type_expr., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 338:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT IDENTIFIER.generic_type_arguments
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      DOLLAR_LBRACKET -> State 184
      generic_type_arguments -> State 411

  State 339:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 412

  State 340:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT.IDENTIFIER generic_type_arguments
      inferred_type_expr: DOT., *
    Reduce:
      * -> [inferred_type_expr]
    Goto:
      IDENTIFIER -> State 338

  State 341:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      field_def: IDENTIFIER type_expr.optional_default
    Reduce:
      * -> [optional_default]
    Goto:
      ASSIGN -> State 349
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252
      optional_default -> State 413

  State 342:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      field_def: UNDERSCORE type_expr., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 343:
    Kernel Items:
      proper_implicit_enum_value_defs: field_def OR.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 237
      UNDERSCORE -> State 238
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 236
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      unsafe_statement -> State 246
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 414
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 344:
    Kernel Items:
      implicit_enum_type_expr: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_type_expr]
    Goto:
      (nil)

  State 345:
    Kernel Items:
      implicit_struct_type_expr: LPAREN implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_type_expr]
    Goto:
      (nil)

  State 346:
    Kernel Items:
      implicit_enum_value_defs: proper_implicit_enum_value_defs NEWLINES., RPAREN
    Reduce:
      RPAREN -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 347:
    Kernel Items:
      proper_implicit_enum_value_defs: proper_implicit_enum_value_defs OR.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 237
      UNDERSCORE -> State 238
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 236
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      unsafe_statement -> State 246
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 415
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 348:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs COMMA.field_def
      implicit_field_defs: proper_implicit_field_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      IDENTIFIER -> State 237
      UNDERSCORE -> State 238
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 236
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      unsafe_statement -> State 246
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 416
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 349:
    Kernel Items:
      optional_default: ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 417

  State 350:
    Kernel Items:
      field_def: type_expr optional_default., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 351:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 418

  State 352:
    Kernel Items:
      binary_type_expr: type_expr binary_type_op returnable_type_expr., *
    Reduce:
      * -> [binary_type_expr]
    Goto:
      (nil)

  State 353:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN expr., *
    Reduce:
      * -> [named_func_def]
    Goto:
      (nil)

  State 354:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.type_expr
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 419
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 355:
    Kernel Items:
      proper_generic_parameter_def_list: generic_parameter_def., *
    Reduce:
      * -> [proper_generic_parameter_def_list]
    Goto:
      (nil)

  State 356:
    Kernel Items:
      generic_parameter_defs: DOLLAR_LBRACKET generic_parameter_def_list.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 420

  State 357:
    Kernel Items:
      proper_generic_parameter_def_list: proper_generic_parameter_def_list.COMMA generic_parameter_def
      generic_parameter_def_list: proper_generic_parameter_def_list., RBRACKET
      generic_parameter_def_list: proper_generic_parameter_def_list.COMMA
    Reduce:
      RBRACKET -> [generic_parameter_def_list]
    Goto:
      COMMA -> State 421

  State 358:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 256
      UNDERSCORE -> State 257
      proper_parameter_def -> State 259
      parameter_def -> State 280
      proper_parameter_defs -> State 282
      parameter_defs -> State 422

  State 359:
    Kernel Items:
      proper_parameter_def: IDENTIFIER ELLIPSIS.type_expr
      proper_parameter_def: IDENTIFIER ELLIPSIS., *
    Reduce:
      * -> [proper_parameter_def]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 423
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 360:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_parameter_def: IDENTIFIER type_expr., *
    Reduce:
      * -> [proper_parameter_def]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 361:
    Kernel Items:
      proper_parameter_def: UNDERSCORE ELLIPSIS., *
      proper_parameter_def: UNDERSCORE ELLIPSIS.type_expr
    Reduce:
      * -> [proper_parameter_def]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 424
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 362:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_parameter_def: UNDERSCORE type_expr., *
    Reduce:
      * -> [proper_parameter_def]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 363:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN.IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 425

  State 364:
    Kernel Items:
      proper_statements: proper_statements NEWLINES statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 365:
    Kernel Items:
      proper_statements: proper_statements SEMICOLON statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 366:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER ASSIGN type_expr., *
    Reduce:
      * -> [type_def]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 367:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER generic_parameter_defs type_expr., *
      type_def: TYPE IDENTIFIER generic_parameter_defs type_expr.IMPLEMENTS type_expr
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 426
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 368:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN.var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 143
      UNDERSCORE -> State 145
      LPAREN -> State 144
      var_pattern -> State 427
      tuple_pattern -> State 146

  State 369:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA.field_var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 268
      UNDERSCORE -> State 145
      LPAREN -> State 144
      ELLIPSIS -> State 267
      var_pattern -> State 271
      tuple_pattern -> State 146
      field_var_pattern -> State 428

  State 370:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns RPAREN., *
    Reduce:
      * -> [tuple_pattern]
    Goto:
      (nil)

  State 371:
    Kernel Items:
      case_pattern: DOT IDENTIFIER implicit_struct_expr., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 372:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER.tuple_pattern
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 144
      tuple_pattern -> State 429

  State 373:
    Kernel Items:
      statement: CASE case_patterns COLON optional_simple_statement., $
    Reduce:
      $ -> [statement]
    Goto:
      (nil)

  State 374:
    Kernel Items:
      case_patterns: case_patterns COMMA case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 375:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 431
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      return_type -> State 430
      func_type_expr -> State 123

  State 376:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs COMMA.parameter_def
      parameter_defs: proper_parameter_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 256
      UNDERSCORE -> State 257
      proper_parameter_def -> State 259
      parameter_def -> State 432

  State 377:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses RPAREN., $
    Reduce:
      $ -> [import_statement]
    Goto:
      (nil)

  State 378:
    Kernel Items:
      proper_import_clauses: proper_import_clauses COMMA.import_clause
      import_clauses: proper_import_clauses COMMA., RPAREN
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      STRING_LITERAL -> State 161
      IDENTIFIER -> State 159
      UNDERSCORE -> State 162
      DOT -> State 158
      import_clause -> State 433

  State 379:
    Kernel Items:
      proper_import_clauses: proper_import_clauses NEWLINES.import_clause
      import_clauses: proper_import_clauses NEWLINES., RPAREN
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      STRING_LITERAL -> State 161
      IDENTIFIER -> State 159
      UNDERSCORE -> State 162
      DOT -> State 158
      import_clause -> State 434

  State 380:
    Kernel Items:
      map_type_expr: LBRACKET type_expr COLON type_expr.RBRACKET
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 435
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 381:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 436

  State 382:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expr., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 383:
    Kernel Items:
      colon_expr: colon_expr COLON expr., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 384:
    Kernel Items:
      colon_expr: expr COLON expr., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 385:
    Kernel Items:
      proper_arguments: proper_arguments COMMA argument., *
    Reduce:
      * -> [proper_arguments]
    Goto:
      (nil)

  State 386:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_type_expr]
    Goto:
      (nil)

  State 387:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs COMMA.field_def
      explicit_field_defs: proper_explicit_field_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 237
      UNDERSCORE -> State 238
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 236
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      unsafe_statement -> State 246
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 437
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 388:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs NEWLINES.field_def
      explicit_field_defs: proper_explicit_field_defs NEWLINES., RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 237
      UNDERSCORE -> State 238
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 236
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      unsafe_statement -> State 246
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 438
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 389:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 439

  State 390:
    Kernel Items:
      proper_type_arguments: proper_type_arguments COMMA.type_expr
      type_arguments: proper_type_arguments COMMA., RBRACKET
    Reduce:
      RBRACKET -> [type_arguments]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 440
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 391:
    Kernel Items:
      generic_type_arguments: DOLLAR_LBRACKET type_arguments RBRACKET., *
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      (nil)

  State 392:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [index_expr]
    Goto:
      (nil)

  State 393:
    Kernel Items:
      call_expr: accessible_expr generic_type_arguments LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 441

  State 394:
    Kernel Items:
      initialize_expr: initializable_type_expr LPAREN arguments RPAREN., *
    Reduce:
      * -> [initialize_expr]
    Goto:
      (nil)

  State 395:
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
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      optional_label_decl -> State 153
      sequence_expr -> State 442
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 396:
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
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      optional_label_decl -> State 153
      sequence_expr -> State 443
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 397:
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
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      optional_label_decl -> State 153
      sequence_expr -> State 444
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 398:
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
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      optional_label_decl -> State 153
      sequence_expr -> State 446
      optional_sequence_expr -> State 445
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 399:
    Kernel Items:
      loop_expr: FOR sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 447

  State 400:
    Kernel Items:
      case_patterns: case_patterns.COMMA case_pattern
      condition: CASE case_patterns.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      COMMA -> State 277
      ASSIGN -> State 448

  State 401:
    Kernel Items:
      if_expr: IF condition statement_block., *
      if_expr: IF condition statement_block.ELSE statement_block
      if_expr: IF condition statement_block.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 449

  State 402:
    Kernel Items:
      switch_expr: SWITCH sequence_expr statement_block., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 403:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_type_expr]
    Goto:
      (nil)

  State 404:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 237
      UNDERSCORE -> State 238
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 236
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      unsafe_statement -> State 246
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 450
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 405:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def OR.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 237
      UNDERSCORE -> State 238
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 236
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      unsafe_statement -> State 246
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 451
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 406:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs NEWLINES.field_def
      explicit_enum_value_defs: proper_explicit_enum_value_defs NEWLINES., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      IDENTIFIER -> State 237
      UNDERSCORE -> State 238
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 236
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      unsafe_statement -> State 246
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 452
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 407:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs OR.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 237
      UNDERSCORE -> State 238
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 236
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      unsafe_statement -> State 246
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 453
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 408:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_decl: ELLIPSIS type_expr., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 409:
    Kernel Items:
      func_type_expr: FUNC LPAREN parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 431
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      return_type -> State 454
      func_type_expr -> State 123

  State 410:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls COMMA.parameter_decl
      parameter_decls: proper_parameter_decls COMMA., RPAREN
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 331
      UNDERSCORE -> State 332
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      ELLIPSIS -> State 330
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 337
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      proper_parameter_def -> State 336
      parameter_decl -> State 455
      func_type_expr -> State 123

  State 411:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT IDENTIFIER generic_type_arguments., *
    Reduce:
      * -> [named_type_expr]
    Goto:
      (nil)

  State 412:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 331
      UNDERSCORE -> State 332
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      ELLIPSIS -> State 330
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 337
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      proper_parameter_def -> State 336
      parameter_decl -> State 333
      proper_parameter_decls -> State 335
      parameter_decls -> State 456
      func_type_expr -> State 123

  State 413:
    Kernel Items:
      field_def: IDENTIFIER type_expr optional_default., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 414:
    Kernel Items:
      proper_implicit_enum_value_defs: field_def OR field_def., *
    Reduce:
      * -> [proper_implicit_enum_value_defs]
    Goto:
      (nil)

  State 415:
    Kernel Items:
      proper_implicit_enum_value_defs: proper_implicit_enum_value_defs OR field_def., *
    Reduce:
      * -> [proper_implicit_enum_value_defs]
    Goto:
      (nil)

  State 416:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [proper_implicit_field_defs]
    Goto:
      (nil)

  State 417:
    Kernel Items:
      optional_default: ASSIGN DEFAULT., *
    Reduce:
      * -> [optional_default]
    Goto:
      (nil)

  State 418:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN explicit_field_defs RPAREN., *
    Reduce:
      * -> [trait_type_expr]
    Goto:
      (nil)

  State 419:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      generic_parameter_def: IDENTIFIER type_expr., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 420:
    Kernel Items:
      generic_parameter_defs: DOLLAR_LBRACKET generic_parameter_def_list RBRACKET., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 421:
    Kernel Items:
      proper_generic_parameter_def_list: proper_generic_parameter_def_list COMMA.generic_parameter_def
      generic_parameter_def_list: proper_generic_parameter_def_list COMMA., RBRACKET
    Reduce:
      RBRACKET -> [generic_parameter_def_list]
    Goto:
      IDENTIFIER -> State 354
      generic_parameter_def -> State 457

  State 422:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 458

  State 423:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_parameter_def: IDENTIFIER ELLIPSIS type_expr., *
    Reduce:
      * -> [proper_parameter_def]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 424:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_parameter_def: UNDERSCORE ELLIPSIS type_expr., *
    Reduce:
      * -> [proper_parameter_def]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 425:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 459

  State 426:
    Kernel Items:
      type_def: TYPE IDENTIFIER generic_parameter_defs type_expr IMPLEMENTS.type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      type_expr -> State 460
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 427:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 428:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 429:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER tuple_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 430:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 461

  State 431:
    Kernel Items:
      return_type: returnable_type_expr., *
    Reduce:
      * -> [return_type]
    Goto:
      (nil)

  State 432:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [proper_parameter_defs]
    Goto:
      (nil)

  State 433:
    Kernel Items:
      proper_import_clauses: proper_import_clauses COMMA import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 434:
    Kernel Items:
      proper_import_clauses: proper_import_clauses NEWLINES import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 435:
    Kernel Items:
      map_type_expr: LBRACKET type_expr COLON type_expr RBRACKET., *
    Reduce:
      * -> [map_type_expr]
    Goto:
      (nil)

  State 436:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA INTEGER_LITERAL RBRACKET., *
    Reduce:
      * -> [array_type_expr]
    Goto:
      (nil)

  State 437:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 438:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 439:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., $
    Reduce:
      $ -> [unsafe_statement]
    Goto:
      (nil)

  State 440:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_type_arguments: proper_type_arguments COMMA type_expr., *
    Reduce:
      * -> [proper_type_arguments]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 441:
    Kernel Items:
      call_expr: accessible_expr generic_type_arguments LPAREN arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 442:
    Kernel Items:
      loop_expr: DO statement_block FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 443:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN sequence_expr., SEMICOLON
    Reduce:
      SEMICOLON -> [for_assignment]
    Goto:
      (nil)

  State 444:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 462

  State 445:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 463

  State 446:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 447:
    Kernel Items:
      loop_expr: FOR sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 448:
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
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      optional_label_decl -> State 153
      sequence_expr -> State 464
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 449:
    Kernel Items:
      if_expr: IF condition statement_block ELSE.statement_block
      if_expr: IF condition statement_block ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 224
      LBRACE -> State 11
      statement_block -> State 466
      if_expr -> State 465

  State 450:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def NEWLINES field_def., *
    Reduce:
      * -> [proper_explicit_enum_value_defs]
    Goto:
      (nil)

  State 451:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def OR field_def., *
    Reduce:
      * -> [proper_explicit_enum_value_defs]
    Goto:
      (nil)

  State 452:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs NEWLINES field_def., *
    Reduce:
      * -> [proper_explicit_enum_value_defs]
    Goto:
      (nil)

  State 453:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs OR field_def., *
    Reduce:
      * -> [proper_explicit_enum_value_defs]
    Goto:
      (nil)

  State 454:
    Kernel Items:
      func_type_expr: FUNC LPAREN parameter_decls RPAREN return_type., *
    Reduce:
      * -> [func_type_expr]
    Goto:
      (nil)

  State 455:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [proper_parameter_decls]
    Goto:
      (nil)

  State 456:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 467

  State 457:
    Kernel Items:
      proper_generic_parameter_def_list: proper_generic_parameter_def_list COMMA generic_parameter_def., *
    Reduce:
      * -> [proper_generic_parameter_def_list]
    Goto:
      (nil)

  State 458:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs LPAREN parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 431
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      return_type -> State 468
      func_type_expr -> State 123

  State 459:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 256
      UNDERSCORE -> State 257
      proper_parameter_def -> State 259
      parameter_def -> State 280
      proper_parameter_defs -> State 282
      parameter_defs -> State 469

  State 460:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER generic_parameter_defs type_expr IMPLEMENTS type_expr., *
    Reduce:
      * -> [type_def]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 461:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 462:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 470

  State 463:
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
      UNDERSCORE -> State 53
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
      var_decl_pattern -> State 104
      var_or_let -> State 24
      optional_label_decl -> State 153
      sequence_expr -> State 446
      optional_sequence_expr -> State 471
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 93
      literal_expr -> State 87
      named_expr -> State 90
      block_expr -> State 70
      initialize_expr -> State 84
      implicit_struct_expr -> State 80
      accessible_expr -> State 105
      access_expr -> State 55
      index_expr -> State 82
      postfixable_expr -> State 95
      postfix_unary_expr -> State 94
      prefixable_expr -> State 98
      prefix_unary_expr -> State 96
      prefix_unary_op -> State 97
      mul_expr -> State 89
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 92
      binary_or_expr -> State 69
      initializable_type_expr -> State 83
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 464:
    Kernel Items:
      condition: CASE case_patterns ASSIGN sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 465:
    Kernel Items:
      if_expr: IF condition statement_block ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 466:
    Kernel Items:
      if_expr: IF condition statement_block ELSE statement_block., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 467:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 431
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      return_type -> State 472
      func_type_expr -> State 123

  State 468:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 473

  State 469:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 474

  State 470:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 471:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 475

  State 472:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 473:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [named_func_def]
    Goto:
      (nil)

  State 474:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 113
      UNDERSCORE -> State 119
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 431
      prefix_unary_type_expr -> State 130
      prefix_unary_type_op -> State 131
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      return_type -> State 476
      func_type_expr -> State 123

  State 475:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 477

  State 476:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 478

  State 477:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 478:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [named_func_def]
    Goto:
      (nil)

Number of states: 478
Number of shift actions: 4523
Number of reduce actions: 429
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 5305
Number of unoptimized shift actions: 46402
Number of unoptimized reduce actions: 34393
*/
