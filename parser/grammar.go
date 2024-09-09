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
	// 365:23: if_expr -> ...
	ToIfExpr(OptionalLabelDecl_ TokenValue, IfExprBody_ Expression) (Expression, error)
}

type IfExprBodyReducer interface {
	// 368:2: if_expr_body -> no_else: ...
	NoElseToIfExprBody(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol) (Expression, error)

	// 369:2: if_expr_body -> if_else: ...
	IfElseToIfExprBody(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol, Else_ TokenValue, StatementBlock_2 GenericSymbol) (Expression, error)

	// 370:2: if_expr_body -> multi_if_else: ...
	MultiIfElseToIfExprBody(If_ TokenValue, Condition_ GenericSymbol, StatementBlock_ GenericSymbol, Else_ TokenValue, IfExpr_ Expression) (Expression, error)
}

type ConditionReducer interface {
	// 373:2: condition -> sequence_expr: ...
	SequenceExprToCondition(SequenceExpr_ Expression) (GenericSymbol, error)

	// 374:2: condition -> case: ...
	CaseToCondition(Case_ TokenValue, CasePatterns_ GenericSymbol, Assign_ TokenValue, SequenceExpr_ Expression) (GenericSymbol, error)
}

type SwitchExprReducer interface {
	// 399:2: switch_expr -> ...
	ToSwitchExpr(OptionalLabelDecl_ TokenValue, Switch_ TokenValue, SequenceExpr_ Expression, StatementBlock_ GenericSymbol) (Expression, error)
}

type LoopExprReducer interface {
	// 412:25: loop_expr -> ...
	ToLoopExpr(OptionalLabelDecl_ TokenValue, LoopExprBody_ Expression) (Expression, error)
}

type LoopExprBodyReducer interface {
	// 415:2: loop_expr_body -> infinite: ...
	InfiniteToLoopExprBody(Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 416:2: loop_expr_body -> do_while: ...
	DoWhileToLoopExprBody(Do_ TokenValue, StatementBlock_ GenericSymbol, For_ TokenValue, SequenceExpr_ Expression) (Expression, error)

	// 417:2: loop_expr_body -> while: ...
	WhileToLoopExprBody(For_ TokenValue, SequenceExpr_ Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 418:2: loop_expr_body -> iterator: ...
	IteratorToLoopExprBody(For_ TokenValue, AssignPattern_ GenericSymbol, In_ TokenValue, SequenceExpr_ Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)

	// 419:2: loop_expr_body -> for: ...
	ForToLoopExprBody(For_ TokenValue, ForAssignment_ GenericSymbol, Semicolon_ TokenValue, OptionalSequenceExpr_ Expression, Semicolon_2 TokenValue, OptionalSequenceExpr_2 Expression, Do_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)
}

type OptionalSequenceExprReducer interface {
	// 422:2: optional_sequence_expr -> sequence_expr: ...
	SequenceExprToOptionalSequenceExpr(SequenceExpr_ Expression) (Expression, error)

	// 423:2: optional_sequence_expr -> nil: ...
	NilToOptionalSequenceExpr() (Expression, error)
}

type ForAssignmentReducer interface {
	// 426:2: for_assignment -> sequence_expr: ...
	SequenceExprToForAssignment(SequenceExpr_ Expression) (GenericSymbol, error)

	// 427:2: for_assignment -> assign: ...
	AssignToForAssignment(AssignPattern_ GenericSymbol, Assign_ TokenValue, SequenceExpr_ Expression) (GenericSymbol, error)
}

type CallExprReducer interface {
	// 434:2: call_expr -> ...
	ToCallExpr(AccessibleExpr_ Expression, GenericTypeArguments_ *TypeArgumentList, Lparen_ TokenValue, Arguments_ *ArgumentList, Rparen_ TokenValue) (Expression, error)
}

type ProperArgumentsReducer interface {
	// 437:2: proper_arguments -> add: ...
	AddToProperArguments(ProperArguments_ *ArgumentList, Comma_ TokenValue, Argument_ *Argument) (*ArgumentList, error)

	// 438:2: proper_arguments -> argument: ...
	ArgumentToProperArguments(Argument_ *Argument) (*ArgumentList, error)
}

type ArgumentsReducer interface {

	// 442:2: arguments -> improper: ...
	ImproperToArguments(ProperArguments_ *ArgumentList, Comma_ TokenValue) (*ArgumentList, error)

	// 443:2: arguments -> nil: ...
	NilToArguments() (*ArgumentList, error)
}

type ArgumentReducer interface {
	// 446:2: argument -> positional: ...
	PositionalToArgument(Expr_ Expression) (*Argument, error)

	// 447:2: argument -> colon_expr: ...
	ColonExprToArgument(ColonExpr_ *ColonExpr) (*Argument, error)

	// 448:2: argument -> named_assignment: ...
	NamedAssignmentToArgument(Identifier_ TokenValue, Assign_ TokenValue, Expr_ Expression) (*Argument, error)

	// 452:2: argument -> vararg_assignment: ...
	VarargAssignmentToArgument(Expr_ Expression, Ellipsis_ TokenValue) (*Argument, error)

	// 455:2: argument -> skip_pattern: ...
	SkipPatternToArgument(Ellipsis_ TokenValue) (*Argument, error)
}

type ColonExprReducer interface {
	// 459:2: colon_expr -> unit_unit_pair: ...
	UnitUnitPairToColonExpr(Colon_ TokenValue) (*ColonExpr, error)

	// 460:2: colon_expr -> expr_unit_pair: ...
	ExprUnitPairToColonExpr(Expr_ Expression, Colon_ TokenValue) (*ColonExpr, error)

	// 461:2: colon_expr -> unit_expr_pair: ...
	UnitExprPairToColonExpr(Colon_ TokenValue, Expr_ Expression) (*ColonExpr, error)

	// 462:2: colon_expr -> expr_expr_pair: ...
	ExprExprPairToColonExpr(Expr_ Expression, Colon_ TokenValue, Expr_2 Expression) (*ColonExpr, error)

	// 463:2: colon_expr -> colon_expr_unit_tuple: ...
	ColonExprUnitTupleToColonExpr(ColonExpr_ *ColonExpr, Colon_ TokenValue) (*ColonExpr, error)

	// 464:2: colon_expr -> colon_expr_expr_tuple: ...
	ColonExprExprTupleToColonExpr(ColonExpr_ *ColonExpr, Colon_ TokenValue, Expr_ Expression) (*ColonExpr, error)
}

type ParseErrorExprReducer interface {
	// 483:32: parse_error_expr -> ...
	ToParseErrorExpr(ParseError_ ParseErrorSymbol) (Expression, error)
}

type LiteralExprReducer interface {
	// 486:2: literal_expr -> TRUE: ...
	TrueToLiteralExpr(True_ TokenValue) (Expression, error)

	// 487:2: literal_expr -> FALSE: ...
	FalseToLiteralExpr(False_ TokenValue) (Expression, error)

	// 488:2: literal_expr -> INTEGER_LITERAL: ...
	IntegerLiteralToLiteralExpr(IntegerLiteral_ TokenValue) (Expression, error)

	// 489:2: literal_expr -> FLOAT_LITERAL: ...
	FloatLiteralToLiteralExpr(FloatLiteral_ TokenValue) (Expression, error)

	// 490:2: literal_expr -> RUNE_LITERAL: ...
	RuneLiteralToLiteralExpr(RuneLiteral_ TokenValue) (Expression, error)

	// 491:2: literal_expr -> STRING_LITERAL: ...
	StringLiteralToLiteralExpr(StringLiteral_ TokenValue) (Expression, error)
}

type NamedExprReducer interface {
	// 494:2: named_expr -> IDENTIFIER: ...
	IdentifierToNamedExpr(Identifier_ TokenValue) (Expression, error)

	// 495:2: named_expr -> UNDERSCORE: ...
	UnderscoreToNamedExpr(Underscore_ TokenValue) (Expression, error)
}

type BlockExprReducer interface {
	// 497:27: block_expr -> ...
	ToBlockExpr(OptionalLabelDecl_ TokenValue, StatementBlock_ GenericSymbol) (Expression, error)
}

type InitializeExprReducer interface {
	// 499:31: initialize_expr -> ...
	ToInitializeExpr(InitializableTypeExpr_ TypeExpression, Lparen_ TokenValue, Arguments_ *ArgumentList, Rparen_ TokenValue) (Expression, error)
}

type ImplicitStructExprReducer interface {
	// 501:36: implicit_struct_expr -> ...
	ToImplicitStructExpr(Lparen_ TokenValue, Arguments_ *ArgumentList, Rparen_ TokenValue) (Expression, error)
}

type AccessExprReducer interface {
	// 509:27: access_expr -> ...
	ToAccessExpr(AccessibleExpr_ Expression, Dot_ TokenValue, Identifier_ TokenValue) (Expression, error)
}

type IndexExprReducer interface {
	// 513:26: index_expr -> ...
	ToIndexExpr(AccessibleExpr_ Expression, Lbracket_ TokenValue, Argument_ *Argument, Rbracket_ TokenValue) (Expression, error)
}

type PostfixUnaryExprReducer interface {
	// 523:34: postfix_unary_expr -> ...
	ToPostfixUnaryExpr(AccessibleExpr_ Expression, PostfixUnaryOp_ TokenValue) (Expression, error)
}

type PrefixUnaryExprReducer interface {
	// 529:33: prefix_unary_expr -> ...
	ToPrefixUnaryExpr(PrefixUnaryOp_ TokenValue, PrefixableExpr_ Expression) (Expression, error)
}

type BinaryMulExprReducer interface {
	// 546:31: binary_mul_expr -> ...
	ToBinaryMulExpr(MulExpr_ Expression, MulOp_ TokenValue, PrefixableExpr_ Expression) (Expression, error)
}

type BinaryAddExprReducer interface {
	// 560:31: binary_add_expr -> ...
	ToBinaryAddExpr(AddExpr_ Expression, AddOp_ TokenValue, MulExpr_ Expression) (Expression, error)
}

type BinaryCmpExprReducer interface {
	// 572:31: binary_cmp_expr -> ...
	ToBinaryCmpExpr(CmpExpr_ Expression, CmpOp_ TokenValue, AddExpr_ Expression) (Expression, error)
}

type BinaryAndExprReducer interface {
	// 586:31: binary_and_expr -> ...
	ToBinaryAndExpr(AndExpr_ Expression, And_ TokenValue, CmpExpr_ Expression) (Expression, error)
}

type BinaryOrExprReducer interface {
	// 592:30: binary_or_expr -> ...
	ToBinaryOrExpr(OrExpr_ Expression, Or_ TokenValue, AndExpr_ Expression) (Expression, error)
}

type SliceTypeExprReducer interface {
	// 607:35: slice_type_expr -> ...
	ToSliceTypeExpr(Lbracket_ TokenValue, TypeExpr_ TypeExpression, Rbracket_ TokenValue) (TypeExpression, error)
}

type ArrayTypeExprReducer interface {
	// 610:2: array_type_expr -> ...
	ToArrayTypeExpr(Lbracket_ TokenValue, TypeExpr_ TypeExpression, Comma_ TokenValue, IntegerLiteral_ TokenValue, Rbracket_ TokenValue) (TypeExpression, error)
}

type MapTypeExprReducer interface {
	// 613:33: map_type_expr -> ...
	ToMapTypeExpr(Lbracket_ TokenValue, TypeExpr_ TypeExpression, Colon_ TokenValue, TypeExpr_2 TypeExpression, Rbracket_ TokenValue) (TypeExpression, error)
}

type NamedTypeExprReducer interface {
	// 627:2: named_type_expr -> local: ...
	LocalToNamedTypeExpr(Identifier_ TokenValue, GenericTypeArguments_ *TypeArgumentList) (TypeExpression, error)

	// 628:2: named_type_expr -> external: ...
	ExternalToNamedTypeExpr(Identifier_ TokenValue, Dot_ TokenValue, Identifier_2 TokenValue, GenericTypeArguments_ *TypeArgumentList) (TypeExpression, error)
}

type InferredTypeExprReducer interface {
	// 636:2: inferred_type_expr -> DOT: ...
	DotToInferredTypeExpr(Dot_ TokenValue) (TypeExpression, error)

	// 637:2: inferred_type_expr -> UNDERSCORE: ...
	UnderscoreToInferredTypeExpr(Underscore_ TokenValue) (TypeExpression, error)
}

type ParseErrorTypeExprReducer interface {
	// 639:41: parse_error_type_expr -> ...
	ToParseErrorTypeExpr(ParseError_ ParseErrorSymbol) (TypeExpression, error)
}

type PrefixUnaryTypeExprReducer interface {
	// 649:2: prefix_unary_type_expr -> ...
	ToPrefixUnaryTypeExpr(PrefixUnaryTypeOp_ TokenValue, ReturnableTypeExpr_ TypeExpression) (TypeExpression, error)
}

type BinaryTypeExprReducer interface {
	// 665:2: binary_type_expr -> ...
	ToBinaryTypeExpr(TypeExpr_ TypeExpression, BinaryTypeOp_ TokenValue, ReturnableTypeExpr_ TypeExpression) (TypeExpression, error)
}

type TypeDefReducer interface {
	// 673:2: type_def -> definition: ...
	DefinitionToTypeDef(Type_ TokenValue, Identifier_ TokenValue, GenericParameterDefs_ GenericSymbol, TypeExpr_ TypeExpression) (SourceDefinition, error)

	// 674:2: type_def -> constrained_def: ...
	ConstrainedDefToTypeDef(Type_ TokenValue, Identifier_ TokenValue, GenericParameterDefs_ GenericSymbol, TypeExpr_ TypeExpression, Implements_ TokenValue, TypeExpr_2 TypeExpression) (SourceDefinition, error)

	// 675:2: type_def -> alias: ...
	AliasToTypeDef(Type_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, TypeExpr_ TypeExpression) (SourceDefinition, error)
}

type GenericParameterDefReducer interface {
	// 683:2: generic_parameter_def -> unconstrained: ...
	UnconstrainedToGenericParameterDef(Identifier_ TokenValue) (GenericSymbol, error)

	// 684:2: generic_parameter_def -> constrained: ...
	ConstrainedToGenericParameterDef(Identifier_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)
}

type GenericParameterDefsReducer interface {
	// 687:2: generic_parameter_defs -> generic: ...
	GenericToGenericParameterDefs(DollarLbracket_ TokenValue, GenericParameterDefList_ GenericSymbol, Rbracket_ TokenValue) (GenericSymbol, error)

	// 688:2: generic_parameter_defs -> nil: ...
	NilToGenericParameterDefs() (GenericSymbol, error)
}

type ProperGenericParameterDefListReducer interface {
	// 691:2: proper_generic_parameter_def_list -> add: ...
	AddToProperGenericParameterDefList(ProperGenericParameterDefList_ GenericSymbol, Comma_ TokenValue, GenericParameterDef_ GenericSymbol) (GenericSymbol, error)

	// 692:2: proper_generic_parameter_def_list -> generic_parameter_def: ...
	GenericParameterDefToProperGenericParameterDefList(GenericParameterDef_ GenericSymbol) (GenericSymbol, error)
}

type GenericParameterDefListReducer interface {

	// 696:2: generic_parameter_def_list -> improper: ...
	ImproperToGenericParameterDefList(ProperGenericParameterDefList_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 697:2: generic_parameter_def_list -> nil: ...
	NilToGenericParameterDefList() (GenericSymbol, error)
}

type GenericTypeArgumentsReducer interface {
	// 700:2: generic_type_arguments -> binding: ...
	BindingToGenericTypeArguments(DollarLbracket_ TokenValue, TypeArguments_ *TypeArgumentList, Rbracket_ TokenValue) (*TypeArgumentList, error)

	// 701:2: generic_type_arguments -> nil: ...
	NilToGenericTypeArguments() (*TypeArgumentList, error)
}

type ProperTypeArgumentsReducer interface {
	// 704:2: proper_type_arguments -> add: ...
	AddToProperTypeArguments(ProperTypeArguments_ *TypeArgumentList, Comma_ TokenValue, TypeExpr_ TypeExpression) (*TypeArgumentList, error)

	// 705:2: proper_type_arguments -> type_expr: ...
	TypeExprToProperTypeArguments(TypeExpr_ TypeExpression) (*TypeArgumentList, error)
}

type TypeArgumentsReducer interface {

	// 709:2: type_arguments -> improper: ...
	ImproperToTypeArguments(ProperTypeArguments_ *TypeArgumentList, Comma_ TokenValue) (*TypeArgumentList, error)

	// 710:2: type_arguments -> nil: ...
	NilToTypeArguments() (*TypeArgumentList, error)
}

type FieldDefReducer interface {
	// 724:2: field_def -> named: ...
	NamedToFieldDef(Identifier_ TokenValue, TypeExpr_ TypeExpression, OptionalDefault_ GenericSymbol) (GenericSymbol, error)

	// 725:2: field_def -> unnamed: ...
	UnnamedToFieldDef(TypeExpr_ TypeExpression, OptionalDefault_ GenericSymbol) (GenericSymbol, error)

	// 726:2: field_def -> struct_padding: ...
	StructPaddingToFieldDef(Underscore_ TokenValue, TypeExpr_ TypeExpression) (GenericSymbol, error)

	// 728:2: field_def -> unsafe_statement: ...
	UnsafeStatementToFieldDef(UnsafeStatement_ Statement) (GenericSymbol, error)
}

type OptionalDefaultReducer interface {
	// 731:2: optional_default -> default: ...
	DefaultToOptionalDefault(Assign_ TokenValue, Default_ TokenValue) (GenericSymbol, error)

	// 732:2: optional_default -> nil: ...
	NilToOptionalDefault() (GenericSymbol, error)
}

type ProperImplicitFieldDefsReducer interface {
	// 735:2: proper_implicit_field_defs -> add: ...
	AddToProperImplicitFieldDefs(ProperImplicitFieldDefs_ GenericSymbol, Comma_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 736:2: proper_implicit_field_defs -> field_def: ...
	FieldDefToProperImplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type ImplicitFieldDefsReducer interface {

	// 740:2: implicit_field_defs -> improper: ...
	ImproperToImplicitFieldDefs(ProperImplicitFieldDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 741:2: implicit_field_defs -> nil: ...
	NilToImplicitFieldDefs() (GenericSymbol, error)
}

type ImplicitStructTypeExprReducer interface {
	// 744:2: implicit_struct_type_expr -> ...
	ToImplicitStructTypeExpr(Lparen_ TokenValue, ImplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ProperExplicitFieldDefsReducer interface {
	// 747:2: proper_explicit_field_defs -> add_implicit: ...
	AddImplicitToProperExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Newlines_ TokenCount, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 748:2: proper_explicit_field_defs -> add_explicit: ...
	AddExplicitToProperExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Comma_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 749:2: proper_explicit_field_defs -> field_def: ...
	FieldDefToProperExplicitFieldDefs(FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type ExplicitFieldDefsReducer interface {

	// 753:2: explicit_field_defs -> improper_implicit: ...
	ImproperImplicitToExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)

	// 754:2: explicit_field_defs -> improper_explicit: ...
	ImproperExplicitToExplicitFieldDefs(ProperExplicitFieldDefs_ GenericSymbol, Comma_ TokenValue) (GenericSymbol, error)

	// 755:2: explicit_field_defs -> nil: ...
	NilToExplicitFieldDefs() (GenericSymbol, error)
}

type ExplicitStructTypeExprReducer interface {
	// 758:2: explicit_struct_type_expr -> ...
	ToExplicitStructTypeExpr(Struct_ TokenValue, Lparen_ TokenValue, ExplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type TraitTypeExprReducer interface {
	// 761:2: trait_type_expr -> ...
	ToTraitTypeExpr(Trait_ TokenValue, Lparen_ TokenValue, ExplicitFieldDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ProperImplicitEnumValueDefsReducer interface {
	// 772:2: proper_implicit_enum_value_defs -> pair: ...
	PairToProperImplicitEnumValueDefs(FieldDef_ GenericSymbol, Or_ TokenValue, FieldDef_2 GenericSymbol) (GenericSymbol, error)

	// 773:2: proper_implicit_enum_value_defs -> add: ...
	AddToProperImplicitEnumValueDefs(ProperImplicitEnumValueDefs_ GenericSymbol, Or_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type ImplicitEnumValueDefsReducer interface {

	// 778:2: implicit_enum_value_defs -> improper: ...
	ImproperToImplicitEnumValueDefs(ProperImplicitEnumValueDefs_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)
}

type ImplicitEnumTypeExprReducer interface {
	// 781:2: implicit_enum_type_expr -> ...
	ToImplicitEnumTypeExpr(Lparen_ TokenValue, ImplicitEnumValueDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ProperExplicitEnumValueDefsReducer interface {
	// 784:2: proper_explicit_enum_value_defs -> explicit_pair: ...
	ExplicitPairToProperExplicitEnumValueDefs(FieldDef_ GenericSymbol, Or_ TokenValue, FieldDef_2 GenericSymbol) (GenericSymbol, error)

	// 785:2: proper_explicit_enum_value_defs -> implicit_pair: ...
	ImplicitPairToProperExplicitEnumValueDefs(FieldDef_ GenericSymbol, Newlines_ TokenCount, FieldDef_2 GenericSymbol) (GenericSymbol, error)

	// 786:2: proper_explicit_enum_value_defs -> explicit_add: ...
	ExplicitAddToProperExplicitEnumValueDefs(ProperExplicitEnumValueDefs_ GenericSymbol, Or_ TokenValue, FieldDef_ GenericSymbol) (GenericSymbol, error)

	// 787:2: proper_explicit_enum_value_defs -> implicit_add: ...
	ImplicitAddToProperExplicitEnumValueDefs(ProperExplicitEnumValueDefs_ GenericSymbol, Newlines_ TokenCount, FieldDef_ GenericSymbol) (GenericSymbol, error)
}

type ExplicitEnumValueDefsReducer interface {

	// 792:2: explicit_enum_value_defs -> improper: ...
	ImproperToExplicitEnumValueDefs(ProperExplicitEnumValueDefs_ GenericSymbol, Newlines_ TokenCount) (GenericSymbol, error)
}

type ExplicitEnumTypeExprReducer interface {
	// 795:2: explicit_enum_type_expr -> ...
	ToExplicitEnumTypeExpr(Enum_ TokenValue, Lparen_ TokenValue, ExplicitEnumValueDefs_ GenericSymbol, Rparen_ TokenValue) (TypeExpression, error)
}

type ReturnTypeReducer interface {
	// 803:2: return_type -> returnable_type_expr: ...
	ReturnableTypeExprToReturnType(ReturnableTypeExpr_ TypeExpression) (TypeExpression, error)

	// 804:2: return_type -> nil: ...
	NilToReturnType() (TypeExpression, error)
}

type ProperParameterDefReducer interface {
	// 807:2: proper_parameter_def -> named_typed_arg: ...
	NamedTypedArgToProperParameterDef(Identifier_ TokenValue, TypeExpr_ TypeExpression) (*Parameter, error)

	// 808:2: proper_parameter_def -> named_typed_vararg: ...
	NamedTypedVarargToProperParameterDef(Identifier_ TokenValue, Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (*Parameter, error)

	// 809:2: proper_parameter_def -> named_inferred_vararg: ...
	NamedInferredVarargToProperParameterDef(Identifier_ TokenValue, Ellipsis_ TokenValue) (*Parameter, error)

	// 810:2: proper_parameter_def -> ignore_typed_arg: ...
	IgnoreTypedArgToProperParameterDef(Underscore_ TokenValue, TypeExpr_ TypeExpression) (*Parameter, error)

	// 811:2: proper_parameter_def -> ignore_inferred_vararg: ...
	IgnoreInferredVarargToProperParameterDef(Underscore_ TokenValue, Ellipsis_ TokenValue) (*Parameter, error)

	// 812:2: proper_parameter_def -> ignore_typed_vararg: ...
	IgnoreTypedVarargToProperParameterDef(Underscore_ TokenValue, Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (*Parameter, error)
}

type ParameterDeclReducer interface {

	// 818:2: parameter_decl -> unnamed_typed_arg: ...
	UnnamedTypedArgToParameterDecl(TypeExpr_ TypeExpression) (*Parameter, error)

	// 819:2: parameter_decl -> unnamed_inferred_vararg: ...
	UnnamedInferredVarargToParameterDecl(Ellipsis_ TokenValue) (*Parameter, error)

	// 820:2: parameter_decl -> unnamed_typed_vararg: ...
	UnnamedTypedVarargToParameterDecl(Ellipsis_ TokenValue, TypeExpr_ TypeExpression) (*Parameter, error)
}

type ParameterDefReducer interface {

	// 828:2: parameter_def -> named_inferred_arg: ...
	NamedInferredArgToParameterDef(Identifier_ TokenValue) (*Parameter, error)

	// 829:2: parameter_def -> ignore_inferred_arg: ...
	IgnoreInferredArgToParameterDef(Underscore_ TokenValue) (*Parameter, error)
}

type ProperParameterDeclsReducer interface {
	// 832:2: proper_parameter_decls -> add: ...
	AddToProperParameterDecls(ProperParameterDecls_ *ParameterList, Comma_ TokenValue, ParameterDecl_ *Parameter) (*ParameterList, error)

	// 833:2: proper_parameter_decls -> parameter_decl: ...
	ParameterDeclToProperParameterDecls(ParameterDecl_ *Parameter) (*ParameterList, error)
}

type ParameterDeclsReducer interface {

	// 837:2: parameter_decls -> improper: ...
	ImproperToParameterDecls(ProperParameterDecls_ *ParameterList, Comma_ TokenValue) (*ParameterList, error)

	// 838:2: parameter_decls -> nil: ...
	NilToParameterDecls() (*ParameterList, error)
}

type ProperParameterDefsReducer interface {
	// 841:2: proper_parameter_defs -> add: ...
	AddToProperParameterDefs(ProperParameterDefs_ *ParameterList, Comma_ TokenValue, ParameterDef_ *Parameter) (*ParameterList, error)

	// 842:2: proper_parameter_defs -> parameter_def: ...
	ParameterDefToProperParameterDefs(ParameterDef_ *Parameter) (*ParameterList, error)
}

type ParameterDefsReducer interface {

	// 846:2: parameter_defs -> improper: ...
	ImproperToParameterDefs(ProperParameterDefs_ *ParameterList, Comma_ TokenValue) (*ParameterList, error)

	// 847:2: parameter_defs -> nil: ...
	NilToParameterDefs() (*ParameterList, error)
}

type FuncTypeExprReducer interface {
	// 850:2: func_type_expr -> ...
	ToFuncTypeExpr(Func_ TokenValue, Lparen_ TokenValue, ParameterDecls_ *ParameterList, Rparen_ TokenValue, ReturnType_ TypeExpression) (TypeExpression, error)
}

type MethodSignatureReducer interface {
	// 861:20: method_signature -> ...
	ToMethodSignature(Func_ TokenValue, Identifier_ TokenValue, Lparen_ TokenValue, ParameterDecls_ *ParameterList, Rparen_ TokenValue, ReturnType_ TypeExpression) (GenericSymbol, error)
}

type NamedFuncDefReducer interface {
	// 865:2: named_func_def -> func_def: ...
	FuncDefToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, GenericParameterDefs_ GenericSymbol, Lparen_ TokenValue, ParameterDefs_ *ParameterList, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 866:2: named_func_def -> method_def: ...
	MethodDefToNamedFuncDef(Func_ TokenValue, Lparen_ TokenValue, ParameterDef_ *Parameter, Rparen_ TokenValue, Identifier_ TokenValue, Lparen_2 TokenValue, ParameterDefs_ *ParameterList, Rparen_2 TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (SourceDefinition, error)

	// 867:2: named_func_def -> alias: ...
	AliasToNamedFuncDef(Func_ TokenValue, Identifier_ TokenValue, Assign_ TokenValue, Expr_ Expression) (SourceDefinition, error)
}

type AnonymousFuncExprReducer interface {
	// 870:2: anonymous_func_expr -> ...
	ToAnonymousFuncExpr(Func_ TokenValue, Lparen_ TokenValue, ParameterDefs_ *ParameterList, Rparen_ TokenValue, ReturnType_ TypeExpression, StatementBlock_ GenericSymbol) (Expression, error)
}

type PackageDefReducer interface {
	// 882:2: package_def -> no_spec: ...
	NoSpecToPackageDef(Package_ TokenValue) (SourceDefinition, error)

	// 883:2: package_def -> with_spec: ...
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
	OptionalLabelDeclReducer
	SequenceExprReducer
	IfExprReducer
	IfExprBodyReducer
	ConditionReducer
	SwitchExprReducer
	LoopExprReducer
	LoopExprBodyReducer
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
	case IfExprBodyType:
		return "if_expr_body"
	case ConditionType:
		return "condition"
	case SwitchExprType:
		return "switch_expr"
	case LoopExprType:
		return "loop_expr"
	case LoopExprBodyType:
		return "loop_expr_body"
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
	case PostfixUnaryOpType:
		return "postfix_unary_op"
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
	IfExprBodyType                    = SymbolId(384)
	ConditionType                     = SymbolId(385)
	SwitchExprType                    = SymbolId(386)
	LoopExprType                      = SymbolId(387)
	LoopExprBodyType                  = SymbolId(388)
	OptionalSequenceExprType          = SymbolId(389)
	ForAssignmentType                 = SymbolId(390)
	CallExprType                      = SymbolId(391)
	ProperArgumentsType               = SymbolId(392)
	ArgumentsType                     = SymbolId(393)
	ArgumentType                      = SymbolId(394)
	ColonExprType                     = SymbolId(395)
	AtomExprType                      = SymbolId(396)
	ParseErrorExprType                = SymbolId(397)
	LiteralExprType                   = SymbolId(398)
	NamedExprType                     = SymbolId(399)
	BlockExprType                     = SymbolId(400)
	InitializeExprType                = SymbolId(401)
	ImplicitStructExprType            = SymbolId(402)
	AccessibleExprType                = SymbolId(403)
	AccessExprType                    = SymbolId(404)
	IndexExprType                     = SymbolId(405)
	PostfixableExprType               = SymbolId(406)
	PostfixUnaryOpType                = SymbolId(407)
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
	PrefixUnaryTypeExprType           = SymbolId(434)
	PrefixUnaryTypeOpType             = SymbolId(435)
	TypeExprType                      = SymbolId(436)
	BinaryTypeExprType                = SymbolId(437)
	BinaryTypeOpType                  = SymbolId(438)
	TypeDefType                       = SymbolId(439)
	GenericParameterDefType           = SymbolId(440)
	GenericParameterDefsType          = SymbolId(441)
	ProperGenericParameterDefListType = SymbolId(442)
	GenericParameterDefListType       = SymbolId(443)
	GenericTypeArgumentsType          = SymbolId(444)
	ProperTypeArgumentsType           = SymbolId(445)
	TypeArgumentsType                 = SymbolId(446)
	FieldDefType                      = SymbolId(447)
	OptionalDefaultType               = SymbolId(448)
	ProperImplicitFieldDefsType       = SymbolId(449)
	ImplicitFieldDefsType             = SymbolId(450)
	ImplicitStructTypeExprType        = SymbolId(451)
	ProperExplicitFieldDefsType       = SymbolId(452)
	ExplicitFieldDefsType             = SymbolId(453)
	ExplicitStructTypeExprType        = SymbolId(454)
	TraitTypeExprType                 = SymbolId(455)
	ProperImplicitEnumValueDefsType   = SymbolId(456)
	ImplicitEnumValueDefsType         = SymbolId(457)
	ImplicitEnumTypeExprType          = SymbolId(458)
	ProperExplicitEnumValueDefsType   = SymbolId(459)
	ExplicitEnumValueDefsType         = SymbolId(460)
	ExplicitEnumTypeExprType          = SymbolId(461)
	ReturnTypeType                    = SymbolId(462)
	ProperParameterDefType            = SymbolId(463)
	ParameterDeclType                 = SymbolId(464)
	ParameterDefType                  = SymbolId(465)
	ProperParameterDeclsType          = SymbolId(466)
	ParameterDeclsType                = SymbolId(467)
	ProperParameterDefsType           = SymbolId(468)
	ParameterDefsType                 = SymbolId(469)
	FuncTypeExprType                  = SymbolId(470)
	MethodSignatureType               = SymbolId(471)
	NamedFuncDefType                  = SymbolId(472)
	AnonymousFuncExprType             = SymbolId(473)
	PackageDefType                    = SymbolId(474)
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
	_ReduceToIfExpr                                               = _ReduceType(109)
	_ReduceNoElseToIfExprBody                                     = _ReduceType(110)
	_ReduceIfElseToIfExprBody                                     = _ReduceType(111)
	_ReduceMultiIfElseToIfExprBody                                = _ReduceType(112)
	_ReduceSequenceExprToCondition                                = _ReduceType(113)
	_ReduceCaseToCondition                                        = _ReduceType(114)
	_ReduceToSwitchExpr                                           = _ReduceType(115)
	_ReduceToLoopExpr                                             = _ReduceType(116)
	_ReduceInfiniteToLoopExprBody                                 = _ReduceType(117)
	_ReduceDoWhileToLoopExprBody                                  = _ReduceType(118)
	_ReduceWhileToLoopExprBody                                    = _ReduceType(119)
	_ReduceIteratorToLoopExprBody                                 = _ReduceType(120)
	_ReduceForToLoopExprBody                                      = _ReduceType(121)
	_ReduceSequenceExprToOptionalSequenceExpr                     = _ReduceType(122)
	_ReduceNilToOptionalSequenceExpr                              = _ReduceType(123)
	_ReduceSequenceExprToForAssignment                            = _ReduceType(124)
	_ReduceAssignToForAssignment                                  = _ReduceType(125)
	_ReduceToCallExpr                                             = _ReduceType(126)
	_ReduceAddToProperArguments                                   = _ReduceType(127)
	_ReduceArgumentToProperArguments                              = _ReduceType(128)
	_ReduceProperArgumentsToArguments                             = _ReduceType(129)
	_ReduceImproperToArguments                                    = _ReduceType(130)
	_ReduceNilToArguments                                         = _ReduceType(131)
	_ReducePositionalToArgument                                   = _ReduceType(132)
	_ReduceColonExprToArgument                                    = _ReduceType(133)
	_ReduceNamedAssignmentToArgument                              = _ReduceType(134)
	_ReduceVarargAssignmentToArgument                             = _ReduceType(135)
	_ReduceSkipPatternToArgument                                  = _ReduceType(136)
	_ReduceUnitUnitPairToColonExpr                                = _ReduceType(137)
	_ReduceExprUnitPairToColonExpr                                = _ReduceType(138)
	_ReduceUnitExprPairToColonExpr                                = _ReduceType(139)
	_ReduceExprExprPairToColonExpr                                = _ReduceType(140)
	_ReduceColonExprUnitTupleToColonExpr                          = _ReduceType(141)
	_ReduceColonExprExprTupleToColonExpr                          = _ReduceType(142)
	_ReduceParseErrorExprToAtomExpr                               = _ReduceType(143)
	_ReduceLiteralExprToAtomExpr                                  = _ReduceType(144)
	_ReduceNamedExprToAtomExpr                                    = _ReduceType(145)
	_ReduceBlockExprToAtomExpr                                    = _ReduceType(146)
	_ReduceAnonymousFuncExprToAtomExpr                            = _ReduceType(147)
	_ReduceInitializeExprToAtomExpr                               = _ReduceType(148)
	_ReduceImplicitStructExprToAtomExpr                           = _ReduceType(149)
	_ReduceToParseErrorExpr                                       = _ReduceType(150)
	_ReduceTrueToLiteralExpr                                      = _ReduceType(151)
	_ReduceFalseToLiteralExpr                                     = _ReduceType(152)
	_ReduceIntegerLiteralToLiteralExpr                            = _ReduceType(153)
	_ReduceFloatLiteralToLiteralExpr                              = _ReduceType(154)
	_ReduceRuneLiteralToLiteralExpr                               = _ReduceType(155)
	_ReduceStringLiteralToLiteralExpr                             = _ReduceType(156)
	_ReduceIdentifierToNamedExpr                                  = _ReduceType(157)
	_ReduceUnderscoreToNamedExpr                                  = _ReduceType(158)
	_ReduceToBlockExpr                                            = _ReduceType(159)
	_ReduceToInitializeExpr                                       = _ReduceType(160)
	_ReduceToImplicitStructExpr                                   = _ReduceType(161)
	_ReduceAtomExprToAccessibleExpr                               = _ReduceType(162)
	_ReduceAccessExprToAccessibleExpr                             = _ReduceType(163)
	_ReduceCallExprToAccessibleExpr                               = _ReduceType(164)
	_ReduceIndexExprToAccessibleExpr                              = _ReduceType(165)
	_ReduceToAccessExpr                                           = _ReduceType(166)
	_ReduceToIndexExpr                                            = _ReduceType(167)
	_ReduceAccessibleExprToPostfixableExpr                        = _ReduceType(168)
	_ReducePostfixUnaryExprToPostfixableExpr                      = _ReduceType(169)
	_ReduceQuestionToPostfixUnaryOp                               = _ReduceType(170)
	_ReduceExclaimToPostfixUnaryOp                                = _ReduceType(171)
	_ReduceToPostfixUnaryExpr                                     = _ReduceType(172)
	_ReducePostfixableExprToPrefixableExpr                        = _ReduceType(173)
	_ReducePrefixUnaryExprToPrefixableExpr                        = _ReduceType(174)
	_ReduceToPrefixUnaryExpr                                      = _ReduceType(175)
	_ReduceNotToPrefixUnaryOp                                     = _ReduceType(176)
	_ReduceBitNegToPrefixUnaryOp                                  = _ReduceType(177)
	_ReduceSubToPrefixUnaryOp                                     = _ReduceType(178)
	_ReduceMulToPrefixUnaryOp                                     = _ReduceType(179)
	_ReduceBitAndToPrefixUnaryOp                                  = _ReduceType(180)
	_ReducePrefixableExprToMulExpr                                = _ReduceType(181)
	_ReduceBinaryMulExprToMulExpr                                 = _ReduceType(182)
	_ReduceToBinaryMulExpr                                        = _ReduceType(183)
	_ReduceMulToMulOp                                             = _ReduceType(184)
	_ReduceDivToMulOp                                             = _ReduceType(185)
	_ReduceModToMulOp                                             = _ReduceType(186)
	_ReduceBitAndToMulOp                                          = _ReduceType(187)
	_ReduceBitLshiftToMulOp                                       = _ReduceType(188)
	_ReduceBitRshiftToMulOp                                       = _ReduceType(189)
	_ReduceMulExprToAddExpr                                       = _ReduceType(190)
	_ReduceBinaryAddExprToAddExpr                                 = _ReduceType(191)
	_ReduceToBinaryAddExpr                                        = _ReduceType(192)
	_ReduceAddToAddOp                                             = _ReduceType(193)
	_ReduceSubToAddOp                                             = _ReduceType(194)
	_ReduceBitOrToAddOp                                           = _ReduceType(195)
	_ReduceBitXorToAddOp                                          = _ReduceType(196)
	_ReduceAddExprToCmpExpr                                       = _ReduceType(197)
	_ReduceBinaryCmpExprToCmpExpr                                 = _ReduceType(198)
	_ReduceToBinaryCmpExpr                                        = _ReduceType(199)
	_ReduceEqualToCmpOp                                           = _ReduceType(200)
	_ReduceNotEqualToCmpOp                                        = _ReduceType(201)
	_ReduceLessToCmpOp                                            = _ReduceType(202)
	_ReduceLessOrEqualToCmpOp                                     = _ReduceType(203)
	_ReduceGreaterToCmpOp                                         = _ReduceType(204)
	_ReduceGreaterOrEqualToCmpOp                                  = _ReduceType(205)
	_ReduceCmpExprToAndExpr                                       = _ReduceType(206)
	_ReduceBinaryAndExprToAndExpr                                 = _ReduceType(207)
	_ReduceToBinaryAndExpr                                        = _ReduceType(208)
	_ReduceAndExprToOrExpr                                        = _ReduceType(209)
	_ReduceBinaryOrExprToOrExpr                                   = _ReduceType(210)
	_ReduceToBinaryOrExpr                                         = _ReduceType(211)
	_ReduceExplicitStructTypeExprToInitializableTypeExpr          = _ReduceType(212)
	_ReduceSliceTypeExprToInitializableTypeExpr                   = _ReduceType(213)
	_ReduceArrayTypeExprToInitializableTypeExpr                   = _ReduceType(214)
	_ReduceMapTypeExprToInitializableTypeExpr                     = _ReduceType(215)
	_ReduceToSliceTypeExpr                                        = _ReduceType(216)
	_ReduceToArrayTypeExpr                                        = _ReduceType(217)
	_ReduceToMapTypeExpr                                          = _ReduceType(218)
	_ReduceInitializableTypeExprToAtomTypeExpr                    = _ReduceType(219)
	_ReduceNamedTypeExprToAtomTypeExpr                            = _ReduceType(220)
	_ReduceInferredTypeExprToAtomTypeExpr                         = _ReduceType(221)
	_ReduceImplicitStructTypeExprToAtomTypeExpr                   = _ReduceType(222)
	_ReduceExplicitEnumTypeExprToAtomTypeExpr                     = _ReduceType(223)
	_ReduceImplicitEnumTypeExprToAtomTypeExpr                     = _ReduceType(224)
	_ReduceTraitTypeExprToAtomTypeExpr                            = _ReduceType(225)
	_ReduceFuncTypeExprToAtomTypeExpr                             = _ReduceType(226)
	_ReduceParseErrorTypeExprToAtomTypeExpr                       = _ReduceType(227)
	_ReduceLocalToNamedTypeExpr                                   = _ReduceType(228)
	_ReduceExternalToNamedTypeExpr                                = _ReduceType(229)
	_ReduceDotToInferredTypeExpr                                  = _ReduceType(230)
	_ReduceUnderscoreToInferredTypeExpr                           = _ReduceType(231)
	_ReduceToParseErrorTypeExpr                                   = _ReduceType(232)
	_ReduceAtomTypeExprToReturnableTypeExpr                       = _ReduceType(233)
	_ReducePrefixUnaryTypeExprToReturnableTypeExpr                = _ReduceType(234)
	_ReduceToPrefixUnaryTypeExpr                                  = _ReduceType(235)
	_ReduceQuestionToPrefixUnaryTypeOp                            = _ReduceType(236)
	_ReduceExclaimToPrefixUnaryTypeOp                             = _ReduceType(237)
	_ReduceBitAndToPrefixUnaryTypeOp                              = _ReduceType(238)
	_ReduceBitNegToPrefixUnaryTypeOp                              = _ReduceType(239)
	_ReduceTildeTildeToPrefixUnaryTypeOp                          = _ReduceType(240)
	_ReduceReturnableTypeExprToTypeExpr                           = _ReduceType(241)
	_ReduceBinaryTypeExprToTypeExpr                               = _ReduceType(242)
	_ReduceToBinaryTypeExpr                                       = _ReduceType(243)
	_ReduceMulToBinaryTypeOp                                      = _ReduceType(244)
	_ReduceAddToBinaryTypeOp                                      = _ReduceType(245)
	_ReduceSubToBinaryTypeOp                                      = _ReduceType(246)
	_ReduceDefinitionToTypeDef                                    = _ReduceType(247)
	_ReduceConstrainedDefToTypeDef                                = _ReduceType(248)
	_ReduceAliasToTypeDef                                         = _ReduceType(249)
	_ReduceUnconstrainedToGenericParameterDef                     = _ReduceType(250)
	_ReduceConstrainedToGenericParameterDef                       = _ReduceType(251)
	_ReduceGenericToGenericParameterDefs                          = _ReduceType(252)
	_ReduceNilToGenericParameterDefs                              = _ReduceType(253)
	_ReduceAddToProperGenericParameterDefList                     = _ReduceType(254)
	_ReduceGenericParameterDefToProperGenericParameterDefList     = _ReduceType(255)
	_ReduceProperGenericParameterDefListToGenericParameterDefList = _ReduceType(256)
	_ReduceImproperToGenericParameterDefList                      = _ReduceType(257)
	_ReduceNilToGenericParameterDefList                           = _ReduceType(258)
	_ReduceBindingToGenericTypeArguments                          = _ReduceType(259)
	_ReduceNilToGenericTypeArguments                              = _ReduceType(260)
	_ReduceAddToProperTypeArguments                               = _ReduceType(261)
	_ReduceTypeExprToProperTypeArguments                          = _ReduceType(262)
	_ReduceProperTypeArgumentsToTypeArguments                     = _ReduceType(263)
	_ReduceImproperToTypeArguments                                = _ReduceType(264)
	_ReduceNilToTypeArguments                                     = _ReduceType(265)
	_ReduceNamedToFieldDef                                        = _ReduceType(266)
	_ReduceUnnamedToFieldDef                                      = _ReduceType(267)
	_ReduceStructPaddingToFieldDef                                = _ReduceType(268)
	_ReduceMethodSignatureToFieldDef                              = _ReduceType(269)
	_ReduceUnsafeStatementToFieldDef                              = _ReduceType(270)
	_ReduceDefaultToOptionalDefault                               = _ReduceType(271)
	_ReduceNilToOptionalDefault                                   = _ReduceType(272)
	_ReduceAddToProperImplicitFieldDefs                           = _ReduceType(273)
	_ReduceFieldDefToProperImplicitFieldDefs                      = _ReduceType(274)
	_ReduceProperImplicitFieldDefsToImplicitFieldDefs             = _ReduceType(275)
	_ReduceImproperToImplicitFieldDefs                            = _ReduceType(276)
	_ReduceNilToImplicitFieldDefs                                 = _ReduceType(277)
	_ReduceToImplicitStructTypeExpr                               = _ReduceType(278)
	_ReduceAddImplicitToProperExplicitFieldDefs                   = _ReduceType(279)
	_ReduceAddExplicitToProperExplicitFieldDefs                   = _ReduceType(280)
	_ReduceFieldDefToProperExplicitFieldDefs                      = _ReduceType(281)
	_ReduceProperExplicitFieldDefsToExplicitFieldDefs             = _ReduceType(282)
	_ReduceImproperImplicitToExplicitFieldDefs                    = _ReduceType(283)
	_ReduceImproperExplicitToExplicitFieldDefs                    = _ReduceType(284)
	_ReduceNilToExplicitFieldDefs                                 = _ReduceType(285)
	_ReduceToExplicitStructTypeExpr                               = _ReduceType(286)
	_ReduceToTraitTypeExpr                                        = _ReduceType(287)
	_ReducePairToProperImplicitEnumValueDefs                      = _ReduceType(288)
	_ReduceAddToProperImplicitEnumValueDefs                       = _ReduceType(289)
	_ReduceProperImplicitEnumValueDefsToImplicitEnumValueDefs     = _ReduceType(290)
	_ReduceImproperToImplicitEnumValueDefs                        = _ReduceType(291)
	_ReduceToImplicitEnumTypeExpr                                 = _ReduceType(292)
	_ReduceExplicitPairToProperExplicitEnumValueDefs              = _ReduceType(293)
	_ReduceImplicitPairToProperExplicitEnumValueDefs              = _ReduceType(294)
	_ReduceExplicitAddToProperExplicitEnumValueDefs               = _ReduceType(295)
	_ReduceImplicitAddToProperExplicitEnumValueDefs               = _ReduceType(296)
	_ReduceProperExplicitEnumValueDefsToExplicitEnumValueDefs     = _ReduceType(297)
	_ReduceImproperToExplicitEnumValueDefs                        = _ReduceType(298)
	_ReduceToExplicitEnumTypeExpr                                 = _ReduceType(299)
	_ReduceReturnableTypeExprToReturnType                         = _ReduceType(300)
	_ReduceNilToReturnType                                        = _ReduceType(301)
	_ReduceNamedTypedArgToProperParameterDef                      = _ReduceType(302)
	_ReduceNamedTypedVarargToProperParameterDef                   = _ReduceType(303)
	_ReduceNamedInferredVarargToProperParameterDef                = _ReduceType(304)
	_ReduceIgnoreTypedArgToProperParameterDef                     = _ReduceType(305)
	_ReduceIgnoreInferredVarargToProperParameterDef               = _ReduceType(306)
	_ReduceIgnoreTypedVarargToProperParameterDef                  = _ReduceType(307)
	_ReduceProperParameterDefToParameterDecl                      = _ReduceType(308)
	_ReduceUnnamedTypedArgToParameterDecl                         = _ReduceType(309)
	_ReduceUnnamedInferredVarargToParameterDecl                   = _ReduceType(310)
	_ReduceUnnamedTypedVarargToParameterDecl                      = _ReduceType(311)
	_ReduceProperParameterDefToParameterDef                       = _ReduceType(312)
	_ReduceNamedInferredArgToParameterDef                         = _ReduceType(313)
	_ReduceIgnoreInferredArgToParameterDef                        = _ReduceType(314)
	_ReduceAddToProperParameterDecls                              = _ReduceType(315)
	_ReduceParameterDeclToProperParameterDecls                    = _ReduceType(316)
	_ReduceProperParameterDeclsToParameterDecls                   = _ReduceType(317)
	_ReduceImproperToParameterDecls                               = _ReduceType(318)
	_ReduceNilToParameterDecls                                    = _ReduceType(319)
	_ReduceAddToProperParameterDefs                               = _ReduceType(320)
	_ReduceParameterDefToProperParameterDefs                      = _ReduceType(321)
	_ReduceProperParameterDefsToParameterDefs                     = _ReduceType(322)
	_ReduceImproperToParameterDefs                                = _ReduceType(323)
	_ReduceNilToParameterDefs                                     = _ReduceType(324)
	_ReduceToFuncTypeExpr                                         = _ReduceType(325)
	_ReduceToMethodSignature                                      = _ReduceType(326)
	_ReduceFuncDefToNamedFuncDef                                  = _ReduceType(327)
	_ReduceMethodDefToNamedFuncDef                                = _ReduceType(328)
	_ReduceAliasToNamedFuncDef                                    = _ReduceType(329)
	_ReduceToAnonymousFuncExpr                                    = _ReduceType(330)
	_ReduceNoSpecToPackageDef                                     = _ReduceType(331)
	_ReduceWithSpecToPackageDef                                   = _ReduceType(332)
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
	case _ReduceToIfExpr:
		return "ToIfExpr"
	case _ReduceNoElseToIfExprBody:
		return "NoElseToIfExprBody"
	case _ReduceIfElseToIfExprBody:
		return "IfElseToIfExprBody"
	case _ReduceMultiIfElseToIfExprBody:
		return "MultiIfElseToIfExprBody"
	case _ReduceSequenceExprToCondition:
		return "SequenceExprToCondition"
	case _ReduceCaseToCondition:
		return "CaseToCondition"
	case _ReduceToSwitchExpr:
		return "ToSwitchExpr"
	case _ReduceToLoopExpr:
		return "ToLoopExpr"
	case _ReduceInfiniteToLoopExprBody:
		return "InfiniteToLoopExprBody"
	case _ReduceDoWhileToLoopExprBody:
		return "DoWhileToLoopExprBody"
	case _ReduceWhileToLoopExprBody:
		return "WhileToLoopExprBody"
	case _ReduceIteratorToLoopExprBody:
		return "IteratorToLoopExprBody"
	case _ReduceForToLoopExprBody:
		return "ForToLoopExprBody"
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
	case _ReduceQuestionToPostfixUnaryOp:
		return "QuestionToPostfixUnaryOp"
	case _ReduceExclaimToPostfixUnaryOp:
		return "ExclaimToPostfixUnaryOp"
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
	case ExprType, SequenceExprType, IfExprType, IfExprBodyType, SwitchExprType, LoopExprType, LoopExprBodyType, OptionalSequenceExprType, CallExprType, AtomExprType, ParseErrorExprType, LiteralExprType, NamedExprType, BlockExprType, InitializeExprType, ImplicitStructExprType, AccessibleExprType, AccessExprType, IndexExprType, PostfixableExprType, PostfixUnaryExprType, PrefixableExprType, PrefixUnaryExprType, MulExprType, BinaryMulExprType, AddExprType, BinaryAddExprType, CmpExprType, BinaryCmpExprType, AndExprType, BinaryAndExprType, OrExprType, BinaryOrExprType, AnonymousFuncExprType:
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
	case CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, EllipsisToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, CallbackOpType, JumpTypeType, UnaryOpAssignType, BinaryOpAssignType, VarOrLetType, OptionalLabelDeclType, PostfixUnaryOpType, PrefixUnaryOpType, MulOpType, AddOpType, CmpOpType, PrefixUnaryTypeOpType, BinaryTypeOpType:
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
	case ExprType, SequenceExprType, IfExprType, IfExprBodyType, SwitchExprType, LoopExprType, LoopExprBodyType, OptionalSequenceExprType, CallExprType, AtomExprType, ParseErrorExprType, LiteralExprType, NamedExprType, BlockExprType, InitializeExprType, ImplicitStructExprType, AccessibleExprType, AccessExprType, IndexExprType, PostfixableExprType, PostfixUnaryExprType, PrefixableExprType, PrefixUnaryExprType, MulExprType, BinaryMulExprType, AddExprType, BinaryAddExprType, CmpExprType, BinaryCmpExprType, AndExprType, BinaryAndExprType, OrExprType, BinaryOrExprType, AnonymousFuncExprType:
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
	case CommentGroupsToken, IntegerLiteralToken, FloatLiteralToken, RuneLiteralToken, StringLiteralToken, IdentifierToken, UnderscoreToken, TrueToken, FalseToken, IfToken, ElseToken, SwitchToken, CaseToken, DefaultToken, ForToken, DoToken, InToken, ReturnToken, BreakToken, ContinueToken, FallthroughToken, PackageToken, ImportToken, UnsafeToken, TypeToken, ImplementsToken, StructToken, EnumToken, TraitToken, FuncToken, AsyncToken, DeferToken, VarToken, LetToken, NotToken, AndToken, OrToken, LabelDeclToken, JumpLabelToken, LbraceToken, RbraceToken, LparenToken, RparenToken, LbracketToken, RbracketToken, DotToken, CommaToken, QuestionToken, SemicolonToken, ColonToken, ExclaimToken, DollarLbracketToken, EllipsisToken, TildeTildeToken, AssignToken, AddAssignToken, SubAssignToken, MulAssignToken, DivAssignToken, ModAssignToken, AddOneAssignToken, SubOneAssignToken, BitNegAssignToken, BitAndAssignToken, BitOrAssignToken, BitXorAssignToken, BitLshiftAssignToken, BitRshiftAssignToken, AddToken, SubToken, MulToken, DivToken, ModToken, BitNegToken, BitAndToken, BitXorToken, BitOrToken, BitLshiftToken, BitRshiftToken, EqualToken, NotEqualToken, LessToken, LessOrEqualToken, GreaterToken, GreaterOrEqualToken, CallbackOpType, JumpTypeType, UnaryOpAssignType, BinaryOpAssignType, VarOrLetType, OptionalLabelDeclType, PostfixUnaryOpType, PrefixUnaryOpType, MulOpType, AddOpType, CmpOpType, PrefixUnaryTypeOpType, BinaryTypeOpType:
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
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExprType
		//line grammar.lr:336:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceSwitchExprToExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExprType
		//line grammar.lr:337:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceLoopExprToExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ExprType
		//line grammar.lr:338:4
		symbol.Expression = args[0].Expression
		err = nil
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
	case _ReduceToIfExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = IfExprType
		symbol.Expression, err = reducer.ToIfExpr(args[0].Value, args[1].Expression)
	case _ReduceNoElseToIfExprBody:
		args := stack[len(stack)-3:]
		stack = stack[:len(stack)-3]
		symbol.SymbolId_ = IfExprBodyType
		symbol.Expression, err = reducer.NoElseToIfExprBody(args[0].Value, args[1].Generic_, args[2].Generic_)
	case _ReduceIfElseToIfExprBody:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = IfExprBodyType
		symbol.Expression, err = reducer.IfElseToIfExprBody(args[0].Value, args[1].Generic_, args[2].Generic_, args[3].Value, args[4].Generic_)
	case _ReduceMultiIfElseToIfExprBody:
		args := stack[len(stack)-5:]
		stack = stack[:len(stack)-5]
		symbol.SymbolId_ = IfExprBodyType
		symbol.Expression, err = reducer.MultiIfElseToIfExprBody(args[0].Value, args[1].Generic_, args[2].Generic_, args[3].Value, args[4].Expression)
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
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = SwitchExprType
		symbol.Expression, err = reducer.ToSwitchExpr(args[0].Value, args[1].Value, args[2].Expression, args[3].Generic_)
	case _ReduceToLoopExpr:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LoopExprType
		symbol.Expression, err = reducer.ToLoopExpr(args[0].Value, args[1].Expression)
	case _ReduceInfiniteToLoopExprBody:
		args := stack[len(stack)-2:]
		stack = stack[:len(stack)-2]
		symbol.SymbolId_ = LoopExprBodyType
		symbol.Expression, err = reducer.InfiniteToLoopExprBody(args[0].Value, args[1].Generic_)
	case _ReduceDoWhileToLoopExprBody:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = LoopExprBodyType
		symbol.Expression, err = reducer.DoWhileToLoopExprBody(args[0].Value, args[1].Generic_, args[2].Value, args[3].Expression)
	case _ReduceWhileToLoopExprBody:
		args := stack[len(stack)-4:]
		stack = stack[:len(stack)-4]
		symbol.SymbolId_ = LoopExprBodyType
		symbol.Expression, err = reducer.WhileToLoopExprBody(args[0].Value, args[1].Expression, args[2].Value, args[3].Generic_)
	case _ReduceIteratorToLoopExprBody:
		args := stack[len(stack)-6:]
		stack = stack[:len(stack)-6]
		symbol.SymbolId_ = LoopExprBodyType
		symbol.Expression, err = reducer.IteratorToLoopExprBody(args[0].Value, args[1].Generic_, args[2].Value, args[3].Expression, args[4].Value, args[5].Generic_)
	case _ReduceForToLoopExprBody:
		args := stack[len(stack)-8:]
		stack = stack[:len(stack)-8]
		symbol.SymbolId_ = LoopExprBodyType
		symbol.Expression, err = reducer.ForToLoopExprBody(args[0].Value, args[1].Generic_, args[2].Value, args[3].Expression, args[4].Value, args[5].Expression, args[6].Value, args[7].Generic_)
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
		//line grammar.lr:441:4
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
		//line grammar.lr:475:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceLiteralExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:476:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceNamedExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:477:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBlockExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:478:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceAnonymousFuncExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:479:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceInitializeExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:480:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceImplicitStructExprToAtomExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomExprType
		//line grammar.lr:481:4
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
		//line grammar.lr:504:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceAccessExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:505:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceCallExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:506:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceIndexExprToAccessibleExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AccessibleExprType
		//line grammar.lr:507:4
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
		//line grammar.lr:516:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReducePostfixUnaryExprToPostfixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixableExprType
		//line grammar.lr:517:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceQuestionToPostfixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixUnaryOpType
		//line grammar.lr:520:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceExclaimToPostfixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PostfixUnaryOpType
		//line grammar.lr:521:4
		symbol.Value = args[0].Value
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
		//line grammar.lr:526:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReducePrefixUnaryExprToPrefixableExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixableExprType
		//line grammar.lr:527:4
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
		//line grammar.lr:532:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:533:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:534:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:537:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToPrefixUnaryOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryOpType
		//line grammar.lr:540:4
		symbol.Value = args[0].Value
		err = nil
	case _ReducePrefixableExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		//line grammar.lr:543:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryMulExprToMulExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulExprType
		//line grammar.lr:544:4
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
		//line grammar.lr:549:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceDivToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:550:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceModToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:551:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:552:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitLshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:553:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitRshiftToMulOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = MulOpType
		//line grammar.lr:554:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceMulExprToAddExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddExprType
		//line grammar.lr:557:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryAddExprToAddExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddExprType
		//line grammar.lr:558:4
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
		//line grammar.lr:563:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:564:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitOrToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:565:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitXorToAddOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AddOpType
		//line grammar.lr:566:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddExprToCmpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpExprType
		//line grammar.lr:569:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryCmpExprToCmpExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpExprType
		//line grammar.lr:570:4
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
		//line grammar.lr:575:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceNotEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:576:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLessToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:577:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceLessOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:578:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceGreaterToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:579:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceGreaterOrEqualToCmpOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = CmpOpType
		//line grammar.lr:580:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceCmpExprToAndExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AndExprType
		//line grammar.lr:583:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryAndExprToAndExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AndExprType
		//line grammar.lr:584:4
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
		//line grammar.lr:589:4
		symbol.Expression = args[0].Expression
		err = nil
	case _ReduceBinaryOrExprToOrExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = OrExprType
		//line grammar.lr:590:4
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
		//line grammar.lr:601:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceSliceTypeExprToInitializableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeExprType
		//line grammar.lr:602:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceArrayTypeExprToInitializableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeExprType
		//line grammar.lr:603:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceMapTypeExprToInitializableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = InitializableTypeExprType
		//line grammar.lr:604:4
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
		//line grammar.lr:616:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceNamedTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:617:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceInferredTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:618:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceImplicitStructTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:619:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceExplicitEnumTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:620:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceImplicitEnumTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:621:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceTraitTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:622:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceFuncTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:623:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceParseErrorTypeExprToAtomTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = AtomTypeExprType
		//line grammar.lr:624:4
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
		//line grammar.lr:645:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReducePrefixUnaryTypeExprToReturnableTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = ReturnableTypeExprType
		//line grammar.lr:646:4
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
		//line grammar.lr:652:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceExclaimToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:653:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitAndToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:654:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceBitNegToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:655:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceTildeTildeToPrefixUnaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = PrefixUnaryTypeOpType
		//line grammar.lr:656:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceReturnableTypeExprToTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TypeExprType
		//line grammar.lr:661:4
		symbol.TypeExpression = args[0].TypeExpression
		err = nil
	case _ReduceBinaryTypeExprToTypeExpr:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = TypeExprType
		//line grammar.lr:662:4
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
		//line grammar.lr:668:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceAddToBinaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryTypeOpType
		//line grammar.lr:669:4
		symbol.Value = args[0].Value
		err = nil
	case _ReduceSubToBinaryTypeOp:
		args := stack[len(stack)-1:]
		stack = stack[:len(stack)-1]
		symbol.SymbolId_ = BinaryTypeOpType
		//line grammar.lr:670:4
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
		//line grammar.lr:695:4
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
		//line grammar.lr:708:4
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
		//line grammar.lr:727:4
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
		//line grammar.lr:739:4
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
		//line grammar.lr:752:4
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
		//line grammar.lr:776:4
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
		//line grammar.lr:790:4
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
		//line grammar.lr:817:4
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
		//line grammar.lr:827:4
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
		//line grammar.lr:836:4
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
		//line grammar.lr:845:4
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
	_GotoState479Action                                                 = &_Action{_ShiftAction, _State479, 0}
	_GotoState480Action                                                 = &_Action{_ShiftAction, _State480, 0}
	_GotoState481Action                                                 = &_Action{_ShiftAction, _State481, 0}
	_GotoState482Action                                                 = &_Action{_ShiftAction, _State482, 0}
	_GotoState483Action                                                 = &_Action{_ShiftAction, _State483, 0}
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
	_ReduceToIfExprAction                                               = &_Action{_ReduceAction, 0, _ReduceToIfExpr}
	_ReduceNoElseToIfExprBodyAction                                     = &_Action{_ReduceAction, 0, _ReduceNoElseToIfExprBody}
	_ReduceIfElseToIfExprBodyAction                                     = &_Action{_ReduceAction, 0, _ReduceIfElseToIfExprBody}
	_ReduceMultiIfElseToIfExprBodyAction                                = &_Action{_ReduceAction, 0, _ReduceMultiIfElseToIfExprBody}
	_ReduceSequenceExprToConditionAction                                = &_Action{_ReduceAction, 0, _ReduceSequenceExprToCondition}
	_ReduceCaseToConditionAction                                        = &_Action{_ReduceAction, 0, _ReduceCaseToCondition}
	_ReduceToSwitchExprAction                                           = &_Action{_ReduceAction, 0, _ReduceToSwitchExpr}
	_ReduceToLoopExprAction                                             = &_Action{_ReduceAction, 0, _ReduceToLoopExpr}
	_ReduceInfiniteToLoopExprBodyAction                                 = &_Action{_ReduceAction, 0, _ReduceInfiniteToLoopExprBody}
	_ReduceDoWhileToLoopExprBodyAction                                  = &_Action{_ReduceAction, 0, _ReduceDoWhileToLoopExprBody}
	_ReduceWhileToLoopExprBodyAction                                    = &_Action{_ReduceAction, 0, _ReduceWhileToLoopExprBody}
	_ReduceIteratorToLoopExprBodyAction                                 = &_Action{_ReduceAction, 0, _ReduceIteratorToLoopExprBody}
	_ReduceForToLoopExprBodyAction                                      = &_Action{_ReduceAction, 0, _ReduceForToLoopExprBody}
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
	_ReduceQuestionToPostfixUnaryOpAction                               = &_Action{_ReduceAction, 0, _ReduceQuestionToPostfixUnaryOp}
	_ReduceExclaimToPostfixUnaryOpAction                                = &_Action{_ReduceAction, 0, _ReduceExclaimToPostfixUnaryOp}
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
	{_State2, SimpleStatementType}:                 _GotoState102Action,
	{_State2, StatementType}:                       _GotoState6Action,
	{_State2, ExprOrImproperStructStatementType}:   _GotoState77Action,
	{_State2, ExprsType}:                           _GotoState78Action,
	{_State2, CallbackOpType}:                      _GotoState72Action,
	{_State2, CallbackOpStatementType}:             _GotoState73Action,
	{_State2, UnsafeStatementType}:                 _GotoState106Action,
	{_State2, JumpStatementType}:                   _GotoState86Action,
	{_State2, JumpTypeType}:                        _GotoState87Action,
	{_State2, FallthroughStatementType}:            _GotoState79Action,
	{_State2, AssignStatementType}:                 _GotoState62Action,
	{_State2, UnaryOpAssignStatementType}:          _GotoState105Action,
	{_State2, BinaryOpAssignStatementType}:         _GotoState68Action,
	{_State2, ImportStatementType}:                 _GotoState82Action,
	{_State2, VarDeclPatternType}:                  _GotoState107Action,
	{_State2, VarOrLetType}:                        _GotoState24Action,
	{_State2, AssignPatternType}:                   _GotoState61Action,
	{_State2, ExprType}:                            _GotoState76Action,
	{_State2, OptionalLabelDeclType}:               _GotoState93Action,
	{_State2, SequenceExprType}:                    _GotoState101Action,
	{_State2, IfExprType}:                          _GotoState80Action,
	{_State2, SwitchExprType}:                      _GotoState104Action,
	{_State2, LoopExprType}:                        _GotoState89Action,
	{_State2, CallExprType}:                        _GotoState71Action,
	{_State2, AtomExprType}:                        _GotoState63Action,
	{_State2, ParseErrorExprType}:                  _GotoState95Action,
	{_State2, LiteralExprType}:                     _GotoState88Action,
	{_State2, NamedExprType}:                       _GotoState92Action,
	{_State2, BlockExprType}:                       _GotoState70Action,
	{_State2, InitializeExprType}:                  _GotoState85Action,
	{_State2, ImplicitStructExprType}:              _GotoState81Action,
	{_State2, AccessibleExprType}:                  _GotoState56Action,
	{_State2, AccessExprType}:                      _GotoState55Action,
	{_State2, IndexExprType}:                       _GotoState83Action,
	{_State2, PostfixableExprType}:                 _GotoState97Action,
	{_State2, PostfixUnaryExprType}:                _GotoState96Action,
	{_State2, PrefixableExprType}:                  _GotoState100Action,
	{_State2, PrefixUnaryExprType}:                 _GotoState98Action,
	{_State2, PrefixUnaryOpType}:                   _GotoState99Action,
	{_State2, MulExprType}:                         _GotoState91Action,
	{_State2, BinaryMulExprType}:                   _GotoState67Action,
	{_State2, AddExprType}:                         _GotoState57Action,
	{_State2, BinaryAddExprType}:                   _GotoState64Action,
	{_State2, CmpExprType}:                         _GotoState74Action,
	{_State2, BinaryCmpExprType}:                   _GotoState66Action,
	{_State2, AndExprType}:                         _GotoState58Action,
	{_State2, BinaryAndExprType}:                   _GotoState65Action,
	{_State2, OrExprType}:                          _GotoState94Action,
	{_State2, BinaryOrExprType}:                    _GotoState69Action,
	{_State2, InitializableTypeExprType}:           _GotoState84Action,
	{_State2, SliceTypeExprType}:                   _GotoState103Action,
	{_State2, ArrayTypeExprType}:                   _GotoState60Action,
	{_State2, MapTypeExprType}:                     _GotoState90Action,
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
	{_State3, VarDeclPatternType}:                  _GotoState107Action,
	{_State3, VarOrLetType}:                        _GotoState24Action,
	{_State3, ExprType}:                            _GotoState7Action,
	{_State3, OptionalLabelDeclType}:               _GotoState93Action,
	{_State3, SequenceExprType}:                    _GotoState109Action,
	{_State3, IfExprType}:                          _GotoState80Action,
	{_State3, SwitchExprType}:                      _GotoState104Action,
	{_State3, LoopExprType}:                        _GotoState89Action,
	{_State3, CallExprType}:                        _GotoState71Action,
	{_State3, AtomExprType}:                        _GotoState63Action,
	{_State3, ParseErrorExprType}:                  _GotoState95Action,
	{_State3, LiteralExprType}:                     _GotoState88Action,
	{_State3, NamedExprType}:                       _GotoState92Action,
	{_State3, BlockExprType}:                       _GotoState70Action,
	{_State3, InitializeExprType}:                  _GotoState85Action,
	{_State3, ImplicitStructExprType}:              _GotoState81Action,
	{_State3, AccessibleExprType}:                  _GotoState108Action,
	{_State3, AccessExprType}:                      _GotoState55Action,
	{_State3, IndexExprType}:                       _GotoState83Action,
	{_State3, PostfixableExprType}:                 _GotoState97Action,
	{_State3, PostfixUnaryExprType}:                _GotoState96Action,
	{_State3, PrefixableExprType}:                  _GotoState100Action,
	{_State3, PrefixUnaryExprType}:                 _GotoState98Action,
	{_State3, PrefixUnaryOpType}:                   _GotoState99Action,
	{_State3, MulExprType}:                         _GotoState91Action,
	{_State3, BinaryMulExprType}:                   _GotoState67Action,
	{_State3, AddExprType}:                         _GotoState57Action,
	{_State3, BinaryAddExprType}:                   _GotoState64Action,
	{_State3, CmpExprType}:                         _GotoState74Action,
	{_State3, BinaryCmpExprType}:                   _GotoState66Action,
	{_State3, AndExprType}:                         _GotoState58Action,
	{_State3, BinaryAndExprType}:                   _GotoState65Action,
	{_State3, OrExprType}:                          _GotoState94Action,
	{_State3, BinaryOrExprType}:                    _GotoState69Action,
	{_State3, InitializableTypeExprType}:           _GotoState84Action,
	{_State3, SliceTypeExprType}:                   _GotoState103Action,
	{_State3, ArrayTypeExprType}:                   _GotoState60Action,
	{_State3, MapTypeExprType}:                     _GotoState90Action,
	{_State3, ExplicitStructTypeExprType}:          _GotoState75Action,
	{_State3, AnonymousFuncExprType}:               _GotoState59Action,
	{_State4, IdentifierToken}:                     _GotoState116Action,
	{_State4, UnderscoreToken}:                     _GotoState122Action,
	{_State4, StructToken}:                         _GotoState50Action,
	{_State4, EnumToken}:                           _GotoState113Action,
	{_State4, TraitToken}:                          _GotoState121Action,
	{_State4, FuncToken}:                           _GotoState115Action,
	{_State4, LparenToken}:                         _GotoState117Action,
	{_State4, LbracketToken}:                       _GotoState42Action,
	{_State4, DotToken}:                            _GotoState112Action,
	{_State4, QuestionToken}:                       _GotoState119Action,
	{_State4, ExclaimToken}:                        _GotoState114Action,
	{_State4, TildeTildeToken}:                     _GotoState120Action,
	{_State4, BitNegToken}:                         _GotoState111Action,
	{_State4, BitAndToken}:                         _GotoState110Action,
	{_State4, ParseErrorToken}:                     _GotoState118Action,
	{_State4, InitializableTypeExprType}:           _GotoState130Action,
	{_State4, SliceTypeExprType}:                   _GotoState103Action,
	{_State4, ArrayTypeExprType}:                   _GotoState60Action,
	{_State4, MapTypeExprType}:                     _GotoState90Action,
	{_State4, AtomTypeExprType}:                    _GotoState123Action,
	{_State4, NamedTypeExprType}:                   _GotoState131Action,
	{_State4, InferredTypeExprType}:                _GotoState129Action,
	{_State4, ParseErrorTypeExprType}:              _GotoState132Action,
	{_State4, ReturnableTypeExprType}:              _GotoState135Action,
	{_State4, PrefixUnaryTypeExprType}:             _GotoState133Action,
	{_State4, PrefixUnaryTypeOpType}:               _GotoState134Action,
	{_State4, TypeExprType}:                        _GotoState8Action,
	{_State4, BinaryTypeExprType}:                  _GotoState124Action,
	{_State4, ImplicitStructTypeExprType}:          _GotoState128Action,
	{_State4, ExplicitStructTypeExprType}:          _GotoState75Action,
	{_State4, TraitTypeExprType}:                   _GotoState136Action,
	{_State4, ImplicitEnumTypeExprType}:            _GotoState127Action,
	{_State4, ExplicitEnumTypeExprType}:            _GotoState125Action,
	{_State4, FuncTypeExprType}:                    _GotoState126Action,
	{_State8, AddToken}:                            _GotoState253Action,
	{_State8, SubToken}:                            _GotoState255Action,
	{_State8, MulToken}:                            _GotoState254Action,
	{_State8, BinaryTypeOpType}:                    _GotoState256Action,
	{_State10, IdentifierToken}:                    _GotoState137Action,
	{_State10, LparenToken}:                        _GotoState138Action,
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
	{_State11, ProperStatementsType}:               _GotoState139Action,
	{_State11, StatementsType}:                     _GotoState141Action,
	{_State11, SimpleStatementType}:                _GotoState102Action,
	{_State11, StatementType}:                      _GotoState140Action,
	{_State11, ExprOrImproperStructStatementType}:  _GotoState77Action,
	{_State11, ExprsType}:                          _GotoState78Action,
	{_State11, CallbackOpType}:                     _GotoState72Action,
	{_State11, CallbackOpStatementType}:            _GotoState73Action,
	{_State11, UnsafeStatementType}:                _GotoState106Action,
	{_State11, JumpStatementType}:                  _GotoState86Action,
	{_State11, JumpTypeType}:                       _GotoState87Action,
	{_State11, FallthroughStatementType}:           _GotoState79Action,
	{_State11, AssignStatementType}:                _GotoState62Action,
	{_State11, UnaryOpAssignStatementType}:         _GotoState105Action,
	{_State11, BinaryOpAssignStatementType}:        _GotoState68Action,
	{_State11, ImportStatementType}:                _GotoState82Action,
	{_State11, VarDeclPatternType}:                 _GotoState107Action,
	{_State11, VarOrLetType}:                       _GotoState24Action,
	{_State11, AssignPatternType}:                  _GotoState61Action,
	{_State11, ExprType}:                           _GotoState76Action,
	{_State11, OptionalLabelDeclType}:              _GotoState93Action,
	{_State11, SequenceExprType}:                   _GotoState101Action,
	{_State11, IfExprType}:                         _GotoState80Action,
	{_State11, SwitchExprType}:                     _GotoState104Action,
	{_State11, LoopExprType}:                       _GotoState89Action,
	{_State11, CallExprType}:                       _GotoState71Action,
	{_State11, AtomExprType}:                       _GotoState63Action,
	{_State11, ParseErrorExprType}:                 _GotoState95Action,
	{_State11, LiteralExprType}:                    _GotoState88Action,
	{_State11, NamedExprType}:                      _GotoState92Action,
	{_State11, BlockExprType}:                      _GotoState70Action,
	{_State11, InitializeExprType}:                 _GotoState85Action,
	{_State11, ImplicitStructExprType}:             _GotoState81Action,
	{_State11, AccessibleExprType}:                 _GotoState56Action,
	{_State11, AccessExprType}:                     _GotoState55Action,
	{_State11, IndexExprType}:                      _GotoState83Action,
	{_State11, PostfixableExprType}:                _GotoState97Action,
	{_State11, PostfixUnaryExprType}:               _GotoState96Action,
	{_State11, PrefixableExprType}:                 _GotoState100Action,
	{_State11, PrefixUnaryExprType}:                _GotoState98Action,
	{_State11, PrefixUnaryOpType}:                  _GotoState99Action,
	{_State11, MulExprType}:                        _GotoState91Action,
	{_State11, BinaryMulExprType}:                  _GotoState67Action,
	{_State11, AddExprType}:                        _GotoState57Action,
	{_State11, BinaryAddExprType}:                  _GotoState64Action,
	{_State11, CmpExprType}:                        _GotoState74Action,
	{_State11, BinaryCmpExprType}:                  _GotoState66Action,
	{_State11, AndExprType}:                        _GotoState58Action,
	{_State11, BinaryAndExprType}:                  _GotoState65Action,
	{_State11, OrExprType}:                         _GotoState94Action,
	{_State11, BinaryOrExprType}:                   _GotoState69Action,
	{_State11, InitializableTypeExprType}:          _GotoState84Action,
	{_State11, SliceTypeExprType}:                  _GotoState103Action,
	{_State11, ArrayTypeExprType}:                  _GotoState60Action,
	{_State11, MapTypeExprType}:                    _GotoState90Action,
	{_State11, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State11, AnonymousFuncExprType}:              _GotoState59Action,
	{_State13, LbraceToken}:                        _GotoState11Action,
	{_State13, StatementBlockType}:                 _GotoState142Action,
	{_State14, IdentifierToken}:                    _GotoState143Action,
	{_State20, NewlinesToken}:                      _GotoState144Action,
	{_State23, AssignToken}:                        _GotoState145Action,
	{_State24, IdentifierToken}:                    _GotoState146Action,
	{_State24, UnderscoreToken}:                    _GotoState148Action,
	{_State24, LparenToken}:                        _GotoState147Action,
	{_State24, VarPatternType}:                     _GotoState150Action,
	{_State24, TuplePatternType}:                   _GotoState149Action,
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
	{_State29, VarToken}:                           _GotoState152Action,
	{_State29, LetToken}:                           _GotoState12Action,
	{_State29, NotToken}:                           _GotoState45Action,
	{_State29, LabelDeclToken}:                     _GotoState41Action,
	{_State29, LparenToken}:                        _GotoState43Action,
	{_State29, LbracketToken}:                      _GotoState42Action,
	{_State29, DotToken}:                           _GotoState151Action,
	{_State29, SubToken}:                           _GotoState51Action,
	{_State29, MulToken}:                           _GotoState44Action,
	{_State29, BitNegToken}:                        _GotoState27Action,
	{_State29, BitAndToken}:                        _GotoState26Action,
	{_State29, GreaterToken}:                       _GotoState37Action,
	{_State29, ParseErrorToken}:                    _GotoState46Action,
	{_State29, CasePatternsType}:                   _GotoState155Action,
	{_State29, VarDeclPatternType}:                 _GotoState107Action,
	{_State29, VarOrLetType}:                       _GotoState24Action,
	{_State29, AssignPatternType}:                  _GotoState153Action,
	{_State29, CasePatternType}:                    _GotoState154Action,
	{_State29, OptionalLabelDeclType}:              _GotoState156Action,
	{_State29, SequenceExprType}:                   _GotoState157Action,
	{_State29, CallExprType}:                       _GotoState71Action,
	{_State29, AtomExprType}:                       _GotoState63Action,
	{_State29, ParseErrorExprType}:                 _GotoState95Action,
	{_State29, LiteralExprType}:                    _GotoState88Action,
	{_State29, NamedExprType}:                      _GotoState92Action,
	{_State29, BlockExprType}:                      _GotoState70Action,
	{_State29, InitializeExprType}:                 _GotoState85Action,
	{_State29, ImplicitStructExprType}:             _GotoState81Action,
	{_State29, AccessibleExprType}:                 _GotoState108Action,
	{_State29, AccessExprType}:                     _GotoState55Action,
	{_State29, IndexExprType}:                      _GotoState83Action,
	{_State29, PostfixableExprType}:                _GotoState97Action,
	{_State29, PostfixUnaryExprType}:               _GotoState96Action,
	{_State29, PrefixableExprType}:                 _GotoState100Action,
	{_State29, PrefixUnaryExprType}:                _GotoState98Action,
	{_State29, PrefixUnaryOpType}:                  _GotoState99Action,
	{_State29, MulExprType}:                        _GotoState91Action,
	{_State29, BinaryMulExprType}:                  _GotoState67Action,
	{_State29, AddExprType}:                        _GotoState57Action,
	{_State29, BinaryAddExprType}:                  _GotoState64Action,
	{_State29, CmpExprType}:                        _GotoState74Action,
	{_State29, BinaryCmpExprType}:                  _GotoState66Action,
	{_State29, AndExprType}:                        _GotoState58Action,
	{_State29, BinaryAndExprType}:                  _GotoState65Action,
	{_State29, OrExprType}:                         _GotoState94Action,
	{_State29, BinaryOrExprType}:                   _GotoState69Action,
	{_State29, InitializableTypeExprType}:          _GotoState84Action,
	{_State29, SliceTypeExprType}:                  _GotoState103Action,
	{_State29, ArrayTypeExprType}:                  _GotoState60Action,
	{_State29, MapTypeExprType}:                    _GotoState90Action,
	{_State29, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State29, AnonymousFuncExprType}:              _GotoState59Action,
	{_State31, ColonToken}:                         _GotoState158Action,
	{_State36, LparenToken}:                        _GotoState159Action,
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
	{_State37, VarDeclPatternType}:                 _GotoState107Action,
	{_State37, VarOrLetType}:                       _GotoState24Action,
	{_State37, OptionalLabelDeclType}:              _GotoState156Action,
	{_State37, SequenceExprType}:                   _GotoState160Action,
	{_State37, CallExprType}:                       _GotoState71Action,
	{_State37, AtomExprType}:                       _GotoState63Action,
	{_State37, ParseErrorExprType}:                 _GotoState95Action,
	{_State37, LiteralExprType}:                    _GotoState88Action,
	{_State37, NamedExprType}:                      _GotoState92Action,
	{_State37, BlockExprType}:                      _GotoState70Action,
	{_State37, InitializeExprType}:                 _GotoState85Action,
	{_State37, ImplicitStructExprType}:             _GotoState81Action,
	{_State37, AccessibleExprType}:                 _GotoState108Action,
	{_State37, AccessExprType}:                     _GotoState55Action,
	{_State37, IndexExprType}:                      _GotoState83Action,
	{_State37, PostfixableExprType}:                _GotoState97Action,
	{_State37, PostfixUnaryExprType}:               _GotoState96Action,
	{_State37, PrefixableExprType}:                 _GotoState100Action,
	{_State37, PrefixUnaryExprType}:                _GotoState98Action,
	{_State37, PrefixUnaryOpType}:                  _GotoState99Action,
	{_State37, MulExprType}:                        _GotoState91Action,
	{_State37, BinaryMulExprType}:                  _GotoState67Action,
	{_State37, AddExprType}:                        _GotoState57Action,
	{_State37, BinaryAddExprType}:                  _GotoState64Action,
	{_State37, CmpExprType}:                        _GotoState74Action,
	{_State37, BinaryCmpExprType}:                  _GotoState66Action,
	{_State37, AndExprType}:                        _GotoState58Action,
	{_State37, BinaryAndExprType}:                  _GotoState65Action,
	{_State37, OrExprType}:                         _GotoState94Action,
	{_State37, BinaryOrExprType}:                   _GotoState69Action,
	{_State37, InitializableTypeExprType}:          _GotoState84Action,
	{_State37, SliceTypeExprType}:                  _GotoState103Action,
	{_State37, ArrayTypeExprType}:                  _GotoState60Action,
	{_State37, MapTypeExprType}:                    _GotoState90Action,
	{_State37, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State37, AnonymousFuncExprType}:              _GotoState59Action,
	{_State39, StringLiteralToken}:                 _GotoState164Action,
	{_State39, IdentifierToken}:                    _GotoState162Action,
	{_State39, UnderscoreToken}:                    _GotoState165Action,
	{_State39, LparenToken}:                        _GotoState163Action,
	{_State39, DotToken}:                           _GotoState161Action,
	{_State39, ImportClauseType}:                   _GotoState166Action,
	{_State42, IdentifierToken}:                    _GotoState116Action,
	{_State42, UnderscoreToken}:                    _GotoState122Action,
	{_State42, StructToken}:                        _GotoState50Action,
	{_State42, EnumToken}:                          _GotoState113Action,
	{_State42, TraitToken}:                         _GotoState121Action,
	{_State42, FuncToken}:                          _GotoState115Action,
	{_State42, LparenToken}:                        _GotoState117Action,
	{_State42, LbracketToken}:                      _GotoState42Action,
	{_State42, DotToken}:                           _GotoState112Action,
	{_State42, QuestionToken}:                      _GotoState119Action,
	{_State42, ExclaimToken}:                       _GotoState114Action,
	{_State42, TildeTildeToken}:                    _GotoState120Action,
	{_State42, BitNegToken}:                        _GotoState111Action,
	{_State42, BitAndToken}:                        _GotoState110Action,
	{_State42, ParseErrorToken}:                    _GotoState118Action,
	{_State42, InitializableTypeExprType}:          _GotoState130Action,
	{_State42, SliceTypeExprType}:                  _GotoState103Action,
	{_State42, ArrayTypeExprType}:                  _GotoState60Action,
	{_State42, MapTypeExprType}:                    _GotoState90Action,
	{_State42, AtomTypeExprType}:                   _GotoState123Action,
	{_State42, NamedTypeExprType}:                  _GotoState131Action,
	{_State42, InferredTypeExprType}:               _GotoState129Action,
	{_State42, ParseErrorTypeExprType}:             _GotoState132Action,
	{_State42, ReturnableTypeExprType}:             _GotoState135Action,
	{_State42, PrefixUnaryTypeExprType}:            _GotoState133Action,
	{_State42, PrefixUnaryTypeOpType}:              _GotoState134Action,
	{_State42, TypeExprType}:                       _GotoState167Action,
	{_State42, BinaryTypeExprType}:                 _GotoState124Action,
	{_State42, ImplicitStructTypeExprType}:         _GotoState128Action,
	{_State42, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State42, TraitTypeExprType}:                  _GotoState136Action,
	{_State42, ImplicitEnumTypeExprType}:           _GotoState127Action,
	{_State42, ExplicitEnumTypeExprType}:           _GotoState125Action,
	{_State42, FuncTypeExprType}:                   _GotoState126Action,
	{_State43, IntegerLiteralToken}:                _GotoState40Action,
	{_State43, FloatLiteralToken}:                  _GotoState35Action,
	{_State43, RuneLiteralToken}:                   _GotoState48Action,
	{_State43, StringLiteralToken}:                 _GotoState49Action,
	{_State43, IdentifierToken}:                    _GotoState170Action,
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
	{_State43, ColonToken}:                         _GotoState168Action,
	{_State43, EllipsisToken}:                      _GotoState169Action,
	{_State43, SubToken}:                           _GotoState51Action,
	{_State43, MulToken}:                           _GotoState44Action,
	{_State43, BitNegToken}:                        _GotoState27Action,
	{_State43, BitAndToken}:                        _GotoState26Action,
	{_State43, GreaterToken}:                       _GotoState37Action,
	{_State43, ParseErrorToken}:                    _GotoState46Action,
	{_State43, VarDeclPatternType}:                 _GotoState107Action,
	{_State43, VarOrLetType}:                       _GotoState24Action,
	{_State43, ExprType}:                           _GotoState174Action,
	{_State43, OptionalLabelDeclType}:              _GotoState93Action,
	{_State43, SequenceExprType}:                   _GotoState109Action,
	{_State43, IfExprType}:                         _GotoState80Action,
	{_State43, SwitchExprType}:                     _GotoState104Action,
	{_State43, LoopExprType}:                       _GotoState89Action,
	{_State43, CallExprType}:                       _GotoState71Action,
	{_State43, ProperArgumentsType}:                _GotoState175Action,
	{_State43, ArgumentsType}:                      _GotoState172Action,
	{_State43, ArgumentType}:                       _GotoState171Action,
	{_State43, ColonExprType}:                      _GotoState173Action,
	{_State43, AtomExprType}:                       _GotoState63Action,
	{_State43, ParseErrorExprType}:                 _GotoState95Action,
	{_State43, LiteralExprType}:                    _GotoState88Action,
	{_State43, NamedExprType}:                      _GotoState92Action,
	{_State43, BlockExprType}:                      _GotoState70Action,
	{_State43, InitializeExprType}:                 _GotoState85Action,
	{_State43, ImplicitStructExprType}:             _GotoState81Action,
	{_State43, AccessibleExprType}:                 _GotoState108Action,
	{_State43, AccessExprType}:                     _GotoState55Action,
	{_State43, IndexExprType}:                      _GotoState83Action,
	{_State43, PostfixableExprType}:                _GotoState97Action,
	{_State43, PostfixUnaryExprType}:               _GotoState96Action,
	{_State43, PrefixableExprType}:                 _GotoState100Action,
	{_State43, PrefixUnaryExprType}:                _GotoState98Action,
	{_State43, PrefixUnaryOpType}:                  _GotoState99Action,
	{_State43, MulExprType}:                        _GotoState91Action,
	{_State43, BinaryMulExprType}:                  _GotoState67Action,
	{_State43, AddExprType}:                        _GotoState57Action,
	{_State43, BinaryAddExprType}:                  _GotoState64Action,
	{_State43, CmpExprType}:                        _GotoState74Action,
	{_State43, BinaryCmpExprType}:                  _GotoState66Action,
	{_State43, AndExprType}:                        _GotoState58Action,
	{_State43, BinaryAndExprType}:                  _GotoState65Action,
	{_State43, OrExprType}:                         _GotoState94Action,
	{_State43, BinaryOrExprType}:                   _GotoState69Action,
	{_State43, InitializableTypeExprType}:          _GotoState84Action,
	{_State43, SliceTypeExprType}:                  _GotoState103Action,
	{_State43, ArrayTypeExprType}:                  _GotoState60Action,
	{_State43, MapTypeExprType}:                    _GotoState90Action,
	{_State43, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State43, AnonymousFuncExprType}:              _GotoState59Action,
	{_State50, LparenToken}:                        _GotoState176Action,
	{_State54, LessToken}:                          _GotoState177Action,
	{_State56, LbracketToken}:                      _GotoState190Action,
	{_State56, DotToken}:                           _GotoState188Action,
	{_State56, QuestionToken}:                      _GotoState193Action,
	{_State56, ExclaimToken}:                       _GotoState189Action,
	{_State56, DollarLbracketToken}:                _GotoState187Action,
	{_State56, AddAssignToken}:                     _GotoState178Action,
	{_State56, SubAssignToken}:                     _GotoState194Action,
	{_State56, MulAssignToken}:                     _GotoState192Action,
	{_State56, DivAssignToken}:                     _GotoState186Action,
	{_State56, ModAssignToken}:                     _GotoState191Action,
	{_State56, AddOneAssignToken}:                  _GotoState179Action,
	{_State56, SubOneAssignToken}:                  _GotoState195Action,
	{_State56, BitNegAssignToken}:                  _GotoState182Action,
	{_State56, BitAndAssignToken}:                  _GotoState180Action,
	{_State56, BitOrAssignToken}:                   _GotoState183Action,
	{_State56, BitXorAssignToken}:                  _GotoState185Action,
	{_State56, BitLshiftAssignToken}:               _GotoState181Action,
	{_State56, BitRshiftAssignToken}:               _GotoState184Action,
	{_State56, UnaryOpAssignType}:                  _GotoState199Action,
	{_State56, BinaryOpAssignType}:                 _GotoState196Action,
	{_State56, PostfixUnaryOpType}:                 _GotoState198Action,
	{_State56, GenericTypeArgumentsType}:           _GotoState197Action,
	{_State57, AddToken}:                           _GotoState200Action,
	{_State57, SubToken}:                           _GotoState203Action,
	{_State57, BitXorToken}:                        _GotoState202Action,
	{_State57, BitOrToken}:                         _GotoState201Action,
	{_State57, AddOpType}:                          _GotoState204Action,
	{_State58, AndToken}:                           _GotoState205Action,
	{_State61, AssignToken}:                        _GotoState206Action,
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
	{_State72, OptionalLabelDeclType}:              _GotoState156Action,
	{_State72, CallExprType}:                       _GotoState208Action,
	{_State72, AtomExprType}:                       _GotoState63Action,
	{_State72, ParseErrorExprType}:                 _GotoState95Action,
	{_State72, LiteralExprType}:                    _GotoState88Action,
	{_State72, NamedExprType}:                      _GotoState92Action,
	{_State72, BlockExprType}:                      _GotoState70Action,
	{_State72, InitializeExprType}:                 _GotoState85Action,
	{_State72, ImplicitStructExprType}:             _GotoState81Action,
	{_State72, AccessibleExprType}:                 _GotoState207Action,
	{_State72, AccessExprType}:                     _GotoState55Action,
	{_State72, IndexExprType}:                      _GotoState83Action,
	{_State72, InitializableTypeExprType}:          _GotoState84Action,
	{_State72, SliceTypeExprType}:                  _GotoState103Action,
	{_State72, ArrayTypeExprType}:                  _GotoState60Action,
	{_State72, MapTypeExprType}:                    _GotoState90Action,
	{_State72, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State72, AnonymousFuncExprType}:              _GotoState59Action,
	{_State74, EqualToken}:                         _GotoState209Action,
	{_State74, NotEqualToken}:                      _GotoState214Action,
	{_State74, LessToken}:                          _GotoState212Action,
	{_State74, LessOrEqualToken}:                   _GotoState213Action,
	{_State74, GreaterToken}:                       _GotoState210Action,
	{_State74, GreaterOrEqualToken}:                _GotoState211Action,
	{_State74, CmpOpType}:                          _GotoState215Action,
	{_State78, CommaToken}:                         _GotoState216Action,
	{_State84, LparenToken}:                        _GotoState217Action,
	{_State87, IntegerLiteralToken}:                _GotoState40Action,
	{_State87, FloatLiteralToken}:                  _GotoState35Action,
	{_State87, RuneLiteralToken}:                   _GotoState48Action,
	{_State87, StringLiteralToken}:                 _GotoState49Action,
	{_State87, IdentifierToken}:                    _GotoState38Action,
	{_State87, UnderscoreToken}:                    _GotoState53Action,
	{_State87, TrueToken}:                          _GotoState52Action,
	{_State87, FalseToken}:                         _GotoState34Action,
	{_State87, StructToken}:                        _GotoState50Action,
	{_State87, FuncToken}:                          _GotoState36Action,
	{_State87, VarToken}:                           _GotoState15Action,
	{_State87, LetToken}:                           _GotoState12Action,
	{_State87, NotToken}:                           _GotoState45Action,
	{_State87, LabelDeclToken}:                     _GotoState41Action,
	{_State87, JumpLabelToken}:                     _GotoState218Action,
	{_State87, LparenToken}:                        _GotoState43Action,
	{_State87, LbracketToken}:                      _GotoState42Action,
	{_State87, SubToken}:                           _GotoState51Action,
	{_State87, MulToken}:                           _GotoState44Action,
	{_State87, BitNegToken}:                        _GotoState27Action,
	{_State87, BitAndToken}:                        _GotoState26Action,
	{_State87, GreaterToken}:                       _GotoState37Action,
	{_State87, ParseErrorToken}:                    _GotoState46Action,
	{_State87, ExprsType}:                          _GotoState219Action,
	{_State87, VarDeclPatternType}:                 _GotoState107Action,
	{_State87, VarOrLetType}:                       _GotoState24Action,
	{_State87, ExprType}:                           _GotoState76Action,
	{_State87, OptionalLabelDeclType}:              _GotoState93Action,
	{_State87, SequenceExprType}:                   _GotoState109Action,
	{_State87, IfExprType}:                         _GotoState80Action,
	{_State87, SwitchExprType}:                     _GotoState104Action,
	{_State87, LoopExprType}:                       _GotoState89Action,
	{_State87, CallExprType}:                       _GotoState71Action,
	{_State87, AtomExprType}:                       _GotoState63Action,
	{_State87, ParseErrorExprType}:                 _GotoState95Action,
	{_State87, LiteralExprType}:                    _GotoState88Action,
	{_State87, NamedExprType}:                      _GotoState92Action,
	{_State87, BlockExprType}:                      _GotoState70Action,
	{_State87, InitializeExprType}:                 _GotoState85Action,
	{_State87, ImplicitStructExprType}:             _GotoState81Action,
	{_State87, AccessibleExprType}:                 _GotoState108Action,
	{_State87, AccessExprType}:                     _GotoState55Action,
	{_State87, IndexExprType}:                      _GotoState83Action,
	{_State87, PostfixableExprType}:                _GotoState97Action,
	{_State87, PostfixUnaryExprType}:               _GotoState96Action,
	{_State87, PrefixableExprType}:                 _GotoState100Action,
	{_State87, PrefixUnaryExprType}:                _GotoState98Action,
	{_State87, PrefixUnaryOpType}:                  _GotoState99Action,
	{_State87, MulExprType}:                        _GotoState91Action,
	{_State87, BinaryMulExprType}:                  _GotoState67Action,
	{_State87, AddExprType}:                        _GotoState57Action,
	{_State87, BinaryAddExprType}:                  _GotoState64Action,
	{_State87, CmpExprType}:                        _GotoState74Action,
	{_State87, BinaryCmpExprType}:                  _GotoState66Action,
	{_State87, AndExprType}:                        _GotoState58Action,
	{_State87, BinaryAndExprType}:                  _GotoState65Action,
	{_State87, OrExprType}:                         _GotoState94Action,
	{_State87, BinaryOrExprType}:                   _GotoState69Action,
	{_State87, InitializableTypeExprType}:          _GotoState84Action,
	{_State87, SliceTypeExprType}:                  _GotoState103Action,
	{_State87, ArrayTypeExprType}:                  _GotoState60Action,
	{_State87, MapTypeExprType}:                    _GotoState90Action,
	{_State87, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State87, AnonymousFuncExprType}:              _GotoState59Action,
	{_State91, MulToken}:                           _GotoState225Action,
	{_State91, DivToken}:                           _GotoState223Action,
	{_State91, ModToken}:                           _GotoState224Action,
	{_State91, BitAndToken}:                        _GotoState220Action,
	{_State91, BitLshiftToken}:                     _GotoState221Action,
	{_State91, BitRshiftToken}:                     _GotoState222Action,
	{_State91, MulOpType}:                          _GotoState226Action,
	{_State93, IfToken}:                            _GotoState229Action,
	{_State93, SwitchToken}:                        _GotoState230Action,
	{_State93, ForToken}:                           _GotoState228Action,
	{_State93, DoToken}:                            _GotoState227Action,
	{_State93, LbraceToken}:                        _GotoState11Action,
	{_State93, StatementBlockType}:                 _GotoState233Action,
	{_State93, IfExprBodyType}:                     _GotoState231Action,
	{_State93, LoopExprBodyType}:                   _GotoState232Action,
	{_State94, OrToken}:                            _GotoState234Action,
	{_State99, IntegerLiteralToken}:                _GotoState40Action,
	{_State99, FloatLiteralToken}:                  _GotoState35Action,
	{_State99, RuneLiteralToken}:                   _GotoState48Action,
	{_State99, StringLiteralToken}:                 _GotoState49Action,
	{_State99, IdentifierToken}:                    _GotoState38Action,
	{_State99, UnderscoreToken}:                    _GotoState53Action,
	{_State99, TrueToken}:                          _GotoState52Action,
	{_State99, FalseToken}:                         _GotoState34Action,
	{_State99, StructToken}:                        _GotoState50Action,
	{_State99, FuncToken}:                          _GotoState36Action,
	{_State99, NotToken}:                           _GotoState45Action,
	{_State99, LabelDeclToken}:                     _GotoState41Action,
	{_State99, LparenToken}:                        _GotoState43Action,
	{_State99, LbracketToken}:                      _GotoState42Action,
	{_State99, SubToken}:                           _GotoState51Action,
	{_State99, MulToken}:                           _GotoState44Action,
	{_State99, BitNegToken}:                        _GotoState27Action,
	{_State99, BitAndToken}:                        _GotoState26Action,
	{_State99, ParseErrorToken}:                    _GotoState46Action,
	{_State99, OptionalLabelDeclType}:              _GotoState156Action,
	{_State99, CallExprType}:                       _GotoState71Action,
	{_State99, AtomExprType}:                       _GotoState63Action,
	{_State99, ParseErrorExprType}:                 _GotoState95Action,
	{_State99, LiteralExprType}:                    _GotoState88Action,
	{_State99, NamedExprType}:                      _GotoState92Action,
	{_State99, BlockExprType}:                      _GotoState70Action,
	{_State99, InitializeExprType}:                 _GotoState85Action,
	{_State99, ImplicitStructExprType}:             _GotoState81Action,
	{_State99, AccessibleExprType}:                 _GotoState108Action,
	{_State99, AccessExprType}:                     _GotoState55Action,
	{_State99, IndexExprType}:                      _GotoState83Action,
	{_State99, PostfixableExprType}:                _GotoState97Action,
	{_State99, PostfixUnaryExprType}:               _GotoState96Action,
	{_State99, PrefixableExprType}:                 _GotoState235Action,
	{_State99, PrefixUnaryExprType}:                _GotoState98Action,
	{_State99, PrefixUnaryOpType}:                  _GotoState99Action,
	{_State99, InitializableTypeExprType}:          _GotoState84Action,
	{_State99, SliceTypeExprType}:                  _GotoState103Action,
	{_State99, ArrayTypeExprType}:                  _GotoState60Action,
	{_State99, MapTypeExprType}:                    _GotoState90Action,
	{_State99, ExplicitStructTypeExprType}:         _GotoState75Action,
	{_State99, AnonymousFuncExprType}:              _GotoState59Action,
	{_State108, LbracketToken}:                     _GotoState190Action,
	{_State108, DotToken}:                          _GotoState188Action,
	{_State108, QuestionToken}:                     _GotoState193Action,
	{_State108, ExclaimToken}:                      _GotoState189Action,
	{_State108, DollarLbracketToken}:               _GotoState187Action,
	{_State108, PostfixUnaryOpType}:                _GotoState198Action,
	{_State108, GenericTypeArgumentsType}:          _GotoState197Action,
	{_State113, LparenToken}:                       _GotoState236Action,
	{_State115, LparenToken}:                       _GotoState237Action,
	{_State116, DotToken}:                          _GotoState238Action,
	{_State116, DollarLbracketToken}:               _GotoState187Action,
	{_State116, GenericTypeArgumentsType}:          _GotoState239Action,
	{_State117, IdentifierToken}:                   _GotoState241Action,
	{_State117, UnderscoreToken}:                   _GotoState242Action,
	{_State117, UnsafeToken}:                       _GotoState54Action,
	{_State117, StructToken}:                       _GotoState50Action,
	{_State117, EnumToken}:                         _GotoState113Action,
	{_State117, TraitToken}:                        _GotoState121Action,
	{_State117, FuncToken}:                         _GotoState240Action,
	{_State117, LparenToken}:                       _GotoState117Action,
	{_State117, LbracketToken}:                     _GotoState42Action,
	{_State117, DotToken}:                          _GotoState112Action,
	{_State117, QuestionToken}:                     _GotoState119Action,
	{_State117, ExclaimToken}:                      _GotoState114Action,
	{_State117, TildeTildeToken}:                   _GotoState120Action,
	{_State117, BitNegToken}:                       _GotoState111Action,
	{_State117, BitAndToken}:                       _GotoState110Action,
	{_State117, ParseErrorToken}:                   _GotoState118Action,
	{_State117, UnsafeStatementType}:               _GotoState250Action,
	{_State117, InitializableTypeExprType}:         _GotoState130Action,
	{_State117, SliceTypeExprType}:                 _GotoState103Action,
	{_State117, ArrayTypeExprType}:                 _GotoState60Action,
	{_State117, MapTypeExprType}:                   _GotoState90Action,
	{_State117, AtomTypeExprType}:                  _GotoState123Action,
	{_State117, NamedTypeExprType}:                 _GotoState131Action,
	{_State117, InferredTypeExprType}:              _GotoState129Action,
	{_State117, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State117, ReturnableTypeExprType}:            _GotoState135Action,
	{_State117, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State117, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State117, TypeExprType}:                      _GotoState249Action,
	{_State117, BinaryTypeExprType}:                _GotoState124Action,
	{_State117, FieldDefType}:                      _GotoState243Action,
	{_State117, ProperImplicitFieldDefsType}:       _GotoState248Action,
	{_State117, ImplicitFieldDefsType}:             _GotoState245Action,
	{_State117, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State117, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State117, TraitTypeExprType}:                 _GotoState136Action,
	{_State117, ProperImplicitEnumValueDefsType}:   _GotoState247Action,
	{_State117, ImplicitEnumValueDefsType}:         _GotoState244Action,
	{_State117, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State117, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State117, FuncTypeExprType}:                  _GotoState126Action,
	{_State117, MethodSignatureType}:               _GotoState246Action,
	{_State121, LparenToken}:                       _GotoState251Action,
	{_State134, IdentifierToken}:                   _GotoState116Action,
	{_State134, UnderscoreToken}:                   _GotoState122Action,
	{_State134, StructToken}:                       _GotoState50Action,
	{_State134, EnumToken}:                         _GotoState113Action,
	{_State134, TraitToken}:                        _GotoState121Action,
	{_State134, FuncToken}:                         _GotoState115Action,
	{_State134, LparenToken}:                       _GotoState117Action,
	{_State134, LbracketToken}:                     _GotoState42Action,
	{_State134, DotToken}:                          _GotoState112Action,
	{_State134, QuestionToken}:                     _GotoState119Action,
	{_State134, ExclaimToken}:                      _GotoState114Action,
	{_State134, TildeTildeToken}:                   _GotoState120Action,
	{_State134, BitNegToken}:                       _GotoState111Action,
	{_State134, BitAndToken}:                       _GotoState110Action,
	{_State134, ParseErrorToken}:                   _GotoState118Action,
	{_State134, InitializableTypeExprType}:         _GotoState130Action,
	{_State134, SliceTypeExprType}:                 _GotoState103Action,
	{_State134, ArrayTypeExprType}:                 _GotoState60Action,
	{_State134, MapTypeExprType}:                   _GotoState90Action,
	{_State134, AtomTypeExprType}:                  _GotoState123Action,
	{_State134, NamedTypeExprType}:                 _GotoState131Action,
	{_State134, InferredTypeExprType}:              _GotoState129Action,
	{_State134, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State134, ReturnableTypeExprType}:            _GotoState252Action,
	{_State134, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State134, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State134, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State134, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State134, TraitTypeExprType}:                 _GotoState136Action,
	{_State134, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State134, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State134, FuncTypeExprType}:                  _GotoState126Action,
	{_State137, DollarLbracketToken}:               _GotoState258Action,
	{_State137, AssignToken}:                       _GotoState257Action,
	{_State137, GenericParameterDefsType}:          _GotoState259Action,
	{_State138, IdentifierToken}:                   _GotoState260Action,
	{_State138, UnderscoreToken}:                   _GotoState261Action,
	{_State138, ProperParameterDefType}:            _GotoState263Action,
	{_State138, ParameterDefType}:                  _GotoState262Action,
	{_State139, NewlinesToken}:                     _GotoState264Action,
	{_State139, SemicolonToken}:                    _GotoState265Action,
	{_State141, RbraceToken}:                       _GotoState266Action,
	{_State143, DollarLbracketToken}:               _GotoState258Action,
	{_State143, AssignToken}:                       _GotoState267Action,
	{_State143, GenericParameterDefsType}:          _GotoState268Action,
	{_State144, CommentGroupsToken}:                _GotoState9Action,
	{_State144, PackageToken}:                      _GotoState13Action,
	{_State144, TypeToken}:                         _GotoState14Action,
	{_State144, FuncToken}:                         _GotoState10Action,
	{_State144, VarToken}:                          _GotoState15Action,
	{_State144, LetToken}:                          _GotoState12Action,
	{_State144, LbraceToken}:                       _GotoState11Action,
	{_State144, DefinitionType}:                    _GotoState269Action,
	{_State144, StatementBlockType}:                _GotoState21Action,
	{_State144, VarDeclPatternType}:                _GotoState23Action,
	{_State144, VarOrLetType}:                      _GotoState24Action,
	{_State144, TypeDefType}:                       _GotoState22Action,
	{_State144, NamedFuncDefType}:                  _GotoState18Action,
	{_State144, PackageDefType}:                    _GotoState19Action,
	{_State145, IntegerLiteralToken}:               _GotoState40Action,
	{_State145, FloatLiteralToken}:                 _GotoState35Action,
	{_State145, RuneLiteralToken}:                  _GotoState48Action,
	{_State145, StringLiteralToken}:                _GotoState49Action,
	{_State145, IdentifierToken}:                   _GotoState38Action,
	{_State145, UnderscoreToken}:                   _GotoState53Action,
	{_State145, TrueToken}:                         _GotoState52Action,
	{_State145, FalseToken}:                        _GotoState34Action,
	{_State145, StructToken}:                       _GotoState50Action,
	{_State145, FuncToken}:                         _GotoState36Action,
	{_State145, VarToken}:                          _GotoState15Action,
	{_State145, LetToken}:                          _GotoState12Action,
	{_State145, NotToken}:                          _GotoState45Action,
	{_State145, LabelDeclToken}:                    _GotoState41Action,
	{_State145, LparenToken}:                       _GotoState43Action,
	{_State145, LbracketToken}:                     _GotoState42Action,
	{_State145, SubToken}:                          _GotoState51Action,
	{_State145, MulToken}:                          _GotoState44Action,
	{_State145, BitNegToken}:                       _GotoState27Action,
	{_State145, BitAndToken}:                       _GotoState26Action,
	{_State145, GreaterToken}:                      _GotoState37Action,
	{_State145, ParseErrorToken}:                   _GotoState46Action,
	{_State145, VarDeclPatternType}:                _GotoState107Action,
	{_State145, VarOrLetType}:                      _GotoState24Action,
	{_State145, ExprType}:                          _GotoState270Action,
	{_State145, OptionalLabelDeclType}:             _GotoState93Action,
	{_State145, SequenceExprType}:                  _GotoState109Action,
	{_State145, IfExprType}:                        _GotoState80Action,
	{_State145, SwitchExprType}:                    _GotoState104Action,
	{_State145, LoopExprType}:                      _GotoState89Action,
	{_State145, CallExprType}:                      _GotoState71Action,
	{_State145, AtomExprType}:                      _GotoState63Action,
	{_State145, ParseErrorExprType}:                _GotoState95Action,
	{_State145, LiteralExprType}:                   _GotoState88Action,
	{_State145, NamedExprType}:                     _GotoState92Action,
	{_State145, BlockExprType}:                     _GotoState70Action,
	{_State145, InitializeExprType}:                _GotoState85Action,
	{_State145, ImplicitStructExprType}:            _GotoState81Action,
	{_State145, AccessibleExprType}:                _GotoState108Action,
	{_State145, AccessExprType}:                    _GotoState55Action,
	{_State145, IndexExprType}:                     _GotoState83Action,
	{_State145, PostfixableExprType}:               _GotoState97Action,
	{_State145, PostfixUnaryExprType}:              _GotoState96Action,
	{_State145, PrefixableExprType}:                _GotoState100Action,
	{_State145, PrefixUnaryExprType}:               _GotoState98Action,
	{_State145, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State145, MulExprType}:                       _GotoState91Action,
	{_State145, BinaryMulExprType}:                 _GotoState67Action,
	{_State145, AddExprType}:                       _GotoState57Action,
	{_State145, BinaryAddExprType}:                 _GotoState64Action,
	{_State145, CmpExprType}:                       _GotoState74Action,
	{_State145, BinaryCmpExprType}:                 _GotoState66Action,
	{_State145, AndExprType}:                       _GotoState58Action,
	{_State145, BinaryAndExprType}:                 _GotoState65Action,
	{_State145, OrExprType}:                        _GotoState94Action,
	{_State145, BinaryOrExprType}:                  _GotoState69Action,
	{_State145, InitializableTypeExprType}:         _GotoState84Action,
	{_State145, SliceTypeExprType}:                 _GotoState103Action,
	{_State145, ArrayTypeExprType}:                 _GotoState60Action,
	{_State145, MapTypeExprType}:                   _GotoState90Action,
	{_State145, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State145, AnonymousFuncExprType}:             _GotoState59Action,
	{_State147, IdentifierToken}:                   _GotoState272Action,
	{_State147, UnderscoreToken}:                   _GotoState148Action,
	{_State147, LparenToken}:                       _GotoState147Action,
	{_State147, EllipsisToken}:                     _GotoState271Action,
	{_State147, VarPatternType}:                    _GotoState275Action,
	{_State147, TuplePatternType}:                  _GotoState149Action,
	{_State147, FieldVarPatternsType}:              _GotoState274Action,
	{_State147, FieldVarPatternType}:               _GotoState273Action,
	{_State150, IdentifierToken}:                   _GotoState116Action,
	{_State150, UnderscoreToken}:                   _GotoState122Action,
	{_State150, StructToken}:                       _GotoState50Action,
	{_State150, EnumToken}:                         _GotoState113Action,
	{_State150, TraitToken}:                        _GotoState121Action,
	{_State150, FuncToken}:                         _GotoState115Action,
	{_State150, LparenToken}:                       _GotoState117Action,
	{_State150, LbracketToken}:                     _GotoState42Action,
	{_State150, DotToken}:                          _GotoState112Action,
	{_State150, QuestionToken}:                     _GotoState119Action,
	{_State150, ExclaimToken}:                      _GotoState114Action,
	{_State150, TildeTildeToken}:                   _GotoState120Action,
	{_State150, BitNegToken}:                       _GotoState111Action,
	{_State150, BitAndToken}:                       _GotoState110Action,
	{_State150, ParseErrorToken}:                   _GotoState118Action,
	{_State150, OptionalTypeExprType}:              _GotoState276Action,
	{_State150, InitializableTypeExprType}:         _GotoState130Action,
	{_State150, SliceTypeExprType}:                 _GotoState103Action,
	{_State150, ArrayTypeExprType}:                 _GotoState60Action,
	{_State150, MapTypeExprType}:                   _GotoState90Action,
	{_State150, AtomTypeExprType}:                  _GotoState123Action,
	{_State150, NamedTypeExprType}:                 _GotoState131Action,
	{_State150, InferredTypeExprType}:              _GotoState129Action,
	{_State150, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State150, ReturnableTypeExprType}:            _GotoState135Action,
	{_State150, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State150, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State150, TypeExprType}:                      _GotoState277Action,
	{_State150, BinaryTypeExprType}:                _GotoState124Action,
	{_State150, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State150, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State150, TraitTypeExprType}:                 _GotoState136Action,
	{_State150, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State150, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State150, FuncTypeExprType}:                  _GotoState126Action,
	{_State151, IdentifierToken}:                   _GotoState278Action,
	{_State152, DotToken}:                          _GotoState279Action,
	{_State155, CommaToken}:                        _GotoState281Action,
	{_State155, ColonToken}:                        _GotoState280Action,
	{_State156, LbraceToken}:                       _GotoState11Action,
	{_State156, StatementBlockType}:                _GotoState233Action,
	{_State158, IntegerLiteralToken}:               _GotoState40Action,
	{_State158, FloatLiteralToken}:                 _GotoState35Action,
	{_State158, RuneLiteralToken}:                  _GotoState48Action,
	{_State158, StringLiteralToken}:                _GotoState49Action,
	{_State158, IdentifierToken}:                   _GotoState38Action,
	{_State158, UnderscoreToken}:                   _GotoState53Action,
	{_State158, TrueToken}:                         _GotoState52Action,
	{_State158, FalseToken}:                        _GotoState34Action,
	{_State158, ReturnToken}:                       _GotoState47Action,
	{_State158, BreakToken}:                        _GotoState28Action,
	{_State158, ContinueToken}:                     _GotoState30Action,
	{_State158, FallthroughToken}:                  _GotoState33Action,
	{_State158, UnsafeToken}:                       _GotoState54Action,
	{_State158, StructToken}:                       _GotoState50Action,
	{_State158, FuncToken}:                         _GotoState36Action,
	{_State158, AsyncToken}:                        _GotoState25Action,
	{_State158, DeferToken}:                        _GotoState32Action,
	{_State158, VarToken}:                          _GotoState15Action,
	{_State158, LetToken}:                          _GotoState12Action,
	{_State158, NotToken}:                          _GotoState45Action,
	{_State158, LabelDeclToken}:                    _GotoState41Action,
	{_State158, LparenToken}:                       _GotoState43Action,
	{_State158, LbracketToken}:                     _GotoState42Action,
	{_State158, SubToken}:                          _GotoState51Action,
	{_State158, MulToken}:                          _GotoState44Action,
	{_State158, BitNegToken}:                       _GotoState27Action,
	{_State158, BitAndToken}:                       _GotoState26Action,
	{_State158, GreaterToken}:                      _GotoState37Action,
	{_State158, ParseErrorToken}:                   _GotoState46Action,
	{_State158, SimpleStatementType}:               _GotoState283Action,
	{_State158, OptionalSimpleStatementType}:       _GotoState282Action,
	{_State158, ExprOrImproperStructStatementType}: _GotoState77Action,
	{_State158, ExprsType}:                         _GotoState78Action,
	{_State158, CallbackOpType}:                    _GotoState72Action,
	{_State158, CallbackOpStatementType}:           _GotoState73Action,
	{_State158, UnsafeStatementType}:               _GotoState106Action,
	{_State158, JumpStatementType}:                 _GotoState86Action,
	{_State158, JumpTypeType}:                      _GotoState87Action,
	{_State158, FallthroughStatementType}:          _GotoState79Action,
	{_State158, AssignStatementType}:               _GotoState62Action,
	{_State158, UnaryOpAssignStatementType}:        _GotoState105Action,
	{_State158, BinaryOpAssignStatementType}:       _GotoState68Action,
	{_State158, VarDeclPatternType}:                _GotoState107Action,
	{_State158, VarOrLetType}:                      _GotoState24Action,
	{_State158, AssignPatternType}:                 _GotoState61Action,
	{_State158, ExprType}:                          _GotoState76Action,
	{_State158, OptionalLabelDeclType}:             _GotoState93Action,
	{_State158, SequenceExprType}:                  _GotoState101Action,
	{_State158, IfExprType}:                        _GotoState80Action,
	{_State158, SwitchExprType}:                    _GotoState104Action,
	{_State158, LoopExprType}:                      _GotoState89Action,
	{_State158, CallExprType}:                      _GotoState71Action,
	{_State158, AtomExprType}:                      _GotoState63Action,
	{_State158, ParseErrorExprType}:                _GotoState95Action,
	{_State158, LiteralExprType}:                   _GotoState88Action,
	{_State158, NamedExprType}:                     _GotoState92Action,
	{_State158, BlockExprType}:                     _GotoState70Action,
	{_State158, InitializeExprType}:                _GotoState85Action,
	{_State158, ImplicitStructExprType}:            _GotoState81Action,
	{_State158, AccessibleExprType}:                _GotoState56Action,
	{_State158, AccessExprType}:                    _GotoState55Action,
	{_State158, IndexExprType}:                     _GotoState83Action,
	{_State158, PostfixableExprType}:               _GotoState97Action,
	{_State158, PostfixUnaryExprType}:              _GotoState96Action,
	{_State158, PrefixableExprType}:                _GotoState100Action,
	{_State158, PrefixUnaryExprType}:               _GotoState98Action,
	{_State158, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State158, MulExprType}:                       _GotoState91Action,
	{_State158, BinaryMulExprType}:                 _GotoState67Action,
	{_State158, AddExprType}:                       _GotoState57Action,
	{_State158, BinaryAddExprType}:                 _GotoState64Action,
	{_State158, CmpExprType}:                       _GotoState74Action,
	{_State158, BinaryCmpExprType}:                 _GotoState66Action,
	{_State158, AndExprType}:                       _GotoState58Action,
	{_State158, BinaryAndExprType}:                 _GotoState65Action,
	{_State158, OrExprType}:                        _GotoState94Action,
	{_State158, BinaryOrExprType}:                  _GotoState69Action,
	{_State158, InitializableTypeExprType}:         _GotoState84Action,
	{_State158, SliceTypeExprType}:                 _GotoState103Action,
	{_State158, ArrayTypeExprType}:                 _GotoState60Action,
	{_State158, MapTypeExprType}:                   _GotoState90Action,
	{_State158, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State158, AnonymousFuncExprType}:             _GotoState59Action,
	{_State159, IdentifierToken}:                   _GotoState260Action,
	{_State159, UnderscoreToken}:                   _GotoState261Action,
	{_State159, ProperParameterDefType}:            _GotoState263Action,
	{_State159, ParameterDefType}:                  _GotoState284Action,
	{_State159, ProperParameterDefsType}:           _GotoState286Action,
	{_State159, ParameterDefsType}:                 _GotoState285Action,
	{_State161, StringLiteralToken}:                _GotoState287Action,
	{_State162, StringLiteralToken}:                _GotoState288Action,
	{_State163, StringLiteralToken}:                _GotoState164Action,
	{_State163, IdentifierToken}:                   _GotoState162Action,
	{_State163, UnderscoreToken}:                   _GotoState165Action,
	{_State163, DotToken}:                          _GotoState161Action,
	{_State163, ImportClauseType}:                  _GotoState289Action,
	{_State163, ProperImportClausesType}:           _GotoState291Action,
	{_State163, ImportClausesType}:                 _GotoState290Action,
	{_State165, StringLiteralToken}:                _GotoState292Action,
	{_State167, RbracketToken}:                     _GotoState295Action,
	{_State167, CommaToken}:                        _GotoState294Action,
	{_State167, ColonToken}:                        _GotoState293Action,
	{_State167, AddToken}:                          _GotoState253Action,
	{_State167, SubToken}:                          _GotoState255Action,
	{_State167, MulToken}:                          _GotoState254Action,
	{_State167, BinaryTypeOpType}:                  _GotoState256Action,
	{_State168, IntegerLiteralToken}:               _GotoState40Action,
	{_State168, FloatLiteralToken}:                 _GotoState35Action,
	{_State168, RuneLiteralToken}:                  _GotoState48Action,
	{_State168, StringLiteralToken}:                _GotoState49Action,
	{_State168, IdentifierToken}:                   _GotoState38Action,
	{_State168, UnderscoreToken}:                   _GotoState53Action,
	{_State168, TrueToken}:                         _GotoState52Action,
	{_State168, FalseToken}:                        _GotoState34Action,
	{_State168, StructToken}:                       _GotoState50Action,
	{_State168, FuncToken}:                         _GotoState36Action,
	{_State168, VarToken}:                          _GotoState15Action,
	{_State168, LetToken}:                          _GotoState12Action,
	{_State168, NotToken}:                          _GotoState45Action,
	{_State168, LabelDeclToken}:                    _GotoState41Action,
	{_State168, LparenToken}:                       _GotoState43Action,
	{_State168, LbracketToken}:                     _GotoState42Action,
	{_State168, SubToken}:                          _GotoState51Action,
	{_State168, MulToken}:                          _GotoState44Action,
	{_State168, BitNegToken}:                       _GotoState27Action,
	{_State168, BitAndToken}:                       _GotoState26Action,
	{_State168, GreaterToken}:                      _GotoState37Action,
	{_State168, ParseErrorToken}:                   _GotoState46Action,
	{_State168, VarDeclPatternType}:                _GotoState107Action,
	{_State168, VarOrLetType}:                      _GotoState24Action,
	{_State168, ExprType}:                          _GotoState296Action,
	{_State168, OptionalLabelDeclType}:             _GotoState93Action,
	{_State168, SequenceExprType}:                  _GotoState109Action,
	{_State168, IfExprType}:                        _GotoState80Action,
	{_State168, SwitchExprType}:                    _GotoState104Action,
	{_State168, LoopExprType}:                      _GotoState89Action,
	{_State168, CallExprType}:                      _GotoState71Action,
	{_State168, AtomExprType}:                      _GotoState63Action,
	{_State168, ParseErrorExprType}:                _GotoState95Action,
	{_State168, LiteralExprType}:                   _GotoState88Action,
	{_State168, NamedExprType}:                     _GotoState92Action,
	{_State168, BlockExprType}:                     _GotoState70Action,
	{_State168, InitializeExprType}:                _GotoState85Action,
	{_State168, ImplicitStructExprType}:            _GotoState81Action,
	{_State168, AccessibleExprType}:                _GotoState108Action,
	{_State168, AccessExprType}:                    _GotoState55Action,
	{_State168, IndexExprType}:                     _GotoState83Action,
	{_State168, PostfixableExprType}:               _GotoState97Action,
	{_State168, PostfixUnaryExprType}:              _GotoState96Action,
	{_State168, PrefixableExprType}:                _GotoState100Action,
	{_State168, PrefixUnaryExprType}:               _GotoState98Action,
	{_State168, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State168, MulExprType}:                       _GotoState91Action,
	{_State168, BinaryMulExprType}:                 _GotoState67Action,
	{_State168, AddExprType}:                       _GotoState57Action,
	{_State168, BinaryAddExprType}:                 _GotoState64Action,
	{_State168, CmpExprType}:                       _GotoState74Action,
	{_State168, BinaryCmpExprType}:                 _GotoState66Action,
	{_State168, AndExprType}:                       _GotoState58Action,
	{_State168, BinaryAndExprType}:                 _GotoState65Action,
	{_State168, OrExprType}:                        _GotoState94Action,
	{_State168, BinaryOrExprType}:                  _GotoState69Action,
	{_State168, InitializableTypeExprType}:         _GotoState84Action,
	{_State168, SliceTypeExprType}:                 _GotoState103Action,
	{_State168, ArrayTypeExprType}:                 _GotoState60Action,
	{_State168, MapTypeExprType}:                   _GotoState90Action,
	{_State168, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State168, AnonymousFuncExprType}:             _GotoState59Action,
	{_State170, AssignToken}:                       _GotoState297Action,
	{_State172, RparenToken}:                       _GotoState298Action,
	{_State173, ColonToken}:                        _GotoState299Action,
	{_State174, ColonToken}:                        _GotoState300Action,
	{_State174, EllipsisToken}:                     _GotoState301Action,
	{_State175, CommaToken}:                        _GotoState302Action,
	{_State176, IdentifierToken}:                   _GotoState241Action,
	{_State176, UnderscoreToken}:                   _GotoState242Action,
	{_State176, UnsafeToken}:                       _GotoState54Action,
	{_State176, StructToken}:                       _GotoState50Action,
	{_State176, EnumToken}:                         _GotoState113Action,
	{_State176, TraitToken}:                        _GotoState121Action,
	{_State176, FuncToken}:                         _GotoState240Action,
	{_State176, LparenToken}:                       _GotoState117Action,
	{_State176, LbracketToken}:                     _GotoState42Action,
	{_State176, DotToken}:                          _GotoState112Action,
	{_State176, QuestionToken}:                     _GotoState119Action,
	{_State176, ExclaimToken}:                      _GotoState114Action,
	{_State176, TildeTildeToken}:                   _GotoState120Action,
	{_State176, BitNegToken}:                       _GotoState111Action,
	{_State176, BitAndToken}:                       _GotoState110Action,
	{_State176, ParseErrorToken}:                   _GotoState118Action,
	{_State176, UnsafeStatementType}:               _GotoState250Action,
	{_State176, InitializableTypeExprType}:         _GotoState130Action,
	{_State176, SliceTypeExprType}:                 _GotoState103Action,
	{_State176, ArrayTypeExprType}:                 _GotoState60Action,
	{_State176, MapTypeExprType}:                   _GotoState90Action,
	{_State176, AtomTypeExprType}:                  _GotoState123Action,
	{_State176, NamedTypeExprType}:                 _GotoState131Action,
	{_State176, InferredTypeExprType}:              _GotoState129Action,
	{_State176, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State176, ReturnableTypeExprType}:            _GotoState135Action,
	{_State176, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State176, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State176, TypeExprType}:                      _GotoState249Action,
	{_State176, BinaryTypeExprType}:                _GotoState124Action,
	{_State176, FieldDefType}:                      _GotoState304Action,
	{_State176, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State176, ProperExplicitFieldDefsType}:       _GotoState305Action,
	{_State176, ExplicitFieldDefsType}:             _GotoState303Action,
	{_State176, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State176, TraitTypeExprType}:                 _GotoState136Action,
	{_State176, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State176, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State176, FuncTypeExprType}:                  _GotoState126Action,
	{_State176, MethodSignatureType}:               _GotoState246Action,
	{_State177, IdentifierToken}:                   _GotoState306Action,
	{_State187, IdentifierToken}:                   _GotoState116Action,
	{_State187, UnderscoreToken}:                   _GotoState122Action,
	{_State187, StructToken}:                       _GotoState50Action,
	{_State187, EnumToken}:                         _GotoState113Action,
	{_State187, TraitToken}:                        _GotoState121Action,
	{_State187, FuncToken}:                         _GotoState115Action,
	{_State187, LparenToken}:                       _GotoState117Action,
	{_State187, LbracketToken}:                     _GotoState42Action,
	{_State187, DotToken}:                          _GotoState112Action,
	{_State187, QuestionToken}:                     _GotoState119Action,
	{_State187, ExclaimToken}:                      _GotoState114Action,
	{_State187, TildeTildeToken}:                   _GotoState120Action,
	{_State187, BitNegToken}:                       _GotoState111Action,
	{_State187, BitAndToken}:                       _GotoState110Action,
	{_State187, ParseErrorToken}:                   _GotoState118Action,
	{_State187, InitializableTypeExprType}:         _GotoState130Action,
	{_State187, SliceTypeExprType}:                 _GotoState103Action,
	{_State187, ArrayTypeExprType}:                 _GotoState60Action,
	{_State187, MapTypeExprType}:                   _GotoState90Action,
	{_State187, AtomTypeExprType}:                  _GotoState123Action,
	{_State187, NamedTypeExprType}:                 _GotoState131Action,
	{_State187, InferredTypeExprType}:              _GotoState129Action,
	{_State187, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State187, ReturnableTypeExprType}:            _GotoState135Action,
	{_State187, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State187, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State187, TypeExprType}:                      _GotoState309Action,
	{_State187, BinaryTypeExprType}:                _GotoState124Action,
	{_State187, ProperTypeArgumentsType}:           _GotoState307Action,
	{_State187, TypeArgumentsType}:                 _GotoState308Action,
	{_State187, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State187, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State187, TraitTypeExprType}:                 _GotoState136Action,
	{_State187, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State187, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State187, FuncTypeExprType}:                  _GotoState126Action,
	{_State188, IdentifierToken}:                   _GotoState310Action,
	{_State190, IntegerLiteralToken}:               _GotoState40Action,
	{_State190, FloatLiteralToken}:                 _GotoState35Action,
	{_State190, RuneLiteralToken}:                  _GotoState48Action,
	{_State190, StringLiteralToken}:                _GotoState49Action,
	{_State190, IdentifierToken}:                   _GotoState170Action,
	{_State190, UnderscoreToken}:                   _GotoState53Action,
	{_State190, TrueToken}:                         _GotoState52Action,
	{_State190, FalseToken}:                        _GotoState34Action,
	{_State190, StructToken}:                       _GotoState50Action,
	{_State190, FuncToken}:                         _GotoState36Action,
	{_State190, VarToken}:                          _GotoState15Action,
	{_State190, LetToken}:                          _GotoState12Action,
	{_State190, NotToken}:                          _GotoState45Action,
	{_State190, LabelDeclToken}:                    _GotoState41Action,
	{_State190, LparenToken}:                       _GotoState43Action,
	{_State190, LbracketToken}:                     _GotoState42Action,
	{_State190, ColonToken}:                        _GotoState168Action,
	{_State190, EllipsisToken}:                     _GotoState169Action,
	{_State190, SubToken}:                          _GotoState51Action,
	{_State190, MulToken}:                          _GotoState44Action,
	{_State190, BitNegToken}:                       _GotoState27Action,
	{_State190, BitAndToken}:                       _GotoState26Action,
	{_State190, GreaterToken}:                      _GotoState37Action,
	{_State190, ParseErrorToken}:                   _GotoState46Action,
	{_State190, VarDeclPatternType}:                _GotoState107Action,
	{_State190, VarOrLetType}:                      _GotoState24Action,
	{_State190, ExprType}:                          _GotoState174Action,
	{_State190, OptionalLabelDeclType}:             _GotoState93Action,
	{_State190, SequenceExprType}:                  _GotoState109Action,
	{_State190, IfExprType}:                        _GotoState80Action,
	{_State190, SwitchExprType}:                    _GotoState104Action,
	{_State190, LoopExprType}:                      _GotoState89Action,
	{_State190, CallExprType}:                      _GotoState71Action,
	{_State190, ArgumentType}:                      _GotoState311Action,
	{_State190, ColonExprType}:                     _GotoState173Action,
	{_State190, AtomExprType}:                      _GotoState63Action,
	{_State190, ParseErrorExprType}:                _GotoState95Action,
	{_State190, LiteralExprType}:                   _GotoState88Action,
	{_State190, NamedExprType}:                     _GotoState92Action,
	{_State190, BlockExprType}:                     _GotoState70Action,
	{_State190, InitializeExprType}:                _GotoState85Action,
	{_State190, ImplicitStructExprType}:            _GotoState81Action,
	{_State190, AccessibleExprType}:                _GotoState108Action,
	{_State190, AccessExprType}:                    _GotoState55Action,
	{_State190, IndexExprType}:                     _GotoState83Action,
	{_State190, PostfixableExprType}:               _GotoState97Action,
	{_State190, PostfixUnaryExprType}:              _GotoState96Action,
	{_State190, PrefixableExprType}:                _GotoState100Action,
	{_State190, PrefixUnaryExprType}:               _GotoState98Action,
	{_State190, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State190, MulExprType}:                       _GotoState91Action,
	{_State190, BinaryMulExprType}:                 _GotoState67Action,
	{_State190, AddExprType}:                       _GotoState57Action,
	{_State190, BinaryAddExprType}:                 _GotoState64Action,
	{_State190, CmpExprType}:                       _GotoState74Action,
	{_State190, BinaryCmpExprType}:                 _GotoState66Action,
	{_State190, AndExprType}:                       _GotoState58Action,
	{_State190, BinaryAndExprType}:                 _GotoState65Action,
	{_State190, OrExprType}:                        _GotoState94Action,
	{_State190, BinaryOrExprType}:                  _GotoState69Action,
	{_State190, InitializableTypeExprType}:         _GotoState84Action,
	{_State190, SliceTypeExprType}:                 _GotoState103Action,
	{_State190, ArrayTypeExprType}:                 _GotoState60Action,
	{_State190, MapTypeExprType}:                   _GotoState90Action,
	{_State190, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State190, AnonymousFuncExprType}:             _GotoState59Action,
	{_State196, IntegerLiteralToken}:               _GotoState40Action,
	{_State196, FloatLiteralToken}:                 _GotoState35Action,
	{_State196, RuneLiteralToken}:                  _GotoState48Action,
	{_State196, StringLiteralToken}:                _GotoState49Action,
	{_State196, IdentifierToken}:                   _GotoState38Action,
	{_State196, UnderscoreToken}:                   _GotoState53Action,
	{_State196, TrueToken}:                         _GotoState52Action,
	{_State196, FalseToken}:                        _GotoState34Action,
	{_State196, StructToken}:                       _GotoState50Action,
	{_State196, FuncToken}:                         _GotoState36Action,
	{_State196, VarToken}:                          _GotoState15Action,
	{_State196, LetToken}:                          _GotoState12Action,
	{_State196, NotToken}:                          _GotoState45Action,
	{_State196, LabelDeclToken}:                    _GotoState41Action,
	{_State196, LparenToken}:                       _GotoState43Action,
	{_State196, LbracketToken}:                     _GotoState42Action,
	{_State196, SubToken}:                          _GotoState51Action,
	{_State196, MulToken}:                          _GotoState44Action,
	{_State196, BitNegToken}:                       _GotoState27Action,
	{_State196, BitAndToken}:                       _GotoState26Action,
	{_State196, GreaterToken}:                      _GotoState37Action,
	{_State196, ParseErrorToken}:                   _GotoState46Action,
	{_State196, VarDeclPatternType}:                _GotoState107Action,
	{_State196, VarOrLetType}:                      _GotoState24Action,
	{_State196, ExprType}:                          _GotoState312Action,
	{_State196, OptionalLabelDeclType}:             _GotoState93Action,
	{_State196, SequenceExprType}:                  _GotoState109Action,
	{_State196, IfExprType}:                        _GotoState80Action,
	{_State196, SwitchExprType}:                    _GotoState104Action,
	{_State196, LoopExprType}:                      _GotoState89Action,
	{_State196, CallExprType}:                      _GotoState71Action,
	{_State196, AtomExprType}:                      _GotoState63Action,
	{_State196, ParseErrorExprType}:                _GotoState95Action,
	{_State196, LiteralExprType}:                   _GotoState88Action,
	{_State196, NamedExprType}:                     _GotoState92Action,
	{_State196, BlockExprType}:                     _GotoState70Action,
	{_State196, InitializeExprType}:                _GotoState85Action,
	{_State196, ImplicitStructExprType}:            _GotoState81Action,
	{_State196, AccessibleExprType}:                _GotoState108Action,
	{_State196, AccessExprType}:                    _GotoState55Action,
	{_State196, IndexExprType}:                     _GotoState83Action,
	{_State196, PostfixableExprType}:               _GotoState97Action,
	{_State196, PostfixUnaryExprType}:              _GotoState96Action,
	{_State196, PrefixableExprType}:                _GotoState100Action,
	{_State196, PrefixUnaryExprType}:               _GotoState98Action,
	{_State196, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State196, MulExprType}:                       _GotoState91Action,
	{_State196, BinaryMulExprType}:                 _GotoState67Action,
	{_State196, AddExprType}:                       _GotoState57Action,
	{_State196, BinaryAddExprType}:                 _GotoState64Action,
	{_State196, CmpExprType}:                       _GotoState74Action,
	{_State196, BinaryCmpExprType}:                 _GotoState66Action,
	{_State196, AndExprType}:                       _GotoState58Action,
	{_State196, BinaryAndExprType}:                 _GotoState65Action,
	{_State196, OrExprType}:                        _GotoState94Action,
	{_State196, BinaryOrExprType}:                  _GotoState69Action,
	{_State196, InitializableTypeExprType}:         _GotoState84Action,
	{_State196, SliceTypeExprType}:                 _GotoState103Action,
	{_State196, ArrayTypeExprType}:                 _GotoState60Action,
	{_State196, MapTypeExprType}:                   _GotoState90Action,
	{_State196, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State196, AnonymousFuncExprType}:             _GotoState59Action,
	{_State197, LparenToken}:                       _GotoState313Action,
	{_State204, IntegerLiteralToken}:               _GotoState40Action,
	{_State204, FloatLiteralToken}:                 _GotoState35Action,
	{_State204, RuneLiteralToken}:                  _GotoState48Action,
	{_State204, StringLiteralToken}:                _GotoState49Action,
	{_State204, IdentifierToken}:                   _GotoState38Action,
	{_State204, UnderscoreToken}:                   _GotoState53Action,
	{_State204, TrueToken}:                         _GotoState52Action,
	{_State204, FalseToken}:                        _GotoState34Action,
	{_State204, StructToken}:                       _GotoState50Action,
	{_State204, FuncToken}:                         _GotoState36Action,
	{_State204, NotToken}:                          _GotoState45Action,
	{_State204, LabelDeclToken}:                    _GotoState41Action,
	{_State204, LparenToken}:                       _GotoState43Action,
	{_State204, LbracketToken}:                     _GotoState42Action,
	{_State204, SubToken}:                          _GotoState51Action,
	{_State204, MulToken}:                          _GotoState44Action,
	{_State204, BitNegToken}:                       _GotoState27Action,
	{_State204, BitAndToken}:                       _GotoState26Action,
	{_State204, ParseErrorToken}:                   _GotoState46Action,
	{_State204, OptionalLabelDeclType}:             _GotoState156Action,
	{_State204, CallExprType}:                      _GotoState71Action,
	{_State204, AtomExprType}:                      _GotoState63Action,
	{_State204, ParseErrorExprType}:                _GotoState95Action,
	{_State204, LiteralExprType}:                   _GotoState88Action,
	{_State204, NamedExprType}:                     _GotoState92Action,
	{_State204, BlockExprType}:                     _GotoState70Action,
	{_State204, InitializeExprType}:                _GotoState85Action,
	{_State204, ImplicitStructExprType}:            _GotoState81Action,
	{_State204, AccessibleExprType}:                _GotoState108Action,
	{_State204, AccessExprType}:                    _GotoState55Action,
	{_State204, IndexExprType}:                     _GotoState83Action,
	{_State204, PostfixableExprType}:               _GotoState97Action,
	{_State204, PostfixUnaryExprType}:              _GotoState96Action,
	{_State204, PrefixableExprType}:                _GotoState100Action,
	{_State204, PrefixUnaryExprType}:               _GotoState98Action,
	{_State204, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State204, MulExprType}:                       _GotoState314Action,
	{_State204, BinaryMulExprType}:                 _GotoState67Action,
	{_State204, InitializableTypeExprType}:         _GotoState84Action,
	{_State204, SliceTypeExprType}:                 _GotoState103Action,
	{_State204, ArrayTypeExprType}:                 _GotoState60Action,
	{_State204, MapTypeExprType}:                   _GotoState90Action,
	{_State204, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State204, AnonymousFuncExprType}:             _GotoState59Action,
	{_State205, IntegerLiteralToken}:               _GotoState40Action,
	{_State205, FloatLiteralToken}:                 _GotoState35Action,
	{_State205, RuneLiteralToken}:                  _GotoState48Action,
	{_State205, StringLiteralToken}:                _GotoState49Action,
	{_State205, IdentifierToken}:                   _GotoState38Action,
	{_State205, UnderscoreToken}:                   _GotoState53Action,
	{_State205, TrueToken}:                         _GotoState52Action,
	{_State205, FalseToken}:                        _GotoState34Action,
	{_State205, StructToken}:                       _GotoState50Action,
	{_State205, FuncToken}:                         _GotoState36Action,
	{_State205, NotToken}:                          _GotoState45Action,
	{_State205, LabelDeclToken}:                    _GotoState41Action,
	{_State205, LparenToken}:                       _GotoState43Action,
	{_State205, LbracketToken}:                     _GotoState42Action,
	{_State205, SubToken}:                          _GotoState51Action,
	{_State205, MulToken}:                          _GotoState44Action,
	{_State205, BitNegToken}:                       _GotoState27Action,
	{_State205, BitAndToken}:                       _GotoState26Action,
	{_State205, ParseErrorToken}:                   _GotoState46Action,
	{_State205, OptionalLabelDeclType}:             _GotoState156Action,
	{_State205, CallExprType}:                      _GotoState71Action,
	{_State205, AtomExprType}:                      _GotoState63Action,
	{_State205, ParseErrorExprType}:                _GotoState95Action,
	{_State205, LiteralExprType}:                   _GotoState88Action,
	{_State205, NamedExprType}:                     _GotoState92Action,
	{_State205, BlockExprType}:                     _GotoState70Action,
	{_State205, InitializeExprType}:                _GotoState85Action,
	{_State205, ImplicitStructExprType}:            _GotoState81Action,
	{_State205, AccessibleExprType}:                _GotoState108Action,
	{_State205, AccessExprType}:                    _GotoState55Action,
	{_State205, IndexExprType}:                     _GotoState83Action,
	{_State205, PostfixableExprType}:               _GotoState97Action,
	{_State205, PostfixUnaryExprType}:              _GotoState96Action,
	{_State205, PrefixableExprType}:                _GotoState100Action,
	{_State205, PrefixUnaryExprType}:               _GotoState98Action,
	{_State205, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State205, MulExprType}:                       _GotoState91Action,
	{_State205, BinaryMulExprType}:                 _GotoState67Action,
	{_State205, AddExprType}:                       _GotoState57Action,
	{_State205, BinaryAddExprType}:                 _GotoState64Action,
	{_State205, CmpExprType}:                       _GotoState315Action,
	{_State205, BinaryCmpExprType}:                 _GotoState66Action,
	{_State205, InitializableTypeExprType}:         _GotoState84Action,
	{_State205, SliceTypeExprType}:                 _GotoState103Action,
	{_State205, ArrayTypeExprType}:                 _GotoState60Action,
	{_State205, MapTypeExprType}:                   _GotoState90Action,
	{_State205, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State205, AnonymousFuncExprType}:             _GotoState59Action,
	{_State206, IntegerLiteralToken}:               _GotoState40Action,
	{_State206, FloatLiteralToken}:                 _GotoState35Action,
	{_State206, RuneLiteralToken}:                  _GotoState48Action,
	{_State206, StringLiteralToken}:                _GotoState49Action,
	{_State206, IdentifierToken}:                   _GotoState38Action,
	{_State206, UnderscoreToken}:                   _GotoState53Action,
	{_State206, TrueToken}:                         _GotoState52Action,
	{_State206, FalseToken}:                        _GotoState34Action,
	{_State206, StructToken}:                       _GotoState50Action,
	{_State206, FuncToken}:                         _GotoState36Action,
	{_State206, VarToken}:                          _GotoState15Action,
	{_State206, LetToken}:                          _GotoState12Action,
	{_State206, NotToken}:                          _GotoState45Action,
	{_State206, LabelDeclToken}:                    _GotoState41Action,
	{_State206, LparenToken}:                       _GotoState43Action,
	{_State206, LbracketToken}:                     _GotoState42Action,
	{_State206, SubToken}:                          _GotoState51Action,
	{_State206, MulToken}:                          _GotoState44Action,
	{_State206, BitNegToken}:                       _GotoState27Action,
	{_State206, BitAndToken}:                       _GotoState26Action,
	{_State206, GreaterToken}:                      _GotoState37Action,
	{_State206, ParseErrorToken}:                   _GotoState46Action,
	{_State206, VarDeclPatternType}:                _GotoState107Action,
	{_State206, VarOrLetType}:                      _GotoState24Action,
	{_State206, ExprType}:                          _GotoState316Action,
	{_State206, OptionalLabelDeclType}:             _GotoState93Action,
	{_State206, SequenceExprType}:                  _GotoState109Action,
	{_State206, IfExprType}:                        _GotoState80Action,
	{_State206, SwitchExprType}:                    _GotoState104Action,
	{_State206, LoopExprType}:                      _GotoState89Action,
	{_State206, CallExprType}:                      _GotoState71Action,
	{_State206, AtomExprType}:                      _GotoState63Action,
	{_State206, ParseErrorExprType}:                _GotoState95Action,
	{_State206, LiteralExprType}:                   _GotoState88Action,
	{_State206, NamedExprType}:                     _GotoState92Action,
	{_State206, BlockExprType}:                     _GotoState70Action,
	{_State206, InitializeExprType}:                _GotoState85Action,
	{_State206, ImplicitStructExprType}:            _GotoState81Action,
	{_State206, AccessibleExprType}:                _GotoState108Action,
	{_State206, AccessExprType}:                    _GotoState55Action,
	{_State206, IndexExprType}:                     _GotoState83Action,
	{_State206, PostfixableExprType}:               _GotoState97Action,
	{_State206, PostfixUnaryExprType}:              _GotoState96Action,
	{_State206, PrefixableExprType}:                _GotoState100Action,
	{_State206, PrefixUnaryExprType}:               _GotoState98Action,
	{_State206, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State206, MulExprType}:                       _GotoState91Action,
	{_State206, BinaryMulExprType}:                 _GotoState67Action,
	{_State206, AddExprType}:                       _GotoState57Action,
	{_State206, BinaryAddExprType}:                 _GotoState64Action,
	{_State206, CmpExprType}:                       _GotoState74Action,
	{_State206, BinaryCmpExprType}:                 _GotoState66Action,
	{_State206, AndExprType}:                       _GotoState58Action,
	{_State206, BinaryAndExprType}:                 _GotoState65Action,
	{_State206, OrExprType}:                        _GotoState94Action,
	{_State206, BinaryOrExprType}:                  _GotoState69Action,
	{_State206, InitializableTypeExprType}:         _GotoState84Action,
	{_State206, SliceTypeExprType}:                 _GotoState103Action,
	{_State206, ArrayTypeExprType}:                 _GotoState60Action,
	{_State206, MapTypeExprType}:                   _GotoState90Action,
	{_State206, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State206, AnonymousFuncExprType}:             _GotoState59Action,
	{_State207, LbracketToken}:                     _GotoState190Action,
	{_State207, DotToken}:                          _GotoState188Action,
	{_State207, DollarLbracketToken}:               _GotoState187Action,
	{_State207, GenericTypeArgumentsType}:          _GotoState197Action,
	{_State215, IntegerLiteralToken}:               _GotoState40Action,
	{_State215, FloatLiteralToken}:                 _GotoState35Action,
	{_State215, RuneLiteralToken}:                  _GotoState48Action,
	{_State215, StringLiteralToken}:                _GotoState49Action,
	{_State215, IdentifierToken}:                   _GotoState38Action,
	{_State215, UnderscoreToken}:                   _GotoState53Action,
	{_State215, TrueToken}:                         _GotoState52Action,
	{_State215, FalseToken}:                        _GotoState34Action,
	{_State215, StructToken}:                       _GotoState50Action,
	{_State215, FuncToken}:                         _GotoState36Action,
	{_State215, NotToken}:                          _GotoState45Action,
	{_State215, LabelDeclToken}:                    _GotoState41Action,
	{_State215, LparenToken}:                       _GotoState43Action,
	{_State215, LbracketToken}:                     _GotoState42Action,
	{_State215, SubToken}:                          _GotoState51Action,
	{_State215, MulToken}:                          _GotoState44Action,
	{_State215, BitNegToken}:                       _GotoState27Action,
	{_State215, BitAndToken}:                       _GotoState26Action,
	{_State215, ParseErrorToken}:                   _GotoState46Action,
	{_State215, OptionalLabelDeclType}:             _GotoState156Action,
	{_State215, CallExprType}:                      _GotoState71Action,
	{_State215, AtomExprType}:                      _GotoState63Action,
	{_State215, ParseErrorExprType}:                _GotoState95Action,
	{_State215, LiteralExprType}:                   _GotoState88Action,
	{_State215, NamedExprType}:                     _GotoState92Action,
	{_State215, BlockExprType}:                     _GotoState70Action,
	{_State215, InitializeExprType}:                _GotoState85Action,
	{_State215, ImplicitStructExprType}:            _GotoState81Action,
	{_State215, AccessibleExprType}:                _GotoState108Action,
	{_State215, AccessExprType}:                    _GotoState55Action,
	{_State215, IndexExprType}:                     _GotoState83Action,
	{_State215, PostfixableExprType}:               _GotoState97Action,
	{_State215, PostfixUnaryExprType}:              _GotoState96Action,
	{_State215, PrefixableExprType}:                _GotoState100Action,
	{_State215, PrefixUnaryExprType}:               _GotoState98Action,
	{_State215, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State215, MulExprType}:                       _GotoState91Action,
	{_State215, BinaryMulExprType}:                 _GotoState67Action,
	{_State215, AddExprType}:                       _GotoState317Action,
	{_State215, BinaryAddExprType}:                 _GotoState64Action,
	{_State215, InitializableTypeExprType}:         _GotoState84Action,
	{_State215, SliceTypeExprType}:                 _GotoState103Action,
	{_State215, ArrayTypeExprType}:                 _GotoState60Action,
	{_State215, MapTypeExprType}:                   _GotoState90Action,
	{_State215, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State215, AnonymousFuncExprType}:             _GotoState59Action,
	{_State216, IntegerLiteralToken}:               _GotoState40Action,
	{_State216, FloatLiteralToken}:                 _GotoState35Action,
	{_State216, RuneLiteralToken}:                  _GotoState48Action,
	{_State216, StringLiteralToken}:                _GotoState49Action,
	{_State216, IdentifierToken}:                   _GotoState38Action,
	{_State216, UnderscoreToken}:                   _GotoState53Action,
	{_State216, TrueToken}:                         _GotoState52Action,
	{_State216, FalseToken}:                        _GotoState34Action,
	{_State216, StructToken}:                       _GotoState50Action,
	{_State216, FuncToken}:                         _GotoState36Action,
	{_State216, VarToken}:                          _GotoState15Action,
	{_State216, LetToken}:                          _GotoState12Action,
	{_State216, NotToken}:                          _GotoState45Action,
	{_State216, LabelDeclToken}:                    _GotoState41Action,
	{_State216, LparenToken}:                       _GotoState43Action,
	{_State216, LbracketToken}:                     _GotoState42Action,
	{_State216, SubToken}:                          _GotoState51Action,
	{_State216, MulToken}:                          _GotoState44Action,
	{_State216, BitNegToken}:                       _GotoState27Action,
	{_State216, BitAndToken}:                       _GotoState26Action,
	{_State216, GreaterToken}:                      _GotoState37Action,
	{_State216, ParseErrorToken}:                   _GotoState46Action,
	{_State216, VarDeclPatternType}:                _GotoState107Action,
	{_State216, VarOrLetType}:                      _GotoState24Action,
	{_State216, ExprType}:                          _GotoState318Action,
	{_State216, OptionalLabelDeclType}:             _GotoState93Action,
	{_State216, SequenceExprType}:                  _GotoState109Action,
	{_State216, IfExprType}:                        _GotoState80Action,
	{_State216, SwitchExprType}:                    _GotoState104Action,
	{_State216, LoopExprType}:                      _GotoState89Action,
	{_State216, CallExprType}:                      _GotoState71Action,
	{_State216, AtomExprType}:                      _GotoState63Action,
	{_State216, ParseErrorExprType}:                _GotoState95Action,
	{_State216, LiteralExprType}:                   _GotoState88Action,
	{_State216, NamedExprType}:                     _GotoState92Action,
	{_State216, BlockExprType}:                     _GotoState70Action,
	{_State216, InitializeExprType}:                _GotoState85Action,
	{_State216, ImplicitStructExprType}:            _GotoState81Action,
	{_State216, AccessibleExprType}:                _GotoState108Action,
	{_State216, AccessExprType}:                    _GotoState55Action,
	{_State216, IndexExprType}:                     _GotoState83Action,
	{_State216, PostfixableExprType}:               _GotoState97Action,
	{_State216, PostfixUnaryExprType}:              _GotoState96Action,
	{_State216, PrefixableExprType}:                _GotoState100Action,
	{_State216, PrefixUnaryExprType}:               _GotoState98Action,
	{_State216, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State216, MulExprType}:                       _GotoState91Action,
	{_State216, BinaryMulExprType}:                 _GotoState67Action,
	{_State216, AddExprType}:                       _GotoState57Action,
	{_State216, BinaryAddExprType}:                 _GotoState64Action,
	{_State216, CmpExprType}:                       _GotoState74Action,
	{_State216, BinaryCmpExprType}:                 _GotoState66Action,
	{_State216, AndExprType}:                       _GotoState58Action,
	{_State216, BinaryAndExprType}:                 _GotoState65Action,
	{_State216, OrExprType}:                        _GotoState94Action,
	{_State216, BinaryOrExprType}:                  _GotoState69Action,
	{_State216, InitializableTypeExprType}:         _GotoState84Action,
	{_State216, SliceTypeExprType}:                 _GotoState103Action,
	{_State216, ArrayTypeExprType}:                 _GotoState60Action,
	{_State216, MapTypeExprType}:                   _GotoState90Action,
	{_State216, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State216, AnonymousFuncExprType}:             _GotoState59Action,
	{_State217, IntegerLiteralToken}:               _GotoState40Action,
	{_State217, FloatLiteralToken}:                 _GotoState35Action,
	{_State217, RuneLiteralToken}:                  _GotoState48Action,
	{_State217, StringLiteralToken}:                _GotoState49Action,
	{_State217, IdentifierToken}:                   _GotoState170Action,
	{_State217, UnderscoreToken}:                   _GotoState53Action,
	{_State217, TrueToken}:                         _GotoState52Action,
	{_State217, FalseToken}:                        _GotoState34Action,
	{_State217, StructToken}:                       _GotoState50Action,
	{_State217, FuncToken}:                         _GotoState36Action,
	{_State217, VarToken}:                          _GotoState15Action,
	{_State217, LetToken}:                          _GotoState12Action,
	{_State217, NotToken}:                          _GotoState45Action,
	{_State217, LabelDeclToken}:                    _GotoState41Action,
	{_State217, LparenToken}:                       _GotoState43Action,
	{_State217, LbracketToken}:                     _GotoState42Action,
	{_State217, ColonToken}:                        _GotoState168Action,
	{_State217, EllipsisToken}:                     _GotoState169Action,
	{_State217, SubToken}:                          _GotoState51Action,
	{_State217, MulToken}:                          _GotoState44Action,
	{_State217, BitNegToken}:                       _GotoState27Action,
	{_State217, BitAndToken}:                       _GotoState26Action,
	{_State217, GreaterToken}:                      _GotoState37Action,
	{_State217, ParseErrorToken}:                   _GotoState46Action,
	{_State217, VarDeclPatternType}:                _GotoState107Action,
	{_State217, VarOrLetType}:                      _GotoState24Action,
	{_State217, ExprType}:                          _GotoState174Action,
	{_State217, OptionalLabelDeclType}:             _GotoState93Action,
	{_State217, SequenceExprType}:                  _GotoState109Action,
	{_State217, IfExprType}:                        _GotoState80Action,
	{_State217, SwitchExprType}:                    _GotoState104Action,
	{_State217, LoopExprType}:                      _GotoState89Action,
	{_State217, CallExprType}:                      _GotoState71Action,
	{_State217, ProperArgumentsType}:               _GotoState175Action,
	{_State217, ArgumentsType}:                     _GotoState319Action,
	{_State217, ArgumentType}:                      _GotoState171Action,
	{_State217, ColonExprType}:                     _GotoState173Action,
	{_State217, AtomExprType}:                      _GotoState63Action,
	{_State217, ParseErrorExprType}:                _GotoState95Action,
	{_State217, LiteralExprType}:                   _GotoState88Action,
	{_State217, NamedExprType}:                     _GotoState92Action,
	{_State217, BlockExprType}:                     _GotoState70Action,
	{_State217, InitializeExprType}:                _GotoState85Action,
	{_State217, ImplicitStructExprType}:            _GotoState81Action,
	{_State217, AccessibleExprType}:                _GotoState108Action,
	{_State217, AccessExprType}:                    _GotoState55Action,
	{_State217, IndexExprType}:                     _GotoState83Action,
	{_State217, PostfixableExprType}:               _GotoState97Action,
	{_State217, PostfixUnaryExprType}:              _GotoState96Action,
	{_State217, PrefixableExprType}:                _GotoState100Action,
	{_State217, PrefixUnaryExprType}:               _GotoState98Action,
	{_State217, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State217, MulExprType}:                       _GotoState91Action,
	{_State217, BinaryMulExprType}:                 _GotoState67Action,
	{_State217, AddExprType}:                       _GotoState57Action,
	{_State217, BinaryAddExprType}:                 _GotoState64Action,
	{_State217, CmpExprType}:                       _GotoState74Action,
	{_State217, BinaryCmpExprType}:                 _GotoState66Action,
	{_State217, AndExprType}:                       _GotoState58Action,
	{_State217, BinaryAndExprType}:                 _GotoState65Action,
	{_State217, OrExprType}:                        _GotoState94Action,
	{_State217, BinaryOrExprType}:                  _GotoState69Action,
	{_State217, InitializableTypeExprType}:         _GotoState84Action,
	{_State217, SliceTypeExprType}:                 _GotoState103Action,
	{_State217, ArrayTypeExprType}:                 _GotoState60Action,
	{_State217, MapTypeExprType}:                   _GotoState90Action,
	{_State217, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State217, AnonymousFuncExprType}:             _GotoState59Action,
	{_State218, IntegerLiteralToken}:               _GotoState40Action,
	{_State218, FloatLiteralToken}:                 _GotoState35Action,
	{_State218, RuneLiteralToken}:                  _GotoState48Action,
	{_State218, StringLiteralToken}:                _GotoState49Action,
	{_State218, IdentifierToken}:                   _GotoState38Action,
	{_State218, UnderscoreToken}:                   _GotoState53Action,
	{_State218, TrueToken}:                         _GotoState52Action,
	{_State218, FalseToken}:                        _GotoState34Action,
	{_State218, StructToken}:                       _GotoState50Action,
	{_State218, FuncToken}:                         _GotoState36Action,
	{_State218, VarToken}:                          _GotoState15Action,
	{_State218, LetToken}:                          _GotoState12Action,
	{_State218, NotToken}:                          _GotoState45Action,
	{_State218, LabelDeclToken}:                    _GotoState41Action,
	{_State218, LparenToken}:                       _GotoState43Action,
	{_State218, LbracketToken}:                     _GotoState42Action,
	{_State218, SubToken}:                          _GotoState51Action,
	{_State218, MulToken}:                          _GotoState44Action,
	{_State218, BitNegToken}:                       _GotoState27Action,
	{_State218, BitAndToken}:                       _GotoState26Action,
	{_State218, GreaterToken}:                      _GotoState37Action,
	{_State218, ParseErrorToken}:                   _GotoState46Action,
	{_State218, ExprsType}:                         _GotoState320Action,
	{_State218, VarDeclPatternType}:                _GotoState107Action,
	{_State218, VarOrLetType}:                      _GotoState24Action,
	{_State218, ExprType}:                          _GotoState76Action,
	{_State218, OptionalLabelDeclType}:             _GotoState93Action,
	{_State218, SequenceExprType}:                  _GotoState109Action,
	{_State218, IfExprType}:                        _GotoState80Action,
	{_State218, SwitchExprType}:                    _GotoState104Action,
	{_State218, LoopExprType}:                      _GotoState89Action,
	{_State218, CallExprType}:                      _GotoState71Action,
	{_State218, AtomExprType}:                      _GotoState63Action,
	{_State218, ParseErrorExprType}:                _GotoState95Action,
	{_State218, LiteralExprType}:                   _GotoState88Action,
	{_State218, NamedExprType}:                     _GotoState92Action,
	{_State218, BlockExprType}:                     _GotoState70Action,
	{_State218, InitializeExprType}:                _GotoState85Action,
	{_State218, ImplicitStructExprType}:            _GotoState81Action,
	{_State218, AccessibleExprType}:                _GotoState108Action,
	{_State218, AccessExprType}:                    _GotoState55Action,
	{_State218, IndexExprType}:                     _GotoState83Action,
	{_State218, PostfixableExprType}:               _GotoState97Action,
	{_State218, PostfixUnaryExprType}:              _GotoState96Action,
	{_State218, PrefixableExprType}:                _GotoState100Action,
	{_State218, PrefixUnaryExprType}:               _GotoState98Action,
	{_State218, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State218, MulExprType}:                       _GotoState91Action,
	{_State218, BinaryMulExprType}:                 _GotoState67Action,
	{_State218, AddExprType}:                       _GotoState57Action,
	{_State218, BinaryAddExprType}:                 _GotoState64Action,
	{_State218, CmpExprType}:                       _GotoState74Action,
	{_State218, BinaryCmpExprType}:                 _GotoState66Action,
	{_State218, AndExprType}:                       _GotoState58Action,
	{_State218, BinaryAndExprType}:                 _GotoState65Action,
	{_State218, OrExprType}:                        _GotoState94Action,
	{_State218, BinaryOrExprType}:                  _GotoState69Action,
	{_State218, InitializableTypeExprType}:         _GotoState84Action,
	{_State218, SliceTypeExprType}:                 _GotoState103Action,
	{_State218, ArrayTypeExprType}:                 _GotoState60Action,
	{_State218, MapTypeExprType}:                   _GotoState90Action,
	{_State218, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State218, AnonymousFuncExprType}:             _GotoState59Action,
	{_State219, CommaToken}:                        _GotoState216Action,
	{_State226, IntegerLiteralToken}:               _GotoState40Action,
	{_State226, FloatLiteralToken}:                 _GotoState35Action,
	{_State226, RuneLiteralToken}:                  _GotoState48Action,
	{_State226, StringLiteralToken}:                _GotoState49Action,
	{_State226, IdentifierToken}:                   _GotoState38Action,
	{_State226, UnderscoreToken}:                   _GotoState53Action,
	{_State226, TrueToken}:                         _GotoState52Action,
	{_State226, FalseToken}:                        _GotoState34Action,
	{_State226, StructToken}:                       _GotoState50Action,
	{_State226, FuncToken}:                         _GotoState36Action,
	{_State226, NotToken}:                          _GotoState45Action,
	{_State226, LabelDeclToken}:                    _GotoState41Action,
	{_State226, LparenToken}:                       _GotoState43Action,
	{_State226, LbracketToken}:                     _GotoState42Action,
	{_State226, SubToken}:                          _GotoState51Action,
	{_State226, MulToken}:                          _GotoState44Action,
	{_State226, BitNegToken}:                       _GotoState27Action,
	{_State226, BitAndToken}:                       _GotoState26Action,
	{_State226, ParseErrorToken}:                   _GotoState46Action,
	{_State226, OptionalLabelDeclType}:             _GotoState156Action,
	{_State226, CallExprType}:                      _GotoState71Action,
	{_State226, AtomExprType}:                      _GotoState63Action,
	{_State226, ParseErrorExprType}:                _GotoState95Action,
	{_State226, LiteralExprType}:                   _GotoState88Action,
	{_State226, NamedExprType}:                     _GotoState92Action,
	{_State226, BlockExprType}:                     _GotoState70Action,
	{_State226, InitializeExprType}:                _GotoState85Action,
	{_State226, ImplicitStructExprType}:            _GotoState81Action,
	{_State226, AccessibleExprType}:                _GotoState108Action,
	{_State226, AccessExprType}:                    _GotoState55Action,
	{_State226, IndexExprType}:                     _GotoState83Action,
	{_State226, PostfixableExprType}:               _GotoState97Action,
	{_State226, PostfixUnaryExprType}:              _GotoState96Action,
	{_State226, PrefixableExprType}:                _GotoState321Action,
	{_State226, PrefixUnaryExprType}:               _GotoState98Action,
	{_State226, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State226, InitializableTypeExprType}:         _GotoState84Action,
	{_State226, SliceTypeExprType}:                 _GotoState103Action,
	{_State226, ArrayTypeExprType}:                 _GotoState60Action,
	{_State226, MapTypeExprType}:                   _GotoState90Action,
	{_State226, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State226, AnonymousFuncExprType}:             _GotoState59Action,
	{_State227, LbraceToken}:                       _GotoState11Action,
	{_State227, StatementBlockType}:                _GotoState322Action,
	{_State228, IntegerLiteralToken}:               _GotoState40Action,
	{_State228, FloatLiteralToken}:                 _GotoState35Action,
	{_State228, RuneLiteralToken}:                  _GotoState48Action,
	{_State228, StringLiteralToken}:                _GotoState49Action,
	{_State228, IdentifierToken}:                   _GotoState38Action,
	{_State228, UnderscoreToken}:                   _GotoState53Action,
	{_State228, TrueToken}:                         _GotoState52Action,
	{_State228, FalseToken}:                        _GotoState34Action,
	{_State228, StructToken}:                       _GotoState50Action,
	{_State228, FuncToken}:                         _GotoState36Action,
	{_State228, VarToken}:                          _GotoState15Action,
	{_State228, LetToken}:                          _GotoState12Action,
	{_State228, NotToken}:                          _GotoState45Action,
	{_State228, LabelDeclToken}:                    _GotoState41Action,
	{_State228, LparenToken}:                       _GotoState43Action,
	{_State228, LbracketToken}:                     _GotoState42Action,
	{_State228, SubToken}:                          _GotoState51Action,
	{_State228, MulToken}:                          _GotoState44Action,
	{_State228, BitNegToken}:                       _GotoState27Action,
	{_State228, BitAndToken}:                       _GotoState26Action,
	{_State228, GreaterToken}:                      _GotoState37Action,
	{_State228, ParseErrorToken}:                   _GotoState46Action,
	{_State228, VarDeclPatternType}:                _GotoState107Action,
	{_State228, VarOrLetType}:                      _GotoState24Action,
	{_State228, AssignPatternType}:                 _GotoState323Action,
	{_State228, OptionalLabelDeclType}:             _GotoState156Action,
	{_State228, SequenceExprType}:                  _GotoState325Action,
	{_State228, ForAssignmentType}:                 _GotoState324Action,
	{_State228, CallExprType}:                      _GotoState71Action,
	{_State228, AtomExprType}:                      _GotoState63Action,
	{_State228, ParseErrorExprType}:                _GotoState95Action,
	{_State228, LiteralExprType}:                   _GotoState88Action,
	{_State228, NamedExprType}:                     _GotoState92Action,
	{_State228, BlockExprType}:                     _GotoState70Action,
	{_State228, InitializeExprType}:                _GotoState85Action,
	{_State228, ImplicitStructExprType}:            _GotoState81Action,
	{_State228, AccessibleExprType}:                _GotoState108Action,
	{_State228, AccessExprType}:                    _GotoState55Action,
	{_State228, IndexExprType}:                     _GotoState83Action,
	{_State228, PostfixableExprType}:               _GotoState97Action,
	{_State228, PostfixUnaryExprType}:              _GotoState96Action,
	{_State228, PrefixableExprType}:                _GotoState100Action,
	{_State228, PrefixUnaryExprType}:               _GotoState98Action,
	{_State228, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State228, MulExprType}:                       _GotoState91Action,
	{_State228, BinaryMulExprType}:                 _GotoState67Action,
	{_State228, AddExprType}:                       _GotoState57Action,
	{_State228, BinaryAddExprType}:                 _GotoState64Action,
	{_State228, CmpExprType}:                       _GotoState74Action,
	{_State228, BinaryCmpExprType}:                 _GotoState66Action,
	{_State228, AndExprType}:                       _GotoState58Action,
	{_State228, BinaryAndExprType}:                 _GotoState65Action,
	{_State228, OrExprType}:                        _GotoState94Action,
	{_State228, BinaryOrExprType}:                  _GotoState69Action,
	{_State228, InitializableTypeExprType}:         _GotoState84Action,
	{_State228, SliceTypeExprType}:                 _GotoState103Action,
	{_State228, ArrayTypeExprType}:                 _GotoState60Action,
	{_State228, MapTypeExprType}:                   _GotoState90Action,
	{_State228, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State228, AnonymousFuncExprType}:             _GotoState59Action,
	{_State229, IntegerLiteralToken}:               _GotoState40Action,
	{_State229, FloatLiteralToken}:                 _GotoState35Action,
	{_State229, RuneLiteralToken}:                  _GotoState48Action,
	{_State229, StringLiteralToken}:                _GotoState49Action,
	{_State229, IdentifierToken}:                   _GotoState38Action,
	{_State229, UnderscoreToken}:                   _GotoState53Action,
	{_State229, TrueToken}:                         _GotoState52Action,
	{_State229, FalseToken}:                        _GotoState34Action,
	{_State229, CaseToken}:                         _GotoState326Action,
	{_State229, StructToken}:                       _GotoState50Action,
	{_State229, FuncToken}:                         _GotoState36Action,
	{_State229, VarToken}:                          _GotoState15Action,
	{_State229, LetToken}:                          _GotoState12Action,
	{_State229, NotToken}:                          _GotoState45Action,
	{_State229, LabelDeclToken}:                    _GotoState41Action,
	{_State229, LparenToken}:                       _GotoState43Action,
	{_State229, LbracketToken}:                     _GotoState42Action,
	{_State229, SubToken}:                          _GotoState51Action,
	{_State229, MulToken}:                          _GotoState44Action,
	{_State229, BitNegToken}:                       _GotoState27Action,
	{_State229, BitAndToken}:                       _GotoState26Action,
	{_State229, GreaterToken}:                      _GotoState37Action,
	{_State229, ParseErrorToken}:                   _GotoState46Action,
	{_State229, VarDeclPatternType}:                _GotoState107Action,
	{_State229, VarOrLetType}:                      _GotoState24Action,
	{_State229, OptionalLabelDeclType}:             _GotoState156Action,
	{_State229, SequenceExprType}:                  _GotoState328Action,
	{_State229, ConditionType}:                     _GotoState327Action,
	{_State229, CallExprType}:                      _GotoState71Action,
	{_State229, AtomExprType}:                      _GotoState63Action,
	{_State229, ParseErrorExprType}:                _GotoState95Action,
	{_State229, LiteralExprType}:                   _GotoState88Action,
	{_State229, NamedExprType}:                     _GotoState92Action,
	{_State229, BlockExprType}:                     _GotoState70Action,
	{_State229, InitializeExprType}:                _GotoState85Action,
	{_State229, ImplicitStructExprType}:            _GotoState81Action,
	{_State229, AccessibleExprType}:                _GotoState108Action,
	{_State229, AccessExprType}:                    _GotoState55Action,
	{_State229, IndexExprType}:                     _GotoState83Action,
	{_State229, PostfixableExprType}:               _GotoState97Action,
	{_State229, PostfixUnaryExprType}:              _GotoState96Action,
	{_State229, PrefixableExprType}:                _GotoState100Action,
	{_State229, PrefixUnaryExprType}:               _GotoState98Action,
	{_State229, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State229, MulExprType}:                       _GotoState91Action,
	{_State229, BinaryMulExprType}:                 _GotoState67Action,
	{_State229, AddExprType}:                       _GotoState57Action,
	{_State229, BinaryAddExprType}:                 _GotoState64Action,
	{_State229, CmpExprType}:                       _GotoState74Action,
	{_State229, BinaryCmpExprType}:                 _GotoState66Action,
	{_State229, AndExprType}:                       _GotoState58Action,
	{_State229, BinaryAndExprType}:                 _GotoState65Action,
	{_State229, OrExprType}:                        _GotoState94Action,
	{_State229, BinaryOrExprType}:                  _GotoState69Action,
	{_State229, InitializableTypeExprType}:         _GotoState84Action,
	{_State229, SliceTypeExprType}:                 _GotoState103Action,
	{_State229, ArrayTypeExprType}:                 _GotoState60Action,
	{_State229, MapTypeExprType}:                   _GotoState90Action,
	{_State229, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State229, AnonymousFuncExprType}:             _GotoState59Action,
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
	{_State230, VarToken}:                          _GotoState15Action,
	{_State230, LetToken}:                          _GotoState12Action,
	{_State230, NotToken}:                          _GotoState45Action,
	{_State230, LabelDeclToken}:                    _GotoState41Action,
	{_State230, LparenToken}:                       _GotoState43Action,
	{_State230, LbracketToken}:                     _GotoState42Action,
	{_State230, SubToken}:                          _GotoState51Action,
	{_State230, MulToken}:                          _GotoState44Action,
	{_State230, BitNegToken}:                       _GotoState27Action,
	{_State230, BitAndToken}:                       _GotoState26Action,
	{_State230, GreaterToken}:                      _GotoState37Action,
	{_State230, ParseErrorToken}:                   _GotoState46Action,
	{_State230, VarDeclPatternType}:                _GotoState107Action,
	{_State230, VarOrLetType}:                      _GotoState24Action,
	{_State230, OptionalLabelDeclType}:             _GotoState156Action,
	{_State230, SequenceExprType}:                  _GotoState329Action,
	{_State230, CallExprType}:                      _GotoState71Action,
	{_State230, AtomExprType}:                      _GotoState63Action,
	{_State230, ParseErrorExprType}:                _GotoState95Action,
	{_State230, LiteralExprType}:                   _GotoState88Action,
	{_State230, NamedExprType}:                     _GotoState92Action,
	{_State230, BlockExprType}:                     _GotoState70Action,
	{_State230, InitializeExprType}:                _GotoState85Action,
	{_State230, ImplicitStructExprType}:            _GotoState81Action,
	{_State230, AccessibleExprType}:                _GotoState108Action,
	{_State230, AccessExprType}:                    _GotoState55Action,
	{_State230, IndexExprType}:                     _GotoState83Action,
	{_State230, PostfixableExprType}:               _GotoState97Action,
	{_State230, PostfixUnaryExprType}:              _GotoState96Action,
	{_State230, PrefixableExprType}:                _GotoState100Action,
	{_State230, PrefixUnaryExprType}:               _GotoState98Action,
	{_State230, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State230, MulExprType}:                       _GotoState91Action,
	{_State230, BinaryMulExprType}:                 _GotoState67Action,
	{_State230, AddExprType}:                       _GotoState57Action,
	{_State230, BinaryAddExprType}:                 _GotoState64Action,
	{_State230, CmpExprType}:                       _GotoState74Action,
	{_State230, BinaryCmpExprType}:                 _GotoState66Action,
	{_State230, AndExprType}:                       _GotoState58Action,
	{_State230, BinaryAndExprType}:                 _GotoState65Action,
	{_State230, OrExprType}:                        _GotoState94Action,
	{_State230, BinaryOrExprType}:                  _GotoState69Action,
	{_State230, InitializableTypeExprType}:         _GotoState84Action,
	{_State230, SliceTypeExprType}:                 _GotoState103Action,
	{_State230, ArrayTypeExprType}:                 _GotoState60Action,
	{_State230, MapTypeExprType}:                   _GotoState90Action,
	{_State230, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State230, AnonymousFuncExprType}:             _GotoState59Action,
	{_State234, IntegerLiteralToken}:               _GotoState40Action,
	{_State234, FloatLiteralToken}:                 _GotoState35Action,
	{_State234, RuneLiteralToken}:                  _GotoState48Action,
	{_State234, StringLiteralToken}:                _GotoState49Action,
	{_State234, IdentifierToken}:                   _GotoState38Action,
	{_State234, UnderscoreToken}:                   _GotoState53Action,
	{_State234, TrueToken}:                         _GotoState52Action,
	{_State234, FalseToken}:                        _GotoState34Action,
	{_State234, StructToken}:                       _GotoState50Action,
	{_State234, FuncToken}:                         _GotoState36Action,
	{_State234, NotToken}:                          _GotoState45Action,
	{_State234, LabelDeclToken}:                    _GotoState41Action,
	{_State234, LparenToken}:                       _GotoState43Action,
	{_State234, LbracketToken}:                     _GotoState42Action,
	{_State234, SubToken}:                          _GotoState51Action,
	{_State234, MulToken}:                          _GotoState44Action,
	{_State234, BitNegToken}:                       _GotoState27Action,
	{_State234, BitAndToken}:                       _GotoState26Action,
	{_State234, ParseErrorToken}:                   _GotoState46Action,
	{_State234, OptionalLabelDeclType}:             _GotoState156Action,
	{_State234, CallExprType}:                      _GotoState71Action,
	{_State234, AtomExprType}:                      _GotoState63Action,
	{_State234, ParseErrorExprType}:                _GotoState95Action,
	{_State234, LiteralExprType}:                   _GotoState88Action,
	{_State234, NamedExprType}:                     _GotoState92Action,
	{_State234, BlockExprType}:                     _GotoState70Action,
	{_State234, InitializeExprType}:                _GotoState85Action,
	{_State234, ImplicitStructExprType}:            _GotoState81Action,
	{_State234, AccessibleExprType}:                _GotoState108Action,
	{_State234, AccessExprType}:                    _GotoState55Action,
	{_State234, IndexExprType}:                     _GotoState83Action,
	{_State234, PostfixableExprType}:               _GotoState97Action,
	{_State234, PostfixUnaryExprType}:              _GotoState96Action,
	{_State234, PrefixableExprType}:                _GotoState100Action,
	{_State234, PrefixUnaryExprType}:               _GotoState98Action,
	{_State234, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State234, MulExprType}:                       _GotoState91Action,
	{_State234, BinaryMulExprType}:                 _GotoState67Action,
	{_State234, AddExprType}:                       _GotoState57Action,
	{_State234, BinaryAddExprType}:                 _GotoState64Action,
	{_State234, CmpExprType}:                       _GotoState74Action,
	{_State234, BinaryCmpExprType}:                 _GotoState66Action,
	{_State234, AndExprType}:                       _GotoState330Action,
	{_State234, BinaryAndExprType}:                 _GotoState65Action,
	{_State234, InitializableTypeExprType}:         _GotoState84Action,
	{_State234, SliceTypeExprType}:                 _GotoState103Action,
	{_State234, ArrayTypeExprType}:                 _GotoState60Action,
	{_State234, MapTypeExprType}:                   _GotoState90Action,
	{_State234, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State234, AnonymousFuncExprType}:             _GotoState59Action,
	{_State236, IdentifierToken}:                   _GotoState241Action,
	{_State236, UnderscoreToken}:                   _GotoState242Action,
	{_State236, UnsafeToken}:                       _GotoState54Action,
	{_State236, StructToken}:                       _GotoState50Action,
	{_State236, EnumToken}:                         _GotoState113Action,
	{_State236, TraitToken}:                        _GotoState121Action,
	{_State236, FuncToken}:                         _GotoState240Action,
	{_State236, LparenToken}:                       _GotoState117Action,
	{_State236, LbracketToken}:                     _GotoState42Action,
	{_State236, DotToken}:                          _GotoState112Action,
	{_State236, QuestionToken}:                     _GotoState119Action,
	{_State236, ExclaimToken}:                      _GotoState114Action,
	{_State236, TildeTildeToken}:                   _GotoState120Action,
	{_State236, BitNegToken}:                       _GotoState111Action,
	{_State236, BitAndToken}:                       _GotoState110Action,
	{_State236, ParseErrorToken}:                   _GotoState118Action,
	{_State236, UnsafeStatementType}:               _GotoState250Action,
	{_State236, InitializableTypeExprType}:         _GotoState130Action,
	{_State236, SliceTypeExprType}:                 _GotoState103Action,
	{_State236, ArrayTypeExprType}:                 _GotoState60Action,
	{_State236, MapTypeExprType}:                   _GotoState90Action,
	{_State236, AtomTypeExprType}:                  _GotoState123Action,
	{_State236, NamedTypeExprType}:                 _GotoState131Action,
	{_State236, InferredTypeExprType}:              _GotoState129Action,
	{_State236, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State236, ReturnableTypeExprType}:            _GotoState135Action,
	{_State236, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State236, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State236, TypeExprType}:                      _GotoState249Action,
	{_State236, BinaryTypeExprType}:                _GotoState124Action,
	{_State236, FieldDefType}:                      _GotoState332Action,
	{_State236, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State236, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State236, TraitTypeExprType}:                 _GotoState136Action,
	{_State236, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State236, ProperExplicitEnumValueDefsType}:   _GotoState333Action,
	{_State236, ExplicitEnumValueDefsType}:         _GotoState331Action,
	{_State236, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State236, FuncTypeExprType}:                  _GotoState126Action,
	{_State236, MethodSignatureType}:               _GotoState246Action,
	{_State237, IdentifierToken}:                   _GotoState335Action,
	{_State237, UnderscoreToken}:                   _GotoState336Action,
	{_State237, StructToken}:                       _GotoState50Action,
	{_State237, EnumToken}:                         _GotoState113Action,
	{_State237, TraitToken}:                        _GotoState121Action,
	{_State237, FuncToken}:                         _GotoState115Action,
	{_State237, LparenToken}:                       _GotoState117Action,
	{_State237, LbracketToken}:                     _GotoState42Action,
	{_State237, DotToken}:                          _GotoState112Action,
	{_State237, QuestionToken}:                     _GotoState119Action,
	{_State237, ExclaimToken}:                      _GotoState114Action,
	{_State237, EllipsisToken}:                     _GotoState334Action,
	{_State237, TildeTildeToken}:                   _GotoState120Action,
	{_State237, BitNegToken}:                       _GotoState111Action,
	{_State237, BitAndToken}:                       _GotoState110Action,
	{_State237, ParseErrorToken}:                   _GotoState118Action,
	{_State237, InitializableTypeExprType}:         _GotoState130Action,
	{_State237, SliceTypeExprType}:                 _GotoState103Action,
	{_State237, ArrayTypeExprType}:                 _GotoState60Action,
	{_State237, MapTypeExprType}:                   _GotoState90Action,
	{_State237, AtomTypeExprType}:                  _GotoState123Action,
	{_State237, NamedTypeExprType}:                 _GotoState131Action,
	{_State237, InferredTypeExprType}:              _GotoState129Action,
	{_State237, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State237, ReturnableTypeExprType}:            _GotoState135Action,
	{_State237, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State237, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State237, TypeExprType}:                      _GotoState341Action,
	{_State237, BinaryTypeExprType}:                _GotoState124Action,
	{_State237, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State237, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State237, TraitTypeExprType}:                 _GotoState136Action,
	{_State237, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State237, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State237, ProperParameterDefType}:            _GotoState340Action,
	{_State237, ParameterDeclType}:                 _GotoState337Action,
	{_State237, ProperParameterDeclsType}:          _GotoState339Action,
	{_State237, ParameterDeclsType}:                _GotoState338Action,
	{_State237, FuncTypeExprType}:                  _GotoState126Action,
	{_State238, IdentifierToken}:                   _GotoState342Action,
	{_State240, IdentifierToken}:                   _GotoState343Action,
	{_State240, LparenToken}:                       _GotoState237Action,
	{_State241, IdentifierToken}:                   _GotoState116Action,
	{_State241, UnderscoreToken}:                   _GotoState122Action,
	{_State241, StructToken}:                       _GotoState50Action,
	{_State241, EnumToken}:                         _GotoState113Action,
	{_State241, TraitToken}:                        _GotoState121Action,
	{_State241, FuncToken}:                         _GotoState115Action,
	{_State241, LparenToken}:                       _GotoState117Action,
	{_State241, LbracketToken}:                     _GotoState42Action,
	{_State241, DotToken}:                          _GotoState344Action,
	{_State241, QuestionToken}:                     _GotoState119Action,
	{_State241, ExclaimToken}:                      _GotoState114Action,
	{_State241, DollarLbracketToken}:               _GotoState187Action,
	{_State241, TildeTildeToken}:                   _GotoState120Action,
	{_State241, BitNegToken}:                       _GotoState111Action,
	{_State241, BitAndToken}:                       _GotoState110Action,
	{_State241, ParseErrorToken}:                   _GotoState118Action,
	{_State241, InitializableTypeExprType}:         _GotoState130Action,
	{_State241, SliceTypeExprType}:                 _GotoState103Action,
	{_State241, ArrayTypeExprType}:                 _GotoState60Action,
	{_State241, MapTypeExprType}:                   _GotoState90Action,
	{_State241, AtomTypeExprType}:                  _GotoState123Action,
	{_State241, NamedTypeExprType}:                 _GotoState131Action,
	{_State241, InferredTypeExprType}:              _GotoState129Action,
	{_State241, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State241, ReturnableTypeExprType}:            _GotoState135Action,
	{_State241, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State241, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State241, TypeExprType}:                      _GotoState345Action,
	{_State241, BinaryTypeExprType}:                _GotoState124Action,
	{_State241, GenericTypeArgumentsType}:          _GotoState239Action,
	{_State241, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State241, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State241, TraitTypeExprType}:                 _GotoState136Action,
	{_State241, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State241, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State241, FuncTypeExprType}:                  _GotoState126Action,
	{_State242, IdentifierToken}:                   _GotoState116Action,
	{_State242, UnderscoreToken}:                   _GotoState122Action,
	{_State242, StructToken}:                       _GotoState50Action,
	{_State242, EnumToken}:                         _GotoState113Action,
	{_State242, TraitToken}:                        _GotoState121Action,
	{_State242, FuncToken}:                         _GotoState115Action,
	{_State242, LparenToken}:                       _GotoState117Action,
	{_State242, LbracketToken}:                     _GotoState42Action,
	{_State242, DotToken}:                          _GotoState112Action,
	{_State242, QuestionToken}:                     _GotoState119Action,
	{_State242, ExclaimToken}:                      _GotoState114Action,
	{_State242, TildeTildeToken}:                   _GotoState120Action,
	{_State242, BitNegToken}:                       _GotoState111Action,
	{_State242, BitAndToken}:                       _GotoState110Action,
	{_State242, ParseErrorToken}:                   _GotoState118Action,
	{_State242, InitializableTypeExprType}:         _GotoState130Action,
	{_State242, SliceTypeExprType}:                 _GotoState103Action,
	{_State242, ArrayTypeExprType}:                 _GotoState60Action,
	{_State242, MapTypeExprType}:                   _GotoState90Action,
	{_State242, AtomTypeExprType}:                  _GotoState123Action,
	{_State242, NamedTypeExprType}:                 _GotoState131Action,
	{_State242, InferredTypeExprType}:              _GotoState129Action,
	{_State242, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State242, ReturnableTypeExprType}:            _GotoState135Action,
	{_State242, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State242, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State242, TypeExprType}:                      _GotoState346Action,
	{_State242, BinaryTypeExprType}:                _GotoState124Action,
	{_State242, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State242, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State242, TraitTypeExprType}:                 _GotoState136Action,
	{_State242, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State242, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State242, FuncTypeExprType}:                  _GotoState126Action,
	{_State243, OrToken}:                           _GotoState347Action,
	{_State244, RparenToken}:                       _GotoState348Action,
	{_State245, RparenToken}:                       _GotoState349Action,
	{_State247, NewlinesToken}:                     _GotoState350Action,
	{_State247, OrToken}:                           _GotoState351Action,
	{_State248, CommaToken}:                        _GotoState352Action,
	{_State249, AssignToken}:                       _GotoState353Action,
	{_State249, AddToken}:                          _GotoState253Action,
	{_State249, SubToken}:                          _GotoState255Action,
	{_State249, MulToken}:                          _GotoState254Action,
	{_State249, BinaryTypeOpType}:                  _GotoState256Action,
	{_State249, OptionalDefaultType}:               _GotoState354Action,
	{_State251, IdentifierToken}:                   _GotoState241Action,
	{_State251, UnderscoreToken}:                   _GotoState242Action,
	{_State251, UnsafeToken}:                       _GotoState54Action,
	{_State251, StructToken}:                       _GotoState50Action,
	{_State251, EnumToken}:                         _GotoState113Action,
	{_State251, TraitToken}:                        _GotoState121Action,
	{_State251, FuncToken}:                         _GotoState240Action,
	{_State251, LparenToken}:                       _GotoState117Action,
	{_State251, LbracketToken}:                     _GotoState42Action,
	{_State251, DotToken}:                          _GotoState112Action,
	{_State251, QuestionToken}:                     _GotoState119Action,
	{_State251, ExclaimToken}:                      _GotoState114Action,
	{_State251, TildeTildeToken}:                   _GotoState120Action,
	{_State251, BitNegToken}:                       _GotoState111Action,
	{_State251, BitAndToken}:                       _GotoState110Action,
	{_State251, ParseErrorToken}:                   _GotoState118Action,
	{_State251, UnsafeStatementType}:               _GotoState250Action,
	{_State251, InitializableTypeExprType}:         _GotoState130Action,
	{_State251, SliceTypeExprType}:                 _GotoState103Action,
	{_State251, ArrayTypeExprType}:                 _GotoState60Action,
	{_State251, MapTypeExprType}:                   _GotoState90Action,
	{_State251, AtomTypeExprType}:                  _GotoState123Action,
	{_State251, NamedTypeExprType}:                 _GotoState131Action,
	{_State251, InferredTypeExprType}:              _GotoState129Action,
	{_State251, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State251, ReturnableTypeExprType}:            _GotoState135Action,
	{_State251, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State251, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State251, TypeExprType}:                      _GotoState249Action,
	{_State251, BinaryTypeExprType}:                _GotoState124Action,
	{_State251, FieldDefType}:                      _GotoState304Action,
	{_State251, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State251, ProperExplicitFieldDefsType}:       _GotoState305Action,
	{_State251, ExplicitFieldDefsType}:             _GotoState355Action,
	{_State251, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State251, TraitTypeExprType}:                 _GotoState136Action,
	{_State251, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State251, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State251, FuncTypeExprType}:                  _GotoState126Action,
	{_State251, MethodSignatureType}:               _GotoState246Action,
	{_State256, IdentifierToken}:                   _GotoState116Action,
	{_State256, UnderscoreToken}:                   _GotoState122Action,
	{_State256, StructToken}:                       _GotoState50Action,
	{_State256, EnumToken}:                         _GotoState113Action,
	{_State256, TraitToken}:                        _GotoState121Action,
	{_State256, FuncToken}:                         _GotoState115Action,
	{_State256, LparenToken}:                       _GotoState117Action,
	{_State256, LbracketToken}:                     _GotoState42Action,
	{_State256, DotToken}:                          _GotoState112Action,
	{_State256, QuestionToken}:                     _GotoState119Action,
	{_State256, ExclaimToken}:                      _GotoState114Action,
	{_State256, TildeTildeToken}:                   _GotoState120Action,
	{_State256, BitNegToken}:                       _GotoState111Action,
	{_State256, BitAndToken}:                       _GotoState110Action,
	{_State256, ParseErrorToken}:                   _GotoState118Action,
	{_State256, InitializableTypeExprType}:         _GotoState130Action,
	{_State256, SliceTypeExprType}:                 _GotoState103Action,
	{_State256, ArrayTypeExprType}:                 _GotoState60Action,
	{_State256, MapTypeExprType}:                   _GotoState90Action,
	{_State256, AtomTypeExprType}:                  _GotoState123Action,
	{_State256, NamedTypeExprType}:                 _GotoState131Action,
	{_State256, InferredTypeExprType}:              _GotoState129Action,
	{_State256, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State256, ReturnableTypeExprType}:            _GotoState356Action,
	{_State256, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State256, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State256, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State256, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State256, TraitTypeExprType}:                 _GotoState136Action,
	{_State256, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State256, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State256, FuncTypeExprType}:                  _GotoState126Action,
	{_State257, IntegerLiteralToken}:               _GotoState40Action,
	{_State257, FloatLiteralToken}:                 _GotoState35Action,
	{_State257, RuneLiteralToken}:                  _GotoState48Action,
	{_State257, StringLiteralToken}:                _GotoState49Action,
	{_State257, IdentifierToken}:                   _GotoState38Action,
	{_State257, UnderscoreToken}:                   _GotoState53Action,
	{_State257, TrueToken}:                         _GotoState52Action,
	{_State257, FalseToken}:                        _GotoState34Action,
	{_State257, StructToken}:                       _GotoState50Action,
	{_State257, FuncToken}:                         _GotoState36Action,
	{_State257, VarToken}:                          _GotoState15Action,
	{_State257, LetToken}:                          _GotoState12Action,
	{_State257, NotToken}:                          _GotoState45Action,
	{_State257, LabelDeclToken}:                    _GotoState41Action,
	{_State257, LparenToken}:                       _GotoState43Action,
	{_State257, LbracketToken}:                     _GotoState42Action,
	{_State257, SubToken}:                          _GotoState51Action,
	{_State257, MulToken}:                          _GotoState44Action,
	{_State257, BitNegToken}:                       _GotoState27Action,
	{_State257, BitAndToken}:                       _GotoState26Action,
	{_State257, GreaterToken}:                      _GotoState37Action,
	{_State257, ParseErrorToken}:                   _GotoState46Action,
	{_State257, VarDeclPatternType}:                _GotoState107Action,
	{_State257, VarOrLetType}:                      _GotoState24Action,
	{_State257, ExprType}:                          _GotoState357Action,
	{_State257, OptionalLabelDeclType}:             _GotoState93Action,
	{_State257, SequenceExprType}:                  _GotoState109Action,
	{_State257, IfExprType}:                        _GotoState80Action,
	{_State257, SwitchExprType}:                    _GotoState104Action,
	{_State257, LoopExprType}:                      _GotoState89Action,
	{_State257, CallExprType}:                      _GotoState71Action,
	{_State257, AtomExprType}:                      _GotoState63Action,
	{_State257, ParseErrorExprType}:                _GotoState95Action,
	{_State257, LiteralExprType}:                   _GotoState88Action,
	{_State257, NamedExprType}:                     _GotoState92Action,
	{_State257, BlockExprType}:                     _GotoState70Action,
	{_State257, InitializeExprType}:                _GotoState85Action,
	{_State257, ImplicitStructExprType}:            _GotoState81Action,
	{_State257, AccessibleExprType}:                _GotoState108Action,
	{_State257, AccessExprType}:                    _GotoState55Action,
	{_State257, IndexExprType}:                     _GotoState83Action,
	{_State257, PostfixableExprType}:               _GotoState97Action,
	{_State257, PostfixUnaryExprType}:              _GotoState96Action,
	{_State257, PrefixableExprType}:                _GotoState100Action,
	{_State257, PrefixUnaryExprType}:               _GotoState98Action,
	{_State257, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State257, MulExprType}:                       _GotoState91Action,
	{_State257, BinaryMulExprType}:                 _GotoState67Action,
	{_State257, AddExprType}:                       _GotoState57Action,
	{_State257, BinaryAddExprType}:                 _GotoState64Action,
	{_State257, CmpExprType}:                       _GotoState74Action,
	{_State257, BinaryCmpExprType}:                 _GotoState66Action,
	{_State257, AndExprType}:                       _GotoState58Action,
	{_State257, BinaryAndExprType}:                 _GotoState65Action,
	{_State257, OrExprType}:                        _GotoState94Action,
	{_State257, BinaryOrExprType}:                  _GotoState69Action,
	{_State257, InitializableTypeExprType}:         _GotoState84Action,
	{_State257, SliceTypeExprType}:                 _GotoState103Action,
	{_State257, ArrayTypeExprType}:                 _GotoState60Action,
	{_State257, MapTypeExprType}:                   _GotoState90Action,
	{_State257, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State257, AnonymousFuncExprType}:             _GotoState59Action,
	{_State258, IdentifierToken}:                   _GotoState358Action,
	{_State258, GenericParameterDefType}:           _GotoState359Action,
	{_State258, ProperGenericParameterDefListType}: _GotoState361Action,
	{_State258, GenericParameterDefListType}:       _GotoState360Action,
	{_State259, LparenToken}:                       _GotoState362Action,
	{_State260, IdentifierToken}:                   _GotoState116Action,
	{_State260, UnderscoreToken}:                   _GotoState122Action,
	{_State260, StructToken}:                       _GotoState50Action,
	{_State260, EnumToken}:                         _GotoState113Action,
	{_State260, TraitToken}:                        _GotoState121Action,
	{_State260, FuncToken}:                         _GotoState115Action,
	{_State260, LparenToken}:                       _GotoState117Action,
	{_State260, LbracketToken}:                     _GotoState42Action,
	{_State260, DotToken}:                          _GotoState112Action,
	{_State260, QuestionToken}:                     _GotoState119Action,
	{_State260, ExclaimToken}:                      _GotoState114Action,
	{_State260, EllipsisToken}:                     _GotoState363Action,
	{_State260, TildeTildeToken}:                   _GotoState120Action,
	{_State260, BitNegToken}:                       _GotoState111Action,
	{_State260, BitAndToken}:                       _GotoState110Action,
	{_State260, ParseErrorToken}:                   _GotoState118Action,
	{_State260, InitializableTypeExprType}:         _GotoState130Action,
	{_State260, SliceTypeExprType}:                 _GotoState103Action,
	{_State260, ArrayTypeExprType}:                 _GotoState60Action,
	{_State260, MapTypeExprType}:                   _GotoState90Action,
	{_State260, AtomTypeExprType}:                  _GotoState123Action,
	{_State260, NamedTypeExprType}:                 _GotoState131Action,
	{_State260, InferredTypeExprType}:              _GotoState129Action,
	{_State260, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State260, ReturnableTypeExprType}:            _GotoState135Action,
	{_State260, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State260, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State260, TypeExprType}:                      _GotoState364Action,
	{_State260, BinaryTypeExprType}:                _GotoState124Action,
	{_State260, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State260, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State260, TraitTypeExprType}:                 _GotoState136Action,
	{_State260, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State260, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State260, FuncTypeExprType}:                  _GotoState126Action,
	{_State261, IdentifierToken}:                   _GotoState116Action,
	{_State261, UnderscoreToken}:                   _GotoState122Action,
	{_State261, StructToken}:                       _GotoState50Action,
	{_State261, EnumToken}:                         _GotoState113Action,
	{_State261, TraitToken}:                        _GotoState121Action,
	{_State261, FuncToken}:                         _GotoState115Action,
	{_State261, LparenToken}:                       _GotoState117Action,
	{_State261, LbracketToken}:                     _GotoState42Action,
	{_State261, DotToken}:                          _GotoState112Action,
	{_State261, QuestionToken}:                     _GotoState119Action,
	{_State261, ExclaimToken}:                      _GotoState114Action,
	{_State261, EllipsisToken}:                     _GotoState365Action,
	{_State261, TildeTildeToken}:                   _GotoState120Action,
	{_State261, BitNegToken}:                       _GotoState111Action,
	{_State261, BitAndToken}:                       _GotoState110Action,
	{_State261, ParseErrorToken}:                   _GotoState118Action,
	{_State261, InitializableTypeExprType}:         _GotoState130Action,
	{_State261, SliceTypeExprType}:                 _GotoState103Action,
	{_State261, ArrayTypeExprType}:                 _GotoState60Action,
	{_State261, MapTypeExprType}:                   _GotoState90Action,
	{_State261, AtomTypeExprType}:                  _GotoState123Action,
	{_State261, NamedTypeExprType}:                 _GotoState131Action,
	{_State261, InferredTypeExprType}:              _GotoState129Action,
	{_State261, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State261, ReturnableTypeExprType}:            _GotoState135Action,
	{_State261, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State261, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State261, TypeExprType}:                      _GotoState366Action,
	{_State261, BinaryTypeExprType}:                _GotoState124Action,
	{_State261, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State261, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State261, TraitTypeExprType}:                 _GotoState136Action,
	{_State261, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State261, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State261, FuncTypeExprType}:                  _GotoState126Action,
	{_State262, RparenToken}:                       _GotoState367Action,
	{_State264, IntegerLiteralToken}:               _GotoState40Action,
	{_State264, FloatLiteralToken}:                 _GotoState35Action,
	{_State264, RuneLiteralToken}:                  _GotoState48Action,
	{_State264, StringLiteralToken}:                _GotoState49Action,
	{_State264, IdentifierToken}:                   _GotoState38Action,
	{_State264, UnderscoreToken}:                   _GotoState53Action,
	{_State264, TrueToken}:                         _GotoState52Action,
	{_State264, FalseToken}:                        _GotoState34Action,
	{_State264, CaseToken}:                         _GotoState29Action,
	{_State264, DefaultToken}:                      _GotoState31Action,
	{_State264, ReturnToken}:                       _GotoState47Action,
	{_State264, BreakToken}:                        _GotoState28Action,
	{_State264, ContinueToken}:                     _GotoState30Action,
	{_State264, FallthroughToken}:                  _GotoState33Action,
	{_State264, ImportToken}:                       _GotoState39Action,
	{_State264, UnsafeToken}:                       _GotoState54Action,
	{_State264, StructToken}:                       _GotoState50Action,
	{_State264, FuncToken}:                         _GotoState36Action,
	{_State264, AsyncToken}:                        _GotoState25Action,
	{_State264, DeferToken}:                        _GotoState32Action,
	{_State264, VarToken}:                          _GotoState15Action,
	{_State264, LetToken}:                          _GotoState12Action,
	{_State264, NotToken}:                          _GotoState45Action,
	{_State264, LabelDeclToken}:                    _GotoState41Action,
	{_State264, LparenToken}:                       _GotoState43Action,
	{_State264, LbracketToken}:                     _GotoState42Action,
	{_State264, SubToken}:                          _GotoState51Action,
	{_State264, MulToken}:                          _GotoState44Action,
	{_State264, BitNegToken}:                       _GotoState27Action,
	{_State264, BitAndToken}:                       _GotoState26Action,
	{_State264, GreaterToken}:                      _GotoState37Action,
	{_State264, ParseErrorToken}:                   _GotoState46Action,
	{_State264, SimpleStatementType}:               _GotoState102Action,
	{_State264, StatementType}:                     _GotoState368Action,
	{_State264, ExprOrImproperStructStatementType}: _GotoState77Action,
	{_State264, ExprsType}:                         _GotoState78Action,
	{_State264, CallbackOpType}:                    _GotoState72Action,
	{_State264, CallbackOpStatementType}:           _GotoState73Action,
	{_State264, UnsafeStatementType}:               _GotoState106Action,
	{_State264, JumpStatementType}:                 _GotoState86Action,
	{_State264, JumpTypeType}:                      _GotoState87Action,
	{_State264, FallthroughStatementType}:          _GotoState79Action,
	{_State264, AssignStatementType}:               _GotoState62Action,
	{_State264, UnaryOpAssignStatementType}:        _GotoState105Action,
	{_State264, BinaryOpAssignStatementType}:       _GotoState68Action,
	{_State264, ImportStatementType}:               _GotoState82Action,
	{_State264, VarDeclPatternType}:                _GotoState107Action,
	{_State264, VarOrLetType}:                      _GotoState24Action,
	{_State264, AssignPatternType}:                 _GotoState61Action,
	{_State264, ExprType}:                          _GotoState76Action,
	{_State264, OptionalLabelDeclType}:             _GotoState93Action,
	{_State264, SequenceExprType}:                  _GotoState101Action,
	{_State264, IfExprType}:                        _GotoState80Action,
	{_State264, SwitchExprType}:                    _GotoState104Action,
	{_State264, LoopExprType}:                      _GotoState89Action,
	{_State264, CallExprType}:                      _GotoState71Action,
	{_State264, AtomExprType}:                      _GotoState63Action,
	{_State264, ParseErrorExprType}:                _GotoState95Action,
	{_State264, LiteralExprType}:                   _GotoState88Action,
	{_State264, NamedExprType}:                     _GotoState92Action,
	{_State264, BlockExprType}:                     _GotoState70Action,
	{_State264, InitializeExprType}:                _GotoState85Action,
	{_State264, ImplicitStructExprType}:            _GotoState81Action,
	{_State264, AccessibleExprType}:                _GotoState56Action,
	{_State264, AccessExprType}:                    _GotoState55Action,
	{_State264, IndexExprType}:                     _GotoState83Action,
	{_State264, PostfixableExprType}:               _GotoState97Action,
	{_State264, PostfixUnaryExprType}:              _GotoState96Action,
	{_State264, PrefixableExprType}:                _GotoState100Action,
	{_State264, PrefixUnaryExprType}:               _GotoState98Action,
	{_State264, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State264, MulExprType}:                       _GotoState91Action,
	{_State264, BinaryMulExprType}:                 _GotoState67Action,
	{_State264, AddExprType}:                       _GotoState57Action,
	{_State264, BinaryAddExprType}:                 _GotoState64Action,
	{_State264, CmpExprType}:                       _GotoState74Action,
	{_State264, BinaryCmpExprType}:                 _GotoState66Action,
	{_State264, AndExprType}:                       _GotoState58Action,
	{_State264, BinaryAndExprType}:                 _GotoState65Action,
	{_State264, OrExprType}:                        _GotoState94Action,
	{_State264, BinaryOrExprType}:                  _GotoState69Action,
	{_State264, InitializableTypeExprType}:         _GotoState84Action,
	{_State264, SliceTypeExprType}:                 _GotoState103Action,
	{_State264, ArrayTypeExprType}:                 _GotoState60Action,
	{_State264, MapTypeExprType}:                   _GotoState90Action,
	{_State264, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State264, AnonymousFuncExprType}:             _GotoState59Action,
	{_State265, IntegerLiteralToken}:               _GotoState40Action,
	{_State265, FloatLiteralToken}:                 _GotoState35Action,
	{_State265, RuneLiteralToken}:                  _GotoState48Action,
	{_State265, StringLiteralToken}:                _GotoState49Action,
	{_State265, IdentifierToken}:                   _GotoState38Action,
	{_State265, UnderscoreToken}:                   _GotoState53Action,
	{_State265, TrueToken}:                         _GotoState52Action,
	{_State265, FalseToken}:                        _GotoState34Action,
	{_State265, CaseToken}:                         _GotoState29Action,
	{_State265, DefaultToken}:                      _GotoState31Action,
	{_State265, ReturnToken}:                       _GotoState47Action,
	{_State265, BreakToken}:                        _GotoState28Action,
	{_State265, ContinueToken}:                     _GotoState30Action,
	{_State265, FallthroughToken}:                  _GotoState33Action,
	{_State265, ImportToken}:                       _GotoState39Action,
	{_State265, UnsafeToken}:                       _GotoState54Action,
	{_State265, StructToken}:                       _GotoState50Action,
	{_State265, FuncToken}:                         _GotoState36Action,
	{_State265, AsyncToken}:                        _GotoState25Action,
	{_State265, DeferToken}:                        _GotoState32Action,
	{_State265, VarToken}:                          _GotoState15Action,
	{_State265, LetToken}:                          _GotoState12Action,
	{_State265, NotToken}:                          _GotoState45Action,
	{_State265, LabelDeclToken}:                    _GotoState41Action,
	{_State265, LparenToken}:                       _GotoState43Action,
	{_State265, LbracketToken}:                     _GotoState42Action,
	{_State265, SubToken}:                          _GotoState51Action,
	{_State265, MulToken}:                          _GotoState44Action,
	{_State265, BitNegToken}:                       _GotoState27Action,
	{_State265, BitAndToken}:                       _GotoState26Action,
	{_State265, GreaterToken}:                      _GotoState37Action,
	{_State265, ParseErrorToken}:                   _GotoState46Action,
	{_State265, SimpleStatementType}:               _GotoState102Action,
	{_State265, StatementType}:                     _GotoState369Action,
	{_State265, ExprOrImproperStructStatementType}: _GotoState77Action,
	{_State265, ExprsType}:                         _GotoState78Action,
	{_State265, CallbackOpType}:                    _GotoState72Action,
	{_State265, CallbackOpStatementType}:           _GotoState73Action,
	{_State265, UnsafeStatementType}:               _GotoState106Action,
	{_State265, JumpStatementType}:                 _GotoState86Action,
	{_State265, JumpTypeType}:                      _GotoState87Action,
	{_State265, FallthroughStatementType}:          _GotoState79Action,
	{_State265, AssignStatementType}:               _GotoState62Action,
	{_State265, UnaryOpAssignStatementType}:        _GotoState105Action,
	{_State265, BinaryOpAssignStatementType}:       _GotoState68Action,
	{_State265, ImportStatementType}:               _GotoState82Action,
	{_State265, VarDeclPatternType}:                _GotoState107Action,
	{_State265, VarOrLetType}:                      _GotoState24Action,
	{_State265, AssignPatternType}:                 _GotoState61Action,
	{_State265, ExprType}:                          _GotoState76Action,
	{_State265, OptionalLabelDeclType}:             _GotoState93Action,
	{_State265, SequenceExprType}:                  _GotoState101Action,
	{_State265, IfExprType}:                        _GotoState80Action,
	{_State265, SwitchExprType}:                    _GotoState104Action,
	{_State265, LoopExprType}:                      _GotoState89Action,
	{_State265, CallExprType}:                      _GotoState71Action,
	{_State265, AtomExprType}:                      _GotoState63Action,
	{_State265, ParseErrorExprType}:                _GotoState95Action,
	{_State265, LiteralExprType}:                   _GotoState88Action,
	{_State265, NamedExprType}:                     _GotoState92Action,
	{_State265, BlockExprType}:                     _GotoState70Action,
	{_State265, InitializeExprType}:                _GotoState85Action,
	{_State265, ImplicitStructExprType}:            _GotoState81Action,
	{_State265, AccessibleExprType}:                _GotoState56Action,
	{_State265, AccessExprType}:                    _GotoState55Action,
	{_State265, IndexExprType}:                     _GotoState83Action,
	{_State265, PostfixableExprType}:               _GotoState97Action,
	{_State265, PostfixUnaryExprType}:              _GotoState96Action,
	{_State265, PrefixableExprType}:                _GotoState100Action,
	{_State265, PrefixUnaryExprType}:               _GotoState98Action,
	{_State265, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State265, MulExprType}:                       _GotoState91Action,
	{_State265, BinaryMulExprType}:                 _GotoState67Action,
	{_State265, AddExprType}:                       _GotoState57Action,
	{_State265, BinaryAddExprType}:                 _GotoState64Action,
	{_State265, CmpExprType}:                       _GotoState74Action,
	{_State265, BinaryCmpExprType}:                 _GotoState66Action,
	{_State265, AndExprType}:                       _GotoState58Action,
	{_State265, BinaryAndExprType}:                 _GotoState65Action,
	{_State265, OrExprType}:                        _GotoState94Action,
	{_State265, BinaryOrExprType}:                  _GotoState69Action,
	{_State265, InitializableTypeExprType}:         _GotoState84Action,
	{_State265, SliceTypeExprType}:                 _GotoState103Action,
	{_State265, ArrayTypeExprType}:                 _GotoState60Action,
	{_State265, MapTypeExprType}:                   _GotoState90Action,
	{_State265, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State265, AnonymousFuncExprType}:             _GotoState59Action,
	{_State267, IdentifierToken}:                   _GotoState116Action,
	{_State267, UnderscoreToken}:                   _GotoState122Action,
	{_State267, StructToken}:                       _GotoState50Action,
	{_State267, EnumToken}:                         _GotoState113Action,
	{_State267, TraitToken}:                        _GotoState121Action,
	{_State267, FuncToken}:                         _GotoState115Action,
	{_State267, LparenToken}:                       _GotoState117Action,
	{_State267, LbracketToken}:                     _GotoState42Action,
	{_State267, DotToken}:                          _GotoState112Action,
	{_State267, QuestionToken}:                     _GotoState119Action,
	{_State267, ExclaimToken}:                      _GotoState114Action,
	{_State267, TildeTildeToken}:                   _GotoState120Action,
	{_State267, BitNegToken}:                       _GotoState111Action,
	{_State267, BitAndToken}:                       _GotoState110Action,
	{_State267, ParseErrorToken}:                   _GotoState118Action,
	{_State267, InitializableTypeExprType}:         _GotoState130Action,
	{_State267, SliceTypeExprType}:                 _GotoState103Action,
	{_State267, ArrayTypeExprType}:                 _GotoState60Action,
	{_State267, MapTypeExprType}:                   _GotoState90Action,
	{_State267, AtomTypeExprType}:                  _GotoState123Action,
	{_State267, NamedTypeExprType}:                 _GotoState131Action,
	{_State267, InferredTypeExprType}:              _GotoState129Action,
	{_State267, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State267, ReturnableTypeExprType}:            _GotoState135Action,
	{_State267, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State267, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State267, TypeExprType}:                      _GotoState370Action,
	{_State267, BinaryTypeExprType}:                _GotoState124Action,
	{_State267, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State267, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State267, TraitTypeExprType}:                 _GotoState136Action,
	{_State267, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State267, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State267, FuncTypeExprType}:                  _GotoState126Action,
	{_State268, IdentifierToken}:                   _GotoState116Action,
	{_State268, UnderscoreToken}:                   _GotoState122Action,
	{_State268, StructToken}:                       _GotoState50Action,
	{_State268, EnumToken}:                         _GotoState113Action,
	{_State268, TraitToken}:                        _GotoState121Action,
	{_State268, FuncToken}:                         _GotoState115Action,
	{_State268, LparenToken}:                       _GotoState117Action,
	{_State268, LbracketToken}:                     _GotoState42Action,
	{_State268, DotToken}:                          _GotoState112Action,
	{_State268, QuestionToken}:                     _GotoState119Action,
	{_State268, ExclaimToken}:                      _GotoState114Action,
	{_State268, TildeTildeToken}:                   _GotoState120Action,
	{_State268, BitNegToken}:                       _GotoState111Action,
	{_State268, BitAndToken}:                       _GotoState110Action,
	{_State268, ParseErrorToken}:                   _GotoState118Action,
	{_State268, InitializableTypeExprType}:         _GotoState130Action,
	{_State268, SliceTypeExprType}:                 _GotoState103Action,
	{_State268, ArrayTypeExprType}:                 _GotoState60Action,
	{_State268, MapTypeExprType}:                   _GotoState90Action,
	{_State268, AtomTypeExprType}:                  _GotoState123Action,
	{_State268, NamedTypeExprType}:                 _GotoState131Action,
	{_State268, InferredTypeExprType}:              _GotoState129Action,
	{_State268, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State268, ReturnableTypeExprType}:            _GotoState135Action,
	{_State268, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State268, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State268, TypeExprType}:                      _GotoState371Action,
	{_State268, BinaryTypeExprType}:                _GotoState124Action,
	{_State268, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State268, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State268, TraitTypeExprType}:                 _GotoState136Action,
	{_State268, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State268, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State268, FuncTypeExprType}:                  _GotoState126Action,
	{_State272, AssignToken}:                       _GotoState372Action,
	{_State274, RparenToken}:                       _GotoState374Action,
	{_State274, CommaToken}:                        _GotoState373Action,
	{_State277, AddToken}:                          _GotoState253Action,
	{_State277, SubToken}:                          _GotoState255Action,
	{_State277, MulToken}:                          _GotoState254Action,
	{_State277, BinaryTypeOpType}:                  _GotoState256Action,
	{_State278, LparenToken}:                       _GotoState43Action,
	{_State278, ImplicitStructExprType}:            _GotoState375Action,
	{_State279, IdentifierToken}:                   _GotoState376Action,
	{_State280, IntegerLiteralToken}:               _GotoState40Action,
	{_State280, FloatLiteralToken}:                 _GotoState35Action,
	{_State280, RuneLiteralToken}:                  _GotoState48Action,
	{_State280, StringLiteralToken}:                _GotoState49Action,
	{_State280, IdentifierToken}:                   _GotoState38Action,
	{_State280, UnderscoreToken}:                   _GotoState53Action,
	{_State280, TrueToken}:                         _GotoState52Action,
	{_State280, FalseToken}:                        _GotoState34Action,
	{_State280, ReturnToken}:                       _GotoState47Action,
	{_State280, BreakToken}:                        _GotoState28Action,
	{_State280, ContinueToken}:                     _GotoState30Action,
	{_State280, FallthroughToken}:                  _GotoState33Action,
	{_State280, UnsafeToken}:                       _GotoState54Action,
	{_State280, StructToken}:                       _GotoState50Action,
	{_State280, FuncToken}:                         _GotoState36Action,
	{_State280, AsyncToken}:                        _GotoState25Action,
	{_State280, DeferToken}:                        _GotoState32Action,
	{_State280, VarToken}:                          _GotoState15Action,
	{_State280, LetToken}:                          _GotoState12Action,
	{_State280, NotToken}:                          _GotoState45Action,
	{_State280, LabelDeclToken}:                    _GotoState41Action,
	{_State280, LparenToken}:                       _GotoState43Action,
	{_State280, LbracketToken}:                     _GotoState42Action,
	{_State280, SubToken}:                          _GotoState51Action,
	{_State280, MulToken}:                          _GotoState44Action,
	{_State280, BitNegToken}:                       _GotoState27Action,
	{_State280, BitAndToken}:                       _GotoState26Action,
	{_State280, GreaterToken}:                      _GotoState37Action,
	{_State280, ParseErrorToken}:                   _GotoState46Action,
	{_State280, SimpleStatementType}:               _GotoState283Action,
	{_State280, OptionalSimpleStatementType}:       _GotoState377Action,
	{_State280, ExprOrImproperStructStatementType}: _GotoState77Action,
	{_State280, ExprsType}:                         _GotoState78Action,
	{_State280, CallbackOpType}:                    _GotoState72Action,
	{_State280, CallbackOpStatementType}:           _GotoState73Action,
	{_State280, UnsafeStatementType}:               _GotoState106Action,
	{_State280, JumpStatementType}:                 _GotoState86Action,
	{_State280, JumpTypeType}:                      _GotoState87Action,
	{_State280, FallthroughStatementType}:          _GotoState79Action,
	{_State280, AssignStatementType}:               _GotoState62Action,
	{_State280, UnaryOpAssignStatementType}:        _GotoState105Action,
	{_State280, BinaryOpAssignStatementType}:       _GotoState68Action,
	{_State280, VarDeclPatternType}:                _GotoState107Action,
	{_State280, VarOrLetType}:                      _GotoState24Action,
	{_State280, AssignPatternType}:                 _GotoState61Action,
	{_State280, ExprType}:                          _GotoState76Action,
	{_State280, OptionalLabelDeclType}:             _GotoState93Action,
	{_State280, SequenceExprType}:                  _GotoState101Action,
	{_State280, IfExprType}:                        _GotoState80Action,
	{_State280, SwitchExprType}:                    _GotoState104Action,
	{_State280, LoopExprType}:                      _GotoState89Action,
	{_State280, CallExprType}:                      _GotoState71Action,
	{_State280, AtomExprType}:                      _GotoState63Action,
	{_State280, ParseErrorExprType}:                _GotoState95Action,
	{_State280, LiteralExprType}:                   _GotoState88Action,
	{_State280, NamedExprType}:                     _GotoState92Action,
	{_State280, BlockExprType}:                     _GotoState70Action,
	{_State280, InitializeExprType}:                _GotoState85Action,
	{_State280, ImplicitStructExprType}:            _GotoState81Action,
	{_State280, AccessibleExprType}:                _GotoState56Action,
	{_State280, AccessExprType}:                    _GotoState55Action,
	{_State280, IndexExprType}:                     _GotoState83Action,
	{_State280, PostfixableExprType}:               _GotoState97Action,
	{_State280, PostfixUnaryExprType}:              _GotoState96Action,
	{_State280, PrefixableExprType}:                _GotoState100Action,
	{_State280, PrefixUnaryExprType}:               _GotoState98Action,
	{_State280, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State280, MulExprType}:                       _GotoState91Action,
	{_State280, BinaryMulExprType}:                 _GotoState67Action,
	{_State280, AddExprType}:                       _GotoState57Action,
	{_State280, BinaryAddExprType}:                 _GotoState64Action,
	{_State280, CmpExprType}:                       _GotoState74Action,
	{_State280, BinaryCmpExprType}:                 _GotoState66Action,
	{_State280, AndExprType}:                       _GotoState58Action,
	{_State280, BinaryAndExprType}:                 _GotoState65Action,
	{_State280, OrExprType}:                        _GotoState94Action,
	{_State280, BinaryOrExprType}:                  _GotoState69Action,
	{_State280, InitializableTypeExprType}:         _GotoState84Action,
	{_State280, SliceTypeExprType}:                 _GotoState103Action,
	{_State280, ArrayTypeExprType}:                 _GotoState60Action,
	{_State280, MapTypeExprType}:                   _GotoState90Action,
	{_State280, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State280, AnonymousFuncExprType}:             _GotoState59Action,
	{_State281, IntegerLiteralToken}:               _GotoState40Action,
	{_State281, FloatLiteralToken}:                 _GotoState35Action,
	{_State281, RuneLiteralToken}:                  _GotoState48Action,
	{_State281, StringLiteralToken}:                _GotoState49Action,
	{_State281, IdentifierToken}:                   _GotoState38Action,
	{_State281, UnderscoreToken}:                   _GotoState53Action,
	{_State281, TrueToken}:                         _GotoState52Action,
	{_State281, FalseToken}:                        _GotoState34Action,
	{_State281, StructToken}:                       _GotoState50Action,
	{_State281, FuncToken}:                         _GotoState36Action,
	{_State281, VarToken}:                          _GotoState152Action,
	{_State281, LetToken}:                          _GotoState12Action,
	{_State281, NotToken}:                          _GotoState45Action,
	{_State281, LabelDeclToken}:                    _GotoState41Action,
	{_State281, LparenToken}:                       _GotoState43Action,
	{_State281, LbracketToken}:                     _GotoState42Action,
	{_State281, DotToken}:                          _GotoState151Action,
	{_State281, SubToken}:                          _GotoState51Action,
	{_State281, MulToken}:                          _GotoState44Action,
	{_State281, BitNegToken}:                       _GotoState27Action,
	{_State281, BitAndToken}:                       _GotoState26Action,
	{_State281, GreaterToken}:                      _GotoState37Action,
	{_State281, ParseErrorToken}:                   _GotoState46Action,
	{_State281, VarDeclPatternType}:                _GotoState107Action,
	{_State281, VarOrLetType}:                      _GotoState24Action,
	{_State281, AssignPatternType}:                 _GotoState153Action,
	{_State281, CasePatternType}:                   _GotoState378Action,
	{_State281, OptionalLabelDeclType}:             _GotoState156Action,
	{_State281, SequenceExprType}:                  _GotoState157Action,
	{_State281, CallExprType}:                      _GotoState71Action,
	{_State281, AtomExprType}:                      _GotoState63Action,
	{_State281, ParseErrorExprType}:                _GotoState95Action,
	{_State281, LiteralExprType}:                   _GotoState88Action,
	{_State281, NamedExprType}:                     _GotoState92Action,
	{_State281, BlockExprType}:                     _GotoState70Action,
	{_State281, InitializeExprType}:                _GotoState85Action,
	{_State281, ImplicitStructExprType}:            _GotoState81Action,
	{_State281, AccessibleExprType}:                _GotoState108Action,
	{_State281, AccessExprType}:                    _GotoState55Action,
	{_State281, IndexExprType}:                     _GotoState83Action,
	{_State281, PostfixableExprType}:               _GotoState97Action,
	{_State281, PostfixUnaryExprType}:              _GotoState96Action,
	{_State281, PrefixableExprType}:                _GotoState100Action,
	{_State281, PrefixUnaryExprType}:               _GotoState98Action,
	{_State281, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State281, MulExprType}:                       _GotoState91Action,
	{_State281, BinaryMulExprType}:                 _GotoState67Action,
	{_State281, AddExprType}:                       _GotoState57Action,
	{_State281, BinaryAddExprType}:                 _GotoState64Action,
	{_State281, CmpExprType}:                       _GotoState74Action,
	{_State281, BinaryCmpExprType}:                 _GotoState66Action,
	{_State281, AndExprType}:                       _GotoState58Action,
	{_State281, BinaryAndExprType}:                 _GotoState65Action,
	{_State281, OrExprType}:                        _GotoState94Action,
	{_State281, BinaryOrExprType}:                  _GotoState69Action,
	{_State281, InitializableTypeExprType}:         _GotoState84Action,
	{_State281, SliceTypeExprType}:                 _GotoState103Action,
	{_State281, ArrayTypeExprType}:                 _GotoState60Action,
	{_State281, MapTypeExprType}:                   _GotoState90Action,
	{_State281, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State281, AnonymousFuncExprType}:             _GotoState59Action,
	{_State285, RparenToken}:                       _GotoState379Action,
	{_State286, CommaToken}:                        _GotoState380Action,
	{_State290, RparenToken}:                       _GotoState381Action,
	{_State291, NewlinesToken}:                     _GotoState383Action,
	{_State291, CommaToken}:                        _GotoState382Action,
	{_State293, IdentifierToken}:                   _GotoState116Action,
	{_State293, UnderscoreToken}:                   _GotoState122Action,
	{_State293, StructToken}:                       _GotoState50Action,
	{_State293, EnumToken}:                         _GotoState113Action,
	{_State293, TraitToken}:                        _GotoState121Action,
	{_State293, FuncToken}:                         _GotoState115Action,
	{_State293, LparenToken}:                       _GotoState117Action,
	{_State293, LbracketToken}:                     _GotoState42Action,
	{_State293, DotToken}:                          _GotoState112Action,
	{_State293, QuestionToken}:                     _GotoState119Action,
	{_State293, ExclaimToken}:                      _GotoState114Action,
	{_State293, TildeTildeToken}:                   _GotoState120Action,
	{_State293, BitNegToken}:                       _GotoState111Action,
	{_State293, BitAndToken}:                       _GotoState110Action,
	{_State293, ParseErrorToken}:                   _GotoState118Action,
	{_State293, InitializableTypeExprType}:         _GotoState130Action,
	{_State293, SliceTypeExprType}:                 _GotoState103Action,
	{_State293, ArrayTypeExprType}:                 _GotoState60Action,
	{_State293, MapTypeExprType}:                   _GotoState90Action,
	{_State293, AtomTypeExprType}:                  _GotoState123Action,
	{_State293, NamedTypeExprType}:                 _GotoState131Action,
	{_State293, InferredTypeExprType}:              _GotoState129Action,
	{_State293, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State293, ReturnableTypeExprType}:            _GotoState135Action,
	{_State293, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State293, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State293, TypeExprType}:                      _GotoState384Action,
	{_State293, BinaryTypeExprType}:                _GotoState124Action,
	{_State293, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State293, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State293, TraitTypeExprType}:                 _GotoState136Action,
	{_State293, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State293, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State293, FuncTypeExprType}:                  _GotoState126Action,
	{_State294, IntegerLiteralToken}:               _GotoState385Action,
	{_State297, IntegerLiteralToken}:               _GotoState40Action,
	{_State297, FloatLiteralToken}:                 _GotoState35Action,
	{_State297, RuneLiteralToken}:                  _GotoState48Action,
	{_State297, StringLiteralToken}:                _GotoState49Action,
	{_State297, IdentifierToken}:                   _GotoState38Action,
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
	{_State297, SubToken}:                          _GotoState51Action,
	{_State297, MulToken}:                          _GotoState44Action,
	{_State297, BitNegToken}:                       _GotoState27Action,
	{_State297, BitAndToken}:                       _GotoState26Action,
	{_State297, GreaterToken}:                      _GotoState37Action,
	{_State297, ParseErrorToken}:                   _GotoState46Action,
	{_State297, VarDeclPatternType}:                _GotoState107Action,
	{_State297, VarOrLetType}:                      _GotoState24Action,
	{_State297, ExprType}:                          _GotoState386Action,
	{_State297, OptionalLabelDeclType}:             _GotoState93Action,
	{_State297, SequenceExprType}:                  _GotoState109Action,
	{_State297, IfExprType}:                        _GotoState80Action,
	{_State297, SwitchExprType}:                    _GotoState104Action,
	{_State297, LoopExprType}:                      _GotoState89Action,
	{_State297, CallExprType}:                      _GotoState71Action,
	{_State297, AtomExprType}:                      _GotoState63Action,
	{_State297, ParseErrorExprType}:                _GotoState95Action,
	{_State297, LiteralExprType}:                   _GotoState88Action,
	{_State297, NamedExprType}:                     _GotoState92Action,
	{_State297, BlockExprType}:                     _GotoState70Action,
	{_State297, InitializeExprType}:                _GotoState85Action,
	{_State297, ImplicitStructExprType}:            _GotoState81Action,
	{_State297, AccessibleExprType}:                _GotoState108Action,
	{_State297, AccessExprType}:                    _GotoState55Action,
	{_State297, IndexExprType}:                     _GotoState83Action,
	{_State297, PostfixableExprType}:               _GotoState97Action,
	{_State297, PostfixUnaryExprType}:              _GotoState96Action,
	{_State297, PrefixableExprType}:                _GotoState100Action,
	{_State297, PrefixUnaryExprType}:               _GotoState98Action,
	{_State297, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State297, MulExprType}:                       _GotoState91Action,
	{_State297, BinaryMulExprType}:                 _GotoState67Action,
	{_State297, AddExprType}:                       _GotoState57Action,
	{_State297, BinaryAddExprType}:                 _GotoState64Action,
	{_State297, CmpExprType}:                       _GotoState74Action,
	{_State297, BinaryCmpExprType}:                 _GotoState66Action,
	{_State297, AndExprType}:                       _GotoState58Action,
	{_State297, BinaryAndExprType}:                 _GotoState65Action,
	{_State297, OrExprType}:                        _GotoState94Action,
	{_State297, BinaryOrExprType}:                  _GotoState69Action,
	{_State297, InitializableTypeExprType}:         _GotoState84Action,
	{_State297, SliceTypeExprType}:                 _GotoState103Action,
	{_State297, ArrayTypeExprType}:                 _GotoState60Action,
	{_State297, MapTypeExprType}:                   _GotoState90Action,
	{_State297, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State297, AnonymousFuncExprType}:             _GotoState59Action,
	{_State299, IntegerLiteralToken}:               _GotoState40Action,
	{_State299, FloatLiteralToken}:                 _GotoState35Action,
	{_State299, RuneLiteralToken}:                  _GotoState48Action,
	{_State299, StringLiteralToken}:                _GotoState49Action,
	{_State299, IdentifierToken}:                   _GotoState38Action,
	{_State299, UnderscoreToken}:                   _GotoState53Action,
	{_State299, TrueToken}:                         _GotoState52Action,
	{_State299, FalseToken}:                        _GotoState34Action,
	{_State299, StructToken}:                       _GotoState50Action,
	{_State299, FuncToken}:                         _GotoState36Action,
	{_State299, VarToken}:                          _GotoState15Action,
	{_State299, LetToken}:                          _GotoState12Action,
	{_State299, NotToken}:                          _GotoState45Action,
	{_State299, LabelDeclToken}:                    _GotoState41Action,
	{_State299, LparenToken}:                       _GotoState43Action,
	{_State299, LbracketToken}:                     _GotoState42Action,
	{_State299, SubToken}:                          _GotoState51Action,
	{_State299, MulToken}:                          _GotoState44Action,
	{_State299, BitNegToken}:                       _GotoState27Action,
	{_State299, BitAndToken}:                       _GotoState26Action,
	{_State299, GreaterToken}:                      _GotoState37Action,
	{_State299, ParseErrorToken}:                   _GotoState46Action,
	{_State299, VarDeclPatternType}:                _GotoState107Action,
	{_State299, VarOrLetType}:                      _GotoState24Action,
	{_State299, ExprType}:                          _GotoState387Action,
	{_State299, OptionalLabelDeclType}:             _GotoState93Action,
	{_State299, SequenceExprType}:                  _GotoState109Action,
	{_State299, IfExprType}:                        _GotoState80Action,
	{_State299, SwitchExprType}:                    _GotoState104Action,
	{_State299, LoopExprType}:                      _GotoState89Action,
	{_State299, CallExprType}:                      _GotoState71Action,
	{_State299, AtomExprType}:                      _GotoState63Action,
	{_State299, ParseErrorExprType}:                _GotoState95Action,
	{_State299, LiteralExprType}:                   _GotoState88Action,
	{_State299, NamedExprType}:                     _GotoState92Action,
	{_State299, BlockExprType}:                     _GotoState70Action,
	{_State299, InitializeExprType}:                _GotoState85Action,
	{_State299, ImplicitStructExprType}:            _GotoState81Action,
	{_State299, AccessibleExprType}:                _GotoState108Action,
	{_State299, AccessExprType}:                    _GotoState55Action,
	{_State299, IndexExprType}:                     _GotoState83Action,
	{_State299, PostfixableExprType}:               _GotoState97Action,
	{_State299, PostfixUnaryExprType}:              _GotoState96Action,
	{_State299, PrefixableExprType}:                _GotoState100Action,
	{_State299, PrefixUnaryExprType}:               _GotoState98Action,
	{_State299, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State299, MulExprType}:                       _GotoState91Action,
	{_State299, BinaryMulExprType}:                 _GotoState67Action,
	{_State299, AddExprType}:                       _GotoState57Action,
	{_State299, BinaryAddExprType}:                 _GotoState64Action,
	{_State299, CmpExprType}:                       _GotoState74Action,
	{_State299, BinaryCmpExprType}:                 _GotoState66Action,
	{_State299, AndExprType}:                       _GotoState58Action,
	{_State299, BinaryAndExprType}:                 _GotoState65Action,
	{_State299, OrExprType}:                        _GotoState94Action,
	{_State299, BinaryOrExprType}:                  _GotoState69Action,
	{_State299, InitializableTypeExprType}:         _GotoState84Action,
	{_State299, SliceTypeExprType}:                 _GotoState103Action,
	{_State299, ArrayTypeExprType}:                 _GotoState60Action,
	{_State299, MapTypeExprType}:                   _GotoState90Action,
	{_State299, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State299, AnonymousFuncExprType}:             _GotoState59Action,
	{_State300, IntegerLiteralToken}:               _GotoState40Action,
	{_State300, FloatLiteralToken}:                 _GotoState35Action,
	{_State300, RuneLiteralToken}:                  _GotoState48Action,
	{_State300, StringLiteralToken}:                _GotoState49Action,
	{_State300, IdentifierToken}:                   _GotoState38Action,
	{_State300, UnderscoreToken}:                   _GotoState53Action,
	{_State300, TrueToken}:                         _GotoState52Action,
	{_State300, FalseToken}:                        _GotoState34Action,
	{_State300, StructToken}:                       _GotoState50Action,
	{_State300, FuncToken}:                         _GotoState36Action,
	{_State300, VarToken}:                          _GotoState15Action,
	{_State300, LetToken}:                          _GotoState12Action,
	{_State300, NotToken}:                          _GotoState45Action,
	{_State300, LabelDeclToken}:                    _GotoState41Action,
	{_State300, LparenToken}:                       _GotoState43Action,
	{_State300, LbracketToken}:                     _GotoState42Action,
	{_State300, SubToken}:                          _GotoState51Action,
	{_State300, MulToken}:                          _GotoState44Action,
	{_State300, BitNegToken}:                       _GotoState27Action,
	{_State300, BitAndToken}:                       _GotoState26Action,
	{_State300, GreaterToken}:                      _GotoState37Action,
	{_State300, ParseErrorToken}:                   _GotoState46Action,
	{_State300, VarDeclPatternType}:                _GotoState107Action,
	{_State300, VarOrLetType}:                      _GotoState24Action,
	{_State300, ExprType}:                          _GotoState388Action,
	{_State300, OptionalLabelDeclType}:             _GotoState93Action,
	{_State300, SequenceExprType}:                  _GotoState109Action,
	{_State300, IfExprType}:                        _GotoState80Action,
	{_State300, SwitchExprType}:                    _GotoState104Action,
	{_State300, LoopExprType}:                      _GotoState89Action,
	{_State300, CallExprType}:                      _GotoState71Action,
	{_State300, AtomExprType}:                      _GotoState63Action,
	{_State300, ParseErrorExprType}:                _GotoState95Action,
	{_State300, LiteralExprType}:                   _GotoState88Action,
	{_State300, NamedExprType}:                     _GotoState92Action,
	{_State300, BlockExprType}:                     _GotoState70Action,
	{_State300, InitializeExprType}:                _GotoState85Action,
	{_State300, ImplicitStructExprType}:            _GotoState81Action,
	{_State300, AccessibleExprType}:                _GotoState108Action,
	{_State300, AccessExprType}:                    _GotoState55Action,
	{_State300, IndexExprType}:                     _GotoState83Action,
	{_State300, PostfixableExprType}:               _GotoState97Action,
	{_State300, PostfixUnaryExprType}:              _GotoState96Action,
	{_State300, PrefixableExprType}:                _GotoState100Action,
	{_State300, PrefixUnaryExprType}:               _GotoState98Action,
	{_State300, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State300, MulExprType}:                       _GotoState91Action,
	{_State300, BinaryMulExprType}:                 _GotoState67Action,
	{_State300, AddExprType}:                       _GotoState57Action,
	{_State300, BinaryAddExprType}:                 _GotoState64Action,
	{_State300, CmpExprType}:                       _GotoState74Action,
	{_State300, BinaryCmpExprType}:                 _GotoState66Action,
	{_State300, AndExprType}:                       _GotoState58Action,
	{_State300, BinaryAndExprType}:                 _GotoState65Action,
	{_State300, OrExprType}:                        _GotoState94Action,
	{_State300, BinaryOrExprType}:                  _GotoState69Action,
	{_State300, InitializableTypeExprType}:         _GotoState84Action,
	{_State300, SliceTypeExprType}:                 _GotoState103Action,
	{_State300, ArrayTypeExprType}:                 _GotoState60Action,
	{_State300, MapTypeExprType}:                   _GotoState90Action,
	{_State300, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State300, AnonymousFuncExprType}:             _GotoState59Action,
	{_State302, IntegerLiteralToken}:               _GotoState40Action,
	{_State302, FloatLiteralToken}:                 _GotoState35Action,
	{_State302, RuneLiteralToken}:                  _GotoState48Action,
	{_State302, StringLiteralToken}:                _GotoState49Action,
	{_State302, IdentifierToken}:                   _GotoState170Action,
	{_State302, UnderscoreToken}:                   _GotoState53Action,
	{_State302, TrueToken}:                         _GotoState52Action,
	{_State302, FalseToken}:                        _GotoState34Action,
	{_State302, StructToken}:                       _GotoState50Action,
	{_State302, FuncToken}:                         _GotoState36Action,
	{_State302, VarToken}:                          _GotoState15Action,
	{_State302, LetToken}:                          _GotoState12Action,
	{_State302, NotToken}:                          _GotoState45Action,
	{_State302, LabelDeclToken}:                    _GotoState41Action,
	{_State302, LparenToken}:                       _GotoState43Action,
	{_State302, LbracketToken}:                     _GotoState42Action,
	{_State302, ColonToken}:                        _GotoState168Action,
	{_State302, EllipsisToken}:                     _GotoState169Action,
	{_State302, SubToken}:                          _GotoState51Action,
	{_State302, MulToken}:                          _GotoState44Action,
	{_State302, BitNegToken}:                       _GotoState27Action,
	{_State302, BitAndToken}:                       _GotoState26Action,
	{_State302, GreaterToken}:                      _GotoState37Action,
	{_State302, ParseErrorToken}:                   _GotoState46Action,
	{_State302, VarDeclPatternType}:                _GotoState107Action,
	{_State302, VarOrLetType}:                      _GotoState24Action,
	{_State302, ExprType}:                          _GotoState174Action,
	{_State302, OptionalLabelDeclType}:             _GotoState93Action,
	{_State302, SequenceExprType}:                  _GotoState109Action,
	{_State302, IfExprType}:                        _GotoState80Action,
	{_State302, SwitchExprType}:                    _GotoState104Action,
	{_State302, LoopExprType}:                      _GotoState89Action,
	{_State302, CallExprType}:                      _GotoState71Action,
	{_State302, ArgumentType}:                      _GotoState389Action,
	{_State302, ColonExprType}:                     _GotoState173Action,
	{_State302, AtomExprType}:                      _GotoState63Action,
	{_State302, ParseErrorExprType}:                _GotoState95Action,
	{_State302, LiteralExprType}:                   _GotoState88Action,
	{_State302, NamedExprType}:                     _GotoState92Action,
	{_State302, BlockExprType}:                     _GotoState70Action,
	{_State302, InitializeExprType}:                _GotoState85Action,
	{_State302, ImplicitStructExprType}:            _GotoState81Action,
	{_State302, AccessibleExprType}:                _GotoState108Action,
	{_State302, AccessExprType}:                    _GotoState55Action,
	{_State302, IndexExprType}:                     _GotoState83Action,
	{_State302, PostfixableExprType}:               _GotoState97Action,
	{_State302, PostfixUnaryExprType}:              _GotoState96Action,
	{_State302, PrefixableExprType}:                _GotoState100Action,
	{_State302, PrefixUnaryExprType}:               _GotoState98Action,
	{_State302, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State302, MulExprType}:                       _GotoState91Action,
	{_State302, BinaryMulExprType}:                 _GotoState67Action,
	{_State302, AddExprType}:                       _GotoState57Action,
	{_State302, BinaryAddExprType}:                 _GotoState64Action,
	{_State302, CmpExprType}:                       _GotoState74Action,
	{_State302, BinaryCmpExprType}:                 _GotoState66Action,
	{_State302, AndExprType}:                       _GotoState58Action,
	{_State302, BinaryAndExprType}:                 _GotoState65Action,
	{_State302, OrExprType}:                        _GotoState94Action,
	{_State302, BinaryOrExprType}:                  _GotoState69Action,
	{_State302, InitializableTypeExprType}:         _GotoState84Action,
	{_State302, SliceTypeExprType}:                 _GotoState103Action,
	{_State302, ArrayTypeExprType}:                 _GotoState60Action,
	{_State302, MapTypeExprType}:                   _GotoState90Action,
	{_State302, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State302, AnonymousFuncExprType}:             _GotoState59Action,
	{_State303, RparenToken}:                       _GotoState390Action,
	{_State305, NewlinesToken}:                     _GotoState392Action,
	{_State305, CommaToken}:                        _GotoState391Action,
	{_State306, GreaterToken}:                      _GotoState393Action,
	{_State307, CommaToken}:                        _GotoState394Action,
	{_State308, RbracketToken}:                     _GotoState395Action,
	{_State309, AddToken}:                          _GotoState253Action,
	{_State309, SubToken}:                          _GotoState255Action,
	{_State309, MulToken}:                          _GotoState254Action,
	{_State309, BinaryTypeOpType}:                  _GotoState256Action,
	{_State311, RbracketToken}:                     _GotoState396Action,
	{_State313, IntegerLiteralToken}:               _GotoState40Action,
	{_State313, FloatLiteralToken}:                 _GotoState35Action,
	{_State313, RuneLiteralToken}:                  _GotoState48Action,
	{_State313, StringLiteralToken}:                _GotoState49Action,
	{_State313, IdentifierToken}:                   _GotoState170Action,
	{_State313, UnderscoreToken}:                   _GotoState53Action,
	{_State313, TrueToken}:                         _GotoState52Action,
	{_State313, FalseToken}:                        _GotoState34Action,
	{_State313, StructToken}:                       _GotoState50Action,
	{_State313, FuncToken}:                         _GotoState36Action,
	{_State313, VarToken}:                          _GotoState15Action,
	{_State313, LetToken}:                          _GotoState12Action,
	{_State313, NotToken}:                          _GotoState45Action,
	{_State313, LabelDeclToken}:                    _GotoState41Action,
	{_State313, LparenToken}:                       _GotoState43Action,
	{_State313, LbracketToken}:                     _GotoState42Action,
	{_State313, ColonToken}:                        _GotoState168Action,
	{_State313, EllipsisToken}:                     _GotoState169Action,
	{_State313, SubToken}:                          _GotoState51Action,
	{_State313, MulToken}:                          _GotoState44Action,
	{_State313, BitNegToken}:                       _GotoState27Action,
	{_State313, BitAndToken}:                       _GotoState26Action,
	{_State313, GreaterToken}:                      _GotoState37Action,
	{_State313, ParseErrorToken}:                   _GotoState46Action,
	{_State313, VarDeclPatternType}:                _GotoState107Action,
	{_State313, VarOrLetType}:                      _GotoState24Action,
	{_State313, ExprType}:                          _GotoState174Action,
	{_State313, OptionalLabelDeclType}:             _GotoState93Action,
	{_State313, SequenceExprType}:                  _GotoState109Action,
	{_State313, IfExprType}:                        _GotoState80Action,
	{_State313, SwitchExprType}:                    _GotoState104Action,
	{_State313, LoopExprType}:                      _GotoState89Action,
	{_State313, CallExprType}:                      _GotoState71Action,
	{_State313, ProperArgumentsType}:               _GotoState175Action,
	{_State313, ArgumentsType}:                     _GotoState397Action,
	{_State313, ArgumentType}:                      _GotoState171Action,
	{_State313, ColonExprType}:                     _GotoState173Action,
	{_State313, AtomExprType}:                      _GotoState63Action,
	{_State313, ParseErrorExprType}:                _GotoState95Action,
	{_State313, LiteralExprType}:                   _GotoState88Action,
	{_State313, NamedExprType}:                     _GotoState92Action,
	{_State313, BlockExprType}:                     _GotoState70Action,
	{_State313, InitializeExprType}:                _GotoState85Action,
	{_State313, ImplicitStructExprType}:            _GotoState81Action,
	{_State313, AccessibleExprType}:                _GotoState108Action,
	{_State313, AccessExprType}:                    _GotoState55Action,
	{_State313, IndexExprType}:                     _GotoState83Action,
	{_State313, PostfixableExprType}:               _GotoState97Action,
	{_State313, PostfixUnaryExprType}:              _GotoState96Action,
	{_State313, PrefixableExprType}:                _GotoState100Action,
	{_State313, PrefixUnaryExprType}:               _GotoState98Action,
	{_State313, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State313, MulExprType}:                       _GotoState91Action,
	{_State313, BinaryMulExprType}:                 _GotoState67Action,
	{_State313, AddExprType}:                       _GotoState57Action,
	{_State313, BinaryAddExprType}:                 _GotoState64Action,
	{_State313, CmpExprType}:                       _GotoState74Action,
	{_State313, BinaryCmpExprType}:                 _GotoState66Action,
	{_State313, AndExprType}:                       _GotoState58Action,
	{_State313, BinaryAndExprType}:                 _GotoState65Action,
	{_State313, OrExprType}:                        _GotoState94Action,
	{_State313, BinaryOrExprType}:                  _GotoState69Action,
	{_State313, InitializableTypeExprType}:         _GotoState84Action,
	{_State313, SliceTypeExprType}:                 _GotoState103Action,
	{_State313, ArrayTypeExprType}:                 _GotoState60Action,
	{_State313, MapTypeExprType}:                   _GotoState90Action,
	{_State313, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State313, AnonymousFuncExprType}:             _GotoState59Action,
	{_State314, MulToken}:                          _GotoState225Action,
	{_State314, DivToken}:                          _GotoState223Action,
	{_State314, ModToken}:                          _GotoState224Action,
	{_State314, BitAndToken}:                       _GotoState220Action,
	{_State314, BitLshiftToken}:                    _GotoState221Action,
	{_State314, BitRshiftToken}:                    _GotoState222Action,
	{_State314, MulOpType}:                         _GotoState226Action,
	{_State315, EqualToken}:                        _GotoState209Action,
	{_State315, NotEqualToken}:                     _GotoState214Action,
	{_State315, LessToken}:                         _GotoState212Action,
	{_State315, LessOrEqualToken}:                  _GotoState213Action,
	{_State315, GreaterToken}:                      _GotoState210Action,
	{_State315, GreaterOrEqualToken}:               _GotoState211Action,
	{_State315, CmpOpType}:                         _GotoState215Action,
	{_State317, AddToken}:                          _GotoState200Action,
	{_State317, SubToken}:                          _GotoState203Action,
	{_State317, BitXorToken}:                       _GotoState202Action,
	{_State317, BitOrToken}:                        _GotoState201Action,
	{_State317, AddOpType}:                         _GotoState204Action,
	{_State319, RparenToken}:                       _GotoState398Action,
	{_State320, CommaToken}:                        _GotoState216Action,
	{_State322, ForToken}:                          _GotoState399Action,
	{_State323, InToken}:                           _GotoState401Action,
	{_State323, AssignToken}:                       _GotoState400Action,
	{_State324, SemicolonToken}:                    _GotoState402Action,
	{_State325, DoToken}:                           _GotoState403Action,
	{_State326, IntegerLiteralToken}:               _GotoState40Action,
	{_State326, FloatLiteralToken}:                 _GotoState35Action,
	{_State326, RuneLiteralToken}:                  _GotoState48Action,
	{_State326, StringLiteralToken}:                _GotoState49Action,
	{_State326, IdentifierToken}:                   _GotoState38Action,
	{_State326, UnderscoreToken}:                   _GotoState53Action,
	{_State326, TrueToken}:                         _GotoState52Action,
	{_State326, FalseToken}:                        _GotoState34Action,
	{_State326, StructToken}:                       _GotoState50Action,
	{_State326, FuncToken}:                         _GotoState36Action,
	{_State326, VarToken}:                          _GotoState152Action,
	{_State326, LetToken}:                          _GotoState12Action,
	{_State326, NotToken}:                          _GotoState45Action,
	{_State326, LabelDeclToken}:                    _GotoState41Action,
	{_State326, LparenToken}:                       _GotoState43Action,
	{_State326, LbracketToken}:                     _GotoState42Action,
	{_State326, DotToken}:                          _GotoState151Action,
	{_State326, SubToken}:                          _GotoState51Action,
	{_State326, MulToken}:                          _GotoState44Action,
	{_State326, BitNegToken}:                       _GotoState27Action,
	{_State326, BitAndToken}:                       _GotoState26Action,
	{_State326, GreaterToken}:                      _GotoState37Action,
	{_State326, ParseErrorToken}:                   _GotoState46Action,
	{_State326, CasePatternsType}:                  _GotoState404Action,
	{_State326, VarDeclPatternType}:                _GotoState107Action,
	{_State326, VarOrLetType}:                      _GotoState24Action,
	{_State326, AssignPatternType}:                 _GotoState153Action,
	{_State326, CasePatternType}:                   _GotoState154Action,
	{_State326, OptionalLabelDeclType}:             _GotoState156Action,
	{_State326, SequenceExprType}:                  _GotoState157Action,
	{_State326, CallExprType}:                      _GotoState71Action,
	{_State326, AtomExprType}:                      _GotoState63Action,
	{_State326, ParseErrorExprType}:                _GotoState95Action,
	{_State326, LiteralExprType}:                   _GotoState88Action,
	{_State326, NamedExprType}:                     _GotoState92Action,
	{_State326, BlockExprType}:                     _GotoState70Action,
	{_State326, InitializeExprType}:                _GotoState85Action,
	{_State326, ImplicitStructExprType}:            _GotoState81Action,
	{_State326, AccessibleExprType}:                _GotoState108Action,
	{_State326, AccessExprType}:                    _GotoState55Action,
	{_State326, IndexExprType}:                     _GotoState83Action,
	{_State326, PostfixableExprType}:               _GotoState97Action,
	{_State326, PostfixUnaryExprType}:              _GotoState96Action,
	{_State326, PrefixableExprType}:                _GotoState100Action,
	{_State326, PrefixUnaryExprType}:               _GotoState98Action,
	{_State326, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State326, MulExprType}:                       _GotoState91Action,
	{_State326, BinaryMulExprType}:                 _GotoState67Action,
	{_State326, AddExprType}:                       _GotoState57Action,
	{_State326, BinaryAddExprType}:                 _GotoState64Action,
	{_State326, CmpExprType}:                       _GotoState74Action,
	{_State326, BinaryCmpExprType}:                 _GotoState66Action,
	{_State326, AndExprType}:                       _GotoState58Action,
	{_State326, BinaryAndExprType}:                 _GotoState65Action,
	{_State326, OrExprType}:                        _GotoState94Action,
	{_State326, BinaryOrExprType}:                  _GotoState69Action,
	{_State326, InitializableTypeExprType}:         _GotoState84Action,
	{_State326, SliceTypeExprType}:                 _GotoState103Action,
	{_State326, ArrayTypeExprType}:                 _GotoState60Action,
	{_State326, MapTypeExprType}:                   _GotoState90Action,
	{_State326, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State326, AnonymousFuncExprType}:             _GotoState59Action,
	{_State327, LbraceToken}:                       _GotoState11Action,
	{_State327, StatementBlockType}:                _GotoState405Action,
	{_State329, LbraceToken}:                       _GotoState11Action,
	{_State329, StatementBlockType}:                _GotoState406Action,
	{_State330, AndToken}:                          _GotoState205Action,
	{_State331, RparenToken}:                       _GotoState407Action,
	{_State332, NewlinesToken}:                     _GotoState408Action,
	{_State332, OrToken}:                           _GotoState409Action,
	{_State333, NewlinesToken}:                     _GotoState410Action,
	{_State333, OrToken}:                           _GotoState411Action,
	{_State334, IdentifierToken}:                   _GotoState116Action,
	{_State334, UnderscoreToken}:                   _GotoState122Action,
	{_State334, StructToken}:                       _GotoState50Action,
	{_State334, EnumToken}:                         _GotoState113Action,
	{_State334, TraitToken}:                        _GotoState121Action,
	{_State334, FuncToken}:                         _GotoState115Action,
	{_State334, LparenToken}:                       _GotoState117Action,
	{_State334, LbracketToken}:                     _GotoState42Action,
	{_State334, DotToken}:                          _GotoState112Action,
	{_State334, QuestionToken}:                     _GotoState119Action,
	{_State334, ExclaimToken}:                      _GotoState114Action,
	{_State334, TildeTildeToken}:                   _GotoState120Action,
	{_State334, BitNegToken}:                       _GotoState111Action,
	{_State334, BitAndToken}:                       _GotoState110Action,
	{_State334, ParseErrorToken}:                   _GotoState118Action,
	{_State334, InitializableTypeExprType}:         _GotoState130Action,
	{_State334, SliceTypeExprType}:                 _GotoState103Action,
	{_State334, ArrayTypeExprType}:                 _GotoState60Action,
	{_State334, MapTypeExprType}:                   _GotoState90Action,
	{_State334, AtomTypeExprType}:                  _GotoState123Action,
	{_State334, NamedTypeExprType}:                 _GotoState131Action,
	{_State334, InferredTypeExprType}:              _GotoState129Action,
	{_State334, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State334, ReturnableTypeExprType}:            _GotoState135Action,
	{_State334, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State334, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State334, TypeExprType}:                      _GotoState412Action,
	{_State334, BinaryTypeExprType}:                _GotoState124Action,
	{_State334, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State334, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State334, TraitTypeExprType}:                 _GotoState136Action,
	{_State334, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State334, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State334, FuncTypeExprType}:                  _GotoState126Action,
	{_State335, IdentifierToken}:                   _GotoState116Action,
	{_State335, UnderscoreToken}:                   _GotoState122Action,
	{_State335, StructToken}:                       _GotoState50Action,
	{_State335, EnumToken}:                         _GotoState113Action,
	{_State335, TraitToken}:                        _GotoState121Action,
	{_State335, FuncToken}:                         _GotoState115Action,
	{_State335, LparenToken}:                       _GotoState117Action,
	{_State335, LbracketToken}:                     _GotoState42Action,
	{_State335, DotToken}:                          _GotoState344Action,
	{_State335, QuestionToken}:                     _GotoState119Action,
	{_State335, ExclaimToken}:                      _GotoState114Action,
	{_State335, DollarLbracketToken}:               _GotoState187Action,
	{_State335, EllipsisToken}:                     _GotoState363Action,
	{_State335, TildeTildeToken}:                   _GotoState120Action,
	{_State335, BitNegToken}:                       _GotoState111Action,
	{_State335, BitAndToken}:                       _GotoState110Action,
	{_State335, ParseErrorToken}:                   _GotoState118Action,
	{_State335, InitializableTypeExprType}:         _GotoState130Action,
	{_State335, SliceTypeExprType}:                 _GotoState103Action,
	{_State335, ArrayTypeExprType}:                 _GotoState60Action,
	{_State335, MapTypeExprType}:                   _GotoState90Action,
	{_State335, AtomTypeExprType}:                  _GotoState123Action,
	{_State335, NamedTypeExprType}:                 _GotoState131Action,
	{_State335, InferredTypeExprType}:              _GotoState129Action,
	{_State335, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State335, ReturnableTypeExprType}:            _GotoState135Action,
	{_State335, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State335, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State335, TypeExprType}:                      _GotoState364Action,
	{_State335, BinaryTypeExprType}:                _GotoState124Action,
	{_State335, GenericTypeArgumentsType}:          _GotoState239Action,
	{_State335, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State335, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State335, TraitTypeExprType}:                 _GotoState136Action,
	{_State335, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State335, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State335, FuncTypeExprType}:                  _GotoState126Action,
	{_State336, IdentifierToken}:                   _GotoState116Action,
	{_State336, UnderscoreToken}:                   _GotoState122Action,
	{_State336, StructToken}:                       _GotoState50Action,
	{_State336, EnumToken}:                         _GotoState113Action,
	{_State336, TraitToken}:                        _GotoState121Action,
	{_State336, FuncToken}:                         _GotoState115Action,
	{_State336, LparenToken}:                       _GotoState117Action,
	{_State336, LbracketToken}:                     _GotoState42Action,
	{_State336, DotToken}:                          _GotoState112Action,
	{_State336, QuestionToken}:                     _GotoState119Action,
	{_State336, ExclaimToken}:                      _GotoState114Action,
	{_State336, EllipsisToken}:                     _GotoState365Action,
	{_State336, TildeTildeToken}:                   _GotoState120Action,
	{_State336, BitNegToken}:                       _GotoState111Action,
	{_State336, BitAndToken}:                       _GotoState110Action,
	{_State336, ParseErrorToken}:                   _GotoState118Action,
	{_State336, InitializableTypeExprType}:         _GotoState130Action,
	{_State336, SliceTypeExprType}:                 _GotoState103Action,
	{_State336, ArrayTypeExprType}:                 _GotoState60Action,
	{_State336, MapTypeExprType}:                   _GotoState90Action,
	{_State336, AtomTypeExprType}:                  _GotoState123Action,
	{_State336, NamedTypeExprType}:                 _GotoState131Action,
	{_State336, InferredTypeExprType}:              _GotoState129Action,
	{_State336, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State336, ReturnableTypeExprType}:            _GotoState135Action,
	{_State336, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State336, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State336, TypeExprType}:                      _GotoState366Action,
	{_State336, BinaryTypeExprType}:                _GotoState124Action,
	{_State336, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State336, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State336, TraitTypeExprType}:                 _GotoState136Action,
	{_State336, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State336, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State336, FuncTypeExprType}:                  _GotoState126Action,
	{_State338, RparenToken}:                       _GotoState413Action,
	{_State339, CommaToken}:                        _GotoState414Action,
	{_State341, AddToken}:                          _GotoState253Action,
	{_State341, SubToken}:                          _GotoState255Action,
	{_State341, MulToken}:                          _GotoState254Action,
	{_State341, BinaryTypeOpType}:                  _GotoState256Action,
	{_State342, DollarLbracketToken}:               _GotoState187Action,
	{_State342, GenericTypeArgumentsType}:          _GotoState415Action,
	{_State343, LparenToken}:                       _GotoState416Action,
	{_State344, IdentifierToken}:                   _GotoState342Action,
	{_State345, AssignToken}:                       _GotoState353Action,
	{_State345, AddToken}:                          _GotoState253Action,
	{_State345, SubToken}:                          _GotoState255Action,
	{_State345, MulToken}:                          _GotoState254Action,
	{_State345, BinaryTypeOpType}:                  _GotoState256Action,
	{_State345, OptionalDefaultType}:               _GotoState417Action,
	{_State346, AddToken}:                          _GotoState253Action,
	{_State346, SubToken}:                          _GotoState255Action,
	{_State346, MulToken}:                          _GotoState254Action,
	{_State346, BinaryTypeOpType}:                  _GotoState256Action,
	{_State347, IdentifierToken}:                   _GotoState241Action,
	{_State347, UnderscoreToken}:                   _GotoState242Action,
	{_State347, UnsafeToken}:                       _GotoState54Action,
	{_State347, StructToken}:                       _GotoState50Action,
	{_State347, EnumToken}:                         _GotoState113Action,
	{_State347, TraitToken}:                        _GotoState121Action,
	{_State347, FuncToken}:                         _GotoState240Action,
	{_State347, LparenToken}:                       _GotoState117Action,
	{_State347, LbracketToken}:                     _GotoState42Action,
	{_State347, DotToken}:                          _GotoState112Action,
	{_State347, QuestionToken}:                     _GotoState119Action,
	{_State347, ExclaimToken}:                      _GotoState114Action,
	{_State347, TildeTildeToken}:                   _GotoState120Action,
	{_State347, BitNegToken}:                       _GotoState111Action,
	{_State347, BitAndToken}:                       _GotoState110Action,
	{_State347, ParseErrorToken}:                   _GotoState118Action,
	{_State347, UnsafeStatementType}:               _GotoState250Action,
	{_State347, InitializableTypeExprType}:         _GotoState130Action,
	{_State347, SliceTypeExprType}:                 _GotoState103Action,
	{_State347, ArrayTypeExprType}:                 _GotoState60Action,
	{_State347, MapTypeExprType}:                   _GotoState90Action,
	{_State347, AtomTypeExprType}:                  _GotoState123Action,
	{_State347, NamedTypeExprType}:                 _GotoState131Action,
	{_State347, InferredTypeExprType}:              _GotoState129Action,
	{_State347, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State347, ReturnableTypeExprType}:            _GotoState135Action,
	{_State347, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State347, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State347, TypeExprType}:                      _GotoState249Action,
	{_State347, BinaryTypeExprType}:                _GotoState124Action,
	{_State347, FieldDefType}:                      _GotoState418Action,
	{_State347, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State347, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State347, TraitTypeExprType}:                 _GotoState136Action,
	{_State347, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State347, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State347, FuncTypeExprType}:                  _GotoState126Action,
	{_State347, MethodSignatureType}:               _GotoState246Action,
	{_State351, IdentifierToken}:                   _GotoState241Action,
	{_State351, UnderscoreToken}:                   _GotoState242Action,
	{_State351, UnsafeToken}:                       _GotoState54Action,
	{_State351, StructToken}:                       _GotoState50Action,
	{_State351, EnumToken}:                         _GotoState113Action,
	{_State351, TraitToken}:                        _GotoState121Action,
	{_State351, FuncToken}:                         _GotoState240Action,
	{_State351, LparenToken}:                       _GotoState117Action,
	{_State351, LbracketToken}:                     _GotoState42Action,
	{_State351, DotToken}:                          _GotoState112Action,
	{_State351, QuestionToken}:                     _GotoState119Action,
	{_State351, ExclaimToken}:                      _GotoState114Action,
	{_State351, TildeTildeToken}:                   _GotoState120Action,
	{_State351, BitNegToken}:                       _GotoState111Action,
	{_State351, BitAndToken}:                       _GotoState110Action,
	{_State351, ParseErrorToken}:                   _GotoState118Action,
	{_State351, UnsafeStatementType}:               _GotoState250Action,
	{_State351, InitializableTypeExprType}:         _GotoState130Action,
	{_State351, SliceTypeExprType}:                 _GotoState103Action,
	{_State351, ArrayTypeExprType}:                 _GotoState60Action,
	{_State351, MapTypeExprType}:                   _GotoState90Action,
	{_State351, AtomTypeExprType}:                  _GotoState123Action,
	{_State351, NamedTypeExprType}:                 _GotoState131Action,
	{_State351, InferredTypeExprType}:              _GotoState129Action,
	{_State351, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State351, ReturnableTypeExprType}:            _GotoState135Action,
	{_State351, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State351, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State351, TypeExprType}:                      _GotoState249Action,
	{_State351, BinaryTypeExprType}:                _GotoState124Action,
	{_State351, FieldDefType}:                      _GotoState419Action,
	{_State351, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State351, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State351, TraitTypeExprType}:                 _GotoState136Action,
	{_State351, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State351, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State351, FuncTypeExprType}:                  _GotoState126Action,
	{_State351, MethodSignatureType}:               _GotoState246Action,
	{_State352, IdentifierToken}:                   _GotoState241Action,
	{_State352, UnderscoreToken}:                   _GotoState242Action,
	{_State352, UnsafeToken}:                       _GotoState54Action,
	{_State352, StructToken}:                       _GotoState50Action,
	{_State352, EnumToken}:                         _GotoState113Action,
	{_State352, TraitToken}:                        _GotoState121Action,
	{_State352, FuncToken}:                         _GotoState240Action,
	{_State352, LparenToken}:                       _GotoState117Action,
	{_State352, LbracketToken}:                     _GotoState42Action,
	{_State352, DotToken}:                          _GotoState112Action,
	{_State352, QuestionToken}:                     _GotoState119Action,
	{_State352, ExclaimToken}:                      _GotoState114Action,
	{_State352, TildeTildeToken}:                   _GotoState120Action,
	{_State352, BitNegToken}:                       _GotoState111Action,
	{_State352, BitAndToken}:                       _GotoState110Action,
	{_State352, ParseErrorToken}:                   _GotoState118Action,
	{_State352, UnsafeStatementType}:               _GotoState250Action,
	{_State352, InitializableTypeExprType}:         _GotoState130Action,
	{_State352, SliceTypeExprType}:                 _GotoState103Action,
	{_State352, ArrayTypeExprType}:                 _GotoState60Action,
	{_State352, MapTypeExprType}:                   _GotoState90Action,
	{_State352, AtomTypeExprType}:                  _GotoState123Action,
	{_State352, NamedTypeExprType}:                 _GotoState131Action,
	{_State352, InferredTypeExprType}:              _GotoState129Action,
	{_State352, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State352, ReturnableTypeExprType}:            _GotoState135Action,
	{_State352, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State352, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State352, TypeExprType}:                      _GotoState249Action,
	{_State352, BinaryTypeExprType}:                _GotoState124Action,
	{_State352, FieldDefType}:                      _GotoState420Action,
	{_State352, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State352, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State352, TraitTypeExprType}:                 _GotoState136Action,
	{_State352, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State352, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State352, FuncTypeExprType}:                  _GotoState126Action,
	{_State352, MethodSignatureType}:               _GotoState246Action,
	{_State353, DefaultToken}:                      _GotoState421Action,
	{_State355, RparenToken}:                       _GotoState422Action,
	{_State358, IdentifierToken}:                   _GotoState116Action,
	{_State358, UnderscoreToken}:                   _GotoState122Action,
	{_State358, StructToken}:                       _GotoState50Action,
	{_State358, EnumToken}:                         _GotoState113Action,
	{_State358, TraitToken}:                        _GotoState121Action,
	{_State358, FuncToken}:                         _GotoState115Action,
	{_State358, LparenToken}:                       _GotoState117Action,
	{_State358, LbracketToken}:                     _GotoState42Action,
	{_State358, DotToken}:                          _GotoState112Action,
	{_State358, QuestionToken}:                     _GotoState119Action,
	{_State358, ExclaimToken}:                      _GotoState114Action,
	{_State358, TildeTildeToken}:                   _GotoState120Action,
	{_State358, BitNegToken}:                       _GotoState111Action,
	{_State358, BitAndToken}:                       _GotoState110Action,
	{_State358, ParseErrorToken}:                   _GotoState118Action,
	{_State358, InitializableTypeExprType}:         _GotoState130Action,
	{_State358, SliceTypeExprType}:                 _GotoState103Action,
	{_State358, ArrayTypeExprType}:                 _GotoState60Action,
	{_State358, MapTypeExprType}:                   _GotoState90Action,
	{_State358, AtomTypeExprType}:                  _GotoState123Action,
	{_State358, NamedTypeExprType}:                 _GotoState131Action,
	{_State358, InferredTypeExprType}:              _GotoState129Action,
	{_State358, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State358, ReturnableTypeExprType}:            _GotoState135Action,
	{_State358, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State358, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State358, TypeExprType}:                      _GotoState423Action,
	{_State358, BinaryTypeExprType}:                _GotoState124Action,
	{_State358, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State358, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State358, TraitTypeExprType}:                 _GotoState136Action,
	{_State358, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State358, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State358, FuncTypeExprType}:                  _GotoState126Action,
	{_State360, RbracketToken}:                     _GotoState424Action,
	{_State361, CommaToken}:                        _GotoState425Action,
	{_State362, IdentifierToken}:                   _GotoState260Action,
	{_State362, UnderscoreToken}:                   _GotoState261Action,
	{_State362, ProperParameterDefType}:            _GotoState263Action,
	{_State362, ParameterDefType}:                  _GotoState284Action,
	{_State362, ProperParameterDefsType}:           _GotoState286Action,
	{_State362, ParameterDefsType}:                 _GotoState426Action,
	{_State363, IdentifierToken}:                   _GotoState116Action,
	{_State363, UnderscoreToken}:                   _GotoState122Action,
	{_State363, StructToken}:                       _GotoState50Action,
	{_State363, EnumToken}:                         _GotoState113Action,
	{_State363, TraitToken}:                        _GotoState121Action,
	{_State363, FuncToken}:                         _GotoState115Action,
	{_State363, LparenToken}:                       _GotoState117Action,
	{_State363, LbracketToken}:                     _GotoState42Action,
	{_State363, DotToken}:                          _GotoState112Action,
	{_State363, QuestionToken}:                     _GotoState119Action,
	{_State363, ExclaimToken}:                      _GotoState114Action,
	{_State363, TildeTildeToken}:                   _GotoState120Action,
	{_State363, BitNegToken}:                       _GotoState111Action,
	{_State363, BitAndToken}:                       _GotoState110Action,
	{_State363, ParseErrorToken}:                   _GotoState118Action,
	{_State363, InitializableTypeExprType}:         _GotoState130Action,
	{_State363, SliceTypeExprType}:                 _GotoState103Action,
	{_State363, ArrayTypeExprType}:                 _GotoState60Action,
	{_State363, MapTypeExprType}:                   _GotoState90Action,
	{_State363, AtomTypeExprType}:                  _GotoState123Action,
	{_State363, NamedTypeExprType}:                 _GotoState131Action,
	{_State363, InferredTypeExprType}:              _GotoState129Action,
	{_State363, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State363, ReturnableTypeExprType}:            _GotoState135Action,
	{_State363, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State363, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State363, TypeExprType}:                      _GotoState427Action,
	{_State363, BinaryTypeExprType}:                _GotoState124Action,
	{_State363, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State363, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State363, TraitTypeExprType}:                 _GotoState136Action,
	{_State363, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State363, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State363, FuncTypeExprType}:                  _GotoState126Action,
	{_State364, AddToken}:                          _GotoState253Action,
	{_State364, SubToken}:                          _GotoState255Action,
	{_State364, MulToken}:                          _GotoState254Action,
	{_State364, BinaryTypeOpType}:                  _GotoState256Action,
	{_State365, IdentifierToken}:                   _GotoState116Action,
	{_State365, UnderscoreToken}:                   _GotoState122Action,
	{_State365, StructToken}:                       _GotoState50Action,
	{_State365, EnumToken}:                         _GotoState113Action,
	{_State365, TraitToken}:                        _GotoState121Action,
	{_State365, FuncToken}:                         _GotoState115Action,
	{_State365, LparenToken}:                       _GotoState117Action,
	{_State365, LbracketToken}:                     _GotoState42Action,
	{_State365, DotToken}:                          _GotoState112Action,
	{_State365, QuestionToken}:                     _GotoState119Action,
	{_State365, ExclaimToken}:                      _GotoState114Action,
	{_State365, TildeTildeToken}:                   _GotoState120Action,
	{_State365, BitNegToken}:                       _GotoState111Action,
	{_State365, BitAndToken}:                       _GotoState110Action,
	{_State365, ParseErrorToken}:                   _GotoState118Action,
	{_State365, InitializableTypeExprType}:         _GotoState130Action,
	{_State365, SliceTypeExprType}:                 _GotoState103Action,
	{_State365, ArrayTypeExprType}:                 _GotoState60Action,
	{_State365, MapTypeExprType}:                   _GotoState90Action,
	{_State365, AtomTypeExprType}:                  _GotoState123Action,
	{_State365, NamedTypeExprType}:                 _GotoState131Action,
	{_State365, InferredTypeExprType}:              _GotoState129Action,
	{_State365, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State365, ReturnableTypeExprType}:            _GotoState135Action,
	{_State365, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State365, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State365, TypeExprType}:                      _GotoState428Action,
	{_State365, BinaryTypeExprType}:                _GotoState124Action,
	{_State365, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State365, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State365, TraitTypeExprType}:                 _GotoState136Action,
	{_State365, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State365, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State365, FuncTypeExprType}:                  _GotoState126Action,
	{_State366, AddToken}:                          _GotoState253Action,
	{_State366, SubToken}:                          _GotoState255Action,
	{_State366, MulToken}:                          _GotoState254Action,
	{_State366, BinaryTypeOpType}:                  _GotoState256Action,
	{_State367, IdentifierToken}:                   _GotoState429Action,
	{_State370, AddToken}:                          _GotoState253Action,
	{_State370, SubToken}:                          _GotoState255Action,
	{_State370, MulToken}:                          _GotoState254Action,
	{_State370, BinaryTypeOpType}:                  _GotoState256Action,
	{_State371, ImplementsToken}:                   _GotoState430Action,
	{_State371, AddToken}:                          _GotoState253Action,
	{_State371, SubToken}:                          _GotoState255Action,
	{_State371, MulToken}:                          _GotoState254Action,
	{_State371, BinaryTypeOpType}:                  _GotoState256Action,
	{_State372, IdentifierToken}:                   _GotoState146Action,
	{_State372, UnderscoreToken}:                   _GotoState148Action,
	{_State372, LparenToken}:                       _GotoState147Action,
	{_State372, VarPatternType}:                    _GotoState431Action,
	{_State372, TuplePatternType}:                  _GotoState149Action,
	{_State373, IdentifierToken}:                   _GotoState272Action,
	{_State373, UnderscoreToken}:                   _GotoState148Action,
	{_State373, LparenToken}:                       _GotoState147Action,
	{_State373, EllipsisToken}:                     _GotoState271Action,
	{_State373, VarPatternType}:                    _GotoState275Action,
	{_State373, TuplePatternType}:                  _GotoState149Action,
	{_State373, FieldVarPatternType}:               _GotoState432Action,
	{_State376, LparenToken}:                       _GotoState147Action,
	{_State376, TuplePatternType}:                  _GotoState433Action,
	{_State379, IdentifierToken}:                   _GotoState116Action,
	{_State379, UnderscoreToken}:                   _GotoState122Action,
	{_State379, StructToken}:                       _GotoState50Action,
	{_State379, EnumToken}:                         _GotoState113Action,
	{_State379, TraitToken}:                        _GotoState121Action,
	{_State379, FuncToken}:                         _GotoState115Action,
	{_State379, LparenToken}:                       _GotoState117Action,
	{_State379, LbracketToken}:                     _GotoState42Action,
	{_State379, DotToken}:                          _GotoState112Action,
	{_State379, QuestionToken}:                     _GotoState119Action,
	{_State379, ExclaimToken}:                      _GotoState114Action,
	{_State379, TildeTildeToken}:                   _GotoState120Action,
	{_State379, BitNegToken}:                       _GotoState111Action,
	{_State379, BitAndToken}:                       _GotoState110Action,
	{_State379, ParseErrorToken}:                   _GotoState118Action,
	{_State379, InitializableTypeExprType}:         _GotoState130Action,
	{_State379, SliceTypeExprType}:                 _GotoState103Action,
	{_State379, ArrayTypeExprType}:                 _GotoState60Action,
	{_State379, MapTypeExprType}:                   _GotoState90Action,
	{_State379, AtomTypeExprType}:                  _GotoState123Action,
	{_State379, NamedTypeExprType}:                 _GotoState131Action,
	{_State379, InferredTypeExprType}:              _GotoState129Action,
	{_State379, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State379, ReturnableTypeExprType}:            _GotoState435Action,
	{_State379, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State379, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State379, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State379, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State379, TraitTypeExprType}:                 _GotoState136Action,
	{_State379, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State379, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State379, ReturnTypeType}:                    _GotoState434Action,
	{_State379, FuncTypeExprType}:                  _GotoState126Action,
	{_State380, IdentifierToken}:                   _GotoState260Action,
	{_State380, UnderscoreToken}:                   _GotoState261Action,
	{_State380, ProperParameterDefType}:            _GotoState263Action,
	{_State380, ParameterDefType}:                  _GotoState436Action,
	{_State382, StringLiteralToken}:                _GotoState164Action,
	{_State382, IdentifierToken}:                   _GotoState162Action,
	{_State382, UnderscoreToken}:                   _GotoState165Action,
	{_State382, DotToken}:                          _GotoState161Action,
	{_State382, ImportClauseType}:                  _GotoState437Action,
	{_State383, StringLiteralToken}:                _GotoState164Action,
	{_State383, IdentifierToken}:                   _GotoState162Action,
	{_State383, UnderscoreToken}:                   _GotoState165Action,
	{_State383, DotToken}:                          _GotoState161Action,
	{_State383, ImportClauseType}:                  _GotoState438Action,
	{_State384, RbracketToken}:                     _GotoState439Action,
	{_State384, AddToken}:                          _GotoState253Action,
	{_State384, SubToken}:                          _GotoState255Action,
	{_State384, MulToken}:                          _GotoState254Action,
	{_State384, BinaryTypeOpType}:                  _GotoState256Action,
	{_State385, RbracketToken}:                     _GotoState440Action,
	{_State391, IdentifierToken}:                   _GotoState241Action,
	{_State391, UnderscoreToken}:                   _GotoState242Action,
	{_State391, UnsafeToken}:                       _GotoState54Action,
	{_State391, StructToken}:                       _GotoState50Action,
	{_State391, EnumToken}:                         _GotoState113Action,
	{_State391, TraitToken}:                        _GotoState121Action,
	{_State391, FuncToken}:                         _GotoState240Action,
	{_State391, LparenToken}:                       _GotoState117Action,
	{_State391, LbracketToken}:                     _GotoState42Action,
	{_State391, DotToken}:                          _GotoState112Action,
	{_State391, QuestionToken}:                     _GotoState119Action,
	{_State391, ExclaimToken}:                      _GotoState114Action,
	{_State391, TildeTildeToken}:                   _GotoState120Action,
	{_State391, BitNegToken}:                       _GotoState111Action,
	{_State391, BitAndToken}:                       _GotoState110Action,
	{_State391, ParseErrorToken}:                   _GotoState118Action,
	{_State391, UnsafeStatementType}:               _GotoState250Action,
	{_State391, InitializableTypeExprType}:         _GotoState130Action,
	{_State391, SliceTypeExprType}:                 _GotoState103Action,
	{_State391, ArrayTypeExprType}:                 _GotoState60Action,
	{_State391, MapTypeExprType}:                   _GotoState90Action,
	{_State391, AtomTypeExprType}:                  _GotoState123Action,
	{_State391, NamedTypeExprType}:                 _GotoState131Action,
	{_State391, InferredTypeExprType}:              _GotoState129Action,
	{_State391, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State391, ReturnableTypeExprType}:            _GotoState135Action,
	{_State391, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State391, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State391, TypeExprType}:                      _GotoState249Action,
	{_State391, BinaryTypeExprType}:                _GotoState124Action,
	{_State391, FieldDefType}:                      _GotoState441Action,
	{_State391, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State391, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State391, TraitTypeExprType}:                 _GotoState136Action,
	{_State391, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State391, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State391, FuncTypeExprType}:                  _GotoState126Action,
	{_State391, MethodSignatureType}:               _GotoState246Action,
	{_State392, IdentifierToken}:                   _GotoState241Action,
	{_State392, UnderscoreToken}:                   _GotoState242Action,
	{_State392, UnsafeToken}:                       _GotoState54Action,
	{_State392, StructToken}:                       _GotoState50Action,
	{_State392, EnumToken}:                         _GotoState113Action,
	{_State392, TraitToken}:                        _GotoState121Action,
	{_State392, FuncToken}:                         _GotoState240Action,
	{_State392, LparenToken}:                       _GotoState117Action,
	{_State392, LbracketToken}:                     _GotoState42Action,
	{_State392, DotToken}:                          _GotoState112Action,
	{_State392, QuestionToken}:                     _GotoState119Action,
	{_State392, ExclaimToken}:                      _GotoState114Action,
	{_State392, TildeTildeToken}:                   _GotoState120Action,
	{_State392, BitNegToken}:                       _GotoState111Action,
	{_State392, BitAndToken}:                       _GotoState110Action,
	{_State392, ParseErrorToken}:                   _GotoState118Action,
	{_State392, UnsafeStatementType}:               _GotoState250Action,
	{_State392, InitializableTypeExprType}:         _GotoState130Action,
	{_State392, SliceTypeExprType}:                 _GotoState103Action,
	{_State392, ArrayTypeExprType}:                 _GotoState60Action,
	{_State392, MapTypeExprType}:                   _GotoState90Action,
	{_State392, AtomTypeExprType}:                  _GotoState123Action,
	{_State392, NamedTypeExprType}:                 _GotoState131Action,
	{_State392, InferredTypeExprType}:              _GotoState129Action,
	{_State392, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State392, ReturnableTypeExprType}:            _GotoState135Action,
	{_State392, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State392, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State392, TypeExprType}:                      _GotoState249Action,
	{_State392, BinaryTypeExprType}:                _GotoState124Action,
	{_State392, FieldDefType}:                      _GotoState442Action,
	{_State392, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State392, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State392, TraitTypeExprType}:                 _GotoState136Action,
	{_State392, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State392, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State392, FuncTypeExprType}:                  _GotoState126Action,
	{_State392, MethodSignatureType}:               _GotoState246Action,
	{_State393, StringLiteralToken}:                _GotoState443Action,
	{_State394, IdentifierToken}:                   _GotoState116Action,
	{_State394, UnderscoreToken}:                   _GotoState122Action,
	{_State394, StructToken}:                       _GotoState50Action,
	{_State394, EnumToken}:                         _GotoState113Action,
	{_State394, TraitToken}:                        _GotoState121Action,
	{_State394, FuncToken}:                         _GotoState115Action,
	{_State394, LparenToken}:                       _GotoState117Action,
	{_State394, LbracketToken}:                     _GotoState42Action,
	{_State394, DotToken}:                          _GotoState112Action,
	{_State394, QuestionToken}:                     _GotoState119Action,
	{_State394, ExclaimToken}:                      _GotoState114Action,
	{_State394, TildeTildeToken}:                   _GotoState120Action,
	{_State394, BitNegToken}:                       _GotoState111Action,
	{_State394, BitAndToken}:                       _GotoState110Action,
	{_State394, ParseErrorToken}:                   _GotoState118Action,
	{_State394, InitializableTypeExprType}:         _GotoState130Action,
	{_State394, SliceTypeExprType}:                 _GotoState103Action,
	{_State394, ArrayTypeExprType}:                 _GotoState60Action,
	{_State394, MapTypeExprType}:                   _GotoState90Action,
	{_State394, AtomTypeExprType}:                  _GotoState123Action,
	{_State394, NamedTypeExprType}:                 _GotoState131Action,
	{_State394, InferredTypeExprType}:              _GotoState129Action,
	{_State394, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State394, ReturnableTypeExprType}:            _GotoState135Action,
	{_State394, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State394, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State394, TypeExprType}:                      _GotoState444Action,
	{_State394, BinaryTypeExprType}:                _GotoState124Action,
	{_State394, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State394, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State394, TraitTypeExprType}:                 _GotoState136Action,
	{_State394, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State394, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State394, FuncTypeExprType}:                  _GotoState126Action,
	{_State397, RparenToken}:                       _GotoState445Action,
	{_State399, IntegerLiteralToken}:               _GotoState40Action,
	{_State399, FloatLiteralToken}:                 _GotoState35Action,
	{_State399, RuneLiteralToken}:                  _GotoState48Action,
	{_State399, StringLiteralToken}:                _GotoState49Action,
	{_State399, IdentifierToken}:                   _GotoState38Action,
	{_State399, UnderscoreToken}:                   _GotoState53Action,
	{_State399, TrueToken}:                         _GotoState52Action,
	{_State399, FalseToken}:                        _GotoState34Action,
	{_State399, StructToken}:                       _GotoState50Action,
	{_State399, FuncToken}:                         _GotoState36Action,
	{_State399, VarToken}:                          _GotoState15Action,
	{_State399, LetToken}:                          _GotoState12Action,
	{_State399, NotToken}:                          _GotoState45Action,
	{_State399, LabelDeclToken}:                    _GotoState41Action,
	{_State399, LparenToken}:                       _GotoState43Action,
	{_State399, LbracketToken}:                     _GotoState42Action,
	{_State399, SubToken}:                          _GotoState51Action,
	{_State399, MulToken}:                          _GotoState44Action,
	{_State399, BitNegToken}:                       _GotoState27Action,
	{_State399, BitAndToken}:                       _GotoState26Action,
	{_State399, GreaterToken}:                      _GotoState37Action,
	{_State399, ParseErrorToken}:                   _GotoState46Action,
	{_State399, VarDeclPatternType}:                _GotoState107Action,
	{_State399, VarOrLetType}:                      _GotoState24Action,
	{_State399, OptionalLabelDeclType}:             _GotoState156Action,
	{_State399, SequenceExprType}:                  _GotoState446Action,
	{_State399, CallExprType}:                      _GotoState71Action,
	{_State399, AtomExprType}:                      _GotoState63Action,
	{_State399, ParseErrorExprType}:                _GotoState95Action,
	{_State399, LiteralExprType}:                   _GotoState88Action,
	{_State399, NamedExprType}:                     _GotoState92Action,
	{_State399, BlockExprType}:                     _GotoState70Action,
	{_State399, InitializeExprType}:                _GotoState85Action,
	{_State399, ImplicitStructExprType}:            _GotoState81Action,
	{_State399, AccessibleExprType}:                _GotoState108Action,
	{_State399, AccessExprType}:                    _GotoState55Action,
	{_State399, IndexExprType}:                     _GotoState83Action,
	{_State399, PostfixableExprType}:               _GotoState97Action,
	{_State399, PostfixUnaryExprType}:              _GotoState96Action,
	{_State399, PrefixableExprType}:                _GotoState100Action,
	{_State399, PrefixUnaryExprType}:               _GotoState98Action,
	{_State399, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State399, MulExprType}:                       _GotoState91Action,
	{_State399, BinaryMulExprType}:                 _GotoState67Action,
	{_State399, AddExprType}:                       _GotoState57Action,
	{_State399, BinaryAddExprType}:                 _GotoState64Action,
	{_State399, CmpExprType}:                       _GotoState74Action,
	{_State399, BinaryCmpExprType}:                 _GotoState66Action,
	{_State399, AndExprType}:                       _GotoState58Action,
	{_State399, BinaryAndExprType}:                 _GotoState65Action,
	{_State399, OrExprType}:                        _GotoState94Action,
	{_State399, BinaryOrExprType}:                  _GotoState69Action,
	{_State399, InitializableTypeExprType}:         _GotoState84Action,
	{_State399, SliceTypeExprType}:                 _GotoState103Action,
	{_State399, ArrayTypeExprType}:                 _GotoState60Action,
	{_State399, MapTypeExprType}:                   _GotoState90Action,
	{_State399, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State399, AnonymousFuncExprType}:             _GotoState59Action,
	{_State400, IntegerLiteralToken}:               _GotoState40Action,
	{_State400, FloatLiteralToken}:                 _GotoState35Action,
	{_State400, RuneLiteralToken}:                  _GotoState48Action,
	{_State400, StringLiteralToken}:                _GotoState49Action,
	{_State400, IdentifierToken}:                   _GotoState38Action,
	{_State400, UnderscoreToken}:                   _GotoState53Action,
	{_State400, TrueToken}:                         _GotoState52Action,
	{_State400, FalseToken}:                        _GotoState34Action,
	{_State400, StructToken}:                       _GotoState50Action,
	{_State400, FuncToken}:                         _GotoState36Action,
	{_State400, VarToken}:                          _GotoState15Action,
	{_State400, LetToken}:                          _GotoState12Action,
	{_State400, NotToken}:                          _GotoState45Action,
	{_State400, LabelDeclToken}:                    _GotoState41Action,
	{_State400, LparenToken}:                       _GotoState43Action,
	{_State400, LbracketToken}:                     _GotoState42Action,
	{_State400, SubToken}:                          _GotoState51Action,
	{_State400, MulToken}:                          _GotoState44Action,
	{_State400, BitNegToken}:                       _GotoState27Action,
	{_State400, BitAndToken}:                       _GotoState26Action,
	{_State400, GreaterToken}:                      _GotoState37Action,
	{_State400, ParseErrorToken}:                   _GotoState46Action,
	{_State400, VarDeclPatternType}:                _GotoState107Action,
	{_State400, VarOrLetType}:                      _GotoState24Action,
	{_State400, OptionalLabelDeclType}:             _GotoState156Action,
	{_State400, SequenceExprType}:                  _GotoState447Action,
	{_State400, CallExprType}:                      _GotoState71Action,
	{_State400, AtomExprType}:                      _GotoState63Action,
	{_State400, ParseErrorExprType}:                _GotoState95Action,
	{_State400, LiteralExprType}:                   _GotoState88Action,
	{_State400, NamedExprType}:                     _GotoState92Action,
	{_State400, BlockExprType}:                     _GotoState70Action,
	{_State400, InitializeExprType}:                _GotoState85Action,
	{_State400, ImplicitStructExprType}:            _GotoState81Action,
	{_State400, AccessibleExprType}:                _GotoState108Action,
	{_State400, AccessExprType}:                    _GotoState55Action,
	{_State400, IndexExprType}:                     _GotoState83Action,
	{_State400, PostfixableExprType}:               _GotoState97Action,
	{_State400, PostfixUnaryExprType}:              _GotoState96Action,
	{_State400, PrefixableExprType}:                _GotoState100Action,
	{_State400, PrefixUnaryExprType}:               _GotoState98Action,
	{_State400, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State400, MulExprType}:                       _GotoState91Action,
	{_State400, BinaryMulExprType}:                 _GotoState67Action,
	{_State400, AddExprType}:                       _GotoState57Action,
	{_State400, BinaryAddExprType}:                 _GotoState64Action,
	{_State400, CmpExprType}:                       _GotoState74Action,
	{_State400, BinaryCmpExprType}:                 _GotoState66Action,
	{_State400, AndExprType}:                       _GotoState58Action,
	{_State400, BinaryAndExprType}:                 _GotoState65Action,
	{_State400, OrExprType}:                        _GotoState94Action,
	{_State400, BinaryOrExprType}:                  _GotoState69Action,
	{_State400, InitializableTypeExprType}:         _GotoState84Action,
	{_State400, SliceTypeExprType}:                 _GotoState103Action,
	{_State400, ArrayTypeExprType}:                 _GotoState60Action,
	{_State400, MapTypeExprType}:                   _GotoState90Action,
	{_State400, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State400, AnonymousFuncExprType}:             _GotoState59Action,
	{_State401, IntegerLiteralToken}:               _GotoState40Action,
	{_State401, FloatLiteralToken}:                 _GotoState35Action,
	{_State401, RuneLiteralToken}:                  _GotoState48Action,
	{_State401, StringLiteralToken}:                _GotoState49Action,
	{_State401, IdentifierToken}:                   _GotoState38Action,
	{_State401, UnderscoreToken}:                   _GotoState53Action,
	{_State401, TrueToken}:                         _GotoState52Action,
	{_State401, FalseToken}:                        _GotoState34Action,
	{_State401, StructToken}:                       _GotoState50Action,
	{_State401, FuncToken}:                         _GotoState36Action,
	{_State401, VarToken}:                          _GotoState15Action,
	{_State401, LetToken}:                          _GotoState12Action,
	{_State401, NotToken}:                          _GotoState45Action,
	{_State401, LabelDeclToken}:                    _GotoState41Action,
	{_State401, LparenToken}:                       _GotoState43Action,
	{_State401, LbracketToken}:                     _GotoState42Action,
	{_State401, SubToken}:                          _GotoState51Action,
	{_State401, MulToken}:                          _GotoState44Action,
	{_State401, BitNegToken}:                       _GotoState27Action,
	{_State401, BitAndToken}:                       _GotoState26Action,
	{_State401, GreaterToken}:                      _GotoState37Action,
	{_State401, ParseErrorToken}:                   _GotoState46Action,
	{_State401, VarDeclPatternType}:                _GotoState107Action,
	{_State401, VarOrLetType}:                      _GotoState24Action,
	{_State401, OptionalLabelDeclType}:             _GotoState156Action,
	{_State401, SequenceExprType}:                  _GotoState448Action,
	{_State401, CallExprType}:                      _GotoState71Action,
	{_State401, AtomExprType}:                      _GotoState63Action,
	{_State401, ParseErrorExprType}:                _GotoState95Action,
	{_State401, LiteralExprType}:                   _GotoState88Action,
	{_State401, NamedExprType}:                     _GotoState92Action,
	{_State401, BlockExprType}:                     _GotoState70Action,
	{_State401, InitializeExprType}:                _GotoState85Action,
	{_State401, ImplicitStructExprType}:            _GotoState81Action,
	{_State401, AccessibleExprType}:                _GotoState108Action,
	{_State401, AccessExprType}:                    _GotoState55Action,
	{_State401, IndexExprType}:                     _GotoState83Action,
	{_State401, PostfixableExprType}:               _GotoState97Action,
	{_State401, PostfixUnaryExprType}:              _GotoState96Action,
	{_State401, PrefixableExprType}:                _GotoState100Action,
	{_State401, PrefixUnaryExprType}:               _GotoState98Action,
	{_State401, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State401, MulExprType}:                       _GotoState91Action,
	{_State401, BinaryMulExprType}:                 _GotoState67Action,
	{_State401, AddExprType}:                       _GotoState57Action,
	{_State401, BinaryAddExprType}:                 _GotoState64Action,
	{_State401, CmpExprType}:                       _GotoState74Action,
	{_State401, BinaryCmpExprType}:                 _GotoState66Action,
	{_State401, AndExprType}:                       _GotoState58Action,
	{_State401, BinaryAndExprType}:                 _GotoState65Action,
	{_State401, OrExprType}:                        _GotoState94Action,
	{_State401, BinaryOrExprType}:                  _GotoState69Action,
	{_State401, InitializableTypeExprType}:         _GotoState84Action,
	{_State401, SliceTypeExprType}:                 _GotoState103Action,
	{_State401, ArrayTypeExprType}:                 _GotoState60Action,
	{_State401, MapTypeExprType}:                   _GotoState90Action,
	{_State401, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State401, AnonymousFuncExprType}:             _GotoState59Action,
	{_State402, IntegerLiteralToken}:               _GotoState40Action,
	{_State402, FloatLiteralToken}:                 _GotoState35Action,
	{_State402, RuneLiteralToken}:                  _GotoState48Action,
	{_State402, StringLiteralToken}:                _GotoState49Action,
	{_State402, IdentifierToken}:                   _GotoState38Action,
	{_State402, UnderscoreToken}:                   _GotoState53Action,
	{_State402, TrueToken}:                         _GotoState52Action,
	{_State402, FalseToken}:                        _GotoState34Action,
	{_State402, StructToken}:                       _GotoState50Action,
	{_State402, FuncToken}:                         _GotoState36Action,
	{_State402, VarToken}:                          _GotoState15Action,
	{_State402, LetToken}:                          _GotoState12Action,
	{_State402, NotToken}:                          _GotoState45Action,
	{_State402, LabelDeclToken}:                    _GotoState41Action,
	{_State402, LparenToken}:                       _GotoState43Action,
	{_State402, LbracketToken}:                     _GotoState42Action,
	{_State402, SubToken}:                          _GotoState51Action,
	{_State402, MulToken}:                          _GotoState44Action,
	{_State402, BitNegToken}:                       _GotoState27Action,
	{_State402, BitAndToken}:                       _GotoState26Action,
	{_State402, GreaterToken}:                      _GotoState37Action,
	{_State402, ParseErrorToken}:                   _GotoState46Action,
	{_State402, VarDeclPatternType}:                _GotoState107Action,
	{_State402, VarOrLetType}:                      _GotoState24Action,
	{_State402, OptionalLabelDeclType}:             _GotoState156Action,
	{_State402, SequenceExprType}:                  _GotoState450Action,
	{_State402, OptionalSequenceExprType}:          _GotoState449Action,
	{_State402, CallExprType}:                      _GotoState71Action,
	{_State402, AtomExprType}:                      _GotoState63Action,
	{_State402, ParseErrorExprType}:                _GotoState95Action,
	{_State402, LiteralExprType}:                   _GotoState88Action,
	{_State402, NamedExprType}:                     _GotoState92Action,
	{_State402, BlockExprType}:                     _GotoState70Action,
	{_State402, InitializeExprType}:                _GotoState85Action,
	{_State402, ImplicitStructExprType}:            _GotoState81Action,
	{_State402, AccessibleExprType}:                _GotoState108Action,
	{_State402, AccessExprType}:                    _GotoState55Action,
	{_State402, IndexExprType}:                     _GotoState83Action,
	{_State402, PostfixableExprType}:               _GotoState97Action,
	{_State402, PostfixUnaryExprType}:              _GotoState96Action,
	{_State402, PrefixableExprType}:                _GotoState100Action,
	{_State402, PrefixUnaryExprType}:               _GotoState98Action,
	{_State402, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State402, MulExprType}:                       _GotoState91Action,
	{_State402, BinaryMulExprType}:                 _GotoState67Action,
	{_State402, AddExprType}:                       _GotoState57Action,
	{_State402, BinaryAddExprType}:                 _GotoState64Action,
	{_State402, CmpExprType}:                       _GotoState74Action,
	{_State402, BinaryCmpExprType}:                 _GotoState66Action,
	{_State402, AndExprType}:                       _GotoState58Action,
	{_State402, BinaryAndExprType}:                 _GotoState65Action,
	{_State402, OrExprType}:                        _GotoState94Action,
	{_State402, BinaryOrExprType}:                  _GotoState69Action,
	{_State402, InitializableTypeExprType}:         _GotoState84Action,
	{_State402, SliceTypeExprType}:                 _GotoState103Action,
	{_State402, ArrayTypeExprType}:                 _GotoState60Action,
	{_State402, MapTypeExprType}:                   _GotoState90Action,
	{_State402, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State402, AnonymousFuncExprType}:             _GotoState59Action,
	{_State403, LbraceToken}:                       _GotoState11Action,
	{_State403, StatementBlockType}:                _GotoState451Action,
	{_State404, CommaToken}:                        _GotoState281Action,
	{_State404, AssignToken}:                       _GotoState452Action,
	{_State405, ElseToken}:                         _GotoState453Action,
	{_State408, IdentifierToken}:                   _GotoState241Action,
	{_State408, UnderscoreToken}:                   _GotoState242Action,
	{_State408, UnsafeToken}:                       _GotoState54Action,
	{_State408, StructToken}:                       _GotoState50Action,
	{_State408, EnumToken}:                         _GotoState113Action,
	{_State408, TraitToken}:                        _GotoState121Action,
	{_State408, FuncToken}:                         _GotoState240Action,
	{_State408, LparenToken}:                       _GotoState117Action,
	{_State408, LbracketToken}:                     _GotoState42Action,
	{_State408, DotToken}:                          _GotoState112Action,
	{_State408, QuestionToken}:                     _GotoState119Action,
	{_State408, ExclaimToken}:                      _GotoState114Action,
	{_State408, TildeTildeToken}:                   _GotoState120Action,
	{_State408, BitNegToken}:                       _GotoState111Action,
	{_State408, BitAndToken}:                       _GotoState110Action,
	{_State408, ParseErrorToken}:                   _GotoState118Action,
	{_State408, UnsafeStatementType}:               _GotoState250Action,
	{_State408, InitializableTypeExprType}:         _GotoState130Action,
	{_State408, SliceTypeExprType}:                 _GotoState103Action,
	{_State408, ArrayTypeExprType}:                 _GotoState60Action,
	{_State408, MapTypeExprType}:                   _GotoState90Action,
	{_State408, AtomTypeExprType}:                  _GotoState123Action,
	{_State408, NamedTypeExprType}:                 _GotoState131Action,
	{_State408, InferredTypeExprType}:              _GotoState129Action,
	{_State408, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State408, ReturnableTypeExprType}:            _GotoState135Action,
	{_State408, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State408, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State408, TypeExprType}:                      _GotoState249Action,
	{_State408, BinaryTypeExprType}:                _GotoState124Action,
	{_State408, FieldDefType}:                      _GotoState454Action,
	{_State408, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State408, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State408, TraitTypeExprType}:                 _GotoState136Action,
	{_State408, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State408, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State408, FuncTypeExprType}:                  _GotoState126Action,
	{_State408, MethodSignatureType}:               _GotoState246Action,
	{_State409, IdentifierToken}:                   _GotoState241Action,
	{_State409, UnderscoreToken}:                   _GotoState242Action,
	{_State409, UnsafeToken}:                       _GotoState54Action,
	{_State409, StructToken}:                       _GotoState50Action,
	{_State409, EnumToken}:                         _GotoState113Action,
	{_State409, TraitToken}:                        _GotoState121Action,
	{_State409, FuncToken}:                         _GotoState240Action,
	{_State409, LparenToken}:                       _GotoState117Action,
	{_State409, LbracketToken}:                     _GotoState42Action,
	{_State409, DotToken}:                          _GotoState112Action,
	{_State409, QuestionToken}:                     _GotoState119Action,
	{_State409, ExclaimToken}:                      _GotoState114Action,
	{_State409, TildeTildeToken}:                   _GotoState120Action,
	{_State409, BitNegToken}:                       _GotoState111Action,
	{_State409, BitAndToken}:                       _GotoState110Action,
	{_State409, ParseErrorToken}:                   _GotoState118Action,
	{_State409, UnsafeStatementType}:               _GotoState250Action,
	{_State409, InitializableTypeExprType}:         _GotoState130Action,
	{_State409, SliceTypeExprType}:                 _GotoState103Action,
	{_State409, ArrayTypeExprType}:                 _GotoState60Action,
	{_State409, MapTypeExprType}:                   _GotoState90Action,
	{_State409, AtomTypeExprType}:                  _GotoState123Action,
	{_State409, NamedTypeExprType}:                 _GotoState131Action,
	{_State409, InferredTypeExprType}:              _GotoState129Action,
	{_State409, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State409, ReturnableTypeExprType}:            _GotoState135Action,
	{_State409, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State409, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State409, TypeExprType}:                      _GotoState249Action,
	{_State409, BinaryTypeExprType}:                _GotoState124Action,
	{_State409, FieldDefType}:                      _GotoState455Action,
	{_State409, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State409, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State409, TraitTypeExprType}:                 _GotoState136Action,
	{_State409, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State409, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State409, FuncTypeExprType}:                  _GotoState126Action,
	{_State409, MethodSignatureType}:               _GotoState246Action,
	{_State410, IdentifierToken}:                   _GotoState241Action,
	{_State410, UnderscoreToken}:                   _GotoState242Action,
	{_State410, UnsafeToken}:                       _GotoState54Action,
	{_State410, StructToken}:                       _GotoState50Action,
	{_State410, EnumToken}:                         _GotoState113Action,
	{_State410, TraitToken}:                        _GotoState121Action,
	{_State410, FuncToken}:                         _GotoState240Action,
	{_State410, LparenToken}:                       _GotoState117Action,
	{_State410, LbracketToken}:                     _GotoState42Action,
	{_State410, DotToken}:                          _GotoState112Action,
	{_State410, QuestionToken}:                     _GotoState119Action,
	{_State410, ExclaimToken}:                      _GotoState114Action,
	{_State410, TildeTildeToken}:                   _GotoState120Action,
	{_State410, BitNegToken}:                       _GotoState111Action,
	{_State410, BitAndToken}:                       _GotoState110Action,
	{_State410, ParseErrorToken}:                   _GotoState118Action,
	{_State410, UnsafeStatementType}:               _GotoState250Action,
	{_State410, InitializableTypeExprType}:         _GotoState130Action,
	{_State410, SliceTypeExprType}:                 _GotoState103Action,
	{_State410, ArrayTypeExprType}:                 _GotoState60Action,
	{_State410, MapTypeExprType}:                   _GotoState90Action,
	{_State410, AtomTypeExprType}:                  _GotoState123Action,
	{_State410, NamedTypeExprType}:                 _GotoState131Action,
	{_State410, InferredTypeExprType}:              _GotoState129Action,
	{_State410, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State410, ReturnableTypeExprType}:            _GotoState135Action,
	{_State410, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State410, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State410, TypeExprType}:                      _GotoState249Action,
	{_State410, BinaryTypeExprType}:                _GotoState124Action,
	{_State410, FieldDefType}:                      _GotoState456Action,
	{_State410, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State410, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State410, TraitTypeExprType}:                 _GotoState136Action,
	{_State410, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State410, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State410, FuncTypeExprType}:                  _GotoState126Action,
	{_State410, MethodSignatureType}:               _GotoState246Action,
	{_State411, IdentifierToken}:                   _GotoState241Action,
	{_State411, UnderscoreToken}:                   _GotoState242Action,
	{_State411, UnsafeToken}:                       _GotoState54Action,
	{_State411, StructToken}:                       _GotoState50Action,
	{_State411, EnumToken}:                         _GotoState113Action,
	{_State411, TraitToken}:                        _GotoState121Action,
	{_State411, FuncToken}:                         _GotoState240Action,
	{_State411, LparenToken}:                       _GotoState117Action,
	{_State411, LbracketToken}:                     _GotoState42Action,
	{_State411, DotToken}:                          _GotoState112Action,
	{_State411, QuestionToken}:                     _GotoState119Action,
	{_State411, ExclaimToken}:                      _GotoState114Action,
	{_State411, TildeTildeToken}:                   _GotoState120Action,
	{_State411, BitNegToken}:                       _GotoState111Action,
	{_State411, BitAndToken}:                       _GotoState110Action,
	{_State411, ParseErrorToken}:                   _GotoState118Action,
	{_State411, UnsafeStatementType}:               _GotoState250Action,
	{_State411, InitializableTypeExprType}:         _GotoState130Action,
	{_State411, SliceTypeExprType}:                 _GotoState103Action,
	{_State411, ArrayTypeExprType}:                 _GotoState60Action,
	{_State411, MapTypeExprType}:                   _GotoState90Action,
	{_State411, AtomTypeExprType}:                  _GotoState123Action,
	{_State411, NamedTypeExprType}:                 _GotoState131Action,
	{_State411, InferredTypeExprType}:              _GotoState129Action,
	{_State411, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State411, ReturnableTypeExprType}:            _GotoState135Action,
	{_State411, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State411, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State411, TypeExprType}:                      _GotoState249Action,
	{_State411, BinaryTypeExprType}:                _GotoState124Action,
	{_State411, FieldDefType}:                      _GotoState457Action,
	{_State411, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State411, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State411, TraitTypeExprType}:                 _GotoState136Action,
	{_State411, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State411, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State411, FuncTypeExprType}:                  _GotoState126Action,
	{_State411, MethodSignatureType}:               _GotoState246Action,
	{_State412, AddToken}:                          _GotoState253Action,
	{_State412, SubToken}:                          _GotoState255Action,
	{_State412, MulToken}:                          _GotoState254Action,
	{_State412, BinaryTypeOpType}:                  _GotoState256Action,
	{_State413, IdentifierToken}:                   _GotoState116Action,
	{_State413, UnderscoreToken}:                   _GotoState122Action,
	{_State413, StructToken}:                       _GotoState50Action,
	{_State413, EnumToken}:                         _GotoState113Action,
	{_State413, TraitToken}:                        _GotoState121Action,
	{_State413, FuncToken}:                         _GotoState115Action,
	{_State413, LparenToken}:                       _GotoState117Action,
	{_State413, LbracketToken}:                     _GotoState42Action,
	{_State413, DotToken}:                          _GotoState112Action,
	{_State413, QuestionToken}:                     _GotoState119Action,
	{_State413, ExclaimToken}:                      _GotoState114Action,
	{_State413, TildeTildeToken}:                   _GotoState120Action,
	{_State413, BitNegToken}:                       _GotoState111Action,
	{_State413, BitAndToken}:                       _GotoState110Action,
	{_State413, ParseErrorToken}:                   _GotoState118Action,
	{_State413, InitializableTypeExprType}:         _GotoState130Action,
	{_State413, SliceTypeExprType}:                 _GotoState103Action,
	{_State413, ArrayTypeExprType}:                 _GotoState60Action,
	{_State413, MapTypeExprType}:                   _GotoState90Action,
	{_State413, AtomTypeExprType}:                  _GotoState123Action,
	{_State413, NamedTypeExprType}:                 _GotoState131Action,
	{_State413, InferredTypeExprType}:              _GotoState129Action,
	{_State413, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State413, ReturnableTypeExprType}:            _GotoState435Action,
	{_State413, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State413, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State413, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State413, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State413, TraitTypeExprType}:                 _GotoState136Action,
	{_State413, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State413, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State413, ReturnTypeType}:                    _GotoState458Action,
	{_State413, FuncTypeExprType}:                  _GotoState126Action,
	{_State414, IdentifierToken}:                   _GotoState335Action,
	{_State414, UnderscoreToken}:                   _GotoState336Action,
	{_State414, StructToken}:                       _GotoState50Action,
	{_State414, EnumToken}:                         _GotoState113Action,
	{_State414, TraitToken}:                        _GotoState121Action,
	{_State414, FuncToken}:                         _GotoState115Action,
	{_State414, LparenToken}:                       _GotoState117Action,
	{_State414, LbracketToken}:                     _GotoState42Action,
	{_State414, DotToken}:                          _GotoState112Action,
	{_State414, QuestionToken}:                     _GotoState119Action,
	{_State414, ExclaimToken}:                      _GotoState114Action,
	{_State414, EllipsisToken}:                     _GotoState334Action,
	{_State414, TildeTildeToken}:                   _GotoState120Action,
	{_State414, BitNegToken}:                       _GotoState111Action,
	{_State414, BitAndToken}:                       _GotoState110Action,
	{_State414, ParseErrorToken}:                   _GotoState118Action,
	{_State414, InitializableTypeExprType}:         _GotoState130Action,
	{_State414, SliceTypeExprType}:                 _GotoState103Action,
	{_State414, ArrayTypeExprType}:                 _GotoState60Action,
	{_State414, MapTypeExprType}:                   _GotoState90Action,
	{_State414, AtomTypeExprType}:                  _GotoState123Action,
	{_State414, NamedTypeExprType}:                 _GotoState131Action,
	{_State414, InferredTypeExprType}:              _GotoState129Action,
	{_State414, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State414, ReturnableTypeExprType}:            _GotoState135Action,
	{_State414, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State414, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State414, TypeExprType}:                      _GotoState341Action,
	{_State414, BinaryTypeExprType}:                _GotoState124Action,
	{_State414, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State414, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State414, TraitTypeExprType}:                 _GotoState136Action,
	{_State414, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State414, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State414, ProperParameterDefType}:            _GotoState340Action,
	{_State414, ParameterDeclType}:                 _GotoState459Action,
	{_State414, FuncTypeExprType}:                  _GotoState126Action,
	{_State416, IdentifierToken}:                   _GotoState335Action,
	{_State416, UnderscoreToken}:                   _GotoState336Action,
	{_State416, StructToken}:                       _GotoState50Action,
	{_State416, EnumToken}:                         _GotoState113Action,
	{_State416, TraitToken}:                        _GotoState121Action,
	{_State416, FuncToken}:                         _GotoState115Action,
	{_State416, LparenToken}:                       _GotoState117Action,
	{_State416, LbracketToken}:                     _GotoState42Action,
	{_State416, DotToken}:                          _GotoState112Action,
	{_State416, QuestionToken}:                     _GotoState119Action,
	{_State416, ExclaimToken}:                      _GotoState114Action,
	{_State416, EllipsisToken}:                     _GotoState334Action,
	{_State416, TildeTildeToken}:                   _GotoState120Action,
	{_State416, BitNegToken}:                       _GotoState111Action,
	{_State416, BitAndToken}:                       _GotoState110Action,
	{_State416, ParseErrorToken}:                   _GotoState118Action,
	{_State416, InitializableTypeExprType}:         _GotoState130Action,
	{_State416, SliceTypeExprType}:                 _GotoState103Action,
	{_State416, ArrayTypeExprType}:                 _GotoState60Action,
	{_State416, MapTypeExprType}:                   _GotoState90Action,
	{_State416, AtomTypeExprType}:                  _GotoState123Action,
	{_State416, NamedTypeExprType}:                 _GotoState131Action,
	{_State416, InferredTypeExprType}:              _GotoState129Action,
	{_State416, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State416, ReturnableTypeExprType}:            _GotoState135Action,
	{_State416, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State416, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State416, TypeExprType}:                      _GotoState341Action,
	{_State416, BinaryTypeExprType}:                _GotoState124Action,
	{_State416, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State416, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State416, TraitTypeExprType}:                 _GotoState136Action,
	{_State416, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State416, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State416, ProperParameterDefType}:            _GotoState340Action,
	{_State416, ParameterDeclType}:                 _GotoState337Action,
	{_State416, ProperParameterDeclsType}:          _GotoState339Action,
	{_State416, ParameterDeclsType}:                _GotoState460Action,
	{_State416, FuncTypeExprType}:                  _GotoState126Action,
	{_State423, AddToken}:                          _GotoState253Action,
	{_State423, SubToken}:                          _GotoState255Action,
	{_State423, MulToken}:                          _GotoState254Action,
	{_State423, BinaryTypeOpType}:                  _GotoState256Action,
	{_State425, IdentifierToken}:                   _GotoState358Action,
	{_State425, GenericParameterDefType}:           _GotoState461Action,
	{_State426, RparenToken}:                       _GotoState462Action,
	{_State427, AddToken}:                          _GotoState253Action,
	{_State427, SubToken}:                          _GotoState255Action,
	{_State427, MulToken}:                          _GotoState254Action,
	{_State427, BinaryTypeOpType}:                  _GotoState256Action,
	{_State428, AddToken}:                          _GotoState253Action,
	{_State428, SubToken}:                          _GotoState255Action,
	{_State428, MulToken}:                          _GotoState254Action,
	{_State428, BinaryTypeOpType}:                  _GotoState256Action,
	{_State429, LparenToken}:                       _GotoState463Action,
	{_State430, IdentifierToken}:                   _GotoState116Action,
	{_State430, UnderscoreToken}:                   _GotoState122Action,
	{_State430, StructToken}:                       _GotoState50Action,
	{_State430, EnumToken}:                         _GotoState113Action,
	{_State430, TraitToken}:                        _GotoState121Action,
	{_State430, FuncToken}:                         _GotoState115Action,
	{_State430, LparenToken}:                       _GotoState117Action,
	{_State430, LbracketToken}:                     _GotoState42Action,
	{_State430, DotToken}:                          _GotoState112Action,
	{_State430, QuestionToken}:                     _GotoState119Action,
	{_State430, ExclaimToken}:                      _GotoState114Action,
	{_State430, TildeTildeToken}:                   _GotoState120Action,
	{_State430, BitNegToken}:                       _GotoState111Action,
	{_State430, BitAndToken}:                       _GotoState110Action,
	{_State430, ParseErrorToken}:                   _GotoState118Action,
	{_State430, InitializableTypeExprType}:         _GotoState130Action,
	{_State430, SliceTypeExprType}:                 _GotoState103Action,
	{_State430, ArrayTypeExprType}:                 _GotoState60Action,
	{_State430, MapTypeExprType}:                   _GotoState90Action,
	{_State430, AtomTypeExprType}:                  _GotoState123Action,
	{_State430, NamedTypeExprType}:                 _GotoState131Action,
	{_State430, InferredTypeExprType}:              _GotoState129Action,
	{_State430, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State430, ReturnableTypeExprType}:            _GotoState135Action,
	{_State430, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State430, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State430, TypeExprType}:                      _GotoState464Action,
	{_State430, BinaryTypeExprType}:                _GotoState124Action,
	{_State430, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State430, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State430, TraitTypeExprType}:                 _GotoState136Action,
	{_State430, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State430, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State430, FuncTypeExprType}:                  _GotoState126Action,
	{_State434, LbraceToken}:                       _GotoState11Action,
	{_State434, StatementBlockType}:                _GotoState465Action,
	{_State444, AddToken}:                          _GotoState253Action,
	{_State444, SubToken}:                          _GotoState255Action,
	{_State444, MulToken}:                          _GotoState254Action,
	{_State444, BinaryTypeOpType}:                  _GotoState256Action,
	{_State448, DoToken}:                           _GotoState466Action,
	{_State449, SemicolonToken}:                    _GotoState467Action,
	{_State452, IntegerLiteralToken}:               _GotoState40Action,
	{_State452, FloatLiteralToken}:                 _GotoState35Action,
	{_State452, RuneLiteralToken}:                  _GotoState48Action,
	{_State452, StringLiteralToken}:                _GotoState49Action,
	{_State452, IdentifierToken}:                   _GotoState38Action,
	{_State452, UnderscoreToken}:                   _GotoState53Action,
	{_State452, TrueToken}:                         _GotoState52Action,
	{_State452, FalseToken}:                        _GotoState34Action,
	{_State452, StructToken}:                       _GotoState50Action,
	{_State452, FuncToken}:                         _GotoState36Action,
	{_State452, VarToken}:                          _GotoState15Action,
	{_State452, LetToken}:                          _GotoState12Action,
	{_State452, NotToken}:                          _GotoState45Action,
	{_State452, LabelDeclToken}:                    _GotoState41Action,
	{_State452, LparenToken}:                       _GotoState43Action,
	{_State452, LbracketToken}:                     _GotoState42Action,
	{_State452, SubToken}:                          _GotoState51Action,
	{_State452, MulToken}:                          _GotoState44Action,
	{_State452, BitNegToken}:                       _GotoState27Action,
	{_State452, BitAndToken}:                       _GotoState26Action,
	{_State452, GreaterToken}:                      _GotoState37Action,
	{_State452, ParseErrorToken}:                   _GotoState46Action,
	{_State452, VarDeclPatternType}:                _GotoState107Action,
	{_State452, VarOrLetType}:                      _GotoState24Action,
	{_State452, OptionalLabelDeclType}:             _GotoState156Action,
	{_State452, SequenceExprType}:                  _GotoState468Action,
	{_State452, CallExprType}:                      _GotoState71Action,
	{_State452, AtomExprType}:                      _GotoState63Action,
	{_State452, ParseErrorExprType}:                _GotoState95Action,
	{_State452, LiteralExprType}:                   _GotoState88Action,
	{_State452, NamedExprType}:                     _GotoState92Action,
	{_State452, BlockExprType}:                     _GotoState70Action,
	{_State452, InitializeExprType}:                _GotoState85Action,
	{_State452, ImplicitStructExprType}:            _GotoState81Action,
	{_State452, AccessibleExprType}:                _GotoState108Action,
	{_State452, AccessExprType}:                    _GotoState55Action,
	{_State452, IndexExprType}:                     _GotoState83Action,
	{_State452, PostfixableExprType}:               _GotoState97Action,
	{_State452, PostfixUnaryExprType}:              _GotoState96Action,
	{_State452, PrefixableExprType}:                _GotoState100Action,
	{_State452, PrefixUnaryExprType}:               _GotoState98Action,
	{_State452, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State452, MulExprType}:                       _GotoState91Action,
	{_State452, BinaryMulExprType}:                 _GotoState67Action,
	{_State452, AddExprType}:                       _GotoState57Action,
	{_State452, BinaryAddExprType}:                 _GotoState64Action,
	{_State452, CmpExprType}:                       _GotoState74Action,
	{_State452, BinaryCmpExprType}:                 _GotoState66Action,
	{_State452, AndExprType}:                       _GotoState58Action,
	{_State452, BinaryAndExprType}:                 _GotoState65Action,
	{_State452, OrExprType}:                        _GotoState94Action,
	{_State452, BinaryOrExprType}:                  _GotoState69Action,
	{_State452, InitializableTypeExprType}:         _GotoState84Action,
	{_State452, SliceTypeExprType}:                 _GotoState103Action,
	{_State452, ArrayTypeExprType}:                 _GotoState60Action,
	{_State452, MapTypeExprType}:                   _GotoState90Action,
	{_State452, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State452, AnonymousFuncExprType}:             _GotoState59Action,
	{_State453, LabelDeclToken}:                    _GotoState41Action,
	{_State453, LbraceToken}:                       _GotoState11Action,
	{_State453, StatementBlockType}:                _GotoState471Action,
	{_State453, OptionalLabelDeclType}:             _GotoState470Action,
	{_State453, IfExprType}:                        _GotoState469Action,
	{_State460, RparenToken}:                       _GotoState472Action,
	{_State462, IdentifierToken}:                   _GotoState116Action,
	{_State462, UnderscoreToken}:                   _GotoState122Action,
	{_State462, StructToken}:                       _GotoState50Action,
	{_State462, EnumToken}:                         _GotoState113Action,
	{_State462, TraitToken}:                        _GotoState121Action,
	{_State462, FuncToken}:                         _GotoState115Action,
	{_State462, LparenToken}:                       _GotoState117Action,
	{_State462, LbracketToken}:                     _GotoState42Action,
	{_State462, DotToken}:                          _GotoState112Action,
	{_State462, QuestionToken}:                     _GotoState119Action,
	{_State462, ExclaimToken}:                      _GotoState114Action,
	{_State462, TildeTildeToken}:                   _GotoState120Action,
	{_State462, BitNegToken}:                       _GotoState111Action,
	{_State462, BitAndToken}:                       _GotoState110Action,
	{_State462, ParseErrorToken}:                   _GotoState118Action,
	{_State462, InitializableTypeExprType}:         _GotoState130Action,
	{_State462, SliceTypeExprType}:                 _GotoState103Action,
	{_State462, ArrayTypeExprType}:                 _GotoState60Action,
	{_State462, MapTypeExprType}:                   _GotoState90Action,
	{_State462, AtomTypeExprType}:                  _GotoState123Action,
	{_State462, NamedTypeExprType}:                 _GotoState131Action,
	{_State462, InferredTypeExprType}:              _GotoState129Action,
	{_State462, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State462, ReturnableTypeExprType}:            _GotoState435Action,
	{_State462, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State462, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State462, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State462, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State462, TraitTypeExprType}:                 _GotoState136Action,
	{_State462, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State462, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State462, ReturnTypeType}:                    _GotoState473Action,
	{_State462, FuncTypeExprType}:                  _GotoState126Action,
	{_State463, IdentifierToken}:                   _GotoState260Action,
	{_State463, UnderscoreToken}:                   _GotoState261Action,
	{_State463, ProperParameterDefType}:            _GotoState263Action,
	{_State463, ParameterDefType}:                  _GotoState284Action,
	{_State463, ProperParameterDefsType}:           _GotoState286Action,
	{_State463, ParameterDefsType}:                 _GotoState474Action,
	{_State464, AddToken}:                          _GotoState253Action,
	{_State464, SubToken}:                          _GotoState255Action,
	{_State464, MulToken}:                          _GotoState254Action,
	{_State464, BinaryTypeOpType}:                  _GotoState256Action,
	{_State466, LbraceToken}:                       _GotoState11Action,
	{_State466, StatementBlockType}:                _GotoState475Action,
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
	{_State467, VarDeclPatternType}:                _GotoState107Action,
	{_State467, VarOrLetType}:                      _GotoState24Action,
	{_State467, OptionalLabelDeclType}:             _GotoState156Action,
	{_State467, SequenceExprType}:                  _GotoState450Action,
	{_State467, OptionalSequenceExprType}:          _GotoState476Action,
	{_State467, CallExprType}:                      _GotoState71Action,
	{_State467, AtomExprType}:                      _GotoState63Action,
	{_State467, ParseErrorExprType}:                _GotoState95Action,
	{_State467, LiteralExprType}:                   _GotoState88Action,
	{_State467, NamedExprType}:                     _GotoState92Action,
	{_State467, BlockExprType}:                     _GotoState70Action,
	{_State467, InitializeExprType}:                _GotoState85Action,
	{_State467, ImplicitStructExprType}:            _GotoState81Action,
	{_State467, AccessibleExprType}:                _GotoState108Action,
	{_State467, AccessExprType}:                    _GotoState55Action,
	{_State467, IndexExprType}:                     _GotoState83Action,
	{_State467, PostfixableExprType}:               _GotoState97Action,
	{_State467, PostfixUnaryExprType}:              _GotoState96Action,
	{_State467, PrefixableExprType}:                _GotoState100Action,
	{_State467, PrefixUnaryExprType}:               _GotoState98Action,
	{_State467, PrefixUnaryOpType}:                 _GotoState99Action,
	{_State467, MulExprType}:                       _GotoState91Action,
	{_State467, BinaryMulExprType}:                 _GotoState67Action,
	{_State467, AddExprType}:                       _GotoState57Action,
	{_State467, BinaryAddExprType}:                 _GotoState64Action,
	{_State467, CmpExprType}:                       _GotoState74Action,
	{_State467, BinaryCmpExprType}:                 _GotoState66Action,
	{_State467, AndExprType}:                       _GotoState58Action,
	{_State467, BinaryAndExprType}:                 _GotoState65Action,
	{_State467, OrExprType}:                        _GotoState94Action,
	{_State467, BinaryOrExprType}:                  _GotoState69Action,
	{_State467, InitializableTypeExprType}:         _GotoState84Action,
	{_State467, SliceTypeExprType}:                 _GotoState103Action,
	{_State467, ArrayTypeExprType}:                 _GotoState60Action,
	{_State467, MapTypeExprType}:                   _GotoState90Action,
	{_State467, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State467, AnonymousFuncExprType}:             _GotoState59Action,
	{_State470, IfToken}:                           _GotoState229Action,
	{_State470, IfExprBodyType}:                    _GotoState231Action,
	{_State472, IdentifierToken}:                   _GotoState116Action,
	{_State472, UnderscoreToken}:                   _GotoState122Action,
	{_State472, StructToken}:                       _GotoState50Action,
	{_State472, EnumToken}:                         _GotoState113Action,
	{_State472, TraitToken}:                        _GotoState121Action,
	{_State472, FuncToken}:                         _GotoState115Action,
	{_State472, LparenToken}:                       _GotoState117Action,
	{_State472, LbracketToken}:                     _GotoState42Action,
	{_State472, DotToken}:                          _GotoState112Action,
	{_State472, QuestionToken}:                     _GotoState119Action,
	{_State472, ExclaimToken}:                      _GotoState114Action,
	{_State472, TildeTildeToken}:                   _GotoState120Action,
	{_State472, BitNegToken}:                       _GotoState111Action,
	{_State472, BitAndToken}:                       _GotoState110Action,
	{_State472, ParseErrorToken}:                   _GotoState118Action,
	{_State472, InitializableTypeExprType}:         _GotoState130Action,
	{_State472, SliceTypeExprType}:                 _GotoState103Action,
	{_State472, ArrayTypeExprType}:                 _GotoState60Action,
	{_State472, MapTypeExprType}:                   _GotoState90Action,
	{_State472, AtomTypeExprType}:                  _GotoState123Action,
	{_State472, NamedTypeExprType}:                 _GotoState131Action,
	{_State472, InferredTypeExprType}:              _GotoState129Action,
	{_State472, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State472, ReturnableTypeExprType}:            _GotoState435Action,
	{_State472, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State472, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State472, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State472, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State472, TraitTypeExprType}:                 _GotoState136Action,
	{_State472, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State472, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State472, ReturnTypeType}:                    _GotoState477Action,
	{_State472, FuncTypeExprType}:                  _GotoState126Action,
	{_State473, LbraceToken}:                       _GotoState11Action,
	{_State473, StatementBlockType}:                _GotoState478Action,
	{_State474, RparenToken}:                       _GotoState479Action,
	{_State476, DoToken}:                           _GotoState480Action,
	{_State479, IdentifierToken}:                   _GotoState116Action,
	{_State479, UnderscoreToken}:                   _GotoState122Action,
	{_State479, StructToken}:                       _GotoState50Action,
	{_State479, EnumToken}:                         _GotoState113Action,
	{_State479, TraitToken}:                        _GotoState121Action,
	{_State479, FuncToken}:                         _GotoState115Action,
	{_State479, LparenToken}:                       _GotoState117Action,
	{_State479, LbracketToken}:                     _GotoState42Action,
	{_State479, DotToken}:                          _GotoState112Action,
	{_State479, QuestionToken}:                     _GotoState119Action,
	{_State479, ExclaimToken}:                      _GotoState114Action,
	{_State479, TildeTildeToken}:                   _GotoState120Action,
	{_State479, BitNegToken}:                       _GotoState111Action,
	{_State479, BitAndToken}:                       _GotoState110Action,
	{_State479, ParseErrorToken}:                   _GotoState118Action,
	{_State479, InitializableTypeExprType}:         _GotoState130Action,
	{_State479, SliceTypeExprType}:                 _GotoState103Action,
	{_State479, ArrayTypeExprType}:                 _GotoState60Action,
	{_State479, MapTypeExprType}:                   _GotoState90Action,
	{_State479, AtomTypeExprType}:                  _GotoState123Action,
	{_State479, NamedTypeExprType}:                 _GotoState131Action,
	{_State479, InferredTypeExprType}:              _GotoState129Action,
	{_State479, ParseErrorTypeExprType}:            _GotoState132Action,
	{_State479, ReturnableTypeExprType}:            _GotoState435Action,
	{_State479, PrefixUnaryTypeExprType}:           _GotoState133Action,
	{_State479, PrefixUnaryTypeOpType}:             _GotoState134Action,
	{_State479, ImplicitStructTypeExprType}:        _GotoState128Action,
	{_State479, ExplicitStructTypeExprType}:        _GotoState75Action,
	{_State479, TraitTypeExprType}:                 _GotoState136Action,
	{_State479, ImplicitEnumTypeExprType}:          _GotoState127Action,
	{_State479, ExplicitEnumTypeExprType}:          _GotoState125Action,
	{_State479, ReturnTypeType}:                    _GotoState481Action,
	{_State479, FuncTypeExprType}:                  _GotoState126Action,
	{_State480, LbraceToken}:                       _GotoState11Action,
	{_State480, StatementBlockType}:                _GotoState482Action,
	{_State481, LbraceToken}:                       _GotoState11Action,
	{_State481, StatementBlockType}:                _GotoState483Action,
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
	{_State80, _EndMarker}:                         _ReduceIfExprToExprAction,
	{_State81, _WildcardMarker}:                    _ReduceImplicitStructExprToAtomExprAction,
	{_State82, _EndMarker}:                         _ReduceImportStatementToStatementAction,
	{_State83, _WildcardMarker}:                    _ReduceIndexExprToAccessibleExprAction,
	{_State85, _WildcardMarker}:                    _ReduceInitializeExprToAtomExprAction,
	{_State86, _EndMarker}:                         _ReduceJumpStatementToSimpleStatementAction,
	{_State87, _EndMarker}:                         _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State87, NewlinesToken}:                      _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State87, RbraceToken}:                        _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State87, SemicolonToken}:                     _ReduceUnlabeledNoValueToJumpStatementAction,
	{_State87, _WildcardMarker}:                    _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State88, _WildcardMarker}:                    _ReduceLiteralExprToAtomExprAction,
	{_State89, _EndMarker}:                         _ReduceLoopExprToExprAction,
	{_State90, _WildcardMarker}:                    _ReduceMapTypeExprToInitializableTypeExprAction,
	{_State91, _WildcardMarker}:                    _ReduceMulExprToAddExprAction,
	{_State92, _WildcardMarker}:                    _ReduceNamedExprToAtomExprAction,
	{_State94, _WildcardMarker}:                    _ReduceOrExprToSequenceExprAction,
	{_State95, _WildcardMarker}:                    _ReduceParseErrorExprToAtomExprAction,
	{_State96, _WildcardMarker}:                    _ReducePostfixUnaryExprToPostfixableExprAction,
	{_State97, _WildcardMarker}:                    _ReducePostfixableExprToPrefixableExprAction,
	{_State98, _WildcardMarker}:                    _ReducePrefixUnaryExprToPrefixableExprAction,
	{_State99, LbraceToken}:                        _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State100, _WildcardMarker}:                   _ReducePrefixableExprToMulExprAction,
	{_State101, AssignToken}:                       _ReduceToAssignPatternAction,
	{_State101, _WildcardMarker}:                   _ReduceSequenceExprToExprAction,
	{_State102, _EndMarker}:                        _ReduceSimpleStatementToStatementAction,
	{_State103, _WildcardMarker}:                   _ReduceSliceTypeExprToInitializableTypeExprAction,
	{_State104, _EndMarker}:                        _ReduceSwitchExprToExprAction,
	{_State105, _EndMarker}:                        _ReduceUnaryOpAssignStatementToSimpleStatementAction,
	{_State106, _EndMarker}:                        _ReduceUnsafeStatementToSimpleStatementAction,
	{_State107, _EndMarker}:                        _ReduceVarDeclPatternToSequenceExprAction,
	{_State108, _WildcardMarker}:                   _ReduceAccessibleExprToPostfixableExprAction,
	{_State108, LparenToken}:                       _ReduceNilToGenericTypeArgumentsAction,
	{_State109, _EndMarker}:                        _ReduceSequenceExprToExprAction,
	{_State110, _WildcardMarker}:                   _ReduceBitAndToPrefixUnaryTypeOpAction,
	{_State111, _WildcardMarker}:                   _ReduceBitNegToPrefixUnaryTypeOpAction,
	{_State112, _WildcardMarker}:                   _ReduceDotToInferredTypeExprAction,
	{_State114, _WildcardMarker}:                   _ReduceExclaimToPrefixUnaryTypeOpAction,
	{_State116, _WildcardMarker}:                   _ReduceNilToGenericTypeArgumentsAction,
	{_State117, RparenToken}:                       _ReduceNilToImplicitFieldDefsAction,
	{_State118, _WildcardMarker}:                   _ReduceToParseErrorTypeExprAction,
	{_State119, _WildcardMarker}:                   _ReduceQuestionToPrefixUnaryTypeOpAction,
	{_State120, _WildcardMarker}:                   _ReduceTildeTildeToPrefixUnaryTypeOpAction,
	{_State122, _WildcardMarker}:                   _ReduceUnderscoreToInferredTypeExprAction,
	{_State123, _WildcardMarker}:                   _ReduceAtomTypeExprToReturnableTypeExprAction,
	{_State124, _WildcardMarker}:                   _ReduceBinaryTypeExprToTypeExprAction,
	{_State125, _WildcardMarker}:                   _ReduceExplicitEnumTypeExprToAtomTypeExprAction,
	{_State126, _WildcardMarker}:                   _ReduceFuncTypeExprToAtomTypeExprAction,
	{_State127, _WildcardMarker}:                   _ReduceImplicitEnumTypeExprToAtomTypeExprAction,
	{_State128, _WildcardMarker}:                   _ReduceImplicitStructTypeExprToAtomTypeExprAction,
	{_State129, _WildcardMarker}:                   _ReduceInferredTypeExprToAtomTypeExprAction,
	{_State130, _WildcardMarker}:                   _ReduceInitializableTypeExprToAtomTypeExprAction,
	{_State131, _WildcardMarker}:                   _ReduceNamedTypeExprToAtomTypeExprAction,
	{_State132, _WildcardMarker}:                   _ReduceParseErrorTypeExprToAtomTypeExprAction,
	{_State133, _WildcardMarker}:                   _ReducePrefixUnaryTypeExprToReturnableTypeExprAction,
	{_State135, _WildcardMarker}:                   _ReduceReturnableTypeExprToTypeExprAction,
	{_State136, _WildcardMarker}:                   _ReduceTraitTypeExprToAtomTypeExprAction,
	{_State137, LparenToken}:                       _ReduceNilToGenericParameterDefsAction,
	{_State139, RbraceToken}:                       _ReduceProperStatementsToStatementsAction,
	{_State140, _WildcardMarker}:                   _ReduceStatementToProperStatementsAction,
	{_State142, _WildcardMarker}:                   _ReduceWithSpecToPackageDefAction,
	{_State143, _WildcardMarker}:                   _ReduceNilToGenericParameterDefsAction,
	{_State144, _EndMarker}:                        _ReduceImproperToDefinitionsAction,
	{_State145, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State146, _WildcardMarker}:                   _ReduceIdentifierToVarPatternAction,
	{_State148, _WildcardMarker}:                   _ReduceUnderscoreToVarPatternAction,
	{_State149, _WildcardMarker}:                   _ReduceTuplePatternToVarPatternAction,
	{_State150, _WildcardMarker}:                   _ReduceNilToOptionalTypeExprAction,
	{_State152, _WildcardMarker}:                   _ReduceVarToVarOrLetAction,
	{_State153, _WildcardMarker}:                   _ReduceAssignPatternToCasePatternAction,
	{_State154, _WildcardMarker}:                   _ReduceCasePatternToCasePatternsAction,
	{_State157, _WildcardMarker}:                   _ReduceToAssignPatternAction,
	{_State158, _EndMarker}:                        _ReduceNilToOptionalSimpleStatementAction,
	{_State158, NewlinesToken}:                     _ReduceNilToOptionalSimpleStatementAction,
	{_State158, RbraceToken}:                       _ReduceNilToOptionalSimpleStatementAction,
	{_State158, SemicolonToken}:                    _ReduceNilToOptionalSimpleStatementAction,
	{_State158, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State159, RparenToken}:                       _ReduceNilToParameterDefsAction,
	{_State160, _EndMarker}:                        _ReduceAssignVarPatternToSequenceExprAction,
	{_State164, _EndMarker}:                        _ReduceStringLiteralToImportClauseAction,
	{_State166, _EndMarker}:                        _ReduceSingleToImportStatementAction,
	{_State168, ColonToken}:                        _ReduceUnitUnitPairToColonExprAction,
	{_State168, CommaToken}:                        _ReduceUnitUnitPairToColonExprAction,
	{_State168, RbracketToken}:                     _ReduceUnitUnitPairToColonExprAction,
	{_State168, RparenToken}:                       _ReduceUnitUnitPairToColonExprAction,
	{_State168, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State169, _WildcardMarker}:                   _ReduceSkipPatternToArgumentAction,
	{_State170, _WildcardMarker}:                   _ReduceIdentifierToNamedExprAction,
	{_State171, _WildcardMarker}:                   _ReduceArgumentToProperArgumentsAction,
	{_State173, _WildcardMarker}:                   _ReduceColonExprToArgumentAction,
	{_State174, _WildcardMarker}:                   _ReducePositionalToArgumentAction,
	{_State175, RparenToken}:                       _ReduceProperArgumentsToArgumentsAction,
	{_State176, RparenToken}:                       _ReduceNilToExplicitFieldDefsAction,
	{_State178, _WildcardMarker}:                   _ReduceAddAssignToBinaryOpAssignAction,
	{_State179, _EndMarker}:                        _ReduceAddOneAssignToUnaryOpAssignAction,
	{_State180, _WildcardMarker}:                   _ReduceBitAndAssignToBinaryOpAssignAction,
	{_State181, _WildcardMarker}:                   _ReduceBitLshiftAssignToBinaryOpAssignAction,
	{_State182, _WildcardMarker}:                   _ReduceBitNegAssignToBinaryOpAssignAction,
	{_State183, _WildcardMarker}:                   _ReduceBitOrAssignToBinaryOpAssignAction,
	{_State184, _WildcardMarker}:                   _ReduceBitRshiftAssignToBinaryOpAssignAction,
	{_State185, _WildcardMarker}:                   _ReduceBitXorAssignToBinaryOpAssignAction,
	{_State186, _WildcardMarker}:                   _ReduceDivAssignToBinaryOpAssignAction,
	{_State187, RbracketToken}:                     _ReduceNilToTypeArgumentsAction,
	{_State189, _WildcardMarker}:                   _ReduceExclaimToPostfixUnaryOpAction,
	{_State190, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State191, _WildcardMarker}:                   _ReduceModAssignToBinaryOpAssignAction,
	{_State192, _WildcardMarker}:                   _ReduceMulAssignToBinaryOpAssignAction,
	{_State193, _WildcardMarker}:                   _ReduceQuestionToPostfixUnaryOpAction,
	{_State194, _WildcardMarker}:                   _ReduceSubAssignToBinaryOpAssignAction,
	{_State195, _EndMarker}:                        _ReduceSubOneAssignToUnaryOpAssignAction,
	{_State196, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State198, _WildcardMarker}:                   _ReduceToPostfixUnaryExprAction,
	{_State199, _EndMarker}:                        _ReduceToUnaryOpAssignStatementAction,
	{_State200, _WildcardMarker}:                   _ReduceAddToAddOpAction,
	{_State201, _WildcardMarker}:                   _ReduceBitOrToAddOpAction,
	{_State202, _WildcardMarker}:                   _ReduceBitXorToAddOpAction,
	{_State203, _WildcardMarker}:                   _ReduceSubToAddOpAction,
	{_State204, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State205, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State206, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State207, LparenToken}:                       _ReduceNilToGenericTypeArgumentsAction,
	{_State208, _EndMarker}:                        _ReduceToCallbackOpStatementAction,
	{_State208, NewlinesToken}:                     _ReduceToCallbackOpStatementAction,
	{_State208, RbraceToken}:                       _ReduceToCallbackOpStatementAction,
	{_State208, SemicolonToken}:                    _ReduceToCallbackOpStatementAction,
	{_State208, _WildcardMarker}:                   _ReduceCallExprToAccessibleExprAction,
	{_State209, _WildcardMarker}:                   _ReduceEqualToCmpOpAction,
	{_State210, _WildcardMarker}:                   _ReduceGreaterToCmpOpAction,
	{_State211, _WildcardMarker}:                   _ReduceGreaterOrEqualToCmpOpAction,
	{_State212, _WildcardMarker}:                   _ReduceLessToCmpOpAction,
	{_State213, _WildcardMarker}:                   _ReduceLessOrEqualToCmpOpAction,
	{_State214, _WildcardMarker}:                   _ReduceNotEqualToCmpOpAction,
	{_State215, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State216, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State217, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State217, RparenToken}:                       _ReduceNilToArgumentsAction,
	{_State218, _EndMarker}:                        _ReduceLabeledNoValueToJumpStatementAction,
	{_State218, NewlinesToken}:                     _ReduceLabeledNoValueToJumpStatementAction,
	{_State218, RbraceToken}:                       _ReduceLabeledNoValueToJumpStatementAction,
	{_State218, SemicolonToken}:                    _ReduceLabeledNoValueToJumpStatementAction,
	{_State218, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State219, _EndMarker}:                        _ReduceUnlabeledValuedToJumpStatementAction,
	{_State220, _WildcardMarker}:                   _ReduceBitAndToMulOpAction,
	{_State221, _WildcardMarker}:                   _ReduceBitLshiftToMulOpAction,
	{_State222, _WildcardMarker}:                   _ReduceBitRshiftToMulOpAction,
	{_State223, _WildcardMarker}:                   _ReduceDivToMulOpAction,
	{_State224, _WildcardMarker}:                   _ReduceModToMulOpAction,
	{_State225, _WildcardMarker}:                   _ReduceMulToMulOpAction,
	{_State226, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State228, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State229, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State230, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State231, _EndMarker}:                        _ReduceToIfExprAction,
	{_State232, _EndMarker}:                        _ReduceToLoopExprAction,
	{_State233, _WildcardMarker}:                   _ReduceToBlockExprAction,
	{_State234, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State235, _WildcardMarker}:                   _ReduceToPrefixUnaryExprAction,
	{_State237, RparenToken}:                       _ReduceNilToParameterDeclsAction,
	{_State239, _WildcardMarker}:                   _ReduceLocalToNamedTypeExprAction,
	{_State241, _WildcardMarker}:                   _ReduceNilToGenericTypeArgumentsAction,
	{_State242, _WildcardMarker}:                   _ReduceUnderscoreToInferredTypeExprAction,
	{_State243, _WildcardMarker}:                   _ReduceFieldDefToProperImplicitFieldDefsAction,
	{_State246, _WildcardMarker}:                   _ReduceMethodSignatureToFieldDefAction,
	{_State247, RparenToken}:                       _ReduceProperImplicitEnumValueDefsToImplicitEnumValueDefsAction,
	{_State248, RparenToken}:                       _ReduceProperImplicitFieldDefsToImplicitFieldDefsAction,
	{_State249, _WildcardMarker}:                   _ReduceNilToOptionalDefaultAction,
	{_State250, _WildcardMarker}:                   _ReduceUnsafeStatementToFieldDefAction,
	{_State251, RparenToken}:                       _ReduceNilToExplicitFieldDefsAction,
	{_State252, _WildcardMarker}:                   _ReduceToPrefixUnaryTypeExprAction,
	{_State253, _WildcardMarker}:                   _ReduceAddToBinaryTypeOpAction,
	{_State254, _WildcardMarker}:                   _ReduceMulToBinaryTypeOpAction,
	{_State255, _WildcardMarker}:                   _ReduceSubToBinaryTypeOpAction,
	{_State257, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State258, RbracketToken}:                     _ReduceNilToGenericParameterDefListAction,
	{_State260, _WildcardMarker}:                   _ReduceNamedInferredArgToParameterDefAction,
	{_State261, _WildcardMarker}:                   _ReduceIgnoreInferredArgToParameterDefAction,
	{_State263, _WildcardMarker}:                   _ReduceProperParameterDefToParameterDefAction,
	{_State264, RbraceToken}:                       _ReduceImproperImplicitToStatementsAction,
	{_State264, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State265, RbraceToken}:                       _ReduceImproperExplicitToStatementsAction,
	{_State265, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State266, _EndMarker}:                        _ReduceToStatementBlockAction,
	{_State269, _WildcardMarker}:                   _ReduceAddToProperDefinitionsAction,
	{_State270, _WildcardMarker}:                   _ReduceGlobalVarAssignmentToDefinitionAction,
	{_State271, _WildcardMarker}:                   _ReduceEllipsisToFieldVarPatternAction,
	{_State272, _WildcardMarker}:                   _ReduceIdentifierToVarPatternAction,
	{_State273, _WildcardMarker}:                   _ReduceFieldVarPatternToFieldVarPatternsAction,
	{_State275, _WildcardMarker}:                   _ReducePositionalToFieldVarPatternAction,
	{_State276, _EndMarker}:                        _ReduceToVarDeclPatternAction,
	{_State277, _WildcardMarker}:                   _ReduceTypeExprToOptionalTypeExprAction,
	{_State278, _WildcardMarker}:                   _ReduceEnumNondataMatchPattenToCasePatternAction,
	{_State280, _EndMarker}:                        _ReduceNilToOptionalSimpleStatementAction,
	{_State280, NewlinesToken}:                     _ReduceNilToOptionalSimpleStatementAction,
	{_State280, RbraceToken}:                       _ReduceNilToOptionalSimpleStatementAction,
	{_State280, SemicolonToken}:                    _ReduceNilToOptionalSimpleStatementAction,
	{_State280, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State281, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State282, _EndMarker}:                        _ReduceDefaultBranchStatementToStatementAction,
	{_State283, _EndMarker}:                        _ReduceSimpleStatementToOptionalSimpleStatementAction,
	{_State284, _WildcardMarker}:                   _ReduceParameterDefToProperParameterDefsAction,
	{_State286, RparenToken}:                       _ReduceProperParameterDefsToParameterDefsAction,
	{_State287, _EndMarker}:                        _ReduceImportToLocalToImportClauseAction,
	{_State288, _EndMarker}:                        _ReduceAliasToImportClauseAction,
	{_State289, _WildcardMarker}:                   _ReduceImportClauseToProperImportClausesAction,
	{_State291, RparenToken}:                       _ReduceProperImportClausesToImportClausesAction,
	{_State292, _EndMarker}:                        _ReduceUnusableImportToImportClauseAction,
	{_State295, _WildcardMarker}:                   _ReduceToSliceTypeExprAction,
	{_State296, _WildcardMarker}:                   _ReduceUnitExprPairToColonExprAction,
	{_State297, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State298, _WildcardMarker}:                   _ReduceToImplicitStructExprAction,
	{_State299, ColonToken}:                        _ReduceColonExprUnitTupleToColonExprAction,
	{_State299, CommaToken}:                        _ReduceColonExprUnitTupleToColonExprAction,
	{_State299, RbracketToken}:                     _ReduceColonExprUnitTupleToColonExprAction,
	{_State299, RparenToken}:                       _ReduceColonExprUnitTupleToColonExprAction,
	{_State299, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State300, ColonToken}:                        _ReduceExprUnitPairToColonExprAction,
	{_State300, CommaToken}:                        _ReduceExprUnitPairToColonExprAction,
	{_State300, RbracketToken}:                     _ReduceExprUnitPairToColonExprAction,
	{_State300, RparenToken}:                       _ReduceExprUnitPairToColonExprAction,
	{_State300, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State301, _WildcardMarker}:                   _ReduceVarargAssignmentToArgumentAction,
	{_State302, RparenToken}:                       _ReduceImproperToArgumentsAction,
	{_State302, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State304, _WildcardMarker}:                   _ReduceFieldDefToProperExplicitFieldDefsAction,
	{_State305, RparenToken}:                       _ReduceProperExplicitFieldDefsToExplicitFieldDefsAction,
	{_State307, RbracketToken}:                     _ReduceProperTypeArgumentsToTypeArgumentsAction,
	{_State309, _WildcardMarker}:                   _ReduceTypeExprToProperTypeArgumentsAction,
	{_State310, _WildcardMarker}:                   _ReduceToAccessExprAction,
	{_State312, _EndMarker}:                        _ReduceToBinaryOpAssignStatementAction,
	{_State313, _WildcardMarker}:                   _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State313, RparenToken}:                       _ReduceNilToArgumentsAction,
	{_State314, _WildcardMarker}:                   _ReduceToBinaryAddExprAction,
	{_State315, _WildcardMarker}:                   _ReduceToBinaryAndExprAction,
	{_State316, _EndMarker}:                        _ReduceToAssignStatementAction,
	{_State317, _WildcardMarker}:                   _ReduceToBinaryCmpExprAction,
	{_State318, _WildcardMarker}:                   _ReduceAddToExprsAction,
	{_State320, _EndMarker}:                        _ReduceLabeledValuedToJumpStatementAction,
	{_State321, _WildcardMarker}:                   _ReduceToBinaryMulExprAction,
	{_State322, _WildcardMarker}:                   _ReduceInfiniteToLoopExprBodyAction,
	{_State325, _WildcardMarker}:                   _ReduceToAssignPatternAction,
	{_State325, SemicolonToken}:                    _ReduceSequenceExprToForAssignmentAction,
	{_State326, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State328, LbraceToken}:                       _ReduceSequenceExprToConditionAction,
	{_State330, _WildcardMarker}:                   _ReduceToBinaryOrExprAction,
	{_State333, RparenToken}:                       _ReduceProperExplicitEnumValueDefsToExplicitEnumValueDefsAction,
	{_State334, _WildcardMarker}:                   _ReduceUnnamedInferredVarargToParameterDeclAction,
	{_State335, _WildcardMarker}:                   _ReduceNilToGenericTypeArgumentsAction,
	{_State336, _WildcardMarker}:                   _ReduceUnderscoreToInferredTypeExprAction,
	{_State337, _WildcardMarker}:                   _ReduceParameterDeclToProperParameterDeclsAction,
	{_State339, RparenToken}:                       _ReduceProperParameterDeclsToParameterDeclsAction,
	{_State340, _WildcardMarker}:                   _ReduceProperParameterDefToParameterDeclAction,
	{_State341, _WildcardMarker}:                   _ReduceUnnamedTypedArgToParameterDeclAction,
	{_State342, _WildcardMarker}:                   _ReduceNilToGenericTypeArgumentsAction,
	{_State344, _WildcardMarker}:                   _ReduceDotToInferredTypeExprAction,
	{_State345, _WildcardMarker}:                   _ReduceNilToOptionalDefaultAction,
	{_State346, _WildcardMarker}:                   _ReduceStructPaddingToFieldDefAction,
	{_State348, _WildcardMarker}:                   _ReduceToImplicitEnumTypeExprAction,
	{_State349, _WildcardMarker}:                   _ReduceToImplicitStructTypeExprAction,
	{_State350, RparenToken}:                       _ReduceImproperToImplicitEnumValueDefsAction,
	{_State352, RparenToken}:                       _ReduceImproperToImplicitFieldDefsAction,
	{_State354, _WildcardMarker}:                   _ReduceUnnamedToFieldDefAction,
	{_State356, _WildcardMarker}:                   _ReduceToBinaryTypeExprAction,
	{_State357, _WildcardMarker}:                   _ReduceAliasToNamedFuncDefAction,
	{_State358, _WildcardMarker}:                   _ReduceUnconstrainedToGenericParameterDefAction,
	{_State359, _WildcardMarker}:                   _ReduceGenericParameterDefToProperGenericParameterDefListAction,
	{_State361, RbracketToken}:                     _ReduceProperGenericParameterDefListToGenericParameterDefListAction,
	{_State362, RparenToken}:                       _ReduceNilToParameterDefsAction,
	{_State363, _WildcardMarker}:                   _ReduceNamedInferredVarargToProperParameterDefAction,
	{_State364, _WildcardMarker}:                   _ReduceNamedTypedArgToProperParameterDefAction,
	{_State365, _WildcardMarker}:                   _ReduceIgnoreInferredVarargToProperParameterDefAction,
	{_State366, _WildcardMarker}:                   _ReduceIgnoreTypedArgToProperParameterDefAction,
	{_State368, _WildcardMarker}:                   _ReduceAddImplicitToProperStatementsAction,
	{_State369, _WildcardMarker}:                   _ReduceAddExplicitToProperStatementsAction,
	{_State370, _WildcardMarker}:                   _ReduceAliasToTypeDefAction,
	{_State371, _WildcardMarker}:                   _ReduceDefinitionToTypeDefAction,
	{_State374, _WildcardMarker}:                   _ReduceToTuplePatternAction,
	{_State375, _WildcardMarker}:                   _ReduceEnumMatchPatternToCasePatternAction,
	{_State377, _EndMarker}:                        _ReduceCaseBranchStatementToStatementAction,
	{_State378, _WildcardMarker}:                   _ReduceMultipleToCasePatternsAction,
	{_State379, LbraceToken}:                       _ReduceNilToReturnTypeAction,
	{_State380, RparenToken}:                       _ReduceImproperToParameterDefsAction,
	{_State381, _EndMarker}:                        _ReduceMultipleToImportStatementAction,
	{_State382, RparenToken}:                       _ReduceExplicitToImportClausesAction,
	{_State383, RparenToken}:                       _ReduceImplicitToImportClausesAction,
	{_State386, _WildcardMarker}:                   _ReduceNamedAssignmentToArgumentAction,
	{_State387, _WildcardMarker}:                   _ReduceColonExprExprTupleToColonExprAction,
	{_State388, _WildcardMarker}:                   _ReduceExprExprPairToColonExprAction,
	{_State389, _WildcardMarker}:                   _ReduceAddToProperArgumentsAction,
	{_State390, _WildcardMarker}:                   _ReduceToExplicitStructTypeExprAction,
	{_State391, RparenToken}:                       _ReduceImproperExplicitToExplicitFieldDefsAction,
	{_State392, RparenToken}:                       _ReduceImproperImplicitToExplicitFieldDefsAction,
	{_State394, RbracketToken}:                     _ReduceImproperToTypeArgumentsAction,
	{_State395, _WildcardMarker}:                   _ReduceBindingToGenericTypeArgumentsAction,
	{_State396, _WildcardMarker}:                   _ReduceToIndexExprAction,
	{_State398, _WildcardMarker}:                   _ReduceToInitializeExprAction,
	{_State399, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State400, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State401, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State402, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State402, SemicolonToken}:                    _ReduceNilToOptionalSequenceExprAction,
	{_State405, _WildcardMarker}:                   _ReduceNoElseToIfExprBodyAction,
	{_State406, _EndMarker}:                        _ReduceToSwitchExprAction,
	{_State407, _WildcardMarker}:                   _ReduceToExplicitEnumTypeExprAction,
	{_State410, RparenToken}:                       _ReduceImproperToExplicitEnumValueDefsAction,
	{_State412, _WildcardMarker}:                   _ReduceUnnamedTypedVarargToParameterDeclAction,
	{_State413, _WildcardMarker}:                   _ReduceNilToReturnTypeAction,
	{_State414, RparenToken}:                       _ReduceImproperToParameterDeclsAction,
	{_State415, _WildcardMarker}:                   _ReduceExternalToNamedTypeExprAction,
	{_State416, RparenToken}:                       _ReduceNilToParameterDeclsAction,
	{_State417, _WildcardMarker}:                   _ReduceNamedToFieldDefAction,
	{_State418, _WildcardMarker}:                   _ReducePairToProperImplicitEnumValueDefsAction,
	{_State419, _WildcardMarker}:                   _ReduceAddToProperImplicitEnumValueDefsAction,
	{_State420, _WildcardMarker}:                   _ReduceAddToProperImplicitFieldDefsAction,
	{_State421, _WildcardMarker}:                   _ReduceDefaultToOptionalDefaultAction,
	{_State422, _WildcardMarker}:                   _ReduceToTraitTypeExprAction,
	{_State423, _WildcardMarker}:                   _ReduceConstrainedToGenericParameterDefAction,
	{_State424, _WildcardMarker}:                   _ReduceGenericToGenericParameterDefsAction,
	{_State425, RbracketToken}:                     _ReduceImproperToGenericParameterDefListAction,
	{_State427, _WildcardMarker}:                   _ReduceNamedTypedVarargToProperParameterDefAction,
	{_State428, _WildcardMarker}:                   _ReduceIgnoreTypedVarargToProperParameterDefAction,
	{_State431, _WildcardMarker}:                   _ReduceNamedToFieldVarPatternAction,
	{_State432, _WildcardMarker}:                   _ReduceAddToFieldVarPatternsAction,
	{_State433, _WildcardMarker}:                   _ReduceEnumVarDeclPatternToCasePatternAction,
	{_State435, _WildcardMarker}:                   _ReduceReturnableTypeExprToReturnTypeAction,
	{_State436, _WildcardMarker}:                   _ReduceAddToProperParameterDefsAction,
	{_State437, _WildcardMarker}:                   _ReduceAddExplicitToProperImportClausesAction,
	{_State438, _WildcardMarker}:                   _ReduceAddImplicitToProperImportClausesAction,
	{_State439, _WildcardMarker}:                   _ReduceToMapTypeExprAction,
	{_State440, _WildcardMarker}:                   _ReduceToArrayTypeExprAction,
	{_State441, _WildcardMarker}:                   _ReduceAddExplicitToProperExplicitFieldDefsAction,
	{_State442, _WildcardMarker}:                   _ReduceAddImplicitToProperExplicitFieldDefsAction,
	{_State443, _EndMarker}:                        _ReduceToUnsafeStatementAction,
	{_State444, _WildcardMarker}:                   _ReduceAddToProperTypeArgumentsAction,
	{_State445, _WildcardMarker}:                   _ReduceToCallExprAction,
	{_State446, _EndMarker}:                        _ReduceDoWhileToLoopExprBodyAction,
	{_State447, SemicolonToken}:                    _ReduceAssignToForAssignmentAction,
	{_State450, DoToken}:                           _ReduceSequenceExprToOptionalSequenceExprAction,
	{_State451, _EndMarker}:                        _ReduceWhileToLoopExprBodyAction,
	{_State452, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State453, IfToken}:                           _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State454, _WildcardMarker}:                   _ReduceImplicitPairToProperExplicitEnumValueDefsAction,
	{_State455, _WildcardMarker}:                   _ReduceExplicitPairToProperExplicitEnumValueDefsAction,
	{_State456, _WildcardMarker}:                   _ReduceImplicitAddToProperExplicitEnumValueDefsAction,
	{_State457, _WildcardMarker}:                   _ReduceExplicitAddToProperExplicitEnumValueDefsAction,
	{_State458, _WildcardMarker}:                   _ReduceToFuncTypeExprAction,
	{_State459, _WildcardMarker}:                   _ReduceAddToProperParameterDeclsAction,
	{_State461, _WildcardMarker}:                   _ReduceAddToProperGenericParameterDefListAction,
	{_State462, LbraceToken}:                       _ReduceNilToReturnTypeAction,
	{_State463, RparenToken}:                       _ReduceNilToParameterDefsAction,
	{_State464, _WildcardMarker}:                   _ReduceConstrainedDefToTypeDefAction,
	{_State465, _WildcardMarker}:                   _ReduceToAnonymousFuncExprAction,
	{_State467, LbraceToken}:                       _ReduceUnlabelledToOptionalLabelDeclAction,
	{_State467, DoToken}:                           _ReduceNilToOptionalSequenceExprAction,
	{_State468, LbraceToken}:                       _ReduceCaseToConditionAction,
	{_State469, _EndMarker}:                        _ReduceMultiIfElseToIfExprBodyAction,
	{_State471, _EndMarker}:                        _ReduceIfElseToIfExprBodyAction,
	{_State472, _WildcardMarker}:                   _ReduceNilToReturnTypeAction,
	{_State475, _EndMarker}:                        _ReduceIteratorToLoopExprBodyAction,
	{_State477, _WildcardMarker}:                   _ReduceToMethodSignatureAction,
	{_State478, _WildcardMarker}:                   _ReduceFuncDefToNamedFuncDefAction,
	{_State479, LbraceToken}:                       _ReduceNilToReturnTypeAction,
	{_State482, _EndMarker}:                        _ReduceForToLoopExprBodyAction,
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
      simple_statement -> State 102
      statement -> State 6
      expr_or_improper_struct_statement -> State 77
      exprs -> State 78
      callback_op -> State 72
      callback_op_statement -> State 73
      unsafe_statement -> State 106
      jump_statement -> State 86
      jump_type -> State 87
      fallthrough_statement -> State 79
      assign_statement -> State 62
      unary_op_assign_statement -> State 105
      binary_op_assign_statement -> State 68
      import_statement -> State 82
      var_decl_pattern -> State 107
      var_or_let -> State 24
      assign_pattern -> State 61
      expr -> State 76
      optional_label_decl -> State 93
      sequence_expr -> State 101
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 56
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 7
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 4:
    Kernel Items:
      #accept: ^.type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 8
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

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
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

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
      IDENTIFIER -> State 137
      LPAREN -> State 138

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
      proper_statements -> State 139
      statements -> State 141
      simple_statement -> State 102
      statement -> State 140
      expr_or_improper_struct_statement -> State 77
      exprs -> State 78
      callback_op -> State 72
      callback_op_statement -> State 73
      unsafe_statement -> State 106
      jump_statement -> State 86
      jump_type -> State 87
      fallthrough_statement -> State 79
      assign_statement -> State 62
      unary_op_assign_statement -> State 105
      binary_op_assign_statement -> State 68
      import_statement -> State 82
      var_decl_pattern -> State 107
      var_or_let -> State 24
      assign_pattern -> State 61
      expr -> State 76
      optional_label_decl -> State 93
      sequence_expr -> State 101
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 56
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
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
      statement_block -> State 142

  State 14:
    Kernel Items:
      type_def: TYPE.IDENTIFIER generic_parameter_defs type_expr
      type_def: TYPE.IDENTIFIER generic_parameter_defs type_expr IMPLEMENTS type_expr
      type_def: TYPE.IDENTIFIER ASSIGN type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 143

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
      NEWLINES -> State 144

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
      ASSIGN -> State 145

  State 24:
    Kernel Items:
      var_decl_pattern: var_or_let.var_pattern optional_type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      UNDERSCORE -> State 148
      LPAREN -> State 147
      var_pattern -> State 150
      tuple_pattern -> State 149

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
      VAR -> State 152
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      DOT -> State 151
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      case_patterns -> State 155
      var_decl_pattern -> State 107
      var_or_let -> State 24
      assign_pattern -> State 153
      case_pattern -> State 154
      optional_label_decl -> State 156
      sequence_expr -> State 157
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
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
      COLON -> State 158

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
      LPAREN -> State 159

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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      optional_label_decl -> State 156
      sequence_expr -> State 160
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
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
      STRING_LITERAL -> State 164
      IDENTIFIER -> State 162
      UNDERSCORE -> State 165
      LPAREN -> State 163
      DOT -> State 161
      import_clause -> State 166

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
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 167
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

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
      IDENTIFIER -> State 170
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
      COLON -> State 168
      ELLIPSIS -> State 169
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 174
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      proper_arguments -> State 175
      arguments -> State 172
      argument -> State 171
      colon_expr -> State 173
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
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
      LPAREN -> State 176

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
      LESS -> State 177

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
      postfix_unary_expr: accessible_expr.postfix_unary_op
    Reduce:
      * -> [postfixable_expr]
      LPAREN -> [generic_type_arguments]
    Goto:
      LBRACKET -> State 190
      DOT -> State 188
      QUESTION -> State 193
      EXCLAIM -> State 189
      DOLLAR_LBRACKET -> State 187
      ADD_ASSIGN -> State 178
      SUB_ASSIGN -> State 194
      MUL_ASSIGN -> State 192
      DIV_ASSIGN -> State 186
      MOD_ASSIGN -> State 191
      ADD_ONE_ASSIGN -> State 179
      SUB_ONE_ASSIGN -> State 195
      BIT_NEG_ASSIGN -> State 182
      BIT_AND_ASSIGN -> State 180
      BIT_OR_ASSIGN -> State 183
      BIT_XOR_ASSIGN -> State 185
      BIT_LSHIFT_ASSIGN -> State 181
      BIT_RSHIFT_ASSIGN -> State 184
      unary_op_assign -> State 199
      binary_op_assign -> State 196
      postfix_unary_op -> State 198
      generic_type_arguments -> State 197

  State 57:
    Kernel Items:
      binary_add_expr: add_expr.add_op mul_expr
      cmp_expr: add_expr., *
    Reduce:
      * -> [cmp_expr]
    Goto:
      ADD -> State 200
      SUB -> State 203
      BIT_XOR -> State 202
      BIT_OR -> State 201
      add_op -> State 204

  State 58:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      or_expr: and_expr., *
    Reduce:
      * -> [or_expr]
    Goto:
      AND -> State 205

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
      ASSIGN -> State 206

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
      optional_label_decl -> State 156
      call_expr -> State 208
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 207
      access_expr -> State 55
      index_expr -> State 83
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
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
      EQUAL -> State 209
      NOT_EQUAL -> State 214
      LESS -> State 212
      LESS_OR_EQUAL -> State 213
      GREATER -> State 210
      GREATER_OR_EQUAL -> State 211
      cmp_op -> State 215

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
      COMMA -> State 216

  State 79:
    Kernel Items:
      simple_statement: fallthrough_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 80:
    Kernel Items:
      expr: if_expr., $
    Reduce:
      $ -> [expr]
    Goto:
      (nil)

  State 81:
    Kernel Items:
      atom_expr: implicit_struct_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 82:
    Kernel Items:
      statement: import_statement., $
    Reduce:
      $ -> [statement]
    Goto:
      (nil)

  State 83:
    Kernel Items:
      accessible_expr: index_expr., *
    Reduce:
      * -> [accessible_expr]
    Goto:
      (nil)

  State 84:
    Kernel Items:
      initialize_expr: initializable_type_expr.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 217

  State 85:
    Kernel Items:
      atom_expr: initialize_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 86:
    Kernel Items:
      simple_statement: jump_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 87:
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
      JUMP_LABEL -> State 218
      LPAREN -> State 43
      LBRACKET -> State 42
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      exprs -> State 219
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 76
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 88:
    Kernel Items:
      atom_expr: literal_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 89:
    Kernel Items:
      expr: loop_expr., $
    Reduce:
      $ -> [expr]
    Goto:
      (nil)

  State 90:
    Kernel Items:
      initializable_type_expr: map_type_expr., *
    Reduce:
      * -> [initializable_type_expr]
    Goto:
      (nil)

  State 91:
    Kernel Items:
      binary_mul_expr: mul_expr.mul_op prefixable_expr
      add_expr: mul_expr., *
    Reduce:
      * -> [add_expr]
    Goto:
      MUL -> State 225
      DIV -> State 223
      MOD -> State 224
      BIT_AND -> State 220
      BIT_LSHIFT -> State 221
      BIT_RSHIFT -> State 222
      mul_op -> State 226

  State 92:
    Kernel Items:
      atom_expr: named_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 93:
    Kernel Items:
      if_expr: optional_label_decl.if_expr_body
      switch_expr: optional_label_decl.SWITCH sequence_expr statement_block
      loop_expr: optional_label_decl.loop_expr_body
      block_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    Goto:
      IF -> State 229
      SWITCH -> State 230
      FOR -> State 228
      DO -> State 227
      LBRACE -> State 11
      statement_block -> State 233
      if_expr_body -> State 231
      loop_expr_body -> State 232

  State 94:
    Kernel Items:
      sequence_expr: or_expr., *
      binary_or_expr: or_expr.OR and_expr
    Reduce:
      * -> [sequence_expr]
    Goto:
      OR -> State 234

  State 95:
    Kernel Items:
      atom_expr: parse_error_expr., *
    Reduce:
      * -> [atom_expr]
    Goto:
      (nil)

  State 96:
    Kernel Items:
      postfixable_expr: postfix_unary_expr., *
    Reduce:
      * -> [postfixable_expr]
    Goto:
      (nil)

  State 97:
    Kernel Items:
      prefixable_expr: postfixable_expr., *
    Reduce:
      * -> [prefixable_expr]
    Goto:
      (nil)

  State 98:
    Kernel Items:
      prefixable_expr: prefix_unary_expr., *
    Reduce:
      * -> [prefixable_expr]
    Goto:
      (nil)

  State 99:
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
      optional_label_decl -> State 156
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 235
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 100:
    Kernel Items:
      mul_expr: prefixable_expr., *
    Reduce:
      * -> [mul_expr]
    Goto:
      (nil)

  State 101:
    Kernel Items:
      assign_pattern: sequence_expr., ASSIGN
      expr: sequence_expr., *
    Reduce:
      * -> [expr]
      ASSIGN -> [assign_pattern]
    Goto:
      (nil)

  State 102:
    Kernel Items:
      statement: simple_statement., $
    Reduce:
      $ -> [statement]
    Goto:
      (nil)

  State 103:
    Kernel Items:
      initializable_type_expr: slice_type_expr., *
    Reduce:
      * -> [initializable_type_expr]
    Goto:
      (nil)

  State 104:
    Kernel Items:
      expr: switch_expr., $
    Reduce:
      $ -> [expr]
    Goto:
      (nil)

  State 105:
    Kernel Items:
      simple_statement: unary_op_assign_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 106:
    Kernel Items:
      simple_statement: unsafe_statement., $
    Reduce:
      $ -> [simple_statement]
    Goto:
      (nil)

  State 107:
    Kernel Items:
      sequence_expr: var_decl_pattern., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 108:
    Kernel Items:
      call_expr: accessible_expr.generic_type_arguments LPAREN arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
      postfixable_expr: accessible_expr., *
      postfix_unary_expr: accessible_expr.postfix_unary_op
    Reduce:
      * -> [postfixable_expr]
      LPAREN -> [generic_type_arguments]
    Goto:
      LBRACKET -> State 190
      DOT -> State 188
      QUESTION -> State 193
      EXCLAIM -> State 189
      DOLLAR_LBRACKET -> State 187
      postfix_unary_op -> State 198
      generic_type_arguments -> State 197

  State 109:
    Kernel Items:
      expr: sequence_expr., $
    Reduce:
      $ -> [expr]
    Goto:
      (nil)

  State 110:
    Kernel Items:
      prefix_unary_type_op: BIT_AND., *
    Reduce:
      * -> [prefix_unary_type_op]
    Goto:
      (nil)

  State 111:
    Kernel Items:
      prefix_unary_type_op: BIT_NEG., *
    Reduce:
      * -> [prefix_unary_type_op]
    Goto:
      (nil)

  State 112:
    Kernel Items:
      inferred_type_expr: DOT., *
    Reduce:
      * -> [inferred_type_expr]
    Goto:
      (nil)

  State 113:
    Kernel Items:
      explicit_enum_type_expr: ENUM.LPAREN explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 236

  State 114:
    Kernel Items:
      prefix_unary_type_op: EXCLAIM., *
    Reduce:
      * -> [prefix_unary_type_op]
    Goto:
      (nil)

  State 115:
    Kernel Items:
      func_type_expr: FUNC.LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 237

  State 116:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_type_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_type_arguments
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      DOT -> State 238
      DOLLAR_LBRACKET -> State 187
      generic_type_arguments -> State 239

  State 117:
    Kernel Items:
      implicit_struct_type_expr: LPAREN.implicit_field_defs RPAREN
      implicit_enum_type_expr: LPAREN.implicit_enum_value_defs RPAREN
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      IDENTIFIER -> State 241
      UNDERSCORE -> State 242
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 240
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      unsafe_statement -> State 250
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      binary_type_expr -> State 124
      field_def -> State 243
      proper_implicit_field_defs -> State 248
      implicit_field_defs -> State 245
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      proper_implicit_enum_value_defs -> State 247
      implicit_enum_value_defs -> State 244
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126
      method_signature -> State 246

  State 118:
    Kernel Items:
      parse_error_type_expr: PARSE_ERROR., *
    Reduce:
      * -> [parse_error_type_expr]
    Goto:
      (nil)

  State 119:
    Kernel Items:
      prefix_unary_type_op: QUESTION., *
    Reduce:
      * -> [prefix_unary_type_op]
    Goto:
      (nil)

  State 120:
    Kernel Items:
      prefix_unary_type_op: TILDE_TILDE., *
    Reduce:
      * -> [prefix_unary_type_op]
    Goto:
      (nil)

  State 121:
    Kernel Items:
      trait_type_expr: TRAIT.LPAREN explicit_field_defs RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 251

  State 122:
    Kernel Items:
      inferred_type_expr: UNDERSCORE., *
    Reduce:
      * -> [inferred_type_expr]
    Goto:
      (nil)

  State 123:
    Kernel Items:
      returnable_type_expr: atom_type_expr., *
    Reduce:
      * -> [returnable_type_expr]
    Goto:
      (nil)

  State 124:
    Kernel Items:
      type_expr: binary_type_expr., *
    Reduce:
      * -> [type_expr]
    Goto:
      (nil)

  State 125:
    Kernel Items:
      atom_type_expr: explicit_enum_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 126:
    Kernel Items:
      atom_type_expr: func_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 127:
    Kernel Items:
      atom_type_expr: implicit_enum_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 128:
    Kernel Items:
      atom_type_expr: implicit_struct_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 129:
    Kernel Items:
      atom_type_expr: inferred_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 130:
    Kernel Items:
      atom_type_expr: initializable_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 131:
    Kernel Items:
      atom_type_expr: named_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 132:
    Kernel Items:
      atom_type_expr: parse_error_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 133:
    Kernel Items:
      returnable_type_expr: prefix_unary_type_expr., *
    Reduce:
      * -> [returnable_type_expr]
    Goto:
      (nil)

  State 134:
    Kernel Items:
      prefix_unary_type_expr: prefix_unary_type_op.returnable_type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 252
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 135:
    Kernel Items:
      type_expr: returnable_type_expr., *
    Reduce:
      * -> [type_expr]
    Goto:
      (nil)

  State 136:
    Kernel Items:
      atom_type_expr: trait_type_expr., *
    Reduce:
      * -> [atom_type_expr]
    Goto:
      (nil)

  State 137:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER.generic_parameter_defs LPAREN parameter_defs RPAREN return_type statement_block
      named_func_def: FUNC IDENTIFIER.ASSIGN expr
    Reduce:
      LPAREN -> [generic_parameter_defs]
    Goto:
      DOLLAR_LBRACKET -> State 258
      ASSIGN -> State 257
      generic_parameter_defs -> State 259

  State 138:
    Kernel Items:
      named_func_def: FUNC LPAREN.parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 260
      UNDERSCORE -> State 261
      proper_parameter_def -> State 263
      parameter_def -> State 262

  State 139:
    Kernel Items:
      proper_statements: proper_statements.NEWLINES statement
      proper_statements: proper_statements.SEMICOLON statement
      statements: proper_statements., RBRACE
      statements: proper_statements.NEWLINES
      statements: proper_statements.SEMICOLON
    Reduce:
      RBRACE -> [statements]
    Goto:
      NEWLINES -> State 264
      SEMICOLON -> State 265

  State 140:
    Kernel Items:
      proper_statements: statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 141:
    Kernel Items:
      statement_block: LBRACE statements.RBRACE
    Reduce:
      (nil)
    Goto:
      RBRACE -> State 266

  State 142:
    Kernel Items:
      package_def: PACKAGE statement_block., *
    Reduce:
      * -> [package_def]
    Goto:
      (nil)

  State 143:
    Kernel Items:
      type_def: TYPE IDENTIFIER.generic_parameter_defs type_expr
      type_def: TYPE IDENTIFIER.generic_parameter_defs type_expr IMPLEMENTS type_expr
      type_def: TYPE IDENTIFIER.ASSIGN type_expr
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      DOLLAR_LBRACKET -> State 258
      ASSIGN -> State 267
      generic_parameter_defs -> State 268

  State 144:
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
      definition -> State 269
      statement_block -> State 21
      var_decl_pattern -> State 23
      var_or_let -> State 24
      type_def -> State 22
      named_func_def -> State 18
      package_def -> State 19

  State 145:
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 270
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 146:
    Kernel Items:
      var_pattern: IDENTIFIER., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 147:
    Kernel Items:
      tuple_pattern: LPAREN.field_var_patterns RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 272
      UNDERSCORE -> State 148
      LPAREN -> State 147
      ELLIPSIS -> State 271
      var_pattern -> State 275
      tuple_pattern -> State 149
      field_var_patterns -> State 274
      field_var_pattern -> State 273

  State 148:
    Kernel Items:
      var_pattern: UNDERSCORE., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 149:
    Kernel Items:
      var_pattern: tuple_pattern., *
    Reduce:
      * -> [var_pattern]
    Goto:
      (nil)

  State 150:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern.optional_type_expr
    Reduce:
      * -> [optional_type_expr]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      optional_type_expr -> State 276
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 277
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 151:
    Kernel Items:
      case_pattern: DOT.IDENTIFIER implicit_struct_expr
      case_pattern: DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 278

  State 152:
    Kernel Items:
      var_or_let: VAR., *
      case_pattern: VAR.DOT IDENTIFIER tuple_pattern
    Reduce:
      * -> [var_or_let]
    Goto:
      DOT -> State 279

  State 153:
    Kernel Items:
      case_pattern: assign_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 154:
    Kernel Items:
      case_patterns: case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 155:
    Kernel Items:
      statement: CASE case_patterns.COLON optional_simple_statement
      case_patterns: case_patterns.COMMA case_pattern
    Reduce:
      (nil)
    Goto:
      COMMA -> State 281
      COLON -> State 280

  State 156:
    Kernel Items:
      block_expr: optional_label_decl.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 233

  State 157:
    Kernel Items:
      assign_pattern: sequence_expr., *
    Reduce:
      * -> [assign_pattern]
    Goto:
      (nil)

  State 158:
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
      simple_statement -> State 283
      optional_simple_statement -> State 282
      expr_or_improper_struct_statement -> State 77
      exprs -> State 78
      callback_op -> State 72
      callback_op_statement -> State 73
      unsafe_statement -> State 106
      jump_statement -> State 86
      jump_type -> State 87
      fallthrough_statement -> State 79
      assign_statement -> State 62
      unary_op_assign_statement -> State 105
      binary_op_assign_statement -> State 68
      var_decl_pattern -> State 107
      var_or_let -> State 24
      assign_pattern -> State 61
      expr -> State 76
      optional_label_decl -> State 93
      sequence_expr -> State 101
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 56
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 159:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 260
      UNDERSCORE -> State 261
      proper_parameter_def -> State 263
      parameter_def -> State 284
      proper_parameter_defs -> State 286
      parameter_defs -> State 285

  State 160:
    Kernel Items:
      sequence_expr: GREATER sequence_expr., $
    Reduce:
      $ -> [sequence_expr]
    Goto:
      (nil)

  State 161:
    Kernel Items:
      import_clause: DOT.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 287

  State 162:
    Kernel Items:
      import_clause: IDENTIFIER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 288

  State 163:
    Kernel Items:
      import_statement: IMPORT LPAREN.import_clauses RPAREN
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 164
      IDENTIFIER -> State 162
      UNDERSCORE -> State 165
      DOT -> State 161
      import_clause -> State 289
      proper_import_clauses -> State 291
      import_clauses -> State 290

  State 164:
    Kernel Items:
      import_clause: STRING_LITERAL., $
    Reduce:
      $ -> [import_clause]
    Goto:
      (nil)

  State 165:
    Kernel Items:
      import_clause: UNDERSCORE.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 292

  State 166:
    Kernel Items:
      import_statement: IMPORT import_clause., $
    Reduce:
      $ -> [import_statement]
    Goto:
      (nil)

  State 167:
    Kernel Items:
      slice_type_expr: LBRACKET type_expr.RBRACKET
      array_type_expr: LBRACKET type_expr.COMMA INTEGER_LITERAL RBRACKET
      map_type_expr: LBRACKET type_expr.COLON type_expr RBRACKET
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 295
      COMMA -> State 294
      COLON -> State 293
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 168:
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 296
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 169:
    Kernel Items:
      argument: ELLIPSIS., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 170:
    Kernel Items:
      argument: IDENTIFIER.ASSIGN expr
      named_expr: IDENTIFIER., *
    Reduce:
      * -> [named_expr]
    Goto:
      ASSIGN -> State 297

  State 171:
    Kernel Items:
      proper_arguments: argument., *
    Reduce:
      * -> [proper_arguments]
    Goto:
      (nil)

  State 172:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 298

  State 173:
    Kernel Items:
      argument: colon_expr., *
      colon_expr: colon_expr.COLON
      colon_expr: colon_expr.COLON expr
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 299

  State 174:
    Kernel Items:
      argument: expr., *
      argument: expr.ELLIPSIS
      colon_expr: expr.COLON
      colon_expr: expr.COLON expr
    Reduce:
      * -> [argument]
    Goto:
      COLON -> State 300
      ELLIPSIS -> State 301

  State 175:
    Kernel Items:
      proper_arguments: proper_arguments.COMMA argument
      arguments: proper_arguments., RPAREN
      arguments: proper_arguments.COMMA
    Reduce:
      RPAREN -> [arguments]
    Goto:
      COMMA -> State 302

  State 176:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN.explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 241
      UNDERSCORE -> State 242
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 240
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      unsafe_statement -> State 250
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      binary_type_expr -> State 124
      field_def -> State 304
      implicit_struct_type_expr -> State 128
      proper_explicit_field_defs -> State 305
      explicit_field_defs -> State 303
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126
      method_signature -> State 246

  State 177:
    Kernel Items:
      unsafe_statement: UNSAFE LESS.IDENTIFIER GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 306

  State 178:
    Kernel Items:
      binary_op_assign: ADD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 179:
    Kernel Items:
      unary_op_assign: ADD_ONE_ASSIGN., $
    Reduce:
      $ -> [unary_op_assign]
    Goto:
      (nil)

  State 180:
    Kernel Items:
      binary_op_assign: BIT_AND_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 181:
    Kernel Items:
      binary_op_assign: BIT_LSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 182:
    Kernel Items:
      binary_op_assign: BIT_NEG_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 183:
    Kernel Items:
      binary_op_assign: BIT_OR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 184:
    Kernel Items:
      binary_op_assign: BIT_RSHIFT_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 185:
    Kernel Items:
      binary_op_assign: BIT_XOR_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 186:
    Kernel Items:
      binary_op_assign: DIV_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 187:
    Kernel Items:
      generic_type_arguments: DOLLAR_LBRACKET.type_arguments RBRACKET
    Reduce:
      RBRACKET -> [type_arguments]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 309
      binary_type_expr -> State 124
      proper_type_arguments -> State 307
      type_arguments -> State 308
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 188:
    Kernel Items:
      access_expr: accessible_expr DOT.IDENTIFIER
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 310

  State 189:
    Kernel Items:
      postfix_unary_op: EXCLAIM., *
    Reduce:
      * -> [postfix_unary_op]
    Goto:
      (nil)

  State 190:
    Kernel Items:
      index_expr: accessible_expr LBRACKET.argument RBRACKET
    Reduce:
      * -> [optional_label_decl]
    Goto:
      INTEGER_LITERAL -> State 40
      FLOAT_LITERAL -> State 35
      RUNE_LITERAL -> State 48
      STRING_LITERAL -> State 49
      IDENTIFIER -> State 170
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
      COLON -> State 168
      ELLIPSIS -> State 169
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 174
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      argument -> State 311
      colon_expr -> State 173
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 191:
    Kernel Items:
      binary_op_assign: MOD_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 192:
    Kernel Items:
      binary_op_assign: MUL_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 193:
    Kernel Items:
      postfix_unary_op: QUESTION., *
    Reduce:
      * -> [postfix_unary_op]
    Goto:
      (nil)

  State 194:
    Kernel Items:
      binary_op_assign: SUB_ASSIGN., *
    Reduce:
      * -> [binary_op_assign]
    Goto:
      (nil)

  State 195:
    Kernel Items:
      unary_op_assign: SUB_ONE_ASSIGN., $
    Reduce:
      $ -> [unary_op_assign]
    Goto:
      (nil)

  State 196:
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 312
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 197:
    Kernel Items:
      call_expr: accessible_expr generic_type_arguments.LPAREN arguments RPAREN
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 313

  State 198:
    Kernel Items:
      postfix_unary_expr: accessible_expr postfix_unary_op., *
    Reduce:
      * -> [postfix_unary_expr]
    Goto:
      (nil)

  State 199:
    Kernel Items:
      unary_op_assign_statement: accessible_expr unary_op_assign., $
    Reduce:
      $ -> [unary_op_assign_statement]
    Goto:
      (nil)

  State 200:
    Kernel Items:
      add_op: ADD., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 201:
    Kernel Items:
      add_op: BIT_OR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 202:
    Kernel Items:
      add_op: BIT_XOR., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 203:
    Kernel Items:
      add_op: SUB., *
    Reduce:
      * -> [add_op]
    Goto:
      (nil)

  State 204:
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
      optional_label_decl -> State 156
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 314
      binary_mul_expr -> State 67
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 205:
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
      optional_label_decl -> State 156
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 315
      binary_cmp_expr -> State 66
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 206:
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 316
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 207:
    Kernel Items:
      call_expr: accessible_expr.generic_type_arguments LPAREN arguments RPAREN
      access_expr: accessible_expr.DOT IDENTIFIER
      index_expr: accessible_expr.LBRACKET argument RBRACKET
    Reduce:
      LPAREN -> [generic_type_arguments]
    Goto:
      LBRACKET -> State 190
      DOT -> State 188
      DOLLAR_LBRACKET -> State 187
      generic_type_arguments -> State 197

  State 208:
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

  State 209:
    Kernel Items:
      cmp_op: EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 210:
    Kernel Items:
      cmp_op: GREATER., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 211:
    Kernel Items:
      cmp_op: GREATER_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 212:
    Kernel Items:
      cmp_op: LESS., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 213:
    Kernel Items:
      cmp_op: LESS_OR_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 214:
    Kernel Items:
      cmp_op: NOT_EQUAL., *
    Reduce:
      * -> [cmp_op]
    Goto:
      (nil)

  State 215:
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
      optional_label_decl -> State 156
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 317
      binary_add_expr -> State 64
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 216:
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 318
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 217:
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
      IDENTIFIER -> State 170
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
      COLON -> State 168
      ELLIPSIS -> State 169
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 174
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      proper_arguments -> State 175
      arguments -> State 319
      argument -> State 171
      colon_expr -> State 173
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 218:
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
      exprs -> State 320
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 76
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 219:
    Kernel Items:
      exprs: exprs.COMMA expr
      jump_statement: jump_type exprs., $
    Reduce:
      $ -> [jump_statement]
    Goto:
      COMMA -> State 216

  State 220:
    Kernel Items:
      mul_op: BIT_AND., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 221:
    Kernel Items:
      mul_op: BIT_LSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 222:
    Kernel Items:
      mul_op: BIT_RSHIFT., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 223:
    Kernel Items:
      mul_op: DIV., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 224:
    Kernel Items:
      mul_op: MOD., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 225:
    Kernel Items:
      mul_op: MUL., *
    Reduce:
      * -> [mul_op]
    Goto:
      (nil)

  State 226:
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
      optional_label_decl -> State 156
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 321
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 227:
    Kernel Items:
      loop_expr_body: DO.statement_block
      loop_expr_body: DO.statement_block FOR sequence_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 322

  State 228:
    Kernel Items:
      loop_expr_body: FOR.sequence_expr DO statement_block
      loop_expr_body: FOR.assign_pattern IN sequence_expr DO statement_block
      loop_expr_body: FOR.for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      assign_pattern -> State 323
      optional_label_decl -> State 156
      sequence_expr -> State 325
      for_assignment -> State 324
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 229:
    Kernel Items:
      if_expr_body: IF.condition statement_block
      if_expr_body: IF.condition statement_block ELSE statement_block
      if_expr_body: IF.condition statement_block ELSE if_expr
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
      CASE -> State 326
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      optional_label_decl -> State 156
      sequence_expr -> State 328
      condition -> State 327
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 230:
    Kernel Items:
      switch_expr: optional_label_decl SWITCH.sequence_expr statement_block
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      optional_label_decl -> State 156
      sequence_expr -> State 329
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 231:
    Kernel Items:
      if_expr: optional_label_decl if_expr_body., $
    Reduce:
      $ -> [if_expr]
    Goto:
      (nil)

  State 232:
    Kernel Items:
      loop_expr: optional_label_decl loop_expr_body., $
    Reduce:
      $ -> [loop_expr]
    Goto:
      (nil)

  State 233:
    Kernel Items:
      block_expr: optional_label_decl statement_block., *
    Reduce:
      * -> [block_expr]
    Goto:
      (nil)

  State 234:
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
      optional_label_decl -> State 156
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 330
      binary_and_expr -> State 65
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 235:
    Kernel Items:
      prefix_unary_expr: prefix_unary_op prefixable_expr., *
    Reduce:
      * -> [prefix_unary_expr]
    Goto:
      (nil)

  State 236:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN.explicit_enum_value_defs RPAREN
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 241
      UNDERSCORE -> State 242
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 240
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      unsafe_statement -> State 250
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      binary_type_expr -> State 124
      field_def -> State 332
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      proper_explicit_enum_value_defs -> State 333
      explicit_enum_value_defs -> State 331
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126
      method_signature -> State 246

  State 237:
    Kernel Items:
      func_type_expr: FUNC LPAREN.parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 335
      UNDERSCORE -> State 336
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      ELLIPSIS -> State 334
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 341
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      proper_parameter_def -> State 340
      parameter_decl -> State 337
      proper_parameter_decls -> State 339
      parameter_decls -> State 338
      func_type_expr -> State 126

  State 238:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT.IDENTIFIER generic_type_arguments
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 342

  State 239:
    Kernel Items:
      named_type_expr: IDENTIFIER generic_type_arguments., *
    Reduce:
      * -> [named_type_expr]
    Goto:
      (nil)

  State 240:
    Kernel Items:
      func_type_expr: FUNC.LPAREN parameter_decls RPAREN return_type
      method_signature: FUNC.IDENTIFIER LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 343
      LPAREN -> State 237

  State 241:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_type_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_type_arguments
      field_def: IDENTIFIER.type_expr optional_default
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 344
      QUESTION -> State 119
      EXCLAIM -> State 114
      DOLLAR_LBRACKET -> State 187
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 345
      binary_type_expr -> State 124
      generic_type_arguments -> State 239
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 242:
    Kernel Items:
      inferred_type_expr: UNDERSCORE., *
      field_def: UNDERSCORE.type_expr
    Reduce:
      * -> [inferred_type_expr]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 346
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 243:
    Kernel Items:
      proper_implicit_field_defs: field_def., *
      proper_implicit_enum_value_defs: field_def.OR field_def
    Reduce:
      * -> [proper_implicit_field_defs]
    Goto:
      OR -> State 347

  State 244:
    Kernel Items:
      implicit_enum_type_expr: LPAREN implicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 348

  State 245:
    Kernel Items:
      implicit_struct_type_expr: LPAREN implicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 349

  State 246:
    Kernel Items:
      field_def: method_signature., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 247:
    Kernel Items:
      proper_implicit_enum_value_defs: proper_implicit_enum_value_defs.OR field_def
      implicit_enum_value_defs: proper_implicit_enum_value_defs., RPAREN
      implicit_enum_value_defs: proper_implicit_enum_value_defs.NEWLINES
    Reduce:
      RPAREN -> [implicit_enum_value_defs]
    Goto:
      NEWLINES -> State 350
      OR -> State 351

  State 248:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs.COMMA field_def
      implicit_field_defs: proper_implicit_field_defs., RPAREN
      implicit_field_defs: proper_implicit_field_defs.COMMA
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      COMMA -> State 352

  State 249:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      field_def: type_expr.optional_default
    Reduce:
      * -> [optional_default]
    Goto:
      ASSIGN -> State 353
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256
      optional_default -> State 354

  State 250:
    Kernel Items:
      field_def: unsafe_statement., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 251:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN.explicit_field_defs RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 241
      UNDERSCORE -> State 242
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 240
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      unsafe_statement -> State 250
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      binary_type_expr -> State 124
      field_def -> State 304
      implicit_struct_type_expr -> State 128
      proper_explicit_field_defs -> State 305
      explicit_field_defs -> State 355
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126
      method_signature -> State 246

  State 252:
    Kernel Items:
      prefix_unary_type_expr: prefix_unary_type_op returnable_type_expr., *
    Reduce:
      * -> [prefix_unary_type_expr]
    Goto:
      (nil)

  State 253:
    Kernel Items:
      binary_type_op: ADD., *
    Reduce:
      * -> [binary_type_op]
    Goto:
      (nil)

  State 254:
    Kernel Items:
      binary_type_op: MUL., *
    Reduce:
      * -> [binary_type_op]
    Goto:
      (nil)

  State 255:
    Kernel Items:
      binary_type_op: SUB., *
    Reduce:
      * -> [binary_type_op]
    Goto:
      (nil)

  State 256:
    Kernel Items:
      binary_type_expr: type_expr binary_type_op.returnable_type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 356
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 257:
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 357
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 258:
    Kernel Items:
      generic_parameter_defs: DOLLAR_LBRACKET.generic_parameter_def_list RBRACKET
    Reduce:
      RBRACKET -> [generic_parameter_def_list]
    Goto:
      IDENTIFIER -> State 358
      generic_parameter_def -> State 359
      proper_generic_parameter_def_list -> State 361
      generic_parameter_def_list -> State 360

  State 259:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 362

  State 260:
    Kernel Items:
      proper_parameter_def: IDENTIFIER.type_expr
      proper_parameter_def: IDENTIFIER.ELLIPSIS type_expr
      proper_parameter_def: IDENTIFIER.ELLIPSIS
      parameter_def: IDENTIFIER., *
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      ELLIPSIS -> State 363
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 364
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 261:
    Kernel Items:
      proper_parameter_def: UNDERSCORE.type_expr
      proper_parameter_def: UNDERSCORE.ELLIPSIS
      proper_parameter_def: UNDERSCORE.ELLIPSIS type_expr
      parameter_def: UNDERSCORE., *
    Reduce:
      * -> [parameter_def]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      ELLIPSIS -> State 365
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 366
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 262:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def.RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 367

  State 263:
    Kernel Items:
      parameter_def: proper_parameter_def., *
    Reduce:
      * -> [parameter_def]
    Goto:
      (nil)

  State 264:
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
      simple_statement -> State 102
      statement -> State 368
      expr_or_improper_struct_statement -> State 77
      exprs -> State 78
      callback_op -> State 72
      callback_op_statement -> State 73
      unsafe_statement -> State 106
      jump_statement -> State 86
      jump_type -> State 87
      fallthrough_statement -> State 79
      assign_statement -> State 62
      unary_op_assign_statement -> State 105
      binary_op_assign_statement -> State 68
      import_statement -> State 82
      var_decl_pattern -> State 107
      var_or_let -> State 24
      assign_pattern -> State 61
      expr -> State 76
      optional_label_decl -> State 93
      sequence_expr -> State 101
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 56
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 265:
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
      simple_statement -> State 102
      statement -> State 369
      expr_or_improper_struct_statement -> State 77
      exprs -> State 78
      callback_op -> State 72
      callback_op_statement -> State 73
      unsafe_statement -> State 106
      jump_statement -> State 86
      jump_type -> State 87
      fallthrough_statement -> State 79
      assign_statement -> State 62
      unary_op_assign_statement -> State 105
      binary_op_assign_statement -> State 68
      import_statement -> State 82
      var_decl_pattern -> State 107
      var_or_let -> State 24
      assign_pattern -> State 61
      expr -> State 76
      optional_label_decl -> State 93
      sequence_expr -> State 101
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 56
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 266:
    Kernel Items:
      statement_block: LBRACE statements RBRACE., $
    Reduce:
      $ -> [statement_block]
    Goto:
      (nil)

  State 267:
    Kernel Items:
      type_def: TYPE IDENTIFIER ASSIGN.type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 370
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 268:
    Kernel Items:
      type_def: TYPE IDENTIFIER generic_parameter_defs.type_expr
      type_def: TYPE IDENTIFIER generic_parameter_defs.type_expr IMPLEMENTS type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 371
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 269:
    Kernel Items:
      proper_definitions: proper_definitions NEWLINES definition., *
    Reduce:
      * -> [proper_definitions]
    Goto:
      (nil)

  State 270:
    Kernel Items:
      definition: var_decl_pattern ASSIGN expr., *
    Reduce:
      * -> [definition]
    Goto:
      (nil)

  State 271:
    Kernel Items:
      field_var_pattern: ELLIPSIS., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 272:
    Kernel Items:
      var_pattern: IDENTIFIER., *
      field_var_pattern: IDENTIFIER.ASSIGN var_pattern
    Reduce:
      * -> [var_pattern]
    Goto:
      ASSIGN -> State 372

  State 273:
    Kernel Items:
      field_var_patterns: field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 274:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns.RPAREN
      field_var_patterns: field_var_patterns.COMMA field_var_pattern
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 374
      COMMA -> State 373

  State 275:
    Kernel Items:
      field_var_pattern: var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 276:
    Kernel Items:
      var_decl_pattern: var_or_let var_pattern optional_type_expr., $
    Reduce:
      $ -> [var_decl_pattern]
    Goto:
      (nil)

  State 277:
    Kernel Items:
      optional_type_expr: type_expr., *
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      * -> [optional_type_expr]
    Goto:
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 278:
    Kernel Items:
      case_pattern: DOT IDENTIFIER.implicit_struct_expr
      case_pattern: DOT IDENTIFIER., *
    Reduce:
      * -> [case_pattern]
    Goto:
      LPAREN -> State 43
      implicit_struct_expr -> State 375

  State 279:
    Kernel Items:
      case_pattern: VAR DOT.IDENTIFIER tuple_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 376

  State 280:
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
      simple_statement -> State 283
      optional_simple_statement -> State 377
      expr_or_improper_struct_statement -> State 77
      exprs -> State 78
      callback_op -> State 72
      callback_op_statement -> State 73
      unsafe_statement -> State 106
      jump_statement -> State 86
      jump_type -> State 87
      fallthrough_statement -> State 79
      assign_statement -> State 62
      unary_op_assign_statement -> State 105
      binary_op_assign_statement -> State 68
      var_decl_pattern -> State 107
      var_or_let -> State 24
      assign_pattern -> State 61
      expr -> State 76
      optional_label_decl -> State 93
      sequence_expr -> State 101
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 56
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 281:
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
      VAR -> State 152
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      DOT -> State 151
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 107
      var_or_let -> State 24
      assign_pattern -> State 153
      case_pattern -> State 378
      optional_label_decl -> State 156
      sequence_expr -> State 157
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 282:
    Kernel Items:
      statement: DEFAULT COLON optional_simple_statement., $
    Reduce:
      $ -> [statement]
    Goto:
      (nil)

  State 283:
    Kernel Items:
      optional_simple_statement: simple_statement., $
    Reduce:
      $ -> [optional_simple_statement]
    Goto:
      (nil)

  State 284:
    Kernel Items:
      proper_parameter_defs: parameter_def., *
    Reduce:
      * -> [proper_parameter_defs]
    Goto:
      (nil)

  State 285:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 379

  State 286:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs.COMMA parameter_def
      parameter_defs: proper_parameter_defs., RPAREN
      parameter_defs: proper_parameter_defs.COMMA
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      COMMA -> State 380

  State 287:
    Kernel Items:
      import_clause: DOT STRING_LITERAL., $
    Reduce:
      $ -> [import_clause]
    Goto:
      (nil)

  State 288:
    Kernel Items:
      import_clause: IDENTIFIER STRING_LITERAL., $
    Reduce:
      $ -> [import_clause]
    Goto:
      (nil)

  State 289:
    Kernel Items:
      proper_import_clauses: import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 290:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 381

  State 291:
    Kernel Items:
      proper_import_clauses: proper_import_clauses.NEWLINES import_clause
      proper_import_clauses: proper_import_clauses.COMMA import_clause
      import_clauses: proper_import_clauses., RPAREN
      import_clauses: proper_import_clauses.NEWLINES
      import_clauses: proper_import_clauses.COMMA
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      NEWLINES -> State 383
      COMMA -> State 382

  State 292:
    Kernel Items:
      import_clause: UNDERSCORE STRING_LITERAL., $
    Reduce:
      $ -> [import_clause]
    Goto:
      (nil)

  State 293:
    Kernel Items:
      map_type_expr: LBRACKET type_expr COLON.type_expr RBRACKET
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 384
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 294:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA.INTEGER_LITERAL RBRACKET
    Reduce:
      (nil)
    Goto:
      INTEGER_LITERAL -> State 385

  State 295:
    Kernel Items:
      slice_type_expr: LBRACKET type_expr RBRACKET., *
    Reduce:
      * -> [slice_type_expr]
    Goto:
      (nil)

  State 296:
    Kernel Items:
      colon_expr: COLON expr., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 297:
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 386
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 298:
    Kernel Items:
      implicit_struct_expr: LPAREN arguments RPAREN., *
    Reduce:
      * -> [implicit_struct_expr]
    Goto:
      (nil)

  State 299:
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 387
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 300:
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 388
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 301:
    Kernel Items:
      argument: expr ELLIPSIS., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 302:
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
      IDENTIFIER -> State 170
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
      COLON -> State 168
      ELLIPSIS -> State 169
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 174
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      argument -> State 389
      colon_expr -> State 173
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 303:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 390

  State 304:
    Kernel Items:
      proper_explicit_field_defs: field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 305:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs.NEWLINES field_def
      proper_explicit_field_defs: proper_explicit_field_defs.COMMA field_def
      explicit_field_defs: proper_explicit_field_defs., RPAREN
      explicit_field_defs: proper_explicit_field_defs.NEWLINES
      explicit_field_defs: proper_explicit_field_defs.COMMA
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      NEWLINES -> State 392
      COMMA -> State 391

  State 306:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER.GREATER STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      GREATER -> State 393

  State 307:
    Kernel Items:
      proper_type_arguments: proper_type_arguments.COMMA type_expr
      type_arguments: proper_type_arguments., RBRACKET
      type_arguments: proper_type_arguments.COMMA
    Reduce:
      RBRACKET -> [type_arguments]
    Goto:
      COMMA -> State 394

  State 308:
    Kernel Items:
      generic_type_arguments: DOLLAR_LBRACKET type_arguments.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 395

  State 309:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_type_arguments: type_expr., *
    Reduce:
      * -> [proper_type_arguments]
    Goto:
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 310:
    Kernel Items:
      access_expr: accessible_expr DOT IDENTIFIER., *
    Reduce:
      * -> [access_expr]
    Goto:
      (nil)

  State 311:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 396

  State 312:
    Kernel Items:
      binary_op_assign_statement: accessible_expr binary_op_assign expr., $
    Reduce:
      $ -> [binary_op_assign_statement]
    Goto:
      (nil)

  State 313:
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
      IDENTIFIER -> State 170
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
      COLON -> State 168
      ELLIPSIS -> State 169
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      var_decl_pattern -> State 107
      var_or_let -> State 24
      expr -> State 174
      optional_label_decl -> State 93
      sequence_expr -> State 109
      if_expr -> State 80
      switch_expr -> State 104
      loop_expr -> State 89
      call_expr -> State 71
      proper_arguments -> State 175
      arguments -> State 397
      argument -> State 171
      colon_expr -> State 173
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 314:
    Kernel Items:
      binary_mul_expr: mul_expr.mul_op prefixable_expr
      binary_add_expr: add_expr add_op mul_expr., *
    Reduce:
      * -> [binary_add_expr]
    Goto:
      MUL -> State 225
      DIV -> State 223
      MOD -> State 224
      BIT_AND -> State 220
      BIT_LSHIFT -> State 221
      BIT_RSHIFT -> State 222
      mul_op -> State 226

  State 315:
    Kernel Items:
      binary_cmp_expr: cmp_expr.cmp_op add_expr
      binary_and_expr: and_expr AND cmp_expr., *
    Reduce:
      * -> [binary_and_expr]
    Goto:
      EQUAL -> State 209
      NOT_EQUAL -> State 214
      LESS -> State 212
      LESS_OR_EQUAL -> State 213
      GREATER -> State 210
      GREATER_OR_EQUAL -> State 211
      cmp_op -> State 215

  State 316:
    Kernel Items:
      assign_statement: assign_pattern ASSIGN expr., $
    Reduce:
      $ -> [assign_statement]
    Goto:
      (nil)

  State 317:
    Kernel Items:
      binary_add_expr: add_expr.add_op mul_expr
      binary_cmp_expr: cmp_expr cmp_op add_expr., *
    Reduce:
      * -> [binary_cmp_expr]
    Goto:
      ADD -> State 200
      SUB -> State 203
      BIT_XOR -> State 202
      BIT_OR -> State 201
      add_op -> State 204

  State 318:
    Kernel Items:
      exprs: exprs COMMA expr., *
    Reduce:
      * -> [exprs]
    Goto:
      (nil)

  State 319:
    Kernel Items:
      initialize_expr: initializable_type_expr LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 398

  State 320:
    Kernel Items:
      exprs: exprs.COMMA expr
      jump_statement: jump_type JUMP_LABEL exprs., $
    Reduce:
      $ -> [jump_statement]
    Goto:
      COMMA -> State 216

  State 321:
    Kernel Items:
      binary_mul_expr: mul_expr mul_op prefixable_expr., *
    Reduce:
      * -> [binary_mul_expr]
    Goto:
      (nil)

  State 322:
    Kernel Items:
      loop_expr_body: DO statement_block., *
      loop_expr_body: DO statement_block.FOR sequence_expr
    Reduce:
      * -> [loop_expr_body]
    Goto:
      FOR -> State 399

  State 323:
    Kernel Items:
      loop_expr_body: FOR assign_pattern.IN sequence_expr DO statement_block
      for_assignment: assign_pattern.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      IN -> State 401
      ASSIGN -> State 400

  State 324:
    Kernel Items:
      loop_expr_body: FOR for_assignment.SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 402

  State 325:
    Kernel Items:
      assign_pattern: sequence_expr., *
      loop_expr_body: FOR sequence_expr.DO statement_block
      for_assignment: sequence_expr., SEMICOLON
    Reduce:
      * -> [assign_pattern]
      SEMICOLON -> [for_assignment]
    Goto:
      DO -> State 403

  State 326:
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
      VAR -> State 152
      LET -> State 12
      NOT -> State 45
      LABEL_DECL -> State 41
      LPAREN -> State 43
      LBRACKET -> State 42
      DOT -> State 151
      SUB -> State 51
      MUL -> State 44
      BIT_NEG -> State 27
      BIT_AND -> State 26
      GREATER -> State 37
      PARSE_ERROR -> State 46
      case_patterns -> State 404
      var_decl_pattern -> State 107
      var_or_let -> State 24
      assign_pattern -> State 153
      case_pattern -> State 154
      optional_label_decl -> State 156
      sequence_expr -> State 157
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 327:
    Kernel Items:
      if_expr_body: IF condition.statement_block
      if_expr_body: IF condition.statement_block ELSE statement_block
      if_expr_body: IF condition.statement_block ELSE if_expr
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 405

  State 328:
    Kernel Items:
      condition: sequence_expr., LBRACE
    Reduce:
      LBRACE -> [condition]
    Goto:
      (nil)

  State 329:
    Kernel Items:
      switch_expr: optional_label_decl SWITCH sequence_expr.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 406

  State 330:
    Kernel Items:
      binary_and_expr: and_expr.AND cmp_expr
      binary_or_expr: or_expr OR and_expr., *
    Reduce:
      * -> [binary_or_expr]
    Goto:
      AND -> State 205

  State 331:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN explicit_enum_value_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 407

  State 332:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def.OR field_def
      proper_explicit_enum_value_defs: field_def.NEWLINES field_def
    Reduce:
      (nil)
    Goto:
      NEWLINES -> State 408
      OR -> State 409

  State 333:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs.OR field_def
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs.NEWLINES field_def
      explicit_enum_value_defs: proper_explicit_enum_value_defs., RPAREN
      explicit_enum_value_defs: proper_explicit_enum_value_defs.NEWLINES
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      NEWLINES -> State 410
      OR -> State 411

  State 334:
    Kernel Items:
      parameter_decl: ELLIPSIS., *
      parameter_decl: ELLIPSIS.type_expr
    Reduce:
      * -> [parameter_decl]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 412
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 335:
    Kernel Items:
      named_type_expr: IDENTIFIER.generic_type_arguments
      named_type_expr: IDENTIFIER.DOT IDENTIFIER generic_type_arguments
      proper_parameter_def: IDENTIFIER.type_expr
      proper_parameter_def: IDENTIFIER.ELLIPSIS type_expr
      proper_parameter_def: IDENTIFIER.ELLIPSIS
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 344
      QUESTION -> State 119
      EXCLAIM -> State 114
      DOLLAR_LBRACKET -> State 187
      ELLIPSIS -> State 363
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 364
      binary_type_expr -> State 124
      generic_type_arguments -> State 239
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 336:
    Kernel Items:
      inferred_type_expr: UNDERSCORE., *
      proper_parameter_def: UNDERSCORE.type_expr
      proper_parameter_def: UNDERSCORE.ELLIPSIS
      proper_parameter_def: UNDERSCORE.ELLIPSIS type_expr
    Reduce:
      * -> [inferred_type_expr]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      ELLIPSIS -> State 365
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 366
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 337:
    Kernel Items:
      proper_parameter_decls: parameter_decl., *
    Reduce:
      * -> [proper_parameter_decls]
    Goto:
      (nil)

  State 338:
    Kernel Items:
      func_type_expr: FUNC LPAREN parameter_decls.RPAREN return_type
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 413

  State 339:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls.COMMA parameter_decl
      parameter_decls: proper_parameter_decls., RPAREN
      parameter_decls: proper_parameter_decls.COMMA
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      COMMA -> State 414

  State 340:
    Kernel Items:
      parameter_decl: proper_parameter_def., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      (nil)

  State 341:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_decl: type_expr., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 342:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT IDENTIFIER.generic_type_arguments
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      DOLLAR_LBRACKET -> State 187
      generic_type_arguments -> State 415

  State 343:
    Kernel Items:
      method_signature: FUNC IDENTIFIER.LPAREN parameter_decls RPAREN return_type
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 416

  State 344:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT.IDENTIFIER generic_type_arguments
      inferred_type_expr: DOT., *
    Reduce:
      * -> [inferred_type_expr]
    Goto:
      IDENTIFIER -> State 342

  State 345:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      field_def: IDENTIFIER type_expr.optional_default
    Reduce:
      * -> [optional_default]
    Goto:
      ASSIGN -> State 353
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256
      optional_default -> State 417

  State 346:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      field_def: UNDERSCORE type_expr., *
    Reduce:
      * -> [field_def]
    Goto:
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 347:
    Kernel Items:
      proper_implicit_enum_value_defs: field_def OR.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 241
      UNDERSCORE -> State 242
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 240
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      unsafe_statement -> State 250
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      binary_type_expr -> State 124
      field_def -> State 418
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126
      method_signature -> State 246

  State 348:
    Kernel Items:
      implicit_enum_type_expr: LPAREN implicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [implicit_enum_type_expr]
    Goto:
      (nil)

  State 349:
    Kernel Items:
      implicit_struct_type_expr: LPAREN implicit_field_defs RPAREN., *
    Reduce:
      * -> [implicit_struct_type_expr]
    Goto:
      (nil)

  State 350:
    Kernel Items:
      implicit_enum_value_defs: proper_implicit_enum_value_defs NEWLINES., RPAREN
    Reduce:
      RPAREN -> [implicit_enum_value_defs]
    Goto:
      (nil)

  State 351:
    Kernel Items:
      proper_implicit_enum_value_defs: proper_implicit_enum_value_defs OR.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 241
      UNDERSCORE -> State 242
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 240
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      unsafe_statement -> State 250
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      binary_type_expr -> State 124
      field_def -> State 419
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126
      method_signature -> State 246

  State 352:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs COMMA.field_def
      implicit_field_defs: proper_implicit_field_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [implicit_field_defs]
    Goto:
      IDENTIFIER -> State 241
      UNDERSCORE -> State 242
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 240
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      unsafe_statement -> State 250
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      binary_type_expr -> State 124
      field_def -> State 420
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126
      method_signature -> State 246

  State 353:
    Kernel Items:
      optional_default: ASSIGN.DEFAULT
    Reduce:
      (nil)
    Goto:
      DEFAULT -> State 421

  State 354:
    Kernel Items:
      field_def: type_expr optional_default., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 355:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN explicit_field_defs.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 422

  State 356:
    Kernel Items:
      binary_type_expr: type_expr binary_type_op returnable_type_expr., *
    Reduce:
      * -> [binary_type_expr]
    Goto:
      (nil)

  State 357:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER ASSIGN expr., *
    Reduce:
      * -> [named_func_def]
    Goto:
      (nil)

  State 358:
    Kernel Items:
      generic_parameter_def: IDENTIFIER., *
      generic_parameter_def: IDENTIFIER.type_expr
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 423
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 359:
    Kernel Items:
      proper_generic_parameter_def_list: generic_parameter_def., *
    Reduce:
      * -> [proper_generic_parameter_def_list]
    Goto:
      (nil)

  State 360:
    Kernel Items:
      generic_parameter_defs: DOLLAR_LBRACKET generic_parameter_def_list.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 424

  State 361:
    Kernel Items:
      proper_generic_parameter_def_list: proper_generic_parameter_def_list.COMMA generic_parameter_def
      generic_parameter_def_list: proper_generic_parameter_def_list., RBRACKET
      generic_parameter_def_list: proper_generic_parameter_def_list.COMMA
    Reduce:
      RBRACKET -> [generic_parameter_def_list]
    Goto:
      COMMA -> State 425

  State 362:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 260
      UNDERSCORE -> State 261
      proper_parameter_def -> State 263
      parameter_def -> State 284
      proper_parameter_defs -> State 286
      parameter_defs -> State 426

  State 363:
    Kernel Items:
      proper_parameter_def: IDENTIFIER ELLIPSIS.type_expr
      proper_parameter_def: IDENTIFIER ELLIPSIS., *
    Reduce:
      * -> [proper_parameter_def]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 427
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 364:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_parameter_def: IDENTIFIER type_expr., *
    Reduce:
      * -> [proper_parameter_def]
    Goto:
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 365:
    Kernel Items:
      proper_parameter_def: UNDERSCORE ELLIPSIS., *
      proper_parameter_def: UNDERSCORE ELLIPSIS.type_expr
    Reduce:
      * -> [proper_parameter_def]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 428
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 366:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_parameter_def: UNDERSCORE type_expr., *
    Reduce:
      * -> [proper_parameter_def]
    Goto:
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 367:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN.IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 429

  State 368:
    Kernel Items:
      proper_statements: proper_statements NEWLINES statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 369:
    Kernel Items:
      proper_statements: proper_statements SEMICOLON statement., *
    Reduce:
      * -> [proper_statements]
    Goto:
      (nil)

  State 370:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER ASSIGN type_expr., *
    Reduce:
      * -> [type_def]
    Goto:
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 371:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER generic_parameter_defs type_expr., *
      type_def: TYPE IDENTIFIER generic_parameter_defs type_expr.IMPLEMENTS type_expr
    Reduce:
      * -> [type_def]
    Goto:
      IMPLEMENTS -> State 430
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 372:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN.var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 146
      UNDERSCORE -> State 148
      LPAREN -> State 147
      var_pattern -> State 431
      tuple_pattern -> State 149

  State 373:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA.field_var_pattern
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 272
      UNDERSCORE -> State 148
      LPAREN -> State 147
      ELLIPSIS -> State 271
      var_pattern -> State 275
      tuple_pattern -> State 149
      field_var_pattern -> State 432

  State 374:
    Kernel Items:
      tuple_pattern: LPAREN field_var_patterns RPAREN., *
    Reduce:
      * -> [tuple_pattern]
    Goto:
      (nil)

  State 375:
    Kernel Items:
      case_pattern: DOT IDENTIFIER implicit_struct_expr., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 376:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER.tuple_pattern
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 147
      tuple_pattern -> State 433

  State 377:
    Kernel Items:
      statement: CASE case_patterns COLON optional_simple_statement., $
    Reduce:
      $ -> [statement]
    Goto:
      (nil)

  State 378:
    Kernel Items:
      case_patterns: case_patterns COMMA case_pattern., *
    Reduce:
      * -> [case_patterns]
    Goto:
      (nil)

  State 379:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 435
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      return_type -> State 434
      func_type_expr -> State 126

  State 380:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs COMMA.parameter_def
      parameter_defs: proper_parameter_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 260
      UNDERSCORE -> State 261
      proper_parameter_def -> State 263
      parameter_def -> State 436

  State 381:
    Kernel Items:
      import_statement: IMPORT LPAREN import_clauses RPAREN., $
    Reduce:
      $ -> [import_statement]
    Goto:
      (nil)

  State 382:
    Kernel Items:
      proper_import_clauses: proper_import_clauses COMMA.import_clause
      import_clauses: proper_import_clauses COMMA., RPAREN
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      STRING_LITERAL -> State 164
      IDENTIFIER -> State 162
      UNDERSCORE -> State 165
      DOT -> State 161
      import_clause -> State 437

  State 383:
    Kernel Items:
      proper_import_clauses: proper_import_clauses NEWLINES.import_clause
      import_clauses: proper_import_clauses NEWLINES., RPAREN
    Reduce:
      RPAREN -> [import_clauses]
    Goto:
      STRING_LITERAL -> State 164
      IDENTIFIER -> State 162
      UNDERSCORE -> State 165
      DOT -> State 161
      import_clause -> State 438

  State 384:
    Kernel Items:
      map_type_expr: LBRACKET type_expr COLON type_expr.RBRACKET
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 439
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 385:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA INTEGER_LITERAL.RBRACKET
    Reduce:
      (nil)
    Goto:
      RBRACKET -> State 440

  State 386:
    Kernel Items:
      argument: IDENTIFIER ASSIGN expr., *
    Reduce:
      * -> [argument]
    Goto:
      (nil)

  State 387:
    Kernel Items:
      colon_expr: colon_expr COLON expr., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 388:
    Kernel Items:
      colon_expr: expr COLON expr., *
    Reduce:
      * -> [colon_expr]
    Goto:
      (nil)

  State 389:
    Kernel Items:
      proper_arguments: proper_arguments COMMA argument., *
    Reduce:
      * -> [proper_arguments]
    Goto:
      (nil)

  State 390:
    Kernel Items:
      explicit_struct_type_expr: STRUCT LPAREN explicit_field_defs RPAREN., *
    Reduce:
      * -> [explicit_struct_type_expr]
    Goto:
      (nil)

  State 391:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs COMMA.field_def
      explicit_field_defs: proper_explicit_field_defs COMMA., RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 241
      UNDERSCORE -> State 242
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 240
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      unsafe_statement -> State 250
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      binary_type_expr -> State 124
      field_def -> State 441
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126
      method_signature -> State 246

  State 392:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs NEWLINES.field_def
      explicit_field_defs: proper_explicit_field_defs NEWLINES., RPAREN
    Reduce:
      RPAREN -> [explicit_field_defs]
    Goto:
      IDENTIFIER -> State 241
      UNDERSCORE -> State 242
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 240
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      unsafe_statement -> State 250
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      binary_type_expr -> State 124
      field_def -> State 442
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126
      method_signature -> State 246

  State 393:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER.STRING_LITERAL
    Reduce:
      (nil)
    Goto:
      STRING_LITERAL -> State 443

  State 394:
    Kernel Items:
      proper_type_arguments: proper_type_arguments COMMA.type_expr
      type_arguments: proper_type_arguments COMMA., RBRACKET
    Reduce:
      RBRACKET -> [type_arguments]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 444
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 395:
    Kernel Items:
      generic_type_arguments: DOLLAR_LBRACKET type_arguments RBRACKET., *
    Reduce:
      * -> [generic_type_arguments]
    Goto:
      (nil)

  State 396:
    Kernel Items:
      index_expr: accessible_expr LBRACKET argument RBRACKET., *
    Reduce:
      * -> [index_expr]
    Goto:
      (nil)

  State 397:
    Kernel Items:
      call_expr: accessible_expr generic_type_arguments LPAREN arguments.RPAREN
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 445

  State 398:
    Kernel Items:
      initialize_expr: initializable_type_expr LPAREN arguments RPAREN., *
    Reduce:
      * -> [initialize_expr]
    Goto:
      (nil)

  State 399:
    Kernel Items:
      loop_expr_body: DO statement_block FOR.sequence_expr
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      optional_label_decl -> State 156
      sequence_expr -> State 446
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 400:
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      optional_label_decl -> State 156
      sequence_expr -> State 447
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 401:
    Kernel Items:
      loop_expr_body: FOR assign_pattern IN.sequence_expr DO statement_block
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      optional_label_decl -> State 156
      sequence_expr -> State 448
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 402:
    Kernel Items:
      loop_expr_body: FOR for_assignment SEMICOLON.optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      optional_label_decl -> State 156
      sequence_expr -> State 450
      optional_sequence_expr -> State 449
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 403:
    Kernel Items:
      loop_expr_body: FOR sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 451

  State 404:
    Kernel Items:
      case_patterns: case_patterns.COMMA case_pattern
      condition: CASE case_patterns.ASSIGN sequence_expr
    Reduce:
      (nil)
    Goto:
      COMMA -> State 281
      ASSIGN -> State 452

  State 405:
    Kernel Items:
      if_expr_body: IF condition statement_block., *
      if_expr_body: IF condition statement_block.ELSE statement_block
      if_expr_body: IF condition statement_block.ELSE if_expr
    Reduce:
      * -> [if_expr_body]
    Goto:
      ELSE -> State 453

  State 406:
    Kernel Items:
      switch_expr: optional_label_decl SWITCH sequence_expr statement_block., $
    Reduce:
      $ -> [switch_expr]
    Goto:
      (nil)

  State 407:
    Kernel Items:
      explicit_enum_type_expr: ENUM LPAREN explicit_enum_value_defs RPAREN., *
    Reduce:
      * -> [explicit_enum_type_expr]
    Goto:
      (nil)

  State 408:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def NEWLINES.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 241
      UNDERSCORE -> State 242
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 240
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      unsafe_statement -> State 250
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      binary_type_expr -> State 124
      field_def -> State 454
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126
      method_signature -> State 246

  State 409:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def OR.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 241
      UNDERSCORE -> State 242
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 240
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      unsafe_statement -> State 250
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      binary_type_expr -> State 124
      field_def -> State 455
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126
      method_signature -> State 246

  State 410:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs NEWLINES.field_def
      explicit_enum_value_defs: proper_explicit_enum_value_defs NEWLINES., RPAREN
    Reduce:
      RPAREN -> [explicit_enum_value_defs]
    Goto:
      IDENTIFIER -> State 241
      UNDERSCORE -> State 242
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 240
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      unsafe_statement -> State 250
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      binary_type_expr -> State 124
      field_def -> State 456
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126
      method_signature -> State 246

  State 411:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs OR.field_def
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 241
      UNDERSCORE -> State 242
      UNSAFE -> State 54
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 240
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      unsafe_statement -> State 250
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 249
      binary_type_expr -> State 124
      field_def -> State 457
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126
      method_signature -> State 246

  State 412:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      parameter_decl: ELLIPSIS type_expr., *
    Reduce:
      * -> [parameter_decl]
    Goto:
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 413:
    Kernel Items:
      func_type_expr: FUNC LPAREN parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 435
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      return_type -> State 458
      func_type_expr -> State 126

  State 414:
    Kernel Items:
      proper_parameter_decls: proper_parameter_decls COMMA.parameter_decl
      parameter_decls: proper_parameter_decls COMMA., RPAREN
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 335
      UNDERSCORE -> State 336
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      ELLIPSIS -> State 334
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 341
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      proper_parameter_def -> State 340
      parameter_decl -> State 459
      func_type_expr -> State 126

  State 415:
    Kernel Items:
      named_type_expr: IDENTIFIER DOT IDENTIFIER generic_type_arguments., *
    Reduce:
      * -> [named_type_expr]
    Goto:
      (nil)

  State 416:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN.parameter_decls RPAREN return_type
    Reduce:
      RPAREN -> [parameter_decls]
    Goto:
      IDENTIFIER -> State 335
      UNDERSCORE -> State 336
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      ELLIPSIS -> State 334
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 341
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      proper_parameter_def -> State 340
      parameter_decl -> State 337
      proper_parameter_decls -> State 339
      parameter_decls -> State 460
      func_type_expr -> State 126

  State 417:
    Kernel Items:
      field_def: IDENTIFIER type_expr optional_default., *
    Reduce:
      * -> [field_def]
    Goto:
      (nil)

  State 418:
    Kernel Items:
      proper_implicit_enum_value_defs: field_def OR field_def., *
    Reduce:
      * -> [proper_implicit_enum_value_defs]
    Goto:
      (nil)

  State 419:
    Kernel Items:
      proper_implicit_enum_value_defs: proper_implicit_enum_value_defs OR field_def., *
    Reduce:
      * -> [proper_implicit_enum_value_defs]
    Goto:
      (nil)

  State 420:
    Kernel Items:
      proper_implicit_field_defs: proper_implicit_field_defs COMMA field_def., *
    Reduce:
      * -> [proper_implicit_field_defs]
    Goto:
      (nil)

  State 421:
    Kernel Items:
      optional_default: ASSIGN DEFAULT., *
    Reduce:
      * -> [optional_default]
    Goto:
      (nil)

  State 422:
    Kernel Items:
      trait_type_expr: TRAIT LPAREN explicit_field_defs RPAREN., *
    Reduce:
      * -> [trait_type_expr]
    Goto:
      (nil)

  State 423:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      generic_parameter_def: IDENTIFIER type_expr., *
    Reduce:
      * -> [generic_parameter_def]
    Goto:
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 424:
    Kernel Items:
      generic_parameter_defs: DOLLAR_LBRACKET generic_parameter_def_list RBRACKET., *
    Reduce:
      * -> [generic_parameter_defs]
    Goto:
      (nil)

  State 425:
    Kernel Items:
      proper_generic_parameter_def_list: proper_generic_parameter_def_list COMMA.generic_parameter_def
      generic_parameter_def_list: proper_generic_parameter_def_list COMMA., RBRACKET
    Reduce:
      RBRACKET -> [generic_parameter_def_list]
    Goto:
      IDENTIFIER -> State 358
      generic_parameter_def -> State 461

  State 426:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 462

  State 427:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_parameter_def: IDENTIFIER ELLIPSIS type_expr., *
    Reduce:
      * -> [proper_parameter_def]
    Goto:
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 428:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_parameter_def: UNDERSCORE ELLIPSIS type_expr., *
    Reduce:
      * -> [proper_parameter_def]
    Goto:
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 429:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER.LPAREN parameter_defs RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      LPAREN -> State 463

  State 430:
    Kernel Items:
      type_def: TYPE IDENTIFIER generic_parameter_defs type_expr IMPLEMENTS.type_expr
    Reduce:
      (nil)
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 135
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      type_expr -> State 464
      binary_type_expr -> State 124
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      func_type_expr -> State 126

  State 431:
    Kernel Items:
      field_var_pattern: IDENTIFIER ASSIGN var_pattern., *
    Reduce:
      * -> [field_var_pattern]
    Goto:
      (nil)

  State 432:
    Kernel Items:
      field_var_patterns: field_var_patterns COMMA field_var_pattern., *
    Reduce:
      * -> [field_var_patterns]
    Goto:
      (nil)

  State 433:
    Kernel Items:
      case_pattern: VAR DOT IDENTIFIER tuple_pattern., *
    Reduce:
      * -> [case_pattern]
    Goto:
      (nil)

  State 434:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 465

  State 435:
    Kernel Items:
      return_type: returnable_type_expr., *
    Reduce:
      * -> [return_type]
    Goto:
      (nil)

  State 436:
    Kernel Items:
      proper_parameter_defs: proper_parameter_defs COMMA parameter_def., *
    Reduce:
      * -> [proper_parameter_defs]
    Goto:
      (nil)

  State 437:
    Kernel Items:
      proper_import_clauses: proper_import_clauses COMMA import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 438:
    Kernel Items:
      proper_import_clauses: proper_import_clauses NEWLINES import_clause., *
    Reduce:
      * -> [proper_import_clauses]
    Goto:
      (nil)

  State 439:
    Kernel Items:
      map_type_expr: LBRACKET type_expr COLON type_expr RBRACKET., *
    Reduce:
      * -> [map_type_expr]
    Goto:
      (nil)

  State 440:
    Kernel Items:
      array_type_expr: LBRACKET type_expr COMMA INTEGER_LITERAL RBRACKET., *
    Reduce:
      * -> [array_type_expr]
    Goto:
      (nil)

  State 441:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs COMMA field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 442:
    Kernel Items:
      proper_explicit_field_defs: proper_explicit_field_defs NEWLINES field_def., *
    Reduce:
      * -> [proper_explicit_field_defs]
    Goto:
      (nil)

  State 443:
    Kernel Items:
      unsafe_statement: UNSAFE LESS IDENTIFIER GREATER STRING_LITERAL., $
    Reduce:
      $ -> [unsafe_statement]
    Goto:
      (nil)

  State 444:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      proper_type_arguments: proper_type_arguments COMMA type_expr., *
    Reduce:
      * -> [proper_type_arguments]
    Goto:
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 445:
    Kernel Items:
      call_expr: accessible_expr generic_type_arguments LPAREN arguments RPAREN., *
    Reduce:
      * -> [call_expr]
    Goto:
      (nil)

  State 446:
    Kernel Items:
      loop_expr_body: DO statement_block FOR sequence_expr., $
    Reduce:
      $ -> [loop_expr_body]
    Goto:
      (nil)

  State 447:
    Kernel Items:
      for_assignment: assign_pattern ASSIGN sequence_expr., SEMICOLON
    Reduce:
      SEMICOLON -> [for_assignment]
    Goto:
      (nil)

  State 448:
    Kernel Items:
      loop_expr_body: FOR assign_pattern IN sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 466

  State 449:
    Kernel Items:
      loop_expr_body: FOR for_assignment SEMICOLON optional_sequence_expr.SEMICOLON optional_sequence_expr DO statement_block
    Reduce:
      (nil)
    Goto:
      SEMICOLON -> State 467

  State 450:
    Kernel Items:
      optional_sequence_expr: sequence_expr., DO
    Reduce:
      DO -> [optional_sequence_expr]
    Goto:
      (nil)

  State 451:
    Kernel Items:
      loop_expr_body: FOR sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr_body]
    Goto:
      (nil)

  State 452:
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      optional_label_decl -> State 156
      sequence_expr -> State 468
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      explicit_struct_type_expr -> State 75
      anonymous_func_expr -> State 59

  State 453:
    Kernel Items:
      if_expr_body: IF condition statement_block ELSE.statement_block
      if_expr_body: IF condition statement_block ELSE.if_expr
    Reduce:
      IF -> [optional_label_decl]
    Goto:
      LABEL_DECL -> State 41
      LBRACE -> State 11
      statement_block -> State 471
      optional_label_decl -> State 470
      if_expr -> State 469

  State 454:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def NEWLINES field_def., *
    Reduce:
      * -> [proper_explicit_enum_value_defs]
    Goto:
      (nil)

  State 455:
    Kernel Items:
      proper_explicit_enum_value_defs: field_def OR field_def., *
    Reduce:
      * -> [proper_explicit_enum_value_defs]
    Goto:
      (nil)

  State 456:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs NEWLINES field_def., *
    Reduce:
      * -> [proper_explicit_enum_value_defs]
    Goto:
      (nil)

  State 457:
    Kernel Items:
      proper_explicit_enum_value_defs: proper_explicit_enum_value_defs OR field_def., *
    Reduce:
      * -> [proper_explicit_enum_value_defs]
    Goto:
      (nil)

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
      RPAREN -> State 472

  State 461:
    Kernel Items:
      proper_generic_parameter_def_list: proper_generic_parameter_def_list COMMA generic_parameter_def., *
    Reduce:
      * -> [proper_generic_parameter_def_list]
    Goto:
      (nil)

  State 462:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs LPAREN parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 435
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      return_type -> State 473
      func_type_expr -> State 126

  State 463:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN.parameter_defs RPAREN return_type statement_block
    Reduce:
      RPAREN -> [parameter_defs]
    Goto:
      IDENTIFIER -> State 260
      UNDERSCORE -> State 261
      proper_parameter_def -> State 263
      parameter_def -> State 284
      proper_parameter_defs -> State 286
      parameter_defs -> State 474

  State 464:
    Kernel Items:
      binary_type_expr: type_expr.binary_type_op returnable_type_expr
      type_def: TYPE IDENTIFIER generic_parameter_defs type_expr IMPLEMENTS type_expr., *
    Reduce:
      * -> [type_def]
    Goto:
      ADD -> State 253
      SUB -> State 255
      MUL -> State 254
      binary_type_op -> State 256

  State 465:
    Kernel Items:
      anonymous_func_expr: FUNC LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [anonymous_func_expr]
    Goto:
      (nil)

  State 466:
    Kernel Items:
      loop_expr_body: FOR assign_pattern IN sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 475

  State 467:
    Kernel Items:
      loop_expr_body: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON.optional_sequence_expr DO statement_block
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
      var_decl_pattern -> State 107
      var_or_let -> State 24
      optional_label_decl -> State 156
      sequence_expr -> State 450
      optional_sequence_expr -> State 476
      call_expr -> State 71
      atom_expr -> State 63
      parse_error_expr -> State 95
      literal_expr -> State 88
      named_expr -> State 92
      block_expr -> State 70
      initialize_expr -> State 85
      implicit_struct_expr -> State 81
      accessible_expr -> State 108
      access_expr -> State 55
      index_expr -> State 83
      postfixable_expr -> State 97
      postfix_unary_expr -> State 96
      prefixable_expr -> State 100
      prefix_unary_expr -> State 98
      prefix_unary_op -> State 99
      mul_expr -> State 91
      binary_mul_expr -> State 67
      add_expr -> State 57
      binary_add_expr -> State 64
      cmp_expr -> State 74
      binary_cmp_expr -> State 66
      and_expr -> State 58
      binary_and_expr -> State 65
      or_expr -> State 94
      binary_or_expr -> State 69
      initializable_type_expr -> State 84
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
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
      if_expr_body: IF condition statement_block ELSE if_expr., $
    Reduce:
      $ -> [if_expr_body]
    Goto:
      (nil)

  State 470:
    Kernel Items:
      if_expr: optional_label_decl.if_expr_body
    Reduce:
      (nil)
    Goto:
      IF -> State 229
      if_expr_body -> State 231

  State 471:
    Kernel Items:
      if_expr_body: IF condition statement_block ELSE statement_block., $
    Reduce:
      $ -> [if_expr_body]
    Goto:
      (nil)

  State 472:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls RPAREN.return_type
    Reduce:
      * -> [return_type]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 435
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      return_type -> State 477
      func_type_expr -> State 126

  State 473:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 478

  State 474:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs.RPAREN return_type statement_block
    Reduce:
      (nil)
    Goto:
      RPAREN -> State 479

  State 475:
    Kernel Items:
      loop_expr_body: FOR assign_pattern IN sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr_body]
    Goto:
      (nil)

  State 476:
    Kernel Items:
      loop_expr_body: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr.DO statement_block
    Reduce:
      (nil)
    Goto:
      DO -> State 480

  State 477:
    Kernel Items:
      method_signature: FUNC IDENTIFIER LPAREN parameter_decls RPAREN return_type., *
    Reduce:
      * -> [method_signature]
    Goto:
      (nil)

  State 478:
    Kernel Items:
      named_func_def: FUNC IDENTIFIER generic_parameter_defs LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [named_func_def]
    Goto:
      (nil)

  State 479:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN.return_type statement_block
    Reduce:
      LBRACE -> [return_type]
    Goto:
      IDENTIFIER -> State 116
      UNDERSCORE -> State 122
      STRUCT -> State 50
      ENUM -> State 113
      TRAIT -> State 121
      FUNC -> State 115
      LPAREN -> State 117
      LBRACKET -> State 42
      DOT -> State 112
      QUESTION -> State 119
      EXCLAIM -> State 114
      TILDE_TILDE -> State 120
      BIT_NEG -> State 111
      BIT_AND -> State 110
      PARSE_ERROR -> State 118
      initializable_type_expr -> State 130
      slice_type_expr -> State 103
      array_type_expr -> State 60
      map_type_expr -> State 90
      atom_type_expr -> State 123
      named_type_expr -> State 131
      inferred_type_expr -> State 129
      parse_error_type_expr -> State 132
      returnable_type_expr -> State 435
      prefix_unary_type_expr -> State 133
      prefix_unary_type_op -> State 134
      implicit_struct_type_expr -> State 128
      explicit_struct_type_expr -> State 75
      trait_type_expr -> State 136
      implicit_enum_type_expr -> State 127
      explicit_enum_type_expr -> State 125
      return_type -> State 481
      func_type_expr -> State 126

  State 480:
    Kernel Items:
      loop_expr_body: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 482

  State 481:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type.statement_block
    Reduce:
      (nil)
    Goto:
      LBRACE -> State 11
      statement_block -> State 483

  State 482:
    Kernel Items:
      loop_expr_body: FOR for_assignment SEMICOLON optional_sequence_expr SEMICOLON optional_sequence_expr DO statement_block., $
    Reduce:
      $ -> [loop_expr_body]
    Goto:
      (nil)

  State 483:
    Kernel Items:
      named_func_def: FUNC LPAREN parameter_def RPAREN IDENTIFIER LPAREN parameter_defs RPAREN return_type statement_block., *
    Reduce:
      * -> [named_func_def]
    Goto:
      (nil)

Number of states: 483
Number of shift actions: 4598
Number of reduce actions: 434
Number of shift/reduce conflicts: 0
Number of reduce/reduce conflicts: 0
Number of unoptimized states: 5377
Number of unoptimized shift actions: 46989
Number of unoptimized reduce actions: 35825
*/
