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

type GenericTypeArgumentsReducer interface {
	// 432:2: generic_type_arguments -> binding: ...
	BindingToGenericTypeArguments(DollarLbracket_ TokenValue, TypeArguments_ *TypeArgumentList, Rbracket_ TokenValue) (*TypeArgumentList, error)

	// 433:2: generic_type_arguments -> nil: ...
	NilToGenericTypeArguments() (*TypeArgumentList, error)
}

type ProperTypeArgumentsReducer interface {
	// 436:2: proper_type_arguments -> add: ...
	AddToProperTypeArguments(ProperTypeArguments_ *TypeArgumentList, Comma_ TokenValue, TypeExpr_ TypeExpression) (*TypeArgumentList, error)

	// 437:2: proper_type_arguments -> type_expr: ...
	TypeExprToProperTypeArguments(TypeExpr_ TypeExpression) (*TypeArgumentList, error)
}

type TypeArgumentsReducer interface {

	// 441:2: type_arguments -> improper: ...
	ImproperToTypeArguments(ProperTypeArguments_ *TypeArgumentList, Comma_ TokenValue) (*TypeArgumentList, error)

	// 442:2: type_arguments -> nil: ...
	NilToTypeArguments() (*TypeArgumentList, error)
}

type ProperArgumentsReducer interface {
	// 445:2: proper_arguments -> add: ...
	AddToProperArguments(ProperArguments_ *ArgumentList, Comma_ TokenValue, Argument_ *Argument) (*ArgumentList, error)

	// 446:2: proper_arguments -> argument: ...
	ArgumentToProperArguments(Argument_ *Argument) (*ArgumentList, error)
}

type ArgumentsReducer interface {

	// 450:2: arguments -> improper: ...
	ImproperToArguments(ProperArguments_ *ArgumentList, Comma_ TokenValue) (*ArgumentList, error)

	// 451:2: arguments -> nil: ...
	NilToArguments() (*ArgumentList, error)
}

type ArgumentReducer interface {
	// 454:2: argument -> positional: ...
	PositionalToArgument(Expr_ Expression) (*Argument, error)

	// 455:2: argument -> colon_expr: ...
	ColonExprToArgument(ColonExpr_ *ColonExpr) (*Argument, error)

	// 456:2: argument -> named_assignment: ...
	NamedAssignmentToArgument(Identifier_ TokenValue, Assign_ TokenValue, Expr_ Expression) (*Argument, error)

	// 460:2: argument -> vararg_assignment: ...
	VarargAssignmentToArgument(Expr_ Expression, Ellipsis_ TokenValue) (*Argument, error)

	// 463:2: argument -> skip_pattern: ...
	SkipPatternToArgument(Ellipsis_ TokenValue) (*Argument, error)
}

type ColonExprReducer interface {
	// 467:2: colon_expr -> unit_unit_pair: ...
	UnitUnitPairToColonExpr(Colon_ TokenValue) (*ColonExpr, error)

	// 468:2: colon_expr -> expr_unit_pair: ...
	ExprUnitPairToColonExpr(Expr_ Expression, Colon_ TokenValue) (*ColonExpr, error)

	// 469:2: colon_expr -> unit_expr_pair: ...
	UnitExprPairToColonExpr(Colon_ TokenValue, Expr_ Expression) (*ColonExpr, error)

	// 470:2: colon_expr -> expr_expr_pair: ...
	ExprExprPairToColonExpr(Expr_ Expression, Colon_ TokenValue, Expr_2 Expression) (*ColonExpr, error)

	// 471:2: colon_expr -> colon_expr_unit_tuple: ...
	ColonExprUnitTupleToColonExpr(ColonExpr_ *ColonExpr, Colon_ TokenValue) (*ColonExpr, error)

	// 472:2: colon_expr -> colon_expr_expr_tuple: ...
	ColonExprExprTupleToColonExpr(ColonExpr_ *ColonExpr, Colon_ TokenValue, Expr_ Expression) (*ColonExpr, error)
}

type ParseErrorExprReducer interface {
	// 491:32: parse_error_expr -> ...
	ToParseErrorExpr(ParseError_ ParseErrorSymbol) (Expression, error)
}

type LiteralExprReducer interface {
	// 494:2: literal_expr -> TRUE: ...
	TrueToLiteralExpr(True_ TokenValue) (Expression, error)

	// 495:2: literal_expr -> FALSE: ...
	FalseToLiteralExpr(False_ TokenValue) (Expression, error)

	// 496:2: literal_expr -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteralExpr(IntegerLiteral_ TokenValue) (Expression, error)

	// 497:2: literal_expr -> FLOAT_LITERAL: ...
	FloatLiteralToLiteralExpr(FloatLiteral_ TokenValue) (Expression, error)

	// 498:2: literal_expr -> RUNE_LITERAL: ...
	RuneLiteralToLiteralExpr(RuneLiteral_ TokenValue) (Expression, error)

	// 499:2: literal_expr -> STRING_LITERAL: ...
	StringLiteralToLiteralExpr(StringLiteral_ TokenValue) (Expression, error)
}

type NamedExprReducer interface {
	// 502:2: named_expr -> IDENTIFIER: ...
	IdentifierToNamedExpr(Identifier_ TokenValue) (Expression, error)

	// 503:2: named_expr -> UNDERSCORE: ...
	UnderscoreToNamedExpr(Underscore_ TokenValue) (Expression, error)
}

type BlockExprReducer interface {
	// 505:27: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)
}

type InitializeExprReducer interface {
	// 507:31: initialize_expr -> ...
	ToInitializeExpr(InitializableTypeExpr_ TypeExpression, Lparen_ TokenValue, Arguments_ *ArgumentList, Rparen_ TokenValue) (Expression, error)
}

type ImplicitStructExprReducer interface {
	// 509:36: implicit_struct_expr -> ...
	ToImplicitStructExpr(Lparen_ TokenValue, Arguments_ *ArgumentList, Rparen_ TokenValue) (Expression, error)
}

type AccessExprReducer interface {
	// 517:27: access_expr -> ...
	ToAccessExpr(AccessibleExpr_ Expression, Dot_ TokenValue, Identifier_ TokenValue) (Expression, error)
}

type IndexExprReducer interface {
	// 521:26: index_expr -> ...
	ToIndexExpr(AccessibleExpr_ Expression, Lbracket_ TokenValue, Argument_ *Argument, Rbracket_ TokenValue) (Expression, error)
}

type PostfixUnaryExprReducer interface {
	// 527:35: postfix_unary_expr -> ...
	ToPostfixUnaryExpr(AccessibleExpr_ Expression, Question_ TokenValue) (Expression, error)
}

type PrefixUnaryExprReducer interface {
	// 533:33: prefix_unary_expr -> ...
	ToPrefixUnaryExpr(PrefixUnaryOp_ TokenValue, PrefixableExpr_ Expression) (Expression, error)
}

type BinaryMulExprReducer interface {
	// 550:31: binary_mul_expr -> ...
	ToBinaryMulExpr(MulExpr_ Expression, MulOp_ TokenValue, PrefixableExpr_ Expression) (Expression, error)
}

type BinaryAddExprReducer interface {
	// 564:31: binary_add_expr -> ...
	ToBinaryAddExpr(AddExpr_ Expression, AddOp_ TokenValue, MulExpr_ Expression) (Expression, error)
}

type BinaryCmpExprReducer interface {
	// 576:31: binary_cmp_expr -> ...
	ToBinaryCmpExpr(CmpExpr_ Expression, CmpOp_ TokenValue, AddExpr_ Expression) (Expression, error)
}

type BinaryAndExprReducer interface {
	// 590:31: binary_and_expr -> ...
	ToBinaryAndExpr(AndExpr_ Expression, And_ TokenValue, CmpExpr_ Expression) (Expression, error)
}

type BinaryOrExprReducer interface {
	// 596:30: binary_or_expr -> ...
	ToBinaryOrExpr(OrExpr_ Expression, Or_ TokenValue, AndExpr_ Expression) (Expression, error)
}

type SliceTypeExprReducer interface {
	// 611:35: slice_type_expr -> ...
	ToSliceTypeExpr(Lbracket_ TokenValue, TypeExpr_ TypeExpression, Rbracket_ TokenValue) (TypeExpression, error)
}

type ArrayTypeExprReducer interface {
	// 614:2: array_type_expr -> ...
	ToArrayTypeExpr(Lbracket_ TokenValue, TypeExpr_ TypeExpression, Comma_ TokenValue, IntegerLiteral_ TokenValue, Rbracket_ TokenValue) (TypeExpression, error)
}

type MapTypeExprReducer interface {
	// 617:33: map_type_expr -> ...
	ToMapTypeExpr(Lbracket_ TokenValue, TypeExpr_ TypeExpression, Colon_ TokenValue, TypeExpr_2 TypeExpression, Rbracket_ TokenValue) (TypeExpression, error)
}

type NamedTypeExprReducer interface {
	// 631:2: named_type_expr -> local: ...
	LocalToNamedTypeExpr(Identifier_ TokenValue, GenericTypeArguments_ *TypeArgumentList) (TypeExpression, error)

	// 632:2: named_type_expr -> external: ...
	ExternalToNamedTypeExpr(Identifier_ TokenValue, Dot_ TokenValue, Identifier_2 TokenValue, GenericTypeArguments_ *TypeArgumentList) (TypeExpression, error)
}

type InferredTypeExprReducer interface {
	// 640:2: inferred_type_expr -> DOT: ...
	DotToInferredTypeExpr(Dot_ TokenValue) (TypeExpression, error)

	// 641:2: inferred_type_expr -> UNDERSCORE: ...
	UnderscoreToInferredTypeExpr(Underscore_ TokenValue) (TypeExpression, error)
}

type ParseErrorTypeExprReducer interface {
	// 643:41: parse_error_type_expr -> ...
	ToParseErrorTypeExpr(ParseError_ ParseErrorSymbol) (TypeExpression, error)
}

type PrefixedUnaryTypeExprReducer interface {
	// 653:2: prefixed_unary_type_expr -> ...
	ToPrefixedUnaryTypeExpr(PrefixUnaryTypeOp_ TokenValue, ReturnableTypeExpr_ TypeExpression) (TypeExpression, error)
}

type BinaryTypeExprReducer interface {
	// 669:2: binary_type_expr -> ...
	ToBinaryTypeExpr(TypeExpr_ TypeExpression, BinaryTypeOp_ TokenValue, ReturnableTypeExpr_ TypeExpression) (TypeExpression, error)
}

type TypeDefReducer interface {
	// 677:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, TypeExpr_ TypeExpression) (SourceDefinition, error)

	// 678:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, TypeExpr_ TypeExpression, Implements_ TokenValue, TypeExpr_2 TypeExpression) (SourceDefinition, error)

	// 679:2: type_def -> alias: ...
	AliasToTypeDef(Type_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, TypeExpr_ TypeExpression) (SourceDefinition, error)
}

type GenericParameterDefReducer interface {
	// 687:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ TokenValue) (GenericSymbol, error)

	// 688:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)
}

type ProperGenericParameterDefsReducer interface {
	// 691:2: proper_generic_parameter_defs -> add: ...
	AddToProperGenericParameterDefs(ProperGenericParameterDefs_ GenericSymbol, Comma_ TokenValue, GenericParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 692:2: proper_generic_parameter_defs -> generic_parameter_def: ...
	GenericParameterDefToProperGenericParameterDefs(GenericParameterDef_ GenericSymbol) (GenericSymbol, error)
}

type GenericParameterDefsReducer interface {

	// 696:2: generic_parameter_defs -> improper: ...
	ImproperToGenericParameterDefs(ProperGenericParameterDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 697:2: generic_parameter_defs -> nil: ...
	NilToGenericParameterDefs() (GenericSymbol, error)
}

type OptionalGenericParametersReducer interface {
	// 700:2: optional_generic_parameters -> generic: ...
	GenericToOptionalGenericParameters(DollarLbracket_ TokenValue, GenericParameterDefs_ GenericSymbol, Rbracket_ TokenValue) (GenericSymbol, error)

	// 701:2: optional_generic_parameters -> nil: ...
	NilToOptionalGenericParameters() (GenericSymbol, error)
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

type ParameterDeclReducer interface {
	// 800:2: parameter_decl -> unnamed_arg: ...
	UnnamedArgToParameterDecl(TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 801:2: parameter_decl -> unnamed_vararg: ...
	UnnamedVarargToParameterDecl(Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 802:2: parameter_decl -> underscore_arg: ...
	UnderscoreArgToParameterDecl(Underscore_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 803:2: parameter_decl -> underscore_vararg: ...
	UnderscoreVarargToParameterDecl(Underscore_ TokenValue, Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 804:2: parameter_decl -> arg: ...
	ArgToParameterDecl(Identifier_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 805:2: parameter_decl -> vararg: ...
	VarargToParameterDecl(Identifier_ TokenValue, Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)
}

type ProperParameterDeclsReducer interface {
	// 808:2: proper_parameter_decls -> add: ...
	AddToProperParameterDecls(ProperParameterDecls_ GenericSymbol, Comma_ TokenValue, ParameterDecl_ GenericSymbol) (GenericSymbol, error)

	// 809:2: proper_parameter_decls -> parameter_decl: ...
	ParameterDeclToProperParameterDecls(ParameterDecl_ GenericSymbol) (GenericSymbol, error)
}

type ParameterDeclsReducer interface {

	// 813:2: parameter_decls -> improper: ...
	ImproperToParameterDecls(ProperParameterDecls_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 814:2: parameter_decls -> nil: ...
	NilToParameterDecls() (GenericSymbol, error)
}

type FuncTypeExprReducer interface {
	// 817:2: func_type_expr -> ...
	ToFuncTypeExpr(Func_ TokenValue, Lparen_ TokenValue, ParameterDecls_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression) (TypeExpression, error)
}

type MethodSignatureReducer interface {
	// 828:20: method_signature -> ...
	ToMethodSignature(Func_ TokenValue, Identifier_ TokenValue, Lparen_ TokenValue, ParameterDecls_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression) (GenericSymbol, error)
}

type ParameterDefReducer interface {
	// 835:2: parameter_def -> inferred_ref_arg: ...
	InferredRefArgToParameterDef(Identifier_ TokenValue) (GenericSymbol, error)

	// 836:2: parameter_def -> inferred_ref_vararg: ...
	InferredRefVarargToParameterDef(Identifier_ TokenValue, Ellipsis_ TokenValue) (GenericSymbol, error)

	// 837:2: parameter_def -> arg: ...
	ArgToParameterDef(Identifier_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 838:2: parameter_def -> vararg: ...
	VarargToParameterDef(Identifier_ TokenValue, Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 839:2: parameter_def -> ignore_inferred_ref_arg: ...
	IgnoreInferredRefArgToParameterDef(Underscore_ TokenValue) (GenericSymbol, error)

	// 840:2: parameter_def -> ignore_inferred_ref_vararg: ...
	IgnoreInferredRefVarargToParameterDef(Underscore_ TokenValue, Ellipsis_ TokenValue) (GenericSymbol, error)

	// 841:2: parameter_def -> ignore_arg: ...
	IgnoreArgToParameterDef(Underscore_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 842:2: parameter_def -> ignore_vararg: ...
	IgnoreVarargToParameterDef(Underscore_ TokenValue, Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)
}

type ProperParameterDefsReducer interface {
	// 845:2: proper_parameter_defs -> add: ...
	AddToProperParameterDefs(ProperParameterDefs_ GenericSymbol, Comma_ TokenValue, ParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 846:2: proper_parameter_defs -> parameter_def: ...
	ParameterDefToProperParameterDefs(ParameterDef_ GenericSymbol) (GenericSymbol, error)
}

type ParameterDefsReducer interface {

	// 850:2: parameter_defs -> improper: ...
	ImproperToParameterDefs(ProperParameterDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 851:2: parameter_defs -> nil: ...
	NilToParameterDefs() (GenericSymbol, error)
}

type NamedFuncDefReducer interface {
	// 854:2: named_func_def -> func_def: ...
	FuncDefToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, Lparen_ TokenValue, ParameterDefs_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 855:2: named_func_def -> method_def: ...
	MethodDefToNamedFuncDef(Func_ TokenValue, Lparen_ TokenValue, ParameterDef_ GenericSymbol, Rparen_ TokenValue, Identifier_ TokenValue, OptionalGenericParameters_ GenericSymbol, Lparen_2 TokenValue, ParameterDefs_ GenericSymbol, Rparen_2 TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 856:2: named_func_def -> alias: ...
	AliasToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, Expr_ Expression) (SourceDefinition, error)
}

type AnonymousFuncExprReducer interface {
	// 860:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ TokenValue, Lparen_ TokenValue, ParameterDefs_ GenericSymbol, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (Expression, error)
}

type PackageDefReducer interface {
	// 872:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ TokenValue) (SourceDefinition, error)

	// 873:2: package_def -> with_spec: ...
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
	GenericTypeArgumentsReducer
	ProperTypeArgumentsReducer
	TypeArgumentsReducer
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
	PrefixedUnaryTypeExprReducer
	BinaryTypeExprReducer
	TypeDefReducer
	GenericParameterDefReducer
	ProperGenericParameterDefsReducer
	GenericParameterDefsReducer
	OptionalGenericParametersReducer
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
	ParameterDeclReducer
	ProperParameterDeclsReducer
	ParameterDeclsReducer
	FuncTypeExprReducer
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
	case GenericTypeArgumentsType:
		return "generic_type_arguments"
	case ProperTypeArgumentsType:
		return "proper_type_arguments"
	case TypeArgumentsType:
		return "type_arguments"
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
	case PrefixedUnaryTypeExprType:
		return "prefixed_unary_type_expr"
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
	case ProperGenericParameterDefsType:
		return "proper_generic_parameter_defs"
	case GenericParameterDefsType:
		return "generic_parameter_defs"
	case OptionalGenericParametersType:
		return "optional_generic_parameters"
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
	case ParameterDeclType:
		return "parameter_decl"
	case ProperParameterDeclsType:
		return "proper_parameter_decls"
	case ParameterDeclsType:
		return "parameter_decls"
	case FuncTypeExprType:
		return "func_type_expr"
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
	GenericTypeArgumentsType          = SymbolId(390)
	ProperTypeArgumentsType           = SymbolId(391)
	TypeArgumentsType                 = SymbolId(392)
	ProperArgumentsType               = SymbolId(393)
	ArgumentsType                     = SymbolId(394)
	ArgumentType                      = SymbolId(395)
	ColonExprType                     = SymbolId(396)
	AtomExprType                      = SymbolId(397)
	ParseErrorExprType                = SymbolId(398)
	LiteralExprType                   = SymbolId(399)
	NamedExprType                     = SymbolId(400)
	BlockExprType                     = SymbolId(401)
	InitializeExprType                = SymbolId(402)
	ImplicitStructExprType            = SymbolId(403)
	AccessibleExprType                = SymbolId(404)
	AccessExprType                    = SymbolId(405)
	IndexExprType                     = SymbolId(406)
	PostfixableExprType               = SymbolId(407)
	PostfixUnaryExprType              = SymbolId(408)
	PrefixableExprType                = SymbolId(409)
	PrefixUnaryExprType               = SymbolId(410)
	PrefixUnaryOpType                 = SymbolId(411)
	MulExprType                       = SymbolId(412)
	BinaryMulExprType                 = SymbolId(413)
	MulOpType                         = SymbolId(414)
	AddExprType                       = SymbolId(415)
	BinaryAddExprType                 = SymbolId(416)
	AddOpType                         = SymbolId(417)
	CmpExprType                       = SymbolId(418)
	BinaryCmpExprType                 = SymbolId(419)
	CmpOpType                         = SymbolId(420)
	AndExprType                       = SymbolId(421)
	BinaryAndExprType                 = SymbolId(422)
	OrExprType                        = SymbolId(423)
	BinaryOrExprType                  = SymbolId(424)
	InitializableTypeExprType         = SymbolId(425)
	SliceTypeExprType                 = SymbolId(426)
	ArrayTypeExprType                 = SymbolId(427)
	MapTypeExprType                   = SymbolId(428)
	AtomTypeExprType                  = SymbolId(429)
	NamedTypeExprType                 = SymbolId(430)
	InferredTypeExprType              = SymbolId(431)
	ParseErrorTypeExprType            = SymbolId(432)
	ReturnableTypeExprType            = SymbolId(433)
	PrefixedUnaryTypeExprType         = SymbolId(434)
	PrefixUnaryTypeOpType             = SymbolId(435)
	TypeExprType                      = SymbolId(436)
	BinaryTypeExprType                = SymbolId(437)
	BinaryTypeOpType                  = SymbolId(438)
	TypeDefType                       = SymbolId(439)
	GenericParameterDefType           = SymbolId(440)
	ProperGenericParameterDefsType    = SymbolId(441)
	GenericParameterDefsType          = SymbolId(442)
	OptionalGenericParametersType     = SymbolId(443)
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
	ParameterDeclType                 = SymbolId(460)
	ProperParameterDeclsType          = SymbolId(461)
	ParameterDeclsType                = SymbolId(462)
	FuncTypeExprType                  = SymbolId(463)
	MethodSignatureType               = SymbolId(464)
	ParameterDefType                  = SymbolId(465)
	ProperParameterDefsType           = SymbolId(466)
	ParameterDefsType                 = SymbolId(467)
	NamedFuncDefType                  = SymbolId(468)
	AnonymousFuncExprType             = SymbolId(469)
	PackageDefType                    = SymbolId(470)
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
	_ReduceAddToProperDefinitions                             = _ReduceType(2)
	_ReduceDefinitionToProperDefinitions                      = _ReduceType(3)
	_ReduceProperDefinitionsToDefinitions                     = _ReduceType(4)
	_ReduceImproperToDefinitions                              = _ReduceType(5)
	_ReduceNilToDefinitions                                   = _ReduceType(6)
	_ReducePackageDefToDefinition                             = _ReduceType(7)
	_ReduceTypeDefToDefinition                                = _ReduceType(8)
	_ReduceNamedFuncDefToDefinition                           = _ReduceType(9)
	_ReduceGlobalVarDefToDefinition                           = _ReduceType(10)
	_ReduceGlobalVarAssignmentToDefinition                    = _ReduceType(11)
	_ReduceStatementBlockToDefinition                         = _ReduceType(12)
	_ReduceCommentGroupsToDefinition                          = _ReduceType(13)
	_ReduceToStatementBlock                                   = _ReduceType(14)
	_ReduceAddImplicitToProperStatements                      = _ReduceType(15)
	_ReduceAddExplicitToProperStatements                      = _ReduceType(16)
	_ReduceStatementToProperStatements                        = _ReduceType(17)
	_ReduceProperStatementsToStatements                       = _ReduceType(18)
	_ReduceImproperImplicitToStatements                       = _ReduceType(19)
	_ReduceImproperExplicitToStatements                       = _ReduceType(20)
	_ReduceNilToStatements                                    = _ReduceType(21)
	_ReduceUnsafeStatementToSimpleStatement                   = _ReduceType(22)
	_ReduceExprOrImproperStructStatementToSimpleStatement     = _ReduceType(23)
	_ReduceCallbackOpStatementToSimpleStatement               = _ReduceType(24)
	_ReduceJumpStatementToSimpleStatement                     = _ReduceType(25)
	_ReduceAssignStatementToSimpleStatement                   = _ReduceType(26)
	_ReduceUnaryOpAssignStatementToSimpleStatement            = _ReduceType(27)
	_ReduceBinaryOpAssignStatementToSimpleStatement           = _ReduceType(28)
	_ReduceFallthroughStatementToSimpleStatement              = _ReduceType(29)
	_ReduceSimpleStatementToStatement                         = _ReduceType(30)
	_ReduceImportStatementToStatement                         = _ReduceType(31)
	_ReduceCaseBranchStatementToStatement                     = _ReduceType(32)
	_ReduceDefaultBranchStatementToStatement                  = _ReduceType(33)
	_ReduceSimpleStatementToOptionalSimpleStatement           = _ReduceType(34)
	_ReduceNilToOptionalSimpleStatement                       = _ReduceType(35)
	_ReduceToExprOrImproperStructStatement                    = _ReduceType(36)
	_ReduceExprToExprs                                        = _ReduceType(37)
	_ReduceAddToExprs                                         = _ReduceType(38)
	_ReduceAsyncToCallbackOp                                  = _ReduceType(39)
	_ReduceDeferToCallbackOp                                  = _ReduceType(40)
	_ReduceToCallbackOpStatement                              = _ReduceType(41)
	_ReduceToUnsafeStatement                                  = _ReduceType(42)
	_ReduceUnlabeledNoValueToJumpStatement                    = _ReduceType(43)
	_ReduceUnlabeledValuedToJumpStatement                     = _ReduceType(44)
	_ReduceLabeledNoValueToJumpStatement                      = _ReduceType(45)
	_ReduceLabeledValuedToJumpStatement                       = _ReduceType(46)
	_ReduceReturnToJumpType                                   = _ReduceType(47)
	_ReduceBreakToJumpType                                    = _ReduceType(48)
	_ReduceContinueToJumpType                                 = _ReduceType(49)
	_ReduceToFallthroughStatement                             = _ReduceType(50)
	_ReduceToAssignStatement                                  = _ReduceType(51)
	_ReduceToUnaryOpAssignStatement                           = _ReduceType(52)
	_ReduceAddOneAssignToUnaryOpAssign                        = _ReduceType(53)
	_ReduceSubOneAssignToUnaryOpAssign                        = _ReduceType(54)
	_ReduceToBinaryOpAssignStatement                          = _ReduceType(55)
	_ReduceAddAssignToBinaryOpAssign                          = _ReduceType(56)
	_ReduceSubAssignToBinaryOpAssign                          = _ReduceType(57)
	_ReduceMulAssignToBinaryOpAssign                          = _ReduceType(58)
	_ReduceDivAssignToBinaryOpAssign                          = _ReduceType(59)
	_ReduceModAssignToBinaryOpAssign                          = _ReduceType(60)
	_ReduceBitNegAssignToBinaryOpAssign                       = _ReduceType(61)
	_ReduceBitAndAssignToBinaryOpAssign                       = _ReduceType(62)
	_ReduceBitOrAssignToBinaryOpAssign                        = _ReduceType(63)
	_ReduceBitXorAssignToBinaryOpAssign                       = _ReduceType(64)
	_ReduceBitLshiftAssignToBinaryOpAssign                    = _ReduceType(65)
	_ReduceBitRshiftAssignToBinaryOpAssign                    = _ReduceType(66)
	_ReduceSingleToImportStatement                            = _ReduceType(67)
	_ReduceMultipleToImportStatement                          = _ReduceType(68)
	_ReduceStringLiteralToImportClause                        = _ReduceType(69)
	_ReduceAliasToImportClause                                = _ReduceType(70)
	_ReduceUnusableImportToImportClause                       = _ReduceType(71)
	_ReduceImportToLocalToImportClause                        = _ReduceType(72)
	_ReduceAddImplicitToProperImportClauses                   = _ReduceType(73)
	_ReduceAddExplicitToProperImportClauses                   = _ReduceType(74)
	_ReduceImportClauseToProperImportClauses                  = _ReduceType(75)
	_ReduceProperImportClausesToImportClauses                 = _ReduceType(76)
	_ReduceImplicitToImportClauses                            = _ReduceType(77)
	_ReduceExplicitToImportClauses                            = _ReduceType(78)
	_ReduceCasePatternToCasePatterns                          = _ReduceType(79)
	_ReduceMultipleToCasePatterns                             = _ReduceType(80)
	_ReduceToVarDeclPattern                                   = _ReduceType(81)
	_ReduceVarToVarOrLet                                      = _ReduceType(82)
	_ReduceLetToVarOrLet                                      = _ReduceType(83)
	_ReduceIdentifierToVarPattern                             = _ReduceType(84)
	_ReduceUnderscoreToVarPattern                             = _ReduceType(85)
	_ReduceTuplePatternToVarPattern                           = _ReduceType(86)
	_ReduceToTuplePattern                                     = _ReduceType(87)
	_ReduceFieldVarPatternToFieldVarPatterns                  = _ReduceType(88)
	_ReduceAddToFieldVarPatterns                              = _ReduceType(89)
	_ReducePositionalToFieldVarPattern                        = _ReduceType(90)
	_ReduceNamedToFieldVarPattern                             = _ReduceType(91)
	_ReduceEllipsisToFieldVarPattern                          = _ReduceType(92)
	_ReduceTypeExprToOptionalTypeExpr                         = _ReduceType(93)
	_ReduceNilToOptionalTypeExpr                              = _ReduceType(94)
	_ReduceToAssignPattern                                    = _ReduceType(95)
	_ReduceAssignPatternToCasePattern                         = _ReduceType(96)
	_ReduceEnumMatchPatternToCasePattern                      = _ReduceType(97)
	_ReduceEnumNondataMatchPattenToCasePattern                = _ReduceType(98)
	_ReduceEnumVarDeclPatternToCasePattern                    = _ReduceType(99)
	_ReduceIfExprToExpr                                       = _ReduceType(100)
	_ReduceSwitchExprToExpr                                   = _ReduceType(101)
	_ReduceLoopExprToExpr                                     = _ReduceType(102)
	_ReduceSequenceExprToExpr                                 = _ReduceType(103)
	_ReduceLabelDeclToOptionalLabelDecl                       = _ReduceType(104)
	_ReduceUnlabelledToOptionalLabelDecl                      = _ReduceType(105)
	_ReduceOrExprToSequenceExpr                               = _ReduceType(106)
	_ReduceVarDeclPatternToSequenceExpr                       = _ReduceType(107)
	_ReduceAssignVarPatternToSequenceExpr                     = _ReduceType(108)
	_ReduceNoElseToIfExpr                                     = _ReduceType(109)
	_ReduceIfElseToIfExpr                                     = _ReduceType(110)
	_ReduceMultiIfElseToIfExpr                                = _ReduceType(111)
	_ReduceSequenceExprToCondition                            = _ReduceType(112)
	_ReduceCaseToCondition                                    = _ReduceType(113)
	_ReduceToSwitchExpr                                       = _ReduceType(114)
	_ReduceInfiniteToLoopExpr                                 = _ReduceType(115)
	_ReduceDoWhileToLoopExpr                                  = _ReduceType(116)
	_ReduceWhileToLoopExpr                                    = _ReduceType(117)
	_ReduceIteratorToLoopExpr                                 = _ReduceType(118)
	_ReduceForToLoopExpr                                      = _ReduceType(119)
	_ReduceSequenceExprToOptionalSequenceExpr                 = _ReduceType(120)
	_ReduceNilToOptionalSequenceExpr                          = _ReduceType(121)
	_ReduceSequenceExprToForAssignment                        = _ReduceType(122)
	_ReduceAssignToForAssignment                              = _ReduceType(123)
	_ReduceToCallExpr                                         = _ReduceType(124)
	_ReduceBindingToGenericTypeArguments                      = _ReduceType(125)
	_ReduceNilToGenericTypeArguments                          = _ReduceType(126)
	_ReduceAddToProperTypeArguments                           = _ReduceType(127)
	_ReduceTypeExprToProperTypeArguments                      = _ReduceType(128)
	_ReduceProperTypeArgumentsToTypeArguments                 = _ReduceType(129)
	_ReduceImproperToTypeArguments                            = _ReduceType(130)
	_ReduceNilToTypeArguments                                 = _ReduceType(131)
	_ReduceAddToProperArguments                               = _ReduceType(132)
	_ReduceArgumentToProperArguments                          = _ReduceType(133)
	_ReduceProperArgumentsToArguments                         = _ReduceType(134)
	_ReduceImproperToArguments                                = _ReduceType(135)
	_ReduceNilToArguments                                     = _ReduceType(136)
	_ReducePositionalToArgument                               = _ReduceType(137)
	_ReduceColonExprToArgument                                = _ReduceType(138)
	_ReduceNamedAssignmentToArgument                          = _ReduceType(139)
	_ReduceVarargAssignmentToArgument                         = _ReduceType(140)
	_ReduceSkipPatternToArgument                              = _ReduceType(141)
	_ReduceUnitUnitPairToColonExpr                            = _ReduceType(142)
	_ReduceExprUnitPairToColonExpr                            = _ReduceType(143)
	_ReduceUnitExprPairToColonExpr                            = _ReduceType(144)
	_ReduceExprExprPairToColonExpr                            = _ReduceType(145)
	_ReduceColonExprUnitTupleToColonExpr                      = _ReduceType(146)
	_ReduceColonExprExprTupleToColonExpr                      = _ReduceType(147)
	_ReduceParseErrorExprToAtomExpr                           = _ReduceType(148)
	_ReduceLiteralExprToAtomExpr                              = _ReduceType(149)
	_ReduceNamedExprToAtomExpr                                = _ReduceType(150)
	_ReduceBlockExprToAtomExpr                                = _ReduceType(151)
	_ReduceAnonymousFuncExprToAtomExpr                        = _ReduceType(152)
	_ReduceInitializeExprToAtomExpr                           = _ReduceType(153)
	_ReduceImplicitStructExprToAtomExpr                       = _ReduceType(154)
	_ReduceToParseErrorExpr                                   = _ReduceType(155)
	_ReduceTrueToLiteralExpr                                  = _ReduceType(156)
	_ReduceFalseToLiteralExpr                                 = _ReduceType(157)
	_ReduceIntegerLiteralToLiteralExpr                        = _ReduceType(158)
	_ReduceFloatLiteralToLiteralExpr                          = _ReduceType(159)
	_ReduceRuneLiteralToLiteralExpr                           = _ReduceType(160)
	_ReduceStringLiteralToLiteralExpr                         = _ReduceType(161)
	_ReduceIdentifierToNamedExpr                              = _ReduceType(162)
	_ReduceUnderscoreToNamedExpr                              = _ReduceType(163)
	_ReduceToBlockExpr                                        = _ReduceType(164)
	_ReduceToInitializeExpr                                   = _ReduceType(165)
	_ReduceToImplicitStructExpr                               = _ReduceType(166)
	_ReduceAtomExprToAccessibleExpr                           = _ReduceType(167)
	_ReduceAccessExprToAccessibleExpr                         = _ReduceType(168)
	_ReduceCallExprToAccessibleExpr                           = _ReduceType(169)
	_ReduceIndexExprToAccessibleExpr                          = _ReduceType(170)
	_ReduceToAccessExpr                                       = _ReduceType(171)
	_ReduceToIndexExpr                                        = _ReduceType(172)
	_ReduceAccessibleExprToPostfixableExpr                    = _ReduceType(173)
	_ReducePostfixUnaryExprToPostfixableExpr                  = _ReduceType(174)
	_ReduceToPostfixUnaryExpr                                 = _ReduceType(175)
	_ReducePostfixableExprToPrefixableExpr                    = _ReduceType(176)
	_ReducePrefixUnaryExprToPrefixableExpr                    = _ReduceType(177)
	_ReduceToPrefixUnaryExpr                                  = _ReduceType(178)
	_ReduceNotToPrefixUnaryOp                                 = _ReduceType(179)
	_ReduceBitNegToPrefixUnaryOp                              = _ReduceType(180)
	_ReduceSubToPrefixUnaryOp                                 = _ReduceType(181)
	_ReduceMulToPrefixUnaryOp                                 = _ReduceType(182)
	_ReduceBitAndToPrefixUnaryOp                              = _ReduceType(183)
	_ReducePrefixableExprToMulExpr                            = _ReduceType(184)
	_ReduceBinaryMulExprToMulExpr                             = _ReduceType(185)
	_ReduceToBinaryMulExpr                                    = _ReduceType(186)
	_ReduceMulToMulOp                                         = _ReduceType(187)
	_ReduceDivToMulOp                                         = _ReduceType(188)
	_ReduceModToMulOp                                         = _ReduceType(189)
	_ReduceBitAndToMulOp                                      = _ReduceType(190)
	_ReduceBitLshiftToMulOp                                   = _ReduceType(191)
	_ReduceBitRshiftToMulOp                                   = _ReduceType(192)
	_ReduceMulExprToAddExpr                                   = _ReduceType(193)
	_ReduceBinaryAddExprToAddExpr                             = _ReduceType(194)
	_ReduceToBinaryAddExpr                                    = _ReduceType(195)
	_ReduceAddToAddOp                                         = _ReduceType(196)
	_ReduceSubToAddOp                                         = _ReduceType(197)
	_ReduceBitOrToAddOp                                       = _ReduceType(198)
	_ReduceBitXorToAddOp                                      = _ReduceType(199)
	_ReduceAddExprToCmpExpr                                   = _ReduceType(200)
	_ReduceBinaryCmpExprToCmpExpr                             = _ReduceType(201)
	_ReduceToBinaryCmpExpr                                    = _ReduceType(202)
	_ReduceEqualToCmpOp                                       = _ReduceType(203)
	_ReduceNotEqualToCmpOp                                    = _ReduceType(204)
	_ReduceLessToCmpOp                                        = _ReduceType(205)
	_ReduceLessOrEqualToCmpOp                                 = _ReduceType(206)
	_ReduceGreaterToCmpOp                                     = _ReduceType(207)
	_ReduceGreaterOrEqualToCmpOp                              = _ReduceType(208)
	_ReduceCmpExprToAndExpr                                   = _ReduceType(209)
	_ReduceBinaryAndExprToAndExpr                             = _ReduceType(210)
	_ReduceToBinaryAndExpr                                    = _ReduceType(211)
	_ReduceAndExprToOrExpr                                    = _ReduceType(212)
	_ReduceBinaryOrExprToOrExpr                               = _ReduceType(213)
	_ReduceToBinaryOrExpr                                     = _ReduceType(214)
	_ReduceExplicitStructTypeExprToInitializableTypeExpr      = _ReduceType(215)
	_ReduceSliceTypeExprToInitializableTypeExpr               = _ReduceType(216)
	_ReduceArrayTypeExprToInitializableTypeExpr               = _ReduceType(217)
	_ReduceMapTypeExprToInitializableTypeExpr                 = _ReduceType(218)
	_ReduceToSliceTypeExpr                                    = _ReduceType(219)
	_ReduceToArrayTypeExpr                                    = _ReduceType(220)
	_ReduceToMapTypeExpr                                      = _ReduceType(221)
	_ReduceInitializableTypeExprToAtomTypeExpr                = _ReduceType(222)
	_ReduceNamedTypeExprToAtomTypeExpr                        = _ReduceType(223)
	_ReduceInferredTypeExprToAtomTypeExpr                     = _ReduceType(224)
	_ReduceImplicitStructTypeExprToAtomTypeExpr               = _ReduceType(225)
	_ReduceExplicitEnumTypeExprToAtomTypeExpr                 = _ReduceType(226)
	_ReduceImplicitEnumTypeExprToAtomTypeExpr                 = _ReduceType(227)
	_ReduceTraitTypeExprToAtomTypeExpr                        = _ReduceType(228)
	_ReduceFuncTypeExprToAtomTypeExpr                         = _ReduceType(229)
	_ReduceParseErrorTypeExprToAtomTypeExpr                   = _ReduceType(230)
	_ReduceLocalToNamedTypeExpr                               = _ReduceType(231)
	_ReduceExternalToNamedTypeExpr                            = _ReduceType(232)
	_ReduceDotToInferredTypeExpr                              = _ReduceType(233)
	_ReduceUnderscoreToInferredTypeExpr                       = _ReduceType(234)
	_ReduceToParseErrorTypeExpr                               = _ReduceType(235)
	_ReduceAtomTypeExprToReturnableTypeExpr                   = _ReduceType(236)
	_ReducePrefixedUnaryTypeExprToReturnableTypeExpr          = _ReduceType(237)
	_ReduceToPrefixedUnaryTypeExpr                            = _ReduceType(238)
	_ReduceQuestionToPrefixUnaryTypeOp                        = _ReduceType(239)
	_ReduceExclaimToPrefixUnaryTypeOp                         = _ReduceType(240)
	_ReduceBitAndToPrefixUnaryTypeOp                          = _ReduceType(241)
	_ReduceBitNegToPrefixUnaryTypeOp                          = _ReduceType(242)
	_ReduceTildeTildeToPrefixUnaryTypeOp                      = _ReduceType(243)
	_ReduceReturnableTypeExprToTypeExpr                       = _ReduceType(244)
	_ReduceBinaryTypeExprToTypeExpr                           = _ReduceType(245)
	_ReduceToBinaryTypeExpr                                   = _ReduceType(246)
	_ReduceMulToBinaryTypeOp                                  = _ReduceType(247)
	_ReduceAddToBinaryTypeOp                                  = _ReduceType(248)
	_ReduceSubToBinaryTypeOp                                  = _ReduceType(249)
	_ReduceDefinitionToTypeDef                                = _ReduceType(250)
	_ReduceConstrainedDefToTypeDef                            = _ReduceType(251)
	_ReduceAliasToTypeDef                                     = _ReduceType(252)
	_ReduceUnconstrainedToGenericParameterDef                 = _ReduceType(253)
	_ReduceConstrainedToGenericParameterDef                   = _ReduceType(254)
	_ReduceAddToProperGenericParameterDefs                    = _ReduceType(255)
	_ReduceGenericParameterDefToProperGenericParameterDefs    = _ReduceType(256)
	_ReduceProperGenericParameterDefsToGenericParameterDefs   = _ReduceType(257)
	_ReduceImproperToGenericParameterDefs                     = _ReduceType(258)
	_ReduceNilToGenericParameterDefs                          = _ReduceType(259)
	_ReduceGenericToOptionalGenericParameters                 = _ReduceType(260)
	_ReduceNilToOptionalGenericParameters                     = _ReduceType(261)
	_ReduceNamedToFieldDef                                    = _ReduceType(262)
	_ReduceUnnamedToFieldDef                                  = _ReduceType(263)
	_ReduceStructPaddingToFieldDef                            = _ReduceType(264)
	_ReduceMethodSignatureToFieldDef                          = _ReduceType(265)
	_ReduceUnsafeStatementToFieldDef                          = _ReduceType(266)
	_ReduceDefaultToOptionalDefault                           = _ReduceType(267)
	_ReduceNilToOptionalDefault                               = _ReduceType(268)
	_ReduceAddToProperImplicitFieldDefs                       = _ReduceType(269)
	_ReduceFieldDefToProperImplicitFieldDefs                  = _ReduceType(270)
	_ReduceProperImplicitFieldDefsToImplicitFieldDefs         = _ReduceType(271)
	_ReduceImproperToImplicitFieldDefs                        = _ReduceType(272)
	_ReduceNilToImplicitFieldDefs                             = _ReduceType(273)
	_ReduceToImplicitStructTypeExpr                           = _ReduceType(274)
	_ReduceAddImplicitToProperExplicitFieldDefs               = _ReduceType(275)
	_ReduceAddExplicitToProperExplicitFieldDefs               = _ReduceType(276)
	_ReduceFieldDefToProperExplicitFieldDefs                  = _ReduceType(277)
	_ReduceProperExplicitFieldDefsToExplicitFieldDefs         = _ReduceType(278)
	_ReduceImproperImplicitToExplicitFieldDefs                = _ReduceType(279)
	_ReduceImproperExplicitToExplicitFieldDefs                = _ReduceType(280)
	_ReduceNilToExplicitFieldDefs                             = _ReduceType(281)
	_ReduceToExplicitStructTypeExpr                           = _ReduceType(282)
	_ReduceToTraitTypeExpr                                    = _ReduceType(283)
	_ReducePairToProperImplicitEnumValueDefs                  = _ReduceType(284)
	_ReduceAddToProperImplicitEnumValueDefs                   = _ReduceType(285)
	_ReduceProperImplicitEnumValueDefsToImplicitEnumValueDefs = _ReduceType(286)
	_ReduceImproperToImplicitEnumValueDefs                    = _ReduceType(287)
	_ReduceToImplicitEnumTypeExpr                             = _ReduceType(288)
	_ReduceExplicitPairToProperExplicitEnumValueDefs          = _ReduceType(289)
	_ReduceImplicitPairToProperExplicitEnumValueDefs          = _ReduceType(290)
	_ReduceExplicitAddToProperExplicitEnumValueDefs           = _ReduceType(291)
	_ReduceImplicitAddToProperExplicitEnumValueDefs           = _ReduceType(292)
	_ReduceProperExplicitEnumValueDefsToExplicitEnumValueDefs = _ReduceType(293)
	_ReduceImproperToExplicitEnumValueDefs                    = _ReduceType(294)
	_ReduceToExplicitEnumTypeExpr                             = _ReduceType(295)
	_ReduceReturnableTypeExprToReturnType                     = _ReduceType(296)
	_ReduceNilToReturnType                                    = _ReduceType(297)
	_ReduceUnnamedArgToParameterDecl                          = _ReduceType(298)
	_ReduceUnnamedVarargToParameterDecl                       = _ReduceType(299)
	_ReduceUnderscoreArgToParameterDecl                       = _ReduceType(300)
	_ReduceUnderscoreVarargToParameterDecl                    = _ReduceType(301)
	_ReduceArgToParameterDecl                                 = _ReduceType(302)
	_ReduceVarargToParameterDecl                              = _ReduceType(303)
	_ReduceAddToProperParameterDecls                          = _ReduceType(304)
	_ReduceParameterDeclToProperParameterDecls                = _ReduceType(305)
	_ReduceProperParameterDeclsToParameterDecls               = _ReduceType(306)
	_ReduceImproperToParameterDecls                           = _ReduceType(307)
	_ReduceNilToParameterDecls                                = _ReduceType(308)
	_ReduceToFuncTypeExpr                                     = _ReduceType(309)
	_ReduceToMethodSignature                                  = _ReduceType(310)
	_ReduceInferredRefArgToParameterDef                       = _ReduceType(311)
	_ReduceInferredRefVarargToParameterDef                    = _ReduceType(312)
	_ReduceArgToParameterDef                                  = _ReduceType(313)
	_ReduceVarargToParameterDef                               = _ReduceType(314)
	_ReduceIgnoreInferredRefArgToParameterDef                 = _ReduceType(315)
	_ReduceIgnoreInferredRefVarargToParameterDef              = _ReduceType(316)
	_ReduceIgnoreArgToParameterDef                            = _ReduceType(317)
	_ReduceIgnoreVarargToParameterDef                         = _ReduceType(318)
	_ReduceAddToProperParameterDefs                           = _ReduceType(319)
	_ReduceParameterDefToProperParameterDefs                  = _ReduceType(320)
	_ReduceProperParameterDefsToParameterDefs                 = _ReduceType(321)
	_ReduceImproperToParameterDefs                            = _ReduceType(322)
	_ReduceNilToParameterDefs                                 = _ReduceType(323)
	_ReduceFuncDefToNamedFuncDef                              = _ReduceType(324)
	_ReduceMethodDefToNamedFuncDef                            = _ReduceType(325)
	_ReduceAliasToNamedFuncDef                                = _ReduceType(326)
	_ReduceToAnonymousFuncExpr                                = _ReduceType(327)
	_ReduceNoSpecToPackageDef                                 = _ReduceType(328)
	_ReduceWithSpecToPackageDef                               = _ReduceType(329)
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
	case _ReducePrefixedUnaryTypeExprToReturnableTypeExpr:
		return "PrefixedUnaryTypeExprToReturnableTypeExpr"
	case _ReduceToPrefixedUnaryTypeExpr:
		return "ToPrefixedUnaryTypeExpr"
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
	case _ReduceUnnamedArgToParameterDecl:
		return "UnnamedArgToParameterDecl"
	case _ReduceUnnamedVarargToParameterDecl:
		return "UnnamedVarargToParameterDecl"
	case _ReduceUnderscoreArgToParameterDecl:
		return "UnderscoreArgToParameterDecl"
	case _ReduceUnderscoreVarargToParameterDecl:
		return "UnderscoreVarargToParameterDecl"
	case _ReduceArgToParameterDecl:
		return "ArgToParameterDecl"
	case _ReduceVarargToParameterDecl:
		return "VarargToParameterDecl"
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
	case _ReduceToFuncTypeExpr:
		return "ToFuncTypeExpr"
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
	case _ReduceIgnoreInferredRefArgToParameterDef:
		return "IgnoreInferredRefArgToParameterDef"
	case _ReduceIgnoreInferredRefVarargToParameterDef:
		return "IgnoreInferredRefVarargToParameterDef"
	case _ReduceIgnoreArgToParameterDef:
		return "IgnoreArgToParameterDef"
	case _ReduceIgnoreVarargToParameterDef:
		return "IgnoreVarargToParameterDef"
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
	_State471 = _StateId(471)
	_State472 = _StateId(472)
	_State473 = _StateId(473)
	_State474 = _StateId(474)
	_State475 = _StateId(475)
	_State476 = _StateId(476)
	_State477 = _StateId(477)
	_State478 = _StateId(478)
	_State479 = _StateId(479)
	_State480 = _StateId(480)
	_State481 = _StateId(481)
	_State482 = _StateId(482)
	_State483 = _StateId(483)
)

type Symbol struct {
	SymbolId_ SymbolId

	Generic_ GenericSymbol

	Argument          *Argument
	ArgumentList      *ArgumentList
	ColonExpr         *ColonExpr
	Count             TokenCount
	Expression        Expression
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
	case OptionalTypeExprType, InitializableTypeExprType, SliceTypeExprType, ArrayTypeExprType, MapTypeExprType, AtomTypeExprType, NamedTypeExprType, InferredTypeExprType, ParseErrorTypeExprType, ReturnableTypeExprType, PrefixedUnaryTypeExprType, TypeExprType, BinaryTypeExprType, ImplicitStructTypeExprType, ExplicitStructTypeExprType, TraitTypeExprType, ImplicitEnumTypeExprType, ExplicitEnumTypeExprType, ReturnTypeType, FuncTypeExprType:
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
	case OptionalTypeExprType, InitializableTypeExprType, SliceTypeExprType, ArrayTypeExprType, MapTypeExprType, AtomTypeExprType, NamedTypeExprType, InferredTypeExprType, ParseErrorTypeExprType, ReturnableTypeExprType, PrefixedUnaryTypeExprType, TypeExprType, BinaryTypeExprType, ImplicitStructTypeExprType, ExplicitStructTypeExprType, TraitTypeExprType, ImplicitEnumTypeExprType, ExplicitEnumTypeExprType, ReturnTypeType, FuncTypeExprType:
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
		//line grammar.lr:440:4
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
		//line grammar.lr:449:4
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
		//line grammar.lr:483:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceLiteralExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:484:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceNamedExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:485:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBlockExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:486:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceAnonymousFuncExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:487:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceInitializeExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:488:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceImplicitStructExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:489:4
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
		//line grammar.lr:512:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceAccessExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:513:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceCallExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:514:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceIndexExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:515:4
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
		//line grammar.lr:524:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReducePostfixUnaryExprToPostfixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixableExprType
		//line grammar.lr:525:4
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
		//line grammar.lr:530:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReducePrefixUnaryExprToPrefixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixableExprType
		//line grammar.lr:531:4
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
		//line grammar.lr:536:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:537:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:538:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:541:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:544:4
		symbol.Value = args[0].Value
		err = nil
	case _ReducePrefixableExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		//line grammar.lr:547:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryMulExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		//line grammar.lr:548:4
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
		//line grammar.lr:553:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDivToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:554:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceModToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:555:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:556:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitLshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:557:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitRshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:558:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulExprToAddExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddExprType
		//line grammar.lr:561:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryAddExprToAddExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddExprType
		//line grammar.lr:562:4
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
		//line grammar.lr:567:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:568:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitOrToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:569:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitXorToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:570:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddExprToCmpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpExprType
		//line grammar.lr:573:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryCmpExprToCmpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpExprType
		//line grammar.lr:574:4
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
		//line grammar.lr:579:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceNotEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:580:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLessToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:581:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLessOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:582:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceGreaterToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:583:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceGreaterOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:584:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceCmpExprToAndExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AndExprType
		//line grammar.lr:587:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryAndExprToAndExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AndExprType
		//line grammar.lr:588:4
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
		//line grammar.lr:593:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryOrExprToOrExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OrExprType
		//line grammar.lr:594:4
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
		//line grammar.lr:605:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceSliceTypeExprToInitializableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeExprType
		//line grammar.lr:606:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceArrayTypeExprToInitializableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeExprType
		//line grammar.lr:607:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceMapTypeExprToInitializableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeExprType
		//line grammar.lr:608:4
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
		//line grammar.lr:620:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceNamedTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:621:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceInferredTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:622:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceImplicitStructTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:623:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceExplicitEnumTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:624:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceImplicitEnumTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:625:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceTraitTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:626:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceFuncTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:627:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceParseErrorTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:628:4
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
		//line grammar.lr:649:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReducePrefixedUnaryTypeExprToReturnableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableTypeExprType
		//line grammar.lr:650:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceToPrefixedUnaryTypeExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = PrefixedUnaryTypeExprType
		symbol.TypeExpression, err = reducer.ToPrefixedUnaryTypeExpr(args[0].Value, args[1].TypeExpression)
	case _ReduceQuestionToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:656:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceExclaimToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:657:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:658:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:659:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceTildeTildeToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:660:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceReturnableTypeExprToTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TypeExprType
		//line grammar.lr:665:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceBinaryTypeExprToTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TypeExprType
		//line grammar.lr:666:4
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
		//line grammar.lr:672:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddToBinaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryTypeOpType
		//line grammar.lr:673:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToBinaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryTypeOpType
		//line grammar.lr:674:4
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
		//line grammar.lr:695:4
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
	case _ReduceUnnamedArgToParameterDecl:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.UnnamedArgToParameterDecl(args[0].TypeExpression)
	case _ReduceUnnamedVarargToParameterDecl:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.UnnamedVarargToParameterDecl(args[0].Value, args[1].TypeExpression)
	case _ReduceUnderscoreArgToParameterDecl:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.UnderscoreArgToParameterDecl(args[0].Value, args[1].TypeExpression)
	case _ReduceUnderscoreVarargToParameterDecl:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDeclType
		symbol.Generic_, err = reducer.UnderscoreVarargToParameterDecl(args[0].Value, args[1].Value, args[2].TypeExpression)
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
		//line grammar.lr:812:4
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
	case _ReduceToFuncTypeExpr:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = FuncTypeExprType
		symbol.TypeExpression, err = reducer.ToFuncTypeExpr(args[0].Value, args[1].Value, args[2].Generic_, args[3].Value, args[4].TypeExpression)
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
	case _ReduceIgnoreInferredRefArgToParameterDef:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.IgnoreInferredRefArgToParameterDef(args[0].Value)
	case _ReduceIgnoreInferredRefVarargToParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.IgnoreInferredRefVarargToParameterDef(args[0].Value, args[1].Value)
	case _ReduceIgnoreArgToParameterDef:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.IgnoreArgToParameterDef(args[0].Value, args[1].TypeExpression)
	case _ReduceIgnoreVarargToParameterDef:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = ParameterDefType
		symbol.Generic_, err = reducer.IgnoreVarargToParameterDef(args[0].Value, args[1].Value, args[2].TypeExpression)
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
		//line grammar.lr:849:4
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
	_GotoState451Action                                             = &_Action{_ShiftAction, _State451, 0}
	_GotoState452Action                                             = &_Action{_ShiftAction, _State452, 0}
	_GotoState453Action                                             = &_Action{_ShiftAction, _State453, 0}
	_GotoState454Action                                             = &_Action{_ShiftAction, _State454, 0}
	_GotoState455Action                                             = &_Action{_ShiftAction, _State455, 0}
	_GotoState456Action                                             = &_Action{_ShiftAction, _State456, 0}
	_GotoState457Action                                             = &_Action{_ShiftAction, _State457, 0}
	_GotoState458Action                                             = &_Action{_ShiftAction, _State458, 0}
	_GotoState459Action                                             = &_Action{_ShiftAction, _State459, 0}
	_GotoState460Action                                             = &_Action{_ShiftAction, _State460, 0}
	_GotoState461Action                                             = &_Action{_ShiftAction, _State461, 0}
	_GotoState462Action                                             = &_Action{_ShiftAction, _State462, 0}
	_GotoState463Action                                             = &_Action{_ShiftAction, _State463, 0}
	_GotoState464Action                                             = &_Action{_ShiftAction, _State464, 0}
	_GotoState465Action                                             = &_Action{_ShiftAction, _State465, 0}
	_GotoState466Action                                             = &_Action{_ShiftAction, _State466, 0}
	_GotoState467Action                                             = &_Action{_ShiftAction, _State467, 0}
	_GotoState468Action                                             = &_Action{_ShiftAction, _State468, 0}
	_GotoState469Action                                             = &_Action{_ShiftAction, _State469, 0}
	_GotoState470Action                                             = &_Action{_ShiftAction, _State470, 0}
	_GotoState471Action                                             = &_Action{_ShiftAction, _State471, 0}
	_GotoState472Action                                             = &_Action{_ShiftAction, _State472, 0}
	_GotoState473Action                                             = &_Action{_ShiftAction, _State473, 0}
	_GotoState474Action                                             = &_Action{_ShiftAction, _State474, 0}
	_GotoState475Action                                             = &_Action{_ShiftAction, _State475, 0}
	_GotoState476Action                                             = &_Action{_ShiftAction, _State476, 0}
	_GotoState477Action                                             = &_Action{_ShiftAction, _State477, 0}
	_GotoState478Action                                             = &_Action{_ShiftAction, _State478, 0}
	_GotoState479Action                                             = &_Action{_ShiftAction, _State479, 0}
	_GotoState480Action                                             = &_Action{_ShiftAction, _State480, 0}
	_GotoState481Action                                             = &_Action{_ShiftAction, _State481, 0}
	_GotoState482Action                                             = &_Action{_ShiftAction, _State482, 0}
	_GotoState483Action                                             = &_Action{_ShiftAction, _State483, 0}
	_ReduceToSourceAction                                           = &_Action{_ReduceAction, 0, _ReduceToSource}
	_ReduceAddToProperDefinitionsAction                             = &_Action{_ReduceAction, 0, _ReduceAddToProperDefinitions}
	_ReduceDefinitionToProperDefinitionsAction                      = &_Action{_ReduceAction, 0, _ReduceDefinitionToProperDefinitions}
	_ReduceProperDefinitionsToDefinitionsAction                     = &_Action{_ReduceAction, 0, _ReduceProperDefinitionsToDefinitions}
	_ReduceImproperToDefinitionsAction                              = &_Action{_ReduceAction, 0, _ReduceImproperToDefinitions}
	_ReduceNilToDefinitionsAction                                   = &_Action{_ReduceAction, 0, _ReduceNilToDefinitions}
	_ReducePackageDefToDefinitionAction                             = &_Action{_ReduceAction, 0, _ReducePackageDefToDefinition}
	_ReduceTypeDefToDefinitionAction                                = &_Action{_ReduceAction, 0, _ReduceTypeDefToDefinition}
	_ReduceNamedFuncDefToDefinitionAction                           = &_Action{_ReduceAction, 0, _ReduceNamedFuncDefToDefinition}
	_ReduceGlobalVarDefToDefinitionAction                           = &_Action{_ReduceAction, 0, _ReduceGlobalVarDefToDefinition}
	_ReduceGlobalVarAssignmentToDefinitionAction                    = &_Action{_ReduceAction, 0, _ReduceGlobalVarAssignmentToDefinition}
	_ReduceStatementBlockToDefinitionAction                         = &_Action{_ReduceAction, 0, _ReduceStatementBlockToDefinition}
	_ReduceCommentGroupsToDefinitionAction                          = &_Action{_ReduceAction, 0, _ReduceCommentGroupsToDefinition}
	_ReduceToStatementBlockAction                                   = &_Action{_ReduceAction, 0, _ReduceToStatementBlock}
	_ReduceAddImplicitToProperStatementsAction                      = &_Action{_ReduceAction, 0, _ReduceAddImplicitToProperStatements}
	_ReduceAddExplicitToProperStatementsAction                      = &_Action{_ReduceAction, 0, _ReduceAddExplicitToProperStatements}
	_ReduceStatementToProperStatementsAction                        = &_Action{_ReduceAction, 0, _ReduceStatementToProperStatements}
	_ReduceProperStatementsToStatementsAction                       = &_Action{_ReduceAction, 0, _ReduceProperStatementsToStatements}
	_ReduceImproperImplicitToStatementsAction                       = &_Action{_ReduceAction, 0, _ReduceImproperImplicitToStatements}
	_ReduceImproperExplicitToStatementsAction                       = &_Action{_ReduceAction, 0, _ReduceImproperExplicitToStatements}
	_ReduceNilToStatementsAction                                    = &_Action{_ReduceAction, 0, _ReduceNilToStatements}
	_ReduceUnsafeStatementToSimpleStatementAction                   = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToSimpleStatement}
	_ReduceExprOrImproperStructStatementToSimpleStatementAction     = &_Action{_ReduceAction, 0, _ReduceExprOrImproperStructStatementToSimpleStatement}
	_ReduceCallbackOpStatementToSimpleStatementAction               = &_Action{_ReduceAction, 0, _ReduceCallbackOpStatementToSimpleStatement}
	_ReduceJumpStatementToSimpleStatementAction                     = &_Action{_ReduceAction, 0, _ReduceJumpStatementToSimpleStatement}
	_ReduceAssignStatementToSimpleStatementAction                   = &_Action{_ReduceAction, 0, _ReduceAssignStatementToSimpleStatement}
	_ReduceUnaryOpAssignStatementToSimpleStatementAction            = &_Action{_ReduceAction, 0, _ReduceUnaryOpAssignStatementToSimpleStatement}
	_ReduceBinaryOpAssignStatementToSimpleStatementAction           = &_Action{_ReduceAction, 0, _ReduceBinaryOpAssignStatementToSimpleStatement}
	_ReduceFallthroughStatementToSimpleStatementAction              = &_Action{_ReduceAction, 0, _ReduceFallthroughStatementToSimpleStatement}
	_ReduceSimpleStatementToStatementAction                         = &_Action{_ReduceAction, 0, _ReduceSimpleStatementToStatement}
	_ReduceImportStatementToStatementAction                         = &_Action{_ReduceAction, 0, _ReduceImportStatementToStatement}
	_ReduceCaseBranchStatementToStatementAction                     = &_Action{_ReduceAction, 0, _ReduceCaseBranchStatementToStatement}
	_ReduceDefaultBranchStatementToStatementAction                  = &_Action{_ReduceAction, 0, _ReduceDefaultBranchStatementToStatement}
	_ReduceSimpleStatementToOptionalSimpleStatementAction           = &_Action{_ReduceAction, 0, _ReduceSimpleStatementToOptionalSimpleStatement}
	_ReduceNilToOptionalSimpleStatementAction                       = &_Action{_ReduceAction, 0, _ReduceNilToOptionalSimpleStatement}
	_ReduceToExprOrImproperStructStatementAction                    = &_Action{_ReduceAction, 0, _ReduceToExprOrImproperStructStatement}
	_ReduceExprToExprsAction                                        = &_Action{_ReduceAction, 0, _ReduceExprToExprs}
	_ReduceAddToExprsAction                                         = &_Action{_ReduceAction, 0, _ReduceAddToExprs}
	_ReduceAsyncToCallbackOpAction                                  = &_Action{_ReduceAction, 0, _ReduceAsyncToCallbackOp}
	_ReduceDeferToCallbackOpAction                                  = &_Action{_ReduceAction, 0, _ReduceDeferToCallbackOp}
	_ReduceToCallbackOpStatementAction                              = &_Action{_ReduceAction, 0, _ReduceToCallbackOpStatement}
	_ReduceToUnsafeStatementAction                                  = &_Action{_ReduceAction, 0, _ReduceToUnsafeStatement}
	_ReduceUnlabeledNoValueToJumpStatementAction                    = &_Action{_ReduceAction, 0, _ReduceUnlabeledNoValueToJumpStatement}
	_ReduceUnlabeledValuedToJumpStatementAction                     = &_Action{_ReduceAction, 0, _ReduceUnlabeledValuedToJumpStatement}
	_ReduceLabeledNoValueToJumpStatementAction                      = &_Action{_ReduceAction, 0, _ReduceLabeledNoValueToJumpStatement}
	_ReduceLabeledValuedToJumpStatementAction                       = &_Action{_ReduceAction, 0, _ReduceLabeledValuedToJumpStatement}
	_ReduceReturnToJumpTypeAction                                   = &_Action{_ReduceAction, 0, _ReduceReturnToJumpType}
	_ReduceBreakToJumpTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceBreakToJumpType}
	_ReduceContinueToJumpTypeAction                                 = &_Action{_ReduceAction, 0, _ReduceContinueToJumpType}
	_ReduceToFallthroughStatementAction                             = &_Action{_ReduceAction, 0, _ReduceToFallthroughStatement}
	_ReduceToAssignStatementAction                                  = &_Action{_ReduceAction, 0, _ReduceToAssignStatement}
	_ReduceToUnaryOpAssignStatementAction                           = &_Action{_ReduceAction, 0, _ReduceToUnaryOpAssignStatement}
	_ReduceAddOneAssignToUnaryOpAssignAction                        = &_Action{_ReduceAction, 0, _ReduceAddOneAssignToUnaryOpAssign}
	_ReduceSubOneAssignToUnaryOpAssignAction                        = &_Action{_ReduceAction, 0, _ReduceSubOneAssignToUnaryOpAssign}
	_ReduceToBinaryOpAssignStatementAction                          = &_Action{_ReduceAction, 0, _ReduceToBinaryOpAssignStatement}
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
	_ReduceSingleToImportStatementAction                            = &_Action{_ReduceAction, 0, _ReduceSingleToImportStatement}
	_ReduceMultipleToImportStatementAction                          = &_Action{_ReduceAction, 0, _ReduceMultipleToImportStatement}
	_ReduceStringLiteralToImportClauseAction                        = &_Action{_ReduceAction, 0, _ReduceStringLiteralToImportClause}
	_ReduceAliasToImportClauseAction                                = &_Action{_ReduceAction, 0, _ReduceAliasToImportClause}
	_ReduceUnusableImportToImportClauseAction                       = &_Action{_ReduceAction, 0, _ReduceUnusableImportToImportClause}
	_ReduceImportToLocalToImportClauseAction                        = &_Action{_ReduceAction, 0, _ReduceImportToLocalToImportClause}
	_ReduceAddImplicitToProperImportClausesAction                   = &_Action{_ReduceAction, 0, _ReduceAddImplicitToProperImportClauses}
	_ReduceAddExplicitToProperImportClausesAction                   = &_Action{_ReduceAction, 0, _ReduceAddExplicitToProperImportClauses}
	_ReduceImportClauseToProperImportClausesAction                  = &_Action{_ReduceAction, 0, _ReduceImportClauseToProperImportClauses}
	_ReduceProperImportClausesToImportClausesAction                 = &_Action{_ReduceAction, 0, _ReduceProperImportClausesToImportClauses}
	_ReduceImplicitToImportClausesAction                            = &_Action{_ReduceAction, 0, _ReduceImplicitToImportClauses}
	_ReduceExplicitToImportClausesAction                            = &_Action{_ReduceAction, 0, _ReduceExplicitToImportClauses}
	_ReduceCasePatternToCasePatternsAction                          = &_Action{_ReduceAction, 0, _ReduceCasePatternToCasePatterns}
	_ReduceMultipleToCasePatternsAction                             = &_Action{_ReduceAction, 0, _ReduceMultipleToCasePatterns}
	_ReduceToVarDeclPatternAction                                   = &_Action{_ReduceAction, 0, _ReduceToVarDeclPattern}
	_ReduceVarToVarOrLetAction                                      = &_Action{_ReduceAction, 0, _ReduceVarToVarOrLet}
	_ReduceLetToVarOrLetAction                                      = &_Action{_ReduceAction, 0, _ReduceLetToVarOrLet}
	_ReduceIdentifierToVarPatternAction                             = &_Action{_ReduceAction, 0, _ReduceIdentifierToVarPattern}
	_ReduceUnderscoreToVarPatternAction                             = &_Action{_ReduceAction, 0, _ReduceUnderscoreToVarPattern}
	_ReduceTuplePatternToVarPatternAction                           = &_Action{_ReduceAction, 0, _ReduceTuplePatternToVarPattern}
	_ReduceToTuplePatternAction                                     = &_Action{_ReduceAction, 0, _ReduceToTuplePattern}
	_ReduceFieldVarPatternToFieldVarPatternsAction                  = &_Action{_ReduceAction, 0, _ReduceFieldVarPatternToFieldVarPatterns}
	_ReduceAddToFieldVarPatternsAction                              = &_Action{_ReduceAction, 0, _ReduceAddToFieldVarPatterns}
	_ReducePositionalToFieldVarPatternAction                        = &_Action{_ReduceAction, 0, _ReducePositionalToFieldVarPattern}
	_ReduceNamedToFieldVarPatternAction                             = &_Action{_ReduceAction, 0, _ReduceNamedToFieldVarPattern}
	_ReduceEllipsisToFieldVarPatternAction                          = &_Action{_ReduceAction, 0, _ReduceEllipsisToFieldVarPattern}
	_ReduceTypeExprToOptionalTypeExprAction                         = &_Action{_ReduceAction, 0, _ReduceTypeExprToOptionalTypeExpr}
	_ReduceNilToOptionalTypeExprAction                              = &_Action{_ReduceAction, 0, _ReduceNilToOptionalTypeExpr}
	_ReduceToAssignPatternAction                                    = &_Action{_ReduceAction, 0, _ReduceToAssignPattern}
	_ReduceAssignPatternToCasePatternAction                         = &_Action{_ReduceAction, 0, _ReduceAssignPatternToCasePattern}
	_ReduceEnumMatchPatternToCasePatternAction                      = &_Action{_ReduceAction, 0, _ReduceEnumMatchPatternToCasePattern}
	_ReduceEnumNondataMatchPattenToCasePatternAction                = &_Action{_ReduceAction, 0, _ReduceEnumNondataMatchPattenToCasePattern}
	_ReduceEnumVarDeclPatternToCasePatternAction                    = &_Action{_ReduceAction, 0, _ReduceEnumVarDeclPatternToCasePattern}
	_ReduceIfExprToExprAction                                       = &_Action{_ReduceAction, 0, _ReduceIfExprToExpr}
	_ReduceSwitchExprToExprAction                                   = &_Action{_ReduceAction, 0, _ReduceSwitchExprToExpr}
	_ReduceLoopExprToExprAction                                     = &_Action{_ReduceAction, 0, _ReduceLoopExprToExpr}
	_ReduceSequenceExprToExprAction                                 = &_Action{_ReduceAction, 0, _ReduceSequenceExprToExpr}
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
	_ReduceBindingToGenericTypeArgumentsAction                      = &_Action{_ReduceAction, 0, _ReduceBindingToGenericTypeArguments}
	_ReduceNilToGenericTypeArgumentsAction                          = &_Action{_ReduceAction, 0, _ReduceNilToGenericTypeArguments}
	_ReduceAddToProperTypeArgumentsAction                           = &_Action{_ReduceAction, 0, _ReduceAddToProperTypeArguments}
	_ReduceTypeExprToProperTypeArgumentsAction                      = &_Action{_ReduceAction, 0, _ReduceTypeExprToProperTypeArguments}
	_ReduceProperTypeArgumentsToTypeArgumentsAction                 = &_Action{_ReduceAction, 0, _ReduceProperTypeArgumentsToTypeArguments}
	_ReduceImproperToTypeArgumentsAction                            = &_Action{_ReduceAction, 0, _ReduceImproperToTypeArguments}
	_ReduceNilToTypeArgumentsAction                                 = &_Action{_ReduceAction, 0, _ReduceNilToTypeArguments}
	_ReduceAddToProperArgumentsAction                               = &_Action{_ReduceAction, 0, _ReduceAddToProperArguments}
	_ReduceArgumentToProperArgumentsAction                          = &_Action{_ReduceAction, 0, _ReduceArgumentToProperArguments}
	_ReduceProperArgumentsToArgumentsAction                         = &_Action{_ReduceAction, 0, _ReduceProperArgumentsToArguments}
	_ReduceImproperToArgumentsAction                                = &_Action{_ReduceAction, 0, _ReduceImproperToArguments}
	_ReduceNilToArgumentsAction                                     = &_Action{_ReduceAction, 0, _ReduceNilToArguments}
	_ReducePositionalToArgumentAction                               = &_Action{_ReduceAction, 0, _ReducePositionalToArgument}
	_ReduceColonExprToArgumentAction                                = &_Action{_ReduceAction, 0, _ReduceColonExprToArgument}
	_ReduceNamedAssignmentToArgumentAction                          = &_Action{_ReduceAction, 0, _ReduceNamedAssignmentToArgument}
	_ReduceVarargAssignmentToArgumentAction                         = &_Action{_ReduceAction, 0, _ReduceVarargAssignmentToArgument}
	_ReduceSkipPatternToArgumentAction                              = &_Action{_ReduceAction, 0, _ReduceSkipPatternToArgument}
	_ReduceUnitUnitPairToColonExprAction                            = &_Action{_ReduceAction, 0, _ReduceUnitUnitPairToColonExpr}
	_ReduceExprUnitPairToColonExprAction                            = &_Action{_ReduceAction, 0, _ReduceExprUnitPairToColonExpr}
	_ReduceUnitExprPairToColonExprAction                            = &_Action{_ReduceAction, 0, _ReduceUnitExprPairToColonExpr}
	_ReduceExprExprPairToColonExprAction                            = &_Action{_ReduceAction, 0, _ReduceExprExprPairToColonExpr}
	_ReduceColonExprUnitTupleToColonExprAction                      = &_Action{_ReduceAction, 0, _ReduceColonExprUnitTupleToColonExpr}
	_ReduceColonExprExprTupleToColonExprAction                      = &_Action{_ReduceAction, 0, _ReduceColonExprExprTupleToColonExpr}
	_ReduceParseErrorExprToAtomExprAction                           = &_Action{_ReduceAction, 0, _ReduceParseErrorExprToAtomExpr}
	_ReduceLiteralExprToAtomExprAction                              = &_Action{_ReduceAction, 0, _ReduceLiteralExprToAtomExpr}
	_ReduceNamedExprToAtomExprAction                                = &_Action{_ReduceAction, 0, _ReduceNamedExprToAtomExpr}
	_ReduceBlockExprToAtomExprAction                                = &_Action{_ReduceAction, 0, _ReduceBlockExprToAtomExpr}
	_ReduceAnonymousFuncExprToAtomExprAction                        = &_Action{_ReduceAction, 0, _ReduceAnonymousFuncExprToAtomExpr}
	_ReduceInitializeExprToAtomExprAction                           = &_Action{_ReduceAction, 0, _ReduceInitializeExprToAtomExpr}
	_ReduceImplicitStructExprToAtomExprAction                       = &_Action{_ReduceAction, 0, _ReduceImplicitStructExprToAtomExpr}
	_ReduceToParseErrorExprAction                                   = &_Action{_ReduceAction, 0, _ReduceToParseErrorExpr}
	_ReduceTrueToLiteralExprAction                                  = &_Action{_ReduceAction, 0, _ReduceTrueToLiteralExpr}
	_ReduceFalseToLiteralExprAction                                 = &_Action{_ReduceAction, 0, _ReduceFalseToLiteralExpr}
	_ReduceIntegerLiteralToLiteralExprAction                        = &_Action{_ReduceAction, 0, _ReduceIntegerLiteralToLiteralExpr}
	_ReduceFloatLiteralToLiteralExprAction                          = &_Action{_ReduceAction, 0, _ReduceFloatLiteralToLiteralExpr}
	_ReduceRuneLiteralToLiteralExprAction                           = &_Action{_ReduceAction, 0, _ReduceRuneLiteralToLiteralExpr}
	_ReduceStringLiteralToLiteralExprAction                         = &_Action{_ReduceAction, 0, _ReduceStringLiteralToLiteralExpr}
	_ReduceIdentifierToNamedExprAction                              = &_Action{_ReduceAction, 0, _ReduceIdentifierToNamedExpr}
	_ReduceUnderscoreToNamedExprAction                              = &_Action{_ReduceAction, 0, _ReduceUnderscoreToNamedExpr}
	_ReduceToBlockExprAction                                        = &_Action{_ReduceAction, 0, _ReduceToBlockExpr}
	_ReduceToInitializeExprAction                                   = &_Action{_ReduceAction, 0, _ReduceToInitializeExpr}
	_ReduceToImplicitStructExprAction                               = &_Action{_ReduceAction, 0, _ReduceToImplicitStructExpr}
	_ReduceAtomExprToAccessibleExprAction                           = &_Action{_ReduceAction, 0, _ReduceAtomExprToAccessibleExpr}
	_ReduceAccessExprToAccessibleExprAction                         = &_Action{_ReduceAction, 0, _ReduceAccessExprToAccessibleExpr}
	_ReduceCallExprToAccessibleExprAction                           = &_Action{_ReduceAction, 0, _ReduceCallExprToAccessibleExpr}
	_ReduceIndexExprToAccessibleExprAction                          = &_Action{_ReduceAction, 0, _ReduceIndexExprToAccessibleExpr}
	_ReduceToAccessExprAction                                       = &_Action{_ReduceAction, 0, _ReduceToAccessExpr}
	_ReduceToIndexExprAction                                        = &_Action{_ReduceAction, 0, _ReduceToIndexExpr}
	_ReduceAccessibleExprToPostfixableExprAction                    = &_Action{_ReduceAction, 0, _ReduceAccessibleExprToPostfixableExpr}
	_ReducePostfixUnaryExprToPostfixableExprAction                  = &_Action{_ReduceAction, 0, _ReducePostfixUnaryExprToPostfixableExpr}
	_ReduceToPostfixUnaryExprAction                                 = &_Action{_ReduceAction, 0, _ReduceToPostfixUnaryExpr}
	_ReducePostfixableExprToPrefixableExprAction                    = &_Action{_ReduceAction, 0, _ReducePostfixableExprToPrefixableExpr}
	_ReducePrefixUnaryExprToPrefixableExprAction                    = &_Action{_ReduceAction, 0, _ReducePrefixUnaryExprToPrefixableExpr}
	_ReduceToPrefixUnaryExprAction                                  = &_Action{_ReduceAction, 0, _ReduceToPrefixUnaryExpr}
	_ReduceNotToPrefixUnaryOpAction                                 = &_Action{_ReduceAction, 0, _ReduceNotToPrefixUnaryOp}
	_ReduceBitNegToPrefixUnaryOpAction                              = &_Action{_ReduceAction, 0, _ReduceBitNegToPrefixUnaryOp}
	_ReduceSubToPrefixUnaryOpAction                                 = &_Action{_ReduceAction, 0, _ReduceSubToPrefixUnaryOp}
	_ReduceMulToPrefixUnaryOpAction                                 = &_Action{_ReduceAction, 0, _ReduceMulToPrefixUnaryOp}
	_ReduceBitAndToPrefixUnaryOpAction                              = &_Action{_ReduceAction, 0, _ReduceBitAndToPrefixUnaryOp}
	_ReducePrefixableExprToMulExprAction                            = &_Action{_ReduceAction, 0, _ReducePrefixableExprToMulExpr}
	_ReduceBinaryMulExprToMulExprAction                             = &_Action{_ReduceAction, 0, _ReduceBinaryMulExprToMulExpr}
	_ReduceToBinaryMulExprAction                                    = &_Action{_ReduceAction, 0, _ReduceToBinaryMulExpr}
	_ReduceMulToMulOpAction                                         = &_Action{_ReduceAction, 0, _ReduceMulToMulOp}
	_ReduceDivToMulOpAction                                         = &_Action{_ReduceAction, 0, _ReduceDivToMulOp}
	_ReduceModToMulOpAction                                         = &_Action{_ReduceAction, 0, _ReduceModToMulOp}
	_ReduceBitAndToMulOpAction                                      = &_Action{_ReduceAction, 0, _ReduceBitAndToMulOp}
	_ReduceBitLshiftToMulOpAction                                   = &_Action{_ReduceAction, 0, _ReduceBitLshiftToMulOp}
	_ReduceBitRshiftToMulOpAction                                   = &_Action{_ReduceAction, 0, _ReduceBitRshiftToMulOp}
	_ReduceMulExprToAddExprAction                                   = &_Action{_ReduceAction, 0, _ReduceMulExprToAddExpr}
	_ReduceBinaryAddExprToAddExprAction                             = &_Action{_ReduceAction, 0, _ReduceBinaryAddExprToAddExpr}
	_ReduceToBinaryAddExprAction                                    = &_Action{_ReduceAction, 0, _ReduceToBinaryAddExpr}
	_ReduceAddToAddOpAction                                         = &_Action{_ReduceAction, 0, _ReduceAddToAddOp}
	_ReduceSubToAddOpAction                                         = &_Action{_ReduceAction, 0, _ReduceSubToAddOp}
	_ReduceBitOrToAddOpAction                                       = &_Action{_ReduceAction, 0, _ReduceBitOrToAddOp}
	_ReduceBitXorToAddOpAction                                      = &_Action{_ReduceAction, 0, _ReduceBitXorToAddOp}
	_ReduceAddExprToCmpExprAction                                   = &_Action{_ReduceAction, 0, _ReduceAddExprToCmpExpr}
	_ReduceBinaryCmpExprToCmpExprAction                             = &_Action{_ReduceAction, 0, _ReduceBinaryCmpExprToCmpExpr}
	_ReduceToBinaryCmpExprAction                                    = &_Action{_ReduceAction, 0, _ReduceToBinaryCmpExpr}
	_ReduceEqualToCmpOpAction                                       = &_Action{_ReduceAction, 0, _ReduceEqualToCmpOp}
	_ReduceNotEqualToCmpOpAction                                    = &_Action{_ReduceAction, 0, _ReduceNotEqualToCmpOp}
	_ReduceLessToCmpOpAction                                        = &_Action{_ReduceAction, 0, _ReduceLessToCmpOp}
	_ReduceLessOrEqualToCmpOpAction                                 = &_Action{_ReduceAction, 0, _ReduceLessOrEqualToCmpOp}
	_ReduceGreaterToCmpOpAction                                     = &_Action{_ReduceAction, 0, _ReduceGreaterToCmpOp}
	_ReduceGreaterOrEqualToCmpOpAction                              = &_Action{_ReduceAction, 0, _ReduceGreaterOrEqualToCmpOp}
	_ReduceCmpExprToAndExprAction                                   = &_Action{_ReduceAction, 0, _ReduceCmpExprToAndExpr}
	_ReduceBinaryAndExprToAndExprAction                             = &_Action{_ReduceAction, 0, _ReduceBinaryAndExprToAndExpr}
	_ReduceToBinaryAndExprAction                                    = &_Action{_ReduceAction, 0, _ReduceToBinaryAndExpr}
	_ReduceAndExprToOrExprAction                                    = &_Action{_ReduceAction, 0, _ReduceAndExprToOrExpr}
	_ReduceBinaryOrExprToOrExprAction                               = &_Action{_ReduceAction, 0, _ReduceBinaryOrExprToOrExpr}
	_ReduceToBinaryOrExprAction                                     = &_Action{_ReduceAction, 0, _ReduceToBinaryOrExpr}
	_ReduceExplicitStructTypeExprToInitializableTypeExprAction      = &_Action{_ReduceAction, 0, _ReduceExplicitStructTypeExprToInitializableTypeExpr}
	_ReduceSliceTypeExprToInitializableTypeExprAction               = &_Action{_ReduceAction, 0, _ReduceSliceTypeExprToInitializableTypeExpr}
	_ReduceArrayTypeExprToInitializableTypeExprAction               = &_Action{_ReduceAction, 0, _ReduceArrayTypeExprToInitializableTypeExpr}
	_ReduceMapTypeExprToInitializableTypeExprAction                 = &_Action{_ReduceAction, 0, _ReduceMapTypeExprToInitializableTypeExpr}
	_ReduceToSliceTypeExprAction                                    = &_Action{_ReduceAction, 0, _ReduceToSliceTypeExpr}
	_ReduceToArrayTypeExprAction                                    = &_Action{_ReduceAction, 0, _ReduceToArrayTypeExpr}
	_ReduceToMapTypeExprAction                                      = &_Action{_ReduceAction, 0, _ReduceToMapTypeExpr}
	_ReduceInitializableTypeExprToAtomTypeExprAction                = &_Action{_ReduceAction, 0, _ReduceInitializableTypeExprToAtomTypeExpr}
	_ReduceNamedTypeExprToAtomTypeExprAction                        = &_Action{_ReduceAction, 0, _ReduceNamedTypeExprToAtomTypeExpr}
	_ReduceInferredTypeExprToAtomTypeExprAction                     = &_Action{_ReduceAction, 0, _ReduceInferredTypeExprToAtomTypeExpr}
	_ReduceImplicitStructTypeExprToAtomTypeExprAction               = &_Action{_ReduceAction, 0, _ReduceImplicitStructTypeExprToAtomTypeExpr}
	_ReduceExplicitEnumTypeExprToAtomTypeExprAction                 = &_Action{_ReduceAction, 0, _ReduceExplicitEnumTypeExprToAtomTypeExpr}
	_ReduceImplicitEnumTypeExprToAtomTypeExprAction                 = &_Action{_ReduceAction, 0, _ReduceImplicitEnumTypeExprToAtomTypeExpr}
	_ReduceTraitTypeExprToAtomTypeExprAction                        = &_Action{_ReduceAction, 0, _ReduceTraitTypeExprToAtomTypeExpr}
	_ReduceFuncTypeExprToAtomTypeExprAction                         = &_Action{_ReduceAction, 0, _ReduceFuncTypeExprToAtomTypeExpr}
	_ReduceParseErrorTypeExprToAtomTypeExprAction                   = &_Action{_ReduceAction, 0, _ReduceParseErrorTypeExprToAtomTypeExpr}
	_ReduceLocalToNamedTypeExprAction                               = &_Action{_ReduceAction, 0, _ReduceLocalToNamedTypeExpr}
	_ReduceExternalToNamedTypeExprAction                            = &_Action{_ReduceAction, 0, _ReduceExternalToNamedTypeExpr}
	_ReduceDotToInferredTypeExprAction                              = &_Action{_ReduceAction, 0, _ReduceDotToInferredTypeExpr}
	_ReduceUnderscoreToInferredTypeExprAction                       = &_Action{_ReduceAction, 0, _ReduceUnderscoreToInferredTypeExpr}
	_ReduceToParseErrorTypeExprAction                               = &_Action{_ReduceAction, 0, _ReduceToParseErrorTypeExpr}
	_ReduceAtomTypeExprToReturnableTypeExprAction                   = &_Action{_ReduceAction, 0, _ReduceAtomTypeExprToReturnableTypeExpr}
	_ReducePrefixedUnaryTypeExprToReturnableTypeExprAction          = &_Action{_ReduceAction, 0, _ReducePrefixedUnaryTypeExprToReturnableTypeExpr}
	_ReduceToPrefixedUnaryTypeExprAction                            = &_Action{_ReduceAction, 0, _ReduceToPrefixedUnaryTypeExpr}
	_ReduceQuestionToPrefixUnaryTypeOpAction                        = &_Action{_ReduceAction, 0, _ReduceQuestionToPrefixUnaryTypeOp}
	_ReduceExclaimToPrefixUnaryTypeOpAction                         = &_Action{_ReduceAction, 0, _ReduceExclaimToPrefixUnaryTypeOp}
	_ReduceBitAndToPrefixUnaryTypeOpAction                          = &_Action{_ReduceAction, 0, _ReduceBitAndToPrefixUnaryTypeOp}
	_ReduceBitNegToPrefixUnaryTypeOpAction                          = &_Action{_ReduceAction, 0, _ReduceBitNegToPrefixUnaryTypeOp}
	_ReduceTildeTildeToPrefixUnaryTypeOpAction                      = &_Action{_ReduceAction, 0, _ReduceTildeTildeToPrefixUnaryTypeOp}
	_ReduceReturnableTypeExprToTypeExprAction                       = &_Action{_ReduceAction, 0, _ReduceReturnableTypeExprToTypeExpr}
	_ReduceBinaryTypeExprToTypeExprAction                           = &_Action{_ReduceAction, 0, _ReduceBinaryTypeExprToTypeExpr}
	_ReduceToBinaryTypeExprAction                                   = &_Action{_ReduceAction, 0, _ReduceToBinaryTypeExpr}
	_ReduceMulToBinaryTypeOpAction                                  = &_Action{_ReduceAction, 0, _ReduceMulToBinaryTypeOp}
	_ReduceAddToBinaryTypeOpAction                                  = &_Action{_ReduceAction, 0, _ReduceAddToBinaryTypeOp}
	_ReduceSubToBinaryTypeOpAction                                  = &_Action{_ReduceAction, 0, _ReduceSubToBinaryTypeOp}
	_ReduceDefinitionToTypeDefAction                                = &_Action{_ReduceAction, 0, _ReduceDefinitionToTypeDef}
	_ReduceConstrainedDefToTypeDefAction                            = &_Action{_ReduceAction, 0, _ReduceConstrainedDefToTypeDef}
	_ReduceAliasToTypeDefAction                                     = &_Action{_ReduceAction, 0, _ReduceAliasToTypeDef}
	_ReduceUnconstrainedToGenericParameterDefAction                 = &_Action{_ReduceAction, 0, _ReduceUnconstrainedToGenericParameterDef}
	_ReduceConstrainedToGenericParameterDefAction                   = &_Action{_ReduceAction, 0, _ReduceConstrainedToGenericParameterDef}
	_ReduceAddToProperGenericParameterDefsAction                    = &_Action{_ReduceAction, 0, _ReduceAddToProperGenericParameterDefs}
	_ReduceGenericParameterDefToProperGenericParameterDefsAction    = &_Action{_ReduceAction, 0, _ReduceGenericParameterDefToProperGenericParameterDefs}
	_ReduceProperGenericParameterDefsToGenericParameterDefsAction   = &_Action{_ReduceAction, 0, _ReduceProperGenericParameterDefsToGenericParameterDefs}
	_ReduceImproperToGenericParameterDefsAction                     = &_Action{_ReduceAction, 0, _ReduceImproperToGenericParameterDefs}
	_ReduceNilToGenericParameterDefsAction                          = &_Action{_ReduceAction, 0, _ReduceNilToGenericParameterDefs}
	_ReduceGenericToOptionalGenericParametersAction                 = &_Action{_ReduceAction, 0, _ReduceGenericToOptionalGenericParameters}
	_ReduceNilToOptionalGenericParametersAction                     = &_Action{_ReduceAction, 0, _ReduceNilToOptionalGenericParameters}
	_ReduceNamedToFieldDefAction                                    = &_Action{_ReduceAction, 0, _ReduceNamedToFieldDef}
	_ReduceUnnamedToFieldDefAction                                  = &_Action{_ReduceAction, 0, _ReduceUnnamedToFieldDef}
	_ReduceStructPaddingToFieldDefAction                            = &_Action{_ReduceAction, 0, _ReduceStructPaddingToFieldDef}
	_ReduceMethodSignatureToFieldDefAction                          = &_Action{_ReduceAction, 0, _ReduceMethodSignatureToFieldDef}
	_ReduceUnsafeStatementToFieldDefAction                          = &_Action{_ReduceAction, 0, _ReduceUnsafeStatementToFieldDef}
	_ReduceDefaultToOptionalDefaultAction                           = &_Action{_ReduceAction, 0, _ReduceDefaultToOptionalDefault}
	_ReduceNilToOptionalDefaultAction                               = &_Action{_ReduceAction, 0, _ReduceNilToOptionalDefault}
	_ReduceAddToProperImplicitFieldDefsAction                       = &_Action{_ReduceAction, 0, _ReduceAddToProperImplicitFieldDefs}
	_ReduceFieldDefToProperImplicitFieldDefsAction                  = &_Action{_ReduceAction, 0, _ReduceFieldDefToProperImplicitFieldDefs}
	_ReduceProperImplicitFieldDefsToImplicitFieldDefsAction         = &_Action{_ReduceAction, 0, _ReduceProperImplicitFieldDefsToImplicitFieldDefs}
	_ReduceImproperToImplicitFieldDefsAction                        = &_Action{_ReduceAction, 0, _ReduceImproperToImplicitFieldDefs}
	_ReduceNilToImplicitFieldDefsAction                             = &_Action{_ReduceAction, 0, _ReduceNilToImplicitFieldDefs}
	_ReduceToImplicitStructTypeExprAction                           = &_Action{_ReduceAction, 0, _ReduceToImplicitStructTypeExpr}
	_ReduceAddImplicitToProperExplicitFieldDefsAction               = &_Action{_ReduceAction, 0, _ReduceAddImplicitToProperExplicitFieldDefs}
	_ReduceAddExplicitToProperExplicitFieldDefsAction               = &_Action{_ReduceAction, 0, _ReduceAddExplicitToProperExplicitFieldDefs}
	_ReduceFieldDefToProperExplicitFieldDefsAction                  = &_Action{_ReduceAction, 0, _ReduceFieldDefToProperExplicitFieldDefs}
	_ReduceProperExplicitFieldDefsToExplicitFieldDefsAction         = &_Action{_ReduceAction, 0, _ReduceProperExplicitFieldDefsToExplicitFieldDefs}
	_ReduceImproperImplicitToExplicitFieldDefsAction                = &_Action{_ReduceAction, 0, _ReduceImproperImplicitToExplicitFieldDefs}
	_ReduceImproperExplicitToExplicitFieldDefsAction                = &_Action{_ReduceAction, 0, _ReduceImproperExplicitToExplicitFieldDefs}
	_ReduceNilToExplicitFieldDefsAction                             = &_Action{_ReduceAction, 0, _ReduceNilToExplicitFieldDefs}
	_ReduceToExplicitStructTypeExprAction                           = &_Action{_ReduceAction, 0, _ReduceToExplicitStructTypeExpr}
	_ReduceToTraitTypeExprAction                                    = &_Action{_ReduceAction, 0, _ReduceToTraitTypeExpr}
	_ReducePairToProperImplicitEnumValueDefsAction                  = &_Action{_ReduceAction, 0, _ReducePairToProperImplicitEnumValueDefs}
	_ReduceAddToProperImplicitEnumValueDefsAction                   = &_Action{_ReduceAction, 0, _ReduceAddToProperImplicitEnumValueDefs}
	_ReduceProperImplicitEnumValueDefsToImplicitEnumValueDefsAction = &_Action{_ReduceAction, 0, _ReduceProperImplicitEnumValueDefsToImplicitEnumValueDefs}
	_ReduceImproperToImplicitEnumValueDefsAction                    = &_Action{_ReduceAction, 0, _ReduceImproperToImplicitEnumValueDefs}
	_ReduceToImplicitEnumTypeExprAction                             = &_Action{_ReduceAction, 0, _ReduceToImplicitEnumTypeExpr}
	_ReduceExplicitPairToProperExplicitEnumValueDefsAction          = &_Action{_ReduceAction, 0, _ReduceExplicitPairToProperExplicitEnumValueDefs}
	_ReduceImplicitPairToProperExplicitEnumValueDefsAction          = &_Action{_ReduceAction, 0, _ReduceImplicitPairToProperExplicitEnumValueDefs}
	_ReduceExplicitAddToProperExplicitEnumValueDefsAction           = &_Action{_ReduceAction, 0, _ReduceExplicitAddToProperExplicitEnumValueDefs}
	_ReduceImplicitAddToProperExplicitEnumValueDefsAction           = &_Action{_ReduceAction, 0, _ReduceImplicitAddToProperExplicitEnumValueDefs}
	_ReduceProperExplicitEnumValueDefsToExplicitEnumValueDefsAction = &_Action{_ReduceAction, 0, _ReduceProperExplicitEnumValueDefsToExplicitEnumValueDefs}
	_ReduceImproperToExplicitEnumValueDefsAction                    = &_Action{_ReduceAction, 0, _ReduceImproperToExplicitEnumValueDefs}
	_ReduceToExplicitEnumTypeExprAction                             = &_Action{_ReduceAction, 0, _ReduceToExplicitEnumTypeExpr}
	_ReduceReturnableTypeExprToReturnTypeAction                     = &_Action{_ReduceAction, 0, _ReduceReturnableTypeExprToReturnType}
	_ReduceNilToReturnTypeAction                                    = &_Action{_ReduceAction, 0, _ReduceNilToReturnType}
	_ReduceUnnamedArgToParameterDeclAction                          = &_Action{_ReduceAction, 0, _ReduceUnnamedArgToParameterDecl}
	_ReduceUnnamedVarargToParameterDeclAction                       = &_Action{_ReduceAction, 0, _ReduceUnnamedVarargToParameterDecl}
	_ReduceUnderscoreArgToParameterDeclAction                       = &_Action{_ReduceAction, 0, _ReduceUnderscoreArgToParameterDecl}
	_ReduceUnderscoreVarargToParameterDeclAction                    = &_Action{_ReduceAction, 0, _ReduceUnderscoreVarargToParameterDecl}
	_ReduceArgToParameterDeclAction                                 = &_Action{_ReduceAction, 0, _ReduceArgToParameterDecl}
	_ReduceVarargToParameterDeclAction                              = &_Action{_ReduceAction, 0, _ReduceVarargToParameterDecl}
	_ReduceAddToProperParameterDeclsAction                          = &_Action{_ReduceAction, 0, _ReduceAddToProperParameterDecls}
	_ReduceParameterDeclToProperParameterDeclsAction                = &_Action{_ReduceAction, 0, _ReduceParameterDeclToProperParameterDecls}
	_ReduceProperParameterDeclsToParameterDeclsAction               = &_Action{_ReduceAction, 0, _ReduceProperParameterDeclsToParameterDecls}
	_ReduceImproperToParameterDeclsAction                           = &_Action{_ReduceAction, 0, _ReduceImproperToParameterDecls}
	_ReduceNilToParameterDeclsAction                                = &_Action{_ReduceAction, 0, _ReduceNilToParameterDecls}
	_ReduceToFuncTypeExprAction                                     = &_Action{_ReduceAction, 0, _ReduceToFuncTypeExpr}
	_ReduceToMethodSignatureAction                                  = &_Action{_ReduceAction, 0, _ReduceToMethodSignature}
	_ReduceInferredRefArgToParameterDefAction                       = &_Action{_ReduceAction, 0, _ReduceInferredRefArgToParameterDef}
	_ReduceInferredRefVarargToParameterDefAction                    = &_Action{_ReduceAction, 0, _ReduceInferredRefVarargToParameterDef}
	_ReduceArgToParameterDefAction                                  = &_Action{_ReduceAction, 0, _ReduceArgToParameterDef}
	_ReduceVarargToParameterDefAction                               = &_Action{_ReduceAction, 0, _ReduceVarargToParameterDef}
	_ReduceIgnoreInferredRefArgToParameterDefAction                 = &_Action{_ReduceAction, 0, _ReduceIgnoreInferredRefArgToParameterDef}
	_ReduceIgnoreInferredRefVarargToParameterDefAction              = &_Action{_ReduceAction, 0, _ReduceIgnoreInferredRefVarargToParameterDef}
	_ReduceIgnoreArgToParameterDefAction                            = &_Action{_ReduceAction, 0, _ReduceIgnoreArgToParameterDef}
	_ReduceIgnoreVarargToParameterDefAction                         = &_Action{_ReduceAction, 0, _ReduceIgnoreVarargToParameterDef}
	_ReduceAddToProperParameterDefsAction                           = &_Action{_ReduceAction, 0, _ReduceAddToProperParameterDefs}
	_ReduceParameterDefToProperParameterDefsAction                  = &_Action{_ReduceAction, 0, _ReduceParameterDefToProperParameterDefs}
	_ReduceProperParameterDefsToParameterDefsAction                 = &_Action{_ReduceAction, 0, _ReduceProperParameterDefsToParameterDefs}
	_ReduceImproperToParameterDefsAction                            = &_Action{_ReduceAction, 0, _ReduceImproperToParameterDefs}
	_ReduceNilToParameterDefsAction                                 = &_Action{_ReduceAction, 0, _ReduceNilToParameterDefs}
	_ReduceFuncDefToNamedFuncDefAction                              = &_Action{_ReduceAction, 0, _ReduceFuncDefToNamedFuncDef}
	_ReduceMethodDefToNamedFuncDefAction                            = &_Action{_ReduceAction, 0, _ReduceMethodDefToNamedFuncDef}
	_ReduceAliasToNamedFuncDefAction                                = &_Action{_ReduceAction, 0, _ReduceAliasToNamedFuncDef}
	_ReduceToAnonymousFuncExprAction                                = &_Action{_ReduceAction, 0, _ReduceToAnonymousFuncExpr}
	_ReduceNoSpecToPackageDefAction                                 = &_Action{_ReduceAction, 0, _ReduceNoSpecToPackageDef}
	_ReduceWithSpecToPackageDefAction                               = &_Action{_ReduceAction, 0, _ReduceWithSpecToPackageDef}
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
	{_State4, PrefixedUnaryTypeExprType}:           _GotoState131Action,
	{_State4, PrefixUnaryTypeOpType}:               _GotoState130Action,
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
	{_State42, PrefixedUnaryTypeExprType}:          _GotoState131Action,
	{_State42, PrefixUnaryTypeOpType}:              _GotoState130Action,
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
	{_State114, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State114, PrefixUnaryTypeOpType}:             _GotoState130Action,
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
	{_State130, IdentifierToken}:                   _GotoState113Action,
	{_State130, UnderscoreToken}:                   _GotoState119Action,
	{_State130, StructToken}:                       _GotoState50Action,
	{_State130, EnumToken}:                         _GotoState110Action,
	{_State130, TraitToken}:                        _GotoState118Action,
	{_State130, FuncToken}:                         _GotoState112Action,
	{_State130, LparenToken}:                       _GotoState114Action,
	{_State130, LbracketToken}:                     _GotoState42Action,
	{_State130, DotToken}:                          _GotoState109Action,
	{_State130, QuestionToken}:                     _GotoState116Action,
	{_State130, ExclaimToken}:                      _GotoState111Action,
	{_State130, TildeTildeToken}:                   _GotoState117Action,
	{_State130, BitNegToken}:                       _GotoState108Action,
	{_State130, BitAndToken}:                       _GotoState107Action,
	{_State130, ParseErrorToken}:                   _GotoState115Action,
	{_State130, InitializableTypeExprType}:         _GotoState127Action,
	{_State130, SliceTypeExprType}:                 _GotoState101Action,
	{_State130, ArrayTypeExprType}:                 _GotoState60Action,
	{_State130, MapTypeExprType}:                   _GotoState88Action,
	{_State130, AtomTypeExprType}:                  _GotoState120Action,
	{_State130, NamedTypeExprType}:                 _GotoState128Action,
	{_State130, InferredTypeExprType}:              _GotoState126Action,
	{_State130, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State130, ReturnableTypeExprType}:            _GotoState248Action,
	{_State130, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State130, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State130, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State130, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State130, TraitTypeExprType}:                 _GotoState133Action,
	{_State130, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State130, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State130, FuncTypeExprType}:                  _GotoState123Action,
	{_State134, DollarLbracketToken}:               _GotoState254Action,
	{_State134, AssignToken}:                       _GotoState253Action,
	{_State134, OptionalGenericParametersType}:     _GotoState255Action,
	{_State135, IdentifierToken}:                   _GotoState256Action,
	{_State135, UnderscoreToken}:                   _GotoState257Action,
	{_State135, ParameterDefType}:                  _GotoState258Action,
	{_State136, NewlinesToken}:                     _GotoState259Action,
	{_State136, SemicolonToken}:                    _GotoState260Action,
	{_State138, RbraceToken}:                       _GotoState261Action,
	{_State140, DollarLbracketToken}:               _GotoState254Action,
	{_State140, AssignToken}:                       _GotoState262Action,
	{_State140, OptionalGenericParametersType}:     _GotoState263Action,
	{_State141, CommentGroupsToken}:                _GotoState9Action,
	{_State141, PackageToken}:                      _GotoState13Action,
	{_State141, TypeToken}:                         _GotoState14Action,
	{_State141, FuncToken}:                         _GotoState10Action,
	{_State141, VarToken}:                          _GotoState15Action,
	{_State141, LetToken}:                          _GotoState12Action,
	{_State141, LbraceToken}:                       _GotoState11Action,
	{_State141, DefinitionType}:                    _GotoState264Action,
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
	{_State142, ExprType}:                          _GotoState265Action,
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
	{_State144, IdentifierToken}:                   _GotoState267Action,
	{_State144, UnderscoreToken}:                   _GotoState145Action,
	{_State144, LparenToken}:                       _GotoState144Action,
	{_State144, EllipsisToken}:                     _GotoState266Action,
	{_State144, VarPatternType}:                    _GotoState270Action,
	{_State144, TuplePatternType}:                  _GotoState146Action,
	{_State144, FieldVarPatternsType}:              _GotoState269Action,
	{_State144, FieldVarPatternType}:               _GotoState268Action,
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
	{_State147, OptionalTypeExprType}:              _GotoState271Action,
	{_State147, InitializableTypeExprType}:         _GotoState127Action,
	{_State147, SliceTypeExprType}:                 _GotoState101Action,
	{_State147, ArrayTypeExprType}:                 _GotoState60Action,
	{_State147, MapTypeExprType}:                   _GotoState88Action,
	{_State147, AtomTypeExprType}:                  _GotoState120Action,
	{_State147, NamedTypeExprType}:                 _GotoState128Action,
	{_State147, InferredTypeExprType}:              _GotoState126Action,
	{_State147, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State147, ReturnableTypeExprType}:            _GotoState132Action,
	{_State147, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State147, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State147, TypeExprType}:                      _GotoState272Action,
	{_State147, BinaryTypeExprType}:                _GotoState121Action,
	{_State147, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State147, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State147, TraitTypeExprType}:                 _GotoState133Action,
	{_State147, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State147, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State147, FuncTypeExprType}:                  _GotoState123Action,
	{_State148, IdentifierToken}:                   _GotoState273Action,
	{_State149, DotToken}:                          _GotoState274Action,
	{_State152, CommaToken}:                        _GotoState276Action,
	{_State152, ColonToken}:                        _GotoState275Action,
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
	{_State155, SimpleStatementType}:               _GotoState278Action,
	{_State155, OptionalSimpleStatementType}:       _GotoState277Action,
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
	{_State156, ParameterDefType}:                  _GotoState279Action,
	{_State156, ProperParameterDefsType}:           _GotoState281Action,
	{_State156, ParameterDefsType}:                 _GotoState280Action,
	{_State158, StringLiteralToken}:                _GotoState282Action,
	{_State159, StringLiteralToken}:                _GotoState283Action,
	{_State160, StringLiteralToken}:                _GotoState161Action,
	{_State160, IdentifierToken}:                   _GotoState159Action,
	{_State160, UnderscoreToken}:                   _GotoState162Action,
	{_State160, DotToken}:                          _GotoState158Action,
	{_State160, ImportClauseType}:                  _GotoState284Action,
	{_State160, ProperImportClausesType}:           _GotoState286Action,
	{_State160, ImportClausesType}:                 _GotoState285Action,
	{_State162, StringLiteralToken}:                _GotoState287Action,
	{_State164, RbracketToken}:                     _GotoState290Action,
	{_State164, CommaToken}:                        _GotoState289Action,
	{_State164, ColonToken}:                        _GotoState288Action,
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
	{_State165, ExprType}:                          _GotoState291Action,
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
	{_State167, AssignToken}:                       _GotoState292Action,
	{_State169, RparenToken}:                       _GotoState293Action,
	{_State170, ColonToken}:                        _GotoState294Action,
	{_State171, ColonToken}:                        _GotoState295Action,
	{_State171, EllipsisToken}:                     _GotoState296Action,
	{_State172, CommaToken}:                        _GotoState297Action,
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
	{_State173, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State173, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State173, TypeExprType}:                      _GotoState245Action,
	{_State173, BinaryTypeExprType}:                _GotoState121Action,
	{_State173, FieldDefType}:                      _GotoState299Action,
	{_State173, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State173, ProperExplicitFieldDefsType}:       _GotoState300Action,
	{_State173, ExplicitFieldDefsType}:             _GotoState298Action,
	{_State173, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State173, TraitTypeExprType}:                 _GotoState133Action,
	{_State173, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State173, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State173, FuncTypeExprType}:                  _GotoState123Action,
	{_State173, MethodSignatureType}:               _GotoState242Action,
	{_State174, IdentifierToken}:                   _GotoState301Action,
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
	{_State184, ProperTypeArgumentsType}:           _GotoState302Action,
	{_State184, TypeArgumentsType}:                 _GotoState303Action,
	{_State184, InitializableTypeExprType}:         _GotoState127Action,
	{_State184, SliceTypeExprType}:                 _GotoState101Action,
	{_State184, ArrayTypeExprType}:                 _GotoState60Action,
	{_State184, MapTypeExprType}:                   _GotoState88Action,
	{_State184, AtomTypeExprType}:                  _GotoState120Action,
	{_State184, NamedTypeExprType}:                 _GotoState128Action,
	{_State184, InferredTypeExprType}:              _GotoState126Action,
	{_State184, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State184, ReturnableTypeExprType}:            _GotoState132Action,
	{_State184, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State184, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State184, TypeExprType}:                      _GotoState304Action,
	{_State184, BinaryTypeExprType}:                _GotoState121Action,
	{_State184, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State184, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State184, TraitTypeExprType}:                 _GotoState133Action,
	{_State184, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State184, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State184, FuncTypeExprType}:                  _GotoState123Action,
	{_State185, IdentifierToken}:                   _GotoState305Action,
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
	{_State186, ArgumentType}:                      _GotoState306Action,
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
	{_State192, ExprType}:                          _GotoState307Action,
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
	{_State193, LparenToken}:                       _GotoState308Action,
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
	{_State199, MulExprType}:                       _GotoState309Action,
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
	{_State200, CmpExprType}:                       _GotoState310Action,
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
	{_State201, ExprType}:                          _GotoState311Action,
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
	{_State210, AddExprType}:                       _GotoState312Action,
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
	{_State211, ExprType}:                          _GotoState313Action,
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
	{_State212, ArgumentsType}:                     _GotoState314Action,
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
	{_State213, ExprsType}:                         _GotoState315Action,
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
	{_State221, PrefixableExprType}:                _GotoState316Action,
	{_State221, PrefixUnaryExprType}:               _GotoState96Action,
	{_State221, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State221, InitializableTypeExprType}:         _GotoState83Action,
	{_State221, SliceTypeExprType}:                 _GotoState101Action,
	{_State221, ArrayTypeExprType}:                 _GotoState60Action,
	{_State221, MapTypeExprType}:                   _GotoState88Action,
	{_State221, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State221, AnonymousFuncExprType}:             _GotoState59Action,
	{_State222, LbraceToken}:                       _GotoState11Action,
	{_State222, StatementBlockType}:                _GotoState317Action,
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
	{_State223, AssignPatternType}:                 _GotoState318Action,
	{_State223, OptionalLabelDeclType}:             _GotoState153Action,
	{_State223, SequenceExprType}:                  _GotoState320Action,
	{_State223, ForAssignmentType}:                 _GotoState319Action,
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
	{_State224, CaseToken}:                         _GotoState321Action,
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
	{_State224, SequenceExprType}:                  _GotoState323Action,
	{_State224, ConditionType}:                     _GotoState322Action,
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
	{_State225, SequenceExprType}:                  _GotoState324Action,
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
	{_State230, AndExprType}:                       _GotoState325Action,
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
	{_State232, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State232, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State232, TypeExprType}:                      _GotoState245Action,
	{_State232, BinaryTypeExprType}:                _GotoState121Action,
	{_State232, FieldDefType}:                      _GotoState327Action,
	{_State232, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State232, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State232, TraitTypeExprType}:                 _GotoState133Action,
	{_State232, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State232, ProperExplicitEnumValueDefsType}:   _GotoState328Action,
	{_State232, ExplicitEnumValueDefsType}:         _GotoState326Action,
	{_State232, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State232, FuncTypeExprType}:                  _GotoState123Action,
	{_State232, MethodSignatureType}:               _GotoState242Action,
	{_State233, IdentifierToken}:                   _GotoState330Action,
	{_State233, UnderscoreToken}:                   _GotoState331Action,
	{_State233, StructToken}:                       _GotoState50Action,
	{_State233, EnumToken}:                         _GotoState110Action,
	{_State233, TraitToken}:                        _GotoState118Action,
	{_State233, FuncToken}:                         _GotoState112Action,
	{_State233, LparenToken}:                       _GotoState114Action,
	{_State233, LbracketToken}:                     _GotoState42Action,
	{_State233, DotToken}:                          _GotoState109Action,
	{_State233, QuestionToken}:                     _GotoState116Action,
	{_State233, ExclaimToken}:                      _GotoState111Action,
	{_State233, EllipsisToken}:                     _GotoState329Action,
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
	{_State233, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State233, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State233, TypeExprType}:                      _GotoState335Action,
	{_State233, BinaryTypeExprType}:                _GotoState121Action,
	{_State233, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State233, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State233, TraitTypeExprType}:                 _GotoState133Action,
	{_State233, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State233, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State233, ParameterDeclType}:                 _GotoState332Action,
	{_State233, ProperParameterDeclsType}:          _GotoState334Action,
	{_State233, ParameterDeclsType}:                _GotoState333Action,
	{_State233, FuncTypeExprType}:                  _GotoState123Action,
	{_State234, IdentifierToken}:                   _GotoState336Action,
	{_State236, IdentifierToken}:                   _GotoState337Action,
	{_State236, LparenToken}:                       _GotoState233Action,
	{_State237, IdentifierToken}:                   _GotoState113Action,
	{_State237, UnderscoreToken}:                   _GotoState119Action,
	{_State237, StructToken}:                       _GotoState50Action,
	{_State237, EnumToken}:                         _GotoState110Action,
	{_State237, TraitToken}:                        _GotoState118Action,
	{_State237, FuncToken}:                         _GotoState112Action,
	{_State237, LparenToken}:                       _GotoState114Action,
	{_State237, LbracketToken}:                     _GotoState42Action,
	{_State237, DotToken}:                          _GotoState338Action,
	{_State237, QuestionToken}:                     _GotoState116Action,
	{_State237, ExclaimToken}:                      _GotoState111Action,
	{_State237, DollarLbracketToken}:               _GotoState184Action,
	{_State237, TildeTildeToken}:                   _GotoState117Action,
	{_State237, BitNegToken}:                       _GotoState108Action,
	{_State237, BitAndToken}:                       _GotoState107Action,
	{_State237, ParseErrorToken}:                   _GotoState115Action,
	{_State237, GenericTypeArgumentsType}:          _GotoState235Action,
	{_State237, InitializableTypeExprType}:         _GotoState127Action,
	{_State237, SliceTypeExprType}:                 _GotoState101Action,
	{_State237, ArrayTypeExprType}:                 _GotoState60Action,
	{_State237, MapTypeExprType}:                   _GotoState88Action,
	{_State237, AtomTypeExprType}:                  _GotoState120Action,
	{_State237, NamedTypeExprType}:                 _GotoState128Action,
	{_State237, InferredTypeExprType}:              _GotoState126Action,
	{_State237, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State237, ReturnableTypeExprType}:            _GotoState132Action,
	{_State237, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State237, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State237, TypeExprType}:                      _GotoState339Action,
	{_State237, BinaryTypeExprType}:                _GotoState121Action,
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
	{_State238, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State238, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State238, TypeExprType}:                      _GotoState340Action,
	{_State238, BinaryTypeExprType}:                _GotoState121Action,
	{_State238, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State238, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State238, TraitTypeExprType}:                 _GotoState133Action,
	{_State238, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State238, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State238, FuncTypeExprType}:                  _GotoState123Action,
	{_State239, OrToken}:                           _GotoState341Action,
	{_State240, RparenToken}:                       _GotoState342Action,
	{_State241, RparenToken}:                       _GotoState343Action,
	{_State243, NewlinesToken}:                     _GotoState344Action,
	{_State243, OrToken}:                           _GotoState345Action,
	{_State244, CommaToken}:                        _GotoState346Action,
	{_State245, AssignToken}:                       _GotoState347Action,
	{_State245, AddToken}:                          _GotoState249Action,
	{_State245, SubToken}:                          _GotoState251Action,
	{_State245, MulToken}:                          _GotoState250Action,
	{_State245, BinaryTypeOpType}:                  _GotoState252Action,
	{_State245, OptionalDefaultType}:               _GotoState348Action,
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
	{_State247, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State247, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State247, TypeExprType}:                      _GotoState245Action,
	{_State247, BinaryTypeExprType}:                _GotoState121Action,
	{_State247, FieldDefType}:                      _GotoState299Action,
	{_State247, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State247, ProperExplicitFieldDefsType}:       _GotoState300Action,
	{_State247, ExplicitFieldDefsType}:             _GotoState349Action,
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
	{_State252, ReturnableTypeExprType}:            _GotoState350Action,
	{_State252, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State252, PrefixUnaryTypeOpType}:             _GotoState130Action,
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
	{_State253, ExprType}:                          _GotoState351Action,
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
	{_State254, IdentifierToken}:                   _GotoState352Action,
	{_State254, GenericParameterDefType}:           _GotoState353Action,
	{_State254, ProperGenericParameterDefsType}:    _GotoState355Action,
	{_State254, GenericParameterDefsType}:          _GotoState354Action,
	{_State255, LparenToken}:                       _GotoState356Action,
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
	{_State256, EllipsisToken}:                     _GotoState357Action,
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
	{_State256, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State256, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State256, TypeExprType}:                      _GotoState358Action,
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
	{_State257, EllipsisToken}:                     _GotoState359Action,
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
	{_State257, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State257, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State257, TypeExprType}:                      _GotoState360Action,
	{_State257, BinaryTypeExprType}:                _GotoState121Action,
	{_State257, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State257, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State257, TraitTypeExprType}:                 _GotoState133Action,
	{_State257, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State257, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State257, FuncTypeExprType}:                  _GotoState123Action,
	{_State258, RparenToken}:                       _GotoState361Action,
	{_State259, IntegerLiteralToken}:               _GotoState40Action,
	{_State259, FloatLiteralToken}:                 _GotoState35Action,
	{_State259, RuneLiteralToken}:                  _GotoState48Action,
	{_State259, StringLiteralToken}:                _GotoState49Action,
	{_State259, IdentifierToken}:                   _GotoState38Action,
	{_State259, UnderscoreToken}:                   _GotoState53Action,
	{_State259, TrueToken}:                         _GotoState52Action,
	{_State259, FalseToken}:                        _GotoState34Action,
	{_State259, CaseToken}:                         _GotoState29Action,
	{_State259, DefaultToken}:                      _GotoState31Action,
	{_State259, ReturnToken}:                       _GotoState47Action,
	{_State259, BreakToken}:                        _GotoState28Action,
	{_State259, ContinueToken}:                     _GotoState30Action,
	{_State259, FallthroughToken}:                  _GotoState33Action,
	{_State259, ImportToken}:                       _GotoState39Action,
	{_State259, UnsafeToken}:                       _GotoState54Action,
	{_State259, StructToken}:                       _GotoState50Action,
	{_State259, FuncToken}:                         _GotoState36Action,
	{_State259, AsyncToken}:                        _GotoState25Action,
	{_State259, DeferToken}:                        _GotoState32Action,
	{_State259, VarToken}:                          _GotoState15Action,
	{_State259, LetToken}:                          _GotoState12Action,
	{_State259, NotToken}:                          _GotoState45Action,
	{_State259, LabelDeclToken}:                    _GotoState41Action,
	{_State259, LparenToken}:                       _GotoState43Action,
	{_State259, LbracketToken}:                     _GotoState42Action,
	{_State259, SubToken}:                          _GotoState51Action,
	{_State259, MulToken}:                          _GotoState44Action,
	{_State259, BitNegToken}:                       _GotoState27Action,
	{_State259, BitAndToken}:                       _GotoState26Action,
	{_State259, GreaterToken}:                      _GotoState37Action,
	{_State259, ParseErrorToken}:                   _GotoState46Action,
	{_State259, SimpleStatementType}:               _GotoState100Action,
	{_State259, StatementType}:                     _GotoState362Action,
	{_State259, ExprOrImproperStructStatementType}: _GotoState77Action,
	{_State259, ExprsType}:                         _GotoState78Action,
	{_State259, CallbackOpType}:                    _GotoState72Action,
	{_State259, CallbackOpStatementType}:           _GotoState73Action,
	{_State259, UnsafeStatementType}:               _GotoState103Action,
	{_State259, JumpStatementType}:                 _GotoState85Action,
	{_State259, JumpTypeType}:                      _GotoState86Action,
	{_State259, FallthroughStatementType}:          _GotoState79Action,
	{_State259, AssignStatementType}:               _GotoState62Action,
	{_State259, UnaryOpAssignStatementType}:        _GotoState102Action,
	{_State259, BinaryOpAssignStatementType}:       _GotoState68Action,
	{_State259, ImportStatementType}:               _GotoState81Action,
	{_State259, VarDeclPatternType}:                _GotoState104Action,
	{_State259, VarOrLetType}:                      _GotoState24Action,
	{_State259, AssignPatternType}:                 _GotoState61Action,
	{_State259, ExprType}:                          _GotoState76Action,
	{_State259, OptionalLabelDeclType}:             _GotoState91Action,
	{_State259, SequenceExprType}:                  _GotoState99Action,
	{_State259, CallExprType}:                      _GotoState71Action,
	{_State259, AtomExprType}:                      _GotoState63Action,
	{_State259, ParseErrorExprType}:                _GotoState93Action,
	{_State259, LiteralExprType}:                   _GotoState87Action,
	{_State259, NamedExprType}:                     _GotoState90Action,
	{_State259, BlockExprType}:                     _GotoState70Action,
	{_State259, InitializeExprType}:                _GotoState84Action,
	{_State259, ImplicitStructExprType}:            _GotoState80Action,
	{_State259, AccessibleExprType}:                _GotoState56Action,
	{_State259, AccessExprType}:                    _GotoState55Action,
	{_State259, IndexExprType}:                     _GotoState82Action,
	{_State259, PostfixableExprType}:               _GotoState95Action,
	{_State259, PostfixUnaryExprType}:              _GotoState94Action,
	{_State259, PrefixableExprType}:                _GotoState98Action,
	{_State259, PrefixUnaryExprType}:               _GotoState96Action,
	{_State259, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State259, MulExprType}:                       _GotoState89Action,
	{_State259, BinaryMulExprType}:                 _GotoState67Action,
	{_State259, AddExprType}:                       _GotoState57Action,
	{_State259, BinaryAddExprType}:                 _GotoState64Action,
	{_State259, CmpExprType}:                       _GotoState74Action,
	{_State259, BinaryCmpExprType}:                 _GotoState66Action,
	{_State259, AndExprType}:                       _GotoState58Action,
	{_State259, BinaryAndExprType}:                 _GotoState65Action,
	{_State259, OrExprType}:                        _GotoState92Action,
	{_State259, BinaryOrExprType}:                  _GotoState69Action,
	{_State259, InitializableTypeExprType}:         _GotoState83Action,
	{_State259, SliceTypeExprType}:                 _GotoState101Action,
	{_State259, ArrayTypeExprType}:                 _GotoState60Action,
	{_State259, MapTypeExprType}:                   _GotoState88Action,
	{_State259, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State259, AnonymousFuncExprType}:             _GotoState59Action,
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
	{_State260, StatementType}:                     _GotoState363Action,
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
	{_State262, IdentifierToken}:                   _GotoState113Action,
	{_State262, UnderscoreToken}:                   _GotoState119Action,
	{_State262, StructToken}:                       _GotoState50Action,
	{_State262, EnumToken}:                         _GotoState110Action,
	{_State262, TraitToken}:                        _GotoState118Action,
	{_State262, FuncToken}:                         _GotoState112Action,
	{_State262, LparenToken}:                       _GotoState114Action,
	{_State262, LbracketToken}:                     _GotoState42Action,
	{_State262, DotToken}:                          _GotoState109Action,
	{_State262, QuestionToken}:                     _GotoState116Action,
	{_State262, ExclaimToken}:                      _GotoState111Action,
	{_State262, TildeTildeToken}:                   _GotoState117Action,
	{_State262, BitNegToken}:                       _GotoState108Action,
	{_State262, BitAndToken}:                       _GotoState107Action,
	{_State262, ParseErrorToken}:                   _GotoState115Action,
	{_State262, InitializableTypeExprType}:         _GotoState127Action,
	{_State262, SliceTypeExprType}:                 _GotoState101Action,
	{_State262, ArrayTypeExprType}:                 _GotoState60Action,
	{_State262, MapTypeExprType}:                   _GotoState88Action,
	{_State262, AtomTypeExprType}:                  _GotoState120Action,
	{_State262, NamedTypeExprType}:                 _GotoState128Action,
	{_State262, InferredTypeExprType}:              _GotoState126Action,
	{_State262, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State262, ReturnableTypeExprType}:            _GotoState132Action,
	{_State262, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State262, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State262, TypeExprType}:                      _GotoState364Action,
	{_State262, BinaryTypeExprType}:                _GotoState121Action,
	{_State262, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State262, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State262, TraitTypeExprType}:                 _GotoState133Action,
	{_State262, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State262, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State262, FuncTypeExprType}:                  _GotoState123Action,
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
	{_State263, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State263, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State263, TypeExprType}:                      _GotoState365Action,
	{_State263, BinaryTypeExprType}:                _GotoState121Action,
	{_State263, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State263, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State263, TraitTypeExprType}:                 _GotoState133Action,
	{_State263, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State263, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State263, FuncTypeExprType}:                  _GotoState123Action,
	{_State267, AssignToken}:                       _GotoState366Action,
	{_State269, RparenToken}:                       _GotoState368Action,
	{_State269, CommaToken}:                        _GotoState367Action,
	{_State272, AddToken}:                          _GotoState249Action,
	{_State272, SubToken}:                          _GotoState251Action,
	{_State272, MulToken}:                          _GotoState250Action,
	{_State272, BinaryTypeOpType}:                  _GotoState252Action,
	{_State273, LparenToken}:                       _GotoState43Action,
	{_State273, ImplicitStructExprType}:            _GotoState369Action,
	{_State274, IdentifierToken}:                   _GotoState370Action,
	{_State275, IntegerLiteralToken}:               _GotoState40Action,
	{_State275, FloatLiteralToken}:                 _GotoState35Action,
	{_State275, RuneLiteralToken}:                  _GotoState48Action,
	{_State275, StringLiteralToken}:                _GotoState49Action,
	{_State275, IdentifierToken}:                   _GotoState38Action,
	{_State275, UnderscoreToken}:                   _GotoState53Action,
	{_State275, TrueToken}:                         _GotoState52Action,
	{_State275, FalseToken}:                        _GotoState34Action,
	{_State275, ReturnToken}:                       _GotoState47Action,
	{_State275, BreakToken}:                        _GotoState28Action,
	{_State275, ContinueToken}:                     _GotoState30Action,
	{_State275, FallthroughToken}:                  _GotoState33Action,
	{_State275, UnsafeToken}:                       _GotoState54Action,
	{_State275, StructToken}:                       _GotoState50Action,
	{_State275, FuncToken}:                         _GotoState36Action,
	{_State275, AsyncToken}:                        _GotoState25Action,
	{_State275, DeferToken}:                        _GotoState32Action,
	{_State275, VarToken}:                          _GotoState15Action,
	{_State275, LetToken}:                          _GotoState12Action,
	{_State275, NotToken}:                          _GotoState45Action,
	{_State275, LabelDeclToken}:                    _GotoState41Action,
	{_State275, LparenToken}:                       _GotoState43Action,
	{_State275, LbracketToken}:                     _GotoState42Action,
	{_State275, SubToken}:                          _GotoState51Action,
	{_State275, MulToken}:                          _GotoState44Action,
	{_State275, BitNegToken}:                       _GotoState27Action,
	{_State275, BitAndToken}:                       _GotoState26Action,
	{_State275, GreaterToken}:                      _GotoState37Action,
	{_State275, ParseErrorToken}:                   _GotoState46Action,
	{_State275, SimpleStatementType}:               _GotoState278Action,
	{_State275, OptionalSimpleStatementType}:       _GotoState371Action,
	{_State275, ExprOrImproperStructStatementType}: _GotoState77Action,
	{_State275, ExprsType}:                         _GotoState78Action,
	{_State275, CallbackOpType}:                    _GotoState72Action,
	{_State275, CallbackOpStatementType}:           _GotoState73Action,
	{_State275, UnsafeStatementType}:               _GotoState103Action,
	{_State275, JumpStatementType}:                 _GotoState85Action,
	{_State275, JumpTypeType}:                      _GotoState86Action,
	{_State275, FallthroughStatementType}:          _GotoState79Action,
	{_State275, AssignStatementType}:               _GotoState62Action,
	{_State275, UnaryOpAssignStatementType}:        _GotoState102Action,
	{_State275, BinaryOpAssignStatementType}:       _GotoState68Action,
	{_State275, VarDeclPatternType}:                _GotoState104Action,
	{_State275, VarOrLetType}:                      _GotoState24Action,
	{_State275, AssignPatternType}:                 _GotoState61Action,
	{_State275, ExprType}:                          _GotoState76Action,
	{_State275, OptionalLabelDeclType}:             _GotoState91Action,
	{_State275, SequenceExprType}:                  _GotoState99Action,
	{_State275, CallExprType}:                      _GotoState71Action,
	{_State275, AtomExprType}:                      _GotoState63Action,
	{_State275, ParseErrorExprType}:                _GotoState93Action,
	{_State275, LiteralExprType}:                   _GotoState87Action,
	{_State275, NamedExprType}:                     _GotoState90Action,
	{_State275, BlockExprType}:                     _GotoState70Action,
	{_State275, InitializeExprType}:                _GotoState84Action,
	{_State275, ImplicitStructExprType}:            _GotoState80Action,
	{_State275, AccessibleExprType}:                _GotoState56Action,
	{_State275, AccessExprType}:                    _GotoState55Action,
	{_State275, IndexExprType}:                     _GotoState82Action,
	{_State275, PostfixableExprType}:               _GotoState95Action,
	{_State275, PostfixUnaryExprType}:              _GotoState94Action,
	{_State275, PrefixableExprType}:                _GotoState98Action,
	{_State275, PrefixUnaryExprType}:               _GotoState96Action,
	{_State275, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State275, MulExprType}:                       _GotoState89Action,
	{_State275, BinaryMulExprType}:                 _GotoState67Action,
	{_State275, AddExprType}:                       _GotoState57Action,
	{_State275, BinaryAddExprType}:                 _GotoState64Action,
	{_State275, CmpExprType}:                       _GotoState74Action,
	{_State275, BinaryCmpExprType}:                 _GotoState66Action,
	{_State275, AndExprType}:                       _GotoState58Action,
	{_State275, BinaryAndExprType}:                 _GotoState65Action,
	{_State275, OrExprType}:                        _GotoState92Action,
	{_State275, BinaryOrExprType}:                  _GotoState69Action,
	{_State275, InitializableTypeExprType}:         _GotoState83Action,
	{_State275, SliceTypeExprType}:                 _GotoState101Action,
	{_State275, ArrayTypeExprType}:                 _GotoState60Action,
	{_State275, MapTypeExprType}:                   _GotoState88Action,
	{_State275, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State275, AnonymousFuncExprType}:             _GotoState59Action,
	{_State276, IntegerLiteralToken}:               _GotoState40Action,
	{_State276, FloatLiteralToken}:                 _GotoState35Action,
	{_State276, RuneLiteralToken}:                  _GotoState48Action,
	{_State276, StringLiteralToken}:                _GotoState49Action,
	{_State276, IdentifierToken}:                   _GotoState38Action,
	{_State276, UnderscoreToken}:                   _GotoState53Action,
	{_State276, TrueToken}:                         _GotoState52Action,
	{_State276, FalseToken}:                        _GotoState34Action,
	{_State276, StructToken}:                       _GotoState50Action,
	{_State276, FuncToken}:                         _GotoState36Action,
	{_State276, VarToken}:                          _GotoState149Action,
	{_State276, LetToken}:                          _GotoState12Action,
	{_State276, NotToken}:                          _GotoState45Action,
	{_State276, LabelDeclToken}:                    _GotoState41Action,
	{_State276, LparenToken}:                       _GotoState43Action,
	{_State276, LbracketToken}:                     _GotoState42Action,
	{_State276, DotToken}:                          _GotoState148Action,
	{_State276, SubToken}:                          _GotoState51Action,
	{_State276, MulToken}:                          _GotoState44Action,
	{_State276, BitNegToken}:                       _GotoState27Action,
	{_State276, BitAndToken}:                       _GotoState26Action,
	{_State276, GreaterToken}:                      _GotoState37Action,
	{_State276, ParseErrorToken}:                   _GotoState46Action,
	{_State276, VarDeclPatternType}:                _GotoState104Action,
	{_State276, VarOrLetType}:                      _GotoState24Action,
	{_State276, AssignPatternType}:                 _GotoState150Action,
	{_State276, CasePatternType}:                   _GotoState372Action,
	{_State276, OptionalLabelDeclType}:             _GotoState153Action,
	{_State276, SequenceExprType}:                  _GotoState154Action,
	{_State276, CallExprType}:                      _GotoState71Action,
	{_State276, AtomExprType}:                      _GotoState63Action,
	{_State276, ParseErrorExprType}:                _GotoState93Action,
	{_State276, LiteralExprType}:                   _GotoState87Action,
	{_State276, NamedExprType}:                     _GotoState90Action,
	{_State276, BlockExprType}:                     _GotoState70Action,
	{_State276, InitializeExprType}:                _GotoState84Action,
	{_State276, ImplicitStructExprType}:            _GotoState80Action,
	{_State276, AccessibleExprType}:                _GotoState105Action,
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
	{_State280, RparenToken}:                       _GotoState373Action,
	{_State281, CommaToken}:                        _GotoState374Action,
	{_State285, RparenToken}:                       _GotoState375Action,
	{_State286, NewlinesToken}:                     _GotoState377Action,
	{_State286, CommaToken}:                        _GotoState376Action,
	{_State288, IdentifierToken}:                   _GotoState113Action,
	{_State288, UnderscoreToken}:                   _GotoState119Action,
	{_State288, StructToken}:                       _GotoState50Action,
	{_State288, EnumToken}:                         _GotoState110Action,
	{_State288, TraitToken}:                        _GotoState118Action,
	{_State288, FuncToken}:                         _GotoState112Action,
	{_State288, LparenToken}:                       _GotoState114Action,
	{_State288, LbracketToken}:                     _GotoState42Action,
	{_State288, DotToken}:                          _GotoState109Action,
	{_State288, QuestionToken}:                     _GotoState116Action,
	{_State288, ExclaimToken}:                      _GotoState111Action,
	{_State288, TildeTildeToken}:                   _GotoState117Action,
	{_State288, BitNegToken}:                       _GotoState108Action,
	{_State288, BitAndToken}:                       _GotoState107Action,
	{_State288, ParseErrorToken}:                   _GotoState115Action,
	{_State288, InitializableTypeExprType}:         _GotoState127Action,
	{_State288, SliceTypeExprType}:                 _GotoState101Action,
	{_State288, ArrayTypeExprType}:                 _GotoState60Action,
	{_State288, MapTypeExprType}:                   _GotoState88Action,
	{_State288, AtomTypeExprType}:                  _GotoState120Action,
	{_State288, NamedTypeExprType}:                 _GotoState128Action,
	{_State288, InferredTypeExprType}:              _GotoState126Action,
	{_State288, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State288, ReturnableTypeExprType}:            _GotoState132Action,
	{_State288, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State288, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State288, TypeExprType}:                      _GotoState378Action,
	{_State288, BinaryTypeExprType}:                _GotoState121Action,
	{_State288, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State288, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State288, TraitTypeExprType}:                 _GotoState133Action,
	{_State288, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State288, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State288, FuncTypeExprType}:                  _GotoState123Action,
	{_State289, IntegerLiteralToken}:               _GotoState379Action,
	{_State292, IntegerLiteralToken}:               _GotoState40Action,
	{_State292, FloatLiteralToken}:                 _GotoState35Action,
	{_State292, RuneLiteralToken}:                  _GotoState48Action,
	{_State292, StringLiteralToken}:                _GotoState49Action,
	{_State292, IdentifierToken}:                   _GotoState38Action,
	{_State292, UnderscoreToken}:                   _GotoState53Action,
	{_State292, TrueToken}:                         _GotoState52Action,
	{_State292, FalseToken}:                        _GotoState34Action,
	{_State292, StructToken}:                       _GotoState50Action,
	{_State292, FuncToken}:                         _GotoState36Action,
	{_State292, VarToken}:                          _GotoState15Action,
	{_State292, LetToken}:                          _GotoState12Action,
	{_State292, NotToken}:                          _GotoState45Action,
	{_State292, LabelDeclToken}:                    _GotoState41Action,
	{_State292, LparenToken}:                       _GotoState43Action,
	{_State292, LbracketToken}:                     _GotoState42Action,
	{_State292, SubToken}:                          _GotoState51Action,
	{_State292, MulToken}:                          _GotoState44Action,
	{_State292, BitNegToken}:                       _GotoState27Action,
	{_State292, BitAndToken}:                       _GotoState26Action,
	{_State292, GreaterToken}:                      _GotoState37Action,
	{_State292, ParseErrorToken}:                   _GotoState46Action,
	{_State292, VarDeclPatternType}:                _GotoState104Action,
	{_State292, VarOrLetType}:                      _GotoState24Action,
	{_State292, ExprType}:                          _GotoState380Action,
	{_State292, OptionalLabelDeclType}:             _GotoState91Action,
	{_State292, SequenceExprType}:                  _GotoState106Action,
	{_State292, CallExprType}:                      _GotoState71Action,
	{_State292, AtomExprType}:                      _GotoState63Action,
	{_State292, ParseErrorExprType}:                _GotoState93Action,
	{_State292, LiteralExprType}:                   _GotoState87Action,
	{_State292, NamedExprType}:                     _GotoState90Action,
	{_State292, BlockExprType}:                     _GotoState70Action,
	{_State292, InitializeExprType}:                _GotoState84Action,
	{_State292, ImplicitStructExprType}:            _GotoState80Action,
	{_State292, AccessibleExprType}:                _GotoState105Action,
	{_State292, AccessExprType}:                    _GotoState55Action,
	{_State292, IndexExprType}:                     _GotoState82Action,
	{_State292, PostfixableExprType}:               _GotoState95Action,
	{_State292, PostfixUnaryExprType}:              _GotoState94Action,
	{_State292, PrefixableExprType}:                _GotoState98Action,
	{_State292, PrefixUnaryExprType}:               _GotoState96Action,
	{_State292, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State292, MulExprType}:                       _GotoState89Action,
	{_State292, BinaryMulExprType}:                 _GotoState67Action,
	{_State292, AddExprType}:                       _GotoState57Action,
	{_State292, BinaryAddExprType}:                 _GotoState64Action,
	{_State292, CmpExprType}:                       _GotoState74Action,
	{_State292, BinaryCmpExprType}:                 _GotoState66Action,
	{_State292, AndExprType}:                       _GotoState58Action,
	{_State292, BinaryAndExprType}:                 _GotoState65Action,
	{_State292, OrExprType}:                        _GotoState92Action,
	{_State292, BinaryOrExprType}:                  _GotoState69Action,
	{_State292, InitializableTypeExprType}:         _GotoState83Action,
	{_State292, SliceTypeExprType}:                 _GotoState101Action,
	{_State292, ArrayTypeExprType}:                 _GotoState60Action,
	{_State292, MapTypeExprType}:                   _GotoState88Action,
	{_State292, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State292, AnonymousFuncExprType}:             _GotoState59Action,
	{_State294, IntegerLiteralToken}:               _GotoState40Action,
	{_State294, FloatLiteralToken}:                 _GotoState35Action,
	{_State294, RuneLiteralToken}:                  _GotoState48Action,
	{_State294, StringLiteralToken}:                _GotoState49Action,
	{_State294, IdentifierToken}:                   _GotoState38Action,
	{_State294, UnderscoreToken}:                   _GotoState53Action,
	{_State294, TrueToken}:                         _GotoState52Action,
	{_State294, FalseToken}:                        _GotoState34Action,
	{_State294, StructToken}:                       _GotoState50Action,
	{_State294, FuncToken}:                         _GotoState36Action,
	{_State294, VarToken}:                          _GotoState15Action,
	{_State294, LetToken}:                          _GotoState12Action,
	{_State294, NotToken}:                          _GotoState45Action,
	{_State294, LabelDeclToken}:                    _GotoState41Action,
	{_State294, LparenToken}:                       _GotoState43Action,
	{_State294, LbracketToken}:                     _GotoState42Action,
	{_State294, SubToken}:                          _GotoState51Action,
	{_State294, MulToken}:                          _GotoState44Action,
	{_State294, BitNegToken}:                       _GotoState27Action,
	{_State294, BitAndToken}:                       _GotoState26Action,
	{_State294, GreaterToken}:                      _GotoState37Action,
	{_State294, ParseErrorToken}:                   _GotoState46Action,
	{_State294, VarDeclPatternType}:                _GotoState104Action,
	{_State294, VarOrLetType}:                      _GotoState24Action,
	{_State294, ExprType}:                          _GotoState381Action,
	{_State294, OptionalLabelDeclType}:             _GotoState91Action,
	{_State294, SequenceExprType}:                  _GotoState106Action,
	{_State294, CallExprType}:                      _GotoState71Action,
	{_State294, AtomExprType}:                      _GotoState63Action,
	{_State294, ParseErrorExprType}:                _GotoState93Action,
	{_State294, LiteralExprType}:                   _GotoState87Action,
	{_State294, NamedExprType}:                     _GotoState90Action,
	{_State294, BlockExprType}:                     _GotoState70Action,
	{_State294, InitializeExprType}:                _GotoState84Action,
	{_State294, ImplicitStructExprType}:            _GotoState80Action,
	{_State294, AccessibleExprType}:                _GotoState105Action,
	{_State294, AccessExprType}:                    _GotoState55Action,
	{_State294, IndexExprType}:                     _GotoState82Action,
	{_State294, PostfixableExprType}:               _GotoState95Action,
	{_State294, PostfixUnaryExprType}:              _GotoState94Action,
	{_State294, PrefixableExprType}:                _GotoState98Action,
	{_State294, PrefixUnaryExprType}:               _GotoState96Action,
	{_State294, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State294, MulExprType}:                       _GotoState89Action,
	{_State294, BinaryMulExprType}:                 _GotoState67Action,
	{_State294, AddExprType}:                       _GotoState57Action,
	{_State294, BinaryAddExprType}:                 _GotoState64Action,
	{_State294, CmpExprType}:                       _GotoState74Action,
	{_State294, BinaryCmpExprType}:                 _GotoState66Action,
	{_State294, AndExprType}:                       _GotoState58Action,
	{_State294, BinaryAndExprType}:                 _GotoState65Action,
	{_State294, OrExprType}:                        _GotoState92Action,
	{_State294, BinaryOrExprType}:                  _GotoState69Action,
	{_State294, InitializableTypeExprType}:         _GotoState83Action,
	{_State294, SliceTypeExprType}:                 _GotoState101Action,
	{_State294, ArrayTypeExprType}:                 _GotoState60Action,
	{_State294, MapTypeExprType}:                   _GotoState88Action,
	{_State294, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State294, AnonymousFuncExprType}:             _GotoState59Action,
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
	{_State295, ExprType}:                          _GotoState382Action,
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
	{_State297, IntegerLiteralToken}:               _GotoState40Action,
	{_State297, FloatLiteralToken}:                 _GotoState35Action,
	{_State297, RuneLiteralToken}:                  _GotoState48Action,
	{_State297, StringLiteralToken}:                _GotoState49Action,
	{_State297, IdentifierToken}:                   _GotoState167Action,
	{_State297, UnderscoreToken}:                   _GotoState53Action,
	{_State297, TrueToken}:                         _GotoState52Action,
	{_State297, FalseToken}:                        _GotoState34Action,
	{_State297, StructToken}:                       _GotoState50Action,
	{_State297, FuncToken}:                         _GotoState36Action,
	{_State297, VarToken}:                          _GotoState15Action,
	{_State297, LetToken}:                          _GotoState12Action,
	{_State297, NotToken}:                          _GotoState45Action,
	{_State297, LabelDeclToken}:                    _GotoState41Action,
	{_State297, LparenToken}:                       _GotoState43Action,
	{_State297, LbracketToken}:                     _GotoState42Action,
	{_State297, ColonToken}:                        _GotoState165Action,
	{_State297, EllipsisToken}:                     _GotoState166Action,
	{_State297, SubToken}:                          _GotoState51Action,
	{_State297, MulToken}:                          _GotoState44Action,
	{_State297, BitNegToken}:                       _GotoState27Action,
	{_State297, BitAndToken}:                       _GotoState26Action,
	{_State297, GreaterToken}:                      _GotoState37Action,
	{_State297, ParseErrorToken}:                   _GotoState46Action,
	{_State297, VarDeclPatternType}:                _GotoState104Action,
	{_State297, VarOrLetType}:                      _GotoState24Action,
	{_State297, ExprType}:                          _GotoState171Action,
	{_State297, OptionalLabelDeclType}:             _GotoState91Action,
	{_State297, SequenceExprType}:                  _GotoState106Action,
	{_State297, CallExprType}:                      _GotoState71Action,
	{_State297, ArgumentType}:                      _GotoState383Action,
	{_State297, ColonExprType}:                     _GotoState170Action,
	{_State297, AtomExprType}:                      _GotoState63Action,
	{_State297, ParseErrorExprType}:                _GotoState93Action,
	{_State297, LiteralExprType}:                   _GotoState87Action,
	{_State297, NamedExprType}:                     _GotoState90Action,
	{_State297, BlockExprType}:                     _GotoState70Action,
	{_State297, InitializeExprType}:                _GotoState84Action,
	{_State297, ImplicitStructExprType}:            _GotoState80Action,
	{_State297, AccessibleExprType}:                _GotoState105Action,
	{_State297, AccessExprType}:                    _GotoState55Action,
	{_State297, IndexExprType}:                     _GotoState82Action,
	{_State297, PostfixableExprType}:               _GotoState95Action,
	{_State297, PostfixUnaryExprType}:              _GotoState94Action,
	{_State297, PrefixableExprType}:                _GotoState98Action,
	{_State297, PrefixUnaryExprType}:               _GotoState96Action,
	{_State297, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State297, MulExprType}:                       _GotoState89Action,
	{_State297, BinaryMulExprType}:                 _GotoState67Action,
	{_State297, AddExprType}:                       _GotoState57Action,
	{_State297, BinaryAddExprType}:                 _GotoState64Action,
	{_State297, CmpExprType}:                       _GotoState74Action,
	{_State297, BinaryCmpExprType}:                 _GotoState66Action,
	{_State297, AndExprType}:                       _GotoState58Action,
	{_State297, BinaryAndExprType}:                 _GotoState65Action,
	{_State297, OrExprType}:                        _GotoState92Action,
	{_State297, BinaryOrExprType}:                  _GotoState69Action,
	{_State297, InitializableTypeExprType}:         _GotoState83Action,
	{_State297, SliceTypeExprType}:                 _GotoState101Action,
	{_State297, ArrayTypeExprType}:                 _GotoState60Action,
	{_State297, MapTypeExprType}:                   _GotoState88Action,
	{_State297, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State297, AnonymousFuncExprType}:             _GotoState59Action,
	{_State298, RparenToken}:                       _GotoState384Action,
	{_State300, NewlinesToken}:                     _GotoState386Action,
	{_State300, CommaToken}:                        _GotoState385Action,
	{_State301, GreaterToken}:                      _GotoState387Action,
	{_State302, CommaToken}:                        _GotoState388Action,
	{_State303, RbracketToken}:                     _GotoState389Action,
	{_State304, AddToken}:                          _GotoState249Action,
	{_State304, SubToken}:                          _GotoState251Action,
	{_State304, MulToken}:                          _GotoState250Action,
	{_State304, BinaryTypeOpType}:                  _GotoState252Action,
	{_State306, RbracketToken}:                     _GotoState390Action,
	{_State308, IntegerLiteralToken}:               _GotoState40Action,
	{_State308, FloatLiteralToken}:                 _GotoState35Action,
	{_State308, RuneLiteralToken}:                  _GotoState48Action,
	{_State308, StringLiteralToken}:                _GotoState49Action,
	{_State308, IdentifierToken}:                   _GotoState167Action,
	{_State308, UnderscoreToken}:                   _GotoState53Action,
	{_State308, TrueToken}:                         _GotoState52Action,
	{_State308, FalseToken}:                        _GotoState34Action,
	{_State308, StructToken}:                       _GotoState50Action,
	{_State308, FuncToken}:                         _GotoState36Action,
	{_State308, VarToken}:                          _GotoState15Action,
	{_State308, LetToken}:                          _GotoState12Action,
	{_State308, NotToken}:                          _GotoState45Action,
	{_State308, LabelDeclToken}:                    _GotoState41Action,
	{_State308, LparenToken}:                       _GotoState43Action,
	{_State308, LbracketToken}:                     _GotoState42Action,
	{_State308, ColonToken}:                        _GotoState165Action,
	{_State308, EllipsisToken}:                     _GotoState166Action,
	{_State308, SubToken}:                          _GotoState51Action,
	{_State308, MulToken}:                          _GotoState44Action,
	{_State308, BitNegToken}:                       _GotoState27Action,
	{_State308, BitAndToken}:                       _GotoState26Action,
	{_State308, GreaterToken}:                      _GotoState37Action,
	{_State308, ParseErrorToken}:                   _GotoState46Action,
	{_State308, VarDeclPatternType}:                _GotoState104Action,
	{_State308, VarOrLetType}:                      _GotoState24Action,
	{_State308, ExprType}:                          _GotoState171Action,
	{_State308, OptionalLabelDeclType}:             _GotoState91Action,
	{_State308, SequenceExprType}:                  _GotoState106Action,
	{_State308, CallExprType}:                      _GotoState71Action,
	{_State308, ProperArgumentsType}:               _GotoState172Action,
	{_State308, ArgumentsType}:                     _GotoState391Action,
	{_State308, ArgumentType}:                      _GotoState168Action,
	{_State308, ColonExprType}:                     _GotoState170Action,
	{_State308, AtomExprType}:                      _GotoState63Action,
	{_State308, ParseErrorExprType}:                _GotoState93Action,
	{_State308, LiteralExprType}:                   _GotoState87Action,
	{_State308, NamedExprType}:                     _GotoState90Action,
	{_State308, BlockExprType}:                     _GotoState70Action,
	{_State308, InitializeExprType}:                _GotoState84Action,
	{_State308, ImplicitStructExprType}:            _GotoState80Action,
	{_State308, AccessibleExprType}:                _GotoState105Action,
	{_State308, AccessExprType}:                    _GotoState55Action,
	{_State308, IndexExprType}:                     _GotoState82Action,
	{_State308, PostfixableExprType}:               _GotoState95Action,
	{_State308, PostfixUnaryExprType}:              _GotoState94Action,
	{_State308, PrefixableExprType}:                _GotoState98Action,
	{_State308, PrefixUnaryExprType}:               _GotoState96Action,
	{_State308, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State308, MulExprType}:                       _GotoState89Action,
	{_State308, BinaryMulExprType}:                 _GotoState67Action,
	{_State308, AddExprType}:                       _GotoState57Action,
	{_State308, BinaryAddExprType}:                 _GotoState64Action,
	{_State308, CmpExprType}:                       _GotoState74Action,
	{_State308, BinaryCmpExprType}:                 _GotoState66Action,
	{_State308, AndExprType}:                       _GotoState58Action,
	{_State308, BinaryAndExprType}:                 _GotoState65Action,
	{_State308, OrExprType}:                        _GotoState92Action,
	{_State308, BinaryOrExprType}:                  _GotoState69Action,
	{_State308, InitializableTypeExprType}:         _GotoState83Action,
	{_State308, SliceTypeExprType}:                 _GotoState101Action,
	{_State308, ArrayTypeExprType}:                 _GotoState60Action,
	{_State308, MapTypeExprType}:                   _GotoState88Action,
	{_State308, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State308, AnonymousFuncExprType}:             _GotoState59Action,
	{_State309, MulToken}:                          _GotoState220Action,
	{_State309, DivToken}:                          _GotoState218Action,
	{_State309, ModToken}:                          _GotoState219Action,
	{_State309, BitAndToken}:                       _GotoState215Action,
	{_State309, BitLshiftToken}:                    _GotoState216Action,
	{_State309, BitRshiftToken}:                    _GotoState217Action,
	{_State309, MulOpType}:                         _GotoState221Action,
	{_State310, EqualToken}:                        _GotoState204Action,
	{_State310, NotEqualToken}:                     _GotoState209Action,
	{_State310, LessToken}:                         _GotoState207Action,
	{_State310, LessOrEqualToken}:                  _GotoState208Action,
	{_State310, GreaterToken}:                      _GotoState205Action,
	{_State310, GreaterOrEqualToken}:               _GotoState206Action,
	{_State310, CmpOpType}:                         _GotoState210Action,
	{_State312, AddToken}:                          _GotoState195Action,
	{_State312, SubToken}:                          _GotoState198Action,
	{_State312, BitXorToken}:                       _GotoState197Action,
	{_State312, BitOrToken}:                        _GotoState196Action,
	{_State312, AddOpType}:                         _GotoState199Action,
	{_State314, RparenToken}:                       _GotoState392Action,
	{_State315, CommaToken}:                        _GotoState211Action,
	{_State317, ForToken}:                          _GotoState393Action,
	{_State318, InToken}:                           _GotoState395Action,
	{_State318, AssignToken}:                       _GotoState394Action,
	{_State319, SemicolonToken}:                    _GotoState396Action,
	{_State320, DoToken}:                           _GotoState397Action,
	{_State321, IntegerLiteralToken}:               _GotoState40Action,
	{_State321, FloatLiteralToken}:                 _GotoState35Action,
	{_State321, RuneLiteralToken}:                  _GotoState48Action,
	{_State321, StringLiteralToken}:                _GotoState49Action,
	{_State321, IdentifierToken}:                   _GotoState38Action,
	{_State321, UnderscoreToken}:                   _GotoState53Action,
	{_State321, TrueToken}:                         _GotoState52Action,
	{_State321, FalseToken}:                        _GotoState34Action,
	{_State321, StructToken}:                       _GotoState50Action,
	{_State321, FuncToken}:                         _GotoState36Action,
	{_State321, VarToken}:                          _GotoState149Action,
	{_State321, LetToken}:                          _GotoState12Action,
	{_State321, NotToken}:                          _GotoState45Action,
	{_State321, LabelDeclToken}:                    _GotoState41Action,
	{_State321, LparenToken}:                       _GotoState43Action,
	{_State321, LbracketToken}:                     _GotoState42Action,
	{_State321, DotToken}:                          _GotoState148Action,
	{_State321, SubToken}:                          _GotoState51Action,
	{_State321, MulToken}:                          _GotoState44Action,
	{_State321, BitNegToken}:                       _GotoState27Action,
	{_State321, BitAndToken}:                       _GotoState26Action,
	{_State321, GreaterToken}:                      _GotoState37Action,
	{_State321, ParseErrorToken}:                   _GotoState46Action,
	{_State321, CasePatternsType}:                  _GotoState398Action,
	{_State321, VarDeclPatternType}:                _GotoState104Action,
	{_State321, VarOrLetType}:                      _GotoState24Action,
	{_State321, AssignPatternType}:                 _GotoState150Action,
	{_State321, CasePatternType}:                   _GotoState151Action,
	{_State321, OptionalLabelDeclType}:             _GotoState153Action,
	{_State321, SequenceExprType}:                  _GotoState154Action,
	{_State321, CallExprType}:                      _GotoState71Action,
	{_State321, AtomExprType}:                      _GotoState63Action,
	{_State321, ParseErrorExprType}:                _GotoState93Action,
	{_State321, LiteralExprType}:                   _GotoState87Action,
	{_State321, NamedExprType}:                     _GotoState90Action,
	{_State321, BlockExprType}:                     _GotoState70Action,
	{_State321, InitializeExprType}:                _GotoState84Action,
	{_State321, ImplicitStructExprType}:            _GotoState80Action,
	{_State321, AccessibleExprType}:                _GotoState105Action,
	{_State321, AccessExprType}:                    _GotoState55Action,
	{_State321, IndexExprType}:                     _GotoState82Action,
	{_State321, PostfixableExprType}:               _GotoState95Action,
	{_State321, PostfixUnaryExprType}:              _GotoState94Action,
	{_State321, PrefixableExprType}:                _GotoState98Action,
	{_State321, PrefixUnaryExprType}:               _GotoState96Action,
	{_State321, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State321, MulExprType}:                       _GotoState89Action,
	{_State321, BinaryMulExprType}:                 _GotoState67Action,
	{_State321, AddExprType}:                       _GotoState57Action,
	{_State321, BinaryAddExprType}:                 _GotoState64Action,
	{_State321, CmpExprType}:                       _GotoState74Action,
	{_State321, BinaryCmpExprType}:                 _GotoState66Action,
	{_State321, AndExprType}:                       _GotoState58Action,
	{_State321, BinaryAndExprType}:                 _GotoState65Action,
	{_State321, OrExprType}:                        _GotoState92Action,
	{_State321, BinaryOrExprType}:                  _GotoState69Action,
	{_State321, InitializableTypeExprType}:         _GotoState83Action,
	{_State321, SliceTypeExprType}:                 _GotoState101Action,
	{_State321, ArrayTypeExprType}:                 _GotoState60Action,
	{_State321, MapTypeExprType}:                   _GotoState88Action,
	{_State321, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State321, AnonymousFuncExprType}:             _GotoState59Action,
	{_State322, LbraceToken}:                       _GotoState11Action,
	{_State322, StatementBlockType}:                _GotoState399Action,
	{_State324, LbraceToken}:                       _GotoState11Action,
	{_State324, StatementBlockType}:                _GotoState400Action,
	{_State325, AndToken}:                          _GotoState200Action,
	{_State326, RparenToken}:                       _GotoState401Action,
	{_State327, NewlinesToken}:                     _GotoState402Action,
	{_State327, OrToken}:                           _GotoState403Action,
	{_State328, NewlinesToken}:                     _GotoState404Action,
	{_State328, OrToken}:                           _GotoState405Action,
	{_State329, IdentifierToken}:                   _GotoState113Action,
	{_State329, UnderscoreToken}:                   _GotoState119Action,
	{_State329, StructToken}:                       _GotoState50Action,
	{_State329, EnumToken}:                         _GotoState110Action,
	{_State329, TraitToken}:                        _GotoState118Action,
	{_State329, FuncToken}:                         _GotoState112Action,
	{_State329, LparenToken}:                       _GotoState114Action,
	{_State329, LbracketToken}:                     _GotoState42Action,
	{_State329, DotToken}:                          _GotoState109Action,
	{_State329, QuestionToken}:                     _GotoState116Action,
	{_State329, ExclaimToken}:                      _GotoState111Action,
	{_State329, TildeTildeToken}:                   _GotoState117Action,
	{_State329, BitNegToken}:                       _GotoState108Action,
	{_State329, BitAndToken}:                       _GotoState107Action,
	{_State329, ParseErrorToken}:                   _GotoState115Action,
	{_State329, InitializableTypeExprType}:         _GotoState127Action,
	{_State329, SliceTypeExprType}:                 _GotoState101Action,
	{_State329, ArrayTypeExprType}:                 _GotoState60Action,
	{_State329, MapTypeExprType}:                   _GotoState88Action,
	{_State329, AtomTypeExprType}:                  _GotoState120Action,
	{_State329, NamedTypeExprType}:                 _GotoState128Action,
	{_State329, InferredTypeExprType}:              _GotoState126Action,
	{_State329, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State329, ReturnableTypeExprType}:            _GotoState132Action,
	{_State329, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State329, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State329, TypeExprType}:                      _GotoState406Action,
	{_State329, BinaryTypeExprType}:                _GotoState121Action,
	{_State329, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State329, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State329, TraitTypeExprType}:                 _GotoState133Action,
	{_State329, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State329, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State329, FuncTypeExprType}:                  _GotoState123Action,
	{_State330, IdentifierToken}:                   _GotoState113Action,
	{_State330, UnderscoreToken}:                   _GotoState119Action,
	{_State330, StructToken}:                       _GotoState50Action,
	{_State330, EnumToken}:                         _GotoState110Action,
	{_State330, TraitToken}:                        _GotoState118Action,
	{_State330, FuncToken}:                         _GotoState112Action,
	{_State330, LparenToken}:                       _GotoState114Action,
	{_State330, LbracketToken}:                     _GotoState42Action,
	{_State330, DotToken}:                          _GotoState338Action,
	{_State330, QuestionToken}:                     _GotoState116Action,
	{_State330, ExclaimToken}:                      _GotoState111Action,
	{_State330, DollarLbracketToken}:               _GotoState184Action,
	{_State330, EllipsisToken}:                     _GotoState407Action,
	{_State330, TildeTildeToken}:                   _GotoState117Action,
	{_State330, BitNegToken}:                       _GotoState108Action,
	{_State330, BitAndToken}:                       _GotoState107Action,
	{_State330, ParseErrorToken}:                   _GotoState115Action,
	{_State330, GenericTypeArgumentsType}:          _GotoState235Action,
	{_State330, InitializableTypeExprType}:         _GotoState127Action,
	{_State330, SliceTypeExprType}:                 _GotoState101Action,
	{_State330, ArrayTypeExprType}:                 _GotoState60Action,
	{_State330, MapTypeExprType}:                   _GotoState88Action,
	{_State330, AtomTypeExprType}:                  _GotoState120Action,
	{_State330, NamedTypeExprType}:                 _GotoState128Action,
	{_State330, InferredTypeExprType}:              _GotoState126Action,
	{_State330, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State330, ReturnableTypeExprType}:            _GotoState132Action,
	{_State330, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State330, PrefixUnaryTypeOpType}:             _GotoState130Action,
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
	{_State331, DotToken}:                          _GotoState109Action,
	{_State331, QuestionToken}:                     _GotoState116Action,
	{_State331, ExclaimToken}:                      _GotoState111Action,
	{_State331, EllipsisToken}:                     _GotoState409Action,
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
	{_State331, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State331, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State331, TypeExprType}:                      _GotoState410Action,
	{_State331, BinaryTypeExprType}:                _GotoState121Action,
	{_State331, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State331, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State331, TraitTypeExprType}:                 _GotoState133Action,
	{_State331, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State331, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State331, FuncTypeExprType}:                  _GotoState123Action,
	{_State333, RparenToken}:                       _GotoState411Action,
	{_State334, CommaToken}:                        _GotoState412Action,
	{_State335, AddToken}:                          _GotoState249Action,
	{_State335, SubToken}:                          _GotoState251Action,
	{_State335, MulToken}:                          _GotoState250Action,
	{_State335, BinaryTypeOpType}:                  _GotoState252Action,
	{_State336, DollarLbracketToken}:               _GotoState184Action,
	{_State336, GenericTypeArgumentsType}:          _GotoState413Action,
	{_State337, LparenToken}:                       _GotoState414Action,
	{_State338, IdentifierToken}:                   _GotoState336Action,
	{_State339, AssignToken}:                       _GotoState347Action,
	{_State339, AddToken}:                          _GotoState249Action,
	{_State339, SubToken}:                          _GotoState251Action,
	{_State339, MulToken}:                          _GotoState250Action,
	{_State339, BinaryTypeOpType}:                  _GotoState252Action,
	{_State339, OptionalDefaultType}:               _GotoState415Action,
	{_State340, AddToken}:                          _GotoState249Action,
	{_State340, SubToken}:                          _GotoState251Action,
	{_State340, MulToken}:                          _GotoState250Action,
	{_State340, BinaryTypeOpType}:                  _GotoState252Action,
	{_State341, IdentifierToken}:                   _GotoState237Action,
	{_State341, UnderscoreToken}:                   _GotoState238Action,
	{_State341, UnsafeToken}:                       _GotoState54Action,
	{_State341, StructToken}:                       _GotoState50Action,
	{_State341, EnumToken}:                         _GotoState110Action,
	{_State341, TraitToken}:                        _GotoState118Action,
	{_State341, FuncToken}:                         _GotoState236Action,
	{_State341, LparenToken}:                       _GotoState114Action,
	{_State341, LbracketToken}:                     _GotoState42Action,
	{_State341, DotToken}:                          _GotoState109Action,
	{_State341, QuestionToken}:                     _GotoState116Action,
	{_State341, ExclaimToken}:                      _GotoState111Action,
	{_State341, TildeTildeToken}:                   _GotoState117Action,
	{_State341, BitNegToken}:                       _GotoState108Action,
	{_State341, BitAndToken}:                       _GotoState107Action,
	{_State341, ParseErrorToken}:                   _GotoState115Action,
	{_State341, UnsafeStatementType}:               _GotoState246Action,
	{_State341, InitializableTypeExprType}:         _GotoState127Action,
	{_State341, SliceTypeExprType}:                 _GotoState101Action,
	{_State341, ArrayTypeExprType}:                 _GotoState60Action,
	{_State341, MapTypeExprType}:                   _GotoState88Action,
	{_State341, AtomTypeExprType}:                  _GotoState120Action,
	{_State341, NamedTypeExprType}:                 _GotoState128Action,
	{_State341, InferredTypeExprType}:              _GotoState126Action,
	{_State341, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State341, ReturnableTypeExprType}:            _GotoState132Action,
	{_State341, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State341, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State341, TypeExprType}:                      _GotoState245Action,
	{_State341, BinaryTypeExprType}:                _GotoState121Action,
	{_State341, FieldDefType}:                      _GotoState416Action,
	{_State341, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State341, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State341, TraitTypeExprType}:                 _GotoState133Action,
	{_State341, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State341, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State341, FuncTypeExprType}:                  _GotoState123Action,
	{_State341, MethodSignatureType}:               _GotoState242Action,
	{_State345, IdentifierToken}:                   _GotoState237Action,
	{_State345, UnderscoreToken}:                   _GotoState238Action,
	{_State345, UnsafeToken}:                       _GotoState54Action,
	{_State345, StructToken}:                       _GotoState50Action,
	{_State345, EnumToken}:                         _GotoState110Action,
	{_State345, TraitToken}:                        _GotoState118Action,
	{_State345, FuncToken}:                         _GotoState236Action,
	{_State345, LparenToken}:                       _GotoState114Action,
	{_State345, LbracketToken}:                     _GotoState42Action,
	{_State345, DotToken}:                          _GotoState109Action,
	{_State345, QuestionToken}:                     _GotoState116Action,
	{_State345, ExclaimToken}:                      _GotoState111Action,
	{_State345, TildeTildeToken}:                   _GotoState117Action,
	{_State345, BitNegToken}:                       _GotoState108Action,
	{_State345, BitAndToken}:                       _GotoState107Action,
	{_State345, ParseErrorToken}:                   _GotoState115Action,
	{_State345, UnsafeStatementType}:               _GotoState246Action,
	{_State345, InitializableTypeExprType}:         _GotoState127Action,
	{_State345, SliceTypeExprType}:                 _GotoState101Action,
	{_State345, ArrayTypeExprType}:                 _GotoState60Action,
	{_State345, MapTypeExprType}:                   _GotoState88Action,
	{_State345, AtomTypeExprType}:                  _GotoState120Action,
	{_State345, NamedTypeExprType}:                 _GotoState128Action,
	{_State345, InferredTypeExprType}:              _GotoState126Action,
	{_State345, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State345, ReturnableTypeExprType}:            _GotoState132Action,
	{_State345, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State345, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State345, TypeExprType}:                      _GotoState245Action,
	{_State345, BinaryTypeExprType}:                _GotoState121Action,
	{_State345, FieldDefType}:                      _GotoState417Action,
	{_State345, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State345, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State345, TraitTypeExprType}:                 _GotoState133Action,
	{_State345, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State345, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State345, FuncTypeExprType}:                  _GotoState123Action,
	{_State345, MethodSignatureType}:               _GotoState242Action,
	{_State346, IdentifierToken}:                   _GotoState237Action,
	{_State346, UnderscoreToken}:                   _GotoState238Action,
	{_State346, UnsafeToken}:                       _GotoState54Action,
	{_State346, StructToken}:                       _GotoState50Action,
	{_State346, EnumToken}:                         _GotoState110Action,
	{_State346, TraitToken}:                        _GotoState118Action,
	{_State346, FuncToken}:                         _GotoState236Action,
	{_State346, LparenToken}:                       _GotoState114Action,
	{_State346, LbracketToken}:                     _GotoState42Action,
	{_State346, DotToken}:                          _GotoState109Action,
	{_State346, QuestionToken}:                     _GotoState116Action,
	{_State346, ExclaimToken}:                      _GotoState111Action,
	{_State346, TildeTildeToken}:                   _GotoState117Action,
	{_State346, BitNegToken}:                       _GotoState108Action,
	{_State346, BitAndToken}:                       _GotoState107Action,
	{_State346, ParseErrorToken}:                   _GotoState115Action,
	{_State346, UnsafeStatementType}:               _GotoState246Action,
	{_State346, InitializableTypeExprType}:         _GotoState127Action,
	{_State346, SliceTypeExprType}:                 _GotoState101Action,
	{_State346, ArrayTypeExprType}:                 _GotoState60Action,
	{_State346, MapTypeExprType}:                   _GotoState88Action,
	{_State346, AtomTypeExprType}:                  _GotoState120Action,
	{_State346, NamedTypeExprType}:                 _GotoState128Action,
	{_State346, InferredTypeExprType}:              _GotoState126Action,
	{_State346, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State346, ReturnableTypeExprType}:            _GotoState132Action,
	{_State346, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State346, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State346, TypeExprType}:                      _GotoState245Action,
	{_State346, BinaryTypeExprType}:                _GotoState121Action,
	{_State346, FieldDefType}:                      _GotoState418Action,
	{_State346, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State346, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State346, TraitTypeExprType}:                 _GotoState133Action,
	{_State346, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State346, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State346, FuncTypeExprType}:                  _GotoState123Action,
	{_State346, MethodSignatureType}:               _GotoState242Action,
	{_State347, DefaultToken}:                      _GotoState419Action,
	{_State349, RparenToken}:                       _GotoState420Action,
	{_State352, IdentifierToken}:                   _GotoState113Action,
	{_State352, UnderscoreToken}:                   _GotoState119Action,
	{_State352, StructToken}:                       _GotoState50Action,
	{_State352, EnumToken}:                         _GotoState110Action,
	{_State352, TraitToken}:                        _GotoState118Action,
	{_State352, FuncToken}:                         _GotoState112Action,
	{_State352, LparenToken}:                       _GotoState114Action,
	{_State352, LbracketToken}:                     _GotoState42Action,
	{_State352, DotToken}:                          _GotoState109Action,
	{_State352, QuestionToken}:                     _GotoState116Action,
	{_State352, ExclaimToken}:                      _GotoState111Action,
	{_State352, TildeTildeToken}:                   _GotoState117Action,
	{_State352, BitNegToken}:                       _GotoState108Action,
	{_State352, BitAndToken}:                       _GotoState107Action,
	{_State352, ParseErrorToken}:                   _GotoState115Action,
	{_State352, InitializableTypeExprType}:         _GotoState127Action,
	{_State352, SliceTypeExprType}:                 _GotoState101Action,
	{_State352, ArrayTypeExprType}:                 _GotoState60Action,
	{_State352, MapTypeExprType}:                   _GotoState88Action,
	{_State352, AtomTypeExprType}:                  _GotoState120Action,
	{_State352, NamedTypeExprType}:                 _GotoState128Action,
	{_State352, InferredTypeExprType}:              _GotoState126Action,
	{_State352, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State352, ReturnableTypeExprType}:            _GotoState132Action,
	{_State352, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State352, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State352, TypeExprType}:                      _GotoState421Action,
	{_State352, BinaryTypeExprType}:                _GotoState121Action,
	{_State352, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State352, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State352, TraitTypeExprType}:                 _GotoState133Action,
	{_State352, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State352, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State352, FuncTypeExprType}:                  _GotoState123Action,
	{_State354, RbracketToken}:                     _GotoState422Action,
	{_State355, CommaToken}:                        _GotoState423Action,
	{_State356, IdentifierToken}:                   _GotoState256Action,
	{_State356, UnderscoreToken}:                   _GotoState257Action,
	{_State356, ParameterDefType}:                  _GotoState279Action,
	{_State356, ProperParameterDefsType}:           _GotoState281Action,
	{_State356, ParameterDefsType}:                 _GotoState424Action,
	{_State357, IdentifierToken}:                   _GotoState113Action,
	{_State357, UnderscoreToken}:                   _GotoState119Action,
	{_State357, StructToken}:                       _GotoState50Action,
	{_State357, EnumToken}:                         _GotoState110Action,
	{_State357, TraitToken}:                        _GotoState118Action,
	{_State357, FuncToken}:                         _GotoState112Action,
	{_State357, LparenToken}:                       _GotoState114Action,
	{_State357, LbracketToken}:                     _GotoState42Action,
	{_State357, DotToken}:                          _GotoState109Action,
	{_State357, QuestionToken}:                     _GotoState116Action,
	{_State357, ExclaimToken}:                      _GotoState111Action,
	{_State357, TildeTildeToken}:                   _GotoState117Action,
	{_State357, BitNegToken}:                       _GotoState108Action,
	{_State357, BitAndToken}:                       _GotoState107Action,
	{_State357, ParseErrorToken}:                   _GotoState115Action,
	{_State357, InitializableTypeExprType}:         _GotoState127Action,
	{_State357, SliceTypeExprType}:                 _GotoState101Action,
	{_State357, ArrayTypeExprType}:                 _GotoState60Action,
	{_State357, MapTypeExprType}:                   _GotoState88Action,
	{_State357, AtomTypeExprType}:                  _GotoState120Action,
	{_State357, NamedTypeExprType}:                 _GotoState128Action,
	{_State357, InferredTypeExprType}:              _GotoState126Action,
	{_State357, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State357, ReturnableTypeExprType}:            _GotoState132Action,
	{_State357, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State357, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State357, TypeExprType}:                      _GotoState425Action,
	{_State357, BinaryTypeExprType}:                _GotoState121Action,
	{_State357, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State357, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State357, TraitTypeExprType}:                 _GotoState133Action,
	{_State357, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State357, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State357, FuncTypeExprType}:                  _GotoState123Action,
	{_State358, AddToken}:                          _GotoState249Action,
	{_State358, SubToken}:                          _GotoState251Action,
	{_State358, MulToken}:                          _GotoState250Action,
	{_State358, BinaryTypeOpType}:                  _GotoState252Action,
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
	{_State359, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State359, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State359, TypeExprType}:                      _GotoState426Action,
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
	{_State361, IdentifierToken}:                   _GotoState427Action,
	{_State364, AddToken}:                          _GotoState249Action,
	{_State364, SubToken}:                          _GotoState251Action,
	{_State364, MulToken}:                          _GotoState250Action,
	{_State364, BinaryTypeOpType}:                  _GotoState252Action,
	{_State365, ImplementsToken}:                   _GotoState428Action,
	{_State365, AddToken}:                          _GotoState249Action,
	{_State365, SubToken}:                          _GotoState251Action,
	{_State365, MulToken}:                          _GotoState250Action,
	{_State365, BinaryTypeOpType}:                  _GotoState252Action,
	{_State366, IdentifierToken}:                   _GotoState143Action,
	{_State366, UnderscoreToken}:                   _GotoState145Action,
	{_State366, LparenToken}:                       _GotoState144Action,
	{_State366, VarPatternType}:                    _GotoState429Action,
	{_State366, TuplePatternType}:                  _GotoState146Action,
	{_State367, IdentifierToken}:                   _GotoState267Action,
	{_State367, UnderscoreToken}:                   _GotoState145Action,
	{_State367, LparenToken}:                       _GotoState144Action,
	{_State367, EllipsisToken}:                     _GotoState266Action,
	{_State367, VarPatternType}:                    _GotoState270Action,
	{_State367, TuplePatternType}:                  _GotoState146Action,
	{_State367, FieldVarPatternType}:               _GotoState430Action,
	{_State370, LparenToken}:                       _GotoState144Action,
	{_State370, TuplePatternType}:                  _GotoState431Action,
	{_State373, IdentifierToken}:                   _GotoState113Action,
	{_State373, UnderscoreToken}:                   _GotoState119Action,
	{_State373, StructToken}:                       _GotoState50Action,
	{_State373, EnumToken}:                         _GotoState110Action,
	{_State373, TraitToken}:                        _GotoState118Action,
	{_State373, FuncToken}:                         _GotoState112Action,
	{_State373, LparenToken}:                       _GotoState114Action,
	{_State373, LbracketToken}:                     _GotoState42Action,
	{_State373, DotToken}:                          _GotoState109Action,
	{_State373, QuestionToken}:                     _GotoState116Action,
	{_State373, ExclaimToken}:                      _GotoState111Action,
	{_State373, TildeTildeToken}:                   _GotoState117Action,
	{_State373, BitNegToken}:                       _GotoState108Action,
	{_State373, BitAndToken}:                       _GotoState107Action,
	{_State373, ParseErrorToken}:                   _GotoState115Action,
	{_State373, InitializableTypeExprType}:         _GotoState127Action,
	{_State373, SliceTypeExprType}:                 _GotoState101Action,
	{_State373, ArrayTypeExprType}:                 _GotoState60Action,
	{_State373, MapTypeExprType}:                   _GotoState88Action,
	{_State373, AtomTypeExprType}:                  _GotoState120Action,
	{_State373, NamedTypeExprType}:                 _GotoState128Action,
	{_State373, InferredTypeExprType}:              _GotoState126Action,
	{_State373, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State373, ReturnableTypeExprType}:            _GotoState433Action,
	{_State373, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State373, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State373, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State373, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State373, TraitTypeExprType}:                 _GotoState133Action,
	{_State373, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State373, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State373, ReturnTypeType}:                    _GotoState432Action,
	{_State373, FuncTypeExprType}:                  _GotoState123Action,
	{_State374, IdentifierToken}:                   _GotoState256Action,
	{_State374, UnderscoreToken}:                   _GotoState257Action,
	{_State374, ParameterDefType}:                  _GotoState434Action,
	{_State376, StringLiteralToken}:                _GotoState161Action,
	{_State376, IdentifierToken}:                   _GotoState159Action,
	{_State376, UnderscoreToken}:                   _GotoState162Action,
	{_State376, DotToken}:                          _GotoState158Action,
	{_State376, ImportClauseType}:                  _GotoState435Action,
	{_State377, StringLiteralToken}:                _GotoState161Action,
	{_State377, IdentifierToken}:                   _GotoState159Action,
	{_State377, UnderscoreToken}:                   _GotoState162Action,
	{_State377, DotToken}:                          _GotoState158Action,
	{_State377, ImportClauseType}:                  _GotoState436Action,
	{_State378, RbracketToken}:                     _GotoState437Action,
	{_State378, AddToken}:                          _GotoState249Action,
	{_State378, SubToken}:                          _GotoState251Action,
	{_State378, MulToken}:                          _GotoState250Action,
	{_State378, BinaryTypeOpType}:                  _GotoState252Action,
	{_State379, RbracketToken}:                     _GotoState438Action,
	{_State385, IdentifierToken}:                   _GotoState237Action,
	{_State385, UnderscoreToken}:                   _GotoState238Action,
	{_State385, UnsafeToken}:                       _GotoState54Action,
	{_State385, StructToken}:                       _GotoState50Action,
	{_State385, EnumToken}:                         _GotoState110Action,
	{_State385, TraitToken}:                        _GotoState118Action,
	{_State385, FuncToken}:                         _GotoState236Action,
	{_State385, LparenToken}:                       _GotoState114Action,
	{_State385, LbracketToken}:                     _GotoState42Action,
	{_State385, DotToken}:                          _GotoState109Action,
	{_State385, QuestionToken}:                     _GotoState116Action,
	{_State385, ExclaimToken}:                      _GotoState111Action,
	{_State385, TildeTildeToken}:                   _GotoState117Action,
	{_State385, BitNegToken}:                       _GotoState108Action,
	{_State385, BitAndToken}:                       _GotoState107Action,
	{_State385, ParseErrorToken}:                   _GotoState115Action,
	{_State385, UnsafeStatementType}:               _GotoState246Action,
	{_State385, InitializableTypeExprType}:         _GotoState127Action,
	{_State385, SliceTypeExprType}:                 _GotoState101Action,
	{_State385, ArrayTypeExprType}:                 _GotoState60Action,
	{_State385, MapTypeExprType}:                   _GotoState88Action,
	{_State385, AtomTypeExprType}:                  _GotoState120Action,
	{_State385, NamedTypeExprType}:                 _GotoState128Action,
	{_State385, InferredTypeExprType}:              _GotoState126Action,
	{_State385, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State385, ReturnableTypeExprType}:            _GotoState132Action,
	{_State385, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State385, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State385, TypeExprType}:                      _GotoState245Action,
	{_State385, BinaryTypeExprType}:                _GotoState121Action,
	{_State385, FieldDefType}:                      _GotoState439Action,
	{_State385, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State385, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State385, TraitTypeExprType}:                 _GotoState133Action,
	{_State385, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State385, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State385, FuncTypeExprType}:                  _GotoState123Action,
	{_State385, MethodSignatureType}:               _GotoState242Action,
	{_State386, IdentifierToken}:                   _GotoState237Action,
	{_State386, UnderscoreToken}:                   _GotoState238Action,
	{_State386, UnsafeToken}:                       _GotoState54Action,
	{_State386, StructToken}:                       _GotoState50Action,
	{_State386, EnumToken}:                         _GotoState110Action,
	{_State386, TraitToken}:                        _GotoState118Action,
	{_State386, FuncToken}:                         _GotoState236Action,
	{_State386, LparenToken}:                       _GotoState114Action,
	{_State386, LbracketToken}:                     _GotoState42Action,
	{_State386, DotToken}:                          _GotoState109Action,
	{_State386, QuestionToken}:                     _GotoState116Action,
	{_State386, ExclaimToken}:                      _GotoState111Action,
	{_State386, TildeTildeToken}:                   _GotoState117Action,
	{_State386, BitNegToken}:                       _GotoState108Action,
	{_State386, BitAndToken}:                       _GotoState107Action,
	{_State386, ParseErrorToken}:                   _GotoState115Action,
	{_State386, UnsafeStatementType}:               _GotoState246Action,
	{_State386, InitializableTypeExprType}:         _GotoState127Action,
	{_State386, SliceTypeExprType}:                 _GotoState101Action,
	{_State386, ArrayTypeExprType}:                 _GotoState60Action,
	{_State386, MapTypeExprType}:                   _GotoState88Action,
	{_State386, AtomTypeExprType}:                  _GotoState120Action,
	{_State386, NamedTypeExprType}:                 _GotoState128Action,
	{_State386, InferredTypeExprType}:              _GotoState126Action,
	{_State386, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State386, ReturnableTypeExprType}:            _GotoState132Action,
	{_State386, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State386, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State386, TypeExprType}:                      _GotoState245Action,
	{_State386, BinaryTypeExprType}:                _GotoState121Action,
	{_State386, FieldDefType}:                      _GotoState440Action,
	{_State386, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State386, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State386, TraitTypeExprType}:                 _GotoState133Action,
	{_State386, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State386, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State386, FuncTypeExprType}:                  _GotoState123Action,
	{_State386, MethodSignatureType}:               _GotoState242Action,
	{_State387, StringLiteralToken}:                _GotoState441Action,
	{_State388, IdentifierToken}:                   _GotoState113Action,
	{_State388, UnderscoreToken}:                   _GotoState119Action,
	{_State388, StructToken}:                       _GotoState50Action,
	{_State388, EnumToken}:                         _GotoState110Action,
	{_State388, TraitToken}:                        _GotoState118Action,
	{_State388, FuncToken}:                         _GotoState112Action,
	{_State388, LparenToken}:                       _GotoState114Action,
	{_State388, LbracketToken}:                     _GotoState42Action,
	{_State388, DotToken}:                          _GotoState109Action,
	{_State388, QuestionToken}:                     _GotoState116Action,
	{_State388, ExclaimToken}:                      _GotoState111Action,
	{_State388, TildeTildeToken}:                   _GotoState117Action,
	{_State388, BitNegToken}:                       _GotoState108Action,
	{_State388, BitAndToken}:                       _GotoState107Action,
	{_State388, ParseErrorToken}:                   _GotoState115Action,
	{_State388, InitializableTypeExprType}:         _GotoState127Action,
	{_State388, SliceTypeExprType}:                 _GotoState101Action,
	{_State388, ArrayTypeExprType}:                 _GotoState60Action,
	{_State388, MapTypeExprType}:                   _GotoState88Action,
	{_State388, AtomTypeExprType}:                  _GotoState120Action,
	{_State388, NamedTypeExprType}:                 _GotoState128Action,
	{_State388, InferredTypeExprType}:              _GotoState126Action,
	{_State388, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State388, ReturnableTypeExprType}:            _GotoState132Action,
	{_State388, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State388, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State388, TypeExprType}:                      _GotoState442Action,
	{_State388, BinaryTypeExprType}:                _GotoState121Action,
	{_State388, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State388, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State388, TraitTypeExprType}:                 _GotoState133Action,
	{_State388, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State388, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State388, FuncTypeExprType}:                  _GotoState123Action,
	{_State391, RparenToken}:                       _GotoState443Action,
	{_State393, IntegerLiteralToken}:               _GotoState40Action,
	{_State393, FloatLiteralToken}:                 _GotoState35Action,
	{_State393, RuneLiteralToken}:                  _GotoState48Action,
	{_State393, StringLiteralToken}:                _GotoState49Action,
	{_State393, IdentifierToken}:                   _GotoState38Action,
	{_State393, UnderscoreToken}:                   _GotoState53Action,
	{_State393, TrueToken}:                         _GotoState52Action,
	{_State393, FalseToken}:                        _GotoState34Action,
	{_State393, StructToken}:                       _GotoState50Action,
	{_State393, FuncToken}:                         _GotoState36Action,
	{_State393, VarToken}:                          _GotoState15Action,
	{_State393, LetToken}:                          _GotoState12Action,
	{_State393, NotToken}:                          _GotoState45Action,
	{_State393, LabelDeclToken}:                    _GotoState41Action,
	{_State393, LparenToken}:                       _GotoState43Action,
	{_State393, LbracketToken}:                     _GotoState42Action,
	{_State393, SubToken}:                          _GotoState51Action,
	{_State393, MulToken}:                          _GotoState44Action,
	{_State393, BitNegToken}:                       _GotoState27Action,
	{_State393, BitAndToken}:                       _GotoState26Action,
	{_State393, GreaterToken}:                      _GotoState37Action,
	{_State393, ParseErrorToken}:                   _GotoState46Action,
	{_State393, VarDeclPatternType}:                _GotoState104Action,
	{_State393, VarOrLetType}:                      _GotoState24Action,
	{_State393, OptionalLabelDeclType}:             _GotoState153Action,
	{_State393, SequenceExprType}:                  _GotoState444Action,
	{_State393, CallExprType}:                      _GotoState71Action,
	{_State393, AtomExprType}:                      _GotoState63Action,
	{_State393, ParseErrorExprType}:                _GotoState93Action,
	{_State393, LiteralExprType}:                   _GotoState87Action,
	{_State393, NamedExprType}:                     _GotoState90Action,
	{_State393, BlockExprType}:                     _GotoState70Action,
	{_State393, InitializeExprType}:                _GotoState84Action,
	{_State393, ImplicitStructExprType}:            _GotoState80Action,
	{_State393, AccessibleExprType}:                _GotoState105Action,
	{_State393, AccessExprType}:                    _GotoState55Action,
	{_State393, IndexExprType}:                     _GotoState82Action,
	{_State393, PostfixableExprType}:               _GotoState95Action,
	{_State393, PostfixUnaryExprType}:              _GotoState94Action,
	{_State393, PrefixableExprType}:                _GotoState98Action,
	{_State393, PrefixUnaryExprType}:               _GotoState96Action,
	{_State393, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State393, MulExprType}:                       _GotoState89Action,
	{_State393, BinaryMulExprType}:                 _GotoState67Action,
	{_State393, AddExprType}:                       _GotoState57Action,
	{_State393, BinaryAddExprType}:                 _GotoState64Action,
	{_State393, CmpExprType}:                       _GotoState74Action,
	{_State393, BinaryCmpExprType}:                 _GotoState66Action,
	{_State393, AndExprType}:                       _GotoState58Action,
	{_State393, BinaryAndExprType}:                 _GotoState65Action,
	{_State393, OrExprType}:                        _GotoState92Action,
	{_State393, BinaryOrExprType}:                  _GotoState69Action,
	{_State393, InitializableTypeExprType}:         _GotoState83Action,
	{_State393, SliceTypeExprType}:                 _GotoState101Action,
	{_State393, ArrayTypeExprType}:                 _GotoState60Action,
	{_State393, MapTypeExprType}:                   _GotoState88Action,
	{_State393, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State393, AnonymousFuncExprType}:             _GotoState59Action,
	{_State394, IntegerLiteralToken}:               _GotoState40Action,
	{_State394, FloatLiteralToken}:                 _GotoState35Action,
	{_State394, RuneLiteralToken}:                  _GotoState48Action,
	{_State394, StringLiteralToken}:                _GotoState49Action,
	{_State394, IdentifierToken}:                   _GotoState38Action,
	{_State394, UnderscoreToken}:                   _GotoState53Action,
	{_State394, TrueToken}:                         _GotoState52Action,
	{_State394, FalseToken}:                        _GotoState34Action,
	{_State394, StructToken}:                       _GotoState50Action,
	{_State394, FuncToken}:                         _GotoState36Action,
	{_State394, VarToken}:                          _GotoState15Action,
	{_State394, LetToken}:                          _GotoState12Action,
	{_State394, NotToken}:                          _GotoState45Action,
	{_State394, LabelDeclToken}:                    _GotoState41Action,
	{_State394, LparenToken}:                       _GotoState43Action,
	{_State394, LbracketToken}:                     _GotoState42Action,
	{_State394, SubToken}:                          _GotoState51Action,
	{_State394, MulToken}:                          _GotoState44Action,
	{_State394, BitNegToken}:                       _GotoState27Action,
	{_State394, BitAndToken}:                       _GotoState26Action,
	{_State394, GreaterToken}:                      _GotoState37Action,
	{_State394, ParseErrorToken}:                   _GotoState46Action,
	{_State394, VarDeclPatternType}:                _GotoState104Action,
	{_State394, VarOrLetType}:                      _GotoState24Action,
	{_State394, OptionalLabelDeclType}:             _GotoState153Action,
	{_State394, SequenceExprType}:                  _GotoState445Action,
	{_State394, CallExprType}:                      _GotoState71Action,
	{_State394, AtomExprType}:                      _GotoState63Action,
	{_State394, ParseErrorExprType}:                _GotoState93Action,
	{_State394, LiteralExprType}:                   _GotoState87Action,
	{_State394, NamedExprType}:                     _GotoState90Action,
	{_State394, BlockExprType}:                     _GotoState70Action,
	{_State394, InitializeExprType}:                _GotoState84Action,
	{_State394, ImplicitStructExprType}:            _GotoState80Action,
	{_State394, AccessibleExprType}:                _GotoState105Action,
	{_State394, AccessExprType}:                    _GotoState55Action,
	{_State394, IndexExprType}:                     _GotoState82Action,
	{_State394, PostfixableExprType}:               _GotoState95Action,
	{_State394, PostfixUnaryExprType}:              _GotoState94Action,
	{_State394, PrefixableExprType}:                _GotoState98Action,
	{_State394, PrefixUnaryExprType}:               _GotoState96Action,
	{_State394, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State394, MulExprType}:                       _GotoState89Action,
	{_State394, BinaryMulExprType}:                 _GotoState67Action,
	{_State394, AddExprType}:                       _GotoState57Action,
	{_State394, BinaryAddExprType}:                 _GotoState64Action,
	{_State394, CmpExprType}:                       _GotoState74Action,
	{_State394, BinaryCmpExprType}:                 _GotoState66Action,
	{_State394, AndExprType}:                       _GotoState58Action,
	{_State394, BinaryAndExprType}:                 _GotoState65Action,
	{_State394, OrExprType}:                        _GotoState92Action,
	{_State394, BinaryOrExprType}:                  _GotoState69Action,
	{_State394, InitializableTypeExprType}:         _GotoState83Action,
	{_State394, SliceTypeExprType}:                 _GotoState101Action,
	{_State394, ArrayTypeExprType}:                 _GotoState60Action,
	{_State394, MapTypeExprType}:                   _GotoState88Action,
	{_State394, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State394, AnonymousFuncExprType}:             _GotoState59Action,
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
	{_State395, SequenceExprType}:                  _GotoState446Action,
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
	{_State396, SequenceExprType}:                  _GotoState448Action,
	{_State396, OptionalSequenceExprType}:          _GotoState447Action,
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
	{_State397, LbraceToken}:                       _GotoState11Action,
	{_State397, StatementBlockType}:                _GotoState449Action,
	{_State398, CommaToken}:                        _GotoState276Action,
	{_State398, AssignToken}:                       _GotoState450Action,
	{_State399, ElseToken}:                         _GotoState451Action,
	{_State402, IdentifierToken}:                   _GotoState237Action,
	{_State402, UnderscoreToken}:                   _GotoState238Action,
	{_State402, UnsafeToken}:                       _GotoState54Action,
	{_State402, StructToken}:                       _GotoState50Action,
	{_State402, EnumToken}:                         _GotoState110Action,
	{_State402, TraitToken}:                        _GotoState118Action,
	{_State402, FuncToken}:                         _GotoState236Action,
	{_State402, LparenToken}:                       _GotoState114Action,
	{_State402, LbracketToken}:                     _GotoState42Action,
	{_State402, DotToken}:                          _GotoState109Action,
	{_State402, QuestionToken}:                     _GotoState116Action,
	{_State402, ExclaimToken}:                      _GotoState111Action,
	{_State402, TildeTildeToken}:                   _GotoState117Action,
	{_State402, BitNegToken}:                       _GotoState108Action,
	{_State402, BitAndToken}:                       _GotoState107Action,
	{_State402, ParseErrorToken}:                   _GotoState115Action,
	{_State402, UnsafeStatementType}:               _GotoState246Action,
	{_State402, InitializableTypeExprType}:         _GotoState127Action,
	{_State402, SliceTypeExprType}:                 _GotoState101Action,
	{_State402, ArrayTypeExprType}:                 _GotoState60Action,
	{_State402, MapTypeExprType}:                   _GotoState88Action,
	{_State402, AtomTypeExprType}:                  _GotoState120Action,
	{_State402, NamedTypeExprType}:                 _GotoState128Action,
	{_State402, InferredTypeExprType}:              _GotoState126Action,
	{_State402, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State402, ReturnableTypeExprType}:            _GotoState132Action,
	{_State402, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State402, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State402, TypeExprType}:                      _GotoState245Action,
	{_State402, BinaryTypeExprType}:                _GotoState121Action,
	{_State402, FieldDefType}:                      _GotoState452Action,
	{_State402, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State402, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State402, TraitTypeExprType}:                 _GotoState133Action,
	{_State402, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State402, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State402, FuncTypeExprType}:                  _GotoState123Action,
	{_State402, MethodSignatureType}:               _GotoState242Action,
	{_State403, IdentifierToken}:                   _GotoState237Action,
	{_State403, UnderscoreToken}:                   _GotoState238Action,
	{_State403, UnsafeToken}:                       _GotoState54Action,
	{_State403, StructToken}:                       _GotoState50Action,
	{_State403, EnumToken}:                         _GotoState110Action,
	{_State403, TraitToken}:                        _GotoState118Action,
	{_State403, FuncToken}:                         _GotoState236Action,
	{_State403, LparenToken}:                       _GotoState114Action,
	{_State403, LbracketToken}:                     _GotoState42Action,
	{_State403, DotToken}:                          _GotoState109Action,
	{_State403, QuestionToken}:                     _GotoState116Action,
	{_State403, ExclaimToken}:                      _GotoState111Action,
	{_State403, TildeTildeToken}:                   _GotoState117Action,
	{_State403, BitNegToken}:                       _GotoState108Action,
	{_State403, BitAndToken}:                       _GotoState107Action,
	{_State403, ParseErrorToken}:                   _GotoState115Action,
	{_State403, UnsafeStatementType}:               _GotoState246Action,
	{_State403, InitializableTypeExprType}:         _GotoState127Action,
	{_State403, SliceTypeExprType}:                 _GotoState101Action,
	{_State403, ArrayTypeExprType}:                 _GotoState60Action,
	{_State403, MapTypeExprType}:                   _GotoState88Action,
	{_State403, AtomTypeExprType}:                  _GotoState120Action,
	{_State403, NamedTypeExprType}:                 _GotoState128Action,
	{_State403, InferredTypeExprType}:              _GotoState126Action,
	{_State403, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State403, ReturnableTypeExprType}:            _GotoState132Action,
	{_State403, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State403, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State403, TypeExprType}:                      _GotoState245Action,
	{_State403, BinaryTypeExprType}:                _GotoState121Action,
	{_State403, FieldDefType}:                      _GotoState453Action,
	{_State403, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State403, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State403, TraitTypeExprType}:                 _GotoState133Action,
	{_State403, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State403, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State403, FuncTypeExprType}:                  _GotoState123Action,
	{_State403, MethodSignatureType}:               _GotoState242Action,
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
	{_State404, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State404, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State404, TypeExprType}:                      _GotoState245Action,
	{_State404, BinaryTypeExprType}:                _GotoState121Action,
	{_State404, FieldDefType}:                      _GotoState454Action,
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
	{_State405, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State405, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State405, TypeExprType}:                      _GotoState245Action,
	{_State405, BinaryTypeExprType}:                _GotoState121Action,
	{_State405, FieldDefType}:                      _GotoState455Action,
	{_State405, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State405, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State405, TraitTypeExprType}:                 _GotoState133Action,
	{_State405, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State405, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State405, FuncTypeExprType}:                  _GotoState123Action,
	{_State405, MethodSignatureType}:               _GotoState242Action,
	{_State406, AddToken}:                          _GotoState249Action,
	{_State406, SubToken}:                          _GotoState251Action,
	{_State406, MulToken}:                          _GotoState250Action,
	{_State406, BinaryTypeOpType}:                  _GotoState252Action,
	{_State407, IdentifierToken}:                   _GotoState113Action,
	{_State407, UnderscoreToken}:                   _GotoState119Action,
	{_State407, StructToken}:                       _GotoState50Action,
	{_State407, EnumToken}:                         _GotoState110Action,
	{_State407, TraitToken}:                        _GotoState118Action,
	{_State407, FuncToken}:                         _GotoState112Action,
	{_State407, LparenToken}:                       _GotoState114Action,
	{_State407, LbracketToken}:                     _GotoState42Action,
	{_State407, DotToken}:                          _GotoState109Action,
	{_State407, QuestionToken}:                     _GotoState116Action,
	{_State407, ExclaimToken}:                      _GotoState111Action,
	{_State407, TildeTildeToken}:                   _GotoState117Action,
	{_State407, BitNegToken}:                       _GotoState108Action,
	{_State407, BitAndToken}:                       _GotoState107Action,
	{_State407, ParseErrorToken}:                   _GotoState115Action,
	{_State407, InitializableTypeExprType}:         _GotoState127Action,
	{_State407, SliceTypeExprType}:                 _GotoState101Action,
	{_State407, ArrayTypeExprType}:                 _GotoState60Action,
	{_State407, MapTypeExprType}:                   _GotoState88Action,
	{_State407, AtomTypeExprType}:                  _GotoState120Action,
	{_State407, NamedTypeExprType}:                 _GotoState128Action,
	{_State407, InferredTypeExprType}:              _GotoState126Action,
	{_State407, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State407, ReturnableTypeExprType}:            _GotoState132Action,
	{_State407, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State407, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State407, TypeExprType}:                      _GotoState456Action,
	{_State407, BinaryTypeExprType}:                _GotoState121Action,
	{_State407, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State407, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State407, TraitTypeExprType}:                 _GotoState133Action,
	{_State407, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State407, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State407, FuncTypeExprType}:                  _GotoState123Action,
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
	{_State409, ReturnableTypeExprType}:            _GotoState132Action,
	{_State409, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State409, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State409, TypeExprType}:                      _GotoState457Action,
	{_State409, BinaryTypeExprType}:                _GotoState121Action,
	{_State409, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State409, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State409, TraitTypeExprType}:                 _GotoState133Action,
	{_State409, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State409, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State409, FuncTypeExprType}:                  _GotoState123Action,
	{_State410, AddToken}:                          _GotoState249Action,
	{_State410, SubToken}:                          _GotoState251Action,
	{_State410, MulToken}:                          _GotoState250Action,
	{_State410, BinaryTypeOpType}:                  _GotoState252Action,
	{_State411, IdentifierToken}:                   _GotoState113Action,
	{_State411, UnderscoreToken}:                   _GotoState119Action,
	{_State411, StructToken}:                       _GotoState50Action,
	{_State411, EnumToken}:                         _GotoState110Action,
	{_State411, TraitToken}:                        _GotoState118Action,
	{_State411, FuncToken}:                         _GotoState112Action,
	{_State411, LparenToken}:                       _GotoState114Action,
	{_State411, LbracketToken}:                     _GotoState42Action,
	{_State411, DotToken}:                          _GotoState109Action,
	{_State411, QuestionToken}:                     _GotoState116Action,
	{_State411, ExclaimToken}:                      _GotoState111Action,
	{_State411, TildeTildeToken}:                   _GotoState117Action,
	{_State411, BitNegToken}:                       _GotoState108Action,
	{_State411, BitAndToken}:                       _GotoState107Action,
	{_State411, ParseErrorToken}:                   _GotoState115Action,
	{_State411, InitializableTypeExprType}:         _GotoState127Action,
	{_State411, SliceTypeExprType}:                 _GotoState101Action,
	{_State411, ArrayTypeExprType}:                 _GotoState60Action,
	{_State411, MapTypeExprType}:                   _GotoState88Action,
	{_State411, AtomTypeExprType}:                  _GotoState120Action,
	{_State411, NamedTypeExprType}:                 _GotoState128Action,
	{_State411, InferredTypeExprType}:              _GotoState126Action,
	{_State411, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State411, ReturnableTypeExprType}:            _GotoState433Action,
	{_State411, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State411, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State411, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State411, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State411, TraitTypeExprType}:                 _GotoState133Action,
	{_State411, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State411, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State411, ReturnTypeType}:                    _GotoState458Action,
	{_State411, FuncTypeExprType}:                  _GotoState123Action,
	{_State412, IdentifierToken}:                   _GotoState330Action,
	{_State412, UnderscoreToken}:                   _GotoState331Action,
	{_State412, StructToken}:                       _GotoState50Action,
	{_State412, EnumToken}:                         _GotoState110Action,
	{_State412, TraitToken}:                        _GotoState118Action,
	{_State412, FuncToken}:                         _GotoState112Action,
	{_State412, LparenToken}:                       _GotoState114Action,
	{_State412, LbracketToken}:                     _GotoState42Action,
	{_State412, DotToken}:                          _GotoState109Action,
	{_State412, QuestionToken}:                     _GotoState116Action,
	{_State412, ExclaimToken}:                      _GotoState111Action,
	{_State412, EllipsisToken}:                     _GotoState329Action,
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
	{_State412, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State412, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State412, TypeExprType}:                      _GotoState335Action,
	{_State412, BinaryTypeExprType}:                _GotoState121Action,
	{_State412, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State412, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State412, TraitTypeExprType}:                 _GotoState133Action,
	{_State412, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State412, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State412, ParameterDeclType}:                 _GotoState459Action,
	{_State412, FuncTypeExprType}:                  _GotoState123Action,
	{_State414, IdentifierToken}:                   _GotoState330Action,
	{_State414, UnderscoreToken}:                   _GotoState331Action,
	{_State414, StructToken}:                       _GotoState50Action,
	{_State414, EnumToken}:                         _GotoState110Action,
	{_State414, TraitToken}:                        _GotoState118Action,
	{_State414, FuncToken}:                         _GotoState112Action,
	{_State414, LparenToken}:                       _GotoState114Action,
	{_State414, LbracketToken}:                     _GotoState42Action,
	{_State414, DotToken}:                          _GotoState109Action,
	{_State414, QuestionToken}:                     _GotoState116Action,
	{_State414, ExclaimToken}:                      _GotoState111Action,
	{_State414, EllipsisToken}:                     _GotoState329Action,
	{_State414, TildeTildeToken}:                   _GotoState117Action,
	{_State414, BitNegToken}:                       _GotoState108Action,
	{_State414, BitAndToken}:                       _GotoState107Action,
	{_State414, ParseErrorToken}:                   _GotoState115Action,
	{_State414, InitializableTypeExprType}:         _GotoState127Action,
	{_State414, SliceTypeExprType}:                 _GotoState101Action,
	{_State414, ArrayTypeExprType}:                 _GotoState60Action,
	{_State414, MapTypeExprType}:                   _GotoState88Action,
	{_State414, AtomTypeExprType}:                  _GotoState120Action,
	{_State414, NamedTypeExprType}:                 _GotoState128Action,
	{_State414, InferredTypeExprType}:              _GotoState126Action,
	{_State414, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State414, ReturnableTypeExprType}:            _GotoState132Action,
	{_State414, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State414, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State414, TypeExprType}:                      _GotoState335Action,
	{_State414, BinaryTypeExprType}:                _GotoState121Action,
	{_State414, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State414, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State414, TraitTypeExprType}:                 _GotoState133Action,
	{_State414, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State414, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State414, ParameterDeclType}:                 _GotoState332Action,
	{_State414, ProperParameterDeclsType}:          _GotoState334Action,
	{_State414, ParameterDeclsType}:                _GotoState460Action,
	{_State414, FuncTypeExprType}:                  _GotoState123Action,
	{_State421, AddToken}:                          _GotoState249Action,
	{_State421, SubToken}:                          _GotoState251Action,
	{_State421, MulToken}:                          _GotoState250Action,
	{_State421, BinaryTypeOpType}:                  _GotoState252Action,
	{_State423, IdentifierToken}:                   _GotoState352Action,
	{_State423, GenericParameterDefType}:           _GotoState461Action,
	{_State424, RparenToken}:                       _GotoState462Action,
	{_State425, AddToken}:                          _GotoState249Action,
	{_State425, SubToken}:                          _GotoState251Action,
	{_State425, MulToken}:                          _GotoState250Action,
	{_State425, BinaryTypeOpType}:                  _GotoState252Action,
	{_State426, AddToken}:                          _GotoState249Action,
	{_State426, SubToken}:                          _GotoState251Action,
	{_State426, MulToken}:                          _GotoState250Action,
	{_State426, BinaryTypeOpType}:                  _GotoState252Action,
	{_State427, DollarLbracketToken}:               _GotoState254Action,
	{_State427, OptionalGenericParametersType}:     _GotoState463Action,
	{_State428, IdentifierToken}:                   _GotoState113Action,
	{_State428, UnderscoreToken}:                   _GotoState119Action,
	{_State428, StructToken}:                       _GotoState50Action,
	{_State428, EnumToken}:                         _GotoState110Action,
	{_State428, TraitToken}:                        _GotoState118Action,
	{_State428, FuncToken}:                         _GotoState112Action,
	{_State428, LparenToken}:                       _GotoState114Action,
	{_State428, LbracketToken}:                     _GotoState42Action,
	{_State428, DotToken}:                          _GotoState109Action,
	{_State428, QuestionToken}:                     _GotoState116Action,
	{_State428, ExclaimToken}:                      _GotoState111Action,
	{_State428, TildeTildeToken}:                   _GotoState117Action,
	{_State428, BitNegToken}:                       _GotoState108Action,
	{_State428, BitAndToken}:                       _GotoState107Action,
	{_State428, ParseErrorToken}:                   _GotoState115Action,
	{_State428, InitializableTypeExprType}:         _GotoState127Action,
	{_State428, SliceTypeExprType}:                 _GotoState101Action,
	{_State428, ArrayTypeExprType}:                 _GotoState60Action,
	{_State428, MapTypeExprType}:                   _GotoState88Action,
	{_State428, AtomTypeExprType}:                  _GotoState120Action,
	{_State428, NamedTypeExprType}:                 _GotoState128Action,
	{_State428, InferredTypeExprType}:              _GotoState126Action,
	{_State428, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State428, ReturnableTypeExprType}:            _GotoState132Action,
	{_State428, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State428, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State428, TypeExprType}:                      _GotoState464Action,
	{_State428, BinaryTypeExprType}:                _GotoState121Action,
	{_State428, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State428, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State428, TraitTypeExprType}:                 _GotoState133Action,
	{_State428, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State428, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State428, FuncTypeExprType}:                  _GotoState123Action,
	{_State432, LbraceToken}:                       _GotoState11Action,
	{_State432, StatementBlockType}:                _GotoState465Action,
	{_State442, AddToken}:                          _GotoState249Action,
	{_State442, SubToken}:                          _GotoState251Action,
	{_State442, MulToken}:                          _GotoState250Action,
	{_State442, BinaryTypeOpType}:                  _GotoState252Action,
	{_State446, DoToken}:                           _GotoState466Action,
	{_State447, SemicolonToken}:                    _GotoState467Action,
	{_State450, IntegerLiteralToken}:               _GotoState40Action,
	{_State450, FloatLiteralToken}:                 _GotoState35Action,
	{_State450, RuneLiteralToken}:                  _GotoState48Action,
	{_State450, StringLiteralToken}:                _GotoState49Action,
	{_State450, IdentifierToken}:                   _GotoState38Action,
	{_State450, UnderscoreToken}:                   _GotoState53Action,
	{_State450, TrueToken}:                         _GotoState52Action,
	{_State450, FalseToken}:                        _GotoState34Action,
	{_State450, StructToken}:                       _GotoState50Action,
	{_State450, FuncToken}:                         _GotoState36Action,
	{_State450, VarToken}:                          _GotoState15Action,
	{_State450, LetToken}:                          _GotoState12Action,
	{_State450, NotToken}:                          _GotoState45Action,
	{_State450, LabelDeclToken}:                    _GotoState41Action,
	{_State450, LparenToken}:                       _GotoState43Action,
	{_State450, LbracketToken}:                     _GotoState42Action,
	{_State450, SubToken}:                          _GotoState51Action,
	{_State450, MulToken}:                          _GotoState44Action,
	{_State450, BitNegToken}:                       _GotoState27Action,
	{_State450, BitAndToken}:                       _GotoState26Action,
	{_State450, GreaterToken}:                      _GotoState37Action,
	{_State450, ParseErrorToken}:                   _GotoState46Action,
	{_State450, VarDeclPatternType}:                _GotoState104Action,
	{_State450, VarOrLetType}:                      _GotoState24Action,
	{_State450, OptionalLabelDeclType}:             _GotoState153Action,
	{_State450, SequenceExprType}:                  _GotoState468Action,
	{_State450, CallExprType}:                      _GotoState71Action,
	{_State450, AtomExprType}:                      _GotoState63Action,
	{_State450, ParseErrorExprType}:                _GotoState93Action,
	{_State450, LiteralExprType}:                   _GotoState87Action,
	{_State450, NamedExprType}:                     _GotoState90Action,
	{_State450, BlockExprType}:                     _GotoState70Action,
	{_State450, InitializeExprType}:                _GotoState84Action,
	{_State450, ImplicitStructExprType}:            _GotoState80Action,
	{_State450, AccessibleExprType}:                _GotoState105Action,
	{_State450, AccessExprType}:                    _GotoState55Action,
	{_State450, IndexExprType}:                     _GotoState82Action,
	{_State450, PostfixableExprType}:               _GotoState95Action,
	{_State450, PostfixUnaryExprType}:              _GotoState94Action,
	{_State450, PrefixableExprType}:                _GotoState98Action,
	{_State450, PrefixUnaryExprType}:               _GotoState96Action,
	{_State450, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State450, MulExprType}:                       _GotoState89Action,
	{_State450, BinaryMulExprType}:                 _GotoState67Action,
	{_State450, AddExprType}:                       _GotoState57Action,
	{_State450, BinaryAddExprType}:                 _GotoState64Action,
	{_State450, CmpExprType}:                       _GotoState74Action,
	{_State450, BinaryCmpExprType}:                 _GotoState66Action,
	{_State450, AndExprType}:                       _GotoState58Action,
	{_State450, BinaryAndExprType}:                 _GotoState65Action,
	{_State450, OrExprType}:                        _GotoState92Action,
	{_State450, BinaryOrExprType}:                  _GotoState69Action,
	{_State450, InitializableTypeExprType}:         _GotoState83Action,
	{_State450, SliceTypeExprType}:                 _GotoState101Action,
	{_State450, ArrayTypeExprType}:                 _GotoState60Action,
	{_State450, MapTypeExprType}:                   _GotoState88Action,
	{_State450, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State450, AnonymousFuncExprType}:             _GotoState59Action,
	{_State451, IfToken}:                           _GotoState224Action,
	{_State451, LbraceToken}:                       _GotoState11Action,
	{_State451, StatementBlockType}:                _GotoState470Action,
	{_State451, IfExprType}:                        _GotoState469Action,
	{_State456, AddToken}:                          _GotoState249Action,
	{_State456, SubToken}:                          _GotoState251Action,
	{_State456, MulToken}:                          _GotoState250Action,
	{_State456, BinaryTypeOpType}:                  _GotoState252Action,
	{_State457, AddToken}:                          _GotoState249Action,
	{_State457, SubToken}:                          _GotoState251Action,
	{_State457, MulToken}:                          _GotoState250Action,
	{_State457, BinaryTypeOpType}:                  _GotoState252Action,
	{_State460, RparenToken}:                       _GotoState471Action,
	{_State462, IdentifierToken}:                   _GotoState113Action,
	{_State462, UnderscoreToken}:                   _GotoState119Action,
	{_State462, StructToken}:                       _GotoState50Action,
	{_State462, EnumToken}:                         _GotoState110Action,
	{_State462, TraitToken}:                        _GotoState118Action,
	{_State462, FuncToken}:                         _GotoState112Action,
	{_State462, LparenToken}:                       _GotoState114Action,
	{_State462, LbracketToken}:                     _GotoState42Action,
	{_State462, DotToken}:                          _GotoState109Action,
	{_State462, QuestionToken}:                     _GotoState116Action,
	{_State462, ExclaimToken}:                      _GotoState111Action,
	{_State462, TildeTildeToken}:                   _GotoState117Action,
	{_State462, BitNegToken}:                       _GotoState108Action,
	{_State462, BitAndToken}:                       _GotoState107Action,
	{_State462, ParseErrorToken}:                   _GotoState115Action,
	{_State462, InitializableTypeExprType}:         _GotoState127Action,
	{_State462, SliceTypeExprType}:                 _GotoState101Action,
	{_State462, ArrayTypeExprType}:                 _GotoState60Action,
	{_State462, MapTypeExprType}:                   _GotoState88Action,
	{_State462, AtomTypeExprType}:                  _GotoState120Action,
	{_State462, NamedTypeExprType}:                 _GotoState128Action,
	{_State462, InferredTypeExprType}:              _GotoState126Action,
	{_State462, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State462, ReturnableTypeExprType}:            _GotoState433Action,
	{_State462, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State462, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State462, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State462, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State462, TraitTypeExprType}:                 _GotoState133Action,
	{_State462, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State462, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State462, ReturnTypeType}:                    _GotoState472Action,
	{_State462, FuncTypeExprType}:                  _GotoState123Action,
	{_State463, LparenToken}:                       _GotoState473Action,
	{_State464, AddToken}:                          _GotoState249Action,
	{_State464, SubToken}:                          _GotoState251Action,
	{_State464, MulToken}:                          _GotoState250Action,
	{_State464, BinaryTypeOpType}:                  _GotoState252Action,
	{_State466, LbraceToken}:                       _GotoState11Action,
	{_State466, StatementBlockType}:                _GotoState474Action,
	{_State467, IntegerLiteralToken}:               _GotoState40Action,
	{_State467, FloatLiteralToken}:                 _GotoState35Action,
	{_State467, RuneLiteralToken}:                  _GotoState48Action,
	{_State467, StringLiteralToken}:                _GotoState49Action,
	{_State467, IdentifierToken}:                   _GotoState38Action,
	{_State467, UnderscoreToken}:                   _GotoState53Action,
	{_State467, TrueToken}:                         _GotoState52Action,
	{_State467, FalseToken}:                        _GotoState34Action,
	{_State467, StructToken}:                       _GotoState50Action,
	{_State467, FuncToken}:                         _GotoState36Action,
	{_State467, VarToken}:                          _GotoState15Action,
	{_State467, LetToken}:                          _GotoState12Action,
	{_State467, NotToken}:                          _GotoState45Action,
	{_State467, LabelDeclToken}:                    _GotoState41Action,
	{_State467, LparenToken}:                       _GotoState43Action,
	{_State467, LbracketToken}:                     _GotoState42Action,
	{_State467, SubToken}:                          _GotoState51Action,
	{_State467, MulToken}:                          _GotoState44Action,
	{_State467, BitNegToken}:                       _GotoState27Action,
	{_State467, BitAndToken}:                       _GotoState26Action,
	{_State467, GreaterToken}:                      _GotoState37Action,
	{_State467, ParseErrorToken}:                   _GotoState46Action,
	{_State467, VarDeclPatternType}:                _GotoState104Action,
	{_State467, VarOrLetType}:                      _GotoState24Action,
	{_State467, OptionalLabelDeclType}:             _GotoState153Action,
	{_State467, SequenceExprType}:                  _GotoState448Action,
	{_State467, OptionalSequenceExprType}:          _GotoState475Action,
	{_State467, CallExprType}:                      _GotoState71Action,
	{_State467, AtomExprType}:                      _GotoState63Action,
	{_State467, ParseErrorExprType}:                _GotoState93Action,
	{_State467, LiteralExprType}:                   _GotoState87Action,
	{_State467, NamedExprType}:                     _GotoState90Action,
	{_State467, BlockExprType}:                     _GotoState70Action,
	{_State467, InitializeExprType}:                _GotoState84Action,
	{_State467, ImplicitStructExprType}:            _GotoState80Action,
	{_State467, AccessibleExprType}:                _GotoState105Action,
	{_State467, AccessExprType}:                    _GotoState55Action,
	{_State467, IndexExprType}:                     _GotoState82Action,
	{_State467, PostfixableExprType}:               _GotoState95Action,
	{_State467, PostfixUnaryExprType}:              _GotoState94Action,
	{_State467, PrefixableExprType}:                _GotoState98Action,
	{_State467, PrefixUnaryExprType}:               _GotoState96Action,
	{_State467, PrefixUnaryOpType}:                 _GotoState97Action,
	{_State467, MulExprType}:                       _GotoState89Action,
	{_State467, BinaryMulExprType}:                 _GotoState67Action,
	{_State467, AddExprType}:                       _GotoState57Action,
	{_State467, BinaryAddExprType}:                 _GotoState64Action,
	{_State467, CmpExprType}:                       _GotoState74Action,
	{_State467, BinaryCmpExprType}:                 _GotoState66Action,
	{_State467, AndExprType}:                       _GotoState58Action,
	{_State467, BinaryAndExprType}:                 _GotoState65Action,
	{_State467, OrExprType}:                        _GotoState92Action,
	{_State467, BinaryOrExprType}:                  _GotoState69Action,
	{_State467, InitializableTypeExprType}:         _GotoState83Action,
	{_State467, SliceTypeExprType}:                 _GotoState101Action,
	{_State467, ArrayTypeExprType}:                 _GotoState60Action,
	{_State467, MapTypeExprType}:                   _GotoState88Action,
	{_State467, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State467, AnonymousFuncExprType}:             _GotoState59Action,
	{_State471, IdentifierToken}:                   _GotoState113Action,
	{_State471, UnderscoreToken}:                   _GotoState119Action,
	{_State471, StructToken}:                       _GotoState50Action,
	{_State471, EnumToken}:                         _GotoState110Action,
	{_State471, TraitToken}:                        _GotoState118Action,
	{_State471, FuncToken}:                         _GotoState112Action,
	{_State471, LparenToken}:                       _GotoState114Action,
	{_State471, LbracketToken}:                     _GotoState42Action,
	{_State471, DotToken}:                          _GotoState109Action,
	{_State471, QuestionToken}:                     _GotoState116Action,
	{_State471, ExclaimToken}:                      _GotoState111Action,
	{_State471, TildeTildeToken}:                   _GotoState117Action,
	{_State471, BitNegToken}:                       _GotoState108Action,
	{_State471, BitAndToken}:                       _GotoState107Action,
	{_State471, ParseErrorToken}:                   _GotoState115Action,
	{_State471, InitializableTypeExprType}:         _GotoState127Action,
	{_State471, SliceTypeExprType}:                 _GotoState101Action,
	{_State471, ArrayTypeExprType}:                 _GotoState60Action,
	{_State471, MapTypeExprType}:                   _GotoState88Action,
	{_State471, AtomTypeExprType}:                  _GotoState120Action,
	{_State471, NamedTypeExprType}:                 _GotoState128Action,
	{_State471, InferredTypeExprType}:              _GotoState126Action,
	{_State471, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State471, ReturnableTypeExprType}:            _GotoState433Action,
	{_State471, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State471, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State471, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State471, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State471, TraitTypeExprType}:                 _GotoState133Action,
	{_State471, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State471, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State471, ReturnTypeType}:                    _GotoState476Action,
	{_State471, FuncTypeExprType}:                  _GotoState123Action,
	{_State472, LbraceToken}:                       _GotoState11Action,
	{_State472, StatementBlockType}:                _GotoState477Action,
	{_State473, IdentifierToken}:                   _GotoState256Action,
	{_State473, UnderscoreToken}:                   _GotoState257Action,
	{_State473, ParameterDefType}:                  _GotoState279Action,
	{_State473, ProperParameterDefsType}:           _GotoState281Action,
	{_State473, ParameterDefsType}:                 _GotoState478Action,
	{_State475, DoToken}:                           _GotoState479Action,
	{_State478, RparenToken}:                       _GotoState480Action,
	{_State479, LbraceToken}:                       _GotoState11Action,
	{_State479, StatementBlockType}:                _GotoState481Action,
	{_State480, IdentifierToken}:                   _GotoState113Action,
	{_State480, UnderscoreToken}:                   _GotoState119Action,
	{_State480, StructToken}:                       _GotoState50Action,
	{_State480, EnumToken}:                         _GotoState110Action,
	{_State480, TraitToken}:                        _GotoState118Action,
	{_State480, FuncToken}:                         _GotoState112Action,
	{_State480, LparenToken}:                       _GotoState114Action,
	{_State480, LbracketToken}:                     _GotoState42Action,
	{_State480, DotToken}:                          _GotoState109Action,
	{_State480, QuestionToken}:                     _GotoState116Action,
	{_State480, ExclaimToken}:                      _GotoState111Action,
	{_State480, TildeTildeToken}:                   _GotoState117Action,
	{_State480, BitNegToken}:                       _GotoState108Action,
	{_State480, BitAndToken}:                       _GotoState107Action,
	{_State480, ParseErrorToken}:                   _GotoState115Action,
	{_State480, InitializableTypeExprType}:         _GotoState127Action,
	{_State480, SliceTypeExprType}:                 _GotoState101Action,
	{_State480, ArrayTypeExprType}:                 _GotoState60Action,
	{_State480, MapTypeExprType}:                   _GotoState88Action,
	{_State480, AtomTypeExprType}:                  _GotoState120Action,
	{_State480, NamedTypeExprType}:                 _GotoState128Action,
	{_State480, InferredTypeExprType}:              _GotoState126Action,
	{_State480, ParseErrorTypeExprType}:            _GotoState129Action,
	{_State480, ReturnableTypeExprType}:            _GotoState433Action,
	{_State480, PrefixedUnaryTypeExprType}:         _GotoState131Action,
	{_State480, PrefixUnaryTypeOpType}:             _GotoState130Action,
	{_State480, ImplicitStructTypeExprType}:        _GotoState125Action,
	{_State480, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State480, TraitTypeExprType}:                 _GotoState133Action,
	{_State480, ImplicitEnumTypeExprType}:          _GotoState124Action,
	{_State480, ExplicitEnumTypeExprType}:          _GotoState122Action,
	{_State480, ReturnTypeType}:                    _GotoState482Action,
	{_State480, FuncTypeExprType}:                  _GotoState123Action,
	{_State482, LbraceToken}:                       _GotoState11Action,
	{_State482, StatementBlockType}:                _GotoState483Action,
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
	{_State131, _WildcardMarker}:                   _ReducePrefixedUnaryTypeExprToReturnableTypeExprAction,
	{_State132, _WildcardMarker}:                   _ReduceReturnableTypeExprToTypeExprAction,
	{_State133, _WildcardMarker}:                   _ReduceTraitTypeExprToAtomTypeExprAction,
	{_State134, LparenToken}:                       _ReduceNilToOptionalGenericParametersAction,
	{_State136, RbraceToken}:                       _ReduceProperStatementsToStatementsAction,
	{_State137, _WildcardMarker}:                   _ReduceStatementToProperStatementsAction,
	{_State139, _WildcardMarker}:                   _ReduceWithSpecToPackageDefAction,
	{_State140, _WildcardMarker}:                   _ReduceNilToOptionalGenericParametersAction,
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
	{_State248, _WildcardMarker}:                   _ReduceToPrefixedUnaryTypeExprAction,
	{_State249, _WildcardMarker}:                   _ReduceAddToBinaryTypeOpAction,
	{_State250, _WildcardMarker}:                   _ReduceMulToBinaryTypeOpAction,
	{_State251, _WildcardMarker}:                   _ReduceSubToBinaryTypeOpAction,
	{_State253, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State254, RbracketToken}:                     _ReduceNilToGenericParameterDefsAction,
	{_State256, _WildcardMarker}:                   _ReduceInferredRefArgToParameterDefAction,
	{_State257, _WildcardMarker}:                   _ReduceIgnoreInferredRefArgToParameterDefAction,
	{_State259, RbraceToken}:                       _ReduceImproperImplicitToStatementsAction,
	{_State259, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State260, RbraceToken}:                       _ReduceImproperExplicitToStatementsAction,
	{_State260, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State261, _EndMarker}:                        _ReduceToStatementBlockAction,
	{_State264, _WildcardMarker}:                   _ReduceAddToProperDefinitionsAction,
	{_State265, _WildcardMarker}:                   _ReduceGlobalVarAssignmentToDefinitionAction,
	{_State266, _WildcardMarker}:                   _ReduceEllipsisToFieldVarPatternAction,
	{_State267, _WildcardMarker}:                   _ReduceIdentifierToVarPatternAction,
	{_State268, _WildcardMarker}:                   _ReduceFieldVarPatternToFieldVarPatternsAction,
	{_State270, _WildcardMarker}:                   _ReducePositionalToFieldVarPatternAction,
	{_State271, _EndMarker}:                        _ReduceToVarDeclPatternAction,
	{_State272, _WildcardMarker}:                   _ReduceTypeExprToOptionalTypeExprAction,
	{_State273, _WildcardMarker}:                   _ReduceEnumNondataMatchPattenToCasePatternAction,
	{_State275, _EndMarker}:                        _ReduceNilToOptionalSimpleStatementAction,
	{_State275, NewlinesToken}:                     _ReduceNilToOptionalSimpleStatementAction,
	{_State275, RbraceToken}:                       _ReduceNilToOptionalSimpleStatementAction,
	{_State275, SemicolonToken}:                    _ReduceNilToOptionalSimpleStatementAction,
	{_State275, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State276, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State277, _EndMarker}:                        _ReduceDefaultBranchStatementToStatementAction,
	{_State278, _EndMarker}:                        _ReduceSimpleStatementToOptionalSimpleStatementAction,
	{_State279, _WildcardMarker}:                   _ReduceParameterDefToProperParameterDefsAction,
	{_State281, RparenToken}:                       _ReduceProperParameterDefsToParameterDefsAction,
	{_State282, _EndMarker}:                        _ReduceImportToLocalToImportClauseAction,
	{_State283, _EndMarker}:                        _ReduceAliasToImportClauseAction,
	{_State284, _WildcardMarker}:                   _ReduceImportClauseToProperImportClausesAction,
	{_State286, RparenToken}:                       _ReduceProperImportClausesToImportClausesAction,
	{_State287, _EndMarker}:                        _ReduceUnusableImportToImportClauseAction,
	{_State290, _WildcardMarker}:                   _ReduceToSliceTypeExprAction,
	{_State291, _WildcardMarker}:                   _ReduceUnitExprPairToColonExprAction,
	{_State292, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State293, _WildcardMarker}:                   _ReduceToImplicitStructExprAction,
	{_State294, ColonToken}:                        _ReduceColonExprUnitTupleToColonExprAction,
	{_State294, CommaToken}:                        _ReduceColonExprUnitTupleToColonExprAction,
	{_State294, RbracketToken}:                     _ReduceColonExprUnitTupleToColonExprAction,
	{_State294, RparenToken}:                       _ReduceColonExprUnitTupleToColonExprAction,
	{_State294, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State295, ColonToken}:                        _ReduceExprUnitPairToColonExprAction,
	{_State295, CommaToken}:                        _ReduceExprUnitPairToColonExprAction,
	{_State295, RbracketToken}:                     _ReduceExprUnitPairToColonExprAction,
	{_State295, RparenToken}:                       _ReduceExprUnitPairToColonExprAction,
	{_State295, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State296, _WildcardMarker}:                   _ReduceVarargAssignmentToArgumentAction,
	{_State297, RparenToken}:                       _ReduceImproperToArgumentsAction,
	{_State297, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State299, _WildcardMarker}:                   _ReduceFieldDefToProperExplicitFieldDefsAction,
	{_State300, RparenToken}:                       _ReduceProperExplicitFieldDefsToExplicitFieldDefsAction,
	{_State302, RbracketToken}:                     _ReduceProperTypeArgumentsToTypeArgumentsAction,
	{_State304, _WildcardMarker}:                   _ReduceTypeExprToProperTypeArgumentsAction,
	{_State305, _WildcardMarker}:                   _ReduceToAccessExprAction,
	{_State307, _EndMarker}:                        _ReduceToBinaryOpAssignStatementAction,
	{_State308, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State308, RparenToken}:                       _ReduceNilToArgumentsAction,
	{_State309, _WildcardMarker}:                   _ReduceToBinaryAddExprAction,
	{_State310, _WildcardMarker}:                   _ReduceToBinaryAndExprAction,
	{_State311, _EndMarker}:                        _ReduceToAssignStatementAction,
	{_State312, _WildcardMarker}:                   _ReduceToBinaryCmpExprAction,
	{_State313, _WildcardMarker}:                   _ReduceAddToExprsAction,
	{_State315, _EndMarker}:                        _ReduceLabeledValuedToJumpStatementAction,
	{_State316, _WildcardMarker}:                   _ReduceToBinaryMulExprAction,
	{_State317, _WildcardMarker}:                   _ReduceInfiniteToLoopExprAction,
	{_State320, _WildcardMarker}:                   _ReduceToAssignPatternAction,
	{_State320, SemicolonToken}:                    _ReduceSequenceExprToForAssignmentAction,
	{_State321, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State323, LbraceToken}:                       _ReduceSequenceExprToConditionAction,
	{_State325, _WildcardMarker}:                   _ReduceToBinaryOrExprAction,
	{_State328, RparenToken}:                       _ReduceProperExplicitEnumValueDefsToExplicitEnumValueDefsAction,
	{_State330, _WildcardMarker}:                   _ReduceNilToGenericTypeArgumentsAction,
	{_State331, _WildcardMarker}:                   _ReduceUnderscoreToInferredTypeExprAction,
	{_State332, _WildcardMarker}:                   _ReduceParameterDeclToProperParameterDeclsAction,
	{_State334, RparenToken}:                       _ReduceProperParameterDeclsToParameterDeclsAction,
	{_State335, _WildcardMarker}:                   _ReduceUnnamedArgToParameterDeclAction,
	{_State336, _WildcardMarker}:                   _ReduceNilToGenericTypeArgumentsAction,
	{_State338, _WildcardMarker}:                   _ReduceDotToInferredTypeExprAction,
	{_State339, _WildcardMarker}:                   _ReduceNilToOptionalDefaultAction,
	{_State340, _WildcardMarker}:                   _ReduceStructPaddingToFieldDefAction,
	{_State342, _WildcardMarker}:                   _ReduceToImplicitEnumTypeExprAction,
	{_State343, _WildcardMarker}:                   _ReduceToImplicitStructTypeExprAction,
	{_State344, RparenToken}:                       _ReduceImproperToImplicitEnumValueDefsAction,
	{_State346, RparenToken}:                       _ReduceImproperToImplicitFieldDefsAction,
	{_State348, _WildcardMarker}:                   _ReduceUnnamedToFieldDefAction,
	{_State350, _WildcardMarker}:                   _ReduceToBinaryTypeExprAction,
	{_State351, _WildcardMarker}:                   _ReduceAliasToNamedFuncDefAction,
	{_State352, _WildcardMarker}:                   _ReduceUnconstrainedToGenericParameterDefAction,
	{_State353, _WildcardMarker}:                   _ReduceGenericParameterDefToProperGenericParameterDefsAction,
	{_State355, RbracketToken}:                     _ReduceProperGenericParameterDefsToGenericParameterDefsAction,
	{_State356, RparenToken}:                       _ReduceNilToParameterDefsAction,
	{_State357, _WildcardMarker}:                   _ReduceInferredRefVarargToParameterDefAction,
	{_State358, _WildcardMarker}:                   _ReduceArgToParameterDefAction,
	{_State359, _WildcardMarker}:                   _ReduceIgnoreInferredRefVarargToParameterDefAction,
	{_State360, _WildcardMarker}:                   _ReduceIgnoreArgToParameterDefAction,
	{_State362, _WildcardMarker}:                   _ReduceAddImplicitToProperStatementsAction,
	{_State363, _WildcardMarker}:                   _ReduceAddExplicitToProperStatementsAction,
	{_State364, _WildcardMarker}:                   _ReduceAliasToTypeDefAction,
	{_State365, _WildcardMarker}:                   _ReduceDefinitionToTypeDefAction,
	{_State368, _WildcardMarker}:                   _ReduceToTuplePatternAction,
	{_State369, _WildcardMarker}:                   _ReduceEnumMatchPatternToCasePatternAction,
	{_State371, _EndMarker}:                        _ReduceCaseBranchStatementToStatementAction,
	{_State372, _WildcardMarker}:                   _ReduceMultipleToCasePatternsAction,
	{_State373, LbraceToken}:                       _ReduceNilToReturnTypeAction,
	{_State374, RparenToken}:                       _ReduceImproperToParameterDefsAction,
	{_State375, _EndMarker}:                        _ReduceMultipleToImportStatementAction,
	{_State376, RparenToken}:                       _ReduceExplicitToImportClausesAction,
	{_State377, RparenToken}:                       _ReduceImplicitToImportClausesAction,
	{_State380, _WildcardMarker}:                   _ReduceNamedAssignmentToArgumentAction,
	{_State381, _WildcardMarker}:                   _ReduceColonExprExprTupleToColonExprAction,
	{_State382, _WildcardMarker}:                   _ReduceExprExprPairToColonExprAction,
	{_State383, _WildcardMarker}:                   _ReduceAddToProperArgumentsAction,
	{_State384, _WildcardMarker}:                   _ReduceToExplicitStructTypeExprAction,
	{_State385, RparenToken}:                       _ReduceImproperExplicitToExplicitFieldDefsAction,
	{_State386, RparenToken}:                       _ReduceImproperImplicitToExplicitFieldDefsAction,
	{_State388, RbracketToken}:                     _ReduceImproperToTypeArgumentsAction,
	{_State389, _WildcardMarker}:                   _ReduceBindingToGenericTypeArgumentsAction,
	{_State390, _WildcardMarker}:                   _ReduceToIndexExprAction,
	{_State392, _WildcardMarker}:                   _ReduceToInitializeExprAction,
	{_State393, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State394, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State395, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State396, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State396, SemicolonToken}:                    _ReduceNilToOptionalSequenceExprAction,
	{_State399, _WildcardMarker}:                   _ReduceNoElseToIfExprAction,
	{_State400, _EndMarker}:                        _ReduceToSwitchExprAction,
	{_State401, _WildcardMarker}:                   _ReduceToExplicitEnumTypeExprAction,
	{_State404, RparenToken}:                       _ReduceImproperToExplicitEnumValueDefsAction,
	{_State406, _WildcardMarker}:                   _ReduceUnnamedVarargToParameterDeclAction,
	{_State408, _WildcardMarker}:                   _ReduceArgToParameterDeclAction,
	{_State410, _WildcardMarker}:                   _ReduceUnderscoreArgToParameterDeclAction,
	{_State411, _WildcardMarker}:                   _ReduceNilToReturnTypeAction,
	{_State412, RparenToken}:                       _ReduceImproperToParameterDeclsAction,
	{_State413, _WildcardMarker}:                   _ReduceExternalToNamedTypeExprAction,
	{_State414, RparenToken}:                       _ReduceNilToParameterDeclsAction,
	{_State415, _WildcardMarker}:                   _ReduceNamedToFieldDefAction,
	{_State416, _WildcardMarker}:                   _ReducePairToProperImplicitEnumValueDefsAction,
	{_State417, _WildcardMarker}:                   _ReduceAddToProperImplicitEnumValueDefsAction,
	{_State418, _WildcardMarker}:                   _ReduceAddToProperImplicitFieldDefsAction,
	{_State419, _WildcardMarker}:                   _ReduceDefaultToOptionalDefaultAction,
	{_State420, _WildcardMarker}:                   _ReduceToTraitTypeExprAction,
	{_State421, _WildcardMarker}:                   _ReduceConstrainedToGenericParameterDefAction,
	{_State422, _WildcardMarker}:                   _ReduceGenericToOptionalGenericParametersAction,
	{_State423, RbracketToken}:                     _ReduceImproperToGenericParameterDefsAction,
	{_State425, _WildcardMarker}:                   _ReduceVarargToParameterDefAction,
	{_State426, _WildcardMarker}:                   _ReduceIgnoreVarargToParameterDefAction,
	{_State427, LparenToken}:                       _ReduceNilToOptionalGenericParametersAction,
	{_State429, _WildcardMarker}:                   _ReduceNamedToFieldVarPatternAction,
	{_State430, _WildcardMarker}:                   _ReduceAddToFieldVarPatternsAction,
	{_State431, _WildcardMarker}:                   _ReduceEnumVarDeclPatternToCasePatternAction,
	{_State433, _WildcardMarker}:                   _ReduceReturnableTypeExprToReturnTypeAction,
	{_State434, _WildcardMarker}:                   _ReduceAddToProperParameterDefsAction,
	{_State435, _WildcardMarker}:                   _ReduceAddExplicitToProperImportClausesAction,
	{_State436, _WildcardMarker}:                   _ReduceAddImplicitToProperImportClausesAction,
	{_State437, _WildcardMarker}:                   _ReduceToMapTypeExprAction,
	{_State438, _WildcardMarker}:                   _ReduceToArrayTypeExprAction,
	{_State439, _WildcardMarker}:                   _ReduceAddExplicitToProperExplicitFieldDefsAction,
	{_State440, _WildcardMarker}:                   _ReduceAddImplicitToProperExplicitFieldDefsAction,
	{_State441, _EndMarker}:                        _ReduceToUnsafeStatementAction,
	{_State442, _WildcardMarker}:                   _ReduceAddToProperTypeArgumentsAction,
	{_State443, _WildcardMarker}:                   _ReduceToCallExprAction,
	{_State444, _EndMarker}:                        _ReduceDoWhileToLoopExprAction,
	{_State445, SemicolonToken}:                    _ReduceAssignToForAssignmentAction,
	{_State448, DoToken}:                           _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State449, _EndMarker}:                        _ReduceWhileToLoopExprAction,
	{_State450, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State452, _WildcardMarker}:                   _ReduceImplicitPairToProperExplicitEnumValueDefsAction,
	{_State453, _WildcardMarker}:                   _ReduceExplicitPairToProperExplicitEnumValueDefsAction,
	{_State454, _WildcardMarker}:                   _ReduceImplicitAddToProperExplicitEnumValueDefsAction,
	{_State455, _WildcardMarker}:                   _ReduceExplicitAddToProperExplicitEnumValueDefsAction,
	{_State456, _WildcardMarker}:                   _ReduceVarargToParameterDeclAction,
	{_State457, _WildcardMarker}:                   _ReduceUnderscoreVarargToParameterDeclAction,
	{_State458, _WildcardMarker}:                   _ReduceToFuncTypeExprAction,
	{_State459, _WildcardMarker}:                   _ReduceAddToProperParameterDeclsAction,
	{_State461, _WildcardMarker}:                   _ReduceAddToProperGenericParameterDefsAction,
	{_State462, LbraceToken}:                       _ReduceNilToReturnTypeAction,
	{_State464, _WildcardMarker}:                   _ReduceConstrainedDefToTypeDefAction,
	{_State465, _WildcardMarker}:                   _ReduceToAnonymousFuncExprAction,
	{_State467, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State467, DoToken}:                           _ReduceNilToOptionalSequenceExprAction,
	{_State468, LbraceToken}:                       _ReduceCaseToConditionAction,
	{_State469, _EndMarker}:                        _ReduceMultiIfElseToIfExprAction,
	{_State470, _EndMarker}:                        _ReduceIfElseToIfExprAction,
	{_State471, _WildcardMarker}:                   _ReduceNilToReturnTypeAction,
	{_State473, RparenToken}:                       _ReduceNilToParameterDefsAction,
	{_State474, _EndMarker}:                        _ReduceIteratorToLoopExprAction,
	{_State476, _WildcardMarker}:                   _ReduceToMethodSignatureAction,
	{_State477, _WildcardMarker}:                   _ReduceFuncDefToNamedFuncDefAction,
	{_State480, LbraceToken}:                       _ReduceNilToReturnTypeAction,
	{_State481, _EndMarker}:                        _ReduceForToLoopExprAction,
	{_State483, _WildcardMarker}:                   _ReduceMethodDefToNamedFuncDefAction,
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
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
      named_func_def: FUNC.IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC.LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
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
      type_def: TYPE.IDENTIFIER optional_generic_parameters type_expr
      type_def: TYPE.IDENTIFIER optional_generic_parameters type_expr IMPLEMENTS type_expr
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
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
      prefixed_unary_type_expr: prefix_unary_type_op.returnable_type_expr
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 131:
    Kernel Items:
      returnable_type_expr: prefixed_unary_type_expr., *
    Reduce:
      * -> [returnable_type_expr]
    Goto:
      (nil)

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
      named_func_def: FUNC IDENTIFIER.optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC IDENTIFIER.ASSIGN expr
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 254
      ASSIGN -> State 253
      optional_generic_parameters -> State 255

  State 135:
    Kernel Items:
      named_func_def: FUNC LPAREN.parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 256
      UNDERSCORE -> State 257
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
      NEWLINES -> State 259
      SEMICOLON -> State 260

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
      RBRACE -> State 261

  State 139:
    Kernel Items:
      package_def: PACKAGE statement_block., *
    Reduce:
      * -> [package_def]
    Goto:
      (nil)

  State 140:
    Kernel Items:
      type_def: TYPE IDENTIFIER.optional_generic_parameters type_expr
      type_def: TYPE IDENTIFIER.optional_generic_parameters type_expr IMPLEMENTS type_expr
      type_def: TYPE IDENTIFIER.ASSIGN type_expr
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 254
      ASSIGN -> State 262
      optional_generic_parameters -> State 263

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
      definition -> State 264
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
      expr -> State 265
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
      IDENTIFIER -> State 267
      UNDERSCORE -> State 145
      LPAREN -> State 144
      ELLIPSIS -> State 266
      var_pattern -> State 270
      tuple_pattern -> State 146
      field_var_patterns -> State 269
      field_var_pattern -> State 268

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
      optional_type_expr -> State 271
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 272
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
      IDENTIFIER -> State 273

  State 149:
    Kernel Items:
      var_or_let: VAR., *
      case_pattern: VAR.DOT IDENTIFIER tuple_pattern
    Reduce:
      * -> [var_or_let]
    Goto:
      DOT -> State 274

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
      COMMA -> State 276
      COLON -> State 275

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
      simple_statement -> State 278
      optional_simple_statement -> State 277
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
      parameter_def -> State 279
      proper_parameter_defs -> State 281
      parameter_defs -> State 280

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
      STRING_LITERAL -> State 282

  State 159:
    Kernel Items:
      import_clause: IDENTIFIER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 283

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
      import_clause -> State 284
      proper_import_clauses -> State 286
      import_clauses -> State 285

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
      STRING_LITERAL -> State 287

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
      RBRACKET -> State 290
      COMMA -> State 289
      COLON -> State 288
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
      expr -> State 291
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
      ASSIGN -> State 292

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
      RPAREN -> State 293

  State 170:
    Kernel Items:
      argument: colon_expr., *
      colon_expr: colon_expr.COLON
      colon_expr: colon_expr.COLON expr
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 294

  State 171:
    Kernel Items:
      argument: expr., *
      argument: expr.ELLIPSIS
      colon_expr: expr.COLON
      colon_expr: expr.COLON expr
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 295
      ELLIPSIS -> State 296

  State 172:
    Kernel Items:
      proper_arguments: proper_arguments.COMMA argument
      arguments: proper_arguments., RPAREN
      arguments: proper_arguments.COMMA
    Reduce:
      RPAREN -> [arguments]
    Goto:
      COMMA -> State 297

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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 299
      implicit_struct_type_expr -> State 125
      proper_explicit_field_defs -> State 300
      explicit_field_defs -> State 298
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
      IDENTIFIER -> State 301

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
      proper_type_arguments -> State 302
      type_arguments -> State 303
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 304
      binary_type_expr -> State 121
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
      IDENTIFIER -> State 305

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
      argument -> State 306
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
      expr -> State 307
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
      LPAREN -> State 308

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
      mul_expr -> State 309
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
      cmp_expr -> State 310
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
      expr -> State 311
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
      add_expr -> State 312
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
      expr -> State 313
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
      arguments -> State 314
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
      exprs -> State 315
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
      prefixable_expr -> State 316
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
      statement_block -> State 317

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
      assign_pattern -> State 318
      optional_label_decl -> State 153
      sequence_expr -> State 320
      for_assignment -> State 319
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
      CASE -> State 321
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
      sequence_expr -> State 323
      condition -> State 322
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
      sequence_expr -> State 324
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
      and_expr -> State 325
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 327
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      proper_explicit_enum_value_defs -> State 328
      explicit_enum_value_defs -> State 326
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 233:
    Kernel Items:
      func_type_expr: FUNC LPAREN.parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 330
      UNDERSCORE -> State 331
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      ELLIPSIS -> State 329
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 335
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      parameter_decl -> State 332
      proper_parameter_decls -> State 334
      parameter_decls -> State 333
      func_type_expr -> State 123

  State 234:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT.IDENTIFIER generic_type_arguments
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 336

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
      IDENTIFIER -> State 337
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
      DOT -> State 338
      QUESTION -> State 116
      EXCLAIM -> State 111
      DOLLAR_LBRACKET -> State 184
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      generic_type_arguments -> State 235
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 339
      binary_type_expr -> State 121
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 340
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
      OR -> State 341

  State 240:
    Kernel Items:
      implicit_enum_type_expr: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 342

  State 241:
    Kernel Items:
      implicit_struct_type_expr: LPAREN implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 343

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
      NEWLINES -> State 344
      OR -> State 345

  State 244:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs.COMMA field_def
      implicit_field_defs: proper_implicit_field_defs., RPAREN
      implicit_field_defs: proper_implicit_field_defs.COMMA
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      COMMA -> State 346

  State 245:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      field_def: type_expr.optional_default
    Reduce:
      * -> [optional_default]
    Goto:
      ASSIGN -> State 347
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252
      optional_default -> State 348

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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 299
      implicit_struct_type_expr -> State 125
      proper_explicit_field_defs -> State 300
      explicit_field_defs -> State 349
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 248:
    Kernel Items:
      prefixed_unary_type_expr: prefix_unary_type_op returnable_type_expr., *
    Reduce:
      * -> [prefixed_unary_type_expr]
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
      returnable_type_expr -> State 350
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
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
      expr -> State 351
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
      optional_generic_parameters: DOLLAR_LBRACKET.generic_parameter_defs RBRACKET
    Reduce:
      RBRACKET -> [generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 352
      generic_parameter_def -> State 353
      proper_generic_parameter_defs -> State 355
      generic_parameter_defs -> State 354

  State 255:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 356

  State 256:
    Kernel Items:
      parameter_def: IDENTIFIER., *
      parameter_def: IDENTIFIER.ELLIPSIS
      parameter_def: IDENTIFIER.type_expr
      parameter_def: IDENTIFIER.ELLIPSIS type_expr
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
      ELLIPSIS -> State 357
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 358
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 257:
    Kernel Items:
      parameter_def: UNDERSCORE., *
      parameter_def: UNDERSCORE.ELLIPSIS
      parameter_def: UNDERSCORE.type_expr
      parameter_def: UNDERSCORE.ELLIPSIS type_expr
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 360
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 258:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def.RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 361

  State 259:
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
      statement -> State 362
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

  State 260:
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
      statement -> State 363
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
      statement_block: LBRACE statements RBRACE., $
    Reduce:
      $ -> [statement_block]
    Goto:
      (nil)

  State 262:
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 364
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 263:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters.type_expr
      type_def: TYPE IDENTIFIER optional_generic_parameters.type_expr IMPLEMENTS type_expr
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 365
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 264:
    Kernel Items:
      proper_definitions: proper_definitions NEWLINES definition., *
    Reduce:
      * -> [proper_definitions]
    Goto:
      (nil)

  State 265:
    Kernel Items:
      definition: var_decl_pattern ASSIGN expr., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 266:
    Kernel Items:
      field_var_pattern: ELLIPSIS., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 267:
    Kernel Items:
      var_pattern: IDENTIFIER., *
      field_var_pattern: IDENTIFIER.ASSIGN var_pattern
    Reduce:
      * -> [var_pattern]
    Goto:
      ASSIGN -> State 366

  State 268:
    Kernel Items:
      field_var_patterns: field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 269:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns.RPAREN
      field_var_patterns: field_var_patterns.COMMA field_var_pattern
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 368
      COMMA -> State 367

  State 270:
    Kernel Items:
      field_var_pattern: var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 271:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern optional_type_expr., $
    Reduce:
      $ -> [var_decl_pattern]
    Goto:
      (nil)

  State 272:
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

  State 273:
    Kernel Items:
      case_pattern: DOT IDENTIFIER.implicit_struct_expr
      case_pattern: DOT IDENTIFIER., *
    Reduce:
      * -> [case_pattern]
    Goto:
      LPAREN -> State 43
      implicit_struct_expr -> State 369

  State 274:
    Kernel Items:
      case_pattern: VAR DOT.IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 370

  State 275:
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
      simple_statement -> State 278
      optional_simple_statement -> State 371
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

  State 276:
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
      case_pattern -> State 372
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

  State 277:
    Kernel Items:
      statement: DEFAULT COLON optional_simple_statement., $
    Reduce:
      $ -> [statement]
    Goto:
      (nil)

  State 278:
    Kernel Items:
      optional_simple_statement: simple_statement., $
    Reduce:
      $ -> [optional_simple_statement]
    Goto:
      (nil)

  State 279:
    Kernel Items:
      proper_parameter_defs: parameter_def., *
    Reduce:
      * -> [proper_parameter_defs]
    Goto:
      (nil)

  State 280:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 373

  State 281:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs.COMMA parameter_def
      parameter_defs: proper_parameter_defs., RPAREN
      parameter_defs: proper_parameter_defs.COMMA
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      COMMA -> State 374

  State 282:
    Kernel Items:
      import_clause: DOT STRING_LITERAL., $
    Reduce:
      $ -> [import_clause]
    Goto:
      (nil)

  State 283:
    Kernel Items:
      import_clause: IDENTIFIER STRING_LITERAL., $
    Reduce:
      $ -> [import_clause]
    Goto:
      (nil)

  State 284:
    Kernel Items:
      proper_import_clauses: import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 285:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 375

  State 286:
    Kernel Items:
      proper_import_clauses: proper_import_clauses.NEWLINES import_clause
      proper_import_clauses: proper_import_clauses.COMMA import_clause
      import_clauses: proper_import_clauses., RPAREN
      import_clauses: proper_import_clauses.NEWLINES
      import_clauses: proper_import_clauses.COMMA
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      NEWLINES -> State 377
      COMMA -> State 376

  State 287:
    Kernel Items:
      import_clause: UNDERSCORE STRING_LITERAL., $
    Reduce:
      $ -> [import_clause]
    Goto:
      (nil)

  State 288:
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 378
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 289:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 379

  State 290:
    Kernel Items:
      slice_type_expr: LBRACKET type_expr RBRACKET., *
    Reduce:
      * -> [slice_type_expr]
    Goto:
      (nil)

  State 291:
    Kernel Items:
      colon_expr: COLON expr., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 292:
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
      expr -> State 380
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

  State 293:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [implicit_struct_expr]
    Goto:
      (nil)

  State 294:
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
      expr -> State 381
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

  State 295:
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

  State 296:
    Kernel Items:
      argument: expr ELLIPSIS., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 297:
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
      argument -> State 383
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

  State 298:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 384

  State 299:
    Kernel Items:
      proper_explicit_field_defs: field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 300:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs.NEWLINES field_def
      proper_explicit_field_defs: proper_explicit_field_defs.COMMA field_def
      explicit_field_defs: proper_explicit_field_defs., RPAREN
      explicit_field_defs: proper_explicit_field_defs.NEWLINES
      explicit_field_defs: proper_explicit_field_defs.COMMA
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      NEWLINES -> State 386
      COMMA -> State 385

  State 301:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 387

  State 302:
    Kernel Items:
      proper_type_arguments: proper_type_arguments.COMMA type_expr
      type_arguments: proper_type_arguments., RBRACKET
      type_arguments: proper_type_arguments.COMMA
    Reduce:
      RBRACKET -> [type_arguments]
    Goto:
      COMMA -> State 388

  State 303:
    Kernel Items:
      generic_type_arguments: DOLLAR_LBRACKET type_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 389

  State 304:
    Kernel Items:
      proper_type_arguments: type_expr., *
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      * -> [proper_type_arguments]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 305:
    Kernel Items:
      access_expr: accessible_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 306:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 390

  State 307:
    Kernel Items:
      binary_op_assign_statement: accessible_expr binary_op_assign expr., $
    Reduce:
      $ -> [binary_op_assign_statement]
    Goto:
      (nil)

  State 308:
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
      arguments -> State 391
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

  State 309:
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

  State 310:
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

  State 311:
    Kernel Items:
      assign_statement: assign_pattern ASSIGN expr., $
    Reduce:
      $ -> [assign_statement]
    Goto:
      (nil)

  State 312:
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

  State 313:
    Kernel Items:
      exprs: exprs COMMA expr., *
    Reduce:
      * -> [exprs]
    Goto:
      (nil)

  State 314:
    Kernel Items:
      initialize_expr: initializable_type_expr LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 392

  State 315:
    Kernel Items:
      exprs: exprs.COMMA expr
      jump_statement: jump_type JUMP_LABEL exprs., $
    Reduce:
      $ -> [jump_statement]
    Goto:
      COMMA -> State 211

  State 316:
    Kernel Items:
      binary_mul_expr: mul_expr mul_op prefixable_expr., *
    Reduce:
      * -> [binary_mul_expr]
    Goto:
      (nil)

  State 317:
    Kernel Items:
      loop_expr: DO statement_block., *
      loop_expr: DO statement_block.FOR sequence_expr
    Reduce:
      * -> [loop_expr]
    Goto:
      FOR -> State 393

  State 318:
    Kernel Items:
      loop_expr: FOR assign_pattern.IN sequence_expr DO statement_block
      for_assignment: assign_pattern.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      IN -> State 395
      ASSIGN -> State 394

  State 319:
    Kernel Items:
      loop_expr: FOR for_assignment.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 396

  State 320:
    Kernel Items:
      assign_pattern: sequence_expr., *
      loop_expr: FOR sequence_expr.DO statement_block
      for_assignment: sequence_expr., SEMICOLON
    Reduce:
      * -> [assign_pattern]
      SEMICOLON -> [for_assignment]
    Goto:
      DO -> State 397

  State 321:
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
      case_patterns -> State 398
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

  State 322:
    Kernel Items:
      if_expr: IF condition.statement_block
      if_expr: IF condition.statement_block ELSE statement_block
      if_expr: IF condition.statement_block ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 399

  State 323:
    Kernel Items:
      condition: sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 324:
    Kernel Items:
      switch_expr: SWITCH sequence_expr.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 400

  State 325:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      binary_or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [binary_or_expr]
    Goto:
      AND -> State 200

  State 326:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 401

  State 327:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def.OR field_def
      proper_explicit_enum_value_defs: field_def.NEWLINES field_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 402
      OR -> State 403

  State 328:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs.OR field_def
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs.NEWLINES field_def
      explicit_enum_value_defs: proper_explicit_enum_value_defs., RPAREN
      explicit_enum_value_defs: proper_explicit_enum_value_defs.NEWLINES
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      NEWLINES -> State 404
      OR -> State 405

  State 329:
    Kernel Items:
      parameter_decl: ELLIPSIS.type_expr
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 406
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 330:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_type_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_type_arguments
      parameter_decl: IDENTIFIER.type_expr
      parameter_decl: IDENTIFIER.ELLIPSIS type_expr
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
      DOT -> State 338
      QUESTION -> State 116
      EXCLAIM -> State 111
      DOLLAR_LBRACKET -> State 184
      ELLIPSIS -> State 407
      TILDE_TILDE -> State 117
      BIT_NEG -> State 108
      BIT_AND -> State 107
      PARSE_ERROR -> State 115
      generic_type_arguments -> State 235
      initializable_type_expr -> State 127
      slice_type_expr -> State 101
      array_type_expr -> State 60
      map_type_expr -> State 88
      atom_type_expr -> State 120
      named_type_expr -> State 128
      inferred_type_expr -> State 126
      parse_error_type_expr -> State 129
      returnable_type_expr -> State 132
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
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
      inferred_type_expr: UNDERSCORE., *
      parameter_decl: UNDERSCORE.type_expr
      parameter_decl: UNDERSCORE.ELLIPSIS type_expr
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
      ELLIPSIS -> State 409
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 410
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 332:
    Kernel Items:
      proper_parameter_decls: parameter_decl., *
    Reduce:
      * -> [proper_parameter_decls]
    Goto:
      (nil)

  State 333:
    Kernel Items:
      func_type_expr: FUNC LPAREN parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 411

  State 334:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls.COMMA parameter_decl
      parameter_decls: proper_parameter_decls., RPAREN
      parameter_decls: proper_parameter_decls.COMMA
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      COMMA -> State 412

  State 335:
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

  State 336:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT IDENTIFIER.generic_type_arguments
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      DOLLAR_LBRACKET -> State 184
      generic_type_arguments -> State 413

  State 337:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 414

  State 338:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT.IDENTIFIER generic_type_arguments
      inferred_type_expr: DOT., *
    Reduce:
      * -> [inferred_type_expr]
    Goto:
      IDENTIFIER -> State 336

  State 339:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      field_def: IDENTIFIER type_expr.optional_default
    Reduce:
      * -> [optional_default]
    Goto:
      ASSIGN -> State 347
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252
      optional_default -> State 415

  State 340:
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

  State 341:
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
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

  State 342:
    Kernel Items:
      implicit_enum_type_expr: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_type_expr]
    Goto:
      (nil)

  State 343:
    Kernel Items:
      implicit_struct_type_expr: LPAREN implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_type_expr]
    Goto:
      (nil)

  State 344:
    Kernel Items:
      implicit_enum_value_defs: proper_implicit_enum_value_defs NEWLINES., RPAREN
    Reduce:
      RPAREN -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 345:
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 417
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 346:
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 418
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 347:
    Kernel Items:
      optional_default: ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 419

  State 348:
    Kernel Items:
      field_def: type_expr optional_default., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 349:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 420

  State 350:
    Kernel Items:
      binary_type_expr: type_expr binary_type_op returnable_type_expr., *
    Reduce:
      * -> [binary_type_expr]
    Goto:
      (nil)

  State 351:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN expr., *
    Reduce:
      * -> [named_func_def]
    Goto:
      (nil)

  State 352:
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 421
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 353:
    Kernel Items:
      proper_generic_parameter_defs: generic_parameter_def., *
    Reduce:
      * -> [proper_generic_parameter_defs]
    Goto:
      (nil)

  State 354:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET generic_parameter_defs.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 422

  State 355:
    Kernel Items:
      proper_generic_parameter_defs: proper_generic_parameter_defs.COMMA generic_parameter_def
      generic_parameter_defs: proper_generic_parameter_defs., RBRACKET
      generic_parameter_defs: proper_generic_parameter_defs.COMMA
    Reduce:
      RBRACKET -> [generic_parameter_defs]
    Goto:
      COMMA -> State 423

  State 356:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 256
      UNDERSCORE -> State 257
      parameter_def -> State 279
      proper_parameter_defs -> State 281
      parameter_defs -> State 424

  State 357:
    Kernel Items:
      parameter_def: IDENTIFIER ELLIPSIS., *
      parameter_def: IDENTIFIER ELLIPSIS.type_expr
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 425
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 358:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_def: IDENTIFIER type_expr., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 359:
    Kernel Items:
      parameter_def: UNDERSCORE ELLIPSIS., *
      parameter_def: UNDERSCORE ELLIPSIS.type_expr
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 426
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
      parameter_def: UNDERSCORE type_expr., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 361:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN.IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 427

  State 362:
    Kernel Items:
      proper_statements: proper_statements NEWLINES statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 363:
    Kernel Items:
      proper_statements: proper_statements SEMICOLON statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 364:
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

  State 365:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER optional_generic_parameters type_expr., *
      type_def: TYPE IDENTIFIER optional_generic_parameters type_expr.IMPLEMENTS type_expr
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 428
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 366:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN.var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 143
      UNDERSCORE -> State 145
      LPAREN -> State 144
      var_pattern -> State 429
      tuple_pattern -> State 146

  State 367:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA.field_var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 267
      UNDERSCORE -> State 145
      LPAREN -> State 144
      ELLIPSIS -> State 266
      var_pattern -> State 270
      tuple_pattern -> State 146
      field_var_pattern -> State 430

  State 368:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns RPAREN., *
    Reduce:
      * -> [tuple_pattern]
    Goto:
      (nil)

  State 369:
    Kernel Items:
      case_pattern: DOT IDENTIFIER implicit_struct_expr., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 370:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER.tuple_pattern
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 144
      tuple_pattern -> State 431

  State 371:
    Kernel Items:
      statement: CASE case_patterns COLON optional_simple_statement., $
    Reduce:
      $ -> [statement]
    Goto:
      (nil)

  State 372:
    Kernel Items:
      case_patterns: case_patterns COMMA case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 373:
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
      returnable_type_expr -> State 433
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      return_type -> State 432
      func_type_expr -> State 123

  State 374:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs COMMA.parameter_def
      parameter_defs: proper_parameter_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 256
      UNDERSCORE -> State 257
      parameter_def -> State 434

  State 375:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses RPAREN., $
    Reduce:
      $ -> [import_statement]
    Goto:
      (nil)

  State 376:
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
      import_clause -> State 435

  State 377:
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
      import_clause -> State 436

  State 378:
    Kernel Items:
      map_type_expr: LBRACKET type_expr COLON type_expr.RBRACKET
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 437
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 379:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 438

  State 380:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expr., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 381:
    Kernel Items:
      colon_expr: colon_expr COLON expr., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 382:
    Kernel Items:
      colon_expr: expr COLON expr., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 383:
    Kernel Items:
      proper_arguments: proper_arguments COMMA argument., *
    Reduce:
      * -> [proper_arguments]
    Goto:
      (nil)

  State 384:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_type_expr]
    Goto:
      (nil)

  State 385:
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 439
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 386:
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 440
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 387:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 441

  State 388:
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 442
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 389:
    Kernel Items:
      generic_type_arguments: DOLLAR_LBRACKET type_arguments RBRACKET., *
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      (nil)

  State 390:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [index_expr]
    Goto:
      (nil)

  State 391:
    Kernel Items:
      call_expr: accessible_expr generic_type_arguments LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 443

  State 392:
    Kernel Items:
      initialize_expr: initializable_type_expr LPAREN arguments RPAREN., *
    Reduce:
      * -> [initialize_expr]
    Goto:
      (nil)

  State 393:
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

  State 394:
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
      sequence_expr -> State 445
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

  State 395:
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
      sequence_expr -> State 446
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
      sequence_expr -> State 448
      optional_sequence_expr -> State 447
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
      loop_expr: FOR sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 449

  State 398:
    Kernel Items:
      case_patterns: case_patterns.COMMA case_pattern
      condition: CASE case_patterns.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      COMMA -> State 276
      ASSIGN -> State 450

  State 399:
    Kernel Items:
      if_expr: IF condition statement_block., *
      if_expr: IF condition statement_block.ELSE statement_block
      if_expr: IF condition statement_block.ELSE if_expr
    Reduce:
      * -> [if_expr]
    Goto:
      ELSE -> State 451

  State 400:
    Kernel Items:
      switch_expr: SWITCH sequence_expr statement_block., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 401:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_type_expr]
    Goto:
      (nil)

  State 402:
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
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

  State 403:
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
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

  State 404:
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 454
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 405:
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 245
      binary_type_expr -> State 121
      field_def -> State 455
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123
      method_signature -> State 242

  State 406:
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

  State 407:
    Kernel Items:
      parameter_decl: IDENTIFIER ELLIPSIS.type_expr
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 456
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 408:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_decl: IDENTIFIER type_expr., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 409:
    Kernel Items:
      parameter_decl: UNDERSCORE ELLIPSIS.type_expr
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 457
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 410:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_decl: UNDERSCORE type_expr., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 411:
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
      returnable_type_expr -> State 433
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      return_type -> State 458
      func_type_expr -> State 123

  State 412:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls COMMA.parameter_decl
      parameter_decls: proper_parameter_decls COMMA., RPAREN
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 330
      UNDERSCORE -> State 331
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      ELLIPSIS -> State 329
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 335
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      parameter_decl -> State 459
      func_type_expr -> State 123

  State 413:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT IDENTIFIER generic_type_arguments., *
    Reduce:
      * -> [named_type_expr]
    Goto:
      (nil)

  State 414:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 330
      UNDERSCORE -> State 331
      STRUCT -> State 50
      ENUM -> State 110
      TRAIT -> State 118
      FUNC -> State 112
      LPAREN -> State 114
      LBRACKET -> State 42
      DOT -> State 109
      QUESTION -> State 116
      EXCLAIM -> State 111
      ELLIPSIS -> State 329
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 335
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      parameter_decl -> State 332
      proper_parameter_decls -> State 334
      parameter_decls -> State 460
      func_type_expr -> State 123

  State 415:
    Kernel Items:
      field_def: IDENTIFIER type_expr optional_default., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 416:
    Kernel Items:
      proper_implicit_enum_value_defs: field_def OR field_def., *
    Reduce:
      * -> [proper_implicit_enum_value_defs]
    Goto:
      (nil)

  State 417:
    Kernel Items:
      proper_implicit_enum_value_defs: proper_implicit_enum_value_defs OR field_def., *
    Reduce:
      * -> [proper_implicit_enum_value_defs]
    Goto:
      (nil)

  State 418:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [proper_implicit_field_defs]
    Goto:
      (nil)

  State 419:
    Kernel Items:
      optional_default: ASSIGN DEFAULT., *
    Reduce:
      * -> [optional_default]
    Goto:
      (nil)

  State 420:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN explicit_field_defs RPAREN., *
    Reduce:
      * -> [trait_type_expr]
    Goto:
      (nil)

  State 421:
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

  State 422:
    Kernel Items:
      optional_generic_parameters: DOLLAR_LBRACKET generic_parameter_defs RBRACKET., *
    Reduce:
      * -> [optional_generic_parameters]
    Goto:
      (nil)

  State 423:
    Kernel Items:
      proper_generic_parameter_defs: proper_generic_parameter_defs COMMA.generic_parameter_def
      generic_parameter_defs: proper_generic_parameter_defs COMMA., RBRACKET
    Reduce:
      RBRACKET -> [generic_parameter_defs]
    Goto:
      IDENTIFIER -> State 352
      generic_parameter_def -> State 461

  State 424:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 462

  State 425:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_def: IDENTIFIER ELLIPSIS type_expr., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 426:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_def: UNDERSCORE ELLIPSIS type_expr., *
    Reduce:
      * -> [parameter_def]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 427:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER.optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      LPAREN -> [optional_generic_parameters]
    Goto:
      DOLLAR_LBRACKET -> State 254
      optional_generic_parameters -> State 463

  State 428:
    Kernel Items:
      type_def: TYPE IDENTIFIER optional_generic_parameters type_expr IMPLEMENTS.type_expr
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
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      type_expr -> State 464
      binary_type_expr -> State 121
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      func_type_expr -> State 123

  State 429:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 430:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 431:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER tuple_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 432:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 465

  State 433:
    Kernel Items:
      return_type: returnable_type_expr., *
    Reduce:
      * -> [return_type]
    Goto:
      (nil)

  State 434:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [proper_parameter_defs]
    Goto:
      (nil)

  State 435:
    Kernel Items:
      proper_import_clauses: proper_import_clauses COMMA import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 436:
    Kernel Items:
      proper_import_clauses: proper_import_clauses NEWLINES import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 437:
    Kernel Items:
      map_type_expr: LBRACKET type_expr COLON type_expr RBRACKET., *
    Reduce:
      * -> [map_type_expr]
    Goto:
      (nil)

  State 438:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA INTEGER_LITERAL RBRACKET., *
    Reduce:
      * -> [array_type_expr]
    Goto:
      (nil)

  State 439:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 440:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 441:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., $
    Reduce:
      $ -> [unsafe_statement]
    Goto:
      (nil)

  State 442:
    Kernel Items:
      proper_type_arguments: proper_type_arguments COMMA type_expr., *
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      * -> [proper_type_arguments]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 443:
    Kernel Items:
      call_expr: accessible_expr generic_type_arguments LPAREN arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 444:
    Kernel Items:
      loop_expr: DO statement_block FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 445:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN sequence_expr., SEMICOLON
    Reduce:
      SEMICOLON -> [for_assignment]
    Goto:
      (nil)

  State 446:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 466

  State 447:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 467

  State 448:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 449:
    Kernel Items:
      loop_expr: FOR sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 450:
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
      sequence_expr -> State 468
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

  State 451:
    Kernel Items:
      if_expr: IF condition statement_block ELSE.statement_block
      if_expr: IF condition statement_block ELSE.if_expr
    Reduce:
      (nil)
    Goto:
      IF -> State 224
      LBRACE -> State 11
      statement_block -> State 470
      if_expr -> State 469

  State 452:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def NEWLINES field_def., *
    Reduce:
      * -> [proper_explicit_enum_value_defs]
    Goto:
      (nil)

  State 453:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def OR field_def., *
    Reduce:
      * -> [proper_explicit_enum_value_defs]
    Goto:
      (nil)

  State 454:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs NEWLINES field_def., *
    Reduce:
      * -> [proper_explicit_enum_value_defs]
    Goto:
      (nil)

  State 455:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs OR field_def., *
    Reduce:
      * -> [proper_explicit_enum_value_defs]
    Goto:
      (nil)

  State 456:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_decl: IDENTIFIER ELLIPSIS type_expr., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 457:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_decl: UNDERSCORE ELLIPSIS type_expr., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 458:
    Kernel Items:
      func_type_expr: FUNC LPAREN parameter_decls RPAREN return_type., *
    Reduce:
      * -> [func_type_expr]
    Goto:
      (nil)

  State 459:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls COMMA parameter_decl., *
    Reduce:
      * -> [proper_parameter_decls]
    Goto:
      (nil)

  State 460:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 471

  State 461:
    Kernel Items:
      proper_generic_parameter_defs: proper_generic_parameter_defs COMMA generic_parameter_def., *
    Reduce:
      * -> [proper_generic_parameter_defs]
    Goto:
      (nil)

  State 462:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN.return_type statement_block
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
      returnable_type_expr -> State 433
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      return_type -> State 472
      func_type_expr -> State 123

  State 463:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 473

  State 464:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER optional_generic_parameters type_expr IMPLEMENTS type_expr., *
    Reduce:
      * -> [type_def]
    Goto:
      ADD -> State 249
      SUB -> State 251
      MUL -> State 250
      binary_type_op -> State 252

  State 465:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 466:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 474

  State 467:
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
      sequence_expr -> State 448
      optional_sequence_expr -> State 475
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

  State 468:
    Kernel Items:
      condition: CASE case_patterns ASSIGN sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 469:
    Kernel Items:
      if_expr: IF condition statement_block ELSE if_expr., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 470:
    Kernel Items:
      if_expr: IF condition statement_block ELSE statement_block., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 471:
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
      returnable_type_expr -> State 433
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      return_type -> State 476
      func_type_expr -> State 123

  State 472:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 477

  State 473:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 256
      UNDERSCORE -> State 257
      parameter_def -> State 279
      proper_parameter_defs -> State 281
      parameter_defs -> State 478

  State 474:
    Kernel Items:
      loop_expr: FOR assign_pattern IN sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 475:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 479

  State 476:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 477:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [named_func_def]
    Goto:
      (nil)

  State 478:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 480

  State 479:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 481

  State 480:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN.return_type statement_block
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
      returnable_type_expr -> State 433
      prefixed_unary_type_expr -> State 131
      prefix_unary_type_op -> State 130
      implicit_struct_type_expr -> State 125
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 133
      implicit_enum_type_expr -> State 124
      explicit_enum_type_expr -> State 122
      return_type -> State 482
      func_type_expr -> State 123

  State 481:
    Kernel Items:
      loop_expr: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 482:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 483

  State 483:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER optional_generic_parameters LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [named_func_def]
    Goto:
      (nil)

Number of states: 483
Number of shift actions: 4601
Number of reduce actions: 431
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 5309
Number of unoptimized shift actions: 46417
Number of unoptimized reduce actions: 34395
*/
